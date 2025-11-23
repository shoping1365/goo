<template>
  <div v-if="hasAccess" class="header-settings-form">
    <div class="form-header">
      <div class="header-title-row">
        <h3>تنظیمات هدر</h3>
        <div class="header-toggle">
          <label class="toggle-label">
            <input v-model="headerData.isActive" type="checkbox" />
            <span class="toggle-text">فعال بودن هدر</span>
          </label>
        </div>
      </div>
    </div>
    
    <div class="form-content">
      <div class="form-group">
        <div class="name-pages-row">
          <div class="name-field">
            <label>نام هدر <span class="required">*</span></label>
            <input 
              v-model="headerData.name" 
              type="text" 
              class="form-control name-input" 
              placeholder="نام هدر را وارد کنید" 
              required 
            />
            <small v-if="!headerData.name.trim()" class="error-text">
              نام هدر الزامی است
            </small>
          </div>
          
          <div class="pages-field">
            <div class="form-row">
              <label>انتخاب صفحات</label>
            </div>
            
            <div class="radio-options">
              <label class="radio-option">
                <input v-model="headerData.pageSelection" type="radio" value="all" />
                <span>کل سایت</span>
              </label>
              
              <label class="radio-option">
                <input v-model="headerData.pageSelection" type="radio" value="specific" />
                <span>صفحات خاص</span>
              </label>
              
              <label class="radio-option">
                <input v-model="headerData.pageSelection" type="radio" value="exclude" />
                <span>مستثنی کردن صفحات</span>
              </label>
            </div>
          </div>
        </div>
      </div>
      
      <div v-if="headerData.pageSelection === 'exclude'" class="form-group">
        <label>صفحات مستثنی</label>
        <textarea 
          v-model="headerData.excludedPages" 
          class="form-control exclude-textarea" 
          placeholder="آدرس صفحات مستثنی را وارد کنید (هر صفحه در یک خط)" 
          rows="2"
        ></textarea>
        <small class="help-text">مثال: 
          http://localhost:3000/admin<br>

        </small>
      </div>
      
      <div v-if="headerData.pageSelection === 'specific'" class="form-group">
        <label>صفحات خاص</label>
        <textarea 
          v-model="headerData.specificPages" 
          class="form-control" 
          placeholder="آدرس صفحات را وارد کنید (هر صفحه در یک خط)" 
          rows="2"
        ></textarea>
        <small class="help-text">مثال: 
          http://localhost:3000/about<br>

        </small>
      </div>

    </div>
  </div>
</template>

<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
</script>

<script setup lang="ts">
import { computed, inject, onMounted, ref, type Ref, watch } from 'vue';
import { useAuth } from '~/composables/useAuth';

// احراز هویت
const { user, isAuthenticated } = useAuth();

// بررسی دسترسی admin
const hasAccess = computed(() => {
  if (!isAuthenticated.value) {
    return false;
  }

  const userRole = user.value?.role?.toLowerCase() || '';
  const adminRoles = ['admin', 'developer'];
  return adminRoles.includes(userRole);
});

// بررسی احراز هویت و دسترسی admin - نمایش 404 در صورت عدم دسترسی
const checkAuth = async (): Promise<void> => {
  if (!hasAccess.value) {
    await navigateTo('/404', { external: false });
  }
};

// بررسی احراز هویت در هنگام mount
onMounted(async () => {
  await checkAuth();
});

// بررسی احراز هویت هنگام تغییر وضعیت احراز هویت
watch([isAuthenticated, hasAccess], async () => {
  if (!hasAccess.value) {
    await checkAuth();
  }
});

interface HeaderData {
  name?: string
  description?: string
  pageSelection?: string
  specificPages?: string
  excludedPages?: string
  isActive?: boolean
  [key: string]: unknown
}

const headerData = inject<Ref<HeaderData>>('headerData', ref({
  name: '',
  description: '',
  pageSelection: 'all',
  specificPages: '',
  excludedPages: '',
  isActive: true
}))





</script>

<style scoped>
.header-settings-form {
  background: white;
  border-radius: 6px;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
  overflow: hidden;
}

.form-header {
  background: #f8fafc;
  padding: 12px 16px;
  border-bottom: 1px solid #e2e8f0;
}

.header-title-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.form-header h3 {
  margin: 0;
  font-size: 16px;
  font-weight: 600;
  color: #1e293b;
}

.header-toggle {
  display: flex;
  align-items: center;
}

.toggle-label {
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
  margin: 0;
}

.toggle-label input[type="checkbox"] {
  width: 16px;
  height: 16px;
  margin: 0;
  cursor: pointer;
  accent-color: #3b82f6;
}

.toggle-text {
  font-size: 13px;
  color: #374151;
  cursor: pointer;
  font-weight: 500;
}

.form-content {
  padding: 16px;
}

.form-group {
  margin-bottom: 16px;
}

.form-group:last-child {
  margin-bottom: 0;
}

.name-pages-row {
  display: flex;
  flex-direction: row;
  gap: 0;
  align-items: flex-start;
  justify-content: space-between;
}

.name-field {
  flex: 0 0 220px;
  order: 1;
}

.pages-field {
  flex: 0 0 300px;
  margin-right: 0 !important;
  margin-left: 50px !important;
  padding-right: 0 !important;
  align-self: stretch;
  order: 2;
}

.name-input {
  width: 100%;
  height: 36px;
  line-height: 1.2;
  padding-top: 8px;
  padding-bottom: 8px;
}

.exclude-textarea {
  width: 100%;
  height: 36px;
  min-height: 36px;
  resize: vertical;
  line-height: 1.2;
  padding-top: 8px;
  padding-bottom: 8px;
}

.form-row {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 8px;
}

label {
  display: block;
  margin-bottom: 6px;
  font-size: 13px;
  font-weight: 500;
  color: #374151;
}

.required {
  color: #ef4444;
  margin-right: 4px;
}

.form-control {
  width: 100%;
  padding: 8px 10px;
  border: 1px solid #d1d5db;
  border-radius: 4px;
  font-size: 13px;
  color: #374151;
  background: white;
  transition: border-color 0.2s ease;
}

.form-control:focus {
  outline: none;
  border-color: #3b82f6;
  box-shadow: 0 0 0 2px rgba(59, 130, 246, 0.1);
}

.form-control::placeholder {
  color: #9ca3af;
}

textarea.form-control {
  resize: vertical;
  min-height: 60px;
}

.radio-options {
  display: flex;
  flex-direction: row;
  gap: 20px;
  margin-top: 8px;
  flex-wrap: nowrap;
  align-items: center;
}

.radio-option {
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
  padding: 4px 0;
  white-space: nowrap;
}

.radio-option input[type="radio"] {
  width: 14px;
  height: 14px;
  margin: 0;
  cursor: pointer;
  accent-color: #3b82f6;
}

.radio-option span {
  font-size: 13px;
  color: #374151;
  cursor: pointer;
  white-space: nowrap;
}

.checkbox-label {
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
  padding: 4px 0;
}

.checkbox-label input[type="checkbox"] {
  width: 14px;
  height: 14px;
  margin: 0;
  cursor: pointer;
  accent-color: #3b82f6;
}

.checkbox-label span {
  font-size: 13px;
  color: #374151;
  cursor: pointer;
}

.help-text {
  display: block;
  margin-top: 4px;
  font-size: 11px;
  color: #6b7280;
  font-style: italic;
}

.error-text {
  display: block;
  margin-top: 4px;
  font-size: 11px;
  color: #ef4444;
}

/* Responsive styles for main form */
@media (max-width: 640px) {
  .header-title-row {
    flex-direction: column;
    gap: 12px;
    align-items: flex-start;
  }
  
  .name-pages-row {
    flex-direction: column;
    gap: 16px;
  }
  
  .name-field {
    flex: none;
  }
}
</style>
