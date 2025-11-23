<template>
  <div v-if="hasAccess" class="bg-white rounded-lg shadow-md p-6">
    <h3 class="text-lg font-semibold text-gray-800 mb-4">مدیریت کمپین‌ها سرویس Torb</h3>
    
    <!-- آمار کمپین‌ها -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
      <div class="bg-blue-50 p-6 rounded-lg">
        <div class="text-sm text-blue-600">کل کمپین‌ها</div>
        <div class="text-2xl font-bold text-blue-800">۲۴</div>
      </div>
      <div class="bg-green-50 p-6 rounded-lg">
        <div class="text-sm text-green-600">کمپین‌های فعال</div>
        <div class="text-2xl font-bold text-green-800">۸</div>
      </div>
      <div class="bg-purple-50 p-6 rounded-lg">
        <div class="text-sm text-purple-600">نرخ موفقیت</div>
        <div class="text-2xl font-bold text-purple-800">۷۸.۵٪</div>
      </div>
      <div class="bg-orange-50 p-6 rounded-lg">
        <div class="text-sm text-orange-600">درآمد کل</div>
        <div class="text-2xl font-bold text-orange-800">۱۲,۴۵۰,۰۰۰ تومان</div>
      </div>
    </div>

    <!-- فیلترهای کمپین -->
    <div class="mb-6">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
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
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">بازه زمانی</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option>همه</option>
            <option>امروز</option>
            <option>هفته گذشته</option>
            <option>ماه گذشته</option>
          </select>
        </div>
        <div class="flex items-end">
          <button class="w-full px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors">
            اعمال فیلتر
          </button>
        </div>
      </div>
    </div>

    <!-- جدول کمپین‌ها -->
    <div class="mb-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">لیست کمپین‌ها</h4>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نام کمپین</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ شروع</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">بودجه</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملکرد</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">کمپین تابستانه ۱۴۰۳</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">ایمیل</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-green-100 text-green-800">فعال</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۱۴۰۳/۰۴/۱۵</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۵,۰۰۰,۰۰۰ تومان</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="flex-1 bg-gray-200 rounded-full h-2 w-16">
                    <div class="bg-green-600 h-2 rounded-full" style="width: 85%"></div>
                  </div>
                  <span class="mr-2 text-xs text-gray-600">۸۵٪</span>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">ویرایش</button>
                <button class="text-red-600 hover:text-red-900">حذف</button>
              </td>
            </tr>
            <tr>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">پیشنهاد ویژه محصولات</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">پیامک</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-blue-100 text-blue-800">تکمیل شده</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۱۴۰۳/۰۴/۱۰</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۳,۵۰۰,۰۰۰ تومان</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="flex-1 bg-gray-200 rounded-full h-2 w-16">
                    <div class="bg-blue-600 h-2 rounded-full" style="width: 92%"></div>
                  </div>
                  <span class="mr-2 text-xs text-gray-600">۹۲٪</span>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">ویرایش</button>
                <button class="text-red-600 hover:text-red-900">حذف</button>
              </td>
            </tr>
            <tr>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">اعلان تخفیف ۲۰٪</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">اعلان</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-yellow-100 text-yellow-800">در انتظار</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۱۴۰۳/۰۴/۱۸</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۲,۰۰۰,۰۰۰ تومان</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="flex-1 bg-gray-200 rounded-full h-2 w-16">
                    <div class="bg-gray-400 h-2 rounded-full" style="width: 0%"></div>
                  </div>
                  <span class="mr-2 text-xs text-gray-600">۰٪</span>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">ویرایش</button>
                <button class="text-red-600 hover:text-red-900">حذف</button>
              </td>
            </tr>
            <tr>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">خوش‌آمدگویی مشتریان جدید</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">ایمیل</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-green-100 text-green-800">فعال</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۱۴۰۳/۰۴/۰۵</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۱,۵۰۰,۰۰۰ تومان</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="flex-1 bg-gray-200 rounded-full h-2 w-16">
                    <div class="bg-green-600 h-2 rounded-full" style="width: 78%"></div>
                  </div>
                  <span class="mr-2 text-xs text-gray-600">۷۸٪</span>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">ویرایش</button>
                <button class="text-red-600 hover:text-red-900">حذف</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- نمودار عملکرد -->
    <div class="mb-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">نمودار عملکرد کمپین‌ها</h4>
      <div class="bg-gray-50 p-6 rounded-lg h-64 flex items-center justify-center">
        <div class="text-center">
          <div class="text-gray-500 mb-2">نمودار عملکرد</div>
          <div class="text-sm text-gray-400">نمایش روند عملکرد کمپین‌ها در طول زمان</div>
        </div>
      </div>
    </div>

    <!-- آمار تفصیلی -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
      <div class="border border-gray-200 rounded-lg p-6">
        <h5 class="font-medium text-gray-900 mb-3">بهترین کمپین‌ها</h5>
        <div class="space-y-3">
          <div class="flex items-center justify-between">
            <div>
              <h6 class="font-medium text-gray-900">پیشنهاد ویژه محصولات</h6>
              <p class="text-sm text-gray-600">نرخ تبدیل: ۱۲.۵٪</p>
            </div>
            <span class="text-green-600 font-medium">۹۲٪</span>
          </div>
          <div class="flex items-center justify-between">
            <div>
              <h6 class="font-medium text-gray-900">کمپین تابستانه ۱۴۰۳</h6>
              <p class="text-sm text-gray-600">نرخ تبدیل: ۸.۷٪</p>
            </div>
            <span class="text-green-600 font-medium">۸۵٪</span>
          </div>
          <div class="flex items-center justify-between">
            <div>
              <h6 class="font-medium text-gray-900">خوش‌آمدگویی مشتریان</h6>
              <p class="text-sm text-gray-600">نرخ تبدیل: ۱۵.۲٪</p>
            </div>
            <span class="text-green-600 font-medium">۷۸٪</span>
          </div>
        </div>
      </div>

      <div class="border border-gray-200 rounded-lg p-6">
        <h5 class="font-medium text-gray-900 mb-3">آمار کلی</h5>
        <div class="space-y-3">
          <div class="flex items-center justify-between">
            <span class="text-gray-600">میانگین نرخ باز شدن</span>
            <span class="font-medium">۷۴.۸٪</span>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-gray-600">میانگین نرخ کلیک</span>
            <span class="font-medium">۲۸.۵٪</span>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-gray-600">میانگین نرخ تبدیل</span>
            <span class="font-medium">۱۲.۱٪</span>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-gray-600">ROI متوسط</span>
            <span class="font-medium">۳.۲</span>
          </div>
        </div>
      </div>
    </div>

    <!-- دکمه‌های عملیات -->
    <div class="flex justify-end space-x-3 space-x-reverse">
      <button class="px-4 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400 transition-colors">
        ایجاد کمپین جدید
      </button>
      <button class="px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700 transition-colors">
        صادر کردن گزارش
      </button>
      <button class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors">
        بروزرسانی
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
// کامپوننت مدیریت کمپین‌ها سرویس Torb
// این کامپوننت مدیریت و نظارت بر کمپین‌های بازاریابی را فراهم می‌کند
</script> 
