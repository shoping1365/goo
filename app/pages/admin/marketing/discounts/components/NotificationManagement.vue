<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200">
    <!-- هدر بخش -->
    <div class="p-6 border-b border-gray-200">
      <div class="flex items-center justify-between">
        <div>
          <h2 class="text-lg font-semibold text-gray-900">مدیریت اعلان‌ها و یادآوری‌ها</h2>
          <p class="text-gray-600 mt-1">مدیریت اعلان‌ها، یادآوری‌ها و زمان‌بندی ارسال پیام‌ها</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors" @click="showNotificationForm = true">
            <span class="i-heroicons-plus ml-2"></span>
            افزودن اعلان جدید
          </button>
        </div>
      </div>
    </div>

    <!-- آمار اعلان‌ها -->
    <div class="p-6 border-b border-gray-200">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div class="bg-blue-50 p-6 rounded-lg">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center ml-3">
              <span class="i-heroicons-bell text-blue-600 text-xl"></span>
            </div>
            <div>
              <p class="text-sm text-blue-600">اعلان‌های فعال</p>
              <p class="text-2xl font-bold text-blue-900">{{ stats.activeNotifications }}</p>
            </div>
          </div>
        </div>
        <div class="bg-green-50 p-6 rounded-lg">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-green-100 rounded-lg flex items-center justify-center ml-3">
              <span class="i-heroicons-clock text-green-600 text-xl"></span>
            </div>
            <div>
              <p class="text-sm text-green-600">یادآوری‌های برنامه‌ریزی شده</p>
              <p class="text-2xl font-bold text-green-900">{{ stats.scheduledReminders }}</p>
            </div>
          </div>
        </div>
        <div class="bg-purple-50 p-6 rounded-lg">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-purple-100 rounded-lg flex items-center justify-center ml-3">
              <span class="i-heroicons-check-circle text-purple-600 text-xl"></span>
            </div>
            <div>
              <p class="text-sm text-purple-600">نرخ تحویل موفق</p>
              <p class="text-2xl font-bold text-purple-900">{{ stats.deliveryRate }}%</p>
            </div>
          </div>
        </div>
        <div class="bg-orange-50 p-6 rounded-lg">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-orange-100 rounded-lg flex items-center justify-center ml-3">
              <span class="i-heroicons-cursor-arrow-rays text-orange-600 text-xl"></span>
            </div>
            <div>
              <p class="text-sm text-orange-600">نرخ تعامل</p>
              <p class="text-2xl font-bold text-orange-900">{{ stats.engagementRate }}%</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تب‌های اعلان‌ها -->
    <div class="border-b border-gray-200">
      <div class="flex border-b border-gray-200 overflow-x-auto">
        <button
v-for="tab in tabs" :key="tab.value" :class="['px-6 py-3 -mb-px font-medium text-sm focus:outline-none whitespace-nowrap', activeTab === tab.value ? 'border-b-2 border-blue-600 text-blue-700' : 'text-gray-500 hover:text-blue-600']"
          @click="activeTab = tab.value">
          {{ tab.label }}
        </button>
      </div>
    </div>

    <!-- محتوای تب‌ها -->
    <div class="p-6">
      <!-- اعلان‌های فعال -->
      <div v-if="activeTab === 'notifications'" class="space-y-6">
        <div class="flex justify-between items-center">
          <div class="flex items-center space-x-4 space-x-reverse">
            <select v-model="filterType" class="px-3 py-2 border border-gray-300 rounded-lg text-sm">
              <option value="">همه انواع</option>
              <option value="email">ایمیل</option>
              <option value="sms">پیامک</option>
              <option value="push">اعلان</option>
              <option value="in-app">درون برنامه</option>
            </select>
            <select v-model="filterStatus" class="px-3 py-2 border border-gray-300 rounded-lg text-sm">
              <option value="">همه وضعیت‌ها</option>
              <option value="active">فعال</option>
              <option value="paused">متوقف</option>
              <option value="draft">پیش‌نویس</option>
            </select>
          </div>
          <div class="flex items-center space-x-2 space-x-reverse">
            <button class="px-3 py-1 bg-green-100 text-green-700 rounded text-sm hover:bg-green-200" @click="bulkAction('activate')">
              فعال کردن
            </button>
            <button class="px-3 py-1 bg-yellow-100 text-yellow-700 rounded text-sm hover:bg-yellow-200" @click="bulkAction('pause')">
              متوقف کردن
            </button>
            <button class="px-3 py-1 bg-red-100 text-red-700 rounded text-sm hover:bg-red-200" @click="bulkAction('delete')">
              حذف
            </button>
          </div>
        </div>

        <div class="space-y-4">
          <div v-for="notification in filteredNotifications" :key="notification.id" class="border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow">
            <div class="flex items-center justify-between">
              <div class="flex items-center space-x-4 space-x-reverse">
                <input v-model="selectedNotifications" type="checkbox" :value="notification.id" class="rounded border-gray-300">
                <div class="w-10 h-10 rounded-full flex items-center justify-center" :style="{ backgroundColor: notification.color + '20', color: notification.color }">
                  <span class="i-heroicons-bell text-lg"></span>
                </div>
                <div>
                  <h4 class="font-medium text-gray-900">{{ notification.title }}</h4>
                  <p class="text-sm text-gray-500">{{ notification.type }} • {{ notification.trigger }}</p>
                </div>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <span :class="['px-2 py-1 rounded-full text-xs', getStatusClass(notification.status)]">
                  {{ getStatusText(notification.status) }}
                </span>
                <button class="text-blue-600 hover:text-blue-900" @click="editNotification(notification)">
                  <span class="i-heroicons-pencil-square"></span>
                </button>
                <button class="text-green-600 hover:text-green-900" @click="duplicateNotification(notification)">
                  <span class="i-heroicons-document-duplicate"></span>
                </button>
                <button class="text-red-600 hover:text-red-900" @click="deleteNotification(notification)">
                  <span class="i-heroicons-trash"></span>
                </button>
              </div>
            </div>
            
            <div class="mt-4 grid grid-cols-1 md:grid-cols-3 gap-6 text-sm">
              <div>
                <span class="text-gray-500">آخرین ارسال:</span>
                <span class="font-medium mr-2">{{ formatDate(notification.lastSent) }}</span>
              </div>
              <div>
                <span class="text-gray-500">تعداد ارسال:</span>
                <span class="font-medium mr-2">{{ notification.sentCount }}</span>
              </div>
              <div>
                <span class="text-gray-500">نرخ باز شدن:</span>
                <span class="font-medium mr-2">{{ notification.openRate }}%</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- یادآوری‌ها -->
      <div v-if="activeTab === 'reminders'" class="space-y-6">
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <div v-for="reminder in reminders" :key="reminder.id" class="border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow">
            <div class="flex items-center justify-between mb-4">
              <div class="flex items-center">
                <div class="w-10 h-10 rounded-full flex items-center justify-center ml-3" :style="{ backgroundColor: reminder.color + '20', color: reminder.color }">
                  <span class="i-heroicons-clock text-lg"></span>
                </div>
                <div>
                  <h4 class="font-medium text-gray-900">{{ reminder.title }}</h4>
                  <p class="text-sm text-gray-500">{{ reminder.type }}</p>
                </div>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <button class="text-blue-600 hover:text-blue-900" @click="editReminder(reminder)">
                  <span class="i-heroicons-pencil-square"></span>
                </button>
                <button class="text-red-600 hover:text-red-900" @click="deleteReminder(reminder)">
                  <span class="i-heroicons-trash"></span>
                </button>
              </div>
            </div>
            
            <div class="space-y-3">
              <div class="text-sm text-gray-600">{{ reminder.description }}</div>
              <div class="flex justify-between text-sm">
                <span class="text-gray-500">زمان ارسال:</span>
                <span class="font-medium">{{ reminder.schedule }}</span>
              </div>
              <div class="flex justify-between text-sm">
                <span class="text-gray-500">آخرین اجرا:</span>
                <span class="font-medium">{{ formatDate(reminder.lastExecuted) }}</span>
              </div>
              <div class="flex justify-between text-sm">
                <span class="text-gray-500">وضعیت:</span>
                <span :class="['px-2 py-1 rounded-full text-xs', getStatusClass(reminder.status)]">
                  {{ getStatusText(reminder.status) }}
                </span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- زمان‌بندی -->
      <div v-if="activeTab === 'scheduling'" class="space-y-6">
        <div class="bg-blue-50 p-6 rounded-lg">
          <h4 class="text-md font-medium text-blue-900 mb-2">زمان‌بندی هوشمند</h4>
          <p class="text-sm text-blue-700">تنظیم زمان‌بندی بهینه برای ارسال اعلان‌ها بر اساس رفتار کاربران</p>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <div class="border border-gray-200 rounded-lg p-6">
            <h5 class="font-medium text-gray-900 mb-4">زمان‌بندی بر اساس منطقه زمانی</h5>
            <div class="space-y-4">
              <div v-for="timezone in timezoneSettings" :key="timezone.id" class="flex items-center justify-between p-3 bg-gray-50 rounded">
                <div>
                  <span class="text-sm font-medium text-gray-700">{{ timezone.name }}</span>
                  <p class="text-xs text-gray-500">{{ timezone.time }}</p>
                </div>
                <div class="flex items-center space-x-2 space-x-reverse">
                  <input v-model="timezone.time" type="time" class="px-2 py-1 border border-gray-300 rounded text-sm">
                  <button class="text-blue-600 hover:text-blue-900" @click="updateTimezone(timezone)">
                    <span class="i-heroicons-check"></span>
                  </button>
                </div>
              </div>
            </div>
          </div>

          <div class="border border-gray-200 rounded-lg p-6">
            <h5 class="font-medium text-gray-900 mb-4">زمان‌بندی بر اساس رفتار</h5>
            <div class="space-y-4">
              <div v-for="behavior in behaviorSettings" :key="behavior.id" class="flex items-center justify-between p-3 bg-gray-50 rounded">
                <div>
                  <span class="text-sm font-medium text-gray-700">{{ behavior.name }}</span>
                  <p class="text-xs text-gray-500">{{ behavior.description }}</p>
                </div>
                <div class="flex items-center space-x-2 space-x-reverse">
                  <select v-model="behavior.delay" class="px-2 py-1 border border-gray-300 rounded text-sm">
                    <option value="1">1 ساعت</option>
                    <option value="3">3 ساعت</option>
                    <option value="6">6 ساعت</option>
                    <option value="24">1 روز</option>
                  </select>
                  <button class="text-blue-600 hover:text-blue-900" @click="updateBehavior(behavior)">
                    <span class="i-heroicons-check"></span>
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- قالب‌های اعلان -->
      <div v-if="activeTab === 'templates'" class="space-y-6">
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <div v-for="template in notificationTemplates" :key="template.id" class="border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow">
            <div class="flex items-center justify-between mb-4">
              <div class="flex items-center">
                <div class="w-10 h-10 rounded-full flex items-center justify-center ml-3" :style="{ backgroundColor: template.color + '20', color: template.color }">
                  <span class="i-heroicons-document-text text-lg"></span>
                </div>
                <div>
                  <h4 class="font-medium text-gray-900">{{ template.name }}</h4>
                  <p class="text-sm text-gray-500">{{ template.category }}</p>
                </div>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <button class="text-green-600 hover:text-green-900" @click="useTemplate(template)">
                  <span class="i-heroicons-plus"></span>
                </button>
                <button class="text-blue-600 hover:text-blue-900" @click="editTemplate(template)">
                  <span class="i-heroicons-pencil-square"></span>
                </button>
              </div>
            </div>
            
            <div class="space-y-3">
              <div class="text-sm text-gray-600 line-clamp-3">{{ template.content }}</div>
              <div class="flex justify-between text-sm">
                <span class="text-gray-500">استفاده شده:</span>
                <span class="font-medium">{{ template.usageCount }}</span>
              </div>
              <div class="flex justify-between text-sm">
                <span class="text-gray-500">نرخ موفقیت:</span>
                <span class="font-medium">{{ template.successRate }}%</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- مودال فرم اعلان -->
    <div v-if="showNotificationForm" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg shadow-xl max-w-4xl w-full mx-4 max-h-[90vh] overflow-y-auto">
        <div class="p-6 border-b border-gray-200">
          <div class="flex items-center justify-between">
            <h3 class="text-lg font-semibold text-gray-900">
              {{ editingNotification ? 'ویرایش اعلان' : 'افزودن اعلان جدید' }}
            </h3>
            <button class="text-gray-400 hover:text-gray-600" @click="closeForm">
              <span class="i-heroicons-x-mark text-xl"></span>
            </button>
          </div>
        </div>
        
        <form class="p-6 space-y-6" @submit.prevent="handleSubmit">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">عنوان اعلان *</label>
              <input v-model="form.title" type="text" required class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="عنوان اعلان">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نوع اعلان *</label>
              <select v-model="form.type" required class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                <option value="email">ایمیل</option>
                <option value="sms">پیامک</option>
                <option value="push">اعلان</option>
                <option value="in-app">درون برنامه</option>
              </select>
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">محتوا *</label>
            <textarea v-model="form.content" rows="4" required class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="محتوای اعلان"></textarea>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">محرک اعلان *</label>
              <select v-model="form.trigger" required class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                <option value="coupon_created">ایجاد کوپن</option>
                <option value="coupon_expiring">انقضای کوپن</option>
                <option value="campaign_started">شروع کمپین</option>
                <option value="campaign_ending">پایان کمپین</option>
                <option value="user_inactive">عدم فعالیت کاربر</option>
                <option value="custom">سفارشی</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
              <select v-model="form.status" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                <option value="active">فعال</option>
                <option value="paused">متوقف</option>
                <option value="draft">پیش‌نویس</option>
              </select>
            </div>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">رنگ اعلان</label>
              <input v-model="form.color" type="color" class="w-full h-10 border border-gray-300 rounded-lg">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">اولویت</label>
              <select v-model="form.priority" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                <option value="low">کم</option>
                <option value="medium">متوسط</option>
                <option value="high">زیاد</option>
                <option value="urgent">فوری</option>
              </select>
            </div>
          </div>
        </form>
        
        <div class="p-6 border-t border-gray-200 flex justify-end space-x-3 space-x-reverse">
          <button class="px-6 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200 transition-colors" @click="closeForm">
            انصراف
          </button>
          <button class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors" @click="handleSubmit">
            ذخیره
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed } from 'vue'

const activeTab = ref('notifications')
const showNotificationForm = ref(false)
interface NotificationData {
  id?: number | string
  [key: string]: unknown
}
const editingNotification = ref<NotificationData | null>(null)
const selectedNotifications = ref<number[]>([])
const filterType = ref('')
const filterStatus = ref('')

const tabs = [
  { value: 'notifications', label: 'اعلان‌های فعال' },
  { value: 'reminders', label: 'یادآوری‌ها' },
  { value: 'scheduling', label: 'زمان‌بندی' },
  { value: 'templates', label: 'قالب‌های اعلان' }
]

const stats = ref({
  activeNotifications: 8,
  scheduledReminders: 12,
  deliveryRate: 94.5,
  engagementRate: 23.8
})

const notifications = ref([
  {
    id: 1,
    title: 'یادآوری انقضای کوپن',
    type: 'ایمیل',
    trigger: 'انقضای کوپن',
    content: 'کوپن شما تا 24 ساعت دیگر منقضی می‌شود',
    color: '#3b82f6',
    status: 'active',
    lastSent: '2024-01-15T10:30:00',
    sentCount: 156,
    openRate: 23.5
  },
  {
    id: 2,
    title: 'شروع کمپین جدید',
    type: 'پیامک',
    trigger: 'شروع کمپین',
    content: 'کمپین تخفیف ویژه شروع شده است',
    color: '#10b981',
    status: 'active',
    lastSent: '2024-01-14T09:15:00',
    sentCount: 89,
    openRate: 18.2
  }
])

const reminders = ref([
  {
    id: 1,
    title: 'یادآوری هفتگی',
    type: 'ایمیل',
    description: 'ارسال خلاصه کوپن‌های فعال',
    color: '#3b82f6',
    status: 'active',
    schedule: 'هر جمعه ساعت 10 صبح',
    lastExecuted: '2024-01-12T10:00:00'
  },
  {
    id: 2,
    title: 'یادآوری انقضا',
    type: 'پیامک',
    description: 'یادآوری کوپن‌های در حال انقضا',
    color: '#f59e0b',
    status: 'active',
    schedule: 'هر روز ساعت 6 عصر',
    lastExecuted: '2024-01-15T18:00:00'
  }
])

const timezoneSettings = ref([
  { id: 1, name: 'تهران', time: '10:00' },
  { id: 2, name: 'اصفهان', time: '09:30' },
  { id: 3, name: 'مشهد', time: '10:30' }
])

const behaviorSettings = ref([
  { id: 1, name: 'عدم فعالیت', description: 'بعد از 7 روز عدم فعالیت', delay: '24' },
  { id: 2, name: 'ترک سبد خرید', description: 'بعد از ترک سبد خرید', delay: '3' },
  { id: 3, name: 'بازدید محصول', description: 'بعد از بازدید محصول', delay: '6' }
])

const notificationTemplates = ref([
  {
    id: 1,
    name: 'خوشامدگویی',
    category: 'کاربر جدید',
    content: 'به خانواده ما خوش آمدید! کوپن ویژه شما آماده است.',
    color: '#3b82f6',
    usageCount: 234,
    successRate: 85
  },
  {
    id: 2,
    name: 'یادآوری انقضا',
    category: 'انقضا',
    content: 'کوپن شما تا {{hours}} ساعت دیگر منقضی می‌شود.',
    color: '#f59e0b',
    usageCount: 156,
    successRate: 72
  }
])

const form = reactive({
  title: '',
  type: 'email',
  content: '',
  trigger: 'coupon_created',
  status: 'active',
  color: '#3b82f6',
  priority: 'medium'
})

const filteredNotifications = computed(() => {
  let filtered = notifications.value
  
  if (filterType.value) {
    filtered = filtered.filter(n => n.type === filterType.value)
  }
  
  if (filterStatus.value) {
    filtered = filtered.filter(n => n.status === filterStatus.value)
  }
  
  return filtered
})

const formatDate = (date: string): string => {
  return new Intl.DateTimeFormat('fa-IR').format(new Date(date))
}

const getStatusClass = (status: string): string => {
  const classes = {
    active: 'bg-green-100 text-green-800',
    paused: 'bg-yellow-100 text-yellow-800',
    draft: 'bg-gray-100 text-gray-800'
  }
  return classes[status as keyof typeof classes] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status: string): string => {
  const texts = {
    active: 'فعال',
    paused: 'متوقف',
    draft: 'پیش‌نویس'
  }
  return texts[status as keyof typeof texts] || status
}

interface Notification {
  id?: number | string
  title?: string
  [key: string]: unknown
}

const editNotification = (notification: Notification) => {
  editingNotification.value = notification
  Object.assign(form, notification)
  showNotificationForm.value = true
}

const duplicateNotification = (notification: Notification) => {
  const duplicate = { ...notification, id: Date.now(), title: `${notification.title} (کپی)` }
  notifications.value.unshift(duplicate)
}

interface NotificationItem {
  id?: number | string
  title?: string
  [key: string]: unknown
}

interface Reminder {
  id?: number | string
  title?: string
  [key: string]: unknown
}

interface Timezone {
  [key: string]: unknown
}

interface Behavior {
  [key: string]: unknown
}

interface Template {
  name?: string
  content?: string
  [key: string]: unknown
}

const deleteNotification = async (notification: NotificationItem) => {
  if (confirm(`آیا از حذف اعلان "${notification.title || 'نامشخص'}" اطمینان دارید؟`)) {
    try {
      const index = notifications.value.findIndex(n => n.id === notification.id)
      if (index !== -1) {
        notifications.value.splice(index, 1)
      }
    } catch (error) {
      console.error('خطا در حذف اعلان:', error)
    }
  }
}

const bulkAction = (action: string) => {
  if (selectedNotifications.value.length === 0) {
    alert('لطفاً اعلان‌هایی را انتخاب کنید')
    return
  }
  
  switch (action) {
    case 'activate':
      notifications.value.forEach(n => {
        if (selectedNotifications.value.includes(n.id)) {
          n.status = 'active'
        }
      })
      break
    case 'pause':
      notifications.value.forEach(n => {
        if (selectedNotifications.value.includes(n.id)) {
          n.status = 'paused'
        }
      })
      break
    case 'delete':
      if (confirm(`آیا از حذف ${selectedNotifications.value.length} اعلان اطمینان دارید؟`)) {
        notifications.value = notifications.value.filter(n => !selectedNotifications.value.includes(n.id))
        selectedNotifications.value = []
      }
      break
  }
}

const editReminder = (_reminder: Reminder) => {
  // پیاده‌سازی ویرایش یادآوری

}

const deleteReminder = async (reminder: Reminder) => {
  if (confirm(`آیا از حذف یادآوری "${reminder.title || 'نامشخص'}" اطمینان دارید؟`)) {
    try {
      const index = reminders.value.findIndex(r => r.id === reminder.id)
      if (index !== -1) {
        reminders.value.splice(index, 1)
      }
    } catch (error) {
      console.error('خطا در حذف یادآوری:', error)
    }
  }
}

const updateTimezone = (_timezone: Timezone) => {
  // پیاده‌سازی به‌روزرسانی منطقه زمانی

}

const updateBehavior = (_behavior: Behavior) => {
  // پیاده‌سازی به‌روزرسانی رفتار

}

const useTemplate = (template: Template) => {
  form.title = template.name
  form.content = template.content
  showNotificationForm.value = true
}

const editTemplate = (_template: Template) => {
  // پیاده‌سازی ویرایش قالب

}

const handleSubmit = async () => {
  if (!form.title || !form.type || !form.content || !form.trigger) {
    alert('لطفاً فیلدهای اجباری را پر کنید')
    return
  }
  
  if (editingNotification.value) {
    Object.assign(editingNotification.value, form)
  } else {
    notifications.value.unshift({
      id: Date.now(),
      ...form,
      lastSent: null,
      sentCount: 0,
      openRate: 0
    })
  }
  closeForm()
}

const closeForm = () => {
  showNotificationForm.value = false
  editingNotification.value = null
  Object.assign(form, { title: '', type: 'email', content: '', trigger: 'coupon_created', status: 'active', color: '#3b82f6', priority: 'medium' })
}
</script>

<!--
  کامپوننت مدیریت اعلان‌ها و یادآوری‌ها
  شامل:
  1. مدیریت اعلان‌های فعال
  2. مدیریت یادآوری‌ها
  3. زمان‌بندی هوشمند
  4. قالب‌های اعلان
  توضیحات کامل به فارسی در کد
--> 
