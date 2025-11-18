<template>
  <div class="space-y-8">
    <!-- Header -->
    <div>
      <h3 class="text-lg font-semibold text-gray-800">تمپلیت‌های نظرسنجی</h3>
      <p class="text-gray-600 text-sm">انتخاب و شخصی‌سازی قالب‌های نظرسنجی برای مشتریان</p>
    </div>

    <!-- Template Selection -->
    <div class="bg-white rounded-lg border border-gray-200 px-4 py-4">
      <h4 class="text-md font-semibold text-gray-800 mb-4">انتخاب تمپلیت فعال</h4>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gapx-4 py-4">
        <div 
          v-for="template in templates" 
          :key="template.id"
          @click="selectTemplate(template.id)"
          :class="[
            'border-2 rounded-lg px-4 py-4 cursor-pointer transition-all hover:shadow-lg',
            selectedTemplate === template.id 
              ? 'border-blue-500 bg-blue-50 shadow-lg' 
              : 'border-gray-200 hover:border-gray-300'
          ]"
        >
          <div class="flex items-center justify-between mb-3">
            <h5 class="font-medium text-gray-900">{{ template.name }}</h5>
            <div v-if="selectedTemplate === template.id" class="w-6 h-6 bg-blue-500 rounded-full flex items-center justify-center">
              <svg class="w-4 h-4 text-white" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd"></path>
              </svg>
            </div>
          </div>
          <div class="text-sm text-gray-600 mb-3">{{ template.description }}</div>
          <div class="flex items-center justify-between text-xs text-gray-500">
            <span>{{ template.category }}</span>
            <span>{{ template.rating }}⭐</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Template Preview -->
    <div class="bg-white rounded-lg border border-gray-200 px-4 py-4">
      <h4 class="text-md font-semibold text-gray-800 mb-4">پیش‌نمایش تمپلیت انتخاب شده</h4>
      <div class="bg-gray-50 rounded-lg px-4 py-4">
        <div class="max-w-md mx-auto">
          <component :is="selectedTemplateComponent" />
        </div>
      </div>
    </div>

    <!-- Template Customization -->
    <div class="bg-white rounded-lg border border-gray-200 px-4 py-4">
      <h4 class="text-md font-semibold text-gray-800 mb-4">شخصی‌سازی تمپلیت</h4>
      <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">رنگ اصلی</label>
          <input 
            v-model="customization.primaryColor" 
            type="color" 
            class="w-full h-10 rounded-lg border border-gray-300"
          >
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">رنگ ثانویه</label>
          <input 
            v-model="customization.secondaryColor" 
            type="color" 
            class="w-full h-10 rounded-lg border border-gray-300"
          >
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">اندازه فونت</label>
          <select v-model="customization.fontSize" class="w-full px-3 py-2 border border-gray-300 rounded-lg">
            <option value="small">کوچک</option>
            <option value="medium">متوسط</option>
            <option value="large">بزرگ</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">سبک دکمه‌ها</label>
          <select v-model="customization.buttonStyle" class="w-full px-3 py-2 border border-gray-300 rounded-lg">
            <option value="rounded">گرد</option>
            <option value="square">مربعی</option>
            <option value="pill">کپسولی</option>
          </select>
        </div>
      </div>
    </div>

    <!-- Save Button -->
    <div class="flex items-center justify-end">
      <button 
        @click="saveTemplate"
        :disabled="saving"
        class="px-6 py-2 bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 text-white rounded-lg text-sm transition-colors flex items-center space-x-2 space-x-reverse"
      >
        <svg v-if="saving" class="animate-spin w-4 h-4" fill="none" viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
        </svg>
        <span>{{ saving ? 'در حال ذخیره...' : 'ذخیره تنظیمات' }}</span>
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

const selectedTemplate = ref('modern')
const saving = ref(false)

const customization = ref({
  primaryColor: '#3B82F6',
  secondaryColor: '#1E40AF',
  fontSize: 'medium',
  buttonStyle: 'rounded'
})

const templates = ref([
  {
    id: 'modern',
    name: 'تمپلیت مدرن',
    description: 'طراحی مدرن با رنگ‌های آبی و انیمیشن‌های نرم',
    category: 'مدرن',
    rating: 4.8
  },
  {
    id: 'elegant',
    name: 'تمپلیت شیک',
    description: 'طراحی شیک و مینیمال با رنگ‌های طلایی',
    category: 'شیک',
    rating: 4.9
  },
  {
    id: 'colorful',
    name: 'تمپلیت رنگی',
    description: 'طراحی پرانرژی با رنگ‌های متنوع و جذاب',
    category: 'رنگی',
    rating: 4.6
  },
  {
    id: 'minimal',
    name: 'تمپلیت مینیمال',
    description: 'طراحی ساده و تمیز با تمرکز بر محتوا',
    category: 'مینیمال',
    rating: 4.7
  },
  {
    id: 'gradient',
    name: 'تمپلیت گرادیانت',
    description: 'طراحی مدرن با پس‌زمینه گرادیانت زیبا',
    category: 'گرادیانت',
    rating: 4.8
  }
])

// Template Components
const ModernTemplate = {
  template: `
    <div class="bg-white rounded-xl shadow-lg overflow-hidden">
      <div class="bg-gradient-to-r from-blue-500 to-blue-600 px-4 py-4 text-white text-center">
        <h2 class="text-2xl font-bold mb-2">نظرسنجی خرید</h2>
        <p class="text-blue-100">تجربه شما برای ما مهم است</p>
      </div>
      <div class="px-4 py-4">
        <div class="mb-6">
          <h3 class="text-lg font-semibold text-gray-800 mb-4">چقدر از خرید خود راضی هستید؟</h3>
          <div class="flex justify-center space-x-2 space-x-reverse">
            <button class="w-12 h-12 rounded-full bg-gray-100 hover:bg-blue-100 text-gray-600 hover:text-blue-600 transition-colors">1</button>
            <button class="w-12 h-12 rounded-full bg-gray-100 hover:bg-blue-100 text-gray-600 hover:text-blue-600 transition-colors">2</button>
            <button class="w-12 h-12 rounded-full bg-gray-100 hover:bg-blue-100 text-gray-600 hover:text-blue-600 transition-colors">3</button>
            <button class="w-12 h-12 rounded-full bg-blue-500 text-white">4</button>
            <button class="w-12 h-12 rounded-full bg-gray-100 hover:bg-blue-100 text-gray-600 hover:text-blue-600 transition-colors">5</button>
          </div>
        </div>
        <div class="mb-6">
          <label class="block text-sm font-medium text-gray-700 mb-2">نظر شما:</label>
          <textarea class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500" rows="3" placeholder="تجربه خود را با ما به اشتراک بگذارید..."></textarea>
        </div>
        <button class="w-full bg-blue-500 hover:bg-blue-600 text-white font-medium py-3 px-4 rounded-lg transition-colors">
          ارسال نظر
        </button>
      </div>
    </div>
  `
}

const ElegantTemplate = {
  template: `
    <div class="bg-white rounded-xl shadow-lg overflow-hidden border border-gray-100">
      <div class="bg-gradient-to-r from-amber-400 to-yellow-500 px-4 py-4 text-white text-center">
        <div class="w-16 h-16 bg-white bg-opacity-20 rounded-full mx-auto mb-4 flex items-center justify-center">
          <svg class="w-8 h-8" fill="currentColor" viewBox="0 0 20 20">
            <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"></path>
          </svg>
        </div>
        <h2 class="text-2xl font-bold mb-2">نظرسنجی کیفیت</h2>
        <p class="text-yellow-100">ارزش‌گذاری شما برای بهبود خدمات</p>
      </div>
      <div class="px-4 py-4">
        <div class="mb-6">
          <h3 class="text-lg font-semibold text-gray-800 mb-4 text-center">امتیاز کلی محصول</h3>
          <div class="flex justify-center space-x-1 space-x-reverse">
            <svg class="w-8 h-8 text-yellow-400" fill="currentColor" viewBox="0 0 20 20"><path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"></path></svg>
            <svg class="w-8 h-8 text-yellow-400" fill="currentColor" viewBox="0 0 20 20"><path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"></path></svg>
            <svg class="w-8 h-8 text-yellow-400" fill="currentColor" viewBox="0 0 20 20"><path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"></path></svg>
            <svg class="w-8 h-8 text-yellow-400" fill="currentColor" viewBox="0 0 20 20"><path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"></path></svg>
            <svg class="w-8 h-8 text-gray-300" fill="currentColor" viewBox="0 0 20 20"><path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"></path></svg>
          </div>
        </div>
        <div class="mb-6">
          <label class="block text-sm font-medium text-gray-700 mb-2">نظر شما:</label>
          <textarea class="w-full px-3 py-2 border border-gray-200 rounded-lg focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500" rows="3" placeholder="تجربه خود را با ما به اشتراک بگذارید..."></textarea>
        </div>
        <button class="w-full bg-gradient-to-r from-amber-400 to-yellow-500 hover:from-amber-500 hover:to-yellow-600 text-white font-medium py-3 px-4 rounded-lg transition-all transform hover:scale-105">
          ارسال نظر
        </button>
      </div>
    </div>
  `
}

const ColorfulTemplate = {
  template: `
    <div class="bg-gradient-to-br from-purple-400 via-pink-500 to-red-500 rounded-xl shadow-lg overflow-hidden">
      <div class="bg-white bg-opacity-10 backdrop-blur-sm px-4 py-4 text-white text-center">
        <div class="w-20 h-20 bg-white bg-opacity-20 rounded-full mx-auto mb-4 flex items-center justify-center">
          <svg class="w-10 h-10" fill="currentColor" viewBox="0 0 20 20">
            <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clip-rule="evenodd"></path>
          </svg>
        </div>
        <h2 class="text-3xl font-bold mb-2">نظرسنجی</h2>
        <p class="text-purple-100 text-lg">تجربه شما مهم است!</p>
      </div>
      <div class="px-4 py-4 bg-white bg-opacity-95 backdrop-blur-sm">
        <div class="mb-6">
          <h3 class="text-xl font-bold text-gray-800 mb-4 text-center">چقدر راضی هستید؟</h3>
          <div class="grid grid-cols-5 gap-2">
            <button class="w-12 h-12 rounded-full bg-red-400 hover:bg-red-500 text-white font-bold transition-all transform hover:scale-110">1</button>
            <button class="w-12 h-12 rounded-full bg-orange-400 hover:bg-orange-500 text-white font-bold transition-all transform hover:scale-110">2</button>
            <button class="w-12 h-12 rounded-full bg-yellow-400 hover:bg-yellow-500 text-white font-bold transition-all transform hover:scale-110">3</button>
            <button class="w-12 h-12 rounded-full bg-green-400 hover:bg-green-500 text-white font-bold transition-all transform hover:scale-110">4</button>
            <button class="w-12 h-12 rounded-full bg-purple-500 text-white font-bold shadow-lg transform scale-110">5</button>
          </div>
        </div>
        <div class="mb-6">
          <label class="block text-sm font-bold text-gray-700 mb-2">نظر شما:</label>
          <textarea class="w-full px-4 py-3 border-2 border-purple-200 rounded-xl focus:ring-2 focus:ring-purple-500 focus:border-purple-500 text-gray-700" rows="3" placeholder="تجربه خود را با ما به اشتراک بگذارید..."></textarea>
        </div>
        <button class="w-full bg-gradient-to-r from-purple-500 to-pink-500 hover:from-purple-600 hover:to-pink-600 text-white font-bold py-4 px-6 rounded-xl transition-all transform hover:scale-105 shadow-lg">
          ارسال نظر ✨
        </button>
      </div>
    </div>
  `
}

const MinimalTemplate = {
  template: `
    <div class="bg-white border border-gray-200 rounded-lg overflow-hidden">
      <div class="px-4 py-4 text-center border-b border-gray-100">
        <h2 class="text-xl font-semibold text-gray-800 mb-1">نظرسنجی</h2>
        <p class="text-gray-600 text-sm">تجربه خرید شما</p>
      </div>
      <div class="px-4 py-4">
        <div class="mb-6">
          <h3 class="text-base font-medium text-gray-800 mb-3">سطح رضایت:</h3>
          <div class="flex items-center justify-center space-x-4 space-x-reverse">
            <label class="flex items-center">
              <input type="radio" name="rating" class="mr-2">
              <span class="text-sm text-gray-600">ضعیف</span>
            </label>
            <label class="flex items-center">
              <input type="radio" name="rating" class="mr-2">
              <span class="text-sm text-gray-600">متوسط</span>
            </label>
            <label class="flex items-center">
              <input type="radio" name="rating" class="mr-2" checked>
              <span class="text-sm text-gray-600">خوب</span>
            </label>
            <label class="flex items-center">
              <input type="radio" name="rating" class="mr-2">
              <span class="text-sm text-gray-600">عالی</span>
            </label>
          </div>
        </div>
        <div class="mb-6">
          <label class="block text-sm font-medium text-gray-700 mb-2">نظر:</label>
          <textarea class="w-full px-3 py-2 border border-gray-300 rounded focus:ring-1 focus:ring-gray-400 focus:border-gray-400" rows="3" placeholder="نظر خود را بنویسید..."></textarea>
        </div>
        <button class="w-full bg-gray-800 hover:bg-gray-900 text-white font-medium py-2 px-4 rounded transition-colors">
          ارسال
        </button>
      </div>
    </div>
  `
}

const GradientTemplate = {
  template: `
    <div class="bg-gradient-to-br from-indigo-500 via-purple-500 to-pink-500 rounded-2xl shadow-2xl overflow-hidden">
      <div class="px-4 py-4 text-white text-center">
        <div class="w-16 h-16 bg-white bg-opacity-20 rounded-full mx-auto mb-4 flex items-center justify-center backdrop-blur-sm">
          <svg class="w-8 h-8" fill="currentColor" viewBox="0 0 20 20">
            <path d="M2 10a8 8 0 018-8v8h8a8 8 0 11-16 0z"></path>
            <path d="M12 2.252A8.014 8.014 0 0117.748 8H12V2.252z"></path>
          </svg>
        </div>
        <h2 class="text-2xl font-bold mb-2">نظرسنجی خرید</h2>
        <p class="text-indigo-100">تجربه شما برای بهبود خدمات</p>
      </div>
      <div class="px-4 py-4 bg-white bg-opacity-10 backdrop-blur-sm">
        <div class="mb-6">
          <h3 class="text-lg font-semibold text-white mb-4 text-center">امتیاز دهید</h3>
          <div class="flex justify-center space-x-3 space-x-reverse">
            <button class="w-10 h-10 rounded-lg bg-white bg-opacity-20 hover:bg-opacity-30 text-white font-bold transition-all transform hover:scale-110">1</button>
            <button class="w-10 h-10 rounded-lg bg-white bg-opacity-20 hover:bg-opacity-30 text-white font-bold transition-all transform hover:scale-110">2</button>
            <button class="w-10 h-10 rounded-lg bg-white bg-opacity-20 hover:bg-opacity-30 text-white font-bold transition-all transform hover:scale-110">3</button>
            <button class="w-10 h-10 rounded-lg bg-white bg-opacity-30 text-white font-bold shadow-lg transform scale-110">4</button>
            <button class="w-10 h-10 rounded-lg bg-white bg-opacity-20 hover:bg-opacity-30 text-white font-bold transition-all transform hover:scale-110">5</button>
          </div>
        </div>
        <div class="mb-6">
          <label class="block text-sm font-medium text-white mb-2">نظر شما:</label>
          <textarea class="w-full px-4 py-3 bg-white bg-opacity-20 backdrop-blur-sm border border-white border-opacity-30 rounded-xl text-white placeholder-white placeholder-opacity-70 focus:ring-2 focus:ring-white focus:border-white" rows="3" placeholder="تجربه خود را با ما به اشتراک بگذارید..."></textarea>
        </div>
        <button class="w-full bg-white bg-opacity-20 hover:bg-opacity-30 backdrop-blur-sm text-white font-semibold py-3 px-6 rounded-xl transition-all transform hover:scale-105 border border-white border-opacity-30">
          ارسال نظر
        </button>
      </div>
    </div>
  `
}

const selectedTemplateComponent = computed(() => {
  switch (selectedTemplate.value) {
    case 'modern': return ModernTemplate
    case 'elegant': return ElegantTemplate
    case 'colorful': return ColorfulTemplate
    case 'minimal': return MinimalTemplate
    case 'gradient': return GradientTemplate
    default: return ModernTemplate
  }
})

const selectTemplate = (templateId: string) => {
  selectedTemplate.value = templateId
}

const saveTemplate = async () => {
  saving.value = true
  try {
    await new Promise(resolve => setTimeout(resolve, 1500))
    alert('تمپلیت با موفقیت ذخیره شد!')
  } finally {
    saving.value = false
  }
}
</script>

<style scoped>
.space-x-reverse > :not([hidden]) ~ :not([hidden]) {
  --tw-space-x-reverse: 1;
}
</style> 
