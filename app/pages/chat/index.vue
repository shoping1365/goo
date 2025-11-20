<template>
  <div class="chat-container" dir="rtl">
    <!-- Sidebar - لیست مکالمات -->
    <div class="chat-sidebar">
      <!-- Header -->
      <div class="sidebar-header">
        <h2 class="sidebar-title">
          <i class="i-heroicons-chat-bubble-left-right"></i>
          چت آنلاین
        </h2>
        <button class="btn-new-chat" @click="startNewChat">
          <i class="i-heroicons-plus"></i>
        </button>
      </div>

      <!-- Search -->
      <div class="search-box">
        <i class="i-heroicons-magnifying-glass search-icon"></i>
        <input
          v-model="searchQuery"
          type="text"
          placeholder="جستجوی مکالمات..."
          class="search-input"
        >
      </div>

      <!-- Tabs -->
      <div class="chat-tabs">
        <button
          v-for="tab in tabs"
          :key="tab.key"
          :class="['tab-btn', { active: activeTab === tab.key }]"
          @click="activeTab = tab.key"
        >
          {{ tab.label }}
          <span v-if="tab.count" class="tab-count">{{ tab.count }}</span>
        </button>
      </div>

      <!-- Chat List -->
      <div class="chat-list">
        <div
          v-for="chat in filteredChats"
          :key="chat.id"
          :class="['chat-item', { active: selectedChat?.id === chat.id, unread: chat.unread > 0 }]"
          @click="selectChat(chat)"
        >
          <div class="chat-avatar">
            <img v-if="chat.avatar" :src="chat.avatar" :alt="chat.name">
            <div v-else class="avatar-placeholder">
              {{ chat.name.charAt(0) }}
            </div>
            <span v-if="chat.online" class="online-badge"></span>
          </div>
          
          <div class="chat-info">
            <div class="chat-header">
              <h3 class="chat-name">{{ chat.name }}</h3>
              <span class="chat-time">{{ formatTime(chat.lastMessageTime) }}</span>
            </div>
            <div class="chat-preview">
              <p class="chat-message">{{ chat.lastMessage }}</p>
              <span v-if="chat.unread > 0" class="unread-badge">{{ chat.unread }}</span>
            </div>
          </div>
        </div>

        <div v-if="filteredChats.length === 0" class="empty-state">
          <i class="i-heroicons-chat-bubble-left-ellipsis"></i>
          <p>مکالمه‌ای یافت نشد</p>
        </div>
      </div>
    </div>

    <!-- Main Chat Area -->
    <div class="chat-main">
      <template v-if="selectedChat">
        <!-- Chat Header -->
        <div class="chat-main-header">
          <div class="chat-user-info">
            <div class="chat-avatar">
              <img v-if="selectedChat.avatar" :src="selectedChat.avatar" :alt="selectedChat.name">
              <div v-else class="avatar-placeholder">
                {{ selectedChat.name.charAt(0) }}
              </div>
              <span v-if="selectedChat.online" class="online-badge"></span>
            </div>
            <div>
              <h3 class="user-name">{{ selectedChat.name }}</h3>
              <p class="user-status">{{ selectedChat.online ? 'آنلاین' : 'آفلاین' }}</p>
            </div>
          </div>
          
          <div class="chat-actions">
            <button class="action-btn" title="جستجو">
              <i class="i-heroicons-magnifying-glass"></i>
            </button>
            <button class="action-btn" title="تماس صوتی">
              <i class="i-heroicons-phone"></i>
            </button>
            <button class="action-btn" title="تماس تصویری">
              <i class="i-heroicons-video-camera"></i>
            </button>
            <button class="action-btn" title="منو">
              <i class="i-heroicons-ellipsis-vertical"></i>
            </button>
          </div>
        </div>

        <!-- Messages Area -->
        <div ref="messagesContainer" class="messages-area">
          <div
            v-for="message in currentMessages"
            :key="message.id"
            :class="['message', message.sender]"
          >
            <div class="message-content">
              <!-- Text Message -->
              <div v-if="message.type === 'text'" class="message-bubble">
                <p class="message-text">{{ message.text }}</p>
                <span class="message-time">{{ message.time }}</span>
                <i v-if="message.sender === 'me'" :class="['message-status', getMessageStatusIcon(message.status)]"></i>
              </div>

              <!-- Image Message -->
              <div v-else-if="message.type === 'image'" class="message-image">
                <img :src="message.imageUrl" :alt="message.caption || 'تصویر'">
                <p v-if="message.caption" class="image-caption">{{ message.caption }}</p>
                <span class="message-time">{{ message.time }}</span>
              </div>

              <!-- File Message -->
              <div v-else-if="message.type === 'file'" class="message-file">
                <div class="file-icon">
                  <i class="i-heroicons-document"></i>
                </div>
                <div class="file-info">
                  <p class="file-name">{{ message.fileName }}</p>
                  <p class="file-size">{{ formatFileSize(message.fileSize) }}</p>
                </div>
                <button class="file-download">
                  <i class="i-heroicons-arrow-down-tray"></i>
                </button>
                <span class="message-time">{{ message.time }}</span>
              </div>

              <!-- Voice Message -->
              <div v-else-if="message.type === 'voice'" class="message-voice">
                <button class="voice-play">
                  <i class="i-heroicons-play"></i>
                </button>
                <div class="voice-wave"></div>
                <span class="voice-duration">{{ message.duration }}</span>
                <span class="message-time">{{ message.time }}</span>
              </div>
            </div>
          </div>

          <div v-if="isTyping" class="typing-indicator">
            <span></span>
            <span></span>
            <span></span>
          </div>
        </div>

        <!-- Input Area -->
        <div class="input-area">
          <button class="input-btn" title="پیوست فایل" @click="showFileMenu = !showFileMenu">
            <i class="i-heroicons-paper-clip"></i>
          </button>
          
          <div v-if="showFileMenu" class="file-menu">
            <button @click="attachImage">
              <i class="i-heroicons-photo"></i>
              تصویر
            </button>
            <button @click="attachFile">
              <i class="i-heroicons-document"></i>
              فایل
            </button>
            <button @click="attachVideo">
              <i class="i-heroicons-video-camera"></i>
              ویدیو
            </button>
          </div>

          <input
            ref="fileInput"
            type="file"
            class="hidden"
            @change="handleFileSelect"
          >

          <div class="input-wrapper">
            <textarea
              v-model="messageText"
              placeholder="پیام خود را بنویسید..."
              class="message-input"
              rows="1"
              @keydown.enter.exact.prevent="sendMessage"
              @input="autoResize"
            ></textarea>
          </div>

          <button class="input-btn" title="ایموجی" @click="showEmojiPicker = !showEmojiPicker">
            <i class="i-heroicons-face-smile"></i>
          </button>

          <button
            v-if="messageText.trim()"
            class="send-btn"
            @click="sendMessage"
          >
            <i class="i-heroicons-paper-airplane"></i>
          </button>
          <button
            v-else
            class="input-btn voice-btn"
            title="پیام صوتی"
            @mousedown="startVoiceRecording"
            @mouseup="stopVoiceRecording"
            @mouseleave="cancelVoiceRecording"
          >
            <i class="i-heroicons-microphone"></i>
          </button>
        </div>

        <!-- Emoji Picker -->
        <div v-if="showEmojiPicker" class="emoji-picker">
          <div class="emoji-grid">
            <button
              v-for="emoji in emojis"
              :key="emoji"
              class="emoji-btn"
              @click="insertEmoji(emoji)"
            >
              {{ emoji }}
            </button>
          </div>
        </div>
      </template>

      <!-- Empty State -->
      <div v-else class="chat-empty-state">
        <i class="i-heroicons-chat-bubble-left-right"></i>
        <h3>یک مکالمه را انتخاب کنید</h3>
        <p>برای شروع چت، یک مکالمه از لیست سمت راست انتخاب کنید</p>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string[] }) => void
</script>

<script setup lang="ts">
import { computed, nextTick, ref, watch } from 'vue';

interface Chat {
  id: number
  name: string
  avatar: string
  lastMessage: string
  lastMessageTime: number
  unread: number
  online: boolean
  type: 'user' | 'group'
}

interface Message {
  id: number
  type: 'text' | 'image' | 'file' | 'voice'
  text?: string
  sender: 'me' | 'them'
  time: string
  status: 'sent' | 'delivered' | 'read'
  fileSize?: number
  fileName?: string
  duration?: string
  imageUrl?: string
  caption?: string
}

definePageMeta({
  layout: 'admin',
  middleware: ['auth']
})

// Data
const searchQuery = ref('')
const activeTab = ref('all')
const selectedChat = ref<Chat | null>(null)
const messageText = ref('')
const showFileMenu = ref(false)
const showEmojiPicker = ref(false)
const isTyping = ref(false)
const isRecording = ref(false)
const messagesContainer = ref<HTMLElement>()
const fileInput = ref<HTMLInputElement>()

// Tabs
const tabs = [
  { key: 'all', label: 'همه', count: 0 },
  { key: 'unread', label: 'خوانده نشده', count: 3 },
  { key: 'groups', label: 'گروه‌ها', count: 0 },
  { key: 'archived', label: 'آرشیو', count: 0 }
]

// Sample Chats Data
const chats = ref<Chat[]>([
  {
    id: 1,
    name: 'محمد احمدی',
    avatar: '',
    lastMessage: 'سلام، وضعیت سفارش من چطوره؟',
    lastMessageTime: new Date().getTime() - 5 * 60 * 1000,
    unread: 2,
    online: true,
    type: 'user'
  },
  {
    id: 2,
    name: 'فاطمه کریمی',
    avatar: '',
    lastMessage: 'ممنون از پشتیبانی عالیتون',
    lastMessageTime: new Date().getTime() - 30 * 60 * 1000,
    unread: 0,
    online: false,
    type: 'user'
  },
  {
    id: 3,
    name: 'علی رضایی',
    avatar: '',
    lastMessage: 'کی محصول ارسال میشه؟',
    lastMessageTime: new Date().getTime() - 2 * 60 * 60 * 1000,
    unread: 1,
    online: true,
    type: 'user'
  },
  {
    id: 4,
    name: 'گروه پشتیبانی',
    avatar: '',
    lastMessage: 'یک تیکت جدید باز شد',
    lastMessageTime: new Date().getTime() - 5 * 60 * 60 * 1000,
    unread: 0,
    online: true,
    type: 'group'
  }
])

// Sample Messages
const messages = ref<{ [key: number]: Message[] }>({
  1: [
    {
      id: 1,
      type: 'text',
      text: 'سلام، وضعیت سفارش من چطوره؟',
      sender: 'them',
      time: '14:20',
      status: 'read'
    },
    {
      id: 2,
      type: 'text',
      text: 'سلام! بذارید چک کنم...',
      sender: 'me',
      time: '14:21',
      status: 'read'
    },
    {
      id: 3,
      type: 'text',
      text: 'سفارش شما در مرحله بسته‌بندی است و فردا ارسال می‌شود.',
      sender: 'me',
      time: '14:22',
      status: 'delivered'
    }
  ],
  2: [
    {
      id: 1,
      type: 'text',
      text: 'ممنون از پشتیبانی عالیتون',
      sender: 'them',
      time: '13:45',
      status: 'read'
    }
  ],
  3: [
    {
      id: 1,
      type: 'text',
      text: 'کی محصول ارسال میشه؟',
      sender: 'them',
      time: '12:30',
      status: 'sent'
    }
  ]
})

// Emojis
const emojis = ['😀', '😃', '😄', '😁', '😅', '😂', '🤣', '😊', '😇', '🙂', '😉', '😌', '😍', '🥰', '😘', '😗', '😙', '😚', '😋', '😛', '😝', '😜', '🤪', '🤨', '🧐', '🤓', '😎', '🥸', '🤩', '🥳', '😏', '😒', '😞', '😔', '😟', '😕', '🙁', '☹️', '😣', '😖', '😫', '😩', '🥺', '😢', '😭', '😤', '😠', '😡', '🤬', '🤯', '😳', '🥵', '🥶', '😱', '😨', '😰', '😥', '😓', '🤗', '🤔', '🤭', '🤫', '🤥', '😶', '😐', '😑', '😬', '🙄', '😯', '😦', '😧', '😮', '😲', '🥱', '😴', '🤤', '😪', '😵', '🤐', '🥴', '🤢', '🤮', '🤧', '😷', '🤒', '🤕', '🤑', '🤠', '👍', '👎', '👌', '✌️', '🤞', '🤟', '🤘', '🤙', '👈', '👉', '👆', '👇', '☝️', '👏', '🙌', '👐', '🤲', '🙏', '✍️', '💪', '🦾', '🦿', '🦵', '🦶', '👂', '🦻', '👃', '🧠', '🫀', '🫁', '🦷', '🦴', '👀', '👁️', '👅', '👄', '💋', '🩸']

// Computed
const filteredChats = computed(() => {
  let filtered = chats.value

  // Filter by search
  if (searchQuery.value) {
    filtered = filtered.filter(chat =>
      chat.name.includes(searchQuery.value) ||
      chat.lastMessage.includes(searchQuery.value)
    )
  }

  // Filter by tab
  if (activeTab.value === 'unread') {
    filtered = filtered.filter(chat => chat.unread > 0)
  } else if (activeTab.value === 'groups') {
    filtered = filtered.filter(chat => chat.type === 'group')
  }

  return filtered
})

const currentMessages = computed(() => {
  if (!selectedChat.value) return []
  return messages.value[selectedChat.value.id] || []
})

// Methods
const selectChat = (chat: Chat) => {
  selectedChat.value = chat
  chat.unread = 0
  nextTick(() => {
    scrollToBottom()
  })
}

const startNewChat = () => {
  // TODO: Implement new chat modal
  // console.log('Start new chat')
}

const sendMessage = () => {
  if (!messageText.value.trim() || !selectedChat.value) return

  const newMessage: Message = {
    id: Date.now(),
    type: 'text',
    text: messageText.value,
    sender: 'me',
    time: new Date().toLocaleTimeString('fa-IR', { hour: '2-digit', minute: '2-digit' }),
    status: 'sent'
  }

  if (!messages.value[selectedChat.value.id]) {
    messages.value[selectedChat.value.id] = []
  }

  messages.value[selectedChat.value.id].push(newMessage)
  messageText.value = ''
  
  // Update chat last message
  const chat = chats.value.find(c => c.id === selectedChat.value.id)
  if (chat) {
    chat.lastMessage = newMessage.text
    chat.lastMessageTime = Date.now()
  }

  nextTick(() => {
    scrollToBottom()
  })
}

const scrollToBottom = () => {
  if (messagesContainer.value) {
    messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight
  }
}

const autoResize = (event: Event) => {
  const textarea = event.target as HTMLTextAreaElement
  textarea.style.height = 'auto'
  textarea.style.height = Math.min(textarea.scrollHeight, 120) + 'px'
}

const formatTime = (timestamp: number) => {
  const now = Date.now()
  const diff = now - timestamp
  const minutes = Math.floor(diff / 60000)
  const hours = Math.floor(diff / 3600000)
  const days = Math.floor(diff / 86400000)

  if (minutes < 1) return 'الان'
  if (minutes < 60) return `${minutes} دقیقه پیش`
  if (hours < 24) return `${hours} ساعت پیش`
  if (days < 7) return `${days} روز پیش`
  
  return new Date(timestamp).toLocaleDateString('fa-IR')
}

const formatFileSize = (bytes: number) => {
  if (bytes < 1024) return bytes + ' B'
  if (bytes < 1024 * 1024) return (bytes / 1024).toFixed(1) + ' KB'
  return (bytes / (1024 * 1024)).toFixed(1) + ' MB'
}

const getMessageStatusIcon = (status: string) => {
  switch (status) {
    case 'sent': return 'i-heroicons-check'
    case 'delivered': return 'i-heroicons-check-check'
    case 'read': return 'i-heroicons-check-check text-blue-500'
    default: return 'i-heroicons-clock'
  }
}

const insertEmoji = (emoji: string) => {
  messageText.value += emoji
  showEmojiPicker.value = false
}

const attachImage = () => {
  if (fileInput.value) {
    fileInput.value.accept = 'image/*'
    fileInput.value.click()
  }
  showFileMenu.value = false
}

const attachFile = () => {
  if (fileInput.value) {
    fileInput.value.accept = '*/*'
    fileInput.value.click()
  }
  showFileMenu.value = false
}

const attachVideo = () => {
  if (fileInput.value) {
    fileInput.value.accept = 'video/*'
    fileInput.value.click()
  }
  showFileMenu.value = false
}

const handleFileSelect = (event: Event) => {
  const input = event.target as HTMLInputElement
  const file = input.files?.[0]
  if (!file) return

  // TODO: Upload file and add to messages
  // console.log('Selected file:', file)
}

const startVoiceRecording = () => {
  isRecording.value = true
  // TODO: Implement voice recording
  // console.log('Start recording')
}

const stopVoiceRecording = () => {
  if (isRecording.value) {
    isRecording.value = false
    // TODO: Stop recording and send
    // console.log('Stop recording')
  }
}

const cancelVoiceRecording = () => {
  if (isRecording.value) {
    isRecording.value = false
    // TODO: Cancel recording
    // console.log('Cancel recording')
  }
}

// Watch for typing
watch(messageText, (newVal) => {
  if (newVal && selectedChat.value) {
    // TODO: Send typing indicator to server
  }
})
</script>

<style scoped>
.chat-container {
  display: flex;
  height: calc(100vh - 64px);
  background: #f5f5f5;
}

/* Sidebar */
.chat-sidebar {
  width: 380px;
  background: white;
  border-left: 1px solid #e5e7eb;
  display: flex;
  flex-direction: column;
}

.sidebar-header {
  padding: 20px;
  border-bottom: 1px solid #e5e7eb;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.sidebar-title {
  font-size: 20px;
  font-weight: 700;
  display: flex;
  align-items: center;
  gap: 8px;
  margin: 0;
}

.btn-new-chat {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background: #3b82f6;
  color: white;
  border: none;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-new-chat:hover {
  background: #2563eb;
  transform: scale(1.05);
}

.search-box {
  padding: 16px 20px;
  border-bottom: 1px solid #e5e7eb;
  position: relative;
}

.search-icon {
  position: absolute;
  right: 36px;
  top: 50%;
  transform: translateY(-50%);
  color: #9ca3af;
  font-size: 20px;
}

.search-input {
  width: 100%;
  padding: 10px 16px 10px 40px;
  border: 1px solid #e5e7eb;
  border-radius: 24px;
  font-size: 14px;
  outline: none;
  transition: all 0.2s;
}

.search-input:focus {
  border-color: #3b82f6;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
}

.chat-tabs {
  display: flex;
  padding: 0 20px;
  gap: 8px;
  border-bottom: 1px solid #e5e7eb;
  background: white;
}

.tab-btn {
  padding: 12px 16px;
  border: none;
  background: none;
  font-size: 14px;
  font-weight: 500;
  color: #6b7280;
  cursor: pointer;
  border-bottom: 2px solid transparent;
  transition: all 0.2s;
  display: flex;
  align-items: center;
  gap: 6px;
}

.tab-btn:hover {
  color: #3b82f6;
}

.tab-btn.active {
  color: #3b82f6;
  border-bottom-color: #3b82f6;
}

.tab-count {
  background: #3b82f6;
  color: white;
  padding: 2px 8px;
  border-radius: 12px;
  font-size: 12px;
}

.chat-list {
  flex: 1;
  overflow-y: auto;
}

.chat-item {
  display: flex;
  gap: 12px;
  padding: 16px 20px;
  cursor: pointer;
  transition: all 0.2s;
  border-bottom: 1px solid #f3f4f6;
}

.chat-item:hover {
  background: #f9fafb;
}

.chat-item.active {
  background: #eff6ff;
  border-right: 3px solid #3b82f6;
}

.chat-item.unread {
  background: #fefefe;
}

.chat-avatar {
  position: relative;
  width: 50px;
  height: 50px;
  flex-shrink: 0;
}

.chat-avatar img,
.avatar-placeholder {
  width: 100%;
  height: 100%;
  border-radius: 50%;
  object-fit: cover;
}

.avatar-placeholder {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
  font-weight: 600;
}

.online-badge {
  position: absolute;
  bottom: 2px;
  left: 2px;
  width: 12px;
  height: 12px;
  background: #10b981;
  border: 2px solid white;
  border-radius: 50%;
}

.chat-info {
  flex: 1;
  min-width: 0;
}

.chat-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 4px;
}

.chat-name {
  font-size: 15px;
  font-weight: 600;
  margin: 0;
  color: #111827;
}

.chat-time {
  font-size: 12px;
  color: #9ca3af;
}

.chat-preview {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.chat-message {
  font-size: 14px;
  color: #6b7280;
  margin: 0;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.unread-badge {
  background: #3b82f6;
  color: white;
  padding: 2px 8px;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 600;
  flex-shrink: 0;
}

.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 60px 20px;
  color: #9ca3af;
}

.empty-state i {
  font-size: 64px;
  margin-bottom: 16px;
}

/* Main Chat */
.chat-main {
  flex: 1;
  display: flex;
  flex-direction: column;
  background: white;
}

.chat-main-header {
  padding: 16px 24px;
  border-bottom: 1px solid #e5e7eb;
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: white;
}

.chat-user-info {
  display: flex;
  align-items: center;
  gap: 12px;
}

.user-name {
  font-size: 16px;
  font-weight: 600;
  margin: 0;
  color: #111827;
}

.user-status {
  font-size: 13px;
  color: #10b981;
  margin: 0;
}

.chat-actions {
  display: flex;
  gap: 8px;
}

.action-btn {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  border: none;
  background: #f3f4f6;
  color: #6b7280;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.2s;
  font-size: 20px;
}

.action-btn:hover {
  background: #e5e7eb;
  color: #374151;
}

.messages-area {
  flex: 1;
  overflow-y: auto;
  padding: 24px;
  background: #f9fafb;
}

.message {
  display: flex;
  margin-bottom: 16px;
}

.message.me {
  justify-content: flex-start;
}

.message.them {
  justify-content: flex-end;
}

.message-content {
  max-width: 60%;
}

.message-bubble {
  padding: 12px 16px;
  border-radius: 18px;
  position: relative;
}

.message.me .message-bubble {
  background: #3b82f6;
  color: white;
  border-bottom-right-radius: 4px;
}

.message.them .message-bubble {
  background: white;
  color: #111827;
  border-bottom-left-radius: 4px;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.05);
}

.message-text {
  margin: 0;
  font-size: 14px;
  line-height: 1.5;
  word-wrap: break-word;
}

.message-time {
  font-size: 11px;
  opacity: 0.7;
  margin-right: 8px;
}

.message-status {
  font-size: 14px;
  margin-right: 4px;
}

.message-image img {
  max-width: 100%;
  border-radius: 12px;
  display: block;
}

.image-caption {
  margin: 8px 0 0;
  font-size: 14px;
}

.message-file {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 16px;
  background: white;
  border-radius: 12px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

.file-icon {
  width: 40px;
  height: 40px;
  background: #3b82f6;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 20px;
}

.file-info {
  flex: 1;
}

.file-name {
  font-size: 14px;
  font-weight: 500;
  margin: 0 0 4px;
}

.file-size {
  font-size: 12px;
  color: #6b7280;
  margin: 0;
}

.file-download {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  border: none;
  background: #f3f4f6;
  color: #3b82f6;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
}

.message-voice {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 16px;
  background: white;
  border-radius: 12px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

.voice-play {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  border: none;
  background: #3b82f6;
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
}

.voice-wave {
  flex: 1;
  height: 24px;
  background: linear-gradient(90deg, #3b82f6 0%, #93c5fd 100%);
  border-radius: 12px;
}

.voice-duration {
  font-size: 12px;
  color: #6b7280;
}

.typing-indicator {
  display: flex;
  gap: 4px;
  padding: 12px 16px;
  background: white;
  border-radius: 18px;
  width: fit-content;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.05);
}

.typing-indicator span {
  width: 8px;
  height: 8px;
  background: #9ca3af;
  border-radius: 50%;
  animation: typing 1.4s infinite;
}

.typing-indicator span:nth-child(2) {
  animation-delay: 0.2s;
}

.typing-indicator span:nth-child(3) {
  animation-delay: 0.4s;
}

@keyframes typing {
  0%, 60%, 100% {
    transform: translateY(0);
  }
  30% {
    transform: translateY(-10px);
  }
}

.input-area {
  padding: 16px 24px;
  border-top: 1px solid #e5e7eb;
  display: flex;
  align-items: flex-end;
  gap: 12px;
  background: white;
  position: relative;
}

.file-menu {
  position: absolute;
  bottom: 70px;
  right: 24px;
  background: white;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  padding: 8px;
  display: flex;
  flex-direction: column;
  gap: 4px;
  z-index: 10;
}

.file-menu button {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 16px;
  border: none;
  background: none;
  cursor: pointer;
  border-radius: 8px;
  font-size: 14px;
  text-align: right;
  transition: all 0.2s;
  color: #374151;
}

.file-menu button:hover {
  background: #f3f4f6;
}

.file-menu button i {
  font-size: 20px;
}

.input-btn {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  border: none;
  background: #f3f4f6;
  color: #6b7280;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.2s;
  font-size: 20px;
  flex-shrink: 0;
}

.input-btn:hover {
  background: #e5e7eb;
  color: #374151;
}

.input-wrapper {
  flex: 1;
  background: #f3f4f6;
  border-radius: 24px;
  padding: 8px 16px;
}

.message-input {
  width: 100%;
  border: none;
  background: none;
  outline: none;
  font-size: 14px;
  resize: none;
  max-height: 120px;
  font-family: inherit;
}

.send-btn {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  border: none;
  background: #3b82f6;
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.2s;
  font-size: 20px;
  flex-shrink: 0;
}

.send-btn:hover {
  background: #2563eb;
  transform: scale(1.05);
}

.voice-btn:active {
  background: #dc2626;
  color: white;
}

.emoji-picker {
  position: absolute;
  bottom: 70px;
  left: 80px;
  width: 320px;
  height: 300px;
  background: white;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  padding: 12px;
  z-index: 10;
  overflow: hidden;
}

.emoji-grid {
  display: grid;
  grid-template-columns: repeat(8, 1fr);
  gap: 4px;
  height: 100%;
  overflow-y: auto;
}

.emoji-btn {
  width: 100%;
  aspect-ratio: 1;
  border: none;
  background: none;
  font-size: 24px;
  cursor: pointer;
  border-radius: 6px;
  transition: all 0.2s;
  display: flex;
  align-items: center;
  justify-content: center;
}

.emoji-btn:hover {
  background: #f3f4f6;
  transform: scale(1.1);
}

.chat-empty-state {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  color: #9ca3af;
}

.chat-empty-state i {
  font-size: 80px;
  margin-bottom: 24px;
  opacity: 0.5;
}

.chat-empty-state h3 {
  font-size: 24px;
  font-weight: 600;
  margin: 0 0 8px;
  color: #6b7280;
}

.chat-empty-state p {
  font-size: 14px;
  margin: 0;
}

.hidden {
  display: none;
}

/* Scrollbar */
.chat-list::-webkit-scrollbar,
.messages-area::-webkit-scrollbar,
.emoji-grid::-webkit-scrollbar {
  width: 6px;
}

.chat-list::-webkit-scrollbar-track,
.messages-area::-webkit-scrollbar-track,
.emoji-grid::-webkit-scrollbar-track {
  background: transparent;
}

.chat-list::-webkit-scrollbar-thumb,
.messages-area::-webkit-scrollbar-thumb,
.emoji-grid::-webkit-scrollbar-thumb {
  background: #d1d5db;
  border-radius: 3px;
}

.chat-list::-webkit-scrollbar-thumb:hover,
.messages-area::-webkit-scrollbar-thumb:hover,
.emoji-grid::-webkit-scrollbar-thumb:hover {
  background: #9ca3af;
}

/* Responsive */
@media (max-width: 768px) {
  .chat-sidebar {
    width: 100%;
    display: none;
  }

  .chat-sidebar.active {
    display: flex;
  }

  .chat-main {
    width: 100%;
  }

  .message-content {
    max-width: 80%;
  }
}
</style>
