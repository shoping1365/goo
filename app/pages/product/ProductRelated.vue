<template>
  <div class="space-y-6" dir="rtl">
    <div v-if="relatedProducts.length" class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
      <NuxtLink
        v-for="relatedProduct in relatedProducts"
        :key="relatedProduct.id"
        :to="`/product/sku-${relatedProduct.sku || relatedProduct.id}${relatedProduct.slug ? '/' + encodeURIComponent(relatedProduct.slug) : ''}`"
        class="bg-white rounded-lg border border-gray-200 overflow-hidden hover:shadow-lg transition-shadow group"
      >
        <div class="aspect-square bg-gray-100 overflow-hidden">
          <img
            :src="toThumbnail(relatedProduct.main_image)"
            :alt="relatedProduct.name"
            class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
          />
        </div>
        <div class="p-6">
          <h3 class="font-medium text-gray-900 text-sm mb-2 line-clamp-2">{{ relatedProduct.name }}</h3>
          <div class="flex items-center justify-between">
            <div class="text-green-600 font-bold text-sm">
              {{ formatPrice(product.price) }} تومان
            </div>
            <div v-if="product.rating" class="flex items-center gap-1">
              <svg class="w-4 h-4 text-yellow-400 fill-current" viewBox="0 0 24 24">
                <path d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"/>
              </svg>
              <span class="text-xs text-gray-600">{{ product.rating }}</span>
            </div>
          </div>
        </div>
      </NuxtLink>
    </div>

    <div v-else class="text-center py-12">
      <div class="text-gray-400 mb-4">
        <svg class="w-16 h-16 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          
        </svg>
      </div>
      <h3 class="text-lg font-medium text-gray-900 mb-2">محصول مرتبطی موجود نیست</h3>
      <p class="text-gray-500">در حال حاضر محصول مرتبطی برای نمایش وجود ندارد.</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface Product {
  id: number
  name: string
  main_image: string
  price: number
  rating?: number
  sku?: string
  slug?: string
}

interface Props {
  product: {
    related_products?: Product[]
  }
}

const props = defineProps<Props>()

const relatedProducts = computed(() => props.product.related_products || [])

const formatPrice = (price: number): string => {
  return new Intl.NumberFormat('fa-IR').format(price)
}

// تابع برای تبدیل تصویر به thumbnail
function toThumbnail(url: string): string {
  if (!url) return '/default-product.jpg'
  
  const dotIdx = url.lastIndexOf('.')
  if (dotIdx === -1) return url + '_thumbnail.webp'
  
  const baseName = url.slice(0, dotIdx)
  const originalExt = url.slice(dotIdx)
  
  // اگر فرمت اصلی jpg یا jpeg است، thumbnail را با webp بساز
  if (originalExt.toLowerCase() === '.jpg' || originalExt.toLowerCase() === '.jpeg') {
    return baseName + '_thumbnail.webp'
  }
  
  // برای سایر فرمت‌ها، همان فرمت را حفظ کن
  return baseName + '_thumbnail' + originalExt
}
</script>

<style scoped>
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style> 
