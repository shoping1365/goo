<template>
  <div class="order-type-limitations">
    <div class="settings-card">
      <div class="card-header">
        <h3>محدودیت‌های نوع سفارش</h3>
        <div class="toggle-switch">
          <input type="checkbox" v-model="enabled" id="orderTypeToggle" />
          <label for="orderTypeToggle"></label>
        </div>
      </div>
      <div v-if="enabled" class="card-body">
        <div class="form-section">
          <h4>نوع‌های سفارش مجاز</h4>
          <div class="order-types">
            <label v-for="type in orderTypes" :key="type.value" class="order-type-checkbox">
              <input
                  type="checkbox"
                  v-model="allowedOrderTypes"
                  :value="type.value"
              />
              {{ type.label }}
            </label>
          </div>
        </div>
        <div class="form-section">
          <h4>نوع‌های سفارش ممنوعه</h4>
          <div class="prohibited-order-types">
            <label v-for="type in orderTypes" :key="type.value" class="order-type-checkbox">
              <input
                  type="checkbox"
                  v-model="prohibitedOrderTypes"
                  :value="type.value"
              />
              {{ type.label }}
            </label>
          </div>
        </div>
        <div class="form-section">
          <h4>پیام هشدار</h4>
          <div class="form-group">
            <textarea
                v-model="warningMessage"
                class="form-input"
                placeholder="پیام نمایشی در صورت انتخاب نوع سفارش غیرمجاز"
                rows="2"
            ></textarea>
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
    default: () => ({})
  }
})

const emit = defineEmits(['update:modelValue'])

const enabled = ref(props.modelValue.enabled ?? true)

const orderTypes = [
  { value: 'normal', label: 'عادی' },
  { value: 'express', label: 'فوری' },
  { value: 'scheduled', label: 'زمان‌بندی‌شده' },
  { value: 'bulk', label: 'گروهی' },
  { value: 'gift', label: 'هدیه' }
]

const allowedOrderTypes = ref(props.modelValue.allowedOrderTypes || ['normal', 'express'])
const prohibitedOrderTypes = ref(props.modelValue.prohibitedOrderTypes || [])
const warningMessage = ref(props.modelValue.warningMessage || 'نوع سفارش انتخاب شده مجاز نیست.')

watch([enabled, allowedOrderTypes, prohibitedOrderTypes, warningMessage], () => {
  emit('update:modelValue', {
    enabled: enabled.value,
    allowedOrderTypes: allowedOrderTypes.value,
    prohibitedOrderTypes: prohibitedOrderTypes.value,
    warningMessage: warningMessage.value
  })
}, { deep: true })
</script>

<style scoped>
.order-type-limitations {
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  overflow: hidden;
}
.settings-card {
  background: #fff;
}
.card-header {
  background: linear-gradient(135deg, #6f42c1 0%, #563d7c 100%);
  color: white;
  padding: 1.5rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.card-header h3 {
  margin: 0;
  font-size: 1.3rem;
}
.toggle-switch {
  position: relative;
  width: 60px;
  height: 30px;
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
  background-color: rgba(255, 255, 255, 0.3);
  transition: 0.4s;
  border-radius: 30px;
}
.toggle-switch label:before {
  position: absolute;
  content: "";
  height: 22px;
  width: 22px;
  left: 4px;
  bottom: 4px;
  background-color: white;
  transition: 0.4s;
  border-radius: 50%;
}
.toggle-switch input:checked + label {
  background-color: #28a745;
}
.toggle-switch input:checked + label:before {
  transform: translateX(30px);
}
.card-body {
  padding: 2rem;
}
.form-section {
  margin-bottom: 2rem;
  padding: 1.5rem;
  background: #f8f9fa;
  border-radius: 8px;
  border: 1px solid #e9ecef;
}
.form-section h4 {
  margin: 0 0 1.5rem 0;
  color: #333;
  font-size: 1.1rem;
  border-bottom: 2px solid #6f42c1;
  padding-bottom: 0.5rem;
}
.order-types,
.prohibited-order-types {
  display: flex;
  flex-wrap: wrap;
  gap: 1rem;
}
.order-type-checkbox {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  cursor: pointer;
  padding: 0.5rem;
  border-radius: 6px;
  transition: background-color 0.3s;
  background: white;
  border: 1px solid #e9ecef;
}
.order-type-checkbox:hover {
  background: #e9ecef;
}
.order-type-checkbox input[type="checkbox"] {
  width: 18px;
  height: 18px;
}
.form-group {
  margin-bottom: 1rem;
}
.form-input {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 0.9rem;
  font-family: inherit;
  transition: border-color 0.3s;
}
.form-input:focus {
  outline: none;
  border-color: #6f42c1;
  box-shadow: 0 0 0 2px rgba(111, 66, 193, 0.25);
}
@media (max-width: 768px) {
  .order-types,
  .prohibited-order-types {
    flex-direction: column;
    gap: 0.5rem;
  }
  .card-body {
    padding: 1rem;
  }
  .form-section {
    padding: 1rem;
  }
}
</style>
