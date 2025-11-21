<template>
  <div class="brands-slider-create-page p-6">
    <div class="max-w-4xl mx-auto">
      <!-- Page Header -->
      <div class="mb-8">
        <div class="flex items-center gap-6 mb-4">
          <NuxtLink 
            to="/admin/content/banners/BrandsSlider"
            class="text-blue-600 hover:text-blue-800 transition-colors"
          >
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"/>
            </svg>
          </NuxtLink>
          <h1 class="text-3xl font-bold text-gray-900">ایجاد اسلایدر برندها</h1>
        </div>
        <p class="text-gray-600">ایجاد ابزارک جدید برای نمایش برندها در قالب اسلایدر</p>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="flex justify-center items-center py-12">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
        <span class="mr-4 text-gray-600">در حال پردازش...</span>
      </div>

      <!-- Error State -->
      <div v-else-if="error" class="bg-red-50 border border-red-200 rounded-lg p-6 mb-6">
        <p class="text-red-800">{{ error }}</p>
        <button
          class="mt-2 text-red-600 hover:text-red-800"
          @click="error = null"
        >
          بستن
        </button>
      </div>

      <!-- Create Form -->
      <div v-else class="bg-white rounded-lg shadow-lg p-6">
        <form class="space-y-6" @submit.prevent="handleSubmit">
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
                  placeholder="مثال: اسلایدر برندهای محبوب"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">کد یکتا</label>
                <input 
                  v-model="form.code"
                  type="text"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  placeholder="مثال: brands-slider"
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
                  <option value="category">صفحات دسته‌بندی</option>
                  <option value="product">صفحات محصول</option>
                  <option value="brands">صفحات برندها</option>
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
                <label class="block text-sm font-medium text-gray-700 mb-2">تعداد برند در هر اسلاید</label>
                <input 
                  v-model.number="form.slidesPerView"
                  type="number"
                  min="1"
                  max="8"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">ارتفاع اسلایدر (پیکسل)</label>
                <input 
                  v-model.number="form.carouselHeight"
                  type="number"
                  min="100"
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

          <!-- Brand Selection -->
          <div class="border-b border-gray-200 pb-6">
            <h2 class="text-xl font-semibold text-gray-800 mb-4">انتخاب برندها</h2>
            <div class="mb-4">
              <label class="block text-sm font-medium text-gray-700 mb-2">روش انتخاب برندها</label>
              <select 
                v-model="form.brandSelectionType"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
                <option value="manual">انتخاب دستی</option>
                <option value="featured">برندهای ویژه</option>
                <option value="popular">برندهای محبوب</option>
                <option value="latest">آخرین برندها</option>
              </select>
            </div>

            <!-- Manual Brand Selection -->
            <div v-if="form.brandSelectionType === 'manual'" class="space-y-4">
              <div class="flex items-center justify-between">
                <h3 class="text-lg font-medium text-gray-800">برندهای انتخاب شده</h3>
                <button 
                  type="button"
                  class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors"
                  @click="addBrandItem"
                >
                  افزودن برند
                </button>
              </div>
              
              <div v-if="form.brands.length === 0" class="text-center py-8 text-gray-500">
                هیچ برندی انتخاب نشده است
              </div>
              
              <div v-else class="space-y-3">
                <div
                  v-for="(brand, index) in form.brands"
                  :key="index"
                  class="flex items-center gap-3 p-3 bg-gray-50 rounded-lg"
                >
                  <div class="flex-1 grid grid-cols-1 md:grid-cols-3 gap-3">
                    <input 
                      v-model="brand.name"
                      type="text"
                      placeholder="نام برند"
                      class="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                    />
                    <input 
                      v-model="brand.logo"
                      type="text"
                      placeholder="آدرس لوگو"
                      class="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                    />
                    <input 
                      v-model="brand.link"
                      type="text"
                      placeholder="لینک برند"
                      class="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                    />
                  </div>
                  <button 
                    type="button"
                    class="text-red-600 hover:text-red-800 p-2"
                    @click="removeBrandItem(index)"
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
              :disabled="loading"
              class="px-6 py-3 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors disabled:opacity-50"
              @click="saveAsDraft"
            >
              {{ loading ? 'در حال ذخیره...' : 'ذخیره به عنوان پیش‌نویس' }}
            </button>
            <button
              type="submit"
              :disabled="loading"
              class="px-6 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors disabled:opacity-50"
            >
              {{ loading ? 'در حال ایجاد...' : 'ایجاد ابزارک' }}
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

// تعریف definePageMeta، useHead و navigateTo برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
declare const useHead: (head: { title?: string }) => void
// navigateTo is not used in this component

const router = useRouter()
// Types
interface Brand {
  name: string
  logo: string
  link: string
}

interface BrandsSliderForm {
  title: string
  code: string
  page: string
  order: number
  slidesPerView: number
  carouselHeight: number
  slideGap: number
  autoplay: boolean
  autoplayDelay: number
  loop: boolean
  showNavigation: boolean
  showIndicators: boolean
  brandSelectionType: 'manual' | 'featured' | 'popular' | 'latest'
  brands: Brand[]
}

// Page meta definition
definePageMeta({
  layout: 'admin',
  middleware: 'admin'
})

// Page title
useHead({
  title: 'ایجاد اسلایدر برندها - پنل ادمین'
})

// State
const loading = ref(false)
const error = ref<string | null>(null)

// Form state
const form = ref<BrandsSliderForm>({
  title: '',
  code: '',
  page: '',
  order: 1,
  slidesPerView: 4,
  carouselHeight: 200,
  slideGap: 20,
  autoplay: true,
  autoplayDelay: 5000,
  loop: true,
  showNavigation: true,
  showIndicators: true,
  brandSelectionType: 'manual',
  brands: []
})

// Methods
const addBrandItem = () => {
  form.value.brands.push({
    name: '',
    logo: '',
    link: ''
  })
}

const removeBrandItem = (index: number) => {
  form.value.brands.splice(index, 1)
}

const handleSubmit = async () => {
  try {
    loading.value = true
    error.value = null

    // Validate form
    if (!form.value.title || !form.value.code || !form.value.page) {
      alert('لطفاً تمام فیلدهای اجباری را پر کنید')
      return
    }

    if (form.value.brandSelectionType === 'manual' && form.value.brands.length === 0) {
      alert('لطفاً حداقل یک برند انتخاب کنید')
      return
    }

    // TODO: جایگزین با API واقعی
    // const { data } = await $fetch('/api/admin/content/banners/brands-slider', {
    //   method: 'POST',
    //   headers: {
    //     'Content-Type': 'application/json'
    //   },
    //   body: form.value
    // })

    // Redirect to management page
    await router.push('/admin/content/banners/BrandsSlider')
  } catch (err) {
    console.error('Error creating widget:', err)
    error.value = err.message || 'خطا در ایجاد ابزارک'
    alert(error.value)
  } finally {
    loading.value = false
  }
}

const saveAsDraft = async () => {
  try {
    loading.value = true
    error.value = null

    // TODO: جایگزین با API واقعی
    // const { data } = await $fetch('/api/admin/content/banners/brands-slider/draft', {
    //   method: 'POST',
    //   headers: {
    //     'Content-Type': 'application/json'
    //   },
    //   body: { ...form.value, status: 'draft' }
    // })

    alert('پیش‌نویس ذخیره شد')
  } catch (err) {
    console.error('Error saving draft:', err)
    error.value = err.message || 'خطا در ذخیره پیش‌نویس'
    alert(error.value)
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.brands-slider-create-page {
  background-color: #f9fafb;
  min-height: 100vh;
}
</style>

