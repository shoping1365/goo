<template>
  <div v-if="hasAccess" class="bg-white rounded-lg shadow-md p-6">
    <h3 class="text-lg font-semibold text-gray-800 mb-4">گزارش‌های جامع سرویس Torb</h3>
    
    <!-- فیلترهای گزارش -->
    <div class="mb-6">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">بازه زمانی</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option>امروز</option>
            <option>هفته گذشته</option>
            <option>ماه گذشته</option>
            <option>سه ماه گذشته</option>
            <option>سال گذشته</option>
            <option>سفارشی</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نوع کمپین</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option>همه</option>
            <option>ایمیل</option>
            <option>پیامک</option>
            <option>اعلان</option>
            <option>بازاریابی</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option>همه</option>
            <option>فعال</option>
            <option>غیرفعال</option>
            <option>در انتظار</option>
            <option>تکمیل شده</option>
          </select>
        </div>
        <div class="flex items-end">
          <button class="w-full px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors">
            اعمال فیلتر
          </button>
        </div>
      </div>
    </div>

    <!-- آمار کلی -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
      <div class="bg-blue-50 p-6 rounded-lg">
        <div class="text-sm text-blue-600">کل کمپین‌ها</div>
        <div class="text-2xl font-bold text-blue-800">۲۴۵</div>
        <div class="text-xs text-blue-600 mt-1">+۱۲٪ نسبت به ماه گذشته</div>
      </div>
      <div class="bg-green-50 p-6 rounded-lg">
        <div class="text-sm text-green-600">نرخ باز شدن</div>
        <div class="text-2xl font-bold text-green-800">۶۸.۵٪</div>
        <div class="text-xs text-green-600 mt-1">+۵.۲٪ نسبت به ماه گذشته</div>
      </div>
      <div class="bg-purple-50 p-6 rounded-lg">
        <div class="text-sm text-purple-600">نرخ کلیک</div>
        <div class="text-2xl font-bold text-purple-800">۲۳.۸٪</div>
        <div class="text-xs text-purple-600 mt-1">+۳.۱٪ نسبت به ماه گذشته</div>
      </div>
      <div class="bg-orange-50 p-6 rounded-lg">
        <div class="text-sm text-orange-600">درآمد کل</div>
        <div class="text-2xl font-bold text-orange-800">۱۲,۴۵۰,۰۰۰ تومان</div>
        <div class="text-xs text-orange-600 mt-1">+۱۸.۷٪ نسبت به ماه گذشته</div>
      </div>
    </div>

    <!-- نمودار عملکرد -->
    <div class="mb-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">نمودار عملکرد کمپین‌ها</h4>
      <div class="bg-gray-50 p-6 rounded-lg h-64 flex items-center justify-center">
        <div class="text-center">
          <div class="text-gray-500 mb-2">نمودار عملکرد</div>
          <div class="text-sm text-gray-400">نمایش روند عملکرد کمپین‌ها در بازه زمانی انتخاب شده</div>
        </div>
      </div>
    </div>

    <!-- جدول گزارش‌های تفصیلی -->
    <div class="mb-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">گزارش‌های تفصیلی</h4>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نام کمپین</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ شروع</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">باز شدن</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کلیک</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">درآمد</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">کمپین تابستانه ۱۴۰۳</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">ایمیل</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۱۴۰۳/۰۴/۱۵</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-green-100 text-green-800">فعال</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۷۲.۳٪</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۲۸.۷٪</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۲,۴۵۰,۰۰۰ تومان</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">دانلود</button>
              </td>
            </tr>
            <tr>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">پیشنهاد ویژه محصولات</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">پیامک</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۱۴۰۳/۰۴/۱۰</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-blue-100 text-blue-800">تکمیل شده</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۶۵.۸٪</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۲۲.۱٪</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۱,۸۷۰,۰۰۰ تومان</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">دانلود</button>
              </td>
            </tr>
            <tr>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">اعلان تخفیف ۲۰٪</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">اعلان</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۱۴۰۳/۰۴/۰۸</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-yellow-100 text-yellow-800">در انتظار</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">-</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">-</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">-</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">دانلود</button>
              </td>
            </tr>
            <tr>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">خوش‌آمدگویی مشتریان جدید</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">ایمیل</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۱۴۰۳/۰۴/۰۵</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-green-100 text-green-800">فعال</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۸۹.۲٪</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۳۴.۵٪</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۳,۲۱۰,۰۰۰ تومان</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">دانلود</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- گزارش‌های تحلیلی -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
      <div class="border border-gray-200 rounded-lg p-6">
        <h5 class="font-medium text-gray-900 mb-3">بهترین زمان ارسال</h5>
        <div class="space-y-2 text-sm">
          <div class="flex justify-between">
            <span class="text-gray-600">دوشنبه ۹ صبح</span>
            <span class="text-green-600">۸۹.۲٪</span>
          </div>
          <div class="flex justify-between">
            <span class="text-gray-600">سه‌شنبه ۱۰ صبح</span>
            <span class="text-green-600">۸۷.۵٪</span>
          </div>
          <div class="flex justify-between">
            <span class="text-gray-600">چهارشنبه ۱۱ صبح</span>
            <span class="text-green-600">۸۵.۸٪</span>
          </div>
          <div class="flex justify-between">
            <span class="text-gray-600">پنج‌شنبه ۹ صبح</span>
            <span class="text-green-600">۸۴.۳٪</span>
          </div>
        </div>
      </div>

      <div class="border border-gray-200 rounded-lg p-6">
        <h5 class="font-medium text-gray-900 mb-3">بهترین موضوعات</h5>
        <div class="space-y-2 text-sm">
          <div class="flex justify-between">
            <span class="text-gray-600">تخفیف ویژه</span>
            <span class="text-green-600">۷۸.۹٪</span>
          </div>
          <div class="flex justify-between">
            <span class="text-gray-600">محصولات جدید</span>
            <span class="text-green-600">۷۲.۴٪</span>
          </div>
          <div class="flex justify-between">
            <span class="text-gray-600">خوش‌آمدگویی</span>
            <span class="text-green-600">۸۹.۲٪</span>
          </div>
          <div class="flex justify-between">
            <span class="text-gray-600">یادآوری</span>
            <span class="text-green-600">۶۸.۷٪</span>
          </div>
        </div>
      </div>
    </div>

    <!-- دکمه‌های عملیات -->
    <div class="flex justify-between items-center">
      <div class="flex space-x-3 space-x-reverse">
        <button class="px-4 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400 transition-colors">
          صادر کردن PDF
        </button>
        <button class="px-4 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400 transition-colors">
          صادر کردن Excel
        </button>
        <button class="px-4 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400 transition-colors">
          ارسال ایمیل
        </button>
      </div>
      <div class="flex space-x-3 space-x-reverse">
        <button class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors">
          بروزرسانی
        </button>
        <button class="px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700 transition-colors">
          ذخیره گزارش
        </button>
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
// کامپوننت گزارش‌های جامع سرویس Torb
// این کامپوننت گزارش‌های تفصیلی و تحلیلی سرویس Torb را نمایش می‌دهد
</script> 
