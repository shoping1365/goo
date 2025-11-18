<template>
  <div class="cache-clear-section">
    <button
      @click="clearCache"
      :disabled="loading"
      class="cache-clear-btn"
      :class="{ 'loading': loading, 'success': success }"
    >
      <span v-if="!loading && !success" class="icon">🗑️</span>
      <span v-if="loading" class="spinner"></span>
      <span v-if="success" class="icon">✅</span>
      <span class="text">
        {{ buttonText }}
      </span>
    </button>
    
    <div v-if="message" class="message" :class="messageType">
      {{ message }}
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'

interface CacheResponse {
  message: string
}

const loading = ref(false)
const success = ref(false)
const message = ref('')
const messageType = ref<'success' | 'error'>('success')

const buttonText = computed(() => {
  if (loading.value) return 'در حال پاک کردن کش...'
  if (success.value) return 'کش پاک شد!'
  return 'پاک کردن کش'
})

const clearCache = async () => {
  if (loading.value) return
  
  loading.value = true
  message.value = ''
  success.value = false
  
  try {
    const response = await $fetch<CacheResponse>('/api/admin/cache/clear', {
      method: 'POST'
    })
    
    success.value = true
    message.value = response?.message || 'کش با موفقیت پاک شد'
    messageType.value = 'success'
    
    // بازگشت به حالت عادی بعد از 3 ثانیه
    setTimeout(() => {
      success.value = false
      message.value = ''
    }, 3000)
    
  } catch (error: any) {
    console.error('خطا در پاک کردن کش:', error)
    message.value = error.data?.message || 'خطا در پاک کردن کش'
    messageType.value = 'error'
    
    setTimeout(() => {
      message.value = ''
    }, 5000)
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.cache-clear-section {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.cache-clear-btn {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 12px 24px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  box-shadow: 0 4px 6px rgba(102, 126, 234, 0.25);
}

.cache-clear-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 6px 12px rgba(102, 126, 234, 0.35);
}

.cache-clear-btn:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.cache-clear-btn.loading {
  background: linear-gradient(135deg, #a0aec0 0%, #718096 100%);
}

.cache-clear-btn.success {
  background: linear-gradient(135deg, #48bb78 0%, #38a169 100%);
}

.icon {
  font-size: 18px;
}

.spinner {
  width: 16px;
  height: 16px;
  border: 2px solid rgba(255, 255, 255, 0.3);
  border-top-color: white;
  border-radius: 50%;
  animation: spin 0.8s linear infinite;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

.message {
  padding: 12px 16px;
  border-radius: 6px;
  font-size: 14px;
  animation: slideIn 0.3s ease;
}

.message.success {
  background: #f0fff4;
  color: #22543d;
  border: 1px solid #9ae6b4;
}

.message.error {
  background: #fff5f5;
  color: #742a2a;
  border: 1px solid #fc8181;
}

@keyframes slideIn {
  from {
    opacity: 0;
    transform: translateY(-10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
