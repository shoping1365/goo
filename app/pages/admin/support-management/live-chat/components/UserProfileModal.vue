<template>
  <div v-if="isOpen && user" class="fixed inset-0 z-50 overflow-y-auto" dir="rtl">
    <div class="flex items-center justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
      <!-- Background overlay -->
      <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" @click="$emit('close')"></div>

      <!-- Modal panel -->
      <div class="inline-block align-bottom bg-white rounded-2xl text-right overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-2xl sm:w-full">
        <div class="bg-white px-6 py-6">
          <!-- Header -->
          <div class="flex items-center justify-between mb-6">
            <h3 class="text-2xl font-bold text-gray-900">پروفایل کاربر</h3>
            <button 
              class="text-gray-400 hover:text-gray-600 transition-colors"
              @click="$emit('close')"
            >
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
              </svg>
            </button>
          </div>

          <!-- User Info -->
          <div class="flex items-center space-x-4 space-x-reverse mb-6">
            <div class="relative">
              <img 
                :src="user.avatar || '/avatars/default.jpg'" 
                :alt="user.name"
                class="w-20 h-20 rounded-full object-cover border-4 border-gray-200"
              >
              <div 
                :class="[
                  user.isOnline ? 'bg-green-400' : 'bg-gray-400',
                  'absolute -bottom-1 -right-1 w-6 h-6 rounded-full border-4 border-white'
                ]"
              ></div>
            </div>
            <div class="flex-1">
              <h4 class="text-xl font-bold text-gray-900">{{ user.name }}</h4>
              <p class="text-gray-600">{{ user.email }}</p>
              <div class="flex items-center mt-2">
                <span 
                  :class="[
                    user.isOnline ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800',
                    'px-2 py-1 text-xs font-medium rounded-full'
                  ]"
                >
                  {{ user.status }}
                </span>
                <span class="text-xs text-gray-500 mr-3">
                  آخرین بازدید: {{ formatLastSeen(user.lastSeen) }}
                </span>
              </div>
            </div>
          </div>

          <!-- User Stats -->
          <div class="grid grid-cols-3 gap-6 mb-6">
            <div class="bg-blue-50 rounded-lg p-6 text-center">
              <div class="text-2xl font-bold text-blue-600">{{ userStats.totalMessages }}</div>
              <div class="text-sm text-blue-600">کل پیام‌ها</div>
            </div>
            <div class="bg-green-50 rounded-lg p-6 text-center">
              <div class="text-2xl font-bold text-green-600">{{ userStats.avgResponseTime }}</div>
              <div class="text-sm text-green-600">میانگین پاسخ (دقیقه)</div>
            </div>
            <div class="bg-purple-50 rounded-lg p-6 text-center">
              <div class="text-2xl font-bold text-purple-600">{{ userStats.satisfactionRate }}%</div>
              <div class="text-sm text-purple-600">رضایت</div>
            </div>
          </div>

          <!-- Quick Actions -->
          <div class="flex space-x-3 space-x-reverse mb-6">
            <button 
              class="flex-1 px-4 py-2 text-sm font-medium text-red-700 bg-red-100 border border-red-300 rounded-lg hover:bg-red-200 transition-colors"
              @click="blockUser"
            >
              <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728L5.636 5.636m12.728 12.728L18.364 5.636M5.636 18.364l12.728-12.728"></path>
              </svg>
              بلاک کردن
            </button>
            <button 
              class="flex-1 px-4 py-2 text-sm font-medium text-blue-700 bg-blue-100 border border-blue-300 rounded-lg hover:bg-blue-200 transition-colors"
              @click="exportChat"
            >
              <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
              </svg>
              خروجی چت
            </button>
            <button 
              class="flex-1 px-4 py-2 text-sm font-medium text-green-700 bg-green-100 border border-green-300 rounded-lg hover:bg-green-200 transition-colors"
              @click="createTicket"
            >
              <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
              </svg>
              تیکت جدید
            </button>
          </div>

          <!-- Recent Messages -->
          <div>
            <h5 class="text-lg font-medium text-gray-900 mb-4">آخرین پیام‌ها</h5>
            <div class="space-y-3 max-h-64 overflow-y-auto">
              <div 
                v-for="message in recentMessages" 
                :key="message.id"
                class="flex items-start space-x-3 space-x-reverse p-3 bg-gray-50 rounded-lg"
              >
                <div 
                  :class="[
                    message.sender === 'user' ? 'bg-blue-500' : 'bg-gray-500',
                    'w-2 h-2 rounded-full mt-2 flex-shrink-0'
                  ]"
                ></div>
                <div class="flex-1 min-w-0">
                  <div class="text-sm text-gray-900">{{ message.content }}</div>
                  <div class="text-xs text-gray-500 mt-1">{{ formatTime(message.timestamp) }}</div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Footer -->
        <div class="bg-gray-50 px-6 py-4 flex justify-end">
          <button 
            class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
            @click="$emit('close')"
          >
            بستن
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
defineEmits(['close'])

interface User {
  name: string
  email: string
  avatar?: string
  isOnline: boolean
  status: string
  lastSeen: Date
}

interface UserStats {
  totalMessages: number | string
  avgResponseTime: number | string
  satisfactionRate: number | string
}

interface Message {
  id: number | string
  sender: string
  content: string
  timestamp: Date
}

defineProps<{
  isOpen: boolean
  user: User
  userStats: UserStats
  recentMessages: Message[]
}>()

// Methods
const formatLastSeen = (timestamp: Date): string => {
  const now = new Date()
  const diff = now.getTime() - timestamp.getTime()
  const minutes = Math.floor(diff / 60000)
  const hours = Math.floor(diff / 3600000)
  const days = Math.floor(diff / 86400000)

  if (days > 0) return `${days} روز پیش`
  if (hours > 0) return `${hours} ساعت پیش`
  if (minutes > 0) return `${minutes} دقیقه پیش`
  return 'همین الان'
}

const formatTime = (timestamp: Date): string => {
  return timestamp.toLocaleTimeString('fa-IR', { 
    hour: '2-digit', 
    minute: '2-digit' 
  })
}

const blockUser = () => {
  // console.log('Blocking user:', props.user?.name)
  // Implementation for blocking user
}

const exportChat = () => {
  // console.log('Exporting chat for user:', props.user?.name)
  // Implementation for exporting chat
}

const createTicket = () => {
  // console.log('Creating ticket for user:', props.user?.name)
  // Implementation for creating ticket
}
</script>

<style scoped>
/* Custom styles for modal */
</style>
