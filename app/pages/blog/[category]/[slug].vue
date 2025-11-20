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
                  <NuxtLink :to="`/blog/${category?.value?.slug}`" class="mr-4 text-sm font-medium text-gray-500 hover:text-gray-700">
                    {{ category?.value?.name }}
                  </NuxtLink>
                </div>
              </li>
              <li>
                <div class="flex items-center">
                  <svg class="flex-shrink-0 h-5 w-5 text-gray-300" fill="currentColor" viewBox="0 0 20 20">
                    <path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd" />
                  </svg>
                  <span class="mr-4 text-sm font-medium text-gray-500">{{ post?.title }}</span>
                </div>
              </li>
            </ol>
          </nav>
        </div>
      </div>
    </div>

    <!-- Main Content -->
    <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <!-- نشانگر پیش‌نمایش -->
      <div v-if="route.query.preview" class="bg-yellow-100 border border-yellow-400 text-yellow-700 px-4 py-3 rounded mb-6">
        <div class="flex items-center">
          <svg class="w-5 h-5 mr-2" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd" d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z" clip-rule="evenodd"></path>
          </svg>
          <span class="font-medium">پیش‌نمایش</span>
        </div>
        <p class="text-sm mt-1">این نوشته در حالت پیش‌نمایش است و هنوز منتشر نشده است.</p>
      </div>
      
      <article v-if="post" class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
        <!-- Featured Image -->
        <div v-if="post.featured_image" class="aspect-w-16 aspect-h-9">
          <img 
            :src="post.featured_image" 
            :alt="post.title"
            class="w-full h-64 object-cover"
          >
        </div>

        <!-- Article Header -->
        <div class="px-4 py-4">
          <div class="flex items-center text-sm text-gray-500 mb-4">
            <span>{{ formatDate(post.created_at) }}</span>
            <span class="mx-2">•</span>
            <span>{{ category?.value?.name }}</span>
            <span v-if="post.author" class="mx-2">•</span>
            <span v-if="post.author">{{ post.author }}</span>
          </div>

          <h1 class="text-3xl font-bold text-gray-900 mb-4">{{ post.title }}</h1>
          
          <div v-if="post.excerpt" class="text-lg text-gray-600 mb-6 leading-relaxed">
            {{ post.excerpt }}
          </div>
        </div>

        <!-- Article Content -->
        <div class="px-6 pb-6">
          <!-- eslint-disable-next-line vue/no-v-html -->
          <div class="prose prose-lg max-w-none" v-html="sanitizedContent"></div>
        </div>

        <!-- Article Footer -->
        <div class="px-6 py-4 bg-gray-50 border-t border-gray-200">
          <div class="flex items-center justify-between">
            <div class="flex items-center space-x-4 space-x-reverse">
              <span class="text-sm text-gray-500">دسته‌بندی:</span>
              <NuxtLink 
                :to="`/blog/${category?.value?.slug}`"
                class="text-sm font-medium text-blue-600 hover:text-blue-800"
              >
                {{ category?.value?.name }}
              </NuxtLink>
            </div>
            
            <div class="flex items-center space-x-4 space-x-reverse">
              <button class="text-gray-400 hover:text-gray-600">
                <svg class="h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.367 2.684 3 3 0 00-5.367-2.684z"></path>
                </svg>
              </button>
              <button class="text-gray-400 hover:text-gray-600">
                <svg class="h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"></path>
                </svg>
              </button>
            </div>
          </div>
        </div>
      </article>

      <!-- Related Posts -->
      <div v-if="relatedPostsList.length > 0" class="mt-8">
        <h2 class="text-2xl font-bold text-gray-900 mb-6">نوشته‌های مرتبط</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gapx-4 py-4">
          <div 
            v-for="relatedPost in relatedPostsList" 
            :key="relatedPost.id"
            class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden hover:shadow-md transition-shadow"
          >
            <NuxtLink :to="`/blog/${category?.value?.slug}/${relatedPost.slug}`">
              <div v-if="relatedPost.featured_image" class="aspect-w-16 aspect-h-9">
                <img 
                  :src="relatedPost.featured_image" 
                  :alt="relatedPost.title"
                  class="w-full h-48 object-cover"
                >
              </div>
              <div class="px-4 py-4">
                <h3 class="text-lg font-semibold text-gray-900 mb-2 line-clamp-2">{{ relatedPost.title }}</h3>
                <p v-if="relatedPost.excerpt" class="text-gray-600 text-sm line-clamp-3">{{ relatedPost.excerpt }}</p>
                <div class="mt-3 text-sm text-gray-500">{{ formatDate(relatedPost.created_at) }}</div>
              </div>
            </NuxtLink>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string }) => void
declare const useFetch: <T = unknown>(url: string, options?: { transform?: (data: unknown) => T }) => Promise<{ data: { value: T }; pending: { value: boolean }; error: { value: Error | null } }>
declare const useHead: (head: { title?: string; meta?: Array<{ name?: string; content?: string }> }) => void
declare const useRoute: () => { params: Record<string, string>; query: Record<string, string> }
declare const $fetch: <T = unknown>(url: string) => Promise<T>
declare const createError: (options: { statusCode: number; statusMessage: string }) => Error

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
import DOMPurify from 'dompurify';
import { computed, ref } from 'vue';

definePageMeta({
  layout: 'default'
})

const route = useRoute()
const { category: categorySlug, slug: postSlug } = route.params

const sanitizedContent = computed(() => {
  if (!post.value?.content) return '';
  return DOMPurify.sanitize(post.value.content);
});

// دریافت اطلاعات دسته‌بندی
const { data: category } = await useFetch<Category | null>(`/api/post-categories?all=1`, {
  transform: (data: unknown) => {
    if (Array.isArray(data)) {
      return (data as Category[]).find(cat => cat.slug === categorySlug) || null
    }
    return null
  }
})

// دریافت اطلاعات نوشته به صورت مستقیم
const post = ref<Post | null>(null)
try {
  // بررسی پارامتر preview از URL
  const isPreview = route.query.preview === '1' || route.query.preview === 'true'
  const apiUrl = isPreview ? `/api/posts/${postSlug}?preview=1` : `/api/posts/${postSlug}`
  
  post.value = await $fetch<Post>(apiUrl)
} catch (_e) {
  throw createError({ statusCode: 404, statusMessage: 'نوشته یافت نشد' })
}

// دریافت نوشته‌های مرتبط
const { data: relatedPosts } = await useFetch<Post[]>(`/api/posts?limit=4&exclude=${post.value?.id}`)

// Computed برای unwrap کردن relatedPosts
const relatedPostsList = computed(() => {
  return relatedPosts.value || []
})

// تنظیم SEO
useHead({
  title: post.value ? `${post.value.title} - ${category.value?.name || ''}` : 'نوشته',
  meta: [
    {
      name: 'description',
      content: post.value?.excerpt || post.value?.content?.substring(0, 160) || ''
    },
    // در حالت پیش‌نمایش، موتورهای جستجو را مسدود کن
    ...(route.query.preview ? [
      { name: 'robots', content: 'noindex, nofollow' },
      { name: 'googlebot', content: 'noindex, nofollow' }
    ] : [])
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