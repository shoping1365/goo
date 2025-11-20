<template>
  <Teleport to="body">
    <Transition name="confirm-dialog">
      <div v-if="isVisible" class="fixed inset-0 z-50 flex items-center justify-center">
        <!-- Backdrop -->
        <div 
          class="absolute inset-0 bg-black/50 backdrop-blur-sm"
          @click="handleBackdropClick"
        ></div>
        
        <!-- Dialog -->
        <div class="relative bg-white dark:bg-gray-800 rounded-lg shadow-xl max-w-md w-full mx-4 px-4 py-4">
          <!-- Header -->
          <div class="flex items-center justify-between mb-4">
            <button
              class="text-gray-400 hover:text-gray-600 dark:hover:text-gray-300 transition-colors"
              @click="handleCancel"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
              </svg>
            </button>
            <h3 class="text-lg font-semibold text-gray-900 dark:text-white">
              {{ currentTitle }}
            </h3>
          </div>
          
          <!-- Content -->
          <div class="mb-6">
            <p class="text-gray-700 dark:text-gray-300 whitespace-pre-line text-right">
              {{ currentMessage }}
            </p>
          </div>
          
          <!-- Actions -->
          <div class="flex justify-between space-x-3 rtl:space-x-reverse">
            <button
              class="px-4 py-2 text-sm font-medium text-gray-700 dark:text-gray-300 bg-gray-100 dark:bg-gray-700 hover:bg-gray-200 dark:hover:bg-gray-600 rounded-md transition-colors"
              @click="handleCancel"
            >
              {{ currentCancelText }}
            </button>
            <button
              class="px-4 py-2 text-sm font-medium text-white bg-red-600 hover:bg-red-700 rounded-md transition-colors"
              @click="handleConfirm"
            >
              {{ currentConfirmText }}
            </button>
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup lang="ts">
import { ref } from 'vue'

interface ConfirmDialogOptions {
  title?: string
  message: string
  confirmText?: string
  cancelText?: string
  type?: 'danger' | 'warning' | 'info'
}

// Props برای تنظیم مقادیر پیش‌فرض
const defaultTitle = ref('تأیید عملیات')
const defaultConfirmText = ref('تأیید')
const defaultCancelText = ref('انصراف')
const defaultType = ref<'danger' | 'warning' | 'info'>('danger')

// State برای تنظیم مقادیر dynamic
const currentTitle = ref('')
const currentMessage = ref('')
const currentConfirmText = ref('')
const currentCancelText = ref('')
const currentType = ref<'danger' | 'warning' | 'info'>('danger')

// Emits
const emit = defineEmits<{
  (e: 'confirm'): void
  (e: 'cancel'): void
}>()

const isVisible = ref(false)
let resolvePromise: ((value: boolean) => void) | null = null

// تابع اصلی که جایگزین confirm() می‌شود
const show = (options: ConfirmDialogOptions): Promise<boolean> => {
  return new Promise((resolve) => {
    currentTitle.value = options.title || defaultTitle.value
    currentMessage.value = options.message
    currentConfirmText.value = options.confirmText || defaultConfirmText.value
    currentCancelText.value = options.cancelText || defaultCancelText.value
    currentType.value = options.type || defaultType.value
    resolvePromise = resolve
    isVisible.value = true
  })
}

// تابع‌های مدیریت رویدادها
const handleConfirm = () => {
  isVisible.value = false
  if (resolvePromise) {
    resolvePromise(true)
    resolvePromise = null
  }
  emit('confirm')
}

const handleCancel = () => {
  isVisible.value = false
  if (resolvePromise) {
    resolvePromise(false)
    resolvePromise = null
  }
  emit('cancel')
}

const handleBackdropClick = () => {
  handleCancel()
}

// در معرض قرار دادن تابع show
defineExpose({
  show
})
</script>

<style scoped>
.confirm-dialog-enter-active,
.confirm-dialog-leave-active {
  transition: all 0.2s ease;
}

.confirm-dialog-enter-from,
.confirm-dialog-leave-to {
  opacity: 0;
  transform: scale(0.95);
}

.confirm-dialog-enter-to,
.confirm-dialog-leave-from {
  opacity: 1;
  transform: scale(1);
}
</style>
