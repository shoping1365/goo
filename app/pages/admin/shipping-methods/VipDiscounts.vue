<template>
  <div class="vip-discounts">
    <div class="section-header">
      <h3>تخفیف‌های مشتریان VIP</h3>
      <p>مدیریت و تعریف تخفیف‌های ویژه برای مشتریان VIP</p>
    </div>

    <!-- VIP Discount List -->
    <div class="discount-list">
      <h4>لیست تخفیف‌های فعال</h4>
      <table>
        <thead>
        <tr>
          <th>عنوان تخفیف</th>
          <th>درصد تخفیف</th>
          <th>روش ارسال</th>
          <th>تاریخ شروع</th>
          <th>تاریخ پایان</th>
          <th>وضعیت</th>
          <th>عملیات</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="discount in discounts" :key="discount.id">
          <td>{{ discount.title }}</td>
          <td>{{ discount.percent }}%</td>
          <td>{{ discount.shippingMethod }}</td>
          <td>{{ formatDate(discount.startDate) }}</td>
          <td>{{ formatDate(discount.endDate) }}</td>
          <td>
              <span :class="discount.active ? 'active' : 'inactive'">
                {{ discount.active ? 'فعال' : 'غیرفعال' }}
              </span>
          </td>
          <td>
            <button class="btn btn-secondary" @click="editDiscount(discount.id)">
              <i class="fas fa-edit"></i>
              ویرایش
            </button>
            <button class="btn" :class="discount.active ? 'btn-warning' : 'btn-success'" @click="toggleDiscount(discount.id)">
              <i :class="discount.active ? 'fas fa-pause' : 'fas fa-play'"></i>
              {{ discount.active ? 'غیرفعال' : 'فعال' }}
            </button>
            <button class="btn btn-danger" @click="deleteDiscount(discount.id)">
              <i class="fas fa-trash"></i>
              حذف
            </button>
          </td>
        </tr>
        </tbody>
      </table>
    </div>

    <!-- Add New Discount -->
    <div class="add-discount-section">
      <h4>افزودن تخفیف جدید</h4>
      <div class="add-discount-form">
        <div class="form-row">
          <div class="form-group">
            <label>عنوان تخفیف:</label>
            <input v-model="newDiscount.title" type="text" placeholder="مثال: تخفیف ویژه VIP">
          </div>
          <div class="form-group">
            <label>درصد تخفیف:</label>
            <input v-model="newDiscount.percent" type="number" min="1" max="100" placeholder="10">
          </div>
        </div>
        <div class="form-row">
          <div class="form-group">
            <label>روش ارسال:</label>
            <select v-model="newDiscount.shippingMethod">
              <option v-for="method in shippingMethods" :key="method.id" :value="method.name">
                {{ method.name }}
              </option>
            </select>
          </div>
          <div class="form-group">
            <label>تاریخ شروع:</label>
            <input v-model="newDiscount.startDate" type="date">
          </div>
          <div class="form-group">
            <label>تاریخ پایان:</label>
            <input v-model="newDiscount.endDate" type="date">
          </div>
        </div>
        <div class="form-actions">
          <button class="btn btn-success" @click="addDiscount">
            <i class="fas fa-plus"></i>
            افزودن تخفیف
          </button>
          <button class="btn btn-secondary" @click="resetNewDiscount">
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

const discounts = ref([
  {
    id: 'd1',
    title: 'تخفیف ویژه VIP زمستان',
    percent: 15,
    shippingMethod: 'ارسال سریع',
    startDate: '2024-01-01',
    endDate: '2024-02-01',
    active: true
  },
  {
    id: 'd2',
    title: 'تخفیف تابستانی VIP',
    percent: 10,
    shippingMethod: 'ارسال استاندارد',
    startDate: '2024-06-01',
    endDate: '2024-07-01',
    active: false
  }
])

const newDiscount = ref({
  title: '',
  percent: 10,
  shippingMethod: 'ارسال استاندارد',
  startDate: '',
  endDate: ''
})

const formatDate = (date: string) => {
  if (!date) return '-'
  return new Intl.DateTimeFormat('fa-IR').format(new Date(date))
}

const addDiscount = () => {
  if (!newDiscount.value.title || !newDiscount.value.percent || !newDiscount.value.startDate || !newDiscount.value.endDate) {
    alert('لطفاً تمام فیلدها را پر کنید')
    return
  }
  discounts.value.push({
    id: `d${Date.now()}`,
    ...newDiscount.value,
    active: true
  })
  resetNewDiscount()
}

const resetNewDiscount = () => {
  newDiscount.value.title = ''
  newDiscount.value.percent = 10
  newDiscount.value.shippingMethod = 'ارسال استاندارد'
  newDiscount.value.startDate = ''
  newDiscount.value.endDate = ''
}

const editDiscount = (id: string) => {
  // Implementation for editing discount
  alert('ویرایش تخفیف: ' + id)
}

const toggleDiscount = (id: string) => {
  const discount = discounts.value.find(d => d.id === id)
  if (discount) discount.active = !discount.active
}

const deleteDiscount = (id: string) => {
  if (confirm('آیا از حذف این تخفیف اطمینان دارید؟')) {
    discounts.value = discounts.value.filter(d => d.id !== id)
  }
}
</script>

<style scoped>
.vip-discounts {
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

.discount-list {
  margin-bottom: 30px;
}

.discount-list h4 {
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

.add-discount-section {
  background: #f8f9fa;
  padding: 20px;
  border-radius: 8px;
  border: 2px dashed #dee2e6;
}

.add-discount-section h4 {
  margin-bottom: 20px;
  color: #2c3e50;
}

.add-discount-form {
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
