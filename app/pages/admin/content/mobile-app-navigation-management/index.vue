<template>
  <div class="mobile-app-navigation-management-page">
    <div class="header-bg">
      <div class="page-header-flex">
        <h1 class="page-title">مدیریت ناوبری موبایل و اپلیکیشن</h1>
        <button 
          v-if="canCreate" 
          class="btn btn-primary new-item-btn" 
          @click="addNewMobileAppNavigation"
        >
          افزودن ناوبری موبایل و اپلیکیشن
        </button>
      </div>
    </div>

    <!-- Delete Confirm Modal -->
    <DeleteConfirmModal
      ref="deleteModal"
      single-delete-title="تایید حذف ناوبری"
      :single-delete-message="deleteMessage"
      @single-delete="handleDeleteConfirm"
      @close-single="handleDeleteCancel"
    />

    <!-- Toast Messages -->
    <Toast
      v-if="toastMessage"
      :message="toastMessage"
      :type="toastType"
      @close="toastMessage = ''"
    />

    <div class="navigations-section">
      
      <!-- Loading indicator -->
      <div v-if="loading" class="loading-state">
        <p>در حال بارگذاری...</p>
      </div>
      
      <!-- Error state -->
      <div v-else-if="error" class="error-state">
        <p>{{ error }}</p>
        <button class="btn btn-primary" @click="loadMobileAppNavigations">تلاش مجدد</button>
      </div>
      
      <!-- Empty state -->
      <div v-else-if="mobileAppNavigations.length === 0" class="empty-state">
      </div>
      
      <!-- Navigations list -->
      <div v-else class="navigations-list">
        <div v-for="navigation in mobileAppNavigations" :key="navigation.id" class="navigation-item">
          <div class="navigation-info">
            <div class="navigation-title-row">
              <h4>{{ navigation.name }}</h4>
              <span v-if="navigation.is_active" class="status-badge status-active">فعال</span>
              <span v-else class="status-badge status-inactive">غیرفعال</span>
            </div>
            <p v-if="navigation.description" class="navigation-description">{{ navigation.description }}</p>
            <div class="navigation-meta">
              <span class="meta-item">
                <strong>پلتفرم:</strong> {{ getPlatformLabel(navigation.platform) }}
              </span>
              <span class="meta-item">
                <strong>تعداد آیتم‌ها:</strong> {{ navigation.items?.length || 0 }}
              </span>
              <span class="meta-item">
                <strong>تاریخ ایجاد:</strong> {{ formatDate(navigation.created_at || navigation.createdAt) }}
              </span>
            </div>
          </div>
          <div class="navigation-actions">
            <button 
              v-if="canView" 
              class="btn btn-secondary btn-sm" 
              title="پیش‌نمایش"
              @click="previewMobileAppNavigation(navigation)"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
              </svg>
            </button>
            <button 
              v-if="canEdit" 
              class="btn btn-primary btn-sm" 
              title="ویرایش"
              @click="editMobileAppNavigation(navigation)"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
              </svg>
            </button>
            <button 
              v-if="canDelete" 
              class="btn btn-danger btn-sm" 
              title="حذف"
              @click="deleteMobileAppNavigation(navigation)"
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
          <h3>پیش‌نمایش ناوبری موبایل و اپلیکیشن</h3>
          <button class="btn btn-secondary" @click="closePreview">بستن</button>
        </div>
        <div class="modal-body">
          <div class="preview-container">
            <div class="mobile-preview-device">
              <div class="mobile-screen">
                <!-- پیش‌نمایش واقعی بر اساس ساختار آیتم‌ها -->
                <div v-if="previewNavigation?.items && previewNavigation.items.length > 0" class="preview-navigation-container">
                  <div 
                    v-for="(item, itemIndex) in previewNavigation.items" 
                    :key="itemIndex"
                    class="preview-navigation-item"
                    :style="getNavigationItemStyle(item)"
                  >
                    <span v-if="item.type === 'home'" class="preview-nav-icon">🏠</span>
                    <span v-else-if="item.type === 'category'" class="preview-nav-icon">📂</span>
                    <span v-else-if="item.type === 'search'" class="preview-nav-icon">🔍</span>
                    <span v-else-if="item.type === 'cart'" class="preview-nav-icon">🛒</span>
                    <span v-else-if="item.type === 'profile'" class="preview-nav-icon">👤</span>
                    <span v-else-if="item.type === 'menu'" class="preview-nav-icon">☰</span>
                    <span v-else-if="item.type === 'favorite'" class="preview-nav-icon">❤️</span>
                    <span v-else-if="item.type === 'orders'" class="preview-nav-icon">📦</span>
                    <span v-else-if="item.type === 'support'" class="preview-nav-icon">💬</span>
                    <span v-else class="preview-nav-icon">📱</span>
                    
                    <span class="preview-nav-label">{{ item.label || item.name || 'آیتم' }}</span>
                    
                    <span v-if="item.badge" class="preview-nav-badge">{{ item.badge }}</span>
                  </div>
                </div>

                <!-- پیش‌نمایش پیش‌فرض در صورت عدم وجود آیتم -->
                <div v-else class="preview-placeholder">
                  <p>ناوبری موبایل و اپلیکیشن: {{ previewNavigation?.name }}</p>
                  <p class="preview-no-items">هیچ آیتمی تعریف نشده است</p>
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
import { computed, onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import DeleteConfirmModal from '~/components/common/DeleteConfirmModal.vue'
import Toast from '~/components/common/Toast.vue'

import { useAuth } from '~/composables/useAuth'

// تعریف definePageMeta برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void

definePageMeta({
  layout: 'admin-main'
})

const router = useRouter()

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { hasPermission } = useAuth()

// بررسی پرمیژن‌های مورد نیاز
const canView = computed(() => hasPermission('mobile_app_navigation'))
const canCreate = computed(() => hasPermission('mobile_app_navigation'))
const canEdit = computed(() => hasPermission('mobile_app_navigation'))
const canDelete = computed(() => hasPermission('mobile_app_navigation'))

// State
const mobileAppNavigations = ref([])
const loading = ref(false)
const error = ref('')
const showPreview = ref(false)
const previewNavigation = ref(null)

// Delete modal refs
const deleteModal = ref()
const deleteMessage = ref('')
const itemToDelete = ref(null)

// Toast refs
const toastMessage = ref('')
const toastType = ref<'success' | 'error' | 'warning' | 'info'>('info')

// Methods
const loadMobileAppNavigations = async () => {
  loading.value = true
  error.value = ''
  
  try {
    const response = await fetch('/api/admin/mobile-app-navigation-settings', {
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json'
      }
    })
    
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }
    
    const json = await response.json()

    // بررسی ساختار داده‌ها
    const data = json?.data?.data || []


    mobileAppNavigations.value = data
  } catch (err: unknown) {
    console.error('Full error:', err)
    const message = (err as { message?: string })?.message || 'خطا در بارگذاری ناوبری‌های موبایل و اپلیکیشن'
    error.value = message
  } finally {
    loading.value = false
  }
}

const addNewMobileAppNavigation = () => {
  router.push('/admin/content/mobile-app-navigation-management/create')
}

const editMobileAppNavigation = (navigation: { id: number | string }) => {
  router.push(`/admin/content/mobile-app-navigation-management/edit/${navigation.id}`)
}

const previewMobileAppNavigation = (navigation: Record<string, unknown>) => {
  previewNavigation.value = navigation
  showPreview.value = true
}

const closePreview = () => {
  showPreview.value = false
  previewNavigation.value = null
}

const deleteMobileAppNavigation = async (navigation: { id: number | string; name?: string }) => {

  itemToDelete.value = navigation
  // استفاده از فیلدهای مختلف برای نام
    const navigationName = navigation.name || (navigation as { title?: string }).title || (navigation as { label?: string }).label || 'این ناوبری'
  deleteMessage.value = `آیا مطمئن هستید که می‌خواهید ناوبری "${navigationName}" را حذف کنید؟`


  if (deleteModal.value) {

    deleteModal.value.openDeleteConfirm(navigation.id)
  } else {
    console.error('کامپوننت DeleteConfirmModal موجود نیست!')
  }
}

const handleDeleteConfirm = async (id: number | string) => {
  if (!itemToDelete.value) return


  try {
    const response = await fetch(`/api/admin/mobile-app-navigation-settings/${id}`, {
      method: 'DELETE',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json'
      }
    })

    if (!response.ok) {
      const errorText = await response.text()
      console.error('خطا در حذف:', errorText)
      throw new Error(`HTTP error! status: ${response.status}`)
    }
    
    const json = await response.json()

    if (json?.success) {
      // حذف از لیست محلی
      const index = mobileAppNavigations.value.findIndex((n: { id: number | string }) => n.id === id)
      if (index > -1) {
        mobileAppNavigations.value.splice(index, 1)
      }
      // نمایش پیام موفقیت با Toast
      showToast('ناوبری موبایل و اپلیکیشن با موفقیت حذف شد', 'success')
    } else {
      showToast(json?.message || 'خطا در حذف ناوبری موبایل و اپلیکیشن', 'error')
    }
  } catch (err: unknown) {
    console.error('خطا در حذف:', err)
    const message = (err as { message?: string })?.message || 'خطا در حذف ناوبری موبایل و اپلیکیشن'
    showToast(message, 'error')
  } finally {
    itemToDelete.value = null
  }
}

const handleDeleteCancel = () => {
  itemToDelete.value = null
}

// Toast methods
const showToast = (message: string, type: 'success' | 'error' | 'warning' | 'info' = 'info') => {
  toastMessage.value = message
  toastType.value = type
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

const getNavigationItemStyle = (item: Record<string, unknown>) => {
  return {
    display: 'flex',
    alignItems: 'center',
    gap: '8px',
    padding: '12px 16px',
    backgroundColor: String(item.backgroundColor || '#f8f9fa'),
    borderRadius: '8px',
    marginBottom: '8px',
    fontSize: String(item.fontSize || '14px'),
    color: String(item.color || '#333333'),
    fontWeight: String(item.fontWeight || 'normal')
  } as Record<string, string>
}

// Lifecycle
onMounted(() => {
  loadMobileAppNavigations()
})
</script>

<style scoped>
.mobile-app-navigation-management-page {
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

.navigations-section h3 {
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

.navigations-list {
  display: grid;
  gap: 20px;
}

.navigation-item {
  background: white;
  border: 1px solid #e5e7eb;
  border-radius: 12px;
  padding: 24px;
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  transition: all 0.3s ease;
}

.navigation-item:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  border-color: #d1d5db;
}

.navigation-info {
  flex: 1;
}

.navigation-title-row {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 8px;
}

.navigation-title-row h4 {
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

.navigation-description {
  color: #6b7280;
  margin: 8px 0 16px 0;
  line-height: 1.5;
}

.navigation-meta {
  display: flex;
  gap: 24px;
  flex-wrap: wrap;
}

.meta-item {
  font-size: 0.875rem;
  color: #6b7280;
}

.navigation-actions {
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

.preview-navigation-container {
  width: 100%;
  padding-top: 20px;
}

.preview-navigation-item {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 12px 16px;
  background: #f8f9fa;
  border-radius: 8px;
  margin-bottom: 8px;
  border: 1px solid #e5e7eb;
  transition: all 0.2s ease;
}

.preview-navigation-item:hover {
  background: #e9ecef;
  border-color: #d1d5db;
}

.preview-nav-icon {
  font-size: 18px;
  width: 24px;
  text-align: center;
}

.preview-nav-label {
  flex: 1;
  font-size: 14px;
  color: #374151;
  font-weight: 500;
}

.preview-nav-badge {
  background: #dc2626;
  color: white;
  border-radius: 50%;
  width: 20px;
  height: 20px;
  font-size: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
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

.preview-no-items {
  font-size: 0.9rem;
  color: #9ca3af;
  margin-top: 8px;
}
</style>

