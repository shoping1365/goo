<template>
  <div class="min-h-screen bg-gray-50" dir="rtl">
    <main class="max-w-5xl mx-auto px-4 sm:px-6 lg:px-8 py-10">
      <section class="mb-8">
        <h1 class="text-2xl font-bold text-gray-900">نتایج جستجو</h1>
        <p v-if="searchQuery" class="text-gray-600 mt-2">
          نمایش {{ totalResults }} نتیجه برای "{{ searchQuery }}"
          <span v-if="searchTime" class="text-sm text-gray-500">
            (در {{ searchTime }} میلی‌ثانیه)
          </span>
        </p>
        <p v-else class="text-gray-600 mt-2">
          برای مشاهده نتایج، عبارت مورد نظر خود را جستجو کنید.
        </p>
      </section>

      <section>
        <div v-if="isLoading" class="flex items-center justify-center py-16">
          <div class="text-center">
            <svg class="animate-spin h-8 w-8 text-blue-500 mx-auto mb-4" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            <p class="text-gray-600">در حال جستجو...</p>
          </div>
        </div>

        <div v-else-if="searchQuery && paginatedResults.length === 0" class="text-center py-16">
          <svg class="w-16 h-16 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.172 16.172a4 4 0 015.656 0M9 12h6m-6-4h6m2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          <h3 class="text-lg font-medium text-gray-900 mb-2">نتیجه‌ای یافت نشد</h3>
          <p class="text-gray-600 mb-4">
            برای "{{ searchQuery }}" نتیجه‌ای یافت نشد. کلمات کلیدی دیگری امتحان کنید.
          </p>
          <div class="space-y-2">
            <p class="text-sm text-gray-500">پیشنهادات:</p>
            <ul class="text-sm text-gray-600 space-y-1">
              <li>• املای کلمات را بررسی کنید</li>
              <li>• از کلمات کلیدی عمومی‌تر استفاده کنید</li>
              <li>• تعداد کلمات کلیدی را کاهش دهید</li>
            </ul>
          </div>
        </div>

        <div v-else-if="paginatedResults.length > 0" class="space-y-6">
          <article
            v-for="result in paginatedResults"
            :key="`${result.type}-${result.id}`"
            class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 hover:shadow-md transition-shadow"
          >
            <div class="flex items-start space-x-4 space-x-reverse">
              <div class="flex-shrink-0">
                <img
                  v-if="result.image"
                  :src="result.image"
                  :alt="result.title"
                  class="w-20 h-20 object-cover rounded-lg"
                />
                <div
                  v-else
                  class="w-20 h-20 bg-gray-200 rounded-lg flex items-center justify-center"
                >
                  <svg class="w-8 h-8 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                  </svg>
                </div>
              </div>

              <div class="flex-1 min-w-0">
                <div class="flex items-start justify-between">
                  <div class="flex-1">
                    <div class="flex items-center space-x-2 space-x-reverse mb-2">
                      <h3 class="text-lg font-semibold text-gray-900">
                        <NuxtLink :to="result.url" class="hover:text-blue-600 transition-colors">
                          {{ result.title }}
                        </NuxtLink>
                      </h3>
                      <span
                        :class="[
                          'text-xs px-2 py-1 rounded-full',
                          getTypeBadgeClass(result.type)
                        ]"
                      >
                        {{ getTypeLabel(result.type) }}
                      </span>
                    </div>

                    <p v-if="result.description" class="text-gray-600 mb-3 line-clamp-2">
                      {{ result.description }}
                    </p>

                    <div class="flex items-center space-x-4 space-x-reverse text-sm text-gray-500">
                      <span v-if="result.price" class="text-green-600 font-medium">
                        {{ formatPrice(result.price) }}
                      </span>

                      <span v-if="result.data?.sku" class="flex items-center">
                        <svg class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"></path>
                        </svg>
                        {{ result.data.sku }}
                      </span>

                      <span v-if="result.data?.category?.name" class="flex items-center">
                        <svg class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"></path>
                        </svg>
                        {{ result.data.category.name }}
                      </span>
                    </div>
                  </div>

                  <div class="flex items-center space-x-2 space-x-reverse">
                    <NuxtLink
                      :to="result.url"
                      class="px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 transition-colors text-sm"
                    >
                      مشاهده
                    </NuxtLink>
                  </div>
                </div>
              </div>
            </div>
          </article>

          <div v-if="totalPages > 1" class="flex items-center justify-center space-x-2 space-x-reverse">
            <button
              @click="currentPage = Math.max(1, currentPage - 1)"
              :disabled="currentPage === 1"
              :class="[
                'px-3 py-2 rounded-lg text-sm transition-colors',
                currentPage === 1
                  ? 'bg-gray-100 text-gray-400 cursor-not-allowed'
                  : 'bg-white border border-gray-300 text-gray-700 hover:bg-gray-50'
              ]"
            >
              قبلی
            </button>

            <div class="flex items-center space-x-1 space-x-reverse">
              <template v-for="(page, index) in visiblePages" :key="`${page}-${index}`">
                <button
                  v-if="page !== '...'"
                  @click="currentPage = page"
                  :class="[
                    'px-3 py-2 rounded-lg text-sm transition-colors',
                    page === currentPage
                      ? 'bg-blue-500 text-white'
                      : 'bg-white border border-gray-300 text-gray-700 hover:bg-gray-50'
                  ]"
                >
                  {{ page }}
                </button>
                <span v-else class="px-2 text-sm text-gray-400">...</span>
              </template>
            </div>

            <button
              @click="currentPage = Math.min(totalPages, currentPage + 1)"
              :disabled="currentPage === totalPages"
              :class="[
                'px-3 py-2 rounded-lg text-sm transition-colors',
                currentPage === totalPages
                  ? 'bg-gray-100 text-gray-400 cursor-not-allowed'
                  : 'bg-white border border-gray-300 text-gray-700 hover:bg-gray-50'
              ]"
            >
              بعدی
            </button>
          </div>
        </div>
      </section>
    </main>
  </div>
</template>

<script lang="ts">
declare const useRoute: () => { query: Record<string, string | string[] | undefined>; fullPath: string; params: Record<string, string> }
declare const $fetch: <T = unknown>(url: string, options?: { query?: Record<string, string | number> }) => Promise<T>
declare const useHead: (head: () => { title?: string; meta?: Array<{ name?: string; content?: string }> }) => void
</script>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue';

type SearchType = 'all' | 'products' | 'posts' | 'categories' | 'brands'
type PaginationItem = number | '...'

const route = useRoute()

const searchQuery = ref('')
const searchType = ref<SearchType>('all')
const searchResults = ref<any[]>([])
const isLoading = ref(false)
const searchTime = ref(0)
const currentPage = ref(1)
const itemsPerPage = ref(10)

const allowedTypes: SearchType[] = ['all', 'products', 'posts', 'categories', 'brands']

const totalResults = computed(() => searchResults.value.length)
const totalPages = computed(() => Math.ceil(totalResults.value / itemsPerPage.value) || 0)
const paginatedResults = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage.value
  const end = start + itemsPerPage.value
  return searchResults.value.slice(start, end)
})

const visiblePages = computed<PaginationItem[]>(() => {
  const total = totalPages.value
  const current = currentPage.value
  if (total <= 7) return Array.from({ length: total }, (_, i) => i + 1)
  if (current < 5) return [1, 2, 3, 4, 5, '...', total]
  if (current > total - 4) return [1, '...', total - 4, total - 3, total - 2, total - 1, total]
  return [1, '...', current - 1, current, current + 1, '...', total]
})

const performSearch = async () => {
  const query = searchQuery.value.trim()
  if (query.length < 2) {
    searchResults.value = []
    searchTime.value = 0
    return
  }

  try {
    isLoading.value = true
    const startTime = Date.now()

    const params: Record<string, string | number> = { q: query, limit: 100 }
    if (searchType.value !== 'all') {
      params.type = searchType.value
    }

    const response = await $fetch('/api/search', { query: params }) as any
    if (response?.success) {
      searchResults.value = response.data.map((item: any) => ({
        ...item,
        image: resolveImage(item)
      }))
      searchTime.value = Date.now() - startTime
      currentPage.value = 1
    } else {
      searchResults.value = []
      searchTime.value = 0
    }
  } catch (error) {
    searchResults.value = []
    searchTime.value = 0
  } finally {
    isLoading.value = false
  }
}

const applyRouteQuery = () => {
  const queryParam = (route.query.q as string) || ''
  const typeParam = route.query.type as SearchType | undefined

  searchQuery.value = queryParam
  searchType.value = typeParam && allowedTypes.includes(typeParam) ? typeParam : 'all'
  currentPage.value = 1

  if (queryParam) {
    performSearch()
  } else {
    searchResults.value = []
    searchTime.value = 0
  }
}

watch(currentPage, () => {
  window.scrollTo({ top: 0, behavior: 'smooth' })
})

watch(
  () => route.fullPath,
  () => {
    applyRouteQuery()
  }
)

onMounted(() => {
  applyRouteQuery()
})

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

const formatPrice = (price: number) => {
  return new Intl.NumberFormat('fa-IR').format(price) + ' تومان'
}

const resolveImage = (item: any) => {
  const candidates = [
    item.image,
    item.thumbnail,
    item.image_url,
    item.main_image,
    item.featured_image,
    item.cover,
    item.data?.thumbnail,
    item.data?.image,
    item.data?.image_url,
  ]

  const src = candidates.find((val) => typeof val === 'string' && val.trim().length > 0)
  if (!src) {
    return '/statics/images/default-image_100.png'
  }
  return src
}

// SEO Head
useHead(() => ({
  title: searchQuery.value
    ? `نتایج جستجو برای "${searchQuery.value}"`
    : 'نتایج جستجو',
  meta: [
    {
      name: 'description',
      content: searchQuery.value
        ? `نمایش نتایج مرتبط با "${searchQuery.value}"`
        : 'نمایش نتایج جستجو'
    }
  ]
}))
</script>

<style scoped>
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>