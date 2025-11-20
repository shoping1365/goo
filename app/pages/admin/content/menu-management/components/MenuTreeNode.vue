<template>
  <div
class="menu-tree-node flex items-center gap-1 my-1 rounded hover:bg-gray-50 group"
       :style="{marginRight: depth * 18 + 'px'}"
       :draggable="true"
       @dragstart="handleDragStart"
       @dragover.prevent="handleDragOver"
       @drop="handleDrop"
  >
    <span class="cursor-grab text-gray-400 hover:text-blue-500 px-1"><i class="fas fa-grip-vertical"></i></span>
    <span class="font-semibold text-gray-700 truncate max-w-[120px]">{{ item.title }}</span>
    <span v-if="item.badge" class="bg-red-100 text-red-700 rounded px-1 text-[10px] ml-1">{{ item.badge }}</span>
    <button class="text-blue-500 hover:text-blue-700 px-1" @click="$emit('edit', item)"><i class="fas fa-edit"></i></button>
    <button 
      v-if="canDeleteMenuItem"
      class="text-red-500 hover:text-red-700 px-1" 
      @click="$emit('delete', item)"
    >
      <i class="fas fa-trash"></i>
    </button>
    <span class="flex-1"></span>
    <!-- زیرمنوها -->
    <div v-if="item.children && item.children.length && depth+1 < maxDepth" class="w-full">
      <MenuTreeNode
        v-for="(child, cidx) in item.children"
        :key="child.id"
        :item="child"
        :index="cidx"
        :depth="depth+1"
        :max-depth="maxDepth"
        :items="item.children"
        @edit="$emit('edit', $event)"
        @delete="$emit('delete', $event)"
        @drag-start="$emit('drag-start', $event)"
        @drop="$emit('drop', $event)"
        @drag-over="$emit('drag-over', $event)"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useAuth } from '~/composables/useAuth'
import MenuTreeNode from './MenuTreeNode.vue'

interface MenuItem {
  id: string | number
  title: string
  badge?: string
  children?: MenuItem[]
}

const props = defineProps<{
  item: MenuItem
  index: number
  depth: number
  maxDepth: number
  items: MenuItem[]
}>()

const emit = defineEmits<{
  edit: [item: MenuItem]
  delete: [item: MenuItem]
  'drag-start': [payload: { item: MenuItem, from: number[], depth: number }]
  drop: [payload: { item: MenuItem, to: number[], depth: number }]
  'drag-over': [payload: { item: MenuItem, to: number[], depth: number }]
}>()

const { user, hasPermission } = useAuth()

const canDeleteMenuItem = computed(() => hasPermission('menu.delete'))

function handleDragStart(e) {
  e.dataTransfer.effectAllowed = 'move'
  emit('drag-start', { item: props.item, from: getPath(), depth: props.depth })
}
function handleDrop(e) {
  emit('drop', { item: props.item, to: getPath(), depth: props.depth })
}
function handleDragOver(e) {
  emit('drag-over', { item: props.item, to: getPath(), depth: props.depth })
}
function getPath() {
  // مسیر این آیتم در آرایه والد
  const path = [props.index]
  // در Vue 3، $parent در دسترس نیست، بنابراین از props استفاده می‌کنیم
  // این یک راه‌حل ساده است که ممکن است نیاز به بهبود داشته باشد
  return path
}
</script>

<style scoped>
.menu-tree-node {
  padding: 2px 4px;
  min-height: 28px;
  user-select: none;
}
</style> 

