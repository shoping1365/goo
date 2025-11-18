<template>
  <div v-if="notification" :class="['notification', notification.type]">
    {{ notification.message }}
  </div>
</template>

<script setup>
defineProps({
  notification: {
    type: Object,
    default: null
  }
})
</script>

<style scoped>
.notification {
  position: fixed;
  bottom: 20px;
  right: 20px;
  padding: 12px 20px;
  border-radius: 16px;
  color: white;
  font-size: 13px;
  font-weight: 500;
  z-index: 1000;
  backdrop-filter: blur(20px);
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.2);
  animation: notificationSlideIn 0.3s ease-out;
}

@keyframes notificationSlideIn {
  from {
    opacity: 0;
    transform: translateX(100px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
}

.notification.success {
  background: linear-gradient(135deg, #4ade80, #22c55e);
}

.notification.info {
  background: linear-gradient(135deg, #667eea, #764ba2);
}

.notification.warning {
  background: linear-gradient(135deg, #ffa726, #ffcc02);
}

.notification.error {
  background: linear-gradient(135deg, #ff6b6b, #ff8787);
}
</style> 
