<template>
  <div class="space-y-6" dir="rtl">
    <!-- Reviews Summary -->
    <div class="bg-gradient-to-r from-blue-50 to-purple-50 rounded-lg px-4 py-4">
      <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
        <!-- Overall Rating -->
        <div class="text-center">
          <div class="text-4xl font-bold text-gray-900 mb-2">{{ averageRating.toFixed(1) }}</div>
          <div class="flex items-center justify-center gap-1 mb-2">
            <span v-for="i in 5" :key="i" class="text-yellow-400">
              <svg 
                :class="[
                  'w-6 h-6',
                  i <= Math.round(averageRating) ? 'fill-current' : 'fill-gray-200'
                ]"
                viewBox="0 0 24 24"
              >
                <path d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"/>
              </svg>
            </span>
          </div>
          <div class="text-gray-600">
            بر اساس {{ totalReviews }} نظر
          </div>
        </div>

        <!-- Rating Breakdown -->
        <div class="space-y-2">
          <div v-for="rating in [5, 4, 3, 2, 1]" :key="rating" class="flex items-center gap-3">
            <span class="text-sm font-medium w-8">{{ rating }}</span>
            <svg class="w-4 h-4 text-yellow-400 fill-current" viewBox="0 0 24 24">
              <path d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"/>
            </svg>
            <div class="flex-1 bg-gray-200 rounded-full h-2">
              <div 
                class="bg-yellow-400 h-2 rounded-full transition-all"
                :style="{ width: `${getRatingPercentage(rating)}%` }"
              ></div>
            </div>
            <span class="text-sm text-gray-500 w-12">{{ getRatingCount(rating) }}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Filters and Actions -->
    <div class="flex flex-col sm:flex-row gapx-4 py-4 items-start sm:items-center justify-between">
      <div class="flex flex-wrap gap-3">
        <!-- Rating Filter -->
        <select 
          v-model="selectedRating"
          class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
        >
          <option value="">همه امتیازها</option>
          <option v-for="rating in [5, 4, 3, 2, 1]" :key="rating" :value="rating">
            {{ rating }} ستاره
          </option>
        </select>

        <!-- Sort Filter -->
        <select 
          v-model="sortBy"
          class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
        >
          <option value="newest">جدیدترین</option>
          <option value="oldest">قدیمی‌ترین</option>
          <option value="highest">بالاترین امتیاز</option>
          <option value="lowest">پایین‌ترین امتیاز</option>
          <option value="helpful">مفیدترین</option>
        </select>

        <!-- Verified Filter -->
        <label class="flex items-center gap-2 px-4 py-2 border border-gray-300 rounded-lg cursor-pointer">
          <input 
            v-model="showVerifiedOnly"
            type="checkbox"
            class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
          />
          <span class="text-sm">فقط خریداران</span>
        </label>
      </div>

      <!-- Add Review Button -->
      <button
        @click="onAddReviewClick"
        class="bg-blue-600 hover:bg-blue-700 text-white px-6 py-2 rounded-lg font-medium transition-colors flex items-center gap-2"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
        </svg>
        نظر جدید
      </button>
    </div>

    <!-- Reviews List -->
    <div class="space-y-4">
      <div 
        v-for="review in filteredReviews" 
        :key="review.id"
        class="bg-white rounded-lg border border-gray-200 px-4 py-4 hover:shadow-md transition-shadow"
      >
        <!-- Review Header -->
        <div class="flex items-start justify-between mb-4">
          <div class="flex items-start gapx-4 py-4">
            <!-- Avatar -->
            <div class="w-12 h-12 bg-gradient-to-r from-blue-500 to-purple-500 rounded-full flex items-center justify-center text-white font-bold">
              {{ review.user_name ? review.user_name.charAt(0).toUpperCase() : '؟' }}
            </div>
            
            <!-- User Info -->
            <div>
              <div class="flex items-center gap-2">
                <span class="font-medium text-gray-900">{{ review.user_name }}</span>
                <span v-if="review.is_verified" class="inline-flex items-center gap-1 text-green-600 text-sm">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                  </svg>
                  خریدار محصول
                </span>
              </div>
              
              <!-- Rating and Date -->
              <div class="flex items-center gap-3 mt-1">
                <div class="flex items-center gap-1">
                  <span v-for="i in 5" :key="i" class="text-yellow-400">
                    <svg 
                      :class="[
                        'w-4 h-4',
                        i <= review.rating ? 'fill-current' : 'fill-gray-200'
                      ]"
                      viewBox="0 0 24 24"
                    >
                      <path d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"/>
                    </svg>
                  </span>
                </div>
                <span class="text-sm text-gray-500">{{ formatDate(review.created_at) }}</span>
              </div>
            </div>
          </div>

          <!-- Review Actions -->
          <div class="flex items-center gap-2">
            <button
              @click="toggleHelpful(review)"
              :class="[
                'flex items-center gap-1 px-3 py-1 rounded-lg text-sm transition-colors',
                review.user_found_helpful 
                  ? 'bg-green-100 text-green-700' 
                  : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
              ]"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 10h4.764a2 2 0 011.789 2.894l-3.5 7A2 2 0 0115.263 21h-4.017c-.163 0-.326-.02-.485-.06L7 20m7-10V5a2 2 0 00-2-2h-.095c-.5 0-.905.405-.905.905 0 .714-.211 1.412-.608 2.006L7 11v9m7-10h-2M7 20H5a2 2 0 01-2-2v-6a2 2 0 012-2h2.5"></path>
              </svg>
              {{ review.helpful_count || 0 }}
            </button>
          </div>
        </div>

        <!-- Review Title -->
        <h4 v-if="review.title" class="font-medium text-gray-900 mb-2">
          {{ review.title }}
        </h4>

        <!-- Review Content -->
        <div class="text-gray-700 mb-4 leading-relaxed">
          {{ review.comment }}
        </div>

        <!-- Review Images -->
        <div v-if="review.images && review.images.length" class="flex gap-2 mb-4 overflow-x-auto">
          <img 
            v-for="(image, index) in review.images"
            :key="index"
            :src="image"
            :alt="`تصویر نظر ${index + 1}`"
            class="w-20 h-20 object-cover rounded-lg cursor-pointer hover:opacity-75 transition-opacity"
            @click="openImageModal(image)"
          />
        </div>

        <!-- Pros and Cons -->
        <div v-if="review.pros || review.cons" class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4 mb-4">
          <div v-if="review.pros" class="bg-green-50 rounded-lg p-3">
            <h5 class="font-medium text-green-800 mb-2">نکات مثبت:</h5>
            <p class="text-green-700 text-sm">{{ review.pros }}</p>
          </div>
          
          <div v-if="review.cons" class="bg-red-50 rounded-lg p-3">
            <h5 class="font-medium text-red-800 mb-2">نکات منفی:</h5>
            <p class="text-red-700 text-sm">{{ review.cons }}</p>
          </div>
        </div>

        <!-- Admin/Seller Reply -->
        <div v-if="review.reply" class="bg-blue-50 rounded-lg px-4 py-4 mr-8">
          <div class="flex items-center gap-2 mb-2">
            <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
            </svg>
            <span class="font-medium text-blue-800">پاسخ فروشنده</span>
            <span class="text-sm text-blue-600">{{ formatDate(review.reply.created_at) }}</span>
          </div>
          <p class="text-blue-700">{{ review.reply.content }}</p>
        </div>
      </div>
    </div>

    <!-- Empty State -->
    <div v-if="!reviews.length" class="text-center py-12">
      <div class="text-gray-400 mb-4">
        <svg class="w-16 h-16 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
        </svg>
      </div>
      <h3 class="text-lg font-medium text-gray-900 mb-2">هنوز نظری ثبت نشده</h3>
      <p class="text-gray-500 mb-4">اولین نفری باشید که نظرتان را درباره این محصول ثبت می‌کنید.</p>
      <button
        @click="onAddReviewClick"
        class="bg-blue-600 hover:bg-blue-700 text-white px-6 py-3 rounded-lg font-medium transition-colors"
      >
        ثبت نظر
      </button>
    </div>

    <!-- Review Form Modal -->
    <div v-if="showReviewForm" class="fixed inset-0 z-50 bg-black bg-opacity-50 flex items-center justify-center px-4 py-4">
      <div class="bg-white rounded-lg shadow-xl w-full max-w-2xl max-h-[90vh] overflow-y-auto">
        <div class="px-4 py-4">
          <div class="flex items-center justify-between mb-6">
            <h3 class="text-xl font-bold text-gray-900">نظر شما درباره این محصول</h3>
            <button
              @click="closeReviewForm"
              class="text-gray-400 hover:text-gray-600"
            >
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
              </svg>
            </button>
          </div>

          <form @submit.prevent="submitReview" class="space-y-6">
            <!-- Rating -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">امتیاز شما</label>
              <div class="flex items-center gap-1">
                <button
                  v-for="i in 5"
                  :key="i"
                  type="button"
                  @click="newReview.rating = i"
                  class="text-3xl transition-colors"
                  :class="i <= newReview.rating ? 'text-yellow-400' : 'text-gray-300 hover:text-yellow-300'"
                >
                  ★
                </button>
                <span class="mr-3 text-sm text-gray-600">
                  {{ getRatingText(newReview.rating) }}
                </span>
              </div>
            </div>

            <!-- Title -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">عنوان نظر</label>
              <input
                v-model="newReview.title"
                type="text"
                placeholder="عنوان کوتاهی برای نظرتان بنویسید..."
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>

            <!-- Comment -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نظر شما *</label>
              <textarea
                v-model="newReview.comment"
                rows="4"
                placeholder="نظر خود را درباره این محصول بنویسید..."
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 resize-none"
              ></textarea>
            </div>

            <!-- File Upload (Images & Videos) -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">تصاویر / ویدیوها (حداکثر ۷ فایل)</label>
              <input 
                ref="fileInput"
                type="file" 
                multiple 
                accept="image/*,video/mp4,video/webm,video/quicktime" 
                class="block w-full text-sm text-gray-700 file:mr-4 file:py-2 file:px-4 file:border-0 file:text-sm file:font-semibold file:bg-blue-50 file:text-blue-700 hover:file:bg-blue-100"/>
            </div>

            <!-- Pros -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نکات مثبت</label>
              <textarea
                v-model="newReview.pros"
                rows="2"
                placeholder="چه چیزهایی در این محصول خوب بود؟"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 resize-none"
              ></textarea>
            </div>

            <!-- Cons -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نکات منفی</label>
              <textarea
                v-model="newReview.cons"
                rows="2"
                placeholder="چه چیزهایی قابل بهبود است؟"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 resize-none"
              ></textarea>
            </div>

            <!-- Anonymous option -->
            <div class="flex items-center gap-2">
              <input v-model="newReview.anonymous" type="checkbox" id="anon" class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500" />
              <label for="anon" class="text-sm text-gray-700">ثبت نظر به صورت ناشناس</label>
            </div>

            <!-- Submit Button -->
            <div class="flex gap-3">
              <button
                type="submit"
                :disabled="submittingReview || !newReview.comment || !newReview.rating"
                class="flex-1 bg-blue-600 hover:bg-blue-700 disabled:bg-gray-300 text-white py-3 rounded-lg font-medium transition-colors"
              >
                {{ submittingReview ? 'در حال ثبت...' : 'ثبت نظر' }}
              </button>
              <button
                type="button"
                @click="closeReviewForm"
                class="px-6 py-3 border border-gray-300 rounded-lg hover:bg-gray-50 transition-colors"
              >
                انصراف
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>

    <!-- Image Modal -->
    <div v-if="showImageModal" class="fixed inset-0 z-50 bg-black bg-opacity-75 flex items-center justify-center px-4 py-4" @click="closeImageModal">
      <img 
        :src="modalImage" 
        :alt="product.name"
        class="max-w-full max-h-full object-contain"
        @click.stop
      />
      <button 
        @click="closeImageModal"
        class="absolute topx-4 py-4 left-4 bg-black bg-opacity-50 hover:bg-opacity-75 text-white p-2 rounded-full"
      >
        <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
        </svg>
      </button>
    </div>
  </div>
</template>

<script lang="ts">
declare const navigateTo: (to: string) => Promise<void>
declare const useFetch: <T = unknown>(url: string, options?: { method?: string; query?: Record<string, string | number | undefined>; server?: boolean; lazy?: boolean; key?: string; immediate?: boolean; body?: FormData | unknown }) => Promise<{ data: { value: T }; error: { value: Error | null } }>
declare const useAuth: () => { isAuthenticated: { value: boolean } }
</script>

<script setup lang="ts">
import { computed, onMounted, reactive, ref, watch } from 'vue';

// Auth disabled
const { isAuthenticated } = useAuth()
interface Review {
  id: number
  user_name: string
  rating: number
  title?: string
  comment: string
  pros?: string
  cons?: string
  images?: string[]
  is_verified: boolean
  helpful_count: number
  user_found_helpful: boolean
  created_at: string
  reply?: {
    content: string
    created_at: string
  }
}

interface Product {
  id: number
  name: string
  reviews?: Review[]
}

interface Props {
  product: Product
}

const props = defineProps<Props>()

// State
const reviews = ref<Review[]>([])
const selectedRating = ref('')
const sortBy = ref('newest')
const showVerifiedOnly = ref(false)
const loading = ref(false)

// Review Form
const showReviewForm = ref(false)
const submittingReview = ref(false)
const newReview = reactive({
  rating: 0,
  title: '',
  comment: '',
  pros: '',
  cons: '',
  anonymous: false
})

// Image Modal
const showImageModal = ref(false)
const modalImage = ref('')

// Ref for file input
const fileInput = ref<HTMLInputElement | null>(null)
// Auth disabled
const onAddReviewClick = async () => {
  if (!isAuthenticated.value) {
    await navigateTo(`/auth/login?callbackUrl=${encodeURIComponent(location.pathname + location.search)}`)
    return
  }
  showReviewForm.value = true
}

// Computed
const averageRating = computed(() => {
  if (!reviews.value.length) return 0
  const sum = reviews.value.reduce((acc, review) => acc + review.rating, 0)
  return sum / reviews.value.length
})

const totalReviews = computed(() => reviews.value.length)

const filteredReviews = computed(() => {
  let filtered = [...reviews.value]

  // Filter by rating
  if (selectedRating.value) {
    filtered = filtered.filter(review => review.rating === parseInt(selectedRating.value))
  }

  // Filter by verified buyers
  if (showVerifiedOnly.value) {
    filtered = filtered.filter(review => review.is_verified)
  }

  // Sort
  filtered.sort((a, b) => {
    switch (sortBy.value) {
      case 'newest':
        return new Date(b.created_at).getTime() - new Date(a.created_at).getTime()
      case 'oldest':
        return new Date(a.created_at).getTime() - new Date(b.created_at).getTime()
      case 'highest':
        return b.rating - a.rating
      case 'lowest':
        return a.rating - b.rating
      case 'helpful':
        return (b.helpful_count || 0) - (a.helpful_count || 0)
      default:
        return 0
    }
  })

  return filtered
})

// Methods
const fetchReviews = async () => {
  loading.value = true
  const { data, error } = await useFetch(`/api/products/${props.product.id}/reviews`, {
    method: 'GET',
    query: {
      status: 'approved',
      rating: selectedRating.value || undefined
    },
    server: false,
    lazy: true,
    key: `product-reviews-${props.product.id}-${selectedRating.value}`,
  })
  reviews.value = (data.value as any)?.reviews || []
  loading.value = false
}

const getRatingCount = (rating: number): number => {
  return reviews.value.filter(review => review.rating === rating).length
}

const getRatingPercentage = (rating: number): number => {
  const count = getRatingCount(rating)
  return totalReviews.value > 0 ? (count / totalReviews.value) * 100 : 0
}

const formatDate = (dateString: string): string => {
  const date = new Date(dateString)
  return date.toLocaleDateString('fa-IR')
}

const getRatingText = (rating: number): string => {
  const texts = {
    1: 'خیلی بد',
    2: 'بد',
    3: 'متوسط',
    4: 'خوب',
    5: 'عالی'
  }
  return texts[rating] || ''
}

const toggleHelpful = async (review: Review) => {
  const { error } = await useFetch(`/api/reviews/${review.id}/helpful`, { method: 'POST', immediate: true })
  if (!error.value) {
    review.user_found_helpful = !review.user_found_helpful
    review.helpful_count += review.user_found_helpful ? 1 : -1
  }
}

const submitReview = async () => {
  if (!newReview.comment || !newReview.rating) return

  submittingReview.value = true
  try {
    const formData = new FormData()
    formData.append('rating', String(newReview.rating))
    formData.append('comment', newReview.comment)
    formData.append('product_id', String(props.product.id))
    formData.append('anonymous', newReview.anonymous ? 'true' : 'false')
    if (newReview.title) formData.append('title', newReview.title)
    if (newReview.pros) formData.append('pros', newReview.pros)
    if (newReview.cons) formData.append('cons', newReview.cons)

    const files = fileInput.value?.files
    if (files) {
      Array.from(files).slice(0, 7).forEach(f => {
        if (f.type.startsWith('image/')) formData.append('images', f)
        else if (f.type.startsWith('video/')) formData.append('videos', f)
      })
    }

    const { data, error } = await useFetch('/api/reviews', {
      method: 'POST',
      body: formData,
      immediate: true
    })
    if (!error.value && data.value) {
      reviews.value.unshift(data.value as unknown as Review)
      closeReviewForm()
    }
  } finally {
    submittingReview.value = false
  }

  // Reset file input
  if (fileInput.value) fileInput.value.value = ''
}

const closeReviewForm = () => {
  showReviewForm.value = false
  Object.assign(newReview, {
    rating: 0,
    title: '',
    comment: '',
    pros: '',
    cons: '',
    anonymous: false
  })
}

const openImageModal = (image: string) => {
  modalImage.value = image
  showImageModal.value = true
}

const closeImageModal = () => {
  showImageModal.value = false
  modalImage.value = ''
}

// Lifecycle
onMounted(() => {
  fetchReviews()
})

// Watch for filter changes
watch([selectedRating, showVerifiedOnly], () => {
  fetchReviews()
})
</script>

<style scoped>
/* Custom scrollbar for image gallery */
.overflow-x-auto::-webkit-scrollbar {
  height: 4px;
}

.overflow-x-auto::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 2px;
}

.overflow-x-auto::-webkit-scrollbar-thumb {
  background: #c1c1c1;
  border-radius: 2px;
}

.overflow-x-auto::-webkit-scrollbar-thumb:hover {
  background: #a1a1a1;
} 
</style> 
