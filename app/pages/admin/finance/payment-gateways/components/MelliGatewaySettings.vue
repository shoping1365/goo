<template>
  <div class="bg-white rounded-lg shadow-md p-6">
    <!-- Header -->
    <div class="flex items-center justify-between mb-6">
      <div class="flex items-center space-x-3">
        <div class="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center">
          <Icon name="ph:bank" class="w-6 h-6 text-blue-600" />
        </div>
        <div>
          <h2 class="text-xl font-bold text-gray-900">درگاه پرداخت ملی</h2>
          <p class="text-sm text-gray-500">تنظیمات درگاه پرداخت بانک ملی ایران</p>
        </div>
      </div>
      <div class="flex items-center space-x-2">
        <span 
          :class="[
            'px-3 py-1 rounded-full text-xs font-medium',
            gatewayStatus === 'active' 
              ? 'bg-green-100 text-green-800' 
              : 'bg-red-100 text-red-800'
          ]"
        >
          {{ gatewayStatus === 'active' ? 'فعال' : 'غیرفعال' }}
        </span>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="flex items-center justify-center py-8">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
      <span class="mr-3 text-gray-600">در حال بارگذاری...</span>
    </div>

    <!-- Settings Form -->
    <form v-else @submit.prevent="saveSettings" class="space-y-6">
      <!-- Terminal ID -->
      <div>
        <label for="terminal_id" class="block text-sm font-medium text-gray-700 mb-2">
          شماره ترمینال
        </label>
        <input
          id="terminal_id"
          v-model="form.merchant_id"
          type="text"
          class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          placeholder="شماره ترمینال خود را وارد کنید"
          required
        />
      </div>

      <!-- Merchant ID -->
      <div>
        <label for="merchant_id" class="block text-sm font-medium text-gray-700 mb-2">
          مرچنت ID
        </label>
        <input
          id="merchant_id"
          v-model="form.public_key"
          type="text"
          class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          placeholder="مرچنت ID خود را وارد کنید"
          required
        />
      </div>

      <!-- Private Key -->
      <div>
        <label for="private_key" class="block text-sm font-medium text-gray-700 mb-2">
          کلید رمزنگاری
        </label>
        <input
          id="private_key"
          v-model="form.private_key"
          type="password"
          class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          placeholder="کلید رمزنگاری خود را وارد کنید"
          required
        />
      </div>

      <!-- Callback URL -->
      <div>
        <label for="callback_url" class="block text-sm font-medium text-gray-700 mb-2">
          آدرس Callback
        </label>
        <input
          id="callback_url"
          v-model="form.callback_url"
          type="url"
          class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          placeholder="https://yourdomain.com/api/payments/melli/callback"
          required
        />
      </div>

      <!-- Test Mode Toggle -->
      <div class="flex items-center justify-between">
        <div>
          <label class="text-sm font-medium text-gray-700">حالت تست</label>
          <p class="text-xs text-gray-500">برای تست درگاه از محیط تستی استفاده کنید</p>
        </div>
        <button
          type="button"
          @click="form.is_test_mode = !form.is_test_mode"
          :class="[
            'relative inline-flex h-6 w-11 items-center rounded-full transition-colors focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2',
            form.is_test_mode ? 'bg-blue-600' : 'bg-gray-200'
          ]"
        >
          <span
            :class="[
              'inline-block h-4 w-4 transform rounded-full bg-white transition-transform',
              form.is_test_mode ? 'translate-x-6' : 'translate-x-1'
            ]"
          />
        </button>
      </div>

      <!-- Active Status Toggle -->
      <div class="flex items-center justify-between">
        <div>
          <label class="text-sm font-medium text-gray-700">وضعیت فعال</label>
          <p class="text-xs text-gray-500">فعال یا غیرفعال کردن درگاه</p>
        </div>
        <button
          type="button"
          @click="form.is_active = !form.is_active"
          :class="[
            'relative inline-flex h-6 w-11 items-center rounded-full transition-colors focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2',
            form.is_active ? 'bg-green-600' : 'bg-gray-200'
          ]"
        >
          <span
            :class="[
              'inline-block h-4 w-4 transform rounded-full bg-white transition-transform',
              form.is_active ? 'translate-x-6' : 'translate-x-1'
            ]"
          />
        </button>
      </div>

      <!-- Action Buttons -->
      <div class="flex items-center justify-between pt-6 border-t border-gray-200">
        <div class="flex items-center space-x-3">
          <button
            type="button"
            @click="testConnection"
            :disabled="testing"
            class="px-4 py-2 text-sm font-medium text-blue-600 bg-blue-50 border border-blue-200 rounded-md hover:bg-blue-100 focus:outline-none focus:ring-2 focus:ring-blue-500 disabled:opacity-50 disabled:cursor-not-allowed"
          >
            <Icon v-if="testing" name="ph:spinner" class="w-4 h-4 animate-spin" />
            <Icon v-else name="ph:plug" class="w-4 h-4" />
            <span class="mr-2">{{ testing ? 'در حال تست...' : 'تست اتصال' }}</span>
          </button>
        </div>

        <div class="flex items-center space-x-3">
          <button
            type="button"
            @click="resetForm"
            class="px-4 py-2 text-sm font-medium text-gray-600 bg-gray-50 border border-gray-200 rounded-md hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-500"
          >
            بازنشانی
          </button>
          <button
            type="submit"
            :disabled="saving"
            class="px-6 py-2 text-sm font-medium text-white bg-blue-600 border border-transparent rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 disabled:opacity-50 disabled:cursor-not-allowed"
          >
            <Icon v-if="saving" name="ph:spinner" class="w-4 h-4 animate-spin" />
            <Icon v-else name="ph:floppy-disk" class="w-4 h-4" />
            <span class="mr-2">{{ saving ? 'در حال ذخیره...' : 'ذخیره تنظیمات' }}</span>
          </button>
        </div>
      </div>
    </form>

    <!-- Test Result Modal -->
    <div v-if="showTestResult" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg p-6 max-w-md w-full mx-4">
        <div class="flex items-center space-x-3 mb-4">
          <Icon 
            :name="testResult.success ? 'ph:check-circle' : 'ph:x-circle'" 
            :class="[
              'w-6 h-6',
              testResult.success ? 'text-green-600' : 'text-red-600'
            ]" 
          />
          <h3 class="text-lg font-medium text-gray-900">
            {{ testResult.success ? 'تست موفق' : 'تست ناموفق' }}
          </h3>
        </div>
        <p class="text-gray-600 mb-6">{{ testResult.message }}</p>
        <div class="flex justify-end">
          <button
            @click="showTestResult = false"
            class="px-4 py-2 text-sm font-medium text-gray-600 bg-gray-50 border border-gray-200 rounded-md hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-500"
          >
            بستن
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue'
import { useToast } from '~/composables/useToast'

// Types
interface GatewaySettings {
  merchant_id: string
  public_key: string
  private_key: string
  callback_url: string
  is_test_mode: boolean
  is_active: boolean
}

interface TestResult {
  success: boolean
  message: string
}

// Reactive data
const loading = ref(false)
const saving = ref(false)
const testing = ref(false)
const showTestResult = ref(false)
const gatewayStatus = ref('inactive')

const form = reactive<GatewaySettings>({
  merchant_id: '',
  public_key: '',
  private_key: '',
  callback_url: '',
  is_test_mode: false,
  is_active: false
})

const testResult = reactive<TestResult>({
  success: false,
  message: ''
})

// Methods
const toast = useToast()

const loadSettings = async () => {
  try {
    loading.value = true
    interface SettingsData {
      merchant_id?: string
      api_keys?: {
        public_key?: string
        private_key?: string
      }
      api_endpoints?: {
        callback?: string
      }
      is_test_mode?: boolean
      status?: string
    }
    interface ApiResponse {
      success?: boolean
      data?: SettingsData
    }
    const response = await $fetch<ApiResponse>('/api/admin/payment-gateways/melli')
    
    if (response.success && response.data) {
      const data = response.data
      form.merchant_id = data.merchant_id || ''
      form.public_key = data.api_keys?.public_key || ''
      form.private_key = data.api_keys?.private_key || ''
      form.callback_url = data.api_endpoints?.callback || ''
      form.is_test_mode = data.is_test_mode || false
      form.is_active = data.status === 'active'
      gatewayStatus.value = data.status || 'inactive'
    }
  } catch (error) {
    const errorMessage = error instanceof Error ? error.message : 'خطای نامشخص'
    console.error('خطا در بارگذاری تنظیمات:', errorMessage)
    toast.showError('خطا در بارگذاری تنظیمات درگاه ملی')
  } finally {
    loading.value = false
  }
}

const saveSettings = async () => {
  try {
    saving.value = true
    interface ApiResponse {
      success?: boolean
      message?: string
    }
    const response = await $fetch<ApiResponse>('/api/admin/payment-gateways/melli', {
      method: 'PUT',
      body: form
    })
    
    if (response.success) {
      toast.showSuccess('تنظیمات درگاه ملی با موفقیت ذخیره شد')
      await loadSettings() // بارگذاری مجدد تنظیمات
    }
  } catch (error) {
    const errorMessage = error instanceof Error ? error.message : 'خطای نامشخص'
    console.error('خطا در ذخیره تنظیمات:', errorMessage)
    toast.showError('خطا در ذخیره تنظیمات درگاه ملی')
  } finally {
    saving.value = false
  }
}

const testConnection = async () => {
  try {
    testing.value = true
    interface ApiResponse {
      success?: boolean
      message?: string
    }
    const response = await $fetch<ApiResponse>('/api/admin/payment-gateways/melli/test', {
      method: 'POST',
      body: {
        merchant_id: form.merchant_id,
        public_key: form.public_key,
        private_key: form.private_key,
        is_test_mode: form.is_test_mode
      }
    })
    
    testResult.success = response.success || false
    testResult.message = response.message || 'تست اتصال انجام شد'
    showTestResult.value = true
  } catch (error) {
    const errorMessage = error instanceof Error ? error.message : 'خطای نامشخص'
    console.error('خطا در تست اتصال:', errorMessage)
    testResult.success = false
    testResult.message = errorMessage || 'خطا در تست اتصال'
    showTestResult.value = true
  } finally {
    testing.value = false
  }
}

const resetForm = () => {
  loadSettings()
}

// Lifecycle
onMounted(() => {
  loadSettings()
})
</script> 
