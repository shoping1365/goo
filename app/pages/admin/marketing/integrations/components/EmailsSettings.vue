<template>
  <div v-if="hasAccess" class="bg-white rounded-lg shadow-md p-6">
    <h3 class="text-lg font-semibold text-gray-800 mb-4">تنظیمات سرویس Emails</h3>
    
    <!-- تنظیمات ایمیل -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
      <div class="space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">SMTP Server</label>
          <input type="text" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500" value="smtp.gmail.com">
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">Port</label>
          <input type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500" value="587">
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">Username</label>
          <input type="email" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500" value="noreply@example.com">
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">Password</label>
          <input type="password" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500" value="********">
        </div>
      </div>

      <div class="space-y-4">
        <div class="flex items-center justify-between">
          <div>
            <h5 class="font-medium text-gray-900">SSL/TLS</h5>
            <p class="text-sm text-gray-600">فعال‌سازی اتصال امن</p>
          </div>
          <label class="relative inline-flex items-center cursor-pointer">
            <input type="checkbox" checked class="sr-only peer">
            <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
          </label>
        </div>
        
        <div class="flex items-center justify-between">
          <div>
            <h5 class="font-medium text-gray-900">تست اتصال</h5>
            <p class="text-sm text-gray-600">آزمایش تنظیمات SMTP</p>
          </div>
          <button class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors">
            تست
          </button>
        </div>

        <div class="flex items-center justify-between">
          <div>
            <h5 class="font-medium text-gray-900">ذخیره تنظیمات</h5>
            <p class="text-sm text-gray-600">ذخیره تغییرات</p>
          </div>
          <button class="px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700 transition-colors">
            ذخیره
          </button>
        </div>
      </div>
    </div>

    <!-- دکمه‌های عملیات -->
    <div class="flex justify-end space-x-3 space-x-reverse">
      <button class="px-4 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400 transition-colors">
        بازنشانی
      </button>
      <button class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors">
        بروزرسانی
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, watch } from 'vue';
import { useAuth } from '~/composables/useAuth';

declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>;

// احراز هویت
const { user, isAuthenticated } = useAuth();

// بررسی دسترسی admin
const hasAccess = computed(() => {
  if (!isAuthenticated.value) {
    return false;
  }

  const userRole = user.value?.role?.toLowerCase() || '';
  const adminRoles = ['admin', 'developer', 'super_admin', 'manager', 'operator'];
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
// کامپوننت تنظیمات سرویس Emails
// این کامپوننت تنظیمات SMTP و ایمیل را فراهم می‌کند
</script> 
