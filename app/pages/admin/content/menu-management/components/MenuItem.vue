<template>
  <div
    class="menu-item"
    :style="{ 
      marginRight: `${depth * 40}px`,
      transform: isDragging ? 'opacity: 0.4' : ''
    }"
    draggable="true"
    @dragstart="handleDragStart"
    @dragend="handleDragEnd"
  >
    <!-- Menu Item Header -->
    <div class="menu-item-inner bg-white rounded-lg border border-gray-200 shadow-sm hover:shadow-md transition-shadow duration-200">
      <div class="p-6">
      <div class="flex items-center justify-between">
        <!-- Left Side: Drag Handle + Icon + Content -->
        <div class="flex items-center space-x-3 space-x-reverse flex-1 min-w-0">
          <!-- Drag Handle -->
          <div class="text-gray-400 cursor-move hover:text-gray-600 transition-colors">
            <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
              <path d="M7 2a2 2 0 1 1 .001 4.001A2 2 0 0 1 7 2zm0 6a2 2 0 1 1 .001 4.001A2 2 0 0 1 7 8zm0 6a2 2 0 1 1 .001 4.001A2 2 0 0 1 7 14zm6-8a2 2 0 1 1-.001-4.001A2 2 0 0 1 13 6zm0 2a2 2 0 1 1 .001 4.001A2 2 0 0 1 13 8zm0 6a2 2 0 1 1 .001 4.001A2 2 0 0 1 13 14z"/>
            </svg>
          </div>

          <!-- Icon/Image -->
          <div class="flex-shrink-0">
            <div v-if="getItemIcon(item)" class="w-8 h-8 bg-blue-50 rounded-lg flex items-center justify-center text-blue-600">
              <i :class="getItemIcon(item)" class="text-lg"></i>
            </div>
            <div v-else-if="getItemImage(item)" class="w-8 h-8 rounded-lg overflow-hidden border border-gray-200">
              <img
                :src="getItemImage(item).url"
                :alt="item.title"
                class="w-full h-full object-cover"
              />
            </div>
            <div v-else class="w-8 h-8 bg-gray-100 rounded-lg flex items-center justify-center">
              <svg class="w-4 h-4 text-gray-400" fill="currentColor" viewBox="0 0 20 20">
                <path d="M3 4a1 1 0 011-1h12a1 1 0 011 1v2a1 1 0 01-1 1H4a1 1 0 01-1-1V4zM3 10a1 1 0 011-1h6a1 1 0 011 1v6a1 1 0 01-1 1H4a1 1 0 01-1-1v-6zM14 9a1 1 0 00-1 1v6a1 1 0 001 1h2a1 1 0 001-1v-6a1 1 0 00-1-1h-2z"/>
              </svg>
            </div>
          </div>

          <!-- Content -->
          <div class="flex-1 min-w-0">
            <div class="flex items-center space-x-2 space-x-reverse">
              <h4 class="font-medium text-gray-900 truncate">{{ item.title }}</h4>
              <!-- Badge if exists -->
              <span v-if="item.badge" :class="getBadgeClasses(item.badgeColor)" class="px-2 py-0.5 text-xs rounded-full font-medium">
                {{ item.badge }}
              </span>
            </div>
            <div class="flex items-center space-x-3 space-x-reverse mt-1">
              <span class="text-sm text-gray-500 truncate">{{ item.path }}</span>
              <!-- Item Type Badge -->
              <span class="px-2 py-0.5 text-xs bg-gray-100 text-gray-600 rounded-full">
                {{ getItemTypeLabel(item.itemType) }}
              </span>
              <!-- Child count indicator -->
              <span v-if="item.children && item.children.length > 0" class="px-2 py-0.5 text-xs bg-blue-100 text-blue-800 rounded-full">
                {{ item.children.length }} زیرمنو
              </span>
            </div>
          </div>
        </div>

        <!-- Right Side: Status + Actions -->
        <div class="flex items-center space-x-2 space-x-reverse flex-shrink-0">
          <!-- Status Badge (Clickable to open settings) -->
          <button 
            :class="item.enabled ? 'bg-green-100 text-green-800 hover:bg-green-200' : 'bg-red-100 text-red-800 hover:bg-red-200'"
            class="px-2 py-1 text-xs rounded-full font-medium transition-colors cursor-pointer"
            :title="item.enabled ? 'فعال - کلیک برای تنظیمات' : 'غیرفعال - کلیک برای تنظیمات'"
            @click="$emit('toggle-expanded')"
          >
            {{ item.enabled ? 'فعال' : 'غیرفعال' }}
          </button>

          <!-- Mega Menu Indicator -->
          <span v-if="item.isMegaMenu" class="px-2 py-1 text-xs bg-purple-100 text-purple-800 rounded-full font-medium">
            مگا
          </span>

          <!-- Expand/Collapse Button -->
          <button
            class="p-1 text-gray-500 hover:text-gray-700 hover:bg-gray-100 rounded transition-colors"
            :title="item.expanded ? 'بستن تنظیمات' : 'باز کردن تنظیمات'"
            @click="$emit('toggle-expanded')"
          >
            <svg 
              class="w-4 h-4 transition-transform duration-200" 
              :class="{ 'rotate-180': item.expanded }"
              fill="currentColor" 
              viewBox="0 0 20 20"
            >
              <path fill-rule="evenodd" d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z" clip-rule="evenodd"/>
            </svg>
          </button>

          <!-- Remove Button -->
          <button
            class="p-1 text-red-500 hover:text-red-700 hover:bg-red-50 rounded transition-colors"
            title="حذف آیتم"
            @click="$emit('remove-item')"
          >
            <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd"/>
            </svg>
          </button>
        </div>
      </div>
    </div>

      <!-- Menu Item Expanded Content -->
      <MenuItemSettings
        v-if="item.expanded"
        :item="item"
        @update-item="$emit('update-item', $event)"
        @open-icon-selector="$emit('open-icon-selector')"
        @open-image-selector="$emit('open-image-selector')"
      />
    </div>

    <!-- Children Items (Recursive) -->
    <div v-if="item.children && item.children.length > 0" class="menu-item-children mt-2">
      <MenuItem
        v-for="(child, childIndex) in item.children"
        :key="child.id ?? child.clientId ?? childIndex"
        :item="child"
        :index="childIndex"
        :depth="depth + 1"
        :parent-path="currentPath"
        @update-item="$emit('update-child', currentPath, childIndex, $event)"
        @remove-item="$emit('remove-child', currentPath, childIndex)"
        @toggle-expanded="$emit('toggle-child-expanded', currentPath, childIndex)"
        @drop-item="$emit('drop-item', $event)"
      />
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import MenuItemSettings from './MenuItemSettings.vue'

const props = defineProps({
  item: Object,
  index: Number,
  depth: {
    type: Number,
    default: 0
  },
  parentPath: {
    type: Array,
    default: () => []
  }
})

const _emit = defineEmits([
  'toggle-expanded',
  'drop-item',
  'remove-item',
  'update-item',
  'update-child',
  'remove-child',
  'toggle-child-expanded',
  'open-icon-selector',
  'open-image-selector',
])

const _dragIndent = ref(null)
const isDragging = ref(false)

const currentPath = computed(() => [...props.parentPath, props.index])

const handleDragStart = (e) => {
  isDragging.value = true
  e.dataTransfer.effectAllowed = 'move'
  e.dataTransfer.setData('application/json', JSON.stringify({
    path: currentPath.value,
    item: props.item,
    depth: props.depth
  }))
  
  // Set drag image
  e.dataTransfer.setDragImage(e.target, 20, 20)
}

const handleDragEnd = () => {
  isDragging.value = false
}

const getItemIcon = (item) => {
  if (!item.icon) return null
  if (typeof item.icon === 'string' && !item.icon.startsWith('{')) {
    return item.icon
  }
  try {
    const parsed = JSON.parse(item.icon)
    if (parsed.type === 'image') {
      return null
    }
  } catch (_e) {
    return item.icon
  }
  return null
}

const getItemImage = (item) => {
  if (!item.icon) return null
  try {
    const parsed = JSON.parse(item.icon)
    if (parsed.type === 'image') {
      return parsed
    }
  } catch (_e) {}
  return null
}

const getBadgeClasses = (color) => {
  const colorClasses = {
    red: 'bg-red-100 text-red-800',
    green: 'bg-green-100 text-green-800',
    blue: 'bg-blue-100 text-blue-800',
    yellow: 'bg-yellow-100 text-yellow-800',
    purple: 'bg-purple-100 text-purple-800',
    gray: 'bg-gray-100 text-gray-800'
  }
  return colorClasses[color] || colorClasses.gray
}

const getItemTypeLabel = (itemType) => {
  const typeLabels = {
    page: 'برگه',
    post: 'نوشته',
    custom: 'سفارشی',
    category: 'دسته‌بندی',
    product_category: 'دسته محصول'
  }
  return typeLabels[itemType] || 'نامشخص'
}
</script>

<style scoped>
.menu-item {
  position: relative;
  margin-bottom: 8px;
  transition: margin-right 0.2s ease-in-out;
  cursor: grab;
}

.menu-item:active {
  cursor: grabbing;
}

.menu-item.is-dragging {
  opacity: 0.5;
  cursor: grabbing;
}

/* فقط نمایش جهت indent */
.menu-item.indent-right .menu-item-inner {
  transform: translateX(32px);
  background: #dbeafe;
  border-color: #3b82f6;
}

.menu-item.indent-right .menu-item-inner::after {
  content: '← زیرمنو';
  position: absolute;
  top: 50%;
  left: -60px;
  transform: translateY(-50%);
  background: #3b82f6;
  color: white;
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 11px;
  font-weight: bold;
}

.menu-item.indent-left .menu-item-inner {
  transform: translateX(-32px);
  background: #fef3c7;
  border-color: #f59e0b;
}

.menu-item.indent-left .menu-item-inner::after {
  content: 'سر دسته →';
  position: absolute;
  top: 50%;
  right: -70px;
  transform: translateY(-50%);
  background: #f59e0b;
  color: white;
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 11px;
  font-weight: bold;
}

.menu-item-inner {
  transition: all 0.2s ease-in-out;
}

.menu-item-inner:hover {
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.menu-item-children {
  padding-right: 24px;
  border-right: 2px dashed #e5e7eb;
}
</style>
