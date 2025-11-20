<template>
  <div class="bg-white rounded-lg shadow-md p-6">
    <div class="flex items-center justify-between mb-6">
      <h2 class="text-xl font-bold text-gray-800">تنظیمات درگاه پرداخت استرایپ</h2>
      <div class="flex items-center space-x-2 space-x-reverse">
        <button
          :disabled="testing"
          class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 disabled:opacity-50"
          @click="testConnection"
        >
          {{ testing ? 'در حال تست...' : 'تست اتصال' }}
        </button>
        <button
          :disabled="saving"
          class="px-4 py-2 bg-green-500 text-white rounded-md hover:bg-green-600 disabled:opacity-50"
          @click="saveSettings"
        >
          {{ saving ? 'در حال ذخیره...' : 'ذخیره تنظیمات' }}
        </button>
      </div>
    </div>

    <div v-if="message" :class="messageClass" class="mb-4 p-3 rounded-md">
      {{ message }}
    </div>

    <form class="space-y-6" @submit.prevent="saveSettings">
      <!-- تنظیمات اصلی -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            نام درگاه
          </label>
          <input
            v-model="settings.name"
            type="text"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            placeholder="استرایپ"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            وضعیت
          </label>
          <select
            v-model="settings.is_active"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            <option :value="true">فعال</option>
            <option :value="false">غیرفعال</option>
          </select>
        </div>
      </div>

      <!-- تنظیمات API -->
      <div class="border-t pt-6">
        <h3 class="text-lg font-semibold text-gray-800 mb-4">تنظیمات API</h3>
        
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Publishable Key
            </label>
            <input
              v-model="settings.publishable_key"
              type="text"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="pk_test_..."
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Secret Key
            </label>
            <input
              v-model="settings.secret_key"
              type="password"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="sk_test_..."
            />
          </div>
        </div>

        <div class="mt-6">
          <label class="block text-sm font-medium text-gray-700 mb-2">
            حالت اجرا
          </label>
          <select
            v-model="settings.mode"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            <option value="test">تست (Test)</option>
            <option value="live">عملیاتی (Live)</option>
          </select>
        </div>
      </div>

      <!-- تنظیمات اضافی -->
      <div class="border-t pt-6">
        <h3 class="text-lg font-semibold text-gray-800 mb-4">تنظیمات اضافی</h3>
        
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              ارز پیش‌فرض
            </label>
            <select
              v-model="settings.currency"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            >
              <option value="usd">دلار آمریکا (USD)</option>
              <option value="eur">یورو (EUR)</option>
              <option value="gbp">پوند انگلیس (GBP)</option>
              <option value="jpy">ین ژاپن (JPY)</option>
            </select>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              زبان پیش‌فرض
            </label>
            <select
              v-model="settings.language"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            >
              <option value="en">انگلیسی</option>
              <option value="fa">فارسی</option>
              <option value="auto">خودکار</option>
            </select>
          </div>
        </div>

        <div class="mt-6 space-y-4">
          <label class="flex items-center">
            <input
              v-model="settings.auto_capture"
              type="checkbox"
              class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
            />
            <span class="mr-2 text-sm text-gray-700">تسویه خودکار پرداخت‌ها</span>
          </label>

          <label class="flex items-center">
            <input
              v-model="settings.enable_webhooks"
              type="checkbox"
              class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
            />
            <span class="mr-2 text-sm text-gray-700">فعال‌سازی Webhook</span>
          </label>
        </div>
      </div>
    </form>
  </div>
</template>

<script setup>
const { $toast } = useNuxtApp()

// Reactive data
const settings = ref({
  name: 'استرایپ',
  is_active: false,
  publishable_key: '',
  secret_key: '',
  mode: 'test',
  currency: 'usd',
  language: 'en',
  auto_capture: false,
  enable_webhooks: true
})

const saving = ref(false)
const testing = ref(false)
const message = ref('')
const messageClass = ref('')

// Load settings on mount
onMounted(async () => {
  await loadSettings()
})

// Load gateway settings
const loadSettings = async () => {
  try {
    const response = await $fetch('/api/admin/payment-gateways/stripe.settings.get')
    if (response.success) {
      Object.assign(settings.value, response.data)
    }
  } catch (error) {
    showMessage('خطا در بارگذاری تنظیمات', 'error')
  }
}

// Save settings
const saveSettings = async () => {
  saving.value = true
  try {
    const response = await $fetch('/api/admin/payment-gateways/stripe.settings.put', {
      method: 'PUT',
      body: settings.value
    })
    
    if (response.success) {
      showMessage('تنظیمات با موفقیت ذخیره شد', 'success')
    } else {
      showMessage(response.message || 'خطا در ذخیره تنظیمات', 'error')
    }
  } catch (error) {
    showMessage('خطا در ذخیره تنظیمات', 'error')
  } finally {
    saving.value = false
  }
}

// Test connection
const testConnection = async () => {
  testing.value = true
  try {
    const response = await $fetch('/api/admin/payment-gateways/stripe.test-connection.post', {
      method: 'POST'
    })
    
    if (response.success) {
      showMessage('اتصال به درگاه با موفقیت برقرار شد', 'success')
    } else {
      showMessage(response.message || 'خطا در اتصال به درگاه', 'error')
    }
  } catch (error) {
    showMessage('خطا در تست اتصال', 'error')
  } finally {
    testing.value = false
  }
}

// Show message
const showMessage = (msg, type) => {
  message.value = msg
  messageClass.value = type === 'success' 
    ? 'bg-green-100 border border-green-400 text-green-700' 
    : 'bg-red-100 border border-red-400 text-red-700'
  
  setTimeout(() => {
    message.value = ''
    messageClass.value = ''
  }, 5000)
}
</script> 
