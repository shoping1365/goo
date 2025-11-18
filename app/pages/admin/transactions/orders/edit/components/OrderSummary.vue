<template>
  <div>
    <!-- اطلاعات کلی سفارش -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-blue-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
        </svg>
        خلاصه سفارش
      </h3>
      
      <div class="grid grid-cols-1 md:grid-cols-3 gapx-4 py-4">
        <!-- اطلاعات سفارش -->
        <div class="space-y-3">
          <h4 class="text-md font-medium text-gray-700">اطلاعات سفارش</h4>
          <div class="space-y-2 text-sm">
            <div class="flex justify-between">
              <span class="text-gray-600">شماره سفارش:</span>
              <span class="font-medium">{{ orderData.orderNumber }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-600">تاریخ سفارش:</span>
              <span class="font-medium">{{ formatDate(orderData.createdAt) }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-600">وضعیت:</span>
              <span :class="getStatusClasses(orderData.status.status)" class="px-2 py-1 rounded-full text-xs font-medium">
                {{ getStatusText(orderData.status.status) }}
              </span>
            </div>
          </div>
        </div>

        <!-- اطلاعات مشتری -->
        <div class="space-y-3">
          <h4 class="text-md font-medium text-gray-700">اطلاعات مشتری</h4>
          <div class="space-y-2 text-sm">
            <div class="flex justify-between">
              <span class="text-gray-600">نام:</span>
              <span class="font-medium">{{ orderData.customer.name }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-600">تلفن:</span>
              <span class="font-medium">{{ orderData.customer.phone }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-600">ایمیل:</span>
              <span class="font-medium">{{ orderData.customer.email }}</span>
            </div>
          </div>
        </div>

        <!-- اطلاعات پرداخت -->
        <div class="space-y-3">
          <h4 class="text-md font-medium text-gray-700">اطلاعات پرداخت</h4>
          <div class="space-y-2 text-sm">
            <div class="flex justify-between">
              <span class="text-gray-600">روش پرداخت:</span>
              <span class="font-medium">{{ getPaymentMethodText(orderData.status.paymentMethod) }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-600">وضعیت پرداخت:</span>
              <span :class="getPaymentStatusClasses(orderData.status.paymentStatus)" class="px-2 py-1 rounded-full text-xs font-medium">
                {{ getPaymentStatusText(orderData.status.paymentStatus) }}
              </span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-600">روش ارسال:</span>
              <span class="font-medium">{{ getShippingMethodText(orderData.status.shippingMethod) }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="mb-6"></div>

    <!-- خلاصه مالی -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-green-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
        </svg>
        خلاصه مالی
      </h3>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
        <!-- جدول محصولات -->
        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">محصولات سفارش</h4>
          <div class="space-y-2">
            <div v-for="item in orderData.items" :key="item.sku" class="flex justify-between items-center p-2 bg-gray-50 rounded">
              <div class="flex-1">
                <div class="text-sm font-medium">{{ item.name }}</div>
                <div class="text-xs text-gray-500">{{ item.sku }} - {{ item.quantity }} عدد</div>
              </div>
              <div class="text-sm font-medium">{{ formatPrice(item.totalPrice) }} تومان</div>
            </div>
          </div>
        </div>

        <!-- محاسبات مالی -->
        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">محاسبات</h4>
          <div class="space-y-2">
            <div class="flex justify-between text-sm">
              <span class="text-gray-600">جمع کل محصولات:</span>
              <span class="font-medium">{{ formatPrice(subtotal) }} تومان</span>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-600">تخفیف کل:</span>
              <span class="font-medium text-green-600">-{{ formatPrice(totalDiscount) }} تومان</span>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-600">هزینه ارسال:</span>
              <span class="font-medium">{{ formatPrice(orderData.shippingCost) }} تومان</span>
            </div>
            <div class="border-t pt-2 flex justify-between text-lg font-bold">
              <span>مبلغ نهایی:</span>
              <span class="text-blue-600">{{ formatPrice(totalAmount) }} تومان</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="mb-6"></div>

    <!-- آدرس ارسال -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-purple-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"></path>
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"></path>
        </svg>
        آدرس ارسال
      </h3>
      
      <div class="bg-gray-50 px-4 py-4 rounded-lg">
        <div class="text-sm">
          <div class="font-medium mb-2">{{ orderData.customer.name }}</div>
          <div class="text-gray-600">{{ orderData.customer.address.fullAddress }}</div>
          <div class="text-gray-600">{{ orderData.customer.address.city }} - {{ getProvinceText(orderData.customer.address.province) }}</div>
          <div class="text-gray-600">کد پستی: {{ orderData.customer.address.postalCode }}</div>
        </div>
      </div>
    </div>

    <div class="mb-6"></div>

    <!-- تاریخچه وضعیت -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-orange-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
        </svg>
        تاریخچه تغییرات
      </h3>
      
      <div class="space-y-3">
        <div 
          v-for="(history, index) in orderData.status.history" 
          :key="index"
          class="flex items-center p-3 bg-gray-50 rounded-lg"
        >
          <div class="w-3 h-3 rounded-full mr-3" :class="getStatusColor(history.status)"></div>
          <div class="flex-1">
            <div class="text-sm font-medium text-gray-900">{{ getStatusText(history.status) }}</div>
            <div class="text-xs text-gray-500">{{ formatDate(history.date) }}</div>
          </div>
          <div class="text-xs text-gray-500">{{ history.user }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

// تعریف props
const props = defineProps({
  orderData: {
    type: Object,
    required: true
  }
})

// محاسبات
const subtotal = computed(() => {
  return props.orderData.items.reduce((sum, item) => sum + (item.unitPrice * item.quantity), 0)
})

const totalDiscount = computed(() => {
  return props.orderData.items.reduce((sum, item) => sum + (item.discount || 0), 0)
})

const totalAmount = computed(() => {
  return subtotal.value - totalDiscount.value + props.orderData.shippingCost
})

// توابع کمکی
const getStatusText = (status) => {
  const statusMap = {
    'pending': 'در انتظار پرداخت',
    'paid': 'پرداخت شده',
    'processing': 'در حال پردازش',
    'shipped': 'ارسال شده',
    'delivered': 'تحویل داده شده',
    'cancelled': 'لغو شده',
    'refunded': 'بازگشت وجه'
  }
  return statusMap[status] || status
}

const getStatusClasses = (status) => {
  const classes = {
    'pending': 'bg-yellow-100 text-yellow-800',
    'paid': 'bg-green-100 text-green-800',
    'processing': 'bg-blue-100 text-blue-800',
    'shipped': 'bg-purple-100 text-purple-800',
    'delivered': 'bg-green-100 text-green-800',
    'cancelled': 'bg-red-100 text-red-800',
    'refunded': 'bg-gray-100 text-gray-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getStatusColor = (status) => {
  const colors = {
    'pending': 'bg-yellow-400',
    'paid': 'bg-green-400',
    'processing': 'bg-blue-400',
    'shipped': 'bg-purple-400',
    'delivered': 'bg-green-500',
    'cancelled': 'bg-red-400',
    'refunded': 'bg-gray-400'
  }
  return colors[status] || 'bg-gray-400'
}

const getPaymentMethodText = (method) => {
  const methodMap = {
    'online': 'پرداخت آنلاین',
    'cash': 'پرداخت نقدی',
    'bank_transfer': 'انتقال بانکی',
    'wallet': 'کیف پول',
    'gift_card': 'کارت هدیه'
  }
  return methodMap[method] || 'نامشخص'
}

const getPaymentStatusText = (status) => {
  const statusMap = {
    'pending': 'در انتظار پرداخت',
    'paid': 'پرداخت شده',
    'failed': 'ناموفق',
    'refunded': 'بازگشت وجه'
  }
  return statusMap[status] || status
}

const getPaymentStatusClasses = (status) => {
  const classes = {
    'pending': 'bg-yellow-100 text-yellow-800',
    'paid': 'bg-green-100 text-green-800',
    'failed': 'bg-red-100 text-red-800',
    'refunded': 'bg-gray-100 text-gray-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getShippingMethodText = (method) => {
  const methodMap = {
    'post': 'پست',
    'express': 'پست پیشتاز',
    'courier': 'پیک موتوری',
    'pickup': 'دریافت حضوری'
  }
  return methodMap[method] || 'نامشخص'
}

const getProvinceText = (province) => {
  const provinceMap = {
    'tehran': 'تهران',
    'isfahan': 'اصفهان',
    'shiraz': 'شیراز',
    'mashhad': 'مشهد',
    'tabriz': 'تبریز'
  }
  return provinceMap[province] || province
}

const formatPrice = (price) => {
  if (!price) return '0'
  return new Intl.NumberFormat('fa-IR').format(price)
}

const formatDate = (date) => {
  if (!date) return 'نامشخص'
  return new Date(date).toLocaleDateString('fa-IR')
}
</script> 
