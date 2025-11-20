<template>
  <div class="backup-restore">
    <div class="section-header">
      <h3>پشتیبان‌گیری و بازیابی نظرسنجی‌ها</h3>
      <p>مدیریت پشتیبان‌گیری و بازیابی داده‌های نظرسنجی‌ها</p>
    </div>

    <!-- Backup Status Overview -->
    <div class="backup-status-overview">
      <div class="status-card last-backup">
        <div class="status-icon">
          <i class="fas fa-database"></i>
        </div>
        <div class="status-content">
          <h4>آخرین پشتیبان‌گیری</h4>
          <div class="date">{{ formatDateTime(lastBackup.date) }}</div>
          <div class="status-text">{{ lastBackup.status }}</div>
        </div>
      </div>

      <div class="status-card backup-size">
        <div class="status-icon">
          <i class="fas fa-hdd"></i>
        </div>
        <div class="status-content">
          <h4>اندازه پشتیبان</h4>
          <div class="size">{{ formatFileSize(lastBackup.size) }}</div>
          <div class="status-text">{{ lastBackup.records }} رکورد</div>
        </div>
      </div>

      <div class="status-card auto-backup">
        <div class="status-icon">
          <i class="fas fa-clock"></i>
        </div>
        <div class="status-content">
          <h4>پشتیبان‌گیری خودکار</h4>
          <div class="status">{{ autoBackup.enabled ? 'فعال' : 'غیرفعال' }}</div>
          <div class="status-text">{{ autoBackup.frequency }}</div>
        </div>
      </div>

      <div class="status-card storage-usage">
        <div class="status-icon">
          <i class="fas fa-chart-pie"></i>
        </div>
        <div class="status-content">
          <h4>استفاده از فضای ذخیره</h4>
          <div class="usage">{{ storageUsage.used }} / {{ storageUsage.total }}</div>
          <div class="progress-bar">
            <div class="progress-fill" :style="{ width: storageUsage.percentage + '%' }"></div>
          </div>
        </div>
      </div>
    </div>

    <!-- Backup Actions -->
    <div class="backup-actions">
      <h4>عملیات پشتیبان‌گیری</h4>
      <div class="actions-grid">
        <div class="action-card">
          <div class="action-header">
            <i class="fas fa-download"></i>
            <h5>ایجاد پشتیبان</h5>
          </div>
          <p>ایجاد پشتیبان کامل از تمام داده‌های نظرسنجی‌ها</p>
          <div class="action-options">
            <label class="checkbox-label">
              <input v-model="backupOptions.includeFiles" type="checkbox">
              <span>شامل فایل‌های پیوست</span>
            </label>
            <label class="checkbox-label">
              <input v-model="backupOptions.compress" type="checkbox">
              <span>فشرده‌سازی</span>
            </label>
            <label class="checkbox-label">
              <input v-model="backupOptions.encrypt" type="checkbox">
              <span>رمزگذاری</span>
            </label>
          </div>
          <button class="btn btn-primary" :disabled="isBackingUp" @click="createBackup">
            <i v-if="isBackingUp" class="fas fa-spinner fa-spin"></i>
            <i v-else class="fas fa-download"></i>
            {{ isBackingUp ? 'در حال پشتیبان‌گیری...' : 'ایجاد پشتیبان' }}
          </button>
        </div>

        <div class="action-card">
          <div class="action-header">
            <i class="fas fa-upload"></i>
            <h5>بازیابی</h5>
          </div>
          <p>بازیابی داده‌ها از فایل پشتیبان</p>
          <div class="file-upload">
            <input ref="restoreFile" type="file" accept=".zip,.sql,.json" style="display: none;" @change="handleFileSelect">
            <button class="btn btn-secondary" @click="restoreFile?.click()">
              <i class="fas fa-folder-open"></i>
              انتخاب فایل
            </button>
            <div v-if="selectedFile" class="selected-file">
              <i class="fas fa-file"></i>
              <span>{{ selectedFile.name }}</span>
              <button class="btn-clear" @click="clearSelectedFile">
                <i class="fas fa-times"></i>
              </button>
            </div>
          </div>
          <button class="btn btn-warning" :disabled="!selectedFile || isRestoring" @click="restoreBackup">
            <i v-if="isRestoring" class="fas fa-spinner fa-spin"></i>
            <i v-else class="fas fa-upload"></i>
            {{ isRestoring ? 'در حال بازیابی...' : 'بازیابی' }}
          </button>
        </div>

        <div class="action-card">
          <div class="action-header">
            <i class="fas fa-cog"></i>
            <h5>تنظیمات خودکار</h5>
          </div>
          <p>تنظیم پشتیبان‌گیری خودکار</p>
          <div class="auto-settings">
            <label class="toggle-switch">
              <input v-model="autoBackup.enabled" type="checkbox">
              <span class="slider"></span>
              <span class="label">پشتیبان‌گیری خودکار</span>
            </label>
            <div class="form-group">
              <label>فرکانس:</label>
              <select v-model="autoBackup.frequency" :disabled="!autoBackup.enabled">
                <option value="daily">روزانه</option>
                <option value="weekly">هفتگی</option>
                <option value="monthly">ماهانه</option>
              </select>
            </div>
            <div class="form-group">
              <label>زمان:</label>
              <input v-model="autoBackup.time" type="time" :disabled="!autoBackup.enabled">
            </div>
          </div>
          <button class="btn btn-success" @click="saveAutoSettings">
            <i class="fas fa-save"></i>
            ذخیره تنظیمات
          </button>
        </div>
      </div>
    </div>

    <!-- Backup History -->
    <div class="backup-history">
      <h4>تاریخچه پشتیبان‌گیری</h4>
      <div class="history-controls">
        <div class="search-box">
          <input v-model="searchQuery" type="text" placeholder="جستجو در تاریخچه...">
          <i class="fas fa-search"></i>
        </div>
        <div class="filter-controls">
          <select v-model="statusFilter">
            <option value="all">همه وضعیت‌ها</option>
            <option value="success">موفق</option>
            <option value="failed">ناموفق</option>
            <option value="in_progress">در حال انجام</option>
          </select>
        </div>
      </div>

      <div class="history-table">
        <table>
          <thead>
            <tr>
              <th>تاریخ</th>
              <th>نوع</th>
              <th>اندازه</th>
              <th>وضعیت</th>
              <th>توضیحات</th>
              <th>عملیات</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="backup in filteredBackups" :key="backup.id">
              <td>{{ formatDateTime(backup.date) }}</td>
              <td>
                <span class="backup-type" :class="backup.type">
                  {{ getBackupTypeText(backup.type) }}
                </span>
              </td>
              <td>{{ formatFileSize(backup.size) }}</td>
              <td>
                <span :class="['status-badge', backup.status]">
                  {{ getStatusText(backup.status) }}
                </span>
              </td>
              <td>{{ backup.description }}</td>
              <td>
                <div class="action-buttons">
                  <button class="btn-icon" title="دانلود" @click="downloadBackup(backup)">
                    <i class="fas fa-download"></i>
                  </button>
                  <button v-if="hasPermission('backup.delete')" class="btn-icon" title="حذف" @click="deleteBackup(backup.id)">
                    <i class="fas fa-trash"></i>
                  </button>
                  <button class="btn-icon" title="جزئیات" @click="viewBackupDetails(backup)">
                    <i class="fas fa-eye"></i>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Storage Management -->
    <div class="storage-management">
      <h4>مدیریت فضای ذخیره</h4>
      <div class="storage-info">
        <div class="storage-item">
          <div class="storage-label">فضای کل:</div>
          <div class="storage-value">{{ formatFileSize(storageUsage.total) }}</div>
        </div>
        <div class="storage-item">
          <div class="storage-label">فضای استفاده شده:</div>
          <div class="storage-value">{{ formatFileSize(storageUsage.used) }}</div>
        </div>
        <div class="storage-item">
          <div class="storage-label">فضای آزاد:</div>
          <div class="storage-value">{{ formatFileSize(storageUsage.free) }}</div>
        </div>
        <div class="storage-item">
          <div class="storage-label">درصد استفاده:</div>
          <div class="storage-value">{{ storageUsage.percentage }}%</div>
        </div>
      </div>
      
      <div class="cleanup-actions">
        <button class="btn btn-danger" @click="cleanupOldBackups">
          <i class="fas fa-broom"></i>
          پاکسازی پشتیبان‌های قدیمی
        </button>
        <button class="btn btn-info" @click="optimizeStorage">
          <i class="fas fa-compress"></i>
          بهینه‌سازی فضای ذخیره
        </button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const useAuth: () => { user: { id?: number; name?: string; email?: string } | null; hasPermission: (perm: string) => boolean }
</script>

<script setup lang="ts">
import { computed, onMounted, reactive, ref } from 'vue';

// استفاده از کامپوزابل احراز هویت
useAuth()
const restoreFile = ref<HTMLInputElement | null>(null)


// Reactive data
const isBackingUp = ref(false)
const isRestoring = ref(false)
const selectedFile = ref(null)
const searchQuery = ref('')
const statusFilter = ref('all')

const lastBackup = reactive({
  date: new Date(Date.now() - 86400000), // 1 day ago
  status: 'موفق',
  size: 52428800, // 50MB
  records: 15420
})

const autoBackup = reactive({
  enabled: true,
  frequency: 'daily',
  time: '02:00'
})

const storageUsage = reactive({
  total: 107374182400, // 100GB
  used: 32212254720, // 30GB
  free: 75161927680, // 70GB
  percentage: 30
})

const backupOptions = reactive({
  includeFiles: true,
  compress: true,
  encrypt: false
})

const backupHistory = ref([
  {
    id: 1,
    date: new Date(Date.now() - 86400000),
    type: 'full',
    size: 52428800,
    status: 'success',
    description: 'پشتیبان‌گیری کامل روزانه'
  },
  {
    id: 2,
    date: new Date(Date.now() - 172800000),
    type: 'incremental',
    size: 10485760,
    status: 'success',
    description: 'پشتیبان‌گیری افزایشی'
  },
  {
    id: 3,
    date: new Date(Date.now() - 259200000),
    type: 'full',
    size: 52428800,
    status: 'failed',
    description: 'خطا در پشتیبان‌گیری'
  }
])

// Computed properties
const filteredBackups = computed(() => {
  let backups = backupHistory.value
  
  if (searchQuery.value) {
    backups = backups.filter(backup => 
      backup.description.includes(searchQuery.value) ||
      getBackupTypeText(backup.type).includes(searchQuery.value)
    )
  }
  
  if (statusFilter.value !== 'all') {
    backups = backups.filter(backup => backup.status === statusFilter.value)
  }
  
  return backups
})

// Methods
function formatDateTime(date) {
  return new Intl.DateTimeFormat('fa-IR', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  }).format(date)
}

function formatFileSize(bytes) {
  if (bytes === 0) return '0 B'
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB', 'TB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

function getBackupTypeText(type) {
  const types = {
    full: 'کامل',
    incremental: 'افزایشی',
    differential: 'تفاضلی'
  }
  return types[type] || type
}

function getStatusText(status) {
  const statuses = {
    success: 'موفق',
    failed: 'ناموفق',
    in_progress: 'در حال انجام'
  }
  return statuses[status] || status
}

function createBackup() {
  isBackingUp.value = true
  // console.log('Creating backup with options:', backupOptions)
  
  // Simulate backup process
  setTimeout(() => {
    isBackingUp.value = false
    // console.log('Backup completed successfully')
  }, 3000)
}

function handleFileSelect(event) {
  const file = event.target.files[0]
  if (file) {
    selectedFile.value = file
  }
}

function clearSelectedFile() {
  selectedFile.value = null
  if (restoreFile.value) {
    restoreFile.value.value = ''
  }
}

function restoreBackup() {
  if (!selectedFile.value) return
  
  isRestoring.value = true
  // console.log('Restoring from file:', selectedFile.value.name)
  
  // Simulate restore process
  setTimeout(() => {
    isRestoring.value = false
    clearSelectedFile()
    // console.log('Restore completed successfully')
  }, 5000)
}

function saveAutoSettings() {
  // console.log('Saving auto backup settings:', autoBackup)
}

function downloadBackup(backup) {
  // console.log('Downloading backup:', backup.id)
}

function deleteBackup(id) {
  if (confirm('آیا از حذف این پشتیبان اطمینان دارید؟')) {
    // console.log('Deleting backup:', id)
  }
}

function viewBackupDetails(backup) {
  // console.log('Viewing backup details:', backup)
}

function cleanupOldBackups() {
  if (confirm('آیا از پاکسازی پشتیبان‌های قدیمی اطمینان دارید؟')) {
    // console.log('Cleaning up old backups')
  }
}

function optimizeStorage() {
  // console.log('Optimizing storage')
}

onMounted(() => {
  // console.log('BackupRestore component mounted')
})
</script>

<style scoped>
.backup-restore {
  direction: rtl;
  padding: 1rem;
}

.section-header {
  margin-bottom: 2rem;
}

.section-header h3 {
  margin: 0 0 0.5rem 0;
  font-size: 1.5rem;
  color: #333;
}

.section-header p {
  margin: 0;
  color: #666;
}

.backup-status-overview {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1rem;
  margin-bottom: 2rem;
}

.status-card {
  background: white;
  border-radius: 12px;
  padding: 1.5rem;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  display: flex;
  align-items: center;
  gap: 1rem;
}

.status-icon {
  width: 50px;
  height: 50px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.5rem;
  background: #e3f2fd;
  color: #1976d2;
}

.status-content h4 {
  margin: 0 0 0.5rem 0;
  font-size: 1rem;
  color: #333;
}

.status-content .date,
.status-content .size,
.status-content .status,
.status-content .usage {
  font-size: 1.2rem;
  font-weight: bold;
  color: #333;
  margin-bottom: 0.25rem;
}

.status-content .status-text {
  font-size: 0.9rem;
  color: #666;
}

.progress-bar {
  width: 100%;
  height: 8px;
  background: #e9ecef;
  border-radius: 4px;
  overflow: hidden;
  margin-top: 0.5rem;
}

.progress-fill {
  height: 100%;
  background: #28a745;
  transition: width 0.3s ease;
}

.backup-actions {
  margin-bottom: 2rem;
}

.backup-actions h4 {
  margin: 0 0 1rem 0;
  font-size: 1.2rem;
  color: #333;
}

.actions-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 1rem;
}

.action-card {
  background: white;
  border-radius: 12px;
  padding: 1.5rem;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.action-header {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin-bottom: 1rem;
}

.action-header i {
  font-size: 1.5rem;
  color: #007bff;
}

.action-header h5 {
  margin: 0;
  font-size: 1.1rem;
  color: #333;
}

.action-card p {
  margin: 0 0 1rem 0;
  color: #666;
}

.action-options {
  margin-bottom: 1rem;
}

.checkbox-label {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin-bottom: 0.5rem;
  cursor: pointer;
}

.checkbox-label input[type="checkbox"] {
  width: 16px;
  height: 16px;
}

.file-upload {
  margin-bottom: 1rem;
}

.selected-file {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin-top: 0.5rem;
  padding: 0.5rem;
  background: #f8f9fa;
  border-radius: 6px;
}

.btn-clear {
  background: none;
  border: none;
  color: #dc3545;
  cursor: pointer;
  padding: 0.25rem;
}

.auto-settings {
  margin-bottom: 1rem;
}

.toggle-switch {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin-bottom: 1rem;
  cursor: pointer;
}

.toggle-switch input {
  display: none;
}

.slider {
  position: relative;
  width: 50px;
  height: 24px;
  background: #ccc;
  border-radius: 12px;
  transition: 0.3s;
}

.slider:before {
  content: '';
  position: absolute;
  width: 18px;
  height: 18px;
  left: 3px;
  bottom: 3px;
  background: white;
  border-radius: 50%;
  transition: 0.3s;
}

.toggle-switch input:checked + .slider {
  background: #28a745;
}

.toggle-switch input:checked + .slider:before {
  transform: translateX(26px);
}

.form-group {
  margin-bottom: 1rem;
}

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: 500;
  color: #555;
}

.form-group input,
.form-group select {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 0.9rem;
}

.btn {
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 0.9rem;
  display: flex;
  align-items: center;
  gap: 0.5rem;
  transition: all 0.2s;
}

.btn-primary {
  background: #007bff;
  color: white;
}

.btn-secondary {
  background: #6c757d;
  color: white;
}

.btn-warning {
  background: #ffc107;
  color: #212529;
}

.btn-success {
  background: #28a745;
  color: white;
}

.btn-danger {
  background: #dc3545;
  color: white;
}

.btn-info {
  background: #17a2b8;
  color: white;
}

.btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.backup-history {
  margin-bottom: 2rem;
}

.backup-history h4 {
  margin: 0 0 1rem 0;
  font-size: 1.2rem;
  color: #333;
}

.history-controls {
  display: flex;
  gap: 1rem;
  margin-bottom: 1rem;
  flex-wrap: wrap;
}

.search-box {
  position: relative;
  flex: 1;
  min-width: 250px;
}

.search-box input {
  width: 100%;
  padding: 0.75rem 2.5rem 0.75rem 0.75rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 0.9rem;
}

.search-box i {
  position: absolute;
  right: 0.75rem;
  top: 50%;
  transform: translateY(-50%);
  color: #666;
}

.filter-controls select {
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 0.9rem;
}

.history-table {
  background: white;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.history-table table {
  width: 100%;
  border-collapse: collapse;
}

.history-table th,
.history-table td {
  padding: 1rem;
  text-align: right;
  border-bottom: 1px solid #e9ecef;
}

.history-table th {
  background: #f8f9fa;
  font-weight: 600;
  color: #555;
}

.backup-type {
  padding: 0.25rem 0.5rem;
  border-radius: 12px;
  font-size: 0.8rem;
}

.backup-type.full {
  background: #d4edda;
  color: #155724;
}

.backup-type.incremental {
  background: #cce7ff;
  color: #004085;
}

.backup-type.differential {
  background: #fff3cd;
  color: #856404;
}

.status-badge {
  padding: 0.25rem 0.5rem;
  border-radius: 12px;
  font-size: 0.8rem;
}

.status-badge.success {
  background: #d4edda;
  color: #155724;
}

.status-badge.failed {
  background: #f8d7da;
  color: #721c24;
}

.status-badge.in_progress {
  background: #fff3cd;
  color: #856404;
}

.action-buttons {
  display: flex;
  gap: 0.5rem;
}

.btn-icon {
  background: none;
  border: none;
  color: #007bff;
  cursor: pointer;
  padding: 0.25rem;
  border-radius: 4px;
  transition: background 0.2s;
}

.btn-icon:hover {
  background: #f8f9fa;
}

.storage-management {
  background: white;
  border-radius: 12px;
  padding: 1.5rem;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.storage-management h4 {
  margin: 0 0 1rem 0;
  font-size: 1.2rem;
  color: #333;
}

.storage-info {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 1rem;
  margin-bottom: 1rem;
}

.storage-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0.75rem;
  background: #f8f9fa;
  border-radius: 6px;
}

.storage-label {
  font-weight: 500;
  color: #555;
}

.storage-value {
  font-weight: bold;
  color: #333;
}

.cleanup-actions {
  display: flex;
  gap: 1rem;
  flex-wrap: wrap;
}

@media (max-width: 768px) {
  .backup-status-overview {
    grid-template-columns: 1fr;
  }
  
  .actions-grid {
    grid-template-columns: 1fr;
  }
  
  .history-controls {
    flex-direction: column;
  }
  
  .search-box {
    min-width: auto;
  }
  
  .history-table {
    font-size: 0.8rem;
  }
  
  .history-table th,
  .history-table td {
    padding: 0.5rem;
  }
  
  .storage-info {
    grid-template-columns: 1fr;
  }
  
  .cleanup-actions {
    flex-direction: column;
  }
}
</style> 
