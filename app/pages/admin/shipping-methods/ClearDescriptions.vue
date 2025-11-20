<template>
  <div class="clear-descriptions">
    <div class="section-header">
      <h3>توضیحات واضح</h3>
      <p>مدیریت توضیحات روشن و قابل فهم برای روش‌های ارسال</p>
    </div>

    <!-- Shipping Method Descriptions -->
    <div class="descriptions-list">
      <h4>لیست توضیحات روش‌های ارسال</h4>
      <table>
        <thead>
        <tr>
          <th>روش ارسال</th>
          <th>عنوان</th>
          <th>توضیحات کوتاه</th>
          <th>توضیحات کامل</th>
          <th>زبان</th>
          <th>وضعیت</th>
          <th>عملیات</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="description in descriptions" :key="description.id">
          <td>{{ description.shippingMethod }}</td>
          <td>{{ description.title }}</td>
          <td>{{ description.shortDescription }}</td>
          <td>{{ description.fullDescription.substring(0, 100) }}...</td>
          <td>{{ description.language }}</td>
          <td>
            <span :class="['status', description.status]">{{ description.statusText }}</span>
          </td>
          <td>
            <button class="btn-edit" @click="editDescription(description)">ویرایش</button>
            <button class="btn-delete" @click="deleteDescription(description.id)">حذف</button>
          </td>
        </tr>
        </tbody>
      </table>
    </div>

    <!-- Add New Description -->
    <div class="add-description-section">
      <h4>افزودن توضیح جدید</h4>
      <form @submit.prevent="addNewDescription">
        <div class="form-row">
          <div class="form-group">
            <label>روش ارسال:</label>
            <select v-model="newDescription.shippingMethod" required>
              <option value="">انتخاب کنید</option>
              <option value="express">ارسال سریع</option>
              <option value="standard">ارسال استاندارد</option>
              <option value="economy">ارسال اقتصادی</option>
              <option value="same_day">ارسال همان روز</option>
            </select>
          </div>
          <div class="form-group">
            <label>زبان:</label>
            <select v-model="newDescription.language" required>
              <option value="fa">فارسی</option>
              <option value="en">انگلیسی</option>
              <option value="ar">عربی</option>
            </select>
          </div>
        </div>
        <div class="form-group">
          <label>عنوان:</label>
          <input v-model="newDescription.title" type="text" required>
        </div>
        <div class="form-group">
          <label>توضیحات کوتاه:</label>
          <textarea v-model="newDescription.shortDescription" rows="2" maxlength="200" placeholder="توضیح کوتاه و مختصر"></textarea>
          <div class="char-count">{{ newDescription.shortDescription.length }}/200</div>
        </div>
        <div class="form-group">
          <label>توضیحات کامل:</label>
          <textarea v-model="newDescription.fullDescription" rows="5" placeholder="توضیح کامل و جامع"></textarea>
        </div>
        <button type="submit" class="btn-primary">افزودن توضیح</button>
      </form>
    </div>

    <!-- Description Preview -->
    <div class="description-preview">
      <h4>پیش‌نمایش توضیحات</h4>
      <div class="preview-container">
        <div class="preview-tabs">
          <button
              v-for="lang in languages"
              :key="lang.code"
              :class="['tab-button', { active: selectedLanguage === lang.code }]"
              @click="selectedLanguage = lang.code"
          >
            {{ lang.name }}
          </button>
        </div>
        <div class="preview-content">
          <div v-for="method in shippingMethods" :key="method.id" class="method-preview">
            <h5>{{ method.name }}</h5>
            <div class="method-description">
              <p><strong>توضیح کوتاه:</strong> {{ getDescription(method.id, 'short') }}</p>
              <p><strong>توضیح کامل:</strong> {{ getDescription(method.id, 'full') }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { reactive, ref } from 'vue'

const selectedLanguage = ref('fa')

const newDescription = reactive({
  shippingMethod: '',
  language: 'fa',
  title: '',
  shortDescription: '',
  fullDescription: ''
})

const descriptions = ref([
  {
    id: 1,
    shippingMethod: 'express',
    title: 'ارسال سریع',
    shortDescription: 'ارسال در کمترین زمان ممکن',
    fullDescription: 'این روش ارسال برای افرادی که نیاز به دریافت سریع محصول دارند مناسب است. زمان تحویل بین 1 تا 3 روز کاری می‌باشد.',
    language: 'fa',
    status: 'active',
    statusText: 'فعال'
  },
  {
    id: 2,
    shippingMethod: 'standard',
    title: 'ارسال استاندارد',
    shortDescription: 'ارسال معمولی با قیمت مناسب',
    fullDescription: 'روش ارسال استاندارد که برای اکثر سفارشات مناسب است. زمان تحویل بین 3 تا 7 روز کاری می‌باشد.',
    language: 'fa',
    status: 'active',
    statusText: 'فعال'
  }
])

const languages = ref([
  { code: 'fa', name: 'فارسی' },
  { code: 'en', name: 'English' },
  { code: 'ar', name: 'العربية' }
])

const shippingMethods = ref([
  { id: 'express', name: 'ارسال سریع' },
  { id: 'standard', name: 'ارسال استاندارد' },
  { id: 'economy', name: 'ارسال اقتصادی' },
  { id: 'same_day', name: 'ارسال همان روز' }
])

function getDescription(methodId, type) {
  const desc = descriptions.value.find(d => d.shippingMethod === methodId && d.language === selectedLanguage.value)
  if (!desc) return 'توضیحی موجود نیست'
  return type === 'short' ? desc.shortDescription : desc.fullDescription
}

function editDescription(_description) {
  // console.log('Editing description:', description)
}

function deleteDescription(_id) {
  // console.log('Deleting description:', id)
}

function addNewDescription() {
  // console.log('Adding new description:', newDescription)
  // Reset form
  Object.assign(newDescription, {
    shippingMethod: '',
    language: 'fa',
    title: '',
    shortDescription: '',
    fullDescription: ''
  })
}
</script>

<style scoped>
.clear-descriptions {
  direction: rtl;
  padding: 1rem;
}

.section-header {
  margin-bottom: 2rem;
}

.section-header h3 {
  margin: 0 0 0.5rem 0;
  font-size: 1.5rem;
  color: #333;
}

.section-header p {
  margin: 0;
  color: #666;
}

.descriptions-list {
  margin-bottom: 2rem;
}

.descriptions-list h4 {
  margin: 0 0 1rem 0;
  font-size: 1.2rem;
  color: #333;
}

.descriptions-list table {
  width: 100%;
  border-collapse: collapse;
  background: white;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.descriptions-list th,
.descriptions-list td {
  padding: 1rem;
  text-align: right;
  border-bottom: 1px solid #e9ecef;
}

.descriptions-list th {
  background: #f8f9fa;
  font-weight: 600;
  color: #555;
}

.status {
  padding: 0.25rem 0.5rem;
  border-radius: 12px;
  font-size: 0.8rem;
}

.status.active {
  background: #d4edda;
  color: #155724;
}

.status.inactive {
  background: #f8d7da;
  color: #721c24;
}

.btn-edit,
.btn-delete {
  padding: 0.25rem 0.5rem;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 0.8rem;
  margin-left: 0.5rem;
}

.btn-edit {
  background: #007bff;
  color: white;
}

.btn-delete {
  background: #dc3545;
  color: white;
}

.add-description-section {
  background: #f8f9fa;
  border-radius: 12px;
  padding: 1.5rem;
  margin-bottom: 2rem;
}

.add-description-section h4 {
  margin: 0 0 1rem 0;
  font-size: 1.2rem;
  color: #333;
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1rem;
  margin-bottom: 1rem;
}

.form-group {
  margin-bottom: 1rem;
}

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: 500;
  color: #555;
}

.form-group input,
.form-group select,
.form-group textarea {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 0.9rem;
  font-family: inherit;
}

.form-group textarea {
  resize: vertical;
  min-height: 80px;
}

.char-count {
  text-align: left;
  font-size: 0.8rem;
  color: #666;
  margin-top: 0.25rem;
}

.btn-primary {
  background: #007bff;
  color: white;
  border: none;
  padding: 0.75rem 1.5rem;
  border-radius: 6px;
  cursor: pointer;
  font-size: 0.9rem;
}

.description-preview {
  background: white;
  border-radius: 12px;
  padding: 1.5rem;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.description-preview h4 {
  margin: 0 0 1rem 0;
  font-size: 1.2rem;
  color: #333;
}

.preview-tabs {
  display: flex;
  gap: 0.5rem;
  margin-bottom: 1rem;
  border-bottom: 1px solid #e9ecef;
}

.tab-button {
  background: none;
  border: none;
  padding: 0.75rem 1rem;
  cursor: pointer;
  border-bottom: 2px solid transparent;
  color: #666;
}

.tab-button.active {
  color: #007bff;
  border-bottom-color: #007bff;
}

.method-preview {
  margin-bottom: 1.5rem;
  padding: 1rem;
  background: #f8f9fa;
  border-radius: 8px;
}

.method-preview h5 {
  margin: 0 0 0.5rem 0;
  color: #333;
}

.method-description p {
  margin: 0.25rem 0;
  color: #666;
}

@media (max-width: 768px) {
  .form-row {
    grid-template-columns: 1fr;
  }
  
  .descriptions-list table {
    font-size: 0.8rem;
  }
  
  .descriptions-list th,
  .descriptions-list td {
    padding: 0.5rem;
  }
}
</style> 
