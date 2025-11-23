<template>
  <div v-if="hasAccess" class="bg-white rounded-xl shadow p-6 h-80 flex flex-col">
    <div class="font-bold text-gray-700 mb-2">محصولات پرفروش</div>
    <div class="flex-1 overflow-y-auto space-y-3">
      <div v-for="i in 8" :key="i" class="flex items-center justify-between p-2 bg-gray-50 rounded-lg">
        <div class="flex items-center">
          <div class="w-8 h-8 bg-green-100 rounded-full flex items-center justify-center text-green-600 font-bold text-sm ml-2">{{i}}</div>
          <div>
            <div class="font-semibold text-sm">محصول شماره {{i}}</div>
            <div class="text-xs text-gray-500">{{50 - i*5}} فروش</div>
          </div>
        </div>
        <div class="text-left">
          <div class="font-bold text-green-600">{{(100 - i*10).toLocaleString('fa-IR')}}K</div>
          <div class="text-xs text-gray-400">تومان</div>
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

defineProps<{ products: unknown[] }>()

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
