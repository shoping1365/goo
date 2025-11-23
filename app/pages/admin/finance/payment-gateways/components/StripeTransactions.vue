<template>
  <div v-if="hasAccess" class="bg-white rounded-lg shadow-md p-6">
    <div class="flex items-center justify-between mb-6">
      <h2 class="text-xl font-bold text-gray-800">تراکنش‌های درگاه استرایپ</h2>
      <div class="flex items-center space-x-2 space-x-reverse">
        <button
          :disabled="loading"
          class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 disabled:opacity-50"
          @click="loadTransactions"
        >
          {{ loading ? 'در حال بارگذاری...' : 'بارگذاری مجدد' }}
        </button>
      </div>
    </div>

    <!-- فیلترها -->
    <div class="mb-6 p-6 bg-gray-50 rounded-lg">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
          <select
            v-model="filters.status"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            @change="loadTransactions"
          >
            <option value="">همه</option>
            <option value="requires_payment_method">نیاز به روش پرداخت</option>
            <option value="requires_confirmation">نیاز به تایید</option>
            <option value="requires_action">نیاز به اقدام</option>
            <option value="processing">در حال پردازش</option>
            <option value="requires_capture">نیاز به تسویه</option>
            <option value="canceled">لغو شده</option>
            <option value="succeeded">موفق</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">از تاریخ</label>
          <input
            v-model="filters.start_date"
            type="date"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            @change="loadTransactions"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تا تاریخ</label>
          <input
            v-model="filters.end_date"
            type="date"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            @change="loadTransactions"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">جستجو</label>
          <input
            v-model="filters.search"
            type="text"
            placeholder="Payment Intent ID یا Customer ID"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
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
              Payment Intent ID
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
          <tr v-for="transaction in transactions" :key="transaction.id || transaction.payment_intent_id" class="hover:bg-gray-50">
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
              {{ transaction.payment_intent_id }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
              {{ formatAmount(transaction.amount) }} {{ transaction.currency }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span :class="getStatusClass(transaction.status)" class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full">
                {{ getStatusText(transaction.status) }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
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
                  v-if="transaction.status === 'succeeded'"
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

    <!-- پیام خالی -->
    <div v-if="transactions.length === 0 && !loading" class="text-center py-8">
      <p class="text-gray-500">هیچ تراکنشی یافت نشد</p>
    </div>

    <!-- صفحه‌بندی -->
    <div v-if="pagination.total_pages > 1" class="mt-6 flex items-center justify-between">
      <div class="text-sm text-gray-700">
        نمایش {{ pagination.from }} تا {{ pagination.to }} از {{ pagination.total }} تراکنش
      </div>
      <div class="flex items-center space-x-2 space-x-reverse">
        <button
          :disabled="pagination.current_page <= 1"
          class="px-3 py-2 border border-gray-300 rounded-md text-sm font-medium text-gray-700 hover:bg-gray-50 disabled:opacity-50"
          @click="changePage(pagination.current_page - 1)"
        >
          قبلی
        </button>
        <span class="px-3 py-2 text-sm text-gray-700">
          صفحه {{ pagination.current_page }} از {{ pagination.total_pages }}
        </span>
        <button
          :disabled="pagination.current_page >= pagination.total_pages"
          class="px-3 py-2 border border-gray-300 rounded-md text-sm font-medium text-gray-700 hover:bg-gray-50 disabled:opacity-50"
          @click="changePage(pagination.current_page + 1)"
        >
          بعدی
        </button>
      </div>
    </div>

    <!-- مودال جزئیات تراکنش -->
    <div v-if="selectedTransaction" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
      <div class="relative top-20 mx-auto p-5 border w-full max-w-md sm:max-w-lg md:max-w-xl shadow-lg rounded-md bg-white">
        <div class="mt-3">
          <h3 class="text-lg font-medium text-gray-900 mb-4">جزئیات تراکنش</h3>
          <div class="space-y-3">
            <div>
              <span class="font-medium">Payment Intent ID:</span>
              <span class="mr-2">{{ selectedTransaction.payment_intent_id }}</span>
            </div>
            <div>
              <span class="font-medium">مبلغ:</span>
              <span class="mr-2">{{ formatAmount(selectedTransaction.amount) }} {{ selectedTransaction.currency }}</span>
            </div>
            <div>
              <span class="font-medium">وضعیت:</span>
              <span class="mr-2">{{ getStatusText(selectedTransaction.status) }}</span>
            </div>
            <div>
              <span class="font-medium">تاریخ:</span>
              <span class="mr-2">{{ formatDate(selectedTransaction.created_at) }}</span>
            </div>
            <div v-if="selectedTransaction.customer_id">
              <span class="font-medium">Customer ID:</span>
              <span class="mr-2">{{ selectedTransaction.customer_id }}</span>
            </div>
            <div v-if="selectedTransaction.description">
              <span class="font-medium">توضیحات:</span>
              <span class="mr-2">{{ selectedTransaction.description }}</span>
            </div>
          </div>
          <div class="mt-6 flex justify-end">
            <button
              class="px-4 py-2 bg-gray-500 text-white rounded-md hover:bg-gray-600"
              @click="selectedTransaction = null"
            >
              بستن
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
</script>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue';
import { useAuth } from '~/composables/useAuth';
import { useToast } from '~/composables/useToast';

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

// بررسی احراز هویت و دسترسی admin - نمایش 404 در صورت عدم دسترسی
const checkAuth = async (): Promise<void> => {
  if (!hasAccess.value) {
    await navigateTo('/404', { external: false });
  }
};

// بررسی احراز هویت در هنگام mount
onMounted(async () => {
  await checkAuth();
});

// بررسی احراز هویت هنگام تغییر وضعیت احراز هویت
watch([isAuthenticated, hasAccess], async () => {
  if (!hasAccess.value) {
    await checkAuth();
  }
});

const $toast = useToast()

// Reactive data
const transactions = ref([])
const loading = ref(false)
const selectedTransaction = ref(null)
const filters = ref({
  status: '',
  start_date: '',
  end_date: '',
  search: ''
})
const pagination = ref({
  current_page: 1,
  total_pages: 1,
  total: 0,
  from: 0,
  to: 0
})

// Load transactions on mount
onMounted(async () => {
  await loadTransactions()
})

// Load transactions
const loadTransactions = async () => {
  loading.value = true
  try {
    interface Transaction {
      id?: string | number
      payment_intent_id?: string
      amount?: number
      status?: string
      [key: string]: unknown
    }
    interface TransactionsData {
      transactions: Transaction[]
      pagination: {
        current_page: number
        total_pages: number
        total: number
        from: number
        to: number
      }
    }
    interface ApiResponse {
      success?: boolean
      data?: TransactionsData
    }
    const query = {
      page: pagination.value.current_page,
      ...filters.value
    }
    
    const response = await $fetch<ApiResponse>('/api/admin/payment-gateways/stripe.transactions.get', {
      query
    })
    
    if (response.success && response.data) {
      transactions.value = response.data.transactions
      pagination.value = response.data.pagination
    }
  } catch (_error) {
    if ($toast && typeof $toast.showError === 'function') {
      $toast.showError('خطا در بارگذاری تراکنش‌ها')
    }
  } finally {
    loading.value = false
  }
}

// Change page
const changePage = (page) => {
  pagination.value.current_page = page
  loadTransactions()
}

// Debounced search
let searchTimeout
const debounceSearch = () => {
  clearTimeout(searchTimeout)
  searchTimeout = setTimeout(() => {
    pagination.value.current_page = 1
    loadTransactions()
  }, 500)
}

// View transaction details
const viewTransaction = (transaction) => {
  selectedTransaction.value = transaction
}

// Refund transaction
const refundTransaction = async (transaction: { payment_intent_id?: string; amount?: number }) => {
  if (!confirm('آیا از بازگشت وجه این تراکنش اطمینان دارید؟')) {
    return
  }
  
  try {
    interface ApiResponse {
      success?: boolean
      message?: string
    }
    const response = await $fetch<ApiResponse>('/api/admin/payment-gateways/stripe.refund.post', {
      method: 'POST',
      body: {
        payment_intent_id: transaction.payment_intent_id,
        amount: transaction.amount
      }
    })
    
    if (response.success) {
      if ($toast && typeof $toast.showSuccess === 'function') {
        $toast.showSuccess('بازگشت وجه با موفقیت انجام شد')
      }
      loadTransactions()
    } else {
      if ($toast && typeof $toast.showError === 'function') {
        $toast.showError(response.message || 'خطا در بازگشت وجه')
      }
    }
  } catch (_error) {
    if ($toast && typeof $toast.showError === 'function') {
      $toast.showError('خطا در بازگشت وجه')
    }
  }
}

// Format amount
const formatAmount = (amount) => {
  return new Intl.NumberFormat('fa-IR').format(amount / 100)
}

// Format date
const formatDate = (date) => {
  return new Date(date).toLocaleDateString('fa-IR')
}

// Get status class
const getStatusClass = (status) => {
  switch (status) {
    case 'succeeded':
      return 'bg-green-100 text-green-800'
    case 'processing':
      return 'bg-blue-100 text-blue-800'
    case 'requires_payment_method':
    case 'requires_confirmation':
    case 'requires_action':
      return 'bg-yellow-100 text-yellow-800'
    case 'canceled':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

// Get status text
const getStatusText = (status) => {
  switch (status) {
    case 'succeeded':
      return 'موفق'
    case 'processing':
      return 'در حال پردازش'
    case 'requires_payment_method':
      return 'نیاز به روش پرداخت'
    case 'requires_confirmation':
      return 'نیاز به تایید'
    case 'requires_action':
      return 'نیاز به اقدام'
    case 'requires_capture':
      return 'نیاز به تسویه'
    case 'canceled':
      return 'لغو شده'
    default:
      return status
  }
}
</script> 
