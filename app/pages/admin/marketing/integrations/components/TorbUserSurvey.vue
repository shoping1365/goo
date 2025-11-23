<template>
  <div v-if="hasAccess" class="bg-white rounded-lg shadow-md p-6">
    <h3 class="text-lg font-semibold text-gray-800 mb-4">نظرسنجی کاربران سرویس Torb</h3>
    
    <!-- آمار نظرسنجی -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
      <div class="bg-blue-50 p-6 rounded-lg">
        <div class="text-sm text-blue-600">کل نظرسنجی‌ها</div>
        <div class="text-2xl font-bold text-blue-800">۴۵</div>
      </div>
      <div class="bg-green-50 p-6 rounded-lg">
        <div class="text-sm text-green-600">نظرسنجی‌های فعال</div>
        <div class="text-2xl font-bold text-green-800">۱۲</div>
      </div>
      <div class="bg-purple-50 p-6 rounded-lg">
        <div class="text-sm text-purple-600">کل پاسخ‌ها</div>
        <div class="text-2xl font-bold text-purple-800">۲,۴۵۰</div>
      </div>
      <div class="bg-orange-50 p-6 rounded-lg">
        <div class="text-sm text-orange-600">نرخ مشارکت</div>
        <div class="text-2xl font-bold text-orange-800">۶۸.۵٪</div>
      </div>
    </div>

    <!-- فیلترهای نظرسنجی -->
    <div class="mb-6">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نوع نظرسنجی</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option>همه</option>
            <option>رضایت مشتری</option>
            <option>کیفیت محصول</option>
            <option>خدمات پس از فروش</option>
            <option>تجربه خرید</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option>همه</option>
            <option>فعال</option>
            <option>غیرفعال</option>
            <option>در انتظار</option>
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

    <!-- جدول نظرسنجی‌ها -->
    <div class="mb-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">لیست نظرسنجی‌ها</h4>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عنوان نظرسنجی</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ شروع</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">پاسخ‌ها</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">رضایت از کیفیت محصولات</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">رضایت مشتری</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-green-100 text-green-800">فعال</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۱۴۰۳/۰۴/۱۰</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۴۵۰</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">ویرایش</button>
                <button class="text-red-600 hover:text-red-900">حذف</button>
              </td>
            </tr>
            <tr>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">تجربه خرید آنلاین</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">تجربه خرید</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-green-100 text-green-800">فعال</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۱۴۰۳/۰۴/۰۸</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۳۲۰</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">ویرایش</button>
                <button class="text-red-600 hover:text-red-900">حذف</button>
              </td>
            </tr>
            <tr>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">کیفیت خدمات پس از فروش</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">خدمات پس از فروش</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-yellow-100 text-yellow-800">در انتظار</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۱۴۰۳/۰۴/۰۵</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۱۸۰</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">ویرایش</button>
                <button class="text-red-600 hover:text-red-900">حذف</button>
              </td>
            </tr>
            <tr>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">رضایت از قیمت‌گذاری</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">رضایت مشتری</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-red-100 text-red-800">غیرفعال</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۱۴۰۳/۰۳/۲۸</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۲۸۰</td>
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

    <!-- نمودار نتایج -->
    <div class="mb-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">نمودار نتایج نظرسنجی</h4>
      <div class="bg-gray-50 p-6 rounded-lg h-64 flex items-center justify-center">
        <div class="text-center">
          <div class="text-gray-500 mb-2">نمودار نتایج</div>
          <div class="text-sm text-gray-400">نمایش نتایج نظرسنجی‌ها به صورت نمودار</div>
        </div>
      </div>
    </div>

    <!-- آمار پاسخ‌ها -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-6">
      <div class="border border-gray-200 rounded-lg p-6">
        <h5 class="font-medium text-gray-900 mb-3">رضایت کلی</h5>
        <div class="text-3xl font-bold text-green-600 mb-2">۸.۷</div>
        <div class="text-sm text-gray-600">از ۱۰ امتیاز</div>
        <div class="mt-2">
          <div class="flex items-center">
            <div class="flex-1 bg-gray-200 rounded-full h-2">
              <div class="bg-green-600 h-2 rounded-full" style="width: 87%"></div>
            </div>
            <span class="mr-2 text-sm text-gray-600">۸۷٪</span>
          </div>
        </div>
      </div>

      <div class="border border-gray-200 rounded-lg p-6">
        <h5 class="font-medium text-gray-900 mb-3">توصیه به دیگران</h5>
        <div class="text-3xl font-bold text-blue-600 mb-2">۹.۲</div>
        <div class="text-sm text-gray-600">از ۱۰ امتیاز</div>
        <div class="mt-2">
          <div class="flex items-center">
            <div class="flex-1 bg-gray-200 rounded-full h-2">
              <div class="bg-blue-600 h-2 rounded-full" style="width: 92%"></div>
            </div>
            <span class="mr-2 text-sm text-gray-600">۹۲٪</span>
          </div>
        </div>
      </div>

      <div class="border border-gray-200 rounded-lg p-6">
        <h5 class="font-medium text-gray-900 mb-3">احتمال خرید مجدد</h5>
        <div class="text-3xl font-bold text-purple-600 mb-2">۸.۹</div>
        <div class="text-sm text-gray-600">از ۱۰ امتیاز</div>
        <div class="mt-2">
          <div class="flex items-center">
            <div class="flex-1 bg-gray-200 rounded-full h-2">
              <div class="bg-purple-600 h-2 rounded-full" style="width: 89%"></div>
            </div>
            <span class="mr-2 text-sm text-gray-600">۸۹٪</span>
          </div>
        </div>
      </div>
    </div>

    <!-- نظرات کاربران -->
    <div class="mb-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">آخرین نظرات کاربران</h4>
      <div class="space-y-4">
        <div class="border border-gray-200 rounded-lg p-6">
          <div class="flex items-start justify-between mb-2">
            <div class="flex items-center">
              <div class="w-8 h-8 bg-blue-500 rounded-full flex items-center justify-center text-white text-sm font-medium">ع</div>
              <div class="mr-3">
                <h5 class="font-medium text-gray-900">علی احمدی</h5>
                <div class="flex items-center text-sm text-gray-500">
                  <span class="text-yellow-400">★★★★★</span>
                  <span class="mr-2">۵ ستاره</span>
                </div>
              </div>
            </div>
            <span class="text-xs text-gray-500">۲ روز پیش</span>
          </div>
          <p class="text-gray-700">کیفیت محصولات عالی است و خدمات پس از فروش بسیار خوبی دارید. حتماً دوباره خرید خواهم کرد.</p>
        </div>

        <div class="border border-gray-200 rounded-lg p-6">
          <div class="flex items-start justify-between mb-2">
            <div class="flex items-center">
              <div class="w-8 h-8 bg-green-500 rounded-full flex items-center justify-center text-white text-sm font-medium">س</div>
              <div class="mr-3">
                <h5 class="font-medium text-gray-900">سارا محمدی</h5>
                <div class="flex items-center text-sm text-gray-500">
                  <span class="text-yellow-400">★★★★☆</span>
                  <span class="mr-2">۴ ستاره</span>
                </div>
              </div>
            </div>
            <span class="text-xs text-gray-500">۳ روز پیش</span>
          </div>
          <p class="text-gray-700">قیمت‌ها مناسب است اما می‌تواند بهتر باشد. کیفیت محصولات خوب است.</p>
        </div>

        <div class="border border-gray-200 rounded-lg p-6">
          <div class="flex items-start justify-between mb-2">
            <div class="flex items-center">
              <div class="w-8 h-8 bg-purple-500 rounded-full flex items-center justify-center text-white text-sm font-medium">م</div>
              <div class="mr-3">
                <h5 class="font-medium text-gray-900">محمد رضایی</h5>
                <div class="flex items-center text-sm text-gray-500">
                  <span class="text-yellow-400">★★★★★</span>
                  <span class="mr-2">۵ ستاره</span>
                </div>
              </div>
            </div>
            <span class="text-xs text-gray-500">۱ هفته پیش</span>
          </div>
          <p class="text-gray-700">تجربه خرید بسیار خوبی داشتم. سایت سریع و کاربرپسند است.</p>
        </div>
      </div>
    </div>

    <!-- دکمه‌های عملیات -->
    <div class="flex justify-end space-x-3 space-x-reverse">
      <button class="px-4 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400 transition-colors">
        ایجاد نظرسنجی جدید
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
// کامپوننت نظرسنجی کاربران سرویس Torb
// این کامپوننت مدیریت نظرسنجی‌ها و تحلیل نتایج را فراهم می‌کند
</script> 
