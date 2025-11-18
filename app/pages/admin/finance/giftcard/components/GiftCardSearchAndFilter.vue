<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex justify-between items-center">
        <div>
          <h2 class="text-xl font-bold text-gray-900">جستجو و فیلتر</h2>
          <p class="text-gray-600 mt-1">جستجو و فیلتر پیشرفته گیفت کارت‌ها</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button 
            @click="activeTab = 'quick'"
            :class="{
              'bg-blue-600 text-white': activeTab === 'quick',
              'bg-gray-200 text-gray-700 hover:bg-gray-300': activeTab !== 'quick'
            }"
            class="px-4 py-2 text-sm font-medium rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
          >
            جستجوی سریع
          </button>
          <button 
            @click="activeTab = 'advanced'"
            :class="{
              'bg-blue-600 text-white': activeTab === 'advanced',
              'bg-gray-200 text-gray-700 hover:bg-gray-300': activeTab !== 'advanced'
            }"
            class="px-4 py-2 text-sm font-medium rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
          >
            جستجوی پیشرفته
          </button>
          <button 
            @click="activeTab = 'saved'"
            :class="{
              'bg-blue-600 text-white': activeTab === 'saved',
              'bg-gray-200 text-gray-700 hover:bg-gray-300': activeTab !== 'saved'
            }"
            class="px-4 py-2 text-sm font-medium rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
          >
            فیلترهای ذخیره شده
          </button>
        </div>
      </div>
    </div>

    <!-- تب‌های محتوا -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200">
      <!-- تب جستجوی سریع -->
      <div v-if="activeTab === 'quick'">
        <GiftCardQuickSearch 
          :gift-cards="giftCards"
          @view-details="handleViewDetails"
          @edit-card="handleEditCard"
          @show-advanced-search="activeTab = 'advanced'"
        />
      </div>

      <!-- تب جستجوی پیشرفته -->
      <div v-if="activeTab === 'advanced'">
        <GiftCardAdvancedSearch 
          :gift-cards="giftCards"
          @view-details="handleViewDetails"
          @edit-card="handleEditCard"
        />
      </div>

      <!-- تب فیلترهای ذخیره شده -->
      <div v-if="activeTab === 'saved'">
        <GiftCardSavedFilters 
          :current-filters="currentFilters"
          @apply-filter="handleApplyFilter"
        />
      </div>
    </div>

    <!-- آمار جستجو -->
    <div v-if="searchStats.totalResults > 0" class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-medium text-gray-900 mb-4">آمار جستجو</h3>
      
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div class="text-center">
          <div class="text-2xl font-bold text-blue-600">{{ searchStats.totalResults }}</div>
          <div class="text-sm text-gray-600">کل نتایج</div>
        </div>
        <div class="text-center">
          <div class="text-2xl font-bold text-green-600">{{ searchStats.activeCards }}</div>
          <div class="text-sm text-gray-600">کارت‌های فعال</div>
        </div>
        <div class="text-center">
          <div class="text-2xl font-bold text-yellow-600">{{ searchStats.expiringCards }}</div>
          <div class="text-sm text-gray-600">در حال انقضا</div>
        </div>
        <div class="text-center">
          <div class="text-2xl font-bold text-red-600">{{ searchStats.expiredCards }}</div>
          <div class="text-sm text-gray-600">منقضی شده</div>
        </div>
      </div>

      <!-- نمودار توزیع وضعیت‌ها -->
      <div class="mt-6">
        <h4 class="text-sm font-medium text-gray-900 mb-3">توزیع وضعیت‌ها</h4>
        <div class="flex items-center space-x-2 space-x-reverse">
          <div 
            v-for="status in searchStats.statusDistribution" 
            :key="status.name"
            class="flex items-center space-x-2 space-x-reverse"
          >
            <div :class="getStatusColorClass(status.name)" class="w-3 h-3 rounded-full"></div>
            <span class="text-sm text-gray-600">{{ status.label }}: {{ status.count }}</span>
          </div>
        </div>
      </div>

      <!-- نمودار توزیع مبالغ -->
      <div class="mt-6">
        <h4 class="text-sm font-medium text-gray-900 mb-3">توزیع مبالغ</h4>
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div class="text-center p-3 bg-gray-50 rounded-lg">
            <div class="text-lg font-bold text-gray-900">{{ formatCurrency(searchStats.amountStats.min) }}</div>
            <div class="text-xs text-gray-600">کمترین مبلغ</div>
          </div>
          <div class="text-center p-3 bg-gray-50 rounded-lg">
            <div class="text-lg font-bold text-gray-900">{{ formatCurrency(searchStats.amountStats.avg) }}</div>
            <div class="text-xs text-gray-600">میانگین مبلغ</div>
          </div>
          <div class="text-center p-3 bg-gray-50 rounded-lg">
            <div class="text-lg font-bold text-gray-900">{{ formatCurrency(searchStats.amountStats.max) }}</div>
            <div class="text-xs text-gray-600">بیشترین مبلغ</div>
          </div>
        </div>
      </div>
    </div>

    <!-- عملیات دسته‌ای -->
    <div v-if="searchResults.length > 0" class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-4">
        <h3 class="text-lg font-medium text-gray-900">عملیات دسته‌ای</h3>
        <div class="flex items-center space-x-2 space-x-reverse">
          <span class="text-sm text-gray-500">{{ selectedCards.length }} کارت انتخاب شده</span>
          <button 
            @click="selectAllCards"
            class="text-sm text-blue-600 hover:text-blue-800"
          >
            انتخاب همه
          </button>
          <button 
            @click="clearSelection"
            class="text-sm text-gray-600 hover:text-gray-800"
          >
            پاک کردن انتخاب
          </button>
        </div>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <button 
          @click="bulkExport"
          :disabled="selectedCards.length === 0"
          class="px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-lg hover:bg-green-700 disabled:opacity-50 disabled:cursor-not-allowed focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2"
        >
          خروجی Excel
        </button>
        
        <button 
          @click="bulkSendEmail"
          :disabled="selectedCards.length === 0"
          class="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
        >
          ارسال ایمیل
        </button>
        
        <button 
          @click="bulkSendSMS"
          :disabled="selectedCards.length === 0"
          class="px-4 py-2 bg-purple-600 text-white text-sm font-medium rounded-lg hover:bg-purple-700 disabled:opacity-50 disabled:cursor-not-allowed focus:outline-none focus:ring-2 focus:ring-purple-500 focus:ring-offset-2"
        >
          ارسال پیامک
        </button>
        
        <button 
          @click="bulkUpdateStatus"
          :disabled="selectedCards.length === 0"
          class="px-4 py-2 bg-yellow-600 text-white text-sm font-medium rounded-lg hover:bg-yellow-700 disabled:opacity-50 disabled:cursor-not-allowed focus:outline-none focus:ring-2 focus:ring-yellow-500 focus:ring-offset-2"
        >
          تغییر وضعیت
        </button>
      </div>
    </div>

    <!-- نتایج جستجو -->
    <div v-if="searchResults.length > 0" class="bg-white rounded-lg shadow-sm border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <h3 class="text-lg font-medium text-gray-900">نتایج جستجو</h3>
          <div class="flex items-center space-x-4 space-x-reverse">
            <div class="flex items-center space-x-2 space-x-reverse">
              <span class="text-sm text-gray-500">مرتب‌سازی:</span>
              <select
                v-model="sortBy"
                @change="handleSort"
                class="px-3 py-1 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="createdAt">تاریخ ایجاد</option>
                <option value="amount">مبلغ</option>
                <option value="expiryDate">تاریخ انقضا</option>
                <option value="code">کد کارت</option>
                <option value="recipientName">نام گیرنده</option>
              </select>
              <select
                v-model="sortOrder"
                @change="handleSort"
                class="px-3 py-1 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="desc">نزولی</option>
                <option value="asc">صعودی</option>
              </select>
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <span class="text-sm text-gray-500">نمایش:</span>
              <select
                v-model="viewMode"
                class="px-3 py-1 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="list">لیست</option>
                <option value="grid">شبکه</option>
                <option value="table">جدول</option>
              </select>
            </div>
          </div>
        </div>
      </div>

      <!-- نمایش نتایج -->
      <div class="p-6">
        <!-- نمایش لیست -->
        <div v-if="viewMode === 'list'" class="space-y-4">
          <div 
            v-for="card in paginatedResults" 
            :key="card.id"
            class="flex items-center justify-between p-6 border border-gray-200 rounded-lg hover:bg-gray-50 transition-colors duration-200"
          >
            <div class="flex items-center space-x-4 space-x-reverse">
              <input
                v-model="selectedCards"
                :value="card.id"
                type="checkbox"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <div :class="getStatusColorClass(card.status)" class="w-3 h-3 rounded-full"></div>
              <div>
                <h4 class="text-sm font-medium text-gray-900">{{ card.code }}</h4>
                <p class="text-xs text-gray-500">{{ card.recipientName }} - {{ card.recipientEmail }}</p>
              </div>
            </div>
            <div class="flex items-center space-x-4 space-x-reverse">
              <div class="text-right">
                <p class="text-sm font-medium text-gray-900">{{ formatCurrency(card.amount) }}</p>
                <p class="text-xs text-gray-500">{{ getStatusLabel(card.status) }}</p>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <button 
                  @click="handleViewDetails(card)"
                  class="text-blue-600 hover:text-blue-800 text-sm"
                >
                  جزئیات
                </button>
                <button 
                  @click="handleEditCard(card)"
                  class="text-green-600 hover:text-green-800 text-sm"
                >
                  ویرایش
                </button>
              </div>
            </div>
          </div>
        </div>

        <!-- نمایش شبکه -->
        <div v-if="viewMode === 'grid'" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <div 
            v-for="card in paginatedResults" 
            :key="card.id"
            class="border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow duration-200"
          >
            <div class="flex items-center justify-between mb-3">
              <input
                v-model="selectedCards"
                :value="card.id"
                type="checkbox"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <div :class="getStatusColorClass(card.status)" class="w-3 h-3 rounded-full"></div>
            </div>
            <h4 class="text-sm font-medium text-gray-900 mb-2">{{ card.code }}</h4>
            <p class="text-xs text-gray-500 mb-3">{{ card.recipientName }}</p>
            <div class="space-y-2 mb-4">
              <div class="flex justify-between text-xs">
                <span class="text-gray-500">مبلغ:</span>
                <span class="font-medium">{{ formatCurrency(card.amount) }}</span>
              </div>
              <div class="flex justify-between text-xs">
                <span class="text-gray-500">وضعیت:</span>
                <span class="font-medium">{{ getStatusLabel(card.status) }}</span>
              </div>
              <div class="flex justify-between text-xs">
                <span class="text-gray-500">انقضا:</span>
                <span class="font-medium">{{ formatDate(card.expiryDate) }}</span>
              </div>
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <button 
                @click="handleViewDetails(card)"
                class="flex-1 px-3 py-1 bg-blue-600 text-white text-xs font-medium rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
              >
                جزئیات
              </button>
              <button 
                @click="handleEditCard(card)"
                class="flex-1 px-3 py-1 bg-green-600 text-white text-xs font-medium rounded-md hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2"
              >
                ویرایش
              </button>
            </div>
          </div>
        </div>

        <!-- نمایش جدول -->
        <div v-if="viewMode === 'table'" class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                  <input
                    v-model="selectAll"
                    @change="toggleSelectAll"
                    type="checkbox"
                    class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                  />
                </th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کد کارت</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">گیرنده</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مبلغ</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ انقضا</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="card in paginatedResults" :key="card.id" class="hover:bg-gray-50">
                <td class="px-6 py-4 whitespace-nowrap">
                  <input
                    v-model="selectedCards"
                    :value="card.id"
                    type="checkbox"
                    class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                  />
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{{ card.code }}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ card.recipientName }}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ formatCurrency(card.amount) }}</td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span :class="getStatusBadgeClass(card.status)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                    {{ getStatusLabel(card.status) }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ formatDate(card.expiryDate) }}</td>
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <button 
                      @click="handleViewDetails(card)"
                      class="text-blue-600 hover:text-blue-900"
                    >
                      جزئیات
                    </button>
                    <button 
                      @click="handleEditCard(card)"
                      class="text-green-600 hover:text-green-900"
                    >
                      ویرایش
                    </button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <!-- صفحه‌بندی -->
        <div v-if="totalPages > 1" class="mt-6 flex items-center justify-between">
          <div class="flex items-center space-x-2 space-x-reverse">
            <button 
              @click="previousPage"
              :disabled="currentPage === 1"
              class="px-3 py-1 border border-gray-300 rounded-md text-sm disabled:opacity-50 disabled:cursor-not-allowed"
            >
              قبلی
            </button>
            <span class="text-sm text-gray-700">
              صفحه {{ currentPage }} از {{ totalPages }}
            </span>
            <button 
              @click="nextPage"
              :disabled="currentPage === totalPages"
              class="px-3 py-1 border border-gray-300 rounded-md text-sm disabled:opacity-50 disabled:cursor-not-allowed"
            >
              بعدی
            </button>
          </div>
          <div class="flex items-center space-x-2 space-x-reverse">
            <span class="text-sm text-gray-500">نمایش:</span>
            <select 
              v-model="pageSize"
              @change="currentPage = 1"
              class="px-2 py-1 border border-gray-300 rounded-md text-sm"
            >
              <option value="10">10</option>
              <option value="25">25</option>
              <option value="50">50</option>
              <option value="100">100</option>
            </select>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted, watch } from 'vue'

// کامپوننت‌های مورد نیاز
import GiftCardQuickSearch from './GiftCardQuickSearch.vue'
import GiftCardAdvancedSearch from './GiftCardAdvancedSearch.vue'
import GiftCardSavedFilters from './GiftCardSavedFilters.vue'

// Props
const props = defineProps<{
  giftCards: any[]
}>()

// Emits
const emit = defineEmits<{
  'view-details': [card: any]
  'edit-card': [card: any]
}>()

// Reactive data
const activeTab = ref('quick')
const currentPage = ref(1)
const pageSize = ref(25)
const sortBy = ref('createdAt')
const sortOrder = ref('desc')
const viewMode = ref('list')
const selectedCards = ref([])
const searchResults = ref([])
const currentFilters = ref({})

// آمار جستجو
const searchStats = reactive({
  totalResults: 0,
  activeCards: 0,
  expiringCards: 0,
  expiredCards: 0,
  statusDistribution: [],
  amountStats: {
    min: 0,
    avg: 0,
    max: 0
  }
})

// Computed properties
const selectAll = computed({
  get: () => selectedCards.value.length === searchResults.value.length && searchResults.value.length > 0,
  set: (value) => {
    if (value) {
      selectedCards.value = searchResults.value.map(card => card.id)
    } else {
      selectedCards.value = []
    }
  }
})

const totalPages = computed(() => Math.ceil(searchResults.value.length / pageSize.value))

const paginatedResults = computed(() => {
  const start = (currentPage.value - 1) * pageSize.value
  const end = start + pageSize.value
  return searchResults.value.slice(start, end)
})

// Methods
const handleViewDetails = (card: any) => {
  emit('view-details', card)
}

const handleEditCard = (card: any) => {
  emit('edit-card', card)
}

const handleApplyFilter = (filters: any) => {
  currentFilters.value = filters
  applyFilters(filters)
}

const applyFilters = (filters: any) => {
  // اینجا منطق اعمال فیلترها اضافه می‌شود
  searchResults.value = [...props.giftCards] // فعلاً همه کارت‌ها
  updateSearchStats()
  currentPage.value = 1
}

const handleSort = () => {
  searchResults.value.sort((a, b) => {
    let aValue = a[sortBy.value]
    let bValue = b[sortBy.value]
    
    if (sortBy.value === 'createdAt' || sortBy.value === 'expiryDate') {
      aValue = new Date(aValue).getTime()
      bValue = new Date(bValue).getTime()
    }
    
    if (sortOrder.value === 'asc') {
      return aValue > bValue ? 1 : -1
    } else {
      return aValue < bValue ? 1 : -1
    }
  })
}

const toggleSelectAll = () => {
  if (selectAll.value) {
    selectedCards.value = []
  } else {
    selectedCards.value = searchResults.value.map(card => card.id)
  }
}

const selectAllCards = () => {
  selectedCards.value = searchResults.value.map(card => card.id)
}

const clearSelection = () => {
  selectedCards.value = []
}

const bulkExport = () => {
  console.log('خروجی Excel برای کارت‌های انتخاب شده:', selectedCards.value)
}

const bulkSendEmail = () => {
  console.log('ارسال ایمیل برای کارت‌های انتخاب شده:', selectedCards.value)
}

const bulkSendSMS = () => {
  console.log('ارسال پیامک برای کارت‌های انتخاب شده:', selectedCards.value)
}

const bulkUpdateStatus = () => {
  console.log('تغییر وضعیت کارت‌های انتخاب شده:', selectedCards.value)
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

const updateSearchStats = () => {
  const results = searchResults.value
  
  searchStats.totalResults = results.length
  searchStats.activeCards = results.filter(card => card.status === 'active').length
  searchStats.expiringCards = results.filter(card => {
    const expiryDate = new Date(card.expiryDate)
    const now = new Date()
    const daysUntilExpiry = (expiryDate.getTime() - now.getTime()) / (1000 * 60 * 60 * 24)
    return daysUntilExpiry <= 7 && daysUntilExpiry > 0
  }).length
  searchStats.expiredCards = results.filter(card => {
    const expiryDate = new Date(card.expiryDate)
    const now = new Date()
    return expiryDate < now
  }).length

  // توزیع وضعیت‌ها
  const statusCounts = {}
  results.forEach(card => {
    statusCounts[card.status] = (statusCounts[card.status] || 0) + 1
  })
  
  searchStats.statusDistribution = Object.entries(statusCounts).map(([status, count]) => ({
    name: status,
    label: getStatusLabel(status),
    count
  }))

  // آمار مبالغ
  const amounts = results.map(card => card.amount)
  searchStats.amountStats.min = Math.min(...amounts)
  searchStats.amountStats.max = Math.max(...amounts)
  searchStats.amountStats.avg = Math.round(amounts.reduce((sum, amount) => sum + amount, 0) / amounts.length)
}

const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

const formatDate = (date: Date) => {
  return new Intl.DateTimeFormat('fa-IR').format(date)
}

const getStatusColorClass = (status: string) => {
  const classes = {
    active: 'bg-green-500',
    used: 'bg-blue-500',
    expired: 'bg-red-500',
    suspended: 'bg-yellow-500',
    cancelled: 'bg-gray-500'
  }
  return classes[status] || 'bg-gray-500'
}

const getStatusBadgeClass = (status: string) => {
  const classes = {
    active: 'bg-green-100 text-green-800',
    used: 'bg-blue-100 text-blue-800',
    expired: 'bg-red-100 text-red-800',
    suspended: 'bg-yellow-100 text-yellow-800',
    cancelled: 'bg-gray-100 text-gray-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getStatusLabel = (status: string) => {
  const labels = {
    active: 'فعال',
    used: 'استفاده شده',
    expired: 'منقضی شده',
    suspended: 'معلق',
    cancelled: 'لغو شده'
  }
  return labels[status] || status
}

// Watchers
watch(searchResults, () => {
  updateSearchStats()
}, { deep: true })

// Lifecycle
onMounted(() => {
  searchResults.value = [...props.giftCards]
  updateSearchStats()
  console.log('Gift card search and filter component mounted')
})
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 
