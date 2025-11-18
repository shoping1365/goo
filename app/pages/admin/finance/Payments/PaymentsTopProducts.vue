<template>
  <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-6">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">پرفروش‌ترین محصولات</h3>
        <p class="text-sm text-gray-500 mt-1">Top Selling Products</p>
      </div>
      <div class="flex items-center space-x-2 space-x-reverse">
        <button 
          @click="changePeriod('week')"
          :class="[
            'px-3 py-1 text-sm rounded-lg transition-colors',
            period === 'week' ? 'bg-blue-500 text-white' : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
          ]"
        >
          هفته
        </button>
        <button 
          @click="changePeriod('month')"
          :class="[
            'px-3 py-1 text-sm rounded-lg transition-colors',
            period === 'month' ? 'bg-blue-500 text-white' : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
          ]"
        >
          ماه
        </button>
        <button 
          @click="changePeriod('year')"
          :class="[
            'px-3 py-1 text-sm rounded-lg transition-colors',
            period === 'year' ? 'bg-blue-500 text-white' : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
          ]"
        >
          سال
        </button>
      </div>
    </div>

    <div class="space-y-4">
      <!-- محصول اول -->
      <div class="flex items-center justify-between p-6 bg-gradient-to-r from-yellow-50 to-orange-50 rounded-lg border border-yellow-200">
        <div class="flex items-center">
          <div class="w-12 h-12 bg-yellow-500 rounded-xl flex items-center justify-center ml-3">
            <span class="text-white font-bold text-lg">۱</span>
          </div>
          <div>
            <h4 class="font-medium text-gray-900">گوشی هوشمند سامسونگ Galaxy S24</h4>
            <p class="text-sm text-gray-500">کد محصول: PRD-001</p>
          </div>
        </div>
        <div class="text-left">
          <div class="text-lg font-bold text-gray-900">{{ formatPrice(topProducts[0].revenue) }}</div>
          <div class="text-sm text-gray-600">{{ topProducts[0].sales }} فروش</div>
        </div>
      </div>

      <!-- محصول دوم -->
      <div class="flex items-center justify-between p-6 bg-gradient-to-r from-gray-50 to-blue-50 rounded-lg border border-gray-200">
        <div class="flex items-center">
          <div class="w-12 h-12 bg-gray-500 rounded-xl flex items-center justify-center ml-3">
            <span class="text-white font-bold text-lg">۲</span>
          </div>
          <div>
            <h4 class="font-medium text-gray-900">لپ‌تاپ اپل MacBook Pro 14</h4>
            <p class="text-sm text-gray-500">کد محصول: PRD-002</p>
          </div>
        </div>
        <div class="text-left">
          <div class="text-lg font-bold text-gray-900">{{ formatPrice(topProducts[1].revenue) }}</div>
          <div class="text-sm text-gray-600">{{ topProducts[1].sales }} فروش</div>
        </div>
      </div>

      <!-- محصول سوم -->
      <div class="flex items-center justify-between p-6 bg-gradient-to-r from-orange-50 to-red-50 rounded-lg border border-orange-200">
        <div class="flex items-center">
          <div class="w-12 h-12 bg-orange-500 rounded-xl flex items-center justify-center ml-3">
            <span class="text-white font-bold text-lg">۳</span>
          </div>
          <div>
            <h4 class="font-medium text-gray-900">ساعت هوشمند اپل Watch Series 9</h4>
            <p class="text-sm text-gray-500">کد محصول: PRD-003</p>
          </div>
        </div>
        <div class="text-left">
          <div class="text-lg font-bold text-gray-900">{{ formatPrice(topProducts[2].revenue) }}</div>
          <div class="text-sm text-gray-600">{{ topProducts[2].sales }} فروش</div>
        </div>
      </div>

      <!-- محصول چهارم -->
      <div class="flex items-center justify-between p-6 bg-gradient-to-r from-green-50 to-emerald-50 rounded-lg border border-green-200">
        <div class="flex items-center">
          <div class="w-12 h-12 bg-green-500 rounded-xl flex items-center justify-center ml-3">
            <span class="text-white font-bold text-lg">۴</span>
          </div>
          <div>
            <h4 class="font-medium text-gray-900">هدفون Sony WH-1000XM5</h4>
            <p class="text-sm text-gray-500">کد محصول: PRD-004</p>
          </div>
        </div>
        <div class="text-left">
          <div class="text-lg font-bold text-gray-900">{{ formatPrice(topProducts[3].revenue) }}</div>
          <div class="text-sm text-gray-600">{{ topProducts[3].sales }} فروش</div>
        </div>
      </div>

      <!-- محصول پنجم -->
      <div class="flex items-center justify-between p-6 bg-gradient-to-r from-purple-50 to-indigo-50 rounded-lg border border-purple-200">
        <div class="flex items-center">
          <div class="w-12 h-12 bg-purple-500 rounded-xl flex items-center justify-center ml-3">
            <span class="text-white font-bold text-lg">۵</span>
          </div>
          <div>
            <h4 class="font-medium text-gray-900">تبلت اپل iPad Pro 12.9</h4>
            <p class="text-sm text-gray-500">کد محصول: PRD-005</p>
          </div>
        </div>
        <div class="text-left">
          <div class="text-lg font-bold text-gray-900">{{ formatPrice(topProducts[4].revenue) }}</div>
          <div class="text-sm text-gray-600">{{ topProducts[4].sales }} فروش</div>
        </div>
      </div>
    </div>

    <!-- آمار خلاصه -->
    <div class="mt-6 pt-4 border-t border-gray-200">
      <h4 class="font-medium text-gray-900 mb-3">آمار کلی</h4>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div class="text-center p-6 bg-blue-50 rounded-lg">
          <div class="text-2xl font-bold text-blue-600">{{ formatPrice(totalRevenue) }}</div>
          <div class="text-sm text-blue-500">کل درآمد</div>
        </div>
        <div class="text-center p-6 bg-green-50 rounded-lg">
          <div class="text-2xl font-bold text-green-600">{{ totalSales }}</div>
          <div class="text-sm text-green-500">کل فروش</div>
        </div>
        <div class="text-center p-6 bg-purple-50 rounded-lg">
          <div class="text-2xl font-bold text-purple-600">{{ formatPrice(averageRevenue) }}</div>
          <div class="text-sm text-purple-500">میانگین درآمد</div>
        </div>
      </div>
    </div>

    <!-- نمودار توزیع فروش -->
    <div class="mt-6 pt-4 border-t border-gray-200">
      <h4 class="font-medium text-gray-900 mb-3">توزیع فروش بر اساس محصول</h4>
      <div class="h-32 bg-gray-50 rounded-lg flex items-center justify-center">
        <div class="text-center">
          <svg class="w-12 h-12 text-gray-300 mx-auto mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 3.055A9.001 9.001 0 1020.945 13H11V3.055z"></path>
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.488 9H15V3.512A9.025 9.025 0 0120.488 9z"></path>
          </svg>
          <p class="text-sm text-gray-500">نمودار توزیع فروش محصولات</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

const period = ref('month')

// داده‌های پرفروش‌ترین محصولات بر اساس دوره
const topProducts = computed(() => {
  if (period.value === 'week') {
    return [
      { name: 'گوشی هوشمند سامسونگ Galaxy S24', code: 'PRD-001', revenue: 12500000, sales: 25 },
      { name: 'لپ‌تاپ اپل MacBook Pro 14', code: 'PRD-002', revenue: 8900000, sales: 8 },
      { name: 'ساعت هوشمند اپل Watch Series 9', code: 'PRD-003', revenue: 6700000, sales: 45 },
      { name: 'هدفون Sony WH-1000XM5', code: 'PRD-004', revenue: 4500000, sales: 30 },
      { name: 'تبلت اپل iPad Pro 12.9', code: 'PRD-005', revenue: 3800000, sales: 6 }
    ]
  } else if (period.value === 'month') {
    return [
      { name: 'گوشی هوشمند سامسونگ Galaxy S24', code: 'PRD-001', revenue: 45600000, sales: 89 },
      { name: 'لپ‌تاپ اپل MacBook Pro 14', code: 'PRD-002', revenue: 32400000, sales: 28 },
      { name: 'ساعت هوشمند اپل Watch Series 9', code: 'PRD-003', revenue: 23400000, sales: 156 },
      { name: 'هدفون Sony WH-1000XM5', code: 'PRD-004', revenue: 16700000, sales: 112 },
      { name: 'تبلت اپل iPad Pro 12.9', code: 'PRD-005', revenue: 14500000, sales: 23 }
    ]
  } else {
    return [
      { name: 'گوشی هوشمند سامسونگ Galaxy S24', code: 'PRD-001', revenue: 234000000, sales: 456 },
      { name: 'لپ‌تاپ اپل MacBook Pro 14', code: 'PRD-002', revenue: 189000000, sales: 167 },
      { name: 'ساعت هوشمند اپل Watch Series 9', code: 'PRD-003', revenue: 145000000, sales: 789 },
      { name: 'هدفون Sony WH-1000XM5', code: 'PRD-004', revenue: 123000000, sales: 567 },
      { name: 'تبلت اپل iPad Pro 12.9', code: 'PRD-005', revenue: 98000000, sales: 145 }
    ]
  }
})

// محاسبه آمار
const totalRevenue = computed(() => {
  return topProducts.value.reduce((sum, product) => sum + product.revenue, 0)
})

const totalSales = computed(() => {
  return topProducts.value.reduce((sum, product) => sum + product.sales, 0)
})

const averageRevenue = computed(() => {
  return totalRevenue.value / topProducts.value.length
})

const changePeriod = (newPeriod: string) => {
  period.value = newPeriod
}

const formatPrice = (price: number) => {
  if (price >= 1000000) {
    return (price / 1000000).toFixed(1) + 'M'
  } else if (price >= 1000) {
    return (price / 1000).toFixed(0) + 'K'
  }
  return new Intl.NumberFormat('fa-IR').format(price)
}
</script>

<!--
  کامپوننت پرفروش‌ترین محصولات
  - نمایش ۵ محصول برتر
  - فیلتر بر اساس دوره (هفته/ماه/سال)
  - آمار خلاصه
  - نمودار توزیع فروش
  - کاملاً ریسپانسیو
  توضیحات کامل به فارسی در کد
--> 
