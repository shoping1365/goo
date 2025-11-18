<template>
  <div class="performance-optimizer bg-white rounded-2xl shadow-xl border border-gray-100 p-8">
    <div class="flex items-center mb-6">
      <div class="p-3 bg-gradient-to-r from-emerald-400 to-emerald-600 rounded-xl shadow-lg">
        <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
        </svg>
      </div>
      <div class="mr-4">
        <h2 class="text-2xl font-bold text-gray-900">بهینه‌سازی عملکرد</h2>
        <p class="text-gray-600 mt-1">مدیریت کش، فشرده‌سازی و بهینه‌سازی سرعت</p>
      </div>
    </div>

    <!-- Performance Stats -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
      <div class="bg-gradient-to-r from-emerald-50 to-emerald-100 rounded-xl p-6 border border-emerald-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-emerald-600">سرعت بارگذاری</p>
            <p class="text-3xl font-bold text-emerald-900">{{ performanceStats.loadTime }}s</p>
          </div>
          <div class="p-3 bg-emerald-500 rounded-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
            </svg>
          </div>
        </div>
        <div class="mt-4 text-sm text-emerald-700">
          {{ performanceStats.loadTimeStatus }}
        </div>
      </div>

      <div class="bg-gradient-to-r from-blue-50 to-blue-100 rounded-xl p-6 border border-blue-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-blue-600">نرخ کش</p>
            <p class="text-3xl font-bold text-blue-900">{{ performanceStats.cacheHitRate }}%</p>
          </div>
          <div class="p-3 bg-blue-500 rounded-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7v10c0 2.21 3.582 4 8 4s8-1.79 8-4V7M4 7c0 2.21 3.582 4 8 4s8-1.79 8-4M4 7c0-2.21 3.582-4 8-4s8 1.79 8 4"></path>
            </svg>
          </div>
        </div>
        <div class="mt-4 text-sm text-blue-700">
          {{ performanceStats.cacheSize }}MB حجم کش
        </div>
      </div>

      <div class="bg-gradient-to-r from-purple-50 to-purple-100 rounded-xl p-6 border border-purple-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-purple-600">فشرده‌سازی</p>
            <p class="text-3xl font-bold text-purple-900">{{ performanceStats.compressionRatio }}%</p>
          </div>
          <div class="p-3 bg-purple-500 rounded-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
            </svg>
          </div>
        </div>
        <div class="mt-4 text-sm text-purple-700">
          {{ performanceStats.savedBandwidth }}MB صرفه‌جویی
        </div>
      </div>

      <div class="bg-gradient-to-r from-orange-50 to-orange-100 rounded-xl p-6 border border-orange-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-orange-600">درخواست‌های همزمان</p>
            <p class="text-3xl font-bold text-orange-900">{{ performanceStats.concurrentRequests }}</p>
          </div>
          <div class="p-3 bg-orange-500 rounded-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
          </div>
        </div>
        <div class="mt-4 text-sm text-orange-700">
          {{ performanceStats.avgResponseTime }}ms میانگین
        </div>
      </div>
    </div>

    <!-- Caching Settings -->
    <div class="bg-emerald-50 rounded-xl p-6 mb-8">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات کش</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">کش مرورگر</label>
          <div class="space-y-2">
            <label class="flex items-center">
              <input type="checkbox" v-model="cachingSettings.browserCache.enabled" class="rounded border-gray-300">
              <span class="text-sm text-gray-700 mr-2">فعال کردن کش مرورگر</span>
            </label>
            <div>
              <label class="text-sm text-gray-600">مدت زمان کش (ساعت)</label>
              <input v-model.number="cachingSettings.browserCache.duration" type="number" min="1" max="8760" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-emerald-500">
            </div>
            <select v-model="cachingSettings.browserCache.strategy" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-emerald-500">
              <option value="aggressive">تهاجمی</option>
              <option value="moderate">متوسط</option>
              <option value="conservative">محافظه‌کارانه</option>
            </select>
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">کش سرور</label>
          <div class="space-y-2">
            <label class="flex items-center">
              <input type="checkbox" v-model="cachingSettings.serverCache.enabled" class="rounded border-gray-300">
              <span class="text-sm text-gray-700 mr-2">فعال کردن کش سرور</span>
            </label>
            <div>
              <label class="text-sm text-gray-600">حداکثر حجم کش (MB)</label>
              <input v-model.number="cachingSettings.serverCache.maxSize" type="number" min="100" max="10000" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-emerald-500">
            </div>
            <select v-model="cachingSettings.serverCache.type" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-emerald-500">
              <option value="redis">Redis</option>
              <option value="memcached">Memcached</option>
              <option value="file">فایل</option>
              <option value="memory">حافظه</option>
            </select>
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">کش CDN</label>
          <div class="space-y-2">
            <label class="flex items-center">
              <input type="checkbox" v-model="cachingSettings.cdnCache.enabled" class="rounded border-gray-300">
              <span class="text-sm text-gray-700 mr-2">فعال کردن کش CDN</span>
            </label>
            <div>
              <label class="text-sm text-gray-600">مدت زمان کش CDN (ساعت)</label>
              <input v-model.number="cachingSettings.cdnCache.duration" type="number" min="1" max="8760" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-emerald-500">
            </div>
            <label class="flex items-center">
              <input type="checkbox" v-model="cachingSettings.cdnCache.autoPurge" class="rounded border-gray-300">
              <span class="text-sm text-gray-700 mr-2">پاک‌سازی خودکار</span>
            </label>
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تنظیمات پیشرفته</label>
          <div class="space-y-2">
            <label class="flex items-center">
              <input type="checkbox" v-model="cachingSettings.advanced.compression" class="rounded border-gray-300">
              <span class="text-sm text-gray-700 mr-2">فشرده‌سازی کش</span>
            </label>
            <label class="flex items-center">
              <input type="checkbox" v-model="cachingSettings.advanced.encryption" class="rounded border-gray-300">
              <span class="text-sm text-gray-700 mr-2">رمزگذاری کش</span>
            </label>
            <label class="flex items-center">
              <input type="checkbox" v-model="cachingSettings.advanced.monitoring" class="rounded border-gray-300">
              <span class="text-sm text-gray-700 mr-2">مانیتورینگ کش</span>
            </label>
          </div>
        </div>
      </div>
    </div>

    <!-- Compression Settings -->
    <div class="bg-blue-50 rounded-xl p-6 mb-8">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات فشرده‌سازی</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">فشرده‌سازی Gzip</label>
          <div class="space-y-2">
            <label class="flex items-center">
              <input type="checkbox" v-model="compressionSettings.gzip.enabled" class="rounded border-gray-300">
              <span class="text-sm text-gray-700 mr-2">فعال کردن Gzip</span>
            </label>
            <select v-model="compressionSettings.gzip.level" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
              <option value="1">سطح ۱ (سریع)</option>
              <option value="3">سطح ۳ (متوسط)</option>
              <option value="6">سطح ۶ (بهینه)</option>
              <option value="9">سطح ۹ (حداکثر)</option>
            </select>
            <div class="text-sm text-gray-600">
              فایل‌های بزرگتر از {{ compressionSettings.gzip.minSize }}KB فشرده می‌شوند
            </div>
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">فشرده‌سازی Brotli</label>
          <div class="space-y-2">
            <label class="flex items-center">
              <input type="checkbox" v-model="compressionSettings.brotli.enabled" class="rounded border-gray-300">
              <span class="text-sm text-gray-700 mr-2">فعال کردن Brotli</span>
            </label>
            <select v-model="compressionSettings.brotli.level" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
              <option value="1">سطح ۱ (سریع)</option>
              <option value="4">سطح ۴ (متوسط)</option>
              <option value="7">سطح ۷ (بهینه)</option>
              <option value="11">سطح ۱۱ (حداکثر)</option>
            </select>
            <div class="text-sm text-gray-600">
              فایل‌های بزرگتر از {{ compressionSettings.brotli.minSize }}KB فشرده می‌شوند
            </div>
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">فشرده‌سازی تصاویر</label>
          <div class="space-y-2">
            <label class="flex items-center">
              <input type="checkbox" v-model="compressionSettings.imageCompression.enabled" class="rounded border-gray-300">
              <span class="text-sm text-gray-700 mr-2">فشرده‌سازی خودکار تصاویر</span>
            </label>
            <select v-model="compressionSettings.imageCompression.quality" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
              <option value="60">کیفیت پایین (۶۰%)</option>
              <option value="75">کیفیت متوسط (۷۵%)</option>
              <option value="85">کیفیت بالا (۸۵%)</option>
              <option value="95">کیفیت عالی (۹۵%)</option>
            </select>
            <label class="flex items-center">
              <input type="checkbox" v-model="compressionSettings.imageCompression.webp" class="rounded border-gray-300">
              <span class="text-sm text-gray-700 mr-2">تبدیل به WebP</span>
            </label>
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تنظیمات پیشرفته</label>
          <div class="space-y-2">
            <label class="flex items-center">
              <input type="checkbox" v-model="compressionSettings.advanced.minifyCSS" class="rounded border-gray-300">
              <span class="text-sm text-gray-700 mr-2">فشرده‌سازی CSS</span>
            </label>
            <label class="flex items-center">
              <input type="checkbox" v-model="compressionSettings.advanced.minifyJS" class="rounded border-gray-300">
              <span class="text-sm text-gray-700 mr-2">فشرده‌سازی JavaScript</span>
            </label>
            <label class="flex items-center">
              <input type="checkbox" v-model="compressionSettings.advanced.minifyHTML" class="rounded border-gray-300">
              <span class="text-sm text-gray-700 mr-2">فشرده‌سازی HTML</span>
            </label>
          </div>
        </div>
      </div>
    </div>

    <!-- Load Balancing -->
    <div class="bg-purple-50 rounded-xl p-6 mb-8">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">توزیع بار</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">فعال‌سازی توزیع بار</label>
          <div class="space-y-2">
            <label class="flex items-center">
              <input type="checkbox" v-model="loadBalancingSettings.enabled" class="rounded border-gray-300">
              <span class="text-sm text-gray-700 mr-2">فعال کردن توزیع بار</span>
            </label>
            <select v-model="loadBalancingSettings.algorithm" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500">
              <option value="round_robin">Round Robin</option>
              <option value="least_connections">کمترین اتصال</option>
              <option value="ip_hash">توزیع بر اساس IP</option>
              <option value="weighted">وزن‌دار</option>
            </select>
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">سرورهای موجود</label>
          <div class="space-y-2">
            <div v-for="server in loadBalancingSettings.servers" :key="server.id" class="flex items-center justify-between p-2 bg-white rounded border">
              <div>
                <span class="text-sm font-medium">{{ server.name }}</span>
                <span class="text-xs text-gray-500 mr-2">{{ server.url }}</span>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <span :class="getServerStatusClass(server.status)" class="px-2 py-1 text-xs rounded-full">
                  {{ getServerStatusText(server.status) }}
                </span>
                <button @click="removeServer(server.id)" class="text-red-500 hover:text-red-700">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                  </svg>
                </button>
              </div>
            </div>
            <button @click="addServer" class="w-full px-3 py-2 border border-dashed border-gray-300 rounded-lg text-gray-500 hover:border-gray-400 hover:text-gray-700">
              + افزودن سرور جدید
            </button>
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تنظیمات سلامت</label>
          <div class="space-y-2">
            <div>
              <label class="text-sm text-gray-600">فاصله بررسی سلامت (ثانیه)</label>
              <input v-model.number="loadBalancingSettings.healthCheck.interval" type="number" min="5" max="300" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500">
            </div>
            <div>
              <label class="text-sm text-gray-600">تایم‌اوت بررسی (ثانیه)</label>
              <input v-model.number="loadBalancingSettings.healthCheck.timeout" type="number" min="1" max="60" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500">
            </div>
            <label class="flex items-center">
              <input type="checkbox" v-model="loadBalancingSettings.healthCheck.autoRemove" class="rounded border-gray-300">
              <span class="text-sm text-gray-700 mr-2">حذف خودکار سرورهای غیرفعال</span>
            </label>
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تنظیمات پیشرفته</label>
          <div class="space-y-2">
            <label class="flex items-center">
              <input type="checkbox" v-model="loadBalancingSettings.advanced.sessionSticky" class="rounded border-gray-300">
              <span class="text-sm text-gray-700 mr-2">چسبندگی نشست</span>
            </label>
            <label class="flex items-center">
              <input type="checkbox" v-model="loadBalancingSettings.advanced.sslTermination" class="rounded border-gray-300">
              <span class="text-sm text-gray-700 mr-2">پایان SSL</span>
            </label>
            <label class="flex items-center">
              <input type="checkbox" v-model="loadBalancingSettings.advanced.rateLimiting" class="rounded border-gray-300">
              <span class="text-sm text-gray-700 mr-2">محدودیت نرخ</span>
            </label>
          </div>
        </div>
      </div>
    </div>

    <!-- Performance Monitoring -->
    <div class="bg-gray-50 rounded-xl p-6 mb-8">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">مانیتورینگ عملکرد</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <h4 class="font-medium text-gray-900 mb-3">آمار عملکرد</h4>
          <div class="space-y-3">
            <div class="flex justify-between items-center">
              <span class="text-sm text-gray-600">میانگین زمان پاسخ</span>
              <span class="text-sm font-medium">{{ performanceMonitoring.avgResponseTime }}ms</span>
            </div>
            <div class="flex justify-between items-center">
              <span class="text-sm text-gray-600">حداکثر زمان پاسخ</span>
              <span class="text-sm font-medium">{{ performanceMonitoring.maxResponseTime }}ms</span>
            </div>
            <div class="flex justify-between items-center">
              <span class="text-sm text-gray-600">درخواست‌های موفق</span>
              <span class="text-sm font-medium">{{ performanceMonitoring.successRate }}%</span>
            </div>
            <div class="flex justify-between items-center">
              <span class="text-sm text-gray-600">خطاهای ۵۰۰</span>
              <span class="text-sm font-medium text-red-600">{{ performanceMonitoring.error500Count }}</span>
            </div>
          </div>
        </div>

        <div>
          <h4 class="font-medium text-gray-900 mb-3">آخرین هشدارها</h4>
          <div class="space-y-2">
            <div v-for="alert in performanceMonitoring.alerts" :key="alert.id" class="flex items-center justify-between p-2 bg-white rounded border">
              <div>
                <span class="text-sm font-medium">{{ alert.message }}</span>
                <span class="text-xs text-gray-500 mr-2">{{ formatTime(alert.timestamp) }}</span>
              </div>
              <span :class="getAlertLevelClass(alert.level)" class="px-2 py-1 text-xs rounded-full">
                {{ getAlertLevelText(alert.level) }}
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Action Buttons -->
    <div class="flex justify-between">
      <button @click="runPerformanceTest" class="px-6 py-3 bg-emerald-600 text-white rounded-lg hover:bg-emerald-700 transition-colors">
        تست عملکرد
      </button>
      
      <div class="flex space-x-3 space-x-reverse">
        <button @click="clearCache" class="px-6 py-3 bg-orange-600 text-white rounded-lg hover:bg-orange-700 transition-colors">
          پاک کردن کش
        </button>
        <button @click="optimizePerformance" class="px-6 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors">
          بهینه‌سازی خودکار
        </button>
        <button @click="saveSettings" class="px-6 py-3 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors">
          ذخیره تنظیمات
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

interface PerformanceStats {
  loadTime: number
  loadTimeStatus: string
  cacheHitRate: number
  cacheSize: number
  compressionRatio: number
  savedBandwidth: number
  concurrentRequests: number
  avgResponseTime: number
}

interface CachingSettings {
  browserCache: {
    enabled: boolean
    duration: number
    strategy: string
  }
  serverCache: {
    enabled: boolean
    maxSize: number
    type: string
  }
  cdnCache: {
    enabled: boolean
    duration: number
    autoPurge: boolean
  }
  advanced: {
    compression: boolean
    encryption: boolean
    monitoring: boolean
  }
}

interface CompressionSettings {
  gzip: {
    enabled: boolean
    level: number
    minSize: number
  }
  brotli: {
    enabled: boolean
    level: number
    minSize: number
  }
  imageCompression: {
    enabled: boolean
    quality: number
    webp: boolean
  }
  advanced: {
    minifyCSS: boolean
    minifyJS: boolean
    minifyHTML: boolean
  }
}

interface LoadBalancingSettings {
  enabled: boolean
  algorithm: string
  servers: Array<{
    id: number
    name: string
    url: string
    status: 'active' | 'inactive' | 'error'
  }>
  healthCheck: {
    interval: number
    timeout: number
    autoRemove: boolean
  }
  advanced: {
    sessionSticky: boolean
    sslTermination: boolean
    rateLimiting: boolean
  }
}

interface PerformanceMonitoring {
  avgResponseTime: number
  maxResponseTime: number
  successRate: number
  error500Count: number
  alerts: Array<{
    id: number
    message: string
    level: 'info' | 'warning' | 'error'
    timestamp: Date
  }>
}

const performanceStats = ref<PerformanceStats>({
  loadTime: 1.2,
  loadTimeStatus: 'عالی',
  cacheHitRate: 87.5,
  cacheSize: 256,
  compressionRatio: 65.3,
  savedBandwidth: 1247,
  concurrentRequests: 45,
  avgResponseTime: 180
})

const cachingSettings = ref<CachingSettings>({
  browserCache: {
    enabled: true,
    duration: 24,
    strategy: 'moderate'
  },
  serverCache: {
    enabled: true,
    maxSize: 1024,
    type: 'redis'
  },
  cdnCache: {
    enabled: true,
    duration: 168,
    autoPurge: true
  },
  advanced: {
    compression: true,
    encryption: false,
    monitoring: true
  }
})

const compressionSettings = ref<CompressionSettings>({
  gzip: {
    enabled: true,
    level: 6,
    minSize: 1024
  },
  brotli: {
    enabled: true,
    level: 7,
    minSize: 1024
  },
  imageCompression: {
    enabled: true,
    quality: 85,
    webp: true
  },
  advanced: {
    minifyCSS: true,
    minifyJS: true,
    minifyHTML: true
  }
})

const loadBalancingSettings = ref<LoadBalancingSettings>({
  enabled: true,
  algorithm: 'round_robin',
  servers: [
    { id: 1, name: 'سرور اصلی', url: 'https://server1.example.com', status: 'active' },
    { id: 2, name: 'سرور پشتیبان', url: 'https://server2.example.com', status: 'active' },
    { id: 3, name: 'سرور تست', url: 'https://server3.example.com', status: 'inactive' }
  ],
  healthCheck: {
    interval: 30,
    timeout: 5,
    autoRemove: true
  },
  advanced: {
    sessionSticky: true,
    sslTermination: false,
    rateLimiting: true
  }
})

const performanceMonitoring = ref<PerformanceMonitoring>({
  avgResponseTime: 180,
  maxResponseTime: 2500,
  successRate: 99.8,
  error500Count: 3,
  alerts: [
    {
      id: 1,
      message: 'زمان پاسخ بالا در سرور ۲',
      level: 'warning',
      timestamp: new Date(Date.now() - 5 * 60 * 1000)
    },
    {
      id: 2,
      message: 'کش سرور پر شده است',
      level: 'info',
      timestamp: new Date(Date.now() - 15 * 60 * 1000)
    },
    {
      id: 3,
      message: 'خطای ۵۰۰ در سرور ۳',
      level: 'error',
      timestamp: new Date(Date.now() - 30 * 60 * 1000)
    }
  ]
})

// Methods
const getServerStatusClass = (status: string) => {
  switch (status) {
    case 'active':
      return 'bg-green-100 text-green-800'
    case 'inactive':
      return 'bg-gray-100 text-gray-800'
    case 'error':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getServerStatusText = (status: string) => {
  switch (status) {
    case 'active':
      return 'فعال'
    case 'inactive':
      return 'غیرفعال'
    case 'error':
      return 'خطا'
    default:
      return 'نامشخص'
  }
}

const getAlertLevelClass = (level: string) => {
  switch (level) {
    case 'info':
      return 'bg-blue-100 text-blue-800'
    case 'warning':
      return 'bg-yellow-100 text-yellow-800'
    case 'error':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getAlertLevelText = (level: string) => {
  switch (level) {
    case 'info':
      return 'اطلاعات'
    case 'warning':
      return 'هشدار'
    case 'error':
      return 'خطا'
    default:
      return 'نامشخص'
  }
}

const formatTime = (timestamp: Date) => {
  return new Intl.DateTimeFormat('fa-IR', {
    hour: '2-digit',
    minute: '2-digit'
  }).format(timestamp)
}

const addServer = () => {
  const newServer = {
    id: Date.now(),
    name: `سرور جدید ${loadBalancingSettings.value.servers.length + 1}`,
    url: 'https://newserver.example.com',
    status: 'inactive' as const
  }
  loadBalancingSettings.value.servers.push(newServer)
}

const removeServer = (serverId: number) => {
  const index = loadBalancingSettings.value.servers.findIndex(s => s.id === serverId)
  if (index > -1) {
    loadBalancingSettings.value.servers.splice(index, 1)
  }
}

const runPerformanceTest = () => {
  console.log('Running performance test...')
}

const clearCache = () => {
  if (confirm('آیا از پاک کردن تمام کش‌ها اطمینان دارید؟')) {
    console.log('Clearing all caches...')
  }
}

const optimizePerformance = () => {
  console.log('Running automatic performance optimization...')
}

const saveSettings = () => {
  console.log('Saving performance settings...')
}

// Expose methods for parent component
defineExpose({
  cachingSettings,
  compressionSettings,
  loadBalancingSettings,
  performanceStats,
  performanceMonitoring
})
</script>

<style scoped>
.performance-optimizer {
  transition: all 0.3s ease;
}

/* Hover effects */
.hover\:bg-gray-50:hover {
  background-color: #f9fafb;
}

/* Status badge animations */
.bg-green-100,
.bg-red-100,
.bg-yellow-100,
.bg-gray-100,
.bg-blue-100 {
  transition: all 0.2s ease;
}

/* Button hover effects */
button:hover {
  transform: translateY(-1px);
  transition: transform 0.2s ease;
}

/* Server status indicators */
.server-status {
  transition: all 0.3s ease;
}

/* Responsive adjustments */
@media (max-width: 768px) {
  .grid {
    grid-template-columns: 1fr;
  }
}
</style> 
