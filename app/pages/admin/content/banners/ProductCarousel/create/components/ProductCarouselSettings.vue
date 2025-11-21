<template>
  <div class="widget-settings-panel bg-white rounded-lg shadow mb-6">
    <div class="px-6 py-4 border-b border-gray-200">
      <h2 class="text-lg font-medium text-gray-900">تنظیمات ویجت</h2>
    </div>
    <div class="p-6">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <!-- اطلاعات پایه -->
        <div class="space-y-4">
          <h3 class="text-sm font-medium text-gray-900">اطلاعات پایه</h3>
          <div>
            <label class="block text-sm font-medium text-gray-900 mb-2">عنوان ویجت</label>
            <input
              v-model="localFormData.title"
              type="text"
              placeholder="عنوان ویجت را وارد کنید"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 text-gray-900 placeholder-gray-600 font-medium"
              required
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-900 mb-2">توضیحات</label>
            <textarea
              v-model="localFormData.description"
              rows="3"
              placeholder="توضیحات ویجت (اختیاری)"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 text-gray-900 placeholder-gray-600 font-medium"
            ></textarea>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-900 mb-2">صفحه نمایش</label>
            <select
              v-model="localFormData.page"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 text-gray-900 placeholder-gray-600 font-medium"
            >
              <option value="home">صفحه اصلی</option>
              <option value="category">صفحه دسته‌بندی</option>
              <option value="product">صفحه محصول</option>
              <option value="other">سایر صفحات</option>
            </select>
          </div>
        </div>

        <!-- تنظیمات محتوا -->
        <div class="space-y-4">
          <h3 class="text-sm font-medium text-gray-900">تنظیمات محتوا</h3>
          <div>
            <label class="block text-sm font-medium text-gray-900 mb-2">عنوان کاروسل</label>
            <input
              v-model="localConfig.title"
              type="text"
              placeholder="عنوان کاروسل را وارد کنید"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 text-gray-900 placeholder-gray-600 font-medium"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-900 mb-2">زیرعنوان</label>
            <input
              v-model="localConfig.subtitle"
              type="text"
              placeholder="زیرعنوان کاروسل (اختیاری)"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 text-gray-900 placeholder-gray-600 font-medium"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-900 mb-2">دسته‌بندی محصولات</label>
            <CategorySelector
              v-model="localConfig.categoryId"
              :categories="props.categories"
              placeholder="همه محصولات"
            />
            <div class="mt-2 text-xs text-gray-700">
              تعداد دسته‌بندی‌ها: {{ props.categories?.length || 0 }}
              <span v-if="props.categories?.length > 0" class="text-green-600">✅</span>
              <span v-else class="text-red-600">❌</span>
            </div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-900 mb-2">تعداد محصولات</label>
            <input
              v-model.number="localConfig.productCount"
              type="number"
              min="1"
              max="50"
              placeholder="12"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 text-gray-900 placeholder-gray-600 font-medium"
            />
          </div>
        </div>

        <!-- تنظیمات نمایش -->
        <div class="space-y-4">
          <h3 class="text-sm font-medium text-gray-900">تنظیمات نمایش</h3>
          <div>
            <label class="block text-sm font-medium text-gray-900 mb-2">تعداد اسلاید در هر صفحه</label>
            <select
              v-model.number="localConfig.slidesPerView"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 text-gray-900 placeholder-gray-600 font-medium"
            >
              <option :value="1">1 اسلاید</option>
              <option :value="2">2 اسلاید</option>
              <option :value="3">3 اسلاید</option>
              <option :value="4">4 اسلاید</option>
              <option :value="5">5 اسلاید</option>
              <option :value="6">6 اسلاید</option>
            </select>
          </div>

          <div class="space-y-3">
            <div class="flex items-center">
              <input
                v-model="localConfig.showPrice"
                type="checkbox"
                class="h-4 w-4 text-purple-600 focus:ring-purple-500 border-gray-300 rounded"
              />
              <label class="mr-2 text-sm text-gray-900">نمایش قیمت</label>
            </div>
            <div class="flex items-center">
              <input
                v-model="localConfig.showRating"
                type="checkbox"
                class="h-4 w-4 text-purple-600 focus:ring-purple-500 border-gray-300 rounded"
              />
              <label class="mr-2 text-sm text-gray-900">نمایش امتیاز</label>
            </div>
            <div class="flex items-center">
              <input
                v-model="localConfig.showDiscount"
                type="checkbox"
                class="h-4 w-4 text-purple-600 focus:ring-purple-500 border-gray-300 rounded"
              />
              <label class="mr-2 text-sm text-gray-900">نمایش تخفیف</label>
            </div>
          </div>
        </div>

        <!-- تنظیمات کاروسل و رنگ -->
        <div class="space-y-4">
          <h3 class="text-sm font-medium text-gray-900">تنظیمات کاروسل</h3>

          <!-- تنظیمات نمایش -->
          <div class="space-y-3">
            <div class="flex items-center">
              <input
                v-model="localConfig.showNavigation"
                type="checkbox"
                class="h-4 w-4 text-purple-600 focus:ring-purple-500 border-gray-300 rounded"
              />
              <label class="mr-2 text-sm text-gray-900">نمایش ناوبری</label>
            </div>
            <div class="flex items-center">
              <input
                v-model="localConfig.showIndicators"
                type="checkbox"
                class="h-4 w-4 text-purple-600 focus:ring-purple-500 border-gray-300 rounded"
              />
              <label class="mr-2 text-sm text-gray-900">نمایش نشانگرها</label>
            </div>
            <div class="flex items-center">
              <input
                v-model="localConfig.showControls"
                type="checkbox"
                class="h-4 w-4 text-purple-600 focus:ring-purple-500 border-gray-300 rounded"
              />
              <label class="mr-2 text-sm text-gray-900">نمایش کنترل‌ها</label>
            </div>
          </div>

          <div class="flex items-center">
            <input
              v-model="localConfig.autoplay"
              type="checkbox"
              class="h-4 w-4 text-purple-600 focus:ring-purple-500 border-gray-300 rounded"
            />
            <label class="mr-2 text-sm text-gray-900">پخش خودکار</label>
          </div>

          <div v-if="localConfig.autoplay">
            <label class="block text-sm font-medium text-gray-900 mb-2">تاخیر پخش خودکار (میلی‌ثانیه)</label>
            <input
              v-model.number="localConfig.autoplayDelay"
              type="number"
              min="1000"
              max="10000"
              step="500"
              placeholder="3000"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 text-gray-900 placeholder-gray-600 font-medium"
            />
          </div>

          <div class="flex items-center">
            <input
              v-model="localConfig.wide_bg"
              type="checkbox"
              class="h-4 w-4 text-purple-600 focus:ring-purple-500 border-gray-300 rounded"
            />
            <label class="mr-2 text-sm text-gray-900">پس‌زمینه تمام عرض</label>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-900 mb-2">رنگ پس‌زمینه</label>
            <input
              v-model="localConfig.bg_color"
              type="color"
              class="w-full h-10 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-900 mb-2">فاصله داخلی</label>
            <input
              v-model="localConfig.padding"
              type="text"
              placeholder="مثال: 24px"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 text-gray-900 placeholder-gray-600 font-medium"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-900 mb-2">فاصله خارجی</label>
            <input
              v-model="localConfig.margin"
              type="text"
              placeholder="مثال: 0px"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 text-gray-900 placeholder-gray-600 font-medium"
            />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import type { ProductCarouselConfig } from '~/types/widget'
import CategorySelector from './CategorySelector.vue'

interface Props {
  formData: {
    title: string
    description: string
    page: string
  }
  config: ProductCarouselConfig
  categories: any[]
}

const props = defineProps<Props>()

const emit = defineEmits<{
  (e: 'update:formData', value: Props['formData']): void
  (e: 'update:config', value: ProductCarouselConfig): void
}>()

// Proxy for formData
const localFormData = computed({
  get: () => new Proxy(props.formData, {
    set(obj, prop, value) {
      emit('update:formData', { ...obj, [prop]: value })
      return true
    }
  }),
  set: (val) => emit('update:formData', val)
})

// Proxy for config
const localConfig = computed({
  get: () => new Proxy(props.config, {
    set(obj, prop, value) {
      emit('update:config', { ...obj, [prop]: value })
      return true
    }
  }),
  set: (val) => emit('update:config', val)
})

// State برای dropdown دسته‌بندی (در صورت نیاز)
const showCategoryDropdown = ref(false)
const categorySearch = ref('')

// متدهای مدیریت dropdown (در صورت نیاز)
const closeCategoryDropdown = () => {
  showCategoryDropdown.value = false
  categorySearch.value = ''
}
</script>

<style scoped>
/* بهبود وضوح متن در فیلدهای ورودی */
input, textarea, select {
  color: #111827 !important; /* text-gray-900 */
  font-weight: 500 !important;
}

input::placeholder, textarea::placeholder {
  color: #4b5563 !important; /* text-gray-600 */
  font-weight: 400 !important;
}

/* بهبود وضوح متن در select options */
select option {
  color: #111827 !important;
  font-weight: 500 !important;
}

/* بهبود وضوح متن در فیلدهای focus */
input:focus, textarea:focus, select:focus {
  color: #111827 !important;
  font-weight: 500 !important;
}
</style>

<style scoped>
/* استایل‌های لازم برای ProductCarouselSettings */
</style>
