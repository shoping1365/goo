<template>
  <div class="footer-address-widget" :style="widgetStyle">
    <div class="text-center">
      <!-- عنوان بخش -->
      <h4 v-if="title" class="text-white font-semibold mb-4 text-lg">{{ title }}</h4>
      
      <!-- آدرس کامل -->
      <div v-if="address" class="space-y-3">
        <!-- آدرس اصلی -->
        <div class="flex items-center justify-center space-x-2 space-x-reverse text-white">
          <span class="text-lg">📍</span>
          <p class="text-sm leading-relaxed">{{ address }}</p>
        </div>
        
        <!-- شهر و استان -->
        <div v-if="city || state" class="flex items-center justify-center space-x-2 space-x-reverse text-white">
          <span class="text-lg">🏙️</span>
          <p class="text-sm">
            {{ [city, state].filter(Boolean).join('، ') }}
          </p>
        </div>
        
        <!-- کد پستی -->
        <div v-if="postalCode" class="flex items-center justify-center space-x-2 space-x-reverse text-white">
          <span class="text-lg">📮</span>
          <p class="text-sm">کد پستی: {{ postalCode }}</p>
        </div>
        
        <!-- کشور -->
        <div v-if="country" class="flex items-center justify-center space-x-2 space-x-reverse text-white">
          <span class="text-lg">🌍</span>
          <p class="text-sm">{{ country }}</p>
        </div>
      </div>
      
      <!-- نقشه -->
      <div v-if="showMap && (latitude && longitude)" class="mt-4">
        <div class="bg-white bg-opacity-10 rounded-lg p-3">
          <div class="w-full h-32 bg-gray-300 rounded flex items-center justify-center">
            <span class="text-gray-600 text-sm">نقشه: {{ latitude }}, {{ longitude }}</span>
          </div>
          <button
            class="mt-2 w-full px-3 py-2 bg-blue-600 hover:bg-blue-700 text-white text-sm rounded transition-colors"
            @click="openMap"
          >
            مشاهده در نقشه
          </button>
        </div>
      </div>
      
      <!-- پیام خالی -->
      <div v-if="!address" class="text-white text-opacity-60 text-sm">
        آدرس تعریف نشده است
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface Props {
  title?: string
  address?: string
  city?: string
  state?: string
  postalCode?: string
  country?: string
  latitude?: number
  longitude?: number
  showMap?: boolean
  paddingRight?: number
  paddingLeft?: number
  align?: 'left' | 'center' | 'right'
}

const props = withDefaults(defineProps<Props>(), {
  title: 'آدرس ما',
  address: '',
  city: '',
  state: '',
  postalCode: '',
  country: 'ایران',
  latitude: 0,
  longitude: 0,
  showMap: false,
  paddingRight: 0,
  paddingLeft: 0,
  align: 'center'
})

// باز کردن نقشه
const openMap = () => {
  if (props.latitude && props.longitude) {
    const mapUrl = `https://www.google.com/maps?q=${props.latitude},${props.longitude}`
    window.open(mapUrl, '_blank')
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
.footer-address-widget {
  transition: all 0.2s ease;
}

.footer-address-widget button:hover {
  transform: translateY(-1px);
}
</style>
