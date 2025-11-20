<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex justify-between items-center">
        <div>
          <h2 class="text-2xl font-bold text-gray-900">داشبورد گیفت کارت</h2>
          <p class="text-gray-600 mt-1">نمای کلی از وضعیت و عملکرد سیستم گیفت کارت</p>
        </div>
      </div>
    </div>



    <!-- نمودارها و آمار -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- نمودار وضعیت کارت‌ها -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-medium text-gray-900 mb-4">وضعیت کارت‌ها</h3>
        <div class="h-[350px]">
          <canvas ref="statusChart" class="w-full h-full"></canvas>
        </div>
        <div class="mt-4 grid grid-cols-2 gap-6">
          <div v-for="status in chartData.statuses" :key="status.label" class="flex items-center">
            <div :class="status.color" class="w-3 h-3 rounded-full ml-2"></div>
            <span class="text-sm text-gray-600">{{ status.label }}</span>
            <span class="text-sm font-medium text-gray-900 mr-auto">{{ status.value }}</span>
          </div>
        </div>
      </div>

      <!-- نمودار روند ایجاد کارت‌ها -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-medium text-gray-900 mb-4">روند ایجاد کارت‌ها (30 روز گذشته)</h3>
        <div class="h-[350px]">
          <canvas ref="trendChart" class="w-full h-full"></canvas>
        </div>
      </div>
    </div>

    <!-- آمار مالی -->
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <!-- نمودار توزیع مبلغ -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-medium text-gray-900 mb-4">توزیع مبلغ کارت‌ها</h3>
        <div class="h-[350px]">
          <canvas ref="amountChart" class="w-full h-full"></canvas>
        </div>
      </div>

      <!-- آمار استفاده -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-medium text-gray-900 mb-4">آمار استفاده</h3>
        <div class="space-y-4">
          <div class="flex justify-between items-center">
            <span class="text-sm text-gray-600">کارت‌های استفاده شده</span>
            <span class="text-sm font-medium text-gray-900">{{ formatNumber(stats.usedCards) }}</span>
          </div>
          <div class="w-full bg-gray-200 rounded-full h-2">
            <div class="bg-blue-600 h-2 rounded-full" :style="{ width: stats.usageRate + '%' }"></div>
          </div>
          <div class="flex justify-between items-center">
            <span class="text-sm text-gray-600">نرخ استفاده</span>
            <span class="text-sm font-medium text-gray-900">{{ stats.usageRate }}%</span>
          </div>
          
          <div class="border-t pt-4 mt-4">
            <div class="flex justify-between items-center">
              <span class="text-sm text-gray-600">میانگین مبلغ استفاده</span>
              <span class="text-sm font-medium text-gray-900">{{ formatCurrency(stats.avgUsageAmount) }}</span>
            </div>
          </div>
        </div>
      </div>

      <!-- آمار تحویل -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-medium text-gray-900 mb-4">آمار تحویل</h3>
        <div class="space-y-4">
          <div class="flex justify-between items-center">
            <span class="text-sm text-gray-600">تحویل شده</span>
            <span class="text-sm font-medium text-gray-900">{{ formatNumber(stats.deliveredCards) }}</span>
          </div>
          <div class="w-full bg-gray-200 rounded-full h-2">
            <div class="bg-green-600 h-2 rounded-full" :style="{ width: stats.deliveryRate + '%' }"></div>
          </div>
          <div class="flex justify-between items-center">
            <span class="text-sm text-gray-600">نرخ تحویل</span>
            <span class="text-sm font-medium text-gray-900">{{ stats.deliveryRate }}%</span>
          </div>
          
          <div class="border-t pt-4 mt-4 space-y-2">
            <div class="flex justify-between items-center">
              <span class="text-sm text-gray-600">ایمیل</span>
              <span class="text-sm font-medium text-gray-900">{{ stats.emailDelivery }}%</span>
            </div>
            <div class="flex justify-between items-center">
              <span class="text-sm text-gray-600">پیامک</span>
              <span class="text-sm font-medium text-gray-900">{{ stats.smsDelivery }}%</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- جدول آخرین فعالیت‌ها -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex justify-between items-center mb-4">
        <h3 class="text-lg font-medium text-gray-900">آخرین فعالیت‌ها</h3>
        <button 
          class="text-sm text-blue-600 hover:text-blue-700 font-medium"
          @click="viewAllActivities"
        >
          مشاهده همه
        </button>
      </div>
      
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کارت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کاربر</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مبلغ</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="activity in recentActivities" :key="activity.id">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="">
                    <div class="h-10 w-10 rounded-lg bg-gray-200 flex items-center justify-center">
                      <span class="text-sm font-medium text-gray-700">{{ activity.cardCode.substring(0, 4) }}</span>
                    </div>
                  </div>
                  <div class="mr-4">
                    <div class="text-sm font-medium text-gray-900">{{ activity.cardCode }}</div>
                    <div class="text-sm text-gray-500">{{ activity.cardType }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="text-sm text-gray-900">{{ activity.action }}</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-900">{{ activity.userName }}</div>
                <div class="text-sm text-gray-500">{{ activity.userEmail }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="text-sm font-medium text-gray-900">{{ formatCurrency(activity.amount) }}</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="text-sm text-gray-500">{{ formatDate(activity.date) }}</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusClass(activity.status)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                  {{ getStatusLabel(activity.status) }}
                </span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, nextTick } from 'vue'
import Chart from 'chart.js/auto'

// Props
const props = defineProps<{
  giftCards: any[]
}>()

// Emits
const emit = defineEmits<{
  'view-activities': []
}>()

// Reactive data
const statusChart = ref<HTMLCanvasElement>()
const trendChart = ref<HTMLCanvasElement>()
const amountChart = ref<HTMLCanvasElement>()

const stats = reactive({
  totalCards: 1250,
  newCardsToday: 15,
  activeCards: 980,
  activePercentage: 78.4,
  expiringSoon: 45,
  totalValue: 125000000,
  valueGrowth: 12.5,
  usedCards: 320,
  usageRate: 25.6,
  avgUsageAmount: 85000,
  deliveredCards: 1100,
  deliveryRate: 88.0,
  emailDelivery: 65,
  smsDelivery: 35
})

const chartData = reactive({
  statuses: [
    { label: 'فعال', value: 980, color: 'bg-green-500' },
    { label: 'استفاده شده', value: 320, color: 'bg-blue-500' },
    { label: 'منقضی شده', value: 150, color: 'bg-red-500' },
    { label: 'معلق', value: 25, color: 'bg-yellow-500' }
  ]
})

const recentActivities = ref([
  {
    id: 1,
    cardCode: 'GC-2024-001',
    cardType: 'دیجیتال',
    action: 'ایجاد کارت',
    userName: 'علی احمدی',
    userEmail: 'ali@example.com',
    amount: 100000,
    date: '2024-01-15T10:30:00',
    status: 'success'
  },
  {
    id: 2,
    cardCode: 'GC-2024-002',
    cardType: 'فیزیکی',
    action: 'استفاده',
    userName: 'فاطمه محمدی',
    userEmail: 'fateme@example.com',
    amount: 75000,
    date: '2024-01-15T09:15:00',
    status: 'success'
  },
  {
    id: 3,
    cardCode: 'GC-2024-003',
    cardType: 'مجازی',
    action: 'تحویل',
    userName: 'محمد رضایی',
    userEmail: 'mohammad@example.com',
    amount: 150000,
    date: '2024-01-15T08:45:00',
    status: 'pending'
  },
  {
    id: 4,
    cardCode: 'GC-2024-004',
    cardType: 'دیجیتال',
    action: 'لغو',
    userName: 'زهرا کریمی',
    userEmail: 'zahra@example.com',
    amount: 50000,
    date: '2024-01-15T07:20:00',
    status: 'cancelled'
  },
  {
    id: 5,
    cardCode: 'GC-2024-005',
    cardType: 'فیزیکی',
    action: 'انقضا',
    userName: 'سیستم',
    userEmail: 'system@example.com',
    amount: 200000,
    date: '2024-01-15T06:00:00',
    status: 'expired'
  }
])

// Methods
const refreshData = () => {
  console.log('به‌روزرسانی داده‌های داشبورد')
  // اینجا می‌توانید API call برای به‌روزرسانی داده‌ها اضافه کنید
}

const viewAllActivities = () => {
  emit('view-activities')
}

const formatNumber = (num: number) => {
  return new Intl.NumberFormat('fa-IR').format(num)
}

const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleDateString('fa-IR', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

const getStatusClass = (status: string) => {
  switch (status) {
    case 'success':
      return 'bg-green-100 text-green-800'
    case 'pending':
      return 'bg-yellow-100 text-yellow-800'
    case 'cancelled':
      return 'bg-red-100 text-red-800'
    case 'expired':
      return 'bg-gray-100 text-gray-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getStatusLabel = (status: string) => {
  switch (status) {
    case 'success':
      return 'موفق'
    case 'pending':
      return 'در انتظار'
    case 'cancelled':
      return 'لغو شده'
    case 'expired':
      return 'منقضی'
    default:
      return 'نامشخص'
  }
}

const initCharts = async () => {
  await nextTick()
  
  // نمودار وضعیت کارت‌ها (Doughnut)
  if (statusChart.value) {
    new Chart(statusChart.value, {
      type: 'doughnut',
      data: {
        labels: chartData.statuses.map(s => s.label),
        datasets: [{
          data: chartData.statuses.map(s => s.value),
          backgroundColor: ['#10B981', '#3B82F6', '#EF4444', '#F59E0B'],
          borderWidth: 0
        }]
      },
      options: {
        responsive: true,
        maintainAspectRatio: false,
        plugins: {
          legend: {
            display: false
          }
        }
      }
    })
  }
  
  // نمودار روند (Line)
  if (trendChart.value) {
    const labels = Array.from({ length: 30 }, (_, i) => {
      const date = new Date()
      date.setDate(date.getDate() - (29 - i))
      return date.toLocaleDateString('fa-IR', { month: 'short', day: 'numeric' })
    })
    
    new Chart(trendChart.value, {
      type: 'line',
      data: {
        labels: labels,
        datasets: [{
          label: 'کارت‌های جدید',
          data: Array.from({ length: 30 }, () => Math.floor(Math.random() * 20) + 5),
          borderColor: '#3B82F6',
          backgroundColor: 'rgba(59, 130, 246, 0.1)',
          tension: 0.4,
          fill: true
        }]
      },
      options: {
        responsive: true,
        maintainAspectRatio: false,
        plugins: {
          legend: {
            display: false
          }
        },
        scales: {
          y: {
            beginAtZero: true,
            ticks: {
              stepSize: 5
            }
          }
        }
      }
    })
  }
  
  // نمودار توزیع مبلغ (Bar)
  if (amountChart.value) {
    new Chart(amountChart.value, {
      type: 'bar',
      data: {
        labels: ['کمتر از 50K', '50K-100K', '100K-200K', '200K-500K', 'بیش از 500K'],
        datasets: [{
          label: 'تعداد کارت‌ها',
          data: [150, 300, 400, 250, 150],
          backgroundColor: '#8B5CF6',
          borderRadius: 4
        }]
      },
      options: {
        responsive: true,
        maintainAspectRatio: false,
        plugins: {
          legend: {
            display: false
          }
        },
        scales: {
          y: {
            beginAtZero: true
          }
        }
      }
    })
  }
}

// Lifecycle
onMounted(() => {
  initCharts()
  console.log('Gift card dashboard component mounted')
})
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 
