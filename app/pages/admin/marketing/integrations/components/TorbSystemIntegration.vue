<template>
  <div v-if="hasAccess" class="space-y-4">
    <h2 class="text-xl font-bold">یکپارچه‌سازی سیستم ترب</h2>
    <p class="text-gray-600">در این بخش می‌توانید تنظیمات و وضعیت یکپارچه‌سازی سرویس ترب با سایر سیستم‌ها را مدیریت کنید.</p>
    <!-- فرم یکپارچه‌سازی -->
    <form class="space-y-3">
      <div>
        <label class="block mb-1">انتخاب سیستم مقصد</label>
        <select class="w-full border rounded px-3 py-2">
          <option>سیستم سفارشات</option>
          <option>سیستم کاربران</option>
          <option>سیستم گزارش‌گیری</option>
        </select>
      </div>
      <button type="submit" class="bg-blue-600 text-white px-4 py-2 rounded">ذخیره تنظیمات</button>
    </form>
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
// این کامپوننت یکپارچه‌سازی سیستم ترب را مدیریت می‌کند
</script>

<!-- توضیحات: این کامپوننت برای مدیریت یکپارچه‌سازی سیستم ترب استفاده می‌شود. --> 
