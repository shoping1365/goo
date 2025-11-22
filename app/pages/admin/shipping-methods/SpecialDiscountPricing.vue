<template>
  <div class="special-discount-pricing">
    <div class="pricing-card">
      <div class="card-header">
        <h3>تخفیف‌های خاص</h3>
        <div class="toggle-switch">
          <input id="discountToggle" v-model="enabled" type="checkbox" />
          <label for="discountToggle"></label>
        </div>
      </div>
      
      <div v-if="enabled" class="card-body">
        <div class="discount-list">
          <div class="list-header">
            <h4>کدهای تخفیف</h4>
            <button class="add-btn" @click="addDiscount">
              <span>+</span> افزودن تخفیف
            </button>
          </div>
          
          <div v-for="(discount, index) in discounts" :key="index" class="discount-item">
            <div class="discount-header">
              <h5>تخفیف {{ index + 1 }}</h5>
              <button class="remove-btn" @click="removeDiscount(index)">
                حذف
              </button>
            </div>
            
            <div class="discount-form">
              <div class="form-row">
                <div class="form-group">
                  <label>کد تخفیف</label>
                  <input 
                    v-model="discount.code" 
                    type="text" 
                    class="form-input"
                    placeholder="مثلاً SHIP20"
                  />
                </div>
                <div class="form-group">
                  <label>نوع تخفیف</label>
                  <select v-model="discount.type" class="form-select">
                    <option value="percentage">درصدی</option>
                    <option value="fixed">مبلغ ثابت</option>
                  </select>
                </div>
              </div>
              
              <div class="form-row">
                <div class="form-group">
                  <label>مقدار تخفیف</label>
                  <div class="discount-input-group">
                    <input 
                      v-model.number="discount.value" 
                      type="number" 
                      min="0" 
                      :max="discount.type === 'percentage' ? 100 : null"
                      step="1"
                      class="form-input"
                      placeholder="0"
                    />
                    <span class="unit-label">
                      {{ discount.type === 'percentage' ? '%' : 'تومان' }}
                    </span>
                  </div>
                </div>
                <div class="form-group">
                  <label>حداقل ارزش سفارش</label>
                  <input 
                    v-model.number="discount.minOrderValue" 
                    type="number" 
                    min="0" 
                    step="1000"
                    class="form-input"
                    placeholder="0"
                  />
                </div>
              </div>
              
              <div class="form-row">
                <div class="form-group">
                  <label>تاریخ شروع</label>
                  <input 
                    v-model="discount.startDate" 
                    type="date" 
                    class="form-input"
                  />
                </div>
                <div class="form-group">
                  <label>تاریخ پایان</label>
                  <input 
                    v-model="discount.endDate" 
                    type="date" 
                    class="form-input"
                  />
                </div>
              </div>
              
              <div class="form-group">
                <label>توضیحات</label>
                <textarea 
                  v-model="discount.description" 
                  class="form-textarea"
                  placeholder="توضیح مختصر درباره این تخفیف"
                ></textarea>
              </div>
            </div>
          </div>
        </div>
        
        <div class="preview-section">
          <h4>پیش‌نمایش تخفیف</h4>
          <div class="preview-controls">
            <div class="preview-row">
              <label>کد تخفیف:</label>
              <input 
                v-model="sampleCode" 
                type="text" 
                class="preview-input"
                placeholder="کد تخفیف را وارد کنید"
              />
            </div>
            <div class="preview-row">
              <label>ارزش سفارش:</label>
              <input 
                v-model.number="sampleOrderValue" 
                type="number" 
                min="0" 
                step="1000"
                class="preview-input"
                placeholder="0"
              />
            </div>
          </div>
          <div class="preview-result">
            <span>تخفیف اعمال شده:</span>
            <strong>{{ calculateDiscount(sampleCode, sampleOrderValue) }}</strong>
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
      discounts: []
    })
  }
})

const emit = defineEmits(['update:modelValue'])

const enabled = ref(props.modelValue.enabled)
const discounts = ref(props.modelValue.discounts.length ? props.modelValue.discounts : [
  {
    code: 'SHIP20',
    type: 'percentage',
    value: 20,
    minOrderValue: 500000,
    startDate: '',
    endDate: '',
    description: 'تخفیف 20 درصدی ارسال برای سفارش‌های بالای 500 هزار تومان'
  }
])
const sampleCode = ref('SHIP20')
const sampleOrderValue = ref(1000000)

watch([enabled, discounts], () => {
  emit('update:modelValue', {
    enabled: enabled.value,
    discounts: discounts.value
  })
}, { deep: true })

function addDiscount() {
  discounts.value.push({
    code: '',
    type: 'percentage',
    value: 0,
    minOrderValue: 0,
    startDate: '',
    endDate: '',
    description: ''
  })
}

function removeDiscount(index) {
  discounts.value.splice(index, 1)
}

function calculateDiscount(code, orderValue) {
  if (!code || !orderValue) return '0 تومان'
  
  const discount = discounts.value.find(d => d.code === code)
  if (!discount) return 'کد تخفیف نامعتبر'
  
  if (orderValue < discount.minOrderValue) {
    return `حداقل سفارش ${new Intl.NumberFormat('fa-IR').format(discount.minOrderValue)} تومان`
  }
  
  let discountAmount = 0
  if (discount.type === 'percentage') {
    discountAmount = (orderValue * discount.value) / 100
  } else {
    discountAmount = discount.value
  }
  
  return `${new Intl.NumberFormat('fa-IR').format(discountAmount)} تومان`
}
</script>

<style scoped>
.special-discount-pricing {
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

.discount-list {
  margin-bottom: 1.5rem;
}

.list-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}

.list-header h4 {
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

.discount-item {
  background: white;
  border-radius: 8px;
  border: 1px solid #e9ecef;
  margin-bottom: 1rem;
  overflow: hidden;
}

.discount-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem 1.5rem;
  background: #f8f9fa;
  border-bottom: 1px solid #e9ecef;
}

.discount-header h5 {
  margin: 0;
  font-size: 1rem;
  color: #333;
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

.discount-form {
  padding: 1.5rem;
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

.discount-input-group {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.discount-input-group .form-input {
  flex: 1;
}

.unit-label {
  font-size: 0.9rem;
  color: #666;
  white-space: nowrap;
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
  margin-bottom: 1rem;
}

.preview-row {
  display: flex;
  align-items: center;
  gap: 1rem;
  margin-bottom: 0.8rem;
}

.preview-row label {
  font-weight: 500;
  color: #555;
  min-width: 120px;
}

.preview-input {
  flex: 1;
  max-width: 200px;
  padding: 0.5rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  text-align: left;
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
  color: #28a745;
  font-size: 1.1rem;
}
</style> 
