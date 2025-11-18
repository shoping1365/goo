<template>
  <nav class="main-menu-vertical" :class="themeClass">
    <ul class="menu-list">
      <li 
        v-for="item in filteredMenuItems" 
        :key="item.id"
        class="menu-item"
        :class="{ 
          'has-children': item.children && item.children.length > 0,
          'expanded': expandedItems.includes(item.id)
        }"
      >
        <div class="menu-link-wrapper">
          <NuxtLink 
            v-if="item.path" 
            :to="item.path" 
            class="menu-link"
            :class="{ active: isActive(item.path) }"
          >
            <i v-if="settings.showIcons && item.icon" :class="item.icon" class="menu-icon"></i>
            <span class="menu-title">{{ item.title }}</span>
            <span v-if="settings.showBadges && item.badge" class="menu-badge">{{ item.badge }}</span>
          </NuxtLink>
          
          <div v-else class="menu-link menu-parent">
            <i v-if="settings.showIcons && item.icon" :class="item.icon" class="menu-icon"></i>
            <span class="menu-title">{{ item.title }}</span>
            <span v-if="settings.showBadges && item.badge" class="menu-badge">{{ item.badge }}</span>
          </div>
          
          <button 
            v-if="item.children && item.children.length > 0"
            @click="toggleExpand(item.id)"
            class="expand-button"
            :class="{ rotated: expandedItems.includes(item.id) }"
          >
            <i class="fas fa-chevron-left expand-arrow"></i>
          </button>
        </div>
        
        <!-- Submenu -->
        <ul 
          v-if="item.children && item.children.length > 0"
          class="submenu"
          :class="{ expanded: expandedItems.includes(item.id) }"
        >
          <li v-for="child in item.children" :key="child.id" class="submenu-item">
            <NuxtLink 
              v-if="child.path"
              :to="child.path" 
              class="submenu-link"
              :class="{ active: isActive(child.path) }"
            >
              <i v-if="settings.showIcons && child.icon" :class="child.icon" class="menu-icon"></i>
              <span class="menu-title">{{ child.title }}</span>
              <span v-if="settings.showBadges && child.badge" class="menu-badge">{{ child.badge }}</span>
            </NuxtLink>
            <div v-else class="submenu-link">
              <i v-if="settings.showIcons && child.icon" :class="child.icon" class="menu-icon"></i>
              <span class="menu-title">{{ child.title }}</span>
              <span v-if="settings.showBadges && child.badge" class="menu-badge"></span>
            </div>
          </li>
        </ul>
      </li>
    </ul>
  </nav>
</template>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'
import { useRoute } from 'vue-router'
import { useAuthState } from '~/composables/useAuthState'


interface MenuChild {
  id: string
  title: string
  path?: string
  icon?: string
  badge?: string
  enabled: boolean
  permission?: string
}

interface MenuItem {
  id: string
  title: string
  path?: string
  icon?: string
  badge?: string
  enabled: boolean
  permission?: string
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
const expandedItems = ref<string[]>([])
const authState = useAuthState()

// Helper function برای بررسی مجوز
const checkPermission = (permission: string): boolean => {
  return authState.hasPermission(permission)
}

const themeClass = computed(() => {
  return `theme-${props.settings.theme}`
})

function isActive(path?: string): boolean {
  if (!path) return false
  return route.path === path || route.path.startsWith(path + '/')
}

function toggleExpand(itemId: string) {
  const index = expandedItems.value.indexOf(itemId)
  if (index > -1) {
    expandedItems.value.splice(index, 1)
  } else {
    expandedItems.value.push(itemId)
  }
}

// Auto expand active parent items
function checkActiveParents() {
  props.menuItems.forEach(item => {
    if (item.children && item.children.some(child => isActive(child.path))) {
      if (!expandedItems.value.includes(item.id)) {
        expandedItems.value.push(item.id)
      }
    }
  })
}

// Check on mount and route change
onMounted(() => {
  checkActiveParents()
})

watch(() => route.path, () => {
  checkActiveParents()
})

// فیلتر منوها بر اساس مجوز
const filteredMenuItems = computed(() => {
  return props.menuItems
    .filter(item => !item.permission || checkPermission(item.permission || ''))
    .map(item => ({
      ...item,
      children: item.children ? item.children.filter(child => !child.permission || checkPermission(child.permission || '')) : []
    }))
})
</script>

<style scoped>
.main-menu-vertical {
  width: 100%;
  background: linear-gradient(180deg, #2c3e50 0%, #34495e 100%);
  border-radius: 16px;
  padding: 16px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.15);
  min-height: 400px;
}

.menu-list {
  list-style: none;
  margin: 0;
  padding: 0;
}

.menu-list > li {
  margin-bottom: 8px;
}

.menu-list > li:last-child {
  margin-bottom: 0;
}

.menu-item {
  width: 100%;
}

.menu-link-wrapper {
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 100%;
}

.menu-link {
  display: flex;
  align-items: center;
  padding: 16px 20px;
  color: #ecf0f1;
  text-decoration: none;
  flex: 1;
  border-radius: 12px;
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
  background: linear-gradient(135deg, rgba(52, 152, 219, 0.2) 0%, rgba(155, 89, 182, 0.2) 100%);
  opacity: 0;
  transition: opacity 0.3s ease;
}

.menu-link:hover::before {
  opacity: 1;
}

.menu-link:hover {
  transform: translateX(-4px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
}

.menu-link.active {
  background: linear-gradient(135deg, #3498db 0%, #9b59b6 100%);
  color: white;
  font-weight: 600;
  box-shadow: 0 4px 12px rgba(52, 152, 219, 0.3);
}

.menu-parent {
  cursor: default;
  flex: 1;
}

.menu-icon {
  margin-right: 12px;
  font-size: 18px;
  opacity: 0.9;
  flex-shrink: 0;
  width: 20px;
  text-align: center;
}

.menu-title {
  font-size: 14px;
  font-weight: 500;
  flex: 1;
  position: relative;
  z-index: 1;
  text-align: right;
}

.menu-badge {
  background: linear-gradient(45deg, #e74c3c, #c0392b);
  color: white;
  font-size: 11px;
  padding: 4px 8px;
  border-radius: 12px;
  margin-left: 8px;
  font-weight: 600;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  flex-shrink: 0;
  position: relative;
  z-index: 1;
}

.expand-button {
  padding: 8px;
  color: #bdc3c7;
  transition: all 0.3s ease;
  border-radius: 8px;
  flex-shrink: 0;
  background: transparent;
  border: none;
  cursor: pointer;
  margin-right: 8px;
}

.expand-button:hover {
  color: #ecf0f1;
  background: rgba(255, 255, 255, 0.1);
}

.expand-arrow {
  font-size: 14px;
  transition: transform 0.3s ease;
}

.expand-button.rotated .expand-arrow {
  transform: rotate(90deg);
}

.submenu {
  list-style: none;
  margin: 0;
  padding: 0;
  overflow: hidden;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  max-height: 0;
  opacity: 0;
  margin-top: 8px;
}

.submenu.expanded {
  max-height: 500px;
  opacity: 1;
  padding-bottom: 8px;
}

.submenu-item {
  width: 100%;
}

.submenu-link {
  display: flex;
  align-items: center;
  padding: 12px 16px;
  margin-left: 20px;
  color: #bdc3c7;
  text-decoration: none;
  font-size: 13px;
  border-radius: 8px;
  transition: all 0.2s ease;
  position: relative;
  overflow: hidden;
}

.submenu-link::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(52, 152, 219, 0.1);
  opacity: 0;
  transition: opacity 0.2s ease;
}

.submenu-link:hover::before {
  opacity: 1;
}

.submenu-link:hover {
  color: #ecf0f1;
  transform: translateX(4px);
}

.submenu-link.active {
  background: linear-gradient(135deg, #3498db 0%, #9b59b6 100%);
  color: white;
  font-weight: 600;
}

/* Theme Variations */
.theme-modern {
  background: linear-gradient(180deg, #2c3e50 0%, #34495e 100%);
}

.theme-minimal {
  background: transparent;
  box-shadow: none;
  padding: 0;
}

.theme-minimal .menu-link {
  color: #2c3e50;
  background: transparent;
  border-left: 3px solid transparent;
  border-radius: 0;
  padding: 16px 20px;
}

.theme-minimal .menu-link:hover,
.theme-minimal .menu-link.active {
  color: #3498db;
  background: transparent;
  border-left-color: #3498db;
  transform: none;
  box-shadow: none;
}

.theme-minimal .submenu-link {
  color: #7f8c8d;
  background: transparent;
  border-left: 2px solid transparent;
  border-radius: 0;
}

.theme-minimal .submenu-link:hover,
.theme-minimal .submenu-link.active {
  color: #3498db;
  background: transparent;
  border-left-color: #3498db;
  transform: none;
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
  margin-bottom: 4px;
}

.theme-classic .menu-link:hover {
  background: #e9ecef;
  transform: none;
}

.theme-classic .submenu-link {
  color: #6c757d;
  background: #f8f9fa;
  border: 1px solid #e9ecef;
  margin-bottom: 2px;
}

/* Responsive adjustments */
@media (max-width: 768px) {
  .menu-icon {
    margin-right: 8px;
  }
  
  .menu-title {
    font-size: 13px;
  }
  
  .menu-badge {
    font-size: 10px;
    padding: 3px 6px;
  }
  
  .menu-link {
    padding: 14px 16px;
  }
  
  .submenu-link {
    padding: 10px 14px;
    font-size: 12px;
  }
}
</style> 
