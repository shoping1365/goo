<template>
  <div v-if="hasAccess" class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-6">
      <h3 class="text-lg font-semibold text-gray-900">تنظیمات سیستم</h3>
      <button class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700" @click="saveAllSettings">
        ذخیره همه تنظیمات
      </button>
    </div>

    <!-- General Settings -->
    <div class="mb-8">
      <h4 class="text-md font-semibold text-gray-900 mb-4">تنظیمات کلی</h4>
      <div class="space-y-4">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداقل مبلغ قابل اقساط (ریال)</label>
            <input v-model="settings.general.minAmount" type="number" class="w-full border border-gray-300 rounded-md px-3 py-2">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر مبلغ قابل اقساط (ریال)</label>
            <input v-model="settings.general.maxAmount" type="number" class="w-full border border-gray-300 rounded-md px-3 py-2">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداقل تعداد اقساط</label>
            <input v-model="settings.general.minInstallments" type="number" class="w-full border border-gray-300 rounded-md px-3 py-2">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر تعداد اقساط</label>
            <input v-model="settings.general.maxInstallments" type="number" class="w-full border border-gray-300 rounded-md px-3 py-2">
          </div>
        </div>
        
        <div class="flex items-center">
          <input v-model="settings.general.autoApproval" type="checkbox" class="rounded border-gray-300">
          <label class="mr-2 text-sm font-medium text-gray-700">تایید خودکار درخواست‌ها</label>
        </div>
        
        <div class="flex items-center">
          <input v-model="settings.general.requireCreditCheck" type="checkbox" class="rounded border-gray-300">
          <label class="mr-2 text-sm font-medium text-gray-700">الزامی بودن بررسی اعتباری</label>
        </div>
      </div>
    </div>

    <!-- Interest Settings -->
    <div class="mb-8">
      <h4 class="text-md font-semibold text-gray-900 mb-4">تنظیمات سود</h4>
      <div class="space-y-4">
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نرخ سود پایه (%)</label>
            <input v-model="settings.interest.baseRate" type="number" step="0.1" class="w-full border border-gray-300 rounded-md px-3 py-2">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نرخ سود ریسک بالا (%)</label>
            <input v-model="settings.interest.highRiskRate" type="number" step="0.1" class="w-full border border-gray-300 rounded-md px-3 py-2">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نرخ سود VIP (%)</label>
            <input v-model="settings.interest.vipRate" type="number" step="0.1" class="w-full border border-gray-300 rounded-md px-3 py-2">
          </div>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نحوه محاسبه سود</label>
          <select v-model="settings.interest.calculationMethod" class="w-full md:w-1/2 border border-gray-300 rounded-md px-3 py-2">
            <option value="simple">سود ساده</option>
            <option value="compound">سود مرکب</option>
            <option value="declining">کاهنده</option>
          </select>
        </div>
      </div>
    </div>

    <!-- Notification Settings -->
    <div class="mb-8">
      <h4 class="text-md font-semibold text-gray-900 mb-4">تنظیمات اطلاع‌رسانی</h4>
      <div class="space-y-4">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="flex items-center">
              <input v-model="settings.notifications.smsEnabled" type="checkbox" class="rounded border-gray-300">
              <span class="mr-2 text-sm font-medium text-gray-700">ارسال پیامک</span>
            </label>
          </div>
          <div>
            <label class="flex items-center">
              <input v-model="settings.notifications.emailEnabled" type="checkbox" class="rounded border-gray-300">
              <span class="mr-2 text-sm font-medium text-gray-700">ارسال ایمیل</span>
            </label>
          </div>
          <div>
            <label class="flex items-center">
              <input v-model="settings.notifications.pushEnabled" type="checkbox" class="rounded border-gray-300">
              <span class="mr-2 text-sm font-medium text-gray-700">اعلان‌های Push</span>
            </label>
          </div>
          <div>
            <label class="flex items-center">
              <input v-model="settings.notifications.adminNotifications" type="checkbox" class="rounded border-gray-300">
              <span class="mr-2 text-sm font-medium text-gray-700">اطلاع‌رسانی به مدیران</span>
            </label>
          </div>
        </div>
        
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">روزهای قبل از سررسید برای یادآوری</label>
            <input v-model="settings.notifications.reminderDays" type="number" class="w-full border border-gray-300 rounded-md px-3 py-2">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">ساعت ارسال یادآوری</label>
            <input v-model="settings.notifications.reminderTime" type="time" class="w-full border border-gray-300 rounded-md px-3 py-2">
          </div>
        </div>
      </div>
    </div>

    <!-- Security Settings -->
    <div class="mb-8">
      <h4 class="text-md font-semibold text-gray-900 mb-4">تنظیمات امنیتی</h4>
      <div class="space-y-4">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر تلاش ورود ناموفق</label>
            <input v-model="settings.security.maxLoginAttempts" type="number" class="w-full border border-gray-300 rounded-md px-3 py-2">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مدت قفل شدن حساب (دقیقه)</label>
            <input v-model="settings.security.lockoutDuration" type="number" class="w-full border border-gray-300 rounded-md px-3 py-2">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مدت اعتبار کد تایید (دقیقه)</label>
            <input v-model="settings.security.otpValidityMinutes" type="number" class="w-full border border-gray-300 rounded-md px-3 py-2">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">سطح رمزنگاری</label>
            <select v-model="settings.security.encryptionLevel" class="w-full border border-gray-300 rounded-md px-3 py-2">
              <option value="basic">پایه</option>
              <option value="standard">استاندارد</option>
              <option value="high">بالا</option>
            </select>
          </div>
        </div>
        
        <div class="space-y-2">
          <label class="flex items-center">
            <input v-model="settings.security.twoFactorAuth" type="checkbox" class="rounded border-gray-300">
            <span class="mr-2 text-sm font-medium text-gray-700">احراز هویت دو مرحله‌ای</span>
          </label>
          <label class="flex items-center">
            <input v-model="settings.security.ipWhitelist" type="checkbox" class="rounded border-gray-300">
            <span class="mr-2 text-sm font-medium text-gray-700">محدودیت IP</span>
          </label>
          <label class="flex items-center">
            <input v-model="settings.security.sessionTimeout" type="checkbox" class="rounded border-gray-300">
            <span class="mr-2 text-sm font-medium text-gray-700">انقضای خودکار جلسه</span>
          </label>
        </div>
      </div>
    </div>

    <!-- Backup Settings -->
    <div class="mb-8">
      <h4 class="text-md font-semibold text-gray-900 mb-4">تنظیمات پشتیبان‌گیری</h4>
      <div class="space-y-4">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">دوره پشتیبان‌گیری خودکار</label>
            <select v-model="settings.backup.autoBackupInterval" class="w-full border border-gray-300 rounded-md px-3 py-2">
              <option value="daily">روزانه</option>
              <option value="weekly">هفتگی</option>
              <option value="monthly">ماهانه</option>
              <option value="disabled">غیرفعال</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تعداد فایل‌های پشتیبان نگهداری شده</label>
            <input v-model="settings.backup.retentionCount" type="number" class="w-full border border-gray-300 rounded-md px-3 py-2">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مسیر ذخیره پشتیبان</label>
            <input v-model="settings.backup.backupPath" type="text" class="w-full border border-gray-300 rounded-md px-3 py-2">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">فرمت فایل پشتیبان</label>
            <select v-model="settings.backup.backupFormat" class="w-full border border-gray-300 rounded-md px-3 py-2">
              <option value="sql">SQL</option>
              <option value="json">JSON</option>
              <option value="csv">CSV</option>
            </select>
          </div>
        </div>
        
        <div class="flex items-center space-x-4 space-x-reverse">
          <button class="bg-green-600 text-white px-4 py-2 rounded-md hover:bg-green-700" @click="createBackup">
            ایجاد پشتیبان فوری
          </button>
          <button class="bg-yellow-600 text-white px-4 py-2 rounded-md hover:bg-yellow-700" @click="restoreBackup">
            بازیابی از پشتیبان
          </button>
        </div>
      </div>
    </div>

    <!-- API Settings -->
    <div class="mb-8">
      <h4 class="text-md font-semibold text-gray-900 mb-4">تنظیمات API</h4>
      <div class="space-y-4">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حد درخواست در ساعت</label>
            <input v-model="settings.api.rateLimit" type="number" class="w-full border border-gray-300 rounded-md px-3 py-2">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">زمان انتظار (ثانیه)</label>
            <input v-model="settings.api.timeout" type="number" class="w-full border border-gray-300 rounded-md px-3 py-2">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نسخه API</label>
            <select v-model="settings.api.version" class="w-full border border-gray-300 rounded-md px-3 py-2">
              <option value="v1">ورژن 1</option>
              <option value="v2">ورژن 2</option>
              <option value="v3">ورژن 3 (آزمایشی)</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">فرمت پاسخ پیش‌فرض</label>
            <select v-model="settings.api.defaultFormat" class="w-full border border-gray-300 rounded-md px-3 py-2">
              <option value="json">JSON</option>
              <option value="xml">XML</option>
            </select>
          </div>
        </div>
        
        <div class="space-y-2">
          <label class="flex items-center">
            <input v-model="settings.api.enableLogging" type="checkbox" class="rounded border-gray-300">
            <span class="mr-2 text-sm font-medium text-gray-700">فعال‌سازی لاگ API</span>
          </label>
          <label class="flex items-center">
            <input v-model="settings.api.enableCors" type="checkbox" class="rounded border-gray-300">
            <span class="mr-2 text-sm font-medium text-gray-700">فعال‌سازی CORS</span>
          </label>
          <label class="flex items-center">
            <input v-model="settings.api.requireApiKey" type="checkbox" class="rounded border-gray-300">
            <span class="mr-2 text-sm font-medium text-gray-700">الزامی بودن API Key</span>
          </label>
        </div>
      </div>
    </div>

    <!-- Current System Status -->
    <div class="bg-blue-50 rounded-lg p-6">
      <h4 class="text-md font-semibold text-gray-900 mb-3">وضعیت فعلی سیستم</h4>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6 text-sm">
        <div>
          <span class="text-gray-600 block">درخواست‌های فعال</span>
          <span class="font-medium">{{ systemStatus.activeRequests }}</span>
        </div>
        <div>
          <span class="text-gray-600 block">آخرین پشتیبان‌گیری</span>
          <span class="font-medium">{{ systemStatus.lastBackup }}</span>
        </div>
        <div>
          <span class="text-gray-600 block">وضعیت سیستم</span>
          <span class="font-medium text-green-600">{{ systemStatus.systemHealth }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuth } from '~/composables/useAuth'

// احراز هویت
const { user, isAuthenticated } = useAuth()

// بررسی دسترسی admin
const hasAccess = computed(() => {
  if (!isAuthenticated.value) {
    return false
  }

  const userRole = user.value?.role?.toLowerCase() || ''
  const adminRoles = ['admin', 'developer']
  return adminRoles.includes(userRole)
})

const router = useRouter()

const settings = ref({
  general: {
    minAmount: 5000000,
    maxAmount: 500000000,
    minInstallments: 3,
    maxInstallments: 36,
    autoApproval: false,
    requireCreditCheck: true
  },
  interest: {
    baseRate: 18,
    highRiskRate: 24,
    vipRate: 15,
    calculationMethod: 'declining'
  },
  notifications: {
    smsEnabled: true,
    emailEnabled: true,
    pushEnabled: false,
    adminNotifications: true,
    reminderDays: 3,
    reminderTime: '09:00'
  },
  security: {
    maxLoginAttempts: 3,
    lockoutDuration: 30,
    otpValidityMinutes: 5,
    encryptionLevel: 'standard',
    twoFactorAuth: true,
    ipWhitelist: false,
    sessionTimeout: true
  },
  backup: {
    autoBackupInterval: 'daily',
    retentionCount: 7,
    backupPath: '/backups/installments',
    backupFormat: 'sql'
  },
  api: {
    rateLimit: 1000,
    timeout: 30,
    version: 'v2',
    defaultFormat: 'json',
    enableLogging: true,
    enableCors: true,
    requireApiKey: true
  }
})

const systemStatus = ref({
  activeRequests: 156,
  lastBackup: '1403/08/15 02:00',
  systemHealth: 'سالم'
})

const saveAllSettings = async () => {
  try {
    await $fetch('/api/admin/installment-payments/settings', {
      method: 'PUT',
      body: settings.value
    })
    alert('تمام تنظیمات با موفقیت ذخیره شد')
  } catch (error) {
    console.error('خطا در ذخیره تنظیمات:', error)
    alert('خطا در ذخیره تنظیمات')
  }
}

const createBackup = async () => {
  try {
    interface ApiResponse {
      success?: boolean
      message?: string
    }
    const response = await $fetch<ApiResponse>('/api/admin/installment-payments/backup/create', {
      method: 'POST'
    })
    
    if (response.success) {
      alert('پشتیبان‌گیری با موفقیت انجام شد')
      systemStatus.value.lastBackup = new Date().toLocaleDateString('fa-IR')
    }
  } catch (error) {
    console.error('خطا در ایجاد پشتیبان:', error)
    alert('خطا در ایجاد پشتیبان')
  }
}

const restoreBackup = () => {
  if (confirm('آیا مطمئن هستید که می‌خواهید از پشتیبان بازیابی کنید؟ این عمل غیرقابل برگشت است.')) {
    // Navigate to restore page or open modal
    router.push('/admin/finance/installment-payments/backup/restore')
  }
}

onMounted(() => {
  // Load current settings
})
</script>
