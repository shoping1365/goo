<template>
  <div>
    <div class="flex flex-col gap-8">
      <!-- عنوان و دکمه افزودن اسلاید جدید -->
      <div class="flex justify-between items-center">
        <h3 class="text-xl font-bold text-gray-800">اسلایدهای اسلایدر</h3>
        <button 
          type="button" 
          class="bg-green-500 text-white px-4 py-2 rounded-lg font-semibold hover:bg-green-600 transition-colors flex items-center gap-2"
          @click="addNewSlide"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
          </svg>
          افزودن اسلاید جدید
        </button>
      </div>

      <!-- تنظیمات اسلایدر -->
      <div class="bg-gray-50 border border-gray-200 rounded-lg px-4 py-4">
        <h4 class="text-lg font-semibold text-gray-800 mb-4">تنظیمات اسلایدر</h4>
        
        <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
          <!-- تنظیمات پس زمینه -->
          <div class="space-y-4">
            <label class="font-bold text-sm text-gray-700">پس زمینه</label>
            
            <!-- فعال/غیرفعال کردن پس زمینه -->
            <div class="flex items-center gap-3">
              <input 
                type="checkbox" 
                id="bgEnabled"
                v-model="sliderSettings.backgroundColorEnabled"
                class="w-4 h-4 text-blue-600 rounded focus:ring-blue-500"
              />
              <label for="bgEnabled" class="text-sm text-gray-700">فعال کردن پس زمینه</label>
            </div>
            
            <!-- انتخاب رنگ پس زمینه -->
            <div v-if="sliderSettings.backgroundColorEnabled" class="space-y-2">
              <label class="text-sm text-gray-600">رنگ پس زمینه</label>
              <div class="flex items-center gap-3">
                <input 
                  type="color" 
                  v-model="sliderSettings.backgroundColor"
                  class="w-12 h-10 border border-gray-300 rounded cursor-pointer"
                />
                <span class="text-sm text-gray-600">{{ sliderSettings.backgroundColor }}</span>
              </div>
            </div>
          </div>
          
          <!-- تنظیمات ارتفاع -->
          <div class="space-y-4">
            <label class="font-bold text-sm text-gray-700">ارتفاع اسلایدر</label>
            <div class="space-y-2">
              <label class="text-sm text-gray-600">انتخاب ارتفاع (پیکسل)</label>
              <div class="flex gap-2">
                <button 
                  v-for="height in [400, 500, 600]" 
                  :key="height"
                  @click="sliderSettings.height = height"
                  :class="[
                    'px-4 py-2 rounded-lg font-medium transition-colors',
                    sliderSettings.height === height 
                      ? 'bg-blue-600 text-white' 
                      : 'bg-gray-200 text-gray-700 hover:bg-gray-300'
                  ]"
                >
                  {{ height }}px
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- لیست اسلایدها -->
      <div v-if="slides.length === 0" class="text-center py-12 text-gray-500">
        هنوز اسلایدی اضافه نشده است. روی دکمه "افزودن اسلاید جدید" کلیک کنید.
      </div>

      <!-- نمایش اسلایدها -->
      <div v-for="(slide, index) in slides" :key="index" class="bg-white border-2 border-blue-400 shadow-lg rounded-2xl px-4 py-4">
        <!-- هدر اسلاید -->
        <div class="flex justify-between items-center mb-4">
          <h4 class="text-lg font-semibold text-blue-800">اسلاید {{ index + 1 }}</h4>
          <div class="flex gap-2">
            <button 
              v-if="index > 0"
              @click="moveSlide(index, 'up')" 
              class="bg-gray-500 text-white p-2 rounded hover:bg-gray-600 transition-colors"
              title="انتقال به بالا"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 15l7-7 7 7" />
              </svg>
            </button>
            <button 
              v-if="index < slides.length - 1"
              @click="moveSlide(index, 'down')" 
              class="bg-gray-500 text-white p-2 rounded hover:bg-gray-600 transition-colors"
              title="انتقال به پایین"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
              </svg>
            </button>
            <button 
              @click="removeSlide(index)" 
              class="bg-red-500 text-white p-2 rounded hover:bg-red-600 transition-colors"
              title="حذف اسلاید"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
              </svg>
            </button>
          </div>
        </div>

        <!-- محتوای اسلاید -->
        <div class="grid grid-cols-1 lg:grid-cols-2 gapx-4 py-4">
          <!-- بخش دسکتاپ -->
          <div class="space-y-4">
            <label class="font-bold text-sm text-gray-700">تصویر دسکتاپ</label>
            <div class="flex flex-col items-center w-full">
              <div v-if="slide.desktopImage" class="mb-2 w-full">
                <img :src="slide.desktopImage" alt="پیش نمایش تصویر دسکتاپ" class="w-full h-32 object-cover rounded border" />
              </div>
              <div v-else class="w-full h-32 flex items-center justify-center bg-gray-100 rounded border text-gray-400 text-sm">
                پیش‌نمایش تصویر دسکتاپ
              </div>
              <button 
                type="button" 
                class="mt-2 bg-blue-500 text-white px-3 py-1 rounded text-sm hover:bg-blue-600 transition-colors"
                @click="openMediaLibrary('desktop', index)"
              >
                انتخاب تصویر دسکتاپ
              </button>
            </div>
            <input 
              v-model="slide.desktopTitle" 
              type="text" 
              class="w-full border border-gray-300 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-400" 
              placeholder="عنوان تصویر دسکتاپ" 
            />
            <input 
              v-model="slide.desktopLink" 
              type="text" 
              class="w-full border border-gray-300 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-400" 
              placeholder="لینک تصویر دسکتاپ (اختیاری)" 
            />
          </div>

          <!-- بخش موبایل -->
          <div class="space-y-4">
            <label class="font-bold text-sm text-gray-700">تصویر موبایل</label>
            <div class="flex flex-col items-center w-full">
              <div v-if="slide.mobileImage" class="mb-2 w-full">
                <img :src="slide.mobileImage" alt="پیش نمایش تصویر موبایل" class="w-full h-32 object-cover rounded border" />
              </div>
              <div v-else class="w-full h-32 flex items-center justify-center bg-gray-100 rounded border text-gray-400 text-sm">
                پیش‌نمایش تصویر موبایل
              </div>
              <button 
                type="button" 
                class="mt-2 bg-blue-500 text-white px-3 py-1 rounded text-sm hover:bg-blue-600 transition-colors"
                @click="openMediaLibrary('mobile', index)"
              >
                انتخاب تصویر موبایل
              </button>
            </div>
            <input 
              v-model="slide.mobileTitle" 
              type="text" 
              class="w-full border border-gray-300 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-400" 
              placeholder="عنوان تصویر موبایل" 
            />
            <input 
              v-model="slide.mobileLink" 
              type="text" 
              class="w-full border border-gray-300 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-400" 
              placeholder="لینک تصویر موبایل (اختیاری)" 
            />
          </div>
        </div>

        <!-- توضیحات -->
        <div class="mt-4">
          <label class="font-bold text-sm text-gray-700">توضیحات (اختیاری)</label>
          <textarea 
            v-model="slide.description" 
            rows="2"
            class="w-full mt-1 border border-gray-300 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-400" 
            placeholder="توضیحات اسلاید را وارد کنید"
          ></textarea>
        </div>
      </div>

      <!-- دکمه ذخیره -->
      <div class="flex justify-end mt-6" v-if="slides.length > 0">
        <button 
          type="button" 
          class="bg-purple-600 text-white px-8 py-3 rounded-lg font-bold hover:bg-purple-700 transition-colors"
          @click="handleSave"
          :disabled="!isFormValid"
        >
          ذخیره اسلایدر
        </button>
      </div>
    </div>

    <!-- مودال کتابخانه رسانه -->
    <MediaLibraryModal 
      v-model="showMediaLibrary" 
      file-type="image" 
      default-category="banners" 
      @confirm="onSelectFromLibrary" 
    />
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import MediaLibraryModal from '~/components/media/MediaLibraryModal.vue'

const props = defineProps<{
  modelValue?: boolean
}>()
const emit = defineEmits(['save'])

// State
const slides = ref<Array<{
  desktopImage: string | null
  mobileImage: string | null
  desktopTitle: string
  mobileTitle: string
  desktopLink: string
  mobileLink: string
  description: string
}>>([])

// تنظیمات اسلایدر
const sliderSettings = ref({
  backgroundColor: '#ffffff', // پس زمینه پیش‌فرض
  backgroundColorEnabled: true, // آیا پس زمینه فعال باشه
  height: 500 // ارتفاع پیش‌فرض
})

const showMediaLibrary = ref(false)
const currentMediaType = ref<'desktop' | 'mobile'>('desktop')
const currentSlideIndex = ref(0)

// Computed
const isFormValid = computed(() => {
  console.log('isFormValid computed - slides:', slides.value)
  
  if (slides.value.length === 0) {
    console.log('هیچ اسلایدی وجود ندارد')
    return false
  }
  
  const validSlides = slides.value.every((slide, index) => {
    // برای هر اسلاید، حداقل یکی (دسکتاپ یا موبایل) باید کامل باشه
    const desktopValid = slide.desktopImage && slide.desktopTitle
    const mobileValid = slide.mobileImage && slide.mobileTitle
    
    console.log(`اسلاید ${index + 1}:`, {
      desktopValid,
      mobileValid,
      desktopImage: slide.desktopImage,
      desktopTitle: slide.desktopTitle,
      mobileImage: slide.mobileImage,
      mobileTitle: slide.mobileTitle
    })
    
    // حداقل یکی باید کامل باشه
    const isValid = desktopValid || mobileValid
    
    if (!isValid) {
      console.log(`اسلاید ${index + 1} معتبر نیست - هیچ کدام کامل نیست`)
    }
    
    return isValid
  })
  
  console.log('isFormValid result:', validSlides)
  return validSlides
})

// Methods
const addNewSlide = () => {
  slides.value.push({
    desktopImage: null,
    mobileImage: null,
    desktopTitle: '',
    mobileTitle: '',
    desktopLink: '',
    mobileLink: '',
    description: ''
  })
}

const removeSlide = (index: number) => {
  if (slides.value.length > 1) {
    slides.value.splice(index, 1)
  }
}

const moveSlide = (index: number, direction: 'up' | 'down') => {
  if (direction === 'up' && index > 0) {
    const temp = slides.value[index]
    slides.value[index] = slides.value[index - 1]
    slides.value[index - 1] = temp
  } else if (direction === 'down' && index < slides.value.length - 1) {
    const temp = slides.value[index]
    slides.value[index] = slides.value[index + 1]
    slides.value[index + 1] = temp
  }
}

const openMediaLibrary = (type: 'desktop' | 'mobile', slideIndex: number) => {
  currentMediaType.value = type
  currentSlideIndex.value = slideIndex
  showMediaLibrary.value = true
}

const onSelectFromLibrary = (files: any[]) => {
  if (files && files.length > 0) {
    const slide = slides.value[currentSlideIndex.value]
    const file = files[0]
    
    if (currentMediaType.value === 'desktop') {
      slide.desktopImage = file.url
    } else {
      slide.mobileImage = file.url
    }
  }
  showMediaLibrary.value = false
}

const handleSave = () => {
  console.log('handleSave اجرا شد')
  console.log('isFormValid:', isFormValid.value)
  console.log('slides:', slides.value)
  console.log('sliderSettings:', sliderSettings.value)
  
  if (!isFormValid.value) {
    console.log('فرم معتبر نیست!')
    return
  }
  
  emit('save', {
    slides: slides.value,
    settings: {
      backgroundColor: sliderSettings.value.backgroundColorEnabled ? sliderSettings.value.backgroundColor : null,
      backgroundColorEnabled: sliderSettings.value.backgroundColorEnabled,
      height: sliderSettings.value.height
    }
  })
}

// Initialize with one slide
addNewSlide()
</script> 