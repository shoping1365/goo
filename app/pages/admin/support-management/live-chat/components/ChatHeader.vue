<template>
  <div 
    class="chat-header"
    @mousedown="$emit('startDrag', $event)"
  >
    <div class="flex justify-between items-center">
      <div class="flex items-center gap-3">
        <div class="header-icon">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
            <path stroke-linecap="round" stroke-linejoin="round" d="M8.625 9.75a.375.375 0 11-.75 0 .375.375 0 01.75 0zm0 0H8.25m4.125 0a.375.375 0 11-.75 0 .375.375 0 01.75 0zm0 0H12m4.125 0a.375.375 0 11-.75 0 .375.375 0 01.75 0zm0 0h-.375m5.618-1.5h2.207a1.5 1.5 0 010 3h-2.207a1.5 1.5 0 01-1.342-.83l-1.158-2.315a1.5 1.5 0 000-1.34l1.158-2.315A1.5 1.5 0 0119.902 6z" />
          </svg>
        </div>
        <div>
          <h3 class="text-white font-bold text-lg">چت پیشرفته ادمین</h3>
          <p class="text-white text-xs opacity-80">سیستم مدیریت چت هوشمند</p>
        </div>
      </div>
      <div class="flex items-center gap-3">
        <div class="waiting-counter">
          <span class="counter-number">{{ waitingCount }}</span>
          <span class="counter-label">منتظر</span>
        </div>
        <button @click="$emit('minimize')" class="text-white hover:text-gray-300" title="کوچک کردن">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-4 h-4">
            <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 12h-15" />
          </svg>
        </button>
        <button @click="$emit('close')" class="text-white hover:text-gray-300" title="بستن">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-4 h-4">
            <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  waitingCount: {
    type: Number,
    default: 0
  }
})

defineEmits(['startDrag', 'minimize', 'close'])
</script>

<style scoped>
.chat-header {
  background: linear-gradient(45deg, #667eea, #764ba2);
  padding: 16px 20px;
  cursor: move;
  user-select: none;
  border-bottom: none;
  backdrop-filter: blur(20px);
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
}

.header-icon {
  width: 44px;
  height: 44px;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  backdrop-filter: blur(10px);
  border: 2px solid rgba(255, 255, 255, 0.3);
}

.waiting-counter {
  background: rgba(255, 255, 255, 0.2);
  padding: 8px 12px;
  border-radius: 12px;
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.3);
  text-align: center;
  min-width: 60px;
}

.counter-number {
  display: block;
  color: white;
  font-size: 18px;
  font-weight: bold;
  line-height: 1;
}

.counter-label {
  display: block;
  color: rgba(255, 255, 255, 0.8);
  font-size: 10px;
  margin-top: 2px;
}
</style> 
