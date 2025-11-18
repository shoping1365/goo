<template>
  <div class="main-site-menu" :class="containerClass">
    <!-- Horizontal Menu -->
    <div v-if="layout === 'horizontal' || layout === 'both'" class="horizontal-menu-container">
      <MainMenuHorizontal 
        :menu-items="menuItems"
        :settings="horizontalSettings"
      />
    </div>

    <!-- Vertical Menu -->
    <div v-if="layout === 'vertical' || layout === 'both'" class="vertical-menu-container">
      <MainMenuVertical 
        :menu-items="menuItems"
        :settings="verticalSettings"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import MainMenuHorizontal from './MainMenuHorizontal.vue'
import MainMenuVertical from './MainMenuVertical.vue'

interface MenuChild {
  id: string
  title: string
  path?: string
  icon?: string
  badge?: string
  enabled: boolean
}

interface MenuItem {
  id: string
  title: string
  path?: string
  icon?: string
  badge?: string
  enabled: boolean
  children?: MenuChild[]
}

interface MenuSettings {
  theme: string
  showIcons: boolean
  showBadges: boolean
}

interface Props {
  menuItems: MenuItem[]
  layout?: 'horizontal' | 'vertical' | 'both'
  theme?: string
  showIcons?: boolean
  showBadges?: boolean
  responsive?: boolean
  rtl?: boolean
}

const props = defineProps<Props>()

const horizontalSettings = computed<MenuSettings>(() => ({
  theme: props.theme,
  showIcons: props.showIcons,
  showBadges: props.showBadges
}))

const verticalSettings = computed<MenuSettings>(() => ({
  theme: props.theme,
  showIcons: props.showIcons,
  showBadges: props.showBadges
}))

const containerClass = computed(() => ({
  'layout-horizontal': props.layout === 'horizontal',
  'layout-vertical': props.layout === 'vertical',
  'layout-both': props.layout === 'both',
  'responsive': props.responsive,
  'rtl': props.rtl
}))
</script>

<style scoped>
.main-site-menu {
  width: 100%;
}

.layout-horizontal {
  /* Horizontal layout styles */
}

.layout-vertical {
  /* Vertical layout styles */
}

.layout-both {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.layout-both .horizontal-menu-container {
  order: 1;
}

.layout-both .vertical-menu-container {
  order: 2;
  max-width: 300px;
  margin: 0 auto;
}

.horizontal-menu-container {
  width: 100%;
}

.vertical-menu-container {
  width: 100%;
}

/* Responsive Design */
@media (max-width: 768px) {
  .layout-both {
    flex-direction: column;
    gap: 16px;
  }
  
  .layout-both .vertical-menu-container {
    max-width: 100%;
  }
}

/* RTL Support */
.rtl {
  direction: rtl;
}

.rtl .layout-both .horizontal-menu-container {
  order: 1;
}

.rtl .layout-both .vertical-menu-container {
  order: 2;
}

/* Theme-specific container styles */
.main-site-menu[class*="theme-modern"] {
  /* Modern theme container styles */
}

.main-site-menu[class*="theme-minimal"] {
  /* Minimal theme container styles */
}

.main-site-menu[class*="theme-classic"] {
  /* Classic theme container styles */
}
</style> 
