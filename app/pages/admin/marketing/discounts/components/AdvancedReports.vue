<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200">
    <!-- هدر بخش -->
    <div class="p-6 border-b border-gray-200">
      <div class="flex items-center justify-between">
        <div>
          <h2 class="text-lg font-semibold text-gray-900">سیستم گزارش‌گیری پیشرفته</h2>
          <p class="text-gray-600 mt-1">گزارش‌های جامع، تحلیل‌های پیشرفته و داشبوردهای تعاملی</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors" @click="exportReport">
            <span class="i-heroicons-arrow-down-tray ml-2"></span>
            صادرات گزارش
          </button>
          <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors" @click="scheduleReport">
            <span class="i-heroicons-clock ml-2"></span>
            برنامه‌ریزی گزارش
          </button>
        </div>
      </div>
    </div>

    <!-- فیلترهای گزارش -->
    <div class="p-6 border-b border-gray-200 bg-gray-50">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">بازه زمانی</label>
          <select v-model="reportFilters.timeRange" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm" @change="updateReport">
            <option value="today">امروز</option>
            <option value="week">هفته جاری</option>
            <option value="month">ماه جاری</option>
            <option value="quarter">سه ماهه</option>
            <option value="year">سال جاری</option>
            <option value="custom">سفارشی</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نوع گزارش</label>
          <select v-model="reportFilters.reportType" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm" @change="updateReport">
            <option value="performance">عملکرد</option>
            <option value="financial">مالی</option>
            <option value="user">کاربران</option>
            <option value="campaign">کمپین‌ها</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">دسته‌بندی</label>
          <select v-model="reportFilters.category" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm" @change="updateReport">
            <option value="">همه</option>
            <option value="coupons">کوپن‌ها</option>
            <option value="campaigns">کمپین‌ها</option>
            <option value="users">کاربران</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">فرمت</label>
          <select v-model="reportFilters.format" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
            <option value="pdf">PDF</option>
            <option value="excel">Excel</option>
            <option value="csv">CSV</option>
            <option value="json">JSON</option>
          </select>
        </div>
      </div>
    </div>

    <!-- تب‌های گزارش -->
    <div class="border-b border-gray-200">
      <div class="flex border-b border-gray-200 overflow-x-auto">
        <button
v-for="tab in tabs" :key="tab.value" :class="['px-6 py-3 -mb-px font-medium text-sm focus:outline-none whitespace-nowrap', activeTab === tab.value ? 'border-b-2 border-blue-600 text-blue-700' : 'text-gray-500 hover:text-blue-600']"
          @click="activeTab = tab.value">
          {{ tab.label }}
        </button>
      </div>
    </div>

    <!-- محتوای تب‌ها -->
    <div class="p-6">
      <!-- داشبورد کلی -->
      <div v-if="activeTab === 'dashboard'" class="space-y-6">
        <!-- KPI های کلیدی -->
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
          <div v-for="kpi in kpis" :key="kpi.name" class="bg-white border border-gray-200 rounded-lg p-6">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-sm text-gray-600">{{ kpi.name }}</p>
                <p class="text-2xl font-bold text-gray-900">{{ kpi.value }}</p>
              </div>
              <div class="w-12 h-12 rounded-lg flex items-center justify-center" :style="{ backgroundColor: kpi.color + '20', color: kpi.color }">
                <span :class="kpi.icon + ' text-xl'"></span>
              </div>
            </div>
            <div class="mt-4 flex items-center">
              <span :class="['text-sm', kpi.trend > 0 ? 'text-green-600' : 'text-red-600']">
                {{ kpi.trend > 0 ? '+' : '' }}{{ kpi.trend }}%
              </span>
              <span class="text-sm text-gray-500 mr-2">نسبت به دوره قبل</span>
            </div>
          </div>
        </div>

        <!-- نمودارهای اصلی -->
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <div class="bg-white border border-gray-200 rounded-lg p-6">
            <h4 class="text-lg font-medium text-gray-900 mb-4">روند فروش</h4>
            <div class="h-64 flex items-center justify-center text-gray-400">[نمودار روند فروش]</div>
          </div>
          <div class="bg-white border border-gray-200 rounded-lg p-6">
            <h4 class="text-lg font-medium text-gray-900 mb-4">توزیع کوپن‌ها</h4>
            <div class="h-64 flex items-center justify-center text-gray-400">[نمودار دایره‌ای توزیع]</div>
          </div>
        </div>

        <!-- جدول عملکرد -->
        <div class="bg-white border border-gray-200 rounded-lg p-6">
          <h4 class="text-lg font-medium text-gray-900 mb-4">عملکرد کمپین‌ها</h4>
          <div class="overflow-x-auto">
            <table class="w-full text-sm">
              <thead>
                <tr class="border-b border-gray-200">
                  <th class="text-right py-3 px-4">کمپین</th>
                  <th class="text-right py-3 px-4">بودجه</th>
                  <th class="text-right py-3 px-4">مصرف شده</th>
                  <th class="text-right py-3 px-4">نرخ تبدیل</th>
                  <th class="text-right py-3 px-4">ROI</th>
                  <th class="text-right py-3 px-4">وضعیت</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="campaign in campaignPerformance" :key="campaign.id" class="border-b border-gray-100">
                  <td class="py-3 px-4">{{ campaign.name }}</td>
                  <td class="py-3 px-4">{{ formatCurrency(campaign.budget) }}</td>
                  <td class="py-3 px-4">{{ formatCurrency(campaign.spent) }}</td>
                  <td class="py-3 px-4">{{ campaign.conversionRate }}%</td>
                  <td class="py-3 px-4">{{ campaign.roi }}%</td>
                  <td class="py-3 px-4">
                    <span :class="['px-2 py-1 rounded-full text-xs', getStatusClass(campaign.status)]">
                      {{ getStatusText(campaign.status) }}
                    </span>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>

      <!-- گزارش‌های مالی -->
      <div v-if="activeTab === 'financial'" class="space-y-6">
        <div class="bg-blue-50 p-6 rounded-lg">
          <h4 class="text-md font-medium text-blue-900 mb-2">گزارش‌های مالی</h4>
          <p class="text-sm text-blue-700">تحلیل دقیق عملکرد مالی و ROI کمپین‌ها</p>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <!-- خلاصه مالی -->
          <div class="bg-white border border-gray-200 rounded-lg p-6">
            <h5 class="font-medium text-gray-900 mb-4">خلاصه مالی</h5>
            <div class="space-y-4">
              <div v-for="item in financialSummary" :key="item.name" class="flex items-center justify-between">
                <span class="text-sm text-gray-700">{{ item.name }}</span>
                <div class="flex items-center space-x-2 space-x-reverse">
                  <span class="font-medium text-gray-900">{{ formatCurrency(item.value) }}</span>
                  <span :class="['text-xs', item.trend > 0 ? 'text-green-600' : 'text-red-600']">
                    {{ item.trend > 0 ? '+' : '' }}{{ item.trend }}%
                  </span>
                </div>
              </div>
            </div>
          </div>

          <!-- نمودار ROI -->
          <div class="bg-white border border-gray-200 rounded-lg p-6">
            <h5 class="font-medium text-gray-900 mb-4">ROI کمپین‌ها</h5>
            <div class="h-48 flex items-center justify-center text-gray-400">[نمودار ROI]</div>
          </div>
        </div>

        <!-- جدول جزئیات مالی -->
        <div class="bg-white border border-gray-200 rounded-lg p-6">
          <h5 class="font-medium text-gray-900 mb-4">جزئیات مالی</h5>
          <div class="overflow-x-auto">
            <table class="w-full text-sm">
              <thead>
                <tr class="border-b border-gray-200">
                  <th class="text-right py-3 px-4">دوره</th>
                  <th class="text-right py-3 px-4">درآمد</th>
                  <th class="text-right py-3 px-4">هزینه</th>
                  <th class="text-right py-3 px-4">سود</th>
                  <th class="text-right py-3 px-4">ROI</th>
                  <th class="text-right py-3 px-4">نرخ رشد</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="period in financialDetails" :key="period.period" class="border-b border-gray-100">
                  <td class="py-3 px-4">{{ period.period }}</td>
                  <td class="py-3 px-4">{{ formatCurrency(period.revenue) }}</td>
                  <td class="py-3 px-4">{{ formatCurrency(period.cost) }}</td>
                  <td class="py-3 px-4">{{ formatCurrency(period.profit) }}</td>
                  <td class="py-3 px-4">{{ period.roi }}%</td>
                  <td class="py-3 px-4">
                    <span :class="['text-sm', period.growth > 0 ? 'text-green-600' : 'text-red-600']">
                      {{ period.growth > 0 ? '+' : '' }}{{ period.growth }}%
                    </span>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>

      <!-- گزارش‌های کاربران -->
      <div v-if="activeTab === 'users'" class="space-y-6">
        <div class="bg-green-50 p-6 rounded-lg">
          <h4 class="text-md font-medium text-green-900 mb-2">گزارش‌های کاربران</h4>
          <p class="text-sm text-green-700">تحلیل رفتار کاربران و الگوهای استفاده از کوپن‌ها</p>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <!-- آمار کاربران -->
          <div class="bg-white border border-gray-200 rounded-lg p-6">
            <h5 class="font-medium text-gray-900 mb-4">آمار کاربران</h5>
            <div class="space-y-4">
              <div v-for="stat in userStats" :key="stat.name" class="flex items-center justify-between">
                <span class="text-sm text-gray-700">{{ stat.name }}</span>
                <div class="flex items-center space-x-2 space-x-reverse">
                  <div class="w-32 bg-gray-200 rounded-full h-2">
                    <div class="h-2 rounded-full bg-green-600" :style="{ width: stat.percentage + '%' }"></div>
                  </div>
                  <span class="text-sm text-gray-600">{{ stat.value }}</span>
                </div>
              </div>
            </div>
          </div>

          <!-- نمودار رفتار کاربران -->
          <div class="bg-white border border-gray-200 rounded-lg p-6">
            <h5 class="font-medium text-gray-900 mb-4">رفتار کاربران</h5>
            <div class="h-48 flex items-center justify-center text-gray-400">[نمودار رفتار کاربران]</div>
          </div>
        </div>

        <!-- جدول کاربران برتر -->
        <div class="bg-white border border-gray-200 rounded-lg p-6">
          <h5 class="font-medium text-gray-900 mb-4">کاربران برتر</h5>
          <div class="overflow-x-auto">
            <table class="w-full text-sm">
              <thead>
                <tr class="border-b border-gray-200">
                  <th class="text-right py-3 px-4">کاربر</th>
                  <th class="text-right py-3 px-4">تعداد کوپن استفاده شده</th>
                  <th class="text-right py-3 px-4">مجموع خرید</th>
                  <th class="text-right py-3 px-4">تخفیف کل</th>
                  <th class="text-right py-3 px-4">آخرین خرید</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="user in topUsers" :key="user.id" class="border-b border-gray-100">
                  <td class="py-3 px-4">{{ user.name }}</td>
                  <td class="py-3 px-4">{{ user.couponsUsed }}</td>
                  <td class="py-3 px-4">{{ formatCurrency(user.totalSpent) }}</td>
                  <td class="py-3 px-4">{{ formatCurrency(user.totalDiscount) }}</td>
                  <td class="py-3 px-4">{{ formatDate(user.lastPurchase) }}</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>

      <!-- گزارش‌های سفارشی -->
      <div v-if="activeTab === 'custom'" class="space-y-6">
        <div class="bg-purple-50 p-6 rounded-lg">
          <h4 class="text-md font-medium text-purple-900 mb-2">گزارش‌های سفارشی</h4>
          <p class="text-sm text-purple-700">ایجاد گزارش‌های سفارشی با فیلترهای پیشرفته</p>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <!-- قالب‌های گزارش -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h5 class="font-medium text-gray-900 mb-4">قالب‌های گزارش</h5>
            <div class="space-y-4">
              <div v-for="template in reportTemplates" :key="template.id" class="border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow cursor-pointer" @click="useReportTemplate(template)">
                <div class="flex items-center justify-between mb-2">
                  <h6 class="font-medium text-gray-900">{{ template.name }}</h6>
                  <span class="text-xs bg-blue-100 text-blue-800 px-2 py-1 rounded">{{ template.type }}</span>
                </div>
                <p class="text-sm text-gray-600 mb-3">{{ template.description }}</p>
                <div class="flex justify-between text-xs text-gray-500">
                  <span>آخرین استفاده: {{ formatDate(template.lastUsed) }}</span>
                  <span>استفاده شده: {{ template.usageCount }}</span>
                </div>
              </div>
            </div>
          </div>

          <!-- فیلترهای پیشرفته -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h5 class="font-medium text-gray-900 mb-4">فیلترهای پیشرفته</h5>
            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">محدوده مبلغ</label>
                <div class="grid grid-cols-2 gap-2">
                  <input v-model="customFilters.minAmount" type="number" placeholder="حداقل" class="px-3 py-2 border border-gray-300 rounded-lg text-sm">
                  <input v-model="customFilters.maxAmount" type="number" placeholder="حداکثر" class="px-3 py-2 border border-gray-300 rounded-lg text-sm">
                </div>
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">نوع کوپن</label>
                <div class="space-y-2">
                  <label v-for="type in couponTypes" :key="type" class="flex items-center">
                    <input v-model="customFilters.couponTypes" :value="type" type="checkbox" class="rounded border-gray-300 ml-2">
                    <span class="text-sm text-gray-700">{{ type }}</span>
                  </label>
                </div>
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
                <select v-model="customFilters.status" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
                  <option value="">همه</option>
                  <option value="active">فعال</option>
                  <option value="expired">منقضی شده</option>
                  <option value="used">استفاده شده</option>
                </select>
              </div>
              
              <button class="w-full px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors" @click="generateCustomReport">
                تولید گزارش سفارشی
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- برنامه‌ریزی گزارش‌ها -->
      <div v-if="activeTab === 'scheduling'" class="space-y-6">
        <div class="bg-orange-50 p-6 rounded-lg">
          <h4 class="text-md font-medium text-orange-900 mb-2">برنامه‌ریزی گزارش‌ها</h4>
          <p class="text-sm text-orange-700">تنظیم ارسال خودکار گزارش‌ها در زمان‌های مشخص</p>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <!-- گزارش‌های برنامه‌ریزی شده -->
          <div class="border border-gray-200 rounded-lg p-6">
            <div class="flex items-center justify-between mb-4">
              <h5 class="font-medium text-gray-900">گزارش‌های برنامه‌ریزی شده</h5>
              <button class="px-3 py-1 bg-blue-600 text-white rounded text-sm hover:bg-blue-700" @click="showScheduleForm = true">
                <span class="i-heroicons-plus ml-1"></span>
                افزودن
              </button>
            </div>
            
            <div class="space-y-4">
              <div v-for="schedule in scheduledReports" :key="schedule.id" class="border border-gray-200 rounded-lg p-6">
                <div class="flex items-center justify-between mb-2">
                  <h6 class="font-medium text-gray-900">{{ schedule.name }}</h6>
                  <span :class="['px-2 py-1 rounded-full text-xs', getStatusClass(schedule.status)]">
                    {{ getStatusText(schedule.status) }}
                  </span>
                </div>
                <p class="text-sm text-gray-600 mb-3">{{ schedule.description }}</p>
                <div class="flex justify-between text-sm">
                  <span class="text-gray-500">فرکانس: {{ schedule.frequency }}</span>
                  <span class="text-gray-500">آخرین ارسال: {{ formatDate(schedule.lastSent) }}</span>
                </div>
                <div class="mt-3 flex items-center space-x-2 space-x-reverse">
                  <button class="text-blue-600 hover:text-blue-900 text-sm" @click="editSchedule(schedule)">
                    <span class="i-heroicons-pencil-square ml-1"></span>
                    ویرایش
                  </button>
                  <button class="text-red-600 hover:text-red-900 text-sm" @click="deleteSchedule(schedule)">
                    <span class="i-heroicons-trash ml-1"></span>
                    حذف
                  </button>
                </div>
              </div>
            </div>
          </div>

          <!-- تنظیمات ارسال -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h5 class="font-medium text-gray-900 mb-4">تنظیمات ارسال</h5>
            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">فرمت پیش‌فرض</label>
                <select v-model="schedulingSettings.defaultFormat" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
                  <option value="pdf">PDF</option>
                  <option value="excel">Excel</option>
                  <option value="csv">CSV</option>
                </select>
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">زمان ارسال</label>
                <input v-model="schedulingSettings.sendTime" type="time" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm">
              </div>
              
              <div class="flex items-center justify-between">
                <div>
                  <span class="text-sm font-medium text-gray-700">فشرده‌سازی فایل</span>
                  <p class="text-xs text-gray-500">فشرده کردن فایل‌های بزرگ</p>
                </div>
                <input v-model="schedulingSettings.compression" type="checkbox" class="rounded border-gray-300">
              </div>
              
              <div class="flex items-center justify-between">
                <div>
                  <span class="text-sm font-medium text-gray-700">هشدار خطا</span>
                  <p class="text-xs text-gray-500">ارسال هشدار در صورت خطا</p>
                </div>
                <input v-model="schedulingSettings.errorAlert" type="checkbox" class="rounded border-gray-300">
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
</script>

<script setup lang="ts">
import { ref, reactive } from 'vue'

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

const activeTab = ref('dashboard')
const showScheduleForm = ref(false)

const tabs = [
  { value: 'dashboard', label: 'داشبورد کلی' },
  { value: 'financial', label: 'گزارش‌های مالی' },
  { value: 'users', label: 'گزارش‌های کاربران' },
  { value: 'custom', label: 'گزارش‌های سفارشی' },
  { value: 'scheduling', label: 'برنامه‌ریزی' }
]

const reportFilters = reactive({
  timeRange: 'month',
  reportType: 'performance',
  category: '',
  format: 'pdf'
})

const kpis = ref([
  {
    name: 'کل فروش',
    value: '۲,۵۴۰,۰۰۰,۰۰۰ تومان',
    trend: 12.5,
    color: '#3b82f6',
    icon: 'i-heroicons-currency-dollar'
  },
  {
    name: 'تعداد کوپن‌های استفاده شده',
    value: '۱۵,۶۷۸',
    trend: 8.3,
    color: '#10b981',
    icon: 'i-heroicons-ticket'
  },
  {
    name: 'نرخ تبدیل',
    value: '۲۳.۵٪',
    trend: -2.1,
    color: '#f59e0b',
    icon: 'i-heroicons-chart-bar'
  },
  {
    name: 'ROI کل',
    value: '۱۸۷٪',
    trend: 15.7,
    color: '#8b5cf6',
    icon: 'i-heroicons-trending-up'
  }
])

const campaignPerformance = ref([
  {
    id: 1,
    name: 'کمپین زمستانه',
    budget: 50000000,
    spent: 45000000,
    conversionRate: 18.5,
    roi: 145,
    status: 'active'
  },
  {
    id: 2,
    name: 'تخفیف ویژه',
    budget: 30000000,
    spent: 28000000,
    conversionRate: 22.1,
    roi: 167,
    status: 'completed'
  }
])

const financialSummary = ref([
  { name: 'درآمد کل', value: 2540000000, trend: 12.5 },
  { name: 'هزینه کل', value: 850000000, trend: 8.2 },
  { name: 'سود خالص', value: 1690000000, trend: 15.8 },
  { name: 'ROI متوسط', value: 187, trend: 9.3 }
])

const financialDetails = ref([
  {
    period: 'ماه جاری',
    revenue: 850000000,
    cost: 280000000,
    profit: 570000000,
    roi: 203,
    growth: 12.5
  },
  {
    period: 'ماه قبل',
    revenue: 750000000,
    cost: 260000000,
    profit: 490000000,
    roi: 188,
    growth: 8.3
  }
])

const userStats = ref([
  { name: 'کاربران فعال', value: '۱۲,۴۵۶', percentage: 78 },
  { name: 'کاربران جدید', value: '۲,۳۴۵', percentage: 45 },
  { name: 'نرخ بازگشت', value: '۶۷٪', percentage: 67 },
  { name: 'میانگین خرید', value: '۱۸۵,۰۰۰ تومان', percentage: 82 }
])

const topUsers = ref([
  {
    id: 1,
    name: 'علی احمدی',
    couponsUsed: 15,
    totalSpent: 2500000,
    totalDiscount: 375000,
    lastPurchase: '2024-01-15T10:30:00'
  },
  {
    id: 2,
    name: 'فاطمه محمدی',
    couponsUsed: 12,
    totalSpent: 1800000,
    totalDiscount: 270000,
    lastPurchase: '2024-01-14T15:45:00'
  }
])

const reportTemplates = ref([
  {
    id: 1,
    name: 'گزارش عملکرد ماهانه',
    type: 'performance',
    description: 'گزارش جامع عملکرد ماهانه کوپن‌ها و کمپین‌ها',
    lastUsed: '2024-01-15T10:30:00',
    usageCount: 45
  },
  {
    id: 2,
    name: 'گزارش مالی فصلی',
    type: 'financial',
    description: 'گزارش مالی دقیق با تحلیل ROI و سودآوری',
    lastUsed: '2024-01-10T09:15:00',
    usageCount: 23
  }
])

const couponTypes = ['درصدی', 'مبلغ ثابت', 'ارسال رایگان', 'کد تخفیف']

const customFilters = reactive({
  minAmount: '',
  maxAmount: '',
  couponTypes: [] as string[],
  status: ''
})

const scheduledReports = ref([
  {
    id: 1,
    name: 'گزارش هفتگی',
    description: 'گزارش عملکرد هفتگی کوپن‌ها',
    frequency: 'هفتگی',
    status: 'active',
    lastSent: '2024-01-15T10:30:00'
  },
  {
    id: 2,
    name: 'گزارش مالی ماهانه',
    description: 'گزارش مالی ماهانه',
    frequency: 'ماهانه',
    status: 'active',
    lastSent: '2024-01-01T09:00:00'
  }
])

const schedulingSettings = reactive({
  defaultFormat: 'pdf',
  sendTime: '09:00',
  compression: true,
  errorAlert: true
})

const formatDate = (date: string): string => {
  return new Intl.DateTimeFormat('fa-IR').format(new Date(date))
}

const formatCurrency = (amount: number): string => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

const getStatusClass = (status: string): string => {
  const classes = {
    active: 'bg-green-100 text-green-800',
    completed: 'bg-blue-100 text-blue-800',
    paused: 'bg-yellow-100 text-yellow-800'
  }
  return classes[status as keyof typeof classes] || 'bg-gray-100 text-gray-800'
}

const getStatusText = (status: string): string => {
  const texts = {
    active: 'فعال',
    completed: 'تکمیل شده',
    paused: 'متوقف'
  }
  return texts[status as keyof typeof texts] || status
}

const updateReport = () => {
  // پیاده‌سازی به‌روزرسانی گزارش بر اساس فیلترها

}

const exportReport = () => {
  // پیاده‌سازی صادرات گزارش

}

const scheduleReport = () => {
  showScheduleForm.value = true
}

interface ReportTemplate {
  id?: number | string
  name?: string
  [key: string]: unknown
}

interface ReportSchedule {
  id?: number | string
  name?: string
  [key: string]: unknown
}

const useReportTemplate = (_template: ReportTemplate) => {
  // پیاده‌سازی استفاده از قالب گزارش

}

const generateCustomReport = () => {
  // پیاده‌سازی تولید گزارش سفارشی

}

const editSchedule = (_schedule: ReportSchedule) => {
  // پیاده‌سازی ویرایش برنامه‌ریزی

}

const deleteSchedule = (schedule: ReportSchedule) => {
  if (confirm(`آیا از حذف برنامه‌ریزی "${schedule.name}" اطمینان دارید؟`)) {
    const index = scheduledReports.value.findIndex(s => s.id === schedule.id)
    if (index !== -1) {
      scheduledReports.value.splice(index, 1)
    }
  }
}
</script>

<!--
  کامپوننت سیستم گزارش‌گیری پیشرفته
  شامل:
  1. داشبورد کلی با KPI ها
  2. گزارش‌های مالی دقیق
  3. گزارش‌های کاربران
  4. گزارش‌های سفارشی
  5. برنامه‌ریزی گزارش‌ها
  توضیحات کامل به فارسی در کد
--> 
