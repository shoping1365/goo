<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-6">
      <h3 class="text-lg font-semibold text-gray-900">تحلیل مشتریان</h3>
      <div class="flex items-center space-x-2 space-x-reverse">
        <select v-model="customerType" class="text-sm border border-gray-300 rounded-md px-3 py-1">
          <option value="all">همه مشتریان</option>
          <option value="new">مشتریان جدید</option>
          <option value="returning">مشتریان تکراری</option>
          <option value="vip">مشتریان VIP</option>
        </select>
        <button class="text-blue-600 hover:text-blue-800" @click="refreshData">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
        </button>
      </div>
    </div>

    <!-- Customer Demographics -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
      <div class="bg-blue-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-blue-600 mb-1">{{ demographics.totalCustomers }}</div>
        <div class="text-sm text-gray-600">کل مشتریان</div>
      </div>
      <div class="bg-green-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-green-600 mb-1">{{ demographics.avgAge }}</div>
        <div class="text-sm text-gray-600">میانگین سن</div>
      </div>
      <div class="bg-yellow-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-yellow-600 mb-1">{{ demographics.avgIncome }}</div>
        <div class="text-sm text-gray-600">میانگین درآمد</div>
      </div>
      <div class="bg-purple-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-purple-600 mb-1">{{ demographics.repeatRate }}%</div>
        <div class="text-sm text-gray-600">نرخ تکرار</div>
      </div>
    </div>

    <!-- Customer List -->
    <div class="space-y-4">
      <div v-for="customer in customers" :key="customer.id" class="bg-gray-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div class="flex items-center">
            <div class="w-12 h-12 bg-blue-100 rounded-full flex items-center justify-center mr-4">
              <span class="text-blue-600 font-semibold">{{ customer.name.charAt(0) }}</span>
            </div>
            <div>
              <h5 class="text-lg font-semibold text-gray-900">{{ customer.name }}</h5>
              <p class="text-sm text-gray-600">{{ customer.phone }} | {{ customer.nationalId }}</p>
              <div class="flex items-center space-x-2 space-x-reverse mt-1">
                <span class="px-2 py-1 text-xs rounded-full" :class="getCustomerTypeClass(customer.type)">
                  {{ getCustomerTypeLabel(customer.type) }}
                </span>
                <span class="px-2 py-1 text-xs bg-gray-200 text-gray-700 rounded-full">
                  امتیاز: {{ customer.creditScore }}
                </span>
              </div>
            </div>
          </div>
          <div class="text-right">
            <div class="text-lg font-bold text-gray-900">{{ customer.totalRequests }}</div>
            <div class="text-sm text-gray-500">درخواست</div>
          </div>
        </div>
        
        <div class="mt-4 grid grid-cols-1 md:grid-cols-4 gap-6 text-sm">
          <div>
            <span class="text-gray-600 block">کل مبلغ</span>
            <span class="font-medium">{{ formatCurrency(customer.totalAmount) }}</span>
          </div>
          <div>
            <span class="text-gray-600 block">نرخ تایید</span>
            <span class="font-medium">{{ customer.approvalRate }}%</span>
          </div>
          <div>
            <span class="text-gray-600 block">آخرین درخواست</span>
            <span class="font-medium">{{ customer.lastRequest }}</span>
          </div>
          <div>
            <span class="text-gray-600 block">وضعیت</span>
            <span class="font-medium" :class="customer.status === 'active' ? 'text-green-600' : 'text-red-600'">
              {{ customer.status === 'active' ? 'فعال' : 'غیرفعال' }}
            </span>
          </div>
        </div>

        <div class="mt-3 flex items-center space-x-4 space-x-reverse">
          <button class="text-blue-600 hover:text-blue-800 text-sm font-medium" @click="viewCustomerDetails(customer)">
            مشاهده جزئیات
          </button>
          <button class="text-green-600 hover:text-green-800 text-sm font-medium" @click="customerHistory(customer)">
            تاریخچه
          </button>
          <button class="text-purple-600 hover:text-purple-800 text-sm font-medium" @click="creditCheck(customer)">
            بررسی اعتبار
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'

// تعریف navigateTo برای Nuxt 3
declare const _navigateTo: (to: string) => Promise<void>

const router = useRouter()

interface Customer {
  id: number
  name: string
  phone: string
  nationalId: string
  type: 'new' | 'returning' | 'vip'
  creditScore: number
  totalRequests: number
  totalAmount: number
  approvalRate: number
  lastRequest: string
  status: 'active' | 'inactive'
}

const customerType = ref('all')
const demographics = ref({
  totalCustomers: 1250,
  avgAge: 32,
  avgIncome: '25M',
  repeatRate: 68
})

const customers = ref<Customer[]>([
  {
    id: 1,
    name: 'علی احمدی',
    phone: '09123456789',
    nationalId: '1234567890',
    type: 'vip',
    creditScore: 85,
    totalRequests: 5,
    totalAmount: 150000000,
    approvalRate: 100,
    lastRequest: '1403/08/15',
    status: 'active'
  }
])

const formatCurrency = (amount: number): string => {
  return new Intl.NumberFormat('fa-IR', {
    style: 'currency',
    currency: 'IRR',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount)
}

const getCustomerTypeClass = (type: string): string => {
  switch (type) {
    case 'new': return 'bg-blue-100 text-blue-800'
    case 'returning': return 'bg-green-100 text-green-800'
    case 'vip': return 'bg-purple-100 text-purple-800'
    default: return 'bg-gray-100 text-gray-800'
  }
}

const getCustomerTypeLabel = (type: string): string => {
  switch (type) {
    case 'new': return 'جدید'
    case 'returning': return 'تکراری'
    case 'vip': return 'VIP'
    default: return 'عادی'
  }
}

const viewCustomerDetails = (customer: Customer) => {
  router.push(`/admin/users/${customer.id}`)
}

const customerHistory = (customer: Customer) => {
  router.push(`/admin/finance/installment-payments/customers/${customer.id}/history`)
}

const creditCheck = (customer: Customer) => {
  router.push(`/admin/finance/installment-payments/customers/${customer.id}/credit`)
}

const refreshData = () => {
  // Refresh logic
}

onMounted(() => {
  refreshData()
})
</script>
