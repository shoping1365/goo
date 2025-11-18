<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">فیلترهای پیشرفته</h4>
        <p class="text-sm text-gray-600 mt-1">مدیریت و تنظیم فیلترهای پیشرفته</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          @click="saveFilterPreset"
          class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
          </svg>
          ذخیره فیلتر
        </button>
      </div>
    </div>

    <!-- فیلترهای فعال -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">فیلترهای فعال</h5>
      <div class="flex flex-wrap gap-2">
        <div
          v-for="filter in activeFilters"
          :key="filter.id"
          class="inline-flex items-center px-3 py-1 rounded-full text-sm bg-blue-100 text-blue-800"
        >
          <span>{{ filter.label }}: {{ filter.value }}</span>
          <button
            @click="removeFilter(filter.id)"
            class="mr-2 text-blue-600 hover:text-blue-800"
          >
            <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
        <button
          v-if="activeFilters.length > 0"
          @click="clearAllFilters"
          class="inline-flex items-center px-3 py-1 rounded-full text-sm bg-gray-100 text-gray-700 hover:bg-gray-200"
        >
          پاک کردن همه
        </button>
      </div>
    </div>

    <!-- انواع فیلتر -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">انواع فیلتر</h5>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div
          v-for="filterType in filterTypes"
          :key="filterType.id"
          class="border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow"
        >
          <div class="flex items-center justify-between mb-3">
            <h6 class="text-sm font-medium text-gray-900">{{ filterType.name }}</h6>
            <button
              @click="addFilter(filterType)"
              class="text-blue-600 hover:text-blue-800"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
              </svg>
            </button>
          </div>
          <p class="text-xs text-gray-600 mb-3">{{ filterType.description }}</p>
          <div class="flex items-center text-xs text-gray-500">
            <svg class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="filterType.icon" />
            </svg>
            {{ filterType.examples }}
          </div>
        </div>
      </div>
    </div>

    <!-- فیلترهای سفارشی -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">فیلترهای سفارشی</h5>
      <div class="space-y-4">
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نام فیلتر</label>
            <input
              v-model="customFilter.name"
              type="text"
              class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="نام فیلتر سفارشی"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع فیلتر</label>
            <select
              v-model="customFilter.type"
              class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="text">متنی</option>
              <option value="number">عددی</option>
              <option value="date">تاریخ</option>
              <option value="select">انتخابی</option>
              <option value="range">بازه</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">فیلد هدف</label>
            <select
              v-model="customFilter.field"
              class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="transaction_id">شناسه تراکنش</option>
              <option value="amount">مبلغ</option>
              <option value="status">وضعیت</option>
              <option value="date">تاریخ</option>
              <option value="gateway">درگاه</option>
            </select>
          </div>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مقدار پیش‌فرض</label>
            <input
              v-model="customFilter.defaultValue"
              type="text"
              class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="مقدار پیش‌فرض"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">عملگر</label>
            <select
              v-model="customFilter.operator"
              class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="equals">برابر</option>
              <option value="contains">شامل</option>
              <option value="starts_with">شروع با</option>
              <option value="ends_with">پایان با</option>
              <option value="greater_than">بزرگتر از</option>
              <option value="less_than">کوچکتر از</option>
            </select>
          </div>
        </div>

        <div class="flex items-center justify-end space-x-3 space-x-reverse">
          <button
            @click="resetCustomFilter"
            class="px-4 py-2 border border-gray-300 rounded-md text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            بازنشانی
          </button>
          <button
            @click="createCustomFilter"
            class="px-4 py-2 border border-transparent rounded-md text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            ایجاد فیلتر
          </button>
        </div>
      </div>
    </div>

    <!-- فیلترهای ذخیره شده -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">فیلترهای ذخیره شده</h5>
      </div>
      <div class="p-6">
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <div
            v-for="preset in savedPresets"
            :key="preset.id"
            class="border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow"
          >
            <div class="flex items-center justify-between mb-3">
              <h6 class="text-sm font-medium text-gray-900">{{ preset.name }}</h6>
              <div class="flex items-center space-x-2 space-x-reverse">
                <button
                  @click="applyPreset(preset)"
                  class="text-green-600 hover:text-green-800"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                  </svg>
                </button>
                <button
                  @click="editPreset(preset)"
                  class="text-blue-600 hover:text-blue-800"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                  </svg>
                </button>
                <button
                  @click="deletePreset(preset)"
                  class="text-red-600 hover:text-red-800"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                  </svg>
                </button>
              </div>
            </div>
            <div class="text-xs text-gray-600 mb-2">{{ preset.description }}</div>
            <div class="flex items-center justify-between text-xs text-gray-500">
              <span>{{ preset.filterCount }} فیلتر</span>
              <span>{{ preset.lastUsed }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات فیلتر -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">تنظیمات فیلتر</h5>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر فیلترهای همزمان</label>
          <input
            v-model="filterSettings.maxFilters"
            type="number"
            min="1"
            max="20"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نمایش فیلترهای پیشنهادی</label>
          <select
            v-model="filterSettings.showSuggestions"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="always">همیشه</option>
            <option value="onFocus">در زمان فوکوس</option>
            <option value="never">هرگز</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">ذخیره خودکار فیلترها</label>
          <select
            v-model="filterSettings.autoSave"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="enabled">فعال</option>
            <option value="disabled">غیرفعال</option>
            <option value="prompt">درخواست تایید</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نمایش تعداد نتایج</label>
          <select
            v-model="filterSettings.showResultCount"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="always">همیشه</option>
            <option value="onChange">در زمان تغییر</option>
            <option value="never">هرگز</option>
          </select>
        </div>
      </div>

      <div class="mt-4 space-y-3">
        <label class="flex items-center">
          <input
            v-model="filterSettings.rememberFilters"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">به خاطر سپردن فیلترها</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="filterSettings.highlightResults"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">برجسته کردن نتایج</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="filterSettings.animateFilters"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">انیمیشن فیلترها</span>
        </label>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// فیلترهای فعال
const activeFilters = ref([
  {
    id: 1,
    label: 'وضعیت',
    value: 'موفق'
  },
  {
    id: 2,
    label: 'مبلغ',
    value: 'بیش از ۱ میلیون'
  }
])

// انواع فیلتر
const filterTypes = ref([
  {
    id: 'text',
    name: 'فیلتر متنی',
    description: 'جستجو در متن و محتوا',
    examples: 'شناسه، نام، توضیحات',
    icon: 'M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z'
  },
  {
    id: 'number',
    name: 'فیلتر عددی',
    description: 'فیلتر بر اساس اعداد',
    examples: 'مبلغ، تعداد، درصد',
    icon: 'M9 7h6m0 10v-3m-3 3h.01M9 17h.01M9 14h.01M12 14h.01M15 11h.01M12 11h.01M9 11h.01M7 21h10a2 2 0 002-2V5a2 2 0 00-2-2H7a2 2 0 00-2 2v14a2 2 0 002 2z'
  },
  {
    id: 'date',
    name: 'فیلتر تاریخ',
    description: 'فیلتر بر اساس تاریخ',
    examples: 'تاریخ تراکنش، بازه زمانی',
    icon: 'M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z'
  },
  {
    id: 'select',
    name: 'فیلتر انتخابی',
    description: 'انتخاب از لیست گزینه‌ها',
    examples: 'وضعیت، درگاه، نوع',
    icon: 'M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2'
  },
  {
    id: 'range',
    name: 'فیلتر بازه',
    description: 'فیلتر در بازه مشخص',
    examples: 'مبلغ، تاریخ، تعداد',
    icon: 'M7 12l3-3 3 3 4-4M8 21l4-4 4 4M3 4h18M4 4h16v12a1 1 0 01-1 1H5a1 1 0 01-1-1V4z'
  },
  {
    id: 'boolean',
    name: 'فیلتر منطقی',
    description: 'فیلتر بله/خیر',
    examples: 'فعال، تایید شده، مهم',
    icon: 'M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z'
  }
])

// فیلتر سفارشی
const customFilter = ref({
  name: '',
  type: 'text',
  field: 'transaction_id',
  defaultValue: '',
  operator: 'equals'
})

// تنظیمات فیلتر
const filterSettings = ref({
  maxFilters: 10,
  showSuggestions: 'onFocus',
  autoSave: 'enabled',
  showResultCount: 'always',
  rememberFilters: true,
  highlightResults: true,
  animateFilters: true
})

// فیلترهای ذخیره شده
const savedPresets = ref([
  {
    id: 1,
    name: 'تراکنش‌های موفق',
    description: 'فیلتر تراکنش‌های موفق در هفته گذشته',
    filterCount: 3,
    lastUsed: '۲ ساعت پیش'
  },
  {
    id: 2,
    name: 'تراکنش‌های با مبلغ بالا',
    description: 'تراکنش‌های بالای ۵ میلیون تومان',
    filterCount: 2,
    lastUsed: '۱ روز پیش'
  },
  {
    id: 3,
    name: 'خطاهای اتصال',
    description: 'فیلتر خطاهای اتصال به درگاه‌ها',
    filterCount: 4,
    lastUsed: '۳ روز پیش'
  }
])

// عملیات
const removeFilter = (filterId: number) => {
  activeFilters.value = activeFilters.value.filter(f => f.id !== filterId)
}

const clearAllFilters = () => {
  activeFilters.value = []
}

const addFilter = (filterType: any) => {
  console.log('اضافه کردن فیلتر:', filterType)
}

const resetCustomFilter = () => {
  customFilter.value = {
    name: '',
    type: 'text',
    field: 'transaction_id',
    defaultValue: '',
    operator: 'equals'
  }
}

const createCustomFilter = () => {
  console.log('ایجاد فیلتر سفارشی:', customFilter.value)
}

const saveFilterPreset = () => {
  console.log('ذخیره فیلتر پیش‌فرض')
}

const applyPreset = (preset: any) => {
  console.log('اعمال فیلتر پیش‌فرض:', preset)
}

const editPreset = (preset: any) => {
  console.log('ویرایش فیلتر پیش‌فرض:', preset)
}

const deletePreset = (preset: any) => {
  console.log('حذف فیلتر پیش‌فرض:', preset)
}
</script>

<!--
  کامپوننت فیلترهای پیشرفته
  شامل:
  1. فیلترهای فعال
  2. انواع فیلتر
  3. فیلترهای سفارشی
  4. فیلترهای ذخیره شده
  5. تنظیمات فیلتر
  6. طراحی مدرن و کاملاً ریسپانسیو
--> 
