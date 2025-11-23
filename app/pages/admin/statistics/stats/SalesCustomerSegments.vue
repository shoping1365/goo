<template>
  <div v-if="hasAccess" class="bg-white rounded-xl shadow p-6">
    <div class="font-bold text-gray-700 mb-4">بخش‌بندی مشتریان</div>
    <div class="space-y-3">
      <div class="flex justify-between items-center">
        <span class="text-gray-600">مشتریان جدید</span>
        <div class="flex items-center">
          <div class="w-20 h-2 bg-gray-200 rounded-full ml-2">
            <div class="w-3/5 h-full bg-blue-500 rounded-full"></div>
          </div>
          <span class="text-sm font-bold">60%</span>
        </div>
      </div>
      <div class="flex justify-between items-center">
        <span class="text-gray-600">مشتریان بازگشتی</span>
        <div class="flex items-center">
          <div class="w-20 h-2 bg-gray-200 rounded-full ml-2">
            <div class="w-2/5 h-full bg-green-500 rounded-full"></div>
          </div>
          <span class="text-sm font-bold">35%</span>
        </div>
      </div>
      <div class="flex justify-between items-center">
        <span class="text-gray-600">مشتریان VIP</span>
        <div class="flex items-center">
          <div class="w-20 h-2 bg-gray-200 rounded-full ml-2">
            <div class="w-1/10 h-full bg-purple-500 rounded-full"></div>
          </div>
          <span class="text-sm font-bold">5%</span>
        </div>
      </div>
      <div class="mt-4 p-3 bg-purple-50 rounded-lg">
        <div class="text-sm font-semibold text-purple-700">درآمد مشتریان VIP</div>
        <div class="text-lg font-bold text-purple-600">45% از کل درآمد</div>
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

defineProps<{ segments: unknown[] }>()

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
