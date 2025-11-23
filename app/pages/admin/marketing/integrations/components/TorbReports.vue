<template>
  <div v-if="hasAccess" class="space-y-4">
    <h2 class="text-xl font-bold">گزارشات سرویس ترب</h2>
    <p class="text-gray-600">در این بخش می‌توانید گزارشات مربوط به سرویس ترب را مشاهده کنید.</p>
    <!-- جدول گزارشات ترب -->
    <table class="w-full border rounded">
      <thead>
        <tr class="bg-gray-100">
          <th class="p-2">تاریخ</th>
          <th class="p-2">نوع گزارش</th>
          <th class="p-2">وضعیت</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td class="p-2">1403/02/01</td>
          <td class="p-2">گزارش قیمت</td>
          <td class="p-2 text-green-600">موفق</td>
        </tr>
        <tr>
          <td class="p-2">1403/02/02</td>
          <td class="p-2">گزارش موجودی</td>
          <td class="p-2 text-red-600">خطا</td>
        </tr>
      </tbody>
    </table>
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
// این کامپوننت گزارشات سرویس ترب را نمایش می‌دهد
</script>

<!-- توضیحات: این کامپوننت برای نمایش گزارشات سرویس ترب استفاده می‌شود. --> 
