<template>
  <div v-if="isOpen" class="fixed inset-0 z-50 overflow-y-auto">
    <div class="flex items-center justify-center min-h-screen px-4 pt-4 pb-20 text-center sm:block sm:p-0">
      <!-- Overlay -->
      <div class="fixed inset-0 transition-opacity bg-gray-500 bg-opacity-75" @click="$emit('close')"></div>

      <!-- Modal -->
      <div class="inline-block w-full max-w-4xl px-6 py-6 my-8 overflow-hidden text-right transition-all transform bg-white rounded-lg shadow-xl align-middle">
        <!-- Header -->
        <div class="flex items-center justify-between mb-6">
          <h3 class="text-lg font-medium text-gray-900">
            {{ editingTest ? 'ویرایش تست' : 'ایجاد تست جدید' }}
          </h3>
          <button @click="$emit('close')" class="text-gray-400 hover:text-gray-600">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        <!-- Form -->
        <form @submit.prevent="handleSubmit" class="space-y-6">
          <!-- اطلاعات پایه -->
          <div class="bg-gray-50 rounded-lg p-6">
            <h4 class="text-md font-semibold text-gray-900 mb-4">اطلاعات پایه</h4>
            
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <!-- نام تست -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">نام تست *</label>
                <input
                  v-model="form.name"
                  type="text"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  placeholder="نام تست را وارد کنید"
                />
              </div>

              <!-- نوع تست -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">نوع تست *</label>
                <select
                  v-model="form.type"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                >
                  <option value="">انتخاب کنید</option>
                  <option value="page">صفحه</option>
                  <option value="button">دکمه</option>
                  <option value="price">قیمت</option>
                  <option value="image">تصویر</option>
                  <option value="product">محصول</option>
                </select>
              </div>

              <!-- تاریخ شروع -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">تاریخ شروع *</label>
                <input
                  v-model="form.startDate"
                  type="date"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                />
              </div>

              <!-- تاریخ پایان -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">تاریخ پایان *</label>
                <input
                  v-model="form.endDate"
                  type="date"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                />
              </div>

              <!-- درصد تقسیم ترافیک -->
              <div class="md:col-span-2">
                <label class="block text-sm font-medium text-gray-700 mb-1">درصد تقسیم ترافیک</label>
                <div class="flex items-center space-x-4 space-x-reverse">
                  <div class="flex-1">
                    <label class="block text-xs text-gray-500 mb-1">نسخه A</label>
                    <input
                      v-model.number="form.trafficSplitA"
                      type="number"
                      min="0"
                      max="100"
                      class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                    />
                  </div>
                  <div class="flex-1">
                    <label class="block text-xs text-gray-500 mb-1">نسخه B</label>
                    <input
                      v-model.number="form.trafficSplitB"
                      type="number"
                      min="0"
                      max="100"
                      class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                    />
                  </div>
                </div>
                <p class="text-xs text-gray-500 mt-1">مجموع باید 100 باشد</p>
              </div>
            </div>

            <!-- توضیحات -->
            <div class="mt-4">
              <label class="block text-sm font-medium text-gray-700 mb-1">توضیحات</label>
              <textarea
                v-model="form.description"
                rows="3"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="توضیحات تست را وارد کنید"
              ></textarea>
            </div>
          </div>

          <!-- تنظیمات پیشرفته -->
          <div class="bg-gray-50 rounded-lg p-6">
            <h4 class="text-md font-semibold text-gray-900 mb-4">تنظیمات پیشرفته</h4>
            
            <!-- معیارهای موفقیت -->
            <div class="mb-4">
              <label class="block text-sm font-medium text-gray-700 mb-2">معیارهای موفقیت</label>
              <div class="space-y-2">
                <label class="flex items-center">
                  <input
                    v-model="form.successMetrics"
                    type="checkbox"
                    value="conversion_rate"
                    class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                  />
                  <span class="mr-2 text-sm text-gray-700">نرخ تبدیل (Conversion Rate)</span>
                </label>
                <label class="flex items-center">
                  <input
                    v-model="form.successMetrics"
                    type="checkbox"
                    value="revenue"
                    class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                  />
                  <span class="mr-2 text-sm text-gray-700">درآمد (Revenue)</span>
                </label>
                <label class="flex items-center">
                  <input
                    v-model="form.successMetrics"
                    type="checkbox"
                    value="click_rate"
                    class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                  />
                  <span class="mr-2 text-sm text-gray-700">نرخ کلیک (Click Rate)</span>
                </label>
              </div>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
              <!-- حداقل نمونه‌گیری -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">حداقل نمونه‌گیری</label>
                <input
                  v-model.number="form.minSampleSize"
                  type="number"
                  min="100"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  placeholder="1000"
                />
              </div>

              <!-- سطح اطمینان -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">سطح اطمینان (%)</label>
                <select
                  v-model="form.confidenceLevel"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                >
                  <option value="90">90%</option>
                  <option value="95">95%</option>
                  <option value="99">99%</option>
                </select>
              </div>

              <!-- قدرت تست -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">قدرت تست (%)</label>
                <select
                  v-model="form.testPower"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                >
                  <option value="80">80%</option>
                  <option value="85">85%</option>
                  <option value="90">90%</option>
                </select>
              </div>
            </div>
          </div>

          <!-- فیلترهای کاربران -->
          <div class="bg-gray-50 rounded-lg p-6">
            <h4 class="text-md font-semibold text-gray-900 mb-4">فیلترهای کاربران</h4>
            
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <!-- محدوده سنی -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">محدوده سنی</label>
                <div class="flex items-center space-x-2 space-x-reverse">
                  <input
                    v-model.number="form.ageRange.min"
                    type="number"
                    min="0"
                    max="100"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                    placeholder="18"
                  />
                  <span class="text-gray-500">تا</span>
                  <input
                    v-model.number="form.ageRange.max"
                    type="number"
                    min="0"
                    max="100"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                    placeholder="65"
                  />
                </div>
              </div>

              <!-- جنسیت -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">جنسیت</label>
                <select
                  v-model="form.gender"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                >
                  <option value="">همه</option>
                  <option value="male">مرد</option>
                  <option value="female">زن</option>
                </select>
              </div>

              <!-- موقعیت جغرافیایی -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">موقعیت جغرافیایی</label>
                <select
                  v-model="form.location"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                >
                  <option value="">همه مناطق</option>
                  <option value="tehran">تهران</option>
                  <option value="isfahan">اصفهان</option>
                  <option value="mashhad">مشهد</option>
                  <option value="tabriz">تبریز</option>
                  <option value="other">سایر شهرها</option>
                </select>
              </div>

              <!-- نوع دستگاه -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">نوع دستگاه</label>
                <select
                  v-model="form.deviceType"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                >
                  <option value="">همه دستگاه‌ها</option>
                  <option value="mobile">موبایل</option>
                  <option value="desktop">دسکتاپ</option>
                  <option value="tablet">تبلت</option>
                </select>
              </div>
            </div>
          </div>

          <!-- دکمه‌های عملیات -->
          <div class="flex items-center justify-end space-x-4 space-x-reverse pt-4 border-t">
            <button
              type="button"
              @click="$emit('close')"
              class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-lg hover:bg-gray-50"
            >
              انصراف
            </button>
            <button
              type="submit"
              class="px-4 py-2 text-sm font-medium text-white bg-blue-600 border border-transparent rounded-lg hover:bg-blue-700"
            >
              {{ editingTest ? 'ذخیره تغییرات' : 'ایجاد تست' }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
// Props
const props = defineProps<{
  isOpen: boolean
  editingTest?: any
}>()

// Events
const emit = defineEmits(['close', 'submit'])

// فرم داده‌ها
const form = ref({
  name: '',
  description: '',
  type: '',
  startDate: '',
  endDate: '',
  trafficSplitA: 50,
  trafficSplitB: 50,
  successMetrics: ['conversion_rate'],
  minSampleSize: 1000,
  confidenceLevel: 95,
  testPower: 80,
  ageRange: {
    min: null,
    max: null
  },
  gender: '',
  location: '',
  deviceType: ''
})

// بارگذاری داده‌های ویرایش
watch(() => props.editingTest, (test) => {
  if (test) {
    form.value = {
      name: test.name || '',
      description: test.description || '',
      type: test.type || '',
      startDate: test.startDate || '',
      endDate: test.endDate || '',
      trafficSplitA: test.trafficSplitA || 50,
      trafficSplitB: test.trafficSplitB || 50,
      successMetrics: test.successMetrics || ['conversion_rate'],
      minSampleSize: test.minSampleSize || 1000,
      confidenceLevel: test.confidenceLevel || 95,
      testPower: test.testPower || 80,
      ageRange: {
        min: test.ageRange?.min || null,
        max: test.ageRange?.max || null
      },
      gender: test.gender || '',
      location: test.location || '',
      deviceType: test.deviceType || ''
    }
  } else {
    // ریست فرم برای تست جدید
    form.value = {
      name: '',
      description: '',
      type: '',
      startDate: '',
      endDate: '',
      trafficSplitA: 50,
      trafficSplitB: 50,
      successMetrics: ['conversion_rate'],
      minSampleSize: 1000,
      confidenceLevel: 95,
      testPower: 80,
      ageRange: {
        min: null,
        max: null
      },
      gender: '',
      location: '',
      deviceType: ''
    }
  }
}, { immediate: true })

// اعتبارسنجی و ارسال فرم
const handleSubmit = () => {
  // اعتبارسنجی درصد تقسیم ترافیک
  if (form.value.trafficSplitA + form.value.trafficSplitB !== 100) {
    alert('مجموع درصد تقسیم ترافیک باید 100 باشد')
    return
  }

  // اعتبارسنجی تاریخ
  if (new Date(form.value.startDate) >= new Date(form.value.endDate)) {
    alert('تاریخ پایان باید بعد از تاریخ شروع باشد')
    return
  }

  emit('submit', form.value)
}

// تنظیم خودکار درصد تقسیم ترافیک
watch(() => form.value.trafficSplitA, (newVal) => {
  form.value.trafficSplitB = 100 - newVal
})

watch(() => form.value.trafficSplitB, (newVal) => {
  form.value.trafficSplitA = 100 - newVal
})
</script> 
