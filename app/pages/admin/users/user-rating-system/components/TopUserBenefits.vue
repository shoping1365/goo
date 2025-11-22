<template>
  <div class="space-y-6">
    <!-- تنظیمات مزایا -->
    <div class="bg-white rounded-lg shadow">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-medium text-gray-900">تنظیمات مزایای کاربران برتر</h3>
      </div>
      <div class="px-4 py-4">
        <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
          <!-- مزایای تخفیف -->
          <div class="space-y-4">
            <h4 class="font-medium text-gray-900">مزایای تخفیف</h4>
            <div class="space-y-3">
              <div class="flex items-center justify-between p-3 bg-green-50 rounded-lg">
                <div>
                  <span class="text-sm font-medium text-gray-900">تخفیف ویژه</span>
                  <p class="text-xs text-gray-500">درصد تخفیف برای کاربران برتر</p>
                </div>
                <input v-model="benefits.discount.percentage" type="number" min="0" max="100" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
              </div>
              <div class="flex items-center justify-between p-3 bg-green-50 rounded-lg">
                <div>
                  <span class="text-sm font-medium text-gray-900">حداقل مبلغ سفارش</span>
                  <p class="text-xs text-gray-500">برای اعمال تخفیف</p>
                </div>
                <input v-model="benefits.discount.minOrderAmount" type="number" min="0" class="w-24 px-2 py-1 border border-gray-300 rounded text-center">
              </div>
              <div class="flex items-center justify-between p-3 bg-green-50 rounded-lg">
                <div>
                  <span class="text-sm font-medium text-gray-900">حداکثر تخفیف</span>
                  <p class="text-xs text-gray-500">مبلغ حداکثر تخفیف</p>
                </div>
                <input v-model="benefits.discount.maxDiscountAmount" type="number" min="0" class="w-24 px-2 py-1 border border-gray-300 rounded text-center">
              </div>
            </div>
          </div>

          <!-- مزایای ارسال -->
          <div class="space-y-4">
            <h4 class="font-medium text-gray-900">مزایای ارسال</h4>
            <div class="space-y-3">
              <div class="flex items-center justify-between p-3 bg-blue-50 rounded-lg">
                <div>
                  <span class="text-sm font-medium text-gray-900">ارسال رایگان</span>
                  <p class="text-xs text-gray-500">برای سفارش‌های بالای مبلغ مشخص</p>
                </div>
                <input v-model="benefits.shipping.freeShippingThreshold" type="number" min="0" class="w-24 px-2 py-1 border border-gray-300 rounded text-center">
              </div>
              <div class="flex items-center justify-between p-3 bg-blue-50 rounded-lg">
                <div>
                  <span class="text-sm font-medium text-gray-900">ارسال سریع</span>
                  <p class="text-xs text-gray-500">کاهش زمان ارسال (روز)</p>
                </div>
                <input v-model="benefits.shipping.fastShippingDays" type="number" min="0" max="7" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
              </div>
              <div class="flex items-center justify-between p-3 bg-blue-50 rounded-lg">
                <div>
                  <span class="text-sm font-medium text-gray-900">پیگیری پیشرفته</span>
                  <p class="text-xs text-gray-500">دسترسی به اطلاعات ارسال پیشرفته</p>
                </div>
                <label class="flex items-center">
                  <input v-model="benefits.shipping.advancedTracking" type="checkbox" class="mr-2 rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                  <span class="text-sm text-gray-700">فعال</span>
                </label>
              </div>
            </div>
          </div>
        </div>

        <!-- مزایای پشتیبانی -->
        <div class="mt-6">
          <h4 class="font-medium text-gray-900 mb-4">مزایای پشتیبانی</h4>
          <div class="grid grid-cols-1 md:grid-cols-3 gapx-4 py-4">
            <div class="flex items-center justify-between p-3 bg-purple-50 rounded-lg">
              <div>
                <span class="text-sm font-medium text-gray-900">پشتیبانی ویژه</span>
                <p class="text-xs text-gray-500">دسترسی به خط پشتیبانی ویژه</p>
              </div>
              <label class="flex items-center">
                <input v-model="benefits.support.prioritySupport" type="checkbox" class="mr-2 rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <span class="text-sm text-gray-700">فعال</span>
              </label>
            </div>
            <div class="flex items-center justify-between p-3 bg-purple-50 rounded-lg">
              <div>
                <span class="text-sm font-medium text-gray-900">زمان پاسخ سریع</span>
                <p class="text-xs text-gray-500">حداکثر زمان پاسخ (ساعت)</p>
              </div>
              <input v-model="benefits.support.responseTime" type="number" min="1" max="24" class="w-20 px-2 py-1 border border-gray-300 rounded text-center">
            </div>
            <div class="flex items-center justify-between p-3 bg-purple-50 rounded-lg">
              <div>
                <span class="text-sm font-medium text-gray-900">چت زنده</span>
                <p class="text-xs text-gray-500">دسترسی به چت زنده</p>
              </div>
              <label class="flex items-center">
                <input v-model="benefits.support.liveChat" type="checkbox" class="mr-2 rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <span class="text-sm text-gray-700">فعال</span>
              </label>
            </div>
          </div>
        </div>

        <!-- مزایای محتوا -->
        <div class="mt-6">
          <h4 class="font-medium text-gray-900 mb-4">مزایای محتوا</h4>
          <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
            <div class="flex items-center justify-between p-3 bg-yellow-50 rounded-lg">
              <div>
                <span class="text-sm font-medium text-gray-900">دسترسی زودهنگام</span>
                <p class="text-xs text-gray-500">دسترسی به محصولات جدید</p>
              </div>
              <label class="flex items-center">
                <input v-model="benefits.content.earlyAccess" type="checkbox" class="mr-2 rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <span class="text-sm text-gray-700">فعال</span>
              </label>
            </div>
            <div class="flex items-center justify-between p-3 bg-yellow-50 rounded-lg">
              <div>
                <span class="text-sm font-medium text-gray-900">محتوای انحصاری</span>
                <p class="text-xs text-gray-500">دسترسی به محتوای ویژه</p>
              </div>
              <label class="flex items-center">
                <input v-model="benefits.content.exclusiveContent" type="checkbox" class="mr-2 rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                <span class="text-sm text-gray-700">فعال</span>
              </label>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- لیست کاربران برتر -->
    <div class="bg-white rounded-lg shadow">
      <div class="px-6 py-4 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <h3 class="text-lg font-medium text-gray-900">کاربران برتر</h3>
          <div class="flex space-x-3 space-x-reverse">
            <select v-model="topUserFilter" class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
              <option value="all">همه سطوح</option>
              <option value="gold">طلایی</option>
              <option value="platinum">پلاتینیوم</option>
              <option value="diamond">الماس</option>
            </select>
            <button class="bg-green-600 text-white px-4 py-2 rounded-md hover:bg-green-700 transition-colors" @click="exportTopUsers">
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
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">سطح</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">امتیاز</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مزایای فعال</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">آخرین استفاده</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="user in filteredTopUsers" :key="user.id" class="hover:bg-gray-50">
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
                <span :class="getLevelClass(user.level)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                  {{ getLevelText(user.level) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="text-sm font-medium text-gray-900">{{ user.score }}</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex flex-wrap gap-1">
                  <span v-if="user.activeBenefits.discount" class="text-xs bg-green-100 text-green-800 px-2 py-1 rounded">تخفیف</span>
                  <span v-if="user.activeBenefits.freeShipping" class="text-xs bg-blue-100 text-blue-800 px-2 py-1 rounded">ارسال رایگان</span>
                  <span v-if="user.activeBenefits.prioritySupport" class="text-xs bg-purple-100 text-purple-800 px-2 py-1 rounded">پشتیبانی ویژه</span>
                  <span v-if="user.activeBenefits.earlyAccess" class="text-xs bg-yellow-100 text-yellow-800 px-2 py-1 rounded">دسترسی زودهنگام</span>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ formatDate(user.lastBenefitUsage) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 mr-3" @click="viewUserBenefits(user)">مزایا</button>
                <button class="text-green-600 hover:text-green-900" @click="editUserBenefits(user)">ویرایش</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- آمار استفاده از مزایا -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gapx-4 py-4">
      <!-- نمودار استفاده از مزایا -->
      <div class="bg-white rounded-lg shadow">
        <div class="px-6 py-4 border-b border-gray-200">
          <h3 class="text-lg font-medium text-gray-900">آمار استفاده از مزایا</h3>
        </div>
        <div class="px-4 py-4">
          <div class="space-y-4">
            <div v-for="benefit in benefitUsageStats" :key="benefit.name" class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
              <div class="flex items-center">
                <div :class="benefit.color" class="w-4 h-4 rounded-full mr-3"></div>
                <div>
                  <div class="text-sm font-medium text-gray-900">{{ benefit.name }}</div>
                  <div class="text-xs text-gray-500">{{ benefit.description }}</div>
                </div>
              </div>
              <div class="text-right">
                <div class="text-lg font-semibold text-gray-900">{{ benefit.usageCount }}</div>
                <div class="text-sm text-gray-500">{{ benefit.percentage }}%</div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- مزایای محبوب -->
      <div class="bg-white rounded-lg shadow">
        <div class="px-6 py-4 border-b border-gray-200">
          <h3 class="text-lg font-medium text-gray-900">مزایای محبوب</h3>
        </div>
        <div class="px-4 py-4">
          <div class="space-y-4">
            <div v-for="benefit in popularBenefits" :key="benefit.name" class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
              <div>
                <div class="text-sm font-medium text-gray-900">{{ benefit.name }}</div>
                <div class="text-xs text-gray-500">{{ benefit.description }}</div>
              </div>
              <div class="text-right">
                <div class="text-lg font-semibold text-gray-900">{{ benefit.usageCount }}</div>
                <div class="text-sm text-gray-500">استفاده</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تاریخچه استفاده از مزایا -->
    <div class="bg-white rounded-lg shadow">
      <div class="px-6 py-4 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <h3 class="text-lg font-medium text-gray-900">تاریخچه استفاده از مزایا</h3>
          <div class="flex space-x-3 space-x-reverse">
            <select v-model="benefitHistoryFilter" class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
              <option value="all">همه مزایا</option>
              <option value="discount">تخفیف</option>
              <option value="shipping">ارسال</option>
              <option value="support">پشتیبانی</option>
              <option value="content">محتوا</option>
            </select>
            <button class="bg-green-600 text-white px-4 py-2 rounded-md hover:bg-green-700 transition-colors" @click="exportBenefitHistory">
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
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مزیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">جزئیات</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مبلغ</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="record in filteredBenefitHistory" :key="record.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <img :src="record.user.avatar || '/default-avatar.png'" :alt="record.user.name" class="w-8 h-8 rounded-full mr-3">
                  <div>
                    <div class="text-sm font-medium text-gray-900">{{ record.user.name }}</div>
                    <div class="text-sm text-gray-500">{{ record.user.email }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getBenefitClass(record.benefit)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                  {{ getBenefitText(record.benefit) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-900">{{ record.details }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="text-sm font-medium text-gray-900">{{ record.amount }} تومان</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ formatDate(record.createdAt) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusClass(record.status)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                  {{ getStatusText(record.status) }}
                </span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, watch } from 'vue';

interface User {
  id: number;
  name: string;
  email: string;
  avatar: string;
  level: string;
  score: number;
  activeBenefits: Record<string, boolean>;
  lastBenefitUsage: string;
  [key: string]: unknown;
}

interface BenefitHistory {
  id: number;
  user: { name: string; email: string; avatar: string };
  benefit: string;
  details: string;
  amount: number;
  status: string;
  createdAt: string;
}

// Props
defineProps<{
  users: User[]
}>()

// Emits
const emit = defineEmits<{
  saveBenefits: [data: Record<string, unknown>]
  exportData: [type: string, data: Record<string, unknown>]
}>()

// Reactive data
const topUserFilter = ref('all')
const benefitHistoryFilter = ref('all')

const benefits = ref({
  discount: {
    percentage: 15,
    minOrderAmount: 100000,
    maxDiscountAmount: 50000
  },
  shipping: {
    freeShippingThreshold: 200000,
    fastShippingDays: 2,
    advancedTracking: true
  },
  support: {
    prioritySupport: true,
    responseTime: 2,
    liveChat: true
  },
  content: {
    earlyAccess: true,
    exclusiveContent: true
  }
})

// Sample data
const topUsers = ref<User[]>([
  {
    id: 1,
    name: 'علی احمدی',
    email: 'ali@example.com',
    avatar: '/avatars/ali.jpg',
    level: 'diamond',
    score: 2500,
    activeBenefits: {
      discount: true,
      freeShipping: true,
      prioritySupport: true,
      earlyAccess: true
    },
    lastBenefitUsage: '2024-01-15T10:30:00Z'
  },
  {
    id: 2,
    name: 'فاطمه محمدی',
    email: 'fateme@example.com',
    avatar: '/avatars/fateme.jpg',
    level: 'platinum',
    score: 1500,
    activeBenefits: {
      discount: true,
      freeShipping: true,
      prioritySupport: false,
      earlyAccess: false
    },
    lastBenefitUsage: '2024-01-14T15:45:00Z'
  }
])

const benefitUsageStats = ref([
  { name: 'تخفیف ویژه', description: 'استفاده از تخفیف', usageCount: 45, percentage: 75, color: 'bg-green-400' },
  { name: 'ارسال رایگان', description: 'استفاده از ارسال رایگان', usageCount: 32, percentage: 53, color: 'bg-blue-400' },
  { name: 'پشتیبانی ویژه', description: 'استفاده از پشتیبانی ویژه', usageCount: 18, percentage: 30, color: 'bg-purple-400' },
  { name: 'دسترسی زودهنگام', description: 'استفاده از دسترسی زودهنگام', usageCount: 25, percentage: 42, color: 'bg-yellow-400' }
])

const popularBenefits = ref([
  { name: 'تخفیف ویژه', description: '15% تخفیف روی سفارش‌ها', usageCount: 45 },
  { name: 'ارسال رایگان', description: 'ارسال رایگان برای سفارش‌های بالای 200 هزار تومان', usageCount: 32 },
  { name: 'پشتیبانی ویژه', description: 'پشتیبانی 24/7 با اولویت بالا', usageCount: 18 },
  { name: 'دسترسی زودهنگام', description: 'دسترسی به محصولات جدید قبل از عموم', usageCount: 25 }
])

const benefitHistory = ref<BenefitHistory[]>([
  {
    id: 1,
    user: { name: 'علی احمدی', email: 'ali@example.com', avatar: '/avatars/ali.jpg' },
    benefit: 'discount',
    details: 'تخفیف 15% روی سفارش #12345',
    amount: 15000,
    status: 'used',
    createdAt: '2024-01-15T10:30:00Z'
  },
  {
    id: 2,
    user: { name: 'فاطمه محمدی', email: 'fateme@example.com', avatar: '/avatars/fateme.jpg' },
    benefit: 'shipping',
    details: 'ارسال رایگان برای سفارش #12346',
    amount: 25000,
    status: 'used',
    createdAt: '2024-01-14T15:45:00Z'
  }
])

// Computed properties
const filteredTopUsers = computed(() => {
  let filtered = topUsers.value

  if (topUserFilter.value !== 'all') {
    filtered = filtered.filter(user => user.level === topUserFilter.value)
  }

  return filtered
})

const filteredBenefitHistory = computed(() => {
  let filtered = benefitHistory.value

  if (benefitHistoryFilter.value !== 'all') {
    filtered = filtered.filter(record => record.benefit === benefitHistoryFilter.value)
  }

  return filtered
})

// Methods
const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleDateString('fa-IR')
}

const getLevelClass = (level: string) => {
  switch (level) {
    case 'gold':
      return 'bg-yellow-100 text-yellow-800'
    case 'platinum':
      return 'bg-blue-100 text-blue-800'
    case 'diamond':
      return 'bg-purple-100 text-purple-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getLevelText = (level: string) => {
  switch (level) {
    case 'gold':
      return 'طلایی'
    case 'platinum':
      return 'پلاتینیوم'
    case 'diamond':
      return 'الماس'
    default:
      return 'نامشخص'
  }
}

const getBenefitClass = (benefit: string) => {
  switch (benefit) {
    case 'discount':
      return 'bg-green-100 text-green-800'
    case 'shipping':
      return 'bg-blue-100 text-blue-800'
    case 'support':
      return 'bg-purple-100 text-purple-800'
    case 'content':
      return 'bg-yellow-100 text-yellow-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getBenefitText = (benefit: string) => {
  switch (benefit) {
    case 'discount':
      return 'تخفیف'
    case 'shipping':
      return 'ارسال'
    case 'support':
      return 'پشتیبانی'
    case 'content':
      return 'محتوا'
    default:
      return 'نامشخص'
  }
}

const getStatusClass = (status: string) => {
  switch (status) {
    case 'used':
      return 'bg-green-100 text-green-800'
    case 'pending':
      return 'bg-yellow-100 text-yellow-800'
    case 'expired':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getStatusText = (status: string) => {
  switch (status) {
    case 'used':
      return 'استفاده شده'
    case 'pending':
      return 'در انتظار'
    case 'expired':
      return 'منقضی شده'
    default:
      return 'نامشخص'
  }
}

const viewUserBenefits = (_user: User) => {
  // View user benefits
}

const editUserBenefits = (_user: User) => {
  // Edit user benefits
}

const exportTopUsers = () => {
  emit('exportData', 'topUsers', {
    users: filteredTopUsers.value,
    filter: topUserFilter.value
  })
}

const exportBenefitHistory = () => {
  emit('exportData', 'benefitHistory', {
    history: filteredBenefitHistory.value,
    filter: benefitHistoryFilter.value
  })
}

// Watchers
watch(benefits, (newBenefits) => {
  emit('saveBenefits', newBenefits)
}, { deep: true })
</script> 
