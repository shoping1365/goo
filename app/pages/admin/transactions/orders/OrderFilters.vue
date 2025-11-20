<template>
  <div class="bg-gradient-to-br from-gray-50 to-gray-100 rounded-xl shadow-lg border border-gray-200 px-4 py-4">
    <!-- Advanced Filters -->
              <div class="flex items-center justify-between mb-6">
        <div class="flex items-center">
          <button class="w-8 h-8 bg-gradient-to-r from-blue-500 to-purple-600 rounded-lg flex items-center justify-center ml-3 text-white hover:bg-blue-600 transition-colors" @click="showAdvanced = !showAdvanced">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.207A1 1 0 013 6.5V4z"></path>
            </svg>
      </button>
        </div>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gapx-4 py-4">
      <!-- Search -->
      <div class="space-y-2">
        <label class="block text-sm font-semibold text-gray-700">جستجو</label>
        <div class="relative">
        <input
          v-model="filters.search"
          type="text"
          placeholder="شماره سفارش، نام مشتری..."
            class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all duration-200 bg-white shadow-sm"
          @input="applyFilters"
        />
          <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
            <svg class="h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- Status Filter -->
      <div class="space-y-2">
        <label class="block text-sm font-semibold text-gray-700">وضعیت</label>
        <select
          v-model="filters.status"
          class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all duration-200 bg-white shadow-sm"
          @change="applyFilters"
        >
          <option value="">همه وضعیت‌ها</option>
          <option value="pending">در انتظار پرداخت</option>
          <option value="paid">پرداخت شده</option>
          <option value="processing">در حال پردازش</option>
          <option value="shipped">ارسال شده</option>
          <option value="delivered">تحویل داده شده</option>
          <option value="cancelled">لغو شده</option>
          <option value="refunded">بازگشت وجه</option>
        </select>
      </div>

      <!-- Date Range -->
      <div class="space-y-2">
        <label class="block text-sm font-semibold text-gray-700">از تاریخ</label>
        <input
          v-model="filters.startDate"
          type="date"
          class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all duration-200 bg-white shadow-sm"
          @change="applyFilters"
        />
      </div>

      <div class="space-y-2">
        <label class="block text-sm font-semibold text-gray-700">تا تاریخ</label>
        <input
          v-model="filters.endDate"
          type="date"
          class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all duration-200 bg-white shadow-sm"
          @change="applyFilters"
        />
      </div>

        <!-- Amount Range -->
      <div class="space-y-2">
        <label class="block text-sm font-semibold text-gray-700">حداقل مبلغ (تومان)</label>
          <input
            v-model="filters.minAmount"
            type="number"
            placeholder="0"
          class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all duration-200 bg-white shadow-sm"
            @input="applyFilters"
          />
        </div>

      <div class="space-y-2">
        <label class="block text-sm font-semibold text-gray-700">حداکثر مبلغ (تومان)</label>
          <input
            v-model="filters.maxAmount"
            type="number"
            placeholder="نامحدود"
          class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all duration-200 bg-white shadow-sm"
            @input="applyFilters"
          />
        </div>

        <!-- Payment Method -->
      <div class="space-y-2">
        <label class="block text-sm font-semibold text-gray-700">روش پرداخت</label>
          <select
            v-model="filters.paymentMethod"
          class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all duration-200 bg-white shadow-sm"
            @change="applyFilters"
          >
            <option value="">همه روش‌ها</option>
            <option value="online">پرداخت آنلاین</option>
            <option value="cash">پرداخت نقدی</option>
            <option value="bank_transfer">انتقال بانکی</option>
          </select>
        </div>

        <!-- City -->
      <div class="space-y-2">
        <label class="block text-sm font-semibold text-gray-700">شهر</label>
          <input
            v-model="filters.city"
            type="text"
            placeholder="نام شهر"
          class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all duration-200 bg-white shadow-sm"
            @input="applyFilters"
          />
        </div>

        <!-- Sort By -->
      <div class="space-y-2">
        <label class="block text-sm font-semibold text-gray-700">مرتب‌سازی بر اساس</label>
          <select
            v-model="filters.sortBy"
          class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all duration-200 bg-white shadow-sm"
            @change="applyFilters"
          >
            <option value="created_at">تاریخ سفارش</option>
            <option value="total_amount">مبلغ سفارش</option>
            <option value="order_number">شماره سفارش</option>
            <option value="customer_name">نام مشتری</option>
          </select>
        </div>

        <!-- Sort Order -->
      <div class="space-y-2">
        <label class="block text-sm font-semibold text-gray-700">ترتیب</label>
          <select
            v-model="filters.sortOrder"
          class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all duration-200 bg-white shadow-sm"
            @change="applyFilters"
          >
            <option value="desc">نزولی</option>
            <option value="asc">صعودی</option>
          </select>
      </div>
    </div>

    <!-- Active Filters -->
    <div v-if="hasActiveFilters" class="mt-6 pt-6 border-t border-gray-200">
      <div class="flex items-center space-x-2 flex-wrap gap-2">
        <span class="text-sm font-semibold text-gray-700">فیلترهای انتخاب شده:</span>
        <span
          v-for="(value, key) in activeFilters"
          :key="key"
          class="inline-flex items-center px-3 py-1.5 rounded-full text-xs font-medium bg-gradient-to-r from-blue-100 to-blue-200 text-blue-800 border border-blue-300 shadow-sm"
        >
          {{ getFilterLabel(key, value) }}
          <button class="mr-1 text-blue-600 hover:text-blue-800 transition-colors" @click="removeFilter(key)">
            <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, ref, watch } from 'vue'

const props = defineProps({
  initialFilters: {
    type: Object,
    default: () => ({})
  }
})

const emit = defineEmits(['filter-change'])

const showAdvanced = ref(false)

const filters = ref({
  search: '',
  status: '',
  startDate: '',
  endDate: '',
  minAmount: '',
  maxAmount: '',
  paymentMethod: '',
  city: '',
  sortBy: 'created_at',
  sortOrder: 'desc',
  ...props.initialFilters
})

const activeFilters = computed(() => {
  const active = {}
  Object.entries(filters.value).forEach(([key, value]) => {
    if (value && value !== '' && key !== 'sortBy' && key !== 'sortOrder') {
      active[key] = value
    }
  })
  return active
})

const hasActiveFilters = computed(() => {
  return Object.keys(activeFilters.value).length > 0
})

const applyFilters = () => {
  emit('filter-change', { ...filters.value })
}

const _clearFilters = () => {
  filters.value = {
    search: '',
    status: '',
    startDate: '',
    endDate: '',
    minAmount: '',
    maxAmount: '',
    paymentMethod: '',
    city: '',
    sortBy: 'created_at',
    sortOrder: 'desc'
  }
  applyFilters()
}

const removeFilter = (key) => {
  filters.value[key] = ''
  applyFilters()
}

const getFilterLabel = (key, value) => {
  const labels = {
    search: `جستجو: ${value}`,
    status: `وضعیت: ${getStatusText(value)}`,
    startDate: `از تاریخ: ${value}`,
    endDate: `تا تاریخ: ${value}`,
    minAmount: `حداقل مبلغ: ${value} تومان`,
    maxAmount: `حداکثر مبلغ: ${value} تومان`,
    paymentMethod: `روش پرداخت: ${getPaymentMethodText(value)}`,
    city: `شهر: ${value}`
  }
  return labels[key] || `${key}: ${value}`
}

const getStatusText = (status) => {
  const statusMap = {
    'pending': 'در انتظار پرداخت',
    'paid': 'پرداخت شده',
    'processing': 'در حال پردازش',
    'shipped': 'ارسال شده',
    'delivered': 'تحویل داده شده',
    'cancelled': 'لغو شده',
    'refunded': 'بازگشت وجه'
  }
  return statusMap[status] || status
}

const getPaymentMethodText = (method) => {
  const methodMap = {
    'online': 'پرداخت آنلاین',
    'cash': 'پرداخت نقدی',
    'bank_transfer': 'انتقال بانکی'
  }
  return methodMap[method] || method
}

// Watch for initial filters changes
watch(() => props.initialFilters, (newFilters) => {
  filters.value = { ...filters.value, ...newFilters }
}, { deep: true })
</script> 
