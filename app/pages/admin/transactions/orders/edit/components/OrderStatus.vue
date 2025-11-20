<template>
  <div class="space-y-6">
    <!-- وضعیت فعلی سفارش -->
    <div class="px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-purple-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
        </svg>
        وضعیت فعلی سفارش
      </h3>

      <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
        <!-- وضعیت اصلی -->
        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">وضعیت اصلی</h4>
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">وضعیت سفارش</label>
              <select 
                v-model="orderStatus.status"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="processing_queue">در صف پردازش</option>
                <option value="awaiting_payment">در انتظار پرداخت</option>
                <option value="pending_review">در انتظار بررسی</option>
                <option value="processing">در حال پردازش</option>
                <option value="ready_to_ship">آماده ارسال</option>
                <option value="shipped">ارسال شده</option>
                <option value="delivered">تحویل داده شده</option>
                <option value="cancelled">لغو شده</option>
                <option value="returned">بازگشت</option>
                <option value="refunded">بازپرداخت شده</option>
                <option value="on_hold">معلق</option>
                <option value="fraud">مشکوک به تقلب</option>
              </select>
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">اولویت سفارش</label>
              <select 
                v-model="orderStatus.priority"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="low">کم</option>
                <option value="normal">عادی</option>
                <option value="high">بالا</option>
                <option value="urgent">فوری</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">دلیل تغییر وضعیت</label>
              <select 
                v-model="orderStatus.statusReason"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="">انتخاب دلیل</option>
                <option value="customer_request">درخواست مشتری</option>
                <option value="stock_issue">مشکل موجودی</option>
                <option value="payment_issue">مشکل پرداخت</option>
                <option value="shipping_issue">مشکل ارسال</option>
                <option value="quality_issue">مشکل کیفیت</option>
                <option value="fraud_detection">تشخیص تقلب</option>
                <option value="system_error">خطای سیستم</option>
                <option value="other">سایر</option>
              </select>
            </div>
          </div>
        </div>

        <!-- اطلاعات پرداخت -->
        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">اطلاعات پرداخت</h4>
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">نحوه پرداخت</label>
              <select 
                v-model="orderStatus.paymentMethod"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="online">پرداخت آنلاین</option>
                <option value="cash">پرداخت نقدی</option>
                <option value="bank_transfer">انتقال بانکی</option>
                <option value="wallet">کیف پول</option>
                <option value="gift_card">کارت هدیه</option>
                <option value="installment">قسطی</option>
                <option value="cod">پرداخت در محل</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">وضعیت پرداخت</label>
              <select 
                v-model="orderStatus.paymentStatus"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="awaiting_payment">در انتظار پرداخت</option>
                <option value="partial">پرداخت جزئی</option>
                <option value="paid">پرداخت شده</option>
                <option value="failed">ناموفق</option>
                <option value="refunded">بازگشت وجه</option>
                <option value="disputed">مورد مناقشه</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">تاریخ پرداخت</label>
              <input 
                v-model="orderStatus.paymentDate"
                type="datetime-local" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
          </div>
        </div>
      </div>
    </div>



    <!-- اطلاعات بازگشت -->
    <div v-if="orderStatus.status === 'returned' || orderStatus.status === 'partially_returned'" class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-red-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h10a8 8 0 018 8v2M3 10l6 6m-6-6l6-6"></path>
        </svg>
        اطلاعات بازگشت
      </h3>

      <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">جزئیات بازگشت</h4>
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">دلیل بازگشت</label>
              <select 
                v-model="orderStatus.returnReason"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="">انتخاب دلیل</option>
                <option value="wrong_size">سایز اشتباه</option>
                <option value="defective">معیوب</option>
                <option value="not_as_described">مطابق توضیحات نبود</option>
                <option value="changed_mind">تغییر نظر</option>
                <option value="duplicate_order">سفارش تکراری</option>
                <option value="shipping_damage">آسیب در حمل</option>
                <option value="late_delivery">تاخیر در تحویل</option>
                <option value="other">سایر</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">وضعیت بازگشت</label>
              <select 
                v-model="orderStatus.returnStatus"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="pending">در انتظار بررسی</option>
                <option value="approved">تایید شده</option>
                <option value="rejected">رد شده</option>
                <option value="completed">تکمیل شده</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">تاریخ درخواست بازگشت</label>
              <input 
                v-model="orderStatus.returnRequestDate"
                type="datetime-local" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
          </div>
        </div>

        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">بازپرداخت</h4>
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">مبلغ بازپرداخت</label>
              <input 
                v-model="orderStatus.refundAmount"
                type="number" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="0"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">روش بازپرداخت</label>
              <select 
                v-model="orderStatus.refundMethod"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="original_method">روش اصلی پرداخت</option>
                <option value="wallet">کیف پول</option>
                <option value="gift_card">کارت هدیه</option>
                <option value="bank_transfer">انتقال بانکی</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">تاریخ بازپرداخت</label>
              <input 
                v-model="orderStatus.refundDate"
                type="datetime-local" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تاریخچه تغییرات وضعیت -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-orange-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
        </svg>
        تاریخچه تغییرات وضعیت
      </h3>
      
      <div class="space-y-3">
        <div 
          v-for="(history, index) in orderStatus.history" 
          :key="index"
          class="flex items-center p-3 bg-gray-50 rounded-lg"
        >
          <div class="w-3 h-3 rounded-full mr-3" :class="getStatusColor(history.status)"></div>
          <div class="flex-1">
            <div class="text-sm font-medium text-gray-900">{{ getStatusText(history.status) }}</div>
            <div class="text-xs text-gray-500">{{ formatDate(history.date) }} - {{ history.user }}</div>
            <div v-if="history.note" class="text-xs text-gray-600 mt-1">{{ history.note }}</div>
          </div>
          <div class="text-xs text-gray-500">{{ formatTime(history.date) }}</div>
        </div>
      </div>

      <!-- دکمه افزودن تاریخچه -->
      <div class="mt-4">
        <button 
          class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors flex items-center text-sm"
          @click="addHistoryEntry"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
          </svg>
          افزودن تاریخچه
        </button>
      </div>
    </div>

    <!-- دکمه‌های عملیات -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <div class="flex justify-between items-center">
        <h3 class="text-lg font-semibold text-gray-900 flex items-center">
          <svg class="w-5 h-5 text-blue-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"></path>
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
          </svg>
          عملیات وضعیت
        </h3>
        
        <div class="flex space-x-3 space-x-reverse">
          <button 
            :disabled="isLoading"
            class="px-4 py-2 bg-gray-500 text-white rounded-lg hover:bg-gray-600 transition-colors flex items-center text-sm disabled:opacity-50"
            @click="resetChanges"
          >
            <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
            </svg>
            بازنشانی
          </button>
          
          <button 
            :disabled="isLoading || !hasChanges"
            class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors flex items-center text-sm disabled:opacity-50"
            @click="saveStatusChanges"
          >
            <svg v-if="isLoading" class="w-4 h-4 ml-2 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
            </svg>
            <svg v-else class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
            </svg>
            {{ isLoading ? 'در حال ذخیره...' : 'ذخیره تغییرات' }}
          </button>
        </div>
      </div>
    </div>

    <!-- یادداشت و توضیحات -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-green-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
        </svg>
        یادداشت و توضیحات
      </h3>
      
      <div class="space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-600 mb-2">یادداشت داخلی</label>
          <textarea 
            v-model="orderStatus.internalNotes"
            rows="3"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            placeholder="یادداشت‌های داخلی برای کارکنان..."
          ></textarea>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-600 mb-2">یادداشت مشتری</label>
          <textarea 
            v-model="orderStatus.customerNotes"
            rows="3"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            placeholder="یادداشت‌های قابل مشاهده برای مشتری..."
          ></textarea>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-600 mb-2">یادداشت عمومی</label>
          <textarea 
            v-model="orderStatus.notes"
            rows="3"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            placeholder="یادداشت‌های عمومی سفارش..."
          ></textarea>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, ref, watch } from 'vue'

// تعریف props
const props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({
      status: 'processing_queue',
      priority: 'normal',
      statusReason: '',
      paymentMethod: 'online',
      paymentStatus: 'awaiting_payment',
      paymentDate: '',
      returnReason: '',
      returnStatus: '',
      returnRequestDate: '',
      refundAmount: 0,
      refundMethod: 'original_method',
      refundDate: '',
      internalNotes: '',
      customerNotes: '',
      notes: '',
      history: []
    })
  },
  orderId: {
    type: [String, Number],
    required: true
  }
})

// تعریف emit
const emit = defineEmits(['update:modelValue', 'status-updated'])

// متغیرهای محلی
const orderStatus = ref({ ...props.modelValue })
const originalStatus = ref({ ...props.modelValue })
const isLoading = ref(false)

// محاسبه تغییرات
const hasChanges = computed(() => {
  return orderStatus.value.status !== originalStatus.value.status ||
         orderStatus.value.paymentStatus !== originalStatus.value.paymentStatus ||
         orderStatus.value.priority !== originalStatus.value.priority ||
         orderStatus.value.paymentMethod !== originalStatus.value.paymentMethod
})

// توابع کمکی
const getStatusText = (status) => {
  const statusMap = {
    'processing_queue': 'در صف پردازش',
    'awaiting_payment': 'در انتظار پرداخت',
    'pending_review': 'در انتظار بررسی',
    'processing': 'در حال پردازش',
    'ready_to_ship': 'آماده ارسال',
    'shipped': 'ارسال شده',
    'delivered': 'تحویل داده شده',
    'cancelled': 'لغو شده',
    'returned': 'بازگشت',
    'refunded': 'بازپرداخت شده',
    'on_hold': 'معلق',
    'fraud': 'مشکوک به تقلب',
    'paid': 'پرداخت شده',
    'failed': 'ناموفق',
    'partial': 'پرداخت جزئی',
    'disputed': 'مورد مناقشه',
    'pending': 'در انتظار پرداخت' // برای سازگاری با مقادیر قدیمی
  }
  return statusMap[status] || status
}

const getStatusColor = (status) => {
  const colors = {
    'processing_queue': 'bg-gray-400',
    'awaiting_payment': 'bg-yellow-400',
    'pending_review': 'bg-orange-400',
    'processing': 'bg-blue-400',
    'ready_to_ship': 'bg-purple-400',
    'shipped': 'bg-blue-500',
    'delivered': 'bg-green-500',
    'cancelled': 'bg-red-600',
    'returned': 'bg-red-400',
    'refunded': 'bg-gray-400',
    'on_hold': 'bg-yellow-500',
    'fraud': 'bg-red-700',
    'paid': 'bg-green-400',
    'failed': 'bg-red-500',
    'partial': 'bg-yellow-300',
    'disputed': 'bg-orange-500',
    'pending': 'bg-yellow-400' // برای سازگاری با مقادیر قدیمی
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

const addHistoryEntry = () => {
  const newEntry = {
    status: orderStatus.value.status,
    date: new Date().toISOString(),
    user: 'مدیر سیستم',
    note: ''
  }
  orderStatus.value.history.unshift(newEntry)
}

// ذخیره تغییرات وضعیت
const saveStatusChanges = async () => {
  if (!hasChanges.value) return
  
  isLoading.value = true
  
  try {
    const updateData = {
      status: orderStatus.value.status,
      paymentStatus: orderStatus.value.paymentStatus,
      notes: orderStatus.value.internalNotes || ''
    }
    
    // console.log('ارسال درخواست به‌روزرسانی وضعیت:', updateData)
    
    const response = await $fetch(`/api/admin/orders/${props.orderId}/status`, {
      method: 'PUT',
      body: updateData
    })
    
    if (response && response.success) {
      // به‌روزرسانی وضعیت اصلی
      originalStatus.value = { ...orderStatus.value }
      
      // اضافه کردن به تاریخچه
      const historyEntry = {
        status: orderStatus.value.status,
        date: new Date().toISOString(),
        user: 'مدیر سیستم',
        note: `وضعیت به ${getStatusText(orderStatus.value.status)} تغییر یافت`
      }
      orderStatus.value.history.unshift(historyEntry)
      
      // اطلاع‌رسانی به والد
      emit('status-updated', {
        status: orderStatus.value.status,
        paymentStatus: orderStatus.value.paymentStatus,
        updatedAt: new Date().toISOString()
      })
      
      // console.log('وضعیت سفارش با موفقیت به‌روزرسانی شد')
    } else {
      throw new Error('پاسخ نامعتبر از سرور')
    }
  } catch (error) {
    console.error('خطا در به‌روزرسانی وضعیت:', error)
    // نمایش پیام خطا به کاربر
    alert('خطا در به‌روزرسانی وضعیت سفارش. لطفاً دوباره تلاش کنید.')
  } finally {
    isLoading.value = false
  }
}

// بازنشانی تغییرات
const resetChanges = () => {
  orderStatus.value = { ...originalStatus.value }
}

// نظارت بر تغییرات
watch(orderStatus, (newValue) => {
  emit('update:modelValue', newValue)
}, { deep: true })

// نظارت بر تغییرات props
watch(() => props.modelValue, (newValue) => {
  orderStatus.value = { ...newValue }
  originalStatus.value = { ...newValue }
}, { deep: true })
</script> 
