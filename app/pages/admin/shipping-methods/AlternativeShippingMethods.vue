<template>
  <div class="alternative-shipping-methods">
    <div class="section-header">
      <h3>روش‌های ارسال جایگزین</h3>
      <p>مدیریت روش‌های ارسال جایگزین در شرایط اضطراری</p>
    </div>

    <!-- Alternative Methods List -->
    <div class="alternative-methods-list">
      <h4>لیست روش‌های جایگزین</h4>
      <table>
        <thead>
        <tr>
          <th>روش اصلی</th>
          <th>روش جایگزین</th>
          <th>شرط فعال‌سازی</th>
          <th>اولویت</th>
          <th>وضعیت</th>
          <th>عملیات</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="method in alternativeMethods" :key="method.id">
          <td>{{ method.primaryMethod }}</td>
          <td>{{ method.alternativeMethod }}</td>
          <td>{{ method.activationCondition }}</td>
          <td>{{ method.priority }}</td>
          <td>
              <span :class="method.active ? 'active' : 'inactive'">
                {{ method.active ? 'فعال' : 'غیرفعال' }}
              </span>
          </td>
          <td>
            <button class="btn btn-secondary" @click="editMethod(method.id)">
              <i class="fas fa-edit"></i>
              ویرایش
            </button>
            <button class="btn" :class="method.active ? 'btn-warning' : 'btn-success'" @click="toggleMethod(method.id)">
              <i :class="method.active ? 'fas fa-pause' : 'fas fa-play'"></i>
              {{ method.active ? 'غیرفعال' : 'فعال' }}
            </button>
            <button class="btn btn-danger" @click="deleteMethod(method.id)">
              <i class="fas fa-trash"></i>
              حذف
            </button>
          </td>
        </tr>
        </tbody>
      </table>
    </div>

    <!-- Add New Alternative Method -->
    <div class="add-method-section">
      <h4>افزودن روش جایگزین جدید</h4>
      <div class="add-method-form">
        <div class="form-row">
          <div class="form-group">
            <label>روش اصلی:</label>
            <select v-model="newMethod.primaryMethod">
              <option v-for="method in shippingMethods" :key="method.id" :value="method.name">
                {{ method.name }}
              </option>
            </select>
          </div>
          <div class="form-group">
            <label>روش جایگزین:</label>
            <select v-model="newMethod.alternativeMethod">
              <option v-for="method in shippingMethods" :key="method.id" :value="method.name">
                {{ method.name }}
              </option>
            </select>
          </div>
        </div>
        <div class="form-row">
          <div class="form-group">
            <label>شرط فعال‌سازی:</label>
            <input v-model="newMethod.activationCondition" type="text" placeholder="مثال: عدم موجودی">
          </div>
          <div class="form-group">
            <label>اولویت:</label>
            <select v-model="newMethod.priority">
              <option value="1">۱ (بالاترین)</option>
              <option value="2">۲</option>
              <option value="3">۳</option>
            </select>
          </div>
        </div>
        <div class="form-actions">
          <button class="btn btn-success" @click="addMethod">
            <i class="fas fa-plus"></i>
            افزودن روش
          </button>
          <button class="btn btn-secondary" @click="resetNewMethod">
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

const alternativeMethods = ref([
  {
    id: 'a1',
    primaryMethod: 'ارسال سریع',
    alternativeMethod: 'ارسال استاندارد',
    activationCondition: 'عدم موجودی',
    priority: 1,
    active: true
  },
  {
    id: 'a2',
    primaryMethod: 'ارسال رایگان',
    alternativeMethod: 'ارسال استاندارد',
    activationCondition: 'محدودیت منطقه',
    priority: 2,
    active: false
  }
])

const newMethod = ref({
  primaryMethod: 'ارسال استاندارد',
  alternativeMethod: 'ارسال سریع',
  activationCondition: '',
  priority: 1
})

const addMethod = () => {
  if (!newMethod.value.activationCondition) {
    alert('لطفاً تمام فیلدها را پر کنید')
    return
  }
  alternativeMethods.value.push({
    id: `a${Date.now()}`,
    ...newMethod.value,
    active: true
  })
  resetNewMethod()
}

const resetNewMethod = () => {
  newMethod.value.primaryMethod = 'ارسال استاندارد'
  newMethod.value.alternativeMethod = 'ارسال سریع'
  newMethod.value.activationCondition = ''
  newMethod.value.priority = 1
}

const editMethod = (id: string) => {
  alert('ویرایش روش: ' + id)
}

const toggleMethod = (id: string) => {
  const method = alternativeMethods.value.find(m => m.id === id)
  if (method) method.active = !method.active
}

const deleteMethod = (id: string) => {
  if (confirm('آیا از حذف این روش اطمینان دارید؟')) {
    alternativeMethods.value = alternativeMethods.value.filter(m => m.id !== id)
  }
}
</script>

<style scoped>
.alternative-shipping-methods {
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

.alternative-methods-list {
  margin-bottom: 30px;
}

.alternative-methods-list h4 {
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

.add-method-section {
  background: #f8f9fa;
  padding: 20px;
  border-radius: 8px;
  border: 2px dashed #dee2e6;
}

.add-method-section h4 {
  margin-bottom: 20px;
  color: #2c3e50;
}

.add-method-form {
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
