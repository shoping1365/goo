<template>
  <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
    <!-- هدر جدول -->
    <div class="bg-gradient-to-r from-green-50 to-emerald-50 px-4 py-3 border-b border-gray-200">
      <div class="flex items-center justify-between">
        <div class="flex items-center">
          <div class="w-6 h-6 bg-green-500 rounded-md flex items-center justify-center ml-2">
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 8h14M5 8a2 2 0 110-4h14a2 2 0 110 4M5 8v10a2 2 0 002 2h10a2 2 0 002-2V8m-14 0h14"></path>
            </svg>
          </div>
          <h3 class="text-sm font-semibold text-gray-900">سفارشات ارسال شده</h3>
        </div>
        <div class="flex items-center space-x-2 space-x-reverse">
          <button @click="exportOrders" class="text-green-600 hover:text-green-800 text-sm font-medium">
            <svg class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
            </svg>
            خروجی
          </button>
        </div>
      </div>
    </div>



    <!-- جدول -->
    <div class="overflow-x-auto">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">شماره سفارش</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مشتری</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مبلغ</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ ارسال</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">شرکت ارسال</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">شماره پیگیری</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-for="order in paginatedOrders" :key="order.id" class="hover:bg-gray-50">
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
              #{{ order.orderNumber || 'نامشخص' }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <div>
                <div class="text-sm font-medium text-gray-900">{{ order.customerName || 'نامشخص' }}</div>
                <div class="text-sm text-gray-500">{{ order.customerPhone || 'نامشخص' }}</div>
              </div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
              {{ formatPrice(order.totalAmount || 0) }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
              {{ formatDate(order.shippedAt) }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
              {{ getCarrierText(order.carrier || 'نامشخص') }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
              <span class="font-mono">{{ order.trackingNumber }}</span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span :class="getStatusClass(order.status)" class="px-2 py-1 text-xs font-medium rounded-full">
                {{ getStatusText(order.status) }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
              <div class="flex items-center space-x-2">
                <button @click="viewOrderDetails(order)" class="text-blue-600 hover:text-blue-900">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                  </svg>
                </button>
                <button @click="trackOrder(order)" class="text-green-600 hover:text-green-900">رهگیری</button>
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
  orders: {
    type: Array,
    default: () => []
  }
})

// متغیرهای صفحه‌بندی
const currentPage = ref(1)
const itemsPerPage = ref(10)

// متغیرهای مودال
const isModalOpen = ref(false)
const selectedOrder = ref(null)

// محاسبه سفارشات فیلتر شده
const filteredOrders = computed(() => {
  return props.orders
})

// محاسبه تعداد صفحات
const totalPages = computed(() => {
  return Math.ceil(filteredOrders.value.length / itemsPerPage.value)
})

// محاسبه سفارشات نمایش داده شده در صفحه فعلی
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

const editOrder = (order) => {
  console.log('ویرایش سفارش:', order)
  // اینجا می‌توانید کاربر را به صفحه ویرایش هدایت کنید
}

// متدهای عملیاتی
const exportOrders = () => {
  console.log('خروجی Excel سفارشات ارسال شده...')
}

const trackOrder = (order) => {
  console.log('رهگیری سفارش:', order)
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
const formatDate = (date) => {
  return new Date(date).toLocaleDateString('fa-IR')
}

const formatPrice = (price) => {
  return new Intl.NumberFormat('fa-IR').format(price || 0) + ' تومان'
}

const getStatusClass = (status) => {
  const classes = {
    shipped: 'bg-blue-100 text-blue-800',
    in_transit: 'bg-yellow-100 text-yellow-800',
    delivered: 'bg-green-100 text-green-800',
    returned: 'bg-red-100 text-red-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status) => {
  const texts = {
    shipped: 'ارسال شده',
    in_transit: 'در راه',
    delivered: 'تحویل شده',
    returned: 'مرجوع شده'
  }
  return texts[status] || status
}

const getCarrierText = (carrier) => {
  const texts = {
    post: 'پست ایران',
    tipax: 'تیپاکس',
    delivery: 'پیک موتوری'
  }
  return texts[carrier] || carrier
}
</script> 
