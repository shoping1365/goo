          <template>
          <!-- Modal for adding/editing slide -->
          <div v-if="isVisible" class="fixed inset-0 flex items-center justify-center z-50 px-4 py-4" @click="closeModal">
                    <div class="bg-white rounded-lg border-4 border-blue-300 max-w-2xl w-full max-h-[90vh] overflow-y-auto" style="box-shadow: 0 0 20px 8px rgba(173, 216, 230, 0.3), 0 0 40px 12px rgba(173, 216, 230, 0.15);" @click.stop>
               <div class="px-4 py-4">
                       <div class="flex justify-between items-center mb-6">
          <h3 class="text-xl font-bold text-gray-800" style="font-family: 'Yekan', sans-serif;">
            {{ isEditing ? 'ویرایش اسلاید' : 'افزودن اسلاید جدید' }}
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
                                <!-- عنوان اسلاید -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1" style="font-family: 'Yekan', sans-serif;">عنوان اسلاید *</label>
              <input
                ref="titleInputRef"
                v-model="slideData.title"
                type="text"
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
                placeholder="عنوان اسلاید را وارد کنید"
                style="font-family: 'Yekan', sans-serif;"
                :class="{ 'border-red-500 focus:ring-red-500': titleError }"
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
                    v-model="slideData.description"
                    rows="3"
                    class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
                    placeholder="توضیحات اسلاید را وارد کنید"
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
                    <div class="border-2 border-dashed rounded-lg px-4 py-4 text-center transition-colors"
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
                      v-model="slideData.link"
                      type="url"
                      class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
                      placeholder="https://example.com"
                      style="font-family: 'Yekan', sans-serif;"
                    />
                    </div>

                    <!-- باز کردن در صفحه جدید -->
                    <div v-if="slideData.link" class="flex items-center gap-3">
                    <input
                      type="checkbox"
                      id="openInNewTabCheckbox"
                      v-model="slideData.openInNewTab"
                      class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 focus:ring-2"
                    />
                    <label for="openInNewTabCheckbox" class="text-sm font-medium text-gray-700" style="font-family: 'Yekan', sans-serif;">
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
                         v-model="showTitleValue"
                         class="w-4 h-4 text-purple-600 bg-gray-100 border-gray-300 rounded focus:ring-purple-500 focus:ring-2"
                    />
                    <label for="showTitleCheckbox" class="text-sm font-medium text-gray-700">
                         نمایش عنوان در اسلایدر
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
          import { ref, computed } from 'vue'
          import type { SlideItem } from '~/types/widget'

          // Props
          interface Props {
          isVisible: boolean
          isEditing: boolean
          slideData: SlideItem
          showTitle: boolean
          deviceType?: 'desktop' | 'mobile'
          }

          const props = withDefaults(defineProps<Props>(), {
          isVisible: false,
          isEditing: false,
          slideData: () => ({
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
  'update:slideData': [slideData: SlideItem]
  'save': [slideData: SlideItem, showTitle: boolean]
  'open-media-library': []
  'remove-image': []
}>()

          // Refs for form elements
          const titleInputRef = ref<HTMLInputElement | null>(null)

          // Error states
          const titleError = ref('')
          const imageError = ref('')

          // Computed property for current image based on device type
          const currentImage = computed(() => {
            return props.deviceType === 'mobile' ? props.slideData.mobile_image : props.slideData.image
          })

          // Computed property for showTitle with getter/setter
          const showTitleValue = computed({
          get: () => props.showTitle,
          set: (value: boolean) => emit('update:showTitle', value)
          })

          // Methods
          const closeModal = () => {
            // Clear errors when closing
            titleError.value = ''
            imageError.value = ''
            emit('update:isVisible', false)
          }

          const validateTitle = () => {
            if (!props.slideData.title?.trim()) {
              titleError.value = 'لطفاً عنوان اسلاید را وارد کنید.'
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
              const titleInput = titleInputRef.value
              if (titleInput) {
                titleInput.focus()
              }
              return
            }

            // Check for required image - at least one image must be uploaded
            if (!currentImage.value) {
              imageError.value = 'لطفاً تصویر را آپلود کنید.'
              // Scroll to the image section
              const imageSection = document.querySelector('.space-y-4')
              if (imageSection) {
                imageSection.scrollIntoView({ behavior: 'smooth', block: 'center' })
              }
              return
            }

            emit('save', props.slideData, props.showTitle)
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
      // Create updated slide data with new image based on device type
      const updatedSlideData = {
        ...props.slideData,
        // بر اساس deviceType، فیلد مناسب را تغییر بده
        ...(props.deviceType === 'mobile' 
          ? { mobile_image: result }
          : { image: result }
        )
      }
      emit('update:slideData', updatedSlideData)
    }
    reader.readAsDataURL(file)
  }
}
          </script>
