<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex justify-between items-center">
        <div>
          <h2 class="text-xl font-bold text-gray-900">لیست گیفت کارت‌ها</h2>
          <p class="text-gray-600 mt-1">مدیریت و مشاهده تمام گیفت کارت‌های سیستم</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button 
            class="px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2"
            @click="exportToExcel"
          >
            <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
            </svg>
            خروجی Excel
          </button>
          <button 
            class="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
            @click="refreshData"
          >
            <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
            </svg>
            به‌روزرسانی
          </button>
        </div>
      </div>
    </div>

    <!-- فیلترها و جستجو -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <!-- جستجو -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">جستجو</label>
          <input
            v-model="searchQuery"
            type="text"
            placeholder="جستجو بر اساس کد، خریدار، گیرنده..."
            class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
          />
        </div>
        
        <!-- فیلتر وضعیت -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
          <select
            v-model="statusFilter"
            class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه وضعیت‌ها</option>
            <option value="active">فعال</option>
            <option value="used">استفاده شده</option>
            <option value="expired">منقضی شده</option>
            <option value="suspended">معلق</option>
            <option value="cancelled">لغو شده</option>
          </select>
        </div>
        
        <!-- فیلتر نوع -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نوع کارت</label>
          <select
            v-model="typeFilter"
            class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه انواع</option>
            <option value="digital">دیجیتال</option>
            <option value="physical">فیزیکی</option>
            <option value="virtual">مجازی</option>
          </select>
        </div>
        
        <!-- فیلتر تاریخ -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ ایجاد</label>
          <select
            v-model="dateFilter"
            class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه تاریخ‌ها</option>
            <option value="today">امروز</option>
            <option value="week">هفته جاری</option>
            <option value="month">ماه جاری</option>
            <option value="last30">30 روز گذشته</option>
          </select>
        </div>
      </div>
    </div>

    <!-- جدول -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th 
                v-for="column in columns" 
                :key="column.key"
                class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider cursor-pointer hover:bg-gray-100"
                :class="{ 'bg-blue-50': sortKey === column.key }"
                @click="sortBy(column.key)"
              >
                <div class="flex items-center justify-between">
                  <span>{{ column.label }}</span>
                  <div class="flex flex-col">
                    <svg 
                      class="w-3 h-3 text-gray-400" 
                      :class="{ 'text-blue-600': sortKey === column.key && sortOrder === 'asc' }"
                      fill="currentColor" 
                      viewBox="0 0 20 20"
                    >
                      <path fill-rule="evenodd" d="M14.707 12.707a1 1 0 01-1.414 0L10 9.414l-3.293 3.293a1 1 0 01-1.414-1.414l4-4a1 1 0 011.414 0l4 4a1 1 0 010 1.414z" clip-rule="evenodd"></path>
                    </svg>
                    <svg 
                      class="w-3 h-3 text-gray-400 -mt-1" 
                      :class="{ 'text-blue-600': sortKey === column.key && sortOrder === 'desc' }"
                      fill="currentColor" 
                      viewBox="0 0 20 20"
                    >
                      <path fill-rule="evenodd" d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z" clip-rule="evenodd"></path>
                    </svg>
                  </div>
                </div>
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                عملیات
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="card in paginatedCards" :key="card.id" class="hover:bg-gray-50">
              <!-- کد کارت -->
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="">
                    <div class="h-10 w-10 rounded-lg bg-gradient-to-br from-blue-500 to-purple-600 flex items-center justify-center">
                      <span class="text-white text-sm font-bold">{{ card.code.substring(0, 4) }}</span>
                    </div>
                  </div>
                  <div class="mr-4">
                    <div class="text-sm font-medium text-gray-900">{{ card.code }}</div>
                    <div class="text-sm text-gray-500">{{ card.type }}</div>
                  </div>
                </div>
              </td>
              
              <!-- خریدار -->
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-900">{{ String(card.buyerName || '') }}</div>
                <div class="text-sm text-gray-500">{{ String(card.buyerEmail || '') }}</div>
              </td>
              
              <!-- گیرنده -->
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-900">{{ String(card.recipientName || '') }}</div>
                <div class="text-sm text-gray-500">{{ String(card.recipientEmail || '') }}</div>
              </td>
              
              <!-- مبلغ -->
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm font-medium text-gray-900">{{ formatCurrency(Number(card.amount || 0)) }}</div>
                <div class="text-sm text-gray-500">{{ formatCurrency(Number(card.remainingAmount || 0)) }} باقی‌مانده</div>
              </td>
              
              <!-- تاریخ ایجاد -->
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-900">{{ formatDate(String(card.createdAt || '')) }}</div>
                <div class="text-sm text-gray-500">{{ formatTime(String(card.createdAt || '')) }}</div>
              </td>
              
              <!-- تاریخ انقضا -->
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-900">{{ formatDate(String(card.expiresAt || '')) }}</div>
                <div class="text-sm text-gray-500">{{ getDaysUntilExpiry(String(card.expiresAt || '')) }}</div>
              </td>
              
              <!-- وضعیت -->
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusClass(String(card.status || ''))" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                  {{ getStatusLabel(String(card.status || '')) }}
                </span>
                <div v-if="Number(card.usageCount || 0) > 0" class="text-xs text-gray-500 mt-1">
                  {{ Number(card.usageCount || 0) }} بار استفاده
                </div>
              </td>
              
              <!-- عملیات -->
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <div class="flex items-center space-x-2 space-x-reverse">
                  <button 
                    class="text-blue-600 hover:text-blue-900"
                    title="مشاهده جزئیات"
                    @click="viewCard(card)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                    </svg>
                  </button>
                  
                  <button 
                    class="text-green-600 hover:text-green-900"
                    title="ویرایش"
                    @click="editCard(card)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
                    </svg>
                  </button>
                  
                  <button 
                    class="text-purple-600 hover:text-purple-900"
                    title="کپی"
                    @click="duplicateCard(card)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
                    </svg>
                  </button>
                  
                  <button 
                    v-if="canDeleteGiftCard"
                    class="text-red-600 hover:text-red-900"
                    title="حذف"
                    @click="deleteCard(card)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      
      <!-- صفحه‌بندی -->
      <div class="bg-white px-4 py-3 flex items-center justify-between border-t border-gray-200 sm:px-6">
        <div class="flex-1 flex justify-between sm:hidden">
          <button 
            :disabled="currentPage === 1"
            class="relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
            @click="previousPage"
          >
            قبلی
          </button>
          <button 
            :disabled="currentPage === totalPages"
            class="mr-3 relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
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
              <span class="font-medium">{{ filteredCards.length }}</span>
              نتیجه
            </p>
          </div>
          <div>
            <nav class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px" aria-label="Pagination">
              <button 
                :disabled="currentPage === 1"
                class="relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
                @click="previousPage"
              >
                <span class="sr-only">قبلی</span>
                <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
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
                :disabled="currentPage === totalPages"
                class="relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
                @click="nextPage"
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
  'view-card': [card: GiftCard]
  'edit-card': [card: GiftCard]
  'duplicate-card': [card: GiftCard]
  'delete-card': [card: GiftCard]
}>()

// Reactive data
const searchQuery = ref('')
const statusFilter = ref('')
const typeFilter = ref('')
const dateFilter = ref('')
const sortKey = ref('createdAt')
const sortOrder = ref('desc')
const currentPage = ref(1)
const itemsPerPage = ref(10)

const columns = [
  { key: 'code', label: 'کد کارت' },
  { key: 'buyerName', label: 'خریدار' },
  { key: 'recipientName', label: 'گیرنده' },
  { key: 'amount', label: 'مبلغ' },
  { key: 'createdAt', label: 'تاریخ ایجاد' },
  { key: 'expiresAt', label: 'تاریخ انقضا' },
  { key: 'status', label: 'وضعیت' }
]

// Computed properties
const filteredCards = computed(() => {
  let filtered = [...props.giftCards]
  
  // فیلتر جستجو
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(card => 
      String(card.code || '').toLowerCase().includes(query) ||
      String(card.buyerName || '').toLowerCase().includes(query) ||
      String(card.buyerEmail || '').toLowerCase().includes(query) ||
      String(card.recipientName || '').toLowerCase().includes(query) ||
      String(card.recipientEmail || '').toLowerCase().includes(query)
    )
  }
  
  // فیلتر وضعیت
  if (statusFilter.value) {
    filtered = filtered.filter(card => card.status === statusFilter.value)
  }
  
  // فیلتر نوع
  if (typeFilter.value) {
    filtered = filtered.filter(card => card.type === typeFilter.value)
  }
  
  // فیلتر تاریخ
  if (dateFilter.value) {
    const today = new Date()
    const cardDate = new Date()
    
    switch (dateFilter.value) {
      case 'today':
        filtered = filtered.filter(card => {
          cardDate.setTime(new Date(String(card.createdAt || '')).getTime())
          return cardDate.toDateString() === today.toDateString()
        })
        break
      case 'week':
        const weekAgo = new Date(today.getTime() - 7 * 24 * 60 * 60 * 1000)
        filtered = filtered.filter(card => new Date(String(card.createdAt || '')).getTime() >= weekAgo.getTime())
        break
      case 'month':
        const monthAgo = new Date(today.getTime() - 30 * 24 * 60 * 60 * 1000)
        filtered = filtered.filter(card => new Date(String(card.createdAt || '')).getTime() >= monthAgo.getTime())
        break
      case 'last30':
        const thirtyDaysAgo = new Date(today.getTime() - 30 * 24 * 60 * 60 * 1000)
        filtered = filtered.filter(card => new Date(String(card.createdAt || '')).getTime() >= thirtyDaysAgo.getTime())
        break
    }
  }
  
  // مرتب‌سازی
  filtered.sort((a, b) => {
    let aValue = a[sortKey.value]
    let bValue = b[sortKey.value]
    
    if (sortKey.value === 'amount' || sortKey.value === 'remainingAmount') {
      aValue = parseFloat(String(aValue || 0)) || 0
      bValue = parseFloat(String(bValue || 0)) || 0
    } else if (sortKey.value === 'createdAt' || sortKey.value === 'expiresAt') {
      aValue = new Date(String(aValue || '')).getTime()
      bValue = new Date(String(bValue || '')).getTime()
    } else {
      aValue = String(aValue || '').toLowerCase()
      bValue = String(bValue || '').toLowerCase()
    }
    
    if (sortOrder.value === 'asc') {
      return aValue > bValue ? 1 : -1
    } else {
      return aValue < bValue ? 1 : -1
    }
  })
  
  return filtered
})

const totalPages = computed(() => Math.ceil(filteredCards.value.length / itemsPerPage.value))

const startIndex = computed(() => (currentPage.value - 1) * itemsPerPage.value)

const endIndex = computed(() => Math.min(startIndex.value + itemsPerPage.value, filteredCards.value.length))

const paginatedCards = computed(() => {
  return filteredCards.value.slice(startIndex.value, endIndex.value)
})

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
const sortBy = (key: string) => {
  if (sortKey.value === key) {
    sortOrder.value = sortOrder.value === 'asc' ? 'desc' : 'asc'
  } else {
    sortKey.value = key
    sortOrder.value = 'asc'
  }
  currentPage.value = 1
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

const viewCard = (card: GiftCard) => {
  emit('view-card', card)
}

const editCard = (card: GiftCard) => {
  emit('edit-card', card)
}

const duplicateCard = (card: GiftCard) => {
  emit('duplicate-card', card)
}

const deleteCard = (card: GiftCard) => {
  if (confirm(`آیا از حذف کارت ${String(card.code || '')} اطمینان دارید؟`)) {
    emit('delete-card', card)
  }
}

const refreshData = () => {
  // اینجا می‌توانید API call برای به‌روزرسانی داده‌ها اضافه کنید
}

const exportToExcel = () => {
  // اینجا می‌توانید منطق خروجی Excel اضافه کنید
}

const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleDateString('fa-IR')
}

const formatTime = (dateString: string) => {
  return new Date(dateString).toLocaleTimeString('fa-IR', {
    hour: '2-digit',
    minute: '2-digit'
  })
}

const getDaysUntilExpiry = (dateString: string) => {
  const expiryDate = new Date(dateString)
  const today = new Date()
  const diffTime = expiryDate.getTime() - today.getTime()
  const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24))
  
  if (diffDays < 0) {
    return `${Math.abs(diffDays)} روز گذشته`
  } else if (diffDays === 0) {
    return 'امروز'
  } else if (diffDays === 1) {
    return 'فردا'
  } else {
    return `${diffDays} روز دیگر`
  }
}

const getStatusClass = (status: string) => {
  switch (status) {
    case 'active':
      return 'bg-green-100 text-green-800'
    case 'used':
      return 'bg-blue-100 text-blue-800'
    case 'expired':
      return 'bg-red-100 text-red-800'
    case 'suspended':
      return 'bg-yellow-100 text-yellow-800'
    case 'cancelled':
      return 'bg-gray-100 text-gray-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getStatusLabel = (status: string) => {
  switch (status) {
    case 'active':
      return 'فعال'
    case 'used':
      return 'استفاده شده'
    case 'expired':
      return 'منقضی شده'
    case 'suspended':
      return 'معلق'
    case 'cancelled':
      return 'لغو شده'
    default:
      return 'نامشخص'
  }
}

// Lifecycle
onMounted(() => {
})
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 
