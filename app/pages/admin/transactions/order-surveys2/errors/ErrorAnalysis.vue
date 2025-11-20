<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="flex items-center justify-between">
      <div>
        <h3 class="text-lg font-semibold text-gray-800">تحلیل خطا</h3>
        <p class="text-gray-600 text-sm">بررسی دقیق علل خطا و راه‌حل‌های پیشنهادی</p>
      </div>
      <button 
        class="text-gray-400 hover:text-gray-600 transition-colors"
        @click="$emit('close')"
      >
        <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
        </svg>
      </button>
    </div>

    <!-- Error Details -->
    <div class="bg-white rounded-lg border border-gray-200 px-4 py-4">
      <h4 class="text-md font-semibold text-gray-800 mb-4">جزئیات خطا</h4>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">شماره سفارش</label>
            <div class="text-sm text-gray-900 font-medium">#{{ errorDetails.orderNumber }}</div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">مشتری</label>
            <div class="text-sm text-gray-900">{{ errorDetails.customerName }}</div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">شماره تلفن</label>
            <div class="text-sm text-gray-900">{{ errorDetails.phoneNumber }}</div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">نوع خطا</label>
            <span 
              :class="[
                'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium',
                getErrorTypeClass(errorDetails.errorType)
              ]"
            >
              {{ getErrorTypeText(errorDetails.errorType) }}
            </span>
          </div>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">تاریخ خطا</label>
            <div class="text-sm text-gray-900">{{ formatDateTime(errorDetails.errorDate) }}</div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">تعداد تلاش</label>
            <div class="text-sm text-gray-900">{{ errorDetails.retryCount }} / {{ maxRetries }}</div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">وضعیت</label>
            <span 
              :class="[
                'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium',
                getStatusClass(errorDetails.status)
              ]"
            >
              {{ getStatusText(errorDetails.status) }}
            </span>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">اولویت</label>
            <div class="flex items-center">
              <span 
                :class="[
                  'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium',
                  getPriorityClass(errorDetails.priority)
                ]"
              >
                {{ getPriorityText(errorDetails.priority) }}
              </span>
            </div>
          </div>
        </div>
      </div>
      
      <div class="mt-6">
        <label class="block text-sm font-medium text-gray-700 mb-1">پیام خطا</label>
        <div class="bg-red-50 border border-red-200 rounded-lg p-3">
          <div class="text-sm text-red-800">{{ errorDetails.errorMessage }}</div>
        </div>
      </div>
    </div>

    <!-- Error Analysis -->
    <div class="bg-white rounded-lg border border-gray-200 px-4 py-4">
      <h4 class="text-md font-semibold text-gray-800 mb-4">تحلیل علل خطا</h4>
      
      <div class="space-y-4">
        <div 
          v-for="(analysis, index) in errorAnalysis" 
          :key="index"
          class="border border-gray-200 rounded-lg px-4 py-4"
        >
          <div class="flex items-start justify-between">
            <div class="flex items-center space-x-3 space-x-reverse">
              <div 
                :class="[
                  'w-3 h-3 rounded-full',
                  analysis.probability === 'high' ? 'bg-red-500' : 
                  analysis.probability === 'medium' ? 'bg-yellow-500' : 'bg-green-500'
                ]"
              ></div>
              <h5 class="font-medium text-gray-900">{{ analysis.cause }}</h5>
            </div>
            <span 
              :class="[
                'inline-flex items-center px-2 py-1 rounded-full text-xs font-medium',
                analysis.probability === 'high' ? 'bg-red-100 text-red-800' : 
                analysis.probability === 'medium' ? 'bg-yellow-100 text-yellow-800' : 'bg-green-100 text-green-800'
              ]"
            >
              {{ getProbabilityText(analysis.probability) }}
            </span>
          </div>
          <p class="text-sm text-gray-600 mt-2">{{ analysis.description }}</p>
          <div v-if="analysis.solution" class="mt-3">
            <div class="text-sm font-medium text-gray-700 mb-1">راه‌حل پیشنهادی:</div>
            <div class="text-sm text-gray-600">{{ analysis.solution }}</div>
          </div>
        </div>
      </div>
    </div>

    <!-- Resolution Actions -->
    <div class="bg-white rounded-lg border border-gray-200 px-4 py-4">
      <h4 class="text-md font-semibold text-gray-800 mb-4">اقدامات حل مشکل</h4>
      
      <div class="space-y-4">
        <div class="flex items-center space-x-4 space-x-reverse">
          <button 
            :disabled="resending"
            class="px-4 py-2 bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 text-white rounded-lg text-sm transition-colors flex items-center space-x-2 space-x-reverse"
            @click="resendSMS"
          >
            <svg v-if="resending" class="animate-spin w-4 h-4" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            <svg v-else class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
            </svg>
            <span>{{ resending ? 'در حال ارسال...' : 'ارسال مجدد' }}</span>
          </button>
          
          <button 
            class="px-4 py-2 bg-green-600 hover:bg-green-700 text-white rounded-lg text-sm transition-colors"
            @click="markAsResolved"
          >
            علامت‌گذاری به عنوان حل شده
          </button>
          
          <button 
            class="px-4 py-2 bg-purple-600 hover:bg-purple-700 text-white rounded-lg text-sm transition-colors"
            @click="updatePhoneNumber"
          >
            بروزرسانی شماره تلفن
          </button>
          
          <button 
            class="px-4 py-2 bg-red-600 hover:bg-red-700 text-white rounded-lg text-sm transition-colors"
            @click="excludeFromSMS"
          >
            حذف از لیست SMS
          </button>
        </div>
        
        <div class="flex items-center space-x-4 space-x-reverse">
          <button 
            class="px-4 py-2 bg-gray-600 hover:bg-gray-700 text-white rounded-lg text-sm transition-colors"
            @click="addToBlacklist"
          >
            افزودن به لیست سیاه
          </button>
          
          <button 
            class="px-4 py-2 bg-orange-600 hover:bg-orange-700 text-white rounded-lg text-sm transition-colors"
            @click="createTicket"
          >
            ایجاد تیکت پشتیبانی
          </button>
        </div>
      </div>
    </div>

    <!-- Error History -->
    <div class="bg-white rounded-lg border border-gray-200 px-4 py-4">
      <h4 class="text-md font-semibold text-gray-800 mb-4">تاریخچه خطاها</h4>
      
      <div class="space-y-3">
        <div 
          v-for="(history, index) in errorHistory" 
          :key="index"
          class="flex items-center space-x-4 space-x-reverse p-3 bg-gray-50 rounded-lg"
        >
          <div class="flex-shrink-0">
            <div 
              :class="[
                'w-2 h-2 rounded-full',
                history.status === 'success' ? 'bg-green-500' : 'bg-red-500'
              ]"
            ></div>
          </div>
          
          <div class="flex-1">
            <div class="text-sm font-medium text-gray-900">{{ history.action }}</div>
            <div class="text-xs text-gray-500">{{ formatDateTime(history.timestamp) }}</div>
          </div>
          
          <div class="text-xs text-gray-500">
            {{ history.details }}
          </div>
        </div>
      </div>
    </div>

    <!-- Notes -->
    <div class="bg-white rounded-lg border border-gray-200 px-4 py-4">
      <h4 class="text-md font-semibold text-gray-800 mb-4">یادداشت‌ها</h4>
      
      <div class="space-y-3">
        <div 
          v-for="(note, index) in notes" 
          :key="index"
          class="border-r-4 border-blue-500 pr-4"
        >
          <div class="text-sm text-gray-900">{{ note.content }}</div>
          <div class="text-xs text-gray-500 mt-1">{{ formatDateTime(note.timestamp) }} - {{ note.author }}</div>
        </div>
      </div>
      
      <div class="mt-4">
        <textarea 
          v-model="newNote"
          placeholder="افزودن یادداشت جدید..."
          class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 resize-none"
          rows="3"
        ></textarea>
        <button 
          :disabled="!newNote.trim()"
          class="mt-2 px-4 py-2 bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 text-white rounded-lg text-sm transition-colors"
          @click="addNote"
        >
          افزودن یادداشت
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'

interface ErrorDetails {
  orderNumber: string
  customerName: string
  phoneNumber: string
  errorType: string
  errorMessage: string
  errorDate: string
  retryCount: number
  status: string
  priority: string
}

interface ErrorAnalysis {
  cause: string
  probability: 'high' | 'medium' | 'low'
  description: string
  solution?: string
}

interface ErrorHistory {
  action: string
  timestamp: string
  status: 'success' | 'failed'
  details: string
}

interface Note {
  content: string
  timestamp: string
  author: string
}

const props = defineProps<{
  errorId: number
}>()

const emit = defineEmits<{
  'close': []
  'resend': [errorId: number]
  'resolve': [errorId: number]
  'update-phone': [errorId: number, newPhone: string]
  'exclude': [errorId: number]
  'blacklist': [errorId: number]
  'create-ticket': [errorId: number]
}>()

// Reactive data
const resending = ref(false)
const newNote = ref('')
const maxRetries = 3

const errorDetails = reactive<ErrorDetails>({
  orderNumber: '12345',
  customerName: 'علی احمدی',
  phoneNumber: '09123456789',
  errorType: 'invalid_number',
  errorMessage: 'شماره تلفن وارد شده معتبر نیست یا در لیست سیاه قرار دارد',
  errorDate: '2024-01-15T10:30:00Z',
  retryCount: 2,
  status: 'pending',
  priority: 'medium'
})

const errorAnalysis = reactive<ErrorAnalysis[]>([
  {
    cause: 'شماره تلفن نامعتبر',
    probability: 'high',
    description: 'شماره تلفن وارد شده در فرمت صحیح نیست یا وجود ندارد',
    solution: 'بررسی و تصحیح شماره تلفن مشتری'
  },
  {
    cause: 'شماره در لیست سیاه',
    probability: 'medium',
    description: 'مشتری قبلاً درخواست لغو دریافت SMS کرده است',
    solution: 'حذف از لیست سیاه یا استفاده از روش‌های دیگر ارتباط'
  },
  {
    cause: 'مشکل در سرویس SMS',
    probability: 'low',
    description: 'سرویس SMS موقتاً در دسترس نیست',
    solution: 'تلاش مجدد بعد از چند دقیقه'
  }
])

const errorHistory = reactive<ErrorHistory[]>([
  {
    action: 'تلاش اول ارسال',
    timestamp: '2024-01-15T10:30:00Z',
    status: 'failed',
    details: 'خطای شماره نامعتبر'
  },
  {
    action: 'تلاش دوم ارسال',
    timestamp: '2024-01-15T10:35:00Z',
    status: 'failed',
    details: 'خطای شماره نامعتبر'
  },
  {
    action: 'بررسی دستی',
    timestamp: '2024-01-15T11:00:00Z',
    status: 'success',
    details: 'توسط اپراتور بررسی شد'
  }
])

const notes = reactive<Note[]>([
  {
    content: 'مشتری تماس گرفت و شماره جدید ارائه داد',
    timestamp: '2024-01-15T11:30:00Z',
    author: 'اپراتور 1'
  },
  {
    content: 'شماره در سیستم بروزرسانی شد',
    timestamp: '2024-01-15T11:35:00Z',
    author: 'سیستم'
  }
])

// Methods
const resendSMS = async () => {
  resending.value = true
  try {
    emit('resend', props.errorId)
    await new Promise(resolve => setTimeout(resolve, 2000))
  } finally {
    resending.value = false
  }
}

const markAsResolved = () => {
  emit('resolve', props.errorId)
}

const updatePhoneNumber = () => {
  const newPhone = prompt('شماره تلفن جدید را وارد کنید:')
  if (newPhone) {
    emit('update-phone', props.errorId, newPhone)
  }
}

const excludeFromSMS = () => {
  if (confirm('آیا مطمئن هستید که می‌خواهید این مشتری را از لیست SMS حذف کنید؟')) {
    emit('exclude', props.errorId)
  }
}

const addToBlacklist = () => {
  if (confirm('آیا مطمئن هستید که می‌خواهید این شماره را به لیست سیاه اضافه کنید؟')) {
    emit('blacklist', props.errorId)
  }
}

const createTicket = () => {
  emit('create-ticket', props.errorId)
}

const addNote = () => {
  if (newNote.value.trim()) {
    notes.push({
      content: newNote.value,
      timestamp: new Date().toISOString(),
      author: 'اپراتور فعلی'
    })
    newNote.value = ''
  }
}

// Utility functions
const formatDateTime = (dateTime: string) => {
  return new Date(dateTime).toLocaleString('fa-IR')
}

const getErrorTypeClass = (errorType: string) => {
  const classes = {
    invalid_number: 'bg-red-100 text-red-800',
    service_unavailable: 'bg-yellow-100 text-yellow-800',
    quota_exceeded: 'bg-orange-100 text-orange-800',
    network_error: 'bg-purple-100 text-purple-800',
    timeout: 'bg-gray-100 text-gray-800'
  }
  return classes[errorType as keyof typeof classes] || 'bg-gray-100 text-gray-800'
}

const getErrorTypeText = (errorType: string) => {
  const texts = {
    invalid_number: 'شماره نامعتبر',
    service_unavailable: 'سرویس در دسترس نیست',
    quota_exceeded: 'محدودیت ارسال',
    network_error: 'خطای شبکه',
    timeout: 'زمان انتظار'
  }
  return texts[errorType as keyof typeof texts] || errorType
}

const getStatusClass = (status: string) => {
  const classes = {
    pending: 'bg-yellow-100 text-yellow-800',
    retrying: 'bg-blue-100 text-blue-800',
    resolved: 'bg-green-100 text-green-800',
    permanent_failure: 'bg-red-100 text-red-800'
  }
  return classes[status as keyof typeof classes] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status: string) => {
  const texts = {
    pending: 'در انتظار',
    retrying: 'در حال تلاش مجدد',
    resolved: 'حل شده',
    permanent_failure: 'خطای دائمی'
  }
  return texts[status as keyof typeof texts] || status
}

const getPriorityClass = (priority: string) => {
  const classes = {
    high: 'bg-red-100 text-red-800',
    medium: 'bg-yellow-100 text-yellow-800',
    low: 'bg-green-100 text-green-800'
  }
  return classes[priority as keyof typeof classes] || 'bg-gray-100 text-gray-800'
}

const getPriorityText = (priority: string) => {
  const texts = {
    high: 'بالا',
    medium: 'متوسط',
    low: 'پایین'
  }
  return texts[priority as keyof typeof texts] || priority
}

const getProbabilityText = (probability: string) => {
  const texts = {
    high: 'احتمال بالا',
    medium: 'احتمال متوسط',
    low: 'احتمال پایین'
  }
  return texts[probability as keyof typeof texts] || probability
}
</script>

<style scoped>
.space-x-reverse > :not([hidden]) ~ :not([hidden]) {
  --tw-space-x-reverse: 1;
}
</style> 
