<template>
  <div v-if="hasAccess" :class="['rounded-xl p-6 shadow bg-white flex flex-col items-center', 'border-t-4', 'border-' + color + '-500']">
    <div class="mb-2">
      <span :class="['mdi', icon, 'text-2xl', 'text-' + color + '-500']"></span>
    </div>
    <div class="text-lg font-bold text-gray-800">{{ value }}</div>
    <div class="text-xs text-gray-500 mt-1">{{ title }}</div>
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

defineProps<{ title: string, value: unknown, icon: string, color: string }>()
</script> 
