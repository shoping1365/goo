<template>
  <div class="min-h-screen">
    <!-- Header -->
    <div class="bg-white shadow-sm border-b border-gray-200 sticky top-0 z-30">
      <div class="px-6 py-4">
        <div class="flex items-center justify-between">
          <div>
            <h1 class="text-2xl font-bold text-gray-900">ویرایش نوشته</h1>
            <p class="text-sm text-gray-500 mt-1">ویرایش و بروزرسانی محتوای موجود</p>
          </div>
          
          <div class="flex items-center gap-3">
            <button 
              @click="previewPost"
              :disabled="isSaving"
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-purple-400 to-purple-600 hover:from-purple-500 hover:to-purple-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-purple-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105 disabled:opacity-50"
            >
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
              </svg>
              پیش‌نمایش
            </button>
            
            <button 
              @click="saveAndContinueEditing"
              :disabled="isSaving"
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-blue-400 to-blue-600 hover:from-blue-500 hover:to-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105 disabled:opacity-50"
            >
              <svg v-if="isSaving" class="w-5 h-5 ml-2 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
              </svg>
              <svg v-else class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
              </svg>
              {{ isSaving ? 'در حال ذخیره...' : 'ذخیره و ادامه ویرایش' }}
            </button>
            
            <button 
              @click="saveAndGoToList"
              :disabled="isSaving"
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-green-400 to-green-600 hover:from-green-500 hover:to-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105 disabled:opacity-50"
            >
              <svg v-if="isSaving" class="w-5 h-5 ml-2 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
              </svg>
              <svg v-else class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
              </svg>
              {{ isSaving ? 'در حال ذخیره...' : 'ذخیره' }}
            </button>
            
            <NuxtLink 
              to="/admin/post-management/all-posts"
              class="inline-flex items-center px-4 py-2 border border-gray-200 rounded-lg bg-white hover:bg-gray-50 transition-all shadow-md"
            >
              <svg class="w-5 h-5 ml-2 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
              </svg>
              <span class="text-gray-700">بازگشت</span>
            </NuxtLink>
          </div>
        </div>
      </div>
    </div>

    <!-- Main Content -->
    <div class="px-4 py-4 w-full">
      <div class="grid grid-cols-1 lg:grid-cols-5 gapx-4 py-4 w-full items-start">
        <!-- فرم اصلی -->
        <div class="lg:col-span-4 space-y-6 w-full">
          <!-- عنوان -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 px-4 py-4 w-full">
            <label class="block text-sm font-medium text-gray-700 mb-2">عنوان نوشته</label>
            <input 
              v-model="postForm.title"
              type="text"
              placeholder="عنوان نوشته را وارد کنید..."
              class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors text-right"
              dir="rtl"
            >
          </div>

          <!-- خلاصه نوشته -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 px-4 py-4 w-full mt-0">
            <div class="flex items-center justify-between mb-2">
              <label class="block text-sm font-medium text-gray-700">خلاصه نوشته</label>
              <TitleInserter 
                :excerpt="postForm.excerpt"
                @update:excerpt="(newExcerpt) => postForm.excerpt = newExcerpt"
              />
            </div>
            <RichTextEditor 
              v-model="postForm.excerpt"
              :lang="'fa'"
              :direction="'rtl'"
              :height="420"
            />
          </div>


        </div>

        <!-- ساید بار -->
        <div class="space-y-4 w-full h-fit">
          <!-- وضعیت انتشار -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 px-4 py-4 w-full">
            <h3 class="text-base font-semibold text-gray-900 mb-2">وضعیت انتشار</h3>
            <div class="space-y-2">
              <label class="flex items-center">
                <input 
                  v-model="postForm.status"
                  type="radio"
                  value="draft"
                  class="w-4 h-4 text-blue-600 border-gray-300 focus:ring-blue-500"
                >
                <span class="mr-2 text-xs text-gray-700">پیش‌نویس</span>
              </label>
              <label class="flex items-center">
                <input 
                  v-model="postForm.status"
                  type="radio"
                  value="published"
                  class="w-4 h-4 text-blue-600 border-gray-300 focus:ring-blue-500"
                >
                <span class="mr-2 text-xs text-gray-700">منتشر شده</span>
              </label>
              <label class="flex items-center">
                <input 
                  v-model="postForm.status"
                  type="radio"
                  value="scheduled"
                  class="w-4 h-4 text-blue-600 border-gray-300 focus:ring-blue-500"
                >
                <span class="mr-2 text-xs text-gray-700">زمان‌بندی شده</span>
              </label>
            </div>
          </div>

          <!-- دسته‌بندی -->
          <div v-if="categories.length > 0" class="bg-white rounded-lg shadow-sm border border-gray-200 p-3 w-full">
            <h3 class="text-base font-semibold text-gray-900 mb-2">دسته‌بندی</h3>
            <select 
              v-model="postForm.category_id"
              class="w-full px-2 py-1 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-right text-sm"
            >
              <option value="">بدون دسته‌بندی</option>
              <option v-for="category in categories" :key="category.id" :value="category.id">
                {{ category.name }}
              </option>
            </select>
          </div>
          
          <!-- پیام عدم وجود دسته‌بندی -->
          <div v-else class="bg-yellow-50 rounded-lg border border-yellow-200 p-3 w-full">
            <div class="flex items-center">
              <svg class="w-5 h-5 text-yellow-400 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
              </svg>
              <div>
                <h3 class="text-sm font-medium text-yellow-800">هیچ دسته‌بندی‌ای وجود ندارد</h3>
                <p class="text-sm text-yellow-700 mt-1">
                  برای دسته‌بندی کردن نوشته‌ها، ابتدا 
                  <NuxtLink to="/admin/post-management/categories" class="underline hover:no-underline">
                    دسته‌بندی جدید ایجاد کنید
                  </NuxtLink>
                </p>
              </div>
            </div>
          </div>

          <!-- تصویر شاخص -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-3 w-full">
            <h3 class="text-base font-semibold text-gray-900 mb-2">تصویر شاخص</h3>
            <div class="space-y-2">
              <div class="relative w-full h-64 md:h-80">
                <template v-if="postForm.featured_image">
                  <img :src="postForm.featured_image" alt="تصویر شاخص" class="w-full h-full object-cover rounded-lg">
                  <button 
                    @click="removeFeaturedImage"
                    class="absolute top-2 right-2 bg-red-500 text-white p-1 rounded-full hover:bg-red-600 transition-colors"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                    </svg>
                  </button>
                </template>
                <template v-else>
                  <button 
                    @click="showMediaLibrary = true"
                    class="w-full h-full flex flex-col items-center justify-center border-2 border-dashed border-gray-300 rounded-lg hover:border-blue-500 transition-colors text-gray-600 hover:text-blue-600 text-sm"
                  >
                    <svg class="w-8 h-8 mb-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
                    </svg>
                    انتخاب تصویر شاخص
                  </button>
                </template>
              </div>
            </div>
          </div>

        </div>
      </div>

                      <!-- محتوای اصلی تمام عرض -->
                <div class="mt-6 w-full">
                  <div class="bg-white rounded-lg shadow-sm border border-gray-200 px-4 py-4 w-full">
                    <div class="flex items-center justify-between mb-2">
                      <label class="block text-sm font-medium text-gray-700">محتوای اصلی</label>
                    </div>
                    
                    <RichTextEditor
                      v-model="postForm.content"
                      :lang="'fa'"
                      :direction="'rtl'"
                      :height="500"
                      @h-tag-selected="handleHTagSelected"
                    />
                  </div>
                </div>

      <!-- Dropdown پیشنهاد عناوین -->
      <TitleSuggestionDropdown
        :show-dropdown="showTitleDropdown"
        :excerpt="postForm.excerpt"
        :position="dropdownPosition"
        @title-selected="handleTitleSelected"
        @close="showTitleDropdown = false"
      />

      <!-- تنظیمات SEO -->
      <div class="mt-6 w-full">
        <SeoTab 
          v-model:meta-title="postForm.meta_title"
          v-model:meta-description="postForm.meta_description"
          :slug="postForm.slug"
          :enable-og-type="true"
          :og-title="postForm.og_title"
          :og-description="postForm.og_description"
          :og-type="postForm.og_type"
          :og-site-name="postForm.og_site_name"
          :og-image="postForm.og_image"
          :robots-meta="postForm.robots_meta"
          :post-form="postForm"
          @update:slug="updateSlug"
          @update:og-title="(val) => { postForm.og_title = val }"
          @update:og-description="(val) => { postForm.og_description = val }"
          @update:og-type="(val) => { postForm.og_type = val }"
          @update:og-site-name="(val) => { postForm.og_site_name = val }"
          @update:og-image="updateOgImage"
          @update:robots-meta="(val) => { postForm.robots_meta = val }"
          @select-og-image="showOgImageModal = true"
        />
      </div>



    </div>

    <!-- Media Library Modal برای تصویر شاخص -->
    <MediaLibraryModal 
      v-model="showMediaLibrary"
      :file-type="'image'"
      :default-category="'library'"
      @confirm="onImageSelected"
    />
    <MediaLibraryModal 
      v-model="showOgImageModal"
      :file-type="'image'"
      :default-category="'library'"
      @confirm="onOgImageSelected"
    />
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
declare const useHead: (head: { title?: string }) => void
declare const navigateTo: (to: string) => Promise<void>
</script>

<script setup lang="ts">
import { computed, onMounted, reactive, ref, watch } from 'vue'
import { useRoute } from 'vue-router'
import RichTextEditor from '~/components/common/RichTextEditor.vue'
import MediaLibraryModal from '~/components/media/MediaLibraryModal.vue'
import TitleInserter from '~/pages/admin/post-management/TitleInserter.vue'
import TitleSuggestionDropdown from '~/pages/admin/post-management/TitleSuggestionDropdown.vue'
import SeoTab from '../add-post/SeoTab.vue'

import { useSlugManagement } from '~/composables/useSlugManagement'
// تعریف interface برای نوشته
interface Post {
  id?: number
  title: string
  excerpt: string
  content: string
  status: 'draft' | 'published' | 'scheduled'
  category_id: number | null
  featured_image: string
  og_image: string // تصویر Open Graph
  og_title?: string // عنوان Open Graph
  og_description?: string // توضیحات Open Graph
  og_type?: string // نوع محتوا Open Graph
  og_site_name?: string // نام سایت Open Graph
  slug: string
  published_at?: string
  meta_title?: string

  meta_description?: string
  meta_keywords?: string
  
  // فیلدهای SEO جدید
  robots_meta?: string // robots meta tag content
  index_status?: string
  follow_status?: string
  
  // فیلدهای جدید محاسباتی
  word_count?: number
  reading_time?: number
  
  // فیلدهای مخصوص مقالات آموزشی
  difficulty?: string
  estimated_time?: number
  prerequisites?: string
  
  // فیلدهای مخصوص مقالات خبری
  news_category?: string
  location?: string
  event_date?: string
  
  // فیلدهای مخصوص مقالات بررسی محصول
  product_price?: number
  rating?: number
  pros?: string
  cons?: string
  
  // فیلدهای عمومی
  tags?: string[]
  language?: string
  is_accessible_for_free?: boolean
}

// تعریف interface برای دسته‌بندی
interface Category {
  id: number
  name: string
  slug: string
}

definePageMeta({
  layout: 'admin-main'
})

useHead({
  title: 'ویرایش نوشته - پنل مدیریت'
})

// Route params
const route = useRoute()
const postId = route.params.id as string

// Reactive data
const isSaving = ref(false)
const categories = ref<Category[]>([])
const showMediaLibrary = ref(false)
const showOgImageModal = ref(false) // Modal برای انتخاب تصویر Open Graph
const showTitleDropdown = ref(false) // برای نمایش dropdown پیشنهاد عناوین
const dropdownPosition = ref({ x: 0, y: 0 }) // موقعیت dropdown


// متغیرهای جدید
const tagsInput = ref('')
const helpField = ref('')

// استفاده از composable
const { slugify } = useSlugManagement()

// فرم نوشته
const postForm = reactive<Post>({
  id: undefined, // شناسه نوشته برای مدیریت عناوین
  title: '',
  excerpt: '',
  content: '',
  status: 'draft',
  category_id: null,
  featured_image: '',
  og_image: '', // تصویر Open Graph
  og_title: '', // عنوان Open Graph
  og_description: '', // توضیحات Open Graph
  og_type: 'article', // نوع محتوا Open Graph
  og_site_name: '', // نام سایت Open Graph
  slug: '',
  meta_title: '',

  meta_description: '',
  
  // فیلدهای SEO جدید
  robots_meta: 'index,follow',
  
  // فیلدهای جدید محاسباتی
  word_count: 0,

  
  // فیلدهای مخصوص مقالات آموزشی
  difficulty: 'medium',
  estimated_time: 0,
  prerequisites: '',
  
  // فیلدهای مخصوص مقالات خبری
  news_category: '',
  location: '',
  event_date: '',
  
  // فیلدهای مخصوص مقالات بررسی محصول
  product_price: 0,
  rating: 0,
  pros: '',
  cons: '',
  
  // فیلدهای عمومی
  tags: [],
  language: 'fa',
  is_accessible_for_free: true,
  meta_keywords: '',
  index_status: 'index',
  follow_status: 'follow',
  reading_time: 0,
})

// متغیرهای واکنشی برای تشخیص تغییر دستی هر فیلد
const seoTitleTouched = ref(false)
const slugTouched = ref(false)
const ogImageTouched = ref(false)
const ogTypeTouched = ref(false)
const metaDescriptionTouched = ref(false)

// همگام‌سازی واکنشی og_image با featured_image فقط در همین صفحه
watch(() => postForm.featured_image, (newFeatured) => {
  if (!ogImageTouched.value) {
    postForm.og_image = newFeatured
  }
})

// watch برای تنظیم touched وقتی meta_description تغییر می‌کند
watch(() => postForm.meta_description, () => {
  metaDescriptionTouched.value = true
})

// watch برای تنظیم touched وقتی meta_description تغییر می‌کند
watch(() => postForm.meta_description, () => {
  metaDescriptionTouched.value = true
})

// watch برای تنظیم touched وقتی og_type تغییر می‌کند (فقط وقتی کاربر دستی تغییر دهد)
watch(() => postForm.og_type, (newVal, oldVal) => {
  // اگر تغییر از طریق sync نبوده، touched را تنظیم کن
  if (oldVal !== undefined && newVal !== oldVal) {
    ogTypeTouched.value = true
  }
})




// watch برای تنظیم touched وقتی meta_title تغییر می‌کند
watch(() => postForm.meta_title, () => {
  seoTitleTouched.value = true
})



// watch برای تنظیم touched وقتی slug تغییر می‌کند
watch(() => postForm.slug, () => {
  slugTouched.value = true
})

// watch برای تنظیم touched وقتی og_image تغییر می‌کند (فقط وقتی کاربر دستی تغییر دهد)
watch(() => postForm.og_image, (newVal, oldVal) => {
  // اگر تغییر از طریق sync نبوده، touched را تنظیم کن
  if (oldVal !== undefined && newVal !== oldVal) {
    ogImageTouched.value = true
  }
})

// watch برای محاسبه خودکار تعداد کلمات و زمان مطالعه
watch([() => postForm.content, () => postForm.excerpt], () => {
  calculateWordCountAndReadingTime()
})

// تابع ریست کردن فیلدهای touched
const resetTouchedFields = () => {
  seoTitleTouched.value = false
  slugTouched.value = false


  ogImageTouched.value = false

  ogTypeTouched.value = false
  metaDescriptionTouched.value = false
}

// همگام‌سازی خودکار فیلدها با عنوان نوشته
watch(() => postForm.title, (newTitle, oldTitle) => {
  // اگر عنوان خالی شد، فیلدهای touched را ریست کن و فیلدها را خالی کن
  if (!newTitle || newTitle.trim() === '') {
    resetTouchedFields()
    postForm.meta_title = ''
    postForm.slug = ''


    postForm.meta_description = ''

    return
  }
  
  // اگر عنوان از خالی به غیر خالی تغییر کرد، فیلدهای touched را ریست کن
  if ((!oldTitle || oldTitle.trim() === '') && newTitle && newTitle.trim() !== '') {
    resetTouchedFields()
  }
  
  // همگام‌سازی فیلدها - SEO Title همیشه از عنوان اصلی گرفته می‌شود
  postForm.meta_title = String(newTitle)
  if (!slugTouched.value) postForm.slug = slugify(String(newTitle))
  

  // meta_description را هم با عنوان همگام کن اگر خالی باشد
  if (!postForm.meta_description && !metaDescriptionTouched.value) {
    postForm.meta_description = String(newTitle)
  }
  
})



// همگام‌سازی og_image با featured_image
watch(() => postForm.featured_image, (newImage) => {
  if (newImage && !ogImageTouched.value) {
    postForm.og_image = String(newImage)
  }
})

// Methods
const fetchCategories = async () => {
  try {
    interface ApiResponse {
      data?: Category[]
    }
    const response = await $fetch<ApiResponse>('/api/post-categories')
    if (response.data) {
      categories.value = response.data
    }
  } catch (error) {
    console.error('خطا در دریافت دسته‌بندی‌ها:', error)
    // در صورت خطا، آرایه خالی قرار بده
    categories.value = []
  }
}

const fetchPost = async () => {
  try {
    interface PostResponse {
      id?: number
      title?: string
      excerpt?: string
      content?: string
      status?: string
      category_id?: number | null
      featured_image?: string
      slug?: string
      meta_title?: string
      meta_description?: string
      og_title?: string
      og_description?: string
      og_image?: string
      og_type?: string
      og_site_name?: string
      robots_meta?: string
      tags?: string[]
    }
    const post = await $fetch<PostResponse>(`/api/posts/${postId}`)
    
    // پر کردن فرم با داده‌های موجود
    if (post.id) postForm.id = post.id
    postForm.title = post.title || ''
    postForm.excerpt = post.excerpt || ''
    postForm.content = post.content || ''
    postForm.status = (post.status as 'draft' | 'published' | 'scheduled') || 'draft'
    postForm.category_id = post.category_id || null
    postForm.featured_image = post.featured_image || ''
    postForm.slug = post.slug || ''
    postForm.meta_title = post.meta_title || ''

    postForm.meta_description = post.meta_description || ''
    postForm.og_title = post.og_title || ''
    postForm.og_description = post.og_description || ''
    postForm.og_image = post.og_image || ''
    postForm.og_type = post.og_type || 'article'
    postForm.og_site_name = post.og_site_name || ''
    postForm.robots_meta = post.robots_meta || 'index,follow'
    postForm.tags = post.tags || []
    
    // محاسبه تعداد کلمات و زمان مطالعه
    calculateWordCountAndReadingTime()
  } catch (error) {
    console.error('خطا در دریافت نوشته:', error)
    alert('خطا در دریافت نوشته')
    navigateTo('/admin/post-management/all-posts')
  }
}

const selectFeaturedImage = () => {
  showMediaLibrary.value = true
}

const onImageSelected = (selectedFiles: any[]) => {
  if (selectedFiles.length > 0) {
    postForm.featured_image = selectedFiles[0].url
  }
  showMediaLibrary.value = false
}

const removeFeaturedImage = () => {
  // اگر og_image sync است، آن را هم پاک کن
  if (!ogImageTouched.value) {
    postForm.og_image = ''
  }
  postForm.featured_image = ''
}

// انتخاب تصویر Open Graph
const onOgImageSelected = (selectedFiles: any[]) => {
  if (selectedFiles.length > 0) {
    postForm.og_image = selectedFiles[0].url
    ogImageTouched.value = true // کاربر دستی og_image را تغییر داد
  }
  showOgImageModal.value = false
}

// توابع مربوط به تگ‌ها
const addTag = () => {
  const tag = tagsInput.value.trim()
  if (tag && !postForm.tags?.includes(tag)) {
    if (!postForm.tags) postForm.tags = []
    postForm.tags.push(tag)
    tagsInput.value = ''
  }
}

const removeTag = (index: number) => {
  if (postForm.tags) {
    postForm.tags.splice(index, 1)
  }
}

// تابع محاسبه تعداد کلمات و زمان مطالعه
const calculateWordCountAndReadingTime = () => {
  const content = postForm.content || ''
  const excerpt = postForm.excerpt || ''
  const fullText = content + ' ' + excerpt
  
  // حذف HTML tags
  const cleanText = fullText.replace(/<[^>]*>/g, '')
  
  // شمارش کلمات (فاصله‌ها و کاراکترهای خاص)
  const words = cleanText.trim().split(/\s+/).filter(word => word.length > 0)
  postForm.word_count = words.length
  
  // محاسبه زمان مطالعه (میانگین 200 کلمه در دقیقه)
  postForm.reading_time = Math.ceil(words.length / 200)
}

// داده‌های SEO برای پیش‌نمایش
const seoDataForPreview = computed(() => ({
  title: postForm.title,
  excerpt: postForm.excerpt,
  content: postForm.content,
  slug: postForm.slug,
  meta_title: postForm.meta_title,
  meta_description: postForm.meta_description,
  meta_keywords: postForm.meta_keywords,
  og_image: postForm.og_image,
  og_type: postForm.og_type,
  schema_type: 'Article',
  index_status: postForm.index_status,
  follow_status: postForm.follow_status,
  author_name: 'نویسنده', // این مقدار باید از سیستم احراز هویت گرفته شود
  category_name: categories.value.find(c => c.id === postForm.category_id)?.name || '',
  word_count: postForm.word_count,
  reading_time: postForm.reading_time,
  tags: postForm.tags,
  language: postForm.language,
  site_url: 'https://example.com' // این مقدار باید از تنظیمات گرفته شود
}))

const updateSlug = (newSlug: string) => {
  postForm.slug = newSlug
}

const updateOgImage = (newOgImage: string) => {
  postForm.og_image = newOgImage
}



function showHelp(field: string) {
  helpField.value = helpField.value === field ? '' : field
}

const savePost = async (): Promise<{ id?: number } | null> => {
  if (!postForm.title.trim()) {
    alert('لطفاً عنوان نوشته را وارد کنید')
    return null
  }

  // اگر اسلاگ خالی بود، از عنوان تولید کن
  if (!postForm.slug.trim()) {
    postForm.slug = slugify(postForm.title)
    updateSlug(postForm.slug)
  }
  
  // بررسی الزامی بودن URL
  if (!postForm.slug.trim()) {
    alert('URL الزامی است. لطفاً URL را وارد کنید یا از عنوان تولید کنید.')
    return null
  }

  isSaving.value = true
  
  try {
    // آماده‌سازی داده‌ها برای ارسال
    const postData: any = {
      title: postForm.title,
      excerpt: postForm.excerpt,
      content: postForm.content,
      status: postForm.status,
      category_id: postForm.category_id || null,
      featured_image: postForm.featured_image,
      og_image: postForm.og_image, // تصویر Open Graph
      og_type: postForm.og_type, // نوع محتوا Open Graph
      slug: postForm.slug,
      meta_title: postForm.meta_title,

      meta_description: postForm.meta_description,
      
      // فیلدهای SEO جدید
      robots_meta: postForm.robots_meta,
      
      // فیلدهای جدید
      word_count: postForm.word_count,
      reading_time: postForm.reading_time,
      difficulty: postForm.difficulty,
      estimated_time: postForm.estimated_time,
      prerequisites: postForm.prerequisites,
      news_category: postForm.news_category,
      location: postForm.location,
      event_date: postForm.event_date,
      product_price: postForm.product_price,
      rating: postForm.rating,
      pros: postForm.pros,
      cons: postForm.cons,
      tags: postForm.tags,
      language: postForm.language,
      is_accessible_for_free: true, // همیشه true باشد
    }

    interface ApiResponse {
      id?: number
    }
    const response = await $fetch<ApiResponse>(`/api/posts/${postId}`, {
      method: 'PUT',
      body: postData
    })

    if (response) {
      return response
    }
    return null
  } catch (error) {
    console.error('خطا در ذخیره نوشته:', error)
    alert('خطا در ذخیره نوشته')
    return null
  } finally {
    isSaving.value = false
  }
}

// ذخیره و ادامه ویرایش
const saveAndContinueEditing = async () => {
  try {
    const response = await savePost()
    if (response) {
      alert('نوشته با موفقیت ذخیره و آماده ویرایش است!')
    }
  } catch (error) {
    console.error('خطا در ذخیره و ادامه ویرایش:', error)
  }
}

// ذخیره و هدایت به همه نوشته‌ها
const saveAndGoToList = async () => {
  try {
    const response = await savePost()
    if (response) {
      alert('نوشته با موفقیت ذخیره شد')
      // هدایت به صفحه همه نوشته‌ها
      navigateTo('/admin/post-management/all-posts')
    }
  } catch (error) {
    console.error('خطا در ذخیره و هدایت:', error)
  }
}

// مدیریت انتخاب تگ h
const handleHTagSelected = (event: MouseEvent) => {
  // نمایش dropdown در موقعیت کلیک
  dropdownPosition.value = {
    x: event.clientX,
    y: event.clientY
  }
  showTitleDropdown.value = true
}

// مدیریت انتخاب عنوان از dropdown
const handleTitleSelected = (title: string, titleId: number) => {
  const anchor = `id-${titleId}`
  const linkText = `<a href="#${anchor}" class="text-blue-600 hover:text-blue-800">${title}</a>`
  
  // درج لینک در محتوای اصلی
  if (postForm.content) {
    postForm.content += `\n\n${linkText}`
  } else {
    postForm.content = linkText
  }
  
  showTitleDropdown.value = false
}

// پیش‌نمایش نوشته
const previewPost = async () => {
  try {
    // ابتدا نوشته را ذخیره کن
    const savedPost = await savePost()
    
    if (savedPost && savedPost.id) {
      // اگر ذخیره موفق بود، پیش‌نمایش را باز کن
      if (postForm.slug && getCategorySlug(postForm.category_id)) {
        window.open(`/blog/${getCategorySlug(postForm.category_id)}/${postForm.slug}`, '_blank')
      } else if (postForm.slug) {
        window.open(`/blog/${postForm.slug}`, '_blank')
      }
    }
  } catch (error) {
    console.error('خطا در پیش‌نمایش:', error)
  }
}

// تابع کمکی برای گرفتن slug دسته‌بندی
const getCategorySlug = (categoryId: number | null) => {
  if (!categoryId) return ''
  const category = categories.value.find(c => c.id === categoryId)
  return category ? category.slug : ''
}



// Lifecycle
onMounted(async () => {
  fetchPost()
  fetchCategories()
})
</script> 