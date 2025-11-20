<template>
  <div class="min-h-screen" dir="rtl">
    <!-- Header -->
    <div class="bg-white border-b border-gray-200 px-6 py-4">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-2xl font-bold text-indigo-700">تحلیل رفتار کاربران</h1>
          <p class="text-gray-500 mt-1">تحلیل جامع تعامل، مسیر کاربری و عملکرد صفحات</p>
        </div>
        <div class="flex gap-2">
          <button class="px-4 py-2 bg-indigo-500 text-white rounded-lg shadow hover:bg-indigo-600 transition">
            گزارش تفصیلی
          </button>
          <button class="px-4 py-2 bg-green-500 text-white rounded-lg shadow hover:bg-green-600 transition">
            خروجی اکسل
          </button>
          <button class="px-4 py-2 bg-blue-500 text-white rounded-lg shadow hover:bg-blue-600 transition">
            تحلیل مسیر
          </button>
        </div>
      </div>
    </div>

    <!-- Filters -->
    <div class="px-6 py-4 bg-white border-b">
      <div class="flex flex-wrap gap-6">
        <select class="px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-indigo-500">
          <option>بازه زمانی</option>
          <option>امروز</option>
          <option>هفته گذشته</option>
          <option>ماه گذشته</option>
          <option>۳ ماه گذشته</option>
        </select>
        <select class="px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-indigo-500">
          <option>نوع کاربر</option>
          <option>همه</option>
          <option>جدید</option>
          <option>بازگشتی</option>
          <option>ثبت‌نام شده</option>
        </select>
        <select class="px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-indigo-500">
          <option>دستگاه</option>
          <option>همه</option>
          <option>دسکتاپ</option>
          <option>موبایل</option>
          <option>تبلت</option>
        </select>
        <select class="px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-indigo-500">
          <option>مرورگر</option>
          <option>همه</option>
          <option>Chrome</option>
          <option>Firefox</option>
          <option>Safari</option>
        </select>
        <input type="date" class="px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-indigo-500" value="2024-01-01">
        <input type="date" class="px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-indigo-500" value="2024-12-31">
        <button class="px-4 py-2 bg-indigo-500 text-white rounded-lg text-sm hover:bg-indigo-600 transition">اعمال فیلتر</button>
        <button class="px-4 py-2 bg-gray-100 text-gray-700 rounded-lg text-sm hover:bg-gray-200 transition">پاک کردن</button>
      </div>
    </div>
    
    <!-- Summary Cards -->
    <div class="px-6 py-6">
      <div class="grid grid-cols-2 md:grid-cols-4 xl:grid-cols-8 gap-6">
        <div class="bg-gradient-to-br from-indigo-50 to-indigo-100 rounded-lg shadow-sm p-6 border border-indigo-200">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-xs text-indigo-600">کل جلسات</p>
              <p class="text-xl font-bold text-indigo-700">{{ stats.totalSessions.toLocaleString('fa-IR') }}</p>
            </div>
            <div class="bg-indigo-200 p-2 rounded-full">
              <svg class="w-5 h-5 text-indigo-700" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M10 2a1 1 0 00-1 1v1a1 1 0 102 0V3a1 1 0 00-1-1zM4 4a1 1 0 00-1 1v1a1 1 0 102 0V5a1 1 0 00-1-1zm12 0a1 1 0 00-1 1v1a1 1 0 102 0V5a1 1 0 00-1-1zM4 8a1 1 0 00-1 1v1a1 1 0 102 0V9a1 1 0 00-1-1zm12 0a1 1 0 00-1 1v1a1 1 0 102 0V9a1 1 0 00-1-1zM4 12a1 1 0 00-1 1v1a1 1 0 102 0v-1a1 1 0 00-1-1zm12 0a1 1 0 00-1 1v1a1 1 0 102 0v-1a1 1 0 00-1-1z" clip-rule="evenodd"/>
              </svg>
            </div>
          </div>
        </div>
        
        <div class="bg-gradient-to-br from-blue-50 to-blue-100 rounded-lg shadow-sm p-6 border border-blue-200">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-xs text-blue-600">متوسط مدت جلسه</p>
              <p class="text-xl font-bold text-blue-700">{{ stats.avgSessionDuration }}</p>
            </div>
            <div class="bg-blue-200 p-2 rounded-full">
              <svg class="w-5 h-5 text-blue-700" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-12a1 1 0 10-2 0v4a1 1 0 00.293.707l2.828 2.829a1 1 0 101.415-1.415L11 9.586V6z" clip-rule="evenodd"/>
              </svg>
            </div>
          </div>
        </div>
        
        <div class="bg-gradient-to-br from-green-50 to-green-100 rounded-lg shadow-sm p-6 border border-green-200">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-xs text-green-600">نرخ تعامل</p>
              <p class="text-xl font-bold text-green-700">{{ stats.engagementRate }}%</p>
            </div>
            <div class="bg-green-200 p-2 rounded-full">
              <svg class="w-5 h-5 text-green-700" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M3 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm0 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm0 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm0 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z" clip-rule="evenodd"/>
              </svg>
            </div>
          </div>
        </div>
        
        <div class="bg-gradient-to-br from-yellow-50 to-yellow-100 rounded-lg shadow-sm p-6 border border-yellow-200">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-xs text-yellow-600">نرخ پرش</p>
              <p class="text-xl font-bold text-yellow-700">{{ stats.bounceRate }}%</p>
            </div>
            <div class="bg-yellow-200 p-2 rounded-full">
              <svg class="w-5 h-5 text-yellow-700" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd"/>
              </svg>
            </div>
          </div>
        </div>
        
        <div class="bg-gradient-to-br from-purple-50 to-purple-100 rounded-lg shadow-sm p-6 border border-purple-200">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-xs text-purple-600">صفحات/جلسه</p>
              <p class="text-xl font-bold text-purple-700">{{ stats.pagesPerSession }}</p>
            </div>
            <div class="bg-purple-200 p-2 rounded-full">
              <svg class="w-5 h-5 text-purple-700" fill="currentColor" viewBox="0 0 20 20">
                <path d="M9 2a1 1 0 000 2h2a1 1 0 100-2H9z"/>
                <path fill-rule="evenodd" d="M4 5a2 2 0 012-2v1a1 1 0 102 0V3a2 2 0 012 2v6a2 2 0 01-2 2H6a2 2 0 01-2-2V5zm3 4a1 1 0 000 2h.01a1 1 0 100-2H7zm3 0a1 1 0 000 2h3a1 1 0 100-2h-3zm-3 4a1 1 0 100 2h.01a1 1 0 100-2H7zm3 0a1 1 0 100 2h3a1 1 0 100-2h-3z" clip-rule="evenodd"/>
              </svg>
            </div>
          </div>
        </div>
        
        <div class="bg-gradient-to-br from-red-50 to-red-100 rounded-lg shadow-sm p-6 border border-red-200">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-xs text-red-600">کاربران فعال</p>
              <p class="text-xl font-bold text-red-700">{{ stats.activeUsers.toLocaleString('fa-IR') }}</p>
            </div>
            <div class="bg-red-200 p-2 rounded-full">
              <svg class="w-5 h-5 text-red-700" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M10 9a3 3 0 100-6 3 3 0 000 6zm-7 9a7 7 0 1114 0H3z" clip-rule="evenodd"/>
              </svg>
            </div>
          </div>
        </div>
        
        <div class="bg-gradient-to-br from-pink-50 to-pink-100 rounded-lg shadow-sm p-6 border border-pink-200">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-xs text-pink-600">نرخ تبدیل</p>
              <p class="text-xl font-bold text-pink-700">{{ stats.conversionRate }}%</p>
            </div>
            <div class="bg-pink-200 p-2 rounded-full">
              <svg class="w-5 h-5 text-pink-700" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M3 3a1 1 0 000 2v8a2 2 0 002 2h2.586l-1.293 1.293a1 1 0 101.414 1.414L10 15.414l2.293 2.293a1 1 0 001.414-1.414L12.414 15H15a2 2 0 002-2V5a1 1 0 100-2H3zm11.707 4.707a1 1 0 00-1.414-1.414L10 9.586 8.707 8.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"/>
              </svg>
            </div>
          </div>
        </div>
        
        <div class="bg-gradient-to-br from-gray-50 to-gray-100 rounded-lg shadow-sm p-6 border border-gray-200">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-xs text-gray-600">امروز</p>
              <p class="text-xl font-bold text-gray-700">{{ stats.todaySessions.toLocaleString('fa-IR') }}</p>
            </div>
            <div class="bg-gray-200 p-2 rounded-full">
              <svg class="w-5 h-5 text-gray-700" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M6 2a1 1 0 00-1 1v1H4a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-1V3a1 1 0 10-2 0v1H7V3a1 1 0 00-1-1zm0 5a1 1 0 000 2h8a1 1 0 100-2H6z" clip-rule="evenodd"/>
              </svg>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Content Sections -->
    <div class="px-6 space-y-6">
      <!-- Peak Hours & Device Analysis -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- Peak Hours -->
        <div class="bg-white rounded-lg shadow-sm p-6">
          <h3 class="text-lg font-semibold mb-4">ساعات پیک فعالیت</h3>
          <div class="space-y-3">
            <div v-for="hour in peakHours" :key="hour.time" class="p-3 bg-gray-50 rounded-lg">
              <div class="flex justify-between items-center mb-2">
                <span class="font-medium">{{ hour.time }}</span>
                <span class="text-indigo-600 font-bold">{{ hour.sessions.toLocaleString('fa-IR') }}</span>
              </div>
              <div class="w-full bg-gray-200 rounded-full h-2">
                <div class="bg-indigo-500 h-2 rounded-full" :style="`width: ${hour.percentage}%`"></div>
              </div>
              <div class="text-xs text-gray-600 mt-1">{{ hour.percentage }}% از کل جلسات</div>
            </div>
          </div>
        </div>

        <!-- Device Analysis -->
        <div class="bg-white rounded-lg shadow-sm p-6">
          <h3 class="text-lg font-semibold mb-4">تحلیل دستگاه‌ها</h3>
          <div class="space-y-4">
            <div v-for="device in deviceAnalysis" :key="device.type" class="p-6 bg-gradient-to-r from-indigo-50 to-indigo-100 rounded-lg border border-indigo-200">
              <div class="flex justify-between items-center mb-2">
                <span class="font-medium text-indigo-800">{{ device.type }}</span>
                <span class="text-indigo-700 font-bold">{{ device.percentage }}%</span>
              </div>
              <div class="grid grid-cols-3 gap-6 text-sm">
                <div>
                  <p class="text-gray-600">جلسات</p>
                  <p class="font-bold">{{ device.sessions.toLocaleString('fa-IR') }}</p>
                </div>
                <div>
                  <p class="text-gray-600">مدت جلسه</p>
                  <p class="font-bold">{{ device.avgDuration }}</p>
                </div>
                <div>
                  <p class="text-gray-600">نرخ پرش</p>
                  <p class="font-bold">{{ device.bounceRate }}%</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Page Performance & User Journey -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- Top Pages -->
        <div class="bg-white rounded-lg shadow-sm p-6">
          <h3 class="text-lg font-semibold mb-4">پرفورمنس صفحات</h3>
          <div class="space-y-3">
            <div v-for="(page, index) in topPages" :key="page.url" class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
              <div class="flex items-center gap-3">
                <span class="flex items-center justify-center w-6 h-6 bg-indigo-500 text-white rounded-full text-xs">{{ index + 1 }}</span>
                <div>
                  <p class="font-medium text-sm">{{ page.title }}</p>
                  <p class="text-xs text-gray-600">{{ page.url }}</p>
                </div>
              </div>
              <div class="text-left">
                <p class="text-sm font-bold">{{ page.views.toLocaleString('fa-IR') }} بازدید</p>
                <p class="text-xs text-gray-600">{{ page.avgTime }} زمان</p>
              </div>
            </div>
          </div>
        </div>

        <!-- User Journey -->
        <div class="bg-white rounded-lg shadow-sm p-6">
          <h3 class="text-lg font-semibold mb-4">مسیر کاربری</h3>
          <div class="space-y-3">
            <div v-for="journey in userJourneys" :key="journey.id" class="p-3 border border-gray-200 rounded-lg">
              <div class="flex justify-between items-center mb-2">
                <span class="font-medium">{{ journey.name }}</span>
                <span class="text-indigo-600 font-bold">{{ journey.users.toLocaleString('fa-IR') }}</span>
              </div>
              <div class="flex items-center space-x-2 text-sm text-gray-600">
                <span>{{ journey.steps[0] }}</span>
                <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd"/>
                </svg>
                <span>{{ journey.steps[1] }}</span>
                <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd"/>
                </svg>
                <span>{{ journey.steps[2] }}</span>
              </div>
              <div class="text-xs text-gray-500 mt-1">نرخ تکمیل: {{ journey.completionRate }}%</div>
            </div>
          </div>
        </div>
      </div>

      <!-- Engagement Metrics -->
      <div class="bg-white rounded-lg shadow-sm p-6">
        <h3 class="text-lg font-semibold mb-4">معیارهای تعامل</h3>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
          <div class="text-center p-6 bg-gradient-to-br from-blue-50 to-blue-100 rounded-lg border border-blue-200">
            <div class="text-2xl font-bold text-blue-700 mb-1">{{ engagementMetrics.avgScrollDepth }}%</div>
            <div class="text-sm text-blue-600">عمق اسکرول</div>
          </div>
          <div class="text-center p-6 bg-gradient-to-br from-green-50 to-green-100 rounded-lg border border-green-200">
            <div class="text-2xl font-bold text-green-700 mb-1">{{ engagementMetrics.avgTimeOnPage }}</div>
            <div class="text-sm text-green-600">زمان در صفحه</div>
          </div>
          <div class="text-center p-6 bg-gradient-to-br from-purple-50 to-purple-100 rounded-lg border border-purple-200">
            <div class="text-2xl font-bold text-purple-700 mb-1">{{ engagementMetrics.clickRate }}%</div>
            <div class="text-sm text-purple-600">نرخ کلیک</div>
          </div>
          <div class="text-center p-6 bg-gradient-to-br from-orange-50 to-orange-100 rounded-lg border border-orange-200">
            <div class="text-2xl font-bold text-orange-700 mb-1">{{ engagementMetrics.returnRate }}%</div>
            <div class="text-sm text-orange-600">نرخ بازگشت</div>
          </div>
        </div>
      </div>

      <!-- Session Details Table -->
      <div class="bg-white rounded-lg shadow-sm p-6">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold">جزئیات جلسات کاربران</h3>
          <span class="bg-indigo-100 text-indigo-800 px-2 py-1 rounded-full text-xs">{{ sessionDetails.length }} جلسه</span>
        </div>
        <div class="overflow-x-auto">
          <table class="w-full text-sm">
            <thead class="bg-gray-50">
              <tr class="text-right">
                <th class="px-4 py-2">کاربر</th>
                <th class="px-4 py-2">دستگاه</th>
                <th class="px-4 py-2">مرورگر</th>
                <th class="px-4 py-2">مدت جلسه</th>
                <th class="px-4 py-2">صفحات بازدید شده</th>
                <th class="px-4 py-2">آخرین فعالیت</th>
                <th class="px-4 py-2">وضعیت</th>
                <th class="px-4 py-2">عملیات</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="session in sessionDetails" :key="session.id" class="border-t">
                <td class="px-4 py-3">
                  <div>
                    <p class="font-medium">{{ session.user }}</p>
                    <p class="text-xs text-gray-600">{{ session.ip }}</p>
                  </div>
                </td>
                <td class="px-4 py-3">
                  <span :class="getDeviceClass(session.device)" class="px-2 py-1 rounded-full text-xs">
                    {{ session.device }}
                  </span>
                </td>
                <td class="px-4 py-3">{{ session.browser }}</td>
                <td class="px-4 py-3">{{ session.duration }}</td>
                <td class="px-4 py-3">{{ session.pagesViewed }}</td>
                <td class="px-4 py-3">{{ session.lastActivity }}</td>
                <td class="px-4 py-3">
                  <span :class="getSessionStatusClass(session.status)" class="px-2 py-1 rounded-full text-xs">
                    {{ session.status }}
                  </span>
                </td>
                <td class="px-4 py-3">
                  <div class="flex gap-2">
                    <button class="text-indigo-600 hover:text-indigo-800 text-xs">مشاهده</button>
                    <button class="text-blue-600 hover:text-blue-800 text-xs">تحلیل</button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
// declare const useAuth: () => { user: unknown; hasPermission: (perm: string) => boolean }
declare const useHead: (head: { title?: string; meta?: Record<string, unknown>[] }) => void
</script>

<script setup lang="ts">

// Set admin layout
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
// const { user, hasPermission: _hasPermission } = useAuth()

// Sample data
const stats = {
  totalSessions: 28547,
  avgSessionDuration: '3:24',
  engagementRate: 72.8,
  bounceRate: 34.2,
  pagesPerSession: 4.2,
  activeUsers: 12456,
  conversionRate: 2.8,
  todaySessions: 2345
}

const peakHours = [
  { time: '۱۴:۰۰ - ۱۵:۰۰', sessions: 2847, percentage: 35 },
  { time: '۲۰:۰۰ - ۲۱:۰۰', sessions: 2634, percentage: 32 },
  { time: '۱۹:۰۰ - ۲۰:۰۰', sessions: 2456, percentage: 30 },
  { time: '۱۵:۰۰ - ۱۶:۰۰', sessions: 2298, percentage: 28 },
  { time: '۲۱:۰۰ - ۲۲:۰۰', sessions: 2156, percentage: 26 }
]

const deviceAnalysis = [
  { type: 'موبایل', percentage: 58, sessions: 16557, avgDuration: '2:45', bounceRate: 42 },
  { type: 'دسکتاپ', percentage: 35, sessions: 9991, avgDuration: '4:12', bounceRate: 28 },
  { type: 'تبلت', percentage: 7, sessions: 1999, avgDuration: '3:18', bounceRate: 38 }
]

const topPages = [
  { url: '/', title: 'صفحه اصلی', views: 15420, avgTime: '2:15' },
  { url: '/products', title: 'محصولات', views: 12340, avgTime: '3:45' },
  { url: '/product/123', title: 'گوشی آیفون ۱۳', views: 8920, avgTime: '4:20' },
  { url: '/cart', title: 'سبد خرید', views: 6540, avgTime: '1:30' },
  { url: '/checkout', title: 'پرداخت', views: 4320, avgTime: '2:45' }
]

const userJourneys = [
  { id: 1, name: 'خرید محصول', users: 3456, steps: ['صفحه اصلی', 'محصولات', 'سبد خرید'], completionRate: 68 },
  { id: 2, name: 'ثبت‌نام', users: 2345, steps: ['صفحه اصلی', 'ثبت‌نام', 'تایید'], completionRate: 45 },
  { id: 3, name: 'جستجو', users: 5678, steps: ['صفحه اصلی', 'جستجو', 'نتایج'], completionRate: 82 }
]

const engagementMetrics = {
  avgScrollDepth: 65,
  avgTimeOnPage: '2:45',
  clickRate: 12.5,
  returnRate: 34.8
}

const sessionDetails = [
  { id: 1, user: 'کاربر ناشناس', ip: '192.168.1.1', device: 'موبایل', browser: 'Chrome', duration: '4:12', pagesViewed: 6, lastActivity: '۲ دقیقه پیش', status: 'فعال' },
  { id: 2, user: 'احمد احمدی', ip: '192.168.1.2', device: 'دسکتاپ', browser: 'Firefox', duration: '8:45', pagesViewed: 12, lastActivity: '۵ دقیقه پیش', status: 'فعال' },
  { id: 3, user: 'کاربر ناشناس', ip: '192.168.1.3', device: 'تبلت', browser: 'Safari', duration: '2:30', pagesViewed: 3, lastActivity: '۱۰ دقیقه پیش', status: 'خاتمه یافته' },
  { id: 4, user: 'مریم مرادی', ip: '192.168.1.4', device: 'موبایل', browser: 'Chrome', duration: '6:20', pagesViewed: 8, lastActivity: '۱۵ دقیقه پیش', status: 'خاتمه یافته' },
  { id: 5, user: 'کاربر ناشناس', ip: '192.168.1.5', device: 'دسکتاپ', browser: 'Edge', duration: '1:45', pagesViewed: 2, lastActivity: '۲۰ دقیقه پیش', status: 'خاتمه یافته' }
]

const getDeviceClass = (device: string) => {
  switch (device) {
    case 'موبایل':
      return 'bg-green-100 text-green-800'
    case 'دسکتاپ':
      return 'bg-blue-100 text-blue-800'
    case 'تبلت':
      return 'bg-purple-100 text-purple-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getSessionStatusClass = (status: string) => {
  switch (status) {
    case 'فعال':
      return 'bg-green-100 text-green-800'
    case 'خاتمه یافته':
      return 'bg-gray-100 text-gray-800'
    case 'معلق':
      return 'bg-yellow-100 text-yellow-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

// Meta and SEO
useHead({
  title: 'تحلیل رفتار کاربران - پنل مدیریت',
  meta: [
    { name: 'description', content: 'تحلیل جامع تعامل، مسیر کاربری و عملکرد صفحات' }
  ]
})
</script> 

