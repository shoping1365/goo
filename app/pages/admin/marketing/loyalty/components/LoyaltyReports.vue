<template>
  <div class="bg-white rounded-lg shadow p-6">
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-xl font-semibold text-gray-900">گزارش‌ها و تحلیل</h2>
      <div class="flex space-x-2 space-x-reverse">
        <select 
          v-model="reportType"
          class="border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
        >
          <option value="activity">گزارش فعالیت</option>
          <option value="points">گزارش امتیازات</option>
          <option value="rewards">گزارش پاداش‌ها</option>
          <option value="levels">گزارش سطوح</option>
        </select>
        <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
          دانلود گزارش
        </button>
      </div>
    </div>

    <!-- نمودارها و آمار -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- نمودار روند امتیازات -->
      <div class="bg-gray-50 p-6 rounded-lg">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">روند امتیازات</h3>
        <div class="h-64 flex items-center justify-center text-gray-500">
          نمودار روند امتیازات (در حال توسعه)
        </div>
      </div>

      <!-- نمودار توزیع سطوح -->
      <div class="bg-gray-50 p-6 rounded-lg">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">توزیع سطوح</h3>
        <div class="h-64 flex items-center justify-center text-gray-500">
          نمودار توزیع سطوح (در حال توسعه)
        </div>
      </div>

      <!-- آمار تفصیلی -->
      <div class="bg-gray-50 p-6 rounded-lg">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">آمار تفصیلی</h3>
        <div class="space-y-4">
          <div class="flex justify-between items-center">
            <span class="text-gray-600">میانگین امتیاز روزانه:</span>
            <span class="font-semibold text-gray-900">{{ detailedStats.dailyAverage }}</span>
          </div>
          <div class="flex justify-between items-center">
            <span class="text-gray-600">میانگین امتیاز ماهانه:</span>
            <span class="font-semibold text-gray-900">{{ detailedStats.monthlyAverage }}</span>
          </div>
          <div class="flex justify-between items-center">
            <span class="text-gray-600">نرخ تبدیل امتیاز:</span>
            <span class="font-semibold text-gray-900">{{ detailedStats.conversionRate }}%</span>
          </div>
          <div class="flex justify-between items-center">
            <span class="text-gray-600">میانگین زمان ارتقا سطح:</span>
            <span class="font-semibold text-gray-900">{{ detailedStats.avgUpgradeTime }} روز</span>
          </div>
        </div>
      </div>

      <!-- برترین اعضا -->
      <div class="bg-gray-50 p-6 rounded-lg">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">برترین اعضا</h3>
        <div class="space-y-3">
          <div v-for="(member, index) in topMembers" :key="member.id" class="flex items-center justify-between">
            <div class="flex items-center">
              <span class="text-lg font-bold text-gray-400 w-6">{{ index + 1 }}</span>
              <img class="w-8 h-8 rounded-full ml-3" :src="member.avatar" :alt="member.name">
              <span class="text-sm font-medium text-gray-900">{{ member.name }}</span>
            </div>
            <span class="text-sm font-semibold text-gray-900">{{ member.points.toLocaleString() }} امتیاز</span>
          </div>
        </div>
      </div>
    </div>

    <!-- جدول گزارش تفصیلی -->
    <div class="mt-8">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">گزارش تفصیلی</h3>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                تاریخ
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                امتیازات توزیع شده
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                امتیازات استفاده شده
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                اعضای جدید
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                ارتقا سطح
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                پاداش‌های استفاده شده
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="report in detailedReports" :key="report.date">
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ formatDate(report.date) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ report.pointsIssued.toLocaleString() }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ report.pointsUsed.toLocaleString() }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ report.newMembers }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ report.levelUpgrades }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ report.rewardsUsed }}
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// نوع گزارش
const reportType = ref('activity')

// آمار تفصیلی
const detailedStats = ref({
  dailyAverage: 1250,
  monthlyAverage: 37500,
  conversionRate: 62,
  avgUpgradeTime: 45
})

// برترین اعضا
const topMembers = ref([
  {
    id: 1,
    name: 'علی احمدی',
    avatar: '/default-avatar.png',
    points: 15000
  },
  {
    id: 2,
    name: 'فاطمه محمدی',
    avatar: '/default-avatar.png',
    points: 8500
  },
  {
    id: 3,
    name: 'محمد رضایی',
    avatar: '/default-avatar.png',
    points: 7200
  },
  {
    id: 4,
    name: 'زهرا کریمی',
    avatar: '/default-avatar.png',
    points: 6800
  },
  {
    id: 5,
    name: 'حسین نوری',
    avatar: '/default-avatar.png',
    points: 5500
  }
])

// گزارش‌های تفصیلی
const detailedReports = ref([
  {
    date: '2024-01-10',
    pointsIssued: 1250,
    pointsUsed: 850,
    newMembers: 12,
    levelUpgrades: 3,
    rewardsUsed: 8
  },
  {
    date: '2024-01-09',
    pointsIssued: 1180,
    pointsUsed: 920,
    newMembers: 8,
    levelUpgrades: 2,
    rewardsUsed: 6
  },
  {
    date: '2024-01-08',
    pointsIssued: 1320,
    pointsUsed: 780,
    newMembers: 15,
    levelUpgrades: 4,
    rewardsUsed: 10
  },
  {
    date: '2024-01-07',
    pointsIssued: 980,
    pointsUsed: 650,
    newMembers: 6,
    levelUpgrades: 1,
    rewardsUsed: 4
  },
  {
    date: '2024-01-06',
    pointsIssued: 1450,
    pointsUsed: 1100,
    newMembers: 18,
    levelUpgrades: 5,
    rewardsUsed: 12
  }
])

// فرمت تاریخ
function formatDate(dateString: string) {
  const date = new Date(dateString)
  return date.toLocaleDateString('fa-IR')
}
</script> 
