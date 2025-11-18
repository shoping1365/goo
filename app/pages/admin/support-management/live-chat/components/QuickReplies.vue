<template>
  <div class="quick-replies">
    <div class="quick-replies-header">
      <h4>پاسخ‌های سریع</h4>
      <button @click="isExpanded = !isExpanded" class="expand-btn">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-4 h-4" :class="{ 'rotate-180': isExpanded }">
          <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 8.25l-7.5 7.5-7.5-7.5" />
        </svg>
      </button>
    </div>
    
    <div v-if="isExpanded" class="quick-replies-list">
      <button 
        v-for="reply in quickReplies" 
        :key="reply.id"
        @click="$emit('selectReply', reply.text)"
        class="quick-reply-btn"
      >
        <span class="reply-icon">{{ reply.icon }}</span>
        <span class="reply-text">{{ reply.text }}</span>
      </button>
    </div>
  </div>
</template>

<script setup>
const isExpanded = ref(true)

const quickReplies = ref([
  {
    id: 1,
    icon: '👋',
    text: 'سلام! چطور می‌تونم کمکتون کنم؟'
  },
  {
    id: 2,
    icon: '⏰',
    text: 'لطفاً چند دقیقه صبر کنید...'
  },
  {
    id: 3,
    icon: '✅',
    text: 'مشکل شما حل شد!'
  },
  {
    id: 4,
    icon: '📞',
    text: 'برای پیگیری با تلفن تماس بگیرید'
  },
  {
    id: 5,
    icon: '📧',
    text: 'ایمیل شما را برای بررسی ارسال کردم'
  },
  {
    id: 6,
    icon: '💳',
    text: 'مشکل پرداخت شما بررسی می‌شود'
  },
  {
    id: 7,
    icon: '🙏',
    text: 'متشکرم از صبر شما'
  },
  {
    id: 8,
    icon: '❓',
    text: 'سوال دیگری دارید؟'
  }
])

defineEmits(['selectReply'])
</script>

<style scoped>
.quick-replies {
  background: rgba(255, 255, 255, 0.9);
  border-radius: 12px;
  margin: 8px;
  backdrop-filter: blur(10px);
  border: 1px solid rgba(102, 126, 234, 0.1);
}

.quick-replies-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  border-bottom: 1px solid rgba(102, 126, 234, 0.1);
}

.quick-replies-header h4 {
  margin: 0;
  font-size: 14px;
  font-weight: 600;
  color: #4a5568;
}

.expand-btn {
  background: none;
  border: none;
  cursor: pointer;
  padding: 4px;
  border-radius: 50%;
  color: #667eea;
  transition: all 0.3s ease;
}

.expand-btn:hover {
  background: rgba(102, 126, 234, 0.1);
}

.expand-btn svg {
  transition: transform 0.3s ease;
}

.rotate-180 {
  transform: rotate(180deg);
}

.quick-replies-list {
  padding: 8px;
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 6px;
  max-height: 200px;
  overflow-y: auto;
}

.quick-reply-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 12px;
  background: rgba(102, 126, 234, 0.05);
  border: 1px solid rgba(102, 126, 234, 0.1);
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s ease;
  text-align: right;
  font-size: 12px;
}

.quick-reply-btn:hover {
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.1), rgba(118, 75, 162, 0.1));
  border-color: #667eea;
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.2);
}

.reply-icon {
  font-size: 14px;
  flex-shrink: 0;
}

.reply-text {
  color: #4a5568;
  line-height: 1.3;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
}
</style> 
