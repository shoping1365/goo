<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200">
    <!-- هدر بخش -->
    <div class="p-6 border-b border-gray-200">
      <div class="flex items-center justify-between">
        <div>
          <h2 class="text-lg font-semibold text-gray-900">یکپارچه‌سازی با سیستم‌های دیگر</h2>
          <p class="text-gray-600 mt-1">مدیریت اتصال و همگام‌سازی با سرویس‌های خارجی (API، CRM، پیامک، ایمیل و ...)
          </p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button @click="showIntegrationForm = true" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors">
            <span class="i-heroicons-plus ml-2"></span>
            افزودن یکپارچه‌سازی
          </button>
        </div>
      </div>
    </div>

    <!-- آمار یکپارچه‌سازی -->
    <div class="p-6 border-b border-gray-200">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div class="bg-blue-50 p-6 rounded-lg">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center ml-3">
              <span class="i-heroicons-link text-blue-600 text-xl"></span>
            </div>
            <div>
              <p class="text-sm text-blue-600">اتصال‌های فعال</p>
              <p class="text-2xl font-bold text-blue-900">{{ stats.activeIntegrations }}</p>
            </div>
          </div>
        </div>
        <div class="bg-green-50 p-6 rounded-lg">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-green-100 rounded-lg flex items-center justify-center ml-3">
              <span class="i-heroicons-server-stack text-green-600 text-xl"></span>
            </div>
            <div>
              <p class="text-sm text-green-600">سرویس‌های متصل</p>
              <p class="text-2xl font-bold text-green-900">{{ stats.connectedServices }}</p>
            </div>
          </div>
        </div>
        <div class="bg-purple-50 p-6 rounded-lg">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-purple-100 rounded-lg flex items-center justify-center ml-3">
              <span class="i-heroicons-envelope text-purple-600 text-xl"></span>
            </div>
            <div>
              <p class="text-sm text-purple-600">ارسال پیام موفق</p>
              <p class="text-2xl font-bold text-purple-900">{{ stats.successfulMessages }}</p>
            </div>
          </div>
        </div>
        <div class="bg-orange-50 p-6 rounded-lg">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-orange-100 rounded-lg flex items-center justify-center ml-3">
              <span class="i-heroicons-exclamation-triangle text-orange-600 text-xl"></span>
            </div>
            <div>
              <p class="text-sm text-orange-600">خطاهای اتصال</p>
              <p class="text-2xl font-bold text-orange-900">{{ stats.connectionErrors }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- لیست یکپارچه‌سازی‌ها -->
    <div class="overflow-x-auto">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نام سرویس</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">آخرین همگام‌سازی</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-for="integration in integrations" :key="integration.id" class="hover:bg-gray-50">
            <td class="px-6 py-4 whitespace-nowrap">
              <div>
                <div class="text-sm font-medium text-gray-900">{{ integration.name }}</div>
                <div class="text-sm text-gray-500">{{ integration.description }}</div>
              </div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span :class="`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${getTypeBadgeClass(integration.type)}`">
                {{ getTypeName(integration.type) }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span :class="`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${getStatusBadgeClass(integration.status)}`">
                {{ getStatusName(integration.status) }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
              {{ formatDateTime(integration.lastSync) }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
              <div class="flex items-center space-x-2 space-x-reverse">
                <button @click="editIntegration(integration)" class="text-blue-600 hover:text-blue-900">
                  <span class="i-heroicons-pencil-square"></span>
                </button>
                <button @click="syncIntegration(integration)" class="text-green-600 hover:text-green-900">
                  <span class="i-heroicons-arrow-path"></span>
                </button>
                <button @click="deleteIntegration(integration)" class="text-red-600 hover:text-red-900">
                  <span class="i-heroicons-trash"></span>
                </button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- پیام عدم وجود نتیجه -->
    <div v-if="integrations.length === 0" class="text-center py-12">
      <div class="w-16 h-16 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-4">
        <span class="i-heroicons-link text-gray-400 text-xl"></span>
      </div>
      <h3 class="text-lg font-medium text-gray-900 mb-2">هیچ یکپارچه‌سازی یافت نشد</h3>
      <p class="text-gray-600">هنوز هیچ سرویس خارجی متصل نشده است.</p>
    </div>

    <!-- مودال فرم یکپارچه‌سازی -->
    <div v-if="showIntegrationForm" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg shadow-xl max-w-2xl w-full mx-4 max-h-[90vh] overflow-y-auto">
        <div class="p-6 border-b border-gray-200">
          <div class="flex items-center justify-between">
            <h3 class="text-lg font-semibold text-gray-900">
              {{ editingIntegration ? 'ویرایش یکپارچه‌سازی' : 'افزودن یکپارچه‌سازی جدید' }}
            </h3>
            <button @click="closeForm" class="text-gray-400 hover:text-gray-600">
              <span class="i-heroicons-x-mark text-xl"></span>
            </button>
          </div>
        </div>
        <form @submit.prevent="handleSubmit" class="p-6 space-y-6">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نام سرویس *</label>
              <input v-model="form.name" type="text" required class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="مثال: پیامک ملی، Mailgun، HubSpot">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نوع سرویس *</label>
              <select v-model="form.type" required class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                <option value="sms">پیامک</option>
                <option value="email">ایمیل</option>
                <option value="crm">CRM</option>
                <option value="erp">ERP</option>
                <option value="webhook">وب‌هوک</option>
                <option value="custom">سفارشی</option>
              </select>
            </div>
            <div class="md:col-span-2">
              <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات</label>
              <textarea v-model="form.description" rows="2" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="توضیحات سرویس"></textarea>
            </div>
            <div class="md:col-span-2">
              <label class="block text-sm font-medium text-gray-700 mb-2">آدرس/کلید API *</label>
              <input v-model="form.apiKey" type="text" required class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="API Key یا Endpoint">
            </div>
            <div class="md:col-span-2">
              <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
              <select v-model="form.status" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                <option value="active">فعال</option>
                <option value="inactive">غیرفعال</option>
              </select>
            </div>
          </div>
        </form>
        <div class="p-6 border-t border-gray-200 flex justify-end space-x-3 space-x-reverse">
          <button @click="closeForm" class="px-6 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200 transition-colors">
            انصراف
          </button>
          <button @click="handleSubmit" class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors">
            ذخیره
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'

const showIntegrationForm = ref(false)
const editingIntegration = ref<any>(null)
const stats = ref({
  activeIntegrations: 3,
  connectedServices: 5,
  successfulMessages: 1247,
  connectionErrors: 12
})

const integrations = ref([
  {
    id: 1,
    name: 'پیامک ملی',
    description: 'ارسال پیامک به کاربران',
    type: 'sms',
    status: 'active',
    lastSync: '2024-01-15T10:30:00'
  },
  {
    id: 2,
    name: 'Mailgun',
    description: 'ارسال ایمیل‌های اطلاع‌رسانی',
    type: 'email',
    status: 'active',
    lastSync: '2024-01-14T09:15:00'
  },
  {
    id: 3,
    name: 'HubSpot',
    description: 'یکپارچه‌سازی با CRM',
    type: 'crm',
    status: 'inactive',
    lastSync: '2024-01-10T14:20:00'
  }
])

const form = reactive({
  name: '',
  type: 'sms',
  description: '',
  apiKey: '',
  status: 'active'
})

const getTypeName = (type: string): string => {
  switch (type) {
    case 'sms': return 'پیامک'
    case 'email': return 'ایمیل'
    case 'crm': return 'CRM'
    case 'erp': return 'ERP'
    case 'webhook': return 'وب‌هوک'
    case 'custom': return 'سفارشی'
    default: return type
  }
}

const getTypeBadgeClass = (type: string): string => {
  switch (type) {
    case 'sms': return 'bg-blue-100 text-blue-700'
    case 'email': return 'bg-green-100 text-green-700'
    case 'crm': return 'bg-purple-100 text-purple-700'
    case 'erp': return 'bg-orange-100 text-orange-700'
    case 'webhook': return 'bg-pink-100 text-pink-700'
    case 'custom': return 'bg-gray-100 text-gray-700'
    default: return 'bg-gray-100 text-gray-700'
  }
}

const getStatusName = (status: string): string => {
  switch (status) {
    case 'active': return 'فعال'
    case 'inactive': return 'غیرفعال'
    default: return status
  }
}

const getStatusBadgeClass = (status: string): string => {
  switch (status) {
    case 'active': return 'bg-green-100 text-green-700'
    case 'inactive': return 'bg-gray-100 text-gray-700'
    default: return 'bg-gray-100 text-gray-700'
  }
}

const formatDateTime = (date: string): string => {
  return new Intl.DateTimeFormat('fa-IR', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit'
  }).format(new Date(date))
}

const editIntegration = (integration: any) => {
  editingIntegration.value = integration
  Object.assign(form, integration)
  showIntegrationForm.value = true
}

const syncIntegration = async (integration: any) => {
  // TODO: فراخوانی API برای همگام‌سازی
  alert(`همگام‌سازی با سرویس ${integration.name}`)
}

const deleteIntegration = async (integration: any) => {
  if (confirm(`آیا از حذف یکپارچه‌سازی "${integration.name}" اطمینان دارید؟`)) {
    try {
      // TODO: فراخوانی API برای حذف
      const index = integrations.value.findIndex(i => i.id === integration.id)
      if (index !== -1) {
        integrations.value.splice(index, 1)
      }
    } catch (error) {
      console.error('خطا در حذف یکپارچه‌سازی:', error)
    }
  }
}

const handleSubmit = async () => {
  if (!form.name || !form.type || !form.apiKey) {
    alert('لطفاً فیلدهای اجباری را پر کنید')
    return
  }
  if (editingIntegration.value) {
    Object.assign(editingIntegration.value, form)
  } else {
    integrations.value.unshift({
      id: Date.now(),
      ...form,
      lastSync: new Date().toISOString()
    })
  }
  closeForm()
}

const closeForm = () => {
  showIntegrationForm.value = false
  editingIntegration.value = null
  Object.assign(form, { name: '', type: 'sms', description: '', apiKey: '', status: 'active' })
}

onMounted(async () => {
  // TODO: فراخوانی API برای دریافت یکپارچه‌سازی‌ها
})
</script>

<!--
  کامپوننت یکپارچه‌سازی با سیستم‌های دیگر
  شامل:
  1. مدیریت اتصال به سرویس‌های پیامک، ایمیل، CRM، ERP، وب‌هوک و ...
  2. افزودن، ویرایش، حذف و همگام‌سازی سرویس‌ها
  3. آمار اتصال‌ها و پیام‌های موفق/ناموفق
  توضیحات کامل به فارسی در کد
--> 
