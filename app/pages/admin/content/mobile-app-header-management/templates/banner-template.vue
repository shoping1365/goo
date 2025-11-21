<template>
  <div class="banner-header">
    <div class="search-section">
      <NuxtLink v-if="header?.show_notifications" to="/notifications" class="notification-icon">🔔</NuxtLink>
      <div class="search-box">
        <div class="logo-right">
          <img v-if="header?.logo_url" :src="header.logo_url" :alt="header.logo_alt || 'لوگو'" class="header-logo" />
          <span v-else class="logo-text">{{ header?.logo_text || 'لوگو' }}</span>
        </div>
        <MobileHeaderSearchBox
          :placeholder="header?.search_placeholder || 'جستجو...'"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import MobileHeaderSearchBox from '~/components/widgets/mobile/MobileHeaderSearchBox.vue';

defineProps({
  header: {
    type: Object,
    required: true
  }
})
</script>

<style scoped>
.banner-header {
  display: flex;
  flex-direction: column;
  width: 100%;
  gap: 8px;
  padding: 16px 8px 0 8px;
  min-height: 80px;
  justify-content: center;
}

.search-section {
  display: flex;
  align-items: center;
  gap: 8px;
}


.search-box {
  flex: 1;
  position: relative;
  display: flex;
  align-items: center;
  gap: 8px;
}

:deep(.mobile-search-box input) {
  background: transparent;
  border: none;
  box-shadow: none;
  text-align: right;
  direction: rtl;
}

.logo-right {
  display: flex;
  align-items: center;
  gap: 6px;
  flex-shrink: 0;
}

.header-logo {
  height: 24px;
  width: auto;
  object-fit: contain;
}

.logo-text {
  font-size: 11px;
  color: #374151;
  font-weight: 600;
}

.notification-icon {
  width: 32px;
  height: 32px;
  border-radius: 6px;
  background: #f3f4f6;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 14px;
  color: #6b7280;
  cursor: pointer;
  transition: background-color 0.2s;
}

.notification-icon:hover {
  background: #e5e7eb;
}
</style>
