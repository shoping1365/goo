<template>
  <div class="delay-notifications">
    <div class="section-header">
      <h3>اعلان‌های تاخیر</h3>
      <p>مدیریت اعلان‌های خودکار برای تاخیر در ارسال و تحویل</p>
    </div>

    <!-- Delay Notification Rules -->
    <div class="delay-rules">
      <h4>قوانین اعلان تاخیر</h4>
      <table>
        <thead>
        <tr>
          <th>عنوان قانون</th>
          <th>نوع تاخیر</th>
          <th>زمان تاخیر</th>
          <th>مخاطب</th>
          <th>کانال ارسال</th>
          <th>وضعیت</th>
          <th>عملیات</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="rule in delayRules" :key="rule.id">
          <td>{{ rule.title }}</td>
          <td>{{ rule.delayType }}</td>
          <td>{{ rule.delayTime }}</td>
          <td>{{ rule.recipient }}</td>
          <td>{{ rule.channel }}</td>
          <td>
              <span :class="rule.active ? 'active' : 'inactive'">
                {{ rule.active ? 'فعال' : 'غیرفعال' }}
              </span>
          </td>
          <td>
            <button class="btn btn-secondary" @click="editRule(rule.id)">
              <i class="fas fa-edit"></i>
              ویرایش
            </button>
            <button class="btn" :class="rule.active ? 'btn-warning' : 'btn-success'" @click="toggleRule(rule.id)">
              <i :class="rule.active ? 'fas fa-pause' : 'fas fa-play'"></i>
              {{ rule.active ? 'غیرفعال' : 'فعال' }}
            </button>
            <button class="btn btn-danger" @click="deleteRule(rule.id)">
              <i class="fas fa-trash"></i>
              حذف
            </button>
          </td>
        </tr>
        </tbody>
      </table>
    </div>

    <!-- Add New Delay Rule -->
    <div class="add-rule-section">
      <h4>افزودن قانون جدید</h4>
      <div class="add-rule-form">
        <div class="form-row">
          <div class="form-group">
            <label>عنوان قانون:</label>
            <input v-model="newRule.title" type="text" placeholder="مثال: تاخیر در ارسال">
          </div>
          <div class="form-group">
            <label>نوع تاخیر:</label>
            <select v-model="newRule.delayType">
              <option value="shipping">تاخیر در ارسال</option>
              <option value="delivery">تاخیر در تحویل</option>
              <option value="processing">تاخیر در پردازش</option>
            </select>
          </div>
        </div>
        <div class="form-row">
          <div class="form-group">
            <label>زمان تاخیر:</label>
            <input v-model="newRule.delayTime" type="text" placeholder="مثال: ۲۴ ساعت">
          </div>
          <div class="form-group">
            <label>مخاطب:</label>
            <select v-model="newRule.recipient">
              <option value="customer">مشتری</option>
              <option value="admin">مدیر</option>
              <option value="both">هر دو</option>
            </select>
          </div>
          <div class="form-group">
            <label>کانال ارسال:</label>
            <select v-model="newRule.channel">
              <option value="email">ایمیل</option>
              <option value="sms">پیامک</option>
              <option value="both">هر دو</option>
            </select>
          </div>
        </div>
        <div class="form-actions">
          <button class="btn btn-success" @click="addRule">
            <i class="fas fa-plus"></i>
            افزودن قانون
          </button>
          <button class="btn btn-secondary" @click="resetNewRule">
            <i class="fas fa-undo"></i>
            بازنشانی
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
</script>

<script setup lang="ts">
import { ref } from 'vue';

definePageMeta({ layout: 'admin-main', middleware: 'admin' });

const delayRules = ref([
  {
    id: 'dr1',
    title: 'تاخیر در ارسال',
    delayType: 'shipping',
    delayTime: '۲۴ ساعت',
    recipient: 'customer',
    channel: 'email',
    active: true
  },
  {
    id: 'dr2',
    title: 'تاخیر در تحویل',
    delayType: 'delivery',
    delayTime: '۴۸ ساعت',
    recipient: 'both',
    channel: 'both',
    active: false
  }
])

const newRule = ref({
  title: '',
  delayType: 'shipping',
  delayTime: '',
  recipient: 'customer',
  channel: 'email'
})

const addRule = () => {
  if (!newRule.value.title || !newRule.value.delayTime) {
    alert('لطفاً تمام فیلدها را پر کنید')
    return
  }
  delayRules.value.push({
    id: `dr${Date.now()}`,
    ...newRule.value,
    active: true
  })
  resetNewRule()
}

const resetNewRule = () => {
  newRule.value.title = ''
  newRule.value.delayType = 'shipping'
  newRule.value.delayTime = ''
  newRule.value.recipient = 'customer'
  newRule.value.channel = 'email'
}

const editRule = (id: string) => {
  alert('ویرایش قانون: ' + id)
}

const toggleRule = (id: string) => {
  const rule = delayRules.value.find(r => r.id === id)
  if (rule) rule.active = !rule.active
}

const deleteRule = (id: string) => {
  if (confirm('آیا از حذف این قانون اطمینان دارید؟')) {
    delayRules.value = delayRules.value.filter(r => r.id !== id)
  }
}
</script>

<style scoped>
.delay-notifications {
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

.delay-rules {
  margin-bottom: 30px;
}

.delay-rules h4 {
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
