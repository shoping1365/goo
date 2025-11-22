<template>
  <div class="min-h-screen">
    <div class="max-w-6xl mx-auto px-4 py-4">
      <!-- Gateway Form -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200">
        <GatewayForm 
          mode="create"
          @save="handleSave as (data: { [key: string]: unknown }) => Promise<void>"
          @cancel="router.push('/admin/finance/payment-gateways')"
        />
      </div>
    </div>

    <!-- Notification Toast -->
    <NotificationToast
      :show="showNotification"
      :type="notificationType"
      :title="notificationTitle"
      :message="notificationMessage"
      @close="showNotification = false"
    />
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

// Import کامپوننت‌ها
import NotificationToast from '~/components/admin/common/NotificationToast.vue'
import GatewayForm from './components/GatewayForm.vue'

// Type declaration for Nuxt 4 auto-imported functions
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
declare const useHead: (head: { title?: string }) => void

const router = useRouter()

// متغیرهای reactive برای notification
const showNotification = ref(false)
const notificationType = ref<'success' | 'error'>('success')
const notificationTitle = ref('')
const notificationMessage = ref('')

// تعریف متا برای صفحه
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// عنوان صفحه
useHead({
  title: 'افزودن درگاه پرداخت جدید - پنل مدیریت'
})

// تابع نمایش notification
const showToast = (type: 'success' | 'error', title: string, message: string) => {
  notificationType.value = type
  notificationTitle.value = title
  notificationMessage.value = message
  showNotification.value = true
}

// تابع ذخیره درگاه جدید
const handleSave = async (data: { [key: string]: unknown }) => {
  try {

    const response = await $fetch('/api/payment-gateways', {
      method: 'POST',
      body: data
    }) as { success?: boolean; message?: string }

    // بررسی response از backend
    if (response.success || response.message) {
      showToast('success', 'موفقیت', response.message || 'درگاه پرداخت با موفقیت ایجاد شد')
      setTimeout(() => {
        router.push('/admin/finance/payment-gateways')
      }, 1500)
    } else {
      showToast('error', 'خطا', 'خطا در ایجاد درگاه پرداخت')
    }
  } catch (err) {
    const errorMessage = err instanceof Error ? err.message : 'خطای نامشخص'
    console.error('خطا در ایجاد درگاه:', errorMessage)
    showToast('error', 'خطا', 'خطا در ایجاد درگاه پرداخت')
  }
}
</script>

<!--
  صفحه افزودن درگاه پرداخت جدید
  شامل:
  1. Breadcrumb برای ناوبری بهتر
  2. هدر زیبا با آیکون و توضیحات
  3. فرم کامل برای افزودن درگاه
  4. مدیریت خطاها و پیام‌های موفقیت
  5. طراحی مدرن و ریسپانسیو
--> 
