<template>
  <div class="bg-white/80 backdrop-blur-md rounded-2xl shadow-lg border border-gray-200/50 p-6">
    <div class="flex items-center justify-between mb-4">
      <h3 class="text-lg font-semibold text-gray-900 flex items-center">
        <svg class="w-5 h-5 ml-2 text-blue-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"></path>
        </svg>
        چت‌بات هوشمند
      </h3>
      <div class="flex items-center space-x-2 space-x-reverse">
        <span 
          :class="[
            isActive ? 'bg-green-400' : 'bg-gray-400',
            'w-2 h-2 rounded-full animate-pulse'
          ]"
        ></span>
        <span class="text-sm text-gray-600">{{ isActive ? 'فعال' : 'غیرفعال' }}</span>
      </div>
    </div>

    <!-- Bot Status -->
    <div class="bg-gradient-to-r from-blue-50 to-purple-50 rounded-lg p-6 mb-4">
      <div class="flex items-center justify-between">
        <div>
          <div class="text-sm font-medium text-gray-900">وضعیت چت‌بات</div>
          <div class="text-xs text-gray-600 mt-1">
            {{ botStats.totalResponses }} پاسخ خودکار در ۲۴ ساعت گذشته
          </div>
        </div>
        <button 
          @click="toggleBot"
          :class="[
            isActive ? 'bg-red-500 hover:bg-red-600' : 'bg-green-500 hover:bg-green-600',
            'px-3 py-1 text-xs font-medium text-white rounded-lg transition-colors'
          ]"
        >
          {{ isActive ? 'غیرفعال کردن' : 'فعال کردن' }}
        </button>
      </div>
    </div>

    <!-- Quick Responses -->
    <div class="space-y-3">
      <h4 class="text-sm font-medium text-gray-900">پاسخ‌های سریع</h4>
      <div class="grid grid-cols-2 gap-2">
        <button 
          v-for="response in quickResponses"
          :key="response.id"
          @click="sendBotResponse(response)"
          class="p-2 text-xs bg-gray-100 hover:bg-gray-200 text-gray-700 rounded-lg transition-colors text-right"
        >
          {{ response.text }}
        </button>
      </div>
    </div>

    <!-- Bot Settings -->
    <div class="mt-4 pt-4 border-t border-gray-200">
      <h4 class="text-sm font-medium text-gray-900 mb-3">تنظیمات</h4>
      <div class="space-y-2">
        <div class="flex items-center justify-between">
          <span class="text-xs text-gray-600">پاسخ خودکار</span>
          <button 
            @click="settings.autoResponse = !settings.autoResponse"
            :class="[
              settings.autoResponse ? 'bg-blue-600' : 'bg-gray-200',
              'relative inline-flex h-4 w-8 items-center rounded-full transition-colors'
            ]"
          >
            <span 
              :class="[
                settings.autoResponse ? 'translate-x-4' : 'translate-x-1',
                'inline-block h-2 w-2 transform rounded-full bg-white transition-transform'
              ]"
            ></span>
          </button>
        </div>
        <div class="flex items-center justify-between">
          <span class="text-xs text-gray-600">تشخیص احساسات</span>
          <button 
            @click="settings.sentimentAnalysis = !settings.sentimentAnalysis"
            :class="[
              settings.sentimentAnalysis ? 'bg-blue-600' : 'bg-gray-200',
              'relative inline-flex h-4 w-8 items-center rounded-full transition-colors'
            ]"
          >
            <span 
              :class="[
                settings.sentimentAnalysis ? 'translate-x-4' : 'translate-x-1',
                'inline-block h-2 w-2 transform rounded-full bg-white transition-transform'
              ]"
            ></span>
          </button>
        </div>
        <div class="flex items-center justify-between">
          <span class="text-xs text-gray-600">ترجمه خودکار</span>
          <button 
            @click="settings.autoTranslate = !settings.autoTranslate"
            :class="[
              settings.autoTranslate ? 'bg-blue-600' : 'bg-gray-200',
              'relative inline-flex h-4 w-8 items-center rounded-full transition-colors'
            ]"
          >
            <span 
              :class="[
                settings.autoTranslate ? 'translate-x-4' : 'translate-x-1',
                'inline-block h-2 w-2 transform rounded-full bg-white transition-transform'
              ]"
            ></span>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// Reactive data
const isActive = ref(true)

// Bot stats
const botStats = ref({
  totalResponses: 247,
  accuracy: 94.5,
  avgResponseTime: 1.2
})

// Settings
const settings = ref({
  autoResponse: true,
  sentimentAnalysis: true,
  autoTranslate: false
})

// Quick responses
const quickResponses = ref([
  {
    id: 1,
    text: 'سلام! چطور می‌تونم کمکتون کنم؟',
    category: 'greeting'
  },
  {
    id: 2,
    text: 'متشکرم از تماس شما',
    category: 'thanks'
  },
  {
    id: 3,
    text: 'لطفاً کمی صبر کنید',
    category: 'wait'
  },
  {
    id: 4,
    text: 'سوال خوبی پرسیدید',
    category: 'acknowledgment'
  },
  {
    id: 5,
    text: 'آیا سوال دیگری دارید؟',
    category: 'followup'
  },
  {
    id: 6,
    text: 'خوشحالم که تونستم کمکتون کنم',
    category: 'closing'
  }
])

// Methods
const toggleBot = () => {
  isActive.value = !isActive.value
  console.log('Bot toggled:', isActive.value)
}

const sendBotResponse = (response: any) => {
  if (!isActive.value) return
  
  console.log('Bot response sent:', response.text)
}

// Auto-response logic
const generateAutoResponse = (message: string): string => {
  const keywords = {
    'سلام': 'سلام! خوشحالم که با شما صحبت می‌کنم. چطور می‌تونم کمکتون کنم؟',
    'قیمت': 'قیمت محصولات ما بر اساس مدل و ویژگی‌ها متفاوت است. کدوم محصول مد نظرتون هست؟',
    'تخفیف': 'در حال حاضر تخفیف‌های ویژه‌ای روی محصولات ما اعمال شده. می‌تونم راهنماییتون کنم.',
    'موجودی': 'برای اطلاع از موجودی محصول مورد نظر، لطفاً کد محصول رو بگید.',
    'ارسال': 'زمان ارسال بسته به آدرس شما بین ۱ تا ۳ روز کاری است.',
    'بازگشت': 'طبق قوانین ما، تا ۷ روز امکان بازگشت محصول وجود دارد.',
    'گارانتی': 'تمام محصولات ما دارای گارانتی معتبر هستند.',
    'مشکل': 'متأسفم که با مشکل مواجه شدید. لطفاً جزئیات بیشتری ارائه دهید تا بتونم کمکتون کنم.'
  }

  for (const [key, response] of Object.entries(keywords)) {
    if (message.includes(key)) {
      return response
    }
  }

  return 'متوجه سوال شما شدم. لطفاً کمی صبر کنید تا یکی از همکارانم به شما پاسخ دهد.'
}

// Expose methods for parent component
defineExpose({
  generateAutoResponse,
  isActive
})
</script>

<style scoped>
/* Custom styles for chatbot */
</style> 
