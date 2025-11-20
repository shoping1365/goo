<template>
  <div class="shipping-method-images">
    <div class="section-header">
      <h3>عکس‌های روش ارسال</h3>
      <p>مدیریت و بارگذاری تصاویر برای هر روش ارسال</p>
    </div>

    <!-- Images List -->
    <div class="images-list">
      <h4>لیست تصاویر روش‌های ارسال</h4>
      <table>
        <thead>
        <tr>
          <th>روش ارسال</th>
          <th>تصویر</th>
          <th>عنوان</th>
          <th>وضعیت</th>
          <th>عملیات</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="image in images" :key="image.id">
          <td>{{ image.shippingMethod }}</td>
          <td>
            <img :src="image.url" :alt="image.title" class="method-image" />
          </td>
          <td>{{ image.title }}</td>
          <td>
            <span :class="['status', image.status]">{{ image.statusText }}</span>
          </td>
          <td>
            <button class="btn-edit" @click="editImage(image)">ویرایش</button>
            <button class="btn-delete" @click="deleteImage(image.id)">حذف</button>
          </td>
        </tr>
        </tbody>
      </table>
    </div>

    <!-- Add New Image -->
    <div class="add-image-section">
      <h4>افزودن تصویر جدید</h4>
      <form @submit.prevent="addNewImage">
        <div class="form-row">
          <div class="form-group">
            <label>روش ارسال:</label>
            <select v-model="newImage.shippingMethod" required>
              <option value="">انتخاب کنید</option>
              <option value="express">ارسال سریع</option>
              <option value="standard">ارسال استاندارد</option>
              <option value="economy">ارسال اقتصادی</option>
              <option value="same_day">ارسال همان روز</option>
            </select>
          </div>
          <div class="form-group">
            <label>عنوان:</label>
            <input v-model="newImage.title" type="text" required>
          </div>
        </div>
        <div class="form-group">
          <label>انتخاب تصویر:</label>
          <input type="file" accept="image/*" required @change="onFileChange">
        </div>
        <button type="submit" class="btn-primary">افزودن تصویر</button>
      </form>
    </div>
  </div>
</template>

<script setup>
import { reactive, ref } from 'vue'

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
// const { user, hasPermission } = useAuth()

const images = ref([
  {
    id: 1,
    shippingMethod: 'ارسال سریع',
    url: '/images/shipping/express.png',
    title: 'ارسال سریع',
    status: 'active',
    statusText: 'فعال'
  },
  {
    id: 2,
    shippingMethod: 'ارسال استاندارد',
    url: '/images/shipping/standard.png',
    title: 'ارسال استاندارد',
    status: 'active',
    statusText: 'فعال'
  }
])

const newImage = reactive({
  shippingMethod: '',
  title: '',
  file: null
})

const onFileChange = (e) => {
  newImage.file = e.target.files[0]
}

const editImage = (_image) => {
  // Edit image logic
}

const deleteImage = (_id) => {
  // Delete image logic
}

const addNewImage = () => {
  // Add new image logic
}
</script>

<style scoped>
.shipping-method-images {
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

.images-list,
.add-image-section {
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

.method-image {
  width: 60px;
  height: 40px;
  object-fit: contain;
  border-radius: 4px;
  border: 1px solid #eee;
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

