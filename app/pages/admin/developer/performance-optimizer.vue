<template>
  <div class="p-6">
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900 mb-2">بهینه‌ساز عملکرد</h1>
      <p class="text-gray-600">تنظیمات بهینه‌سازی سایت، کش و عملکرد</p>
    </div>

    <!-- دکمه‌های اصلی -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6">
      <h2 class="text-lg font-semibold text-gray-900 mb-4">عملیات سریع</h2>
      <div class="flex flex-wrap gap-6">
        <button 
          :disabled="isClearing"
          class="bg-red-600 hover:bg-red-700 text-white font-medium py-3 px-6 rounded-lg transition-colors disabled:opacity-50"
          @click="clearCache"
        >
          {{ isClearing ? 'در حال پاک کردن...' : 'پاک کردن کش' }}
        </button>
        
        <button 
          :disabled="isReloading"
          class="bg-blue-600 hover:bg-blue-700 text-white font-medium py-3 px-6 rounded-lg transition-colors disabled:opacity-50"
          @click="reloadServer"
        >
          {{ isReloading ? 'در حال ری‌استارت...' : 'ری‌استارت سرور' }}
        </button>
        
        <button 
          :disabled="isAnalyzing"
          class="bg-green-600 hover:bg-green-700 text-white font-medium py-3 px-6 rounded-lg transition-colors disabled:opacity-50"
          @click="analyzeBundle"
        >
          {{ isAnalyzing ? 'در حال تحلیل...' : 'تحلیل باندل' }}
        </button>
      </div>
    </div>

    <!-- تنظیمات کش -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6">
      <div class="flex items-center justify-between mb-4">
        <h2 class="text-lg font-semibold text-gray-900">تنظیمات کش</h2>
        <div class="flex items-center">
          <input 
            id="cache-enabled" 
            v-model="settings.cache.enabled" 
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <label for="cache-enabled" class="mr-2 text-sm font-medium text-gray-700">فعال</label>
        </div>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            مدت زمان کش استاتیک (ثانیه)
            <span class="text-gray-500 text-xs">- فایل‌های CSS, JS, تصاویر</span>
          </label>
          <input 
            v-model="settings.cache.staticMaxAge" 
            type="number" 
            min="0"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
          <p class="text-xs text-gray-500 mt-1">پیشنهادی: 31536000 (یک سال)</p>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            مدت زمان کش API (ثانیه)
            <span class="text-gray-500 text-xs">- درخواست‌های API</span>
          </label>
          <input 
            v-model="settings.cache.apiMaxAge" 
            type="number" 
            min="0"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
          <p class="text-xs text-gray-500 mt-1">پیشنهادی: 0 (بدون کش برای API)</p>
        </div>
      </div>
    </div>

    <!-- تنظیمات فشردگی -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6">
      <div class="flex items-center justify-between mb-4">
        <h2 class="text-lg font-semibold text-gray-900">فشردگی</h2>
        <div class="flex items-center">
          <input 
            id="compression-enabled" 
            v-model="settings.compression.enabled" 
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <label for="compression-enabled" class="mr-2 text-sm font-medium text-gray-700">فعال</label>
        </div>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            نوع فشردگی
            <span class="text-gray-500 text-xs">- Brotli بهتر از Gzip</span>
          </label>
          <select 
            v-model="settings.compression.algorithm"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            <option value="brotliCompress">Brotli (پیشنهادی)</option>
            <option value="gzip">Gzip</option>
          </select>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            حداقل سایز فایل (KB)
            <span class="text-gray-500 text-xs">- فقط فایل‌های بزرگتر فشرده شوند</span>
          </label>
          <input 
            v-model="settings.compression.threshold" 
            type="number" 
            min="1"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
          <p class="text-xs text-gray-500 mt-1">پیشنهادی: 20 (20KB)</p>
        </div>
      </div>
    </div>

    <!-- تنظیمات Lazy Loading -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6">
      <div class="flex items-center justify-between mb-4">
        <h2 class="text-lg font-semibold text-gray-900">بارگذاری تنبل (Lazy Loading)</h2>
        <div class="flex items-center">
          <input 
            id="lazy-enabled" 
            v-model="settings.lazyLoading.enabled" 
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <label for="lazy-enabled" class="mr-2 text-sm font-medium text-gray-700">فعال</label>
        </div>
      </div>
      
      <div class="space-y-6">
        <!-- تنظیمات پیش‌فرض -->
        <div class="space-y-4">
          <div class="flex items-center">
            <input 
              id="lazy-admin" 
              v-model="settings.lazyLoading.adminComponents" 
              type="checkbox"
              class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
            />
            <label for="lazy-admin" class="mr-3 text-sm font-medium text-gray-700">
              کامپوننت‌های ادمین
              <span class="text-gray-500 text-xs block">- بارگذاری کامپوننت‌های بخش ادمین فقط در صورت نیاز</span>
            </label>
          </div>
          
          <div class="flex items-center">
            <input 
              id="lazy-libs" 
              v-model="settings.lazyLoading.heavyLibraries" 
              type="checkbox"
              class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
            />
            <label for="lazy-libs" class="mr-3 text-sm font-medium text-gray-700">
              کتابخانه‌های سنگین
              <span class="text-gray-500 text-xs block">- TinyMCE، Chart.js، Swiper و...</span>
            </label>
          </div>
          
          <div class="flex items-center">
            <input 
              id="lazy-images" 
              v-model="settings.lazyLoading.images" 
              type="checkbox"
              class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
            />
            <label for="lazy-images" class="mr-3 text-sm font-medium text-gray-700">
              تصاویر
              <span class="text-gray-500 text-xs block">- استفاده از Nuxt Image با WebP/AVIF</span>
            </label>
          </div>
        </div>

        <!-- تنظیمات lazy loading دلخواه -->
        <div>
          <h4 class="font-medium text-gray-900 mb-3">کامپوننت‌های دلخواه</h4>
          
          <!-- لیست کامپوننت های lazy -->
          <div class="bg-gray-50 p-6 rounded-md mb-4">
            <h5 class="font-medium text-gray-800 mb-2">کامپوننت‌های Lazy:</h5>
            <div class="space-y-2">
              <div
v-for="(component, index) in settings.lazyLoading.customComponents" :key="index" 
                   class="flex items-center justify-between bg-white p-2 rounded border">
                <div class="flex-1">
                  <span class="font-medium text-sm">{{ component.name }}</span>
                  <span class="text-gray-500 text-xs block">{{ component.path }}</span>
                </div>
                <div class="flex items-center space-x-2 space-x-reverse">
                  <span :class="component.enabled ? 'text-green-600' : 'text-gray-400'" class="text-xs">
                    {{ component.enabled ? 'فعال' : 'غیرفعال' }}
                  </span>
                  <button 
                    :class="component.enabled ? 'bg-green-600 hover:bg-green-700' : 'bg-gray-400 hover:bg-gray-500'"
                    class="text-white px-2 py-1 rounded text-xs"
                    @click="toggleLazyComponent(index)"
                  >
                    {{ component.enabled ? 'غیرفعال' : 'فعال' }}
                  </button>
                  <button 
                    class="text-red-600 hover:text-red-800 text-xs"
                    @click="removeLazyComponent(index)"
                  >
                    حذف
                  </button>
                </div>
              </div>
            </div>
          </div>

          <!-- افزودن کامپوننت جدید -->
          <div class="space-y-2">
            <div class="flex gap-2">
              <input 
                v-model="newLazyComponent.name"
                type="text" 
                placeholder="نام کامپوننت"
                class="flex-1 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 text-sm"
              />
              <input 
                v-model="newLazyComponent.path"
                type="text" 
                placeholder="مسیر کامپوننت"
                class="flex-1 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 text-sm"
              />
              <button 
                class="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-md text-sm"
                @click="addLazyComponent"
              >
                افزودن
              </button>
            </div>
            <p class="text-xs text-gray-500">مثال: ProductEditor | /components/admin/ProductEditor.vue</p>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات Bundle -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6">
      <h2 class="text-lg font-semibold text-gray-900 mb-4">تقسیم‌بندی باندل</h2>
      
      <div class="space-y-6">
        <!-- تنظیمات پیش‌فرض -->
        <div class="space-y-4">
          <div class="flex items-center">
            <input 
              id="vendor-chunk" 
              v-model="settings.bundleSplitting.vendorChunk" 
              type="checkbox"
              class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
            />
            <label for="vendor-chunk" class="mr-3 text-sm font-medium text-gray-700">
              جداسازی کتابخانه‌ها (Vendor Chunk)
              <span class="text-gray-500 text-xs block">- تمام پکیج‌های node_modules در فایل جداگانه</span>
            </label>
          </div>
          
          <div class="flex items-center">
            <input 
              id="admin-chunk" 
              v-model="settings.bundleSplitting.adminChunk" 
              type="checkbox"
              class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
            />
            <label for="admin-chunk" class="mr-3 text-sm font-medium text-gray-700">
              جداسازی بخش ادمین
              <span class="text-gray-500 text-xs block">- کدهای ادمین در فایل جداگانه تا سایت اصلی سبک‌تر باشد</span>
            </label>
          </div>
          
          <div class="flex items-center">
            <input 
              id="editor-chunk" 
              v-model="settings.bundleSplitting.editorChunk" 
              type="checkbox"
              class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
            />
            <label for="editor-chunk" class="mr-3 text-sm font-medium text-gray-700">
              جداسازی ویرایشگر
              <span class="text-gray-500 text-xs block">- TinyMCE و ویرایشگرهای متن در فایل جداگانه</span>
            </label>
          </div>
        </div>

        <!-- تنظیمات chunk های دلخواه -->
        <div>
          <h4 class="font-medium text-gray-900 mb-3">Chunk های دلخواه</h4>
          
          <!-- لیست chunk های موجود -->
          <div class="bg-gray-50 p-6 rounded-md mb-4">
            <h5 class="font-medium text-gray-800 mb-2">Chunk های تنظیم شده:</h5>
            <div class="space-y-2">
              <div
v-for="(chunk, index) in settings.bundleSplitting.customChunks" :key="index" 
                   class="flex items-center justify-between bg-white p-2 rounded border">
                <div class="flex-1">
                  <span class="font-medium text-sm">{{ chunk.name }}</span>
                  <span class="text-gray-500 text-xs block">{{ chunk.pattern }}</span>
                </div>
                <button 
                  class="text-red-600 hover:text-red-800 text-xs"
                  @click="removeCustomChunk(index)"
                >
                  حذف
                </button>
              </div>
            </div>
          </div>

          <!-- افزودن chunk جدید -->
          <div class="space-y-2">
            <div class="flex gap-2">
              <input 
                v-model="newChunk.name"
                type="text" 
                placeholder="نام chunk (مثل: charts)"
                class="flex-1 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 text-sm"
              />
              <input 
                v-model="newChunk.pattern"
                type="text" 
                placeholder="الگو (مثل: chart.js)"
                class="flex-1 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 text-sm"
              />
              <button 
                class="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-md text-sm"
                @click="addCustomChunk"
              >
                افزودن
              </button>
            </div>
            <p class="text-xs text-gray-500">مثال: charts | chart.js یا maps | leaflet</p>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات SSR -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6">
      <h2 class="text-lg font-semibold text-gray-900 mb-4">رندرینگ سمت سرور (SSR)</h2>
      
      <div class="space-y-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حالت پیش‌فرض</label>
          <select 
            v-model="settings.ssr.defaultMode"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            <option value="ssr">SSR (رندر در سرور)</option>
            <option value="spa">SPA (رندر در کلاینت)</option>
            <option value="hybrid">Hybrid (ترکیبی)</option>
          </select>
          <p class="text-xs text-gray-500 mt-1">SSR برای SEO بهتر، SPA برای UX بهتر پس از بارگذاری</p>
        </div>

        <!-- تنظیمات صفحات خاص -->
        <div>
          <h4 class="font-medium text-gray-900 mb-3">تنظیمات صفحات خاص</h4>
          
          <!-- لیست صفحات موجود -->
          <div class="bg-gray-50 p-6 rounded-md mb-4">
            <h5 class="font-medium text-gray-800 mb-2">صفحات تنظیم شده:</h5>
            <div class="space-y-2">
              <div
v-for="(route, index) in settings.ssr.customRoutes" :key="index" 
                   class="flex items-center justify-between bg-white p-2 rounded border">
                <div class="flex-1">
                  <span class="font-mono text-sm">{{ route.path }}</span>
                </div>
                <div class="flex items-center space-x-2 space-x-reverse">
                  <select 
                    v-model="route.mode"
                    class="text-xs border border-gray-300 rounded px-2 py-1"
                  >
                    <option value="ssr">SSR</option>
                    <option value="spa">SPA</option>
                  </select>
                  <button 
                    class="text-red-600 hover:text-red-800 text-xs"
                    @click="removeCustomRoute(index)"
                  >
                    حذف
                  </button>
                </div>
              </div>
            </div>
          </div>

          <!-- افزودن صفحه جدید -->
          <div class="flex gap-2">
            <input 
              v-model="newRoute.path"
              type="text" 
              placeholder="/path/to/page یا /path/**"
              class="flex-1 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 text-sm"
            />
            <select 
              v-model="newRoute.mode"
              class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 text-sm"
            >
              <option value="ssr">SSR</option>
              <option value="spa">SPA</option>
            </select>
            <button 
              class="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-md text-sm"
              @click="addCustomRoute"
            >
              افزودن
            </button>
          </div>
          <p class="text-xs text-gray-500 mt-1">مثال: /blog/**, /admin/products, /user/profile</p>
        </div>
      </div>
    </div>

    <!-- دکمه ذخیره -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex justify-between items-center">
        <div>
          <h3 class="text-lg font-semibold text-gray-900">اعمال تنظیمات</h3>
          <p class="text-sm text-gray-600">تغییرات پس از ذخیره و ری‌استارت سرور اعمال می‌شود</p>
        </div>
        
        <div class="flex gap-3">
          <button 
            class="bg-gray-100 hover:bg-gray-200 text-gray-700 font-medium py-2 px-4 rounded-lg transition-colors"
            @click="resetSettings"
          >
            بازگردانی پیش‌فرض
          </button>
          
          <button 
            :disabled="isSaving"
            class="bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-6 rounded-lg transition-colors disabled:opacity-50"
            @click="saveSettings"
          >
            {{ isSaving ? 'در حال ذخیره...' : 'ذخیره تنظیمات' }}
          </button>
        </div>
      </div>
    </div>

    <!-- آمار عملکرد -->
    <div v-if="stats" class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mt-6">
      <h2 class="text-lg font-semibold text-gray-900 mb-4">آمار عملکرد</h2>
      
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div class="bg-green-50 p-6 rounded-lg">
          <div class="text-2xl font-bold text-green-600">{{ stats.bundleSize }}</div>
          <div class="text-sm text-green-700">حجم باندل اولیه</div>
        </div>
        
        <div class="bg-blue-50 p-6 rounded-lg">
          <div class="text-2xl font-bold text-blue-600">{{ stats.loadTime }}</div>
          <div class="text-sm text-blue-700">زمان بارگذاری</div>
        </div>
        
        <div class="bg-purple-50 p-6 rounded-lg">
          <div class="text-2xl font-bold text-purple-600">{{ stats.cacheHitRate }}</div>
          <div class="text-sm text-purple-700">نرخ بازدید کش</div>
        </div>
        
        <div class="bg-orange-50 p-6 rounded-lg">
          <div class="text-2xl font-bold text-orange-600">{{ stats.compressionRatio }}</div>
          <div class="text-sm text-orange-700">نسبت فشردگی</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { onMounted, reactive, ref } from 'vue'

definePageMeta({
  layout: 'admin-main',
  middleware: ['developer-only']
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
// const { user, hasPermission } = useAuth()

// Loading states
const isClearing = ref(false)
const isReloading = ref(false)
const isAnalyzing = ref(false)
const isSaving = ref(false)

// Settings object
const settings = reactive({
  cache: {
    enabled: true,
    staticMaxAge: 31536000, // 1 year
    apiMaxAge: 0
  },
  compression: {
    enabled: true,
    algorithm: 'brotliCompress',
    threshold: 20 // KB
  },
  lazyLoading: {
    enabled: true,
    adminComponents: true,
    heavyLibraries: true,
    images: true,
    customComponents: [
      { name: 'TinyMCE Editor', path: '/components/admin/TinyMCEEditor.vue', enabled: true },
      { name: 'Chart Component', path: '/components/admin/ChartComponent.vue', enabled: true }
    ]
  },
  bundleSplitting: {
    vendorChunk: true,
    adminChunk: true,
    editorChunk: true,
    customChunks: [
      { name: 'charts', pattern: 'chart.js' },
      { name: 'editor', pattern: '@tinymce' }
    ]
  },
  ssr: {
    defaultMode: 'hybrid',
    customRoutes: [
      { path: '/admin/**', mode: 'spa' },
      { path: '/auth/**', mode: 'spa' },
      { path: '/cart/**', mode: 'spa' },
      { path: '/product/**', mode: 'ssr' },
      { path: '/blog/**', mode: 'ssr' }
    ]
  }
})

// New items for adding
const newRoute = ref({ path: '', mode: 'ssr' })
const newChunk = ref({ name: '', pattern: '' })
const newLazyComponent = ref({ name: '', path: '' })

// Stats
const stats = ref(null)

// Load current settings
onMounted(async () => {
  await loadSettings()
  await loadStats()
})

const loadSettings = async () => {
  try {
    const response = await $fetch('/api/admin/performance/settings')
    if (response.success) {
      Object.assign(settings, response.data)
    }
  } catch (error) {
    console.error('خطا در بارگذاری تنظیمات:', error)
  }
}

const loadStats = async () => {
  try {
    const response = await $fetch('/api/admin/performance/stats')
    if (response.success) {
      stats.value = response.data
    }
  } catch (error) {
    console.error('خطا در بارگذاری آمار:', error)
  }
}

const clearCache = async () => {
  isClearing.value = true
  try {
    const response = await $fetch('/api/admin/performance/clear-cache', {
      method: 'POST'
    })
    
    if (response.success) {
      alert('کش با موفقیت پاک شد!')
      await loadStats()
    } else {
      alert('خطا در پاک کردن کش: ' + response.message)
    }
  } catch (error) {
    alert('خطا در پاک کردن کش: ' + error.message)
  } finally {
    isClearing.value = false
  }
}

const reloadServer = async () => {
  isReloading.value = true
  try {
    const response = await $fetch('/api/admin/performance/reload-server', {
      method: 'POST'
    })
    
    if (response.success) {
      alert('سرور در حال ری‌استارت است...')
      // Reload page after 3 seconds
      setTimeout(() => {
        window.location.reload()
      }, 3000)
    } else {
      alert('خطا در ری‌استارت سرور: ' + response.message)
    }
  } catch (error) {
    alert('خطا در ری‌استارت سرور: ' + error.message)
  } finally {
    isReloading.value = false
  }
}

const analyzeBundle = async () => {
  isAnalyzing.value = true
  try {
    const response = await $fetch('/api/admin/performance/analyze-bundle', {
      method: 'POST'
    })
    
    if (response.success) {
      // Open bundle analyzer in new tab
      window.open('/bundle-analyzer', '_blank')
    } else {
      alert('خطا در تحلیل باندل: ' + response.message)
    }
  } catch (error) {
    alert('خطا در تحلیل باندل: ' + error.message)
  } finally {
    isAnalyzing.value = false
  }
}

const saveSettings = async () => {
  isSaving.value = true
  try {
    const response = await $fetch('/api/admin/performance/settings', {
      method: 'PUT',
      body: settings
    })
    
    if (response.success) {
      alert('تنظیمات ذخیره شد! برای اعمال تغییرات سرور را ری‌استارت کنید.')
    } else {
      alert('خطا در ذخیره تنظیمات: ' + response.message)
    }
  } catch (error) {
    alert('خطا در ذخیره تنظیمات: ' + error.message)
  } finally {
    isSaving.value = false
  }
}

const resetSettings = () => {
  if (confirm('آیا مطمئن هستید که می‌خواهید تنظیمات را به حالت پیش‌فرض بازگردانید؟')) {
    // Reset to default values
    Object.assign(settings, {
      cache: {
        enabled: true,
        staticMaxAge: 31536000,
        apiMaxAge: 0
      },
      compression: {
        enabled: true,
        algorithm: 'brotliCompress',
        threshold: 20
      },
      lazyLoading: {
        enabled: true,
        adminComponents: true,
        heavyLibraries: true,
        images: true,
        customComponents: []
      },
      bundleSplitting: {
        vendorChunk: true,
        adminChunk: true,
        editorChunk: true,
        customChunks: []
      },
      ssr: {
        defaultMode: 'hybrid',
        customRoutes: []
      }
    })
  }
}

// SSR Route Management
const addCustomRoute = () => {
  if (newRoute.value.path.trim()) {
    settings.ssr.customRoutes.push({
      path: newRoute.value.path.trim(),
      mode: newRoute.value.mode
    })
    newRoute.value = { path: '', mode: 'ssr' }
  }
}

const removeCustomRoute = (index) => {
  settings.ssr.customRoutes.splice(index, 1)
}

// Bundle Chunk Management
const addCustomChunk = () => {
  if (newChunk.value.name.trim() && newChunk.value.pattern.trim()) {
    settings.bundleSplitting.customChunks.push({
      name: newChunk.value.name.trim(),
      pattern: newChunk.value.pattern.trim()
    })
    newChunk.value = { name: '', pattern: '' }
  }
}

const removeCustomChunk = (index) => {
  settings.bundleSplitting.customChunks.splice(index, 1)
}

// Lazy Component Management
const addLazyComponent = () => {
  if (newLazyComponent.value.name.trim() && newLazyComponent.value.path.trim()) {
    settings.lazyLoading.customComponents.push({
      name: newLazyComponent.value.name.trim(),
      path: newLazyComponent.value.path.trim(),
      enabled: true
    })
    newLazyComponent.value = { name: '', path: '' }
  }
}

const removeLazyComponent = (index) => {
  settings.lazyLoading.customComponents.splice(index, 1)
}

const toggleLazyComponent = (index) => {
  settings.lazyLoading.customComponents[index].enabled = !settings.lazyLoading.customComponents[index].enabled
}
</script>
