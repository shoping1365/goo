<template>
  <div class="space-y-6">
    <!-- آمار کلی -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gapx-4 py-4 mb-6">
      <!-- کل سفارشات ارسال شده -->
      <div class="bg-gradient-to-br from-green-50 to-green-100 border border-green-200 rounded-lg px-4 py-4 hover:shadow-md transition-all duration-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-green-700 mb-1">کل سفارشات ارسال شده</p>
            <p class="text-2xl font-bold text-green-900">{{ stats.totalShipped?.toLocaleString('fa-IR') || '0' }}</p>
            <p class="text-xs text-green-600 mt-1">در 30 روز گذشته</p>
          </div>
          <div class="w-12 h-12 bg-green-500 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 8h14M5 8a2 2 0 110-4h14a2 2 0 110 4M5 8v10a2 2 0 002 2h10a2 2 0 002-2V8m-14 0h14"></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- مبلغ کل ارسال شده -->
      <div class="bg-gradient-to-br from-blue-50 to-blue-100 border border-blue-200 rounded-lg px-4 py-4 hover:shadow-md transition-all duration-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-blue-700 mb-1">مبلغ کل ارسال شده</p>
            <p class="text-2xl font-bold text-blue-900">{{ formatPrice(stats.totalAmount || 0) }}</p>
            <p class="text-xs text-blue-600 mt-1">در 30 روز گذشته</p>
          </div>
          <div class="w-12 h-12 bg-blue-500 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- نرخ تحویل موفق -->
      <div class="bg-gradient-to-br from-purple-50 to-purple-100 border border-purple-200 rounded-lg px-4 py-4 hover:shadow-md transition-all duration-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-purple-700 mb-1">نرخ تحویل موفق</p>
            <p class="text-2xl font-bold text-purple-900">{{ stats.deliverySuccessRate || 0 }}%</p>
            <p class="text-xs text-purple-600 mt-1">سفارشات تحویل شده</p>
          </div>
          <div class="w-12 h-12 bg-purple-500 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- متوسط زمان تحویل -->
      <div class="bg-gradient-to-br from-orange-50 to-orange-100 border border-orange-200 rounded-lg px-4 py-4 hover:shadow-md transition-all duration-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-orange-700 mb-1">متوسط زمان تحویل</p>
            <p class="text-2xl font-bold text-orange-900">{{ stats.avgDeliveryTime || 0 }} روز</p>
            <p class="text-xs text-orange-600 mt-1">از ارسال تا تحویل</p>
          </div>
          <div class="w-12 h-12 bg-orange-500 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
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
      <!-- توزیع وضعیت ارسال -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
        <div class="bg-gradient-to-r from-green-50 to-emerald-50 px-4 py-3 border-b border-gray-200">
          <div class="flex items-center">
            <div class="w-6 h-6 bg-green-500 rounded-md flex items-center justify-center ml-2">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
              </svg>
            </div>
            <h3 class="text-sm font-semibold text-gray-900">توزیع وضعیت ارسال</h3>
          </div>
        </div>

        <div class="px-4 py-4">
          <div class="space-y-4">
            <div
              v-for="status in shippingStatuses"
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
                    class="h-2 rounded-full"
                    :style="{ width: `${status.percentage}%`, backgroundColor: status.color }"
                  ></div>
                </div>
                <span class="text-sm font-medium text-gray-900 w-12 text-left">{{ status.percentage }}%</span>
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
                <div class="w-3 h-3 rounded-full ml-3" :style="{ backgroundColor: method.color || '#3B82F6' }"></div>
                <span class="text-sm text-gray-700">{{ methodName }}</span>
              </div>
              <div class="flex items-center space-x-3 space-x-reverse">
                <div class="w-32 bg-gray-200 rounded-full h-2">
                  <div
                    class="h-2 rounded-full"
                    :style="{ width: `${method.percentage || 0}%`, backgroundColor: method.color || '#3B82F6' }"
                  ></div>
                </div>
                <div class="text-left">
                  <div class="text-sm font-medium text-gray-900">{{ method.count || 0 }}</div>
                  <div class="text-xs text-gray-500">{{ method.percentage || 0 }}%</div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- آمار زمانی ارسال -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
        <div class="bg-gradient-to-r from-blue-50 to-indigo-50 px-4 py-3 border-b border-gray-200">
          <div class="flex items-center">
            <div class="w-6 h-6 bg-blue-500 rounded-md flex items-center justify-center ml-2">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
            <h3 class="text-sm font-semibold text-gray-900">آمار زمانی ارسال</h3>
          </div>
        </div>

        <div class="px-4 py-4">
          <div class="grid grid-cols-2 gapx-4 py-4">
            <div class="text-center">
              <div class="text-2xl font-bold text-green-600">{{ timeStats.avgDeliveryTime }}</div>
              <div class="text-xs text-gray-600 mt-1">متوسط زمان تحویل</div>
            </div>
            <div class="text-center">
              <div class="text-2xl font-bold text-blue-600">{{ timeStats.fastestDelivery }}</div>
              <div class="text-xs text-gray-600 mt-1">سریع ترین تحویل</div>
            </div>
            <div class="text-center">
              <div class="text-2xl font-bold text-purple-600">{{ timeStats.shippingPerDay }}</div>
              <div class="text-xs text-gray-600 mt-1">ارسال در روز</div>
            </div>
            <div class="text-center">
              <div class="text-2xl font-bold text-orange-600">{{ timeStats.longestDelivery }}</div>
              <div class="text-xs text-gray-600 mt-1">طولانی ترین تحویل</div>
            </div>
          </div>
        </div>
      </div>
    </div>



    <!-- نمودارها -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gapx-4 py-4 mb-6">
      <!-- نمودار روند ارسال (7 روز) -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
        <div class="bg-gradient-to-r from-green-50 to-emerald-50 px-4 py-3 border-b border-gray-200">
          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-6 h-6 bg-green-500 rounded-md flex items-center justify-center ml-2">
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 12l3-3 3 3 4-4M8 21l4-4 4 4M3 4h18M4 4h16v12a1 1 0 01-1 1H5a1 1 0 01-1-1V4z"></path>
                </svg>
              </div>
              <h3 class="text-sm font-semibold text-gray-900">روند ارسال در ۷ روز گذشته</h3>
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <button
                v-for="period in chartPeriods"
                :key="period.id"
                :class="[
                  'px-3 py-1 text-xs font-medium rounded-md transition-colors',
                  selectedPeriod === period.id
                    ? 'bg-green-500 text-white'
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
                class="w-full bg-gradient-to-t from-green-500 to-green-400 rounded-t"
                :style="{ height: `${(day.value / maxChartValue) * 200}px` }"
              ></div>
              <div class="text-xs font-medium text-gray-700 mt-1">{{ day.value }}</div>
            </div>
          </div>
        </div>
      </div>

      <!-- نمودار دایره‌ای روش‌های ارسال -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
        <div class="bg-gradient-to-r from-purple-50 to-violet-50 px-4 py-3 border-b border-gray-200">
          <div class="flex items-center">
            <div class="w-6 h-6 bg-purple-500 rounded-md flex items-center justify-center ml-2">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
              </svg>
            </div>
            <h3 class="text-sm font-semibold text-gray-900">توزیع روش‌های ارسال</h3>
          </div>
        </div>

        <div class="px-4 py-4">
          <div class="flex items-center justify-center">
            <!-- نمودار دایره‌ای -->
            <div class="relative w-48 h-48">
              <svg class="w-full h-full transform -rotate-90" viewBox="0 0 100 100">
                <circle
                  cx="50"
                  cy="50"
                  r="40"
                  fill="none"
                  stroke="#e5e7eb"
                  stroke-width="8"
                />
                <circle
                  v-for="(method, methodName, index) in stats.shippingMethods"
                  :key="methodName"
                  cx="50"
                  cy="50"
                  r="40"
                  fill="none"
                  :stroke="method.color || '#3B82F6'"
                  stroke-width="8"
                  :stroke-dasharray="`${(method.percentage || 0) * 2.51} ${(100 - (method.percentage || 0)) * 2.51}`"
                  :stroke-dashoffset="getCircleOffset(index)"
                  class="transition-all duration-500"
                />
              </svg>
              
              <!-- مرکز نمودار -->
              <div class="absolute inset-0 flex items-center justify-center">
                <div class="text-center">
                  <div class="text-2xl font-bold text-gray-900">{{ stats.totalShipped || 0 }}</div>
                  <div class="text-xs text-gray-500">کل سفارشات</div>
                </div>
              </div>
            </div>
          </div>

          <!-- راهنمای نمودار -->
          <div class="mt-6 grid grid-cols-2 gap-3">
            <div
              v-for="(method, methodName) in stats.shippingMethods"
              :key="methodName"
              class="flex items-center space-x-2 space-x-reverse"
            >
                <div class="w-3 h-3 rounded-full" :style="{ backgroundColor: method.color || '#3B82F6' }"></div>
              <div class="flex-1">
                <div class="text-sm font-medium text-gray-900">{{ methodName }}</div>
                <div class="text-xs text-gray-500">{{ method.count || 0 }} سفارش ({{ method.percentage || 0 }}%)</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- سفارشات اخیر ارسال شده -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
      <div class="bg-gradient-to-r from-slate-50 to-gray-50 px-4 py-3 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <div class="flex items-center">
            <div class="w-6 h-6 bg-slate-600 rounded-md flex items-center justify-center ml-2">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16"></path>
              </svg>
            </div>
            <h3 class="text-sm font-semibold text-gray-900">سفارشات اخیر ارسال شده</h3>
          </div>
          
          <div class="flex items-center space-x-2 space-x-reverse">
            <div class="relative">
              <input
                v-model="searchTerm"
                type="text"
                placeholder="جستجو در سفارشات..."
                class="w-48 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-green-500 text-sm"
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
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">روش ارسال</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ ارسال</th>
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
                  <div class="text-sm font-medium text-gray-900">{{ order.customerName }}</div>
                  <div class="text-sm text-gray-500">{{ order.customerPhone }}</div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ formatPrice(order.totalAmount || 0) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusClass(order.status)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ getStatusText(order.status) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ order.shippingMethod }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ formatDate(order.shippedAt) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <div class="flex items-center space-x-2">
                  <button class="text-blue-600 hover:text-blue-900" @click="viewOrderDetails(order)">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                    </svg>
                  </button>
                  <button class="text-green-600 hover:text-green-900" @click="trackOrder(order)">پیگیری</button>
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
        :total="filteredRecentOrders.length"
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

// Props
const props = defineProps({
  stats: {
    type: Object,
    default: () => ({
      totalShipped: 0,
      delivered: 0,
      inTransit: 0,
      totalAmount: 0,
      deliverySuccessRate: 0,
      avgDeliveryTime: 0
    })
  },
  orders: {
    type: Array,
    default: () => []
  }
})

// دوره‌های نمودار
const chartPeriods = ref([
  { id: '7d', name: '7 روز' },
  { id: '30d', name: '30 روز' },
  { id: '90d', name: '90 روز' }
])

const selectedPeriod = ref('7d')

// داده‌های نمودار 7 روزه
const chartData = ref([
  { date: 'شنبه', value: 12 },
  { date: 'یکشنبه', value: 18 },
  { date: 'دوشنبه', value: 15 },
  { date: 'سه‌شنبه', value: 22 },
  { date: 'چهارشنبه', value: 19 },
  { date: 'پنج‌شنبه', value: 25 },
  { date: 'جمعه', value: 16 }
])



// وضعیت‌های ارسال
const shippingStatuses = ref([
  { id: 1, name: 'تحویل شده', percentage: 65, color: '#10B981' },
  { id: 2, name: 'در راه', percentage: 25, color: '#F59E0B' },
  { id: 3, name: 'در انتظار ارسال', percentage: 10, color: '#6B7280' }
])

// آمار زمانی
const timeStats = ref({
  avgDeliveryTime: '3.2 روز',
  fastestDelivery: '1 روز',
  shippingPerDay: '45 سفارش',
  longestDelivery: '7 روز'
})

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



const filteredRecentOrders = computed(() => {
  if (!searchTerm.value) return props.orders
  
  return props.orders.filter(order => 
    order.orderNumber.toLowerCase().includes(searchTerm.value.toLowerCase()) ||
    order.customerName.toLowerCase().includes(searchTerm.value.toLowerCase())
  )
})

const totalPages = computed(() => Math.ceil(filteredRecentOrders.value.length / itemsPerPage.value))

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
const trackOrder = (_order) => {
  // اینجا می‌توانید کاربر را به صفحه پیگیری هدایت کنید
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

const formatDate = (date) => {
  return new Date(date).toLocaleDateString('fa-IR')
}

const getStatusClass = (status) => {
  const classes = {
    shipped: 'bg-green-100 text-green-800',
    delivered: 'bg-blue-100 text-blue-800',
    in_transit: 'bg-yellow-100 text-yellow-800',
    pending: 'bg-gray-100 text-gray-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status) => {
  const texts = {
    shipped: 'ارسال شده',
    delivered: 'تحویل شده',
    in_transit: 'در راه',
    pending: 'در انتظار ارسال'
  }
  return texts[status] || status
}

// تابع محاسبه محبوب‌ترین روش ارسال
const getMostPopularMethod = () => {
  if (!props.stats.shippingMethods) {
    return { name: 'نامشخص', count: 0 }
  }

  const methods = Object.entries(props.stats.shippingMethods)
  const mostPopular = methods.reduce((max, [name, data]) => {
    return (data?.count || 0) > max.count ? { name, count: data.count || 0 } : max
  }, { name: '', count: 0 })

  return mostPopular
}

// تابع محاسبه offset نمودار دایره‌ای
const getCircleOffset = (index) => {
  if (!props.stats.shippingMethods) return 0

  const methods = Object.values(props.stats.shippingMethods)
  let offset = 0

  for (let i = 0; i < index; i++) {
    offset += (methods[i]?.percentage || 0) * 2.51
  }

  return offset
}
</script> 
