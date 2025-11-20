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
        <!-- Quick Actions -->
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
          <label class="block text-sm font-medium text-gray-700 mb-2 whitespace-nowrap">نام منو</label>
          <input
            v-model="menu.name"
            type="text"
            placeholder="نام منو را وارد کنید"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            @input="$emit('update-slug')"
          />
        </div>

        <!-- Menu Slug Input -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2 whitespace-nowrap">شناسه منو (Slug)</label>
          <input
            v-model="menu.slug"
            type="text"
            placeholder="شناسه منو را وارد کنید"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>
      </div>

      <!-- Menu Items - لیست ساده flat مثل وردپرس -->
      <div 
        v-if="!hasItems" 
        class="text-center py-12 border-2 border-dashed border-gray-300 rounded-lg"
      >
        <div class="text-gray-500 mb-4">
          <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16" />
          </svg>
        </div>
        <h3 class="text-lg font-medium text-gray-900 mb-2">منوی شما خالی است</h3>
        <p class="text-gray-500">برای شروع، آیتم‌هایی را از پنل سمت چپ اضافه کنید</p>
      </div>
      <div 
        v-else 
        class="menu-container space-y-2"
      >
        <!-- لیست flat - همه آیتم‌ها در یک سطح -->
        <MenuItemFlat
          v-for="(item, index) in flatItems"
          :key="item.id ?? item.clientId ?? index"
          :item="item"
          :index="index"
          :all-items="flatItems"
          @update-item="handleUpdateFlatItem(index, $event)"
          @remove-item="removeFlatItem(index)"
          @toggle-expanded="toggleFlatItem(index)"
          @move-item="moveFlatItem"
        />
      </div>

      <!-- Save Button -->
      <div class="mt-6 pt-6 border-t border-gray-200">
        <button
          :disabled="!menu.name || isSaving"
          class="bg-green-200 hover:bg-green-300 text-green-900 px-6 py-2 rounded-lg font-bold transition-colors duration-150 border border-green-300 disabled:opacity-50 disabled:cursor-not-allowed"
          style="box-shadow: 0 2px 12px 0 rgba(34,197,94,0.10);"
          @click="$emit('save')"
        >
          {{ isSaving ? 'در حال ذخیره...' : 'ذخیره منو' }}
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, ref, onMounted } from 'vue'
import MenuItem from './MenuItem.vue'

const props = defineProps({
  menu: {
    type: Object,
    default: () => ({ items: [] }),
  },
  isSaving: Boolean,
})

const menu = computed({
  get: () => props.menu ?? { items: [] },
  set: (value) => {
    if (props.menu && value) {
      Object.assign(props.menu, value)
    }
  },
})

const menuItems = computed(() => {
  const rawItems = menu.value?.items
  return Array.isArray(rawItems) ? rawItems : []
})

const hasItems = computed(() => menuItems.value.length > 0)

const renderReady = ref(false)
onMounted(() => {
  renderReady.value = true
})

const emit = defineEmits([
  'update-slug',
  'update-item',
  'remove-item',
  'toggle-expanded',
  'save',
])

const dropIndicator = ref({ show: false, index: -1, indent: 0 })

const handleContainerDragOver = (e) => {
  e.preventDefault()
  e.dataTransfer.dropEffect = 'move'
  
  // محاسبه indent بر اساس فاصله از راست container
  const containerRect = e.currentTarget.getBoundingClientRect()
  const offsetFromRight = containerRect.right - e.clientX
  
  // هر 40px = 1 depth
  const indent = Math.floor(offsetFromRight / 40)
  const clampedIndent = Math.max(0, Math.min(indent, 2)) // max 2 levels
  
  // پیدا کردن index بر اساس موقعیت Y
  const items = e.currentTarget.querySelectorAll('.menu-item')
  let targetIndex = items.length
  
  items.forEach((item, idx) => {
    const rect = item.getBoundingClientRect()
    if (e.clientY < rect.top + rect.height / 2 && targetIndex === items.length) {
      targetIndex = idx
    }
  })
  
  dropIndicator.value = { show: true, index: targetIndex, indent: clampedIndent }
}

const handleContainerDrop = (e) => {
  e.preventDefault()
  
  try {
    const data = JSON.parse(e.dataTransfer.getData('application/json'))
    const containerRect = e.currentTarget.getBoundingClientRect()
    const offsetFromRight = containerRect.right - e.clientX
    
    // محاسبه depth بر اساس فاصله از راست
    const newDepth = Math.max(0, Math.min(Math.floor(offsetFromRight / 40), 2))
    
    // پیدا کردن target index
    const items = e.currentTarget.querySelectorAll('.menu-item')
    let targetIndex = items.length
    
    items.forEach((item, idx) => {
      const rect = item.getBoundingClientRect()
      if (e.clientY < rect.top + rect.height / 2 && targetIndex === items.length) {
        targetIndex = idx
      }
    })
    
    handleDropItem({
      draggedPath: data.path,
      targetIndex: targetIndex,
      newDepth: newDepth
    })
  } catch (error) {
    console.error('Drop error:', error)
  }
  
  dropIndicator.value = { show: false, index: -1, indent: 0 }
}

const handleDropItem = ({ draggedPath, targetIndex, newDepth }) => {
  console.log('🎯 WordPress-style drop:', { draggedPath, targetIndex, newDepth })
  
  // حذف آیتم از جای قبلی
  const flatItems = []
  const extractItems = (items, parentId = null) => {
    items.forEach(item => {
      flatItems.push({ ...item, parentId })
      if (item.children) {
        extractItems(item.children, item.id || item.clientId)
      }
    })
  }
  extractItems(menu.value.items)
  
  // پیدا کردن آیتمی که drag شده
  let draggedItem = menu.value.items[draggedPath[0]]
  if (draggedPath.length > 1) {
    let current = draggedItem
    for (let i = 1; i < draggedPath.length; i++) {
      current = current.children[draggedPath[i]]
    }
    draggedItem = current
  }
  
  if (!draggedItem) return
  
  // حذف از array
  if (draggedPath.length === 1) {
    menu.value.items.splice(draggedPath[0], 1)
  } else {
    let parent = menu.value.items[draggedPath[0]]
    for (let i = 1; i < draggedPath.length - 1; i++) {
      parent = parent.children[draggedPath[i]]
    }
    parent.children.splice(draggedPath[draggedPath.length - 1], 1)
  }
  
  // تنظیم depth و parentId
  if (newDepth === 0) {
    draggedItem.parentId = null
    // اضافه کردن به root در index مشخص
    menu.value.items.splice(targetIndex, 0, draggedItem)
  } else {
    // پیدا کردن parent مناسب
    if (targetIndex > 0) {
      const parentItem = menu.value.items[targetIndex - 1]
      draggedItem.parentId = parentItem.id || parentItem.clientId
      if (!parentItem.children) parentItem.children = []
      parentItem.children.push(draggedItem)
    }
  }
  
  console.log('✅ New structure:', menu.value.items)
}

const handleUpdateItem = (index, updatedItem) => {
  emit('update-item', index, updatedItem)
}
</script>
