<template>
  <div class="bg-white rounded-2xl shadow-xl border border-gray-100 p-8 mb-8">
    <div class="flex items-center mb-6">
      <div class="p-3 bg-gradient-to-r from-pink-400 to-blue-600 rounded-xl shadow-lg">
        <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16.24 7.76a6 6 0 11-8.48 0M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707" />
        </svg>
      </div>
      <div class="mr-4">
        <h2 class="text-2xl font-bold text-gray-900">سوالات تخصصی</h2>
        <p class="text-gray-600 mt-1">پاسخ به سوالات تخصصی درباره تجربه خرید</p>
      </div>
    </div>

    <div class="space-y-8">
      <!-- رضایت از روش پرداخت -->
      <div class="bg-gradient-to-r from-blue-50 to-cyan-50 rounded-xl p-6 border border-blue-100">
        <div class="flex items-center mb-4">
          <div class="p-2 bg-blue-500 rounded-lg mr-3">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900">رضایت از روش پرداخت</h3>
        </div>
        <div class="flex gap-3">
          <button
              v-for="option in paymentSatisfactionOptions"
              :key="option.value"
              @click="answers.paymentSatisfaction = option.value"
              :class="['px-5 py-2 rounded-lg font-medium border transition', answers.paymentSatisfaction === option.value ? 'bg-blue-600 text-white border-blue-700' : 'bg-white text-gray-700 border-gray-300 hover:bg-blue-50']"
          >
            {{ option.label }}
          </button>
        </div>
      </div>

      <!-- رضایت از روش ارسال -->
      <div class="bg-gradient-to-r from-green-50 to-emerald-50 rounded-xl p-6 border border-green-100">
        <div class="flex items-center mb-4">
          <div class="p-2 bg-green-500 rounded-lg mr-3">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900">رضایت از روش ارسال</h3>
        </div>
        <div class="flex gap-3">
          <button
              v-for="option in shippingSatisfactionOptions"
              :key="option.value"
              @click="answers.shippingSatisfaction = option.value"
              :class="['px-5 py-2 rounded-lg font-medium border transition', answers.shippingSatisfaction === option.value ? 'bg-green-600 text-white border-green-700' : 'bg-white text-gray-700 border-gray-300 hover:bg-green-50']"
          >
            {{ option.label }}
          </button>
        </div>
      </div>

      <!-- کیفیت ارتباطات -->
      <div class="bg-gradient-to-r from-purple-50 to-violet-50 rounded-xl p-6 border border-purple-100">
        <div class="flex items-center mb-4">
          <div class="p-2 bg-purple-500 rounded-lg mr-3">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900">کیفیت ارتباطات (پیامک، ایمیل)</h3>
        </div>
        <div class="flex gap-3">
          <button
              v-for="option in communicationQualityOptions"
              :key="option.value"
              @click="answers.communicationQuality = option.value"
              :class="['px-5 py-2 rounded-lg font-medium border transition', answers.communicationQuality === option.value ? 'bg-purple-600 text-white border-purple-700' : 'bg-white text-gray-700 border-gray-300 hover:bg-purple-50']"
          >
            {{ option.label }}
          </button>
        </div>
      </div>

      <!-- دقت در اطلاعات محصول -->
      <div class="bg-gradient-to-r from-orange-50 to-amber-50 rounded-xl p-6 border border-orange-100">
        <div class="flex items-center mb-4">
          <div class="p-2 bg-orange-500 rounded-lg mr-3">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900">دقت در اطلاعات محصول</h3>
        </div>
        <div class="flex gap-3">
          <button
              v-for="option in productInfoAccuracyOptions"
              :key="option.value"
              @click="answers.productInfoAccuracy = option.value"
              :class="['px-5 py-2 rounded-lg font-medium border transition', answers.productInfoAccuracy === option.value ? 'bg-orange-600 text-white border-orange-700' : 'bg-white text-gray-700 border-gray-300 hover:bg-orange-50']"
          >
            {{ option.label }}
          </button>
        </div>
      </div>

      <!-- سرعت پاسخگویی پشتیبانی -->
      <div class="bg-gradient-to-r from-red-50 to-pink-50 rounded-xl p-6 border border-red-100">
        <div class="flex items-center mb-4">
          <div class="p-2 bg-red-500 rounded-lg mr-3">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 5.636l-3.536 3.536m0 5.656l3.536 3.536M9.172 9.172L5.636 5.636m3.536 9.192L5.636 18.364M12 2.25a9.75 9.75 0 109.75 9.75A9.75 9.75 0 0012 2.25z"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900">سرعت پاسخگویی پشتیبانی</h3>
        </div>
        <div class="flex gap-3">
          <button
              v-for="option in supportResponseSpeedOptions"
              :key="option.value"
              @click="answers.supportResponseSpeed = option.value"
              :class="['px-5 py-2 rounded-lg font-medium border transition', answers.supportResponseSpeed === option.value ? 'bg-red-600 text-white border-red-700' : 'bg-white text-gray-700 border-gray-300 hover:bg-red-50']"
          >
            {{ option.label }}
          </button>
        </div>
      </div>
    </div>

    <!-- خلاصه پاسخ‌ها -->
    <div class="mt-8 bg-gradient-to-r from-gray-50 to-slate-50 rounded-xl p-6 border border-gray-200">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">خلاصه پاسخ‌های تخصصی</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-6">
        <div v-for="(answer, key) in answers" :key="key" class="text-center p-3 bg-white rounded-lg border border-gray-200">
          <div class="text-sm text-gray-600 mb-1">{{ getQuestionName(key) }}</div>
          <div class="font-semibold text-gray-900">{{ getAnswerText(key, answer) || 'پاسخ داده نشده' }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

interface SpecializedAnswers {
  paymentSatisfaction: string
  shippingSatisfaction: string
  communicationQuality: string
  productInfoAccuracy: string
  supportResponseSpeed: string
}

const answers = ref<SpecializedAnswers>({
  paymentSatisfaction: '',
  shippingSatisfaction: '',
  communicationQuality: '',
  productInfoAccuracy: '',
  supportResponseSpeed: ''
})

const paymentSatisfactionOptions = [
  { value: 'excellent', label: 'عالی' },
  { value: 'good', label: 'خوب' },
  { value: 'average', label: 'متوسط' },
  { value: 'poor', label: 'ضعیف' }
]

const shippingSatisfactionOptions = [
  { value: 'excellent', label: 'عالی' },
  { value: 'good', label: 'خوب' },
  { value: 'average', label: 'متوسط' },
  { value: 'poor', label: 'ضعیف' }
]

const communicationQualityOptions = [
  { value: 'excellent', label: 'عالی' },
  { value: 'good', label: 'خوب' },
  { value: 'average', label: 'متوسط' },
  { value: 'poor', label: 'ضعیف' }
]

const productInfoAccuracyOptions = [
  { value: 'very_accurate', label: 'کاملاً دقیق' },
  { value: 'accurate', label: 'دقیق' },
  { value: 'average', label: 'متوسط' },
  { value: 'inaccurate', label: 'غیر دقیق' }
]

const supportResponseSpeedOptions = [
  { value: 'very_fast', label: 'خیلی سریع' },
  { value: 'fast', label: 'سریع' },
  { value: 'average', label: 'متوسط' },
  { value: 'slow', label: 'کند' }
]

const getQuestionName = (key: string): string => {
  const names = {
    paymentSatisfaction: 'روش پرداخت',
    shippingSatisfaction: 'روش ارسال',
    communicationQuality: 'ارتباطات',
    productInfoAccuracy: 'دقت اطلاعات',
    supportResponseSpeed: 'پاسخگویی'
  }
  return names[key] || key
}

const getAnswerText = (key: string, value: string): string => {
  const allOptions = [
    ...paymentSatisfactionOptions,
    ...shippingSatisfactionOptions,
    ...communicationQualityOptions,
    ...productInfoAccuracyOptions,
    ...supportResponseSpeedOptions
  ]
  const option = allOptions.find(opt => opt.value === value)
  return option ? option.label : ''
}
</script>

<style scoped>
button:focus {
  outline: 2px solid #6366f1;
  outline-offset: 2px;
}

/* Custom hover effects */
.hover\:bg-blue-50:hover {
  background-color: #eff6ff;
}

.hover\:bg-green-50:hover {
  background-color: #f0fdf4;
}

.hover\:bg-purple-50:hover {
  background-color: #faf5ff;
}

.hover\:bg-orange-50:hover {
  background-color: #fff7ed;
}

.hover\:bg-red-50:hover {
  background-color: #fef2f2;
}
</style>
