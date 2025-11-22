<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200">
    <!-- هدر بخش -->
    <div class="px-6 py-4 border-b border-gray-200">
      <div class="flex items-center justify-between">
        <div>
          <h3 class="text-lg font-semibold text-gray-900">تنظیمات پیشرفته</h3>
          <p class="text-sm text-gray-600 mt-1">تنظیمات پیشرفته امنیتی و مدیریتی اتصال حسابداری</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button
            :disabled="isSaving"
            class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
            @click="saveAllSettings"
          >
            <svg v-if="isSaving" class="w-4 h-4 ml-2 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
            </svg>
            <svg v-else class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
            </svg>
            {{ isSaving ? 'در حال ذخیره...' : 'ذخیره همه تنظیمات' }}
          </button>
          <button
            class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-500"
            @click="resetToDefaults"
          >
            <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
            </svg>
            بازنشانی به پیش‌فرض
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
      <!-- تب تنظیمات امنیتی -->
      <div v-if="activeTab === 'security'" class="space-y-6">
        <SecuritySettings />
      </div>

      <!-- تب مدیریت کلیدهای API -->
      <div v-if="activeTab === 'api-keys'" class="space-y-6">
        <ApiKeyManagement />
      </div>

      <!-- تب تنظیمات رمزگذاری -->
      <div v-if="activeTab === 'encryption'" class="space-y-6">
        <EncryptionSettings />
      </div>

      <!-- تب مدیریت مجوزهای دسترسی -->
      <div v-if="activeTab === 'permissions'" class="space-y-6">
        <PermissionManagement />
      </div>

      <!-- تب تنظیمات پشتیبان‌گیری -->
      <div v-if="activeTab === 'backup'" class="space-y-6">
        <BackupSettings />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
</script>

<script setup lang="ts">
import { ref } from 'vue';

definePageMeta({ layout: 'admin-main', middleware: 'admin' });

// Import کامپوننت‌های زیربخش‌ها
// @ts-ignore
import SecuritySettings from './settings/SecuritySettings.vue'
// @ts-ignore
import ApiKeyManagement from './settings/ApiKeyManagement.vue'
// @ts-ignore
import EncryptionSettings from './settings/EncryptionSettings.vue'
// @ts-ignore
import PermissionManagement from './settings/PermissionManagement.vue'
// @ts-ignore
import BackupSettings from './settings/BackupSettings.vue'

// تعریف تب‌های مختلف
const tabs = ref([
  { id: 'security', name: 'تنظیمات امنیتی' },
  { id: 'api-keys', name: 'مدیریت کلیدهای API' },
  { id: 'encryption', name: 'تنظیمات رمزگذاری' },
  { id: 'permissions', name: 'مدیریت مجوزها' },
  { id: 'backup', name: 'تنظیمات پشتیبان‌گیری' }
])

// تب فعال
const activeTab = ref('security')

// وضعیت ذخیره
const isSaving = ref(false)

// ذخیره همه تنظیمات
const saveAllSettings = async () => {
  try {
    isSaving.value = true
    // TODO: ذخیره همه تنظیمات

    // شبیه‌سازی ذخیره
    await new Promise(resolve => setTimeout(resolve, 2000))

  } catch (error) {
    console.error('خطا در ذخیره تنظیمات:', error)
  } finally {
    isSaving.value = false
  }
}

// بازنشانی به تنظیمات پیش‌فرض
const resetToDefaults = () => {
  // TODO: بازنشانی به تنظیمات پیش‌فرض

}
</script>

<!--
  کامپوننت تنظیمات پیشرفته
  شامل ۵ بخش اصلی:
  1. تنظیمات امنیتی
  2. مدیریت کلیدهای API
  3. تنظیمات رمزگذاری
  4. مدیریت مجوزهای دسترسی
  5. تنظیمات پشتیبان‌گیری
  طراحی مدرن با تب‌بندی
  کاملاً ریسپانسیو
--> 
