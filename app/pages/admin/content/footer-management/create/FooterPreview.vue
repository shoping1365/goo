.chip-value {
  font-family: 'Fira Code', monospace;
  direction: ltr;
}

@media (max-width: 768px) {
  .info-capsule {
    min-width: calc(50% - 8px);
  }

  .width-items-inline {
    gap: 6px;
  }

  .width-chip {
    font-size: 10px;
    padding: 4px 8px;
  }
}
<template>
  <div class="preview-section" :class="{ 'full-width': isEditing }">
    <div class="footer-info">
      <div class="footer-form-row"></div>
    </div>

    <h3>پیش نمایش فوتر</h3>
    
    <!-- نمایش پیش‌نمایش لایه‌های ایجاد شده فقط وقتی که در حال ویرایش لایه نیستیم -->
    <div v-if="!showLayerSettings && createdLayers.length" class="footer-preview">
      <div class="created-layers-preview">
        <div
          v-for="(layer, layerIndex) in createdLayers"
          :key="layer.id || layerIndex"
          class="preview-layer"
          :style="getLayerStyle(layer, layerIndex)"
        >
          <div class="layer-name-preview">{{ layer.name || 'لایه بدون نام' }}</div>
          <div class="layer-items-preview-horizontal">
            <div
              v-for="(itemObj, itemIndex) in parseLayerItems(layer.items)"
              :key="itemIndex"
              class="layer-item-preview-box"
              :style="getPreviewItemStyle(itemObj, layer.items, itemIndex)"
            >
              <div class="item-preview-content">
                <div class="item-preview-name">{{ getItemDisplayName(itemObj) }}</div>
                <div class="item-preview-width">{{ getPreviewWidth(itemObj, layer.items) }}%</div>
              </div>
            </div>
          </div>
          <div
            v-if="layer.showSeparator && layerIndex < createdLayers.length - 1"
            class="preview-separator-horizontal"
            :style="getSeparatorStyle(layer)"
          ></div>
        </div>
      </div>
    </div>

    <!-- نمایش پیش‌نمایش زنده فقط در حال ویرایش/افزودن لایه -->
    <div v-if="showLayerSettings" class="footer-preview">
      <div class="footer-preview-actions">
        <button type="button" class="btn btn-outline items-selector" @click="openItemsModal">
          <span>{{ getSelectedItemsText() }}</span>
          <span>📋</span>
        </button>
      </div>

      <div class="preview-footer">
        <div class="preview-footer-label">
          <strong>پیش‌نمایش زنده لایه:</strong>
        </div>
        <div class="preview-container" :style="getActiveLayerStyle()">
          <div
            ref="previewItemsContainer"
            class="preview-items-container"
            :style="getContainerAlignment()"
          >
            <template v-for="(itemObj, itemIndex) in currentItems" :key="itemIndex">
              <div
                class="preview-section-item"
                :class="{ 
                  'is-active': activeIndex === itemIndex,
                  'is-dragging': draggingIndex === itemIndex,
                  'drag-over': dragOverIndex === itemIndex
                }"
                :data-align="itemObj.align || 'center'"
                :style="getItemStyle(itemObj, itemIndex)"
                :draggable="true"
                @click="toggleItemSettings(itemIndex)"
                @dragstart="handleDragStart(itemIndex, $event)"
                @dragend="handleDragEnd"
                @dragover.prevent="handleDragOver(itemIndex, $event)"
                @drop="handleDrop(itemIndex, $event)"
              >
                <div class="item-preview-content">
                  <div class="drag-handle">⋮⋮</div>
                  <div class="item-preview-icon">{{ getItemIcon(itemObj) }}</div>
                  <div class="item-preview-name">{{ getItemDisplayName(itemObj) }}</div>
                  <div class="item-preview-width">{{ formatWidth(itemObj.width) }}%</div>
                </div>
              </div>
            </template>
          </div>
        </div>

        <!-- تنظیمات آیتم انتخاب‌شده -->
        <div v-if="activeIndex !== null && currentItems[activeIndex]" class="item-settings-panel">
          <div class="settings-header">
            <h4>تنظیمات {{ getItemDisplayName(currentItems[activeIndex]) }}</h4>
            <button class="close-settings-btn" @click="closeActiveSettings">✕</button>
          </div>

          <div class="settings-content">
            <div class="settings-grid">
              <div class="form-field">
                <label>چینش محتوا</label>
                <select v-model="currentItems[activeIndex].align" class="form-control">
                  <option value="right">راست</option>
                  <option value="center">وسط</option>
                  <option value="left">چپ</option>
                </select>
              </div>

              <div class="form-field">
                <label>عرض (%)</label>
                <input
                  type="number"
                  v-model.number="currentItems[activeIndex].width"
                  min="5"
                  max="95"
                  step="0.1"
                  class="form-control"
                  @blur="normalizeWidth(activeIndex)"
                  dir="ltr"
                />
              </div>

              <div class="form-field">
                <label>پدینگ راست (px)</label>
                <input
                  type="number"
                  v-model.number="currentItems[activeIndex].paddingRight"
                  min="0"
                  max="200"
                  class="form-control"
                  dir="ltr"
                />
              </div>

              <div class="form-field">
                <label>پدینگ چپ (px)</label>
                <input
                  type="number"
                  v-model.number="currentItems[activeIndex].paddingLeft"
                  min="0"
                  max="200"
                  class="form-control"
                  dir="ltr"
                />
              </div>

              <div
                v-if="currentItems[activeIndex].id === 'image'"
                class="form-field form-field--full"
              >
                <label>تصویر</label>
                <div v-if="currentItems[activeIndex].imageUrl" class="selected-image-preview">
                  <img :src="currentItems[activeIndex].imageUrl" alt="تصویر انتخاب‌شده" />
                  <button @click="removeSelectedImage" class="btn btn-sm btn-outline">حذف تصویر</button>
                </div>
                <button v-else @click="openMediaLibrary" class="btn btn-sm btn-primary">انتخاب تصویر</button>
              </div>

              <div
                v-if="currentItems[activeIndex].id === 'copyright'"
                class="form-field form-field--full rich-text-editor-wrapper"
              >
                <label>متن کپی‌رایت</label>
                <ClientOnly>
                  <template #default>
                    <Suspense>
                      <template #default>
                        <RichTextEditor
                          v-model="currentItems[activeIndex].props.text"
                          :height="220"
                          lang="fa"
                          direction="rtl"
                        />
                      </template>
                      <template #fallback>
                        <div class="editor-loading">در حال بارگذاری ویرایشگر...</div>
                      </template>
                    </Suspense>
                  </template>
                  <template #fallback>
                    <textarea
                      v-model="currentItems[activeIndex].props.text"
                      class="form-control"
                      rows="5"
                      placeholder="© 2025 شرکت شما. تمامی حقوق محفوظ است."
                    ></textarea>
                  </template>
                </ClientOnly>
              </div>
            </div>

            <button class="remove-item-btn" @click="removeItem(activeIndex)">
              حذف آیتم
            </button>
          </div>
        </div>

        <MediaLibraryModal
          v-if="activeIndex !== null && currentItems[activeIndex]?.id === 'image'"
          v-model="showMediaLibrary"
          :model-selected="selectedMediaIds"
          default-category="banners"
          file-type="image"
          context-title="فوتر - انتخاب تصویر"
          @confirm="onMediaSelected"
        />
      </div>

      <div v-if="showLayerSettings" class="layer-info">
        <div class="info-row">
          <div class="info-capsule">
            <span class="info-label">نام لایه</span>
            <span class="info-value">{{ newLayer.name || 'بدون نام' }}</span>
          </div>
          <div class="info-capsule">
            <span class="info-label">عرض</span>
            <span class="info-value">{{ newLayer.width }}%</span>
          </div>
          <div class="info-capsule">
            <span class="info-label">ارتفاع</span>
            <span class="info-value">{{ newLayer.height }}px</span>
          </div>
          <div class="info-capsule">
            <span class="info-label">شفافیت</span>
            <span class="info-value">{{ newLayer.opacity }}%</span>
          </div>
          <div class="info-capsule">
            <span class="info-label">تعداد آیتم</span>
            <span class="info-value">{{ currentItems.length }}</span>
          </div>
          <div v-if="currentItems.length" class="info-capsule emphasis">
            <span class="info-label">مجموع عرض‌ها</span>
            <span :class="['info-value', 'total-width', { valid: isTotalWidthBalanced, invalid: !isTotalWidthBalanced }]">
              {{ totalWidth.toFixed(1) }}%
            </span>
          </div>
          <button
            v-if="currentItems.length > 1"
            class="btn btn-sm btn-outline capsule-action"
            type="button"
            @click="resetItemWidths"
          >
            تنظیم مساوی
          </button>
        </div>
      </div>

      <div
        v-if="!showLayerSettings && !createdLayers.length"
        class="empty-preview"
      >
        <div class="empty-message">
          <p>هیچ لایه‌ای ایجاد نشده است</p>
          <p>برای شروع، یک لایه جدید ایجاد کنید</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, inject, onBeforeUnmount, ref, watch, defineAsyncComponent } from 'vue'
import type { CSSProperties, Ref } from 'vue'
import MediaLibraryModal from '~/components/media/MediaLibraryModal.vue'

const RichTextEditor = defineAsyncComponent(() => import('~/components/common/RichTextEditor.vue'))

type FooterItem = {
  id?: string
  name?: string
  icon?: string
  width?: number
  align?: string
  bgColor?: string
  paddingRight?: number
  paddingLeft?: number
  props?: Record<string, any>
  config?: Record<string, any>
  [key: string]: any
}

type SocialEntry = {
  id: string
  platform: string
  label: string
  url: string
  enabled: boolean
  openInNewTab: boolean
  title?: string
  href?: string
  visible?: boolean
}

const createdLayers = inject<Ref<any[]>>('createdLayers', ref([]))
const newLayer = inject<Ref<any>>('newLayer', ref({ items: [] }))
const showLayerSettings = inject<Ref<boolean>>('showLayerSettings', ref(false))
const openItemsModal = inject<() => void>('openItemsModal', () => {})
const getSelectedItemsText = inject<() => string>('getSelectedItemsText', () => 'انتخاب آیتم‌ها')
const isEditing = inject<Ref<boolean>>('isEditing', ref(false))
const availableItemsSource = inject<any>('availableItems', [])

const previewItemsContainer = ref<HTMLElement | null>(null)
const activeIndex = ref<number | null>(null)
const isResizing = ref(false)
const resizeStartX = ref(0)
const resizeItemIndex = ref(-1)
const containerWidth = ref(1)
const initialWidths = ref<number[]>([])

// Drag & Drop state
const draggingIndex = ref<number | null>(null)
const dragOverIndex = ref<number | null>(null)

const showMediaLibrary = ref(false)
const selectedMediaIds = ref<number[]>([])

const SOCIAL_PRESETS = [
  { platform: 'instagram', label: 'اینستاگرام' },
  { platform: 'telegram', label: 'تلگرام' },
  { platform: 'twitter', label: 'توییتر' },
  { platform: 'linkedin', label: 'لینکدین' },
  { platform: 'youtube', label: 'یوتیوب' },
  { platform: 'facebook', label: 'فیسبوک' },
  { platform: 'whatsapp', label: 'واتساپ' }
]

const LEGACY_SOCIAL_KEYS = ['instagram', 'telegram', 'twitter', 'linkedin', 'youtube', 'facebook', 'whatsapp'] as const

const resolvedAvailableItems = computed(() => {
  if (availableItemsSource && typeof availableItemsSource === 'object' && 'value' in availableItemsSource) {
    return Array.isArray(availableItemsSource.value) ? availableItemsSource.value : []
  }
  return Array.isArray(availableItemsSource) ? availableItemsSource : []
})

const currentItems = computed<FooterItem[]>(() => {
  const items = newLayer.value?.items
  if (Array.isArray(items)) return items as FooterItem[]
  if (typeof items === 'string') {
    try {
      return JSON.parse(items)
    } catch {
      return []
    }
  }
  return []
})

const layerOpacity = computed(() => Number(newLayer.value?.opacity ?? 100))
const layerHeight = computed(() => Number(newLayer.value?.height ?? 50))
const layerWidth = computed(() => Number(newLayer.value?.width ?? 100))

const fallbackWidth = computed(() => {
  const count = currentItems.value.length || 1
  return Number((WIDTH_LIMIT / count).toFixed(2))
})

const WIDTH_LIMIT = 100
const BASE_MIN_WIDTH = 5
let adjustingWidths = false

const totalWidth = computed(() => {
  if (!currentItems.value.length) return 0
  const total = currentItems.value.reduce((sum, item) => sum + (typeof item.width === 'number' ? item.width : fallbackWidth.value), 0)
  return Number(total.toFixed(2))
})

const isTotalWidthBalanced = computed(() => {
  if (currentItems.value.length <= 1) return true
  return Math.abs(totalWidth.value - WIDTH_LIMIT) < 0.01
})

watch(
  currentItems,
  items => {
    if (!items.length) return
    enforceWidthBudget()
  },
  { deep: true, immediate: true }
)

function parseLayerItems(items: any): FooterItem[] {
  if (Array.isArray(items)) return items
  if (typeof items === 'string') {
    try {
      return JSON.parse(items)
    } catch {
      return []
    }
  }
  return []
}

function getPreviewWidth(item: FooterItem, original: any): string {
  if (typeof item.width === 'number') return item.width.toFixed(1)
  const list = parseLayerItems(original)
  if (!list.length) return '0.0'
  return (100 / list.length).toFixed(1)
}

function formatWidth(width?: number) {
  if (typeof width !== 'number' || Number.isNaN(width)) return fallbackWidth.value.toFixed(1)
  return width.toFixed(1)
}

function ensureItemDefaults(item: FooterItem) {
  if (!item.align) item.align = 'center'
  if (typeof item.paddingRight !== 'number' || Number.isNaN(item.paddingRight)) item.paddingRight = 0
  if (typeof item.paddingLeft !== 'number' || Number.isNaN(item.paddingLeft)) item.paddingLeft = 0
  ensureItemProps(item)
}

function ensureItemProps(item: FooterItem) {
  if (!item) return
  if (!item.props || typeof item.props !== 'object') {
    item.props = {}
  }
  if (item.id === 'copyright') {
    if (typeof item.props.text !== 'string') {
      const year = new Date().getFullYear()
      item.props.text = `© ${year} نام شرکت. تمامی حقوق محفوظ است.`
    }
  }
}

function getEffectiveMinWidth(count: number): number {
  if (count <= 0) return BASE_MIN_WIDTH
  const minCandidate = Math.min(BASE_MIN_WIDTH, WIDTH_LIMIT / count)
  return Number(minCandidate.toFixed(2))
}

function resolveBackground(color?: string | null): string {
  if (typeof color === 'string' && color.trim().length) return color
  return 'transparent'
}

function getMaxWidthPerItem(count: number, minWidth: number): number {
  if (count <= 1) return WIDTH_LIMIT
  const maxCandidate = WIDTH_LIMIT - minWidth * (count - 1)
  return Number(Math.max(maxCandidate, minWidth).toFixed(2))
}

function clampWidth(value: number, minWidth: number, maxWidth: number): number {
  if (!Number.isFinite(value)) return minWidth
  return Number(Math.min(Math.max(value, minWidth), maxWidth).toFixed(2))
}

function enforceWidthBudgetInternal(changedIndex?: number) {
  const items = currentItems.value
  const count = items.length
  if (count === 0) return

  const minWidth = getEffectiveMinWidth(count)
  const maxWidth = getMaxWidthPerItem(count, minWidth)
  const indices = items.map((_, idx) => idx)
  const pivot = typeof changedIndex === 'number' ? changedIndex : indices[0]

  if (count === 1) {
    ensureItemDefaults(items[0])
    const requested = typeof items[0].width === 'number' ? items[0].width : WIDTH_LIMIT
    items[0].width = clampWidth(requested, minWidth, maxWidth)
    return
  }

  indices.forEach(idx => {
    const item = items[idx]
    ensureItemDefaults(item)
    const requested = typeof item.width === 'number' ? item.width : WIDTH_LIMIT / count
    item.width = clampWidth(requested, minWidth, maxWidth)
  })

  if (pivot >= 0 && pivot < items.length) {
    items[pivot].width = clampWidth(items[pivot].width, minWidth, maxWidth)
  }

  let total = Number(items.reduce((sum, item) => sum + item.width, 0).toFixed(2))

  if (total > WIDTH_LIMIT) {
    let excess = Number((total - WIDTH_LIMIT).toFixed(2))
    const reduceOrder = indices.filter(idx => idx !== pivot).concat(pivot)
    for (const idx of reduceOrder) {
      if (excess <= 0.01) break
      const item = items[idx]
      const reducible = Number((item.width - minWidth).toFixed(2))
      if (reducible <= 0) continue
      const delta = Math.min(reducible, excess)
      item.width = Number((item.width - delta).toFixed(2))
      excess = Number((excess - delta).toFixed(2))
    }
  } else if (total < WIDTH_LIMIT) {
    let deficit = Number((WIDTH_LIMIT - total).toFixed(2))
    const increaseOrder = indices.filter(idx => idx !== pivot).concat(pivot)
    for (const idx of increaseOrder) {
      if (deficit <= 0.01) break
      const item = items[idx]
      const room = Number((maxWidth - item.width).toFixed(2))
      if (room <= 0) continue
      const delta = Math.min(room, deficit)
      item.width = Number((item.width + delta).toFixed(2))
      deficit = Number((deficit - delta).toFixed(2))
    }
  }

  total = Number(items.reduce((sum, item) => sum + item.width, 0).toFixed(2))
  let diff = Number((WIDTH_LIMIT - total).toFixed(2))
  if (Math.abs(diff) > 0.05) {
    const adjustOrder = indices.filter(idx => idx !== pivot).concat(pivot)
    for (const idx of adjustOrder) {
      const candidate = Number((items[idx].width + diff).toFixed(2))
      if (candidate >= minWidth - 0.01 && candidate <= maxWidth + 0.01) {
        items[idx].width = clampWidth(candidate, minWidth, maxWidth)
        diff = 0
        break
      }
    }

    if (Math.abs(diff) > 0.01) {
      const fallbackIdx = pivot >= 0 ? pivot : 0
      const item = items[fallbackIdx]
      item.width = clampWidth(item.width + diff, minWidth, maxWidth)
    }
  }
}

function enforceWidthBudget(changedIndex?: number) {
  if (adjustingWidths) {
    enforceWidthBudgetInternal(changedIndex)
    return
  }
  adjustingWidths = true
  try {
    enforceWidthBudgetInternal(changedIndex)
  } finally {
    adjustingWidths = false
  }
}

function getLayerStyle(layer: any, index: number): CSSProperties {
  return {
    backgroundColor: resolveBackground(layer.color),
    opacity: ((layer.opacity ?? 100) as number) / 100,
    height: `${Math.max(Math.min(layer.height ?? 80, 180), 60)}px`,
    width: `${layer.width ?? 100}%`,
    marginBottom: index < createdLayers.value.length - 1 ? '8px' : '0',
    direction: (layer.direction || 'rtl') as CSSProperties['direction']
  }
}

function getPreviewItemStyle(item: FooterItem, original: any, itemIndex: number): CSSProperties {
  const list = parseLayerItems(original)
  const width = typeof item.width === 'number' ? item.width : (list.length ? 100 / list.length : 100)
  return {
    width: `${width}%`,
    backgroundColor: 'transparent',
    borderRight: itemIndex < list.length - 1 ? '1px solid rgba(0,0,0,0.1)' : 'none'
  }
}

function getSeparatorStyle(layer: any): CSSProperties {
  const style: CSSProperties = {
    borderTopStyle: 'solid',
    borderTopColor: layer.separatorColor || '#e9ecef',
    borderTopWidth: `${layer.separatorWidth || 1}px`,
    opacity: ((layer.separatorOpacity ?? 100) as number) / 100,
    margin: '8px 0'
  }
  if (layer.separatorType === 'dashed' || layer.separatorType === 'dotted' || layer.separatorType === 'double') {
    style.borderTopStyle = layer.separatorType
  }
  return style
}

function getActiveLayerStyle(): CSSProperties {
  return {
    backgroundColor: resolveBackground(newLayer.value?.color),
    opacity: layerOpacity.value / 100,
    height: `${Math.max(layerHeight.value, 80)}px`,
    width: `${layerWidth.value}%`,
    direction: (newLayer.value?.direction || 'rtl') as CSSProperties['direction'],
    padding: '0 24px'
  }
}

function getContainerAlignment(): CSSProperties {
  if (!currentItems.value.length) {
    return { display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100%' }
  }
  const firstAlign = currentItems.value[0]?.align || 'center'
  const allSame = currentItems.value.every(item => (item.align || 'center') === firstAlign)
  if (allSame) {
    if (firstAlign === 'left') return { display: 'flex', justifyContent: 'flex-start', alignItems: 'center', height: '100%' }
    if (firstAlign === 'right') return { display: 'flex', justifyContent: 'flex-end', alignItems: 'center', height: '100%' }
    return { display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100%' }
  }
  return { display: 'flex', justifyContent: 'space-between', alignItems: 'center', height: '100%' }
}

function getItemAlignment(align: string) {
  if (align === 'left') return 'flex-start'
  if (align === 'right') return 'flex-end'
  return 'center'
}

function getItemStyle(item: FooterItem, index: number): CSSProperties {
  const width = typeof item.width === 'number' ? item.width : fallbackWidth.value
  return {
    width: `${Math.max(width, 5)}%`,
    backgroundColor: 'transparent',
    borderRight: index < currentItems.value.length - 1 ? '2px dashed rgba(55,65,81,0.35)' : 'none',
    borderLeft: index > 0 ? '2px dashed rgba(55,65,81,0.35)' : '2px dashed rgba(55,65,81,0.35)',
    borderTop: '2px dashed rgba(55,65,81,0.35)',
    borderBottom: '2px dashed rgba(55,65,81,0.35)',
    paddingRight: `${item.paddingRight || 0}px`,
    paddingLeft: `${item.paddingLeft || 0}px`,
    display: 'flex',
    alignItems: 'center',
    justifyContent: getItemAlignment(item.align || 'center'),
    cursor: 'pointer'
  }
}

function resolveItemId(item: FooterItem | string): string {
  if (typeof item === 'string') return item
  if (item?.id) return item.id
  if (item?.name) return item.name
  return 'item'
}

function getItemDisplayName(item: FooterItem | string) {
  if (typeof item === 'object' && item !== null) {
    if (item.name) return item.name
    if (item.id === 'trust' && Array.isArray(item.props?.badges)) {
      const count = item.props.badges.length
      return count ? `نشان‌های اعتماد (${count})` : 'نشان‌های اعتماد'
    }
    if (item.config && Array.isArray(item.config.menuIds)) {
      const ids = item.config.menuIds
      if (!ids.length) return 'منو (بدون انتخاب)'
      if (ids.length === 1) return `منو: ${ids[0]}`
      return `منو (${ids.length} مورد)`
    }
  }
  const itemId = resolveItemId(item)
  const registry = resolvedAvailableItems.value
  const found = registry.find((entry: any) => entry.id === itemId)
  return found?.name || itemId
}

function getItemIcon(item: FooterItem | string) {
  const id = resolveItemId(item)
  switch (id) {
    case 'logo':
      return '🏢'
    case 'image':
      return '🖼️'
    case 'custom':
      return '🧩'
    case 'menu':
      return '☰'
    case 'language':
      return '🌐'
    case 'currency':
      return '💰'
    case 'social':
      return '📱'
    case 'contact':
      return '☎️'
    case 'about':
      return 'ℹ️'
    case 'working-hours':
      return '🕐'
    case 'newsletter':
      return '✉️'
    case 'trust':
      return '🛡️'
    case 'links':
      return '🔗'
    case 'back-to-top':
      return '⬆️'
    default:
      return '📦'
  }
}

function toggleItemSettings(index: number) {
  if (activeIndex.value === index) {
    activeIndex.value = null
    return
  }
  activeIndex.value = index
  const item = currentItems.value[index]
  if (item) {
    ensureItemDefaults(item)
  }
}

// Drag & Drop functions
function handleDragStart(index: number, event: DragEvent) {
  draggingIndex.value = index
  if (event.dataTransfer) {
    event.dataTransfer.effectAllowed = 'move'
    event.dataTransfer.setData('text/plain', index.toString())
  }
  // اضافه کردن کلاس به body برای نشان دادن حالت drag
  document.body.classList.add('is-dragging-item')
}

function handleDragEnd() {
  draggingIndex.value = null
  dragOverIndex.value = null
  document.body.classList.remove('is-dragging-item')
}

function handleDragOver(index: number, event: DragEvent) {
  if (draggingIndex.value === null || draggingIndex.value === index) return
  dragOverIndex.value = index
}

function handleDrop(targetIndex: number, event: DragEvent) {
  event.preventDefault()
  
  if (draggingIndex.value === null || draggingIndex.value === targetIndex) {
    handleDragEnd()
    return
  }
  
  const items = [...currentItems.value]
  const draggedItem = items[draggingIndex.value]
  
  // حذف آیتم از موقعیت قبلی
  items.splice(draggingIndex.value, 1)
  
  // اضافه کردن آیتم به موقعیت جدید
  const newIndex = draggingIndex.value < targetIndex ? targetIndex - 1 : targetIndex
  items.splice(newIndex, 0, draggedItem)
  
  // آپدیت لیست آیتم‌ها
  if (newLayer.value) {
    newLayer.value.items = items
  }
  
  // اگر آیتم فعلی در حال ویرایش بود، ایندکس آن را آپدیت کن
  if (activeIndex.value === draggingIndex.value) {
    activeIndex.value = newIndex
  } else if (activeIndex.value !== null) {
    if (draggingIndex.value < activeIndex.value && newIndex >= activeIndex.value) {
      activeIndex.value--
    } else if (draggingIndex.value > activeIndex.value && newIndex <= activeIndex.value) {
      activeIndex.value++
    }
  }
  
  handleDragEnd()
}

function ensureWidthsSnapshot() {
  initialWidths.value = currentItems.value.map(item => (typeof item.width === 'number' ? item.width : fallbackWidth.value))
}

function startResize(itemIndex: number, event: MouseEvent | TouchEvent) {
  if (!previewItemsContainer.value || !currentItems.value[itemIndex + 1]) return
  isResizing.value = true
  resizeItemIndex.value = itemIndex
  const clientX = 'touches' in event ? event.touches[0].clientX : event.clientX
  resizeStartX.value = clientX
  containerWidth.value = previewItemsContainer.value.offsetWidth || 1
  ensureWidthsSnapshot()
  document.addEventListener('mousemove', handleResize, { passive: false })
  document.addEventListener('mouseup', stopResize)
  document.addEventListener('touchmove', handleResize, { passive: false })
  document.addEventListener('touchend', stopResize)
  document.body.style.cursor = 'col-resize'
  document.body.style.userSelect = 'none'
}

function handleResize(event: MouseEvent | TouchEvent) {
  if (!isResizing.value) return
  const nextIndex = resizeItemIndex.value + 1
  const current = currentItems.value[resizeItemIndex.value]
  const next = currentItems.value[nextIndex]
  if (!current || !next) return
  const clientX = 'touches' in event ? event.touches[0].clientX : event.clientX
  const deltaX = clientX - resizeStartX.value
  const deltaPercent = (deltaX / (containerWidth.value || 1)) * 100
  let currentWidth = initialWidths.value[resizeItemIndex.value] + deltaPercent
  let nextWidth = initialWidths.value[nextIndex] - deltaPercent
  const minWidth = getEffectiveMinWidth(currentItems.value.length)
  const maxWidth = getMaxWidthPerItem(currentItems.value.length, minWidth)
  currentWidth = Math.max(minWidth, Math.min(maxWidth, currentWidth))
  nextWidth = Math.max(minWidth, Math.min(maxWidth, nextWidth))
  const originalSum = initialWidths.value[resizeItemIndex.value] + initialWidths.value[nextIndex]
  const adjustedSum = currentWidth + nextWidth
  const diff = originalSum - adjustedSum
  if (Math.abs(diff) > 0.001) {
    currentWidth += diff / 2
    nextWidth += diff / 2
  }
  current.width = Number(currentWidth.toFixed(2))
  next.width = Number(nextWidth.toFixed(2))
  enforceWidthBudget(resizeItemIndex.value)
  event.preventDefault()
}

function stopResize() {
  if (!isResizing.value) return
  isResizing.value = false
  resizeItemIndex.value = -1
  document.removeEventListener('mousemove', handleResize)
  document.removeEventListener('mouseup', stopResize)
  document.removeEventListener('touchmove', handleResize)
  document.removeEventListener('touchend', stopResize)
  document.body.style.cursor = ''
  document.body.style.userSelect = ''
  enforceWidthBudget()
}

onBeforeUnmount(() => {
  stopResize()
})

function normalizeWidth(index: number | null) {
  if (index === null) return
  const item = currentItems.value[index]
  if (!item) return
  if (typeof item.width !== 'number' || Number.isNaN(item.width)) {
    item.width = fallbackWidth.value
  }
  const minWidth = getEffectiveMinWidth(currentItems.value.length)
  const maxWidth = getMaxWidthPerItem(currentItems.value.length, minWidth)
  item.width = clampWidth(item.width, minWidth, maxWidth)
  enforceWidthBudget(index)
}

function resetItemWidths() {
  if (!currentItems.value.length) return
  const equalWidth = Number((WIDTH_LIMIT / currentItems.value.length).toFixed(2))
  currentItems.value.forEach(item => {
    item.width = equalWidth
  })
  enforceWidthBudget()
}

function removeItem(index: number | null) {
  if (index === null) return
  currentItems.value.splice(index, 1)
  activeIndex.value = null
  if (!currentItems.value.length) return
  const equalWidth = Number((WIDTH_LIMIT / currentItems.value.length).toFixed(2))
  currentItems.value.forEach(item => {
    if (typeof item.width !== 'number' || Number.isNaN(item.width)) {
      item.width = equalWidth
    }
  })
  enforceWidthBudget()
}

function openMediaLibrary() {
  showMediaLibrary.value = true
}

function onMediaSelected(mediaList: any[]) {
  if (!mediaList.length || activeIndex.value === null) {
    showMediaLibrary.value = false
    return
  }
  const media = mediaList[0]
  const item = currentItems.value[activeIndex.value]
  if (item && item.id === 'image') {
    item.imageUrl = media.url
    item.imageName = media.name
    item.imageId = media.id
  }
  showMediaLibrary.value = false
}

function removeSelectedImage() {
  if (activeIndex.value === null) return
  const item = currentItems.value[activeIndex.value]
  if (item && item.id === 'image') {
    delete item.imageUrl
    delete item.imageName
    delete item.imageId
  }
}

function closeActiveSettings() {
  activeIndex.value = null
}

// Social media functions
const isActiveItemSocial = computed(() => {
  if (activeIndex.value === null) return false
  const item = currentItems.value[activeIndex.value]
  return item?.id === 'social'
})

const activeSocials = computed<SocialEntry[]>({
  get() {
    if (activeIndex.value === null) return []
    const item = currentItems.value[activeIndex.value]
    if (!item || item.id !== 'social') return []
    if (!item.props) item.props = {}
    if (!Array.isArray(item.props.socials)) {
      item.props.socials = []
    }
    return item.props.socials
  },
  set(newVal: SocialEntry[]) {
    if (activeIndex.value === null) return
    const item = currentItems.value[activeIndex.value]
    if (!item || item.id !== 'social') return
    if (!item.props) item.props = {}
    item.props.socials = newVal
  }
})

function handleSocialPlatformChange(social: SocialEntry, index: number) {
  const preset = SOCIAL_PRESETS.find(p => p.platform === social.platform)
  if (preset) {
    social.label = preset.label
  }
}

function addActiveSocial() {
  if (activeIndex.value === null) return
  const item = currentItems.value[activeIndex.value]
  if (!item || item.id !== 'social') return
  if (!item.props) item.props = {}
  if (!Array.isArray(item.props.socials)) {
    item.props.socials = []
  }
  item.props.socials.push({
    id: `social-${Date.now()}`,
    platform: '',
    label: '',
    url: '',
    enabled: true,
    openInNewTab: true
  })
}

function removeActiveSocial(index: number) {
  if (activeIndex.value === null) return
  const item = currentItems.value[activeIndex.value]
  if (!item || item.id !== 'social') return
  if (!item.props || !Array.isArray(item.props.socials)) return
  item.props.socials.splice(index, 1)
}

function resetActiveSocials() {
  if (activeIndex.value === null) return
  const item = currentItems.value[activeIndex.value]
  if (!item || item.id !== 'social') return
  if (!item.props) item.props = {}
  item.props.socials = SOCIAL_PRESETS.map(preset => ({
    id: `social-${preset.platform}`,
    platform: preset.platform,
    label: preset.label,
    url: '',
    enabled: true,
    openInNewTab: true
  }))
}
</script>

<style scoped>
.preview-section {
  background: #fff;
  border-radius: 10px;
  box-shadow: 0 2px 10px rgba(15, 23, 42, 0.08);
  padding: 24px;
  margin-bottom: 24px;
  width: 100%;
}

.preview-section.full-width {
  width: 100%;
  max-width: none;
}

.preview-section h3 {
  font-size: 18px;
  font-weight: 600;
  color: #1f2937;
  margin-bottom: 16px;
}

.footer-preview-actions {
  display: flex;
  justify-content: flex-end;
  margin-bottom: 16px;
}

.btn-outline {
  background: #f3f4f6;
  border: 1px solid #d1d5db;
  color: #374151;
  padding: 8px 16px;
  border-radius: 6px;
  font-size: 14px;
  cursor: pointer;
  transition: background-color 0.2s, border-color 0.2s;
}

.btn-outline:hover {
  background: #e5e7eb;
  border-color: #9ca3af;
}

.items-selector {
  display: flex;
  align-items: center;
  gap: 8px;
}

.footer-preview {
  border: 2px dashed #d1d5db;
  border-radius: 8px;
  padding: 24px;
  min-height: 220px;
  background: #f9fafb;
}

.created-layers-preview {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.preview-layer {
  border-radius: 8px;
  border: 1px solid rgba(0, 0, 0, 0.08);
  padding: 16px;
  transition: box-shadow 0.2s;
}

.preview-layer:hover {
  box-shadow: 0 6px 18px rgba(15, 23, 42, 0.08);
}

.layer-name-preview {
  font-size: 12px;
  font-weight: 600;
  color: #4b5563;
  margin-bottom: 8px;
}

.layer-items-preview-horizontal {
  display: flex;
  min-height: 44px;
  border-radius: 6px;
  overflow: hidden;
}

.layer-item-preview-box {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 8px;
  font-size: 12px;
  color: #1f2937;
  text-align: center;
}

.item-preview-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 2px;
}

.item-preview-icon {
  font-size: 16px;
}

.item-preview-name {
  font-weight: 500;
  color: #111827;
}

.item-preview-width {
  font-size: 10px;
  color: #6b7280;
}

.preview-separator-horizontal {
  width: 100%;
}

.preview-footer {
  margin-top: 24px;
}

.preview-footer-label {
  font-size: 14px;
  color: #374151;
  margin-bottom: 12px;
}

.preview-container {
  border: 1px dashed rgba(59, 130, 246, 0.4);
  border-radius: 10px;
  background: #f3f4f6;
  padding: 12px 24px;
}

.preview-items-container {
  display: flex;
  align-items: stretch;
  height: 100%;
  min-height: 64px;
  gap: 0;
}

.preview-section-item {
  min-height: 56px;
  border-radius: 8px;
  background: #e5e7eb;
  transition: background-color 0.2s, box-shadow 0.2s, transform 0.2s;
  cursor: grab;
  position: relative;
}

.preview-section-item:active {
  cursor: grabbing;
}

.preview-section-item.is-active {
  background: #dbeafe;
  box-shadow: inset 0 0 0 2px #3b82f6;
}

.preview-section-item.is-dragging {
  opacity: 0.5;
  transform: scale(0.95);
  cursor: grabbing;
}

.preview-section-item.drag-over {
  background: #bfdbfe;
  box-shadow: inset 0 0 0 2px #60a5fa;
}

.drag-handle {
  position: absolute;
  left: 8px;
  top: 50%;
  transform: translateY(-50%);
  font-size: 16px;
  color: #9ca3af;
  cursor: grab;
  padding: 4px;
  user-select: none;
}

.drag-handle:active {
  cursor: grabbing;
}

.item-preview-content {
  padding-left: 28px;
}

body.is-dragging-item {
  cursor: grabbing !important;
  user-select: none;
}

.resize-handle {
  width: 10px;
  cursor: col-resize;
  background: repeating-linear-gradient(
    90deg,
    rgba(148, 163, 184, 0.4),
    rgba(148, 163, 184, 0.4) 2px,
    transparent 2px,
    transparent 4px
  );
  margin: 0 4px;
  border-radius: 4px;
}

.item-settings-panel {
  margin-top: 16px;
  border: 1px solid #d1d5db;
  border-radius: 8px;
  background: #ffffff;
  padding: 16px;
  box-shadow: 0 6px 16px rgba(15, 23, 42, 0.06);
}

.settings-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.settings-header h4 {
  margin: 0;
  font-size: 14px;
  font-weight: 600;
  color: #1f2937;
}

.close-settings-btn {
  background: none;
  border: none;
  font-size: 16px;
  cursor: pointer;
  color: #6b7280;
}

.settings-content {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.settings-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(160px, 1fr));
  gap: 16px;
}

.form-field {
  display: flex;
  flex-direction: column;
  gap: 6px;
  font-size: 12px;
  color: #374151;
}

.form-field--full {
  grid-column: 1 / -1;
}

.form-field label {
  font-weight: 600;
  color: #111827;
}

.form-control {
  width: 100%;
  padding: 8px 10px;
  border: 1px solid #d1d5db;
  border-radius: 6px;
  font-size: 13px;
  background: #fff;
  transition: border-color 0.2s, box-shadow 0.2s;
}

.form-control:focus {
  border-color: #3b82f6;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.2);
  outline: none;
}

.rich-text-editor-wrapper {
  width: 100%;
}

.editor-loading {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 16px;
  font-size: 13px;
  color: #6b7280;
  background: #f9fafb;
  border: 1px dashed #d1d5db;
  border-radius: 8px;
}

.remove-item-btn {
  align-self: flex-start;
  background: #fee2e2;
  border: 1px solid #ef4444;
  color: #b91c1c;
  padding: 6px 12px;
  border-radius: 6px;
  font-size: 12px;
  cursor: pointer;
  transition: background-color 0.2s;
}

.remove-item-btn:hover {
  background: #fecaca;
}

.layer-info {
  margin-top: 20px;
  border-top: 1px dashed #d1d5db;
  padding-top: 12px;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.info-row {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  align-items: center;
}

.info-capsule {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  justify-content: center;
  gap: 4px;
  min-width: 120px;
  padding: 8px 12px;
  border-radius: 8px;
  background: #f9fafb;
  border: 1px solid #e5e7eb;
  font-size: 12px;
  color: #374151;
}

.info-capsule.emphasis {
  background: rgba(34, 197, 94, 0.08);
  border-color: rgba(34, 197, 94, 0.4);
}

.info-label {
  font-weight: 600;
  color: #111827;
}

.info-value {
  font-size: 13px;
  color: #1f2937;
}

.total-width.valid {
  color: #15803d;
}

.total-width.invalid {
  color: #b91c1c;
}

.capsule-action {
  height: 40px;
  padding: 0 14px;
  display: flex;
  align-items: center;
}

.width-items-inline {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.width-items-inline.compact {
  margin-top: 4px;
}

.width-chip {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 6px 10px;
  border-radius: 999px;
  background: #eef2ff;
  color: #1d4ed8;
  font-size: 11px;
}

.chip-name {
  font-weight: 600;
}

.chip-value {
  font-family: 'Fira Code', monospace;
  direction: ltr;
}

.empty-preview {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 200px;
}

.empty-message {
  text-align: center;
  color: #6b7280;
}

.image-upload-section {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.image-upload-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.selected-image-preview {
  display: flex;
  align-items: center;
  gap: 12px;
}

.selected-image-preview img {
  width: 64px;
  height: 64px;
  object-fit: cover;
  border-radius: 6px;
}

.image-info {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 12px;
  color: #374151;
}

.remove-image-btn {
  background: none;
  border: none;
  font-size: 14px;
  cursor: pointer;
  color: #ef4444;
}

.image-upload-guide {
  font-size: 12px;
  color: #6b7280;
}

@media (max-width: 768px) {
  .preview-container {
    padding: 12px;
  }
  .preview-items-container {
    flex-direction: column;
    gap: 8px;
  }
  .preview-section-item {
    width: 100% !important;
    border: 2px dashed rgba(55, 65, 81, 0.35);
  }
  .resize-handle {
    display: none;
  }
}
</style>


