<template>
  <div v-if="hasAccess" class="bg-white rounded-xl shadow p-6 h-80 flex flex-col">
    <div class="font-bold text-gray-700 mb-2">نمودار فروش</div>
    <div class="flex-1 flex items-center justify-center">
      <!-- Placeholder for chart -->
      <div class="w-full h-64 bg-gradient-to-r from-green-100 to-green-200 rounded-lg flex items-center justify-center text-green-400 text-2xl">
        [نمودار فروش روزانه]
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

defineProps<{ data: unknown[] }>()

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
