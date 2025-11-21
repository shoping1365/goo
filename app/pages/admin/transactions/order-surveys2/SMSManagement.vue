<template>
  <div class="bg-white rounded-xl shadow-lg border border-gray-200 overflow-hidden" dir="rtl">
    <!-- Header Section -->
    <div class="bg-gradient-to-r from-green-600 to-blue-600 px-6 py-4">
      <div class="flex items-center justify-between">
        <div class="flex items-center space-x-3 space-x-reverse">
          <div class="p-2 bg-white/20 rounded-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8"></path>
            </svg>
          </div>
          <div>
            <h2 class="text-xl font-bold text-white">مدیریت ارسال SMS</h2>
            <p class="text-green-100 text-sm">کنترل و مدیریت ارسال پیامک‌های نظرسنجی</p>
          </div>
        </div>
        
        <div class="flex items-center space-x-3 space-x-reverse">
          <button 
            class="px-4 py-2 bg-white/20 hover:bg-white/30 text-white rounded-lg font-medium transition-colors flex items-center space-x-2 space-x-reverse"
            @click="showSMSStatus"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
            <span>وضعیت ارسال</span>
          </button>
          
          <button 
            class="px-4 py-2 bg-white/20 hover:bg-white/30 text-white rounded-lg font-medium transition-colors flex items-center space-x-2 space-x-reverse"
            @click="showSMSTemplate"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
            </svg>
            <span>قالب پیام</span>
          </button>
        </div>
      </div>
    </div>

    <!-- Tab Navigation -->
    <div class="border-b border-gray-200">
      <nav class="flex space-x-8 space-x-reverse px-6">
        <button 
          v-for="tab in tabs" 
          :key="tab.id"
          :class="[
            'py-4 px-1 border-b-2 font-medium text-sm transition-colors',
            activeTab === tab.id
              ? 'border-green-500 text-green-600'
              : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
          ]"
          @click="activeTab = tab.id"
        >
          <div class="flex items-center space-x-2 space-x-reverse">
            <component :is="tab.icon" class="w-5 h-5" />
            <span>{{ tab.name }}</span>
          </div>
        </button>
      </nav>
    </div>

    <!-- Tab Content -->
    <div class="px-4 py-4">
      <!-- Manual Sending Tab -->
      <div v-if="activeTab === 'manual'" class="space-y-6">
        <ManualSending 
          :orders="orders"
          :selected-orders="selectedOrders"
          @send-single="sendSingleSMS"
          @send-bulk="sendBulkSMS"
          @select-orders="updateSelectedOrders"
        />
      </div>

      <!-- Automatic Sending Tab -->
      <div v-else-if="activeTab === 'automatic'" class="space-y-6">
        <AutomaticSending 
          :settings="autoSettings"
          @update-settings="updateAutoSettings"
          @test-schedule="testSchedule"
        />
      </div>

      <!-- Error Management Tab -->
      <div v-else-if="activeTab === 'errors'" class="space-y-6">
        <ErrorManagement 
          :failed-s-m-s="failedSMS"
          @resend="resendFailedSMS"
          @retry-all="retryAllFailed"
          @analyze-error="analyzeError"
        />
      </div>

      <!-- SMS Templates Tab -->
      <div v-else-if="activeTab === 'templates'" class="space-y-6">
        <SMSTemplates 
          :templates="smsTemplates"
          @save-template="saveTemplate"
          @delete-template="deleteTemplate"
          @test-template="testTemplate"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue'
import ErrorManagement from './errors/ErrorManagement.vue'
import AutomaticSending from './sending/AutomaticSending.vue'
import ManualSending from './sending/ManualSending.vue'
import SMSTemplates from './templates/SMSTemplates.vue'

// Icons
const ManualIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8"></path></svg>`
}

const AutomaticIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>`
}

const ErrorIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>`
}

const TemplateIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path></svg>`
}

// Reactive data
const activeTab = ref('manual')
const showStatusModal = ref(false)
const showTemplateModal = ref(false)
const _currentTemplate = ref(null)

const tabs = [
  { id: 'manual', name: 'ارسال دستی', icon: ManualIcon },
  { id: 'automatic', name: 'ارسال خودکار', icon: AutomaticIcon },
  { id: 'errors', name: 'مدیریت خطاها', icon: ErrorIcon },
  { id: 'templates', name: 'قالب‌های پیام', icon: TemplateIcon }
]

const orders = ref([
  {
    id: 1,
    orderNumber: 'ORD-001',
    customerName: 'علی احمدی',
    phoneNumber: '09123456789',
    orderDate: '2024-01-15',
    smsStatus: 'not_sent'
  },
  {
    id: 2,
    orderNumber: 'ORD-002',
    customerName: 'فاطمه محمدی',
    phoneNumber: '09187654321',
    orderDate: '2024-01-16',
    smsStatus: 'not_sent'
  }
])

const selectedOrders = ref<number[]>([])

const autoSettings = reactive({
  enabled: true,
  delayDays: 3,
  dailyLimit: 100,
  allowedHours: {
    start: '09:00',
    end: '21:00'
  },
  allowedDays: [0, 1, 2, 3, 4, 5, 6], // همه روزهای هفته
  timezone: 'Asia/Tehran',
  rules: [
    {
      id: 1,
      name: 'سفارشات جدید',
      condition: 'order_completed',
      delay: 3,
      priority: 'normal',
      enabled: true
    },
    {
      id: 2,
      name: 'سفارشات با مبلغ بالا',
      condition: 'high_value',
      delay: 1,
      priority: 'high',
      enabled: true
    }
  ]
})

const failedSMS = ref([
  {
    id: 1,
    orderNumber: 'ORD-003',
    customerName: 'محمد رضایی',
    phoneNumber: '09111111111',
    errorType: 'invalid_number',
    errorMessage: 'شماره تلفن نامعتبر است',
    retryCount: 2,
    lastAttempt: '2024-01-17 10:30:00',
    status: 'failed'
  }
])

const smsTemplates = ref([
  {
    id: 1,
    name: 'قالب پیش‌فرض',
    content: 'سلام {نام} عزیز، از خرید شما متشکریم. لطفاً تجربه خود را با ما به اشتراک بگذارید: {لینک}',
    variables: ['نام', 'لینک'],
    isDefault: true
  }
])

const _smsStatusData = ref({
  totalSent: 0,
  successful: 0,
  failed: 0,
  pending: 0
})

// Methods
const sendSingleSMS = async (_orderId: number) => {
  try {
    // API call implementation
  } catch {
    // console.error('Error sending SMS:', error)
  }
}

const sendBulkSMS = async (_orderIds: number[]) => {
  try {
    // API call implementation
  } catch {
    // console.error('Error sending bulk SMS:', error)
  }
}

const updateSelectedOrders = (orders: number[]) => {
  selectedOrders.value = orders
}

interface AutoSettings {
  enabled: boolean;
  delayDays: number;
  dailyLimit: number;
  allowedHours: {
    start: string;
    end: string;
  };
  allowedDays: number[];
  timezone: string;
  rules: {
    id: number;
    name: string;
    condition: string;
    delay: number;
    priority: string;
    enabled: boolean;
  }[];
}

const updateAutoSettings = (settings: AutoSettings) => {
  Object.assign(autoSettings, settings)
}

const testSchedule = () => {
}

const resendFailedSMS = async (_smsId: number) => {
  try {
    // API call implementation
  } catch {
  }
}

const retryAllFailed = async () => {
  try {
    // API call implementation
  } catch {
  }
}

const analyzeError = (_errorId: number) => {
}

const saveTemplate = (_template: Record<string, unknown>) => {
}

const deleteTemplate = (_templateId: number) => {
}

const testTemplate = (_templateId: number) => {
}

const showSMSStatus = () => {
  showStatusModal.value = true
}

const showSMSTemplate = () => {
  showTemplateModal.value = true
}

// Lifecycle
onMounted(() => {
  // Load initial data
})
</script>

<style scoped>
.space-x-reverse > :not([hidden]) ~ :not([hidden]) {
  --tw-space-x-reverse: 1;
}
</style> 
