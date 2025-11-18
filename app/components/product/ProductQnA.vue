<template>
  <div class="qna-root">
    <h3 class="qna-title">پرسش‌های کاربران</h3>
    <div v-if="loading" class="qna-loading">در حال بارگذاری...</div>
    <div v-else>
      <div v-if="questions.length === 0" class="qna-empty">هنوز پرسشی ثبت نشده است.</div>
      <div class="qna-list">
        <div v-for="q in questions" :key="q.id" class="qna-card">
          <div class="qna-question-row">
            <span class="qna-question-label">سؤال:</span>
            <span class="qna-question-text">{{ q.question }}</span>
          </div>
          <div v-if="q.answer" class="qna-answer-row">
            <span class="qna-answer-label">پاسخ ادمین:</span>
            <span class="qna-answer-text">{{ q.answer }}</span>
          </div>
        </div>
      </div>
    </div>
    <div class="qna-form-card">
      <h4 class="qna-form-title">ثبت پرسش جدید</h4>
      <div v-if="!isAuthenticated && !isLoading">
        <input v-model="guestName" class="qna-input" placeholder="نام و نام خانوادگی" />
        <input v-model="guestPhone" class="qna-input" placeholder="شماره تلفن" />
      </div>
      <textarea v-model="newQuestion" class="qna-textarea" placeholder="پرسش خود را بنویسید..." rows="3"></textarea>
      <div class="qna-public-row">
        <label class="qna-public-label">
          <input type="checkbox" v-model="isPublic" />
          <span>به صورت عمومی در سایت منتشر شود</span>
        </label>
      </div>
      <button @click="submitQuestion" :disabled="submitting || !newQuestion || (!isAuthenticated && (!guestName || !guestPhone))" class="qna-submit-btn">
        <span v-if="submitting">در حال ارسال...</span>
        <span v-else>ارسال پرسش</span>
      </button>
      <div v-if="error" class="qna-error">{{ error }}</div>
      <div v-if="success" class="qna-success">پرسش شما ثبت شد و پس از تأیید نمایش داده خواهد شد.</div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useNuxtApp } from '#app'


const props = defineProps({
  productId: {
    type: [String, Number],
    required: true
  }
})

// احراز هویت - غیرفعال شده
const user = null
const isAuthenticated = false
const fetchUser = null
const isLoading = ref(false)

const questions = ref([])
const loading = ref(true)
const newQuestion = ref('')
const submitting = ref(false)
const error = ref('')
const success = ref(false)
const guestName = ref('')
const guestPhone = ref('')
const isPublic = ref(true)

const fetchQuestions = async () => {
  loading.value = true
  error.value = ''
  try {
    const { data, error: fetchError } = await useFetch(`/api/products/${props.productId}/questions`)
    if (fetchError.value) throw fetchError.value
    questions.value = data.value || []
  } catch (e) {
    console.error('Error fetching questions:', e)
    error.value = 'خطا در دریافت پرسش‌ها'
  } finally {
    loading.value = false
  }
}

const submitQuestion = async () => {
  if (!newQuestion.value) return
  if (!isAuthenticated.value) {
    if (!guestName.value || !guestPhone.value) {
      error.value = 'نام و شماره تلفن الزامی است.'
      return
    }
  }
  submitting.value = true
  error.value = ''
  success.value = false
  try {
    const body = {
      question: newQuestion.value,
      is_anonymous: !isPublic.value,
    }
    
    if (!isAuthenticated.value) {
      body.customer_name = guestName.value
      body.customer_mobile = guestPhone.value
    }
    
    const { error: postError } = await useFetch(`/api/products/${props.productId}/questions`, {
      method: 'POST',
      body
    })
    if (postError.value) throw postError.value
    success.value = true
    newQuestion.value = ''
    guestName.value = ''
    guestPhone.value = ''
    isPublic.value = true
    await fetchQuestions()
  } catch (e) {
    console.error('Error submitting question:', e)
    error.value = 'خطا در ثبت پرسش'
  } finally {
    submitting.value = false
  }
}

// یک‌بار دریافت نام (اگر کاربر لاگین است و نام ندارد)
const ensureProfileNameOnce = async () => {
  try {
    if (!isAuthenticated.value) return
    // اطمینان از تازه بودن اطلاعات کاربر
    if (!user.value) {
      await fetchUser()
    }
    if (!user.value?.id) return
    const askedKey = `profile-name-asked:${user.value.id}`
    if (import.meta.client && localStorage.getItem(askedKey) === '1') return
    const currentName = user.value?.name || ''
    const currentMobile = user.value?.mobile || ''
    // اگر نام خالی است، یک‌بار بپرس
      if (!currentName || String(currentName).trim() === '') {
      if (!import.meta.client) return
      const fullName = window.prompt('لطفاً نام و نام خانوادگی خود را وارد کنید:') || ''
      if (fullName.trim()) {
        const { $fetch } = useNuxtApp()
        await $fetch(`/api/users/${user.value.id}`, {
          method: 'PUT',
          body: { name: fullName.trim() }
        })
        await fetchUser()
      }
      // حتی در صورت انصراف، یک‌بار پرسیده شده است
      if (import.meta.client) localStorage.setItem(askedKey, '1')
    }
    // اگر موبایل خالی بود (بعید)، یک‌بار بپرسیم
      if (!currentMobile || String(currentMobile).trim() === '') {
      if (!import.meta.client) return
      const mobile = window.prompt('شماره موبایل خود را وارد کنید:') || ''
      if (mobile.trim()) {
        const { $fetch } = useNuxtApp()
        await $fetch(`/api/users/${user.value.id}`, {
          method: 'PUT',
          body: { mobile: mobile.trim() }
        })
        await fetchUser()
      }
    }
  } catch (e) {
    // نادیده گرفتن خطا برای جلوگیری از اختلال در تجربه کاربری
    console.error('ensureProfileNameOnce error:', e)
  }
}

onMounted(async () => {
  isLoading.value = true
  try {
    await fetchQuestions()
    await ensureProfileNameOnce()
  } finally {
    isLoading.value = false
  }
})
</script>

<style scoped>
.qna-root {
  max-width: 600px;
  margin: 0 auto;
  padding: 2rem 0 3rem 0;
}
.qna-title {
  font-size: 1.4rem;
  font-weight: bold;
  margin-bottom: 1.5rem;
  text-align: center;
  color: #2563eb;
}
.qna-loading {
  text-align: center;
  color: #888;
  margin: 2rem 0;
}
.qna-empty {
  text-align: center;
  color: #aaa;
  margin: 2rem 0;
}
.qna-list {
  display: flex;
  flex-direction: column;
  gap: 1.2rem;
  margin-bottom: 2.5rem;
}
.qna-card {
  background: #f9fafb;
  border-radius: 12px;
  box-shadow: 0 2px 8px 0 #0001;
  padding: 1.1rem 1.2rem;
  border: 1px solid #e5e7eb;
  transition: box-shadow 0.2s;
}
.qna-card:hover {
  box-shadow: 0 4px 16px 0 #0002;
}
.qna-question-row {
  display: flex;
  align-items: flex-start;
  gap: 0.5rem;
  margin-bottom: 0.5rem;
}
.qna-question-label {
  color: #2563eb;
  font-weight: 500;
  min-width: 48px;
}
.qna-question-text {
  color: #222;
  font-size: 1rem;
  font-weight: 400;
}
.qna-answer-row {
  display: flex;
  align-items: flex-start;
  gap: 0.5rem;
  background: #e0f2fe;
  border-radius: 8px;
  padding: 0.5rem 0.7rem;
  margin-top: 0.3rem;
}
.qna-answer-label {
  color: #0284c7;
  font-weight: 500;
  min-width: 80px;
}
.qna-answer-text {
  color: #0369a1;
  font-size: 0.98rem;
}
.qna-form-card {
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 2px 8px 0 #0001;
  padding: 1.5rem 1.2rem 1.2rem 1.2rem;
  border: 1px solid #e5e7eb;
  margin-top: 2rem;
}
.qna-form-title {
  font-size: 1.1rem;
  font-weight: 500;
  margin-bottom: 0.7rem;
  color: #2563eb;
}
.qna-textarea {
  width: 100%;
  border-radius: 8px;
  border: 1px solid #cbd5e1;
  padding: 0.7rem;
  font-size: 1rem;
  margin-bottom: 0.7rem;
  resize: vertical;
  background: #f8fafc;
  transition: border 0.2s;
}
.qna-textarea:focus {
  border: 1.5px solid #2563eb;
  outline: none;
  background: #fff;
}
.qna-submit-btn {
  background: linear-gradient(90deg, #2563eb 60%, #0284c7 100%);
  color: #fff;
  border: none;
  border-radius: 8px;
  padding: 0.6rem 1.5rem;
  font-size: 1rem;
  font-weight: 500;
  cursor: pointer;
  transition: background 0.2s, box-shadow 0.2s;
  box-shadow: 0 1px 4px 0 #0001;
}
.qna-submit-btn:disabled {
  background: #cbd5e1;
  color: #fff;
  cursor: not-allowed;
}
.qna-error {
  color: #dc2626;
  margin-top: 0.7rem;
  text-align: right;
}
.qna-success {
  color: #16a34a;
  margin-top: 0.7rem;
  text-align: right;
}
.qna-input {
  width: 100%;
  border-radius: 8px;
  border: 1px solid #cbd5e1;
  padding: 0.7rem;
  font-size: 1rem;
  margin-bottom: 0.7rem;
  background: #f8fafc;
  transition: border 0.2s;
}
.qna-input:focus {
  border: 1.5px solid #2563eb;
  outline: none;
  background: #fff;
}
.qna-public-row {
  margin: 0.7rem 0 1rem 0;
  display: flex;
  align-items: center;
}
.qna-public-label {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-size: 0.98rem;
  color: #2563eb;
  font-weight: 500;
}
</style> 