<template>
  <div>
    <!-- آمار کلی -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gapx-4 py-4 mb-6">
      <!-- کل سفارشات در حال انجام -->
      <div class="bg-gradient-to-br from-blue-50 to-blue-100 border border-blue-200 rounded-lg px-4 py-4 hover:shadow-md transition-all duration-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-blue-700 mb-1">کل سفارشات در حال انجام</p>
            <p class="text-2xl font-bold text-blue-900">{{ stats.totalProcessing.toLocaleString('fa-IR') }}</p>
            <p class="text-xs text-blue-600 mt-1">در حال پردازش</p>
          </div>
          <div class="w-12 h-12 bg-blue-500 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- مبلغ کل در حال انجام -->
      <div class="bg-gradient-to-br from-green-50 to-green-100 border border-green-200 rounded-lg px-4 py-4 hover:shadow-md transition-all duration-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-green-700 mb-1">مبلغ کل در حال انجام</p>
            <p class="text-2xl font-bold text-green-900">{{ formatPrice(stats.totalAmount) }}</p>
            <p class="text-xs text-green-600 mt-1">در حال پردازش</p>
          </div>
          <div class="w-12 h-12 bg-green-500 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- نرخ تکمیل -->
      <div class="bg-gradient-to-br from-purple-50 to-purple-100 border border-purple-200 rounded-lg px-4 py-4 hover:shadow-md transition-all duration-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-purple-700 mb-1">نرخ تکمیل</p>
            <p class="text-2xl font-bold text-purple-900">{{ stats.completionRate }}%</p>
            <p class="text-xs text-purple-600 mt-1">سفارشات تکمیل شده</p>
          </div>
          <div class="w-12 h-12 bg-purple-500 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- متوسط زمان پردازش -->
      <div class="bg-gradient-to-br from-orange-50 to-orange-100 border border-orange-200 rounded-lg px-4 py-4 hover:shadow-md transition-all duration-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-orange-700 mb-1">متوسط زمان پردازش</p>
            <p class="text-2xl font-bold text-orange-900">{{ stats.avgProcessingTime }} ساعت</p>
            <p class="text-xs text-orange-600 mt-1">از شروع تا تکمیل</p>
          </div>
          <div class="w-12 h-12 bg-orange-500 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- آمار تفصیلی -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gapx-4 py-4 mb-6">
      <!-- توزیع مراحل پردازش -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
        <div class="bg-gradient-to-r from-blue-50 to-indigo-50 px-4 py-3 border-b border-gray-200">
          <div class="flex items-center">
            <div class="w-6 h-6 bg-blue-500 rounded-md flex items-center justify-center ml-2">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
              </svg>
            </div>
            <h3 class="text-sm font-semibold text-gray-900">توزیع مراحل پردازش</h3>
          </div>
        </div>

        <div class="px-4 py-4">
          <div class="space-y-4">
            <div
              v-for="stage in processingStages"
              :key="stage.id"
              class="flex items-center justify-between"
            >
              <div class="flex items-center">
                <div class="w-3 h-3 rounded-full ml-3" :style="{ backgroundColor: stage.color }"></div>
                <span class="text-sm text-gray-700">{{ stage.name }}</span>
              </div>
              <div class="flex items-center space-x-3 space-x-reverse">
                <div class="w-32 bg-gray-200 rounded-full h-2">
                  <div
                    class="h-2 rounded-full"
                    :style="{ width: `${stage.percentage}%`, backgroundColor: stage.color }"
                  ></div>
                </div>
                <span class="text-sm font-medium text-gray-900 w-12 text-left">{{ stage.percentage }}%</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- آمار زمانی پردازش -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
        <div class="bg-gradient-to-r from-green-50 to-emerald-50 px-4 py-3 border-b border-gray-200">
          <div class="flex items-center">
            <div class="w-6 h-6 bg-green-500 rounded-md flex items-center justify-center ml-2">
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
              <div class="text-2xl font-bold text-green-600">{{ formatTime(timeStats.avgProcessingTime) }}</div>
              <div class="text-xs text-gray-600 mt-1">متوسط زمان پردازش</div>
            </div>
            <div class="text-center">
              <div class="text-2xl font-bold text-blue-600">{{ formatTime(timeStats.fastestProcessing) }}</div>
              <div class="text-xs text-gray-600 mt-1">سریع ترین پردازش</div>
            </div>
            <div class="text-center">
              <div class="text-2xl font-bold text-purple-600">{{ Math.round(timeStats.processingPerDay) }}</div>
              <div class="text-xs text-gray-600 mt-1">پردازش در روز</div>
            </div>
            <div class="text-center">
              <div class="text-2xl font-bold text-orange-600">{{ formatTime(timeStats.longestProcessing) }}</div>
              <div class="text-xs text-gray-600 mt-1">طولانی ترین پردازش</div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- نمودار روند 30 روزه -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4 mb-6">
      <div class="flex items-center justify-between mb-4">
        <h3 class="text-lg font-semibold text-gray-900">روند پردازش 30 روزه</h3>
        <div class="flex items-center space-x-2 space-x-reverse">
          <span class="text-sm text-gray-500">واحد: تعداد سفارشات</span>
        </div>
      </div>

      <div class="relative h-64">
        <!-- نمایش پیام عدم وجود داده -->
        <div v-if="trendData.length === 0" class="flex items-center justify-center h-full text-gray-500">
          <div class="text-center">
            <svg class="w-12 h-12 mx-auto mb-4 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
            <p class="text-sm">هیچ داده‌ای برای نمایش وجود ندارد</p>
          </div>
        </div>

        <!-- نمودار زمانی -->
        <div v-else>
          <!-- خطوط راهنما -->
          <div class="absolute inset-0 flex flex-col justify-between text-xs text-gray-400">
            <div class="border-b border-gray-100 pb-1">{{ maxTrendValue }}</div>
            <div class="border-b border-gray-100 pb-1">{{ Math.floor(maxTrendValue * 0.75) }}</div>
            <div class="border-b border-gray-100 pb-1">{{ Math.floor(maxTrendValue * 0.5) }}</div>
            <div class="border-b border-gray-100 pb-1">{{ Math.floor(maxTrendValue * 0.25) }}</div>
            <div class="pb-1">0</div>
          </div>

          <!-- ستون‌های نمودار -->
          <div class="absolute inset-0 flex items-end justify-between space-x-1 space-x-reverse pr-16">
            <div
              v-for="(day, index) in trendData"
              :key="index"
              class="flex-1 flex flex-col items-center relative group"
            >
              <!-- ستون -->
              <div
                class="w-full bg-gradient-to-t from-blue-500 to-blue-400 rounded-t transition-all duration-200 hover:from-blue-600 hover:to-blue-500"
                :style="{ height: maxTrendValue > 0 ? `${(day.value / maxTrendValue) * 240}px` : '0px' }"
              ></div>

              <!-- نام روز -->
              <div class="text-xs text-gray-600 mt-2 text-center">{{ day.name }}</div>

              <!-- مقدار در tooltip -->
              <div class="absolute bottom-full mb-2 opacity-0 group-hover:opacity-100 transition-opacity duration-200 bg-gray-800 text-white text-xs px-2 py-1 rounded whitespace-nowrap z-10">
                {{ day.value }} سفارش
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- راهنمای نمودار -->
      <div class="mt-4 text-xs text-gray-500 text-center">
        مقادیر در tooltip نمایش داده می‌شوند
      </div>
    </div>

    <!-- سفارشات اخیر در حال انجام -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
      <div class="bg-gradient-to-r from-slate-50 to-gray-50 px-4 py-3 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <div class="flex items-center">
            <div class="w-6 h-6 bg-slate-600 rounded-md flex items-center justify-center ml-2">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16"></path>
              </svg>
            </div>
            <h3 class="text-sm font-semibold text-gray-900">سفارشات اخیر در حال انجام</h3>
          </div>
          
          <div class="flex items-center space-x-2 space-x-reverse">
            <div class="relative">
              <input
                v-model="searchTerm"
                type="text"
                placeholder="جستجو در سفارشات..."
                class="w-48 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-sm"
              />
              <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                <svg class="h-4 w-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
                </svg>
              </div>
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
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ شروع</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="order in filteredRecentOrders" :key="order.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                #{{ order.orderNumber }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div>
                  <div class="text-sm font-medium text-gray-900">{{ order.customerName || 'نامشخص' }}</div>
                  <div class="text-sm text-gray-500">{{ order.customerPhone || 'شماره نامشخص' }}</div>
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
                  <button class="text-green-600 hover:text-green-900" @click="updateProgress(order)">بروزرسانی</button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
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
// Import کامپوننت مودال
import OrderDetailsModal from '~/components/admin/modals/OrderDetailsModal.vue'

// داده‌های آماری
const stats = ref({
  totalProcessing: 0,
  totalAmount: 0,
  completionRate: 0,
  avgProcessingTime: 0
})

// آمار زمانی واقعی
const timeStats = ref({
  avgProcessingTime: 0,
  fastestProcessing: 0,
  processingPerDay: 0,
  longestProcessing: 0
})

// مراحل پردازش واقعی
const processingStages = ref([])

// داده‌های روند واقعی
const trendData = ref([])

// سفارشات اخیر واقعی
const recentOrders = ref([])

const isLoading = ref(false)
const error = ref(null)

// متغیر جستجو
const searchTerm = ref('')

// متغیرهای مودال
const isModalOpen = ref(false)
const selectedOrder = ref(null)

// تابع دریافت آمار از API
const fetchStats = async () => {
  isLoading.value = true
  error.value = null
  
  try {
    const data = await $fetch('/api/admin/orders/processing/stats')
    
    // آمار اصلی
    stats.value = {
      totalProcessing: data.totalProcessing || 0,
      totalAmount: data.totalAmount || 0,
      completionRate: data.completionRate || 0,
      avgProcessingTime: data.avgProcessingTime || 0
    }

    // آمار زمانی
    if (data.timeStats) {
      timeStats.value = {
        avgProcessingTime: data.timeStats.avgProcessingTime || 0,
        fastestProcessing: data.timeStats.fastestProcessing || 0,
        processingPerDay: data.timeStats.processingPerDay || 0,
        longestProcessing: data.timeStats.longestProcessing || 0
      }
    }

    // مراحل پردازش از توزیع وضعیت‌ها
    if (data.statusDistribution && data.statusDistribution.length > 0) {
      const total = data.statusDistribution.reduce((sum, item) => sum + item.count, 0)
      processingStages.value = data.statusDistribution.map((item, index) => ({
        id: index + 1,
        name: getStatusText(item.status),
        percentage: total > 0 ? Math.round((item.count / total) * 100) : 0,
        color: getStatusColor(item.status)
      }))
    } else {
      // اگر داده‌ای وجود ندارد، نمایش پیام مناسب
      processingStages.value = [{
        id: 1,
        name: 'هیچ سفارشی در حال انجام نیست',
        percentage: 100,
        color: '#6B7280'
      }]
    }

    // داده‌های روند
    if (data.trendData && data.trendData.length > 0) {
      trendData.value = data.trendData.map((item, index) => ({
        name: (index + 1).toString(),
        value: item.count
      }))
    } else {
      trendData.value = []
    }

    // سفارشات اخیر
    if (data.recentOrders && data.recentOrders.length > 0) {
      recentOrders.value = data.recentOrders.map(order => ({
        id: order.id,
        orderNumber: order.orderNumber,
        customerName: order.customerName || 'نامشخص',
        customerPhone: order.customerPhone || 'شماره نامشخص',
        totalAmount: order.totalAmount,
        status: order.status,
        paymentMethod: order.paymentMethod,
        createdAt: order.createdAt
      }))
    } else {
      recentOrders.value = []
    }

  } catch (err) {
    error.value = 'خطا در دریافت آمار'
    console.error('Error fetching stats:', err)
    
    // تنظیم مقادیر پیش‌فرض در صورت خطا
    processingStages.value = [{
      id: 1,
      name: 'خطا در دریافت داده‌ها',
      percentage: 100,
      color: '#EF4444'
    }]
    trendData.value = []
    recentOrders.value = []
  } finally {
    isLoading.value = false
  }
}

// دریافت آمار هنگام بارگذاری کامپوننت
onMounted(() => {
  fetchStats()
})

// محاسبات
const maxTrendValue = computed(() => {
  if (trendData.value.length === 0) return 0
  return Math.max(...trendData.value.map(day => day.value))
})

const filteredRecentOrders = computed(() => {
  if (!searchTerm.value) return recentOrders.value
  
  return recentOrders.value.filter(order => 
    order.orderNumber.toLowerCase().includes(searchTerm.value.toLowerCase()) ||
    order.customerName.toLowerCase().includes(searchTerm.value.toLowerCase()) ||
    order.customerPhone.includes(searchTerm.value)
  )
})

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

const updateProgress = (_order) => {
  // Update progress logic
}

// متدهای کمکی
const formatPrice = (price) => {
  return new Intl.NumberFormat('fa-IR').format(price) + ' تومان'
}

const formatDate = (date) => {
  return new Date(date).toLocaleDateString('fa-IR')
}

const getStatusClass = (status) => {
  const classes = {
    processing: 'bg-blue-100 text-blue-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getPaymentMethodText = (method) => {
  const texts = {
    online: 'پرداخت آنلاین',
    bank_transfer: 'انتقال بانکی',
    wallet: 'کیف پول',
    gift_card: 'گیفت کارت',
    cash: 'پرداخت نقدی'
  }
  return texts[method] || 'نامشخص'
}

// تابع کمکی برای رنگ وضعیت
const getStatusColor = (status) => {
  const colors = {
    processing: '#3B82F6',
    processing_queue: '#6B7280',
    awaiting_payment: '#F59E0B',
    pending_review: '#8B5CF6',
    ready_to_ship: '#10B981',
    shipped: '#059669',
    delivered: '#047857',
    cancelled: '#EF4444'
  }
  return colors[status] || '#6B7280'
}

// تابع کمکی برای متن وضعیت
const getStatusText = (status) => {
  const texts = {
    processing: 'در حال انجام',
    processing_queue: 'در صف پردازش',
    awaiting_payment: 'در انتظار پرداخت',
    pending_review: 'در انتظار بررسی',
    ready_to_ship: 'آماده ارسال',
    shipped: 'ارسال شده',
    delivered: 'تحویل شده',
    cancelled: 'لغو شده'
  }
  return texts[status] || status
}

// تابع کمکی برای فرمت زمان
const formatTime = (hours) => {
  if (!hours || hours === 0) return '0 ساعت'
  
  if (hours < 1) {
    const minutes = Math.round(hours * 60)
    return `${minutes} دقیقه`
  } else if (hours < 24) {
    const wholeHours = Math.floor(hours)
    const minutes = Math.round((hours - wholeHours) * 60)
    if (minutes > 0) {
      return `${wholeHours} ساعت ${minutes} دقیقه`
    }
    return `${wholeHours} ساعت`
  } else {
    const days = Math.floor(hours / 24)
    const remainingHours = Math.floor(hours % 24)
    if (remainingHours > 0) {
      return `${days} روز ${remainingHours} ساعت`
    }
    return `${days} روز`
  }
}
</script> 
