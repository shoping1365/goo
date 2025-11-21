<template>
  <div class="max-w-3xl mx-auto py-8 px-4">
    <nav class="mb-4 text-sm text-gray-500 flex items-center space-x-2 rtl:space-x-reverse">
      <NuxtLink to="/blog" class="hover:underline">بلاگ</NuxtLink>
      <span>/</span>
      <span v-if="post">{{ post.title }}</span>
    </nav>

    <div v-if="pending" class="text-center text-gray-400 py-20">در حال بارگذاری...</div>

    <div v-else-if="error" class="text-center text-red-500 py-20">
      <h1 class="text-2xl font-bold mb-4">نوشته یافت نشد</h1>
      <p class="text-gray-600 mb-6">متأسفانه نوشته‌ای با این آدرس وجود ندارد.</p>
      <NuxtLink to="/blog" class="bg-blue-600 text-white px-6 py-3 rounded-lg hover:bg-blue-700 transition-colors">
        بازگشت به بلاگ
      </NuxtLink>
    </div>

    <div v-else-if="post">
      <!-- نشانگر پیش‌نمایش -->
      <div v-if="route.query.preview" class="bg-yellow-100 border border-yellow-400 text-yellow-700 px-4 py-3 rounded mb-4">
        <div class="flex items-center">
          <svg class="w-5 h-5 mr-2" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd" d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z" clip-rule="evenodd"></path>
          </svg>
          <span class="font-medium">پیش‌نمایش</span>
        </div>
        <p class="text-sm mt-1">این نوشته در حالت پیش‌نمایش است و هنوز منتشر نشده است.</p>
      </div>
      
      <h1 class="text-3xl font-bold mb-2">{{ post.title }}</h1>
      <div v-if="post.published_at" class="text-xs text-gray-400 mb-4">منتشر شده در {{ toPersianDate(post.published_at) }}</div>
      <img v-if="post.featured_image" :src="post.featured_image" :alt="post.title" class="rounded-lg mb-6 w-full object-cover max-h-96">
      
      <!-- eslint-disable-next-line vue/no-v-html -->
      <div v-if="post.excerpt" class="prose max-w-none mb-6 text-gray-600 border-r-4 border-gray-200 pr-4" v-html="sanitizedExcerpt"></div>
      
      <!-- eslint-disable-next-line vue/no-v-html -->
      <div class="prose prose-lg max-w-none mb-8" v-html="sanitizedContent"></div>

      <div v-if="relatedPosts.length" class="mt-12">
        <h2 class="text-lg font-semibold mb-4">نوشته‌های مرتبط</h2>
        <ul class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <li v-for="item in relatedPosts" :key="item.id" class="bg-gray-50 rounded-lg p-6 border hover:shadow">
            <NuxtLink :to="getPostLink(item)" class="font-bold hover:underline">{{ item.title }}</NuxtLink>
            <div class="text-xs text-gray-400 mt-1">{{ toPersianDate(item.published_at) }}</div>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const useHead: (head: { title?: string; meta?: Array<{ name?: string; content?: string }> }) => void
declare const useRoute: () => { params: Record<string, string>; query: Record<string, string> }
declare const $fetch: <T = unknown>(url: string) => Promise<T>

interface Post {
  id?: number | string
  slug?: string
  title?: string
  excerpt?: string
  content?: string
  published_at?: string
  featured_image?: string
  meta_title?: string
  meta_description?: string
  meta_keywords?: string
  category?: {
    slug?: string
  }
}
</script>

<script setup lang="ts">
import DOMPurify from 'dompurify';
import { computed, ref, watch } from 'vue';
import { toPersianDate } from '~/utils/dateUtils';

const route = useRoute()
const slug = route.params.slug

const post = ref<Post | null>(null)
const pending = ref(false)
const error = ref<{ statusCode: number; statusMessage: string } | null>(null)

const fetchPost = async () => {
  try {
    pending.value = true
    const isPreview = route.query.preview === '1' || route.query.preview === 'true'
    const apiUrl = isPreview ? `/api/posts/${slug}?preview=1` : `/api/posts/${slug}`
    
    post.value = await $fetch<Post>(apiUrl)
    error.value = null
  } catch (e: unknown) {
    // console.error('❌ Error fetching post:', e)
    const err = e as { statusCode?: number }
    if (err.statusCode === 404) {
      error.value = { statusCode: 404, statusMessage: 'نوشته یافت نشد' }
    } else {
      error.value = { statusCode: 500, statusMessage: 'خطا در دریافت اطلاعات نوشته' }
    }
    post.value = null
  } finally {
    pending.value = false
  }
}

// همیشه پست را fetch کن
await fetchPost()

const sanitizedExcerpt = computed(() => {
  if (!post.value?.excerpt) return '';
  return DOMPurify.sanitize(post.value.excerpt);
});

const sanitizedContent = computed(() => {
  if (!post.value?.content) return '';
  return DOMPurify.sanitize(post.value.content);
});

const relatedPosts = ref<Post[]>([])

watch(post, async (val) => {
  if (!val?.id) return
  try {
    const res = await $fetch<Post[]>(`/api/posts?limit=4&exclude=${val.id}`)
    relatedPosts.value = Array.isArray(res) ? res.filter((p: Post) => p.id !== val.id) : []
  } catch (_err) {
    // console.error('خطا در دریافت نوشته‌های مرتبط:', err)
  }
}, { immediate: true })

useHead({
  title: post.value ? (post.value.meta_title || post.value.title || 'نوشته') : 'نوشته',
  meta: [
    { name: 'description', content: post.value?.meta_description || '' },
    { name: 'keywords', content: post.value?.meta_keywords || '' },
    ...(route.query.preview ? [
      { name: 'robots', content: 'noindex, nofollow' },
      { name: 'googlebot', content: 'noindex, nofollow' }
    ] : [])
  ]
})

const getPostLink = (item: Post) => {
  if (item.category?.slug) {
    return `/blog/${item.category.slug}/${item.slug}`
  }
  return `/blog/${item.slug}`
}
</script>