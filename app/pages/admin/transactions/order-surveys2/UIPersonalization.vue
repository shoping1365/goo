<template>
  <div class="space-y-8">
    <!-- Header -->
    <div>
      <h3 class="text-lg font-semibold text-gray-800">شخصی‌سازی رابط کاربری</h3>
      <p class="text-gray-600 text-sm">تنظیم تم، زبان و چیدمان داشبورد</p>
    </div>

    <!-- Theme Settings -->
    <div class="bg-white rounded-lg border border-gray-200 px-4 py-4">
      <h4 class="text-md font-semibold text-gray-800 mb-4">تنظیمات تم</h4>
      <div class="grid grid-cols-1 md:grid-cols-3 gapx-4 py-4">
        <div 
          v-for="theme in themes" 
          :key="theme.id"
          @click="selectTheme(theme.id)"
          :class="[
            'border-2 rounded-lg px-4 py-4 cursor-pointer transition-all',
            selectedTheme === theme.id 
              ? 'border-blue-500 bg-blue-50' 
              : 'border-gray-200 hover:border-gray-300'
          ]"
        >
          <div class="flex items-center justify-between mb-3">
            <h5 class="font-medium text-gray-900">{{ theme.name }}</h5>
            <div v-if="selectedTheme === theme.id" class="w-5 h-5 bg-blue-500 rounded-full flex items-center justify-center">
              <svg class="w-3 h-3 text-white" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd"></path>
              </svg>
            </div>
          </div>
          <div class="text-sm text-gray-600 mb-3">{{ theme.description }}</div>
          <div class="flex items-center space-x-2 space-x-reverse">
            <div :class="theme.preview" class="w-4 h-4 rounded border"></div>
            <span class="text-xs text-gray-500">پیش‌نمایش</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Language Settings -->
    <div class="bg-white rounded-lg border border-gray-200 px-4 py-4">
      <h4 class="text-md font-semibold text-gray-800 mb-4">تنظیمات زبان</h4>
      <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">زبان رابط کاربری</label>
          <select v-model="selectedLanguage" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
            <option v-for="lang in languages" :key="lang.code" :value="lang.code">
              {{ lang.name }} ({{ lang.nativeName }})
            </option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">منطقه زمانی</label>
          <select v-model="selectedTimezone" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
            <option v-for="tz in timezones" :key="tz.value" :value="tz.value">
              {{ tz.label }}
            </option>
          </select>
        </div>
      </div>
    </div>

    <!-- Dashboard Customization -->
    <div class="bg-white rounded-lg border border-gray-200 px-4 py-4">
      <h4 class="text-md font-semibold text-gray-800 mb-4">شخصی‌سازی داشبورد</h4>
      <div class="space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">ویجت‌های نمایش داده شده</label>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-3">
            <div v-for="widget in availableWidgets" :key="widget.id" class="flex items-center">
              <input 
                :id="widget.id"
                type="checkbox" 
                v-model="widget.enabled"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              >
              <label :for="widget.id" class="mr-2 text-sm text-gray-700">{{ widget.name }}</label>
            </div>
          </div>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">ترتیب نمایش</label>
          <div class="grid grid-cols-1 md:grid-cols-3 gapx-4 py-4">
            <div 
              v-for="layout in layouts" 
              :key="layout.id"
              @click="selectLayout(layout.id)"
              :class="[
                'border-2 rounded-lg p-3 cursor-pointer transition-all',
                selectedLayout === layout.id 
                  ? 'border-blue-500 bg-blue-50' 
                  : 'border-gray-200 hover:border-gray-300'
              ]"
            >
              <div class="text-sm font-medium text-gray-900 mb-2">{{ layout.name }}</div>
              <div class="text-xs text-gray-500">{{ layout.description }}</div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Display Settings -->
    <div class="bg-white rounded-lg border border-gray-200 px-4 py-4">
      <h4 class="text-md font-semibold text-gray-800 mb-4">تنظیمات نمایش</h4>
      <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">اندازه فونت</label>
          <select v-model="fontSize" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
            <option value="small">کوچک</option>
            <option value="medium">متوسط</option>
            <option value="large">بزرگ</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تراکم محتوا</label>
          <select v-model="density" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
            <option value="compact">فشرده</option>
            <option value="comfortable">راحت</option>
            <option value="spacious">گسترده</option>
          </select>
        </div>
      </div>
    </div>

    <!-- Save Button -->
    <div class="flex items-center justify-end">
      <button 
        @click="saveSettings"
        :disabled="saving"
        class="px-6 py-2 bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 text-white rounded-lg text-sm transition-colors flex items-center space-x-2 space-x-reverse"
      >
        <svg v-if="saving" class="animate-spin w-4 h-4" fill="none" viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
        </svg>
        <span>{{ saving ? 'در حال ذخیره...' : 'ذخیره تنظیمات' }}</span>
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const selectedTheme = ref('light')
const selectedLanguage = ref('fa')
const selectedTimezone = ref('Asia/Tehran')
const selectedLayout = ref('default')
const fontSize = ref('medium')
const density = ref('comfortable')
const saving = ref(false)

const themes = ref([
  {
    id: 'light',
    name: 'تم روشن',
    description: 'پیش‌فرض - مناسب برای استفاده در روز',
    preview: 'bg-white border-gray-300'
  },
  {
    id: 'dark',
    name: 'تم تاریک',
    description: 'مناسب برای استفاده در شب و کاهش خستگی چشم',
    preview: 'bg-gray-800 border-gray-600'
  },
  {
    id: 'auto',
    name: 'خودکار',
    description: 'تغییر خودکار بر اساس تنظیمات سیستم',
    preview: 'bg-gradient-to-r from-white to-gray-800 border-gray-400'
  }
])

const languages = ref([
  { code: 'fa', name: 'Persian', nativeName: 'فارسی' },
  { code: 'en', name: 'English', nativeName: 'English' },
  { code: 'ar', name: 'Arabic', nativeName: 'العربية' }
])

const timezones = ref([
  { value: 'Asia/Tehran', label: 'تهران (UTC+3:30)' },
  { value: 'UTC', label: 'UTC (UTC+0)' },
  { value: 'America/New_York', label: 'نیویورک (UTC-5)' },
  { value: 'Europe/London', label: 'لندن (UTC+0)' }
])

const availableWidgets = ref([
  { id: 'stats', name: 'آمار کلی', enabled: true },
  { id: 'recent-orders', name: 'سفارشات اخیر', enabled: true },
  { id: 'sms-status', name: 'وضعیت پیامک‌ها', enabled: true },
  { id: 'responses', name: 'پاسخ‌های جدید', enabled: false },
  { id: 'errors', name: 'خطاهای ارسال', enabled: true },
  { id: 'charts', name: 'نمودارها', enabled: true }
])

const layouts = ref([
  {
    id: 'default',
    name: 'پیش‌فرض',
    description: 'چیدمان استاندارد 3 ستونه'
  },
  {
    id: 'compact',
    name: 'فشرده',
    description: 'چیدمان 4 ستونه برای نمایش بیشتر'
  },
  {
    id: 'wide',
    name: 'گسترده',
    description: 'چیدمان 2 ستونه با فضای بیشتر'
  }
])

const selectTheme = (themeId: string) => {
  selectedTheme.value = themeId
}

const selectLayout = (layoutId: string) => {
  selectedLayout.value = layoutId
}

const saveSettings = async () => {
  saving.value = true
  try {
    await new Promise(resolve => setTimeout(resolve, 1500))
    alert('تنظیمات شخصی‌سازی با موفقیت ذخیره شد!')
  } finally {
    saving.value = false
  }
}
</script>

<style scoped>
.space-x-reverse > :not([hidden]) ~ :not([hidden]) {
  --tw-space-x-reverse: 1;
}
</style> 
