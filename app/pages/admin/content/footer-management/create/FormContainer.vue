<template>
  <div class="form-container">
    <div class="layers-management">
      <div class="layers-header">
        <h3>مدیریت لایه‌ها</h3>
        <div class="layers-actions">
          <button
            class="btn btn-primary"
            @click="showLayerSettings = true"
          >
            افزودن لایه جدید
          </button>
          <button
            v-if="createdLayers.length > 0"
            class="btn btn-danger"
            @click="clearAllLayers"
          >
            پاک کردن همه لایه‌ها
          </button>
        </div>
      </div>
      
      <!-- لیست لایه‌های ایجاد شده -->
      <div v-if="createdLayers.length > 0" class="layers-list">
        <div
          v-for="(layer, index) in createdLayers"
          :key="layer.id || index"
          class="layer-item"
          :class="{ 
            'is-dragging': draggingIndex === index,
            'drag-over': dragOverIndex === index
          }"
          :draggable="true"
          @dragstart="handleDragStart(index, $event)"
          @dragend="handleDragEnd"
          @dragover.prevent="handleDragOver(index, $event)"
          @drop="handleDrop(index, $event)"
        >
          <div class="drag-handle">⋮⋮</div>
          <div class="layer-info">
            <div class="layer-name">{{ String(layer.name || 'لایه بدون نام') }}</div>
            <div class="layer-details">
              <span class="detail-item">عرض: {{ typeof layer.width === 'number' ? layer.width : 100 }}%</span>
              <span class="detail-item">ارتفاع: {{ typeof layer.height === 'number' ? layer.height : 50 }}px</span>
              <span class="detail-item">آیتم‌ها: {{ Array.isArray(layer.items) ? layer.items.length : (typeof layer.items === 'string' ? (() => { try { return JSON.parse(layer.items).length } catch { return 0 } })() : 0) }}</span>
            </div>
          </div>
          
          <div class="layer-actions">
            <button
              class="btn btn-sm btn-secondary"
              @click="editLayer(layer)"
            >
              ویرایش
            </button>
            <button
              class="btn btn-sm btn-danger"
              @click="deleteLayer(layer.id as string | number | undefined)"
            >
              حذف
            </button>
          </div>
        </div>
      </div>
      
      <!-- پیام خالی بودن -->
      <div v-else class="empty-layers">
        <div class="empty-message">
          <p>هیچ لایه‌ای ایجاد نشده است</p>
          <p>برای شروع، روی "افزودن لایه جدید" کلیک کنید</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { inject, ref, type Ref } from 'vue'

// Inject functions and data from parent
interface LayerItem {
  id?: string | number
  name?: string
  width?: number
  height?: number
  items?: unknown[] | string
  [key: string]: unknown
}
const createdLayers = inject<Ref<LayerItem[]>>('createdLayers', ref([]))
const showLayerSettings = inject<Ref<boolean>>('showLayerSettings')!
const editLayer = inject<(layer: LayerItem) => void>('editLayer')!
const deleteLayer = inject<(id: string | number | undefined) => void>('deleteLayer')!
const clearAllLayers = inject<() => void>('clearAllLayers')!

// Drag & Drop state
const draggingIndex = ref<number | null>(null)
const dragOverIndex = ref<number | null>(null)

// Drag & Drop functions
function handleDragStart(index: number, event: DragEvent) {
  draggingIndex.value = index
  if (event.dataTransfer) {
    event.dataTransfer.effectAllowed = 'move'
    event.dataTransfer.setData('text/plain', index.toString())
  }
  document.body.classList.add('is-dragging-layer')
}

function handleDragEnd() {
  draggingIndex.value = null
  dragOverIndex.value = null
  document.body.classList.remove('is-dragging-layer')
}

function handleDragOver(index: number, _event: DragEvent) {
  if (draggingIndex.value === null || draggingIndex.value === index) return
  dragOverIndex.value = index
}

function handleDrop(targetIndex: number, event: DragEvent) {
  event.preventDefault()
  
  if (draggingIndex.value === null || draggingIndex.value === targetIndex) {
    handleDragEnd()
    return
  }
  
  const layers = [...createdLayers.value]
  const draggedLayer = layers[draggingIndex.value]
  
  // حذف لایه از موقعیت قبلی
  layers.splice(draggingIndex.value, 1)
  
  // اضافه کردن لایه به موقعیت جدید
  const newIndex = draggingIndex.value < targetIndex ? targetIndex - 1 : targetIndex
  layers.splice(newIndex, 0, draggedLayer)
  
  // آپدیت لیست لایه‌ها
  createdLayers.value = layers
  
  // ذخیره در sessionStorage اگر در حال ویرایش نیستیم
  if (typeof window !== 'undefined' && !window.location.pathname.includes('edit')) {
    sessionStorage.setItem('footerLayers', JSON.stringify(layers))
  }
  
  handleDragEnd()
}
</script>

<style scoped>
.form-container {
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  padding: 24px;
  margin-bottom: 24px;
}

.layers-management h3 {
  margin: 0 0 20px 0;
  font-size: 18px;
  font-weight: 600;
  color: #1f2937;
}

.layers-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.layers-actions {
  display: flex;
  gap: 12px;
}

.btn {
  padding: 10px 16px;
  border: none;
  border-radius: 6px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
  text-decoration: none;
  display: inline-flex;
  align-items: center;
  justify-content: center;
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

.btn-sm {
  padding: 6px 12px;
  font-size: 12px;
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
  padding-right: 48px;
  border: 1px solid #e5e7eb;
  border-radius: 6px;
  background: #f9fafb;
  transition: all 0.2s;
  cursor: grab;
  position: relative;
}

.layer-item:active {
  cursor: grabbing;
}

.layer-item.is-dragging {
  opacity: 0.5;
  transform: scale(0.98);
  cursor: grabbing;
}

.layer-item.drag-over {
  background: #dbeafe;
  border-color: #3b82f6;
  transform: translateY(-2px);
  box-shadow: 0 4px 6px rgba(59, 130, 246, 0.2);
}

.drag-handle {
  position: absolute;
  right: 16px;
  top: 50%;
  transform: translateY(-50%);
  font-size: 18px;
  color: #9ca3af;
  cursor: grab;
  padding: 4px;
  user-select: none;
}

.drag-handle:active {
  cursor: grabbing;
}

:global(body.is-dragging-layer) {
  cursor: grabbing !important;
  user-select: none;
}

.layer-info {
  flex: 1;
}

.layer-name {
  font-weight: 600;
  color: #1f2937;
  margin-bottom: 4px;
}

.layer-details {
  display: flex;
  gap: 16px;
}

.detail-item {
  font-size: 12px;
  color: #6b7280;
}

.layer-actions {
  display: flex;
  gap: 8px;
}

.empty-layers {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 48px 24px;
  border: 2px dashed #d1d5db;
  border-radius: 8px;
  background: #f9fafb;
}

.empty-message {
  text-align: center;
  color: #6b7280;
}

.empty-message p {
  margin: 4px 0;
  font-size: 14px;
}

@media (max-width: 768px) {
  .layers-header {
    flex-direction: column;
    align-items: stretch;
    gap: 16px;
  }
  
  .layers-actions {
    justify-content: stretch;
  }
  
  .btn {
    flex: 1;
  }
  
  .layer-item {
    flex-direction: column;
    align-items: stretch;
    gap: 12px;
  }
  
  .layer-actions {
    justify-content: stretch;
  }
  
  .layer-actions .btn {
    flex: 1;
  }
}
</style>


