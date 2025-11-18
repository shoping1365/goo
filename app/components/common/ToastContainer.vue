<template>
  <div class="fixed top-6 right-4 z-50 space-y-2">
    <div v-for="toast in toasts" :key="toast.id" class="transition-all duration-300">
      <Toast
        :message="toast.message"
        :type="toast.type"
        :duration="toast.duration"
        @close="removeToast(toast.id)"
      />
    </div>
  </div>
</template>

<script setup>
import Toast from '@/components/common/Toast.vue'
import { useToast } from '~/composables/useToast'

const { toasts, removeToast } = useToast()
</script> 
