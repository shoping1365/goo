<template>
  <div class="compensation-management">
    <div class="section-header">
      <h3>جبران خسارت</h3>
      <p>مدیریت جبران خسارت برای تاخیرها و مشکلات ارسال</p>
    </div>

    <!-- Compensation Rules -->
    <div class="compensation-rules">
      <h4>قوانین جبران خسارت</h4>
      <table>
        <thead>
        <tr>
          <th>عنوان قانون</th>
          <th>نوع مشکل</th>
          <th>مقدار جبران</th>
          <th>شرط فعال‌سازی</th>
          <th>وضعیت</th>
          <th>عملیات</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="rule in compensationRules" :key="rule.id">
          <td>{{ rule.title }}</td>
          <td>{{ rule.issueType }}</td>
          <td>{{ rule.compensationAmount }}</td>
          <td>{{ rule.activationCondition }}</td>
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
            <label>عنوان قانون:</label>
            <input v-model="newRule.title" type="text" required>
          </div>
          <div class="form-group">
            <label>نوع مشکل:</label>
            <select v-model="newRule.issueType" required>
              <option value="delay">تاخیر در ارسال</option>
              <option value="damage">آسیب به کالا</option>
              <option value="loss">گم شدن کالا</option>
              <option value="wrong_delivery">تحویل اشتباه</option>
            </select>
          </div>
        </div>
        <div class="form-row">
          <div class="form-group">
            <label>مقدار جبران:</label>
            <input v-model="newRule.compensationAmount" type="text" placeholder="مثال: 10% تخفیف">
          </div>
          <div class="form-group">
            <label>شرط فعال‌سازی:</label>
            <input v-model="newRule.activationCondition" type="text" placeholder="مثال: تاخیر بیش از 2 روز">
          </div>
        </div>
        <button type="submit" class="btn-primary">افزودن قانون</button>
      </form>
    </div>

    <!-- Compensation History -->
    <div class="compensation-history">
      <h4>تاریخچه جبران خسارت</h4>
      <table>
        <thead>
        <tr>
          <th>شماره سفارش</th>
          <th>مشکل</th>
          <th>مقدار جبران</th>
          <th>تاریخ</th>
          <th>وضعیت</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="compensation in compensationHistory" :key="compensation.id">
          <td>{{ compensation.orderNumber }}</td>
          <td>{{ compensation.issue }}</td>
          <td>{{ compensation.amount }}</td>
          <td>{{ compensation.date }}</td>
          <td>
            <span :class="['status', compensation.status]">{{ compensation.statusText }}</span>
          </td>
        </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'

const compensationRules = ref([
  {
    id: 1,
    title: 'جبران تاخیر 2 روزه',
    issueType: 'تاخیر در ارسال',
    compensationAmount: '10% تخفیف',
    activationCondition: 'تاخیر بیش از 2 روز',
    status: 'active',
    statusText: 'فعال'
  },
  {
    id: 2,
    title: 'جبران آسیب کالا',
    issueType: 'آسیب به کالا',
    compensationAmount: '100% بازگشت وجه',
    activationCondition: 'آسیب قابل مشاهده',
    status: 'active',
    statusText: 'فعال'
  }
])

const compensationHistory = ref([
  {
    id: 1,
    orderNumber: 'ORD-001',
    issue: 'تاخیر در ارسال',
    amount: '10% تخفیف',
    date: '1402/10/15',
    status: 'completed',
    statusText: 'تکمیل شده'
  }
])

const newRule = reactive({
  title: '',
  issueType: '',
  compensationAmount: '',
  activationCondition: ''
})

const editRule = (rule) => {
  // Edit rule logic
}

const deleteRule = (id) => {
  // Delete rule logic
}

const addNewRule = () => {
  // Add new rule logic
}
</script>

<style scoped>
.compensation-management {
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

.compensation-rules,
.add-rule-section,
.compensation-history {
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

.status.completed {
  background-color: #d1ecf1;
  color: #0c5460;
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
.form-group select {
  width: 100%;
  padding: 8px 12px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 14px;
}

.btn-primary {
  background-color: #28a745;
  color: white;
  border: none;
  padding: 10px 20px;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
}

.btn-primary:hover {
  background-color: #218838;
}
</style>
