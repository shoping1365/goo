<template>
  <nav class="main-menu-horizontal" :class="themeClass">
    <ul class="menu-list">
      <li 
        v-for="item in menuItems" 
        :key="item.id"
        class="menu-item"
        :class="{ 'has-children': item.children && item.children.length > 0 }"
        @mouseenter="showDropdown(item.id)"
        @mouseleave="hideDropdown(item.id)"
      >
        <NuxtLink 
          v-if="item.path" 
          :to="item.path" 
          class="menu-link"
          :class="{ active: isActive(item.path) }"
        >
          <i v-if="settings.showIcons && item.icon" :class="item.icon" class="menu-icon"></i>
          <span class="menu-title">{{ item.title }}</span>
          <span v-if="settings.showBadges && item.badge" class="menu-badge">{{ item.badge }}</span>
          <i v-if="item.children && item.children.length > 0" class="fas fa-chevron-down dropdown-arrow"></i>
        </NuxtLink>
        
        <span v-else class="menu-link menu-parent">
          <i v-if="settings.showIcons && item.icon" :class="item.icon" class="menu-icon"></i>
          <span class="menu-title">{{ item.title }}</span>
          <span v-if="settings.showBadges && item.badge" class="menu-badge">{{ item.badge }}</span>
          <i v-if="item.children && item.children.length > 0" class="fas fa-chevron-down dropdown-arrow"></i>
        </span>
        
        <!-- Dropdown Menu -->
        <ul 
          v-if="item.children && item.children.length > 0"
          class="dropdown-menu"
          :class="{ show: activeDropdown === item.id }"
        >
          <li v-for="child in item.children" :key="child.id" class="dropdown-item">
            <NuxtLink 
              v-if="child.path"
              :to="child.path" 
              class="dropdown-link"
              :class="{ active: isActive(child.path) }"
            >
              <i v-if="settings.showIcons && child.icon" :class="child.icon" class="menu-icon"></i>
              <span class="menu-title">{{ child.title }}</span>
              <span v-if="settings.showBadges && child.badge" class="menu-badge">{{ child.badge }}</span>
            </NuxtLink>
            <span v-else class="dropdown-link">
              <i v-if="settings.showIcons && child.icon" :class="child.icon" class="menu-icon"></i>
              <span class="menu-title">{{ child.title }}</span>
              <span v-if="settings.showBadges && child.badge" class="menu-badge">{{ child.badge }}</span>
            </span>
          </li>
        </ul>
      </li>
    </ul>
  </nav>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import { useRoute } from 'vue-router'

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

const props = defineProps<{
  menuItems: MenuItem[]
  settings: MenuSettings
}>()

const route = useRoute()
const activeDropdown = ref<string | null>(null)

const themeClass = computed(() => {
  return `theme-${props.settings.theme}`
})

function isActive(path?: string): boolean {
  if (!path) return false
  return route.path === path
}

function showDropdown(itemId: string) {
  activeDropdown.value = itemId
}

function hideDropdown(itemId: string) {
  setTimeout(() => {
    if (activeDropdown.value === itemId) {
      activeDropdown.value = null
    }
  }, 200)
}
</script>

<style scoped>
.main-menu-horizontal {
  width: 100%;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 12px;
  padding: 8px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
}

.menu-list {
  display: flex;
  align-items: center;
  list-style: none;
  margin: 0;
  padding: 0;
  gap: 4px;
  direction: rtl;
}

.menu-item {
  position: relative;
}

.menu-link {
  display: flex;
  align-items: center;
  padding: 12px 20px;
  color: white;
  text-decoration: none;
  white-space: nowrap;
  border-radius: 8px;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  font-weight: 500;
  position: relative;
  overflow: hidden;
}

.menu-link::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(255, 255, 255, 0.1);
  opacity: 0;
  transition: opacity 0.3s ease;
}

.menu-link:hover::before {
  opacity: 1;
}

.menu-link:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.menu-link.active {
  background: rgba(255, 255, 255, 0.2);
  backdrop-filter: blur(10px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
}

.menu-parent {
  cursor: default;
}

.menu-icon {
  margin-right: 8px;
  font-size: 16px;
  opacity: 0.9;
}

.menu-title {
  font-size: 14px;
  font-weight: 500;
  position: relative;
  z-index: 1;
}

.menu-badge {
  background: linear-gradient(45deg, #ff6b6b, #ee5a24);
  color: white;
  font-size: 11px;
  padding: 4px 8px;
  border-radius: 12px;
  margin-left: 8px;
  font-weight: 600;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  position: relative;
  z-index: 1;
}

.dropdown-arrow {
  margin-left: 8px;
  font-size: 12px;
  transition: transform 0.3s ease;
  opacity: 0.8;
}

.menu-item:hover .dropdown-arrow {
  transform: rotate(180deg);
}

.dropdown-menu {
  position: absolute;
  top: calc(100% + 8px);
  left: 0;
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(20px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 12px;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.15);
  min-width: 200px;
  opacity: 0;
  visibility: hidden;
  transform: translateY(-10px) scale(0.95);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  z-index: 1000;
  list-style: none;
  margin: 0;
  padding: 8px;
}

.dropdown-menu.show {
  opacity: 1;
  visibility: visible;
  transform: translateY(0) scale(1);
}

.dropdown-item {
  list-style: none;
}

.dropdown-link {
  display: flex;
  align-items: center;
  padding: 12px 16px;
  color: #374151;
  text-decoration: none;
  font-size: 14px;
  border-radius: 8px;
  transition: all 0.2s ease;
  margin-bottom: 2px;
}

.dropdown-link:hover {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  transform: translateX(4px);
}

.dropdown-link.active {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  font-weight: 600;
}

.dropdown-link:last-child {
  margin-bottom: 0;
}

/* Theme Variations */
.theme-modern {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.theme-minimal {
  background: transparent;
  box-shadow: none;
  padding: 0;
}

.theme-minimal .menu-link {
  color: #374151;
  background: transparent;
  border-bottom: 2px solid transparent;
  border-radius: 0;
  padding: 16px 20px;
}

.theme-minimal .menu-link:hover,
.theme-minimal .menu-link.active {
  color: #667eea;
  background: transparent;
  border-bottom-color: #667eea;
  transform: none;
  box-shadow: none;
}

.theme-classic {
  background: #f8f9fa;
  border: 1px solid #e9ecef;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.theme-classic .menu-link {
  color: #495057;
  background: white;
  border: 1px solid #e9ecef;
  margin: 0 2px;
}

.theme-classic .menu-link:hover {
  background: #e9ecef;
  transform: none;
}

/* Responsive adjustments */
@media (max-width: 768px) {
  .menu-list {
    flex-direction: column;
    gap: 8px;
  }
  
  .menu-item {
    width: 100%;
  }
  
  .menu-link {
    justify-content: space-between;
    padding: 16px 20px;
  }
  
  .dropdown-menu {
    position: relative;
    top: 0;
    right: 0;
    border: none;
    box-shadow: none;
    background: rgba(102, 126, 234, 0.1);
    margin-left: 20px;
    margin-top: 8px;
  }
  
  .dropdown-menu.show {
    display: block;
  }
  
  .dropdown-link:hover {
    transform: none;
  }
}
</style> 
