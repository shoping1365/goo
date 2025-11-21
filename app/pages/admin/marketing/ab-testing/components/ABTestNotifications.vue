<template>
  <div class="fixed bottom-4 left-4 z-50 space-y-2">
    <!-- هشدارهای تست -->
    <div v-for="alert in alerts" :key="alert.id" class="max-w-sm bg-white rounded-lg shadow-lg border-l-4 p-6" :class="getAlertBorderClass(alert.type)">
      <div class="flex items-start">
        <div class="flex-shrink-0">
          <svg class="w-5 h-5" :class="getAlertIconClass(alert.type)" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path v-if="alert.type === 'warning'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z" />
            <path v-else-if="alert.type === 'success'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            <path v-else-if="alert.type === 'error'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            <path v-else stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
        </div>
        <div class="mr-3 flex-1">
          <h4 class="text-sm font-medium" :class="getAlertTitleClass(alert.type)">
            {{ alert.title }}
          </h4>
          <p class="text-sm text-gray-600 mt-1">
            {{ alert.message }}
          </p>
          <div v-if="alert.action" class="mt-3">
            <button class="text-sm font-medium" :class="getAlertActionClass(alert.type)" @click="handleAlertAction(alert)">
              {{ alert.action.text }}
            </button>
          </div>
        </div>
        <div class="flex-shrink-0">
          <button class="text-gray-400 hover:text-gray-600" @click="removeAlert(alert.id)">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
      </div>
    </div>
  </div>

  <!-- اعلان‌های نتایج (Toast) -->
  <div class="fixed top-6 right-4 z-50 space-y-2">
    <div v-for="notification in notifications" :key="notification.id" class="max-w-sm bg-white rounded-lg shadow-lg border-l-4 p-6 transform transition-all duration-300" :class="getNotificationBorderClass(notification.type)">
      <div class="flex items-start">
        <div class="flex-shrink-0">
          <svg class="w-5 h-5" :class="getNotificationIconClass(notification.type)" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path v-if="notification.type === 'winner'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            <path v-else-if="notification.type === 'completion'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            <path v-else stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
        </div>
        <div class="mr-3 flex-1">
          <h4 class="text-sm font-medium" :class="getNotificationTitleClass(notification.type)">
            {{ notification.title }}
          </h4>
          <p class="text-sm text-gray-600 mt-1">
            {{ notification.message }}
          </p>
        </div>
        <div class="flex-shrink-0">
          <button class="text-gray-400 hover:text-gray-600" @click="removeNotification(notification.id)">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
// State
const alerts = ref([
  {
    id: 1,
    type: 'warning',
    title: 'تست نزدیک به پایان',
    message: 'تست "دکمه خرید صفحه محصول" تا 2 روز دیگر به پایان می‌رسد.',
    action: {
      text: 'مشاهده تست',
      handler: () => {}
    }
  },
  {
    id: 2,
    type: 'success',
    title: 'نتایج معنادار',
    message: 'تست "قیمت محصولات" نتایج آماری معنادار نشان می‌دهد.',
    action: {
      text: 'مشاهده نتایج',
      handler: () => {}
    }
  },
  {
    id: 3,
    type: 'error',
    title: 'خطا در تست',
    message: 'تست "تصویر صفحه اصلی" با خطا مواجه شده است.',
    action: {
      text: 'بررسی خطا',
      handler: () => {}
    }
  }
])

const notifications = ref([
  {
    id: 1,
    type: 'winner',
    title: 'برنده تست اعلام شد',
    message: 'نسخه B در تست "دکمه خرید" با 28.5% بهبود برنده شد.',
    timestamp: new Date()
  },
  {
    id: 2,
    type: 'completion',
    title: 'تست تکمیل شد',
    message: 'تست "قیمت محصولات" با موفقیت به پایان رسید.',
    timestamp: new Date()
  }
])

// کلاس‌های هشدارها
const getAlertBorderClass = (type: string) => {
  const classes = {
    warning: 'border-yellow-400',
    success: 'border-green-400',
    error: 'border-red-400',
    info: 'border-blue-400'
  }
  return classes[type as keyof typeof classes] || 'border-gray-400'
}

const getAlertIconClass = (type: string) => {
  const classes = {
    warning: 'text-yellow-400',
    success: 'text-green-400',
    error: 'text-red-400',
    info: 'text-blue-400'
  }
  return classes[type as keyof typeof classes] || 'text-gray-400'
}

const getAlertTitleClass = (type: string) => {
  const classes = {
    warning: 'text-yellow-800',
    success: 'text-green-800',
    error: 'text-red-800',
    info: 'text-blue-800'
  }
  return classes[type as keyof typeof classes] || 'text-gray-800'
}

const getAlertActionClass = (type: string) => {
  const classes = {
    warning: 'text-yellow-600 hover:text-yellow-500',
    success: 'text-green-600 hover:text-green-500',
    error: 'text-red-600 hover:text-red-500',
    info: 'text-blue-600 hover:text-blue-500'
  }
  return classes[type as keyof typeof classes] || 'text-gray-600 hover:text-gray-500'
}

// کلاس‌های اعلان‌ها
const getNotificationBorderClass = (type: string) => {
  const classes = {
    winner: 'border-green-400',
    completion: 'border-blue-400',
    error: 'border-red-400'
  }
  return classes[type as keyof typeof classes] || 'border-gray-400'
}

const getNotificationIconClass = (type: string) => {
  const classes = {
    winner: 'text-green-400',
    completion: 'text-blue-400',
    error: 'text-red-400'
  }
  return classes[type as keyof typeof classes] || 'text-gray-400'
}

const getNotificationTitleClass = (type: string) => {
  const classes = {
    winner: 'text-green-800',
    completion: 'text-blue-800',
    error: 'text-red-800'
  }
  return classes[type as keyof typeof classes] || 'text-gray-800'
}

// حذف هشدار
const removeAlert = (id: number) => {
  alerts.value = alerts.value.filter(alert => alert.id !== id)
}

// حذف اعلان
const removeNotification = (id: number) => {
  notifications.value = notifications.value.filter(notification => notification.id !== id)
}

interface AlertAction {
  handler?: () => void
  [key: string]: unknown
}

interface Alert {
  id?: number | string
  action?: AlertAction
  [key: string]: unknown
}

interface Notification {
  id?: number | string
  [key: string]: unknown
}

// اجرای عملیات هشدار
const handleAlertAction = (alert: Alert) => {
  if (alert.action && alert.action.handler) {
    alert.action.handler()
  }
  if (alert.id) {
    removeAlert(alert.id)
  }
}

// اضافه کردن هشدار جدید
const addAlert = (alert: Alert) => {
  const newAlert = {
    id: Date.now(),
    ...alert
  }
  alerts.value.unshift(newAlert)
  
  // حذف خودکار بعد از 10 ثانیه
  setTimeout(() => {
    if (newAlert.id) {
      removeAlert(newAlert.id)
    }
  }, 10000)
}

// اضافه کردن اعلان جدید
const addNotification = (notification: Notification) => {
  const newNotification = {
    id: Date.now(),
    timestamp: new Date(),
    ...notification
  }
  notifications.value.unshift(newNotification)
  
  // حذف خودکار بعد از 5 ثانیه
  setTimeout(() => {
    removeNotification(newNotification.id)
  }, 5000)
}

// Expose methods for parent components
defineExpose({
  addAlert,
  addNotification
})
</script> 
