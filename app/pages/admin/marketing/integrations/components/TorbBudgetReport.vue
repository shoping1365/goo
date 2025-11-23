<template>
  <div v-if="hasAccess" class="space-y-4">
    <h2 class="text-xl font-bold">گزارش بودجه ترب</h2>
    <p class="text-gray-600">در این بخش می‌توانید گزارش‌های بودجه سرویس ترب را مشاهده کنید.</p>
    <!-- جدول بودجه ترب -->
    <table class="w-full border rounded">
      <thead>
        <tr class="bg-gray-100">
          <th class="p-2">دوره</th>
          <th class="p-2">بودجه تخصیص یافته</th>
          <th class="p-2">بودجه مصرف شده</th>
          <th class="p-2">مانده بودجه</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td class="p-2">فروردین 1403</td>
          <td class="p-2">۵۰٬۰۰۰٬۰۰۰</td>
          <td class="p-2">۳۵٬۰۰۰٬۰۰۰</td>
          <td class="p-2">۱۵٬۰۰۰٬۰۰۰</td>
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
// این کامپوننت گزارش بودجه سرویس ترب را نمایش می‌دهد
</script>

<!-- توضیحات: این کامپوننت برای نمایش گزارش بودجه ترب استفاده می‌شود. --> 
