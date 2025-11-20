<template>
  <div class="special-order-conditions">
    <div class="settings-card">
      <div class="card-header">
        <h3>شرایط خاص سفارش</h3>
        <div class="toggle-switch">
          <input id="specialOrderToggle" v-model="enabled" type="checkbox" />
          <label for="specialOrderToggle"></label>
        </div>
      </div>
      <div v-if="enabled" class="card-body">
        <!-- سفارش فوری -->
        <div class="form-section">
          <h4>سفارش فوری</h4>
          <label class="checkbox-label">
            <input v-model="express.enabled" type="checkbox" />
            فعال‌سازی سفارش فوری
          </label>
          <div v-if="express.enabled" class="express-options">
            <div class="form-row">
              <div class="form-group">
                <label>حداقل زمان قبل از ارسال (ساعت)</label>
                <input
                    v-model.number="express.minPreparationHours"
                    type="number"
                    min="0"
                    step="0.5"
                    class="form-input"
                    placeholder="0"
                />
              </div>
              <div class="form-group">
                <label>هزینه اضافی (تومان)</label>
                <input
                    v-model.number="express.additionalCost"
                    type="number"
                    min="0"
                    step="1000"
                    class="form-input"
                    placeholder="0"
                />
              </div>
            </div>
            <div class="form-group">
              <label>پیام سفارش فوری</label>
              <textarea
                  v-model="express.message"
                  class="form-input"
                  placeholder="پیام نمایشی برای سفارش فوری"
                  rows="2"
              ></textarea>
            </div>
          </div>
        </div>
        <!-- سفارش گروهی -->
        <div class="form-section">
          <h4>سفارش گروهی</h4>
          <label class="checkbox-label">
            <input v-model="bulk.enabled" type="checkbox" />
            فعال‌سازی سفارش گروهی
          </label>
          <div v-if="bulk.enabled" class="bulk-options">
            <div class="form-row">
              <div class="form-group">
                <label>حداقل تعداد برای سفارش گروهی</label>
                <input
                    v-model.number="bulk.minQuantity"
                    type="number"
                    min="2"
                    step="1"
                    class="form-input"
                    placeholder="10"
                />
              </div>
              <div class="form-group">
                <label>تخفیف سفارش گروهی (%)</label>
                <input
                    v-model.number="bulk.discountPercentage"
                    type="number"
                    min="0"
                    max="100"
                    step="1"
                    class="form-input"
                    placeholder="0"
                />
              </div>
            </div>
            <div class="form-group">
              <label>پیام سفارش گروهی</label>
              <textarea
                  v-model="bulk.message"
                  class="form-input"
                  placeholder="پیام نمایشی برای سفارش گروهی"
                  rows="2"
              ></textarea>
            </div>
          </div>
        </div>
        <!-- سفارش تکرارشونده -->
        <div class="form-section">
          <h4>سفارش تکرارشونده</h4>
          <label class="checkbox-label">
            <input v-model="recurring.enabled" type="checkbox" />
            فعال‌سازی سفارش تکرارشونده
          </label>
          <div v-if="recurring.enabled" class="recurring-options">
            <div class="form-row">
              <div class="form-group">
                <label>دوره تکرار (روز)</label>
                <input
                    v-model.number="recurring.intervalDays"
                    type="number"
                    min="1"
                    step="1"
                    class="form-input"
                    placeholder="30"
                />
              </div>
              <div class="form-group">
                <label>حداکثر تکرار</label>
                <input
                    v-model.number="recurring.maxRepeats"
                    type="number"
                    min="1"
                    step="1"
                    class="form-input"
                    placeholder="12"
                />
              </div>
            </div>
            <div class="form-group">
              <label>پیام سفارش تکرارشونده</label>
              <textarea
                  v-model="recurring.message"
                  class="form-input"
                  placeholder="پیام نمایشی برای سفارش تکرارشونده"
                  rows="2"
              ></textarea>
            </div>
          </div>
        </div>
        <!-- سفارش هدیه -->
        <div class="form-section">
          <h4>سفارش هدیه</h4>
          <label class="checkbox-label">
            <input v-model="gift.enabled" type="checkbox" />
            فعال‌سازی سفارش هدیه
          </label>
          <div v-if="gift.enabled" class="gift-options">
            <div class="form-group">
              <label>پیام سفارش هدیه</label>
              <textarea
                  v-model="gift.message"
                  class="form-input"
                  placeholder="پیام نمایشی برای سفارش هدیه"
                  rows="2"
              ></textarea>
            </div>
          </div>
        </div>
        <!-- پیام کلی شرایط خاص -->
        <div class="form-section">
          <h4>پیام کلی شرایط خاص</h4>
          <div class="form-group">
            <textarea
                v-model="generalMessage"
                class="form-input"
                placeholder="پیام نمایشی کلی برای شرایط خاص سفارش"
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

const express = ref({
  enabled: props.modelValue.express?.enabled ?? false,
  minPreparationHours: props.modelValue.express?.minPreparationHours || 0,
  additionalCost: props.modelValue.express?.additionalCost || 0,
  message: props.modelValue.express?.message || ''
})

const bulk = ref({
  enabled: props.modelValue.bulk?.enabled ?? false,
  minQuantity: props.modelValue.bulk?.minQuantity || 10,
  discountPercentage: props.modelValue.bulk?.discountPercentage || 0,
  message: props.modelValue.bulk?.message || ''
})

const recurring = ref({
  enabled: props.modelValue.recurring?.enabled ?? false,
  intervalDays: props.modelValue.recurring?.intervalDays || 30,
  maxRepeats: props.modelValue.recurring?.maxRepeats || 12,
  message: props.modelValue.recurring?.message || ''
})

const gift = ref({
  enabled: props.modelValue.gift?.enabled ?? false,
  message: props.modelValue.gift?.message || ''
})

const generalMessage = ref(props.modelValue.generalMessage || '')

watch([enabled, express, bulk, recurring, gift, generalMessage], () => {
  emit('update:modelValue', {
    enabled: enabled.value,
    express: express.value,
    bulk: bulk.value,
    recurring: recurring.value,
    gift: gift.value,
    generalMessage: generalMessage.value
  })
}, { deep: true })
</script>

<style scoped>
.special-order-conditions {
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  overflow: hidden;
}
.settings-card {
  background: #fff;
}
.card-header {
  background: linear-gradient(135deg, #fd7e14 0%, #e8590c 100%);
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
  border-bottom: 2px solid #fd7e14;
  padding-bottom: 0.5rem;
}
.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1.5rem;
  margin-bottom: 1rem;
}
.form-group {
  margin-bottom: 1rem;
}
.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  color: #333;
  font-weight: 500;
  font-size: 0.9rem;
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
  border-color: #fd7e14;
  box-shadow: 0 0 0 2px rgba(253, 126, 20, 0.25);
}
.checkbox-label {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  cursor: pointer;
  margin-bottom: 1rem;
  font-weight: 500;
}
.checkbox-label input[type="checkbox"] {
  width: 18px;
  height: 18px;
}
.express-options,
.bulk-options,
.recurring-options,
.gift-options {
  background: white;
  border: 1px solid #e9ecef;
  border-radius: 8px;
  padding: 1rem;
  margin-top: 1rem;
}
@media (max-width: 768px) {
  .form-row {
    grid-template-columns: 1fr;
  }
  .card-body {
    padding: 1rem;
  }
  .form-section {
    padding: 1rem;
  }
}
</style>
