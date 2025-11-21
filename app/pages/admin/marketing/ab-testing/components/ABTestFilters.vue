<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6">
    <div class="flex items-center justify-between mb-4">
      <h3 class="text-lg font-semibold text-gray-900">فیلترها و جستجو</h3>
      <button 
        class="text-sm text-gray-500 hover:text-gray-700 flex items-center" 
        @click="clearAllFilters"
      >
        <svg class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
        </svg>
        پاک کردن همه فیلترها
      </button>
    </div>

    <!-- جستجوی هوشمند -->
    <div class="mb-6">
      <div class="relative">
        <input
          v-model="searchQuery"
          type="text"
          placeholder="جستجو در نام تست‌ها، توضیحات و نتایج..."
          class="w-full pl-10 pr-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          @input="handleSearch"
        />
        <svg class="absolute left-3 top-3.5 w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
        </svg>
      </div>
    </div>

    <!-- فیلترهای پیشرفته -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <!-- فیلتر نوع تست -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">نوع تست</label>
        <select 
          v-model="filters.testType" 
          class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          @change="applyFilters"
        >
          <option value="">همه انواع</option>
          <option value="page">صفحه</option>
          <option value="button">دکمه</option>
          <option value="price">قیمت</option>
          <option value="image">تصویر</option>
          <option value="product">محصول</option>
        </select>
      </div>

      <!-- فیلتر وضعیت -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
        <select 
          v-model="filters.status" 
          class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          @change="applyFilters"
        >
          <option value="">همه وضعیت‌ها</option>
          <option value="active">فعال</option>
          <option value="inactive">غیرفعال</option>
          <option value="pending">در انتظار</option>
          <option value="completed">تکمیل شده</option>
        </select>
      </div>

      <!-- فیلتر تاریخ شروع -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ شروع</label>
        <input
          v-model="filters.startDate"
          type="date"
          class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          @change="applyFilters"
        />
      </div>

      <!-- فیلتر تاریخ پایان -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ پایان</label>
        <input
          v-model="filters.endDate"
          type="date"
          class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          @change="applyFilters"
        />
      </div>
    </div>

    <!-- فیلترهای اضافی -->
    <div class="mt-4 grid grid-cols-1 md:grid-cols-3 gap-6">
      <!-- فیلتر نتیجه -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">نتیجه</label>
        <select 
          v-model="filters.result" 
          class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          @change="applyFilters"
        >
          <option value="">همه نتایج</option>
          <option value="winner">برنده</option>
          <option value="loser">بازنده</option>
          <option value="inconclusive">نامشخص</option>
          <option value="running">در حال اجرا</option>
        </select>
      </div>

      <!-- فیلتر نرخ تبدیل -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">نرخ تبدیل (حداقل)</label>
        <input
          v-model="filters.minConversionRate"
          type="number"
          min="0"
          max="100"
          step="0.1"
          placeholder="0.0"
          class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          @input="applyFilters"
        />
      </div>

      <!-- فیلتر تعداد شرکت‌کنندگان -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">حداقل شرکت‌کننده</label>
        <input
          v-model="filters.minParticipants"
          type="number"
          min="0"
          placeholder="0"
          class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          @input="applyFilters"
        />
      </div>
    </div>

    <!-- فیلترهای پیشرفته -->
    <div class="mt-6">
      <button 
        class="text-sm text-blue-600 hover:text-blue-800 flex items-center"
        @click="toggleAdvancedFilters"
      >
        <svg 
          :class="['w-4 h-4 ml-1 transition-transform', { 'rotate-90': showAdvancedFilters }]" 
          fill="none" 
          stroke="currentColor" 
          viewBox="0 0 24 24"
        >
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
        </svg>
        فیلترهای پیشرفته
      </button>

      <div v-if="showAdvancedFilters" class="mt-4 p-6 bg-gray-50 rounded-lg">
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <!-- فیلتر سطح اطمینان -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">سطح اطمینان (حداقل)</label>
            <select 
              v-model="filters.minConfidenceLevel" 
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              @change="applyFilters"
            >
              <option value="">همه سطوح</option>
              <option value="90">90%</option>
              <option value="95">95%</option>
              <option value="99">99%</option>
            </select>
          </div>

          <!-- فیلتر معیار موفقیت -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">معیار موفقیت</label>
            <select 
              v-model="filters.successMetric" 
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              @change="applyFilters"
            >
              <option value="">همه معیارها</option>
              <option value="conversion_rate">نرخ تبدیل</option>
              <option value="revenue">درآمد</option>
              <option value="click_rate">نرخ کلیک</option>
              <option value="engagement">تعامل</option>
            </select>
          </div>

          <!-- فیلتر ایجادکننده -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">ایجادکننده</label>
            <select 
              v-model="filters.createdBy" 
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              @change="applyFilters"
            >
              <option value="">همه کاربران</option>
              <option value="admin">مدیر</option>
              <option value="marketing">بازاریابی</option>
              <option value="developer">توسعه‌دهنده</option>
            </select>
          </div>
        </div>
      </div>
    </div>

    <!-- نمایش فیلترهای فعال -->
    <div v-if="activeFiltersCount > 0" class="mt-4">
      <div class="flex items-center flex-wrap gap-2">
        <span class="text-sm text-gray-600">فیلترهای فعال:</span>
        <div 
          v-for="(filter, key) in activeFilters" 
          :key="key"
          class="inline-flex items-center px-3 py-1 rounded-full text-sm bg-blue-100 text-blue-800"
        >
          {{ filter.label }}: {{ filter.value }}
          <button 
            class="mr-1 text-blue-600 hover:text-blue-800"
            @click="removeFilter(String(key))"
          >
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
import { computed, ref, watch } from 'vue';
// Props
interface FilterValue {
  [key: string]: unknown
}

interface Props {
  modelValue?: FilterValue
}

const props = withDefaults(defineProps<Props>(), {
  modelValue: () => ({})
})

// Emits
interface Filters {
  [key: string]: unknown
}

const emit = defineEmits<{
  'update:modelValue': [value: FilterValue]
  'filter-change': [filters: Filters]
}>()

// State
const searchQuery = ref('')
const showAdvancedFilters = ref(false)

// فیلترها
const filters = ref({
  testType: '',
  status: '',
  startDate: '',
  endDate: '',
  result: '',
  minConversionRate: '',
  minParticipants: '',
  minConfidenceLevel: '',
  successMetric: '',
  createdBy: ''
})

// محاسبه تعداد فیلترهای فعال
const activeFiltersCount = computed(() => {
  return Object.values(filters.value).filter(value => value !== '').length
})

// نمایش فیلترهای فعال
const activeFilters = computed(() => {
  const active: Record<string, { label: string; value: string }> = {}
  
  if (filters.value.testType) {
    active.testType = {
      label: 'نوع تست',
      value: getTestTypeLabel(filters.value.testType)
    }
  }
  
  if (filters.value.status) {
    active.status = {
      label: 'وضعیت',
      value: getStatusLabel(filters.value.status)
    }
  }
  
  if (filters.value.startDate) {
    active.startDate = {
      label: 'تاریخ شروع',
      value: filters.value.startDate
    }
  }
  
  if (filters.value.endDate) {
    active.endDate = {
      label: 'تاریخ پایان',
      value: filters.value.endDate
    }
  }
  
  if (filters.value.result) {
    active.result = {
      label: 'نتیجه',
      value: getResultLabel(filters.value.result)
    }
  }
  
  if (filters.value.minConversionRate) {
    active.minConversionRate = {
      label: 'حداقل نرخ تبدیل',
      value: `${filters.value.minConversionRate.toString()}%`
    }
  }
  
  if (filters.value.minParticipants) {
    active.minParticipants = {
      label: 'حداقل شرکت‌کننده',
      value: filters.value.minParticipants
    }
  }
  
  return active
})

// توابع کمکی
const getTestTypeLabel = (type: string) => {
  const labels: Record<string, string> = {
    page: 'صفحه',
    button: 'دکمه',
    price: 'قیمت',
    image: 'تصویر',
    product: 'محصول'
  }
  return labels[type] || type
}

const getStatusLabel = (status: string) => {
  const labels: Record<string, string> = {
    active: 'فعال',
    inactive: 'غیرفعال',
    pending: 'در انتظار',
    completed: 'تکمیل شده'
  }
  return labels[status] || status
}

const getResultLabel = (result: string) => {
  const labels: Record<string, string> = {
    winner: 'برنده',
    loser: 'بازنده',
    inconclusive: 'نامشخص',
    running: 'در حال اجرا'
  }
  return labels[result] || result
}

// توگل فیلترهای پیشرفته
const toggleAdvancedFilters = () => {
  showAdvancedFilters.value = !showAdvancedFilters.value
}

// اعمال فیلترها
const applyFilters = () => {
  const allFilters = {
    search: searchQuery.value,
    ...filters.value
  }
  
  emit('update:modelValue', allFilters)
  emit('filter-change', allFilters)
}

// جستجو
let searchTimeout: NodeJS.Timeout
const handleSearch = () => {
  clearTimeout(searchTimeout)
  searchTimeout = setTimeout(() => {
    applyFilters()
  }, 300)
}

// حذف فیلتر
const removeFilter = (filterKey: string) => {
  if (filterKey in filters.value) {
    filters.value[filterKey] = ''
  }
  applyFilters()
}

// پاک کردن همه فیلترها
const clearAllFilters = () => {
  searchQuery.value = ''
  filters.value = {
    testType: '',
    status: '',
    startDate: '',
    endDate: '',
    result: '',
    minConversionRate: '',
    minParticipants: '',
    minConfidenceLevel: '',
    successMetric: '',
    createdBy: ''
  }
  applyFilters()
}

// Watch برای تغییرات props
watch(() => props.modelValue, (newValue) => {
  if (newValue) {
    searchQuery.value = newValue.search || ''
    filters.value = {
      testType: newValue.testType || '',
      status: newValue.status || '',
      startDate: newValue.startDate || '',
      endDate: newValue.endDate || '',
      result: newValue.result || '',
      minConversionRate: newValue.minConversionRate || '',
      minParticipants: newValue.minParticipants || '',
      minConfidenceLevel: newValue.minConfidenceLevel || '',
      successMetric: newValue.successMetric || '',
      createdBy: newValue.createdBy || ''
    }
  }
}, { immediate: true })
</script> 
