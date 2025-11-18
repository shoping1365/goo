import { defineNuxtPlugin } from 'nuxt/app'
import 'vue3-toastify/dist/index.css'
import Vue3Toastify, { toast, type ToastContainerOptions } from 'vue3-toastify'

// پلاگین تنظیمات Toast
export default defineNuxtPlugin((nuxtApp) => {
  const options: ToastContainerOptions = {
    autoClose: 6000,
    position: 'top-center',
    rtl: true,
    pauseOnHover: true,
    closeOnClick: true,
    hideProgressBar: false,
    theme: 'colored'
  }
  nuxtApp.vueApp.use(Vue3Toastify, options)
    // expose toast function
    ; (nuxtApp as any).$toast = toast
    ; (nuxtApp.vueApp.config.globalProperties as any).$toast = toast
  nuxtApp.provide('toast', toast)
  if (typeof window !== 'undefined') (window as any).$toast = toast

  // Ensure Persian font and centered text inside toasts
  if (typeof window !== 'undefined') {
    const style = document.createElement('style')
    style.textContent = `
      .Toastify__toast { font-family: 'IRANSansWeb', 'IRANSansWebFaNum', Tahoma, Arial, sans-serif; }
      .Toastify__toast-body { text-align: center; width: 100%; }
    `
    document.head.appendChild(style)
  }
})


