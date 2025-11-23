<template>
  <div v-if="hasAccess" class="overflow-x-auto">
    <table class="min-w-full text-xs text-right rtl border-collapse">
      <thead>
      <tr class="bg-gray-100 border-b border-blue-200">
        <th class="px-3 py-2 text-gray-600 font-normal text-right">نقش مشتری</th>
        <th class="px-3 py-2 text-gray-600 font-normal text-right">حداقل تعداد</th>
        <th class="px-3 py-2 text-gray-600 font-normal text-right">قیمت واحد</th>
        <th class="px-3 py-2 text-gray-600 font-normal text-right">تخفیف</th>
        <th class="px-3 py-2 text-gray-600 font-normal text-right">تاریخ شروع</th>
        <th class="px-3 py-2 text-gray-600 font-normal text-right">تاریخ پایان</th>
        <th class="px-3 py-2 text-gray-600 font-normal text-right w-20">ویرایش</th>
        <th class="px-3 py-2 text-gray-600 font-normal text-right w-16">حذف</th>
      </tr>
      </thead>
      <tbody>
      <tr class="border-b hover:bg-gray-50 border-blue-200">
        <td class="px-3 py-2 text-right">عمده فروشان</td>
        <td class="px-3 py-2 text-right">10</td>
        <td class="px-3 py-2 text-right">45,000</td>
        <td class="px-3 py-2 text-right">10%</td>
        <td class="px-3 py-2 text-right">1403/08/01</td>
        <td class="px-3 py-2 text-right">1403/12/29</td>
        <td class="px-3 py-2 text-center">
          <button class="bg-yellow-400 text-white rounded px-2 py-1 text-xs">ویرایش</button>
        </td>
        <td class="px-3 py-2 text-center">
          <button class="bg-red-500 text-white rounded px-2 py-1 text-xs">حذف</button>
        </td>
      </tr>
      </tbody>
    </table>
  </div>
</template>

<script lang="ts">
declare const createError: (options: { statusCode: number; statusMessage: string }) => Error
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
const checkAuth = (): void => {
  if (!hasAccess.value) {
    // استفاده از createError که به صورت خودکار توسط Nuxt import می‌شود
    throw createError({
      statusCode: 404,
      statusMessage: 'Not Found'
    });
  }
};

// بررسی احراز هویت در هنگام mount
onMounted(() => {
  checkAuth();
});

// بررسی احراز هویت هنگام تغییر وضعیت احراز هویت
watch([isAuthenticated, hasAccess], () => {
  if (!hasAccess.value) {
    checkAuth();
  }
});

// اسکلت اولیه، بعداً دیتا و عملکرد اضافه می‌شود
</script> 
