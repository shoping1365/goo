<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">هشدارهای خطا</h4>
        <p class="text-sm text-gray-600 mt-1">مدیریت و تنظیم هشدارهای خطاهای اتصال حسابداری</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
          @click="testAlerts"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          تست هشدارها
        </button>
        <button
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="saveAlertSettings"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
          </svg>
          ذخیره تنظیمات
        </button>
      </div>
    </div>

    <!-- آمار هشدارها -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="bg-red-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-red-600">{{ alertStats.activeAlerts }}</div>
            <div class="text-sm text-red-700">هشدارهای فعال</div>
          </div>
          <div class="w-10 h-10 bg-red-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-yellow-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-yellow-600">{{ alertStats.todayAlerts }}</div>
            <div class="text-sm text-yellow-700">هشدارهای امروز</div>
          </div>
          <div class="w-10 h-10 bg-yellow-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-blue-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-blue-600">{{ alertStats.alertChannels }}</div>
            <div class="text-sm text-blue-700">کانال‌های هشدار</div>
          </div>
          <div class="w-10 h-10 bg-blue-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-5 5v-5zM4 19h6v-2H4v2zM4 15h6v-2H4v2zM4 11h6V9H4v2zM4 7h6V5H4v2zM10 7h10V5H10v2zM10 11h10V9H10v2zM10 15h10v-2H10v2zM10 19h10v-2H10v2z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-green-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-green-600">{{ alertStats.responseRate }}</div>
            <div class="text-sm text-green-700">نرخ پاسخ</div>
          </div>
          <div class="w-10 h-10 bg-green-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات هشدارها -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- انواع هشدارها -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h5 class="text-md font-medium text-gray-900 mb-4">انواع هشدارها</h5>
        <div class="space-y-4">
          <div>
            <label class="flex items-center">
              <input
                v-model="alertSettings.connectionAlerts"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">هشدارهای اتصال</span>
            </label>
            <p class="text-xs text-gray-500 mt-1">هشدار در صورت قطع شدن اتصال به نرم‌افزارهای حسابداری</p>
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="alertSettings.syncAlerts"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">هشدارهای همگام‌سازی</span>
            </label>
            <p class="text-xs text-gray-500 mt-1">هشدار در صورت خطا در فرآیند همگام‌سازی</p>
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="alertSettings.securityAlerts"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">هشدارهای امنیتی</span>
            </label>
            <p class="text-xs text-gray-500 mt-1">هشدار در صورت تشخیص فعالیت مشکوک یا ناامن</p>
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="alertSettings.performanceAlerts"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">هشدارهای عملکرد</span>
            </label>
            <p class="text-xs text-gray-500 mt-1">هشدار در صورت کند شدن یا مشکل در عملکرد</p>
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="alertSettings.dataAlerts"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">هشدارهای داده</span>
            </label>
            <p class="text-xs text-gray-500 mt-1">هشدار در صورت خطا در داده‌ها یا عدم تطابق</p>
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="alertSettings.systemAlerts"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">هشدارهای سیستم</span>
            </label>
            <p class="text-xs text-gray-500 mt-1">هشدار در صورت مشکل در سیستم یا سرویس‌ها</p>
          </div>
        </div>
      </div>

      <!-- کانال‌های هشدار -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h5 class="text-md font-medium text-gray-900 mb-4">کانال‌های هشدار</h5>
        <div class="space-y-4">
          <div>
            <label class="flex items-center">
              <input
                v-model="alertSettings.emailAlerts"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">هشدارهای ایمیل</span>
            </label>
            <div v-if="alertSettings.emailAlerts" class="mt-2">
              <input
                v-model="alertSettings.emailAddresses"
                type="text"
                class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="آدرس‌های ایمیل (جدا شده با کاما)"
              />
            </div>
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="alertSettings.smsAlerts"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">هشدارهای پیامک</span>
            </label>
            <div v-if="alertSettings.smsAlerts" class="mt-2">
              <input
                v-model="alertSettings.phoneNumbers"
                type="text"
                class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="شماره‌های تلفن (جدا شده با کاما)"
              />
            </div>
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="alertSettings.pushAlerts"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">هشدارهای فوری</span>
            </label>
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="alertSettings.webhookAlerts"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">هشدارهای Webhook</span>
            </label>
            <div v-if="alertSettings.webhookAlerts" class="mt-2">
              <input
                v-model="alertSettings.webhookUrl"
                type="url"
                class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="آدرس Webhook"
              />
            </div>
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="alertSettings.slackAlerts"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">هشدارهای Slack</span>
            </label>
            <div v-if="alertSettings.slackAlerts" class="mt-2">
              <input
                v-model="alertSettings.slackWebhook"
                type="url"
                class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="آدرس Webhook Slack"
              />
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات پیشرفته هشدار -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- تنظیمات زمان‌بندی -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h5 class="text-md font-medium text-gray-900 mb-4">تنظیمات زمان‌بندی</h5>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">فاصله هشدار (دقیقه)</label>
            <input
              v-model.number="alertSettings.alertInterval"
              type="number"
              min="1"
              max="1440"
              class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
            <p class="text-xs text-gray-500 mt-1">حداقل فاصله بین هشدارهای مشابه</p>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">ساعات کاری</label>
            <div class="grid grid-cols-2 gap-2">
              <input
                v-model="alertSettings.workStartTime"
                type="time"
                class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
              <input
                v-model="alertSettings.workEndTime"
                type="time"
                class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="alertSettings.workHoursOnly"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">فقط در ساعات کاری</span>
            </label>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">روزهای هفته</label>
            <div class="grid grid-cols-7 gap-1">
              <label v-for="day in weekDays" :key="day.value" class="flex items-center">
                <input
                  v-model="alertSettings.workDays"
                  :value="day.value"
                  type="checkbox"
                  class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                />
                <span class="mr-1 text-xs text-gray-700">{{ day.label }}</span>
              </label>
            </div>
          </div>
        </div>
      </div>

      <!-- تنظیمات پیشرفته -->
      <div class="bg-white rounded-lg border border-gray-200 p-6">
        <h5 class="text-md font-medium text-gray-900 mb-4">تنظیمات پیشرفته</h5>
        <div class="space-y-4">
          <div>
            <label class="flex items-center">
              <input
                v-model="alertSettings.escalationAlerts"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">هشدارهای تشدید</span>
            </label>
            <p class="text-xs text-gray-500 mt-1">ارسال هشدار به مدیران در صورت عدم پاسخ</p>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر هشدار روزانه</label>
            <input
              v-model.number="alertSettings.maxDailyAlerts"
              type="number"
              min="1"
              max="100"
              class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="alertSettings.groupAlerts"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">گروه‌بندی هشدارها</span>
            </label>
            <p class="text-xs text-gray-500 mt-1">ارسال هشدارهای مشابه در یک پیام</p>
          </div>

          <div>
            <label class="flex items-center">
              <input
                v-model="alertSettings.retryAlerts"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
              <span class="mr-2 text-sm text-gray-700">تلاش مجدد ارسال</span>
            </label>
            <p class="text-xs text-gray-500 mt-1">تلاش مجدد در صورت عدم ارسال موفق هشدار</p>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تعداد تلاش‌های مجدد</label>
            <input
              v-model.number="alertSettings.retryCount"
              type="number"
              min="1"
              max="5"
              class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
          </div>
        </div>
      </div>
    </div>

    <!-- هشدارهای فعال -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">هشدارهای فعال</h5>
      </div>
      <div class="overflow-x-auto">
        <table class="w-full text-sm">
          <thead>
            <tr class="border-b border-gray-200 bg-gray-50">
              <th class="text-right py-3 px-4 font-medium text-gray-600">نوع هشدار</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">پیام</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">اولویت</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">زمان ارسال</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">وضعیت</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">عملیات</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="alert in activeAlerts" :key="alert.id" class="border-b border-gray-100 hover:bg-gray-50">
              <td class="py-3 px-4">
                <div class="flex items-center space-x-3 space-x-reverse">
                  <div
                    class="w-3 h-3 rounded-full"
                    :class="getAlertTypeColor(alert.type)"
                  ></div>
                  <span class="font-medium text-gray-900">{{ alert.type }}</span>
                </div>
              </td>
              <td class="py-3 px-4">
                <div class="text-sm text-gray-900">{{ alert.message }}</div>
                <div class="text-xs text-gray-500">{{ alert.details }}</div>
              </td>
              <td class="py-3 px-4">
                <span
                  class="px-2 py-1 rounded-full text-xs font-medium"
                  :class="getPriorityClass(alert.priority)"
                >
                  {{ getPriorityLabel(alert.priority) }}
                </span>
              </td>
              <td class="py-3 px-4">
                <div class="text-sm text-gray-900">{{ alert.sentDate }}</div>
                <div class="text-xs text-gray-500">{{ alert.sentTime }}</div>
              </td>
              <td class="py-3 px-4">
                <span
                  class="px-2 py-1 rounded-full text-xs font-medium"
                  :class="getStatusClass(alert.status)"
                >
                  {{ getStatusLabel(alert.status) }}
                </span>
              </td>
              <td class="py-3 px-4">
                <div class="flex items-center space-x-2 space-x-reverse">
                  <button
                    class="p-1 text-green-600 hover:text-green-800"
                    title="تأیید"
                    @click="acknowledgeAlert(alert)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                    </svg>
                  </button>
                  <button
                    class="p-1 text-yellow-600 hover:text-yellow-800"
                    title="رد"
                    @click="dismissAlert(alert)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                    </svg>
                  </button>
                  <button
                    class="p-1 text-blue-600 hover:text-blue-800"
                    title="جزئیات"
                    @click="viewAlertDetails(alert)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// آمار هشدارها
const alertStats = ref({
  activeAlerts: 5,
  todayAlerts: 12,
  alertChannels: 4,
  responseRate: 85.5
})

// تنظیمات هشدارها
const alertSettings = ref({
  connectionAlerts: true,
  syncAlerts: true,
  securityAlerts: true,
  performanceAlerts: true,
  dataAlerts: true,
  systemAlerts: false,
  emailAlerts: true,
  emailAddresses: 'admin@example.com, manager@example.com',
  smsAlerts: false,
  phoneNumbers: '',
  pushAlerts: true,
  webhookAlerts: false,
  webhookUrl: '',
  slackAlerts: false,
  slackWebhook: '',
  alertInterval: 15,
  workStartTime: '08:00',
  workEndTime: '18:00',
  workHoursOnly: false,
  workDays: ['sat', 'sun', 'mon', 'tue', 'wed', 'thu'],
  escalationAlerts: true,
  maxDailyAlerts: 50,
  groupAlerts: true,
  retryAlerts: true,
  retryCount: 3
})

// روزهای هفته
const weekDays = ref([
  { value: 'sat', label: 'ش' },
  { value: 'sun', label: 'ی' },
  { value: 'mon', label: 'د' },
  { value: 'tue', label: 'س' },
  { value: 'wed', label: 'چ' },
  { value: 'thu', label: 'پ' },
  { value: 'fri', label: 'ج' }
])

// هشدارهای فعال
const activeAlerts = ref([
  {
    id: 1,
    type: 'خطای اتصال',
    message: 'اتصال به نرم‌افزار هلو قطع شد',
    details: 'Timeout در اتصال به سرور هلو',
    priority: 'high',
    sentDate: 'امروز',
    sentTime: '۱۴:۳۰',
    status: 'sent'
  },
  {
    id: 2,
    type: 'خطای همگام‌سازی',
    message: 'خطا در همگام‌سازی تراکنش‌ها',
    details: 'خطای فرمت در داده‌های ارسالی',
    priority: 'medium',
    sentDate: 'امروز',
    sentTime: '۱۳:۱۵',
    status: 'pending'
  },
  {
    id: 3,
    type: 'هشدار امنیتی',
    message: 'تلاش ورود ناموفق',
    details: '۳ تلاش ورود ناموفق از IP مشکوک',
    priority: 'high',
    sentDate: 'دیروز',
    sentTime: '۱۸:۴۵',
    status: 'acknowledged'
  }
])

// رنگ نوع هشدار
const getAlertTypeColor = (type: string) => {
  const colors = {
    'خطای اتصال': 'bg-red-500',
    'خطای همگام‌سازی': 'bg-yellow-500',
    'هشدار امنیتی': 'bg-red-600',
    'خطای عملکرد': 'bg-orange-500',
    'خطای داده': 'bg-blue-500',
    'خطای سیستم': 'bg-purple-500'
  }
  return colors[type] || 'bg-gray-500'
}

// کلاس اولویت
const getPriorityClass = (priority: string) => {
  const classes = {
    high: 'bg-red-100 text-red-700',
    medium: 'bg-yellow-100 text-yellow-700',
    low: 'bg-green-100 text-green-700'
  }
  return classes[priority] || 'bg-gray-100 text-gray-700'
}

// برچسب اولویت
const getPriorityLabel = (priority: string) => {
  const labels = {
    high: 'زیاد',
    medium: 'متوسط',
    low: 'کم'
  }
  return labels[priority] || priority
}

// کلاس وضعیت
const getStatusClass = (status: string) => {
  const classes = {
    sent: 'bg-blue-100 text-blue-700',
    pending: 'bg-yellow-100 text-yellow-700',
    acknowledged: 'bg-green-100 text-green-700',
    dismissed: 'bg-gray-100 text-gray-700'
  }
  return classes[status] || 'bg-gray-100 text-gray-700'
}

// برچسب وضعیت
const getStatusLabel = (status: string) => {
  const labels = {
    sent: 'ارسال شده',
    pending: 'در انتظار',
    acknowledged: 'تأیید شده',
    dismissed: 'رد شده'
  }
  return labels[status] || status
}

// تست هشدارها
const testAlerts = () => {
  // TODO: تست هشدارها

}

// ذخیره تنظیمات هشدار
const saveAlertSettings = () => {
  // TODO: ذخیره تنظیمات هشدار

}

interface Alert {
  id?: number | string
  [key: string]: unknown
}

// تأیید هشدار
const acknowledgeAlert = (_alert: Alert) => {
  // TODO: تأیید هشدار

}

// رد هشدار
const dismissAlert = (_alert: Alert) => {
  // TODO: رد هشدار

}

// مشاهده جزئیات هشدار
const viewAlertDetails = (_alert: Alert) => {
  // TODO: نمایش جزئیات هشدار

}
</script>

<!--
  کامپوننت هشدارهای خطا
  شامل:
  1. آمار هشدارها
  2. تنظیمات انواع هشدارها
  3. کانال‌های هشدار
  4. تنظیمات زمان‌بندی
  5. تنظیمات پیشرفته
  6. جدول هشدارهای فعال
  7. طراحی مدرن و کاملاً ریسپانسیو
--> 
