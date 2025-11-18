<template>
  <div class="bg-white rounded-lg shadow-md p-6">
    <div class="flex items-center justify-between mb-6">
      <h2 class="text-xl font-bold text-gray-800">تراکنش‌های درگاه پی‌پال</h2>
      <div class="flex items-center space-x-2 space-x-reverse">
        <button
          @click="loadTransactions"
          :disabled="loading"
          class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 disabled:opacity-50"
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
            @change="loadTransactions"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            <option value="">همه</option>
            <option value="pending">در انتظار</option>
            <option value="completed">تکمیل شده</option>
            <option value="failed">ناموفق</option>
            <option value="refunded">بازگشت شده</option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">از تاریخ</label>
          <input
            v-model="filters.start_date"
            type="date"
            @change="loadTransactions"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تا تاریخ</label>
          <input
            v-model="filters.end_date"
            type="date"
            @change="loadTransactions"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">جستجو</label>
          <input
            v-model="filters.search"
            type="text"
            @input="debounceSearch"
            placeholder="شماره تراکنش یا ایمیل"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
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
              شماره تراکنش
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
                  @click="viewTransaction(transaction)"
                  class="text-blue-600 hover:text-blue-900"
                >
                  مشاهده
                </button>
                <button
                  v-if="transaction.status === 'completed'"
                  @click="refundTransaction(transaction)"
                  class="text-red-600 hover:text-red-900"
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
          @click="changePage(pagination.current_page - 1)"
          :disabled="pagination.current_page <= 1"
          class="px-3 py-2 border border-gray-300 rounded-md text-sm font-medium text-gray-700 hover:bg-gray-50 disabled:opacity-50"
        >
          قبلی
        </button>
        <span class="px-3 py-2 text-sm text-gray-700">
          صفحه {{ pagination.current_page }} از {{ pagination.total_pages }}
        </span>
        <button
          @click="changePage(pagination.current_page + 1)"
          :disabled="pagination.current_page >= pagination.total_pages"
          class="px-3 py-2 border border-gray-300 rounded-md text-sm font-medium text-gray-700 hover:bg-gray-50 disabled:opacity-50"
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
              <span class="font-medium">شماره تراکنش:</span>
              <span class="mr-2">{{ selectedTransaction.transaction_id }}</span>
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
            <div v-if="selectedTransaction.description">
              <span class="font-medium">توضیحات:</span>
              <span class="mr-2">{{ selectedTransaction.description }}</span>
            </div>
          </div>
          <div class="mt-6 flex justify-end">
            <button
              @click="selectedTransaction = null"
              class="px-4 py-2 bg-gray-500 text-white rounded-md hover:bg-gray-600"
            >
              بستن
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
const { $toast } = useNuxtApp()

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
    const query = {
      page: pagination.value.current_page,
      ...filters.value
    }
    
    const response = await $fetch('/api/admin/payment-gateways/paypal.transactions.get', {
      query
    })
    
    if (response.success) {
      transactions.value = response.data.transactions
      pagination.value = response.data.pagination
    }
  } catch (error) {
    $toast.error('خطا در بارگذاری تراکنش‌ها')
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
const refundTransaction = async (transaction) => {
  if (!confirm('آیا از بازگشت وجه این تراکنش اطمینان دارید؟')) {
    return
  }
  
  try {
    const response = await $fetch('/api/admin/payment-gateways/paypal.refund.post', {
      method: 'POST',
      body: {
        transaction_id: transaction.transaction_id,
        amount: transaction.amount
      }
    })
    
    if (response.success) {
      $toast.success('بازگشت وجه با موفقیت انجام شد')
      loadTransactions()
    } else {
      $toast.error(response.message || 'خطا در بازگشت وجه')
    }
  } catch (error) {
    $toast.error('خطا در بازگشت وجه')
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
    case 'completed':
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

// Get status text
const getStatusText = (status) => {
  switch (status) {
    case 'completed':
      return 'تکمیل شده'
    case 'pending':
      return 'در انتظار'
    case 'failed':
      return 'ناموفق'
    case 'refunded':
      return 'بازگشت شده'
    default:
      return status
  }
}
</script> 
