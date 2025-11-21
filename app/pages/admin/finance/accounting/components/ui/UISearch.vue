<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">جستجوی سریع</h4>
        <p class="text-sm text-gray-600 mt-1">مدیریت و تنظیم جستجوی پیشرفته</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="clearSearchHistory"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
          </svg>
          پاک کردن تاریخچه
        </button>
      </div>
    </div>

    <!-- جستجوی اصلی -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">جستجوی اصلی</h5>
      <div class="relative">
        <div class="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
          <svg class="h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
        </div>
        <input
          v-model="searchQuery"
          type="text"
          class="block w-full pr-10 pl-3 py-3 border border-gray-300 rounded-md leading-5 bg-white placeholder-gray-500 focus:outline-none focus:placeholder-gray-400 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          placeholder="جستجو در تراکنش‌ها، اتصالات، خطاها..."
          @input="handleSearch"
          @focus="showSuggestions = true"
        />
      </div>

      <!-- پیشنهادات جستجو -->
      <div v-if="showSuggestions && searchSuggestions.length > 0" class="mt-2 bg-white border border-gray-200 rounded-md shadow-lg">
        <div class="py-2">
          <div
            v-for="suggestion in searchSuggestions"
            :key="suggestion.id"
            class="px-4 py-2 hover:bg-gray-100 cursor-pointer flex items-center"
            @click="selectSuggestion(suggestion)"
          >
            <svg class="w-4 h-4 ml-3 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="suggestion.icon" />
            </svg>
            <div>
              <div class="text-sm font-medium text-gray-900">{{ suggestion.title }}</div>
              <div class="text-xs text-gray-500">{{ suggestion.description }}</div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- فیلترهای جستجو -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">فیلترهای جستجو</h5>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نوع محتوا</label>
          <div class="space-y-2">
            <label class="flex items-center">
              <input
                v-model="searchFilters.contentTypes"
                value="transactions"
                type="checkbox"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <span class="mr-3 text-sm text-gray-700">تراکنش‌ها</span>
            </label>
            <label class="flex items-center">
              <input
                v-model="searchFilters.contentTypes"
                value="connections"
                type="checkbox"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <span class="mr-3 text-sm text-gray-700">اتصالات</span>
            </label>
            <label class="flex items-center">
              <input
                v-model="searchFilters.contentTypes"
                value="errors"
                type="checkbox"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <span class="mr-3 text-sm text-gray-700">خطاها</span>
            </label>
            <label class="flex items-center">
              <input
                v-model="searchFilters.contentTypes"
                value="logs"
                type="checkbox"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <span class="mr-3 text-sm text-gray-700">لاگ‌ها</span>
            </label>
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">بازه زمانی</label>
          <select
            v-model="searchFilters.timeRange"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="all">همه زمان‌ها</option>
            <option value="1hour">۱ ساعت گذشته</option>
            <option value="24hours">۲۴ ساعت گذشته</option>
            <option value="7days">۷ روز گذشته</option>
            <option value="30days">۳۰ روز گذشته</option>
            <option value="custom">سفارشی</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">مرتب‌سازی</label>
          <select
            v-model="searchFilters.sortBy"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="relevance">مرتبط‌ترین</option>
            <option value="date_desc">جدیدترین</option>
            <option value="date_asc">قدیمی‌ترین</option>
            <option value="name_asc">نام (الف-ی)</option>
            <option value="name_desc">نام (ی-الف)</option>
          </select>
        </div>
      </div>
    </div>

    <!-- تاریخچه جستجو -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">تاریخچه جستجو</h5>
      </div>
      <div class="p-6">
        <div v-if="searchHistory.length === 0" class="text-center py-8">
          <svg class="w-12 h-12 mx-auto text-gray-400 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          <p class="text-gray-500">تاریخچه جستجویی وجود ندارد</p>
        </div>
        <div v-else class="space-y-3">
          <div
            v-for="item in searchHistory"
            :key="item.id"
            class="flex items-center justify-between p-3 border border-gray-200 rounded-lg hover:bg-gray-50"
          >
            <div class="flex items-center">
              <svg class="w-4 h-4 ml-3 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
              <div>
                <div class="text-sm font-medium text-gray-900">{{ item.query }}</div>
                <div class="text-xs text-gray-500">{{ item.timestamp }} - {{ item.resultCount }} نتیجه</div>
              </div>
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <button
                class="text-blue-600 hover:text-blue-800"
                @click="repeatSearch(item)"
              >
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                </svg>
              </button>
              <button
                class="text-red-600 hover:text-red-800"
                @click="removeFromHistory(item.id)"
              >
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                </svg>
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- جستجوهای محبوب -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">جستجوهای محبوب</h5>
      <div class="flex flex-wrap gap-2">
        <button
          v-for="popular in popularSearches"
          :key="popular.id"
          class="inline-flex items-center px-3 py-1 rounded-full text-sm bg-blue-100 text-blue-800 hover:bg-blue-200 transition-colors"
          @click="selectPopularSearch(popular)"
        >
          <svg class="w-3 h-3 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
          {{ popular.query }}
        </button>
      </div>
    </div>

    <!-- تنظیمات جستجو -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">تنظیمات جستجو</h5>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر نتایج</label>
          <input
            v-model="searchSettings.maxResults"
            type="number"
            min="10"
            max="1000"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تاخیر جستجو</label>
          <select
            v-model="searchSettings.searchDelay"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="0">بدون تاخیر</option>
            <option value="300">۳۰۰ میلی‌ثانیه</option>
            <option value="500">۵۰۰ میلی‌ثانیه</option>
            <option value="1000">۱ ثانیه</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نمایش پیشنهادات</label>
          <select
            v-model="searchSettings.showSuggestions"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="always">همیشه</option>
            <option value="onFocus">در زمان فوکوس</option>
            <option value="onInput">در زمان تایپ</option>
            <option value="never">هرگز</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">ذخیره تاریخچه</label>
          <select
            v-model="searchSettings.saveHistory"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="all">همه جستجوها</option>
            <option value="successful">فقط جستجوهای موفق</option>
            <option value="none">هیچ‌کدام</option>
          </select>
        </div>
      </div>

      <div class="mt-4 space-y-3">
        <label class="flex items-center">
          <input
            v-model="searchSettings.fuzzySearch"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">جستجوی فازی (جستجوی مشابه)</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="searchSettings.highlightResults"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">برجسته کردن نتایج</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="searchSettings.autoComplete"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">تکمیل خودکار</span>
        </label>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// جستجوی اصلی
const searchQuery = ref('')
const showSuggestions = ref(false)

// فیلترهای جستجو
const searchFilters = ref({
  contentTypes: ['transactions', 'connections', 'errors'],
  timeRange: 'all',
  sortBy: 'relevance'
})

// تنظیمات جستجو
const searchSettings = ref({
  maxResults: 100,
  searchDelay: 300,
  showSuggestions: 'onFocus',
  saveHistory: 'all',
  fuzzySearch: true,
  highlightResults: true,
  autoComplete: true
})

// پیشنهادات جستجو
const searchSuggestions = ref([
  {
    id: 1,
    title: 'تراکنش‌های موفق',
    description: 'جستجو در تراکنش‌های موفق',
    icon: 'M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z'
  },
  {
    id: 2,
    title: 'خطاهای اتصال',
    description: 'جستجو در خطاهای اتصال',
    icon: 'M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z'
  },
  {
    id: 3,
    title: 'درگاه زرین‌پال',
    description: 'جستجو در تراکنش‌های زرین‌پال',
    icon: 'M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z'
  }
])

// تاریخچه جستجو
const searchHistory = ref([
  {
    id: 1,
    query: 'تراکنش‌های موفق',
    timestamp: '۲ ساعت پیش',
    resultCount: 45
  },
  {
    id: 2,
    query: 'خطاهای اتصال',
    timestamp: '۱ روز پیش',
    resultCount: 12
  },
  {
    id: 3,
    query: 'درگاه زرین‌پال',
    timestamp: '۳ روز پیش',
    resultCount: 78
  }
])

// جستجوهای محبوب
const popularSearches = ref([
  { id: 1, query: 'تراکنش‌های موفق' },
  { id: 2, query: 'خطاهای اتصال' },
  { id: 3, query: 'درگاه زرین‌پال' },
  { id: 4, query: 'تراکنش‌های بالای ۱ میلیون' },
  { id: 5, query: 'خطاهای هفته گذشته' },
  { id: 6, query: 'اتصالات غیرفعال' }
])

// عملیات
const handleSearch = () => {

}

interface Suggestion {
  id?: number | string
  title?: string
  [key: string]: unknown
}

interface SearchItem {
  id?: number | string
  query?: string
  [key: string]: unknown
}

interface PopularSearch {
  id?: number | string
  query?: string
  [key: string]: unknown
}

const selectSuggestion = (suggestion: Suggestion) => {
  searchQuery.value = suggestion.title || ''
  showSuggestions.value = false

}

const repeatSearch = (item: SearchItem) => {
  searchQuery.value = item.query || ''

}

const removeFromHistory = (id: number) => {
  searchHistory.value = searchHistory.value.filter(item => item.id !== id)
}

const selectPopularSearch = (popular: PopularSearch) => {
  searchQuery.value = popular.query || ''

}

const clearSearchHistory = () => {
  searchHistory.value = []

}
</script>

<!--
  کامپوننت جستجوی سریع
  شامل:
  1. جستجوی اصلی
  2. فیلترهای جستجو
  3. تاریخچه جستجو
  4. جستجوهای محبوب
  5. تنظیمات جستجو
  6. طراحی مدرن و کاملاً ریسپانسیو
--> 
