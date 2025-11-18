<template>
  <div class="product-card bg-white rounded-lg shadow-md overflow-hidden" style="font-family: 'IranYekan', 'Vazir', 'IRANSansX', 'Tahoma', sans-serif;">
    <div class="relative">
      <img
        :src="product.image"
        :alt="product.title"
        :title="product.title"
        class="w-full h-48 object-cover"
      />
      <div v-if="product.discount" class="absolute top-2 right-2 bg-red-500 text-white px-2 py-1 rounded">
        {{ product.discount }}% تخفیف
      </div>
      <div v-if="product.isNew" class="absolute top-2 left-2 bg-green-500 text-white px-2 py-1 rounded">
        جدید
      </div>
    </div>
    <div class="p-6">
      <h3 class="text-lg font-semibold mb-2">{{ product.title }}</h3>
      <p class="text-gray-600 text-sm mb-4">{{ product.title }}</p>
      <div class="flex justify-between items-center">
        <div>
          <span class="text-lg font-bold text-primary">{{ formatPrice(product.price) }}</span>
          <span v-if="product.original_price && product.original_price > product.price" class="text-sm text-gray-500 line-through mr-2">
            {{ formatPrice(product.original_price) }}
          </span>
        </div>
        <div class="flex items-center">
          <span class="text-yellow-500 mr-1">★</span>
          <span class="text-sm">{{ product.rating }}</span>
        </div>
      </div>
      <div class="mt-4">
        <button
          @click="addToCart"
          class="w-full bg-primary text-white py-2 rounded hover:bg-primary-dark transition-colors"
        >
          افزودن به سبد خرید
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Product } from '~/types/product'

const props = defineProps<{
  product: any
}>()

const emit = defineEmits<{
  (e: 'add-to-cart', product: any): void
}>()

const formatPrice = (price: number) => {
  return new Intl.NumberFormat('fa-IR', {
    style: 'currency',
    currency: 'IRR'
  }).format(price)
}

const calculateOriginalPrice = (price: number, discount: number) => {
  return Math.round(price / (1 - discount / 100))
}

const addToCart = () => {
  emit('add-to-cart', props.product)
}
</script> 
