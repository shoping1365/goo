<template>
  <div class="footer-newsletter-widget" :style="widgetStyle">
    <div class="text-center">
      <!-- عنوان بخش -->
      <h4 v-if="title" class="text-white font-semibold mb-4 text-lg">{{ title }}</h4>
      
      <!-- توضیحات -->
      <p v-if="description" class="text-white text-opacity-80 mb-6 text-sm leading-relaxed">
        {{ description }}
      </p>
      
      <!-- فرم خبرنامه -->
      <form class="space-y-4" @submit.prevent="handleSubmit">
        <!-- ایمیل -->
        <div class="relative">
          <input
            v-model="email"
            type="email"
            :placeholder="emailPlaceholder"
            required
            class="w-full px-4 py-3 bg-white bg-opacity-20 border border-white border-opacity-30 rounded-lg text-white placeholder-white placeholder-opacity-60 focus:outline-none focus:ring-2 focus:ring-white focus:ring-opacity-50 focus:border-white transition-all duration-200"
            :class="{ 'border-red-400 bg-red-500 bg-opacity-20': emailError }"
          />
          
          <!-- آیکون ایمیل -->
          <div class="absolute left-3 top-1/2 transform -translate-y-1/2 text-white text-opacity-60">
            <span class="text-lg">📧</span>
          </div>
          
          <!-- پیام خطا -->
          <div v-if="emailError" class="absolute -bottom-6 left-0 text-red-300 text-xs">
            {{ emailError }}
          </div>
        </div>
        
        <!-- دکمه عضویت -->
        <button
          type="submit"
          :disabled="isSubmitting"
          class="w-full px-6 py-3 bg-gradient-to-r from-blue-500 to-purple-600 hover:from-blue-600 hover:to-purple-700 text-white font-medium rounded-lg transition-all duration-200 transform hover:scale-105 disabled:opacity-50 disabled:cursor-not-allowed disabled:transform-none"
        >
          <span v-if="!isSubmitting" class="flex items-center justify-center space-x-2 space-x-reverse">
            <span>📬</span>
            <span>{{ subscribeButtonText }}</span>
          </span>
          <span v-else class="flex items-center justify-center space-x-2 space-x-reverse">
            <div class="w-4 h-4 border-2 border-white border-t-transparent rounded-full animate-spin"></div>
            <span>در حال عضویت...</span>
          </span>
        </button>
        
        <!-- پیام موفقیت -->
        <div v-if="successMessage" class="bg-green-500 bg-opacity-20 border border-green-400 rounded-lg p-3 text-green-300 text-sm">
          <div class="flex items-center justify-center space-x-2 space-x-reverse">
            <span>✅</span>
            <span>{{ successMessage }}</span>
          </div>
        </div>
        
        <!-- پیام خطا -->
        <div v-if="errorMessage" class="bg-red-500 bg-opacity-20 border border-red-400 rounded-lg p-3 text-red-300 text-sm">
          <div class="flex items-center justify-center space-x-2 space-x-reverse">
            <span>❌</span>
            <span>{{ errorMessage }}</span>
          </div>
        </div>
      </form>
      
      <!-- اطلاعات اضافی -->
      <div v-if="showAdditionalInfo" class="mt-6 text-white text-opacity-60 text-xs space-y-2">
        <div v-if="privacyPolicy" class="flex items-center justify-center space-x-2 space-x-reverse">
          <span>🔒</span>
          <a :href="privacyPolicy" class="hover:text-white transition-colors underline">
            حریم خصوصی
          </a>
        </div>
        
        <div v-if="termsOfService" class="flex items-center justify-center space-x-2 space-x-reverse">
          <span>📋</span>
          <a :href="termsOfService" class="hover:text-white transition-colors underline">
            شرایط استفاده
          </a>
        </div>
        
        <div v-if="unsubscribeInfo" class="text-center">
          <span>💡</span>
          <span>{{ unsubscribeInfo }}</span>
        </div>
      </div>
      
      <!-- آمار -->
      <div v-if="showStats && (subscriberCount || monthlyNewsletters)" class="mt-6 pt-4 border-t border-white border-opacity-20">
        <div class="grid grid-cols-2 gap-6 text-white text-opacity-80">
          <div v-if="subscriberCount" class="text-center">
            <div class="text-lg font-semibold">{{ subscriberCount.toLocaleString() }}</div>
            <div class="text-xs">عضو خبرنامه</div>
          </div>
          <div v-if="monthlyNewsletters" class="text-center">
            <div class="text-lg font-semibold">{{ monthlyNewsletters }}</div>
            <div class="text-xs">خبرنامه ماهانه</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, watch } from 'vue'

interface Props {
  title?: string
  description?: string
  emailPlaceholder?: string
  subscribeButtonText?: string
  privacyPolicy?: string
  termsOfService?: string
  unsubscribeInfo?: string
  showAdditionalInfo?: boolean
  showStats?: boolean
  subscriberCount?: number
  monthlyNewsletters?: number
  paddingRight?: number
  paddingLeft?: number
  align?: 'left' | 'center' | 'right'
}

const props = withDefaults(defineProps<Props>(), {
  title: 'عضویت در خبرنامه',
  description: 'برای دریافت آخرین اخبار و تخفیف‌ها، در خبرنامه ما عضو شوید',
  emailPlaceholder: 'ایمیل خود را وارد کنید',
  subscribeButtonText: 'عضویت در خبرنامه',
  privacyPolicy: '',
  termsOfService: '',
  unsubscribeInfo: 'می‌توانید هر زمان که خواستید از خبرنامه خارج شوید',
  showAdditionalInfo: true,
  showStats: false,
  subscriberCount: 0,
  monthlyNewsletters: 0,
  paddingRight: 0,
  paddingLeft: 0,
  align: 'center'
})

// متغیرهای محلی
const email = ref('')
const emailError = ref('')
const isSubmitting = ref(false)
const successMessage = ref('')
const errorMessage = ref('')

// اعتبارسنجی ایمیل
const validateEmail = (email: string): boolean => {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  return emailRegex.test(email)
}

// مدیریت ارسال فرم
const handleSubmit = async () => {
  // پاک کردن پیام‌های قبلی
  emailError.value = ''
  successMessage.value = ''
  errorMessage.value = ''
  
  // اعتبارسنجی
  if (!email.value.trim()) {
    emailError.value = 'لطفاً ایمیل خود را وارد کنید'
    return
  }
  
  if (!validateEmail(email.value.trim())) {
    emailError.value = 'لطفاً یک ایمیل معتبر وارد کنید'
    return
  }
  
  // شروع فرآیند عضویت
  isSubmitting.value = true
  
  try {
    // شبیه‌سازی API call
    await new Promise(resolve => setTimeout(resolve, 1500))
    
    // موفقیت
    successMessage.value = 'با موفقیت در خبرنامه عضو شدید!'
    email.value = ''
    
    // پاک کردن پیام موفقیت بعد از 5 ثانیه
    setTimeout(() => {
      successMessage.value = ''
    }, 5000)
    
  } catch (_error) {
    // خطا
    errorMessage.value = 'خطا در عضویت. لطفاً دوباره تلاش کنید.'
    
    // پاک کردن پیام خطا بعد از 5 ثانیه
    setTimeout(() => {
      errorMessage.value = ''
    }, 5000)
  } finally {
    isSubmitting.value = false
  }
}

// پاک کردن خطاها هنگام تغییر ایمیل
watch(email, () => {
  if (emailError.value) {
    emailError.value = ''
  }
})

// استایل کامپوننت بر اساس چینش
const widgetStyle = computed(() => ({
  paddingRight: `${props.paddingRight}px`,
  paddingLeft: `${props.paddingLeft}px`,
  display: 'flex',
  alignItems: 'center',
  justifyContent: getJustifyContent(props.align),
  width: '100%',
  height: '100%'
}))

// تابع تعیین justify-content بر اساس چینش
function getJustifyContent(align: string): string {
  switch (align) {
    case 'left':
      return 'flex-start'  // در RTL: چپ = flex-start
    case 'center':
      return 'center'
    case 'right':
      return 'flex-end'  // در RTL: راست = flex-end
    default:
      return 'center'
  }
}
</script>

<style scoped>
.footer-newsletter-widget {
  transition: all 0.2s ease;
}

.footer-newsletter-widget input:focus {
  background-color: rgba(255, 255, 255, 0.3);
}

.footer-newsletter-widget button:hover:not(:disabled) {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
}

.footer-newsletter-widget a:hover {
  text-decoration: none;
}

/* انیمیشن برای دکمه */
@keyframes scale {
  0% { transform: scale(1); }
  50% { transform: scale(1.05); }
  100% { transform: scale(1); }
}

.footer-newsletter-widget button:active:not(:disabled) {
  animation: scale 0.1s ease-in-out;
}
</style>
