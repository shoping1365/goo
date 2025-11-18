<template>
  <div class="min-h-screen">
    <!-- Header -->
    <div class="bg-white shadow-sm border-b border-gray-200">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6">
        <div class="flex items-center justify-between">
          <div>
            <h1 class="text-2xl font-bold text-gray-900">مدیریت اعلان‌ها</h1>
            <p class="mt-1 text-sm text-gray-600">مدیریت و تنظیم اعلان‌های سیستم</p>
          </div>
          <div class="flex items-center space-x-3 space-x-reverse">
            <button @click="showCreateModal = true" class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors">
              ایجاد اعلان جدید
            </button>
            <button @click="bulkAction" class="bg-gray-600 text-white px-4 py-2 rounded-md hover:bg-gray-700 transition-colors">
              عملیات گروهی
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Main Content -->
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <!-- فیلترها -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 mb-6">
        <div class="p-6">
          <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نوع اعلان</label>
              <select v-model="filters.type" class="block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50">
                <option value="">همه</option>
                <option value="sms">پیامک</option>
                <option value="email">ایمیل</option>
                <option value="push">اعلان مرورگر</option>
                <option value="system">سیستمی</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
              <select v-model="filters.status" class="block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50">
                <option value="">همه</option>
                <option value="active">فعال</option>
                <option value="inactive">غیرفعال</option>
                <option value="draft">پیش‌نویس</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">اولویت</label>
              <select v-model="filters.priority" class="block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50">
                <option value="">همه</option>
                <option value="high">بالا</option>
                <option value="medium">متوسط</option>
                <option value="low">پایین</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">جستجو</label>
              <input type="text" v-model="filters.search" placeholder="جستجو در عنوان یا محتوا..." class="block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50">
            </div>
          </div>
        </div>
      </div>

      <!-- آمار کلی -->
      <div class="relative mb-6">
        <div class="absolute left-0 right-0 top-2 bottom-2 rounded-2xl shadow-lg py-6" style="z-index:0;"></div>
        <div class="relative grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6" style="z-index:1;">
          <!-- کارت کل اعلان‌ها -->
          <div class="bg-gradient-to-br from-blue-50 to-indigo-100 overflow-hidden shadow-md rounded-lg cursor-pointer border border-blue-200 transition-all duration-200 hover:shadow-lg hover:scale-105">
            <div class="p-3">
              <div class="flex items-center justify-between">
                <div class="flex-1">
                  <dl>
                    <dt class="text-xs font-medium text-blue-600 truncate">کل اعلان‌ها</dt>
                    <dd class="text-base font-bold text-blue-800">{{ stats.totalAlerts }}</dd>
                  </dl>
                </div>
                <div class="flex-shrink-0 mr-3">
                  <div class="w-8 h-8 bg-gradient-to-br from-blue-400 to-indigo-500 rounded-lg flex items-center justify-center shadow-md">
                    <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-5 5v-5zM4 19h6v-2H4v2zM20 4v6h-2V4h2zM4 4v6h2V4H4z"></path>
                    </svg>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- کارت اعلان‌های فعال -->
          <div class="bg-gradient-to-br from-emerald-50 to-green-100 overflow-hidden shadow-md rounded-lg cursor-pointer border border-emerald-200 transition-all duration-200 hover:shadow-lg hover:scale-105">
            <div class="p-3">
              <div class="flex items-center justify-between">
                <div class="flex-1">
                  <dl>
                    <dt class="text-xs font-medium text-emerald-600 truncate">فعال</dt>
                    <dd class="text-base font-bold text-emerald-800">{{ stats.activeAlerts }}</dd>
                  </dl>
                </div>
                <div class="flex-shrink-0 mr-3">
                  <div class="w-8 h-8 bg-gradient-to-br from-emerald-400 to-green-500 rounded-lg flex items-center justify-center shadow-md">
                    <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                    </svg>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- کارت امروز ارسال شده -->
          <div class="bg-gradient-to-br from-amber-50 to-orange-100 overflow-hidden shadow-md rounded-lg cursor-pointer border border-amber-200 transition-all duration-200 hover:shadow-lg hover:scale-105">
            <div class="p-3">
              <div class="flex items-center justify-between">
                <div class="flex-1">
                  <dl>
                    <dt class="text-xs font-medium text-amber-600 truncate">امروز ارسال شده</dt>
                    <dd class="text-base font-bold text-amber-800">{{ stats.todaySent }}</dd>
                  </dl>
                </div>
                <div class="flex-shrink-0 mr-3">
                  <div class="w-8 h-8 bg-gradient-to-br from-amber-400 to-orange-500 rounded-lg flex items-center justify-center shadow-md">
                    <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                    </svg>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- کارت خطاها -->
          <div class="bg-gradient-to-br from-rose-50 to-pink-100 overflow-hidden shadow-md rounded-lg cursor-pointer border border-rose-200 transition-all duration-200 hover:shadow-lg hover:scale-105">
            <div class="p-3">
              <div class="flex items-center justify-between">
                <div class="flex-1">
                  <dl>
                    <dt class="text-xs font-medium text-rose-600 truncate">خطاها</dt>
                    <dd class="text-base font-bold text-rose-800">{{ stats.errors }}</dd>
                  </dl>
                </div>
                <div class="flex-shrink-0 mr-3">
                  <div class="w-8 h-8 bg-gradient-to-br from-rose-400 to-red-500 rounded-lg flex items-center justify-center shadow-md">
                    <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                    </svg>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- جدول اعلان‌ها -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200">
        <div class="px-6 py-4 border-b border-gray-200">
          <div class="flex items-center justify-between">
            <h3 class="text-lg font-medium text-gray-900">لیست اعلان‌ها</h3>
            <div class="flex items-center space-x-3 space-x-reverse">
              <span class="text-sm text-gray-500">{{ filteredAlerts.length }} اعلان</span>
              <button @click="selectAll" class="text-sm text-blue-600 hover:text-blue-800">
                انتخاب همه
              </button>
            </div>
          </div>
        </div>
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                  <input type="checkbox" v-model="selectAllAlerts" @change="toggleSelectAll" class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50">
                </th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عنوان</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">اولویت</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ ایجاد</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="alert in filteredAlerts" :key="alert.id" class="hover:bg-gray-50">
                <td class="px-6 py-4 whitespace-nowrap">
                  <input type="checkbox" v-model="selectedAlerts" :value="alert.id" class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50">
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <div>
                    <div class="text-sm font-medium text-gray-900">{{ alert.title }}</div>
                    <div class="text-sm text-gray-500">{{ alert.description }}</div>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium" :class="getTypeClass(alert.type)">
                    {{ getTypeLabel(alert.type) }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium" :class="getStatusClass(alert.status)">
                    {{ getStatusLabel(alert.status) }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium" :class="getPriorityClass(alert.priority)">
                    {{ getPriorityLabel(alert.priority) }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ alert.createdAt }}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <button @click="editAlert(alert)" class="text-blue-600 hover:text-blue-900">ویرایش</button>
                    <button @click="toggleAlertStatus(alert)" class="text-green-600 hover:text-green-900">
                      {{ alert.status === 'active' ? 'غیرفعال' : 'فعال' }}
                    </button>
                    <button 
                  v-if="canDeleteAlert"
                  @click="deleteAlert(alert.id)" 
                  class="text-red-600 hover:text-red-900"
                >
                  حذف
                </button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Pagination -->
      <Pagination
        v-if="totalAlerts > 0"
        :current-page="currentPage"
        :total-pages="totalPages"
        :total="totalAlerts"
        :items-per-page="pageSize"
        @page-changed="handlePageChange"
        @items-per-page-changed="handleItemsPerPageChange"
      />
    </div>

    <!-- Modal ایجاد/ویرایش اعلان -->
    <div v-if="showCreateModal" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
      <div class="relative top-20 mx-auto p-5 border w-full max-w-md sm:max-w-lg md:max-w-xl shadow-lg rounded-md bg-white">
        <div class="mt-3">
          <h3 class="text-lg font-medium text-gray-900 mb-4">{{ editingAlert ? 'ویرایش اعلان' : 'ایجاد اعلان جدید' }}</h3>
          <form @submit.prevent="saveAlert" class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">عنوان</label>
              <input type="text" v-model="alertForm.title" required class="block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات</label>
              <textarea v-model="alertForm.description" rows="3" class="block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"></textarea>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نوع اعلان</label>
              <select v-model="alertForm.type" required class="block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50">
                <option value="sms">پیامک</option>
                <option value="email">ایمیل</option>
                <option value="push">اعلان مرورگر</option>
                <option value="system">سیستمی</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">اولویت</label>
              <select v-model="alertForm.priority" required class="block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50">
                <option value="low">پایین</option>
                <option value="medium">متوسط</option>
                <option value="high">بالا</option>
              </select>
            </div>
            <div>
              <label class="flex items-center">
                <input type="checkbox" v-model="alertForm.isActive" class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50">
                <span class="mr-3 text-sm text-gray-900">فعال</span>
              </label>
            </div>
            <div class="flex items-center justify-end space-x-3 space-x-reverse">
              <button type="button" @click="showCreateModal = false" class="px-4 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400">
                انصراف
              </button>
              <button type="submit" class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700">
                {{ editingAlert ? 'ویرایش' : 'ایجاد' }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
</script>

<script setup lang="ts">
import { ref, reactive, computed } from 'vue'
import { useAuth } from '~/composables/useAuth'
import Pagination from '~/components/admin/common/Pagination.vue'

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user, hasPermission } = useAuth()

// تعریف interface ها
interface Alert {
  id: number
  title: string
  description: string
  type: 'sms' | 'email' | 'push' | 'system'
  status: 'active' | 'inactive' | 'draft'
  priority: 'low' | 'medium' | 'high'
  createdAt: string
  isActive: boolean
}

interface AlertForm {
  title: string
  description: string
  type: 'sms' | 'email' | 'push' | 'system'
  priority: 'low' | 'medium' | 'high'
  isActive: boolean
}

interface Filters {
  type: string
  status: string
  priority: string
  search: string
}

interface Stats {
  totalAlerts: number
  activeAlerts: number
  todaySent: number
  errors: number
}

// متغیرهای reactive
const alerts = ref<Alert[]>([
  {
    id: 1,
    title: 'اعلان موجودی محصول',
    description: 'اعلان خودکار برای موجودی محصولات',
    type: 'sms',
    status: 'active',
    priority: 'high',
    createdAt: '2024/01/15',
    isActive: true
  },
  {
    id: 2,
    title: 'تایید سفارش',
    description: 'اعلان تایید سفارش برای مشتریان',
    type: 'email',
    status: 'active',
    priority: 'medium',
    createdAt: '2024/01/14',
    isActive: true
  },
  {
    id: 3,
    title: 'تخفیف ویژه',
    description: 'اعلان تخفیف‌های ویژه',
    type: 'push',
    status: 'draft',
    priority: 'low',
    createdAt: '2024/01/13',
    isActive: false
  },
  {
    id: 4,
    title: 'خطای سیستم',
    description: 'اعلان خطاهای سیستمی',
    type: 'system',
    status: 'active',
    priority: 'high',
    createdAt: '2024/01/12',
    isActive: true
  }
])

const stats = reactive<Stats>({
  totalAlerts: 4,
  activeAlerts: 3,
  todaySent: 156,
  errors: 2
})

const filters = reactive<Filters>({
  type: '',
  status: '',
  priority: '',
  search: ''
})

const selectedAlerts = ref<number[]>([])
const selectAllAlerts = ref(false)
const showCreateModal = ref(false)
const editingAlert = ref<Alert | null>(null)
const currentPage = ref(1)
const pageSize = ref(10)

const alertForm = reactive<AlertForm>({
  title: '',
  description: '',
  type: 'sms',
  priority: 'medium',
  isActive: true
})

// Computed properties
const filteredAlerts = computed(() => {
  let filtered = alerts.value

  if (filters.type) {
    filtered = filtered.filter(alert => alert.type === filters.type)
  }

  if (filters.status) {
    filtered = filtered.filter(alert => alert.status === filters.status)
  }

  if (filters.priority) {
    filtered = filtered.filter(alert => alert.priority === filters.priority)
  }

  if (filters.search) {
    const search = filters.search.toLowerCase()
    filtered = filtered.filter(alert => 
      alert.title.toLowerCase().includes(search) || 
      alert.description.toLowerCase().includes(search)
    )
  }

  return filtered
})

const totalAlerts = computed(() => filteredAlerts.value.length)
const totalPages = computed(() => Math.ceil(totalAlerts.value / pageSize.value))

// Computed برای چک کردن پرمیژن حذف
const canDeleteAlert = computed(() => hasPermission('alert.delete'))

// متدها
const getTypeClass = (type: string) => {
  switch (type) {
    case 'sms':
      return 'bg-blue-100 text-blue-800'
    case 'email':
      return 'bg-green-100 text-green-800'
    case 'push':
      return 'bg-purple-100 text-purple-800'
    case 'system':
      return 'bg-gray-100 text-gray-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getTypeLabel = (type: string) => {
  switch (type) {
    case 'sms':
      return 'پیامک'
    case 'email':
      return 'ایمیل'
    case 'push':
      return 'اعلان مرورگر'
    case 'system':
      return 'سیستمی'
    default:
      return type
  }
}

const getStatusClass = (status: string) => {
  switch (status) {
    case 'active':
      return 'bg-green-100 text-green-800'
    case 'inactive':
      return 'bg-red-100 text-red-800'
    case 'draft':
      return 'bg-yellow-100 text-yellow-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getStatusLabel = (status: string) => {
  switch (status) {
    case 'active':
      return 'فعال'
    case 'inactive':
      return 'غیرفعال'
    case 'draft':
      return 'پیش‌نویس'
    default:
      return status
  }
}

const getPriorityClass = (priority: string) => {
  switch (priority) {
    case 'high':
      return 'bg-red-100 text-red-800'
    case 'medium':
      return 'bg-yellow-100 text-yellow-800'
    case 'low':
      return 'bg-green-100 text-green-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getPriorityLabel = (priority: string) => {
  switch (priority) {
    case 'high':
      return 'بالا'
    case 'medium':
      return 'متوسط'
    case 'low':
      return 'پایین'
    default:
      return priority
  }
}

const toggleSelectAll = () => {
  if (selectAllAlerts.value) {
    selectedAlerts.value = filteredAlerts.value.map(alert => alert.id)
  } else {
    selectedAlerts.value = []
  }
}

const selectAll = () => {
  selectAllAlerts.value = true
  selectedAlerts.value = filteredAlerts.value.map(alert => alert.id)
}

const editAlert = (alert: Alert) => {
  editingAlert.value = alert
  alertForm.title = alert.title
  alertForm.description = alert.description
  alertForm.type = alert.type
  alertForm.priority = alert.priority
  alertForm.isActive = alert.isActive
  showCreateModal.value = true
}

const saveAlert = async () => {
  try {
    if (editingAlert.value) {
      // ویرایش اعلان موجود
      const index = alerts.value.findIndex(a => a.id === editingAlert.value!.id)
      if (index !== -1) {
        alerts.value[index] = {
          ...alerts.value[index],
          ...alertForm,
          status: alertForm.isActive ? 'active' : 'inactive'
        }
      }
    } else {
      // ایجاد اعلان جدید
      const newAlert: Alert = {
        id: Date.now(),
        ...alertForm,
        status: alertForm.isActive ? 'active' : 'inactive',
        createdAt: new Date().toLocaleDateString('fa-IR')
      }
      alerts.value.unshift(newAlert)
    }

    showCreateModal.value = false
    editingAlert.value = null
    resetAlertForm()
    console.log('اعلان با موفقیت ذخیره شد')
  } catch (error) {
    console.error('خطا در ذخیره اعلان:', error)
    console.log('خطا در ذخیره اعلان')
  }
}

const resetAlertForm = () => {
  alertForm.title = ''
  alertForm.description = ''
  alertForm.type = 'sms'
  alertForm.priority = 'medium'
  alertForm.isActive = true
}

const toggleAlertStatus = async (alert: Alert) => {
  try {
    const newStatus = alert.status === 'active' ? 'inactive' : 'active'
    const index = alerts.value.findIndex(a => a.id === alert.id)
    if (index !== -1) {
      alerts.value[index].status = newStatus
      alerts.value[index].isActive = newStatus === 'active'
    }
    console.log(`وضعیت اعلان به ${getStatusLabel(newStatus)} تغییر یافت`)
  } catch (error) {
    console.error('خطا در تغییر وضعیت:', error)
    console.log('خطا در تغییر وضعیت')
  }
}

const deleteAlert = async (alertId: number) => {
  if (confirm('آیا مطمئن هستید که می‌خواهید این اعلان را حذف کنید؟')) {
    try {
      alerts.value = alerts.value.filter(alert => alert.id !== alertId)
      selectedAlerts.value = selectedAlerts.value.filter(id => id !== alertId)
      console.log('اعلان با موفقیت حذف شد')
    } catch (error) {
      console.error('خطا در حذف اعلان:', error)
      console.log('خطا در حذف اعلان')
    }
  }
}

const bulkAction = () => {
  if (selectedAlerts.value.length === 0) {
    console.log('لطفاً ابتدا اعلان‌هایی را انتخاب کنید')
    return
  }

  const action = prompt('عملیات مورد نظر را انتخاب کنید:\n1. فعال کردن\n2. غیرفعال کردن\n3. حذف')
  
  switch (action) {
    case '1':
      selectedAlerts.value.forEach(id => {
        const alertItem = alerts.value.find(a => a.id === id)
        if (alertItem) {
          alertItem.status = 'active'
          alertItem.isActive = true
        }
      })
      break
    case '2':
      selectedAlerts.value.forEach(id => {
        const alertItem = alerts.value.find(a => a.id === id)
        if (alertItem) {
          alertItem.status = 'inactive'
          alertItem.isActive = false
        }
      })
      break
    case '3':
      if (confirm('آیا مطمئن هستید که می‌خواهید اعلان‌های انتخاب شده را حذف کنید؟')) {
        alerts.value = alerts.value.filter(alertItem => !selectedAlerts.value.includes(alertItem.id))
        selectedAlerts.value = []
      }
      break
  }
}

const handlePageChange = (newPage: number) => {
  currentPage.value = newPage
}

const handleItemsPerPageChange = (newItemsPerPage: number) => {
  pageSize.value = newItemsPerPage
  currentPage.value = 1 // بازگشت به صفحه اول
}
</script>

