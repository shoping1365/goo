<template>
  <!-- هدر موبایل -->
  <div
class="mobile-header relative" :class="{
    'full-width': config.banner_width === 800,
    'center-width': config.banner_width === 600
  }" :style="{ 
    height: `${config.height || 400}px`,
    width: config.banner_width === 800 ? '100%' : config.banner_width === 600 ? `${config.center_width || 1000}px` : '100%',
    maxWidth: config.banner_width === 600 ? `${config.center_width || 1000}px` : '100%',
    margin: config.banner_width === 600 ? '0 auto' : '0',
    backgroundColor: config.bg_enabled ? config.bg_color : 'transparent',
    paddingTop: config.padding_top ? `${config.padding_top}px` : '0',
    paddingBottom: config.padding_bottom ? `${config.padding_bottom}px` : '0',
    marginLeft: config.banner_width === 800 && config.margin_left ? `${config.margin_left}px` : '0',
    marginRight: config.banner_width === 800 && config.margin_right ? `${config.margin_right}px` : '0',
    '--desktop-height': `${config.height || 400}px`,
    '--mobile-height': `${config.mobile_height || config.height || 400}px`
  }">

    <!-- عنوان هدر (در صورت وجود) -->
    <div v-if="config.title" class="w-full text-center font-bold text-xl md:text-2xl mb-2 mt-4 text-gray-800">
      {{ config.title }}
    </div>

    <!-- نمایش هدرهای موبایل -->
    <div v-if="config.headers && config.headers.length > 0" class="relative">
      <div 
        v-for="(header, index) in config.headers" 
        :key="header.id || index"
        class="mobile-header-item"
        :style="{ height: `${config.height || 400}px` }"
      >
        <!-- پیش‌نمایش هدر موبایل -->
        <div class="mobile-preview-device">
          <div class="mobile-screen" :style="{ backgroundColor: header.background_color || '#ffffff', color: header.text_color || '#000000' }">
            <!-- لوگو -->
            <div v-if="header.logo_url" class="preview-logo-container">
              <img :src="header.logo_url" :alt="header.logo_alt || 'لوگو'" class="preview-logo-image" />
            </div>
            
            <!-- آیتم‌های هدر -->
            <div class="preview-items-container">
              <!-- جستجو -->
              <div v-if="header.show_search" class="preview-search">
                <span class="preview-search-icon">🔍</span>
                <span class="preview-search-text">جستجو...</span>
              </div>
              
              <!-- سبد خرید -->
              <div v-if="header.show_cart" class="preview-cart">
                <span class="preview-cart-icon">🛒</span>
              </div>
              
              <!-- منوی کاربر -->
              <div v-if="header.show_user_menu" class="preview-user">
                <span class="preview-user-icon">👤</span>
              </div>
              
              <!-- اعلان‌ها -->
              <div v-if="header.show_notifications" class="preview-notification">
                <span class="preview-notification-icon">🔔</span>
              </div>
              
              <!-- دکمه منو -->
              <div v-if="header.show_menu_button" class="preview-menu">
                <span class="preview-menu-icon">☰</span>
              </div>
            </div>
          </div>
        </div>

        <!-- اطلاعات هدر -->
        <div class="header-info">
          <h3 class="header-name">{{ header.name }}</h3>
          <p v-if="header.description" class="header-description">{{ header.description }}</p>
          <div class="header-meta">
            <span class="meta-item">
              <strong>نوع:</strong> {{ getHeaderTypeLabel(header.header_type) }}
            </span>
            <span class="meta-item">
              <strong>پلتفرم:</strong> {{ getPlatformLabel(header.platform) }}
            </span>
            <span class="meta-item">
              <strong>وضعیت:</strong> 
              <span :class="header.is_active ? 'status-active' : 'status-inactive'">
                {{ header.is_active ? 'فعال' : 'غیرفعال' }}
              </span>
            </span>
          </div>
        </div>
      </div>
    </div>

    <!-- اگر هدر وجود نداشت، پیام جایگزین -->
    <div v-else class="w-full flex items-center justify-center bg-gray-200 text-gray-400 rounded-lg mt-4" :style="{ height: `${config.height || 400}px` }">
      هدر موبایل انتخاب نشده است
    </div>

  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  widget: {
    type: Object,
    required: true
  }
})

// Config با محافظت از SSR
const config = computed(() => {
  if (!props.widget || !props.widget.config) {
    return {
      headers: [],
      banner_width: 800,
      center_width: 1000,
      height: 400,
      mobile_height: 400,
      show_title: true,
      show_description: true,
      bg_enabled: false,
      bg_color: '#ffffff',
      padding_top: 0,
      padding_bottom: 0,
      margin_left: 0,
      margin_right: 0,
      title: ''
    }
  }

  const widgetConfig = props.widget.config
  
  return {
    headers: Array.isArray(widgetConfig.headers) ? widgetConfig.headers : [],
    banner_width: widgetConfig.banner_width || 800,
    center_width: widgetConfig.center_width || 1000,
    height: widgetConfig.height || 400,
    mobile_height: widgetConfig.mobile_height || widgetConfig.height || 400,
    show_title: widgetConfig.show_title !== undefined ? widgetConfig.show_title : true,
    show_description: widgetConfig.show_description !== undefined ? widgetConfig.show_description : true,
    bg_enabled: widgetConfig.bg_enabled || false,
    bg_color: widgetConfig.bg_color || '#ffffff',
    padding_top: widgetConfig.padding_top || 0,
    padding_bottom: widgetConfig.padding_bottom || 0,
    margin_left: widgetConfig.margin_left || 0,
    margin_right: widgetConfig.margin_right || 0,
    title: widgetConfig.title || ''
  }
})

// توابع کمکی با محافظت از خطا
const getPlatformLabel = (platform) => {
  if (!platform) return 'نامشخص'
  const labels = {
    'mobile': 'موبایل',
    'app': 'اپلیکیشن',
    'both': 'هر دو'
  }
  return labels[platform] || platform
}

const getHeaderTypeLabel = (headerType) => {
  if (!headerType) return 'نامشخص'
  const labels = {
    'default': 'پیش‌فرض',
    'minimal': 'مینیمال',
    'full': 'کامل'
  }
  return labels[headerType] || headerType
}
</script>

<style scoped>
.mobile-header {
  overflow: hidden;
}

.mobile-header-item {
  background: white;
  border-radius: 12px;
  padding: 20px;
  transition: all 0.3s ease;
  width: 100%;
  height: auto;
  object-fit: cover;
}

.mobile-header-item:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

/* تنظیم ارتفاع موبایل */
.responsive-header-image {
  height: var(--desktop-height);
}

@media (max-width: 767px) {
  .responsive-header-image {
    height: var(--mobile-height) ;
  }
}

/* برای حالت تمام عرض */
.mobile-header.full-width {
  width: 100vw ;
  max-width: 100vw;
  margin-left: calc(-50vw + 50%) ;
  margin-right: calc(-50vw + 50%) ;
}

/* برای حالت وسط */
.mobile-header.center-width {
  margin-left: auto ;
  margin-right: auto ;
  display: block ;
}

.mobile-preview-device {
  background: linear-gradient(145deg, #2d3748, #1a202c);
  border-radius: 20px;
  padding: 6px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
  max-width: 250px;
  width: 100%;
  position: relative;
  border: 2px solid #4a5568;
  margin: 0 auto 15px auto;
}

.mobile-preview-device::before {
  content: '';
  position: absolute;
  top: 15px;
  left: 50%;
  transform: translateX(-50%);
  width: 40px;
  height: 3px;
  background: #4a5568;
  border-radius: 2px;
}

.mobile-screen {
  background: white;
  border-radius: 18px;
  padding: 15px;
  min-height: 200px;
  position: relative;
  overflow: hidden;
  padding-top: 30px;
}

.mobile-screen::before {
  content: '2:39 PM';
  position: absolute;
  top: 8px;
  left: 15px;
  color: #374151;
  font-size: 12px;
  font-weight: 600;
  z-index: 10;
}

.mobile-screen::after {
  content: '●●●●●';
  position: absolute;
  top: 8px;
  right: 15px;
  color: #374151;
  font-size: 10px;
  z-index: 10;
}

.preview-logo-container {
  text-align: center;
  margin-bottom: 15px;
}

.preview-logo-image {
  max-width: 80px;
  max-height: 30px;
  object-fit: contain;
}

.preview-items-container {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
  padding: 10px 0;
}

/* استایل‌های آیتم‌های مختلف */
.preview-logo {
  font-weight: bold;
  color: #f97316;
  font-size: 1rem;
}

.preview-search {
  background: white;
  border-radius: 15px;
  padding: 6px 10px;
  display: flex;
  align-items: center;
  gap: 4px;
  border: 1px solid #d1d5db;
  flex: 1;
  min-width: 80px;
}

.preview-search-icon {
  font-size: 10px;
  color: #6b7280;
}

.preview-search-text {
  font-size: 10px;
  color: #9ca3af;
}

.preview-cart {
  position: relative;
  background: #f3f4f6;
  border-radius: 15px;
  padding: 6px 10px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.preview-cart-icon {
  font-size: 12px;
}

.preview-cart-badge {
  position: absolute;
  top: -2px;
  right: -2px;
  background: #10b981;
  color: white;
  border-radius: 50%;
  width: 12px;
  height: 12px;
  font-size: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.preview-user {
  background: #e5e7eb;
  border-radius: 15px;
  padding: 6px 10px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.preview-user-icon {
  font-size: 12px;
}

.preview-notification {
  position: relative;
  background: #f3f4f6;
  border-radius: 15px;
  padding: 6px 10px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.preview-notification-icon {
  font-size: 12px;
}

.preview-notification-badge {
  position: absolute;
  top: -2px;
  right: -2px;
  background: #dc2626;
  color: white;
  border-radius: 50%;
  width: 12px;
  height: 12px;
  font-size: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.preview-menu {
  background: #e5e7eb;
  border-radius: 6px;
  padding: 6px 10px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.preview-menu-icon {
  font-size: 12px;
}

.preview-text {
  font-size: 12px;
  color: #374151;
}

.preview-banner {
  background: #fef3c7;
  border: 1px solid #f59e0b;
  border-radius: 6px;
  padding: 6px 10px;
  width: 100%;
  text-align: center;
}

.preview-banner-text {
  font-size: 10px;
  color: #374151;
  line-height: 1.2;
}

.preview-default {
  font-size: 12px;
  color: #6b7280;
}

.header-info {
  text-align: center;
}

.header-name {
  font-size: 1.1rem;
  font-weight: 600;
  color: #111827;
  margin: 0 0 8px 0;
}

.header-description {
  color: #6b7280;
  margin: 0 0 12px 0;
  line-height: 1.4;
  font-size: 0.9rem;
}

.header-meta {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.meta-item {
  font-size: 0.8rem;
  color: #6b7280;
}

.status-active {
  color: #166534;
  font-weight: 600;
}

.status-inactive {
  color: #dc2626;
  font-weight: 600;
}

.empty-state {
  text-align: center;
  padding: 40px;
  color: #6b7280;
}
</style>