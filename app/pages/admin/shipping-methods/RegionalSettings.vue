<template>
  <div class="regional-settings">
    <div class="settings-card">
      <div class="card-header">
        <h3>تنظیمات منطقه‌ای</h3>
        <div class="toggle-switch">
          <input id="regionalToggle" v-model="enabled" type="checkbox" />
          <label for="regionalToggle"></label>
        </div>
      </div>
      
      <div v-if="enabled" class="card-body">
        <div class="form-section">
          <h4>محدوده جغرافیایی</h4>
          
          <div class="form-row">
            <div class="form-group">
              <label>نوع محدوده</label>
              <select v-model="regionType" class="form-select">
                <option value="province">استان</option>
                <option value="city">شهر</option>
                <option value="district">محله</option>
                <option value="custom">محدوده سفارشی</option>
              </select>
            </div>
            <div class="form-group">
              <label>واحد فاصله</label>
              <select v-model="distanceUnit" class="form-select">
                <option value="km">کیلومتر</option>
                <option value="m">متر</option>
                <option value="mile">مایل</option>
              </select>
            </div>
          </div>
        </div>
        
        <div class="form-section">
          <h4>مناطق تحت پوشش</h4>
          <div class="regions-list">
            <div v-for="(region, index) in coveredRegions" :key="index" class="region-item">
              <div class="region-info">
                <span class="region-name">{{ region.name }}</span>
                <span class="region-type">{{ getRegionTypeLabel(region.type) }}</span>
              </div>
              <button class="remove-btn" @click="removeRegion(index)">حذف</button>
            </div>
            <button class="add-region-btn" @click="addRegion">
              <span>+</span> افزودن منطقه
            </button>
          </div>
        </div>
        
        <div class="form-section">
          <h4>مناطق مستثنی</h4>
          <div class="excluded-regions">
            <div v-for="(region, index) in excludedRegions" :key="index" class="region-item">
              <div class="region-info">
                <span class="region-name">{{ region.name }}</span>
                <span class="region-type">{{ getRegionTypeLabel(region.type) }}</span>
              </div>
              <button class="remove-btn" @click="removeExcludedRegion(index)">حذف</button>
            </div>
            <button class="add-region-btn" @click="addExcludedRegion">
              <span>+</span> افزودن منطقه مستثنی
            </button>
          </div>
        </div>
        
        <div class="form-section">
          <h4>محدودیت‌های فاصله</h4>
          <div class="form-row">
            <div class="form-group">
              <label>حداکثر فاصله</label>
              <div class="distance-input-group">
                <input 
                  v-model.number="maxDistance" 
                  type="number" 
                  min="0" 
                  step="0.1"
                  class="form-input"
                  placeholder="0"
                />
                <span class="unit-label">{{ getDistanceUnitLabel() }}</span>
              </div>
            </div>
            <div class="form-group">
              <label>حداقل فاصله</label>
              <div class="distance-input-group">
                <input 
                  v-model.number="minDistance" 
                  type="number" 
                  min="0" 
                  step="0.1"
                  class="form-input"
                  placeholder="0"
                />
                <span class="unit-label">{{ getDistanceUnitLabel() }}</span>
              </div>
            </div>
          </div>
        </div>
        
        <div class="form-section">
          <h4>نقشه مناطق</h4>
          <div class="map-container">
            <div class="map-placeholder">
              <div class="map-icon">🗺️</div>
              <p>نقشه مناطق تحت پوشش</p>
              <button class="map-btn">نمایش نقشه</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'

const props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({
      enabled: false,
      regionType: 'province',
      distanceUnit: 'km',
      coveredRegions: [],
      excludedRegions: [],
      maxDistance: 0,
      minDistance: 0
    })
  }
})

const emit = defineEmits(['update:modelValue'])

const enabled = ref(props.modelValue.enabled)
const regionType = ref(props.modelValue.regionType)
const distanceUnit = ref(props.modelValue.distanceUnit)
const coveredRegions = ref(props.modelValue.coveredRegions.length ? props.modelValue.coveredRegions : [
  { name: 'تهران', type: 'province' },
  { name: 'اصفهان', type: 'province' }
])
const excludedRegions = ref(props.modelValue.excludedRegions.length ? props.modelValue.excludedRegions : [])
const maxDistance = ref(props.modelValue.maxDistance)
const minDistance = ref(props.modelValue.minDistance)

watch([enabled, regionType, distanceUnit, coveredRegions, excludedRegions, maxDistance, minDistance], () => {
  emit('update:modelValue', {
    enabled: enabled.value,
    regionType: regionType.value,
    distanceUnit: distanceUnit.value,
    coveredRegions: coveredRegions.value,
    excludedRegions: excludedRegions.value,
    maxDistance: maxDistance.value,
    minDistance: minDistance.value
  })
}, { deep: true })

function addRegion() {
  coveredRegions.value.push({
    name: 'منطقه جدید',
    type: regionType.value
  })
}

function removeRegion(index) {
  coveredRegions.value.splice(index, 1)
}

function addExcludedRegion() {
  excludedRegions.value.push({
    name: 'منطقه مستثنی',
    type: regionType.value
  })
}

function removeExcludedRegion(index) {
  excludedRegions.value.splice(index, 1)
}

function getRegionTypeLabel(type) {
  const labels = {
    'province': 'استان',
    'city': 'شهر',
    'district': 'محله',
    'custom': 'سفارشی'
  }
  return labels[type] || type
}

function getDistanceUnitLabel() {
  const labels = {
    'km': 'کیلومتر',
    'm': 'متر',
    'mile': 'مایل'
  }
  return labels[distanceUnit.value] || distanceUnit.value
}
</script>

<style scoped>
.regional-settings {
  direction: rtl;
}

.settings-card {
  background: #f8f9fa;
  border-radius: 12px;
  border: 1px solid #e9ecef;
  overflow: hidden;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem 1.5rem;
  background: #fff;
  border-bottom: 1px solid #e9ecef;
}

.card-header h3 {
  margin: 0;
  font-size: 1.1rem;
  font-weight: 600;
  color: #333;
}

.toggle-switch {
  position: relative;
  width: 50px;
  height: 24px;
}

.toggle-switch input {
  opacity: 0;
  width: 0;
  height: 0;
}

.toggle-switch label {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: #ccc;
  transition: .4s;
  border-radius: 24px;
}

.toggle-switch label:before {
  position: absolute;
  content: "";
  height: 18px;
  width: 18px;
  left: 3px;
  bottom: 3px;
  background-color: white;
  transition: .4s;
  border-radius: 50%;
}

.toggle-switch input:checked + label {
  background-color: #2196F3;
}

.toggle-switch input:checked + label:before {
  transform: translateX(26px);
}

.card-body {
  padding: 1.5rem;
}

.form-section {
  margin-bottom: 2rem;
}

.form-section h4 {
  margin: 0 0 1rem 0;
  font-size: 1rem;
  color: #555;
  font-weight: 600;
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1rem;
  margin-bottom: 1rem;
}

.form-group {
  margin-bottom: 1rem;
}

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: 500;
  color: #555;
  font-size: 0.9rem;
}

.form-input, .form-select {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 0.9rem;
  background: #fafbfc;
}

.form-input:focus, .form-select:focus {
  border-color: #2196F3;
  outline: none;
}

.distance-input-group {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.distance-input-group .form-input {
  flex: 1;
}

.unit-label {
  font-size: 0.9rem;
  color: #666;
  white-space: nowrap;
}

.regions-list, .excluded-regions {
  background: white;
  border-radius: 8px;
  border: 1px solid #e9ecef;
  padding: 1rem;
}

.region-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0.8rem;
  background: #f8f9fa;
  border-radius: 6px;
  margin-bottom: 0.5rem;
  border: 1px solid #e9ecef;
}

.region-info {
  display: flex;
  flex-direction: column;
  gap: 0.2rem;
}

.region-name {
  font-weight: 500;
  color: #333;
}

.region-type {
  font-size: 0.8rem;
  color: #666;
}

.remove-btn {
  background: #dc3545;
  color: white;
  border: none;
  padding: 0.4rem 0.8rem;
  border-radius: 4px;
  cursor: pointer;
  font-size: 0.8rem;
}

.remove-btn:hover {
  background: #c82333;
}

.add-region-btn {
  background: #28a745;
  color: white;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 6px;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-size: 0.9rem;
  width: 100%;
  justify-content: center;
}

.add-region-btn:hover {
  background: #218838;
}

.map-container {
  background: white;
  border-radius: 8px;
  border: 1px solid #e9ecef;
  overflow: hidden;
}

.map-placeholder {
  padding: 3rem 2rem;
  text-align: center;
  background: #f8f9fa;
}

.map-icon {
  font-size: 3rem;
  margin-bottom: 1rem;
}

.map-placeholder p {
  margin: 0 0 1rem 0;
  color: #666;
}

.map-btn {
  background: #2196F3;
  color: white;
  border: none;
  padding: 0.5rem 1.5rem;
  border-radius: 6px;
  cursor: pointer;
  font-size: 0.9rem;
}

.map-btn:hover {
  background: #1976d2;
}
</style> 
