<template>
  <div class="category-box-create-page min-h-screen">
    <!-- هدر صفحه -->
    <div class="bg-white shadow-sm border-2 border-blue-200 rounded-lg mx-4 mt-4">
      <div class="w-full px-4 sm:px-6 lg:px-8">
        <div class="flex items-center justify-between h-20 text-right">
          <div class="flex items-center">
            <h1 class="text-xl font-semibold text-gray-900">ویرایش ویجت دسته‌بندی</h1>
          </div>
          <div class="flex items-center gap-8">
            <TemplateButton
              :disabled="isSaving"
              bg-gradient="bg-green-600"
              text-color="text-white"
              border-color="border border-transparent"
              hover-class="hover:bg-green-700"
              focus-class="focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
              size="medium"
              custom-class="flex items-center gap-2"
              @click="saveWidget"
            >
              <i class="fas fa-check"></i>
              {{ isSaving ? 'در حال ذخیره...' : 'ذخیره' }}
            </TemplateButton>
            <NuxtLink to="/admin/content/banners">
              <TemplateButton
                bg-gradient="bg-gradient-to-b from-purple-400 to-purple-600"
                text-color="text-white"
                border-color="border border-transparent"
                hover-class="hover:from-purple-500 hover:to-purple-700"
                focus-class="focus:ring-2 focus:ring-offset-2 focus:ring-purple-500"
                size="medium"
                custom-class="flex items-center gap-2"
              >
                <i class="fas fa-arrow-right"></i>
                بازگشت به ابزارک‌ها
              </TemplateButton>
            </NuxtLink>
          </div>
        </div>
      </div>
    </div>

    <div class="w-full px-4 sm:px-6 lg:px-8 py-8">
      <!-- تنظیمات ویجت -->
      <CategoryBoxSettings :form-data="formData" :config="config" />

      <!-- مدیریت دسته‌بندی‌ها -->
      <div class="bg-white rounded-lg shadow mb-6">
        <div class="px-6 py-4 border-b border-gray-200">
          <div class="flex items-center justify-between">
            <h2 class="text-lg font-medium text-gray-900">مدیریت دسته‌بندی‌ها</h2>
            <button
              class="px-4 py-2 text-sm font-medium text-white bg-purple-600 border border-transparent rounded-md hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-purple-500"
              @click="addCategory"
            >
              افزودن دسته‌بندی
            </button>
          </div>
        </div>
        <div class="p-6">
          <div v-if="config.categories.length === 0" class="text-center py-8 text-gray-500">
            <i class="fas fa-tags text-4xl mb-4"></i>
            <p>هنوز دسته‌بندی‌ای اضافه نشده است</p>
            <p class="text-sm">برای شروع روی دکمه "افزودن دسته‌بندی" کلیک کنید</p>
          </div>
          <div v-else class="space-y-4">
            <div
              v-for="(category, index) in config.categories"
              :key="index"
              class="border border-gray-200 rounded-lg p-6"
            >
              <!-- تست: دکمه حذف در بالای هر دسته‌بندی -->
              <div class="flex justify-end mb-2">
                <button
                  class="px-3 py-1 text-sm text-red-600 hover:text-red-800 hover:bg-red-50 rounded-md border border-red-200"
                  @click="removeCategory(index)"
                >
                  حذف
                </button>
              </div>
              
              <div class="flex items-start space-x-4">
                <!-- تصویر دسته‌بندی -->
                <div class="flex-shrink-0">
                  <div class="w-16 h-16 bg-gray-100 rounded-lg overflow-hidden">
                    <img
                      v-if="category.image"
                      :src="category.image"
                      :alt="category.title"
                      class="w-full h-full object-cover"
                    />
                    <div v-else class="w-full h-full flex items-center justify-center text-gray-400">
                      <i class="fas fa-image"></i>
                    </div>
                  </div>
                  <button
                    type="button"
                    class="mt-2 px-3 py-1.5 text-xs font-semibold text-white bg-emerald-500 border border-transparent rounded-md hover:bg-emerald-600 focus:outline-none focus:ring-2 focus:ring-emerald-500 focus:ring-offset-2 shadow-md transition-all duration-200 flex items-center gap-1"
                    @click="openMediaModal(index)"
                  >
                    <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
                    </svg>
                    انتخاب تصویر
                  </button>
                </div>
                
                <!-- اطلاعات دسته‌بندی -->
                <div class="flex-1 min-w-0">
                  <div class="space-y-4">
                    <!-- عنوان و دسته‌بندی در یک ردیف -->
                    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                      <div>
                        <label class="block text-sm font-medium text-gray-700 mb-1">عنوان</label>
                        <input
                          v-model="category.title"
                          type="text"
                          placeholder="عنوان دسته‌بندی"
                          class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500"
                        />
                      </div>
                      <div>
                        <label class="block text-sm font-medium text-gray-700 mb-1">دسته‌بندی</label>
                        <div class="relative">
                          <input 
                            v-model="category.searchTerm"
                            :data-category-index="index"
                            type="text"
                            placeholder="جستجو در دسته‌بندی‌ها..."
                            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500"
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
                                  <span v-if="cat.product_count > 0" class="text-xs text-gray-500">
                                    {{ cat.product_count }} محصول
                                  </span>
                                </div>
                              </div>
                            </div>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- پیش‌نمایش -->
      <div class="bg-white rounded-lg shadow">
        <div class="px-6 py-4 border-b border-gray-200">
          <h2 class="text-lg font-medium text-gray-900">پیش‌نمایش</h2>
        </div>
        <div class="p-6">
          <CategoryBoxPreview :config="config" />
        </div>
      </div>
    </div>

    <!-- مودال انتخاب رسانه -->
    <MediaLibraryModal
      v-model="showMediaModal"
      :default-category="'banners'"
      :file-type="'image'"
      :max-select="1"
      @confirm="onMediaSelected"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import type { CategoryConfig, CategoryItem } from '~/types/widget'
import { useWidgetRegistry } from '~/composables/useWidgetRegistry'
import { useWidget } from '~/composables/useWidget'
import CategoryBoxPreview from '~/components/admin/CategoryBoxPreview.vue'
import CategoryBoxSettings from './components/CategoryBoxSettings.vue'
import MediaLibraryModal from '~/components/media/MediaLibraryModal.vue'

import { useRoute, useRouter } from 'vue-router'

// تعریف definePageMeta و navigateTo برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
declare const navigateTo: (to: string) => Promise<void>

const router = useRouter()

 definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// استفاده از composables
const { createWidget } = useWidget()

// Route parameters
const route = useRoute()
const widgetId = route.params.id as string

// State
const isSaving = ref(false)
const showMediaModal = ref(false)
const selectedCategoryIndex = ref<number | null>(null)
const isLoading = ref(true)

// All categories for search
const allCategories = ref<any[]>([])

// Computed filtered categories
const filteredCategories = computed(() => {
  const activeCategory = config.value.categories.find(cat => cat.showDropdown)
  if (!activeCategory || !activeCategory.searchTerm.trim()) {
    return allCategories.value.slice(0, 10) // Show first 10 when no search
  }
  
  const term = activeCategory.searchTerm.toLowerCase()
  return allCategories.value.filter(category => 
    category.name.toLowerCase().includes(term) ||
    (category.parent_name && category.parent_name.toLowerCase().includes(term))
  )
})

// فرم داده‌ها
const formData = ref({
  title: '',
  type: 'category-box' as const,
  status: 'active' as const,
  page: 'home' as const
})

// تنظیمات ویجت
const config = ref<CategoryConfig>({
  categories: [],
  layout: 'grid',
  vertical_layout: false,
  show_product_count: true,
  columns: 3,
  mobile_columns: 2,
  box_width: 'center',
  show_title: true,
  title: 'دسته‌بندی‌های محبوب',
  title_color: '#1f2937',
  card_style: 'rounded',
  image_size: 'medium',
  border_radius: 'medium',
  text_alignment: 'center',
  background_color: '#ffffff',
  show_background: false,
  full_width_background: false,
  show_border: false,
  border_color: '#e5e7eb',
  border_width: 'medium',
  category_source: 'manual',
  padding: '24px',
  margin: '0px'
})

// Methods
const addCategory = () => {
  // بر اساس منبع انتخاب شده، دسته‌بندی‌ها را اضافه کن
  switch (config.value.category_source) {
    case 'best_selling':
      // پرفروش‌ترین دسته‌بندی‌ها
      config.value.categories.push({
        id: Date.now(),
        title: 'پرفروش‌ترین دسته‌بندی‌ها',
        description: 'دسته‌بندی‌های پرفروش بر اساس فروش',
        image: '/default-product.svg',
        link: '/categories/best-selling',
        product_count: 0,
        searchTerm: '',
        showDropdown: false,
        selectedCategory: null
      } as any)
      break
    case 'user_search':
      // بر اساس سرچ کاربر
      config.value.categories.push({
        id: Date.now(),
        title: 'دسته‌بندی‌های جستجو شده',
        description: 'دسته‌بندی‌های محبوب در جستجوها',
        image: '/default-product.svg',
        link: '/categories/searched',
        product_count: 0,
        searchTerm: '',
        showDropdown: false,
        selectedCategory: null
      } as any)
      break
    case 'user_interests':
      // بر اساس علایق کاربر
      config.value.categories.push({
        id: Date.now(),
        title: 'دسته‌بندی‌های مورد علاقه',
        description: 'دسته‌بندی‌های بر اساس علایق کاربران',
        image: '/default-product.svg',
        link: '/categories/interests',
        product_count: 0,
        searchTerm: '',
        showDropdown: false,
        selectedCategory: null
      } as any)
      break
    case 'manual':
    default:
      // انتخاب دستی
      config.value.categories.push({
        id: Date.now(),
        title: '',
        description: '',
        image: '',
        link: '',
        product_count: 0,
        searchTerm: '',
        showDropdown: false,
        selectedCategory: null
      } as any)
      break
  }
}

// Select category for specific item
const selectCategoryForItem = (index: number, category: any) => {
  const item = config.value.categories[index]
  item.selectedCategory = category
  item.title = category.name
  item.link = category.slug || category.name
  item.searchTerm = category.name
  item.showDropdown = false
  console.log('دسته‌بندی انتخاب شد:', category)
}

// Hide category dropdown with delay
const hideCategoryDropdown = (index: number) => {
  setTimeout(() => {
    config.value.categories[index].showDropdown = false
  }, 200)
}

// Load widget data
const loadWidget = async () => {
  try {
    console.log('در حال لود ویجت با ID:', widgetId)
    const response = await $fetch(`/api/admin/widgets/${widgetId}`) as any
    console.log('پاسخ API:', response)
    
    if (response && response.data) {
      const widget = response.data
      console.log('داده‌های ویجت:', widget)
      
      // Load form data
      formData.value.title = widget.title || ''
      formData.value.page = widget.page || 'home'
      
      // Load config
      config.value = {
        ...config.value,
        ...widget.config
      }
      
      console.log('config لود شده:', config.value)
    } else if (response && response.config) {
      // Load form data
      formData.value.title = response.title || ''
      formData.value.page = response.page || 'home'
      
      // Load config
      config.value = {
        ...config.value,
        ...response.config
      }
      
      console.log('config لود شده (fallback):', config.value)
    } else {
      console.log('هیچ داده‌ای دریافت نشد')
    }
  } catch (error) {
    console.error('خطا در دریافت ویجت:', error)
    alert('خطا در دریافت ویجت. لطفاً دوباره تلاش کنید.')
  }
}

// Load categories on mount
onMounted(async () => {
  try {
    // Load widget data first
    await loadWidget()
    
    // Then load categories
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
  } finally {
    isLoading.value = false
  }
})

// Media modal functions
const openMediaModal = (index: number) => {
  selectedCategoryIndex.value = index
  showMediaModal.value = true
}

const closeMediaModal = () => {
  showMediaModal.value = false
  selectedCategoryIndex.value = null
}

const onMediaSelected = (selectedMedia: any[]) => {
  if (selectedCategoryIndex.value !== null && selectedMedia.length > 0) {
    const media = selectedMedia[0]
    config.value.categories[selectedCategoryIndex.value].image = media.url
    closeMediaModal()
  }
}

const removeCategory = (index: number) => {
  config.value.categories.splice(index, 1)
}

const editCategory = (index: number) => {
  const category = config.value.categories[index]
  
  // باز کردن dropdown برای ویرایش
  category.showDropdown = true
  
  // پاک کردن search term برای نمایش همه گزینه‌ها
  category.searchTerm = ''
  
  // اسکرول به input field
  setTimeout(() => {
    const inputElement = document.querySelector(`input[data-category-index="${index}"]`) as HTMLInputElement
    if (inputElement) {
      inputElement.focus()
    }
  }, 100)
}

const saveWidget = async () => {
  try {
    isSaving.value = true

    // اعتبارسنجی
    if (!formData.value.title) {
      alert('لطفاً عنوان ویجت را وارد کنید')
      return
    }

    if (config.value.categories.length === 0) {
      alert('لطفاً حداقل یک دسته‌بندی اضافه کنید')
      return
    }

    // آماده‌سازی داده‌ها برای آپدیت
    const widgetData = {
      title: formData.value.title,
      type: formData.value.type,
      status: formData.value.status,
      page: formData.value.page,
      config: config.value
    }

    // آپدیت ویجت
    const response = await $fetch(`/api/admin/widgets/${widgetId}`, {
      method: 'PUT',
      body: widgetData
    })

    if (response) {
      // هدایت به صفحه لیست ویجت‌ها
      await router.push('/admin/content/banners')
    }

  } catch (error) {
    console.error('خطا در ذخیره ویجت:', error)
    alert('خطا در ذخیره ویجت. لطفاً دوباره تلاش کنید.')
  } finally {
    isSaving.value = false
  }
}

const saveAsDraft = async () => {
  try {
    isSaving.value = true

    // ذخیره به عنوان پیش‌نویس
    const widgetData = {
      title: formData.value.title || 'پیش‌نویس ویجت دسته‌بندی',
      type: formData.value.type,
      status: 'draft' as const,
      page: formData.value.page,
      config: config.value
    }

    await createWidget(widgetData)
    alert('پیش‌نویس ذخیره شد')

  } catch (error) {
    console.error('خطا در ذخیره پیش‌نویس:', error)
    alert('خطا در ذخیره پیش‌نویس')
  } finally {
    isSaving.value = false
  }
}
</script>

<style scoped>
.category-box-create-page {
  /* پس‌زمینه حذف شد */
}
</style>
