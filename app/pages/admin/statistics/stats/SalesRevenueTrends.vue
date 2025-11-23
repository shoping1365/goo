<template>
  <div v-if="hasAccess" class="bg-white rounded-xl shadow p-6">
    <div class="font-bold text-gray-700 mb-4">روند درآمد</div>
    <div class="space-y-4">
      <div class="flex justify-between items-center p-3 bg-green-50 rounded-lg">
        <div>
          <div class="font-semibold text-green-700">درآمد این ماه</div>
          <div class="text-sm text-gray-600">نسبت به ماه قبل</div>
        </div>
        <div class="text-right">
          <div class="font-bold text-green-600">+12.5%</div>
          <div class="text-xs text-green-500">رشد</div>
        </div>
      </div>
      <div class="flex justify-between items-center p-3 bg-blue-50 rounded-lg">
        <div>
          <div class="font-semibold text-blue-700">درآمد این هفته</div>
          <div class="text-sm text-gray-600">نسبت به هفته قبل</div>
        </div>
        <div class="text-right">
          <div class="font-bold text-blue-600">+8.3%</div>
          <div class="text-xs text-blue-500">رشد</div>
        </div>
      </div>
      <div class="flex justify-between items-center p-3 bg-orange-50 rounded-lg">
        <div>
          <div class="font-semibold text-orange-700">میانگین روزانه</div>
          <div class="text-sm text-gray-600">30 روز گذشته</div>
        </div>
        <div class="text-right">
          <div class="font-bold text-orange-600">29.7K</div>
          <div class="text-xs text-orange-500">تومان</div>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
</script>

<script setup lang="ts">
import { computed, onMounted, watch } from 'vue'
import { useAuth } from '~/composables/useAuth'

defineProps<{ trends: unknown[] }>()

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
