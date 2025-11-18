<template>
  <div class="space-y-6">
    <!-- تنظیمات سیستم هشدار -->
    <div class="bg-white rounded-lg shadow">
      <div class="px-6 py-4 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <h3 class="text-lg font-medium text-gray-900">تنظیمات سیستم هشدار</h3>
          <label class="flex items-center">
            <input v-model="alertSystemEnabled" type="checkbox" class="mr-2 rounded border-gray-300 text-blue-600 focus:ring-blue-500">
            <span class="text-sm text-gray-700">فعال‌سازی سیستم هشدار</span>
          </label>
        </div>
      </div>
      <div class="px-4 py-4">
        <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
          <!-- آستانه‌های هشدار -->
          <div class="space-y-4">
            <h4 class="font-medium text-gray-900">آستانه‌های هشدار</h4>
            <div class="space-y-3">
              <div class="flex items-center justify-between p-3 bg-yellow-50 rounded-lg">
                <div>
                  <span class="text-sm font-medium text-gray-900">هشدار اولیه</span>
                  <p class="text-xs text-gray-500">امتیاز منفی برای هشدار اولیه</p>
                </div>
                <input v-model="alertThresholds.warning" type="number" min="-1000" max="0" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
              </div>
              <div class="flex items-center justify-between p-3 bg-orange-50 rounded-lg">
                <div>
                  <span class="text-sm font-medium text-gray-900">هشدار جدی</span>
                  <p class="text-xs text-gray-500">امتیاز منفی برای هشدار جدی</p>
                </div>
                <input v-model="alertThresholds.serious" type="number" min="-1000" max="0" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
              </div>
              <div class="flex items-center justify-between p-3 bg-red-50 rounded-lg">
                <div>
                  <span class="text-sm font-medium text-gray-900">مسدودسازی خودکار</span>
                  <p class="text-xs text-gray-500">امتیاز منفی برای مسدودسازی</p>
                </div>
                <input v-model="alertThresholds.blocking" type="number" min="-1000" max="0" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
              </div>
            </div>
          </div>

          <!-- تنظیمات اعلان‌ها -->
          <div class="space-y-4">
            <h4 class="font-medium text-gray-900">تنظیمات اعلان‌ها</h4>
            <div class="space-y-3">
              <label class="flex items-center">
                <input v-model="notificationSettings.email" type="checkbox" class="mr-2 rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <span class="text-sm text-gray-700">اعلان از طریق ایمیل</span>
              </label>
              <label class="flex items-center">
                <input v-model="notificationSettings.sms" type="checkbox" class="mr-2 rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <span class="text-sm text-gray-700">اعلان از طریق پیامک</span>
              </label>
              <label class="flex items-center">
                <input v-model="notificationSettings.push" type="checkbox" class="mr-2 rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <span class="text-sm text-gray-700">اعلان push</span>
              </label>
              <label class="flex items-center">
                <input v-model="notificationSettings.dashboard" type="checkbox" class="mr-2 rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <span class="text-sm text-gray-700">نمایش در داشبورد</span>
              </label>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- کاربران در معرض خطر -->
    <div class="bg-white rounded-lg shadow">
      <div class="px-6 py-4 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <h3 class="text-lg font-medium text-gray-900">کاربران در معرض خطر</h3>
          <div class="flex space-x-3 space-x-reverse">
            <select v-model="riskFilter" class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
              <option value="all">همه سطوح</option>
              <option value="warning">هشدار اولیه</option>
              <option value="serious">هشدار جدی</option>
              <option value="blocking">در آستانه مسدودسازی</option>
            </select>
            <button @click="sendBulkAlerts" class="bg-red-600 text-white px-4 py-2 rounded-md hover:bg-red-700 transition-colors">
              ارسال هشدار گروهی
            </button>
          </div>
        </div>
      </div>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                <input type="checkbox" v-model="selectAll" @change="toggleSelectAll" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کاربر</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">امتیاز فعلی</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">سطح خطر</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">آخرین فعالیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="user in filteredRiskUsers" :key="user.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <input type="checkbox" v-model="selectedUsers" :value="user.id" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <img :src="user.avatar || '/default-avatar.png'" :alt="user.name" class="w-8 h-8 rounded-full mr-3">
                  <div>
                    <div class="text-sm font-medium text-gray-900">{{ user.name }}</div>
                    <div class="text-sm text-gray-500">{{ user.email }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getScoreClass(user.score)" class="text-sm font-medium">
                  {{ user.score }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getRiskLevelClass(user.riskLevel)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                  {{ getRiskLevelText(user.riskLevel) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ formatDate(user.lastActivity) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button @click="sendAlert(user)" class="text-blue-600 hover:text-blue-900 mr-3">ارسال هشدار</button>
                <button @click="viewUserDetails(user)" class="text-green-600 hover:text-green-900">جزئیات</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- تاریخچه هشدارها -->
    <div class="bg-white rounded-lg shadow">
      <div class="px-6 py-4 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <h3 class="text-lg font-medium text-gray-900">تاریخچه هشدارها</h3>
          <div class="flex space-x-3 space-x-reverse">
            <select v-model="alertHistoryFilter" class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
              <option value="all">همه هشدارها</option>
              <option value="warning">هشدارهای اولیه</option>
              <option value="serious">هشدارهای جدی</option>
              <option value="blocking">هشدارهای مسدودسازی</option>
            </select>
            <button @click="exportAlertHistory" class="bg-green-600 text-white px-4 py-2 rounded-md hover:bg-green-700 transition-colors">
              خروجی اکسل
            </button>
          </div>
        </div>
      </div>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کاربر</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع هشدار</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">امتیاز</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="alert in filteredAlertHistory" :key="alert.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <img :src="alert.user.avatar || '/default-avatar.png'" :alt="alert.user.name" class="w-8 h-8 rounded-full mr-3">
                  <div>
                    <div class="text-sm font-medium text-gray-900">{{ alert.user.name }}</div>
                    <div class="text-sm text-gray-500">{{ alert.user.email }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getAlertTypeClass(alert.type)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                  {{ getAlertTypeText(alert.type) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getScoreClass(alert.score)" class="text-sm font-medium">
                  {{ alert.score }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusClass(alert.status)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                  {{ getStatusText(alert.status) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ formatDate(alert.createdAt) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button @click="viewAlertDetails(alert)" class="text-blue-600 hover:text-blue-900">جزئیات</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- قالب‌های پیام هشدار -->
    <div class="bg-white rounded-lg shadow">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-medium text-gray-900">قالب‌های پیام هشدار</h3>
      </div>
      <div class="px-4 py-4">
        <div class="space-y-6">
          <div v-for="template in alertTemplates" :key="template.id" class="border border-gray-200 rounded-lg px-4 py-4">
            <div class="flex items-center justify-between mb-4">
              <h4 class="font-medium text-gray-900">{{ template.name }}</h4>
              <div class="flex space-x-2 space-x-reverse">
                <button @click="editTemplate(template)" class="text-blue-600 hover:text-blue-900">ویرایش</button>
                <button @click="testTemplate(template)" class="text-green-600 hover:text-green-900">تست</button>
              </div>
            </div>
            <div class="space-y-2">
              <div>
                <label class="block text-sm font-medium text-gray-700">موضوع:</label>
                <p class="text-sm text-gray-900">{{ template.subject }}</p>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700">متن پیام:</label>
                <p class="text-sm text-gray-900">{{ template.message }}</p>
              </div>
              <div class="flex items-center space-x-4 space-x-reverse">
                <span class="text-xs text-gray-500">کانال‌های ارسال:</span>
                <span v-if="template.channels.email" class="text-xs bg-blue-100 text-blue-800 px-2 py-1 rounded">ایمیل</span>
                <span v-if="template.channels.sms" class="text-xs bg-green-100 text-green-800 px-2 py-1 rounded">پیامک</span>
                <span v-if="template.channels.push" class="text-xs bg-purple-100 text-purple-800 px-2 py-1 rounded">Push</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Modal جزئیات هشدار -->
    <div v-if="showAlertDetailsModal" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
      <div class="relative top-20 mx-auto p-5 border w-full max-w-md sm:max-w-lg md:max-w-xl shadow-lg rounded-md bg-white">
        <div class="mt-3">
          <h3 class="text-lg font-medium text-gray-900 mb-4">جزئیات هشدار</h3>
          <div v-if="selectedAlert" class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700">کاربر:</label>
              <p class="text-sm text-gray-900">{{ selectedAlert.user.name }}</p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">نوع هشدار:</label>
              <p class="text-sm text-gray-900">{{ getAlertTypeText(selectedAlert.type) }}</p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">امتیاز:</label>
              <p class="text-sm text-gray-900">{{ selectedAlert.score }}</p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">وضعیت:</label>
              <p class="text-sm text-gray-900">{{ getStatusText(selectedAlert.status) }}</p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">تاریخ:</label>
              <p class="text-sm text-gray-900">{{ formatDate(selectedAlert.createdAt) }}</p>
            </div>
            <div v-if="selectedAlert.description">
              <label class="block text-sm font-medium text-gray-700">توضیحات:</label>
              <p class="text-sm text-gray-900">{{ selectedAlert.description }}</p>
            </div>
          </div>
          <div class="flex justify-end mt-6">
            <button @click="showAlertDetailsModal = false" class="px-4 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400">بستن</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, watch } from 'vue';

// Props
const props = defineProps<{
  users: any[]
}>()

// Emits
const emit = defineEmits<{
  sendAlert: [data: any]
  sendBulkAlerts: [data: any]
  exportHistory: [data: any]
}>()

// Reactive data
const alertSystemEnabled = ref(true)
const riskFilter = ref('all')
const alertHistoryFilter = ref('all')
const selectAll = ref(false)
const selectedUsers = ref([])
const showAlertDetailsModal = ref(false)
const selectedAlert = ref(null)

const alertThresholds = ref({
  warning: -20,
  serious: -50,
  blocking: -100
})

const notificationSettings = ref({
  email: true,
  sms: false,
  push: true,
  dashboard: true
})

// Sample data
const riskUsers = ref([
  {
    id: 1,
    name: 'علی احمدی',
    email: 'ali@example.com',
    avatar: '/avatars/ali.jpg',
    score: -25,
    riskLevel: 'warning',
    lastActivity: '2024-01-15T10:30:00Z'
  },
  {
    id: 2,
    name: 'فاطمه محمدی',
    email: 'fateme@example.com',
    avatar: '/avatars/fateme.jpg',
    score: -75,
    riskLevel: 'serious',
    lastActivity: '2024-01-14T15:45:00Z'
  },
  {
    id: 3,
    name: 'محمد رضایی',
    email: 'mohammad@example.com',
    avatar: '/avatars/mohammad.jpg',
    score: -120,
    riskLevel: 'blocking',
    lastActivity: '2024-01-13T09:20:00Z'
  }
])

const alertHistory = ref([
  {
    id: 1,
    user: { name: 'علی احمدی', email: 'ali@example.com', avatar: '/avatars/ali.jpg' },
    type: 'warning',
    score: -25,
    status: 'sent',
    description: 'هشدار اولیه برای امتیاز منفی',
    createdAt: '2024-01-15T10:30:00Z'
  },
  {
    id: 2,
    user: { name: 'فاطمه محمدی', email: 'fateme@example.com', avatar: '/avatars/fateme.jpg' },
    type: 'serious',
    score: -75,
    status: 'sent',
    description: 'هشدار جدی برای امتیاز منفی بالا',
    createdAt: '2024-01-14T15:45:00Z'
  }
])

const alertTemplates = ref([
  {
    id: 1,
    name: 'هشدار اولیه',
    subject: 'هشدار امتیاز منفی',
    message: 'کاربر گرامی، امتیاز شما به دلیل رفتار نامناسب کاهش یافته است. لطفاً قوانین سایت را رعایت کنید.',
    channels: { email: true, sms: false, push: true }
  },
  {
    id: 2,
    name: 'هشدار جدی',
    subject: 'هشدار جدی - امتیاز منفی بالا',
    message: 'کاربر گرامی، امتیاز شما به دلیل تخلفات مکرر به شدت کاهش یافته است. در صورت ادامه، حساب شما مسدود خواهد شد.',
    channels: { email: true, sms: true, push: true }
  },
  {
    id: 3,
    name: 'هشدار مسدودسازی',
    subject: 'هشدار نهایی - مسدودسازی حساب',
    message: 'کاربر گرامی، به دلیل تخلفات مکرر و امتیاز منفی بالا، حساب شما مسدود شده است.',
    channels: { email: true, sms: true, push: true }
  }
])

// Computed properties
const filteredRiskUsers = computed(() => {
  let filtered = riskUsers.value

  if (riskFilter.value !== 'all') {
    filtered = filtered.filter(user => user.riskLevel === riskFilter.value)
  }

  return filtered
})

const filteredAlertHistory = computed(() => {
  let filtered = alertHistory.value

  if (alertHistoryFilter.value !== 'all') {
    filtered = filtered.filter(alert => alert.type === alertHistoryFilter.value)
  }

  return filtered
})

// Methods
const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleDateString('fa-IR')
}

const getScoreClass = (score: number) => {
  if (score >= 0) return 'text-green-600'
  if (score >= -50) return 'text-yellow-600'
  return 'text-red-600'
}

const getRiskLevelClass = (level: string) => {
  switch (level) {
    case 'warning':
      return 'bg-yellow-100 text-yellow-800'
    case 'serious':
      return 'bg-orange-100 text-orange-800'
    case 'blocking':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getRiskLevelText = (level: string) => {
  switch (level) {
    case 'warning':
      return 'هشدار اولیه'
    case 'serious':
      return 'هشدار جدی'
    case 'blocking':
      return 'در آستانه مسدودسازی'
    default:
      return 'نامشخص'
  }
}

const getAlertTypeClass = (type: string) => {
  switch (type) {
    case 'warning':
      return 'bg-yellow-100 text-yellow-800'
    case 'serious':
      return 'bg-orange-100 text-orange-800'
    case 'blocking':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getAlertTypeText = (type: string) => {
  switch (type) {
    case 'warning':
      return 'هشدار اولیه'
    case 'serious':
      return 'هشدار جدی'
    case 'blocking':
      return 'هشدار مسدودسازی'
    default:
      return 'نامشخص'
  }
}

const getStatusClass = (status: string) => {
  switch (status) {
    case 'sent':
      return 'bg-green-100 text-green-800'
    case 'pending':
      return 'bg-yellow-100 text-yellow-800'
    case 'failed':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getStatusText = (status: string) => {
  switch (status) {
    case 'sent':
      return 'ارسال شده'
    case 'pending':
      return 'در انتظار'
    case 'failed':
      return 'ناموفق'
    default:
      return 'نامشخص'
  }
}

const toggleSelectAll = () => {
  if (selectAll.value) {
    selectedUsers.value = filteredRiskUsers.value.map(user => user.id)
  } else {
    selectedUsers.value = []
  }
}

const sendAlert = (user: any) => {
  emit('sendAlert', {
    userId: user.id,
    type: user.riskLevel,
    score: user.score
  })
}

const sendBulkAlerts = () => {
  if (selectedUsers.value.length === 0) {
    alert('لطفاً کاربران را انتخاب کنید')
    return
  }

  emit('sendBulkAlerts', {
    userIds: selectedUsers.value,
    type: 'bulk'
  })
}

const viewUserDetails = (user: any) => {
  // Navigate to user details page
  console.log('مشاهده جزئیات کاربر:', user)
}

const viewAlertDetails = (alert: any) => {
  selectedAlert.value = alert
  showAlertDetailsModal.value = true
}

const exportAlertHistory = () => {
  emit('exportHistory', {
    history: filteredAlertHistory.value,
    filter: alertHistoryFilter.value
  })
}

const editTemplate = (template: any) => {
  console.log('ویرایش قالب:', template)
}

const testTemplate = (template: any) => {
  console.log('تست قالب:', template)
}

// Watchers
watch(selectedUsers, (newValue) => {
  selectAll.value = newValue.length === filteredRiskUsers.value.length && newValue.length > 0
})
</script> 
