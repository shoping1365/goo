<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
    <div class="flex items-center space-x-4 space-x-reverse">
      <!-- جستجوی اصلی -->
      <div class="flex-1">
        <div class="relative">
          <input
            v-model="searchQuery"
            type="text"
            :placeholder="placeholder"
            class="w-full px-4 py-2 pr-12 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            @input="handleSearch"
            @keyup.enter="performSearch"
          />
          <div class="absolute inset-y-0 right-0 pr-3 flex items-center">
            <svg class="h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- فیلترهای سریع -->
      <div class="flex items-center space-x-2 space-x-reverse">
        <select
          v-model="searchType"
          class="px-3 py-2 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          @change="handleSearchTypeChange"
        >
          <option value="all">همه</option>
          <option value="code">کد کارت</option>
          <option value="buyer">خریدار</option>
          <option value="recipient">گیرنده</option>
          <option value="email">ایمیل</option>
          <option value="phone">شماره تماس</option>
        </select>

        <select
          v-model="statusFilter"
          class="px-3 py-2 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          @change="handleStatusFilterChange"
        >
          <option value="">همه وضعیت‌ها</option>
          <option value="active">فعال</option>
          <option value="used">استفاده شده</option>
          <option value="expired">منقضی شده</option>
          <option value="suspended">معلق</option>
        </select>

        <button 
          class="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
          @click="performSearch"
        >
          جستجو
        </button>

        <button 
          class="px-4 py-2 bg-gray-600 text-white text-sm font-medium rounded-lg hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
          @click="clearSearch"
        >
          پاک کردن
        </button>
      </div>
    </div>

    <!-- پیشنهادات جستجو -->
    <div v-if="showSuggestions && searchSuggestions.length > 0" class="mt-2">
      <div class="bg-gray-50 border border-gray-200 rounded-lg p-2">
        <div class="text-xs text-gray-500 mb-2">پیشنهادات:</div>
        <div class="space-y-1">
          <button
            v-for="suggestion in searchSuggestions"
            :key="suggestion.id"
            class="w-full text-right px-3 py-2 text-sm text-gray-700 hover:bg-gray-100 rounded-md transition-colors duration-200"
            @click="selectSuggestion(suggestion)"
          >
            <div class="flex items-center justify-between">
              <span>{{ suggestion.displayText }}</span>
              <span class="text-xs text-gray-500">{{ suggestion.type }}</span>
            </div>
          </button>
        </div>
      </div>
    </div>

    <!-- نتایج جستجوی سریع -->
    <div v-if="showQuickResults && quickResults.length > 0" class="mt-4">
      <div class="border-t border-gray-200 pt-4">
        <div class="flex items-center justify-between mb-3">
          <h4 class="text-sm font-medium text-gray-900">نتایج جستجو ({{ quickResults.length }})</h4>
          <button 
            class="text-sm text-blue-600 hover:text-blue-800"
            @click="showAdvancedSearch"
          >
            جستجوی پیشرفته
          </button>
        </div>
        
        <div class="space-y-2 max-h-64 overflow-y-auto">
          <div 
            v-for="result in quickResults" 
            :key="result.id"
            class="flex items-center justify-between p-3 border border-gray-200 rounded-lg hover:bg-gray-50 transition-colors duration-200"
          >
            <div class="flex items-center space-x-3 space-x-reverse">
              <div :class="getStatusColorClass(result.status)" class="w-2 h-2 rounded-full"></div>
              <div>
                <div class="text-sm font-medium text-gray-900">{{ result.code }}</div>
                <div class="text-xs text-gray-500">{{ result.recipientName }} - {{ result.recipientEmail }}</div>
              </div>
            </div>
            <div class="flex items-center space-x-2 space-x-reverse">
              <div class="text-right">
                <div class="text-sm font-medium text-gray-900">{{ formatCurrency(result.amount) }}</div>
                <div class="text-xs text-gray-500">{{ getStatusLabel(result.status) }}</div>
              </div>
              <div class="flex items-center space-x-1 space-x-reverse">
                <button 
                  class="text-blue-600 hover:text-blue-800 text-xs"
                  title="مشاهده جزئیات"
                  @click="viewCardDetails(result)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                  </svg>
                </button>
                <button 
                  class="text-green-600 hover:text-green-800 text-xs"
                  title="ویرایش"
                  @click="editCard(result)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- پیام خالی -->
    <div v-if="showQuickResults && quickResults.length === 0 && searchQuery" class="mt-4">
      <div class="text-center py-4">
        <svg class="mx-auto h-8 w-8 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
        </svg>
        <p class="mt-2 text-sm text-gray-500">نتیجه‌ای برای "{{ searchQuery }}" یافت نشد</p>
        <button 
          class="mt-2 text-sm text-blue-600 hover:text-blue-800"
          @click="showAdvancedSearch"
        >
          جستجوی پیشرفته را امتحان کنید
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, onUnmounted, ref, watch } from 'vue';

// Props
interface GiftCard {
  id?: number | string
  code?: string
  recipientName?: string
  recipientEmail?: string
  buyerName?: string
  recipientPhone?: string
  status?: string
  [key: string]: unknown
}

const props = defineProps<{
  giftCards: GiftCard[]
}>()

// Emits
const emit = defineEmits<{
  'view-details': [card: GiftCard]
  'edit-card': [card: GiftCard]
  'show-advanced-search': []
}>()

// Reactive data
const searchQuery = ref('')
const searchType = ref('all')
const statusFilter = ref('')
const showSuggestions = ref(false)
const showQuickResults = ref(false)

// Computed properties
const placeholder = computed(() => {
  const placeholders = {
    all: 'جستجو در همه فیلدها...',
    code: 'جستجو بر اساس کد کارت...',
    buyer: 'جستجو بر اساس نام خریدار...',
    recipient: 'جستجو بر اساس نام گیرنده...',
    email: 'جستجو بر اساس ایمیل...',
    phone: 'جستجو بر اساس شماره تماس...'
  }
  return placeholders[searchType.value] || placeholders.all
})

const searchSuggestions = computed(() => {
  if (!searchQuery.value || searchQuery.value.length < 2) return []
  
  const suggestions = []
  const query = searchQuery.value.toLowerCase()
  
  props.giftCards.forEach(card => {
    // پیشنهاد بر اساس کد کارت
    if (card.code.toLowerCase().includes(query)) {
      suggestions.push({
        id: `code-${card.id}`,
        displayText: card.code,
        type: 'کد کارت',
        card: card
      })
    }
    
    // پیشنهاد بر اساس نام گیرنده
    if (card.recipientName?.toLowerCase().includes(query)) {
      suggestions.push({
        id: `recipient-${card.id}`,
        displayText: card.recipientName,
        type: 'گیرنده',
        card: card
      })
    }
    
    // پیشنهاد بر اساس ایمیل
    if (card.recipientEmail?.toLowerCase().includes(query)) {
      suggestions.push({
        id: `email-${card.id}`,
        displayText: card.recipientEmail,
        type: 'ایمیل',
        card: card
      })
    }
    
    // پیشنهاد بر اساس نام خریدار
    if (card.buyerName?.toLowerCase().includes(query)) {
      suggestions.push({
        id: `buyer-${card.id}`,
        displayText: card.buyerName,
        type: 'خریدار',
        card: card
      })
    }
  })
  
  // حذف موارد تکراری
  const uniqueSuggestions = suggestions.filter((suggestion, index, self) => 
    index === self.findIndex(s => s.id === suggestion.id)
  )
  
  return uniqueSuggestions.slice(0, 5) // حداکثر 5 پیشنهاد
})

const quickResults = computed(() => {
  if (!searchQuery.value) return []
  
  let results = props.giftCards
  
  // فیلتر بر اساس نوع جستجو
  if (searchType.value !== 'all') {
    const query = searchQuery.value.toLowerCase()
    
    switch (searchType.value) {
      case 'code':
        results = results.filter(card => 
          card.code.toLowerCase().includes(query)
        )
        break
      case 'buyer':
        results = results.filter(card => 
          card.buyerName?.toLowerCase().includes(query)
        )
        break
      case 'recipient':
        results = results.filter(card => 
          card.recipientName?.toLowerCase().includes(query)
        )
        break
      case 'email':
        results = results.filter(card => 
          card.recipientEmail?.toLowerCase().includes(query)
        )
        break
      case 'phone':
        results = results.filter(card => 
          card.recipientPhone?.includes(searchQuery.value)
        )
        break
    }
  } else {
    // جستجو در همه فیلدها
    const query = searchQuery.value.toLowerCase()
    results = results.filter(card => 
      card.code.toLowerCase().includes(query) ||
      card.recipientName?.toLowerCase().includes(query) ||
      card.recipientEmail?.toLowerCase().includes(query) ||
      card.buyerName?.toLowerCase().includes(query) ||
      card.recipientPhone?.includes(searchQuery.value)
    )
  }
  
  // فیلتر بر اساس وضعیت
  if (statusFilter.value) {
    results = results.filter(card => card.status === statusFilter.value)
  }
  
  return results.slice(0, 10) // حداکثر 10 نتیجه
})

// Methods
const handleSearch = () => {
  showSuggestions.value = searchQuery.value.length >= 2
  showQuickResults.value = searchQuery.value.length >= 2
}

const performSearch = () => {
  showSuggestions.value = false
  showQuickResults.value = true

}

const handleSearchTypeChange = () => {
  if (searchQuery.value) {
    handleSearch()
  }
}

const handleStatusFilterChange = () => {
  if (searchQuery.value) {
    handleSearch()
  }
}

interface Suggestion {
  displayText?: string
  [key: string]: unknown
}

const selectSuggestion = (suggestion: Suggestion) => {
  searchQuery.value = suggestion.displayText || ''
  showSuggestions.value = false
  performSearch()
}

const clearSearch = () => {
  searchQuery.value = ''
  statusFilter.value = ''
  searchType.value = 'all'
  showSuggestions.value = false
  showQuickResults.value = false
}

const showAdvancedSearch = () => {
  emit('show-advanced-search')
}

const viewCardDetails = (card: GiftCard) => {
  emit('view-details', card)
}

const editCard = (card: GiftCard) => {
  emit('edit-card', card)
}

const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
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
watch(searchQuery, (newValue) => {
  if (newValue.length < 2) {
    showSuggestions.value = false
    showQuickResults.value = false
  }
})

// Event listeners
const handleClickOutside = (event: Event) => {
  const target = event.target as Element
  if (!target.closest('.search-container')) {
    showSuggestions.value = false
  }
}

// Lifecycle
onMounted(() => {
  document.addEventListener('click', handleClickOutside)

})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
.search-container {
  position: relative;
}
</style> 
