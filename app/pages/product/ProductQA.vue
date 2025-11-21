<template>
  <div class="space-y-4" dir="rtl">
    <!-- Header -->
    <div class="flex items-center justify-between">
      <h3 class="text-lg font-bold text-gray-900 flex items-center gap-2">
        <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
        </svg>
        پرسش و پاسخ
        <span v-if="questions.length > 0" class="text-sm text-gray-500">({{ questions.length }})</span>
      </h3>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="space-y-3">
      <div v-for="i in 3" :key="i" class="bg-white rounded-lg border border-gray-200 px-4 py-4">
        <div class="animate-pulse">
          <div class="h-4 bg-gray-200 rounded w-3/4 mb-2"></div>
          <div class="h-3 bg-gray-200 rounded w-1/2"></div>
        </div>
      </div>
    </div>

    <!-- Questions List -->
    <div v-else-if="questions?.length" class="space-y-3">
      <div
        v-for="(question, index) in questions"
        :key="question.id || index"
        class="bg-white rounded-lg border border-gray-200 overflow-hidden"
      >
        <div class="px-6 py-4">
          <h4 class="font-medium text-gray-900 mb-2">
            {{ question.question }}
          </h4>
          <div class="flex items-center gapx-4 py-4 text-xs text-gray-500 mb-2">
            <span>{{ question.customer_name || 'کاربر ناشناس' }}</span>
            <span>{{ formatDate(question.created_at) }}</span>
            <span v-if="question.category" class="bg-blue-100 text-blue-700 px-2 py-1 rounded">
              {{ question.category }}
            </span>
          </div>
          <div v-if="question.answer" class="text-gray-700 leading-relaxed">
            {{ question.answer }}
          </div>
          <div v-else class="text-gray-500 italic text-sm">
            هنوز پاسخی برای این سوال ثبت نشده است.
          </div>
          <div v-if="question.answered_by" class="text-xs text-gray-500 mt-2">
            پاسخ داده شده توسط: {{ question.answered_by }}
            <span v-if="question.answered_at"> در {{ formatDate(question.answered_at) }}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Empty State -->
    <div v-else-if="!loading" class="bg-gray-50 rounded-lg p-12 text-center">
      <svg class="w-16 h-16 text-gray-300 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
      </svg>
      <h3 class="text-lg font-medium text-gray-900 mb-2">سوالی ثبت نشده است</h3>
      <p class="text-gray-500">هنوز سوالی برای این محصول ثبت نشده است.</p>
    </div>

    <!-- Ask Question Form -->
    <div class="bg-blue-50 rounded-lg px-4 py-4 border border-blue-200">
      <h4 class="font-medium text-blue-900 mb-4 flex items-center gap-2">
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
        </svg>
        سوال خود را بپرسید
      </h4>
      
      <form class="space-y-4" @submit.prevent="submitQuestion">
        <div>
          <label for="question" class="block text-sm font-medium text-blue-900 mb-2">سوال شما</label>
          <textarea
            id="question"
            v-model="newQuestion.question"
            rows="3"
            class="w-full px-3 py-2 border border-blue-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 resize-none"
            placeholder="سوال خود را اینجا بنویسید..."
            required
          ></textarea>
        </div>
        
        <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
          <div>
            <label for="name" class="block text-sm font-medium text-blue-900 mb-2">نام (اختیاری)</label>
            <input
              id="name"
              v-model="newQuestion.customer_name"
              type="text"
              class="w-full px-3 py-2 border border-blue-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="نام شما"
            />
          </div>
          
          <div>
            <label for="mobile" class="block text-sm font-medium text-blue-900 mb-2">شماره موبایل (اختیاری)</label>
            <input
              id="mobile"
              v-model="newQuestion.customer_mobile"
              type="tel"
              class="w-full px-3 py-2 border border-blue-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="09xxxxxxxxx"
            />
          </div>
        </div>
        
        <div class="flex items-center gapx-4 py-4">
          <label class="flex items-center gap-2">
            <input
              v-model="newQuestion.is_anonymous"
              type="checkbox"
              class="rounded border-blue-300 text-blue-600 focus:ring-blue-500"
            />
            <span class="text-sm text-blue-900">ناشناس بمانم</span>
          </label>
        </div>
        
        <div class="flex gap-3">
          <button
            type="submit"
            :disabled="submitting"
            class="bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 text-white px-6 py-2 rounded-lg text-sm font-medium transition-colors flex items-center gap-2"
          >
            <svg v-if="submitting" class="w-4 h-4 animate-spin" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            {{ submitting ? 'در حال ثبت...' : 'ثبت سوال' }}
          </button>
          <button
            type="button"
            class="bg-white border border-blue-300 text-blue-700 hover:bg-blue-50 px-6 py-2 rounded-lg text-sm font-medium transition-colors"
            @click="resetForm"
          >
            پاک کردن
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { onMounted, reactive, ref } from 'vue'

const props = defineProps({
  product: {
    type: Object,
    required: true
  }
})

// State
const questions = ref([])
const loading = ref(false)
const submitting = ref(false)

const newQuestion = reactive({
  question: '',
  customer_name: '',
  customer_mobile: '',
  is_anonymous: false
})

// Methods
const fetchQuestions = async () => {
  if (!props.product?.id) return
  
  try {
    loading.value = true
    const response = await $fetch(`/api/products/${props.product.id}/questions`)
    questions.value = response || []
  } catch (error) {
    console.error('Error fetching questions:', error)
  } finally {
    loading.value = false
  }
}

const submitQuestion = async () => {
  if (!props.product?.id || !newQuestion.question.trim()) return
  
  try {
    submitting.value = true
    await $fetch('/api/questions', {
      method: 'POST',
      body: {
        question: newQuestion.question.trim(),
        product_id: props.product.id,
        customer_name: newQuestion.customer_name.trim(),
        customer_mobile: newQuestion.customer_mobile.trim(),
        is_anonymous: newQuestion.is_anonymous
      }
    })
    
    // Reset form
    resetForm()
    
    // Refresh questions
    await fetchQuestions()
    
    // Show success message
    alert('سوال شما با موفقیت ثبت شد و پس از بررسی پاسخ داده خواهد شد.')
  } catch (error) {
    console.error('Error submitting question:', error)
    alert('خطا در ثبت سوال. لطفاً دوباره تلاش کنید.')
  } finally {
    submitting.value = false
  }
}

const resetForm = () => {
  newQuestion.question = ''
  newQuestion.customer_name = ''
  newQuestion.customer_mobile = ''
  newQuestion.is_anonymous = false
}

const formatDate = (dateString) => {
  if (!dateString) return ''
  return new Date(dateString).toLocaleDateString('fa-IR')
}

// Lifecycle
onMounted(() => {
  fetchQuestions()
})
</script> 