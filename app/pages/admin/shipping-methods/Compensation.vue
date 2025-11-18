<template>
  <div class="compensation">
    <div class="section-header">
      <h3>جبران خسارت</h3>
      <p>مدیریت قوانین و روش‌های جبران خسارت برای مشتریان</p>
    </div>

    <!-- Compensation Rules -->
    <div class="compensation-rules">
      <h4>قوانین جبران خسارت</h4>
      <table>
        <thead>
        <tr>
          <th>عنوان قانون</th>
          <th>نوع خسارت</th>
          <th>مبلغ جبران</th>
          <th>شرط اعمال</th>
          <th>وضعیت</th>
          <th>عملیات</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="rule in compensationRules" :key="rule.id">
          <td>{{ rule.title }}</td>
          <td>{{ rule.damageType }}</td>
          <td>{{ rule.compensationAmount }}</td>
          <td>{{ rule.condition }}</td>
          <td>
              <span :class="rule.active ? 'active' : 'inactive'">
                {{ rule.active ? 'فعال' : 'غیرفعال' }}
              </span>
          </td>
          <td>
            <button @click="editRule(rule.id)" class="btn btn-secondary">
              <i class="fas fa-edit"></i>
              ویرایش
            </button>
            <button @click="toggleRule(rule.id)" class="btn" :class="rule.active ? 'btn-warning' : 'btn-success'">
              <i :class="rule.active ? 'fas fa-pause' : 'fas fa-play'"></i>
              {{ rule.active ? 'غیرفعال' : 'فعال' }}
            </button>
            <button @click="deleteRule(rule.id)" class="btn btn-danger">
              <i class="fas fa-trash"></i>
              حذف
            </button>
          </td>
        </tr>
        </tbody>
      </table>
    </div>

    <!-- Add New Compensation Rule -->
    <div class="add-rule-section">
      <h4>افزودن قانون جدید</h4>
      <div class="add-rule-form">
        <div class="form-row">
          <div class="form-group">
            <label>عنوان قانون:</label>
            <input type="text" v-model="newRule.title" placeholder="مثال: جبران تاخیر">
          </div>
          <div class="form-group">
            <label>نوع خسارت:</label>
            <select v-model="newRule.damageType">
              <option value="delay">تاخیر</option>
              <option value="damage">آسیب</option>
              <option value="loss">گم شدن</option>
            </select>
          </div>
        </div>
        <div class="form-row">
          <div class="form-group">
            <label>مبلغ جبران:</label>
            <input type="text" v-model="newRule.compensationAmount" placeholder="مثال: ۱۰۰,۰۰۰ تومان">
          </div>
          <div class="form-group">
            <label>شرط اعمال:</label>
            <input type="text" v-model="newRule.condition" placeholder="مثال: تاخیر بیش از ۳ روز">
          </div>
        </div>
        <div class="form-actions">
          <button @click="addRule" class="btn btn-success">
            <i class="fas fa-plus"></i>
            افزودن قانون
          </button>
          <button @click="resetNewRule" class="btn btn-secondary">
            <i class="fas fa-undo"></i>
            بازنشانی
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const compensationRules = ref([
  {
    id: 'cr1',
    title: 'جبران تاخیر',
    damageType: 'delay',
    compensationAmount: '۱۰۰,۰۰۰ تومان',
    condition: 'تاخیر بیش از ۳ روز',
    active: true
  },
  {
    id: 'cr2',
    title: 'جبران آسیب',
    damageType: 'damage',
    compensationAmount: '۵۰٪ ارزش کالا',
    condition: 'آسیب به کالا',
    active: false
  }
])

const newRule = ref({
  title: '',
  damageType: 'delay',
  compensationAmount: '',
  condition: ''
})

const addRule = () => {
  if (!newRule.value.title || !newRule.value.compensationAmount || !newRule.value.condition) {
    alert('لطفاً تمام فیلدها را پر کنید')
    return
  }
  compensationRules.value.push({
    id: `cr${Date.now()}`,
    ...newRule.value,
    active: true
  })
  resetNewRule()
}

const resetNewRule = () => {
  newRule.value.title = ''
  newRule.value.damageType = 'delay'
  newRule.value.compensationAmount = ''
  newRule.value.condition = ''
}

const editRule = (id: string) => {
  alert('ویرایش قانون: ' + id)
}

const toggleRule = (id: string) => {
  const rule = compensationRules.value.find(r => r.id === id)
  if (rule) rule.active = !rule.active
}

const deleteRule = (id: string) => {
  if (confirm('آیا از حذف این قانون اطمینان دارید؟')) {
    compensationRules.value = compensationRules.value.filter(r => r.id !== id)
  }
}
</script>

<style scoped>
.compensation {
  padding: 20px;
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.section-header {
  margin-bottom: 30px;
  text-align: center;
}

.section-header h3 {
  color: #2c3e50;
  margin-bottom: 10px;
  font-size: 24px;
}

.section-header p {
  color: #7f8c8d;
  font-size: 14px;
}

.compensation-rules {
  margin-bottom: 30px;
}

.compensation-rules h4 {
  margin-bottom: 15px;
  color: #2c3e50;
}

table {
  width: 100%;
  border-collapse: collapse;
  background: white;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 2px 4px rgba(0,0,0,0.05);
}

th, td {
  padding: 12px;
  text-align: center;
  border-bottom: 1px solid #eee;
}

th {
  background: #f8f9fa;
  font-weight: 600;
  color: #2c3e50;
}

.active {
  color: #27ae60;
  font-weight: 600;
}

.inactive {
  color: #e74c3c;
  font-weight: 600;
}

.add-rule-section {
  background: #f8f9fa;
  padding: 20px;
  border-radius: 8px;
  border: 2px dashed #dee2e6;
}

.add-rule-section h4 {
  margin-bottom: 20px;
  color: #2c3e50;
}

.add-rule-form {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.form-row {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 15px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 5px;
}

.form-group label {
  font-weight: 600;
  color: #2c3e50;
  font-size: 14px;
}

.form-group input,
.form-group select {
  padding: 8px 12px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 14px;
}

.form-actions {
  display: flex;
  gap: 10px;
  justify-content: flex-end;
}

.btn {
  padding: 8px 16px;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 14px;
  display: flex;
  align-items: center;
  gap: 6px;
  transition: all 0.3s ease;
}

.btn-success {
  background: #27ae60;
  color: white;
}

.btn-success:hover {
  background: #229954;
}

.btn-secondary {
  background: #6c757d;
  color: white;
}

.btn-secondary:hover {
  background: #5a6268;
}

.btn-warning {
  background: #f39c12;
  color: white;
}

.btn-warning:hover {
  background: #e67e22;
}

.btn-danger {
  background: #e74c3c;
  color: white;
}

.btn-danger:hover {
  background: #c0392b;
}
</style>
