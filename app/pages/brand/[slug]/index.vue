<template>
  <div class="max-w-6xl mx-auto py-10 px-4 lg:px-0" v-if="brand">
    <div class="flex items-center gapx-4 py-4 mb-8">
      <img :src="brand.image_url || '/statics/images/default-image_100.png'" alt="Brand" class="w-24 h-24 object-contain border rounded" />
      <h1 class="text-3xl font-bold">{{ brand.name }}</h1>
    </div>
    <!-- 
      ⚠️ امنیت XSS: استفاده از v-html خطرناک است!
      
      این کد محتوای HTML را بدون sanitization نمایش می‌دهد که می‌تواند منجر به حملات XSS شود.
      
      ✅ راه حل صحیح:
      1. قبل از استفاده از v-html، محتوا را با کتابخانه sanitization (مثل DOMPurify) پاکسازی کنید
      2. یا از {{ }} به جای v-html استفاده کنید اگر HTML نیاز نیست
      3. محتوای کاربر را هرگز بدون sanitization در v-html قرار ندهید
      
      مثال صحیح:
      import DOMPurify from 'dompurify'
      const sanitizedDescription = DOMPurify.sanitize(brand.description)
      <div v-html="sanitizedDescription"></div>
    -->
    <div class="prose max-w-none" v-html="brand.description"></div>
  </div>
  <div v-else class="py-20 text-center text-gray-500">برند یافت نشد</div>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router'
// useAsyncData is auto-imported in Nuxt 3
import { onMounted, ref } from 'vue'

const route = useRoute()
const slug = route.params.slug as string

const brand = ref(null)
const pending = ref(false)
const error = ref(null)

const fetchBrand = async () => {
  try {
    pending.value = true
    const isPreview = route.query.preview === '1' || route.query.preview === 'true'
    const apiUrl = isPreview ? `/api/brand/${slug}?preview=1` : `/api/brand/${slug}`
    brand.value = await $fetch(apiUrl)
    error.value = null
  } catch (e) {
    brand.value = null
    error.value = e
  } finally {
    pending.value = false
  }
}

if (import.meta.client && (route.query.preview === '1' || route.query.preview === 'true')) {
  onMounted(fetchBrand)
} else if (import.meta.server && !(route.query.preview === '1' || route.query.preview === 'true')) {
  await fetchBrand()
}
</script>

<style scoped>
.prose img {
  max-width: 100%;
  height: auto;
}
</style> 