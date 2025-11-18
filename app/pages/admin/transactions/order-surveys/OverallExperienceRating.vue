<template>
  <div class="bg-white rounded-2xl shadow-xl border border-gray-100 p-8 mb-8">
    <div class="flex items-center mb-6">
      <div class="p-3 bg-gradient-to-r from-yellow-400 to-yellow-600 rounded-xl shadow-lg">
        <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17.75l-6.172 3.247 1.179-6.873L2 9.753l6.908-1.004L12 2.25l3.092 6.499L22 9.753l-5.007 4.371 1.179 6.873z" />
        </svg>
      </div>
      <div class="mr-4">
        <h2 class="text-2xl font-bold text-gray-900">امتیازدهی کلی تجربه خرید</h2>
        <p class="text-gray-600 mt-1">نظر خود را درباره تجربه کلی خرید ثبت کنید</p>
      </div>
    </div>

    <!-- امتیاز کلی (ستاره‌ای) -->
    <div class="mb-8">
      <label class="block text-lg font-medium text-gray-800 mb-2">امتیاز کلی شما</label>
      <div class="flex items-center gap-2">
        <button
          v-for="i in 5"
          :key="i"
          type="button"
          @click="overallRating = i"
          class="text-4xl transition-colors focus:outline-none"
          :class="i <= overallRating ? 'text-yellow-400' : 'text-gray-300 hover:text-yellow-300'"
        >★</button>
        <span class="ml-4 text-lg text-gray-700 font-semibold">{{ ratingText }}</span>
      </div>
    </div>

    <!-- رضایت از فرآیند خرید -->
    <div class="mb-8">
      <label class="block text-lg font-medium text-gray-800 mb-2">رضایت از فرآیند خرید</label>
      <div class="flex gap-3">
        <button
          v-for="(option, idx) in satisfactionOptions"
          :key="idx"
          @click="satisfaction = option.value"
          :class="['px-5 py-2 rounded-lg font-medium border transition', satisfaction === option.value ? 'bg-blue-600 text-white border-blue-700' : 'bg-white text-gray-700 border-gray-300 hover:bg-blue-50']"
        >
          {{ option.label }}
        </button>
      </div>
    </div>

    <!-- احتمال خرید مجدد -->
    <div class="mb-8">
      <label class="block text-lg font-medium text-gray-800 mb-2">احتمال خرید مجدد</label>
      <div class="flex gap-3">
        <button
          v-for="(option, idx) in repurchaseOptions"
          :key="idx"
          @click="repurchase = option.value"
          :class="['px-5 py-2 rounded-lg font-medium border transition', repurchase === option.value ? 'bg-green-600 text-white border-green-700' : 'bg-white text-gray-700 border-gray-300 hover:bg-green-50']"
        >
          {{ option.label }}
        </button>
      </div>
    </div>

    <!-- توصیه به دیگران -->
    <div>
      <label class="block text-lg font-medium text-gray-800 mb-2">آیا این فروشگاه را به دیگران توصیه می‌کنید؟</label>
      <div class="flex gap-3">
        <button
          v-for="(option, idx) in recommendOptions"
          :key="idx"
          @click="recommend = option.value"
          :class="['px-5 py-2 rounded-lg font-medium border transition', recommend === option.value ? 'bg-purple-600 text-white border-purple-700' : 'bg-white text-gray-700 border-gray-300 hover:bg-purple-50']"
        >
          {{ option.label }}
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

const overallRating = ref(0)
const satisfaction = ref('')
const repurchase = ref('')
const recommend = ref('')

const satisfactionOptions = [
  { value: 'excellent', label: 'عالی' },
  { value: 'good', label: 'خوب' },
  { value: 'average', label: 'متوسط' },
  { value: 'poor', label: 'ضعیف' }
]

const repurchaseOptions = [
  { value: 'very_likely', label: 'قطعاً' },
  { value: 'likely', label: 'احتمالاً' },
  { value: 'maybe', label: 'شاید' },
  { value: 'unlikely', label: 'بعید است' }
]

const recommendOptions = [
  { value: 'yes', label: 'بله' },
  { value: 'maybe', label: 'شاید' },
  { value: 'no', label: 'خیر' }
]

const ratingText = computed(() => {
  switch (overallRating.value) {
    case 1: return 'خیلی بد'
    case 2: return 'بد'
    case 3: return 'متوسط'
    case 4: return 'خوب'
    case 5: return 'عالی'
    default: return 'امتیاز دهید'
  }
})
</script>

<style scoped>
button:focus {
  outline: 2px solid #6366f1;
  outline-offset: 2px;
}
</style> 
