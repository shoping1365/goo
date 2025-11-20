<template>
  <div class="space-y-8">
    <!-- بخش بهینه‌سازی -->
    <section>
      <div class="bg-gradient-to-r from-teal-50 to-cyan-50 rounded-lg p-6 mb-6">
        <h2 class="text-2xl font-bold text-gray-900 mb-2">⚡ بهینه‌سازی</h2>
        <p class="text-gray-600">بهینه‌سازی عملکرد سیستم و بهبود سرعت بارگذاری</p>
      </div>
      
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- کش کردن داده‌ها -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">کش کردن داده‌ها</h3>
            <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">فعال</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">نرخ ضربه (Hit Rate):</span>
              <span class="text-lg font-bold text-green-600">{{ optimizationStats.cacheHitRate }}%</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">حجم کش:</span>
              <span class="text-lg font-bold text-blue-600">{{ optimizationStats.cacheSize }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">آیتم‌های کش شده:</span>
              <span class="text-lg font-bold text-purple-600">{{ optimizationStats.cachedItems }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">زمان انقضای کش:</span>
              <span class="text-sm font-medium text-gray-900">{{ optimizationStats.cacheExpiry }}</span>
            </div>
          </div>
          
          <!-- نوار پیشرفت کش -->
          <div class="mt-4">
            <div class="flex justify-between text-sm text-gray-500 mb-1">
              <span>استفاده از کش</span>
              <span>{{ optimizationStats.cacheUsage }}%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div
class="bg-gradient-to-r from-green-500 to-blue-500 h-2 rounded-full transition-all duration-300"
                   :style="{ width: optimizationStats.cacheUsage + '%' }"></div>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-green-600 text-white text-sm rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500">
              پاک کردن کش
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              تنظیمات
            </button>
          </div>
        </div>

        <!-- بهینه‌سازی کوئری‌ها -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">بهینه‌سازی کوئری‌ها</h3>
            <span class="text-xs bg-blue-100 text-blue-800 rounded-full px-3 py-1">فعال</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">میانگین زمان کوئری:</span>
              <span class="text-lg font-bold text-green-600">{{ optimizationStats.avgQueryTime }}ms</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">کوئری‌های کند:</span>
              <span class="text-lg font-bold text-red-600">{{ optimizationStats.slowQueries }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">کوئری‌های بهینه شده:</span>
              <span class="text-lg font-bold text-blue-600">{{ optimizationStats.optimizedQueries }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">اندیس‌های فعال:</span>
              <span class="text-lg font-bold text-purple-600">{{ optimizationStats.activeIndexes }}</span>
            </div>
          </div>
          
          <!-- نمودار عملکرد کوئری -->
          <div class="mt-4">
            <div class="flex justify-between text-sm text-gray-500 mb-1">
              <span>بهبود عملکرد</span>
              <span>{{ optimizationStats.performanceImprovement }}%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div
class="bg-gradient-to-r from-blue-500 to-purple-500 h-2 rounded-full transition-all duration-300"
                   :style="{ width: optimizationStats.performanceImprovement + '%' }"></div>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
              تحلیل کوئری
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              بهینه‌سازی
            </button>
          </div>
        </div>

        <!-- فشرده‌سازی تصاویر -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">فشرده‌سازی تصاویر</h3>
            <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">فعال</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">تصاویر فشرده شده:</span>
              <span class="text-lg font-bold text-green-600">{{ optimizationStats.compressedImages }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">صرفه‌جویی در حجم:</span>
              <span class="text-lg font-bold text-blue-600">{{ optimizationStats.storageSaved }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">درصد فشرده‌سازی:</span>
              <span class="text-lg font-bold text-purple-600">{{ optimizationStats.compressionRatio }}%</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">در انتظار فشرده‌سازی:</span>
              <span class="text-lg font-bold text-orange-600">{{ optimizationStats.pendingCompression }}</span>
            </div>
          </div>
          
          <!-- نوار پیشرفت فشرده‌سازی -->
          <div class="mt-4">
            <div class="flex justify-between text-sm text-gray-500 mb-1">
              <span>پیشرفت فشرده‌سازی</span>
              <span>{{ optimizationStats.compressionProgress }}%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div
class="bg-gradient-to-r from-purple-500 to-pink-500 h-2 rounded-full transition-all duration-300"
                   :style="{ width: optimizationStats.compressionProgress + '%' }"></div>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-purple-600 text-white text-sm rounded-lg hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-purple-500">
              فشرده‌سازی دستی
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              تنظیمات
            </button>
          </div>
        </div>

        <!-- CDN برای فایل‌ها -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">CDN برای فایل‌ها</h3>
            <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">فعال</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">فایل‌های CDN:</span>
              <span class="text-lg font-bold text-green-600">{{ optimizationStats.cdnFiles }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">ترافیک CDN:</span>
              <span class="text-lg font-bold text-blue-600">{{ optimizationStats.cdnTraffic }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">سرعت بارگذاری:</span>
              <span class="text-lg font-bold text-purple-600">{{ optimizationStats.loadSpeed }}ms</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">نقاط توزیع:</span>
              <span class="text-lg font-bold text-orange-600">{{ optimizationStats.cdnNodes }}</span>
            </div>
          </div>
          
          <!-- نوار پیشرفت CDN -->
          <div class="mt-4">
            <div class="flex justify-between text-sm text-gray-500 mb-1">
              <span>پوشش CDN</span>
              <span>{{ optimizationStats.cdnCoverage }}%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div
class="bg-gradient-to-r from-orange-500 to-red-500 h-2 rounded-full transition-all duration-300"
                   :style="{ width: optimizationStats.cdnCoverage + '%' }"></div>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-orange-600 text-white text-sm rounded-lg hover:bg-orange-700 focus:outline-none focus:ring-2 focus:ring-orange-500">
              مدیریت CDN
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              آمار
            </button>
          </div>
        </div>
      </div>
    </section>

    <!-- بخش مانیتورینگ -->
    <section>
      <div class="bg-gradient-to-r from-purple-50 to-indigo-50 rounded-lg p-6 mb-6">
        <h2 class="text-2xl font-bold text-gray-900 mb-2">📊 مانیتورینگ</h2>
        <p class="text-gray-600">نظارت بر عملکرد سیستم و هشدار خطاها</p>
      </div>
      
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- نظارت بر عملکرد -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">نظارت بر عملکرد</h3>
            <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">فعال</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">CPU استفاده شده:</span>
              <span class="text-lg font-bold text-blue-600">{{ monitoringStats.cpuUsage }}%</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">حافظه استفاده شده:</span>
              <span class="text-lg font-bold text-purple-600">{{ monitoringStats.memoryUsage }}%</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">دیسک استفاده شده:</span>
              <span class="text-lg font-bold text-orange-600">{{ monitoringStats.diskUsage }}%</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">شبکه:</span>
              <span class="text-lg font-bold text-green-600">{{ monitoringStats.networkSpeed }}</span>
            </div>
          </div>
          
          <!-- نمودار عملکرد -->
          <div class="mt-4">
            <div class="flex justify-between text-sm text-gray-500 mb-1">
              <span>وضعیت کلی</span>
              <span>{{ monitoringStats.overallHealth }}%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div
class="bg-gradient-to-r from-green-500 to-blue-500 h-2 rounded-full transition-all duration-300"
                   :style="{ width: monitoringStats.overallHealth + '%' }"></div>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
              مشاهده جزئیات
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              تنظیمات
            </button>
          </div>
        </div>

        <!-- هشدار خطاها -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">هشدار خطاها</h3>
            <span class="text-xs bg-red-100 text-red-800 rounded-full px-3 py-1">فعال</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">خطاهای بحرانی:</span>
              <span class="text-lg font-bold text-red-600">{{ monitoringStats.criticalErrors }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">خطاهای هشدار:</span>
              <span class="text-lg font-bold text-orange-600">{{ monitoringStats.warningErrors }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">خطاهای اطلاعاتی:</span>
              <span class="text-lg font-bold text-blue-600">{{ monitoringStats.infoErrors }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">آخرین خطا:</span>
              <span class="text-sm text-gray-900">{{ monitoringStats.lastError }}</span>
            </div>
          </div>
          
          <!-- نوار پیشرفت خطاها -->
          <div class="mt-4">
            <div class="flex justify-between text-sm text-gray-500 mb-1">
              <span>نرخ خطا</span>
              <span>{{ monitoringStats.errorRate }}%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div
class="bg-gradient-to-r from-red-500 to-orange-500 h-2 rounded-full transition-all duration-300"
                   :style="{ width: monitoringStats.errorRate + '%' }"></div>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-red-600 text-white text-sm rounded-lg hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500">
              مشاهده خطاها
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              تنظیمات هشدار
            </button>
          </div>
        </div>

        <!-- گزارش‌های عملکرد -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">گزارش‌های عملکرد</h3>
            <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">فعال</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">گزارش‌های تولید شده:</span>
              <span class="text-lg font-bold text-green-600">{{ monitoringStats.generatedReports }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">گزارش‌های خودکار:</span>
              <span class="text-lg font-bold text-blue-600">{{ monitoringStats.autoReports }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">گزارش‌های دستی:</span>
              <span class="text-lg font-bold text-purple-600">{{ monitoringStats.manualReports }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">آخرین گزارش:</span>
              <span class="text-sm text-gray-900">{{ monitoringStats.lastReport }}</span>
            </div>
          </div>
          
          <!-- نوار پیشرفت گزارش‌ها -->
          <div class="mt-4">
            <div class="flex justify-between text-sm text-gray-500 mb-1">
              <span>پوشش گزارش‌ها</span>
              <span>{{ monitoringStats.reportCoverage }}%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div
class="bg-gradient-to-r from-green-500 to-blue-500 h-2 rounded-full transition-all duration-300"
                   :style="{ width: monitoringStats.reportCoverage + '%' }"></div>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-green-600 text-white text-sm rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500">
              ایجاد گزارش
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              تاریخچه
            </button>
          </div>
        </div>

        <!-- تحلیل ترافیک -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-gray-900">تحلیل ترافیک</h3>
            <span class="text-xs bg-blue-100 text-blue-800 rounded-full px-3 py-1">فعال</span>
          </div>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">درخواست‌های امروز:</span>
              <span class="text-lg font-bold text-blue-600">{{ monitoringStats.requestsToday }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">کاربران فعال:</span>
              <span class="text-lg font-bold text-green-600">{{ monitoringStats.activeUsers }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">پهنای باند استفاده شده:</span>
              <span class="text-lg font-bold text-purple-600">{{ monitoringStats.bandwidthUsed }}</span>
            </div>
            
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-600">میانگین زمان پاسخ:</span>
              <span class="text-lg font-bold text-orange-600">{{ monitoringStats.avgResponseTime }}ms</span>
            </div>
          </div>
          
          <!-- نوار پیشرفت ترافیک -->
          <div class="mt-4">
            <div class="flex justify-between text-sm text-gray-500 mb-1">
              <span>بارگذاری سیستم</span>
              <span>{{ monitoringStats.systemLoad }}%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div
class="bg-gradient-to-r from-purple-500 to-pink-500 h-2 rounded-full transition-all duration-300"
                   :style="{ width: monitoringStats.systemLoad + '%' }"></div>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="px-4 py-2 bg-purple-600 text-white text-sm rounded-lg hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-purple-500">
              تحلیل ترافیک
            </button>
            <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
              گزارش‌ها
            </button>
          </div>
        </div>
      </div>
    </section>

    <!-- داشبورد کلی عملکرد -->
    <section>
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-xl font-semibold text-gray-900 mb-4">📈 آمار کلی عملکرد</h3>
        
        <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
          <div class="text-center p-6 bg-green-50 rounded-lg">
            <div class="text-2xl font-bold text-green-600">{{ overallStats.performanceScore }}%</div>
            <div class="text-sm text-gray-600">امتیاز عملکرد</div>
          </div>
          
          <div class="text-center p-6 bg-blue-50 rounded-lg">
            <div class="text-2xl font-bold text-blue-600">{{ overallStats.uptime }}%</div>
            <div class="text-sm text-gray-600">زمان فعالیت</div>
          </div>
          
          <div class="text-center p-6 bg-purple-50 rounded-lg">
            <div class="text-2xl font-bold text-purple-600">{{ overallStats.responseTime }}ms</div>
            <div class="text-sm text-gray-600">زمان پاسخ</div>
          </div>
          
          <div class="text-center p-6 bg-orange-50 rounded-lg">
            <div class="text-2xl font-bold text-orange-600">{{ overallStats.errorRate }}%</div>
            <div class="text-sm text-gray-600">نرخ خطا</div>
          </div>
        </div>
        
        <!-- نمودار روند عملکرد -->
        <div class="mt-6">
          <h4 class="text-sm font-medium text-gray-900 mb-3">روند عملکرد (7 روز گذشته)</h4>
          <div class="flex items-end space-x-2 space-x-reverse h-32 overflow-x-auto">
            <div v-for="(day, index) in performanceTrend" :key="index" class="flex-shrink-0 flex flex-col items-center min-w-16">
              <div
class="w-full bg-gray-200 rounded-t relative"
                   :style="{ height: getChartHeight(day.performance) + 'px' }">
                <div
class="w-full bg-gradient-to-t from-blue-500 to-purple-500 rounded-t transition-all duration-300 absolute bottom-0"
                     :style="{ height: getChartHeight(day.performance) + 'px' }"></div>
              </div>
              <span class="text-xs text-gray-500 mt-1 text-center">{{ day.label }}</span>
              <span class="text-xs text-gray-400 mt-1 text-center">{{ day.performance }}%</span>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
// آمار بهینه‌سازی
const optimizationStats = {
  cacheHitRate: 87,
  cacheSize: '2.4 GB',
  cachedItems: 15420,
  cacheExpiry: '1 ساعت',
  cacheUsage: 78,
  avgQueryTime: 45,
  slowQueries: 23,
  optimizedQueries: 12450,
  activeIndexes: 156,
  performanceImprovement: 34,
  compressedImages: 8920,
  storageSaved: '1.2 GB',
  compressionRatio: 65,
  pendingCompression: 234,
  compressionProgress: 89,
  cdnFiles: 5670,
  cdnTraffic: '45.2 GB',
  loadSpeed: 120,
  cdnNodes: 12,
  cdnCoverage: 95
}

// آمار مانیتورینگ
const monitoringStats = {
  cpuUsage: 45,
  memoryUsage: 67,
  diskUsage: 78,
  networkSpeed: '125 Mbps',
  overallHealth: 89,
  criticalErrors: 3,
  warningErrors: 12,
  infoErrors: 45,
  lastError: '2 ساعت پیش',
  errorRate: 2.3,
  generatedReports: 156,
  autoReports: 134,
  manualReports: 22,
  lastReport: 'امروز 14:30',
  reportCoverage: 92,
  requestsToday: 45670,
  activeUsers: 1234,
  bandwidthUsed: '2.1 GB',
  avgResponseTime: 180,
  systemLoad: 67
}

// آمار کلی
const overallStats = {
  performanceScore: 94,
  uptime: 99.8,
  responseTime: 180,
  errorRate: 2.3
}

// روند عملکرد
const performanceTrend = [
  { label: 'شنبه', performance: 92 },
  { label: 'یکشنبه', performance: 89 },
  { label: 'دوشنبه', performance: 95 },
  { label: 'سه‌شنبه', performance: 91 },
  { label: 'چهارشنبه', performance: 88 },
  { label: 'پنج‌شنبه', performance: 93 },
  { label: 'جمعه', performance: 94 }
]

// تابع محاسبه ارتفاع مناسب برای نمودار روند عملکرد
const getChartHeight = (percentage: number) => {
  const maxHeight = 80; // حداکثر ارتفاع مناسب برای نمودار
  const minHeight = 10; // حداقل ارتفاع مناسب برای نمودار
  return Math.max(minHeight, percentage * (maxHeight / 100));
}
</script>

<!--
  مستندسازی:
  این کامپوننت شامل دو بخش اصلی است:
  1. بهینه‌سازی: کش کردن، بهینه‌سازی کوئری‌ها، فشرده‌سازی تصاویر، CDN
  2. مانیتورینگ: نظارت بر عملکرد، هشدار خطاها، گزارش‌ها، تحلیل ترافیک
  
  تمام توضیحات به فارسی و با طراحی ریسپانسیو ارائه شده است.
  نمودارها و آمار دقیق برای نمایش وضعیت عملکرد استفاده شده‌اند.
--> 
