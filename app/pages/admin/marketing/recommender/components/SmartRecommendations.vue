<template>
  <div class="space-y-8">
    <!-- هدر بخش -->
    <div class="flex justify-between items-center">
      <div>
        <h2 class="text-2xl font-bold text-gray-900">پیشنهاد محصولات هوشمند</h2>
        <p class="mt-1 text-sm text-gray-600">مدیریت انواع پیشنهادات محصولات بر اساس الگوریتم‌های مختلف</p>
      </div>
      <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
        <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
        </svg>
        افزودن پیشنهاد جدید
      </button>
    </div>

    <!-- آمار پیشنهادات -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-6">
      <div class="bg-gradient-to-r from-blue-500 to-blue-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">تاریخچه خرید</p>
            <p class="text-2xl font-bold">1,247</p>
            <p class="text-xs opacity-75">پیشنهاد امروز</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-green-500 to-green-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">سبد خرید</p>
            <p class="text-2xl font-bold">892</p>
            <p class="text-xs opacity-75">پیشنهاد امروز</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4m0 0L7 13m0 0l-2.5 5M7 13l2.5 5m6-5v6a2 2 0 01-2 2H9a2 2 0 01-2-2v-6m8 0V9a2 2 0 00-2-2H9a2 2 0 00-2 2v4.01"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-purple-500 to-purple-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">مرور صفحات</p>
            <p class="text-2xl font-bold">2,156</p>
            <p class="text-xs opacity-75">پیشنهاد امروز</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-yellow-500 to-yellow-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">امتیازات</p>
            <p class="text-2xl font-bold">1,734</p>
            <p class="text-xs opacity-75">پیشنهاد امروز</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-red-500 to-red-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">زمان‌بندی</p>
            <p class="text-2xl font-bold">567</p>
            <p class="text-xs opacity-75">پیشنهاد امروز</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات پیشنهادات -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <!-- پیشنهاد بر اساس تاریخچه خرید -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-blue-100 rounded-lg">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">تاریخچه خرید</h3>
          </div>
          <label class="relative inline-flex items-center cursor-pointer">
            <input v-model="settings.purchaseHistory.enabled" type="checkbox" class="sr-only peer">
            <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
          </label>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تعداد پیشنهادات</label>
            <input v-model="settings.purchaseHistory.count" type="range" min="1" max="20" class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
            <div class="flex justify-between text-xs text-gray-500 mt-1">
              <span>1</span>
              <span>{{ settings.purchaseHistory.count }}</span>
              <span>20</span>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">وزن اهمیت</label>
            <div class="flex items-center space-x-4 space-x-reverse">
              <input v-model="settings.purchaseHistory.weight" type="range" min="0" max="100" class="flex-1 h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
              <span class="text-sm font-medium text-gray-900 w-12">{{ settings.purchaseHistory.weight }}%</span>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">بازه زمانی</label>
            <select v-model="settings.purchaseHistory.timeRange" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
              <option value="7">7 روز گذشته</option>
              <option value="30">30 روز گذشته</option>
              <option value="90">90 روز گذشته</option>
              <option value="365">1 سال گذشته</option>
            </select>
          </div>
        </div>
      </div>

      <!-- پیشنهاد بر اساس سبد خرید -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-green-100 rounded-lg">
              <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4m0 0L7 13m0 0l-2.5 5M7 13l2.5 5m6-5v6a2 2 0 01-2 2H9a2 2 0 01-2-2v-6m8 0V9a2 2 0 00-2-2H9a2 2 0 00-2 2v4.01"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">سبد خرید فعلی</h3>
          </div>
          <label class="relative inline-flex items-center cursor-pointer">
            <input v-model="settings.cartBased.enabled" type="checkbox" class="sr-only peer">
            <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-green-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-green-600"></div>
          </label>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع پیشنهاد</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input v-model="settings.cartBased.types" type="checkbox" value="complementary" class="rounded border-gray-300 text-green-600 focus:ring-green-500">
                <span class="mr-2 text-sm text-gray-700">محصولات تکمیلی</span>
              </label>
              <label class="flex items-center">
                <input v-model="settings.cartBased.types" type="checkbox" value="related" class="rounded border-gray-300 text-green-600 focus:ring-green-500">
                <span class="mr-2 text-sm text-gray-700">محصولات مرتبط</span>
              </label>
              <label class="flex items-center">
                <input v-model="settings.cartBased.types" type="checkbox" value="upsell" class="rounded border-gray-300 text-green-600 focus:ring-green-500">
                <span class="mr-2 text-sm text-gray-700">محصولات ارتقایی</span>
              </label>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداقل تعداد محصول</label>
            <input v-model="settings.cartBased.minItems" type="number" min="1" max="10" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">وزن اهمیت</label>
            <div class="flex items-center space-x-4 space-x-reverse">
              <input v-model="settings.cartBased.weight" type="range" min="0" max="100" class="flex-1 h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
              <span class="text-sm font-medium text-gray-900 w-12">{{ settings.cartBased.weight }}%</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- پیشنهادات فعال -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">پیشنهادات فعال</h3>
      </div>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع پیشنهاد</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نرخ کلیک</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نرخ تبدیل</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">آخرین به‌روزرسانی</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="recommendation in activeRecommendations" :key="recommendation.id">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="p-2 rounded-lg" :class="getTypeColor(recommendation.type).bg">
                    <svg class="w-5 h-5" :class="getTypeColor(recommendation.type).text" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="getTypeIcon(recommendation.type)"></path>
                    </svg>
                  </div>
                  <div class="mr-3">
                    <div class="text-sm font-medium text-gray-900">{{ recommendation.name }}</div>
                    <div class="text-sm text-gray-500">{{ recommendation.description }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full" :class="recommendation.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'">
                  {{ recommendation.status === 'active' ? 'فعال' : 'غیرفعال' }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ recommendation.clickRate }}%</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ recommendation.conversionRate }}%</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ recommendation.lastUpdated }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-4">ویرایش</button>
                <button class="text-red-600 hover:text-red-900">حذف</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
// تنظیمات پیشنهادات
const settings = ref({
  purchaseHistory: {
    enabled: true,
    count: 5,
    weight: 75,
    timeRange: '30'
  },
  cartBased: {
    enabled: true,
    types: ['complementary', 'related'],
    minItems: 2,
    weight: 60
  }
})

// پیشنهادات فعال
const activeRecommendations = ref([
  {
    id: 1,
    name: 'پیشنهاد بر اساس تاریخچه',
    description: 'محصولاتی که کاربران مشابه خریده‌اند',
    type: 'purchase-history',
    status: 'active',
    clickRate: 24.5,
    conversionRate: 18.7,
    lastUpdated: '2 ساعت پیش'
  },
  {
    id: 2,
    name: 'پیشنهاد سبد خرید',
    description: 'محصولات تکمیلی برای سبد خرید فعلی',
    type: 'cart-based',
    status: 'active',
    clickRate: 31.2,
    conversionRate: 22.1,
    lastUpdated: '1 ساعت پیش'
  },
  {
    id: 3,
    name: 'پیشنهاد مرور صفحات',
    description: 'محصولات مرتبط با صفحات مشاهده شده',
    type: 'page-view',
    status: 'active',
    clickRate: 19.8,
    conversionRate: 15.3,
    lastUpdated: '3 ساعت پیش'
  },
  {
    id: 4,
    name: 'پیشنهاد امتیازات',
    description: 'محصولات با امتیاز بالا در دسته‌بندی مشابه',
    type: 'rating-based',
    status: 'active',
    clickRate: 28.7,
    conversionRate: 20.4,
    lastUpdated: '4 ساعت پیش'
  }
])

// توابع کمکی
const getTypeColor = (type: string) => {
  const colors = {
    'purchase-history': { bg: 'bg-blue-100', text: 'text-blue-600' },
    'cart-based': { bg: 'bg-green-100', text: 'text-green-600' },
    'page-view': { bg: 'bg-purple-100', text: 'text-purple-600' },
    'rating-based': { bg: 'bg-yellow-100', text: 'text-yellow-600' },
    'time-based': { bg: 'bg-red-100', text: 'text-red-600' }
  }
  return colors[type] || colors['purchase-history']
}

const getTypeIcon = (type: string) => {
  const icons = {
    'purchase-history': 'M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2',
    'cart-based': 'M3 3h2l.4 2M7 13h10l4-8H5.4m0 0L7 13m0 0l-2.5 5M7 13l2.5 5m6-5v6a2 2 0 01-2 2H9a2 2 0 01-2-2v-6m8 0V9a2 2 0 00-2-2H9a2 2 0 00-2 2v4.01',
    'page-view': 'M15 12a3 3 0 11-6 0 3 3 0 016 0z M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z',
    'rating-based': 'M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z',
    'time-based': 'M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z'
  }
  return icons[type] || icons['purchase-history']
}
</script> 
