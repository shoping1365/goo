<template>
  <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-6">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">روش‌های پرداخت</h3>
        <p class="text-sm text-gray-500 mt-1">Payment Methods Statistics</p>
      </div>
      <div class="flex items-center space-x-2 space-x-reverse">
        <button class="text-sm text-blue-600 hover:text-blue-800">مشاهده جزئیات</button>
      </div>
    </div>

    <!-- آمار روش‌های پرداخت -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div 
        v-for="method in paymentMethods" 
        :key="method.id"
        class="bg-gray-50 rounded-lg p-6 hover:bg-gray-100 transition-colors cursor-pointer"
      >
        <div class="flex items-center justify-between mb-3">
          <div class="flex items-center">
            <div class="w-8 h-8 rounded-lg bg-blue-100 flex items-center justify-center ml-3">
              <svg class="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"></path>
              </svg>
            </div>
            <div>
              <div class="font-medium text-gray-900">{{ method.name }}</div>
              <div class="text-sm text-gray-500">{{ method.gateway }}</div>
            </div>
          </div>
          <div class="text-right">
            <div class="text-lg font-bold text-gray-900">{{ formatPrice(method.totalAmount) }}</div>
            <div class="text-xs text-gray-500">{{ method.count }} تراکنش</div>
          </div>
        </div>
        
        <!-- نوار پیشرفت -->
        <div class="w-full bg-gray-200 rounded-full h-2">
          <div 
            class="bg-blue-600 h-2 rounded-full transition-all duration-300"
            :style="{ width: method.percentage + '%' }"
          ></div>
        </div>
        
        <div class="flex items-center justify-between mt-2">
          <div class="text-xs text-gray-500">{{ method.percentage }}% از کل</div>
          <div class="text-xs text-green-600">{{ method.successRate }}% موفقیت</div>
        </div>
      </div>
    </div>

    <!-- خلاصه آمار -->
    <div class="mt-6 pt-6 border-t border-gray-200">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div class="text-center">
          <div class="text-2xl font-bold text-blue-600">{{ totalMethods }}</div>
          <div class="text-sm text-gray-500">روش پرداخت فعال</div>
        </div>
        <div class="text-center">
          <div class="text-2xl font-bold text-green-600">{{ formatPrice(totalAmount) }}</div>
          <div class="text-sm text-gray-500">مجموع مبالغ</div>
        </div>
        <div class="text-center">
          <div class="text-2xl font-bold text-purple-600">{{ totalTransactions }}</div>
          <div class="text-sm text-gray-500">کل تراکنش‌ها</div>
        </div>
        <div class="text-center">
          <div class="text-2xl font-bold text-orange-600">{{ averageSuccessRate }}%</div>
          <div class="text-sm text-gray-500">میانگین موفقیت</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

// داده‌های روش‌های پرداخت
const paymentMethods = ref([
  {
    id: 1,
    name: 'زرین‌پال',
    gateway: 'ZarinPal',
    totalAmount: 45000000,
    count: 125,
    percentage: 45,
    successRate: 92
  },
  {
    id: 2,
    name: 'نکست‌پی',
    gateway: 'NextPay',
    totalAmount: 32000000,
    count: 89,
    percentage: 32,
    successRate: 88
  },
  {
    id: 3,
    name: 'کارت به کارت',
    gateway: 'Card to Card',
    totalAmount: 18000000,
    count: 45,
    percentage: 18,
    successRate: 95
  },
  {
    id: 4,
    name: 'کیف پول',
    gateway: 'Wallet',
    totalAmount: 8000000,
    count: 23,
    percentage: 8,
    successRate: 100
  },
  {
    id: 5,
    name: 'کارت اعتباری',
    gateway: 'Credit Card',
    totalAmount: 12000000,
    count: 34,
    percentage: 12,
    successRate: 85
  },
  {
    id: 6,
    name: 'پرداخت نقدی',
    gateway: 'Cash on Delivery',
    totalAmount: 5000000,
    count: 15,
    percentage: 5,
    successRate: 100
  }
])

// محاسبه آمار کلی
const totalMethods = computed(() => paymentMethods.value.length)

const totalAmount = computed(() => {
  return paymentMethods.value.reduce((sum, method) => sum + method.totalAmount, 0)
})

const totalTransactions = computed(() => {
  return paymentMethods.value.reduce((sum, method) => sum + method.count, 0)
})

const averageSuccessRate = computed(() => {
  const total = paymentMethods.value.reduce((sum, method) => sum + method.successRate, 0)
  return Math.round(total / paymentMethods.value.length)
})

// فرمت کردن قیمت
const formatPrice = (price: number) => {
  if (price >= 1000000) {
    return (price / 1000000).toFixed(1) + 'M'
  } else if (price >= 1000) {
    return (price / 1000).toFixed(0) + 'K'
  }
  return price.toString()
}
</script> 
