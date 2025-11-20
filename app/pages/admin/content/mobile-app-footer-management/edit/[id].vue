<template>
  <div class="edit-mobile-app-footer-page">
    <div class="header-bg">
      <div class="page-header-flex">
        <div>
          <h1 class="page-title">ویرایش فوتر موبایل و اپلیکیشن</h1>
          <p class="page-subtitle">ویرایش فوتر موبایل و اپلیکیشن: {{ mobileAppFooter?.name }}</p>
        </div>
        <button class="btn btn-secondary" @click="goBack">بازگشت</button>
      </div>
    </div>

    <div v-if="loading" class="loading-state">
      <p>در حال بارگذاری...</p>
    </div>

    <div v-else-if="error" class="error-state">
      <p>{{ error }}</p>
      <button class="btn btn-primary" @click="loadMobileAppFooter">تلاش مجدد</button>
    </div>

    <div v-else class="form-container">
      <form class="mobile-app-footer-form" @submit.prevent="updateMobileAppFooter">
        <!-- اطلاعات اصلی -->
        <div class="form-section">
          <h3>اطلاعات اصلی</h3>
          <div class="form-grid">
            <div class="form-group">
              <label for="name">نام فوتر *</label>
              <input 
                id="name"
                v-model="formData.name" 
                type="text" 
                required
                placeholder="نام فوتر موبایل و اپلیکیشن"
                class="form-input"
              >
            </div>
            <div class="form-group">
              <label for="platform">پلتفرم *</label>
              <select id="platform" v-model="formData.platform" required class="form-select">
                <option value="mobile">موبایل</option>
                <option value="app">اپلیکیشن</option>
                <option value="both">هر دو</option>
              </select>
            </div>
          </div>
          <div class="form-group">
            <label for="description">توضیحات</label>
            <textarea 
              id="description"
              v-model="formData.description" 
              rows="3"
              placeholder="توضیحات فوتر موبایل و اپلیکیشن"
              class="form-textarea"
            ></textarea>
          </div>
        </div>

        <!-- تنظیمات نمایش -->
        <div class="form-section">
          <h3>تنظیمات نمایش</h3>
          <div class="form-group">
            <label for="pageSelection">انتخاب صفحات</label>
            <select id="pageSelection" v-model="formData.pageSelection" class="form-select">
              <option value="all">همه صفحات</option>
              <option value="specific">صفحات خاص</option>
              <option value="exclude">مستثنی کردن صفحات</option>
            </select>
          </div>
          <div v-if="formData.pageSelection === 'specific'" class="form-group">
            <label for="specificPages">صفحات خاص</label>
            <textarea 
              id="specificPages"
              v-model="formData.specificPages" 
              rows="2"
              placeholder="آدرس صفحات مورد نظر (هر خط یک آدرس)"
              class="form-textarea"
            ></textarea>
          </div>
          <div v-if="formData.pageSelection === 'exclude'" class="form-group">
            <label for="excludedPages">صفحات مستثنی</label>
            <textarea 
              id="excludedPages"
              v-model="formData.excludedPages" 
              rows="2"
              placeholder="آدرس صفحات مستثنی (هر خط یک آدرس)"
              class="form-textarea"
            ></textarea>
          </div>
          <div class="form-group">
            <label class="checkbox-label">
              <input v-model="formData.isActive" type="checkbox">
              <span>فعال</span>
            </label>
          </div>
        </div>

        <!-- دکمه‌های عملیات -->
        <div class="form-actions">
          <button type="button" class="btn btn-secondary" @click="goBack">لغو</button>
          <button type="submit" class="btn btn-primary" :disabled="updating">
            <span v-if="updating">در حال به‌روزرسانی...</span>
            <span v-else>به‌روزرسانی فوتر موبایل و اپلیکیشن</span>
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';

// تعریف definePageMeta برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void

definePageMeta({
  layout: 'admin-main'
})

const router = useRouter()
const route = useRoute()

// State
const loading = ref(false)
const updating = ref(false)
const error = ref('')
const mobileAppFooter = ref(null)
const formData = ref({
  name: '',
  description: '',
  platform: 'mobile',
  pageSelection: 'all',
  specificPages: '',
  excludedPages: '',
  isActive: true
})

// Methods
const goBack = () => {
  router.push('/admin/content/mobile-app-footer-management')
}

const loadMobileAppFooter = async () => {
  loading.value = true
  error.value = ''
  
  try {
    const response = await $fetch<{ data?: { name?: string; description?: string; platform?: string; pageSelection?: string; specificPages?: string; excludedPages?: string; isActive?: boolean } }>(`/api/admin/mobile-app-footer-settings/${route.params.id}`)
    const data = response.data
    if (data) {
      mobileAppFooter.value = data
      formData.value = {
        name: data.name || '',
        description: data.description || '',
        platform: data.platform || 'mobile',
        pageSelection: data.pageSelection || 'all',
        specificPages: data.specificPages || '',
        excludedPages: data.excludedPages || '',
        isActive: data.isActive !== false
      }
    }
  } catch (err: any) {
    error.value = err.data?.message || 'خطا در بارگذاری فوتر موبایل و اپلیکیشن'
    console.error('Error loading mobile app footer:', err)
  } finally {
    loading.value = false
  }
}

const updateMobileAppFooter = async () => {
  updating.value = true
  
  try {
    const response = await $fetch<{ data?: unknown }>(`/api/admin/mobile-app-footer-settings/${route.params.id}`, {
      method: 'PUT',
      body: formData.value
    })
    const data = response.data
    
    alert('فوتر موبایل و اپلیکیشن با موفقیت به‌روزرسانی شد')
    router.push('/admin/content/mobile-app-footer-management')
  } catch (err: any) {
    alert(err.data?.message || 'خطا در به‌روزرسانی فوتر موبایل و اپلیکیشن')
    console.error('Error updating mobile app footer:', err)
  } finally {
    updating.value = false
  }
}

// Lifecycle
onMounted(() => {
  loadMobileAppFooter()
})
</script>

<style scoped>
.edit-mobile-app-footer-page {
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
  margin: 0 0 8px 0;
}

.page-subtitle {
  margin: 0;
  opacity: 0.9;
  font-size: 1rem;
}

.loading-state, .error-state {
  text-align: center;
  padding: 40px;
  color: #6b7280;
}

.error-state {
  color: #dc2626;
}

.form-container {
  background: white;
  border-radius: 12px;
  padding: 30px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

.mobile-app-footer-form {
  max-width: 800px;
}

.form-section {
  margin-bottom: 40px;
}

.form-section h3 {
  font-size: 1.25rem;
  color: #374151;
  margin-bottom: 20px;
  padding-bottom: 10px;
  border-bottom: 2px solid #e5e7eb;
}

.form-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 20px;
}

.form-group {
  margin-bottom: 20px;
}

.form-group label {
  display: block;
  margin-bottom: 6px;
  font-weight: 600;
  color: #374151;
}

.form-input, .form-select, .form-textarea {
  width: 100%;
  padding: 12px 16px;
  border: 1px solid #d1d5db;
  border-radius: 8px;
  font-size: 1rem;
  transition: border-color 0.2s ease;
}

.form-input:focus, .form-select:focus, .form-textarea:focus {
  outline: none;
  border-color: #3b82f6;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
}

.form-textarea {
  resize: vertical;
  min-height: 80px;
}

.checkbox-label {
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
}

.checkbox-label input[type="checkbox"] {
  width: auto;
  margin: 0;
}

.form-actions {
  display: flex;
  gap: 16px;
  justify-content: flex-end;
  padding-top: 20px;
  border-top: 1px solid #e5e7eb;
}

.btn {
  padding: 12px 24px;
  border-radius: 8px;
  font-weight: 600;
  transition: all 0.2s ease;
  border: none;
  cursor: pointer;
  font-size: 1rem;
}

.btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.btn-primary {
  background: #3b82f6;
  color: white;
}

.btn-primary:hover:not(:disabled) {
  background: #2563eb;
}

.btn-secondary {
  background: #6b7280;
  color: white;
}

.btn-secondary:hover {
  background: #4b5563;
}

@media (max-width: 768px) {
  .form-grid {
    grid-template-columns: 1fr;
  }
  
  .page-header-flex {
    flex-direction: column;
    gap: 20px;
    text-align: center;
  }
  
  .form-actions {
    flex-direction: column;
  }
}
</style>
