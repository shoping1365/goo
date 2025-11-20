<template>
  <ClientOnly>
    <!-- Chat Button (when closed) -->
    <button v-if="!isVisible" class="chat-button" @click="openWidget">
      <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
      </svg>
      <div v-if="unreadCount > 0" class="absolute -top-2 -right-2 w-6 h-6 bg-red-500 text-white text-xs rounded-full flex items-center justify-center font-bold animate-pulse">
        {{ unreadCount > 99 ? '99+' : unreadCount }}
      </div>
    </button>

    <!-- Chat Widget (when open) -->
    <div v-else class="chat-widget">
      <div class="flex items-center justify-between p-6 border-b border-gray-200/50 bg-gradient-to-r from-blue-500 to-blue-600 text-white rounded-t-2xl">
        <div class="flex items-center gap-3">
          <div class="relative">
            <div class="w-8 h-8 bg-white/20 rounded-full flex items-center justify-center">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
              </svg>
            </div>
            <!-- Connection Status Indicator -->
            <div class="absolute -bottom-1 -right-1 w-3 h-3 rounded-full border-2 border-white" :class="connectionStatus === 'connected' ? 'bg-green-400' : connectionStatus === 'connecting' ? 'bg-yellow-400' : 'bg-red-400'"></div>
          </div>
          <div>
            <h3 class="text-sm font-semibold">{{ isAdmin ? 'Admin Chat' : 'Online Support' }}</h3>
            <p v-if="isAdmin" class="text-xs opacity-80">{{ onlineAgents }} agents online</p>
            <p v-else class="text-xs opacity-80">{{ connectionStatus === 'connected' ? 'متصل' : connectionStatus === 'connecting' ? 'در حال اتصال...' : 'قطع شده' }}</p>
          </div>
        </div>
        <button class="w-8 h-8 hover:bg-red-500/20 rounded-lg flex items-center justify-center transition-colors" @click.stop="closeWidget">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
        </button>
      </div>
      
      <div class="flex flex-col h-[calc(100%-4rem)]">
        <div ref="messagesContainer" class="flex-1 overflow-y-auto p-6 space-y-3">
          <div v-if="messages.length === 0 && connectionStatus === 'connected'" class="text-center py-8">
            <div class="w-16 h-16 bg-blue-100 rounded-full flex items-center justify-center mx-auto mb-4">
              <svg class="w-8 h-8 text-blue-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
              </svg>
            </div>
            <h4 class="text-lg font-semibold text-gray-900 mb-2">{{ isAdmin ? 'Admin Chat' : 'سلام! 👋' }}</h4>
            <p class="text-sm text-gray-600 mb-4">{{ isAdmin ? 'چت با کاربران' : 'چطور می‌تونیم کمکتون کنیم؟' }}</p>
          </div>
          
          <!-- Connection Status Messages -->
          <div v-if="connectionStatus === 'connecting'" class="text-center py-4">
            <div class="inline-flex items-center gap-2 text-gray-500">
              <div class="w-4 h-4 border-2 border-blue-500 border-t-transparent rounded-full animate-spin"></div>
              <span>در حال اتصال به سرور چت...</span>
            </div>
          </div>
          
          <div v-if="connectionStatus === 'disconnected'" class="text-center py-4">
            <div class="inline-flex items-center gap-2 text-red-500">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
              <span>اتصال قطع شده. تلاش مجدد...</span>
            </div>
            <p v-if="lastWsError" class="text-xs text-gray-500 mt-2">{{ lastWsError }}</p>
          </div>
          
          <!-- Rate Limit Warning -->
          <div v-if="rateLimitExceeded" class="text-center py-2">
            <div class="inline-flex items-center gap-2 text-orange-600 bg-orange-50 px-3 py-2 rounded-lg">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
              </svg>
              <span>نرخ ارسال پیام بیش از حد مجاز است. لطفاً کمی صبر کنید.</span>
            </div>
          </div>
          
          <!-- Security Warning -->
          <div v-if="securityWarning" class="text-center py-2">
            <div class="inline-flex items-center gap-2 px-3 py-2 rounded-lg" :class="isBlocked ? 'text-red-700 bg-red-100' : 'text-amber-700 bg-amber-100'">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"></path>
              </svg>
              <span>{{ securityWarning }}</span>
            </div>
          </div>
          
          <!-- Blocked Warning -->
          <div v-if="isBlocked" class="text-center py-4">
            <div class="inline-flex flex-col items-center gap-3 text-red-700 bg-red-50 px-4 py-3 rounded-lg">
              <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728L5.636 5.636m12.728 12.728L18.364 5.636M5.636 18.364l12.728-12.728"></path>
              </svg>
              <div class="text-center">
                <p class="font-medium">دسترسی محدود شده</p>
                <p class="text-sm">به دلیل ارسال محتوای مخرب، دسترسی شما موقتاً محدود شده است.</p>
              </div>
            </div>
          </div>
          
          <div v-for="message in messages" :key="message.id" class="flex" :class="message.sender === 'user' ? 'justify-end' : 'justify-start'">
            <div class="max-w-[80%]">
              <div :class="[message.sender === 'user' ? 'bg-blue-500 text-white rounded-2xl rounded-br-md' : 'bg-gray-100 text-gray-900 rounded-2xl rounded-bl-md','px-4 py-3 shadow-sm']">
                <div class="text-sm">{{ message.content }}</div>
                <div class="flex items-center justify-between mt-2">
                  <span class="text-xs opacity-70">{{ formatTime(message.timestamp) }}</span>
                  <span v-if="message.status" class="text-xs opacity-70">{{ message.status }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
        
        <div class="p-6 border-t border-gray-200/50 bg-gray-50/50">
          <div class="flex items-end gap-3">
            <div class="flex-1 relative">
              <textarea 
                v-model="newMessage" 
                ref="messageInput"
                placeholder="پیام خود را بنویسید..."
                class="w-full px-4 py-3 border rounded-2xl focus:ring-2 focus:ring-blue-500 focus:border-transparent resize-none bg-white" 
                :class="securityWarning ? 'border-red-300' : 'border-gray-300'" 
                rows="1"
                :disabled="connectionStatus !== 'connected' || rateLimitExceeded || isBlocked || securityCooldown" 
                maxlength="1000"
                @keydown.enter.prevent="sendMessage"
                @input="validateInput"
              ></textarea>
              <!-- Character count and security indicator -->
              <div class="flex justify-between items-center mt-1 text-xs">
                <span v-if="newMessage.length > 800" :class="newMessage.length > 1000 ? 'text-red-500' : 'text-orange-500'">
                  {{ newMessage.length }}/1000
                </span>
                <span v-if="securityCooldown" class="text-orange-600">
                  🛡️ در حالت امنیتی
                </span>
              </div>
            </div>
            <button 
              :disabled="!newMessage.trim() || connectionStatus !== 'connected' || rateLimitExceeded" 
              class="px-4 py-3 bg-blue-500 text-white rounded-2xl hover:bg-blue-600 disabled:opacity-50 disabled:cursor-not-allowed transition-all duration-200 shadow-lg hover:shadow-xl flex-shrink-0" 
              @click="sendMessage"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8"></path>
              </svg>
            </button>
          </div>
        </div>
      </div>
    </div>
  </ClientOnly>
</template>

<script setup lang="ts">
import { useRuntimeConfig } from 'nuxt/app'
import { computed, nextTick, onMounted, onUnmounted, ref, watch } from 'vue'
import { useRoute } from 'vue-router'
import { generateSecurityReport, logSecurityEvent, validateTextMessage } from '~/utils/messageValidator'
import { parseWsClose, parseWsError } from '~/utils/wsErrors'

const route = useRoute()
const user = ref(null)
const isAuthenticated = computed(() => false)

const isAdmin = computed(() => route.path.startsWith('/admin'))

const isVisible = ref(false)
const newMessage = ref('')
const onlineAgents = ref(0)
const unreadCount = ref(0)
const messages = ref<any[]>([])
const messageInput = ref<HTMLTextAreaElement>()
const messagesContainer = ref<HTMLElement>()

// WebSocket connection state
const ws = ref<WebSocket | null>(null)
const connectionStatus = ref<'disconnected' | 'connecting' | 'connected'>('disconnected')
const rateLimitExceeded = ref(false)
const reconnectAttempts = ref(0)
const maxReconnectAttempts = 5
const reconnectInterval = ref<NodeJS.Timeout | null>(null)
const lastWsError = ref('')

// امنیت و اعتبارسنجی
const securityWarning = ref('')
const isBlocked = ref(false)
const securityCooldown = ref(false)

// Get WebSocket URL from environment
const getWebSocketUrl = () => {
  const config = useRuntimeConfig()
  const wsBase = config.public.wsBase as string | undefined
  
  // If wsBase is not provided, use current host
  if (!wsBase) {
    const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:'
    return `${protocol}//${window.location.host}/api/chat/ws`
  }
  
  // If wsBase is already a full URL, use it
  if (typeof wsBase === 'string' && (wsBase.startsWith('ws://') || wsBase.startsWith('wss://'))) {
    return `${wsBase}/api/chat/ws`
  }
  
  // Otherwise, construct from current protocol and host
  const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:'
  const host = wsBase || window.location.host
  return `${protocol}//${host}/api/chat/ws`
}

// Get authentication token
const getAuthToken = () => {
  // کوکی HttpOnly به صورت خودکار توسط مرورگر ارسال می‌شود، نیازی به توکن سمت کلاینت نیست
  // اگر برای WebSocket نیاز به پارامتر داریم، آن را از کوکی refresh/session استخراج نمی‌کنیم
  return undefined
}

  // Connect to WebSocket
  const connectWebSocket = () => {
  if (ws.value?.readyState === WebSocket.OPEN) {
    return
  }

  connectionStatus.value = 'connecting'
  lastWsError.value = ''
  
    try {
    const wsUrl = getWebSocketUrl()
    const url = wsUrl
    
    ws.value = new WebSocket(url)
    
    ws.value.onopen = () => {
      console.log('WebSocket connected')
      connectionStatus.value = 'connected'
      reconnectAttempts.value = 0
      rateLimitExceeded.value = false
      lastWsError.value = ''
    }
    
      ws.value.onmessage = (event) => {
      try {
        const data = JSON.parse(event.data)
        handleWebSocketMessage(data)
      } catch (error) {
        console.error('Error parsing WebSocket message:', error)
      }
    }
    
    ws.value.onclose = (event) => {
      console.log('WebSocket disconnected:', event.code, event.reason)
      connectionStatus.value = 'disconnected'
      const info = parseWsClose(event)
      lastWsError.value = `${info.userMessage}`
      
      // Auto-reconnect logic
      if (reconnectAttempts.value < maxReconnectAttempts && !event.wasClean) {
        reconnectAttempts.value++
        const delay = Math.min(1000 * Math.pow(2, reconnectAttempts.value), 30000) // Exponential backoff, max 30s
        
        reconnectInterval.value = setTimeout(() => {
          console.log(`Attempting to reconnect (${reconnectAttempts.value}/${maxReconnectAttempts})`)
          connectWebSocket()
        }, delay)
      }
    }
    
    ws.value.onerror = (error) => {
      console.error('WebSocket error:', error)
      connectionStatus.value = 'disconnected'
      lastWsError.value = parseWsError(error as any)
    }
    
  } catch (error) {
    console.error('Error creating WebSocket connection:', error)
    connectionStatus.value = 'disconnected'
    lastWsError.value = parseWsError(error as any)
  }
}

// Handle incoming WebSocket messages
  const handleWebSocketMessage = (data: any) => {
  switch (data.type) {
      case 'message':
      const message = {
        id: Date.now() + Math.random(),
        sender: data.sender === 'user' ? 'user' : 'agent',
        content: data.content,
        timestamp: new Date(data.timestamp || Date.now()),
        status: 'delivered'
      }
      messages.value.push(message)
      
      // Increment unread count if chat is closed
      if (!isVisible.value && message.sender === 'agent') {
        unreadCount.value++
      }
      
      // Scroll to bottom
      nextTick(() => {
        scrollToBottom()
      })
      break
      
      case 'error':
      if (data.message === 'rate limit exceeded') {
        rateLimitExceeded.value = true
        setTimeout(() => {
          rateLimitExceeded.value = false
        }, 5000) // Show warning for 5 seconds
      }
      break
      
      case 'new_message':
        // فرمت جدید بک‌اند
        if (data.data) {
          const msg = data.data
          const mapped = {
            id: msg.id || Date.now() + Math.random(),
            sender: msg.sender_type === 'customer' ? 'user' : 'agent',
            content: msg.message,
            timestamp: new Date(msg.created_at || Date.now()),
            status: 'delivered'
          }
          messages.value.push(mapped)
          if (!isVisible.value && mapped.sender === 'agent') {
            unreadCount.value++
          }
          nextTick(() => {
            scrollToBottom()
          })
        }
        break
      
    case 'stats':
      if (data.onlineAgents !== undefined) {
        onlineAgents.value = data.onlineAgents
      }
      break
      
    case 'ping':
      // Respond to ping with pong
      if (ws.value?.readyState === WebSocket.OPEN) {
        ws.value.send(JSON.stringify({ type: 'pong' }))
      }
      break
  }
}

// Send message via WebSocket
  const sendMessage = () => {
  if (!newMessage.value.trim() || connectionStatus.value !== 'connected' || rateLimitExceeded.value || isBlocked.value || securityCooldown.value) {
    return
  }
  
  const messageContent = newMessage.value.trim()
  
  // اعتبارسنجی امنیتی پیام
  const validationResult = validateTextMessage(messageContent)
  
  if (!validationResult.isValid) {
    // نمایش هشدار امنیتی
    securityWarning.value = validationResult.errors[0]
    
    // لاگ امنیتی
    const securityReport = generateSecurityReport(validationResult, user.value)
    logSecurityEvent(securityReport)
    
    // در صورت تهدید بالا، کاربر را مسدود کن
    if (validationResult.riskLevel === 'critical') {
      isBlocked.value = true
      securityWarning.value = 'پیام شما حاوی محتوای مخرب است. دسترسی شما محدود شده است.'
      
      // مسدود کردن موقت (10 دقیقه)
      setTimeout(() => {
        isBlocked.value = false
        securityWarning.value = ''
      }, 10 * 60 * 1000)
      
      newMessage.value = ''
      return
    }
    
    // در صورت تهدید متوسط، cooldown کوتاه
    if (validationResult.riskLevel === 'high') {
      securityCooldown.value = true
      setTimeout(() => {
        securityCooldown.value = false
        securityWarning.value = ''
      }, 30000) // 30 ثانیه
      
      newMessage.value = ''
      return
    }
    
    // پاک کردن هشدار بعد از 5 ثانیه
    setTimeout(() => {
      securityWarning.value = ''
    }, 5000)
    
    newMessage.value = ''
    return
  }
  
  // استفاده از محتوای پاکسازی شده
  const sanitizedContent = validationResult.sanitizedContent || messageContent
  newMessage.value = ''
  
  // بررسی طول نهایی
  if (sanitizedContent.length > 1000) {
    securityWarning.value = 'پیام بیش از حد طولانی است (حداکثر 1000 کاراکتر)'
    setTimeout(() => {
      securityWarning.value = ''
    }, 5000)
    return
  }
  
  // Add message to local state immediately
  const message = {
    id: Date.now() + Math.random(),
    sender: 'user' as const,
    content: sanitizedContent,
    timestamp: new Date(),
    status: 'sending'
  }
  messages.value.push(message)
  
  // Scroll to bottom
  nextTick(() => {
    scrollToBottom()
  })
  
    // Send via WebSocket (فرمت جدید)
    if (ws.value?.readyState === WebSocket.OPEN) {
    try {
        ws.value.send(JSON.stringify({ type: 'send_message', data: { content: sanitizedContent } }))
      
      // Update message status to sent
      message.status = 'sent'
    } catch (error) {
      console.error('Error sending message:', error)
      message.status = 'failed'
    }
  } else {
    message.status = 'failed'
  }
}

// اعتبارسنجی ورودی در زمان واقعی
const validateInput = () => {
  if (!newMessage.value) {
    securityWarning.value = ''
    return
  }
  
  // بررسی سریع الگوهای مخرب در حین تایپ
  const content = newMessage.value
  
  // بررسی script tags
  if (/<script/gi.test(content)) {
    securityWarning.value = 'استفاده از تگ script مجاز نیست'
    return
  }
  
  // بررسی javascript: URLs
  if (/javascript:/gi.test(content)) {
    securityWarning.value = 'استفاده از لینک‌های JavaScript مجاز نیست'
    return
  }
  
  // بررسی HTML injection
  if (/<iframe|<object|<embed/gi.test(content)) {
    securityWarning.value = 'استفاده از تگ‌های HTML خاص مجاز نیست'
    return
  }
  
  // پاک کردن هشدار اگر محتوا امن است
  securityWarning.value = ''
}

// Scroll messages container to bottom
const scrollToBottom = () => {
  if (messagesContainer.value) {
    messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight
  }
}

const openWidget = () => {
  isVisible.value = true
  unreadCount.value = 0
  
  // Connect to WebSocket when opening chat
  if (connectionStatus.value === 'disconnected') {
    connectWebSocket()
  }
  
  nextTick(() => {
    messageInput.value?.focus()
    scrollToBottom()
  })
}

const closeWidget = () => {
  isVisible.value = false
}

const formatTime = (timestamp: Date) => {
  return new Date(timestamp).toLocaleTimeString('fa-IR', { hour: '2-digit', minute: '2-digit' })
}

// Watch for authentication changes
watch(isAuthenticated, (newValue) => {
  if (newValue && isVisible.value) {
    // Reconnect with new auth token
    if (ws.value) {
      ws.value.close()
    }
    connectWebSocket()
  }
})

onMounted(async () => {
  // Show unread indicator after 5 seconds if chat is not open
  setTimeout(() => {
    if (!isVisible.value && connectionStatus.value === 'connected') {
      unreadCount.value = 1
    }
  }, 5000)
})

onUnmounted(() => {
  // Clean up WebSocket connection
  if (ws.value) {
    ws.value.close()
  }
  
  if (reconnectInterval.value) {
    clearTimeout(reconnectInterval.value)
  }
})
</script>

<style scoped>
.chat-widget {
  position: fixed !important;
  bottom: 24px !important;
  right: 24px !important;
  left: auto !important;
  top: auto !important;
  width: 400px !important;
  height: 600px !important;
  max-width: calc(100vw - 48px) !important;
  max-height: calc(100vh - 48px) !important;
  background: white !important;
  border-radius: 1rem !important;
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25) !important;
  border: 1px solid rgba(229, 231, 235, 0.5) !important;
  z-index: 9999 !important;
  direction: ltr !important;
  transform: none !important;
  margin: 0 !important;
  padding: 0 !important;
}

.chat-button {
  position: fixed !important;
  bottom: 24px !important;
  right: 24px !important;
  left: auto !important;
  top: auto !important;
  width: 64px !important;
  height: 64px !important;
  background: linear-gradient(135deg, #3b82f6, #1d4ed8) !important;
  color: white !important;
  border-radius: 50% !important;
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25) !important;
  border: none !important;
  cursor: pointer !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  z-index: 9998 !important;
  transition: all 0.3s ease !important;
  transform: none !important;
  margin: 0 !important;
  padding: 0 !important;
}

.chat-button:hover {
  transform: scale(1.1) !important;
  box-shadow: 0 32px 64px -12px rgba(0, 0, 0, 0.35) !important;
}

.chat-button, .chat-widget {
  font-family: inherit !important;
}
</style> 
