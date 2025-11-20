<template>
  <div>
    <!-- آمار کلی سفارشات مسترد شده -->
    <div class="rounded-lg shadow-md border border-gray-200 overflow-hidden mb-6">
      <div class="bg-gradient-to-r from-purple-50 to-indigo-50 px-4 py-3 border-b border-gray-200">
        <div class="flex items-center">
          <div class="w-6 h-6 bg-purple-500 rounded-md flex items-center justify-center ml-2">
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h10a8 8 0 018 8v2M3 10l6 6m-6-6l6-6"></path>
            </svg>
          </div>
          <h3 class="text-sm font-semibold text-gray-900">آمار کلی سفارشات مسترد شده</h3>
        </div>
      </div>

      <div class="px-4 py-4">
        <div class="grid grid-cols-1 md:grid-cols-4 gapx-4 py-4">
          <!-- کل سفارشات مسترد شده -->
          <div class="bg-gradient-to-br from-purple-50 to-purple-100 border border-purple-200 rounded-lg px-4 py-4 hover:shadow-md transition-all duration-200">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-purple-700 mb-1">کل سفارشات مسترد شده</p>
                <p class="text-2xl font-bold text-purple-900">{{ stats.totalRefunded?.toLocaleString('fa-IR') || '0' }}</p>
                <p class="text-xs text-purple-600 mt-1">در 30 روز گذشته</p>
              </div>
              <div class="w-12 h-12 bg-purple-500 rounded-lg flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h10a8 8 0 018 8v2M3 10l6 6m-6-6l6-6"></path>
                </svg>
              </div>
            </div>
          </div>

          <!-- مبلغ کل مسترد شده -->
          <div class="bg-gradient-to-br from-green-50 to-green-100 border border-green-200 rounded-lg px-4 py-4 hover:shadow-md transition-all duration-200">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-green-700 mb-1">مبلغ کل مسترد شده</p>
                <p class="text-2xl font-bold text-green-900">{{ formatPrice(stats.totalRefundAmount || 0) }}</p>
                <p class="text-xs text-green-600 mt-1">در 30 روز گذشته</p>
              </div>
              <div class="w-12 h-12 bg-green-500 rounded-lg flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
                </svg>
              </div>
            </div>
          </div>

          <!-- نرخ مسترد -->
          <div class="bg-gradient-to-br from-blue-50 to-blue-100 border border-blue-200 rounded-lg px-4 py-4 hover:shadow-md transition-all duration-200">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-blue-700 mb-1">نرخ مسترد</p>
                <p class="text-2xl font-bold text-blue-900">{{ stats.refundRate }}%</p>
                <p class="text-xs text-blue-600 mt-1">از کل سفارشات</p>
              </div>
              <div class="w-12 h-12 bg-blue-500 rounded-lg flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
                </svg>
              </div>
            </div>
          </div>

          <!-- متوسط زمان مسترد -->
          <div class="bg-gradient-to-br from-orange-50 to-orange-100 border border-orange-200 rounded-lg px-4 py-4 hover:shadow-md transition-all duration-200">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-orange-700 mb-1">متوسط زمان مسترد</p>
                <p class="text-2xl font-bold text-orange-900">{{ stats.avgRefundTime || 0 }} روز</p>
                <p class="text-xs text-orange-600 mt-1">از درخواست تا پرداخت</p>
              </div>
              <div class="w-12 h-12 bg-orange-500 rounded-lg flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- آمار تفصیلی -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gapx-4 py-4 mb-6">
      <!-- توزیع دلایل مسترد -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
        <div class="bg-gradient-to-r from-red-50 to-pink-50 px-4 py-3 border-b border-gray-200">
          <div class="flex items-center">
            <div class="w-6 h-6 bg-red-500 rounded-md flex items-center justify-center ml-2">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.172 16.172a4 4 0 015.656 0M9 12h6m-6-4h6m2 5.291A7.962 7.962 0 0112 15c-2.34 0-4.47-.881-6.08-2.33"></path>
              </svg>
            </div>
            <h3 class="text-sm font-semibold text-gray-900">توزیع دلایل مسترد</h3>
          </div>
        </div>

        <div class="px-4 py-4">
          <div class="space-y-4">
            <div
              v-for="reason in refundReasons"
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
                    class="h-2 rounded-full transition-all duration-300"
                    :style="{ width: `${reason.percentage}%`, backgroundColor: reason.color }"
                  ></div>
                </div>
                <span class="text-sm font-medium text-gray-900 w-12 text-left">{{ reason.percentage }}%</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- آمار زمانی مسترد -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
        <div class="bg-gradient-to-r from-blue-50 to-indigo-50 px-4 py-3 border-b border-gray-200">
          <div class="flex items-center">
            <div class="w-6 h-6 bg-blue-500 rounded-md flex items-center justify-center ml-2">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
            <h3 class="text-sm font-semibold text-gray-900">آمار زمانی مسترد</h3>
          </div>
        </div>

        <div class="px-4 py-4">
          <div class="grid grid-cols-2 gapx-4 py-4">
            <div class="text-center">
              <div class="text-2xl font-bold text-green-600">{{ timeStats.avgRefundTime }}</div>
              <div class="text-xs text-gray-600 mt-1">متوسط زمان مسترد</div>
            </div>
            <div class="text-center">
              <div class="text-2xl font-bold text-blue-600">{{ timeStats.fastestRefund }}</div>
              <div class="text-xs text-gray-600 mt-1">سریع ترین مسترد</div>
            </div>
            <div class="text-center">
              <div class="text-2xl font-bold text-purple-600">{{ timeStats.refundsPerDay }}</div>
              <div class="text-xs text-gray-600 mt-1">مسترد در روز</div>
            </div>
            <div class="text-center">
              <div class="text-2xl font-bold text-orange-600">{{ timeStats.longestRefund }}</div>
              <div class="text-xs text-gray-600 mt-1">دیرترین مسترد</div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- روش‌های پرداخت مسترد -->
    <div class="mb-6">
      <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
        <div class="bg-gradient-to-r from-teal-50 to-cyan-50 px-4 py-3 border-b border-gray-200">
          <div class="flex items-center">
            <div class="w-6 h-6 bg-teal-500 rounded-md flex items-center justify-center ml-2">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"></path>
              </svg>
            </div>
            <h3 class="text-sm font-semibold text-gray-900">روش‌های پرداخت مسترد</h3>
          </div>
        </div>

        <div class="px-4 py-4">
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gapx-4 py-4">
            <div
              v-for="method in paymentMethods"
              :key="method.id"
              class="bg-white border border-gray-200 rounded-lg px-4 py-4 shadow-sm hover:shadow-md transition-shadow"
            >
              <div class="flex items-center justify-between mb-3">
                <div class="flex items-center">
                  <div class="w-4 h-4 rounded-full ml-2" :style="{ backgroundColor: method.color }"></div>
                  <span class="text-sm font-medium text-gray-900">{{ method.name }}</span>
                </div>
                <span class="text-lg font-bold" :style="{ color: method.color }">{{ method.percentage }}%</span>
              </div>
              <div class="w-full bg-gray-200 rounded-full h-2">
                <div
                  class="h-2 rounded-full transition-all duration-300"
                  :style="{ width: `${method.percentage}%`, backgroundColor: method.color }"
                ></div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- نمودار روند مسترد -->
    <div class="rounded-lg shadow-md border border-gray-200 overflow-hidden mb-6">
      <div class="bg-gradient-to-r from-indigo-50 to-purple-50 px-4 py-3 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <div class="flex items-center">
            <div class="w-6 h-6 bg-indigo-500 rounded-md flex items-center justify-center ml-2">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 12l3-3 3 3 4-4M8 21l4-4 4 4M3 4h18M4 4h16v12a1 1 0 01-1 1H5a1 1 0 01-1-1V4z"></path>
              </svg>
            </div>
            <h3 class="text-sm font-semibold text-gray-900">آمار روند مسترد</h3>
          </div>
          <div class="flex items-center space-x-2 space-x-reverse">
            <button
              v-for="period in chartPeriods"
              :key="period.id"
              :class="[
                'px-3 py-1 text-xs font-medium rounded-md transition-colors',
                selectedPeriod === period.id
                  ? 'bg-indigo-500 text-white'
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
              class="w-full bg-gradient-to-t from-indigo-500 to-indigo-400 rounded-t"
              :style="{ height: `${(day.value / maxChartValue) * 200}px` }"
            ></div>
            <div class="text-xs font-medium text-gray-700 mt-1">{{ day.value }}</div>
          </div>
        </div>
      </div>
    </div>





    <!-- جدول سفارشات مسترد شده -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200">
                        <div class="px-4 py-4 border-b border-gray-200">
                    <div class="flex items-center justify-between">
                      <h3 class="text-lg font-semibold text-gray-900">سفارشات مسترد شده اخیر</h3>
                      <div class="relative">
                        <input
                          v-model="searchTerm"
                          type="text"
                          placeholder="جستجو در سفارشات..."
                          class="w-48 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500"
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
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">دلیل مسترد</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">روش انتقال</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ مسترد</th>
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
                <span :class="getReasonClass(order.refundReason)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ getReasonText(order.refundReason) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ getRefundMethodText(order.refundMethod) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ formatDate(order.refundedAt) }}
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
                  <AdminPagination 
                    :current-page="currentPage"
                    :total-pages="totalPages"
                    :total="refundedOrders.length"
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
// Import کامپوننت مودال
import OrderDetailsModal from '~/components/admin/modals/OrderDetailsModal.vue'

// مقداردهی اولیه و دریافت داده‌ها
onMounted(() => {
  fetchStats()
})

// داده‌های آماری
const stats = ref({})

// متغیرهای مدیریت وضعیت بارگذاری
const loading = ref(false)
const error = ref(null)

// تابع دریافت آمار مسترد
const fetchStats = async () => {
  try {
    loading.value = true
    error.value = null

    const { data } = await $fetch('/api/admin/orders/refunded/stats', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    })

    // تبدیل داده‌های API به فرمت مورد نیاز کامپوننت
    stats.value = {
      totalRefunded: data.totalRefunded || 0,
      totalRefundAmount: (data.todaySales || 0) + (data.weeklySales || 0) + (data.monthlySales || 0),
      refundRate: data.averageOrder || 0,
      avgRefundTime: 0 // TODO: محاسبه متوسط زمان مسترد
    }

  } catch (err) {
    console.error('خطا در دریافت آمار:', err)
    error.value = err.message || 'خطا در دریافت آمار'
  } finally {
    loading.value = false
  }
}

// دوره‌های نمودار
const chartPeriods = ref([
  { id: '7d', name: '7 روز' },
  { id: '30d', name: '30 روز' },
  { id: '90d', name: '90 روز' }
])

const selectedPeriod = ref('7d')

// داده‌های نمودار
const chartData = ref([
  { date: 'شنبه', value: 12 },
  { date: 'یکشنبه', value: 18 },
  { date: 'دوشنبه', value: 15 },
  { date: 'سه‌شنبه', value: 22 },
  { date: 'چهارشنبه', value: 19 },
  { date: 'پنج‌شنبه', value: 25 },
  { date: 'جمعه', value: 16 }
])

// دلایل مسترد
const refundReasons = ref([
  { id: 1, name: 'درخواست مشتری', percentage: 45, color: '#EF4444' },
  { id: 2, name: 'عدم موجودی', percentage: 28, color: '#F59E0B' },
  { id: 3, name: 'مشکل پرداخت', percentage: 18, color: '#EAB308' },
  { id: 4, name: 'مشکل سیستمی', percentage: 9, color: '#6B7280' }
])

// آمار زمانی مسترد
const timeStats = ref({
  avgRefundTime: '2.3 ساعت',
  fastestRefund: '15 دقیقه',
  refundsPerDay: '123',
  longestRefund: '24 ساعت'
})

// روش‌های پرداخت
const paymentMethods = ref([
  { id: 1, name: 'انتقال بانکی', percentage: 40, color: '#3B82F6' },
  { id: 2, name: 'کیف پول', percentage: 30, color: '#10B981' },
  { id: 3, name: 'گیفت کارت', percentage: 20, color: '#F59E0B' },
  { id: 4, name: 'کارت اعتباری', percentage: 10, color: '#EF4444' }
])



// داده‌های سفارشات مسترد شده
const refundedOrders = ref([
  {
    id: 1,
    orderNumber: 'ORD-001',
    customerName: 'علی احمدی',
    customerPhone: '09123456789',
    totalAmount: 1250000,
    refundReason: 'customer',
    refundMethod: 'bank_transfer',
    refundedAt: '2024-01-15'
  },
  {
    id: 2,
    orderNumber: 'ORD-002',
    customerName: 'فاطمه محمدی',
    customerPhone: '09187654321',
    totalAmount: 890000,
    refundReason: 'inventory',
    refundMethod: 'wallet',
    refundedAt: '2024-01-14'
  },
  {
    id: 3,
    orderNumber: 'ORD-003',
    customerName: 'محمد رضایی',
    customerPhone: '09111222333',
    totalAmount: 2100000,
    refundReason: 'payment',
    refundMethod: 'gift_card',
    refundedAt: '2024-01-13'
  },
  {
    id: 4,
    orderNumber: 'ORD-004',
    customerName: 'زهرا کریمی',
    customerPhone: '09144555666',
    totalAmount: 1560000,
    refundReason: 'system',
    refundMethod: 'credit_card',
    refundedAt: '2024-01-12'
  },
  {
    id: 5,
    orderNumber: 'ORD-005',
    customerName: 'حسین نوری',
    customerPhone: '09177888999',
    totalAmount: 980000,
    refundReason: 'customer',
    refundMethod: 'bank_transfer',
    refundedAt: '2024-01-11'
  }
])

// محاسبه حداکثر مقدار نمودار
const maxChartValue = computed(() => {
  return Math.max(...chartData.value.map(item => item.value))
})

// تابع فرمت قیمت
const formatPrice = (price) => {
  return new Intl.NumberFormat('fa-IR').format(price) + ' تومان'
}

// تابع فرمت تاریخ
const formatDate = (date) => {
  return new Date(date).toLocaleDateString('fa-IR')
}

// تابع دریافت کلاس دلیل مسترد
const getReasonClass = (reason) => {
  const classes = {
    customer: 'bg-red-100 text-red-800',
    inventory: 'bg-orange-100 text-orange-800',
    payment: 'bg-yellow-100 text-yellow-800',
    system: 'bg-gray-100 text-gray-800'
  }
  return classes[reason] || 'bg-gray-100 text-gray-800'
}

            // تابع دریافت متن دلیل مسترد
            const getReasonText = (reason) => {
              const texts = {
                customer: 'درخواست مشتری',
                inventory: 'عدم موجودی',
                payment: 'مشکل پرداخت',
                system: 'مشکل سیستمی'
              }
              return texts[reason] || 'نامشخص'
            }

            // تابع دریافت متن روش مسترد
            const getRefundMethodText = (method) => {
              const texts = {
                bank_transfer: 'انتقال بانکی',
                wallet: 'کیف پول',
                gift_card: 'گیفت کارت',
                credit_card: 'کارت اعتباری',
                cash: 'نقدی',
                check: 'چک'
              }
              return texts[method] || 'نامشخص'
            }

            // متغیرهای صفحه‌بندی و جستجو
            const currentPage = ref(1)
            const itemsPerPage = ref(5)
            const searchTerm = ref('')

            // متغیرهای مودال
            const isModalOpen = ref(false)
            const selectedOrder = ref(null)

            // محاسبه تعداد صفحات
            const totalPages = computed(() => {
              return Math.ceil(filteredOrders.value.length / itemsPerPage.value)
            })

            // محاسبه سفارشات فیلتر شده بر اساس جستجو
            const filteredOrders = computed(() => {
              if (!searchTerm.value) {
                return refundedOrders.value
              }
              return refundedOrders.value.filter(order => 
                order.orderNumber.toLowerCase().includes(searchTerm.value.toLowerCase()) ||
                order.customerName.toLowerCase().includes(searchTerm.value.toLowerCase()) ||
                order.customerPhone.includes(searchTerm.value)
              )
            })

            // محاسبه سفارشات نمایش داده شده در صفحه فعلی
            const paginatedOrders = computed(() => {
              const start = (currentPage.value - 1) * itemsPerPage.value
              const end = start + itemsPerPage.value
              return filteredOrders.value.slice(start, end)
            })

            // تابع تغییر صفحه
            const handlePageChange = (page) => {
              currentPage.value = page
            }

            // تابع تغییر تعداد آیتم در هر صفحه
            const handleItemsPerPageChange = (newItemsPerPage) => {
              itemsPerPage.value = newItemsPerPage
              currentPage.value = 1 // بازگشت به صفحه اول
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
