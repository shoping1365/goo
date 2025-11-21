<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200">
    <!-- هدر بخش -->
    <div class="px-6 py-4 border-b border-gray-200">
      <div class="flex items-center justify-between">
        <div>
          <h3 class="text-lg font-semibold text-gray-900">مدیریت نسخه‌ها</h3>
          <p class="text-sm text-gray-600 mt-1">نگهداری، بازگشت و مقایسه نسخه‌های مختلف اتصال حسابداری</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button
            class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
            @click="createBackup"
          >
            <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
            </svg>
            ایجاد نسخه پشتیبان
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
          :class="[
            activeTab === tab.id
              ? 'border-blue-500 text-blue-600'
              : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300',
            'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm'
          ]"
          @click="activeTab = tab.id"
        >
          {{ tab.name }}
        </button>
      </nav>
    </div>

    <!-- محتوای تب‌ها -->
    <div class="p-6">
      <div v-if="activeTab === 'versions'">
        <VersionList />
      </div>
      <div v-if="activeTab === 'rollback'">
        <VersionRollback />
      </div>
      <div v-if="activeTab === 'compare'">
        <VersionCompare />
      </div>
      <div v-if="activeTab === 'archive'">
        <VersionArchive />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
// @ts-ignore
import VersionList from './VersionList.vue'
// @ts-ignore
import VersionRollback from './VersionRollback.vue'
// @ts-ignore
import VersionCompare from './VersionCompare.vue'
// @ts-ignore
import VersionArchive from './VersionArchive.vue'

const tabs = ref([
  { id: 'versions', name: 'نگهداری نسخه‌ها' },
  { id: 'rollback', name: 'بازگشت به نسخه قبلی' },
  { id: 'compare', name: 'مقایسه نسخه‌ها' },
  { id: 'archive', name: 'آرشیو نسخه‌های قدیمی' }
])
const activeTab = ref('versions')

// ایجاد نسخه پشتیبان
const createBackup = () => {

}
</script>

<!--
  کامپوننت اصلی مدیریت نسخه‌ها
  شامل:
  1. نگهداری نسخه‌های مختلف اتصال
  2. بازگشت به نسخه قبلی
  3. مقایسه نسخه‌ها
  4. آرشیو نسخه‌های قدیمی
  طراحی مدرن و کاملاً ریسپانسیو
  توضیحات کامل به فارسی
--> 
