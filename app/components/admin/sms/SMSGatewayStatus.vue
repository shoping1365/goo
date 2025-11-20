<template>
  <div class="rounded-xl shadow-lg p-6 bg-white" dir="rtl">
    <!-- فقط عنوان آمار بالای کارت -->
    <div class="mb-4">
      <p class="text-gray-600 text-sm">مانیتورینگ لحظه‌ای و آمار عملکرد درگاه‌ها</p>
    </div>
    <!-- کارت وضعیت درگاه -->
    <div class="grid grid-cols-1 gap-6 mb-6">
      <div class="rounded-lg p-6 border border-gray-200 min-h-[300px] flex flex-col justify-between bg-white">
        <div>
          <div class="flex items-center justify-between mb-4">
            <div class="flex items-center space-x-3 space-x-reverse">
              <div class="w-4 h-4 rounded-full" :class="getStatusColor(gateway.status)"></div>
              <div>
                <h4 class="font-semibold text-gray-800">{{ gateway.name || 'درگاه بدون نام' }}</h4>
                <p class="text-sm text-gray-600">اولویت {{ gateway.priority || 1 }}</p>
              </div>
            </div>
            <div class="text-right">
              <p class="text-sm text-gray-600">آخرین بررسی</p>
              <p class="text-xs text-gray-500">{{ formatTime(gateway.lastCheck) }}</p>
            </div>
          </div>
          <!-- برای ملی پیامک: نمایش 4 ستون -->
          <div v-if="gateway.type === 'meli_payamak'" class="grid grid-cols-4 gap-6 mb-4">
            <div class="text-center">
              <p class="text-2xl font-bold" :class="getStatusTextColor(gateway.status)">{{ gateway.successCount || 0 }}</p>
              <p class="text-xs text-gray-600">موفق</p>
            </div>
            <div class="text-center">
              <p class="text-2xl font-bold text-red-600">{{ gateway.failureCount || 0 }}</p>
              <p class="text-xs text-gray-600">ناموفق</p>
            </div>
            <div class="text-center cursor-pointer" title="کلیک برای به‌روزرسانی اطلاعات" @click="fetchMeliPayamakInfo(gateway.id)">
              <div v-if="isLoadingBalance" class="inline-block animate-spin rounded-full h-6 w-6 border-b-2 border-blue-600"></div>
              <div v-else>
                <p class="text-2xl font-bold text-blue-600">{{ meliPayamakInfo.remaining_sms || 0 }}</p>
                <p class="text-xs text-gray-600">SMS باقی‌مانده</p>
              </div>
            </div>
            <div class="text-center cursor-pointer" title="کلیک برای به‌روزرسانی اطلاعات" @click="fetchMeliPayamakInfo(gateway.id)">
              <div v-if="isLoadingBalance" class="inline-block animate-spin rounded-full h-6 w-6 border-b-2 border-green-600"></div>
              <div v-else>
                <p class="text-2xl font-bold text-green-600">{{ formatBalance(meliPayamakInfo.credit || 0) }}</p>
                <p class="text-xs text-gray-600">موجودی ریالی</p>
              </div>
            </div>
          </div>
          
          <!-- برای سایر درگاه‌ها: نمایش 3 ستون -->
          <div v-else class="grid grid-cols-3 gap-6 mb-4">
            <div class="text-center">
              <p class="text-2xl font-bold" :class="getStatusTextColor(gateway.status)">{{ gateway.successCount || 0 }}</p>
              <p class="text-xs text-gray-600">موفق</p>
            </div>
            <div class="text-center">
              <p class="text-2xl font-bold text-red-600">{{ gateway.failureCount || 0 }}</p>
              <p class="text-xs text-gray-600">ناموفق</p>
            </div>
            <div class="text-center cursor-pointer" title="کلیک برای به‌روزرسانی موجودی" @click="fetchGatewayBalance(gateway.id)">
              <p class="text-2xl font-bold text-blue-600 hover:text-blue-800 transition-colors">
                <span v-if="isLoadingBalance" class="inline-block animate-spin rounded-full h-6 w-6 border-b-2 border-blue-600"></span>
                <span v-else>{{ formatBalance(gatewayBalance) }}</span>
              </p>
              <p class="text-xs text-gray-600">موجودی ریالی</p>
            </div>
          </div>
          <div class="mb-3">
            <div class="flex justify-between text-xs text-gray-600 mb-1">
              <span>نرخ موفقیت</span>
              <span>{{ calculateSuccessRate() }}%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div
class="h-2 rounded-full transition-all duration-300" 
                   :class="getSuccessRateColor(calculateSuccessRate())"
                   :style="{ width: calculateSuccessRate() + '%' }"></div>
            </div>
          </div>
        </div>
        <div class="flex flex-col md:flex-row-reverse gap-6 mt-10 items-stretch">
          <button class="w-full md:w-48 px-10 py-3 rounded-lg text-base font-bold bg-gradient-to-r from-blue-500 to-blue-700 text-white shadow-md hover:from-blue-600 hover:to-blue-800 transition flex items-center justify-center gap-2" @click="testGateway(gateway.id)">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8h2a2 2 0 012 2v8a2 2 0 01-2 2H5a2 2 0 01-2-2v-8a2 2 0 012-2h2" /><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 3h-6a2 2 0 00-2 2v3h10V5a2 2 0 00-2-2z" /></svg>
            تست اتصال
          </button>
          <button class="w-full md:w-48 px-10 py-3 rounded-lg text-base font-bold bg-gradient-to-r from-green-500 to-green-700 text-white shadow-md hover:from-green-600 hover:to-green-800 transition flex items-center justify-center gap-2" @click="testSendSMS(gateway.id)">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" /></svg>
            تست ارسال
          </button>
          <button
class="w-full md:w-48 px-10 py-3 rounded-lg text-base font-bold shadow-md transition flex items-center justify-center gap-2" :class="(gateway.status === 'active' || gateway.is_active) ? 'bg-gradient-to-r from-red-500 to-red-700 text-white hover:from-red-600 hover:to-red-800' : 'bg-gradient-to-r from-green-400 to-green-600 text-white hover:from-green-500 hover:to-green-700'"
                  @click="toggleGateway(gateway.id)">
            <svg v-if="gateway.status === 'active' || gateway.is_active" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path></svg>
            <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg>
            {{ (gateway.status === 'active' || gateway.is_active) ? 'غیرفعال' : 'فعال' }}
          </button>
          <NuxtLink
            :to="`/admin/notifications/system-settings-create/${gateway.id}`"
            class="w-full md:w-48 px-10 py-3 rounded-lg text-base font-bold bg-gradient-to-r from-purple-500 to-purple-700 text-white hover:from-purple-600 hover:to-purple-800 border border-purple-300 shadow-md transition flex items-center justify-center gap-2"
          >
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24"><circle cx="12" cy="12" r="4" stroke="currentColor" stroke-width="2" fill="none" /><path stroke="currentColor" stroke-width="2" fill="none" d="M4.93 4.93l2.12 2.12M2 12h3M4.93 19.07l2.12-2.12M12 22v-3M19.07 19.07l-2.12-2.12M22 12h-3M19.07 4.93l-2.12 2.12M12 2v3" /></svg>
            تنظیمات
          </NuxtLink>
          <button class="w-full md:w-48 px-10 py-3 rounded-lg text-base font-bold bg-gradient-to-r from-red-600 to-red-800 text-white shadow-md hover:from-red-700 hover:to-red-900 transition flex items-center justify-center gap-2" @click="deleteGateway(gateway.id)">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" /></svg>
            حذف سخت
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'

// تعریف interface برای gateway
interface Gateway {
  id: number
  name?: string
  priority?: number
  status?: 'active' | 'inactive' | 'maintenance'
  successCount?: number
  failureCount?: number
  successRate?: number
  lastCheck?: Date
  is_active?: boolean
  type?: string
  api_url?: string
  sender_number?: string
  pattern_based?: boolean
  created_at?: string
  updated_at?: string
  [key: string]: any
}

// تعریف props برای دریافت gateway از parent component
const props = defineProps<{
  gateway: Gateway
}>()

// تعریف emit برای اطلاع‌رسانی به parent component
const emit = defineEmits<{
  gatewayDeleted: [id: number]
}>()

// متغیر reactive برای موجودی درگاه
const gatewayBalance = ref<number>(0)
const isLoadingBalance = ref<boolean>(false)

// متغیر reactive برای اطلاعات ملی پیامک
const meliPayamakInfo = ref<{remaining_sms: number, credit: number}>({
  remaining_sms: 0,
  credit: 0
})

// محاسبه نرخ موفقیت
const calculateSuccessRate = () => {
  const success = props.gateway.successCount || 0
  const failure = props.gateway.failureCount || 0
  const total = success + failure
  return total > 0 ? Math.round((success / total) * 100) : 0
}

const getStatusColor = (status?: string) => {
  switch (status) {
    case 'active': return 'bg-green-500'
    case 'inactive': return 'bg-red-500'
    case 'maintenance': return 'bg-yellow-500'
    default: return 'bg-gray-500'
  }
}

const getStatusTextColor = (status?: string) => {
  switch (status) {
    case 'active': return 'text-green-600'
    case 'inactive': return 'text-red-600'
    case 'maintenance': return 'text-yellow-600'
    default: return 'text-gray-600'
  }
}

const getSuccessRateColor = (rate: number) => {
  if (rate >= 95) return 'bg-green-500'
  if (rate >= 80) return 'bg-yellow-500'
  return 'bg-red-500'
}

const formatTime = (date: Date | string | undefined) => {
  if (!date) return 'نامشخص'
  const dateObj = typeof date === 'string' ? new Date(date) : date
  return new Intl.DateTimeFormat('fa-IR', {
    hour: '2-digit',
    minute: '2-digit'
  }).format(dateObj)
}

const formatBalance = (balance: number | undefined) => {
  if (balance === undefined || balance === null) return '0'
  // نمایش تمام ارقام موجودی با جداکننده کاما
  return new Intl.NumberFormat('en-US', {
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(balance)
}

// دریافت موجودی درگاه از API
const fetchGatewayBalance = async (id: number) => {
  isLoadingBalance.value = true
  try {
    const response = await $fetch<{status: string, data: {balance: number}}>(`/api/sms-gateways/${id}/balance`)
    if (response.status === 'success' && response.data) {
      gatewayBalance.value = response.data.balance || 0
    } else {
      gatewayBalance.value = 0
    }
  } catch (error) {
    // خطا در دریافت موجودی درگاه
    gatewayBalance.value = 0
  } finally {
    isLoadingBalance.value = false
  }
}

// دریافت اطلاعات کامل ملی پیامک
const fetchMeliPayamakInfo = async (id: number) => {
  isLoadingBalance.value = true
  try {
    const response = await $fetch<{status: string, data: {remaining_sms: number, credit: number}}>(`/api/sms-gateways/${id}/melipayamak-info`)
    if (response.status === 'success' && response.data) {
      meliPayamakInfo.value = {
        remaining_sms: response.data.remaining_sms || 0,
        credit: response.data.credit || 0
      }
    } else {
      meliPayamakInfo.value = { remaining_sms: 0, credit: 0 }
    }
  } catch (error) {
    // خطا در دریافت اطلاعات ملی پیامک
    meliPayamakInfo.value = { remaining_sms: 0, credit: 0 }
  } finally {
    isLoadingBalance.value = false
  }
}

const testGateway = async (id: number) => {
  try {
    const response = await $fetch<{status: string, message?: string, error_message?: string}>(`/api/sms-gateways/${id}/test-connection`, {
      method: 'POST'
    })
    
    if (response.status === 'success') {
      alert('✅ تست اتصال موفق بود!\n\n' + (response.message || 'درگاه به درستی پیکربندی شده و قابل استفاده است.'))
    } else {
      alert('❌ خطا در تست اتصال!\n\n' + (response.error_message || 'لطفاً تنظیمات درگاه را بررسی کنید.'))
    }
  } catch (error: any) {
    let errorMessage = 'خطا در تست اتصال!'
    
    // استخراج پیام خطا از response
    if (error.response && error.response._data) {
      const errorData = error.response._data
      if (errorData.error_message) {
        errorMessage = errorData.error_message
      } else if (errorData.message) {
        errorMessage = errorData.message
      }
    }
    
    alert('❌ خطا در تست اتصال!\n\n' + errorMessage)
  }
}

const testSendSMS = async (id: number) => {
  // درخواست شماره موبایل و متن پیام از کاربر
  const mobile = prompt('لطفاً شماره موبایل گیرنده را وارد کنید:')
  if (!mobile) return
  
  const message = prompt('لطفاً متن پیام را وارد کنید:')
  if (!message) return
  
  try {
    const response = await $fetch(`/api/sms-gateways/${id}/send-test`, {
      method: 'POST',
      body: {
        mobile: mobile,
        message: message
      }
    })
    alert('پیامک با موفقیت ارسال شد!')
  } catch (error) {
    // خطا در ارسال پیامک
    alert('خطا در ارسال پیامک!')
  }
}

const toggleGateway = async (id: number) => {
  try {
    const currentStatus = props.gateway.status === 'active' || props.gateway.is_active
    const newStatus = currentStatus ? 'inactive' : 'active'
    const newIsActive = !currentStatus
    
    await $fetch(`/api/sms-gateways/${id}`, {
      method: 'PUT',
      body: {
        ...props.gateway,
        status: newStatus,
        is_active: newIsActive
      }
    })
    // به‌روزرسانی وضعیت در parent component
    props.gateway.status = newStatus
    props.gateway.is_active = newIsActive
    alert(`درگاه ${newIsActive ? 'فعال' : 'غیرفعال'} شد!`)
  } catch (error) {
    alert('خطا در تغییر وضعیت درگاه!')
  }
}

// حذف سخت درگاه
const deleteGateway = async (id: number) => {
  // تأیید حذف از کاربر
  const confirmed = confirm(`آیا مطمئن هستید که می‌خواهید درگاه "${props.gateway.name || 'بدون نام'}" را به طور کامل حذف کنید؟\n\nاین عملیات غیرقابل بازگشت است!`)
  
  if (!confirmed) {
    return
  }
  
  try {
    await $fetch(`/api/sms-gateways/${id}`, {
      method: 'DELETE'
    })
    
    // اطلاع‌رسانی به parent component
    emit('gatewayDeleted', id)
    alert('درگاه با موفقیت حذف شد!')
  } catch (error) {
    // خطا در حذف درگاه
    alert('خطا در حذف درگاه! لطفاً دوباره تلاش کنید.')
  }
}

onMounted(() => {
  // دریافت موجودی درگاه در بارگذاری
  fetchGatewayBalance(props.gateway.id)
  
  // به‌روزرسانی زمان آخرین بررسی هر 60 ثانیه
  setInterval(() => {
    if (props.gateway.lastCheck) {
      props.gateway.lastCheck = new Date()
    }
  }, 60000)
  
  // به‌روزرسانی موجودی هر 5 دقیقه
  setInterval(() => {
    fetchGatewayBalance(props.gateway.id)
  }, 300000)
})
</script>

<style scoped>
.space-x-reverse > :not([hidden]) ~ :not([hidden]) {
  --tw-space-x-reverse: 1;
}
</style> 
