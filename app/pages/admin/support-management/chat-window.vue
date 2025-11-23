<template>
  <div class="chat-window" dir="rtl">
    <!-- Header -->
    <AdminChatChatHeader 
      :waiting-count="waitingUsers.length"
      @start-drag="startDrag"
      @minimize="minimizeWindow"
      @close="closeWindow"
    />

    <!-- Telegram Layout -->
    <div class="telegram-layout">
      <!-- Right Sidebar (optional) -->
      <!-- Sidebar component not available; layout simplified to chat area only. -->

      <!-- Left Chat Area -->
      <AdminChatChatArea 
        :selected-user="selectedUser"
        :messages="messages"
        :new-message="newMessage"
        @end-chat="endChat"
        @transfer-chat="transferChat"
        @attach-file="attachFile"
        @send-message="sendMessage"
        @voice-message="handleVoiceMessage"
        @update:new-message="newMessage = $event"
      />
    </div>

    <!-- Notifications -->
    <AdminChatChatNotification :notification="notification" />
    
    <!-- File Upload Modal -->
    <AdminChatFileUploader 
      :show="showFileUploader"
      @close="showFileUploader = false"
      @upload="handleFileUpload"
    />
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string | false; middleware?: string | string[] }) => void
</script>

<script setup lang="ts">
import { onMounted, onUnmounted, ref } from 'vue'
// استفاده از auto-import Nuxt: کامپوننت‌های پوشه components/admin/chat با پیشوند AdminChat*

// Disable layout
definePageMeta({
  layout: false,
  middleware: 'admin'
})

const waitingUsers = ref([])

// Reactive data
const selectedUser = ref(null)
const newMessage = ref('')
const notification = ref(null)
const showFileUploader = ref(false)

// Sample data - now handled by Sidebar component

const messages = ref([
  {
    id: 1,
    type: 'text',
    text: 'سلام، ممکنه کمک کنید؟',
    sender: 'user',
    time: '14:30'
  },
  {
    id: 2,
    type: 'text',
    text: 'سلام! البته، چطور می‌تونم کمکتون کنم؟',
    sender: 'admin',
    time: '14:31'
  },
  {
    id: 3,
    type: 'file',
    fileName: 'گزارش فروش.pdf',
    fileSize: 1024000,
    fileType: 'application/pdf',
    fileUrl: '#',
    sender: 'user',
    time: '14:32'
  },
  {
    id: 4,
    type: 'voice',
    audioUrl: 'data:audio/wav;base64,',
    duration: 15,
    sender: 'admin', 
    time: '14:33'
  }
])

// Computed - now handled by Sidebar component

// Methods
const selectUser = (chat) => {
  selectedUser.value = {
    id: chat.id,
    name: `کاربر ${chat.userId}`,
    location: 'ایران',
    device: 'موبایل',
    browser: 'Chrome'
  }
  showNotification('شروع چت با ' + selectedUser.value.name, 'success')
}
const _selectUser = selectUser

const endChat = () => {
  if (selectedUser.value) {
    showNotification('چت با ' + selectedUser.value.name + ' پایان یافت', 'info')
    selectedUser.value = null
  }
}

const transferChat = () => {
  if (selectedUser.value) {
    showNotification('چت انتقال داده شد', 'warning')
  }
}

const attachFile = () => {
  showFileUploader.value = true
}

const handleFileUpload = (files) => {
  files.forEach(file => {
    const fileMessage = {
      id: Date.now() + Math.random(),
      type: 'file',
      fileName: file.name,
      fileSize: file.size,
      fileType: file.type,
      fileUrl: file.url,
      sender: 'admin',
      time: new Date().toLocaleTimeString('fa-IR', { hour: '2-digit', minute: '2-digit' })
    }
    messages.value.push(fileMessage)
  })
  showNotification(`${files.length} فایل ارسال شد`, 'success')
}

const handleVoiceMessage = (voiceData) => {
  const voiceMessage = {
    id: Date.now(),
    type: 'voice',
    audioUrl: voiceData.url,
    duration: voiceData.duration,
    sender: 'admin',
    time: new Date().toLocaleTimeString('fa-IR', { hour: '2-digit', minute: '2-digit' })
  }
  messages.value.push(voiceMessage)
  showNotification('پیام صوتی ارسال شد', 'success')
}

const sendMessage = (messageText) => {
  if (!selectedUser.value) return
  
  const message = {
    id: Date.now(),
    type: 'text',
    text: messageText,
    sender: 'admin',
    time: new Date().toLocaleTimeString('fa-IR', { hour: '2-digit', minute: '2-digit' })
  }
  
  messages.value.push(message)
  newMessage.value = ''
  
  // Auto response simulation
  setTimeout(() => {
    const response = {
      id: Date.now() + 1,
      type: 'text',
      text: 'متشکرم از پاسخ شما',
      sender: 'user',
      time: new Date().toLocaleTimeString('fa-IR', { hour: '2-digit', minute: '2-digit' })
    }
    messages.value.push(response)
  }, 1000)
}

const showNotification = (message, type = 'info') => {
  notification.value = { message, type }
  setTimeout(() => {
    notification.value = null
  }, 3000)
}

// Event handlers for Sidebar
const handleCategoryChange = (category) => {
  showNotification(`دسته‌بندی به ${category} تغییر کرد`, 'info')
}
const _handleCategoryChange = handleCategoryChange

const handleSearchChange = (_query) => {
}
const _handleSearchChange = handleSearchChange

// Window controls
const isDragging = ref(false)
const dragOffset = ref({ x: 0, y: 0 })

const startDrag = (event) => {
  isDragging.value = true
  dragOffset.value = {
    x: event.clientX - window.screenX,
    y: event.clientY - window.screenY
  }
  
  document.addEventListener('mousemove', drag)
  document.addEventListener('mouseup', stopDrag)
}

const drag = (event) => {
  if (isDragging.value) {
    window.moveTo(
      event.screenX - dragOffset.value.x,
      event.screenY - dragOffset.value.y
    )
  }
}

const stopDrag = () => {
  isDragging.value = false
  document.removeEventListener('mousemove', drag)
  document.removeEventListener('mouseup', stopDrag)
}

const minimizeWindow = () => {
  if (import.meta.client) {
    window.blur()
  }
}

const closeWindow = () => {
  if (import.meta.client) {
    window.close()
  }
}

// Cleanup
onUnmounted(() => {
  document.removeEventListener('mousemove', drag)
  document.removeEventListener('mouseup', stopDrag)
})

onMounted(() => {
  document.title = 'چت ادمین - پیشرفته'
  document.body.style.overflow = 'hidden'
  document.body.style.margin = '0'
  document.body.style.padding = '0'
})
</script>

<style scoped>
.chat-window {
  width: 100vw;
  height: 100vh;
  display: flex;
  flex-direction: column;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  font-family: 'IRANSans', sans-serif;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3);
}

.telegram-layout {
  flex: 1;
  display: flex;
  height: 100%;
  overflow: hidden;
}
</style> 
