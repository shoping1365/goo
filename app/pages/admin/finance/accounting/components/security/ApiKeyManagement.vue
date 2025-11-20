<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">مدیریت کلیدهای API</h4>
        <p class="text-sm text-gray-600 mt-1">مدیریت کلیدهای API و مجوزهای دسترسی</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="generateNewApiKey"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
          </svg>
          تولید کلید جدید
        </button>
      </div>
    </div>

    <!-- آمار کلیدها -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">آمار کلیدهای API</h5>
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">کل کلیدها</h6>
            <div class="w-3 h-3 bg-blue-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-blue-600">{{ apiStats.total }}</div>
          <div class="text-xs text-gray-500 mt-1">کلید</div>
        </div>

        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">کلیدهای فعال</h6>
            <div class="w-3 h-3 bg-green-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-green-600">{{ apiStats.active }}</div>
          <div class="text-xs text-gray-500 mt-1">کلید</div>
        </div>

        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">درخواست‌های امروز</h6>
            <div class="w-3 h-3 bg-yellow-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-yellow-600">{{ apiStats.todayRequests }}</div>
          <div class="text-xs text-gray-500 mt-1">درخواست</div>
        </div>

        <div class="p-6 border border-gray-200 rounded-lg">
          <div class="flex items-center justify-between mb-2">
            <h6 class="text-sm font-medium text-gray-900">نرخ موفقیت</h6>
            <div class="w-3 h-3 bg-green-500 rounded-full"></div>
          </div>
          <div class="text-2xl font-bold text-green-600">{{ apiStats.successRate }}%</div>
          <div class="text-xs text-gray-500 mt-1">موفق</div>
        </div>
      </div>
    </div>

    <!-- لیست کلیدهای API -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">کلیدهای API</h5>
      </div>
      <div class="p-6">
        <div class="space-y-4">
          <div
            v-for="key in apiKeys"
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
                <div class="text-sm font-medium text-gray-900">{{ key.lastUsed }}</div>
                <div class="text-xs text-gray-500">آخرین استفاده</div>
              </div>
              <div class="text-right">
                <div class="text-sm font-medium text-gray-900">{{ key.requests }}</div>
                <div class="text-xs text-gray-500">درخواست</div>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <button
                  class="text-blue-600 hover:text-blue-800"
                  @click="viewKey(key)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                  </svg>
                </button>
                <button
                  class="text-green-600 hover:text-green-800"
                  @click="editKey(key)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                  </svg>
                </button>
                <button
                  class="text-red-600 hover:text-red-800"
                  @click="revokeKey(key)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728L5.636 5.636m12.728 12.728L18.364 5.636M5.636 18.364l12.728-12.728" />
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات API -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">تنظیمات API</h5>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر درخواست در دقیقه</label>
          <input
            v-model="apiSettings.rateLimit"
            type="number"
            min="10"
            max="1000"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">مدت اعتبار کلید</label>
          <select
            v-model="apiSettings.keyExpiry"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="30">۳۰ روز</option>
            <option value="90">۹۰ روز</option>
            <option value="180">۱۸۰ روز</option>
            <option value="365">۱ سال</option>
            <option value="0">بدون انقضا</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">طول کلید API</label>
          <select
            v-model="apiSettings.keyLength"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="32">۳۲ کاراکتر</option>
            <option value="64">۶۴ کاراکتر</option>
            <option value="128">۱۲۸ کاراکتر</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نوع کلید</label>
          <select
            v-model="apiSettings.keyType"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="alphanumeric">حروف و اعداد</option>
            <option value="hex">هگزادسیمال</option>
            <option value="base64">Base64</option>
          </select>
        </div>
      </div>

      <div class="mt-4 space-y-3">
        <label class="flex items-center">
          <input
            v-model="apiSettings.requireHttps"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">اجبار استفاده از HTTPS</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="apiSettings.logRequests"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">ثبت تمام درخواست‌ها</span>
        </label>

        <label class="flex items-center">
          <input
            v-model="apiSettings.notifyOnExpiry"
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <span class="mr-3 text-sm text-gray-700">اعلان قبل از انقضا</span>
        </label>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// آمار API
const apiStats = ref({
  total: 8,
  active: 6,
  todayRequests: 1247,
  successRate: 98.5
})

// کلیدهای API
const apiKeys = ref([
  {
    id: 1,
    name: 'کلید اصلی سیستم',
    description: 'کلید API برای دسترسی کامل به سیستم',
    status: 'active',
    lastUsed: '۲ ساعت پیش',
    requests: 15420
  },
  {
    id: 2,
    name: 'کلید گزارش‌گیری',
    description: 'کلید API برای دسترسی به گزارش‌ها',
    status: 'active',
    lastUsed: '۱ روز پیش',
    requests: 3240
  },
  {
    id: 3,
    name: 'کلید تست',
    description: 'کلید API برای تست و توسعه',
    status: 'inactive',
    lastUsed: '۱ هفته پیش',
    requests: 1250
  }
])

// تنظیمات API
const apiSettings = ref({
  rateLimit: 100,
  keyExpiry: 90,
  keyLength: 64,
  keyType: 'alphanumeric',
  requireHttps: true,
  logRequests: true,
  notifyOnExpiry: true
})

// رنگ‌های وضعیت
const getKeyStatusColor = (status: string) => {
  const colors = {
    active: 'bg-green-500',
    inactive: 'bg-gray-500',
    expired: 'bg-red-500'
  }
  return colors[status] || 'bg-gray-500'
}

// عملیات
const generateNewApiKey = () => {
  console.log('تولید کلید API جدید')
}

const viewKey = (key: any) => {
  console.log('مشاهده کلید:', key)
}

const editKey = (key: any) => {
  console.log('ویرایش کلید:', key)
}

const revokeKey = (key: any) => {
  console.log('لغو کلید:', key)
}
</script>

<!--
  کامپوننت مدیریت کلیدهای API
  شامل:
  1. آمار کلیدها
  2. لیست کلیدهای API
  3. تنظیمات API
  4. طراحی مدرن و کاملاً ریسپانسیو
--> 
