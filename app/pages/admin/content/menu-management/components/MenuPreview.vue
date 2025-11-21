<template>
  <div class="lg:col-span-3 lg:w-[65%]">
    <div v-if="!renderReady" class="bg-white rounded-lg shadow-sm p-6 animate-pulse space-y-4">
      <div class="h-6 bg-gray-200 rounded" />
      <div class="h-10 bg-gray-100 rounded" />
      <div class="h-10 bg-gray-100 rounded" />
      <div class="h-32 bg-gray-50 rounded" />
      <div class="h-12 bg-gray-100 rounded" />
    </div>

    <div v-else class="bg-white rounded-lg shadow-sm p-6">
      <div class="flex items-center justify-between mb-4">
  <h3 class="text-lg font-semibold text-gray-900">ساختار منو</h3>
        <button
          class="px-3 py-1.5 bg-green-600 hover:bg-green-700 text-white text-sm rounded-lg font-medium transition-colors"
          title="فعال کردن همه آیتم‌ها"
          @click="enableAllItems"
        >
          ✓ فعال کردن همه
        </button>
      </div>
      
      <!-- Menu Name Input -->
      <div class="mb-6 grid gap-4 md:grid-cols-2 md:items-end">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نام منو</label>
          <input
            v-model="menu.name"
            type="text"
            placeholder="نام منو را وارد کنید"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            @input="$emit('update-slug')"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">شناسه منو (Slug)</label>
          <input
            v-model="menu.slug"
            type="text"
            placeholder="شناسه منو"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>
      </div>

      <!-- Flat Menu Items List -->
      <div v-if="!hasItems" class="text-center py-12 border-2 border-dashed border-gray-300 rounded-lg">
        <div class="text-gray-500 mb-4">
          <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16" />
          </svg>
        </div>
        <h3 class="text-lg font-medium text-gray-900 mb-2">منوی شما خالی است</h3>
        <p class="text-gray-500">برای شروع، آیتم‌هایی را از پنل سمت چپ اضافه کنید</p>
      </div>

      <div v-else class="space-y-2 min-h-[200px]">
        <MenuItemFlat
          v-for="(item, index) in flatItems"
          :key="item.id ?? item.clientId ?? index"
          :item="item"
          :index="index"
          :total-items="flatItems.length"
          @update-item="handleUpdateFlatItem(index, $event)"
          @remove-item="removeFlatItem(index)"
          @toggle-expanded="toggleFlatItem(index)"
          @move-up="moveItemUp"
          @move-down="moveItemDown"
          @make-root="makeItemRoot"
          @make-child="makeItemChild"
        />
      </div>

      <!-- Save Button -->
      <div class="mt-6 pt-6 border-t border-gray-200">
        <button
          :disabled="!menu.name || isSaving"
          class="bg-green-200 hover:bg-green-300 text-green-900 px-6 py-2 rounded-lg font-bold transition-colors duration-150 border border-green-300 disabled:opacity-50 disabled:cursor-not-allowed"
          @click="handleSave"
        >
          {{ isSaving ? 'در حال ذخیره...' : 'ذخیره منو' }}
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, onMounted, ref } from 'vue'
import MenuItemFlat from './MenuItemFlat.vue'

const props = defineProps({
  menu: { type: Object, default: () => ({ items: [] }) },
  isSaving: Boolean,
})

const emit = defineEmits(['update-slug', 'update-item', 'remove-item', 'toggle-expanded', 'save', 'update:menu'])

const menu = computed({
  get: () => new Proxy(props.menu ?? { items: [] }, {
    set(target, key, value) {
      emit('update:menu', { ...target, [key]: value })
      return true
    }
  }),
  set: (value) => {
    emit('update:menu', value)
  },
})

// تبدیل tree به flat list
const flatItems = computed(() => {
  const result = []
  const flatten = (items, depth = 0) => {
    if (!Array.isArray(items)) return
    items.forEach(item => {
      item.depth = depth
      result.push(item)
      if (item.children?.length > 0) flatten(item.children, depth + 1)
    })
  }
  if (Array.isArray(menu.value?.items)) flatten(menu.value.items)
  return result
})

const hasItems = computed(() => flatItems.value.length > 0)
const renderReady = ref(false)

const getItemId = (item) => item?.id ?? item?.clientId ?? null

const refreshMenuStructure = () => {
  const assign = (items, parent = null, depth = 0) => {
    if (!Array.isArray(items)) return
    items.forEach((item, index) => {
      item.order = index + 1
      item.depth = depth
      if (parent) {
        const parentDbId = typeof parent.id === 'number' ? parent.id : null
        const parentClientKey = typeof parent.clientId === 'number' ? parent.clientId : parentDbId
        item.parentId = parentDbId
        item.parentClientId = parentClientKey
      } else {
        item.parentId = null
        item.parentClientId = null
      }
      if (Array.isArray(item.children) && item.children.length > 0) {
        assign(item.children, item, depth + 1)
      } else if (!Array.isArray(item.children)) {
        item.children = []
      }
    })
  }
  assign(menu.value.items ?? [], null, 0)
}

onMounted(() => {
  renderReady.value = true
  refreshMenuStructure()
})

const handleSave = () => {
  refreshMenuStructure()
  emit('save')
}

const enableAllItems = () => {
  const setEnabled = (items) => {
    items.forEach(item => {
      item.enabled = true
      if (item.children) setEnabled(item.children)
    })
  }
  setEnabled(menu.value.items)
}

const handleUpdateFlatItem = (index, updatedItem) => {
  const findAndUpdate = (items) => {
    for (const item of items) {
      if ((item.id && item.id === updatedItem.id) || (item.clientId && item.clientId === updatedItem.clientId)) {
        Object.assign(item, updatedItem)
        return true
      }
      if (item.children && findAndUpdate(item.children)) return true
    }
    return false
  }
  findAndUpdate(menu.value.items)
}

const snapshotBranch = (node) => {
  if (!node) return null
  const base = {
    id: typeof node.id === 'number' ? node.id : null,
    clientId: typeof node.clientId === 'number' ? node.clientId : null,
    children: []
  }
  if (Array.isArray(node.children) && node.children.length > 0) {
    base.children = node.children
      .map(child => snapshotBranch(child))
      .filter(Boolean)
  }
  return base
}

const removeFlatItem = (index) => {
  const itemToRemove = flatItems.value[index]
  if (!itemToRemove) return
  const removedSnapshot = snapshotBranch(itemToRemove)
  const removeFromTree = (items) => {
    for (let i = 0; i < items.length; i++) {
      const item = items[i]
      if ((item.id && item.id === itemToRemove.id) || (item.clientId && item.clientId === itemToRemove.clientId)) {
        items.splice(i, 1)
        return true
      }
      if (item.children && removeFromTree(item.children)) return true
    }
    return false
  }
  removeFromTree(menu.value.items)
  refreshMenuStructure()
  if (removedSnapshot) emit('remove-item', removedSnapshot)
}

const toggleFlatItem = (index) => {
  const item = flatItems.value[index]
  if (!item) return
  emit('toggle-expanded', item)
}

const findItemContext = (items, targetId, ancestors = []) => {
  if (!Array.isArray(items)) return null
  for (let i = 0; i < items.length; i++) {
    const current = items[i]
    const currentId = getItemId(current)
    if (currentId === targetId) {
      const parent = ancestors.length ? ancestors[ancestors.length - 1] : null
      return {
        item: current,
        parent,
        siblings: items,
        index: i,
        ancestors
      }
    }
    if (current.children?.length) {
      const result = findItemContext(current.children, targetId, [...ancestors, current])
      if (result) return result
    }
  }
  return null
}

const moveItemUp = (index) => {
  const target = flatItems.value[index]
  if (!target) return
  const context = findItemContext(menu.value.items, getItemId(target))
  if (!context) return
  const { siblings, index: currentIndex } = context
  if (currentIndex <= 0) return
  const [item] = siblings.splice(currentIndex, 1)
  siblings.splice(currentIndex - 1, 0, item)
  refreshMenuStructure()
}

const moveItemDown = (index) => {
  const target = flatItems.value[index]
  if (!target) return
  const context = findItemContext(menu.value.items, getItemId(target))
  if (!context) return
  const { siblings, index: currentIndex } = context
  if (currentIndex >= siblings.length - 1) return
  const [item] = siblings.splice(currentIndex, 1)
  siblings.splice(currentIndex + 1, 0, item)
  refreshMenuStructure()
}

const makeItemRoot = (index) => {
  const target = flatItems.value[index]
  if (!target) return
  const targetId = getItemId(target)
  const context = findItemContext(menu.value.items, targetId)
  if (!context || !context.parent) return

  const { siblings, index: currentIndex, parent, ancestors } = context
  const [item] = siblings.splice(currentIndex, 1)
  if (!item) return

  const grandParent = ancestors.length > 1 ? ancestors[ancestors.length - 2] : null
  if (grandParent) {
    const grandId = getItemId(grandParent)
    const grandContext = findItemContext(menu.value.items, grandId)
    if (!grandContext) return
    if (!Array.isArray(grandContext.item.children)) grandContext.item.children = []
    const parentId = getItemId(parent)
    const parentIndex = grandContext.item.children.findIndex(child => getItemId(child) === parentId)
    const insertIndex = parentIndex >= 0 ? parentIndex + 1 : grandContext.item.children.length
    grandContext.item.children.splice(insertIndex, 0, item)
  } else {
    const rootItems = menu.value.items ?? []
    const parentId = getItemId(parent)
    const parentIndex = rootItems.findIndex(rootItem => getItemId(rootItem) === parentId)
    const insertIndex = parentIndex >= 0 ? parentIndex + 1 : rootItems.length
    rootItems.splice(insertIndex, 0, item)
    menu.value.items = rootItems
  }

  refreshMenuStructure()
}

const makeItemChild = (index) => {
  if (index <= 0) return
  const flat = flatItems.value
  const target = flat[index]
  if (!target) return

  const targetDepth = Math.max(0, Number(target?.depth) || 0)
  let parentCandidate = null
  let scanIndex = index - 1

  while (scanIndex >= 0) {
    const candidate = flat[scanIndex]
    if (!candidate) {
      scanIndex--
      continue
    }
    const candidateId = getItemId(candidate)
    const targetId = getItemId(target)
    if (candidateId === targetId) {
      scanIndex--
      continue
    }
    const candidateDepth = Math.max(0, Number(candidate?.depth) || 0)
    if (candidateDepth < targetDepth + 1) {
      parentCandidate = candidate
      break
    }
    scanIndex--
  }

  if (!parentCandidate) return

  const targetContext = findItemContext(menu.value.items, getItemId(target))
  const parentContext = findItemContext(menu.value.items, getItemId(parentCandidate))
  if (!targetContext || !parentContext) return

  const { siblings, index: currentIndex } = targetContext
  const [item] = siblings.splice(currentIndex, 1)
  if (!item) return

  if (!Array.isArray(parentContext.item.children)) parentContext.item.children = []
  parentContext.item.children.push(item)
  refreshMenuStructure()
}
</script>
