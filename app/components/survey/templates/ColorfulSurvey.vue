<template>
  <div class="bg-gradient-to-br from-purple-400 via-pink-500 to-red-500 rounded-xl shadow-lg overflow-hidden">
    <div class="bg-white bg-opacity-10 backdrop-blur-sm p-6 text-white text-center">
      <div class="w-20 h-20 bg-white bg-opacity-20 rounded-full mx-auto mb-4 flex items-center justify-center">
        <svg class="w-10 h-10" fill="currentColor" viewBox="0 0 20 20">
          <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clip-rule="evenodd"></path>
        </svg>
      </div>
      <h2 class="text-3xl font-bold mb-2">نظرسنجی</h2>
      <p class="text-purple-100 text-lg">تجربه شما مهم است!</p>
    </div>
    
    <div class="p-6 bg-white bg-opacity-95 backdrop-blur-sm">
      <!-- Order Info -->
      <div class="bg-gradient-to-r from-purple-100 to-pink-100 rounded-xl p-6 mb-6 border border-purple-200">
        <h3 class="font-bold text-purple-800 mb-2">اطلاعات سفارش</h3>
        <div class="grid grid-cols-2 gap-6 text-sm">
          <div>
            <span class="text-purple-600 font-medium">محصول:</span>
            <span class="font-bold text-purple-800 mr-2">{{ order.productName }}</span>
          </div>
          <div>
            <span class="text-purple-600 font-medium">مبلغ:</span>
            <span class="font-bold text-purple-800 mr-2">{{ formatPrice(order.totalAmount) }}</span>
          </div>
        </div>
      </div>

      <!-- Rating -->
      <div class="mb-6">
        <h3 class="text-xl font-bold text-gray-800 mb-4 text-center">چقدر راضی هستید؟</h3>
        <div class="grid grid-cols-5 gap-2">
          <button 
            v-for="rating in 5" 
            :key="rating"
            :class="[
              'w-12 h-12 rounded-full text-white font-bold transition-all transform hover:scale-110',
              selectedRating === rating 
                ? 'shadow-lg transform scale-110' 
                : ''
            ]"
            :style="{
              backgroundColor: rating === 1 ? '#ef4444' : 
                             rating === 2 ? '#f97316' : 
                             rating === 3 ? '#eab308' : 
                             rating === 4 ? '#22c55e' : '#8b5cf6'
            }"
            @click="selectedRating = rating"
          >
            {{ rating }}
          </button>
        </div>
        <div class="text-center mt-3 text-sm font-medium text-gray-700">
          {{ getRatingText(selectedRating) }}
        </div>
      </div>

      <!-- Questions -->
      <div class="space-y-4 mb-6">
        <div>
          <label class="block text-sm font-bold text-gray-700 mb-2">کیفیت محصول:</label>
          <div class="grid grid-cols-2 gap-2">
            <button 
              v-for="option in qualityOptions" 
              :key="option.value"
              :class="[
                'px-4 py-3 rounded-xl text-sm font-bold transition-all transform hover:scale-105',
                qualityRating === option.value 
                  ? 'bg-gradient-to-r from-purple-500 to-pink-500 text-white shadow-lg' 
                  : 'bg-gray-100 hover:bg-purple-100 text-gray-700 hover:text-purple-700'
              ]"
              @click="qualityRating = option.value"
            >
              {{ option.label }}
            </button>
          </div>
        </div>

        <div>
          <label class="block text-sm font-bold text-gray-700 mb-2">سرعت ارسال:</label>
          <div class="grid grid-cols-3 gap-2">
            <button 
              v-for="option in deliveryOptions" 
              :key="option.value"
              :class="[
                'px-3 py-2 rounded-xl text-sm font-bold transition-all transform hover:scale-105',
                deliveryRating === option.value 
                  ? 'bg-gradient-to-r from-purple-500 to-pink-500 text-white shadow-lg' 
                  : 'bg-gray-100 hover:bg-purple-100 text-gray-700 hover:text-purple-700'
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
        <label class="block text-sm font-bold text-gray-700 mb-2">نظر شما:</label>
        <textarea 
          v-model="comment"
          class="w-full px-4 py-3 border-2 border-purple-200 rounded-xl focus:ring-2 focus:ring-purple-500 focus:border-purple-500 text-gray-700 font-medium" 
          rows="3" 
          placeholder="تجربه خود را با ما به اشتراک بگذارید..."
        ></textarea>
      </div>

      <!-- Submit Button -->
      <button 
        :disabled="!selectedRating || submitting"
        class="w-full bg-gradient-to-r from-purple-500 to-pink-500 hover:from-purple-600 hover:to-pink-600 disabled:from-gray-400 disabled:to-gray-500 text-white font-bold py-4 px-6 rounded-xl transition-all transform hover:scale-105 shadow-lg flex items-center justify-center space-x-2 space-x-reverse"
        @click="submitSurvey"
      >
        <svg v-if="submitting" class="animate-spin w-5 h-5" fill="none" viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
        </svg>
        <span>{{ submitting ? 'در حال ارسال...' : 'ارسال نظر ✨' }}</span>
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

const props = defineProps<Props>()
const emit = defineEmits<{
  submit: [data: any]
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
