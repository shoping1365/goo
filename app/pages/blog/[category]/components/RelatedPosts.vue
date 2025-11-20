<template>
  <div v-if="posts && posts.length" class="mt-8">
    <h2 class="text-2xl font-bold text-gray-900 mb-6">نوشته‌های مرتبط</h2>
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gapx-4 py-4">
      <div v-for="post in posts" :key="post.id" class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden hover:shadow-md transition-shadow">
        <NuxtLink :to="`/blog/${category}/${post.slug}`">
          <div v-if="post.featured_image" class="aspect-w-16 aspect-h-9">
            <img :src="post.featured_image" :alt="post.title" class="w-full h-48 object-cover">
          </div>
          <div class="px-4 py-4">
            <h3 class="text-lg font-semibold text-gray-900 mb-2 line-clamp-2">{{ post.title }}</h3>
            <p v-if="post.excerpt" class="text-gray-600 text-sm line-clamp-3">{{ post.excerpt }}</p>
            <div class="mt-3 text-sm text-gray-500">{{ formatDate(post.created_at) }}</div>
          </div>
        </NuxtLink>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
interface Post {
  id: number | string
  slug: string
  title: string
  excerpt?: string
  featured_image?: string
  created_at: string
}

defineProps<{
  posts?: Post[]
  category?: string
  formatDate?: (date: string) => string
}>()
</script> 