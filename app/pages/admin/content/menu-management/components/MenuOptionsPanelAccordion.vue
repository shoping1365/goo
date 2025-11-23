<template>
  <div v-if="hasAccess" class="border-b">
    <button type="button" class="flex items-center justify-between w-full px-4 py-2 text-xs font-bold text-gray-700 focus:outline-none hover:bg-gray-50" @click="toggle">
      <span>{{ title }}</span>
      <span>{{ isOpen ? '▲' : '▼' }}</span>
    </button>
    <div v-if="isOpen" class="px-4 pb-3">
      <slot />
    </div>
  </div>
</template>

<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
</script>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue';
import { useAuth } from '~/composables/useAuth';

// احراز هویت
const { user, isAuthenticated } = useAuth();

// بررسی دسترسی admin
const hasAccess = computed(() => {
  if (!isAuthenticated.value) {
    return false;
  }

  const userRole = user.value?.role?.toLowerCase() || '';
  const adminRoles = ['admin', 'developer'];
  return adminRoles.includes(userRole);
});

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

const props = defineProps({
  title: String,
  open: Boolean
})

const isOpen = ref(props.open)

watch(() => props.open, (val) => {
  isOpen.value = val
})

function toggle() {
  isOpen.value = !isOpen.value
}
</script> 
