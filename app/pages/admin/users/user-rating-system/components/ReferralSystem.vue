<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="bg-white rounded-lg shadow px-4 py-4">
      <div class="flex items-center justify-between">
        <div>
          <h2 class="text-xl font-bold text-gray-900">سیستم ارجاع پیشرفته</h2>
          <p class="mt-1 text-sm text-gray-500">مدیریت ارجاعات کاربران و امتیازدهی بر اساس کیفیت ارجاع</p>
        </div>
        <div class="flex space-x-3 space-x-reverse">
          <button class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors" @click="generateReferralCodes">
            تولید کد ارجاع
          </button>
          <button class="bg-green-600 text-white px-4 py-2 rounded-md hover:bg-green-700 transition-colors" @click="exportReferralReport">
            خروجی گزارش
          </button>
        </div>
      </div>
    </div>

    <!-- Referral Settings -->
    <div class="bg-white rounded-lg shadow px-4 py-4">
      <h3 class="text-lg font-medium text-gray-900 mb-4">تنظیمات سیستم ارجاع</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gapx-4 py-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">امتیاز برای ارجاع موفق</label>
          <input v-model="referralSettings.successfulReferralScore" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">امتیاز برای خرید ارجاع شده</label>
          <input v-model="referralSettings.purchaseBonusScore" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">درصد تخفیف برای ارجاع کننده</label>
          <input v-model="referralSettings.referrerDiscount" type="number" step="0.1" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">درصد تخفیف برای ارجاع شده</label>
          <input v-model="referralSettings.referredDiscount" type="number" step="0.1" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداقل مبلغ خرید برای اعتبار</label>
          <input v-model="referralSettings.minPurchaseAmount" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">مدت اعتبار کد ارجاع (روز)</label>
          <input v-model="referralSettings.codeExpiryDays" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
      </div>
      
      <div class="mt-6 space-y-4">
        <label class="flex items-center">
          <input v-model="referralSettings.autoGenerateCodes" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
          <span class="mr-2 text-sm text-gray-700">تولید خودکار کد ارجاع برای کاربران جدید</span>
        </label>
        <label class="flex items-center">
          <input v-model="referralSettings.trackReferralQuality" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
          <span class="mr-2 text-sm text-gray-700">پیگیری کیفیت ارجاعات</span>
        </label>
        <label class="flex items-center">
          <input v-model="referralSettings.notifyOnReferral" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
          <span class="mr-2 text-sm text-gray-700">اعلان در صورت ارجاع موفق</span>
        </label>
      </div>
      
      <div class="mt-4 flex space-x-3 space-x-reverse">
        <button class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors" @click="saveReferralSettings">
          ذخیره تنظیمات
        </button>
        <button class="bg-purple-600 text-white px-4 py-2 rounded-md hover:bg-purple-700 transition-colors" @click="testReferralSystem">
          تست سیستم
        </button>
      </div>
    </div>

    <!-- Referral Statistics -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gapx-4 py-4">
      <div class="bg-white rounded-lg shadow px-4 py-4">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-green-100 rounded-full flex items-center justify-center">
              <svg class="w-5 h-5 text-green-600" fill="currentColor" viewBox="0 0 20 20">
                <path d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-500">ارجاعات موفق</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.successfulReferrals }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow px-4 py-4">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center">
              <svg class="w-5 h-5 text-blue-600" fill="currentColor" viewBox="0 0 20 20">
                <path d="M13 6a3 3 0 11-6 0 3 3 0 016 0zM18 8a2 2 0 11-4 0 2 2 0 014 0zM14 15a4 4 0 00-8 0v3h8v-3z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-500">کاربران ارجاع شده</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.referredUsers }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow px-4 py-4">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-yellow-100 rounded-full flex items-center justify-center">
              <svg class="w-5 h-5 text-yellow-600" fill="currentColor" viewBox="0 0 20 20">
                <path d="M8.433 7.418c.155-.103.346-.196.567-.267v1.698a2.305 2.305 0 01-.567-.267C8.07 8.34 8 8.114 8 8c0-.114.07-.34.433-.582zM11 12.849v-1.698c.22.071.412.164.567.267.364.243.433.468.433.582 0 .114-.07.34-.433.582a2.305 2.305 0 01-.567.267z"></path>
                <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-13a1 1 0 10-2 0v.092a4.535 4.535 0 00-1.676.662C6.602 6.234 6 7.009 6 8c0 .99.602 1.765 1.324 2.246.48.32 1.054.545 1.676.662v1.941c-.391-.127-.68-.317-.843-.504a1 1 0 10-1.51 1.31c.562.649 1.413 1.076 2.353 1.253V15a1 1 0 102 0v-.092a4.535 4.535 0 001.676-.662C13.398 13.766 14 12.991 14 12c0-.99-.602-1.765-1.324-2.246A4.535 4.535 0 0011 9.092V7.151c.391.127.68.317.843.504a1 1 0 101.511-1.31c-.563-.649-1.413-1.076-2.354-1.253V5z" clip-rule="evenodd"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-500">درآمد ارجاع</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.referralRevenue.toLocaleString() }} تومان</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow px-4 py-4">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-purple-100 rounded-full flex items-center justify-center">
              <svg class="w-5 h-5 text-purple-600" fill="currentColor" viewBox="0 0 20 20">
                <path d="M9 6a3 3 0 11-6 0 3 3 0 016 0zM17 6a3 3 0 11-6 0 3 3 0 016 0zM12.93 17c.046-.327.07-.66.07-1a6.97 6.97 0 00-1.5-4.33A5 5 0 0119 16v1h-6.07zM6 11a5 5 0 015 5v1H1v-1a5 5 0 015-5z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-500">نرخ تبدیل</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.conversionRate }}%</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Referral Management -->
    <div class="bg-white rounded-lg shadow px-4 py-4">
      <h3 class="text-lg font-medium text-gray-900 mb-4">مدیریت ارجاعات</h3>
      
      <!-- Search and Filter -->
      <div class="flex flex-col sm:flex-row gapx-4 py-4 mb-6">
        <div class="flex-1">
          <input v-model="searchQuery" type="text" placeholder="جستجو در ارجاعات..." class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
        </div>
        <select v-model="filterStatus" class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
          <option value="">همه وضعیت‌ها</option>
          <option value="successful">موفق</option>
          <option value="pending">در انتظار</option>
          <option value="expired">منقضی شده</option>
          <option value="cancelled">لغو شده</option>
        </select>
        <select v-model="filterQuality" class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
          <option value="">همه کیفیت‌ها</option>
          <option value="high">عالی</option>
          <option value="medium">متوسط</option>
          <option value="low">ضعیف</option>
        </select>
      </div>

      <!-- Referrals Table -->
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">ارجاع کننده</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">ارجاع شده</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کد ارجاع</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کیفیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">امتیاز</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="referral in paginatedReferrals" :key="referral.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <img :src="referral.referrerAvatar" :alt="referral.referrerName" class="w-8 h-8 rounded-full">
                  <div class="mr-3">
                    <div class="text-sm font-medium text-gray-900">{{ referral.referrerName }}</div>
                    <div class="text-sm text-gray-500">{{ referral.referrerEmail }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <img :src="referral.referredAvatar" :alt="referral.referredName" class="w-8 h-8 rounded-full">
                  <div class="mr-3">
                    <div class="text-sm font-medium text-gray-900">{{ referral.referredName }}</div>
                    <div class="text-sm text-gray-500">{{ referral.referredEmail }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-900 font-mono">{{ referral.referralCode }}</div>
                <div class="text-sm text-gray-500">{{ referral.codeType }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ formatDate(referral.createdAt) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusClass(referral.status)" class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium">
                  {{ getStatusText(referral.status) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getQualityClass(referral.quality)" class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium">
                  {{ getQualityText(referral.quality) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ referral.score }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3" @click="viewReferralDetails(referral)">مشاهده</button>
                <button class="text-green-600 hover:text-green-900 ml-3" @click="editReferral(referral)">ویرایش</button>
                <button class="text-red-600 hover:text-red-900" @click="cancelReferral(referral)">لغو</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Pagination -->
      <div class="flex items-center justify-between mt-6">
        <div class="flex items-center space-x-2 space-x-reverse">
          <span class="text-sm text-gray-700">نمایش {{ pagination.start }} تا {{ pagination.end }} از {{ pagination.total }} ارجاع</span>
        </div>
        <div class="flex space-x-2 space-x-reverse">
          <button :disabled="currentPage === 1" class="px-3 py-1 border border-gray-300 rounded-md text-sm disabled:opacity-50" @click="previousPage">
            قبلی
          </button>
          <button v-for="page in visiblePages" :key="page" :class="page === currentPage ? 'bg-blue-600 text-white' : 'bg-white text-gray-700'" class="px-3 py-1 border border-gray-300 rounded-md text-sm" @click="goToPage(page)">
            {{ page }}
          </button>
          <button :disabled="currentPage === totalPages" class="px-3 py-1 border border-gray-300 rounded-md text-sm disabled:opacity-50" @click="nextPage">
            بعدی
          </button>
        </div>
      </div>
    </div>

    <!-- Referral Codes Management -->
    <div class="bg-white rounded-lg shadow px-4 py-4">
      <h3 class="text-lg font-medium text-gray-900 mb-4">مدیریت کدهای ارجاع</h3>
      
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gapx-4 py-4">
        <div v-for="code in referralCodes" :key="code.id" class="border border-gray-200 rounded-lg px-4 py-4">
          <div class="flex items-center justify-between mb-3">
            <div class="flex items-center">
              <img :src="code.userAvatar" :alt="code.userName" class="w-8 h-8 rounded-full">
              <div class="mr-3">
                <div class="text-sm font-medium text-gray-900">{{ code.userName }}</div>
                <div class="text-sm text-gray-500">{{ code.userEmail }}</div>
              </div>
            </div>
            <span :class="getCodeStatusClass(code.status)" class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium">
              {{ getCodeStatusText(code.status) }}
            </span>
          </div>
          
          <div class="space-y-2">
            <div>
              <span class="text-sm text-gray-500">کد ارجاع:</span>
              <div class="text-sm font-mono text-gray-900 bg-gray-100 p-2 rounded">{{ code.code }}</div>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-500">استفاده شده:</span>
              <span class="text-gray-900">{{ code.usageCount }} بار</span>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-500">تاریخ انقضا:</span>
              <span class="text-gray-900">{{ formatDate(code.expiryDate) }}</span>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="text-blue-600 hover:text-blue-900 text-sm" @click="copyCode(code.code)">کپی</button>
            <button class="text-green-600 hover:text-green-900 text-sm" @click="regenerateCode(code.id)">تولید مجدد</button>
            <button class="text-red-600 hover:text-red-900 text-sm" @click="deactivateCode(code.id)">غیرفعال</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue';

interface Referral {
  id: number;
  referrerName: string;
  referrerEmail: string;
  referrerAvatar: string;
  referredName: string;
  referredEmail: string;
  referredAvatar: string;
  referralCode: string;
  codeType: string;
  createdAt: string;
  status: string;
  quality: string;
  score: number;
  [key: string]: unknown;
}

interface ReferralCode {
  id: number;
  userName: string;
  userEmail: string;
  userAvatar: string;
  code: string;
  status: string;
  usageCount: number;
  expiryDate: string;
}

// Props and Emits
const props = defineProps<{
  referrals?: Referral[]
}>()

const emit = defineEmits<{
  saveSettings: [settings: Record<string, unknown>]
  generateCodes: [users: Record<string, unknown>[]]
  exportReport: [data: Record<string, unknown>]
}>()

// Reactive data
const searchQuery = ref('')
const filterStatus = ref('')
const filterQuality = ref('')
const currentPage = ref(1)
const itemsPerPage = 10

// Referral Settings
const referralSettings = ref({
  successfulReferralScore: 50,
  purchaseBonusScore: 20,
  referrerDiscount: 10,
  referredDiscount: 5,
  minPurchaseAmount: 100000,
  codeExpiryDays: 30,
  autoGenerateCodes: true,
  trackReferralQuality: true,
  notifyOnReferral: true
})

// Sample referrals data
const localReferrals = ref<Referral[]>([
  {
    id: 1,
    referrerName: 'علی احمدی',
    referrerEmail: 'ali@example.com',
    referrerAvatar: '/avatars/ali.jpg',
    referredName: 'فاطمه محمدی',
    referredEmail: 'fateme@example.com',
    referredAvatar: '/avatars/fateme.jpg',
    referralCode: 'ALI2024',
    codeType: 'شخصی',
    createdAt: '2024-01-15T10:30:00Z',
    status: 'successful',
    quality: 'high',
    score: 70
  },
  {
    id: 2,
    referrerName: 'محمد رضایی',
    referrerEmail: 'mohammad@example.com',
    referrerAvatar: '/avatars/mohammad.jpg',
    referredName: 'زهرا کریمی',
    referredEmail: 'zahra@example.com',
    referredAvatar: '/avatars/zahra.jpg',
    referralCode: 'MOH2024',
    codeType: 'شخصی',
    createdAt: '2024-01-14T15:45:00Z',
    status: 'pending',
    quality: 'medium',
    score: 30
  }
])

// Sample referral codes
const referralCodes = ref<ReferralCode[]>([
  {
    id: 1,
    userName: 'علی احمدی',
    userEmail: 'ali@example.com',
    userAvatar: '/avatars/ali.jpg',
    code: 'ALI2024',
    status: 'active',
    usageCount: 5,
    expiryDate: '2024-02-15T10:30:00Z'
  },
  {
    id: 2,
    userName: 'فاطمه محمدی',
    userEmail: 'fateme@example.com',
    userAvatar: '/avatars/fateme.jpg',
    code: 'FAT2024',
    status: 'active',
    usageCount: 2,
    expiryDate: '2024-02-14T15:45:00Z'
  }
])

// Statistics
const stats = ref({
  successfulReferrals: 1250,
  referredUsers: 980,
  referralRevenue: 45000000,
  conversionRate: 78.4
})

// Pagination
const pagination = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage + 1
  const end = Math.min(currentPage.value * itemsPerPage, filteredReferrals.value.length)
  return {
    start,
    end,
    total: filteredReferrals.value.length
  }
})

const totalPages = computed(() => Math.ceil(filteredReferrals.value.length / itemsPerPage))

const visiblePages = computed(() => {
  const pages = []
  const start = Math.max(1, currentPage.value - 2)
  const end = Math.min(totalPages.value, currentPage.value + 2)
  
  for (let i = start; i <= end; i++) {
    pages.push(i)
  }
  return pages
})

// Filtered referrals
const filteredReferrals = computed(() => {
  let filtered = props.referrals || localReferrals.value

  if (searchQuery.value) {
    filtered = filtered.filter(referral => 
      referral.referrerName.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      referral.referredName.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      referral.referralCode.toLowerCase().includes(searchQuery.value.toLowerCase())
    )
  }

  if (filterStatus.value) {
    filtered = filtered.filter(referral => referral.status === filterStatus.value)
  }

  if (filterQuality.value) {
    filtered = filtered.filter(referral => referral.quality === filterQuality.value)
  }

  return filtered
})

const paginatedReferrals = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage
  const end = start + itemsPerPage
  return filteredReferrals.value.slice(start, end)
})

// Methods
const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleDateString('fa-IR')
}

const getStatusClass = (status: string) => {
  switch (status) {
    case 'successful':
      return 'bg-green-100 text-green-800'
    case 'pending':
      return 'bg-yellow-100 text-yellow-800'
    case 'expired':
      return 'bg-red-100 text-red-800'
    case 'cancelled':
      return 'bg-gray-100 text-gray-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getStatusText = (status: string) => {
  switch (status) {
    case 'successful':
      return 'موفق'
    case 'pending':
      return 'در انتظار'
    case 'expired':
      return 'منقضی شده'
    case 'cancelled':
      return 'لغو شده'
    default:
      return 'نامشخص'
  }
}

const getQualityClass = (quality: string) => {
  switch (quality) {
    case 'high':
      return 'bg-green-100 text-green-800'
    case 'medium':
      return 'bg-yellow-100 text-yellow-800'
    case 'low':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getQualityText = (quality: string) => {
  switch (quality) {
    case 'high':
      return 'عالی'
    case 'medium':
      return 'متوسط'
    case 'low':
      return 'ضعیف'
    default:
      return 'نامشخص'
  }
}

const getCodeStatusClass = (status: string) => {
  switch (status) {
    case 'active':
      return 'bg-green-100 text-green-800'
    case 'inactive':
      return 'bg-red-100 text-red-800'
    case 'expired':
      return 'bg-yellow-100 text-yellow-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getCodeStatusText = (status: string) => {
  switch (status) {
    case 'active':
      return 'فعال'
    case 'inactive':
      return 'غیرفعال'
    case 'expired':
      return 'منقضی شده'
    default:
      return 'نامشخص'
  }
}

const saveReferralSettings = () => {
  // API call to save referral settings
  emit('saveSettings', referralSettings.value)
}

const testReferralSystem = () => {
  // Test referral system functionality
}

const generateReferralCodes = () => {
  // Generate referral codes for users
  emit('generateCodes', [])
}

const exportReferralReport = () => {
  // Export referral report
}

const viewReferralDetails = (_referral: Referral) => {
  // Open referral details modal
}

const editReferral = (_referral: Referral) => {
  // Edit referral details
}

const cancelReferral = (referral: Referral) => {
  referral.status = 'cancelled'
  // API call to cancel referral
}

const copyCode = (code: string) => {
  navigator.clipboard.writeText(code)
}

const regenerateCode = (_codeId: number) => {
  // Regenerate referral code
}

const deactivateCode = (_codeId: number) => {
  // Deactivate referral code
}

const previousPage = () => {
  if (currentPage.value > 1) {
    currentPage.value--
  }
}

const nextPage = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value++
  }
}

const goToPage = (page: number) => {
  currentPage.value = page
}

// Lifecycle
// onMounted removed as it only contained console.log
</script> 
