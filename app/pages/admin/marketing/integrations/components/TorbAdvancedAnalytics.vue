<template>
  <div v-if="hasAccess" class="space-y-4">
    <h2 class="text-xl font-bold">تحلیل‌های پیشرفته ترب</h2>
    <p class="text-gray-600">در این بخش می‌توانید تحلیل‌های آماری و داده‌ای سرویس ترب را مشاهده کنید.</p>
    <!-- نمودار یا جدول تحلیل‌ها -->
    <div class="bg-white border rounded p-6 text-center text-gray-500">نمودار و داده‌های تحلیلی (در حال توسعه)</div>
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
// این کامپوننت تحلیل‌های پیشرفته سرویس ترب را نمایش می‌دهد
</script>

<!-- توضیحات: این کامپوننت برای نمایش تحلیل‌های پیشرفته ترب استفاده می‌شود. --> 
