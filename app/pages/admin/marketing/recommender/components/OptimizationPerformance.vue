<template>
  <div class="space-y-8">
    <!-- هدر بخش -->
    <div class="flex justify-between items-center">
      <div>
        <h2 class="text-2xl font-bold text-gray-900">بهینه‌سازی و عملکرد</h2>
        <p class="mt-1 text-sm text-gray-600">مدیریت عملکرد و بهینه‌سازی سیستم پیشنهادات</p>
      </div>
      <button class="bg-orange-600 text-white px-4 py-2 rounded-lg hover:bg-orange-700 transition-colors">
        <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
        </svg>
        بهینه‌سازی خودکار
      </button>
    </div>

    <!-- آمار عملکرد -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-6">
      <div class="bg-gradient-to-r from-green-500 to-green-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">زمان پاسخ</p>
            <p class="text-2xl font-bold">45ms</p>
            <p class="text-xs opacity-75">میانگین</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-blue-500 to-blue-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">کش هیت</p>
            <p class="text-2xl font-bold">94.2%</p>
            <p class="text-xs opacity-75">نرخ موفقیت</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12h14M5 12a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v4a2 2 0 01-2 2M5 12a2 2 0 00-2 2v4a2 2 0 002 2h14a2 2 0 002-2v-4a2 2 0 00-2-2m-2-4h.01M17 16h.01"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-purple-500 to-purple-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">پردازش موازی</p>
            <p class="text-2xl font-bold">8</p>
            <p class="text-xs opacity-75">Thread فعال</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-yellow-500 to-yellow-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">کوئری/ثانیه</p>
            <p class="text-2xl font-bold">1,234</p>
            <p class="text-xs opacity-75">QPS</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-red-500 to-red-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">مقیاس‌پذیری</p>
            <p class="text-2xl font-bold">85%</p>
            <p class="text-xs opacity-75">آماده رشد</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 11l5-5m0 0l5 5m-5-5v12"></path>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات کش -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <!-- کش پیشنهادات -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-blue-100 rounded-lg">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12h14M5 12a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v4a2 2 0 01-2 2M5 12a2 2 0 00-2 2v4a2 2 0 002 2h14a2 2 0 002-2v-4a2 2 0 00-2-2m-2-4h.01M17 16h.01"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">کش پیشنهادات</h3>
          </div>
          <label class="relative inline-flex items-center cursor-pointer">
            <input v-model="cache.enabled" type="checkbox" class="sr-only peer">
            <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
          </label>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">زمان انقضا (دقیقه)</label>
            <input v-model="cache.expirationMinutes" type="number" min="1" max="1440" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر اندازه (MB)</label>
            <input v-model="cache.maxSizeMB" type="number" min="10" max="10000" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">استراتژی حذف</label>
            <select v-model="cache.evictionPolicy" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
              <option value="lru">LRU (Least Recently Used)</option>
              <option value="lfu">LFU (Least Frequently Used)</option>
              <option value="fifo">FIFO (First In First Out)</option>
              <option value="random">Random</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">فشرده‌سازی</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input v-model="cache.compression" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <span class="mr-2 text-sm text-gray-700">فعال‌سازی فشرده‌سازی</span>
              </label>
            </div>
          </div>
        </div>
      </div>

      <!-- پردازش موازی -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-purple-100 rounded-lg">
              <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">پردازش موازی</h3>
          </div>
          <label class="relative inline-flex items-center cursor-pointer">
            <input v-model="parallel.enabled" type="checkbox" class="sr-only peer">
            <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-purple-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-purple-600"></div>
          </label>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تعداد Thread</label>
            <input v-model="parallel.threadCount" type="range" min="1" max="16" class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
            <div class="flex justify-between text-xs text-gray-500 mt-1">
              <span>1</span>
              <span>{{ parallel.threadCount }}</span>
              <span>16</span>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر درخواست همزمان</label>
            <input v-model="parallel.maxConcurrentRequests" type="number" min="10" max="1000" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Timeout پردازش (ثانیه)</label>
            <input v-model="parallel.timeoutSeconds" type="number" min="1" max="300" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">استراتژی پردازش</label>
            <select v-model="parallel.strategy" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500">
              <option value="round_robin">Round Robin</option>
              <option value="least_loaded">Least Loaded</option>
              <option value="priority">Priority Based</option>
              <option value="adaptive">Adaptive</option>
            </select>
          </div>
        </div>
      </div>
    </div>

    <!-- بهینه‌سازی کوئری -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <!-- بهینه‌سازی دیتابیس -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-green-100 rounded-lg">
              <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7v10c0 2.21 3.582 4 8 4s8-1.79 8-4V7M4 7c0 2.21 3.582 4 8 4s8-1.79 8-4M4 7c0-2.21 3.582-4 8-4s8 1.79 8 4m0 5c0 2.21-3.582 4-8 4s-8-1.79-8-4"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">بهینه‌سازی کوئری</h3>
          </div>
          <button class="bg-green-600 text-white px-3 py-1 rounded-lg text-sm hover:bg-green-700 transition-colors">
            بهینه‌سازی
          </button>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">اندیس‌گذاری خودکار</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input v-model="database.autoIndexing" type="checkbox" class="rounded border-gray-300 text-green-600 focus:ring-green-500">
                <span class="mr-2 text-sm text-gray-700">فعال‌سازی اندیس‌گذاری خودکار</span>
              </label>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تحلیل کوئری</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input v-model="database.queryAnalysis" type="checkbox" class="rounded border-gray-300 text-green-600 focus:ring-green-500">
                <span class="mr-2 text-sm text-gray-700">تحلیل خودکار کوئری‌های کند</span>
              </label>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Connection Pool</label>
            <input v-model="database.connectionPoolSize" type="number" min="5" max="100" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Query Cache</label>
            <input v-model="database.queryCacheSize" type="number" min="10" max="1000" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500">
          </div>
        </div>
      </div>

      <!-- مقیاس‌پذیری -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-red-100 rounded-lg">
              <svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 11l5-5m0 0l5 5m-5-5v12"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">مقیاس‌پذیری</h3>
          </div>
          <button class="bg-red-600 text-white px-3 py-1 rounded-lg text-sm hover:bg-red-700 transition-colors">
            تنظیم
          </button>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Auto Scaling</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input v-model="scaling.autoScaling" type="checkbox" class="rounded border-gray-300 text-red-600 focus:ring-red-500">
                <span class="mr-2 text-sm text-gray-700">مقیاس‌بندی خودکار</span>
              </label>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداقل Instance</label>
            <input v-model="scaling.minInstances" type="number" min="1" max="10" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر Instance</label>
            <input v-model="scaling.maxInstances" type="number" min="1" max="50" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500">
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">آستانه CPU (%)</label>
            <input v-model="scaling.cpuThreshold" type="number" min="10" max="90" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500">
          </div>
        </div>
      </div>
    </div>

    <!-- Backup و Recovery -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <div class="flex items-center">
          <div class="p-2 bg-yellow-100 rounded-lg">
            <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-8l-4-4m0 0L8 8m4-4v12"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900 mr-3">Backup و Recovery</h3>
        </div>
        <div class="flex space-x-3 space-x-reverse">
          <button class="bg-yellow-600 text-white px-4 py-2 rounded-lg hover:bg-yellow-700 transition-colors">
            ایجاد Backup
          </button>
          <button class="bg-gray-600 text-white px-4 py-2 rounded-lg hover:bg-gray-700 transition-colors">
            بازیابی
          </button>
        </div>
      </div>
      
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">فرکانس Backup</label>
            <select v-model="backup.frequency" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500">
              <option value="hourly">ساعتی</option>
              <option value="daily">روزانه</option>
              <option value="weekly">هفتگی</option>
              <option value="monthly">ماهانه</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نگهداری Backup</label>
            <input v-model="backup.retentionDays" type="number" min="1" max="365" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500">
            <p class="text-xs text-gray-500 mt-1">روز</p>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">فشرده‌سازی</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input v-model="backup.compression" type="checkbox" class="rounded border-gray-300 text-yellow-600 focus:ring-yellow-500">
                <span class="mr-2 text-sm text-gray-700">فشرده‌سازی Backup</span>
              </label>
            </div>
          </div>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">رمزنگاری</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input v-model="backup.encryption" type="checkbox" class="rounded border-gray-300 text-yellow-600 focus:ring-yellow-500">
                <span class="mr-2 text-sm text-gray-700">رمزنگاری Backup</span>
              </label>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مکان ذخیره</label>
            <select v-model="backup.storageLocation" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500">
              <option value="local">محلی</option>
              <option value="cloud">ابر</option>
              <option value="hybrid">ترکیبی</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تست بازیابی</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input v-model="backup.testRecovery" type="checkbox" class="rounded border-gray-300 text-yellow-600 focus:ring-yellow-500">
                <span class="mr-2 text-sm text-gray-700">تست خودکار بازیابی</span>
              </label>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- گزارش‌های عملکرد -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">گزارش‌های عملکرد</h3>
      </div>
      <div class="p-6">
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">معیار</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مقدار فعلی</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">هدف</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">روند</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="metric in performanceMetrics" :key="metric.id">
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="flex items-center">
                    <div class="w-8 h-8 rounded-lg mr-3 flex items-center justify-center" :class="metric.bgColor">
                      <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="metric.icon"></path>
                      </svg>
                    </div>
                    <div>
                      <div class="text-sm font-medium text-gray-900">{{ metric.name }}</div>
                      <div class="text-sm text-gray-500">{{ metric.description }}</div>
                    </div>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ metric.currentValue }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ metric.targetValue }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full" :class="metric.statusClass">
                    {{ metric.status }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  <div class="flex items-center">
                    <svg class="w-4 h-4 mr-1" :class="metric.trendColor" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="metric.trendIcon"></path>
                    </svg>
                    {{ metric.trend }}
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
// تنظیمات کش
const cache = ref({
  enabled: true,
  expirationMinutes: 30,
  maxSizeMB: 500,
  evictionPolicy: 'lru',
  compression: true
})

// تنظیمات پردازش موازی
const parallel = ref({
  enabled: true,
  threadCount: 8,
  maxConcurrentRequests: 100,
  timeoutSeconds: 30,
  strategy: 'round_robin'
})

// تنظیمات دیتابیس
const database = ref({
  autoIndexing: true,
  queryAnalysis: true,
  connectionPoolSize: 20,
  queryCacheSize: 100
})

// تنظیمات مقیاس‌پذیری
const scaling = ref({
  autoScaling: true,
  minInstances: 2,
  maxInstances: 10,
  cpuThreshold: 70
})

// تنظیمات Backup
const backup = ref({
  frequency: 'daily',
  retentionDays: 30,
  compression: true,
  encryption: true,
  storageLocation: 'cloud',
  testRecovery: true
})

// معیارهای عملکرد
const performanceMetrics = ref([
  {
    id: 1,
    name: 'زمان پاسخ',
    description: 'میانگین زمان پاسخ API',
    currentValue: '45ms',
    targetValue: '< 50ms',
    status: 'عالی',
    statusClass: 'bg-green-100 text-green-800',
    trend: '+2ms',
    trendColor: 'text-red-500',
    trendIcon: 'M13 7h8m0 0v8m0-8l-8 8-4-4-6 6',
    bgColor: 'bg-green-500',
    icon: 'M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z'
  },
  {
    id: 2,
    name: 'کش هیت',
    description: 'نرخ موفقیت کش',
    currentValue: '94.2%',
    targetValue: '> 90%',
    status: 'عالی',
    statusClass: 'bg-green-100 text-green-800',
    trend: '+1.2%',
    trendColor: 'text-green-500',
    trendIcon: 'M13 7h8m0 0v8m0-8l-8 8-4-4-6 6',
    bgColor: 'bg-blue-500',
    icon: 'M5 12h14M5 12a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v4a2 2 0 01-2 2M5 12a2 2 0 00-2 2v4a2 2 0 002 2h14a2 2 0 002-2v-4a2 2 0 00-2-2m-2-4h.01M17 16h.01'
  },
  {
    id: 3,
    name: 'QPS',
    description: 'درخواست در ثانیه',
    currentValue: '1,234',
    targetValue: '> 1,000',
    status: 'خوب',
    statusClass: 'bg-yellow-100 text-yellow-800',
    trend: '+45',
    trendColor: 'text-green-500',
    trendIcon: 'M13 7h8m0 0v8m0-8l-8 8-4-4-6 6',
    bgColor: 'bg-purple-500',
    icon: ''
  },
  {
    id: 4,
    name: 'مقیاس‌پذیری',
    description: 'آمادگی برای رشد',
    currentValue: '85%',
    targetValue: '> 80%',
    status: 'خوب',
    statusClass: 'bg-yellow-100 text-yellow-800',
    trend: '+5%',
    trendColor: 'text-green-500',
    trendIcon: 'M13 7h8m0 0v8m0-8l-8 8-4-4-6 6',
    bgColor: 'bg-red-500',
    icon: 'M7 11l5-5m0 0l5 5m-5-5v12'
  }
])
</script> 
