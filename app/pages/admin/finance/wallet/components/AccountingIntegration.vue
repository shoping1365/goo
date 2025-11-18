<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="bg-gradient-to-r from-teal-50 to-cyan-50 rounded-lg p-6">
      <h2 class="text-2xl font-bold text-gray-900 mb-2">🧮 اتصال به سیستم حسابداری</h2>
      <p class="text-gray-600">یکپارچه‌سازی با نرم‌افزارهای حسابداری و ثبت خودکار تراکنش‌ها</p>
    </div>

    <!-- وضعیت اتصال -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <!-- وضعیت کلی -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">وضعیت کلی</h3>
          <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">متصل</span>
        </div>
        <div class="text-3xl font-bold text-green-600 mb-2">{{ integrationStats.connectionStatus }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-green-500">{{ integrationStats.uptime }}%</span>
          <span class="mx-2">آپتایم</span>
        </div>
      </div>

      <!-- تراکنش‌های ثبت شده -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">تراکنش‌های ثبت شده</h3>
          <span class="text-xs bg-blue-100 text-blue-800 rounded-full px-3 py-1">امروز</span>
        </div>
        <div class="text-3xl font-bold text-blue-600 mb-2">{{ integrationStats.syncedTransactions }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-blue-500">{{ integrationStats.syncRate }}%</span>
          <span class="mx-2">نرخ همگام‌سازی</span>
        </div>
      </div>

      <!-- خطاهای اتصال -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">خطاهای اتصال</h3>
          <span class="text-xs bg-red-100 text-red-800 rounded-full px-3 py-1">هشدار</span>
        </div>
        <div class="text-3xl font-bold text-red-600 mb-2">{{ integrationStats.connectionErrors }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-red-500">{{ integrationStats.errorRate }}%</span>
          <span class="mx-2">نرخ خطا</span>
        </div>
      </div>

      <!-- آخرین همگام‌سازی -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">آخرین همگام‌سازی</h3>
          <span class="text-xs bg-yellow-100 text-yellow-800 rounded-full px-3 py-1">اخیر</span>
        </div>
        <div class="text-3xl font-bold text-yellow-600 mb-2">{{ integrationStats.lastSync }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-yellow-500">{{ integrationStats.syncDuration }}</span>
          <span class="mx-2">مدت زمان</span>
        </div>
      </div>
    </div>

    <!-- نرم‌افزارهای حسابداری -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-xl font-semibold text-gray-900">نرم‌افزارهای حسابداری</h3>
        <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
          افزودن نرم‌افزار جدید
        </button>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div v-for="software in accountingSoftware" :key="software.id" 
             class="border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow">
          <div class="flex items-center justify-between mb-3">
            <div class="flex items-center">
              <div class="w-10 h-10 bg-gray-100 rounded-lg flex items-center justify-center mr-3">
                <span class="text-lg">{{ software.icon }}</span>
              </div>
              <div>
                <h4 class="font-medium text-gray-900">{{ software.name }}</h4>
                <p class="text-sm text-gray-500">{{ software.version }}</p>
              </div>
            </div>
            <span :class="getConnectionStatusClass(software.status)" class="px-2 py-1 text-xs rounded-full">
              {{ software.status }}
            </span>
          </div>
          
          <div class="space-y-2 mb-4">
            <div class="flex justify-between text-sm">
              <span class="text-gray-600">نوع اتصال:</span>
              <span class="font-medium">{{ software.connectionType }}</span>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-600">آخرین همگام‌سازی:</span>
              <span class="font-medium">{{ software.lastSync }}</span>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-600">تراکنش‌های امروز:</span>
              <span class="font-medium">{{ software.todayTransactions }}</span>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-600">وضعیت API:</span>
              <span :class="software.apiStatus === 'فعال' ? 'text-green-600' : 'text-red-600'" class="font-medium">
                {{ software.apiStatus }}
              </span>
            </div>
          </div>
          
          <div class="flex space-x-2 space-x-reverse">
            <button class="flex-1 px-3 py-2 bg-blue-100 text-blue-700 text-sm rounded-lg hover:bg-blue-200">
              تنظیمات
            </button>
            <button @click="toggleGateway(software)" :class="software.status === 'فعال' ? 'bg-red-100 text-red-700 hover:bg-red-200' : 'bg-green-100 text-green-700 hover:bg-green-200'"
                    class="flex-1 px-3 py-2 text-sm rounded-lg">
              {{ software.status === 'فعال' ? 'غیرفعال' : 'فعال' }}
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات اتصال -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- تنظیمات API -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات API</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">آدرس سرور</label>
            <input type="text" :value="apiSettings.serverUrl" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">کلید API</label>
            <input type="password" :value="apiSettings.apiKey" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">رمز عبور</label>
            <input type="password" :value="apiSettings.password" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">بازه همگام‌سازی (دقیقه)</label>
            <input type="number" :value="apiSettings.syncInterval" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <div class="flex items-center">
            <input type="checkbox" :checked="apiSettings.autoSync" 
                   class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
            <label class="mr-2 text-sm text-gray-700">همگام‌سازی خودکار</label>
          </div>
          <button class="w-full px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
            ذخیره تنظیمات
          </button>
        </div>
      </div>

      <!-- تنظیمات کدهای حساب -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات کدهای حساب</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حساب کیف پول</label>
            <input type="text" :value="accountSettings.walletAccount" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حساب بانک</label>
            <input type="text" :value="accountSettings.bankAccount" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حساب کارمزد</label>
            <input type="text" :value="accountSettings.feeAccount" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مرکز هزینه</label>
            <input type="text" :value="accountSettings.costCenter" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">پروژه</label>
            <input type="text" :value="accountSettings.project" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <button class="w-full px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
            ذخیره تنظیمات
          </button>
        </div>
      </div>
    </div>

    <!-- جدول تراکنش‌های همگام شده -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-xl font-semibold text-gray-900">تراکنش‌های همگام شده</h3>
        <div class="flex space-x-2 space-x-reverse">
          <button class="px-4 py-2 bg-green-600 text-white text-sm rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500">
            همگام‌سازی دستی
          </button>
          <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
            خروجی اکسل
          </button>
        </div>
      </div>
      
      <div class="overflow-x-auto">
        <table class="w-full text-sm text-right">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-4 py-3 text-gray-700 font-medium">شماره تراکنش</th>
              <th class="px-4 py-3 text-gray-700 font-medium">نوع تراکنش</th>
              <th class="px-4 py-3 text-gray-700 font-medium">مبلغ</th>
              <th class="px-4 py-3 text-gray-700 font-medium">نرم‌افزار حسابداری</th>
              <th class="px-4 py-3 text-gray-700 font-medium">کد حساب</th>
              <th class="px-4 py-3 text-gray-700 font-medium">تاریخ همگام‌سازی</th>
              <th class="px-4 py-3 text-gray-700 font-medium">وضعیت</th>
              <th class="px-4 py-3 text-gray-700 font-medium">عملیات</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="transaction in syncedTransactions" :key="transaction.id" class="hover:bg-gray-50">
              <td class="px-4 py-3 text-gray-900 font-medium">{{ transaction.transactionId }}</td>
              <td class="px-4 py-3">
                <span :class="getTransactionTypeClass(transaction.type)" class="px-2 py-1 text-xs rounded-full">
                  {{ transaction.type }}
                </span>
              </td>
              <td class="px-4 py-3 font-medium" :class="transaction.amount > 0 ? 'text-green-600' : 'text-red-600'">
                {{ transaction.amount > 0 ? '+' : '' }}{{ formatCurrency(transaction.amount) }}
              </td>
              <td class="px-4 py-3 text-gray-600">{{ transaction.accountingSoftware }}</td>
              <td class="px-4 py-3 text-gray-600">{{ transaction.accountCode }}</td>
              <td class="px-4 py-3 text-gray-600">{{ transaction.syncDate }}</td>
              <td class="px-4 py-3">
                <span :class="getSyncStatusClass(transaction.status)" class="px-2 py-1 text-xs rounded-full">
                  {{ transaction.status }}
                </span>
              </td>
              <td class="px-4 py-3">
                <div class="flex space-x-2 space-x-reverse">
                  <button class="text-blue-600 hover:text-blue-800 text-sm">جزئیات</button>
                  <button v-if="transaction.status === 'ناموفق'" class="text-green-600 hover:text-green-800 text-sm">تلاش مجدد</button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- صفحه‌بندی -->
      <div class="flex items-center justify-between mt-6">
        <div class="text-sm text-gray-700">
          نمایش {{ paginationInfo.start }} تا {{ paginationInfo.end }} از {{ paginationInfo.total }} تراکنش
        </div>
        <div class="flex space-x-2 space-x-reverse">
          <button class="px-3 py-1 text-sm bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200">قبلی</button>
          <button class="px-3 py-1 text-sm bg-blue-600 text-white rounded-lg">1</button>
          <button class="px-3 py-1 text-sm bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200">2</button>
          <button class="px-3 py-1 text-sm bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200">3</button>
          <button class="px-3 py-1 text-sm bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200">بعدی</button>
        </div>
      </div>
    </div>

    <!-- آمار همگام‌سازی -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- نمودار وضعیت همگام‌سازی -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">وضعیت همگام‌سازی</h3>
        <div class="space-y-4">
          <div v-for="status in syncStatusStats" :key="status.name" class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-4 h-4 rounded-full mr-3" :style="{ backgroundColor: status.color }"></div>
              <span class="text-sm text-gray-700">{{ status.name }}</span>
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <span class="text-sm font-medium text-gray-900">{{ status.count }}</span>
              <span class="text-sm text-gray-500">({{ status.percentage }}%)</span>
            </div>
          </div>
        </div>
      </div>

      <!-- نمودار تراکنش‌ها بر اساس نرم‌افزار -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">تراکنش‌ها بر اساس نرم‌افزار</h3>
        <div class="space-y-4">
          <div v-for="software in softwareStats" :key="software.name" class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-4 h-4 rounded-full mr-3" :style="{ backgroundColor: software.color }"></div>
              <span class="text-sm text-gray-700">{{ software.name }}</span>
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <span class="text-sm font-medium text-gray-900">{{ software.count }}</span>
              <span class="text-sm text-gray-500">({{ software.percentage }}%)</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// آمار اتصال
const integrationStats = {
  connectionStatus: 'متصل',
  uptime: 99.8,
  syncedTransactions: 15420,
  syncRate: 98.5,
  connectionErrors: 23,
  errorRate: 0.15,
  lastSync: '2 دقیقه پیش',
  syncDuration: '45 ثانیه'
}

// تنظیمات API
const apiSettings = {
  serverUrl: 'https://api.accounting.com',
  apiKey: 'sk-1234567890abcdef',
  password: '********',
  syncInterval: 5,
  autoSync: true
}

// تنظیمات حساب
const accountSettings = {
  walletAccount: '1101',
  bankAccount: '1201',
  feeAccount: '5201',
  costCenter: 'CC001',
  project: 'PRJ001'
}

// فهرست درگاه‌ها از بک‌اند
import { ref, watchEffect } from 'vue'
declare const useFetch: <T>(url: string, options?: unknown) => Promise<{ data: { value: T }; refresh: () => Promise<void> }>
const accountingSoftware = ref<any[]>([])
const { data: gateways, refresh } = await useFetch('/api/payment-gateways', {
  credentials: 'include',
  key: 'payment-gateways-list',
  defaultCache: true,
  swr: true,
})
watchEffect(()=>{
  const res:any = gateways.value
  const list = Array.isArray(res?.data) ? res.data : []
  accountingSoftware.value = list.map((g:any)=>({
    id: g.id,
    name: g.name || g.english_name || g.type,
    version: g.type?.toUpperCase?.() || '',
    icon: '🏦',
    status: mapGatewayStatusFa(g.status),
    connectionType: 'API',
    lastSync: '-',
    todayTransactions: g.today_transactions || 0,
    apiStatus: g.status === 'active' ? 'فعال' : 'غیرفعال',
  }))
})

function mapGatewayStatusFa(st: string){
  if (st === 'active') return 'فعال'
  if (st === 'maintenance') return 'در حال تعمیر'
  return 'غیرفعال'
}

async function toggleGateway(item: any){
  const newStatus = item.status === 'فعال' ? 'inactive' : 'active'
  await $fetch(`/api/payment-gateways/${item.id}/status`, { method: 'PUT', body: { status: newStatus }, credentials: 'include' })
  await refresh()
}

// تراکنش‌های همگام شده
const syncedTransactions = [
  {
    id: 1,
    transactionId: 'TXN-2024-001',
    type: 'شارژ',
    amount: 5000000,
    accountingSoftware: 'هلو',
    accountCode: '1101',
    syncDate: '1402/10/15 - 14:30',
    status: 'موفق'
  },
  {
    id: 2,
    transactionId: 'TXN-2024-002',
    type: 'برداشت',
    amount: -2500000,
    accountingSoftware: 'پارسیان',
    accountCode: '1201',
    syncDate: '1402/10/15 - 13:45',
    status: 'موفق'
  },
  {
    id: 3,
    transactionId: 'TXN-2024-003',
    type: 'کارمزد',
    amount: -50000,
    accountingSoftware: 'هلو',
    accountCode: '5201',
    syncDate: '1402/10/15 - 12:20',
    status: 'ناموفق'
  },
  {
    id: 4,
    transactionId: 'TXN-2024-004',
    type: 'شارژ',
    amount: 15000000,
    accountingSoftware: 'قیاس',
    accountCode: '1101',
    syncDate: '1402/10/15 - 11:15',
    status: 'موفق'
  },
  {
    id: 5,
    transactionId: 'TXN-2024-005',
    type: 'انتقال',
    amount: -8000000,
    accountingSoftware: 'پارسیان',
    accountCode: '1101',
    syncDate: '1402/10/15 - 10:30',
    status: 'موفق'
  }
]

// اطلاعات صفحه‌بندی
const paginationInfo = {
  start: 1,
  end: 10,
  total: 15420
}

// آمار وضعیت همگام‌سازی
const syncStatusStats = [
  { name: 'موفق', count: 15180, percentage: 98.4, color: '#10B981' },
  { name: 'ناموفق', count: 240, percentage: 1.6, color: '#EF4444' }
]

// آمار نرم‌افزارها
const softwareStats = [
  { name: 'هلو', count: 6500, percentage: 42, color: '#10B981' },
  { name: 'پارسیان', count: 5200, percentage: 34, color: '#3B82F6' },
  { name: 'قیاس', count: 2800, percentage: 18, color: '#F59E0B' },
  { name: 'سپیدار', count: 920, percentage: 6, color: '#8B5CF6' }
]

// تابع فرمت کردن ارز
const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

// تابع کلاس وضعیت اتصال
const getConnectionStatusClass = (status: string) => {
  switch (status) {
    case 'متصل':
      return 'bg-green-100 text-green-800'
    case 'قطع اتصال':
      return 'bg-red-100 text-red-800'
    case 'در حال اتصال':
      return 'bg-yellow-100 text-yellow-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

// تابع کلاس نوع تراکنش
const getTransactionTypeClass = (type: string) => {
  switch (type) {
    case 'شارژ':
      return 'bg-green-100 text-green-800'
    case 'برداشت':
      return 'bg-red-100 text-red-800'
    case 'کارمزد':
      return 'bg-orange-100 text-orange-800'
    case 'انتقال':
      return 'bg-blue-100 text-blue-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

// تابع کلاس وضعیت همگام‌سازی
const getSyncStatusClass = (status: string) => {
  switch (status) {
    case 'موفق':
      return 'bg-green-100 text-green-800'
    case 'ناموفق':
      return 'bg-red-100 text-red-800'
    case 'در حال پردازش':
      return 'bg-yellow-100 text-yellow-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}
</script>

<!--
  مستندسازی:
  این کامپوننت شامل بخش‌های زیر است:
  1. وضعیت اتصال به سیستم‌های حسابداری
  2. مدیریت نرم‌افزارهای حسابداری
  3. تنظیمات API و کدهای حساب
  4. جدول تراکنش‌های همگام شده
  5. نمودارهای تحلیلی همگام‌سازی
  
  تمام توضیحات به فارسی و با طراحی ریسپانسیو ارائه شده است.
--> 
