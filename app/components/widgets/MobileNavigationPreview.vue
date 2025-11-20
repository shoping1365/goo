<template>
  <div class="mobile-navigation-preview">
    <div class="preview-header">
      <h3 class="preview-title">پیش‌نمایش ناوبری موبایل</h3>
      <div class="preview-subtitle">{{ navigationData?.name || 'نام ناوبری' }}</div>
    </div>

    <!-- پیش‌نمایش موبایل -->
    <div class="preview-mobile-container">
      <div class="preview-mobile-screen">
        <!-- شبیه‌سازی نوار وضعیت -->
        <div class="preview-status-bar">
          <span class="preview-time">2:39 PM</span>
          <div class="preview-signal">●●●●●</div>
        </div>

        <!-- محتوای صفحه -->
        <div class="preview-page-content">
          <!-- هدر ساده -->
          <div
class="preview-simple-header" :style="{ 
            backgroundColor: navigationData?.background_color || '#f8fafc',
            color: navigationData?.text_color || '#000000'
          }">
            <div class="preview-header-content">
              <div class="preview-header-logo">{{ navigationData?.name || 'نام سایت' }}</div>
              <div class="preview-header-icons">
                <div class="preview-header-icon">🔍</div>
                <div class="preview-header-icon">🛒</div>
              </div>
            </div>
          </div>

          <!-- محتوای نمونه -->
          <div class="preview-content-area">
            <div class="preview-section">
              <h3 class="preview-section-title">دسته بندی های محبوب</h3>
              <div class="preview-products">
                <div class="preview-product-card">
                  <div class="preview-product-icon">🔧</div>
                  <span class="preview-product-text">پرچ کن و مهره پرچ</span>
                </div>
                <div class="preview-product-card">
                  <div class="preview-product-icon">🔩</div>
                  <span class="preview-product-text">آچار و بکس</span>
                </div>
                <div class="preview-product-card">
                  <div class="preview-product-icon">⚡</div>
                  <span class="preview-product-text">دستگاه جوشکاری</span>
                </div>
              </div>
            </div>
          </div>

          <!-- ناوبری پایینی -->
          <div class="preview-bottom-nav">
            <div 
              v-for="item in navigationItems" 
              :key="item.id"
              class="preview-bottom-nav-item"
              :class="{ 'active': item.id === 'home' }"
            >
              <div class="preview-nav-icon">
                <!-- آیکون خانه -->
                <svg v-if="getItemIcon(item.id) === 'home'" class="icon-svg" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"></path>
                </svg>
                
                <!-- آیکون دسته‌ها (شبکه) -->
                <svg v-else-if="getItemIcon(item.id) === 'grid'" class="icon-svg" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z"></path>
                </svg>
                
                <!-- آیکون سبد خرید -->
                <div v-else-if="getItemIcon(item.id) === 'shopping-cart'" class="cart-icon-container">
                  <svg class="icon-svg" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4m0 0L7 13m0 0l-2.5 5M7 13l2.5 5m0 0h9"></path>
                  </svg>
                  <div class="cart-badge">0</div>
                </div>
                
                <!-- آیکون تماس -->
                <svg v-else-if="getItemIcon(item.id) === 'phone'" class="icon-svg" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z"></path>
                </svg>
                
                <!-- آیکون ورود -->
                <svg v-else-if="getItemIcon(item.id) === 'user'" class="icon-svg" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
                </svg>
                
                <!-- آیکون حالت تاریک -->
                <svg v-else-if="getItemIcon(item.id) === 'moon'" class="icon-svg" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z"></path>
                </svg>
              </div>
              <span class="preview-nav-text">{{ item.title }}</span>
            </div>
          </div>
        </div>

      </div>
    </div>

    <!-- اطلاعات ناوبری -->
    <div class="preview-info mt-auto">
      <div class="info-item">
        <span class="info-label">پلتفرم:</span>
        <span class="info-value">{{ platformText }}</span>
      </div>
      <div class="info-item">
        <span class="info-label">وضعیت:</span>
        <span class="info-value" :class="{ 'active': navigationData?.is_active, 'inactive': !navigationData?.is_active }">
          {{ navigationData?.is_active ? 'فعال' : 'غیرفعال' }}
        </span>
      </div>
      <div class="info-item">
        <span class="info-label">صفحات:</span>
        <span class="info-value">{{ pageSelectionText }}</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

// Props
const props = defineProps({
  navigationData: {
    type: Object,
    default: () => ({})
  }
})

// Static items to prevent memory leaks
const ALL_NAVIGATION_ITEMS = {
  'home': { id: 'home', title: 'خانه' },
  'categories': { id: 'categories', title: 'دسته‌ها' },
  'cart': { id: 'cart', title: 'سبد خرید' },
  'contact': { id: 'contact', title: 'تماس' },
  'login': { id: 'login', title: 'ورود' },
  'dark-mode': { id: 'dark-mode', title: 'حالت تاریک' }
}

// Computed properties for better performance
const navigationItems = computed(() => {
  const selectedItems = props.navigationData?.navigation_items || []
  return selectedItems.map(itemId => ALL_NAVIGATION_ITEMS[itemId]).filter(Boolean)
})

// Static mappings to prevent memory leaks
const PLATFORM_MAP = {
  'mobile': 'موبایل',
  'app': 'اپلیکیشن',
  'both': 'هر دو'
}

const PAGE_SELECTION_MAP = {
  'all': 'همه صفحات',
  'specific': 'صفحات خاص',
  'excluded': 'مستثنی شده'
}

const platformText = computed(() => {
  return PLATFORM_MAP[props.navigationData?.platform] || 'نامشخص'
})

const pageSelectionText = computed(() => {
  return PAGE_SELECTION_MAP[props.navigationData?.page_selection] || 'همه صفحات'
})

// Static icon mapping to prevent memory leaks
const ICON_MAP = {
  'home': 'home',
  'categories': 'grid',
  'cart': 'shopping-cart',
  'contact': 'phone',
  'login': 'user',
  'dark-mode': 'moon'
}

// Methods
const getItemIcon = (itemId) => {
  return ICON_MAP[itemId] || 'grid'
}
</script>

<style scoped>
.mobile-navigation-preview {
  background: white;
  border-radius: 12px;
  padding: 24px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
  height: 100%;
  min-height: 600px;
  position: sticky;
  top: 20px;
  display: flex;
  flex-direction: column;
}

.preview-header {
  text-align: center;
  margin-bottom: 24px;
  padding-bottom: 16px;
  border-bottom: 2px solid #e5e7eb;
}

.preview-title {
  font-size: 1.25rem;
  font-weight: bold;
  color: #374151;
  margin: 0 0 8px 0;
}

.preview-subtitle {
  font-size: 0.9rem;
  color: #6b7280;
  margin: 0;
}

.preview-mobile-container {
  display: flex;
  justify-content: center;
  margin-bottom: 24px;
}

.preview-mobile-screen {
  background: linear-gradient(145deg, #2d3748, #1a202c);
  border-radius: 20px;
  padding: 6px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.3);
  max-width: 280px;
  width: 100%;
  position: relative;
  border: 2px solid #4a5568;
  min-height: 500px;
}

.preview-mobile-screen::before {
  content: '';
  position: absolute;
  top: 15px;
  left: 50%;
  transform: translateX(-50%);
  width: 50px;
  height: 3px;
  background: #4a5568;
  border-radius: 2px;
}

.preview-status-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px 16px;
  background: white;
  border-radius: 15px 15px 0 0;
  font-size: 12px;
  color: #374151;
  font-weight: 600;
}

.preview-time {
  font-size: 12px;
}

.preview-signal {
  font-size: 10px;
}

.preview-page-content {
  background: white;
  border-radius: 0 0 15px 15px;
  min-height: 400px;
  display: flex;
  flex-direction: column;
}

.preview-simple-header {
  background: #f8fafc;
  padding: 12px 16px;
  border-bottom: 1px solid #e5e7eb;
}

.preview-header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.preview-header-logo {
  font-weight: bold;
  color: #f97316;
  font-size: 1.1rem;
}

.preview-header-icons {
  display: flex;
  gap: 8px;
}

.preview-header-icon {
  width: 28px;
  height: 28px;
  border-radius: 6px;
  background: rgba(255, 255, 255, 0.8);
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 12px;
}

.preview-content-area {
  flex: 1;
  padding: 16px;
}

.preview-section {
  margin-bottom: 20px;
}

.preview-section-title {
  font-size: 14px;
  color: #ef4444;
  font-weight: 600;
  margin: 0 0 12px 0;
  text-align: right;
}

.preview-products {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 8px;
}

.preview-product-card {
  background: #f9fafb;
  border-radius: 8px;
  padding: 12px 8px;
  text-align: center;
  border: 1px solid #e5e7eb;
}

.preview-product-icon {
  font-size: 20px;
  margin-bottom: 6px;
}

.preview-product-text {
  font-size: 9px;
  color: #374151;
  font-weight: 500;
  line-height: 1.3;
}

.preview-bottom-nav {
  background: white;
  border-top: 1px solid #e5e7eb;
  display: flex;
  justify-content: space-around;
  padding: 8px 0;
  margin-top: auto;
}

.preview-bottom-nav-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 4px;
  padding: 6px 8px;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.2s;
  min-width: 50px;
}

.preview-bottom-nav-item.active {
  background: #fef2f2;
}

.preview-bottom-nav-item.active .preview-nav-text {
  color: #ef4444;
}

.preview-nav-icon {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
}

.icon-svg {
  width: 20px;
  height: 20px;
  stroke: currentColor;
}

.cart-icon-container {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
}

.cart-badge {
  position: absolute;
  top: -4px;
  right: -4px;
  background: #ef4444;
  color: white;
  border-radius: 50%;
  width: 12px;
  height: 12px;
  font-size: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
}

.preview-nav-text {
  font-size: 9px;
  color: #6b7280;
  font-weight: 500;
  text-align: center;
}

.preview-info {
  background: #f9fafb;
  border-radius: 8px;
  padding: 16px;
  border: 1px solid #e5e7eb;
}

.info-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}

.info-item:last-child {
  margin-bottom: 0;
}

.info-label {
  font-size: 12px;
  color: #6b7280;
  font-weight: 500;
}

.info-value {
  font-size: 12px;
  color: #374151;
  font-weight: 600;
}

.info-value.active {
  color: #10b981;
}

.info-value.inactive {
  color: #ef4444;
}

@media (max-width: 768px) {
  .mobile-navigation-preview {
    padding: 16px;
  }
  
  .preview-mobile-screen {
    max-width: 250px;
  }
  
  .preview-products {
    grid-template-columns: repeat(2, 1fr);
  }
  
  .preview-bottom-nav-item {
    min-width: 40px;
    padding: 4px 6px;
  }
  
  .icon-svg {
    width: 18px;
    height: 18px;
  }
  
  .cart-badge {
    width: 10px;
    height: 10px;
    font-size: 7px;
  }
  
  .preview-nav-text {
    font-size: 8px;
  }
}
</style>