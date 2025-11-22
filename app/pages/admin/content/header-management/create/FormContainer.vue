<template>
  <div class="form-container" :class="{ 'full-width': isEditing }">
    <div class="form-section">
      <div class="form-header">
        <button class="btn btn-primary" @click="handleAddNewLayer">افزودن لایه جدید</button>
        <button v-if="Array.isArray(createdLayers) && createdLayers.length > 0" class="btn btn-danger" @click="confirmClearAllLayers">پاک کردن همه لایه‌ها</button>
      </div>
      
      <!-- نمایش لایه‌های ایجاد شده -->
      <div v-if="Array.isArray(createdLayers) && createdLayers.length > 0" class="created-layers">
        <h4>لایه‌های ایجاد شده</h4>
        <div class="layers-list">
          <div 
            v-for="(layer, idx) in createdLayers" 
            :key="(layer as { id?: number | string }).id"
            class="layer-item"
            draggable="true"
            @dragstart="onDragStart(idx)"
            @dragover.prevent="onDragOver(idx)"
            @drop="onDrop(idx)"
            @dragend="onDragEnd"
          >
            <div class="layer-info">
              <div class="layer-name">{{ ((layer as { name?: string }).name) || 'لایه بدون نام' }}</div>
              <div class="layer-details">
                <span class="detail-item">نام لایه: {{ ((layer as { name?: string }).name) || 'بدون نام' }}</span>
                <span class="detail-item">عرض: {{ ((layer as { width?: number }).width) || 0 }}%</span>
                <span class="detail-item">ارتفاع: {{ ((layer as { height?: number }).height) || 0 }}px</span>
                <span class="detail-item">شفافیت: {{ ((layer as { opacity?: number }).opacity) || 0 }}%</span>
                <span class="detail-item">تعداد ردیف: {{ ((layer as { rowCount?: number }).rowCount) || 0 }}</span>
                <span class="detail-item">ایتم ها: {{ Array.isArray((layer as { items?: unknown[] }).items) ? ((layer as { items?: unknown[] }).items?.length || 0) : 0 }} مورد انتخاب شده</span>
              </div>
            </div>
            <div class="layer-actions">
              <button class="btn btn-sm btn-outline" @click="editLayer(layer as HeaderLayer)">ویرایش</button>
              <button class="btn btn-sm btn-danger" @click="deleteLayer((layer as { id?: number | string }).id)">حذف</button>
            </div>
          </div>
        </div>
      </div>
      
      <!-- کامپوننت هشدار -->
      <div v-if="showWarningModal" class="warning-modal-overlay" @click="closeWarningModal">
        <div class="warning-modal" @click.stop>
          <div class="warning-header">
            <div class="warning-icon">⚠️</div>
            <h3 class="warning-title">هشدار</h3>
          </div>
          <div class="warning-content">
            <p class="warning-message">قبل از ایجاد لایه جدید، لایه در حال ویرایش را بروزرسانی کنید یا از ویرایش انصراف دهید.</p>
          </div>
          <div class="warning-actions">
            <button class="btn btn-secondary" @click="closeWarningModal">انصراف</button>
            <button class="btn btn-primary" @click="continueEditing">ادامه ویرایش</button>
          </div>
        </div>
      </div>
      

      

      

    </div>
  </div>
</template>

<script setup lang="ts">
import type { Ref } from 'vue'
import { inject, ref } from 'vue'

interface HeaderLayer {
  id?: string | number
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
  direction: string
  mobileResponsive: boolean
  tabletResponsive: boolean
  paddingRight: number
  paddingLeft: number
  paddingTop: number
  paddingBottom: number
  styleSettings?: Record<string, unknown>
  boxWidths?: number[]
  createdAt?: string
  updatedAt?: string
  [key: string]: unknown
}

// Inject data and functions from parent
// const headerData = inject<Ref<unknown>>('headerData', ref({}))
// const saveHeader = inject<() => void>('saveHeader', () => {})
// const openItemsModal = inject<() => void>('openItemsModal', () => {})
// const getSelectedItemsText = inject<() => string>('getSelectedItemsText', () => 'انتخاب ایتم ها')
const showLayerSettings = inject<Ref<boolean>>('showLayerSettings', ref(false))
const newLayer = inject<Ref<unknown>>('newLayer', ref({}))
// const resetLayerForm = inject<() => void>('resetLayerForm', () => {})
const createdLayers = inject<Ref<unknown[]>>('createdLayers', ref([]))
const clearAllLayers = inject<() => void>('clearAllLayers', () => {})
const isEditing = inject<Ref<boolean>>('isEditing', ref(false))

// متغیرهای مربوط به هشدار
const showWarningModal = ref(false)
const dragIndex = ref<number|null>(null)





function editLayer(layer: HeaderLayer | Record<string, unknown>) {
  // کپی کردن اطلاعات لایه به فرم ویرایش
  Object.assign(newLayer.value, layer)
  // Ensure the ID is preserved for editing
  ;(newLayer.value as Record<string, unknown>).id = (layer as { id?: number | string }).id
  showLayerSettings.value = true
}

function deleteLayer(layerId: number | string | undefined) {
  if (layerId === undefined) return
  const index = createdLayers.value.findIndex(layer => ((layer as { id?: number | string }).id) === layerId)
  if (index > -1) {
    // const deletedLayer = createdLayers.value.splice(index, 1)[0]
    createdLayers.value.splice(index, 1)


  } else {
    console.error('لایه با ID مورد نظر پیدا نشد:', layerId)
  }
}

function confirmClearAllLayers() {
  if (confirm('آیا مطمئن هستید که می‌خواهید همه لایه‌ها را پاک کنید؟')) {
    clearAllLayers()
  }
}

// تابع بررسی و نمایش هشدار
function handleAddNewLayer() {
  // بررسی اینکه آیا لایه‌ای در حال ویرایش است
  if (newLayer.value && ((newLayer.value as { id?: number | string }).id)) {
    showWarningModal.value = true
  } else {
    // اگر لایه‌ای در حال ویرایش نیست، فرم جدید را باز کن
    showLayerSettings.value = true
  }
}

// بستن مودال هشدار
function closeWarningModal() {
  showWarningModal.value = false
}

// ادامه ویرایش لایه فعلی
function continueEditing() {
  showWarningModal.value = false
  // فرم ویرایش باز می‌ماند
}

function onDragStart(index: number) {
  dragIndex.value = index
}
function onDragOver(_index: number) {
  // Highlighting can be added if needed
}
function onDrop(index: number) {
  if (dragIndex.value === null || dragIndex.value === index) return
  const moved = createdLayers.value.splice(dragIndex.value, 1)[0]
  createdLayers.value.splice(index, 0, moved)
  dragIndex.value = null
}
function onDragEnd() {
  dragIndex.value = null
}






</script>

<style scoped>
/* استایل‌های کامپوننت هشدار */
.warning-modal-overlay {
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

.warning-modal {
  background: white;
  border-radius: 12px;
  padding: 24px;
  max-width: 400px;
  width: 90%;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.2);
  text-align: center;
}

.warning-header {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 20px;
}

.warning-icon {
  font-size: 48px;
  margin-bottom: 12px;
}

.warning-title {
  font-size: 20px;
  font-weight: 600;
  color: #2c3e50;
  margin: 0;
}

.warning-content {
  margin-bottom: 24px;
}

.warning-message {
  font-size: 16px;
  color: #495057;
  line-height: 1.5;
  margin: 0;
}

.warning-actions {
  display: flex;
  gap: 12px;
  justify-content: center;
}

.warning-actions .btn {
  padding: 10px 20px;
  border: none;
  border-radius: 8px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
  font-size: 14px;
}

.warning-actions .btn-primary {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}

.warning-actions .btn-primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3);
}

.warning-actions .btn-secondary {
  background: #e9ecef;
  color: #6c757d;
  border: 1px solid #ced4da;
}

.warning-actions .btn-secondary:hover {
  background: #f8f9fa;
  color: #495057;
}

/* استایل‌های کامپوننت پیام موفقیت */
.success-modal-overlay {
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

.success-modal {
  background: white;
  border-radius: 12px;
  padding: 24px;
  max-width: 400px;
  width: 90%;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.2);
  text-align: center;
}

.success-header {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 20px;
}

.success-icon {
  font-size: 48px;
  margin-bottom: 12px;
}

.success-title {
  font-size: 20px;
  font-weight: 600;
  color: #2c3e50;
  margin: 0;
}

.success-content {
  margin-bottom: 24px;
}

.success-message {
  font-size: 16px;
  color: #495057;
  line-height: 1.5;
  margin: 0;
}

.success-actions {
  display: flex;
  justify-content: center;
}

.success-actions .btn {
  padding: 10px 20px;
  border: none;
  border-radius: 8px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
  font-size: 14px;
}

.success-actions .btn-primary {
  background: linear-gradient(135deg, #28a745 0%, #20c997 100%);
  color: white;
}

.success-actions .btn-primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(40, 167, 69, 0.3);
}



.form-container {
  display: flex;
  gap: 20px;
  align-items: flex-start;
}

/* حالت تمام عرض برای ویرایش */
.form-container.full-width {
  width: 100%;
  max-width: none;
}

.form-section {
  background: #fff;
  border-radius: 10px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.04);
  padding: 24px;
  flex: 1;
  width: 100%;
}

/* حالت تمام عرض برای ویرایش */
.form-container.full-width .form-section {
  width: 100%;
  max-width: none;
}

.form-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
  padding-bottom: 16px;
  border-bottom: 1px solid #e9ecef;
}

.btn-danger {
  background: #dc3545;
  color: white;
  border: 1px solid #dc3545;
}

.btn-danger:hover {
  background: #c82333;
  border-color: #bd2130;
}

.layer-hint {
  margin-top: 20px;
  padding: 20px;
  background: #f8f9fa;
  border-radius: 8px;
  border: 1px solid #e9ecef;
  text-align: center;
}

.layer-hint p {
  margin: 0 0 8px 0;
  color: #6c757d;
  font-size: 14px;
}

.layer-hint ul {
  margin: 8px 0 0 0;
  padding-left: 20px;
  text-align: left;
}

.layer-hint li {
  margin-bottom: 4px;
  color: #6c757d;
}

.created-layers {
  margin-top: 20px;
}

.created-layers h4 {
  font-size: 16px;
  font-weight: 600;
  color: #2c3e50;
  margin-bottom: 16px;
}

.layers-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.layer-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px;
  background: #f8f9fa;
  border-radius: 8px;
  border: 1px solid #e9ecef;
  transition: all 0.2s ease;
}

.layer-item:hover {
  background: #e9ecef;
  transform: translateY(-1px);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.layer-info {
  flex: 1;
}

.layer-name {
  font-weight: 600;
  color: #2c3e50;
  margin-bottom: 8px;
  font-size: 14px;
}

.layer-details {
  display: flex;
  gap: 16px;
  flex-wrap: wrap;
}

.layer-details span {
  font-size: 12px;
  color: #6c757d;
  background: #fff;
  padding: 4px 8px;
  border-radius: 4px;
  border: 1px solid #dee2e6;
}

.detail-item {
  white-space: nowrap;
  font-size: 11px;
  padding: 3px 6px;
  background: #f8f9fa;
  border: 1px solid #e9ecef;
  border-radius: 3px;
  color: #495057;
}

.layer-actions {
  display: flex;
  gap: 8px;
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

.btn-success {
  background: linear-gradient(135deg, #28a745 0%, #20c997 100%);
  color: white;
}

.btn-success:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(40, 167, 69, 0.3);
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

.btn-sm {
  padding: 6px 12px;
  font-size: 12px;
}

.btn-outline {
  background: transparent;
  color: #667eea;
  border: 1px solid #667eea;
}

.btn-outline:hover {
  background: #667eea;
  color: white;
}

.btn-danger {
  background: linear-gradient(135deg, #dc3545 0%, #c82333 100%);
  color: white;
}

.btn-danger:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(220, 53, 69, 0.3);
}

.form-group {
  margin-bottom: 16px;
}

.form-group label {
  display: block;
  margin-bottom: 8px;
  font-weight: 500;
  color: #2c3e50;
}

.form-control {
  width: 100%;
  padding: 10px 12px;
  border: 1px solid #ced4da;
  border-radius: 6px;
  font-size: 14px;
  transition: border-color 0.2s ease;
}

.form-control:focus {
  outline: none;
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

/* Responsive styles */
@media (max-width: 768px) {
  .form-container {
    flex-direction: column;
    gap: 16px;
  }
  
  .form-header {
    flex-direction: column;
    gap: 12px;
  }
  
  .layer-item {
    flex-direction: column;
    gap: 12px;
    align-items: flex-start;
  }
  
  .layer-actions {
    width: 100%;
    justify-content: flex-end;
  }
  
  .created-layers {
    margin-top: 16px;
  }
  
  .layers-list {
    gap: 12px;
  }
  
  .layer-details {
    flex-direction: column;
    gap: 8px;
  }
  
  .detail-item {
    font-size: 12px;
  }
}

/* Mobile Responsive */
@media (max-width: 480px) {
  .form-container {
    padding: 16px;
  }
  
  .form-header {
    gap: 8px;
  }
  
  .btn {
    font-size: 14px;
    padding: 12px 16px;
  }
  
  .layer-item {
    padding: 12px;
  }
  
  .layer-name {
    font-size: 16px;
    margin-bottom: 8px;
  }
  
  .layer-details {
    font-size: 12px;
  }
  
  .detail-item {
    padding: 4px 8px;
    background: #f8f9fa;
    border-radius: 4px;
    display: inline-block;
    margin: 2px;
  }
}

/* Tablet Responsive */
@media (max-width: 1024px) and (min-width: 769px) {
  .form-container {
    padding: 20px;
  }
  
  .layer-item {
    padding: 16px;
  }
}

/* RTL Support */
.form-container {
  direction: rtl;
  text-align: right;
}

.layer-item {
  direction: rtl;
  text-align: right;
}

.layer-details {
  direction: rtl;
  text-align: right;
}
</style> 
