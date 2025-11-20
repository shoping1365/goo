<template>
  <div class="flex flex-col items-end text-right px-4 py-4 border-b border-gray-100 bg-gradient-to-l from-blue-50 to-white">
    <!-- عنوان مقاله -->
    <h1 class="text-3xl font-bold text-gray-900 mb-2">{{ title }}</h1>
    <!-- اطلاعات نویسنده و تاریخ -->
    <div class="flex flex-row-reverse items-center text-sm text-gray-500 gap-2">
      <span v-if="author">نویسنده: {{ author }}</span>
      <span v-if="author">|</span>
      <span>{{ date }}</span>
      <span>|</span>
      <span>{{ category }}</span>
    </div>
    <div v-if="excerpt" class="text-lg text-gray-600 mt-4 leading-relaxed">{{ excerpt }}</div>
  </div>
</template>

<script setup lang="ts">
// props
defineProps({
  title: String,
  author: String,
  date: String,
  category: String,
  excerpt: String
})
</script> 