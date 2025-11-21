<template>
  <div class="mobile-app-header-management-page">
    <div class="header-bg">
      <div class="page-header-flex">
        <h1 class="page-title">مدیریت هدر موبایل و اپلیکیشن</h1>
        <button class="btn btn-primary new-item-btn" @click="addNewMobileAppHeader">افزودن هدر موبایل و اپلیکیشن</button>
      </div>
    </div>

    <div class="headers-section">
      
      <!-- Loading indicator -->
      <div v-if="loading" class="loading-state">
        <p>در حال بارگذاری...</p>
      </div>
      
      <!-- Error state -->
      <div v-else-if="error" class="error-state">
        <p>{{ error }}</p>
        <button class="btn btn-primary" @click="loadMobileAppHeaders">تلاش مجدد</button>
      </div>
      
      <!-- Empty state -->
      <div v-else-if="mobileAppHeaders.length === 0" class="empty-state">
      </div>
      
      <!-- Headers list -->
      <div v-else class="headers-list">
        <div v-for="header in mobileAppHeaders" :key="header.id" class="header-item">
          <div class="header-info">
            <div class="header-title-row">
              <h4>{{ header.name }}</h4>
              <span v-if="header.is_active" class="status-badge status-active">فعال</span>
              <span v-else class="status-badge status-inactive">غیرفعال</span>
            </div>
            <p v-if="header.description" class="header-description">{{ header.description }}</p>
            <div class="header-meta">
              <span class="meta-item">
                <strong>پلتفرم:</strong> {{ getPlatformLabel(header.platform) }}
              </span>
              <span class="meta-item">
                <strong>تعداد لایه‌ها:</strong> {{ header.layers?.length || 0 }}
              </span>
              <span class="meta-item">
                <strong>تاریخ ایجاد:</strong> {{ formatDate(header.created_at || header.createdAt) }}
              </span>
            </div>
          </div>
          <div class="header-actions">
            <button 
              class="btn btn-secondary btn-sm" 
              title="پیش‌نمایش"
              @click="previewMobileAppHeader(header)"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
              </svg>
            </button>
            <button 
              class="btn btn-primary btn-sm" 
              title="ویرایش"
              @click="editMobileAppHeader(header)"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
              </svg>
            </button>
            <button 
              class="btn btn-danger btn-sm" 
              title="حذف"
              @click="deleteMobileAppHeader(header)"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
              </svg>
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Preview Modal -->
    <div v-if="showPreview" class="modal-overlay" @click="closePreview">
      <div class="modal-content" @click.stop>
        <div class="modal-header">
          <h3>پیش‌نمایش هدر موبایل و اپلیکیشن</h3>
          <button class="btn btn-secondary" @click="closePreview">بستن</button>
        </div>
        <div class="modal-body">
          <div class="preview-container">
            <div class="mobile-preview-device">
              <div class="mobile-screen">
                <!-- پیش‌نمایش واقعی بر اساس ساختار لایه‌ها -->
                <div v-if="previewHeader?.layers && previewHeader.layers.length > 0" class="preview-header-container">
                  <div 
                    v-for="(layer, layerIndex) in previewHeader.layers" 
                    :key="layerIndex"
                    class="preview-layer"
                    :style="getLayerStyle(layer)"
                  >
                    <div 
                      v-for="(item, itemIndex) in getLayerItems(layer)" 
                      :key="itemIndex"
                      class="preview-item"
                      :style="getItemStyle(item)"
                    >
                      <span v-if="item.type === 'logo'" class="preview-logo">{{ item.text || 'لوگو' }}</span>
                      <span v-else-if="item.type === 'search'" class="preview-search">
                        <span class="preview-search-icon">🔍</span>
                        <span class="preview-search-text">{{ item.text || 'جستجو...' }}</span>
                      </span>
                      <span v-else-if="item.type === 'cart'" class="preview-cart">
                        <span class="preview-cart-icon">🛒</span>
                        <span v-if="item.badge" class="preview-cart-badge">{{ item.badge }}</span>
                      </span>
                      <span v-else-if="item.type === 'user'" class="preview-user">
                        <span class="preview-user-icon">👤</span>
                      </span>
                      <span v-else-if="item.type === 'notification'" class="preview-notification">
                        <span class="preview-notification-icon">🔔</span>
                        <span v-if="item.badge" class="preview-notification-badge">{{ item.badge }}</span>
                      </span>
                      <span v-else-if="item.type === 'menu'" class="preview-menu">
                        <span class="preview-menu-icon">☰</span>
                      </span>
                      <span v-else-if="item.type === 'text'" class="preview-text">{{ item.text || 'متن' }}</span>
                      <span v-else-if="item.type === 'banner'" class="preview-banner">
                        <span class="preview-banner-text">{{ item.text || 'بنر تبلیغاتی' }}</span>
                      </span>
                      <span v-else class="preview-default">{{ item.text || 'آیتم' }}</span>
                    </div>
                    <div v-if="layer.show_separator" class="preview-separator" :style="getSeparatorStyle(layer)"></div>
                  </div>
                </div>

                <!-- پیش‌نمایش پیش‌فرض در صورت عدم وجود لایه -->
                <div v-else class="preview-placeholder">
                  <p>هدر موبایل و اپلیکیشن: {{ previewHeader?.name }}</p>
                  <p class="preview-no-layers">هیچ لایه‌ای تعریف نشده است</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';

definePageMeta({
  layout: 'admin-main'
})

interface Item {
  type: string
  text?: string
  badge?: string
  fontSize?: string
  color?: string
  fontWeight?: string
  [key: string]: unknown
}

interface Layer {
  items?: string
  color?: string
  opacity?: number
  padding_top?: number
  padding_right?: number
  padding_bottom?: number
  padding_left?: number
  height?: number
  show_separator?: boolean
  separator_width?: number
  separator_color?: string
  separator_opacity?: number
  [key: string]: unknown
}

interface MobileAppHeader {
  id: number | string
  name: string
  is_active: boolean
  description?: string
  platform: string
  created_at?: string
  createdAt?: string
  layers?: Layer[]
  [key: string]: unknown
}

const router = useRouter()

// State
const mobileAppHeaders = ref<MobileAppHeader[]>([])
const loading = ref(false)
const error = ref('')
const showPreview = ref(false)
const previewHeader = ref<MobileAppHeader | null>(null)

// Methods
const loadMobileAppHeaders = async () => {
  loading.value = true
  error.value = ''
  
  try {
    const response = await $fetch<{ data?: MobileAppHeader[] }>('/api/admin/mobile-app-header-settings')
    const data = response.data
    mobileAppHeaders.value = data || []
  } catch (err: unknown) {
    const message = (err as { data?: { message?: string } })?.data?.message || 'خطا در بارگذاری هدرهای موبایل و اپلیکیشن'
    error.value = message
    console.error('Error loading mobile app headers:', err)
  } finally {
    loading.value = false
  }
}

const addNewMobileAppHeader = () => {
  router.push('/admin/content/mobile-app-header-management/create')
}

const editMobileAppHeader = (header: MobileAppHeader) => {
  router.push(`/admin/content/mobile-app-header-management/edit/${header.id}`)
}

const previewMobileAppHeader = (header: MobileAppHeader) => {
  previewHeader.value = header
  showPreview.value = true
}

const closePreview = () => {
  showPreview.value = false
  previewHeader.value = null
}

const deleteMobileAppHeader = async (header: MobileAppHeader) => {
  if (!confirm(`آیا مطمئن هستید که می‌خواهید هدر "${header.name}" را حذف کنید؟`)) {
    return
  }
  
  try {
    await $fetch(`/api/admin/mobile-app-header-settings/${header.id}`, {
      method: 'DELETE'
    })
    
    // حذف از لیست محلی
    const index = mobileAppHeaders.value.findIndex((h) => h.id === header.id)
    if (index > -1) {
      mobileAppHeaders.value.splice(index, 1)
    }
    
    alert('هدر موبایل و اپلیکیشن با موفقیت حذف شد')
  } catch (err: unknown) {
    const message = (err as { data?: { message?: string } })?.data?.message || 'خطا در حذف هدر موبایل و اپلیکیشن'
    alert(message)
    console.error('Error deleting mobile app header:', err)
  }
}

const getPlatformLabel = (platform: string) => {
  const labels: Record<string, string> = {
    'mobile': 'موبایل',
    'app': 'اپلیکیشن',
    'both': 'هر دو'
  }
  return labels[platform] || platform
}

const formatDate = (dateString: string) => {
  if (!dateString) return '-'
  const date = new Date(dateString)
  return date.toLocaleDateString('fa-IR')
}

const getLayerItems = (layer: Layer): Item[] => {
  if (!layer.items) return []
  
  try {
    return JSON.parse(layer.items)
  } catch (error) {
    console.error('Error parsing layer items:', error)
    return []
  }
}

const getLayerStyle = (layer: Layer): Record<string, string | number> => {
  return {
    backgroundColor: (layer.color as string) || '#ffffff',
    opacity: (layer.opacity as number) || 1,
    padding: `${layer.padding_top || 10}px ${layer.padding_right || 20}px ${layer.padding_bottom || 10}px ${layer.padding_left || 20}px`,
    minHeight: `${layer.height || 50}px`,
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'space-between',
    flexWrap: 'wrap' as const,
    gap: '8px'
  }
}

const getItemStyle = (item: Item) => {
  return {
    display: 'flex',
    alignItems: 'center',
    gap: '4px',
    fontSize: (item.fontSize as string) || '14px',
    color: (item.color as string) || '#333333',
    fontWeight: (item.fontWeight as string) || 'normal'
  }
}

const getSeparatorStyle = (layer: Layer) => {
  return {
    width: '100%',
    height: `${layer.separator_width || 1}px`,
    backgroundColor: (layer.separator_color as string) || '#000000',
    opacity: (layer.separator_opacity as number) || 1,
    margin: '8px 0'
  }
}

// Lifecycle
onMounted(() => {
  loadMobileAppHeaders()
})
</script>

<style scoped>
.mobile-app-header-management-page {
  padding: 20px;
}

.header-bg {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  padding: 30px;
  border-radius: 12px;
  margin-bottom: 30px;
}

.page-header-flex {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.page-title {
  font-size: 2rem;
  font-weight: bold;
  margin: 0;
}

.new-item-btn {
  background: rgba(255, 255, 255, 0.2);
  border: 2px solid rgba(255, 255, 255, 0.3);
  color: white;
  padding: 12px 24px;
  border-radius: 8px;
  font-weight: 600;
  transition: all 0.3s ease;
}

.new-item-btn:hover {
  background: rgba(255, 255, 255, 0.3);
  border-color: rgba(255, 255, 255, 0.5);
}

.headers-section h3 {
  font-size: 1.5rem;
  margin-bottom: 20px;
  color: #374151;
}

.loading-state, .error-state, .empty-state {
  text-align: center;
  padding: 40px;
  color: #6b7280;
}

.error-state {
  color: #dc2626;
}

.headers-list {
  display: grid;
  gap: 20px;
}

.header-item {
  background: white;
  border: 1px solid #e5e7eb;
  border-radius: 12px;
  padding: 24px;
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  transition: all 0.3s ease;
}

.header-item:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  border-color: #d1d5db;
}

.header-info {
  flex: 1;
}

.header-title-row {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 8px;
}

.header-title-row h4 {
  margin: 0;
  font-size: 1.25rem;
  color: #111827;
}

.status-badge {
  padding: 4px 12px;
  border-radius: 20px;
  font-size: 0.875rem;
  font-weight: 600;
}

.status-active {
  background: #dcfce7;
  color: #166534;
}

.status-inactive {
  background: #fee2e2;
  color: #dc2626;
}

.header-description {
  color: #6b7280;
  margin: 8px 0 16px 0;
  line-height: 1.5;
}

.header-meta {
  display: flex;
  gap: 24px;
  flex-wrap: wrap;
}

.meta-item {
  font-size: 0.875rem;
  color: #6b7280;
}

.header-actions {
  display: flex;
  gap: 8px;
}

.btn {
  padding: 8px 16px;
  border-radius: 6px;
  font-weight: 500;
  transition: all 0.2s ease;
  border: none;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 6px;
}

.btn-sm {
  padding: 6px 12px;
  font-size: 0.875rem;
}

.btn-primary {
  background: #3b82f6;
  color: white;
}

.btn-primary:hover {
  background: #2563eb;
}

.btn-secondary {
  background: #6b7280;
  color: white;
}

.btn-secondary:hover {
  background: #4b5563;
}

.btn-danger {
  background: #dc2626;
  color: white;
}

.btn-danger:hover {
  background: #b91c1c;
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal-content {
  background: white;
  border-radius: 12px;
  max-width: 800px;
  width: 90%;
  max-height: 90vh;
  overflow: hidden;
}

.modal-header {
  padding: 20px 24px;
  border-bottom: 1px solid #e5e7eb;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.modal-header h3 {
  margin: 0;
  font-size: 1.25rem;
  color: #111827;
}

.modal-body {
  padding: 24px;
}

.preview-container {
  min-height: 400px;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 20px;
}

.mobile-preview-device {
  background: linear-gradient(145deg, #2d3748, #1a202c);
  border-radius: 30px;
  padding: 8px;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3);
  max-width: 350px;
  width: 100%;
  position: relative;
  border: 2px solid #4a5568;
  min-height: 600px;
}

.mobile-preview-device::before {
  content: '';
  position: absolute;
  top: 20px;
  left: 50%;
  transform: translateX(-50%);
  width: 60px;
  height: 4px;
  background: #4a5568;
  border-radius: 2px;
}

.mobile-screen {
  background: white;
  border-radius: 25px;
  padding: 20px;
  min-height: 580px;
  position: relative;
  overflow: hidden;
  padding-top: 50px;
}

.mobile-screen::before {
  content: '2:39 PM';
  position: absolute;
  top: 10px;
  left: 20px;
  color: #374151;
  font-size: 14px;
  font-weight: 600;
  z-index: 10;
}

.mobile-screen::after {
  content: '●●●●●';
  position: absolute;
  top: 10px;
  right: 20px;
  color: #374151;
  font-size: 12px;
  z-index: 10;
}

.preview-header-container {
  width: 100%;
}

.preview-layer {
  border-radius: 8px;
  margin-bottom: 8px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(0, 0, 0, 0.1);
}

.preview-item {
  display: flex;
  align-items: center;
  gap: 4px;
}

.preview-separator {
  width: 100%;
  margin: 8px 0;
}

.preview-placeholder {
  text-align: center;
  padding: 60px 20px;
  color: #6b7280;
  background: white;
  border-radius: 12px;
  margin-bottom: 20px;
}

.preview-placeholder p {
  font-size: 1.1rem;
  margin: 0;
}

/* استایل‌های پیش‌نمایش - نمونه کلاسیک */
.preview-top-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}

.preview-logo {
  font-weight: bold;
  color: #f97316;
  font-size: 1.4rem;
}

.preview-icons {
  display: flex;
  gap: 6px;
}

.preview-icon {
  width: 36px;
  height: 36px;
  border-radius: 8px;
  background: #e5e7eb;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 16px;
}

.preview-search-section {
  display: flex;
  align-items: center;
  gap: 6px;
}

.preview-search-bar {
  flex: 1;
  background: white;
  border-radius: 24px;
  padding: 12px 16px;
  display: flex;
  align-items: center;
  gap: 8px;
  border: 1px solid #d1d5db;
}

.preview-search-icon {
  font-size: 14px;
  color: #6b7280;
}

.preview-search-text {
  font-size: 14px;
  color: #9ca3af;
}

.preview-cart-icon {
  width: 44px;
  height: 44px;
  border-radius: 22px;
  background: #f3f4f6;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 18px;
  position: relative;
}

.preview-cart-icon::after {
  content: '1';
  position: absolute;
  top: -2px;
  right: -2px;
  background: #10b981;
  color: white;
  border-radius: 50%;
  width: 18px;
  height: 18px;
  font-size: 11px;
  display: flex;
  align-items: center;
  justify-content: center;
}

/* استایل‌های پیش‌نمایش - نمونه با بنر */
.preview-banner-section {
  margin-bottom: 8px;
}

.preview-banner {
  background: #fef3c7;
  border: 1px solid #f59e0b;
  border-radius: 8px;
  padding: 12px 16px;
  position: relative;
  min-height: 50px;
  display: flex;
  align-items: center;
}

.preview-banner-content {
  flex: 1;
}

.preview-banner-text {
  font-size: 13px;
  color: #374151;
  line-height: 1.4;
}

.preview-paperclip {
  position: absolute;
  top: -2px;
  right: 12px;
  font-size: 16px;
  color: #6b7280;
}

.preview-search-section-2 {
  display: flex;
  align-items: center;
  gap: 8px;
}

.preview-logo-left-2 {
  font-size: 1rem;
  color: #374151;
  font-weight: 600;
  flex-shrink: 0;
}

.preview-search-bar-2 {
  flex: 1;
  background: white;
  border-radius: 24px;
  padding: 12px 16px;
  display: flex;
  align-items: center;
  gap: 8px;
  border: 1px solid #d1d5db;
}

.preview-search-text-2 {
  font-size: 13px;
  color: #6b7280;
  flex: 1;
}

.preview-notification-icon {
  width: 44px;
  height: 44px;
  border-radius: 8px;
  background: #f3f4f6;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 18px;
  color: #6b7280;
}

/* استایل‌های پیش‌نمایش - نمونه مینیمال */
.preview-search-section-3 {
  display: flex;
  align-items: center;
}

.preview-search-bar-3 {
  width: 100%;
  background: white;
  border-radius: 12px;
  padding: 16px 20px;
  display: flex;
  align-items: center;
  gap: 12px;
  border: 1px solid #d1d5db;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.preview-logo-left {
  display: flex;
  align-items: center;
  gap: 8px;
  flex-shrink: 0;
}

.preview-logo-text {
  font-size: 13px;
  color: #1f2937;
  font-weight: 600;
}

.preview-search-text-3 {
  font-size: 14px;
  color: #9ca3af;
  flex: 1;
  text-align: right;
}

.preview-search-icon-3 {
  font-size: 18px;
  color: #6b7280;
  flex-shrink: 0;
}

/* استایل‌های آیتم‌های مختلف */
.preview-logo {
  font-weight: bold;
  color: #f97316;
  font-size: 1.2rem;
}

.preview-search {
  background: white;
  border-radius: 20px;
  padding: 8px 12px;
  display: flex;
  align-items: center;
  gap: 6px;
  border: 1px solid #d1d5db;
  flex: 1;
  min-width: 120px;
}

.preview-search-icon {
  font-size: 12px;
  color: #6b7280;
}

.preview-search-text {
  font-size: 12px;
  color: #9ca3af;
}

.preview-cart {
  position: relative;
  background: #f3f4f6;
  border-radius: 20px;
  padding: 8px 12px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.preview-cart-icon {
  font-size: 16px;
}

.preview-cart-badge {
  position: absolute;
  top: -2px;
  right: -2px;
  background: #10b981;
  color: white;
  border-radius: 50%;
  width: 16px;
  height: 16px;
  font-size: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.preview-user {
  background: #e5e7eb;
  border-radius: 20px;
  padding: 8px 12px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.preview-user-icon {
  font-size: 16px;
}

.preview-notification {
  position: relative;
  background: #f3f4f6;
  border-radius: 20px;
  padding: 8px 12px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.preview-notification-icon {
  font-size: 16px;
}

.preview-notification-badge {
  position: absolute;
  top: -2px;
  right: -2px;
  background: #dc2626;
  color: white;
  border-radius: 50%;
  width: 16px;
  height: 16px;
  font-size: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.preview-menu {
  background: #e5e7eb;
  border-radius: 8px;
  padding: 8px 12px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.preview-menu-icon {
  font-size: 16px;
}

.preview-text {
  font-size: 14px;
  color: #374151;
}

.preview-banner {
  background: #fef3c7;
  border: 1px solid #f59e0b;
  border-radius: 8px;
  padding: 8px 12px;
  width: 100%;
  text-align: center;
}

.preview-banner-text {
  font-size: 12px;
  color: #374151;
  line-height: 1.3;
}

.preview-default {
  font-size: 14px;
  color: #6b7280;
}

.preview-no-layers {
  font-size: 0.9rem;
  color: #9ca3af;
  margin-top: 8px;
}
</style>
