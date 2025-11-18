<template>
  <transition name="fade">
    <div v-if="visible" class="w-full bg-red-50 border border-red-200 text-red-800 p-3 flex items-start justify-between gap-3">
      <div class="flex-1">
        <div class="font-bold mb-1">{{ bannerTitle || title || 'خطا' }}</div>
        <div class="text-sm whitespace-pre-line">{{ bannerMessage || message }}</div>
        <div v-if="meta" class="text-xs text-red-600 mt-1">{{ meta }}</div>
      </div>
      <div class="flex items-center gap-2">
        <button class="text-xs underline" @click="copyDetails">کپی جزئیات</button>
        <button class="text-xs underline" @click="dismiss">بستن</button>
      </div>
    </div>
  </transition>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'

const props = withDefaults(defineProps<{ title?: string, message?: string, isOffline?: boolean }>(), {
  title: '',
  message: '',
  isOffline: false
})
const emit = defineEmits<{(e:'dismiss'):void}>()

const visible = ref(false)
const meta = ref('')
const bannerTitle = ref<string>('')
const bannerMessage = ref<string>('')

function dismiss(){
  visible.value = false
  emit('dismiss')
}

function copyDetails(){
  const payload = (window as any).__lastError || { title: bannerTitle.value || props.title, message: bannerMessage.value || props.message, meta: meta.value }
  const text = JSON.stringify(payload, null, 2)
  navigator.clipboard?.writeText(text).catch(()=>{})
}

function onLastError(ev: any){
  try {
    const d = ev?.detail || {}
    // صف موقت در sessionStorage برای گزارش‌های بعدی
    const qKey = '__error_queue__'
    const qRaw = sessionStorage.getItem(qKey)
    const q = qRaw ? JSON.parse(qRaw) : []
    q.push({ ts: Date.now(), ...d })
    sessionStorage.setItem(qKey, JSON.stringify(q.slice(-100))) // فقط 100 مورد آخر

    visible.value = true
    bannerTitle.value = typeof d.title === 'string' && d.title ? d.title : 'خطا'
    bannerMessage.value = typeof d.message === 'string' ? d.message : ''
    meta.value = d.meta || ''
  } catch {}
}

onMounted(()=>{
  window.addEventListener('app:last-error', onLastError as any)
  window.addEventListener('offline', ()=>{
    visible.value = true
    bannerTitle.value = 'عدم دسترسی به اینترنت'
    bannerMessage.value = 'اتصال اینترنت برقرار نیست. لطفاً اتصال خود را بررسی کنید.'
    meta.value = ''
  })
  window.addEventListener('online', ()=>{
    // خاموش‌کردن بنر در بازگشت آنلاین تنها اگر پیامی در صف نباشد
    // اینجا کاری نمی‌کنیم تا خطاهای اخیر دیده شوند
  })
})
onUnmounted(()=>{
  window.removeEventListener('app:last-error', onLastError as any)
})
</script>

<style scoped>
.fade-enter-active, .fade-leave-active { transition: opacity .2s; }
.fade-enter-from, .fade-leave-to { opacity: 0; }
</style>