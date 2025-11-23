<template>
  <div v-if="hasAccess" class="page-header">
    <div class="header-content">
      <div class="header-title">
        <h1>{{ isEditing ? 'ویرایش فوتر' : 'ایجاد فوتر جدید' }}</h1>
        <p class="header-description">
          {{ isEditing ? 'ویرایش تنظیمات و محتوای فوتر موجود' : 'تنظیم محتوای فوتر جدید برای سایت' }}
        </p>
      </div>
      
      <div class="header-actions">
        <button
          class="btn btn-secondary"
          @click="goBack"
        >
          انصراف
        </button>
        <button
          class="btn btn-primary"
          @click="saveFooter"
        >
          {{ isEditing ? 'به‌روزرسانی فوتر' : 'ذخیره فوتر' }}
        </button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
</script>

<script setup lang="ts">
import { computed, inject, onMounted, watch } from 'vue'
import { useAuth } from '~/composables/useAuth'

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
const checkAuth = async (): Promise<void> => {
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

// Inject functions and data from parent
const isEditing = inject<boolean | undefined>('isEditing')
const saveFooter = inject<(() => void) | undefined>('saveFooter')
const goBack = inject<(() => void) | undefined>('goBack')
</script>

<style scoped>
.page-header {
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  padding: 24px;
  margin-bottom: 24px;
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  gap: 24px;
}

.header-title h1 {
  margin: 0 0 8px 0;
  font-size: 24px;
  font-weight: 700;
  color: #1f2937;
}

.header-description {
  margin: 0;
  color: #6b7280;
  font-size: 14px;
  line-height: 1.5;
}

.header-actions {
  display: flex;
  gap: 12px;
  flex-shrink: 0;
}

.btn {
  padding: 10px 20px;
  border: none;
  border-radius: 6px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
  text-decoration: none;
  display: inline-flex;
  align-items: center;
  justify-content: center;
}

.btn-primary {
  background: #3b82f6;
  color: white;
}

.btn-primary:hover {
  background: #2563eb;
}

.btn-secondary {
  background: #6b7280;
  color: white;
}

.btn-secondary:hover {
  background: #4b5563;
}

@media (max-width: 768px) {
  .header-content {
    flex-direction: column;
    align-items: stretch;
  }
  
  .header-actions {
    justify-content: stretch;
  }
  
  .btn {
    flex: 1;
  }
}
</style>


