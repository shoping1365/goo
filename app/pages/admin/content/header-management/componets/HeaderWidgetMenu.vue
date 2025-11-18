<template>
  <nav class="main-nav" @mouseleave="closeAllDropdowns">
    <!-- بارگذاری -->
    <div v-if="isLoading" class="nav-loading">در حال بارگذاری منو...</div>
    
    <!-- منوی اصلی -->
    <ul v-else-if="menuItems.length > 0" class="nav-menu">
      <li
        v-for="item in menuItems"
        :key="item.id"
        class="nav-item"
        :class="{ 'has-children': hasChildren(item), 'is-mega': item.isMegaMenu }"
        @mouseenter="openDropdown(item)"
      >
        <!-- آیتم اصلی -->
        <NuxtLink
          v-if="item.path"
          :to="item.path"
          :target="item.openInNewTab ? '_blank' : '_self'"
          class="nav-link"
          :class="{ 'is-active': isActive(item.path) }"
        >
          <i v-if="item.icon && item.iconType === 'icon'" :class="item.icon" class="nav-icon"></i>
          <img v-else-if="item.imageURL && item.iconType === 'image'" :src="item.imageURL" :alt="item.title" class="nav-image" />
          <span class="nav-text">{{ item.title }}</span>
          <span v-if="item.badge" class="nav-badge" :style="{ backgroundColor: item.badgeColor }">{{ item.badge }}</span>
          <span v-if="item.isNew" class="nav-new-tag">جدید</span>
          <svg v-if="hasChildren(item)" class="dropdown-arrow" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/>
          </svg>
        </NuxtLink>

        <!-- زیرمنوی معمولی -->
        <ul v-if="hasChildren(item) && !item.isMegaMenu" class="dropdown-menu" :class="{ 'is-open': openDropdownId === item.id }">
          <li v-for="child in item.children" :key="child.id" class="dropdown-item">
            <NuxtLink
              :to="child.path"
              :target="child.openInNewTab ? '_blank' : '_self'"
              class="dropdown-link"
              :class="{ 'is-active': isActive(child.path) }"
            >
              <i v-if="child.icon && child.iconType === 'icon'" :class="child.icon" class="dropdown-icon"></i>
              <img v-else-if="child.imageURL && child.iconType === 'image'" :src="child.imageURL" :alt="child.title" class="dropdown-image" />
              <span>{{ child.title }}</span>
              <span v-if="child.badge" class="dropdown-badge" :style="{ backgroundColor: child.badgeColor }">{{ child.badge }}</span>
              <span v-if="child.isNew" class="dropdown-new-tag">جدید</span>
            </NuxtLink>
            
            <!-- زیرمنوی سطح سوم -->
            <ul v-if="hasChildren(child)" class="sub-dropdown-menu">
              <li v-for="grandChild in child.children" :key="grandChild.id" class="sub-dropdown-item">
                <NuxtLink
                  :to="grandChild.path"
                  :target="grandChild.openInNewTab ? '_blank' : '_self'"
                  class="sub-dropdown-link"
                >
                  {{ grandChild.title }}
                </NuxtLink>
              </li>
            </ul>
          </li>
        </ul>

        <!-- مگامنو -->
        <div v-else-if="item.isMegaMenu && hasChildren(item)" class="mega-menu" :class="{ 'is-open': openDropdownId === item.id }" :style="{ width: item.megaWidth || '100%' }">
          <div class="mega-container" :style="{ gridTemplateColumns: `repeat(${item.megaColumns || 4}, 1fr)` }">
            <div v-for="child in item.children" :key="child.id" class="mega-column">
              <h3 v-if="child.title" class="mega-title">
                <NuxtLink :to="child.path" class="mega-title-link">
                  {{ child.title }}
                </NuxtLink>
              </h3>
              <img v-if="child.imageURL && child.featured" :src="child.imageURL" :alt="child.title" class="mega-featured-image" />
              <p v-if="child.description" class="mega-description">{{ child.description }}</p>
              
              <!-- زیرآیتم‌های مگامنو -->
              <ul v-if="hasChildren(child)" class="mega-list">
                <li v-for="grandChild in child.children" :key="grandChild.id" class="mega-list-item">
                  <NuxtLink :to="grandChild.path" class="mega-link">
                    {{ grandChild.title }}
                  </NuxtLink>
                </li>
              </ul>
            </div>
          </div>
        </div>
      </li>
    </ul>
    
    <!-- خالی -->
    <div v-else class="nav-empty">هیچ آیتمی در منو وجود ندارد</div>
  </nav>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useMenuManagement } from '~/composables/useMenuManagement'

const props = defineProps<{
  menuId?: number // ID منوی مورد نظر
  menuSlug?: string // Slug منوی مورد نظر
}>()

const route = useRoute()
const { menus, fetchMenus, isLoading } = useMenuManagement()

const openDropdownId = ref<number | null>(null)

// پیدا کردن منوی مورد نظر
const selectedMenu = computed(() => {
  console.log('🔍 All menus:', menus.value)
  console.log('🔍 Looking for menuId:', props.menuId, 'menuSlug:', props.menuSlug)
  
  if (!menus.value.length) {
    console.log('⚠️ No menus loaded')
    return null
  }
  
  // اگر menuId داده شده، از آن استفاده کن
  if (props.menuId) {
    const found = menus.value.find(m => m.id === props.menuId && m.enabled)
    console.log('🎯 Found by ID:', found)
    return found
  }
  
  // اگر menuSlug داده شده، از آن استفاده کن
  if (props.menuSlug) {
    const found = menus.value.find(m => m.slug === props.menuSlug && m.enabled)
    console.log('🎯 Found by slug:', found)
    return found
  }
  
  // وگرنه اولین منوی فعال رو برگردون
  const firstActive = menus.value.find(m => m.enabled)
  console.log('🎯 Using first active menu:', firstActive)
  return firstActive
})

// آیتم‌های منو (فقط آیتم‌های فعال و سطح اول)
const menuItems = computed(() => {
  console.log('🔍 HeaderWidgetMenu - selectedMenu:', selectedMenu.value)
  console.log('🔍 HeaderWidgetMenu - items:', selectedMenu.value?.items)
  
  if (!selectedMenu.value?.items) {
    console.log('⚠️ No menu or items found')
    return []
  }
  
  const items = selectedMenu.value.items
    .filter(item => item.enabled && !item.parentId)
    .sort((a, b) => (a.order || 0) - (b.order || 0))
  
  console.log('✅ Filtered menu items:', items)
  return items
})

const hasChildren = (item: any) => {
  return item.children && item.children.length > 0
}

const isActive = (path: string) => {
  return route.path === path
}

const openDropdown = (item: any) => {
  if (hasChildren(item)) {
    openDropdownId.value = item.id
  }
}

const closeAllDropdowns = () => {
  openDropdownId.value = null
}

// بارگذاری منوها
onMounted(async () => {
  if (menus.value.length === 0) {
    await fetchMenus()
  }
})
</script>

<style scoped>
.main-nav {
  position: relative;
  height: 100%;
  display: flex;
  align-items: center;
}

.nav-loading,
.nav-empty {
  padding: 0 1rem;
  color: #9ca3af;
  font-size: 0.875rem;
}

/* منوی اصلی */
.nav-menu {
  display: flex;
  list-style: none;
  margin: 0;
  padding: 0;
  gap: 0.5rem;
  height: 100%;
}

.nav-item {
  position: relative;
  display: flex;
  align-items: center;
  height: 100%;
}

.nav-link {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.75rem 1rem;
  color: #374151;
  text-decoration: none;
  font-weight: 500;
  transition: all 0.2s;
  height: 100%;
  white-space: nowrap;
}

.nav-link:hover,
.nav-link.is-active {
  color: #3b82f6;
  background: #eff6ff;
}

.nav-icon,
.dropdown-icon {
  font-size: 1.125rem;
}

.nav-image,
.dropdown-image {
  width: 20px;
  height: 20px;
  object-fit: contain;
}

.nav-badge,
.dropdown-badge {
  font-size: 0.625rem;
  padding: 0.125rem 0.375rem;
  border-radius: 9999px;
  background: #ef4444;
  color: white;
  font-weight: bold;
}

.nav-new-tag,
.dropdown-new-tag {
  font-size: 0.625rem;
  padding: 0.125rem 0.375rem;
  border-radius: 0.25rem;
  background: #10b981;
  color: white;
  font-weight: bold;
}

.dropdown-arrow {
  width: 16px;
  height: 16px;
  transition: transform 0.2s;
}

.nav-item:hover .dropdown-arrow {
  transform: rotate(180deg);
}

/* زیرمنوی معمولی */
.dropdown-menu {
  position: absolute;
  top: 100%;
  right: 0;
  min-width: 220px;
  background: white;
  border: 1px solid #e5e7eb;
  border-radius: 0.5rem;
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
  list-style: none;
  margin: 0;
  padding: 0.5rem 0;
  opacity: 0;
  visibility: hidden;
  transform: translateY(-10px);
  transition: all 0.2s;
  z-index: 50;
}

.nav-item:hover .dropdown-menu {
  opacity: 1;
  visibility: visible;
  transform: translateY(0);
}

.dropdown-item {
  position: relative;
}

.dropdown-link {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.75rem 1rem;
  color: #374151;
  text-decoration: none;
  transition: all 0.2s;
}

.dropdown-link:hover,
.dropdown-link.is-active {
  background: #f3f4f6;
  color: #3b82f6;
}

/* زیرمنوی سطح سوم */
.sub-dropdown-menu {
  position: absolute;
  top: 0;
  left: 100%;
  min-width: 200px;
  background: white;
  border: 1px solid #e5e7eb;
  border-radius: 0.5rem;
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
  list-style: none;
  margin: 0;
  padding: 0.5rem 0;
  opacity: 0;
  visibility: hidden;
  transition: all 0.2s;
}

.dropdown-item:hover .sub-dropdown-menu {
  opacity: 1;
  visibility: visible;
}

.sub-dropdown-link {
  display: block;
  padding: 0.75rem 1rem;
  color: #374151;
  text-decoration: none;
  transition: all 0.2s;
}

.sub-dropdown-link:hover {
  background: #f3f4f6;
  color: #3b82f6;
}

/* مگامنو */
.mega-menu {
  position: absolute;
  top: 100%;
  right: 0;
  background: white;
  border: 1px solid #e5e7eb;
  border-radius: 0.5rem;
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
  padding: 2rem;
  opacity: 0;
  visibility: hidden;
  transform: translateY(-10px);
  transition: all 0.2s;
  z-index: 50;
  max-width: 1200px;
}

.nav-item:hover .mega-menu {
  opacity: 1;
  visibility: visible;
  transform: translateY(0);
}

.mega-container {
  display: grid;
  gap: 2rem;
}

.mega-column {
  min-width: 0;
}

.mega-title {
  font-size: 1rem;
  font-weight: 600;
  margin: 0 0 1rem 0;
  color: #111827;
}

.mega-title-link {
  color: inherit;
  text-decoration: none;
  transition: color 0.2s;
}

.mega-title-link:hover {
  color: #3b82f6;
}

.mega-featured-image {
  width: 100%;
  height: auto;
  border-radius: 0.5rem;
  margin-bottom: 1rem;
}

.mega-description {
  font-size: 0.875rem;
  color: #6b7280;
  margin: 0 0 1rem 0;
  line-height: 1.5;
}

.mega-list {
  list-style: none;
  margin: 0;
  padding: 0;
}

.mega-list-item {
  margin-bottom: 0.5rem;
}

.mega-link {
  display: block;
  padding: 0.5rem 0;
  color: #374151;
  text-decoration: none;
  transition: color 0.2s;
  font-size: 0.875rem;
}

.mega-link:hover {
  color: #3b82f6;
}

/* Responsive */
@media (max-width: 1024px) {
  .mega-menu {
    max-width: 90vw;
  }
  
  .mega-container {
    grid-template-columns: repeat(2, 1fr) !important;
  }
}

@media (max-width: 768px) {
  .nav-menu {
    flex-direction: column;
    height: auto;
  }
  
  .dropdown-menu,
  .mega-menu {
    position: static;
    opacity: 1;
    visibility: visible;
    transform: none;
  }
  
  .mega-container {
    grid-template-columns: 1fr !important;
  }
}
</style>
