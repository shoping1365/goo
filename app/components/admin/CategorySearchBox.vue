<template>
  <div class="relative">
    <div class="relative">
      <input
        v-model="searchTerm"
        @input="handleSearch"
        @focus="showDropdown = true"
        @blur="handleBlur"
        type="text"
        :placeholder="placeholder"
        class="w-full px-3 py-2 pr-10 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 text-sm"
      />
      <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
        <svg class="h-4 w-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
        </svg>
      </div>
      <button
        v-if="selectedCategory"
        @click="clearSelection"
        class="absolute inset-y-0 right-0 pr-3 flex items-center text-gray-400 hover:text-gray-600"
      >
        <svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
        </svg>
      </button>
    </div>

    <!-- Dropdown -->
    <div
      v-if="showDropdown && (searchTerm || categories.length > 0)"
      class="absolute z-50 w-full mt-1 bg-white border border-gray-300 rounded-md shadow-lg max-h-60 overflow-y-auto"
    >
      <div v-if="loading" class="px-3 py-2 text-sm text-gray-500 text-center">
        در حال جستجو...
      </div>
      
      <div v-else-if="filteredCategories.length === 0 && searchTerm" class="px-3 py-2 text-sm text-gray-500 text-center">
        دسته‌بندی یافت نشد
      </div>
      
      <div v-else>
        <div
          v-for="category in filteredCategories"
          :key="category.id"
          @click="selectCategory(category)"
          class="px-3 py-2 hover:bg-gray-100 cursor-pointer text-sm border-b border-gray-100 last:border-b-0"
        >
          <div class="flex items-center justify-between">
            <div>
              <div class="font-medium text-gray-900">{{ category.name }}</div>
              <div v-if="category.parent_name && category.parent_name !== '-'" class="text-xs text-gray-500">
                زیرمجموعه: {{ category.parent_name }}
              </div>
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <span
                class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium"
                :class="category.parent_id ? 'bg-green-100 text-green-800' : 'bg-purple-100 text-purple-800'"
              >
                {{ category.parent_id ? 'فرعی' : 'اصلی' }}
              </span>
              <span v-if="category.product_count > 0" class="text-xs text-gray-500">
                {{ category.product_count }} محصول
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch, onMounted } from 'vue'

interface Category {
  id: number
  name: string
  parent_id?: number
  parent_name?: string
  product_count?: number
  published?: boolean
}

const props = defineProps<{
  placeholder?: string
  value?: Category | null
}>()

const emit = defineEmits<{
  'update:value': [category: Category | null]
  'change': [category: Category | null]
}>()

const searchTerm = ref('')
const showDropdown = ref(false)
const loading = ref(false)
const categories = ref<Category[]>([])
const selectedCategory = ref<Category | null>(props.value || null)

// Watch for external value changes
watch(() => props.value, (newValue) => {
  selectedCategory.value = newValue
  if (newValue) {
    searchTerm.value = newValue.name
  } else {
    searchTerm.value = ''
  }
}, { immediate: true })

// Computed filtered categories
const filteredCategories = computed(() => {
  if (!searchTerm.value.trim()) {
    return categories.value.slice(0, 10) // Show first 10 when no search
  }
  
  const term = searchTerm.value.toLowerCase()
  return categories.value.filter(category => 
    category.name.toLowerCase().includes(term) ||
    (category.parent_name && category.parent_name.toLowerCase().includes(term))
  )
})

// Search handler with debounce
let searchTimeout: NodeJS.Timeout
const handleSearch = () => {
  clearTimeout(searchTimeout)
  searchTimeout = setTimeout(() => {
    if (searchTerm.value.trim()) {
      fetchCategories()
    }
  }, 300)
}

// Handle blur with delay to allow click
const handleBlur = () => {
  setTimeout(() => {
    showDropdown.value = false
  }, 200)
}

// Fetch categories
const fetchCategories = async () => {
  if (loading.value) return
  
  loading.value = true
  try {
    const response = await $fetch('/api/admin/product-categories?all=1')
    const raw = (response as any).data || []
    
    // Add parent_name for display
    raw.forEach((cat: any) => {
      if (cat.parent_id) {
        const parent = raw.find((c: any) => c.id === cat.parent_id)
        cat.parent_name = parent ? parent.name : '-'
      } else {
        cat.parent_name = '-'
      }
    })
    
    categories.value = raw
  } catch (error) {
    console.error('خطا در دریافت دسته‌بندی‌ها:', error)
    categories.value = []
  } finally {
    loading.value = false
  }
}

// Select category
const selectCategory = (category: Category) => {
  selectedCategory.value = category
  searchTerm.value = category.name
  showDropdown.value = false
  
  emit('update:value', category)
  emit('change', category)
}

// Clear selection
const clearSelection = () => {
  selectedCategory.value = null
  searchTerm.value = ''
  showDropdown.value = false
  
  emit('update:value', null)
  emit('change', null)
}

// Load initial categories
onMounted(() => {
  fetchCategories()
})
</script>
