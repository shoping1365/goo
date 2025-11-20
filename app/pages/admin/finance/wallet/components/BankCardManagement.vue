<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="bg-gradient-to-r from-orange-50 to-red-50 rounded-lg p-6">
      <h2 class="text-2xl font-bold text-gray-900 mb-2">🏦 مدیریت کارت‌های بانکی</h2>
      <p class="text-gray-600">مدیریت کارت‌های ثبت شده و تأیید آنها</p>
    </div>

    <!-- آمار کارت‌ها -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <!-- کل کارت‌ها -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">کل کارت‌ها</h3>
          <span class="text-xs bg-blue-100 text-blue-800 rounded-full px-3 py-1">ثبت شده</span>
        </div>
        <div class="text-3xl font-bold text-blue-600 mb-2">{{ cardStats.totalCards }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-green-500">+{{ cardStats.dailyGrowth }}%</span>
          <span class="mx-2">از دیروز</span>
        </div>
      </div>

      <!-- کارت‌های تأیید شده -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">تأیید شده</h3>
          <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">فعال</span>
        </div>
        <div class="text-3xl font-bold text-green-600 mb-2">{{ cardStats.verifiedCards }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-green-500">{{ cardStats.verificationRate }}%</span>
          <span class="mx-2">نرخ تأیید</span>
        </div>
      </div>

      <!-- کارت‌های در انتظار -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">در انتظار تأیید</h3>
          <span class="text-xs bg-yellow-100 text-yellow-800 rounded-full px-3 py-1">انتظار</span>
        </div>
        <div class="text-3xl font-bold text-yellow-600 mb-2">{{ cardStats.pendingCards }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-yellow-500">{{ cardStats.pendingPercentage }}%</span>
          <span class="mx-2">از کل کارت‌ها</span>
        </div>
      </div>

      <!-- کارت‌های مسدود -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">مسدود شده</h3>
          <span class="text-xs bg-red-100 text-red-800 rounded-full px-3 py-1">مسدود</span>
        </div>
        <div class="text-3xl font-bold text-red-600 mb-2">{{ cardStats.blockedCards }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-red-500">{{ cardStats.blockedPercentage }}%</span>
          <span class="mx-2">از کل کارت‌ها</span>
        </div>
      </div>
    </div>

    <!-- فیلترها -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-6">
        <!-- وضعیت -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
          <select v-model="statusFilter" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500" @change="refresh()">
            <option value="">همه</option>
            <option value="active">تأیید شده</option>
            <option value="pending">در انتظار تأیید</option>
            <option value="blocked">مسدود شده</option>
            <option value="rejected">رد شده</option>
          </select>
        </div>

        <!-- بانک -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">بانک</label>
          <select v-model="bankFilter" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500" @change="refresh()">
            <option value="">همه</option>
            <option value="melli">ملی</option>
            <option value="mellat">ملت</option>
            <option value="parsian">پارسیان</option>
            <option value="sepah">سپه</option>
            <option value="tejarat">تجارت</option>
          </select>
        </div>

        <!-- نوع کارت -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نوع کارت</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option value="">همه</option>
            <option value="debit">کارت عابربانک</option>
            <option value="credit">کارت اعتباری</option>
            <option value="prepaid">کارت پیش‌پرداخت</option>
          </select>
        </div>

        <!-- جستجو -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">جستجو</label>
          <input
type="text" placeholder="نام کاربر یا شماره کارت" 
                 class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
      </div>

      <div class="flex justify-between items-center">
        <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
          اعمال فیلتر
        </button>
        <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
          پاک کردن فیلترها
        </button>
      </div>
    </div>

    <!-- جدول کارت‌های بانکی -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-xl font-semibold text-gray-900">کارت‌های بانکی</h3>
        <div class="flex space-x-2 space-x-reverse">
          <button class="px-4 py-2 bg-green-600 text-white text-sm rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500">
            تأیید همه
          </button>
          <button class="px-4 py-2 bg-red-600 text-white text-sm rounded-lg hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500">
            مسدود کردن همه
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
              <th class="px-4 py-3 text-gray-700 font-medium">
                <input type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
              </th>
              <th class="px-4 py-3 text-gray-700 font-medium">کاربر</th>
              <th class="px-4 py-3 text-gray-700 font-medium">شماره کارت</th>
              <th class="px-4 py-3 text-gray-700 font-medium">بانک</th>
              <th class="px-4 py-3 text-gray-700 font-medium">نوع کارت</th>
              <th class="px-4 py-3 text-gray-700 font-medium">تاریخ ثبت</th>
              <th class="px-4 py-3 text-gray-700 font-medium">وضعیت</th>
              <th class="px-4 py-3 text-gray-700 font-medium">عملیات</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="card in bankCards" :key="card.id" class="hover:bg-gray-50">
              <td class="px-4 py-3">
                <input type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
              </td>
              <td class="px-4 py-3">
                <div class="flex items-center">
                  <div class="w-8 h-8 bg-gray-200 rounded-full flex items-center justify-center text-xs font-medium text-gray-600">
                    {{ card.userInitials }}
                  </div>
                  <div class="mr-3">
                    <div class="text-sm font-medium text-gray-900">{{ card.userName }}</div>
                    <div class="text-xs text-gray-500">{{ card.userEmail }}</div>
                  </div>
                </div>
              </td>
              <td class="px-4 py-3 text-gray-900 font-medium">{{ card.cardNumber }}</td>
              <td class="px-4 py-3">
                <div class="flex items-center">
                  <div class="w-6 h-6 bg-gray-100 rounded flex items-center justify-center mr-2">
                    <span class="text-xs">{{ card.bankIcon }}</span>
                  </div>
                  <span class="text-sm text-gray-700">{{ card.bankName }}</span>
                </div>
              </td>
              <td class="px-4 py-3">
                <span :class="getCardTypeClass(card.cardType)" class="px-2 py-1 text-xs rounded-full">
                  {{ card.cardType }}
                </span>
              </td>
              <td class="px-4 py-3 text-gray-600">{{ card.registrationDate }}</td>
              <td class="px-4 py-3">
                <span :class="getStatusClass(card.status)" class="px-2 py-1 text-xs rounded-full">
                  {{ card.status }}
                </span>
              </td>
              <td class="px-4 py-3">
                <div class="flex space-x-2 space-x-reverse">
                  <button class="text-blue-600 hover:text-blue-800 text-sm">جزئیات</button>
                  <button v-if="card.status === 'در انتظار تأیید'" class="text-green-600 hover:text-green-800 text-sm" @click="verifyCard(card.id)">تأیید</button>
                  <button v-if="card.status === 'در انتظار تأیید'" class="text-red-600 hover:text-red-800 text-sm" @click="rejectCard(card.id)">رد</button>
                  <button v-if="card.status === 'تأیید شده'" class="text-red-600 hover:text-red-800 text-sm" @click="blockCard(card.id)">مسدود</button>
                  <button v-if="card.status === 'مسدود شده'" class="text-green-600 hover:text-green-800 text-sm" @click="unblockCard(card.id)">باز کردن</button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- صفحه‌بندی -->
      <div class="flex items-center justify-between mt-6">
        <div class="text-sm text-gray-700">
          نمایش {{ paginationInfo.start }} تا {{ paginationInfo.end }} از {{ paginationInfo.total }} کارت
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

    <!-- آمار بانک‌ها -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- نمودار کارت‌ها بر اساس بانک -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">کارت‌ها بر اساس بانک</h3>
        <div class="space-y-4">
          <div v-for="bank in bankStats" :key="bank.name" class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-4 h-4 rounded-full mr-3" :style="{ backgroundColor: bank.color }"></div>
              <span class="text-sm text-gray-700">{{ bank.name }}</span>
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <span class="text-sm font-medium text-gray-900">{{ bank.count }}</span>
              <span class="text-sm text-gray-500">({{ bank.percentage }}%)</span>
            </div>
          </div>
        </div>
      </div>

      <!-- نمودار کارت‌ها بر اساس وضعیت -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">کارت‌ها بر اساس وضعیت</h3>
        <div class="space-y-4">
          <div v-for="status in statusStats" :key="status.name" class="flex items-center justify-between">
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
    </div>

    <!-- تنظیمات امنیتی -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات امنیتی کارت‌ها</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر کارت برای هر کاربر</label>
            <input
type="number" :value="securitySettings.maxCardsPerUser" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداقل مبلغ تراکنش</label>
            <input
type="number" :value="securitySettings.minTransactionAmount" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر مبلغ تراکنش</label>
            <input
type="number" :value="securitySettings.maxTransactionAmount" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">زمان تأیید خودکار (ساعت)</label>
            <input
type="number" :value="securitySettings.autoVerificationTime" 
                   class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
        </div>
        
        <div class="space-y-4">
          <div class="flex items-center">
            <input
type="checkbox" :checked="securitySettings.requireCardVerification" 
                   class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
            <label class="mr-2 text-sm text-gray-700">نیاز به تأیید کارت</label>
          </div>
          <div class="flex items-center">
            <input
type="checkbox" :checked="securitySettings.enableCardBlocking" 
                   class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
            <label class="mr-2 text-sm text-gray-700">امکان مسدود کردن کارت</label>
          </div>
          <button class="w-full px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
            ذخیره تنظیمات
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref, watchEffect } from 'vue';
declare const useFetch: <T>(url: string, options?: unknown) => Promise<{ data: { value: T }; refresh: () => Promise<void> }>
// آمار کارت‌ها از API ادمین بارگذاری می‌شود
const cardStats = reactive({
  totalCards: 0,
  dailyGrowth: 0,
  verifiedCards: 0,
  verificationRate: 0,
  pendingCards: 0,
  pendingPercentage: 0,
  blockedCards: 0,
  blockedPercentage: 0
})

// تنظیمات امنیتی
const securitySettings = {
  maxCardsPerUser: 5,
  minTransactionAmount: 10000,
  maxTransactionAmount: 50000000,
  autoVerificationTime: 24,
  requireCardVerification: true,
  enableCardBlocking: true
}

// کارت‌های بانکی از API ادمین
const bankCards = ref<any[]>([])
const page = ref(1)
const pageSize = ref(20)
const statusFilter = ref('')
const bankFilter = ref('')
const { data, refresh } = await useFetch('/api/admin/bank-accounts', {
  method: 'GET',
  query: { page, pageSize, status: statusFilter, bank: bankFilter },
  credentials: 'include',
  key: () => `admin-bank-accounts-${page.value}-${pageSize.value}-${statusFilter.value}-${bankFilter.value}`,
  defaultCache: true,
})
const paginationInfo = reactive({ start: 1, end: 0, total: 0 })
watchEffect(() => {
  const res: any = data.value
  if (!res) return
  const items = Array.isArray(res.items) ? res.items : []
  bankCards.value = items.map((x: any) => ({
    id: x.id,
    userName: x.user_name || x.user_id,
    userEmail: x.user_email || '-',
    userInitials: (x.user_name || '-').slice(0,2),
    cardNumber: x.iban ? `${x.iban.slice(0,6)}****${x.iban.slice(-4)}` : (x.card_number_last4 ? `****-${x.card_number_last4}` : '-'),
    bankName: x.bank_name || '-',
    bankIcon: '🏦',
    cardType: 'کارت عابربانک',
    registrationDate: x.created_at || '-',
    status: x.status === 'active' ? 'تأیید شده' : (x.status === 'blocked' ? 'مسدود شده' : (x.verified_at ? 'تأیید شده' : 'در انتظار تأیید'))
  }))
  paginationInfo.total = Number(res.total || 0)
  paginationInfo.start = (page.value-1) * pageSize.value + 1
  paginationInfo.end = Math.min(page.value * pageSize.value, paginationInfo.total)
  // آمار ساده سمت‌کلاینت
  cardStats.totalCards = paginationInfo.total
  const verified = bankCards.value.filter((c:any)=>c.status==='تأیید شده').length
  const blocked = bankCards.value.filter((c:any)=>c.status==='مسدود شده').length
  const pending = Math.max(0, cardStats.totalCards - verified - blocked)
  cardStats.verifiedCards = verified
  cardStats.blockedCards = blocked
  cardStats.pendingCards = pending
  const tot = Math.max(1, cardStats.totalCards)
  cardStats.verificationRate = Math.round((verified/tot)*1000)/10
  cardStats.blockedPercentage = Math.round((blocked/tot)*1000)/10
  cardStats.pendingPercentage = Math.round((pending/tot)*1000)/10
})

async function verifyCard(id: number) {
  await $fetch(`/api/admin/bank-accounts/${id}/verify`, { method: 'POST', credentials: 'include' })
  await refresh()
}
async function blockCard(id: number) {
  await $fetch(`/api/admin/bank-accounts/${id}/block`, { method: 'POST', credentials: 'include' })
  await refresh()
}
async function unblockCard(id: number) {
  await $fetch(`/api/admin/bank-accounts/${id}/unblock`, { method: 'POST', credentials: 'include' })
  await refresh()
}
async function rejectCard(id: number) {
  await $fetch(`/api/admin/bank-accounts/${id}/reject`, { method: 'POST', credentials: 'include' })
  await refresh()
}

// اطلاعات صفحه‌بندی مقداردهی در بالا انجام می‌شود

// آمار بانک‌ها
const bankStats = [
  { name: 'ملی', count: 4250, percentage: 28, color: '#10B981' },
  { name: 'ملت', count: 3850, percentage: 25, color: '#3B82F6' },
  { name: 'پارسیان', count: 2150, percentage: 14, color: '#F59E0B' },
  { name: 'سپه', count: 1850, percentage: 12, color: '#EF4444' },
  { name: 'تجارت', count: 1650, percentage: 11, color: '#8B5CF6' },
  { name: 'سایر', count: 1670, percentage: 10, color: '#6B7280' }
]

// آمار وضعیت‌ها
const statusStats = [
  { name: 'تأیید شده', count: 14250, percentage: 92.4, color: '#10B981' },
  { name: 'در انتظار تأیید', count: 890, percentage: 5.8, color: '#F59E0B' },
  { name: 'مسدود شده', count: 280, percentage: 1.8, color: '#EF4444' }
]

// تابع کلاس نوع کارت
const getCardTypeClass = (cardType: string) => {
  switch (cardType) {
    case 'کارت عابربانک':
      return 'bg-blue-100 text-blue-800'
    case 'کارت اعتباری':
      return 'bg-green-100 text-green-800'
    case 'کارت پیش‌پرداخت':
      return 'bg-purple-100 text-purple-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

// تابع کلاس وضعیت
const getStatusClass = (status: string) => {
  switch (status) {
    case 'تأیید شده':
      return 'bg-green-100 text-green-800'
    case 'در انتظار تأیید':
      return 'bg-yellow-100 text-yellow-800'
    case 'مسدود شده':
      return 'bg-red-100 text-red-800'
    case 'رد شده':
      return 'bg-gray-100 text-gray-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}
</script>

<!--
  مستندسازی:
  این کامپوننت شامل بخش‌های زیر است:
  1. آمار کلی کارت‌های بانکی (کل، تأیید شده، در انتظار، مسدود)
  2. فیلترهای پیشرفته برای جستجو
  3. جدول کارت‌های بانکی با امکان تأیید/رد/مسدود کردن
  4. نمودارهای تحلیلی بر اساس بانک و وضعیت
  5. تنظیمات امنیتی کارت‌ها
  
  تمام توضیحات به فارسی و با طراحی ریسپانسیو ارائه شده است.
--> 
