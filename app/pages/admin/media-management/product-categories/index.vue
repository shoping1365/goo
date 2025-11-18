<template>
  <div>
    <!-- استفاده از همان کامپوننت library با دسته‌بندی product-categories -->
    <MediaLibrary category="product-categories" />
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
declare const useHead: (head: { title?: string }) => void
</script>

<script setup lang="ts">
import MediaLibrary from '../library/index.vue';

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

useHead({
  title: 'کتابخانه رسانه - دسته‌بندی محصولات - پنل مدیریت'
})
</script>
