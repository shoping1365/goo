<template>
  <div class="w-full px-4 py-4" dir="rtl">
    <!-- Page Header -->
    <div class="bg-white shadow-sm border-b border-gray-200 mb-4">
      <div class="w-full px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center py-3">
          <div>
            <h1 class="text-lg font-bold text-gray-900">مسدودی</h1>
            <p class="text-xs text-gray-600 mt-1">مدیریت IP ها و کاربران مسدود شده</p>
          </div>
          <div class="flex">
            <!-- دکمه خروجی اکسل -->
            <button
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-green-400 to-green-600 hover:from-green-500 hover:to-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105 ml-8"
              @click="exportToExcel"
            >
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
              </svg>
              خروجی اکسل
            </button>
            
            <!-- دکمه بلاک کردن لیست IP -->
            <button
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-orange-400 to-orange-600 hover:from-orange-500 hover:to-orange-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-orange-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105 ml-8"
              @click="showBulkBlockModal = true"
            >
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
              بلاک کردن لیست IP
            </button>
            
            <button
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-red-400 to-red-600 hover:from-red-500 hover:to-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105"
              @click="showAddBlockModal = true"
            >
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
              </svg>
              افزودن مسدودی جدید
            </button>
          </div>
        </div>
      </div>
    </div>

    <div class="mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-6">
        <TemplateCard
          title="IP های مسدود شده"
          :value="blockedIPs"
          variant="red"
        />
        <TemplateCard
          title="کاربران مسدود شده"
          :value="blockedUsers"
          variant="orange"
        />
        <TemplateCard
          title="مسدودی‌های موقت"
          :value="temporaryBlocks"
          variant="purple"
        />
        <TemplateCard
          title="مسدودی‌های دائمی"
          :value="permanentBlocks"
          variant="gray"
        />
      </div>
    </div>



    <!-- جدول IP های مسدود شده -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">IP های مسدود شده</h3>
      </div>
      <div class="overflow-y-auto min-h-[300px] md:min-h-[400px] max-h-[70vh]" dir="ltr"> <!-- اسکرول عمودی واکنش‌گرا با جهت LTR -->
        <table class="min-w-full divide-y divide-gray-200" dir="rtl">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">IP آدرس</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">دلیل</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع مسدودی</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ مسدودی</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">انقضا</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="ip in blockedIPsList.slice(0, blockedIPsItemsPerPage)" :key="ip.id">
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{{ ip.address }}</td>
                <td class="px-6 py-4 text-sm text-gray-900">{{ ip.reason }}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  <div class="flex items-center">
                    <span v-if="ip.blockedBy === 'system'" class="i-heroicons-cpu-chip text-blue-500 ml-1"></span>
                    <span v-else class="i-heroicons-user text-green-500 ml-1"></span>
                    {{ ip.blockedBy === 'system' ? 'سیستم' : 'مدیر' }}
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ ip.blockedAt }}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ ip.expiresAt }}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                  <button v-if="hasPermission('security.delete')" class="text-green-600 hover:text-green-900" @click="removeBlock(ip.id)">حذف مسدودی</button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
      <Pagination
        :current-page="blockedIPsCurrentPage"
        :total-pages="blockedIPsTotalPages"
        :total="blockedIPsList.length"
        :items-per-page="blockedIPsItemsPerPage"
        class="mb-8"
        @items-per-page-changed="blockedIPsItemsPerPage = $event"
      />
    </div>

    <!-- جدول کاربران مسدود شده -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">کاربران مسدود شده</h3>
      </div>
      <div class="overflow-y-auto min-h-[300px] md:min-h-[400px] max-h-[70vh]" dir="ltr"> <!-- اسکرول عمودی واکنش‌گرا با جهت LTR -->
        <table class="min-w-full divide-y divide-gray-200" dir="rtl">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کاربر</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">دلیل</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع مسدودی</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ مسدودی</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">انقضا</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="user in blockedUsersList.slice(0, blockedUsersItemsPerPage)" :key="user.id">
                <td class="px-6 py-4 whitespace-nowrap max-w-xs">
                  <div class="flex items-center">
                    <div class="flex-shrink-0 h-8 w-8">
                      <div class="h-8 w-8 rounded-full bg-gray-300 flex items-center justify-center">
                        <span class="text-sm font-medium text-gray-700">{{ user.name.charAt(0) }}</span>
                      </div>
                    </div>
                    <div class="mr-3 min-w-0 flex-1">
                      <div class="text-sm font-medium text-gray-900 truncate">{{ user.name }}</div>
                      <div class="text-sm text-gray-500 truncate">{{ user.email }}</div>
                    </div>
                  </div>
                </td>
                <td class="px-6 py-4 text-sm text-gray-900">{{ user.reason }}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  <div class="flex items-center">
                    <span v-if="user.blockedBy === 'system'" class="i-heroicons-cpu-chip text-blue-500 ml-1"></span>
                    <span v-else class="i-heroicons-user text-green-500 ml-1"></span>
                    {{ user.blockedBy === 'system' ? 'سیستم' : 'مدیر' }}
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ user.blockedAt }}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ user.expiresAt }}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                  <button v-if="hasPermission('security.delete')" class="text-green-600 hover:text-green-900" @click="removeUserBlock(user.id)">حذف مسدودی</button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
      <Pagination
        :current-page="blockedUsersCurrentPage"
        :total-pages="blockedUsersTotalPages"
        :total="blockedUsersList.length"
        :items-per-page="blockedUsersItemsPerPage"
        class="mb-12"
        @items-per-page-changed="blockedUsersItemsPerPage = $event"
      />


  <!-- مودال بلاک کردن لیست IP -->
  <div v-if="showBulkBlockModal" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50" @click="showBulkBlockModal = false">
    <div class="relative top-20 mx-auto p-5 border w-11/12 md:w-3/4 lg:w-1/2 shadow-lg rounded-md bg-white" @click.stop>
      <div class="mt-3">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">بلاک کردن لیست IP</h3>
          <button class="text-gray-400 hover:text-gray-600" @click="showBulkBlockModal = false">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>
        
        <div class="mb-4">
          <label class="block text-sm font-medium text-gray-700 mb-2">لیست IP های غیرمجاز (هر IP در یک خط)</label>
          <textarea 
            v-model="bulkIPs" 
            rows="8" 
            placeholder="مثال:
192.168.1.100
10.0.0.50
172.16.0.25
203.0.113.10" 
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 font-mono text-sm"
          ></textarea>
        </div>
        
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مدت زمان مسدودی</label>
            <select v-model="bulkBlockDuration" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
              <option value="1h">1 ساعت</option>
              <option value="24h">24 ساعت</option>
              <option value="7d">7 روز</option>
              <option value="30d">30 روز</option>
              <option value="permanent">دائمی</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">دلیل مسدودی</label>
            <input v-model="bulkBlockReason" type="text" placeholder="دلیل مسدودی" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
        </div>
        
        <div class="flex justify-end space-x-2 space-x-reverse">
          <button class="px-4 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400 focus:outline-none focus:ring-2 focus:ring-gray-500" @click="showBulkBlockModal = false">
            انصراف
          </button>
          <button class="px-4 py-2 bg-orange-600 text-white rounded-md hover:bg-orange-700 focus:outline-none focus:ring-2 focus:ring-orange-500" @click="bulkBlockIPs">
            بلاک کردن IP ها
          </button>
        </div>
      </div>
    </div>
  </div>

  <!-- مودال افزودن مسدودی جدید -->
  <div v-if="showAddBlockModal" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50" @click="showAddBlockModal = false">
    <div class="relative top-20 mx-auto p-5 border w-11/12 md:w-3/4 lg:w-1/2 shadow-lg rounded-md bg-white" @click.stop>
      <div class="mt-3">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">افزودن مسدودی جدید</h3>
          <button class="text-gray-400 hover:text-gray-600" @click="showAddBlockModal = false">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>
        
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع مسدودی</label>
            <select v-model="newBlock.type" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
              <option value="ip">IP آدرس</option>
              <option value="user">کاربر</option>
              <option value="range">محدوده IP</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مقدار</label>
            <input v-model="newBlock.value" type="text" placeholder="IP یا نام کاربری" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مدت زمان</label>
            <select v-model="newBlock.duration" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
              <option value="1h">1 ساعت</option>
              <option value="24h">24 ساعت</option>
              <option value="7d">7 روز</option>
              <option value="30d">30 روز</option>
              <option value="permanent">دائمی</option>
            </select>
          </div>
        </div>
        
        <div class="mb-4">
          <label class="block text-sm font-medium text-gray-700 mb-2">دلیل مسدودی</label>
          <textarea v-model="newBlock.reason" rows="3" placeholder="دلیل مسدودی را وارد کنید" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"></textarea>
        </div>
        
        <div class="flex justify-end space-x-2 space-x-reverse">
          <button class="px-4 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400 focus:outline-none focus:ring-2 focus:ring-gray-500" @click="showAddBlockModal = false">
            انصراف
          </button>
          <button class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500" @click="addBlock">
            افزودن مسدودی
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
declare const useAuth: () => { user: { id?: number; name?: string; email?: string } | null; hasPermission: (perm: string) => boolean }
</script>

<script setup lang="ts">
import TemplateCard from '@/components/common/TemplateCard.vue';
import { computed, ref } from 'vue';
import Pagination from '~/components/admin/common/Pagination.vue';

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// استفاده از کامپوزابل احراز هویت
// Auth disabled
const { hasPermission } = useAuth()

// داده‌های نمونه
const blockedIPs = ref(15)
const blockedUsers = ref(8)
const temporaryBlocks = ref(12)
const permanentBlocks = ref(5)

// مودال‌ها
const showAddBlockModal = ref(false)
const showBulkBlockModal = ref(false)

// متغیرهای بلاک کردن لیست IP
const bulkIPs = ref('')
const bulkBlockDuration = ref('24h')
const bulkBlockReason = ref('')

// فقط متغیرهای مربوط به تعداد آیتم در هر صفحه را نگه دار و متغیرهای صفحه فعلی و computedهای صفحه‌بندی را حذف کن
const blockedIPsItemsPerPage = ref(10)
const blockedUsersItemsPerPage = ref(10)
const blockedIPsCurrentPage = ref(1)
const blockedUsersCurrentPage = ref(1)

const blockedIPsTotalPages = computed(() => Math.ceil(blockedIPsList.value.length / blockedIPsItemsPerPage.value))
const blockedUsersTotalPages = computed(() => Math.ceil(blockedUsersList.value.length / blockedUsersItemsPerPage.value))

const newBlock = ref({
  type: 'ip',
  value: '',
  duration: '24h',
  reason: ''
})

const blockedIPsList = ref([
  {
    id: 1,
    address: '192.168.1.100',
    reason: 'تلاش برای ورود غیرمجاز',
    blockedAt: '1402/10/15 14:30',
    expiresAt: '1402/10/16 14:30',
    blockedBy: 'user'
  },
  {
    id: 2,
    address: '10.0.0.50',
    reason: 'حمله DDoS',
    blockedAt: '1402/10/14 09:15',
    expiresAt: 'دائمی',
    blockedBy: 'system'
  },
  {
    id: 3,
    address: '172.16.0.25',
    reason: 'اسکن پورت',
    blockedAt: '1402/10/13 22:45',
    expiresAt: '1402/10/20 22:45',
    blockedBy: 'user'
  },
  {
    id: 4,
    address: '203.0.113.10',
    reason: 'فعالیت مشکوک',
    blockedAt: '1402/10/12 16:20',
    expiresAt: 'دائمی',
    blockedBy: 'system'
  },
  {
    id: 5,
    address: '198.51.100.75',
    reason: 'تلاش برای دسترسی غیرمجاز',
    blockedAt: '1402/10/11 11:30',
    expiresAt: '1402/10/18 11:30',
    blockedBy: 'user'
  },
  {
    id: 6,
    address: '8.8.8.8',
    reason: 'حمله Brute Force',
    blockedAt: '1402/10/10 08:15',
    expiresAt: 'دائمی',
    blockedBy: 'system'
  },
  {
    id: 7,
    address: '1.1.1.1',
    reason: 'اسپم',
    blockedAt: '1402/10/09 19:45',
    expiresAt: '1402/10/16 19:45',
    blockedBy: 'user'
  },
  {
    id: 8,
    address: '208.67.222.222',
    reason: 'فعالیت غیرعادی',
    blockedAt: '1402/10/08 13:20',
    expiresAt: '1402/10/15 13:20',
    blockedBy: 'system'
  },
  {
    id: 9,
    address: '185.199.108.153',
    reason: 'اسکن امنیتی',
    blockedAt: '1402/10/07 20:10',
    expiresAt: '1402/10/14 20:10',
    blockedBy: 'system'
  },
  {
    id: 10,
    address: '140.82.112.4',
    reason: 'حمله XSS',
    blockedAt: '1402/10/06 16:45',
    expiresAt: 'دائمی',
    blockedBy: 'user'
  }
])

const blockedUsersList = ref([
  {
    id: 1,
    name: 'احمد محمدی',
    email: 'ahmad@example.com',
    reason: 'نقض قوانین سایت',
    blockedAt: '1402/10/13 16:45',
    expiresAt: '1402/11/13 16:45',
    blockedBy: 'user'
  },
  {
    id: 2,
    name: 'سارا احمدی',
    email: 'sara@example.com',
    reason: 'ارسال اسپم',
    blockedAt: '1402/10/12 11:20',
    expiresAt: 'دائمی',
    blockedBy: 'system'
  },
  {
    id: 3,
    name: 'محمد رضایی',
    email: 'mohammad@example.com',
    reason: 'فعالیت مشکوک',
    blockedAt: '1402/10/11 09:30',
    expiresAt: '1402/11/11 09:30',
    blockedBy: 'user'
  },
  {
    id: 4,
    name: 'فاطمه کریمی',
    email: 'fateme@example.com',
    reason: 'نقض قوانین',
    blockedAt: '1402/10/10 14:15',
    expiresAt: 'دائمی',
    blockedBy: 'system'
  },
  {
    id: 5,
    name: 'علی نوری',
    email: 'ali@example.com',
    reason: 'ارسال محتوای نامناسب',
    blockedAt: '1402/10/09 18:45',
    expiresAt: '1402/11/09 18:45',
    blockedBy: 'user'
  },
  {
    id: 6,
    name: 'زهرا صالحی',
    email: 'zahra@example.com',
    reason: 'فعالیت غیرمجاز',
    blockedAt: '1402/10/08 12:30',
    expiresAt: '1402/11/08 12:30',
    blockedBy: 'system'
  },
  {
    id: 7,
    name: 'حسین رضوی',
    email: 'hossein@example.com',
    reason: 'نقض حریم خصوصی',
    blockedAt: '1402/10/07 15:20',
    expiresAt: 'دائمی',
    blockedBy: 'user'
  },
  {
    id: 8,
    name: 'مریم جعفری',
    email: 'maryam@example.com',
    reason: 'ارسال پیام‌های توهین‌آمیز',
    blockedAt: '1402/10/06 10:45',
    expiresAt: '1402/11/06 10:45',
    blockedBy: 'system'
  },
  {
    id: 9,
    name: 'رضا کریمی',
    email: 'reza@example.com',
    reason: 'تلاش برای هک',
    blockedAt: '1402/10/05 08:15',
    expiresAt: 'دائمی',
    blockedBy: 'system'
  },
  {
    id: 10,
    name: 'نرگس احمدی',
    email: 'narges@example.com',
    reason: 'فعالیت مشکوک',
    blockedAt: '1402/10/04 19:30',
    expiresAt: '1402/11/04 19:30',
    blockedBy: 'user'
  }
])


function addBlock() {
  // منطق افزودن مسدودی
  alert('مسدودی جدید اضافه شد')
  newBlock.value = {
    type: 'ip',
    value: '',
    duration: '24h',
    reason: ''
  }
  showAddBlockModal.value = false
}

function removeBlock(id) {
  // منطق حذف مسدودی IP
  alert('مسدودی IP حذف شد')
}

function removeUserBlock(id) {
  // منطق حذف مسدودی کاربر
  alert('مسدودی کاربر حذف شد')
}

// تابع خروجی اکسل
function exportToExcel() {
  // ایجاد داده‌های اکسل
  const excelData = []
  
  // اضافه کردن هدر
  excelData.push(['IP آدرس', 'دلیل', 'نوع مسدودی', 'تاریخ مسدودی', 'انقضا'])
  
  // اضافه کردن داده‌های IP ها
  blockedIPsList.value.forEach(ip => {
    excelData.push([
      ip.address,
      ip.reason,
      ip.blockedBy === 'system' ? 'سیستم' : 'مدیر',
      ip.blockedAt,
      ip.expiresAt
    ])
  })
  
  // اضافه کردن خط خالی
  excelData.push([])
  excelData.push(['کاربر', 'ایمیل', 'دلیل', 'نوع مسدودی', 'تاریخ مسدودی', 'انقضا'])
  
  // اضافه کردن داده‌های کاربران
  blockedUsersList.value.forEach(user => {
    excelData.push([
      user.name,
      user.email,
      user.reason,
      user.blockedBy === 'system' ? 'سیستم' : 'مدیر',
      user.blockedAt,
      user.expiresAt
    ])
  })
  
  // تبدیل به CSV
  const csvContent = excelData.map(row => row.map(cell => `"${cell}"`).join(',')).join('\n')
  
  // ایجاد فایل و دانلود
  const blob = new Blob(['\ufeff' + csvContent], { type: 'text/csv;charset=utf-8;' })
  const link = document.createElement('a')
  const url = URL.createObjectURL(blob)
  link.setAttribute('href', url)
  link.setAttribute('download', `مسدودی‌ها_${new Date().toLocaleDateString('fa-IR')}.csv`)
  link.style.visibility = 'hidden'
  document.body.appendChild(link)
  link.click()
  document.body.removeChild(link)
}

// تابع بلاک کردن لیست IP
function bulkBlockIPs() {
  if (!bulkIPs.value.trim()) {
    alert('لطفاً لیست IP ها را وارد کنید')
    return
  }
  
  if (!bulkBlockReason.value.trim()) {
    alert('لطفاً دلیل مسدودی را وارد کنید')
    return
  }
  
  // تقسیم IP ها بر اساس خط جدید
  const ipList = bulkIPs.value.split('\n').map(ip => ip.trim()).filter(ip => ip)
  
  if (ipList.length === 0) {
    alert('هیچ IP معتبری یافت نشد')
    return
  }
  
  // اعتبارسنجی IP ها
  const validIPs = ipList.filter(ip => {
    const ipRegex = /^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$/
    return ipRegex.test(ip)
  })
  
  if (validIPs.length === 0) {
    alert('هیچ IP معتبری یافت نشد')
    return
  }
  
  // اضافه کردن IP ها به لیست مسدودی‌ها
  const newId = Math.max(...blockedIPsList.value.map(ip => ip.id)) + 1
  const currentDate = new Date().toLocaleString('fa-IR')
  
  validIPs.forEach((ip, index) => {
    const newBlock = {
      id: newId + index,
      address: ip,
      reason: bulkBlockReason.value,
      blockedAt: currentDate,
      expiresAt: bulkBlockDuration.value === 'permanent' ? 'دائمی' : 'موقت',
      blockedBy: 'user'
    }
    blockedIPsList.value.unshift(newBlock)
  })
  
  // به‌روزرسانی آمار
  blockedIPs.value += validIPs.length
  
  // پاک کردن فرم
  bulkIPs.value = ''
  bulkBlockReason.value = ''
  bulkBlockDuration.value = '24h'
  showBulkBlockModal.value = false
  
  alert(`${validIPs.length} IP با موفقیت مسدود شد`)
}
</script>

<style scoped>
.scrollbar-right {
  direction: rtl;
}

.scrollbar-right > div {
  direction: ltr;
}

/* برای مرورگرهای WebKit (Chrome, Safari) */
.scrollbar-right::-webkit-scrollbar {
  width: 8px;
}

.scrollbar-right::-webkit-scrollbar-track {
  background: #f1f1f1;
}

.scrollbar-right::-webkit-scrollbar-thumb {
  background: #888;
  border-radius: 4px;
}

.scrollbar-right::-webkit-scrollbar-thumb:hover {
  background: #555;
}
</style> 
