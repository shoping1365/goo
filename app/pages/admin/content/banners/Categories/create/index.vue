<template>
  <div class="categories-create-page p-6">
    <div class="max-w-4xl mx-auto">
      <!-- Page Header -->
      <div class="mb-8">
        <div class="flex items-center gap-6 mb-4">
          <NuxtLink 
            to="/admin/content/banners/Categories"
            class="text-blue-600 hover:text-blue-800 transition-colors"
          >
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"/>
            </svg>
          </NuxtLink>
          <h1 class="text-3xl font-bold text-gray-900">ایجاد باکس دسته‌بندی</h1>
        </div>
        <p class="text-gray-600">ایجاد ابزارک جدید برای نمایش دسته‌بندی‌ها</p>
      </div>

      <!-- Create Form -->
      <div class="bg-white rounded-lg shadow-lg p-6">
        <form class="space-y-6" @submit.prevent="handleSubmit">
          <!-- Basic Information -->
          <div class="border-b border-gray-200 pb-6">
            <h2 class="text-xl font-semibold text-gray-800 mb-4">اطلاعات پایه</h2>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">عنوان ابزارک</label>
                <input 
                  v-model="form.title"
                  type="text"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  placeholder="مثال: باکس دسته‌بندی‌های اصلی"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">کد یکتا</label>
                <input 
                  v-model="form.code"
                  type="text"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  placeholder="مثال: categories-box"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">صفحه نمایش</label>
                <select 
                  v-model="form.page"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                >
                  <option value="">انتخاب کنید</option>
                  <option value="home">صفحه اصلی</option>
                  <option value="category">صفحات دسته‌بندی</option>
                  <option value="sidebar">نوار کناری</option>
                  <option value="footer">فوتر</option>
                </select>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">ترتیب نمایش</label>
                <input 
                  v-model.number="form.order"
                  type="number"
                  min="1"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                />
              </div>
            </div>
          </div>

          <!-- Layout Configuration -->
          <div class="border-b border-gray-200 pb-6">
            <h2 class="text-xl font-semibold text-gray-800 mb-4">تنظیمات چیدمان</h2>
            <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">تعداد ستون</label>
                <select 
                  v-model="form.columns"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                >
                  <option value="1">1 ستون</option>
                  <option value="2">2 ستون</option>
                  <option value="3">3 ستون</option>
                  <option value="4">4 ستون</option>
                  <option value="6">6 ستون</option>
                </select>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر تعداد دسته‌بندی</label>
                <input 
                  v-model.number="form.maxCategories"
                  type="number"
                  min="1"
                  max="50"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">ارتفاع آیتم (پیکسل)</label>
                <input 
                  v-model.number="form.itemHeight"
                  type="number"
                  min="60"
                  max="200"
                  step="10"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                />
              </div>
            </div>
            
            <div class="mt-4 grid grid-cols-1 md:grid-cols-3 gap-6">
              <label class="flex items-center">
                <input 
                  v-model="form.showImages"
                  type="checkbox"
                  class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                />
                <span class="ml-2 text-sm text-gray-700">نمایش تصاویر</span>
              </label>
              <label class="flex items-center">
                <input 
                  v-model="form.showCounts"
                  type="checkbox"
                  class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                />
                <span class="ml-2 text-sm text-gray-700">نمایش تعداد محصولات</span>
              </label>
              <label class="flex items-center">
                <input 
                  v-model="form.showDescriptions"
                  type="checkbox"
                  class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                />
                <span class="ml-2 text-sm text-gray-700">نمایش توضیحات</span>
              </label>
            </div>
          </div>

          <!-- Category Selection -->
          <div class="border-b border-gray-200 pb-6">
            <h2 class="text-xl font-semibold text-gray-800 mb-4">انتخاب دسته‌بندی‌ها</h2>
            <div class="mb-4">
              <label class="block text-sm font-medium text-gray-700 mb-2">روش انتخاب دسته‌بندی‌ها</label>
              <select 
                v-model="form.categorySelectionType"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
                <option value="manual">انتخاب دستی</option>
                <option value="featured">دسته‌بندی‌های ویژه</option>
                <option value="popular">دسته‌بندی‌های محبوب</option>
                <option value="latest">آخرین دسته‌بندی‌ها</option>
                <option value="parent">فقط دسته‌بندی‌های اصلی</option>
              </select>
            </div>

            <!-- Manual Category Selection -->
            <div v-if="form.categorySelectionType === 'manual'" class="space-y-4">
              <div class="flex items-center justify-between">
                <h3 class="text-lg font-medium text-gray-800">دسته‌بندی‌های انتخاب شده</h3>
                <button 
                  type="button"
                  class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors"
                  @click="addCategoryItem"
                >
                  افزودن دسته‌بندی
                </button>
              </div>
              
              <div v-if="form.categories.length === 0" class="text-center py-8 text-gray-500">
                هیچ دسته‌بندی انتخاب نشده است
              </div>
              
              <div v-else class="space-y-3">
                <div 
                  v-for="(category, index) in form.categories"
                  :key="index"
                  class="flex items-center gap-3 p-3 bg-gray-50 rounded-lg"
                >
                  <div class="flex-1 grid grid-cols-1 md:grid-cols-4 gap-3">
                    <input 
                      v-model="category.name"
                      type="text"
                      placeholder="نام دسته‌بندی"
                      class="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                    />
                    <input 
                      v-model="category.image"
                      type="text"
                      placeholder="آدرس تصویر"
                      class="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                    />
                    <div class="relative">
                      <input 
                        v-model="category.searchTerm"
                        type="text"
                        placeholder="جستجو در دسته‌بندی‌ها..."
                        class="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                        @focus="category.showDropdown = true"
                        @blur="hideCategoryDropdown(index)"
                      />
                      <div v-if="category.showDropdown && filteredCategories.length > 0" class="absolute z-50 w-full mt-1 bg-white border border-gray-300 rounded-md shadow-lg max-h-60 overflow-y-auto">
                        <div
                          v-for="cat in filteredCategories"
                          :key="cat.id"
                          class="px-3 py-2 hover:bg-gray-100 cursor-pointer text-sm border-b border-gray-100 last:border-b-0"
                          @click="selectCategoryForItem(index, cat)"
                        >
                          <div class="flex items-center justify-between">
                            <div>
                              <div class="font-medium text-gray-900">{{ cat.name }}</div>
                              <div v-if="cat.parent_name && cat.parent_name !== '-'" class="text-xs text-gray-500">
                                زیرمجموعه: {{ cat.parent_name }}
                              </div>
                            </div>
                            <div class="flex items-center space-x-2 space-x-reverse">
                              <span
                                class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium"
                                :class="cat.parent_id ? 'bg-green-100 text-green-800' : 'bg-purple-100 text-purple-800'"
                              >
                                {{ cat.parent_id ? 'فرعی' : 'اصلی' }}
                              </span>
                              <span v-if="typeof cat.product_count === 'number' && cat.product_count > 0" class="text-xs text-gray-500">
                                {{ cat.product_count }} محصول
                              </span>
                            </div>
                          </div>
                        </div>
                      </div>
                    </div>
                    <input 
                      v-model="category.description"
                      type="text"
                      placeholder="توضیحات کوتاه"
                      class="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                    />
                  </div>
                  <button 
                    type="button"
                    class="text-red-600 hover:text-red-800 p-2"
                    @click="removeCategoryItem(index)"
                  >
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"/>
                    </svg>
                  </button>
                </div>
              </div>
            </div>
          </div>

          <!-- Form Actions -->
          <div class="flex gap-3 justify-end">
            <button 
              type="button"
              class="px-6 py-3 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors"
              @click="saveAsDraft"
            >
              ذخیره به عنوان پیش‌نویس
            </button>
            <button 
              type="submit"
              class="px-6 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
            >
              ایجاد ابزارک
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';
import { useAuth } from '~/composables/useAuth';

// تعریف definePageMeta، useHead و navigateTo برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
declare const useHead: (head: { title?: string }) => void
declare const navigateTo: (to: string) => Promise<void>

const router = useRouter()

// Page meta definition
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user, hasPermission } = useAuth()

// Page title
useHead({
  title: 'ایجاد باکس دسته‌بندی - پنل ادمین'
})

// Form state
const form = ref({
  title: '',
  code: '',
  page: '',
  order: 1,
  columns: 3,
  maxCategories: 12,
  itemHeight: 120,
  showImages: true,
  showCounts: true,
  showDescriptions: false,
  categorySelectionType: 'manual',
  categories: [] as Array<{
    name: string
    image: string
    link: string
    description: string
    searchTerm: string
    showDropdown: boolean
    selectedCategory: any
  }>
})

// All categories for search
interface Category {
  id: number
  name: string
  [key: string]: unknown
}
const allCategories = ref<Category[]>([])

// Computed filtered categories
const filteredCategories = computed(() => {
  const activeCategory = form.value.categories.find(cat => cat.showDropdown)
  if (!activeCategory || !activeCategory.searchTerm.trim()) {
    return allCategories.value.slice(0, 10) // Show first 10 when no search
  }
  
  const term = activeCategory.searchTerm.toLowerCase()
  return allCategories.value.filter(category => {
    const parentName = category.parent_name as string | undefined
    return category.name.toLowerCase().includes(term) ||
      (parentName && typeof parentName === 'string' && parentName.toLowerCase().includes(term))
  })
})

// Methods
const addCategoryItem = () => {
  form.value.categories.push({
    name: '',
    image: '',
    link: '',
    description: '',
    searchTerm: '',
    showDropdown: false,
    selectedCategory: null
  })
}

const removeCategoryItem = (index: number) => {
  form.value.categories.splice(index, 1)
}

// Select category for specific item
const selectCategoryForItem = (index: number, category: any) => {
  const item = form.value.categories[index]
  item.selectedCategory = category
  item.name = category.name
  item.link = category.slug || category.name
  item.searchTerm = category.name
  item.showDropdown = false
  console.log('دسته‌بندی انتخاب شد:', category)
}

// Hide category dropdown with delay
const hideCategoryDropdown = (index: number) => {
  setTimeout(() => {
    form.value.categories[index].showDropdown = false
  }, 200)
}

// Load categories on mount
onMounted(async () => {
  try {
    const response = await $fetch('/api/admin/product-categories?all=1')
    let raw = []
    if (Array.isArray(response)) {
      raw = response
    } else if (Array.isArray((response as any)?.data)) {
      raw = (response as any).data
    } else {
      raw = []
    }
    
    // Add parent_name for display
    raw.forEach((cat: any) => {
      if (cat.parent_id) {
        const parent = raw.find((c: any) => c.id === cat.parent_id)
        cat.parent_name = parent ? parent.name : '-'
      } else {
        cat.parent_name = '-'
      }
    })
    
    allCategories.value = raw
  } catch (error) {
    console.error('خطا در دریافت دسته‌بندی‌ها:', error)
    allCategories.value = []
  }
})

const handleSubmit = async () => {
  try {
    // Validate form
    if (!form.value.title || !form.value.code || !form.value.page) {
      alert('لطفاً تمام فیلدهای اجباری را پر کنید')
      return
    }

    if (form.value.categorySelectionType === 'manual' && form.value.categories.length === 0) {
      alert('لطفاً حداقل یک دسته‌بندی انتخاب کنید')
      return
    }

    // Here you would call your API to create the widget
    // Creating categories widget
    
    // Redirect to management page
    await router.push('/admin/content/banners/Categories')
  } catch (error) {
    // Error creating widget
    alert('خطا در ایجاد ابزارک')
  }
}

const saveAsDraft = async () => {
  try {
    // Save as draft logic
    // Saving as draft
    alert('پیش‌نویس ذخیره شد')
  } catch (error) {
    // Error saving draft
    alert('خطا در ذخیره پیش‌نویس')
  }
}
</script>

<style scoped>
.categories-create-page {
  background-color: #f9fafb;
  min-height: 100vh;
}
</style>


