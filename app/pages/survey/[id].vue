<template>
  <div class="min-h-screen bg-white py-8">
    <div class="max-w-2xl mx-auto px-4">
      <!-- Loading State -->
      <div v-if="loading" class="text-center py-12">
        <div class="animate-spin w-12 h-12 border-4 border-blue-500 border-t-transparent rounded-full mx-auto mb-4"></div>
        <p class="text-gray-600">در حال بارگذاری نظرسنجی...</p>
      </div>

      <!-- Survey Form -->
      <div v-else-if="surveyData" class="space-y-6">
        <!-- Header -->
        <div class="text-center">
          <h1 class="text-2xl font-bold text-gray-900 mb-2">نظرسنجی خرید</h1>
          <p class="text-gray-600">تجربه شما برای بهبود خدمات ما مهم است</p>
        </div>

        <!-- Survey Form -->
        <SurveyForm 
          :order="surveyData.order" 
          :template="surveyData.template"
          @submit="handleSurveySubmit"
        />

        <!-- Success Message -->
        <div v-if="submitted" class="bg-green-50 border border-green-200 rounded-lg p-6 text-center">
          <div class="w-16 h-16 bg-green-100 rounded-full mx-auto mb-4 flex items-center justify-center">
            <svg class="w-8 h-8 text-green-600" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-green-800 mb-2">نظرسنجی با موفقیت ارسال شد!</h3>
          <p class="text-green-600">از مشارکت شما در بهبود خدمات متشکریم</p>
        </div>
      </div>

      <!-- Error State -->
      <div v-else class="text-center py-12">
        <div class="w-16 h-16 bg-red-100 rounded-full mx-auto mb-4 flex items-center justify-center">
          <svg class="w-8 h-8 text-red-600" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clip-rule="evenodd"></path>
          </svg>
        </div>
        <h3 class="text-lg font-semibold text-red-800 mb-2">نظرسنجی یافت نشد</h3>
        <p class="text-red-600">لینک نظرسنجی نامعتبر است یا منقضی شده</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'
import SurveyForm from '~/components/survey/SurveyForm.vue'

interface Order {
  id: string
  customerName: string
  productName: string
  orderDate: string
  totalAmount: number
}

interface SurveyData {
  order: Order
  template: string
}

const route = useRoute()
const loading = ref(true)
const submitted = ref(false)
const surveyData = ref<SurveyData | null>(null)

// Mock data - در واقعیت از API دریافت می‌شود
const mockSurveyData = {
  'survey-123': {
    order: {
      id: 'order-123',
      customerName: 'علی احمدی',
      productName: 'لپ‌تاپ اپل مک‌بوک پرو',
      orderDate: '2024-01-15',
      totalAmount: 85000000
    },
    template: 'modern'
  },
  'survey-456': {
    order: {
      id: 'order-456',
      customerName: 'فاطمه محمدی',
      productName: 'گوشی سامسونگ Galaxy S24',
      orderDate: '2024-01-14',
      totalAmount: 45000000
    },
    template: 'elegant'
  },
  'survey-789': {
    order: {
      id: 'order-789',
      customerName: 'محمد رضایی',
      productName: 'ساعت هوشمند اپل واچ',
      orderDate: '2024-01-13',
      totalAmount: 25000000
    },
    template: 'colorful'
  }
}

const fetchSurveyData = async () => {
  loading.value = true
  try {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    const surveyId = route.params.id as string
    const data = mockSurveyData[surveyId as keyof typeof mockSurveyData]
    
    if (data) {
      surveyData.value = data
    }
  } catch {
    // Error fetching survey data
  } finally {
    loading.value = false
  }
}

const handleSurveySubmit = async (_data: Record<string, unknown>) => {
  try {
    // Simulate API call to save survey response
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    // Survey submitted
    submitted.value = true
    
    // Optional: Redirect after a delay
    setTimeout(() => {
      // window.location.href = '/thank-you'
    }, 3000)
  } catch {
    // Error submitting survey
  }
}

onMounted(() => {
  fetchSurveyData()
})
</script>

<style scoped>
/* Custom styles if needed */
</style> 