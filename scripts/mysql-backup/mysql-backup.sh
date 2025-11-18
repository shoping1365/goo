#!/usr/bin/env bash
set -euo pipefail
set +m

# ============================================================
# MySQL FTP Backup Utility (single-file edition)
# Configure everything in this section, then deploy the script.
# ============================================================

# Interval in hours between automatic backup runs when the script is
# launched in daemon mode (or via the systemd service).
BACKUP_INTERVAL_HOURS=${BACKUP_INTERVAL_HOURS:-2}

# Maximum number of concurrent mysqldump processes.
PARALLEL_JOBS=${PARALLEL_JOBS:-1}

# Worker threads for archive compression (requires 7z for multi-threading).
ZIP_THREADS=${ZIP_THREADS:-20}

# Retain log files for this many days.
LOG_RETENTION_DAYS=${LOG_RETENTION_DAYS:-15}

# Password used to encrypt ZIP archives. Leave empty to disable encryption.
ARCHIVE_PASSWORD=${ARCHIVE_PASSWORD:-"uS7^hL!xP9@rZ3*qD1#vM6&bT8\$wY0%kF"}

# Optional label to identify the host in logs/notifications.
SERVER_NAME=${SERVER_NAME:-$(hostname -f 2>/dev/null || hostname || "unknown-host")}

# Telegram notification settings.
TELEGRAM_BOT_TOKEN=${TELEGRAM_BOT_TOKEN:-"8028767427:AAE66PL0nEIlhfUHAovjzDRJoADHE77QwyY"}
TELEGRAM_CHAT_ID=${TELEGRAM_CHAT_ID:-"504690554"}
TELEGRAM_DISABLE_NOTIFICATIONS=${TELEGRAM_DISABLE_NOTIFICATIONS:-false}

# FTP destination information.
FTP_HOST="45.92.92.3"
FTP_PORT=${FTP_PORT:-21}
FTP_USER="mohamma2"
FTP_PASS="z@DPmzQFT9tN"
FTP_BASE_DIR="/home/mohamma2/database/myserver"

# Where logs should be stored on the server running the script.
LOG_DIR=${LOG_DIR:-/var/log/mysql-backup}

# Location to copy/install the script when using --install-service.
INSTALL_PATH=${INSTALL_PATH:-/opt/mysql-backup/mysql-backup.sh}

# Name of the systemd unit managed by --install-service/--remove-service.
SYSTEMD_UNIT_NAME=${SYSTEMD_UNIT_NAME:-mysql-backup.service}

# Define databases as multi-line entries for readability.
# Required keys per entry: NAME, USER, PASSWORD.
# Optional keys: HOST (default localhost), REMOTE_DIR (default NAME).
#
# ⚠️ امنیت: ذخیره password در کد منبع خطرناک است!
#
# ✅ راه حل صحیح:
# 1. از environment variables استفاده کنید: PASSWORD=${DB_PASSWORD:-"default"}
# 2. از secret management systems استفاده کنید (مثل HashiCorp Vault, AWS Secrets Manager)
# 3. از فایل‌های config خارج از repository استفاده کنید که در .gitignore هستند
# 4. هرگز password ها را در کد منبع commit نکنید
#
# مثال صحیح:
# DATABASES=(
#   "NAME=my_database
#    USER=my_user
#    PASSWORD=${MY_DB_PASSWORD}
#    HOST=localhost"
# )
#
# سپس در environment variables تنظیم کنید:
# export MY_DB_PASSWORD="secure_password_here"
#
DATABASES=(
  "NAME=abanwtr_database
USER=abanwtr_user
PASSWORD=r}i*i*2dHdMdbP3d
HOST=localhost
REMOTE_DIR=abanwtr"

  "NAME=banooyekermani_database
USER=banooyekermani_user
PASSWORD=KZH#hP=sN%3}RVYm
HOST=localhost
REMOTE_DIR=banooyekermani"

  "NAME=esfahantours_database
USER=esfahantours_user
PASSWORD=WoT{p*W7ne7E_3Fu}~
HOST=localhost
REMOTE_DIR=esfahantours"

  "NAME=iranxia_database
USER=iranxia_user
PASSWORD=UW!{EgO)8-#y7A}Bcs
HOST=localhost
REMOTE_DIR=iranxia"

  "NAME=jirjor_database
USER=jirjor_user
PASSWORD=wy*}6U%gLyfz+3T9DB
HOST=localhost
REMOTE_DIR=jirjor"

  "NAME=kimiakesht_database
USER=kimiakesht_user
PASSWORD=,~oBU+R[~2hv
HOST=localhost
REMOTE_DIR=kimiakesht"

  "NAME=miteck_database
USER=miteck_user
PASSWORD=[g2~cMR,QWMMnsK8sP
HOST=localhost
REMOTE_DIR=miteck"

  "NAME=rahesaba_database
USER=rahesaba_user
PASSWORD=OWavz!meUu!m)Q*}
HOST=localhost
REMOTE_DIR=rahesaba"

  "NAME=ravanekhoob_database
USER=ravanekhoob_user
PASSWORD=CF=RrGXK?z0OXN+u
HOST=localhost
REMOTE_DIR=ravanekhoob"
)

# ============================================================

SCRIPT_PATH=$(readlink -f "$0" 2>/dev/null || realpath "$0")
SCRIPT_DIR=$(cd "$(dirname "$SCRIPT_PATH")" && pwd)
LOG_FILE=""

ensure_positive_integer() {
  local name="$1"
  local value="$2"
  if ! [[ "$value" =~ ^[0-9]+$ ]] || (( value < 1 )); then
    echo "Invalid value for $name: $value" >&2
    exit 1
  fi
}

check_prereqs() {
  local missing=()
  for cmd in mysqldump curl mktemp zip; do
    if ! command -v "$cmd" >/dev/null 2>&1; then
      missing+=("$cmd")
    fi
  done
  if (( ${#missing[@]} > 0 )); then
    echo "Missing required commands: ${missing[*]}" >&2
    exit 1
  fi
}

log() {
  local message="$1"
  local ts
  ts=$(date +"%Y-%m-%d %H:%M:%S")
  if [[ -n "$LOG_FILE" ]]; then
    echo "[$ts] $message" >> "$LOG_FILE"
  fi
}

log_and_print() {
  local message="$1"
  local ts
  ts=$(date +"%Y-%m-%d %H:%M:%S")
  echo "$message"
  if [[ -n "$LOG_FILE" ]]; then
    echo "[$ts] $message" >> "$LOG_FILE"
  fi
}

telegram_notify() {
  local message="$1"
  [[ -z "$TELEGRAM_BOT_TOKEN" || -z "$TELEGRAM_CHAT_ID" ]] && return 0

  local url="https://api.telegram.org/bot${TELEGRAM_BOT_TOKEN}/sendMessage"
  local decorated="[${SERVER_NAME}] $message"
  
  local -a payload=(
    --silent --show-error --fail --max-time 10
    --data "chat_id=${TELEGRAM_CHAT_ID}"
    --data-urlencode "text=${decorated}"
  )
  [[ "$TELEGRAM_DISABLE_NOTIFICATIONS" == "true" ]] && payload+=(--data "disable_notification=true")
  
  # Send silently to telegram, don't show in terminal
  curl "${payload[@]}" "$url" >>"${LOG_FILE:-/dev/null}" 2>&1 || true
  return 0
}

trim_spaces() {
  local value="$1"
  value="${value#${value%%[![:space:]]*}}"
  value="${value%${value##*[![:space:]]}}"
  printf '%s' "$value"
}

parse_entry_field() {
  local entry="$1"
  local wanted_key="$2"
  local key value
  while IFS='=' read -r key value; do
    [[ -z "$key" ]] && continue
    key=$(trim_spaces "$key")
    [[ -z "$key" || ${key:0:1} == "#" ]] && continue
    value=$(trim_spaces "$value")
    if [[ "$key" == "$wanted_key" ]]; then
      printf '%s' "$value"
      return 0
    fi
  done <<< "$entry"
  return 1
}

build_remote_url() {
  local remote_dir="$1"
  local file_name="$2"

  local base_dir="/${FTP_BASE_DIR#/}"
  if [[ -n "$remote_dir" ]]; then
    base_dir="${base_dir%/}/$(echo "$remote_dir" | sed 's#^/*##')"
  fi
  echo "ftp://$FTP_HOST:${FTP_PORT:-21}${base_dir}/${file_name}"
}

emit_progress_line() {
  local pipe="$1"
  shift
  local message="$*"
  if [[ -n "$pipe" && -p "$pipe" ]]; then
    printf '%s\n' "$message" >>"$pipe"
  else
    printf '%s\n' "$message"
  fi
}

cleanup_progress_channel() {
  local pipe="$1"
  local reader_pid="$2"
  if [[ -n "$reader_pid" ]]; then
    wait "$reader_pid" 2>/dev/null || true
  fi
  if [[ -n "$pipe" ]]; then
    rm -f "$pipe" 2>/dev/null || true
  fi
}

filesize_bytes() {
  local target="$1"
  if command -v stat >/dev/null 2>&1; then
    local size
    if size=$(stat -c%s "$target" 2>/dev/null); then
      printf '%s' "$size"
      return 0
    fi
    if size=$(stat -f%z "$target" 2>/dev/null); then
      printf '%s' "$size"
      return 0
    fi
  fi
  wc -c <"$target" 2>/dev/null || echo 0
}

now_ns() {
  local ts
  ts=$(date +%s%N 2>/dev/null || true)
  if [[ "$ts" =~ ^[0-9]+$ ]]; then
    printf '%s' "$ts"
  else
    printf '%s' $(( $(date +%s) * 1000000000 ))
  fi
}

elapsed_seconds() {
  local start_ns="$1"
  local end_ns="$2"
  awk -v s="$start_ns" -v e="$end_ns" 'BEGIN {
    if (e <= s) {
      printf "0.000"
    } else {
      printf "%.3f", (e - s) / 1000000000
    }
  }'
}

calc_speed_mib_per_s() {
  local bytes="$1"
  local seconds="$2"
  awk -v b="$bytes" -v s="$seconds" 'BEGIN {
    if (s <= 0 || b < 0) {
      printf "0.00"
    } else {
      printf "%.2f", (b / 1048576) / s
    }
  }'
}

run_single_backup() {
  local entry="$1"
  local slot_number="${2:-0}"
  local db_name db_user db_pass db_host remote_dir
  local safe_name ts dump_file archive_name archive_file remote_url
  local dump_start_ns dump_end_ns dump_duration_sec
  local archive_start_ns archive_end_ns archive_duration_sec
  local upload_start_ns upload_end_ns upload_duration_sec
  local job_temp_dir db_version

  dump_duration_sec="0.000"
  archive_duration_sec="0.000"
  upload_duration_sec="0.000"

  db_name=$(parse_entry_field "$entry" NAME || true)
  db_user=$(parse_entry_field "$entry" USER || true)
  db_pass=$(parse_entry_field "$entry" PASSWORD || true)
  db_host=$(parse_entry_field "$entry" HOST || true)
  remote_dir=$(parse_entry_field "$entry" REMOTE_DIR || true)

  db_name=${db_name:-}
  db_user=${db_user:-}
  db_pass=${db_pass:-}
  db_host=${db_host:-localhost}
  remote_dir=${remote_dir:-$db_name}

  if [[ -z "$db_name" || -z "$db_user" || -z "$db_pass" ]]; then
    log "Skipping entry with missing required fields (NAME/USER/PASSWORD)"
    return 1
  fi

  # Get MariaDB/MySQL version
  db_version=$(mysql -h "$db_host" -u "$db_user" -p"$db_pass" -sN -e "SELECT VERSION();" 2>/dev/null | awk '{print $1}' | tr -d '\n' || echo "unknown")
  db_version=$(echo "$db_version" | sed 's/[^0-9.]//g')

  # Create a unique temp directory for this job
  job_temp_dir=$(mktemp -d)
  trap "rm -rf '$job_temp_dir'" EXIT

  safe_name=$(echo "$db_name" | tr ' ' '_')
  ts=$(date +"%Y-%m-%d_%H-%M-%S")
  dump_file="$job_temp_dir/${safe_name}_${ts}.sql"
  archive_name="${safe_name}_${ts}_mariadb-${db_version}.zip"
  archive_file="$job_temp_dir/$archive_name"
  remote_url=$(build_remote_url "$remote_dir" "$archive_name")
  telegram_notify "🚀 شروع بک‌آپ برای ${db_name}"

  log "Dumping $db_name (host=$db_host user=$db_user) [Slot $slot_number]"
  log_and_print "[$slot_number] → Starting backup: ${db_name} (MariaDB ${db_version})"
  dump_start_ns=$(now_ns)

  local dump_exit_code=0
  timeout 300 mysqldump \
    -h "$db_host" \
    -u "$db_user" \
    -p"$db_pass" \
    --single-transaction \
    --routines \
    --triggers \
    --events \
    --add-drop-table \
    --skip-lock-tables \
    --max-allowed-packet=512M \
    --net-buffer-length=16384 \
    "$db_name" 2>>"$LOG_FILE" >"$dump_file"
  dump_exit_code=$?

  if (( dump_exit_code == 0 )); then
    dump_end_ns=$(now_ns)
    dump_duration_sec=$(elapsed_seconds "$dump_start_ns" "$dump_end_ns")
    local dump_size_human=$(du -h "$dump_file" | awk '{print $1}')
    log_and_print "[$slot_number] ✓ Dump complete: ${db_name} | Time: ${dump_duration_sec}s | Size: ${dump_size_human}"
  else
    log_and_print "[$slot_number] ✗ Dump FAILED: $db_name"
    telegram_notify "❌ بک‌آپ ${db_name} در مرحله Dump شکست خورد"
    rm -f "$dump_file"
    return 1
  fi

  log "Archiving $archive_name [Slot $slot_number]"
  log_and_print "[$slot_number] → Compressing ${db_name}..."
  archive_start_ns=$(now_ns)
  local archive_exit_code=0

  local dump_size_bytes
  dump_size_bytes=$(filesize_bytes "$dump_file")
  dump_size_bytes=${dump_size_bytes:-0}

  # Use standard zip command with maximum compression
  if [[ -n "$ARCHIVE_PASSWORD" ]]; then
    zip -9 -j -P "$ARCHIVE_PASSWORD" "$archive_file" "$dump_file" >>"$LOG_FILE" 2>&1
    archive_exit_code=$?
  else
    zip -9 -j "$archive_file" "$dump_file" >>"$LOG_FILE" 2>&1
    archive_exit_code=$?
  fi

  if (( archive_exit_code == 0 )); then
    archive_end_ns=$(now_ns)
    archive_duration_sec=$(elapsed_seconds "$archive_start_ns" "$archive_end_ns")
    local archive_size_human=$(du -h "$archive_file" | awk '{print $1}')
    local archive_size_bytes=$(filesize_bytes "$archive_file")
    local compression_speed=$(calc_speed_mib_per_s "$archive_size_bytes" "$archive_duration_sec")
    log_and_print "[$slot_number] ✓ Compression complete: ${db_name} | Time: ${archive_duration_sec}s | Size: ${archive_size_human} | Speed: ${compression_speed} MiB/s"
  else
    log_and_print "[$slot_number] ✗ Compression FAILED: $db_name"
    telegram_notify "❌ بک‌آپ ${db_name} در مرحله ZIP شکست خورد"
    rm -f "$dump_file" "$archive_file"
    return 1
  fi

  rm -f "$dump_file"

  log "Uploading to $remote_url [Slot $slot_number]"
  log_and_print "[$slot_number] → Uploading ${db_name} to FTP..."
  upload_start_ns=$(now_ns)
  local upload_exit_code=0
  local archive_size_bytes
  archive_size_bytes=$(filesize_bytes "$archive_file")
  archive_size_bytes=${archive_size_bytes:-0}

  curl --silent --show-error --ftp-create-dirs \
    --user "$FTP_USER:$FTP_PASS" \
    -T "$archive_file" \
    "$remote_url" >>"$LOG_FILE" 2>&1
  upload_exit_code=$?

  if (( upload_exit_code == 0 )); then
    upload_end_ns=$(now_ns)
    upload_duration_sec=$(elapsed_seconds "$upload_start_ns" "$upload_end_ns")
    local upload_speed_mib
    upload_speed_mib=$(calc_speed_mib_per_s "$archive_size_bytes" "$upload_duration_sec")
    log_and_print "[$slot_number] ✓ Upload complete: ${db_name} | Time: ${upload_duration_sec}s | Speed: ${upload_speed_mib} MiB/s"
    log_and_print "[$slot_number] ═══════════════════════════════════════════════════"
    rm -f "$archive_file"
    local remote_path
    remote_path="${remote_url#ftp://$FTP_HOST:${FTP_PORT:-21}}"
    telegram_notify "✅ بک‌آپ ${db_name} با موفقیت پایان یافت\n⏱ Dump: ${dump_duration_sec}s | ZIP: ${archive_duration_sec}s | Upload: ${upload_duration_sec}s\n🚀 سرعت آپلود: ${upload_speed_mib} MiB/s\n📁 ${remote_path}"
  else
    log_and_print "[$slot_number] ✗ Upload FAILED: $db_name (file kept: $archive_file)"
    telegram_notify "❌ بک‌آپ ${db_name} در مرحله Upload شکست خورد"
    return 1
  fi
}

run_all_backups() {
  check_prereqs
  ensure_positive_integer "PARALLEL_JOBS" "$PARALLEL_JOBS"
  ensure_positive_integer "LOG_RETENTION_DAYS" "$LOG_RETENTION_DAYS"
  ensure_positive_integer "ZIP_THREADS" "$ZIP_THREADS"

  mkdir -p "$LOG_DIR"
  LOG_FILE="$LOG_DIR/backup_$(date +"%Y-%m-%d_%H-%M-%S").log"

  if (( ZIP_THREADS > 1 )) && ! command -v 7z >/dev/null 2>&1; then
    log "ZIP_THREADS=$ZIP_THREADS set but 7z not found; falling back to single-thread zip"
  fi

  log "Starting backup run on $SERVER_NAME (parallel $PARALLEL_JOBS, zip threads $ZIP_THREADS)"
  log_and_print "═══════════════════════════════════════════════════"
  log_and_print "MySQL Backup Started"
  log_and_print "Server: $SERVER_NAME"
  log_and_print "Parallel Jobs: $PARALLEL_JOBS | Compression Threads: $ZIP_THREADS"
  log_and_print "Total Databases: ${#DATABASES[@]}"
  log_and_print "═══════════════════════════════════════════════════"
  local run_start run_end run_duration
  run_start=$(date +%s)
  local total_databases=${#DATABASES[@]}
  
  telegram_notify "🚀 اجرای بک‌آپ آغاز شد (${total_databases} دیتابیس)"

  if (( total_databases == 0 )); then
    log "No databases defined; nothing to do"
    telegram_notify "ℹ️ هیچ دیتابیسی برای بک‌آپ تعریف نشده بود"
    return 0
  fi

  local -a pids=()
  local -i failures=0
  local -i success_count=0
  local -i current_slot=1
  declare -A job_names=()
  declare -A job_start=()
  declare -A job_slots=()

  for entry in "${DATABASES[@]}"; do
    local db_name
    db_name=$(parse_entry_field "$entry" NAME || echo "(unknown)")
    log "Queued $db_name for backup [Slot $current_slot]"

    run_single_backup "$entry" "$current_slot" &
    local pid=$!
    pids+=($pid)
    job_names[$pid]="$db_name"
    job_start[$pid]=$(date +%s)
    job_slots[$pid]=$current_slot

    if (( PARALLEL_JOBS > 0 && ${#pids[@]} >= PARALLEL_JOBS )); then
      log "Reached parallel limit ($PARALLEL_JOBS); waiting for a job to finish"
    fi

    while (( PARALLEL_JOBS > 0 && ${#pids[@]} >= PARALLEL_JOBS )); do
      local wait_start wait_end wait_duration
      wait_start=$(date +%s)
      
      # Wait for any job to finish (-n option)
      local finished_pid
      wait -n -p finished_pid "${pids[@]}"
      local exit_code=$?
      
      wait_end=$(date +%s)
      wait_duration=$((wait_end - wait_start))
      
      if (( exit_code == 0 )); then
        log "Job ${job_names[$finished_pid]} [Slot ${job_slots[$finished_pid]}] finished successfully after $((wait_end - job_start[$finished_pid]))s (waited ${wait_duration}s for slot)"
      else
        log "Job ${job_names[$finished_pid]} [Slot ${job_slots[$finished_pid]}] failed after $((wait_end - job_start[$finished_pid]))s (waited ${wait_duration}s for slot)"
      fi
      
      # Remove finished job from arrays
      unset job_names[$finished_pid]
      unset job_start[$finished_pid]
      unset job_slots[$finished_pid]
      
      # Remove finished PID from pids array
      local new_pids=()
      for p in "${pids[@]}"; do
        [[ "$p" != "$finished_pid" ]] && new_pids+=("$p")
      done
      pids=("${new_pids[@]}")
    done

    ((current_slot++))
  done

  for pid in "${pids[@]}"; do
    local wait_end total_duration
    if wait "$pid"; then
      wait_end=$(date +%s)
      total_duration=$((wait_end - job_start[$pid]))
      log "Job ${job_names[$pid]} [Slot ${job_slots[$pid]}] finished successfully after ${total_duration}s"
      ((success_count++))
    else
      wait_end=$(date +%s)
      total_duration=$((wait_end - job_start[$pid]))
      log "Job ${job_names[$pid]} [Slot ${job_slots[$pid]}] failed after ${total_duration}s"
      ((failures++))
    fi
    unset job_names[$pid]
    unset job_start[$pid]
    unset job_slots[$pid]
  done

  run_end=$(date +%s)
  run_duration=$((run_end - run_start))

  log "Run completed in ${run_duration}s with $failures failure(s)"
  log_and_print "═══════════════════════════════════════════════════"
  log_and_print "Backup Run Completed"
  log_and_print "Total Time: ${run_duration}s"
  log_and_print "Success: ${success_count} | Failed: ${failures}"
  log_and_print "═══════════════════════════════════════════════════"
  
  telegram_notify "📊 اجرای بک‌آپ تمام شد در ${run_duration}s\n✅ موفق: ${success_count}\n❌ ناموفق: ${failures}"
  
  find "$LOG_DIR" -type f -name "backup_*.log" -mtime +"$LOG_RETENTION_DAYS" -delete || true

  if (( failures > 0 )); then
    return 1
  fi
  return 0
}

run_daemon() {
  ensure_positive_integer "BACKUP_INTERVAL_HOURS" "$BACKUP_INTERVAL_HOURS"
  while true; do
    if ! run_all_backups; then
      log "One or more backups failed"
    fi
    sleep "$((BACKUP_INTERVAL_HOURS * 3600))"
  done
}

require_root() {
  if (( EUID != 0 )); then
    echo "This action requires root privileges." >&2
    exit 1
  fi
}

install_service() {
  require_root
  local target_path="${1:-$INSTALL_PATH}"
  local target_dir
  target_dir=$(dirname "$target_path")

  mkdir -p "$target_dir"
  install -m 700 "$SCRIPT_PATH" "$target_path"

  local unit_path="/etc/systemd/system/$SYSTEMD_UNIT_NAME"
  cat >"$unit_path" <<EOF
[Unit]
Description=MySQL FTP Backup Service
After=network-online.target
Wants=network-online.target

[Service]
Type=simple
ExecStart=$target_path --daemon
Restart=always
RestartSec=30
User=root
Group=root
WorkingDirectory=$(dirname "$target_path")

[Install]
WantedBy=multi-user.target
EOF

  systemctl daemon-reload
  systemctl enable --now "$SYSTEMD_UNIT_NAME"
  echo "Service installed and started: $SYSTEMD_UNIT_NAME"
}

remove_service() {
  require_root
  local unit_path="/etc/systemd/system/$SYSTEMD_UNIT_NAME"
  if [[ -f "$unit_path" ]]; then
    systemctl disable --now "$SYSTEMD_UNIT_NAME" || true
    rm -f "$unit_path"
    systemctl daemon-reload
    echo "Service removed: $SYSTEMD_UNIT_NAME"
  else
    echo "Service unit not found: $unit_path" >&2
    exit 1
  fi
}

show_help() {
  cat <<EOF
Usage: $(basename "$0") [COMMAND]

Commands:
  --once                Run a single backup cycle (default).
  --daemon              Run continuously, sleeping BACKUP_INTERVAL_HOURS between runs.
  --install-service [path]
                        Copy the script to [path] (default: $INSTALL_PATH) and install
                        a systemd service named $SYSTEMD_UNIT_NAME.
  --remove-service      Disable and remove the systemd service.
  --help                Show this help message.

Configuration is defined at the top of the script. Environment variables with
the same names can override defaults at runtime.
EOF
}

main() {
  local cmd="${1:---once}"
  case "$cmd" in
    --once)
      run_all_backups
      ;;
    --daemon)
      run_daemon
      ;;
    --install-service)
      install_service "${2:-}"
      ;;
    --remove-service)
      remove_service
      ;;
    --help|-h)
      show_help
      ;;
    *)
      echo "Unknown command: $cmd" >&2
      show_help
      exit 1
      ;;
  esac
}

main "$@"
