<template>
  <div class="selection-guide">
    <div class="section-header">
      <h3>راهنمای انتخاب</h3>
      <p>راهنمای انتخاب بهترین روش ارسال برای مشتریان</p>
    </div>

    <!-- Selection Criteria -->
    <div class="selection-criteria">
      <h4>معیارهای انتخاب</h4>
      <div class="criteria-grid">
        <div v-for="criterion in criteria" :key="criterion.id" class="criterion-card">
          <div class="criterion-icon">
            <i :class="criterion.icon"></i>
          </div>
          <div class="criterion-content">
            <h5>{{ criterion.title }}</h5>
            <p>{{ criterion.description }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Selection Rules -->
    <div class="selection-rules">
      <h4>قوانین انتخاب</h4>
      <table>
        <thead>
        <tr>
          <th>شرط</th>
          <th>روش پیشنهادی</th>
          <th>دلیل</th>
          <th>وضعیت</th>
          <th>عملیات</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="rule in selectionRules" :key="rule.id">
          <td>{{ rule.condition }}</td>
          <td>{{ rule.recommendedMethod }}</td>
          <td>{{ rule.reason }}</td>
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
            <label>شرط:</label>
            <input v-model="newRule.condition" type="text" required>
          </div>
          <div class="form-group">
            <label>روش پیشنهادی:</label>
            <select v-model="newRule.recommendedMethod" required>
              <option value="">انتخاب کنید</option>
              <option value="express">ارسال سریع</option>
              <option value="standard">ارسال استاندارد</option>
              <option value="economy">ارسال اقتصادی</option>
            </select>
          </div>
        </div>
        <div class="form-group">
          <label>دلیل:</label>
          <textarea v-model="newRule.reason" rows="2" required></textarea>
        </div>
        <button type="submit" class="btn-primary">افزودن قانون</button>
      </form>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
</script>

<script setup lang="ts">
import { reactive, ref } from 'vue'

definePageMeta({ layout: 'admin-main', middleware: 'admin' })

const criteria = ref([
  {
    id: 1,
    title: 'سرعت',
    description: 'زمان تحویل مورد نیاز',
    icon: 'fas fa-clock'
  },
  {
    id: 2,
    title: 'هزینه',
    description: 'بودجه در نظر گرفته شده',
    icon: 'fas fa-dollar-sign'
  },
  {
    id: 3,
    title: 'امنیت',
    description: 'میزان امنیت مورد نیاز',
    icon: 'fas fa-shield-alt'
  }
])

const selectionRules = ref([
  {
    id: 1,
    condition: 'ارزش سفارش > 1000 تومان',
    recommendedMethod: 'ارسال سریع',
    reason: 'ارزش بالای سفارش نیاز به امنیت بیشتر دارد',
    status: 'active',
    statusText: 'فعال'
  }
])

const newRule = reactive({
  condition: '',
  recommendedMethod: '',
  reason: ''
})

const editRule = (_rule) => {
  // Edit rule logic
}

const deleteRule = (_id) => {
  // Delete rule logic
}

const addNewRule = () => {
  // Add new rule logic
}
</script>

<style scoped>
.selection-guide {
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

.selection-criteria,
.selection-rules,
.add-rule-section {
  background: white;
  border-radius: 8px;
  padding: 20px;
  margin-bottom: 20px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.criteria-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 20px;
  margin-top: 15px;
}

.criterion-card {
  display: flex;
  align-items: center;
  padding: 15px;
  border: 1px solid #ecf0f1;
  border-radius: 8px;
  background-color: #f8f9fa;
}

.criterion-icon {
  font-size: 24px;
  color: #007bff;
  margin-left: 15px;
}

.criterion-content h5 {
  margin: 0 0 5px 0;
  color: #2c3e50;
}

.criterion-content p {
  margin: 0;
  color: #7f8c8d;
  font-size: 14px;
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
