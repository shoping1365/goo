<template>
  <div class="p-6" dir="rtl">
    <!-- هدر بنفش -->
    <div class="bg-gradient-to-r from-purple-600 to-purple-700 text-white p-6 rounded-lg mb-6">
      <h1 class="text-2xl font-bold mb-2">آمار اپراتورها</h1>
      <p class="text-purple-100">با انتخاب دوره زمانی دلخواه، می‌توانید کل و میانگین زمان فعالیت، تعداد مکالمات قبول شده و جزئیات آماری هر اپراتور را بررسی کنید.</p>
    </div>

    <!-- انتخاب بازه زمانی و خروجی -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6 mb-8">
      <div class="flex justify-between items-center">
        <h3 class="text-2xl font-bold text-gray-900 mb-4">آمار اپراتورها</h3>
        
        <div class="flex items-center space-x-4 space-x-reverse">
          <!-- بازه زمانی -->
          <div class="bg-gray-50 rounded-lg p-6 flex items-center space-x-2 space-x-reverse text-gray-600">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
            </svg>
            <span class="text-sm">۱۴۰۴/۰۵/۰۱ - ۱۴۰۴/۰۴/۲۶</span>
          </div>
          
          <!-- دکمه خروجی CSV -->
          <button class="bg-blue-500 text-white px-4 py-2 rounded-lg hover:bg-blue-600 transition-colors flex items-center space-x-2 space-x-reverse">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
            </svg>
            <span>خروجی csv</span>
          </button>
        </div>
      </div>
    </div>

    <!-- آمار کلی -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
      <TemplateCard
        title="کل اپراتورها"
        :value="totalOperators"
        variant="blue"
      />
      <TemplateCard
        title="اپراتورهای آنلاین"
        :value="onlineOperators"
        variant="green"
      />
      <TemplateCard
        title="میانگین رضایت"
        :value="avgSatisfaction"
        variant="cyan"
      />
      <TemplateCard
        title="میانگین زمان پاسخ"
        :value="avgResponseTime"
        variant="orange"
      />
    </div>

    <!-- فیلترها -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 mb-8">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">فیلترها و جستجو</h3>
      </div>
      <div class="p-6">
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">دوره زمانی</label>
            <select class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm">
              <option value="today">امروز</option>
              <option value="week">هفته گذشته</option>
              <option value="month">ماه گذشته</option>
              <option value="quarter">سه ماهه</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">سطح اپراتور</label>
            <select class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm">
              <option value="">همه</option>
              <option value="senior">ارشد</option>
              <option value="junior">جونیور</option>
              <option value="trainee">کارآموز</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
            <select class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm">
              <option value="">همه</option>
              <option value="online">آنلاین</option>
              <option value="busy">مشغول</option>
              <option value="offline">آفلاین</option>
            </select>
          </div>
        </div>
      </div>
    </div>

    <!-- نمودارها -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8 mb-8">
      <!-- نمودار عملکرد اپراتورها -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200">
        <div class="px-6 py-4 border-b border-gray-200">
          <h3 class="text-lg font-semibold text-gray-900">عملکرد اپراتورها</h3>
        </div>
        <div class="p-6">
          <div class="h-64 bg-gray-50 rounded-lg flex items-center justify-center">
            <div class="text-center">
              <svg class="w-16 h-16 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
              </svg>
              <p class="text-gray-500">نمودار عملکرد اپراتورها</p>
              <p class="text-sm text-gray-400">بر اساس تعداد تیکت‌های حل شده</p>
            </div>
          </div>
        </div>
      </div>

      <!-- نمودار رضایت اپراتورها -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200">
        <div class="px-6 py-4 border-b border-gray-200">
          <h3 class="text-lg font-semibold text-gray-900">رضایت اپراتورها</h3>
        </div>
        <div class="p-6">
          <div class="h-64 bg-gray-50 rounded-lg flex items-center justify-center">
            <div class="text-center">
              <svg class="w-16 h-16 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"></path>
              </svg>
              <p class="text-gray-500">نمودار رضایت اپراتورها</p>
              <p class="text-sm text-gray-400">بر اساس امتیاز مشتریان</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- جدول آمار اپراتورها -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200">
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">اپراتور</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کل زمان فعالیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">میانگین زمان فعالیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مکالمات قبول شده</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">میانگین زمان پاسخگویی</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">میانگین زمان بستن مکالمه</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">درصد رضایت مندی</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">امتیاز از ۱۰۰</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <!-- ردیف میانگین -->
            <tr class="bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">میانگین</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">-</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">-</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">۵۱ مکالمه</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">-</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">-</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">-</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">-</td>
            </tr>
            
            <!-- ردیف اپراتور آذری -->
            <tr>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="">
                    <div class="h-10 w-10 rounded-full bg-orange-400 flex items-center justify-center">
                      <span class="text-white text-sm">🐮</span>
                    </div>
                  </div>
                  <div class="mr-4">
                    <div class="text-sm font-medium text-gray-900">۱. آذری</div>
                  </div>
                  <div class="flex-shrink-0">
                    <div class="w-5 h-5 bg-gray-300 rounded-full flex items-center justify-center">
                      <span class="text-gray-600 text-xs">i</span>
                    </div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">۴:۴۸</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">۰:۲۶</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">۵۱ مکالمه</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">۰:۰</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">۰:۰</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">-</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">-</td>
            </tr>
            
            <!-- ردیف‌های اضافی برای نمونه -->
            <tr v-for="operator in operatorStats" :key="operator.id">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="">
                    <div class="h-10 w-10 rounded-full bg-gray-300 flex items-center justify-center">
                      <span class="text-sm font-medium text-gray-700">{{ operator.name.charAt(0) }}</span>
                    </div>
                  </div>
                  <div class="mr-4">
                    <div class="text-sm font-medium text-gray-900">{{ operator.name }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ operator.totalActivityTime || '۳:۲۵' }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ operator.avgActivityTime || '۰:۱۸' }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ operator.acceptedCalls || '۳۸ مکالمه' }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ operator.avgResponseTime || '۰:۱۵' }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ operator.avgCallClosingTime || '۰:۴۵' }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ operator.satisfactionPercentage || '۹۲%' }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ operator.scoreOutOf100 || '۸۵' }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup>
import TemplateCard from '@/components/common/TemplateCard.vue'

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// داده‌های نمونه
const totalOperators = ref(8)
const onlineOperators = ref(6)
const avgSatisfaction = ref('4.7/5')
const avgResponseTime = ref('2.3 دقیقه')

const operatorStats = ref([
  {
    id: 1,
    name: 'علی احمدی',
    level: 'ارشد',
    status: 'online',
    resolvedTickets: 45,
    avgResponseTime: '2.1 دقیقه',
    satisfaction: 4.8,
    workHours: '8.5 ساعت'
  },
  {
    id: 2,
    name: 'فاطمه محمدی',
    level: 'ارشد',
    status: 'busy',
    resolvedTickets: 38,
    avgResponseTime: '2.8 دقیقه',
    satisfaction: 4.6,
    workHours: '7.8 ساعت'
  },
  {
    id: 3,
    name: 'محمد رضایی',
    level: 'جونیور',
    status: 'online',
    resolvedTickets: 32,
    avgResponseTime: '3.2 دقیقه',
    satisfaction: 4.4,
    workHours: '8.2 ساعت'
  },
  {
    id: 4,
    name: 'سارا کریمی',
    level: 'کارآموز',
    status: 'offline',
    resolvedTickets: 28,
    avgResponseTime: '4.1 دقیقه',
    satisfaction: 4.2,
    workHours: '6.5 ساعت'
  },
  {
    id: 5,
    name: 'حسن نوری',
    level: 'ارشد',
    status: 'online',
    resolvedTickets: 41,
    avgResponseTime: '1.9 دقیقه',
    satisfaction: 4.9,
    workHours: '8.0 ساعت'
  }
])

// function getStatusClass(status) {
//   switch (status) {
//     case 'online':
//       return 'bg-green-100 text-green-800'
//     case 'busy':
//       return 'bg-yellow-100 text-yellow-800'
//     case 'offline':
//       return 'bg-red-100 text-red-800'
//     default:
//       return 'bg-gray-100 text-gray-800'
//   }
// }

// function getStatusText(status) {
//   switch (status) {
//     case 'online':
//       return 'آنلاین'
//     case 'busy':
//       return 'مشغول'
//     case 'offline':
//       return 'آفلاین'
//     default:
//       return 'نامشخص'
//   }
// }
</script> 
