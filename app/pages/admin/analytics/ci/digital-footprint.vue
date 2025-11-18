<template>
  <div class="min-h-screen bg-gray-50" dir="rtl">
    <!-- Header -->
    <div class="bg-white shadow-sm border-b border-gray-200 mb-6">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center py-4">
          <div>
            <h1 class="text-2xl font-bold text-gray-900">پایش فعالیت‌های دیجیتال</h1>
            <p class="text-sm text-gray-600 mt-1">ردپای دیجیتال و تحلیل فعالیت‌های آنلاین</p>
          </div>
          <div class="flex items-center space-x-3 space-x-reverse">
            <button 
              @click="refreshData"
              class="inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-lg text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 shadow-md transition-all duration-200"
            >
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
              </svg>
              به‌روزرسانی
            </button>
            <button 
              @click="showSettings = true"
              class="inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-lg text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 shadow-md transition-all duration-200"
            >
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"></path>
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
              </svg>
              تنظیمات
            </button>
          </div>
        </div>
      </div>
    </div>

    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <!-- Tab Navigation -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 mb-6">
        <div class="border-b border-gray-200">
          <nav class="flex space-x-8 space-x-reverse px-6" aria-label="Tabs">
            <button
              v-for="tab in tabs"
              :key="tab.id"
              @click="activeTab = tab.id"
              :class="[
                activeTab === tab.id
                  ? 'border-blue-500 text-blue-600'
                  : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300',
                'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm transition-colors duration-200'
              ]"
            >
              <div class="flex items-center space-x-2 space-x-reverse">
                <component :is="tab.icon" class="w-5 h-5" />
                <span>{{ tab.name }}</span>
                <span v-if="tab.badge" class="ml-2 inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800">
                  {{ tab.badge }}
                </span>
              </div>
            </button>
          </nav>
        </div>
      </div>

      <!-- Tab Content -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200">
        <div class="p-6">
          <!-- Website Traffic Tab -->
          <div v-if="activeTab === 'traffic'" class="space-y-6">
            <WebsiteTrafficMonitoring />
          </div>

          <!-- SEO Analysis Tab -->
          <div v-if="activeTab === 'seo'" class="space-y-6">
            <SeoAnalysis />
          </div>

          <!-- Digital Advertising Tab -->
          <div v-if="activeTab === 'advertising'" class="space-y-6">
            <DigitalAdvertising />
          </div>

          <!-- Social Media Tab -->
          <div v-if="activeTab === 'social'" class="space-y-6">
            <SocialMediaMonitoring />
          </div>

          <!-- Content Analysis Tab -->
          <div v-if="activeTab === 'content'" class="space-y-6">
            <ContentAnalysis />
          </div>
        </div>
      </div>
    </div>

    <!-- Settings Modal -->
    <div v-if="showSettings" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
      <div class="relative top-20 mx-auto p-5 border w-96 shadow-lg rounded-md bg-white">
        <div class="mt-3 text-center">
          <h3 class="text-lg font-medium text-gray-900 mb-4">تنظیمات پایش دیجیتال</h3>
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">فاصله به‌روزرسانی (دقیقه)</label>
              <input 
                v-model="updateInterval" 
                type="number" 
                min="1" 
                max="60"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نمایش اعلان‌ها</label>
              <div class="flex items-center space-x-4 space-x-reverse">
                <label class="flex items-center">
                  <input type="checkbox" v-model="showNotifications" class="ml-2">
                  <span class="text-sm text-gray-700">فعال</span>
                </label>
              </div>
            </div>
          </div>
          <div class="flex justify-center space-x-3 space-x-reverse mt-6">
            <button 
              @click="saveSettings"
              class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500"
            >
              ذخیره
            </button>
            <button 
              @click="showSettings = false"
              class="px-4 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400 focus:outline-none focus:ring-2 focus:ring-gray-500"
            >
              انصراف
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, defineComponent } from 'vue'
import WebsiteTrafficMonitoring from './components/WebsiteTrafficMonitoring.vue'
import SeoAnalysis from './components/SeoAnalysis.vue'
import DigitalAdvertising from './components/DigitalAdvertising.vue'
import SocialMediaMonitoring from './components/SocialMediaMonitoring.vue'
import ContentAnalysis from './components/ContentAnalysis.vue'

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user, hasPermission } = useAuth()

// آیکون‌های تب‌ها
const TrafficIcon = defineComponent({
  template: `
    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
    </svg>
  `
})

const SeoIcon = defineComponent({
  template: `
    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
    </svg>
  `
})

const AdvertisingIcon = defineComponent({
  template: `
    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 3.055A9.001 9.001 0 1020.945 13H11V3.055z"></path>
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.488 9H15V3.512A9.025 9.025 0 0120.488 9z"></path>
    </svg>
  `
})

const SocialIcon = defineComponent({
  template: `
    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 8h10M7 12h4m1 8l-4-4H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-3l-4 4z"></path>
    </svg>
  `
})

const ContentIcon = defineComponent({
  template: `
    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
      
    </svg>
  `
})

// تعریف تب‌های مختلف
const tabs = [
  {
    id: 'traffic',
    name: 'ترافیک وب‌سایت',
    icon: TrafficIcon,
    badge: null
  },
  {
    id: 'seo',
    name: 'تحلیل SEO',
    icon: SeoIcon,
    badge: 'جدید'
  },
  {
    id: 'advertising',
    name: 'تبلیغات دیجیتال',
    icon: AdvertisingIcon,
    badge: null
  },
  {
    id: 'social',
    name: 'شبکه‌های اجتماعی',
    icon: SocialIcon,
    badge: null
  },
  {
    id: 'content',
    name: 'تحلیل محتوا',
    icon: ContentIcon,
    badge: null
  }
]

// متغیرهای reactive
const activeTab = ref('traffic')
const showSettings = ref(false)
const updateInterval = ref(15)
const showNotifications = ref(true)

// توابع
const refreshData = () => {
  // به‌روزرسانی داده‌ها
  console.log('Refreshing digital footprint data...')
}

const saveSettings = () => {
  // ذخیره تنظیمات
  console.log('Saving settings...', { updateInterval: updateInterval.value, showNotifications: showNotifications.value })
  showSettings.value = false
}
</script>

