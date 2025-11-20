<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200">
    <!-- هدر بخش -->
    <div class="px-6 py-4 border-b border-gray-200">
      <div class="flex items-center justify-between">
        <div>
          <h3 class="text-lg font-semibold text-gray-900">مدیریت داده‌ها</h3>
          <p class="text-sm text-gray-600 mt-1">پشتیبان‌گیری، بازیابی و مدیریت داده‌های حسابداری</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button
            class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
            @click="createBackup"
          >
            <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7H5a2 2 0 00-2 2v9a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-3m-1 4l-3 3m0 0l-3-3m3 3V4" />
            </svg>
            ایجاد پشتیبان
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
      <div v-if="activeTab === 'backup'">
        <BackupRestore />
      </div>
      <div v-if="activeTab === 'cleanup'">
        <DataCleanup />
      </div>
      <div v-if="activeTab === 'archive'">
        <DataArchive />
      </div>
      <div v-if="activeTab === 'migration'">
        <DataMigration />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
// @ts-ignore
import BackupRestore from './BackupRestore.vue'
// @ts-ignore
import DataCleanup from './DataCleanup.vue'
// @ts-ignore
import DataArchive from './DataArchive.vue'
// @ts-ignore
import DataMigration from './DataMigration.vue'

const tabs = ref([
  { id: 'backup', name: 'پشتیبان‌گیری و بازیابی' },
  { id: 'cleanup', name: 'پاکسازی داده‌ها' },
  { id: 'archive', name: 'آرشیو داده‌ها' },
  { id: 'migration', name: 'انتقال داده‌ها' }
])
const activeTab = ref('backup')

// ایجاد پشتیبان
const createBackup = () => {
  console.log('ایجاد پشتیبان')
}
</script>

<!--
  کامپوننت اصلی مدیریت داده‌ها
  شامل:
  1. پشتیبان‌گیری و بازیابی
  2. پاکسازی داده‌ها
  3. آرشیو داده‌ها
  4. انتقال داده‌ها
  5. طراحی مدرن و کاملاً ریسپانسیو
  توضیحات کامل به فارسی
--> 
