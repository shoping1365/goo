<template>
  <div v-if="modelValue" class="fixed inset-0 z-50 flex items-center justify-center">
    <div class="bg-white rounded-lg shadow-lg p-8 w-full max-w-4xl min-h-[420px] min-h-[60vh] relative flex flex-col">
      <button @click="$emit('update:modelValue', false)" class="absolute left-4 top-6 text-gray-400 hover:text-gray-600 text-xl">&times;</button>
      <h3 class="text-lg font-bold mb-4 text-right">{{ title }}</h3>
      <div class="flex-1 flex flex-col">
        <slot></slot>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  modelValue: boolean;
  title: string;
}>();

defineEmits<{
  (e: 'update:modelValue', value: boolean): void;
}>();
</script> 
