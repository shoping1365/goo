<template>
  <div class="chat-area">
    <div v-if="selectedUser" class="chat-container">
      
      <!-- Chat Header -->
      <div class="chat-user-header">
        <div class="chat-user-info-main">
          <div class="user-avatar-large">{{ selectedUser.name[0] }}</div>
          <div class="user-info-text">
            <h4>{{ selectedUser.name }}</h4>
            <p>{{ selectedUser.location }} - {{ selectedUser.device }} - {{ selectedUser.browser }}</p>
          </div>
        </div>
        <div class="chat-actions">
          <button class="action-btn end-chat" @click="$emit('endChat')">پایان</button>
          <button class="action-btn transfer" @click="$emit('transferChat')">انتقال</button>
        </div>
      </div>

      <!-- Messages Area -->
      <div class="messages-area" ref="messagesContainer">
        <div v-for="message in messages" :key="message.id" :class="['message', message.sender === 'admin' ? 'outgoing' : 'incoming']">
          <div class="message-bubble">
            <p v-if="message.type === 'text'">{{ message.text }}</p>
            <div v-else-if="message.type === 'voice'" class="voice-message">
              <audio :src="message.audioUrl" controls class="voice-player"></audio>
              <span class="voice-duration">{{ formatDuration(message.duration) }}</span>
            </div>
            <div v-else-if="message.type === 'file'" class="file-message">
              <div class="file-icon">{{ getFileIcon(message.fileType) }}</div>
              <div class="file-details">
                <span class="file-name">{{ message.fileName }}</span>
                <span class="file-size">{{ formatFileSize(message.fileSize) }}</span>
              </div>
              <a :href="message.fileUrl" download class="download-btn">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-4 h-4">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M3 16.5v2.25A2.25 2.25 0 005.25 21h13.5A2.25 2.25 0 0021 18.75V16.5M16.5 12L12 16.5m0 0L7.5 12m4.5 4.5V3" />
                </svg>
              </a>
            </div>
            <span class="message-time">{{ message.time }}</span>
          </div>
        </div>
        
        <!-- Typing Indicator -->
        <TypingIndicator 
          v-if="selectedUser"
          :is-typing="isUserTyping" 
          :user-name="selectedUser.name" 
        />
      </div>

      <!-- Quick Replies -->
      <QuickReplies @select-reply="handleQuickReply" />
      
      <!-- Chat Input -->
      <div class="chat-input-area">
        <div class="input-container">
          <button class="attachment-btn" @click="$emit('attachFile')">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-4 h-4">
              <path stroke-linecap="round" stroke-linejoin="round" d="m18.375 12.739-7.693 7.693a4.5 4.5 0 0 1-6.364-6.364l10.94-10.94A3 3 0 1 1 19.5 7.372L8.552 18.32m.009-.01-.01.01m5.699-9.941-7.81 7.81a1.5 1.5 0 0 0 2.112 2.13" />
            </svg>
          </button>
          <input 
            v-model="localNewMessage" 
            @keypress.enter="sendMessage"
            type="text" 
            placeholder="پیام خود را بنویسید..." 
            class="message-input"
          >
          <VoiceRecorder @voice-message="handleVoiceMessage" />
          <button @click="sendMessage" class="send-btn">ارسال</button>
        </div>
      </div>
    </div>
    
    <div v-else class="no-chat-selected">
      <div class="no-chat-icon">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-8 h-8">
          <path stroke-linecap="round" stroke-linejoin="round" d="M8.625 12a.375.375 0 11-.75 0 .375.375 0 01.75 0zm0 0H8.25m4.125 0a.375.375 0 11-.75 0 .375.375 0 01.75 0zm0 0H12m4.125 0a.375.375 0 11-.75 0 .375.375 0 01.75 0zm0 0h-.375M21 12c0 4.556-4.03 8.25-9 8.25a9.764 9.764 0 01-2.555-.337A5.972 5.972 0 015.41 20.97a5.969 5.969 0 01-.474-.065 4.48 4.48 0 00.978-2.025c.09-.457-.133-.901-.467-1.226C3.93 16.178 3 14.189 3 12c0-4.556 4.03-8.25 9-8.25s9 3.694 9 8.25z" />
        </svg>
      </div>
      <h3>هیچ چتی انتخاب نشده</h3>
      <p>برای شروع گفتگو، یک کاربر از صف انتظار انتخاب کنید</p>
    </div>
  </div>
</template>

<script setup>
import TypingIndicator from './TypingIndicator.vue'
import QuickReplies from './QuickReplies.vue'
import VoiceRecorder from './VoiceRecorder.vue'

const props = defineProps({
  selectedUser: Object,
  messages: Array,
  newMessage: String
})

const emit = defineEmits(['endChat', 'transferChat', 'attachFile', 'sendMessage', 'update:newMessage', 'voiceMessage'])

const localNewMessage = ref(props.newMessage || '')
const messagesContainer = ref(null)
const isUserTyping = ref(false)

// Simulate user typing
onMounted(() => {
  if (props.selectedUser) {
    // Random typing simulation
    const typingTimer = setInterval(() => {
      if (Math.random() > 0.7) {
        isUserTyping.value = true
        setTimeout(() => {
          isUserTyping.value = false
        }, 2000 + Math.random() * 3000)
      }
    }, 8000)
    
    onUnmounted(() => {
      clearInterval(typingTimer)
    })
  }
})

// Watch for prop changes
watch(() => props.newMessage, (newVal) => {
  localNewMessage.value = newVal
})

// Auto scroll to bottom when new messages arrive
watch(() => props.messages, () => {
  nextTick(() => {
    if (messagesContainer.value) {
      messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight
    }
  })
}, { deep: true })

const sendMessage = () => {
  if (localNewMessage.value.trim()) {
    emit('sendMessage', localNewMessage.value.trim())
    localNewMessage.value = ''
    emit('update:newMessage', '')
  }
}

const handleQuickReply = (text) => {
  localNewMessage.value = text
  sendMessage()
}

const handleVoiceMessage = (voiceData) => {
  emit('voiceMessage', voiceData)
}

const formatDuration = (seconds) => {
  const mins = Math.floor(seconds / 60)
  const secs = seconds % 60
  return `${mins}:${secs.toString().padStart(2, '0')}`
}

const getFileIcon = (type) => {
  if (type?.startsWith('image/')) return '🖼️'
  if (type?.startsWith('video/')) return '🎥'
  if (type?.startsWith('audio/')) return '🎵'
  if (type?.includes('pdf')) return '📄'
  if (type?.includes('word') || type?.includes('doc')) return '📝'
  return '📎'
}

const formatFileSize = (bytes) => {
  if (bytes === 0) return '0 Bytes'
  const k = 1024
  const sizes = ['Bytes', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}
</script>

<style scoped>
.chat-area {
  flex: 1;
  display: flex;
  flex-direction: column;
  background: rgba(255, 255, 255, 0.98);
  border-radius: 0 0 12px 12px;
  backdrop-filter: blur(20px);
  margin: 0 8px 8px 8px;
}

.chat-container {
  height: 100%;
  display: flex;
  flex-direction: column;
}

.chat-user-header {
  padding: 16px 20px;
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.1), rgba(118, 75, 162, 0.1));
  border-bottom: 1px solid rgba(255, 255, 255, 0.2);
  display: flex;
  justify-content: space-between;
  align-items: center;
  backdrop-filter: blur(10px);
  border-radius: 12px;
  margin: 8px;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.05);
}

.chat-user-info-main {
  display: flex;
  align-items: center;
  gap: 12px;
}

.user-avatar-large {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  background: linear-gradient(135deg, #667eea, #764ba2);
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-weight: bold;
  font-size: 18px;
  box-shadow: 0 8px 24px rgba(102, 126, 234, 0.3);
  border: 3px solid rgba(255, 255, 255, 0.3);
}

.user-info-text h4 {
  font-size: 16px;
  font-weight: 600;
  color: #2d3748;
  margin-bottom: 4px;
}

.user-info-text p {
  font-size: 12px;
  color: #718096;
}

.chat-actions {
  display: flex;
  gap: 8px;
}

.action-btn {
  padding: 8px 12px;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  font-size: 12px;
  font-weight: 500;
  transition: all 0.3s ease;
}

.action-btn.end-chat {
  background: linear-gradient(135deg, #ff6b6b, #ff8787);
  color: white;
  box-shadow: 0 4px 16px rgba(255, 107, 107, 0.4);
}

.action-btn.transfer {
  background: linear-gradient(135deg, #ffa726, #ffcc02);
  color: white;
  box-shadow: 0 4px 16px rgba(255, 167, 38, 0.4);
}

.action-btn:hover {
  transform: translateY(-2px);
}

.messages-area {
  flex: 1;
  padding: 20px;
  overflow-y: auto;
  background: linear-gradient(135deg, rgba(248, 249, 250, 0.8), rgba(255, 255, 255, 0.9));
  display: flex;
  flex-direction: column;
  gap: 16px;
  margin: 8px;
  border-radius: 12px;
}

.message {
  display: flex;
}

.message.incoming {
  justify-content: flex-start;
}

.message.outgoing {
  justify-content: flex-end;
}

.message-bubble {
  max-width: 75%;
  padding: 12px 16px;
  border-radius: 20px;
  position: relative;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.1);
  backdrop-filter: blur(10px);
  animation: messageSlideIn 0.3s ease-out;
}

@keyframes messageSlideIn {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.message.incoming .message-bubble {
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.15), rgba(118, 75, 162, 0.15));
  color: #4c63d2;
  border: 1px solid rgba(102, 126, 234, 0.2);
}

.message.outgoing .message-bubble {
  background: linear-gradient(135deg, #667eea, #764ba2);
  color: white;
  box-shadow: 0 8px 24px rgba(102, 126, 234, 0.3);
}

.message-time {
  display: block;
  font-size: 10px;
  opacity: 0.7;
  margin-top: 4px;
}

.voice-message {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.voice-player {
  width: 200px;
  height: 30px;
  border-radius: 15px;
  outline: none;
  background: rgba(255, 255, 255, 0.2);
}

.voice-player::-webkit-media-controls-panel {
  background: transparent;
}

.voice-duration {
  font-size: 10px;
  opacity: 0.8;
  color: inherit;
}

.file-message {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 8px;
  background: rgba(255, 255, 255, 0.1);
  border-radius: 8px;
  min-width: 200px;
}

.file-icon {
  font-size: 24px;
  flex-shrink: 0;
}

.file-details {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.file-name {
  font-size: 13px;
  font-weight: 500;
  color: inherit;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.file-size {
  font-size: 10px;
  opacity: 0.7;
  color: inherit;
}

.download-btn {
  background: rgba(255, 255, 255, 0.2);
  border: none;
  border-radius: 50%;
  padding: 8px;
  cursor: pointer;
  color: inherit;
  transition: all 0.3s ease;
  text-decoration: none;
  display: flex;
  align-items: center;
  justify-content: center;
}

.download-btn:hover {
  background: rgba(255, 255, 255, 0.3);
  transform: scale(1.1);
}

.chat-input-area {
  padding: 20px;
  border-top: none;
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(20px);
  margin: 8px;
  border-radius: 12px;
  box-shadow: 0 -4px 16px rgba(0, 0, 0, 0.05);
}

.input-container {
  display: flex;
  gap: 12px;
  align-items: center;
  background: rgba(255, 255, 255, 0.8);
  padding: 8px;
  border-radius: 24px;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.1);
}

.attachment-btn {
  padding: 10px;
  border: none;
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.2), rgba(118, 75, 162, 0.2));
  border-radius: 50%;
  cursor: pointer;
  transition: all 0.3s ease;
  color: #667eea;
}

.attachment-btn:hover {
  transform: scale(1.1);
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3);
}

.message-input {
  flex: 1;
  padding: 12px 16px;
  border: none;
  border-radius: 20px;
  outline: none;
  font-size: 14px;
  background: transparent;
  color: #2d3748;
}

.message-input::placeholder {
  color: #a0aec0;
}

.send-btn {
  padding: 10px 16px;
  background: linear-gradient(135deg, #667eea, #764ba2);
  color: white;
  border: none;
  border-radius: 20px;
  cursor: pointer;
  transition: all 0.3s ease;
  box-shadow: 0 4px 16px rgba(102, 126, 234, 0.4);
}

.send-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 24px rgba(102, 126, 234, 0.5);
}

.no-chat-selected {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  text-align: center;
  color: #64748b;
  padding: 40px;
}

.no-chat-icon {
  opacity: 0.4;
  margin-bottom: 20px;
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.1), rgba(118, 75, 162, 0.1));
  width: 80px;
  height: 80px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #667eea;
}

.no-chat-selected h3 {
  color: #4a5568;
  font-size: 18px;
  font-weight: 600;
  margin-bottom: 8px;
}

.no-chat-selected p {
  color: #718096;
  font-size: 14px;
  line-height: 1.5;
}
</style> 
