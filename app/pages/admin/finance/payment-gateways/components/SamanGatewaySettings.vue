<template>
  <div class="bg-white rounded-lg shadow-lg p-6">
    <div class="flex items-center justify-between mb-6">
      <div class="flex items-center space-x-3 space-x-reverse">
        <div class="w-12 h-12 bg-blue-600 rounded-lg flex items-center justify-center">
          <span class="text-white font-bold text-lg">سامان</span>
        </div>
        <div>
          <h3 class="text-lg font-semibold text-gray-900">تنظیمات درگاه سامان</h3>
          <p class="text-sm text-gray-500">مدیریت تنظیمات و پارامترهای درگاه پرداخت بانک سامان</p>
        </div>
      </div>
      <div class="flex items-center space-x-2 space-x-reverse">
        <button
          @click="testConnection"
          :disabled="testing"
          class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:opacity-50"
        >
          <span v-if="testing" class="i-heroicons-arrow-path animate-spin mr-2"></span>
          {{ testing ? 'در حال تست...' : 'تست اتصال' }}
        </button>
        <button
          @click="saveSettings"
          :disabled="saving"
          class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 disabled:opacity-50"
        >
          <span v-if="saving" class="i-heroicons-arrow-path animate-spin mr-2"></span>
          {{ saving ? 'در حال ذخیره...' : 'ذخیره تنظیمات' }}
        </button>
      </div>
    </div>

    <div v-if="loading" class="flex justify-center py-8">
      <div class="i-heroicons-arrow-path animate-spin text-2xl text-blue-600"></div>
    </div>

    <form v-else @submit.prevent="saveSettings" class="space-y-6">
      <!-- تنظیمات اصلی -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">شماره ترمینال</label>
          <input
            v-model="settings.merchant_id"
            type="text"
            required
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            placeholder="شماره ترمینال سامان"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">کلید خصوصی</label>
          <input
            v-model="settings.private_key"
            type="password"
            required
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            placeholder="کلید خصوصی سامان"
          />
        </div>
      </div>

      <!-- تنظیمات اختیاری -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">کلید عمومی</label>
          <input
            v-model="settings.public_key"
            type="text"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            placeholder="کلید عمومی (اختیاری)"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">کلید تست</label>
          <input
            v-model="settings.test_key"
            type="text"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            placeholder="کلید تست (اختیاری)"
          />
        </div>
      </div>

      <!-- محدودیت‌های مبلغ -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداقل مبلغ (تومان)</label>
          <input
            v-model.number="settings.min_amount"
            type="number"
            min="0"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            placeholder="1000"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر مبلغ (تومان)</label>
          <input
            v-model.number="settings.max_amount"
            type="number"
            min="0"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            placeholder="100000000"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">کارمزد (درصد)</label>
          <input
            v-model.number="settings.fee"
            type="number"
            min="0"
            max="100"
            step="0.1"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            placeholder="0"
          />
        </div>
      </div>

      <!-- تنظیمات وضعیت -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">اولویت</label>
          <input
            v-model.number="settings.priority"
            type="number"
            min="1"
            max="100"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            placeholder="1"
          />
        </div>

        <div class="flex items-center space-x-4 space-x-reverse">
          <label class="flex items-center">
            <input
              v-model="settings.is_active"
              type="checkbox"
              class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
            />
            <span class="mr-2 text-sm font-medium text-gray-700">فعال</span>
          </label>

          <label class="flex items-center">
            <input
              v-model="settings.is_test_mode"
              type="checkbox"
              class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
            />
            <span class="mr-2 text-sm font-medium text-gray-700">حالت تست</span>
          </label>
        </div>
      </div>

      <!-- اطلاعات درگاه -->
      <div class="bg-gray-50 rounded-lg p-6">
        <h4 class="text-sm font-medium text-gray-900 mb-3">اطلاعات درگاه سامان</h4>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6 text-sm">
          <div>
            <span class="text-gray-600">نوع درگاه:</span>
            <span class="mr-2 font-medium">سامان (SEP)</span>
          </div>
          <div>
            <span class="text-gray-600">پشتیبانی از بازگشت وجه:</span>
            <span class="mr-2 font-medium text-green-600">بله</span>
          </div>
          <div>
            <span class="text-gray-600">پشتیبانی از موجودی:</span>
            <span class="mr-2 font-medium text-red-600">خیر</span>
          </div>
          <div>
            <span class="text-gray-600">حالت تست:</span>
            <span class="mr-2 font-medium text-green-600">پشتیبانی می‌شود</span>
          </div>
        </div>
      </div>
    </form>

    <!-- پیام‌های نتیجه -->
    <div v-if="message" class="mt-4 p-6 rounded-lg" :class="messageType === 'success' ? 'bg-green-50 text-green-800' : 'bg-red-50 text-red-800'">
      {{ message }}
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'

// Type declaration for Nuxt 4 auto-imported functions
declare const definePageMeta: (meta: { layout?: string }) => void

definePageMeta({
  layout: 'admin-main'
})

// متغیرهای reactive
const loading = ref(false)
const saving = ref(false)
const testing = ref(false)
const message = ref('')
const messageType = ref<'success' | 'error'>('success')

// تنظیمات درگاه
const settings = ref({
  merchant_id: '',
  private_key: '',
  public_key: '',
  test_key: '',
  min_amount: 1000,
  max_amount: 100000000,
  fee: 0,
  priority: 1,
  is_active: false,
  is_test_mode: false
})

// دریافت تنظیمات
const fetchSettings = async () => {
  loading.value = true
  try {
    interface SettingsData {
      merchant_id?: string
      api_keys?: {
        private_key?: string
        public_key?: string
        test_key?: string
      }
      min_amount?: number
      max_amount?: number
      fee?: number
      priority?: number
      status?: string
      is_test_mode?: boolean
    }
    interface ApiResponse {
      success?: boolean
      data?: SettingsData
    }
    const response = await $fetch<ApiResponse>('/api/admin/payment-gateways/saman.settings')
    if (response.success && response.data) {
      const data = response.data
      settings.value = {
        merchant_id: data.merchant_id || '',
        private_key: data.api_keys?.private_key || '',
        public_key: data.api_keys?.public_key || '',
        test_key: data.api_keys?.test_key || '',
        min_amount: data.min_amount || 1000,
        max_amount: data.max_amount || 100000000,
        fee: data.fee || 0,
        priority: data.priority || 1,
        is_active: data.status === 'active',
        is_test_mode: data.is_test_mode || false
      }
    }
  } catch (error) {
    const errorMessage = error instanceof Error ? error.message : 'خطای نامشخص'
    console.error('خطا در دریافت تنظیمات:', errorMessage)
    message.value = 'خطا در دریافت تنظیمات درگاه سامان'
    messageType.value = 'error'
  } finally {
    loading.value = false
  }
}

// ذخیره تنظیمات
const saveSettings = async () => {
  saving.value = true
  message.value = ''
  
  try {
    interface ApiResponse {
      success?: boolean
      message?: string
    }
    const response = await $fetch<ApiResponse>('/api/admin/payment-gateways/saman.settings', {
      method: 'PUT',
      body: settings.value
    })
    
    if (response.success) {
      message.value = 'تنظیمات درگاه سامان با موفقیت ذخیره شد'
      messageType.value = 'success'
    } else {
      message.value = response.message || 'خطا در ذخیره تنظیمات'
      messageType.value = 'error'
    }
  } catch (error) {
    const errorMessage = error instanceof Error ? error.message : 'خطای نامشخص'
    console.error('خطا در ذخیره تنظیمات:', errorMessage)
    message.value = 'خطا در ذخیره تنظیمات درگاه سامان'
    messageType.value = 'error'
  } finally {
    saving.value = false
  }
}

// تست اتصال
const testConnection = async () => {
  testing.value = true
  message.value = ''
  
  try {
    interface ApiResponse {
      success?: boolean
      message?: string
    }
    const response = await $fetch<ApiResponse>('/api/admin/payment-gateways/saman.test-connection', {
      method: 'POST'
    })
    
    if (response.success) {
      message.value = 'اتصال به درگاه سامان موفقیت‌آمیز بود'
      messageType.value = 'success'
    } else {
      message.value = response.message || 'خطا در تست اتصال'
      messageType.value = 'error'
    }
  } catch (error) {
    const errorMessage = error instanceof Error ? error.message : 'خطای نامشخص'
    console.error('خطا در تست اتصال:', errorMessage)
    message.value = 'خطا در تست اتصال به درگاه سامان'
    messageType.value = 'error'
  } finally {
    testing.value = false
  }
}

// دریافت تنظیمات در لود صفحه
onMounted(() => {
  fetchSettings()
})
</script> 
 
 
