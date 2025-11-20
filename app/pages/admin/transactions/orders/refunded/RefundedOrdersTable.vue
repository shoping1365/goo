<template>
  <div>
    <!-- فیلترهای پیشرفته -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden mb-6">
      <div class="bg-gradient-to-r from-indigo-50 to-purple-50 px-4 py-3 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <div class="flex items-center">
            <div class="w-6 h-6 bg-indigo-500 rounded-md flex items-center justify-center ml-2">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.207A1 1 0 013 6.5V4z"></path>
              </svg>
            </div>
            <h3 class="text-sm font-semibold text-gray-900">فیلترهای پیشرفته</h3>
          </div>
          <button class="text-sm text-indigo-600 hover:text-indigo-800 transition-colors font-medium hover:bg-indigo-50 px-3 py-1 rounded-lg" @click="showFilters = !showFilters">
            {{ showFilters ? 'مخفی کردن' : 'نمایش' }}
          </button>
        </div>
      </div>
      
      <div v-if="showFilters" class="p-6">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <!-- فیلتر دلیل مسترد -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">دلیل مسترد</label>
            <select v-model="filters.refundReason" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500">
              <option value="">همه دلایل</option>
              <option value="customer">درخواست مشتری</option>
              <option value="inventory">عدم موجودی</option>
              <option value="payment">مشکل پرداخت</option>
              <option value="system">مشکل سیستمی</option>
            </select>
          </div>
          
          <!-- فیلتر روش مسترد -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">روش مسترد</label>
            <select v-model="filters.refundMethod" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500">
              <option value="">همه روش‌ها</option>
              <option value="bank_transfer">انتقال بانکی</option>
              <option value="wallet">کیف پول</option>
              <option value="gift_card">گیفت کارت</option>
              <option value="credit_card">کارت اعتباری</option>
              <option value="cash">نقدی</option>
              <option value="check">چک</option>
            </select>
          </div>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mt-4">
          <!-- فیلتر مبلغ -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مبلغ</label>
            <select v-model="filters.amountRange" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500">
              <option value="">همه مبالغ</option>
              <option value="0-500000">کمتر از 500 هزار تومان</option>
              <option value="500000-1000000">500 هزار تا 1 میلیون تومان</option>
              <option value="1000000-2000000">1 تا 2 میلیون تومان</option>
              <option value="2000000+">بیش از 2 میلیون تومان</option>
            </select>
          </div>
          
          <!-- فیلتر تاریخ مسترد -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ مسترد</label>
            <select v-model="filters.dateRange" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500">
              <option value="">همه تاریخ‌ها</option>
              <option value="today">امروز</option>
              <option value="yesterday">دیروز</option>
              <option value="week">هفته گذشته</option>
              <option value="month">ماه گذشته</option>
            </select>
          </div>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mt-4">
          <!-- فیلتر تاریخ مسترد از -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ مسترد از</label>
            <input v-model="filters.dateFrom" type="date" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500">
          </div>

          <!-- فیلتر تاریخ مسترد تا -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ مسترد تا</label>
            <input v-model="filters.dateTo" type="date" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500">
          </div>
        </div>

        <!-- دکمه‌های عملیات فیلتر -->
        <div class="flex items-center justify-between mt-4 pt-4 border-t border-gray-200">
          <div class="flex items-center space-x-2 space-x-reverse">
            <button
              class="inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
              @click="clearFilters"
            >
              پاک کردن فیلترها
            </button>
          </div>
          <div class="text-sm text-gray-500">
            {{ filteredOrders.length }} سفارش یافت شد
          </div>
        </div>
      </div>
    </div>
      
    <!-- جدول سفارشات -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
      <div class="bg-gradient-to-r from-slate-50 to-gray-50 px-4 py-3 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <div class="flex items-center">
            <div class="w-6 h-6 bg-slate-600 rounded-md flex items-center justify-center ml-2">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16"></path>
              </svg>
            </div>
            <h3 class="text-sm font-semibold text-gray-900">لیست سفارشات مسترد شده</h3>
          </div>
          
          <div class="relative">
            <input 
              v-model="searchTerm"
              type="text" 
              class="block w-64 pl-8 pr-3 py-2 border border-gray-300 rounded-md leading-5 bg-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 text-sm shadow-sm" 
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

      <!-- محتوای جدول -->
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">شماره سفارش</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مشتری</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مبلغ</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">دلیل مسترد</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">روش مسترد</th>
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
// Import کامپوننت مودال
import OrderDetailsModal from '~/components/admin/modals/OrderDetailsModal.vue'

// داده‌های واقعی از API
const orders = ref([])

// متغیرهای مدیریت وضعیت بارگذاری
const loading = ref(false)
const error = ref(null)

// تابع دریافت لیست سفارشات مسترد
const fetchOrders = async () => {
  try {
    loading.value = true
    error.value = null

    const { data } = await $fetch('/api/admin/orders/refunded', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    })

    orders.value = (data || []).map(order => ({
      id: order.id || 0,
      orderNumber: order.orderNumber || 'نامشخص',
      customerName: order.customerName || 'نامشخص',
      customerPhone: order.customerPhone || 'نامشخص',
      totalAmount: order.totalAmount || 0,
      refundReason: 'نامشخص', // TODO: اضافه کردن فیلد refundReason به مدل Order
      refundMethod: 'bank_transfer', // TODO: اضافه کردن فیلد refundMethod به مدل Order
      refundedAt: order.createdAt ? new Date(order.createdAt).toLocaleDateString('fa-IR') : 'نامشخص'
    }))

  } catch (err) {
    console.error('خطا در دریافت سفارشات:', err)
    error.value = err.message || 'خطا در دریافت سفارشات'
  } finally {
    loading.value = false
  }
}

// مقداردهی اولیه و دریافت داده‌ها
onMounted(() => {
  fetchOrders()
})

// متغیرهای فیلتر و جستجو
const showFilters = ref(false)
const searchTerm = ref('')
const filters = ref({
  refundReason: '',
  refundMethod: '',
  amountRange: '',
  dateRange: ''
})

// متغیرهای صفحه‌بندی
const currentPage = ref(1)
const itemsPerPage = ref(10)

// متغیرهای مودال
const isModalOpen = ref(false)
const selectedOrder = ref(null)

// فیلتر کردن سفارشات
const filteredOrders = computed(() => {
  let filtered = orders.value

  // فیلتر جستجو
  if (searchTerm.value) {
    const search = searchTerm.value.toLowerCase()
    filtered = filtered.filter(order => 
      order.orderNumber.toLowerCase().includes(search) ||
      order.customerName.toLowerCase().includes(search)
    )
  }

  // فیلتر دلیل مسترد
  if (filters.value.refundReason) {
    filtered = filtered.filter(order => order.refundReason === filters.value.refundReason)
  }

  // فیلتر روش مسترد
  if (filters.value.refundMethod) {
    filtered = filtered.filter(order => order.refundMethod === filters.value.refundMethod)
  }

  // فیلتر مبلغ
  if (filters.value.amountRange) {
    const [min, max] = filters.value.amountRange.split('-')
    if (max === '+') {
      // بیش از مقدار مشخص
      filtered = filtered.filter(order => order.totalAmount >= parseInt(min))
    } else {
      // بین دو مقدار
      filtered = filtered.filter(order => 
        order.totalAmount >= parseInt(min) && order.totalAmount <= parseInt(max)
      )
    }
  }

  // فیلتر تاریخ مسترد
  if (filters.value.dateRange) {
    const today = new Date()
    const orderDate = new Date()
    
    switch (filters.value.dateRange) {
      case 'today':
        filtered = filtered.filter(order => {
          orderDate.setTime(new Date(order.refundedAt).getTime())
          return orderDate.toDateString() === today.toDateString()
        })
        break
      case 'yesterday':
        const yesterday = new Date(today)
        yesterday.setDate(yesterday.getDate() - 1)
        filtered = filtered.filter(order => {
          orderDate.setTime(new Date(order.refundedAt).getTime())
          return orderDate.toDateString() === yesterday.toDateString()
        })
        break
      case 'week':
        const weekAgo = new Date(today)
        weekAgo.setDate(weekAgo.getDate() - 7)
        filtered = filtered.filter(order => new Date(order.refundedAt) >= weekAgo)
        break
      case 'month':
        const monthAgo = new Date(today)
        monthAgo.setMonth(monthAgo.getMonth() - 1)
        filtered = filtered.filter(order => new Date(order.refundedAt) >= monthAgo)
        break
    }
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
const clearFilters = () => {
  filters.value = {
    refundReason: '',
    refundMethod: '',
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

// متدهای صفحه‌بندی
const handlePageChange = (page) => {
  currentPage.value = page
}

const handleItemsPerPageChange = (newItemsPerPage) => {
  itemsPerPage.value = newItemsPerPage
  currentPage.value = 1 // بازگشت به صفحه اول
}

// متدهای کمکی
const formatDate = (date) => {
  return new Date(date).toLocaleDateString('fa-IR')
}

const formatPrice = (price) => {
  return new Intl.NumberFormat('fa-IR').format(price) + ' تومان'
}

// تابع دریافت متن وضعیت
const _getStatusText = (status) => {
  const texts = {
    approved: 'تأیید شده',
    processing: 'در حال پردازش',
    completed: 'مسترد شده',
    rejected: 'رد شده'
  }
  return texts[status] || status
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
</script> 
