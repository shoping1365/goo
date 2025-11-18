<template>
  <div class="space-y-8">
    <!-- هدر بخش -->
    <div class="relative overflow-hidden bg-gradient-to-br from-blue-50 via-indigo-50 to-purple-50 rounded-2xl p-8">
      <div class="absolute top-0 right-0 w-32 h-32 bg-gradient-to-bl from-blue-200/30 to-purple-200/30 rounded-full -translate-y-16 translate-x-16"></div>
      <div class="absolute bottom-0 left-0 w-24 h-24 bg-gradient-to-tr from-indigo-200/30 to-pink-200/30 rounded-full translate-y-12 -translate-x-12"></div>
      
      <div class="relative z-10">
        <div class="flex items-center mb-4">
          <div class="w-12 h-12 bg-gradient-to-br from-blue-500 to-purple-600 rounded-xl flex items-center justify-center mr-4">
            <i class="i-heroicons-code-bracket text-white text-xl"></i>
          </div>
          <div>
            <h2 class="text-3xl font-bold bg-gradient-to-r from-gray-800 to-gray-600 bg-clip-text text-transparent">تنظیمات API</h2>
            <p class="text-gray-600 mt-1">مدیریت تنظیمات API های خارجی و سرویس‌های هوش مصنوعی</p>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات OpenAI -->
    <div class="bg-white rounded-2xl shadow-lg border border-gray-100 overflow-hidden">
      <div class="bg-gradient-to-r from-green-500 to-emerald-600 px-6 py-4">
        <div class="flex items-center">
          <i class="i-heroicons-sparkles text-white text-xl mr-3"></i>
          <h3 class="text-xl font-bold text-white">تنظیمات OpenAI</h3>
        </div>
      </div>
      
      <div class="p-6 space-y-6">
        <!-- اطلاعات حساب -->
        <div class="bg-gradient-to-r from-green-50 to-emerald-50 rounded-xl p-6 border border-green-200">
          <div class="flex items-center justify-between mb-4">
            <h4 class="text-lg font-semibold text-gray-900 flex items-center">
              <i class="i-heroicons-credit-card text-green-600 mr-2"></i>
              اطلاعات حساب
            </h4>
            <button 
              @click="fetchUsageData"
              :disabled="fetchingUsage"
              class="flex items-center px-3 py-2 text-sm font-medium text-green-600 bg-green-50 border border-green-200 rounded-lg hover:bg-green-100 hover:border-green-300 transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
            >
              <i v-if="fetchingUsage" class="i-heroicons-arrow-path animate-spin mr-2"></i>
              <i v-else class="i-heroicons-arrow-path mr-2"></i>
              {{ fetchingUsage ? 'در حال دریافت...' : 'بروزرسانی آمار' }}
            </button>
          </div>
          <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
            <div class="bg-white rounded-lg p-6 border border-green-200">
              <div class="flex items-center justify-between mb-2">
                <span class="text-sm font-medium text-gray-600">میزان شارژ حساب</span>
                <i class="i-heroicons-currency-dollar text-green-600"></i>
              </div>
              <div class="text-2xl font-bold text-gray-900">{{ openAISettings.accountBalance || '0.00' }} $</div>
              <div class="text-xs text-gray-500 mt-1">آخرین بروزرسانی: {{ openAISettings.lastBalanceUpdate || 'نامشخص' }}</div>
              <div class="text-xs text-blue-500 mt-1" v-if="openAISettings.accountBalance === '0.00'">
                <i class="i-heroicons-information-circle mr-1"></i>
                برخی API Key ها دسترسی به اطلاعات billing ندارند
              </div>
            </div>
            
            <div class="bg-white rounded-lg p-6 border border-green-200">
              <div class="flex items-center justify-between mb-2">
                <span class="text-sm font-medium text-gray-600">میزان مصرف ماهانه</span>
                <i class="i-heroicons-chart-bar text-blue-600"></i>
              </div>
              <div class="text-2xl font-bold text-gray-900">{{ openAISettings.monthlyUsage || '0.00' }} $</div>
              <div class="text-xs text-gray-500 mt-1">از ابتدای ماه جاری</div>
            </div>
            
            <div class="bg-white rounded-lg p-6 border border-green-200">
              <div class="flex items-center justify-between mb-2">
                <span class="text-sm font-medium text-gray-600">تعداد درخواست‌ها</span>
                <i class="i-heroicons-arrow-path text-purple-600"></i>
              </div>
              <div class="text-2xl font-bold text-gray-900">{{ openAISettings.totalRequests || '0' }}</div>
              <div class="text-xs text-gray-500 mt-1">امروز: {{ openAISettings.todayRequests || '0' }}</div>
            </div>
          </div>
        </div>

        <!-- تنظیمات کلید و آدرس -->
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <div class="group">
            <label class="block text-sm font-semibold text-gray-700 mb-3 flex items-center">
              <i class="i-heroicons-key text-green-500 mr-2"></i>
              کلید API
            </label>
            <div class="relative">
              <input 
                v-model="openAISettings.apiKey" 
                :type="showApiKey ? 'text' : 'password'"
                class="w-full px-4 py-3 pr-12 border-2 border-gray-200 rounded-xl focus:outline-none focus:border-green-500 focus:ring-4 focus:ring-green-100 transition-all duration-200 group-hover:border-gray-300"
                placeholder="sk-..."
              >
              <button 
                type="button"
                @click="showApiKey = !showApiKey"
                class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 hover:text-gray-600 transition-colors"
              >
                <i :class="showApiKey ? 'i-heroicons-eye-slash' : 'i-heroicons-eye'" class="text-lg"></i>
              </button>
            </div>
            <p class="text-xs text-gray-500 mt-1">کلید API شما از OpenAI</p>
          </div>
          
          <div class="group">
            <label class="block text-sm font-semibold text-gray-700 mb-3 flex items-center">
              <i class="i-heroicons-globe-alt text-green-500 mr-2"></i>
              آدرس API
            </label>
            <input 
              v-model="openAISettings.apiUrl" 
              type="url" 
              class="w-full px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:border-green-500 focus:ring-4 focus:ring-green-100 transition-all duration-200 group-hover:border-gray-300"
              placeholder="https://api.openai.com/v1"
            >
            <p class="text-xs text-gray-500 mt-1">آدرس پایه API (پیش‌فرض: OpenAI)</p>
          </div>
        </div>

        

        <!-- لیست تمام مدل‌های موجود (تاشو) -->
        <div class="bg-gradient-to-r from-blue-50 to-indigo-50 rounded-xl p-6 border border-blue-200">
          <h4 class="text-lg font-semibold text-gray-900 mb-4 flex items-center" @click="isModelsCollapsed = !isModelsCollapsed">
            <i class="i-heroicons-list-bullet text-blue-600 mr-2"></i>
            لیست تمام مدل‌های OpenAI
          </h4>
          <div v-show="!isModelsCollapsed" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            <div 
              v-for="model in availableModels" 
              :key="model.id"
              class="bg-white rounded-lg p-6 border border-gray-200 hover:border-blue-300 transition-all duration-200"
            >
              <div class="flex items-start justify-between mb-2">
                <h5 class="font-semibold text-gray-900">{{ model.name }}</h5>
                <span :class="getModelStatusClass(model.is_active)" class="px-2 py-1 rounded-full text-xs font-medium">
                  {{ model.is_active ? 'فعال' : 'غیرفعال' }}
                </span>
              </div>
              <p class="text-sm text-gray-600 mb-2">{{ model.description }}</p>
              <div class="flex items-center justify-between text-xs text-gray-500">
                <span>{{ model.category }}</span>
                <span>{{ model.cost_per_1k }}</span>
              </div>
              <div class="mt-2 text-xs text-gray-500">
                حداکثر توکن: {{ model.max_tokens.toLocaleString() }}
              </div>
            </div>
          </div>
        </div>

        <!-- تنظیمات مدل برای هر بخش -->
        <div class="bg-gradient-to-r from-purple-50 to-pink-50 rounded-xl p-6 border border-purple-200">
          <h4 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
            <i class="i-heroicons-cog-6-tooth text-purple-600 mr-2"></i>
            تنظیمات مدل برای هر بخش
          </h4>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div 
              v-for="(section, sectionKey) in sectionConfigs" 
              :key="sectionKey"
              class="bg-white rounded-lg p-3 border border-gray-200"
            >
              <div class="flex items-center justify-between mb-3">
                <h5 class="font-semibold text-gray-900">{{ section.title }}</h5>
                <div class="flex items-center gap-3">
                  <NuxtLink v-if="section.link" :to="section.link" class="text-xs text-blue-600 hover:underline">
                    {{ section.linkText || 'مشاهده' }}
                  </NuxtLink>
                  <input 
                    v-model="section.isEnabled" 
                    type="checkbox" 
                    class="w-4 h-4 text-purple-600 border-gray-300 rounded focus:ring-purple-500"
                  >
                </div>
              </div>
              <p class="text-sm text-gray-600 mb-3">{{ section.description }}</p>
              
              <div class="space-y-3">
                <div class="grid grid-cols-1 md:grid-cols-2 gap-3 items-end">
                  <div>
                    <label class="block text-xs font-medium text-gray-700 mb-1">مدل انتخاب شده</label>
                    <select 
                      v-model="section.model" 
                      :disabled="!section.isEnabled"
                      class="w-full px-3 py-1.5 text-sm border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500 disabled:opacity-50"
                    >
                      <option v-for="model in availableModels" :key="model.id" :value="model.id">
                        {{ model.name }}
                      </option>
                    </select>
                  </div>
                  <div>
                    <label class="block text-xs font-medium text-gray-700 mb-1">حداکثر توکن</label>
                    <input 
                      v-model="section.maxTokens" 
                      type="number" 
                      min="100" 
                      max="128000"
                      :disabled="!section.isEnabled"
                      class="w-full px-3 py-1.5 text-sm border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500 disabled:opacity-50"
                    >
                  </div>
                </div>
                
                <div class="mt-1">
                  <label class="block text-xs font-medium text-gray-700 mb-1">دمای پاسخ‌دهی</label>
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <input 
                      v-model="section.temperature" 
                      type="range" 
                      min="0" 
                      max="2" 
                      step="0.1"
                      :disabled="!section.isEnabled"
                      class="flex-1 h-1.5 bg-gray-200 rounded-lg appearance-none cursor-pointer slider disabled:opacity-50"
                    >
                    <span class="text-xs text-gray-600 w-8">{{ section.temperature }}</span>
                  </div>
                </div>
                
                
              </div>
            </div>
          </div>
        </div>

        <!-- محدودیت‌ها -->
        <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
          <div class="group">
            <label class="block text-sm font-semibold text-gray-700 mb-3 flex items-center">
              <i class="i-heroicons-arrow-path text-green-500 mr-2"></i>
              محدودیت درخواست (دقیقه)
            </label>
            <input 
              v-model="openAISettings.rateLimit" 
              type="number" 
              min="1" 
              max="1000"
              class="w-full px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:border-green-500 focus:ring-4 focus:ring-green-100 transition-all duration-200 group-hover:border-gray-300"
              placeholder="60"
            >
          </div>
          
          <div class="group">
            <label class="block text-sm font-semibold text-gray-700 mb-3 flex items-center">
              <i class="i-heroicons-clock text-green-500 mr-2"></i>
              تایم‌اوت (ثانیه)
            </label>
            <input 
              v-model="openAISettings.timeout" 
              type="number" 
              min="5" 
              max="120"
              class="w-full px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:border-green-500 focus:ring-4 focus:ring-green-100 transition-all duration-200 group-hover:border-gray-300"
              placeholder="30"
            >
          </div>
          
          <div class="group">
            <label class="block text-sm font-semibold text-gray-700 mb-3 flex items-center">
              <i class="i-heroicons-currency-dollar text-green-500 mr-2"></i>
              حداکثر هزینه روزانه ($)
            </label>
            <input 
              v-model="openAISettings.maxDailyCost" 
              type="number" 
              min="0" 
              max="1000"
              step="0.01"
              class="w-full px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:border-green-500 focus:ring-4 focus:ring-green-100 transition-all duration-200 group-hover:border-gray-300"
              placeholder="10.00"
            >
          </div>
        </div>

        <!-- صفحات مصرف کننده -->
        <div class="group">
          <label class="block text-sm font-semibold text-gray-700 mb-3 flex items-center">
            <i class="i-heroicons-document-text text-green-500 mr-2"></i>
            صفحات مصرف کننده API
          </label>
          <div class="space-y-3">
            <div v-for="(page, index) in openAISettings.consumingPages" :key="index" class="flex items-center gap-3">
              <input 
                v-model="openAISettings.consumingPages[index]" 
                type="text" 
                class="flex-1 px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:border-green-500 focus:ring-4 focus:ring-green-100 transition-all duration-200"
                placeholder="مثلاً: /admin/seo/content-generation"
              >
              <button 
                type="button" 
                @click="removeConsumingPage(index)"
                class="px-3 py-3 bg-gradient-to-r from-red-500 to-red-600 text-white rounded-xl hover:from-red-600 hover:to-red-700 transition-all duration-200"
              >
                <i class="i-heroicons-trash text-sm"></i>
              </button>
            </div>
            <button 
              type="button" 
              @click="addConsumingPage"
              class="w-full px-4 py-3 bg-gradient-to-r from-green-500 to-green-600 text-white rounded-xl hover:from-green-600 hover:to-green-700 transition-all duration-200 shadow-md hover:shadow-lg"
            >
              <i class="i-heroicons-plus mr-2"></i>
              افزودن صفحه جدید
            </button>
          </div>
        </div>

        <!-- وضعیت فعال‌سازی -->
        <div class="flex items-center p-6 bg-gradient-to-r from-green-50 to-emerald-50 rounded-xl border border-green-200">
          <input 
            v-model="openAISettings.enabled" 
            type="checkbox" 
            id="openaiEnabled"
            class="w-5 h-5 text-green-600 border-gray-300 rounded focus:ring-green-500"
          >
          <label for="openaiEnabled" class="mr-3 text-sm font-semibold text-gray-700 flex items-center">
            <i class="i-heroicons-check-circle text-green-500 mr-2"></i>
            فعال‌سازی OpenAI
          </label>
          <div class="mr-auto">
            <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">
              <i class="i-heroicons-information-circle mr-1"></i>
              API OpenAI در سیستم فعال خواهد بود
            </span>
          </div>
        </div>
      </div>
    </div>

    <!-- دکمه‌های عملیات -->
    <div class="bg-gradient-to-r from-gray-50 to-gray-100 rounded-2xl p-6">
      <div class="flex flex-col sm:flex-row justify-end space-y-3 sm:space-y-0 sm:space-x-3 sm:space-x-reverse">
        <button 
          type="button" 
          @click="testOpenAIConnection"
          :disabled="testingConnection"
          class="px-8 py-3 bg-gradient-to-r from-blue-500 to-indigo-600 text-white rounded-xl hover:from-blue-600 hover:to-indigo-700 disabled:opacity-50 transition-all duration-200 shadow-lg hover:shadow-xl font-semibold flex items-center justify-center"
        >
          <i v-if="testingConnection" class="i-heroicons-arrow-path animate-spin mr-2"></i>
          <i v-else class="i-heroicons-arrow-path mr-2"></i>
          {{ testingConnection ? 'در حال تست...' : 'تست اتصال' }}
        </button>
        <button 
          type="button" 
          @click="resetApiSettings"
          class="px-8 py-3 border-2 border-gray-300 text-gray-700 rounded-xl hover:bg-gray-50 hover:border-gray-400 transition-all duration-200 font-semibold flex items-center justify-center"
        >
          <i class="i-heroicons-arrow-path mr-2"></i>
          بازنشانی تنظیمات
        </button>
        <button 
          type="button" 
          @click="saveApiSettings"
          :disabled="savingApi"
          class="px-8 py-3 bg-gradient-to-r from-green-500 to-emerald-600 text-white rounded-xl hover:from-green-600 hover:to-emerald-700 disabled:opacity-50 transition-all duration-200 shadow-lg hover:shadow-xl font-semibold flex items-center justify-center"
        >
          <i v-if="savingApi" class="i-heroicons-arrow-path animate-spin mr-2"></i>
          <i v-else class="i-heroicons-check mr-2"></i>
          {{ savingApi ? 'در حال ذخیره...' : 'ذخیره تنظیمات API' }}
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  openAISettings: {
    type: Object,
    required: true
  },
  availableModels: {
    type: Array,
    required: true
  },
  sectionConfigs: {
    type: Object,
    required: true
  },
  showApiKey: {
    type: Boolean,
    default: false
  },
  fetchingUsage: {
    type: Boolean,
    default: false
  },
  testingConnection: {
    type: Boolean,
    default: false
  },
  savingApi: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['save', 'reset', 'testConnection', 'fetchUsage', 'toggleApiKey', 'addConsumingPage', 'removeConsumingPage'])

// وضعیت تاشو بودن لیست مدل‌ها (به صورت پیش‌فرض بسته است)
const isModelsCollapsed = ref(true)

const saveApiSettings = () => {
  emit('save')
}

const resetApiSettings = () => {
  emit('reset')
}

const testOpenAIConnection = () => {
  emit('testConnection')
}

const fetchUsageData = () => {
  emit('fetchUsage')
}

const toggleApiKey = () => {
  emit('toggleApiKey')
}

const addConsumingPage = () => {
  emit('addConsumingPage')
}

const removeConsumingPage = (index) => {
  emit('removeConsumingPage', index)
}

const getModelStatusClass = (isActive) => {
  return isActive ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800'
}
</script>

<style scoped>
.slider {
  -webkit-appearance: none;
  appearance: none;
  background: transparent;
  cursor: pointer;
}

/* WebKit (Chrome/Edge/Safari) visible blue track */
.slider::-webkit-slider-runnable-track {
  background: linear-gradient(to right, #93c5fd, #3b82f6);
  height: 6px;
  border-radius: 9999px;
}

.slider::-webkit-slider-thumb {
  -webkit-appearance: none;
  appearance: none;
  background: #10b981;
  height: 20px;
  width: 20px;
  border-radius: 50%;
  margin-top: -7px; /* align with 6px track */
  cursor: pointer;
  transition: all 0.2s ease;
}

.slider::-webkit-slider-thumb:hover {
  background: #059669;
  transform: scale(1.1);
}

/* Firefox visible blue track */
.slider::-moz-range-track {
  background: linear-gradient(to right, #93c5fd, #3b82f6);
  height: 6px;
  border-radius: 9999px;
  border: none;
}

.slider::-moz-range-thumb {
  background: #10b981;
  height: 20px;
  width: 20px;
  border-radius: 50%;
  border: none;
  cursor: pointer;
  border: none;
  transition: all 0.2s ease;
}

.slider::-moz-range-thumb:hover {
  background: #059669;
  transform: scale(1.1);
}

.slider:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.slider:disabled::-webkit-slider-thumb {
  background: #9ca3af;
  cursor: not-allowed;
}

.slider:disabled::-moz-range-thumb {
  background: #9ca3af;
  cursor: not-allowed;
}
</style> 
