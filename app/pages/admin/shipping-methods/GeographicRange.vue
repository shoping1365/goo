<template>
  <div class="geographic-range">
    <div class="settings-card">
      <div class="card-header">
        <h3>محدوده جغرافیایی</h3>
        <div class="toggle-switch">
          <input id="geographicToggle" v-model="enabled" type="checkbox" />
          <label for="geographicToggle"></label>
        </div>
      </div>
      
      <div v-if="enabled" class="card-body">
        <div class="form-section">
          <h4>نوع محدوده</h4>
          <div class="form-row">
            <div class="form-group">
              <label>سطح محدوده</label>
              <select v-model="rangeType" class="form-select" @change="onRangeTypeChange">
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
          <h4>انتخاب مناطق</h4>
          <div class="regions-selection">
            <div v-if="rangeType === 'province'" class="province-selection">
              <h5>انتخاب استان‌ها</h5>
              <div class="provinces-grid">
                <label v-for="province in provinces" :key="province.value" class="province-checkbox">
                  <input 
                    v-model="selectedProvinces" 
                    type="checkbox" 
                    :value="province.value"
                  />
                  {{ province.label }}
                </label>
              </div>
            </div>
            
            <div v-if="rangeType === 'city'" class="city-selection">
              <h5>انتخاب شهرها</h5>
              <div class="cities-grid">
                <label v-for="city in cities" :key="city.value" class="city-checkbox">
                  <input 
                    v-model="selectedCities" 
                    type="checkbox" 
                    :value="city.value"
                  />
                  {{ city.label }}
                </label>
              </div>
            </div>
            
            <div v-if="rangeType === 'district'" class="district-selection">
              <h5>انتخاب محله‌ها</h5>
              <div class="districts-grid">
                <label v-for="district in districts" :key="district.value" class="district-checkbox">
                  <input 
                    v-model="selectedDistricts" 
                    type="checkbox" 
                    :value="district.value"
                  />
                  {{ district.label }}
                </label>
              </div>
            </div>
            
            <div v-if="rangeType === 'custom'" class="custom-range">
              <h5>محدوده سفارشی</h5>
              <div class="custom-range-form">
                <div class="form-row">
                  <div class="form-group">
                    <label>مرکز محدوده (طول جغرافیایی)</label>
                    <input 
                      v-model.number="customRange.centerLng" 
                      type="number" 
                      step="0.000001"
                      class="form-input"
                      placeholder="51.3890"
                    />
                  </div>
                  <div class="form-group">
                    <label>مرکز محدوده (عرض جغرافیایی)</label>
                    <input 
                      v-model.number="customRange.centerLat" 
                      type="number" 
                      step="0.000001"
                      class="form-input"
                      placeholder="35.6892"
                    />
                  </div>
                </div>
                <div class="form-group">
                  <label>شعاع محدوده</label>
                  <div class="radius-input-group">
                    <input 
                      v-model.number="customRange.radius" 
                      type="number" 
                      min="0" 
                      step="0.1"
                      class="form-input"
                      placeholder="10"
                    />
                    <span class="unit-label">{{ getDistanceUnitLabel() }}</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        
        <div class="form-section">
          <h4>پیش‌نمایش محدوده</h4>
          <div class="range-preview">
            <div class="preview-info">
              <div class="preview-item">
                <span>نوع محدوده:</span>
                <strong>{{ getRangeTypeLabel() }}</strong>
              </div>
              <div class="preview-item">
                <span>تعداد مناطق انتخاب شده:</span>
                <strong>{{ getSelectedCount() }}</strong>
              </div>
              <div v-if="rangeType === 'custom'" class="preview-item">
                <span>شعاع محدوده:</span>
                <strong>{{ customRange.radius }} {{ getDistanceUnitLabel() }}</strong>
              </div>
            </div>
            <div class="map-preview">
              <div class="map-placeholder">
                <div class="map-icon">🗺️</div>
                <p>نقشه محدوده جغرافیایی</p>
                <button class="view-map-btn">نمایش نقشه</button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
</script>

<script setup lang="ts">
import { ref, watch } from 'vue';

definePageMeta({ layout: 'admin-main', middleware: 'admin' });

const props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({
      enabled: false,
      rangeType: 'province',
      distanceUnit: 'km',
      selectedProvinces: [],
      selectedCities: [],
      selectedDistricts: [],
      customRange: {
        centerLng: 51.3890,
        centerLat: 35.6892,
        radius: 10
      }
    })
  }
})

const emit = defineEmits(['update:modelValue'])

const enabled = ref(props.modelValue.enabled)
const rangeType = ref(props.modelValue.rangeType)
const distanceUnit = ref(props.modelValue.distanceUnit)
const selectedProvinces = ref(props.modelValue.selectedProvinces)
const selectedCities = ref(props.modelValue.selectedCities)
const selectedDistricts = ref(props.modelValue.selectedDistricts)
const customRange = ref(props.modelValue.customRange)

// داده‌های نمونه
const provinces = [
  { value: 'tehran', label: 'تهران' },
  { value: 'isfahan', label: 'اصفهان' },
  { value: 'shiraz', label: 'شیراز' },
  { value: 'mashhad', label: 'مشهد' },
  { value: 'tabriz', label: 'تبریز' },
  { value: 'kerman', label: 'کرمان' },
  { value: 'yazd', label: 'یزد' },
  { value: 'kermanshah', label: 'کرمانشاه' }
]

const cities = [
  { value: 'tehran_city', label: 'شهر تهران' },
  { value: 'isfahan_city', label: 'شهر اصفهان' },
  { value: 'shiraz_city', label: 'شهر شیراز' },
  { value: 'mashhad_city', label: 'شهر مشهد' },
  { value: 'tabriz_city', label: 'شهر تبریز' }
]

const districts = [
  { value: 'tehran_center', label: 'مرکز تهران' },
  { value: 'tehran_north', label: 'شمال تهران' },
  { value: 'tehran_south', label: 'جنوب تهران' },
  { value: 'tehran_east', label: 'شرق تهران' },
  { value: 'tehran_west', label: 'غرب تهران' }
]

watch([enabled, rangeType, distanceUnit, selectedProvinces, selectedCities, selectedDistricts, customRange], () => {
  emit('update:modelValue', {
    enabled: enabled.value,
    rangeType: rangeType.value,
    distanceUnit: distanceUnit.value,
    selectedProvinces: selectedProvinces.value,
    selectedCities: selectedCities.value,
    selectedDistricts: selectedDistricts.value,
    customRange: customRange.value
  })
}, { deep: true })

function onRangeTypeChange() {
  // پاک کردن انتخاب‌های قبلی هنگام تغییر نوع محدوده
  selectedProvinces.value = []
  selectedCities.value = []
  selectedDistricts.value = []
}

function getRangeTypeLabel() {
  const labels = {
    'province': 'استان',
    'city': 'شهر',
    'district': 'محله',
    'custom': 'محدوده سفارشی'
  }
  return labels[rangeType.value] || rangeType.value
}

function getDistanceUnitLabel() {
  const labels = {
    'km': 'کیلومتر',
    'm': 'متر',
    'mile': 'مایل'
  }
  return labels[distanceUnit.value] || distanceUnit.value
}

function getSelectedCount() {
  switch (rangeType.value) {
    case 'province':
      return selectedProvinces.value.length
    case 'city':
      return selectedCities.value.length
    case 'district':
      return selectedDistricts.value.length
    case 'custom':
      return 1
    default:
      return 0
  }
}
</script>

<style scoped>
.geographic-range {
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

.form-section h5 {
  margin: 0 0 1rem 0;
  font-size: 0.9rem;
  color: #666;
  font-weight: 500;
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

.radius-input-group {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.radius-input-group .form-input {
  flex: 1;
}

.unit-label {
  font-size: 0.9rem;
  color: #666;
  white-space: nowrap;
}

.provinces-grid, .cities-grid, .districts-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 0.5rem;
  background: white;
  border-radius: 8px;
  border: 1px solid #e9ecef;
  padding: 1rem;
}

.province-checkbox, .city-checkbox, .district-checkbox {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-size: 0.9rem;
  cursor: pointer;
  padding: 0.5rem;
  background: #f8f9fa;
  border-radius: 4px;
  border: 1px solid #e9ecef;
  transition: all 0.2s;
}

.province-checkbox:hover, .city-checkbox:hover, .district-checkbox:hover {
  background: #e9ecef;
}

.province-checkbox input[type="checkbox"], .city-checkbox input[type="checkbox"], .district-checkbox input[type="checkbox"] {
  width: 16px;
  height: 16px;
}

.custom-range-form {
  background: white;
  border-radius: 8px;
  border: 1px solid #e9ecef;
  padding: 1rem;
}

.range-preview {
  background: white;
  border-radius: 8px;
  border: 1px solid #e9ecef;
  padding: 1rem;
}

.preview-info {
  margin-bottom: 1rem;
}

.preview-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0.5rem 0;
  border-bottom: 1px solid #f1f3f4;
}

.preview-item:last-child {
  border-bottom: none;
}

.preview-item span {
  color: #666;
  font-size: 0.9rem;
}

.preview-item strong {
  color: #333;
  font-weight: 600;
}

.map-preview {
  border-top: 1px solid #e9ecef;
  padding-top: 1rem;
}

.map-placeholder {
  padding: 2rem;
  text-align: center;
  background: #f8f9fa;
  border-radius: 6px;
  border: 1px solid #e9ecef;
}

.map-icon {
  font-size: 2rem;
  margin-bottom: 0.5rem;
}

.map-placeholder p {
  margin: 0 0 1rem 0;
  color: #666;
  font-size: 0.9rem;
}

.view-map-btn {
  background: #2196F3;
  color: white;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 6px;
  cursor: pointer;
  font-size: 0.9rem;
}

.view-map-btn:hover {
  background: #1976d2;
}
</style> 
