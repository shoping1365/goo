<template>
  <div class="classic-header">
    <div class="top-bar">
      <div class="icons-section">
        <!-- نمایش لینک حساب کاربری فقط برای کاربران لاگین شده -->
        <NuxtLink v-if="header?.show_user_menu && isAuthenticated" to="/account" class="icon user-icon">👤</NuxtLink>
        <!-- نمایش لینک ورود برای کاربران غیر لاگین -->
        <NuxtLink v-else-if="header?.show_user_menu && !isAuthenticated" to="/auth/login" class="icon user-icon">👤</NuxtLink>
        <div v-if="header?.show_notifications" class="icon moon-icon">🌙</div>
      </div>
      <div class="logo-section">
        <img v-if="header?.logo_url" :src="header.logo_url" :alt="header.logo_alt || 'لوگو'" class="header-logo" />
        <span v-else class="logo-placeholder">{{ header?.logo_text || 'لوگو' }}</span>
      </div>
    </div>
    <div class="search-section">
      <NuxtLink v-if="header?.show_cart" to="/Cart" class="cart-icon">🛒</NuxtLink>
      <div class="search-box">
        <MobileHeaderSearchBox
          :placeholder="header?.search_placeholder || 'جستجو...'"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import MobileHeaderSearchBox from '~/components/widgets/mobile/MobileHeaderSearchBox.vue'
import { useAuthState } from '~/composables/useAuthState'

defineProps({
  header: {
    type: Object,
    required: true
  }
})

// دریافت وضعیت احراز هویت
const { isAuthenticated } = useAuthState()
</script>

<style scoped>
.classic-header {
  display: flex;
  flex-direction: column;
  width: 100%;
  gap: 12px;
  padding: 12px 8px 0 8px;
  min-height: 80px;
}

.top-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.logo-section {
  flex-shrink: 0;
}

.header-logo {
  height: 28px;
  width: auto;
  object-fit: contain;
}

.logo-placeholder {
  display: inline-block;
  padding: 4px 8px;
  background: #f3f4f6;
  border-radius: 4px;
  font-size: 11px;
  color: #6b7280;
  font-weight: 500;
}

.icons-section {
  display: flex;
  gap: 8px;
}

.icon {
  width: 32px;
  height: 32px;
  border-radius: 6px;
  background: #e5e7eb;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 14px;
  cursor: pointer;
  transition: background-color 0.2s;
}

.icon:hover {
  background: #d1d5db;
}

.search-section {
  display: flex;
  align-items: center;
  gap: 8px;
}

.search-box {
  flex: 1;
  position: relative;
}

.cart-icon {
  width: 36px;
  height: 36px;
  border-radius: 18px;
  background: #f3f4f6;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 16px;
  cursor: pointer;
  transition: background-color 0.2s;
  position: relative;
  color: #000000;
}

.cart-icon:hover {
  background: #e5e7eb;
}
</style>
