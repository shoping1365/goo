<template>
  <div class="space-y-6">
    <!-- تنظیمات کلی سیستم -->
    <div class="bg-white rounded-lg shadow">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-medium text-gray-900">تنظیمات کلی سیستم</h3>
      </div>
      <div class="px-4 py-4">
        <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
          <div class="space-y-4">
            <div class="flex items-center justify-between px-4 py-4 bg-gray-50 rounded-lg">
              <div>
                <h4 class="font-medium text-gray-900">فعال‌سازی سیستم رتبه‌بندی</h4>
                <p class="text-sm text-gray-500">فعال یا غیرفعال کردن کل سیستم</p>
              </div>
              <label class="flex items-center">
                <input v-model="settings.isSystemEnabled" type="checkbox" class="mr-2 rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <span class="text-sm text-gray-700">فعال</span>
              </label>
            </div>

            <div class="flex items-center justify-between px-4 py-4 bg-gray-50 rounded-lg">
              <div>
                <h4 class="font-medium text-gray-900">امتیازدهی خودکار</h4>
                <p class="text-sm text-gray-500">فعال‌سازی سیستم امتیازدهی خودکار</p>
              </div>
              <label class="flex items-center">
                <input v-model="settings.autoScoringEnabled" type="checkbox" class="mr-2 rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <span class="text-sm text-gray-700">فعال</span>
              </label>
            </div>

            <div class="flex items-center justify-between px-4 py-4 bg-gray-50 rounded-lg">
              <div>
                <h4 class="font-medium text-gray-900">اعلان‌های خودکار</h4>
                <p class="text-sm text-gray-500">ارسال اعلان برای تغییرات امتیاز</p>
              </div>
              <label class="flex items-center">
                <input v-model="settings.autoNotifications" type="checkbox" class="mr-2 rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <span class="text-sm text-gray-700">فعال</span>
              </label>
            </div>
          </div>

          <div class="space-y-4">
            <div class="flex items-center justify-between px-4 py-4 bg-gray-50 rounded-lg">
              <div>
                <h4 class="font-medium text-gray-900">نمایش امتیاز به کاربران</h4>
                <p class="text-sm text-gray-500">نمایش امتیاز در پروفایل کاربران</p>
              </div>
              <label class="flex items-center">
                <input v-model="settings.showScoreToUsers" type="checkbox" class="mr-2 rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <span class="text-sm text-gray-700">فعال</span>
              </label>
            </div>

            <div class="flex items-center justify-between px-4 py-4 bg-gray-50 rounded-lg">
              <div>
                <h4 class="font-medium text-gray-900">نمایش رتبه به کاربران</h4>
                <p class="text-sm text-gray-500">نمایش سطح رتبه‌بندی به کاربران</p>
              </div>
              <label class="flex items-center">
                <input v-model="settings.showLevelToUsers" type="checkbox" class="mr-2 rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <span class="text-sm text-gray-700">فعال</span>
              </label>
            </div>

            <div class="flex items-center justify-between px-4 py-4 bg-gray-50 rounded-lg">
              <div>
                <h4 class="font-medium text-gray-900">سیستم هشدار</h4>
                <p class="text-sm text-gray-500">فعال‌سازی هشدار برای امتیازات منفی</p>
              </div>
              <label class="flex items-center">
                <input v-model="settings.warningSystemEnabled" type="checkbox" class="mr-2 rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <span class="text-sm text-gray-700">فعال</span>
              </label>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات سطوح رتبه‌بندی -->
    <div class="bg-white rounded-lg shadow">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-medium text-gray-900">تنظیمات سطوح رتبه‌بندی</h3>
      </div>
      <div class="px-4 py-4">
        <div class="space-y-4">
          <div v-for="(level, index) in levelSettings" :key="index" class="border border-gray-200 rounded-lg px-4 py-4">
            <div class="flex items-center justify-between mb-4">
              <div class="flex items-center">
                <div :class="level.color" class="w-6 h-6 rounded-full mr-3"></div>
                <h4 class="font-medium text-gray-900">{{ level.name }}</h4>
              </div>
              <label class="flex items-center">
                <input v-model="level.enabled" type="checkbox" class="mr-2 rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <span class="text-sm text-gray-700">فعال</span>
              </label>
            </div>
            <div class="grid grid-cols-1 md:grid-cols-3 gapx-4 py-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">حداقل امتیاز</label>
                <input v-model="level.minScore" type="number" min="0" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">حداکثر امتیاز</label>
                <input v-model="level.maxScore" type="number" min="0" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">مزایا</label>
                <select v-model="level.benefits" multiple class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
                  <option value="discount">تخفیف ویژه</option>
                  <option value="early_access">دسترسی زودهنگام</option>
                  <option value="priority_support">پشتیبانی ویژه</option>
                  <option value="free_shipping">ارسال رایگان</option>
                  <option value="exclusive_content">محتوای انحصاری</option>
                </select>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات آستانه‌ها -->
    <div class="bg-white rounded-lg shadow">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-medium text-gray-900">تنظیمات آستانه‌ها</h3>
      </div>
      <div class="px-4 py-4">
        <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
          <div class="space-y-4">
            <div class="flex items-center justify-between px-4 py-4 bg-yellow-50 rounded-lg">
              <div>
                <h4 class="font-medium text-gray-900">حداقل امتیاز برای کاربر برتر</h4>
                <p class="text-sm text-gray-500">امتیاز مورد نیاز برای دسترسی به مزایای ویژه</p>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <input v-model="settings.minTopUserScore" type="number" min="0" max="1000" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
                <span class="text-sm text-gray-600">امتیاز</span>
              </div>
            </div>

            <div class="flex items-center justify-between px-4 py-4 bg-green-50 rounded-lg">
              <div>
                <h4 class="font-medium text-gray-900">امتیاز برای ارتقاء خودکار</h4>
                <p class="text-sm text-gray-500">امتیاز مورد نیاز برای ارتقاء خودکار سطح</p>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <input v-model="settings.autoUpgradeThreshold" type="number" min="0" max="1000" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
                <span class="text-sm text-gray-600">امتیاز</span>
              </div>
            </div>
          </div>

          <div class="space-y-4">
            <div class="flex items-center justify-between px-4 py-4 bg-red-50 rounded-lg">
              <div>
                <h4 class="font-medium text-gray-900">آستانه مسدودسازی</h4>
                <p class="text-sm text-gray-500">امتیاز منفی برای مسدودسازی خودکار</p>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <input v-model="settings.blockThreshold" type="number" min="-1000" max="0" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
                <span class="text-sm text-gray-600">امتیاز</span>
              </div>
            </div>

            <div class="flex items-center justify-between px-4 py-4 bg-orange-50 rounded-lg">
              <div>
                <h4 class="font-medium text-gray-900">آستانه هشدار</h4>
                <p class="text-sm text-gray-500">امتیاز منفی برای ارسال هشدار</p>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <input v-model="settings.warningThreshold" type="number" min="-1000" max="0" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
                <span class="text-sm text-gray-600">امتیاز</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات محدودیت‌ها -->
    <div class="bg-white rounded-lg shadow">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-medium text-gray-900">تنظیمات محدودیت‌ها</h3>
      </div>
      <div class="px-4 py-4">
        <div class="grid grid-cols-1 md:grid-cols-3 gapx-4 py-4">
          <div class="space-y-4">
            <h4 class="font-medium text-gray-900">محدودیت‌های روزانه</h4>
            <div class="space-y-3">
              <div class="flex items-center justify-between">
                <span class="text-sm text-gray-600">حداکثر امتیاز روزانه</span>
                <input v-model="settings.dailyScoreLimit" type="number" min="0" max="1000" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
              </div>
              <div class="flex items-center justify-between">
                <span class="text-sm text-gray-600">حداکثر نظرات روزانه</span>
                <input v-model="settings.dailyReviewLimit" type="number" min="0" max="100" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
              </div>
              <div class="flex items-center justify-between">
                <span class="text-sm text-gray-600">حداکثر ارجاعات روزانه</span>
                <input v-model="settings.dailyReferralLimit" type="number" min="0" max="50" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
              </div>
            </div>
          </div>

          <div class="space-y-4">
            <h4 class="font-medium text-gray-900">محدودیت‌های ماهانه</h4>
            <div class="space-y-3">
              <div class="flex items-center justify-between">
                <span class="text-sm text-gray-600">حداکثر امتیاز ماهانه</span>
                <input v-model="settings.monthlyScoreLimit" type="number" min="0" max="10000" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
              </div>
              <div class="flex items-center justify-between">
                <span class="text-sm text-gray-600">حداکثر خریدهای ماهانه</span>
                <input v-model="settings.monthlyPurchaseLimit" type="number" min="0" max="1000" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
              </div>
              <div class="flex items-center justify-between">
                <span class="text-sm text-gray-600">حداکثر فعالیت‌های ماهانه</span>
                <input v-model="settings.monthlyActivityLimit" type="number" min="0" max="10000" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
              </div>
            </div>
          </div>

          <div class="space-y-4">
            <h4 class="font-medium text-gray-900">محدودیت‌های کلی</h4>
            <div class="space-y-3">
              <div class="flex items-center justify-between">
                <span class="text-sm text-gray-600">حداکثر امتیاز کل</span>
                <input v-model="settings.maxTotalScore" type="number" min="0" max="100000" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
              </div>
              <div class="flex items-center justify-between">
                <span class="text-sm text-gray-600">حداقل سن حساب</span>
                <input v-model="settings.minAccountAge" type="number" min="0" max="365" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
              </div>
              <div class="flex items-center justify-between">
                <span class="text-sm text-gray-600">حداقل فعالیت‌ها</span>
                <input v-model="settings.minActivities" type="number" min="0" max="1000" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات اعلان‌ها -->
    <div class="bg-white rounded-lg shadow">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-medium text-gray-900">تنظیمات اعلان‌ها</h3>
      </div>
      <div class="px-4 py-4">
        <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
          <div class="space-y-4">
            <h4 class="font-medium text-gray-900">اعلان‌های کاربران</h4>
            <div class="space-y-3">
              <label class="flex items-center">
                <input v-model="settings.notifications.scoreChange" type="checkbox" class="mr-2 rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <span class="text-sm text-gray-700">تغییر امتیاز</span>
              </label>
              <label class="flex items-center">
                <input v-model="settings.notifications.levelUpgrade" type="checkbox" class="mr-2 rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <span class="text-sm text-gray-700">ارتقاء سطح</span>
              </label>
              <label class="flex items-center">
                <input v-model="settings.notifications.warning" type="checkbox" class="mr-2 rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <span class="text-sm text-gray-700">هشدار امتیاز منفی</span>
              </label>
              <label class="flex items-center">
                <input v-model="settings.notifications.blocking" type="checkbox" class="mr-2 rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <span class="text-sm text-gray-700">مسدودسازی حساب</span>
              </label>
            </div>
          </div>

          <div class="space-y-4">
            <h4 class="font-medium text-gray-900">اعلان‌های مدیران</h4>
            <div class="space-y-3">
              <label class="flex items-center">
                <input v-model="settings.adminNotifications.newTopUser" type="checkbox" class="mr-2 rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <span class="text-sm text-gray-700">کاربر برتر جدید</span>
              </label>
              <label class="flex items-center">
                <input v-model="settings.adminNotifications.userBlocked" type="checkbox" class="mr-2 rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <span class="text-sm text-gray-700">مسدودسازی کاربر</span>
              </label>
              <label class="flex items-center">
                <input v-model="settings.adminNotifications.systemAlert" type="checkbox" class="mr-2 rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <span class="text-sm text-gray-700">هشدار سیستم</span>
              </label>
              <label class="flex items-center">
                <input v-model="settings.adminNotifications.dailyReport" type="checkbox" class="mr-2 rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <span class="text-sm text-gray-700">گزارش روزانه</span>
              </label>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- دکمه‌های عملیات -->
    <div class="flex justify-end space-x-3 space-x-reverse">
      <button class="px-6 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400 transition-colors" @click="resetSettings">
        بازنشانی
      </button>
      <button class="px-6 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors" @click="saveSettings">
        ذخیره تنظیمات
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';

// Props
const props = defineProps<{
  initialSettings?: Record<string, unknown>
}>()

// Emits
const emit = defineEmits<{
  saveSettings: [settings: Record<string, unknown>]
}>()

// Reactive data
const settings = ref({
  // تنظیمات کلی
  isSystemEnabled: true,
  autoScoringEnabled: true,
  autoNotifications: true,
  showScoreToUsers: true,
  showLevelToUsers: true,
  warningSystemEnabled: true,

  // آستانه‌ها
  minTopUserScore: 500,
  autoUpgradeThreshold: 100,
  blockThreshold: -50,
  warningThreshold: -20,

  // محدودیت‌های روزانه
  dailyScoreLimit: 100,
  dailyReviewLimit: 10,
  dailyReferralLimit: 5,

  // محدودیت‌های ماهانه
  monthlyScoreLimit: 2000,
  monthlyPurchaseLimit: 50,
  monthlyActivityLimit: 500,

  // محدودیت‌های کلی
  maxTotalScore: 50000,
  minAccountAge: 7,
  minActivities: 10,

  // اعلان‌های کاربران
  notifications: {
    scoreChange: true,
    levelUpgrade: true,
    warning: true,
    blocking: true
  },

  // اعلان‌های مدیران
  adminNotifications: {
    newTopUser: true,
    userBlocked: true,
    systemAlert: true,
    dailyReport: false
  }
})

const levelSettings = ref([
  {
    name: 'برنزی',
    color: 'bg-orange-400',
    enabled: true,
    minScore: 0,
    maxScore: 100,
    benefits: ['discount']
  },
  {
    name: 'نقره‌ای',
    color: 'bg-gray-400',
    enabled: true,
    minScore: 101,
    maxScore: 500,
    benefits: ['discount', 'free_shipping']
  },
  {
    name: 'طلایی',
    color: 'bg-yellow-400',
    enabled: true,
    minScore: 501,
    maxScore: 1000,
    benefits: ['discount', 'free_shipping', 'priority_support']
  },
  {
    name: 'پلاتینیوم',
    color: 'bg-blue-400',
    enabled: true,
    minScore: 1001,
    maxScore: 2000,
    benefits: ['discount', 'free_shipping', 'priority_support', 'early_access']
  },
  {
    name: 'الماس',
    color: 'bg-purple-400',
    enabled: true,
    minScore: 2001,
    maxScore: 9999,
    benefits: ['discount', 'free_shipping', 'priority_support', 'early_access', 'exclusive_content']
  }
])

// Methods
const saveSettings = () => {
  const allSettings = {
    ...settings.value,
    levels: levelSettings.value
  }
  emit('saveSettings', allSettings)
}

const resetSettings = () => {
  // بازنشانی به تنظیمات پیش‌فرض
  if (props.initialSettings) {
    settings.value = { ...props.initialSettings }
  }
}

// Lifecycle
onMounted(() => {
  if (props.initialSettings) {
    settings.value = { ...settings.value, ...props.initialSettings }
  }
})
</script> 
