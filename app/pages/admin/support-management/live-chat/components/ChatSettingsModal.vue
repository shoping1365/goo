<template>
  <div v-if="isOpen" class="fixed inset-0 z-50 overflow-y-auto" dir="rtl">
    <div class="flex items-center justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
      <!-- Background overlay -->
      <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" @click="$emit('close')"></div>

      <!-- Modal panel -->
      <div class="inline-block align-bottom bg-white rounded-2xl text-right overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-2xl sm:w-full">
        <div class="bg-white px-6 py-6">
          <!-- Header -->
          <div class="flex items-center justify-between mb-6">
            <h3 class="text-2xl font-bold text-gray-900">تنظیمات چت</h3>
            <button 
              @click="$emit('close')"
              class="text-gray-400 hover:text-gray-600 transition-colors"
            >
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
              </svg>
            </button>
          </div>

          <!-- Settings Tabs -->
          <div class="border-b border-gray-200 mb-6">
            <nav class="flex space-x-8 space-x-reverse">
              <button
                v-for="tab in tabs"
                :key="tab.id"
                @click="activeTab = tab.id"
                :class="[
                  activeTab === tab.id
                    ? 'border-blue-500 text-blue-600'
                    : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300',
                  'whitespace-nowrap py-2 px-1 border-b-2 font-medium text-sm transition-colors'
                ]"
              >
                {{ tab.title }}
              </button>
            </nav>
          </div>

          <!-- General Settings -->
          <div v-if="activeTab === 'general'" class="space-y-6">
            <div>
              <h4 class="text-lg font-medium text-gray-900 mb-4">تنظیمات عمومی</h4>
              <div class="space-y-4">
                <div class="flex items-center justify-between">
                  <div>
                    <label class="text-sm font-medium text-gray-700">وضعیت آنلاین</label>
                    <p class="text-sm text-gray-500">نمایش وضعیت آنلاین شما به کاربران</p>
                  </div>
                  <button 
                    @click="settings.onlineStatus = !settings.onlineStatus"
                    :class="[
                      settings.onlineStatus ? 'bg-blue-600' : 'bg-gray-200',
                      'relative inline-flex h-6 w-11 items-center rounded-full transition-colors'
                    ]"
                  >
                    <span 
                      :class="[
                        settings.onlineStatus ? 'translate-x-6' : 'translate-x-1',
                        'inline-block h-4 w-4 transform rounded-full bg-white transition-transform'
                      ]"
                    ></span>
                  </button>
                </div>

                <div class="flex items-center justify-between">
                  <div>
                    <label class="text-sm font-medium text-gray-700">اعلان‌های صوتی</label>
                    <p class="text-sm text-gray-500">پخش صدای اعلان برای پیام‌های جدید</p>
                  </div>
                  <button 
                    @click="settings.soundNotifications = !settings.soundNotifications"
                    :class="[
                      settings.soundNotifications ? 'bg-blue-600' : 'bg-gray-200',
                      'relative inline-flex h-6 w-11 items-center rounded-full transition-colors'
                    ]"
                  >
                    <span 
                      :class="[
                        settings.soundNotifications ? 'translate-x-6' : 'translate-x-1',
                        'inline-block h-4 w-4 transform rounded-full bg-white transition-transform'
                      ]"
                    ></span>
                  </button>
                </div>

                <div class="flex items-center justify-between">
                  <div>
                    <label class="text-sm font-medium text-gray-700">لرزش</label>
                    <p class="text-sm text-gray-500">لرزش دستگاه برای پیام‌های جدید</p>
                  </div>
                  <button 
                    @click="settings.vibration = !settings.vibration"
                    :class="[
                      settings.vibration ? 'bg-blue-600' : 'bg-gray-200',
                      'relative inline-flex h-6 w-11 items-center rounded-full transition-colors'
                    ]"
                  >
                    <span 
                      :class="[
                        settings.vibration ? 'translate-x-6' : 'translate-x-1',
                        'inline-block h-4 w-4 transform rounded-full bg-white transition-transform'
                      ]"
                    ></span>
                  </button>
                </div>

                <div>
                  <label class="text-sm font-medium text-gray-700">زبان</label>
                  <select v-model="settings.language" class="mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm rounded-md">
                    <option value="fa">فارسی</option>
                    <option value="en">English</option>
                    <option value="ar">العربية</option>
                  </select>
                </div>
              </div>
            </div>
          </div>

          <!-- Notifications -->
          <div v-if="activeTab === 'notifications'" class="space-y-6">
            <div>
              <h4 class="text-lg font-medium text-gray-900 mb-4">تنظیمات اعلان‌ها</h4>
              <div class="space-y-4">
                <div class="flex items-center justify-between">
                  <div>
                    <label class="text-sm font-medium text-gray-700">اعلان‌های ایمیل</label>
                    <p class="text-sm text-gray-500">ارسال اعلان‌ها به ایمیل شما</p>
                  </div>
                  <button 
                    @click="settings.emailNotifications = !settings.emailNotifications"
                    :class="[
                      settings.emailNotifications ? 'bg-blue-600' : 'bg-gray-200',
                      'relative inline-flex h-6 w-11 items-center rounded-full transition-colors'
                    ]"
                  >
                    <span 
                      :class="[
                        settings.emailNotifications ? 'translate-x-6' : 'translate-x-1',
                        'inline-block h-4 w-4 transform rounded-full bg-white transition-transform'
                      ]"
                    ></span>
                  </button>
                </div>

                <div class="flex items-center justify-between">
                  <div>
                    <label class="text-sm font-medium text-gray-700">اعلان‌های Push</label>
                    <p class="text-sm text-gray-500">اعلان‌های Push برای پیام‌های جدید</p>
                  </div>
                  <button 
                    @click="settings.pushNotifications = !settings.pushNotifications"
                    :class="[
                      settings.pushNotifications ? 'bg-blue-600' : 'bg-gray-200',
                      'relative inline-flex h-6 w-11 items-center rounded-full transition-colors'
                    ]"
                  >
                    <span 
                      :class="[
                        settings.pushNotifications ? 'translate-x-6' : 'translate-x-1',
                        'inline-block h-4 w-4 transform rounded-full bg-white transition-transform'
                      ]"
                    ></span>
                  </button>
                </div>

                <div>
                  <label class="text-sm font-medium text-gray-700">ساعات کاری</label>
                  <div class="mt-2 grid grid-cols-2 gap-6">
                    <div>
                      <label class="text-xs text-gray-500">شروع</label>
                      <input 
                        v-model="settings.workHours.start" 
                        type="time" 
                        class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                      >
                    </div>
                    <div>
                      <label class="text-xs text-gray-500">پایان</label>
                      <input 
                        v-model="settings.workHours.end" 
                        type="time" 
                        class="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                      >
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Appearance -->
          <div v-if="activeTab === 'appearance'" class="space-y-6">
            <div>
              <h4 class="text-lg font-medium text-gray-900 mb-4">ظاهر و تم</h4>
              <div class="space-y-4">
                <div>
                  <label class="text-sm font-medium text-gray-700">تم</label>
                  <div class="mt-2 grid grid-cols-3 gap-3">
                    <button 
                      @click="settings.theme = 'light'"
                      :class="[
                        settings.theme === 'light' ? 'ring-2 ring-blue-500' : '',
                        'relative p-6 bg-white border border-gray-300 rounded-lg hover:bg-gray-50'
                      ]"
                    >
                      <div class="w-full h-8 bg-gray-100 rounded"></div>
                      <div class="mt-2 text-sm font-medium text-gray-900">روشن</div>
                    </button>
                    <button 
                      @click="settings.theme = 'dark'"
                      :class="[
                        settings.theme === 'dark' ? 'ring-2 ring-blue-500' : '',
                        'relative p-6 bg-gray-800 border border-gray-600 rounded-lg hover:bg-gray-700'
                      ]"
                    >
                      <div class="w-full h-8 bg-gray-600 rounded"></div>
                      <div class="mt-2 text-sm font-medium text-white">تاریک</div>
                    </button>
                    <button 
                      @click="settings.theme = 'auto'"
                      :class="[
                        settings.theme === 'auto' ? 'ring-2 ring-blue-500' : '',
                        'relative p-6 bg-gradient-to-r from-gray-100 to-gray-800 border border-gray-300 rounded-lg hover:bg-gray-50'
                      ]"
                    >
                      <div class="w-full h-8 bg-gradient-to-r from-gray-200 to-gray-600 rounded"></div>
                      <div class="mt-2 text-sm font-medium text-gray-900">خودکار</div>
                    </button>
                  </div>
                </div>

                <div>
                  <label class="text-sm font-medium text-gray-700">اندازه فونت</label>
                  <select v-model="settings.fontSize" class="mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm rounded-md">
                    <option value="small">کوچک</option>
                    <option value="medium">متوسط</option>
                    <option value="large">بزرگ</option>
                  </select>
                </div>

                <div>
                  <label class="text-sm font-medium text-gray-700">رنگ اصلی</label>
                  <div class="mt-2 flex space-x-2 space-x-reverse">
                    <button 
                      v-for="color in themeColors"
                      :key="color.name"
                      @click="settings.primaryColor = color.value"
                      :class="[
                        settings.primaryColor === color.value ? 'ring-2 ring-offset-2 ring-blue-500' : '',
                        'w-8 h-8 rounded-full border-2 border-white shadow-sm'
                      ]"
                      :style="{ backgroundColor: color.value }"
                    ></button>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Privacy -->
          <div v-if="activeTab === 'privacy'" class="space-y-6">
            <div>
              <h4 class="text-lg font-medium text-gray-900 mb-4">حریم خصوصی</h4>
              <div class="space-y-4">
                <div class="flex items-center justify-between">
                  <div>
                    <label class="text-sm font-medium text-gray-700">نمایش آخرین بازدید</label>
                    <p class="text-sm text-gray-500">نمایش زمان آخرین بازدید شما به کاربران</p>
                  </div>
                  <button 
                    @click="settings.showLastSeen = !settings.showLastSeen"
                    :class="[
                      settings.showLastSeen ? 'bg-blue-600' : 'bg-gray-200',
                      'relative inline-flex h-6 w-11 items-center rounded-full transition-colors'
                    ]"
                  >
                    <span 
                      :class="[
                        settings.showLastSeen ? 'translate-x-6' : 'translate-x-1',
                        'inline-block h-4 w-4 transform rounded-full bg-white transition-transform'
                      ]"
                    ></span>
                  </button>
                </div>

                <div class="flex items-center justify-between">
                  <div>
                    <label class="text-sm font-medium text-gray-700">نمایش وضعیت تایپ</label>
                    <p class="text-sm text-gray-500">نمایش وضعیت تایپ کردن شما</p>
                  </div>
                  <button 
                    @click="settings.showTypingStatus = !settings.showTypingStatus"
                    :class="[
                      settings.showTypingStatus ? 'bg-blue-600' : 'bg-gray-200',
                      'relative inline-flex h-6 w-11 items-center rounded-full transition-colors'
                    ]"
                  >
                    <span 
                      :class="[
                        settings.showTypingStatus ? 'translate-x-6' : 'translate-x-1',
                        'inline-block h-4 w-4 transform rounded-full bg-white transition-transform'
                      ]"
                    ></span>
                  </button>
                </div>

                <div class="flex items-center justify-between">
                  <div>
                    <label class="text-sm font-medium text-gray-700">ذخیره تاریخچه</label>
                    <p class="text-sm text-gray-500">ذخیره تاریخچه چت‌ها</p>
                  </div>
                  <button 
                    @click="settings.saveChatHistory = !settings.saveChatHistory"
                    :class="[
                      settings.saveChatHistory ? 'bg-blue-600' : 'bg-gray-200',
                      'relative inline-flex h-6 w-11 items-center rounded-full transition-colors'
                    ]"
                  >
                    <span 
                      :class="[
                        settings.saveChatHistory ? 'translate-x-6' : 'translate-x-1',
                        'inline-block h-4 w-4 transform rounded-full bg-white transition-transform'
                      ]"
                    ></span>
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Footer -->
        <div class="bg-gray-50 px-6 py-4 flex justify-end space-x-3 space-x-reverse">
          <button 
            @click="$emit('close')"
            class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            انصراف
          </button>
          <button 
            @click="saveSettings"
            class="px-4 py-2 text-sm font-medium text-white bg-blue-600 border border-transparent rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            ذخیره تنظیمات
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const props = defineProps<{
  isOpen: boolean
}>()

const emit = defineEmits<{
  close: []
}>()

// Reactive data
const activeTab = ref('general')

// Settings
const settings = ref({
  onlineStatus: true,
  soundNotifications: true,
  vibration: false,
  language: 'fa',
  emailNotifications: true,
  pushNotifications: true,
  workHours: {
    start: '09:00',
    end: '18:00'
  },
  theme: 'light',
  fontSize: 'medium',
  primaryColor: '#3B82F6',
  showLastSeen: true,
  showTypingStatus: true,
  saveChatHistory: true
})

// Theme colors
const themeColors = ref([
  { name: 'آبی', value: '#3B82F6' },
  { name: 'سبز', value: '#10B981' },
  { name: 'بنفش', value: '#8B5CF6' },
  { name: 'قرمز', value: '#EF4444' },
  { name: 'نارنجی', value: '#F59E0B' },
  { name: 'صورتی', value: '#EC4899' }
])

// Tabs
const tabs = ref([
  { id: 'general', title: 'عمومی' },
  { id: 'notifications', title: 'اعلان‌ها' },
  { id: 'appearance', title: 'ظاهر' },
  { id: 'privacy', title: 'حریم خصوصی' }
])

// Methods
const saveSettings = () => {
  // Save settings to localStorage or API
  localStorage.setItem('chatSettings', JSON.stringify(settings.value))
  console.log('Settings saved:', settings.value)
  emit('close')
}
</script>

<style scoped>
/* Custom styles for modal */
</style> 
