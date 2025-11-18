<template>
  <div>
    <!-- Container برای Turnstile widget -->
    <div 
      ref="turnstileContainer" 
      class="turnstile-container"
      :class="{ 'loading': isLoading }"
    ></div>
    
    <!-- نمایش وضعیت بارگذاری -->
    <div v-if="isLoading" class="text-sm text-gray-500 mt-2">
      در حال بارگذاری Turnstile...
    </div>
    
    <!-- نمایش خطا -->
    <div v-if="error" class="text-sm text-red-500 mt-2">
      {{ error }}
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  siteKey: {
    type: String,
    required: true
  },
  theme: {
    type: String,
    default: 'light',
    validator: (value) => ['light', 'dark'].includes(value)
  },
  size: {
    type: String,
    default: 'normal',
    validator: (value) => ['normal', 'compact'].includes(value)
  },
  language: {
    type: String,
    default: 'fa'
  }
})

const emit = defineEmits(['token', 'expired', 'error'])

const turnstileContainer = ref(null)
const isLoading = ref(true)
const error = ref('')
const widgetId = ref(null)

// تنظیمات Turnstile
const turnstileOptions = {
  sitekey: props.siteKey,
  theme: props.theme,
  size: props.size,
  language: props.language,
  callback: (token) => {
    emit('token', token)
  },
  'expired-callback': () => {
    emit('expired')
  },
  'error-callback': () => {
    error.value = 'خطا در بارگذاری Turnstile'
    emit('error', error.value)
  }
}

// بارگذاری Turnstile script
const loadTurnstileScript = () => {
  return new Promise((resolve, reject) => {
    // بررسی اینکه آیا script قبلاً بارگذاری شده
    if (window.turnstile) {
      resolve()
      return
    }

    const script = document.createElement('script')
    script.src = 'https://challenges.cloudflare.com/turnstile/v0/api.js'
    script.async = true
    script.defer = true
    
    script.onload = () => resolve()
    script.onerror = () => reject(new Error('خطا در بارگذاری Turnstile script'))
    
    document.head.appendChild(script)
  })
}

// راه‌اندازی Turnstile widget
const initTurnstile = async () => {
  try {
    isLoading.value = true
    error.value = ''
    
    await loadTurnstileScript()
    
    // صبر کردن تا DOM آماده شود
    await nextTick()
    
    if (turnstileContainer.value && window.turnstile) {
      widgetId.value = window.turnstile.render(turnstileContainer.value, turnstileOptions)
    }
    
    isLoading.value = false
  } catch (err) {
    error.value = err.message
    isLoading.value = false
    emit('error', err.message)
  }
}

// بازنشانی widget
const reset = () => {
  if (widgetId.value && window.turnstile) {
    window.turnstile.reset(widgetId.value)
  }
}

// دریافت توکن فعلی
const getToken = () => {
  if (widgetId.value && window.turnstile) {
    return window.turnstile.getResponse(widgetId.value)
  }
  return null
}

// expose methods
defineExpose({
  reset,
  getToken
})

// راه‌اندازی اولیه
onMounted(() => {
  initTurnstile()
})

// تمیزکاری
onUnmounted(() => {
  if (widgetId.value && window.turnstile) {
    window.turnstile.remove(widgetId.value)
  }
})
</script>

<style scoped>
.turnstile-container {
  min-height: 65px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.turnstile-container.loading {
  background-color: #f3f4f6;
  border-radius: 8px;
}
</style> 