<template>
  <div class="min-h-screen">
    <div class="max-w-6xl mx-auto px-4 py-4">
      <!-- Loading State -->
      <div v-if="loading" class="bg-white rounded-xl shadow-sm border border-gray-200 p-12 text-center">
        <div class="inline-flex flex-col items-center space-y-4">
          <div class="w-16 h-16 bg-blue-100 rounded-full flex items-center justify-center">
            <span class="i-heroicons-arrow-path animate-spin text-2xl text-blue-600"></span>
          </div>
          <div>
            <h3 class="text-lg font-medium text-gray-900 mb-1">در حال بارگذاری...</h3>
            <p class="text-gray-600">در حال دریافت اطلاعات درگاه پرداخت</p>
          </div>
        </div>
      </div>

      <!-- Error State -->
      <div v-else-if="error" class="bg-white rounded-xl shadow-sm border border-gray-200 p-12 text-center">
        <div class="inline-flex flex-col items-center space-y-4">
          <div class="w-16 h-16 bg-red-100 rounded-full flex items-center justify-center">
            <span class="i-heroicons-exclamation-triangle text-2xl text-red-600"></span>
          </div>
          <div>
            <h3 class="text-lg font-medium text-gray-900 mb-1">خطا در بارگذاری اطلاعات</h3>
            <p class="text-gray-600 mb-6">{{ error }}</p>
            <button 
              @click="loadGateway" 
              class="inline-flex items-center px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
            >
              <span class="i-heroicons-arrow-path ml-2"></span>
              تلاش مجدد
            </button>
          </div>
        </div>
      </div>

      <!-- Gateway Form -->
      <div v-else-if="gateway" class="bg-white rounded-xl shadow-sm border border-gray-200">
        <GatewayForm 
          :gateway="gateway"
          mode="edit"
          @save="handleSave"
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
import { defineAsyncComponent, onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'

// Import کامپوننت‌ها - lazy loading برای بهینه‌سازی
const GatewayForm = defineAsyncComponent(() => import('./components/GatewayForm.vue'))
const NotificationToast = defineAsyncComponent(() => import('~/components/admin/common/NotificationToast.vue'))

const router = useRouter()
const route = useRoute()

// متغیرهای reactive
const loading = ref(true)
const error = ref('')
const gateway = ref<any>(null)

// متغیرهای reactive برای notification
const showNotification = ref(false)
const notificationType = ref<'success' | 'error'>('success')
const notificationTitle = ref('')
const notificationMessage = ref('')

// Type declaration for Nuxt 4 auto-imported functions
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
declare const useHead: (head: { title?: string }) => void

// تعریف متا برای صفحه
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// عنوان صفحه - ثابت برای جلوگیری از re-render
useHead({
  title: 'ویرایش درگاه پرداخت'
})

// تابع نمایش notification
const showToast = (type: 'success' | 'error', title: string, message: string) => {
  notificationType.value = type
  notificationTitle.value = title
  notificationMessage.value = message
  showNotification.value = true
}

// دریافت route یکبار در ابتدا
const gatewayId = route.params.id as string

// تابع بارگذاری اطلاعات درگاه - بهینه‌سازی شده
const loadGateway = async () => {
  try {
    loading.value = true
    error.value = ''
    
    const apiUrl = `/api/payment-gateways/${gatewayId}`
    
    // استفاده از $fetch ساده و سریع
    const response: any = await $fetch(apiUrl, {
      timeout: 5000
    })
    
    if (response.status === 'success' && response.data) {
      gateway.value = response.data
    } else {
      error.value = 'درگاه پرداخت یافت نشد'
    }
  } catch (err: any) {
    error.value = err.message || 'خطا در بارگذاری اطلاعات درگاه'
  } finally {
    loading.value = false
  }
}

// تابع ذخیره تغییرات
const handleSave = async (updatedGateway: any) => {
  try {
    const apiUrl = `/api/payment-gateways/${gatewayId}`
    
    const response: any = await $fetch(apiUrl, {
      method: 'PUT',
      body: updatedGateway
    })
    
    if (response.status === 'success') {
      showToast('success', 'موفقیت', 'درگاه پرداخت با موفقیت بروزرسانی شد')
      setTimeout(() => {
        router.push('/admin/finance/payment-gateways')
      }, 1500)
    } else {
      showToast('error', 'خطا', 'خطا در بروزرسانی درگاه پرداخت')
    }
  } catch (err: any) {
    showToast('error', 'خطا', 'خطا در بروزرسانی درگاه پرداخت')
  }
}

// بارگذاری اطلاعات درگاه در زمان mount
onMounted(() => {
  loadGateway()
})
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