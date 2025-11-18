<template>
  <div class="footer-working-hours-widget" :style="widgetStyle">
    <div class="text-center">
      <!-- عنوان بخش -->
      <h4 v-if="title" class="footer-widget-title font-semibold mb-4 text-lg">{{ title }}</h4>
      
      <!-- ساعات کاری -->
      <div v-if="workingSchedules && workingSchedules.length > 0" class="working-hours-simple">
        <div 
          v-for="(schedule, index) in workingSchedules" 
          :key="index"
          class="schedule-item"
        >
          {{ schedule }}
        </div>
      </div>
      
      <!-- پیش‌فرض -->
      <div v-else class="schedule-item">
        ساعات کاری تعریف نشده است
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface Props {
  title?: string
  workingSchedules?: string[]
  paddingRight?: number
  paddingLeft?: number
  align?: 'left' | 'center' | 'right'
}

const props = withDefaults(defineProps<Props>(), {
  title: 'ساعات کاری',
  workingSchedules: () => [],
  paddingRight: 0,
  paddingLeft: 0,
  align: 'center'
})

// استایل کامپوننت بر اساس چینش
const widgetStyle = computed(() => ({
  paddingRight: `${props.paddingRight}px`,
  paddingLeft: `${props.paddingLeft}px`,
  display: 'flex',
  alignItems: 'center',
  justifyContent: getJustifyContent(props.align),
  width: '100%',
  height: '100%'
}))

// تابع تعیین justify-content بر اساس چینش
function getJustifyContent(align: string): string {
  switch (align) {
    case 'left':
      return 'flex-start'
    case 'center':
      return 'center'
    case 'right':
      return 'flex-end'
    default:
      return 'center'
  }
}
</script>

<style scoped>
.footer-working-hours-widget {
  transition: all 0.2s ease;
}

/* رنگ متن پیش‌فرض - تیره برای پس‌زمینه روشن */
.footer-widget-title,
.footer-widget-text {
  color: #1f2937; /* gray-800 */
}

.working-hours-simple {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.schedule-item {
  font-size: 0.9rem;
  color: #1f2937;
  direction: rtl;
  text-align: right;
  line-height: 1.6;
}
</style>
