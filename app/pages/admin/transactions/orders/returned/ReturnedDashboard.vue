<template>
  <div v-if="hasAccess" class="space-y-6">
    <!-- آمار کلی -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gapx-4 py-4">
      <!-- کل سفارشات مرجوع شده -->
      <div class="bg-gradient-to-br from-red-50 to-red-100 border border-red-200 rounded-lg px-4 py-4 hover:shadow-md transition-all duration-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-red-700 mb-1">کل سفارشات مرجوع شده</p>
            <p class="text-2xl font-bold text-red-900">{{ stats.totalReturned?.toLocaleString('fa-IR') || '0' }}</p>
          </div>
          <div class="w-12 h-12 bg-red-500 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h10a8 8 0 018 8v2M3 10l6 6m-6-6l6-6"></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- در انتظار بررسی -->
      <div class="bg-gradient-to-br from-yellow-50 to-yellow-100 border border-yellow-200 rounded-lg px-4 py-4 hover:shadow-md transition-all duration-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-yellow-700 mb-1">در انتظار بررسی</p>
            <p class="text-2xl font-bold text-yellow-900">{{ stats.pendingReview?.toLocaleString('fa-IR') || '0' }}</p>
          </div>
          <div class="w-12 h-12 bg-yellow-500 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- تایید شده -->
      <div class="bg-gradient-to-br from-green-50 to-green-100 border border-green-200 rounded-lg px-4 py-4 hover:shadow-md transition-all duration-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-green-700 mb-1">تایید شده</p>
            <p class="text-2xl font-bold text-green-900">{{ stats.approved?.toLocaleString('fa-IR') || '0' }}</p>
          </div>
          <div class="w-12 h-12 bg-green-500 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- محبوب‌ترین روش ارسال -->
      <div class="bg-gradient-to-br from-indigo-50 to-indigo-100 border border-indigo-200 rounded-lg px-4 py-4 hover:shadow-md transition-all duration-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-indigo-700 mb-1">محبوب‌ترین روش ارسال</p>
            <p class="text-lg font-bold text-indigo-900">{{ getMostPopularMethod().name }}</p>
            <p class="text-xs text-indigo-600 mt-1">{{ getMostPopularMethod().count }} سفارش</p>
          </div>
          <div class="w-12 h-12 bg-indigo-500 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- آمار تفصیلی -->
    <div class="grid grid-cols-1 lg:grid-cols-3 gapx-4 py-4 mb-6">
      <!-- توزیع دلایل مرجوع -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
        <div class="bg-gradient-to-r from-red-50 to-pink-50 px-4 py-3 border-b border-gray-200">
          <div class="flex items-center">
            <div class="w-6 h-6 bg-red-500 rounded-md flex items-center justify-center ml-2">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
              </svg>
            </div>
            <h3 class="text-sm font-semibold text-gray-900">دلایل مرجوع</h3>
          </div>
        </div>

        <div class="px-4 py-4">
          <div class="space-y-4">
            <div
              v-for="reason in returnReasons"
              :key="reason.id"
              class="flex items-center justify-between"
            >
              <div class="flex items-center">
                <div class="w-3 h-3 rounded-full ml-3" :style="{ backgroundColor: reason.color }"></div>
                <span class="text-sm text-gray-700">{{ reason.name }}</span>
              </div>
              <div class="flex items-center space-x-3 space-x-reverse">
                <div class="w-32 bg-gray-200 rounded-full h-2">
                  <div
                    class="h-2 rounded-full"
                    :style="{ width: `${reason.percentage}%`, backgroundColor: reason.color }"
                  ></div>
                </div>
                <span class="text-sm font-medium text-gray-900 w-12 text-left">{{ reason.percentage }}%</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- آمار روش‌های ارسال -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
        <div class="bg-gradient-to-r from-purple-50 to-violet-50 px-4 py-3 border-b border-gray-200">
          <div class="flex items-center">
            <div class="w-6 h-6 bg-purple-500 rounded-md flex items-center justify-center ml-2">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
              </svg>
            </div>
            <h3 class="text-sm font-semibold text-gray-900">روش‌های ارسال</h3>
          </div>
        </div>

        <div class="px-4 py-4">
          <div class="space-y-4">
            <div
              v-for="(method, methodName) in stats.shippingMethods"
              :key="methodName"
              class="flex items-center justify-between"
            >
              <div class="flex items-center">
                <div class="w-3 h-3 rounded-full ml-3" :style="{ backgroundColor: method.color }"></div>
                <span class="text-sm text-gray-700">{{ methodName }}</span>
              </div>
              <div class="flex items-center space-x-3 space-x-reverse">
                <div class="w-32 bg-gray-200 rounded-full h-2">
                  <div
                    class="h-2 rounded-full"
                    :style="{ width: `${method.percentage}%`, backgroundColor: method.color }"
                  ></div>
                </div>
                <div class="text-left">
                  <div class="text-sm font-medium text-gray-900">{{ method.count }}</div>
                  <div class="text-xs text-gray-500">{{ method.percentage }}%</div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- آمار وضعیت مرجوع -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
        <div class="bg-gradient-to-r from-orange-50 to-amber-50 px-4 py-3 border-b border-gray-200">
          <div class="flex items-center">
            <div class="w-6 h-6 bg-orange-500 rounded-md flex items-center justify-center ml-2">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
            <h3 class="text-sm font-semibold text-gray-900">وضعیت مرجوع</h3>
          </div>
        </div>

        <div class="px-4 py-4">
          <div class="grid grid-cols-2 gapx-4 py-4">
            <div class="text-center">
              <div class="text-2xl font-bold text-yellow-600">{{ stats.pendingReview }}</div>
              <div class="text-xs text-gray-600 mt-1">در انتظار بررسی</div>
            </div>
            <div class="text-center">
              <div class="text-2xl font-bold text-green-600">{{ stats.approved }}</div>
              <div class="text-xs text-gray-600 mt-1">تایید شده</div>
            </div>
            <div class="text-center">
              <div class="text-2xl font-bold text-red-600">{{ stats.totalReturned - stats.approved - stats.pendingReview }}</div>
              <div class="text-xs text-gray-600 mt-1">رد شده</div>
            </div>
            <div class="text-center">
              <div class="text-2xl font-bold text-blue-600">{{ Math.round((stats.approved / stats.totalReturned) * 100) }}%</div>
              <div class="text-xs text-gray-600 mt-1">نرخ تایید</div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- نمودار روند مرجوع -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
      <div class="bg-gradient-to-r from-red-50 to-pink-50 px-4 py-3 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <div class="flex items-center">
            <div class="w-6 h-6 bg-red-500 rounded-md flex items-center justify-center ml-2">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 12l3-3 3 3 4-4M8 21l4-4 4 4M3 4h18M4 4h16v12a1 1 0 01-1 1H5a1 1 0 01-1-1V4z"></path>
              </svg>
            </div>
            <h3 class="text-sm font-semibold text-gray-900">روند مرجوع در ۷ روز گذشته</h3>
          </div>
          <div class="flex items-center space-x-2 space-x-reverse">
            <button
              v-for="period in chartPeriods"
              :key="period.id"
              :class="[
                'px-3 py-1 text-xs font-medium rounded-md transition-colors',
                selectedPeriod === period.id
                  ? 'bg-red-500 text-white'
                  : 'bg-white text-gray-600 hover:bg-gray-50'
              ]"
              @click="selectedPeriod = period.id"
            >
              {{ period.name }}
            </button>
          </div>
        </div>
      </div>

      <div class="px-4 py-4">
        <div class="h-64 flex items-end justify-between space-x-2 space-x-reverse">
          <div
            v-for="(day, index) in chartData"
            :key="index"
            class="flex-1 flex flex-col items-center"
          >
            <div class="text-xs text-gray-500 mb-2">{{ day.date }}</div>
            <div
              class="w-full bg-gradient-to-t from-red-500 to-red-400 rounded-t"
              :style="{ height: `${(day.value / maxChartValue) * 200}px` }"
            ></div>
            <div class="text-xs font-medium text-gray-700 mt-1">{{ day.value }}</div>
          </div>
        </div>
      </div>
    </div>

    <!-- جدول سفارشات اخیر -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200">
      <div class="px-4 py-4 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <h3 class="text-lg font-semibold text-gray-900">سفارشات مرجوع شده اخیر</h3>
          <div class="relative">
            <input
              v-model="searchTerm"
              type="text"
              placeholder="جستجو در سفارشات..."
              class="w-48 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-red-500 focus:border-red-500"
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
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">دلیل مرجوع</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
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
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ getReasonText(order.reason) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusClass(order.status)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ getStatusText(order.status) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <div class="flex items-center space-x-2">
                  <button class="text-blue-600 hover:text-blue-900" @click="viewOrderDetails(order)">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                    </svg>
                  </button>
                  <button class="text-green-600 hover:text-green-900" @click="processReturn(order)">پردازش</button>
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
        :total="filteredOrders.length"
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

<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
</script>

<script setup lang="ts">
// Import کامپوننت‌ها
import { computed, onMounted, ref, watch } from 'vue'
import Pagination from '~/components/admin/common/Pagination.vue'
import OrderDetailsModal from '~/components/admin/modals/OrderDetailsModal.vue'
import { useAuth } from '~/composables/useAuth'

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

// Props
interface ReturnedOrder {
  id: string | number
  orderNumber: string
  customerName: string
  customerPhone: string
  totalAmount: number
  reason: string
  status: string
}

interface ShippingMethod {
  count: number
  percentage: number
  color: string
}

const props = defineProps<{
  stats: {
    totalReturned: number
    pendingReview: number
    approved: number
    shippingMethods?: Record<string, ShippingMethod>
  }
  orders: ReturnedOrder[]
}>()

// دوره‌های نمودار
const chartPeriods = ref([
  { id: '7d', name: '7 روز' },
  { id: '30d', name: '30 روز' },
  { id: '90d', name: '90 روز' }
])

const selectedPeriod = ref('7d')

// داده‌های نمودار
const chartData = ref([
  { date: 'شنبه', value: 3 },
  { date: 'یکشنبه', value: 5 },
  { date: 'دوشنبه', value: 2 },
  { date: 'سه‌شنبه', value: 7 },
  { date: 'چهارشنبه', value: 4 },
  { date: 'پنج‌شنبه', value: 6 },
  { date: 'جمعه', value: 3 }
])

// دلایل مرجوع
const returnReasons = ref([
  { id: 1, name: 'آسیب دیده', percentage: 35, color: '#EF4444' },
  { id: 2, name: 'محصول اشتباه', percentage: 25, color: '#F59E0B' },
  { id: 3, name: 'عدم رضایت', percentage: 20, color: '#8B5CF6' },
  { id: 4, name: 'مشکل سایز', percentage: 15, color: '#3B82F6' },
  { id: 5, name: 'دیگر', percentage: 5, color: '#6B7280' }
])

// متغیر جستجو
const searchTerm = ref('')

// متغیرهای مودال
const isModalOpen = ref(false)
const selectedOrder = ref(null)

// متغیرهای صفحه‌بندی
const currentPage = ref(1)
const itemsPerPage = ref(5)

// محاسبات
const maxChartValue = computed(() => {
  return Math.max(...chartData.value.map(item => item.value))
})

const filteredOrders = computed(() => {
  if (!searchTerm.value) return props.orders
  
  return props.orders.filter(order => 
    order.orderNumber.toLowerCase().includes(searchTerm.value.toLowerCase()) ||
    order.customerName.toLowerCase().includes(searchTerm.value.toLowerCase())
  )
})

const totalPages = computed(() => Math.ceil(filteredOrders.value.length / itemsPerPage.value))

const paginatedOrders = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage.value
  const end = start + itemsPerPage.value
  return filteredOrders.value.slice(start, end)
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

// متدهای عملیاتی
const processReturn = (_order) => {
  // Process return logic
}

// متدهای صفحه‌بندی
const handlePageChange = (page) => {
  currentPage.value = page
}

const handleItemsPerPageChange = (newItemsPerPage) => {
  itemsPerPage.value = newItemsPerPage
  currentPage.value = 1
}

// متدهای کمکی
const formatPrice = (price) => {
  return new Intl.NumberFormat('fa-IR').format(price || 0) + ' تومان'
}

const getStatusClass = (status) => {
  const classes = {
    pending: 'bg-yellow-100 text-yellow-800',
    approved: 'bg-green-100 text-green-800',
    rejected: 'bg-red-100 text-red-800',
    completed: 'bg-blue-100 text-blue-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status) => {
  const texts = {
    pending: 'در انتظار بررسی',
    approved: 'تایید شده',
    rejected: 'رد شده',
    completed: 'تکمیل شده'
  }
  return texts[status] || status
}

const getReasonText = (reason) => {
  const texts = {
    damaged: 'آسیب دیده',
    wrong_item: 'محصول اشتباه',
    not_satisfied: 'عدم رضایت',
    size_issue: 'مشکل سایز'
  }
  return texts[reason] || reason
}

// تابع محاسبه محبوب‌ترین روش ارسال
const getMostPopularMethod = () => {
  if (!props.stats.shippingMethods) {
    return { name: 'نامشخص', count: 0 }
  }
  
  const methods = Object.entries(props.stats.shippingMethods)
  const mostPopular = methods.reduce((max, [name, data]) => {
    return data.count > max.count ? { name, count: data.count } : max
  }, { name: '', count: 0 })
  
  return mostPopular
}
</script> 
