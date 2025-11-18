<template>
  <div v-if="isOpen" class="fixed inset-0 z-50 overflow-y-auto">
    <div class="flex items-center justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
      <!-- پس‌زمینه -->
      <div class="fixed inset-0 transition-opacity" @click="closeModal">
        <div class="absolute inset-0 bg-gray-500 opacity-75"></div>
      </div>

      <!-- مودال -->
      <div class="inline-block align-bottom bg-white rounded-lg text-right overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-4xl sm:w-full">
        <!-- هدر مودال -->
        <div class="bg-gradient-to-r from-blue-50 to-indigo-50 px-6 py-4 border-b border-gray-200">
          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-8 h-8 bg-blue-500 rounded-lg flex items-center justify-center ml-3">
                <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                </svg>
              </div>
              <div>
                <h3 class="text-lg font-semibold text-gray-900">جزئیات سفارش #{{ order?.orderNumber }}</h3>
                <p class="text-sm text-gray-600">مشاهده اطلاعات کامل سفارش</p>
              </div>
            </div>
            <button @click="closeModal" class="text-gray-400 hover:text-gray-600 transition-colors">
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
              </svg>
            </button>
          </div>
        </div>

        <!-- محتوای مودال -->
        <div class="px-6 py-4 max-h-96 overflow-y-auto">
          <div v-if="order" class="space-y-6">
            <!-- اطلاعات کلی سفارش -->
            <div class="bg-gray-50 rounded-lg p-6">
              <h4 class="text-md font-semibold text-gray-900 mb-3">اطلاعات کلی سفارش</h4>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  <label class="block text-sm font-medium text-gray-700">شماره سفارش:</label>
                  <p class="text-sm text-gray-900 font-medium">#{{ order.orderNumber }}</p>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700">تاریخ سفارش:</label>
                  <p class="text-sm text-gray-900">{{ formatDate(order.startedAt || order.createdAt) }}</p>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700">وضعیت:</label>
                  <span :class="getStatusClass(order.status)" class="px-2 py-1 text-xs font-medium rounded-full">
                    {{ getStatusText(order.status) }}
                  </span>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700">مبلغ کل:</label>
                  <p class="text-sm text-gray-900 font-medium">{{ formatPrice(order.totalAmount) }}</p>
                </div>
              </div>
            </div>

            <!-- اطلاعات مشتری -->
            <div class="bg-gray-50 rounded-lg p-6">
              <h4 class="text-md font-semibold text-gray-900 mb-3">اطلاعات مشتری</h4>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  <label class="block text-sm font-medium text-gray-700">نام مشتری:</label>
                  <p class="text-sm text-gray-900">{{ order.customerName }}</p>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700">شماره تماس:</label>
                  <p class="text-sm text-gray-900">{{ order.customerPhone }}</p>
                </div>
                <div v-if="order.customerEmail">
                  <label class="block text-sm font-medium text-gray-700">ایمیل:</label>
                  <p class="text-sm text-gray-900">{{ order.customerEmail }}</p>
                </div>
                <div v-if="order.customerAddress">
                  <label class="block text-sm font-medium text-gray-700">آدرس:</label>
                  <p class="text-sm text-gray-900">{{ order.customerAddress }}</p>
                </div>
              </div>
            </div>

            <!-- جزئیات پردازش (برای سفارشات در حال انجام) -->
            <div v-if="order.stage || order.progress" class="bg-gray-50 rounded-lg p-6">
              <h4 class="text-md font-semibold text-gray-900 mb-3">وضعیت پردازش</h4>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div v-if="order.stage">
                  <label class="block text-sm font-medium text-gray-700">مرحله فعلی:</label>
                  <span :class="getStageClass(order.stage)" class="px-2 py-1 text-xs font-medium rounded-full">
                    {{ getStageText(order.stage) }}
                  </span>
                </div>
                <div v-if="order.progress">
                  <label class="block text-sm font-medium text-gray-700">پیشرفت:</label>
                  <div class="flex items-center mt-1">
                    <div class="w-24 bg-gray-200 rounded-full h-2 ml-2">
                      <div
                        class="h-2 rounded-full bg-blue-500"
                        :style="{ width: `${order.progress}%` }"
                      ></div>
                    </div>
                    <span class="text-xs text-gray-500">{{ order.progress }}%</span>
                  </div>
                </div>
              </div>
            </div>

            <!-- اطلاعات مسترد/مرجوع (برای سفارشات مسترد/مرجوع) -->
            <div v-if="order.refundReason || order.returnReason" class="bg-gray-50 rounded-lg p-6">
              <h4 class="text-md font-semibold text-gray-900 mb-3">
                {{ order.refundReason ? 'اطلاعات مسترد' : 'اطلاعات مرجوع' }}
              </h4>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div v-if="order.refundReason">
                  <label class="block text-sm font-medium text-gray-700">دلیل مسترد:</label>
                  <p class="text-sm text-gray-900">{{ getRefundReasonText(order.refundReason) }}</p>
                </div>
                <div v-if="order.returnReason">
                  <label class="block text-sm font-medium text-gray-700">دلیل مرجوع:</label>
                  <p class="text-sm text-gray-900">{{ getReturnReasonText(order.returnReason) }}</p>
                </div>
                <div v-if="order.refundMethod">
                  <label class="block text-sm font-medium text-gray-700">روش مسترد:</label>
                  <p class="text-sm text-gray-900">{{ getRefundMethodText(order.refundMethod) }}</p>
                </div>
                <div v-if="order.refundedAt || order.returnedAt">
                  <label class="block text-sm font-medium text-gray-700">تاریخ {{ order.refundedAt ? 'مسترد' : 'مرجوع' }}:</label>
                  <p class="text-sm text-gray-900">{{ formatDate(order.refundedAt || order.returnedAt) }}</p>
                </div>
              </div>
            </div>

            <!-- اطلاعات ارسال (برای سفارشات ارسال شده) -->
            <div v-if="order.shippingMethod || order.trackingNumber" class="bg-gray-50 rounded-lg p-6">
              <h4 class="text-md font-semibold text-gray-900 mb-3">اطلاعات ارسال</h4>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div v-if="order.shippingMethod">
                  <label class="block text-sm font-medium text-gray-700">روش ارسال:</label>
                  <p class="text-sm text-gray-900">{{ order.shippingMethod }}</p>
                </div>
                <div v-if="order.trackingNumber">
                  <label class="block text-sm font-medium text-gray-700">شماره پیگیری:</label>
                  <p class="text-sm text-gray-900 font-mono">{{ order.trackingNumber }}</p>
                </div>
                <div v-if="order.shippedAt">
                  <label class="block text-sm font-medium text-gray-700">تاریخ ارسال:</label>
                  <p class="text-sm text-gray-900">{{ formatDate(order.shippedAt) }}</p>
                </div>
                <div v-if="order.estimatedDelivery">
                  <label class="block text-sm font-medium text-gray-700">تاریخ تحویل تخمینی:</label>
                  <p class="text-sm text-gray-900">{{ formatDate(order.estimatedDelivery) }}</p>
                </div>
              </div>
            </div>

            <!-- محصولات سفارش -->
            <div class="bg-gray-50 rounded-lg p-6">
              <h4 class="text-md font-semibold text-gray-900 mb-3">محصولات سفارش</h4>
              <div class="space-y-3">
                <div v-for="(item, index) in order.items || []" :key="index" class="flex items-center justify-between p-3 bg-white rounded border">
                  <div class="flex items-center">
                    <div class="w-12 h-12 bg-gray-200 rounded-lg flex items-center justify-center ml-3">
                      <svg class="w-6 h-6 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
                      </svg>
                    </div>
                    <div>
                      <p class="text-sm font-medium text-gray-900">{{ item.name || 'محصول ' + (index + 1) }}</p>
                      <p class="text-xs text-gray-500">تعداد: {{ item.quantity || 1 }}</p>
                    </div>
                  </div>
                  <div class="text-left">
                    <p class="text-sm font-medium text-gray-900">{{ formatPrice(item.price || 0) }}</p>
                  </div>
                </div>
              </div>
            </div>

            <!-- یادداشت‌ها -->
            <div v-if="order.notes" class="bg-gray-50 rounded-lg p-6">
              <h4 class="text-md font-semibold text-gray-900 mb-3">یادداشت‌ها</h4>
              <p class="text-sm text-gray-700">{{ order.notes }}</p>
            </div>
          </div>

          <!-- حالت بارگذاری -->
          <div v-else class="flex items-center justify-center py-8">
            <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-500"></div>
            <span class="mr-3 text-sm text-gray-500">در حال بارگذاری...</span>
          </div>
        </div>

        <!-- فوتر مودال -->
        <div class="bg-gray-50 px-6 py-3 border-t border-gray-200">
          <div class="flex justify-end space-x-2 space-x-reverse">
            <button @click="closeModal" class="px-4 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400 transition-colors">
              بستن
            </button>
            <button @click="editOrder" class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 transition-colors">
              ویرایش سفارش
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// Props
const props = defineProps({
  isOpen: {
    type: Boolean,
    default: false
  },
  order: {
    type: Object,
    default: null
  }
})

// Emits
const emit = defineEmits(['close', 'edit'])

// متدهای عملیاتی
const closeModal = () => {
  emit('close')
}

const editOrder = () => {
  emit('edit', props.order)
}

// متدهای کمکی
const formatDate = (date) => {
  if (!date) return '-'
  return new Date(date).toLocaleDateString('fa-IR')
}

const formatPrice = (price) => {
  return new Intl.NumberFormat('fa-IR').format(price || 0) + ' تومان'
}

const getStatusClass = (status) => {
  const classes = {
    pending: 'bg-yellow-100 text-yellow-800',
    processing: 'bg-blue-100 text-blue-800',
    shipped: 'bg-green-100 text-green-800',
    delivered: 'bg-green-100 text-green-800',
    returned: 'bg-orange-100 text-orange-800',
    refunded: 'bg-red-100 text-red-800',
    cancelled: 'bg-gray-100 text-gray-800',
    fraud: 'bg-red-100 text-red-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status) => {
  const texts = {
    pending: 'در انتظار',
    processing: 'در حال پردازش',
    shipped: 'ارسال شده',
    delivered: 'تحویل داده شده',
    returned: 'مرجوع شده',
    refunded: 'مسترد شده',
    cancelled: 'لغو شده',
    fraud: 'تشخیص تقلب'
  }
  return texts[status] || status
}

const getStageClass = (stage) => {
  const classes = {
    confirmation: 'bg-blue-100 text-blue-800',
    preparation: 'bg-green-100 text-green-800',
    packaging: 'bg-yellow-100 text-yellow-800',
    ready_to_ship: 'bg-purple-100 text-purple-800'
  }
  return classes[stage] || 'bg-gray-100 text-gray-800'
}

const getStageText = (stage) => {
  const texts = {
    confirmation: 'تایید سفارش',
    preparation: 'آماده‌سازی',
    packaging: 'بسته‌بندی',
    ready_to_ship: 'آماده ارسال'
  }
  return texts[stage] || stage
}

const getRefundReasonText = (reason) => {
  const texts = {
    customer: 'درخواست مشتری',
    out_of_stock: 'عدم موجودی',
    payment_issue: 'مشکل پرداخت',
    system_issue: 'مشکل سیستمی'
  }
  return texts[reason] || reason
}

const getReturnReasonText = (reason) => {
  const texts = {
    damaged: 'آسیب دیده',
    wrong_item: 'محصول اشتباه',
    not_satisfied: 'عدم رضایت',
    size_issue: 'مشکل سایز'
  }
  return texts[reason] || reason
}

const getRefundMethodText = (method) => {
  const texts = {
    bank_transfer: 'انتقال بانکی',
    wallet: 'کیف پول',
    gift_card: 'گیفت کارت',
    credit_card: 'کارت اعتباری'
  }
  return texts[method] || method
}
</script> 