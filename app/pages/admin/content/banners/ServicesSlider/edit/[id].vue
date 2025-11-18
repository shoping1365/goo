<template>
  <div class="services-slider-edit-page p-6">
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
          <h1 class="text-3xl font-bold text-gray-900">ویرایش اسلایدر خدمات</h1>
        </div>
        <p class="text-gray-600">ویرایش ابزارک اسلایدر خدمات</p>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="flex justify-center items-center py-12">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
      </div>

      <!-- Edit Form -->
      <div v-else-if="slider" class="bg-white rounded-lg shadow-lg p-6">
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
                  <option value="home">صفحه اصلی</option>
                  <option value="other">سایر بخش‌ها</option>
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
                  min="150"
                  max="500"
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
                    <input 
                      v-model="service.description"
                      type="text"
                      placeholder="توضیحات"
                      class="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                    />
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
              @click="resetForm"
              class="px-6 py-3 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors"
            >
              بازنشانی تغییرات
            </button>
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
              به‌روزرسانی ابزارک
            </button>
          </div>
        </form>
      </div>

      <!-- Error State -->
      <div v-else class="text-center py-12 bg-white rounded-lg shadow-sm">
        <div class="w-24 h-24 mx-auto mb-4 bg-red-100 rounded-full flex items-center justify-center">
          <svg class="w-12 h-12 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"/>
          </svg>
        </div>
        <h3 class="text-lg font-semibold text-gray-800 mb-2">خطا در بارگذاری</h3>
        <p class="text-gray-600 mb-4">اسلایدر خدمات مورد نظر یافت نشد</p>
        <NuxtLink 
          to="/admin/content/banners/ServicesSlider"
          class="bg-blue-600 hover:bg-blue-700 text-white px-6 py-3 rounded-lg font-semibold transition-colors duration-200 inline-flex items-center gap-2"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"/>
          </svg>
          بازگشت به لیست
        </NuxtLink>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';

// تعریف definePageMeta و useHead و navigateTo برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
declare const useHead: (head: { title?: string }) => void
declare const navigateTo: (to: string) => Promise<void>

// Page meta definition
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// Page title
useHead({
  title: 'ویرایش اسلایدر خدمات - پنل ادمین'
})

// Route
const route = useRoute()
const router = useRouter()
const sliderId = route.params.id as string

// State
const loading = ref(true)
const slider = ref(null)
const form = ref({
  title: '',
  code: '',
  page: '',
  order: 1,
  slidesPerView: 3,
  carouselHeight: 200,
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

// Mock data for demonstration
const mockSlider = {
  id: sliderId,
  title: 'اسلایدر خدمات ما',
  code: 'our-services',
  page: 'home',
  order: 3,
  slidesPerView: 3,
  carouselHeight: 200,
  slideGap: 20,
  autoplay: true,
  autoplayDelay: 5000,
  loop: true,
  showNavigation: true,
  showIndicators: true,
  serviceSelectionType: 'manual',
  services: [
    { title: 'پشتیبانی 24/7', icon: 'fas fa-headset', link: '/services/support', description: 'پشتیبانی شبانه‌روزی' },
    { title: 'ارسال رایگان', icon: 'fas fa-shipping-fast', link: '/services/free-shipping', description: 'ارسال رایگان برای سفارشات بالای 500 هزار تومان' }
  ]
}

// Methods
const loadSlider = async () => {
  try {
    loading.value = true
    // Here you would call your API to load the slider
    // For now, using mock data
    await new Promise(resolve => setTimeout(resolve, 1000))
    slider.value = mockSlider
    populateForm()
  } catch (error) {
    console.error('Error loading slider:', error)
  } finally {
    loading.value = false
  }
}

const populateForm = () => {
  if (!slider.value) return
  
  form.value = {
    title: slider.value.title,
    code: slider.value.code,
    page: slider.value.page,
    order: slider.value.order,
    slidesPerView: slider.value.slidesPerView,
    carouselHeight: slider.value.carouselHeight,
    slideGap: slider.value.slideGap,
    autoplay: slider.value.autoplay,
    autoplayDelay: slider.value.autoplayDelay,
    loop: slider.value.loop,
    showNavigation: slider.value.showNavigation,
    showIndicators: slider.value.showIndicators,
    serviceSelectionType: slider.value.serviceSelectionType,
    services: [...slider.value.services]
  }
}

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

    // Here you would call your API to update the slider
    console.log('Updating services slider widget:', form.value)
    
    // Redirect to management page
    await router.push('/admin/content/banners/ServicesSlider')
  } catch (error) {
    console.error('Error updating widget:', error)
    alert('خطا در به‌روزرسانی ابزارک')
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

const resetForm = () => {
  populateForm()
}

// Lifecycle
onMounted(() => {
  loadSlider()
})
</script>

<style scoped>
.services-slider-edit-page {
  background-color: #f9fafb;
  min-height: 100vh;
}
</style>

