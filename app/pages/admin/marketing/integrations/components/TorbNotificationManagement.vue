<template>
  <div v-if="hasAccess" class="bg-white rounded-lg shadow-md p-6">
    <h3 class="text-lg font-semibold text-gray-800 mb-4">مدیریت اعلان‌ها سرویس Torb</h3>
    
    <!-- آمار اعلان‌ها -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
      <div class="bg-blue-50 p-6 rounded-lg">
        <div class="text-sm text-blue-600">کل اعلان‌ها</div>
        <div class="text-2xl font-bold text-blue-800">۲,۸۴۵</div>
      </div>
      <div class="bg-green-50 p-6 rounded-lg">
        <div class="text-sm text-green-600">ارسال شده</div>
        <div class="text-2xl font-bold text-green-800">۲,۶۷۰</div>
      </div>
      <div class="bg-purple-50 p-6 rounded-lg">
        <div class="text-sm text-purple-600">در انتظار</div>
        <div class="text-2xl font-bold text-purple-800">۱۲۵</div>
      </div>
      <div class="bg-orange-50 p-6 rounded-lg">
        <div class="text-sm text-orange-600">نرخ باز شدن</div>
        <div class="text-2xl font-bold text-orange-800">۷۸.۵٪</div>
      </div>
    </div>

    <!-- فیلترهای اعلان‌ها -->
    <div class="mb-6">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نوع اعلان</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option>همه</option>
            <option>تخفیف</option>
            <option>محصول جدید</option>
            <option>اخبار</option>
            <option>سیستم</option>
            <option>تبلیغات</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option>همه</option>
            <option>ارسال شده</option>
            <option>در انتظار</option>
            <option>ناموفق</option>
            <option>لغو شده</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ ارسال</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option>همه</option>
            <option>امروز</option>
            <option>هفته گذشته</option>
            <option>ماه گذشته</option>
            <option>سال گذشته</option>
          </select>
        </div>
        <div class="flex items-end">
          <button class="w-full px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors">
            اعمال فیلتر
          </button>
        </div>
      </div>
    </div>

    <!-- جدول اعلان‌ها -->
    <div class="mb-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">لیست اعلان‌ها</h4>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عنوان</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">گیرندگان</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نرخ باز شدن</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ ارسال</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr>
              <td class="px-6 py-4 whitespace-nowrap">
                <div>
                  <div class="text-sm font-medium text-gray-900">تخفیف ویژه ۲۰٪ روی تمام محصولات</div>
                  <div class="text-sm text-gray-500">کد تخفیف: SPRING20</div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-purple-100 text-purple-800">تخفیف</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۱۲,۴۵۰</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-green-100 text-green-800">ارسال شده</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۸۵.۲٪</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">امروز ۱۰:۳۰</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">ویرایش</button>
                <button class="text-red-600 hover:text-red-900">حذف</button>
              </td>
            </tr>
            <tr>
              <td class="px-6 py-4 whitespace-nowrap">
                <div>
                  <div class="text-sm font-medium text-gray-900">محصولات جدید اضافه شد</div>
                  <div class="text-sm text-gray-500">۵۰ محصول جدید در دسته الکترونیک</div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-blue-100 text-blue-800">محصول جدید</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۸,۲۳۰</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-green-100 text-green-800">ارسال شده</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۷۲.۸٪</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">دیروز ۱۶:۴۵</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">ویرایش</button>
                <button class="text-red-600 hover:text-red-900">حذف</button>
              </td>
            </tr>
            <tr>
              <td class="px-6 py-4 whitespace-nowrap">
                <div>
                  <div class="text-sm font-medium text-gray-900">بروزرسانی سیستم</div>
                  <div class="text-sm text-gray-500">سیستم در تاریخ ۱۴۰۳/۰۳/۱۵ بروزرسانی می‌شود</div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-gray-100 text-gray-800">سیستم</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۲,۱۵۰</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-yellow-100 text-yellow-800">در انتظار</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">-</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">فردا ۰۸:۰۰</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">ویرایش</button>
                <button class="text-red-600 hover:text-red-900">لغو</button>
              </td>
            </tr>
            <tr>
              <td class="px-6 py-4 whitespace-nowrap">
                <div>
                  <div class="text-sm font-medium text-gray-900">اخبار فناوری</div>
                  <div class="text-sm text-gray-500">آخرین اخبار دنیای تکنولوژی</div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-green-100 text-green-800">اخبار</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۱۵,۶۷۰</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-red-100 text-red-800">ناموفق</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۴۵.۳٪</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">۲ روز پیش</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">ارسال مجدد</button>
                <button class="text-red-600 hover:text-red-900">حذف</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- نمودار اعلان‌ها -->
    <div class="mb-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">نمودار عملکرد اعلان‌ها</h4>
      <div class="bg-gray-50 p-6 rounded-lg h-64 flex items-center justify-center">
        <div class="text-center">
          <div class="text-gray-500 mb-2">نمودار عملکرد</div>
          <div class="text-sm text-gray-400">نمایش نرخ باز شدن و کلیک اعلان‌ها</div>
        </div>
      </div>
    </div>

    <!-- آمار تفصیلی -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
      <div class="border border-gray-200 rounded-lg p-6">
        <h5 class="font-medium text-gray-900 mb-3">عملکرد اعلان‌ها</h5>
        <div class="space-y-3">
          <div class="flex items-center justify-between">
            <div>
              <h6 class="font-medium text-gray-900">تخفیف</h6>
              <p class="text-sm text-gray-600">۸۵.۲٪ نرخ باز شدن</p>
            </div>
            <span class="text-purple-600 font-medium">۸۵.۲٪</span>
          </div>
          <div class="flex items-center justify-between">
            <div>
              <h6 class="font-medium text-gray-900">محصول جدید</h6>
              <p class="text-sm text-gray-600">۷۲.۸٪ نرخ باز شدن</p>
            </div>
            <span class="text-blue-600 font-medium">۷۲.۸٪</span>
          </div>
          <div class="flex items-center justify-between">
            <div>
              <h6 class="font-medium text-gray-900">اخبار</h6>
              <p class="text-sm text-gray-600">۴۵.۳٪ نرخ باز شدن</p>
            </div>
            <span class="text-green-600 font-medium">۴۵.۳٪</span>
          </div>
        </div>
      </div>

      <div class="border border-gray-200 rounded-lg p-6">
        <h5 class="font-medium text-gray-900 mb-3">آمار کلی</h5>
        <div class="space-y-3">
          <div class="flex items-center justify-between">
            <span class="text-gray-600">میانگین نرخ باز شدن</span>
            <span class="font-medium">۷۸.۵٪</span>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-gray-600">میانگین نرخ کلیک</span>
            <span class="font-medium">۱۲.۳٪</span>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-gray-600">زمان بهینه ارسال</span>
            <span class="font-medium">۱۰:۰۰ - ۱۴:۰۰</span>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-gray-600">نرخ لغو اشتراک</span>
            <span class="font-medium">۲.۱٪</span>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات اعلان‌ها -->
    <div class="mb-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">تنظیمات اعلان‌ها</h4>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="space-y-4">
          <div class="flex items-center justify-between">
            <div>
              <h5 class="font-medium text-gray-900">اعلان‌های خودکار</h5>
              <p class="text-sm text-gray-600">ارسال خودکار اعلان‌های تخفیف</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input type="checkbox" checked class="sr-only peer">
              <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
            </label>
          </div>
          
          <div class="flex items-center justify-between">
            <div>
              <h5 class="font-medium text-gray-900">اعلان‌های محصول جدید</h5>
              <p class="text-sm text-gray-600">اعلان محصولات جدید به کاربران</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input type="checkbox" checked class="sr-only peer">
              <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
            </label>
          </div>

          <div class="flex items-center justify-between">
            <div>
              <h5 class="font-medium text-gray-900">اعلان‌های سیستم</h5>
              <p class="text-sm text-gray-600">اعلان بروزرسانی‌ها و تغییرات سیستم</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input type="checkbox" class="sr-only peer">
              <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
            </label>
          </div>
        </div>

        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر اعلان روزانه</label>
            <input type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500" value="3">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">زمان ارسال (ساعت)</label>
            <input type="time" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500" value="10:00">
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداقل فاصله (ساعت)</label>
            <input type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500" value="4">
          </div>
        </div>
      </div>
    </div>

    <!-- دکمه‌های عملیات -->
    <div class="flex justify-end space-x-3 space-x-reverse">
      <button class="px-4 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400 transition-colors">
        ایجاد اعلان جدید
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
// کامپوننت مدیریت اعلان‌ها سرویس Torb
// این کامپوننت مدیریت و نظارت بر اعلان‌ها را فراهم می‌کند
</script> 
