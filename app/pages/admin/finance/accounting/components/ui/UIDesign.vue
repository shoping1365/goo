<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">طراحی مدرن و ریسپانسیو</h4>
        <p class="text-sm text-gray-600 mt-1">تنظیمات ظاهری و ریسپانسیو رابط کاربری</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          @click="previewDesign"
          class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
          </svg>
          پیش‌نمایش
        </button>
      </div>
    </div>

    <!-- تنظیمات تم -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">تنظیمات تم</h5>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تم اصلی</label>
          <div class="grid grid-cols-3 gap-2">
            <button
              v-for="theme in themes"
              :key="theme.id"
              @click="selectedTheme = theme.id"
              class="relative p-3 border-2 rounded-lg transition-all"
              :class="selectedTheme === theme.id ? 'border-blue-500' : 'border-gray-200 hover:border-gray-300'"
            >
              <div class="w-full h-8 rounded" :class="theme.colors.primary"></div>
              <div class="mt-2 text-xs text-center">{{ theme.name }}</div>
              <div v-if="selectedTheme === theme.id" class="absolute top-1 right-1">
                <svg class="w-4 h-4 text-blue-500" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
                </svg>
              </div>
            </button>
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حالت نمایش</label>
          <div class="space-y-2">
            <label class="flex items-center">
              <input
                v-model="uiSettings.darkMode"
                type="checkbox"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <span class="mr-3 text-sm text-gray-700">حالت تاریک</span>
            </label>
            <label class="flex items-center">
              <input
                v-model="uiSettings.compactMode"
                type="checkbox"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <span class="mr-3 text-sm text-gray-700">حالت فشرده</span>
            </label>
            <label class="flex items-center">
              <input
                v-model="uiSettings.animations"
                type="checkbox"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <span class="mr-3 text-sm text-gray-700">انیمیشن‌ها</span>
            </label>
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">اندازه فونت</label>
          <select
            v-model="uiSettings.fontSize"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="small">کوچک</option>
            <option value="medium">متوسط</option>
            <option value="large">بزرگ</option>
            <option value="xlarge">خیلی بزرگ</option>
          </select>
        </div>
      </div>
    </div>

    <!-- تنظیمات رنگ‌ها -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">تنظیمات رنگ‌ها</h5>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">رنگ اصلی</label>
          <div class="flex items-center space-x-3 space-x-reverse">
            <input
              v-model="uiSettings.primaryColor"
              type="color"
              class="w-12 h-10 border border-gray-300 rounded-md"
            />
            <span class="text-sm text-gray-600">{{ uiSettings.primaryColor }}</span>
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">رنگ ثانویه</label>
          <div class="flex items-center space-x-3 space-x-reverse">
            <input
              v-model="uiSettings.secondaryColor"
              type="color"
              class="w-12 h-10 border border-gray-300 rounded-md"
            />
            <span class="text-sm text-gray-600">{{ uiSettings.secondaryColor }}</span>
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">رنگ پس‌زمینه</label>
          <div class="flex items-center space-x-3 space-x-reverse">
            <input
              v-model="uiSettings.backgroundColor"
              type="color"
              class="w-12 h-10 border border-gray-300 rounded-md"
            />
            <span class="text-sm text-gray-600">{{ uiSettings.backgroundColor }}</span>
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">رنگ متن</label>
          <div class="flex items-center space-x-3 space-x-reverse">
            <input
              v-model="uiSettings.textColor"
              type="color"
              class="w-12 h-10 border border-gray-300 rounded-md"
            />
            <span class="text-sm text-gray-600">{{ uiSettings.textColor }}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات ریسپانسیو -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">تنظیمات ریسپانسیو</h5>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نقطه شکست موبایل</label>
          <input
            v-model="uiSettings.mobileBreakpoint"
            type="number"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            placeholder="768"
          />
          <p class="text-xs text-gray-500 mt-1">پیکسل</p>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نقطه شکست تبلت</label>
          <input
            v-model="uiSettings.tabletBreakpoint"
            type="number"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            placeholder="1024"
          />
          <p class="text-xs text-gray-500 mt-1">پیکسل</p>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">اندازه ستون‌ها در موبایل</label>
          <select
            v-model="uiSettings.mobileColumns"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="1">۱ ستون</option>
            <option value="2">۲ ستون</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">اندازه ستون‌ها در تبلت</label>
          <select
            v-model="uiSettings.tabletColumns"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="2">۲ ستون</option>
            <option value="3">۳ ستون</option>
            <option value="4">۴ ستون</option>
          </select>
        </div>
      </div>
    </div>

    <!-- پیش‌نمایش ریسپانسیو -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">پیش‌نمایش ریسپانسیو</h5>
      <div class="flex items-center space-x-4 space-x-reverse">
        <button
          v-for="device in devices"
          :key="device.id"
          @click="selectedDevice = device.id"
          class="flex items-center px-4 py-2 border rounded-md transition-colors"
          :class="selectedDevice === device.id ? 'border-blue-500 bg-blue-50' : 'border-gray-300 hover:bg-gray-50'"
        >
          <svg class="w-4 h-4 ml-2" :class="device.icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="device.path" />
          </svg>
          {{ device.name }}
        </button>
      </div>

      <div class="mt-4 border border-gray-300 rounded-lg overflow-hidden">
        <div class="bg-gray-100 px-4 py-2 border-b border-gray-300">
          <div class="flex items-center justify-between">
            <span class="text-sm font-medium text-gray-700">{{ getDeviceName(selectedDevice) }}</span>
            <div class="flex items-center space-x-2 space-x-reverse">
              <div class="w-3 h-3 bg-red-500 rounded-full"></div>
              <div class="w-3 h-3 bg-yellow-500 rounded-full"></div>
              <div class="w-3 h-3 bg-green-500 rounded-full"></div>
            </div>
          </div>
        </div>
        <div class="p-6 bg-white" :class="getDeviceClass(selectedDevice)">
          <div class="grid gap-6" :class="getGridClass(selectedDevice)">
            <div class="bg-blue-100 p-6 rounded-lg text-center">کارت ۱</div>
            <div class="bg-green-100 p-6 rounded-lg text-center">کارت ۲</div>
            <div class="bg-yellow-100 p-6 rounded-lg text-center">کارت ۳</div>
            <div class="bg-purple-100 p-6 rounded-lg text-center">کارت ۴</div>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات اضافی -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">تنظیمات اضافی</h5>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">فاصله بین المان‌ها</label>
          <select
            v-model="uiSettings.spacing"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="compact">فشرده</option>
            <option value="normal">عادی</option>
            <option value="loose">گسترده</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">گردی گوشه‌ها</label>
          <select
            v-model="uiSettings.borderRadius"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="none">بدون گردی</option>
            <option value="small">کوچک</option>
            <option value="medium">متوسط</option>
            <option value="large">بزرگ</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">سایه المان‌ها</label>
          <select
            v-model="uiSettings.shadow"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="none">بدون سایه</option>
            <option value="small">سایه کوچک</option>
            <option value="medium">سایه متوسط</option>
            <option value="large">سایه بزرگ</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">سرعت انیمیشن</label>
          <select
            v-model="uiSettings.animationSpeed"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="slow">کند</option>
            <option value="normal">عادی</option>
            <option value="fast">سریع</option>
          </select>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// تم‌های موجود
const themes = ref([
  {
    id: 'blue',
    name: 'آبی',
    colors: { primary: 'bg-blue-500' }
  },
  {
    id: 'green',
    name: 'سبز',
    colors: { primary: 'bg-green-500' }
  },
  {
    id: 'purple',
    name: 'بنفش',
    colors: { primary: 'bg-purple-500' }
  }
])

// تنظیمات UI
const uiSettings = ref({
  darkMode: false,
  compactMode: false,
  animations: true,
  fontSize: 'medium',
  primaryColor: '#3B82F6',
  secondaryColor: '#6B7280',
  backgroundColor: '#FFFFFF',
  textColor: '#1F2937',
  mobileBreakpoint: 768,
  tabletBreakpoint: 1024,
  mobileColumns: 1,
  tabletColumns: 2,
  spacing: 'normal',
  borderRadius: 'medium',
  shadow: 'medium',
  animationSpeed: 'normal'
})

// انتخاب تم
const selectedTheme = ref('blue')

// دستگاه‌های پیش‌نمایش
const devices = ref([
  {
    id: 'mobile',
    name: 'موبایل',
    icon: 'text-gray-600',
    path: 'M12 18h.01M8 21h8a2 2 0 002-2V5a2 2 0 00-2-2H8a2 2 0 00-2 2v14a2 2 0 002 2z'
  },
  {
    id: 'tablet',
    name: 'تبلت',
    icon: 'text-gray-600',
    path: 'M12 18h.01M7 21h10a2 2 0 002-2V5a2 2 0 00-2-2H7a2 2 0 00-2 2v14a2 2 0 002 2z'
  },
  {
    id: 'desktop',
    name: 'دسکتاپ',
    icon: 'text-gray-600',
    path: 'M9.75 17L9 20l-1 1h8l-1-1-.75-3M3 13h18M5 17h14a2 2 0 002-2V5a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z'
  }
])

const selectedDevice = ref('desktop')

// نام دستگاه
const getDeviceName = (deviceId: string) => {
  const device = devices.value.find(d => d.id === deviceId)
  return device ? device.name : 'دسکتاپ'
}

// کلاس دستگاه
const getDeviceClass = (deviceId: string) => {
  const classes = {
    mobile: 'max-w-sm mx-auto',
    tablet: 'max-w-md mx-auto',
    desktop: 'w-full'
  }
  return classes[deviceId] || 'w-full'
}

// کلاس گرید
const getGridClass = (deviceId: string) => {
  const classes = {
    mobile: 'grid-cols-1',
    tablet: 'grid-cols-2',
    desktop: 'grid-cols-4'
  }
  return classes[deviceId] || 'grid-cols-4'
}

// پیش‌نمایش طراحی
const previewDesign = () => {
  console.log('پیش‌نمایش طراحی:', uiSettings.value)
}
</script>

<!--
  کامپوننت طراحی مدرن و ریسپانسیو
  شامل:
  1. تنظیمات تم
  2. تنظیمات رنگ‌ها
  3. تنظیمات ریسپانسیو
  4. پیش‌نمایش دستگاه‌ها
  5. تنظیمات اضافی
  6. طراحی مدرن و کاملاً ریسپانسیو
--> 
