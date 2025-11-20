<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">مدیریت کلیدهای API</h4>
        <p class="text-sm text-gray-600 mt-1">مدیریت و نظارت بر کلیدهای API نرم‌افزارهای حسابداری</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
          @click="showCreateModal = true"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
          </svg>
          ایجاد کلید جدید
        </button>
        <button
          class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="refreshApiKeys"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
          بروزرسانی
        </button>
      </div>
    </div>

    <!-- آمار کلیدهای API -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="bg-blue-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-blue-600">{{ apiStats.totalKeys }}</div>
            <div class="text-sm text-blue-700">کل کلیدها</div>
          </div>
          <div class="w-10 h-10 bg-blue-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 7a2 2 0 012 2m4 0a6 6 0 01-7.743 5.743L11 17H9v2H7v2H4a1 1 0 01-1-1v-2.586a1 1 0 01.293-.707l5.964-5.964A6 6 0 1121 9z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-green-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-green-600">{{ apiStats.activeKeys }}</div>
            <div class="text-sm text-green-700">کلیدهای فعال</div>
          </div>
          <div class="w-10 h-10 bg-green-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-yellow-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-yellow-600">{{ apiStats.expiredKeys }}</div>
            <div class="text-sm text-yellow-700">کلیدهای منقضی شده</div>
          </div>
          <div class="w-10 h-10 bg-yellow-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-red-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-red-600">{{ apiStats.revokedKeys }}</div>
            <div class="text-sm text-red-700">کلیدهای لغو شده</div>
          </div>
          <div class="w-10 h-10 bg-red-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- جدول کلیدهای API -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">کلیدهای API</h5>
      </div>
      <div class="overflow-x-auto">
        <table class="w-full text-sm">
          <thead>
            <tr class="border-b border-gray-200 bg-gray-50">
              <th class="text-right py-3 px-4 font-medium text-gray-600">نام کلید</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">نرم‌افزار</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">کلید API</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">وضعیت</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">تاریخ انقضا</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">آخرین استفاده</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">عملیات</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="key in apiKeys" :key="key.id" class="border-b border-gray-100 hover:bg-gray-50">
              <td class="py-3 px-4">
                <div>
                  <div class="font-medium text-gray-900">{{ key.name }}</div>
                  <div class="text-xs text-gray-500">{{ key.description }}</div>
                </div>
              </td>
              <td class="py-3 px-4">
                <div class="flex items-center gap-2">
                  <img :src="key.softwareLogo" :alt="key.softwareName" class="w-6 h-6 rounded" />
                  <span class="text-sm text-gray-900">{{ key.softwareName }}</span>
                </div>
              </td>
              <td class="py-3 px-4">
                <div class="flex items-center space-x-2 space-x-reverse">
                  <code class="text-xs bg-gray-100 px-2 py-1 rounded">{{ key.maskedKey }}</code>
                  <button
                    class="p-1 text-blue-600 hover:text-blue-800"
                    :title="key.showKey ? 'مخفی کردن' : 'نمایش'"
                    @click="toggleKeyVisibility(key)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path v-if="key.showKey" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.878 9.878L3 3m6.878 6.878L21 21" />
                      <path v-else stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                    </svg>
                  </button>
                </div>
                <div v-if="key.showKey" class="mt-2">
                  <code class="text-xs bg-yellow-100 px-2 py-1 rounded break-all">{{ key.fullKey }}</code>
                </div>
              </td>
              <td class="py-3 px-4">
                <span
                  class="px-2 py-1 rounded-full text-xs font-medium"
                  :class="getKeyStatusClass(key.status)"
                >
                  {{ getKeyStatusLabel(key.status) }}
                </span>
              </td>
              <td class="py-3 px-4">
                <div class="text-sm text-gray-900">{{ key.expiryDate }}</div>
                <div class="text-xs text-gray-500">{{ key.daysLeft }}</div>
              </td>
              <td class="py-3 px-4">
                <div class="text-sm text-gray-900">{{ key.lastUsed }}</div>
                <div class="text-xs text-gray-500">{{ key.usageCount }} بار استفاده</div>
              </td>
              <td class="py-3 px-4">
                <div class="flex items-center space-x-2 space-x-reverse">
                  <button
                    class="p-1 text-green-600 hover:text-green-800"
                    title="تولید مجدد"
                    @click="regenerateKey(key)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                    </svg>
                  </button>
                  <button
                    class="p-1 text-blue-600 hover:text-blue-800"
                    title="ویرایش"
                    @click="editKey(key)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                    </svg>
                  </button>
                  <button
                    class="p-1 text-red-600 hover:text-red-800"
                    title="لغو کلید"
                    @click="revokeKey(key)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728L5.636 5.636m12.728 12.728L18.364 5.636M5.636 18.364l12.728-12.728" />
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- نمودار استفاده از کلیدها -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h5 class="text-md font-medium text-gray-900 mb-4">استفاده از کلیدهای API</h5>
        <div class="space-y-4">
          <div v-for="usage in apiKeyUsage" :key="usage.software" class="flex items-center justify-between">
            <div class="flex items-center space-x-3 space-x-reverse">
              <img :src="usage.logo" :alt="usage.software" class="w-6 h-6 rounded" />
              <span class="text-sm text-gray-700">{{ usage.software }}</span>
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <div class="w-24 bg-gray-200 rounded-full h-2">
                <div
                  class="h-2 rounded-full bg-blue-500"
                  :style="{ width: usage.percentage + '%' }"
                ></div>
              </div>
              <span class="text-sm text-gray-900 w-8">{{ usage.percentage }}%</span>
            </div>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h5 class="text-md font-medium text-gray-900 mb-4">وضعیت کلیدها</h5>
        <div class="space-y-4">
          <div v-for="status in keyStatusDistribution" :key="status.status" class="flex items-center justify-between">
            <div class="flex items-center space-x-3 space-x-reverse">
              <div class="w-3 h-3 rounded-full" :class="getKeyStatusColor(status.status)"></div>
              <span class="text-sm text-gray-700">{{ getKeyStatusLabel(status.status) }}</span>
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <div class="w-24 bg-gray-200 rounded-full h-2">
                <div
                  class="h-2 rounded-full"
                  :class="getKeyStatusColor(status.status)"
                  :style="{ width: status.percentage + '%' }"
                ></div>
              </div>
              <span class="text-sm text-gray-900 w-8">{{ status.percentage }}%</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- مودال ایجاد کلید جدید -->
    <div v-if="showCreateModal" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
      <div class="relative top-20 mx-auto p-5 border w-full max-w-md sm:max-w-lg md:max-w-xl shadow-lg rounded-md bg-white">
        <div class="mt-3">
          <h3 class="text-lg font-medium text-gray-900 mb-4">ایجاد کلید API جدید</h3>
          <form class="space-y-4" @submit.prevent="createApiKey">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نام کلید</label>
              <input
                v-model="newKey.name"
                type="text"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="نام کلید API"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات</label>
              <textarea
                v-model="newKey.description"
                rows="3"
                class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="توضیحات کلید"
              ></textarea>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نرم‌افزار</label>
              <select
                v-model="newKey.software"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="">انتخاب نرم‌افزار</option>
                <option value="helo">هلو</option>
                <option value="parsian">پارسیان</option>
                <option value="sepidar">سپیدار</option>
                <option value="ghias">قیاس</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">مدت اعتبار (روز)</label>
              <input
                v-model.number="newKey.validityDays"
                type="number"
                min="1"
                max="365"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>

            <div class="flex items-center justify-end space-x-3 space-x-reverse">
              <button
                type="button"
                class="px-4 py-2 border border-gray-300 rounded-md text-sm font-medium text-gray-700 hover:bg-gray-50"
                @click="showCreateModal = false"
              >
                انصراف
              </button>
              <button
                type="submit"
                class="px-4 py-2 bg-blue-600 text-white rounded-md text-sm font-medium hover:bg-blue-700"
              >
                ایجاد کلید
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// وضعیت مودال
const showCreateModal = ref(false)

// آمار کلیدهای API
const apiStats = ref({
  totalKeys: 12,
  activeKeys: 8,
  expiredKeys: 3,
  revokedKeys: 1
})

// کلیدهای API
const apiKeys = ref([
  {
    id: 1,
    name: 'کلید اصلی هلو',
    description: 'کلید API برای اتصال به نرم‌افزار هلو',
    softwareName: 'هلو',
    softwareLogo: '/images/helo-logo.png',
    maskedKey: 'helo_****_****_****',
    fullKey: 'helo_api_key_1234567890abcdef',
    showKey: false,
    status: 'active',
    expiryDate: '۱۴۰۳/۰۶/۱۵',
    daysLeft: '۳۰ روز باقی‌مانده',
    lastUsed: 'امروز ۱۴:۳۰',
    usageCount: 1250
  },
  {
    id: 2,
    name: 'کلید پارسیان',
    description: 'کلید API برای اتصال به نرم‌افزار پارسیان',
    softwareName: 'پارسیان',
    softwareLogo: '/images/parsian-logo.png',
    maskedKey: 'parsian_****_****',
    fullKey: 'parsian_api_key_abcdef123456',
    showKey: false,
    status: 'active',
    expiryDate: '۱۴۰۳/۰۵/۲۰',
    daysLeft: '۵ روز باقی‌مانده',
    lastUsed: 'دیروز ۱۸:۴۵',
    usageCount: 890
  },
  {
    id: 3,
    name: 'کلید تست سپیدار',
    description: 'کلید تست برای نرم‌افزار سپیدار',
    softwareName: 'سپیدار',
    softwareLogo: '/images/sepidar-logo.png',
    maskedKey: 'sepidar_****_****',
    fullKey: 'sepidar_test_key_789xyz',
    showKey: false,
    status: 'expired',
    expiryDate: '۱۴۰۳/۰۴/۱۰',
    daysLeft: 'منقضی شده',
    lastUsed: 'هفته گذشته',
    usageCount: 45
  }
])

// استفاده از کلیدهای API
const apiKeyUsage = ref([
  { software: 'هلو', logo: '/images/helo-logo.png', percentage: 45 },
  { software: 'پارسیان', logo: '/images/parsian-logo.png', percentage: 35 },
  { software: 'سپیدار', logo: '/images/sepidar-logo.png', percentage: 15 },
  { software: 'قیاس', logo: '/images/ghias-logo.png', percentage: 5 }
])

// توزیع وضعیت کلیدها
const keyStatusDistribution = ref([
  { status: 'active', percentage: 67 },
  { status: 'expired', percentage: 25 },
  { status: 'revoked', percentage: 8 }
])

// کلید جدید
const newKey = ref({
  name: '',
  description: '',
  software: '',
  validityDays: 90
})

// کلاس وضعیت کلید
const getKeyStatusClass = (status: string) => {
  const classes = {
    active: 'bg-green-100 text-green-700',
    expired: 'bg-yellow-100 text-yellow-700',
    revoked: 'bg-red-100 text-red-700'
  }
  return classes[status] || 'bg-gray-100 text-gray-700'
}

// رنگ وضعیت کلید
const getKeyStatusColor = (status: string) => {
  const colors = {
    active: 'bg-green-500',
    expired: 'bg-yellow-500',
    revoked: 'bg-red-500'
  }
  return colors[status] || 'bg-gray-500'
}

// برچسب وضعیت کلید
const getKeyStatusLabel = (status: string) => {
  const labels = {
    active: 'فعال',
    expired: 'منقضی شده',
    revoked: 'لغو شده'
  }
  return labels[status] || status
}

// تغییر نمایش کلید
const toggleKeyVisibility = (key: any) => {
  key.showKey = !key.showKey
}

// بروزرسانی کلیدهای API
const refreshApiKeys = () => {
  // TODO: بروزرسانی کلیدهای API
  console.log('بروزرسانی کلیدهای API')
}

// تولید مجدد کلید
const regenerateKey = (key: any) => {
  // TODO: تولید مجدد کلید
  console.log('تولید مجدد کلید:', key)
}

// ویرایش کلید
const editKey = (key: any) => {
  // TODO: ویرایش کلید
  console.log('ویرایش کلید:', key)
}

// لغو کلید
const revokeKey = (key: any) => {
  // TODO: لغو کلید
  console.log('لغو کلید:', key)
}

// ایجاد کلید جدید
const createApiKey = () => {
  // TODO: ایجاد کلید جدید
  console.log('ایجاد کلید جدید:', newKey.value)
  showCreateModal.value = false
  
  // پاک کردن فرم
  newKey.value = {
    name: '',
    description: '',
    software: '',
    validityDays: 90
  }
}
</script>

<!--
  کامپوننت مدیریت کلیدهای API
  شامل:
  1. آمار کلیدهای API
  2. جدول کلیدها
  3. نمودارهای تحلیلی
  4. مودال ایجاد کلید جدید
  5. عملیات مدیریت کلیدها
  6. طراحی مدرن و کاملاً ریسپانسیو
--> 
