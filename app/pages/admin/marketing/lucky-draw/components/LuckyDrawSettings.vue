<template>
  <div class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
    <div class="relative top-20 mx-auto p-5 border w-11/12 md:w-4/5 lg:w-3/4 shadow-lg rounded-md bg-white">
      <div class="mt-3">
        <!-- هدر مودال -->
        <div class="flex justify-between items-center mb-6">
          <h3 class="text-lg font-medium text-gray-900">تنظیمات گرونه شانس</h3>
          <button 
            class="text-gray-400 hover:text-gray-600"
            @click="$emit('close')"
          >
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>

        <!-- فرم تنظیمات -->
        <form class="space-y-8" @submit.prevent="saveSettings">
          <!-- تنظیمات عمومی -->
          <div class="bg-gray-50 p-6 rounded-lg">
            <h4 class="text-lg font-medium text-gray-900 mb-4">تنظیمات عمومی</h4>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">شروع خودکار گرونه‌ها</label>
                <select 
                  v-model="settings.autoStart"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                  <option value="enabled">فعال</option>
                  <option value="disabled">غیرفعال</option>
                  <option value="manual">دستی</option>
                </select>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">پایان خودکار گرونه‌ها</label>
                <select 
                  v-model="settings.autoEnd"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                  <option value="enabled">فعال</option>
                  <option value="disabled">غیرفعال</option>
                  <option value="manual">دستی</option>
                </select>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر گرونه‌های همزمان</label>
                <input 
                  v-model.number="settings.maxConcurrentDraws"
                  type="number"
                  min="1"
                  max="10"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">فاصله زمانی بین گرونه‌ها (دقیقه)</label>
                <input 
                  v-model.number="settings.timeBetweenDraws"
                  type="number"
                  min="0"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                />
              </div>
            </div>
          </div>

          <!-- تنظیمات امنیتی -->
          <div class="bg-gray-50 p-6 rounded-lg">
            <h4 class="text-lg font-medium text-gray-900 mb-4">تنظیمات امنیتی</h4>
            <div class="space-y-4">
              <div class="flex items-center">
                <input 
                  id="preventMultipleEntries"
                  v-model="settings.preventMultipleEntries"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label for="preventMultipleEntries" class="mr-2 text-sm text-gray-700">جلوگیری از شرکت چندباره</label>
              </div>
              <div class="flex items-center">
                <input 
                  id="requireVerification"
                  v-model="settings.requireVerification"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label for="requireVerification" class="mr-2 text-sm text-gray-700">نیاز به تایید شرکت</label>
              </div>
              <div class="flex items-center">
                <input 
                  id="logAllActions"
                  v-model="settings.logAllActions"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label for="logAllActions" class="mr-2 text-sm text-gray-700">ثبت تمام عملیات</label>
              </div>
              <div class="flex items-center">
                <input 
                  id="captchaEnabled"
                  v-model="settings.captchaEnabled"
                  type="checkbox"
                  class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
                <label for="captchaEnabled" class="mr-2 text-sm text-gray-700">فعال‌سازی کپچا</label>
              </div>
            </div>
          </div>

          <!-- تنظیمات اعلان‌ها -->
          <div class="bg-gray-50 p-6 rounded-lg">
            <h4 class="text-lg font-medium text-gray-900 mb-4">تنظیمات اعلان‌ها</h4>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">اعلان‌های ایمیل</label>
                <select 
                  v-model="settings.emailNotifications"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                  <option value="all">همه</option>
                  <option value="winners">فقط برندگان</option>
                  <option value="none">هیچ‌کدام</option>
                </select>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">اعلان‌های پیامک</label>
                <select 
                  v-model="settings.smsNotifications"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                  <option value="all">همه</option>
                  <option value="winners">فقط برندگان</option>
                  <option value="none">هیچ‌کدام</option>
                </select>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">اعلان‌های push</label>
                <select 
                  v-model="settings.pushNotifications"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                  <option value="all">همه</option>
                  <option value="winners">فقط برندگان</option>
                  <option value="none">هیچ‌کدام</option>
                </select>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">اعلان‌های ادمین</label>
                <select 
                  v-model="settings.adminNotifications"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                  <option value="all">همه</option>
                  <option value="important">فقط موارد مهم</option>
                  <option value="none">هیچ‌کدام</option>
                </select>
              </div>
            </div>
          </div>

          <!-- تنظیمات پیشرفته -->
          <div class="bg-gray-50 p-6 rounded-lg">
            <h4 class="text-lg font-medium text-gray-900 mb-4">تنظیمات پیشرفته</h4>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">تست A/B</label>
                <select 
                  v-model="settings.abTesting"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                  <option value="enabled">فعال</option>
                  <option value="disabled">غیرفعال</option>
                </select>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">هوش مصنوعی</label>
                <select 
                  v-model="settings.aiEnabled"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                  <option value="enabled">فعال</option>
                  <option value="disabled">غیرفعال</option>
                </select>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">تحلیل رفتار</label>
                <select 
                  v-model="settings.behaviorAnalytics"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                  <option value="enabled">فعال</option>
                  <option value="disabled">غیرفعال</option>
                </select>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">بهینه‌سازی خودکار</label>
                <select 
                  v-model="settings.autoOptimization"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                  <option value="enabled">فعال</option>
                  <option value="disabled">غیرفعال</option>
                </select>
              </div>
            </div>
          </div>

          <!-- تنظیمات ظاهری -->
          <div class="bg-gray-50 p-6 rounded-lg">
            <h4 class="text-lg font-medium text-gray-900 mb-4">تنظیمات ظاهری</h4>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">تم پیش‌فرض</label>
                <select 
                  v-model="settings.defaultTheme"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                  <option value="light">روشن</option>
                  <option value="dark">تیره</option>
                  <option value="auto">خودکار</option>
                </select>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">زبان پیش‌فرض</label>
                <select 
                  v-model="settings.defaultLanguage"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                  <option value="fa">فارسی</option>
                  <option value="en">انگلیسی</option>
                  <option value="ar">عربی</option>
                </select>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">اندازه فونت</label>
                <select 
                  v-model="settings.fontSize"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                  <option value="small">کوچک</option>
                  <option value="medium">متوسط</option>
                  <option value="large">بزرگ</option>
                </select>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">انیمیشن‌ها</label>
                <select 
                  v-model="settings.animations"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                  <option value="enabled">فعال</option>
                  <option value="disabled">غیرفعال</option>
                  <option value="reduced">کاهش یافته</option>
                </select>
              </div>
            </div>
          </div>

          <!-- دکمه‌های عملیات -->
          <div class="flex justify-end space-x-3 space-x-reverse">
            <button 
              type="button"
              class="px-4 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400 transition-colors"
              @click="resetToDefaults"
            >
              بازنشانی به پیش‌فرض
            </button>
            <button 
              type="submit"
              :disabled="isSaving"
              class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors disabled:opacity-50"
            >
              {{ isSaving ? 'در حال ذخیره...' : 'ذخیره تنظیمات' }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
const emit = defineEmits(['close'])

// متغیرهای reactive
const isSaving = ref(false)

// تنظیمات
const settings = ref({
  // تنظیمات عمومی
  autoStart: 'enabled',
  autoEnd: 'enabled',
  maxConcurrentDraws: 5,
  timeBetweenDraws: 30,
  
  // تنظیمات امنیتی
  preventMultipleEntries: true,
  requireVerification: false,
  logAllActions: true,
  captchaEnabled: true,
  
  // تنظیمات اعلان‌ها
  emailNotifications: 'winners',
  smsNotifications: 'winners',
  pushNotifications: 'all',
  adminNotifications: 'important',
  
  // تنظیمات پیشرفته
  abTesting: 'disabled',
  aiEnabled: 'enabled',
  behaviorAnalytics: 'enabled',
  autoOptimization: 'disabled',
  
  // تنظیمات ظاهری
  defaultTheme: 'light',
  defaultLanguage: 'fa',
  fontSize: 'medium',
  animations: 'enabled'
})

// توابع
const saveSettings = async () => {
  try {
    isSaving.value = true
    
    // شبیه‌سازی API call
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    // ذخیره تنظیمات
    alert('تنظیمات با موفقیت ذخیره شدند')
    
    emit('close')
  } catch (_error) {
    alert('خطا در ذخیره تنظیمات')
  } finally {
    isSaving.value = false
  }
}

const resetToDefaults = () => {
  if (confirm('آیا از بازنشانی تمام تنظیمات به مقادیر پیش‌فرض اطمینان دارید؟')) {
    // بازنشانی به مقادیر پیش‌فرض
    settings.value = {
      autoStart: 'enabled',
      autoEnd: 'enabled',
      maxConcurrentDraws: 5,
      timeBetweenDraws: 30,
      preventMultipleEntries: true,
      requireVerification: false,
      logAllActions: true,
      captchaEnabled: true,
      emailNotifications: 'winners',
      smsNotifications: 'winners',
      pushNotifications: 'all',
      adminNotifications: 'important',
      abTesting: 'disabled',
      aiEnabled: 'enabled',
      behaviorAnalytics: 'enabled',
      autoOptimization: 'disabled',
      defaultTheme: 'light',
      defaultLanguage: 'fa',
      fontSize: 'medium',
      animations: 'enabled'
    }
    
    alert('تنظیمات به مقادیر پیش‌فرض بازنشانی شدند')
  }
}
</script>
