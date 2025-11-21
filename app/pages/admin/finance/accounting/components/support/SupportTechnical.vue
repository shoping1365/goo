<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">پشتیبانی فنی</h4>
        <p class="text-sm text-gray-600 mt-1">ارتباط مستقیم با تیم پشتیبانی فنی</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
          @click="startLiveChat"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
          </svg>
          چت زنده
        </button>
      </div>
    </div>

    <!-- اطلاعات تماس -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
      <div class="bg-blue-50 rounded-lg p-6 text-center">
        <div class="w-12 h-12 bg-blue-500 rounded-lg flex items-center justify-center mx-auto mb-4">
          <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z" />
          </svg>
        </div>
        <h5 class="font-medium text-blue-900 mb-2">تماس تلفنی</h5>
        <p class="text-sm text-blue-700 mb-3">شنبه تا چهارشنبه ۸ صبح تا ۶ عصر</p>
        <a href="tel:+982188765432" class="text-lg font-bold text-blue-600 hover:text-blue-800">
          ۰۲۱-۸۸۷۶۵۴۳۲
        </a>
      </div>

      <div class="bg-green-50 rounded-lg p-6 text-center">
        <div class="w-12 h-12 bg-green-500 rounded-lg flex items-center justify-center mx-auto mb-4">
          <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 4.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
          </svg>
        </div>
        <h5 class="font-medium text-green-900 mb-2">ایمیل پشتیبانی</h5>
        <p class="text-sm text-green-700 mb-3">پاسخ در کمتر از ۲۴ ساعت</p>
        <a href="mailto:support@example.com" class="text-lg font-bold text-green-600 hover:text-green-800">
          support@example.com
        </a>
      </div>

      <div class="bg-purple-50 rounded-lg p-6 text-center">
        <div class="w-12 h-12 bg-purple-500 rounded-lg flex items-center justify-center mx-auto mb-4">
          <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
          </svg>
        </div>
        <h5 class="font-medium text-purple-900 mb-2">چت زنده</h5>
        <p class="text-sm text-purple-700 mb-3">پاسخ فوری در ساعات کاری</p>
        <button
          class="text-lg font-bold text-purple-600 hover:text-purple-800"
          @click="startLiveChat"
        >
          شروع چت
        </button>
      </div>
    </div>

    <!-- فرم ارسال تیکت -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">ارسال تیکت پشتیبانی</h5>
      <form class="space-y-4" @submit.prevent="submitTicket">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نام و نام خانوادگی</label>
            <input
              v-model="ticketForm.name"
              type="text"
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">ایمیل</label>
            <input
              v-model="ticketForm.email"
              type="email"
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">موضوع</label>
          <select
            v-model="ticketForm.subject"
            required
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">انتخاب کنید</option>
            <option value="connection">مشکل در اتصال</option>
            <option value="sync">مشکل در همگام‌سازی</option>
            <option value="configuration">مشکل در پیکربندی</option>
            <option value="performance">مشکل در عملکرد</option>
            <option value="other">سایر موارد</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">اولویت</label>
          <select
            v-model="ticketForm.priority"
            required
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="low">کم</option>
            <option value="medium">متوسط</option>
            <option value="high">زیاد</option>
            <option value="urgent">فوری</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات</label>
          <textarea
            v-model="ticketForm.description"
            rows="4"
            required
            placeholder="مشکل خود را به طور کامل توضیح دهید..."
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          ></textarea>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">فایل پیوست (اختیاری)</label>
          <input
            type="file"
            multiple
            accept=".jpg,.jpeg,.png,.gif,.pdf,.doc,.docx,.txt"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            @change="handleFileUpload"
          />
          <p class="text-xs text-gray-500 mt-1">حداکثر ۵ فایل، هر فایل حداکثر ۱۰ مگابایت</p>
        </div>

        <div class="flex items-center justify-end">
          <button
            type="submit"
            :disabled="isSubmitting"
            class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            <svg v-if="isSubmitting" class="w-4 h-4 ml-2 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
            </svg>
            {{ isSubmitting ? 'در حال ارسال...' : 'ارسال تیکت' }}
          </button>
        </div>
      </form>
    </div>

    <!-- تیکت‌های اخیر -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">تیکت‌های اخیر</h5>
      </div>
      <div class="divide-y divide-gray-200">
        <div
          v-for="ticket in recentTickets"
          :key="ticket.id"
          class="p-6 hover:bg-gray-50 transition-colors"
        >
          <div class="flex items-center justify-between">
            <div class="flex-1">
              <div class="flex items-center space-x-3 space-x-reverse">
                <h6 class="text-sm font-medium text-gray-900">{{ ticket.subject }}</h6>
                <span
                  class="px-2 py-1 text-xs font-medium rounded-full"
                  :class="getPriorityClass(ticket.priority)"
                >
                  {{ getPriorityLabel(ticket.priority) }}
                </span>
                <span
                  class="px-2 py-1 text-xs font-medium rounded-full"
                  :class="getStatusClass(ticket.status)"
                >
                  {{ getStatusLabel(ticket.status) }}
                </span>
              </div>
              <p class="text-sm text-gray-600 mt-1">{{ ticket.description }}</p>
              <div class="flex items-center space-x-4 space-x-reverse mt-2 text-xs text-gray-500">
                <span>شماره تیکت: #{{ ticket.id }}</span>
                <span>{{ ticket.createdAt }}</span>
                <span v-if="ticket.lastUpdate">آخرین بروزرسانی: {{ ticket.lastUpdate }}</span>
              </div>
            </div>
            <button
              class="inline-flex items-center px-3 py-1 border border-transparent text-xs font-medium rounded-md text-blue-700 bg-blue-100 hover:bg-blue-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
              @click="viewTicket(ticket)"
            >
              مشاهده
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// وضعیت ارسال
const isSubmitting = ref(false)

// فرم تیکت
const ticketForm = ref({
  name: '',
  email: '',
  subject: '',
  priority: 'medium',
  description: '',
  attachments: [] as File[]
})

// تیکت‌های اخیر
const recentTickets = ref([
  {
    id: 1001,
    subject: 'مشکل در اتصال به نرم‌افزار هلو',
    description: 'اتصال به نرم‌افزار هلو پس از ۳۰ ثانیه timeout می‌شود',
    priority: 'high',
    status: 'open',
    createdAt: '۲ روز پیش',
    lastUpdate: '۱ روز پیش'
  },
  {
    id: 1002,
    subject: 'خطا در همگام‌سازی تراکنش‌ها',
    description: 'برخی تراکنش‌ها به درستی همگام نمی‌شوند',
    priority: 'medium',
    status: 'in_progress',
    createdAt: '۱ هفته پیش',
    lastUpdate: '۳ روز پیش'
  },
  {
    id: 1003,
    subject: 'درخواست راهنمای پیکربندی',
    description: 'نیاز به راهنمایی برای تنظیمات پیشرفته',
    priority: 'low',
    status: 'resolved',
    createdAt: '۲ هفته پیش',
    lastUpdate: '۱ هفته پیش'
  }
])

// کلاس اولویت
const getPriorityClass = (priority: string) => {
  const classes = {
    low: 'bg-green-100 text-green-700',
    medium: 'bg-yellow-100 text-yellow-700',
    high: 'bg-orange-100 text-orange-700',
    urgent: 'bg-red-100 text-red-700'
  }
  return classes[priority] || 'bg-gray-100 text-gray-700'
}

// برچسب اولویت
const getPriorityLabel = (priority: string) => {
  const labels = {
    low: 'کم',
    medium: 'متوسط',
    high: 'زیاد',
    urgent: 'فوری'
  }
  return labels[priority] || priority
}

// کلاس وضعیت
const getStatusClass = (status: string) => {
  const classes = {
    open: 'bg-blue-100 text-blue-700',
    in_progress: 'bg-yellow-100 text-yellow-700',
    resolved: 'bg-green-100 text-green-700',
    closed: 'bg-gray-100 text-gray-700'
  }
  return classes[status] || 'bg-gray-100 text-gray-700'
}

// برچسب وضعیت
const getStatusLabel = (status: string) => {
  const labels = {
    open: 'باز',
    in_progress: 'در حال بررسی',
    resolved: 'حل شده',
    closed: 'بسته'
  }
  return labels[status] || status
}

// شروع چت زنده
const startLiveChat = () => {
  // TODO: شروع چت زنده

}

// ارسال تیکت
const submitTicket = async () => {
  try {
    isSubmitting.value = true
    // TODO: ارسال تیکت به سرور

    // شبیه‌سازی ارسال
    await new Promise(resolve => setTimeout(resolve, 2000))
    
    // پاک کردن فرم
    ticketForm.value = {
      name: '',
      email: '',
      subject: '',
      priority: 'medium',
      description: '',
      attachments: []
    }

  } catch (error) {
    console.error('خطا در ارسال تیکت:', error)
  } finally {
    isSubmitting.value = false
  }
}

// آپلود فایل
const handleFileUpload = (event: Event) => {
  const target = event.target as HTMLInputElement
  if (target.files) {
    ticketForm.value.attachments = Array.from(target.files)
  }
}

interface Ticket {
  id: number | string
  [key: string]: unknown
}

// مشاهده تیکت
const viewTicket = (_ticket: Ticket) => {
  // TODO: مشاهده جزئیات تیکت

}
</script>

<!--
  کامپوننت پشتیبانی فنی
  شامل:
  1. اطلاعات تماس
  2. فرم ارسال تیکت
  3. تیکت‌های اخیر
  4. چت زنده
  5. طراحی مدرن و کاملاً ریسپانسیو
--> 
