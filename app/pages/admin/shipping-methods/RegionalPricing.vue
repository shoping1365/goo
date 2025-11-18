<template>
  <div class="regional-pricing">
    <div class="pricing-card">
      <div class="card-header">
        <h3>قیمت منطقه‌ای</h3>
        <div class="toggle-switch">
          <input type="checkbox" v-model="enabled" id="regionalToggle" />
          <label for="regionalToggle"></label>
        </div>
      </div>
      
      <div v-if="enabled" class="card-body">
        <div class="regional-settings">
          <div class="form-group">
            <label>نوع منطقه‌بندی</label>
            <select v-model="regionType" class="region-type-select">
              <option value="province">استان</option>
              <option value="city">شهر</option>
              <option value="postal_code">کد پستی</option>
              <option value="custom">سفارشی</option>
            </select>
          </div>
          
          <div class="regions-section">
            <div class="table-header">
              <h4>مناطق و قیمت‌ها</h4>
              <button @click="addRegion" class="add-btn">
                <span>+</span> افزودن منطقه
              </button>
            </div>
            
            <div class="regions-table">
              <div class="table-row header">
                <div class="col">نام منطقه</div>
                <div class="col">کد</div>
                <div class="col">قیمت</div>
                <div class="col">وضعیت</div>
                <div class="col actions">عملیات</div>
              </div>
              
              <div v-for="(region, index) in regions" :key="index" class="table-row">
                <div class="col">
                  <input 
                    v-model="region.name" 
                    type="text"
                    class="region-input"
                    placeholder="نام منطقه"
                  />
                </div>
                <div class="col">
                  <input 
                    v-model="region.code" 
                    type="text"
                    class="region-input"
                    placeholder="کد منطقه"
                  />
                </div>
                <div class="col">
                  <div class="price-input-group">
                    <input 
                      v-model.number="region.price" 
                      type="number" 
                      min="0" 
                      step="1000"
                      class="price-input"
                      placeholder="0"
                    />
                    <span class="currency-label">تومان</span>
                  </div>
                </div>
                <div class="col">
                  <select v-model="region.active" class="status-select">
                    <option :value="true">فعال</option>
                    <option :value="false">غیرفعال</option>
                  </select>
                </div>
                <div class="col actions">
                  <button @click="removeRegion(index)" class="remove-btn">
                    حذف
                  </button>
                </div>
              </div>
            </div>
          </div>
          
          <div class="bulk-actions">
            <h4>عملیات گروهی</h4>
            <div class="bulk-controls">
              <div class="bulk-input-group">
                <label>افزودن قیمت به همه:</label>
                <input 
                  v-model.number="bulkPrice" 
                  type="number" 
                  min="0" 
                  step="1000"
                  class="bulk-input"
                  placeholder="0"
                />
                <span class="currency-label">تومان</span>
                <button @click="applyBulkPrice" class="bulk-btn">
                  اعمال
                </button>
              </div>
            </div>
          </div>
          
          <div class="preview-section">
            <h4>پیش‌نمایش</h4>
            <div class="preview-controls">
              <label>انتخاب منطقه:</label>
              <select v-model="selectedRegion" class="preview-select">
                <option value="">انتخاب کنید</option>
                <option v-for="region in regions" :key="region.code" :value="region.code">
                  {{ region.name }}
                </option>
              </select>
            </div>
            <div v-if="selectedRegion" class="preview-result">
              <span>قیمت ارسال:</span>
              <strong>{{ getRegionPrice(selectedRegion) }} تومان</strong>
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
      regions: []
    })
  }
})

const emit = defineEmits(['update:modelValue'])

const enabled = ref(props.modelValue.enabled)
const regionType = ref(props.modelValue.regionType)
const regions = ref(props.modelValue.regions.length ? props.modelValue.regions : [
  { name: 'تهران', code: 'tehran', price: 50000, active: true },
  { name: 'اصفهان', code: 'isfahan', price: 60000, active: true },
  { name: 'شیراز', code: 'shiraz', price: 70000, active: true }
])
const selectedRegion = ref('')
const bulkPrice = ref(0)

watch([enabled, regionType, regions], () => {
  emit('update:modelValue', {
    enabled: enabled.value,
    regionType: regionType.value,
    regions: regions.value
  })
}, { deep: true })

function addRegion() {
  regions.value.push({
    name: '',
    code: '',
    price: 50000,
    active: true
  })
}

function removeRegion(index) {
  if (regions.value.length > 1) {
    regions.value.splice(index, 1)
  }
}

function applyBulkPrice() {
  if (bulkPrice.value > 0) {
    regions.value.forEach(region => {
      region.price = bulkPrice.value
    })
  }
}

function getRegionPrice(code) {
  const region = regions.value.find(r => r.code === code)
  if (region && region.active) {
    return new Intl.NumberFormat('fa-IR').format(region.price)
  }
  return 'تعریف نشده'
}
</script>

<style scoped>
.regional-pricing {
  direction: rtl;
}

.pricing-card {
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

.form-group {
  margin-bottom: 1.5rem;
}

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: 500;
  color: #555;
}

.region-type-select {
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 8px;
  background: white;
  min-width: 120px;
}

.regions-section {
  margin-bottom: 1.5rem;
}

.table-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}

.table-header h4 {
  margin: 0;
  font-size: 1rem;
  color: #555;
}

.add-btn {
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
}

.add-btn:hover {
  background: #218838;
}

.regions-table {
  background: white;
  border-radius: 8px;
  border: 1px solid #e9ecef;
  overflow: hidden;
}

.table-row {
  display: grid;
  grid-template-columns: 1.5fr 1fr 1.5fr 1fr 100px;
  gap: 1rem;
  padding: 1rem;
  align-items: center;
}

.table-row.header {
  background: #f8f9fa;
  font-weight: 600;
  color: #555;
  border-bottom: 1px solid #e9ecef;
}

.table-row:not(.header) {
  border-bottom: 1px solid #f1f3f4;
}

.table-row:last-child {
  border-bottom: none;
}

.region-input {
  width: 100%;
  padding: 0.5rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 0.9rem;
  text-align: right;
}

.price-input-group {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.price-input {
  flex: 1;
  padding: 0.5rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 0.9rem;
  text-align: left;
}

.currency-label {
  font-size: 0.8rem;
  color: #666;
  white-space: nowrap;
}

.status-select {
  padding: 0.5rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  background: white;
  font-size: 0.9rem;
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

.bulk-actions {
  margin-bottom: 1.5rem;
  padding: 1rem;
  background: #f8f9fa;
  border-radius: 8px;
  border: 1px solid #e9ecef;
}

.bulk-actions h4 {
  margin: 0 0 1rem 0;
  font-size: 1rem;
  color: #555;
}

.bulk-controls {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.bulk-input-group {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.bulk-input-group label {
  font-weight: 500;
  color: #555;
  white-space: nowrap;
}

.bulk-input {
  width: 120px;
  padding: 0.5rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  text-align: left;
}

.bulk-btn {
  background: #007bff;
  color: white;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 6px;
  cursor: pointer;
  font-size: 0.9rem;
}

.bulk-btn:hover {
  background: #0056b3;
}

.preview-section {
  margin-top: 1.5rem;
  padding-top: 1.5rem;
  border-top: 1px solid #e9ecef;
}

.preview-section h4 {
  margin: 0 0 1rem 0;
  font-size: 1rem;
  color: #666;
}

.preview-controls {
  display: flex;
  align-items: center;
  gap: 1rem;
  margin-bottom: 1rem;
}

.preview-controls label {
  font-weight: 500;
  color: #555;
}

.preview-select {
  padding: 0.5rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  background: white;
  min-width: 150px;
}

.preview-result {
  background: white;
  padding: 1rem;
  border-radius: 8px;
  border: 1px solid #e9ecef;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.preview-result strong {
  color: #2196F3;
  font-size: 1.1rem;
}
</style> 
