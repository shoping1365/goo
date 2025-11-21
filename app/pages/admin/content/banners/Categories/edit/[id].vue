<template>
  <div class="categories-edit-page p-6">
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
          <h1 class="text-3xl font-bold text-gray-900">ویرایش باکس دسته‌بندی</h1>
        </div>
        <p class="text-gray-600">ویرایش ابزارک باکس دسته‌بندی</p>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="flex justify-center items-center py-12">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
      </div>

      <!-- Edit Form -->
      <div v-else-if="categoryBox" class="bg-white rounded-lg shadow-lg p-6">
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
                  placeholder="مثال: باکس دسته‌بندی‌های محبوب"
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
                  <option value="home">صفحه اصلی</option>
                  <option value="other">سایر بخش‌ها</option>
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
                <label class="block text-sm font-medium text-gray-700 mb-2">تعداد ستون‌ها</label>
                <select 
                  v-model="form.columns"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                >
                  <option value="2">2 ستون</option>
                  <option value="3">3 ستون</option>
                  <option value="4">4 ستون</option>
                  <option value="5">5 ستون</option>
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
                  min="80"
                  max="300"
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
                <option value="parent">دسته‌بندی‌های اصلی</option>
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
                    <input 
                      v-model="category.link"
                      type="text"
                      placeholder="لینک دسته‌بندی"
                      class="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                    />
                    <input 
                      v-model="category.description"
                      type="text"
                      placeholder="توضیحات"
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
              @click="resetForm"
            >
              بازنشانی تغییرات
            </button>
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
              به‌روزرسانی ابزارک
            </button>
          </div>
        </form>
      </div>

      <!-- Error State -->
      <div v-else class="text-center py-12 bg-white rounded-lg shadow-sm">
        <div class="w-24 h-24 mx-auto mb-4 bg-red-100 rounded-full flex items-center justify-center">
          <svg class="w-12 h-12 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"/>
          </svg>
        </div>
        <h3 class="text-lg font-semibold text-gray-800 mb-2">خطا در بارگذاری</h3>
        <p class="text-gray-600 mb-4">باکس دسته‌بندی مورد نظر یافت نشد</p>
        <NuxtLink 
          to="/admin/content/banners/Categories"
          class="bg-blue-600 hover:bg-blue-700 text-white px-6 py-3 rounded-lg font-semibold transition-colors duration-200 inline-flex items-center gap-2"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"/>
          </svg>
          بازگشت به لیست
        </NuxtLink>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// Page meta definition
definePageMeta({
  layout: 'admin',
  middleware: 'admin'
})

import { onMounted, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';

// تعریف definePageMeta، useHead و navigateTo برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
declare const useHead: (head: { title?: string }) => void
// navigateTo is not used in this component

// Page title
useHead({
  title: 'ویرایش باکس دسته‌بندی - پنل ادمین'
})

// Route
const route = useRoute()
const router = useRouter()
const categoryBoxId = route.params.id as string

// State
const loading = ref(true)
interface CategoryBoxData {
  id: string
  title: string
  code: string
  page: string
  order: number
  columns: string
  maxCategories: number
  itemHeight: number
  showImages: boolean
  showCounts: boolean
  showDescriptions: boolean
  categorySelectionType: string
  categories: Array<{
    name: string
    image: string
    link: string
    description: string
  }>
}

const categoryBox = ref<CategoryBoxData | null>(null)
const form = ref({
  title: '',
  code: '',
  page: '',
  order: 1,
  columns: '4',
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
  }>
})

// Mock data for demonstration
const mockCategoryBox = {
  id: categoryBoxId,
  title: 'باکس دسته‌بندی‌های محبوب',
  code: 'popular-categories',
  page: 'home',
  order: 2,
  columns: '4',
  maxCategories: 12,
  itemHeight: 120,
  showImages: true,
  showCounts: true,
  showDescriptions: false,
  categorySelectionType: 'manual',
  categories: [
    { name: 'الکترونیک', image: '/images/cat1.png', link: '/category/electronics', description: 'محصولات الکترونیکی' },
    { name: 'پوشاک', image: '/images/cat2.png', link: '/category/clothing', description: 'لباس و کفش' }
  ]
}

// Methods
const loadCategoryBox = async () => {
  try {
    loading.value = true
    // Here you would call your API to load the category box
    // For now, using mock data
    await new Promise(resolve => setTimeout(resolve, 1000))
    categoryBox.value = mockCategoryBox
    populateForm()
  } catch (error) {
    console.error('Error loading category box:', error)
  } finally {
    loading.value = false
  }
}

const populateForm = () => {
  if (!categoryBox.value) return
  
  const boxData = categoryBox.value
  form.value = {
    title: boxData.title,
    code: boxData.code,
    page: boxData.page,
    order: boxData.order,
    columns: boxData.columns,
    maxCategories: boxData.maxCategories,
    itemHeight: boxData.itemHeight,
    showImages: boxData.showImages,
    showCounts: boxData.showCounts,
    showDescriptions: boxData.showDescriptions,
    categorySelectionType: boxData.categorySelectionType,
    categories: [...boxData.categories]
  }
}

const addCategoryItem = () => {
  form.value.categories.push({
    name: '',
    image: '',
    link: '',
    description: ''
  })
}

const removeCategoryItem = (index: number) => {
  form.value.categories.splice(index, 1)
}

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

    // Here you would call your API to update the category box
    
    // Redirect to management page
    await router.push('/admin/content/banners/Categories')
  } catch (error) {
    console.error('Error updating widget:', error)
    alert('خطا در به‌روزرسانی ابزارک')
  }
}

const saveAsDraft = async () => {
  try {
    // Save as draft logic
    alert('پیش‌نویس ذخیره شد')
  } catch (error) {
    console.error('Error saving draft:', error)
    alert('خطا در ذخیره پیش‌نویس')
  }
}

const resetForm = () => {
  populateForm()
}

// Lifecycle
onMounted(() => {
  loadCategoryBox()
})
</script>

<style scoped>
.categories-edit-page {
  background-color: #f9fafb;
  min-height: 100vh;
}
</style>

