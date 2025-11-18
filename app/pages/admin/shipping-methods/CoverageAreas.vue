<template>
  <div class="coverage-areas rtl">
    <div class="coverage-header">
      <h3>مدیریت مناطق تحت پوشش</h3>
      <div class="stats">
        <span class="stat">مناطق تحت پوشش: {{ coveredAreas.length }}</span>
        <span class="stat">مناطق مستثنی: {{ excludedAreas.length }}</span>
      </div>
    </div>

    <div class="coverage-content">
      <!-- نقشه تعاملی -->
      <div class="map-section">
        <h4>نقشه مناطق</h4>
        <div class="interactive-map">
          <div class="map-placeholder">
            <div class="map-icon">🗺️</div>
            <h5>نقشه تعاملی</h5>
            <p>برای انتخاب مناطق روی نقشه کلیک کنید</p>

            <div class="map-legend">
              <div class="legend-item">
                <span class="legend-color covered"></span>
                <span>مناطق تحت پوشش</span>
              </div>
              <div class="legend-item">
                <span class="legend-color excluded"></span>
                <span>مناطق مستثنی</span>
              </div>
            </div>

            <div class="map-controls">
              <button @click="selectByProvince" class="map-btn">انتخاب استان</button>
              <button @click="selectByCity" class="map-btn">انتخاب شهر</button>
              <button @click="selectByDistrict" class="map-btn">انتخاب محله</button>
            </div>
          </div>
        </div>
      </div>

      <!-- لیست مناطق -->
      <div class="areas-section">
        <div class="tabs">
          <button @click="activeTab = 'covered'" :class="['tab', { active: activeTab === 'covered' }]">
            مناطق تحت پوشش
          </button>
          <button @click="activeTab = 'excluded'" :class="['tab', { active: activeTab === 'excluded' }]">
            مناطق مستثنی
          </button>
        </div>

        <div class="areas-list">
          <div class="list-header">
            <h4>{{ activeTab === 'covered' ? 'مناطق تحت پوشش' : 'مناطق مستثنی' }}</h4>
            <button @click="addArea" class="add-btn">+ افزودن</button>
          </div>

          <div class="areas-grid">
            <div v-for="(area, index) in currentAreas" :key="index" class="area-card">
              <div class="area-header">
                <span class="area-type">{{ getAreaTypeLabel(area.type) }}</span>
                <div class="actions">
                  <button @click="editArea(index)" class="edit-btn">✏️</button>
                  <button @click="removeArea(index)" class="delete-btn">🗑️</button>
                </div>
              </div>

              <div class="area-body">
                <h5>{{ area.name }}</h5>
                <p v-if="area.description">{{ area.description }}</p>

                <div class="area-info">
                  <span>اولویت: {{ area.priority }}</span>
                  <span v-if="area.distance">فاصله: {{ area.distance }} کیلومتر</span>
                </div>

                <div v-if="area.subAreas && area.subAreas.length" class="sub-areas">
                  <span>زیرمناطق:</span>
                  <div class="sub-areas-tags">
                    <span v-for="subArea in area.subAreas" :key="subArea" class="tag">
                      {{ subArea }}
                    </span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- مودال -->
    <div v-if="showModal" class="modal" @click="closeModal">
      <div class="modal-content" @click.stop>
        <div class="modal-header">
          <h4>{{ isEditing ? 'ویرایش منطقه' : 'افزودن منطقه' }}</h4>
          <button @click="closeModal" class="close-btn">×</button>
        </div>

        <div class="modal-body">
          <div class="form-group">
            <label>نام منطقه</label>
            <input v-model="editingArea.name" type="text" placeholder="نام منطقه" />
          </div>

          <div class="form-group">
            <label>نوع منطقه</label>
            <select v-model="editingArea.type">
              <option value="province">استان</option>
              <option value="city">شهر</option>
              <option value="district">محله</option>
              <option value="custom">سفارشی</option>
            </select>
          </div>

          <div class="form-group">
            <label>توضیحات</label>
            <textarea v-model="editingArea.description" placeholder="توضیحات"></textarea>
          </div>

          <div class="form-row">
            <div class="form-group">
              <label>اولویت</label>
              <input v-model.number="editingArea.priority" type="number" min="1" />
            </div>
            <div class="form-group">
              <label>فاصله (کیلومتر)</label>
              <input v-model.number="editingArea.distance" type="number" min="0" />
            </div>
          </div>
        </div>

        <div class="modal-footer">
          <button @click="closeModal" class="btn-cancel">انصراف</button>
          <button @click="saveArea" class="btn-save">ذخیره</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'

const props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({
      coveredAreas: [],
      excludedAreas: []
    })
  }
})

const emit = defineEmits(['update:modelValue'])

const activeTab = ref('covered')
const showModal = ref(false)
const isEditing = ref(false)
const editingIndex = ref(-1)

const coveredAreas = ref(props.modelValue.coveredAreas.length ? props.modelValue.coveredAreas : [
  { name: 'تهران', type: 'province', description: 'مرکز استان تهران', priority: 1, distance: 0, subAreas: ['شهریار', 'ورامین'] },
  { name: 'اصفهان', type: 'province', description: 'استان اصفهان', priority: 2, distance: 450, subAreas: ['کاشان', 'نجف‌آباد'] }
])

const excludedAreas = ref(props.modelValue.excludedAreas.length ? props.modelValue.excludedAreas : [
  { name: 'مناطق کوهستانی', type: 'custom', description: 'عدم دسترسی', priority: 1, distance: 0 }
])

const editingArea = ref({
  name: '',
  type: 'province',
  description: '',
  priority: 1,
  distance: 0
})

const currentAreas = computed(() => activeTab.value === 'covered' ? coveredAreas.value : excludedAreas.value)

function selectByProvince() {
  console.log('انتخاب بر اساس استان')
}

function selectByCity() {
  console.log('انتخاب بر اساس شهر')
}

function selectByDistrict() {
  console.log('انتخاب بر اساس محله')
}

function addArea() {
  isEditing.value = false
  editingIndex.value = -1
  editingArea.value = { name: '', type: 'province', description: '', priority: 1, distance: 0 }
  showModal.value = true
}

function editArea(index) {
  isEditing.value = true
  editingIndex.value = index
  editingArea.value = { ...currentAreas.value[index] }
  showModal.value = true
}

function removeArea(index) {
  if (confirm('آیا از حذف این منطقه اطمینان دارید؟')) {
    if (activeTab.value === 'covered') {
      coveredAreas.value.splice(index, 1)
    } else {
      excludedAreas.value.splice(index, 1)
    }
  }
}

function saveArea() {
  if (!editingArea.value.name.trim()) {
    alert('لطفاً نام منطقه را وارد کنید')
    return
  }

  if (isEditing.value) {
    if (activeTab.value === 'covered') {
      coveredAreas.value[editingIndex.value] = { ...editingArea.value }
    } else {
      excludedAreas.value[editingIndex.value] = { ...editingArea.value }
    }
  } else {
    if (activeTab.value === 'covered') {
      coveredAreas.value.push({ ...editingArea.value })
    } else {
      excludedAreas.value.push({ ...editingArea.value })
    }
  }

  closeModal()
}

function closeModal() {
  showModal.value = false
}

function getAreaTypeLabel(type) {
  const labels = { province: 'استان', city: 'شهر', district: 'محله', custom: 'سفارشی' }
  return labels[type] || type
}
</script>

<style scoped>
.coverage-areas {
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  overflow: hidden;
}

.coverage-header {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  padding: 1.5rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.coverage-header h3 {
  margin: 0;
  font-size: 1.3rem;
}

.stats {
  display: flex;
  gap: 2rem;
}

.stat {
  text-align: center;
}

.coverage-content {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 2rem;
  padding: 2rem;
}

.map-section {
  background: #f8f9fa;
  border-radius: 8px;
  padding: 1.5rem;
}

.map-section h4 {
  margin: 0 0 1rem 0;
  color: #333;
}

.interactive-map {
  background: white;
  border-radius: 8px;
  border: 1px solid #e9ecef;
  padding: 2rem;
  text-align: center;
}

.map-icon {
  font-size: 3rem;
  margin-bottom: 1rem;
}

.map-legend {
  display: flex;
  justify-content: center;
  gap: 2rem;
  margin: 2rem 0;
}

.legend-item {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.legend-color {
  width: 16px;
  height: 16px;
  border-radius: 4px;
}

.legend-color.covered { background: #28a745; }
.legend-color.excluded { background: #dc3545; }

.map-controls {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.map-btn {
  background: #17a2b8;
  color: white;
  border: none;
  padding: 0.75rem;
  border-radius: 6px;
  cursor: pointer;
}

.areas-section {
  background: #f8f9fa;
  border-radius: 8px;
  padding: 1.5rem;
}

.tabs {
  display: flex;
  margin-bottom: 1.5rem;
  background: white;
  border-radius: 8px;
  padding: 0.25rem;
}

.tab {
  flex: 1;
  background: none;
  border: none;
  padding: 0.75rem;
  border-radius: 6px;
  cursor: pointer;
  color: #666;
}

.tab.active {
  background: #007bff;
  color: white;
}

.list-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}

.list-header h4 {
  margin: 0;
  color: #333;
}

.add-btn {
  background: #28a745;
  color: white;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 6px;
  cursor: pointer;
}

.areas-grid {
  display: grid;
  gap: 1rem;
  max-height: 60vh;
  overflow-y: auto;
}
@media (min-width: 768px) {
  .areas-grid {
    max-height: 70vh;
  }
}

.area-card {
  background: white;
  border-radius: 8px;
  border: 1px solid #e9ecef;
  overflow: hidden;
}

.area-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0.75rem 1rem;
  background: #f8f9fa;
  border-bottom: 1px solid #e9ecef;
}

.area-type {
  background: #e3f2fd;
  color: #1976d2;
  padding: 0.25rem 0.75rem;
  border-radius: 20px;
  font-size: 0.8rem;
}

.actions {
  display: flex;
  gap: 0.5rem;
}

.edit-btn, .delete-btn {
  background: none;
  border: none;
  cursor: pointer;
  padding: 0.25rem;
  border-radius: 4px;
}

.edit-btn:hover { background: #f8f9fa; }
.delete-btn:hover { background: #ffebee; }

.area-body {
  padding: 1rem;
}

.area-body h5 {
  margin: 0 0 0.5rem 0;
  color: #333;
}

.area-body p {
  margin: 0 0 1rem 0;
  color: #666;
  font-size: 0.9rem;
}

.area-info {
  display: flex;
  gap: 1rem;
  margin-bottom: 1rem;
  font-size: 0.85rem;
  color: #666;
}

.sub-areas {
  border-top: 1px solid #e9ecef;
  padding-top: 0.75rem;
}

.sub-areas span {
  display: block;
  margin-bottom: 0.5rem;
  font-size: 0.85rem;
  color: #666;
}

.sub-areas-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 0.25rem;
}

.tag {
  background: #e9ecef;
  color: #495057;
  padding: 0.25rem 0.5rem;
  border-radius: 12px;
  font-size: 0.8rem;
}

.modal {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal-content {
  background: white;
  border-radius: 12px;
  width: 90%;
  max-width: 500px;
  max-height: 80vh;
  overflow-y: auto;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.3);
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1.5rem;
  border-bottom: 1px solid #e9ecef;
  background: #f8f9fa;
  border-radius: 12px 12px 0 0;
}

.modal-header h4 {
  margin: 0;
  color: #333;
}

.close-btn {
  background: none;
  border: none;
  font-size: 1.5rem;
  cursor: pointer;
  color: #666;
  padding: 0;
  width: 30px;
  height: 30px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
}

.close-btn:hover {
  background: #e9ecef;
}

.modal-body {
  padding: 1.5rem;
}

.form-group {
  margin-bottom: 1.5rem;
}

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  color: #333;
  font-weight: 500;
}

.form-group input,
.form-group select,
.form-group textarea {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 0.9rem;
  font-family: inherit;
}

.form-group textarea {
  resize: vertical;
  min-height: 80px;
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1rem;
}

.modal-footer {
  display: flex;
  justify-content: flex-end;
  gap: 1rem;
  padding: 1rem 1.5rem;
  border-top: 1px solid #e9ecef;
  background: #f8f9fa;
  border-radius: 0 0 12px 12px;
}

.btn-cancel,
.btn-save {
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-weight: 500;
}

.btn-cancel {
  background: #6c757d;
  color: white;
}

.btn-cancel:hover {
  background: #5a6268;
}

.btn-save {
  background: #28a745;
  color: white;
}

.btn-save:hover {
  background: #218838;
}

@media (max-width: 768px) {
  .coverage-content {
    grid-template-columns: 1fr;
  }
  
  .modal-content {
    width: 95%;
    margin: 1rem;
  }
  
  .form-row {
    grid-template-columns: 1fr;
  }
  
  .stats {
    flex-direction: column;
    gap: 0.5rem;
  }
  
  .coverage-header {
    flex-direction: column;
    text-align: center;
    gap: 1rem;
  }
}
</style>
