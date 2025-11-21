<template>
  <div class="header-image-widget" :style="widgetStyle">
    <img 
      v-if="finalImageUrl"
      :src="finalImageUrl" 
      :alt="finalAltText"
      class="header-image"
      @error="handleImageError"
    />
    <div v-else class="image-placeholder">
      <div class="placeholder-icon">📷</div>
      <div class="placeholder-text">عکس انتخاب نشده</div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface Props {
  imageUrl?: string
  altText?: string
  borderRadius?: number
  objectFit?: 'cover' | 'contain' | 'fill' | 'scale-down'
  paddingRight?: number
  paddingLeft?: number
  // props اضافی که از آیتم هدر می‌آیند
  imageId?: number
  imageName?: string
}

const props = withDefaults(defineProps<Props>(), {
  imageUrl: '',
  altText: 'عکس هدر',
  borderRadius: 8,
  objectFit: 'cover',
  paddingRight: 0,
  paddingLeft: 0,
  imageId: undefined,
  imageName: ''
})

// استفاده از imageName اگر imageUrl موجود نباشد
const finalImageUrl = computed(() => {
  return props.imageUrl || ''
})

// استفاده از imageName برای alt text
const finalAltText = computed(() => {
  return props.imageName || props.altText || 'عکس هدر'
})

// استایل کامپوننت
const widgetStyle = computed(() => {
  const styles: Record<string, string> = {
    width: '100%',
    height: '100%'
  }
  
  // اگر padding تنظیم شده، اضافه کن
  if (props.paddingRight > 0) {
    styles.paddingRight = `${props.paddingRight}px`
  }
  if (props.paddingLeft > 0) {
    styles.paddingLeft = `${props.paddingLeft}px`
  }
  
  return styles
})

// مدیریت خطای بارگذاری عکس
const handleImageError = (_event: Event) => {
  // می‌توانید اینجا منطق جایگزینی عکس را اضافه کنید
}
</script>

<style scoped>
.header-image-widget {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: stretch;
}

.header-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  object-position: center;
  display: block;
}

.image-placeholder {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  background-color: #f5f5f5;
  border: 2px dashed #ddd;
  color: #666;
  font-size: 14px;
}

.placeholder-icon {
  font-size: 24px;
  margin-bottom: 8px;
}

.placeholder-text {
  font-size: 12px;
  text-align: center;
}
</style>
