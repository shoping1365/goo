<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">مدیریت بازپرداخت‌های گیفت کارت</h3>
        <p class="text-sm text-gray-600">درخواست، تایید و مدیریت بازپرداخت‌های گیفت کارت</p>
      </div>
      <div class="flex gap-2">
        <button
          @click="showNewRefundModal = true"
          class="inline-flex items-center px-4 py-2 bg-red-600 text-white text-sm font-medium rounded-lg hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2"
        >
          <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h10a8 8 0 018 8v2M3 10l6 6m-6-6l6-6"></path>
          </svg>
          درخواست بازپرداخت جدید
        </button>
        <button
          @click="exportRefunds"
          class="inline-flex items-center px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
        >
          <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
          </svg>
          خروجی گزارش
        </button>
      </div>
    </div>

    <!-- آمار بازپرداخت‌ها -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="bg-white p-6 rounded-lg border border-gray-200">
        <div class="flex items-center">
          <div class="p-2 bg-yellow-100 rounded-lg">
            <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
          <div class="ml-3">
            <p class="text-sm font-medium text-gray-600">در انتظار بررسی</p>
            <p class="text-2xl font-bold text-gray-900">{{ pendingRefunds.length }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white p-6 rounded-lg border border-gray-200">
        <div class="flex items-center">
          <div class="p-2 bg-green-100 rounded-lg">
            <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
          <div class="ml-3">
            <p class="text-sm font-medium text-gray-600">تایید شده</p>
            <p class="text-2xl font-bold text-gray-900">{{ approvedRefunds.length }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white p-6 rounded-lg border border-gray-200">
        <div class="flex items-center">
          <div class="p-2 bg-red-100 rounded-lg">
            <svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </div>
          <div class="ml-3">
            <p class="text-sm font-medium text-gray-600">رد شده</p>
            <p class="text-2xl font-bold text-gray-900">{{ rejectedRefunds.length }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white p-6 rounded-lg border border-gray-200">
        <div class="flex items-center">
          <div class="p-2 bg-blue-100 rounded-lg">
            <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
            </svg>
          </div>
          <div class="ml-3">
            <p class="text-sm font-medium text-gray-600">کل مبلغ</p>
            <p class="text-2xl font-bold text-gray-900">{{ formatCurrency(totalRefundAmount) }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- فیلترها و جستجو -->
    <div class="bg-white p-6 rounded-lg border border-gray-200">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <div>
          <input
            v-model="searchQuery"
            type="text"
            placeholder="جستجو در بازپرداخت‌ها..."
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          />
        </div>
        <div>
          <select
            v-model="statusFilter"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          >
            <option value="">همه وضعیت‌ها</option>
            <option value="pending">در انتظار</option>
            <option value="approved">تایید شده</option>
            <option value="rejected">رد شده</option>
            <option value="completed">انجام شده</option>
          </select>
        </div>
        <div>
          <select
            v-model="reasonFilter"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          >
            <option value="">همه دلایل</option>
            <option value="duplicate_purchase">خرید تکراری</option>
            <option value="wrong_amount">مبلغ اشتباه</option>
            <option value="technical_issue">مشکل فنی</option>
            <option value="customer_request">درخواست مشتری</option>
            <option value="expired_card">کارت منقضی شده</option>
          </select>
        </div>
        <div>
          <input
            v-model="dateFilter"
            type="date"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          />
        </div>
      </div>
    </div>

    <!-- جدول بازپرداخت‌ها -->
    <div class="bg-white rounded-lg border border-gray-200 overflow-hidden">
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                شناسه
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                کد گیفت کارت
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                مبلغ
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                درخواست کننده
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                دلیل
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
            <tr
              v-for="refund in filteredRefunds"
              :key="refund.id"
              class="hover:bg-gray-50"
            >
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                #{{ refund.id }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ refund.giftCardCode }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ formatCurrency(refund.amount) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="flex-shrink-0 h-8 w-8">
                    <div class="h-8 w-8 rounded-full bg-gray-300 flex items-center justify-center">
                      <span class="text-sm font-medium text-gray-700">
                        {{ refund.requester.charAt(0) }}
                      </span>
                    </div>
                  </div>
                  <div class="mr-3">
                    <div class="text-sm font-medium text-gray-900">{{ refund.requester }}</div>
                    <div class="text-sm text-gray-500">{{ refund.requesterEmail }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="text-sm text-gray-900">{{ getReasonText(refund.reason) }}</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span
                  :class="{
                    'bg-yellow-100 text-yellow-800': refund.status === 'pending',
                    'bg-green-100 text-green-800': refund.status === 'approved',
                    'bg-red-100 text-red-800': refund.status === 'rejected',
                    'bg-blue-100 text-blue-800': refund.status === 'completed'
                  }"
                  class="px-2 py-1 text-xs font-medium rounded-full"
                >
                  {{ getStatusText(refund.status) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ formatDate(refund.createdAt) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <div class="flex gap-2">
                  <button
                    @click="viewRefundDetails(refund)"
                    class="text-blue-600 hover:text-blue-900"
                  >
                    جزئیات
                  </button>
                  <button
                    v-if="refund.status === 'pending'"
                    @click="approveRefund(refund)"
                    class="text-green-600 hover:text-green-900"
                  >
                    تایید
                  </button>
                  <button
                    v-if="refund.status === 'pending'"
                    @click="rejectRefund(refund)"
                    class="text-red-600 hover:text-red-900"
                  >
                    رد
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- صفحه‌بندی -->
      <div class="bg-white px-4 py-3 border-t border-gray-200 sm:px-6">
        <div class="flex items-center justify-between">
          <div class="flex-1 flex justify-between sm:hidden">
            <button
              @click="previousPage"
              :disabled="currentPage === 1"
              class="relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50"
            >
              قبلی
            </button>
            <button
              @click="nextPage"
              :disabled="currentPage >= totalPages"
              class="mr-3 relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50"
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
                <span class="font-medium">{{ Math.min(currentPage * pageSize, filteredRefunds.length) }}</span>
                از
                <span class="font-medium">{{ filteredRefunds.length }}</span>
                نتیجه
              </p>
            </div>
            <div>
              <nav class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px">
                <button
                  @click="previousPage"
                  :disabled="currentPage === 1"
                  class="relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50"
                >
                  قبلی
                </button>
                <button
                  @click="nextPage"
                  :disabled="currentPage >= totalPages"
                  class="relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50"
                >
                  بعدی
                </button>
              </nav>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- مودال درخواست بازپرداخت جدید -->
    <GiftCardRefundRequestModal
      v-if="showNewRefundModal"
      @close="showNewRefundModal = false"
      @created="handleRefundCreated"
    />

    <!-- مودال جزئیات بازپرداخت -->
    <GiftCardRefundDetailsModal
      v-if="showDetailsModal"
      :refund="selectedRefund"
      @close="showDetailsModal = false"
    />
  </div>
</template>

<script setup>
// تعریف متغیرهای reactive
const searchQuery = ref('')
const statusFilter = ref('')
const reasonFilter = ref('')
const dateFilter = ref('')
const showNewRefundModal = ref(false)
const showDetailsModal = ref(false)
const selectedRefund = ref(null)
const currentPage = ref(1)
const pageSize = ref(10)

// داده‌های نمونه بازپرداخت‌ها
const refunds = ref([
  {
    id: 1,
    giftCardCode: 'GC-2024-001',
    amount: 500000,
    requester: 'احمد محمدی',
    requesterEmail: 'ahmad@example.com',
    reason: 'duplicate_purchase',
    status: 'pending',
    createdAt: new Date(Date.now() - 3600000),
    description: 'کارت به اشتباه دو بار خریداری شد',
    transactionId: 'TXN-001',
    approvedBy: null,
    approvedAt: null,
    rejectionReason: null
  },
  {
    id: 2,
    giftCardCode: 'GC-2024-002',
    amount: 250000,
    requester: 'فاطمه احمدی',
    requesterEmail: 'fateme@example.com',
    reason: 'wrong_amount',
    status: 'approved',
    createdAt: new Date(Date.now() - 7200000),
    description: 'مبلغ اشتباه وارد شد',
    transactionId: 'TXN-002',
    approvedBy: 'مدیر سیستم',
    approvedAt: new Date(Date.now() - 3600000),
    rejectionReason: null
  },
  {
    id: 3,
    giftCardCode: 'GC-2024-003',
    amount: 1000000,
    requester: 'علی رضایی',
    requesterEmail: 'ali@example.com',
    reason: 'technical_issue',
    status: 'rejected',
    createdAt: new Date(Date.now() - 10800000),
    description: 'مشکل فنی در پرداخت',
    transactionId: 'TXN-003',
    approvedBy: null,
    approvedAt: null,
    rejectionReason: 'مشکل فنی تایید نشد'
  },
  {
    id: 4,
    giftCardCode: 'GC-2024-004',
    amount: 750000,
    requester: 'مریم کریمی',
    requesterEmail: 'maryam@example.com',
    reason: 'customer_request',
    status: 'completed',
    createdAt: new Date(Date.now() - 14400000),
    description: 'درخواست بازپرداخت از سوی مشتری',
    transactionId: 'TXN-004',
    approvedBy: 'مدیر سیستم',
    approvedAt: new Date(Date.now() - 7200000),
    rejectionReason: null
  },
  {
    id: 5,
    giftCardCode: 'GC-2024-005',
    amount: 300000,
    requester: 'سارا احمدی',
    requesterEmail: 'sara@example.com',
    reason: 'expired_card',
    status: 'pending',
    createdAt: new Date(Date.now() - 18000000),
    description: 'کارت قبل از استفاده منقضی شد',
    transactionId: 'TXN-005',
    approvedBy: null,
    approvedAt: null,
    rejectionReason: null
  }
])

// محاسبه بازپرداخت‌های در انتظار
const pendingRefunds = computed(() => 
  refunds.value.filter(r => r.status === 'pending')
)

// محاسبه بازپرداخت‌های تایید شده
const approvedRefunds = computed(() => 
  refunds.value.filter(r => r.status === 'approved')
)

// محاسبه بازپرداخت‌های رد شده
const rejectedRefunds = computed(() => 
  refunds.value.filter(r => r.status === 'rejected')
)

// محاسبه کل مبلغ بازپرداخت‌ها
const totalRefundAmount = computed(() => 
  refunds.value.reduce((sum, r) => sum + r.amount, 0)
)

// فیلتر کردن بازپرداخت‌ها
const filteredRefunds = computed(() => {
  return refunds.value.filter(refund => {
    const matchesSearch = 
      refund.requester.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      refund.giftCardCode.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      refund.description.toLowerCase().includes(searchQuery.value.toLowerCase())
    
    const matchesStatus = !statusFilter.value || refund.status === statusFilter.value
    const matchesReason = !reasonFilter.value || refund.reason === reasonFilter.value
    
    const matchesDate = !dateFilter.value || 
      refund.createdAt.toISOString().split('T')[0] === dateFilter.value
    
    return matchesSearch && matchesStatus && matchesReason && matchesDate
  })
})

// محاسبه تعداد صفحات
const totalPages = computed(() => 
  Math.ceil(filteredRefunds.value.length / pageSize.value)
)

// تابع دریافت متن دلیل
const getReasonText = (reason) => {
  const reasonMap = {
    duplicate_purchase: 'خرید تکراری',
    wrong_amount: 'مبلغ اشتباه',
    technical_issue: 'مشکل فنی',
    customer_request: 'درخواست مشتری',
    expired_card: 'کارت منقضی شده'
  }
  return reasonMap[reason] || reason
}

// تابع دریافت متن وضعیت
const getStatusText = (status) => {
  const statusMap = {
    pending: 'در انتظار',
    approved: 'تایید شده',
    rejected: 'رد شده',
    completed: 'انجام شده'
  }
  return statusMap[status] || status
}

// تابع فرمت کردن مبلغ
const formatCurrency = (amount) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

// تابع فرمت کردن تاریخ
const formatDate = (date) => {
  return new Intl.DateTimeFormat('fa-IR', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  }).format(date)
}

// تابع مشاهده جزئیات بازپرداخت
const viewRefundDetails = (refund) => {
  selectedRefund.value = refund
  showDetailsModal.value = true
}

// تابع تایید بازپرداخت
const approveRefund = (refund) => {
  if (confirm(`آیا از تایید بازپرداخت #${refund.id} به مبلغ ${formatCurrency(refund.amount)} اطمینان دارید؟`)) {
    refund.status = 'approved'
    refund.approvedBy = 'مدیر سیستم'
    refund.approvedAt = new Date()
    alert('بازپرداخت با موفقیت تایید شد')
  }
}

// تابع رد بازپرداخت
const rejectRefund = (refund) => {
  const rejectionReason = prompt('دلیل رد بازپرداخت را وارد کنید:')
  if (rejectionReason) {
    refund.status = 'rejected'
    refund.rejectionReason = rejectionReason
    alert('بازپرداخت رد شد')
  }
}

// تابع خروجی گزارش
const exportRefunds = () => {
  // اینجا کد export را اضافه کنید
  console.log('Exporting refunds')
  alert('گزارش بازپرداخت‌ها در حال دانلود است')
}

// تابع مدیریت ایجاد بازپرداخت جدید
const handleRefundCreated = (newRefund) => {
  refunds.value.unshift({
    ...newRefund,
    id: Date.now(),
    createdAt: new Date(),
    status: 'pending'
  })
  showNewRefundModal.value = false
}

// توابع صفحه‌بندی
const nextPage = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value++
  }
}

const previousPage = () => {
  if (currentPage.value > 1) {
    currentPage.value--
  }
}
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 
