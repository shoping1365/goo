<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex justify-between items-center">
        <div>
          <h2 class="text-xl font-bold text-gray-900">مدیریت تراکنش‌های گیفت کارت</h2>
          <p class="text-gray-600 mt-1">ثبت، ردیابی و مدیریت تمام تراکنش‌های مربوط به گیفت کارت‌ها</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button 
            @click="showCreateModal = true"
            class="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
          >
            ثبت تراکنش جدید
          </button>
          <button 
            @click="exportTransactions"
            class="px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2"
          >
            خروجی Excel
          </button>
        </div>
      </div>
    </div>

    <!-- آمار کلی -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-blue-100 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">کل تراکنش‌ها</p>
            <p class="text-2xl font-bold text-gray-900">{{ totalTransactions }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-green-100 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">تراکنش‌های موفق</p>
            <p class="text-2xl font-bold text-gray-900">{{ successfulTransactions }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-red-100 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">تراکنش‌های ناموفق</p>
            <p class="text-2xl font-bold text-gray-900">{{ failedTransactions }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-yellow-100 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">در انتظار تایید</p>
            <p class="text-2xl font-bold text-gray-900">{{ pendingTransactions }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- فیلترها و جستجو -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">نوع تراکنش</label>
          <select 
            v-model="filters.type"
            class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه</option>
            <option value="purchase">خرید</option>
            <option value="usage">استفاده</option>
            <option value="refund">بازپرداخت</option>
            <option value="expiry">انقضا</option>
          </select>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">وضعیت</label>
          <select 
            v-model="filters.status"
            class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه</option>
            <option value="successful">موفق</option>
            <option value="failed">ناموفق</option>
            <option value="pending">در انتظار</option>
            <option value="cancelled">لغو شده</option>
          </select>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">تاریخ از</label>
          <input
            v-model="filters.dateFrom"
            type="date"
            class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
          />
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">تاریخ تا</label>
          <input
            v-model="filters.dateTo"
            type="date"
            class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
          />
        </div>
      </div>
      
      <div class="mt-4 flex items-center space-x-3 space-x-reverse">
        <input
          v-model="filters.search"
          type="text"
          placeholder="جستجو بر اساس کد کارت، شماره سفارش یا توضیحات..."
          class="flex-1 px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
        />
        <button 
          @click="applyFilters"
          class="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
        >
          اعمال فیلتر
        </button>
        <button 
          @click="clearFilters"
          class="px-4 py-2 bg-gray-600 text-white text-sm font-medium rounded-lg hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
        >
          پاک کردن
        </button>
      </div>
    </div>

    <!-- جدول تراکنش‌ها -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-medium text-gray-900">لیست تراکنش‌ها</h3>
      </div>
      
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                کد تراکنش
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                کد کارت
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                نوع
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
            <tr v-for="transaction in filteredTransactions" :key="transaction.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                {{ transaction.transactionId }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ transaction.cardCode }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getTypeBadgeClass(transaction.type)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                  {{ getTypeLabel(transaction.type) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ formatCurrency(transaction.amount) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusBadgeClass(transaction.status)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                  {{ getStatusLabel(transaction.status) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ formatDate(transaction.date) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button 
                  @click="viewTransactionDetails(transaction)"
                  class="text-blue-600 hover:text-blue-900 ml-3"
                >
                  جزئیات
                </button>
                <button 
                  v-if="transaction.status === 'failed'"
                  @click="retryTransaction(transaction)"
                  class="text-green-600 hover:text-green-900 ml-3"
                >
                  تلاش مجدد
                </button>
                <button 
                  v-if="transaction.status === 'successful' && transaction.type === 'usage'"
                  @click="openRefundModal(transaction)"
                  class="text-red-600 hover:text-red-900 ml-3"
                >
                  بازپرداخت
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      
      <!-- صفحه‌بندی -->
      <div class="bg-white px-4 py-3 flex items-center justify-between border-t border-gray-200 sm:px-6">
        <div class="flex-1 flex justify-between sm:hidden">
          <button 
            @click="previousPage"
            :disabled="currentPage === 1"
            class="relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
          >
            قبلی
          </button>
          <button 
            @click="nextPage"
            :disabled="currentPage === totalPages"
            class="mr-3 relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
          >
            بعدی
          </button>
        </div>
        <div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
          <div>
            <p class="text-sm text-gray-700">
              نمایش 
              <span class="font-medium">{{ (currentPage - 1) * pageSize + 1 }}</span>
              تا 
              <span class="font-medium">{{ Math.min(currentPage * pageSize, totalTransactions) }}</span>
              از 
              <span class="font-medium">{{ totalTransactions }}</span>
              نتیجه
            </p>
          </div>
          <div>
            <nav class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px" aria-label="Pagination">
              <button 
                @click="previousPage"
                :disabled="currentPage === 1"
                class="relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
              >
                <span class="sr-only">قبلی</span>
                <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                  <path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd" />
                </svg>
              </button>
              <button 
                v-for="page in visiblePages" 
                :key="page"
                @click="goToPage(page)"
                :class="{
                  'bg-blue-50 border-blue-500 text-blue-600': page === currentPage,
                  'bg-white border-gray-300 text-gray-500 hover:bg-gray-50': page !== currentPage
                }"
                class="relative inline-flex items-center px-4 py-2 border text-sm font-medium"
              >
                {{ page }}
              </button>
              <button 
                @click="nextPage"
                :disabled="currentPage === totalPages"
                class="relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
              >
                <span class="sr-only">بعدی</span>
                <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                  <path fill-rule="evenodd" d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z" clip-rule="evenodd" />
                </svg>
              </button>
            </nav>
          </div>
        </div>
      </div>
    </div>

    <!-- مودال‌ها -->
    <GiftCardTransactionCreateModal 
      v-if="showCreateModal"
      @close="showCreateModal = false"
      @created="handleTransactionCreated"
    />

    <GiftCardTransactionDetailsModal 
      v-if="showDetailsModal"
      :transaction="selectedTransaction"
      @close="showDetailsModal = false"
    />

    <GiftCardRefundRequestModal 
      v-if="showRefundModal"
      :transaction="selectedTransaction"
      @close="showRefundModal = false"
      @refund-requested="handleRefundRequested"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted } from 'vue'

// کامپوننت‌های مورد نیاز
import GiftCardTransactionCreateModal from './GiftCardTransactionCreateModal.vue'
import GiftCardTransactionDetailsModal from './GiftCardTransactionDetailsModal.vue'
import GiftCardRefundRequestModal from './GiftCardRefundRequestModal.vue'

// Reactive data
const showCreateModal = ref(false)
const showDetailsModal = ref(false)
const showRefundModal = ref(false)
const selectedTransaction = ref(null)
const currentPage = ref(1)
const pageSize = ref(10)

// فیلترها
const filters = reactive({
  type: '',
  status: '',
  dateFrom: '',
  dateTo: '',
  search: ''
})

// داده‌های نمونه
const transactions = ref([
  {
    id: 1,
    transactionId: 'TXN-2024-001',
    cardCode: 'GC-BIRTHDAY-2024-001',
    type: 'purchase',
    amount: 500000,
    status: 'successful',
    date: new Date(Date.now() - 86400000),
    description: 'خرید گیفت کارت تولد',
    orderId: 'ORD-001',
    customerName: 'علی احمدی',
    customerEmail: 'ali@example.com',
    paymentMethod: 'online',
    gateway: 'zarinpal',
    referenceId: 'REF-123456'
  },
  {
    id: 2,
    transactionId: 'TXN-2024-002',
    cardCode: 'GC-WEDDING-2024-002',
    type: 'usage',
    amount: 250000,
    status: 'successful',
    date: new Date(Date.now() - 172800000),
    description: 'استفاده از گیفت کارت عروسی',
    orderId: 'ORD-002',
    customerName: 'فاطمه محمدی',
    customerEmail: 'fateme@example.com',
    remainingAmount: 750000,
    usageCount: 1
  },
  {
    id: 3,
    transactionId: 'TXN-2024-003',
    cardCode: 'GC-NEWYEAR-2024-003',
    type: 'purchase',
    amount: 200000,
    status: 'failed',
    date: new Date(Date.now() - 259200000),
    description: 'خرید گیفت کارت سال نو',
    orderId: 'ORD-003',
    customerName: 'حسن رضایی',
    customerEmail: 'hasan@example.com',
    paymentMethod: 'online',
    gateway: 'zarinpal',
    errorMessage: 'خطا در پرداخت'
  },
  {
    id: 4,
    transactionId: 'TXN-2024-004',
    cardCode: 'GC-BIRTHDAY-2024-004',
    type: 'refund',
    amount: 100000,
    status: 'pending',
    date: new Date(Date.now() - 345600000),
    description: 'درخواست بازپرداخت',
    orderId: 'ORD-004',
    customerName: 'مریم کریمی',
    customerEmail: 'maryam@example.com',
    refundReason: 'خرید اشتباه',
    originalTransactionId: 'TXN-2024-001'
  }
])

// Computed properties
const totalTransactions = computed(() => transactions.value.length)
const successfulTransactions = computed(() => transactions.value.filter(t => t.status === 'successful').length)
const failedTransactions = computed(() => transactions.value.filter(t => t.status === 'failed').length)
const pendingTransactions = computed(() => transactions.value.filter(t => t.status === 'pending').length)

const filteredTransactions = computed(() => {
  let filtered = transactions.value

  if (filters.type) {
    filtered = filtered.filter(t => t.type === filters.type)
  }

  if (filters.status) {
    filtered = filtered.filter(t => t.status === filters.status)
  }

  if (filters.dateFrom) {
    filtered = filtered.filter(t => new Date(t.date) >= new Date(filters.dateFrom))
  }

  if (filters.dateTo) {
    filtered = filtered.filter(t => new Date(t.date) <= new Date(filters.dateTo))
  }

  if (filters.search) {
    const search = filters.search.toLowerCase()
    filtered = filtered.filter(t => 
      t.cardCode.toLowerCase().includes(search) ||
      t.transactionId.toLowerCase().includes(search) ||
      t.orderId.toLowerCase().includes(search) ||
      t.description.toLowerCase().includes(search)
    )
  }

  return filtered
})

const totalPages = computed(() => Math.ceil(filteredTransactions.value.length / pageSize.value))

const visiblePages = computed(() => {
  const pages = []
  const start = Math.max(1, currentPage.value - 2)
  const end = Math.min(totalPages.value, currentPage.value + 2)
  
  for (let i = start; i <= end; i++) {
    pages.push(i)
  }
  
  return pages
})

// Methods
const applyFilters = () => {
  currentPage.value = 1
  // اعمال فیلترها
  console.log('فیلترها اعمال شد:', filters)
}

const clearFilters = () => {
  Object.assign(filters, {
    type: '',
    status: '',
    dateFrom: '',
    dateTo: '',
    search: ''
  })
  currentPage.value = 1
}

const getTypeBadgeClass = (type: string) => {
  const classes = {
    purchase: 'bg-blue-100 text-blue-800',
    usage: 'bg-green-100 text-green-800',
    refund: 'bg-yellow-100 text-yellow-800',
    expiry: 'bg-red-100 text-red-800'
  }
  return classes[type] || 'bg-gray-100 text-gray-800'
}

const getTypeLabel = (type: string) => {
  const labels = {
    purchase: 'خرید',
    usage: 'استفاده',
    refund: 'بازپرداخت',
    expiry: 'انقضا'
  }
  return labels[type] || type
}

const getStatusBadgeClass = (status: string) => {
  const classes = {
    successful: 'bg-green-100 text-green-800',
    failed: 'bg-red-100 text-red-800',
    pending: 'bg-yellow-100 text-yellow-800',
    cancelled: 'bg-gray-100 text-gray-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getStatusLabel = (status: string) => {
  const labels = {
    successful: 'موفق',
    failed: 'ناموفق',
    pending: 'در انتظار',
    cancelled: 'لغو شده'
  }
  return labels[status] || status
}

const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

const formatDate = (date: Date) => {
  return new Intl.DateTimeFormat('fa-IR').format(date)
}

const viewTransactionDetails = (transaction: any) => {
  selectedTransaction.value = transaction
  showDetailsModal.value = true
}

const retryTransaction = async (transaction: any) => {
  if (confirm('آیا مطمئن هستید که می‌خواهید این تراکنش را دوباره تلاش کنید؟')) {
    try {
      // شبیه‌سازی retry
      await new Promise(resolve => setTimeout(resolve, 1000))
      
      // به‌روزرسانی وضعیت
      transaction.status = 'successful'
      transaction.date = new Date()
      
      alert('تراکنش با موفقیت انجام شد')
    } catch (error) {
      alert('خطا در انجام تراکنش')
    }
  }
}

const openRefundModal = (transaction: any) => {
  selectedTransaction.value = transaction
  showRefundModal.value = true
}

const exportTransactions = () => {
  // شبیه‌سازی export
  console.log('خروجی Excel در حال آماده‌سازی...')
  alert('فایل Excel آماده شد')
}

const handleTransactionCreated = (transaction: any) => {
  transactions.value.unshift(transaction)
  showCreateModal.value = false
}

const handleRefundRequested = (refundData: any) => {
  // ایجاد تراکنش بازپرداخت
  const refundTransaction = {
    id: transactions.value.length + 1,
    transactionId: `TXN-${Date.now()}`,
    cardCode: refundData.cardCode,
    type: 'refund',
    amount: refundData.amount,
    status: 'pending',
    date: new Date(),
    description: `بازپرداخت: ${refundData.reason}`,
    orderId: refundData.orderId,
    customerName: refundData.customerName,
    customerEmail: refundData.customerEmail,
    refundReason: refundData.reason,
    originalTransactionId: refundData.originalTransactionId
  }
  
  transactions.value.unshift(refundTransaction)
  showRefundModal.value = false
  alert('درخواست بازپرداخت ثبت شد')
}

const previousPage = () => {
  if (currentPage.value > 1) {
    currentPage.value--
  }
}

const nextPage = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value++
  }
}

const goToPage = (page: number) => {
  currentPage.value = page
}

// Lifecycle
onMounted(() => {
  console.log('Gift card transactions component mounted')
})
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 
