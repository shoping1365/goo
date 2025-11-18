<template>
  <div class="prohibited-products">
    <div class="settings-card">
      <div class="card-header">
        <h3>محصولات ممنوعه</h3>
        <div class="toggle-switch">
          <input type="checkbox" v-model="enabled" id="prohibitedToggle" />
          <label for="prohibitedToggle"></label>
        </div>
      </div>

      <div v-if="enabled" class="card-body">
        <div class="form-section">
          <h4>دسته‌بندی‌های ممنوعه</h4>
          <div class="category-list">
            <div v-for="(category, index) in prohibited.categories" :key="index" class="category-item">
              <div class="form-row">
                <div class="form-group">
                  <label>نام دسته‌بندی</label>
                  <input
                    v-model="category.name"
                    type="text"
                    class="form-input"
                    placeholder="نام دسته‌بندی"
                  />
                </div>
                <div class="form-group">
                  <label>دلیل ممنوعیت</label>
                  <select v-model="category.reason" class="form-input">
                    <option value="legal">ممنوعیت قانونی</option>
                    <option value="safety">مشکل ایمنی</option>
                    <option value="logistics">مشکل لجستیکی</option>
                    <option value="policy">سیاست شرکت</option>
                  </select>
                </div>
                <button @click="removeCategory(index)" class="remove-btn">حذف</button>
              </div>
              <div class="form-group">
                <label>توضیحات</label>
                <textarea
                  v-model="category.description"
                  class="form-input"
                  placeholder="توضیحات بیشتر"
                  rows="2"
                ></textarea>
              </div>
            </div>
            <button @click="addCategory" class="add-btn">+ افزودن دسته‌بندی</button>
          </div>
        </div>

        <div class="form-section">
          <h4>کلمات کلیدی ممنوعه</h4>
          <div class="keywords-list">
            <div v-for="(keyword, index) in prohibited.keywords" :key="index" class="keyword-item">
              <div class="form-row">
                <div class="form-group">
                  <label>کلمه کلیدی</label>
                  <input
                    v-model="keyword.value"
                    type="text"
                    class="form-input"
                    placeholder="کلمه کلیدی"
                  />
                </div>
                <div class="form-group">
                  <label>شدت ممنوعیت</label>
                  <select v-model="keyword.severity" class="form-input">
                    <option value="high">بالا</option>
                    <option value="medium">متوسط</option>
                    <option value="low">کم</option>
                  </select>
                </div>
                <button @click="removeKeyword(index)" class="remove-btn">حذف</button>
              </div>
            </div>
            <button @click="addKeyword" class="add-btn">+ افزودن کلمه کلیدی</button>
          </div>
        </div>

        <div class="form-section">
          <h4>تنظیمات اضافی</h4>
          <div class="additional-settings">
            <div class="form-group">
              <label>نحوه بررسی</label>
              <div class="check-methods">
                <label class="radio-label">
                  <input type="radio" v-model="prohibited.checkMethod" value="automatic" />
                  بررسی خودکار
                </label>
                <label class="radio-label">
                  <input type="radio" v-model="prohibited.checkMethod" value="manual" />
                  بررسی دستی
                </label>
                <label class="radio-label">
                  <input type="radio" v-model="prohibited.checkMethod" value="both" />
                  هر دو روش
                </label>
              </div>
            </div>

            <div class="form-group">
              <label>اقدام در صورت یافتن محصول ممنوعه</label>
              <select v-model="prohibited.action" class="form-input">
                <option value="reject">رد سفارش</option>
                <option value="warning">هشدار</option>
                <option value="review">بررسی دستی</option>
                <option value="custom">اقدام سفارشی</option>
              </select>
            </div>

            <div class="form-group">
              <label>پیام هشدار</label>
              <textarea
                v-model="prohibited.warningMessage"
                class="form-input"
                placeholder="پیام نمایشی برای محصولات ممنوعه"
                rows="3"
              ></textarea>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'

const props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({})
  }
})

const emit = defineEmits(['update:modelValue'])

const enabled = ref(props.modelValue.enabled ?? true)

const prohibited = ref({
  categories: props.modelValue.categories || [
    { name: 'مواد مخدر', reason: 'legal', description: 'ممنوعیت قانونی' },
    { name: 'سلاح گرم', reason: 'legal', description: 'ممنوعیت قانونی' },
    { name: 'مواد انفجاری', reason: 'safety', description: 'خطر ایمنی' }
  ],
  keywords: props.modelValue.keywords || [
    { value: 'مواد مخدر', severity: 'high' },
    { value: 'سلاح', severity: 'high' },
    { value: 'انفجاری', severity: 'high' }
  ],
  checkMethod: props.modelValue.checkMethod || 'automatic',
  action: props.modelValue.action || 'reject',
  warningMessage: props.modelValue.warningMessage || 'این محصول در لیست محصولات ممنوعه قرار دارد.'
})

function addCategory() {
  prohibited.value.categories.push({ name: '', reason: 'legal', description: '' })
}

function removeCategory(index) {
  prohibited.value.categories.splice(index, 1)
}

function addKeyword() {
  prohibited.value.keywords.push({ value: '', severity: 'medium' })
}

function removeKeyword(index) {
  prohibited.value.keywords.splice(index, 1)
}

// Watch for changes and emit updates
watch([enabled, prohibited], () => {
  emit('update:modelValue', {
    enabled: enabled.value,
    ...prohibited.value
  })
}, { deep: true })
</script>

<style scoped>
.prohibited-products {
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  overflow: hidden;
}

.settings-card {
  background: #fff;
}

.card-header {
  background: linear-gradient(135deg, #dc3545 0%, #c82333 100%);
  color: white;
  padding: 1.5rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.card-header h3 {
  margin: 0;
  font-size: 1.3rem;
}

.toggle-switch {
  position: relative;
  width: 60px;
  height: 30px;
}

.toggle-switch input {
  opacity: 0;
  width: 0;
  height: 0;
}

.toggle-switch label {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(255, 255, 255, 0.3);
  transition: 0.4s;
  border-radius: 30px;
}

.toggle-switch label:before {
  position: absolute;
  content: "";
  height: 22px;
  width: 22px;
  left: 4px;
  bottom: 4px;
  background-color: white;
  transition: 0.4s;
  border-radius: 50%;
}

.toggle-switch input:checked + label {
  background-color: #28a745;
}

.toggle-switch input:checked + label:before {
  transform: translateX(30px);
}

.card-body {
  padding: 2rem;
}

.form-section {
  margin-bottom: 2rem;
  padding: 1.5rem;
  background: #f8f9fa;
  border-radius: 8px;
  border: 1px solid #e9ecef;
}

.form-section h4 {
  margin: 0 0 1.5rem 0;
  color: #333;
  font-size: 1.1rem;
  border-bottom: 2px solid #dc3545;
  padding-bottom: 0.5rem;
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr auto;
  gap: 1.5rem;
  margin-bottom: 1rem;
  align-items: end;
}

.form-group {
  margin-bottom: 1rem;
}

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  color: #333;
  font-weight: 500;
  font-size: 0.9rem;
}

.form-input {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 0.9rem;
  font-family: inherit;
  transition: border-color 0.3s;
}

.form-input:focus {
  outline: none;
  border-color: #dc3545;
  box-shadow: 0 0 0 2px rgba(220, 53, 69, 0.25);
}

.category-item,
.keyword-item {
  background: white;
  border: 1px solid #e9ecef;
  border-radius: 8px;
  padding: 1rem;
  margin-bottom: 1rem;
  border-left: 4px solid #dc3545;
}

.remove-btn {
  background: #dc3545;
  color: white;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 6px;
  cursor: pointer;
  font-size: 0.8rem;
  align-self: end;
  transition: background-color 0.3s;
}

.remove-btn:hover {
  background: #c82333;
}

.add-btn {
  background: #28a745;
  color: white;
  border: none;
  padding: 0.75rem 1.5rem;
  border-radius: 6px;
  cursor: pointer;
  font-size: 0.9rem;
  margin-top: 1rem;
  transition: background-color 0.3s;
}

.add-btn:hover {
  background: #218838;
}

.radio-label {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  cursor: pointer;
  margin-bottom: 0.5rem;
  font-size: 0.9rem;
}

.radio-label input[type="radio"] {
  width: 16px;
  height: 16px;
}

.check-methods {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  margin-top: 0.5rem;
}

@media (max-width: 768px) {
  .form-row {
    grid-template-columns: 1fr;
  }

  .card-body {
    padding: 1rem;
  }

  .form-section {
    padding: 1rem;
  }
}
</style> 
