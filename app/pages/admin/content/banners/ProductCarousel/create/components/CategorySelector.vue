<template>
  <div class="relative">
    <button
      type="button"
      class="w-full px-3 py-2 border border-gray-300 rounded-md bg-white text-left focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-purple-500 text-gray-900 font-medium"
      :class="{ 'border-purple-500 ring-2 ring-purple-500': showDropdown }"
      @click="toggleDropdown"
    >
      <div class="flex items-center justify-between">
        <span class="text-gray-900">
          {{ selectedCategoryName || placeholder }}
        </span>
        <i
          class="fas fa-chevron-down transition-transform duration-200"
          :class="{ 'rotate-180': showDropdown }"
        ></i>
      </div>
    </button>

    <!-- Dropdown -->
    <div
      v-if="showDropdown"
      class="absolute z-50 w-full mt-1 bg-white border border-gray-300 rounded-md shadow-lg max-h-60 overflow-auto"
    >
      <!-- جستجو -->
      <div class="p-2 border-b border-gray-200">
        <div class="relative">
          <input
            v-model="searchQuery"
            type="text"
            placeholder="جستجوی دسته‌بندی..."
            class="w-full px-3 py-2 pl-10 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 text-gray-900 placeholder-gray-600 font-medium"
          />
          <i class="fas fa-search absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400"></i>
        </div>
      </div>

      <!-- گزینه همه محصولات -->
      <div
        class="px-4 py-3 cursor-pointer hover:bg-gray-50 border-b border-gray-200"
        :class="{ 'bg-purple-50 text-purple-700': modelValue === null }"
        @click="selectCategory(null)"
      >
        <div class="flex items-center">
          <i class="fas fa-th-large ml-2 text-gray-400"></i>
          <span class="font-medium text-gray-900">{{ placeholder }}</span>
        </div>
        <p class="text-sm text-gray-700 mt-1">نمایش تمام محصولات بدون فیلتر دسته‌بندی</p>
      </div>

      <!-- لیست دسته‌بندی‌ها -->
      <div v-for="category in filteredCategories" :key="category.id">
        <div
          class="px-4 py-3 cursor-pointer hover:bg-gray-50"
          :class="{ 'bg-purple-50 text-purple-700': modelValue === category.id }"
          @click="selectCategory(category.id)"
        >
          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <i class="fas fa-folder ml-2 text-gray-400"></i>
              <span class="font-medium text-gray-900">{{ category.name }}</span>
            </div>
            <span v-if="category.product_count" class="text-xs bg-gray-100 px-2 py-1 rounded-full">
              {{ category.product_count }} محصول
            </span>
          </div>
          <p v-if="category.description" class="text-sm text-gray-700 mt-1">{{ category.description }}</p>
        </div>
      </div>

      <!-- پیام عدم وجود نتیجه -->
      <div v-if="filteredCategories.length === 0 && searchQuery.trim()" class="px-4 py-6 text-center text-gray-700">
        <i class="fas fa-search text-2xl mb-2"></i>
        <p>دسته‌بندی یافت نشد</p>
        <p class="text-sm">برای عبارت "{{ searchQuery }}" نتیجه‌ای پیدا نشد</p>
      </div>
    </div>

    <!-- Overlay برای بستن dropdown -->
    <div
      v-if="showDropdown"
      class="fixed inset-0 z-40"
      @click="closeDropdown"
    ></div>

    <!-- نمایش دسته‌بندی انتخاب شده -->
    <div v-if="selectedCategory" class="mt-2 p-3 bg-blue-50 rounded-md border border-blue-200">
      <div class="flex items-center justify-between">
        <div class="flex items-center">
          <i class="fas fa-check-circle text-green-500 ml-2"></i>
          <span class="text-sm font-medium text-gray-900">دسته‌بندی انتخاب شده:</span>
        </div>
        <button
          class="text-blue-600 hover:text-blue-800 text-sm"
          @click="selectCategory(null)"
        >
          <i class="fas fa-times"></i>
        </button>
      </div>
      <p class="text-sm text-gray-800 mt-1">{{ selectedCategory.name }}</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'

interface Props {
  modelValue: number | null
  categories: any[]
  placeholder?: string
}

interface Emits {
  (e: 'update:modelValue', value: number | null): void
}

const props = defineProps<Props>()
const emit = defineEmits<Emits>()

// State
const showDropdown = ref(false)
const searchQuery = ref('')// محاسبه دسته‌بندی فیلتر شده
const filteredCategories = computed(() => {
  if (!searchQuery.value || searchQuery.value.trim() === '') {
    return props.categories || []
  }

  const searchTerm = searchQuery.value.toLowerCase().trim()

  if (!props.categories || !Array.isArray(props.categories)) {
    return []
  }

  const filtered = props.categories.filter(category => {
    if (!category) return false

    const nameMatch = category.name && typeof category.name === 'string' &&
                     category.name.toLowerCase().includes(searchTerm)
    const descMatch = category.description && typeof category.description === 'string' &&
                     category.description.toLowerCase().includes(searchTerm)
    return nameMatch || descMatch
  })

  return filtered
})

// محاسبه نام دسته‌بندی انتخاب شده
const selectedCategoryName = computed(() => {
  if (props.modelValue === null) {
    return props.placeholder || 'انتخاب کنید'
  }

  const category = props.categories.find(c => c.id === props.modelValue)
  return category ? category.name : 'دسته‌بندی نامشخص'
})

// محاسبه دسته‌بندی انتخاب شده
const selectedCategory = computed(() => {
  return props.categories.find(c => c.id === props.modelValue)
})

// متدهای مدیریت dropdown
const toggleDropdown = () => {
  showDropdown.value = !showDropdown.value
  if (showDropdown.value) {
    searchQuery.value = ''
  }
}

const closeDropdown = () => {
  showDropdown.value = false
  searchQuery.value = ''
}

const selectCategory = (categoryId: number | null) => {
  emit('update:modelValue', categoryId)
  closeDropdown()
}
</script>

<style scoped>
.category-dropdown {
  /* Dropdown specific styles */
  position: relative;
  width: 100%;
}

.category-item {
  transition: all 0.2s ease;
  padding: 12px 16px;
  cursor: pointer;
  border-bottom: 1px solid #f3f4f6;
}

/* جستجوی دسته‌بندی */
.category-search {
  position: relative;
}

.category-search i {
  position: absolute;
  left: 12px;
  top: 50%;
  transform: translateY(-50%);
  color: #9ca3af;
}

/* لیست دسته‌بندی‌ها */
.categories-list {
  max-height: 300px;
  overflow-y: auto;
}

.category-item {
  padding: 12px 16px;
  cursor: pointer;
  transition: all 0.2s ease;
  border-bottom: 1px solid #f3f4f6;
}

.category-item:hover {
  background-color: #f9fafb;
}

.category-item.selected {
  background-color: #ddd6fe;
  color: #6d28d9;
}

.category-item .category-info {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.category-item .category-name {
  font-weight: 500;
}

.category-item .category-count {
  background-color: #e5e7eb;
  padding: 2px 8px;
  border-radius: 12px;
  font-size: 12px;
  color: #6b7280;
}

.category-item.selected .category-count {
  background-color: #c4b5fd;
  color: #6d28d9;
}

.category-item:hover {
  background-color: #f9fafb;
}

.category-item.selected {
  background-color: #e0e7ff;
  color: #3730a3;
}

.category-search {
  border-bottom: 1px solid #e5e7eb;
}

.category-search input {
  border: 1px solid #d1d5db;
  border-radius: 6px;
}

.category-search input:focus {
  border-color: #8b5cf6;
  box-shadow: 0 0 0 3px rgba(139, 92, 246, 0.1);
}

.category-count {
  background-color: #f3f4f6;
  color: #6b7280;
  padding: 2px 6px;
  border-radius: 12px;
  font-size: 11px;
}

.selected-category-card {
  background-color: #eff6ff;
  border-color: #3b82f6;
}

.selected-category-card .selected-category-name {
  color: #1d4ed8;
}

/* Animation classes */
.dropdown-enter-active,
.dropdown-leave-active {
  transition: all 0.3s ease;
}

.dropdown-enter-from,
.dropdown-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}

/* Custom scrollbar for dropdown */
.category-dropdown::-webkit-scrollbar {
  width: 6px;
}

.category-dropdown::-webkit-scrollbar-track {
  background: #f1f5f9;
}

.category-dropdown::-webkit-scrollbar-thumb {
  background: #cbd5e1;
  border-radius: 3px;
}

.category-dropdown::-webkit-scrollbar-thumb:hover {
  background: #94a3b8;
}

/* بهبود وضوح متن در فیلدهای ورودی */
input, button {
  color: #111827 !important; /* text-gray-900 */
  font-weight: 500 !important;
}

input::placeholder {
  color: #4b5563 !important; /* text-gray-600 */
  font-weight: 400 !important;
}

/* بهبود وضوح متن در فیلدهای focus */
input:focus, button:focus {
  color: #111827 !important;
  font-weight: 500 !important;
}
</style>
