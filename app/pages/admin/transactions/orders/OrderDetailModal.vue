<template>
  <div v-if="isOpen" class="fixed inset-0 z-50 overflow-y-auto">
    <!-- Backdrop -->
    <div class="fixed inset-0 bg-black bg-opacity-50 transition-opacity" @click="closeModal"></div>
    
    <!-- Modal -->
    <div class="flex min-h-full items-center justify-center px-4 py-4">
      <div class="relative w-full max-w-6xl bg-white rounded-lg shadow-xl">
        <!-- Header -->
        <div class="flex items-center justify-between px-4 py-4 border-b border-gray-200">
          <div class="flex items-center space-x-3">
            <div class="w-10 h-10 bg-blue-100 rounded-full flex items-center justify-center">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
              </svg>
            </div>
            <div>
              <h3 class="text-lg font-semibold text-gray-900">جزئیات سفارش</h3>
              <p class="text-sm text-gray-500">شماره سفارش: {{ order?.orderNumber || 'نامشخص' }}</p>
            </div>
          </div>
          <button class="text-gray-400 hover:text-gray-600 transition-colors" @click="closeModal">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>

        <!-- Content -->
        <div class="px-4 py-4 max-h-[70vh] overflow-y-auto">
          <div v-if="!order" class="text-center py-8">
            <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600 mx-auto"></div>
            <p class="mt-4 text-gray-500">در حال بارگذاری...</p>
          </div>

          <div v-else class="space-y-6">
            <!-- Order Status & Summary -->
            <div class="grid grid-cols-1 md:grid-cols-4 gap-3 py-3">
              <!-- Status Card -->
              <div :class="getStatusCardClass(order.status) + ' px-3 py-3 rounded-lg border'">
                <div class="flex items-center justify-between">
                  <div>
                    <p class="text-xs font-medium text-blue-600">وضعیت سفارش</p>
                    <p class="text-sm font-semibold text-blue-900 mt-1">
                      {{ getStatusText(order.status) }}
                    </p>
                  </div>
                  <div class="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center">
                    <svg class="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                    </svg>
                  </div>
                </div>
              </div>

              <!-- Payment Status Card -->
              <div class="bg-gradient-to-r from-orange-50 to-yellow-50 px-3 py-3 rounded-lg border border-orange-200">
                <div class="flex items-center justify-between">
                  <div>
                    <p class="text-xs font-medium text-orange-600">وضعیت پرداخت</p>
                    <p class="text-sm font-semibold text-orange-900 mt-1">
                      {{ getPaymentStatusText(order.paymentStatus || order.payment?.status || order.PaymentStatus) }}
                    </p>
                  </div>
                  <div class="w-8 h-8 bg-orange-100 rounded-full flex items-center justify-center">
                    <svg class="w-4 h-4 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"></path>
                    </svg>
                  </div>
                </div>
              </div>

              <!-- Total Amount -->
              <div class="bg-gradient-to-r from-green-50 to-emerald-50 px-3 py-3 rounded-lg border border-green-200">
                <div class="flex items-center justify-between">
                  <div>
                    <p class="text-xs font-medium text-green-600">مبلغ کل</p>
                    <p class="text-sm font-semibold text-green-900 mt-1">
                      {{ formatPrice(order.totalAmount) }} تومان
                    </p>
                  </div>
                  <div class="w-8 h-8 bg-green-100 rounded-full flex items-center justify-center">
                    <svg class="w-4 h-4 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
                    </svg>
                  </div>
                </div>
              </div>

              <!-- Order Date -->
              <div class="bg-gradient-to-r from-purple-50 to-violet-50 px-3 py-3 rounded-lg border border-purple-200">
                <div class="flex items-center justify-between">
                  <div>
                    <p class="text-xs font-medium text-purple-600">تاریخ سفارش</p>
                    <p class="text-sm font-semibold text-purple-900 mt-1">
                      <PersianDateFormatter :date-string="order.createdAt || order.created_at || order.CreatedAt" format="HH:MM - YYYY/MM/DD" />
                    </p>
                  </div>
                  <div class="w-8 h-8 bg-purple-100 rounded-full flex items-center justify-center">
                    <svg class="w-4 h-4 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                    </svg>
                  </div>
                </div>
              </div>
            </div>

            <!-- Customer Information -->
            <div class="bg-white border border-gray-200 rounded-lg px-4 py-4">
              <h4 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
                <svg class="w-5 h-5 text-blue-600 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
                </svg>
                اطلاعات مشتری
              </h4>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <!-- Personal Info -->
                <div class="space-y-3">
                  <div class="flex justify-between items-center py-2 border-b border-blue-100">
                    <span class="text-sm text-gray-500">نام و نام خانوادگی</span>
                    <span class="font-medium text-gray-900">{{ order.customer?.name || order.customerName || order.CustomerName || 'نامشخص' }}</span>
                  </div>
                  <div class="flex justify-between items-center py-2 border-b border-blue-100">
                    <span class="text-sm text-gray-500">ایمیل</span>
                    <span class="font-medium text-gray-900">{{ order.customer?.email || order.customerEmail || order.CustomerEmail || 'نامشخص' }}</span>
                  </div>
                  <div class="flex justify-between items-center py-2 border-b border-blue-100">
                    <span class="text-sm text-gray-500">کد ملی</span>
                    <span class="font-medium text-gray-900">{{ order.customer?.nationalCode || order.customerNationalCode || order.CustomerNationalCode || 'نامشخص' }}</span>
                  </div>
                </div>
                
                <!-- Contact Info -->
                <div class="space-y-3">
                  <div class="flex justify-between items-center py-2 border-b border-blue-100">
                    <span class="text-sm text-gray-500">شماره تلفن</span>
                    <span class="font-medium text-gray-900">{{ order.customer?.phone || order.customerPhone || order.CustomerPhone || 'نامشخص' }}</span>
                  </div>
                  <div class="flex justify-between items-center py-2 border-b border-blue-100">
                    <span class="text-sm text-gray-500">شماره موبایل</span>
                    <span class="font-medium text-gray-900">{{ order.customer?.mobile || order.customerMobile || order.CustomerMobile || order.customerPhone || order.CustomerPhone || 'نامشخص' }}</span>
                  </div>
                  <div class="flex justify-between items-center py-2 border-b border-blue-100">
                    <span class="text-sm text-gray-500">آدرس IP</span>
                    <span class="font-medium text-gray-900">{{ order.customerIP || 'نامشخص' }}</span>
                  </div>
                </div>
              </div>
            </div>

            <!-- Shipping Information -->
            <div class="bg-white border border-gray-200 rounded-lg px-4 py-4">
              <h4 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
                <svg class="w-5 h-5 text-green-600 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"></path>
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"></path>
                </svg>
                اطلاعات ارسال
              </h4>
              <div class="space-y-4">
                <!-- ردیف اول: استان و شهر -->
                <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                  <div class="flex justify-between items-center py-2 border-b border-blue-100">
                    <span class="text-sm text-gray-500">استان</span>
                    <span class="font-medium text-gray-900">{{ order.shipping?.province || order.shippingProvince || order.ShippingProvince || 'نامشخص' }}</span>
                  </div>
                  <div class="flex justify-between items-center py-2 border-b border-blue-100">
                    <span class="text-sm text-gray-500">شهر</span>
                    <span class="font-medium text-gray-900">{{ order.shipping?.city || order.shippingCity || order.ShippingCity || 'نامشخص' }}</span>
                  </div>
                </div>
                
                <!-- ردیف دوم: آدرس کامل -->
                <div class="py-2 border-b border-blue-100">
                  <div class="flex justify-between items-start">
                    <span class="text-sm text-gray-500">آدرس کامل</span>
                    <span class="font-medium text-gray-900 text-right flex-1 mr-3 leading-relaxed">{{ order.shipping?.address || order.shippingAddress || order.ShippingAddress || 'نامشخص' }}</span>
                  </div>
                </div>
                
                <!-- ردیف سوم: تلفن‌ها و کد پستی -->
                <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
                  <div class="flex justify-between items-center py-2 border-b border-blue-100">
                    <span class="text-sm text-gray-500">تلفن همراه</span>
                    <span class="font-medium text-gray-900">{{ order.shipping?.mobile || order.shippingMobile || order.ShippingMobile || order.customer?.mobile || order.customerMobile || order.CustomerMobile || 'نامشخص' }}</span>
                  </div>
                  <div class="flex justify-between items-center py-2 border-b border-blue-100">
                    <span class="text-sm text-gray-500">تلفن ثابت</span>
                    <span class="font-medium text-gray-900">{{ order.shipping?.phone || order.shippingPhone || order.ShippingPhone || order.customer?.phone || order.customerPhone || order.CustomerPhone || 'نامشخص' }}</span>
                  </div>
                  <div class="flex justify-between items-center py-2 border-b border-blue-100">
                    <span class="text-xs text-gray-500">کد پستی</span>
                    <span class="text-xs font-medium text-gray-600">{{ order.shipping?.postalCode || order.shippingPostalCode || order.ShippingPostalCode || 'نامشخص' }}</span>
                  </div>
                </div>
              </div>
            </div>

            <!-- Order Items -->
            <div class="bg-white border border-gray-200 rounded-lg px-4 py-4">
              <h4 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
                <svg class="w-5 h-5 text-orange-600 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
                </svg>
                محصولات سفارش
              </h4>
              <div v-if="orderItemsLoading" class="text-center py-4">
                <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-orange-600 mx-auto"></div>
                <p class="mt-2 text-gray-500">در حال بارگذاری...</p>
              </div>
              <div v-else-if="orderItems && orderItems.length > 0" class="space-y-4">
                <div v-for="(item, index) in orderItems" :key="index" class="flex items-center space-x-4 px-4 py-4 bg-gray-50 rounded-lg">
                  <div class="w-16 h-16 bg-gray-200 rounded-lg flex items-center justify-center overflow-hidden">
                    <img v-if="item.image || item.product_image" :src="item.image || item.product_image" :alt="item.name || item.product_name || 'محصول'" class="w-full h-full object-cover" />
                    <svg v-else class="w-8 h-8 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                    </svg>
                  </div>
                  <div class="flex-1">
                    <h5 class="font-medium text-gray-900">{{ item.name || item.product_name || `محصول ${item.productId || item.product_id || index + 1}` }}</h5>
                    <p class="text-sm text-gray-500">SKU: {{ item.sku || item.product_sku || 'نامشخص' }}</p>
                  </div>
                  <div class="text-right">
                    <p class="font-medium text-gray-900">{{ formatPrice(item.finalPrice || item.final_price || 0) }} تومان</p>
                    <p class="text-sm text-gray-500">تعداد: {{ item.quantity || 0 }}</p>
                  </div>
                </div>
              </div>
              <div v-else class="text-center py-4 text-gray-500">
                محصولی یافت نشد
              </div>
            </div>

            <!-- Order Timeline -->
            <div class="bg-white border border-gray-200 rounded-lg px-4 py-4">
              <h4 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
                <svg class="w-5 h-5 text-purple-600 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
                تاریخچه سفارش
              </h4>
              <div v-if="orderTimeline && (orderTimeline as any).length > 0" class="space-y-4">
                <div v-for="(event, index) in orderTimeline" :key="index" class="flex items-start space-x-4 space-x-reverse">
                  <div :class="getTimelineIconClass((event as any).type)" class="w-10 h-10 rounded-full flex items-center justify-center flex-shrink-0">
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="getTimelineIcon((event as any).type)"></path>
                    </svg>
                  </div>
                  <div class="flex-1">
                    <p class="font-medium text-gray-900">{{ (event as any).title }}</p>
                    <p class="text-sm text-gray-500">
                      <PersianDateFormatter v-if="(event as any).date" :date-string="(event as any).date" format="HH:MM - YYYY/MM/DD" />
                      <span v-else>نامشخص</span>
                    </p>
                    <p v-if="(event as any).description" class="text-sm text-gray-600 mt-1">{{ (event as any).description }}</p>
                  </div>
                </div>
              </div>
              <div v-else class="text-center py-4 text-gray-500">
                <svg class="w-12 h-12 text-gray-300 mx-auto mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
                تاریخچه‌ای برای این سفارش ثبت نشده است
              </div>
            </div>

            <!-- Payment Information -->
            <div class="bg-white border border-gray-200 rounded-lg px-4 py-4">
              <h4 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
                <svg class="w-5 h-5 text-green-600 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"></path>
                </svg>
                اطلاعات پرداخت
              </h4>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <!-- Payment Details -->
                <div class="space-y-3">
                  <div class="flex justify-between items-center py-2 border-b border-blue-100">
                    <span class="text-sm text-gray-500">روش پرداخت</span>
                    <span class="font-medium text-gray-900">{{ getPaymentMethodText(order.paymentMethod || order.payment?.method || order.PaymentMethod) || 'نامشخص' }}</span>
                  </div>
                  <div class="flex justify-between items-center py-2 border-b border-blue-100">
                    <span class="text-sm text-gray-500">وضعیت پرداخت</span>
                    <span :class="getPaymentStatusClass(order.paymentStatus || order.payment?.status || order.PaymentStatus)" class="px-2 py-1 text-xs font-medium rounded-full">
                      {{ getPaymentStatusText(order.paymentStatus || order.payment?.status || order.PaymentStatus) || 'نامشخص' }}
                    </span>
                  </div>
                  <div class="flex justify-between items-center py-2 border-b border-blue-100">
                    <span class="text-sm text-gray-500">درگاه پرداخت</span>
                    <span class="font-medium text-gray-900">{{ getPaymentGatewayText(order.paymentMethod || order.payment?.method || order.PaymentMethod) || 'نامشخص' }}</span>
                  </div>
                  <div class="flex justify-between items-center py-2 border-b border-blue-100">
                    <span class="text-sm text-gray-500">شماره تراکنش</span>
                    <span class="font-medium text-gray-900">{{ order.payment?.transactionId || order.transactionId || order.TransactionId || 'نامشخص' }}</span>
                  </div>
                </div>
                
                <!-- Device & System Info -->
                <div class="space-y-3">
                  <div class="flex justify-between items-center py-2 border-b border-blue-100">
                    <span class="text-sm text-gray-500">نوع دستگاه</span>
                    <span class="font-medium text-gray-900">{{ parseUserAgent(order.userAgent || order.user_agent || order.UserAgent).device }}</span>
                  </div>
                  <div class="flex justify-between items-center py-2 border-b border-blue-100">
                    <span class="text-sm text-gray-500">سیستم عامل</span>
                    <span class="font-medium text-gray-900">{{ parseUserAgent(order.userAgent || order.user_agent || order.UserAgent).os }}</span>
                  </div>
                  <div class="flex justify-between items-center py-2 border-b border-blue-100">
                    <span class="text-sm text-gray-500">مرورگر</span>
                    <span class="font-medium text-gray-900">{{ parseUserAgent(order.userAgent || order.user_agent || order.UserAgent).browser }}</span>
                  </div>
                  <div class="flex justify-between items-center py-2 border-b border-blue-100">
                    <span class="text-sm text-gray-500">آدرس IP</span>
                    <span class="font-medium text-gray-900">{{ order.customerIP || 'نامشخص' }}</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Footer -->
        <div class="flex items-center justify-end space-x-3 px-4 py-4 border-t border-gray-200">
          <button class="px-4 py-2 text-gray-700 bg-gray-100 hover:bg-gray-200 rounded-lg transition-colors" @click="closeModal">
            بستن
          </button>
          <button class="px-4 py-2 text-white bg-blue-600 hover:bg-blue-700 rounded-lg transition-colors flex items-center" @click="printOrder">
            <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 17h2a2 2 0 002-2v-4a2 2 0 00-2-2H5a2 2 0 00-2 2v4a2 2 0 002 2h2m2 4h6a2 2 0 002-2v-4a2 2 0 00-2-2H9a2 2 0 00-2 2v4a2 2 0 002 2zm8-12V5a2 2 0 00-2-2H9a2 2 0 00-2 2v4h10z"></path>
            </svg>
            چاپ سفارش
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
// تعریف interface ها
interface ApiResponse<T> {
  data?: T
  success?: boolean
  message?: string
}

interface OrderItem {
  id?: string
  productName?: string
  product_name?: string
  name?: string
  quantity: number
  price?: number
  finalPrice?: number
  final_price?: number
  total?: number
  productId?: string
  product_id?: string
  sku?: string
  product_sku?: string
  image?: string
  product_image?: string
}

interface OrderItemsResponse {
  items: OrderItem[]
}

interface TimelineEvent {
  type: string
  title: string
  date: string
  description?: string
}

export default {
  name: 'OrderDetailModal'
}
</script>

<script setup lang="ts">
import PersianDateFormatter from '@/components/common/PersianDateFormatter.vue'
import { computed, ref, watch } from 'vue'

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

const emit = defineEmits(['close'])

// Reactive variables for order items
const orderItems = ref<OrderItem[]>([])
const orderItemsLoading = ref(false)

const closeModal = () => {
  emit('close')
}

// Fetch order items when modal opens
const fetchOrderItems = async () => {
  // بررسی ID های مختلف
  const orderId = props.order?.ID || props.order?.id || props.order?.order_id
  
  if (!orderId) {
    return
  }
  
  orderItemsLoading.value = true
  
  try {
    const response: ApiResponse<OrderItemsResponse> = await $fetch(`/api/admin/orders/${orderId}/items`)
    
    if (response && response.success && response.data && response.data.items) {
      orderItems.value = response.data.items
    } else {
      orderItems.value = []
    }
  } catch (error) {
    console.error('❌ خطا در دریافت آیتم‌ها:', error);
    orderItems.value = []
  } finally {
    orderItemsLoading.value = false
  }
}

// Watch for order changes and fetch items
watch(() => props.order, (newOrder) => {
  const orderId = newOrder?.ID || newOrder?.id || newOrder?.order_id
  if (orderId) {
    fetchOrderItems()
  }
}, { immediate: true })

// Computed property برای تولید timeline
const orderTimeline = computed((): TimelineEvent[] => {
  if (!props.order) return []
  
  const timeline: TimelineEvent[] = []
  const createdDate = props.order.createdAt || props.order.created_at || props.order.CreatedAt
  const updatedDate = props.order.updatedAt || props.order.updated_at || props.order.UpdatedAt
  const status = props.order.status || props.order.Status
  const paymentStatus = props.order.paymentStatus || props.order.payment_status || props.order.PaymentStatus
  
  // رویداد ایجاد سفارش
  if (createdDate) {
    timeline.push({
      type: 'created',
      title: 'سفارش ثبت شد',
      date: createdDate,
      description: `سفارش با شماره ${props.order.orderNumber || props.order.order_number || props.order.OrderNumber} ثبت شد`
    })
  }
  
  // رویداد پرداخت
  if (paymentStatus === 'paid') {
    timeline.push({
      type: 'payment',
      title: 'پرداخت انجام شد',
      date: updatedDate || createdDate,
      description: 'پرداخت سفارش با موفقیت انجام شد'
    })
  } else if (paymentStatus === 'failed') {
    timeline.push({
      type: 'payment_failed',
      title: 'پرداخت ناموفق',
      date: updatedDate || createdDate,
      description: 'پرداخت سفارش با خطا مواجه شد'
    })
  }
  
  // رویداد وضعیت سفارش
  if (status === 'processing' || status === 'processing_queue') {
    timeline.push({
      type: 'processing',
      title: 'در حال پردازش',
      date: updatedDate || createdDate,
      description: 'سفارش در حال پردازش است'
    })
  } else if (status === 'shipped') {
    timeline.push({
      type: 'shipped',
      title: 'ارسال شد',
      date: updatedDate || createdDate,
      description: props.order.trackingCode || props.order.tracking_code ? 
        `سفارش با کد رهگیری ${props.order.trackingCode || props.order.tracking_code} ارسال شد` : 
        'سفارش ارسال شد'
    })
  } else if (status === 'delivered') {
    timeline.push({
      type: 'delivered',
      title: 'تحویل داده شد',
      date: updatedDate || createdDate,
      description: 'سفارش با موفقیت تحویل مشتری شد'
    })
  } else if (status === 'cancelled') {
    timeline.push({
      type: 'cancelled',
      title: 'لغو شد',
      date: updatedDate || createdDate,
      description: 'سفارش لغو شد'
    })
  } else if (status === 'refunded') {
    timeline.push({
      type: 'refunded',
      title: 'بازگشت وجه',
      date: updatedDate || createdDate,
      description: 'مبلغ سفارش به حساب مشتری بازگشت داده شد'
    })
  }
  
  // مرتب‌سازی بر اساس تاریخ
  return timeline.sort((a, b) => new Date(a.date).getTime() - new Date(b.date).getTime())
})

const parseUserAgent = (userAgent: string | undefined) => {
  if (!userAgent) return { device: 'نامشخص', os: 'نامشخص', browser: 'نامشخص' }
  
  const ua = userAgent.toLowerCase()
  
  // تشخیص دستگاه
  let device = 'نامشخص'
  if (ua.includes('mobile') || ua.includes('android') || ua.includes('iphone')) {
    device = 'موبایل'
  } else if (ua.includes('tablet') || ua.includes('ipad')) {
    device = 'تبلت'
  } else if (ua.includes('desktop') || ua.includes('windows') || ua.includes('macintosh') || ua.includes('linux')) {
    device = 'دسکتاپ'
  }
  
  // تشخیص سیستم عامل
  let os = 'نامشخص'
  if (ua.includes('windows nt 10')) os = 'Windows 10/11'
  else if (ua.includes('windows nt 6.3')) os = 'Windows 8.1'
  else if (ua.includes('windows nt 6.2')) os = 'Windows 8'
  else if (ua.includes('windows nt 6.1')) os = 'Windows 7'
  else if (ua.includes('windows')) os = 'Windows'
  else if (ua.includes('macintosh') || ua.includes('mac os')) os = 'macOS'
  else if (ua.includes('linux')) os = 'Linux'
  else if (ua.includes('android')) os = 'Android'
  else if (ua.includes('iphone') || ua.includes('ipad')) os = 'iOS'
  
  // تشخیص مرورگر
  let browser = 'نامشخص'
  if (ua.includes('chrome') && !ua.includes('edge')) browser = 'Chrome'
  else if (ua.includes('firefox')) browser = 'Firefox'
  else if (ua.includes('safari') && !ua.includes('chrome')) browser = 'Safari'
  else if (ua.includes('edge')) browser = 'Edge'
  else if (ua.includes('opera')) browser = 'Opera'
  else if (ua.includes('msie') || ua.includes('trident')) browser = 'Internet Explorer'
  
  return { device, os, browser }
}

const getPaymentMethodText = (method: string | undefined) => {
  const methodMap = {
    'online': 'پرداخت آنلاین',
    'cash': 'پرداخت نقدی',
    'cod': 'پرداخت در محل',
    'bank_transfer': 'انتقال بانکی',
    'wallet': 'کیف پول',
    'gift_card': 'کارت هدیه',
    'melli': 'درگاه ملی',
    'parsian': 'درگاه پارسیان',
    'saman': 'درگاه سامان',
    'zarinpal': 'زرین‌پال',
    'paypal': 'پی‌پال',
    'stripe': 'استرایپ',
    'mellat': 'درگاه ملت'
  }
  return methodMap[method] || method || 'نامشخص'
}

const getPaymentGatewayText = (method: string | undefined) => {
  const gatewayMap = {
    'melli': 'درگاه ملی',
    'parsian': 'درگاه پارسیان',
    'saman': 'درگاه سامان',
    'zarinpal': 'زرین‌پال',
    'paypal': 'پی‌پال',
    'stripe': 'استرایپ',
    'mellat': 'درگاه ملت',
    'online': 'پرداخت آنلاین',
    'cash': 'پرداخت نقدی',
    'cod': 'پرداخت در محل',
    'bank_transfer': 'انتقال بانکی',
    'wallet': 'کیف پول',
    'gift_card': 'کارت هدیه'
  }
  return gatewayMap[method] || method || 'نامشخص'
}

// کامپوزابل تاریخ شمسی در صورت نیاز استفاده می‌شود

const getPaymentStatusText = (status: string | undefined) => {
  const statusMap = {
    'pending': 'در انتظار پرداخت',
    'paid': 'پرداخت شده',
    'failed': 'ناموفق',
    'cancelled': 'لغو شده',
    'refunded': 'بازگشت وجه',
    'awaiting_payment': 'در انتظار پرداخت',
    'processing': 'در حال پردازش',
    'completed': 'تکمیل شده'
  }
  return statusMap[status] || status || 'نامشخص'
}

const getPaymentStatusClass = (status: string | undefined) => {
  const classMap = {
    'pending': 'bg-yellow-100 text-yellow-800',
    'paid': 'bg-green-100 text-green-800',
    'failed': 'bg-red-100 text-red-800',
    'cancelled': 'bg-red-100 text-red-800',
    'refunded': 'bg-gray-100 text-gray-800',
    'awaiting_payment': 'bg-yellow-100 text-yellow-800',
    'processing': 'bg-blue-100 text-blue-800',
    'completed': 'bg-green-100 text-green-800'
  }
  return classMap[status] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status: string | undefined) => {
  const statusMap = {
    'pending': 'در انتظار پرداخت',
    'paid': 'پرداخت شده',
    'processing': 'در حال پردازش',
    'processing_queue': 'در صف پردازش',
    'shipped': 'ارسال شده',
    'delivered': 'تحویل داده شده',
    'cancelled': 'لغو شده',
    'refunded': 'بازگشت وجه'
  }
  return statusMap[status] || status
}

// وضعیت رنگ کارت وضعیت سفارش را بر اساس وضعیت برگردان
const getStatusCardClass = (status: string | undefined) => {
  switch (status) {
    case 'processing_queue':
      return 'bg-gradient-to-r from-yellow-50 to-yellow-100 border-yellow-200';
    case 'pending':
      return 'bg-gradient-to-r from-blue-50 to-indigo-50 border-blue-200';
    case 'processing':
      return 'bg-gradient-to-r from-blue-50 to-indigo-50 border-blue-200';
    case 'shipped':
      return 'bg-gradient-to-r from-green-50 to-emerald-50 border-green-200';
    case 'delivered':
      return 'bg-gradient-to-r from-purple-50 to-violet-50 border-purple-200';
    case 'cancelled':
      return 'bg-gradient-to-r from-red-50 to-red-100 border-red-200';
    case 'refunded':
      return 'bg-gradient-to-r from-gray-50 to-gray-100 border-gray-200';
    default:
      return 'bg-gradient-to-r from-blue-50 to-indigo-50 border-blue-200';
  }
}

const formatPrice = (price: number | string | undefined) => {
  if (!price) return '0'
  const numPrice = typeof price === 'string' ? parseFloat(price) : price
  return new Intl.NumberFormat('fa-IR').format(numPrice)
}

const _formatDate = (date: string | Date | undefined) => {
  if (!date) return 'نامشخص'
  
  try {
    // اگر تاریخ به صورت string است، آن را به Date تبدیل کن
    const dateObj = typeof date === 'string' ? new Date(date) : date
    
    // بررسی معتبر بودن تاریخ
    if (isNaN(dateObj.getTime())) {
      return 'نامشخص'
    }
    
    return dateObj.toLocaleDateString('fa-IR', {
      year: 'numeric',
      month: 'long',
      day: 'numeric'
    })
  } catch (error) {
    return 'نامشخص'
  }
}

const printOrder = () => {
  window.print()
}

// دریافت آیکون timeline بر اساس نوع رویداد
const getTimelineIcon = (type: string) => {
  const icons = {
    'created': 'M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z',
    'payment': 'M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z',
    'payment_failed': 'M6 18L18 6M6 6l12 12',
    'processing': 'M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15',
    'shipped': 'M5 8h14M5 8a2 2 0 110-4h14a2 2 0 110 4M5 8v10a2 2 0 002 2h10a2 2 0 002-2V8m-9 4h4',
    'delivered': 'M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z',
    'cancelled': 'M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z',
    'refunded': 'M3 10h10a8 8 0 018 8v2M3 10l6 6m-6-6l6-6'
  }
  return icons[type] || 'M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z'
}

// دریافت کلاس آیکون timeline
const getTimelineIconClass = (type: string) => {
  const classes = {
    'created': 'bg-blue-100 text-blue-600',
    'payment': 'bg-green-100 text-green-600',
    'payment_failed': 'bg-red-100 text-red-600',
    'processing': 'bg-yellow-100 text-yellow-600',
    'shipped': 'bg-indigo-100 text-indigo-600',
    'delivered': 'bg-green-100 text-green-600',
    'cancelled': 'bg-red-100 text-red-600',
    'refunded': 'bg-gray-100 text-gray-600'
  }
  return classes[type] || 'bg-blue-100 text-blue-600'
}
</script>

<style scoped>
@media print {
  .fixed {
    position: static ;
  }
  
  .bg-black {
    background: none ;
  }
  
  .shadow-xl {
    box-shadow: none ;
  }
  
  .border {
    border: 1px solid #000 ;
  }
}
</style> 
