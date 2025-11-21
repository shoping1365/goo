<template>
  <div class="menu-tree-editor bg-white rounded-lg shadow-md p-3 text-xs">
    <div class="font-bold text-gray-700 text-center border-b pb-2 mb-2">ساختار منو</div>
    <div v-if="items.length === 0" class="text-center text-gray-400 py-8">
      <i class="fas fa-list-ul text-2xl mb-2"></i>
      <div>هیچ آیتمی در منو وجود ندارد</div>
      <div class="text-xs mt-1">از پنل سمت چپ آیتم‌ها را به منو اضافه کنید</div>
    </div>
    <MenuTreeNode
      v-for="(item, idx) in items"
      :key="item.id"
      :item="item"
      :index="idx"
      :depth="0"
      :max-depth="maxDepth"
      :items="items"
      @edit="editItem"
      @delete="deleteItem"
      @drag-start="onDragStart"
      @drop="onDrop"
      @drag-over="onDragOver"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import MenuTreeNode from './MenuTreeNode.vue'

interface MenuItem {
  id: string | number
  title: string
  badge?: string
  children?: MenuItem[]
}

const props = defineProps<{
  items: MenuItem[]
  maxDepth?: number
}>()

const maxDepth = computed(() => props.maxDepth || 3)

interface DragInfo {
  item: unknown
  from: number[]
  to: number[]
  depth: number
}

const dragInfo = ref<DragInfo | null>(null)

function onDragStart(payload) {
  dragInfo.value = payload
}
function onDrop(payload) {
  if (!dragInfo.value) return
  // اگر عمق مجاز نیست، لغو
  if (payload.depth >= props.maxDepth) return
  // حذف آیتم از محل قبلی
  let arr = props.items
  let parent = null
  let idx = dragInfo.value.from[0]
  if (dragInfo.value.from.length > 1) {
    parent = findParent(props.items, dragInfo.value.from)
    arr = parent.children
    idx = dragInfo.value.from[dragInfo.value.from.length-1]
  }
  const [moved] = arr.splice(idx, 1)
  // اضافه به محل جدید
  let targetArr = props.items
  if (payload.to.length > 1) {
    const targetParent = findParent(props.items, payload.to)
    targetArr = targetParent.children
  }
  targetArr.splice(payload.to[payload.to.length-1], 0, moved)
  dragInfo.value = null
  // emit('update:items', props.items)
}
function onDragOver() {}
function editItem(_item: unknown) {
  // emit('edit', item)
}
function deleteItem(_item: unknown) {
  // emit('delete', item)
}

function findParent(arr, path) {
  let parent = null
  let cur = arr
  for (let i = 0; i < path.length-1; i++) {
    parent = cur[path[i]]
    if (!parent.children) parent.children = []
    cur = parent.children
  }
  return parent
}
</script>

<style scoped>
.menu-tree-editor {
  min-width: 260px;
  max-width: 340px;
  font-size: 13px;
}
</style> 
