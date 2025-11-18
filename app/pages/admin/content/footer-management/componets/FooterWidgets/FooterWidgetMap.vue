<template>
  <div class="footer-map-widget" :style="widgetStyle">
    <div class="text-center">
      <!-- عنوان بخش -->
      <h4 v-if="title" class="text-white font-semibold mb-4 text-lg">{{ title }}</h4>
      
      <!-- نقشه -->
      <div v-if="latitude && longitude" class="space-y-3">
        <!-- پیش‌نمایش نقشه -->
        <div class="bg-white bg-opacity-10 rounded-lg p-3">
          <div class="w-full h-40 bg-gray-300 rounded flex items-center justify-center relative overflow-hidden">
            <!-- نقشه ساده -->
            <div class="absolute inset-0 bg-gradient-to-br from-blue-400 to-green-400 flex items-center justify-center">
              <div class="text-white text-center">
                <div class="text-2xl mb-2">📍</div>
                <div class="text-sm font-medium">موقعیت ما</div>
                <div class="text-xs opacity-80 mt-1">
                  {{ latitude.toFixed(6) }}, {{ longitude.toFixed(6) }}
                </div>
              </div>
            </div>
            
            <!-- نشانگر موقعیت -->
            <div class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2">
              <div class="w-6 h-6 bg-red-500 rounded-full border-2 border-white shadow-lg animate-pulse"></div>
            </div>
          </div>
          
          <!-- دکمه‌های عملیات -->
          <div class="flex space-x-2 space-x-reverse mt-3">
            <button
              @click="openGoogleMaps"
              class="flex-1 px-3 py-2 bg-blue-600 hover:bg-blue-700 text-white text-sm rounded transition-colors"
            >
              مشاهده در نقشه
            </button>
            <button
              @click="copyCoordinates"
              class="px-3 py-2 bg-gray-600 hover:bg-gray-700 text-white text-sm rounded transition-colors"
              :title="copyButtonText"
            >
              {{ copyButtonText }}
            </button>
          </div>
        </div>
        
        <!-- اطلاعات موقعیت -->
        <div v-if="showLocationInfo" class="text-white text-sm space-y-2">
          <div v-if="address" class="flex items-center justify-center space-x-2 space-x-reverse">
            <span class="text-lg">📍</span>
            <p>{{ address }}</p>
          </div>
          
          <div v-if="city || state" class="flex items-center justify-center space-x-2 space-x-reverse">
            <span class="text-lg">🏙️</span>
            <p>{{ [city, state].filter(Boolean).join('، ') }}</p>
          </div>
          
          <div v-if="postalCode" class="flex items-center justify-center space-x-2 space-x-reverse">
            <span class="text-lg">📮</span>
            <p>کد پستی: {{ postalCode }}</p>
          </div>
        </div>
      </div>
      
      <!-- پیام خالی -->
      <div v-else class="text-white text-opacity-60 text-sm">
        <div class="bg-white bg-opacity-10 rounded-lg p-6">
          <div class="text-4xl mb-3">🗺️</div>
          <p>مختصات نقشه تعریف نشده است</p>
          <p class="text-xs mt-2">لطفاً عرض و طول جغرافیایی را تنظیم کنید</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'

interface Props {
  title?: string
  latitude?: number
  longitude?: number
  address?: string
  city?: string
  state?: string
  postalCode?: string
  showLocationInfo?: boolean
  paddingRight?: number
  paddingLeft?: number
  align?: 'left' | 'center' | 'right'
}

const props = withDefaults(defineProps<Props>(), {
  title: 'موقعیت ما',
  latitude: 0,
  longitude: 0,
  address: '',
  city: '',
  state: '',
  postalCode: '',
  showLocationInfo: true,
  paddingRight: 0,
  paddingLeft: 0,
  align: 'center'
})

// متغیرهای محلی
const copyButtonText = ref('کپی مختصات')
const hasCopied = ref(false)

// باز کردن نقشه گوگل
const openGoogleMaps = () => {
  if (props.latitude && props.longitude) {
    const mapUrl = `https://www.google.com/maps?q=${props.latitude},${props.longitude}`
    window.open(mapUrl, '_blank')
  }
}

// کپی کردن مختصات
const copyCoordinates = async () => {
  if (props.latitude && props.longitude) {
    const coordinates = `${props.latitude}, ${props.longitude}`
    
    try {
      await navigator.clipboard.writeText(coordinates)
      copyButtonText.value = 'کپی شد!'
      hasCopied.value = true
      
      // بازگرداندن متن دکمه بعد از 2 ثانیه
      setTimeout(() => {
        copyButtonText.value = 'کپی مختصات'
        hasCopied.value = false
      }, 2000)
      
    } catch (error) {
      // روش جایگزین برای مرورگرهای قدیمی
      const textArea = document.createElement('textarea')
      textArea.value = coordinates
      document.body.appendChild(textArea)
      textArea.select()
      document.execCommand('copy')
      document.body.removeChild(textArea)
      
      copyButtonText.value = 'کپی شد!'
      hasCopied.value = true
      
      setTimeout(() => {
        copyButtonText.value = 'کپی مختصات'
        hasCopied.value = false
      }, 2000)
    }
  }
}

// استایل کامپوننت بر اساس چینش
const widgetStyle = computed(() => ({
  paddingRight: `${props.paddingRight}px`,
  paddingLeft: `${props.paddingLeft}px`,
  display: 'flex',
  alignItems: 'center',
  justifyContent: getJustifyContent(props.align),
  width: '100%',
  height: '100%'
}))

// تابع تعیین justify-content بر اساس چینش
function getJustifyContent(align: string): string {
  switch (align) {
    case 'left':
      return 'flex-start'  // در RTL: چپ = flex-start
    case 'center':
      return 'center'
    case 'right':
      return 'flex-end'  // در RTL: راست = flex-end
    default:
      return 'center'
  }
}
</script>

<style scoped>
.footer-map-widget {
  transition: all 0.2s ease;
}

.footer-map-widget button:hover {
  transform: translateY(-1px);
}

.footer-map-widget .animate-pulse {
  animation: pulse 2s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}

@keyframes pulse {
  0%, 100% {
    opacity: 1;
  }
  50% {
    opacity: .5;
  }
}
</style>
