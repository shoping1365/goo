<template>
  <div class="min-h-screen bg-gray-50" dir="rtl">
    <!-- Header -->
    <div class="bg-white shadow-sm border-b border-gray-200 mb-6">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center py-4">
          <div>
            <h1 class="text-2xl font-bold text-gray-900">پروفایل رقبا</h1>
            <p class="text-sm text-gray-600 mt-1">تحلیل کامل رقبا و استراتژی‌های آن‌ها</p>
          </div>
          <div class="flex items-center space-x-3 space-x-reverse">
            <button 
              @click="refreshData"
              :disabled="isLoading"
              class="inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-lg text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 shadow-md transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed"
            >
              <svg v-if="isLoading" class="animate-spin -ml-1 mr-3 h-5 w-5 text-gray-700" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              <svg v-else class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
              </svg>
              {{ isLoading ? 'در حال به‌روزرسانی...' : 'به‌روزرسانی' }}
            </button>
            <button 
              @click="showAddCompetitorModal = true"
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 shadow-md transition-all duration-200"
            >
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
              </svg>
              افزودن رقیب جدید
            </button>
          </div>
        </div>
      </div>
    </div>

    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <!-- Filters and Search -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6 mb-6">
        <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
          <!-- Search -->
          <div class="md:col-span-2">
            <label class="block text-sm font-medium text-gray-700 mb-2">جستجو</label>
            <div class="relative">
              <input
                v-model="searchQuery"
                type="text"
                placeholder="جستجو بر اساس نام، محصول یا بازار..."
                class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
              <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                <svg class="h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
                </svg>
              </div>
            </div>
          </div>

          <!-- Threat Level Filter -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">سطح تهدید</label>
            <select
              v-model="selectedThreatLevel"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="">همه سطوح</option>
              <option value="low">کم</option>
              <option value="medium">متوسط</option>
              <option value="high">زیاد</option>
            </select>
          </div>

          <!-- Market Filter -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">بازار هدف</label>
            <select
              v-model="selectedMarket"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="">همه بازارها</option>
              <option value="iran">ایران</option>
              <option value="turkey">ترکیه</option>
              <option value="uae">امارات</option>
              <option value="iraq">عراق</option>
              <option value="afghanistan">افغانستان</option>
            </select>
          </div>
        </div>
      </div>

      <!-- Competitors Grid -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-8">
        <div 
          v-for="competitor in filteredCompetitors" 
          :key="competitor.id"
          class="bg-white rounded-xl shadow-sm border border-gray-200 hover:shadow-md transition-all duration-200 cursor-pointer"
          @click="selectCompetitor(competitor)"
        >
          <!-- Header -->
          <div class="p-6 border-b border-gray-100">
            <div class="flex items-start justify-between mb-4">
              <div class="flex-1">
                <h3 class="text-lg font-semibold text-gray-900 mb-2">{{ competitor.name }}</h3>
                <div class="flex items-center space-x-2 space-x-reverse">
                  <span class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium" :class="getThreatLevelBadgeClass(competitor.threatLevel)">
                    {{ getThreatLevelText(competitor.threatLevel) }}
                  </span>
                  <span class="text-sm text-gray-500">{{ competitor.companySize }}</span>
                </div>
              </div>
              <div class="flex space-x-2 space-x-reverse">
                <button 
                  @click.stop="editCompetitor(competitor)"
                  class="p-2 text-gray-400 hover:text-blue-600 hover:bg-blue-50 rounded-lg transition-colors"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
                  </svg>
                </button>
                <button 
                  @click.stop="deleteCompetitor(competitor.id)"
                  class="p-2 text-gray-400 hover:text-red-600 hover:bg-red-50 rounded-lg transition-colors"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                  </svg>
                </button>
              </div>
            </div>

            <!-- Basic Info -->
            <div class="space-y-2">
              <div class="flex items-center text-sm text-gray-600">
                <svg class="w-4 h-4 ml-2 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 12a9 9 0 01-9 9m9-9a9 9 0 00-9-9m9 9H3m9 9v-9m0-9v9"></path>
                </svg>
                <a :href="competitor.website" target="_blank" class="text-blue-600 hover:text-blue-700">{{ competitor.website }}</a>
              </div>
              <div class="flex items-center text-sm text-gray-600">
                <svg class="w-4 h-4 ml-2 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                </svg>
                تأسیس: {{ competitor.foundedYear }}
              </div>
            </div>
          </div>

          <!-- Content -->
          <div class="p-6">
            <!-- Target Markets -->
            <div class="mb-4">
              <h4 class="text-sm font-medium text-gray-700 mb-2">بازارهای هدف</h4>
              <div class="flex flex-wrap gap-1">
                <span 
                  v-for="market in competitor.targetMarkets.slice(0, 3)" 
                  :key="market"
                  class="inline-flex items-center px-2 py-1 rounded-md text-xs font-medium bg-blue-100 text-blue-800"
                >
                  {{ market }}
                </span>
                <span v-if="competitor.targetMarkets.length > 3" class="inline-flex items-center px-2 py-1 rounded-md text-xs font-medium bg-gray-100 text-gray-800">
                  +{{ competitor.targetMarkets.length - 3 }}
                </span>
              </div>
            </div>

            <!-- Key Products -->
            <div class="mb-4">
              <h4 class="text-sm font-medium text-gray-700 mb-2">محصولات کلیدی</h4>
              <div class="flex flex-wrap gap-1">
                <span 
                  v-for="product in competitor.keyProducts.slice(0, 2)" 
                  :key="product"
                  class="inline-flex items-center px-2 py-1 rounded-md text-xs font-medium bg-green-100 text-green-800"
                >
                  {{ product }}
                </span>
                <span v-if="competitor.keyProducts.length > 2" class="inline-flex items-center px-2 py-1 rounded-md text-xs font-medium bg-gray-100 text-gray-800">
                  +{{ competitor.keyProducts.length - 2 }}
                </span>
              </div>
            </div>

            <!-- Recent Changes -->
            <div>
              <h4 class="text-sm font-medium text-gray-700 mb-2">تغییرات اخیر</h4>
              <div class="space-y-1">
                <div 
                  v-for="change in competitor.recentChanges.slice(0, 2)" 
                  :key="change"
                  class="text-xs text-gray-600 flex items-start"
                >
                  <span class="w-1.5 h-1.5 bg-blue-500 rounded-full mt-1.5 ml-2 flex-shrink-0"></span>
                  {{ change }}
                </div>
                <div v-if="competitor.recentChanges.length > 2" class="text-xs text-gray-500 text-center">
                  و {{ competitor.recentChanges.length - 2 }} مورد دیگر
                </div>
              </div>
            </div>
          </div>

          <!-- Footer -->
          <div class="px-6 py-3 bg-gray-50 rounded-b-xl">
            <button 
              @click.stop="selectCompetitor(competitor)"
              class="w-full text-sm font-medium text-blue-600 hover:text-blue-700 text-center"
            >
              مشاهده جزئیات کامل
            </button>
          </div>
        </div>
      </div>

      <!-- Empty State -->
      <div v-if="filteredCompetitors.length === 0" class="text-center py-12">
        <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path>
        </svg>
        <h3 class="mt-2 text-sm font-medium text-gray-900">هیچ رقیبی یافت نشد</h3>
        <p class="mt-1 text-sm text-gray-500">با تغییر فیلترها یا افزودن رقیب جدید شروع کنید.</p>
        <div class="mt-6">
          <button 
            @click="showAddCompetitorModal = true"
            class="inline-flex items-center px-4 py-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
            </svg>
            افزودن رقیب جدید
          </button>
        </div>
      </div>
    </div>

    <!-- Competitor Detail Modal -->
    <CompetitorDetailModal 
      v-if="selectedCompetitor"
      :competitor="selectedCompetitor"
      @close="selectedCompetitor = null"
      @edit="editCompetitor"
    />

    <!-- Add/Edit Competitor Modal -->
    <CompetitorFormModal 
      v-if="showAddCompetitorModal"
      :competitor="editingCompetitor"
      @close="closeCompetitorModal"
      @save="saveCompetitor"
    />

    <!-- Delete Confirmation Modal -->
    <div v-if="showDeleteModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg p-6 w-full max-w-md mx-4">
        <div class="flex items-center mb-4">
          <div class="w-10 h-10 bg-red-100 rounded-full flex items-center justify-center ml-3">
            <svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900">حذف رقیب</h3>
        </div>
        <p class="text-sm text-gray-600 mb-6">
          آیا مطمئن هستید که می‌خواهید این رقیب را حذف کنید؟ این عملیات غیرقابل بازگشت است.
        </p>
        <div class="flex justify-end space-x-3 space-x-reverse">
          <button 
            @click="showDeleteModal = false"
            class="px-4 py-2 border border-gray-300 text-sm font-medium rounded-lg text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            انصراف
          </button>
          <button 
            @click="confirmDeleteCompetitor"
            class="px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-red-600 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
          >
            حذف
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useAuth } from '~/composables/useAuth'
import { useCompetitiveIntelligence, type CompetitorProfile } from '../composables/useCompetitiveIntelligence'

// تعریف definePageMeta برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user, hasPermission } = useAuth()

// Composables
const {
  competitors,
  isLoading,
  refreshData: refreshCI,
  deleteCompetitor: deleteCompetitorFromStore,
  addCompetitor,
  updateCompetitor
} = useCompetitiveIntelligence()

const refreshData = refreshCI

// State
const searchQuery = ref('')
const selectedThreatLevel = ref('')
const selectedMarket = ref('')
const selectedCompetitor = ref<CompetitorProfile | null>(null)
const showAddCompetitorModal = ref(false)
const editingCompetitor = ref<CompetitorProfile | null>(null)
const showDeleteModal = ref(false)
const competitorToDelete = ref<number | null>(null)

// Computed
const filteredCompetitors = computed(() => {
  let filtered = competitors.value

  // Search filter
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(comp => 
      comp.name.toLowerCase().includes(query) ||
      comp.keyProducts.some(product => product.toLowerCase().includes(query)) ||
      comp.targetMarkets.some(market => market.toLowerCase().includes(query))
    )
  }

  // Threat level filter
  if (selectedThreatLevel.value) {
    filtered = filtered.filter(comp => comp.threatLevel === selectedThreatLevel.value)
  }

  // Market filter
  if (selectedMarket.value) {
    filtered = filtered.filter(comp => 
      comp.targetMarkets.some(market => 
        market.toLowerCase().includes(selectedMarket.value.toLowerCase())
      )
    )
  }

  return filtered
})

// Methods
const selectCompetitor = (competitor: CompetitorProfile) => {
  selectedCompetitor.value = competitor
}

const editCompetitor = (competitor: CompetitorProfile) => {
  editingCompetitor.value = { ...competitor }
  showAddCompetitorModal.value = true
}

const deleteCompetitor = (id: number) => {
  competitorToDelete.value = id
  showDeleteModal.value = true
}

const confirmDeleteCompetitor = () => {
  if (competitorToDelete.value) {
    deleteCompetitorFromStore(competitorToDelete.value)
    showDeleteModal.value = false
    competitorToDelete.value = null
  }
}

const closeCompetitorModal = () => {
  showAddCompetitorModal.value = false
  editingCompetitor.value = null
}

const saveCompetitor = (competitorData: Omit<CompetitorProfile, 'id'>) => {
  if (editingCompetitor.value) {
    // Update existing competitor
    updateCompetitor(editingCompetitor.value.id, competitorData)
  } else {
    // Add new competitor
    addCompetitor(competitorData)
  }
  closeCompetitorModal()
}

const getThreatLevelBadgeClass = (level: string) => {
  switch (level) {
    case 'low': return 'bg-green-100 text-green-800'
    case 'medium': return 'bg-yellow-100 text-yellow-800'
    case 'high': return 'bg-red-100 text-red-800'
    default: return 'bg-gray-100 text-gray-800'
  }
}

const getThreatLevelText = (level: string) => {
  switch (level) {
    case 'low': return 'کم'
    case 'medium': return 'متوسط'
    case 'high': return 'زیاد'
    default: return 'نامشخص'
  }
}

// Lifecycle
onMounted(() => {
  // Load initial data if needed
})
</script>

<style scoped>
/* Custom styles */
</style>

