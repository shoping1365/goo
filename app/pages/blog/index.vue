<template>
  <div class="max-w-3xl mx-auto py-8 px-4">
    <h1 class="text-3xl font-bold mb-8 text-center">بلاگ</h1>
    <div v-if="posts.length">
      <ul class="space-y-8">
        <li v-for="post in posts" :key="post.id" class="bg-white rounded-lg shadow px-4 py-4 border hover:shadow-md transition">
          <NuxtLink :to="`/blog/${post.slug}`" class="block">
            <h2 class="text-2xl font-semibold mb-2">{{ post.title }}</h2>
            <div class="text-xs text-gray-400 mb-2">منتشر شده در {{ toPersianDate(post.published_at) }}</div>
            <p class="text-gray-600 mb-2 line-clamp-3">{{ post.excerpt }}</p>
            <span class="text-blue-600 hover:underline text-sm">مشاهده نوشته</span>
          </NuxtLink>
        </li>
      </ul>
    </div>
    <div v-else class="text-center text-gray-500 py-20">هیچ نوشته‌ای منتشر نشده است.</div>
  </div>
</template>

<script lang="ts">
declare const useAsyncData: <T = unknown>(key: string, fn: () => Promise<T>) => Promise<{ data: { value: T } }>
declare const $fetch: <T = unknown>(url: string) => Promise<T>

interface Post {
  id: number | string
  slug: string
  title: string
  excerpt?: string
  published_at?: string
}
</script>

<script setup lang="ts">
// این صفحه لیست پست‌های منتشرشده را نمایش می‌دهد
import { computed } from 'vue'
import { toPersianDate } from '~/utils/dateUtils'

// دریافت پست‌ها از API
const { data: postsData } = await useAsyncData<Post[]>('posts', async () => {
  try {
    // فقط پست‌های منتشرشده را دریافت کن
    const result = await $fetch<Post[]>('/api/posts')
    return Array.isArray(result) ? result : []
  } catch (e) {
    return []
  }
})

const posts = computed(() => postsData.value || [])
</script>

<!--
  استایل‌ها با Tailwind v4 و کاملاً واکنش‌گرا
  هیچ استایل سفارشی اضافه نشده است
--> 