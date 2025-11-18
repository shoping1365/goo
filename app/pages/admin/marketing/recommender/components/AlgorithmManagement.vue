<template>
  <div class="space-y-8">
    <!-- هدر بخش -->
    <div class="flex justify-between items-center">
      <div>
        <h2 class="text-2xl font-bold text-gray-900">مدیریت الگوریتم‌ها</h2>
        <p class="mt-1 text-sm text-gray-600">مدیریت و تنظیم الگوریتم‌های پیشنهادات هوشمند</p>
      </div>
      <button class="bg-indigo-600 text-white px-4 py-2 rounded-lg hover:bg-indigo-700 transition-colors">
        <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
        </svg>
        افزودن الگوریتم جدید
      </button>
    </div>

    <!-- آمار الگوریتم‌ها -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <div class="bg-gradient-to-r from-indigo-500 to-indigo-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">الگوریتم‌های فعال</p>
            <p class="text-2xl font-bold">8</p>
            <p class="text-xs opacity-75">از 12 الگوریتم</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-green-500 to-green-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">میانگین دقت</p>
            <p class="text-2xl font-bold">87.3%</p>
            <p class="text-xs opacity-75">+2.1% از ماه گذشته</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-purple-500 to-purple-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">زمان پاسخ</p>
            <p class="text-2xl font-bold">45ms</p>
            <p class="text-xs opacity-75">-12ms از ماه گذشته</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-yellow-500 to-yellow-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">درخواست‌های روزانه</p>
            <p class="text-2xl font-bold">1.2M</p>
            <p class="text-xs opacity-75">+15.3% از هفته گذشته</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- الگوریتم‌های موجود -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <!-- Collaborative Filtering -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-blue-100 rounded-lg">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">Collaborative Filtering</h3>
          </div>
          <label class="relative inline-flex items-center cursor-pointer">
            <input type="checkbox" v-model="algorithms.collaborative.enabled" class="sr-only peer">
            <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
          </label>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع الگوریتم</label>
            <select v-model="algorithms.collaborative.type" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
              <option value="user-based">User-Based</option>
              <option value="item-based">Item-Based</option>
              <option value="matrix-factorization">Matrix Factorization</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تعداد همسایه‌ها</label>
            <input type="range" v-model="algorithms.collaborative.neighbors" min="5" max="50" class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
            <div class="flex justify-between text-xs text-gray-500 mt-1">
              <span>5</span>
              <span>{{ algorithms.collaborative.neighbors }}</span>
              <span>50</span>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">وزن اهمیت</label>
            <div class="flex items-center space-x-4 space-x-reverse">
              <input type="range" v-model="algorithms.collaborative.weight" min="0" max="100" class="flex-1 h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
              <span class="text-sm font-medium text-gray-900 w-12">{{ algorithms.collaborative.weight }}%</span>
            </div>
          </div>
          
          <div class="flex items-center justify-between text-sm">
            <span class="text-gray-600">دقت: {{ algorithms.collaborative.accuracy }}%</span>
            <span class="text-gray-600">زمان پاسخ: {{ algorithms.collaborative.responseTime }}ms</span>
          </div>
        </div>
      </div>

      <!-- Content-Based Filtering -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-green-100 rounded-lg">
              <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">Content-Based Filtering</h3>
          </div>
          <label class="relative inline-flex items-center cursor-pointer">
            <input type="checkbox" v-model="algorithms.contentBased.enabled" class="sr-only peer">
            <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-green-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-green-600"></div>
          </label>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">ویژگی‌های مورد استفاده</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input type="checkbox" v-model="algorithms.contentBased.features" value="category" class="rounded border-gray-300 text-green-600 focus:ring-green-500">
                <span class="mr-2 text-sm text-gray-700">دسته‌بندی</span>
              </label>
              <label class="flex items-center">
                <input type="checkbox" v-model="algorithms.contentBased.features" value="brand" class="rounded border-gray-300 text-green-600 focus:ring-green-500">
                <span class="mr-2 text-sm text-gray-700">برند</span>
              </label>
              <label class="flex items-center">
                <input type="checkbox" v-model="algorithms.contentBased.features" value="price" class="rounded border-gray-300 text-green-600 focus:ring-green-500">
                <span class="mr-2 text-sm text-gray-700">قیمت</span>
              </label>
              <label class="flex items-center">
                <input type="checkbox" v-model="algorithms.contentBased.features" value="tags" class="rounded border-gray-300 text-green-600 focus:ring-green-500">
                <span class="mr-2 text-sm text-gray-700">برچسب‌ها</span>
              </label>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">آستانه تشابه</label>
            <input type="range" v-model="algorithms.contentBased.similarityThreshold" min="0.1" max="1" step="0.1" class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
            <div class="flex justify-between text-xs text-gray-500 mt-1">
              <span>0.1</span>
              <span>{{ algorithms.contentBased.similarityThreshold }}</span>
              <span>1.0</span>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">وزن اهمیت</label>
            <div class="flex items-center space-x-4 space-x-reverse">
              <input type="range" v-model="algorithms.contentBased.weight" min="0" max="100" class="flex-1 h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
              <span class="text-sm font-medium text-gray-900 w-12">{{ algorithms.contentBased.weight }}%</span>
            </div>
          </div>
          
          <div class="flex items-center justify-between text-sm">
            <span class="text-gray-600">دقت: {{ algorithms.contentBased.accuracy }}%</span>
            <span class="text-gray-600">زمان پاسخ: {{ algorithms.contentBased.responseTime }}ms</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Hybrid Approach -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <div class="flex items-center">
          <div class="p-2 bg-purple-100 rounded-lg">
            <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900 mr-3">Hybrid Approach</h3>
        </div>
        <label class="relative inline-flex items-center cursor-pointer">
          <input type="checkbox" v-model="algorithms.hybrid.enabled" class="sr-only peer">
          <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-purple-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-purple-600"></div>
        </label>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">روش ترکیب</label>
          <select v-model="algorithms.hybrid.method" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500">
            <option value="weighted">Weighted Average</option>
            <option value="switching">Switching</option>
            <option value="cascade">Cascade</option>
            <option value="ensemble">Ensemble</option>
          </select>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">وزن Collaborative</label>
          <input type="range" v-model="algorithms.hybrid.collaborativeWeight" min="0" max="100" class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
          <div class="flex justify-between text-xs text-gray-500 mt-1">
            <span>0%</span>
            <span>{{ algorithms.hybrid.collaborativeWeight }}%</span>
            <span>100%</span>
          </div>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">وزن Content-Based</label>
          <input type="range" v-model="algorithms.hybrid.contentWeight" min="0" max="100" class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
          <div class="flex justify-between text-xs text-gray-500 mt-1">
            <span>0%</span>
            <span>{{ algorithms.hybrid.contentWeight }}%</span>
            <span>100%</span>
          </div>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">آستانه تصمیم‌گیری</label>
          <input type="range" v-model="algorithms.hybrid.decisionThreshold" min="0.1" max="1" step="0.1" class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
          <div class="flex justify-between text-xs text-gray-500 mt-1">
            <span>0.1</span>
            <span>{{ algorithms.hybrid.decisionThreshold }}</span>
            <span>1.0</span>
          </div>
        </div>
      </div>
      
      <div class="mt-6 flex items-center justify-between text-sm">
        <span class="text-gray-600">دقت کلی: {{ algorithms.hybrid.accuracy }}%</span>
        <span class="text-gray-600">زمان پاسخ: {{ algorithms.hybrid.responseTime }}ms</span>
        <span class="text-gray-600">بهبود نسبت به بهترین الگوریتم: +{{ algorithms.hybrid.improvement }}%</span>
      </div>
    </div>

    <!-- Real-time Learning -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <div class="flex items-center">
          <div class="p-2 bg-yellow-100 rounded-lg">
            <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900 mr-3">Real-time Learning</h3>
        </div>
        <label class="relative inline-flex items-center cursor-pointer">
          <input type="checkbox" v-model="algorithms.realTimeLearning.enabled" class="sr-only peer">
          <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-yellow-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-yellow-600"></div>
        </label>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نرخ یادگیری</label>
          <input type="range" v-model="algorithms.realTimeLearning.learningRate" min="0.001" max="0.1" step="0.001" class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
          <div class="flex justify-between text-xs text-gray-500 mt-1">
            <span>0.001</span>
            <span>{{ algorithms.realTimeLearning.learningRate }}</span>
            <span>0.1</span>
          </div>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">فرکانس به‌روزرسانی</label>
          <select v-model="algorithms.realTimeLearning.updateFrequency" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500">
            <option value="realtime">Real-time</option>
            <option value="hourly">هر ساعت</option>
            <option value="daily">روزانه</option>
            <option value="weekly">هفتگی</option>
          </select>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حافظه کوتاه‌مدت</label>
          <input type="number" v-model="algorithms.realTimeLearning.shortTermMemory" min="100" max="10000" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500">
          <p class="text-xs text-gray-500 mt-1">تعداد آخرین تعاملات ذخیره شده</p>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">آستانه تغییر</label>
          <input type="range" v-model="algorithms.realTimeLearning.changeThreshold" min="0.01" max="0.5" step="0.01" class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
          <div class="flex justify-between text-xs text-gray-500 mt-1">
            <span>0.01</span>
            <span>{{ algorithms.realTimeLearning.changeThreshold }}</span>
            <span>0.5</span>
          </div>
        </div>
      </div>
      
      <div class="mt-6 p-6 bg-yellow-50 rounded-lg">
        <div class="flex items-center">
          <svg class="w-5 h-5 text-yellow-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
          </svg>
          <div>
            <p class="text-sm font-medium text-yellow-800">آخرین به‌روزرسانی: {{ algorithms.realTimeLearning.lastUpdate }}</p>
            <p class="text-xs text-yellow-600">تعداد تغییرات امروز: {{ algorithms.realTimeLearning.todayChanges }}</p>
          </div>
        </div>
      </div>
    </div>


  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
// تنظیمات الگوریتم‌ها
const algorithms = ref({
  collaborative: {
    enabled: true,
    type: 'user-based',
    neighbors: 20,
    weight: 60,
    accuracy: 82.5,
    responseTime: 45
  },
  contentBased: {
    enabled: true,
    features: ['category', 'brand', 'price'],
    similarityThreshold: 0.7,
    weight: 40,
    accuracy: 78.3,
    responseTime: 32
  },
  hybrid: {
    enabled: true,
    method: 'weighted',
    collaborativeWeight: 60,
    contentWeight: 40,
    decisionThreshold: 0.6,
    accuracy: 87.3,
    responseTime: 52,
    improvement: 4.8
  },
  realTimeLearning: {
    enabled: true,
    learningRate: 0.01,
    updateFrequency: 'realtime',
    shortTermMemory: 1000,
    changeThreshold: 0.1,
    lastUpdate: '2 دقیقه پیش',
    todayChanges: 156
  },

})
</script> 
