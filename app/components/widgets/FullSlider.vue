<template>
  <div
    class="w-screen left-1/2 -translate-x-1/2 relative py-6"
    style="position: relative; margin-left: calc(50% - 50vw); margin-right: calc(50% - 50vw); width: 100vw; max-width: 100vw;"
  >
    <div v-if="hasSlides" class="slider-container">
      <Swiper
        :modules="swiperModules"
        :slides-per-view="1"
        :space-between="0"
        :loop="canLoop"
        :autoplay="canLoop ? {
          delay: 5000,
          disableOnInteraction: false,
        } : false"
        :navigation="canLoop"
        :pagination="canLoop ? { clickable: true } : false"
        :effect="'fade'"
        :fade-effect="{ crossFade: true }"
        :class="`w-full rounded-lg overflow-hidden`"
        :style="`height: ${sliderHeight}px`"
      >
        <SwiperSlide
          v-for="(slideItem, index) in slidesList"
          :key="index"
          class="relative"
        >
          <a 
            v-if="getSlideItem(slideItem).link"
            :href="getSlideItem(slideItem).link"
            class="block h-full"
          >
            <img 
              :src="getSlideItem(slideItem).image" 
              :alt="getSlideItem(slideItem).title"
              class="w-full h-full object-scale-down"
            />
            <div 
              v-if="getSlideItem(slideItem).title" 
              class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-6"
            >
              <h3 class="text-white text-lg md:text-xl lg:text-2xl font-bold">{{ getSlideItem(slideItem).title }}</h3>
              <p v-if="getSlideItem(slideItem).description" class="text-white/90 text-sm md:text-base mt-1">{{ getSlideItem(slideItem).description }}</p>
            </div>
          </a>
          <div v-else class="block h-full">
            <img 
              :src="getSlideItem(slideItem).image" 
              :alt="getSlideItem(slideItem).title"
              class="w-full h-full object-scale-down"
            />
            <div 
              v-if="getSlideItem(slideItem).title" 
              class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-6"
            >
              <h3 class="text-white text-lg md:text-xl lg:text-2xl font-bold">{{ getSlideItem(slideItem).title }}</h3>
              <p v-if="getSlideItem(slideItem).description" class="text-white/90 text-sm md:text-base mt-1">{{ getSlideItem(slideItem).description }}</p>
            </div>
          </div>
        </SwiperSlide>
      </Swiper>
    </div>
    <div v-else class="bg-gray-200 rounded-lg flex items-center justify-center" :style="`height: ${sliderHeight}px`">
      <p class="text-gray-500">هیچ اسلایدری یافت نشد</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import type { SlideItem, SliderConfig, Widget } from '~/types/widget'

interface Props {
  widget: Widget
}

const props = defineProps<Props>()

// Import Swiper modules
import { Autoplay, EffectFade, Navigation, Pagination } from 'swiper/modules';

const swiperModules = [
  Autoplay,
  Navigation, 
  Pagination,
  EffectFade
]

// Get config with proper typing
const config = computed<SliderConfig>(() => {
  const defaultConfig: SliderConfig = {
    slider_count: 5,
    wide_bg: true,
    bg_color: '#ffffff',
    banner_position: 'right',
    display_order: 'asc',
    height: 400,
    slides: [],
    side_banners: [],
    bg_enabled: false,
    padding_top: 0,
    padding_bottom: 0,
    margin_left: 0,
    margin_right: 0,
    autoplay_delay: 5000,
    autoplay_enabled: true,
    navigation_enabled: true
  }
  return { ...defaultConfig, ...props.widget.config } as SliderConfig
})

// Helper function to get slides array with proper typing
const getSlidesArray = (): SlideItem[] => {
  return (config.value.slides || []) as SlideItem[]
}

// Helper function to type slide item in template
const getSlideItem = (item: any): SlideItem => {
  return item as SlideItem
}

// Computed properties for template access
const slides = computed<SlideItem[]>(() => getSlidesArray())
const slidesList = computed<SlideItem[]>(() => getSlidesArray())
const sliderHeight = computed<number>(() => config.value.height || 400)
const hasSlides = computed<boolean>(() => getSlidesArray().length > 0)
const canLoop = computed<boolean>(() => getSlidesArray().length > 1)
</script>

<style scoped>
.full-slider {
  background: #f8fafc;
}

/* Custom Swiper styles */
:deep(.swiper-button-next),
:deep(.swiper-button-prev) {
  color: white;
  background: rgba(0, 0, 0, 0.5);
  border-radius: 9999px;
  width: 2.5rem;
  height: 2.5rem;
  transition: all 0.3s;
  margin-top: 0;
  transform: translateY(-50%);
}

:deep(.swiper-button-next:hover),
:deep(.swiper-button-prev:hover) {
  background: rgba(0, 0, 0, 0.7);
}

:deep(.swiper-button-next:after),
:deep(.swiper-button-prev:after) {
  font-size: 0.875rem;
}

:deep(.swiper-pagination-bullet) {
  background: rgba(255, 255, 255, 0.5);
  opacity: 1;
}

:deep(.swiper-pagination-bullet-active) {
  background: white;
}
</style> 