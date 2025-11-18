<template>
  <div class="footer-widget footer-widget--working-hours">
    <h4 v-if="viewModel.title" class="footer-widget__title">{{ viewModel.title }}</h4>
    <div class="footer-widget__content">
      <div v-if="viewModel.schedules.length > 0" class="working-hours-table">
        <div 
          v-for="(schedule, index) in viewModel.schedules" 
          :key="index" 
          class="working-hours-row"
        >
          <span class="hours-text">{{ schedule }}</span>
        </div>
      </div>
      <div v-else class="footer-widget__item footer-widget__item--empty">
        ساعات کاری تعریف نشده است.
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = withDefaults(defineProps<{
  title?: string
  workingSchedules?: string[]
  simpleHours?: string
  showStatus?: boolean
}>(), {
  title: 'ساعات کاری',
  workingSchedules: () => [],
  showStatus: true
})

const viewModel = computed(() => {
  const schedules = (props.workingSchedules || []).filter(s => s && s.trim())
  
  return {
    title: props.title?.trim() || '',
    schedules,
    hasAny: schedules.length > 0
  }
})
</script>

<style scoped>
.footer-widget {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.footer-widget__title {
  margin: 0;
  font-size: 0.95rem;
  font-weight: 600;
  color: var(--footer-widget-text-color, #111827);
}

.footer-widget__content {
  display: flex;
  flex-direction: column;
}

.working-hours-table {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.working-hours-row {
  font-size: 0.9rem;
  color: var(--footer-widget-text-color, #374151);
  direction: rtl;
  text-align: right;
}

.hours-text {
  display: block;
  line-height: 1.6;
}

.footer-widget__item--empty {
  font-style: italic;
  color: var(--footer-widget-muted-color, #9ca3af);
  font-size: 0.9rem;
}
</style>
