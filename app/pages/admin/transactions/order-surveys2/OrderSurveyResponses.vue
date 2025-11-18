<template>
  <div class="space-y-6">
    <div class="flex items-center space-x-4 space-x-reverse mb-4">
      <button 
        :class="['px-4 py-2 rounded-lg text-sm font-medium', tab==='list' ? 'bg-blue-600 text-white' : 'bg-gray-100 text-gray-700']"
        @click="tab='list'"
      >لیست پاسخ‌ها</button>
      <button 
        :class="['px-4 py-2 rounded-lg text-sm font-medium', tab==='analysis' ? 'bg-blue-600 text-white' : 'bg-gray-100 text-gray-700']"
        @click="tab='analysis'"
      >تحلیل پاسخ‌ها</button>
    </div>
    <div>
      <ResponseList v-if="tab==='list'" />
      <ResponseAnalysis v-else />
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref } from 'vue'
import ResponseList from './response/ResponseList.vue'
import ResponseAnalysis from './response/ResponseAnalysis.vue'
const tab = ref<'list'|'analysis'>('list')
</script> 
