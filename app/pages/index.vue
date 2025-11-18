<template>
  <div class="min-h-screen bg-white" dir="rtl">
    <!-- Header Section -->
    
    <WidgetArea page="home" />


    
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { ssr?: boolean }) => void
declare const useHead: (head: { title?: string; meta?: Array<{ name?: string; content?: string }> }) => void
</script>

<script setup lang="ts">
import WidgetArea from '~/components/common/WidgetArea.vue';


// تعریف متادیتای صفحه
definePageMeta({
  ssr: true
})

// تنظیم هد صفحه
useHead({
  title: 'فروشگاه آنلاین - صفحه اصلی',
  meta: [
    { name: 'description', content: 'صفحه اصلی فروشگاه آنلاین' }
  ]
})
</script> 
