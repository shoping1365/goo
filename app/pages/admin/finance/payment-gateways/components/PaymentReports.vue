<template>
  <div class="space-y-6">
    <!-- کارت‌های آمار گزارشات -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <!-- تعداد درگاه‌های فعال -->
      <div class="bg-white p-2 rounded-xl shadow-sm">
        <div class="bg-gradient-to-br from-emerald-50 to-emerald-100 border border-emerald-200 rounded-lg p-6">
          <div class="flex items-center justify-between">
            <div class="flex-1">
              <dl>
                <dt class="text-xs font-medium truncate text-emerald-600">درگاه‌های فعال</dt>
                <dd class="text-base font-bold text-emerald-800">{{ activeRealGateways.length }}</dd>
              </dl>
            </div>
            <div class="flex-shrink-0 mr-3">
              <div class="w-8 h-8 rounded-lg flex items-center justify-center bg-gradient-to-br from-emerald-400 to-emerald-500">
                🔗
              </div>
            </div>
          </div>
          <div class="mt-3 flex items-center text-sm">
            📈
            <span class="text-green-600">+2</span>
            <span class="text-emerald-600 mr-2">نسبت به ماه قبل</span>
          </div>
        </div>
      </div>

      <!-- میانگین کارمزد -->
      <div class="bg-white p-2 rounded-xl shadow-sm">
        <div class="bg-gradient-to-br from-orange-50 to-orange-100 border border-orange-200 rounded-lg p-6">
          <div class="flex items-center justify-between">
            <div class="flex-1">
              <dl>
                <dt class="text-xs font-medium truncate text-orange-600">میانگین کارمزد</dt>
                <dd class="text-base font-bold text-orange-800">{{ averageFee }}%</dd>
              </dl>
            </div>
            <div class="flex-shrink-0 mr-3">
              <div class="w-8 h-8 rounded-lg flex items-center justify-center bg-gradient-to-br from-orange-400 to-orange-500">
                💸
              </div>
            </div>
          </div>
          <div class="mt-3 flex items-center text-sm">
            📈
            <span class="text-green-600">+0.3%</span>
            <span class="text-orange-600 mr-2">نسبت به ماه قبل</span>
          </div>
        </div>
      </div>

      <!-- بالاترین مبلغ تراکنش -->
      <div class="bg-white p-2 rounded-xl shadow-sm">
        <div class="bg-gradient-to-br from-pink-50 to-pink-100 border border-pink-200 rounded-lg p-6">
          <div class="flex items-center justify-between">
            <div class="flex-1">
              <dl>
                <dt class="text-xs font-medium truncate text-pink-600">بیشترین تراکنش</dt>
                <dd class="text-base font-bold text-pink-800">{{ formatCurrency(highestTransaction) }}</dd>
              </dl>
            </div>
            <div class="flex-shrink-0 mr-3">
              <div class="w-8 h-8 rounded-lg flex items-center justify-center bg-gradient-to-br from-pink-400 to-pink-500">
                🏆
              </div>
            </div>
          </div>
          <div class="mt-3 flex items-center text-sm">
            📈
            <span class="text-green-600">+15%</span>
            <span class="text-pink-600 mr-2">نسبت به ماه قبل</span>
          </div>
        </div>
      </div>

      <!-- تعداد تراکنش‌های امروز -->
      <div class="bg-white p-2 rounded-xl shadow-sm">
        <div class="bg-gradient-to-br from-indigo-50 to-indigo-100 border border-indigo-200 rounded-lg p-6">
          <div class="flex items-center justify-between">
            <div class="flex-1">
              <dl>
                <dt class="text-xs font-medium truncate text-indigo-600">تراکنش‌های امروز</dt>
                <dd class="text-base font-bold text-indigo-800">{{ todayTransactions }}</dd>
              </dl>
            </div>
            <div class="flex-shrink-0 mr-3">
              <div class="w-8 h-8 rounded-lg flex items-center justify-center bg-gradient-to-br from-indigo-400 to-indigo-500">
                📅
              </div>
            </div>
          </div>
          <div class="mt-3 flex items-center text-sm">
            📈
            <span class="text-green-600">+8</span>
            <span class="text-indigo-600 mr-2">نسبت به دیروز</span>
          </div>
        </div>
      </div>
    </div>

    <!-- هدر بخش - منتقل شده به پایین -->
    <div class="flex items-center justify-between bg-white p-6 rounded-xl shadow-sm">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">گزارشات پرداخت</h3>
        <p class="text-sm text-gray-600">آمار و گزارشات جامع تراکنش‌های پرداخت</p>
      </div>
      <div class="flex items-center">
        <TemplateButton 
          @click="refreshGateways" 
          :disabled="loadingGateways || loadingTransactions"
          bgGradient="bg-gradient-to-r from-gray-100 to-gray-200"
          textColor="text-gray-700"
          borderColor="border border-gray-200"
          hoverClass="hover:from-gray-200 hover:to-gray-300"
          focusClass="focus:ring-2 focus:ring-gray-200 focus:ring-offset-2"
          size="medium"
          :customClass="(loadingGateways || loadingTransactions) ? 'opacity-50 cursor-not-allowed' : ''"
        >
          <svg 
            :class="['w-4 h-4 inline-block mr-1', (loadingGateways || loadingTransactions) ? 'animate-spin' : '']" 
            fill="none" 
            stroke="currentColor" 
            viewBox="0 0 24 24"
          >
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
          </svg>
          {{ (loadingGateways || loadingTransactions) ? 'در حال بروزرسانی...' : 'بروزرسانی' }}
        </TemplateButton>
        <div class="ml-8"></div>
        <select v-model="selectedPeriod" class="px-4 py-3 border border-gray-200 rounded-lg text-sm bg-white">
          <option value="today">امروز</option>
          <option value="week">هفته جاری</option>
          <option value="month">ماه جاری</option>
          <option value="quarter">سه ماهه</option>
          <option value="year">سال جاری</option>
        </select>
        <div class="ml-8"></div>
        <button
            @click="exportReport"
            class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-green-500 to-emerald-600 hover:from-green-600 hover:to-emerald-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105"
        >
          <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
          </svg>
          خروجی
        </button>
      </div>
    </div>

    <!-- نمودار تراکنش‌ها -->
    <div class="bg-white p-6 rounded-xl shadow-sm">
      <div class="flex items-center justify-between mb-6">
        <h4 class="text-lg font-semibold text-gray-900">نمودار تراکنش‌ها</h4>
        <div class="flex items-center">
          <TemplateButton 
            @click="selectedChartType = 'transactions'"
            :bgGradient="selectedChartType === 'transactions' ? 'bg-gradient-to-r from-blue-100 to-blue-200' : 'bg-gradient-to-r from-gray-100 to-gray-200'"
            :textColor="selectedChartType === 'transactions' ? 'text-blue-700' : 'text-gray-700'"
            :borderColor="selectedChartType === 'transactions' ? 'border border-blue-200' : 'border border-gray-200'"
            :hoverClass="selectedChartType === 'transactions' ? 'hover:from-blue-200 hover:to-blue-300' : 'hover:from-gray-200 hover:to-gray-300'"
            :focusClass="selectedChartType === 'transactions' ? 'focus:ring-2 focus:ring-blue-200 focus:ring-offset-2' : 'focus:ring-2 focus:ring-gray-200 focus:ring-offset-2'"
            size="medium"
          >
            تعداد تراکنش‌ها
          </TemplateButton>
          <div class="ml-4"></div>
          <TemplateButton 
            @click="selectedChartType = 'amounts'"
            :bgGradient="selectedChartType === 'amounts' ? 'bg-gradient-to-r from-blue-100 to-blue-200' : 'bg-gradient-to-r from-gray-100 to-gray-200'"
            :textColor="selectedChartType === 'amounts' ? 'text-blue-700' : 'text-gray-700'"
            :borderColor="selectedChartType === 'amounts' ? 'border border-blue-200' : 'border border-gray-200'"
            :hoverClass="selectedChartType === 'amounts' ? 'hover:from-blue-200 hover:to-blue-300' : 'hover:from-gray-200 hover:to-gray-300'"
            :focusClass="selectedChartType === 'amounts' ? 'focus:ring-2 focus:ring-blue-200 focus:ring-offset-2' : 'focus:ring-2 focus:ring-gray-200 focus:ring-offset-2'"
            size="medium"
          >
            مبالغ
          </TemplateButton>
        </div>
      </div>
      
      <!-- نمودار شبیه‌سازی شده -->
      <div class="h-64 bg-gray-50 rounded-lg flex items-center justify-center">
        <div class="text-center">
          📊
          <p class="text-gray-500">نمودار {{ selectedChartType === 'transactions' ? 'تراکنش‌ها' : 'مبالغ' }} در {{ selectedPeriod }}</p>
          <p class="text-sm text-gray-400 mt-1">این بخش به کتابخانه نمودار متصل خواهد شد</p>
        </div>
      </div>
    </div>
      </div>

    <!-- آمار درگاه‌های پرداخت - در انتهای صفحه -->
    <div class="bg-white rounded-xl shadow-sm mt-8">
      <div class="p-6 border-b border-gray-200">
        <h4 class="text-lg font-semibold text-gray-900">آمار درگاه‌های پرداخت</h4>
      </div>
      
      <div class="p-6">
        <div :class="['grid gap-6', gatewayGridCols]">
          <div v-for="gateway in activeRealGateways" :key="gateway.id" class="border border-gray-200 rounded-lg p-6">
            <div class="flex items-center justify-between mb-3">
              <div class="flex items-center">
                <div :class="['w-8 h-8 rounded-full flex items-center justify-center text-white text-sm font-medium', gateway.color]">
                  {{ gateway.icon }}
                </div>
                <span class="mr-3 font-medium text-gray-900">{{ gateway.name }}</span>
              </div>
              <span :class="[
                'px-2 py-1 text-xs font-semibold rounded-full',
                gateway.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'
              ]">
                {{ gateway.status === 'active' ? 'فعال' : 'غیرفعال' }}
              </span>
            </div>
            
            <div class="space-y-2">
              <div class="flex justify-between text-sm">
                <span class="text-gray-600">تراکنش‌ها:</span>
                <span class="font-medium">{{ formatNumber(gateway.total_transactions) }}</span>
              </div>
              <div class="flex justify-between text-sm">
                <span class="text-gray-600">مبلغ کل:</span>
                <span class="font-medium">{{ formatCurrency(gateway.total_revenue) }}</span>
              </div>
              <div class="flex justify-between text-sm">
                <span class="text-gray-600">نرخ موفقیت:</span>
                <span class="font-medium">{{ gateway.successRate }}%</span>
              </div>
              <div class="flex justify-between text-sm">
                <span class="text-gray-600">کارمزد:</span>
                <span class="font-medium">{{ formatCurrency(gateway.fee) }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </template>

<script setup lang="ts">
import { computed, onMounted, onUnmounted, ref } from 'vue'
import TemplateButton from '~/components/common/TemplateButton.vue'
// آیکون‌ها به صورت مستقیم استفاده می‌شوند

// Types - مطابق با GatewayList.vue
interface PaymentGateway {
  id: number
  name: string
  english_name: string
  type: string
  status: string
  icon: string
  color: string
  fee: number
  min_amount: number
  max_amount: number
  today_transactions: number
  today_revenue: number
  total_transactions: number
  total_revenue: number
  is_test_mode: boolean
  successRate?: number
}

// متغیرهای reactive
const selectedPeriod = ref('month')
const selectedChartType = ref('transactions')

// نوع‌های نمودار
const chartTypes = [
  { value: 'transactions', label: 'تعداد تراکنش‌ها' },
  { value: 'amounts', label: 'مبالغ' }
]

// آمار محاسبه شده برای کارت‌های گزارشات
const averageFee = computed(() => {
  if (activeRealGateways.value.length === 0) return 0
  const totalFee = activeRealGateways.value.reduce((sum, gateway) => sum + gateway.fee, 0)
  return Math.round((totalFee / activeRealGateways.value.length) * 10) / 10
})

const highestTransaction = computed(() => {
  if (allTransactions.value.length === 0) return 0
  return Math.max(...allTransactions.value.map(t => t.amount || 0))
})

const todayTransactions = computed(() => {
  const today = new Date().toDateString()
  return allTransactions.value.filter(t => {
    const transactionDate = new Date(t.created_at).toDateString()
    return transactionDate === today
  }).length
})

// تنظیم تعداد ستون‌ها بر اساس تعداد درگاه‌ها
const gatewayGridCols = computed(() => {
  const count = activeRealGateways.value.length
  if (count === 1) return 'grid-cols-1'
  if (count === 2) return 'grid-cols-1 md:grid-cols-2'
  if (count === 3) return 'grid-cols-1 md:grid-cols-3'
  if (count === 4) return 'grid-cols-1 md:grid-cols-2 lg:grid-cols-4'
  if (count === 5) return 'grid-cols-1 md:grid-cols-2 lg:grid-cols-5'
  if (count === 6) return 'grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-6'
  return 'grid-cols-1 md:grid-cols-2 lg:grid-cols-3' // پیش‌فرض
})

// تراکنش‌های واقعی از API
const allTransactions = ref<any[]>([])
const loadingTransactions = ref(false)

// درگاه‌های واقعی از API - دقیقاً همان API داشبورد
const actualGateways = ref<PaymentGateway[]>([])
const loadingGateways = ref(true)

// دریافت درگاه‌های واقعی از API - دقیقاً همان روش داشبورد
const fetchActualGateways = async () => {
  try {
    loadingGateways.value = true
    const response: any = await $fetch('/api/payment-gateways')
    actualGateways.value = response.data || []
    
    console.log('🔍 درگاه‌های دریافت شده از API:', actualGateways.value)
    
    // اگر API در دسترس نباشد، از داده‌های نمونه استفاده کن
    if (!actualGateways.value.length) {
      console.log('⚠️ API در دسترس نیست، استفاده از داده‌های نمونه')
      actualGateways.value = [
        {
          id: 1,
          name: 'زرین‌پال',
          english_name: 'ZarinPal',
          type: 'iranian',
          status: 'active',
          icon: 'زر',
          color: 'bg-green-500',
          fee: 1.5,
          min_amount: 1000,
          max_amount: 50000000,
          today_transactions: 45,
          today_revenue: 12500000,
          total_transactions: 1250,
          total_revenue: 450000000,
          is_test_mode: false
        },
        {
          id: 2,
          name: 'نکست‌پی',
          english_name: 'NextPay',
          type: 'iranian',
          status: 'active',
          icon: 'ن',
          color: 'bg-blue-500',
          fee: 2.0,
          min_amount: 1000,
          max_amount: 100000000,
          today_transactions: 32,
          today_revenue: 8900000,
          total_transactions: 890,
          total_revenue: 320000000,
          is_test_mode: false
        },
        {
          id: 3,
          name: 'کیف پول دیجیتال',
          english_name: 'Digital Wallet',
          type: 'digital',
          status: 'active',
          icon: 'ک',
          color: 'bg-violet-500',
          fee: 0.5,
          min_amount: 1000,
          max_amount: 20000000,
          today_transactions: 28,
          today_revenue: 6500000,
          total_transactions: 650,
          total_revenue: 180000000,
          is_test_mode: false
        }
      ]
    }
  } catch (error) {
    console.error('❌ خطا در دریافت درگاه‌ها:', error)
    // در صورت خطا، از داده‌های نمونه استفاده کن
    actualGateways.value = [
      {
        id: 1,
        name: 'زرین‌پال',
        english_name: 'ZarinPal',
        type: 'iranian',
        status: 'active',
        icon: 'زر',
        color: 'bg-green-500',
        fee: 1.5,
        min_amount: 1000,
        max_amount: 50000000,
        today_transactions: 45,
        today_revenue: 12500000,
        total_transactions: 1250,
        total_revenue: 450000000,
        is_test_mode: false
      },
      {
        id: 2,
        name: 'نکست‌پی',
        english_name: 'NextPay',
        type: 'iranian',
        status: 'active',
        icon: 'ن',
        color: 'bg-blue-500',
        fee: 2.0,
        min_amount: 1000,
        max_amount: 100000000,
        today_transactions: 32,
        today_revenue: 8900000,
        total_transactions: 890,
        total_revenue: 320000000,
        is_test_mode: false
      }
    ]
  } finally {
    loadingGateways.value = false
  }
}

// درگاه‌های فعال که واقعاً در سیستم وجود دارند - دقیقاً همان فیلتر داشبورد
const activeRealGateways = computed(() => {
  return actualGateways.value.filter(gateway => gateway.status === 'active')
})

// دریافت تراکنش‌های واقعی از API
const fetchAllTransactions = async () => {
  try {
    loadingTransactions.value = true
    const response: any = await $fetch('/api/payments/admin/transactions', {
      query: {
        limit: 100 // دریافت 100 تراکنش اخیر
      }
    })
    
    if (response.data) {
      allTransactions.value = response.data
    }
  } catch (error) {
    console.error('خطا در دریافت تراکنش‌ها:', error)
    allTransactions.value = []
  } finally {
    loadingTransactions.value = false
  }
}

// توابع کمکی
const formatNumber = (num: number): string => {
  return new Intl.NumberFormat('fa-IR').format(num)
}

const formatCurrency = (amount: number): string => {
  return new Intl.NumberFormat('fa-IR', {
    style: 'currency',
    currency: 'IRR',
    minimumFractionDigits: 0
  }).format(amount)
}

const formatDate = (dateString: string): string => {
  return new Date(dateString).toLocaleDateString('fa-IR', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

const getStatusText = (status: string): string => {
  const statusMap = {
    success: 'موفق',
    failed: 'ناموفق',
    pending: 'در انتظار',
    refunded: 'بازگشت وجه',
    cancelled: 'لغو شده'
  }
  return statusMap[status] || status
}

const getStatusClass = (status: string): string => {
  switch (status) {
    case 'success':
      return 'bg-green-100 text-green-800'
    case 'failed':
      return 'bg-red-100 text-red-800'
    case 'pending':
      return 'bg-yellow-100 text-yellow-800'
    case 'refunded':
      return 'bg-purple-100 text-purple-800'
    case 'cancelled':
      return 'bg-gray-100 text-gray-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getGatewayColor = (type: string): string => {
  switch (type) {
    case 'zarinpal':
      return 'bg-purple-600'
    case 'mellat':
      return 'bg-blue-600'
    case 'saman':
      return 'bg-green-600'
    case 'parsian':
      return 'bg-orange-600'
    case 'melli':
      return 'bg-red-600'
    default:
      return 'bg-gray-600'
  }
}

const getGatewayIcon = (type: string): string => {
  switch (type) {
    case 'zarinpal':
      return 'زر'
    case 'mellat':
      return 'مل'
    case 'saman':
      return 'سا'
    case 'parsian':
      return 'پا'
    case 'melli':
      return 'مل'
    default:
      return '💳'
  }
}

// توابع عملیات
const exportReport = () => {
  // منطق خروجی اکسل
  console.log('خروجی گزارش برای دوره:', selectedPeriod.value)
}

const viewAllTransactions = () => {
  // هدایت به صفحه همه تراکنش‌ها
  console.log('مشاهده همه تراکنش‌ها')
}

const viewTransaction = (transaction: any) => {
  // نمایش جزئیات تراکنش
  console.log('جزئیات تراکنش:', transaction)
}

// تابع برای دریافت تراکنش‌های یک درگاه خاص
const getGatewayTransactions = (gatewayId: number) => {
  return allTransactions.value.filter(
    (transaction) => transaction.gateway_id === gatewayId
  )
}

// تراکنش‌های اخیر برای نمایش در جدول بالا
const recentTransactions = computed(() => {
  return allTransactions.value.slice(0, 10) // 10 تراکنش اخیر
})

// بارگذاری درگاه‌ها و تراکنش‌ها هنگام لود کامپوننت
let refreshInterval: ReturnType<typeof setInterval> | null = null

onMounted(() => {
  fetchActualGateways()
  fetchAllTransactions()
  
  // شروع بررسی تغییرات هر 30 ثانیه
  refreshInterval = setInterval(() => {
    fetchActualGateways()
    fetchAllTransactions()
  }, 30000)
})

onUnmounted(() => {
  if (refreshInterval) {
    clearInterval(refreshInterval)
    refreshInterval = null
  }
})

// تابع بروزرسانی درگاه‌ها
const refreshGateways = () => {
  fetchActualGateways()
  fetchAllTransactions()
}
</script> 
