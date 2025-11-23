<template>
  <div v-if="hasAccess" class="meta-box">
    <button
      class="meta-box__header"
      @click="$emit('toggle')"
    >
      <span class="meta-box__title">{{ title }}</span>
      <svg
        class="meta-box__chevron"
        :class="{ 'meta-box__chevron--open': isOpen }"
        viewBox="0 0 20 20"
      >
        <path d="M5.23 7.21a.75.75 0 011.06.02L10 10.94l3.71-3.71a.75.75 0 111.06 1.06l-4.24 4.24a.75.75 0 01-1.06 0L5.21 8.27a.75.75 0 01.02-1.06z" fill="currentColor" />
      </svg>
    </button>
    
    <div
      v-if="isOpen"
      class="meta-box__body"
    >
      <slot />
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

defineProps({
  title: {
    type: String,
    required: true
  },
  isOpen: {
    type: Boolean,
    default: false
  }
})

defineEmits(['toggle'])
</script>

<style scoped>
.meta-box {
  border: 1px solid #d7d7d7;
  border-radius: 8px;
  background: #fff;
}

.meta-box__header {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0.9rem 1rem;
  background: #f8f9fb;
  border: none;
  border-radius: 8px 8px 0 0;
  font-weight: 600;
  color: #1f2937;
  text-align: right;
}

.meta-box__header:hover {
  background: #eef1f6;
}

.meta-box__title {
  flex: 1;
}

.meta-box__chevron {
  width: 18px;
  height: 18px;
  transition: transform 0.2s ease;
  color: #6b7280;
}

.meta-box__chevron--open {
  transform: rotate(180deg);
}

.meta-box__body {
  border-top: 1px solid #e5e7eb;
  padding: 1rem;
}
</style>
