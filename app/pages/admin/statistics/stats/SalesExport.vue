<template>
  <div v-if="hasAccess" class="bg-white rounded-xl shadow p-6 flex gap-6 items-center">
    <div class="font-bold text-gray-700">خروجی گرفتن از داده‌های فروش</div>
    <button class="px-4 py-2 bg-green-500 text-white rounded-lg shadow hover:bg-green-600 transition">Excel</button>
    <button class="px-4 py-2 bg-blue-500 text-white rounded-lg shadow hover:bg-blue-600 transition">CSV</button>
    <button class="px-4 py-2 bg-red-500 text-white rounded-lg shadow hover:bg-red-600 transition">PDF</button>
    <button class="px-4 py-2 bg-purple-500 text-white rounded-lg shadow hover:bg-purple-600 transition">گزارش مالی</button>
  </div>
</template>
<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
</script>

<script setup lang="ts">
import { computed, onMounted, watch } from 'vue';
import { useAuth } from '~/composables/useAuth';

defineProps<{ sales: unknown[] }>()

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
