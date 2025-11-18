<template>
  <div class="bg-white rounded-xl shadow-lg overflow-hidden">
    <div class="px-6 py-4 border-b border-gray-200 flex items-center justify-between">
      <h3 class="text-lg font-bold text-gray-900">جزئیات بازدیدها</h3>
      <div class="flex items-center gap-3">
        <div class="text-sm text-gray-500">
          {{ visits.length.toLocaleString('fa-IR') }} رکورد
        </div>
        <button class="text-sm text-blue-600 hover:text-blue-700 font-medium">
          مشاهده همه
        </button>
      </div>
    </div>

    <div class="overflow-x-auto">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              زمان بازدید
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              آدرس IP
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              موقعیت
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              دستگاه
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              مرورگر
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              صفحه
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              منبع
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              مدت زمان
            </th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr
            v-for="visit in visits"
            :key="visit.id"
            class="hover:bg-gray-50 transition-colors duration-200"
          >
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
              {{ visit.timestamp }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm font-mono text-gray-600">
              {{ visit.ip }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="flex items-center gap-2">
                <div class="w-6 h-4 bg-gradient-to-r from-green-500 to-red-500 rounded-sm flex items-center justify-center">
                  <span class="text-white text-xs font-bold">IR</span>
                </div>
                <span class="text-sm text-gray-900">{{ visit.city }}</span>
              </div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                    :class="getDeviceBadgeClass(visit.device)">
                {{ visit.device }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
              {{ visit.browser }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <a href="#" class="text-sm text-blue-600 hover:text-blue-700 hover:underline">
                {{ visit.page }}
              </a>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                    :class="getReferrerBadgeClass(visit.referrer)">
                {{ visit.referrer }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-600">
              {{ visit.duration }}
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Pagination -->
    <div class="px-6 py-4 border-t border-gray-200 flex items-center justify-between">
      <div class="text-sm text-gray-500">
        نمایش 1 تا {{ visits.length }} از {{ visits.length }} رکورد
      </div>
      <div class="flex items-center gap-2">
        <button class="px-3 py-1 text-sm text-gray-600 border border-gray-300 rounded hover:bg-gray-50 disabled:opacity-50" disabled>
          قبلی
        </button>
        <span class="px-3 py-1 text-sm bg-blue-600 text-white rounded">1</span>
        <button class="px-3 py-1 text-sm text-gray-600 border border-gray-300 rounded hover:bg-gray-50">
          بعدی
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
interface Visit {
  id: number
  timestamp: string
  ip: string
  city: string
  device: string
  browser: string
  page: string
  referrer: string
  duration: string
}

const props = defineProps<{
  visits: Visit[]
}>()

const getDeviceBadgeClass = (device: string) => {
  const classes = {
    'موبایل': 'bg-blue-100 text-blue-800',
    'دسکتاپ': 'bg-green-100 text-green-800',
    'تبلت': 'bg-purple-100 text-purple-800'
  }
  return classes[device] || 'bg-gray-100 text-gray-800'
}

const getReferrerBadgeClass = (referrer: string) => {
  const classes = {
    'Google': 'bg-red-100 text-red-800',
    'مستقیم': 'bg-gray-100 text-gray-800',
    'Instagram': 'bg-pink-100 text-pink-800',
    'Telegram': 'bg-blue-100 text-blue-800',
    'Facebook': 'bg-indigo-100 text-indigo-800',
    'Twitter': 'bg-cyan-100 text-cyan-800'
  }
  return classes[referrer] || 'bg-gray-100 text-gray-800'
}
</script> 
