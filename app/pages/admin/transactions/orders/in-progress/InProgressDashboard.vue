<template>
  <div>
    <!-- آمار کلی سفارشات در صف پردازش -->
    <div class="rounded-lg shadow-md border border-gray-200 overflow-hidden mb-6">
      <div class="bg-gradient-to-r from-blue-50 to-indigo-50 px-4 py-3 border-b border-gray-200">
        <div class="flex items-center">
          <div class="w-6 h-6 bg-blue-500 rounded-md flex items-center justify-center ml-2">
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
          <h3 class="text-sm font-semibold text-gray-900">آمار کلی سفارشات در صف پردازش</h3>
        </div>
      </div>

      <div class="px-4 py-4">
        <div class="grid grid-cols-1 md:grid-cols-4 gapx-4 py-4">
          <!-- کل سفارشات در صف -->
          <div class="bg-gradient-to-br from-blue-50 to-blue-100 border border-blue-200 rounded-lg px-4 py-4 hover:shadow-md transition-all duration-200">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-blue-700 mb-1">کل سفارشات در صف</p>
                <p class="text-2xl font-bold text-blue-900">{{ stats.totalInQueue.toLocaleString('fa-IR') }}</p>
                <p class="text-xs text-blue-600 mt-1">در حال پردازش</p>
              </div>
              <div class="w-12 h-12 bg-blue-500 rounded-lg flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  
                </svg>
              </div>
            </div>
          </div>

          <!-- در انتظار پرداخت -->
          <div class="bg-gradient-to-br from-yellow-50 to-yellow-100 border border-yellow-200 rounded-lg px-4 py-4 hover:shadow-md transition-all duration-200">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-yellow-700 mb-1">در انتظار پرداخت</p>
                <p class="text-2xl font-bold text-yellow-900">{{ stats.pendingPayment.toLocaleString('fa-IR') }}</p>
                <p class="text-xs text-yellow-600 mt-1">نیاز به تایید</p>
              </div>
              <div class="w-12 h-12 bg-yellow-500 rounded-lg flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"></path>
                </svg>
              </div>
            </div>
          </div>

          <!-- در انتظار بررسی -->
          <div class="bg-gradient-to-br from-orange-50 to-orange-100 border border-orange-200 rounded-lg px-4 py-4 hover:shadow-md transition-all duration-200">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-orange-700 mb-1">در انتظار بررسی</p>
                <p class="text-2xl font-bold text-orange-900">{{ stats.pendingReview.toLocaleString('fa-IR') }}</p>
                <p class="text-xs text-orange-600 mt-1">نیاز به بررسی</p>
              </div>
              <div class="w-12 h-12 bg-orange-500 rounded-lg flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                </svg>
              </div>
            </div>
          </div>

          <!-- متوسط زمان پردازش -->
          <div class="bg-gradient-to-br from-green-50 to-green-100 border border-green-200 rounded-lg px-4 py-4 hover:shadow-md transition-all duration-200">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-green-700 mb-1">متوسط زمان پردازش</p>
                <p class="text-2xl font-bold text-green-900">{{ stats.avgProcessingTime }} ساعت</p>
                <p class="text-xs text-green-600 mt-1">از ثبت تا تکمیل</p>
              </div>
              <div class="w-12 h-12 bg-green-500 rounded-lg flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
                </svg>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- آمار تفصیلی -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gapx-4 py-4 mb-6">
      <!-- توزیع وضعیت‌ها -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
        <div class="bg-gradient-to-r from-purple-50 to-indigo-50 px-4 py-3 border-b border-gray-200">
          <div class="flex items-center">
            <div class="w-6 h-6 bg-purple-500 rounded-md flex items-center justify-center ml-2">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
              </svg>
            </div>
            <h3 class="text-sm font-semibold text-gray-900">توزیع وضعیت‌ها</h3>
          </div>
        </div>

        <div class="px-4 py-4">
          <div class="space-y-4">
            <div
              v-for="status in orderStatuses"
              :key="status.id"
              class="flex items-center justify-between"
            >
              <div class="flex items-center">
                <div class="w-3 h-3 rounded-full ml-3" :style="{ backgroundColor: status.color }"></div>
                <span class="text-sm text-gray-700">{{ status.name }}</span>
              </div>
              <div class="flex items-center space-x-3 space-x-reverse">
                <div class="w-32 bg-gray-200 rounded-full h-2">
                  <div
                    class="h-2 rounded-full transition-all duration-300"
                    :style="{ width: `${status.percentage}%`, backgroundColor: status.color }"
                  ></div>
                </div>
                <span class="text-sm font-medium text-gray-900 w-12 text-left">{{ status.percentage }}%</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- آمار زمانی پردازش -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
        <div class="bg-gradient-to-r from-blue-50 to-indigo-50 px-4 py-3 border-b border-gray-200">
          <div class="flex items-center">
            <div class="w-6 h-6 bg-blue-500 rounded-md flex items-center justify-center ml-2">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
            <h3 class="text-sm font-semibold text-gray-900">آمار زمانی پردازش</h3>
          </div>
        </div>

        <div class="px-4 py-4">
          <div class="grid grid-cols-2 gapx-4 py-4">
            <div class="text-center">
              <div class="text-2xl font-bold text-green-600">{{ timeStats.fastestProcessing }}</div>
              <div class="text-xs text-gray-600 mt-1">سریع‌ترین پردازش</div>
            </div>
            <div class="text-center">
              <div class="text-2xl font-bold text-blue-600">{{ timeStats.averageProcessingTime }}</div>
              <div class="text-xs text-gray-600 mt-1">متوسط زمان پردازش</div>
            </div>
            <div class="text-center">
              <div class="text-2xl font-bold text-orange-600">{{ timeStats.longestProcessing }}</div>
              <div class="text-xs text-gray-600 mt-1">دیرترین پردازش</div>
            </div>
            <div class="text-center">
              <div class="text-2xl font-bold text-purple-600">{{ timeStats.processedPerDay }}</div>
              <div class="text-xs text-gray-600 mt-1">پردازش در روز</div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- نمودارهای گزارشات -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gapx-4 py-4 mb-6">
      <!-- نمودار روند پردازش -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">روند پردازش در 30 روز گذشته</h3>
          <div class="flex items-center space-x-2 space-x-reverse">
            <button
              v-for="period in chartPeriods"
              :key="period.id"
              :class="[
                'px-3 py-1 text-xs font-medium rounded-md transition-colors',
                selectedChartPeriod === period.id
                  ? 'bg-blue-500 text-white'
                  : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
              ]"
              @click="selectedChartPeriod = period.id"
            >
              {{ period.name }}
            </button>
          </div>
        </div>
        <div class="h-64 flex items-end justify-between space-x-1 space-x-reverse">
          <div
            v-for="(day, index) in trendChartData"
            :key="index"
            class="flex-1 flex flex-col items-center"
          >
            <div class="text-xs text-gray-500 mb-2">{{ day.date }}</div>
            <div
              class="w-full bg-gradient-to-t from-blue-500 to-blue-400 rounded-t"
              :style="{ height: `${(day.value / maxTrendValue) * 200}px` }"
            ></div>
            <div class="text-xs font-medium text-gray-700 mt-1">{{ day.value }}</div>
          </div>
        </div>
      </div>

      <!-- نمودار توزیع روش‌های پرداخت -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">توزیع روش‌های پرداخت</h3>
        <div class="space-y-4">
          <div
            v-for="method in paymentMethods"
            :key="method.id"
            class="flex items-center justify-between"
          >
            <div class="flex items-center">
              <div class="w-4 h-4 rounded-full ml-3" :style="{ backgroundColor: method.color }"></div>
              <span class="text-sm text-gray-700">{{ method.name }}</span>
            </div>
            <div class="flex items-center space-x-3 space-x-reverse">
              <div class="w-32 bg-gray-200 rounded-full h-3">
                <div
                  class="h-3 rounded-full"
                  :style="{ width: `${method.percentage}%`, backgroundColor: method.color }"
                ></div>
              </div>
              <span class="text-sm font-medium text-gray-900 w-12 text-left">{{ method.percentage }}%</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- جدول سفارشات در صف پردازش -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200">
      <div class="px-4 py-4 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <h3 class="text-lg font-semibold text-gray-900">سفارشات در صف پردازش اخیر</h3>
          <div class="relative">
            <input
              v-model="searchTerm"
              type="text"
              placeholder="جستجو در سفارشات..."
              class="w-48 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
            <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
              <svg class="h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
              </svg>
            </div>
          </div>
        </div>
      </div>

      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">شماره سفارش</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مشتری</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مبلغ</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">روش پرداخت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ ثبت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="order in paginatedOrders" :key="order.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                #{{ order.orderNumber }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div>
                  <div class="text-sm font-medium text-gray-900">{{ order.customerName }}</div>
                  <div class="text-sm text-gray-500">{{ order.customerPhone }}</div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ formatPrice(order.totalAmount) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusClass(order.status)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ getStatusText(order.status) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ getPaymentMethodText(order.paymentMethod) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ formatDate(order.createdAt) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <div class="flex items-center space-x-2">
                  <button class="text-blue-600 hover:text-blue-900" @click="viewOrderDetails(order)">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                    </svg>
                  </button>
                  <button class="text-green-600 hover:text-green-900">پردازش</button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      
      <!-- صفحه‌بندی -->
      <Pagination 
        :current-page="currentPage"
        :total-pages="totalPages"
        :total="inProgressOrders.length"
        :items-per-page="itemsPerPage"
        @page-changed="handlePageChange"
        @items-per-page-changed="handleItemsPerPageChange"
      />
    </div>

    <!-- مودال جزئیات سفارش -->
    <OrderDetailsModal 
      :is-open="isModalOpen"
      :order="selectedOrder"
      @close="closeModal"
      @edit="editOrder"
    />
  </div>
</template>

<script setup>
// Import کامپوننت‌ها
import Pagination from '~/components/admin/common/Pagination.vue'
import OrderDetailsModal from '~/components/admin/modals/OrderDetailsModal.vue'

// داده‌های آماری
const stats = ref({
  totalInQueue: 0,
  pendingPayment: 0,
  pendingReview: 0,
  avgProcessingTime: 0
})

const isLoading = ref(false)
const error = ref(null)

// تابع دریافت آمار از API
const fetchStats = async () => {
  isLoading.value = true
  error.value = null
  
  try {
    const data = await $fetch('/api/admin/orders/in-progress/stats')
    
    stats.value = {
      totalInQueue: data.totalInQueue || 0,
      pendingPayment: data.pendingPayment || 0,
      pendingReview: data.pendingReview || 0,
      avgProcessingTime: data.avgProcessingTime || 0
    }
  } catch (err) {
    error.value = 'خطا در دریافت آمار'
    console.error('Error fetching stats:', err)
  } finally {
    isLoading.value = false
  }
}

// دریافت آمار هنگام بارگذاری کامپوننت
onMounted(() => {
  fetchStats()
})

// دوره‌های نمودار
const chartPeriods = ref([
  { id: '7d', name: '7 روز' },
  { id: '30d', name: '30 روز' },
  { id: '90d', name: '90 روز' }
])

const selectedChartPeriod = ref('30d')

// داده‌های نمودار روند
const trendChartData = ref([
  { date: '1', value: 15 },
  { date: '2', value: 22 },
  { date: '3', value: 18 },
  { date: '4', value: 25 },
  { date: '5', value: 21 },
  { date: '6', value: 28 },
  { date: '7', value: 24 },
  { date: '8', value: 31 },
  { date: '9', value: 27 },
  { date: '10', value: 34 },
  { date: '11', value: 30 },
  { date: '12', value: 37 },
  { date: '13', value: 33 },
  { date: '14', value: 40 },
  { date: '15', value: 36 },
  { date: '16', value: 43 },
  { date: '17', value: 39 },
  { date: '18', value: 45 },
  { date: '19', value: 42 },
  { date: '20', value: 48 },
  { date: '21', value: 44 },
  { date: '22', value: 51 },
  { date: '23', value: 47 },
  { date: '24', value: 54 },
  { date: '25', value: 50 },
  { date: '26', value: 57 },
  { date: '27', value: 53 },
  { date: '28', value: 60 },
  { date: '29', value: 56 },
  { date: '30', value: 63 }
])

// وضعیت‌های سفارش
const orderStatuses = ref([
  { id: 1, name: 'در انتظار پرداخت', percentage: 38, color: '#F59E0B' },
  { id: 2, name: 'در انتظار بررسی', percentage: 29, color: '#F97316' },
  { id: 3, name: 'در حال پردازش', percentage: 20, color: '#3B82F6' },
  { id: 4, name: 'آماده ارسال', percentage: 13, color: '#10B981' }
])

// آمار زمانی پردازش
const timeStats = ref({
  fastestProcessing: '45 دقیقه',
  averageProcessingTime: '3.2 ساعت',
  longestProcessing: '12 ساعت',
  processedPerDay: 156
})

// روش‌های پرداخت
const paymentMethods = ref([
  { id: 1, name: 'پرداخت آنلاین', percentage: 45, color: '#3B82F6' },
  { id: 2, name: 'انتقال بانکی', percentage: 25, color: '#10B981' },
  { id: 3, name: 'کیف پول', percentage: 20, color: '#F59E0B' },
  { id: 4, name: 'گیفت کارت', percentage: 10, color: '#EF4444' }
])

// داده‌های سفارشات در صف پردازش
const inProgressOrders = ref([])

// متغیرهای صفحه‌بندی و جستجو
const currentPage = ref(1)
const itemsPerPage = ref(5)
const searchTerm = ref('')

// متغیرهای مودال
const isModalOpen = ref(false)
const selectedOrder = ref(null)

// محاسبات صفحه‌بندی
const totalPages = computed(() => Math.ceil(filteredOrders.value.length / itemsPerPage.value))

const filteredOrders = computed(() => {
  if (!searchTerm.value) return inProgressOrders.value
  
  return inProgressOrders.value.filter(order => 
    order.orderNumber.toLowerCase().includes(searchTerm.value.toLowerCase()) ||
    order.customerName.toLowerCase().includes(searchTerm.value.toLowerCase()) ||
    order.customerPhone.includes(searchTerm.value)
  )
})

const paginatedOrders = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage.value
  const end = start + itemsPerPage.value
  return filteredOrders.value.slice(start, end)
})

const maxTrendValue = computed(() => Math.max(...trendChartData.value.map(day => day.value)))

// متدهای صفحه‌بندی
const handlePageChange = (page) => {
  currentPage.value = page
}

const handleItemsPerPageChange = (newItemsPerPage) => {
  itemsPerPage.value = newItemsPerPage
  currentPage.value = 1
}

// متدهای کمکی
const formatDate = (date) => {
  return new Date(date).toLocaleDateString('fa-IR')
}

const formatPrice = (price) => {
  return new Intl.NumberFormat('fa-IR').format(price) + ' تومان'
}

const getStatusClass = (status) => {
  const classes = {
    processing_queue: 'bg-gray-100 text-gray-800',
    awaiting_payment: 'bg-yellow-100 text-yellow-800',
    pending_review: 'bg-orange-100 text-orange-800',
    processing: 'bg-blue-100 text-blue-800',
    ready_to_ship: 'bg-green-100 text-green-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status) => {
  const texts = {
    processing_queue: 'در صف پردازش',
    awaiting_payment: 'در انتظار پرداخت',
    pending_review: 'در انتظار بررسی',
    processing: 'در حال پردازش',
    ready_to_ship: 'آماده ارسال'
  }
  return texts[status] || status
}

const getPaymentMethodText = (method) => {
  const texts = {
    online: 'پرداخت آنلاین',
    bank_transfer: 'انتقال بانکی',
    wallet: 'کیف پول',
    gift_card: 'گیفت کارت'
  }
  return texts[method] || 'نامشخص'
}

// متدهای مودال
const viewOrderDetails = (order) => {
  selectedOrder.value = order
  isModalOpen.value = true
}

const closeModal = () => {
  isModalOpen.value = false
  selectedOrder.value = null
}

const editOrder = (_order) => {
  // اینجا می‌توانید کاربر را به صفحه ویرایش هدایت کنید
}
</script> 
