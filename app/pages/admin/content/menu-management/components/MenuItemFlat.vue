<template>
  <div class="menu-item-flat" @contextmenu.prevent="openContextMenu">
    <div
      class="item-inner bg-white border border-gray-200 rounded-lg p-4 transition-all"
      :style="{ marginRight: `${item.depth * 40}px` }"
    >
      <div class="flex items-center gap-3">
        <div class="flex flex-col gap-1">
          <button
            type="button"
            class="control-button"
            :disabled="disableMoveUp"
            title="انتقال به بالا"
            @click.stop="$emit('move-up', index)"
          >
            ▲
          </button>
          <button
            type="button"
            class="control-button"
            :disabled="disableMoveDown"
            title="انتقال به پایین"
            @click.stop="$emit('move-down', index)"
          >
            ▼
          </button>
        </div>

        <div class="text-sm font-bold text-blue-500 w-6 flex-shrink-0 text-center select-none">
          {{ item.depth > 0 ? '↳'.repeat(item.depth) : '●' }}
        </div>

        <div class="flex-1 min-w-0 select-none">
          <div class="flex items-center gap-2">
            <span class="font-medium text-gray-900 truncate">{{ item.title }}</span>
            <span v-if="item.badge" class="px-2 py-0.5 text-xs rounded-full bg-blue-100 text-blue-800">{{ item.badge }}</span>
          </div>
          <div class="text-sm text-gray-500 truncate">{{ item.path }}</div>
        </div>

        <button
          type="button"
          :class="item.enabled ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'"
          class="px-2 py-1 text-xs rounded-full font-medium transition-colors"
          @click.stop="$emit('update-item', { ...item, enabled: !item.enabled })"
        >
          {{ item.enabled ? 'فعال' : 'غیرفعال' }}
        </button>

        <button
          type="button"
          class="p-1.5 text-gray-500 hover:text-gray-700 hover:bg-gray-100 rounded transition-colors"
          title="تنظیمات"
          @click.stop="$emit('toggle-expanded')"
        >
          <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
            <path d="M10 6a2 2 0 110-4 2 2 0 010 4zM10 12a2 2 0 110-4 2 2 0 010 4zM10 18a2 2 0 110-4 2 2 0 010 4z"/>
          </svg>
        </button>

        <button
          type="button"
          class="p-1.5 text-red-500 hover:text-red-700 hover:bg-red-50 rounded transition-colors"
          title="حذف"
          @click.stop="$emit('remove-item')"
        >
          <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd"/>
          </svg>
        </button>
      </div>

      <MenuItemSettings
        v-if="item.expanded"
        :item="item"
        class="mt-4 pt-4 border-t border-gray-200"
        @update-item="$emit('update-item', $event)"
      />
    </div>

    <div
      v-if="showContextMenu"
      class="context-menu"
      :style="{ top: `${contextMenuPosition.y}px`, left: `${contextMenuPosition.x}px` }"
      @click.stop
    >
      <button type="button" @click="handleMakeRoot">سر دسته شود</button>
      <button type="button" :disabled="index === 0" @click="handleMakeChild">زیر منو شود</button>
    </div>
  </div>
</template>

<script setup>
import { computed, onBeforeUnmount, onMounted, ref } from 'vue'
import MenuItemSettings from './MenuItemSettings.vue'

const props = defineProps({
  item: { type: Object, required: true },
  index: { type: Number, required: true },
  totalItems: { type: Number, required: true }
})

const emit = defineEmits([
  'update-item',
  'remove-item',
  'toggle-expanded',
  'move-up',
  'move-down',
  'make-root',
  'make-child'
])

const showContextMenu = ref(false)
const contextMenuPosition = ref({ x: 0, y: 0 })

const disableMoveUp = computed(() => props.index === 0)
const disableMoveDown = computed(() => props.index >= props.totalItems - 1)

const closeContextMenu = () => {
  showContextMenu.value = false
}

const openContextMenu = (event) => {
  contextMenuPosition.value = { x: event.clientX, y: event.clientY }
  showContextMenu.value = true
}

const handleMakeRoot = () => {
  emit('make-root', props.index)
  closeContextMenu()
}

const handleMakeChild = () => {
  emit('make-child', props.index)
  closeContextMenu()
}

onMounted(() => {
  window.addEventListener('click', closeContextMenu)
  window.addEventListener('scroll', closeContextMenu, true)
})

onBeforeUnmount(() => {
  window.removeEventListener('click', closeContextMenu)
  window.removeEventListener('scroll', closeContextMenu, true)
})
</script>

<style scoped>
.menu-item-flat {
  position: relative;
  margin-bottom: 8px;
}

.item-inner {
  cursor: default;
  transition: all 0.2s ease;
}

.item-inner:hover {
  box-shadow: 0 8px 16px rgba(15, 23, 42, 0.08);
}

.control-button {
  width: 28px;
  height: 28px;
  border: 1px solid #d1d5db;
  border-radius: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 12px;
  color: #4b5563;
  background: #fff;
  transition: all 0.2s ease;
}

.control-button:hover:not(:disabled) {
  background: #eff6ff;
  border-color: #60a5fa;
  color: #2563eb;
}

.control-button:disabled {
  opacity: 0.3;
  cursor: not-allowed;
}

.context-menu {
  position: fixed;
  z-index: 50;
  background: #1f2937;
  color: #fff;
  border-radius: 8px;
  padding: 4px;
  box-shadow: 0 12px 40px rgba(15, 23, 42, 0.4);
  min-width: 140px;
}

.context-menu button {
  width: 100%;
  text-align: right;
  padding: 8px 12px;
  font-size: 13px;
  border-radius: 6px;
  transition: background 0.2s ease;
}

.context-menu button:hover:not(:disabled) {
  background: rgba(255, 255, 255, 0.12);
}

.context-menu button:disabled {
  opacity: 0.4;
  cursor: not-allowed;
}
</style>
