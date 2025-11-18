<script setup>
// صفحه redirect برای مسیر /admin/attributes
// این صفحه کاربر را به مسیر صحیح /admin/product-management/attributes هدایت می‌کند

definePageMeta({
  middleware: 'redirect'
})
</script>

<template>
  <div class="flex items-center justify-center min-h-screen">
    <div class="text-center">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600 mx-auto mb-4"></div>
      <p class="text-gray-600">در حال انتقال...</p>
    </div>
  </div>
</template> 
