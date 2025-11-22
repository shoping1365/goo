<template>
  <ClientOnly>
    <!-- Page Header -->
    <div class="bg-white shadow-sm border-b border-gray-200 mb-4">
      <div class="w-full px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center py-3">
          <div class="flex items-center">
            <NuxtLink 
              to="/admin/transactions/orders" 
              class="text-gray-500 hover:text-gray-700 mr-4 transition-colors"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
              </svg>
            </NuxtLink>
            <div>
              <h1 class="text-lg font-bold text-gray-900">ویرایش سفارشات</h1>
              <p class="text-xs text-gray-600 mt-1">انتخاب سفارش برای ویرایش</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="w-full px-4 py-4">
      <!-- فیلترها -->
      <OrderFilters 
        v-model="filters"
        @apply-filters="applyFilters"
      />

      <!-- جدول سفارشات -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200">
        <div class="px-6 py-4 border-b border-gray-200">
          <div class="flex items-center justify-between">
            <h3 class="text-lg font-semibold text-gray-900">لیست سفارشات</h3>
            <div class="text-sm text-gray-500">
              {{ filteredOrders.length }} سفارش یافت شد
            </div>
          </div>
        </div>

        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
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
                  تاریخ
                </th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                  عملیات
                </th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="order in paginatedOrders" :key="order.id" class="hover:bg-gray-50 transition-colors">
                <!-- شماره سفارش -->
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="flex items-center">
                    <div class="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center mr-3">
                      <svg class="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                      </svg>
                    </div>
                    <div>
                      <div class="text-sm font-medium text-gray-900">{{ order.orderNumber }}</div>
                      <div class="text-sm text-gray-500">{{ order.items?.length || 0 }} محصول</div>
                    </div>
                  </div>
                </td>

                <!-- مشتری -->
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="flex items-center">
                    <div class="w-8 h-8 bg-gray-100 rounded-full flex items-center justify-center mr-3">
                      <svg class="w-4 h-4 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
                      </svg>
                    </div>
                    <div>
                      <div class="text-sm font-medium text-gray-900">{{ order.customer?.name || 'نامشخص' }}</div>
                      <div class="text-sm text-gray-500">{{ order.customer?.phone || 'نامشخص' }}</div>
                    </div>
                  </div>
                </td>

                <!-- مبلغ -->
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="text-sm font-medium text-gray-900">{{ formatPrice(order.totalAmount) }} تومان</div>
                </td>

                <!-- نحوه پرداخت -->
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="flex items-center">
                    <div class="w-6 h-6 bg-green-100 rounded-full flex items-center justify-center mr-2">
                      <svg class="w-3 h-3 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"></path>
                      </svg>
                    </div>
                    <span class="text-sm text-gray-900">{{ getPaymentMethodText(order.payment?.method) }}</span>
                  </div>
                </td>

                <!-- وضعیت -->
                <td class="px-6 py-4 whitespace-nowrap">
                  <span :class="getStatusClasses(order.status)" class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium">
                    {{ getStatusText(order.status) }}
                  </span>
                </td>

                <!-- تاریخ -->
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ formatPersianDateTime(order.createdAt) }}
                </td>

                <!-- عملیات -->
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                  <NuxtLink 
                    :to="`/admin/transactions/orders/edit/${order.id}`"
                    class="inline-flex items-center px-3 py-1.5 border border-transparent text-sm font-medium rounded-md text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 transition-colors"
                  >
                    <svg class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
                    </svg>
                    ویرایش
                  </NuxtLink>
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <!-- صفحه‌بندی -->
        <div class="px-6 py-4 border-t border-gray-200">
          <div class="flex items-center justify-between">
            <div class="text-sm text-gray-700">
              نمایش {{ startIndex + 1 }} تا {{ endIndex }} از {{ filteredOrders.length }} سفارش
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <button 
                :disabled="currentPage === 1"
                class="px-3 py-1 border border-gray-300 rounded-md text-sm disabled:opacity-50 disabled:cursor-not-allowed hover:bg-gray-50"
                @click="previousPage"
              >
                قبلی
              </button>
              <span class="px-3 py-1 text-sm text-gray-700">
                صفحه {{ currentPage }} از {{ totalPages }}
              </span>
              <button 
                :disabled="currentPage === totalPages"
                class="px-3 py-1 border border-gray-300 rounded-md text-sm disabled:opacity-50 disabled:cursor-not-allowed hover:bg-gray-50"
                @click="nextPage"
              >
                بعدی
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </ClientOnly>
</template>

// تابع تبدیل تاریخ میلادی به شمسی
const formatPersianDate = (dateString) => {
  if (!dateString) return 'نامشخص'
  
  try {
    const PersianDate = require('persian-date')
    const date = new Date(dateString)
    const persianDate = new PersianDate(date)
    return `${persianDate.year()}/${persianDate.month().toString().padStart(2, '0')}/${persianDate.date().toString().padStart(2, '0')}`
  } catch (error) {
    return 'نامشخص'
  }
}

// تابع تبدیل تاریخ و زمان میلادی به شمسی
const formatPersianDateTime = (dateString) => {
  if (!dateString) return 'نامشخص'
  
  try {
    const PersianDate = require('persian-date')
    const date = new Date(dateString)
    const persianDate = new PersianDate(date)
    const hours = date.getHours().toString().padStart(2, '0')
    const minutes = date.getMinutes().toString().padStart(2, '0')
    return `${persianDate.year()}/${persianDate.month().toString().padStart(2, '0')}/${persianDate.date().toString().padStart(2, '0')} - ${hours}:${minutes}`
  } catch (error) {
    return 'نامشخص'
  }
}

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
</script>

<script setup>
import { computed, onMounted, ref } from 'vue'

definePageMeta({ layout: 'admin-main', middleware: 'admin' });

// Import کامپوننت
import OrderFilters from './components/OrderFilters.vue'

definePageMeta({
  layout: 'admin-main'
})

// متغیرهای فیلتر
const filters = ref({
  searchTerm: '',
  statusFilter: '',
  dateFrom: '',
  dateTo: ''
})

// متغیرهای صفحه‌بندی
const currentPage = ref(1)
const itemsPerPage = ref(10)

// داده‌های نمونه
const orders = ref([
  {
    id: 1,
    orderNumber: 'ORD-2024-001',
    customer: {
      name: 'علی احمدی',
      phone: '09123456789'
    },
    totalAmount: 2500000,
    payment: {
      method: 'online'
    },
    status: 'paid',
    createdAt: '2024-01-15T10:30:00',
    items: [
      { productName: 'لپ‌تاپ اپل مک‌بوک پرو', sku: 'MBP-001', quantity: 1, price: 2500000 }
    ]
  },
  {
    id: 2,
    orderNumber: 'ORD-2024-002',
    customer: {
      name: 'فاطمه محمدی',
      phone: '09187654321'
    },
    totalAmount: 1800000,
    payment: {
      method: 'cash'
    },
    status: 'processing',
    createdAt: '2024-01-16T14:20:00',
    items: [
      { productName: 'گوشی سامسونگ گلکسی S24', sku: 'S24-001', quantity: 1, price: 1800000 }
    ]
  },
  {
    id: 3,
    orderNumber: 'ORD-2024-003',
    customer: {
      name: 'محمد رضایی',
      phone: '09111223344'
    },
    totalAmount: 3200000,
    payment: {
      method: 'bank_transfer'
    },
    status: 'shipped',
    createdAt: '2024-01-17T09:15:00',
    items: [
      { productName: 'تبلت اپل آیپد پرو', sku: 'IPAD-001', quantity: 1, price: 3200000 }
    ]
  }
])

// فیلتر کردن سفارشات
const filteredOrders = computed(() => {
  return orders.value.filter(order => {
    const matchesSearch = !filters.value.searchTerm || 
      order.orderNumber.toLowerCase().includes(filters.value.searchTerm.toLowerCase()) ||
      order.customer.name.toLowerCase().includes(filters.value.searchTerm.toLowerCase())
    
    const matchesStatus = !filters.value.statusFilter || order.status === filters.value.statusFilter
    
    const matchesDate = (!filters.value.dateFrom || new Date(order.createdAt) >= new Date(filters.value.dateFrom)) &&
                       (!filters.value.dateTo || new Date(order.createdAt) <= new Date(filters.value.dateTo))
    
    return matchesSearch && matchesStatus && matchesDate
  })
})

// محاسبه صفحه‌بندی
const totalPages = computed(() => Math.ceil(filteredOrders.value.length / itemsPerPage.value))
const startIndex = computed(() => (currentPage.value - 1) * itemsPerPage.value)
const endIndex = computed(() => Math.min(startIndex.value + itemsPerPage.value, filteredOrders.value.length))

// سفارشات صفحه فعلی
const paginatedOrders = computed(() => {
  return filteredOrders.value.slice(startIndex.value, endIndex.value)
})

// توابع صفحه‌بندی
const nextPage = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value++
  }
}

const previousPage = () => {
  if (currentPage.value > 1) {
    currentPage.value--
  }
}

// اعمال فیلترها
const applyFilters = () => {
  currentPage.value = 1
}

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

const formatPrice = (price) => {
  if (!price) return '0'
  return new Intl.NumberFormat('fa-IR').format(price)
}

const _formatDate = (date) => {
  if (!date) return 'نامشخص'
  return new Date(date).toLocaleDateString('fa-IR')
}

onMounted(() => {
  // در اینجا می‌توانید API call برای دریافت سفارشات انجام دهید
})
</script> 
