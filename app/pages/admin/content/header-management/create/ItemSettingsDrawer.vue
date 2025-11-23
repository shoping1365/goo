<template>
  <div v-if="hasAccess" class="drawer-overlay" @click.self="$emit('close')">
    <div class="drawer-panel">
      <h3 class="drawer-title">تنظیمات آیتم</h3>

      <div class="form-row">
        <label>چینش:</label>
        <select v-model="local.align">
          <option value="right">راست</option>
          <option value="center">وسط</option>
          <option value="left">چپ</option>
        </select>
      </div>

      <div class="form-row">
        <label>عرض (%):</label>
        <input v-model.number="local.width" type="number" min="5" max="100" />
      </div>

      <div class="drawer-actions">
        <button class="btn-primary" @click="apply">ذخیره</button>
        <button class="btn-secondary" @click="$emit('close')">انصراف</button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
</script>

<script setup lang="ts">
import { computed, onMounted, reactive, watch } from 'vue';
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

interface ItemSettings {
  align?: 'left' | 'center' | 'right'
  width?: number
}

const props = defineProps<{ modelValue: ItemSettings | undefined }>()
const emit = defineEmits(['update', 'close'])

const local = reactive<ItemSettings>({
  align: props.modelValue?.align ?? 'right',
  width: props.modelValue?.width ?? 25
})

function apply() {
  emit('update', { ...props.modelValue, ...local })
  emit('close')
}
</script>

<style scoped>
.drawer-overlay {
  position: fixed;
  inset: 0;
  background: rgba(0,0,0,.3);
  display:flex;justify-content:flex-end;z-index:1200;
}
.drawer-panel {
  width: 280px;
  background:#fff;
  height:100%;
  padding:20px;
  overflow:auto;
}
.drawer-title {font-weight:700;margin-bottom:16px;}
.form-row{margin-bottom:12px;display:flex;flex-direction:column;font-size:14px;}
.form-row label{margin-bottom:4px;}
.drawer-actions{display:flex;gap:8px;}
.btn-primary{background:#667eea;color:#fff;padding:6px 12px;border-radius:6px;}
.btn-secondary{background:#e9ecef;padding:6px 12px;border-radius:6px;}
</style>
