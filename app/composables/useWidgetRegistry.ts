/**
 * Dynamic Widget Registry for SSR-safe widget loading
 * No hardcoded imports, fully dynamic component resolution
 */

import { readonly, ref } from 'vue'

export interface WidgetDefinition {
  type: string
  name: string
  description: string
  category: string
  icon: string
  componentPath: string
  defaultConfig: Record<string, unknown>
  configSchema?: Record<string, unknown>
}

// Widget registry - can be extended dynamically
const widgetRegistry = ref<Record<string, WidgetDefinition>>({
  'main-slider-side-banner': {
    type: 'main-slider-side-banner',
    name: 'اسلایدر اصلی + بنر کناری',
    description: 'اسلایدر با بنرهای کناری',
    category: 'slider',
    icon: 'fas fa-images',
    componentPath: 'components/widgets/MainSliderSideBanner.vue',
    defaultConfig: {
      slider_count: 5,
      wide_bg: true,
      bg_color: '#ffffff',
      banner_position: 'right',
      slides: [],
      side_banners: []
    }
  },
  'single-slider-side': {
    type: 'single-slider-side',
    name: 'اسلایدر تکی کناری',
    description: 'اسلایدر با یک بنر کناری',
    category: 'slider',
    icon: 'fas fa-image',
    componentPath: 'components/widgets/SingleSliderSide.vue',
    defaultConfig: {
      wide_bg: true,
      bg_color: '#ffffff',
      slides: [],
      side_banner: null
    }
  },
  'full-slider': {
    type: 'full-slider',
    name: 'اسلایدر تمام عرض',
    description: 'اسلایدر در تمام عرض صفحه',
    category: 'slider',
    icon: 'fas fa-image',
    componentPath: 'components/widgets/FullSlider.vue',
    defaultConfig: {
      slider_count: 5,
      wide_bg: true,
      bg_color: '#ffffff',
      slides: []
    }
  },
  'product-carousel': {
    type: 'product-carousel',
    name: 'کاروسل محصولات',
    description: 'نمایش محصولات در قالب کاروسل',
    category: 'product',
    icon: 'fas fa-shopping-cart',
    componentPath: 'components/widgets/ProductCarousel.vue',
    defaultConfig: {
      title: 'محصولات پیشنهادی',
      subtitle: 'بهترین محصولات برای شما',
      categoryId: null,
      productCount: 12,
      slidesPerView: 4,
      showPrice: true,
      showRating: true,
      showDiscount: true,
      autoplay: true,
      autoplayDelay: 3000,
      wide_bg: false,
      bg_color: '#ffffff',
      background_color: '#ffffff',
      padding: '24px',
      margin: '0px',
      showNavigation: true,
      showIndicators: true,
      showControls: false,
      navigationStyle: 'default',
      navigationSize: 40,
      indicatorStyle: 'default',
      indicatorColor: '#3b82f6'
    }
  },
  'double-banner': {
    type: 'double-banner',
    name: 'بنر دوتایی',
    description: 'دو بنر در کنار هم',
    category: 'banner',
    icon: 'fas fa-th-large',
    componentPath: '~/components/widgets/DoubleBanner.vue',
    defaultConfig: {
      wide_bg: false,
      bg_color: '#ffffff',
      banners: []
    }
  },
  'triple-banner': {
    type: 'triple-banner',
    name: 'بنر سه‌تایی',
    description: 'سه بنر در کنار هم',
    category: 'banner',
    icon: 'fas fa-th',
    componentPath: '~/components/widgets/TripleBanner.vue',
    defaultConfig: {
      wide_bg: false,
      bg_color: '#ffffff',
      banners: []
    }
  },
  'quad-banner': {
    type: 'quad-banner',
    name: 'بنر چهارتایی',
    description: 'چهار بنر در کنار هم',
    category: 'banner',
    icon: 'fas fa-th',
    componentPath: '~/components/widgets/QuadBanner.vue',
    defaultConfig: {
      wide_bg: false,
      bg_color: '#ffffff',
      banners: []
    }
  },
  'full-banner': {
    type: 'full-banner',
    name: 'بنر تکی',
    description: 'بنر تکی تمام عرض',
    category: 'banner',
    icon: 'fas fa-image',
    componentPath: '~/components/widgets/FullBanner.vue',
    defaultConfig: {
      height: 400,
      show_title: true,
      show_description: true,
      bg_enabled: false,
      bg_color: '#ffffff',
      slides: []
    }
  },

  'category-box': {
    type: 'category-box',
    name: 'باکس دسته‌بندی',
    description: 'نمایش دسته‌بندی‌ها',
    category: 'category',
    icon: 'fas fa-tags',
    componentPath: '~/components/widgets/CategoryBox.vue',
    defaultConfig: {
      categories: [],
      layout: 'grid',
      vertical_layout: false,
      show_product_count: true,
      columns: 3,
      show_title: true,
      title: 'دسته‌بندی‌های محبوب',
      title_color: '#1f2937',
      card_style: 'default',
      image_size: 'medium',
      border_radius: 'medium',
      text_alignment: 'center',
      background_color: '#ffffff',
      show_background: false,
      full_width_background: false,
      show_border: false,
      border_color: '#e5e7eb',
      border_width: 'medium',
      category_source: 'manual',
      padding: '24px',
      margin: '0px'
    }
  },

  'mobile-header': {
    type: 'mobile-header',
    name: 'هدر موبایل',
    description: 'نمایش هدرهای موبایل و اپلیکیشن',
    category: 'layout',
    icon: 'fas fa-mobile-alt',
    componentPath: '~/components/widgets/MobileHeader.vue',
    defaultConfig: {
      headers: [],
      layout: 'grid',
      columns: 3,
      mobile_columns: 1,
      show_title: true,
      title: 'هدرهای موبایل و اپلیکیشن',
      title_color: '#1f2937',
      card_style: 'default',
      background_color: '#ffffff',
      show_background: false,
      show_border: false,
      border_color: '#e5e7eb',
      border_width: 'medium',
      padding: '24px',
      margin: '0px',
      box_width: 'center',
      center_width: 1200,
      custom_class: ''
    }
  }
})

export const useWidgetRegistry = () => {
  /**
   * Register a new widget type
   */
  const registerWidget = (definition: WidgetDefinition) => {
    widgetRegistry.value[definition.type] = definition
  }

  /**
   * Get widget definition by type
   */
  const getWidgetDefinition = (type: string): WidgetDefinition | null => {
    return widgetRegistry.value[type] || null
  }

  /**
   * Get all registered widgets
   */
  const getAllWidgets = (): WidgetDefinition[] => {
    return Object.values(widgetRegistry.value)
  }

  /**
   * Get widgets by category
   */
  const getWidgetsByCategory = (category: string): WidgetDefinition[] => {
    return Object.values(widgetRegistry.value).filter(w => w.category === category)
  }

  /**
   * Get widget categories
   */
  const getCategories = (): string[] => {
    const categories = new Set(Object.values(widgetRegistry.value).map(w => w.category))
    return Array.from(categories)
  }

  /**
   * Dynamically load widget component
   */
  const loadWidgetComponent = async (type: string) => {
    const definition = getWidgetDefinition(type)
    if (!definition) {
      throw new Error(`Widget type "${type}" not found in registry`)
    }

    try {
      // Dynamic import based on component path
      const component = await import(/* @vite-ignore */ definition.componentPath)
      return component.default || component
    } catch (error) {
      // Failed to load widget component
      throw error
    }
  }

  /**
   * Create widget with default config
   */
  const createWidgetConfig = (type: string, overrides: Record<string, unknown> = {}) => {
    const definition = getWidgetDefinition(type)
    if (!definition) {
      throw new Error(`Widget type "${type}" not found`)
    }

    return {
      ...definition.defaultConfig,
      ...overrides
    }
  }

  /**
   * Validate widget config against schema
   */
  const validateWidgetConfig = (type: string, config: Record<string, unknown>): boolean => {
    const definition = getWidgetDefinition(type)
    if (!definition?.configSchema) {
      return true // No schema = valid
    }

    // TODO: Implement JSON schema validation
    return true
  }

  return {
    registerWidget,
    getWidgetDefinition,
    getAllWidgets,
    getWidgetsByCategory,
    getCategories,
    loadWidgetComponent,
    createWidgetConfig,
    validateWidgetConfig,
    widgetRegistry: readonly(widgetRegistry)
  }
}