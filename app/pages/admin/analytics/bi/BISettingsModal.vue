<template>
  <div v-if="isOpen" class="fixed inset-0 z-50 overflow-y-auto">
    <div class="flex items-center justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
      <!-- Background overlay -->
      <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" @click="$emit('close')"></div>

      <!-- Modal panel -->
      <div class="inline-block align-bottom bg-white rounded-lg text-right overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-4xl sm:w-full">
        <div class="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
          <!-- Header -->
          <div class="flex items-center justify-between mb-6">
            <h3 class="text-lg font-semibold text-gray-900">تنظیمات هوش تجاری</h3>
            <button class="text-gray-400 hover:text-gray-600" @click="$emit('close')">
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
                :class="[
                  activeTab === tab.id
                    ? 'border-blue-500 text-blue-600'
                    : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300',
                  'whitespace-nowrap py-2 px-1 border-b-2 font-medium text-sm transition-colors'
                ]"
                @click="activeTab = tab.id"
              >
                {{ tab.title }}
              </button>
            </nav>
          </div>

          <!-- Tab Content -->
          <div class="space-y-6">
            <!-- General Settings -->
            <div v-if="activeTab === 'general'">
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">بازه زمانی پیش‌فرض</label>
                  <select v-model="settings.defaultTimeRange" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                    <option value="today">امروز</option>
                    <option value="week">هفته اخیر</option>
                    <option value="month">ماه اخیر</option>
                    <option value="quarter">سه ماهه</option>
                  </select>
                </div>

                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">زبان گزارش‌ها</label>
                  <select v-model="settings.language" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                    <option value="fa">فارسی</option>
                    <option value="en">انگلیسی</option>
                  </select>
                </div>

                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">واحد پول</label>
                  <select v-model="settings.currency" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                    <option value="IRR">ریال</option>
                    <option value="IRT">تومان</option>
                    <option value="USD">دلار</option>
                  </select>
                </div>

                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">منطقه زمانی</label>
                  <select v-model="settings.timezone" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                    <option value="Asia/Tehran">تهران (GMT+3:30)</option>
                    <option value="UTC">UTC</option>
                  </select>
                </div>
              </div>
            </div>

            <!-- Alerts & Notifications -->
            <div v-if="activeTab === 'alerts'">
              <div class="space-y-6">
                <div>
                  <h4 class="text-md font-medium text-gray-900 mb-4">هشدارهای خودکار</h4>
                  <div class="space-y-4">
                    <div class="flex items-center justify-between p-6 border border-gray-200 rounded-lg">
                      <div>
                        <h5 class="font-medium text-gray-900">افت نرخ تبدیل</h5>
                        <p class="text-sm text-gray-500">هشدار در صورت افت بیش از ۲۰٪ نرخ تبدیل</p>
                      </div>
                      <label class="relative inline-flex items-center cursor-pointer">
                        <input v-model="settings.alerts.conversionDrop" type="checkbox" class="sr-only peer">
                        <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
                      </label>
                    </div>

                    <div class="flex items-center justify-between p-6 border border-gray-200 rounded-lg">
                      <div>
                        <h5 class="font-medium text-gray-900">کالاهای کم‌موجود</h5>
                        <p class="text-sm text-gray-500">هشدار برای کالاهایی که موجودی آن‌ها کمتر از ۱۰ عدد است</p>
                      </div>
                      <label class="relative inline-flex items-center cursor-pointer">
                        <input v-model="settings.alerts.lowStock" type="checkbox" class="sr-only peer">
                        <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
                      </label>
                    </div>

                    <div class="flex items-center justify-between p-6 border border-gray-200 rounded-lg">
                      <div>
                        <h5 class="font-medium text-gray-900">افزایش نرخ ریزش</h5>
                        <p class="text-sm text-gray-500">هشدار در صورت افزایش بیش از ۱۰٪ نرخ ریزش مشتریان</p>
                      </div>
                      <label class="relative inline-flex items-center cursor-pointer">
                        <input v-model="settings.alerts.churnIncrease" type="checkbox" class="sr-only peer">
                        <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
                      </label>
                    </div>
                  </div>
                </div>

                <div>
                  <h4 class="text-md font-medium text-gray-900 mb-4">کانال‌های اطلاع‌رسانی</h4>
                  <div class="space-y-3">
                    <label class="flex items-center">
                      <input v-model="settings.notifications.email" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                      <span class="mr-2 text-sm text-gray-700">ایمیل</span>
                    </label>
                    <label class="flex items-center">
                      <input v-model="settings.notifications.sms" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                      <span class="mr-2 text-sm text-gray-700">پیامک</span>
                    </label>
                    <label class="flex items-center">
                      <input v-model="settings.notifications.push" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                      <span class="mr-2 text-sm text-gray-700">اعلان مرورگر</span>
                    </label>
                  </div>
                </div>
              </div>
            </div>

            <!-- Automated Reports -->
            <div v-if="activeTab === 'reports'">
              <div class="space-y-6">
                <div>
                  <h4 class="text-md font-medium text-gray-900 mb-4">گزارش‌های خودکار</h4>
                  <div class="space-y-4">
                    <div v-for="report in automatedReports" :key="report.id" class="border border-gray-200 rounded-lg p-6">
                      <div class="flex items-center justify-between mb-3">
                        <div>
                          <h5 class="font-medium text-gray-900">{{ report.title }}</h5>
                          <p class="text-sm text-gray-500">{{ report.description }}</p>
                        </div>
                        <label class="relative inline-flex items-center cursor-pointer">
                          <input v-model="report.enabled" type="checkbox" class="sr-only peer">
                          <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
                        </label>
                      </div>
                      
                      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                        <div>
                          <label class="block text-sm font-medium text-gray-700 mb-1">فرکانس</label>
                          <select v-model="report.frequency" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                            <option value="daily">روزانه</option>
                            <option value="weekly">هفتگی</option>
                            <option value="monthly">ماهانه</option>
                          </select>
                        </div>
                        
                        <div>
                          <label class="block text-sm font-medium text-gray-700 mb-1">زمان ارسال</label>
                          <input v-model="report.time" type="time" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                        </div>
                      </div>
                      
                      <div class="mt-3">
                        <label class="block text-sm font-medium text-gray-700 mb-1">گیرندگان</label>
                        <input v-model="report.recipients" type="email" placeholder="ایمیل‌ها را با کاما جدا کنید" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <!-- Data Sources -->
            <div v-if="activeTab === 'data'">
              <div class="space-y-6">
                <div>
                  <h4 class="text-md font-medium text-gray-900 mb-4">منابع داده</h4>
                  <div class="space-y-4">
                    <div class="flex items-center justify-between p-6 border border-gray-200 rounded-lg">
                      <div>
                        <h5 class="font-medium text-gray-900">دیتابیس اصلی</h5>
                        <p class="text-sm text-gray-500">اتصال به دیتابیس MySQL</p>
                      </div>
                      <span class="text-green-600 text-sm font-medium">متصل</span>
                    </div>

                    <div class="flex items-center justify-between p-6 border border-gray-200 rounded-lg">
                      <div>
                        <h5 class="font-medium text-gray-900">Redis Cache</h5>
                        <p class="text-sm text-gray-500">کش داده‌ها برای عملکرد بهتر</p>
                      </div>
                      <span class="text-green-600 text-sm font-medium">متصل</span>
                    </div>

                    <div class="flex items-center justify-between p-6 border border-gray-200 rounded-lg">
                      <div>
                        <h5 class="font-medium text-gray-900">Google Analytics</h5>
                        <p class="text-sm text-gray-500">داده‌های ترافیک و رفتار کاربران</p>
                      </div>
                      <button class="text-blue-600 hover:text-blue-800 text-sm font-medium">اتصال</button>
                    </div>
                  </div>
                </div>

                <div>
                  <h4 class="text-md font-medium text-gray-900 mb-4">تنظیمات به‌روزرسانی</h4>
                  <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                    <div>
                      <label class="block text-sm font-medium text-gray-700 mb-2">فرکانس به‌روزرسانی</label>
                      <select v-model="settings.dataRefresh" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                        <option value="5">هر ۵ دقیقه</option>
                        <option value="15">هر ۱۵ دقیقه</option>
                        <option value="30">هر ۳۰ دقیقه</option>
                        <option value="60">هر ساعت</option>
                      </select>
                    </div>

                    <div>
                      <label class="block text-sm font-medium text-gray-700 mb-2">حفظ داده‌ها</label>
                      <select v-model="settings.dataRetention" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                        <option value="30">۳۰ روز</option>
                        <option value="90">۹۰ روز</option>
                        <option value="365">۱ سال</option>
                        <option value="forever">برای همیشه</option>
                      </select>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Footer -->
        <div class="bg-gray-50 px-4 py-3 sm:px-6 sm:flex sm:flex-row-reverse">
          <button class="w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-blue-600 text-base font-medium text-white hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 sm:mr-3 sm:w-auto sm:text-sm" @click="saveSettings">
            ذخیره تنظیمات
          </button>
          <button class="mt-3 w-full inline-flex justify-center rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 sm:mt-0 sm:mr-3 sm:w-auto sm:text-sm" @click="$emit('close')">
            انصراف
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useAuth } from '~/composables/useAuth';

// تعریف definePageMeta برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user: _user, hasPermission: _hasPermission } = useAuth()

// Props
defineProps<{
  isOpen: boolean
}>()

// Emits
const emit = defineEmits<{
  close: []
}>()

// Reactive data
const activeTab = ref('general')

const tabs = [
  { id: 'general', title: 'عمومی' },
  { id: 'alerts', title: 'هشدارها' },
  { id: 'reports', title: 'گزارش‌ها' },
  { id: 'data', title: 'داده‌ها' }
]

const settings = ref({
  defaultTimeRange: 'month',
  language: 'fa',
  currency: 'IRT',
  timezone: 'Asia/Tehran',
  dataRefresh: 15,
  dataRetention: 90,
  alerts: {
    conversionDrop: true,
    lowStock: true,
    churnIncrease: false
  },
  notifications: {
    email: true,
    sms: false,
    push: true
  }
})

const automatedReports = ref([
  {
    id: 1,
    title: 'گزارش فروش روزانه',
    description: 'خلاصه‌ای از فروش و عملکرد روزانه',
    enabled: true,
    frequency: 'daily',
    time: '08:00',
    recipients: 'manager@example.com, sales@example.com'
  },
  {
    id: 2,
    title: 'گزارش موجودی هفتگی',
    description: 'وضعیت موجودی و کالاهای کم‌موجود',
    enabled: true,
    frequency: 'weekly',
    time: '09:00',
    recipients: 'inventory@example.com'
  },
  {
    id: 3,
    title: 'گزارش مشتریان ماهانه',
    description: 'تحلیل رفتار و نرخ ریزش مشتریان',
    enabled: false,
    frequency: 'monthly',
    time: '10:00',
    recipients: 'marketing@example.com'
  }
])

// Methods
const saveSettings = () => {

  // Implementation for saving settings
  emit('close')
}
</script> 
