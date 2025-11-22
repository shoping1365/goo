<template>
  <div class="p-6" dir="rtl">
    <div class="mb-6 bg-white p-6 rounded-lg shadow-md border border-gray-200">
      <h1 class="text-2xl font-bold text-gray-900 mb-2">آمار پشتیبانی</h1>
      <p class="text-gray-600">تحلیل و گزارش‌گیری از عملکرد سیستم پشتیبانی</p>
    </div>

    <!-- منوی تب‌ها -->
    <div class="mb-8">
      <nav class="flex flex-wrap gap-2 border-b border-gray-200">
        <button
v-for="tab in tabs" :key="tab.key" :class="['px-4 py-2 text-sm font-medium rounded-t-md focus:outline-none', activeTab === tab.key ? 'bg-white border-x border-t border-b-0 border-gray-200 text-blue-600' : 'bg-gray-100 text-gray-600 hover:bg-gray-200']"
          @click="activeTab = tab.key">
          {{ tab.label }}
        </button>
      </nav>
    </div>

    <!-- محتوای تب فعال -->
    <div>
      <div v-if="activeTab === 'main'">
        <!-- محتوای فعلی آمار کلی -->
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
          <TemplateCard
            title="تیکت‌های باز"
            :value="openTickets"
            variant="red"
          />
          <TemplateCard
            title="تیکت‌های حل شده"
            :value="resolvedTickets"
            variant="green"
          />
          <TemplateCard
            title="زمان پاسخ متوسط"
            :value="avgResponseTime"
            variant="blue"
          />
          <TemplateCard
            title="رضایت مشتریان"
            :value="customerSatisfaction"
            variant="cyan"
          />
        </div>
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-8 mb-8">
          <div class="bg-white rounded-lg shadow-md border border-gray-200">
            <div class="px-6 py-4 border-b border-gray-200">
              <h3 class="text-lg font-semibold text-gray-900">تیکت‌ها در طول زمان</h3>
            </div>
            <div class="p-6">
              <div class="h-64 bg-gray-50 rounded-lg flex items-center justify-center">
                <div class="text-center">
                  <svg class="w-16 h-16 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
                  </svg>
                  <p class="text-gray-500">نمودار تیکت‌ها در طول زمان</p>
                  <p class="text-sm text-gray-400">آخرین 30 روز</p>
                </div>
              </div>
            </div>
          </div>
          <div class="bg-white rounded-lg shadow-md border border-gray-200">
            <div class="px-6 py-4 border-b border-gray-200">
              <h3 class="text-lg font-semibold text-gray-900">توزیع تیکت‌ها بر اساس دسته‌بندی</h3>
            </div>
            <div class="p-6">
              <div class="h-64 bg-gray-50 rounded-lg flex items-center justify-center">
                <div class="text-center">
                  <svg class="w-16 h-16 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 3.055A9.001 9.001 0 1020.945 13H11V3.055z"></path>
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.488 9H15V3.512A9.025 9.025 0 0120.488 9z"></path>
                  </svg>
                  <p class="text-gray-500">نمودار دایره‌ای توزیع تیکت‌ها</p>
                  <p class="text-sm text-gray-400">بر اساس دسته‌بندی</p>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="grid grid-cols-1 lg:grid-cols-3 gap-8 mb-8">
          <div class="bg-white rounded-lg shadow-md border border-gray-200">
            <div class="px-6 py-4 border-b border-gray-200">
              <h3 class="text-lg font-semibold text-gray-900">عملکرد اپراتورها</h3>
            </div>
            <div class="p-6">
              <div class="space-y-4">
                <div v-for="operator in operatorStats" :key="operator.id" class="flex items-center justify-between">
                  <div class="flex items-center">
                    <div class="w-8 h-8 rounded-full bg-gray-300 flex items-center justify-center mr-3">
                      <span class="text-xs font-medium text-gray-700">{{ operator.name.charAt(0) }}</span>
                    </div>
                    <div>
                      <div class="text-sm font-medium text-gray-900">{{ operator.name }}</div>
                      <div class="text-xs text-gray-500">{{ operator.tickets }} تیکت</div>
                    </div>
                  </div>
                  <div class="text-right">
                    <div class="text-sm font-medium text-gray-900">{{ operator.avgTime }}</div>
                    <div class="text-xs text-gray-500">زمان پاسخ</div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="bg-white rounded-lg shadow-md border border-gray-200">
            <div class="px-6 py-4 border-b border-gray-200">
              <h3 class="text-lg font-semibold text-gray-900">دسته‌بندی تیکت‌ها</h3>
            </div>
            <div class="p-6">
              <div class="space-y-4">
                <div v-for="category in categoryStats" :key="category.name" class="flex items-center justify-between">
                  <div class="flex items-center">
                    <div :class="category.color" class="w-3 h-3 rounded-full mr-3"></div>
                    <span class="text-sm text-gray-900">{{ category.name }}</span>
                  </div>
                  <div class="text-right">
                    <div class="text-sm font-medium text-gray-900">{{ category.count }}</div>
                    <div class="text-xs text-gray-500">{{ category.percentage }}%</div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="bg-white rounded-lg shadow-md border border-gray-200">
            <div class="px-6 py-4 border-b border-gray-200">
              <h3 class="text-lg font-semibold text-gray-900">روند زمانی</h3>
            </div>
            <div class="p-6">
              <div class="space-y-4">
                <div v-for="trend in timeTrends" :key="trend.period" class="flex items-center justify-between">
                  <span class="text-sm text-gray-900">{{ trend.period }}</span>
                  <div class="text-right">
                    <div class="text-sm font-medium text-gray-900">{{ trend.tickets }}</div>
                    <div class="text-xs text-gray-500">تیکت</div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="bg-white rounded-lg shadow-md border border-gray-200">
          <div class="px-6 py-4 border-b border-gray-200">
            <h3 class="text-lg font-semibold text-gray-900">آمار تفصیلی</h3>
          </div>
          <div class="overflow-x-auto">
            <table class="min-w-full divide-y divide-gray-200">
              <thead class="bg-gray-50">
                <tr>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">شاخص</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">امروز</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">هفته گذشته</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">ماه گذشته</th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تغییر</th>
                </tr>
              </thead>
              <tbody class="bg-white divide-y divide-gray-200">
                <tr v-for="metric in detailedStats" :key="metric.name">
                  <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{{ metric.name }}</td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ metric.today }}</td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ metric.lastWeek }}</td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ metric.lastMonth }}</td>
                  <td class="px-6 py-4 whitespace-nowrap">
                    <span :class="getChangeClass(metric.change)" class="text-sm font-medium">
                      {{ metric.change > 0 ? '+' : '' }}{{ metric.change }}%
                    </span>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
      <CallStatistics v-else-if="activeTab === 'call-statistics'" />
      <OperatorStatistics v-else-if="activeTab === 'operator-statistics'" />
      <CustomerSatisfaction v-else-if="activeTab === 'customer-satisfaction'" />
      <CallArchive v-else-if="activeTab === 'call-archive'" />
      <UserArchive v-else-if="activeTab === 'user-archive'" />
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
</script>

<script setup>
import CallArchive from './call-archive.vue'
import CallStatistics from './call-statistics.vue'
import CustomerSatisfaction from './customer-satisfaction.vue'
import OperatorStatistics from './operator-statistics.vue'
import UserArchive from './user-archive.vue'

definePageMeta({ layout: 'admin-main', middleware: 'admin' })

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user: _user, hasPermission: _hasPermission } = useAuth()

const tabs = [
  { key: 'main', label: 'آمار کلی' },
  { key: 'call-statistics', label: 'آمار مکالمات' },
  { key: 'operator-statistics', label: 'آمار اپراتورها' },
  { key: 'customer-satisfaction', label: 'آمار رضایت مشتری' },
  { key: 'call-archive', label: 'آرشیو مکالمات' },
  { key: 'user-archive', label: 'آرشیو کاربران' },
]
const activeTab = ref('main')

// داده‌های نمونه
const openTickets = ref(24)
const resolvedTickets = ref(156)
const avgResponseTime = ref('2.5 دقیقه')
const customerSatisfaction = ref('4.8/5')

const operatorStats = ref([
  { id: 1, name: 'علی احمدی', tickets: 45, avgTime: '2.1 دقیقه' },
  { id: 2, name: 'فاطمه محمدی', tickets: 38, avgTime: '2.8 دقیقه' },
  { id: 3, name: 'محمد رضایی', tickets: 32, avgTime: '3.2 دقیقه' },
  { id: 4, name: 'سارا کریمی', tickets: 28, avgTime: '2.5 دقیقه' },
  { id: 5, name: 'حسن نوری', tickets: 41, avgTime: '1.9 دقیقه' }
])

const categoryStats = ref([
  { name: 'فنی', count: 45, percentage: 35, color: 'bg-blue-500' },
  { name: 'پرداخت', count: 28, percentage: 22, color: 'bg-green-500' },
  { name: 'حساب کاربری', count: 25, percentage: 19, color: 'bg-purple-500' },
  { name: 'عمومی', count: 20, percentage: 15, color: 'bg-orange-500' },
  { name: 'سایر', count: 12, percentage: 9, color: 'bg-gray-500' }
])

const timeTrends = ref([
  { period: 'امروز', tickets: 24 },
  { period: 'دیروز', tickets: 31 },
  { period: '2 روز پیش', tickets: 28 },
  { period: '3 روز پیش', tickets: 35 },
  { period: '4 روز پیش', tickets: 22 },
  { period: '5 روز پیش', tickets: 29 },
  { period: '6 روز پیش', tickets: 26 }
])

const detailedStats = ref([
  { name: 'تیکت‌های جدید', today: 24, lastWeek: 156, lastMonth: 642, change: 15 },
  { name: 'تیکت‌های حل شده', today: 22, lastWeek: 148, lastMonth: 598, change: -8 },
  { name: 'زمان پاسخ متوسط', today: '2.5 دقیقه', lastWeek: '2.8 دقیقه', lastMonth: '3.1 دقیقه', change: -12 },
  { name: 'رضایت مشتریان', today: '4.8/5', lastWeek: '4.7/5', lastMonth: '4.6/5', change: 4 },
  { name: 'تیکت‌های باز', today: 24, lastWeek: 18, lastMonth: 44, change: 33 }
])

function getChangeClass(change) {
  if (change > 0) {
    return 'text-green-600'
  } else if (change < 0) {
    return 'text-red-600'
  } else {
    return 'text-gray-600'
  }
}
</script> 

