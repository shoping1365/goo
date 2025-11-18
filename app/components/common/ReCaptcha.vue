<template>
  <div ref="recaptchaContainer"></div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'

const props = defineProps<{
  sitekey: string
  theme?: 'light' | 'dark'
  size?: 'normal' | 'compact'
  tabindex?: number
}>()

const emit = defineEmits<{
  (e: 'verify', token: string): void
  (e: 'expire'): void
  (e: 'fail'): void
}>()

const recaptchaContainer = ref<HTMLElement | null>(null)
let recaptchaId: number | null = null

const renderRecaptcha = () => {
  if (!recaptchaContainer.value) return

  // @ts-ignore
  recaptchaId = window.grecaptcha.render(recaptchaContainer.value, {
    sitekey: props.sitekey,
    theme: props.theme || 'light',
    size: props.size || 'normal',
    tabindex: props.tabindex || 0,
    callback: (token: string) => emit('verify', token),
    'expired-callback': () => emit('expire'),
    'error-callback': () => emit('fail')
  })
}

const resetRecaptcha = () => {
  if (recaptchaId !== null) {
    // @ts-ignore
    window.grecaptcha.reset(recaptchaId)
  }
}

onMounted(() => {
  // Load reCAPTCHA script if not already loaded
  if (!document.querySelector('script[src*="recaptcha/api.js"]')) {
    const script = document.createElement('script')
    script.src = `https://www.google.com/recaptcha/api.js?render=explicit`
    script.async = true
    script.defer = true
    script.onload = renderRecaptcha
    document.head.appendChild(script)
  } else {
    renderRecaptcha()
  }
})

// Watch for prop changes
watch(() => props.theme, resetRecaptcha)
watch(() => props.size, resetRecaptcha)
</script> 
