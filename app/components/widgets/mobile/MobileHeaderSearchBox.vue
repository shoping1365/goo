<template>
  <div ref="searchContainer" class="mobile-search">
    <div class="mobile-search__input">
      <input
        ref="searchInput"
        v-model="searchQuery"
        type="text"
        :placeholder="placeholder"
        autocomplete="off"
        dir="rtl"
        @input="handleSearchInput"
        @focus="handleFocus"
        @blur="handleBlur"
        @keydown="handleKeydown"
      />
      <button type="button" class="mobile-search__icon" @mousedown.prevent @click="submitQuery">
        <svg class="icon" viewBox="0 0 24 24" fill="none" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
        </svg>
      </button>
    </div>

    <transition name="fade">
      <div
        v-if="showPanel"
        class="mobile-search__panel"
      >
        <div v-if="!hasTypedQuery" class="mobile-search__default">
          <div v-if="recentSearches.length" class="mobile-search__section">
            <div class="mobile-search__section-header">
              <span>آخرین جستجوها</span>
              <button type="button" @click="clearRecentSearches">پاک کردن</button>
            </div>
            <div class="mobile-search__chips">
              <button v-for="recent in recentSearches" :key="`recent-${recent}`" type="button" @mousedown.prevent @click="triggerQuickSearch(recent)">
                {{ recent }}
              </button>
            </div>
          </div>

          <div v-if="popularSearches.length" class="mobile-search__section">
            <div class="mobile-search__section-header">
              <span>جستجوهای پرطرفدار</span>
            </div>
            <div class="mobile-search__chips">
              <button v-for="(popular, index) in popularSearches" :key="`popular-${index}`" type="button" @mousedown.prevent @click="triggerQuickSearch(popular.query)">
                {{ popular.label }}
              </button>
            </div>
          </div>
        </div>

        <div v-else class="mobile-search__results">
          <div v-if="isLoading" class="mobile-search__loading">
            <svg class="spinner" viewBox="0 0 50 50">
              <circle class="path" cx="25" cy="25" r="20" fill="none" stroke-width="5" />
            </svg>
            <span>در حال جستجو...</span>
          </div>

          <template v-else>
            <div v-if="suggestions.length" class="mobile-search__section">
              <div class="mobile-search__section-header">
                <span>پیشنهادها</span>
              </div>
              <ul class="mobile-search__list">
                <li
                  v-for="(suggestion, index) in suggestions"
                  :key="`suggestion-${index}`"
                  :class="{ active: highlightedIndex === index }"
                  @mousedown.prevent
                  @mouseenter="highlightedIndex = index"
                  @click="selectSuggestion(suggestion)"
                >
                  <svg class="icon" viewBox="0 0 24 24" fill="none" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                  </svg>
                  <span>{{ suggestion }}</span>
                </li>
              </ul>
            </div>

            <div v-if="searchResults.length" class="mobile-search__section">
              <div class="mobile-search__section-header">
                <span>نتایج مرتبط</span>
              </div>
              <ul class="mobile-search__results-list">
                <li
                  v-for="(result, index) in searchResults"
                  :key="`${result.type}-${result.id}`"
                  :class="{ active: highlightedIndex === suggestions.length + index }"
                  @mousedown.prevent
                  @mouseenter="highlightedIndex = suggestions.length + index"
                  @click="selectResult(result)"
                >
                  <div class="thumb">
                    <img v-if="result.image" :src="result.image" :alt="result.title" loading="lazy" />
                    <div v-else class="placeholder">
                      <svg class="icon" viewBox="0 0 24 24" fill="none" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                      </svg>
                    </div>
                  </div>
                  <div class="meta">
                    <h4>{{ result.title }}</h4>
                    <p v-if="result.description">{{ result.description }}</p>
                    <div class="meta__footer">
                      <span class="badge">{{ getTypeLabel(result.type) }}</span>
                      <span v-if="result.price" class="price">{{ formatPrice(result.price) }}</span>
                    </div>
                  </div>
                </li>
              </ul>
              <button type="button" class="mobile-search__all" @click="viewAllResults">
                مشاهده همه نتایج ({{ totalResults }})
              </button>
            </div>

            <div v-if="!suggestions.length && !searchResults.length" class="mobile-search__empty">
              نتیجه‌ای یافت نشد
            </div>
          </template>
        </div>
      </div>
    </transition>
  </div>
</template>

<script setup lang="ts">
import { computed, nextTick, onMounted, onUnmounted, ref, watch } from 'vue'
import { useRouter } from 'vue-router'

interface PopularSearchItem {
  label: string
  query: string
  context?: string
}

interface SearchResult {
  id?: string | number
  type?: 'product' | 'category' | 'post' | 'brand' | string
  title?: string
  name?: string
  description?: string
  url?: string
  image?: string
  price?: number | string | null
  final_price?: number | string | null
  sale_price?: number | string | null
  slug?: string
  link?: string
  data?: {
    title?: string
    name?: string
    description?: string
    price?: number | string
    final_price?: number | string
    image?: string
    image_url?: string
    main_image?: string
    thumbnail?: string
    thumbnail_url?: string
    thumb?: string
  }
  media?: Array<{
    thumbnail?: string
    thumbnail_url?: string
    small?: string
    url?: string
  }>
  thumbnail?: string
  thumbnail_url?: string
  image_thumbnail?: string
  small_image?: string
  thumb?: string
  image_url?: string
  main_image?: string
  featured_image?: string
  cover?: string
  picture?: string
  excerpt?: string
}

interface SearchResponse {
  success?: boolean
  data?: SearchResult[] | unknown[]
}

interface SuggestionsResponse {
  success?: boolean
  data?: string[]
}

interface WindowWithTimeout extends Window {
  mobileSearchTimeout?: ReturnType<typeof setTimeout>
}

withDefaults(defineProps<{ placeholder?: string }>(), {
  placeholder: 'جستجو...'
})

const router = useRouter()

const searchQuery = ref('')
const querySnapshot = ref('')
const recentSearches = ref<string[]>([])
const suggestions = ref<string[]>([])
const searchResults = ref<SearchResult[]>([])
const popularSearches = ref<PopularSearchItem[]>([])
const isLoading = ref(false)
const showPanel = ref(false)
const highlightedIndex = ref(-1)
const isInputFocused = ref(false)
const searchContainer = ref<HTMLElement | null>(null)
const searchInput = ref<HTMLInputElement | null>(null)

const hasTypedQuery = computed(() => searchQuery.value.trim().length > 0)
const totalResults = computed(() => searchResults.value.length)

const handleSearchInput = () => {
  const current = searchQuery.value.trim()

  if (!current) {
    suggestions.value = []
    searchResults.value = []
    querySnapshot.value = ''
    highlightedIndex.value = -1
    isLoading.value = false
    showPanel.value = isInputFocused.value
    return
  }

  querySnapshot.value = current
  showPanel.value = true
  highlightedIndex.value = -1

  const windowWithTimeout = window as WindowWithTimeout
  clearTimeout(windowWithTimeout.mobileSearchTimeout)
  windowWithTimeout.mobileSearchTimeout = setTimeout(async () => {
    await Promise.all([
      fetchSuggestions(current),
      fetchSearchResults(current)
    ])
  }, 220)
}

const fetchSuggestions = async (keyword: string) => {
  const trimmed = keyword.trim()
  if (!trimmed) return

  try {
    isLoading.value = true
    const response = await $fetch<SuggestionsResponse>('/api/search/suggestions', {
      query: { q: trimmed }
    })

    if (trimmed !== searchQuery.value.trim()) return

    if (response?.success && Array.isArray(response.data)) {
      suggestions.value = response.data.slice(0, 5)
    } else {
      suggestions.value = []
    }
  } catch (_error) {
    suggestions.value = []
  } finally {
    isLoading.value = false
  }
}

const fetchSearchResults = async (keyword: string) => {
  const trimmed = keyword.trim()
  if (!trimmed) return

  try {
    const response = await $fetch<SearchResponse>('/api/search', {
      query: { q: trimmed, limit: 8 }
    })

    if (trimmed !== searchQuery.value.trim()) return

    if (response?.success && Array.isArray(response.data)) {
      searchResults.value = response.data.map((item: SearchResult) => ({
        ...item,
        title: item.title || item.name || item.data?.title || item.data?.name || 'بدون عنوان',
        description: item.description || item.excerpt || item.data?.description || '',
        price: item.price || item.final_price || item.sale_price || item.data?.price || null,
        image: resolveImage(item),
        url: item.url || item.link || (item.slug ? `/${item.type}/${item.slug}` : '#')
      }))
    } else {
      searchResults.value = []
    }
  } catch (_error) {
    searchResults.value = []
  }
}

const resolveImage = (item: SearchResult) => {
  const take = (values: unknown[]) => {
    for (const value of values) {
      if (typeof value === 'string' && value.trim()) {
        return value.trim()
      }
    }
    return ''
  }

  const main = take([
    item.thumbnail,
    item.thumbnail_url,
    item.image_thumbnail,
    item.image,
    item.image_url,
    item.main_image,
    item.featured_image,
    item.cover,
    item.data?.thumbnail,
    item.data?.image,
    item.data?.image_url
  ])

  return main || '/statics/images/default-image_100.png'
}

const handleFocus = () => {
  isInputFocused.value = true
  showPanel.value = true
}

const handleBlur = () => {
  setTimeout(() => {
    isInputFocused.value = false
    showPanel.value = false
    highlightedIndex.value = -1
  }, 200)
}

const handleKeydown = (event: KeyboardEvent) => {
  const total = suggestions.value.length + searchResults.value.length
  switch (event.key) {
    case 'ArrowDown':
      event.preventDefault()
      highlightedIndex.value = Math.min(highlightedIndex.value + 1, total - 1)
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
          const selected = searchResults.value[highlightedIndex.value - suggestions.value.length]
          if (selected) {
            selectResult(selected)
          }
        }
      } else {
        submitQuery()
      }
      break
    case 'Escape':
      showPanel.value = false
      highlightedIndex.value = -1
      searchInput.value?.blur()
      break
  }
}

const selectSuggestion = (suggestion: string) => {
  searchQuery.value = suggestion
  addRecentSearch(suggestion)
  submitQuery()
}

const selectResult = (result: SearchResult) => {
  addRecentSearch(result.title || querySnapshot.value)
  showPanel.value = false
  searchInput.value?.blur()
  if (result.url) {
    router.push(result.url)
  }
  nextTick(() => {
    searchQuery.value = ''
  })
}

const submitQuery = () => {
  const query = searchQuery.value.trim() || querySnapshot.value.trim()
  if (!query) return

  addRecentSearch(query)
  showPanel.value = false
  router.push({ path: '/search-results', query: { q: query } })
  nextTick(() => {
    searchQuery.value = ''
  })
}

const viewAllResults = () => {
  submitQuery()
}

const triggerQuickSearch = async (term: string) => {
  const normalized = term.trim()
  if (!normalized) return
  searchQuery.value = normalized
  await nextTick()
  handleSearchInput()
}

const addRecentSearch = (query: string) => {
  const trimmed = query.trim()
  if (!trimmed) return

  const unique = new Set(recentSearches.value)
  unique.delete(trimmed)
  recentSearches.value = [trimmed, ...Array.from(unique)].slice(0, 8)

  if (typeof localStorage !== 'undefined') {
    localStorage.setItem('mobile-recent-searches', JSON.stringify(recentSearches.value))
  }
}

const loadRecentSearches = () => {
  try {
    if (typeof localStorage === 'undefined') return
    const raw = localStorage.getItem('mobile-recent-searches')
    if (!raw) return
    const parsed = JSON.parse(raw)
    if (Array.isArray(parsed)) {
      recentSearches.value = parsed.filter((item: unknown) => typeof item === 'string' && item.trim()).slice(0, 8)
    }
  } catch (_error) {
    recentSearches.value = []
  }
}

const clearRecentSearches = () => {
  recentSearches.value = []
  if (typeof localStorage !== 'undefined') {
    localStorage.removeItem('mobile-recent-searches')
  }
}

const loadPopularSearches = async () => {
  try {
    interface PopularSearchResponse {
      success?: boolean
      data?: Array<{
        query?: string
        term?: string
        keyword?: string
        context?: string
        category?: string
      }>
    }
    const response = await $fetch<PopularSearchResponse>('/api/search/popular')
    if (response?.success && Array.isArray(response.data)) {
      popularSearches.value = response.data.slice(0, 8).map((item) => ({
        label: item.query || item.term || item.keyword,
        query: item.query || item.term || item.keyword,
        context: item.context || item.category || ''
      }))
    }
  } catch (_error) {
    popularSearches.value = []
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

const formatPrice = (price: number | string | null | undefined): string => {
  if (price === null || price === undefined) return '0 تومان'
  const numPrice = typeof price === 'string' ? parseFloat(price) || 0 : price
  return new Intl.NumberFormat('fa-IR').format(numPrice) + ' تومان'
}

const handleClickOutside = (event: Event) => {
  if (searchContainer.value && !searchContainer.value.contains(event.target as Node)) {
    showPanel.value = false
    highlightedIndex.value = -1
  }
}

watch(searchQuery, (value) => {
  if (!value.trim()) {
    suggestions.value = []
    searchResults.value = []
  }
})

onMounted(() => {
  loadRecentSearches()
  loadPopularSearches()
  document.addEventListener('click', handleClickOutside)
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>

<style scoped>
.mobile-search {
  position: relative;
  width: 100%;
}

.mobile-search__input {
  display: flex;
  align-items: center;
  background: #f8f9fa;
  border-radius: 20px;
  border: 1px solid #e5e7eb;
  padding: 6px 12px;
  gap: 8px;
}

.mobile-search__input input {
  flex: 1;
  border: none;
  background: transparent;
  outline: none;
  font-size: 14px;
  color: #374151;
}

.mobile-search__input input::placeholder {
  color: #9ca3af;
}

.mobile-search__icon {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  background: #2563eb;
  color: #ffffff;
  border-radius: 999px;
  width: 32px;
  height: 32px;
  border: none;
}

.mobile-search__icon .icon {
  width: 18px;
  height: 18px;
}

.mobile-search__panel {
  position: absolute;
  top: calc(100% + 8px);
  right: 0;
  left: 0;
  background: #ffffff;
  border-radius: 16px;
  border: 1px solid #e5e7eb;
  box-shadow: 0 20px 45px rgba(15, 23, 42, 0.12);
  padding: 12px;
  z-index: 9999;
  max-height: 26rem;
  overflow-y: auto;
}

.mobile-search__default,
.mobile-search__results {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.mobile-search__section {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.mobile-search__section-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-size: 12px;
  color: #6b7280;
}

.mobile-search__section-header button {
  background: transparent;
  border: none;
  font-size: 11px;
  color: #ef4444;
}

.mobile-search__chips {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.mobile-search__chips button {
  border: 1px solid #e5e7eb;
  border-radius: 999px;
  padding: 6px 12px;
  background: #ffffff;
  font-size: 12px;
  color: #4b5563;
}

.mobile-search__list,
.mobile-search__results-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.mobile-search__list li,
.mobile-search__results-list li {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 10px;
  border-radius: 12px;
  font-size: 13px;
  color: #374151;
}

.mobile-search__list li.active,
.mobile-search__results-list li.active,
.mobile-search__list li:hover,
.mobile-search__results-list li:hover {
  background: #eff6ff;
}

.mobile-search__list .icon {
  width: 16px;
  height: 16px;
  color: #9ca3af;
}

.mobile-search__results-list .thumb {
  width: 48px;
  height: 48px;
  border-radius: 12px;
  overflow: hidden;
  background: #f3f4f6;
  flex-shrink: 0;
}

.mobile-search__results-list .thumb img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.mobile-search__results-list .placeholder {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 100%;
  color: #9ca3af;
}

.mobile-search__results-list .meta {
  flex: 1;
  min-width: 0;
}

.mobile-search__results-list .meta h4 {
  font-size: 14px;
  color: #111827;
  margin-bottom: 4px;
  line-height: 1.4;
}

.mobile-search__results-list .meta p {
  font-size: 12px;
  color: #6b7280;
  line-height: 1.6;
  line-clamp: 2;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.mobile-search__results-list .meta__footer {
  margin-top: 6px;
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 11px;
}

.mobile-search__results-list .badge {
  background: #eff6ff;
  color: #2563eb;
  padding: 4px 8px;
  border-radius: 999px;
}

.mobile-search__results-list .price {
  color: #16a34a;
  font-weight: 600;
}

.mobile-search__all {
  width: 100%;
  border: none;
  border-radius: 12px;
  padding: 10px;
  text-align: center;
  font-size: 13px;
  color: #2563eb;
  background: #eff6ff;
}

.mobile-search__empty {
  text-align: center;
  font-size: 13px;
  color: #9ca3af;
  padding: 16px 0;
}

.mobile-search__loading {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  font-size: 13px;
  color: #6b7280;
}

.spinner {
  width: 20px;
  height: 20px;
  animation: spin 1s linear infinite;
}

.spinner .path {
  stroke: #2563eb;
  stroke-linecap: round;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}
</style>
