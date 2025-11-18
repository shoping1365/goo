<template>
  <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
    <!-- هدر -->
    <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-6 mb-6">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">مدیریت اتصالات حسابداری</h3>
        <p class="text-sm text-gray-600">اتصال و تنظیم نرم‌افزارهای حسابداری و ERP</p>
      </div>
      
      <!-- دکمه افزودن اتصال جدید -->
      <button 
        @click="showAddConnectionModal = true"
        class="inline-flex items-center gap-2 px-4 py-2 bg-green-600 hover:bg-green-700 text-white rounded-lg transition-colors duration-200"
      >
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
        </svg>
        افزودن اتصال جدید
      </button>
    </div>

    <!-- فیلترهای اتصال -->
    <div class="bg-gray-50 rounded-lg p-6 mb-6">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نوع نرم‌افزار</label>
          <select 
            v-model="connectionFilters.softwareType"
            @change="filterConnections"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه انواع</option>
            <option value="helo">هلو</option>
            <option value="parsian">پارسیان</option>
            <option value="sepidar">سپیدار</option>
            <option value="ghias">قیاس</option>
            <option value="other">سایر</option>
          </select>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
          <select 
            v-model="connectionFilters.status"
            @change="filterConnections"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه وضعیت‌ها</option>
            <option value="active">فعال</option>
            <option value="inactive">غیرفعال</option>
            <option value="error">خطا</option>
            <option value="testing">در حال تست</option>
          </select>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">جستجو</label>
          <input 
            v-model="connectionFilters.search"
            type="text"
            placeholder="جستجو در اتصالات..."
            @input="filterConnections"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>
        
        <div class="flex items-end">
          <button 
            @click="clearFilters"
            class="w-full px-4 py-2 bg-gray-200 hover:bg-gray-300 text-gray-700 rounded-lg transition-colors duration-200"
          >
            پاک کردن فیلترها
          </button>
        </div>
      </div>
    </div>

    <!-- جدول اتصالات -->
    <div class="overflow-x-auto">
      <table class="w-full text-sm">
        <thead>
          <tr class="border-b border-gray-200 bg-gray-50">
            <th class="text-right py-3 px-4 font-medium text-gray-600">نرم‌افزار</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">نام اتصال</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">آدرس سرور</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">وضعیت</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">آخرین همگام‌سازی</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">عملیات</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="connection in filteredConnections" :key="connection.id" class="border-b border-gray-100 hover:bg-gray-50">
            <td class="py-3 px-4">
              <div class="flex items-center gap-3">
                <img :src="connection.logo" :alt="connection.softwareName" class="w-8 h-8 rounded-lg" />
                <div>
                  <div class="font-medium text-gray-900">{{ connection.softwareName }}</div>
                  <div class="text-xs text-gray-500">{{ connection.version }}</div>
                </div>
              </div>
            </td>
            <td class="py-3 px-4 text-gray-900">{{ connection.connectionName }}</td>
            <td class="py-3 px-4 text-gray-900 font-mono text-xs">{{ connection.serverUrl }}</td>
            <td class="py-3 px-4">
              <div class="flex items-center gap-2">
                <div 
                  class="w-2 h-2 rounded-full"
                  :class="getStatusColor(connection.status)"
                ></div>
                <span 
                  class="px-2 py-1 rounded-full text-xs font-medium"
                  :class="getStatusClass(connection.status)"
                >
                  {{ getStatusLabel(connection.status) }}
                </span>
              </div>
            </td>
            <td class="py-3 px-4 text-gray-900">{{ connection.lastSync }}</td>
            <td class="py-3 px-4">
              <div class="flex items-center gap-2">
                <button 
                  @click="testConnection(connection)"
                  class="p-1 text-blue-600 hover:text-blue-800"
                  title="تست اتصال"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                  </svg>
                </button>
                <button 
                  @click="editConnection(connection)"
                  class="p-1 text-yellow-600 hover:text-yellow-800"
                  title="ویرایش"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
                  </svg>
                </button>
                <button 
                  @click="toggleConnection(connection)"
                  :class="connection.status === 'active' ? 'text-red-600 hover:text-red-800' : 'text-green-600 hover:text-green-800'"
                  class="p-1"
                  :title="connection.status === 'active' ? 'غیرفعال کردن' : 'فعال کردن'"
                >
                  <svg v-if="connection.status === 'active'" class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728L5.636 5.636m12.728 12.728L18.364 5.636M5.636 18.364l12.728-12.728"></path>
                  </svg>
                  <svg v-else class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
                  </svg>
                </button>
                <button 
                  v-if="canDeleteConnection"
                  @click="deleteConnection(connection)"
                  class="p-1 text-red-600 hover:text-red-800"
                  title="حذف"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                  </svg>
                </button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- مودال افزودن اتصال جدید -->
    <div v-if="showAddConnectionModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-xl p-6 w-full max-w-2xl mx-4 max-h-[90vh] overflow-y-auto">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">افزودن اتصال جدید</h3>
          <button @click="showAddConnectionModal = false" class="text-gray-400 hover:text-gray-600">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>

        <form @submit.prevent="addConnection" class="space-y-4">
          <!-- انتخاب نرم‌افزار -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نرم‌افزار حسابداری</label>
            <select 
              v-model="newConnection.softwareType"
              @change="updateSoftwareInfo"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              required
            >
              <option value="">انتخاب نرم‌افزار</option>
              <option value="helo">هلو</option>
              <option value="parsian">پارسیان</option>
              <option value="sepidar">سپیدار</option>
              <option value="ghias">قیاس</option>
              <option value="other">سایر</option>
            </select>
          </div>

          <!-- نام اتصال -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نام اتصال</label>
            <input 
              v-model="newConnection.connectionName"
              type="text"
              placeholder="مثال: اتصال هلو اصلی"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              required
            />
          </div>

          <!-- آدرس سرور -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">آدرس سرور</label>
            <input 
              v-model="newConnection.serverUrl"
              type="url"
              placeholder="https://example.com/api"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              required
            />
          </div>

          <!-- کلید API -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">کلید API</label>
            <div class="relative">
              <input 
                v-model="newConnection.apiKey"
                :type="showApiKey ? 'text' : 'password'"
                placeholder="کلید API را وارد کنید"
                class="w-full px-3 py-2 pr-10 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                required
              />
              <button 
                @click="showApiKey = !showApiKey"
                type="button"
                class="absolute left-2 top-1/2 transform -translate-y-1/2 text-gray-400 hover:text-gray-600"
              >
                <svg v-if="showApiKey" class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                </svg>
                <svg v-else class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.878 9.878L3 3m6.878 6.878L21 21"></path>
                </svg>
              </button>
            </div>
          </div>

          <!-- تنظیمات اضافی -->
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">پورت</label>
              <input 
                v-model.number="newConnection.port"
                type="number"
                placeholder="80"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">Timeout (ثانیه)</label>
              <input 
                v-model.number="newConnection.timeout"
                type="number"
                placeholder="30"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
          </div>

          <!-- تنظیمات همگام‌سازی -->
          <div class="bg-gray-50 rounded-lg p-6">
            <h4 class="font-medium text-gray-900 mb-3">تنظیمات همگام‌سازی</h4>
            <div class="space-y-3">
              <label class="flex items-center">
                <input 
                  v-model="newConnection.autoSync"
                  type="checkbox"
                  class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                />
                <span class="mr-2 text-sm text-gray-700">همگام‌سازی خودکار</span>
              </label>
              <label class="flex items-center">
                <input 
                  v-model="newConnection.syncInvoices"
                  type="checkbox"
                  class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                />
                <span class="mr-2 text-sm text-gray-700">همگام‌سازی فاکتورها</span>
              </label>
              <label class="flex items-center">
                <input 
                  v-model="newConnection.syncCustomers"
                  type="checkbox"
                  class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                />
                <span class="mr-2 text-sm text-gray-700">همگام‌سازی مشتریان</span>
              </label>
              <label class="flex items-center">
                <input 
                  v-model="newConnection.syncProducts"
                  type="checkbox"
                  class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                />
                <span class="mr-2 text-sm text-gray-700">همگام‌سازی محصولات</span>
              </label>
            </div>
          </div>

          <!-- دکمه‌ها -->
          <div class="flex gap-3 pt-4">
            <button 
              type="submit"
              class="flex-1 px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg transition-colors duration-200"
            >
              افزودن اتصال
            </button>
            <button 
              @click="showAddConnectionModal = false"
              type="button"
              class="flex-1 px-4 py-2 bg-gray-200 hover:bg-gray-300 text-gray-700 rounded-lg transition-colors duration-200"
            >
              انصراف
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';

import { useAuth } from '~/composables/useAuth';

// تعریف definePageMeta برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void

// احراز هویت غیرفعال شده است - محدودیت دسترسی حذف شده
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user, hasPermission } = useAuth()

// Computed برای چک کردن پرمیژن حذف
const canDeleteConnection = computed(() => {
  // TODO: Implement permission check
  return true // Placeholder
})

// فیلترهای اتصال
const connectionFilters = ref({
  softwareType: '',
  status: '',
  search: ''
})

// مودال افزودن اتصال
const showAddConnectionModal = ref(false)
const showApiKey = ref(false)

// اتصال جدید
const newConnection = ref({
  softwareType: '',
  connectionName: '',
  serverUrl: '',
  apiKey: '',
  port: 80,
  timeout: 30,
  autoSync: true,
  syncInvoices: true,
  syncCustomers: true,
  syncProducts: true
})

// اتصالات موجود
const connections = ref([
  {
    id: 1,
    softwareType: 'helo',
    softwareName: 'هلو',
    version: 'نسخه 2024',
    connectionName: 'اتصال هلو اصلی',
    serverUrl: 'https://helo-server.local',
    status: 'active',
    lastSync: '5 دقیقه پیش',
    logo: '/images/helo-logo.png'
  },
  {
    id: 2,
    softwareType: 'parsian',
    softwareName: 'پارسیان',
    version: 'نسخه 2023',
    connectionName: 'اتصال پارسیان',
    serverUrl: 'https://parsian-accounting.com',
    status: 'active',
    lastSync: '10 دقیقه پیش',
    logo: '/images/parsian-logo.png'
  },
  {
    id: 3,
    softwareType: 'sepidar',
    softwareName: 'سپیدار',
    version: 'نسخه 2024',
    connectionName: 'اتصال سپیدار',
    serverUrl: 'https://sepidar-server.local',
    status: 'error',
    lastSync: '1 ساعت پیش',
    logo: '/images/sepidar-logo.png'
  }
])

// اتصالات فیلتر شده
const filteredConnections = computed(() => {
  return connections.value.filter(connection => {
    if (connectionFilters.value.softwareType && connection.softwareType !== connectionFilters.value.softwareType) return false
    if (connectionFilters.value.status && connection.status !== connectionFilters.value.status) return false
    if (connectionFilters.value.search) {
      const search = connectionFilters.value.search.toLowerCase()
      return connection.softwareName.toLowerCase().includes(search) || 
             connection.connectionName.toLowerCase().includes(search) ||
             connection.serverUrl.toLowerCase().includes(search)
    }
    return true
  })
})

// رنگ وضعیت
const getStatusColor = (status: string) => {
  const colors = {
    active: 'bg-green-500',
    inactive: 'bg-gray-500',
    error: 'bg-red-500',
    testing: 'bg-yellow-500'
  }
  return colors[status] || 'bg-gray-500'
}

// کلاس وضعیت
const getStatusClass = (status: string) => {
  const classes = {
    active: 'bg-green-100 text-green-700',
    inactive: 'bg-gray-100 text-gray-700',
    error: 'bg-red-100 text-red-700',
    testing: 'bg-yellow-100 text-yellow-700'
  }
  return classes[status] || 'bg-gray-100 text-gray-700'
}

// برچسب وضعیت
const getStatusLabel = (status: string) => {
  const labels = {
    active: 'فعال',
    inactive: 'غیرفعال',
    error: 'خطا',
    testing: 'در حال تست'
  }
  return labels[status] || status
}

// فیلتر کردن اتصالات
const filterConnections = () => {
  // TODO: اعمال فیلترها
  console.log('فیلترهای اتصال اعمال شد:', connectionFilters.value)
}

// پاک کردن فیلترها
const clearFilters = () => {
  connectionFilters.value = {
    softwareType: '',
    status: '',
    search: ''
  }
}

// بروزرسانی اطلاعات نرم‌افزار
const updateSoftwareInfo = () => {
  // TODO: بروزرسانی اطلاعات نرم‌افزار بر اساس نوع انتخاب شده
  console.log('نوع نرم‌افزار انتخاب شد:', newConnection.value.softwareType)
}

// افزودن اتصال جدید
const addConnection = async () => {
  try {
    // TODO: ارسال درخواست به API
    console.log('اتصال جدید اضافه شد:', newConnection.value)
    
    // ریست فرم
    newConnection.value = {
      softwareType: '',
      connectionName: '',
      serverUrl: '',
      apiKey: '',
      port: 80,
      timeout: 30,
      autoSync: true,
      syncInvoices: true,
      syncCustomers: true,
      syncProducts: true
    }
    
    showAddConnectionModal.value = false
  } catch (error) {
    console.error('خطا در افزودن اتصال:', error)
  }
}

// تست اتصال
const testConnection = async (connection: any) => {
  try {
    // TODO: تست اتصال
    console.log('تست اتصال شروع شد:', connection)
  } catch (error) {
    console.error('خطا در تست اتصال:', error)
  }
}

// ویرایش اتصال
const editConnection = (connection: any) => {
  // TODO: ویرایش اتصال
  console.log('ویرایش اتصال:', connection)
}

// تغییر وضعیت اتصال
const toggleConnection = async (connection: any) => {
  try {
    // TODO: تغییر وضعیت اتصال
    console.log('تغییر وضعیت اتصال:', connection)
  } catch (error) {
    console.error('خطا در تغییر وضعیت اتصال:', error)
  }
}

// حذف اتصال
const deleteConnection = async (connection: any) => {
  if (confirm('آیا از حذف این اتصال اطمینان دارید؟')) {
    try {
      // TODO: حذف اتصال
      console.log('اتصال حذف شد:', connection)
    } catch (error) {
      console.error('خطا در حذف اتصال:', error)
    }
  }
}

// بارگذاری اولیه
onMounted(() => {
  // TODO: بارگذاری اتصالات از API
})
</script>

<!--
  کامپوننت مدیریت اتصالات حسابداری
  شامل:
  1. نمایش لیست اتصالات
  2. افزودن اتصال جدید
  3. ویرایش و حذف اتصالات
  4. تست اتصال
  5. فیلترهای پیشرفته
  6. تنظیمات API
  7. طراحی مدرن و کاملاً ریسپانسیو
--> 

