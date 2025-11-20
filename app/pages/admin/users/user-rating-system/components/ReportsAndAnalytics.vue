<template>
  <div class="space-y-6">
    <!-- کارت‌های آمار -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <div class="bg-white rounded-lg shadow p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-blue-100 rounded-md flex items-center justify-center">
              <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197m13.5-9a2.5 2.5 0 11-5 0 2.5 2.5 0 015 0z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-500">کل کاربران</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.totalUsers }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-green-100 rounded-md flex items-center justify-center">
              <svg class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-500">میانگین امتیاز</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.averageScore }}/1000</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-yellow-100 rounded-md flex items-center justify-center">
              <svg class="w-5 h-5 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-500">کاربران برتر</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.topUsers }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-red-100 rounded-md flex items-center justify-center">
              <svg class="w-5 h-5 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-500">کاربران مسدود</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.blockedUsers }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- نمودار توزیع امتیازات -->
    <div class="bg-white rounded-lg shadow">
      <div class="px-6 py-4 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <h3 class="text-lg font-medium text-gray-900">توزیع امتیازات</h3>
          <div class="flex space-x-3 space-x-reverse">
            <select v-model="selectedPeriod" class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
              <option value="all">همه زمان‌ها</option>
              <option value="month">ماه جاری</option>
              <option value="week">هفته جاری</option>
              <option value="day">امروز</option>
            </select>
            <button class="bg-green-600 text-white px-4 py-2 rounded-md hover:bg-green-700 transition-colors" @click="exportChart">
              خروجی نمودار
            </button>
          </div>
        </div>
      </div>
      <div class="p-6">
        <div class="space-y-4">
          <div v-for="(level, index) in levelDistribution" :key="index" class="flex items-center">
            <span class="w-20 text-sm text-gray-600">{{ level.name }}</span>
            <div class="flex-1 mx-4 bg-gray-200 rounded-full h-3">
              <div :class="level.color" class="h-3 rounded-full transition-all duration-300" :style="{ width: level.percentage + '%' }"></div>
            </div>
            <span class="w-16 text-sm text-gray-600 text-left">{{ level.count }} نفر</span>
            <span class="w-12 text-sm text-gray-500 text-left">({{ level.percentage }}%)</span>
          </div>
        </div>
      </div>
    </div>

    <!-- نمودار روند تغییرات -->
    <div class="bg-white rounded-lg shadow">
      <div class="px-6 py-4 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <h3 class="text-lg font-medium text-gray-900">روند تغییرات امتیاز</h3>
          <div class="flex space-x-3 space-x-reverse">
            <select v-model="trendPeriod" class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
              <option value="7">7 روز گذشته</option>
              <option value="30">30 روز گذشته</option>
              <option value="90">90 روز گذشته</option>
            </select>
          </div>
        </div>
      </div>
      <div class="p-6">
        <div class="h-64 flex items-end justify-between space-x-2 space-x-reverse">
          <div v-for="(point, index) in trendData" :key="index" class="flex-1 flex flex-col items-center">
            <div class="w-full bg-blue-200 rounded-t transition-all duration-300 hover:bg-blue-300" :style="{ height: (point.value / maxTrendValue) * 200 + 'px' }"></div>
            <span class="text-xs text-gray-500 mt-2">{{ point.label }}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- گزارش فعالیت‌های کاربران برتر -->
    <div class="bg-white rounded-lg shadow">
      <div class="px-6 py-4 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <h3 class="text-lg font-medium text-gray-900">فعالیت‌های کاربران برتر</h3>
          <button class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors" @click="exportTopUsersReport">
            خروجی اکسل
          </button>
        </div>
      </div>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کاربر</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">امتیاز</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">فعالیت‌های این ماه</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">آخرین فعالیت</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="user in topUsers" :key="user.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="">
                    <img class="h-10 w-10 rounded-full" :src="user.avatar || '/default-avatar.png'" :alt="user.name">
                  </div>
                  <div class="mr-4">
                    <div class="text-sm font-medium text-gray-900">{{ user.name }}</div>
                    <div class="text-sm text-gray-500">{{ user.email }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="text-sm font-medium text-gray-900">{{ user.score }}</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-900">
                  <div>نظرات: {{ user.monthlyReviews }}</div>
                  <div>خریدها: {{ user.monthlyPurchases }}</div>
                  <div>ارجاعات: {{ user.monthlyReferrals }}</div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ formatDate(user.lastActivity) }}
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- گزارش‌های پیشرفته -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- گزارش فعالیت‌های روزانه -->
      <div class="bg-white rounded-lg shadow">
        <div class="px-6 py-4 border-b border-gray-200">
          <h3 class="text-lg font-medium text-gray-900">فعالیت‌های روزانه</h3>
        </div>
        <div class="p-6">
          <div class="space-y-4">
            <div v-for="activity in dailyActivities" :key="activity.type" class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
              <div>
                <h4 class="font-medium text-gray-900">{{ activity.name }}</h4>
                <p class="text-sm text-gray-500">{{ activity.description }}</p>
              </div>
              <div class="text-right">
                <div class="text-lg font-semibold text-gray-900">{{ activity.count }}</div>
                <div class="text-sm text-gray-500">{{ activity.change }}%</div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- گزارش سطوح رتبه‌بندی -->
      <div class="bg-white rounded-lg shadow">
        <div class="px-6 py-4 border-b border-gray-200">
          <h3 class="text-lg font-medium text-gray-900">آمار سطوح رتبه‌بندی</h3>
        </div>
        <div class="p-6">
          <div class="space-y-4">
            <div v-for="level in levelStats" :key="level.name" class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
              <div class="flex items-center">
                <div :class="level.color" class="w-4 h-4 rounded-full mr-3"></div>
                <div>
                  <h4 class="font-medium text-gray-900">{{ level.name }}</h4>
                  <p class="text-sm text-gray-500">امتیاز: {{ level.minScore }} - {{ level.maxScore }}</p>
                </div>
              </div>
              <div class="text-right">
                <div class="text-lg font-semibold text-gray-900">{{ level.count }}</div>
                <div class="text-sm text-gray-500">{{ level.percentage }}%</div>
              </div>
            </div>
          </div>
        </div>
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
  score: number;
  monthlyReviews: number;
  monthlyPurchases: number;
  monthlyReferrals: number;
  lastActivity: string;
  [key: string]: unknown;
}

interface Stats {
  totalUsers: number;
  averageScore: number;
  topUsers: number;
  blockedUsers: number;
}

interface LevelDistribution {
  name: string;
  count: number;
  percentage: number;
  color: string;
}

interface TrendData {
  label: string;
  value: number;
}

interface Activity {
  type: string;
  name: string;
  description: string;
  count: number;
  change: number;
}

interface LevelStat {
  name: string;
  minScore: number;
  maxScore: number;
  count: number;
  percentage: number;
  color: string;
}

// Props
defineProps<{
  users: User[]
}>()

// Emits
const emit = defineEmits<{
  exportData: [type: string, data: Record<string, unknown>]
}>()

// Reactive data
const selectedPeriod = ref('all')
const trendPeriod = ref('30')

// Sample data
const stats = ref<Stats>({
  totalUsers: 1250,
  averageScore: 450,
  topUsers: 89,
  blockedUsers: 12
})

const levelDistribution = ref<LevelDistribution[]>([
  { name: 'برنزی', count: 450, percentage: 36, color: 'bg-orange-400' },
  { name: 'نقره‌ای', count: 380, percentage: 30.4, color: 'bg-gray-400' },
  { name: 'طلایی', count: 250, percentage: 20, color: 'bg-yellow-400' },
  { name: 'پلاتینیوم', count: 120, percentage: 9.6, color: 'bg-blue-400' },
  { name: 'الماس', count: 50, percentage: 4, color: 'bg-purple-400' }
])

const trendData = ref<TrendData[]>([
  { label: '1', value: 120 },
  { label: '2', value: 135 },
  { label: '3', value: 110 },
  { label: '4', value: 145 },
  { label: '5', value: 160 },
  { label: '6', value: 140 },
  { label: '7', value: 155 }
])

const topUsers = ref<User[]>([
  {
    id: 1,
    name: 'علی احمدی',
    email: 'ali@example.com',
    avatar: '/avatars/ali.jpg',
    score: 850,
    monthlyReviews: 15,
    monthlyPurchases: 8,
    monthlyReferrals: 3,
    lastActivity: '2024-01-15T10:30:00Z'
  },
  {
    id: 2,
    name: 'فاطمه محمدی',
    email: 'fateme@example.com',
    avatar: '/avatars/fateme.jpg',
    score: 720,
    monthlyReviews: 12,
    monthlyPurchases: 6,
    monthlyReferrals: 2,
    lastActivity: '2024-01-14T15:45:00Z'
  }
])

const dailyActivities = ref<Activity[]>([
  { type: 'reviews', name: 'نظرات جدید', description: 'تعداد نظرات ثبت شده', count: 45, change: 12 },
  { type: 'purchases', name: 'خریدهای جدید', description: 'تعداد خریدهای انجام شده', count: 23, change: 8 },
  { type: 'referrals', name: 'ارجاعات جدید', description: 'تعداد ارجاعات موفق', count: 8, change: 25 },
  { type: 'logins', name: 'ورودهای جدید', description: 'تعداد ورودهای روزانه', count: 156, change: -5 }
])

const levelStats = ref<LevelStat[]>([
  { name: 'برنزی', minScore: 0, maxScore: 100, count: 450, percentage: 36, color: 'bg-orange-400' },
  { name: 'نقره‌ای', minScore: 101, maxScore: 500, count: 380, percentage: 30.4, color: 'bg-gray-400' },
  { name: 'طلایی', minScore: 501, maxScore: 1000, count: 250, percentage: 20, color: 'bg-yellow-400' },
  { name: 'پلاتینیوم', minScore: 1001, maxScore: 2000, count: 120, percentage: 9.6, color: 'bg-blue-400' },
  { name: 'الماس', minScore: 2001, maxScore: 9999, count: 50, percentage: 4, color: 'bg-purple-400' }
])

// Computed properties
const maxTrendValue = computed(() => {
  return Math.max(...trendData.value.map(item => item.value))
})

// Methods
const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleDateString('fa-IR')
}

const exportChart = () => {
  emit('exportData', 'chart', {
    levelDistribution: levelDistribution.value,
    period: selectedPeriod.value
  })
}

const exportTopUsersReport = () => {
  emit('exportData', 'topUsers', {
    users: topUsers.value,
    period: selectedPeriod.value
  })
}

// Watchers
watch(selectedPeriod, (_newPeriod) => {
  // در اینجا می‌توانید داده‌ها را بر اساس دوره انتخاب شده فیلتر کنید
})

watch(trendPeriod, (_newPeriod) => {
  // در اینجا می‌توانید داده‌های روند را بر اساس دوره انتخاب شده به‌روزرسانی کنید
})
</script> 
