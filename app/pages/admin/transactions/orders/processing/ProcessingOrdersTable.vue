<template>
  <div>
    <!-- فیلترهای پیشرفته -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden mb-6">
      <div class="bg-gradient-to-r from-blue-50 to-indigo-50 px-4 py-3 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <div class="flex items-center">
            <div class="w-6 h-6 bg-blue-500 rounded-md flex items-center justify-center ml-2">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.207A1 1 0 013 6.5V4z"></path>
              </svg>
            </div>
            <h3 class="text-sm font-semibold text-gray-900">فیلترهای پیشرفته</h3>
          </div>
          <button class="text-sm text-blue-600 hover:text-blue-800 transition-colors font-medium hover:bg-blue-50 px-3 py-1 rounded-lg" @click="showFilters = !showFilters">
            {{ showFilters ? 'مخفی کردن' : 'نمایش' }}
          </button>
        </div>
      </div>
      
      <div v-if="showFilters" class="p-6">
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
          <!-- فیلتر مرحله پردازش -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت سفارش</label>
            <select v-model="filters.stage" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
              <option value="">همه وضعیت‌ها</option>
              <option value="processing">در حال انجام</option>
            </select>
          </div>

          <!-- فیلتر پیشرفت -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">میزان پیشرفت</label>
            <select v-model="filters.progress" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
              <option value="">همه سطوح</option>
              <option value="0-25">0 تا 25%</option>
              <option value="25-50">25 تا 50%</option>
              <option value="50-75">50 تا 75%</option>
              <option value="75-100">75 تا 100%</option>
            </select>
          </div>

          <!-- فیلتر محدوده مبلغ -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">محدوده مبلغ</label>
            <select v-model="filters.amountRange" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
              <option value="">همه مبالغ</option>
              <option value="0-500000">کمتر از 500 هزار تومان</option>
              <option value="500000-1000000">500 هزار تا 1 میلیون تومان</option>
              <option value="1000000-5000000">1 تا 5 میلیون تومان</option>
              <option value="5000000+">بیش از 5 میلیون تومان</option>
            </select>
          </div>

          <!-- فیلتر بازه زمانی -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">بازه زمانی شروع</label>
            <select v-model="filters.dateRange" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
              <option value="">همه زمان‌ها</option>
              <option value="today">امروز</option>
              <option value="week">هفته جاری</option>
              <option value="month">ماه جاری</option>
              <option value="custom">انتخاب دستی</option>
            </select>
          </div>
        </div>

        <div class="mt-4 flex justify-end">
          <button class="px-4 py-2 bg-gray-100 text-gray-700 rounded-md hover:bg-gray-200 transition-colors" @click="clearFilters">
            پاک کردن فیلترها
          </button>
        </div>
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
            <h3 class="text-sm font-semibold text-gray-900">لیست سفارشات در حال انجام</h3>
          </div>
          
          <div class="flex items-center space-x-2 space-x-reverse">
            <!-- Bulk Actions -->
            <div v-if="selectedOrders.length > 0" class="flex items-center space-x-2 space-x-reverse bg-blue-50 rounded-md px-2 py-1.5 border border-blue-200">
              <span class="text-xs text-blue-700 font-medium">{{ selectedOrders.length }} مورد انتخاب شده</span>
              <select 
                v-model="bulkAction"
                class="text-xs border border-blue-300 rounded px-2 py-1 bg-blue-50"
                @change="executeBulkAction"
              >
                <option value="">عملیات گروهی</option>
                <option value="move_to_packaging">انتقال به بسته‌بندی</option>
                <option value="move_to_ready">آماده ارسال</option>
                <option value="pause">متوقف کردن</option>
                <option value="priority">اولویت بالا</option>
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
              <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                <svg class="h-4 w-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
                </svg>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Loading State -->
      <div v-if="isLoading" class="flex justify-center items-center py-8">
        <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-500"></div>
        <span class="mr-2 text-gray-600">در حال بارگذاری...</span>
      </div>

      <!-- Error State -->
      <div v-else-if="error" class="text-center py-8">
        <div class="text-red-600 mb-4">{{ error }}</div>
        <button class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600" @click="fetchProcessingOrders">
          تلاش مجدد
        </button>
      </div>

      <!-- Orders Table -->
      <div v-else class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                <input 
                  v-model="selectAll" 
                  type="checkbox"
                  class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                  @change="toggleSelectAll"
                />
              </th>
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
            <tr v-if="paginatedOrders.length === 0">
              <td colspan="8" class="px-6 py-8 text-center text-gray-500">
                هیچ سفارشی در حال انجام یافت نشد
              </td>
            </tr>
            <tr v-for="order in paginatedOrders" :key="order.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <input 
                  v-model="selectedOrders" 
                  type="checkbox"
                  :value="order.id"
                  class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                />
              </td>
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
                {{ formatPersianDateTime(order.createdAt) }}
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
                  <button class="text-orange-600 hover:text-orange-900" @click="pauseOrder(order)">متوقف</button>
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

<script setup>
// Import کامپوننت‌ها
import Pagination from '~/components/admin/common/Pagination.vue'
import OrderDetailsModal from '~/components/admin/modals/OrderDetailsModal.vue'

// متغیرهای فیلتر
const showFilters = ref(false)
const filters = ref({
  stage: '',
  progress: '',
  amountRange: '',
  dateRange: ''
})

// متغیرهای جستجو و انتخاب
const searchTerm = ref('')
const selectedOrders = ref([])
const selectAll = ref(false)
const bulkAction = ref('')

// متغیرهای صفحه‌بندی
const currentPage = ref(1)
const itemsPerPage = ref(10)

// متغیرهای مودال
const isModalOpen = ref(false)
const selectedOrder = ref(null)

// داده‌های سفارشات در حال انجام
const processingOrders = ref([])
const isLoading = ref(false)
const error = ref(null)

// تابع دریافت داده‌ها از API
const fetchProcessingOrders = async () => {
  isLoading.value = true
  error.value = null
  
  try {
    const response = await $fetch('/api/admin/orders/processing', {
      query: {
        page: currentPage.value,
        pageSize: itemsPerPage.value
      }
    })
    
    // بررسی ساختار response و استخراج داده‌ها
    let orders = []
    
    if (response && typeof response === 'object') {
      if (response.data && Array.isArray(response.data)) {
        orders = response.data
      } else if (Array.isArray(response)) {
        orders = response
      } else if (response.orders && Array.isArray(response.orders)) {
        orders = response.orders
      }
    }
    
    // تبدیل داده‌ها به فرمت مورد نیاز
    processingOrders.value = orders.map(order => ({
      id: order.id,
      orderNumber: order.orderNumber || order.trackingCode || `#${order.id}`,
      customerName: order.customerName || 'نامشخص',
      customerPhone: order.customerPhone || 'شماره نامشخص',
      totalAmount: order.totalAmount || order.finalAmount || 0,
      status: order.status || 'processing',
      paymentMethod: order.paymentMethod || 'online',
      createdAt: order.createdAt,
      itemsCount: order.itemsCount || 0,
      shippingAddress: order.shippingAddress || '',
      shippingCity: order.shippingCity || '',
      shippingProvince: order.shippingProvince || '',
      shippingPostalCode: order.shippingPostalCode || '',
      recipientName: order.recipientName || '',
      recipientPhone: order.recipientPhone || ''
    }))
    
  } catch (err) {
    error.value = 'خطا در دریافت داده‌های سفارشات'
    console.error('Error fetching processing orders:', err)
  } finally {
    isLoading.value = false
  }
}

// دریافت داده‌ها هنگام بارگذاری کامپوننت
onMounted(() => {
  fetchProcessingOrders()
})

// نظارت بر تغییرات صفحه و بروزرسانی داده‌ها
watch([currentPage, itemsPerPage], () => {
  fetchProcessingOrders()
})

// محاسبات فیلتر و جستجو
const filteredOrders = computed(() => {
  let filtered = processingOrders.value

  // فیلتر بر اساس جستجو
  if (searchTerm.value) {
    filtered = filtered.filter(order => 
      order.orderNumber.toLowerCase().includes(searchTerm.value.toLowerCase()) ||
      order.customerName.toLowerCase().includes(searchTerm.value.toLowerCase()) ||
      order.customerPhone.includes(searchTerm.value)
    )
  }

  // فیلتر بر اساس مرحله
  if (filters.value.stage) {
    filtered = filtered.filter(order => order.stage === filters.value.stage)
  }

  // فیلتر بر اساس پیشرفت
  if (filters.value.progress) {
    const [min, max] = filters.value.progress.split('-').map(Number)
    filtered = filtered.filter(order => {
      if (max) {
        return order.progress >= min && order.progress <= max
      } else {
        return order.progress >= min
      }
    })
  }

  // فیلتر بر اساس محدوده مبلغ
  if (filters.value.amountRange) {
    const [min, max] = filters.value.amountRange.split('-').map(Number)
    filtered = filtered.filter(order => {
      if (max) {
        return order.totalAmount >= min && order.totalAmount <= max
      } else {
        return order.totalAmount >= min
      }
    })
  }

  return filtered
})

// محاسبات صفحه‌بندی
const totalPages = computed(() => Math.ceil(filteredOrders.value.length / itemsPerPage.value))

const paginatedOrders = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage.value
  const end = start + itemsPerPage.value
  return filteredOrders.value.slice(start, end)
})

// متدهای عملیاتی
const clearFilters = () => {
  filters.value = {
    stage: '',
    progress: '',
    amountRange: '',
    dateRange: ''
  }
  searchTerm.value = ''
  currentPage.value = 1
}

// نظارت بر تغییرات فیلترها و بازگشت به صفحه اول
watch([filters, searchTerm], () => {
  currentPage.value = 1
}, { deep: true })

// متدهای انتخاب
const toggleSelectAll = () => {
  if (selectAll.value) {
    selectedOrders.value = paginatedOrders.value.map(order => order.id)
  } else {
    selectedOrders.value = []
  }
}

// نظارت بر تغییرات انتخاب
watch(selectedOrders, (newSelection) => {
  selectAll.value = newSelection.length === paginatedOrders.value.length && paginatedOrders.value.length > 0
})

const executeBulkAction = () => {
  if (!bulkAction.value) return
  
  // اینجا می‌توانید منطق عملیات گروهی را پیاده‌سازی کنید
  bulkAction.value = ''
  selectedOrders.value = []
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

// متدهای عملیات
const updateProgress = (_order) => {
  // Update progress logic
}

const pauseOrder = (_order) => {
  // Pause order logic
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
  return new Intl.NumberFormat('fa-IR').format(price) + ' تومان'
}

const getStatusClass = (status) => {
  const classes = {
    processing: 'bg-blue-100 text-blue-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status) => {
  const texts = {
    processing: 'در حال انجام'
  }
  return texts[status] || status
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
</script> 
