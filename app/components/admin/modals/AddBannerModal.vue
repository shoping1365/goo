<template>
  <div>
    <!-- تب‌های انتخاب نوع -->
    <div class="flex border-b border-gray-200 mb-6">
      <button 
        :class="[
          'px-6 py-3 font-medium text-sm border-b-2 transition-colors',
          activeTab === 'banner' 
            ? 'border-blue-500 text-blue-600' 
            : 'border-transparent text-gray-500 hover:text-gray-700'
        ]"
        @click="activeTab = 'banner'"
      >
        بنر ساده
      </button>
      <button 
        :class="[
          'px-6 py-3 font-medium text-sm border-b-2 transition-colors',
          activeTab === 'slider' 
            ? 'border-blue-500 text-blue-600' 
            : 'border-transparent text-gray-500 hover:text-gray-700'
        ]"
        @click="activeTab = 'slider'"
      >
        اسلایدر
      </button>
    </div>

    <!-- محتوای تب بنر -->
    <div v-if="activeTab === 'banner'" class="flex flex-col gap-8">
      <!-- بخش بنر -->
      <div class="bg-white border-2 border-blue-400 shadow-lg rounded-2xl p-6 flex flex-col gap-6 transition-all">
        <div class="flex flex-row items-center gap-2">
          <label class="font-bold text-sm mb-1">تصویر بنر</label>
          <div class="flex-1"></div>
          <label class="bg-gray-200 text-gray-700 px-3 py-1 rounded cursor-pointer text-center ml-2">
            افزودن مستقیم
            <input type="file" accept="image/*" class="hidden" @change="onBannerImageChange" />
          </label>
          <button type="button" class="bg-blue-500 text-white px-3 py-1 rounded" @click="showBannerLibrary = true">افزودن از کتابخانه</button>
        </div>
        <div class="flex flex-col items-center w-full mt-4">
          <div v-if="bannerImage" class="mb-2 w-full">
            <img :src="bannerImage" alt="پیش نمایش تصویر بنر" class="w-full max-w-md h-32 object-cover rounded border" />
          </div>
          <div v-else class="w-full max-w-md h-32 flex items-center justify-center bg-gray-100 rounded border text-gray-400">پیش‌نمایش تصویر بنر</div>
          <div v-if="bannerImageInfo" class="text-xs text-gray-500 mt-2">
            سایز: {{ bannerImageInfo.width }}x{{ bannerImageInfo.height }} px &nbsp; | &nbsp;
            حجم: {{ bannerImageInfo.sizeKB }} KB
          </div>
          <input v-model="bannerTitle" type="text" class="mt-3 w-full max-w-md border border-gray-300 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-400" placeholder="عنوان بنر" />
          <input v-model="bannerLink" type="text" class="mt-2 w-full max-w-md border border-gray-300 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-400" placeholder="لینک بنر (اختیاری)" />
        </div>
      </div>
      <div class="flex justify-end mt-2">
        <button type="button" class="bg-purple-600 text-white px-8 py-2 rounded-md font-bold hover:bg-purple-700 transition-colors" @click="handleSave">ذخیره</button>
      </div>
    </div>

    <!-- محتوای تب اسلایدر -->
    <div v-if="activeTab === 'slider'" class="flex flex-col gap-8">
      <AddSliderModal @save="onSliderSave" />
    </div>

    <!-- مودال کتابخانه بنر -->
    <MediaLibraryModal v-model="showBannerLibrary" file-type="image" default-category="banners" @confirm="onSelectBannerFromLibrary" />
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import MediaLibraryModal from '~/components/media/MediaLibraryModal.vue'
import AddSliderModal from './AddSliderModal.vue'

const emit = defineEmits(['save'])

// تب فعال
const activeTab = ref<'banner' | 'slider'>('banner')

const bannerImage = ref<string | null>(null)
const bannerImageInfo = ref<{ width: number, height: number, sizeKB: number } | null>(null)
const bannerTitle = ref('')
const bannerLink = ref('')
const showBannerLibrary = ref(false)

function onBannerImageChange(e: Event) {
  const file = (e.target as HTMLInputElement).files?.[0]
  if (file) {
    bannerImage.value = URL.createObjectURL(file)
    const img = new window.Image()
    img.onload = () => {
      bannerImageInfo.value = {
        width: img.width,
        height: img.height,
        sizeKB: Math.round(file.size / 1024)
      }
    }
    img.src = bannerImage.value
  } else {
    bannerImage.value = null
    bannerImageInfo.value = null
  }
}

interface MediaFile {
  id: number
  url: string
  thumbnail: string
  type: string
  name: string
  size: number
  category: string
}

function onSelectBannerFromLibrary(files: MediaFile[]) {
  if (files && files.length > 0) {
    bannerImage.value = files[0].url
    const img = new window.Image()
    img.onload = () => {
      bannerImageInfo.value = {
        width: img.width,
        height: img.height,
        sizeKB: Math.round((files[0].size || 0) / 1024)
      }
    }
    img.src = files[0].url
  }
}

function handleSave() {
  emit('save', {
    type: 'banner',
    data: {
      image: bannerImage.value,
      title: bannerTitle.value,
      link: bannerLink.value
    }
  })
  // پاک کردن مقادیر پس از ذخیره (اختیاری)
  bannerImage.value = null
  bannerImageInfo.value = null
  bannerTitle.value = ''
  bannerLink.value = ''
}

interface SliderData {
  slides: Array<{
    desktopImage: string | null
    mobileImage: string | null
    desktopTitle: string
    mobileTitle: string
    desktopLink: string
    mobileLink: string
    description: string
  }>
}

function onSliderSave(sliderData: SliderData) {
  emit('save', {
    type: 'slider',
    data: sliderData
  })
}
</script> 