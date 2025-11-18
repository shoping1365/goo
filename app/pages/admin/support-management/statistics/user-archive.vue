<template>
  <div class="p-6" dir="rtl">
    <div class="mb-6 bg-white p-6 rounded-lg shadow-md border border-gray-200">
      <h1 class="text-2xl font-bold text-gray-900 mb-2">آرشیو کاربران</h1>
      <p class="text-gray-600">مدیریت و جستجو در آرشیو کاربران پشتیبانی</p>
    </div>

    <!-- آمار کلی -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
      <TemplateCard
        title="کل کاربران آرشیو شده"
        :value="totalArchivedUsers"
        variant="blue"
      />
      <TemplateCard
        title="کاربران فعال"
        :value="activeUsers"
        variant="green"
      />
      <TemplateCard
        title="کاربران غیرفعال"
        :value="inactiveUsers"
        variant="red"
      />
      <TemplateCard
        title="کاربران مسدود شده"
        :value="blockedUsers"
        variant="orange"
      />
    </div>

    <!-- فیلترها -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 mb-8">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">فیلترها و جستجو</h3>
      </div>
      <div class="p-6">
        <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نام یا ایمیل</label>
            <input type="text" placeholder="جستجو..." class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
            <select class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm">
              <option value="">همه</option>
              <option value="active">فعال</option>
              <option value="inactive">غیرفعال</option>
              <option value="blocked">مسدود شده</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ عضویت</label>
            <select class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm">
              <option value="">همه</option>
              <option value="today">امروز</option>
              <option value="week">هفته گذشته</option>
              <option value="month">ماه گذشته</option>
              <option value="year">سال گذشته</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع کاربر</label>
            <select class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm">
              <option value="">همه</option>
              <option value="regular">عادی</option>
              <option value="premium">پریمیوم</option>
              <option value="vip">VIP</option>
            </select>
          </div>
        </div>
      </div>
    </div>

    <!-- عملیات دسته‌ای -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 mb-8">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">عملیات دسته‌ای</h3>
      </div>
      <div class="p-6">
        <div class="flex flex-wrap gap-6">
          <button class="bg-green-500 text-white px-4 py-2 rounded-md hover:bg-green-600 transition-colors">
            فعال‌سازی انتخاب شده
          </button>
          <button class="bg-red-500 text-white px-4 py-2 rounded-md hover:bg-red-600 transition-colors">
            غیرفعال کردن انتخاب شده
          </button>
          <button class="bg-orange-500 text-white px-4 py-2 rounded-md hover:bg-orange-600 transition-colors">
            مسدود کردن انتخاب شده
          </button>
          <button class="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600 transition-colors">
            ارسال ایمیل دسته‌ای
          </button>
        </div>
      </div>
    </div>

    <!-- جدول آرشیو کاربران -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200 flex justify-between items-center">
        <h3 class="text-lg font-semibold text-gray-900">لیست کاربران آرشیو شده</h3>
        <div class="flex space-x-2 space-x-reverse">
          <button class="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600 transition-colors">
            دانلود گزارش
          </button>
          <button class="bg-gray-500 text-white px-4 py-2 rounded-md hover:bg-gray-600 transition-colors">
            پاک کردن آرشیو
          </button>
        </div>
      </div>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                <input type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کاربر</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">ایمیل</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ عضویت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">آخرین ورود</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="user in archivedUsers" :key="user.id">
              <td class="px-6 py-4 whitespace-nowrap">
                <input type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="">
                    <div class="h-10 w-10 rounded-full bg-gray-300 flex items-center justify-center">
                      <span class="text-sm font-medium text-gray-700">{{ user.name.charAt(0) }}</span>
                    </div>
                  </div>
                  <div class="mr-4">
                    <div class="text-sm font-medium text-gray-900">{{ user.name }}</div>
                    <div class="text-sm text-gray-500">{{ user.phone }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ user.email }}</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getUserTypeClass(user.type)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ getUserTypeText(user.type) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ user.joinDate }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ user.lastLogin }}</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusClass(user.status)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ getStatusText(user.status) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <div class="flex space-x-2 space-x-reverse">
                  <button class="text-blue-600 hover:text-blue-900 transition-colors" title="مشاهده جزئیات">
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                    </svg>
                  </button>
                  <button class="text-green-600 hover:text-green-900 transition-colors" title="فعال‌سازی">
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                    </svg>
                  </button>
                  <button class="text-orange-600 hover:text-orange-900 transition-colors" title="مسدود کردن">
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728L5.636 5.636m12.728 12.728L18.364 5.636M5.636 18.364l12.728-12.728"></path>
                    </svg>
                  </button>
                  <button class="text-red-600 hover:text-red-900 transition-colors" title="حذف">
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- صفحه‌بندی -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 mt-8">
      <div class="px-6 py-4 flex items-center justify-between">
        <div class="text-sm text-gray-700">
          نمایش <span class="font-medium">1</span> تا <span class="font-medium">10</span> از <span class="font-medium">156</span> نتیجه
        </div>
        <div class="flex space-x-2 space-x-reverse">
          <button class="px-3 py-2 text-sm font-medium text-gray-500 bg-white border border-gray-300 rounded-md hover:bg-gray-50">
            قبلی
          </button>
          <button class="px-3 py-2 text-sm font-medium text-white bg-blue-600 border border-blue-600 rounded-md">
            1
          </button>
          <button class="px-3 py-2 text-sm font-medium text-gray-500 bg-white border border-gray-300 rounded-md hover:bg-gray-50">
            2
          </button>
          <button class="px-3 py-2 text-sm font-medium text-gray-500 bg-white border border-gray-300 rounded-md hover:bg-gray-50">
            3
          </button>
          <button class="px-3 py-2 text-sm font-medium text-gray-500 bg-white border border-gray-300 rounded-md hover:bg-gray-50">
            بعدی
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import TemplateCard from '@/components/common/TemplateCard.vue'

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// داده‌های نمونه
const totalArchivedUsers = ref(156)
const activeUsers = ref(89)
const inactiveUsers = ref(45)
const blockedUsers = ref(22)

const archivedUsers = ref([
  {
    id: 1,
    name: 'احمد محمدی',
    email: 'ahmad@example.com',
    phone: '09123456789',
    type: 'regular',
    joinDate: '1401/03/15',
    lastLogin: '1402/10/15',
    status: 'active'
  },
  {
    id: 2,
    name: 'فاطمه کریمی',
    email: 'fateme@example.com',
    phone: '09387654321',
    type: 'premium',
    joinDate: '1401/05/20',
    lastLogin: '1402/10/14',
    status: 'active'
  },
  {
    id: 3,
    name: 'محمد رضایی',
    email: 'mohammad@example.com',
    phone: '09111111111',
    type: 'vip',
    joinDate: '1400/12/10',
    lastLogin: '1402/10/13',
    status: 'inactive'
  },
  {
    id: 4,
    name: 'سارا احمدی',
    email: 'sara@example.com',
    phone: '09222222222',
    type: 'regular',
    joinDate: '1401/08/05',
    lastLogin: '1402/10/12',
    status: 'blocked'
  },
  {
    id: 5,
    name: 'علی نوری',
    email: 'ali@example.com',
    phone: '09333333333',
    type: 'premium',
    joinDate: '1401/02/28',
    lastLogin: '1402/10/11',
    status: 'active'
  },
  {
    id: 6,
    name: 'زهرا محمدی',
    email: 'zahra@example.com',
    phone: '09444444444',
    type: 'regular',
    joinDate: '1401/07/12',
    lastLogin: '1402/10/10',
    status: 'inactive'
  },
  {
    id: 7,
    name: 'حسن کریمی',
    email: 'hasan@example.com',
    phone: '09555555555',
    type: 'vip',
    joinDate: '1400/11/03',
    lastLogin: '1402/10/09',
    status: 'active'
  },
  {
    id: 8,
    name: 'مریم رضایی',
    email: 'maryam@example.com',
    phone: '09666666666',
    type: 'regular',
    joinDate: '1401/09/18',
    lastLogin: '1402/10/08',
    status: 'blocked'
  },
  {
    id: 9,
    name: 'رضا احمدی',
    email: 'reza@example.com',
    phone: '09777777777',
    type: 'premium',
    joinDate: '1401/04/25',
    lastLogin: '1402/10/07',
    status: 'active'
  },
  {
    id: 10,
    name: 'نرگس نوری',
    email: 'narges@example.com',
    phone: '09888888888',
    type: 'regular',
    joinDate: '1401/06/30',
    lastLogin: '1402/10/06',
    status: 'inactive'
  }
])

function getUserTypeClass(type) {
  switch (type) {
    case 'regular':
      return 'bg-gray-100 text-gray-800'
    case 'premium':
      return 'bg-blue-100 text-blue-800'
    case 'vip':
      return 'bg-purple-100 text-purple-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

function getUserTypeText(type) {
  switch (type) {
    case 'regular':
      return 'عادی'
    case 'premium':
      return 'پریمیوم'
    case 'vip':
      return 'VIP'
    default:
      return 'نامشخص'
  }
}

function getStatusClass(status) {
  switch (status) {
    case 'active':
      return 'bg-green-100 text-green-800'
    case 'inactive':
      return 'bg-red-100 text-red-800'
    case 'blocked':
      return 'bg-orange-100 text-orange-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

function getStatusText(status) {
  switch (status) {
    case 'active':
      return 'فعال'
    case 'inactive':
      return 'غیرفعال'
    case 'blocked':
      return 'مسدود شده'
    default:
      return 'نامشخص'
  }
}
</script> 
