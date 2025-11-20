<template>
  <div class="p-6" dir="rtl">
    <div class="mb-6 bg-white p-6 rounded-lg shadow-sm border border-gray-200">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-2xl font-bold text-gray-900 mb-2">تایید دو مرحله‌ای</h1>
          <p class="text-gray-600">مدیریت احراز هویت دو مرحله‌ای و امنیت حساب‌های کاربری</p>
        </div>
        <div class="flex items-center gap-3">
          <button 
            class="px-4 py-2 text-blue-600 bg-blue-50 border-2 border-blue-200 rounded-lg hover:bg-blue-100 hover:border-blue-300 transition-all duration-200 flex items-center shadow-sm hover:shadow-md" 
            @click="navigateToTwoFactorLogin"
          >
            <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"></path>
            </svg>
            ورود دو مرحله‌ای
          </button>
        <button 
          class="px-4 py-2 text-green-600 bg-green-50 border-2 border-green-200 rounded-lg hover:bg-green-100 hover:border-green-300 transition-all duration-200 flex items-center shadow-sm hover:shadow-md" 
          @click="exportTwoFactorData"
        >
          <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
          </svg>
          خروجی اکسل
        </button>
        </div>
      </div>
    </div>

    <!-- آمار کلی -->
    <div class="relative mb-8">
      <div class="relative grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6" style="z-index:1;">
        <TemplateCard
          title="فعال شده"
          :value="twoFactorStats.enabled"
          variant="green"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"></path>
            </svg>
          </template>
        </TemplateCard>
        <TemplateCard
          title="غیرفعال"
          :value="twoFactorStats.disabled"
          variant="red"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </template>
        </TemplateCard>
        <TemplateCard
          title="در انتظار تایید"
          :value="twoFactorStats.pending"
          variant="yellow"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </template>
        </TemplateCard>
        <TemplateCard
          title="کل کاربران"
          :value="twoFactorStats.total"
          variant="cyan"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197m13.5-9a2.5 2.5 0 11-5 0 2.5 2.5 0 015 0z"></path>
            </svg>
          </template>
        </TemplateCard>
      </div>
    </div>

    <!-- فیلترها و جستجوی پیشرفته -->
    <div class="bg-gradient-to-br from-blue-50 to-cyan-50 shadow-lg rounded-2xl mb-8 border border-blue-100">
      <div class="px-6 py-4 border-b border-blue-100 flex items-center gap-2">
        <div class="w-8 h-8 bg-gradient-to-br from-purple-400 to-purple-600 rounded-lg flex items-center justify-center shadow-md">
          <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.707A1 1 0 013 7V4z"></path>
          </svg>
        </div>
        <h3 class="text-lg font-bold text-blue-900">فیلترها و جستجوی پیشرفته</h3>
      </div>
      <div class="p-6">
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-6 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نام کاربر</label>
            <input v-model="filters.user" type="text" placeholder="جستجو..." class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
            <select v-model="filters.status" class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm">
              <option value="">همه</option>
              <option value="enabled">فعال</option>
              <option value="disabled">غیرفعال</option>
              <option value="pending">در انتظار</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع احراز هویت</label>
            <select v-model="filters.authType" class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm">
              <option value="">همه</option>
              <option value="sms">پیامک</option>
              <option value="email">ایمیل</option>
              <option value="app">اپلیکیشن</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ فعال‌سازی</label>
            <input v-model="filters.activationDate" type="date" class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">آخرین ورود</label>
            <input v-model="filters.lastLogin" type="date" class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نقش کاربر</label>
            <select v-model="filters.role" class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm">
              <option value="">همه</option>
              <option value="admin">مدیر</option>
              <option value="user">کاربر</option>
              <option value="moderator">ناظر</option>
            </select>
          </div>
        </div>
      </div>
    </div>

    <!-- جدول تایید دو مرحله‌ای -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">مدیریت تایید دو مرحله‌ای</h3>
      </div>
      <div class="overflow-x-auto overflow-y-auto custom-scrollbar" :style="{ maxHeight: `${itemsPerPage * 60 + 100}px` }">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کاربر</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع احراز هویت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ فعال‌سازی</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">آخرین ورود</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نقش</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="user in paginatedTwoFactorUsers" :key="user.id" :class="getUserRowClass(user)">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="flex-shrink-0 h-8 w-8">
                    <div class="h-8 w-8 rounded-full bg-gray-300 flex items-center justify-center">
                      <span class="text-sm font-medium text-gray-700">{{ user.name.charAt(0) }}</span>
                    </div>
                  </div>
                  <div class="mr-3">
                    <div class="text-sm font-medium text-gray-900">{{ user.name }}</div>
                    <div class="text-sm text-gray-500">{{ user.email }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusClass(user.status)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ getStatusText(user.status) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                <div class="flex items-center">
                  <span class="mr-2">{{ getAuthTypeIcon(user.authType) }}</span>
                  {{ getAuthTypeText(user.authType) }}
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ user.activationDate || '-' }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ user.lastLogin || '-' }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ getRoleText(user.role) }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <div class="flex flex-col items-center space-y-1">
                  <button 
                    class="text-blue-600 hover:text-blue-900 transition-colors" 
                    title="مشاهده جزئیات"
                    @click="viewUserDetails(user)"
                  >
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                    </svg>
                  </button>
                  <button 
                    v-if="user.status === 'disabled'" 
                    class="text-green-600 hover:text-green-900 transition-colors" 
                    title="فعال کردن تایید دو مرحله‌ای"
                    @click="enableTwoFactor(user)"
                  >
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"></path>
                    </svg>
                  </button>
                  <button 
                    v-if="user.status === 'enabled'" 
                    class="text-red-600 hover:text-red-900 transition-colors" 
                    title="غیرفعال کردن تایید دو مرحله‌ای"
                    @click="disableTwoFactor(user)"
                  >
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      
      <!-- کامپوننت صفحه‌بندی -->
      <AdminPagination 
        :current-page="currentPage"
        :total-pages="totalPages"
        :total="twoFactorUsers.length"
        :items-per-page="itemsPerPage"
        @page-changed="handlePageChange"
        @items-per-page-changed="handleItemsPerPageChange"
      />
    </div>

    <!-- مودال ورود دو مرحله‌ای -->
    <div v-if="showTwoFactorLoginModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50" @click="showTwoFactorLoginModal = false">
      <div class="bg-white rounded-lg shadow-xl max-w-md w-full mx-4" @click.stop>
        <div class="px-6 py-4 border-b border-gray-200">
          <div class="flex items-center justify-between">
            <h3 class="text-lg font-semibold text-gray-900">ورود دو مرحله‌ای</h3>
            <button 
              class="text-gray-400 hover:text-gray-600 transition-colors"
              @click="showTwoFactorLoginModal = false"
            >
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
              </svg>
            </button>
          </div>
        </div>
        
        <div class="px-6 py-4">
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700 mb-2">کد تایید</label>
            <input 
              type="text" 
              placeholder="کد 6 رقمی را وارد کنید"
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-center text-lg tracking-widest"
              maxlength="6"
            >
          </div>
          
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700 mb-2">روش تایید</label>
            <select class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500">
              <option value="sms">پیامک</option>
              <option value="email">ایمیل</option>
              <option value="app">اپلیکیشن احراز هویت</option>
            </select>
          </div>
          
          <div class="flex items-center justify-between text-sm text-gray-600 mb-4">
            <span>زمان باقی‌مانده: <span class="font-mono">02:45</span></span>
            <button class="text-blue-600 hover:text-blue-800 transition-colors">ارسال مجدد</button>
          </div>
        </div>
        
        <div class="px-6 py-4 bg-gray-50 rounded-b-lg flex gap-3">
          <button 
            class="flex-1 px-4 py-2 text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50 transition-colors"
            @click="showTwoFactorLoginModal = false"
          >
            انصراف
          </button>
          <button 
            class="flex-1 px-4 py-2 text-white bg-blue-600 rounded-md hover:bg-blue-700 transition-colors"
          >
            تایید ورود
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import TemplateCard from '@/components/common/TemplateCard.vue'
definePageMeta({ layout: 'admin-main' })

// آمار تایید دو مرحله‌ای
const twoFactorStats = ref({
  enabled: 156,
  disabled: 89,
  pending: 12,
  total: 257
})

// متغیرهای صفحه‌بندی
const currentPage = ref(1)
const itemsPerPage = ref(10)

// متغیر کنترل نمایش مودال ورود دو مرحله‌ای
const showTwoFactorLoginModal = ref(false)

const filters = ref({
  user: '',
  status: '',
  authType: '',
  activationDate: '',
  lastLogin: '',
  role: ''
})

// محاسبه داده‌های صفحه‌بندی شده
const paginatedTwoFactorUsers = computed(() => {
  const startIndex = (currentPage.value - 1) * itemsPerPage.value
  const endIndex = startIndex + itemsPerPage.value
  return twoFactorUsers.value.slice(startIndex, endIndex)
})

// محاسبه تعداد کل صفحات
const totalPages = computed(() => {
  return Math.ceil(twoFactorUsers.value.length / itemsPerPage.value)
})

const twoFactorUsers = ref([
  {
    id: 1,
    name: 'علی احمدی',
    email: 'ali@example.com',
    status: 'enabled',
    authType: 'sms',
    activationDate: '1402/10/15',
    lastLogin: '1402/10/20 14:30',
    role: 'admin'
  },
  {
    id: 2,
    name: 'فاطمه محمدی',
    email: 'fateme@example.com',
    status: 'disabled',
    authType: 'email',
    activationDate: null,
    lastLogin: '1402/10/19 09:15',
    role: 'user'
  },
  {
    id: 3,
    name: 'محمد رضایی',
    email: 'mohammad@example.com',
    status: 'enabled',
    authType: 'app',
    activationDate: '1402/10/10',
    lastLogin: '1402/10/20 16:45',
    role: 'moderator'
  },
  {
    id: 4,
    name: 'سارا کریمی',
    email: 'sara@example.com',
    status: 'pending',
    authType: 'sms',
    activationDate: null,
    lastLogin: '1402/10/18 11:20',
    role: 'user'
  },
  {
    id: 5,
    name: 'احمد نوری',
    email: 'ahmad@example.com',
    status: 'enabled',
    authType: 'email',
    activationDate: '1402/10/05',
    lastLogin: '1402/10/20 13:10',
    role: 'admin'
  },
  {
    id: 6,
    name: 'زهرا صالحی',
    email: 'zahra@example.com',
    status: 'disabled',
    authType: 'app',
    activationDate: null,
    lastLogin: '1402/10/17 15:30',
    role: 'user'
  }
])

function getStatusClass(status) {
  switch (status) {
    case 'enabled':
      return 'bg-green-100 text-green-800'
    case 'disabled':
      return 'bg-red-100 text-red-800'
    case 'pending':
      return 'bg-yellow-100 text-yellow-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

function getStatusText(status) {
  switch (status) {
    case 'enabled':
      return 'فعال'
    case 'disabled':
      return 'غیرفعال'
    case 'pending':
      return 'در انتظار'
    default:
      return 'نامشخص'
  }
}

function getAuthTypeIcon(authType) {
  switch (authType) {
    case 'sms':
      return '📱'
    case 'email':
      return '📧'
    case 'app':
      return '🔐'
    default:
      return '❓'
  }
}

function getAuthTypeText(authType) {
  switch (authType) {
    case 'sms':
      return 'پیامک'
    case 'email':
      return 'ایمیل'
    case 'app':
      return 'اپلیکیشن'
    default:
      return 'نامشخص'
  }
}

function getRoleText(role) {
  switch (role) {
    case 'admin':
      return 'مدیر'
    case 'user':
      return 'کاربر'
    case 'moderator':
      return 'ناظر'
    default:
      return 'نامشخص'
  }
}

function getUserRowClass(user) {
  if (user.status === 'pending') {
    return 'bg-yellow-50'
  }
  return ''
}

function exportTwoFactorData() {
  // منطق خروجی اکسل
  alert('داده‌های تایید دو مرحله‌ای صادر شد')
}

function navigateToTwoFactorLogin() {
  // باز کردن مودال ورود دو مرحله‌ای
  showTwoFactorLoginModal.value = true
}

function viewUserDetails(user) {
  // منطق نمایش جزئیات کاربر
  alert(`جزئیات کاربر ${user.name} نمایش داده می‌شود`)
}

function enableTwoFactor(user) {
  // منطق فعال کردن تایید دو مرحله‌ای
  user.status = 'enabled'
  user.activationDate = new Date().toLocaleDateString('fa-IR')
  alert(`تایید دو مرحله‌ای برای ${user.name} فعال شد`)
}

function disableTwoFactor(user) {
  // منطق غیرفعال کردن تایید دو مرحله‌ای
  user.status = 'disabled'
  user.activationDate = null
  alert(`تایید دو مرحله‌ای برای ${user.name} غیرفعال شد`)
}

// توابع صفحه‌بندی
function handlePageChange(page) {
  currentPage.value = page
}

function handleItemsPerPageChange(newItemsPerPage) {
  itemsPerPage.value = newItemsPerPage
  currentPage.value = 1 // بازگشت به صفحه اول
}
</script>

<style scoped>
.custom-scrollbar {
  scrollbar-width: thin;
  scrollbar-color: #cbd5e0 #f7fafc;
  direction: rtl;
}

.custom-scrollbar::-webkit-scrollbar {
  width: 8px;
  height: 8px;
}

.custom-scrollbar::-webkit-scrollbar-track {
  background: #f7fafc;
  border-radius: 4px;
}

.custom-scrollbar::-webkit-scrollbar-thumb {
  background: #cbd5e0;
  border-radius: 4px;
}

.custom-scrollbar::-webkit-scrollbar-thumb:hover {
  background: #a0aec0;
}

.custom-scrollbar::-webkit-scrollbar-corner {
  background: #f7fafc;
}

/* قرار دادن اسکرول در سمت چپ */
.custom-scrollbar {
  direction: ltr;
}

.custom-scrollbar > * {
  direction: rtl;
}
</style> 
