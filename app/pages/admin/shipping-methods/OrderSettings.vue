<template>
  <div class="order-settings">
    <div class="settings-card">
      <div class="card-header">
        <h3>تنظیمات سفارش</h3>
        <div class="toggle-switch">
          <input type="checkbox" v-model="enabled" id="orderToggle" />
          <label for="orderToggle"></label>
        </div>
      </div>
      
      <div v-if="enabled" class="card-body">
        <div class="form-section">
          <h4>محدودیت‌های ارزش سفارش</h4>
          <div class="form-row">
            <div class="form-group">
              <label>حداقل ارزش سفارش</label>
              <div class="value-input-group">
                <input 
                  v-model.number="orderValueLimits.min" 
                  type="number" 
                  min="0" 
                  step="1000"
                  class="form-input"
                  placeholder="0"
                />
                <select v-model="currency" class="currency-select">
                  <option value="IRR">تومان</option>
                  <option value="USD">دلار</option>
                  <option value="EUR">یورو</option>
                </select>
              </div>
            </div>
            <div class="form-group">
              <label>حداکثر ارزش سفارش</label>
              <div class="value-input-group">
                <input 
                  v-model.number="orderValueLimits.max" 
                  type="number" 
                  min="0" 
                  step="1000"
                  class="form-input"
                  placeholder="0"
                />
                <select v-model="currency" class="currency-select">
                  <option value="IRR">تومان</option>
                  <option value="USD">دلار</option>
                  <option value="EUR">یورو</option>
                </select>
              </div>
            </div>
          </div>
        </div>
        
        <div class="form-section">
          <h4>محدودیت‌های تعداد آیتم</h4>
          <div class="form-row">
            <div class="form-group">
              <label>حداقل تعداد آیتم</label>
              <input 
                v-model.number="itemCountLimits.min" 
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
                v-model.number="itemCountLimits.max" 
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
          <h4>محدودیت‌های نوع سفارش</h4>
          <div class="order-type-limitations">
            <div class="form-group">
              <label>نوع‌های سفارش مجاز</label>
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
            <div class="form-group">
              <label>نوع‌های سفارش ممنوعه</label>
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
          </div>
        </div>
        
        <div class="form-section">
          <h4>شرایط خاص سفارش</h4>
          <div class="special-conditions">
            <div class="form-group">
              <label>سفارش‌های فوری</label>
              <div class="express-orders">
                <label class="checkbox-label">
                  <input type="checkbox" v-model="expressOrders.enabled" />
                  فعال‌سازی سفارش‌های فوری
                </label>
                <div v-if="expressOrders.enabled" class="express-options">
                  <div class="form-row">
                    <div class="form-group">
                      <label>حداقل زمان قبل از ارسال (ساعت)</label>
                      <input 
                        v-model.number="expressOrders.minPreparationHours" 
                        type="number" 
                        min="0" 
                        step="0.5"
                        class="form-input"
                        placeholder="0"
                      />
                    </div>
                    <div class="form-group">
                      <label>هزینه اضافی</label>
                      <div class="price-input-group">
                        <input 
                          v-model.number="expressOrders.additionalCost" 
                          type="number" 
                          min="0" 
                          step="1000"
                          class="form-input"
                          placeholder="0"
                        />
                        <span class="currency-label">تومان</span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            
            <div class="form-group">
              <label>سفارش‌های گروهی</label>
              <div class="bulk-orders">
                <label class="checkbox-label">
                  <input type="checkbox" v-model="bulkOrders.enabled" />
                  فعال‌سازی سفارش‌های گروهی
                </label>
                <div v-if="bulkOrders.enabled" class="bulk-options">
                  <div class="form-row">
                    <div class="form-group">
                      <label>حداقل تعداد برای سفارش گروهی</label>
                      <input 
                        v-model.number="bulkOrders.minQuantity" 
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
                        v-model.number="bulkOrders.discountPercentage" 
                        type="number" 
                        min="0" 
                        max="100"
                        step="1"
                        class="form-input"
                        placeholder="0"
                      />
                    </div>
                  </div>
                </div>
              </div>
            </div>
            
            <div class="form-group">
              <label>سفارش‌های تکرار شونده</label>
              <div class="recurring-orders">
                <label class="checkbox-label">
                  <input type="checkbox" v-model="recurringOrders.enabled" />
                  فعال‌سازی سفارش‌های تکرار شونده
                </label>
                <div v-if="recurringOrders.enabled" class="recurring-options">
                  <div class="form-row">
                    <div class="form-group">
                      <label>حداقل فاصله بین سفارش‌ها (روز)</label>
                      <input 
                        v-model.number="recurringOrders.minIntervalDays" 
                        type="number" 
                        min="1" 
                        step="1"
                        class="form-input"
                        placeholder="7"
                      />
                    </div>
                    <div class="form-group">
                      <label>تخفیف سفارش تکرار شونده (%)</label>
                      <input 
                        v-model.number="recurringOrders.discountPercentage" 
                        type="number" 
                        min="0" 
                        max="100"
                        step="1"
                        class="form-input"
                        placeholder="0"
                      />
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        
        <div class="form-section">
          <h4>شرایط اضافی</h4>
          <div class="additional-conditions">
            <div class="form-group">
              <label>شرایط خاص سفارش</label>
              <textarea 
                v-model="specialConditions" 
                class="form-textarea"
                placeholder="شرایط و محدودیت‌های خاص برای سفارش‌ها..."
              ></textarea>
            </div>
            <div class="form-group">
              <label>پیام‌های هشدار</label>
              <textarea 
                v-model="warningMessages" 
                class="form-textarea"
                placeholder="پیام‌های هشدار برای مشتریان..."
              ></textarea>
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
      currency: 'IRR',
      orderValueLimits: { min: 0, max: 0 },
      itemCountLimits: { min: 1, max: 0 },
      orderTypes: [
        { value: 'standard', label: 'سفارش عادی' },
        { value: 'express', label: 'سفارش فوری' },
        { value: 'scheduled', label: 'سفارش زمان‌بندی شده' },
        { value: 'pickup', label: 'دریافت از فروشگاه' },
        { value: 'bulk', label: 'سفارش گروهی' }
      ],
      allowedOrderTypes: ['standard', 'express'],
      prohibitedOrderTypes: [],
      expressOrders: {
        enabled: false,
        minPreparationHours: 2,
        additionalCost: 50000
      },
      bulkOrders: {
        enabled: false,
        minQuantity: 10,
        discountPercentage: 10
      },
      recurringOrders: {
        enabled: false,
        minIntervalDays: 7,
        discountPercentage: 5
      },
      specialConditions: '',
      warningMessages: ''
    })
  }
})

const emit = defineEmits(['update:modelValue'])

const enabled = ref(props.modelValue.enabled)
const currency = ref(props.modelValue.currency)
const orderValueLimits = ref(props.modelValue.orderValueLimits)
const itemCountLimits = ref(props.modelValue.itemCountLimits)
const orderTypes = ref(props.modelValue.orderTypes)
const allowedOrderTypes = ref(props.modelValue.allowedOrderTypes)
const prohibitedOrderTypes = ref(props.modelValue.prohibitedOrderTypes)
const expressOrders = ref(props.modelValue.expressOrders)
const bulkOrders = ref(props.modelValue.bulkOrders)
const recurringOrders = ref(props.modelValue.recurringOrders)
const specialConditions = ref(props.modelValue.specialConditions)
const warningMessages = ref(props.modelValue.warningMessages)

watch([enabled, currency, orderValueLimits, itemCountLimits, allowedOrderTypes, prohibitedOrderTypes, expressOrders, bulkOrders, recurringOrders, specialConditions, warningMessages], () => {
  emit('update:modelValue', {
    enabled: enabled.value,
    currency: currency.value,
    orderValueLimits: orderValueLimits.value,
    itemCountLimits: itemCountLimits.value,
    orderTypes: orderTypes.value,
    allowedOrderTypes: allowedOrderTypes.value,
    prohibitedOrderTypes: prohibitedOrderTypes.value,
    expressOrders: expressOrders.value,
    bulkOrders: bulkOrders.value,
    recurringOrders: recurringOrders.value,
    specialConditions: specialConditions.value,
    warningMessages: warningMessages.value
  })
}, { deep: true })
</script>

<style scoped>
.order-settings {
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

.form-input, .form-select, .form-textarea {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 0.9rem;
  background: #fafbfc;
}

.form-input:focus, .form-select:focus, .form-textarea:focus {
  border-color: #2196F3;
  outline: none;
}

.form-textarea {
  resize: vertical;
  min-height: 80px;
}

.value-input-group, .price-input-group {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.value-input-group .form-input, .price-input-group .form-input {
  flex: 1;
}

.currency-select {
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  background: white;
  min-width: 100px;
}

.currency-label {
  font-size: 0.9rem;
  color: #666;
  white-space: nowrap;
}

.order-types, .prohibited-order-types {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 0.5rem;
  margin-top: 0.5rem;
}

.order-type-checkbox {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-size: 0.9rem;
  cursor: pointer;
  padding: 0.5rem;
  background: white;
  border-radius: 4px;
  border: 1px solid #e9ecef;
}

.order-type-checkbox input[type="checkbox"] {
  width: 16px;
  height: 16px;
}

.checkbox-label {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  cursor: pointer;
  font-weight: 500;
  color: #333;
}

.checkbox-label input[type="checkbox"] {
  width: 16px;
  height: 16px;
}

.express-options, .bulk-options, .recurring-options {
  margin-top: 1rem;
  padding: 1rem;
  background: #f8f9fa;
  border-radius: 6px;
  border: 1px solid #e9ecef;
}

.additional-conditions {
  background: white;
  border-radius: 8px;
  border: 1px solid #e9ecef;
  padding: 1rem;
}
</style> 
