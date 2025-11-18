<template>
  <div class="emergency-settings">
<div class="section-header">
  <h3>تنظیمات اضطراری</h3>
  <p>مدیریت تنظیمات و قوانین در شرایط اضطراری</p>
</div>

<!-- Emergency Settings List -->
<div class="emergency-settings-list">
  <h4>لیست تنظیمات اضطراری</h4>
  <table>
    <thead>
    <tr>
      <th>عنوان تنظیم</th>
      <th>نوع</th>
      <th>شرط فعال‌سازی</th>
      <th>اقدام</th>
      <th>وضعیت</th>
      <th>عملیات</th>
    </tr>
    </thead>
    <tbody>
    <tr v-for="setting in emergencySettings" :key="setting.id">
      <td>{{ setting.title }}</td>
      <td>{{ setting.type }}</td>
      <td>{{ setting.activationCondition }}</td>
      <td>{{ setting.action }}</td>
      <td>
              <span :class="setting.active ? 'active' : 'inactive'">
                {{ setting.active ? 'فعال' : 'غیرفعال' }}
              </span>
      </td>
      <td>
        <button @click="editSetting(setting.id)" class="btn btn-secondary">
          <i class="fas fa-edit"></i>
          ویرایش
        </button>
        <button @click="toggleSetting(setting.id)" class="btn" :class="setting.active ? 'btn-warning' : 'btn-success'">
          <i :class="setting.active ? 'fas fa-pause' : 'fas fa-play'"></i>
          {{ setting.active ? 'غیرفعال' : 'فعال' }}
        </button>
        <button @click="deleteSetting(setting.id)" class="btn btn-danger">
          <i class="fas fa-trash"></i>
          حذف
        </button>
      </td>
    </tr>
    </tbody>
  </table>
</div>

<!-- Add New Emergency Setting -->
<div class="add-setting-section">
  <h4>افزودن تنظیم اضطراری جدید</h4>
  <div class="add-setting-form">
    <div class="form-row">
      <div class="form-group">
        <label>عنوان تنظیم:</label>
        <input type="text" v-model="newSetting.title" placeholder="مثال: محدودیت ارسال">
      </div>
      <div class="form-group">
        <label>نوع:</label>
        <select v-model="newSetting.type">
          <option value="restriction">محدودیت</option>
          <option value="notification">اعلان</option>
          <option value="alternative">جایگزین</option>
        </select>
      </div>
    </div>
    <div class="form-row">
      <div class="form-group">
        <label>شرط فعال‌سازی:</label>
        <input type="text" v-model="newSetting.activationCondition" placeholder="مثال: تاخیر بیش از ۳ روز">
      </div>
      <div class="form-group">
        <label>اقدام:</label>
        <input type="text" v-model="newSetting.action" placeholder="مثال: ارسال اعلان">
      </div>
    </div>
    <div class="form-actions">
      <button @click="addSetting" class="btn btn-success">
        <i class="fas fa-plus"></i>
        افزودن تنظیم
      </button>
      <button @click="resetNewSetting" class="btn btn-secondary">
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

const emergencySettings = ref([
  {
    id: 'e1',
    title: 'محدودیت ارسال',
    type: 'restriction',
    activationCondition: 'تاخیر بیش از ۳ روز',
    action: 'غیرفعال کردن ارسال',
    active: true
  },
  {
    id: 'e2',
    title: 'اعلان اضطراری',
    type: 'notification',
    activationCondition: 'خطا در سیستم',
    action: 'ارسال اعلان',
    active: false
  }
])

const newSetting = ref({
  title: '',
  type: 'restriction',
  activationCondition: '',
  action: ''
})

const addSetting = () => {
  if (!newSetting.value.title || !newSetting.value.activationCondition || !newSetting.value.action) {
    alert('لطفاً تمام فیلدها را پر کنید')
    return
  }
  emergencySettings.value.push({
    id: `e${Date.now()}`,
    ...newSetting.value,
    active: true
  })
  resetNewSetting()
}

const resetNewSetting = () => {
  newSetting.value.title = ''
  newSetting.value.type = 'restriction'
  newSetting.value.activationCondition = ''
  newSetting.value.action = ''
}

const editSetting = (id: string) => {
  alert('ویرایش تنظیم: ' + id)
}

const toggleSetting = (id: string) => {
  const setting = emergencySettings.value.find(s => s.id === id)
  if (setting) setting.active = !setting.active
}

const deleteSetting = (id: string) => {
  if (confirm('آیا از حذف این تنظیم اطمینان دارید؟')) {
    emergencySettings.value = emergencySettings.value.filter(s => s.id !== id)
  }
}
</script>

<style scoped>
.emergency-settings {
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

.emergency-settings-list {
  margin-bottom: 30px;
}

.emergency-settings-list h4 {
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

.add-setting-section {
  background: #f8f9fa;
  padding: 20px;
  border-radius: 8px;
  border: 2px dashed #dee2e6;
}

.add-setting-section h4 {
  margin-bottom: 20px;
  color: #2c3e50;
}

.add-setting-form {
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
