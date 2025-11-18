<template>
  <button @click="exportExcel" title="اکسپورت به Excel"
    class="bg-green-500 hover:bg-green-600 text-white rounded-lg p-2 transition-colors duration-150 shadow flex items-center justify-center"
  >
    <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16l4-4 4 4m0 0V4m0 12v4a2 2 0 01-2 2H6a2 2 0 01-2-2V6a2 2 0 012-2h7" />
    </svg>
  </button>
</template>
<script setup lang="ts">
import { useUserStore } from '~/stores/userStore';
const userStore = useUserStore();
function exportExcel() {
  userStore.exportExcel();
}
</script> 
