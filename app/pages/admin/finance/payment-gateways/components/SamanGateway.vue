<template>
  <div class="bg-white rounded-lg shadow-md p-6">
    <!-- هدر کامپوننت -->
    <div class="flex items-center justify-between mb-6">
      <div class="flex items-center space-x-3">
        <div class="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center">
          <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"></path>
          </svg>
        </div>
        <div>
          <h3 class="text-lg font-semibold text-gray-900">درگاه پرداخت سامان</h3>
          <p class="text-sm text-gray-500">مدیریت تنظیمات و عملیات درگاه سامان</p>
        </div>
      </div>
      <div class="flex items-center space-x-2">
        <span
class="px-3 py-1 text-xs font-medium rounded-full" 
              :class="gateway?.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'">
          {{ gateway?.status === 'active' ? 'فعال' : 'غیرفعال' }}
        </span>
      </div>
    </div>

    <!-- فرم تنظیمات -->
    <form class="space-y-6" @submit.prevent="saveSettings">
      <!-- اطلاعات پایه -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نام درگاه</label>
          <input 
            v-model="form.name" 
            type="text" 
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            placeholder="سامان"
            required
          />
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">مرچنت ID</label>
          <input 
            v-model="form.merchantId" 
            type="text" 
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            placeholder="00000000-0000-0000-0000-000000000000"
            required
          />
        </div>
      </div>

      <!-- رمزهای API -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">رمز خصوصی</label>
          <input 
            v-model="form.privateKey" 
            type="password" 
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            placeholder="رمز خصوصی سامان"
            required
          />
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">رمز عمومی</label>
          <input 
            v-model="form.publicKey" 
            type="text" 
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            placeholder="رمز عمومی سامان"
          />
        </div>
      </div>

      <!-- آدرس‌های API -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">آدرس پرداخت</label>
          <input 
            v-model="form.paymentUrl" 
            type="url" 
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            placeholder="https://sep.shaparak.ir/api/v1/Payment/GetToken"
            required
          />
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">آدرس تایید</label>
          <input 
            v-model="form.verifyUrl" 
            type="url" 
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            placeholder="https://acquirer.samanepay.com/payments/referencepayment.asmx"
            required
          />
        </div>
      </div>

      <!-- محدودیت‌های مبلغ -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداقل مبلغ (تومان)</label>
          <input 
            v-model.number="form.minAmount" 
            type="number" 
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            placeholder="1000"
            min="0"
          />
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر مبلغ (تومان)</label>
          <input 
            v-model.number="form.maxAmount" 
            type="number" 
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            placeholder="50000000"
            min="0"
          />
        </div>
      </div>

      <!-- تنظیمات اضافی -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div class="flex items-center">
          <input 
            v-model="form.isTestMode" 
            type="checkbox" 
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <label class="mr-2 block text-sm text-gray-900">حالت تست</label>
        </div>
        
        <div class="flex items-center">
          <input 
            v-model="form.isActive" 
            type="checkbox" 
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <label class="mr-2 block text-sm text-gray-900">فعال</label>
        </div>
        
        <div class="flex items-center">
          <input 
            v-model="form.autoVerify" 
            type="checkbox" 
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <label class="mr-2 block text-sm text-gray-900">تایید خودکار</label>
        </div>
      </div>

      <!-- دکمه‌های عملیات -->
      <div class="flex items-center justify-between pt-6 border-t border-gray-200">
        <div class="flex items-center space-x-3">
          <button 
            type="button"
            :disabled="testing"
            class="px-4 py-2 text-sm font-medium text-blue-700 bg-blue-100 border border-blue-300 rounded-md hover:bg-blue-200 focus:outline-none focus:ring-2 focus:ring-blue-500 disabled:opacity-50"
            @click="testConnection"
          >
            <span v-if="testing">در حال تست...</span>
            <span v-else>تست اتصال</span>
          </button>
          
          <button 
            type="button"
            class="px-4 py-2 text-sm font-medium text-gray-700 bg-gray-100 border border-gray-300 rounded-md hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500"
            @click="resetForm"
          >
            بازنشانی
          </button>
        </div>
        
        <button 
          type="submit"
          :disabled="saving"
          class="px-6 py-2 text-sm font-medium text-white bg-blue-600 border border-transparent rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 disabled:opacity-50"
        >
          <span v-if="saving">در حال ذخیره...</span>
          <span v-else>ذخیره تنظیمات</span>
        </button>
      </div>
    </form>

    <!-- نتایج تست -->
    <div v-if="testResult" class="mt-6 p-6 rounded-md" :class="testResult.success ? 'bg-green-50 border border-green-200' : 'bg-red-50 border border-red-200'">
      <div class="flex">
        <div class="flex-shrink-0">
          <svg v-if="testResult.success" class="h-5 w-5 text-green-400" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"></path>
          </svg>
          <svg v-else class="h-5 w-5 text-red-400" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd"></path>
          </svg>
        </div>
        <div class="mr-3">
          <h3 class="text-sm font-medium" :class="testResult.success ? 'text-green-800' : 'text-red-800'">
            {{ testResult.success ? 'تست موفقیت‌آمیز' : 'تست ناموفق' }}
          </h3>
          <div class="mt-2 text-sm" :class="testResult.success ? 'text-green-700' : 'text-red-700'">
            {{ testResult.message }}
          </div>
        </div>
      </div>
    </div>

    <!-- آمار و اطلاعات -->
    <div v-if="gateway" class="mt-8 grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="bg-gray-50 rounded-lg p-6">
        <div class="text-2xl font-bold text-gray-900">{{ gateway.totalTransactions || 0 }}</div>
        <div class="text-sm text-gray-500">کل تراکنش‌ها</div>
      </div>
      
      <div class="bg-green-50 rounded-lg p-6">
        <div class="text-2xl font-bold text-green-600">{{ gateway.successfulTransactions || 0 }}</div>
        <div class="text-sm text-green-500">تراکنش‌های موفق</div>
      </div>
      
      <div class="bg-red-50 rounded-lg p-6">
        <div class="text-2xl font-bold text-red-600">{{ gateway.failedTransactions || 0 }}</div>
        <div class="text-sm text-red-500">تراکنش‌های ناموفق</div>
      </div>
      
      <div class="bg-blue-50 rounded-lg p-6">
        <div class="text-2xl font-bold text-blue-600">{{ formatCurrency(gateway.totalAmount || 0) }}</div>
        <div class="text-sm text-blue-500">کل مبلغ</div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'

// تعریف interface برای درگاه پرداخت
interface PaymentGateway {
  id: number
  name: string
  type: string
  merchantId: string
  status: string
  isTestMode: boolean
  isActive: boolean
  minAmount: number
  maxAmount: number
  totalTransactions: number
  successfulTransactions: number
  failedTransactions: number
  totalAmount: number
  apiEndpoints: {
    payment: string
    verify: string
  }
  apiKeys: {
    publicKey: string
    privateKey: string
  }
}

// متغیرهای reactive
const gateway = ref<PaymentGateway | null>(null)
const saving = ref(false)
const testing = ref(false)
const testResult = ref<{ success: boolean; message: string } | null>(null)

// فرم تنظیمات
const form = reactive({
  name: '',
  merchantId: '',
  privateKey: '',
  publicKey: '',
  paymentUrl: '',
  verifyUrl: '',
  minAmount: 1000,
  maxAmount: 50000000,
  isTestMode: false,
  isActive: true,
  autoVerify: false
})

// بارگذاری تنظیمات درگاه
const loadGateway = async () => {
  try {
    const response = await $fetch('/api/payments/admin/gateway/saman') as any
    if (response.success) {
      gateway.value = response.data
      // پر کردن فرم با داده‌های موجود
      form.name = gateway.value.name || 'سامان'
      form.merchantId = gateway.value.merchantId || ''
      form.privateKey = gateway.value.apiKeys?.privateKey || ''
      form.publicKey = gateway.value.apiKeys?.publicKey || ''
      form.paymentUrl = gateway.value.apiEndpoints?.payment || 'https://sep.shaparak.ir/api/v1/Payment/GetToken'
      form.verifyUrl = gateway.value.apiEndpoints?.verify || 'https://acquirer.samanepay.com/payments/referencepayment.asmx'
      form.minAmount = gateway.value.minAmount || 1000
      form.maxAmount = gateway.value.maxAmount || 50000000
      form.isTestMode = gateway.value.isTestMode || false
      form.isActive = gateway.value.isActive || true
    }
  } catch (error) {
    console.error('خطا در بارگذاری درگاه سامان:', error)
  }
}

// ذخیره تنظیمات
const saveSettings = async () => {
  saving.value = true
  try {
    const response = await $fetch('/api/payments/admin/gateway/saman', {
      method: 'POST',
      body: {
        name: form.name,
        type: 'saman',
        merchantId: form.merchantId,
        apiKeys: {
          publicKey: form.publicKey,
          privateKey: form.privateKey
        },
        apiEndpoints: {
          payment: form.paymentUrl,
          verify: form.verifyUrl
        },
        minAmount: form.minAmount,
        maxAmount: form.maxAmount,
        isTestMode: form.isTestMode,
        isActive: form.isActive,
        status: form.isActive ? 'active' : 'inactive'
      }
    }) as any

    if (response.success) {
      await loadGateway()
      // نمایش پیام موفقیت
      alert('تنظیمات با موفقیت ذخیره شد')
    }
  } catch (error: any) {
    console.error('خطا در ذخیره تنظیمات:', error)
    alert('خطا در ذخیره تنظیمات: ' + (error.statusMessage || 'خطای نامشخص'))
  } finally {
    saving.value = false
  }
}

// تست اتصال
const testConnection = async () => {
  testing.value = true
  testResult.value = null
  
  try {
    const response = await $fetch('/api/payments/admin/gateway/saman/test') as any
    testResult.value = {
      success: response.success,
      message: response.message
    }
  } catch (error: any) {
    testResult.value = {
      success: false,
      message: error.statusMessage || 'خطا در تست اتصال'
    }
  } finally {
    testing.value = false
  }
}

// بازنشانی فرم
const resetForm = () => {
  if (gateway.value) {
    form.name = gateway.value.name || 'سامان'
    form.merchantId = gateway.value.merchantId || ''
    form.privateKey = gateway.value.apiKeys?.privateKey || ''
    form.publicKey = gateway.value.apiKeys?.publicKey || ''
    form.paymentUrl = gateway.value.apiEndpoints?.payment || 'https://sep.shaparak.ir/api/v1/Payment/GetToken'
    form.verifyUrl = gateway.value.apiEndpoints?.verify || 'https://acquirer.samanepay.com/payments/referencepayment.asmx'
    form.minAmount = gateway.value.minAmount || 1000
    form.maxAmount = gateway.value.maxAmount || 50000000
    form.isTestMode = gateway.value.isTestMode || false
    form.isActive = gateway.value.isActive || true
  }
}

// فرمت کردن مبلغ
const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

// بارگذاری اولیه
onMounted(() => {
  loadGateway()
})
</script> 
