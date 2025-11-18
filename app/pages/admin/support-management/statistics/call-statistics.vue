<template>
  <div class="p-6" dir="rtl">
    <div class="mb-6 bg-white p-6 rounded-lg shadow-md border border-gray-200">
      <h1 class="text-2xl font-bold text-gray-900 mb-2">آمار مکالمات</h1>
      <p class="text-gray-600">تحلیل و گزارش‌گیری از مکالمات پشتیبانی</p>
    </div>

    <!-- آمار کلی -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
      <TemplateCard
        title="کل مکالمات"
        :value="totalCalls"
        variant="blue"
      />
      <TemplateCard
        title="مکالمات موفق"
        :value="successfulCalls"
        variant="green"
      />
      <TemplateCard
        title="مکالمات ناموفق"
        :value="failedCalls"
        variant="red"
      />
      <TemplateCard
        title="میانگین مدت مکالمه"
        :value="avgCallDuration"
        variant="cyan"
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
            <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ شروع</label>
            <input type="date" class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ پایان</label>
            <input type="date" class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
            <select class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm">
              <option value="">همه</option>
              <option value="successful">موفق</option>
              <option value="failed">ناموفق</option>
              <option value="missed">از دست رفته</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">اپراتور</label>
            <select class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm">
              <option value="">همه</option>
              <option value="operator1">علی احمدی</option>
              <option value="operator2">فاطمه محمدی</option>
              <option value="operator3">محمد رضایی</option>
            </select>
          </div>
        </div>
      </div>
    </div>

    <!-- نمودارها -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8 mb-8">
      <!-- نمودار مکالمات در طول زمان -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200">
        <div class="px-6 py-4 border-b border-gray-200">
          <h3 class="text-lg font-semibold text-gray-900">مکالمات در طول زمان</h3>
        </div>
        <div class="p-6">
          <div class="h-64 bg-gray-50 rounded-lg flex items-center justify-center">
            <div class="text-center">
              <svg class="w-16 h-16 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z"></path>
              </svg>
              <p class="text-gray-500">نمودار مکالمات در طول زمان</p>
              <p class="text-sm text-gray-400">آخرین 30 روز</p>
            </div>
          </div>
        </div>
      </div>

      <!-- نمودار توزیع وضعیت مکالمات -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200">
        <div class="px-6 py-4 border-b border-gray-200">
          <h3 class="text-lg font-semibold text-gray-900">توزیع وضعیت مکالمات</h3>
        </div>
        <div class="p-6">
          <div class="h-64 bg-gray-50 rounded-lg flex items-center justify-center">
            <div class="text-center">
              <svg class="w-16 h-16 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 3.055A9.001 9.001 0 1020.945 13H11V3.055z"></path>
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.488 9H15V3.512A9.025 9.025 0 0120.488 9z"></path>
              </svg>
              <p class="text-gray-500">نمودار دایره‌ای وضعیت مکالمات</p>
              <p class="text-sm text-gray-400">موفق، ناموفق، از دست رفته</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- جدول مکالمات -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200 flex justify-between items-center">
        <h3 class="text-lg font-semibold text-gray-900">لیست مکالمات</h3>
        <button class="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600 transition-colors">
          دانلود گزارش
        </button>
      </div>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">شماره تماس</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">اپراتور</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ و زمان</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مدت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="call in calls" :key="call.id">
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{{ call.phoneNumber }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ call.operator }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ call.dateTime }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ call.duration }}</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusClass(call.status)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ getStatusText(call.status) }}
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
                  <button class="text-green-600 hover:text-green-900 transition-colors" title="پخش">
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.828 14.828a4 4 0 01-5.656 0M9 10h1m4 0h1m-6 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup>
import TemplateCard from '@/components/common/TemplateCard.vue'

definePageMeta({ layout: 'admin-main' })

// داده‌های نمونه
const totalCalls = ref(1245)
const successfulCalls = ref(1180)
const failedCalls = ref(45)
const avgCallDuration = ref('8.5 دقیقه')

const calls = ref([
  {
    id: 1,
    phoneNumber: '09123456789',
    operator: 'علی احمدی',
    dateTime: '1402/10/15 - 14:30',
    duration: '12 دقیقه',
    status: 'successful'
  },
  {
    id: 2,
    phoneNumber: '09387654321',
    operator: 'فاطمه محمدی',
    dateTime: '1402/10/15 - 15:45',
    duration: '8 دقیقه',
    status: 'successful'
  },
  {
    id: 3,
    phoneNumber: '09111111111',
    operator: 'محمد رضایی',
    dateTime: '1402/10/15 - 16:20',
    duration: '5 دقیقه',
    status: 'failed'
  },
  {
    id: 4,
    phoneNumber: '09222222222',
    operator: 'سارا کریمی',
    dateTime: '1402/10/15 - 17:10',
    duration: '15 دقیقه',
    status: 'successful'
  },
  {
    id: 5,
    phoneNumber: '09333333333',
    operator: 'حسن نوری',
    dateTime: '1402/10/15 - 18:30',
    duration: '0 دقیقه',
    status: 'missed'
  }
])

function getStatusClass(status) {
  switch (status) {
    case 'successful':
      return 'bg-green-100 text-green-800'
    case 'failed':
      return 'bg-red-100 text-red-800'
    case 'missed':
      return 'bg-yellow-100 text-yellow-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

function getStatusText(status) {
  switch (status) {
    case 'successful':
      return 'موفق'
    case 'failed':
      return 'ناموفق'
    case 'missed':
      return 'از دست رفته'
    default:
      return 'نامشخص'
  }
}
</script> 
