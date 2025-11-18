<template>
  <div class="bg-white rounded-lg shadow-md border border-gray-200 mb-6">
    <!-- هدر فیلتر -->
    <div class="bg-gradient-to-r from-blue-50 to-indigo-50 px-4 py-3 border-b border-gray-200">
      <div class="flex items-center justify-between">
        <div class="flex items-center">
          <div class="w-6 h-6 bg-blue-500 rounded-md flex items-center justify-center ml-2">
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.207A1 1 0 013 6.5V4z"></path>
            </svg>
          </div>
          <h3 class="text-sm font-semibold text-gray-900">فیلتر پیشرفته</h3>
        </div>
        <div class="flex items-center space-x-2 space-x-reverse">
          <button 
            @click="toggleAdvancedFilters"
            class="text-blue-600 hover:text-blue-800 text-sm font-medium flex items-center"
          >
            <svg class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
            </svg>
            {{ showAdvancedFilters ? 'بستن' : 'نمایش' }} فیلترها
          </button>
        </div>
      </div>
    </div>

    <!-- فیلترهای پیشرفته -->
    <div v-if="showAdvancedFilters" class="px-4 py-4">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gapx-4 py-4">
        <!-- جستجو -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">جستجو</label>
          <div class="relative">
            <input
              v-model="filters.searchTerm"
              type="text"
              placeholder="شماره سفارش، نام مشتری..."
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
            <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
              <svg class="h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
              </svg>
            </div>
          </div>
        </div>

        <!-- وضعیت -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
          <select 
            v-model="filters.status" 
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه وضعیت‌ها</option>
            <option value="shipped">ارسال شده</option>
            <option value="in_transit">در راه</option>
            <option value="delivered">تحویل شده</option>
            <option value="returned">مرجوع شده</option>
          </select>
        </div>

        <!-- شرکت ارسال -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">شرکت ارسال</label>
          <select 
            v-model="filters.carrier" 
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه شرکت‌ها</option>
            <option value="post">پست ایران</option>
            <option value="tipax">تیپاکس</option>
            <option value="delivery">پیک موتوری</option>
            <option value="snap">اسنپ</option>
          </select>
        </div>

        <!-- محدوده مبلغ -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">محدوده مبلغ (تومان)</label>
          <div class="grid grid-cols-2 gap-2">
            <input
              v-model="filters.minAmount"
              type="number"
              placeholder="حداقل"
              class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
            <input
              v-model="filters.maxAmount"
              type="number"
              placeholder="حداکثر"
              class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
          </div>
        </div>

        <!-- تاریخ ارسال از -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ ارسال از</label>
          <input
            v-model="filters.shippedFrom"
            type="date"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>

        <!-- تاریخ ارسال تا -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ ارسال تا</label>
          <input
            v-model="filters.shippedTo"
            type="date"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>

        <!-- تاریخ تحویل از -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ تحویل از</label>
          <input
            v-model="filters.deliveredFrom"
            type="date"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>

        <!-- تاریخ تحویل تا -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ تحویل تا</label>
          <input
            v-model="filters.deliveredTo"
            type="date"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>

        <!-- منطقه ارسال -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">منطقه ارسال</label>
          <select 
            v-model="filters.region" 
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه مناطق</option>
            <option value="tehran">تهران</option>
            <option value="isfahan">اصفهان</option>
            <option value="shiraz">شیراز</option>
            <option value="mashhad">مشهد</option>
            <option value="tabriz">تبریز</option>
            <option value="other">سایر</option>
          </select>
        </div>

        <!-- نوع بسته‌بندی -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نوع بسته‌بندی</label>
          <select 
            v-model="filters.packageType" 
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه انواع</option>
            <option value="small">کوچک</option>
            <option value="medium">متوسط</option>
            <option value="large">بزرگ</option>
            <option value="fragile">شکننده</option>
          </select>
        </div>

        <!-- اولویت ارسال -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">اولویت ارسال</label>
          <select 
            v-model="filters.priority" 
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه اولویت‌ها</option>
            <option value="normal">عادی</option>
            <option value="express">سریع</option>
            <option value="urgent">فوری</option>
          </select>
        </div>

        <!-- وضعیت پیگیری -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت پیگیری</label>
          <select 
            v-model="filters.trackingStatus" 
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه وضعیت‌ها</option>
            <option value="active">فعال</option>
            <option value="inactive">غیرفعال</option>
            <option value="expired">منقضی شده</option>
          </select>
        </div>
      </div>

      <!-- دکمه‌های عملیات -->
      <div class="flex items-center justify-between mt-6 pt-6 border-t border-gray-200">
        <div class="flex items-center space-x-2 space-x-reverse">
          <button 
            @click="applyFilters"
            class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 transition-colors"
          >
            اعمال فیلترها
          </button>
          <button 
            @click="clearAllFilters"
            class="px-4 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2 transition-colors"
          >
            پاک کردن همه
          </button>
        </div>
        
        <div class="text-sm text-gray-500">
          {{ activeFiltersCount }} فیلتر فعال
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// Props
const props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({})
  }
})

// Emits
const emit = defineEmits(['update:modelValue', 'filter'])

// متغیرهای محلی
const showAdvancedFilters = ref(false)

// فیلترها
const filters = ref({
  searchTerm: '',
  status: '',
  carrier: '',
  minAmount: '',
  maxAmount: '',
  shippedFrom: '',
  shippedTo: '',
  deliveredFrom: '',
  deliveredTo: '',
  region: '',
  packageType: '',
  priority: '',
  trackingStatus: ''
})

// محاسبه تعداد فیلترهای فعال
const activeFiltersCount = computed(() => {
  return Object.values(filters.value).filter(value => value !== '').length
})

// متدها
const toggleAdvancedFilters = () => {
  showAdvancedFilters.value = !showAdvancedFilters.value
}

const applyFilters = () => {
  emit('update:modelValue', { ...filters.value })
  emit('filter', { ...filters.value })
}

const clearAllFilters = () => {
  filters.value = {
    searchTerm: '',
    status: '',
    carrier: '',
    minAmount: '',
    maxAmount: '',
    shippedFrom: '',
    shippedTo: '',
    deliveredFrom: '',
    deliveredTo: '',
    region: '',
    packageType: '',
    priority: '',
    trackingStatus: ''
  }
  applyFilters()
}

// نظارت بر تغییرات props
watch(() => props.modelValue, (newValue) => {
  if (newValue) {
    filters.value = { ...newValue }
  }
}, { deep: true })

// مقداردهی اولیه
onMounted(() => {
  if (props.modelValue) {
    filters.value = { ...props.modelValue }
  }
})
</script> 
