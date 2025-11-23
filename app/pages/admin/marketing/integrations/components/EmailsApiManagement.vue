<template>
  <div v-if="hasAccess" class="bg-white rounded-lg shadow-md p-6">
    <h3 class="text-lg font-semibold text-gray-800 mb-4">مدیریت API سرویس Emails</h3>
    
    <!-- آمار API -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
      <div class="bg-blue-50 p-6 rounded-lg">
        <div class="text-sm text-blue-600">کل درخواست‌ها</div>
        <div class="text-2xl font-bold text-blue-800">۱۸,۷۲۰</div>
      </div>
      <div class="bg-green-50 p-6 rounded-lg">
        <div class="text-sm text-green-600">درخواست‌های موفق</div>
        <div class="text-2xl font-bold text-green-800">۱۷,۸۹۰</div>
      </div>
      <div class="bg-orange-50 p-6 rounded-lg">
        <div class="text-sm text-orange-600">نرخ خطا</div>
        <div class="text-2xl font-bold text-orange-800">۴.۴٪</div>
      </div>
      <div class="bg-purple-50 p-6 rounded-lg">
        <div class="text-sm text-purple-600">زمان پاسخ متوسط</div>
        <div class="text-2xl font-bold text-purple-800">۱۸۰ms</div>
      </div>
    </div>

    <!-- کلیدهای API -->
    <div class="mb-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">کلیدهای API</h4>
      <div class="space-y-4">
        <div class="border border-gray-200 rounded-lg p-6">
          <div class="flex items-center justify-between mb-3">
            <div>
              <h5 class="font-medium text-gray-900">کلید اصلی</h5>
              <p class="text-sm text-gray-600">کلید API برای دسترسی کامل به سرویس ایمیل</p>
            </div>
            <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-green-100 text-green-800">فعال</span>
          </div>
          <div class="flex space-x-2 space-x-reverse">
            <input type="password" class="flex-1 px-3 py-2 border border-gray-300 rounded-md" value="emails_live_abcdef1234567890" readonly>
            <button class="px-3 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400 transition-colors">کپی</button>
            <button class="px-3 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors">تست</button>
            <button class="px-3 py-2 bg-red-600 text-white rounded-md hover:bg-red-700 transition-colors">بازنشانی</button>
          </div>
        </div>

        <div class="border border-gray-200 rounded-lg p-6">
          <div class="flex items-center justify-between mb-3">
            <div>
              <h5 class="font-medium text-gray-900">کلید تست</h5>
              <p class="text-sm text-gray-600">کلید API برای محیط تست و توسعه</p>
            </div>
            <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-yellow-100 text-yellow-800">تست</span>
          </div>
          <div class="flex space-x-2 space-x-reverse">
            <input type="password" class="flex-1 px-3 py-2 border border-gray-300 rounded-md" value="emails_test_1234567890abcdef" readonly>
            <button class="px-3 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400 transition-colors">کپی</button>
            <button class="px-3 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors">تست</button>
            <button class="px-3 py-2 bg-red-600 text-white rounded-md hover:bg-red-700 transition-colors">بازنشانی</button>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات API -->
    <div class="mb-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">تنظیمات API</h4>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">آدرس پایه API</label>
            <input type="url" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500" value="https://api.emails.com/v1">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">محدودیت درخواست (در دقیقه)</label>
            <input type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500" value="1500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تایم‌اوت (ثانیه)</label>
            <input type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500" value="25">
          </div>
        </div>

        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نسخه API</label>
            <select class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
              <option>v1.0 (پایدار)</option>
              <option>v1.1 (بتا)</option>
              <option>v2.0 (آلفا)</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">فرمت پاسخ</label>
            <select class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
              <option>JSON</option>
              <option>XML</option>
              <option>CSV</option>
            </select>
          </div>
          <div class="flex items-center">
            <input type="checkbox" checked class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded">
            <label class="mr-2 block text-sm text-gray-900">فعال‌سازی CORS</label>
          </div>
        </div>
      </div>
    </div>

    <!-- لاگ‌های API -->
    <div class="mb-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">آخرین درخواست‌های API</h4>
      <div class="bg-gray-50 p-6 rounded-lg max-h-48 overflow-y-auto">
        <div class="space-y-2 text-sm">
          <div class="flex items-center justify-between">
            <div class="flex items-center space-x-2 space-x-reverse">
              <span class="text-green-600">[۱۴:۳۲]</span>
              <span class="text-gray-700">POST /api/v1/emails/send</span>
            </div>
            <span class="text-xs text-green-600">۲۰۱ Created</span>
          </div>
          <div class="flex items-center justify-between">
            <div class="flex items-center space-x-2 space-x-reverse">
              <span class="text-blue-600">[۱۴:۳۱]</span>
              <span class="text-gray-700">GET /api/v1/templates</span>
            </div>
            <span class="text-xs text-green-600">۲۰۰ OK</span>
          </div>
          <div class="flex items-center justify-between">
            <div class="flex items-center space-x-2 space-x-reverse">
              <span class="text-red-600">[۱۴:۳۰]</span>
              <span class="text-gray-700">PUT /api/v1/campaigns/456</span>
            </div>
            <span class="text-xs text-red-600">۴۰۰ Bad Request</span>
          </div>
          <div class="flex items-center justify-between">
            <div class="flex items-center space-x-2 space-x-reverse">
              <span class="text-green-600">[۱۴:۲۹]</span>
              <span class="text-gray-700">DELETE /api/v1/emails/789</span>
            </div>
            <span class="text-xs text-green-600">۲۰۰ OK</span>
          </div>
          <div class="flex items-center justify-between">
            <div class="flex items-center space-x-2 space-x-reverse">
              <span class="text-orange-600">[۱۴:۲۸]</span>
              <span class="text-gray-700">GET /api/v1/analytics</span>
            </div>
            <span class="text-xs text-orange-600">۴۰۳ Forbidden</span>
          </div>
        </div>
      </div>
    </div>

    <!-- مستندات API -->
    <div class="mb-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">مستندات API</h4>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="border border-gray-200 rounded-lg p-6">
          <h5 class="font-medium text-gray-900 mb-2">نقاط پایانی اصلی</h5>
          <div class="space-y-2 text-sm">
            <div class="flex justify-between">
              <span class="text-gray-600">POST /emails/send</span>
              <span class="text-green-600">فعال</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-600">GET /templates</span>
              <span class="text-green-600">فعال</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-600">PUT /campaigns</span>
              <span class="text-yellow-600">بتا</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-600">GET /analytics</span>
              <span class="text-green-600">فعال</span>
            </div>
          </div>
        </div>

        <div class="border border-gray-200 rounded-lg p-6">
          <h5 class="font-medium text-gray-900 mb-2">کدهای وضعیت</h5>
          <div class="space-y-2 text-sm">
            <div class="flex justify-between">
              <span class="text-gray-600">۲۰۰ - موفق</span>
              <span class="text-green-600">عملیات موفق</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-600">۲۰۱ - ایجاد شد</span>
              <span class="text-green-600">ایمیل ارسال شد</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-600">۴۰۰ - خطای درخواست</span>
              <span class="text-red-600">داده‌های نامعتبر</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-600">۴۰۱ - غیرمجاز</span>
              <span class="text-red-600">کلید API نامعتبر</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- دکمه‌های عملیات -->
    <div class="flex justify-end space-x-3 space-x-reverse">
      <button class="px-4 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400 transition-colors">
        دانلود مستندات
      </button>
      <button class="px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700 transition-colors">
        ذخیره تنظیمات
      </button>
      <button class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors">
        تست API
      </button>
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
// کامپوننت مدیریت API سرویس Emails
// این کامپوننت مدیریت کلیدهای API، تنظیمات و مستندات سرویس Emails را فراهم می‌کند
</script> 
