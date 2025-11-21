<template>
  <div class="minimal-header">
    <div class="search-section">
      <div class="search-box">
        <MobileHeaderSearchBox
          :placeholder="header?.search_placeholder || 'جستجو'"
        />
        <div class="logo-section">
          <img v-if="header?.logo_url" :src="header.logo_url" :alt="header.logo_alt || 'لوگو'" class="header-logo" />
          <span v-else class="logo-placeholder">{{ header?.logo_text || 'نام هدر' }}</span>
        </div>
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
.minimal-header {
  display: flex;
  align-items: center;
  width: 100%;
  padding: 12px 8px 0 8px;
  min-height: 80px;
}

.search-section {
  flex: 1;
}

.search-box {
  position: relative;
  display: flex;
  align-items: center;
  gap: 12px;
  flex-direction: row-reverse;
}

.logo-section {
  flex-shrink: 0;
}

.header-logo {
  height: 24px;
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

</style>
