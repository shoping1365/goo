<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6">
    <div class="flex items-center justify-between mb-6">
      <h3 class="text-lg font-semibold text-gray-900">یکپارچه‌سازی</h3>
      <button 
        @click="refreshIntegrations" 
        class="px-3 py-2 text-sm bg-blue-50 text-blue-700 rounded-lg hover:bg-blue-100 flex items-center"
      >
        <svg class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
        </svg>
        به‌روزرسانی
      </button>
    </div>

    <!-- Google Analytics -->
    <div class="mb-8">
      <div class="flex items-center justify-between mb-4">
        <div class="flex items-center">
          <div class="w-10 h-10 bg-blue-500 rounded-lg flex items-center justify-center mr-3">
            <svg class="w-6 h-6 text-white" fill="currentColor" viewBox="0 0 24 24">
              <path d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"/>
              <path d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"/>
              <path d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"/>
              <path d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"/>
            </svg>
          </div>
          <div>
            <h4 class="text-md font-medium text-gray-900">Google Analytics</h4>
            <p class="text-sm text-gray-500">پیگیری نتایج تست‌ها در Google Analytics</p>
          </div>
        </div>
        <div class="flex items-center space-x-2 space-x-reverse">
          <span :class="[
            'px-2 py-1 text-xs rounded-full',
            googleAnalytics.connected ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'
          ]">
            {{ googleAnalytics.connected ? 'متصل' : 'قطع' }}
          </span>
          <button 
            @click="toggleGoogleAnalytics"
            :class="[
              'px-3 py-1 text-sm rounded-lg',
              googleAnalytics.connected 
                ? 'bg-red-50 text-red-700 hover:bg-red-100' 
                : 'bg-green-50 text-green-700 hover:bg-green-100'
            ]"
          >
            {{ googleAnalytics.connected ? 'قطع اتصال' : 'اتصال' }}
          </button>
        </div>
      </div>

      <div v-if="googleAnalytics.connected" class="bg-gray-50 rounded-lg p-6">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Property ID</label>
            <input 
              v-model="googleAnalytics.propertyId" 
              type="text" 
              placeholder="G-XXXXXXXXXX"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Measurement ID</label>
            <input 
              v-model="googleAnalytics.measurementId" 
              type="text" 
              placeholder="G-XXXXXXXXXX"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
          </div>
        </div>
        
        <div class="mt-4">
          <label class="flex items-center">
            <input 
              v-model="googleAnalytics.trackEvents" 
              type="checkbox" 
              class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
            />
            <span class="text-sm text-gray-700 mr-2">پیگیری رویدادهای تست</span>
          </label>
        </div>
        
        <button 
          @click="saveGoogleAnalytics" 
          class="mt-4 px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700"
        >
          ذخیره تنظیمات
        </button>
      </div>
    </div>

    <!-- Facebook Pixel -->
    <div class="mb-8">
      <div class="flex items-center justify-between mb-4">
        <div class="flex items-center">
          <div class="w-10 h-10 bg-blue-600 rounded-lg flex items-center justify-center mr-3">
            <svg class="w-6 h-6 text-white" fill="currentColor" viewBox="0 0 24 24">
              <path d="M24 12.073c0-6.627-5.373-12-12-12s-12 5.373-12 12c0 5.99 4.388 10.954 10.125 11.854v-8.385H7.078v-3.47h3.047V9.43c0-3.007 1.792-4.669 4.533-4.669 1.312 0 2.686.235 2.686.235v2.953H15.83c-1.491 0-1.956.925-1.956 1.874v2.25h3.328l-.532 3.47h-2.796v8.385C19.612 23.027 24 18.062 24 12.073z"/>
            </svg>
          </div>
          <div>
            <h4 class="text-md font-medium text-gray-900">Facebook Pixel</h4>
            <p class="text-sm text-gray-500">پیگیری تبدیل‌ها در Facebook Ads</p>
          </div>
        </div>
        <div class="flex items-center space-x-2 space-x-reverse">
          <span :class="[
            'px-2 py-1 text-xs rounded-full',
            facebookPixel.connected ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'
          ]">
            {{ facebookPixel.connected ? 'متصل' : 'قطع' }}
          </span>
          <button 
            @click="toggleFacebookPixel"
            :class="[
              'px-3 py-1 text-sm rounded-lg',
              facebookPixel.connected 
                ? 'bg-red-50 text-red-700 hover:bg-red-100' 
                : 'bg-green-50 text-green-700 hover:bg-green-100'
            ]"
          >
            {{ facebookPixel.connected ? 'قطع اتصال' : 'اتصال' }}
          </button>
        </div>
      </div>

      <div v-if="facebookPixel.connected" class="bg-gray-50 rounded-lg p-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Pixel ID</label>
          <input 
            v-model="facebookPixel.pixelId" 
            type="text" 
            placeholder="123456789012345"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>
        
        <div class="mt-4">
          <label class="flex items-center">
            <input 
              v-model="facebookPixel.trackConversions" 
              type="checkbox" 
              class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
            />
            <span class="text-sm text-gray-700 mr-2">پیگیری تبدیل‌ها</span>
          </label>
        </div>
        
        <div class="mt-2">
          <label class="flex items-center">
            <input 
              v-model="facebookPixel.trackCustomEvents" 
              type="checkbox" 
              class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
            />
            <span class="text-sm text-gray-700 mr-2">پیگیری رویدادهای سفارشی</span>
          </label>
        </div>
        
        <button 
          @click="saveFacebookPixel" 
          class="mt-4 px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700"
        >
          ذخیره تنظیمات
        </button>
      </div>
    </div>

    <!-- CRM Integration -->
    <div class="mb-8">
      <div class="flex items-center justify-between mb-4">
        <div class="flex items-center">
          <div class="w-10 h-10 bg-purple-500 rounded-lg flex items-center justify-center mr-3">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
            </svg>
          </div>
          <div>
            <h4 class="text-md font-medium text-gray-900">سیستم CRM</h4>
            <p class="text-sm text-gray-500">ارسال نتایج تست‌ها به سیستم CRM</p>
          </div>
        </div>
        <div class="flex items-center space-x-2 space-x-reverse">
          <span :class="[
            'px-2 py-1 text-xs rounded-full',
            crmIntegration.connected ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'
          ]">
            {{ crmIntegration.connected ? 'متصل' : 'قطع' }}
          </span>
          <button 
            @click="toggleCRMIntegration"
            :class="[
              'px-3 py-1 text-sm rounded-lg',
              crmIntegration.connected 
                ? 'bg-red-50 text-red-700 hover:bg-red-100' 
                : 'bg-green-50 text-green-700 hover:bg-green-100'
            ]"
          >
            {{ crmIntegration.connected ? 'قطع اتصال' : 'اتصال' }}
          </button>
        </div>
      </div>

      <div v-if="crmIntegration.connected" class="bg-gray-50 rounded-lg p-6">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">نوع CRM</label>
            <select 
              v-model="crmIntegration.type" 
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="salesforce">Salesforce</option>
              <option value="hubspot">HubSpot</option>
              <option value="pipedrive">Pipedrive</option>
              <option value="custom">سفارشی</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">API Key</label>
            <input 
              v-model="crmIntegration.apiKey" 
              type="password" 
              placeholder="API Key"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
          </div>
        </div>
        
        <div class="mt-4">
          <label class="flex items-center">
            <input 
              v-model="crmIntegration.syncResults" 
              type="checkbox" 
              class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
            />
            <span class="text-sm text-gray-700 mr-2">همگام‌سازی نتایج تست‌ها</span>
          </label>
        </div>
        
        <div class="mt-2">
          <label class="flex items-center">
            <input 
              v-model="crmIntegration.syncLeads" 
              type="checkbox" 
              class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
            />
            <span class="text-sm text-gray-700 mr-2">همگام‌سازی لیدها</span>
          </label>
        </div>
        
        <button 
          @click="saveCRMIntegration" 
          class="mt-4 px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700"
        >
          ذخیره تنظیمات
        </button>
      </div>
    </div>

    <!-- Email Marketing -->
    <div class="mb-8">
      <div class="flex items-center justify-between mb-4">
        <div class="flex items-center">
          <div class="w-10 h-10 bg-orange-500 rounded-lg flex items-center justify-center mr-3">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 4.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
            </svg>
          </div>
          <div>
            <h4 class="text-md font-medium text-gray-900">ایمیل مارکتینگ</h4>
            <p class="text-sm text-gray-500">ارسال نتایج تست‌ها به سیستم ایمیل مارکتینگ</p>
          </div>
        </div>
        <div class="flex items-center space-x-2 space-x-reverse">
          <span :class="[
            'px-2 py-1 text-xs rounded-full',
            emailMarketing.connected ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'
          ]">
            {{ emailMarketing.connected ? 'متصل' : 'قطع' }}
          </span>
          <button 
            @click="toggleEmailMarketing"
            :class="[
              'px-3 py-1 text-sm rounded-lg',
              emailMarketing.connected 
                ? 'bg-red-50 text-red-700 hover:bg-red-100' 
                : 'bg-green-50 text-green-700 hover:bg-green-100'
            ]"
          >
            {{ emailMarketing.connected ? 'قطع اتصال' : 'اتصال' }}
          </button>
        </div>
      </div>

      <div v-if="emailMarketing.connected" class="bg-gray-50 rounded-lg p-6">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">نوع سرویس</label>
            <select 
              v-model="emailMarketing.service" 
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="mailchimp">Mailchimp</option>
              <option value="sendgrid">SendGrid</option>
              <option value="convertkit">ConvertKit</option>
              <option value="custom">سفارشی</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">API Key</label>
            <input 
              v-model="emailMarketing.apiKey" 
              type="password" 
              placeholder="API Key"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
          </div>
        </div>
        
        <div class="mt-4">
          <label class="block text-sm font-medium text-gray-700 mb-1">لیست ایمیل</label>
          <input 
            v-model="emailMarketing.listId" 
            type="text" 
            placeholder="List ID"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>
        
        <div class="mt-4">
          <label class="flex items-center">
            <input 
              v-model="emailMarketing.sendReports" 
              type="checkbox" 
              class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
            />
            <span class="text-sm text-gray-700 mr-2">ارسال گزارش‌های هفتگی</span>
          </label>
        </div>
        
        <button 
          @click="saveEmailMarketing" 
          class="mt-4 px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700"
        >
          ذخیره تنظیمات
        </button>
      </div>
    </div>

    <!-- تست اتصال -->
    <div class="bg-blue-50 rounded-lg p-6">
      <h4 class="text-md font-medium text-gray-900 mb-3">تست اتصال</h4>
      <p class="text-sm text-gray-600 mb-4">برای اطمینان از صحت تنظیمات، اتصال‌ها را تست کنید:</p>
      
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-3">
        <button 
          @click="testConnection('google')"
          :disabled="!googleAnalytics.connected"
          :class="[
            'px-3 py-2 text-sm rounded-lg flex items-center justify-center',
            googleAnalytics.connected 
              ? 'bg-blue-100 text-blue-700 hover:bg-blue-200' 
              : 'bg-gray-100 text-gray-400 cursor-not-allowed'
          ]"
        >
          <svg class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          Google Analytics
        </button>
        
        <button 
          @click="testConnection('facebook')"
          :disabled="!facebookPixel.connected"
          :class="[
            'px-3 py-2 text-sm rounded-lg flex items-center justify-center',
            facebookPixel.connected 
              ? 'bg-blue-100 text-blue-700 hover:bg-blue-200' 
              : 'bg-gray-100 text-gray-400 cursor-not-allowed'
          ]"
        >
          <svg class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          Facebook Pixel
        </button>
        
        <button 
          @click="testConnection('crm')"
          :disabled="!crmIntegration.connected"
          :class="[
            'px-3 py-2 text-sm rounded-lg flex items-center justify-center',
            crmIntegration.connected 
              ? 'bg-blue-100 text-blue-700 hover:bg-blue-200' 
              : 'bg-gray-100 text-gray-400 cursor-not-allowed'
          ]"
        >
          <svg class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          CRM
        </button>
        
        <button 
          @click="testConnection('email')"
          :disabled="!emailMarketing.connected"
          :class="[
            'px-3 py-2 text-sm rounded-lg flex items-center justify-center',
            emailMarketing.connected 
              ? 'bg-blue-100 text-blue-700 hover:bg-blue-200' 
              : 'bg-gray-100 text-gray-400 cursor-not-allowed'
          ]"
        >
          <svg class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          Email
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
// Props
interface Props {
  modelValue?: any
}

const props = withDefaults(defineProps<Props>(), {
  modelValue: () => ({})
})

// Emits
const emit = defineEmits<{
  'update:modelValue': [value: any]
  'integration-changed': [integration: string, data: any]
}>()

// State
const googleAnalytics = ref({
  connected: false,
  propertyId: '',
  measurementId: '',
  trackEvents: true
})

const facebookPixel = ref({
  connected: false,
  pixelId: '',
  trackConversions: true,
  trackCustomEvents: false
})

const crmIntegration = ref({
  connected: false,
  type: 'salesforce',
  apiKey: '',
  syncResults: true,
  syncLeads: false
})

const emailMarketing = ref({
  connected: false,
  service: 'mailchimp',
  apiKey: '',
  listId: '',
  sendReports: true
})

// توابع
const refreshIntegrations = () => {
  // در حالت واقعی، وضعیت اتصال‌ها را از API دریافت می‌کنیم
  console.log('به‌روزرسانی اتصال‌ها')
}

const toggleGoogleAnalytics = () => {
  googleAnalytics.value.connected = !googleAnalytics.value.connected
  emit('integration-changed', 'google-analytics', googleAnalytics.value)
}

const toggleFacebookPixel = () => {
  facebookPixel.value.connected = !facebookPixel.value.connected
  emit('integration-changed', 'facebook-pixel', facebookPixel.value)
}

const toggleCRMIntegration = () => {
  crmIntegration.value.connected = !crmIntegration.value.connected
  emit('integration-changed', 'crm', crmIntegration.value)
}

const toggleEmailMarketing = () => {
  emailMarketing.value.connected = !emailMarketing.value.connected
  emit('integration-changed', 'email-marketing', emailMarketing.value)
}

const saveGoogleAnalytics = () => {
  console.log('ذخیره تنظیمات Google Analytics:', googleAnalytics.value)
  emit('integration-changed', 'google-analytics', googleAnalytics.value)
}

const saveFacebookPixel = () => {
  console.log('ذخیره تنظیمات Facebook Pixel:', facebookPixel.value)
  emit('integration-changed', 'facebook-pixel', facebookPixel.value)
}

const saveCRMIntegration = () => {
  console.log('ذخیره تنظیمات CRM:', crmIntegration.value)
  emit('integration-changed', 'crm', crmIntegration.value)
}

const saveEmailMarketing = () => {
  console.log('ذخیره تنظیمات Email Marketing:', emailMarketing.value)
  emit('integration-changed', 'email-marketing', emailMarketing.value)
}

const testConnection = (type: string) => {
  console.log(`تست اتصال ${type}`)
  // در حالت واقعی، اتصال را تست می‌کنیم
  setTimeout(() => {
    alert(`اتصال ${type} با موفقیت تست شد!`)
  }, 1000)
}

// Watch برای تغییرات props
watch(() => props.modelValue, (newValue) => {
  if (newValue) {
    if (newValue.googleAnalytics) {
      googleAnalytics.value = { ...googleAnalytics.value, ...newValue.googleAnalytics }
    }
    if (newValue.facebookPixel) {
      facebookPixel.value = { ...facebookPixel.value, ...newValue.facebookPixel }
    }
    if (newValue.crmIntegration) {
      crmIntegration.value = { ...crmIntegration.value, ...newValue.crmIntegration }
    }
    if (newValue.emailMarketing) {
      emailMarketing.value = { ...emailMarketing.value, ...newValue.emailMarketing }
    }
  }
}, { immediate: true })
</script> 
