<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-6">
      <h3 class="text-lg font-semibold text-gray-900">تحلیل مبالغ</h3>
      <div class="flex items-center space-x-2 space-x-reverse">
        <select v-model="amountRange" class="text-sm border border-gray-300 rounded-md px-3 py-1">
          <option value="all">همه مبالغ</option>
          <option value="low">کمتر از 10 میلیون</option>
          <option value="medium">10 تا 50 میلیون</option>
          <option value="high">بیش از 50 میلیون</option>
        </select>
        <button @click="refreshData" class="text-blue-600 hover:text-blue-800">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
        </button>
      </div>
    </div>

    <!-- Amount Summary -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
      <div class="bg-blue-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-blue-600 mb-1">{{ formatCurrency(amounts.totalAmount) }}</div>
        <div class="text-sm text-gray-600">کل مبلغ</div>
      </div>
      <div class="bg-green-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-green-600 mb-1">{{ formatCurrency(amounts.avgAmount) }}</div>
        <div class="text-sm text-gray-600">میانگین مبلغ</div>
      </div>
      <div class="bg-yellow-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-yellow-600 mb-1">{{ formatCurrency(amounts.maxAmount) }}</div>
        <div class="text-sm text-gray-600">بالاترین مبلغ</div>
      </div>
      <div class="bg-purple-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-purple-600 mb-1">{{ formatCurrency(amounts.minAmount) }}</div>
        <div class="text-sm text-gray-600">کمترین مبلغ</div>
      </div>
    </div>

    <!-- Amount Ranges -->
    <div class="mb-6">
      <h4 class="text-md font-semibold text-gray-900 mb-4">بازه‌های مبلغ</h4>
      <div class="space-y-4">
        <div v-for="range in amountRanges" :key="range.label" class="bg-gray-50 rounded-lg p-6">
          <div class="flex items-center justify-between mb-3">
            <h5 class="text-md font-semibold text-gray-900">{{ range.label }}</h5>
            <span class="text-sm font-medium text-gray-600">{{ range.count }} درخواست</span>
          </div>
          
          <div class="grid grid-cols-1 md:grid-cols-4 gap-6 text-sm mb-4">
            <div>
              <span class="text-gray-600 block">کل مبلغ</span>
              <span class="font-medium">{{ formatCurrency(range.totalAmount) }}</span>
            </div>
            <div>
              <span class="text-gray-600 block">نرخ تایید</span>
              <span class="font-medium">{{ range.approvalRate }}%</span>
            </div>
            <div>
              <span class="text-gray-600 block">میانگین اقساط</span>
              <span class="font-medium">{{ range.avgInstallments }} قسط</span>
            </div>
            <div>
              <span class="text-gray-600 block">سهم از کل</span>
              <span class="font-medium">{{ range.percentage }}%</span>
            </div>
          </div>

          <div class="mb-3">
            <div class="flex items-center justify-between text-sm mb-1">
              <span class="text-gray-600">سهم از کل مبلغ</span>
              <span class="font-medium">{{ range.percentage }}%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div class="bg-blue-500 h-2 rounded-full" :style="{ width: range.percentage + '%' }"></div>
            </div>
          </div>

          <div class="flex items-center justify-between">
            <button @click="viewRangeDetails(range)" class="text-blue-600 hover:text-blue-800 text-sm font-medium">
              جزئیات بیشتر
            </button>
            <div class="flex items-center" :class="range.trend > 0 ? 'text-green-600' : range.trend < 0 ? 'text-red-600' : 'text-gray-600'">
              <svg v-if="range.trend > 0" class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
              </svg>
              <svg v-else-if="range.trend < 0" class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 17h8m0 0V9m0 8l-8-8-4 4-6-6" />
              </svg>
              <svg v-else class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h8" />
              </svg>
              <span class="text-sm font-medium">
                {{ range.trend > 0 ? '+' : '' }}{{ range.trend }}%
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Installment Analysis -->
    <div class="mb-6">
      <h4 class="text-md font-semibold text-gray-900 mb-4">تحلیل بر اساس تعداد اقساط</h4>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div v-for="installment in installmentAnalysis" :key="installment.count" class="bg-gray-50 rounded-lg p-6">
          <div class="text-center mb-3">
            <div class="text-2xl font-bold text-blue-600 mb-1">{{ installment.count }}</div>
            <div class="text-sm text-gray-600">قسط</div>
          </div>
          
          <div class="space-y-2 text-sm">
            <div class="flex justify-between">
              <span class="text-gray-600">درخواست‌ها:</span>
              <span class="font-medium">{{ installment.requests }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-600">کل مبلغ:</span>
              <span class="font-medium">{{ formatCurrency(installment.totalAmount) }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-600">نرخ تایید:</span>
              <span class="font-medium">{{ installment.approvalRate }}%</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-600">میانگین هر قسط:</span>
              <span class="font-medium">{{ formatCurrency(installment.avgPerInstallment) }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Risk Analysis by Amount -->
    <div>
      <h4 class="text-md font-semibold text-gray-900 mb-4">تحلیل ریسک بر اساس مبلغ</h4>
      <div class="bg-gray-50 rounded-lg p-6">
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div class="text-center p-6 bg-green-100 rounded-lg">
            <div class="text-2xl font-bold text-green-600 mb-1">{{ riskAnalysis.lowRisk.percentage }}%</div>
            <div class="text-sm text-gray-600 mb-2">ریسک پایین</div>
            <div class="text-xs text-gray-500">مبالغ کمتر از 20 میلیون</div>
          </div>
          <div class="text-center p-6 bg-yellow-100 rounded-lg">
            <div class="text-2xl font-bold text-yellow-600 mb-1">{{ riskAnalysis.mediumRisk.percentage }}%</div>
            <div class="text-sm text-gray-600 mb-2">ریسک متوسط</div>
            <div class="text-xs text-gray-500">مبالغ 20 تا 100 میلیون</div>
          </div>
          <div class="text-center p-6 bg-red-100 rounded-lg">
            <div class="text-2xl font-bold text-red-600 mb-1">{{ riskAnalysis.highRisk.percentage }}%</div>
            <div class="text-sm text-gray-600 mb-2">ریسک بالا</div>
            <div class="text-xs text-gray-500">مبالغ بیش از 100 میلیون</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'

// تعریف navigateTo برای Nuxt 3
declare const navigateTo: (to: string) => Promise<void>

const router = useRouter()

interface AmountRange {
  label: string
  count: number
  totalAmount: number
  approvalRate: number
  avgInstallments: number
  percentage: number
  trend: number
}

interface InstallmentData {
  count: number
  requests: number
  totalAmount: number
  approvalRate: number
  avgPerInstallment: number
}

const amountRange = ref('all')

const amounts = ref({
  totalAmount: 125000000000,
  avgAmount: 25000000,
  maxAmount: 500000000,
  minAmount: 5000000
})

const amountRanges = ref<AmountRange[]>([
  {
    label: 'کمتر از 10 میلیون تومان',
    count: 450,
    totalAmount: 3500000000,
    approvalRate: 85,
    avgInstallments: 6,
    percentage: 28,
    trend: 12
  },
  {
    label: '10 تا 30 میلیون تومان',
    count: 320,
    totalAmount: 6400000000,
    approvalRate: 78,
    avgInstallments: 12,
    percentage: 51,
    trend: -5
  },
  {
    label: '30 تا 50 میلیون تومان',
    count: 180,
    totalAmount: 7200000000,
    approvalRate: 65,
    avgInstallments: 18,
    percentage: 14,
    trend: 8
  },
  {
    label: 'بیش از 50 میلیون تومان',
    count: 95,
    totalAmount: 9500000000,
    approvalRate: 55,
    avgInstallments: 24,
    percentage: 7,
    trend: -2
  }
])

const installmentAnalysis = ref<InstallmentData[]>([
  {
    count: 3,
    requests: 245,
    totalAmount: 2450000000,
    approvalRate: 88,
    avgPerInstallment: 3333333
  },
  {
    count: 6,
    requests: 385,
    totalAmount: 7700000000,
    approvalRate: 82,
    avgPerInstallment: 3333333
  },
  {
    count: 12,
    requests: 298,
    totalAmount: 11920000000,
    approvalRate: 75,
    avgPerInstallment: 3333333
  },
  {
    count: 18,
    requests: 156,
    totalAmount: 9360000000,
    approvalRate: 68,
    avgPerInstallment: 3333333
  },
  {
    count: 24,
    requests: 89,
    totalAmount: 8544000000,
    approvalRate: 62,
    avgPerInstallment: 4000000
  },
  {
    count: 36,
    requests: 45,
    totalAmount: 6750000000,
    approvalRate: 55,
    avgPerInstallment: 4166667
  }
])

const riskAnalysis = ref({
  lowRisk: {
    percentage: 65
  },
  mediumRisk: {
    percentage: 28
  },
  highRisk: {
    percentage: 7
  }
})

const formatCurrency = (amount: number): string => {
  return new Intl.NumberFormat('fa-IR', {
    style: 'currency',
    currency: 'IRR',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount)
}

const viewRangeDetails = (range: AmountRange) => {
  router.push(`/admin/finance/installment-payments/amounts/${range.label}`)
}

const refreshData = () => {
  // Refresh data
}

onMounted(() => {
  refreshData()
})
</script>
