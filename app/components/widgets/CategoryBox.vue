<template>
  <div class="category-box w-full" :class="containerClasses">
    <div class="py-6">
      <!-- عنوان بخش -->
      <div v-if="config.show_title && config.title" class="mb-6">
        <h2 
          class="text-xl font-bold text-right"
          :style="{ color: config.title_color || '#1f2937' }"
        >
          {{ config.title }}
        </h2>
      </div>
      
      <!-- تصاویر دسته‌بندی -->
      <div 
        v-if="displayCategories && displayCategories.length > 0" 
        :class="gridClasses"
        :style="borderStyle"
      >
        <div 
          v-for="(category, index) in displayCategories" 
          :key="category.id"
          class="w-full min-w-0 flex-shrink-0"
          @click="navigateToCategory(category.link)"
        >
          <!-- تصویر دسته‌بندی -->
          <div :class="imageContainerClasses">
            <img 
              :src="category.image || '/default-product.svg'" 
              :alt="category.title"
              :class="imageClasses"
            />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import type { CategoryConfig, Widget } from '~/types/widget'

const router = useRouter()

interface Props {
  widget: Widget
}

const props = defineProps<Props>()

// استخراج تنظیمات از widget config
const config = computed<CategoryConfig>(() => ({
  categories: [],
  layout: 'grid',
  vertical_layout: false,
  show_product_count: true,
  columns: 3,
  mobile_columns: 2,
  box_width: 'center',
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
  margin: '0px',
  ...props.widget.config
} as CategoryConfig))

// استفاده از داده‌های واقعی کاربر
const displayCategories = computed(() => {
  return config.value.categories || []
})

// کلاس‌های داینامیک برای container
const containerClasses = computed(() => {
  const classes = ['w-full']
  
  // تنظیمات اندازه باکس
  if (config.value.box_width === 'center') {
    classes.push('max-w-7xl mx-auto px-4 sm:px-6 lg:px-8')
  } else {
    classes.push('px-4 sm:px-6 lg:px-8')
  }
  
  if (config.value.background_color) {
    classes.push(`bg-[${config.value.background_color}]`)
  }
  
  if (config.value.custom_class) {
    classes.push(config.value.custom_class)
  }
  
  return classes.join(' ')
})

// کلاس‌های داینامیک برای grid container
const gridContainerClasses = computed(() => {
  return 'grid gap-6'
})

// کلاس‌های داینامیک برای grid
const gridClasses = computed(() => {
  const baseClasses = 'grid gap-6 w-full'
  
  // اضافه کردن کلاس‌های خط دور کل مجموعه اگر فعال باشد
  let borderClasses = ''
  if (config.value.show_border) {
    const borderWidthClasses = {
      thin: 'border',
      medium: 'border-2',
      thick: 'border-4'
    }
    const borderWidth = borderWidthClasses[config.value.border_width as keyof typeof borderWidthClasses] || 'border-2'
    borderClasses = `${borderWidth} rounded-lg p-6`
  }
  
  // اگر چیدمان عمودی فعال باشد یا layout لیست باشد
  if (config.value.vertical_layout || config.value.layout === 'list') {
    return `${baseClasses} grid-cols-1 ${borderClasses}`
  }
  
  // کلاس‌های responsive برای موبایل و دسکتاپ
  const mobileColumns = config.value.mobile_columns || 2
  const desktopColumns = config.value.columns || 3
  
  // استفاده از کلاس‌های ثابت Tailwind
  const mobileClasses = {
    1: 'grid-cols-1',
    2: 'grid-cols-2', 
    3: 'grid-cols-3',
    4: 'grid-cols-4'
  }
  
  const desktopClasses = {
    1: 'sm:grid-cols-1',
    2: 'sm:grid-cols-2',
    3: 'sm:grid-cols-3', 
    4: 'sm:grid-cols-4',
    5: 'sm:grid-cols-5',
    6: 'sm:grid-cols-6'
  }
  
  const mobileClass = mobileClasses[mobileColumns as keyof typeof mobileClasses] || 'grid-cols-2'
  const desktopClass = desktopClasses[desktopColumns as keyof typeof desktopClasses] || 'sm:grid-cols-3'
  
  return `${baseClasses} ${mobileClass} ${desktopClass} ${borderClasses}`
})

// استایل داینامیک برای خط دور
const borderStyle = computed(() => {
  if (config.value.show_border && config.value.border_color) {
    return {
      borderColor: config.value.border_color
    }
  }
  return {}
})


// کلاس‌های داینامیک برای container تصویر
const imageContainerClasses = computed(() => {
  const sizeClasses = {
    small: 'w-16 h-16 sm:w-24 sm:h-24',    // موبایل: 64px، دسکتاپ: 96px
    medium: 'w-20 h-20 sm:w-32 sm:h-32',   // موبایل: 80px، دسکتاپ: 128px  
    large: 'w-24 h-24 sm:w-40 sm:h-40'     // موبایل: 96px، دسکتاپ: 160px
  }
  
  const radiusClasses = {
    none: '',
    small: 'rounded-md',
    medium: 'rounded-xl',
    large: 'rounded-2xl'
  }
  
  const baseClasses = `${sizeClasses[config.value.image_size as keyof typeof sizeClasses] || 'w-20 h-20 sm:w-32 sm:h-32'} mx-auto mb-3`
  const radiusClass = radiusClasses[config.value.border_radius as keyof typeof radiusClasses] || 'rounded-xl'
  
  return `${baseClasses} ${radiusClass} overflow-hidden`
})

// کلاس‌های داینامیک برای تصویر
const imageClasses = computed(() => {
  const baseClasses = 'w-full h-full object-cover'
  
  const radiusClasses = {
    none: '',
    small: 'rounded-lg',
    medium: 'rounded-xl',
    large: 'rounded-2xl'
  }
  
  const radiusClass = radiusClasses[config.value.border_radius as keyof typeof radiusClasses] || 'rounded-xl'
  
  return `${baseClasses} ${radiusClass}`
})


// تابع ناوبری به دسته‌بندی
const navigateToCategory = (link: string) => {
  if (link) {
    router.push(link)
  }
}
</script> 