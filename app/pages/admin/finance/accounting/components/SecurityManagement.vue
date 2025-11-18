<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200">
    <!-- هدر بخش -->
    <div class="px-6 py-4 border-b border-gray-200">
      <div class="flex items-center justify-between">
        <div>
          <h3 class="text-lg font-semibold text-gray-900">مدیریت امنیت</h3>
          <p class="text-sm text-gray-600 mt-1">تنظیمات امنیتی، کلیدهای API و رمزگذاری</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button
            @click="generateNewKey"
            class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
          >
            <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 7a2 2 0 012 2m4 0a6 6 0 01-7.743 5.743L11 17H9v2H7v2H4a1 1 0 01-1-1v-2.586a1 1 0 01.293-.707l5.964-5.964A6 6 0 1121 9z" />
            </svg>
            تولید کلید جدید
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
      <div v-if="activeTab === 'security'">
        <SecuritySettings />
      </div>
      <div v-if="activeTab === 'api'">
        <ApiKeyManagement />
      </div>
      <div v-if="activeTab === 'encryption'">
        <EncryptionSettings />
      </div>
      <div v-if="activeTab === 'permissions'">
        <PermissionManagement />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
// @ts-ignore
import SecuritySettings from './security/SecuritySettings.vue'
// @ts-ignore
import ApiKeyManagement from './security/ApiKeyManagement.vue'
// @ts-ignore
import EncryptionSettings from './security/EncryptionSettings.vue'
// @ts-ignore
import PermissionManagement from './security/PermissionManagement.vue'

const tabs = ref([
  { id: 'security', name: 'تنظیمات امنیتی' },
  { id: 'api', name: 'مدیریت کلیدهای API' },
  { id: 'encryption', name: 'تنظیمات رمزگذاری' },
  { id: 'permissions', name: 'مدیریت مجوزها' }
])
const activeTab = ref('security')

// تولید کلید جدید
const generateNewKey = () => {
  console.log('تولید کلید جدید')
}
</script>

<!--
  کامپوننت اصلی مدیریت امنیت
  شامل:
  1. تنظیمات امنیتی
  2. مدیریت کلیدهای API
  3. تنظیمات رمزگذاری
  4. مدیریت مجوزهای دسترسی
  5. طراحی مدرن و کاملاً ریسپانسیو
  توضیحات کامل به فارسی
--> 
