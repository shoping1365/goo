<template>
  <div class="suggested-shipping-methods">
    <div class="section-header">
      <h3>روش‌های ارسال پیشنهادی</h3>
      <p>پیشنهاد هوشمند روش ارسال بر اساس نوع مشتری و سفارش</p>
    </div>

    <!-- Suggestion Rules List -->
    <div class="rules-list">
      <h4>قوانین پیشنهاد</h4>
      <table>
        <thead>
        <tr>
          <th>عنوان قانون</th>
          <th>شرط</th>
          <th>روش پیشنهادی</th>
          <th>اولویت</th>
          <th>وضعیت</th>
          <th>عملیات</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="rule in suggestionRules" :key="rule.id">
          <td>{{ rule.title }}</td>
          <td>{{ rule.condition }}</td>
          <td>{{ rule.suggestedMethod }}</td>
          <td>{{ rule.priority }}</td>
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

    <!-- Add New Suggestion Rule -->
    <div class="add-rule-section">
      <h4>افزودن قانون جدید</h4>
      <div class="add-rule-form">
        <div class="form-row">
          <div class="form-group">
            <label>عنوان قانون:</label>
            <input type="text" v-model="newRule.title" placeholder="مثال: سفارشات VIP">
          </div>
          <div class="form-group">
            <label>شرط:</label>
            <input type="text" v-model="newRule.condition" placeholder="مثال: گروه مشتری = VIP">
          </div>
        </div>
        <div class="form-row">
          <div class="form-group">
            <label>روش پیشنهادی:</label>
            <select v-model="newRule.suggestedMethod">
              <option v-for="method in shippingMethods" :key="method.id" :value="method.name">
                {{ method.name }}
              </option>
            </select>
          </div>
          <div class="form-group">
            <label>اولویت:</label>
            <select v-model="newRule.priority">
              <option value="1">۱ (بالاترین)</option>
              <option value="2">۲</option>
              <option value="3">۳</option>
            </select>
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

const shippingMethods = ref([
  { id: 'standard', name: 'ارسال استاندارد' },
  { id: 'express', name: 'ارسال سریع' },
  { id: 'free', name: 'ارسال رایگان' }
])

const suggestionRules = ref([
  {
    id: 'r1',
    title: 'VIP',
    condition: 'گروه مشتری = VIP',
    suggestedMethod: 'ارسال سریع',
    priority: 1,
    active: true
  },
  {
    id: 'r2',
    title: 'سفارشات مهمان',
    condition: 'گروه مشتری = مهمان',
    suggestedMethod: 'ارسال استاندارد',
    priority: 2,
    active: true
  }
])

const newRule = ref({
  title: '',
  condition: '',
  suggestedMethod: 'ارسال استاندارد',
  priority: 1
})

const addRule = () => {
  if (!newRule.value.title || !newRule.value.condition) {
    alert('لطفاً تمام فیلدها را پر کنید')
    return
  }
  suggestionRules.value.push({
    id: `r${Date.now()}`,
    ...newRule.value,
    active: true
  })
  resetNewRule()
}

const resetNewRule = () => {
  newRule.value.title = ''
  newRule.value.condition = ''
  newRule.value.suggestedMethod = 'ارسال استاندارد'
  newRule.value.priority = 1
}

const editRule = (id: string) => {
  alert('ویرایش قانون: ' + id)
}

const toggleRule = (id: string) => {
  const rule = suggestionRules.value.find(r => r.id === id)
  if (rule) rule.active = !rule.active
}

const deleteRule = (id: string) => {
  if (confirm('آیا از حذف این قانون اطمینان دارید؟')) {
    suggestionRules.value = suggestionRules.value.filter(r => r.id !== id)
  }
}
</script>

<style scoped>
.suggested-shipping-methods {
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

.rules-list {
  margin-bottom: 30px;
}

.rules-list h4 {
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
