<template>
  <div v-if="hasAccess" class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-6">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">تنظیمات درگاه‌های پرداخت</h3>
        <p class="text-sm text-gray-500 mt-1">Payment Gateway Settings</p>
      </div>
      <button 
        class="inline-flex items-center px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 transition-colors"
        @click="addGateway"
      >
        <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
        </svg>
        افزودن درگاه
      </button>
    </div>

    <div class="space-y-4">
      <!-- زرین‌پال -->
      <div class="p-6 border border-gray-200 rounded-lg">
        <div class="flex items-center justify-between mb-4">
          <div class="flex items-center">
            <div class="w-10 h-10 bg-green-500 rounded-lg flex items-center justify-center ml-3">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"></path>
              </svg>
            </div>
            <div>
              <h4 class="font-medium text-gray-900">زرین‌پال</h4>
              <p class="text-sm text-gray-500">ZarinPal Gateway</p>
            </div>
          </div>
          <div class="flex items-center space-x-2 space-x-reverse">
            <span class="inline-flex items-center px-2 py-1 text-xs font-semibold rounded-full bg-green-100 text-green-800">
              فعال
            </span>
            <button class="text-gray-400 hover:text-gray-600">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z"></path>
              </svg>
            </button>
          </div>
        </div>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Merchant ID</label>
            <input type="text" value="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm" readonly />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">وضعیت اتصال</label>
            <div class="flex items-center">
              <div class="w-2 h-2 bg-green-500 rounded-full ml-2"></div>
              <span class="text-sm text-green-600">متصل</span>
            </div>
          </div>
        </div>
      </div>

      <!-- نکست‌پی -->
      <div class="p-6 border border-gray-200 rounded-lg">
        <div class="flex items-center justify-between mb-4">
          <div class="flex items-center">
            <div class="w-10 h-10 bg-blue-500 rounded-lg flex items-center justify-center ml-3">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
              </svg>
            </div>
            <div>
              <h4 class="font-medium text-gray-900">نکست‌پی</h4>
              <p class="text-sm text-gray-500">NextPay Gateway</p>
            </div>
          </div>
          <div class="flex items-center space-x-2 space-x-reverse">
            <span class="inline-flex items-center px-2 py-1 text-xs font-semibold rounded-full bg-green-100 text-green-800">
              فعال
            </span>
            <button class="text-gray-400 hover:text-gray-600">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z"></path>
              </svg>
            </button>
          </div>
        </div>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">API Key</label>
            <input type="text" value="np_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm" readonly />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">وضعیت اتصال</label>
            <div class="flex items-center">
              <div class="w-2 h-2 bg-green-500 rounded-full ml-2"></div>
              <span class="text-sm text-green-600">متصل</span>
            </div>
          </div>
        </div>
      </div>

      <!-- آیدی‌پی -->
      <div class="p-6 border border-gray-200 rounded-lg">
        <div class="flex items-center justify-between mb-4">
          <div class="flex items-center">
            <div class="w-10 h-10 bg-purple-500 rounded-lg flex items-center justify-center ml-3">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"></path>
              </svg>
            </div>
            <div>
              <h4 class="font-medium text-gray-900">آیدی‌پی</h4>
              <p class="text-sm text-gray-500">IDPay Gateway</p>
            </div>
          </div>
          <div class="flex items-center space-x-2 space-x-reverse">
            <span class="inline-flex items-center px-2 py-1 text-xs font-semibold rounded-full bg-red-100 text-red-800">
              غیرفعال
            </span>
            <button class="text-gray-400 hover:text-gray-600">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z"></path>
              </svg>
            </button>
          </div>
        </div>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">API Key</label>
            <input type="text" placeholder="کلید API را وارد کنید" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">وضعیت اتصال</label>
            <div class="flex items-center">
              <div class="w-2 h-2 bg-red-500 rounded-full ml-2"></div>
              <span class="text-sm text-red-600">قطع</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات عمومی -->
    <div class="mt-6 pt-4 border-t border-gray-200">
      <h4 class="font-medium text-gray-900 mb-3">تنظیمات عمومی</h4>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">درگاه پیش‌فرض</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
            <option value="zarinpal">زرین‌پال</option>
            <option value="nextpay">نکست‌پی</option>
            <option value="idpay">آیدی‌پی</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر مبلغ تراکنش</label>
          <input type="number" value="50000000" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500" />
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
</script>

<script setup lang="ts">
import { computed, onMounted, watch } from 'vue';
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

const addGateway = () => {
  // TODO: باز کردن مودال افزودن درگاه جدید

}
</script>

<!--
  کامپوننت تنظیمات درگاه‌های پرداخت
  - مدیریت درگاه‌های فعال/غیرفعال
  - تنظیمات API هر درگاه
  - وضعیت اتصال
  - تنظیمات عمومی
  - کاملاً ریسپانسیو
  توضیحات کامل به فارسی در کد
--> 
