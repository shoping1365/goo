<template>
  <div class="item-count-limits">
    <div class="settings-card">
      <div class="card-header">
        <h3>محدودیت تعداد آیتم در سفارش</h3>
        <div class="toggle-switch">
          <input id="itemCountToggle" v-model="enabled" type="checkbox" />
          <label for="itemCountToggle"></label>
        </div>
      </div>
      <div v-if="enabled" class="card-body">
        <div class="form-section">
          <h4>تعداد آیتم مجاز در هر سفارش</h4>
          <div class="form-row">
            <div class="form-group">
              <label>حداقل تعداد آیتم</label>
              <input
                  v-model.number="itemCount.min"
                  type="number"
                  min="1"
                  step="1"
                  class="form-input"
                  placeholder="1"
              />
            </div>
            <div class="form-group">
              <label>حداکثر تعداد آیتم</label>
              <input
                  v-model.number="itemCount.max"
                  type="number"
                  min="1"
                  step="1"
                  class="form-input"
                  placeholder="0"
              />
            </div>
          </div>
        </div>
        <div class="form-section">
          <h4>تنظیمات اضافی</h4>
          <div class="form-group">
            <label>پیام هشدار</label>
            <textarea
                v-model="itemCount.warningMessage"
                class="form-input"
                placeholder="پیام نمایشی در صورت نقض محدودیت"
                rows="2"
            ></textarea>
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
    default: () => ({})
  }
})

const emit = defineEmits(['update:modelValue'])

const enabled = ref(props.modelValue.enabled ?? true)

const itemCount = ref({
  min: props.modelValue.min || 1,
  max: props.modelValue.max || 100,
  warningMessage: props.modelValue.warningMessage || 'تعداد آیتم سفارش باید در بازه مجاز باشد.'
})

watch([enabled, itemCount], () => {
  emit('update:modelValue', {
    enabled: enabled.value,
    ...itemCount.value
  })
}, { deep: true })
</script>

<style scoped>
.item-count-limits {
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  overflow: hidden;
}
.settings-card {
  background: #fff;
}
.card-header {
  background: linear-gradient(135deg, #007bff 0%, #0056b3 100%);
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
  border-bottom: 2px solid #007bff;
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
  border-color: #007bff;
  box-shadow: 0 0 0 2px rgba(0, 123, 255, 0.25);
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
