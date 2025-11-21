<template>
  <div class="create-footer-page">
    <div class="main-content">
      <PageHeader />
      
      <!-- بخش تنظیمات فوتر -->
      <div class="footer-settings-container">
        <FooterSettingsForm />
      </div>

      <!-- Layout دو ستونه برای ویرایش لایه -->
      <div v-if="showLayerSettings" class="layer-edit-layout">
        <FooterSettingsSidebar />
        <div class="preview-column">
          <FooterPreview />
        </div>
      </div>

      <!-- بخش مدیریت لایه‌ها و پیش‌نمایش عمومی -->
      <div v-else class="normal-layout">
        <FormContainer />
        <FooterPreview />
      </div>
    </div>

    <ItemsSelectionModal />

    <!-- Toast Notification -->
    <div
v-if="showToast" 
         class="fixed top-6 right-4 z-50 transition-all duration-300"
         :class="toastType === 'success' ? 'bg-green-500' : 'bg-red-500'">
      <div class="text-white px-6 py-3 rounded-lg shadow-lg font-iranyekan">
        {{ toastMessage }}
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, provide, ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import FooterPreview from './FooterPreview.vue'
import FooterSettingsForm from './FooterSettingsForm.vue'
import FooterSettingsSidebar from './FooterSettingsSidebar.vue'
import FormContainer from './FormContainer.vue'
import ItemsSelectionModal from './ItemsSelectionModal.vue'
import PageHeader from './PageHeader.vue'

// @ts-ignore
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// Shared state for all components
const footerData = ref({
  name: '',
  description: '',
  pageSelection: 'all', // 'all', 'specific', 'exclude'
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

// آرایه لایه‌های ایجاد شده
const createdLayers = ref<Record<string, unknown>[]>([])

const newLayer = ref<Record<string, unknown>>({
  name: '',
  width: 100,
  height: 50,
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
})

const availableItems = [
  { id: 'logo', name: 'لوگو', path: '/', icon: 'heroicons:building-office-2' },
  { id: 'image', name: 'آپلود عکس', path: '/image', icon: 'heroicons:photo' },
  { id: 'custom', name: 'باکس سفارشی', path: '/custom', icon: 'heroicons:cube' },
  { id: 'menu', name: 'منو', path: '/menu', icon: 'heroicons:bars-3' },
  { id: 'contact', name: 'اطلاعات تماس', path: '/contact', icon: 'heroicons:phone' },
  { id: 'about', name: 'درباره ما', path: '/about', icon: 'heroicons:information-circle' },
  { id: 'working-hours', name: 'ساعات کاری', path: '/working-hours', icon: 'heroicons:clock' },
  { id: 'language', name: 'انتخاب زبان', path: '/language', icon: 'heroicons:globe-alt' },
  { id: 'currency', name: 'انتخاب ارز', path: '/currency', icon: 'heroicons:currency-dollar' },
  { id: 'social', name: 'شبکه‌های اجتماعی', path: '/social', icon: 'heroicons:share' },
  { id: 'trust', name: 'نشان‌های اعتماد', path: '/trust', icon: 'heroicons:shield-check' }
]

const route = useRoute()
const router = useRouter()
const isEditing = computed(() => !!route.query.id)

// Toast notification state
const showToast = ref(false)
const toastMessage = ref('')
const toastType = ref('success')

// نمایش پیام toast
const showToastMessage = (message, type = 'success') => {
  toastMessage.value = message
  toastType.value = type
  showToast.value = true
  
  setTimeout(() => {
    showToast.value = false
  }, 3000)
}

async function loadExistingFooter() {
  if (isEditing.value) {
    try {
      const footerId = parseInt(route.query.id as string)
      
  const response = await $fetch(`/api/admin/footer-settings/${footerId}`) as Record<string, unknown>
      
      if (response && typeof response === 'object' && 'success' in response && response.success && 'data' in response && response.data) {
        const existingFooter = response.data as Record<string, unknown>
        
        footerData.value = {
          name: existingFooter.name || '',
          description: existingFooter.description || '',
          pageSelection: existingFooter.page_selection || 'all',
          specificPages: existingFooter.specific_pages || '',
          excludedPages: existingFooter.excluded_pages || '',
          isActive: existingFooter.is_active !== undefined ? existingFooter.is_active : true
        }
        
        createdLayers.value = []
        
        if (existingFooter.layers && existingFooter.layers.length > 0) {
          createdLayers.value = existingFooter.layers.map(l => ({
            ...l,
            name: l.name || '',
            width: l.width || 100,
            height: l.height || 50,
            rowCount: l.row_count || 1,
            color: l.color || '#ffffff',
            opacity: (l.opacity || 1.0) * 100,
            showSeparator: l.show_separator || false,
            separatorType: l.separator_type || 'solid',
            separatorColor: l.separator_color || '#e9ecef',
            separatorOpacity: (l.separator_opacity || 0.2) * 100,
            separatorWidth: l.separator_width || 1,
            items: Array.isArray(l.items)
              ? l.items
              : (typeof l.items === 'string' && l.items.trim() !== ''
                  ? JSON.parse(l.items)
                  : []),
            boxWidths: Array.isArray(l.boxWidths) ? l.boxWidths : (typeof l.boxWidths === 'string' && l.boxWidths ? JSON.parse(l.boxWidths) : undefined)
          }))
        }
      } else {
        console.error('پاسخ API موفق نبود:', response)
      }
    } catch (error) {
      console.error('خطا در بارگذاری فوتر موجود:', error)
    }
  } else {
    createdLayers.value = []
  }
}

onMounted(async () => {
  await loadExistingFooter()
  
  if (!isEditing.value) {
    sessionStorage.removeItem('footerLayers')
    createdLayers.value = []
    footerData.value = {
      name: '',
      description: '',
      pageSelection: 'all',
      specificPages: '',
      excludedPages: '',
      isActive: true
    }
  }
})

watch(() => newLayer.value, () => {
}, { deep: true })

watch(() => newLayer.value.color, (_newColor) => {
}, { immediate: true })

watch(() => newLayer.value.opacity, (_newOpacity) => {
}, { immediate: true })

// async function loadLayersFromDatabase() {
//   try {
//     if (isEditing.value) {
//       try {
//         const footerId = parseInt(route.query.id as string)
        
//   const response = await $fetch(`/api/admin/footer-settings/${footerId}/layers`) as any
        
//         if (response && response.data && Array.isArray(response.data)) {
          
//           const layersFromBackend = response.data.map(layer => {
//             return {
//               id: layer.id,
//               name: layer.name,
//               width: layer.width,
//               height: layer.height,
//               rowCount: layer.row_count,
//               color: layer.color,
//               opacity: (layer.opacity || 1.0) * 100,
//               showSeparator: layer.show_separator || false,
//               separatorType: layer.separator_type || 'solid',
//               separatorColor: layer.separator_color || '#e9ecef',
//               separatorOpacity: (layer.separator_opacity || 0.2) * 100,
//               separatorWidth: layer.separator_width || 1,
//               items: typeof layer.items === 'string' ? JSON.parse(layer.items) : layer.items || []
//             }
//           })
          
//           createdLayers.value = layersFromBackend
//           return true
//         } else {
//         }
//       } catch (error) {
//         console.error('❌ خطا در بارگذاری لایه‌ها از دیتابیس:', error)
//       }
      
//       return true
//     } else {
//       const savedLayers = sessionStorage.getItem('footerLayers')
      
//       if (savedLayers && savedLayers.trim() !== '') {
//         const parsedLayers = JSON.parse(savedLayers)
        
//         if (Array.isArray(parsedLayers) && parsedLayers.length > 0) {
//           createdLayers.value = parsedLayers
//           return true
//         }
//       }
      
//       return false
//     }
//   } catch (error) {
//     console.error('❌ خطا در بارگذاری لایه‌ها از session storage:', error)
//     return false
//   }
// }

async function saveLayer() {
  if (!newLayer.value.name.trim()) {
    showToastMessage('لطفاً نام لایه را وارد کنید!', 'error')
    return
  }
  
  try {
    if (newLayer.value.id) {
      const idx = createdLayers.value.findIndex(l => l.id === newLayer.value.id)
      if (idx > -1) {
        createdLayers.value[idx] = {
          ...JSON.parse(JSON.stringify(newLayer.value)),
          updatedAt: new Date().toISOString()
        }
      }
    } else {
      const layerToSave = {
        ...JSON.parse(JSON.stringify(newLayer.value)),
        id: `temp_${Date.now()}`,
        createdAt: new Date().toISOString()
      }
      createdLayers.value.push(layerToSave)
    }

    if (!isEditing.value) {
      sessionStorage.setItem('footerLayers', JSON.stringify(createdLayers.value))
    }

    if (newLayer.value.id) {
      showToastMessage('✅ لایه با موفقیت به‌روزرسانی شد!', 'success')
      showLayerSettings.value = false
    } else {
      resetLayerForm()
      showToastMessage('✅ لایه با موفقیت ذخیره شد!', 'success')
      showLayerSettings.value = false
    }
  } catch (error) {
    console.error('خطا در ذخیره لایه:', error)
    showToastMessage('خطا در ذخیره لایه: ' + error.message, 'error')
  }
}

async function deleteLayer(layerId) {
  const index = createdLayers.value.findIndex(l => l.id === layerId)
  if (index > -1) {
    createdLayers.value.splice(index, 1)
    if (!isEditing.value) {
      sessionStorage.setItem('footerLayers', JSON.stringify(createdLayers.value))
    }
  }
}

function cancelLayer() {
  showLayerSettings.value = false
  resetLayerForm()
}

function resetLayerForm() {
  newLayer.value = {
    name: '',
    width: 100,
    height: 50,
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
  }
  delete newLayer.value.id
}

function openItemsModal() {
  showItemsModal.value = true
}

function closeItemsModal() {
  showItemsModal.value = false
}

function toggleItem(itemId) {
  const item = availableItems.find(ai => ai.id === itemId)
  if (!item) return
  
  const existingIndex = newLayer.value.items.findIndex(i => 
    (typeof i === 'string' && i === itemId) || 
    (typeof i === 'object' && i.id === itemId)
  )
  
  if (existingIndex > -1) {
    newLayer.value.items.splice(existingIndex, 1)
    const totalAfter = newLayer.value.items.length
    if (totalAfter) {
      newLayer.value.items.forEach(it => {
        it.width = Math.round((100 / totalAfter) * 100) / 100
      })
    }
  } else {
    newLayer.value.items.push({
      id: item.id,
      name: item.name,
      path: item.path,
      icon: item.icon,
      align: 'right',
      bgColor: 'transparent',
      width: 0,
      items: [item.id]
    })
    const total = newLayer.value.items.length
    newLayer.value.items.forEach(it => {
      it.width = Math.round((100 / total) * 100) / 100
    })
  }
  
  const totalWidth = newLayer.value.items.reduce((sum, item) => sum + (item.width || 0), 0)
  if (Math.abs(totalWidth - 100) > 0.01) {
    const lastIndex = newLayer.value.items.length - 1
    if (lastIndex >= 0) {
      const currentTotal = newLayer.value.items.reduce((sum, item, index) => 
        index === lastIndex ? sum : sum + (item.width || 0), 0)
      newLayer.value.items[lastIndex].width = Math.round((100 - currentTotal) * 100) / 100
    }
  }
}

function isItemSelected(itemId) {
  return newLayer.value.items.some(item => 
    (typeof item === 'string' && item === itemId) || 
    (typeof item === 'object' && item.id === itemId)
  )
}

function getSelectedItemsText() {
  if (newLayer.value.items.length === 0) {
    return 'انتخاب ایتم ها'
  }
  return `${newLayer.value.items.length} ایتم انتخاب شده`
}

async function clearAllLayers() {
  createdLayers.value = []
  sessionStorage.removeItem('footerLayers')
}

async function saveFooter() {
  if (!footerData.value.name.trim()) {
    showToastMessage('لطفاً نام فوتر را وارد کنید!', 'error')
    return
  }
  if (createdLayers.value.length === 0) {
    showToastMessage('لطفاً حداقل یک لایه ایجاد کنید!', 'error')
    return
  }

  const footerPayload: Record<string, unknown> = {
    name: footerData.value.name,
    description: footerData.value.description,
    page_selection: footerData.value.pageSelection,
    specific_pages: footerData.value.specificPages,
    excluded_pages: footerData.value.excludedPages,
    is_active: footerData.value.isActive,
    layers: createdLayers.value.map(layer => {
      // ساخت styleSettings به صورت JSONB
      const styleSettings = {
        border: {
          enabled: layer.enableBorder || false,
          position: layer.borderPosition || 'all',
          color: layer.borderColor || '#e5e7eb',
          width: layer.borderWidth || 1,
          style: layer.borderStyle || 'solid'
        },
        shadow: {
          enabled: layer.enableShadow || false,
          intensity: layer.shadowIntensity || 'md',
          direction: layer.shadowDirection || 'top'
        },
        layout: {
          direction: layer.direction || 'rtl',
          mobileResponsive: layer.mobileResponsive !== undefined ? layer.mobileResponsive : true,
          tabletResponsive: layer.tabletResponsive !== undefined ? layer.tabletResponsive : true
        }
      }
      
      const l: Record<string, unknown> = {
        name: layer.name,
        width: layer.width,
        height: layer.height,
        row_count: layer.rowCount,
        color: layer.color,
        opacity: layer.opacity / 100.0,
        showSeparator: layer.showSeparator,
        separatorType: layer.separatorType,
        separatorColor: layer.separatorColor,
        separatorOpacity: layer.separatorOpacity / 100.0,
        separatorWidth: layer.separatorWidth,
        items: JSON.stringify(layer.items),
        styleSettings: JSON.stringify(styleSettings)
      }
      if (layer.boxWidths) l.box_widths = JSON.stringify(layer.boxWidths)
      if (typeof layer.id === 'string' && layer.id.startsWith('temp_')) delete l.id
      return l
    })
  }



  try {
    let response: unknown
    if (isEditing.value) {
      const footerId = parseInt(route.query.id as string)
  response = await fetch(`/api/admin/footer-settings/${footerId}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(footerPayload)
      })
      
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`)
      }
      
      response = await response.json()
    } else {
      response = await fetch('/api/admin/footer-settings', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(footerPayload)
      })
      
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`)
      }
      
      response = await response.json()
    }

    if (response?.success) {
      showToastMessage(isEditing.value ? 'فوتر با موفقیت بروزرسانی شد!' : 'فوتر با موفقیت ذخیره شد!', 'success')
      sessionStorage.removeItem('footerLayers')
      goBack()
    } else {
      showToastMessage('خطا در ذخیره فوتر', 'error')
    }
  } catch (error: unknown) {
    console.error('saveFooter error:', error)
    const err = error as { data?: { message?: string }, message?: string }
    showToastMessage('خطا در ذخیره فوتر: ' + (err?.data?.message || err?.message || error), 'error')
  }
}

function goBack() {
  router.push('/admin/content/footer-management')
}

provide('saveFooter', saveFooter)
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
provide('footerData', footerData)
provide('layerData', layerData)
provide('showItemsModal', showItemsModal)

function editLayer(layer) {
  const layerCopy = JSON.parse(JSON.stringify(layer))
  
  Object.keys(layerCopy).forEach(key => {
    if (key in newLayer.value) {
      newLayer.value[key] = layerCopy[key]
    }
  })
  
  showLayerSettings.value = true
}
</script>

<style scoped>
.create-footer-page {
  min-height: 100vh;
  background: #f3f4f6;
  padding: 20px;
}

.main-content {
  max-width: 1400px;
  margin: 0 auto;
}

.footer-settings-container {
  margin-bottom: 24px;
}

/* Layout دو ستونه برای ویرایش لایه */
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

/* Layout عادی */
.normal-layout {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

/* Responsive styles */
@media (max-width: 1024px) {
  .layer-edit-layout {
    grid-template-columns: 1fr;
    gap: 16px;
  }
}

@media (max-width: 768px) {
  .create-footer-page {
    padding: 12px;
  }
  
  .layer-edit-layout {
    grid-template-columns: 1fr;
  }
}

/* Toast notification styles */
.font-iranyekan {
  font-family: 'Yekan', 'Tahoma', sans-serif;
}
</style>
