<template>
  <div class="edit-footer-page">
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
    <div v-if="showToast" 
         class="fixed top-6 right-4 z-50 transition-all duration-300"
         :class="toastType === 'success' ? 'bg-green-500' : 'bg-red-500'">
      <div class="text-white px-6 py-3 rounded-lg shadow-lg font-iranyekan">
        {{ toastMessage }}
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, provide, computed, onMounted, watch } from 'vue'
import FooterSettingsSidebar from '../create/FooterSettingsSidebar.vue'
import PageHeader from '../create/PageHeader.vue'
import FooterPreview from '../create/FooterPreview.vue'
import FormContainer from '../create/FormContainer.vue'
import ItemsSelectionModal from '../create/ItemsSelectionModal.vue'
import FooterSettingsForm from '../create/FooterSettingsForm.vue'
import { useRoute, useRouter } from 'vue-router'

// @ts-ignore
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// Route and Router
const route = useRoute()
const router = useRouter()
const footerId = route.params.id as string

// Shared state for all components
const footerData = ref({
  name: '',
  description: '',
  pageSelection: 'all', // 'all', 'specific', 'exclude'
  specificPages: '',
  excludedPages: '',
  isActive: true
})

// Debug: watch تغییرات footerData
watch(footerData, (newVal) => {
  console.log('🔄 footerData در صفحه ویرایش تغییر کرد:', newVal)
}, { deep: true })

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
const createdLayers = ref<any[]>([])

const newLayer = ref<any>({
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
  separatorType: 'line',
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
  { id: 'language', name: 'انتخاب زبان', path: '/language', icon: 'heroicons:globe-alt' },
  { id: 'social', name: 'شبکه‌های اجتماعی', path: '/social', icon: 'heroicons:share' },
  { id: 'contact', name: 'اطلاعات تماس', path: '/contact', icon: 'heroicons:envelope' },
  { id: 'about', name: 'درباره ما', path: '/about', icon: 'heroicons:information-circle' },
  { id: 'working-hours', name: 'ساعات کاری', path: '/working-hours', icon: 'heroicons:clock' },
  { id: 'newsletter', name: 'خبرنامه', path: '/newsletter', icon: 'heroicons:newspaper' },
  { id: 'copyright', name: 'کپی‌رایت', path: '/copyright', icon: 'heroicons:document-text' },
  { id: 'links', name: 'لینک‌های مفید', path: '/links', icon: 'heroicons:link' },
  { id: 'trust', name: 'نشان‌های اعتماد', path: '/trust', icon: 'heroicons:shield-check' },
  { id: 'back-to-top', name: 'بازگشت به بالا', path: '/back-to-top', icon: 'heroicons:arrow-up' }
]

// Toast state
const showToast = ref(false)
const toastMessage = ref('')
const toastType = ref<'success' | 'error'>('success')

// Computed
const isEditing = computed(() => footerId && footerId !== 'create')

// Methods
const showToastMessage = (message: string, type: 'success' | 'error' = 'success') => {
  toastMessage.value = message
  toastType.value = type
  showToast.value = true
  
  setTimeout(() => {
    showToast.value = false
  }, 3000)
}

// توابع مورد نیاز کامپوننت‌های فرزند
const editLayer = (layer: any) => {
  // کپی کردن تنظیمات لایه به newLayer برای ویرایش
  Object.assign(newLayer.value, layer)
  showLayerSettings.value = true
}

const deleteLayer = (layerId: string) => {
  const index = createdLayers.value.findIndex(layer => layer.id === layerId)
  if (index !== -1) {
    createdLayers.value.splice(index, 1)
    showToastMessage('لایه با موفقیت حذف شد', 'success')
  }
}

const clearAllLayers = () => {
  if (confirm('آیا مطمئن هستید که می‌خواهید همه لایه‌ها را پاک کنید؟')) {
    createdLayers.value = []
    showToastMessage('همه لایه‌ها پاک شدند', 'success')
  }
}

const saveLayer = () => {
  if (!newLayer.value.name.trim()) {
    showToastMessage('نام لایه الزامی است', 'error')
    return
  }

  // اگر لایه جدید است، ID جدید ایجاد کن
  if (!newLayer.value.id || newLayer.value.id.toString().startsWith('temp_')) {
    newLayer.value.id = `temp_${Date.now()}`
    createdLayers.value.push({ ...newLayer.value })
  } else {
    // اگر لایه موجود است، آن را آپدیت کن
    const index = createdLayers.value.findIndex(layer => layer.id === newLayer.value.id)
    if (index !== -1) {
      createdLayers.value[index] = { ...newLayer.value }
    }
  }

  // ریست کردن newLayer
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

  showLayerSettings.value = false
  showToastMessage('لایه با موفقیت ذخیره شد', 'success')
}

const cancelLayerEdit = () => {
  showLayerSettings.value = false
  // ریست کردن newLayer
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
}

const getSelectedItemsText = (items: any[]) => {
  if (!items || items.length === 0) return 'هیچ آیتمی انتخاب نشده'
  
  const itemNames = items.map(item => {
    if (typeof item === 'string') return item
    if (item.name) return item.name
    if (item.id) return item.id
    return 'آیتم ناشناس'
  })
  
  return itemNames.join(', ')
}

const openItemsModal = () => {
  showItemsModal.value = true
}

const closeItemsModal = () => {
  showItemsModal.value = false
}

const addItemToLayer = (item: any) => {
  if (!newLayer.value.items) {
    newLayer.value.items = []
  }
  
  // بررسی اینکه آیا آیتم قبلاً اضافه شده یا نه
  const exists = newLayer.value.items.some(existingItem => 
    (typeof existingItem === 'string' && existingItem === item.id) ||
    (typeof existingItem === 'object' && existingItem.id === item.id)
  )
  
  if (!exists) {
    newLayer.value.items.push(item)
    showToastMessage(`${item.name} به لایه اضافه شد`, 'success')
  } else {
    showToastMessage('این آیتم قبلاً در لایه موجود است', 'error')
  }
}

const removeItemFromLayer = (itemIndex: number) => {
  if (newLayer.value.items && newLayer.value.items[itemIndex]) {
    const removedItem = newLayer.value.items[itemIndex]
    newLayer.value.items.splice(itemIndex, 1)
    showToastMessage(`${typeof removedItem === 'string' ? removedItem : removedItem.name} از لایه حذف شد`, 'success')
  }
}

// توابع مورد نیاز ItemsSelectionModal
const toggleItem = (itemId: string) => {
  if (!newLayer.value.items) {
    newLayer.value.items = []
  }
  
  const itemIndex = newLayer.value.items.findIndex(item => 
    (typeof item === 'string' && item === itemId) ||
    (typeof item === 'object' && item.id === itemId)
  )
  
  if (itemIndex !== -1) {
    // حذف آیتم اگر قبلاً انتخاب شده
    newLayer.value.items.splice(itemIndex, 1)
    showToastMessage('آیتم از انتخاب حذف شد', 'success')
  } else {
    // اضافه کردن آیتم جدید
    const item = availableItems.find(ai => ai.id === itemId)
    if (item) {
      newLayer.value.items.push(item)
      showToastMessage(`${item.name} انتخاب شد`, 'success')
    }
  }
}

const isItemSelected = (itemId: string): boolean => {
  if (!newLayer.value.items) return false
  
  return newLayer.value.items.some(item => 
    (typeof item === 'string' && item === itemId) ||
    (typeof item === 'object' && item.id === itemId)
  )
}

const loadFooter = async () => {
  try {
    console.log('🔄 شروع بارگذاری فوتر...')
    console.log('📝 footerId:', footerId)
    
    if (!footerId || footerId === 'create') {
      console.log('❌ footerId نامعتبر است')
      return
    }

    console.log('🌐 درخواست API به:', `/api/admin/footer-settings/${footerId}`)
    const response = await fetch(`/api/admin/footer-settings/${footerId}`)
    
    console.log('📡 Response status:', response.status)
    console.log('📡 Response headers:', Object.fromEntries(response.headers.entries()))
    
    if (!response.ok) {
      throw new Error(`خطای HTTP: ${response.status}`)
    }

    const data = await response.json()
    console.log('📦 Response data:', data)
    
    if (data.success && data.data) {
      const footer = data.data
      console.log('✅ فوتر دریافت شد:', footer)
      console.log('🔍 ساختار داده‌های backend:', {
        name: footer.name,
        description: footer.description,
        page_selection: footer.page_selection,
        is_active: footer.is_active,
        layers_count: footer.layers?.length || 0,
        keys: Object.keys(footer)
      })
      
      // تبدیل داده‌ها از backend به frontend
      footerData.value = {
        name: footer.name || '',
        description: footer.description || '',
        pageSelection: footer.page_selection || 'all',
        specificPages: footer.specific_pages || '',
        excludedPages: footer.excluded_pages || '',
        isActive: footer.is_active !== undefined ? footer.is_active : true
      }
      
      console.log('🔄 footerData به‌روزرسانی شد:', footerData.value)

      // تبدیل لایه‌ها
      const layers = footer.layers || []
      if (layers && Array.isArray(layers)) {
        console.log('🔧 لایه‌ها دریافت شدند:', layers)
        createdLayers.value = layers.map((layer: any) => {
          // Parse styleSettings از JSONB
          let styleSettings: any = { border: {}, shadow: {}, layout: {} }
          if (layer.styleSettings) {
            try {
              styleSettings = typeof layer.styleSettings === 'string' 
                ? JSON.parse(layer.styleSettings) 
                : layer.styleSettings
            } catch (e) {
              console.error('خطا در parse کردن styleSettings:', e)
            }
          }
          
          return {
            id: layer.id || `temp_${Date.now()}`,
            name: layer.name || '',
            width: layer.width || 100,
            height: layer.height || 50,
            rowCount: layer.row_count || 1,
            color: layer.color || '#ffffff',
            opacity: (layer.opacity || 1.0) * 100,
            
            // Border از styleSettings
            enableBorder: styleSettings.border?.enabled || false,
            borderPosition: styleSettings.border?.position || 'all',
            borderColor: styleSettings.border?.color || '#e5e7eb',
            borderWidth: styleSettings.border?.width || 1,
            borderStyle: styleSettings.border?.style || 'solid',
            
            // Shadow از styleSettings
            enableShadow: styleSettings.shadow?.enabled || false,
            shadowIntensity: styleSettings.shadow?.intensity || 'md',
            shadowDirection: styleSettings.shadow?.direction || 'top',
            
            // Layout از styleSettings
            direction: styleSettings.layout?.direction || 'rtl',
            mobileResponsive: styleSettings.layout?.mobileResponsive !== undefined ? styleSettings.layout.mobileResponsive : true,
            tabletResponsive: styleSettings.layout?.tabletResponsive !== undefined ? styleSettings.layout.tabletResponsive : true,
            
            // Separator (قدیمی - از فیلدهای مستقیم)
            showSeparator: layer.showSeparator || false,
            separatorType: layer.separatorType || 'solid',
            separatorColor: layer.separatorColor || '#e9ecef',
            separatorOpacity: (layer.separatorOpacity || 0.2) * 100,
            separatorWidth: layer.separatorWidth || 1,
            
            items: layer.items ? (typeof layer.items === 'string' ? JSON.parse(layer.items) : layer.items) : [],
            paddingRight: layer.padding_right || 0,
            paddingLeft: layer.padding_left || 0,
            boxWidths: layer.box_widths ? (typeof layer.box_widths === 'string' ? JSON.parse(layer.box_widths) : layer.box_widths) : undefined
          }
        })
        console.log('✅ لایه‌ها به‌روزرسانی شدند:', createdLayers.value)
      } else {
        console.log('ℹ️ هیچ لایه‌ای یافت نشد. فیلدهای موجود:', Object.keys(footer))
        console.log('🔍 footer.layers:', footer.layers)
        console.log('🔍 footer.Layers:', footer.Layers)
      }
    } else {
      console.log('❌ داده‌های نامعتبر:', data)
      throw new Error(data.message || 'خطا در دریافت اطلاعات فوتر')
    }
  } catch (err: any) {
    console.error('💥 خطا در بارگذاری فوتر:', err)
    console.error('💥 Error details:', {
      message: err.message,
      stack: err.stack
    })
    showToastMessage(err.message || 'خطا در بارگذاری اطلاعات فوتر', 'error')
  }
}

const saveFooter = async () => {
  try {
    // ساخت payload با snake_case
    const footerPayload: any = {
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
        
        const l: any = {
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
          styleSettings: JSON.stringify(styleSettings),
          padding_right: layer.paddingRight,
          padding_left: layer.paddingLeft
        }
        if (layer.boxWidths) l.box_widths = JSON.stringify(layer.boxWidths)
        if (typeof layer.id === 'string' && layer.id.startsWith('temp_')) delete l.id
        else if (layer.id && layer.id !== 'temp_' + Date.now()) l.id = layer.id
        return l
      })
    }

    console.log('Payload for update:', footerPayload)

    const response = await fetch(`/api/admin/footer-settings/${footerId}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(footerPayload)
    })

    if (!response.ok) {
      const errorData = await response.json().catch(() => ({}))
      throw new Error(errorData.message || `خطای HTTP: ${response.status}`)
    }

    const result = await response.json()
    
    if (result.success) {
      showToastMessage('فوتر با موفقیت ویرایش شد', 'success')
      setTimeout(() => {
        goBack()
      }, 1500)
    } else {
      throw new Error(result.message || 'خطا در ویرایش فوتر')
    }
  } catch (err: any) {
    console.error('خطا در ذخیره فوتر:', err)
    showToastMessage(err.message || 'خطا در ذخیره فوتر', 'error')
  }
}

const goBack = () => {
  router.push('/admin/content/footer-management')
}

// Provide data to child components
provide('footerData', footerData)
provide('layerData', layerData)
provide('createdLayers', createdLayers)
provide('newLayer', newLayer)
provide('availableItems', availableItems)
provide('showItemsModal', showItemsModal)
provide('showLayerSettings', showLayerSettings)
provide('saveFooter', saveFooter)
provide('isEditing', isEditing)
provide('goBack', goBack)

// Provide functions to child components
provide('editLayer', editLayer)
provide('deleteLayer', deleteLayer)
provide('clearAllLayers', clearAllLayers)
provide('saveLayer', saveLayer)
provide('cancelLayerEdit', cancelLayerEdit)
provide('cancelLayer', cancelLayerEdit) // FooterSettingsSidebar از cancelLayer استفاده می‌کند
provide('getSelectedItemsText', getSelectedItemsText)
provide('openItemsModal', openItemsModal)
provide('closeItemsModal', closeItemsModal)
provide('addItemToLayer', addItemToLayer)
provide('removeItemFromLayer', removeItemFromLayer)
provide('toggleItem', toggleItem)
provide('isItemSelected', isItemSelected)

// Lifecycle
onMounted(() => {
  console.log('🚀 صفحه ویرایش فوتر بارگذاری شد')
  console.log('📝 footerId در onMounted:', footerId)
  
  if (footerId && footerId !== 'create') {
    console.log('✅ شروع بارگذاری فوتر...')
    loadFooter()
  } else {
    console.log('❌ footerId نامعتبر برای بارگذاری')
  }
})
</script>

<style scoped>
.edit-footer-page {
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

/* Responsive */
@media (max-width: 1024px) {
  .layer-edit-layout {
    grid-template-columns: 1fr;
    gap: 16px;
  }
}

/* Toast styles */
.toast-enter-active, .toast-leave-active {
  transition: all 0.3s ease;
}

.toast-enter-from, .toast-leave-to {
  opacity: 0;
  transform: translateY(-20px);
}
</style>
