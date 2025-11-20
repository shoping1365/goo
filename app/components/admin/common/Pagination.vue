<template>
  <div class="bg-white px-4 py-3 border-t border-gray-200 sm:px-6">
    <div class="flex items-center justify-between">
      <div class="flex-1 flex justify-between sm:hidden">
        <button 
          :disabled="currentPage === 1"
          class="relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
          @click="previousPage"
        >
          قبلی
        </button>
        <button 
          :disabled="currentPage === totalPages"
          class="mr-3 relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
          @click="nextPage"
        >
          بعدی
        </button>
      </div>
            <div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-end">
 
          <div class="flex items-center space-x-4">
          <!-- انتخاب تعداد آیتم در هر صفحه -->
    <div class="flex items-center space-x-2">
      <span class="text-sm text-gray-700">نمایش:</span>
      <select 
        :value="props.perPage"
        class="text-sm border border-gray-300 rounded-md px-2 py-1 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
        @change="handleSelectChange"
      >
        <option :value="5">5</option>
        <option :value="10">10</option>
        <option :value="20">20</option>
        <option :value="50">50</option>
        <option :value="100">100</option>
      </select>
    </div>
          
          <!-- دکمه‌های صفحه‌بندی -->
          <nav class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px">
            <button 
              :disabled="currentPage === 1"
              class="relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
              @click="previousPage"
            >
              <svg class="h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"></path>
              </svg>
            </button>
            
            <button 
              v-for="page in visiblePages" 
              :key="page"
              :class="[
                page === currentPage 
                  ? 'bg-blue-50 border-blue-500 text-blue-600' 
                  : 'bg-white border-gray-300 text-gray-500 hover:bg-gray-50',
                'relative inline-flex items-center px-4 py-2 border text-sm font-medium'
              ]"
              @click="goToPage(page)"
            >
              {{ page }}
            </button>
            
            <button 
              :disabled="currentPage === totalPages"
              class="relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
              @click="nextPage"
            >
              <svg class="h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
              </svg>
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
  perPage?: number
}>()

const emit = defineEmits(['page-changed', 'per-page-changed'])

// Handle select change
const handleSelectChange = (event: Event) => {
  const target = event.target as HTMLSelectElement
  const value = Number(target.value)
  emit('per-page-changed', value)
}

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
  const ipp = props.perPage || 10
  const from = (props.currentPage - 1) * ipp + 1
  const to = Math.min(props.currentPage * ipp, props.total)
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

