<template>
  <div class="min-h-screen bg-gray-50 py-8">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <!-- هدر صفحه -->
      <div class="mb-8 text-right">
        <h1 class="text-3xl font-bold text-gray-900 mb-2">خلاصه فعالیت ها</h1>
        <p class="text-gray-600">مروری بر فعالیت‌ها و وضعیت حساب کاربری شما</p>
      </div>

      <!-- لایوت اصلی با سایدبار -->
      <div class="flex flex-col lg:flex-row gap-8 justify-center">
        <!-- محتوای اصلی -->
        <div class="flex-1 max-w-4xl" dir="rtl">
          <!-- بخش سفارش‌های من -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-8 text-right">
            <div class="flex items-center justify-between mb-4">
              <h2 class="text-xl font-semibold text-gray-900">سفارش‌های من</h2>
              <NuxtLink to="/account/orders" class="text-blue-600 hover:text-blue-700 text-sm font-medium transition-colors">
                مشاهده همه
              </NuxtLink>
            </div>

            <!-- وضعیت بارگذاری -->
            <div v-if="ordersLoading" class="py-6 text-gray-500">در حال بارگذاری ...</div>

            <!-- لیست سفارش‌ها (حداکثر 2 مورد) -->
            <div v-else-if="orders.length" class="divide-y divide-gray-100">
              <div v-for="o in orders.slice(0, 2)" :key="o.id" class="py-4">
                <div class="flex items-start justify-between gap-6">
                  <div class="flex-1">
                    <div class="flex items-center gap-3">
                      <span class="text-sm text-gray-500">شماره سفارش:</span>
                      <span class="text-sm font-semibold text-gray-900">{{ o.order_number }}</span>
                      <span class="text-xs text-gray-400">•</span>
                      <span class="text-xs text-gray-500">{{ formatDateFa(o.created_at) }}</span>
                    </div>
                    <div class="mt-2 flex items-center gap-2 flex-wrap">
                      <span
class="inline-flex items-center px-2 py-0.5 rounded-full text-xs border"
                            :class="statusClass(o.status)">
                        {{ statusLabel(o.status) }}
                      </span>
                      <span class="text-sm text-gray-700">مبلغ: {{ formatCurrencyFa(o.final_amount ?? o.total_amount) }}</span>
                    </div>

                    <!-- تصاویر آیتم‌ها -->
                    <div v-if="o.items?.length" class="mt-3 flex items-center gap-2">
                      <img v-for="(it, idx) in o.items.slice(0, 4)" :key="idx" :src="it.product_image || '/default-product.svg'" :alt="it.product_name" class="w-10 h-10 rounded object-cover">
                      <span v-if="o.items.length > 4" class="text-xs text-gray-500">+{{ o.items.length - 4 }}</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <!-- خالی -->
            <div v-else class="py-6 text-gray-500">در حال حاضر سفارشی برای نمایش وجود ندارد.</div>
          </div>

          <!-- بخش از لیست‌های شما (Wishlist) -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-8 text-right">
            <div class="flex items-center justify-between mb-4">
              <h2 class="text-xl font-semibold text-gray-900">از لیست‌های شما</h2>
              <NuxtLink to="/account/wishlist" class="text-blue-600 hover:text-blue-700 text-sm font-medium transition-colors">
                مشاهده همه
              </NuxtLink>
            </div>

            <!-- وضعیت بارگذاری -->
            <div v-if="wishlistLoading" class="py-6 text-gray-500">در حال بارگذاری ...</div>

            <!-- لیست علاقه‌مندی‌ها (حداکثر 3 مورد) -->
            <div v-else-if="wishlist.length" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
              <div v-for="p in wishlist.slice(0, 3)" :key="p.id" class="rounded-lg p-3 flex items-center gap-3 hover:shadow-sm transition border border-blue-100 hover:border-blue-200">
                <img :src="p.image || '/default-product.svg'" :alt="p.name" class="w-14 h-14 rounded object-cover">
                <div class="flex-1 min-w-0">
                  <div class="text-sm font-medium text-gray-900 truncate">{{ p.name }}</div>
                  <div class="text-xs text-gray-500 mt-1">
                    <span v-if="p.price != null">{{ formatCurrencyFa(p.price) }}</span>
                    <span v-else>{{ p.stock_status || 'قیمت نامشخص' }}</span>
                  </div>
                </div>
              </div>
            </div>

            <!-- خالی -->
            <div v-else class="py-6 text-gray-500">لیست علاقه‌مندی‌های شما خالی است.</div>
          </div>

          <!-- بخش خریدهای پرتکرار شما -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 text-right">
            <div class="mb-4">
              <h2 class="text-xl font-semibold text-gray-900">خریدهای پرتکرار شما</h2>
            </div>
            <div class="py-6 text-gray-500">داده‌ای برای نمایش وجود ندارد.</div>
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
import AccountSidebar from '~/components/account/AccountSidebar.vue'
// احراز هویت - غیرفعال شده
// const user = null
// const isAuthenticated = false
// const checkAuth = null

// تعریف متا صفحه
definePageMeta({
  layout: 'default',
  middleware: ['auth']
})

// تنظیم عنوان صفحه
useHead({
  title: 'حساب کاربری - فروشگاه آنلاین',
  meta: [
    { name: 'description', content: 'مدیریت حساب کاربری، سفارشات و تنظیمات شخصی' }
  ]
})

// داده‌ها
const orders = ref([])
const ordersLoading = ref(false)
const wishlist = ref([])
const wishlistLoading = ref(false)

function formatDateFa(dateStr) {
  try { return new Date(dateStr).toLocaleDateString('fa-IR') } catch { return '' }
}
function formatCurrencyFa(val) {
  if (val == null) return ''
  try { return `${Number(val).toLocaleString('fa-IR')} تومان` } catch { return `${val} تومان` }
}
function statusLabel(s) {
  const map = { pending: 'در حال پردازش', paid: 'پرداخت شده', shipped: 'ارسال شده', delivered: 'تحویل شده', canceled: 'لغو شده' }
  return map[s] || s || 'نامشخص'
}
function statusClass(s) {
  const base = 'border-gray-300 text-gray-600'
  if (s === 'paid' || s === 'delivered') return 'border-green-300 text-green-700 bg-green-50'
  if (s === 'shipped') return 'border-blue-300 text-blue-700 bg-blue-50'
  if (s === 'canceled') return 'border-red-300 text-red-700 bg-red-50'
  return base
}

async function loadOrders() {
  ordersLoading.value = true
  try {
    const res = await $fetch('/api/orders', { credentials: 'include' })
    orders.value = Array.isArray(res?.data) ? res.data : (Array.isArray(res) ? res : [])
  } catch (e) {
    console.error('Failed to load orders', e)
    orders.value = []
  } finally {
    ordersLoading.value = false
  }
}

async function loadWishlist() {
  wishlistLoading.value = true
  try {
    // فقط اگر کاربر احراز هویت شده باشد، wishlist را بارگذاری کن
    const { isAuthenticated } = useAuthState()
    if (!isAuthenticated.value) {
      wishlist.value = []
      return
    }

    const res = await $fetch('/api/wishlist', { credentials: 'include' })
    wishlist.value = Array.isArray(res?.data) ? res.data : (Array.isArray(res) ? res : [])
  } catch (e) {
    console.error('Failed to load wishlist', e)
    wishlist.value = []
  } finally {
    wishlistLoading.value = false
  }
}

onMounted(() => { loadOrders(); loadWishlist() })
</script>

<style scoped>
/* استایل‌های اضافی برای بهبود ظاهر */
.space-x-reverse > :not([hidden]) ~ :not([hidden]) {
  --tw-space-x-reverse: 1;
  margin-right: calc(1rem * var(--tw-space-x-reverse));
  margin-left: calc(1rem * calc(1 - var(--tw-space-x-reverse)));
}

/* مخفی کردن اسکرول‌بار */
.scrollbar-hide {
  -ms-overflow-style: none;
  scrollbar-width: none;
}

.scrollbar-hide::-webkit-scrollbar {
  display: none;
}
</style>
