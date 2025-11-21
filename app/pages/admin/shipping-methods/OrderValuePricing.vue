<template>
     <div class="order-value-pricing">
       <div class="section-header">
         <h3>قیمت‌گذاری بر اساس ارزش سفارش</h3>
         <p>تعریف قیمت‌های ارسال بر اساس ارزش کل سفارش</p>
       </div>
   
       <!-- Pricing Rules List -->
       <div class="pricing-rules">
         <h4>قوانین قیمت‌گذاری</h4>
         <table>
           <thead>
             <tr>
               <th>از ارزش</th>
               <th>تا ارزش</th>
               <th>قیمت ارسال</th>
               <th>نوع محاسبه</th>
               <th>وضعیت</th>
               <th>عملیات</th>
             </tr>
           </thead>
           <tbody>
             <tr v-for="rule in pricingRules" :key="rule.id">
               <td>{{ formatCurrency(rule.minValue) }}</td>
               <td>{{ formatCurrency(rule.maxValue) }}</td>
               <td>{{ formatCurrency(rule.shippingCost) }}</td>
               <td>{{ rule.calculationType }}</td>
               <td>
                 <span :class="['status', rule.status]">{{ rule.statusText }}</span>
               </td>
               <td>
                 <button class="btn-edit" @click="editRule(rule)">ویرایش</button>
                 <button class="btn-delete" @click="deleteRule(rule.id)">حذف</button>
               </td>
             </tr>
           </tbody>
         </table>
       </div>
   
       <!-- Add New Rule -->
       <div class="add-rule-section">
         <h4>افزودن قانون جدید</h4>
         <form @submit.prevent="addNewRule">
           <div class="form-row">
             <div class="form-group">
               <label>از ارزش:</label>
               <input v-model.number="newRule.minValue" type="number" min="0" step="1000" required>
             </div>
             <div class="form-group">
               <label>تا ارزش:</label>
               <input v-model.number="newRule.maxValue" type="number" min="0" step="1000" required>
             </div>
           </div>
           <div class="form-row">
             <div class="form-group">
               <label>قیمت ارسال:</label>
               <input v-model.number="newRule.shippingCost" type="number" min="0" step="1000" required>
             </div>
             <div class="form-group">
               <label>نوع محاسبه:</label>
               <select v-model="newRule.calculationType" required>
                 <option value="fixed">ثابت</option>
                 <option value="percentage">درصدی</option>
                 <option value="tiered">پلکانی</option>
               </select>
             </div>
           </div>
           <button type="submit" class="btn-primary">افزودن قانون</button>
         </form>
       </div>
   
       <!-- Pricing Calculator -->
       <div class="pricing-calculator">
         <h4>محاسبه‌گر قیمت</h4>
         <div class="calculator-form">
           <div class="form-group">
             <label>ارزش سفارش:</label>
             <input v-model.number="calculator.orderValue" type="number" min="0" step="1000" @input="calculateShipping">
           </div>
           <div v-if="calculatedShipping" class="calculation-result">
             <h5>نتیجه محاسبه:</h5>
             <div class="result-item">
               <span>قیمت ارسال:</span>
               <span>{{ formatCurrency(calculatedShipping) }}</span>
             </div>
             <div class="result-item">
               <span>قانون اعمال شده:</span>
               <span>{{ appliedRule }}</span>
             </div>
           </div>
         </div>
       </div>
     </div>
</template>

<script setup>
import { reactive, ref } from 'vue'

const pricingRules = ref([
  {
    id: 1,
    minValue: 0,
    maxValue: 100000,
    shippingCost: 15000,
    calculationType: 'fixed',
    status: 'active',
    statusText: 'فعال'
  },
  {
    id: 2,
    minValue: 100000,
    maxValue: 500000,
    shippingCost: 10000,
    calculationType: 'fixed',
    status: 'active',
    statusText: 'فعال'
  },
  {
    id: 3,
    minValue: 500000,
    maxValue: 1000000,
    shippingCost: 5000,
    calculationType: 'fixed',
    status: 'active',
    statusText: 'فعال'
  }
])

const newRule = reactive({
  minValue: 0,
  maxValue: 0,
  shippingCost: 0,
  calculationType: 'fixed'
})

const calculator = reactive({
  orderValue: 0
})

const calculatedShipping = ref(0)
const appliedRule = ref('')

function formatCurrency(value) {
  return new Intl.NumberFormat('fa-IR', {
    style: 'currency',
    currency: 'IRR'
  }).format(value)
}

function editRule(_rule) {
}

function deleteRule(_id) {
}

function addNewRule() {
  // Reset form
  Object.assign(newRule, {
    minValue: 0,
    maxValue: 0,
    shippingCost: 0,
    calculationType: 'fixed'
  })
}

function calculateShipping() {
  const value = calculator.orderValue
  const rule = pricingRules.value.find(r => 
    value >= r.minValue && value <= r.maxValue && r.status === 'active'
  )
  
  if (rule) {
    calculatedShipping.value = rule.shippingCost
    appliedRule.value = `از ${formatCurrency(rule.minValue)} تا ${formatCurrency(rule.maxValue)}`
  } else {
    calculatedShipping.value = 0
    appliedRule.value = 'قانونی یافت نشد'
  }
}
</script>

<style scoped>
.order-value-pricing {
  direction: rtl;
  padding: 1rem;
}

.section-header {
  margin-bottom: 2rem;
}

.section-header h3 {
  margin: 0 0 0.5rem 0;
  font-size: 1.5rem;
  color: #333;
}

.section-header p {
  margin: 0;
  color: #666;
}

.pricing-rules {
  margin-bottom: 2rem;
}

.pricing-rules h4 {
  margin: 0 0 1rem 0;
  font-size: 1.2rem;
  color: #333;
}

.pricing-rules table {
  width: 100%;
  border-collapse: collapse;
  background: white;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.pricing-rules th,
.pricing-rules td {
  padding: 1rem;
  text-align: right;
  border-bottom: 1px solid #e9ecef;
}

.pricing-rules th {
  background: #f8f9fa;
  font-weight: 600;
  color: #555;
}

.status {
  padding: 0.25rem 0.5rem;
  border-radius: 12px;
  font-size: 0.8rem;
}

.status.active {
  background: #d4edda;
  color: #155724;
}

.status.inactive {
  background: #f8d7da;
  color: #721c24;
}

.btn-edit,
.btn-delete {
  padding: 0.25rem 0.5rem;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 0.8rem;
  margin-left: 0.5rem;
}

.btn-edit {
  background: #007bff;
  color: white;
}

.btn-delete {
  background: #dc3545;
  color: white;
}

.add-rule-section {
  background: #f8f9fa;
  border-radius: 12px;
  padding: 1.5rem;
  margin-bottom: 2rem;
}

.add-rule-section h4 {
  margin: 0 0 1rem 0;
  font-size: 1.2rem;
  color: #333;
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
}

.form-group input,
.form-group select {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 0.9rem;
  font-family: inherit;
}

.btn-primary {
  background: #007bff;
  color: white;
  border: none;
  padding: 0.75rem 1.5rem;
  border-radius: 6px;
  cursor: pointer;
  font-size: 0.9rem;
}

.pricing-calculator {
  background: white;
  border-radius: 12px;
  padding: 1.5rem;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.pricing-calculator h4 {
  margin: 0 0 1rem 0;
  font-size: 1.2rem;
  color: #333;
}

.calculator-form {
  max-width: 400px;
}

.calculation-result {
  margin-top: 1rem;
  padding: 1rem;
  background: #f8f9fa;
  border-radius: 8px;
}

.calculation-result h5 {
  margin: 0 0 0.5rem 0;
  color: #333;
}

.result-item {
  display: flex;
  justify-content: space-between;
  margin-bottom: 0.5rem;
}

.result-item:last-child {
  margin-bottom: 0;
}

@media (max-width: 768px) {
  .form-row {
    grid-template-columns: 1fr;
  }
  
  .pricing-rules table {
    font-size: 0.8rem;
  }
  
  .pricing-rules th,
  .pricing-rules td {
    padding: 0.5rem;
  }
}
</style>
