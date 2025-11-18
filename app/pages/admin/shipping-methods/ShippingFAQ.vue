<template>
  <div class="shipping-faq">
    <div class="section-header">
      <h3>سوالات متداول (FAQ)</h3>
      <p>پاسخ به سوالات رایج درباره روش‌های ارسال</p>
    </div>

    <!-- FAQ List -->
    <div class="faq-list">
      <h4>لیست سوالات متداول</h4>
      <div v-for="faq in faqs" :key="faq.id" class="faq-item">
        <div class="faq-question" @click="faq.open = !faq.open">
          <span>{{ faq.question }}</span>
          <span class="toggle-icon">{{ faq.open ? '-' : '+' }}</span>
        </div>
        <div v-if="faq.open" class="faq-answer">
          {{ faq.answer }}
        </div>
      </div>
    </div>

    <!-- Add New FAQ -->
    <div class="add-faq-section">
      <h4>افزودن سوال جدید</h4>
      <form @submit.prevent="addNewFaq">
        <div class="form-group">
          <label>سوال:</label>
          <input v-model="newFaq.question" type="text" required>
        </div>
        <div class="form-group">
          <label>پاسخ:</label>
          <textarea v-model="newFaq.answer" rows="3" required></textarea>
        </div>
        <button type="submit" class="btn-primary">افزودن سوال</button>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'

const faqs = ref([
  { id: 1, question: 'چگونه می‌توانم روش ارسال را تغییر دهم؟', answer: 'در صفحه سبد خرید، گزینه تغییر روش ارسال را انتخاب کنید.', open: false },
  { id: 2, question: 'آیا ارسال سریع شامل همه شهرها می‌شود؟', answer: 'خیر، ارسال سریع فقط برای شهرهای منتخب فعال است.', open: false }
])

const newFaq = reactive({
  question: '',
  answer: ''
})

const addNewFaq = () => {
  // Add new FAQ logic
}
</script>

<style scoped>
.shipping-faq {
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

.faq-list {
  background: white;
  border-radius: 8px;
  padding: 20px;
  margin-bottom: 20px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.faq-item {
  border-bottom: 1px solid #ecf0f1;
  padding: 12px 0;
}

.faq-question {
  cursor: pointer;
  font-weight: 600;
  color: #007bff;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.faq-answer {
  margin-top: 8px;
  color: #2c3e50;
}

.add-faq-section {
  background: white;
  border-radius: 8px;
  padding: 20px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.form-group {
  margin-bottom: 15px;
}

.form-group label {
  display: block;
  margin-bottom: 5px;
  font-weight: 500;
  color: #2c3e50;
}

.form-group input,
.form-group textarea {
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

.toggle-icon {
  font-size: 20px;
  color: #888;
}
</style>
