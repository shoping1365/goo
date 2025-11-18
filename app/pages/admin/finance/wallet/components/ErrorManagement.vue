<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="bg-gradient-to-r from-red-50 to-pink-50 rounded-lg p-6">
      <h2 class="text-2xl font-bold text-gray-900 mb-2">⚠️ مدیریت خطاها</h2>
      <p class="text-gray-600">مدیریت و رفع خطاهای سیستم کیف پول</p>
    </div>

    <!-- آمار کلی خطاها -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <!-- خطاهای امروز -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">خطاهای امروز</h3>
          <span class="text-xs bg-red-100 text-red-800 rounded-full px-3 py-1">هشدار</span>
        </div>
        <div class="text-3xl font-bold text-red-600 mb-2">{{ errorStats.todayErrors }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-red-500">+{{ errorStats.errorGrowth }}%</span>
          <span class="mx-2">از دیروز</span>
        </div>
      </div>

      <!-- خطاهای حل شده -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">خطاهای حل شده</h3>
          <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">موفق</span>
        </div>
        <div class="text-3xl font-bold text-green-600 mb-2">{{ errorStats.resolvedErrors }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-green-500">{{ errorStats.resolutionRate }}%</span>
          <span class="mx-2">نرخ حل</span>
        </div>
      </div>

      <!-- خطاهای در انتظار -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">در انتظار حل</h3>
          <span class="text-xs bg-yellow-100 text-yellow-800 rounded-full px-3 py-1">انتظار</span>
        </div>
        <div class="text-3xl font-bold text-yellow-600 mb-2">{{ errorStats.pendingErrors }}</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-yellow-500">{{ errorStats.pendingPercentage }}%</span>
          <span class="mx-2">از کل خطاها</span>
        </div>
      </div>

      <!-- زمان متوسط حل -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">زمان حل متوسط</h3>
          <span class="text-xs bg-blue-100 text-blue-800 rounded-full px-3 py-1">بهبود</span>
        </div>
        <div class="text-3xl font-bold text-blue-600 mb-2">{{ errorStats.avgResolutionTime }}m</div>
        <div class="flex items-center text-sm text-gray-600">
          <span class="text-blue-500">-{{ errorStats.timeImprovement }}%</span>
          <span class="mx-2">از هفته قبل</span>
        </div>
      </div>
    </div>

    <!-- فیلترهای خطا -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">فیلترهای خطا</h3>
      
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نوع خطا</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option value="all">همه خطاها</option>
            <option value="transaction">تراکنش ناموفق</option>
            <option value="gateway">خطای درگاه پرداخت</option>
            <option value="network">مشکل شبکه</option>
            <option value="system">خطای سیستمی</option>
            <option value="security">خطای امنیتی</option>
          </select>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">سطح اهمیت</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option value="all">همه سطوح</option>
            <option value="critical">بحرانی</option>
            <option value="high">بالا</option>
            <option value="medium">متوسط</option>
            <option value="low">پایین</option>
          </select>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option value="all">همه وضعیت‌ها</option>
            <option value="open">باز</option>
            <option value="in_progress">در حال حل</option>
            <option value="resolved">حل شده</option>
            <option value="closed">بسته</option>
          </select>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">بازه زمانی</label>
          <select class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option value="today">امروز</option>
            <option value="week">هفته جاری</option>
            <option value="month" selected>ماه جاری</option>
            <option value="quarter">سه ماهه</option>
          </select>
        </div>
      </div>
      
      <div class="mt-4 flex space-x-2 space-x-reverse">
        <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
          اعمال فیلتر
        </button>
        <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
          پاک کردن
        </button>
        <button class="px-4 py-2 bg-green-600 text-white text-sm rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500">
          حل خودکار
        </button>
      </div>
    </div>

    <!-- جدول خطاها -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-xl font-semibold text-gray-900">لیست خطاها</h3>
        <div class="flex space-x-2 space-x-reverse">
          <button class="px-4 py-2 bg-red-600 text-white text-sm rounded-lg hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500">
            حل دسته‌ای
          </button>
          <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
            خروجی
          </button>
        </div>
      </div>
      
      <div class="overflow-x-auto">
        <table class="w-full text-sm text-right">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-4 py-3 text-gray-700 font-medium">
                <input type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
              </th>
              <th class="px-4 py-3 text-gray-700 font-medium">شناسه خطا</th>
              <th class="px-4 py-3 text-gray-700 font-medium">نوع خطا</th>
              <th class="px-4 py-3 text-gray-700 font-medium">توضیحات</th>
              <th class="px-4 py-3 text-gray-700 font-medium">سطح اهمیت</th>
              <th class="px-4 py-3 text-gray-700 font-medium">وضعیت</th>
              <th class="px-4 py-3 text-gray-700 font-medium">تاریخ</th>
              <th class="px-4 py-3 text-gray-700 font-medium">عملیات</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="error in errors" :key="error.id" class="hover:bg-gray-50">
              <td class="px-4 py-3">
                <input type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
              </td>
              <td class="px-4 py-3 text-gray-900 font-medium">{{ error.id }}</td>
              <td class="px-4 py-3">
                <span :class="getErrorTypeClass(error.type)" class="px-2 py-1 text-xs rounded-full">
                  {{ error.type }}
                </span>
              </td>
              <td class="px-4 py-3 text-gray-600">{{ error.description }}</td>
              <td class="px-4 py-3">
                <span :class="getSeverityClass(error.severity)" class="px-2 py-1 text-xs rounded-full">
                  {{ error.severity }}
                </span>
              </td>
              <td class="px-4 py-3">
                <span :class="getStatusClass(error.status)" class="px-2 py-1 text-xs rounded-full">
                  {{ error.status }}
                </span>
              </td>
              <td class="px-4 py-3 text-gray-600">{{ error.date }}</td>
              <td class="px-4 py-3">
                <div class="flex space-x-2 space-x-reverse">
                  <button class="text-blue-600 hover:text-blue-800 text-sm">جزئیات</button>
                  <button v-if="error.status === 'باز'" class="text-green-600 hover:text-green-800 text-sm">حل</button>
                  <button v-if="error.status === 'باز'" class="text-yellow-600 hover:text-yellow-800 text-sm">در حال حل</button>
                  <button class="text-red-600 hover:text-red-800 text-sm">حذف</button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      
      <!-- صفحه‌بندی -->
      <div class="mt-6 flex items-center justify-between">
        <div class="text-sm text-gray-600">
          نمایش {{ paginationInfo.start }} تا {{ paginationInfo.end }} از {{ paginationInfo.total }} خطا
        </div>
        <div class="flex space-x-2 space-x-reverse">
          <button class="px-3 py-1 text-sm bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200">قبلی</button>
          <button class="px-3 py-1 text-sm bg-blue-100 text-blue-800 rounded-lg">1</button>
          <button class="px-3 py-1 text-sm bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200">2</button>
          <button class="px-3 py-1 text-sm bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200">3</button>
          <button class="px-3 py-1 text-sm bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200">بعدی</button>
        </div>
      </div>
    </div>

    <!-- راه‌حل‌های خودکار -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- تنظیمات راه‌حل خودکار -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات راه‌حل خودکار</h3>
        
        <div class="space-y-4">
          <div class="flex items-center justify-between">
            <div>
              <div class="text-sm font-medium text-gray-900">حل خودکار خطاهای شبکه</div>
              <div class="text-xs text-gray-500">تلاش مجدد خودکار برای خطاهای شبکه</div>
            </div>
            <input type="checkbox" :checked="autoFixSettings.networkErrors" 
                   class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
          </div>
          
          <div class="flex items-center justify-between">
            <div>
              <div class="text-sm font-medium text-gray-900">بازپرداخت خودکار</div>
              <div class="text-xs text-gray-500">بازپرداخت خودکار تراکنش‌های ناموفق</div>
            </div>
            <input type="checkbox" :checked="autoFixSettings.autoRefund" 
                   class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
          </div>
          
          <div class="flex items-center justify-between">
            <div>
              <div class="text-sm font-medium text-gray-900">تأیید خودکار</div>
              <div class="text-xs text-gray-500">تأیید خودکار تراکنش‌های مشکوک</div>
            </div>
            <input type="checkbox" :checked="autoFixSettings.autoApprove" 
                   class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
          </div>
          
          <div class="flex items-center justify-between">
            <div>
              <div class="text-sm font-medium text-gray-900">اعلان خودکار</div>
              <div class="text-xs text-gray-500">ارسال اعلان برای خطاهای بحرانی</div>
            </div>
            <input type="checkbox" :checked="autoFixSettings.autoNotification" 
                   class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
          </div>
        </div>
        
        <div class="mt-4">
          <button class="w-full px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
            ذخیره تنظیمات
          </button>
        </div>
      </div>

      <!-- آمار راه‌حل‌های خودکار -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">آمار راه‌حل‌های خودکار</h3>
        
        <div class="space-y-4">
          <div class="flex items-center justify-between p-3 bg-green-50 rounded-lg">
            <div>
              <div class="text-sm font-medium text-green-900">خطاهای حل شده خودکار</div>
              <div class="text-xs text-green-600">امروز</div>
            </div>
            <div class="text-right">
              <div class="text-lg font-bold text-green-600">{{ autoFixStats.autoResolved }}</div>
              <div class="text-xs text-green-500">{{ autoFixStats.autoResolveRate }}%</div>
            </div>
          </div>
          
          <div class="flex items-center justify-between p-3 bg-blue-50 rounded-lg">
            <div>
              <div class="text-sm font-medium text-blue-900">بازپرداخت‌های خودکار</div>
              <div class="text-xs text-blue-600">امروز</div>
            </div>
            <div class="text-right">
              <div class="text-lg font-bold text-blue-600">{{ autoFixStats.autoRefunds }}</div>
              <div class="text-xs text-blue-500">{{ formatCurrency(autoFixStats.autoRefundAmount) }}</div>
            </div>
          </div>
          
          <div class="flex items-center justify-between p-3 bg-purple-50 rounded-lg">
            <div>
              <div class="text-sm font-medium text-purple-900">تأییدهای خودکار</div>
              <div class="text-xs text-purple-600">امروز</div>
            </div>
            <div class="text-right">
              <div class="text-lg font-bold text-purple-600">{{ autoFixStats.autoApprovals }}</div>
              <div class="text-xs text-purple-500">{{ autoFixStats.autoApprovalRate }}%</div>
            </div>
          </div>
          
          <div class="flex items-center justify-between p-3 bg-orange-50 rounded-lg">
            <div>
              <div class="text-sm font-medium text-orange-900">اعلان‌های ارسال شده</div>
              <div class="text-xs text-orange-600">امروز</div>
            </div>
            <div class="text-right">
              <div class="text-lg font-bold text-orange-600">{{ autoFixStats.notificationsSent }}</div>
              <div class="text-xs text-orange-500">{{ autoFixStats.notificationRate }}%</div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- نمودار روند خطاها -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-xl font-semibold text-gray-900">روند خطاها (7 روز گذشته)</h3>
        <div class="flex space-x-2 space-x-reverse">
          <button class="px-3 py-1 text-sm bg-blue-100 text-blue-800 rounded-lg hover:bg-blue-200">روزانه</button>
          <button class="px-3 py-1 text-sm bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200">هفتگی</button>
          <button class="px-3 py-1 text-sm bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200">ماهانه</button>
        </div>
      </div>
      
      <div class="h-64 flex items-end space-x-2 space-x-reverse overflow-x-auto">
        <div v-for="(day, index) in errorTrend" :key="index" class="flex-shrink-0 flex flex-col items-center min-w-12">
          <div class="w-full bg-gray-200 rounded-t relative"
               :style="{ height: getChartHeight(day.errors) + 'px' }">
            <div class="w-full bg-gradient-to-t from-red-500 to-pink-500 rounded-t transition-all duration-300 absolute bottom-0"
                 :style="{ height: getChartHeight(day.errors) + 'px' }"></div>
          </div>
          <span class="text-xs text-gray-500 mt-1 text-center">{{ day.date }}</span>
          <span class="text-xs text-gray-400 mt-1 text-center">{{ day.errors }} خطا</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// آمار کلی خطاها
const errorStats = {
  todayErrors: 45,
  errorGrowth: 12.5,
  resolvedErrors: 38,
  resolutionRate: 84.4,
  pendingErrors: 7,
  pendingPercentage: 15.6,
  avgResolutionTime: 25,
  timeImprovement: 8.3
}

// خطاها
const errors = [
  {
    id: 'ERR-001',
    type: 'تراکنش ناموفق',
    description: 'خطا در اتصال به درگاه پرداخت',
    severity: 'بالا',
    status: 'باز',
    date: '1402/10/15 14:30'
  },
  {
    id: 'ERR-002',
    type: 'خطای درگاه پرداخت',
    description: 'پاسخ نامعتبر از درگاه پرداخت',
    severity: 'بحرانی',
    status: 'در حال حل',
    date: '1402/10/15 13:45'
  },
  {
    id: 'ERR-003',
    type: 'مشکل شبکه',
    description: 'اتصال شبکه قطع شد',
    severity: 'متوسط',
    status: 'حل شده',
    date: '1402/10/15 12:20'
  },
  {
    id: 'ERR-004',
    type: 'خطای سیستمی',
    description: 'خطا در پردازش تراکنش',
    severity: 'بالا',
    status: 'باز',
    date: '1402/10/15 11:15'
  },
  {
    id: 'ERR-005',
    type: 'خطای امنیتی',
    description: 'تلاش غیرمجاز برای دسترسی',
    severity: 'بحرانی',
    status: 'بسته',
    date: '1402/10/15 10:30'
  }
]

// تنظیمات راه‌حل خودکار
const autoFixSettings = {
  networkErrors: true,
  autoRefund: true,
  autoApprove: false,
  autoNotification: true
}

// آمار راه‌حل‌های خودکار
const autoFixStats = {
  autoResolved: 32,
  autoResolveRate: 71.1,
  autoRefunds: 15,
  autoRefundAmount: 85000000,
  autoApprovals: 8,
  autoApprovalRate: 17.8,
  notificationsSent: 45,
  notificationRate: 100
}

// روند خطاها (7 روز)
const errorTrend = [
  { errors: 38, date: 'شنبه' },
  { errors: 42, date: 'یکشنبه' },
  { errors: 35, date: 'دوشنبه' },
  { errors: 48, date: 'سه‌شنبه' },
  { errors: 41, date: 'چهارشنبه' },
  { errors: 39, date: 'پنج‌شنبه' },
  { errors: 45, date: 'جمعه' }
]

// اطلاعات صفحه‌بندی
const paginationInfo = {
  start: 1,
  end: 10,
  total: 45
}

// تابع فرمت کردن ارز
const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

// تابع کلاس نوع خطا
const getErrorTypeClass = (type: string) => {
  const classes = {
    'تراکنش ناموفق': 'bg-red-100 text-red-800',
    'خطای درگاه پرداخت': 'bg-orange-100 text-orange-800',
    'مشکل شبکه': 'bg-yellow-100 text-yellow-800',
    'خطای سیستمی': 'bg-purple-100 text-purple-800',
    'خطای امنیتی': 'bg-pink-100 text-pink-800'
  }
  return classes[type] || 'bg-gray-100 text-gray-800'
}

// تابع کلاس سطح اهمیت
const getSeverityClass = (severity: string) => {
  const classes = {
    'بحرانی': 'bg-red-100 text-red-800',
    'بالا': 'bg-orange-100 text-orange-800',
    'متوسط': 'bg-yellow-100 text-yellow-800',
    'پایین': 'bg-green-100 text-green-800'
  }
  return classes[severity] || 'bg-gray-100 text-gray-800'
}

// تابع کلاس وضعیت
const getStatusClass = (status: string) => {
  const classes = {
    'باز': 'bg-red-100 text-red-800',
    'در حال حل': 'bg-yellow-100 text-yellow-800',
    'حل شده': 'bg-green-100 text-green-800',
    'بسته': 'bg-gray-100 text-gray-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

// تابع محاسبه ارتفاع مناسب برای نمودار
const getChartHeight = (errors: number) => {
  const maxErrors = Math.max(...errorTrend.map(item => item.errors));
  const minErrors = Math.min(...errorTrend.map(item => item.errors));
  const range = maxErrors - minErrors;
  const height = 200; // حداکثر ارتفاع نمودار

  if (range === 0) return height;

  const percentage = ((errors - minErrors) / range) * 100;
  return (percentage / 100) * height;
}
</script>

<!--
  مستندسازی:
  این کامپوننت شامل مدیریت خطاهای کیف پول است که شامل:
  1. آمار کلی خطاها: خطاهای امروز، حل شده، در انتظار، زمان حل
  2. فیلترهای خطا: نوع، سطح اهمیت، وضعیت، بازه زمانی
  3. جدول خطاها: نمایش و مدیریت خطاها
  4. راه‌حل‌های خودکار: تنظیمات و آمار
  5. نمودار روند خطاها: نمایش روند خطاها در زمان
  
  تمام بخش‌ها به صورت ریسپانسیو و با طراحی مدرن ارائه شده‌اند.
--> 
