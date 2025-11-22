<template>
  <ClientOnly>
    <!-- Page Header -->
    <div class="bg-white shadow-sm border-b border-gray-200 mb-4">
      <div class="w-full px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center py-3">
          <div>
            <h1 class="text-lg font-bold text-gray-900">سفارشات مرجوع شده</h1>
            <p class="text-xs text-gray-600 mt-1">مدیریت و پیگیری سفارشات مرجوع شده</p>
          </div>
          <div class="flex space-x-2 space-x-reverse">
            <button
              class="inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-lg text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 shadow-md transition-all duration-200"
              @click="refreshData"
            >
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
              </svg>
              بروزرسانی
            </button>
          </div>
        </div>
      </div>
    </div>

    <div class="w-full px-4 py-4">
      <!-- تب‌های اصلی -->
      <div class="bg-white border-b border-gray-200 mb-6">
        <nav class="-mb-px flex px-6 overflow-x-auto">
          <button
            v-for="tab in tabs"
            :key="tab.id"
            :class="[
              activeTab === tab.id
                ? 'border-red-500 text-red-600'
                : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300',
            'flex-1 whitespace-nowrap py-4 px-4 border-b-2 font-medium text-sm text-center'
            ]"
            @click="activeTab = tab.id"
          >
            {{ tab.name }}
          </button>
        </nav>
      </div>

      <div class="px-4 py-4">
        <!-- تب داشبورد -->
        <div v-if="activeTab === 'dashboard'">
          <ReturnedDashboard 
            :stats="stats"
            :orders="orders"
          />
        </div>

        <!-- تب سفارشات -->
        <div v-if="activeTab === 'returned-orders'">
          <ReturnedOrdersTable 
            :orders="orders"
          />
        </div>

        <!-- تب گزارشات -->
        <div v-if="activeTab === 'reports'">
          <ReturnedReports />
        </div>
      </div>
    </div>
  </ClientOnly>
</template>

<script setup>
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// Import کامپوننت‌ها
import ReturnedDashboard from './ReturnedDashboard.vue'
import ReturnedOrdersTable from './ReturnedOrdersTable.vue'
import ReturnedReports from './ReturnedReports.vue'

// تب‌های صفحه
const tabs = ref([
  { id: 'dashboard', name: 'داشبورد' },
  { id: 'returned-orders', name: 'سفارشات مرجوع شده' },
  { id: 'reports', name: 'گزارشات' }
])

const activeTab = ref('dashboard')

// داده‌های واقعی از API
const stats = ref({})
const orders = ref([])

// متغیرهای مدیریت وضعیت بارگذاری
const loading = ref(false)
const error = ref(null)

// تابع دریافت داده‌های آمار
const fetchStats = async () => {
  try {
    loading.value = true
    error.value = null

    const { data } = await $fetch('/api/admin/orders/returned/stats', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    })

    // تبدیل داده‌های API به فرمت مورد نیاز کامپوننت
    stats.value = {
      totalReturned: data.totalReturned || 0,
      pendingReview: data.pendingReview || 0,
      approved: data.approved || 0,
      shippingMethods: {}
    }

    // تبدیل shippingMethods به فرمت مورد نیاز
    if (data.shippingMethods && Array.isArray(data.shippingMethods)) {
      data.shippingMethods.forEach(method => {
        if (method && method.shipping_method) {
          stats.value.shippingMethods[method.shipping_method] = {
            count: method.count || 0,
            percentage: method.percentage || 0,
            color: '#3B82F6' // رنگ پیش‌فرض
          }
        }
      })
    }

    // تنظیم رنگ‌های مختلف برای شرکت‌های حمل و نقل
    const colors = ['#3B82F6', '#10B981', '#F59E0B', '#8B5CF6', '#6B7280']
    let colorIndex = 0
    Object.keys(stats.value.shippingMethods || {}).forEach(method => {
      if (stats.value.shippingMethods && stats.value.shippingMethods[method] && !stats.value.shippingMethods[method].color) {
        stats.value.shippingMethods[method].color = colors[colorIndex % colors.length]
        colorIndex++
      }
    })

  } catch (err) {
    console.error('خطا در دریافت آمار:', err)
    error.value = err.message || 'خطا در دریافت آمار'
  }
}

// تابع دریافت لیست سفارشات
const fetchOrders = async () => {
  try {
    loading.value = true
    error.value = null

    const { data } = await $fetch('/api/admin/orders/returned', {
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
      returnedAt: order.createdAt ? new Date(order.createdAt).toLocaleDateString('fa-IR') : 'نامشخص',
      reason: 'نامشخص', // TODO: اضافه کردن فیلد reason به مدل Order
      status: order.status || 'نامشخص'
    }))

  } catch (err) {
    console.error('خطا در دریافت سفارشات:', err)
    error.value = err.message || 'خطا در دریافت سفارشات'
  } finally {
    loading.value = false
  }
}

// تابع بروزرسانی داده‌ها
const refreshData = async () => {
  await Promise.all([fetchStats(), fetchOrders()])
}

// تنظیم عنوان صفحه
useHead({
  title: 'سفارشات مرجوع شده - پنل مدیریت',
  meta: [
    { name: 'description', content: 'مدیریت سفارشات مرجوع شده و گزارشات مربوطه' }
  ]
})
</script> 
