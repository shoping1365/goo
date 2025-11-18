<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-4">
      <h3 class="text-lg font-semibold text-gray-900">فیلترها</h3>
      <div class="flex items-center space-x-2 space-x-reverse">
        <button @click="clearFilters" class="text-sm text-gray-600 hover:text-gray-800">
          پاک کردن فیلترها
        </button>
        <button @click="applyFilters" class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 text-sm">
          اعمال فیلترها
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <!-- Status Filter -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
        <select v-model="filters.status" class="w-full border border-gray-300 rounded-md px-3 py-2 text-sm">
          <option value="">همه وضعیت‌ها</option>
          <option value="pending">در انتظار بررسی</option>
          <option value="approved">تایید شده</option>
          <option value="rejected">رد شده</option>
          <option value="completed">تکمیل شده</option>
          <option value="cancelled">لغو شده</option>
        </select>
      </div>

      <!-- Credit Score Range -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">امتیاز اعتباری</label>
        <select v-model="filters.creditScore" class="w-full border border-gray-300 rounded-md px-3 py-2 text-sm">
          <option value="">همه امتیازات</option>
          <option value="excellent">عالی (80-100)</option>
          <option value="good">خوب (60-79)</option>
          <option value="fair">متوسط (40-59)</option>
          <option value="poor">ضعیف (0-39)</option>
        </select>
      </div>

      <!-- Amount Range -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">مبلغ</label>
        <select v-model="filters.amountRange" class="w-full border border-gray-300 rounded-md px-3 py-2 text-sm">
          <option value="">همه مبالغ</option>
          <option value="low">کمتر از 10 میلیون</option>
          <option value="medium">10 تا 50 میلیون</option>
          <option value="high">بیش از 50 میلیون</option>
        </select>
      </div>

      <!-- Installment Count -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">تعداد اقساط</label>
        <select v-model="filters.installmentCount" class="w-full border border-gray-300 rounded-md px-3 py-2 text-sm">
          <option value="">همه</option>
          <option value="3">3 قسط</option>
          <option value="6">6 قسط</option>
          <option value="12">12 قسط</option>
          <option value="24">24 قسط</option>
          <option value="36">36 قسط</option>
        </select>
      </div>
    </div>

    <!-- Advanced Filters -->
    <div class="mt-6 pt-6 border-t border-gray-200">
      <div class="flex items-center justify-between mb-4">
        <h4 class="text-md font-medium text-gray-900">فیلترهای پیشرفته</h4>
        <button @click="showAdvancedFilters = !showAdvancedFilters" class="text-blue-600 hover:text-blue-800 text-sm">
          {{ showAdvancedFilters ? 'مخفی کردن' : 'نمایش' }}
        </button>
      </div>

      <div v-if="showAdvancedFilters" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <!-- Date Range -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">بازه زمانی</label>
          <div class="grid grid-cols-2 gap-2">
            <input 
              v-model="filters.startDate" 
              type="date" 
              class="border border-gray-300 rounded-md px-3 py-2 text-sm"
              placeholder="از تاریخ"
            >
            <input 
              v-model="filters.endDate" 
              type="date" 
              class="border border-gray-300 rounded-md px-3 py-2 text-sm"
              placeholder="تا تاریخ"
            >
          </div>
        </div>

        <!-- Customer Type -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نوع مشتری</label>
          <select v-model="filters.customerType" class="w-full border border-gray-300 rounded-md px-3 py-2 text-sm">
            <option value="">همه</option>
            <option value="new">مشتری جدید</option>
            <option value="returning">مشتری تکراری</option>
            <option value="vip">مشتری VIP</option>
          </select>
        </div>

        <!-- Product Category -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">دسته‌بندی محصول</label>
          <select v-model="filters.productCategory" class="w-full border border-gray-300 rounded-md px-3 py-2 text-sm">
            <option value="">همه دسته‌بندی‌ها</option>
            <option value="electronics">الکترونیک</option>
            <option value="furniture">مبل و دکوراسیون</option>
            <option value="clothing">پوشاک</option>
            <option value="automotive">خودرو</option>
            <option value="real_estate">املاک</option>
          </select>
        </div>

        <!-- Risk Level -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">سطح ریسک</label>
          <select v-model="filters.riskLevel" class="w-full border border-gray-300 rounded-md px-3 py-2 text-sm">
            <option value="">همه سطوح</option>
            <option value="low">ریسک پایین</option>
            <option value="medium">ریسک متوسط</option>
            <option value="high">ریسک بالا</option>
          </select>
        </div>

        <!-- Payment Method -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">روش پرداخت</label>
          <select v-model="filters.paymentMethod" class="w-full border border-gray-300 rounded-md px-3 py-2 text-sm">
            <option value="">همه روش‌ها</option>
            <option value="bank_transfer">انتقال بانکی</option>
            <option value="check">چک</option>
            <option value="cash">نقدی</option>
            <option value="online">آنلاین</option>
          </select>
        </div>

        <!-- Approval Method -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">روش تایید</label>
          <select v-model="filters.approvalMethod" class="w-full border border-gray-300 rounded-md px-3 py-2 text-sm">
            <option value="">همه روش‌ها</option>
            <option value="automatic">خودکار</option>
            <option value="manual">دستی</option>
            <option value="review">بررسی تخصصی</option>
          </select>
        </div>
      </div>
    </div>

    <!-- Search -->
    <div class="mt-6 pt-6 border-t border-gray-200">
      <div class="flex items-center space-x-4 space-x-reverse">
        <div class="flex-1">
          <label class="block text-sm font-medium text-gray-700 mb-2">جستجو</label>
          <div class="relative">
            <input 
              v-model="filters.search" 
              type="text" 
              class="w-full border border-gray-300 rounded-md pl-10 pr-3 py-2 text-sm"
              placeholder="جستجو بر اساس نام، کد ملی، شماره موبایل..."
            >
            <div class="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
              <svg class="h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
            </div>
          </div>
        </div>
        <div class="flex items-end">
          <button @click="applyFilters" class="bg-blue-600 text-white px-6 py-2 rounded-md hover:bg-blue-700 text-sm">
            جستجو
          </button>
        </div>
      </div>
    </div>

    <!-- Active Filters -->
    <div v-if="hasActiveFilters" class="mt-4 pt-4 border-t border-gray-200">
      <div class="flex items-center space-x-2 space-x-reverse flex-wrap">
        <span class="text-sm text-gray-600">فیلترهای فعال:</span>
        <div v-for="(filter, key) in activeFilters" :key="key" class="inline-flex items-center bg-blue-100 text-blue-800 text-xs px-2 py-1 rounded-full">
          <span>{{ filter.label }}: {{ filter.value }}</span>
          <button @click="removeFilter(key)" class="mr-1 text-blue-600 hover:text-blue-800">
            <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'

interface Filters {
  status: string
  creditScore: string
  amountRange: string
  installmentCount: string
  startDate: string
  endDate: string
  customerType: string
  productCategory: string
  riskLevel: string
  paymentMethod: string
  approvalMethod: string
  search: string
}

const emit = defineEmits<{
  filterChange: [filters: Filters]
}>()

const showAdvancedFilters = ref(false)

const filters = ref<Filters>({
  status: '',
  creditScore: '',
  amountRange: '',
  installmentCount: '',
  startDate: '',
  endDate: '',
  customerType: '',
  productCategory: '',
  riskLevel: '',
  paymentMethod: '',
  approvalMethod: '',
  search: ''
})

const filterLabels = {
  status: {
    pending: 'در انتظار بررسی',
    approved: 'تایید شده',
    rejected: 'رد شده',
    completed: 'تکمیل شده',
    cancelled: 'لغو شده'
  },
  creditScore: {
    excellent: 'عالی (80-100)',
    good: 'خوب (60-79)',
    fair: 'متوسط (40-59)',
    poor: 'ضعیف (0-39)'
  },
  amountRange: {
    low: 'کمتر از 10 میلیون',
    medium: '10 تا 50 میلیون',
    high: 'بیش از 50 میلیون'
  },
  customerType: {
    new: 'مشتری جدید',
    returning: 'مشتری تکراری',
    vip: 'مشتری VIP'
  },
  productCategory: {
    electronics: 'الکترونیک',
    furniture: 'مبل و دکوراسیون',
    clothing: 'پوشاک',
    automotive: 'خودرو',
    real_estate: 'املاک'
  },
  riskLevel: {
    low: 'ریسک پایین',
    medium: 'ریسک متوسط',
    high: 'ریسک بالا'
  },
  paymentMethod: {
    bank_transfer: 'انتقال بانکی',
    check: 'چک',
    cash: 'نقدی',
    online: 'آنلاین'
  },
  approvalMethod: {
    automatic: 'خودکار',
    manual: 'دستی',
    review: 'بررسی تخصصی'
  }
}

const activeFilters = computed(() => {
  const active: Record<string, { label: string; value: string }> = {}
  
  Object.entries(filters.value).forEach(([key, value]) => {
    if (value && key !== 'startDate' && key !== 'endDate' && key !== 'search') {
      const label = getFilterLabel(key)
      if (label) {
        active[key] = { label, value: getFilterValue(key, value) }
      }
    }
  })

  if (filters.value.startDate || filters.value.endDate) {
    active.dateRange = {
      label: 'بازه زمانی',
      value: `${filters.value.startDate || 'از ابتدا'} تا ${filters.value.endDate || 'تا انتها'}`
    }
  }

  if (filters.value.search) {
    active.search = {
      label: 'جستجو',
      value: filters.value.search
    }
  }

  return active
})

const hasActiveFilters = computed(() => {
  return Object.keys(activeFilters.value).length > 0
})

const getFilterLabel = (key: string): string => {
  const labels: Record<string, string> = {
    status: 'وضعیت',
    creditScore: 'امتیاز اعتباری',
    amountRange: 'مبلغ',
    installmentCount: 'تعداد اقساط',
    customerType: 'نوع مشتری',
    productCategory: 'دسته‌بندی محصول',
    riskLevel: 'سطح ریسک',
    paymentMethod: 'روش پرداخت',
    approvalMethod: 'روش تایید'
  }
  return labels[key] || key
}

const getFilterValue = (key: string, value: string): string => {
  const labelMap = filterLabels[key as keyof typeof filterLabels]
  if (labelMap && labelMap[value as keyof typeof labelMap]) {
    return labelMap[value as keyof typeof labelMap]
  }
  return value
}

const applyFilters = () => {
  emit('filterChange', { ...filters.value })
}

const clearFilters = () => {
  filters.value = {
    status: '',
    creditScore: '',
    amountRange: '',
    installmentCount: '',
    startDate: '',
    endDate: '',
    customerType: '',
    productCategory: '',
    riskLevel: '',
    paymentMethod: '',
    approvalMethod: '',
    search: ''
  }
  applyFilters()
}

const removeFilter = (key: string) => {
  if (key === 'dateRange') {
    filters.value.startDate = ''
    filters.value.endDate = ''
  } else if (key === 'search') {
    filters.value.search = ''
  } else {
    filters.value[key as keyof Filters] = ''
  }
  applyFilters()
}

// Auto-apply filters when they change
watch(filters, () => {
  applyFilters()
}, { deep: true })
</script> 
