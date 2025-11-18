<template>
  <div class="bg-white rounded-2xl shadow-xl border border-gray-100 p-8 mb-8">
    <div class="flex items-center mb-6">
      <div class="p-3 bg-gradient-to-r from-amber-400 to-orange-600 rounded-xl shadow-lg">
        <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
        </svg>
      </div>
      <div class="mr-4">
        <h2 class="text-2xl font-bold text-gray-900">بخش اختیاری</h2>
        <p class="text-gray-600 mt-1">موارد اضافی برای بهبود تجربه شما</p>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <!-- ستون اول -->
      <div class="space-y-6">
        <!-- آپلود تصاویر -->
        <div class="bg-gradient-to-r from-blue-50 to-cyan-50 rounded-xl p-6 border border-blue-100">
          <div class="flex items-center mb-4">
            <div class="p-2 bg-blue-500 rounded-lg mr-3">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900">آپلود تصاویر</h3>
          </div>

          <div class="space-y-4">
            <div class="border-2 border-dashed border-gray-300 rounded-lg p-6 text-center hover:border-blue-400 transition-colors">
              <input
                  ref="imageInput"
                  type="file"
                  multiple
                  accept="image/*"
                  @change="handleImageUpload"
                  class="hidden"
              />
              <button
                @click="() => imageInput?.click()"
                  class="flex flex-col items-center space-y-2 space-y-reverse"
              >
                <svg class="w-12 h-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"></path>
                </svg>
                <span class="text-gray-600">برای آپلود تصویر کلیک کنید</span>
                <span class="text-sm text-gray-400">حداکثر 5 تصویر، هر کدام کمتر از 5MB</span>
              </button>
            </div>

            <!-- نمایش تصاویر انتخاب شده -->
            <div v-if="selectedImages.length > 0" class="grid grid-cols-3 gap-2">
              <div
                  v-for="(image, index) in selectedImages"
                  :key="index"
                  class="relative group"
              >
                <img
                    :src="image.preview"
                    :alt="`تصویر ${index + 1}`"
                    class="w-full h-20 object-cover rounded-lg"
                />
                <button
                    @click="removeImage(index)"
                    class="absolute top-1 right-1 bg-red-500 text-white rounded-full w-6 h-6 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity"
                >
                  ×
                </button>
              </div>
            </div>
          </div>
        </div>

        <!-- آپلود ویدیو -->
        <div class="bg-gradient-to-r from-purple-50 to-pink-50 rounded-xl p-6 border border-purple-100">
          <div class="flex items-center mb-4">
            <div class="p-2 bg-purple-500 rounded-lg mr-3">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900">آپلود ویدیو</h3>
          </div>

          <div class="border-2 border-dashed border-gray-300 rounded-lg p-6 text-center hover:border-purple-400 transition-colors">
            <input
                ref="videoInput"
                type="file"
                accept="video/*"
                @change="handleVideoUpload"
                class="hidden"
            />
            <button
                @click="() => videoInput?.click()"
                class="flex flex-col items-center space-y-2 space-y-reverse"
            >
              <svg class="w-12 h-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"></path>
              </svg>
              <span class="text-gray-600">برای آپلود ویدیو کلیک کنید</span>
              <span class="text-sm text-gray-400">حداکثر 50MB، فرمت‌های MP4, AVI</span>
            </button>
          </div>

          <!-- نمایش ویدیو انتخاب شده -->
          <div v-if="selectedVideo" class="mt-4">
            <video
                :src="selectedVideo.preview"
                controls
                class="w-full rounded-lg"
            ></video>
            <button
                @click="removeVideo"
                class="mt-2 px-3 py-1 bg-red-500 text-white rounded text-sm hover:bg-red-600 transition-colors"
            >
              حذف ویدیو
            </button>
          </div>
        </div>
      </div>

      <!-- ستون دوم -->
      <div class="space-y-6">
        <!-- تنظیمات حریم خصوصی -->
        <div class="bg-gradient-to-r from-green-50 to-emerald-50 rounded-xl p-6 border border-green-100">
          <div class="flex items-center mb-4">
            <div class="p-2 bg-green-500 rounded-lg mr-3">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900">تنظیمات حریم خصوصی</h3>
          </div>

          <div class="space-y-4">
            <label class="flex items-center space-x-3 space-x-reverse cursor-pointer">
              <input
                  v-model="privacySettings.allowPublication"
                  type="checkbox"
                  class="w-5 h-5 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
              />
              <span class="text-gray-700">اجازه انتشار نظر در سایت</span>
            </label>

            <label class="flex items-center space-x-3 space-x-reverse cursor-pointer">
              <input
                  v-model="privacySettings.anonymousReview"
                  type="checkbox"
                  class="w-5 h-5 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
              />
              <span class="text-gray-700">نظر ناشناس (بدون نمایش نام)</span>
            </label>

            <label class="flex items-center space-x-3 space-x-reverse cursor-pointer">
              <input
                  v-model="privacySettings.allowContact"
                  type="checkbox"
                  class="w-5 h-5 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
              />
              <span class="text-gray-700">اجازه تماس برای پیگیری</span>
            </label>
          </div>
        </div>

        <!-- خبرنامه و اطلاع‌رسانی -->
        <div class="bg-gradient-to-r from-orange-50 to-amber-50 rounded-xl p-6 border border-orange-100">
          <div class="flex items-center mb-4">
            <div class="p-2 bg-orange-500 rounded-lg mr-3">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-5 5v-5zM4 19h6v-2H4v2zM4 15h6v-2H4v2zM4 11h6V9H4v2zM4 7h6V5H4v2z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900">خبرنامه و اطلاع‌رسانی</h3>
          </div>

          <div class="space-y-4">
            <label class="flex items-center space-x-3 space-x-reverse cursor-pointer">
              <input
                  v-model="newsletterSettings.receiveNewsletter"
                  type="checkbox"
                  class="w-5 h-5 text-orange-600 border-gray-300 rounded focus:ring-orange-500"
              />
              <span class="text-gray-700">دریافت خبرنامه و تخفیف‌ها</span>
            </label>

            <label class="flex items-center space-x-3 space-x-reverse cursor-pointer">
              <input
                  v-model="newsletterSettings.productUpdates"
                  type="checkbox"
                  class="w-5 h-5 text-orange-600 border-gray-300 rounded focus:ring-orange-500"
              />
              <span class="text-gray-700">اطلاع از محصولات جدید</span>
            </label>

            <label class="flex items-center space-x-3 space-x-reverse cursor-pointer">
              <input
                  v-model="newsletterSettings.surveyInvitations"
                  type="checkbox"
                  class="w-5 h-5 text-orange-600 border-gray-300 rounded focus:ring-orange-500"
              />
              <span class="text-gray-700">دعوت به نظرسنجی‌های آینده</span>
            </label>
          </div>
        </div>

        <!-- اطلاعات تماس اختیاری -->
        <div class="bg-gradient-to-r from-teal-50 to-cyan-50 rounded-xl p-6 border border-teal-100">
          <div class="flex items-center mb-4">
            <div class="p-2 bg-teal-500 rounded-lg mr-3">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 4.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900">اطلاعات تماس اختیاری</h3>
          </div>

          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">شماره تلفن</label>
              <input
                  v-model="contactInfo.phone"
                  type="tel"
                  placeholder="09xxxxxxxxx"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-teal-500 focus:border-teal-500"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">آدرس ایمیل</label>
              <input
                  v-model="contactInfo.email"
                  type="email"
                  placeholder="example@email.com"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-teal-500 focus:border-teal-500"
              />
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// Template refs
const imageInput = ref<HTMLInputElement>()
const videoInput = ref<HTMLInputElement>()

// Reactive data for file uploads
const selectedImages = ref<Array<{ preview: string; file: File }>>([])
const selectedVideo = ref<{ preview: string; file: File } | null>(null)

// Reactive data for privacy settings
const privacySettings = ref({
  allowPublication: false,
  anonymousReview: false,
  allowContact: true
})

// Reactive data for newsletter settings
const newsletterSettings = ref({
  receiveNewsletter: false,
  productUpdates: false,
  surveyInvitations: false
})

// Reactive data for contact info
const contactInfo = ref({
  phone: '',
  email: ''
})

// File upload handlers
const handleImageUpload = (event: Event) => {
  const target = event.target as HTMLInputElement
  if (target.files && target.files.length > 0) {
    const files = Array.from(target.files)
    files.forEach(file => {
      if (file.type.startsWith('image/') && selectedImages.value.length < 5) {
        const reader = new FileReader()
        reader.onload = (e) => {
          selectedImages.value.push({
            preview: e.target?.result as string,
            file: file
          })
        }
        reader.readAsDataURL(file)
      }
    })
  }
}

const handleVideoUpload = (event: Event) => {
  const target = event.target as HTMLInputElement
  if (target.files && target.files.length > 0) {
    const file = target.files[0]
    if (file.type.startsWith('video/')) {
      const reader = new FileReader()
      reader.onload = (e) => {
        selectedVideo.value = {
          preview: e.target?.result as string,
          file: file
        }
      }
      reader.readAsDataURL(file)
    }
  }
}

const removeImage = (index: number) => {
  selectedImages.value.splice(index, 1)
}

const removeVideo = () => {
  selectedVideo.value = null
}
</script>

<style scoped>
/* Custom file upload styling */
.border-dashed {
  border-style: dashed;
}

.border-dashed:hover {
  border-color: rgb(147 197 253);
}

/* Smooth transitions */
.transition-colors {
  transition: all 0.2s ease-in-out;
}

/* Custom checkbox styling */
input[type="checkbox"] {
  accent-color: rgb(59 130 246);
}

/* Focus states */
input:focus, textarea:focus {
  outline: none;
  ring: 2px;
  ring-color: rgb(59 130 246);
}
</style>
