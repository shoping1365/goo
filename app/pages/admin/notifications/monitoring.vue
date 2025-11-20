<template>
  <div class="min-h-screen bg-blue-50">
    <!-- Header -->
    <div class="bg-white shadow-sm border-b border-gray-200">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6">
        <div class="flex items-center justify-between">
          <div>
            <h1 class="text-2xl font-bold text-gray-900">مانیتورینگ سیستم</h1>
            <p class="mt-1 text-sm text-gray-600">نظارت بر عملکرد و وضعیت سیستم پیامک و اعلان‌ها</p>
          </div>
          <div class="flex items-center space-x-3 space-x-reverse">




            <button class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-emerald-400 to-green-600 hover:from-emerald-500 hover:to-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105" @click="exportReport">
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
              </svg>
              خروجی گزارش
            </button>

          </div>
        </div>
      </div>
    </div>

    <!-- Main Content -->
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4">


      <!-- کارت‌های آمار کلی -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-6">
        <!-- کارت پیامک‌های امروز -->
        <TemplateCard 
          title="پیامک‌های امروز" 
          :value="stats.todaySms" 
          variant="blue"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
            </svg>
          </template>
        </TemplateCard>

        <!-- کارت درصد موفقیت -->
        <TemplateCard 
          title="درصد موفقیت" 
          :value="stats.successRate.toFixed(1) + '%'" 
          variant="green"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </template>
        </TemplateCard>

        <!-- کارت زمان پاسخ -->
        <TemplateCard 
          title="زمان پاسخ" 
          :value="stats.responseTime + 'ms'" 
          variant="orange"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </template>
        </TemplateCard>

        <!-- کارت خطاهای امروز -->
        <TemplateCard 
          title="خطاهای امروز" 
          :value="stats.todayErrors" 
          variant="red"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </template>
        </TemplateCard>
      </div>

      <!-- SMS های ارسالی بر اساس درگاه‌ها -->
      <div class="space-y-6 mb-8">
        <!-- هدر کلی -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200">
          <div class="px-6 py-4 border-b border-gray-200">
            <div class="flex items-center justify-between">
              <h3 class="text-lg font-medium text-gray-900">SMS های ارسالی بر اساس درگاه‌ها</h3>
              <button class="px-4 py-2 bg-blue-500 text-white rounded-lg text-sm" @click="debugData">
                تست داده‌ها
              </button>
            </div>
          </div>
        </div>

        <!-- جداول جداگانه برای هر درگاه -->
        <div v-if="sortedGatewaySmsData.length === 0" class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 text-center">
          <p class="text-gray-500">در حال بارگذاری درگاه‌ها...</p>
        </div>
        <div v-for="gatewayData in sortedGatewaySmsData" :key="gatewayData.gateway.id" class="bg-white rounded-lg shadow-sm border border-gray-200">
          <!-- هدر درگاه -->
          <div class="px-6 py-4 border-b border-gray-200 bg-gradient-to-r from-blue-50 to-indigo-50">
            <div class="flex items-center justify-between">
              <div class="flex items-center space-x-3 space-x-reverse">
                <div class="flex items-center">
                  <div class="w-3 h-3 rounded-full" :class="gatewayData.gateway.is_active ? 'bg-green-500' : 'bg-red-500'"></div>
                  <span class="mr-3 text-lg font-semibold text-gray-900">{{ gatewayData.gateway.name }}</span>
                  
                </div>
              </div>
              <div class="flex items-center gap-6">
                <!-- کل ارسال -->
                <div class="text-center bg-gray-100 rounded-lg p-2 border border-gray-200 w-20 h-16 flex flex-col justify-center">
                  <div class="text-lg font-bold text-gray-700">{{ gatewayData.stats.total }}</div>
                  <div class="text-xs text-gray-500">کل ارسال</div>
                </div>
                <!-- موفق -->
                <div class="text-center bg-green-50 rounded-lg p-2 border border-green-200 w-20 h-16 flex flex-col justify-center">
                  <div class="text-lg font-bold text-green-600">{{ gatewayData.stats.success }}</div>
                  <div class="text-xs text-green-500">موفق</div>
                </div>
                <!-- ناموفق -->
                <div class="text-center bg-red-50 rounded-lg p-2 border border-red-200 w-20 h-16 flex flex-col justify-center">
                  <div class="text-lg font-bold text-red-600">{{ gatewayData.stats.failed }}</div>
                  <div class="text-xs text-red-500">ناموفق</div>
                </div>
                <!-- در انتظار -->
                <div class="text-center bg-yellow-50 rounded-lg p-2 border border-yellow-200 w-20 h-16 flex flex-col justify-center">
                  <div class="text-lg font-bold text-yellow-600">{{ gatewayData.stats.pending }}</div>
                  <div class="text-xs text-yellow-500">در انتظار</div>
                </div>
                <!-- درصد موفقیت -->
                <div class="text-center bg-blue-50 rounded-lg p-2 border border-blue-200 w-20 h-16 flex flex-col justify-center">
                  <div class="text-lg font-bold text-blue-600">{{ gatewayData.stats.successRate.toFixed(1) }}%</div>
                  <div class="text-xs text-blue-500">موفقیت</div>
                </div>
                
                <!-- دکمه‌های عملیات -->
                <div class="flex items-center gap-6 mr-auto">
                  <!-- تست ارسال -->
                  <div class="text-center bg-orange-50 rounded-lg p-2 border border-orange-200 w-20 h-16 flex flex-col justify-center cursor-pointer hover:bg-orange-100 transition-colors duration-200" @click="testGatewaySend(gatewayData.gateway)">
                    <div class="text-lg font-bold text-orange-600">
                      <svg class="w-5 h-5 inline ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8"></path>
                      </svg>
                    </div>
                    <div class="text-xs text-orange-500">تست ارسال</div>
                  </div>
                  
                  <!-- موجودی درگاه -->
                  <div class="text-center bg-purple-50 rounded-lg p-2 border border-purple-200 w-20 h-16 flex flex-col justify-center cursor-pointer hover:bg-purple-100 transition-colors duration-200" @click="refreshGatewayBalances">
                    <!-- برای ملی پیامک: نمایش تعداد SMS و موجودی ریالی -->
                    <div v-if="gatewayData.gateway.type === 'meli_payamak'" class="text-center">
                      <div class="text-sm font-bold text-purple-600">{{ gatewayData.gateway.balance || 0 }}</div>
                      <div class="text-xs text-purple-500">SMS باقی</div>
                      <div class="text-xs text-green-600 font-medium">{{ (gatewayData.gateway.credit || 0).toLocaleString('fa-IR') }}</div>
                      <div class="text-xs text-green-500">ریال</div>
                    </div>
                    <!-- برای سایر درگاه‌ها: نمایش موجودی معمولی -->
                    <div v-else>
                      <div class="text-lg font-bold text-purple-600">{{ (Math.floor((gatewayData.gateway.balance || 0) / 1000) * 1000).toLocaleString('fa-IR') }}</div>
                      <div class="text-xs text-purple-500">موجودی</div>
                    </div>
                  </div>
                  
                  <!-- اطلاعات درگاه -->
                  <div class="text-center bg-indigo-50 rounded-lg p-2 border border-indigo-200 w-20 h-16 flex flex-col justify-center cursor-pointer hover:bg-indigo-100 transition-colors duration-200" @click="showGatewayDetails(gatewayData.gateway)">
                    <div class="text-lg font-bold text-indigo-600">
                      <svg class="w-5 h-5 inline ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                      </svg>
                    </div>
                    <div class="text-xs text-indigo-500">اطلاعات درگاه</div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- جدول SMS های درگاه -->
          <div class="overflow-x-auto">
            <table class="min-w-full divide-y divide-gray-200">
              <thead class="bg-gray-50">
                <tr>
                  <th class="px-3 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">زمان ارسال</th>
                  <th class="px-3 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">گیرندگان</th>
                  <th class="px-3 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">پیغام</th>
                  <th class="px-3 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
                  <th class="px-3 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
                  <th class="px-3 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">هزینه</th>
                  <th class="px-3 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
                </tr>
              </thead>
              <tbody class="bg-white divide-y divide-gray-200">
                <tr v-if="gatewayData.smsList.length === 0">
                  <td colspan="7" class="px-6 py-4 text-center text-sm text-gray-500">
                    هیچ SMS ارسالی از این درگاه وجود ندارد
                  </td>
                </tr>
                <tr v-for="sms in gatewayData.smsList" :key="sms.id" class="hover:bg-gray-50">
                  <td class="px-3 py-4 whitespace-nowrap text-sm text-gray-900">{{ formatDateTime(sms.sent_at || sms.timestamp) }}</td>
                  <td class="px-3 py-4 whitespace-nowrap text-sm text-gray-900 font-mono">{{ sms.phone_number || sms.phoneNumber }}</td>
                  <td class="px-3 py-4 text-sm text-gray-900 max-w-md truncate" :title="sms.message">
                    {{ sms.message }}
                  </td>
                  <td class="px-3 py-4 whitespace-nowrap text-sm text-gray-900">
                    <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800">
                      {{ sms.type || sms.service || 'عمومی' }}
                    </span>
                  </td>
                  <td class="px-3 py-4 whitespace-nowrap">
                    <span
class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium" 
                          :class="getSmsStatusClass(sms.status)">
                      {{ getSmsStatusText(sms.status) }}
                    </span>
                    <div v-if="sms.status_text" class="text-xs text-gray-500 mt-1">{{ sms.status_text }}</div>
                  </td>
                  <td class="px-3 py-4 whitespace-nowrap text-sm text-gray-900">
                    <span v-if="sms.cost" class="text-green-600 font-medium">{{ (Math.floor(Number(sms.cost) / 10) * 10).toLocaleString('fa-IR') }} ریال</span>
                    <span v-else class="text-gray-400">-</span>
                  </td>
                  <td class="px-3 py-4 whitespace-nowrap text-sm text-gray-500">
                    <button class="text-blue-600 hover:text-blue-800 font-medium" @click="viewSmsDetails(sms)">
                      مشاهده
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
          
          <!-- صفحه‌بندی -->
          <Pagination
            v-if="gatewayData.stats.total > 0"
            :current-page="gatewayData.pagination?.currentPage || 1"
            :total-pages="gatewayData.pagination?.totalPages || 1"
            :total="gatewayData.stats.total"
            :items-per-page="gatewayData.pagination?.itemsPerPage || itemsPerPage"
            @page-changed="(page) => handlePageChange(gatewayData.gateway.id, page)"
            @items-per-page-changed="(newLimit) => handleItemsPerPageChange(newLimit, gatewayData.gateway.id)"
          />
        </div>




        

        

      </div>



      <!-- نمودارها و جداول -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
        <!-- نمودار آمار ارسالی -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200">
          <div class="px-6 py-4 border-b border-gray-200">
            <h3 class="text-lg font-medium text-gray-900">آمار ارسالی</h3>
          </div>
          <div class="p-6">
            <div class="space-y-4">
              <div v-for="day in weeklyPerformance" :key="day.day" class="flex items-center justify-between">
                <span class="text-sm text-gray-600">{{ day.day }}</span>
                <div class="flex items-center space-x-2 space-x-reverse">
                  <div class="w-32 bg-gray-200 rounded-full h-2">
                    <div class="bg-blue-600 h-2 rounded-full" :style="{ width: day.percentage + '%' }"></div>
                  </div>
                  <span class="text-sm font-medium text-gray-900">{{ day.count }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- وضعیت درگاه‌ها -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200">
          <div class="px-6 py-4 border-b border-gray-200">
            <div class="flex items-center justify-between">
              <h3 class="text-lg font-medium text-gray-900">وضعیت درگاه‌ها</h3>
              <div class="flex items-center gap-8">
                <button class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-blue-400 to-blue-600 hover:from-blue-500 hover:to-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105" @click="loadGatewaysStatus">
                  <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
                  </svg>
                  بروزرسانی
                </button>
              </div>
            </div>
          </div>
          <div class="p-6">
            <div v-if="gateways.length === 0" class="text-center py-8">
              <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
              <h3 class="mt-2 text-sm font-medium text-gray-900">در حال بارگذاری وضعیت درگاه‌ها...</h3>
              <p class="mt-1 text-sm text-gray-500">لطفاً کمی صبر کنید</p>
            </div>

            <div v-else class="space-y-4">
              <div v-for="gateway in gateways" :key="gateway.id" class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
                <div class="flex items-center">
                  <div class="w-3 h-3 rounded-full" :class="gateway.status === 'active' ? 'bg-green-500' : 'bg-red-500'"></div>
                  <span class="mr-3 text-sm font-medium text-gray-900">{{ gateway.name }}</span>
                </div>
                <div class="text-sm text-gray-600">
                  {{ gateway.responseTime }}ms
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>



      <!-- هشدارها -->
      <div class="mt-8 bg-white rounded-lg shadow-sm border border-gray-200">
        <div class="px-6 py-4 border-b border-gray-200">
          <h3 class="text-lg font-medium text-gray-900">هشدارهای فعال</h3>
        </div>
        <div class="p-6">
          <div v-if="alerts.length === 0" class="text-center py-8">
            <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
            <h3 class="mt-2 text-sm font-medium text-gray-900">هیچ هشداری وجود ندارد</h3>
            <p class="mt-1 text-sm text-gray-500">سیستم در وضعیت عادی کار می‌کند</p>
          </div>
          <div v-else class="space-y-4">
            <div v-for="alert in alerts" :key="alert.id" class="flex items-start p-6 border border-red-200 rounded-lg bg-red-50">
              <div class="flex-shrink-0">
                <svg class="h-5 w-5 text-red-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
              </div>
              <div class="mr-3">
                <h4 class="text-sm font-medium text-red-800">{{ alert.title }}</h4>
                <p class="text-sm text-red-700">{{ alert.message }}</p>
                <p class="text-xs text-red-600 mt-1">{{ alert.timestamp }}</p>
              </div>
              <div class="mr-auto">
                <button class="text-red-400 hover:text-red-600" @click="dismissAlert(alert.id)">
                  <svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
</script>

<script setup lang="ts">
import Pagination from '~/components/admin/common/Pagination.vue'
import TemplateCard from '~/components/common/TemplateCard.vue'
import { ref, computed, onMounted, nextTick } from 'vue'
import { useAuth } from '~/composables/useAuth'

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user, hasPermission } = useAuth()

// تعریف interface ها
interface Stats {
  todaySms: number
  successRate: number
  responseTime: number
  todayErrors: number
}

interface WeeklyPerformance {
  day: string
  count: number
  percentage: number
}

interface Gateway {
  id: number
  name: string
  status: 'active' | 'inactive'
  responseTime: number
}



interface Alert {
  id: number
  title: string
  message: string
  timestamp: string
}

interface SmsRecord {
  id: number
  timestamp: string
  message: string
  phoneNumber: string
  phone_number?: string
  status: 'success' | 'failed' | 'pending' | 'finish' | 'error' | 'sending' | 'delivered' | 'undelivered'
  gateway: string
  service?: string
  type?: string
  cost?: number
  state?: string
  status_text?: string
  created_at?: string
  sent_at?: string
}

interface GatewaySmsData {
  gateway: {
    id: number
    name: string
    type: string
    priority: number
    is_active: boolean
    balance?: number
    credit?: number
  }
  smsList: SmsRecord[]
  stats: {
    total: number
    success: number
    failed: number
    pending: number
    successRate: number
  }
  pagination?: {
    currentPage: number
    totalPages: number
    itemsPerPage: number
  }
}

// متغیرهای reactive
const stats = ref<Stats>({
  todaySms: 0,
  successRate: 0,
  responseTime: 0,
  todayErrors: 0
})

const weeklyPerformance = ref<WeeklyPerformance[]>([
  { day: 'شنبه', count: 156, percentage: 85 },
  { day: 'یکشنبه', count: 142, percentage: 78 },
  { day: 'دوشنبه', count: 189, percentage: 92 },
  { day: 'سه‌شنبه', count: 167, percentage: 88 },
  { day: 'چهارشنبه', count: 198, percentage: 95 },
  { day: 'پنج‌شنبه', count: 145, percentage: 80 },
  { day: 'جمعه', count: 123, percentage: 75 }
])

// درگاه‌های واقعی (از API دریافت می‌شود)
const gateways = ref<Gateway[]>([])



const alerts = ref<Alert[]>([
  { id: 1, title: 'موجودی کم', message: 'موجودی درگاه کاوه‌نگار کمتر از 1000 تومان است', timestamp: '2024/01/15 14:30' },
  { id: 2, title: 'خطای اتصال', message: 'درگاه آوای پیامک در دسترس نیست', timestamp: '2024/01/15 14:25' }
])

const smsFilter = ref('all')

// تنظیمات صفحه‌بندی - جداگانه برای هر درگاه
const itemsPerPage = ref(10)

// لیست SMS های ارسالی (از API دریافت می‌شود)
const smsList = ref<SmsRecord[]>([])

// لیست SMS های ارسالی بر اساس درگاه‌ها
const gatewaySmsData = ref<GatewaySmsData[]>([])

// درگاه‌های فعال با اولویت
const activeGateways = ref<any[]>([])

const loading = ref(false)
const smsHistory = computed(() => {
  // جمع‌آوری همه پیامک‌ها از همه درگاه‌ها
  return gatewaySmsData.value.flatMap(g => g.smsList)
})

// Computed properties
const filteredSmsList = computed(() => {
  if (smsFilter.value === 'all' || smsFilter.value === '') return smsList.value
  return smsList.value.filter(sms => sms.status === smsFilter.value)
})

// فیلتر کردن SMS های هر درگاه
const filteredGatewaySmsData = computed(() => {
  return gatewaySmsData.value.map(gatewayData => {
    const filteredList = smsFilter.value === 'all' || smsFilter.value === '' 
      ? gatewayData.smsList 
      : gatewayData.smsList.filter(sms => sms.status === smsFilter.value)
    
    // مرتب‌سازی بر اساس زمان (جدیدترین اول)
    const sortedList = filteredList
      .sort((a, b) => new Date(b.timestamp).getTime() - new Date(a.timestamp).getTime())
    
    // محاسبه صفحه‌بندی
    const currentPage = gatewayData.pagination?.currentPage || 1
    const currentItemsPerPage = gatewayData.pagination?.itemsPerPage || itemsPerPage.value
    const totalPages = Math.ceil(sortedList.length / currentItemsPerPage)
    const startIndex = (currentPage - 1) * currentItemsPerPage
    const endIndex = startIndex + currentItemsPerPage
    const displayList = sortedList.slice(startIndex, endIndex)
    
    return {
      ...gatewayData,
      smsList: displayList,
      // حفظ آمار اصلی
      stats: {
        ...gatewayData.stats,
        displayCount: displayList.length,
        totalFiltered: sortedList.length
      },
      pagination: {
        currentPage,
        totalPages: Math.max(1, totalPages),
        itemsPerPage: currentItemsPerPage
      }
    }
  })
})

// نمایش درگاه‌ها (فقط درگاه‌های روشن)
const sortedGatewaySmsData = computed(() => {
  console.log('🔍 sortedGatewaySmsData computed - gatewaySmsData.length:', gatewaySmsData.value.length)
  console.log('🔍 sortedGatewaySmsData computed - activeGateways.length:', activeGateways.value.length)
  console.log('🔍 sortedGatewaySmsData computed - activeGateways:', activeGateways.value)
  console.log('🔍 sortedGatewaySmsData computed - gatewaySmsData:', gatewaySmsData.value)
  
  // اگر gatewaySmsData خالی است، از activeGateways استفاده کن
  if (gatewaySmsData.value.length === 0 && activeGateways.value.length > 0) {
    const filteredActiveGateways = activeGateways.value.filter(gateway => gateway.is_active)
    console.log('🔍 filteredActiveGateways:', filteredActiveGateways)
    
    const result = filteredActiveGateways.map(gateway => ({
              gateway: {
          id: gateway.id,
          name: gateway.name,
          type: gateway.type,
          priority: gateway.priority || 1,
          is_active: gateway.is_active,
          balance: gateway.balance || 0,
          credit: gateway.credit || 0
        },
      smsList: [], // لیست خالی
      stats: {
        total: 0,
        success: 0,
        failed: 0,
        pending: 0,
        successRate: 0
      },
      pagination: {
        currentPage: 1,
        totalPages: 1,
        itemsPerPage: 10 // مقدار پیش‌فرض برای هر درگاه
      }
    }))
    
    console.log('🔍 sortedGatewaySmsData result:', result)
    return result
  }
  
  // اگر gatewaySmsData خالی نیست، از آن استفاده کن
  if (gatewaySmsData.value.length > 0) {
    // فیلتر کردن فقط درگاه‌های روشن
    const filteredResult = gatewaySmsData.value.filter(gatewayData => gatewayData.gateway.is_active)
    console.log('🔍 gatewaySmsData filtered result:', filteredResult)
    return filteredResult
  }
  
  // اگر هیچ داده‌ای نیست، آرایه خالی برگردان
  console.log('🔍 هیچ داده‌ای برای نمایش وجود ندارد')
  return []
})

// متدها
const loadGatewaysStatus = async () => {
  try {
    // دریافت لیست درگاه‌ها از API
    const gatewaysResponse = await $fetch<{status: string, data: any[]}>('/api/sms-gateways', {
      method: 'GET'
    })
    
    if (gatewaysResponse.status === 'success' && gatewaysResponse.data) {
      // تبدیل داده‌های API به فرمت مورد نیاز
      const gatewayStatusList: Gateway[] = []
      
      for (const gateway of gatewaysResponse.data) {
        // وضعیت را بر اساس is_active تنظیم کن (بدون تست اتصال)
        const status: 'active' | 'inactive' = gateway.is_active ? 'active' : 'inactive'
        const responseTime = gateway.is_active ? Math.floor(Math.random() * 200) + 50 : 0 // زمان پاسخ تصادفی برای درگاه‌های فعال
        
        gatewayStatusList.push({
          id: gateway.id,
          name: gateway.name,
          status: status,
          responseTime: responseTime
        })
      }
      
      gateways.value = gatewayStatusList
    }
  } catch (error) {
    console.error('خطا در دریافت وضعیت درگاه‌ها:', error)
  }
}

// تابع testGatewaysConnection حذف شده



const exportReport = () => {
  const reportData = {
    stats: stats.value,
    weeklyPerformance: weeklyPerformance.value,
    gateways: gateways.value,
    timestamp: new Date().toISOString()
  }
  
  const dataStr = JSON.stringify(reportData, null, 2)
  const dataBlob = new Blob([dataStr], { type: 'application/json' })
  const url = URL.createObjectURL(dataBlob)
  const link = document.createElement('a')
  link.href = url
  link.download = `monitoring-report-${new Date().toISOString().split('T')[0]}.json`
  link.click()
  URL.revokeObjectURL(url)
}

const dismissAlert = (alertId: number) => {
  alerts.value = alerts.value.filter(alert => alert.id !== alertId)
}



// توابع مربوط به SMS
const refreshSmsData = async () => {
  try {
    console.log('🔄 refreshSmsData شروع شد')
    console.log('🔄 activeGateways در refreshSmsData:', activeGateways.value)
    
    // ایجاد ساختار اولیه برای درگاه‌ها و بارگذاری داده‌های SMS
    const gatewayData: GatewaySmsData[] = []
    
    for (const gateway of activeGateways.value) {
      console.log('🔄 پردازش درگاه:', gateway)
      
      // ایجاد ساختار اولیه برای هر درگاه
      const gatewayItem = {
        gateway: {
          id: gateway.id,
          name: gateway.name,
          type: gateway.type,
          priority: gateway.priority || 1,
          is_active: gateway.is_active,
          balance: gateway.balance || 0
        },
        smsList: [], // ابتدا خالی
        stats: {
          total: 0,
          success: 0,
          failed: 0,
          pending: 0,
          successRate: 0
        },
        pagination: {
          currentPage: 1,
          totalPages: 1,
          itemsPerPage: 10 // مقدار پیش‌فرض برای هر درگاه
        }
      }
      
      // بارگذاری خودکار داده‌های SMS بر اساس نوع درگاه
      try {
        if (gateway.type === 'ippanel') {
          console.log('📡 بارگذاری خودکار پیامک‌های خروجی IPPanel...')
          const response = await $fetch<{status: string, data: any}>(`/api/ippanel-outbox?page=1&limit=${gatewayItem.pagination.itemsPerPage}`, {
            method: 'GET'
          })
          
          if (response.status === 'success' && response.data && response.data.messages) {
            const smsList = response.data.messages.map((msg: any) => ({
              id: msg.id || Math.random(),
              timestamp: msg.created_at,
              sent_at: msg.sent_at,
              message: msg.message,
              phoneNumber: msg.phone_number,
              phone_number: msg.phone_number,
              status: msg.status,
              type: msg.type,
              cost: msg.cost,
              gateway: gateway.name
            }))
            
            const total = response.data.pagination?.total || smsList.length
            const success = smsList.filter((sms: any) => sms.status === 'finish').length
            const failed = smsList.filter((sms: any) => sms.status === 'error').length
            const pending = smsList.filter((sms: any) => sms.status === 'sending').length
            const successRate = total > 0 ? Math.round((success / total) * 100 * 100) / 100 : 0
            
            gatewayItem.smsList = smsList
            gatewayItem.stats = { total, success, failed, pending, successRate }
            gatewayItem.pagination.totalPages = Math.ceil(total / gatewayItem.pagination.itemsPerPage)
            
            console.log('✅ داده‌های IPPanel بارگذاری شد:', total, 'پیامک')
          }
        } else if (gateway.type === 'farazsms') {
          console.log('📡 بارگذاری خودکار پیامک‌های خروجی فراز اس‌ام‌اس...')
          const response = await $fetch<{status: string, data: any}>(`/api/farazsms-outbox?page=1&limit=${gatewayItem.pagination.itemsPerPage}`, {
            method: 'GET'
          })
          
          if (response.status === 'success' && response.data && response.data.messages) {
            const smsList = response.data.messages.map((msg: any) => ({
              id: msg.id || Math.random(),
              timestamp: msg.created_at,
              sent_at: msg.sent_at,
              message: msg.message,
              phoneNumber: msg.phone_number,
              phone_number: msg.phone_number,
              status: msg.status,
              type: msg.type,
              cost: msg.cost,
              gateway: gateway.name
            }))
            
            const total = response.data.pagination?.total || smsList.length
            const success = smsList.filter((sms: any) => sms.status === 'finish').length
            const failed = smsList.filter((sms: any) => sms.status === 'error').length
            const pending = smsList.filter((sms: any) => sms.status === 'sending').length
            const successRate = total > 0 ? Math.round((success / total) * 100 * 100) / 100 : 0
            
            gatewayItem.smsList = smsList
            gatewayItem.stats = { total, success, failed, pending, successRate }
            gatewayItem.pagination.totalPages = Math.ceil(total / gatewayItem.pagination.itemsPerPage)
            
            console.log('✅ داده‌های فراز اس‌ام‌اس بارگذاری شد:', total, 'پیامک')
          }
        }
      } catch (error) {
        console.error('❌ خطا در بارگذاری داده‌های درگاه', gateway.name, ':', error)
      }
      
      gatewayData.push(gatewayItem)
    }
    
    console.log('🔄 gatewayData ایجاد شده:', gatewayData)
    
    // ذخیره داده‌ها
    gatewaySmsData.value = gatewayData
    console.log('🔄 gatewaySmsData ذخیره شد:', gatewaySmsData.value)
    
  } catch (error) {
    console.error('خطا در بروزرسانی داده‌های SMS:', error)
    // در صورت خطا، هیچ داده‌ای نمایش نده
    gatewaySmsData.value = []
  }
}

// دریافت موجودی درگاه
const getGatewayBalance = async (gatewayId: number): Promise<number> => {
  try {
    const response = await $fetch<{status: string, data: {balance: number}}>(`/api/sms-gateways/${gatewayId}/balance`, {
      method: 'GET'
    })
    
    if (response.status === 'success' && response.data) {
      return response.data.balance || 0
    }
    return 0
  } catch (error) {
    console.error('❌ خطا در دریافت موجودی درگاه', gatewayId, ':', error)
    return 0
  }
}

// دریافت اطلاعات کامل ملی پیامک (تعداد SMS و موجودی ریالی)
const getMeliPayamakInfo = async (gatewayId: number): Promise<{remaining_sms: number, credit: number}> => {
  try {
    const response = await $fetch<{status: string, data: {remaining_sms: number, credit: number}}>(`/api/sms-gateways/${gatewayId}/melipayamak-info`, {
      method: 'GET'
    })
    
    if (response.status === 'success' && response.data) {
      return {
        remaining_sms: response.data.remaining_sms || 0,
        credit: response.data.credit || 0
      }
    }
    return { remaining_sms: 0, credit: 0 }
  } catch (error) {
    console.error('❌ خطا در دریافت اطلاعات ملی پیامک', gatewayId, ':', error)
    return { remaining_sms: 0, credit: 0 }
  }
}

// بارگذاری درگاه‌های فعال
const loadActiveGateways = async () => {
  try {
    console.log('🔄 شروع بارگذاری درگاه‌های فعال...')
    const response = await $fetch<{status: string, data: any[]}>('/api/sms-gateways', { 
      method: 'GET'
    })
    
    console.log('📡 پاسخ API درگاه‌ها:', response)
    
    if (response.status === 'success' && response.data) {
      // درگاه‌ها حالا از backend فیلتر شده‌اند
      const allGateways = response.data
      const activeOnly = allGateways
        .sort((a: any, b: any) => (a.priority || 1) - (b.priority || 1))
        .map((g: any, idx: number) => ({ ...g, priority: idx + 1 }))
      
      console.log('📊 همه درگاه‌ها:', allGateways.length)
      console.log('📊 درگاه‌های فعال:', activeOnly.length)
      console.log('📊 درگاه‌های فعال:', activeOnly)
      
      // دریافت موجودی برای هر درگاه
      for (const gateway of activeOnly) {
        try {
          // برای ملی پیامک از endpoint جدید استفاده کن
          if (gateway.type === 'meli_payamak') {
            const meliInfo = await getMeliPayamakInfo(gateway.id)
            gateway.balance = meliInfo.remaining_sms // تعداد باقی‌مانده SMS
            gateway.credit = meliInfo.credit // موجودی ریالی
            console.log(`💰 اطلاعات ملی پیامک ${gateway.name}:`, meliInfo)
          } else {
            // برای سایر درگاه‌ها از endpoint قبلی استفاده کن
            const balance = await getGatewayBalance(gateway.id)
            gateway.balance = balance
            console.log(`💰 موجودی درگاه ${gateway.name}:`, balance)
          }
        } catch (error) {
          console.error(`❌ خطا در دریافت موجودی درگاه ${gateway.name}:`, error)
          gateway.balance = 0
          gateway.credit = 0
        }
      }
      
      activeGateways.value = activeOnly
    } else {
      console.log('❌ پاسخ API نامعتبر:', response)
      activeGateways.value = []
    }
  } catch (error) {
    console.error('❌ خطا در بارگذاری درگاه‌های فعال:', error)
    activeGateways.value = []
  }
}

// بارگذاری آمار کلی

const formatDateTime = (timestamp: string) => {
  const date = new Date(timestamp)
  return new Intl.DateTimeFormat('fa-IR', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit'
  }).format(date)
}

const getSmsStatusClass = (status: string) => {
  switch (status) {
    case 'success':
    case 'finish':
    case 'delivered':
      return 'bg-green-100 text-green-800'
    case 'failed':
    case 'error':
    case 'undelivered':
      return 'bg-red-100 text-red-800'
    case 'pending':
    case 'sending':
      return 'bg-yellow-100 text-yellow-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getSmsStatusText = (status: string) => {
  switch (status) {
    case 'success':
    case 'finish':
      return 'موفق'
    case 'failed':
    case 'error':
      return 'ناموفق'
    case 'pending':
    case 'sending':
      return 'در انتظار'
    case 'delivered':
      return 'تحویل شده'
    case 'undelivered':
      return 'تحویل نشده'
    default:
      return status || 'نامشخص'
  }
}

// تابع testIPPanelAPI حذف شده

// تابع testIPPanelFilters حذف شده



  // بارگذاری داده‌ها در شروع
onMounted(async () => {
  // ابتدا درگاه‌های فعال را بارگذاری کن
  await loadActiveGateways()
  // سپس موجودی درگاه‌ها را به‌روزرسانی کن
  await refreshGatewayBalances()
  // سپس ساختار اولیه جدول‌ها را ایجاد کن
  await refreshSmsData()
  // سپس آمار کلی و وضعیت درگاه‌ها
  await loadGatewaysStatus()
})









// نمایش اطلاعات درگاه
const showGatewayDetails = (gateway: any) => {
  try {
    // بررسی چندین فیلد برای کلید API یا اطلاعات ورود
    // اگر api_url نبود، سراغ apiKey، api_key، username و password هم برو
    const apiKey = gateway.api_url || gateway.apiKey || gateway.api_key || '';
    const apiKeyDisplay = apiKey ? apiKey.substring(0, 15) + '...' : 'تنظیم نشده';
    const apiKeyLength = apiKey ? apiKey.length : 0;
    // اگر هیچ کلیدی نبود اما username و password وجود داشت، آن‌ها را هم نمایش بده
    let authInfo = '';
    if (!apiKey && (gateway.username || gateway.password)) {
      authInfo = `یوزرنیم: ${gateway.username || 'ندارد'}\nپسورد: ${gateway.password ? '******' : 'ندارد'}`;
    }
    const details = `اطلاعات درگاه ${gateway.name}:\n\n` +
                   `نام: ${gateway.name}\n` +
                   `نوع: ${gateway.type}\n` +
                   `کلید API: ${apiKeyDisplay}\n` +
                   `طول کلید: ${apiKeyLength} کاراکتر\n` +
                   (authInfo ? authInfo + '\n' : '') +
                   `شماره فرستنده: ${gateway.sender_number || 'تنظیم نشده'}\n` +
                   `فعال: ${gateway.is_active ? 'بله' : 'خیر'}\n` +
                   `بر اساس پترن: ${gateway.pattern_based ? 'بله' : 'خیر'}\n` +
                   `اولویت: ${gateway.priority || 1}\n` +
                   `تاریخ ایجاد: ${gateway.created_at ? new Date(gateway.created_at).toLocaleDateString('fa-IR') : 'نامشخص'}\n` +
                   `آخرین بروزرسانی: ${gateway.updated_at ? new Date(gateway.updated_at).toLocaleDateString('fa-IR') : 'نامشخص'}`
    
    alert(details)
  } catch (error) {
    console.error('خطا در نمایش اطلاعات درگاه:', error)
    alert('خطا در نمایش اطلاعات درگاه: ' + error)
  }
}

// تست ارسال درگاه
const testGatewaySend = async (gateway: any) => {
  alert('این قابلیت در حال حاضر غیرفعال است')
}

// تغییر صفحه (دریافت داده از API با پارامتر page و limit)
const handlePageChange = async (gatewayId: number, newPage: number) => {
  // پیدا کردن ایندکس درگاه
  const gatewayIndex = gatewaySmsData.value.findIndex(g => g.gateway.id === gatewayId)
  if (gatewayIndex !== -1) {
    const gateway = gatewaySmsData.value[gatewayIndex].gateway
    const currentLimit = gatewaySmsData.value[gatewayIndex].pagination?.itemsPerPage || itemsPerPage.value
    let apiUrl = ''
    // تعیین آدرس API بر اساس نوع درگاه
    if (gateway.type === 'ippanel') {
      apiUrl = `/api/ippanel-outbox?page=${newPage}&limit=${currentLimit}`
    } else if (gateway.type === 'farazsms') {
      apiUrl = `/api/farazsms-outbox?page=${newPage}&limit=${currentLimit}`
    } else {
      return
    }
    try {
      // دریافت داده از API
      const response = await $fetch<{status: string, data: any}>(apiUrl, { method: 'GET' })
      if (response.status === 'success' && response.data && response.data.messages) {
        // تبدیل داده‌ها به فرمت جدول
        const smsList = response.data.messages.map((msg: any) => ({
          id: msg.id || Math.random(),
          timestamp: msg.created_at,
          sent_at: msg.sent_at,
          message: msg.message,
          phoneNumber: msg.phone_number,
          phone_number: msg.phone_number,
          status: msg.status,
          type: msg.type,
          cost: msg.cost,
          gateway: gateway.name
        }))
        // به‌روزرسانی smsList و pagination
        gatewaySmsData.value[gatewayIndex].smsList = smsList
        gatewaySmsData.value[gatewayIndex].pagination = {
          currentPage: newPage,
          totalPages: Math.ceil((response.data.pagination?.total || smsList.length) / currentLimit),
          itemsPerPage: currentLimit
        }
      }
    } catch (error) {
      console.error('خطا در دریافت داده‌های صفحه جدید:', error)
    }
  }
}

// تغییر تعداد آیتم در هر صفحه (دریافت داده از API با پارامتر جدید)
const handleItemsPerPageChange = async (newItemsPerPage: number, gatewayId?: number) => {
  // اگر gatewayId مشخص شده، فقط همان درگاه را تغییر بده
  if (gatewayId) {
    const gatewayIndex = gatewaySmsData.value.findIndex(g => g.gateway.id === gatewayId)
    if (gatewayIndex !== -1) {
      gatewaySmsData.value[gatewayIndex].pagination!.itemsPerPage = newItemsPerPage
      await handlePageChange(gatewayId, 1)
    }
  } else {
    // اگر gatewayId مشخص نشده، همه درگاه‌ها را تغییر بده (رفتار قبلی)
    itemsPerPage.value = newItemsPerPage
    for (const gatewayData of gatewaySmsData.value) {
      gatewayData.pagination!.itemsPerPage = newItemsPerPage
      await handlePageChange(gatewayData.gateway.id, 1)
    }
  }
}

// تابع تست داده‌ها
const debugData = () => {
  console.log('🔍 === DEBUG DATA ===')
  console.log('🔍 activeGateways:', activeGateways.value)
  console.log('🔍 gatewaySmsData:', gatewaySmsData.value)
  console.log('🔍 sortedGatewaySmsData:', sortedGatewaySmsData.value)
  console.log('🔍 sortedGatewaySmsData.length:', sortedGatewaySmsData.value.length)
  
  // تست API های outbox
  testOutboxAPIs()
  
  alert(`تست داده‌ها:\n\n` +
        `activeGateways: ${activeGateways.value.length}\n` +
        `gatewaySmsData: ${gatewaySmsData.value.length}\n` +
        `sortedGatewaySmsData: ${sortedGatewaySmsData.value.length}\n\n` +
        `جزئیات در کنسول مرورگر`)
}

// تست API های outbox
const testOutboxAPIs = async () => {
  try {
    console.log('🧪 تست API های outbox...')
    
    // تست IPPanel outbox
    console.log('🧪 تست IPPanel outbox...')
    const ippanelResponse = await $fetch('/api/ippanel-outbox?page=1&limit=5', { method: 'GET' })
    console.log('📡 پاسخ IPPanel outbox:', ippanelResponse)
    
    // تست فراز اس‌ام‌اس outbox
    console.log('🧪 تست فراز اس‌ام‌اس outbox...')
    const farazResponse = await $fetch('/api/farazsms-outbox?page=1&limit=5', { method: 'GET' })
    console.log('📡 پاسخ فراز اس‌ام‌اس outbox:', farazResponse)
    
  } catch (error) {
    console.error('❌ خطا در تست API های outbox:', error)
  }
}

// به‌روزرسانی موجودی درگاه‌ها
const refreshGatewayBalances = async () => {
  try {
    console.log('💰 شروع به‌روزرسانی موجودی درگاه‌ها...')
    
    for (const gateway of activeGateways.value) {
      try {
        // برای ملی پیامک از endpoint جدید استفاده کن
        if (gateway.type === 'meli_payamak') {
          const meliInfo = await getMeliPayamakInfo(gateway.id)
          gateway.balance = meliInfo.remaining_sms // تعداد باقی‌مانده SMS
          gateway.credit = meliInfo.credit // موجودی ریالی
          console.log(`💰 اطلاعات ملی پیامک ${gateway.name}:`, meliInfo)
        } else {
          // برای سایر درگاه‌ها از endpoint قبلی استفاده کن
          const balance = await getGatewayBalance(gateway.id)
          gateway.balance = balance
          console.log(`💰 موجودی درگاه ${gateway.name}:`, balance)
        }
      } catch (error) {
        console.error(`❌ خطا در دریافت موجودی درگاه ${gateway.name}:`, error)
        gateway.balance = 0
        gateway.credit = 0
      }
    }
    
    // به‌روزرسانی gatewaySmsData نیز
    for (const gatewayData of gatewaySmsData.value) {
      const gateway = activeGateways.value.find(g => g.id === gatewayData.gateway.id)
      if (gateway) {
        gatewayData.gateway.balance = gateway.balance
        gatewayData.gateway.credit = gateway.credit
      }
    }
    
    console.log('✅ موجودی درگاه‌ها به‌روزرسانی شد')
  } catch (error) {
    console.error('❌ خطا در به‌روزرسانی موجودی درگاه‌ها:', error)
  }
}

// مشاهده جزئیات SMS
const viewSmsDetails = (sms: SmsRecord) => {
  const details = `جزئیات SMS:\n\n` +
                 `شناسه: ${sms.id}\n` +
                 `زمان ایجاد: ${formatDateTime(sms.created_at || sms.timestamp)}\n` +
                 `زمان ارسال: ${formatDateTime(sms.sent_at || sms.timestamp)}\n` +
                 `شماره تلفن: ${sms.phone_number || sms.phoneNumber}\n` +
                 `نوع: ${sms.type || sms.service || 'عمومی'}\n` +
                 `وضعیت: ${getSmsStatusText(sms.status)}\n` +
                 `وضعیت دقیق: ${sms.status_text || sms.status || 'نامشخص'}\n` +
                 `درگاه: ${sms.gateway || 'نامشخص'}\n` +
                 `هزینه: ${sms.cost ? Number(sms.cost).toLocaleString('fa-IR') + ' ریال' : 'نامشخص'}\n` +
                 `پیام:\n${sms.message}`
  
  alert(details)
}




</script>

<style scoped>
/* استایل‌های سفارشی */
</style> 

