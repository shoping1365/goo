<template>
    <header v-if="shouldShowHeader && activeHeader && activeHeader.layers" :style="headerStyle as any" class="sticky-header">

  <!-- حذف نمایش اطلاعات دیباگ لایه -->
    <div v-for="layer in activeHeader.layers" :key="layer.id" :style="getLayerStyle(layer)">
      <template v-for="(item, idx) in getLayerItems(layer)" :key="item.id || idx">
        <div :style="getItemStyle(item)" class="header-item-wrapper">
          <component
            :is="resolveWidgetComponent(item)"
            v-if="resolveWidgetComponent(item)"
            v-bind="{
              ...(typeof item === 'object' ? (item.props || {}) : {}),
              paddingRight: typeof item.paddingRight === 'string' ? parseInt(item.paddingRight) || 0 : item.paddingRight || 0,
              paddingLeft: typeof item.paddingLeft === 'string' ? parseInt(item.paddingLeft) || 0 : item.paddingLeft || 0,
              align: item.align,
              // ارسال props مخصوص عکس
              ...(item.id === 'image' ? {
                imageUrl: item.imageUrl,
                imageName: item.imageName,
                imageId: item.imageId
              } : {})
            }"
          />

          <NuxtLink
            v-else-if="item.path"
            :to="item.path"
            class="header-item-link"
          >
            {{ item.text || item.name || item.type || item }}
          </NuxtLink>

          <div v-else-if="typeof item === 'string'" class="header-item">
            {{ item }}
          </div>
        </div>
        
        <!-- جداکننده بین آیتم‌ها -->
        <div 
          v-if="layer.showSeparator && idx < getLayerItems(layer).length - 1" 
          class="header-separator"
          :style="getSeparatorStyle(layer)"
        ></div>
      </template>
    </div>
  </header>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import { useRoute, useState } from 'nuxt/app'
import { usePublicHeaders } from '~/composables/usePublicHeaders'
import { useDeviceDetection } from '~/composables/useDeviceDetection'

const route = useRoute()
const { loadHeaders, getActiveHeader, getHeaderForPage } = usePublicHeaders()
const { isDesktop } = useDeviceDetection()
const activeHeader = ref<any>(null)
const headerRef = ref<HTMLElement>()
const isSticky = ref(false)



// بررسی اینکه آیا هدر باید نمایش داده شود یا نه
const shouldShowHeader = computed(() => {
  // هدر دسکتاپ فقط در دستگاه‌های دسکتاپ نمایش داده می‌شود
  if (!isDesktop.value) {
    return false
  }

  // استفاده از state میدل‌ور
  const headerDisplayState = useState('headerDisplayState', () => ({
    shouldShow: true,
    headerData: null
  }))
  
  // اگر state میدل‌ور موجود باشد، از آن استفاده کن
  if (headerDisplayState.value && headerDisplayState.value.shouldShow !== undefined) {
    return headerDisplayState.value.shouldShow
  }
  
  // اگر activeHeader موجود باشد، هدر را نمایش بده
  if (activeHeader.value && activeHeader.value.layers) {
    return true
  }
  
  // در حالت پیش‌فرض، هدر را نمایش بده
  return true
})

const headerStyle = computed(() => {
  let headerHeight = '80px'
  
  if (activeHeader.value && activeHeader.value.layers) {
    const totalHeight = activeHeader.value.layers.reduce((total, layer) => total + (layer.height || 80), 0)
    headerHeight = totalHeight + 'px'
  }
  
  return {
    height: headerHeight
  }
})

function getLayerStyle(layer: any) {
  // سبک پایه هدر را تعریف می‌کنیم و سپس هر استایل سفارشی درج‌شده در لایه را روی آن اعمال می‌کنیم
  const baseStyle: Record<string, any> = {
    height: (layer.height ?? 80) + 'px',
    backgroundColor: layer.color || '#ffffff',
    width: layer.width ? layer.width + '%' : '100%',
    display: 'flex',
    alignItems: 'stretch',
    justifyContent: getLayerJustifyContent(layer),
    direction: layer.direction || 'rtl',
    textAlign: layer.direction === 'ltr' ? ('left' as const) : ('right' as const),
    opacity: layer.opacity || 1,
    transition: 'all 0.3s ease'
  }

  Object.assign(baseStyle, buildBorderStyles(resolveBorderConfig(layer)))

  const shadowStyles = resolveShadowStyles(layer)
  if (shadowStyles.boxShadow) {
    baseStyle.boxShadow = shadowStyles.boxShadow
    baseStyle.position = shadowStyles.position
    baseStyle.zIndex = shadowStyles.zIndex
  } else {
    baseStyle.boxShadow = 'none'
  }

  // لایه ممکن است یک آبجکت استایل یا یک رشته JSON باشد
  let customStyle: Record<string, any> = {}
  if (layer.style) {
    if (typeof layer.style === 'string') {
      try {
        customStyle = JSON.parse(layer.style)
      } catch {
        customStyle = {}
      }
    } else if (typeof layer.style === 'object') {
      customStyle = layer.style
    }
  }

  return { ...baseStyle, ...customStyle }
}

function resolveToggle(value: unknown): boolean {
  if (typeof value === 'boolean') return value
  if (typeof value === 'number') return value === 1
  if (typeof value === 'string') {
    const normalized = value.trim().toLowerCase()
    return normalized === 'true' || normalized === '1' || normalized === 'yes'
  }
  return false
}

/**
 * تعیین چینش لایه بر اساس آیتم‌ها و تنظیمات
 */
function getLayerJustifyContent(layer: any): string {
  // اگر لایه دارای آیتم‌هایی با چینش مشخص باشد
  if (layer?.items) {
    try {
      const items = typeof layer.items === 'string'
        ? JSON.parse(layer.items)
        : Array.isArray(layer.items)
          ? layer.items
          : []

      if (Array.isArray(items) && items.length > 0) {
        const firstItem = items[0]
        const firstAlign = firstItem?.align || 'center'

        const allSameAlign = items.every((item: any) => (item?.align || 'center') === firstAlign)

        if (allSameAlign) {
          switch (firstAlign) {
            case 'left':
              return 'flex-start'
            case 'center':
              return 'center'
            case 'right':
              return 'flex-end'
            default:
              return 'center'
          }
        }

        return 'space-between'
      }
    } catch (error) {
      // خطا در پردازش آیتم‌های لایه
    }
  }
  
  // اگر آیتم‌ها وجود نداشته باشند، از direction استفاده می‌کنیم
  return layer.direction === 'ltr' ? 'flex-start' : 'flex-end'
}

function parseStyleSettings(layer: any) {
  const raw = layer?.styleSettings ?? layer?.style_settings
  if (!raw) return null
  if (typeof raw === 'string') {
    try {
      return JSON.parse(raw)
    } catch (error) {
      console.warn('Header.vue - failed to parse styleSettings:', error)
      return null
    }
  }
  if (typeof raw === 'object') return raw
  return null
}

function resolveBorderConfig(layer: any) {
  const styleSettings = parseStyleSettings(layer) || {}
  const border = styleSettings.border || {}
  const enabled = resolveToggle(
    layer?.borderEnabled ??
    layer?.border_enabled ??
    layer?.enableBorder ??
    layer?.enable_border ??
    border.enabled
  )
  return {
    enabled,
    position: layer?.borderPosition ?? layer?.border_position ?? border.position ?? 'all',
    color: layer?.borderColor ?? layer?.border_color ?? border.color ?? '#e5e7eb',
    width: Number(layer?.borderWidth ?? layer?.border_width ?? border.width ?? 1) || 1,
    style: layer?.borderStyle ?? layer?.border_style ?? border.style ?? 'solid'
  }
}

function buildBorderStyles(config: { enabled: boolean; position: string; color: string; width: number; style: string }) {
  if (!config.enabled) return {}
  const borderValue = `${config.width}px ${config.style} ${config.color}`
  const styles: Record<string, string> = {
    borderRadius: '8px'
  }

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

const SHADOW_PRESETS: Record<string, Record<string, string>> = {
  sm: {
    top: '0 -2px 4px -1px rgba(0, 0, 0, 0.06)',
    bottom: '0 2px 4px -1px rgba(0, 0, 0, 0.06)',
    both: '0 -2px 4px -1px rgba(0, 0, 0, 0.06), 0 2px 4px -1px rgba(0, 0, 0, 0.06)'
  },
  md: {
    top: '0 -4px 6px -1px rgba(0, 0, 0, 0.1), 0 -2px 4px -1px rgba(0, 0, 0, 0.06)',
    bottom: '0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06)',
    both: '0 -4px 6px -1px rgba(0, 0, 0, 0.1), 0 4px 6px -1px rgba(0, 0, 0, 0.1)'
  },
  lg: {
    top: '0 -10px 15px -3px rgba(0, 0, 0, 0.1), 0 -4px 6px -2px rgba(0, 0, 0, 0.05)',
    bottom: '0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05)',
    both: '0 -10px 15px -3px rgba(0, 0, 0, 0.1), 0 10px 15px -3px rgba(0, 0, 0, 0.1)'
  },
  xl: {
    top: '0 -20px 25px -5px rgba(0, 0, 0, 0.1), 0 -10px 10px -5px rgba(0, 0, 0, 0.04)',
    bottom: '0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04)',
    both: '0 -20px 25px -5px rgba(0, 0, 0, 0.1), 0 20px 25px -5px rgba(0, 0, 0, 0.1)'
  }
}

function resolveShadowConfig(layer: any) {
  const styleSettings = parseStyleSettings(layer) || {}
  const shadow = styleSettings.shadow || {}
  const rawIntensity = typeof layer?.shadowIntensity === 'string'
    ? layer.shadowIntensity
    : (typeof layer?.shadow_intensity === 'string' ? layer.shadow_intensity : shadow.intensity)
  const intensity = rawIntensity && rawIntensity.trim() !== '' ? rawIntensity : 'md'
  const rawDirection = typeof layer?.shadowDirection === 'string'
    ? layer.shadowDirection
    : (typeof layer?.shadow_direction === 'string' ? layer.shadow_direction : shadow.direction)
  const direction = rawDirection && rawDirection.trim() !== '' ? rawDirection : 'top'
  const enabled = resolveToggle(
    layer?.elevationEnabled ??
    layer?.elevation_enabled ??
    layer?.enableShadow ??
    layer?.enable_shadow ??
    shadow.enabled
  ) && intensity !== 'none'

  return {
    enabled,
    intensity,
    direction
  }
}

function resolveShadowStyles(layer: any) {
  const shadowConfig = resolveShadowConfig(layer)
  if (!shadowConfig.enabled) return { boxShadow: '', position: '', zIndex: undefined as number | undefined }

  const preset = SHADOW_PRESETS[shadowConfig.intensity]
  const boxShadow = preset?.[shadowConfig.direction] || preset?.top || '0 24px 48px -32px rgba(15, 23, 42, 0.35)'

  return {
    boxShadow,
    position: 'relative',
    zIndex: 5
  }
}

function getLayerItems(layer) {
  if (!layer || !layer.items) return []
  try {
    if (typeof layer.items === 'string') {
      if (layer.items.trim() === '') return []
      const parsed = JSON.parse(layer.items)
      return Array.isArray(parsed) ? parsed : []
    }
    return Array.isArray(layer.items) ? layer.items : []
  } catch {
    return []
  }
}

/**
 * تعیین نام کامپوننت برای رندر ویجت هدر بر اساس فیلد type یا component
 */
function resolveWidgetComponent(item: any): string | null {
  let raw: string | undefined
  if (typeof item === 'string') {
    raw = item
  } else {
    raw = item?.component || item?.type
  }

  if (!raw && typeof item === 'object' && item.id) raw = item.id

  if (!raw) return null

  const pascal = toPascalCase(raw)
  const key = `HeaderWidget${pascal}`

  return widgetRegistry[key] ?? null
}

function toPascalCase(str: string): string {
  return str
    .replace(/[^a-zA-Z0-9]+/g, ' ') // جداکننده‌ها را فاصله می‌کنیم
    .split(' ')
    .map(segment => segment.charAt(0).toUpperCase() + segment.slice(1))
    .join('')
}

// تابع تعیین justify-content بر اساس چینش
function getJustifyContent(align: string): string {
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

function getItemStyle(item: any) {
  if (typeof item !== 'object') return {}
  const style: Record<string, any> = {
    display: 'flex',
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: getJustifyContent(item.align || 'center'),
    height: '100%'
  }
  
  // تنظیم ارتفاع برای آیتم عکس
  if (item.type === 'image' || item.component === 'image') {
    style.alignItems = 'stretch'
  }
  
  // تنظیم عرض آیتم
  if (item.width) {
    style.flex = `0 0 ${item.width}%`
    style.width = `${item.width}%`
  } else {
    style.flex = '1 1 auto'
  }
  
  // تنظیم پس‌زمینه
  if (item.bgColor && item.bgColor !== 'transparent') {
    style.backgroundColor = item.bgColor
  }
  
  // تنظیم پدینگ راست و چپ
  if (item.paddingRight !== undefined) {
    style.paddingRight = `${item.paddingRight}px`
  }
  if (item.paddingLeft !== undefined) {
    style.paddingLeft = `${item.paddingLeft}px`
  }
  
  // برای متن، align اعمال شود
  if (item.align) {
    style.textAlign = item.align === 'left' ? 'left' : item.align === 'right' ? 'right' : 'center'
  }
  
  return style
}

/**
 * تعیین استایل جداکننده لایه
 */
function getSeparatorStyle(layer: any) {
  const style: Record<string, any> = {}
  
  // تنظیم نوع خط
  if (layer.separatorType) {
    switch (layer.separatorType) {
      case 'dashed':
        style.borderLeftStyle = 'dashed'
        break
      case 'dotted':
        style.borderLeftStyle = 'dotted'
        break
      case 'double':
        style.borderLeftStyle = 'double'
        break
      default:
        style.borderLeftStyle = 'solid'
    }
  } else {
    style.borderLeftStyle = 'solid'
  }
  
 
  
  return style
}

// @ts-ignore - vite/nuxt import meta glob
const widgetModules = import.meta.glob('./HeaderWidget*.vue', { eager: true })
const widgetRegistry: Record<string, any> = {}
for (const path in widgetModules) {
  const mod: any = (widgetModules as any)[path]
  if (mod && mod.default) {
    const name = path.split('/').pop()?.replace('.vue', '') || ''
    widgetRegistry[name] = mod.default
  }
}


onMounted(async () => {
  try {
    await loadHeaders()
    activeHeader.value = getActiveHeader()
    // اضافه کردن watch برای تغییرات route
    watch(() => route.path, async (newPath) => {
      // بررسی مجدد اینکه آیا هدر باید نمایش داده شود
      if (activeHeader.value) {
        try {
          const headerForPage = getHeaderForPage(newPath)
          // اگر هدر برای صفحه جدید وجود نداشته باشد، هدر را پاک کن
          if (!headerForPage) {
            activeHeader.value = null
          } else {
            // اگر هدر برای صفحه جدید وجود داشته باشد، آن را تنظیم کن
            activeHeader.value = headerForPage
          }
        } catch (error) {
          // در صورت خطا، هدر را پاک کن
          activeHeader.value = null
        }
      }
    })
  } catch (error) {
    activeHeader.value = null
  }
})
</script>

<style scoped>

.header-item-wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
  box-sizing: border-box;
  overflow: hidden;
}

.header-item-wrapper :deep(img) {
  max-height: 100%;
  height: 100%;
  width: auto;
  object-fit: contain;
}

/* هدر همیشه بالاتر از آیتم‌های صفحه باشد */

/* استایل‌های responsive */
@media (max-width: 768px) {

}
</style>
