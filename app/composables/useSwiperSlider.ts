import type { Swiper } from 'swiper/types'
import { onUnmounted, ref } from 'vue'

export function useSwiperSlider() {
     const swiperInstance = ref<Swiper | null>(null)
     const isInitialized = ref(false)

     const onSwiper = (swiper: Swiper) => {
          swiperInstance.value = swiper
          isInitialized.value = true

          // تنظیمات پایه Swiper
          if (swiper) {
               // بروزرسانی ساده Swiper
               setTimeout(() => {
                    swiper.update()
               }, 100)
          }
     }

     const onSlideChange = () => {
          if (swiperInstance.value) {
               // Slide changed
          }
     }

     const goToSlide = (index: number) => {
          if (swiperInstance.value) {
               swiperInstance.value.slideTo(index)
          }
     }

     const nextSlide = () => {
          if (swiperInstance.value) {
               swiperInstance.value.slideNext()
          }
     }

     const prevSlide = () => {
          if (swiperInstance.value) {
               swiperInstance.value.slidePrev()
          }
     }

     const startAutoplay = () => {
          if (swiperInstance.value?.autoplay) {
               swiperInstance.value.autoplay.start()
          }
     }

     const stopAutoplay = () => {
          if (swiperInstance.value?.autoplay) {
               swiperInstance.value.autoplay.stop()
          }
     }

     const destroySwiper = () => {
          if (swiperInstance.value) {
               swiperInstance.value.destroy(true, true)
               swiperInstance.value = null
               isInitialized.value = false
          }
     }

     onUnmounted(() => {
          destroySwiper()
     })

     return {
          swiperInstance,
          isInitialized,
          onSwiper,
          onSlideChange,
          goToSlide,
          nextSlide,
          prevSlide,
          startAutoplay,
          stopAutoplay,
          destroySwiper
     }
}
