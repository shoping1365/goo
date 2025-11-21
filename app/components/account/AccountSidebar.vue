<template>
  <div class="account-sidebar bg-white rounded-lg shadow-sm border border-gray-200 p-6 text-right">
    <!-- بخش کیف پول و دیجی کلاب -->
    <div class="mb-8">
      <!-- کیف پول -->
      <div class="bg-gradient-to-r from-blue-50 to-indigo-50 rounded-lg p-6 mb-4 text-right">
        <div class="flex items-center justify-between mb-2">
          <span class="inline-flex items-baseline gap-1">
            <span class="text-[10px] sm:text-xs text-gray-600">تومان</span>
            <span class="text-sm sm:text-base font-medium text-gray-800">{{ walletBalance || '0' }}</span>
          </span>
          <h3 class="text-base font-semibold text-gray-900">کیف پول</h3>
        </div>
        <NuxtLink to="/account/wallet" class="text-blue-600 hover:text-blue-700 text-sm font-medium transition-colors inline-flex items-center gap-1 flex-row-reverse">
          <span>افزایش موجودی</span>
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
          </svg>
        </NuxtLink>
      </div>

      <!-- دیجی کلاب -->
      <div class="bg-gradient-to-r from-purple-50 to-pink-50 rounded-lg p-6 text-right">
        <div class="flex items-baseline justify-between mb-2">
          <span class="text-sm sm:text-base font-medium text-gray-800">{{ digiClubPoints || '177' }}</span>
          <span class="text-[10px] sm:text-xs text-gray-600">امتیاز</span>
        </div>
      </div>
    </div>

    <!-- منوی اصلی -->
    <nav class="space-y-2">
      <template v-for="item in menuItems" :key="item.id">
        <!-- لینک عادی -->
        <NuxtLink 
          v-if="item.path && !item.action"
          :to="item.path"
          class="flex items-center space-x-3 space-x-reverse p-3 rounded-lg transition-colors relative"
          :class="[
            $route.path === item.path 
              ? 'bg-red-50 text-red-700 border-r-4 border-red-500'
              : 'text-gray-700 hover:bg-gray-50'
          ]"
        >
          <!-- آیکون -->
          <div class="flex-shrink-0 relative">
            <component :is="getIcon(item.icon)" class="w-5 h-5" />
            <!-- نشانگر پیام‌های جدید -->
            <div 
              v-if="item.badge" 
              class="absolute -top-1 -right-1 w-4 h-4 bg-red-500 text-white text-xs rounded-full flex items-center justify-center"
            >
              {{ item.badge }}
            </div>
          </div>
          
          <!-- عنوان -->
          <span class="font-medium text-right flex-1">{{ item.title }}</span>
        </NuxtLink>

        <!-- دکمه logout -->
        <button 
          v-else-if="item.action === 'logout'"
          class="flex items-center space-x-3 space-x-reverse p-3 rounded-lg transition-colors relative w-full text-right text-gray-700 hover:bg-gray-50"
          @click="handleLogout"
        >
          <!-- آیکون -->
          <div class="flex-shrink-0 relative">
            <component :is="getIcon(item.icon)" class="w-5 h-5" />
          </div>
          
          <!-- عنوان -->
          <span class="font-medium text-right flex-1">{{ item.title }}</span>
        </button>
      </template>
    </nav>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuth } from '~/composables/useAuth'

const router = useRouter()
const { logout } = useAuth()

// تعریف آیتم‌های منو
const menuItems = ref([
  {
    id: 'dashboard',
    title: 'خلاصه فعالیت ها',
    path: '/account',
    icon: 'HomeIcon',
    badge: null
  },
  {
    id: 'orders',
    title: 'سفارش ها',
    path: '/account/orders',
    icon: 'ShoppingBagIcon',
    badge: null
  },
  {
    id: 'wishlist',
    title: 'لیست های من',
    path: '/account/wishlist',
    icon: 'HeartIcon',
    badge: null
  },
  {
    id: 'reviews',
    title: 'دیدگاه ها و پرسشها',
    path: '/account/reviews',
    icon: 'ChatIcon',
    badge: null
  },
  {
    id: 'addresses',
    title: 'آدرس ها',
    path: '/account/addresses',
    icon: 'LocationIcon',
    badge: null
  },
  {
    id: 'gift-cards',
    title: 'کارت های هدیه',
    path: '/account/gift-cards',
    icon: 'GiftIcon',
    badge: null
  },
  {
    id: 'banking',
    title: 'اطلاعات بانکی',
    path: '/account/banking',
    icon: 'BankIcon',
    badge: null
  },
  {
    id: 'messages',
    title: 'پیام ها',
    path: '/account/messages',
    icon: 'BellIcon',
    badge: null
  },
  {
    id: 'recent-views',
    title: 'بازدیدهای اخیر',
    path: '/account/recent-views',
    icon: 'ClockIcon',
    badge: null
  },
  {
    id: 'profile',
    title: 'اطلاعات حساب کاربری',
    path: '/account/profile',
    icon: 'UserIcon',
    badge: null
  },
  {
    id: 'logout',
    title: 'خروج از حساب کاربری',
    path: null,
    icon: 'LogoutIcon',
    badge: null,
    action: 'logout'
  }
])

// داده‌های کیف پول و دیجی کلاب
const walletBalance = ref('0')
const digiClubPoints = ref('177')

// تابع logout
const handleLogout = async () => {
  try {
    // Auth disabled
await logout()
  } catch (error) {
    console.error('خطا در logout:', error)
    // در صورت خطا، به صفحه اصلی برو
    await router.push('/')
  }
}

// آیکون‌های SVG
const HomeIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"></path>
  </svg>`
}

const ShoppingBagIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z"></path>
  </svg>`
}

const HeartIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"></path>
  </svg>`
}

const ChatIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
  </svg>`
}

const LocationIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"></path>
    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"></path>
  </svg>`
}

const GiftIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v13m0-13V6a2 2 0 112 2h-2zm0 0V5.5A2.5 2.5 0 109.5 8H12zm-7 4h14M5 12a2 2 0 110-4h14a2 2 0 110 4M5 12v7a2 2 0 002 2h10a2 2 0 002-2v-7"></path>
  </svg>`
}

const BellIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-5 5v-5zM4.83 2.83A4 4 0 0011 7.16V11a4 4 0 001.17 2.83l.83.83a4 4 0 001.17 2.83V19a4 4 0 01-4 4H5a4 4 0 01-4-4v-2.17a4 4 0 001.17-2.83l.83-.83A4 4 0 0011 11V7.16a4 4 0 00-1.17-2.83z"></path>
  </svg>`
}

const ClockIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
  </svg>`
}

const UserIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
  </svg>`
}

const LogoutIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"></path>
  </svg>`
}

const BankIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"></path>
  </svg>`
}

// نگاشت آیکون‌ها برای استفاده در template
interface IconComponent {
  template: string
}

const iconMap: Record<string, IconComponent> = {
  HomeIcon,
  ShoppingBagIcon,
  HeartIcon,
  ChatIcon,
  LocationIcon,
  GiftIcon,
  BellIcon,
  ClockIcon,
  UserIcon,
  LogoutIcon,
  BankIcon
}

// تابع برای دریافت آیکون
const getIcon = (iconName: string) => {
  return iconMap[iconName] || HomeIcon
}
</script>

<style scoped>
.account-sidebar {
  min-width: 280px;
  max-width: 320px;
}

/* استایل برای آیتم فعال - انتقال نوار قرمز به سمت راست */
.account-sidebar nav a.router-link-active {
  background-color: #fef2f2;
  color: #dc2626;
  border-right: 4px solid #dc2626;
}
</style>
