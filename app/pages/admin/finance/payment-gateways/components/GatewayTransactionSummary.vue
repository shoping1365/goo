<template>
  <div class="space-y-4">
    <!-- هدر بخش تراکنش‌ها -->
    <div class="flex items-center justify-between">
      <h4 class="text-sm font-semibold text-gray-700">تراکنش‌های اخیر</h4>
      <button 
        @click="showAllTransactions = !showAllTransactions"
        class="text-blue-600 hover:text-blue-700 text-xs font-medium"
      >
        {{ showAllTransactions ? 'نمایش خلاصه' : 'مشاهده همه' }}
      </button>
    </div>

    <!-- خلاصه تراکنش‌ها -->
    <div v-if="!showAllTransactions" class="space-y-3">
      <div v-if="loading" class="text-center py-4">
        <Icon name="heroicons:arrow-path" class="w-5 h-5 animate-spin text-blue-600 mx-auto" />
        <p class="text-xs text-gray-500 mt-1">در حال بارگذاری...</p>
      </div>
      
      <div v-else-if="recentTransactions.length === 0" class="text-center py-4">
        <Icon name="heroicons:credit-card" class="w-8 h-8 text-gray-300 mx-auto" />
        <p class="text-xs text-gray-500 mt-1">هیچ تراکنشی یافت نشد</p>
      </div>
      
      <div v-else class="space-y-2">
        <div 
          v-for="transaction in recentTransactions.slice(0, 3)" 
          :key="transaction.id"
          class="flex items-center justify-between p-2 bg-gray-50 rounded-lg hover:bg-gray-100 transition-colors cursor-pointer"
          @click="viewTransaction(transaction)"
        >
          <div class="flex items-center space-x-3 space-x-reverse">
            <div 
              :class="[
                'w-2 h-2 rounded-full',
                getStatusColor(transaction.status)
              ]"
            ></div>
            <div>
              <p class="text-xs font-medium text-gray-900">{{ transaction.transaction_id }}</p>
              <p class="text-xs text-gray-500">{{ formatDate(transaction.created_at) }}</p>
            </div>
          </div>
          <div class="text-left">
            <p class="text-xs font-semibold text-gray-900">{{ formatCurrency(transaction.amount) }}</p>
            <p class="text-xs text-gray-500">{{ getStatusText(transaction.status) }}</p>
          </div>
        </div>
        
        <div v-if="recentTransactions.length > 3" class="text-center">
          <button 
            @click="showAllTransactions = true"
            class="text-xs text-blue-600 hover:text-blue-700"
          >
            +{{ recentTransactions.length - 3 }} تراکنش دیگر
          </button>
        </div>
      </div>
    </div>

    <!-- جدول کامل تراکنش‌ها -->
    <div v-else class="space-y-4">
      <!-- فیلترهای سریع -->
      <div class="flex items-center space-x-2 space-x-reverse">
        <select
          v-model="filters.status"
          @change="fetchTransactions"
          class="px-2 py-1 text-xs border border-gray-300 rounded focus:ring-1 focus:ring-blue-500"
        >
          <option value="">همه وضعیت‌ها</option>
          <option value="success">موفق</option>
          <option value="failed">ناموفق</option>
          <option value="pending">در انتظار</option>
        </select>
        
        <button
          @click="fetchTransactions"
          :disabled="loading"
          class="px-2 py-1 text-xs bg-blue-600 text-white rounded hover:bg-blue-700 disabled:opacity-50"
        >
          <Icon 
            :name="loading ? 'heroicons:arrow-path' : 'heroicons:arrow-path'" 
            :class="['w-3 h-3', loading ? 'animate-spin' : '']" 
          />
        </button>
      </div>

      <!-- جدول تراکنش‌ها -->
      <div class="overflow-x-auto">
        <table class="min-w-full text-xs">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-2 py-1 text-right font-medium text-gray-500">شناسه</th>
              <th class="px-2 py-1 text-right font-medium text-gray-500">مبلغ</th>
              <th class="px-2 py-1 text-right font-medium text-gray-500">وضعیت</th>
              <th class="px-2 py-1 text-right font-medium text-gray-500">تاریخ</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="transaction in recentTransactions" :key="transaction.id" class="hover:bg-gray-50">
              <td class="px-2 py-1 text-gray-900 font-medium">{{ transaction.transaction_id }}</td>
              <td class="px-2 py-1 text-gray-900">{{ formatCurrency(transaction.amount) }}</td>
              <td class="px-2 py-1">
                <span
                  :class="[
                    'px-1 py-0.5 text-xs rounded-full',
                    getStatusClass(transaction.status)
                  ]"
                >
                  {{ getStatusText(transaction.status) }}
                </span>
              </td>
              <td class="px-2 py-1 text-gray-500">{{ formatDate(transaction.created_at) }}</td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- صفحه‌بندی ساده -->
      <div v-if="total > filters.limit" class="flex items-center justify-between text-xs">
        <span class="text-gray-500">
          نمایش {{ (currentPage - 1) * filters.limit + 1 }} تا {{ Math.min(currentPage * filters.limit, total) }} از {{ total }}
        </span>
        <div class="flex items-center space-x-1 space-x-reverse">
          <button
            @click="changePage(currentPage - 1)"
            :disabled="currentPage === 1"
            class="px-2 py-1 border border-gray-300 rounded hover:bg-gray-50 disabled:opacity-50"
          >
            قبلی
          </button>
          <span class="px-2 py-1">{{ currentPage }}</span>
          <button
            @click="changePage(currentPage + 1)"
            :disabled="currentPage >= totalPages"
            class="px-2 py-1 border border-gray-300 rounded hover:bg-gray-50 disabled:opacity-50"
          >
            بعدی
          </button>
        </div>
      </div>
    </div>

    <!-- مودال مشاهده تراکنش -->
    <div v-if="selectedTransaction" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg p-6 max-w-md w-full mx-4">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold">جزئیات تراکنش</h3>
          <button @click="selectedTransaction = null" class="text-gray-400 hover:text-gray-600">
            <Icon name="heroicons:x-mark" class="w-5 h-5" />
          </button>
        </div>
        
        <div class="space-y-3">
          <div class="grid grid-cols-2 gap-3 text-sm">
            <div>
              <label class="block text-xs font-medium text-gray-700">شناسه تراکنش</label>
              <p class="text-gray-900">{{ selectedTransaction.transaction_id }}</p>
            </div>
            <div>
              <label class="block text-xs font-medium text-gray-700">مبلغ</label>
              <p class="text-gray-900">{{ formatCurrency(selectedTransaction.amount) }}</p>
            </div>
            <div>
              <label class="block text-xs font-medium text-gray-700">وضعیت</label>
              <p class="text-gray-900">
                <span :class="getStatusClass(selectedTransaction.status)">
                  {{ getStatusText(selectedTransaction.status) }}
                </span>
              </p>
            </div>
            <div>
              <label class="block text-xs font-medium text-gray-700">تاریخ</label>
              <p class="text-gray-900">{{ formatDate(selectedTransaction.created_at) }}</p>
            </div>
          </div>
          
          <div v-if="selectedTransaction.description" class="text-sm">
            <label class="block text-xs font-medium text-gray-700">توضیحات</label>
            <p class="text-gray-900">{{ selectedTransaction.description }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'

// تعریف interface ها
interface Transaction {
  id: number
  transaction_id: string
  order_id: string
  amount: number
  status: string
  description?: string
  created_at: string
  updated_at: string
}

// Props
const props = defineProps<{
  gatewayId: number
}>()

// متغیرهای reactive
const recentTransactions = ref<Transaction[]>([])
const loading = ref(false)
const selectedTransaction = ref<Transaction | null>(null)
const showAllTransactions = ref(false)
const currentPage = ref(1)
const total = ref(0)
const totalPages = ref(0)

// فیلترها
const filters = ref({
  status: '',
  limit: 10
})

// Methods
const fetchTransactions = async () => {
  try {
    loading.value = true
    const response: any = await $fetch('/api/payments/admin/transactions', {
      query: {
        gateway_id: props.gatewayId,
        page: currentPage.value,
        limit: filters.value.limit,
        status: filters.value.status
      }
    })
    
    if (response.data) {
      recentTransactions.value = response.data
      total.value = response.total || 0
      totalPages.value = Math.ceil(total.value / filters.value.limit)
    }
  } catch (error) {
    console.error('خطا در دریافت تراکنش‌ها:', error)
  } finally {
    loading.value = false
  }
}

const changePage = (page: number) => {
  if (page >= 1 && page <= totalPages.value) {
    currentPage.value = page
    fetchTransactions()
  }
}

const viewTransaction = (transaction: Transaction) => {
  selectedTransaction.value = transaction
}

const getStatusClass = (status: string) => {
  switch (status) {
    case 'success':
      return 'bg-green-100 text-green-800'
    case 'failed':
      return 'bg-red-100 text-red-800'
    case 'pending':
      return 'bg-yellow-100 text-yellow-800'
    case 'refunded':
      return 'bg-purple-100 text-purple-800'
    case 'cancelled':
      return 'bg-gray-100 text-gray-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getStatusColor = (status: string) => {
  switch (status) {
    case 'success':
      return 'bg-green-500'
    case 'failed':
      return 'bg-red-500'
    case 'pending':
      return 'bg-yellow-500'
    case 'refunded':
      return 'bg-purple-500'
    case 'cancelled':
      return 'bg-gray-500'
    default:
      return 'bg-gray-500'
  }
}

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
    case 'cancelled':
      return 'لغو شده'
    default:
      return 'نامشخص'
  }
}

const formatCurrency = (amount: number) => {
  if (!amount) return '0 تومان'
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

const formatDate = (dateString: string) => {
  if (!dateString) return '-'
  const date = new Date(dateString)
  return new Intl.DateTimeFormat('fa-IR', {
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  }).format(date)
}

// Lifecycle
onMounted(() => {
  fetchTransactions()
})
</script> 
