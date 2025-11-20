<template>
  <div class="space-y-6">
    <!-- اطلاعات درگاه پرداخت -->
    <div class="px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-blue-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"></path>
        </svg>
        اطلاعات درگاه پرداخت
      </h3>

      <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
        <!-- درگاه اصلی -->
        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">درگاه پرداخت اصلی</h4>
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">نوع درگاه</label>
              <select 
                v-model="paymentGateways.primaryGateway"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="">انتخاب درگاه</option>
                <option value="zarinpal">زرین‌پال</option>
                <option value="mellat">ملت</option>
                <option value="parsian">پارسیان</option>
                <option value="saman">سامان</option>
                <option value="pasargad">پاسارگاد</option>
                <option value="saderat">صادرات</option>
                <option value="melli">ملی</option>
                <option value="sepah">سپه</option>
                <option value="keshavarzi">کشاورزی</option>
                <option value="tejarat">تجارت</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">شماره تراکنش</label>
              <input 
                v-model="paymentGateways.transactionId"
                type="text" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="شماره تراکنش درگاه"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">شماره ارجاع</label>
              <input 
                v-model="paymentGateways.referenceNumber"
                type="text" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="شماره ارجاع بانکی"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">وضعیت تراکنش</label>
              <select 
                v-model="paymentGateways.transactionStatus"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="pending">در انتظار پرداخت</option>
                <option value="processing">در حال پردازش</option>
                <option value="completed">تکمیل شده</option>
                <option value="failed">ناموفق</option>
                <option value="cancelled">لغو شده</option>
                <option value="refunded">بازگشت وجه</option>
              </select>
            </div>
          </div>
        </div>

        <!-- اطلاعات پرداخت -->
        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">جزئیات پرداخت</h4>
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">مبلغ تراکنش</label>
              <input 
                v-model="paymentGateways.transactionAmount"
                type="number" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="0"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">کارمزد درگاه</label>
              <input 
                v-model="paymentGateways.gatewayFee"
                type="number" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="0"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">تاریخ تراکنش</label>
              <input 
                v-model="paymentGateways.transactionDate"
                type="datetime-local" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">کد خطا (در صورت وجود)</label>
              <input 
                v-model="paymentGateways.errorCode"
                type="text" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="کد خطای درگاه"
              />
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- درگاه‌های جایگزین -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-green-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7h12m0 0l-4-4m4 4l-4 4m0 6H4m0 0l4 4m-4-4l4-4"></path>
        </svg>
        درگاه‌های جایگزین
      </h3>

      <div class="space-y-4">
        <div 
          v-for="(gateway, index) in paymentGateways.alternativeGateways" 
          :key="index"
          class="border border-gray-200 rounded-lg px-4 py-4"
        >
          <div class="grid grid-cols-1 md:grid-cols-3 gapx-4 py-4">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-1">درگاه جایگزین</label>
              <select 
                v-model="gateway.name"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="">انتخاب درگاه</option>
                <option value="zarinpal">زرین‌پال</option>
                <option value="mellat">ملت</option>
                <option value="parsian">پارسیان</option>
                <option value="saman">سامان</option>
                <option value="pasargad">پاسارگاد</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-1">وضعیت</label>
              <select 
                v-model="gateway.status"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="available">در دسترس</option>
                <option value="unavailable">غیرفعال</option>
                <option value="maintenance">در حال تعمیر</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-1">اولویت</label>
              <input 
                v-model="gateway.priority"
                type="number" 
                min="1"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="1"
              />
            </div>
          </div>
          <div class="mt-3">
            <label class="block text-sm font-medium text-gray-600 mb-1">یادداشت</label>
            <textarea 
              v-model="gateway.notes"
              rows="2"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="یادداشت مربوط به درگاه..."
            ></textarea>
          </div>
        </div>

        <!-- دکمه افزودن درگاه جایگزین -->
        <div class="mt-4">
          <button 
            class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors flex items-center text-sm"
            @click="addAlternativeGateway"
          >
            <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
            </svg>
            افزودن درگاه جایگزین
          </button>
        </div>
      </div>
    </div>

    <!-- تنظیمات امنیتی -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-red-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"></path>
        </svg>
        تنظیمات امنیتی
      </h3>

      <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">احراز هویت</h4>
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">نوع احراز هویت</label>
              <select 
                v-model="paymentGateways.authenticationType"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="sms">پیامک</option>
                <option value="email">ایمیل</option>
                <option value="app">اپلیکیشن</option>
                <option value="card">کارت بانکی</option>
                <option value="biometric">احراز هویت زیستی</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">وضعیت احراز هویت</label>
              <select 
                v-model="paymentGateways.authenticationStatus"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="pending">در انتظار</option>
                <option value="verified">تایید شده</option>
                <option value="failed">ناموفق</option>
                <option value="expired">منقضی شده</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">تاریخ احراز هویت</label>
              <input 
                v-model="paymentGateways.authenticationDate"
                type="datetime-local" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
          </div>
        </div>

        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">امنیت تراکنش</h4>
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">سطح امنیت</label>
              <select 
                v-model="paymentGateways.securityLevel"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="low">کم</option>
                <option value="medium">متوسط</option>
                <option value="high">بالا</option>
                <option value="maximum">حداکثر</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">وضعیت فیلتر امنیتی</label>
              <select 
                v-model="paymentGateways.securityFilterStatus"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="active">فعال</option>
                <option value="inactive">غیرفعال</option>
                <option value="suspicious">مشکوک</option>
                <option value="blocked">مسدود</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">IP مشتری</label>
              <input 
                v-model="paymentGateways.clientIP"
                type="text" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="192.168.1.1"
              />
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- لاگ‌های تراکنش -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-orange-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
        </svg>
        لاگ‌های تراکنش
      </h3>
      
      <div class="space-y-3">
        <div 
          v-for="(log, index) in paymentGateways.transactionLogs" 
          :key="index"
          class="flex items-center p-3 bg-gray-50 rounded-lg"
        >
          <div class="w-3 h-3 rounded-full mr-3" :class="getLogStatusColor(log.status)"></div>
          <div class="flex-1">
            <div class="text-sm font-medium text-gray-900">{{ log.action }}</div>
            <div class="text-xs text-gray-500">{{ formatDate(log.date) }} - {{ log.gateway }}</div>
            <div v-if="log.details" class="text-xs text-gray-600 mt-1">{{ log.details }}</div>
          </div>
          <div class="text-right">
            <div class="text-xs text-gray-500">{{ formatTime(log.date) }}</div>
            <div class="text-xs text-gray-500">{{ log.status }}</div>
          </div>
        </div>
      </div>

      <!-- دکمه افزودن لاگ -->
      <div class="mt-4">
        <button 
          class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors flex items-center text-sm"
          @click="addTransactionLog"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
          </svg>
          افزودن لاگ تراکنش
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'

// تعریف props
const props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({
      primaryGateway: 'zarinpal',
      transactionId: '',
      referenceNumber: '',
      transactionStatus: 'completed',
      transactionAmount: 0,
      gatewayFee: 0,
      transactionDate: '',
      errorCode: '',
      alternativeGateways: [],
      authenticationType: 'sms',
      authenticationStatus: 'verified',
      authenticationDate: '',
      securityLevel: 'high',
      securityFilterStatus: 'active',
      clientIP: '',
      transactionLogs: []
    })
  }
})

// تعریف emit
const emit = defineEmits(['update:modelValue'])

// متغیر محلی
const paymentGateways = ref({ ...props.modelValue })

// توابع کمکی
const getLogStatusColor = (status) => {
  const colors = {
    'success': 'bg-green-400',
    'error': 'bg-red-400',
    'warning': 'bg-yellow-400',
    'info': 'bg-blue-400',
    'pending': 'bg-gray-400'
  }
  return colors[status] || 'bg-gray-400'
}

const formatDate = (date) => {
  if (!date) return 'نامشخص'
  return new Date(date).toLocaleDateString('fa-IR')
}

const formatTime = (date) => {
  if (!date) return ''
  return new Date(date).toLocaleTimeString('fa-IR', { 
    hour: '2-digit', 
    minute: '2-digit' 
  })
}

const addAlternativeGateway = () => {
  paymentGateways.value.alternativeGateways.push({
    name: '',
    status: 'available',
    priority: 1,
    notes: ''
  })
}

const addTransactionLog = () => {
  const newEntry = {
    action: 'تراکنش جدید',
    date: new Date().toISOString(),
    gateway: paymentGateways.value.primaryGateway,
    status: 'success',
    details: 'ثبت تراکنش جدید'
  }
  paymentGateways.value.transactionLogs.unshift(newEntry)
}

// نظارت بر تغییرات
watch(paymentGateways, (newValue) => {
  emit('update:modelValue', newValue)
}, { deep: true })
</script> 
