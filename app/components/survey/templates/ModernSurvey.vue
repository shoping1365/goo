<template>
  <div class="bg-white rounded-xl shadow-lg overflow-hidden">
    <div class="bg-gradient-to-r from-blue-500 to-blue-600 p-6 text-white text-center">
      <h2 class="text-2xl font-bold mb-2">نظرسنجی خرید</h2>
      <p class="text-blue-100">تجربه شما برای ما مهم است</p>
    </div>
    
    <div class="p-6">
      <!-- Order Info -->
      <div class="bg-gray-50 rounded-lg p-6 mb-6">
        <h3 class="font-semibold text-gray-800 mb-2">اطلاعات سفارش</h3>
        <div class="grid grid-cols-2 gap-6 text-sm">
          <div>
            <span class="text-gray-600">محصول:</span>
            <span class="font-medium text-gray-800 mr-2">{{ order.productName }}</span>
          </div>
          <div>
            <span class="text-gray-600">مبلغ:</span>
            <span class="font-medium text-gray-800 mr-2">{{ formatPrice(order.totalAmount) }}</span>
          </div>
        </div>
      </div>

      <!-- Rating -->
      <div class="mb-6">
        <h3 class="text-lg font-semibold text-gray-800 mb-4">چقدر از خرید خود راضی هستید؟</h3>
        <div class="flex justify-center space-x-2 space-x-reverse">
          <button 
            v-for="rating in 5" 
            :key="rating"
            :class="[
              'w-12 h-12 rounded-full transition-colors font-medium',
              selectedRating === rating 
                ? 'bg-blue-500 text-white' 
                : 'bg-gray-100 hover:bg-blue-100 text-gray-600 hover:text-blue-600'
            ]"
            @click="selectedRating = rating"
          >
            {{ rating }}
          </button>
        </div>
        <div class="text-center mt-2 text-sm text-gray-600">
          {{ getRatingText(selectedRating) }}
        </div>
      </div>

      <!-- Questions -->
      <div class="space-y-4 mb-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">کیفیت محصول چطور بود؟</label>
          <div class="flex space-x-2 space-x-reverse">
            <button 
              v-for="option in qualityOptions" 
              :key="option.value"
              :class="[
                'px-3 py-2 rounded-lg text-sm transition-colors',
                qualityRating === option.value 
                  ? 'bg-blue-500 text-white' 
                  : 'bg-gray-100 hover:bg-gray-200 text-gray-700'
              ]"
              @click="qualityRating = option.value"
            >
              {{ option.label }}
            </button>
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">سرعت ارسال چطور بود؟</label>
          <div class="flex space-x-2 space-x-reverse">
            <button 
              v-for="option in deliveryOptions" 
              :key="option.value"
              :class="[
                'px-3 py-2 rounded-lg text-sm transition-colors',
                deliveryRating === option.value 
                  ? 'bg-blue-500 text-white' 
                  : 'bg-gray-100 hover:bg-gray-200 text-gray-700'
              ]"
              @click="deliveryRating = option.value"
            >
              {{ option.label }}
            </button>
          </div>
        </div>
      </div>

      <!-- Comment -->
      <div class="mb-6">
        <label class="block text-sm font-medium text-gray-700 mb-2">نظر شما:</label>
        <textarea 
          v-model="comment"
          class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500" 
          rows="3" 
          placeholder="تجربه خود را با ما به اشتراک بگذارید..."
        ></textarea>
      </div>

      <!-- Submit Button -->
      <button 
        :disabled="!selectedRating || submitting"
        class="w-full bg-blue-500 hover:bg-blue-600 disabled:bg-gray-400 text-white font-medium py-3 px-4 rounded-lg transition-colors flex items-center justify-center space-x-2 space-x-reverse"
        @click="submitSurvey"
      >
        <svg v-if="submitting" class="animate-spin w-4 h-4" fill="none" viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
        </svg>
        <span>{{ submitting ? 'در حال ارسال...' : 'ارسال نظر' }}</span>
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

interface Order {
  id: string
  customerName: string
  productName: string
  orderDate: string
  totalAmount: number
}

interface Props {
  order: Order
}

interface SurveyData {
  rating: number
  qualityRating: string
  deliveryRating: string
  comment: string
  submittedAt: string
}

defineProps<Props>()
const emit = defineEmits<{
  submit: [data: SurveyData]
}>()

const selectedRating = ref(0)
const qualityRating = ref('')
const deliveryRating = ref('')
const comment = ref('')
const submitting = ref(false)

const qualityOptions = [
  { value: 'excellent', label: 'عالی' },
  { value: 'good', label: 'خوب' },
  { value: 'average', label: 'متوسط' },
  { value: 'poor', label: 'ضعیف' }
]

const deliveryOptions = [
  { value: 'fast', label: 'سریع' },
  { value: 'normal', label: 'معمولی' },
  { value: 'slow', label: 'کند' }
]

const formatPrice = (price: number) => {
  return new Intl.NumberFormat('fa-IR').format(price) + ' تومان'
}

const getRatingText = (rating: number) => {
  const texts = {
    1: 'خیلی ضعیف',
    2: 'ضعیف',
    3: 'متوسط',
    4: 'خوب',
    5: 'عالی'
  }
  return texts[rating as keyof typeof texts] || ''
}

const submitSurvey = async () => {
  if (!selectedRating.value) return
  
  submitting.value = true
  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    emit('submit', {
      rating: selectedRating.value,
      qualityRating: qualityRating.value,
      deliveryRating: deliveryRating.value,
      comment: comment.value,
      submittedAt: new Date().toISOString()
    })
  } finally {
    submitting.value = false
  }
}
</script>

<style scoped>
.space-x-reverse > :not([hidden]) ~ :not([hidden]) {
  --tw-space-x-reverse: 1;
}
</style> 
