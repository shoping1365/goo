<template>
  <client-only>
    <div class="p-6">
      <!-- مدیریت کاربران در بالای کارت‌ها -->
      <div class="w-full mb-4">
        <div class="bg-white rounded-xl shadow p-2 flex flex-row items-center gap-6 w-full justify-between" style="box-shadow: 0 2px 12px 0 rgba(0,0,0,0.06);">
          <div class="flex flex-row items-center gap-6 w-full">
            <div class="bg-pink-100 text-pink-700 px-6 py-2 rounded-xl font-bold text-xl flex items-center border border-pink-100" style="box-shadow: 0 2px 12px 0 rgba(236, 72, 153, 0.10);">
              مدیریت کاربران
            </div>
          </div>
          <!-- دکمه تنظیمات آستانه آنلاین/آفلاین -->
          <div class="relative">
            <button
              class="inline-flex items-center px-3 py-2 border border-gray-200 rounded-lg bg-white hover:bg-gray-50 transition-all ml-2 shadow-md"
              title="تنظیمات"
              @click="settingsOpen = !settingsOpen"
            >
              <svg class="w-5 h-5 text-gray-500" fill="none" viewBox="0 0 24 24">
                <circle cx="12" cy="12" r="4" stroke="currentColor" stroke-width="2" fill="none" />
                <path stroke="currentColor" stroke-width="2" fill="none" d="M4.93 4.93l2.12 2.12M2 12h3M4.93 19.07l2.12-2.12M12 22v-3M19.07 19.07l-2.12-2.12M22 12h-3M19.07 4.93l-2.12 2.12M12 2v3" />
              </svg>
            </button>
            <div v-if="settingsOpen" class="absolute left-0 mt-2 w-64 bg-white border border-gray-200 rounded-lg shadow-lg p-3 z-10">
              <div class="text-sm font-bold mb-2">تنظیم آستانه آنلاین بودن (ثانیه)</div>
              <div class="flex items-center gap-2">
                <input type="number" min="5" step="5" v-model.number="onlineThresholdSec" class="border border-gray-200 rounded px-2 py-1 w-24" />
                <button class="px-3 py-1 border border-gray-200 rounded bg-white hover:bg-gray-50" @click="saveThreshold">ذخیره</button>
              </div>
              <div class="text-xs text-gray-500 mt-2">مبنای محاسبه وضعیت آنلاین/آفلاین در لیست کاربران</div>
            </div>
          </div>
        </div>
      </div>

      <UserStatsCards 
        :users="users" 
        :filters="filters" 
        @update:filters="filters = $event"
        @export-excel="exportExcel"
      />

      <UserFilterBar :filters="filters" @update:filters="filters = $event" />

      <!-- باکس عملیات زیر فیلترها -->
      <div class="w-full my-2">
        <div class="bg-white rounded-xl shadow flex flex-row-reverse justify-between items-center w-full pr-4 py-2" style="box-shadow: 0 2px 12px 0 rgba(0,0,0,0.06);">
          <div>
            <button
              v-if="selectedUserIds.length > 0 && hasPermission('user.delete')"
              @click="deleteSelectedUsers"
              class="px-6 py-2 rounded-xl font-bold text-red-700 bg-red-100 border border-red-100 hover:bg-red-200 transition-colors text-base flex items-center"
              style="box-shadow: 0 2px 12px 0 rgba(239, 68, 68, 0.10);"
            >
              حذف کاربر
            </button>
          </div>
          <NewUserBtn @open-modal="adminRegisterModalOpen = true" />
        </div>
      </div>

      <div class="bg-white rounded-xl shadow overflow-hidden">
        <UserTable :users="pagedUsers" :filters="filters" @select-user="openUserDetail" @selectionChanged="selectedUserIds = $event" />
        <Pagination
          class="border-t"
          :current-page="currentPage"
          :total-pages="totalPages"
          :total="filteredUsers.length"
          :items-per-page="itemsPerPage"
          @page-changed="(p:number)=>{ currentPage=p }"
          @items-per-page-changed="(n:number)=>{ itemsPerPage=n; currentPage=1 }"
        />
      </div>
      <UserDetailDrawer :user="selectedUser" :open="detailOpen" @close="detailOpen = false" />
      <AdminRegisterModal
        :open="adminRegisterModalOpen"
        @close="adminRegisterModalOpen = false"
        @success="fetchUsers()"
      />
    </div>
  </client-only>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
declare const $fetch: <T = unknown>(url: string, options?: { method?: string }) => Promise<T>
</script>

<script setup lang="ts">
import axios from 'axios';
import { computed, onMounted, ref, watch } from 'vue';
import Pagination from '~/components/admin/common/Pagination.vue';
import { useAuth } from '~/composables/useAuth';
import AdminRegisterModal from './AdminRegisterModal.vue';
import NewUserBtn from './NewUserBtn.vue';
import UserDetailDrawer from './UserDetailDrawer.vue';
import UserFilterBar from './UserFilterBar.vue';
import UserStatsCards from './UserStatsCards.vue';
import UserTable from './UserTable.vue';

definePageMeta({ 
  layout: 'admin-main',
  middleware: 'admin' // طبق مستندات ADMIN_MIDDLEWARE_SECURITY.md
})

// Interface برای User - متناسب با UserTable component
interface User {
  id: number | string;
  ID?: number | string;
  name: string; // required
  fullName?: string;
  username?: string;
  email?: string;
  mobile?: string;
  phone?: string;
  mobileNumber?: string;
  mobile_number?: string;
  landline?: string;
  landlineNumber?: string;
  landline_number?: string;
  nationalCode?: string;
  national_code?: string;
  selectedAddress?: string;
  selected_address?: string;
  registered?: string;
  registeredAt?: string;
  registered_at?: string;
  lastSeen?: string;
  last_seen?: string;
  lastSeenAt?: string;
  last_seen_at?: string;
  online?: boolean;
  role?: string;
  role_name?: string;
  roleName?: string;
  roles?: string[];
  status?: string;
  twofa?: boolean;
}

// استفاده از کامپوزابل احراز هویت برای بررسی پرمیژن‌ها
const { hasPermission } = useAuth()

// User state management without Pinia
const users = ref<User[]>([]);
// آستانه آنلاین بودن بر حسب ثانیه (تنظیم‌پذیر از طریق دکمه چرخ‌دنده)
const settingsOpen = ref(false);
const onlineThresholdSec = ref<number>(120);
const loading = ref(false);
const filters = ref({
  status: '',
  role: '',
  onlineOnly: false
});
const selectedUser = ref(null);
const detailOpen = ref(false);
const adminRegisterModalOpen = ref(false);
const selectedUserIds = ref<number[]>([]);

// صفحه‌بندی (پیش‌فرض 10)
const currentPage = ref(1)
const itemsPerPage = ref(10)

// فیلتر کردن کاربران بر اساس فیلترهای اعمال شده
const filteredUsers = computed(() => {
  let filtered = users.value;
  
  // فیلتر آنلاین
  if (filters.value.onlineOnly) {
    filtered = filtered.filter(u => u.online);
  }
  
  // فیلتر وضعیت
  if (filters.value.status) {
    filtered = filtered.filter(u => u.status === filters.value.status);
  }
  
  // فیلتر نقش
  if (filters.value.role) {
    filtered = filtered.filter(u => u.role === filters.value.role || u.roleName === filters.value.role);
  }
  
  return filtered;
});

const totalPages = computed(() => Math.max(1, Math.ceil(filteredUsers.value.length / itemsPerPage.value)))
const pagedUsers = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage.value
  return filteredUsers.value.slice(start, start + itemsPerPage.value)
})

function openUserDetail(user: User) {
  fetchUserDetail(user.id);
}

async function fetchUserDetail(id: number | string) {
  try {
    const detail = await $fetch<User>(`/api/users/${id}`);
    
    selectedUser.value = {
      ...detail,
      id: detail.id || detail.ID || id,
      name: detail.name || detail.fullName || detail.username || '-',
      registered: detail.registered || detail.registeredAt || detail.registered_at || null,
      lastSeen: detail.lastSeen || detail.last_seen || detail.lastSeenAt || detail.last_seen_at || null,
      online: false,
      mobile: detail.mobile || detail.phone || detail.mobileNumber || detail.mobile_number || '-',
      landline: detail.landline || detail.landlineNumber || detail.landline_number || '-',
      nationalCode: detail.nationalCode || detail.national_code || '-',
      selectedAddress: detail.selectedAddress || detail.selected_address || '-',
    };
    detailOpen.value = true;
  } catch (err) {
    console.error('Error fetching user detail:', err);
    alert('خطا در دریافت اطلاعات کاربر');
  }
}

// محاسبه وضعیت آنلاین بودن کاربر بر اساس آخرین بازدید و آستانه تنظیم‌شده
function computeOnline(lastSeen: string | undefined | null, thresholdSec: number): boolean {
  // اگر مقدار زمان آخرین فعالیت وجود نداشته باشد، آفلاین در نظر گرفته می‌شود
  if (!lastSeen) return false;
  const lastSeenDate = new Date(lastSeen);
  const now = new Date();
  return (now.getTime() - lastSeenDate.getTime()) < thresholdSec * 1000;
}

// نرمال‌سازی تاریخ‌های برگشتی از بک‌اند (string | object(sql.NullTime) | number)
function normalizeDate(input: any): string | null {
  if (!input) return null;
  try {
    // اگر رشته ISO است
    if (typeof input === 'string') {
      const fixed = input.replace(/(\.\d{3})\d+(Z|[+\-]\d{2}:?\d{2})$/, '$1$2');
      const d = new Date(fixed);
      if (!Number.isNaN(d.getTime())) return d.toISOString();
      // تلاش ثانویه: جایگزینی فاصله بین تاریخ و زمان با T
      const alt = new Date(fixed.replace(' ', 'T'));
      if (!Number.isNaN(alt.getTime())) return alt.toISOString();
      return null;
    }
    // اگر شیء شبیه sql.NullTime باشد
    if (typeof input === 'object') {
      if (input.Time) {
        const d = new Date(input.Time);
        return Number.isNaN(d.getTime()) ? null : d.toISOString();
      }
      if (input.seconds) {
        const d = new Date(Number(input.seconds) * 1000);
        return Number.isNaN(d.getTime()) ? null : d.toISOString();
      }
    }
    // اگر عدد یونیکس میلی‌ثانیه باشد
    if (typeof input === 'number') {
      const d = new Date(input > 1e12 ? input : input * 1000);
      return Number.isNaN(d.getTime()) ? null : d.toISOString();
    }
  } catch {}
  return null;
}

// ذخیره آستانه در localStorage و بازمحاسبه وضعیت آنلاین
function saveThreshold() {
  try {
    localStorage.setItem('users_online_threshold_sec', String(onlineThresholdSec.value));
  } catch {}
  // اطلاع‌رسانی به سایر بخش‌ها (مانند پلاگین heartbeat) برای همگام‌سازی بازه ارسال ضربان
  try {
    window.dispatchEvent(new CustomEvent('users_online_threshold_changed', { detail: onlineThresholdSec.value }));
  } catch {}
  // بلافاصله یک ضربان سمت سرور ارسال کن تا اثر تغییر دیده شود
  try {
    fetch('/api/users/heartbeat', { method: 'PUT', credentials: 'same-origin' });
  } catch {}
  recomputeOnlineStatuses();
  settingsOpen.value = false;
}

// بازمحاسبه وضعیت آنلاین برای کل لیست فعلی کاربران
function recomputeOnlineStatuses() {
  users.value = users.value.map((u: any) => {
    const lastSeen = u.lastSeen || u.last_seen || u.lastSeenAt || u.last_seen_at;
    return { ...u, online: computeOnline(lastSeen, onlineThresholdSec.value) };
  });
}

async function fetchUsers() {
  loading.value = true;
  try {
    const response = await $fetch<User[]>('/api/users');
    console.log('RAW users from backend:', response);
    
    // Map users to ensure required fields exist
    users.value = response.map(u => {
      const lastSeenRaw = u.lastSeen || u.last_seen || u.lastSeenAt || u.last_seen_at;
      const registeredRaw = u.registered || u.registeredAt || u.registered_at;
      const lastSeen = normalizeDate(lastSeenRaw);
      const registered = normalizeDate(registeredRaw);
      const online = typeof u.online === 'boolean' ? u.online : computeOnline(lastSeen, onlineThresholdSec.value);
      
      return {
        ...u,
        id: u.id || u.ID || 0,
        name: u.name || u.fullName || u.username || '-', // name is required
        email: u.email || '-',
        roleName: u.role_name || u.roleName || (Array.isArray(u.roles) ? u.roles[0] : u.role) || undefined,
        role: u.role || u.role_name || u.roleName || (Array.isArray(u.roles) ? u.roles[0] : undefined),
        registered: registered || undefined,
        lastSeen: lastSeen || undefined,
        online,
        mobile: u.mobile || u.phone || u.mobileNumber || u.mobile_number || undefined,
        landline: u.landline || u.landlineNumber || u.landline_number || undefined,
        nationalCode: u.nationalCode || u.national_code || undefined,
        selectedAddress: u.selectedAddress || u.selected_address || undefined,
      } as User;
    });
    
    console.log('users after mapping:', users.value);
  } catch (error) {
    console.error('Failed to fetch users:', error);
  } finally {
    loading.value = false;
  }
}

async function deleteSelectedUsers() {
  if (!selectedUserIds.value.length) return;
  if (!confirm('آیا از حذف کاربران انتخاب شده مطمئن هستید؟')) return;
  try {
    await Promise.all(selectedUserIds.value.map(id => axios.delete(`/api/users/${id}`)));
    selectedUserIds.value = [];
    fetchUsers();
  } catch (err) {
    alert('خطا در حذف کاربران');
  }
}

async function exportExcel() {
  try {
    const response = await $fetch('/api/users/export', { method: 'GET' }) as Blob;
    // Handle Excel download
    const blob = new Blob([response], { type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet' });
    const url = window.URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = 'users.xlsx';
    a.click();
    window.URL.revokeObjectURL(url);
  } catch (error) {
    console.error('Export failed:', error);
  }
}

onMounted(() => {
  // خواندن مقدار ذخیره‌شده آستانه از localStorage (پیش‌فرض 120 ثانیه)
  try {
    const saved = localStorage.getItem('users_online_threshold_sec');
    if (saved) {
      const parsed = parseInt(saved, 10);
      if (!Number.isNaN(parsed) && parsed > 0) onlineThresholdSec.value = parsed;
    }
  } catch {}

  fetchUsers();

  // وقتی ضربان موفق ارسال شد، لیست کاربران را تازه کن تا lastSeen و وضعیت دقیق شود
  try {
    window.addEventListener('users:heartbeat-ok', () => {
      fetchUsers();
    });
  } catch {}
});

// بازمحاسبه خودکار در صورت تغییر آستانه (بدون نیاز به دکمه ذخیره)
watch(onlineThresholdSec, () => {
  recomputeOnlineStatuses();
});
</script>
