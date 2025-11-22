<template>
  <div v-if="hasAccess" class="flex items-center justify-between mb-6">
    <div>
      <h1 class="text-2xl font-bold text-blue-700">آمار بازدید سایت</h1>
      <p class="text-gray-500 mt-1">گزارش کامل بازدید کاربران، موقعیت، روندها و تحلیل پیشرفته</p>
    </div>
    <div class="flex gap-2">
      <button class="px-4 py-2 bg-blue-500 text-white rounded-lg shadow hover:bg-blue-600 transition">خروجی اکسل</button>
      <button class="px-4 py-2 bg-gray-100 text-gray-700 rounded-lg shadow hover:bg-gray-200 transition">بروزرسانی</button>
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
const checkAuth = async () => {
  if (!hasAccess.value) {
    await navigateTo('/404', { external: false })
  }
}

// بررسی احراز هویت در هنگام mount
onMounted(async () => {
  await checkAuth()
})

// بررسی احراز هویت هنگام تغییر وضعیت احراز هویت
watch([isAuthenticated, hasAccess], async () => {
  if (!hasAccess.value) {
    await checkAuth()
  }
})
</script> 
