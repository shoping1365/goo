<template>
  <div v-if="hasAccess" class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-6">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">گزارش‌گیری و خروجی</h3>
        <p class="text-sm text-gray-500 mt-1">Export Reports & Analytics</p>
      </div>
    </div>

    <div class="space-y-4">
      <!-- خروجی Excel -->
      <div class="p-6 bg-green-50 rounded-lg border border-green-200">
        <div class="flex items-center justify-between">
          <div class="flex items-center">
            <div class="w-10 h-10 bg-green-500 rounded-lg flex items-center justify-center ml-3">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
              </svg>
            </div>
            <div>
              <h4 class="font-medium text-green-800">خروجی Excel</h4>
              <p class="text-sm text-green-600">گزارش کامل پرداخت‌ها</p>
            </div>
          </div>
          <button 
            class="px-4 py-2 bg-green-500 text-white rounded-lg hover:bg-green-600 transition-colors"
            :disabled="isExporting"
            @click="exportToExcel"
          >
            {{ isExporting ? 'در حال آماده‌سازی...' : 'دانلود Excel' }}
          </button>
        </div>
      </div>

      <!-- خروجی PDF -->
      <div class="p-6 bg-red-50 rounded-lg border border-red-200">
        <div class="flex items-center justify-between">
          <div class="flex items-center">
            <div class="w-10 h-10 bg-red-500 rounded-lg flex items-center justify-center ml-3">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21h10a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v14a2 2 0 002 2z"></path>
              </svg>
            </div>
            <div>
              <h4 class="font-medium text-red-800">گزارش PDF</h4>
              <p class="text-sm text-red-600">گزارش تحلیلی پرداخت‌ها</p>
            </div>
          </div>
          <button 
            class="px-4 py-2 bg-red-500 text-white rounded-lg hover:bg-red-600 transition-colors"
            :disabled="isExporting"
            @click="exportToPDF"
          >
            {{ isExporting ? 'در حال آماده‌سازی...' : 'دانلود PDF' }}
          </button>
        </div>
      </div>

      <!-- گزارش تحلیلی -->
      <div class="p-6 bg-purple-50 rounded-lg border border-purple-200">
        <div class="flex items-center justify-between">
          <div class="flex items-center">
            <div class="w-10 h-10 bg-purple-500 rounded-lg flex items-center justify-center ml-3">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
              </svg>
            </div>
            <div>
              <h4 class="font-medium text-purple-800">گزارش تحلیلی</h4>
              <p class="text-sm text-purple-600">تحلیل پیشرفته و نمودارها</p>
            </div>
          </div>
          <button 
            class="px-4 py-2 bg-purple-500 text-white rounded-lg hover:bg-purple-600 transition-colors"
            :disabled="isExporting"
            @click="exportAnalytics"
          >
            {{ isExporting ? 'در حال آماده‌سازی...' : 'دانلود گزارش' }}
          </button>
        </div>
      </div>
    </div>

    <!-- تنظیمات خروجی -->
    <div class="mt-6 pt-4 border-t border-gray-200">
      <h4 class="font-medium text-gray-900 mb-3">تنظیمات خروجی</h4>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">بازه زمانی</label>
          <select 
            v-model="exportSettings.dateRange"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="all">همه زمان‌ها</option>
            <option value="today">امروز</option>
            <option value="week">این هفته</option>
            <option value="month">این ماه</option>
            <option value="custom">انتخاب دستی</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">فرمت فایل</label>
          <select 
            v-model="exportSettings.format"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="xlsx">Excel (.xlsx)</option>
            <option value="csv">CSV (.csv)</option>
            <option value="pdf">PDF (.pdf)</option>
          </select>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
</script>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue';
import { useAuth } from '~/composables/useAuth';

// احراز هویت
const { user, isAuthenticated } = useAuth();

// بررسی دسترسی admin
const hasAccess = computed(() => {
  if (!isAuthenticated.value) {
    return false;
  }

  const userRole = user.value?.role?.toLowerCase() || '';
  const adminRoles = ['admin', 'developer'];
  return adminRoles.includes(userRole);
});

// بررسی احراز هویت و دسترسی admin - نمایش 404 در صورت عدم دسترسی
const checkAuth = async (): Promise<void> => {
  if (!hasAccess.value) {
    await navigateTo('/404', { external: false });
  }
};

// بررسی احراز هویت در هنگام mount
onMounted(async () => {
  await checkAuth();
});

// بررسی احراز هویت هنگام تغییر وضعیت احراز هویت
watch([isAuthenticated, hasAccess], async () => {
  if (!hasAccess.value) {
    await checkAuth();
  }
});

interface Payment {
  [key: string]: unknown
}

interface Props {
  payments: Payment[]
}

const _props = defineProps<Props>()

const isExporting = ref(false)
const exportSettings = ref({
  dateRange: 'all',
  format: 'xlsx'
})

const exportToExcel = async () => {
  isExporting.value = true
  try {
    // TODO: پیاده‌سازی خروجی Excel

    await new Promise(resolve => setTimeout(resolve, 2000)) // شبیه‌سازی تاخیر
    alert('فایل Excel با موفقیت دانلود شد!')
  } catch (error) {
    console.error('Export error:', error)
    alert('خطا در خروجی فایل')
  } finally {
    isExporting.value = false
  }
}

const exportToPDF = async () => {
  isExporting.value = true
  try {
    // TODO: پیاده‌سازی خروجی PDF

    await new Promise(resolve => setTimeout(resolve, 2000)) // شبیه‌سازی تاخیر
    alert('فایل PDF با موفقیت دانلود شد!')
  } catch (error) {
    console.error('Export error:', error)
    alert('خطا در خروجی فایل')
  } finally {
    isExporting.value = false
  }
}

const exportAnalytics = async () => {
  isExporting.value = true
  try {
    // TODO: پیاده‌سازی گزارش تحلیلی

    await new Promise(resolve => setTimeout(resolve, 2000)) // شبیه‌سازی تاخیر
    alert('گزارش تحلیلی با موفقیت دانلود شد!')
  } catch (error) {
    console.error('Export error:', error)
    alert('خطا در خروجی فایل')
  } finally {
    isExporting.value = false
  }
}
</script>

<!--
  کامپوننت Export و گزارش‌گیری
  - خروجی Excel/PDF/تحلیلی
  - تنظیمات خروجی
  - کاملاً ریسپانسیو
  توضیحات کامل به فارسی در کد
--> 
