<template>
  <div class="create-header-page">
    <div class="main-content" :class="{ 'full-width': isEditing }">
      <PageHeader />

      <div class="header-settings-container">
        <HeaderSettingsForm />
      </div>

      <div v-if="showLayerSettings" class="layer-edit-layout">
        <HeaderSettingsSidebar />
        <div class="preview-column">
          <HeaderPreview />
        </div>
      </div>

      <div v-else class="normal-layout">
        <FormContainer />
        <HeaderPreview />
      </div>
    </div>

    <ItemsSelectionModal />

    <div
      v-if="showToast"
      class="toast-notification"
      :class="toastType === 'success' ? 'toast-success' : 'toast-error'"
    >
      {{ toastMessage }}
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, provide, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import HeaderSettingsSidebar from './HeaderSettingsSidebar.vue'
import PageHeader from './PageHeader.vue'
import HeaderPreview from './HeaderPreview.vue'
import FormContainer from './FormContainer.vue'
import ItemsSelectionModal from './ItemsSelectionModal.vue'
import HeaderSettingsForm from './HeaderSettingsForm.vue'

// @ts-ignore
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

interface HeaderLayerItem {
  id: string
  name?: string
  path?: string
  icon?: string
  align?: string
  width?: number
  paddingRight?: number
  paddingLeft?: number
  imageUrl?: string
  imageName?: string
  [key: string]: any
}

interface HeaderLayer {
  id?: string | number
  name: string
  width: number
  height: number
  rowCount: number
  color: string
  opacity: number
  enableBorder: boolean
  borderPosition: string
  borderColor: string
  borderWidth: number
  borderStyle: string
  enableShadow: boolean
  shadowIntensity: string
  shadowDirection: string
  showSeparator: boolean
  separatorType: string
  separatorColor: string
  separatorOpacity: number
  separatorWidth: number
  items: HeaderLayerItem[]
  direction: string
  mobileResponsive: boolean
  tabletResponsive: boolean
  paddingRight: number
  paddingLeft: number
  paddingTop: number
  paddingBottom: number
  styleSettings?: Record<string, any>
  boxWidths?: number[]
  createdAt?: string
  updatedAt?: string
}

const headerData = ref({
  name: '',
  description: '',
  pageSelection: 'all',
  specificPages: '',
  excludedPages: '',
  isActive: true
})

const layerData = ref({
  rowCount: 1,
  layerWidth: 100,
  layerHeight: 50,
  items: [],
  layerColor: '#ffffff',
  layerOpacity: 100,
  showSeparator: false
})

const showItemsModal = ref(false)
const showLayerSettings = ref(false)
const createdLayers = ref<HeaderLayer[]>([])
const newLayer = ref<HeaderLayer>(createBaseLayer())

const availableItems = [
  { id: 'logo', name: 'لوگو', path: '/', icon: 'heroicons:building-office-2' },
  { id: 'search', name: 'باکس جستجو', path: '/search', icon: 'heroicons:magnifying-glass' },
  { id: 'auth', name: 'ورود و ثبت‌نام', path: '/auth/login', icon: 'heroicons:user-circle' },
  { id: 'cart', name: 'سبد خرید', path: '/cart', icon: 'heroicons:shopping-cart' },
  { id: 'notification', name: 'اعلان‌ها', path: '/notifications', icon: 'heroicons:bell' },
  { id: 'menu', name: 'منوی اصلی', path: '/menu', icon: 'heroicons:bars-3' },
  { id: 'language', name: 'انتخاب زبان', path: '/language', icon: 'heroicons:globe-alt' },
  { id: 'currency', name: 'انتخاب ارز', path: '/currency', icon: 'heroicons:currency-dollar' },
  { id: 'image', name: 'آپلود عکس', path: '/media', icon: 'heroicons:photo' },
  { id: 'custom', name: 'بلوک سفارشی', path: '/custom', icon: 'heroicons:cube' },
  { id: 'social', name: 'شبکه‌های اجتماعی', path: '/social', icon: 'heroicons:share' }
]

const route = useRoute()
const router = useRouter()
const isEditing = computed(() => !!route.query.id)

const showToast = ref(false)
const toastMessage = ref('')
const toastType = ref<'success' | 'error'>('success')

function showToastMessage(message: string, type: 'success' | 'error' = 'success') {
  toastMessage.value = message
  toastType.value = type
  showToast.value = true
  setTimeout(() => {
    showToast.value = false
  }, 3000)
}

function createBaseLayer(): HeaderLayer {
  return {
    name: '',
    width: 100,
    height: 60,
    rowCount: 1,
    color: '#ffffff',
    opacity: 100,
    enableBorder: false,
    borderPosition: 'all',
    borderColor: '#e5e7eb',
    borderWidth: 1,
    borderStyle: 'solid',
    enableShadow: false,
    shadowIntensity: 'md',
    shadowDirection: 'top',
    showSeparator: false,
    separatorType: 'solid',
    separatorColor: '#e9ecef',
    separatorOpacity: 100,
    separatorWidth: 1,
    items: [],
    direction: 'rtl',
    mobileResponsive: true,
    tabletResponsive: true,
    paddingRight: 0,
    paddingLeft: 0,
    paddingTop: 0,
    paddingBottom: 0,
    styleSettings: {}
  }
}

function parseBoolean(value: any, fallback = false): boolean {
  if (typeof value === 'boolean') return value
  if (typeof value === 'number') return value === 1
  if (typeof value === 'string') {
    const normalized = value.trim().toLowerCase()
    if (['true', '1', 'yes', 'on'].includes(normalized)) return true
    if (['false', '0', 'no', 'off'].includes(normalized)) return false
  }
  return fallback
}

function parseJSON<T = any>(value: any): T | null {
  if (!value) return null
  if (typeof value === 'object') return value as T
  if (typeof value === 'string') {
    try {
      return JSON.parse(value) as T
    } catch (error) {
      console.warn('parseJSON failed:', error)
      return null
    }
  }
  return null
}

function normalizePercent(raw: any, fallback = 100): number {
  if (raw === undefined || raw === null || raw === '') return fallback
  const numeric = Number(raw)
  if (Number.isNaN(numeric)) return fallback
  if (numeric <= 1 && numeric >= 0) return Math.round(numeric * 100)
  if (numeric < 0) return 0
  if (numeric > 100) return 100
  return Math.round(numeric)
}

function parseStyleSettings(layer: any) {
  return (
    parseJSON(layer?.styleSettings) ||
    parseJSON(layer?.style_settings) ||
    layer?.styleSettings ||
    layer?.style_settings ||
    {}
  ) as Record<string, any>
}

function normalizeItems(rawItems: any): HeaderLayerItem[] {
  if (Array.isArray(rawItems)) {
    return rawItems.map(item => (typeof item === 'object' ? { ...item } : { id: item }))
  }
  if (typeof rawItems === 'string' && rawItems.trim() !== '') {
    const parsed = parseJSON(rawItems)
    if (Array.isArray(parsed)) {
      return parsed.map(item => (typeof item === 'object' ? { ...item } : { id: item }))
    }
  }
  return []
}

function normalizeLayer(rawLayer: any): HeaderLayer {
  const styleSettings = parseStyleSettings(rawLayer)
  const borderSettings = styleSettings.border || {}
  const shadowSettings = styleSettings.shadow || {}
  const layoutSettings = styleSettings.layout || {}

  const normalized: HeaderLayer = {
    id: rawLayer.id,
    name: rawLayer.name || '',
    width: Number(rawLayer.width ?? layoutSettings.width ?? 100),
    height: Number(rawLayer.height ?? layoutSettings.height ?? 60),
    rowCount: Number(rawLayer.rowCount ?? rawLayer.row_count ?? layoutSettings.rowCount ?? 1),
    color: rawLayer.color || layoutSettings.backgroundColor || '#ffffff',
    opacity: normalizePercent(rawLayer.opacity ?? rawLayer.opacity_percentage, 100),
    enableBorder: parseBoolean(
      rawLayer.enableBorder ?? 
      rawLayer.enable_border ?? 
      rawLayer.borderEnabled ?? 
      borderSettings.enabled, 
      false
    ),
    borderPosition: rawLayer.borderPosition ?? rawLayer.border_position ?? borderSettings.position ?? 'all',
    borderColor: rawLayer.borderColor ?? rawLayer.border_color ?? borderSettings.color ?? '#e5e7eb',
    borderWidth: Number(rawLayer.borderWidth ?? rawLayer.border_width ?? borderSettings.width ?? 1),
    borderStyle: rawLayer.borderStyle ?? rawLayer.border_style ?? borderSettings.style ?? 'solid',
    enableShadow: parseBoolean(
      rawLayer.enableShadow ??
        rawLayer.enable_shadow ??
        rawLayer.shadowEnabled ??
        rawLayer.elevationEnabled ??
        shadowSettings.enabled,
      false
    ),
    shadowIntensity: rawLayer.shadowIntensity ?? rawLayer.shadow_intensity ?? shadowSettings.intensity ?? 'md',
    shadowDirection: rawLayer.shadowDirection ?? rawLayer.shadow_direction ?? shadowSettings.direction ?? 'top',
    showSeparator: parseBoolean(rawLayer.showSeparator ?? rawLayer.show_separator ?? rawLayer.separatorEnabled ?? layoutSettings.showSeparator, false),
    separatorType: rawLayer.separatorType ?? rawLayer.separator_type ?? layoutSettings.separatorType ?? 'solid',
    separatorColor: rawLayer.separatorColor ?? rawLayer.separator_color ?? layoutSettings.separatorColor ?? '#e9ecef',
    separatorOpacity: normalizePercent(rawLayer.separatorOpacity ?? rawLayer.separator_opacity ?? layoutSettings.separatorOpacity, 100),
    separatorWidth: Number(rawLayer.separatorWidth ?? rawLayer.separator_width ?? layoutSettings.separatorWidth ?? 1),
    items: normalizeItems(rawLayer.items),
    direction: rawLayer.direction ?? layoutSettings.direction ?? 'rtl',
    mobileResponsive: parseBoolean(rawLayer.mobileResponsive ?? rawLayer.mobile_responsive ?? layoutSettings.mobileResponsive, true),
    tabletResponsive: parseBoolean(rawLayer.tabletResponsive ?? rawLayer.tablet_responsive ?? layoutSettings.tabletResponsive, true),
    paddingRight: Number(rawLayer.paddingRight ?? rawLayer.padding_right ?? layoutSettings.paddingRight ?? 0),
    paddingLeft: Number(rawLayer.paddingLeft ?? rawLayer.padding_left ?? layoutSettings.paddingLeft ?? 0),
    paddingTop: Number(rawLayer.paddingTop ?? rawLayer.padding_top ?? layoutSettings.paddingTop ?? 0),
    paddingBottom: Number(rawLayer.paddingBottom ?? rawLayer.padding_bottom ?? layoutSettings.paddingBottom ?? 0),
    styleSettings: styleSettings || {},
    boxWidths: Array.isArray(rawLayer.boxWidths)
      ? [...rawLayer.boxWidths]
      : typeof rawLayer.boxWidths === 'string'
        ? parseJSON(rawLayer.boxWidths) || undefined
        : undefined,
    createdAt: rawLayer.createdAt ?? rawLayer.created_at,
    updatedAt: rawLayer.updatedAt ?? rawLayer.updated_at
  }

  // ❌ حذف شد: distributeItemWidths - کاربر باید دستی تنظیم کنه
  // if (normalized.items.length) {
  //   distributeItemWidths(normalized.items)
  // }

  return normalized
}

function buildStyleSettings(layer: HeaderLayer) {
  return {
    separator: {
      show: !!layer.showSeparator,
      type: layer.separatorType || 'solid',
      color: layer.separatorColor || '#e9ecef',
      opacity: (layer.separatorOpacity || 100) / 100,
      width: layer.separatorWidth || 1
    }
  }
}

function distributeItemWidths(items: HeaderLayerItem[]) {
  // ❌ این تابع دیگه استفاده نمیشه - برای سازگاری نگه داشته شده
  // کاربر باید دستی عرض هر آیتم رو تنظیم کنه
  if (!items.length) return
  
  // فقط برای آیتم‌هایی که width ندارند، یه مقدار پیش‌فرض بده
  items.forEach(item => {
    if (!item.width || item.width === 0) {
      item.width = 20 // عرض پیش‌فرض 20%
    }
  })
}

function persistLayersToSession() {
  if (isEditing.value) return
  if (typeof window === 'undefined') return
  sessionStorage.setItem('headerLayers', JSON.stringify(createdLayers.value))
}

function loadLayersFromSession() {
  if (isEditing.value) return
  if (typeof window === 'undefined') return
  const saved = sessionStorage.getItem('headerLayers')
  if (!saved) return
  try {
    const parsed = JSON.parse(saved)
    if (Array.isArray(parsed)) {
      createdLayers.value = parsed.map(normalizeLayer)
    }
  } catch (error) {
    console.warn('Failed to parse session header layers:', error)
  }
}

async function loadExistingHeader() {
  if (!isEditing.value) {
    createdLayers.value = []
    if (typeof window !== 'undefined') {
      sessionStorage.removeItem('headerLayers')
    }
    return
  }

  try {
    const headerId = parseInt(route.query.id as string, 10)
    if (Number.isNaN(headerId)) return

    const response = await fetch(`/api/admin/header-settings/${headerId}`)
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }
    const payload = await response.json()
    const existingHeader = payload?.data ?? payload

    headerData.value = {
      name: existingHeader.name || '',
      description: existingHeader.description || '',
      pageSelection: existingHeader.page_selection || existingHeader.pageSelection || 'all',
      specificPages: existingHeader.specific_pages || existingHeader.specificPages || '',
      excludedPages: existingHeader.excluded_pages || existingHeader.excludedPages || '',
      isActive: parseBoolean(existingHeader.is_active ?? existingHeader.isActive, true)
    }

    if (Array.isArray(existingHeader.layers)) {
      createdLayers.value = existingHeader.layers.map(normalizeLayer)
    } else {
      createdLayers.value = []
    }
  } catch (error) {
    console.error('Failed to load existing header:', error)
    showToastMessage('بارگذاری اطلاعات هدر با خطا مواجه شد.', 'error')
  }
}

function resetLayerForm() {
  newLayer.value = createBaseLayer()
  delete newLayer.value.id
}

function cancelLayer() {
  showLayerSettings.value = false
  resetLayerForm()
}

function openItemsModal() {
  showItemsModal.value = true
}

function closeItemsModal() {
  showItemsModal.value = false
}

function toggleItem(itemId: string) {
  const existingIndex = newLayer.value.items.findIndex(item => item.id === itemId)
  if (existingIndex > -1) {
    newLayer.value.items.splice(existingIndex, 1)
  } else {
    const item = availableItems.find(entry => entry.id === itemId)
    if (!item) return
    newLayer.value.items.push({
      id: item.id,
      name: item.name,
      path: item.path,
      icon: item.icon,
      align: 'center',
      width: 20  // عرض پیش‌فرض 20% - کاربر باید دستی تنظیم کنه
    })
  }

  // ❌ حذف شد: distributeItemWidths - کاربر باید دستی تنظیم کنه
  // if (newLayer.value.items.length) {
  //   distributeItemWidths(newLayer.value.items)
  // }
}

function isItemSelected(itemId: string) {
  return newLayer.value.items.some(item => item.id === itemId)
}

function getSelectedItemsText() {
  if (!newLayer.value.items.length) return 'هیچ آیتمی انتخاب نشده است'
  return `${newLayer.value.items.length} آیتم انتخاب شده`
}

function editLayer(layer: HeaderLayer) {
  const layerCopy = JSON.parse(JSON.stringify(layer)) as HeaderLayer
  newLayer.value = normalizeLayer(layerCopy)
  showLayerSettings.value = true
}

function deleteLayer(layerId: string | number) {
  const index = createdLayers.value.findIndex(layer => layer.id === layerId)
  if (index === -1) return
  createdLayers.value.splice(index, 1)
  persistLayersToSession()
}

async function saveLayer() {
  if (!newLayer.value.name.trim()) {
    showToastMessage('نام لایه نمی‌تواند خالی باشد.', 'error')
    return
  }

  if (!newLayer.value.items.length) {
    showToastMessage('حداقل یک آیتم باید برای لایه انتخاب شود.', 'error')
    return
  }

  // ❌ حذف شد: distributeItemWidths - کاربر باید دستی عرض‌ها رو تنظیم کنه
  // distributeItemWidths(newLayer.value.items)

  const isUpdatingExistingLayer = Boolean(newLayer.value.id)
  const layerToStore: HeaderLayer = {
    ...JSON.parse(JSON.stringify(newLayer.value)),
    styleSettings: buildStyleSettings(newLayer.value)
  }

  if (isUpdatingExistingLayer && layerToStore.id) {
    const index = createdLayers.value.findIndex(layer => layer.id === layerToStore.id)
    if (index > -1) {
      createdLayers.value.splice(index, 1, layerToStore)
    }
  } else {
    layerToStore.id = `temp_${Date.now()}`
    createdLayers.value.push(layerToStore)
  }

  persistLayersToSession()

  showToastMessage(isUpdatingExistingLayer ? 'لایه با موفقیت به‌روزرسانی شد.' : 'لایه جدید با موفقیت اضافه شد.', 'success')
  showLayerSettings.value = false
  resetLayerForm()
}

function clearAllLayers() {
  createdLayers.value = []
  if (!isEditing.value && typeof window !== 'undefined') {
    sessionStorage.removeItem('headerLayers')
  }
}

async function saveHeader() {
  if (!headerData.value.name.trim()) {
    showToastMessage('نام هدر نمی‌تواند خالی باشد.', 'error')
    return
  }
  if (!createdLayers.value.length) {
    showToastMessage('حداقل یک لایه برای ساخت هدر لازم است.', 'error')
    return
  }

  const headerPayload: any = {
    name: headerData.value.name,
    description: headerData.value.description,
    page_selection: headerData.value.pageSelection,
    specific_pages: headerData.value.specificPages,
    excluded_pages: headerData.value.excludedPages,
    is_active: headerData.value.isActive,
    layers: createdLayers.value.map(layer => {
      const styleSettings = buildStyleSettings(layer)
      const payload: any = {
        name: layer.name,
        width: layer.width,
        height: layer.height,
        row_count: layer.rowCount,
        color: layer.color,
        opacity: (layer.opacity || 100) / 100,
        direction: layer.direction,
        padding_right: layer.paddingRight,
        padding_left: layer.paddingLeft,
        padding_top: layer.paddingTop,
        padding_bottom: layer.paddingBottom,
        mobile_responsive: layer.mobileResponsive,
        tablet_responsive: layer.tabletResponsive,
        items: JSON.stringify(layer.items),
        style_settings: JSON.stringify(styleSettings),
        enable_border: layer.enableBorder,
        border_position: layer.borderPosition,
        border_color: layer.borderColor,
        border_width: layer.borderWidth,
        border_style: layer.borderStyle,
        enable_shadow: layer.enableShadow,
        shadow_intensity: layer.shadowIntensity,
        shadow_direction: layer.shadowDirection
      }

      if (layer.boxWidths) {
        payload.boxWidths = JSON.stringify(layer.boxWidths)
      }

      if (layer.id && !(typeof layer.id === 'string' && layer.id.startsWith('temp_'))) {
        payload.id = layer.id
      }

      return payload
    })
  }

  try {
    let result: any
    if (isEditing.value) {
      const headerId = parseInt(route.query.id as string, 10)
      const res = await fetch(`/api/admin/header-settings/${headerId}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(headerPayload)
      })
      if (!res.ok) {
        throw new Error(`HTTP error! status: ${res.status}`)
      }
      result = await res.json()
    } else {
      const res = await fetch('/api/admin/header-settings', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(headerPayload)
      })
      if (!res.ok) {
        throw new Error(`HTTP error! status: ${res.status}`)
      }
      result = await res.json()
    }

    if (result?.success) {
      showToastMessage(isEditing.value ? 'هدر با موفقیت به‌روزرسانی شد.' : 'هدر با موفقیت ایجاد شد.', 'success')
      if (typeof window !== 'undefined') {
        sessionStorage.removeItem('headerLayers')
      }
      goBack()
    } else {
      showToastMessage('ذخیره هدر با خطا مواجه شد.', 'error')
    }
  } catch (error: any) {
    console.error('saveHeader error:', error)
    showToastMessage(`ذخیره هدر با خطا مواجه شد: ${error?.message || error}`, 'error')
  }
}

function goBack() {
  router.push('/admin/content/header-management')
}

onMounted(async () => {
  await loadExistingHeader()
  loadLayersFromSession()
})

provide('saveHeader', saveHeader)
provide('goBack', goBack)
provide('saveLayer', saveLayer)
provide('cancelLayer', cancelLayer)
provide('resetLayerForm', resetLayerForm)
provide('openItemsModal', openItemsModal)
provide('closeItemsModal', closeItemsModal)
provide('toggleItem', toggleItem)
provide('isItemSelected', isItemSelected)
provide('getSelectedItemsText', getSelectedItemsText)
provide('clearAllLayers', clearAllLayers)
provide('editLayer', editLayer)
provide('deleteLayer', deleteLayer)
provide('newLayer', newLayer)
provide('createdLayers', createdLayers)
provide('showLayerSettings', showLayerSettings)
provide('isEditing', isEditing)
provide('availableItems', availableItems)
provide('headerData', headerData)
provide('layerData', layerData)
provide('showItemsModal', showItemsModal)
provide('showToastMessage', showToastMessage)
</script>

<style scoped>
.create-header-page {
  min-height: 100vh;
  background: #f3f4f6;
  padding: 20px;
}

.main-content {
  max-width: 1400px;
  margin: 0 auto;
}

.main-content.full-width {
  max-width: none;
}

.header-settings-container {
  margin-bottom: 24px;
}

.layer-edit-layout {
  display: grid;
  grid-template-columns: 320px 1fr;
  gap: 24px;
  margin-bottom: 24px;
  align-items: start;
}

.preview-column {
  background: white;
  border-radius: 8px;
  padding: 24px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  min-height: 400px;
}

.normal-layout {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.toast-notification {
  position: fixed;
  top: 24px;
  right: 24px;
  z-index: 1000;
  padding: 12px 18px;
  border-radius: 8px;
  color: #fff;
  font-weight: 500;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.15);
  transition: opacity 0.3s ease;
}

.toast-success {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
}

.toast-error {
  background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%);
}

@media (max-width: 1024px) {
  .layer-edit-layout {
    grid-template-columns: 1fr;
    gap: 16px;
  }
}

@media (max-width: 768px) {
  .create-header-page {
    padding: 12px;
  }

  .layer-edit-layout {
    grid-template-columns: 1fr;
  }
}
</style>

