<template>
  <div v-if="hasAccess" class="bg-white rounded-lg shadow-sm border border-gray-200 relative">
    <!-- Loading Overlay -->
    <div v-if="loading" class="absolute inset-0 bg-white bg-opacity-75 flex items-center justify-center z-10 rounded-lg">
      <div class="flex items-center space-x-2">
        <svg class="animate-spin h-6 w-6 text-blue-600" fill="none" viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
        </svg>
        <span class="text-sm text-gray-600">در حال بارگذاری...</span>
      </div>
    </div>
    <!-- Table Header -->
    <div class="px-6 py-4 border-b border-gray-200">
      <div class="flex items-center justify-between">
        <h3 class="text-lg font-semibold text-gray-900">لیست سفارشات</h3>
        <div class="flex items-center space-x-3 space-x-reverse">
          <!-- دکمه رفرش -->
          <button 
            :disabled="loading" 
            class="p-2 text-gray-500 hover:text-gray-700 hover:bg-gray-100 rounded-lg transition-colors disabled:opacity-50 disabled:cursor-not-allowed relative group"
            @click="$emit('refresh')"
          >
            <svg 
              :class="['w-5 h-5', { 'animate-spin': loading }]" 
              fill="none" 
              stroke="currentColor" 
              viewBox="0 0 24 24"
            >
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
            </svg>
            <!-- Tooltip -->
            <div class="absolute bottom-full left-1/2 transform -translate-x-1/2 mb-2 px-2 py-1 text-xs text-white bg-gray-900 rounded opacity-0 group-hover:opacity-100 transition-opacity pointer-events-none whitespace-nowrap">
              بارگذاری مجدد
            </div>
          </button>
        </div>
      </div>
      <div v-if="selectedOrderIds.length > 2" class="mt-4 flex justify-end">
        <button class="px-4 py-2 bg-red-600 text-white rounded hover:bg-red-700 transition-colors" @click="handleBulkDelete">حذف سفارشات انتخاب شده</button>
      </div>
    </div>

    <!-- Table -->
    <div class="overflow-x-auto">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-2 py-3 text-center">
              <input type="checkbox" :checked="isAllSelected" @change="toggleSelectAll" />
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              شماره سفارش
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              مشتری
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              مبلغ
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              نحوه پرداخت
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              وضعیت
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              صحت سفارش
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              تاریخ
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              عملیات
            </th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-for="order in orders" :key="order.id" class="hover:bg-gray-50 transition-colors">
            <td class="px-2 py-4 text-center">
              <input v-model="selectedOrderIds" type="checkbox" :value="order.id" />
            </td>
            <!-- Order Number -->
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="flex items-center">
                <div class="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center mr-3">
                  <svg class="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                  </svg>
                </div>
                <div>
                  <div class="text-sm font-medium text-gray-900">{{ order.orderNumber }}</div>
                  <div class="text-sm text-gray-500">{{ order.itemsCount || 0 }} محصول</div>
                </div>
              </div>
            </td>

            <!-- Customer -->
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="flex items-center">
                <div class="w-8 h-8 bg-gray-100 rounded-full flex items-center justify-center mr-3">
                  <svg class="w-4 h-4 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
                  </svg>
                </div>
                <div>
                  <div class="text-sm font-medium text-gray-900">{{ order.customerName || 'نامشخص' }}</div>
                  <div class="text-sm text-gray-500">{{ order.customerPhone || 'نامشخص' }}</div>
                </div>
              </div>
            </td>

            <!-- Amount -->
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="text-sm font-medium text-gray-900">{{ formatPrice(order.totalAmount) }} تومان</div>
            </td>

            <!-- Payment Method -->
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="flex items-center">
                <div class="w-6 h-6 bg-green-100 rounded-full flex items-center justify-center mr-2">
                  <svg class="w-3 h-3 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"></path>
                  </svg>
                </div>
                <span class="text-sm text-gray-900">{{ getPaymentMethodText(order.paymentMethod) }}</span>
              </div>
            </td>

            <!-- Status -->
            <td class="px-6 py-4 whitespace-nowrap">
              <span :class="getStatusClasses(order.status)" class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium">
                {{ getStatusText(order.status) }}
              </span>
            </td>

            <!-- Order Integrity -->
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="flex items-center">
                <span :class="getIntegrityClasses(order.orderIntegrity)" class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium">
                  <svg v-if="order.orderIntegrity === 'suspicious'" class="w-3 h-3 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                  </svg>
                  <svg v-else-if="order.orderIntegrity === 'fraud_detected'" class="w-3 h-3 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
                  </svg>
                  <svg v-else-if="order.orderIntegrity === 'verified'" class="w-3 h-3 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                  </svg>
                  {{ getIntegrityText(order.orderIntegrity) }}
                </span>
              </div>
            </td>

             <!-- Date -->
             <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
               <PersianDateFormatter :date-string="order.createdAt" format="HH:MM - YYYY/MM/DD" />
             </td>

            <!-- Actions -->
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
              <div class="flex items-center space-x-2">
                <!-- View Details -->
                <button class="text-blue-600 hover:text-blue-900 p-1 rounded transition-colors" @click="$emit('view-details', order)">
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                  </svg>
                </button>

                <!-- Edit Order -->
                <NuxtLink :to="`/admin/transactions/orders/edit/${order.id}`" class="text-green-600 hover:text-green-900 p-1 rounded transition-colors">
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
                  </svg>
                </NuxtLink>

                <!-- Quick Actions -->
                <button class="text-gray-600 hover:text-gray-900 p-1 rounded transition-colors" @click="$emit('quick-actions', order)">
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z"></path>
                  </svg>
                </button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Empty State -->
    <div v-if="orders.length === 0" class="text-center py-12">
      <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
      </svg>
      <h3 class="mt-2 text-sm font-medium text-gray-900">هیچ سفارشی یافت نشد</h3>
      <p class="mt-1 text-sm text-gray-500">سفارش جدیدی ایجاد کنید یا فیلترها را تغییر دهید.</p>
    </div>
  </div>
</template>

<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
</script>

<script setup lang="ts">
import PersianDateFormatter from '@/components/common/PersianDateFormatter.vue';
import { computed, onMounted, ref, watch } from 'vue';
import { useAuth } from '~/composables/useAuth';

// احراز هویت
const { user, isAuthenticated } = useAuth()

// بررسی دسترسی admin
const hasAccess = computed(() => {
  if (!isAuthenticated.value) {
    return false
  }

  const userRole = user.value?.role?.toLowerCase() || ''
  const adminRoles = ['admin', 'developer']
  return adminRoles.includes(userRole)
})

// بررسی احراز هویت و دسترسی admin - نمایش 404 در صورت عدم دسترسی
const checkAuth = async () => {
  if (!hasAccess.value) {
    await navigateTo('/404', { external: false })
  }
}

// بررسی احراز هویت در هنگام mount
onMounted(async () => {
  await checkAuth()
})

// بررسی احراز هویت هنگام تغییر وضعیت احراز هویت
watch([isAuthenticated, hasAccess], async () => {
  if (!hasAccess.value) {
    await checkAuth()
  }
})

interface Order {
  id: string | number
  orderNumber: string
  itemsCount?: number
  customerName: string
  customerPhone?: string
  totalAmount: number
  paymentMethod?: string
  status: string
  orderIntegrity?: string
  createdAt: string
  [key: string]: unknown
}

const props = defineProps<{
  orders: Order[]
  loading?: boolean
}>()

// استفاده از کامپوزابل تاریخ شمسی
// const { } = usePersianDate()

const emit = defineEmits(['view-details', 'edit-order', 'quick-actions', 'refresh', 'export', 'bulk-delete'])

const selectedOrderIds = ref([])
const isAllSelected = computed(() => props.orders.length > 0 && selectedOrderIds.value.length === props.orders.length)

function toggleSelectAll() {
  if (isAllSelected.value) {
    selectedOrderIds.value = []
  } else {
    selectedOrderIds.value = props.orders.map(o => o.id)
  }
}

function handleBulkDelete() {
  emit('bulk-delete', selectedOrderIds.value)
  selectedOrderIds.value = []
}

const getStatusText = (status) => {
  const statusMap = {
    'pending': 'در انتظار پرداخت',
    'paid': 'پرداخت شده',
    'processing_queue': 'در صف پردازش',
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
    'processing_queue': 'bg-yellow-100 text-yellow-800',
    'processing': 'bg-blue-100 text-blue-800',
    'shipped': 'bg-purple-100 text-purple-800',
    'delivered': 'bg-green-100 text-green-800',
    'cancelled': 'bg-red-100 text-red-800',
    'refunded': 'bg-gray-100 text-gray-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getPaymentMethodText = (method) => {
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

const getIntegrityText = (integrity) => {
  const integrityMap = {
    'verified': 'تایید شده',
    'suspicious': 'مشکوک',
    'fraud_detected': 'تشخیص تقلب',
    'pending_review': 'در انتظار بررسی'
  }
  return integrityMap[integrity] || 'نامشخص'
}

const getIntegrityClasses = (integrity) => {
  const classes = {
    'verified': 'bg-green-100 text-green-800',
    'suspicious': 'bg-orange-100 text-orange-800',
    'fraud_detected': 'bg-red-100 text-red-800',
    'pending_review': 'bg-yellow-100 text-yellow-800'
  }
  return classes[integrity] || 'bg-gray-100 text-gray-800'
}

const formatPrice = (price) => {
  if (!price) return '0'
  return new Intl.NumberFormat('fa-IR').format(price)
}
</script>
