<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-6">
      <h3 class="text-lg font-bold text-gray-800">پیش‌نمایش SEO</h3>
      <button 
        :disabled="isLoading"
        class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed text-sm"
        @click="refreshPreview"
      >
        <span v-if="isLoading">در حال بارگذاری...</span>
        <span v-else>بروزرسانی</span>
      </button>
    </div>

    <!-- پیش‌نمایش Google -->
    <div class="mb-8">
      <h4 class="text-md font-semibold text-gray-700 mb-3 flex items-center">
        <svg class="w-5 h-5 mr-2" viewBox="0 0 24 24" fill="currentColor">
          <path d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"/>
          <path d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"/>
          <path d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"/>
          <path d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"/>
        </svg>
        پیش‌نمایش Google
      </h4>
      <div class="border border-gray-300 rounded-lg p-6 bg-gray-50">
        <div v-if="previewData" class="space-y-2">
          <div class="text-blue-600 text-sm">{{ previewData.canonical_url }}</div>
          <div class="text-xl text-blue-800 font-medium">{{ previewData.meta_title || 'عنوان صفحه' }}</div>
          <div class="text-sm text-gray-600">{{ previewData.meta_description || 'توضیحات صفحه' }}</div>
        </div>
        <div v-else class="text-gray-500 text-sm">اطلاعات پیش‌نمایش در دسترس نیست</div>
      </div>
    </div>

    <!-- پیش‌نمایش Facebook -->
    <div class="mb-8">
      <h4 class="text-md font-semibold text-gray-700 mb-3 flex items-center">
        <svg class="w-5 h-5 mr-2" viewBox="0 0 24 24" fill="currentColor">
          <path d="M24 12.073c0-6.627-5.373-12-12-12s-12 5.373-12 12c0 5.99 4.388 10.954 10.125 11.854v-8.385H7.078v-3.47h3.047V9.43c0-3.007 1.792-4.669 4.533-4.669 1.312 0 2.686.235 2.686.235v2.953H15.83c-1.491 0-1.956.925-1.956 1.874v2.25h3.328l-.532 3.47h-2.796v8.385C19.612 23.027 24 18.062 24 12.073z"/>
        </svg>
        پیش‌نمایش Facebook
      </h4>
      <div class="border border-gray-300 rounded-lg p-6 bg-gray-50">
        <div v-if="previewData" class="space-y-2">
          <div class="text-sm text-gray-500">{{ previewData.og_site_name }}</div>
          <div class="text-lg font-medium">{{ previewData.og_title || 'عنوان Open Graph' }}</div>
          <div class="text-sm text-gray-600">{{ previewData.og_description || 'توضیحات Open Graph' }}</div>
          <div v-if="seoData.featured_image" class="mt-2">
            <img :src="seoData.featured_image" alt="تصویر Open Graph" class="w-full h-32 object-cover rounded" />
          </div>
        </div>
        <div v-else class="text-gray-500 text-sm">اطلاعات پیش‌نمایش در دسترس نیست</div>
      </div>
    </div>

    <!-- پیش‌نمایش Twitter -->
    <div class="mb-8">
      <h4 class="text-md font-semibold text-gray-700 mb-3 flex items-center">
        <svg class="w-5 h-5 mr-2" viewBox="0 0 24 24" fill="currentColor">
          <path d="M23.953 4.57a10 10 0 01-2.825.775 4.958 4.958 0 002.163-2.723c-.951.555-2.005.959-3.127 1.184a4.92 4.92 0 00-8.384 4.482C7.69 8.095 4.067 6.13 1.64 3.162a4.822 4.822 0 00-.666 2.475c0 1.71.87 3.213 2.188 4.096a4.904 4.904 0 01-2.228-.616v.06a4.923 4.923 0 003.946 4.827 4.996 4.996 0 01-2.212.085 4.936 4.936 0 004.604 3.417 9.867 9.867 0 01-6.102 2.105c-.39 0-.779-.023-1.17-.067a13.995 13.995 0 007.557 2.209c9.053 0 13.998-7.496 13.998-13.985 0-.21 0-.42-.015-.63A9.935 9.935 0 0024 4.59z"/>
        </svg>
        پیش‌نمایش Twitter
      </h4>
      <div class="border border-gray-300 rounded-lg p-6 bg-gray-50">
        <div v-if="previewData" class="space-y-2">
          <div class="text-lg font-medium">{{ previewData.og_title || 'عنوان Twitter' }}</div>
          <div class="text-sm text-gray-600">{{ previewData.og_description || 'توضیحات Twitter' }}</div>
          <div v-if="seoData.featured_image" class="mt-2">
            <img :src="seoData.featured_image" alt="تصویر Twitter" class="w-full h-32 object-cover rounded" />
          </div>
        </div>
        <div v-else class="text-gray-500 text-sm">اطلاعات پیش‌نمایش در دسترس نیست</div>
      </div>
    </div>

    <!-- نمایش تگ‌های تولید شده -->
    <div class="space-y-6">
      <div>
        <h4 class="text-md font-semibold text-gray-700 mb-3">تگ‌های HTML Meta</h4>
        <div class="bg-gray-900 text-green-400 p-6 rounded-lg overflow-x-auto">
          <pre class="text-sm">{{ seoTags.html_meta_tags || 'تگ‌های HTML Meta تولید نشده' }}</pre>
        </div>
      </div>

      <div>
        <h4 class="text-md font-semibold text-gray-700 mb-3">تگ‌های Open Graph</h4>
        <div class="bg-gray-900 text-green-400 p-6 rounded-lg overflow-x-auto">
          <pre class="text-sm">{{ seoTags.open_graph_tags || 'تگ‌های Open Graph تولید نشده' }}</pre>
        </div>
      </div>

      <div>
        <h4 class="text-md font-semibold text-gray-700 mb-3">تگ‌های Twitter Card</h4>
        <div class="bg-gray-900 text-green-400 p-6 rounded-lg overflow-x-auto">
          <pre class="text-sm">{{ seoTags.twitter_card_tags || 'تگ‌های Twitter Card تولید نشده' }}</pre>
        </div>
      </div>

      <div>
        <h4 class="text-md font-semibold text-gray-700 mb-3">JSON-LD Schema</h4>
        <div class="bg-gray-900 text-green-400 p-6 rounded-lg overflow-x-auto">
          <pre class="text-sm">{{ seoTags.json_ld_schema || 'JSON-LD Schema تولید نشده' }}</pre>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch, onMounted } from 'vue'
import { useSeoGenerator, type SeoData, type SeoTags, type SeoPreview } from '~/composables/useSeoGenerator'

interface Props {
  seoData: SeoData
}

const props = defineProps<Props>()

const { generateAllSeoTags, previewSeoTags } = useSeoGenerator()

const seoTags = ref<SeoTags>({
  html_meta_tags: '',
  open_graph_tags: '',
  twitter_card_tags: '',
  json_ld_schema: '',
  og_title: '',
  og_description: '',
  og_site_name: ''
})

const previewData = ref<SeoPreview['preview'] | null>(null)
const isLoading = ref(false)

// تولید تگ‌های SEO
const generateSeo = async () => {
  isLoading.value = true
  try {
    // تولید تگ‌های SEO به صورت محلی
    seoTags.value = generateAllSeoTags(props.seoData)
    
    // دریافت پیش‌نمایش از سرور
    const preview = await previewSeoTags(props.seoData)
    if (preview) {
      previewData.value = preview.preview
    }
  } catch (_error) {
    // خطا در تولید SEO
  } finally {
    isLoading.value = false
  }
}

// بروزرسانی پیش‌نمایش
const refreshPreview = () => {
  generateSeo()
}

// نظارت بر تغییرات داده‌های SEO
watch(() => props.seoData, () => {
  generateSeo()
}, { deep: true })

// تولید اولیه
onMounted(() => {
  generateSeo()
})
</script> 