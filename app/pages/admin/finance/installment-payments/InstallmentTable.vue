<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200">
    <div class="overflow-x-auto">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مشتری</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">محصول</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مبلغ</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">اقساط</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">امتیاز اعتباری</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-for="installment in installments" :key="installment.id">
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="text-sm font-medium text-gray-900">{{ installment.customer_name }}</div>
              <div class="text-sm text-gray-500">{{ installment.national_id }}</div>
              <div class="text-sm text-gray-500">{{ installment.mobile }}</div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="text-sm text-gray-900">{{ installment.product_name }}</div>
              <div class="text-sm text-gray-500">{{ installment.product_sku }}</div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="text-sm text-gray-900">{{ formatCurrency(installment.amount) }}</div>
              <div class="text-sm text-gray-500">{{ formatCurrency(installment.monthly_payment) }} ماهانه</div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="text-sm text-gray-900">{{ installment.installment_count }} قسط</div>
              <div class="text-sm text-gray-500">{{ installment.installment_count }} ماه</div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span class="px-2 py-1 text-xs font-medium rounded-full" :class="getStatusClass(installment.status)">
                {{ getStatusText(installment.status) }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="text-sm text-gray-900">{{ installment.credit_score }}</div>
              <div class="text-sm text-gray-500" :class="getCreditScoreColor(installment.credit_score)">
                {{ getCreditScoreText(installment.credit_score) }}
              </div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="text-sm text-gray-900">{{ formatDate(installment.created_at) }}</div>
              <div class="text-sm text-gray-500">{{ formatTime(installment.created_at) }}</div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
              <button
                class="text-blue-600 hover:text-blue-900 ml-3"
                @click="viewDetails(installment)"
              >
                جزئیات
              </button>
              <button
                v-if="installment.status === 'pending'"
                class="text-green-600 hover:text-green-900 ml-3"
                @click="approveInstallment(installment.id)"
              >
                تایید
              </button>
              <button
                v-if="installment.status === 'pending'"
                class="text-red-600 hover:text-red-900"
                @click="rejectInstallment(installment.id)"
              >
                رد
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Pagination -->
    <div v-if="pagination" class="bg-white px-4 py-3 flex items-center justify-between border-t border-gray-200 sm:px-6">
      <div class="flex-1 flex justify-between sm:hidden">
        <button
          :disabled="pagination.page <= 1"
          class="relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
          @click="changePage(pagination.page - 1)"
        >
          قبلی
        </button>
        <button
          :disabled="pagination.page >= pagination.totalPages"
          class="mr-3 relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
          @click="changePage(pagination.page + 1)"
        >
          بعدی
        </button>
      </div>
      <div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
        <div>
          <p class="text-sm text-gray-700">
            نمایش
            <span class="font-medium">{{ (pagination.page - 1) * pagination.limit + 1 }}</span>
            تا
            <span class="font-medium">{{ Math.min(pagination.page * pagination.limit, pagination.total) }}</span>
            از
            <span class="font-medium">{{ pagination.total }}</span>
            نتیجه
          </p>
        </div>
        <div>
          <nav class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px" aria-label="Pagination">
            <button
              :disabled="pagination.page <= 1"
              class="relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
              @click="changePage(pagination.page - 1)"
            >
              <span class="sr-only">قبلی</span>
              <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                <path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd" />
              </svg>
            </button>
            
            <template v-for="page in getPageNumbers()" :key="page">
              <button
                v-if="page !== '...'"
                :class="[
                  page === pagination.page
                    ? 'z-10 bg-blue-50 border-blue-500 text-blue-600'
                    : 'bg-white border-gray-300 text-gray-500 hover:bg-gray-50',
                  'relative inline-flex items-center px-4 py-2 border text-sm font-medium'
                ]"
                @click="changePage(page)"
              >
                {{ page }}
              </button>
              <span
                v-else
                class="relative inline-flex items-center px-4 py-2 border border-gray-300 bg-white text-sm font-medium text-gray-700"
              >
                ...
              </span>
            </template>
            
            <button
              :disabled="pagination.page >= pagination.totalPages"
              class="relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
              @click="changePage(pagination.page + 1)"
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
</template>

<script setup lang="ts">
import { } from 'vue'

interface Installment {
  id: string
  customer_name: string
  national_id: string
  mobile: string
  product_name: string
  product_sku: string
  amount: number
  monthly_payment: number
  installment_count: number
  status: string
  credit_score: number
  created_at: string
}

interface Pagination {
  page: number
  limit: number
  total: number
  totalPages: number
}

const props = defineProps<{
  installments: Installment[]
  pagination?: Pagination
}>()

const emit = defineEmits<{
  pageChange: [page: number]
  viewDetails: [installment: Installment]
  approve: [id: string]
  reject: [id: string]
}>()

// Utility functions
const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

const formatDate = (date: string) => {
  return new Date(date).toLocaleDateString('fa-IR')
}

const formatTime = (date: string) => {
  return new Date(date).toLocaleTimeString('fa-IR')
}

const getStatusClass = (status: string) => {
  switch (status) {
    case 'approved':
      return 'bg-green-100 text-green-800'
    case 'pending':
      return 'bg-yellow-100 text-yellow-800'
    case 'rejected':
      return 'bg-red-100 text-red-800'
    case 'completed':
      return 'bg-blue-100 text-blue-800'
    case 'overdue':
      return 'bg-orange-100 text-orange-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getStatusText = (status: string) => {
  switch (status) {
    case 'approved':
      return 'تایید شده'
    case 'pending':
      return 'در انتظار'
    case 'rejected':
      return 'رد شده'
    case 'completed':
      return 'تکمیل شده'
    case 'overdue':
      return 'معوق'
    default:
      return 'نامشخص'
  }
}

const getCreditScoreColor = (score: number) => {
  if (score >= 700) return 'text-green-600'
  if (score >= 500) return 'text-yellow-600'
  return 'text-red-600'
}

const getCreditScoreText = (score: number) => {
  if (score >= 700) return 'عالی'
  if (score >= 500) return 'متوسط'
  return 'ضعیف'
}

const getPageNumbers = () => {
  if (!props.pagination) return []
  
  const { page, totalPages } = props.pagination
  const pages = []
  
  if (totalPages <= 7) {
    for (let i = 1; i <= totalPages; i++) {
      pages.push(i)
    }
  } else {
    if (page <= 4) {
      for (let i = 1; i <= 5; i++) {
        pages.push(i)
      }
      pages.push('...')
      pages.push(totalPages)
    } else if (page >= totalPages - 3) {
      pages.push(1)
      pages.push('...')
      for (let i = totalPages - 4; i <= totalPages; i++) {
        pages.push(i)
      }
    } else {
      pages.push(1)
      pages.push('...')
      for (let i = page - 1; i <= page + 1; i++) {
        pages.push(i)
      }
      pages.push('...')
      pages.push(totalPages)
    }
  }
  
  return pages
}

// Event handlers
const changePage = (page: number) => {
  emit('pageChange', page)
}

const viewDetails = (installment: Installment) => {
  emit('viewDetails', installment)
}

const approveInstallment = (id: string) => {
  emit('approve', id)
}

const rejectInstallment = (id: string) => {
  emit('reject', id)
}
</script> 
