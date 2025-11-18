<template>
  <div class="container mx-auto px-4 py-8">
    <h1 class="text-3xl font-bold text-center mb-8">برندها</h1>
    
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gapx-4 py-4">
      <!-- اینجا لیست برندها نمایش داده می‌شود -->
      <div class="text-center text-gray-500">
        در حال بارگذاری برندها...
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string }) => void
declare const useHead: (head: { title?: string; meta?: Array<{ name?: string; content?: string }> }) => void
</script>

<script setup lang="ts">
// تعریف متا صفحه
definePageMeta({
  layout: 'default'
})

// تنظیم عنوان صفحه
useHead({
  title: 'برندها - فروشگاه آنلاین',
  meta: [
    { name: 'description', content: 'لیست تمام برندهای موجود در فروشگاه' }
  ]
})
</script>
