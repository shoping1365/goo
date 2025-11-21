<template>
  <div class="bg-white border border-gray-200 rounded-lg overflow-hidden">
    <div class="p-6 text-center border-b border-gray-100">
      <h2 class="text-xl font-semibold text-gray-800 mb-1">نظرسنجی</h2>
      <p class="text-gray-600 text-sm">تجربه خرید شما</p>
    </div>
    
    <div class="p-6">
      <!-- Order Info -->
      <div class="bg-gray-50 rounded p-3 mb-6">
        <div class="text-sm text-gray-600">
          <span>محصول: {{ order.productName }}</span>
          <span class="mx-2">•</span>
          <span>{{ formatPrice(order.totalAmount) }}</span>
        </div>
      </div>

      <!-- Rating -->
      <div class="mb-6">
        <h3 class="text-base font-medium text-gray-800 mb-3">سطح رضایت:</h3>
        <div class="flex items-center justify-center space-x-4 space-x-reverse">
          <label class="flex items-center">
            <input v-model="selectedRating" type="radio" name="rating" value="1" class="mr-2">
            <span class="text-sm text-gray-600">ضعیف</span>
          </label>
          <label class="flex items-center">
            <input v-model="selectedRating" type="radio" name="rating" value="2" class="mr-2">
            <span class="text-sm text-gray-600">متوسط</span>
          </label>
          <label class="flex items-center">
            <input v-model="selectedRating" type="radio" name="rating" value="3" class="mr-2">
            <span class="text-sm text-gray-600">خوب</span>
          </label>
          <label class="flex items-center">
            <input v-model="selectedRating" type="radio" name="rating" value="4" class="mr-2">
            <span class="text-sm text-gray-600">عالی</span>
          </label>
        </div>
      </div>

      <!-- Questions -->
      <div class="space-y-4 mb-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">کیفیت:</label>
          <div class="flex space-x-2 space-x-reverse">
            <button 
              v-for="option in qualityOptions" 
              :key="option.value"
              :class="[
                'px-3 py-1 rounded text-sm transition-colors',
                qualityRating === option.value 
                  ? 'bg-gray-800 text-white' 
                  : 'bg-gray-100 hover:bg-gray-200 text-gray-700'
              ]"
              @click="qualityRating = option.value"
            >
              {{ option.label }}
            </button>
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">ارسال:</label>
          <div class="flex space-x-2 space-x-reverse">
            <button 
              v-for="option in deliveryOptions" 
              :key="option.value"
              :class="[
                'px-3 py-1 rounded text-sm transition-colors',
                deliveryRating === option.value 
                  ? 'bg-gray-800 text-white' 
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
        <label class="block text-sm font-medium text-gray-700 mb-2">نظر:</label>
        <textarea 
          v-model="comment"
          class="w-full px-3 py-2 border border-gray-300 rounded focus:ring-1 focus:ring-gray-400 focus:border-gray-400" 
          rows="3" 
          placeholder="نظر خود را بنویسید..."
        ></textarea>
      </div>

      <!-- Submit Button -->
      <button 
        :disabled="!selectedRating || submitting"
        class="w-full bg-gray-800 hover:bg-gray-900 disabled:bg-gray-400 text-white font-medium py-2 px-4 rounded transition-colors flex items-center justify-center"
        @click="submitSurvey"
      >
        <svg v-if="submitting" class="animate-spin w-4 h-4 mr-2" fill="none" viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
        </svg>
        <span>{{ submitting ? 'در حال ارسال...' : 'ارسال' }}</span>
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

defineProps<Props>()
const emit = defineEmits<{
  submit: [data: Record<string, unknown>]
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

const submitSurvey = async () => {
  if (!selectedRating.value) return
  
  submitting.value = true
  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    emit('submit', {
      rating: Number(selectedRating.value),
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
