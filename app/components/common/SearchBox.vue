<template>
  <div class="relative w-full max-w-3xl" ref="searchContainer">
    <!-- باکس جستجو -->
    <div class="relative w-full">
      <input
        ref="searchInput"
        v-model="searchQuery"
        @input="handleSearchInput"
        @focus="handleFocus"
        @blur="handleBlur"
        @keydown="handleKeydown"
        type="text"
        :placeholder="placeholder"
        class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-200 bg-white"
        dir="rtl"
        autocomplete="off"
      />
      
      <!-- آیکون جستجو -->
      <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
        <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
        </svg>
      </div>
    </div>

    <!-- پیشنهادات جستجو -->
    <div
      v-if="showSuggestions"
      class="absolute top-full left-0 right-0 mt-1 bg-white border border-gray-200 rounded-lg shadow-lg z-[9999] max-h-[28rem] overflow-y-auto"
    >
      <div v-if="!hasTypedQuery" class="p-4 space-y-5">
        <div class="rounded-xl bg-gradient-to-l from-blue-500 via-indigo-500 to-purple-500 p-4 text-white">
          <div class="flex items-center justify-between gap-4">
            <div>
              <div class="text-sm font-semibold">جستجو کنید و زودتر پیدا کنید</div>
              <p class="mt-1 text-xs text-white/80">کلمات محبوب را انتخاب کنید یا خودتان بنویسید</p>
            </div>
            <svg class="w-12 h-12 text-white/70" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
            </svg>
          </div>
        </div>

        <div class="space-y-2">
          <div class="px-1 text-xs text-gray-400">پیشنهادهای آماده</div>
          <div class="flex flex-wrap gap-2">
            <button
              v-for="(preset, idx) in curatedQuickSearches"
              :key="`preset-${idx}`"
              type="button"
              class="rounded-full border border-gray-200 bg-white px-3 py-1 text-xs text-gray-600 hover:border-blue-400 hover:text-blue-600 transition-colors"
              @mousedown.prevent
              @click="triggerQuickSearch(preset)"
            >
              {{ preset }}
            </button>
          </div>
        </div>

        <div v-if="recentSearches.length > 0" class="space-y-2">
          <div class="px-1 text-xs text-gray-400">آخرین جستجوهای شما</div>
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
          <div class="grid gap-2 sm:grid-cols-2">
            <button
              v-for="(popular, index) in popularSearches"
              :key="`popular-${index}`"
              type="button"
              class="flex items-center justify-between rounded-lg border border-gray-200 px-3 py-2 text-left hover:border-blue-400 hover:bg-blue-50 transition-colors"
              @mousedown.prevent
              @click="triggerQuickSearch(popular.query)"
            >
              <div>
                <div class="text-sm font-medium text-gray-700">{{ popular.label }}</div>
                <div v-if="popular.context" class="text-xs text-gray-400 mt-0.5">{{ popular.context }}</div>
              </div>
              <svg class="w-4 h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 5l7 7-7 7"></path>
              </svg>
            </button>
          </div>
        </div>
      </div>

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
              @click="selectSuggestion(suggestion)"
              @mouseenter="highlightedIndex = index"
              :class="[
                'px-4 py-2 cursor-pointer hover:bg-gray-50 transition-colors duration-150 flex items-center',
                highlightedIndex === index ? 'bg-blue-50' : ''
              ]"
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
                  @click="selectResult(result)"
                  @mouseenter="highlightedIndex = suggestions.length + group.offset + index"
                  :class="[
                    'px-4 py-3 cursor-pointer hover:bg-gray-50 transition-colors duration-150 border-b border-gray-100 last:border-b-0',
                    highlightedIndex === suggestions.length + group.offset + index ? 'bg-blue-50' : ''
                  ]"
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
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" د="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6ه.01M6 20ه12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
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

                      <div v-if="hasPrice(result.price)" class="text-xs text-green-600 font-medium mt-1">
                        {{ formatPrice(result.price) }}
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <div class="px-4 py-2 border-t border-gray-100">
              <button
                @click="viewAllResults"
                class="w-full text-center text-sm text-blue-600 hover:text-blue-800 font-medium py-2"
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
              <svg class="w-5 ه-5 text-gray-400 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" د="M9.172 16.172a4 4 0 015.656 0M9 12ه6م-6-4ه6م2 5.291A7.962 7.962 0 014 12ه0c0 3.042 1.135 5.824 3 7.938ل3-2.647z"></path>
              </svg>
              نتیجه‌ای یافت نشد
            </div>
          </div>
        </div>
      </template>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, nextTick, onMounted, onUnmounted, ref, watch } from 'vue'
import { useRouter } from 'vue-router'

// Props
interface Props {
  placeholder?: string
  searchType?: 'all' | 'products' | 'posts' | 'categories' | 'brands'
  showResults?: boolean
  maxSuggestions?: number
}

const props = withDefaults(defineProps<Props>(), {
  placeholder: 'جستجو در محصولات، مقالات و...',
  searchType: 'all',
  showResults: true,
  maxSuggestions: 5
})

// Emits
const emit = defineEmits<{
  'search': [query: string]
  'suggestion-selected': [suggestion: string]
  'result-selected': [result: any]
}>()

// Reactive data
const searchQuery = ref('')
const querySnapshot = ref('')
const recentSearches = ref<string[]>([])
const suggestions = ref<string[]>([])
const searchResults = ref<any[]>([])
const isLoading = ref(false)
const showSuggestions = ref(false)
const showResultsPanel = ref(false)
const highlightedIndex = ref(-1)
const searchContainer = ref<HTMLElement>()
const searchInput = ref<HTMLInputElement>()
const isInputFocused = ref(false)

interface PopularSearchItem {
  label: string
  query: string
  context?: string
}

const curatedQuickSearches = ref<string[]>([
  'ساعت هوشمند',
  'ساعت عقربه ای مردانه',
  'کنسول بازی',
  'کیبورد مکانیکی',
  'گوشی سامسونگ',
  'تلویزیون 4K'
])

const popularSearches = ref<PopularSearchItem[]>([
  { label: 'ساعت هوشمند سامسونگ', query: 'ساعت هوشمند سامسونگ', context: 'در دسته ساعت هوشمند' },
  { label: 'ایرپاد بی سیم', query: 'ایرپاد بی سیم', context: 'در دسته لوازم جانبی صوتی' },
  { label: 'لپ تاپ دانشجویی', query: 'لپ تاپ دانشجویی', context: 'برای کارهای روزمره' },
  { label: 'دوربین عکاسی حرفه ای', query: 'دوربین عکاسی حرفه ای', context: 'برای عکاسان تازه کار' }
])

// Router
const router = useRouter()

// Computed
const totalResults = computed(() => searchResults.value.length)
const hasTypedQuery = computed(() => searchQuery.value.trim().length > 0)
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

// Methods
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

  // تاخیر برای جلوگیری از درخواست‌های مکرر
  clearTimeout((window as any).searchTimeout)
  ;(window as any).searchTimeout = setTimeout(async () => {
    await fetchSuggestions(currentQuery)
    if (props.showResults) {
      await fetchSearchResults(currentQuery)
    }
  }, 250)
}

const addRecentSearch = (query: string) => {
  const key = 'recent-searches'
  try {
    if (!query) return
    const existing = new Set(recentSearches.value)
    existing.delete(query)
    recentSearches.value = [query, ...recentSearches.value.filter(item => item !== query)].slice(0, 10)
    const stored = Array.from(new Set(recentSearches.value))
    if (stored.length > 10) {
      stored.length = 10
    }
    if (typeof localStorage !== 'undefined') {
      localStorage.setItem(key, JSON.stringify(stored))
    }
  } catch (err) {
    // ignore storage errors
  }
}

const loadRecentSearches = () => {
  try {
    if (typeof localStorage === 'undefined') return
    const raw = localStorage.getItem('recent-searches')
    if (!raw) return
    const parsed = JSON.parse(raw)
    if (Array.isArray(parsed)) {
      recentSearches.value = parsed.slice(0, 10)
    }
  } catch (err) {
    recentSearches.value = []
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
    
    if (trimmed !== searchQuery.value.trim()) {
      return
    }

    if (response.success) {
      suggestions.value = response.data.slice(0, props.maxSuggestions)
    } else {
      suggestions.value = []
    }
  } catch (error) {
    // خطا در دریافت پیشنهادات
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
    const params: any = { q: trimmed }
    if (props.searchType !== 'all') {
      params.type = props.searchType
    }
    params.limit = 10

    const response = await $fetch('/api/search', { query: params }) as any

    if (trimmed !== searchQuery.value.trim()) {
      return
    }
    
    if (response.success) {
      const mapped = Array.isArray(response.data)
        ? response.data.map((item: any) => ({
            ...item,
            title: resolveTitle(item),
            description: resolveDescription(item),
            url: resolveUrl(item),
            price: resolvePrice(item),
            image: resolveImage(item)
          }))
        : []

      const categories = mapped.filter((item: any) => item.type === 'category')
      const products = mapped.filter((item: any) => item.type === 'product')
      const others = mapped.filter((item: any) => item.type !== 'category' && item.type !== 'product')

      if (categories.length > 0) {
        searchResults.value = [...categories, ...products, ...others]
      } else if (products.length > 0) {
        searchResults.value = [...products, ...others]
      } else {
        searchResults.value = mapped
      }

      showResultsPanel.value = true
      querySnapshot.value = trimmed
    } else {
      searchResults.value = []
      showResultsPanel.value = false
      querySnapshot.value = trimmed
    }
  } catch (error) {
    // خطا در جستجو
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
  emit('suggestion-selected', suggestion)
  
  // انجام جستجو با پیشنهاد انتخاب شده
  if (props.showResults) {
    fetchSearchResults(suggestion)
  }

  nextTick(() => {
    searchQuery.value = ''
  })
}

const selectResult = (result: any) => {
  showResultsPanel.value = false
  showSuggestions.value = false
  isInputFocused.value = false
  emit('result-selected', result)
  addRecentSearch(result.title || querySnapshot.value)
  
  // هدایت به صفحه نتیجه
  router.push(result.url)

  nextTick(() => {
    searchQuery.value = ''
  })
}

const viewAllResults = () => {
  showResultsPanel.value = false
  showSuggestions.value = false
  isInputFocused.value = false
  emit('search', querySnapshot.value)
  addRecentSearch(querySnapshot.value)
  
  // هدایت به صفحه نتایج جستجو
  const params: any = { q: querySnapshot.value }
  if (props.searchType !== 'all') {
    params.type = props.searchType
  }
  router.push({ path: '/search-results', query: params })

  nextTick(() => {
    searchQuery.value = ''
  })
}



const handleBlur = () => {
  // تاخیر برای اجازه کلیک روی نتایج
  setTimeout(() => {
    isInputFocused.value = false
    showSuggestions.value = false
    showResultsPanel.value = false
    highlightedIndex.value = -1
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
  const labels: Record<string, string> = {
    product: 'محصول',
    post: 'مقاله',
    category: 'دسته‌بندی',
    brand: 'برند'
  }
  return labels[type] || type
}

const getTypeBadgeClass = (type: string) => {
  const classes: Record<string, string> = {
    product: 'bg-blue-100 text-blue-800',
    post: 'bg-green-100 text-green-800',
    category: 'bg-purple-100 text-purple-800',
    brand: 'bg-orange-100 text-orange-800'
  }
  return classes[type] || 'bg-gray-100 text-gray-800'
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

const resolvePrice = (item: any) => {
  const candidates = [
    item.price,
    item.final_price,
    item.sale_price,
    item.min_price,
    item.data?.price,
    item.data?.final_price
  ]

  for (const candidate of candidates) {
    if (candidate === null || candidate === undefined) continue
    const numeric = typeof candidate === 'string' ? parseFloat(candidate) : candidate
    if (typeof numeric === 'number' && isFinite(numeric) && numeric > 0) {
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

const resolveImage = (item: any) => {
  const normalize = (val?: unknown) => {
    if (typeof val !== 'string') return ''
    const trimmed = val.trim()
    return trimmed.length > 0 ? trimmed : ''
  }

  const buildProductThumbnail = (url: string) => {
    if (!url) return ''
    const normalized = normalize(url)
    if (!normalized) return ''

    if (normalized.includes('_thumbnail') || normalized.includes('/thumbnail')) {
      return normalized
    }

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
    const preferredExt = extension.toLowerCase() === '.jpg' || extension.toLowerCase() === '.jpeg' ? '.webp' : extension

    return `${baseName}_thumbnail${preferredExt}${query ? `?${query}` : ''}`
  }

  const takeFirst = (values: unknown[]) => {
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

  return item.type === 'product' ? buildProductThumbnail(selected) : selected
}

// Watchers
watch(searchQuery, (newValue) => {
  if (newValue.trim().length < 1) {
    suggestions.value = []
    searchResults.value = []
  }
})

// Click outside handler
const handleClickOutside = (event: Event) => {
  if (searchContainer.value && !searchContainer.value.contains(event.target as Node)) {
    isInputFocused.value = false
    showSuggestions.value = false
    showResultsPanel.value = false
    highlightedIndex.value = -1
  }
}

const handleFocus = () => {
  isInputFocused.value = true
  showSuggestions.value = true
}

const triggerQuickSearch = async (term: string) => {
  const normalized = term.trim()
  if (!normalized) {
    return
  }
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

// Lifecycle
onMounted(() => {
  document.addEventListener('click', handleClickOutside)
  loadRecentSearches()
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})
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