<template>
  <div class="space-y-6">
    <!-- تیکت‌های پشتیبانی -->
    <div class="px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-blue-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
        </svg>
        تیکت‌های پشتیبانی
      </h3>

      <div class="space-y-4">
        <div 
          v-for="(ticket, index) in supportCommunication.tickets" 
          :key="index"
          class="border border-gray-200 rounded-lg px-4 py-4"
        >
          <div class="grid grid-cols-1 md:grid-cols-3 gapx-4 py-4">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-1">شماره تیکت</label>
              <input 
                v-model="ticket.ticketNumber"
                type="text" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="TKT-2024-001"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-1">موضوع</label>
              <input 
                v-model="ticket.subject"
                type="text" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="موضوع تیکت"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-1">وضعیت</label>
              <select 
                v-model="ticket.status"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="open">باز</option>
                <option value="in_progress">در حال بررسی</option>
                <option value="waiting">در انتظار پاسخ</option>
                <option value="resolved">حل شده</option>
                <option value="closed">بسته</option>
              </select>
            </div>
          </div>
          <div class="mt-3 grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-1">اولویت</label>
              <select 
                v-model="ticket.priority"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="low">کم</option>
                <option value="normal">عادی</option>
                <option value="high">بالا</option>
                <option value="urgent">فوری</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-1">دسته‌بندی</label>
              <select 
                v-model="ticket.category"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="technical">فنی</option>
                <option value="billing">مالی</option>
                <option value="delivery">تحویل</option>
                <option value="product">محصول</option>
                <option value="general">عمومی</option>
              </select>
            </div>
          </div>
          <div class="mt-3">
            <label class="block text-sm font-medium text-gray-600 mb-1">توضیحات</label>
            <textarea 
              v-model="ticket.description"
              rows="3"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="توضیحات تیکت..."
            ></textarea>
          </div>
        </div>

        <!-- دکمه افزودن تیکت -->
        <div class="mt-4">
          <button 
            class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors flex items-center text-sm"
            @click="addTicket"
          >
            <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
            </svg>
            افزودن تیکت جدید
          </button>
        </div>
      </div>
    </div>

    <!-- چت زنده -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-green-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
        </svg>
        چت زنده
      </h3>

      <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
        <!-- اطلاعات چت -->
        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">اطلاعات چت</h4>
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">شماره چت</label>
              <input 
                v-model="supportCommunication.chatInfo.chatId"
                type="text" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="CHAT-2024-001"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">وضعیت چت</label>
              <select 
                v-model="supportCommunication.chatInfo.status"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="active">فعال</option>
                <option value="waiting">در انتظار</option>
                <option value="ended">پایان یافته</option>
                <option value="transferred">انتقال داده شده</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">عامل پشتیبانی</label>
              <input 
                v-model="supportCommunication.chatInfo.agent"
                type="text" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="نام عامل پشتیبانی"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">تاریخ شروع</label>
              <input 
                v-model="supportCommunication.chatInfo.startDate"
                type="datetime-local" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
          </div>
        </div>

        <!-- آمار چت -->
        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">آمار چت</h4>
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">تعداد پیام‌ها</label>
              <input 
                v-model="supportCommunication.chatInfo.messageCount"
                type="number" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="0"
                readonly
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">مدت زمان (دقیقه)</label>
              <input 
                v-model="supportCommunication.chatInfo.duration"
                type="number" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="0"
                readonly
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">امتیاز رضایت</label>
              <select 
                v-model="supportCommunication.chatInfo.satisfactionRating"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="">انتخاب امتیاز</option>
                <option value="1">1 - بسیار ضعیف</option>
                <option value="2">2 - ضعیف</option>
                <option value="3">3 - متوسط</option>
                <option value="4">4 - خوب</option>
                <option value="5">5 - عالی</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">یادداشت چت</label>
              <textarea 
                v-model="supportCommunication.chatInfo.notes"
                rows="3"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="یادداشت‌های مربوط به چت..."
              ></textarea>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تاریخچه پیام‌ها -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-purple-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
        </svg>
        تاریخچه پیام‌ها
      </h3>
      
      <div class="space-y-3 max-h-64 overflow-y-auto">
        <div 
          v-for="(message, index) in supportCommunication.chatMessages" 
          :key="index"
          class="flex items-start p-3 bg-gray-50 rounded-lg"
          :class="message.sender === 'customer' ? 'justify-end' : 'justify-start'"
        >
          <div 
            class="max-w-xs p-3 rounded-lg"
            :class="message.sender === 'customer' ? 'bg-blue-500 text-white' : 'bg-white border border-gray-200'"
          >
            <div class="text-sm font-medium mb-1">
              {{ message.sender === 'customer' ? 'مشتری' : 'پشتیبانی' }}
            </div>
            <div class="text-sm">{{ message.content }}</div>
            <div class="text-xs mt-2 opacity-75">{{ formatTime(message.timestamp) }}</div>
          </div>
        </div>
      </div>

      <!-- دکمه افزودن پیام -->
      <div class="mt-4">
        <button 
          class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors flex items-center text-sm"
          @click="addMessage"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
          </svg>
          افزودن پیام
        </button>
      </div>
    </div>

    <!-- تماس‌های تلفنی -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-orange-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z"></path>
        </svg>
        تماس‌های تلفنی
      </h3>

      <div class="space-y-4">
        <div 
          v-for="(call, index) in supportCommunication.phoneCalls" 
          :key="index"
          class="border border-gray-200 rounded-lg px-4 py-4"
        >
          <div class="grid grid-cols-1 md:grid-cols-3 gapx-4 py-4">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-1">شماره تماس</label>
              <input 
                v-model="call.phoneNumber"
                type="tel" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="09123456789"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-1">نوع تماس</label>
              <select 
                v-model="call.callType"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="incoming">ورودی</option>
                <option value="outgoing">خروجی</option>
                <option value="missed">از دست رفته</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-1">وضعیت</label>
              <select 
                v-model="call.status"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="answered">پاسخ داده شده</option>
                <option value="no_answer">بدون پاسخ</option>
                <option value="busy">مشغول</option>
                <option value="failed">ناموفق</option>
              </select>
            </div>
          </div>
          <div class="mt-3 grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-1">تاریخ و زمان</label>
              <input 
                v-model="call.callDate"
                type="datetime-local" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-1">مدت زمان (ثانیه)</label>
              <input 
                v-model="call.duration"
                type="number" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="0"
              />
            </div>
          </div>
          <div class="mt-3">
            <label class="block text-sm font-medium text-gray-600 mb-1">یادداشت تماس</label>
            <textarea 
              v-model="call.notes"
              rows="2"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="یادداشت‌های مربوط به تماس..."
            ></textarea>
          </div>
        </div>

        <!-- دکمه افزودن تماس -->
        <div class="mt-4">
          <button 
            class="px-4 py-2 bg-orange-600 text-white rounded-lg hover:bg-orange-700 transition-colors flex items-center text-sm"
            @click="addPhoneCall"
          >
            <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
            </svg>
            افزودن تماس جدید
          </button>
        </div>
      </div>
    </div>

    <!-- خلاصه پشتیبانی -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-red-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
        </svg>
        خلاصه پشتیبانی
      </h3>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
        <!-- آمار کلی -->
        <div class="space-y-3">
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">تعداد تیکت‌ها:</span>
            <span class="font-medium">{{ supportCommunication.tickets.length }} تیکت</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">تیکت‌های باز:</span>
            <span class="font-medium text-orange-600">{{ openTicketsCount }} تیکت</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">تیکت‌های حل شده:</span>
            <span class="font-medium text-green-600">{{ resolvedTicketsCount }} تیکت</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">متوسط زمان پاسخ:</span>
            <span class="font-medium">{{ averageResponseTime }} دقیقه</span>
          </div>
        </div>

        <!-- وضعیت ارتباطات -->
        <div class="space-y-3">
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">وضعیت چت:</span>
            <span class="font-medium" :class="getChatStatusColor(supportCommunication.chatInfo.status)">
              {{ getChatStatusText(supportCommunication.chatInfo.status) }}
            </span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">تعداد پیام‌ها:</span>
            <span class="font-medium">{{ supportCommunication.chatInfo.messageCount }} پیام</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">تعداد تماس‌ها:</span>
            <span class="font-medium">{{ supportCommunication.phoneCalls.length }} تماس</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">امتیاز رضایت:</span>
            <span class="font-medium">{{ supportCommunication.chatInfo.satisfactionRating || 'ثبت نشده' }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'

// تعریف props
const props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({
      tickets: [],
      chatInfo: {
        chatId: '',
        status: 'active',
        agent: '',
        startDate: '',
        messageCount: 0,
        duration: 0,
        satisfactionRating: '',
        notes: ''
      },
      chatMessages: [],
      phoneCalls: []
    })
  }
})

// تعریف emit
const emit = defineEmits(['update:modelValue'])

// متغیر محلی
const supportCommunication = ref({ ...props.modelValue })

// محاسبات
const openTicketsCount = computed(() => {
  return supportCommunication.value.tickets.filter(ticket => 
    ticket.status === 'open' || ticket.status === 'in_progress' || ticket.status === 'waiting'
  ).length
})

const resolvedTicketsCount = computed(() => {
  return supportCommunication.value.tickets.filter(ticket => 
    ticket.status === 'resolved' || ticket.status === 'closed'
  ).length
})

const averageResponseTime = computed(() => {
  // محاسبه متوسط زمان پاسخ (نمونه)
  return 15
})

// توابع کمکی
const getChatStatusText = (status) => {
  const statusMap = {
    'active': 'فعال',
    'waiting': 'در انتظار',
    'ended': 'پایان یافته',
    'transferred': 'انتقال داده شده'
  }
  return statusMap[status] || status
}

const getChatStatusColor = (status) => {
  const colors = {
    'active': 'text-green-600',
    'waiting': 'text-yellow-600',
    'ended': 'text-gray-600',
    'transferred': 'text-blue-600'
  }
  return colors[status] || 'text-gray-600'
}

const formatTime = (date) => {
  if (!date) return ''
  return new Date(date).toLocaleTimeString('fa-IR', { 
    hour: '2-digit', 
    minute: '2-digit' 
  })
}

const addTicket = () => {
  supportCommunication.value.tickets.push({
    ticketNumber: '',
    subject: '',
    status: 'open',
    priority: 'normal',
    category: 'general',
    description: ''
  })
}

const addMessage = () => {
  const newMessage = {
    sender: 'support',
    content: 'پیام جدید از پشتیبانی',
    timestamp: new Date().toISOString()
  }
  supportCommunication.value.chatMessages.push(newMessage)
  supportCommunication.value.chatInfo.messageCount++
}

const addPhoneCall = () => {
  supportCommunication.value.phoneCalls.push({
    phoneNumber: '',
    callType: 'incoming',
    status: 'answered',
    callDate: '',
    duration: 0,
    notes: ''
  })
}

// نظارت بر تغییرات
watch(supportCommunication, (newValue) => {
  emit('update:modelValue', newValue)
}, { deep: true })
</script> 
