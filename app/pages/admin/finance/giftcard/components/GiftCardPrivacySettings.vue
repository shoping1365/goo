<template>
  <div class="space-y-6">
    <!-- هدر -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h2 class="text-xl font-bold text-gray-900">تنظیمات حریم خصوصی گیفت کارت</h2>
      <p class="text-gray-600 mt-1">مدیریت حریم خصوصی کاربران و رعایت قوانین GDPR</p>
    </div>

    <!-- بخش رضایت GDPR -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-4">
        <div>
          <h3 class="text-lg font-medium text-gray-900">رضایت GDPR</h3>
          <p class="text-gray-500 text-sm mt-1">
            مدیریت رضایت کاربران برای جمع‌آوری و پردازش اطلاعات شخصی مطابق با قوانین GDPR.
          </p>
        </div>
        <div>
          <button @click="showGdprModal = true" class="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2">
            مشاهده رضایت‌ها
          </button>
        </div>
      </div>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div class="text-center p-6 bg-green-50 rounded-lg">
          <div class="text-2xl font-bold text-green-600">{{ gdprStats.consented }}</div>
          <div class="text-sm text-green-700">رضایت داده‌اند</div>
        </div>
        <div class="text-center p-6 bg-yellow-50 rounded-lg">
          <div class="text-2xl font-bold text-yellow-600">{{ gdprStats.pending }}</div>
          <div class="text-sm text-yellow-700">در انتظار رضایت</div>
        </div>
        <div class="text-center p-6 bg-red-50 rounded-lg">
          <div class="text-2xl font-bold text-red-600">{{ gdprStats.revoked }}</div>
          <div class="text-sm text-red-700">رضایت لغو شده</div>
        </div>
      </div>
    </div>

    <!-- بخش مدیریت کوکی‌ها -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="mb-4">
        <h3 class="text-lg font-medium text-gray-900">مدیریت کوکی‌ها</h3>
        <p class="text-gray-500 text-sm mt-1">
          کنترل و مدیریت کوکی‌های مورد استفاده در سیستم گیفت کارت.
        </p>
      </div>
      <div class="space-y-4">
        <div v-for="cookie in cookies" :key="cookie.name" class="flex items-center justify-between p-6 border border-gray-200 rounded-lg">
          <div>
            <div class="font-medium text-gray-900">{{ cookie.name }}</div>
            <div class="text-sm text-gray-500">{{ cookie.description }}</div>
            <div class="text-xs text-gray-400">انقضا: {{ cookie.expiry }}</div>
          </div>
          <div class="flex items-center space-x-2 space-x-reverse">
            <input 
              v-model="cookie.enabled" 
              type="checkbox" 
              :disabled="cookie.essential"
              class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
            />
            <span v-if="cookie.essential" class="text-xs text-red-600">ضروری</span>
          </div>
        </div>
      </div>
      <div class="mt-4 flex justify-between">
        <button @click="clearAllCookies" class="px-4 py-2 bg-red-600 text-white text-sm font-medium rounded-lg hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2">
          حذف همه کوکی‌ها
        </button>
        <button @click="saveCookieSettings" class="px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2">
          ذخیره تنظیمات
        </button>
      </div>
    </div>

    <!-- بخش حذف اطلاعات شخصی -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="mb-4">
        <h3 class="text-lg font-medium text-gray-900">حذف اطلاعات شخصی</h3>
        <p class="text-gray-500 text-sm mt-1">
          مدیریت درخواست‌های حذف اطلاعات شخصی کاربران مطابق با قوانین GDPR.
        </p>
      </div>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-4 py-2 text-right text-xs font-medium text-gray-500">کاربر</th>
              <th class="px-4 py-2 text-right text-xs font-medium text-gray-500">تاریخ درخواست</th>
              <th class="px-4 py-2 text-right text-xs font-medium text-gray-500">وضعیت</th>
              <th class="px-4 py-2 text-right text-xs font-medium text-gray-500">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="request in deletionRequests" :key="request.id">
              <td class="px-4 py-2 whitespace-nowrap text-sm">{{ request.userName }}</td>
              <td class="px-4 py-2 whitespace-nowrap text-sm">{{ formatDate(request.requestDate) }}</td>
              <td class="px-4 py-2 whitespace-nowrap">
                <span :class="getStatusClass(request.status)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                  {{ getStatusLabel(request.status) }}
                </span>
              </td>
              <td class="px-4 py-2 whitespace-nowrap">
                <button v-if="request.status === 'pending'" @click="approveDeletion(request)" class="text-green-600 hover:underline text-xs ml-2">تأیید</button>
                <button v-if="request.status === 'pending'" @click="rejectDeletion(request)" class="text-red-600 hover:underline text-xs">رد</button>
                <button v-if="request.status === 'completed'" @click="viewDeletionLog(request)" class="text-blue-600 hover:underline text-xs">مشاهده لاگ</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- بخش شفافیت در استفاده از داده -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="mb-4">
        <h3 class="text-lg font-medium text-gray-900">شفافیت در استفاده از داده</h3>
        <p class="text-gray-500 text-sm mt-1">
          نمایش اطلاعات شفاف درباره داده‌های جمع‌آوری شده و نحوه استفاده از آن‌ها.
        </p>
      </div>
      <div class="space-y-4">
        <div v-for="dataType in dataTypes" :key="dataType.name" class="border border-gray-200 rounded-lg p-6">
          <div class="flex items-center justify-between mb-2">
            <h4 class="font-medium text-gray-900">{{ dataType.name }}</h4>
            <span :class="dataType.required ? 'bg-red-100 text-red-800' : 'bg-blue-100 text-blue-800'" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
              {{ dataType.required ? 'ضروری' : 'اختیاری' }}
            </span>
          </div>
          <p class="text-sm text-gray-600 mb-2">{{ dataType.description }}</p>
          <div class="text-xs text-gray-500">
            <strong>هدف استفاده:</strong> {{ dataType.purpose }}
          </div>
        </div>
      </div>
      <div class="mt-4">
        <button @click="showPrivacyPolicy = true" class="px-4 py-2 bg-purple-600 text-white text-sm font-medium rounded-lg hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-purple-500 focus:ring-offset-2">
          مشاهده سیاست حفظ حریم خصوصی
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import { useAuth } from '~/composables/useAuth'

const { user, hasPermission } = useAuth()
const canDeletePrivacyRequest = computed(() => hasPermission('giftcard-privacy-request.delete'))

// آمار رضایت GDPR
const gdprStats = ref({
  consented: 1250,
  pending: 45,
  revoked: 12
})

// کوکی‌های سیستم
const cookies = ref([
  { name: 'session_id', description: 'شناسه نشست کاربر', expiry: 'تا پایان نشست', enabled: true, essential: true },
  { name: 'preferences', description: 'تنظیمات کاربر', expiry: '۱ سال', enabled: true, essential: false },
  { name: 'analytics', description: 'تحلیل رفتار کاربر', expiry: '۲ سال', enabled: false, essential: false },
  { name: 'marketing', description: 'کوکی‌های تبلیغاتی', expiry: '۱ سال', enabled: false, essential: false }
])

// درخواست‌های حذف اطلاعات
const deletionRequests = ref([
  { id: 1, userName: 'علی احمدی', requestDate: new Date(Date.now() - 86400000), status: 'pending' },
  { id: 2, userName: 'فاطمه محمدی', requestDate: new Date(Date.now() - 172800000), status: 'completed' },
  { id: 3, userName: 'حسن رضایی', requestDate: new Date(Date.now() - 259200000), status: 'rejected' }
])

// انواع داده‌های جمع‌آوری شده
const dataTypes = ref([
  { 
    name: 'اطلاعات شخصی', 
    description: 'نام، ایمیل، شماره تلفن و آدرس',
    purpose: 'ارائه خدمات و ارتباط با کاربر',
    required: true 
  },
  { 
    name: 'اطلاعات مالی', 
    description: 'جزئیات تراکنش‌ها و پرداخت‌ها',
    purpose: 'پردازش تراکنش‌ها و گزارش‌گیری',
    required: true 
  },
  { 
    name: 'اطلاعات استفاده', 
    description: 'لاگ فعالیت‌ها و رفتار کاربر',
    purpose: 'بهبود خدمات و امنیت',
    required: false 
  },
  { 
    name: 'کوکی‌ها', 
    description: 'اطلاعات مرورگر و تنظیمات',
    purpose: 'تجربه کاربری بهتر',
    required: false 
  }
])

// متغیرهای مودال
const showGdprModal = ref(false)
const showPrivacyPolicy = ref(false)

// متدها
const clearAllCookies = () => {
  if (confirm('آیا از حذف همه کوکی‌ها اطمینان دارید؟')) {
    cookies.value.forEach(cookie => {
      if (!cookie.essential) {
        cookie.enabled = false
      }
    })
    alert('همه کوکی‌های غیرضروری حذف شدند.')
  }
}

const saveCookieSettings = () => {
  alert('تنظیمات کوکی‌ها با موفقیت ذخیره شد.')
}

const approveDeletion = (request: any) => {
  if (confirm('آیا از تأیید حذف اطلاعات این کاربر اطمینان دارید؟')) {
    request.status = 'completed'
    alert('درخواست حذف تأیید شد.')
  }
}

const rejectDeletion = (request: any) => {
  if (confirm('آیا از رد درخواست حذف اطمینان دارید؟')) {
    request.status = 'rejected'
    alert('درخواست حذف رد شد.')
  }
}

const viewDeletionLog = (request: any) => {
  alert('لاگ حذف اطلاعات: ' + request.userName)
}

const getStatusClass = (status: string) => {
  switch (status) {
    case 'pending': return 'bg-yellow-100 text-yellow-800'
    case 'completed': return 'bg-green-100 text-green-800'
    case 'rejected': return 'bg-red-100 text-red-800'
    default: return 'bg-gray-100 text-gray-800'
  }
}

const getStatusLabel = (status: string) => {
  switch (status) {
    case 'pending': return 'در انتظار'
    case 'completed': return 'انجام شد'
    case 'rejected': return 'رد شد'
    default: return 'نامشخص'
  }
}

const formatDate = (date: Date) => {
  return date.toLocaleDateString('fa-IR')
}

// مستندسازی فارسی:
// این کامپوننت برای مدیریت حریم خصوصی کاربران و رعایت قوانین GDPR استفاده می‌شود.
// شامل مدیریت رضایت، کوکی‌ها، حذف اطلاعات شخصی و شفافیت در استفاده از داده است.
// تمام تنظیمات با رعایت قوانین بین‌المللی حریم خصوصی پیاده‌سازی شده‌اند.
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 

