<template>
  <div class="widget-area">
    <div v-if="loading" class="flex justify-center items-center py-8">
      <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-500"></div>
    </div>


    <div v-else class="space-y-8">
      <WidgetRenderer
        v-for="widget in sortedWidgets"
        :key="widget.id"
        :widget="widget"
        class="widget-item"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import type { Widget, WidgetPage } from '~/types/widget'
import { useWidget } from '~/composables/useWidget'
import WidgetRenderer from './WidgetRenderer.vue'

interface Props {
  page: string
}

const props = defineProps<Props>()

// Composables
const { fetchWidgetsByPage, loading, refreshTrigger } = useWidget()

// State
const widgets = ref<Widget[]>([])
const error = ref<string | null>(null)

// Computed
const sortedWidgets = computed(() => {
  // اطمینان از اینکه widgets.value یک آرایه است
  if (!widgets.value || !Array.isArray(widgets.value)) {
    return []
  }
  
  return [...widgets.value]
    .filter(widget => widget.status === 'active')
    .sort((a, b) => a.order - b.order)
})

// Methods
const loadWidgets = async () => {
  try {
    error.value = null

    // Type safety improvement
    const pageWidgets = await fetchWidgetsByPage(props.page as WidgetPage)

    widgets.value = pageWidgets
  } catch (err: any) {
    // خطا در بارگذاری ابزارک‌ها
    widgets.value = []
    error.value = err.message || 'خطا در بارگذاری ویجت‌ها'
  }
}

// متد برای تازه کردن ویجت‌ها از بیرون (مثل وقتی که در ادمین تغییر داده می‌شود)
const forceRefresh = async () => {
  await loadWidgets()
}

// expose methods for external use
defineExpose({
  refreshWidgets: forceRefresh
})

// Load widgets on mount
onMounted(() => {
  loadWidgets()
})

// Watch for page changes
watch(() => props.page, () => {
  loadWidgets()
})

// Watch for global refresh trigger
watch(refreshTrigger, (newVal, oldVal) => {
  if (newVal > oldVal) {
    loadWidgets()
  }
})
</script>

<style scoped>

</style>