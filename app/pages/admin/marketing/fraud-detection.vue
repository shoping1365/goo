<template>
  <div v-if="hasAccess" class="space-y-6">
    <!-- هدر صفحه -->
    <div class="bg-white rounded-lg shadow-sm p-6">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-2xl font-bold text-gray-900">تشخیص تقلب و امنیت</h1>
          <p class="mt-2 text-gray-600">مدیریت و نظارت بر فعالیت‌های مشکوک و تقلبی در سیستم</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors">
            <span class="i-heroicons-cog-6-tooth ml-2"></span>
            تنظیمات
          </button>
        </div>
      </div>
    </div>

    <!-- کامپوننت تشخیص تقلب -->
    <FraudDetectionSecurity />
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
declare const useHead: (head: { title?: string }) => void
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>;
</script>

<script setup lang="ts">
import { computed, onMounted, watch } from 'vue';
import { useAuth } from '~/composables/useAuth';

// احراز هویت
const { user, isAuthenticated } = useAuth()

// بررسی دسترسی admin
const hasAccess = computed(() => {
  if (!isAuthenticated.value) {
    return false
  }

  const userRole = user.value?.role?.toLowerCase() || ''
  const adminRoles = ['admin', 'developer']
  return adminRoles.includes(userRole)
})

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

// تعریف متا صفحه
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// عنوان صفحه
useHead({
  title: 'تشخیص تقلب و امنیت - پنل مدیریت'
})
</script> 
