import { computed, type Ref } from 'vue'

export interface MobileImageConfig {
     banners: Array<{
          image: string
          mobile_image?: string
          title?: string
          description?: string
     }>
     mobile_image_mode?: 'auto' | 'separate'
     mobile_cropped_image?: string
     mobile_crop_width?: number
     mobile_crop_height?: number
     mobile_height?: number
     banner_width?: number
     center_width?: number
     height?: number
     bg_enabled?: boolean
     bg_color?: string
     padding_top?: number
     padding_bottom?: number
     margin_left?: number
     margin_right?: number
}

export const useMobileImageHandler = (config: Ref<MobileImageConfig>) => {
     // تشخیص نوع دستگاه و دریافت URL مناسب
     const getDeviceSpecificImageUrl = () => {
          if (!config.value.banners || !config.value.banners.length) {
               return ''
          }

          const banner = config.value.banners[0]

          // اگر حالت عکس جداگانه انتخاب شده و عکس موبایل وجود دارد
          if (config.value.mobile_image_mode === 'separate' && banner.mobile_image) {
               return banner.mobile_image
          }

          // اگر برش خودکار اعمال شده (اما URL قدیمی نباشد)
          if (config.value.mobile_cropped_image && !config.value.mobile_cropped_image.includes('auto=true')) {
               return config.value.mobile_cropped_image
          }

          // برای موبایل: از عکس اصلی استفاده کن
          if (typeof window !== 'undefined' && window.innerWidth < 768) {
               return banner.image || ''
          }

          // برای دسکتاپ: عکس اصلی را برگردان
          return banner.image || ''
     }

     // دریافت متن alt برای تصویر
     const getImageAltText = () => {
          if (!config.value.banners || !config.value.banners.length) {
               return ''
          }

          const banner = config.value.banners[0]

          // اگر عنوان وجود دارد، از آن استفاده کن
          if (banner.title && banner.title.trim()) {
               return banner.title.trim()
          }

          // اگر توضیحات وجود دارد، از آن استفاده کن
          if (banner.description && banner.description.trim()) {
               return banner.description.trim()
          }

          // اگر هیچ کدام نبود، alt خالی باشد
          return ''
     }

     // محاسبه ارتفاع responsive
     const getResponsiveHeight = () => {
          if (typeof window === 'undefined') {
               return config.value.height || '200px'
          }

          if (window.innerWidth < 768) {
               return config.value.mobile_height ? `${config.value.mobile_height}px` : '150px'
          }

          return config.value.height ? `${config.value.height}px` : '200px'
     }

     // استایل container با تنظیمات responsive
     const containerStyle = computed(() => {
          const styles: Record<string, string | number> = {
               height: getResponsiveHeight(),
               '--desktop-height': config.value.height ? `${config.value.height}px` : '200px',
               '--mobile-height': config.value.mobile_height ? `${config.value.mobile_height}px` : '150px'
          }

          // تنظیم عرض و موقعیت
          if (config.value.banner_width === 800) {
               // تمام عرض
               styles.width = '100%'
               styles.maxWidth = '100%'
          } else if (config.value.banner_width === 600) {
               // وسط با عرض مشخص
               const centerWidth = config.value.center_width || 1000
               styles.width = `${centerWidth}px`
               styles.maxWidth = `${centerWidth}px`
               styles.margin = '0 auto'
          } else {
               // پیش‌فرض
               styles.width = '100%'
               styles.maxWidth = '100%'
          }

          // اگر پس‌زمینه فعال است، رنگ آن را اضافه کن
          if (config.value.bg_enabled && config.value.bg_color) {
               styles.backgroundColor = config.value.bg_color
          }

          // پدینگ و مارجین
          if (config.value.padding_top !== undefined) {
               styles.paddingTop = `${config.value.padding_top}px`
          }
          if (config.value.padding_bottom !== undefined) {
               styles.paddingBottom = `${config.value.padding_bottom}px`
          }

          // مارجین چپ و راست فقط در حالت تمام عرض اعمال شود
          if (config.value.banner_width === 800) {
               if (config.value.margin_left !== undefined) {
                    styles.marginLeft = `${config.value.margin_left}px`
               }
               if (config.value.margin_right !== undefined) {
                    styles.marginRight = `${config.value.margin_right}px`
               }
          }

          return styles
     })

     return {
          getDeviceSpecificImageUrl,
          getImageAltText,
          getResponsiveHeight,
          containerStyle
     }
}
