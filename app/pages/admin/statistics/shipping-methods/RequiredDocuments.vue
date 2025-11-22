<template>
  <div class="required-documents">
    <div class="section-header">
      <h3>مستندات مورد نیاز</h3>
      <p>مدیریت مستندات مورد نیاز برای ارسال بین‌المللی</p>
    </div>

    <!-- Documents List -->
    <div class="documents-list">
      <h4>لیست مستندات</h4>
      <table>
        <thead>
        <tr>
          <th>نام مستند</th>
          <th>نوع</th>
          <th>کشورهای مورد نیاز</th>
          <th>ضروری/اختیاری</th>
          <th>توضیحات</th>
          <th>وضعیت</th>
          <th>عملیات</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="document in documents" :key="document.id">
          <td>{{ document.name }}</td>
          <td>{{ document.type }}</td>
          <td>{{ document.requiredCountries }}</td>
          <td>
            <span :class="['requirement', document.requirement]">{{ document.requirementText }}</span>
          </td>
          <td>{{ document.description }}</td>
          <td>
            <span :class="['status', document.status]">{{ document.statusText }}</span>
          </td>
          <td>
            <button class="btn-edit" @click="editDocument(document)">ویرایش</button>
            <button class="btn-delete" @click="deleteDocument(document.id)">حذف</button>
          </td>
        </tr>
        </tbody>
      </table>
    </div>

    <!-- Add New Document -->
    <div class="add-document-section">
      <h4>افزودن مستند جدید</h4>
      <form @submit.prevent="addNewDocument">
        <div class="form-row">
          <div class="form-group">
            <label>نام مستند:</label>
            <input v-model="newDocument.name" type="text" required>
          </div>
          <div class="form-group">
            <label>نوع:</label>
            <select v-model="newDocument.type" required>
              <option value="">انتخاب کنید</option>
              <option value="invoice">فاکتور</option>
              <option value="certificate">گواهی</option>
              <option value="permit">مجوز</option>
              <option value="declaration">اظهارنامه</option>
              <option value="other">سایر</option>
            </select>
          </div>
        </div>
        <div class="form-row">
          <div class="form-group">
            <label>کشورهای مورد نیاز:</label>
            <select v-model="newDocument.requiredCountries" multiple>
              <option value="usa">آمریکا</option>
              <option value="uk">انگلیس</option>
              <option value="germany">آلمان</option>
              <option value="france">فرانسه</option>
              <option value="australia">استرالیا</option>
            </select>
          </div>
          <div class="form-group">
            <label>ضروری/اختیاری:</label>
            <select v-model="newDocument.requirement" required>
              <option value="required">ضروری</option>
              <option value="optional">اختیاری</option>
            </select>
          </div>
        </div>
        <div class="form-group">
          <label>توضیحات:</label>
          <textarea v-model="newDocument.description" rows="3" placeholder="توضیحات کامل مستند"></textarea>
        </div>
        <button type="submit" class="btn-primary">افزودن مستند</button>
      </form>
    </div>

    <!-- Document Templates -->
    <div class="document-templates">
      <h4>قالب‌های مستندات</h4>
      <div class="templates-grid">
        <div v-for="template in documentTemplates" :key="template.id" class="template-card">
          <div class="template-header">
            <h5>{{ template.name }}</h5>
            <span :class="['template-type', template.type]">{{ template.typeText }}</span>
          </div>
          <div class="template-content">
            <p>{{ template.description }}</p>
            <div class="template-countries">
              <strong>کشورها:</strong> {{ template.countries }}
            </div>
          </div>
          <div class="template-actions">
            <button class="btn-download" @click="downloadTemplate(template)">دانلود</button>
            <button class="btn-preview" @click="previewTemplate(template)">پیش‌نمایش</button>
          </div>
        </div>
      </div>
    </div>

    <!-- Document Checklist -->
    <div class="document-checklist">
      <h4>چک‌لیست مستندات</h4>
      <div class="checklist-form">
        <div class="form-row">
          <div class="form-group">
            <label>کشور مقصد:</label>
            <select v-model="checklist.country" @change="updateChecklist">
              <option value="">انتخاب کنید</option>
              <option value="usa">آمریکا</option>
              <option value="uk">انگلیس</option>
              <option value="germany">آلمان</option>
              <option value="france">فرانسه</option>
              <option value="australia">استرالیا</option>
            </select>
          </div>
        </div>
        <div v-if="checklist.country" class="checklist-items">
          <h5>مستندات مورد نیاز برای {{ getCountryName(checklist.country) }}:</h5>
          <div v-for="item in checklist.items" :key="item.id" class="checklist-item">
            <label class="checkbox-label">
              <input v-model="item.completed" type="checkbox">
              <span class="checkmark"></span>
              <span class="item-text">{{ item.name }}</span>
              <span :class="['requirement-badge', item.requirement]">{{ item.requirementText }}</span>
            </label>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">

declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
</script>

<script setup lang="ts">
import { reactive, ref } from 'vue';

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

const documents = ref([
  {
    id: 1,
    name: 'فاکتور تجاری',
    type: 'invoice',
    requiredCountries: 'همه کشورها',
    requirement: 'required',
    requirementText: 'ضروری',
    description: 'فاکتور رسمی با جزئیات کامل محصولات',
    status: 'active',
    statusText: 'فعال'
  },
  {
    id: 2,
    name: 'گواهی مبدا',
    type: 'certificate',
    requiredCountries: 'اتحادیه اروپا',
    requirement: 'required',
    requirementText: 'ضروری',
    description: 'گواهی مبدا کالا برای کشورهای اتحادیه اروپا',
    status: 'active',
    statusText: 'فعال'
  }
])

const newDocument = reactive({
  name: '',
  type: '',
  requiredCountries: [],
  requirement: 'required',
  description: ''
})

const documentTemplates = ref([
  {
    id: 1,
    name: 'قالب فاکتور تجاری',
    type: 'invoice',
    typeText: 'فاکتور',
    description: 'قالب استاندارد فاکتور تجاری برای ارسال بین‌المللی',
    countries: 'همه کشورها'
  },
  {
    id: 2,
    name: 'قالب گواهی مبدا',
    type: 'certificate',
    typeText: 'گواهی',
    description: 'قالب گواهی مبدا برای کشورهای اتحادیه اروپا',
    countries: 'اتحادیه اروپا'
  }
])

const checklist = reactive({
  country: '',
  items: []
})

const countryDocuments = {
  usa: [
    { id: 1, name: 'فاکتور تجاری', requirement: 'required', requirementText: 'ضروری', completed: false },
    { id: 2, name: 'برچسب آدرس', requirement: 'required', requirementText: 'ضروری', completed: false },
    { id: 3, name: 'گواهی FDA', requirement: 'optional', requirementText: 'اختیاری', completed: false }
  ],
  uk: [
    { id: 1, name: 'فاکتور تجاری', requirement: 'required', requirementText: 'ضروری', completed: false },
    { id: 2, name: 'گواهی مبدا', requirement: 'required', requirementText: 'ضروری', completed: false },
    { id: 3, name: 'برچسب آدرس', requirement: 'required', requirementText: 'ضروری', completed: false }
  ]
}

function getCountryName(code) {
  const names = {
    usa: 'آمریکا',
    uk: 'انگلیس',
    germany: 'آلمان',
    france: 'فرانسه',
    australia: 'استرالیا'
  }
  return names[code] || code
}

function updateChecklist() {
  if (checklist.country && countryDocuments[checklist.country]) {
    checklist.items = [...countryDocuments[checklist.country]]
  } else {
    checklist.items = []
  }
}

function editDocument(_document) {
}

function deleteDocument(_id) {
}

function addNewDocument() {
  // Reset form
  Object.assign(newDocument, {
    name: '',
    type: '',
    requiredCountries: [],
    requirement: 'required',
    description: ''
  })
}

function downloadTemplate(_template) {
}

function previewTemplate(_template) {
}
</script>

<style scoped>
.required-documents {
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

.documents-list {
  margin-bottom: 2rem;
}

.documents-list h4 {
  margin: 0 0 1rem 0;
  font-size: 1.2rem;
  color: #333;
}

.documents-list table {
  width: 100%;
  border-collapse: collapse;
  background: white;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.documents-list th,
.documents-list td {
  padding: 1rem;
  text-align: right;
  border-bottom: 1px solid #e9ecef;
}

.documents-list th {
  background: #f8f9fa;
  font-weight: 600;
  color: #555;
}

.requirement,
.status {
  padding: 0.25rem 0.5rem;
  border-radius: 12px;
  font-size: 0.8rem;
}

.requirement.required,
.status.active {
  background: #d4edda;
  color: #155724;
}

.requirement.optional {
  background: #fff3cd;
  color: #856404;
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

.add-document-section {
  background: #f8f9fa;
  border-radius: 12px;
  padding: 1.5rem;
  margin-bottom: 2rem;
}

.add-document-section h4 {
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

.btn-primary {
  background: #007bff;
  color: white;
  border: none;
  padding: 0.75rem 1.5rem;
  border-radius: 6px;
  cursor: pointer;
  font-size: 0.9rem;
}

.document-templates {
  margin-bottom: 2rem;
}

.document-templates h4 {
  margin: 0 0 1rem 0;
  font-size: 1.2rem;
  color: #333;
}

.templates-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 1rem;
}

.template-card {
  background: white;
  border-radius: 8px;
  padding: 1rem;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.template-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 0.5rem;
}

.template-header h5 {
  margin: 0;
  color: #333;
}

.template-type {
  padding: 0.25rem 0.5rem;
  border-radius: 12px;
  font-size: 0.8rem;
  background: #e9ecef;
  color: #495057;
}

.template-content p {
  margin: 0 0 0.5rem 0;
  color: #666;
}

.template-countries {
  font-size: 0.9rem;
  color: #555;
}

.template-actions {
  display: flex;
  gap: 0.5rem;
  margin-top: 1rem;
}

.btn-download,
.btn-preview {
  padding: 0.5rem 1rem;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 0.8rem;
}

.btn-download {
  background: #28a745;
  color: white;
}

.btn-preview {
  background: #17a2b8;
  color: white;
}

.document-checklist {
  background: white;
  border-radius: 12px;
  padding: 1.5rem;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.document-checklist h4 {
  margin: 0 0 1rem 0;
  font-size: 1.2rem;
  color: #333;
}

.checklist-items h5 {
  margin: 0 0 1rem 0;
  color: #333;
}

.checklist-item {
  margin-bottom: 0.5rem;
}

.checkbox-label {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  cursor: pointer;
}

.checkbox-label input[type="checkbox"] {
  display: none;
}

.checkmark {
  width: 18px;
  height: 18px;
  border: 2px solid #ddd;
  border-radius: 3px;
  position: relative;
}

.checkbox-label input[type="checkbox"]:checked + .checkmark {
  background: #007bff;
  border-color: #007bff;
}

.checkbox-label input[type="checkbox"]:checked + .checkmark::after {
  content: '✓';
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  color: white;
  font-size: 12px;
}

.item-text {
  flex: 1;
  color: #333;
}

.requirement-badge {
  padding: 0.25rem 0.5rem;
  border-radius: 12px;
  font-size: 0.8rem;
}

.requirement-badge.required {
  background: #d4edda;
  color: #155724;
}

.requirement-badge.optional {
  background: #fff3cd;
  color: #856404;
}

@media (max-width: 768px) {
  .form-row {
    grid-template-columns: 1fr;
  }
  
  .templates-grid {
    grid-template-columns: 1fr;
  }
  
  .documents-list table {
    font-size: 0.8rem;
  }
  
  .documents-list th,
  .documents-list td {
    padding: 0.5rem;
  }
}
</style>
