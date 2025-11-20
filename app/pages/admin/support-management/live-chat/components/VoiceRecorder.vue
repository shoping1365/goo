<template>
  <div class="voice-recorder">
    <button 
      :class="['record-btn', { recording: isRecording }]"
      :disabled="!isSupported"
      @mousedown="startRecording"
      @mouseup="stopRecording"
      @mouseleave="stopRecording"
      @touchstart="startRecording"
      @touchend="stopRecording"
    >
      <svg v-if="!isRecording" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5">
        <path stroke-linecap="round" stroke-linejoin="round" d="M12 18.75a6 6 0 006-6v-1.5m-6 7.5a6 6 0 01-6-6v-1.5m6 7.5v3.75m-3.75 0h7.5M12 15.75a3 3 0 01-3-3V4.5a3 3 0 116 0v8.25a3 3 0 01-3 3z" />
      </svg>
      
      <div v-else class="recording-animation">
        <div class="pulse-ring"></div>
        <svg xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 24 24" class="w-5 h-5">
          <path d="M12 2a3 3 0 013 3v6a3 3 0 01-6 0V5a3 3 0 013-3z"/>
          <path d="M19 10v1a7 7 0 01-14 0v-1"/>
        </svg>
      </div>
    </button>
    
    <div v-if="isRecording" class="recording-info">
      <span class="recording-time">{{ formatTime(recordingTime) }}</span>
      <span class="recording-text">در حال ضبط...</span>
    </div>
  </div>
</template>

<script setup>
const isRecording = ref(false)
const recordingTime = ref(0)
const isSupported = ref(false)
let mediaRecorder = null
let recordingInterval = null

const emit = defineEmits(['voiceMessage'])

onMounted(() => {
  // Check if browser supports MediaRecorder
  isSupported.value = !!(navigator.mediaDevices && MediaRecorder)
})

const startRecording = async () => {
  if (!isSupported.value || isRecording.value) return
  
  try {
    const stream = await navigator.mediaDevices.getUserMedia({ audio: true })
    mediaRecorder = new MediaRecorder(stream)
    
    const chunks = []
    mediaRecorder.ondataavailable = (event) => {
      chunks.push(event.data)
    }
    
    mediaRecorder.onstop = () => {
      const blob = new Blob(chunks, { type: 'audio/webm' })
      const audioUrl = URL.createObjectURL(blob)
      
      emit('voiceMessage', {
        blob,
        url: audioUrl,
        duration: recordingTime.value
      })
      
      // Clean up
      stream.getTracks().forEach(track => track.stop())
    }
    
    mediaRecorder.start()
    isRecording.value = true
    recordingTime.value = 0
    
    // Start timer
    recordingInterval = setInterval(() => {
      recordingTime.value++
      
      // Auto stop after 60 seconds
      if (recordingTime.value >= 60) {
        stopRecording()
      }
    }, 1000)
    
  } catch (error) {
    console.error('Error accessing microphone:', error)
  }
}

const stopRecording = () => {
  if (!isRecording.value) return
  
  mediaRecorder?.stop()
  isRecording.value = false
  
  if (recordingInterval) {
    clearInterval(recordingInterval)
    recordingInterval = null
  }
}

const formatTime = (seconds) => {
  const mins = Math.floor(seconds / 60)
  const secs = seconds % 60
  return `${mins}:${secs.toString().padStart(2, '0')}`
}

onUnmounted(() => {
  if (recordingInterval) {
    clearInterval(recordingInterval)
  }
})
</script>

<style scoped>
.voice-recorder {
  display: flex;
  align-items: center;
  gap: 8px;
}

.record-btn {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  border: none;
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.2), rgba(118, 75, 162, 0.2));
  color: #667eea;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
}

.record-btn:hover:not(:disabled) {
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.3), rgba(118, 75, 162, 0.3));
  transform: scale(1.05);
}

.record-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.record-btn.recording {
  background: linear-gradient(135deg, #ff6b6b, #ff8787);
  color: white;
  animation: recordingPulse 1s infinite ease-in-out;
}

@keyframes recordingPulse {
  0%, 100% {
    transform: scale(1);
  }
  50% {
    transform: scale(1.1);
  }
}

.recording-animation {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
}

.pulse-ring {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 60px;
  height: 60px;
  border: 2px solid rgba(255, 255, 255, 0.3);
  border-radius: 50%;
  animation: pulsate 1.5s infinite;
}

@keyframes pulsate {
  0% {
    transform: translate(-50%, -50%) scale(0.5);
    opacity: 1;
  }
  100% {
    transform: translate(-50%, -50%) scale(1.2);
    opacity: 0;
  }
}

.recording-info {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 2px;
}

.recording-time {
  font-size: 14px;
  font-weight: bold;
  color: #ff6b6b;
  font-family: monospace;
}

.recording-text {
  font-size: 10px;
  color: #718096;
  animation: blink 1s infinite;
}

@keyframes blink {
  0%, 50% {
    opacity: 1;
  }
  51%, 100% {
    opacity: 0.5;
  }
}
</style> 
