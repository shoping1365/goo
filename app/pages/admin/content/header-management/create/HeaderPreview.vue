<template>
  <div class="preview-section" :class="{ 'full-width': isEditing }">
    <div class="header-info">
      <div class="header-form-row">
        
      </div>
    </div>
    
    <h3>پیش نمایش هدر</h3>
    <div class="header-preview-actions" style="display: flex; justify-content: flex-end; margin-bottom: 12px;">
      <button type="button" class="btn btn-outline items-selector" @click="openItemsModal">
        <span>{{ getSelectedItemsText() }}</span>
        <span>📋</span>
      </button>
    </div>
    <div class="header-preview">
      <!-- نمایش لایه‌های ایجاد شده -->
      <div v-if="createdLayers && createdLayers.length > 0" class="created-layers-preview">
        
        <div 
          v-for="(layer, layerIndex) in createdLayers" 
          :key="layer.id"
          class="preview-layer"
          :style="getPreviewLayerStyle(layer, layerIndex)"
        >
          <div class="layer-name-preview">{{ layer.name || 'لایه بدون نام' }}</div>
          
          <!-- نمایش ایتم‌های لایه به صورت افقی -->
          <div class="layer-items-preview-horizontal">
            <div 
              v-for="(itemObj, itemIndex) in (typeof layer.items === 'string' ? JSON.parse(layer.items) : layer.items)" 
              :key="itemIndex"
              class="layer-item-preview-box"
              :style="{
                width: `${itemObj.width || (100 / (typeof layer.items === 'string' ? JSON.parse(layer.items).length : layer.items.length))}%`,
                backgroundColor: 'transparent',
                borderRight: itemIndex < (typeof layer.items === 'string' ? JSON.parse(layer.items).length : layer.items.length) - 1 ? '1px solid rgba(0,0,0,0.1)' : 'none'
              }"
            >
              <div class="item-preview-content">
                <div class="item-preview-name">
                  {{ getItemDisplayName(typeof itemObj === 'string' ? itemObj : (itemObj.id || 'unknown')) }}
                </div>
                <div class="item-preview-width">
                  {{ (itemObj.width || (100 / (typeof layer.items === 'string' ? JSON.parse(layer.items).length : layer.items.length))).toFixed(1) }}%
                </div>
              </div>
            </div>
          </div>
          
          <!-- جداکننده بین لایه‌ها -->
          <div 
            v-if="layer.showSeparator && layerIndex < createdLayers.length - 1" 
            class="preview-separator-horizontal"
            :style="{
              borderTopStyle: layer.separatorType === 'dashed' ? 'dashed' : 
                             layer.separatorType === 'dotted' ? 'dotted' : 
                             layer.separatorType === 'double' ? 'double' : 'solid',
              borderTopColor: layer.separatorColor || '#e9ecef',
              borderTopWidth: (layer.separatorWidth || 1) + 'px',
              opacity: (layer.separatorOpacity || 100) / 100,
              margin: '8px 0'
            }"
          ></div>
        </div>
      </div>
      

      
      <!-- لایه در حال ویرایش یا ایجاد جدید -->
      <div v-if="showLayerSettings" class="preview-header">
        <div class="preview-header-label">
          <strong>پیش‌نمایش زنده لایه:</strong>
        </div>
        <div 
          class="preview-container"
          :style="getActiveLayerStyle()"
        >
          <!-- نمایش ایتم‌های انتخاب شده با قابلیت درگ و دراپ -->
          <div class="preview-items-container" style="display: flex; align-items: center; justify-content: flex-start; gap: 0px;">
            <template v-for="(itemObj, itemIndex) in newLayer.items" :key="itemIndex">
              <div 
                class="preview-section-item"
                :data-align="itemObj.align || 'center'"
                :style="{
                  width: `${Math.max(itemObj.width || 20, 15)}%`,
                  backgroundColor: 'transparent',
                  borderRight: itemIndex < newLayer.items.length - 1 ? '2px dashed #999' : 'none',
                  borderTop: '2px dashed #999',
                  borderBottom: '2px dashed #999',
                  borderLeft: itemIndex > 0 ? '2px dashed #999' : 'none',
                  boxSizing: 'border-box',
                  flex: '0 0 auto',
                  overflow: 'hidden',
                  position: 'relative',
                  paddingRight: `${itemObj.paddingRight || 0}px`,
                  paddingLeft: `${itemObj.paddingLeft || 0}px`
                }"
                @click="toggleItemSettings(itemIndex)"
              >
                <div class="text-center text-xs mb-1 font-medium">
                  {{ getItemDisplayName(typeof itemObj === 'string' ? itemObj : (itemObj.id || 'unknown')) }}
                </div>
                <div class="text-[10px] text-gray-500 mb-1">
                  {{ itemObj.align || 'center' }} / 
                  {{ itemObj.width || (100 / newLayer.items.length) }}%
                </div>
                
                <!-- نمایش آیکون‌های آیتم با چینش -->
                <div class="item-icon-container" :style="getIconContainerStyle(itemObj.align || 'center')">
                  <div v-if="hasItem(itemObj, 'logo')" class="item-icon">🏢</div>
                  <div v-else-if="hasItem(itemObj, 'search')" class="item-icon">🔍</div>
                  <div v-else-if="hasItem(itemObj, 'auth')" class="item-icon">👤</div>
                  <div v-else-if="hasItem(itemObj, 'cart')" class="item-icon">🛒</div>
                  <div v-else-if="hasItem(itemObj, 'notification')" class="item-icon">🔔</div>
                  <div v-else-if="hasItem(itemObj, 'custom')" class="item-icon">📝</div>
                  <div v-else-if="hasItem(itemObj, 'menu')" class="item-icon">☰</div>
                  <div v-else-if="hasItem(itemObj, 'language')" class="item-icon">🌐</div>
                  <div v-else-if="hasItem(itemObj, 'currency')" class="item-icon">💰</div>
                  <div v-else-if="hasItem(itemObj, 'social')" class="item-icon">📱</div>
                  <div v-else class="item-icon">📦</div>
                </div>
              </div>
              
              <!-- دستگیره تغییر اندازه -->
              <div 
                v-if="itemIndex < newLayer.items.length - 1"
                class="resize-handle"
                @mousedown="startResize(itemIndex, $event)"
                @touchstart="startResize(itemIndex, $event)"
                title="کشیدن برای تغییر اندازه"
              ></div>
            </template>
          </div>
        </div>
        
        <!-- تنظیمات آیتم انتخاب‌شده -->
        <div v-if="activeIndex !== null && newLayer.items[activeIndex]" class="item-settings-panel mt-4">
          <div class="settings-header">
            <h4>{{ getItemDisplayName(newLayer.items[activeIndex].id) }}</h4>
            <button class="close-settings-btn" @click="activeIndex = null">✕</button>
          </div>
          
          <div class="settings-content">
            <!-- اگر آیتم عکس است، دکمه آپلود عکس نمایش داده شود -->
            <div v-if="newLayer.items[activeIndex].id === 'image'" class="image-upload-section mb-4">
              <div class="image-upload-header">
                <h5>آپلود عکس</h5>
                <button 
                  type="button"
                  class="btn btn-primary btn-sm"
                  @click="openMediaLibrary"
                >
                  📷 انتخاب از کتابخانه
                </button>
              </div>
              
              <!-- نمایش عکس انتخاب شده -->
              <div v-if="newLayer.items[activeIndex].imageUrl" class="selected-image-preview">
                <img 
                  :src="newLayer.items[activeIndex].imageUrl" 
                  :alt="newLayer.items[activeIndex].imageName || 'عکس انتخاب شده'"
                  class="preview-image"
                />
                <div class="image-info">
                  <span class="image-name">{{ newLayer.items[activeIndex].imageName || 'عکس انتخاب شده' }}</span>
                  <button 
                    @click="removeSelectedImage" 
                    class="remove-image-btn"
                    title="حذف عکس"
                  >
                    ✕
                  </button>
                </div>
              </div>
              
              <!-- پیام راهنما -->
              <div v-else class="image-upload-guide">
                <p class="text-sm text-gray-600">
                  برای نمایش عکس در هدر، ابتدا یک عکس آپلود کنید.
                </p>
              </div>
            </div>
            
            <div class="settings-row-inline">
              <label class="settings-label-inline">
                <span>چینش:</span>
                <select v-model="newLayer.items[activeIndex].align" class="settings-select-small">
                  <option value="left">راست</option>
                  <option value="center">وسط</option>
                  <option value="right">چپ</option>
                </select>
              </label>
              
              <label class="settings-label-inline">
                <span>عرض:</span>
                <input 
                  type="number" 
                  v-model.number="newLayer.items[activeIndex].width"
                  min="5" 
                  max="95" 
                  step="0.1"
                  class="settings-input-small"
                />
                <span class="settings-unit">%</span>
              </label>
            </div>
            
            <!-- پیام راهنما -->
            <div class="width-warning-compact">
              <small style="color: #6c757d; font-size: 11px;">
                💡 برای تغییر عرض آیتم‌ها، عدد را در فیلد عرض تغییر دهید.
              </small>
            </div>
            
            <!-- فیلدهای جدید پدینگ -->
            <div class="settings-row-inline">
              <label class="settings-label-inline">
                <span>پدینگ راست:</span>
                <input 
                  type="number" 
                  v-model.number="newLayer.items[activeIndex].paddingRight" 
                  min="0" 
                  max="200" 
                  class="settings-input-small" 
                />
                <span class="unit">px</span>
              </label>
              
              <label class="settings-label-inline">
                <span>پدینگ چپ:</span>
                <input 
                  type="number" 
                  v-model.number="newLayer.items[activeIndex].paddingLeft" 
                  min="0" 
                  max="200" 
                  class="settings-input-small" 
                />
                <span class="unit">px</span>
              </label>
            </div>
            
            <div class="settings-row-inline">
              <button class="remove-item-btn-inline" @click="removeItem(activeIndex)">
                حذف آیتم
              </button>
            </div>
          </div>
        </div>
        
        <!-- کتابخانه رسانه -->
        <MediaLibraryModal 
          v-model="showMediaLibrary"
          :model-selected="selectedMediaIds"
          default-category="banners"
          file-type="image"
          context-title="هدر - انتخاب عکس"
          @confirm="onMediaSelected"
        />
      </div>
      
      <!-- نمایش اطلاعات لایه -->
        <div v-if="showLayerSettings" class="layer-info">
          <div class="info-row">
            <span class="info-item">نام لایه: {{ newLayer.name || 'بدون نام' }}</span>
            <span class="info-item">عرض: {{ newLayer.width }}%</span>
            <span class="info-item">ارتفاع: {{ newLayer.height }}px</span>
            <span class="info-item">شفافیت: {{ newLayer.opacity }}%</span>
            <span class="info-item">تعداد ردیف: {{ newLayer.items.length }}</span>
            <span class="info-item">ایتم‌ها: {{ newLayer.items.length }} مورد انتخاب شده</span>
            <span v-if="newLayer.id" class="info-item">وضعیت: <span style="color: #28a745;">در حال ویرایش</span></span>
          </div>
          
          <!-- نمایش مجموع عرض‌ها -->
          <div v-if="newLayer.items && newLayer.items.length > 0" class="width-summary">
            <div class="width-summary-header">
              <strong>مجموع عرض‌ها:</strong>
              <span class="total-width" :class="{ 'valid': isTotalWidthValid, 'invalid': !isTotalWidthValid }">
                {{ totalWidth.toFixed(2) }}%
              </span>
            </div>
            
            <!-- نمایش عرض‌ها و دکمه تنظیم خودکار در یک خط -->
            <div class="width-controls-row">
              <div class="width-items-inline">
                <div v-for="(item, index) in newLayer.items" :key="index" class="width-item-inline">
                  <span class="item-name-small">{{ getItemDisplayName(item.id) }}:</span>
                  <span class="item-width-small">{{ item.width.toFixed(2) }}%</span>
                </div>
              </div>
              
              <!-- دکمه تنظیم خودکار عرض‌ها -->
              <button 
                class="btn btn-sm btn-outline reset-widths-btn" 
                @click="resetItemWidths"
                :disabled="isTotalWidthValid"
              >
                تنظیم خودکار
              </button>
            </div>
          </div>
        </div>
    </div>

  </div>
</template>

<script setup lang="ts">
import { inject, ref, computed, watch, onUnmounted, nextTick } from 'vue'
import MediaLibraryModal from '~/components/media/MediaLibraryModal.vue'
// حذف Drawer


// Inject data from parent
const headerData = inject('headerData')
const newLayer = inject<Ref<any>>('newLayer', ref({
  id: '',
  name: '',
  width: 100,
  height: 50,
  rowCount: 1,
  color: '#ffffff',
  opacity: 100,
  showSeparator: false,
  separatorType: 'line',
  separatorColor: '#e9ecef',
  separatorOpacity: 100,
  separatorWidth: 1,
  pageSelection: 'all',
  specificPages: '',
  excludedPages: '',
  items: [],
  paddingRight: 0,
  paddingLeft: 0,
  borderEnabled: false,
  elevationEnabled: false,
}))

// Debug: نمایش اطلاعات newLayer
console.log('HeaderPreview - newLayer:', newLayer.value)
console.log('HeaderPreview - newLayer.items:', newLayer.value?.items)

// Watch برای تغییرات newLayer.items
watch(() => newLayer.value?.items, (items) => {
  console.log('HeaderPreview - newLayer.items changed:', items)
  if (items && items.length > 0) {
    items.forEach((item, index) => {
      console.log(`📦 Item ${index}:`, item)
      console.log(`   - Type: ${typeof item}`)
      console.log(`   - ID: ${item.id}`)
      console.log(`   - Items: ${JSON.stringify(item.items)}`)
      console.log(`   - hasItem(logo): ${hasItem(item, 'logo')}`)
      console.log(`   - hasItem(search): ${hasItem(item, 'search')}`)
      console.log(`   - hasItem(auth): ${hasItem(item, 'auth')}`)
    })
    // ❌ غیرفعال شد: enforceWidthBudget خودکار - کاربر باید دستی تنظیم کنه
    // enforceWidthBudget(typeof activeIndex.value === 'number' ? activeIndex.value : undefined)
  }
}, { deep: true })
import type { Ref } from 'vue'
const createdLayers = inject<Ref<any[]>>('createdLayers', ref([]))

// Debug: نمایش اطلاعات createdLayers
console.log('HeaderPreview - createdLayers:', createdLayers.value)

// Watch برای تغییرات createdLayers
watch(createdLayers, (layers) => {
  console.log('HeaderPreview - createdLayers changed:', layers)
}, { deep: true })
const isEditing = inject('isEditing', ref(false))

// Debug: نمایش اطلاعات isEditing
console.log('HeaderPreview - isEditing:', isEditing.value)
const showLayerSettings = inject<Ref<boolean>>('showLayerSettings', ref(false))

// Debug: نمایش اطلاعات showLayerSettings
console.log('HeaderPreview - showLayerSettings:', showLayerSettings.value)
const openItemsModal = inject('openItemsModal', () => {})
const getSelectedItemsText = inject('getSelectedItemsText', () => 'انتخاب ایتم ها')

// ایندکس آیتم فعال برای تنظیمات
const activeIndex = ref<number | null>(null)

// متغیرهای مربوط به کتابخانه رسانه
const showMediaLibrary = ref(false)
const selectedMediaIds = ref<number[]>([])

function resolveToggle(value: unknown) {
  if (typeof value === 'boolean') return value
  if (typeof value === 'number') return value === 1
  if (typeof value === 'string') {
    const normalized = value.trim().toLowerCase()
    return normalized === 'true' || normalized === '1' || normalized === 'yes'
  }
  return false
}

function parseStyleSettings(layer: any) {
  const raw = layer?.styleSettings ?? layer?.style_settings
  if (!raw) return null
  if (typeof raw === 'string') {
    try {
      return JSON.parse(raw)
    } catch (error) {
      console.warn('Failed to parse styleSettings:', error)
      return null
    }
  }
  if (typeof raw === 'object') return raw
  return null
}

function resolveOpacity(value: unknown) {
  if (typeof value !== 'number') return 1
  if (value <= 1) return Math.min(Math.max(value, 0), 1)
  return Math.min(Math.max(value / 100, 0), 1)
}

function resolveBorderConfig(layer: any) {
  const styleSettings = parseStyleSettings(layer) || {}
  const borderSettings = styleSettings.border || {}
  const enabled = resolveToggle(layer?.borderEnabled ?? layer?.border_enabled ?? layer?.enableBorder ?? borderSettings.enabled)
  return {
    enabled,
    position: layer?.borderPosition ?? borderSettings.position ?? 'all',
    color: layer?.borderColor ?? borderSettings.color ?? '#e5e7eb',
    width: Number(layer?.borderWidth ?? borderSettings.width ?? 1) || 1,
    style: layer?.borderStyle ?? borderSettings.style ?? 'solid'
  }
}

const SHADOW_PRESETS: Record<string, Record<string, string>> = {
  sm: {
    top: '0 -3px 8px -2px rgba(0, 0, 0, 0.15), 0 -1px 4px -1px rgba(0, 0, 0, 0.1)',
    bottom: '0 3px 8px -2px rgba(0, 0, 0, 0.15), 0 1px 4px -1px rgba(0, 0, 0, 0.1)',
    both: '0 -3px 8px -2px rgba(0, 0, 0, 0.15), 0 3px 8px -2px rgba(0, 0, 0, 0.15)'
  },
  md: {
    top: '0 -6px 12px -2px rgba(0, 0, 0, 0.2), 0 -3px 6px -2px rgba(0, 0, 0, 0.12)',
    bottom: '0 6px 12px -2px rgba(0, 0, 0, 0.2), 0 3px 6px -2px rgba(0, 0, 0, 0.12)',
    both: '0 -6px 12px -2px rgba(0, 0, 0, 0.2), 0 6px 12px -2px rgba(0, 0, 0, 0.2)'
  },
  lg: {
    top: '0 -12px 20px -4px rgba(0, 0, 0, 0.25), 0 -6px 10px -3px rgba(0, 0, 0, 0.15)',
    bottom: '0 12px 20px -4px rgba(0, 0, 0, 0.25), 0 6px 10px -3px rgba(0, 0, 0, 0.15)',
    both: '0 -12px 20px -4px rgba(0, 0, 0, 0.25), 0 12px 20px -4px rgba(0, 0, 0, 0.25)'
  },
  xl: {
    top: '0 -24px 36px -8px rgba(0, 0, 0, 0.3), 0 -12px 16px -6px rgba(0, 0, 0, 0.18)',
    bottom: '0 24px 36px -8px rgba(0, 0, 0, 0.3), 0 12px 16px -6px rgba(0, 0, 0, 0.18)',
    both: '0 -24px 36px -8px rgba(0, 0, 0, 0.3), 0 24px 36px -8px rgba(0, 0, 0, 0.3)'
  }
}

function resolveShadowConfig(layer: any) {
  const styleSettings = parseStyleSettings(layer) || {}
  const shadowSettings = styleSettings.shadow || {}
  const rawIntensity = typeof layer?.shadowIntensity === 'string' ? layer.shadowIntensity : shadowSettings.intensity
  const intensity = rawIntensity && rawIntensity.trim() !== '' ? rawIntensity : 'md'
  const rawDirection = typeof layer?.shadowDirection === 'string' ? layer.shadowDirection : shadowSettings.direction
  const direction = rawDirection && rawDirection.trim() !== '' ? rawDirection : 'top'
  const enabled = resolveToggle(layer?.elevationEnabled ?? layer?.elevation_enabled ?? layer?.enableShadow ?? shadowSettings.enabled) && intensity !== 'none'

  return {
    enabled,
    intensity,
    direction
  }
}

function buildBorderStyles(config: { enabled: boolean; position: string; color: string; width: number; style: string }) {
  if (!config.enabled) return {}
  const borderValue = `${config.width}px ${config.style} ${config.color}`
  const styles: Record<string, string> = {}

  switch (config.position) {
    case 'all':
      styles.border = borderValue
      break
    case 'top':
      styles.borderTop = borderValue
      break
    case 'bottom':
      styles.borderBottom = borderValue
      break
    case 'left':
      styles.borderLeft = borderValue
      break
    case 'right':
      styles.borderRight = borderValue
      break
    case 'top-bottom':
      styles.borderTop = borderValue
      styles.borderBottom = borderValue
      break
    default:
      styles.border = borderValue
  }

  return styles
}

function resolveBoxShadow(intensity: string, direction: string) {
  const preset = SHADOW_PRESETS[intensity]
  if (!preset) return 'none'
  return preset[direction] || preset.top || 'none'
}

function getPreviewLayerStyle(layer: any, index: number) {
  const borderConfig = resolveBorderConfig(layer)
  const shadowConfig = resolveShadowConfig(layer)
  const height = typeof layer?.height === 'number' ? layer.height : 50
  const width = typeof layer?.width === 'number' ? layer.width : 100
  const base: Record<string, any> = {
    backgroundColor: layer?.color || 'transparent',
    opacity: resolveOpacity(layer?.opacity),
    height: `${height}px`,
    width: `${width}%`,
    marginBottom: index < createdLayers.value.length - 1 ? '8px' : '0',
    direction: layer?.direction || 'rtl',
    textAlign: (layer?.direction || 'rtl') === 'ltr' ? 'left' : 'right',
    borderRadius: '8px',
    border: borderConfig.enabled ? 'none' : '1px dashed rgba(203, 213, 225, 0.7)'
  }

  Object.assign(base, buildBorderStyles(borderConfig))

  if (shadowConfig.enabled) {
    base.boxShadow = resolveBoxShadow(shadowConfig.intensity, shadowConfig.direction)
    base.position = 'relative'
    base.zIndex = 5
  } else {
    base.boxShadow = 'none'
  }

  return base
}

function getActiveLayerStyle() {
  const borderConfig = resolveBorderConfig(newLayer.value)
  const shadowConfig = resolveShadowConfig(newLayer.value)
  const base: Record<string, any> = {
    backgroundColor: newLayer.value.color || 'transparent',
    opacity: layerOpacity.value / 100,
    height: `${Math.max(layerHeight.value, 80)}px`,
    width: `${layerWidth.value}%`,
    direction: newLayer.value.direction || 'rtl',
    textAlign: (newLayer.value.direction || 'rtl') === 'ltr' ? 'left' : 'right',
    padding: '0 24px',
    borderRadius: '8px',
    border: borderConfig.enabled ? 'none' : '1px dashed rgba(203, 213, 225, 0.7)'
  }

  Object.assign(base, buildBorderStyles(borderConfig))

  if (shadowConfig.enabled) {
    base.boxShadow = resolveBoxShadow(shadowConfig.intensity, shadowConfig.direction)
    base.position = 'relative'
    base.zIndex = 5
  } else {
    base.boxShadow = 'none'
  }

  return base
}

// متغیرهای مربوط به درگ و دراپ
const isResizing = ref(false)
const resizeStartX = ref(0)
const resizeItemIndex = ref(-1)
const containerWidth = ref(0)
const initialWidths = ref<number[]>([])

const WIDTH_LIMIT = 100
const BASE_MIN_WIDTH = 5
let adjustingWidths = false

const fallbackWidth = computed(() => {
  const items = Array.isArray(newLayer.value?.items) ? newLayer.value.items : []
  const count = items.length || 1
  return Number((WIDTH_LIMIT / count).toFixed(2))
})

function getEffectiveMinWidth(count: number): number {
  if (count <= 0) return BASE_MIN_WIDTH
  const minCandidate = Math.min(BASE_MIN_WIDTH, WIDTH_LIMIT / count)
  return Number(minCandidate.toFixed(2))
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
  const items = Array.isArray(newLayer.value?.items) ? newLayer.value.items : []
  const count = items.length
  if (!count) return

  const minWidth = getEffectiveMinWidth(count)
  const maxWidth = getMaxWidthPerItem(count, minWidth)

  if (count === 1) {
    const sole = items[0]
    if (sole) sole.width = WIDTH_LIMIT
    return
  }

  // فقط clamp کردن عرض‌ها بین min و max - بدون تنظیم خودکار
  items.forEach(item => {
    if (!item) return
    const requested = typeof item.width === 'number' ? item.width : fallbackWidth.value
    item.width = clampWidth(requested, minWidth, maxWidth)
  })

  // فقط اگر مجموع بیشتر از 100% شد، کم کنیم
  let total = Number(items.reduce((sum, item) => sum + (item?.width || 0), 0).toFixed(2))

  if (total > WIDTH_LIMIT) {
    let excess = Number((total - WIDTH_LIMIT).toFixed(2))
    // فقط آیتمی که تغییر کرده رو کم می‌کنیم
    if (typeof changedIndex === 'number' && changedIndex >= 0 && changedIndex < count) {
      const item = items[changedIndex]
      if (item) {
        const reducible = Number(((item.width ?? minWidth) - minWidth).toFixed(2))
        if (reducible > 0) {
          const delta = Math.min(reducible, excess)
          item.width = Number(((item.width ?? minWidth) - delta).toFixed(2))
        }
      }
    }
  }
  
  // ❌ حذف شد: اضافه کردن خودکار عرض به آیتم‌های دیگر
  // هیچ تنظیم خودکار برای رسیدن به 100% انجام نمیشه
}

function enforceWidthBudget(changedIndex?: number) {
  if (adjustingWidths) return
  adjustingWidths = true
  try {
    enforceWidthBudgetInternal(changedIndex)
  } finally {
    adjustingWidths = false
  }
}

nextTick(() => {
  if (Array.isArray(newLayer.value?.items) && newLayer.value.items.length) {
    enforceWidthBudget()
  }
})

function toggleItemSettings(idx: number) {
  activeIndex.value = activeIndex.value === idx ? null : idx
}

// تابع اعتبارسنجی عرض (حذف شد - آیتم‌های دیگر تغییر نمی‌کنند)

// توابع مربوط به درگ و دراپ
function startResize(itemIndex: number, event: MouseEvent | TouchEvent) {
  console.log('startResize called!', itemIndex)
  event.preventDefault()
  event.stopPropagation()
  
  isResizing.value = true
  resizeItemIndex.value = itemIndex
  
  // دریافت موقعیت شروع
  const clientX = 'touches' in event ? event.touches[0].clientX : event.clientX
  resizeStartX.value = clientX
  
  // دریافت عرض کانتینر
  const container = document.querySelector('.preview-items-container') as HTMLElement
  if (container) {
    containerWidth.value = container.offsetWidth
  }
  
  // ذخیره عرض‌های اولیه
  initialWidths.value = newLayer.value.items.map((item: any) => {
    return item.width || (100 / newLayer.value.items.length)
  })
  
  // اضافه کردن event listeners
  document.addEventListener('mousemove', handleResize, { passive: false })
  document.addEventListener('mouseup', stopResize)
  document.addEventListener('touchmove', handleResize, { passive: false })
  document.addEventListener('touchend', stopResize)
  
  // تغییر cursor
  document.body.style.cursor = 'col-resize'
  document.body.style.userSelect = 'none'
  
  console.log('Cursor changed to col-resize')
}

function handleResize(event: MouseEvent | TouchEvent) {
  if (!isResizing.value) return
  
  const clientX = 'touches' in event ? event.touches[0].clientX : event.clientX
  const deltaX = clientX - resizeStartX.value
  
  // محاسبه درصد تغییر بر اساس عرض کانتینر
  const deltaPercent = (deltaX / containerWidth.value) * 100
  
  // محاسبه عرض‌های جدید
  const itemIndex = resizeItemIndex.value
  const currentWidth = initialWidths.value[itemIndex]
  const nextWidth = initialWidths.value[itemIndex + 1]
  
  // محاسبه عرض‌های جدید
  let newCurrentWidth = currentWidth + deltaPercent
  let newNextWidth = nextWidth - deltaPercent
  
  // محدودیت‌های حداقل و حداکثر
  newCurrentWidth = Math.max(5, Math.min(95, newCurrentWidth))
  newNextWidth = Math.max(5, Math.min(95, newNextWidth))
  
  // بررسی اینکه مجموع از 100% بیشتر نشود
  newLayer.value.items[itemIndex].width = Math.round(newCurrentWidth * 100) / 100
  newLayer.value.items[itemIndex + 1].width = Math.round(newNextWidth * 100) / 100
  enforceWidthBudget(itemIndex)
}

function stopResize() {
  if (!isResizing.value) return
  isResizing.value = false
  resizeItemIndex.value = -1
  
  // حذف event listeners
  document.removeEventListener('mousemove', handleResize)
  document.removeEventListener('mouseup', stopResize)
  document.removeEventListener('touchmove', handleResize)
  document.removeEventListener('touchend', stopResize)
  
  // بازگرداندن cursor و user-select
  document.body.style.cursor = ''
  document.body.style.userSelect = ''
  
  console.log('Cursor reset')
  enforceWidthBudget()
}



// تابع نمایش نام فارسی ایتم‌ها
function getItemDisplayName(itemId: string): string {
  const itemNames: { [key: string]: string } = {
    'logo': 'لوگو',
    'search': 'جستجو',
    'auth': 'ورود/ثبت‌نام',
    'cart': 'سبد خرید',
    'notification': 'نوتیفیکیشن',
    'custom': 'باکس سفارشی',
    'menu': 'منو',
    'language': 'زبان',
    'currency': 'ارز',
    'social': 'شبکه‌های اجتماعی'
  }
  return itemNames[itemId] || itemId
}

function hasItem(obj: any, id: string) {
  if (!obj) return false
  
  console.log('hasItem called with:', { obj, id })
  
  // اگر obj خودش آیتم مورد نظر است
  if (obj.id === id) {
    console.log('obj.id === id:', true)
    return true
  }
  
  // اگر obj دارای items array است
  if (Array.isArray(obj.items)) {
    const result = obj.items.includes(id)
    console.log('obj.items.includes(id):', result, 'items:', obj.items)
    return result
  }
  
  // اگر obj دارای items array است و id در آن وجود دارد
  if (Array.isArray(obj.items) && obj.items.includes(id)) {
    console.log('obj.items.includes(id):', true, 'items:', obj.items)
    return true
  }
  
  // اگر obj یک string است
  if (typeof obj === 'string' && obj === id) {
    console.log('obj === id (string):', true)
    return true
  }
  
  // اگر obj یک object با properties مختلف است
  if (typeof obj === 'object' && obj !== null) {
    console.log('obj is object, checking properties:', Object.keys(obj))
    // بررسی اینکه آیا obj دارای property مورد نظر است
    if (obj[id] !== undefined) {
      console.log(`obj[${id}] exists:`, obj[id])
      return true
    }
  }
  
  console.log('hasItem returning false')
  return false
}





// Computed properties for reactive styling
const layerBackgroundColor = computed(() => newLayer.value.color || '#ffffff')
const layerOpacity = computed(() => newLayer.value.opacity || 100)
const layerHeight = computed(() => newLayer.value.height || 50)
const layerWidth = computed(() => newLayer.value.width || 100)
const layerDirection = computed(() => newLayer.value.direction || 'rtl')
const layerMobileResponsive = computed(() => newLayer.value.mobileResponsive !== false)
const layerTabletResponsive = computed(() => newLayer.value.tabletResponsive !== false)

// Computed properties for width management
const totalWidth = computed(() => {
  if (!newLayer.value.items || newLayer.value.items.length === 0) return 0
  const total = newLayer.value.items.reduce((sum, item) => sum + (item.width || 0), 0)
  return Number(total.toFixed(2))
})

const isTotalWidthValid = computed(() => {
  return Math.abs(totalWidth.value - 100) < 0.01
})







// Watch for any changes in newLayer to update preview immediately
watch(() => newLayer.value, () => {
  // Force reactivity update
}, { deep: true })

// Cleanup function for component unmount
onUnmounted(() => {
  try {
    // پاک کردن منابع مربوط به کتابخانه رسانه
    if (showMediaLibrary.value) {
      showMediaLibrary.value = false
    }
    
    if (selectedMediaIds.value.length > 0) {
      selectedMediaIds.value = []
    }
    
    console.log('✅ منابع کتابخانه رسانه پاک شدند')
    
  } catch (error) {
    console.error('❌ خطا در پاک کردن منابع:', error)
  }
})



function removeItem(idx: number) {
  if (!Array.isArray(newLayer.value.items)) return
  newLayer.value.items.splice(idx, 1)
  // Recalculate widths
  const total = newLayer.value.items.length
  if (total) {
    newLayer.value.items.forEach(it => {
      it.width = Math.round((100 / total) * 100) / 100
    })
  }
  
  // اطمینان از اینکه مجموع همه عرض‌ها دقیقاً 100% باشد
  const totalWidth = newLayer.value.items.reduce((sum, item) => sum + (item.width || 0), 0)
  if (Math.abs(totalWidth - 100) > 0.01) {
    // تنظیم آخرین آیتم برای رسیدن به دقیقاً 100%
    const lastIndex = newLayer.value.items.length - 1
    if (lastIndex >= 0) {
      const currentTotal = newLayer.value.items.reduce((sum, item, index) => 
        index === lastIndex ? sum : sum + (item.width || 0), 0)
      newLayer.value.items[lastIndex].width = Math.round((100 - currentTotal) * 100) / 100
    }
  }
  
  activeIndex.value = null
  enforceWidthBudget()
}

// تابع تنظیم عرض سایر آیتم‌ها وقتی یکی تغییر می‌کند
function adjustOtherItemWidths(changedIndex: number) {
  if (!Array.isArray(newLayer.value.items) || newLayer.value.items.length <= 1) return
  
  const changedItem = newLayer.value.items[changedIndex]
  if (!changedItem) return
  
  // اطمینان از اینکه عرض منفی نباشد
  if (changedItem.width < 0) {
    changedItem.width = 0
  }
  
  // اگر عرض آیتم تغییر یافته از 100% بیشتر باشد، آن را به 100% محدود کن
  if (changedItem.width > 100) {
    changedItem.width = 100
  }
  
  // محاسبه کل عرض آیتم‌های دیگر
  const otherItems = newLayer.value.items.filter((_, index) => index !== changedIndex)
  const otherItemsTotalWidth = otherItems.reduce((sum, item) => sum + (item.width || 0), 0)
  
  // محاسبه عرض باقی‌مانده برای سایر آیتم‌ها
  const remainingWidth = Math.max(0, 100 - changedItem.width)
  
  if (otherItems.length > 0 && remainingWidth > 0) {
    // تقسیم عرض باقی‌مانده بین سایر آیتم‌ها به نسبت عرض فعلی‌شان
    if (otherItemsTotalWidth > 0) {
      otherItems.forEach(item => {
        const ratio = item.width / otherItemsTotalWidth
        item.width = Math.round(remainingWidth * ratio * 100) / 100
      })
    } else {
      // اگر سایر آیتم‌ها عرض صفر دارند، عرض باقی‌مانده را مساوی تقسیم کن
      const equalWidth = remainingWidth / otherItems.length
      otherItems.forEach(item => {
        item.width = Math.round(equalWidth * 100) / 100
      })
    }
  }
  
  // اطمینان از اینکه مجموع همه عرض‌ها دقیقاً 100% باشد
  const totalWidth = newLayer.value.items.reduce((sum, item) => sum + (item.width || 0), 0)
  if (Math.abs(totalWidth - 100) > 0.01) {
    // تنظیم آخرین آیتم برای رسیدن به دقیقاً 100%
    const lastIndex = newLayer.value.items.length - 1
    if (lastIndex >= 0) {
      const currentTotal = newLayer.value.items.reduce((sum, item, index) => 
        index === lastIndex ? sum : sum + (item.width || 0), 0)
      newLayer.value.items[lastIndex].width = Math.round((100 - currentTotal) * 100) / 100
    }
  }
  
  // اطمینان از اینکه هیچ آیتمی عرض منفی نداشته باشد
  newLayer.value.items.forEach(item => {
    if (item.width < 0) item.width = 0
  })
  
  // نمایش پیام موفقیت
  console.log('✅ عرض آیتم‌ها تنظیم شد. مجموع کل:', totalWidth.toFixed(2) + '%')
  
  enforceWidthBudget(changedIndex)
  
  // نمایش پیام به کاربر (فقط در صورت تغییر قابل توجه)
  const finalTotalWidth = newLayer.value.items.reduce((sum, item) => sum + (item.width || 0), 0)
  if (Math.abs(finalTotalWidth - 100) < 0.01) {
    console.log('✅ مجموع عرض‌ها دقیقاً 100% شد')
  }
}

// تابع تنظیم خودکار عرض‌ها به حالت مساوی
function resetItemWidths() {
  if (!Array.isArray(newLayer.value.items) || newLayer.value.items.length === 0) return
  
  const total = newLayer.value.items.length
  newLayer.value.items.forEach(item => {
    item.width = Math.round((100 / total) * 100) / 100
  })
  
  // نمایش پیام موفقیت
  console.log('✅ عرض‌ها به حالت مساوی تنظیم شدند')
  
  // نمایش پیام به کاربر
  alert('عرض‌ها به حالت مساوی تنظیم شدند!')

  enforceWidthBudget()
}

// تابع تعیین چینش کانتینر بر اساس چینش آیتم‌ها
function getContainerAlignment() {
  if (!newLayer.value.items || newLayer.value.items.length === 0) {
    return { justifyContent: 'center' }
  }
  
  // بررسی اینکه آیا همه آیتم‌ها چینش یکسانی دارند
  const firstAlign = newLayer.value.items[0]?.align || 'center'
  const allSameAlign = newLayer.value.items.every(item => (item.align || 'center') === firstAlign)
  
  if (allSameAlign) {
    switch (firstAlign) {
      case 'left':
        return { justifyContent: 'flex-start' }
      case 'center':
        return { justifyContent: 'center' }
      case 'right':
        return { justifyContent: 'flex-end' }
      default:
        return { justifyContent: 'center' }
    }
  }
  
  // اگر چینش‌ها متفاوت باشند، از space-between استفاده می‌کنیم
  return { justifyContent: 'space-between' }
}

// تابع تعیین چینش هر آیتم جداگانه
function getItemAlignment(align: string): string {
  switch (align) {
    case 'left':
      return 'flex-start'  // در RTL: چپ = flex-start
    case 'center':
      return 'center'
    case 'right':
      return 'flex-end'  // در RTL: راست = flex-end
    default:
      return 'center'
  }
}

// تابع تعیین استایل کانتینر ایکون در پیش‌نمایش
function getIconContainerStyle(align: string) {
  return {
    display: 'flex',
    alignItems: 'center',
    justifyContent: getItemAlignment(align),
    width: '100%',
    height: '100%'
  }
}

// توابع مربوط به کتابخانه رسانه
function openMediaLibrary() {
  showMediaLibrary.value = true
}

function onMediaSelected(selectedMedia: any[]) {
  try {
    if (selectedMedia.length > 0 && activeIndex.value !== null) {
      const currentItem = newLayer.value.items[activeIndex.value]
      if (currentItem && currentItem.id === 'image') {
        const media = selectedMedia[0] // فقط اولین عکس انتخاب شده
        
        // ذخیره اطلاعات عکس در آیتم
        currentItem.imageUrl = media.url
        currentItem.imageName = media.name
        currentItem.imageId = media.id
        
        console.log('✅ عکس از کتابخانه رسانه انتخاب شد:', {
          id: media.id,
          name: media.name,
          url: media.url
        })
        
        // نمایش پیام موفقیت
        alert('عکس با موفقیت انتخاب شد!')
      }
    }
    
    // بستن کتابخانه رسانه
    showMediaLibrary.value = false
    
  } catch (error) {
    console.error('❌ خطا در انتخاب عکس از کتابخانه رسانه:', error)
    alert('خطا در انتخاب عکس: ' + error.message)
  }
}

// تابع حذف عکس انتخاب شده
function removeSelectedImage() {
  if (activeIndex.value === null) return
  
  const currentItem = newLayer.value.items[activeIndex.value]
  if (currentItem && currentItem.id === 'image') {
    delete currentItem.imageUrl
    delete currentItem.imageName
    delete currentItem.imageId
    
   
  }
}

function formatFileSize(bytes: number): string {
  if (bytes === 0) return '0 Bytes'
  const k = 1024
  const sizes = ['Bytes', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}


</script>

<style scoped>
.preview-section {
  background: #fff;
  border-radius: 10px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.04);
  padding: 24px;
  margin-bottom: 24px;
  width: 100%;
}

/* حالت تمام عرض برای ویرایش */
.preview-section.full-width {
  width: 100%;
  max-width: none;
}

.header-info {
  margin-bottom: 24px;
  padding-bottom: 20px;
  border-bottom: 1px solid #e9ecef;
}

.header-form-row {
  display: flex;
  gap: 20px;
  align-items: flex-start;
}

.header-form-row .form-group {
  flex: 1;
  margin-bottom: 0;
}

.header-info label {
  display: block;
  font-weight: 500;
  color: #2c3e50;
  margin-bottom: 8px;
}

.header-info .form-control {
  width: 100%;
  padding: 12px;
  border: 1px solid #ced4da;
  border-radius: 6px;
  font-size: 14px;
  transition: border-color 0.2s;
}

.header-info .form-control:focus {
  outline: none;
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

.preview-section h3 {
  font-size: 18px;
  font-weight: 600;
  color: #2c3e50;
  margin-bottom: 20px;
}

.header-preview {
  border: 2px dashed #e9ecef;
  border-radius: 8px;
  padding: 20px;
  background: #f8f9fa;
}

.preview-header {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
  padding: 16px;
  background: #f5f5f5;
  border-radius: 6px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  position: relative;
  transition: all 0.3s ease;
  margin: 0 auto;
  min-height: 60px;
  width: 100%;
  direction: rtl;
  text-align: right;
  overflow: hidden;
}

.preview-container {
  display: flex;
  flex-direction: column;
  align-items: stretch;
  justify-content: flex-start;
  padding: 0 24px;
  background: transparent;
  border-radius: 6px;
  position: relative;
  transition: all 0.3s ease;
  width: 100%;
  direction: rtl;
  text-align: right;
  overflow: hidden;
}

/* Responsive Design برای لایه‌ها */
.preview-layer {
  transition: all 0.3s ease;
  direction: rtl;
  text-align: right;
}

/* Mobile Responsive */
@media (max-width: 768px) {
  .preview-header {
    padding: 12px;
    min-height: 50px;
    justify-content: center;
  }
  
  .preview-layer {
    width: 100% ;
    margin: 0 0 8px 0;
    flex-direction: column;
    align-items: center;
    text-align: center;
  }
  
  .preview-layer .layer-items-preview {
    flex-direction: column;
    gap: 1px;
  }
  
  /* نمایش عمودی آیتم‌ها در موبایل */
  .layer-items-preview-horizontal {
    flex-direction: column;
    height: auto;
  }
  
  .layer-item-preview-box {
    width: 100% !important;
    border-right: none !important;
    border-bottom: 1px solid rgba(0,0,0,0.1);
    min-height: 30px;
  }
  
  .layer-item-preview-box:last-child {
    border-bottom: none;
  }
}

/* بهبود نمایش در صفحات کوچک */
@media (max-width: 480px) {
  .layer-item-preview-box {
    padding: 6px 2px;
    min-height: 25px;
  }
  
  .item-preview-name {
    font-size: 10px;
  }
  
  .item-preview-width {
    font-size: 8px;
    padding: 1px 3px;
  }
}

/* Tablet Responsive */
@media (max-width: 1024px) and (min-width: 769px) {
  .preview-header {
    padding: 14px;
  }
  
  .preview-layer {
    padding: 0 12px;
  }
  
  /* بهبود نمایش در تبلت */
  .layer-item-preview-box {
    padding: 10px 6px;
    min-height: 35px;
  }
  
  .item-preview-name {
    font-size: 12px;
  }
  
  .item-preview-width {
    font-size: 10px;
    padding: 3px 6px;
  }
}

/* RTL Support */
.preview-layer[style*="direction: rtl"] {
  text-align: right;
  justify-content: flex-end;
}

.preview-layer[style*="direction: ltr"] {
  text-align: left;
  justify-content: flex-start;
}

/* Layer Items Preview */
.layer-items-preview {
  display: flex;
  flex-wrap: wrap;
  gap: 2px;
  align-items: center;
  justify-content: flex-end;
  direction: rtl;
}

.layer-item-tag {
  background: rgba(0,0,0,0.1);
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 12px;
  color: #333;
  white-space: nowrap;
}

/* نمایش افقی آیتم‌های لایه */
.layer-items-preview-horizontal {
  display: flex;
  width: 100%;
  height: 100%;
  align-items: stretch;
  direction: rtl;
}

.layer-item-preview-box {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 8px 4px;
  box-sizing: border-box;
  position: relative;
  min-height: 40px;
  transition: all 0.2s ease;
}

.layer-item-preview-box:hover {
  background-color: rgba(255,255,255,0.1);
}

.item-preview-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 4px;
  text-align: center;
}

.item-preview-name {
  font-size: 11px;
  font-weight: 600;
  color: #2c3e50;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  max-width: 100%;
}

.item-preview-width {
  font-size: 9px;
  color: #667eea;
  font-weight: 500;
  background: rgba(255,255,255,0.8);
  padding: 2px 4px;
  border-radius: 3px;
  border: 1px solid rgba(0,0,0,0.1);
}

.preview-header-label {
  text-align: center;
  margin-bottom: 12px;
  padding: 8px 16px;
  background: #f8f9fa;
  border-radius: 6px;
  border: 1px solid #e9ecef;
  font-size: 14px;
  color: #495057;
}

/* استایل‌های لایه‌های ایجاد شده */
.created-layers-preview {
  display: flex;
  flex-direction: column;
  gap: 8px;
  width: 100%;
}

.preview-layer {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 16px 24px;
  border-radius: 6px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  position: relative;
  transition: all 0.3s ease;
  margin: 0 auto;
  min-height: 60px;
  border: 2px solid rgba(0,0,0,0.1);
  background: transparent;
}

.layer-name-preview {
  font-weight: 600;
  color: #2c3e50;
  margin-bottom: 8px;
  font-size: 14px;
  text-align: right;
  position: absolute;
  top: -25px;
  right: 0;
  background: #fff;
  padding: 4px 8px;
  border-radius: 4px;
  border: 1px solid #e9ecef;
  z-index: 10;
}

.layer-items-preview {
  display: flex;
  flex-wrap: wrap;
  gap: 2px;
  justify-content: center;
}

.layer-item-tag {
  padding: 4px 8px;
  background: rgba(255, 255, 255, 0.9);
  border-radius: 4px;
  font-size: 11px;
  color: #2c3e50;
  border: 1px solid rgba(0, 0, 0, 0.1);
}

.no-layers-message {
  text-align: center;
  padding: 40px 20px;
  color: #6c757d;
  background: #f8f9fa;
  border-radius: 8px;
  border: 2px dashed #dee2e6;
}

.no-layers-message p {
  margin: 8px 0;
  font-size: 14px;
}

.no-layers-message p:first-child {
  font-weight: 600;
  color: #495057;
}

.preview-logo {
  font-weight: bold;
  color: #667eea;
  font-size: 18px;
  max-height: 100%;
  object-fit: contain;
}

.item-icon-container {
  flex: 1;
  min-height: 0;
  overflow: hidden;
}

.item-icon {
  font-size: 24px;
  line-height: 1;
  max-height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.preview-menu {
  display: flex;
  gap: 20px;
}

.preview-menu span {
  color: #6c757d;
  cursor: pointer;
  transition: color 0.2s;
}

.preview-menu span:hover {
  color: #667eea;
}

.preview-actions {
  display: flex;
  gap: 16px;
  align-items: center;
}

.preview-actions span {
  padding: 8px 12px;
  background: #f8f9fa;
  border-radius: 4px;
  color: #6c757d;
  font-size: 14px;
  cursor: pointer;
  transition: all 0.2s;
}

.preview-actions span:hover {
  background: #667eea;
  color: white;
}

/* استایل‌های جدید برای پیش‌نمایش زنده */
.preview-items-container {
  display: flex;
  width: 100%;
  position: relative;
  margin: 0;
  padding: 0;
  overflow: hidden;
  height: 100%;
  align-items: stretch;
  /* justify-content will be set dynamically via JavaScript */
}

.preview-section-item {
  /* سایر استایل‌ها */
  border: none;
  border-radius: 8px;
  min-height: 48px;
  background: #fff;
  margin: 0;
  position: relative;
  transition: border-color 0.2s;
  flex: 1;
  overflow: hidden;
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: stretch;
  justify-content: center;
}

/* استایل دستگیره تغییر اندازه */
.resize-handle {
  width: 8px;
  height: 100%;
  background: transparent;
  cursor: col-resize !important;
  position: relative;
  z-index: 10;
  border-left: 2px solid #9ca3af;
  border-right: 2px solid #9ca3af;
  margin: 0 -2px;
}

.resize-handle:hover {
  border-left: 3px solid #6b7280;
  border-right: 3px solid #6b7280;
  background: rgba(107, 114, 128, 0.1);
}

.resize-handle:active {
  border-left: 4px solid #374151;
  border-right: 4px solid #374151;
  background: rgba(55, 65, 81, 0.2);
}

/* استایل input کوچک */
.settings-input-small {
  width: 200px;
  padding: 12px 18px;
  border: 1px solid #d1d5db;
  border-radius: 4px;
  font-size: 20px;
  text-align: center;
  background: white;
}

.settings-input-small:focus {
  outline: none;
  border-color: #667eea;
  box-shadow: 0 0 0 2px rgba(102, 126, 234, 0.1);
}

.settings-unit {
  font-size: 12px;
  color: #6b7280;
  margin-left: 4px;
}

/* استایل نمایش عرض */
.settings-display {
  font-weight: 600;
  color: #667eea;
  background: #f0f4ff;
  padding: 4px 8px;
  border-radius: 4px;
  border: 1px solid #e0e7ff;
  font-size: 12px;
}

.preview-section-item:last-child {
  border-right: none;
}

.preview-items {
  display: flex;
  flex-wrap: wrap;
  gap: 12px;
  align-items: center;
  justify-content: center;
  width: 100%;
  padding: 8px;
  min-height: 40px;
}

/* استایل‌های آیتم‌ها حالا با Tailwind classes تعریف می‌شن */

.preview-separator-horizontal {
  position: relative;
  width: 100%;
  height: 0;
  margin: 8px 0;
  transition: all 0.3s ease;
  border-top: 1px solid #e9ecef;
}

.preview-separator-vertical {
  position: relative;
  width: 0;
  height: 100%;
  margin: 0 8px;
  transition: all 0.3s ease;
  border-left: 1px solid #e9ecef;
}

/* استایل‌های جداکننده بین بخش‌ها */
.section-divider {
  position: absolute;
  right: -1px;
  top: 0;
  bottom: 0;
  width: 2px;
  background: transparent;
  cursor: col-resize;
  z-index: 10;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s ease;
  border-left: 1px dashed #999;
}

.section-divider:hover {
  background: rgba(102, 126, 234, 0.3);
  width: 4px;
}

.section-divider:hover .resize-handle {
  opacity: 1;
  transform: scale(1.2);
}

.resize-handle {
  font-size: 16px;
  font-weight: bold;
  color: #667eea;
  opacity: 0.6;
  transition: all 0.2s ease;
  user-select: none;
  pointer-events: none;
}

.layer-info {
  margin-top: 16px;
  padding: 16px;
  background: #f8f9fa;
  border-radius: 8px;
  border: 1px solid #e9ecef;
}

.info-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px 0;
  border-bottom: 1px solid #e9ecef;
  font-size: 13px;
}

.info-item:last-child {
  border-bottom: none;
}

.info-item strong {
  color: #2c3e50;
  font-weight: 600;
}

/* استایل‌های نمایش مجموع عرض‌ها */
.width-summary {
  margin-top: 16px;
  padding: 16px;
  background: #f8f9fa;
  border-radius: 8px;
  border: 1px solid #e9ecef;
}

.width-summary-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
  padding-bottom: 8px;
  border-bottom: 1px solid #e9ecef;
}

.total-width {
  font-weight: 600;
  font-size: 14px;
  padding: 4px 8px;
  border-radius: 4px;
}

.total-width.valid {
  background: #d4edda;
  color: #155724;
  border: 1px solid #c3e6cb;
}

.total-width.invalid {
  background: #f8d7da;
  color: #721c24;
  border: 1px solid #f5c6cb;
}

.width-items-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.width-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 6px 8px;
  background: #fff;
  border-radius: 4px;
  border: 1px solid #e9ecef;
  font-size: 12px;
}

.item-name {
  color: #495057;
  font-weight: 500;
}

.item-width {
  color: #667eea;
  font-weight: 600;
  font-family: monospace;
}

/* استایل‌های جدید برای نمایش inline عرض‌ها */
.width-controls-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
  margin-top: 12px;
  padding-top: 12px;
  border-top: 1px solid #e9ecef;
}

.width-items-inline {
  display: flex;
  align-items: center;
  gap: 12px;
  flex: 1;
}

.width-item-inline {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 4px 8px;
  background: #fff;
  border-radius: 4px;
  border: 1px solid #e9ecef;
  font-size: 11px;
  white-space: nowrap;
}

.item-name-small {
  color: #495057;
  font-weight: 500;
}

.item-width-small {
  color: #667eea;
  font-weight: 600;
  font-family: monospace;
  background: #f8f9fa;
  padding: 2px 6px;
  border-radius: 3px;
  border: 1px solid #e9ecef;
}

.reset-widths-btn {
  flex-shrink: 0;
  white-space: nowrap;
}

/* Responsive design برای نمایش inline عرض‌ها */
@media (max-width: 768px) {
  .width-controls-row {
    flex-direction: column;
    gap: 12px;
    align-items: stretch;
  }
  
  .width-items-inline {
    flex-direction: column;
    gap: 8px;
  }
  
  .width-item-inline {
    justify-content: space-between;
    width: 100%;
  }
  
  .reset-widths-btn {
    width: 100%;
  }
}

@media (max-width: 480px) {
  .width-item-inline {
    font-size: 10px;
    padding: 3px 6px;
  }
  
  .item-width-small {
    font-size: 10px;
    padding: 1px 4px;
  }
}

/* حذف شده - جایگزین شده با width-controls-row */

.btn {
  padding: 8px 16px;
  border: 1px solid #667eea;
  border-radius: 6px;
  background: transparent;
  color: #667eea;
  font-size: 12px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
}

.btn:hover:not(:disabled) {
  background: #667eea;
  color: white;
}

.btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.btn-sm {
  padding: 6px 12px;
  font-size: 11px;
}

.btn-outline {
  background: transparent;
  border: 1px solid #667eea;
  color: #667eea;
}

.width-warning {
  margin-top: 4px;
  padding: 8px;
  background: #fff3cd;
  border: 1px solid #ffeaa7;
  border-radius: 4px;
  text-align: center;
}

.width-warning-compact {
  margin-top: 8px;
  padding: 8px 12px;
  background: #fff3cd;
  border: 1px solid #ffeaa7;
  border-radius: 4px;
  text-align: center;
  max-width: 100%;
  font-size: 11px;
  line-height: 1.4;
}

.info-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 0;
  flex-wrap: wrap;
  width: 100%;
}

.info-row .info-item {
  display: inline;
  padding: 0;
  border-bottom: none;
  font-size: 12px;
  color: #495057;
  background: #fff;
  padding: 4px 8px;
  border-radius: 4px;
  border: 1px solid #dee2e6;
  white-space: nowrap;
  flex: 0 0 auto;
}

/* Responsive styles */
@media (max-width: 768px) {
  .header-form-row {
    flex-direction: column;
    gap: 16px;
  }
  
  .header-form-row .form-group {
    margin-bottom: 16px;
  }
  
  .header-form-row .form-group:last-child {
    margin-bottom: 0;
  }
}

.item-settings-dropdown{
  position:absolute;
  top:100%;
  left:50%;
  transform:translateX(-50%);
  margin-top:4px;
  background:#fff;
  border:1px solid #d1d5db;
  border-radius:6px;
  padding:6px 8px;
  width:160px;
  z-index:50;
  box-shadow:0 4px 12px rgba(0,0,0,.08);
  font-size:11px;
}
.item-settings-dropdown label{font-size:11px}
.item-settings-floating{background:#fff;border:1px solid #e5e7eb;border-radius:8px;padding:8px;max-width:200px;box-shadow:0 4px 10px rgba(0,0,0,.06);}

/* استایل‌های جدید برای پنل تنظیمات آیتم */
.item-settings-panel {
  background: #fff;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  padding: 16px;
  box-shadow: 0 4px 10px rgba(0,0,0,.06);
  margin-top: 16px;
  width: 100%;
  max-width: none;
}

.settings-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
  padding-bottom: 12px;
  border-bottom: 1px solid #e5e7eb;
}

.settings-header h4 {
  font-size: 14px;
  font-weight: 600;
  color: #2c3e50;
  margin: 0;
}

.close-settings-btn {
  background: #f3f4f6;
  border: 1px solid #d1d5db;
  border-radius: 4px;
  width: 24px;
  height: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  font-size: 12px;
  color: #6b7280;
  transition: all 0.2s;
}

.close-settings-btn:hover {
  background: #e5e7eb;
  color: #374151;
}

.settings-content {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.settings-row {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.settings-label {
  display: flex;
  flex-direction: column;
  gap: 6px;
  font-size: 13px;
  color: #374151;
}

.settings-label span:first-child {
  font-weight: 500;
  color: #2c3e50;
}

.settings-select {
  padding: 8px 12px;
  border: 1px solid #d1d5db;
  border-radius: 6px;
  font-size: 13px;
  background: #fff;
  transition: border-color 0.2s;
}

.settings-select:focus {
  outline: none;
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

.settings-input {
  padding: 8px 12px;
  border: 1px solid #d1d5db;
  border-radius: 6px;
  font-size: 13px;
  width: 80px;
  text-align: center;
  transition: border-color 0.2s;
}

.settings-input:focus {
  outline: none;
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

.unit {
  font-size: 12px;
  color: #6b7280;
  margin-left: 4px;
}

.color-input-group {
  display: flex;
  align-items: center;
  gap: 12px;
}

.color-picker {
  width: 40px;
  height: 32px;
  border: 1px solid #d1d5db;
  border-radius: 4px;
  cursor: pointer;
  background: transparent;
}

.transparent-checkbox {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 12px;
  color: #6b7280;
  cursor: pointer;
}

.transparent-checkbox input[type="checkbox"] {
  width: 14px;
  height: 14px;
  cursor: pointer;
}

.settings-actions {
  margin-top: 8px;
  padding-top: 12px;
  border-top: 1px solid #e5e7eb;
}

.remove-item-btn {
  background: #ef4444;
  color: #fff;
  border: none;
  border-radius: 6px;
  padding: 8px 16px;
  font-size: 13px;
  cursor: pointer;
  transition: background-color 0.2s;
}

.remove-item-btn:hover {
  background: #dc2626;
}

/* استایل‌های جدید برای تنظیمات inline */
.settings-row-inline {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
  flex-wrap: nowrap;
  width: 100%;
}

.settings-label-inline {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  color: #374151;
  white-space: nowrap;
  flex: 1;
  justify-content: center;
}

.settings-label-inline span:first-child {
  font-weight: 500;
  color: #2c3e50;
}

.settings-select-small {
  padding: 4px 8px;
  border: 1px solid #d1d5db;
  border-radius: 4px;
  font-size: 12px;
  background: #fff;
  transition: border-color 0.2s;
  min-width: 60px;
}

.settings-select-small:focus {
  outline: none;
  border-color: #667eea;
  box-shadow: 0 0 0 2px rgba(102, 126, 234, 0.1);
}

.settings-input-small {
  padding: 12px 18px;
  border: 1px solid #d1d5db;
  border-radius: 4px;
  font-size: 20px;
  width: 200px;
  text-align: center;
  transition: border-color 0.2s;
}

.settings-input-small:focus {
  outline: none;
  border-color: #667eea;
  box-shadow: 0 0 0 2px rgba(102, 126, 234, 0.1);
}

.color-picker-small {
  width: 30px;
  height: 24px;
  border: 1px solid #d1d5db;
  border-radius: 4px;
  cursor: pointer;
  background: transparent;
}

.transparent-checkbox-inline {
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 11px;
  color: #6b7280;
  cursor: pointer;
  white-space: nowrap;
}

.transparent-checkbox-inline input[type="checkbox"] {
  width: 12px;
  height: 12px;
  cursor: pointer;
}

.remove-item-btn-inline {
  background: #ef4444;
  color: #fff;
  border: none;
  border-radius: 4px;
  padding: 4px 12px;
  font-size: 12px;
  cursor: pointer;
  transition: background-color 0.2s;
  white-space: nowrap;
}

.remove-item-btn-inline:hover {
  background: #dc2626;
}

/* استایل خط جداکننده آبی */
.section-divider {
  height: 2px;
  background: #667eea;
  margin: 20px 0;
  border-radius: 2px;
  width: 100%;
  display: block;
}

/* استایل‌های مربوط به آپلود عکس */
.image-upload-section {
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  padding: 16px;
  background: #f9fafb;
}

.image-upload-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.image-upload-header h5 {
  margin: 0;
  font-size: 14px;
  font-weight: 600;
  color: #374151;
}

/* Responsive برای بخش آپلود عکس */
@media (max-width: 768px) {
  .image-upload-section {
    padding: 14px;
  }
  
  .image-upload-header {
    flex-direction: column;
    gap: 8px;
    align-items: stretch;
  }
  
  .image-upload-header h5 {
    font-size: 13px;
    text-align: center;
  }
}

@media (max-width: 480px) {
  .image-upload-section {
    padding: 12px;
  }
  
  .image-upload-header h5 {
    font-size: 12px;
  }
}

.selected-image-preview {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px;
  background: white;
  border-radius: 6px;
  border: 1px solid #e5e7eb;
}

.preview-image {
  width: 60px;
  height: 60px;
  object-fit: cover;
  border-radius: 4px;
  border: 1px solid #d1d5db;
}

/* Responsive برای preview عکس انتخاب شده */
@media (max-width: 768px) {
  .selected-image-preview {
    flex-direction: column;
    gap: 8px;
    text-align: center;
  }
  
  .preview-image {
    width: 80px;
    height: 80px;
  }
}

@media (max-width: 480px) {
  .selected-image-preview {
    padding: 10px;
  }
  
  .preview-image {
    width: 70px;
    height: 70px;
  }
}

.image-info {
  display: flex;
  flex-direction: column;
  gap: 4px;
  flex: 1;
}

.image-name {
  font-size: 13px;
  font-weight: 500;
  color: #374151;
}

.remove-image-btn {
  background: #ef4444;
  color: white;
  border: none;
  border-radius: 4px;
  width: 24px;
  height: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  font-size: 12px;
  transition: background-color 0.2s;
}

.remove-image-btn:hover {
  background: #dc2626;
}

.image-upload-guide {
  padding: 12px;
  background: white;
  border-radius: 6px;
  border: 1px solid #e5e7eb;
}

.image-upload-guide p {
  margin: 0;
  text-align: center;
}

/* Responsive برای راهنمای آپلود */
@media (max-width: 768px) {
  .image-upload-guide {
    padding: 10px;
  }
  
  .image-upload-guide p {
    font-size: 13px;
  }
}

@media (max-width: 480px) {
  .image-upload-guide {
    padding: 8px;
  }
  
  .image-upload-guide p {
    font-size: 12px;
  }
}


</style> 
