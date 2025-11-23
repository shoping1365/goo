<template>
  <div v-if="hasAccess" class="space-y-6">
    <!-- هدر بخش -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between">
        <div>
          <h2 class="text-xl font-bold text-gray-900">اعتبارسنجی خرید اقساطی</h2>
          <p class="text-sm text-gray-500 mt-1">بررسی اعتبار مشتریان و تایید درخواست‌های خرید اقساطی</p>
        </div>
        <button
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-green-600 hover:bg-green-700"
          @click="openBulkCreditCheck"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
          </svg>
          اعتبارسنجی گروهی
        </button>
      </div>
    </div>

    <!-- فرم اعتبارسنجی -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-medium text-gray-900 mb-4">اعتبارسنجی دستی</h3>
      
      <form class="space-y-6" @submit.prevent="submitCreditCheck">
        <!-- اطلاعات مشتری -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">کد ملی</label>
            <input
              v-model="creditCheck.nationalId"
              type="text"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="کد ملی مشتری"
              maxlength="10"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">شماره موبایل</label>
            <input
              v-model="creditCheck.mobile"
              type="text"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="شماره موبایل"
            />
          </div>
        </div>

        <!-- اطلاعات درخواست -->
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مبلغ درخواستی</label>
            <input
              v-model="creditCheck.amount"
              type="number"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="مبلغ به تومان"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تعداد اقساط</label>
            <select
              v-model="creditCheck.installmentCount"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            >
              <option value="">انتخاب کنید</option>
              <option value="3">3 قسط</option>
              <option value="6">6 قسط</option>
              <option value="12">12 قسط</option>
              <option value="24">24 قسط</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">محصول</label>
            <select
              v-model="creditCheck.productId"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            >
              <option value="">انتخاب محصول</option>
              <option v-for="product in products" :key="product.id" :value="product.id">
                {{ product.name }}
              </option>
            </select>
          </div>
        </div>

        <!-- دکمه‌های عملیات -->
        <div class="flex justify-end space-x-3 space-x-reverse">
          <button
            type="button"
            class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500"
            @click="checkCredit"
          >
            بررسی اعتبار
          </button>
          
          <button
            type="submit"
            class="px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500"
          >
            تایید و ثبت
          </button>
        </div>
      </form>
    </div>

    <!-- نتایج اعتبارسنجی -->
    <div v-if="creditResult" class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-medium text-gray-900 mb-4">نتیجه اعتبارسنجی</h3>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <!-- اطلاعات اعتباری -->
        <div class="space-y-4">
          <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <span class="text-sm font-medium text-gray-700">امتیاز اعتباری:</span>
            <span class="text-lg font-bold" :class="getCreditScoreColor(creditResult.score)">
              {{ creditResult.score }}
            </span>
          </div>
          
          <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <span class="text-sm font-medium text-gray-700">وضعیت:</span>
            <span class="px-2 py-1 text-xs font-medium rounded-full" :class="getStatusClass(creditResult.status)">
              {{ creditResult.status }}
            </span>
          </div>
          
          <div class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
            <span class="text-sm font-medium text-gray-700">حداکثر مبلغ:</span>
            <span class="text-lg font-bold text-green-600">{{ formatCurrency(creditResult.maxAmount) }}</span>
          </div>
        </div>
        
        <!-- جزئیات -->
        <div class="space-y-4">
          <div class="p-3 bg-blue-50 rounded-lg">
            <h4 class="text-sm font-medium text-blue-900 mb-2">تاریخچه اعتباری</h4>
            <ul class="text-sm text-blue-800 space-y-1">
              <li>• تعداد وام‌های قبلی: {{ creditResult.history.totalLoans }}</li>
              <li>• وام‌های معوق: {{ creditResult.history.overdueLoans }}</li>
              <li>• مدت زمان عضویت: {{ creditResult.history.membershipDuration }} ماه</li>
            </ul>
          </div>
          
          <div class="p-3 bg-yellow-50 rounded-lg">
            <h4 class="text-sm font-medium text-yellow-900 mb-2">توصیه‌ها</h4>
            <ul class="text-sm text-yellow-800 space-y-1">
              <li v-for="recommendation in creditResult.recommendations" :key="recommendation">
                • {{ recommendation }}
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>

    <!-- لیست درخواست‌های در انتظار -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-4">
        <h3 class="text-lg font-medium text-gray-900">درخواست‌های در انتظار بررسی</h3>
        <button
          class="text-sm text-blue-600 hover:text-blue-700"
          @click="refreshPendingRequests"
        >
          بروزرسانی
        </button>
      </div>
      
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مشتری</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">محصول</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مبلغ</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ درخواست</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="request in pendingRequests" :key="request.id">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm font-medium text-gray-900">{{ request.customerName }}</div>
                <div class="text-sm text-gray-500">{{ request.nationalId }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ request.productName }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ formatCurrency(request.amount) }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ formatDate(request.createdAt) }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button
                  class="text-blue-600 hover:text-blue-900 ml-3"
                  @click="quickCreditCheck(request)"
                >
                  بررسی سریع
                </button>
                <button
                  class="text-green-600 hover:text-green-900"
                  @click="approveRequest(request.id)"
                >
                  تایید
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
</script>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'
import { useAuth } from '~/composables/useAuth'

// احراز هویت
const { user, isAuthenticated } = useAuth()

// بررسی دسترسی admin
const hasAccess = computed(() => {
  if (!isAuthenticated.value) {
    return false
  }

  const userRole = user.value?.role?.toLowerCase() || ''
  const adminRoles = ['admin', 'developer']
  return adminRoles.includes(userRole)
})

// بررسی احراز هویت و دسترسی admin - نمایش 404 در صورت عدم دسترسی
const checkAuth = async () => {
  if (!hasAccess.value) {
    await navigateTo('/404', { external: false })
  }
}

// بررسی احراز هویت هنگام تغییر وضعیت احراز هویت
watch([isAuthenticated, hasAccess], async () => {
  if (!hasAccess.value) {
    await checkAuth()
  }
})

// تعریف interface ها
interface CreditCheck {
  nationalId: string
  mobile: string
  amount: number
  installmentCount: number
  productId: string
}

interface CreditResult {
  score: number
  status: string
  maxAmount: number
  history: {
    totalLoans: number
    overdueLoans: number
    membershipDuration: number
  }
  recommendations: string[]
}

interface PendingRequest {
  id: string
  customerName: string
  nationalId: string
  productName: string
  amount: number
  createdAt: string
}

// داده‌های reactive
const creditCheck = ref<CreditCheck>({
  nationalId: '',
  mobile: '',
  amount: 0,
  installmentCount: 0,
  productId: ''
})

const creditResult = ref<CreditResult | null>(null)
const pendingRequests = ref<PendingRequest[]>([])
const products = ref([])

// متدها
const checkCredit = async () => {
  try {
    const response = await $fetch<{ data?: CreditResult }>('/api/admin/installment-payments/credit-check', {
      method: 'POST',
      body: creditCheck.value
    })
    const data = response.data
    if (data) {
      creditResult.value = data
    }
  } catch (error) {
    console.error('خطا در بررسی اعتبار:', error)
  }
}

const submitCreditCheck = async () => {
  try {
    await $fetch('/api/admin/installment-payments/approve', {
      method: 'POST',
      body: {
        ...creditCheck.value,
        creditResult: creditResult.value
      }
    })
    // پاک کردن فرم
    creditCheck.value = {
      nationalId: '',
      mobile: '',
      amount: 0,
      installmentCount: 0,
      productId: ''
    }
    creditResult.value = null
  } catch (error) {
    console.error('خطا در ثبت اعتبارسنجی:', error)
  }
}

const openBulkCreditCheck = () => {

}

const refreshPendingRequests = async () => {
  try {
    const response = await $fetch<{ data?: PendingRequest[] }>('/api/admin/installment-payments/pending', {
      method: 'GET'
    })
    const data = response.data
    pendingRequests.value = data || []
  } catch (error) {
    console.error('خطا در دریافت درخواست‌های در انتظار:', error)
  }
}

const quickCreditCheck = async (request: PendingRequest) => {
  creditCheck.value.nationalId = request.nationalId
  creditCheck.value.amount = request.amount
  await checkCredit()
}

const approveRequest = async (requestId: string) => {
  try {
    await $fetch(`/api/admin/installment-payments/${requestId}/approve`, {
      method: 'POST'
    })
    await refreshPendingRequests()
  } catch (error) {
    console.error('خطا در تایید درخواست:', error)
  }
}

// Utility functions
const getCreditScoreColor = (score: number) => {
  if (score >= 700) return 'text-green-600'
  if (score >= 500) return 'text-yellow-600'
  return 'text-red-600'
}

const getStatusClass = (status: string) => {
  switch (status) {
    case 'تایید شده':
      return 'bg-green-100 text-green-800'
    case 'نیاز به بررسی بیشتر':
      return 'bg-yellow-100 text-yellow-800'
    case 'رد شده':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

const formatDate = (date: string) => {
  return new Date(date).toLocaleDateString('fa-IR')
}

// دریافت داده‌ها
const fetchProducts = async () => {
  try {
    const response = await $fetch<{ data?: unknown[] }>('/api/admin/products', {
      method: 'GET'
    })
    const data = response.data
    products.value = data || []
  } catch (error) {
    console.error('خطا در دریافت محصولات:', error)
  }
}

onMounted(async () => {
  await checkAuth()
  if (!hasAccess.value) {
    return
  }
  fetchProducts()
  refreshPendingRequests()
})
</script> 
