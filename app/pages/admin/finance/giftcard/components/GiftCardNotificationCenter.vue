<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex justify-between items-center">
        <div>
          <h2 class="text-xl font-bold text-gray-900">مرکز نوتیفیکیشن</h2>
          <p class="text-gray-600 mt-1">مدیریت اطلاع‌رسانی‌های داخلی سیستم گیفت کارت</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button 
            @click="markAllAsRead"
            class="px-4 py-2 bg-gray-600 text-white text-sm font-medium rounded-lg hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
          >
            علامت‌گذاری همه به عنوان خوانده شده
          </button>
          <button 
            @click="refreshNotifications"
            class="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
          >
            به‌روزرسانی
          </button>
        </div>
      </div>
    </div>

    <!-- آمار نوتیفیکیشن‌ها -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-blue-100 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-5 5v-5zM4 19h6a2 2 0 002-2V7a2 2 0 00-2-2H4a2 2 0 00-2 2v10a2 2 0 002 2z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">کل نوتیفیکیشن‌ها</p>
            <p class="text-2xl font-bold text-gray-900">{{ totalNotifications }}</p>
            <p class="text-xs text-blue-600">{{ unreadCount }} خوانده نشده</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-green-100 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">خریدهای جدید</p>
            <p class="text-2xl font-bold text-gray-900">{{ newPurchasesCount }}</p>
            <p class="text-xs text-green-600">امروز</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-yellow-100 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">کارت‌های در حال انقضا</p>
            <p class="text-2xl font-bold text-gray-900">{{ expiringCardsCount }}</p>
            <p class="text-xs text-yellow-600">7 روز آینده</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-red-100 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">تراکنش‌های مشکوک</p>
            <p class="text-2xl font-bold text-gray-900">{{ suspiciousTransactionsCount }}</p>
            <p class="text-xs text-red-600">نیاز به بررسی</p>
          </div>
        </div>
      </div>
    </div>

    <!-- فیلترها -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">نوع نوتیفیکیشن</label>
          <select 
            v-model="filters.type"
            @change="applyFilters"
            class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه</option>
            <option value="purchase">خرید جدید</option>
            <option value="expiry">هشدار انقضا</option>
            <option value="suspicious">تراکنش مشکوک</option>
            <option value="system">سیستم</option>
          </select>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">اولویت</label>
          <select 
            v-model="filters.priority"
            @change="applyFilters"
            class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه</option>
            <option value="high">بالا</option>
            <option value="medium">متوسط</option>
            <option value="low">پایین</option>
          </select>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">وضعیت</label>
          <select 
            v-model="filters.status"
            @change="applyFilters"
            class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه</option>
            <option value="unread">خوانده نشده</option>
            <option value="read">خوانده شده</option>
          </select>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">تاریخ</label>
          <select 
            v-model="filters.date"
            @change="applyFilters"
            class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه</option>
            <option value="today">امروز</option>
            <option value="week">هفته جاری</option>
            <option value="month">ماه جاری</option>
          </select>
        </div>
      </div>
    </div>

    <!-- لیست نوتیفیکیشن‌ها -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-medium text-gray-900">نوتیفیکیشن‌ها</h3>
      </div>
      
      <div class="divide-y divide-gray-200">
        <div 
          v-for="notification in filteredNotifications" 
          :key="notification.id"
          :class="{
            'bg-blue-50': !notification.isRead,
            'hover:bg-gray-50': notification.isRead
          }"
          class="p-6 transition-colors duration-200"
        >
          <div class="flex items-start space-x-4 space-x-reverse">
            <!-- آیکون -->
            <div class="flex-shrink-0">
              <div :class="getNotificationIconClass(notification.type)" class="w-10 h-10 rounded-lg flex items-center justify-center">
                <component :is="getNotificationIcon(notification.type)" class="w-5 h-5" />
              </div>
            </div>
            
            <!-- محتوا -->
            <div class="flex-1 min-w-0">
              <div class="flex items-center justify-between">
                <div class="flex items-center space-x-2 space-x-reverse">
                  <h4 class="text-sm font-medium text-gray-900">{{ notification.title }}</h4>
                  <span :class="getPriorityBadgeClass(notification.priority)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                    {{ getPriorityLabel(notification.priority) }}
                  </span>
                  <span v-if="!notification.isRead" class="inline-flex w-2 h-2 bg-blue-600 rounded-full"></span>
                </div>
                <div class="flex items-center space-x-2 space-x-reverse">
                  <span class="text-xs text-gray-500">{{ formatDate(notification.createdAt) }}</span>
                  <div class="flex items-center space-x-1 space-x-reverse">
                    <button 
                      @click="markAsRead(notification)"
                      v-if="!notification.isRead"
                      class="text-gray-400 hover:text-gray-600"
                      title="علامت‌گذاری به عنوان خوانده شده"
                    >
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                      </svg>
                    </button>
                    <button 
                      @click="deleteNotification(notification)"
                      class="text-gray-400 hover:text-red-600"
                      title="حذف"
                    >
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                      </svg>
                    </button>
                  </div>
                </div>
              </div>
              
              <p class="text-sm text-gray-600 mt-1">{{ notification.message }}</p>
              
              <!-- جزئیات اضافی -->
              <div v-if="notification.details" class="mt-3 bg-gray-50 rounded-lg p-3">
                <div class="grid grid-cols-1 md:grid-cols-2 gap-6 text-sm">
                  <div v-for="(value, key) in notification.details" :key="key" class="flex justify-between">
                    <span class="text-gray-600">{{ getDetailLabel(key) }}:</span>
                    <span class="font-medium text-gray-900">{{ value }}</span>
                  </div>
                </div>
              </div>
              
              <!-- عملیات -->
              <div v-if="notification.actions" class="mt-3 flex items-center space-x-3 space-x-reverse">
                <button 
                  v-for="action in notification.actions"
                  :key="action.id"
                  @click="handleAction(notification, action)"
                  :class="getActionButtonClass(action.type)"
                  class="px-3 py-1 text-xs font-medium rounded-md focus:outline-none focus:ring-2 focus:ring-offset-2"
                >
                  {{ action.label }}
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
      
      <!-- پیام خالی -->
      <div v-if="filteredNotifications.length === 0" class="p-12 text-center">
        <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-5 5v-5zM4 19h6a2 2 0 002-2V7a2 2 0 00-2-2H4a2 2 0 00-2 2v10a2 2 0 002 2z"></path>
        </svg>
        <h3 class="mt-2 text-sm font-medium text-gray-900">نوتیفیکیشنی یافت نشد</h3>
        <p class="mt-1 text-sm text-gray-500">هیچ نوتیفیکیشنی با فیلترهای انتخاب شده وجود ندارد.</p>
      </div>
    </div>

    <!-- صفحه‌بندی -->
    <div v-if="totalPages > 1" class="bg-white px-4 py-3 flex items-center justify-between border-t border-gray-200 sm:px-6">
      <div class="flex-1 flex justify-between sm:hidden">
        <button 
          @click="previousPage"
          :disabled="currentPage === 1"
          class="relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          قبلی
        </button>
        <button 
          @click="nextPage"
          :disabled="currentPage === totalPages"
          class="mr-3 relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          بعدی
        </button>
      </div>
      <div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
        <div>
          <p class="text-sm text-gray-700">
            نمایش 
            <span class="font-medium">{{ (currentPage - 1) * pageSize + 1 }}</span>
            تا 
            <span class="font-medium">{{ Math.min(currentPage * pageSize, totalNotifications) }}</span>
            از 
            <span class="font-medium">{{ totalNotifications }}</span>
            نوتیفیکیشن
          </p>
        </div>
        <div>
          <nav class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px" aria-label="Pagination">
            <button 
              @click="previousPage"
              :disabled="currentPage === 1"
              class="relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
            >
              <span class="sr-only">قبلی</span>
              <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                <path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd" />
              </svg>
            </button>
            <button 
              v-for="page in visiblePages" 
              :key="page"
              @click="goToPage(page)"
              :class="{
                'bg-blue-50 border-blue-500 text-blue-600': page === currentPage,
                'bg-white border-gray-300 text-gray-500 hover:bg-gray-50': page !== currentPage
              }"
              class="relative inline-flex items-center px-4 py-2 border text-sm font-medium"
            >
              {{ page }}
            </button>
            <button 
              @click="nextPage"
              :disabled="currentPage === totalPages"
              class="relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
            >
              <span class="sr-only">بعدی</span>
              <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                <path fill-rule="evenodd" d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z" clip-rule="evenodd" />
              </svg>
            </button>
          </nav>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted } from 'vue'

// Reactive data
const currentPage = ref(1)
const pageSize = ref(10)

// فیلترها
const filters = reactive({
  type: '',
  priority: '',
  status: '',
  date: ''
})

// نوتیفیکیشن‌های نمونه
const notifications = ref([
  {
    id: 1,
    type: 'purchase',
    priority: 'medium',
    title: 'خرید جدید گیفت کارت',
    message: 'یک گیفت کارت جدید با مبلغ 500,000 تومان خریداری شد.',
    isRead: false,
    createdAt: new Date(Date.now() - 1800000),
    details: {
      'مبلغ': '500,000 تومان',
      'خریدار': 'علی احمدی',
      'کد کارت': 'GC-BIRTHDAY-2024-001',
      'روش پرداخت': 'آنلاین'
    },
    actions: [
      { id: 1, label: 'مشاهده جزئیات', type: 'primary' },
      { id: 2, label: 'تأیید', type: 'success' }
    ]
  },
  {
    id: 2,
    type: 'expiry',
    priority: 'high',
    title: 'هشدار انقضای نزدیک',
    message: 'کارت GC-WEDDING-2024-002 تا 3 روز دیگر منقضی می‌شود.',
    isRead: false,
    createdAt: new Date(Date.now() - 3600000),
    details: {
      'کد کارت': 'GC-WEDDING-2024-002',
      'تاریخ انقضا': '1403/12/26',
      'مبلغ باقی‌مانده': '750,000 تومان',
      'گیرنده': 'فاطمه محمدی'
    },
    actions: [
      { id: 1, label: 'تمدید خودکار', type: 'warning' },
      { id: 2, label: 'ارسال یادآوری', type: 'primary' }
    ]
  },
  {
    id: 3,
    type: 'suspicious',
    priority: 'high',
    title: 'تراکنش مشکوک شناسایی شد',
    message: 'تراکنش با مبلغ 2,000,000 تومان از کارت GC-NEWYEAR-2024-003 مشکوک به نظر می‌رسد.',
    isRead: true,
    createdAt: new Date(Date.now() - 7200000),
    details: {
      'کد کارت': 'GC-NEWYEAR-2024-003',
      'مبلغ تراکنش': '2,000,000 تومان',
      'زمان تراکنش': '14:30',
      'IP آدرس': '192.168.1.100'
    },
    actions: [
      { id: 1, label: 'بررسی', type: 'danger' },
      { id: 2, label: 'مسدود کردن', type: 'danger' }
    ]
  },
  {
    id: 4,
    type: 'purchase',
    priority: 'low',
    title: 'خرید جدید گیفت کارت',
    message: 'یک گیفت کارت جدید با مبلغ 200,000 تومان خریداری شد.',
    isRead: true,
    createdAt: new Date(Date.now() - 10800000),
    details: {
      'مبلغ': '200,000 تومان',
      'خریدار': 'حسن رضایی',
      'کد کارت': 'GC-NEWYEAR-2024-004',
      'روش پرداخت': 'آنلاین'
    },
    actions: [
      { id: 1, label: 'مشاهده جزئیات', type: 'primary' }
    ]
  },
  {
    id: 5,
    type: 'system',
    priority: 'medium',
    title: 'به‌روزرسانی سیستم',
    message: 'سیستم گیفت کارت با موفقیت به‌روزرسانی شد.',
    isRead: true,
    createdAt: new Date(Date.now() - 14400000),
    details: {
      'نسخه جدید': '2.1.0',
      'تاریخ به‌روزرسانی': '1403/12/24',
      'مدت زمان': '5 دقیقه'
    },
    actions: []
  }
])

// Computed properties
const totalNotifications = computed(() => notifications.value.length)
const unreadCount = computed(() => notifications.value.filter(n => !n.isRead).length)
const newPurchasesCount = computed(() => notifications.value.filter(n => n.type === 'purchase' && !n.isRead).length)
const expiringCardsCount = computed(() => notifications.value.filter(n => n.type === 'expiry' && !n.isRead).length)
const suspiciousTransactionsCount = computed(() => notifications.value.filter(n => n.type === 'suspicious' && !n.isRead).length)

const filteredNotifications = computed(() => {
  let filtered = notifications.value

  if (filters.type) {
    filtered = filtered.filter(n => n.type === filters.type)
  }

  if (filters.priority) {
    filtered = filtered.filter(n => n.priority === filters.priority)
  }

  if (filters.status === 'read') {
    filtered = filtered.filter(n => n.isRead)
  } else if (filters.status === 'unread') {
    filtered = filtered.filter(n => !n.isRead)
  }

  if (filters.date) {
    const now = new Date()
    const today = new Date(now.getFullYear(), now.getMonth(), now.getDate())
    const weekAgo = new Date(today.getTime() - 7 * 24 * 60 * 60 * 1000)
    const monthAgo = new Date(today.getTime() - 30 * 24 * 60 * 60 * 1000)

    filtered = filtered.filter(n => {
      if (filters.date === 'today') {
        return n.createdAt >= today
      } else if (filters.date === 'week') {
        return n.createdAt >= weekAgo
      } else if (filters.date === 'month') {
        return n.createdAt >= monthAgo
      }
      return true
    })
  }

  return filtered
})

const totalPages = computed(() => Math.ceil(filteredNotifications.value.length / pageSize.value))

const visiblePages = computed(() => {
  const pages = []
  const start = Math.max(1, currentPage.value - 2)
  const end = Math.min(totalPages.value, currentPage.value + 2)
  
  for (let i = start; i <= end; i++) {
    pages.push(i)
  }
  
  return pages
})

// Methods
const applyFilters = () => {
  currentPage.value = 1
  console.log('فیلترها اعمال شد:', filters)
}

const markAsRead = (notification: any) => {
  notification.isRead = true
  console.log('نوتیفیکیشن به عنوان خوانده شده علامت‌گذاری شد:', notification.id)
}

const markAllAsRead = () => {
  notifications.value.forEach(n => n.isRead = true)
  console.log('همه نوتیفیکیشن‌ها به عنوان خوانده شده علامت‌گذاری شدند')
}

const deleteNotification = (notification: any) => {
  const index = notifications.value.findIndex(n => n.id === notification.id)
  if (index > -1) {
    notifications.value.splice(index, 1)
  }
  console.log('نوتیفیکیشن حذف شد:', notification.id)
}

const handleAction = (notification: any, action: any) => {
  console.log('عملیات انجام شد:', { notificationId: notification.id, actionId: action.id, actionLabel: action.label })
  
  // اینجا می‌توانید منطق عملیات مختلف را اضافه کنید
  switch (action.label) {
    case 'مشاهده جزئیات':
      // باز کردن مودال جزئیات
      break
    case 'تأیید':
      // تأیید تراکنش
      break
    case 'بررسی':
      // بررسی تراکنش مشکوک
      break
    case 'مسدود کردن':
      // مسدود کردن کارت
      break
  }
}

const refreshNotifications = () => {
  console.log('نوتیفیکیشن‌ها به‌روزرسانی شدند')
  // اینجا می‌توانید API call برای دریافت نوتیفیکیشن‌های جدید اضافه کنید
}

const formatDate = (date: Date) => {
  return new Intl.DateTimeFormat('fa-IR').format(date)
}

const getNotificationIcon = (type: string) => {
  const icons = {
    purchase: 'ShoppingBagIcon',
    expiry: 'ClockIcon',
    suspicious: 'ExclamationTriangleIcon',
    system: 'CogIcon'
  }
  return icons[type] || 'BellIcon'
}

const getNotificationIconClass = (type: string) => {
  const classes = {
    purchase: 'bg-green-100 text-green-600',
    expiry: 'bg-yellow-100 text-yellow-600',
    suspicious: 'bg-red-100 text-red-600',
    system: 'bg-blue-100 text-blue-600'
  }
  return classes[type] || 'bg-gray-100 text-gray-600'
}

const getPriorityBadgeClass = (priority: string) => {
  const classes = {
    high: 'bg-red-100 text-red-800',
    medium: 'bg-yellow-100 text-yellow-800',
    low: 'bg-green-100 text-green-800'
  }
  return classes[priority] || 'bg-gray-100 text-gray-800'
}

const getPriorityLabel = (priority: string) => {
  const labels = {
    high: 'بالا',
    medium: 'متوسط',
    low: 'پایین'
  }
  return labels[priority] || priority
}

const getDetailLabel = (key: string) => {
  const labels = {
    'مبلغ': 'مبلغ',
    'خریدار': 'خریدار',
    'کد کارت': 'کد کارت',
    'روش پرداخت': 'روش پرداخت',
    'تاریخ انقضا': 'تاریخ انقضا',
    'مبلغ باقی‌مانده': 'مبلغ باقی‌مانده',
    'گیرنده': 'گیرنده',
    'مبلغ تراکنش': 'مبلغ تراکنش',
    'زمان تراکنش': 'زمان تراکنش',
    'IP آدرس': 'IP آدرس',
    'نسخه جدید': 'نسخه جدید',
    'تاریخ به‌روزرسانی': 'تاریخ به‌روزرسانی',
    'مدت زمان': 'مدت زمان'
  }
  return labels[key] || key
}

const getActionButtonClass = (type: string) => {
  const classes = {
    primary: 'bg-blue-600 text-white hover:bg-blue-700 focus:ring-blue-500',
    success: 'bg-green-600 text-white hover:bg-green-700 focus:ring-green-500',
    warning: 'bg-yellow-600 text-white hover:bg-yellow-700 focus:ring-yellow-500',
    danger: 'bg-red-600 text-white hover:bg-red-700 focus:ring-red-500'
  }
  return classes[type] || 'bg-gray-600 text-white hover:bg-gray-700 focus:ring-gray-500'
}

const previousPage = () => {
  if (currentPage.value > 1) {
    currentPage.value--
  }
}

const nextPage = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value++
  }
}

const goToPage = (page: number) => {
  currentPage.value = page
}

// Lifecycle
onMounted(() => {
  console.log('Gift card notification center component mounted')
})
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 
