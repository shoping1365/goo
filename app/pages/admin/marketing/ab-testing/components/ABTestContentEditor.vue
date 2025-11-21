<template>
  <div v-if="isOpen" class="fixed inset-0 z-50 overflow-y-auto">
    <div class="flex items-center justify-center min-h-screen px-4 pt-4 pb-20 text-center sm:block sm:p-0">
      <!-- Overlay -->
      <div class="fixed inset-0 transition-opacity bg-gray-500 bg-opacity-75" @click="$emit('close')"></div>

      <!-- Modal -->
      <div class="inline-block w-full max-w-6xl px-6 py-6 my-8 overflow-hidden text-right transition-all transform bg-white rounded-lg shadow-xl align-middle">
        <!-- Header -->
        <div class="flex items-center justify-between mb-6">
          <h3 class="text-lg font-medium text-gray-900">ویرایشگر محتوا - {{ testName }}</h3>
          <div class="flex items-center space-x-4 space-x-reverse">
            <button class="px-3 py-1 text-sm bg-blue-100 text-blue-700 rounded-lg hover:bg-blue-200" @click="previewContent">
              پیش‌نمایش
            </button>
            <button class="text-gray-400 hover:text-gray-600" @click="$emit('close')">
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>
        </div>

        <!-- Content -->
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <!-- ویرایشگر -->
          <div class="space-y-4">
            <!-- انتخاب نوع محتوا -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نوع محتوا</label>
              <select v-model="contentType" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                <option value="page">صفحه</option>
                <option value="product">محصول</option>
                <option value="button">دکمه</option>
                <option value="image">تصویر</option>
                <option value="price">قیمت</option>
              </select>
            </div>

            <!-- ویرایشگر HTML/CSS برای صفحات -->
            <div v-if="contentType === 'page'" class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">نسخه A - HTML</label>
                <textarea
                  v-model="contentA.html"
                  rows="8"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 font-mono text-sm"
                  placeholder="کد HTML نسخه A را وارد کنید..."
                ></textarea>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">نسخه A - CSS</label>
                <textarea
                  v-model="contentA.css"
                  rows="6"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 font-mono text-sm"
                  placeholder="کد CSS نسخه A را وارد کنید..."
                ></textarea>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">نسخه B - HTML</label>
                <textarea
                  v-model="contentB.html"
                  rows="8"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 font-mono text-sm"
                  placeholder="کد HTML نسخه B را وارد کنید..."
                ></textarea>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">نسخه B - CSS</label>
                <textarea
                  v-model="contentB.css"
                  rows="6"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 font-mono text-sm"
                  placeholder="کد CSS نسخه B را وارد کنید..."
                ></textarea>
              </div>
            </div>

            <!-- ویرایشگر محصولات -->
            <div v-if="contentType === 'product'" class="space-y-4">
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">انتخاب محصول</label>
                  <select v-model="selectedProduct" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                    <option value="">انتخاب کنید</option>
                    <option v-for="product in products" :key="product.id" :value="product.id">
                      {{ product.name }}
                    </option>
                  </select>
                </div>
              </div>
              
              <!-- تغییرات قیمت -->
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">قیمت نسخه A</label>
                  <input
                    v-model.number="contentA.price"
                    type="number"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                    placeholder="قیمت به تومان"
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">قیمت نسخه B</label>
                  <input
                    v-model.number="contentB.price"
                    type="number"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                    placeholder="قیمت به تومان"
                  />
                </div>
              </div>

              <!-- تغییرات تصاویر -->
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">تصویر نسخه A</label>
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <input
                      v-model="contentA.image"
                      type="text"
                      class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                      placeholder="URL تصویر"
                    />
                    <button class="px-3 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200" @click="selectImage('A')">
                      انتخاب
                    </button>
                  </div>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">تصویر نسخه B</label>
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <input
                      v-model="contentB.image"
                      type="text"
                      class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                      placeholder="URL تصویر"
                    />
                    <button class="px-3 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200" @click="selectImage('B')">
                      انتخاب
                    </button>
                  </div>
                </div>
              </div>

              <!-- تغییرات توضیحات -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات نسخه A</label>
                <textarea
                  v-model="contentA.description"
                  rows="3"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  placeholder="توضیحات محصول نسخه A"
                ></textarea>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات نسخه B</label>
                <textarea
                  v-model="contentB.description"
                  rows="3"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  placeholder="توضیحات محصول نسخه B"
                ></textarea>
              </div>
            </div>

            <!-- ویرایشگر دکمه‌ها -->
            <div v-if="contentType === 'button'" class="space-y-4">
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">متن دکمه نسخه A</label>
                  <input
                    v-model="contentA.text"
                    type="text"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                    placeholder="متن دکمه"
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">متن دکمه نسخه B</label>
                  <input
                    v-model="contentB.text"
                    type="text"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                    placeholder="متن دکمه"
                  />
                </div>
              </div>

              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">رنگ دکمه نسخه A</label>
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <input
                      v-model="contentA.color"
                      type="color"
                      class="w-12 h-10 border border-gray-300 rounded-lg"
                    />
                    <input
                      v-model="contentA.color"
                      type="text"
                      class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                      placeholder="#3B82F6"
                    />
                  </div>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">رنگ دکمه نسخه B</label>
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <input
                      v-model="contentB.color"
                      type="color"
                      class="w-12 h-10 border border-gray-300 rounded-lg"
                    />
                    <input
                      v-model="contentB.color"
                      type="text"
                      class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                      placeholder="#10B981"
                    />
                  </div>
                </div>
              </div>

              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">اندازه دکمه نسخه A</label>
                  <select v-model="contentA.size" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                    <option value="small">کوچک</option>
                    <option value="medium">متوسط</option>
                    <option value="large">بزرگ</option>
                  </select>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">اندازه دکمه نسخه B</label>
                  <select v-model="contentB.size" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                    <option value="small">کوچک</option>
                    <option value="medium">متوسط</option>
                    <option value="large">بزرگ</option>
                  </select>
                </div>
              </div>
            </div>

            <!-- ویرایشگر تصاویر -->
            <div v-if="contentType === 'image'" class="space-y-4">
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">تصویر نسخه A</label>
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <input
                      v-model="contentA.imageUrl"
                      type="text"
                      class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                      placeholder="URL تصویر"
                    />
                    <button class="px-3 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200" @click="selectImage('A')">
                      انتخاب
                    </button>
                  </div>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">تصویر نسخه B</label>
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <input
                      v-model="contentB.imageUrl"
                      type="text"
                      class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                      placeholder="URL تصویر"
                    />
                    <button class="px-3 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200" @click="selectImage('B')">
                      انتخاب
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- پیش‌نمایش -->
          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <h4 class="text-md font-semibold text-gray-900">پیش‌نمایش زنده</h4>
              <div class="flex space-x-2 space-x-reverse">
                <button :class="previewVersion === 'A' ? 'bg-blue-600 text-white' : 'bg-gray-200 text-gray-700'" class="px-3 py-1 rounded-lg text-sm" @click="previewVersion = 'A'">
                  نسخه A
                </button>
                <button :class="previewVersion === 'B' ? 'bg-green-600 text-white' : 'bg-gray-200 text-gray-700'" class="px-3 py-1 rounded-lg text-sm" @click="previewVersion = 'B'">
                  نسخه B
                </button>
              </div>
            </div>

            <!-- پیش‌نمایش محتوا -->
            <div class="border border-gray-300 rounded-lg p-6 min-h-96 bg-gray-50">
              <!-- 
                ⚠️ امنیت XSS: استفاده از v-html خطرناک است!
                
                این کد محتوای HTML را بدون sanitization نمایش می‌دهد که می‌تواند منجر به حملات XSS شود.
                
                ✅ راه حل صحیح:
                1. قبل از استفاده از v-html، محتوا را با کتابخانه sanitization (مثل DOMPurify) پاکسازی کنید
                2. یا از {{ }} به جای v-html استفاده کنید اگر HTML نیاز نیست
                3. محتوای کاربر را هرگز بدون sanitization در v-html قرار ندهید
                
                مثال صحیح:
                import DOMPurify from 'dompurify'
                const sanitizedPreviewContent = computed(() => DOMPurify.sanitize(getPreviewContent()))
                <div v-html="sanitizedPreviewContent"></div>
              -->
              <!-- eslint-disable-next-line vue/no-v-html -->
              <div v-if="contentType === 'page'" v-html="getPreviewContent()"></div>
              
              <div v-else-if="contentType === 'product'" class="space-y-4">
                <div class="bg-white rounded-lg p-6 shadow-sm">
                  <img v-if="getCurrentContent().image" :src="getCurrentContent().image" alt="تصویر محصول" class="w-full h-48 object-cover rounded-lg mb-4" />
                  <h3 class="text-lg font-semibold mb-2">نام محصول</h3>
                  <p class="text-gray-600 mb-2">{{ getCurrentContent().description || 'توضیحات محصول' }}</p>
                  <div class="text-xl font-bold text-blue-600">{{ formatPrice(getCurrentContent().price) }}</div>
                </div>
              </div>

              <div v-else-if="contentType === 'button'" class="flex items-center justify-center">
                <button :style="{ backgroundColor: getCurrentContent().color, fontSize: getButtonSize() }" class="px-6 py-3 text-white rounded-lg">
                  {{ getCurrentContent().text || 'دکمه نمونه' }}
                </button>
              </div>

              <div v-else-if="contentType === 'image'" class="flex items-center justify-center">
                <img v-if="getCurrentContent().imageUrl" :src="getCurrentContent().imageUrl" alt="تصویر" class="max-w-full max-h-80 object-contain" />
                <div v-else class="text-gray-400 text-center">
                  <svg class="w-16 h-16 mx-auto mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                  </svg>
                  <p>تصویری انتخاب نشده</p>
                </div>
              </div>

              <div v-else class="text-gray-400 text-center py-12">
                <svg class="w-16 h-16 mx-auto mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                </svg>
                <p>نوع محتوا را انتخاب کنید</p>
              </div>
            </div>
          </div>
        </div>

        <!-- دکمه‌های عملیات -->
        <div class="flex items-center justify-end space-x-4 space-x-reverse pt-6 border-t mt-6">
          <button
            class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-lg hover:bg-gray-50"
            @click="$emit('close')"
          >
            انصراف
          </button>
          <button
            class="px-4 py-2 text-sm font-medium text-white bg-blue-600 border border-transparent rounded-lg hover:bg-blue-700"
            @click="saveContent"
          >
            ذخیره محتوا
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
// Props
const _props = defineProps<{
  isOpen: boolean
  testName?: string
}>()

// Events
const emit = defineEmits(['close', 'save'])

// State
const contentType = ref('page')
const previewVersion = ref('A')
const selectedProduct = ref('')

// محتوای نسخه A و B
const contentA = ref({
  html: '',
  css: '',
  text: '',
  color: '#3B82F6',
  size: 'medium',
  image: '',
  imageUrl: '',
  price: 0,
  description: ''
})

const contentB = ref({
  html: '',
  css: '',
  text: '',
  color: '#10B981',
  size: 'medium',
  image: '',
  imageUrl: '',
  price: 0,
  description: ''
})

// لیست محصولات (mock data)
const products = ref([
  { id: 1, name: 'لپ‌تاپ اپل MacBook Pro' },
  { id: 2, name: 'گوشی سامسونگ Galaxy S24' },
  { id: 3, name: 'ساعت هوشمند اپل Watch' },
  { id: 4, name: 'هدفون Sony WH-1000XM5' }
])

// دریافت محتوای فعلی برای پیش‌نمایش
const getCurrentContent = () => {
  return previewVersion.value === 'A' ? contentA.value : contentB.value
}

// دریافت محتوای پیش‌نمایش برای صفحات
const getPreviewContent = () => {
  const content = getCurrentContent()
  return `<style>${content.css}</style>${content.html}`
}

// اندازه دکمه
const getButtonSize = () => {
  const sizes = {
    small: '14px',
    medium: '16px',
    large: '18px'
  }
  return sizes[getCurrentContent().size as keyof typeof sizes] || '16px'
}

// فرمت کردن قیمت
const formatPrice = (price: number) => {
  if (!price) return 'قیمت تعیین نشده'
  return new Intl.NumberFormat('fa-IR', {
    style: 'currency',
    currency: 'IRR',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(price)
}

// انتخاب تصویر
const selectImage = (version: string) => {
  // در حالت واقعی، modal انتخاب تصویر باز می‌شود
  const url = prompt('URL تصویر را وارد کنید:')
  if (url) {
    if (version === 'A') {
      contentA.value.image = url
      contentA.value.imageUrl = url
    } else {
      contentB.value.image = url
      contentB.value.imageUrl = url
    }
  }
}

// پیش‌نمایش محتوا
const previewContent = () => {
  // منطق پیش‌نمایش در پنجره جدید
  const content = getPreviewContent()
  const newWindow = window.open('', '_blank')
  if (newWindow) {
    newWindow.document.write(content)
    newWindow.document.close()
  }
}

// ذخیره محتوا
const saveContent = () => {
  const content = {
    type: contentType.value,
    versionA: contentA.value,
    versionB: contentB.value
  }
  emit('save', content)
  emit('close')
}
</script> 
