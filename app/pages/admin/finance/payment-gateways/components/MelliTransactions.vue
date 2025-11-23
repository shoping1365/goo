<template>
  <div v-if="hasAccess" class="bg-white rounded-lg shadow-md">
    <!-- Header -->
    <div class="px-6 py-4 border-b border-gray-200">
      <div class="flex items-center justify-between">
        <div>
          <h3 class="text-lg font-medium text-gray-900">تراکنش‌های درگاه ملی</h3>
          <p class="text-sm text-gray-500">لیست تمام تراکنش‌های انجام شده از طریق درگاه ملی</p>
        </div>
        <button
          :disabled="loading"
          class="px-4 py-2 text-sm font-medium text-blue-600 bg-blue-50 border border-blue-200 rounded-md hover:bg-blue-100 focus:outline-none focus:ring-2 focus:ring-blue-500 disabled:opacity-50"
          @click="loadTransactions"
        >
          <Icon v-if="loading" name="ph:spinner" class="w-4 h-4 animate-spin" />
          <Icon v-else name="ph:arrow-clockwise" class="w-4 h-4" />
          <span class="mr-2">{{ loading ? 'در حال بارگذاری...' : 'بارگذاری مجدد' }}</span>
        </button>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="flex items-center justify-center py-12">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
      <span class="mr-3 text-gray-600">در حال بارگذاری تراکنش‌ها...</span>
    </div>

    <!-- Transactions Table -->
    <div v-else-if="transactions.length > 0" class="overflow-x-auto">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              شناسه تراکنش
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              شناسه سفارش
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
          <tr v-for="transaction in transactions" :key="transaction.id" class="hover:bg-gray-50">
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
              {{ transaction.transaction_id }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
              {{ transaction.order_id }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
              {{ formatPrice(transaction.amount) }} ریال
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span
                :class="[
                  'inline-flex px-2 py-1 text-xs font-semibold rounded-full',
                  getStatusClass(transaction.status)
                ]"
              >
                {{ getStatusText(transaction.status) }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
              {{ formatDate(transaction.created_at) }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
              <button
                class="text-blue-600 hover:text-blue-900 mr-3"
                @click="viewTransaction(transaction)"
              >
                <Icon name="ph:eye" class="w-4 h-4" />
              </button>
              <button
                v-if="transaction.status === 'success'"
                class="text-red-600 hover:text-red-900"
                @click="refundTransaction(transaction)"
              >
                <Icon name="ph:arrow-counter-clockwise" class="w-4 h-4" />
              </button>
            </td>
          </tr>
        </tbody>
      </table>

      <!-- Pagination -->
      <div v-if="totalPages > 1" class="px-6 py-4 border-t border-gray-200">
        <div class="flex items-center justify-between">
          <div class="text-sm text-gray-700">
            نمایش {{ (currentPage - 1) * limit + 1 }} تا {{ Math.min(currentPage * limit, total) }} از {{ total }} تراکنش
          </div>
          <div class="flex items-center space-x-2">
            <button
              :disabled="currentPage === 1"
              class="px-3 py-1 text-sm font-medium text-gray-500 bg-white border border-gray-300 rounded-md hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
              @click="changePage(currentPage - 1)"
            >
              قبلی
            </button>
            <span class="px-3 py-1 text-sm font-medium text-gray-700">
              صفحه {{ currentPage }} از {{ totalPages }}
            </span>
            <button
              :disabled="currentPage === totalPages"
              class="px-3 py-1 text-sm font-medium text-gray-500 bg-white border border-gray-300 rounded-md hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
              @click="changePage(currentPage + 1)"
            >
              بعدی
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Empty State -->
    <div v-else class="text-center py-12">
      <Icon name="ph:receipt" class="w-12 h-12 text-gray-400 mx-auto mb-4" />
      <h3 class="text-lg font-medium text-gray-900 mb-2">هیچ تراکنشی یافت نشد</h3>
      <p class="text-gray-500">هنوز هیچ تراکنشی از طریق درگاه ملی انجام نشده است.</p>
    </div>

    <!-- Transaction Detail Modal -->
    <div v-if="showTransactionModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg p-6 max-w-2xl w-full mx-4 max-h-[90vh] overflow-y-auto">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-medium text-gray-900">جزئیات تراکنش</h3>
          <button
            class="text-gray-400 hover:text-gray-600"
            @click="showTransactionModal = false"
          >
            <Icon name="ph:x" class="w-6 h-6" />
          </button>
        </div>
        
        <div v-if="selectedTransaction" class="space-y-4">
          <div class="grid grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700">شناسه تراکنش</label>
              <p class="mt-1 text-sm text-gray-900">{{ selectedTransaction.transaction_id }}</p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">شناسه سفارش</label>
              <p class="mt-1 text-sm text-gray-900">{{ selectedTransaction.order_id }}</p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">مبلغ</label>
              <p class="mt-1 text-sm text-gray-900">{{ formatPrice(selectedTransaction.amount) }} ریال</p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">وضعیت</label>
              <span
                :class="[
                  'inline-flex px-2 py-1 text-xs font-semibold rounded-full mt-1',
                  getStatusClass(selectedTransaction.status)
                ]"
              >
                {{ getStatusText(selectedTransaction.status) }}
              </span>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">تاریخ ایجاد</label>
              <p class="mt-1 text-sm text-gray-900">{{ formatDate(selectedTransaction.created_at) }}</p>
            </div>
            <div v-if="selectedTransaction.updated_at">
              <label class="block text-sm font-medium text-gray-700">تاریخ به‌روزرسانی</label>
              <p class="mt-1 text-sm text-gray-900">{{ formatDate(selectedTransaction.updated_at) }}</p>
            </div>
          </div>
          
          <div v-if="selectedTransaction.description">
            <label class="block text-sm font-medium text-gray-700">توضیحات</label>
            <p class="mt-1 text-sm text-gray-900">{{ selectedTransaction.description }}</p>
          </div>
          
          <div v-if="selectedTransaction.gateway_token">
            <label class="block text-sm font-medium text-gray-700">توکن درگاه</label>
            <p class="mt-1 text-sm text-gray-900 font-mono">{{ selectedTransaction.gateway_token }}</p>
          </div>
        </div>
        
        <div class="flex justify-end mt-6">
          <button
            class="px-4 py-2 text-sm font-medium text-gray-600 bg-gray-50 border border-gray-200 rounded-md hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-500"
            @click="showTransactionModal = false"
          >
            بستن
          </button>
        </div>
      </div>
    </div>

    <!-- Refund Confirmation Modal -->
    <div v-if="showRefundModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg p-6 max-w-md w-full mx-4">
        <div class="flex items-center space-x-3 mb-4">
          <Icon name="ph:warning" class="w-6 h-6 text-yellow-600" />
          <h3 class="text-lg font-medium text-gray-900">تایید بازگشت وجه</h3>
        </div>
        <p class="text-gray-600 mb-6">
          آیا از بازگشت وجه تراکنش 
          <strong>{{ selectedTransaction?.transaction_id }}</strong> 
          به مبلغ 
          <strong>{{ formatPrice(selectedTransaction?.amount || 0) }} ریال</strong> 
          اطمینان دارید؟
        </p>
        <div class="flex justify-end space-x-3">
          <button
            class="px-4 py-2 text-sm font-medium text-gray-600 bg-gray-50 border border-gray-200 rounded-md hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-500"
            @click="showRefundModal = false"
          >
            انصراف
          </button>
          <button
            :disabled="refunding"
            class="px-4 py-2 text-sm font-medium text-white bg-red-600 border border-transparent rounded-md hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500 disabled:opacity-50"
            @click="confirmRefund"
          >
            <Icon v-if="refunding" name="ph:spinner" class="w-4 h-4 animate-spin" />
            <span class="mr-2">{{ refunding ? 'در حال بازگشت...' : 'بازگشت وجه' }}</span>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import { useAuth } from '~/composables/useAuth';
import { useToast } from '~/composables/useToast'

// احراز هویت
const { user, isAuthenticated } = useAuth();

// بررسی دسترسی admin
const hasAccess = computed(() => {
  if (!isAuthenticated.value) {
    return false;
  }

  const userRole = user.value?.role?.toLowerCase() || '';
  const adminRoles = ['admin', 'developer'];
  return adminRoles.includes(userRole);
});

// Types
interface Transaction {
  id: number
  transaction_id: string
  order_id: string
  amount: number
  status: string
  description?: string
  gateway_token?: string
  created_at: string
  updated_at?: string
}

// Reactive data
const loading = ref(false)
const refunding = ref(false)
const showTransactionModal = ref(false)
const showRefundModal = ref(false)
const selectedTransaction = ref<Transaction | null>(null)

const transactions = ref<Transaction[]>([])
const currentPage = ref(1)
const limit = ref(20)
const total = ref(0)
const totalPages = ref(0)

// Methods
const toast = useToast()

const loadTransactions = async () => {
  try {
    loading.value = true
    interface TransactionsData {
      transactions: Transaction[]
      total: number
    }
    interface ApiResponse {
      success?: boolean
      data?: TransactionsData
    }
    const response = await $fetch<ApiResponse>(`/api/admin/payment-gateways/melli/transactions?page=${currentPage.value}&limit=${limit.value}`)
    
    if (response.success && response.data) {
      transactions.value = response.data.transactions || []
      total.value = response.data.total || 0
      totalPages.value = Math.ceil(total.value / limit.value)
    }
  } catch (error) {
    const errorMessage = error instanceof Error ? error.message : 'خطای نامشخص'
    console.error('خطا در بارگذاری تراکنش‌ها:', errorMessage)
    toast.showError('خطا در بارگذاری تراکنش‌های درگاه ملی')
  } finally {
    loading.value = false
  }
}

const changePage = (page: number) => {
  if (page >= 1 && page <= totalPages.value) {
    currentPage.value = page
    loadTransactions()
  }
}

const viewTransaction = (transaction: Transaction) => {
  selectedTransaction.value = transaction
  showTransactionModal.value = true
}

const refundTransaction = (transaction: Transaction) => {
  selectedTransaction.value = transaction
  showRefundModal.value = true
}

const confirmRefund = async () => {
  if (!selectedTransaction.value) return
  
  try {
    refunding.value = true
    interface ApiResponse {
      success?: boolean
      message?: string
    }
    const response = await $fetch<ApiResponse>('/api/payments/melli/refund', {
      method: 'POST',
      body: {
        token: selectedTransaction.value.gateway_token,
        amount: selectedTransaction.value.amount
      }
    })
    
    if (response.success) {
      toast.showSuccess('بازگشت وجه با موفقیت انجام شد')
      showRefundModal.value = false
      loadTransactions() // بارگذاری مجدد تراکنش‌ها
    }
  } catch (error) {
    const errorMessage = error instanceof Error ? error.message : 'خطای نامشخص'
    console.error('خطا در بازگشت وجه:', errorMessage)
    toast.showError('خطا در بازگشت وجه')
  } finally {
    refunding.value = false
  }
}

const formatPrice = (price: number) => {
  return new Intl.NumberFormat('fa-IR').format(price)
}

const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleDateString('fa-IR', {
    year: 'numeric',
    month: 'long',
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
    case 'failed':
      return 'bg-red-100 text-red-800'
    case 'refunded':
      return 'bg-gray-100 text-gray-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getStatusText = (status: string) => {
  switch (status) {
    case 'success':
      return 'موفق'
    case 'pending':
      return 'در انتظار'
    case 'failed':
      return 'ناموفق'
    case 'refunded':
      return 'بازگشت شده'
    default:
      return 'نامشخص'
  }
}

// Lifecycle
onMounted(() => {
  loadTransactions()
})
</script> 
