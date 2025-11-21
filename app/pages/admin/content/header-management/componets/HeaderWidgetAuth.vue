<template>
  <div class="auth-widget" :style="widgetStyle">
    <!-- اگر لاگین نیست -->
    <NuxtLink
      v-if="!isAuthenticated"
      to="/auth/login"
      class="flex items-center justify-center p-2 rounded-lg transition-colors w-full h-full auth-link"
      :class="{
        'justify-start': props.align === 'left',
        'justify-center': props.align === 'center', 
        'justify-end': props.align === 'right'
      }"
    >
      <span class="text-xl">👤</span>
    </NuxtLink>

    <!-- اگر لاگین است -->
    <div v-else class="relative w-full">
      <button 
        class="flex items-center justify-center p-2 rounded-lg transition-colors w-full h-full auth-button"
        :class="{
          'justify-start': props.align === 'left',
          'justify-center': props.align === 'center', 
          'justify-end': props.align === 'right'
        }"
        @click="toggleMenu"
      >
        <span class="text-xl">👤</span>
      </button>
      
      <!-- منوی کشویی -->
      <div
        v-if="isMenuOpen"
        class="absolute top-full mt-2 bg-white text-gray-800 p-2 rounded-lg shadow-xl border border-gray-200 min-w-[200px] z-50"
        :class="{
          'left-0': props.align === 'left',
          'right-0': props.align === 'right',
          'left-1/2 transform -translate-x-1/2': props.align === 'center',
          'sm:min-w-[200px] min-w-[180px]': true
        }"
      >
        <!-- لینک‌های مشترک -->
        <NuxtLink 
          to="/account" 
          class="block px-4 py-2 text-sm hover:bg-gray-50 rounded transition-colors"
          @click="closeMenu"
        >
          👤 حساب کاربری
        </NuxtLink>
        <NuxtLink 
          to="/account/orders" 
          class="block px-4 py-2 text-sm hover:bg-gray-50 rounded transition-colors"
          @click="closeMenu"
        >
          📦 سفارشات
        </NuxtLink>
        <NuxtLink 
          to="/account/messages" 
          class="block px-4 py-2 text-sm hover:bg-gray-50 rounded transition-colors"
          @click="closeMenu"
        >
          💬 پرسش و پاسخ
        </NuxtLink>
        
        <hr class="my-2 border-gray-200">
        <button 
          class="w-full px-4 py-2 text-sm text-red-600 hover:bg-red-50 rounded text-right transition-colors" 
          @click="handleLogout"
        >
          🚪 خروج
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, onUnmounted, ref, watch } from 'vue'
import { useAuthState } from '~/composables/useAuthState'

interface Props {
  paddingRight?: number
  paddingLeft?: number
  align?: 'left' | 'center' | 'right'
}

const props = withDefaults(defineProps<Props>(), {
  paddingRight: 0,
  paddingLeft: 0,
  align: 'center'
})

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

const { isAuthenticated, logout, fetchUser } = useAuthState()

// مدیریت منوی کشویی
const isMenuOpen = ref(false)

const toggleMenu = () => {
  isMenuOpen.value = !isMenuOpen.value
}

// بستن منو وقتی روی خارج از آن کلیک می‌شود
const closeMenu = () => {
  isMenuOpen.value = false
}

// تابع خروج با بستن منو
const handleLogout = async () => {
  closeMenu()
  await logout()
}

const handleClickOutside = (event: Event) => {
  const target = event.target as HTMLElement
  if (!target.closest('.relative')) {
    closeMenu()
  }
}

onMounted(async () => {
  // بررسی وضعیت احراز هویت در زمان بارگذاری
  await fetchUser()
  
  document.addEventListener('click', handleClickOutside)
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})

// Watch برای تغییرات احراز هویت
watch(isAuthenticated, (newValue) => {
  if (!newValue) {
    // اگر کاربر خارج شد، منو را ببند
    closeMenu()
  }
})
</script>

<style scoped>
.auth-widget, .auth-widget * {
  font-family: 'Yekan', Tahoma, Arial, sans-serif ;
}

/* Responsive Design */
@media (max-width: 768px) {
  .auth-link,
  .auth-button {
    padding: 4px;
  }
}

@media (max-width: 480px) {
  .auth-link,
  .auth-button {
    padding: 2px;
  }
}
</style>

