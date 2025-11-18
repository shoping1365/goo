<template>
  <div class="space-y-8">
    <!-- هدر بخش -->
    <div class="flex justify-between items-center">
      <div>
        <h2 class="text-2xl font-bold text-gray-900">تحلیل رفتار کاربران</h2>
        <p class="mt-1 text-sm text-gray-600">تحلیل و بررسی رفتار کاربران برای بهبود پیشنهادات</p>
      </div>
      <button class="bg-purple-600 text-white px-4 py-2 rounded-lg hover:bg-purple-700 transition-colors">
        <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
        </svg>
        تولید گزارش جدید
      </button>
    </div>

    <!-- آمار کلی رفتار -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="p-3 bg-blue-100 rounded-lg">
            <svg class="w-8 h-8 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
            </svg>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">کاربران فعال</p>
            <p class="text-2xl font-semibold text-gray-900">12,847</p>
            <p class="text-sm text-green-600">+5.2% از هفته گذشته</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="p-3 bg-green-100 rounded-lg">
            <svg class="w-8 h-8 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">میانگین زمان جلسه</p>
            <p class="text-2xl font-semibold text-gray-900">8.5 دقیقه</p>
            <p class="text-sm text-green-600">+1.2 دقیقه از هفته گذشته</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="p-3 bg-purple-100 rounded-lg">
            <svg class="w-8 h-8 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">صفحات مشاهده شده</p>
            <p class="text-2xl font-semibold text-gray-900">45,234</p>
            <p class="text-sm text-green-600">+12.3% از هفته گذشته</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="p-3 bg-yellow-100 rounded-lg">
            <svg class="w-8 h-8 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
            </svg>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">نرخ تبدیل</p>
            <p class="text-2xl font-semibold text-gray-900">3.8%</p>
            <p class="text-sm text-green-600">+0.5% از هفته گذشته</p>
          </div>
        </div>
      </div>
    </div>

    <!-- نمودارهای تحلیل رفتار -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <!-- الگوی خرید -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <h3 class="text-lg font-semibold text-gray-900">الگوی خرید</h3>
          <select class="text-sm border border-gray-300 rounded-lg px-3 py-1">
            <option>7 روز گذشته</option>
            <option>30 روز گذشته</option>
            <option>90 روز گذشته</option>
          </select>
        </div>
        
        <div class="space-y-4">
          <div v-for="pattern in purchasePatterns" :key="pattern.id" class="flex items-center justify-between p-6 bg-gray-50 rounded-lg">
            <div class="flex items-center">
              <div class="w-3 h-3 rounded-full mr-3" :class="pattern.color"></div>
              <div>
                <p class="text-sm font-medium text-gray-900">{{ pattern.name }}</p>
                <p class="text-xs text-gray-500">{{ pattern.description }}</p>
              </div>
            </div>
            <div class="text-right">
              <p class="text-sm font-semibold text-gray-900">{{ pattern.percentage }}%</p>
              <p class="text-xs text-gray-500">{{ pattern.count }} کاربر</p>
            </div>
          </div>
        </div>
      </div>

      <!-- زمان‌بندی خرید -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <h3 class="text-lg font-semibold text-gray-900">زمان‌بندی خرید</h3>
          <div class="flex space-x-2 space-x-reverse">
            <button class="px-3 py-1 text-xs bg-blue-100 text-blue-800 rounded-lg">روزانه</button>
            <button class="px-3 py-1 text-xs bg-gray-100 text-gray-600 rounded-lg">هفتگی</button>
          </div>
        </div>
        
        <div class="space-y-4">
          <div v-for="timeSlot in timeSlots" :key="timeSlot.id" class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-8 text-sm text-gray-600">{{ timeSlot.time }}</div>
              <div class="mr-4">
                <div class="w-32 bg-gray-200 rounded-full h-2">
                  <div class="bg-blue-600 h-2 rounded-full" :style="{ width: timeSlot.percentage + '%' }"></div>
                </div>
              </div>
            </div>
            <div class="text-right">
              <p class="text-sm font-medium text-gray-900">{{ timeSlot.percentage }}%</p>
              <p class="text-xs text-gray-500">{{ timeSlot.count }} خرید</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تحلیل دستگاه و مرورگر -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <!-- نوع دستگاه -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-6">نوع دستگاه</h3>
        <div class="space-y-4">
          <div v-for="device in devices" :key="device.id" class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="p-2 rounded-lg mr-3" :class="device.bgColor">
                <svg class="w-5 h-5" :class="device.textColor" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="device.icon"></path>
                </svg>
              </div>
              <div>
                <p class="text-sm font-medium text-gray-900">{{ device.name }}</p>
                <p class="text-xs text-gray-500">{{ device.description }}</p>
              </div>
            </div>
            <div class="text-right">
              <p class="text-sm font-semibold text-gray-900">{{ device.percentage }}%</p>
              <p class="text-xs text-gray-500">{{ device.count }} کاربر</p>
            </div>
          </div>
        </div>
      </div>

      <!-- مرورگرهای محبوب -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-6">مرورگرهای محبوب</h3>
        <div class="space-y-4">
          <div v-for="browser in browsers" :key="browser.id" class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-8 h-8 rounded-lg mr-3 flex items-center justify-center" :class="browser.bgColor">
                <span class="text-white text-xs font-bold">{{ browser.icon }}</span>
              </div>
              <div>
                <p class="text-sm font-medium text-gray-900">{{ browser.name }}</p>
                <p class="text-xs text-gray-500">نسخه {{ browser.version }}</p>
              </div>
            </div>
            <div class="text-right">
              <p class="text-sm font-semibold text-gray-900">{{ browser.percentage }}%</p>
              <p class="text-xs text-gray-500">{{ browser.count }} کاربر</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تحلیل منبع ترافیک -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-lg font-semibold text-gray-900">تحلیل منبع ترافیک</h3>
        <button class="text-blue-600 hover:text-blue-800 text-sm font-medium">مشاهده جزئیات</button>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <div v-for="source in trafficSources" :key="source.id" class="text-center">
          <div class="w-16 h-16 mx-auto mb-4 rounded-full flex items-center justify-center" :class="source.bgColor">
            <svg class="w-8 h-8" :class="source.textColor" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="source.icon"></path>
            </svg>
          </div>
          <h4 class="text-sm font-medium text-gray-900 mb-1">{{ source.name }}</h4>
          <p class="text-2xl font-bold text-gray-900 mb-1">{{ source.percentage }}%</p>
          <p class="text-xs text-gray-500">{{ source.count }} کاربر</p>
          <div class="mt-2">
            <span class="text-xs" :class="source.trend > 0 ? 'text-green-600' : 'text-red-600'">
              {{ source.trend > 0 ? '+' : '' }}{{ source.trend }}%
            </span>
            <span class="text-xs text-gray-500">از هفته گذشته</span>
          </div>
        </div>
      </div>
    </div>

    <!-- تحلیل جغرافیایی -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-lg font-semibold text-gray-900">تحلیل جغرافیایی</h3>
        <div class="flex space-x-2 space-x-reverse">
          <button class="px-3 py-1 text-xs bg-blue-100 text-blue-800 rounded-lg">استان</button>
          <button class="px-3 py-1 text-xs bg-gray-100 text-gray-600 rounded-lg">شهر</button>
        </div>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div v-for="region in topRegions" :key="region.id" class="border border-gray-200 rounded-lg p-6">
          <div class="flex items-center justify-between mb-3">
            <div class="flex items-center">
              <div class="w-8 h-8 bg-blue-100 rounded-lg flex items-center justify-center mr-3">
                <span class="text-blue-600 text-sm font-bold">{{ region.rank }}</span>
              </div>
              <div>
                <p class="text-sm font-medium text-gray-900">{{ region.name }}</p>
                <p class="text-xs text-gray-500">{{ region.province }}</p>
              </div>
            </div>
            <div class="text-right">
              <p class="text-sm font-semibold text-gray-900">{{ region.percentage }}%</p>
              <p class="text-xs text-gray-500">{{ region.count }} کاربر</p>
            </div>
          </div>
          <div class="w-full bg-gray-200 rounded-full h-2">
            <div class="bg-blue-600 h-2 rounded-full" :style="{ width: region.percentage + '%' }"></div>
          </div>
        </div>
      </div>
    </div>

    <!-- رفتارهای غیرعادی -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-lg font-semibold text-gray-900">رفتارهای غیرعادی</h3>
        <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-red-100 text-red-800">
          {{ abnormalBehaviors.length }} مورد جدید
        </span>
      </div>
      
      <div class="space-y-4">
        <div v-for="behavior in abnormalBehaviors" :key="behavior.id" class="flex items-center justify-between p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center">
            <div class="p-2 rounded-lg mr-4" :class="behavior.severity === 'high' ? 'bg-red-100' : behavior.severity === 'medium' ? 'bg-yellow-100' : 'bg-blue-100'">
              <svg class="w-5 h-5" :class="behavior.severity === 'high' ? 'text-red-600' : behavior.severity === 'medium' ? 'text-yellow-600' : 'text-blue-600'" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
              </svg>
            </div>
            <div>
              <p class="text-sm font-medium text-gray-900">{{ behavior.title }}</p>
              <p class="text-xs text-gray-500">{{ behavior.description }}</p>
              <p class="text-xs text-gray-400 mt-1">{{ behavior.time }}</p>
            </div>
          </div>
          <div class="flex space-x-2 space-x-reverse">
            <button class="text-blue-600 hover:text-blue-800 text-sm">بررسی</button>
            <button class="text-gray-600 hover:text-gray-800 text-sm">نادیده گرفتن</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
// الگوهای خرید
const purchasePatterns = ref([
  {
    id: 1,
    name: 'خریداران منظم',
    description: 'کاربرانی که هر هفته خرید می‌کنند',
    percentage: 35,
    count: 4496,
    color: 'bg-blue-500'
  },
  {
    id: 2,
    name: 'خریداران فصلی',
    description: 'کاربرانی که در مناسبت‌ها خرید می‌کنند',
    percentage: 28,
    count: 3597,
    color: 'bg-green-500'
  },
  {
    id: 3,
    name: 'خریداران تصادفی',
    description: 'کاربرانی که گاه‌به‌گاه خرید می‌کنند',
    percentage: 22,
    count: 2826,
    color: 'bg-purple-500'
  },
  {
    id: 4,
    name: 'خریداران جدید',
    description: 'کاربرانی که برای اولین بار خرید کرده‌اند',
    percentage: 15,
    count: 1928,
    color: 'bg-yellow-500'
  }
])

// زمان‌بندی خرید
const timeSlots = ref([
  { id: 1, time: '09:00-12:00', percentage: 15, count: 1927 },
  { id: 2, time: '12:00-15:00', percentage: 25, count: 3212 },
  { id: 3, time: '15:00-18:00', percentage: 30, count: 3854 },
  { id: 4, time: '18:00-21:00', percentage: 20, count: 2568 },
  { id: 5, time: '21:00-24:00', percentage: 10, count: 1284 }
])

// نوع دستگاه
const devices = ref([
  {
    id: 1,
    name: 'موبایل',
    description: 'گوشی‌های هوشمند',
    percentage: 65,
    count: 8349,
    bgColor: 'bg-blue-100',
    textColor: 'text-blue-600',
    icon: 'M12 18h.01M8 21h8a2 2 0 002-2V5a2 2 0 00-2-2H8a2 2 0 00-2 2v14a2 2 0 002 2z'
  },
  {
    id: 2,
    name: 'دسکتاپ',
    description: 'کامپیوترهای شخصی',
    percentage: 25,
    count: 3212,
    bgColor: 'bg-green-100',
    textColor: 'text-green-600',
    icon: 'M9.75 17L9 20l-1 1h8l-1-1-.75-3M3 13h18M5 17h14a2 2 0 002-2V5a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z'
  },
  {
    id: 3,
    name: 'تبلت',
    description: 'تبلت‌ها و آیپد',
    percentage: 10,
    count: 1286,
    bgColor: 'bg-purple-100',
    textColor: 'text-purple-600',
    icon: 'M12 18h.01M7 21h10a2 2 0 002-2V5a2 2 0 00-2-2H7a2 2 0 00-2 2v14a2 2 0 002 2z'
  }
])

// مرورگرهای محبوب
const browsers = ref([
  {
    id: 1,
    name: 'Chrome',
    version: '120.0',
    percentage: 45,
    count: 5781,
    bgColor: 'bg-blue-500',
    icon: 'C'
  },
  {
    id: 2,
    name: 'Safari',
    version: '17.0',
    percentage: 25,
    count: 3212,
    bgColor: 'bg-gray-500',
    icon: 'S'
  },
  {
    id: 3,
    name: 'Firefox',
    version: '121.0',
    percentage: 15,
    count: 1927,
    bgColor: 'bg-orange-500',
    icon: 'F'
  },
  {
    id: 4,
    name: 'Edge',
    version: '120.0',
    percentage: 10,
    count: 1284,
    bgColor: 'bg-blue-600',
    icon: 'E'
  },
  {
    id: 5,
    name: 'سایر',
    version: '-',
    percentage: 5,
    count: 642,
    bgColor: 'bg-gray-400',
    icon: 'O'
  }
])

// منابع ترافیک
const trafficSources = ref([
  {
    id: 1,
    name: 'جستجوی ارگانیک',
    percentage: 40,
    count: 5139,
    trend: 5.2,
    bgColor: 'bg-green-500',
    textColor: 'text-white',
    icon: 'M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z'
  },
  {
    id: 2,
    name: 'شبکه‌های اجتماعی',
    percentage: 25,
    count: 3212,
    trend: 12.8,
    bgColor: 'bg-blue-500',
    textColor: 'text-white',
    icon: 'M7 4V2a1 1 0 011-1h8a1 1 0 011 1v2m-9 0h10m-10 0a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V6a2 2 0 00-2-2'
  },
  {
    id: 3,
    name: 'ایمیل',
    percentage: 20,
    count: 2568,
    trend: -2.1,
    bgColor: 'bg-purple-500',
    textColor: 'text-white',
    icon: 'M3 8l7.89 4.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z'
  },
  {
    id: 4,
    name: 'مستقیم',
    percentage: 15,
    count: 1927,
    trend: 1.5,
    bgColor: 'bg-gray-500',
    textColor: 'text-white',
    icon: 'M13 7l5 5m0 0l-5 5m5-5H6'
  }
])

// مناطق برتر
const topRegions = ref([
  { id: 1, name: 'تهران', province: 'تهران', rank: 1, percentage: 35, count: 4496 },
  { id: 2, name: 'اصفهان', province: 'اصفهان', rank: 2, percentage: 18, count: 2312 },
  { id: 3, name: 'مشهد', province: 'خراسان رضوی', rank: 3, percentage: 15, count: 1927 },
  { id: 4, name: 'شیراز', province: 'فارس', rank: 4, percentage: 12, count: 1542 },
  { id: 5, name: 'تبریز', province: 'آذربایجان شرقی', rank: 5, percentage: 10, count: 1284 },
  { id: 6, name: 'قم', province: 'قم', rank: 6, percentage: 10, count: 1284 }
])

// رفتارهای غیرعادی
const abnormalBehaviors = ref([
  {
    id: 1,
    title: 'کلیک‌های مشکوک',
    description: 'کاربری در 5 دقیقه گذشته 50 بار روی محصولات کلیک کرده است',
    severity: 'high',
    time: '2 دقیقه پیش'
  },
  {
    id: 2,
    title: 'خریدهای غیرعادی',
    description: 'کاربری در یک ساعت 15 محصول مختلف خریده است',
    severity: 'medium',
    time: '15 دقیقه پیش'
  },
  {
    id: 3,
    title: 'جستجوهای مکرر',
    description: 'کاربری 20 بار کلمه کلیدی مشابه را جستجو کرده است',
    severity: 'low',
    time: '1 ساعت پیش'
  }
])
</script> 
