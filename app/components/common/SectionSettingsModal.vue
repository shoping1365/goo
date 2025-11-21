<template>
  <Teleport to="body">
    <Transition name="fade">
      <div v-if="show" class="fixed inset-0 z-50 flex items-center justify-center bg-black/40 backdrop-blur-sm">
        <div class="bg-white border border-gray-200 shadow-2xl rounded-2xl w-full max-w-4xl mx-4 p-6 text-right transition-all duration-300" @click.stop>
          <h2 class="text-2xl font-extrabold mb-2 flex flex-row-reverse items-center justify-between w-full text-right">
            مدیریت نمایش بخش‌های صفحه محصول
            <button class="text-gray-400 hover:text-gray-600" @click="$emit('update:show', false)">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </h2>
          <div class="border-b border-gray-200 mb-4"></div>

          <!-- تب‌های اصلی -->
          <div class="mb-6">
            <div class="flex flex-row-reverse gap-2 overflow-x-auto pb-2">
              <button 
                v-for="tab in mainTabs" 
                :key="tab.value"
                :class="[
                  'px-4 py-2 rounded-lg text-sm font-medium transition-all duration-200 whitespace-nowrap',
                  activeTab === tab.value
                    ? 'bg-blue-600 text-white shadow-lg'
                    : 'bg-gray-100 text-gray-700 hover:bg-gray-200'
                ]"
                @click="activeTab = tab.value"
              >
                {{ tab.label }}
              </button>
            </div>
          </div>

          <!-- محتوای تب اطلاعات محصول -->
          <div v-if="activeTab === 'info'" class="max-h-[60vh] overflow-x-auto overflow-y-auto pr-1">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div v-for="section in infoSections" :key="section.value" class="bg-gray-50 rounded-xl p-6 shadow-sm hover:shadow-lg hover:bg-blue-50 transition-all duration-200 cursor-pointer">
                <label class="flex flex-row-reverse items-center gap-3 cursor-pointer">
                  <input 
                    v-model="sections[section.value]" 
                    type="checkbox" 
                    class="accent-blue-600 w-5 h-5" 
                  />
                  <div class="flex flex-col">
                    <span class="font-bold text-gray-800 text-base text-right">{{ section.label }}</span>
                    <span class="text-xs text-gray-500 mt-1">{{ section.description }}</span>
                  </div>
                  </label>
                </div>
            </div>
          </div>

          <!-- محتوای سایر تب‌ها -->
          <div v-else class="max-h-[60vh] overflow-x-auto overflow-y-auto pr-1">
            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6" dir="rtl">
              <div
v-for="section in getCurrentTabSections()" :key="section.value"
                class="bg-gray-50 rounded-xl p-6 text-right shadow-sm hover:shadow-lg hover:bg-blue-50 transition-all duration-200 cursor-pointer"
                dir="rtl"
              >
                <label class="flex flex-row-reverse items-center gap-3 cursor-pointer w-full">
                  <input 
                    v-model="sections[section.value]" 
                    type="checkbox" 
                    class="accent-blue-600 w-5 h-5" 
                  />
                  <div class="flex flex-col w-full">
                    <span
                      class="font-bold text-gray-800 text-base w-full block text-right"
                      dir="rtl"
                    >{{ section.label }}</span>
                    <span
                      class="text-xs text-gray-500 mt-1 w-full block text-right"
                      dir="rtl"
                    >{{ section.description }}</span>
                  </div>
                </label>
              </div>
            </div>
          </div>

          <div class="flex justify-start gap-3 mt-6 pt-4 border-t">
            <button class="px-4 py-2 rounded-lg text-sm border border-gray-300 bg-white hover:bg-gray-100 transition-colors shadow-sm font-semibold" @click="$emit('update:show', false)">
              انصراف
            </button>
            <button class="px-4 py-2 rounded-lg text-sm bg-blue-600 text-white hover:bg-blue-700 transition-colors shadow-md font-bold" @click="save">
              ذخیره تنظیمات
            </button>
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup>
import { toRefs, ref, onMounted } from 'vue'

const props = defineProps({
  show: {
    type: Boolean,
    required: true
  },
  sections: {
    type: Object,
    default: () => ({})
  },
  sectionLabels: {
    type: Object,
    default: () => ({})
  },
  sectionGroups: {
    type: Array,
    default: () => []
  }
})

const emit = defineEmits(['update:show', 'sections-updated'])

const { show } = toRefs(props)

// تب فعال
const activeTab = ref('info')

// تب‌های اصلی
const mainTabs = ref([
  { label: 'اطلاعات محصول', value: 'info' },
  { label: 'قیمت گذاری', value: 'pricing' },
  { label: 'مدیریت موجودی', value: 'inventory' },
  { label: 'حمل و نقل', value: 'shipping' },
  { label: 'مدیریت تصاویر', value: 'images' },
  { label: 'مدیریت تنوع کالایی', value: 'variants' },
  { label: 'مشخصات فنی محصول', value: 'specs' },
  { label: 'سئو', value: 'seo' },
  { label: 'محصولات مکمل', value: 'related' },
  { label: 'ویدیو', value: 'video' }
])

// بخش‌های تب اطلاعات محصول (فقط بخش‌های اختیاری)
const infoSections = ref([
  {
    value: 'displaySettings',
    label: 'تنظیمات نمایش و انتشار',
    description: 'کنترل نمایش محصول در بخش‌های مختلف سایت'
  },
  {
    value: 'scheduling',
    label: 'زمان‌بندی و تاریخ‌ها',
    description: 'برنامه‌ریزی انتشار و موجودی محصول'
  },
  {
    value: 'management',
    label: 'مدیریت و یادداشت‌ها',
    description: 'اطلاعات اداری و یادداشت‌های داخلی'
  },
  {
    value: 'strengthsWeaknesses',
    label: 'نقاط قوت و ضعف محصول',
    description: 'تحلیل جامع مزایا و معایب محصول برای مشتریان'
  }
])

// تعریف بخش‌های سایر تب‌ها
const otherTabSections = ref({
  pricing: [
    {
      value: 'priceSettings',
      label: 'تنظیمات قیمت و فروش',
      description: 'گزینه‌های فروش، زمان‌بندی قیمت و تنظیمات نمایش'
    },
    {
      value: 'taxSettings',
      label: 'مالیات و هزینه‌های اضافی',
      description: 'تنظیمات مالیات، هزینه بسته‌بندی و خدمات اضافی'
    },
    {
      value: 'tierPricing',
      label: 'قیمت‌گذاری پلکانی و خرید عمده',
      description: 'قیمت‌های ویژه برای خریدهای عمده و مشتریان مختلف'
    },
    {
      value: 'discountCodes',
      label: 'کدهای تخفیف و کوپن‌های مخصوص',
      description: 'مدیریت کدهای تخفیف و کوپن‌های ویژه محصول'
    }
  ],
  inventory: [
    {
      value: 'inventoryReports',
      label: 'گزارشات و هشدارها',
      description: 'نمایش گزارش‌های وضعیت موجودی و هشدارهای کمبود یا مازاد'
    },
    {
      value: 'inventoryMovements',
      label: 'حرکات موجودی و انتقالات',
      description: 'مدیریت و مشاهده انتقالات و حرکات موجودی بین انبارها'
    },
    {
      value: 'inventoryForecast',
      label: 'پیش‌بینی و برنامه‌ریزی موجودی',
      description: 'پیش‌بینی نیاز آینده و برنامه‌ریزی تامین موجودی کالا'
    }
  ],
  shipping: [
    {
      value: 'mainShipping',
      label: 'تنظیمات ارسال اصلی',
      description: 'مدیریت روش‌های ارسال و انتخاب پیش‌فرض'
    },
    {
      value: 'shippingCosts',
      label: 'هزینه‌های ارسال',
      description: 'تعیین هزینه‌های ارسال بر اساس شرایط مختلف'
    },
    {
      value: 'shippingSchedule',
      label: 'زمان‌بندی و تحویل',
      description: 'برنامه‌ریزی زمان ارسال و تحویل سفارش'
    },
    {
      value: 'shippingSpecial',
      label: 'تنظیمات ویژه و محدودیت‌ها',
      description: 'محدودیت‌های ارسال، مناطق خاص و شرایط ویژه'
    },
    {
      value: 'packagingLabeling',
      label: 'بسته‌بندی و برچسب‌گذاری',
      description: 'مدیریت نوع بسته‌بندی و برچسب‌های سفارش'
    }
  ],
  images: [
    {
      value: 'gallerySettings',
      label: 'تنظیمات نمایش و گالری',
      description: 'مدیریت نحوه نمایش تصاویر و گالری محصول برای کاربران'
    }
  ],
  variants: [
    {
      value: 'variantManagement',
      label: 'مدیریت تنوع‌های محصول',
      description: 'مدیریت و افزودن انواع مختلف محصول (رنگ، سایز و ...)' 
    },
    {
      value: 'variantPricing',
      label: 'قیمت‌گذاری تنوع‌ها',
      description: 'تعیین قیمت جداگانه برای هر تنوع محصول'
    },
    {
      value: 'variantImages',
      label: 'تصاویر تنوع‌های محصول',
      description: 'اختصاص تصویر به هر تنوع محصول'
    }
  ],
  specs: [
    {
      value: 'newSpecsManagement',
      label: 'مدیریت ویژگی‌های جدید',
      description: 'مدیریت و افزودن ویژگی‌های جدید برای محصول (مانند جنس، ابعاد و ...)' 
    }
  ],
  seo: [
    {
      value: 'openGraphSocial',
      label: 'Open Graph و شبکه‌های اجتماعی',
      description: 'مدیریت تنظیمات اشتراک‌گذاری محصول در شبکه‌های اجتماعی و Open Graph'
    },
    {
      value: 'schemaStructuredData',
      label: 'Schema و داده‌های ساختاریافته',
      description: 'افزودن و مدیریت داده‌های ساختاریافته (Schema) برای بهبود نمایش در نتایج جستجو'
    },
    {
      value: 'seoAnalytics',
      label: 'آنالیز و گزارش SEO',
      description: 'مشاهده و بررسی وضعیت سئو و دریافت گزارش‌های بهینه‌سازی'
    }
  ],
  related: [
    {
      value: 'relatedProductsDisplay',
      label: 'تنظیمات نمایش محصولات مرتبط',
      description: 'مدیریت نحوه نمایش و انتخاب محصولات مرتبط و مکمل برای این محصول'
    }
  ],
  video: [
    {
      value: 'aparatVideos',
      label: 'ویدیوهای آپارات',
      description: 'نمایش و مدیریت ویدیوهای آپارات مرتبط با این محصول'
    }
  ],
// تنظیمات پیش‌فرض - بخش‌های اجباری همیشه فعال، بخش‌های اختیاری قابل تنظیم
const sections = ref({
  // بخش‌های اجباری اطلاعات محصول (همیشه فعال)
  mainInfo: true,
  technicalInfo: true,
  
  // بخش‌های اختیاری اطلاعات محصول
  displaySettings: true,
  scheduling: true,
  management: true,
  strengthsWeaknesses: true,
  
  // بخش‌های تب قیمت‌گذاری
  mainPrices: true,
  priceSettings: true,
  taxSettings: true,
  tierPricing: true,
  discountCodes: true,
  
  // سایر تب‌ها (فعلاً همه فعال)
  pricing: true,
  inventory: true,
  shipping: true,
  images: true,
  variants: true,
  specs: true,
  seo: true,
  related: true,
  video: true,
  inventoryReports: true,
  inventoryMovements: true,
  inventoryForecast: true,
  mainShipping: true,
  shippingCosts: true,
  shippingSchedule: true,
  shippingSpecial: true,
  packagingLabeling: true,
  gallerySettings: true,
  variantManagement: true,
  variantPricing: true,
  variantImages: true,
  newSpecsManagement: true,
  openGraphSocial: true,
  schemaStructuredData: true,
  seoAnalytics: true,
  relatedProductsDisplay: true,
  aparatVideos: true
})

onMounted(() => {
  // بارگذاری تنظیمات از localStorage
  loadSettings()
})

function loadSettings() {
  if (typeof window !== 'undefined') {
    const savedSettings = localStorage.getItem('productPageSections')
    if (savedSettings) {
      try {
        const parsed = JSON.parse(savedSettings)
        // بخش‌های اجباری همیشه فعال باقی می‌مانند
        sections.value = { 
          ...sections.value, 
          ...parsed,
          // اطمینان از فعال بودن بخش‌های اجباری
          mainInfo: true,
          technicalInfo: true
        }
      } catch {
        // خطا در بارگذاری تنظیمات
      }
    }
  }
}

function getCurrentTabSections() {
  if (activeTab.value === 'info') {
    return infoSections.value
  }
  return otherTabSections.value[activeTab.value] || []
}

function save() {
  // اطمینان از فعال بودن بخش‌های اجباری قبل از ذخیره
  const settingsToSave = {
    ...sections.value,
    mainInfo: true,
    technicalInfo: true
  }
  
  // ذخیره تنظیمات در localStorage
  if (typeof window !== 'undefined') {
    localStorage.setItem('productPageSections', JSON.stringify(settingsToSave))
  }
  
  // ارسال تنظیمات به کامپوننت والد
  emit('sections-updated', settingsToSave)
  emit('update:show', false)
}
</script>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style> 
