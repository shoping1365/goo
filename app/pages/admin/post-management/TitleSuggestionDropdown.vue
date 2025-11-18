<template>
  <div v-if="showDropdown" class="title-suggestion-dropdown">
    <div class="absolute z-50 bg-white border border-gray-200 rounded-lg shadow-lg max-w-xs">
      <div class="p-2">
        <div class="text-xs text-gray-500 mb-2 px-2">انتخاب عنوان از خلاصه نوشته:</div>
        <div class="space-y-1 max-h-40 overflow-y-auto">
          <button
            v-for="(title, index) in availableTitles"
            :key="index"
            @click="selectTitle(title, index + 1)"
            class="w-full text-right px-3 py-2 text-sm text-gray-700 hover:bg-blue-50 hover:text-blue-700 rounded transition-colors"
          >
            {{ title }}
          </button>
        </div>
        <div v-if="availableTitles.length === 0" class="px-3 py-2 text-sm text-gray-500">
          هیچ عنوانی در خلاصه نوشته وجود ندارد
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

// Props
const props = defineProps<{
  showDropdown: boolean
  excerpt: string
  position: { x: number; y: number } | null
}>()

// Emits
const emit = defineEmits<{
  'title-selected': [title: string, titleId: number]
  'close': []
}>()

// استخراج عناوین از خلاصه نوشته
const availableTitles = computed(() => {
  if (!props.excerpt) return []
  
  // تبدیل <br> به خط جدید و تقسیم بر اساس خط جدید
  const textWithNewlines = props.excerpt.replace(/<br\s*\/?>/gi, '\n')
  const lines = textWithNewlines.split('\n').map(line => line.trim()).filter(line => line.length > 0)
  
  // استخراج متن از HTML links
  const titles: string[] = []
  lines.forEach(line => {
    // اگر خط شامل HTML link است
    const linkMatch = line.match(/<a[^>]*>([^<]*)<\/a>/)
    if (linkMatch) {
      titles.push(linkMatch[1]) // متن داخل link
    } else if (line.length < 100 && !line.includes('http')) {
      // اگر خط ساده است و کوتاه
      titles.push(line)
    }
  })
  
  return titles
})

// انتخاب عنوان
const selectTitle = (title: string, titleId: number) => {
  emit('title-selected', title, titleId)
  emit('close')
}
</script>

<style scoped>
.title-suggestion-dropdown {
  position: fixed;
  top: v-bind('position?.y || 0');
  left: v-bind('position?.x || 0');
}

/* استایل برای scrollbar */
.overflow-y-auto {
  scrollbar-width: thin;
  scrollbar-color: #cbd5e0 #f7fafc;
}

.overflow-y-auto::-webkit-scrollbar {
  width: 4px;
}

.overflow-y-auto::-webkit-scrollbar-track {
  background: #f7fafc;
}

.overflow-y-auto::-webkit-scrollbar-thumb {
  background-color: #cbd5e0;
  border-radius: 2px;
}

.overflow-y-auto::-webkit-scrollbar-thumb:hover {
  background-color: #a0aec0;
}
</style> 
