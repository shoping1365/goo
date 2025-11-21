<template>
  <div class="product-carousel-demo-page p-6">
    <div class="max-w-6xl mx-auto">
      <!-- Page Header -->
      <div class="mb-8">
        <div class="flex items-center gap-6 mb-4">
          <NuxtLink 
            to="/admin/content/banners/ProductCarousel"
            class="text-blue-600 hover:text-blue-800 transition-colors"
          >
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"/>
            </svg>
          </NuxtLink>
          <h1 class="text-3xl font-bold text-gray-900">Product Carousel Demo</h1>
        </div>
        <p class="text-gray-600">Preview and test the ProductCarousel component with sample data</p>
      </div>

      <!-- Demo Controls -->
      <div class="bg-white rounded-lg shadow-lg p-6 mb-8">
        <h2 class="text-xl font-semibold text-gray-800 mb-4">Configuration Controls</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Slides Per View</label>
            <input 
              v-model.number="config.slidesPerView"
              type="number"
              min="1"
              max="6"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Autoplay Delay (ms)</label>
            <input 
              v-model.number="config.autoplayDelay"
              type="number"
              min="1000"
              max="10000"
              step="500"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Carousel Height (px)</label>
            <input 
              v-model.number="config.carouselHeight"
              type="number"
              min="300"
              max="600"
              step="50"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Slide Gap (px)</label>
            <input 
              v-model.number="config.slideGap"
              type="number"
              min="0"
              max="50"
              step="5"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            />
          </div>
        </div>
        
        <div class="mt-4 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-6">
          <label class="flex items-center">
            <input 
              v-model="config.autoplay"
              type="checkbox"
              class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
            />
            <span class="ml-2 text-sm text-gray-700">Autoplay</span>
          </label>
          <label class="flex items-center">
            <input 
              v-model="config.loop"
              type="checkbox"
              class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
            />
            <span class="ml-2 text-sm text-gray-700">Loop</span>
          </label>
          <label class="flex items-center">
            <input 
              v-model="config.showNavigation"
              type="checkbox"
              class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
            />
            <span class="ml-2 text-sm text-gray-700">Navigation</span>
          </label>
          <label class="flex items-center">
            <input 
              v-model="config.showIndicators"
              type="checkbox"
              class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
            />
            <span class="ml-2 text-sm text-gray-700">Indicators</span>
          </label>
          <label class="flex items-center">
            <input 
              v-model="config.showControls"
              type="checkbox"
              class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
            />
            <span class="ml-2 text-sm text-gray-700">Controls</span>
          </label>
        </div>
      </div>

      <!-- Product Carousel Demo -->
      <div class="bg-white rounded-lg shadow-lg p-6">
        <h2 class="text-xl font-semibold text-gray-800 mb-4">Product Carousel Preview</h2>
        <ProductCarousel
          :products="sampleProducts"
          :slides-per-view="config.slidesPerView"
          :autoplay="config.autoplay"
          :autoplay-delay="config.autoplayDelay"
          :loop="config.loop"
          :show-navigation="config.showNavigation"
          :show-indicators="config.showIndicators"
          :show-controls="config.showControls"
          :carousel-height="config.carouselHeight"
          :slide-gap="config.slideGap"
          @product-click="handleProductClick"
          @slide-change="handleSlideChange"
          @autoplay-toggle="handleAutoplayToggle"
        />
      </div>

      <!-- Event Log -->
      <div class="bg-white rounded-lg shadow-lg p-6 mt-8">
        <h2 class="text-xl font-semibold text-gray-800 mb-4">Event Log</h2>
        <div class="bg-gray-50 rounded-lg p-6 max-h-64 overflow-y-auto">
          <div v-if="eventLog.length === 0" class="text-gray-500 text-center py-8">
            No events logged yet. Interact with the carousel to see events here.
          </div>
          <div v-else class="space-y-2">
            <div 
              v-for="(event, index) in eventLog.slice().reverse()" 
              :key="index"
              class="text-sm p-2 bg-white rounded border-l-4"
              :class="getEventBorderClass(event.type)"
            >
              <span class="font-medium text-gray-700">{{ event.timestamp }}</span>
              <span class="ml-2 text-gray-600">{{ event.message }}</span>
            </div>
          </div>
        </div>
        <button 
          class="mt-4 px-4 py-2 bg-gray-600 text-white rounded-lg hover:bg-gray-700 transition-colors"
          @click="clearEventLog"
        >
          Clear Log
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// // import ProductCarousel from  '~/components/widgets/ProductCarousel.vue'

// Page meta definition
definePageMeta({
  layout: 'admin',
  middleware: 'admin'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
// const { user, hasPermission } = useAuth()

import { ref } from 'vue'

// تعریف definePageMeta و useHead برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
declare const useHead: (head: { title?: string }) => void

// Page title
useHead({
  title: 'Product Carousel Demo - Admin Panel'
})

// Sample products data
const sampleProducts = ref([
  {
    id: 1,
    name: 'iPhone 15 Pro Max',
    description: 'Latest iPhone with advanced camera system and A17 Pro chip',
    price: 45000000,
    image: 'https://images.unsplash.com/photo-1592750475338-74b7b21085ab?w=400&h=300&fit=crop',
    category: 'Electronics',
    rating: 5,
    stock: 50,
    brand: 'Apple',
    discount: 10,
    isNew: true,
    isFeatured: true
  },
  {
    id: 2,
    name: 'MacBook Air M2',
    description: 'Ultra-thin laptop with M2 chip and all-day battery life',
    price: 35000000,
    image: 'https://images.unsplash.com/photo-1541807084-5c52b6b3adef?w=400&h=300&fit=crop',
    category: 'Electronics',
    rating: 5,
    stock: 30,
    brand: 'Apple',
    discount: 5,
    isNew: false,
    isFeatured: true
  },
  {
    id: 3,
    name: 'AirPods Pro',
    description: 'Active noise cancellation with spatial audio',
    price: 8500000,
    image: 'https://images.unsplash.com/photo-1606220588913-b3aacb4d2f46?w=400&h=300&fit=crop',
    category: 'Electronics',
    rating: 4,
    stock: 100,
    brand: 'Apple',
    discount: 0,
    isNew: false,
    isFeatured: false
  },
  {
    id: 4,
    name: 'iPad Air',
    description: 'Powerful tablet with M1 chip and stunning display',
    price: 22000000,
    image: 'https://images.unsplash.com/photo-1544244015-0df4b3ffc6b0?w=400&h=300&fit=crop',
    category: 'Electronics',
    rating: 4,
    stock: 75,
    brand: 'Apple',
    discount: 15,
    isNew: true,
    isFeatured: false
  },
  {
    id: 5,
    name: 'Apple Watch Series 9',
    description: 'Advanced health monitoring and fitness tracking',
    price: 12000000,
    image: 'https://images.unsplash.com/photo-1523275335684-37898b6baf30?w=400&h=300&fit=crop',
    category: 'Electronics',
    rating: 5,
    stock: 60,
    brand: 'Apple',
    discount: 8,
    isNew: false,
    isFeatured: true
  },
  {
    id: 6,
    name: 'iMac 24"',
    description: 'All-in-one desktop with M1 chip and Retina display',
    price: 28000000,
    image: 'https://images.unsplash.com/photo-1517336714731-489689fd1ca8?w=400&h=300&fit=crop',
    category: 'Electronics',
    rating: 4,
    stock: 25,
    brand: 'Apple',
    discount: 12,
    isNew: false,
    isFeatured: false
  }
])

// Configuration state
const config = ref({
  slidesPerView: 3,
  autoplay: true,
  autoplayDelay: 5000,
  loop: true,
  showNavigation: true,
  showIndicators: true,
  showControls: true,
  carouselHeight: 400,
  slideGap: 20
})

// Event log
const eventLog = ref<Array<{type: string, message: string, timestamp: string}>>([])

// Methods
const logEvent = (type: string, message: string) => {
  const timestamp = new Date().toLocaleTimeString()
  eventLog.value.push({ type, message, timestamp })
  
  // Keep only last 50 events
  if (eventLog.value.length > 50) {
    eventLog.value = eventLog.value.slice(-50)
  }
}

import type { Product } from '~/types/product'

const handleProductClick = (product: Product) => {
  logEvent('product-click', `Clicked on product: ${product.name}`)
}

const handleSlideChange = (index: number) => {
  logEvent('slide-change', `Changed to slide ${index + 1}`)
}

const handleAutoplayToggle = (enabled: boolean) => {
  logEvent('autoplay-toggle', `Autoplay ${enabled ? 'enabled' : 'disabled'}`)
}

const clearEventLog = () => {
  eventLog.value = []
}

const getEventBorderClass = (type: string) => {
  switch (type) {
    case 'product-click':
      return 'border-blue-500'
    case 'slide-change':
      return 'border-green-500'
    case 'autoplay-toggle':
      return 'border-purple-500'
    default:
      return 'border-gray-300'
  }
}
</script>

<style scoped>
.product-carousel-demo-page {
  background-color: #f9fafb;
  min-height: 100vh;
}
</style>

