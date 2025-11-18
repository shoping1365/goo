<template>
  <div class="space-y-8">
    <!-- هدر بخش -->
    <div class="flex justify-between items-center">
      <div>
        <h2 class="text-2xl font-bold text-gray-900">تنظیمات پیشرفته</h2>
        <p class="mt-1 text-sm text-gray-600">تنظیمات پیشرفته و تخصصی سیستم پیشنهادات هوشمند</p>
      </div>
      <div class="flex space-x-3 space-x-reverse">
        <button class="bg-gray-600 text-white px-4 py-2 rounded-lg hover:bg-gray-700 transition-colors">
          <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
          </svg>
          بازنشانی تنظیمات
        </button>
        <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
          <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
          </svg>
          ذخیره تنظیمات
        </button>
      </div>
    </div>

    <!-- تنظیمات کش -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <div class="flex items-center">
          <div class="p-2 bg-blue-100 rounded-lg">
            <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7v10c0 2.21 3.582 4 8 4s8-1.79 8-4V7M4 7c0 2.21 3.582 4 8 4s8-1.79 8-4M4 7c0-2.21 3.582-4 8-4s8 1.79 8 4"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900 mr-3">تنظیمات کش</h3>
        </div>
        <label class="relative inline-flex items-center cursor-pointer">
          <input type="checkbox" v-model="cacheSettings.enabled" class="sr-only peer">
          <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
        </label>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع کش</label>
            <select v-model="cacheSettings.type" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
              <option value="redis">Redis</option>
              <option value="memory">Memory</option>
              <option value="file">File</option>
              <option value="database">Database</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">زمان انقضا (دقیقه)</label>
            <input type="range" v-model="cacheSettings.expiration" min="1" max="1440" class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
            <div class="flex justify-between text-xs text-gray-500 mt-1">
              <span>1 دقیقه</span>
              <span>{{ cacheSettings.expiration }} دقیقه</span>
              <span>24 ساعت</span>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر اندازه کش (MB)</label>
            <input type="range" v-model="cacheSettings.maxSize" min="10" max="1000" class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
            <div class="flex justify-between text-xs text-gray-500 mt-1">
              <span>10MB</span>
              <span>{{ cacheSettings.maxSize }}MB</span>
              <span>1GB</span>
            </div>
          </div>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">استراتژی حذف</label>
            <select v-model="cacheSettings.evictionPolicy" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
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
                <input type="checkbox" v-model="cacheSettings.compression.enabled" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <span class="mr-2 text-sm text-gray-700">فعال‌سازی فشرده‌سازی</span>
              </label>
              <label class="flex items-center">
                <input type="checkbox" v-model="cacheSettings.compression.gzip" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <span class="mr-2 text-sm text-gray-700">استفاده از Gzip</span>
              </label>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">آمار کش</label>
            <div class="bg-gray-50 rounded-lg p-3 space-y-2">
              <div class="flex justify-between text-sm">
                <span class="text-gray-600">نرخ برخورد:</span>
                <span class="font-medium">{{ cacheStats.hitRate }}%</span>
              </div>
              <div class="flex justify-between text-sm">
                <span class="text-gray-600">مقدار استفاده:</span>
                <span class="font-medium">{{ cacheStats.usage }}%</span>
              </div>
              <div class="flex justify-between text-sm">
                <span class="text-gray-600">تعداد کلیدها:</span>
                <span class="font-medium">{{ cacheStats.keyCount }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات امنیت -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <div class="flex items-center">
          <div class="p-2 bg-red-100 rounded-lg">
            <svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900 mr-3">تنظیمات امنیت</h3>
        </div>
        <label class="relative inline-flex items-center cursor-pointer">
          <input type="checkbox" v-model="securitySettings.enabled" class="sr-only peer">
          <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-red-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-red-600"></div>
        </label>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">رمزگذاری داده‌ها</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input type="checkbox" v-model="securitySettings.encryption.enabled" class="rounded border-gray-300 text-red-600 focus:ring-red-500">
                <span class="mr-2 text-sm text-gray-700">فعال‌سازی رمزگذاری</span>
              </label>
              <label class="flex items-center">
                <input type="checkbox" v-model="securitySettings.encryption.atRest" class="rounded border-gray-300 text-red-600 focus:ring-red-500">
                <span class="mr-2 text-sm text-gray-700">رمزگذاری در حالت استراحت</span>
              </label>
              <label class="flex items-center">
                <input type="checkbox" v-model="securitySettings.encryption.inTransit" class="rounded border-gray-300 text-red-600 focus:ring-red-500">
                <span class="mr-2 text-sm text-gray-700">رمزگذاری در انتقال</span>
              </label>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">الگوریتم رمزگذاری</label>
            <select v-model="securitySettings.encryption.algorithm" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500">
              <option value="aes-256">AES-256</option>
              <option value="aes-128">AES-128</option>
              <option value="chacha20">ChaCha20</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">کلید رمزگذاری</label>
            <div class="flex space-x-2 space-x-reverse">
              <input type="password" v-model="securitySettings.encryption.key" class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500" placeholder="کلید رمزگذاری">
              <button class="px-3 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700">تولید</button>
            </div>
          </div>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">محدودیت‌های دسترسی</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input type="checkbox" v-model="securitySettings.accessControl.ipWhitelist" class="rounded border-gray-300 text-red-600 focus:ring-red-500">
                <span class="mr-2 text-sm text-gray-700">لیست سفید IP</span>
              </label>
              <label class="flex items-center">
                <input type="checkbox" v-model="securitySettings.accessControl.rateLimit" class="rounded border-gray-300 text-red-600 focus:ring-red-500">
                <span class="mr-2 text-sm text-gray-700">محدودیت نرخ درخواست</span>
              </label>
              <label class="flex items-center">
                <input type="checkbox" v-model="securitySettings.accessControl.authentication" class="rounded border-gray-300 text-red-600 focus:ring-red-500">
                <span class="mr-2 text-sm text-gray-700">احراز هویت اجباری</span>
              </label>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">محدودیت نرخ (درخواست/دقیقه)</label>
            <input type="range" v-model="securitySettings.accessControl.rateLimitValue" min="10" max="1000" class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
            <div class="flex justify-between text-xs text-gray-500 mt-1">
              <span>10</span>
              <span>{{ securitySettings.accessControl.rateLimitValue }}</span>
              <span>1000</span>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">IP های مجاز</label>
            <textarea v-model="securitySettings.accessControl.allowedIPs" rows="3" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500" placeholder="هر IP در یک خط"></textarea>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات لاگ -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <div class="flex items-center">
          <div class="p-2 bg-green-100 rounded-lg">
            <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900 mr-3">تنظیمات لاگ</h3>
        </div>
        <label class="relative inline-flex items-center cursor-pointer">
          <input type="checkbox" v-model="loggingSettings.enabled" class="sr-only peer">
          <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-green-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-green-600"></div>
        </label>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">سطح لاگ</label>
            <select v-model="loggingSettings.level" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500">
              <option value="debug">Debug</option>
              <option value="info">Info</option>
              <option value="warning">Warning</option>
              <option value="error">Error</option>
              <option value="critical">Critical</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">محل ذخیره لاگ</label>
            <select v-model="loggingSettings.storage" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500">
              <option value="file">فایل</option>
              <option value="database">پایگاه داده</option>
              <option value="syslog">Syslog</option>
              <option value="elasticsearch">Elasticsearch</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر اندازه فایل لاگ (MB)</label>
            <input type="range" v-model="loggingSettings.maxFileSize" min="1" max="100" class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
            <div class="flex justify-between text-xs text-gray-500 mt-1">
              <span>1MB</span>
              <span>{{ loggingSettings.maxFileSize }}MB</span>
              <span>100MB</span>
            </div>
          </div>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع‌های لاگ</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input type="checkbox" v-model="loggingSettings.types.access" class="rounded border-gray-300 text-green-600 focus:ring-green-500">
                <span class="mr-2 text-sm text-gray-700">لاگ دسترسی</span>
              </label>
              <label class="flex items-center">
                <input type="checkbox" v-model="loggingSettings.types.error" class="rounded border-gray-300 text-green-600 focus:ring-green-500">
                <span class="mr-2 text-sm text-gray-700">لاگ خطا</span>
              </label>
              <label class="flex items-center">
                <input type="checkbox" v-model="loggingSettings.types.performance" class="rounded border-gray-300 text-green-600 focus:ring-green-500">
                <span class="mr-2 text-sm text-gray-700">لاگ عملکرد</span>
              </label>
              <label class="flex items-center">
                <input type="checkbox" v-model="loggingSettings.types.security" class="rounded border-gray-300 text-green-600 focus:ring-green-500">
                <span class="mr-2 text-sm text-gray-700">لاگ امنیت</span>
              </label>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">فرمت لاگ</label>
            <select v-model="loggingSettings.format" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500">
              <option value="json">JSON</option>
              <option value="text">Text</option>
              <option value="xml">XML</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">آمار لاگ</label>
            <div class="bg-gray-50 rounded-lg p-3 space-y-2">
              <div class="flex justify-between text-sm">
                <span class="text-gray-600">تعداد لاگ‌های امروز:</span>
                <span class="font-medium">{{ logStats.todayCount }}</span>
              </div>
              <div class="flex justify-between text-sm">
                <span class="text-gray-600">حجم فایل‌های لاگ:</span>
                <span class="font-medium">{{ logStats.totalSize }}MB</span>
              </div>
              <div class="flex justify-between text-sm">
                <span class="text-gray-600">خطاهای امروز:</span>
                <span class="font-medium">{{ logStats.errorCount }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات بهینه‌سازی -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <div class="flex items-center">
          <div class="p-2 bg-purple-100 rounded-lg">
            <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900 mr-3">تنظیمات بهینه‌سازی</h3>
        </div>
        <label class="relative inline-flex items-center cursor-pointer">
          <input type="checkbox" v-model="optimizationSettings.enabled" class="sr-only peer">
          <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-purple-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-purple-600"></div>
        </label>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">بهینه‌سازی الگوریتم</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input type="checkbox" v-model="optimizationSettings.algorithm.parallel" class="rounded border-gray-300 text-purple-600 focus:ring-purple-500">
                <span class="mr-2 text-sm text-gray-700">پردازش موازی</span>
              </label>
              <label class="flex items-center">
                <input type="checkbox" v-model="optimizationSettings.algorithm.batching" class="rounded border-gray-300 text-purple-600 focus:ring-purple-500">
                <span class="mr-2 text-sm text-gray-700">پردازش دسته‌ای</span>
              </label>
              <label class="flex items-center">
                <input type="checkbox" v-model="optimizationSettings.algorithm.caching" class="rounded border-gray-300 text-purple-600 focus:ring-purple-500">
                <span class="mr-2 text-sm text-gray-700">کش نتایج</span>
              </label>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تعداد thread ها</label>
            <input type="range" v-model="optimizationSettings.threads" min="1" max="16" class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
            <div class="flex justify-between text-xs text-gray-500 mt-1">
              <span>1</span>
              <span>{{ optimizationSettings.threads }}</span>
              <span>16</span>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">اندازه batch</label>
            <input type="range" v-model="optimizationSettings.batchSize" min="10" max="1000" class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
            <div class="flex justify-between text-xs text-gray-500 mt-1">
              <span>10</span>
              <span>{{ optimizationSettings.batchSize }}</span>
              <span>1000</span>
            </div>
          </div>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">بهینه‌سازی پایگاه داده</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input type="checkbox" v-model="optimizationSettings.database.indexing" class="rounded border-gray-300 text-purple-600 focus:ring-purple-500">
                <span class="mr-2 text-sm text-gray-700">بهینه‌سازی ایندکس</span>
              </label>
              <label class="flex items-center">
                <input type="checkbox" v-model="optimizationSettings.database.queryOptimization" class="rounded border-gray-300 text-purple-600 focus:ring-purple-500">
                <span class="mr-2 text-sm text-gray-700">بهینه‌سازی کوئری</span>
              </label>
              <label class="flex items-center">
                <input type="checkbox" v-model="optimizationSettings.database.connectionPooling" class="rounded border-gray-300 text-purple-600 focus:ring-purple-500">
                <span class="mr-2 text-sm text-gray-700">Connection Pooling</span>
              </label>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر اتصالات</label>
            <input type="range" v-model="optimizationSettings.maxConnections" min="5" max="100" class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
            <div class="flex justify-between text-xs text-gray-500 mt-1">
              <span>5</span>
              <span>{{ optimizationSettings.maxConnections }}</span>
              <span>100</span>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">آمار عملکرد</label>
            <div class="bg-gray-50 rounded-lg p-3 space-y-2">
              <div class="flex justify-between text-sm">
                <span class="text-gray-600">میانگین زمان پاسخ:</span>
                <span class="font-medium">{{ performanceStats.avgResponseTime }}ms</span>
              </div>
              <div class="flex justify-between text-sm">
                <span class="text-gray-600">درخواست‌های بر ثانیه:</span>
                <span class="font-medium">{{ performanceStats.requestsPerSecond }}</span>
              </div>
              <div class="flex justify-between text-sm">
                <span class="text-gray-600">استفاده از CPU:</span>
                <span class="font-medium">{{ performanceStats.cpuUsage }}%</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
// تنظیمات کش
const cacheSettings = ref({
  enabled: true,
  type: 'redis',
  expiration: 60,
  maxSize: 100,
  evictionPolicy: 'lru',
  compression: {
    enabled: true,
    gzip: true
  }
})

// آمار کش
const cacheStats = ref({
  hitRate: 87.3,
  usage: 65.2,
  keyCount: 1247
})

// تنظیمات امنیت
const securitySettings = ref({
  enabled: true,
  encryption: {
    enabled: true,
    atRest: true,
    inTransit: true,
    algorithm: 'aes-256',
    key: 'your-encryption-key-here'
  },
  accessControl: {
    ipWhitelist: true,
    rateLimit: true,
    authentication: true,
    rateLimitValue: 100,
    allowedIPs: '127.0.0.1\n192.168.1.0/24'
  }
})

// تنظیمات لاگ
const loggingSettings = ref({
  enabled: true,
  level: 'info',
  storage: 'file',
  maxFileSize: 10,
  types: {
    access: true,
    error: true,
    performance: true,
    security: true
  },
  format: 'json'
})

// آمار لاگ
const logStats = ref({
  todayCount: 1247,
  totalSize: 45.2,
  errorCount: 23
})

// تنظیمات بهینه‌سازی
const optimizationSettings = ref({
  enabled: true,
  algorithm: {
    parallel: true,
    batching: true,
    caching: true
  },
  threads: 8,
  batchSize: 100,
  database: {
    indexing: true,
    queryOptimization: true,
    connectionPooling: true
  },
  maxConnections: 20
})

// آمار عملکرد
const performanceStats = ref({
  avgResponseTime: 45,
  requestsPerSecond: 234,
  cpuUsage: 35.2
})
</script> 
