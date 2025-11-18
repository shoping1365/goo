<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">تنظیمات رمزگذاری</h4>
        <p class="text-sm text-gray-600 mt-1">مدیریت الگوریتم‌های رمزگذاری و کلیدها</p>
      </div>
    </div>

    <!-- آمار رمزگذاری -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">وضعیت رمزگذاری</h5>
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">داده‌های رمزگذاری شده</h6>
            <div class="w-3 h-3 bg-green-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-green-600">{{ encryptionStats.encrypted }}</div>
          <div class="text-xs text-gray-500 mt-1">رکورد</div>
        </div>

        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">الگوریتم فعال</h6>
            <div class="w-3 h-3 bg-blue-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-blue-600">{{ encryptionStats.algorithm }}</div>
          <div class="text-xs text-gray-500 mt-1">AES-256</div>
        </div>

        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">کلیدهای فعال</h6>
            <div class="w-3 h-3 bg-yellow-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-yellow-600">{{ encryptionStats.activeKeys }}</div>
          <div class="text-xs text-gray-500 mt-1">کلید</div>
        </div>

        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">آخرین به‌روزرسانی</h6>
            <div class="w-3 h-3 bg-green-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-green-600">{{ encryptionStats.lastUpdate }}</div>
          <div class="text-xs text-gray-500 mt-1">ساعت پیش</div>
        </div>
      </div>
    </div>

    <!-- تنظیمات الگوریتم -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">تنظیمات الگوریتم رمزگذاری</h5>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">الگوریتم رمزگذاری</label>
          <select
            v-model="encryptionSettings.algorithm"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="aes-256">AES-256</option>
            <option value="aes-192">AES-192</option>
            <option value="aes-128">AES-128</option>
            <option value="chacha20">ChaCha20</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حالت رمزگذاری</label>
          <select
            v-model="encryptionSettings.mode"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="gcm">GCM (توصیه شده)</option>
            <option value="cbc">CBC</option>
            <option value="ctr">CTR</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">طول کلید</label>
          <select
            v-model="encryptionSettings.keyLength"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="256">۲۵۶ بیت</option>
            <option value="192">۱۹۲ بیت</option>
            <option value="128">۱۲۸ بیت</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">طول Salt</label>
          <select
            v-model="encryptionSettings.saltLength"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="32">۳۲ بایت</option>
            <option value="16">۱۶ بایت</option>
            <option value="8">۸ بایت</option>
          </select>
        </div>
      </div>
    </div>

    <!-- مدیریت کلیدها -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">مدیریت کلیدهای رمزگذاری</h5>
      </div>
      <div class="p-6">
        <div class="space-y-4">
          <div
            v-for="key in encryptionKeys"
            :key="key.id"
            class="flex items-center justify-between p-6 border border-gray-200 rounded-lg hover:bg-gray-50"
          >
            <div class="flex items-center">
              <div
                class="w-3 h-3 rounded-full ml-3"
                :class="getKeyStatusColor(key.status)"
              ></div>
              <div>
                <h6 class="text-sm font-medium text-gray-900">{{ key.name }}</h6>
                <p class="text-xs text-gray-500">{{ key.description }}</p>
              </div>
            </div>
            <div class="flex items-center space-x-4 space-x-reverse">
              <div class="text-right">
                <div class="text-sm font-medium text-gray-900">{{ key.created }}</div>
                <div class="text-xs text-gray-500">تاریخ ایجاد</div>
              </div>
              <div class="text-right">
                <div class="text-sm font-medium text-gray-900">{{ key.usage }}</div>
                <div class="text-xs text-gray-500">استفاده</div>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <button
                  @click="rotateKey(key)"
                  class="text-blue-600 hover:text-blue-800"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                  </svg>
                </button>
                <button
                  @click="backupKey(key)"
                  class="text-green-600 hover:text-green-800"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7H5a2 2 0 00-2 2v9a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-3m-1 4l-3 3m0 0l-3-3m3 3V4" />
                  </svg>
                </button>
                <button
                  @click="deleteKey(key)"
                  class="text-red-600 hover:text-red-800"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات پیشرفته -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">تنظیمات پیشرفته</h5>
      <div class="space-y-4">
        <label class="flex items-center">
          <input
            v-model="encryptionSettings.autoRotate"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">چرخش خودکار کلیدها</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="encryptionSettings.encryptBackups"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">رمزگذاری پشتیبان‌ها</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="encryptionSettings.encryptLogs"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">رمزگذاری لاگ‌ها</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="encryptionSettings.hardwareAcceleration"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">شتاب سخت‌افزاری</span>
        </label>
      </div>

      <div class="mt-6 grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">فاصله چرخش کلید</label>
          <select
            v-model="encryptionSettings.rotationInterval"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="30">۳۰ روز</option>
            <option value="90">۹۰ روز</option>
            <option value="180">۱۸۰ روز</option>
            <option value="365">۱ سال</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر کلیدهای قدیمی</label>
          <input
            v-model="encryptionSettings.maxOldKeys"
            type="number"
            min="1"
            max="10"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// آمار رمزگذاری
const encryptionStats = ref({
  encrypted: 15420,
  algorithm: 'AES-256',
  activeKeys: 3,
  lastUpdate: 6
})

// کلیدهای رمزگذاری
const encryptionKeys = ref([
  {
    id: 1,
    name: 'کلید اصلی سیستم',
    description: 'کلید رمزگذاری اصلی برای داده‌های حساس',
    status: 'active',
    created: '۱۵ دی ۱۴۰۲',
    usage: 12450
  },
  {
    id: 2,
    name: 'کلید پشتیبان',
    description: 'کلید رمزگذاری پشتیبان',
    status: 'active',
    created: '۱۰ دی ۱۴۰۲',
    usage: 3240
  },
  {
    id: 3,
    name: 'کلید قدیمی',
    description: 'کلید رمزگذاری قدیمی (در حال انقضا)',
    status: 'expiring',
    created: '۱ آذر ۱۴۰۲',
    usage: 1250
  }
])

// تنظیمات رمزگذاری
const encryptionSettings = ref({
  algorithm: 'aes-256',
  mode: 'gcm',
  keyLength: 256,
  saltLength: 32,
  autoRotate: true,
  encryptBackups: true,
  encryptLogs: false,
  hardwareAcceleration: true,
  rotationInterval: 90,
  maxOldKeys: 3
})

// رنگ‌های وضعیت
const getKeyStatusColor = (status: string) => {
  const colors = {
    active: 'bg-green-500',
    expiring: 'bg-yellow-500',
    expired: 'bg-red-500'
  }
  return colors[status] || 'bg-gray-500'
}

// عملیات
const rotateKey = (key: any) => {
  console.log('چرخش کلید:', key)
}

const backupKey = (key: any) => {
  console.log('پشتیبان‌گیری کلید:', key)
}

const deleteKey = (key: any) => {
  console.log('حذف کلید:', key)
}
</script>

<!--
  کامپوننت تنظیمات رمزگذاری
  شامل:
  1. آمار رمزگذاری
  2. تنظیمات الگوریتم
  3. مدیریت کلیدها
  4. تنظیمات پیشرفته
  5. طراحی مدرن و کاملاً ریسپانسیو
--> 
