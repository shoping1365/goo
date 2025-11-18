<template>
  <div v-if="show" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
    <div class="bg-white rounded-lg w-full max-w-4xl mx-4 max-h-[90vh] overflow-hidden flex flex-col">
      <!-- هدر مودال -->
      <div class="bg-gradient-to-r from-blue-600 to-purple-600 text-white p-6">
        <div class="flex justify-between items-center">
          <button @click="$emit('close')" class="text-white hover:text-blue-200 text-sm">
            بازگشت
          </button>
          <div class="flex items-center space-x-4 space-x-reverse">
            <h2 class="text-xl font-bold">{{ isEditing ? 'ویرایش چت بات' : 'ایجاد چت بات جدید' }}</h2>
            <input type="checkbox" v-model="form.isActive" class="rounded border-white text-blue-600 focus:ring-blue-500">
          </div>
        </div>
      </div>

      <!-- محتوای اصلی -->
      <div class="flex-1 p-6 overflow-y-auto">
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <!-- بخش چپ - تنظیمات اصلی -->
          <div class="space-y-6">
            <div>
              <h3 class="text-lg font-medium text-gray-900 mb-4">اطلاعات اصلی</h3>
              
              <!-- نام بات -->
              <div class="mb-4">
                <label class="block text-sm font-medium text-gray-700 mb-2">نام بات *</label>
                <input v-model="form.name" type="text" placeholder="نام چت بات را وارد کنید" class="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
              </div>
              
              <!-- توضیحات -->
              <div class="mb-4">
                <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات</label>
                <textarea v-model="form.description" placeholder="توضیحات بات را وارد کنید" rows="3" class="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"></textarea>
              </div>
              
              <!-- نوع بات -->
              <div class="mb-4">
                <label class="block text-sm font-medium text-gray-700 mb-2">نوع بات *</label>
                <select v-model="form.type" class="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                  <option value="">انتخاب کنید</option>
                  <option value="پشتیبانی عمومی">پشتیبانی عمومی</option>
                  <option value="فروش و مشاوره">فروش و مشاوره</option>
                  <option value="سوالات متداول">سوالات متداول</option>
                  <option value="رزرو و سفارش">رزرو و سفارش</option>
                  <option value="سفارشی">سفارشی</option>
                </select>
              </div>
              
              <!-- آیکون -->
              <div class="mb-4">
                <label class="block text-sm font-medium text-gray-700 mb-2">آیکون</label>
                <div class="grid grid-cols-6 gap-2">
                  <button v-for="icon in availableIcons" :key="icon" 
                          @click="form.icon = icon"
                          :class="['w-12 h-12 rounded-lg flex items-center justify-center text-xl', form.icon === icon ? 'bg-blue-100 border-2 border-blue-500' : 'bg-gray-100 hover:bg-gray-200']">
                    {{ icon }}
                  </button>
                </div>
              </div>
            </div>

            <!-- تنظیمات هوش مصنوعی -->
            <div>
              <h3 class="text-lg font-medium text-gray-900 mb-4">تنظیمات هوش مصنوعی</h3>
              
              <div class="space-y-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">مدل AI</label>
                  <select v-model="form.aiModel" class="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                    <option value="gpt-4">GPT-4 (پیشرفته)</option>
                    <option value="gpt-3.5-turbo">GPT-3.5 Turbo (متوسط)</option>
                    <option value="claude">Claude (تخصصی)</option>
                  </select>
                </div>
                
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">دمای پاسخ‌دهی</label>
                  <div class="flex items-center space-x-3 space-x-reverse">
                    <input v-model="form.temperature" type="range" min="0" max="1" step="0.1" class="flex-1">
                    <span class="text-sm text-gray-600 w-12">{{ form.temperature }}</span>
                  </div>
                  <p class="text-xs text-gray-500 mt-1">مقدار پایین: پاسخ‌های دقیق‌تر، مقدار بالا: پاسخ‌های خلاقانه‌تر</p>
                </div>
                
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر طول پاسخ</label>
                  <input v-model="form.maxTokens" type="number" min="50" max="2000" class="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                </div>
              </div>
            </div>
          </div>

          <!-- بخش راست - تنظیمات پیشرفته -->
          <div class="space-y-6">
            <!-- شخصیت بات -->
            <div>
              <h3 class="text-lg font-medium text-gray-900 mb-4">شخصیت بات</h3>
              
              <div class="mb-4">
                <label class="block text-sm font-medium text-gray-700 mb-2">شخصیت و سبک پاسخ‌دهی</label>
                <textarea v-model="form.personality" placeholder="شخصیت بات را توصیف کنید. مثلاً: دوستانه، حرفه‌ای، شوخ‌طبع..." rows="4" class="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"></textarea>
              </div>
              
              <div class="mb-4">
                <label class="block text-sm font-medium text-gray-700 mb-2">دانش پایه</label>
                <textarea v-model="form.baseKnowledge" placeholder="اطلاعات پایه‌ای که بات باید بداند..." rows="4" class="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"></textarea>
              </div>
            </div>

            <!-- تنظیمات پاسخ‌دهی -->
            <div>
              <h3 class="text-lg font-medium text-gray-900 mb-4">تنظیمات پاسخ‌دهی</h3>
              
              <div class="space-y-3">
                <label class="flex items-center">
                  <input v-model="form.autoRespond" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                  <span class="mr-2 text-sm text-gray-700">پاسخ خودکار</span>
                </label>
                
                <label class="flex items-center">
                  <input v-model="form.escalateToHuman" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                  <span class="mr-2 text-sm text-gray-700">انتقال به اپراتور در صورت عدم اطمینان</span>
                </label>
                
                <label class="flex items-center">
                  <input v-model="form.learnFromConversations" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                  <span class="mr-2 text-sm text-gray-700">یادگیری از مکالمات</span>
                </label>
                
                <label class="flex items-center">
                  <input v-model="form.multiLanguage" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                  <span class="mr-2 text-sm text-gray-700">پشتیبانی از چند زبان</span>
                </label>
              </div>
            </div>

            <!-- تنظیمات زمان‌بندی -->
            <div>
              <h3 class="text-lg font-medium text-gray-900 mb-4">زمان‌بندی فعالیت</h3>
              
              <div class="space-y-3">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">ساعات کاری</label>
                  <div class="grid grid-cols-2 gap-2">
                    <input v-model="form.workStartTime" type="time" class="p-2 border border-gray-300 rounded">
                    <input v-model="form.workEndTime" type="time" class="p-2 border border-gray-300 rounded">
                  </div>
                </div>
                
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">منطقه زمانی</label>
                  <select v-model="form.timezone" class="w-full p-2 border border-gray-300 rounded">
                    <option value="Asia/Tehran">تهران (UTC+3:30)</option>
                    <option value="UTC">UTC</option>
                    <option value="Europe/London">لندن (UTC+0)</option>
                  </select>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- فوتر مودال -->
      <div class="bg-gray-50 px-6 py-4 flex justify-between items-center">
        <button @click="$emit('close')" class="px-4 py-2 text-gray-600 hover:text-gray-800">
          انصراف
        </button>
        <div class="flex space-x-3 space-x-reverse">
          <button @click="$emit('reset')" class="px-4 py-2 text-gray-600 hover:text-gray-800">
            بازنشانی
          </button>
          <button @click="$emit('save')" class="bg-blue-600 text-white px-6 py-2 rounded-lg hover:bg-blue-700">
            {{ isEditing ? 'ویرایش' : 'ایجاد' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// Props
const props = defineProps({
  show: {
    type: Boolean,
    required: true
  },
  isEditing: {
    type: Boolean,
    required: true
  },
  form: {
    type: Object,
    required: true
  }
})

// Emits
const emit = defineEmits(['close', 'save', 'reset'])

// آیکون‌های موجود
const availableIcons = ['🤖', '👾', '🎯', '💡', '🔮', '🌟', '🚀', '🎪', '🎨', '🎭', '🎪', '🎯']
</script> 
