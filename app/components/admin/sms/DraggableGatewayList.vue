<template>
  <div class="space-y-4">
    <!-- پیام وضعیت -->
    <div v-if="isUpdating" class="bg-blue-50 border border-blue-200 rounded-lg p-6">
      <div class="flex items-center">
        <svg class="animate-spin -ml-1 mr-3 h-5 w-5 text-blue-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
        </svg>
        <span class="text-blue-800">در حال به‌روزرسانی اولویت‌ها...</span>
      </div>
    </div>

    <!-- لیست درگاه‌ها -->
    <div class="space-y-4">
      <div
        v-for="(gateway, index) in sortedGateways"
        :key="gateway.id"
        :data-id="gateway.id"
        class="relative"
        draggable="true"
        @dragstart="onDragStart($event, index)"
        @dragover.prevent
        @drop="onDrop($event, index)"
        @dragenter.prevent
        @dragleave.prevent
      >
        <!-- نشانگر اولویت -->
        <div class="absolute top-2 right-2 z-10">
          <div class="bg-purple-600 text-white text-xs font-bold px-2 py-1 rounded-full">
            {{ gateway.priority || index + 1 }}
          </div>
        </div>

        <!-- کارت درگاه -->
        <div
          class="rounded-xl shadow-lg p-6 bg-white border-2 transition-all duration-200 cursor-move"
          :class="{
            'border-purple-300 shadow-xl bg-purple-50': draggedIndex === index,
            'border-gray-200 hover:border-purple-200': draggedIndex !== index
          }"
          :style="{
            transform: draggedIndex === index ? 'scale(1.02) rotate(2deg)' : 'scale(1)',
            opacity: draggedIndex === index ? '0.9' : '1'
          }"
        >

          
          <!-- کارت وضعیت درگاه -->
          <div class="rounded-lg p-6 border border-gray-200 min-h-[300px] flex flex-col justify-between bg-white">
            <div>
              <div class="flex items-center justify-between mb-4">
                <div class="flex items-center">
                  <div class="w-4 h-4 rounded-full ml-8" :class="getStatusColor(gateway.status)"></div>
                  <div>
                    <h4 class="font-semibold text-gray-800">{{ gateway.name || 'درگاه بدون نام' }}</h4>
                  </div>
                </div>
                <div class="text-right">
                  <p class="text-sm font-semibold text-gray-600">آخرین بررسی: {{ formatTime(gateway.lastCheck) }}</p>
                </div>
              </div>
            </div>

            <!-- آمار و دکمه‌ها در پایین -->
            <div>
              <!-- آمار و نرخ موفقیت -->
              <!-- برای ملی پیامک: نمایش 4 ستون -->
              <div v-if="gateway.type === 'meli_payamak'" class="grid grid-cols-4 gap-6 mb-2">
                <div class="text-center">
                  <p class="text-2xl font-bold" :class="getStatusTextColor(gateway.status)">{{ gateway.successCount || 0 }}</p>
                  <p class="text-xs text-gray-600">موفق</p>
                </div>
                <div class="text-center">
                  <p class="text-2xl font-bold text-red-600">{{ gateway.failureCount || 0 }}</p>
                  <p class="text-xs text-gray-600">ناموفق</p>
                </div>
                <div class="text-center cursor-pointer" title="کلیک برای به‌روزرسانی اطلاعات" @click="fetchMeliPayamakInfo(gateway.id)">
                  <div v-if="isLoadingBalance[gateway.id]" class="inline-block animate-spin rounded-full h-6 w-6 border-b-2 border-blue-600"></div>
                  <div v-else>
                    <p class="text-2xl font-bold text-blue-600">{{ meliPayamakInfos[gateway.id]?.remaining_sms || 0 }}</p>
                    <p class="text-xs text-gray-600">SMS باقی‌مانده</p>
                  </div>
                </div>
                <div class="text-center cursor-pointer" title="کلیک برای به‌روزرسانی اطلاعات" @click="fetchMeliPayamakInfo(gateway.id)">
                  <div v-if="isLoadingBalance[gateway.id]" class="inline-block animate-spin rounded-full h-6 w-6 border-b-2 border-green-600"></div>
                  <div v-else>
                    <p class="text-2xl font-bold text-green-600">{{ formatBalance(meliPayamakInfos[gateway.id]?.credit || 0) }}</p>
                    <p class="text-xs text-gray-600">موجودی ریالی</p>
                  </div>
                </div>
              </div>
              
              <!-- برای سایر درگاه‌ها: نمایش 3 ستون -->
              <div v-else class="grid grid-cols-3 gap-6 mb-2">
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
                    <span v-if="isLoadingBalance[gateway.id]" class="inline-block animate-spin rounded-full h-6 w-6 border-b-2 border-blue-600"></span>
                    <span v-else>{{ formatBalance(gatewayBalances[gateway.id]) }}</span>
                  </p>
                  <p class="text-xs text-gray-600">موجودی ریالی</p>
                </div>
              </div>
              <div class="mb-2">
                <div class="flex justify-between text-xs text-gray-600 mb-1">
                  <span>نرخ موفقیت</span>
                  <span>{{ calculateSuccessRate(gateway) }}%</span>
                </div>
                <div class="w-full bg-gray-200 rounded-full h-2">
                  <div
class="h-2 rounded-full transition-all duration-300" 
                       :class="getSuccessRateColor(calculateSuccessRate(gateway))"
                       :style="{ width: calculateSuccessRate(gateway) + '%' }"></div>
                </div>
              </div>

              <!-- دکمه‌ها -->
              <div class="flex flex-col md:flex-row-reverse gap-6 items-stretch">
                <button class="w-full md:w-48 px-10 py-3 rounded-lg text-base font-bold bg-gradient-to-r from-red-600 to-red-800 text-white shadow-md hover:from-red-700 hover:to-red-900 transition flex items-center justify-center gap-2" @click="deleteGateway(gateway.id)">
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" /></svg>
                  حذف
                </button>
                <button
class="w-full md:w-48 px-10 py-3 rounded-lg text-base font-bold shadow-md transition flex items-center justify-center gap-2" :class="(gateway.status === 'active' || gateway.is_active) ? 'bg-gradient-to-r from-yellow-500 to-yellow-700 text-white hover:from-yellow-600 hover:to-yellow-800' : 'bg-gradient-to-r from-green-400 to-green-600 text-white hover:from-green-500 hover:to-green-700'"
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
                                 <button
class="w-full md:w-auto px-6 py-3 rounded-lg text-base font-bold bg-gradient-to-r from-green-500 to-green-700 text-white hover:from-green-600 hover:to-green-800 shadow-md transition flex items-center justify-center whitespace-nowrap" 
                   title="تست ارسال پیامک"
                   @click="testSendSMS(gateway.id)">
                   <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" /></svg>تست ارسال
                 </button>
                <button class="w-full md:w-auto px-6 py-3 rounded-lg text-base font-bold bg-gradient-to-r from-blue-500 to-blue-700 text-white shadow-md hover:from-blue-600 hover:to-blue-800 transition flex items-center justify-center whitespace-nowrap" @click="testGateway(gateway.id)">
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8h2a2 2 0 012 2v8a2 2 0 01-2 2H5a2 2 0 01-2-2v-8a2 2 0 012-2h2" /><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 3h-6a2 2 0 00-2 2v3h10V5a2 2 0 00-2-2z" /></svg>تست اتصال
                </button>
              </div>
            </div>
          </div>
          

        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'

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
  [key: string]: unknown
}

// تعریف props برای دریافت gateways از parent component
const props = defineProps<{
  gateways: Gateway[]
}>()

// تعریف emit برای اطلاع‌رسانی به parent component
const emit = defineEmits<{
  gatewayDeleted: [id: number]
  prioritiesUpdated: [priorities: Array<{id: number, priority: number}>]
}>()

// متغیرهای reactive
const draggedIndex = ref<number | null>(null)
const isUpdating = ref<boolean>(false)
const gatewayBalances = ref<Record<number, number>>({})
const isLoadingBalance = ref<Record<number, boolean>>({})

// متغیر reactive برای اطلاعات ملی پیامک
const meliPayamakInfos = ref<Record<number, {remaining_sms: number, credit: number}>>({})

// بررسی وجود پترن تست برای هر درگاه
const _checkTestPatterns = async () => {
  // این قابلیت غیرفعال شده است
}

// مرتب‌سازی درگاه‌ها بر اساس اولویت
const sortedGateways = computed(() => {
  return [...props.gateways].sort((a, b) => (a.priority || 0) - (b.priority || 0))
})

// دریافت درگاه پیش‌فرض (اولین درگاه فعال با کمترین اولویت)
const _getDefaultGateway = () => {
  const activeGateways = sortedGateways.value.filter(g => g.is_active || g.status === 'active')
  return activeGateways.length > 0 ? activeGateways[0] : null
}

// محاسبه نرخ موفقیت
const calculateSuccessRate = (gateway: Gateway) => {
  const success = gateway.successCount || 0
  const failure = gateway.failureCount || 0
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
  return new Intl.NumberFormat('en-US', {
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(balance)
}

// Drag and Drop functions
const onDragStart = (event: DragEvent, index: number) => {
  if (event.dataTransfer) {
    event.dataTransfer.effectAllowed = 'move'
    event.dataTransfer.setData('text/html', index.toString())
    draggedIndex.value = index
  }
}

const onDrop = async (event: DragEvent, dropIndex: number) => {
  event.preventDefault()
  
  if (draggedIndex.value === null || draggedIndex.value === dropIndex) {
    draggedIndex.value = null
    return
  }

  const draggedGateway = sortedGateways.value[draggedIndex.value]
  const newGateways = [...sortedGateways.value]
  
  // حذف آیتم از موقعیت قبلی
  newGateways.splice(draggedIndex.value, 1)
  // اضافه کردن آیتم در موقعیت جدید
  newGateways.splice(dropIndex, 0, draggedGateway)
  
  // به‌روزرسانی اولویت‌ها
  const priorities = newGateways
    .filter(g => g && typeof g.id === 'number')
    .map((gateway, index) => ({ id: Number(gateway.id), priority: index + 1 }))
  
  // به‌روزرسانی اولویت‌ها
  
  try {
    isUpdating.value = true
    
    await $fetch('/api/sms-gateways/priorities', {
      method: 'PUT',
      body: { priorities }
    })
    
    // اطلاع‌رسانی به parent component بدون refresh
    emit('prioritiesUpdated', priorities)
    
    // نمایش پیام موفقیت
    alert('اولویت‌ها با موفقیت به‌روزرسانی شدند')
  } catch (error) {
    // خطا در به‌روزرسانی اولویت‌ها
    alert('خطا در به‌روزرسانی اولویت‌ها: ' + (error as Error).message)
  } finally {
    isUpdating.value = false
    draggedIndex.value = null
  }
}

// دریافت موجودی درگاه
const fetchGatewayBalance = async (gatewayId: number) => {
  if (isLoadingBalance.value[gatewayId]) return
  
  try {
    isLoadingBalance.value[gatewayId] = true
    const response = await $fetch<{status: string, data: {balance: number}}>(`/api/sms-gateways/${gatewayId}/balance`)
    
    if (response.status === 'success') {
      gatewayBalances.value[gatewayId] = response.data.balance
    }
  } catch (_error) {
    // خطا در دریافت موجودی
  } finally {
    isLoadingBalance.value[gatewayId] = false
  }
}

// دریافت اطلاعات کامل ملی پیامک
const fetchMeliPayamakInfo = async (gatewayId: number) => {
  if (isLoadingBalance.value[gatewayId]) return
  
  try {
    isLoadingBalance.value[gatewayId] = true
    const response = await $fetch<{status: string, data: {remaining_sms: number, credit: number}}>(`/api/sms-gateways/${gatewayId}/melipayamak-info`)
    
    if (response.status === 'success') {
      meliPayamakInfos.value[gatewayId] = {
        remaining_sms: response.data.remaining_sms || 0,
        credit: response.data.credit || 0
      }
    }
  } catch (_error) {
    // خطا در دریافت اطلاعات ملی پیامک
  } finally {
    isLoadingBalance.value[gatewayId] = false
  }
}

// تست اتصال درگاه
const testGateway = async (gatewayId: number) => {
  try {
    const response = await $fetch<{status: string, message?: string, error_message?: string}>(`/api/sms-gateways/${gatewayId}/test-connection`, {
      method: 'POST'
    })
    
    if (response.status === 'success') {
      alert('✅ اتصال درگاه با موفقیت تست شد!\n\n' + (response.message || 'درگاه به درستی پیکربندی شده و قابل استفاده است.'))
    } else {
      alert('❌ خطا در تست اتصال درگاه!\n\n' + (response.error_message || 'لطفاً تنظیمات درگاه را بررسی کنید.'))
    }
  } catch (error: unknown) {
    // خطا در تست اتصال
    let errorMessage = 'خطا در تست اتصال درگاه'
    
    // استخراج پیام خطا از response
    interface ErrorResponse {
      response?: {
        _data?: {
          error_message?: string
          message?: string
        }
      }
    }
    
    if (error && typeof error === 'object' && 'response' in error) {
      const errorResponse = error as ErrorResponse
      const errorData = errorResponse.response?._data
      if (errorData?.error_message) {
        errorMessage = errorData.error_message
      } else if (errorData?.message) {
        errorMessage = errorData.message
      }
    }
    
    alert('❌ خطا در تست اتصال درگاه!\n\n' + errorMessage)
  }
}

// تست ارسال SMS
const testSendSMS = async (gatewayId: number) => {
  try {
    // دریافت شماره موبایل مدیر از تنظیمات فروشگاه
    let adminPhone = ''
    try {
      const settingsResponse = await $fetch<{status: string, data: {adminPhones: string[]}}>('/api/admin/shop-settings')
      if (settingsResponse.status === 'success' && settingsResponse.data?.adminPhones && settingsResponse.data.adminPhones.length > 0) {
        adminPhone = settingsResponse.data.adminPhones[0]
      }
    } catch (error) {
      console.error('خطا در دریافت شماره موبایل مدیر:', error)
    }

         // اگر شماره موبایل مدیر موجود نباشد، از کاربر بخواهیم
     if (!adminPhone) {
       adminPhone = prompt('لطفاً شماره موبایل خود را وارد کنید:')
       if (!adminPhone) {
         alert('شماره موبایل الزامی است')
         return
       }
     }

     // پیدا کردن اطلاعات درگاه
     const gateway = props.gateways.find(g => g.id === gatewayId)
     const gatewayName = gateway?.name || `درگاه ${gatewayId}`

     // ارسال پیامک تست با استفاده از پترن تست ارسال درگاه
     const response = await $fetch<{success: boolean, message?: string}>('/api/admin/sms-patterns/send-by-scope-feature', {
       method: 'POST',
       body: {
         mobile: adminPhone,
         scope: 'manager',
         feature: 'gateway_test',
         gateway_id: gatewayId,
         variables: {
           gateway: gatewayName,
           date: new Date().toLocaleDateString('fa-IR') + ' ' + new Date().toLocaleTimeString('fa-IR')
         }
       }
     })

    if (response.success) {
      alert('✅ پیامک تست با موفقیت ارسال شد!\n\n' + (response.message || 'پیامک تست به شماره ' + adminPhone + ' ارسال شد.'))
    } else {
      alert('❌ خطا در ارسال پیامک تست!\n\n' + (response.message || 'لطفاً تنظیمات درگاه و پترن تست را بررسی کنید.'))
    }
  } catch (error: unknown) {
    // خطا در ارسال پیامک تست
    let errorMessage = 'خطا در ارسال پیامک تست'
    
    // استخراج پیام خطا از response
    interface ErrorResponse {
      response?: {
        _data?: {
          message?: string
        }
      }
    }
    
    if (error && typeof error === 'object' && 'response' in error) {
      const errorResponse = error as ErrorResponse
      const errorData = errorResponse.response?._data
      if (errorData?.message) {
        errorMessage = errorData.message
      }
    }
    
    alert('❌ خطا در ارسال پیامک تست!\n\n' + errorMessage)
  }
}

// فعال/غیرفعال کردن درگاه
const toggleGateway = async (gatewayId: number) => {
  try {
    await $fetch(`/api/sms-gateways/${gatewayId}`, {
      method: 'PATCH',
      body: {
        is_active: false // تغییر وضعیت
      }
    })
    
    alert('وضعیت درگاه با موفقیت تغییر یافت')
    // به جای reload، emit event
    emit('gatewayDeleted', gatewayId)
  } catch (_error) {
    // خطا در تغییر وضعیت درگاه
    alert('خطا در تغییر وضعیت درگاه')
  }
}

// حذف درگاه
const deleteGateway = async (gatewayId: number) => {
  if (!confirm('آیا مطمئن هستید که می‌خواهید این درگاه را حذف کنید؟')) {
    return
  }
  
  try {
    await $fetch(`/api/sms-gateways/${gatewayId}`, {
      method: 'DELETE'
    })
    
    alert('درگاه با موفقیت حذف شد')
    emit('gatewayDeleted', gatewayId)
  } catch (_error) {
    // خطا در حذف درگاه
    alert('خطا در حذف درگاه')
  }
}

// دریافت موجودی همه درگاه‌ها در ابتدا
onMounted(async () => {


  sortedGateways.value.forEach(gateway => {
    // برای ملی پیامک از endpoint جدید استفاده کن
    if (gateway.type === 'meli_payamak') {
      fetchMeliPayamakInfo(gateway.id)
    } else {
      fetchGatewayBalance(gateway.id)
    }
  })
})

// وقتی درگاه‌ها تغییر کردند، دوباره بررسی کن
watch(() => props.gateways, async () => {

}, { deep: true })
</script>

<style scoped>
.space-x-reverse > :not([hidden]) ~ :not([hidden]) {
  --tw-space-x-reverse: 1;
}
</style> 