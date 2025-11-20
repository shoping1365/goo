<template>
  <Transition
    enter-active-class="transition ease-out duration-300"
    enter-from-class="transform translate-y-2 opacity-0 sm:translate-y-0 sm:translate-x-2"
    enter-to-class="transform translate-y-0 opacity-100 sm:translate-x-0"
    leave-active-class="transition ease-in duration-100"
    leave-from-class="opacity-100"
    leave-to-class="opacity-0"
  >
    <div
      v-if="show"
      class="fixed top-6 right-4 z-50 max-w-sm w-full bg-white shadow-lg rounded-lg pointer-events-auto ring-1 ring-black ring-opacity-5 overflow-hidden"
    >
      <div class="p-6">
        <div class="flex items-start">
          <div class="flex-shrink-0">
            <div
              :class="[
                'w-8 h-8 rounded-full flex items-center justify-center',
                type === 'success' ? 'bg-green-100' : 'bg-red-100'
              ]"
            >
              <span
                :class="[
                  type === 'success' ? 'i-heroicons-check text-green-600' : 'i-heroicons-x-mark text-red-600'
                ]"
              ></span>
            </div>
          </div>
          <div class="mr-3 w-0 flex-1 pt-0.5">
            <p class="text-sm font-medium text-gray-900">{{ title }}</p>
            <p class="mt-1 text-sm text-gray-500">{{ message }}</p>
          </div>
          <div class="mr-4 flex flex-shrink-0">
            <button
              class="bg-white rounded-md inline-flex text-gray-400 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
              @click="close"
            >
              <span class="sr-only">بستن</span>
              <span class="i-heroicons-x-mark text-lg"></span>
            </button>
          </div>
        </div>
      </div>
    </div>
  </Transition>
</template>

<script setup lang="ts">
import { watch } from 'vue'

interface Props {
  show: boolean
  type: 'success' | 'error'
  title: string
  message: string
  duration?: number
}

const props = withDefaults(defineProps<Props>(), {
  duration: 5000
})

const emit = defineEmits<{
  close: []
}>()

const close = () => {
  emit('close')
}

// بستن خودکار پس از مدت زمان مشخص شده - بهینه‌سازی شده
let timeoutId: NodeJS.Timeout | null = null

watch(() => props.show, (newShow) => {
  if (timeoutId) {
    clearTimeout(timeoutId)
    timeoutId = null
  }
  
  if (newShow && props.duration > 0) {
    timeoutId = setTimeout(() => {
      close()
    }, props.duration)
  }
})
</script>

<!--
  کامپوننت نمایش پیام‌های موفقیت و خطا
  شامل:
  1. انیمیشن ورود و خروج
  2. آیکون‌های مختلف برای انواع مختلف پیام
  3. بستن خودکار پس از مدت زمان مشخص
  4. طراحی مدرن و زیبا
  5. پشتیبانی از انواع success و error
--> 