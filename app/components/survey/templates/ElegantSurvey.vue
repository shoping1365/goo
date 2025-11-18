<template>
  <div class="bg-white rounded-xl shadow-lg overflow-hidden border border-gray-100">
    <div class="bg-gradient-to-r from-amber-400 to-yellow-500 p-6 text-white text-center">
      <div class="w-16 h-16 bg-white bg-opacity-20 rounded-full mx-auto mb-4 flex items-center justify-center">
        <svg class="w-8 h-8" fill="currentColor" viewBox="0 0 20 20">
          <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"></path>
        </svg>
      </div>
      <h2 class="text-2xl font-bold mb-2">نظرسنجی کیفیت</h2>
      <p class="text-yellow-100">ارزش‌گذاری شما برای بهبود خدمات</p>
    </div>
    
    <div class="p-6">
      <!-- Order Info -->
      <div class="bg-amber-50 border border-amber-200 rounded-lg p-6 mb-6">
        <h3 class="font-semibold text-amber-800 mb-2">اطلاعات سفارش</h3>
        <div class="grid grid-cols-2 gap-6 text-sm">
          <div>
            <span class="text-amber-600">محصول:</span>
            <span class="font-medium text-amber-800 mr-2">{{ order.productName }}</span>
          </div>
          <div>
            <span class="text-amber-600">مبلغ:</span>
            <span class="font-medium text-amber-800 mr-2">{{ formatPrice(order.totalAmount) }}</span>
          </div>
        </div>
      </div>

      <!-- Rating -->
      <div class="mb-6">
        <h3 class="text-lg font-semibold text-gray-800 mb-4 text-center">امتیاز کلی محصول</h3>
        <div class="flex justify-center space-x-1 space-x-reverse mb-3">
          <button 
            v-for="star in 5" 
            :key="star"
            @click="selectedRating = star"
            class="transition-transform hover:scale-110"
          >
            <svg 
              :class="[
                'w-8 h-8',
                star <= selectedRating ? 'text-yellow-400' : 'text-gray-300'
              ]" 
              fill="currentColor" 
              viewBox="0 0 20 20"
            >
              <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"></path>
            </svg>
          </button>
        </div>
        <div class="text-center text-sm text-gray-600">
          {{ getRatingText(selectedRating) }}
        </div>
      </div>

      <!-- Questions -->
      <div class="space-y-4 mb-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">کیفیت محصول:</label>
          <div class="grid grid-cols-2 gap-2">
            <button 
              v-for="option in qualityOptions" 
              :key="option.value"
              @click="qualityRating = option.value"
              :class="[
                'px-4 py-2 rounded-lg text-sm transition-all',
                qualityRating === option.value 
                  ? 'bg-amber-500 text-white shadow-md' 
                  : 'bg-gray-100 hover:bg-amber-50 text-gray-700 hover:text-amber-700'
              ]"
            >
              {{ option.label }}
            </button>
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">سرعت ارسال:</label>
          <div class="grid grid-cols-3 gap-2">
            <button 
              v-for="option in deliveryOptions" 
              :key="option.value"
              @click="deliveryRating = option.value"
              :class="[
                'px-3 py-2 rounded-lg text-sm transition-all',
                deliveryRating === option.value 
                  ? 'bg-amber-500 text-white shadow-md' 
                  : 'bg-gray-100 hover:bg-amber-50 text-gray-700 hover:text-amber-700'
              ]"
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
          class="w-full px-3 py-2 border border-gray-200 rounded-lg focus:ring-2 focus:ring-amber-500 focus:border-amber-500" 
          rows="3" 
          placeholder="تجربه خود را با ما به اشتراک بگذارید..."
        ></textarea>
      </div>

      <!-- Submit Button -->
      <button 
        @click="submitSurvey"
        :disabled="!selectedRating || submitting"
        class="w-full bg-gradient-to-r from-amber-400 to-yellow-500 hover:from-amber-500 hover:to-yellow-600 disabled:from-gray-400 disabled:to-gray-500 text-white font-medium py-3 px-4 rounded-lg transition-all transform hover:scale-105 flex items-center justify-center space-x-2 space-x-reverse"
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
