<template>
  <button
    v-if="hasPermission('user.create')"
    type="button"
    class="bg-green-200 hover:bg-green-300 text-green-900 px-5 py-2 rounded-lg flex items-center justify-center shadow font-bold transition-colors duration-150 border border-green-300 backdrop-blur-sm w-40 relative"
    style="box-shadow: 0 2px 12px 0 rgba(34,197,94,0.10);"
    @click="openNewUserModal"
  >
    
    <span class="mx-auto w-full text-center">کاربر جدید</span>
  </button>
</template>

<script setup lang="ts">
import { useAuth } from '~/composables/useAuth';

// استفاده از کامپوزابل احراز هویت برای بررسی پرمیژن‌ها
const { hasPermission } = useAuth()

const emit = defineEmits(['open-modal']);

function openNewUserModal() {
  emit('open-modal');
}
</script> 
