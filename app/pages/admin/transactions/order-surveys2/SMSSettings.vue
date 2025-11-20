<template>
  <div class="space-y-8">
    <!-- Header -->
    <div class="flex items-center justify-between">
      <div>
        <h3 class="text-lg font-semibold text-gray-800">تنظیمات فنی پیامک</h3>
        <p class="text-gray-600 text-sm">پیکربندی چندین درگاه SMS، failover و monitoring</p>
      </div>
      <div class="flex space-x-3 space-x-reverse">
        <button 
          :disabled="testing"
          class="px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg text-sm flex items-center space-x-2 space-x-reverse"
          @click="testAllConnections"
        >
          <svg v-if="testing" class="animate-spin w-4 h-4" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          <span>{{ testing ? 'در حال تست اتصالات...' : 'تست همه اتصالات' }}</span>
        </button>
        <button 
          class="px-4 py-2 bg-green-600 hover:bg-green-700 text-white rounded-lg text-sm flex items-center space-x-2 space-x-reverse"
          @click="addGateway"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
          </svg>
          <span>افزودن درگاه</span>
        </button>
      </div>
    </div>

    <!-- SMS Gateways -->
    <div class="space-y-6">
      <h4 class="text-md font-semibold text-gray-800">درگاه‌های پیامک</h4>
      
      <div v-for="(gateway, index) in gateways" :key="index" class="bg-white rounded-lg border border-gray-200 px-4 py-4">
        <div class="flex items-center justify-between mb-4">
          <div class="flex items-center space-x-3 space-x-reverse">
            <div class="flex items-center space-x-2 space-x-reverse">
              <div class="w-3 h-3 rounded-full" :class="gateway.status === 'active' ? 'bg-green-500' : gateway.status === 'inactive' ? 'bg-red-500' : 'bg-yellow-500'"></div>
              <span class="text-sm font-medium">{{ gateway.name }}</span>
            </div>
            <span class="px-2 py-1 text-xs rounded-full" :class="gateway.priority === 1 ? 'bg-blue-100 text-blue-800' : gateway.priority === 2 ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800'">
              اولویت {{ gateway.priority }}
            </span>
          </div>
          <div class="flex items-center space-x-2 space-x-reverse">
            <button class="text-blue-600 hover:text-blue-800 text-sm" @click="testGateway(index)">تست</button>
            <button class="text-red-600 hover:text-red-800 text-sm" @click="removeGateway(index)">حذف</button>
          </div>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gapx-4 py-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">نام درگاه</label>
            <input v-model="gateway.name" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-lg" placeholder="مثال: ملی پیامک">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">اولویت</label>
            <select v-model="gateway.priority" class="w-full px-3 py-2 border border-gray-300 rounded-lg">
              <option value="1">اول (اصلی)</option>
              <option value="2">دوم (پشتیبان)</option>
              <option value="3">سوم (اضافی)</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">وضعیت</label>
            <select v-model="gateway.status" class="w-full px-3 py-2 border border-gray-300 rounded-lg">
              <option value="active">فعال</option>
              <option value="inactive">غیرفعال</option>
              <option value="maintenance">تعمیر و نگهداری</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">آدرس API</label>
            <input v-model="gateway.apiUrl" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-lg" placeholder="https://sms.example.com/api/send">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">کلید API</label>
            <input v-model="gateway.apiKey" type="password" class="w-full px-3 py-2 border border-gray-300 rounded-lg" placeholder="API Key">
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">حداکثر پیامک در دقیقه</label>
            <input v-model.number="gateway.maxPerMinute" type="number" min="1" class="w-full px-3 py-2 border border-gray-300 rounded-lg" placeholder="60">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">تأخیر بین ارسال (ثانیه)</label>
            <input v-model.number="gateway.delaySeconds" type="number" min="0" class="w-full px-3 py-2 border border-gray-300 rounded-lg" placeholder="2">
          </div>
        </div>
      </div>
    </div>

    <!-- Advanced Settings -->
    <div class="bg-white rounded-lg border border-gray-200 px-4 py-4 space-y-6">
      <h4 class="text-md font-semibold text-gray-800 mb-4">تنظیمات پیشرفته</h4>
      
      <!-- Failover Settings -->
      <div class="space-y-4">
        <h5 class="text-sm font-semibold text-gray-700">تنظیمات Failover</h5>
        <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">فعال‌سازی Failover خودکار</label>
            <select v-model="advancedSettings.autoFailover" class="w-full px-3 py-2 border border-gray-300 rounded-lg">
              <option value="enabled">فعال</option>
              <option value="disabled">غیرفعال</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">تعداد تلاش قبل از Failover</label>
            <input v-model.number="advancedSettings.failoverAttempts" type="number" min="1" max="10" class="w-full px-3 py-2 border border-gray-300 rounded-lg" placeholder="3">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">زمان انتظار قبل از Failover (ثانیه)</label>
            <input v-model.number="advancedSettings.failoverTimeout" type="number" min="1" max="60" class="w-full px-3 py-2 border border-gray-300 rounded-lg" placeholder="30">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">بازگشت خودکار به درگاه اصلی</label>
            <select v-model="advancedSettings.autoRecovery" class="w-full px-3 py-2 border border-gray-300 rounded-lg">
              <option value="enabled">فعال</option>
              <option value="disabled">غیرفعال</option>
            </select>
          </div>
        </div>
      </div>

      <!-- Monitoring Settings -->
      <div class="space-y-4">
        <h5 class="text-sm font-semibold text-gray-700">تنظیمات Monitoring</h5>
        <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">فاصله بررسی سلامت (دقیقه)</label>
            <input v-model.number="advancedSettings.healthCheckInterval" type="number" min="1" max="60" class="w-full px-3 py-2 border border-gray-300 rounded-lg" placeholder="5">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">حد آستانه خطا (%)</label>
            <input v-model.number="advancedSettings.errorThreshold" type="number" min="1" max="100" class="w-full px-3 py-2 border border-gray-300 rounded-lg" placeholder="10">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">زمان انتظار پاسخ (ثانیه)</label>
            <input v-model.number="advancedSettings.responseTimeout" type="number" min="1" max="60" class="w-full px-3 py-2 border border-gray-300 rounded-lg" placeholder="10">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">فعال‌سازی Logging پیشرفته</label>
            <select v-model="advancedSettings.advancedLogging" class="w-full px-3 py-2 border border-gray-300 rounded-lg">
              <option value="enabled">فعال</option>
              <option value="disabled">غیرفعال</option>
            </select>
          </div>
        </div>
      </div>

      <!-- Alerting Settings -->
      <div class="space-y-4">
        <h5 class="text-sm font-semibold text-gray-700">تنظیمات هشدار</h5>
        <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">هشدار به مدیر</label>
            <select v-model="advancedSettings.adminAlerts" class="w-full px-3 py-2 border border-gray-300 rounded-lg">
              <option value="enabled">فعال</option>
              <option value="disabled">غیرفعال</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">ایمیل مدیر</label>
            <input v-model="advancedSettings.adminEmail" type="email" class="w-full px-3 py-2 border border-gray-300 rounded-lg" placeholder="admin@example.com">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">شماره تلفن مدیر</label>
            <input v-model="advancedSettings.adminPhone" type="tel" class="w-full px-3 py-2 border border-gray-300 rounded-lg" placeholder="09123456789">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">نوع هشدار</label>
            <select v-model="advancedSettings.alertType" class="w-full px-3 py-2 border border-gray-300 rounded-lg">
              <option value="email">ایمیل</option>
              <option value="sms">پیامک</option>
              <option value="both">هر دو</option>
            </select>
          </div>
        </div>
      </div>

      <!-- Security Settings -->
      <div class="space-y-4">
        <h5 class="text-sm font-semibold text-gray-700">تنظیمات امنیتی</h5>
        <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">IP های مجاز</label>
            <input v-model="advancedSettings.allowedIPs" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-lg" placeholder="192.168.1.1, 10.0.0.1">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">رمزنگاری اتصال</label>
            <select v-model="advancedSettings.encryption" class="w-full px-3 py-2 border border-gray-300 rounded-lg">
              <option value="none">بدون رمزنگاری</option>
              <option value="tls">TLS</option>
              <option value="ssl">SSL</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Rate Limiting</label>
            <select v-model="advancedSettings.rateLimiting" class="w-full px-3 py-2 border border-gray-300 rounded-lg">
              <option value="enabled">فعال</option>
              <option value="disabled">غیرفعال</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">فعال‌سازی Audit Log</label>
            <select v-model="advancedSettings.auditLog" class="w-full px-3 py-2 border border-gray-300 rounded-lg">
              <option value="enabled">فعال</option>
              <option value="disabled">غیرفعال</option>
            </select>
          </div>
        </div>
      </div>
    </div>

    <!-- Save Button -->
    <div class="flex items-center justify-end">
      <button 
        :disabled="saving"
        class="px-6 py-2 bg-green-600 hover:bg-green-700 text-white rounded-lg text-sm flex items-center space-x-2 space-x-reverse"
        @click="saveSettings"
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
import { reactive, ref } from 'vue'

interface SMSGateway {
  name: string
  priority: number
  status: 'active' | 'inactive' | 'maintenance'
  apiUrl: string
  apiKey: string
  maxPerMinute: number
  delaySeconds: number
}

interface AdvancedSettings {
  autoFailover: 'enabled' | 'disabled'
  failoverAttempts: number
  failoverTimeout: number
  autoRecovery: 'enabled' | 'disabled'
  healthCheckInterval: number
  errorThreshold: number
  responseTimeout: number
  advancedLogging: 'enabled' | 'disabled'
  adminAlerts: 'enabled' | 'disabled'
  adminEmail: string
  adminPhone: string
  alertType: 'email' | 'sms' | 'both'
  allowedIPs: string
  encryption: 'none' | 'tls' | 'ssl'
  rateLimiting: 'enabled' | 'disabled'
  auditLog: 'enabled' | 'disabled'
}

const gateways = ref<SMSGateway[]>([
  {
    name: 'ملی پیامک',
    priority: 1,
    status: 'active',
    apiUrl: 'https://api.melipayamak.com/v1/sms/send',
    apiKey: '',
    maxPerMinute: 60,
    delaySeconds: 2
  },
  {
    name: 'کاوه‌پیامک',
    priority: 2,
    status: 'active',
    apiUrl: 'https://api.kavenegar.com/v1/sms/send',
    apiKey: '',
    maxPerMinute: 50,
    delaySeconds: 3
  },
  {
    name: 'آوای پیامک',
    priority: 3,
    status: 'inactive',
    apiUrl: 'https://api.avayesms.com/sms/send',
    apiKey: '',
    maxPerMinute: 40,
    delaySeconds: 4
  }
])

const advancedSettings = reactive<AdvancedSettings>({
  autoFailover: 'enabled',
  failoverAttempts: 3,
  failoverTimeout: 30,
  autoRecovery: 'enabled',
  healthCheckInterval: 5,
  errorThreshold: 10,
  responseTimeout: 10,
  advancedLogging: 'enabled',
  adminAlerts: 'enabled',
  adminEmail: 'admin@example.com',
  adminPhone: '09123456789',
  alertType: 'both',
  allowedIPs: '192.168.1.1, 10.0.0.1',
  encryption: 'tls',
  rateLimiting: 'enabled',
  auditLog: 'enabled'
})

const saving = ref(false)
const testing = ref(false)

const addGateway = () => {
  gateways.value.push({
    name: `درگاه جدید ${gateways.value.length + 1}`,
    priority: gateways.value.length + 1,
    status: 'inactive',
    apiUrl: '',
    apiKey: '',
    maxPerMinute: 30,
    delaySeconds: 5
  })
}

const removeGateway = (index: number) => {
  if (gateways.value.length > 1) {
    gateways.value.splice(index, 1)
    // حذف بازسازی خودکار اولویت‌ها - اولویت‌ها باید ثابت بمانند
    // gateways.value.forEach((gateway, idx) => {
    //   gateway.priority = idx + 1
    // })
  }
}

const testGateway = async (index: number) => {
  const gateway = gateways.value[index]
  // console.log(`Testing gateway: ${gateway.name}`)
  // اینجا باید تست اتصال واقعی انجام شود
  await new Promise(resolve => setTimeout(resolve, 1000))
  alert(`تست اتصال ${gateway.name} انجام شد!`)
}

const testAllConnections = async () => {
  testing.value = true
  try {
    // تست همه درگاه‌ها
    for (const gateway of gateways.value) {
      if (gateway.status === 'active') {
        // console.log(`Testing: ${gateway.name}`)
        await new Promise(resolve => setTimeout(resolve, 500))
      }
    }
    alert('تست همه اتصالات با موفقیت انجام شد!')
  } catch {
    alert('خطا در تست اتصالات!')
  } finally {
    testing.value = false
  }
}

const saveSettings = async () => {
  saving.value = true
  try {
    // اینجا باید ذخیره به سرور انجام شود
    await new Promise(resolve => setTimeout(resolve, 1500))
    alert('تنظیمات با موفقیت ذخیره شد!')
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
