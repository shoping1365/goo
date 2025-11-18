<template>
  <div class="header-management-page">
    <div class="header-bg">
      <div class="page-header-flex">
        <h1 class="page-title">مدیریت هدرها</h1>
        <button class="btn btn-primary new-item-btn" @click="addNewHeader">افزودن هدر</button>
      </div>
    </div>

    <div class="headers-section">
      <h3>هدرها</h3>
      
      <!-- Loading indicator -->
      <div v-if="loading" class="loading-state">
        <p>در حال بارگذاری...</p>
      </div>
      
      <!-- Error state -->
      <div v-else-if="error" class="error-state">
        <p>{{ error }}</p>
        <button class="btn btn-primary" @click="loadHeaders">تلاش مجدد</button>
      </div>
      
      <!-- Empty state -->
      <div v-else-if="headers.length === 0" class="empty-state">
        <p>هیچ هدری وجود ندارد</p>
        <button class="btn btn-primary" @click="addNewHeader">افزودن اولین هدر</button>
      </div>
      
      <!-- Headers list -->
      <div v-else class="headers-list">
        <div v-for="header in headers" :key="header.id" class="header-item">
          <div class="header-info">
            <div class="header-title-row">
              <h4>{{ header.name }}</h4>
              <span v-if="header.is_active" class="status-badge status-active">فعال</span>
              <span v-else class="status-badge status-inactive">غیرفعال</span>
            </div>
            <p v-if="header.description" class="header-description">{{ header.description }}</p>
            <div class="header-meta">
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
              @click="previewHeader(header)"
              title="پیش‌نمایش"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
              </svg>
            </button>
            <button 
              class="btn btn-primary btn-sm" 
              @click="editHeader(header.id)"
              title="ویرایش"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
              </svg>
            </button>
            <button 
              class="btn btn-danger btn-sm" 
              @click="handleDeleteHeader(header.id)"
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
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';
import { useHeaders } from '~/composables/useHeaders';

// تعریف definePageMeta و navigateTo برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
declare const navigateTo: (to: string) => Promise<void>

definePageMeta({
  layout: 'admin-main'
})

const router = useRouter()

const { headers, loading, error, loadHeaders, deleteHeader, toggleHeaderStatus } = useHeaders()
const toggleLoading = ref<number | null>(null)

// تابع فرمت کردن تاریخ
function formatDate(dateString: string | undefined): string {
  if (!dateString) return 'نامشخص'
  
  try {
    const date = new Date(dateString)
    if (isNaN(date.getTime())) return 'نامشخص'
    
    return date.toLocaleDateString('fa-IR', {
      year: 'numeric',
      month: 'long',
      day: 'numeric'
    })
  } catch (error) {
    console.error('خطا در فرمت کردن تاریخ:', error)
    return 'نامشخص'
  }
}

// پیش‌نمایش هدر
function previewHeader(header: any) {
  // اینجا می‌توانید modal پیش‌نمایش اضافه کنید
  console.log('پیش‌نمایش هدر:', header)
}

// ویرایش هدر
function editHeader(headerId: number) {
  // هدایت به صفحه ویرایش (فعلاً همان صفحه ایجاد)
  router.push(`/admin/content/header-management/create?id=${headerId}`)
}

function addNewHeader() {
  // هدایت به صفحه جدید برای افزودن هدر
  router.push('/admin/content/header-management/create')
}

// فعال/غیرفعال کردن هدر
async function handleToggleStatus(headerId: number | undefined) {
  if (!headerId) {
    alert('شناسه هدر نامعتبر است')
    return
  }
  
  try {
    toggleLoading.value = headerId
    const success = await toggleHeaderStatus(headerId)
    if (success) {
      // پیام موفقیت نمایش نمی‌دهیم چون وضعیت به صورت بصری تغییر می‌کند
    } else {
      alert('خطا در تغییر وضعیت هدر')
    }
  } catch (error) {
    console.error('خطا در تغییر وضعیت هدر:', error)
    alert('خطا در تغییر وضعیت هدر')
  } finally {
    toggleLoading.value = null
  }
}

// حذف هدر
async function handleDeleteHeader(headerId: number | undefined) {
  if (!headerId) {
    alert('شناسه هدر نامعتبر است')
    return
  }
  
  if (confirm('آیا مطمئن هستید که می‌خواهید این هدر را حذف کنید؟')) {
    try {
      const success = await deleteHeader(headerId)
      if (success) {
        alert('هدر با موفقیت حذف شد')
      } else {
        alert('خطا در حذف هدر')
      }
    } catch (error) {
      console.error('خطا در حذف هدر:', error)
      alert('خطا در حذف هدر')
    }
  }
}

// بارگذاری هدرها هنگام بارگذاری صفحه
onMounted(() => {
  loadHeaders()
})
</script>

<style scoped>
.header-management-page {
  padding: 24px;
  max-width: 1200px;
  margin: 0 auto;
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
  align-items: center;
  justify-content: space-between;
}

.page-title {
  font-size: 2rem;
  font-weight: bold;
  margin: 0;
}

.btn {
  padding: 10px 20px;
  border: none;
  border-radius: 8px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
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

.headers-section {
  background: white;
  border-radius: 12px;
  padding: 30px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
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
</style> 
