<template>
  <div v-if="hasAccess" class="bg-white rounded-lg shadow-md p-6">
    <h3 class="text-lg font-semibold text-gray-800 mb-4">خلاصه بودجه Torb</h3>
    
    <!-- خلاصه کلی بودجه -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-6">
      <div class="bg-blue-50 p-6 rounded-lg">
        <div class="text-sm text-blue-600">بودجه کل</div>
        <div class="text-2xl font-bold text-blue-800">۱۲,۵۰۰,۰۰۰ تومان</div>
      </div>
      <div class="bg-green-50 p-6 rounded-lg">
        <div class="text-sm text-green-600">هزینه شده</div>
        <div class="text-2xl font-bold text-green-800">۸,۲۰۰,۰۰۰ تومان</div>
      </div>
      <div class="bg-orange-50 p-6 rounded-lg">
        <div class="text-sm text-orange-600">باقی‌مانده</div>
        <div class="text-2xl font-bold text-orange-800">۴,۳۰۰,۰۰۰ تومان</div>
      </div>
    </div>

    <!-- نمودار توزیع بودجه -->
    <div class="mb-6">
      <h4 class="text-md font-medium text-gray-700 mb-3">توزیع بودجه بر اساس دسته‌بندی</h4>
      <div class="space-y-3">
        <div class="flex items-center justify-between">
          <span class="text-sm text-gray-600">تبلیغات کلیکی</span>
          <div class="flex items-center space-x-2">
            <div class="w-32 bg-gray-200 rounded-full h-2">
              <div class="bg-blue-500 h-2 rounded-full" style="width: 65%"></div>
            </div>
            <span class="text-sm font-medium">۶۵٪</span>
          </div>
        </div>
        <div class="flex items-center justify-between">
          <span class="text-sm text-gray-600">تبلیغات نمایشی</span>
          <div class="flex items-center space-x-2">
            <div class="w-32 bg-gray-200 rounded-full h-2">
              <div class="bg-green-500 h-2 rounded-full" style="width: 25%"></div>
            </div>
            <span class="text-sm font-medium">۲۵٪</span>
          </div>
        </div>
        <div class="flex items-center justify-between">
          <span class="text-sm text-gray-600">سایر هزینه‌ها</span>
          <div class="flex items-center space-x-2">
            <div class="w-32 bg-gray-200 rounded-full h-2">
              <div class="bg-orange-500 h-2 rounded-full" style="width: 10%"></div>
            </div>
            <span class="text-sm font-medium">۱۰٪</span>
          </div>
        </div>
      </div>
    </div>

    <!-- جدول جزئیات بودجه -->
    <div>
      <h4 class="text-md font-medium text-gray-700 mb-3">جزئیات هزینه‌ها</h4>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">دسته‌بندی</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مبلغ</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">درصد</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">تبلیغات کلیکی</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">۸,۱۲۵,۰۰۰ تومان</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">۶۵٪</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-green-100 text-green-800">فعال</span>
              </td>
            </tr>
            <tr>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">تبلیغات نمایشی</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">۳,۱۲۵,۰۰۰ تومان</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">۲۵٪</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-yellow-100 text-yellow-800">در انتظار</span>
              </td>
            </tr>
            <tr>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">سایر هزینه‌ها</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">۱,۲۵۰,۰۰۰ تومان</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">۱۰٪</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-blue-100 text-blue-800">برنامه‌ریزی شده</span>
              </td>
            </tr>
          </tbody>
        </table>
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
// کامپوننت خلاصه بودجه Torb
// این کامپوننت نمایش خلاصه‌ای از بودجه و هزینه‌های سرویس Torb را فراهم می‌کند
</script> 
