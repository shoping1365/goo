<template>
  <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden mb-6">
    <!-- Header -->
    <div class="bg-gradient-to-r from-indigo-50 to-purple-50 px-4 py-3 border-b border-gray-200">
      <div class="flex items-center justify-between">
        <div class="flex items-center">
          <div class="w-6 h-6 bg-indigo-500 rounded-md flex items-center justify-center ml-2">
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.707A1 1 0 013 7V4z"></path>
            </svg>
          </div>
          <h3 class="text-sm font-semibold text-gray-900">فیلترها و جستجوی پیشرفته</h3>
        </div>
        <button 
          v-if="hasActiveFilters"
          class="text-xs text-red-600 hover:text-red-800 font-medium transition-colors"
          @click="clearAllFilters"
        >
          پاک کردن همه فیلترها
        </button>
      </div>
    </div>

    <!-- فرم فیلترها -->
    <div class="p-6">
      <form class="grid grid-cols-1 md:grid-cols-6 gap-3" @submit.prevent="applyFilters">
        <!-- جستجوی نام محصول -->
        <div>
          <label class="block text-xs font-medium text-gray-700 mb-1">نام محصول</label>
          <input 
            v-model="filters.productName"
            type="text" 
            class="block w-full pl-3 pr-3 py-1.5 border border-gray-300 rounded-md leading-5 bg-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-xs shadow-sm" 
            placeholder="جستجو..."
            dir="rtl"
            @input="onFilterChange"
          />
        </div>

        <!-- جستجوی شناسه محصول -->
        <div>
          <label class="block text-xs font-medium text-gray-700 mb-1">شناسه محصول</label>
          <input
            v-model="filters.productId"
            type="text"
            class="block w-full pl-3 pr-3 py-1.5 border border-gray-300 rounded-md leading-5 bg-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-xs shadow-sm"
            placeholder="شناسه..."
            dir="rtl"
            @input="onFilterChange"
          />
        </div>

        <!-- انتخاب دسته‌بندی -->
        <div>
          <label class="block text-xs font-medium text-gray-700 mb-1">دسته‌بندی</label>
          <SearchableSelect
            v-model="filters.categoryId"
            :options="categoryOptions"
            placeholder="انتخاب یا جستجوی دسته‌بندی..."
            label-key="name"
            value-key="id"
            @update:model-value="onFilterChange"
          />
        </div>

        <!-- انتخاب برند -->
        <div>
          <label class="block text-xs font-medium text-gray-700 mb-1">برند</label>
          <SearchableSelect
            v-model="filters.brandId"
            :options="brandOptions"
            placeholder="انتخاب یا جستجوی برند..."
            label-key="name"
            value-key="id"
            @update:model-value="onFilterChange"
          />
        </div>

        <!-- وضعیت انتشار -->
        <div>
          <label class="block text-xs font-medium text-gray-700 mb-1">وضعیت</label>
          <select 
            v-model="filters.published"
            class="block w-full pl-3 pr-3 py-1.5 border border-gray-300 rounded-md leading-5 bg-white focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-xs shadow-sm"
            @change="onFilterChange"
          >
            <option value="">همه</option>
            <option value="true">منتشر شده</option>
            <option value="false">پیش‌نویس</option>
          </select>
        </div>

        <!-- وضعیت موجودی -->
        <div>
          <label class="block text-xs font-medium text-gray-700 mb-1">موجودی</label>
          <select 
            v-model="filters.stockStatus"
            class="block w-full pl-3 pr-3 py-1.5 border border-gray-300 rounded-md leading-5 bg-white focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-xs shadow-sm"
            @change="onFilterChange"
          >
            <option value="">همه</option>
            <option value="in_stock">موجود (بیش از 10)</option>
            <option value="low_stock">کم‌موجود (1-10)</option>
            <option value="out_of_stock">ناموجود (0)</option>
          </select>
        </div>
      </form>

      <!-- Active Filters Display -->
      <div v-if="hasActiveFilters" class="mt-3 pt-3 border-t border-gray-200">
        <div class="flex items-center gap-2 flex-wrap">
          <span class="text-xs font-medium text-gray-600">فیلترهای فعال:</span>
          
          <span v-if="filters.productName" class="inline-flex items-center px-2 py-1 bg-blue-100 text-blue-800 text-xs rounded-md">
            نام: {{ filters.productName }}
            <button class="mr-1 text-blue-600 hover:text-blue-800" @click="clearFilter('productName')">×</button>
          </span>
          
          <span v-if="filters.categoryId" class="inline-flex items-center px-2 py-1 bg-green-100 text-green-800 text-xs rounded-md">
            دسته: {{ getCategoryName(filters.categoryId) }}
            <button class="mr-1 text-green-600 hover:text-green-800" @click="clearFilter('categoryId')">×</button>
          </span>
          
          <span v-if="filters.brandId" class="inline-flex items-center px-2 py-1 bg-purple-100 text-purple-800 text-xs rounded-md">
            برند: {{ getBrandName(filters.brandId) }}
            <button class="mr-1 text-purple-600 hover:text-purple-800" @click="clearFilter('brandId')">×</button>
          </span>
          
          <span v-if="filters.published" class="inline-flex items-center px-2 py-1 bg-yellow-100 text-yellow-800 text-xs rounded-md">
            وضعیت: {{ filters.published === 'true' ? 'منتشر شده' : 'پیش‌نویس' }}
            <button class="mr-1 text-yellow-600 hover:text-yellow-800" @click="clearFilter('published')">×</button>
          </span>
          
          <span v-if="filters.stockStatus" class="inline-flex items-center px-2 py-1 bg-red-100 text-red-800 text-xs rounded-md">
            موجودی: {{ getStockStatusLabel(filters.stockStatus) }}
            <button class="mr-1 text-red-600 hover:text-red-800" @click="clearFilter('stockStatus')">×</button>
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, computed, watch } from 'vue'
import SearchableSelect from '~/components/admin/common/SearchableSelect.vue'

const emit = defineEmits(['filtersChanged'])

// Filter state
const filters = reactive({
  productName: '',
  productId: '',
  categoryId: '',
  brandId: '',
  published: '',
  stockStatus: ''
})

// Data state
const categories = ref([])
const brands = ref([])
const isLoading = ref(false)

// Computed properties
const hasActiveFilters = computed(() => {
  return filters.productName || 
         filters.productId || 
         filters.categoryId || 
         filters.brandId || 
         filters.published || 
         filters.stockStatus
})

const categoryOptions = computed(() => {
  const allOption = { id: '', name: 'همه دسته‌ها' }
  const categoryList = categories.value.map(category => ({
    id: category.id,
    name: category.name
  }))
  return [allOption, ...categoryList]
})

const brandOptions = computed(() => {
  const allOption = { id: '', name: 'همه برندها' }
  const brandList = brands.value.map(brand => ({
    id: brand.id,
    name: brand.name
  }))
  return [allOption, ...brandList]
})

// Helper functions
const getCategoryName = (id) => {
  const category = categories.value.find(c => c.id === parseInt(id))
  return category?.name || 'نامشخص'
}

const getBrandName = (id) => {
  const brand = brands.value.find(b => b.id === parseInt(id))
  return brand?.name || 'نامشخص'
}

const getStockStatusLabel = (status) => {
  const labels = {
    'in_stock': 'موجود',
    'low_stock': 'کم‌موجود',
    'out_of_stock': 'ناموجود'
  }
  return labels[status] || status
}

// Methods
const loadCategories = async () => {
  try {
    const response = await $fetch('/api/product-categories?all=1')
    categories.value = Array.isArray(response) ? response : (response?.data || [])
  } catch (error) {
    console.error('خطا در دریافت دسته‌بندی‌ها:', error)
    categories.value = []
  }
}

const loadBrands = async () => {
  try {
    const response = await $fetch('/api/brands')
    brands.value = Array.isArray(response) ? response : (response?.data || [])
  } catch (error) {
    console.error('خطا در دریافت برندها:', error)
    brands.value = []
  }
}

const onFilterChange = () => {
  emit('filtersChanged', { ...filters })
}

const applyFilters = () => {
  emit('filtersChanged', { ...filters })
}

const clearFilter = (filterName) => {
  filters[filterName] = ''
  onFilterChange()
}

const clearAllFilters = () => {
  Object.keys(filters).forEach(key => {
    filters[key] = ''
  })
  onFilterChange()
}

// Load data on mount
onMounted(async () => {
  isLoading.value = true
  try {
    console.log('🔄 شروع بارگذاری دسته‌بندی‌ها و برندها...')
    console.log('🔧 SearchableSelect component:', typeof SearchableSelect)
    
    await Promise.all([
      loadCategories(),
      loadBrands()
    ])
    
    console.log('✅ دسته‌بندی‌ها:', categories.value.length, categories.value)
    console.log('✅ برندها:', brands.value.length, brands.value)
    console.log('✅ گزینه‌های دسته‌بندی:', categoryOptions.value)
    console.log('✅ گزینه‌های برند:', brandOptions.value)
    
  } catch (error) {
    console.error('❌ خطا در بارگذاری داده‌ها:', error)
  } finally {
    isLoading.value = false
  }
})

// Watch for filter changes with debouncing for product name
let debounceTimer = null
watch(() => filters.productName, () => {
  if (debounceTimer) {
    clearTimeout(debounceTimer)
  }
  debounceTimer = setTimeout(() => {
    onFilterChange()
  }, 300) // 300ms debounce for product name search
})
</script>

<style scoped>
/* استایل‌های ساده */
</style> 
