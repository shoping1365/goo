<template>
  <div class="min-h-screen bg-gray-50 py-8">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <!-- لایوت اصلی با سایدبار -->
      <div class="flex flex-col lg:flex-row gap-8 justify-center">
        <!-- محتوای اصلی -->
        <div class="flex-1 max-w-4xl">
          <!-- عنوان و تب‌ها داخل کارت سفید -->
          <div class="bg-white rounded-lg border border-gray-200 p-6 mb-6">
            <h1 class="text-sm font-bold text-gray-900 text-right mb-4">دیدگاه ها و پرسش ها</h1>
            <div class="flex justify-end space-x-6 md:space-x-10 space-x-reverse border-b border-gray-200">
              <button
                :class="[
                  'pb-4 px-3 font-medium text-sm border-b-2',
                  activeTab === 'awaiting' ? 'border-red-500 text-red-600' : 'border-transparent text-gray-500 hover:text-gray-700'
                ]"
                @click="switchTab('awaiting')"
              >
                در انتظار دیدگاه
              </button>
              <button
                :class="[
                  'pb-4 px-3 font-medium text-sm border-b-2',
                  activeTab === 'my-reviews' ? 'border-red-500 text-red-600' : 'border-transparent text-gray-500 hover:text-gray-700'
                ]"
                @click="switchTab('my-reviews')"
              >
                دیدگاه های من
              </button>
              <button
                :class="[
                  'pb-4 px-3 font-medium text-sm border-b-2',
                  activeTab === 'my-questions' ? 'border-red-500 text-red-600' : 'border-transparent text-gray-500 hover:text-gray-700'
                ]"
                @click="switchTab('my-questions')"
              >
                پرسشهای من
              </button>
              <button
                :class="[
                  'pb-4 px-3 font-medium text-sm border-b-2',
                  activeTab === 'all' ? 'border-red-500 text-red-600' : 'border-transparent text-gray-500 hover:text-gray-700'
                ]"
                @click="switchTab('all')"
              >
                همه
              </button>
            </div>
          </div>

          <!-- وضعیت بارگذاری/خالی -->
          <div v-if="loading" class="text-center text-gray-500 py-8">در حال بارگذاری ...</div>
          <div v-else-if="items.length === 0" class="text-center text-gray-500 py-12">موردی برای نمایش وجود ندارد.</div>

          <!-- لیست آیتم‌ها -->
          <div v-else class="space-y-4">
            <div v-for="item in items" :key="item.id" class="bg-white rounded-lg border border-gray-200 shadow-sm p-6">
              <div class="flex items-start space-x-4 space-x-reverse">
                <img :src="item.product?.image || item.product?.thumbnail || '/default-product.svg'" :alt="item.product?.name || 'product'" class="w-16 h-16 object-cover rounded-lg">
                <div class="flex-1">
                  <h3 class="text-sm font-medium text-gray-900 mb-1 text-right">{{ item.product?.name || item.title || 'بدون عنوان' }}</h3>
                  <p v-if="item.variant" class="text-xs text-gray-500 mb-2 text-right">{{ item.variant }}</p>
                  <div class="flex items-center text-xs text-gray-500 mb-3">
                    <svg class="w-4 h-4 ml-1" fill="currentColor" viewBox="0 0 20 20">
                      <path fill-rule="evenodd" d="M5.05 4.05a7 7 0 119.9 9.9L10 18.9l-4.95-4.95a7 7 0 010-9.9zM10 11a2 2 0 100-4 2 2 0 000 4z" clip-rule="evenodd" />
                    </svg>
                    <span>{{ formatMeta(item) }}</span>
                  </div>

                  <div v-if="activeTab === 'awaiting'" class="mt-2">
                    <button class="w-full border border-red-500 text-red-600 py-2 px-4 rounded-lg hover:bg-red-50 transition-colors flex items-center justify-center space-x-2 space-x-reverse">
                      <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                        <path fill-rule="evenodd" d="M18 10c0 3.866-3.582 7-8 7a8.841 8.841 0 01-4.083-.98L2 17l1.338-3.123C2.493 12.767 2 11.434 2 10c0-3.866 3.582-7 8-7s8 3.134 8 7zM7 9H5v2h2V9zm8 0h-2v2h2V9zM9 9h2v2H9V9z" clip-rule="evenodd" />
                      </svg>
                      <span>ثبت دیدگاه</span>
                    </button>
                  </div>
                  <div v-else>
                    <p v-if="item.comment" class="text-sm text-gray-700 mb-2 text-right">{{ item.comment }}</p>
                    <div v-if="item.rating" class="text-yellow-500 text-sm text-right">امتیاز: {{ item.rating }}/5</div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- سایدبار -->
        <div class="lg:w-80 flex-shrink-0 lg:ml-8">
          <AccountSidebar />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import AccountSidebar from '@/components/account/AccountSidebar.vue'

// متا
definePageMeta({ layout: 'default' })
useHead({ title: 'دیدگاه ها و پرسش ها - حساب کاربری' })

// تب‌ها
const activeTab = ref('awaiting')

// داده‌ها
const items = ref([])
const loading = ref(false)

function formatMeta(item) {
  const source = item.shop_name || item.seller_name || item.source || ''
  const date = item.created_at || item.createdAt || item.date
  let faDate = ''
  try { faDate = date ? new Date(date).toLocaleDateString('fa-IR') : '' } catch {}
  return [source, faDate].filter(Boolean).join('  •  ')
}

function mapTabToQuery(tab) {
  if (tab === 'awaiting') return { status: 'pending' }
  if (tab === 'my-reviews' ) return { mine: 'true' }
  if (tab === 'my-questions') return { type: 'question', mine: 'true' }
  return {}
}

async function load() {
  loading.value = true
  try {
    const params = mapTabToQuery(activeTab.value)
    const res = await $fetch('/api/reviews', { params, credentials: 'include' })
    let list = []
    if (Array.isArray(res)) list = res
    else if (Array.isArray(res?.items)) list = res.items
    else if (Array.isArray(res?.data)) list = res.data
    else if (Array.isArray(res?.results)) list = res.results
    items.value = list.map(r => ({
      id: r.id || r.review_id || r._id || Math.random().toString(36).slice(2),
      product: r.product || { name: r.product_name, image: r.product_image, thumbnail: r.thumbnail },
      comment: r.comment || r.text || r.body,
      rating: r.rating || r.stars || r.score,
      created_at: r.created_at || r.createdAt || r.date,
      seller_name: r.seller_name || r.shop || r.vendor,
      variant: r.variant || r.sku || ''
    }))
  } catch (e) {
    console.error('Failed to load reviews', e)
    items.value = []
  } finally {
    loading.value = false
  }
}

function switchTab(tab) { if (activeTab.value !== tab) { activeTab.value = tab; load() } }

onMounted(load)
</script>
