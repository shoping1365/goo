<template>
  <div>
    <!-- Page Header -->
    <div class="bg-white shadow-sm border-b border-gray-200 mb-4 rounded-lg">
      <div class="w-full px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center py-3">
        <div>
            <h1 class="text-lg font-bold text-gray-900">در حال انجام</h1>
            <p class="text-xs text-gray-600 mt-1">مدیریت و پیگیری سفارشات در حال انجام</p>
        </div>
          <div class="flex space-x-2 space-x-reverse">
            <button 
              @click="showCreateOrderModal = true"
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-blue-400 to-blue-600 hover:from-blue-500 hover:to-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105"
            >
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5 ml-2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M12 4v16m8-8H4" />
              </svg>
              سفارش جدید
            </button>
            <button 
              @click="exportOrders"
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-green-400 to-green-600 hover:from-green-500 hover:to-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105"
            >
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
              </svg>
              خروجی Excel
            </button>
            <button 
              @click="printOrders"
              class="inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-lg text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 shadow-md transition-all duration-200"
            >
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 17h2a2 2 0 002-2v-4a2 2 0 00-2-2H5a2 2 0 00-2 2v4a2 2 0 002 2h2m2 4h6a2 2 0 002-2v-4a2 2 0 00-2-2H9a2 2 0 00-2 2v4a2 2 0 002 2zm8-12V5a2 2 0 00-2-2H9a2 2 0 00-2 2v4h10z"></path>
              </svg>
              چاپ
          </button>
          </div>
        </div>
      </div>
    </div>

    <!-- آمار سفارشات در حال انجام -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden mb-6">
      <div class="bg-gradient-to-r from-purple-50 to-indigo-50 px-4 py-3 border-b border-gray-200">
        <div class="flex items-center">
          <div class="w-6 h-6 bg-purple-500 rounded-md flex items-center justify-center ml-2">
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
            </svg>
          </div>
          <h3 class="text-sm font-semibold text-gray-900">در حال انجام</h3>
        </div>
      </div>

      <div class="px-4 py-4">
        <div class="grid grid-cols-1 md:grid-cols-3 gapx-4 py-4">
          <!-- سفارشات در حال انجام -->
          <div class="bg-gradient-to-br from-purple-50 to-purple-100 border border-purple-200 rounded-lg px-4 py-4 hover:shadow-md transition-all duration-200">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-purple-700 mb-1">سفارشات در حال انجام</p>
                <p class="text-2xl font-bold text-purple-900">{{ stats.totalProcessing.toLocaleString('fa-IR') }}</p>
          </div>
              <div class="w-12 h-12 bg-purple-500 rounded-lg flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
                </svg>
          </div>
        </div>
      </div>

          <!-- در حال آماده‌سازی -->
          <div class="bg-gradient-to-br from-green-50 to-green-100 border border-green-200 rounded-lg px-4 py-4 hover:shadow-md transition-all duration-200">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-green-700 mb-1">در حال آماده‌سازی</p>
                <p class="text-2xl font-bold text-green-900">{{ stats.preparing.toLocaleString('fa-IR') }}</p>
          </div>
              <div class="w-12 h-12 bg-green-500 rounded-lg flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  
                </svg>
          </div>
        </div>
      </div>

          <!-- آماده ارسال -->
          <div class="bg-gradient-to-br from-blue-50 to-blue-100 border border-blue-200 rounded-lg px-4 py-4 hover:shadow-md transition-all duration-200">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm font-medium text-blue-700 mb-1">آماده ارسال</p>
                <p class="text-2xl font-bold text-blue-900">{{ stats.readyToShip.toLocaleString('fa-IR') }}</p>
              </div>
              <div class="w-12 h-12 bg-blue-500 rounded-lg flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 8h14M5 8a2 2 0 110-4h14a2 2 0 110 4M5 8v10a2 2 0 002 2h10a2 2 0 002-2V8m-14 0h14"></path>
                </svg>
              </div>
          </div>
          </div>


        </div>
      </div>
    </div>

    <!-- فیلترهای پیشرفته -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden mb-6">
      <div class="bg-gradient-to-r from-green-50 to-emerald-50 px-4 py-3 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <div class="flex items-center">
            <div class="w-6 h-6 bg-green-500 rounded-md flex items-center justify-center ml-2">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.207A1 1 0 013 6.5V4z"></path>
              </svg>
            </div>
            <h3 class="text-sm font-semibold text-gray-900">فیلترهای پیشرفته در حال انجام</h3>
          </div>
          <button @click="showFilters = !showFilters" class="text-sm text-green-600 hover:text-green-800 transition-colors font-medium hover:bg-green-50 px-3 py-1 rounded-lg">
            {{ showFilters ? 'مخفی کردن' : 'نمایش' }}
          </button>
        </div>
      </div>

      <div v-if="showFilters" class="px-4 py-4">
        <OrderFilters @filter-change="handleFiltersChange" />
      </div>
      </div>

    <!-- لیست سفارشات -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
      <div class="bg-gradient-to-r from-slate-50 to-gray-50 px-4 py-3 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <div class="flex items-center">
            <div class="w-6 h-6 bg-slate-600 rounded-md flex items-center justify-center ml-2">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16"></path>
              </svg>
            </div>
            <h3 class="text-sm font-semibold text-gray-900">لیست در حال انجام</h3>
          </div>
          
          <div class="flex items-center space-x-2 space-x-reverse">
            <!-- Bulk Actions -->
            <div class="flex items-center space-x-2 space-x-reverse bg-blue-50 rounded-md px-2 py-1.5 border border-blue-200" v-if="selectedOrders.length > 0">
              <span class="text-xs text-blue-700 font-medium">{{ selectedOrders.length }} مورد انتخاب شده</span>
              <select 
                v-model="bulkAction"
                @change="executeBulkAction"
                class="text-xs border border-blue-300 rounded px-2 py-1 bg-blue-50"
              >
                <option value="">عملیات گروهی</option>
                <option value="confirm">تایید سفارشات</option>
                <option value="process">آماده‌سازی</option>
                <option value="ready">آماده ارسال</option>
                <option value="ship">ارسال</option>
                <option value="print">چاپ فاکتور</option>
                <option value="export">خروجی Excel</option>
              </select>
            </div>
            
            <!-- Search -->
            <div class="relative">
              <input 
                v-model="searchTerm"
                type="text" 
                class="block w-56 pl-8 pr-3 py-1.5 border border-gray-300 rounded-md leading-5 bg-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-xs shadow-sm" 
                placeholder="جستجو شماره سفارش، مشتری یا کد پیگیری"
                dir="rtl"
              />
              <div class="absolute inset-y-0 left-0 pl-2 flex items-center pointer-events-none">
                <svg class="h-4 w-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
                </svg>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Table Content -->
      <OrderTable 
        :orders="paginatedOrders"
        @view-details="openOrderDetail"
        @quick-actions="openQuickActions"
        @refresh="loadOrders"
        @export="exportOrders"
      />

      <!-- Pagination -->
      <Pagination 
        :current-page="currentPage"
        :total-pages="totalPages"
        :total="totalItems"
        :items-per-page="itemsPerPage"
        @page-changed="handlePageChange"
        @items-per-page-changed="handleItemsPerPageChange"
      />
    </div>
  </div>
</template>

<script setup>
definePageMeta({
  layout: 'admin-main'
})

// Import components
import OrderFilters from '../OrderFilters.vue'
import OrderTable from '../OrderTable.vue'
import Pagination from '~/components/common/Pagination.vue'

// داده‌های نمونه
const stats = ref({
  totalProcessing: 156,
  preparing: 67,
  readyToShip: 34
})

const orders = ref([
  {
    id: 1,
    orderNumber: 'ORD-101',
    customerName: 'علی احمدی',
    customerPhone: '09123456789',
    createdAt: '2024-01-15',
    totalAmount: 1250000,
    status: 'processing',
    priority: 'high',
    paymentMethod: 'online',
    items: [
      { name: 'لپ تاپ اپل', quantity: 1, price: 1250000 }
    ],
    totalItems: 1
  },
  {
    id: 2,
    orderNumber: 'ORD-102',
    customerName: 'فاطمه محمدی',
    customerPhone: '09187654321',
    createdAt: '2024-01-14',
    totalAmount: 890000,
    status: 'packaging',
    priority: 'medium',
    paymentMethod: 'cash',
    items: [
      { name: 'گوشی سامسونگ', quantity: 1, price: 890000 }
    ],
    totalItems: 1
  },
  {
    id: 3,
    orderNumber: 'ORD-103',
    customerName: 'محمد رضایی',
    customerPhone: '09111222333',
    createdAt: '2024-01-13',
    totalAmount: 2100000,
    status: 'preparing',
    priority: 'low',
    paymentMethod: 'bank',
    items: [
      { name: 'تبلت اپل', quantity: 1, price: 2100000 }
    ],
    totalItems: 1
  }
])

// متغیرهای مورد نیاز
const showCreateOrderModal = ref(false)
const showFilters = ref(false)
const selectedOrder = ref(null)
const selectedOrders = ref([])
const searchTerm = ref('')
const bulkAction = ref('')
const currentPage = ref(1)
const itemsPerPage = ref(10)

// محاسبات
const filteredOrders = computed(() => {
  let filtered = orders.value

  if (searchTerm.value) {
    const search = searchTerm.value.toLowerCase()
    filtered = filtered.filter(order => 
      order.orderNumber.toLowerCase().includes(search) ||
      order.customerName.toLowerCase().includes(search) ||
      order.customerPhone.includes(search)
    )
  }

  return filtered
})

const totalItems = computed(() => filteredOrders.value.length)
const totalPages = computed(() => Math.ceil(totalItems.value / itemsPerPage.value))

const paginatedOrders = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage.value
  const end = start + itemsPerPage.value
  return filteredOrders.value.slice(start, end)
})

// متدهای عملیاتی
const exportOrders = () => {
  console.log('خروجی Excel سفارشات...')
}

const printOrders = () => {
  console.log('چاپ سفارشات...')
}

const handleFiltersChange = (filters) => {
  console.log('فیلترها تغییر کردند:', filters)
  currentPage.value = 1
}

const executeBulkAction = () => {
  if (!bulkAction.value) return
  
  console.log('اجرای عملیات گروهی:', bulkAction.value, 'برای سفارشات:', selectedOrders.value)
  bulkAction.value = ''
  selectedOrders.value = []
}

const openOrderDetail = (order) => {
  selectedOrder.value = order
  console.log('نمایش جزئیات سفارش:', order)
}

const openQuickActions = (order) => {
  selectedOrder.value = order
  console.log('عملیات سریع برای سفارش:', order)
}

const handlePageChange = (page) => {
  currentPage.value = page
}

const handleItemsPerPageChange = (newItemsPerPage) => {
  itemsPerPage.value = newItemsPerPage
  currentPage.value = 1
}

const loadOrders = () => {
  console.log('بارگذاری مجدد سفارشات...')
}
</script> 
