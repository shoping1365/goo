<template>
  <div class="p-6" dir="rtl">
    <div class="mb-6 bg-white p-6 rounded-lg shadow-md border border-gray-200">
      <h1 class="text-2xl font-bold text-gray-900 mb-2">آرشیو مکالمات</h1>
      <p class="text-gray-600">مدیریت و جستجو در آرشیو مکالمات پشتیبانی</p>
    </div>

    <!-- آمار کلی -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
      <TemplateCard
        title="کل مکالمات آرشیو شده"
        :value="totalArchivedCalls"
        variant="blue"
      />
      <TemplateCard
        title="حجم فایل‌ها"
        :value="totalFileSize"
        variant="cyan"
      />
      <TemplateCard
        title="قدیمی‌ترین مکالمه"
        :value="oldestCall"
        variant="orange"
      />
      <TemplateCard
        title="جدیدترین مکالمه"
        :value="newestCall"
        variant="green"
      />
    </div>

    <!-- فیلترها -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 mb-8">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">فیلترها و جستجو</h3>
      </div>
      <div class="p-6">
        <div class="grid grid-cols-1 md:grid-cols-5 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ شروع</label>
            <input type="date" class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ پایان</label>
            <input type="date" class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">شماره تماس</label>
            <input type="text" placeholder="جستجو..." class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm">
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
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
            <select class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm">
              <option value="">همه</option>
              <option value="successful">موفق</option>
              <option value="failed">ناموفق</option>
              <option value="missed">از دست رفته</option>
            </select>
          </div>
        </div>
      </div>
    </div>

    <!-- عملیات دسته‌ای -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 mb-8">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">عملیات دسته‌ای</h3>
      </div>
      <div class="p-6">
        <div class="flex flex-wrap gap-6">
          <button class="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600 transition-colors">
            دانلود انتخاب شده
          </button>
          <button class="bg-green-500 text-white px-4 py-2 rounded-md hover:bg-green-600 transition-colors">
            انتقال به آرشیو
          </button>
          <button class="bg-red-500 text-white px-4 py-2 rounded-md hover:bg-red-600 transition-colors">
            حذف انتخاب شده
          </button>
          <button class="bg-purple-500 text-white px-4 py-2 rounded-md hover:bg-purple-600 transition-colors">
            پخش دسته‌ای
          </button>
        </div>
      </div>
    </div>

    <!-- جدول آرشیو مکالمات -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200 flex justify-between items-center">
        <h3 class="text-lg font-semibold text-gray-900">لیست مکالمات آرشیو شده</h3>
        <div class="flex space-x-2 space-x-reverse">
          <button class="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600 transition-colors">
            دانلود همه
          </button>
          <button class="bg-gray-500 text-white px-4 py-2 rounded-md hover:bg-gray-600 transition-colors">
            پاک کردن آرشیو
          </button>
        </div>
      </div>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                <input type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">شماره تماس</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">اپراتور</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ و زمان</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مدت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">حجم فایل</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="call in archivedCalls" :key="call.id">
              <td class="px-6 py-4 whitespace-nowrap">
                <input type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{{ call.phoneNumber }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ call.operator }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ call.dateTime }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ call.duration }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ call.fileSize }}</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusClass(call.status)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ getStatusText(call.status) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <div class="flex space-x-2 space-x-reverse">
                  <button class="text-blue-600 hover:text-blue-900 transition-colors" title="پخش">
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.828 14.828a4 4 0 01-5.656 0M9 10h1m4 0h1m-6 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                    </svg>
                  </button>
                  <button class="text-green-600 hover:text-green-900 transition-colors" title="دانلود">
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                    </svg>
                  </button>
                  <button class="text-purple-600 hover:text-purple-900 transition-colors" title="مشاهده جزئیات">
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                    </svg>
                  </button>
                  <button class="text-red-600 hover:text-red-900 transition-colors" title="حذف">
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- صفحه‌بندی -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 mt-8">
      <div class="px-6 py-4 flex items-center justify-between">
        <div class="text-sm text-gray-700">
          نمایش <span class="font-medium">1</span> تا <span class="font-medium">10</span> از <span class="font-medium">97</span> نتیجه
        </div>
        <div class="flex space-x-2 space-x-reverse">
          <button class="px-3 py-2 text-sm font-medium text-gray-500 bg-white border border-gray-300 rounded-md hover:bg-gray-50">
            قبلی
          </button>
          <button class="px-3 py-2 text-sm font-medium text-white bg-blue-600 border border-blue-600 rounded-md">
            1
          </button>
          <button class="px-3 py-2 text-sm font-medium text-gray-500 bg-white border border-gray-300 rounded-md hover:bg-gray-50">
            2
          </button>
          <button class="px-3 py-2 text-sm font-medium text-gray-500 bg-white border border-gray-300 rounded-md hover:bg-gray-50">
            3
          </button>
          <button class="px-3 py-2 text-sm font-medium text-gray-500 bg-white border border-gray-300 rounded-md hover:bg-gray-50">
            بعدی
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import TemplateCard from '@/components/common/TemplateCard.vue'

definePageMeta({ layout: 'admin-main' })

// داده‌های نمونه
const totalArchivedCalls = ref(1245)
const totalFileSize = ref('2.3 GB')
const oldestCall = ref('1401/01/01')
const newestCall = ref('1402/10/15')

const archivedCalls = ref([
  {
    id: 1,
    phoneNumber: '09123456789',
    operator: 'علی احمدی',
    dateTime: '1402/10/15 - 14:30',
    duration: '12 دقیقه',
    fileSize: '8.5 MB',
    status: 'successful'
  },
  {
    id: 2,
    phoneNumber: '09387654321',
    operator: 'فاطمه محمدی',
    dateTime: '1402/10/15 - 15:45',
    duration: '8 دقیقه',
    fileSize: '5.2 MB',
    status: 'successful'
  },
  {
    id: 3,
    phoneNumber: '09111111111',
    operator: 'محمد رضایی',
    dateTime: '1402/10/15 - 16:20',
    duration: '5 دقیقه',
    fileSize: '3.8 MB',
    status: 'failed'
  },
  {
    id: 4,
    phoneNumber: '09222222222',
    operator: 'سارا کریمی',
    dateTime: '1402/10/15 - 17:10',
    duration: '15 دقیقه',
    fileSize: '12.1 MB',
    status: 'successful'
  },
  {
    id: 5,
    phoneNumber: '09333333333',
    operator: 'حسن نوری',
    dateTime: '1402/10/15 - 18:30',
    duration: '0 دقیقه',
    fileSize: '0 MB',
    status: 'missed'
  },
  {
    id: 6,
    phoneNumber: '09444444444',
    operator: 'علی احمدی',
    dateTime: '1402/10/14 - 09:15',
    duration: '20 دقیقه',
    fileSize: '15.3 MB',
    status: 'successful'
  },
  {
    id: 7,
    phoneNumber: '09555555555',
    operator: 'فاطمه محمدی',
    dateTime: '1402/10/14 - 11:30',
    duration: '6 دقیقه',
    fileSize: '4.2 MB',
    status: 'successful'
  },
  {
    id: 8,
    phoneNumber: '09666666666',
    operator: 'محمد رضایی',
    dateTime: '1402/10/14 - 13:45',
    duration: '10 دقیقه',
    fileSize: '7.1 MB',
    status: 'successful'
  },
  {
    id: 9,
    phoneNumber: '09777777777',
    operator: 'سارا کریمی',
    dateTime: '1402/10/14 - 16:20',
    duration: '3 دقیقه',
    fileSize: '2.1 MB',
    status: 'failed'
  },
  {
    id: 10,
    phoneNumber: '09888888888',
    operator: 'حسن نوری',
    dateTime: '1402/10/14 - 18:55',
    duration: '18 دقیقه',
    fileSize: '13.7 MB',
    status: 'successful'
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
