<template>
  <div class="px-4 py-3 border-t border-gray-200 sm:px-6">
    <div class="flex items-center justify-between">
      <div class="flex-1 flex justify-between sm:hidden">
        <!-- بخش موبایل حذف شد -->
      </div>
      <div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
        <!-- اطلاعات صفحه‌بندی -->
        <div class="flex items-center text-sm text-gray-700">
          <span>نمایش {{ paginationInfo.from }} تا {{ paginationInfo.to }} از {{ total }} نتیجه</span>
        </div>
        
        <!-- دکمه‌های صفحه‌بندی -->
        <div class="flex items-center space-x-2">
          <!-- انتخاب تعداد آیتم در هر صفحه -->
          <div class="flex items-center space-x-2">
            <span class="text-sm text-gray-700">نمایش:</span>
            <select 
              :value="itemsPerPage || 10"
              class="text-sm border border-gray-300 rounded-md px-2 py-1 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              @change="(event) => $emit('items-per-page-changed', parseInt((event.target as HTMLSelectElement).value))"
            >
              <option v-for="option in itemsPerPageOptions" :key="option" :value="option">
                {{ option }}
              </option>
            </select>
          </div>
          
          <!-- دکمه‌های ناوبری -->
          <nav class="flex items-center space-x-1">
            <!-- دکمه صفحه قبل -->
            <button
              :disabled="currentPage === 1"
              :class="[
                'px-2 py-1 text-sm font-medium rounded-md border',
                currentPage === 1 
                  ? 'text-gray-400 border-gray-300 cursor-not-allowed bg-gray-50' 
                  : 'text-gray-700 border-gray-300 hover:bg-gray-50 hover:text-gray-900'
              ]"
              @click="previousPage"
            >
              قبلی
            </button>
            
            <!-- شماره صفحات -->
            <template v-for="page in visiblePages" :key="page">
              <button
                :class="[
                  'px-3 py-1 text-sm font-medium rounded-md border',
                  page === currentPage
                    ? 'text-blue-600 border-blue-600 bg-blue-50'
                    : 'text-gray-700 border-gray-300 hover:bg-gray-50 hover:text-gray-900'
                ]"
                @click="goToPage(page)"
              >
                {{ page }}
              </button>
            </template>
            
            <!-- دکمه صفحه بعد -->
            <button
              :disabled="currentPage === totalPages"
              :class="[
                'px-2 py-1 text-sm font-medium rounded-md border',
                currentPage === totalPages 
                  ? 'text-gray-400 border-gray-300 cursor-not-allowed bg-gray-50' 
                  : 'text-gray-700 border-gray-300 hover:bg-gray-50 hover:text-gray-900'
              ]"
              @click="nextPage"
            >
              بعدی
            </button>
          </nav>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  currentPage: number
  totalPages: number
  total: number
  itemsPerPage?: number
}>()

const emit = defineEmits(['page-changed', 'items-per-page-changed'])

// گزینه‌های تعداد آیتم در هر صفحه
const itemsPerPageOptions = [5, 10, 20, 50, 100]

// محاسبه صفحات قابل مشاهده
const visiblePages = computed(() => {
  const pages = []
  const maxVisible = 5
  let start = Math.max(1, props.currentPage - Math.floor(maxVisible / 2))
  const end = Math.min(props.totalPages, start + maxVisible - 1)
  
  if (end - start + 1 < maxVisible) {
    start = Math.max(1, end - maxVisible + 1)
  }
  
  for (let i = start; i <= end; i++) {
    pages.push(i)
  }
  
  return pages
})

// اطلاعات صفحه‌بندی
const paginationInfo = computed(() => {
  const from = (props.currentPage - 1) * (props.itemsPerPage || 10) + 1
  const to = Math.min(props.currentPage * (props.itemsPerPage || 10), props.total)
  return { from, to }
})

// متدهای صفحه‌بندی
const previousPage = () => {
  if (props.currentPage > 1) {
    emit('page-changed', props.currentPage - 1)
  }
}

const nextPage = () => {
  if (props.currentPage < props.totalPages) {
    emit('page-changed', props.currentPage + 1)
  }
}

const goToPage = (page: number) => {
  emit('page-changed', page)
}
</script> 