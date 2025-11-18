<template>
  <div class="space-y-8">
    <!-- Header -->
    <div>
      <h3 class="text-lg font-semibold text-gray-800">تنظیمات و اعلان‌ها</h3>
      <p class="text-gray-600 text-sm">مدیریت اعلان‌های سیستم و تنظیمات مربوطه</p>
    </div>
    <!-- Notification Settings -->
    <div class="bg-white rounded-lg border border-gray-200 px-4 py-4 space-y-4">
      <h4 class="text-md font-semibold text-gray-800 mb-2">تنظیمات اعلان</h4>
      <div class="space-y-3">
        <div class="flex items-center justify-between">
          <span class="text-sm text-gray-700">اعلان خطاهای ارسال پیامک</span>
          <input type="checkbox" v-model="settings.smsError" class="form-checkbox h-5 w-5 text-blue-600">
        </div>
        <div class="flex items-center justify-between">
          <span class="text-sm text-gray-700">اعلان پاسخ‌های جدید نظرسنجی</span>
          <input type="checkbox" v-model="settings.newResponse" class="form-checkbox h-5 w-5 text-blue-600">
        </div>
        <div class="flex items-center justify-between">
          <span class="text-sm text-gray-700">اعلان آمار روزانه</span>
          <input type="checkbox" v-model="settings.dailyStats" class="form-checkbox h-5 w-5 text-blue-600">
        </div>
        <div class="flex items-center justify-between">
          <span class="text-sm text-gray-700">دریافت اعلان‌ها از طریق ایمیل</span>
          <input type="checkbox" v-model="settings.email" class="form-checkbox h-5 w-5 text-blue-600">
        </div>
      </div>
      <div class="flex items-center justify-end mt-4">
        <button @click="saveSettings" :disabled="saving" class="px-6 py-2 bg-green-600 hover:bg-green-700 text-white rounded-lg text-sm flex items-center space-x-2 space-x-reverse">
          <svg v-if="saving" class="animate-spin w-4 h-4" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          <span>{{ saving ? 'در حال ذخیره...' : 'ذخیره تنظیمات' }}</span>
        </button>
      </div>
    </div>
    <!-- Recent Notifications -->
    <div class="bg-white rounded-lg border border-gray-200 px-4 py-4">
      <h4 class="text-md font-semibold text-gray-800 mb-4">اعلان‌های اخیر</h4>
      <div v-if="notifications.length" class="space-y-3">
        <div v-for="notif in notifications" :key="notif.id" class="flex items-center gap-3 p-3 rounded-lg border border-gray-100 bg-gray-50">
          <div>
            <span v-if="notif.type==='error'" class="inline-block w-2 h-2 bg-red-500 rounded-full mr-1"></span>
            <span v-else-if="notif.type==='response'" class="inline-block w-2 h-2 bg-blue-500 rounded-full mr-1"></span>
            <span v-else-if="notif.type==='stats'" class="inline-block w-2 h-2 bg-green-500 rounded-full mr-1"></span>
            <span v-else class="inline-block w-2 h-2 bg-gray-400 rounded-full mr-1"></span>
          </div>
          <div class="flex-1">
            <div class="text-sm text-gray-900">{{ notif.message }}</div>
            <div class="text-xs text-gray-500 mt-1">{{ formatDate(notif.date) }}</div>
          </div>
        </div>
      </div>
      <div v-else class="text-sm text-gray-500">اعلانی وجود ندارد.</div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref } from 'vue'
const settings = ref({
  smsError: true,
  newResponse: true,
  dailyStats: false,
  email: false
})
const saving = ref(false)
const notifications = ref([
  { id: 1, type: 'error', message: 'خطا در ارسال پیامک به سفارش #1002', date: '2024-06-07T10:30:00Z' },
  { id: 2, type: 'response', message: 'پاسخ جدید از علی احمدی برای سفارش #1001', date: '2024-06-07T09:20:00Z' },
  { id: 3, type: 'stats', message: 'گزارش آمار روزانه آماده شد.', date: '2024-06-06T23:59:00Z' }
])
const saveSettings = async () => {
  saving.value = true
  try {
    await new Promise(resolve => setTimeout(resolve, 1200))
    alert('تنظیمات اعلان با موفقیت ذخیره شد!')
  } finally {
    saving.value = false
  }
}
const formatDate = (date: string) => new Date(date).toLocaleString('fa-IR')
</script> 
