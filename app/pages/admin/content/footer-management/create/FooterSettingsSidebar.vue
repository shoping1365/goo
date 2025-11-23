<template>
  <div v-if="showLayerSettings && hasAccess" class="settings-sidebar">
    <!-- تنظیمات لایه - فقط وقتی لایه جدید در حال ایجاد است نمایش داده می‌شود -->
    <div class="layer-settings-section">
      <div class="settings-header">
        <h3>{{ newLayer.id ? 'ویرایش لایه' : 'افزودن لایه جدید' }}</h3>
        <button class="close-sidebar-btn" title="بستن" @click="cancelLayer">
          <span class="i-heroicons-x-mark"></span>
        </button>
      </div>

      <div class="form-group">
        <label>نام لایه: <span style="color: red;">*</span></label>
        <input v-model="newLayer.name" type="text" class="form-control" placeholder="نام لایه را وارد کنید" required />
        <small v-if="!newLayer.name.trim()" style="color: red; font-size: 12px;">نام لایه الزامی است</small>
      </div>
      
      <!-- خط جداکننده -->
      <div class="section-divider"></div>
      
      <div class="form-group">
        <label>رنگ پس‌زمینه:</label>
        <div class="color-picker-container">
          <input v-model="newLayer.color" type="color" class="form-control color-input" />
          <span class="color-value">{{ newLayer.color }}</span>
        </div>
      </div>
      
      <!-- خط جداکننده -->
      <div class="section-divider"></div>
      
      <div class="form-group">
        <label>شفافیت:</label>
        <div class="opacity-control">
          <input v-model.number="newLayer.opacity" type="range" min="0" max="100" class="form-control opacity-slider" />
          <span class="opacity-value">{{ newLayer.opacity }}%</span>
        </div>
      </div>
      
      <!-- خط جداکننده -->
      <div class="section-divider"></div>
      
      <!-- تنظیمات ریسپانسیو و راست‌چین -->
      <div class="form-group">
        <label>عرض لایه (%):</label>
        <input v-model.number="newLayer.width" type="number" class="form-control" min="10" max="100" step="5" />
        <small class="form-help">عرض لایه به درصد (10 تا 100)</small>
      </div>
      
      <!-- فاصله بیشتر بین عرض و ارتفاع -->
      <div style="margin-bottom: 20px;"></div>
      
      <div class="form-group">
        <label>ارتفاع لایه (px):</label>
        <input v-model.number="newLayer.height" type="number" class="form-control" min="30" max="200" step="5" />
        <small class="form-help">ارتفاع لایه به پیکسل (30 تا 200)</small>
      </div>
      
      <!-- خط جداکننده -->
      <div class="section-divider"></div>
      
      <!-- تنظیمات حاشیه (Border) -->
      <div class="form-group">
        <label class="checkbox-label">
          <input v-model="newLayer.enableBorder" type="checkbox" />
          فعال کردن حاشیه (Border)
        </label>
      </div>
      
      <div v-if="newLayer.enableBorder" class="border-settings">
        <div class="form-group">
          <label>جهت حاشیه:</label>
          <select v-model="newLayer.borderPosition" class="form-control">
            <option value="all">همه طرف‌ها</option>
            <option value="top">بالا</option>
            <option value="bottom">پایین</option>
            <option value="left">چپ</option>
            <option value="right">راست</option>
            <option value="top-bottom">بالا و پایین</option>
          </select>
        </div>
        
        <div class="form-group">
          <label>رنگ حاشیه:</label>
          <div class="color-picker-container">
            <input v-model="newLayer.borderColor" type="color" class="form-control color-input" />
            <span class="color-value">{{ newLayer.borderColor }}</span>
          </div>
        </div>
        
        <div class="form-group">
          <label>ضخامت حاشیه (px):</label>
          <input v-model.number="newLayer.borderWidth" type="number" class="form-control" min="1" max="10" />
        </div>
        
        <div class="form-group">
          <label>نوع حاشیه:</label>
          <select v-model="newLayer.borderStyle" class="form-control">
            <option value="solid">خط ساده</option>
            <option value="dashed">خط نقطه‌چین</option>
            <option value="dotted">خط نقطه‌ای</option>
            <option value="double">خط دوتایی</option>
          </select>
        </div>
      </div>
      
      <!-- خط جداکننده -->
      <div class="section-divider"></div>
      
      <!-- تنظیمات سایه -->
      <div class="form-group">
        <label class="checkbox-label">
          <input v-model="newLayer.enableShadow" type="checkbox" />
          فعال کردن سایه (Elevation)
        </label>
      </div>
      
      <div v-if="newLayer.enableShadow" class="shadow-settings">
        <div class="form-group">
          <label>شدت سایه:</label>
          <select v-model="newLayer.shadowIntensity" class="form-control">
            <option value="none">بدون سایه</option>
            <option value="sm">ملایم</option>
            <option value="md">متوسط</option>
            <option value="lg">قوی</option>
            <option value="xl">خیلی قوی</option>
          </select>
        </div>
        
        <div class="form-group">
          <label>جهت سایه:</label>
          <select v-model="newLayer.shadowDirection" class="form-control">
            <option value="top">بالا</option>
            <option value="bottom">پایین</option>
            <option value="both">هر دو طرف</option>
          </select>
        </div>
      </div>
      
      <!-- خط جداکننده -->
      <div class="section-divider"></div>
      
      <!-- تنظیمات جداکننده -->
      <div class="form-group">
        <label class="checkbox-label">
          <input v-model="newLayer.showSeparator" type="checkbox" />
          نمایش جداکننده
        </label>
      </div>
      
      <div v-if="newLayer.showSeparator" class="separator-settings">
        <div class="form-group">
          <label>نوع جداکننده:</label>
          <select v-model="newLayer.separatorType" class="form-control">
            <option value="solid">خط ساده</option>
            <option value="dashed">خط نقطه‌چین</option>
            <option value="dotted">خط نقطه‌ای</option>
            <option value="double">خط دوتایی</option>
          </select>
        </div>
        
        <div class="form-group">
          <label>رنگ جداکننده:</label>
          <div class="color-picker-container">
            <input v-model="newLayer.separatorColor" type="color" class="form-control color-input" />
            <span class="color-value">{{ newLayer.separatorColor }}</span>
          </div>
        </div>
        
        <div class="form-group">
          <label>شفافیت جداکننده (%):</label>
          <input v-model="newLayer.separatorOpacity" type="number" class="form-control" min="0" max="100" />
        </div>
        
        <div class="form-group">
          <label>ضخامت جداکننده:</label>
          <input v-model="newLayer.separatorWidth" type="number" class="form-control" min="1" max="10" />
        </div>
      </div>
      
      <!-- خط جداکننده -->
      <div class="section-divider"></div>
      
      <!-- دکمه‌های عملیات -->
      <div class="layer-actions">
        <button class="btn btn-primary" @click="saveLayer">
          <span class="i-heroicons-check mr-1"></span>
          <span v-if="newLayer.id">به‌روزرسانی</span>
          <span v-else>ذخیره</span>
        </button>
        <button class="btn btn-secondary" @click="cancelLayer">
          <span class="i-heroicons-x-mark mr-1"></span>
          انصراف
        </button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
</script>

<script setup lang="ts">
import { computed, inject, onMounted, type Ref, watch } from 'vue'
import { useAuth } from '~/composables/useAuth'

// احراز هویت
const { user, isAuthenticated } = useAuth()

// بررسی دسترسی admin
const hasAccess = computed(() => {
  if (!isAuthenticated.value) {
    return false
  }

  const userRole = user.value?.role?.toLowerCase() || ''
  const adminRoles = ['admin', 'developer']
  return adminRoles.includes(userRole)
})

// بررسی احراز هویت و دسترسی admin - نمایش 404 در صورت عدم دسترسی
const checkAuth = async (): Promise<void> => {
  if (!hasAccess.value) {
    await navigateTo('/404', { external: false })
  }
}

// بررسی احراز هویت در هنگام mount
onMounted(async () => {
  await checkAuth()
})

// بررسی احراز هویت هنگام تغییر وضعیت احراز هویت
watch([isAuthenticated, hasAccess], async () => {
  if (!hasAccess.value) {
    await checkAuth()
  }
})

interface LayerData {
  id?: string
  name: string
  width: number
  height: number
  rowCount: number
  color: string
  opacity: number
  enableBorder: boolean
  borderPosition: string
  borderColor: string
  borderWidth: number
  borderStyle: string
  enableShadow: boolean
  shadowIntensity: string
  shadowDirection: string
  showSeparator: boolean
  separatorType: string
  separatorColor: string
  separatorOpacity: number
  separatorWidth: number
  items: unknown[]
}

// Inject functions and data from parent
const newLayer = inject<Ref<LayerData>>('newLayer')!
const showLayerSettings = inject<Ref<boolean>>('showLayerSettings')!
const saveLayer = inject<() => void>('saveLayer')!
const cancelLayer = inject<() => void>('cancelLayer')!
</script>

<style scoped>
.settings-sidebar {
  background: #fff;
  border-radius: 10px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.08);
  padding: 24px;
  height: fit-content;
  max-height: calc(100vh - 100px);
  overflow-y: auto;
  direction: rtl;
  text-align: right;
}

/* Custom scrollbar */
.settings-sidebar::-webkit-scrollbar {
  width: 6px;
}

.settings-sidebar::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 10px;
}

.settings-sidebar::-webkit-scrollbar-thumb {
  background: #888;
  border-radius: 10px;
}

.settings-sidebar::-webkit-scrollbar-thumb:hover {
  background: #555;
}

.settings-sidebar h3 {
  font-size: 18px;
  font-weight: 700;
  color: #1f2937;
  margin: 0;
}

.settings-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
  padding-bottom: 12px;
  border-bottom: 2px solid #e5e7eb;
}

.close-sidebar-btn {
  background: #f3f4f6;
  border: none;
  width: 32px;
  height: 32px;
  border-radius: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.2s;
  color: #6b7280;
  font-size: 20px;
}

.close-sidebar-btn:hover {
  background: #e5e7eb;
  color: #374151;
}

.settings-sidebar .settings-form {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.settings-sidebar .form-group {
  margin-bottom: 0;
}

/* خط‌های جداکننده بین بخش‌ها */
.section-divider {
  height: 1px;
  background: #667eea;
  margin: 20px 0;
  border-radius: 1px;
}

.settings-sidebar .form-control {
  font-size: 13px;
  padding: 10px 12px;
  width: 100%;
  border: 1px solid #ced4da;
  border-radius: 6px;
  transition: border-color 0.2s;
}

.settings-sidebar .form-control:focus {
  outline: none;
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

.settings-sidebar .color-input {
  height: 40px;
  padding: 4px;
}

.color-picker-container {
  display: flex;
  align-items: center;
  gap: 12px;
}

.color-preview {
  width: 40px;
  height: 40px;
  border-radius: 6px;
  border: 2px solid #ced4da;
  cursor: pointer;
  transition: all 0.2s ease;
}

.color-preview:hover {
  border-color: #667eea;
  transform: scale(1.05);
}

.color-value {
  font-size: 12px;
  color: #6c757d;
  font-weight: 500;
  font-family: monospace;
  min-width: 70px;
}

.opacity-control {
  display: flex;
  align-items: center;
  gap: 12px;
}

.opacity-slider {
  flex: 1;
}

.opacity-value {
  font-size: 11px;
  color: #6c757d;
  font-weight: 500;
  min-width: 35px;
  text-align: center;
}

.settings-sidebar .btn {
  padding: 10px 20px;
  border: none;
  border-radius: 8px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
  font-size: 13px;
}

.settings-sidebar .btn-primary {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}

.settings-sidebar .btn-primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3);
}

.checkbox-label {
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
}

.checkbox-label input[type="checkbox"] {
  margin: 0;
  cursor: pointer;
}

.separator-settings {
  border: 1px solid #e9ecef;
  border-radius: 8px;
  padding: 16px;
  margin-top: 16px;
  background-color: #f8f9fa;
}

.shadow-settings {
  border: 1px solid #e9ecef;
  border-radius: 8px;
  padding: 16px;
  margin-top: 16px;
  background-color: #f8f9fa;
}

.border-settings {
  border: 1px solid #e9ecef;
  border-radius: 8px;
  padding: 16px;
  margin-top: 16px;
  background-color: #f8f9fa;
}

.settings-sidebar .btn-secondary {
  background: #e9ecef;
  color: #6c757d;
  border: 1px solid #ced4da;
}

.settings-sidebar .btn-secondary:hover {
  background: #f8f9fa;
  color: #495057;
}

.settings-sidebar label {
  display: block;
  font-weight: 500;
  color: #2c3e50;
  margin-bottom: 10px;
  font-size: 13px;
  text-align: right;
}

/* بهبود استایل‌های separator */
.separator-settings {
  border-top: 1px solid #e9ecef;
  padding-top: 16px;
  margin-top: 16px;
}

/* بهبود استایل‌های textarea */
textarea.form-control {
  resize: vertical;
  min-height: 60px;
  font-family: inherit;
}

/* بهبود استایل‌های number input */
input[type="number"].form-control {
  -moz-appearance: textfield;
  appearance: textfield;
}

input[type="number"].form-control::-webkit-outer-spin-button,
input[type="number"].form-control::-webkit-inner-spin-button {
  -webkit-appearance: none;
  margin: 0;
}

/* استایل‌های جدید */
.form-help {
  font-size: 11px;
  color: #6c757d;
  margin-top: 4px;
  display: block;
}

.layer-actions {
  display: flex;
  flex-direction: column;
  gap: 10px;
  margin-top: 24px;
  padding-top: 20px;
  border-top: 1px solid #e5e7eb;
}

.layer-actions .btn {
  width: 100%;
  padding: 12px 16px;
  border: none;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 6px;
}

.layer-actions .btn-primary {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  box-shadow: 0 2px 8px rgba(102, 126, 234, 0.3);
}

.layer-actions .btn-primary:hover {
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.4);
}

.layer-actions .btn-secondary {
  background: white;
  color: #6b7280;
  border: 1px solid #e5e7eb;
}

.layer-actions .btn-secondary:hover {
  background: #f9fafb;
  border-color: #d1d5db;
}

.layer-settings-section {
  margin-top: 20px;
  padding-top: 20px;
  border-top: 1px solid #e9ecef;
}

.layer-settings-section h3 {
  font-size: 16px;
  font-weight: 600;
  color: #2c3e50;
  margin-bottom: 16px;
}

/* Responsive Design برای سایدبار */
@media (max-width: 768px) {
  .settings-sidebar {
    flex: 0 0 100%;
    position: static;
    margin-bottom: 20px;
    padding: 16px;
  }
  
  .layer-settings-section {
    margin-top: 16px;
    padding-top: 16px;
  }
  
  .form-group {
    margin-bottom: 16px;
  }
  
  .form-control {
    font-size: 14px;
    padding: 12px;
  }
  
  .btn {
    width: 100%;
    margin-bottom: 8px;
  }
  
  .layer-actions {
    flex-direction: column;
    gap: 8px;
  }
  
  /* تنظیم فاصله خط‌های جداکننده در موبایل */
  .section-divider {
    margin: 16px 0;
  }
}

@media (max-width: 480px) {
  .settings-sidebar {
    padding: 12px;
  }
  
  .form-control {
    font-size: 16px; /* برای موبایل بهتر */
    padding: 14px;
  }
  
  .color-picker-container {
    flex-direction: column;
    gap: 8px;
  }
  
  .opacity-control {
    flex-direction: column;
    gap: 8px;
  }
  
  .layer-actions {
    flex-direction: column;
    gap: 8px;
  }
  
  .btn {
    width: 100%;
    font-size: 14px;
    padding: 12px;
  }
  
  /* تنظیم فاصله خط‌های جداکننده در صفحات کوچک */
  .section-divider {
    margin: 12px 0;
  }
}
</style>
