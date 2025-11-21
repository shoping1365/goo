<template>
  <div dir="rtl">
    <div class="bg-white rounded-lg shadow p-6 mb-8 flex items-center justify-between">
      <div class="flex items-center gap-6 w-full justify-between">
        <h1 class="text-2xl font-bold">ویرایش درگاه پیامک</h1>
        <button class="inline-flex items-center justify-center gap-2 w-40 h-12 text-base px-0 py-0 rounded-lg bg-gradient-to-r from-purple-400 to-purple-600 text-white font-bold shadow hover:from-purple-500 hover:to-purple-700 transition" @click="$router.back()">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" /></svg>
          بازگشت
        </button>
      </div>
    </div>
    
    <!-- نمایش loading state -->
    <div v-if="isLoading || pending" class="text-center py-8">
      <div class="inline-flex items-center px-4 py-2 font-semibold leading-6 text-sm shadow rounded-md text-white bg-purple-500 hover:bg-purple-400 transition ease-in-out duration-150 cursor-not-allowed">
        <svg class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
        </svg>
        در حال بارگذاری اطلاعات درگاه...
      </div>
    </div>
    
    <!-- نمایش خطا -->
    <div v-else-if="error" class="text-center py-8">
      <div class="bg-red-50 border border-red-200 rounded-lg p-6">
        <div class="flex items-center justify-center mb-4">
          <svg class="w-8 h-8 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
        </div>
        <h3 class="text-lg font-medium text-red-800 mb-2">خطا در بارگذاری اطلاعات درگاه</h3>
        <p class="text-red-600 mb-4">متأسفانه در بارگذاری اطلاعات درگاه مشکلی پیش آمده است.</p>
        <button class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-red-600 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500" @click="() => refresh()">
          تلاش مجدد
        </button>
      </div>
    </div>
    
    <!-- محتوای اصلی فرم -->
    <div v-else>
    <div class="bg-white rounded-lg shadow p-6 py-2 mt-8 max-w-4xl mx-auto">
      <!-- انتخاب نوع درگاه پیامک -->
      <div class="mb-6">
        <label class="block mb-6 font-semibold">نوع درگاه پیامک</label>
        <select v-model="selectedGateway" class="form-select w-full rounded-lg border border-purple-400 focus:border-purple-500 focus:ring focus:ring-purple-200">
          <option disabled value="">یک درگاه را انتخاب کنید...</option>
          <option v-for="gateway in gateways" :key="gateway.value" :value="gateway.value">
            {{ gateway.label }}
          </option>
        </select>
      </div>
    </div>
    <!-- تنظیمات اختصاصی هر درگاه در کانتینر جدا -->
    <div v-if="selectedGateway" class="bg-white rounded-2xl shadow-lg p-8 mt-8 max-w-4xl mx-auto">
      <div class="flex items-center justify-between mb-8">
        <h2 class="text-xl font-bold flex items-center gap-2">
          <svg class="w-6 h-6 text-purple-500" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 01.88 7.903A4.001 4.001 0 0112 21a4.001 4.001 0 01-4.88-6.097A4 4 0 018 7V5a4 4 0 118 0v2z" /></svg>
          تنظیمات {{ getGatewayLabel(selectedGateway) }}
        </h2>
        
        <!-- دکمه فعال/غیرفعال کردن درگاه -->
        <div class="flex items-center gap-3">
          <span class="text-sm font-medium text-gray-700">وضعیت درگاه:</span>
          <button 
            :disabled="isUpdatingStatus" 
            class="relative inline-flex h-6 w-11 items-center rounded-full transition-colors focus:outline-none focus:ring-2 focus:ring-purple-500 focus:ring-offset-2"
            :class="gatewayStatus ? 'bg-purple-600' : 'bg-gray-200'"
            @click="toggleGatewayStatus"
          >
            <span 
              class="inline-block h-4 w-4 transform rounded-full bg-white transition-transform"
              :class="gatewayStatus ? 'translate-x-6' : 'translate-x-1'"
            ></span>
          </button>
          <span class="text-sm font-medium" :class="gatewayStatus ? 'text-green-600' : 'text-gray-500'">
            {{ gatewayStatus ? 'فعال' : 'غیرفعال' }}
          </span>
          <div v-if="isUpdatingStatus" class="ml-2">
            <svg class="animate-spin h-4 w-4 text-purple-500" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
          </div>
        </div>
      </div>
      <!-- آدرس API -->
      <div class="mb-8">
        <div class="flex flex-row-reverse items-center gap-3 bg-purple-50 border border-purple-200 rounded-lg p-6">
          <label class="font-bold whitespace-nowrap">:آدرس API</label>
          <input v-model="apiUrl" type="text" class="form-input w-full rounded-lg border-gray-300 focus:border-purple-500 focus:ring focus:ring-purple-200 font-mono text-xs select-all text-left" dir="ltr" placeholder="مثال: https://api.example.com/send" />
        </div>
      </div>

      <!-- کلید API -->
      <div class="mb-8">
        <div class="flex flex-row-reverse items-center gap-3 bg-purple-50 border border-purple-200 rounded-lg p-6">
          <label class="font-bold whitespace-nowrap">:کلید API</label>
          <input v-model="apiKey" type="password" class="form-input w-full rounded-lg border-gray-300 focus:border-purple-500 focus:ring focus:ring-purple-200 font-mono text-xs select-all text-left" dir="ltr" placeholder="کلید API درگاه" />
        </div>
      </div>
      <!-- فیلدهای اضافی هر درگاه -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
        <div v-for="field in getGatewayFields(selectedGateway)" :key="field.name" class="bg-white border-2 border-purple-300 rounded-xl p-6 flex flex-col gap-2">
          <label :for="field.name" class="block mb-2 font-semibold">{{ field.label }}</label>
          <input
              :id="field.name"
              v-model="form[field.name]"
              :type="field.name === 'password' ? 'password' : 'text'"
              class="form-input w-full rounded-lg border-gray-300 focus:border-purple-500 focus:ring focus:ring-purple-200"
              :placeholder="field.placeholder || ''"
          />
        </div>
        <!-- فیلد ارسال بر اساس پترن -->
        <div class="bg-white border-2 border-purple-300 rounded-xl p-6 flex flex-col gap-2">
          <label class="block mb-2 font-semibold">ارسال بر اساس پترن</label>
          <div class="flex gap-6">
            <label class="inline-flex items-center gap-2 cursor-pointer">
              <input v-model="form.pattern_based" type="radio" value="yes" class="form-radio text-purple-500" />
              <span>بله</span>
            </label>
            <label class="inline-flex items-center gap-2 cursor-pointer">
              <input v-model="form.pattern_based" type="radio" value="no" class="form-radio text-purple-500" />
              <span>خیر</span>
            </label>
          </div>
        </div>
      </div>
      <!-- باکس تست ارسال پیامک -->
      <div class="bg-purple-50 border-2 border-purple-300 rounded-2xl p-8 mt-8 mb-8 shadow flex flex-col gap-6">
        <div class="font-bold text-purple-700 text-lg mb-2 flex items-center gap-2">
          <svg class="w-5 h-5 text-purple-500" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" /></svg>
          تست ارسال پیامک
        </div>
        <!-- باکس سفید ورودی‌ها -->
        <div class="bg-white border border-gray-200 rounded-xl shadow-sm p-6 flex flex-col gap-6">
          <div>
            <label class="block mb-2 font-semibold text-gray-700">شماره موبایل گیرنده</label>
            <input v-model="testMobile" type="text" placeholder="مثال: 0912xxxxxxx" class="form-input w-full rounded-lg border-gray-300 focus:border-purple-500 focus:ring focus:ring-purple-200" />
          </div>
          <hr class="my-2 border-t border-purple-300" />
          <div>
            <label class="block mb-2 font-semibold text-gray-700">متن پیامک</label>
            <textarea v-model="testMessage" rows="3" placeholder="متن پیام تستی" class="form-input w-full rounded-lg border-gray-300 focus:border-purple-500 focus:ring focus:ring-purple-200 resize-none"></textarea>
          </div>
        </div>
        <button :disabled="!testMobile || !testMessage || !isFormValid" class="inline-flex items-center justify-center gap-2 w-40 h-12 text-base px-0 py-0 rounded-lg bg-gradient-to-r from-purple-400 to-purple-600 text-white font-bold shadow hover:from-purple-500 hover:to-purple-700 transition disabled:opacity-50 disabled:cursor-not-allowed mt-4 self-start md:self-auto" @click="testSendSms">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" /></svg>
          تست ارسال
        </button>
        <div class="mt-2">
          <div v-if="testSendStatus === 'success'" class="flex items-center gap-2 bg-green-50 border border-green-200 text-green-700 rounded-lg px-4 py-2">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" /></svg>
            پیامک با موفقیت ارسال شد
          </div>
          <div v-if="testSendStatus === 'error'" class="flex items-center gap-2 bg-red-50 border border-red-200 text-red-700 rounded-lg px-4 py-2">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" /></svg>
            ارسال پیامک ناموفق بود
          </div>
        </div>
      </div>
      <!-- دکمه تست اتصال و ثبت درگاه (در انتهای فرم) -->
      <div class="flex flex-col md:flex-row md:items-center md:justify-between gap-6 mt-8">
        <div class="flex items-center gap-6 order-2 md:order-1">
          <button :disabled="!isFormValid" class="inline-flex items-center justify-center gap-2 w-40 h-12 text-base px-0 py-0 rounded-lg bg-gradient-to-r from-blue-400 to-blue-600 text-white font-bold shadow hover:from-blue-500 hover:to-blue-700 transition disabled:opacity-50 disabled:cursor-not-allowed" @click="testConnection">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 11V7m0 8h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
            تست اتصال
          </button>
        </div>
        <div class="flex justify-end order-1 md:order-2 w-full md:w-auto">
          <button :disabled="!selectedGateway || !isFormValid" class="inline-flex items-center justify-center gap-2 w-40 h-12 text-base px-0 py-0 rounded-lg bg-gradient-to-r from-purple-400 to-purple-600 text-white font-bold shadow hover:from-purple-500 hover:to-purple-700 transition disabled:opacity-50 disabled:cursor-not-allowed" @click="submitForm">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" /></svg>
            ثبت درگاه
          </button>
        </div>
      </div>
    </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
declare const useFetch: <T>(url: string, options?: unknown) => Promise<{ data: { value: T }; pending: { value: boolean }; error: { value: unknown }; refresh: () => Promise<void> }>
</script>

<script setup lang="ts">
import { computed, ref, watchEffect } from 'vue';
import { useRoute, useRouter } from 'vue-router';

definePageMeta({ layout: 'admin-main' })

// لیست انواع درگاه و فیلدهای هرکدام
const gateways = [
  {
    value: 'meli_payamak',
    label: 'ملی پیامک',
    fields: [
      { name: 'username', label: 'نام کاربری', placeholder: 'نام کاربری ملی پیامک' },
      { name: 'password', label: 'رمز عبور', placeholder: 'رمز عبور ملی پیامک' },
      { name: 'from', label: 'شماره ارسال‌کننده', placeholder: 'مثال: 1000xxxxxx' }
    ]
  },
  {
    value: 'farazsms',
    label: 'فراز اس‌ام‌اس',
    fields: [
      { name: 'from', label: 'شماره ارسال‌کننده', placeholder: 'مثال: 1000xxxxxx' },
      { name: 'username', label: 'نام کاربری', placeholder: 'نام کاربری فراز اس‌ام‌اس' },
      { name: 'password', label: 'رمز عبور', placeholder: 'رمز عبور فراز اس‌ام‌اس' }
    ]
  },
  {
    value: 'ippanel',
    label: 'آی‌پی‌پنل',
    fields: [
      { name: 'from', label: 'شماره ارسال‌کننده', placeholder: 'مثال: 1000xxxxxx' },
      { name: 'username', label: 'نام کاربری', placeholder: 'نام کاربری آی‌پی‌پنل' },
      { name: 'password', label: 'رمز عبور', placeholder: 'رمز عبور آی‌پی‌پنل' }
    ]
  },
  {
    value: 'kavenegar',
    label: 'کاوه‌نگار',
    fields: [
      { name: 'username', label: 'نام کاربری', placeholder: 'نام کاربری کاوه‌نگار' },
      { name: 'password', label: 'رمز عبور', placeholder: 'رمز عبور کاوه‌نگار' }
    ]
  }
]

const selectedGateway = ref('')
const form = ref<Record<string, string>>({})
const apiUrl = ref('')
const apiKey = ref('')

// متغیرهای مربوط به وضعیت درگاه
const gatewayStatus = ref(false)
const isUpdatingStatus = ref(false)

const route = useRoute()
const id = route.params.id

// تابع‌های کمکی (تعریف شده قبل از استفاده)
function getGatewayLabel(val: string) {
  return gateways.find(g => g.value === val)?.label || ''
}

function getGatewayFields(val: string) {
  return gateways.find(g => g.value === val)?.fields || []
}

// تابع به‌روزرسانی نام درگاه (تعریف شده قبل از استفاده)
const _updateGatewayName = async () => {
  if (!selectedGateway.value) return
  
  try {
    await $fetch(`/api/sms-gateways/${id}`, {
      method: 'PUT',
      body: {
        name: getGatewayLabel(selectedGateway.value)
      }
    })
  } catch (error) {
    console.error('خطا در به‌روزرسانی نام درگاه:', error)
  }
}

interface GatewayData {
  id: number
  name: string
  type: string
  [key: string]: unknown
}

// دریافت اطلاعات درگاه
const { data: gatewayData, pending, error, refresh } = await useFetch<{ data: GatewayData }>(`/api/sms-gateways/${id}`)

// اضافه کردن loading state برای نمایش وضعیت بارگذاری
const isLoading = ref(true)

watchEffect(() => {
  if (gatewayData.value?.data) {
    const gw = gatewayData.value.data
    selectedGateway.value = gw.type || gw.gateway_type
    
    // تنظیم apiUrl و apiKey از دیتابیس
    apiUrl.value = gw.api_url || ''
    apiKey.value = gw.api_key || ''
    
    // اگر نام درگاه در دیتابیس موجود نیست، از نوع درگاه بگیر
    if (!gw.name && selectedGateway.value) {
      // اینجا می‌توانیم نام را به‌روزرسانی کنیم
      // updateGatewayName() // حذف شده تا از خطای lexical declaration جلوگیری شود
    }
    
    // مقداردهی فیلدهای اضافی بر اساس ساختار backend
    const fields = getGatewayFields(gw.type || gw.gateway_type)
    fields.forEach(f => {
      if (f.name === 'from') {
        form.value[f.name] = gw.sender_number || gw.from || gw.sender || ''
      } else {
        form.value[f.name] = gw[f.name] || gw[f.name.toLowerCase()] || ''
      }
    })
    
    // تنظیم pattern_based
    form.value.pattern_based = gw.pattern_based ? 'yes' : 'no'
    
    // تنظیم وضعیت درگاه
    gatewayStatus.value = gw.is_active || false
    
    isLoading.value = false
  } else if (error.value) {
    isLoading.value = false
  }
})

// تابع تغییر وضعیت درگاه
const toggleGatewayStatus = async () => {
  if (isUpdatingStatus.value) return
  
  isUpdatingStatus.value = true
  try {
    const newStatus = !gatewayStatus.value
    
    await $fetch(`/api/sms-gateways/${id}`, {
      method: 'PUT',
      body: {
        is_active: newStatus,
        name: getGatewayLabel(selectedGateway.value) // حفظ نام درگاه
      }
    })
    
    gatewayStatus.value = newStatus
    alert(`درگاه ${newStatus ? 'فعال' : 'غیرفعال'} شد!`)
  } catch (error) {
    console.error('خطا در تغییر وضعیت درگاه:', error)
    alert('خطا در تغییر وضعیت درگاه!')
  } finally {
    isUpdatingStatus.value = false
  }
}

// حذف watch که apiValue را با URL پیش‌فرض جایگزین می‌کند
// watch(selectedGateway, (val) => {
//   apiValue.value = getGatewayApiInfo(val)
// })

// اعتبارسنجی ساده: فیلدهای اضافی باید پر باشند
const isFormValid = computed(() => {
  if (!selectedGateway.value) return false
  const fields = getGatewayFields(selectedGateway.value)
  return fields.every(f => form.value[f.name] && form.value[f.name].trim() !== '')
})

// ثبت فرم (نمونه)
const router = useRouter()

async function submitForm() {
  if (!isFormValid.value) return

  // ساخت payload بر اساس ساختار backend
  interface GatewayPayload {
    type: string
    name: string
    api_url?: string
    api_key?: string
    username?: string
    password?: string
    sender_number?: string
    [key: string]: unknown
  }
  const payload: GatewayPayload = {
    type: selectedGateway.value,
    name: getGatewayLabel(selectedGateway.value), // استفاده از نام نوع درگاه
    api_url: apiUrl.value, // آدرس API
    api_key: apiKey.value, // کلید API
    pattern_based: form.value.pattern_based === 'yes',
    is_active: gatewayStatus.value, // استفاده از وضعیت فعلی
    priority: 1 // پیش‌فرض اولویت 1
  }
  
  // نگاشت فیلدهای اضافی به ساختار backend
  const fields = getGatewayFields(selectedGateway.value)
  fields.forEach(f => {
    if (f.name === 'from') {
      payload.sender_number = form.value[f.name]
    } else {
      payload[f.name] = form.value[f.name]
    }
  })

  try {
    // استفاده از PUT برای به‌روزرسانی درگاه موجود
    await $fetch(`/api/sms-gateways/${id}`, {
      method: 'PUT',
      body: payload
    })
    // موفقیت
    alert('درگاه با موفقیت به‌روزرسانی شد!')
    router.push('/admin/notifications/system-settings')
  } catch (e) {
    console.error('خطا در به‌روزرسانی درگاه:', e)
    alert('خطا در به‌روزرسانی درگاه!')
  }
}

// اطلاعات API برای هر درگاه (در صورت نیاز)
const gatewayApiInfo = {
  meli_payamak: 'https://rest.payamak-panel.com/api/SendSMS/',
  farazsms: 'https://edge.ippanel.com/v1/',
        ippanel: 'https://rest.ippanel.com/v1/messages/patterns/send',
  kavenegar: 'https://api.kavenegar.com/v1/{API-KEY}/sms/send.json'
}
const _getGatewayApiInfo = (val: string) => {
  return gatewayApiInfo[val] || ''
}

// تست اتصال واقعی
const testConnectionStatus = ref('') // '', 'success', 'error'
const testConnectionMessage = ref('') // پیام جزئیات تست اتصال

async function testConnection() {
  if (!isFormValid.value) return

  testConnectionStatus.value = ''
  testConnectionMessage.value = ''

  try {
    const response = await $fetch<{status: string, message?: string, error_message?: string}>(`/api/sms-gateways/${id}/test-connection`, {
      method: 'POST'
    })
    if (response.status === 'success') {
      testConnectionStatus.value = 'success'
      testConnectionMessage.value = response.message || 'اتصال موفق بود!'
      alert('✅ اتصال موفق بود!\n\n' + (response.message || 'درگاه به درستی پیکربندی شده و قابل استفاده است.'))
    } else {
      testConnectionStatus.value = 'error'
      testConnectionMessage.value = response.error_message || 'خطا در اتصال'
      alert('❌ خطا در اتصال!\n\n' + (response.error_message || 'لطفاً تنظیمات درگاه را بررسی کنید.'))
    }
  } catch (e: unknown) {
    interface ErrorResponse {
      response?: {
        _data?: {
          error_message?: string
        }
      }
      message?: string
    }
    const error = e as ErrorResponse
    testConnectionStatus.value = 'error'
    let errorMessage = 'خطا در اتصال!'
    if (error.response && error.response._data) {
      const errorData = error.response._data
      if (errorData.error_message) {
        errorMessage = errorData.error_message
        testConnectionMessage.value = errorData.error_message
      } else if (errorData.message) {
        errorMessage = errorData.message
        testConnectionMessage.value = errorData.message
      }
    }
    alert('❌ خطا در اتصال!\n\n' + errorMessage)
  }
}

// تست ارسال پیامک واقعی
const testMobile = ref('')
const testMessage = ref('')
const testSendStatus = ref('') // '', 'success', 'error'
async function testSendSms() {
  if (!testMobile.value || !testMessage.value || !isFormValid.value) return
  
      testSendStatus.value = ''
    try {
      const response = await $fetch(`/api/sms-gateways/${id}/send-test`, {
        method: 'POST',
        body: {
          mobile: testMobile.value,
          message: testMessage.value
        }
      })
      testSendStatus.value = 'success'
      alert('پیامک با موفقیت ارسال شد!')
    } catch (error) {
      testSendStatus.value = 'error'
      alert('خطا در ارسال پیامک!')
    }
  }
</script>