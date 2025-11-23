<template>
  <div v-if="hasAccess" class="flex items-center justify-between mb-4">
    <h2 class="text-lg font-bold">مدیریت انبارها</h2>
    <div class="flex items-center gap-2">
      <input :value="query" type="text" placeholder="جستجو (نام/کد/شهر)" class="rounded px-3 py-2 text-sm w-64 bg-white shadow-sm border border-transparent focus:border-blue-300 focus:ring-2 focus:ring-blue-200" @input="$emit('update:query', ($event.target as HTMLInputElement).value)"/>
      <slot name="extra-actions"></slot>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, watch } from 'vue';
import { useAuth } from '~/composables/useAuth';

declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>;

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

// emit در template استفاده می‌شود
// eslint-disable-next-line @typescript-eslint/no-unused-vars
const emit = defineEmits<{
  (e: 'update:query', value: string): void
}>()

defineProps<{
  query?: string
}>()
</script>
