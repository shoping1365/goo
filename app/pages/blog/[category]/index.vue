<template>
  <div class="min-h-screen bg-white">
    <!-- Header -->
    <div class="bg-white shadow-sm border-b border-gray-200">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="py-4">
          <nav class="flex" aria-label="Breadcrumb">
            <ol class="flex items-center space-x-4 space-x-reverse">
              <li>
                <div>
                  <NuxtLink to="/" class="text-gray-400 hover:text-gray-500">
                    <svg class="flex-shrink-0 h-5 w-5" fill="currentColor" viewBox="0 0 20 20">
                      <path d="M10.707 2.293a1 1 0 00-1.414 0l-7 7a1 1 0 001.414 1.414L4 10.414V17a1 1 0 001 1h2a1 1 0 001-1v-2a1 1 0 011-1h2a1 1 0 011 1v2a1 1 0 001 1h2a1 1 0 001-1v-6.586l.293.293a1 1 0 001.414-1.414l-7-7z" />
                    </svg>
                  </NuxtLink>
                </div>
              </li>
              <li>
                <div class="flex items-center">
                  <svg class="flex-shrink-0 h-5 w-5 text-gray-300" fill="currentColor" viewBox="0 0 20 20">
                    <path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd" />
                  </svg>
                  <span class="mr-4 text-sm font-medium text-gray-500">{{ categoryData?.name }}</span>
                </div>
              </li>
            </ol>
          </nav>
        </div>
      </div>
    </div>

    <!-- Category Header -->
    <div v-if="categoryData" class="bg-white border-b border-gray-200">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div class="text-center">
          <h1 class="text-3xl font-bold text-gray-900 mb-4">{{ categoryData.name }}</h1>
          <p v-if="categoryData.description" class="text-lg text-gray-600 max-w-2xl mx-auto">
            {{ categoryData.description }}
          </p>
        </div>
      </div>
    </div>

    <!-- Main Content -->
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <div v-if="pending" class="text-center py-12">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600 mx-auto"></div>
        <p class="mt-4 text-gray-600">در حال بارگذاری...</p>
      </div>

      <div v-else-if="error" class="text-center py-12">
        <div class="text-red-600">
          <svg class="mx-auto h-12 w-12" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
          </svg>
          <h3 class="mt-2 text-sm font-medium text-gray-900">خطا در بارگذاری</h3>
          <p class="mt-1 text-sm text-gray-500">{{ error?.value?.message || 'خطای نامشخص' }}</p>
        </div>
      </div>

      <div v-else>
        <!-- Posts Grid -->
        <div v-if="postsList.length > 0" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gapx-4 py-4">
          <article 
            v-for="post in postsList" 
            :key="post.id"
            class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden hover:shadow-md transition-shadow"
          >
            <NuxtLink :to="`/blog/${categoryData?.slug}/${post.slug}`">
              <div v-if="post.featured_image" class="aspect-w-16 aspect-h-9">
                <img 
                  :src="post.featured_image" 
                  :alt="post.title"
                  class="w-full h-48 object-cover"
                >
              </div>
              <div class="px-4 py-4">
                <div class="flex items-center text-sm text-gray-500 mb-3">
                  <span>{{ formatDate(post.created_at) }}</span>
                  <span v-if="post.author" class="mx-2">•</span>
                  <span v-if="post.author">{{ post.author }}</span>
                </div>
                
                <h2 class="text-xl font-semibold text-gray-900 mb-3 line-clamp-2">{{ post.title }}</h2>
                
                <div>
                  <div style="color:red; font-size:12px;">excerpt: {{ post.excerpt }}</div>
                  <div v-html="post.excerpt" style="border:1px dashed #aaa; padding:4px; margin-bottom:8px;"></div>
                </div>
                
                <div class="flex items-center justify-between">
                  <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800">
                    {{ categoryData?.name }}
                  </span>
                  
                  <span class="text-sm text-gray-500">
                    ادامه مطلب
                    <svg class="inline-block w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
                    </svg>
                  </span>
                </div>
              </div>
            </NuxtLink>
          </article>
        </div>

        <!-- Empty State -->
        <div v-else class="text-center py-12">
          <div class="text-gray-400">
            <svg class="mx-auto h-12 w-12" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
            </svg>
            <h3 class="mt-2 text-sm font-medium text-gray-900">نوشته‌ای یافت نشد</h3>
            <p class="mt-1 text-sm text-gray-500">در حال حاضر نوشته‌ای در این دسته‌بندی وجود ندارد.</p>
          </div>
        </div>

                 <!-- Pagination -->
         <div v-if="postsList.length > 0" class="mt-8 flex justify-center">
           <nav class="flex items-center space-x-2 space-x-reverse">
             <span class="text-sm text-gray-500">
               تعداد {{ postsList.length }} نوشته در این دسته‌بندی
             </span>
           </nav>
         </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string }) => void
declare const useFetch: <T = unknown>(url: string, options?: { transform?: (data: any) => T }) => Promise<{ data: { value: T }; pending: { value: boolean }; error: { value: Error | null } }>
declare const useHead: (head: { title?: string; meta?: Array<{ name?: string; content?: string }> }) => void
declare const useRoute: () => { params: Record<string, string>; query: Record<string, string> }

// تعریف interface برای نوشته
interface Post {
  id: number
  title: string
  slug: string
  excerpt: string
  content: string
  status: string
  category_id: number
  featured_image: string
  author: string
  created_at: string
  updated_at: string
}

// تعریف interface برای دسته‌بندی
interface Category {
  id: number
  name: string
  slug: string
  description: string
  is_active: boolean
}
</script>

<script setup lang="ts">
import { computed } from 'vue';

definePageMeta({
  layout: 'default'
})

const route = useRoute()
const { category: categorySlug } = route.params

// دریافت اطلاعات دسته‌بندی
const { data: category, pending: categoryPending, error: categoryError } = await useFetch<Category | null>(`/api/post-categories?all=1`, {
  transform: (data: any[]) => {
    if (Array.isArray(data)) {
      return data.find(cat => cat.slug === categorySlug) as Category | null
    }
    return null
  }
})

// دریافت نوشته‌های دسته‌بندی
const { data: posts, pending, error } = await useFetch<Post[]>(`/api/posts?all=1`, {
  transform: (data: any[]) => {
    if (Array.isArray(data) && category.value) {
      return data.filter(p => p.category_id === category.value?.id && p.status === 'published') as Post[]
    }
    return []
  }
})

// Computed properties برای unwrap کردن reactive refs
const categoryData = computed(() => category.value || null)
const postsList = computed(() => posts.value || [])

// تنظیم SEO
useHead({
  title: category.value ? `${category.value.name} - وبلاگ` : 'دسته‌بندی',
  meta: [
    {
      name: 'description',
      content: category.value?.description || ''
    }
  ]
})

// تابع فرمت تاریخ
const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleDateString('fa-IR', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}
</script> 