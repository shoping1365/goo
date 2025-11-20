<template>
  <div class="bg-white rounded-lg shadow-lg p-6">
    <div class="flex items-center justify-between mb-6">
      <div class="flex items-center space-x-3 space-x-reverse">
        <div class="w-12 h-12 bg-blue-600 rounded-lg flex items-center justify-center">
          <span class="text-white font-bold text-lg">سامان</span>
        </div>
        <div>
          <h3 class="text-lg font-semibold text-gray-900">تراکنش‌های درگاه سامان</h3>
          <p class="text-sm text-gray-500">مشاهده و مدیریت تراکنش‌های پرداخت بانک سامان</p>
        </div>
      </div>
      <div class="flex items-center space-x-2 space-x-reverse">
        <button
          :disabled="loading"
          class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:opacity-50"
          @click="refreshTransactions"
        >
          <span v-if="loading" class="i-heroicons-arrow-path animate-spin mr-2"></span>
          {{ loading ? 'در حال بارگذاری...' : 'بروزرسانی' }}
        </button>
      </div>
    </div>

    <!-- فیلترها -->
    <div class="mb-6 p-6 bg-gray-50 rounded-lg">
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
          <select
            v-model="filters.status"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            @change="fetchTransactions"
          >
            <option value="">همه وضعیت‌ها</option>
            <option value="pending">در انتظار</option>
            <option value="success">موفق</option>
            <option value="failed">ناموفق</option>
            <option value="refunded">بازگشت وجه</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تعداد در صفحه</label>
          <select
            v-model="filters.limit"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            @change="fetchTransactions"
          >
            <option value="10">10</option>
            <option value="20">20</option>
            <option value="50">50</option>
            <option value="100">100</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">جستجو</label>
          <input
            v-model="filters.search"
            type="text"
            placeholder="جستجو در تراکنش‌ها..."
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            @input="debounceSearch"
          />
        </div>
      </div>
    </div>

    <!-- جدول تراکنش‌ها -->
    <div class="overflow-x-auto">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              شناسه تراکنش
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              مبلغ
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              وضعیت
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              تاریخ
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              عملیات
            </th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-if="loading" class="text-center">
            <td colspan="5" class="px-6 py-4">
              <div class="flex justify-center">
                <div class="i-heroicons-arrow-path animate-spin text-2xl text-blue-600"></div>
              </div>
            </td>
          </tr>
          
          <tr v-else-if="transactions.length === 0" class="text-center">
            <td colspan="5" class="px-6 py-4 text-gray-500">
              هیچ تراکنشی یافت نشد
            </td>
          </tr>
          
          <tr v-for="transaction in transactions" :key="transaction.id" class="hover:bg-gray-50">
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
              {{ transaction.transaction_id }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
              {{ formatCurrency(transaction.amount) }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span
                class="inline-flex px-2 py-1 text-xs font-semibold rounded-full"
                :class="getStatusClass(transaction.status)"
              >
                {{ getStatusText(transaction.status) }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
              {{ formatDate(transaction.created_at) }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
              <div class="flex items-center space-x-2 space-x-reverse">
                <button
                  class="text-blue-600 hover:text-blue-900"
                  @click="viewTransaction(transaction)"
                >
                  مشاهده
                </button>
                <button
                  v-if="transaction.status === 'success'"
                  class="text-red-600 hover:text-red-900"
                  @click="refundTransaction(transaction)"
                >
                  بازگشت وجه
                </button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- صفحه‌بندی -->
    <div v-if="total > 0" class="mt-6 flex items-center justify-between">
      <div class="text-sm text-gray-700">
        نمایش {{ (currentPage - 1) * filters.limit + 1 }} تا {{ Math.min(currentPage * filters.limit, total) }} از {{ total }} تراکنش
      </div>
      <div class="flex items-center space-x-2 space-x-reverse">
        <button
          :disabled="currentPage === 1"
          class="px-3 py-2 text-sm border border-gray-300 rounded-lg hover:bg-gray-50 disabled:opacity-50"
          @click="changePage(currentPage - 1)"
        >
          قبلی
        </button>
        
        <span class="px-3 py-2 text-sm text-gray-700">
          صفحه {{ currentPage }} از {{ totalPages }}
        </span>
        
        <button
          :disabled="currentPage === totalPages"
          class="px-3 py-2 text-sm border border-gray-300 rounded-lg hover:bg-gray-50 disabled:opacity-50"
          @click="changePage(currentPage + 1)"
        >
          بعدی
        </button>
      </div>
    </div>

    <!-- مودال مشاهده تراکنش -->
    <div v-if="selectedTransaction" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg p-6 max-w-2xl w-full mx-4 max-h-[90vh] overflow-y-auto">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold">جزئیات تراکنش</h3>
          <button class="text-gray-400 hover:text-gray-600" @click="selectedTransaction = null">
            <span class="i-heroicons-x-mark text-xl"></span>
          </button>
        </div>
        
        <div class="space-y-4">
          <div class="grid grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700">شناسه تراکنش</label>
              <p class="text-sm text-gray-900">{{ selectedTransaction.transaction_id }}</p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">شناسه سفارش</label>
              <p class="text-sm text-gray-900">{{ selectedTransaction.order_id }}</p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">مبلغ</label>
              <p class="text-sm text-gray-900">{{ formatCurrency(selectedTransaction.amount) }}</p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">وضعیت</label>
              <span
                class="inline-flex px-2 py-1 text-xs font-semibold rounded-full"
                :class="getStatusClass(selectedTransaction.status)"
              >
                {{ getStatusText(selectedTransaction.status) }}
              </span>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">روش پرداخت</label>
              <p class="text-sm text-gray-900">{{ selectedTransaction.payment_method }}</p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">تاریخ ایجاد</label>
              <p class="text-sm text-gray-900">{{ formatDate(selectedTransaction.created_at) }}</p>
            </div>
          </div>
          
          <div v-if="selectedTransaction.description">
            <label class="block text-sm font-medium text-gray-700">توضیحات</label>
            <p class="text-sm text-gray-900">{{ selectedTransaction.description }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- پیام‌های نتیجه -->
    <div v-if="message" class="mt-4 p-6 rounded-lg" :class="messageType === 'success' ? 'bg-green-50 text-green-800' : 'bg-red-50 text-red-800'">
      {{ message }}
    </div>
  </div>
</template>

<script setup lang="ts">
import { useDebounceFn } from '@vueuse/core'
import { computed, onMounted, ref } from 'vue'

// Type declaration for Nuxt 4 auto-imported functions
declare const definePageMeta: (meta: { layout?: string }) => void

definePageMeta({
  layout: 'admin-main'
})

// متغیرهای reactive
const loading = ref(false)
const transactions = ref([])
const total = ref(0)
const currentPage = ref(1)
const selectedTransaction = ref(null)
const message = ref('')
const messageType = ref<'success' | 'error'>('success')

// فیلترها
const filters = ref({
  status: '',
  limit: 20,
  search: ''
})

// محاسبه تعداد صفحات
const totalPages = computed(() => Math.ceil(total.value / filters.value.limit))

// دریافت تراکنش‌ها
const fetchTransactions = async () => {
  loading.value = true
  try {
    const params = new URLSearchParams({
      page: currentPage.value.toString(),
      limit: filters.value.limit.toString()
    })
    
    if (filters.value.status) {
      params.append('status', filters.value.status)
    }
    
    interface TransactionsData {
      transactions: unknown[]
      total: number
    }
    interface ApiResponse {
      success?: boolean
      data?: TransactionsData
      message?: string
    }
    const response = await $fetch<ApiResponse>(`/api/admin/payment-gateways/saman.transactions?${params}`)
    
    if (response.success && response.data) {
      transactions.value = response.data.transactions
      total.value = response.data.total
    } else {
      message.value = response.message || 'خطا در دریافت تراکنش‌ها'
      messageType.value = 'error'
    }
  } catch (error) {
    const errorMessage = error instanceof Error ? error.message : 'خطای نامشخص'
    console.error('خطا در دریافت تراکنش‌ها:', errorMessage)
    message.value = 'خطا در دریافت تراکنش‌های درگاه سامان'
    messageType.value = 'error'
  } finally {
    loading.value = false
  }
}

// تغییر صفحه
const changePage = (page: number) => {
  if (page >= 1 && page <= totalPages.value) {
    currentPage.value = page
    fetchTransactions()
  }
}

// بروزرسانی تراکنش‌ها
const refreshTransactions = () => {
  currentPage.value = 1
  fetchTransactions()
}

// مشاهده تراکنش
interface Transaction {
  transaction_id: string
  amount: number
  [key: string]: unknown
}
const viewTransaction = (transaction: Transaction) => {
  selectedTransaction.value = transaction
}

// بازگشت وجه
const refundTransaction = async (transaction: Transaction) => {
  if (!confirm('آیا از بازگشت وجه این تراکنش اطمینان دارید؟')) {
    return
  }
  
  try {
    interface ApiResponse {
      success?: boolean
      message?: string
    }
    const response = await $fetch<ApiResponse>('/api/admin/payment-gateways/saman.refund', {
      method: 'POST',
      body: {
        transaction_id: transaction.transaction_id,
        amount: transaction.amount
      }
    })
    
    if (response.success) {
      message.value = 'بازگشت وجه با موفقیت انجام شد'
      messageType.value = 'success'
      fetchTransactions()
    } else {
      message.value = response.message || 'خطا در بازگشت وجه'
      messageType.value = 'error'
    }
  } catch (error) {
    const errorMessage = error instanceof Error ? error.message : 'خطای نامشخص'
    console.error('خطا در بازگشت وجه:', errorMessage)
    message.value = 'خطا در بازگشت وجه تراکنش'
    messageType.value = 'error'
  }
}

// جستجوی تاخیری
const debounceSearch = useDebounceFn(() => {
  currentPage.value = 1
  fetchTransactions()
}, 500)

// فرمت‌بندی مبلغ
const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

// فرمت‌بندی تاریخ
const formatDate = (date: string) => {
  return new Date(date).toLocaleDateString('fa-IR')
}

// کلاس وضعیت
const getStatusClass = (status: string) => {
  switch (status) {
    case 'success':
      return 'bg-green-100 text-green-800'
    case 'failed':
      return 'bg-red-100 text-red-800'
    case 'pending':
      return 'bg-yellow-100 text-yellow-800'
    case 'refunded':
      return 'bg-gray-100 text-gray-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

// متن وضعیت
const getStatusText = (status: string) => {
  switch (status) {
    case 'success':
      return 'موفق'
    case 'failed':
      return 'ناموفق'
    case 'pending':
      return 'در انتظار'
    case 'refunded':
      return 'بازگشت وجه'
    default:
      return status
  }
}

// دریافت تراکنش‌ها در لود صفحه
onMounted(() => {
  fetchTransactions()
})
</script> 
 
 
