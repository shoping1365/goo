<template>
  <div class="bg-white rounded-lg shadow-md p-6">
    <div class="flex items-center justify-between mb-6">
      <div class="flex items-center space-x-3">
        <div class="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center">
          <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z" />
          </svg>
        </div>
        <div>
          <h3 class="text-lg font-semibold text-gray-900">درگاه پرداخت ملت</h3>
          <p class="text-sm text-gray-500">مدیریت تنظیمات و تست اتصال</p>
        </div>
      </div>
      <div class="flex items-center space-x-2">
        <span
class="px-3 py-1 text-xs font-medium rounded-full" 
              :class="gateway?.is_active ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'">
          {{ gateway?.is_active ? 'فعال' : 'غیرفعال' }}
        </span>
      </div>
    </div>

    <!-- فرم تنظیمات -->
    <form class="space-y-6" @submit.prevent="saveSettings">
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <!-- Terminal ID -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            شماره ترمینال
          </label>
          <input
            v-model="form.merchantId"
            type="text"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            placeholder="مثال: 2768649"
            required
          />
          <p class="mt-1 text-xs text-gray-500">
            شماره ترمینال دریافتی از بانک ملت
          </p>
        </div>

        <!-- Username -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            نام کاربری
          </label>
          <input
            v-model="form.publicKey"
            type="text"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            placeholder="مثال: be27"
            required
          />
          <p class="mt-1 text-xs text-gray-500">
            نام کاربری دریافتی از بانک ملت
          </p>
        </div>

        <!-- Password -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            رمز عبور
          </label>
          <input
            v-model="form.privateKey"
            type="password"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            placeholder="رمز عبور دریافتی از بانک ملت"
            required
          />
          <p class="mt-1 text-xs text-gray-500">
            رمز عبور دریافتی از بانک ملت
          </p>
        </div>

        <!-- Callback URL -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            آدرس Callback
          </label>
          <input
            v-model="form.callbackUrl"
            type="url"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            placeholder="https://example.com/payment/callback"
            required
          />
          <p class="mt-1 text-xs text-gray-500">
            آدرس بازگشت پس از پرداخت
          </p>
        </div>
      </div>

      <!-- تنظیمات پیشرفته -->
      <div class="border-t pt-6">
        <h4 class="text-md font-medium text-gray-900 mb-4">تنظیمات پیشرفته</h4>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <!-- حالت تست -->
          <div class="flex items-center">
            <input
              id="testMode"
              v-model="form.isTestMode"
              type="checkbox"
              class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
            />
            <label for="testMode" class="mr-2 block text-sm text-gray-900">
              حالت تست
            </label>
          </div>

          <!-- وضعیت فعال -->
          <div class="flex items-center">
            <input
              id="isActive"
              v-model="form.isActive"
              type="checkbox"
              class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
            />
            <label for="isActive" class="mr-2 block text-sm text-gray-900">
              فعال
            </label>
          </div>
        </div>
      </div>

      <!-- دکمه‌های عملیات -->
      <div class="flex items-center justify-between pt-6 border-t">
        <div class="flex items-center space-x-3">
          <button
            type="button"
            :disabled="testing"
            class="px-4 py-2 bg-yellow-500 text-white rounded-md hover:bg-yellow-600 focus:outline-none focus:ring-2 focus:ring-yellow-500 disabled:opacity-50"
            @click="testConnection"
          >
            <span v-if="testing" class="flex items-center">
              <svg class="animate-spin -mr-1 ml-2 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              در حال تست...
            </span>
            <span v-else>تست اتصال</span>
          </button>
        </div>

        <div class="flex items-center space-x-3">
          <button
            type="button"
            class="px-4 py-2 bg-gray-500 text-white rounded-md hover:bg-gray-600 focus:outline-none focus:ring-2 focus:ring-gray-500"
            @click="resetForm"
          >
            بازنشانی
          </button>
          <button
            type="submit"
            :disabled="saving"
            class="px-6 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 disabled:opacity-50"
          >
            <span v-if="saving" class="flex items-center">
              <svg class="animate-spin -mr-1 ml-2 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              در حال ذخیره...
            </span>
            <span v-else>ذخیره تنظیمات</span>
          </button>
        </div>
      </div>
    </form>

    <!-- نتایج تست -->
    <div v-if="testResult" class="mt-6 p-6 rounded-md" :class="testResult.success ? 'bg-green-50 border border-green-200' : 'bg-red-50 border border-red-200'">
      <div class="flex">
        <div class="flex-shrink-0">
          <svg v-if="testResult.success" class="h-5 w-5 text-green-400" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
          </svg>
          <svg v-else class="h-5 w-5 text-red-400" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd" />
          </svg>
        </div>
        <div class="mr-3">
          <h3 class="text-sm font-medium" :class="testResult.success ? 'text-green-800' : 'text-red-800'">
            {{ testResult.success ? 'تست موفقیت‌آمیز' : 'تست ناموفق' }}
          </h3>
          <div class="mt-2 text-sm" :class="testResult.success ? 'text-green-700' : 'text-red-700'">
            <p>{{ testResult.message }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- اطلاعات اضافی -->
    <div class="mt-6 bg-gray-50 rounded-lg p-6">
      <h4 class="text-sm font-medium text-gray-900 mb-3">اطلاعات درگاه ملت</h4>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6 text-sm text-gray-600">
        <div>
          <span class="font-medium">آدرس Web Service:</span>
          <span class="block text-xs font-mono bg-gray-100 p-1 rounded mt-1">
            {{ form.isTestMode ? 'https://pgwstest.bpm.bankmellat.ir/pgwchannel/services/pgw' : 'https://bpm.shaparak.ir/pgwchannel/services/pgw' }}
          </span>
        </div>
        <div>
          <span class="font-medium">آدرس پرداخت:</span>
          <span class="block text-xs font-mono bg-gray-100 p-1 rounded mt-1">
            {{ form.isTestMode ? 'https://pgwstest.bpm.bankmellat.ir/pgwchannel/startpay.mellat' : 'https://bpm.shaparak.ir/pgwchannel/startpay.mellat' }}
          </span>
        </div>
        <div>
          <span class="font-medium">پروتکل:</span>
          <span>SOAP Web Service</span>
        </div>
        <div>
          <span class="font-medium">واحد پول:</span>
          <span>تومان</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'

// تعریف interface برای gateway
interface PaymentGateway {
  id: number
  name: string
  type: string
  merchant_id: string
  public_key: string
  private_key: string
  callback_url: string
  is_test_mode: boolean
  is_active: boolean
  created_at: string
  updated_at: string
}

// تعریف interface برای فرم
interface FormData {
  merchantId: string
  publicKey: string
  privateKey: string
  callbackUrl: string
  isTestMode: boolean
  isActive: boolean
}

// تعریف interface برای نتیجه تست
interface TestResult {
  success: boolean
  message: string
}

// متغیرهای reactive
const gateway = ref<PaymentGateway | null>(null)
const saving = ref(false)
const testing = ref(false)
const testResult = ref<TestResult | null>(null)

// فرم تنظیمات
const form = reactive<FormData>({
  merchantId: '',
  publicKey: '',
  privateKey: '',
  callbackUrl: '',
  isTestMode: false,
  isActive: false
})

// بارگذاری تنظیمات درگاه
const loadGatewaySettings = async () => {
  try {
    interface MellatGatewayResponse {
      success?: boolean
      data?: {
        id?: number | string
        [key: string]: unknown
      }
      [key: string]: unknown
    }
    const response = await $fetch<MellatGatewayResponse>('/api/admin/payment-gateways/mellat')
    if (response.success && response.data) {
      gateway.value = response.data as unknown as PaymentGateway
      form.merchantId = gateway.value.merchant_id || ''
      form.publicKey = gateway.value.public_key || ''
      form.privateKey = gateway.value.private_key || ''
      form.callbackUrl = gateway.value.callback_url || ''
      form.isTestMode = gateway.value.is_test_mode || false
      form.isActive = gateway.value.is_active || false
    }
  } catch (error) {
    console.error('خطا در بارگذاری تنظیمات ملت:', error)
  }
}

// ذخیره تنظیمات
const saveSettings = async () => {
  saving.value = true
  try {
    const response = await $fetch('/api/admin/payment-gateways/mellat', {
      method: 'PUT',
      body: {
        merchant_id: form.merchantId,
        public_key: form.publicKey,
        private_key: form.privateKey,
        callback_url: form.callbackUrl,
        is_test_mode: form.isTestMode,
        is_active: form.isActive
      }
    }) as { success?: boolean; [key: string]: unknown }

    if (response.success) {
      // نمایش پیام موفقیت
      alert('تنظیمات ملت با موفقیت ذخیره شد')
      await loadGatewaySettings() // بارگذاری مجدد
    }
  } catch (error: unknown) {
    console.error('خطا در ذخیره تنظیمات ملت:', error)
    const errorData = error as { data?: { message?: string }; message?: string }
    alert(errorData.data?.message || errorData.message || 'خطا در ذخیره تنظیمات')
  } finally {
    saving.value = false
  }
}

// تست اتصال
const testConnection = async () => {
  testing.value = true
  testResult.value = null
  
  try {
    const response = await $fetch('/api/admin/payment-gateways/mellat/test', {
      method: 'POST',
      body: {
        merchant_id: form.merchantId,
        public_key: form.publicKey,
        private_key: form.privateKey,
        is_test_mode: form.isTestMode
      }
    }) as { success?: boolean; [key: string]: unknown }

    testResult.value = {
      success: Boolean(response.success),
      message: String(response.message || (response.success ? 'اتصال موفقیت‌آمیز' : 'اتصال ناموفق'))
    }
  } catch (error: unknown) {
    testResult.value = {
      success: false,
      message: (error as { data?: { message?: string } }).data?.message || 'خطا در تست اتصال'
    }
  } finally {
    testing.value = false
  }
}

// بازنشانی فرم
const resetForm = () => {
  if (gateway.value) {
    form.merchantId = gateway.value.merchant_id || ''
    form.publicKey = gateway.value.public_key || ''
    form.privateKey = gateway.value.private_key || ''
    form.callbackUrl = gateway.value.callback_url || ''
    form.isTestMode = gateway.value.is_test_mode || false
    form.isActive = gateway.value.is_active || false
  }
  testResult.value = null
}

// بارگذاری اولیه
onMounted(() => {
  loadGatewaySettings()
})
</script> 
 
 
