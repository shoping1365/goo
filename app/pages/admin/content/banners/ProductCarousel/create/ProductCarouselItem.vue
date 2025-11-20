<template>
  <div 
    class="product-carousel-item relative flex-shrink-0 cursor-pointer group"
    :style="{ width: `${slideWidth}px` }"
    @click="$emit('click')"
  >
    <!-- کانتینر اصلی آیتم -->
    <div class="product-item-container bg-white rounded-lg shadow-md overflow-hidden transition-all duration-300 group-hover:shadow-xl group-hover:scale-105">
      
      <!-- تصویر محصول -->
      <div class="product-image-container relative overflow-hidden">
        <img 
          :src="product.image" 
          :alt="product.name"
          class="w-full h-48 object-cover transition-transform duration-500 group-hover:scale-110"
          loading="lazy"
        />
        
        <!-- Overlay اطلاعات سریع -->
        <div class="product-quick-info absolute inset-0 bg-black bg-opacity-0 group-hover:bg-opacity-30 transition-all duration-300 flex items-center justify-center">
          <div class="text-white text-center opacity-0 group-hover:opacity-100 transition-opacity duration-300">
            <div class="text-sm font-medium mb-2">مشاهده جزئیات</div>
            <div class="text-xs">کلیک کنید</div>
          </div>
        </div>
        
        <!-- برچسب‌های محصول -->
        <div class="product-badges absolute top-2 right-2 flex flex-col gap-1">
          <!-- برچسب تخفیف -->
          <div 
            v-if="product.discount"
            class="discount-badge bg-red-500 text-white text-xs px-2 py-1 rounded-full font-bold"
          >
            {{ product.discount }}% تخفیف
          </div>
          
          <!-- برچسب جدید -->
          <div 
            v-if="product.isNew"
            class="new-badge bg-green-500 text-white text-xs px-2 py-1 rounded-full font-bold"
          >
            جدید
          </div>
          
          <!-- برچسب فروش ویژه -->
          <div 
            v-if="product.isSpecial"
            class="special-badge bg-purple-500 text-white text-xs px-2 py-1 rounded-full font-bold"
          >
            فروش ویژه
          </div>
        </div>
        
        <!-- دکمه‌های عملیات سریع -->
        <div class="quick-actions absolute top-2 left-2 flex flex-col gap-1 opacity-0 group-hover:opacity-100 transition-opacity duration-300">
          <!-- دکمه علاقه‌مندی -->
          <button 
            class="w-8 h-8 bg-white rounded-full shadow-md flex items-center justify-center hover:bg-red-50 transition-colors"
            :class="{ 'text-red-500': isFavorite, 'text-gray-400': !isFavorite }"
            @click.stop="toggleFavorite"
          >
            <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
              <path v-if="isFavorite" d="M3.172 5.172a4 4 0 015.656 0L10 6.343l1.172-1.171a4 4 0 115.656 5.656L10 17.657l-6.828-6.829a4 4 0 010-5.656z"/>
              <path v-else d="M3.172 5.172a4 4 0 015.656 0L10 6.343l1.172-1.171a4 4 0 115.656 5.656L10 17.657l-6.828-6.829a4 4 0 010-5.656z" fill="none" stroke="currentColor" stroke-width="2"/>
            </svg>
          </button>
          
          <!-- دکمه مقایسه -->
          <button 
            class="w-8 h-8 bg-white rounded-full shadow-md flex items-center justify-center hover:bg-blue-50 transition-colors text-gray-400 hover:text-blue-500"
            @click.stop="addToCompare"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"/>
            </svg>
          </button>
        </div>
      </div>
      
      <!-- اطلاعات محصول -->
      <div class="product-info p-6">
        <!-- دسته‌بندی -->
        <div class="product-category text-xs text-gray-500 mb-2">
          {{ product.category }}
        </div>
        
        <!-- عنوان محصول -->
        <h3 class="product-title text-sm font-semibold text-gray-800 mb-2 line-clamp-2 group-hover:text-blue-600 transition-colors">
          {{ product.name }}
        </h3>
        
        <!-- امتیاز و نظرات -->
        <div class="product-rating flex items-center gap-2 mb-3">
          <!-- ستاره‌ها -->
          <div class="stars flex items-center gap-1">
            <div 
              v-for="star in 5" 
              :key="star"
              class="star"
              :class="star <= product.rating ? 'text-yellow-400' : 'text-gray-300'"
            >
              <svg class="w-3 h-3" fill="currentColor" viewBox="0 0 20 20">
                <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"/>
              </svg>
            </div>
          </div>
          
          <!-- تعداد نظرات -->
          <span class="text-xs text-gray-500">({{ Math.floor(Math.random() * 50) + 10 }})</span>
        </div>
        
        <!-- قیمت -->
        <div class="product-price flex items-center gap-2 mb-3">
          <!-- قیمت اصلی -->
          <span class="current-price text-lg font-bold text-gray-800">
            {{ formatPrice(product.price) }}
          </span>
          
          <!-- قیمت قبلی (در صورت تخفیف) -->
          <span 
            v-if="product.discount"
            class="original-price text-sm text-gray-400 line-through"
          >
            {{ formatPrice(product.price / (1 - product.discount / 100)) }}
          </span>
        </div>
        
        <!-- دکمه افزودن به سبد خرید -->
        <button 
          class="add-to-cart-btn w-full bg-blue-600 text-white py-2 px-4 rounded-lg font-medium hover:bg-blue-700 transition-colors flex items-center justify-center gap-2"
          @click.stop="addToCart"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4m0 0L7 13m0 0l-2.5 5M7 13l2.5 5m6-5v6a2 2 0 01-2 2H9a2 2 0 01-2-2v-6m8 0V9a2 2 0 00-2-2H9a2 2 0 00-2 2v4.01M9 9h.01M15 9h.01"/>
          </svg>
          افزودن به سبد
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import type { Product } from '~/types/product'

// Props
interface Props {
  product: Product
  slideWidth: number
  isActive: boolean
}

const props = defineProps<Props>()

// Emits
const emit = defineEmits<{
  click: []
  'add-to-cart': [product: Product]
  'toggle-favorite': [product: Product, isFavorite: boolean]
  'add-to-compare': [product: Product]
}>()

// State
const isFavorite = ref(false)

// Computed
const hasDiscount = computed(() => {
  return props.product.discount && props.product.discount > 0
})

const discountPercentage = computed(() => {
  return props.product.discount || 0
})

// Methods
const formatPrice = (price: number): string => {
  return new Intl.NumberFormat('fa-IR').format(price) + ' تومان'
}

const addToCart = () => {
  emit('add-to-cart', props.product)
}

const toggleFavorite = () => {
  isFavorite.value = !isFavorite.value
  emit('toggle-favorite', props.product, isFavorite.value)
}

const addToCompare = () => {
  emit('add-to-compare', props.product)
}

// Initialize favorite state from product data
if (props.product.isFavorite !== undefined) {
  isFavorite.value = props.product.isFavorite
}
</script>

<style scoped>
.product-carousel-item {
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.product-item-container {
  border: 1px solid #e5e7eb;
  transition: all 0.3s ease;
}

.product-image-container {
  position: relative;
  overflow: hidden;
}

.product-quick-info {
  background: linear-gradient(to bottom, rgba(0,0,0,0) 0%, rgba(0,0,0,0.7) 100%);
}

.product-badges {
  z-index: 2;
}

.discount-badge {
  background: #ef4444;
  color: white;
  font-size: 10px;
  padding: 2px 6px;
  border-radius: 4px;
  font-weight: bold;
}

.new-badge {
  background: #10b981;
  color: white;
  font-size: 10px;
  padding: 2px 6px;
  border-radius: 4px;
  font-weight: bold;
}

.special-badge {
  background: #8b5cf6;
  color: white;
  font-size: 10px;
  padding: 2px 6px;
  border-radius: 4px;
  font-weight: bold;
}

.quick-actions {
  z-index: 3;
}

.product-info {
  background: white;
}

.product-category {
  color: #6b7280;
  font-size: 11px;
  font-weight: 500;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.product-title {
  color: #1f2937;
  font-weight: 600;
  line-height: 1.3;
  margin-bottom: 8px;
}

.product-rating {
  margin-bottom: 12px;
}

.stars {
  gap: 2px;
}

.star {
  color: #d1d5db;
  transition: color 0.2s ease;
}

.star.text-yellow-400 {
  color: #fbbf24;
}

.product-price {
  margin-bottom: 12px;
}

.current-price {
  color: #1f2937;
  font-weight: 700;
  font-size: 18px;
}

.original-price {
  color: #9ca3af;
  font-size: 14px;
  text-decoration: line-through;
}

.add-to-cart-btn {
  background: #3b82f6;
  border: none;
  border-radius: 6px;
  font-weight: 600;
  transition: all 0.2s ease;
  cursor: pointer;
}

.add-to-cart-btn:hover {
  background: #2563eb;
  transform: translateY(-1px);
}

/* Hover effects */
.group:hover .product-item-container {
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.1);
  transform: translateY(-2px);
}

.group:hover .product-quick-info {
  background: rgba(0, 0, 0, 0.3);
}

.group:hover .quick-actions {
  opacity: 1;
}

/* Responsive styles */
@media (max-width: 768px) {
  .product-carousel-item {
    margin: 0 8px;
  }

  .product-info {
    padding: 12px;
  }

  .current-price {
    font-size: 16px;
  }

  .product-title {
    font-size: 13px;
  }
}

@media (max-width: 480px) {
  .product-carousel-item {
    margin: 0 4px;
  }

  .product-info {
    padding: 8px;
  }

  .product-category {
    font-size: 10px;
  }

  .product-title {
    font-size: 12px;
    margin-bottom: 6px;
  }

  .current-price {
    font-size: 14px;
  }

  .add-to-cart-btn {
    padding: 8px 12px;
    font-size: 12px;
  }
}

/* Line clamp utility */
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

/* Accessibility */
.product-carousel-item:focus-within {
  outline: 2px solid #3b82f6;
  outline-offset: 2px;
}

button:focus {
  outline: 2px solid #3b82f6;
  outline-offset: 2px;
}
</style>
