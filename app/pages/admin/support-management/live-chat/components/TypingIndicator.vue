<template>
  <div v-if="isTyping" class="typing-indicator">
    <div class="typing-avatar">{{ userName[0] }}</div>
    <div class="typing-bubble">
      <div class="typing-dots">
        <div class="dot"></div>
        <div class="dot"></div>
        <div class="dot"></div>
      </div>
      <span class="typing-text">{{ userName }} در حال تایپ...</span>
    </div>
  </div>
</template>

<script setup>
defineProps({
  isTyping: {
    type: Boolean,
    default: false
  },
  userName: {
    type: String,
    default: 'کاربر'
  }
})
</script>

<style scoped>
.typing-indicator {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 8px 16px;
  margin-bottom: 8px;
  animation: fadeIn 0.3s ease-out;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.typing-avatar {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.3), rgba(118, 75, 162, 0.3));
  display: flex;
  align-items: center;
  justify-content: center;
  color: #667eea;
  font-weight: bold;
  font-size: 12px;
}

.typing-bubble {
  background: rgba(102, 126, 234, 0.1);
  border: 1px solid rgba(102, 126, 234, 0.2);
  border-radius: 16px;
  padding: 8px 12px;
  position: relative;
}

.typing-dots {
  display: flex;
  gap: 3px;
  margin-bottom: 4px;
}

.dot {
  width: 6px;
  height: 6px;
  border-radius: 50%;
  background: #667eea;
  animation: typingPulse 1.4s infinite ease-in-out;
}

.dot:nth-child(1) {
  animation-delay: 0s;
}

.dot:nth-child(2) {
  animation-delay: 0.2s;
}

.dot:nth-child(3) {
  animation-delay: 0.4s;
}

@keyframes typingPulse {
  0%, 60%, 100% {
    transform: scale(1);
    opacity: 0.5;
  }
  30% {
    transform: scale(1.2);
    opacity: 1;
  }
}

.typing-text {
  font-size: 11px;
  color: #667eea;
  font-style: italic;
}
</style> 
