<template>
  <div v-if="hasAccess" class="bg-white rounded-xl shadow p-6">
    <div class="font-bold text-gray-700 mb-4">تحلیل پیشرفته فروش</div>
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="bg-green-50 rounded-lg p-6 text-center">
        <div class="font-bold text-green-600 mb-2">Cohort Analysis</div>
        <div class="text-green-400">[تحلیل گروهی مشتریان]</div>
      </div>
      <div class="bg-red-50 rounded-lg p-6 text-center">
        <div class="font-bold text-red-600 mb-2">Churn Rate</div>
        <div class="text-red-400">[نرخ ترک مشتریان]</div>
      </div>
      <div class="bg-purple-50 rounded-lg p-6 text-center">
        <div class="font-bold text-purple-600 mb-2">CLV Analysis</div>
        <div class="text-purple-400">[ارزش طول عمر مشتری]</div>
      </div>
      <div class="bg-blue-50 rounded-lg p-6 text-center">
        <div class="font-bold text-blue-600 mb-2">RFM Segmentation</div>
        <div class="text-blue-400">[بخش‌بندی RFM]</div>
      </div>
    </div>
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mt-4">
      <div class="bg-yellow-50 rounded-lg p-6">
        <div class="font-bold text-yellow-600 mb-2">پیش‌بینی فروش</div>
        <div class="text-sm text-gray-600">بر اساس روندهای گذشته</div>
        <div class="text-lg font-bold text-yellow-600 mt-2">+15% رشد پیش‌بینی شده</div>
      </div>
      <div class="bg-indigo-50 rounded-lg p-6">
        <div class="font-bold text-indigo-600 mb-2">تحلیل فصلی</div>
        <div class="text-sm text-gray-600">الگوهای فروش در طول سال</div>
        <div class="text-lg font-bold text-indigo-600 mt-2">پیک فروش: زمستان</div>
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

defineProps<{ analytics: unknown }>()

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
