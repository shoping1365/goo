<template>
  <!-- Modal for adding/editing banner -->
  <div v-if="isVisible" class="fixed inset-0 flex items-center justify-center z-50 p-6" @click="closeModal">
    <div class="bg-white rounded-lg border-4 border-blue-300 max-w-2xl w-full max-h-[90vh] overflow-y-auto" style="box-shadow: 0 0 20px 8px rgba(173, 216, 230, 0.3), 0 0 40px 12px rgba(173, 216, 230, 0.15);" @click.stop>
      <div class="p-6">
        <div class="flex justify-between items-center mb-6">
          <h3 class="text-xl font-bold text-gray-800" style="font-family: 'Yekan', sans-serif;">
            {{ isEditing ? 'ویرایش بنر' : 'افزودن بنر جدید' }}
          </h3>
          <button
            @click="closeModal"
            class="text-gray-400 hover:text-gray-600"
          >
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>

        <form @submit.prevent="handleSave" class="space-y-4">
          <!-- عنوان بنر -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1" style="font-family: 'Yekan', sans-serif;">عنوان بنر *</label>
            <input
              v-model="bannerData.title"
              type="text"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
              placeholder="عنوان بنر را وارد کنید"
              style="font-family: 'Yekan', sans-serif;"
              :class="{ 'border-red-500 focus:ring-red-500': titleError }"
              ref="titleInputRef"
              @input="validateTitle"
              @blur="validateTitle"
            />
            <!-- Custom Error Message -->
            <p v-if="titleError" class="mt-1 text-sm text-red-600" style="font-family: 'Yekan', sans-serif;">
              {{ titleError }}
            </p>
          </div>

          <!-- توضیحات -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1" style="font-family: 'Yekan', sans-serif;">توضیحات</label>
            <textarea
              v-model="bannerData.description"
              rows="3"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
              placeholder="توضیحات بنر را وارد کنید"
              style="font-family: 'Yekan', sans-serif;"
            ></textarea>
          </div>

          <!-- تصویر -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1" style="font-family: 'Yekan', sans-serif;">
              تصویر {{ deviceType === 'mobile' ? 'موبایل' : 'دسکتاپ' }} *
            </label>
            <div class="space-y-4">
              <!-- Custom File Upload Area -->
              <div class="border-2 border-dashed rounded-lg p-6 text-center transition-colors"
                   :class="imageError ? 'border-red-400 bg-red-50' : 'border-gray-300 hover:border-blue-400'">
                <div v-if="!currentImage" class="space-y-3">
                  <svg class="w-12 h-12 text-gray-400 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  </svg>

                  <!-- Library Selection Button (Primary) -->
                  <button
                    type="button"
                    @click="handleOpenMediaLibrary"
                    class="inline-flex items-center font-medium rounded-lg focus:outline-none transition-all duration-200 shadow-md bg-gradient-to-r from-green-500 to-green-600 text-white hover:from-green-600 hover:to-green-700 px-5 py-3 text-sm cursor-pointer mb-3"
                    style="font-family: 'Yekan', sans-serif;"
                  >
                    <svg class="w-5 h-5 inline-block ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    </svg>
                    انتخاب از کتابخانه تصاویر
                  </button>
                </div>

                <!-- Image preview -->
                <div v-else class="space-y-3">
                  <img
                    :src="currentImage"
                    alt="پیش نمایش"
                    class="max-w-full max-h-48 object-cover rounded mx-auto border border-gray-200"
                  />
                  <div class="flex gap-2 justify-center flex-wrap">
                    <button
                      type="button"
                      @click="handleOpenMediaLibrary"
                      class="bg-green-500 text-white px-4 py-2 rounded-lg text-sm hover:bg-green-600 transition-colors"
                      style="font-family: 'Yekan', sans-serif;"
                    >
                      <svg class="w-4 h-4 inline-block ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      </svg>
                      انتخاب از کتابخانه
                    </button>

                    <button
                      type="button"
                      @click="handleRemoveImage"
                      class="bg-red-500 text-white px-4 py-2 rounded-lg text-sm hover:bg-red-600 transition-colors"
                      style="font-family: 'Yekan', sans-serif;"
                    >
                      <svg class="w-4 h-4 inline-block ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                      </svg>
                      حذف تصویر
                    </button>
                  </div>
                </div>

                <!-- Image Error Message -->
                <p v-if="imageError" class="mt-1 text-sm text-red-600" style="font-family: 'Yekan', sans-serif;">
                  {{ imageError }}
                </p>
              </div>
            </div>
          </div>

          <!-- لینک -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1" style="font-family: 'Yekan', sans-serif;">لینک (اختیاری)</label>
            <input
              v-model="bannerData.link"
              type="url"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
              placeholder="https://example.com"
              style="font-family: 'Yekan', sans-serif;"
            />
          </div>

          <!-- باز کردن در صفحه جدید -->
          <div v-if="bannerData.link" class="flex items-center gap-3">
            <input
              type="checkbox"
              id="openInNewTabBannerCheckbox"
              v-model="bannerData.openInNewTab"
              class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 focus:ring-2"
            />
            <label for="openInNewTabBannerCheckbox" class="text-sm font-medium text-gray-700" style="font-family: 'Yekan', sans-serif;">
              باز کردن در صفحه جدید
            </label>
          </div>

          <!-- دکمه‌های اقدام -->
          <div class="flex justify-between items-center pt-4 border-t border-gray-200">
            <!-- کنترل نمایش عنوان -->
            <div class="flex items-center gap-3">
              <input
                type="checkbox"
                id="showTitleCheckbox"
                v-model="showTitleInBanner"
                class="w-4 h-4 text-purple-600 bg-gray-100 border-gray-300 rounded focus:ring-purple-500 focus:ring-2"
              />
              <label for="showTitleCheckbox" class="text-sm font-medium text-gray-700">
                نمایش عنوان در بنر
              </label>
            </div>

            <!-- دکمه‌های اقدام -->
            <div class="flex gap-3">
              <button
                type="button"
                @click="closeModal"
                class="px-4 py-2 text-gray-600 bg-gray-100 rounded-md hover:bg-gray-200 transition-colors"
                style="font-family: 'Yekan', sans-serif;"
              >
                لغو
              </button>
              <button
                type="submit"
                class="px-6 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors"
                style="font-family: 'Yekan', sans-serif;"
              >
                {{ isEditing ? 'به‌روزرسانی' : 'افزودن' }}
              </button>
            </div>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import type { BannerItem } from '~/types/widget'

// Props
interface Props {
  isVisible: boolean
  isEditing: boolean
  bannerData: BannerItem
  showTitle: boolean
  deviceType?: 'desktop' | 'mobile'
}

const props = withDefaults(defineProps<Props>(), {
  isVisible: false,
  isEditing: false,
  bannerData: () => ({
    title: '',
    description: '',
    image: '',
    mobile_image: '',
    link: '',
    order: 1,
    status: 'active' as const
  }),
  showTitle: true,
  deviceType: 'desktop'
})

// Emits
const emit = defineEmits<{
  'update:isVisible': [value: boolean]
  'update:showTitle': [value: boolean]
  'update:bannerData': [bannerData: BannerItem]
  'save': [bannerData: BannerItem, showTitle: boolean]
  'open-media-library': []
  'remove-image': []
}>()

// Methods
const closeModal = () => {
  // Clear errors when closing
  titleError.value = ''
  imageError.value = ''
  emit('update:isVisible', false)
}

// Refs for form elements
const titleInputRef = ref<HTMLInputElement>()

// Title display control
const showTitleInBanner = ref(props.showTitle)

// Error states
const titleError = ref<string>('')
const imageError = ref<string>('')

// Computed property for current image based on device type
const currentImage = computed(() => {
  return props.deviceType === 'mobile' ? props.bannerData.mobile_image : props.bannerData.image
})

const validateTitle = () => {
  if (!props.bannerData.title?.trim()) {
    titleError.value = 'لطفاً عنوان بنر را وارد کنید.'
  } else {
    titleError.value = ''
  }
}

const handleSave = () => {
  // Validate title field
  validateTitle()

  // Check if there are validation errors
  if (titleError.value) {
    // Focus on the title input to show the error
    if (titleInputRef.value) {
      titleInputRef.value.focus()
    }
    return
  }

  // Check for required image
  if (!currentImage.value) {
    imageError.value = 'تصویر بنر الزامی است'
    // Scroll to the image section
    const imageSection = document.querySelector('.space-y-4')
    if (imageSection) {
      imageSection.scrollIntoView({ behavior: 'smooth', block: 'center' })
    }
    return
  }

  emit('save', props.bannerData, showTitleInBanner.value)
  closeModal()
}

const handleOpenMediaLibrary = () => {
  emit('open-media-library')
}

const handleRemoveImage = () => {
  imageError.value = ''
  emit('remove-image')
}

// Handle custom file selection
const handleFileSelect = (event: Event) => {
  const target = event.target as HTMLInputElement
  const file = target.files?.[0]

  if (file) {
    // Validate file type
    if (!file.type.startsWith('image/')) {
      imageError.value = 'لطفا یک فایل تصویر انتخاب کنید'
      return
    }

    // Validate file size (max 5MB)
    if (file.size > 5 * 1024 * 1024) {
      imageError.value = 'حجم فایل نباید بیشتر از 5 مگابایت باشد'
      return
    }

    // Clear any previous errors
    imageError.value = ''

    const reader = new FileReader()
    reader.onload = (e) => {
      const result = e.target?.result as string
      // Create updated banner data with new image
      const updatedBannerData = {
        ...props.bannerData,
        image: result
      }
      emit('update:bannerData', updatedBannerData)
    }
    reader.readAsDataURL(file)
  }
}

// Initialize showTitleInBanner when component mounts
onMounted(() => {
  showTitleInBanner.value = props.showTitle
})
</script>
