<template>
  <div class="bg-white rounded-lg shadow p-6">
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-xl font-semibold text-gray-900">مدیریت اعلان‌ها</h2>
      <button 
        class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors"
        @click="showNotificationForm = true; editingNotification = null"
      >
        افزودن اعلان جدید
      </button>
    </div>

    <!-- لیست اعلان‌ها -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div 
        v-for="notification in notifications" 
        :key="notification.id"
        class="border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow"
      >
        <div class="flex justify-between items-start mb-3">
          <div class="flex items-center">
            <div :class="getNotificationIconColor(notification.type)" class="p-2 rounded-lg">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="getNotificationIcon(notification.type)" />
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">{{ notification.title }}</h3>
          </div>
          <span :class="notification.active ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'" class="px-2 py-1 text-xs font-semibold rounded-full">
            {{ notification.active ? 'فعال' : 'غیرفعال' }}
          </span>
        </div>
        <p class="text-gray-600 text-sm mb-3">{{ notification.message }}</p>
        <div class="space-y-2 text-sm">
          <div class="flex justify-between">
            <span class="text-gray-500">نوع:</span>
            <span class="font-semibold text-gray-900">{{ getNotificationTypeName(notification.type) }}</span>
          </div>
          <div class="flex justify-between">
            <span class="text-gray-500">تاریخ ارسال:</span>
            <span class="font-semibold text-gray-900">{{ formatDate(notification.date) }}</span>
          </div>
        </div>
        <div class="flex justify-end space-x-2 space-x-reverse mt-4">
          <button 
            class="text-blue-600 hover:text-blue-800 text-sm"
            @click="editingNotification = notification; showNotificationForm = true"
          >
            ویرایش
          </button>
          <button 
            :class="notification.active ? 'text-red-600 hover:text-red-800' : 'text-green-600 hover:text-green-800'"
            class="text-sm"
            @click="toggleNotificationStatus(notification)"
          >
            {{ notification.active ? 'غیرفعال' : 'فعال' }}
          </button>
        </div>
      </div>
    </div>

    <!-- مودال فرم اعلان -->
    <div v-if="showNotificationForm" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
      <div class="relative top-20 mx-auto p-5 border w-full max-w-md sm:max-w-lg md:max-w-xl shadow-lg rounded-md bg-white">
        <div class="mt-3">
          <h3 class="text-lg font-medium text-gray-900 mb-4">
            {{ editingNotification ? 'ویرایش اعلان' : 'افزودن اعلان جدید' }}
          </h3>
          <form @submit.prevent="saveNotification">
            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-700">نوع اعلان</label>
                <select 
                  v-model="notificationForm.type"
                  class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                  required
                >
                  <option value="level_up">ارتقا سطح</option>
                  <option value="points_earned">دریافت امتیاز</option>
                  <option value="points_expiry">انقضای امتیاز</option>
                  <option value="reward_earned">دریافت پاداش</option>
                  <option value="custom">سفارشی</option>
                </select>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700">عنوان اعلان</label>
                <input 
                  v-model="notificationForm.title" 
                  type="text" 
                  class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                  required
                >
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700">متن پیام</label>
                <textarea 
                  v-model="notificationForm.message" 
                  rows="2"
                  class="mt-1 block w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                  required
                ></textarea>
              </div>
            </div>
            <div class="flex justify-end space-x-3 space-x-reverse mt-6">
              <button 
                type="button"
                class="bg-gray-300 text-gray-700 px-4 py-2 rounded-md hover:bg-gray-400 transition-colors"
                @click="showNotificationForm = false"
              >
                انصراف
              </button>
              <button 
                type="submit"
                class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors"
              >
                {{ editingNotification ? 'ویرایش' : 'افزودن' }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'

const notifications = ref([
  {
    id: 1,
    type: 'level_up',
    title: 'ارتقا سطح',
    message: 'شما به سطح طلایی ارتقا یافتید!',
    date: '2024-01-10',
    active: true
  },
  {
    id: 2,
    type: 'points_earned',
    title: 'دریافت امتیاز',
    message: 'شما 100 امتیاز جدید دریافت کردید.',
    date: '2024-01-09',
    active: true
  },
  {
    id: 3,
    type: 'points_expiry',
    title: 'انقضای امتیاز',
    message: 'برخی از امتیازات شما در حال انقضا هستند.',
    date: '2024-01-08',
    active: true
  },
  {
    id: 4,
    type: 'reward_earned',
    title: 'دریافت پاداش',
    message: 'شما یک پاداش جدید دریافت کردید.',
    date: '2024-01-07',
    active: true
  }
])

const showNotificationForm = ref(false)
const editingNotification = ref(null)

const notificationForm = reactive({
  type: 'level_up',
  title: '',
  message: ''
})

function getNotificationIconColor(type) {
  const colors = {
    level_up: 'bg-purple-500',
    points_earned: 'bg-green-500',
    points_expiry: 'bg-yellow-500',
    reward_earned: 'bg-blue-500',
    custom: 'bg-gray-500'
  }
  return colors[type] || 'bg-gray-500'
}

function getNotificationIcon(type) {
  const icons = {
    level_up: 'M5 13l4 4L19 7',
    points_earned: 'M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2',
    points_expiry: 'M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z',
    reward_earned: 'M9 12l2 2 4-4',
    custom: 'M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z'
  }
  return icons[type] || icons.custom
}

function getNotificationTypeName(type) {
  const names = {
    level_up: 'ارتقا سطح',
    points_earned: 'دریافت امتیاز',
    points_expiry: 'انقضای امتیاز',
    reward_earned: 'دریافت پاداش',
    custom: 'سفارشی'
  }
  return names[type] || 'سفارشی'
}

function formatDate(dateString) {
  const date = new Date(dateString)
  return date.toLocaleDateString('fa-IR')
}

function toggleNotificationStatus(notification) {
  notification.active = !notification.active
}

function saveNotification() {
  if (editingNotification.value) {
    const index = notifications.value.findIndex(n => n.id === editingNotification.value.id)
    if (index !== -1) {
      notifications.value[index] = { ...editingNotification.value, ...notificationForm }
    }
  } else {
    const newNotification = {
      id: Date.now(),
      ...notificationForm,
      date: new Date().toISOString().slice(0, 10),
      active: true
    }
    notifications.value.push(newNotification)
  }
  Object.assign(notificationForm, { type: 'level_up', title: '', message: '' })
  showNotificationForm.value = false
  editingNotification.value = null
}
</script> 
