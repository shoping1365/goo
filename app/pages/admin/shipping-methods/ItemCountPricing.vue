<template>
  <div class="item-count-pricing">
    <div class="pricing-card">
      <div class="card-header">
        <h3>قیمت بر اساس تعداد آیتم</h3>
        <div class="toggle-switch">
          <input id="itemCountToggle" v-model="enabled" type="checkbox" />
          <label for="itemCountToggle"></label>
        </div>
      </div>
      
      <div v-if="enabled" class="card-body">
        <div class="table-section">
          <div class="table-header">
            <h4>بازه‌های تعداد آیتم</h4>
            <button class="add-btn" @click="addItemRange">
              <span>+</span> افزودن بازه
            </button>
          </div>
          
          <div class="item-table">
            <div class="table-row header">
              <div class="col">از</div>
              <div class="col">تا</div>
              <div class="col">قیمت</div>
              <div class="col actions">عملیات</div>
            </div>
            
            <div v-for="(range, index) in itemRanges" :key="index" class="table-row">
              <div class="col">
                <input 
                  v-model.number="range.from" 
                  type="number" 
                  min="1" 
                  step="1"
                  class="range-input"
                  placeholder="1"
                />
              </div>
              <div class="col">
                <input 
                  v-model.number="range.to" 
                  type="number" 
                  min="1" 
                  step="1"
                  class="range-input"
                  :placeholder="`${range.from + 2}`"
                />
              </div>
              <div class="col">
                <div class="price-input-group">
                  <input 
                    v-model.number="range.price" 
                    type="number" 
                    min="0" 
                    step="1000"
                    class="price-input"
                    placeholder="0"
                  />
                  <span class="currency-label">تومان</span>
                </div>
              </div>
              <div class="col actions">
                <button class="remove-btn" @click="removeItemRange(index)">
                  حذف
                </button>
              </div>
            </div>
          </div>
        </div>
        
        <div class="preview-section">
          <h4>پیش‌نمایش محاسبه</h4>
          <div class="preview-controls">
            <label>تعداد آیتم نمونه:</label>
            <div class="preview-input-group">
              <input 
                v-model.number="sampleItemCount" 
                type="number" 
                min="1" 
                step="1"
                class="preview-input"
                placeholder="1"
              />
              <span class="unit-label">عدد</span>
            </div>
          </div>
          <div class="preview-result">
            <span>قیمت ارسال:</span>
            <strong>{{ calculatePrice(sampleItemCount) }} تومان</strong>
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
      itemRanges: []
    })
  }
})

const emit = defineEmits(['update:modelValue'])

const enabled = ref(props.modelValue.enabled)
const itemRanges = ref(props.modelValue.itemRanges.length ? props.modelValue.itemRanges : [
  { from: 1, to: 3, price: 25000 },
  { from: 4, to: 10, price: 40000 },
  { from: 11, to: 20, price: 60000 }
])
const sampleItemCount = ref(5)

watch([enabled, itemRanges], () => {
  emit('update:modelValue', {
    enabled: enabled.value,
    itemRanges: itemRanges.value
  })
}, { deep: true })

function addItemRange() {
  const lastRange = itemRanges.value[itemRanges.value.length - 1]
  const newFrom = lastRange ? lastRange.to + 1 : 1
  itemRanges.value.push({
    from: newFrom,
    to: newFrom + 9,
    price: lastRange ? lastRange.price + 10000 : 25000
  })
}

function removeItemRange(index) {
  if (itemRanges.value.length > 1) {
    itemRanges.value.splice(index, 1)
  }
}

function calculatePrice(itemCount) {
  if (!itemCount || itemCount <= 0) return 0
  
  const range = itemRanges.value.find(r => 
    itemCount >= r.from && itemCount <= r.to
  )
  
  if (range) {
    return new Intl.NumberFormat('fa-IR').format(range.price)
  }
  
  return 'تعریف نشده'
}
</script>

<style scoped>
.item-count-pricing {
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

.table-section {
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

.item-table {
  background: white;
  border-radius: 8px;
  border: 1px solid #e9ecef;
  overflow: hidden;
}

.table-row {
  display: grid;
  grid-template-columns: 1fr 1fr 1.5fr 100px;
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

.range-input {
  width: 100%;
  padding: 0.5rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 0.9rem;
  text-align: left;
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

.preview-input-group {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.preview-input {
  width: 100px;
  padding: 0.5rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  text-align: left;
}

.unit-label {
  font-size: 0.9rem;
  color: #666;
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
