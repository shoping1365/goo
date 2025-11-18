<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200">
    <!-- هدر بخش -->
    <div class="px-6 py-4 border-b border-gray-200">
      <div class="flex items-center justify-between">
        <div>
          <h3 class="text-lg font-semibold text-gray-900">رابط کاربری</h3>
          <p class="text-sm text-gray-600 mt-1">بهبود تجربه کاربری و رابط‌های تعاملی</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button
            @click="saveUISettings"
            class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
            </svg>
            ذخیره تنظیمات
          </button>
        </div>
      </div>
    </div>

    <!-- تب‌های بخش‌های مختلف -->
    <div class="border-b border-gray-200">
      <nav class="-mb-px flex space-x-8 space-x-reverse px-6">
        <button
          v-for="tab in tabs"
          :key="tab.id"
          @click="activeTab = tab.id"
          :class="[
            activeTab === tab.id
              ? 'border-blue-500 text-blue-600'
              : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300',
            'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm'
          ]"
        >
          {{ tab.name }}
        </button>
      </nav>
    </div>

    <!-- محتوای تب‌ها -->
    <div class="p-6">
      <div v-if="activeTab === 'design'">
        <UIDesign />
      </div>
      <div v-if="activeTab === 'charts'">
        <UICharts />
      </div>
      <div v-if="activeTab === 'filters'">
        <UIFilters />
      </div>
      <div v-if="activeTab === 'search'">
        <UISearch />
      </div>
      <div v-if="activeTab === 'status'">
        <UIStatus />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
// @ts-ignore
import UIDesign from './UIDesign.vue'
// @ts-ignore
import UICharts from './UICharts.vue'
// @ts-ignore
import UIFilters from './UIFilters.vue'
// @ts-ignore
import UISearch from './UISearch.vue'
// @ts-ignore
import UIStatus from './UIStatus.vue'

const tabs = ref([
  { id: 'design', name: 'طراحی مدرن' },
  { id: 'charts', name: 'نمودارهای تعاملی' },
  { id: 'filters', name: 'فیلترهای پیشرفته' },
  { id: 'search', name: 'جستجوی سریع' },
  { id: 'status', name: 'نمایش وضعیت زنده' }
])
const activeTab = ref('design')

// ذخیره تنظیمات UI
const saveUISettings = () => {
  console.log('ذخیره تنظیمات رابط کاربری')
}
</script>

<!--
  کامپوننت اصلی رابط کاربری
  شامل:
  1. طراحی مدرن و ریسپانسیو
  2. نمودارهای تعاملی
  3. فیلترهای پیشرفته
  4. جستجوی سریع
  5. نمایش وضعیت زنده
  طراحی مدرن و کاملاً ریسپانسیو
  توضیحات کامل به فارسی
--> 
