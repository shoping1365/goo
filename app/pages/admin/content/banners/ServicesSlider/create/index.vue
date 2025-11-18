<template>
  <div class="services-slider-create-page p-6">
    <div class="max-w-4xl mx-auto">
      <!-- Page Header -->
      <div class="mb-8">
        <div class="flex items-center gap-6 mb-4">
          <NuxtLink 
            to="/admin/content/banners/ServicesSlider"
            class="text-blue-600 hover:text-blue-800 transition-colors"
          >
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"/>
            </svg>
          </NuxtLink>
          <h1 class="text-3xl font-bold text-gray-900">ایجاد اسلایدر خدمات</h1>
        </div>
        <p class="text-gray-600">ایجاد ابزارک جدید برای نمایش خدمات در قالب اسلایدر</p>
      </div>

      <!-- Create Form -->
      <div class="bg-white rounded-lg shadow-lg p-6">
        <form @submit.prevent="handleSubmit" class="space-y-6">
          <!-- Basic Information -->
          <div class="border-b border-gray-200 pb-6">
            <h2 class="text-xl font-semibold text-gray-800 mb-4">اطلاعات پایه</h2>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">عنوان ابزارک</label>
                <input 
                  v-model="form.title"
                  type="text"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  placeholder="مثال: اسلایدر خدمات ما"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">کد یکتا</label>
                <input 
                  v-model="form.code"
                  type="text"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  placeholder="مثال: services-slider"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">صفحه نمایش</label>
                <select 
                  v-model="form.page"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                >
                  <option value="">انتخاب کنید</option>
                  <option value="home">صفحه اصلی</option>
                  <option value="services">صفحات خدمات</option>
                  <option value="about">درباره ما</option>
                  <option value="contact">تماس با ما</option>
                </select>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">ترتیب نمایش</label>
                <input 
                  v-model.number="form.order"
                  type="number"
                  min="1"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                />
              </div>
            </div>
          </div>

          <!-- Slider Configuration -->
          <div class="border-b border-gray-200 pb-6">
            <h2 class="text-xl font-semibold text-gray-800 mb-4">تنظیمات اسلایدر</h2>
            <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">تعداد خدمات در هر اسلاید</label>
                <input 
                  v-model.number="form.slidesPerView"
                  type="number"
                  min="1"
                  max="6"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">ارتفاع اسلایدر (پیکسل)</label>
                <input 
                  v-model.number="form.carouselHeight"
                  type="number"
                  min="200"
                  max="600"
                  step="50"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">فاصله بین اسلایدها (پیکسل)</label>
                <input 
                  v-model.number="form.slideGap"
                  type="number"
                  min="0"
                  max="50"
                  step="5"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                />
              </div>
            </div>
            
            <div class="mt-4 grid grid-cols-1 md:grid-cols-4 gap-6">
              <label class="flex items-center">
                <input 
                  v-model="form.autoplay"
                  type="checkbox"
                  class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                />
                <span class="ml-2 text-sm text-gray-700">پخش خودکار</span>
              </label>
              <label class="flex items-center">
                <input 
                  v-model="form.loop"
                  type="checkbox"
                  class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                />
                <span class="ml-2 text-sm text-gray-700">تکرار بی‌نهایت</span>
              </label>
              <label class="flex items-center">
                <input 
                  v-model="form.showNavigation"
                  type="checkbox"
                  class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                />
                <span class="ml-2 text-sm text-gray-700">نمایش دکمه‌های ناوبری</span>
              </label>
              <label class="flex items-center">
                <input 
                  v-model="form.showIndicators"
                  type="checkbox"
                  class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                />
                <span class="ml-2 text-sm text-gray-700">نمایش نشانگرها</span>
              </label>
            </div>

            <div v-if="form.autoplay" class="mt-4">
              <label class="block text-sm font-medium text-gray-700 mb-2">تاخیر پخش خودکار (میلی‌ثانیه)</label>
              <input 
                v-model.number="form.autoplayDelay"
                type="number"
                min="1000"
                max="10000"
                step="500"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              />
            </div>
          </div>

          <!-- Service Selection -->
          <div class="border-b border-gray-200 pb-6">
            <h2 class="text-xl font-semibold text-gray-800 mb-4">انتخاب خدمات</h2>
            <div class="mb-4">
              <label class="block text-sm font-medium text-gray-700 mb-2">روش انتخاب خدمات</label>
              <select 
                v-model="form.serviceSelectionType"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
                <option value="manual">انتخاب دستی</option>
                <option value="featured">خدمات ویژه</option>
                <option value="popular">خدمات محبوب</option>
                <option value="latest">آخرین خدمات</option>
              </select>
            </div>

            <!-- Manual Service Selection -->
            <div v-if="form.serviceSelectionType === 'manual'" class="space-y-4">
              <div class="flex items-center justify-between">
                <h3 class="text-lg font-medium text-gray-800">خدمات انتخاب شده</h3>
                <button 
                  type="button"
                  @click="addServiceItem"
                  class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors"
                >
                  افزودن خدمت
                </button>
              </div>
              
              <div v-if="form.services.length === 0" class="text-center py-8 text-gray-500">
                هیچ خدمتی انتخاب نشده است
              </div>
              
              <div v-else class="space-y-3">
                <div 
                  v-for="(service, index) in form.services"
                  :key="index"
                  class="flex items-center gap-3 p-3 bg-gray-50 rounded-lg"
                >
                  <div class="flex-1 grid grid-cols-1 md:grid-cols-4 gap-3">
                    <input 
                      v-model="service.title"
                      type="text"
                      placeholder="عنوان خدمت"
                      class="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                    />
                    <input 
                      v-model="service.icon"
                      type="text"
                      placeholder="آیکون (نام کلاس CSS)"
                      class="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                    />
                    <input 
                      v-model="service.link"
                      type="text"
                      placeholder="لینک خدمت"
                      class="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                    />
                    <textarea 
                      v-model="service.description"
                      placeholder="توضیحات کوتاه"
                      rows="2"
                      class="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                    ></textarea>
                  </div>
                  <button 
                    type="button"
                    @click="removeServiceItem(index)"
                    class="text-red-600 hover:text-red-800 p-2"
                  >
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"/>
                    </svg>
                  </button>
                </div>
              </div>
            </div>
          </div>

          <!-- Form Actions -->
          <div class="flex gap-3 justify-end">
            <button 
              type="button"
              @click="saveAsDraft"
              class="px-6 py-3 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors"
            >
              ذخیره به عنوان پیش‌نویس
            </button>
            <button 
              type="submit"
              class="px-6 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
            >
              ایجاد ابزارک
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';

// تعریف definePageMeta و useHead و navigateTo برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
declare const useHead: (head: { title?: string }) => void
declare const navigateTo: (to: string) => Promise<void>

// Page meta definition
definePageMeta({
  layout: 'admin',
  middleware: 'admin'
})

// Page title
useHead({
  title: 'ایجاد اسلایدر خدمات - پنل ادمین'
})

const router = useRouter()

// Form state
const form = ref({
  title: '',
  code: '',
  page: '',
  order: 1,
  slidesPerView: 3,
  carouselHeight: 300,
  slideGap: 20,
  autoplay: true,
  autoplayDelay: 5000,
  loop: true,
  showNavigation: true,
  showIndicators: true,
  serviceSelectionType: 'manual',
  services: [] as Array<{
    title: string
    icon: string
    link: string
    description: string
  }>
})

// Methods
const addServiceItem = () => {
  form.value.services.push({
    title: '',
    icon: '',
    link: '',
    description: ''
  })
}

const removeServiceItem = (index: number) => {
  form.value.services.splice(index, 1)
}

const handleSubmit = async () => {
  try {
    // Validate form
    if (!form.value.title || !form.value.code || !form.value.page) {
      alert('لطفاً تمام فیلدهای اجباری را پر کنید')
      return
    }

    if (form.value.serviceSelectionType === 'manual' && form.value.services.length === 0) {
      alert('لطفاً حداقل یک خدمت انتخاب کنید')
      return
    }

    // Here you would call your API to create the widget
    console.log('Creating services slider widget:', form.value)
    
    // Redirect to management page
    await router.push('/admin/content/banners/ServicesSlider')
  } catch (error) {
    console.error('Error creating widget:', error)
    alert('خطا در ایجاد ابزارک')
  }
}

const saveAsDraft = async () => {
  try {
    // Save as draft logic
    console.log('Saving as draft:', form.value)
    alert('پیش‌نویس ذخیره شد')
  } catch (error) {
    console.error('Error saving draft:', error)
    alert('خطا در ذخیره پیش‌نویس')
  }
}
</script>

<style scoped>
.services-slider-create-page {
  background-color: #f9fafb;
  min-height: 100vh;
}
</style>

