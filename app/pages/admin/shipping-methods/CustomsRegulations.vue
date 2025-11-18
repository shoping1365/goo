<template>
  <div class="customs-regulations">
    <div class="section-header">
      <h3>قوانین گمرکی</h3>
      <p>مدیریت قوانین و مقررات گمرکی برای ارسال بین‌المللی</p>
    </div>

    <!-- Customs Rules List -->
    <div class="customs-rules">
      <h4>لیست قوانین گمرکی</h4>
      <table>
        <thead>
        <tr>
          <th>کشور مقصد</th>
          <th>نوع کالا</th>
          <th>محدودیت مقدار</th>
          <th>محدودیت ارزش</th>
          <th>مستندات مورد نیاز</th>
          <th>وضعیت</th>
          <th>عملیات</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="rule in customsRules" :key="rule.id">
          <td>{{ rule.destinationCountry }}</td>
          <td>{{ rule.productType }}</td>
          <td>{{ rule.quantityLimit }}</td>
          <td>{{ rule.valueLimit }}</td>
          <td>{{ rule.requiredDocuments }}</td>
          <td>
            <span :class="['status', rule.status]">{{ rule.statusText }}</span>
          </td>
          <td>
            <button @click="editRule(rule)" class="btn-edit">ویرایش</button>
            <button @click="deleteRule(rule.id)" class="btn-delete">حذف</button>
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
            <label>کشور مقصد:</label>
            <select v-model="newRule.destinationCountry" required>
              <option value="">انتخاب کنید</option>
              <option value="usa">آمریکا</option>
              <option value="uk">انگلیس</option>
              <option value="germany">آلمان</option>
              <option value="france">فرانسه</option>
            </select>
          </div>
          <div class="form-group">
            <label>نوع کالا:</label>
            <select v-model="newRule.productType" required>
              <option value="">انتخاب کنید</option>
              <option value="electronics">الکترونیک</option>
              <option value="clothing">پوشاک</option>
              <option value="food">مواد غذایی</option>
              <option value="cosmetics">لوازم آرایشی</option>
            </select>
          </div>
        </div>
        <div class="form-row">
          <div class="form-group">
            <label>محدودیت مقدار:</label>
            <input v-model="newRule.quantityLimit" type="text" placeholder="مثال: حداکثر 5 عدد">
          </div>
          <div class="form-group">
            <label>محدودیت ارزش:</label>
            <input v-model="newRule.valueLimit" type="text" placeholder="مثال: حداکثر 1000 دلار">
          </div>
        </div>
        <div class="form-group">
          <label>مستندات مورد نیاز:</label>
          <textarea v-model="newRule.requiredDocuments" rows="3" placeholder="لیست مستندات مورد نیاز"></textarea>
        </div>
        <button type="submit" class="btn-primary">افزودن قانون</button>
      </form>
    </div>

    <!-- Customs Calculator -->
    <div class="customs-calculator">
      <h4>محاسبه‌گر گمرکی</h4>
      <div class="calculator-form">
        <div class="form-row">
          <div class="form-group">
            <label>کشور مقصد:</label>
            <select v-model="calculator.destinationCountry">
              <option value="">انتخاب کنید</option>
              <option value="usa">آمریکا</option>
              <option value="uk">انگلیس</option>
              <option value="germany">آلمان</option>
            </select>
          </div>
          <div class="form-group">
            <label>ارزش کالا:</label>
            <input v-model="calculator.productValue" type="number" placeholder="ارزش به دلار">
          </div>
        </div>
        <div class="form-row">
          <div class="form-group">
            <label>وزن کالا:</label>
            <input v-model="calculator.productWeight" type="number" placeholder="وزن به کیلوگرم">
          </div>
          <div class="form-group">
            <label>نوع کالا:</label>
            <select v-model="calculator.productType">
              <option value="">انتخاب کنید</option>
              <option value="electronics">الکترونیک</option>
              <option value="clothing">پوشاک</option>
              <option value="food">مواد غذایی</option>
            </select>
          </div>
        </div>
        <button @click="calculateCustoms" class="btn-calculate">محاسبه</button>
      </div>

      <div v-if="customsCalculation" class="calculation-result">
        <h5>نتیجه محاسبه:</h5>
        <div class="result-item">
          <span>حقوق گمرکی:</span>
          <span>{{ customsCalculation.duty }}</span>
        </div>
        <div class="result-item">
          <span>مالیات:</span>
          <span>{{ customsCalculation.tax }}</span>
        </div>
        <div class="result-item">
          <span>هزینه‌های اضافی:</span>
          <span>{{ customsCalculation.additionalFees }}</span>
        </div>
        <div class="result-item total">
          <span>مجموع:</span>
          <span>{{ customsCalculation.total }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'

const customsRules = ref([
  {
    id: 1,
    destinationCountry: 'آمریکا',
    productType: 'الکترونیک',
    quantityLimit: 'حداکثر 3 عدد',
    valueLimit: 'حداکثر 800 دلار',
    requiredDocuments: 'فاکتور، گواهی اصالت، مجوز واردات',
    status: 'active',
    statusText: 'فعال'
  },
  {
    id: 2,
    destinationCountry: 'انگلیس',
    productType: 'پوشاک',
    quantityLimit: 'حداکثر 10 عدد',
    valueLimit: 'حداکثر 500 پوند',
    requiredDocuments: 'فاکتور، گواهی منشا',
    status: 'active',
    statusText: 'فعال'
  }
])

const newRule = reactive({
  destinationCountry: '',
  productType: '',
  quantityLimit: '',
  valueLimit: '',
  requiredDocuments: ''
})

const calculator = reactive({
  destinationCountry: '',
  productValue: '',
  productWeight: '',
  productType: ''
})

const customsCalculation = ref(null)

const editRule = (rule) => {
  // Edit rule logic
}

const deleteRule = (id) => {
  // Delete rule logic
}

const addNewRule = () => {
  // Add new rule logic
}

const calculateCustoms = () => {
  // Calculate customs logic
  customsCalculation.value = {
    duty: '50 دلار',
    tax: '25 دلار',
    additionalFees: '15 دلار',
    total: '90 دلار'
  }
}
</script>

<style scoped>
.customs-regulations {
  padding: 20px;
}

.section-header {
  margin-bottom: 30px;
}

.section-header h3 {
  color: #2c3e50;
  margin-bottom: 10px;
}

.section-header p {
  color: #7f8c8d;
}

.customs-rules,
.add-rule-section,
.customs-calculator {
  background: white;
  border-radius: 8px;
  padding: 20px;
  margin-bottom: 20px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 15px;
}

th, td {
  padding: 12px;
  text-align: right;
  border-bottom: 1px solid #ecf0f1;
}

th {
  background-color: #f8f9fa;
  font-weight: 600;
  color: #2c3e50;
}

.status {
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 12px;
  font-weight: 500;
}

.status.active {
  background-color: #d4edda;
  color: #155724;
}

.status.inactive {
  background-color: #f8d7da;
  color: #721c24;
}

.btn-edit,
.btn-delete {
  padding: 6px 12px;
  border: none;
  border-radius: 4px;
  margin-left: 5px;
  cursor: pointer;
  font-size: 12px;
}

.btn-edit {
  background-color: #007bff;
  color: white;
}

.btn-delete {
  background-color: #dc3545;
  color: white;
}

.form-row {
  display: flex;
  gap: 20px;
  margin-bottom: 15px;
}

.form-group {
  flex: 1;
}

.form-group label {
  display: block;
  margin-bottom: 5px;
  font-weight: 500;
  color: #2c3e50;
}

.form-group input,
.form-group select,
.form-group textarea {
  width: 100%;
  padding: 8px 12px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 14px;
}

.btn-primary,
.btn-calculate {
  background-color: #28a745;
  color: white;
  border: none;
  padding: 10px 20px;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
}

.btn-primary:hover,
.btn-calculate:hover {
  background-color: #218838;
}

.calculator-form {
  margin-bottom: 20px;
}

.calculation-result {
  background-color: #f8f9fa;
  padding: 15px;
  border-radius: 4px;
  border: 1px solid #dee2e6;
}

.calculation-result h5 {
  margin-bottom: 15px;
  color: #2c3e50;
}

.result-item {
  display: flex;
  justify-content: space-between;
  margin-bottom: 8px;
  padding: 5px 0;
  border-bottom: 1px solid #dee2e6;
}

.result-item.total {
  font-weight: bold;
  border-bottom: none;
  margin-top: 10px;
  padding-top: 10px;
  border-top: 2px solid #007bff;
}
</style>
