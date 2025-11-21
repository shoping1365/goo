<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200">
    <!-- Header -->
    <div class="px-6 py-4 border-b border-gray-200">
      <div class="flex justify-between items-center">
        <h3 class="text-lg font-semibold text-gray-900">لیست گیفت کارت‌ها</h3>
      </div>
    </div>

    <!-- جدول -->
    <div class="overflow-x-auto">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              <div class="flex items-center space-x-1 space-x-reverse">
                <span>کد کارت</span>
                <button class="text-gray-400 hover:text-gray-600" @click="sortBy('code')">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16V4m0 0L3 8m4-4l4 4m6 0v12m0 0l4-4m-4 4l-4-4" />
                  </svg>
                </button>
              </div>
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              گیرنده
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              <div class="flex items-center space-x-1 space-x-reverse">
                <span>مبلغ</span>
                <button class="text-gray-400 hover:text-gray-600" @click="sortBy('amount')">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16V4m0 0L3 8m4-4l4 4m6 0v12m0 0l4-4m-4 4l-4-4" />
                  </svg>
                </button>
              </div>
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              وضعیت
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              نوع
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              <div class="flex items-center space-x-1 space-x-reverse">
                <span>تاریخ انقضا</span>
                <button class="text-gray-400 hover:text-gray-600" @click="sortBy('expiryDate')">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16V4m0 0L3 8m4-4l4 4m6 0v12m0 0l4-4m-4 4l-4-4" />
                  </svg>
                </button>
              </div>
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              عملیات
            </th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-for="giftCard in filteredAndSortedGiftCards" :key="giftCard.id" class="hover:bg-gray-50">
            <!-- کد کارت -->
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="flex items-center">
                <div class="">
                  <div class="h-10 w-10 rounded-lg bg-gradient-to-r from-blue-500 to-purple-600 flex items-center justify-center">
                    <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z" />
                    </svg>
                  </div>
                </div>
                <div class="mr-4">
                  <div class="text-sm font-medium text-gray-900 font-mono">{{ giftCard.code }}</div>
                  <div class="text-sm text-gray-500">{{ formatDate(giftCard.createdAt) }}</div>
                </div>
              </div>
            </td>

            <!-- گیرنده -->
            <td class="px-6 py-4 whitespace-nowrap">
              <div>
                <div class="text-sm font-medium text-gray-900">{{ giftCard.recipientName || giftCard.recipient }}</div>
                <div class="text-sm text-gray-500">{{ giftCard.recipientEmail }}</div>
                <div v-if="giftCard.recipientPhone" class="text-xs text-gray-400">{{ giftCard.recipientPhone }}</div>
              </div>
            </td>

            <!-- مبلغ -->
            <td class="px-6 py-4 whitespace-nowrap">
              <div>
                <div class="text-sm font-medium text-gray-900">{{ formatAmount(giftCard.amount) }} تومان</div>
                <div v-if="giftCard.remainingAmount !== giftCard.amount" class="text-sm text-gray-500">
                  باقی‌مانده: {{ formatAmount(giftCard.remainingAmount) }} تومان
                </div>
                <div v-if="giftCard.usageHistory && giftCard.usageHistory.length > 0" class="text-xs text-blue-600">
                  {{ giftCard.usageHistory.length }} بار استفاده شده
                </div>
              </div>
            </td>

            <!-- وضعیت -->
            <td class="px-6 py-4 whitespace-nowrap">
              <span :class="getStatusClasses(giftCard.status)" class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium">
                <span class="w-2 h-2 rounded-full ml-1" :class="getStatusDotClasses(giftCard.status)"></span>
                {{ getStatusText(giftCard.status) }}
              </span>
              <div v-if="giftCard.status === 'expired'" class="text-xs text-red-600 mt-1">
                {{ getDaysUntilExpiry(giftCard.expiryDate) }} روز گذشته
              </div>
              <div v-else-if="giftCard.status === 'active'" class="text-xs text-green-600 mt-1">
                {{ getDaysUntilExpiry(giftCard.expiryDate) }} روز باقی‌مانده
              </div>
            </td>

            <!-- نوع -->
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="flex items-center">
                <span :class="getTypeClasses(giftCard.type)" class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium">
                  {{ getTypeText(giftCard.type) }}
                </span>
                <div v-if="giftCard.category" class="mr-2">
                  <span class="text-xs text-gray-500">{{ getCategoryText(giftCard.category) }}</span>
                </div>
              </div>
            </td>

            <!-- تاریخ انقضا -->
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
              {{ formatDate(giftCard.expiryDate) }}
            </td>

            <!-- عملیات -->
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
              <div class="flex items-center space-x-2 space-x-reverse">
                <!-- مشاهده جزئیات -->
                <button 
                  class="text-blue-600 hover:text-blue-900 p-1 rounded hover:bg-blue-50"
                  title="مشاهده جزئیات"
                  @click="$emit('view-details', giftCard)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                  </svg>
                </button>

                <!-- ویرایش -->
                <button 
                  class="text-green-600 hover:text-green-900 p-1 rounded hover:bg-green-50"
                  title="ویرایش"
                  @click="$emit('edit-card', giftCard)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                  </svg>
                </button>

                <!-- کپی کد -->
                <button 
                  class="text-purple-600 hover:text-purple-900 p-1 rounded hover:bg-purple-50"
                  title="کپی کد"
                  @click="copyCode(giftCard.code)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z" />
                  </svg>
                </button>

                <!-- ارسال مجدد -->
                <button 
                  v-if="giftCard.deliveryMethod !== 'manual'"
                  class="text-orange-600 hover:text-orange-900 p-1 rounded hover:bg-orange-50"
                  title="ارسال مجدد"
                  @click="resendGiftCard(giftCard)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 4.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                  </svg>
                </button>

                <!-- فعال/غیرفعال -->
                <button 
                  v-if="giftCard.status === 'active'"
                  class="text-yellow-600 hover:text-yellow-900 p-1 rounded hover:bg-yellow-50"
                  title="غیرفعال کردن"
                  @click="deactivateGiftCard(giftCard)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728L5.636 5.636m12.728 12.728L18.364 5.636M5.636 18.364l12.728-12.728" />
                  </svg>
                </button>
                <button 
                  v-else
                  class="text-green-600 hover:text-green-900 p-1 rounded hover:bg-green-50"
                  title="فعال کردن"
                  @click="activateGiftCard(giftCard)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                  </svg>
                </button>

                <!-- حذف -->
                <button 
                  v-if="canDeleteGiftCard"
                  class="text-red-600 hover:text-red-900 p-1 rounded hover:bg-red-50"
                  title="حذف"
                  @click="confirmDelete(giftCard)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                  </svg>
                </button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Pagination -->
    <div class="bg-white px-4 py-3 flex items-center justify-between border-t border-gray-200 sm:px-6">
      <div class="flex-1 flex justify-between sm:hidden">
        <button 
          :disabled="currentPage === 1"
          class="relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50"
          @click="previousPage"
        >
          قبلی
        </button>
        <button 
          :disabled="currentPage >= totalPages"
          class="mr-3 relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50"
          @click="nextPage"
        >
          بعدی
        </button>
      </div>
      <div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
        <div>
          <p class="text-sm text-gray-700">
            نمایش 
            <span class="font-medium">{{ startIndex + 1 }}</span>
            تا 
            <span class="font-medium">{{ endIndex }}</span>
            از 
            <span class="font-medium">{{ totalItems }}</span>
            نتیجه
          </p>
        </div>
        <div>
          <nav class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px" aria-label="Pagination">
            <button 
              :disabled="currentPage === 1"
              class="relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50"
              @click="previousPage"
            >
              <span class="sr-only">قبلی</span>
              <svg class="h-5 w-5" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd" />
              </svg>
            </button>
            <button 
              v-for="page in visiblePages" 
              :key="page"
              :class="[
                page === currentPage 
                  ? 'z-10 bg-blue-50 border-blue-500 text-blue-600' 
                  : 'bg-white border-gray-300 text-gray-500 hover:bg-gray-50',
                'relative inline-flex items-center px-4 py-2 border text-sm font-medium'
              ]"
              @click="goToPage(page)"
            >
              {{ page }}
            </button>
            <button 
              :disabled="currentPage >= totalPages"
              class="relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50"
              @click="nextPage"
            >
              <span class="sr-only">بعدی</span>
              <svg class="h-5 w-5" fill="currentColor" viewBox="0 0 20 20">
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
import { computed, onMounted, ref } from 'vue';
import { useAuth } from '~/composables/useAuth';

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { hasPermission } = useAuth()

// Computed برای چک کردن پرمیژن حذف
const canDeleteGiftCard = computed(() => hasPermission('giftcard.delete'))

interface GiftCard {
  id?: number | string
  code?: string
  amount?: number
  [key: string]: unknown
}

// Props
const props = defineProps<{
  giftCards: GiftCard[]
}>()

// Emits
const emit = defineEmits<{
  'edit-card': [giftCard: GiftCard]
  'delete-card': [giftCard: GiftCard]
  'view-details': [giftCard: GiftCard]
}>()

// Reactive data
const searchQuery = ref('')
const currentPage = ref(1)
const itemsPerPage = ref(10)
const sortField = ref('createdAt')
const sortDirection = ref('desc')

const filters = ref({
  status: '',
  type: '',
  category: ''
})

// Computed properties
const filteredGiftCards = computed(() => {
  let filtered = props.giftCards

  // فیلتر بر اساس جستجو
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(card => 
      card.code.toLowerCase().includes(query) ||
      (card.recipientName || card.recipient || '').toLowerCase().includes(query) ||
      card.recipientEmail.toLowerCase().includes(query)
    )
  }

  // فیلتر بر اساس وضعیت
  if (filters.value.status) {
    filtered = filtered.filter(card => card.status === filters.value.status)
  }

  // فیلتر بر اساس نوع
  if (filters.value.type) {
    filtered = filtered.filter(card => card.type === filters.value.type)
  }

  // فیلتر بر اساس دسته‌بندی
  if (filters.value.category) {
    filtered = filtered.filter(card => card.category === filters.value.category)
  }

  return filtered
})

const sortedGiftCards = computed(() => {
  const sorted = [...filteredGiftCards.value]
  
  sorted.sort((a, b) => {
    let aValue = a[sortField.value]
    let bValue = b[sortField.value]

    // تبدیل تاریخ‌ها
    if (sortField.value === 'createdAt' || sortField.value === 'expiryDate') {
      aValue = new Date(aValue).getTime()
      bValue = new Date(bValue).getTime()
    }

    // تبدیل اعداد
    if (sortField.value === 'amount') {
      aValue = Number(aValue)
      bValue = Number(bValue)
    }

    if (sortDirection.value === 'asc') {
      return aValue > bValue ? 1 : -1
    } else {
      return aValue < bValue ? 1 : -1
    }
  })

  return sorted
})

const paginatedGiftCards = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage.value
  const end = start + itemsPerPage.value
  return sortedGiftCards.value.slice(start, end)
})

const filteredAndSortedGiftCards = computed(() => paginatedGiftCards.value)

const totalItems = computed(() => filteredGiftCards.value.length)
const totalPages = computed(() => Math.ceil(totalItems.value / itemsPerPage.value))
const startIndex = computed(() => (currentPage.value - 1) * itemsPerPage.value)
const endIndex = computed(() => Math.min(startIndex.value + itemsPerPage.value, totalItems.value))

const visiblePages = computed(() => {
  const pages = []
  const maxVisible = 5
  let start = Math.max(1, currentPage.value - Math.floor(maxVisible / 2))
  const end = Math.min(totalPages.value, start + maxVisible - 1)
  
  if (end - start + 1 < maxVisible) {
    start = Math.max(1, end - maxVisible + 1)
  }
  
  for (let i = start; i <= end; i++) {
    pages.push(i)
  }
  
  return pages
})

// Methods
const formatAmount = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount)
}

const formatDate = (date: string | Date) => {
  return new Date(date).toLocaleDateString('fa-IR', {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  })
}

const getStatusText = (status: string) => {
  const statusMap = {
    active: 'فعال',
    used: 'استفاده شده',
    expired: 'منقضی شده',
    inactive: 'غیرفعال'
  }
  return statusMap[status] || status
}

const getStatusClasses = (status: string) => {
  const classesMap = {
    active: 'bg-green-100 text-green-800',
    used: 'bg-blue-100 text-blue-800',
    expired: 'bg-red-100 text-red-800',
    inactive: 'bg-gray-100 text-gray-800'
  }
  return classesMap[status] || 'bg-gray-100 text-gray-800'
}

const getStatusDotClasses = (status: string) => {
  const classesMap = {
    active: 'bg-green-400',
    used: 'bg-blue-400',
    expired: 'bg-red-400',
    inactive: 'bg-gray-400'
  }
  return classesMap[status] || 'bg-gray-400'
}

const getTypeText = (type: string) => {
  const typeMap = {
    digital: 'دیجیتال',
    physical: 'فیزیکی',
    hybrid: 'ترکیبی'
  }
  return typeMap[type] || type
}

const getTypeClasses = (type: string) => {
  const classesMap = {
    digital: 'bg-purple-100 text-purple-800',
    physical: 'bg-orange-100 text-orange-800',
    hybrid: 'bg-indigo-100 text-indigo-800'
  }
  return classesMap[type] || 'bg-gray-100 text-gray-800'
}

const getCategoryText = (category: string) => {
  const categoryMap = {
    birthday: 'تولد',
    wedding: 'عروسی',
    newyear: 'سال نو',
    corporate: 'شرکتی',
    discount: 'تخفیف ویژه',
    general: 'عمومی'
  }
  return categoryMap[category] || category
}

const getDaysUntilExpiry = (expiryDate: string | Date) => {
  const now = new Date()
  const expiry = new Date(expiryDate)
  const diffTime = expiry.getTime() - now.getTime()
  const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24))
  return Math.abs(diffDays)
}

const sortBy = (field: string) => {
  if (sortField.value === field) {
    sortDirection.value = sortDirection.value === 'asc' ? 'desc' : 'asc'
  } else {
    sortField.value = field
    sortDirection.value = 'asc'
  }
}

const copyCode = async (code: string) => {
  try {
    await navigator.clipboard.writeText(code)
    // نمایش پیام موفقیت
    alert('کد کپی شد')
  } catch (err) {
    console.error('خطا در کپی کردن کد:', err)
  }
}

interface GiftCard {
  id?: number | string
  code?: string
  status?: string
  [key: string]: unknown
}

const resendGiftCard = async (giftCard: GiftCard) => {
  if (confirm(`آیا می‌خواهید گیفت کارت ${giftCard.code} مجدداً ارسال شود؟`)) {
    try {
      // در نسخه واقعی: ارسال به API
      alert('گیفت کارت با موفقیت ارسال شد')
    } catch (error) {
      console.error('خطا در ارسال مجدد:', error)
      alert('خطا در ارسال مجدد گیفت کارت')
    }
  }
}

const activateGiftCard = async (giftCard: GiftCard) => {
  if (confirm(`آیا می‌خواهید گیفت کارت ${giftCard.code} فعال شود؟`)) {
    try {
      // در نسخه واقعی: ارسال به API
      giftCard.status = 'active'
      alert('گیفت کارت فعال شد')
    } catch (error) {
      console.error('خطا در فعال‌سازی:', error)
      alert('خطا در فعال‌سازی گیفت کارت')
    }
  }
}

const deactivateGiftCard = async (giftCard: GiftCard) => {
  if (confirm(`آیا می‌خواهید گیفت کارت ${giftCard.code} غیرفعال شود؟`)) {
    try {
      // در نسخه واقعی: ارسال به API
      giftCard.status = 'inactive'
      alert('گیفت کارت غیرفعال شد')
    } catch (error) {
      console.error('خطا در غیرفعال‌سازی:', error)
      alert('خطا در غیرفعال‌سازی گیفت کارت')
    }
  }
}

const confirmDelete = (giftCard: GiftCard) => {
  if (confirm(`آیا از حذف گیفت کارت ${giftCard.code} مطمئن هستید؟ این عملیات غیرقابل بازگشت است.`)) {
    emit('delete-card', giftCard)
  }
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
  // تنظیمات اولیه
})
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 
