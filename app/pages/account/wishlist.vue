<template>
  <div class="min-h-screen bg-gray-50 py-8">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <!-- لایوت اصلی با سایدبار -->
      <div class="flex flex-col lg:flex-row gap-8 justify-center">
        <!-- محتوای اصلی -->
        <div class="flex-1 max-w-4xl" dir="rtl">
          <!-- عنوان و تب‌های ناوبری -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6">
            <h1 class="text-sm font-bold text-gray-900 text-right mb-4">لیست ها</h1>
            <div class="flex border-b border-gray-200">
              <!-- تب لیست علاقه مندی -->
              <button 
                class="pb-4 px-2 text-gray-500 hover:text-gray-700 transition-colors ml-8"
                :class="{ 'text-red-600 border-b-2 border-red-600': activeTab === 'wishlist' }"
                @click="activeTab = 'wishlist'"
              >
                لیست علاقه مندی
              </button>
              
              <!-- تب لیست های دیگر -->
              <button 
                class="pb-4 px-2 text-gray-500 hover:text-gray-700 transition-colors ml-8"
                :class="{ 'text-red-600 border-b-2 border-red-600': activeTab === 'other' }"
                @click="activeTab = 'other'"
              >
                لیست های دیگر
              </button>
              
              <!-- تب اطلاع رسانی ها -->
              <button 
                class="pb-4 px-2 text-gray-500 hover:text-gray-700 transition-colors"
                :class="{ 'text-red-600 border-b-2 border-red-600': activeTab === 'notifications' }"
                @click="activeTab = 'notifications'"
              >
                اطلاع رسانی ها
              </button>
            </div>

            <!-- محتوای تب اطلاع رسانی ها -->
            <div v-if="activeTab === 'notifications'" class="mt-6">
              <div class="flex items-center justify-between p-6 bg-gray-50 rounded-lg">
                <div class="flex-1 text-right">
                  <p class="text-sm text-gray-700">اطلاع رسانی تخفیف و روبه اتمام بودن موجودی این کالاها</p>
                </div>
                <div class="flex-shrink-0 mr-4">
                  <label class="relative inline-flex items-center cursor-pointer">
                    <input type="checkbox" v-model="notificationsEnabled" class="sr-only peer">
                    <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
                  </label>
                </div>
              </div>
            </div>
          </div>

          <!-- بخش مرتب‌سازی و تعداد -->
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6">
            <div class="flex items-center justify-between">
              <!-- گزینه‌های مرتب‌سازی -->
              <div class="flex items-center space-x-3 space-x-reverse">
                <span class="text-sm text-gray-600">مرتب سازی</span>
                <div class="flex space-x-1 space-x-reverse">
                  <button 
                    v-for="sort in sortOptions" 
                    :key="sort.value"
                    @click="currentSort = sort.value"
                    class="px-2 py-1 text-xs transition-colors"
                    :class="currentSort === sort.value ? 'text-red-600 border-b-2 border-red-600' : 'text-gray-500 hover:text-gray-700'"
                  >
                    {{ sort.label }}
                  </button>
                </div>
              </div>
              
              <!-- تعداد کالاها -->
              <div class="text-sm text-gray-600">
                {{ products.length }} کالا
              </div>
            </div>
          </div>

          <!-- کارت‌های محصول -->
          <div v-if="loading" class="text-center text-gray-500">در حال بارگذاری ...</div>
          <div v-else-if="products.length === 0" class="text-center text-gray-500">هیچ محصولی در لیست علاقه‌مندی‌ها یافت نشد.</div>
          <div v-else class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div v-for="product in products" :key="product.id" class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 flex flex-col">
              <div class="mb-4 relative">
                <img :src="getProductImage(product.image)" :alt="product.name" class="w-full h-40 object-contain rounded-lg bg-gray-50">
              </div>
              <h3 class="text-sm font-medium text-gray-900 mb-2 text-right">{{ product.name }}</h3>
              
              <!-- وضعیت موجودی/قیمت یا ناموجود (انتقال ناموجود به بالا برای هم‌تراز شدن دکمه‌ها) -->
              <div class="mb-3">
                <template v-if="product.available">
                  <div v-if="product.stock_status && product.stock_status.includes('تنها')" class="text-xs text-red-500 mb-1 text-right">{{ product.stock_status }}</div>
                  <div class="text-lg font-bold text-gray-900">{{ formatPrice(product.price) }}</div>
                </template>
                <!-- ناموجود دیگر اینجا نمایش داده نمی‌شود -->
              </div>

              <!-- نمایش ناموجود بالای سطل آشغال (سمت چپ) -->
              <div v-if="!product.available" class="text-sm text-gray-500 text-left mb-2">ناموجود</div>

              <!-- دکمه‌های عملیات -->
              <div class="flex space-x-2 space-x-reverse mt-auto">
                <!-- دکمه اضافه به سبد (برای موجود) -->
                <button 
                  v-if="product.available"
                  class="flex-1 bg-red-600 text-white py-2 px-4 rounded-lg hover:bg-red-700 transition-colors text-sm flex items-center justify-center"
                >
                  <svg class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4m0 0L7 13m0 0l-2.5 5M7 13l2.5 5m6-5v6a2 2 0 01-2 2H9a2 2 0 01-2-2v-6m8 0V9a2 2 0 00-2-2H9a2 2 0 00-2 2v4.01"></path>
                  </svg>
                  اضافه به سبد
                </button>
                
                <!-- دکمه کالاهای مشابه (برای ناموجود) -->
                <button 
                  v-if="!product.available"
                  class="flex-1 bg-gray-400 text-white py-2 px-4 rounded-lg hover:bg-gray-500 transition-colors text-sm flex items-center justify-center"
                >
                  کالاهای مشابه
                </button>
                
                <!-- دکمه حذف -->
                <button 
                  @click="toggleFavorite(product.id)"
                  class="w-10 h-10 bg-gray-200 hover:bg-gray-300 rounded-lg flex items-center justify-center transition-colors"
                >
                  <svg class="w-5 h-5 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                  </svg>
                </button>
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
// تعریف متا صفحه
definePageMeta({
  layout: 'default',
  middleware: ['auth']
})

// حالت تب فعال
const activeTab = ref('wishlist')

// تنظیمات اطلاع‌رسانی
const notificationsEnabled = ref(true)

// گزینه‌های مرتب‌سازی
const sortOptions = [
  { label: 'جدیدترین', value: 'newest' },
  { label: 'گران ترین', value: 'expensive' },
  { label: 'ارزان ترین', value: 'cheapest' }
]

const currentSort = ref('newest')

// محصولات علاقه‌مندی
const products = ref([])
const loading = ref(false)

// تابع بارگذاری لیست علاقه‌مندی‌ها
async function loadWishlist() {
  try {
    loading.value = true
    console.log('درخواست دریافت محصولات علاقه‌مندی...')
    const response = await $fetch('/api/wishlist', { credentials: 'include' })
    console.log('پاسخ API علاقه‌مندی‌ها:', response)
    
    if (response.success) {
      products.value = response.data
      console.log('محصولات علاقه‌مندی دریافت شده:', products.value.length)
    } else {
      console.error('خطا در دریافت علاقه‌مندی‌ها:', response.message)
      products.value = []
    }
  } catch (e) {
    console.error('خطا در دریافت علاقه‌مندی‌ها', e)
    products.value = []
  } finally {
    loading.value = false
  }
}

// دریافت محصولات علاقه‌مندی
onMounted(async () => {
  await loadWishlist()
})

// توابع علاقه‌مندی و لیست
async function toggleFavorite(productId) {
  try {
    console.log('محصول به لیست علاقه‌مندی‌ها اضافه/حذف شد:', productId)
    
    // بررسی اینکه آیا محصول در لیست است یا نه
    const isInWishlist = products.value.some(p => p.id === productId)
    
    // دریافت CSRF token
    const { getCSRFToken } = useCSRF()
    const csrfToken = await getCSRFToken()
    
    if (!csrfToken) {
      console.error('CSRF token در دسترس نیست')
      return
    }

    if (isInWishlist) {
      // حذف از لیست
      const response = await $fetch(`/api/wishlist/remove/${productId}`, {
        method: 'DELETE',
        headers: {
          'x-csrf-token': csrfToken
        },
        credentials: 'include'
      })
      
      if (response.success) {
        // حذف از لیست محلی
        products.value = products.value.filter(p => p.id !== productId)
        console.log('محصول از لیست علاقه‌مندی‌ها حذف شد')
      }
    } else {
      // اضافه کردن به لیست
      const response = await $fetch('/api/wishlist/add', {
        method: 'POST',
        headers: {
          'x-csrf-token': csrfToken
        },
        credentials: 'include',
        body: { product_id: productId }
      })
      
      if (response.success) {
        console.log('محصول به لیست علاقه‌مندی‌ها اضافه شد')
        // بارگذاری مجدد لیست
        await loadWishlist()
      }
    }
  } catch (error) {
    console.error('خطا در تغییر وضعیت علاقه‌مندی:', error)
  }
}

function addToList(productId) {
  console.log('محصول به لیست اضافه شد:', productId)
  // TODO: نمایش مودال انتخاب لیست یا ارسال درخواست به سرور
}

function isFavorite(productId) {
  // بررسی اینکه آیا محصول در لیست علاقه‌مندی‌ها است
  return products.value.some(p => p.id === productId)
}

function formatPrice(price) {
  if (!price || isNaN(price)) return 'قیمت نامشخص'
  return new Intl.NumberFormat('fa-IR').format(price) + ' تومان'
}

function getProductImage(imagePath) {
  if (!imagePath) return '/default-product.svg'
  
  // اگر مسیر کامل است
  if (imagePath.startsWith('http')) return imagePath
  
  // اگر مسیر نسبی است
  if (imagePath.startsWith('/uploads/')) return imagePath
  
  // اگر فقط نام فایل است
  return `/uploads/media/products/images/${imagePath}`
}

// تنظیم عنوان صفحه
useHead({
  title: 'لیست های من - فروشگاه آنلاین',
  meta: [
    { name: 'description', content: 'مدیریت لیست‌های علاقه‌مندی و محصولات' }
  ]
})
</script>

<style scoped>
/* استایل‌های اضافی برای بهبود ظاهر */
.space-x-reverse > :not([hidden]) ~ :not([hidden]) {
  --tw-space-x-reverse: 1;
  margin-right: calc(1rem * var(--tw-space-x-reverse));
  margin-left: calc(1rem * calc(1 - var(--tw-space-x-reverse)));
}
</style>
