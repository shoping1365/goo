<template>
  <div class="order-info-header bg-white rounded-xl shadow-lg border border-gray-200 p-6 mb-6">
    <div class="flex items-center justify-between mb-4">
      <div class="flex items-center space-x-3 space-x-reverse">
        <div class="p-2 bg-blue-100 rounded-lg">
          <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
          </svg>
        </div>
        <div>
          <h2 class="text-xl font-bold text-gray-900">اطلاعات سفارش</h2>
          <p class="text-sm text-gray-600">جزئیات سفارش شما</p>
        </div>
      </div>
      
      <div class="flex items-center space-x-2 space-x-reverse">
        <span class="px-3 py-1 bg-green-100 text-green-800 rounded-full text-sm font-medium">
          تحویل شده
        </span>
        <span class="text-sm text-gray-500">تاریخ تحویل: {{ orderInfo.deliveryDate }}</span>
      </div>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-6">
      <div class="bg-gray-50 rounded-lg p-6">
        <div class="flex items-center space-x-2 space-x-reverse mb-2">
          <svg class="w-5 h-5 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 20l4-16m2 16l4-16M6 9h14M4 15h14"></path>
          </svg>
          <span class="text-sm font-medium text-gray-700">شماره سفارش</span>
        </div>
        <p class="text-lg font-bold text-gray-900">{{ orderInfo.orderNumber }}</p>
      </div>

      <div class="bg-gray-50 rounded-lg p-6">
        <div class="flex items-center space-x-2 space-x-reverse mb-2">
          <svg class="w-5 h-5 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
          </svg>
          <span class="text-sm font-medium text-gray-700">تاریخ سفارش</span>
        </div>
        <p class="text-lg font-bold text-gray-900">{{ orderInfo.orderDate }}</p>
      </div>

      <div class="bg-gray-50 rounded-lg p-6">
        <div class="flex items-center space-x-2 space-x-reverse mb-2">
          <svg class="w-5 h-5 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"></path>
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"></path>
          </svg>
          <span class="text-sm font-medium text-gray-700">وضعیت تحویل</span>
        </div>
        <p class="text-lg font-bold text-green-600">{{ orderInfo.deliveryStatus }}</p>
      </div>

      <div class="bg-gray-50 rounded-lg p-6">
        <div class="flex items-center space-x-2 space-x-reverse mb-2">
          <svg class="w-5 h-5 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
          </svg>
          <span class="text-sm font-medium text-gray-700">مبلغ کل</span>
        </div>
        <p class="text-lg font-bold text-gray-900">{{ orderInfo.totalAmount }}</p>
      </div>
    </div>

    <!-- Products List -->
    <div class="border-t border-gray-200 pt-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">محصولات سفارش شده</h3>
      <div class="space-y-3">
        <div v-for="product in orderInfo.products" :key="product.id" class="flex items-center space-x-4 space-x-reverse bg-gray-50 rounded-lg p-6">
          <div class="flex-shrink-0">
            <img :src="product.image" :alt="product.name" class="w-16 h-16 object-cover rounded-lg">
          </div>
          <div class="flex-1 min-w-0">
            <h4 class="text-sm font-medium text-gray-900 truncate">{{ product.name }}</h4>
            <p class="text-sm text-gray-500">{{ product.sku }}</p>
            <div class="flex items-center space-x-4 space-x-reverse mt-1">
              <span class="text-sm text-gray-600">تعداد: {{ product.quantity }}</span>
              <span class="text-sm font-medium text-gray-900">{{ product.price }}</span>
            </div>
          </div>
          <div class="flex-shrink-0">
            <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800">
              {{ product.category }}
            </span>
          </div>
        </div>
      </div>
    </div>

    <!-- Delivery Information -->
    <div class="border-t border-gray-200 pt-4 mt-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">اطلاعات تحویل</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="bg-blue-50 rounded-lg p-6">
          <h4 class="font-medium text-blue-900 mb-2">آدرس تحویل</h4>
          <p class="text-sm text-blue-800">{{ orderInfo.deliveryAddress }}</p>
        </div>
        <div class="bg-green-50 rounded-lg p-6">
          <h4 class="font-medium text-green-900 mb-2">روش ارسال</h4>
          <p class="text-sm text-green-800">{{ orderInfo.shippingMethod }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

interface Product {
  id: number
  name: string
  sku: string
  image: string
  quantity: number
  price: string
  category: string
}

interface OrderInfo {
  orderNumber: string
  orderDate: string
  deliveryDate: string
  deliveryStatus: string
  totalAmount: string
  products: Product[]
  deliveryAddress: string
  shippingMethod: string
}

// Sample data - in real app this would come from props or API
const orderInfo = ref<OrderInfo>({
  orderNumber: 'ORD-2024-001234',
  orderDate: '۱۴۰۲/۱۰/۱۵',
  deliveryDate: '۱۴۰۲/۱۰/۱۸',
  deliveryStatus: 'تحویل شده',
  totalAmount: '۲,۵۰۰,۰۰۰ تومان',
  products: [
    {
      id: 1,
      name: 'لپ تاپ اپل مک‌بوک پرو ۱۳ اینچ',
      sku: 'MBP-13-2023',
      image: '/images/products/macbook-pro.jpg',
      quantity: 1,
      price: '۱,۸۰۰,۰۰۰ تومان',
      category: 'لپ تاپ'
    },
    {
      id: 2,
      name: 'کیف لپ تاپ چرمی برند سامسونیت',
      sku: 'BAG-SAM-001',
      image: '/images/products/laptop-bag.jpg',
      quantity: 1,
      price: '۴۵۰,۰۰۰ تومان',
      category: 'لوازم جانبی'
    },
    {
      id: 3,
      name: 'ماوس بی‌سیم اپل مجیک ماوس',
      sku: 'MAGIC-MOUSE-001',
      image: '/images/products/magic-mouse.jpg',
      quantity: 1,
      price: '۲۵۰,۰۰۰ تومان',
      category: 'لوازم جانبی'
    }
  ],
  deliveryAddress: 'تهران، خیابان ولیعصر، پلاک ۱۲۳، طبقه ۲، واحد ۵',
  shippingMethod: 'ارسال اکسپرس (۲۴ ساعته)'
})

// Expose data for parent component
defineExpose({
  orderInfo
})
</script>

<style scoped>
.order-info-header {
  transition: all 0.3s ease;
}

.order-info-header:hover {
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.1);
}

/* Custom animations */
@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.order-info-header > * {
  animation: fadeInUp 0.6s ease-out;
}

.order-info-header > *:nth-child(2) {
  animation-delay: 0s;
}

.order-info-header > *:nth-child(3) {
  animation-delay: 0s;
}
</style> 
