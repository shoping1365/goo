<template>
  <header class="bg-white shadow-sm sticky top-0 z-40" dir="rtl">
    <div class="container mx-auto px-4 py-3 flex items-center justify-between gapx-4 py-4">
      <!-- Logo / Brand -->
      <NuxtLink to="/" class="text-lg font-bold text-gray-800 hover:text-blue-600 transition-colors">
        فروشگاه
      </NuxtLink>

      <!-- اکشن‌ها -->
      <div class="flex items-center gapx-4 py-4">
        <CartBadge />
        <div v-if="isAuthenticated" class="relative">
          <NuxtLink to="/account" class="flex items-center space-x-2 space-x-reverse p-2 rounded-lg hover:bg-gray-100 transition-colors">
            <div class="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center">
              <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
              </svg>
            </div>
            <span class="text-sm font-medium text-gray-700">{{ user?.name || 'کاربر' }}</span>
          </NuxtLink>
        </div>
      </div>
    </div>
  </header>
</template>

<script setup lang="ts">
import CartBadge from '~/components/cart/CartBadge.vue';
import { useAuth } from '~/composables/useAuth';
const { user, isAuthenticated } = useAuth()
</script> 

