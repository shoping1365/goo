<template>
  <div class="flat-rate-pricing">
    <div class="pricing-card">
      <div class="card-header">
        <h3>قیمت ثابت</h3>
        <div class="toggle-switch">
          <input id="flatRateToggle" v-model="enabled" type="checkbox" />
          <label for="flatRateToggle"></label>
        </div>
      </div>
      
      <div v-if="enabled" class="card-body">
        <div class="form-group">
          <label>مقدار قیمت ثابت</label>
          <div class="price-input-group">
            <input 
              v-model.number="price" 
              type="number" 
              min="0" 
              step="1000" 
              class="price-input"
              placeholder="0"
            />
            <select v-model="currency" class="currency-select">
              <option value="IRR">تومان</option>
              <option value="USD">دلار</option>
              <option value="EUR">یورو</option>
            </select>
          </div>
        </div>
        
        <div class="preview-section">
          <h4>پیش‌نمایش</h4>
          <div class="preview-card">
            <span>قیمت ارسال:</span>
            <strong>{{ formatPrice(price) }} {{ getCurrencyLabel(currency) }}</strong>
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
      price: 0,
      currency: 'IRR'
    })
  }
})

const emit = defineEmits(['update:modelValue'])

const enabled = ref(props.modelValue.enabled)
const price = ref(props.modelValue.price)
const currency = ref(props.modelValue.currency)

watch([enabled, price, currency], () => {
  emit('update:modelValue', {
    enabled: enabled.value,
    price: price.value,
    currency: currency.value
  })
}, { deep: true })

function formatPrice(value) {
  return new Intl.NumberFormat('fa-IR').format(value)
}

function getCurrencyLabel(currency) {
  const labels = {
    'IRR': 'تومان',
    'USD': 'دلار',
    'EUR': 'یورو'
  }
  return labels[currency] || currency
}
</script>

<style scoped>
.flat-rate-pricing {
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

.price-input-group {
  display: flex;
  gap: 0.5rem;
}

.price-input {
  flex: 1;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 8px;
  font-size: 1rem;
  text-align: left;
}

.currency-select {
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 8px;
  background: white;
  min-width: 100px;
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

.preview-card {
  background: white;
  padding: 1rem;
  border-radius: 8px;
  border: 1px solid #e9ecef;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.preview-card strong {
  color: #2196F3;
  font-size: 1.1rem;
}
</style> 
