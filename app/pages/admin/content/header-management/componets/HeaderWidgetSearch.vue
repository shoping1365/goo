<template>
  <!-- باکس جستجو -->
  <div class="header-search-wrapper w-full h-full flex items-center" :style="getItemStyle()">
    <div ref="searchContainer" class="w-[70%] mx-auto relative">
      <!-- Input جستجو -->
      <div class="relative w-full">
        <input
          ref="searchInput"
          v-model="searchQuery"
          type="text"
          placeholder="جستجو در محصولات، مقالات و..."
          class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-200 focus:border-blue-400 transition-all duration-200 bg-white"
          dir="rtl"
          autocomplete="off"
          @input="handleSearchInput"
          @focus="handleFocus"
          @blur="handleBlur"
          @keydown="handleKeydown"
        />
        
        <!-- آیکون جستجو -->
        <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
          <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
          </svg>
        </div>
      </div>

      <!-- پنل پیشنهادات و نتایج -->
      <Teleport to="body">
        <div
          v-show="showSuggestions && searchContainer"
          :style="{
            position: 'fixed',
            top: panelPosition.top + 'px',
            left: panelPosition.left + 'px',
            width: panelPosition.width + 'px',
            zIndex: 99999
          }"
          class="bg-white border border-gray-200 rounded-lg shadow-lg max-h-[28rem] overflow-y-auto"
          dir="rtl"
        >
        <!-- حالت خالی: پنل دیجی‌کالا -->
        <div v-if="!hasTypedQuery" class="p-4 space-y-5">
          <div v-if="recentSearches.length > 0" class="space-y-2">
            <div class="px-1 flex items-center justify-between">
              <span class="text-xs text-gray-400">آخرین جستجوهای شما</span>
              <button
                type="button"
                class="text-red-500 hover:text-red-700 transition-colors p-1"
                title="پاک کردن همه"
                @click="clearRecentSearches"
              >
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                </svg>
              </button>
            </div>
            <div class="flex flex-wrap gap-2">
              <button
                v-for="(recent, index) in recentSearches"
                :key="`recent-${index}`"
                type="button"
                class="rounded-full bg-gray-100 px-3 py-1 text-xs text-gray-600 hover:bg-gray-200 transition-colors"
                @mousedown.prevent
                @click="triggerQuickSearch(recent)"
              >
                {{ recent }}
              </button>
            </div>
          </div>

          <div class="space-y-2">
            <div class="px-1 text-xs text-gray-400">جستجوهای پرطرفدار</div>
            <div class="flex flex-wrap gap-2">
              <button
                v-for="(popular, index) in popularSearches.slice(0, 10)"
                :key="`popular-${index}`"
                type="button"
                class="rounded-full bg-gray-100 px-3 py-1 text-xs text-gray-600 hover:bg-gray-200 transition-colors"
                @mousedown.prevent
                @click="triggerQuickSearch(popular.query)"
              >
                {{ popular.label }}
              </button>
            </div>
          </div>
        </div>

        <!-- حالت تایپ: پیشنهادات و نتایج -->
        <template v-else>
          <div v-if="isLoading" class="p-6 text-center text-gray-500">
            <div class="flex items-center justify-center">
              <svg class="animate-spin h-5 w-5 text-blue-500 ml-2" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              در حال جستجو...
            </div>
          </div>
          <div v-else class="py-2">
            <div v-if="suggestions.length > 0" class="pt-2">
              <div class="px-4 text-xs text-gray-400 mb-1">پیشنهادها</div>
              <div
                v-for="(suggestion, index) in suggestions"
                :key="`suggestion-${index}`"
                :class="[
                  'px-4 py-2 cursor-pointer hover:bg-gray-50 transition-colors duration-150 flex items-center',
                  highlightedIndex === index ? 'bg-blue-50' : ''
                ]"
                @click="selectSuggestion(suggestion)"
                @mouseenter="highlightedIndex = index"
              >
                <svg class="w-4 h-4 text-gray-400 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
                </svg>
                <span class="text-sm text-gray-700">{{ suggestion }}</span>
              </div>
            </div>

            <div
              v-if="showResultsPanel && searchResults.length > 0"
              class="pt-2 mt-2 border-t border-gray-100"
            >
              <div class="px-4 text-xs text-gray-400 mb-2">نتایج مرتبط</div>
              <div class="space-y-4">
                <div
                  v-for="group in groupedResults"
                  :key="group.key"
                  class="space-y-1"
                >
                  <div class="px-4 text-[11px] text-gray-400">{{ group.label }}</div>
                  <div
                    v-for="(result, index) in group.items"
                    :key="`${group.key}-${result.type}-${result.id}`"
                    :class="[
                      'px-4 py-3 cursor-pointer hover:bg-gray-50 transition-colors duration-150 border-b border-gray-100 last:border-b-0',
                      highlightedIndex === suggestions.length + group.offset + index ? 'bg-blue-50' : ''
                    ]"
                    @click="selectResult(result)"
                    @mouseenter="highlightedIndex = suggestions.length + group.offset + index"
                  >
                    <div class="flex items-start space-x-3 space-x-reverse">
                      <div class="flex-shrink-0">
                        <img
                          v-if="result.image"
                          :src="result.image"
                          :alt="result.title"
                          class="w-[50px] h-[50px] object-cover rounded-lg"
                          loading="lazy"
                        />
                        <div
                          v-else
                          class="w-12 h-12 bg-gray-200 rounded-lg flex items-center justify-center"
                        >
                          <svg class="w-6 h-6 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6ه.01M6 20ه12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                          </svg>
                        </div>
                      </div>

                      <div class="flex-1 min-w-0">
                        <div class="flex items-center justify-between">
                          <h4 class="text-sm font-medium text-gray-900 truncate">{{ result.title }}</h4>
                          <span
                            :class="[
                              'text-xs px-2 py-1 rounded-full',
                              getTypeBadgeClass(result.type)
                            ]"
                          >
                            {{ getTypeLabel(result.type) }}
                          </span>
                        </div>

                        <p v-if="result.description" class="text-xs text-gray-500 mt-1 line-clamp-2">
                          {{ result.description }}
                        </p>

                        <div v-if="result.type === 'product' && hasPrice(result.price)" class="text-xs text-green-600 font-medium mt-1">
                          {{ formatPrice(result.price) }}
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              <div class="px-4 py-2 border-t border-gray-100">
                <button
                  class="w-full text-center text-sm text-blue-600 hover:text-blue-800 font-medium py-2"
                  @click="viewAllResults"
                >
                  مشاهده همه نتایج ({{ totalResults }})
                </button>
              </div>
            </div>

            <div
              v-if="suggestions.length === 0 && (!showResultsPanel || searchResults.length === 0)"
              class="p-6 text-center text-gray-500"
            >
              <div class="flex items-center justify-center">
                <svg class="w-5 h-5 text-gray-400 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.172 16.172a4 4 0 015.656 0M9 12h6m-6-4h6m2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
                نتیجه‌ای یافت نشد
              </div>
            </div>
          </div>
        </template>
        </div>
      </Teleport>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, nextTick, onMounted, onUnmounted, ref, watch } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter()

// دریافت props از parent component
const props = defineProps<{
  width?: number
  align?: 'left' | 'center' | 'right'
  bgColor?: string
  paddingRight?: number
  paddingLeft?: number
}>()

// State های جستجو
const searchQuery = ref('')
const querySnapshot = ref('')
const recentSearches = ref<string[]>([])
const suggestions = ref<string[]>([])
const searchResults = ref<any[]>([])
const groupedResults = computed(() => {
  const groups: Array<{ key: string; label: string; items: any[]; offset: number }> = []
  let offset = 0

  const ordered = [
    { key: 'categories', type: 'category', label: 'دسته بندی ها' },
    { key: 'products', type: 'product', label: 'محصولات' }
  ]

  for (const entry of ordered) {
    const items = searchResults.value.filter(item => item.type === entry.type)
    if (items.length) {
      groups.push({ key: entry.key, label: entry.label, items, offset })
      offset += items.length
    }
  }

  const otherItems = searchResults.value.filter(item => !ordered.some(entry => entry.type === item.type))
  if (otherItems.length) {
    groups.push({ key: 'others', label: 'سایر نتایج', items: otherItems, offset })
  }

  return groups
})
const isLoading = ref(false)
const showSuggestions = ref(false)
const showResultsPanel = ref(false)
const highlightedIndex = ref(-1)
const searchContainer = ref<HTMLElement>()
const searchInput = ref<HTMLInputElement>()
const isInputFocused = ref(false)
const panelPosition = ref({ top: 0, left: 0, width: 0 })

interface PopularSearchItem {
  label: string
  query: string
  context?: string
}

const popularSearches = ref<PopularSearchItem[]>([])

const loadPopularSearches = async () => {
  try {
    const response = await $fetch('/api/search/popular') as any
    if (response.success && Array.isArray(response.data)) {
      popularSearches.value = response.data.slice(0, 10).map((item: any) => ({
        label: item.query || item.term || item.keyword,
        query: item.query || item.term || item.keyword,
        context: item.context || item.category || ''
      }))
    }
  } catch (error) {
    // اگر API موجود نیست، از داده‌های پیش‌فرض استفاده کن
    popularSearches.value = [
      { label: 'گوشی سامسونگ', query: 'گوشی سامسونگ', context: 'موبایل' },
      { label: 'لپ تاپ ایسوس', query: 'لپ تاپ ایسوس', context: 'کامپیوتر' },
      { label: 'هدفون بلوتوثی', query: 'هدفون بلوتوثی', context: 'صوتی' },
      { label: 'ساعت هوشمند', query: 'ساعت هوشمند', context: 'پوشیدنی' },
      { label: 'کیبورد گیمینگ', query: 'کیبورد گیمینگ', context: 'جانبی' },
      { label: 'مانیتور', query: 'مانیتور', context: 'نمایشگر' },
      { label: 'موس بی سیم', query: 'موس بی سیم', context: 'جانبی' },
      { label: 'شارژر فست', query: 'شارژر فست', context: 'لوازم جانبی' },
      { label: 'پاوربانک', query: 'پاوربانک', context: 'لوازم جانبی' },
      { label: 'کابل تایپ سی', query: 'کابل تایپ سی', context: 'لوازم جانبی' }
    ]
  }
}

const totalResults = computed(() => searchResults.value.length)
const hasTypedQuery = computed(() => searchQuery.value.trim().length > 0)

const handleSearchInput = async () => {
  const currentQuery = searchQuery.value.trim()
  if (currentQuery.length < 1) {
    suggestions.value = []
    showResultsPanel.value = false
    querySnapshot.value = ''
    highlightedIndex.value = -1
    isLoading.value = false
    showSuggestions.value = isInputFocused.value
    return
  }

  querySnapshot.value = currentQuery
  showSuggestions.value = true
  highlightedIndex.value = -1

  clearTimeout((window as any).searchTimeout)
  ;(window as any).searchTimeout = setTimeout(async () => {
    await fetchSuggestions(currentQuery)
    await fetchSearchResults(currentQuery)
  }, 250)
}

const addRecentSearch = (query: string) => {
  try {
    if (!query || !query.trim()) return
    const trimmedQuery = query.trim()
    
    // حذف تکراری‌ها و اضافه کردن در ابتدای لیست
    const filtered = recentSearches.value.filter(item => item !== trimmedQuery)
    recentSearches.value = [trimmedQuery, ...filtered].slice(0, 8)
    
    if (typeof localStorage !== 'undefined') {
      localStorage.setItem('recent-searches', JSON.stringify(recentSearches.value))
    }
  } catch (err) {
    console.error('Error saving recent search:', err)
  }
}

const loadRecentSearches = () => {
  try {
    if (typeof localStorage === 'undefined') return
    const raw = localStorage.getItem('recent-searches')
    if (!raw) return
    const parsed = JSON.parse(raw)
    if (Array.isArray(parsed)) {
      // فقط جستجوهای منحصر به فرد و حداکثر 8 تا
      recentSearches.value = [...new Set(parsed)].slice(0, 8).filter(item => item && item.trim())
    }
  } catch (err) {
    recentSearches.value = []
  }
}

const clearRecentSearches = () => {
  recentSearches.value = []
  if (typeof localStorage !== 'undefined') {
    localStorage.removeItem('recent-searches')
  }
}

const fetchSuggestions = async (query: string) => {
  const trimmed = query.trim()
  if (!trimmed) {
    suggestions.value = []
    showSuggestions.value = isInputFocused.value
    return
  }

  try {
    isLoading.value = true
    showSuggestions.value = true
    const response = await $fetch('/api/search/suggestions', {
      query: { q: trimmed }
    }) as any
    
    if (trimmed !== searchQuery.value.trim()) return

    if (response.success) {
      suggestions.value = response.data.slice(0, 5)
    } else {
      suggestions.value = []
    }
  } catch (error) {
    suggestions.value = []
  } finally {
    isLoading.value = false
    showSuggestions.value = isInputFocused.value
  }
}

const fetchSearchResults = async (query: string) => {
  const trimmed = query.trim()
  if (!trimmed) {
    searchResults.value = []
    showResultsPanel.value = false
    return
  }

  try {
    const params: any = { q: trimmed, limit: 10 }
    const response = await $fetch('/api/search', { query: params }) as any

    if (trimmed !== searchQuery.value.trim()) return
    
    if (response.success) {
      const mapped = Array.isArray(response.data)
        ? response.data.map((item: any) => {
            const normalizedType = normalizeResultType(item.type || item.data?.type)
            const enrichedItem = { ...item, type: normalizedType }
            return {
              ...enrichedItem,
              title: resolveTitle(enrichedItem),
              description: resolveDescription(enrichedItem),
              url: resolveUrl(enrichedItem),
              price: resolvePrice(enrichedItem),
              image: resolveImage(enrichedItem)
            }
          })
        : []

      const categories = mapped.filter((item: any) => item.type === 'category')
      const products = mapped.filter((item: any) => item.type === 'product')
      const others = mapped.filter((item: any) => item.type !== 'category' && item.type !== 'product')

      // ترتیب: دسته‌بندی‌ها اول، بعد محصولات، بعد بقیه
      searchResults.value = [...categories, ...products, ...others]

      showResultsPanel.value = searchResults.value.length > 0
      querySnapshot.value = trimmed
    } else {
      searchResults.value = []
      showResultsPanel.value = false
      querySnapshot.value = trimmed
    }
  } catch (error) {
    searchResults.value = []
    showResultsPanel.value = false
    querySnapshot.value = trimmed
  }
}

const selectSuggestion = (suggestion: string) => {
  searchQuery.value = suggestion
  showSuggestions.value = false
  querySnapshot.value = suggestion
  addRecentSearch(suggestion)
  fetchSearchResults(suggestion)
  nextTick(() => {
    searchQuery.value = ''
  })
}

const selectResult = (result: any) => {
  showResultsPanel.value = false
  showSuggestions.value = false
  isInputFocused.value = false
  addRecentSearch(result.title || querySnapshot.value)
  router.push(result.url)
  nextTick(() => {
    searchQuery.value = ''
  })
}

const viewAllResults = () => {
  showResultsPanel.value = false
  showSuggestions.value = false
  isInputFocused.value = false
  addRecentSearch(querySnapshot.value)
  router.push({ path: '/search-results', query: { q: querySnapshot.value } })
  nextTick(() => {
    searchQuery.value = ''
  })
}

const handleBlur = () => {
  // تاخیر برای اینکه کلیک روی آیتم‌ها اول اجرا بشه
  setTimeout(() => {
    // فقط اگر واقعاً focus از دست رفته باشه
    if (document.activeElement !== searchInput.value) {
      isInputFocused.value = false
      showSuggestions.value = false
      showResultsPanel.value = false
      highlightedIndex.value = -1
    }
  }, 200)
}

const handleKeydown = (event: KeyboardEvent) => {
  const totalItems = suggestions.value.length + searchResults.value.length
  
  switch (event.key) {
    case 'ArrowDown':
      event.preventDefault()
      highlightedIndex.value = Math.min(highlightedIndex.value + 1, totalItems - 1)
      break
    case 'ArrowUp':
      event.preventDefault()
      highlightedIndex.value = Math.max(highlightedIndex.value - 1, -1)
      break
    case 'Enter':
      event.preventDefault()
      if (highlightedIndex.value >= 0) {
        if (highlightedIndex.value < suggestions.value.length) {
          selectSuggestion(suggestions.value[highlightedIndex.value])
        } else {
          const resultIndex = highlightedIndex.value - suggestions.value.length
          selectResult(searchResults.value[resultIndex])
        }
      } else if (searchQuery.value.trim()) {
        viewAllResults()
      }
      break
    case 'Escape':
      showSuggestions.value = false
      showResultsPanel.value = false
      highlightedIndex.value = -1
      searchInput.value?.blur()
      break
  }
}

const getTypeLabel = (type: string) => {
  const normalized = normalizeResultType(type)
  const labels: Record<string, string> = {
    product: 'محصول',
    post: 'مقاله',
    category: 'دسته‌بندی',
    brand: 'برند'
  }
  return labels[normalized] || normalized
}

const getTypeBadgeClass = (type: string) => {
  const normalized = normalizeResultType(type)
  const classes: Record<string, string> = {
    product: 'bg-blue-100 text-blue-800',
    post: 'bg-green-100 text-green-800',
    category: 'bg-purple-100 text-purple-800',
    brand: 'bg-orange-100 text-orange-800'
  }
  return classes[normalized] || 'bg-gray-100 text-gray-800'
}

const normalizeResultType = (rawType?: string) => {
  if (!rawType) return 'other'
  const lowered = String(rawType).toLowerCase()
  if (lowered === 'categories') return 'category'
  if (lowered === 'products') return 'product'
  if (lowered === 'posts') return 'post'
  return lowered
}

const resolveTitle = (item: any) => {
  return item.title || item.name || item.data?.title || item.data?.name || 'بدون عنوان'
}

const resolveDescription = (item: any) => {
  return item.description || item.excerpt || item.data?.description || ''
}

const resolveUrl = (item: any) => {
  if (item.url || item.link) {
    return item.url || item.link
  }
  if (item.slug) {
    const base = item.type === 'category' ? '/category' : `/${item.type || 'item'}`
    return `${base}/${item.slug}`
  }
  return '#'
}

const parseJsonIfNeeded = (value: string) => {
  const trimmed = value.trim()
  if (!trimmed) return null
  const startsLikeJson = ['[', '{', '"'].includes(trimmed[0])
  const endsLikeJson = [']', '}', '"'].includes(trimmed[trimmed.length - 1])
  if (!startsLikeJson || !endsLikeJson) {
    return null
  }
  try {
    return JSON.parse(trimmed)
  } catch (error) {
    return null
  }
}

const PRICE_KEYS = [
  'price',
  'selling_price',
  'sale_price',
  'final_price',
  'current_price',
  'discounted_price',
  'price_with_discount',
  'price_without_discount',
  'display_price',
  'min_price',
  'max_price'
]

const normalizePriceValue = (value: unknown): number | null => {
  if (value === null || value === undefined) {
    return null
  }

  if (typeof value === 'number') {
    return Number.isFinite(value) ? value : null
  }

  if (Array.isArray(value)) {
    for (const entry of value) {
      const numeric = normalizePriceValue(entry)
      if (numeric !== null) {
        return numeric
      }
    }
    return null
  }

  if (typeof value === 'object') {
    for (const entry of Object.values(value as Record<string, unknown>)) {
      const numeric = normalizePriceValue(entry)
      if (numeric !== null) {
        return numeric
      }
    }
    return null
  }

  if (typeof value === 'string') {
    const trimmed = value.trim()
    if (!trimmed) return null

    const parsedStructure = parseJsonIfNeeded(trimmed)
    if (parsedStructure !== null) {
      return normalizePriceValue(parsedStructure)
    }

    const digits = trimmed.replace(/[^0-9\.\-٬٬۰-۹]/g, '').replace(/[٬,]/g, '')
    if (!digits) return null
    const formatted = digits.replace(/[۰-۹]/g, ch => String('۰۱۲۳۴۵۶۷۸۹'.indexOf(ch)))
    const numeric = parseFloat(formatted)
    return Number.isFinite(numeric) ? numeric : null
  }

  return null
}

const collectPriceCandidates = (item: any): unknown[] => {
  const candidates: unknown[] = []
  const push = (val: unknown) => {
    if (val !== null && val !== undefined) {
      candidates.push(val)
    }
  }

  const sources = [item, item.data, item.meta]
  for (const source of sources) {
    if (!source || typeof source !== 'object') continue
    for (const key of PRICE_KEYS) {
      if (key in source) {
        push((source as Record<string, unknown>)[key])
      }
    }
    for (const [key, value] of Object.entries(source as Record<string, unknown>)) {
      if (key.toLowerCase().includes('price')) {
        push(value)
      }
    }
  }

  push(item.amount)
  push(item.total)

  return candidates
}

const resolvePrice = (item: any) => {
  const candidates = collectPriceCandidates(item)
  for (const candidate of candidates) {
    const numeric = normalizePriceValue(candidate)
    if (numeric !== null && numeric > 0) {
      return Math.round(numeric)
    }
  }
  return null
}

const hasPrice = (price: number | string | null | undefined) => {
  if (price === null || price === undefined) {
    return false
  }
  const numeric = typeof price === 'string' ? parseFloat(price) : price
  return typeof numeric === 'number' && isFinite(numeric) && numeric > 0
}

const formatPrice = (price: number | string | null | undefined) => {
  if (!hasPrice(price)) {
    return ''
  }
  const numeric = typeof price === 'string' ? parseFloat(price) : price!
  return new Intl.NumberFormat('fa-IR').format(Math.round(numeric)) + ' تومان'
}

const stripWrappingQuotes = (value: string) => {
  return value.replace(/^['"]+|['"]+$/g, '')
}

const resolveImage = (item: any) => {
  const normalize = (val?: unknown): string => {
    if (val === null || val === undefined) {
      return ''
    }

    if (typeof val === 'string') {
      const parsed = parseJsonIfNeeded(val)
      if (parsed !== null) {
        return normalize(parsed)
      }
      const stripped = stripWrappingQuotes(val.trim())
      return stripped.length > 0 ? stripped : ''
    }

    if (Array.isArray(val)) {
      return takeFirst(val)
    }

    if (typeof val === 'object') {
      return takeFirst(Object.values(val as Record<string, unknown>))
    }

    return ''
  }

  const buildProductThumbnail = (url: string) => {
    if (!url) return ''
    const normalized = normalize(url)
    if (!normalized) return ''

    // اگر قبلاً thumbnail هست، همون رو برگردون
    if (normalized.includes('_thumbnail') || normalized.includes('/thumbnail')) {
      return normalized
    }

    // اگر URL خارجی هست (شروع با http)، همون رو برگردون
    if (normalized.startsWith('http://') || normalized.startsWith('https://')) {
      return normalized
    }

    const [basePath, query = ''] = normalized.split('?', 2)
    const dotIndex = basePath.lastIndexOf('.')
    
    if (dotIndex === -1) {
      return `${basePath}_thumbnail.webp${query ? `?${query}` : ''}`
    }

    const baseName = basePath.slice(0, dotIndex)
    const extension = basePath.slice(dotIndex)
    
    // تبدیل به webp برای بهینه‌سازی
    const thumbnailUrl = `${baseName}_thumbnail.webp${query ? `?${query}` : ''}`
    
    return thumbnailUrl
  }

  const takeFirst = (values: unknown[]): string => {
    for (const value of values) {
      if (Array.isArray(value)) {
        for (const inner of value) {
          const normalized = normalize(inner)
          if (normalized) {
            return normalized
          }
        }
        continue
      }

      if (value && typeof value === 'object') {
        const mediaObj = value as Record<string, unknown>
        const nested = takeFirst([
          mediaObj.thumbnail,
          mediaObj.thumbnail_url,
          mediaObj.small,
          mediaObj.small_url,
          mediaObj.url,
          mediaObj.image
        ])
        if (nested) {
          return nested
        }
      }

      const normalized = normalize(value)
      if (normalized) {
        return normalized
      }
    }
    return ''
  }

  // اولویت با thumbnail
  const thumbnailCandidates = [
    item.thumbnail,
    item.thumbnail_url,
    item.image_thumbnail,
    item.small_image,
    item.thumb,
    item.data?.thumbnail,
    item.data?.thumbnail_url,
    item.data?.thumb,
    item.media?.[0]?.thumbnail,
    item.media?.[0]?.thumbnail_url,
    item.media?.[0]?.small,
    item.media?.[0]?.url
  ]

  let selected = takeFirst(thumbnailCandidates)

  if (!selected) {
    // اگر thumbnail نبود، از تصویر اصلی استفاده کن
    const fallbackCandidates = [
      item.image,
      item.image_url,
      item.main_image,
      item.featured_image,
      item.cover,
      item.picture,
      item.media,
      item.data?.image,
      item.data?.image_url,
      item.data?.main_image
    ]

    const fallback = takeFirst(fallbackCandidates)
    if (fallback) {
      selected = item.type === 'product' ? buildProductThumbnail(fallback) : fallback
    }
  }

  if (!selected) {
    return '/statics/images/default-image_100.png'
  }

  // برای محصولات، thumbnail بساز
  return item.type === 'product' ? buildProductThumbnail(selected) : selected
}

watch(searchQuery, (newValue) => {
  if (newValue.trim().length < 1) {
    suggestions.value = []
    searchResults.value = []
  }
})

const handleClickOutside = (event: Event) => {
  const target = event.target as Node
  
  // اگر کلیک داخل searchContainer بود، هیچ کاری نکن
  if (searchContainer.value && searchContainer.value.contains(target)) {
    return
  }
  
  isInputFocused.value = false
  showSuggestions.value = false
  showResultsPanel.value = false
  highlightedIndex.value = -1
}

const handleFocus = () => {
  isInputFocused.value = true
  showSuggestions.value = true
  updatePanelPosition()
  
  // اگر چیزی تایپ نشده، فقط پنل خالی رو نشون بده
  if (!searchQuery.value.trim()) {
    suggestions.value = []
    searchResults.value = []
    showResultsPanel.value = false
  }
}

const updatePanelPosition = () => {
  if (!searchContainer.value) return
  const rect = searchContainer.value.getBoundingClientRect()
  panelPosition.value = {
    top: rect.bottom + window.scrollY + 4,
    left: rect.left + window.scrollX,
    width: rect.width
  }
}

const triggerQuickSearch = async (term: string) => {
  const normalized = term.trim()
  if (!normalized) return
  
  if (!isInputFocused.value) {
    searchInput.value?.focus()
    isInputFocused.value = true
  }
  showSuggestions.value = true
  searchQuery.value = normalized
  highlightedIndex.value = -1
  addRecentSearch(normalized)
  await nextTick()
  handleSearchInput()
}

onMounted(() => {
  document.addEventListener('click', handleClickOutside)
  loadRecentSearches()
  loadPopularSearches()
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})

// تابع تعیین استایل آیتم
const getItemStyle = () => {
  const style: Record<string, any> = {}
  
  if (props.width) {
    style.width = `${props.width}%`
    style.flex = `0 0 ${props.width}%`
  }
  
  if (props.bgColor && props.bgColor !== 'transparent') {
    style.backgroundColor = props.bgColor
  }
  
  if (props.paddingRight !== undefined) {
    style.paddingRight = `${props.paddingRight}px`
  }
  if (props.paddingLeft !== undefined) {
    style.paddingLeft = `${props.paddingLeft}px`
  }
  
  style.display = 'flex'
  style.alignItems = 'center'
  
  switch (props.align) {
    case 'left':
      style.justifyContent = 'flex-start'
      break
    case 'center':
      style.justifyContent = 'center'
      break
    case 'right':
      style.justifyContent = 'flex-end'
      break
    default:
      style.justifyContent = 'center'
  }
  
  return style
}
</script>

<style scoped>
.line-clamp-2 {
  display: -webkit-box;
  line-clamp: 2;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>
