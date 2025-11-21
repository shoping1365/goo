<template>
  <div class="category-box-preview w-full" :class="containerClasses">
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
        :class="gridClasses"
        :style="borderStyle"
      >
        <div 
          v-for="category in displayCategories" 
          :key="category.id"
          class="w-full min-w-0 flex-shrink-0"
        >
          <!-- تصویر دسته‌بندی -->
          <div :class="imageContainerClasses">
            <img 
              v-if="category.image"
              :src="category.image" 
              :alt="category.title"
              :class="imageClasses"
            />
            <div v-else class="w-full h-full flex items-center justify-center bg-gray-100 text-gray-400">
              <i class="fas fa-image text-2xl"></i>
            </div>
          </div>
        </div>
      </div>

      <!-- پیام خالی -->
      <div v-if="displayCategories.length === 0" class="text-center py-8 text-gray-500">
        <i class="fas fa-tags text-4xl mb-4"></i>
        <p>هنوز دسته‌بندی‌ای اضافه نشده است</p>
        <p class="text-sm">برای مشاهده پیش‌نمایش، دسته‌بندی‌ها را اضافه کنید</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import type { CategoryConfig } from '~/types/widget';

interface Props {
  config: CategoryConfig
}

const props = defineProps<Props>()

// تنظیمات پیش‌فرض
const config = computed(() => ({
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
  margin: '0px',
  ...props.config
}))

// استفاده از داده‌های واقعی کاربر
const displayCategories = computed(() => {
  return config.value.categories || []
})

// کلاس‌های داینامیک برای container
const containerClasses = computed(() => {
  const classes = ['w-full']
  
  // تنظیمات اندازه باکس
  if (props.config.box_width === 'center') {
    classes.push('max-w-7xl mx-auto px-4 sm:px-6 lg:px-8')
  } else {
    classes.push('px-4 sm:px-6 lg:px-8')
  }
  
  if (props.config.background_color) {
    classes.push(`bg-[${props.config.background_color}]`)
  }
  
  if (props.config.custom_class) {
    classes.push(props.config.custom_class)
  }
  
  return classes.join(' ')
})

// کلاس‌های داینامیک برای grid
const gridClasses = computed(() => {
  const baseClasses = 'grid gap-6 w-full'
  
  // اضافه کردن کلاس‌های خط دور کل مجموعه اگر فعال باشد
  let borderClasses = ''
  if (props.config.show_border) {
    const borderWidthClasses = {
      thin: 'border',
      medium: 'border-2',
      thick: 'border-4'
    }
    const borderWidth = borderWidthClasses[props.config.border_width as keyof typeof borderWidthClasses] || 'border-2'
    borderClasses = `${borderWidth} rounded-lg p-6`
  }
  
  // اگر چیدمان عمودی فعال باشد یا layout لیست باشد
  if (props.config.vertical_layout || props.config.layout === 'list') {
    return `${baseClasses} grid-cols-1 ${borderClasses}`
  }
  
  // کلاس‌های responsive برای موبایل و دسکتاپ
  const mobileColumns = props.config.mobile_columns || 2
  const desktopColumns = props.config.columns || 3
  
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
  if (props.config.show_border && props.config.border_color) {
    return {
      borderColor: props.config.border_color
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
  
  const baseClasses = `${sizeClasses[props.config.image_size as keyof typeof sizeClasses] || 'w-20 h-20 sm:w-32 sm:h-32'} mx-auto mb-3`
  const radiusClass = radiusClasses[props.config.border_radius as keyof typeof radiusClasses] || 'rounded-xl'
  
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
  
  const radiusClass = radiusClasses[props.config.border_radius as keyof typeof radiusClasses] || 'rounded-xl'
  
  return `${baseClasses} ${radiusClass}`
})

</script>

