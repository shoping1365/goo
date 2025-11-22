<template>
  <div class="widget-container" :class="containerClasses">
    <Suspense>
      <component
        :is="dynamicComponent"
        v-bind="widgetProps"
        :key="`${widget.id}-${widget.updated_at}`"
        :widget="widget"
        :widget-type="widget.type"
      />
      <template #fallback>
        <div class="widget-loading p-6">
          <WidgetLoading />
        </div>
      </template>
    </Suspense>
  </div>
</template>

<script setup lang="ts">
import { computed, defineAsyncComponent, h } from 'vue'
import type { Widget } from '~/types/widget'
import WidgetNotFound from '~/components/common/WidgetNotFound.vue'
import WidgetLoading from '~/components/common/WidgetLoading.vue'
import WidgetError from '~/components/common/WidgetError.vue'

// Props
interface Props {
  widget: Widget
}

const props = defineProps<Props>()

// Dynamic component loading with async/await
const dynamicComponent = computed(() => {
  interface ComponentModule {
    default: unknown
  }
  
  const componentMap: Record<string, () => Promise<ComponentModule>> = {
    'main-slider-side-banner': () => import('~/components/widgets/MainSliderSideBanner.vue'),
    'single-slider-side': () => import('~/components/widgets/SingleSliderSide.vue'),

    'full-banner': () => import('~/components/widgets/FullBanner.vue'),

    'double-banner': () => import('~/components/widgets/DoubleBanner.vue'),
    'triple-banner': () => import('~/components/widgets/TripleBanner.vue'),
    'quad-banner': () => import('~/components/widgets/QuadBanner.vue'),
    'penta-banner': () => import('~/components/widgets/PentaBanner.vue'),
    'brands-slider': () => import('~/components/widgets/BrandsSlider.vue'),
    'services-slider': () => import('~/components/widgets/ServicesSlider.vue'),
    'category-box': () => import('~/components/widgets/CategoryBox.vue'),
    'blog-posts': () => import('~/components/widgets/BlogPosts.vue'),
    'product-carousel': () => import('~/components/widgets/ProductCarousel.vue'),
    'mobile-header': () => import('~/components/widgets/MobileHeaderPreview.vue')
  }
  
  const componentLoader = componentMap[props.widget.type]
  
  if (!componentLoader) {
    return WidgetNotFound
  }
  
  // Create error component wrapper that receives widget-type prop
  const ErrorComponentWrapper = {
    props: {
      widgetType: {
        type: String,
        default: 'unknown'
      }
    },
    setup(props: { widgetType: string }) {
      return () => h(WidgetError, { widgetType: props.widgetType })
    }
  }
  
  return defineAsyncComponent({
    loader: componentLoader,
    loadingComponent: WidgetLoading,
    errorComponent: ErrorComponentWrapper,
    delay: 0,
    timeout: 0 // حذف کامل timeout
  })
})

// Dynamic container classes based on widget config
const containerClasses = computed(() => {
  interface WidgetConfig {
    bg_enabled?: boolean
    bg_color?: string
    [key: string]: unknown
  }
  
  const config = props.widget.config as WidgetConfig | null | undefined
  const classes = ['widget-container']
  
  // Add background color only if bg_enabled is true
  if (config?.bg_enabled === true && config?.bg_color && config.bg_color !== '#ffffff') {
    classes.push('widget-custom-bg')
  }
  
  return classes
})

// Methods
const _refreshWidget = () => {
  // Force component re-render by changing the key
  // Component will re-render automatically due to key change
}

// Dynamic props based on widget configuration
const widgetProps = computed(() => {
  const baseProps = {
    id: props.widget.id,
    code: props.widget.code,
    title: props.widget.title,
    description: props.widget.description,
    config: props.widget.config,
    order: props.widget.order,
    status: props.widget.status
  }
  
  // Parse config if it's a string
  if (typeof props.widget.config === 'string') {
    try {
      baseProps.config = JSON.parse(props.widget.config)
    } catch (_e) {
      // Failed to parse widget config
    }
  }
  
  return baseProps
})

// Dynamic background color for v-bind
const widgetBgColor = computed(() => {
  interface WidgetConfig {
    bg_enabled?: boolean
    bg_color?: string
    [key: string]: unknown
  }
  
  const config = widgetProps.value.config as WidgetConfig | null | undefined
  // Only return background color if bg_enabled is true
  if (config?.bg_enabled === true && config?.bg_color) {
    return config.bg_color
  }
  return 'transparent'
})
</script>

<style scoped>
.widget-container {
  width: 100%;
  max-width: 100vw;
  margin: 0;
  padding: 0;
  overflow-x: hidden;
  box-sizing: border-box;
  display: flex;
  justify-content: center;
}

.widget-loading {
  padding: 1rem;
  width: 100%;
  overflow-x: hidden;
}

.widget-custom-bg {
  background-color: v-bind('widgetBgColor');
  width: 100%;
  overflow-x: hidden;
}
</style> 