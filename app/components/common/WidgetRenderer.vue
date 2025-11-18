<template>
  <div class="widget-container" :class="containerClasses">
    <Suspense>
      <component
        :is="dynamicComponent"
        :widget="widget"
        v-bind="widgetProps"
        :key="`${widget.id}-${widget.updated_at}`"
      />
      <template #fallback>
        <div class="widget-loading p-6">
          <div class="animate-pulse bg-gray-200 rounded h-32"></div>
        </div>
      </template>
    </Suspense>
  </div>
</template>

<script setup lang="ts">
import { computed, defineComponent, defineAsyncComponent } from 'vue'
import type { Widget } from '~/types/widget'

// Props
interface Props {
  widget: Widget
}

const props = defineProps<Props>()

// Dynamic component loading with async/await
const dynamicComponent = computed(() => {
  const componentMap: Record<string, () => Promise<any>> = {
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
    return defineComponent({
      template: `<div class="text-red-500 p-6 border border-red-300 rounded">
        <h3 class="font-bold">Widget نادیده</h3>
        <p>نوع ویجت "${props.widget.type}" پیدا نشد.</p>
      </div>`
    })
  }
  
  return defineAsyncComponent({
    loader: componentLoader,
    loadingComponent: defineComponent({
      template: '<div class="animate-pulse bg-gray-200 rounded h-32"></div>'
    }),
    errorComponent: defineComponent({
      template: `<div class="text-red-500 p-6 border border-red-300 rounded">
        Failed to load widget: ${props.widget.type}
      </div>`
    }),
    delay: 0,
    timeout: 0 // حذف کامل timeout
  })
})

// Dynamic container classes based on widget config
const containerClasses = computed(() => {
  const config = props.widget.config as any
  const classes = ['widget-container']
  
  // Add background color only if bg_enabled is true
  if (config?.bg_enabled === true && config?.bg_color && config.bg_color !== '#ffffff') {
    classes.push('widget-custom-bg')
  }
  
  return classes
})

// Methods
const refreshWidget = () => {
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
    } catch (e) {
      // Failed to parse widget config
    }
  }
  
  return baseProps
})

// Dynamic background color for v-bind
const widgetBgColor = computed(() => {
  const config = widgetProps.value.config as any
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