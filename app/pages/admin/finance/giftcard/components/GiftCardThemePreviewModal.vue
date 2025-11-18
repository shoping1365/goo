<template>
  <div class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
    <div class="relative top-10 mx-auto p-5 border w-11/12 max-w-6xl shadow-lg rounded-md bg-white">
      <!-- هدر مودال -->
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-lg font-semibold text-gray-900">پیش‌نمایش تم گیفت کارت</h3>
        <button
          @click="$emit('close')"
          class="text-gray-400 hover:text-gray-600"
        >
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
        </button>
      </div>

      <!-- محتوای مودال -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
        <!-- پیش‌نمایش تم -->
        <div class="space-y-6">
          <div>
            <h4 class="text-sm font-medium text-gray-700 mb-3">پیش‌نمایش تم</h4>
            <div
              class="p-8 rounded-lg text-center transform transition-all duration-300 hover:scale-105"
              :style="{
                backgroundColor: theme.backgroundColor,
                color: theme.textColor,
                fontFamily: theme.fontFamily,
                fontSize: theme.fontSize,
                border: `3px solid ${theme.borderColor}`,
                borderRadius: theme.borderRadius,
                boxShadow: theme.shadow
              }"
            >
              <!-- آیکون و عنوان -->
              <div class="text-6xl mb-4">{{ theme.icon }}</div>
              <h2 class="text-2xl font-bold mb-3">{{ theme.name }}</h2>
              <p class="text-sm opacity-80 mb-6">{{ theme.description }}</p>

              <!-- اطلاعات کارت -->
              <div class="space-y-4">
                <div class="p-6 rounded-lg" :style="{ backgroundColor: theme.primaryColor, color: 'white' }">
                  <div class="text-lg font-bold">مبلغ: ۵۰,۰۰۰ تومان</div>
                  <div class="text-sm opacity-90">کد: GC-2024-001</div>
                </div>

                <div class="grid grid-cols-2 gap-3 text-sm">
                  <div class="p-3 rounded-lg bg-white bg-opacity-20">
                    <div class="font-semibold">تاریخ انقضا</div>
                    <div>۱۴۰۳/۱۲/۲۹</div>
                  </div>
                  <div class="p-3 rounded-lg bg-white bg-opacity-20">
                    <div class="font-semibold">وضعیت</div>
                    <div>فعال</div>
                  </div>
                </div>

                <div class="p-3 rounded-lg bg-white bg-opacity-10">
                  <div class="text-sm">
                    <div class="font-semibold mb-1">پیام شخصی:</div>
                    <div class="italic">"تولدت مبارک! امیدوارم سالی پر از شادی و موفقیت داشته باشی."</div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- تنظیمات نمایش -->
          <div class="bg-gray-50 p-6 rounded-lg">
            <h4 class="text-sm font-medium text-gray-700 mb-3">تنظیمات نمایش</h4>
            <div class="grid grid-cols-2 gap-6">
              <div>
                <label class="block text-xs text-gray-600 mb-1">اندازه نمایش</label>
                <select
                  v-model="displaySize"
                  class="w-full px-2 py-1 text-sm border border-gray-300 rounded focus:outline-none focus:ring-1 focus:ring-blue-500"
                >
                  <option value="small">کوچک</option>
                  <option value="medium">متوسط</option>
                  <option value="large">بزرگ</option>
                </select>
              </div>
              <div>
                <label class="block text-xs text-gray-600 mb-1">نوع نمایش</label>
                <select
                  v-model="displayType"
                  class="w-full px-2 py-1 text-sm border border-gray-300 rounded focus:outline-none focus:ring-1 focus:ring-blue-500"
                >
                  <option value="card">کارت</option>
                  <option value="email">ایمیل</option>
                  <option value="print">چاپ</option>
                </select>
              </div>
            </div>
          </div>
        </div>

        <!-- اطلاعات تم -->
        <div class="space-y-6">
          <!-- اطلاعات کلی -->
          <div class="bg-white border border-gray-200 rounded-lg p-6">
            <h4 class="text-sm font-medium text-gray-700 mb-3">اطلاعات تم</h4>
            <div class="space-y-3">
              <div class="flex justify-between">
                <span class="text-sm text-gray-600">نام:</span>
                <span class="text-sm font-medium">{{ theme.name }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-sm text-gray-600">دسته‌بندی:</span>
                <span class="text-sm font-medium">{{ theme.category }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-sm text-gray-600">وضعیت:</span>
                <span
                  :class="{
                    'text-green-600': theme.status === 'active',
                    'text-red-600': theme.status === 'inactive',
                    'text-yellow-600': theme.status === 'pending'
                  }"
                  class="text-sm font-medium"
                >
                  {{ getStatusText(theme.status) }}
                </span>
              </div>
              <div class="flex justify-between">
                <span class="text-sm text-gray-600">آیکون:</span>
                <span class="text-2xl">{{ theme.icon }}</span>
              </div>
            </div>
          </div>

          <!-- رنگ‌ها -->
          <div class="bg-white border border-gray-200 rounded-lg p-6">
            <h4 class="text-sm font-medium text-gray-700 mb-3">رنگ‌های تم</h4>
            <div class="space-y-3">
              <div class="flex items-center justify-between">
                <span class="text-sm text-gray-600">رنگ اصلی:</span>
                <div class="flex items-center">
                  <div
                    class="w-6 h-6 rounded-full mr-2 border border-gray-300"
                    :style="{ backgroundColor: theme.primaryColor }"
                  ></div>
                  <span class="text-sm font-mono">{{ theme.primaryColor }}</span>
                </div>
              </div>
              <div class="flex items-center justify-between">
                <span class="text-sm text-gray-600">پس‌زمینه:</span>
                <div class="flex items-center">
                  <div
                    class="w-6 h-6 rounded-full mr-2 border border-gray-300"
                    :style="{ backgroundColor: theme.backgroundColor }"
                  ></div>
                  <span class="text-sm font-mono">{{ theme.backgroundColor }}</span>
                </div>
              </div>
              <div class="flex items-center justify-between">
                <span class="text-sm text-gray-600">متن:</span>
                <div class="flex items-center">
                  <div
                    class="w-6 h-6 rounded-full mr-2 border border-gray-300"
                    :style="{ backgroundColor: theme.textColor }"
                  ></div>
                  <span class="text-sm font-mono">{{ theme.textColor }}</span>
                </div>
              </div>
              <div class="flex items-center justify-between">
                <span class="text-sm text-gray-600">حاشیه:</span>
                <div class="flex items-center">
                  <div
                    class="w-6 h-6 rounded-full mr-2 border border-gray-300"
                    :style="{ backgroundColor: theme.borderColor }"
                  ></div>
                  <span class="text-sm font-mono">{{ theme.borderColor }}</span>
                </div>
              </div>
            </div>
          </div>

          <!-- تنظیمات فونت -->
          <div class="bg-white border border-gray-200 rounded-lg p-6">
            <h4 class="text-sm font-medium text-gray-700 mb-3">تنظیمات فونت</h4>
            <div class="space-y-3">
              <div class="flex justify-between">
                <span class="text-sm text-gray-600">فونت:</span>
                <span class="text-sm font-medium">{{ theme.fontFamily }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-sm text-gray-600">اندازه:</span>
                <span class="text-sm font-medium">{{ theme.fontSize }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-sm text-gray-600">شعاع گوشه:</span>
                <span class="text-sm font-medium">{{ theme.borderRadius }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-sm text-gray-600">انیمیشن:</span>
                <span class="text-sm font-medium">{{ getAnimationText(theme.animation) }}</span>
              </div>
            </div>
          </div>

          <!-- آمار استفاده -->
          <div class="bg-white border border-gray-200 rounded-lg p-6">
            <h4 class="text-sm font-medium text-gray-700 mb-3">آمار استفاده</h4>
            <div class="grid grid-cols-2 gap-6">
              <div class="text-center">
                <div class="text-2xl font-bold text-blue-600">{{ theme.usageCount }}</div>
                <div class="text-xs text-gray-500">استفاده شده</div>
              </div>
              <div class="text-center">
                <div class="text-2xl font-bold text-green-600">{{ theme.rating }}</div>
                <div class="text-xs text-gray-500">امتیاز</div>
              </div>
            </div>
          </div>

          <!-- دکمه‌های عملیات -->
          <div class="bg-white border border-gray-200 rounded-lg p-6">
            <h4 class="text-sm font-medium text-gray-700 mb-3">عملیات</h4>
            <div class="grid grid-cols-2 gap-3">
              <button
                @click="downloadTheme"
                class="px-3 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
              >
                دانلود
              </button>
              <button
                @click="shareTheme"
                class="px-3 py-2 bg-green-600 text-white text-sm font-medium rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2"
              >
                اشتراک‌گذاری
              </button>
              <button
                @click="printTheme"
                class="px-3 py-2 bg-purple-600 text-white text-sm font-medium rounded-lg hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-purple-500 focus:ring-offset-2"
              >
                چاپ
              </button>
              <button
                @click="copyThemeCode"
                class="px-3 py-2 bg-gray-600 text-white text-sm font-medium rounded-lg hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
              >
                کپی کد
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- دکمه بستن -->
      <div class="flex justify-end mt-6">
        <button
          @click="$emit('close')"
          class="px-4 py-2 bg-gray-300 text-gray-700 text-sm font-medium rounded-lg hover:bg-gray-400 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
        >
          بستن
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
// تعریف props و emits
const props = defineProps({
  theme: {
    type: Object,
    required: true
  }
})

const emit = defineEmits(['close'])

// تعریف متغیرهای reactive
const displaySize = ref('medium')
const displayType = ref('card')

// تابع دریافت متن وضعیت
const getStatusText = (status) => {
  const statusMap = {
    active: 'فعال',
    inactive: 'غیرفعال',
    pending: 'در انتظار'
  }
  return statusMap[status] || status
}

// تابع دریافت متن انیمیشن
const getAnimationText = (animation) => {
  const animationMap = {
    none: 'بدون انیمیشن',
    fadeIn: 'ظهور تدریجی',
    slideIn: 'لغزش',
    bounce: 'پرش',
    pulse: 'ضربان',
    rotate: 'چرخش'
  }
  return animationMap[animation] || animation
}

// تابع دانلود تم
const downloadTheme = () => {
  // ایجاد کد CSS برای تم
  const cssCode = `
/* تم ${props.theme.name} */
.gift-card-theme-${props.theme.id} {
  background-color: ${props.theme.backgroundColor};
  color: ${props.theme.textColor};
  font-family: ${props.theme.fontFamily};
  font-size: ${props.theme.fontSize};
  border: 2px solid ${props.theme.borderColor};
  border-radius: ${props.theme.borderRadius};
  box-shadow: ${props.theme.shadow};
}

.gift-card-theme-${props.theme.id} .primary-color {
  background-color: ${props.theme.primaryColor};
  color: white;
}

${props.theme.customCSS}
  `.trim()

  // ایجاد فایل برای دانلود
  const blob = new Blob([cssCode], { type: 'text/css' })
  const url = URL.createObjectURL(blob)
  const a = document.createElement('a')
  a.href = url
  a.download = `theme-${props.theme.id}.css`
  document.body.appendChild(a)
  a.click()
  document.body.removeChild(a)
  URL.revokeObjectURL(url)

  alert('فایل CSS تم با موفقیت دانلود شد')
}

// تابع اشتراک‌گذاری تم
const shareTheme = () => {
  const shareData = {
    title: `تم گیفت کارت: ${props.theme.name}`,
    text: props.theme.description,
    url: window.location.href
  }

  if (navigator.share) {
    navigator.share(shareData)
  } else {
    // کپی لینک به کلیپ‌بورد
    navigator.clipboard.writeText(window.location.href)
    alert('لینک تم در کلیپ‌بورد کپی شد')
  }
}

// تابع چاپ تم
const printTheme = () => {
  const printWindow = window.open('', '_blank')
  const printContent = `
    <!DOCTYPE html>
    <html dir="rtl" lang="fa">
    <head>
      <meta charset="UTF-8">
      <title>${props.theme.name}</title>
      <style>
        body { margin: 0; padding: 20px; font-family: ${props.theme.fontFamily}; }
        .gift-card {
          background-color: ${props.theme.backgroundColor};
          color: ${props.theme.textColor};
          border: 2px solid ${props.theme.borderColor};
          border-radius: ${props.theme.borderRadius};
          padding: 40px;
          text-align: center;
          max-width: 400px;
          margin: 0 auto;
        }
        .icon { font-size: 48px; margin-bottom: 20px; }
        .title { font-size: 24px; font-weight: bold; margin-bottom: 10px; }
        .description { font-size: 14px; opacity: 0.8; margin-bottom: 30px; }
        .amount {
          background-color: ${props.theme.primaryColor};
          color: white;
          padding: 15px;
          border-radius: 8px;
          margin-bottom: 20px;
        }
        @media print {
          body { padding: 0; }
          .gift-card { box-shadow: none; }
        }
      </style>
    </head>
    <body>
      <div class="gift-card">
        <div class="icon">${props.theme.icon}</div>
        <div class="title">${props.theme.name}</div>
        <div class="description">${props.theme.description}</div>
        <div class="amount">
          <div style="font-size: 18px; font-weight: bold;">مبلغ: ۵۰,۰۰۰ تومان</div>
          <div style="font-size: 12px; opacity: 0.9;">کد: GC-2024-001</div>
        </div>
        <div style="font-size: 12px; opacity: 0.7;">
          تاریخ انقضا: ۱۴۰۳/۱۲/۲۹
        </div>
      </div>
    </body>
    </html>
  `
  
  printWindow.document.write(printContent)
  printWindow.document.close()
  printWindow.print()
}

// تابع کپی کد تم
const copyThemeCode = () => {
  const themeCode = JSON.stringify(props.theme, null, 2)
  navigator.clipboard.writeText(themeCode)
  alert('کد تم در کلیپ‌بورد کپی شد')
}
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 
