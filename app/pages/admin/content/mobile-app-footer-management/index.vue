<template>
  <div class="mobile-app-footer-management-page">
    <div class="header-bg">
      <div class="page-header-flex">
        <h1 class="page-title">مدیریت فوتر موبایل و اپلیکیشن</h1>
        <button class="btn btn-primary new-item-btn" @click="addNewMobileAppFooter">افزودن فوتر موبایل و اپلیکیشن</button>
      </div>
    </div>

    <div class="footers-section">
      
      <!-- Loading indicator -->
      <div v-if="loading" class="loading-state">
        <p>در حال بارگذاری...</p>
      </div>
      
      <!-- Error state -->
      <div v-else-if="error" class="error-state">
        <p>{{ error }}</p>
        <button class="btn btn-primary" @click="loadMobileAppFooters">تلاش مجدد</button>
      </div>
      
      <!-- Empty state -->
      <div v-else-if="mobileAppFooters.length === 0" class="empty-state">
      </div>
      
      <!-- Footers list -->
      <div v-else class="footers-list">
        <div v-for="footer in mobileAppFooters" :key="footer.id" class="footer-item">
          <div class="footer-info">
            <div class="footer-title-row">
              <h4>{{ footer.name }}</h4>
              <span v-if="footer.is_active" class="status-badge status-active">فعال</span>
              <span v-else class="status-badge status-inactive">غیرفعال</span>
            </div>
            <p v-if="footer.description" class="footer-description">{{ footer.description }}</p>
            <div class="footer-meta">
              <span class="meta-item">
                <strong>پلتفرم:</strong> {{ getPlatformLabel(footer.platform) }}
              </span>
              <span class="meta-item">
                <strong>تعداد لایه‌ها:</strong> {{ footer.layers?.length || 0 }}
              </span>
              <span class="meta-item">
                <strong>تاریخ ایجاد:</strong> {{ formatDate(footer.created_at || footer.createdAt) }}
              </span>
            </div>
          </div>
          <div class="footer-actions">
            <button 
              class="btn btn-secondary btn-sm" 
              @click="previewMobileAppFooter(footer)"
              title="پیش‌نمایش"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
              </svg>
            </button>
            <button 
              class="btn btn-primary btn-sm" 
              @click="editMobileAppFooter(footer)"
              title="ویرایش"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
              </svg>
            </button>
            <button 
              class="btn btn-danger btn-sm" 
              @click="deleteMobileAppFooter(footer)"
              title="حذف"
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
          <h3>پیش‌نمایش فوتر موبایل و اپلیکیشن</h3>
          <button class="btn btn-secondary" @click="closePreview">بستن</button>
        </div>
        <div class="modal-body">
          <div class="preview-container">
            <!-- اینجا باید کامپوننت پیش‌نمایش فوتر موبایل و اپلیکیشن قرار گیرد -->
            <p>پیش‌نمایش فوتر موبایل و اپلیکیشن: {{ previewFooter?.name }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';

// تعریف definePageMeta برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void

definePageMeta({
  layout: 'admin-main'
})

const router = useRouter()

// State
const mobileAppFooters = ref([])
const loading = ref(false)
const error = ref('')
const showPreview = ref(false)
const previewFooter = ref(null)

// Methods
const loadMobileAppFooters = async () => {
  loading.value = true
  error.value = ''
  
  try {
    const response = await $fetch<{ data?: unknown[] }>('/api/admin/mobile-app-footer-settings')
    const data = response.data
    mobileAppFooters.value = data || []
  } catch (err: any) {
    error.value = err.data?.message || 'خطا در بارگذاری فوترهای موبایل و اپلیکیشن'
    console.error('Error loading mobile app footers:', err)
  } finally {
    loading.value = false
  }
}

const addNewMobileAppFooter = () => {
  router.push('/admin/content/mobile-app-footer-management/create')
}

const editMobileAppFooter = (footer: any) => {
  router.push(`/admin/content/mobile-app-footer-management/edit/${footer.id}`)
}

const previewMobileAppFooter = (footer: any) => {
  previewFooter.value = footer
  showPreview.value = true
}

const closePreview = () => {
  showPreview.value = false
  previewFooter.value = null
}

const deleteMobileAppFooter = async (footer: any) => {
  if (!confirm(`آیا مطمئن هستید که می‌خواهید فوتر "${footer.name}" را حذف کنید؟`)) {
    return
  }
  
  try {
    await $fetch(`/api/admin/mobile-app-footer-settings/${footer.id}`, {
      method: 'DELETE'
    })
    
    // حذف از لیست محلی
    const index = mobileAppFooters.value.findIndex((f: any) => f.id === footer.id)
    if (index > -1) {
      mobileAppFooters.value.splice(index, 1)
    }
    
    alert('فوتر موبایل و اپلیکیشن با موفقیت حذف شد')
  } catch (err: any) {
    alert(err.data?.message || 'خطا در حذف فوتر موبایل و اپلیکیشن')
    console.error('Error deleting mobile app footer:', err)
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

// Lifecycle
onMounted(() => {
  loadMobileAppFooters()
})
</script>

<style scoped>
.mobile-app-footer-management-page {
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

.footers-section h3 {
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

.footers-list {
  display: grid;
  gap: 20px;
}

.footer-item {
  background: white;
  border: 1px solid #e5e7eb;
  border-radius: 12px;
  padding: 24px;
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  transition: all 0.3s ease;
}

.footer-item:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  border-color: #d1d5db;
}

.footer-info {
  flex: 1;
}

.footer-title-row {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 8px;
}

.footer-title-row h4 {
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

.footer-description {
  color: #6b7280;
  margin: 8px 0 16px 0;
  line-height: 1.5;
}

.footer-meta {
  display: flex;
  gap: 24px;
  flex-wrap: wrap;
}

.meta-item {
  font-size: 0.875rem;
  color: #6b7280;
}

.footer-actions {
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
  min-height: 200px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #6b7280;
}
</style>
