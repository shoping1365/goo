<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200">
    <div class="p-6 border-b border-gray-200">
      <div class="flex items-center justify-between">
        <div>
          <h2 class="text-lg font-semibold text-gray-900">تنظیمات کلی درگاه‌های پرداخت</h2>
          <p class="text-gray-600 mt-1">مدیریت تنظیمات عمومی و پیش‌فرض درگاه‌های پرداخت</p>
        </div>
        <TemplateButton 
          :disabled="saving" 
          bg-gradient="bg-gradient-to-r from-blue-500 to-indigo-600" 
          text-color="text-white"
          border-color="border border-blue-500"
          hover-class="hover:from-blue-600 hover:to-indigo-700"
          focus-class="focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
          size="medium"
          @click="saveSettings"
        >
          <span v-if="saving" class="i-heroicons-arrow-path animate-spin ml-2"></span>
          <span v-else class="i-heroicons-check ml-2"></span>
          {{ saving ? 'در حال ذخیره...' : 'ذخیره تنظیمات' }}
        </TemplateButton>
      </div>
    </div>

    <div class="p-6">
      <!-- تب‌های تنظیمات -->
      <div class="flex border-b border-gray-200 mb-6">
        <button
v-for="tab in settingsTabs" :key="tab.value" :class="['px-4 py-2 -mb-px font-medium text-sm focus:outline-none', activeTab === tab.value ? 'border-b-2 border-blue-600 text-blue-700' : 'text-gray-500 hover:text-blue-600']"
          @click="activeTab = tab.value">
          {{ tab.label }}
        </button>
      </div>

      <!-- محتوای تب‌ها -->
      <div>
        <!-- تنظیمات عمومی -->
        <div v-if="activeTab === 'general'" class="space-y-6">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">درگاه پیش‌فرض</label>
              <select v-model="settings.defaultGateway" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                <option value="">انتخاب کنید</option>
                <option v-for="gateway in availableGateways" :key="gateway.id" :value="gateway.id">
                  {{ gateway.name }}
                </option>
              </select>
              <p class="text-xs text-gray-500 mt-1">درگاه پیش‌فرض برای پرداخت‌های جدید</p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر تعداد درگاه‌های همزمان</label>
              <input v-model.number="settings.maxConcurrentGateways" type="number" min="1" max="10" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
              <p class="text-xs text-gray-500 mt-1">تعداد درگاه‌هایی که همزمان می‌توانند فعال باشند</p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">زمان انقضای پیش‌فرض تراکنش (دقیقه)</label>
              <input v-model.number="settings.defaultTransactionExpiry" type="number" min="5" max="1440" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر تلاش مجاز پیش‌فرض</label>
              <input v-model.number="settings.defaultMaxAttempts" type="number" min="1" max="10" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
            </div>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">حداقل مبلغ پیش‌فرض (تومان)</label>
              <input v-model.number="settings.defaultMinAmount" type="number" min="0" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر مبلغ پیش‌فرض (تومان)</label>
              <input v-model.number="settings.defaultMaxAmount" type="number" min="0" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
            </div>
          </div>

          <div class="space-y-4">
            <div class="flex items-center">
              <input id="autoRetry" v-model="settings.autoRetryFailed" type="checkbox" class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500">
              <label for="autoRetry" class="mr-2 text-sm font-medium text-gray-700">تلاش مجدد خودکار برای تراکنش‌های ناموفق</label>
            </div>
            <div class="flex items-center">
              <input id="fallback" v-model="settings.enableFallback" type="checkbox" class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500">
              <label for="fallback" class="mr-2 text-sm font-medium text-gray-700">فعال‌سازی درگاه پشتیبان</label>
            </div>
            <div class="flex items-center">
              <input id="testMode" v-model="settings.enableTestMode" type="checkbox" class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500">
              <label for="testMode" class="mr-2 text-sm font-medium text-gray-700">فعال‌سازی حالت تست</label>
            </div>
          </div>
        </div>

        <!-- تنظیمات امنیتی -->
        <div v-if="activeTab === 'security'" class="space-y-6">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">IP Whitelist عمومی</label>
              <textarea v-model="settings.globalIpWhitelist" rows="3" placeholder="مثال: 192.168.1.1, 10.0.0.0/24" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"></textarea>
              <p class="text-xs text-gray-500 mt-1">آدرس‌های IP مجاز برای همه درگاه‌ها</p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر تلاش مجاز در ساعت</label>
              <input v-model.number="settings.maxAttemptsPerHour" type="number" min="1" max="100" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">زمان مسدودیت IP (دقیقه)</label>
              <input v-model.number="settings.ipBlockDuration" type="number" min="5" max="1440" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">حداقل طول کلید API</label>
              <input v-model.number="settings.minApiKeyLength" type="number" min="8" max="64" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
            </div>
          </div>

          <div class="space-y-4">
            <div class="flex items-center">
              <input id="rateLimit" v-model="settings.enableRateLimiting" type="checkbox" class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500">
              <label for="rateLimit" class="mr-2 text-sm font-medium text-gray-700">فعال‌سازی محدودیت نرخ درخواست</label>
            </div>
            <div class="flex items-center">
              <input id="logTransactions" v-model="settings.logAllTransactions" type="checkbox" class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500">
              <label for="logTransactions" class="mr-2 text-sm font-medium text-gray-700">ثبت تمام تراکنش‌ها در لاگ</label>
            </div>
          </div>
        </div>

        <!-- تنظیمات مالی -->
        <div v-if="activeTab === 'financial'" class="space-y-6">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">کارمزد پیش‌فرض (%)</label>
              <input v-model.number="settings.defaultFee" type="number" min="0" max="100" step="0.1" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">حداقل کارمزد (تومان)</label>
              <input v-model.number="settings.minFee" type="number" min="0" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر کارمزد (تومان)</label>
              <input v-model.number="settings.maxFee" type="number" min="0" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">واحد ارز پیش‌فرض</label>
              <select v-model="settings.defaultCurrency" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                <option value="IRR">ریال ایران</option>
                <option value="USD">دلار آمریکا</option>
                <option value="EUR">یورو</option>
                <option value="GBP">پوند انگلیس</option>
              </select>
            </div>
          </div>

          <div class="space-y-4">
            <div class="flex items-center">
              <input id="autoFee" v-model="settings.autoCalculateFee" type="checkbox" class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500">
              <label for="autoFee" class="mr-2 text-sm font-medium text-gray-700">محاسبه خودکار کارمزد</label>
            </div>
            <div class="flex items-center">
              <input id="roundFees" v-model="settings.roundFees" type="checkbox" class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500">
              <label for="roundFees" class="mr-2 text-sm font-medium text-gray-700">گرد کردن کارمزدها</label>
            </div>
          </div>
        </div>

        <!-- تنظیمات اعلان‌ها -->
        <div v-if="activeTab === 'notifications'" class="space-y-6">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">ایمیل ادمین</label>
              <input v-model="settings.adminEmail" type="email" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">شماره تلفن ادمین</label>
              <input v-model="settings.adminPhone" type="tel" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
            </div>
          </div>

          <div class="space-y-4">
            <div class="flex items-center">
              <input id="gatewayFailure" v-model="settings.notifyOnGatewayFailure" type="checkbox" class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500">
              <label for="gatewayFailure" class="mr-2 text-sm font-medium text-gray-700">اعلان در صورت خرابی درگاه</label>
            </div>
            <div class="flex items-center">
              <input id="highVolume" v-model="settings.notifyOnHighVolume" type="checkbox" class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500">
              <label for="highVolume" class="mr-2 text-sm font-medium text-gray-700">اعلان در صورت حجم بالای تراکنش</label>
            </div>
            <div class="flex items-center">
              <input id="suspicious" v-model="settings.notifyOnSuspiciousActivity" type="checkbox" class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500">
              <label for="suspicious" class="mr-2 text-sm font-medium text-gray-700">اعلان در صورت فعالیت مشکوک</label>
            </div>
            <div class="flex items-center">
              <input id="lowBalance" v-model="settings.notifyOnLowBalance" type="checkbox" class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500">
              <label for="lowBalance" class="mr-2 text-sm font-medium text-gray-700">اعلان در صورت کمبود موجودی</label>
            </div>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">حد آستانه حجم تراکنش</label>
              <input v-model.number="settings.highVolumeThreshold" type="number" min="100" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
              <p class="text-xs text-gray-500 mt-1">تعداد تراکنش در ساعت</p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">حد آستانه موجودی کم</label>
              <input v-model.number="settings.lowBalanceThreshold" type="number" min="0" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
              <p class="text-xs text-gray-500 mt-1">مبلغ به تومان</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import TemplateButton from '~/components/common/TemplateButton.vue'

// تعریف interface برای تنظیمات
interface GlobalSettings {
  // تنظیمات عمومی
  defaultGateway: string
  maxConcurrentGateways: number
  defaultTransactionExpiry: number
  defaultMaxAttempts: number
  defaultMinAmount: number
  defaultMaxAmount: number
  autoRetryFailed: boolean
  enableFallback: boolean
  enableTestMode: boolean

  // تنظیمات امنیتی
  globalIpWhitelist: string
  maxAttemptsPerHour: number
  ipBlockDuration: number
  minApiKeyLength: number
  enableRateLimiting: boolean
  logAllTransactions: boolean

  // تنظیمات مالی
  defaultFee: number
  minFee: number
  maxFee: number
  defaultCurrency: string
  autoCalculateFee: boolean
  roundFees: boolean

  // تنظیمات اعلان‌ها
  adminEmail: string
  adminPhone: string
  notifyOnGatewayFailure: boolean
  notifyOnHighVolume: boolean
  notifyOnSuspiciousActivity: boolean
  notifyOnLowBalance: boolean
  highVolumeThreshold: number
  lowBalanceThreshold: number
}

interface AvailableGateway {
  id: string
  name: string
  englishName: string
}

// متغیرهای reactive
const activeTab = ref('general')
const saving = ref(false)

// تب‌های تنظیمات
const settingsTabs = [
  { value: 'general', label: 'تنظیمات عمومی' },
  { value: 'security', label: 'امنیت' },
  { value: 'financial', label: 'مالی' },
  { value: 'notifications', label: 'اعلان‌ها' }
]

// تنظیمات پیش‌فرض
const settings = ref<GlobalSettings>({
  // تنظیمات عمومی
  defaultGateway: '',
  maxConcurrentGateways: 5,
  defaultTransactionExpiry: 30,
  defaultMaxAttempts: 3,
  defaultMinAmount: 10000,
  defaultMaxAmount: 50000000,
  autoRetryFailed: true,
  enableFallback: true,
  enableTestMode: false,

  // تنظیمات امنیتی
  globalIpWhitelist: '',
  maxAttemptsPerHour: 50,
  ipBlockDuration: 60,
  minApiKeyLength: 16,
  enableRateLimiting: true,
  logAllTransactions: true,

  // تنظیمات مالی
  defaultFee: 1.5,
  minFee: 1000,
  maxFee: 100000,
  defaultCurrency: 'IRR',
  autoCalculateFee: true,
  roundFees: true,

  // تنظیمات اعلان‌ها
  adminEmail: 'admin@example.com',
  adminPhone: '+989123456789',
  notifyOnGatewayFailure: true,
  notifyOnHighVolume: true,
  notifyOnSuspiciousActivity: true,
  notifyOnLowBalance: true,
  highVolumeThreshold: 1000,
  lowBalanceThreshold: 1000000
})

// درگاه‌های موجود
const availableGateways = ref<AvailableGateway[]>([
  { id: '1', name: 'زرین‌پال', englishName: 'ZarinPal' },
  { id: '2', name: 'نکست‌پی', englishName: 'NextPay' },
  { id: '3', name: 'آیدی‌پی', englishName: 'IDPay' },
  { id: '4', name: 'پی‌ایر', englishName: 'PayIR' },
  { id: '5', name: 'پی‌پینگ', englishName: 'PayPing' }
])

// توابع عملیات
const saveSettings = async () => {
  saving.value = true
  try {
    // TODO: فراخوانی API برای ذخیره تنظیمات
    // await $fetch('/api/admin/payment-gateways/global-settings', {
    //   method: 'PUT',
    //   body: settings.value
    // })
    
    await new Promise(resolve => setTimeout(resolve, 1000))

  } catch (error) {
    console.error('خطا در ذخیره تنظیمات:', error)
  } finally {
    saving.value = false
  }
}

// بارگذاری تنظیمات
onMounted(() => {
  // TODO: فراخوانی API برای دریافت تنظیمات
  // const response = await $fetch('/api/admin/payment-gateways/global-settings')
  // settings.value = response.settings
  // availableGateways.value = response.gateways
})
</script>

<!--
  کامپوننت تنظیمات کلی درگاه‌های پرداخت
  شامل:
  1. تنظیمات عمومی (درگاه پیش‌فرض، محدودیت‌ها)
  2. تنظیمات امنیتی (IP Whitelist، محدودیت‌ها)
  3. تنظیمات مالی (کارمزد، ارز)
  4. تنظیمات اعلان‌ها (ایمیل، تلفن، آستانه‌ها)
  5. مدیریت state و validation
  6. طراحی ریسپانسیو
  توضیحات کامل به فارسی در کد
--> 
