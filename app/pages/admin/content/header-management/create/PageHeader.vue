<template>
  <div class="header-bg" :class="{ 'full-width': isEditing }">
    <div class="page-header-flex">
      <h1 class="page-title">{{ isEditing ? 'ویرایش هدر' : 'ایجاد هدر جدید' }}</h1>
      <div class="header-actions">
        <button class="btn btn-primary" @click="saveHeader">{{ isEditing ? 'به‌روزرسانی' : 'ذخیره' }}</button>
        <button class="btn btn-secondary" @click="goBack">برگشت</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { inject } from 'vue'

// Inject functions from parent
const saveHeader = inject<(() => void) | undefined>('saveHeader')
const goBack = inject<(() => void) | undefined>('goBack')
const isEditing = inject<boolean | undefined>('isEditing')
</script>

<style scoped>
.header-bg {
  background: #fff;
  border-radius: 10px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.04);
  padding: 18px;
  margin-bottom: 24px;
  width: 100%; /* اطمینان از تمام عرض بودن */
}

/* حالت تمام عرض برای ویرایش */
.header-bg.full-width {
  width: 100%;
  max-width: none;
}

.page-header-flex {
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 100%; /* اطمینان از تمام عرض بودن */
}

.header-actions {
  display: flex;
  gap: 12px;
}

.page-title {
  font-size: 22px;
  font-weight: 700;
  color: #2c3e50;
  margin: 0;
  flex: 1; /* اجازه دادن به عنوان برای اشغال فضای بیشتر */
}

.btn {
  padding: 10px 20px;
  border: none;
  border-radius: 8px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
}

.btn-primary {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}

.btn-primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3);
}

.btn-secondary {
  background: #e9ecef;
  color: #6c757d;
  border: 1px solid #ced4da;
}

.btn-secondary:hover {
  background: #f8f9fa;
  color: #495057;
}
</style> 
