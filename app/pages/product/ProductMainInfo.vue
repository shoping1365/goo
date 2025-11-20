<template>
  <div class="grid grid-cols-1 lg:grid-cols-2 gap-8 px-4 py-4" dir="rtl">
    <!-- Image Gallery -->
    <div class="space-y-4">
      <div class="relative bg-gray-100 rounded-lg overflow-hidden aspect-square">
        <img 
          :src="mainImage" 
          :alt="product.name"
          class="w-full h-full object-cover"
        />
        
        <!-- Invisible clickable area for modal -->
        <div 
          class="absolute inset-0 cursor-pointer z-5"
          @click="openImageModal"
        ></div>
        
        <!-- Navigation Arrows -->
        <button 
          v-if="availableImages && availableImages.length > 1"
          class="absolute left-2 top-1/2 -translate-y-1/2 z-20 bg-blue-500 hover:bg-blue-600 rounded-full p-2 transition-all"
          @click.stop="previousImage"
        >
          <svg class="w-4 h-4 text-white" fill="currentColor" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
          </svg>
        </button>
        
        <button 
          v-if="availableImages && availableImages.length > 1"
          class="absolute right-2 top-1/2 -translate-y-1/2 z-20 bg-blue-500 hover:bg-blue-600 rounded-full p-2 transition-all"
          @click.stop="nextImage"
        >
          <svg class="w-4 h-4 text-white" fill="currentColor" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
          </svg>
        </button>
        
        <!-- آیکون‌های علاقه‌مندی و لیست -->
        <div class="absolute top-6 right-4 z-20 flex space-x-2 space-x-reverse">
          <!-- آیکون قلب برای علاقه‌مندی -->
          <button 
            class="w-10 h-10 bg-white rounded-full shadow-md flex items-center justify-center hover:bg-gray-50 transition-colors"
            @click="toggleFavorite"
          >
            <svg 
              class="w-5 h-5" 
              :class="isFavorite ? 'text-red-500 fill-current' : 'text-gray-400'"
              fill="none" 
              stroke="currentColor" 
              viewBox="0 0 24 24"
            >
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"></path>
            </svg>
          </button>
          <!-- آیکون لیست برای اضافه کردن به لیست -->
          <button 
            class="w-10 h-10 bg-white rounded-full shadow-md flex items-center justify-center hover:bg-gray-50 transition-colors"
            @click="addToList"
          >
            <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
            </svg>
          </button>
        </div>

        <!-- Admin Edit Button -->
        <div v-if="isAdmin" class="absolute top-6 left-4 z-20">
          <NuxtLink 
            :to="`/admin/product-management/products/edit/${product.id}`"
            target="_blank"
            class="bg-blue-600 hover:bg-blue-700 text-white p-2 rounded-full shadow-lg hover:shadow-xl transition-all duration-200 group flex items-center justify-center"
            title="ویرایش محصول"
            style="width: 2.5rem; height: 2.5rem;"
          >
            <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
              <path d="M13.586 3.586a2 2 0 112.828 2.828l-.793.793-2.828-2.828.793-.793zM11.379 5.793L3 14.172V17h2.828l8.38-8.379-2.83-2.828z"/>
            </svg>
          </NuxtLink>
        </div>
        
        <div v-if="showDiscountBadge" class="absolute topx-4 py-4 right-4">
          <span class="bg-red-500 text-white px-3 py-1 rounded-full text-sm font-bold">
            {{ discountPercentage }}% تخفیف
          </span>
        </div>

        <div class="absolute bottom-4 right-4 flex flex-col items-end gap-1">
          <span
v-if="!product.call_for_price" :class="[
            'px-3 py-1 rounded-full text-sm font-medium',
            isInStock ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'
          ]">
            {{ isInStock ? 'موجود' : 'ناموجود' }}
          </span>
        </div>
      </div>

      <div v-if="availableImages && availableImages.length > 1" class="flex gap-2 justify-center pb-2">
        <button
          v-for="(image, index) in availableImages.slice(0, Math.max(1, availableImages.length - 3))"
          :key="index"
          :class="[
            'w-12 h-12 rounded-lg border-2 overflow-hidden transition-all cursor-pointer relative',
            currentImageIndex === index ? 'border-blue-500' : 'border-gray-200 hover:border-gray-300'
          ]"
          @click="index === 0 ? openImageModal() : setMainImage(index)"
        >
          <img
:src="image" :class="[
            'w-full h-full object-cover transition-all',
            index === 0 ? 'blur-sm' : ''
          ]" />
          <div v-if="index === 0" class="absolute inset-0 flex items-center justify-center bg-black bg-opacity-50">
            <span class="text-white text-xs font-bold">---</span>
          </div>
        </button>
      </div>
    </div>

    <!-- Product Info -->
    <div class="space-y-6">
      <div>
        <h1 class="text-2xl lg:text-3xl font-bold text-gray-900 mb-2">
          {{ product.name }}
        </h1>
        <div v-if="product.sku" class="text-sm text-gray-500">
          شناسه کالا: <span class="font-mono bg-gray-100 px-2 py-0.5 rounded">{{ product.sku }}</span>
        </div>
        <div v-if="product.category && product.category.name" class="text-sm text-gray-500 mt-1">
          دسته‌بندی: 
          <NuxtLink 
            :to="`/product-category/${product.category.slug}`"
            class="font-mono bg-gray-100 px-2 py-0.5 rounded hover:bg-blue-100 hover:text-blue-700 transition-colors cursor-pointer"
          >
            {{ product.category.name }}
          </NuxtLink>
        </div>
        <div v-if="product.brand && product.brand.name" class="text-sm text-gray-500 mt-1">
          برند: 
          <NuxtLink 
            :to="`/brand/${product.brand.slug}`"
            class="font-mono bg-gray-100 px-2 py-0.5 rounded hover:bg-blue-100 hover:text-blue-700 transition-colors cursor-pointer"
          >
            {{ product.brand.name }}
          </NuxtLink>
        </div>
      </div>

      <div class="bg-gray-50 rounded-lg px-4 py-4">
        <!-- اگر سیاست‌های نمایش قیمت اجازه ندهند -->
        <template v-if="!canShowPrice">
          <span class="text-gray-600 text-sm">{{ hiddenPriceMessage }}</span>
        </template>
        <template v-else-if="product.call_for_price">
          <span class="text-yellow-700 text-lg font-bold">برای استعلام قیمت تماس بگیرید</span>
        </template>
        <template v-else-if="product.disable_buy_button">
          <span class="text-red-600 text-lg font-bold">ناموجود</span>
        </template>
        <template v-else>
          <!-- نمایش قیمت جاری و قیمت خط‌خورده (در صورت وجود) -->
          <div class="space-y-2">
            <div class="flex items-center gap-3">
              <span class="text-2xl font-bold text-green-600">
                {{ formatPrice(currentPrice) }} تومان
              </span>
              <span v-if="comparePrice" class="text-lg text-gray-500 line-through">
                {{ formatPrice(comparePrice) }} تومان
              </span>
              <!-- برچسب واحد قیمت در صورت تنظیم -->
              <span v-if="unitLabelText" class="text-xs text-gray-600">{{ unitLabelText }}</span>
            </div>
          </div>
        </template>
      </div>

      <div v-if="canAddToCart" class="space-y-3">
        <div v-if="shippingLabel" class="text-xs text-gray-700 bg-blue-50 border border-blue-100 rounded px-3 py-2 text-right">
          {{ shippingLabel }}
        </div>
        <button
          :disabled="!isInStock && !canPreorder"
          class="w-full py-3 px-6 rounded-lg font-semibold transition-all duration-200 bg-green-600 hover:bg-green-700 disabled:bg-gray-400 text-white"
          @click="$emit('add-to-cart')"
        >
          {{ addToCartLabel }}
        </button>
        <div
          v-if="product.track_inventory && product.show_stock_to_customer && product.stock_quantity !== undefined"
          class="text-xs text-gray-600 text-center"
        >
          تعداد موجودی در انبار: {{ product.stock_quantity }}
        </div>
      </div>
    </div>
  </div>

  <!-- مودال تصویر محصول -->
  <ProductImageModal 
    :is-open="showImageModal"
    :images="availableImages"
    :current-index="currentImageIndex"
    :product-name="product.name"
    @close="closeImageModal"
  />
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import ProductImageModal from '~/components/product/ProductImageModal.vue'
import { useAdmin } from '~/composables/useAdmin'
import { useAuth } from '~/composables/useAuth'
import { toLargeImage } from '~/utils/imageUtils'

interface ProductImage {
  id: number
  product_id: number
  image_url: string
  sort_order: number
}

interface Product {
  id: number
  name: string
  price: number
  sale_price?: number
  discount_price?: number
  old_price?: number
  // تنظیمات نمایش قیمت
  show_price_to_customer?: boolean
  hide_price_until_login?: boolean
  unit_label?: string
  price_per_unit?: number
  allow_preorder?: boolean
  vip_only?: boolean
  show_discount_percent?: boolean
  image?: string
  Images?: ProductImage[]
  images?: string[]
  main_image?: string
  in_stock?: boolean
  stock_quantity: number
  sku?: string
  category?: {
    name: string
    slug: string
  }
  brand?: {
    name: string
    slug: string
  }
  show_stock_to_customer: boolean
  track_inventory?: boolean
  call_for_price?: boolean
  disable_buy_button?: boolean
  sale_start_at?: string
  sale_end_at?: string
  special_offers?: {
    price: number
    sort_order: number
  }[]
}

interface Props {
  product: Product
}

const props = defineProps<Props>()
defineEmits(['add-to-cart'])

const currentImageIndex = ref(0)
// برچسب «ارسال از انبار پیش‌فرض»
const defaultWarehouseName = ref<string>('')
const shippingEnabled = ref<boolean>(false)
// مودال تصویر
const showImageModal = ref(false)
// وضعیت علاقه‌مندی
const isFavorite = ref(false)
const shippingLabel = computed(() => {
  if (!shippingEnabled.value || !defaultWarehouseName.value) return ''
  return `ارسال از انبار ${defaultWarehouseName.value}`
})

onMounted(async () => {
  try {
    // از کوئری id در صفحه پدر استفاده می‌شود؛ اگر نبود از props.product.id استفاده کن
    const id = (props.product as Record<string, unknown>)?.id
    if (!id) return
    const inv = await $fetch<Record<string, unknown>>(`/api/product-inventories/${id}`)
    // در حال حاضر shipping_enabled را از UI ادمین می‌گیریم؛ اگر نبود، فرض را false می‌گذاریم
    shippingEnabled.value = !!inv?.shipping_enabled || false
    const wid = Number(inv?.default_warehouse_id || 0)
    if (wid > 0) {
      try {
        // از API عمومی موجودی انبارها استفاده می‌کنیم تا نیازی به مسیر ادمین نباشد
        const stocksResponse = await $fetch<Record<string, unknown> | { data: Record<string, unknown>[] }>(`/api/product-warehouse-stocks`, {
          params: { product_id: id }
        })
        const entries = (Array.isArray((stocksResponse as { data: unknown[] })?.data) ? (stocksResponse as { data: unknown[] }).data : stocksResponse) as Record<string, unknown>[]
        const w = (entries || []).find((x: Record<string, unknown>) => Number(x.warehouse_id) === wid)
        defaultWarehouseName.value = (w?.warehouse_name as string) || ''
      } catch (error) {
        console.error('خطا در دریافت نام انبار پیش‌فرض:', error)
        defaultWarehouseName.value = ''
      }
    }
    
    // بررسی وضعیت علاقه‌مندی محصول - فقط برای کاربران احراز هویت شده
    // @ts-ignore - Nuxt auto-imports
    const { isAuthenticated } = useAuthState()
    if (isAuthenticated.value) {
      try {
        const wishlistResponse = await $fetch('/api/wishlist', { credentials: 'include' })
        if ((wishlistResponse as Record<string, unknown>).success) {
          const isInWishlist = ((wishlistResponse as Record<string, unknown>).data as Record<string, unknown>[]).some((item: Record<string, unknown>) => item.id === props.product.id)
          isFavorite.value = isInWishlist
        }
      } catch (error) {
        console.error('خطا در بررسی وضعیت علاقه‌مندی:', error)
        isFavorite.value = false
      }
    } else {
      isFavorite.value = false
    }
  } catch {}
})

// Admin check
const { isAdmin } = useAdmin()

// Builds array of image URLs regardless of structure
const availableImages = computed(() => {
  if (props.product.images && props.product.images.length > 0) {
    const first = props.product.images[0]
    if (typeof first === 'string') {
      // Array of plain URLs (string)
      return (props.product.images as unknown as string[]).map(url => toLargeImage(url))
    } else if (first && (first as Record<string, unknown>).url) {
      // Array of objects having url field (Media schema)
      return (props.product.images as Record<string, unknown>[]).map(i => toLargeImage((i.url || i.image_url) as string))
    } else {
      // Array of objects with image_url
      return (props.product.images as Record<string, unknown>[]).map(i => toLargeImage((i.image_url || i.url) as string))
    }
  }
  if (props.product.Images && props.product.Images.length > 0) {
    return props.product.Images.map(img => toLargeImage(((img as Record<string, unknown>).image_url || (img as Record<string, unknown>).url) as string))
  }
  return []
})

const mainImage = computed(() => {
  if (props.product.images && props.product.images.length > 0) {
    const first = props.product.images[0]
    if (typeof first === 'string') {
      return toLargeImage(props.product.images[currentImageIndex.value] as unknown as string)
    }
    const itm = props.product.images[currentImageIndex.value] as Record<string, unknown>
    return toLargeImage((itm.url || itm.image_url) as string)
  }
  if (props.product.Images && props.product.Images.length > 0) {
    const itm = props.product.Images[currentImageIndex.value] as Record<string, unknown>
    return toLargeImage((itm.image_url || itm.url) as string)
  }
  // Fallback to main_image or image
  if (props.product.main_image) {
    return toLargeImage(props.product.main_image)
  }
  if (props.product.image) {
    return toLargeImage(props.product.image)
  }
  return '/default-product.jpg'
})

// محاسبه فعال بودن بازه زمانی فروش ویژه
const isSaleScheduleActive = computed<boolean>(() => {
  try {
    const start = (props.product as Record<string, unknown>).sale_start_at
    const end = (props.product as Record<string, unknown>).sale_end_at
    const now = Date.now()
    if (start) {
      const st = new Date(start as string).getTime()
      if (now < st) return false
    }
    if (end) {
      const ed = new Date(end as string).getTime()
      if (now > ed) return false
    }
    return true
  } catch { return true }
})

// انتخاب بهترین قیمت از special_offers اگر فعال باشد
const currentPrice = computed<number>(() => {
  // 1) اگر فروش ویژه زمان‌بندی شده فعال است و special_offers وجود دارد، پله اول را در نظر بگیر
  if (isSaleScheduleActive.value && Array.isArray((props.product as Record<string, unknown>).special_offers) && (props.product as Record<string, unknown>).special_offers.length > 0) {
    const sorted = [...((props.product as Record<string, unknown>).special_offers as Record<string, unknown>[])].sort((a: Record<string, unknown>, b: Record<string, unknown>) => ((a.sort_order as number) || 0) - ((b.sort_order as number) || 0))
    const first = sorted[0]
    if (first && (first.price as number) > 0) return first.price as number
  }
  // 2) در غیر این صورت از sale_price یا discount_price استفاده کن
  const sale = props.product.sale_price || props.product.discount_price
  if (sale && sale > 0) return sale
  // 3) fallback به price
  return props.product.price
})

// قیمتی که باید خط‌خورده نمایش داده شود
const comparePrice = computed<number>(() => {
  // اگر فروش ویژه فعال است، اول old_price را بررسی کن؛ در غیر اینصورت price خط می‌خورد
  if (props.product.sale_price || props.product.discount_price) {
    if (props.product.old_price && props.product.old_price > currentPrice.value) return props.product.old_price
    return props.product.price
  }
  // اگر فروش ویژه فعال نیست ولی old_price موجود است و با price متفاوت است، همان را خط بزن
  if (props.product.old_price && props.product.old_price !== props.product.price) return props.product.old_price
  return 0
})

const discountPercentage = computed(() => {
  const base = (props.product.old_price && props.product.old_price > currentPrice.value)
    ? props.product.old_price
    : props.product.price
  const sale = currentPrice.value
  if (base && sale && sale < base) {
    return Math.round(((base - sale) / base) * 100)
  }
  return 0
})

const formatPrice = (price: number): string => {
  return new Intl.NumberFormat('fa-IR').format(price)
}

const setMainImage = (index: number) => {
  currentImageIndex.value = index
}

// موجودی واقعی انبار
const isInStock = computed(() => {
  if (props.product.track_inventory === true) {
    return props.product.in_stock === true || props.product.stock_quantity > 0
  }
  return true
})

// نبود موجودی وقتی رهگیری موجودی فعال است
const isOutOfStock = computed<boolean>(() => {
  const qty = Number((props.product as Record<string, unknown>)?.stock_quantity ?? 0)
  const inFlag = (props.product as Record<string, unknown>)?.in_stock === true
  if (props.product.track_inventory === true) {
    return !(inFlag || qty > 0)
  }
  // حتی اگر رهگیری موجودی غیرفعال باشد، اگر یکی از شاخص‌ها عدم موجودی را نشان دهد، ناموجود تلقی کن
  if (qty <= 0 || (props.product as Record<string, unknown>)?.in_stock === false) return true
  return false
})

// آیا امکان پیش‌خرید فعال است؟
const canPreorder = computed<boolean>(() => {
  return !!props.product.allow_preorder && !isInStock.value
})

// نمایش برچسب تخفیف فقط در صورت فعال بودن گزینه مدیریتی
const showDiscountBadge = computed<boolean>(() => {
  return !!props.product.show_discount_percent && discountPercentage.value > 0
})

// برچسب واحد قیمت در صورت ارسال از بک‌اند
const unitLabelText = computed<string>(() => {
  // نمونه: "قیمت هر {{unit_label}}"
  if (props.product.unit_label) {
    return `قیمت بر حسب ${props.product.unit_label}`
  }
  return ''
})

// چک دسترسی کاربر برای VIP و ورود
const { user, isAuthenticated } = useAuth()
const isVip = computed<boolean>(() => {
  try {
    return (user.value as Record<string, unknown>)?.role?.toString?.().toLowerCase() === 'vip'
  } catch {
    return false
  }
})

// آیا اجازه نمایش قیمت داریم؟
const canShowPrice = computed<boolean>(() => {
  // اگر به‌صورت صریح گفته شده که به مشتری نمایش داده نشود
  if (props.product.show_price_to_customer === false) return false
  // اگر تا ورود باید مخفی باشد
  if (props.product.hide_price_until_login === true && !isAuthenticated.value) return false
  // اگر فقط VIP
  if (props.product.vip_only === true && !isVip.value) return false
  return true
})

// پیام جایگزین برای پنهان‌سازی قیمت
const hiddenPriceMessage = computed<string>(() => {
  if (props.product.vip_only === true && !isVip.value) return 'این محصول فقط برای اعضای VIP قابل مشاهده است'
  if (props.product.hide_price_until_login === true && !isAuthenticated.value) return 'برای مشاهده قیمت وارد حساب کاربری شوید'
  return 'قیمت برای مشتریان نمایش داده نمی‌شود'
})

// اجازه نمایش دکمه خرید
const canAddToCart = computed<boolean>(() => {
  if (!canShowPrice.value) return false
  // «تماس برای قیمت» یا «دکمه خرید غیرفعال» → کل بخش پنهان
  if (!!props.product.call_for_price || !!props.product.disable_buy_button) return false
  // فقط VIP
  if (props.product.vip_only === true && !isVip.value) return false
  // اگر موجود نیست و پیش‌خرید هم مجاز نیست → پنهان
  if (isOutOfStock.value && !canPreorder.value) return false
  return true
})

// متن دکمه خرید
const addToCartLabel = computed<string>(() => {
  if (canPreorder.value) return 'پیش‌خرید'
  return isInStock.value ? 'افزودن به سبد خرید' : 'ناموجود'
})

// تابع برای باز کردن مودال تصویر
function openImageModal() {
  showImageModal.value = true
}

// تابع برای بستن مودال تصویر
function closeImageModal() {
  showImageModal.value = false
}

// توابع علاقه‌مندی و لیست
async function toggleFavorite() {
  try {
    if (isFavorite.value) {
      // حذف از لیست علاقه‌مندی‌ها
      const response = await $fetch(`/api/wishlist/remove/${props.product.id}`, {
        method: 'DELETE',
        credentials: 'include'
      })
      
      if ((response as Record<string, unknown>).success) {
        isFavorite.value = false
        // console.log('محصول از لیست علاقه‌مندی‌ها حذف شد:', props.product.id)
      }
    } else {
      // افزودن به لیست علاقه‌مندی‌ها
      const response = await $fetch('/api/wishlist/add', {
        method: 'POST',
        credentials: 'include',
        body: { product_id: props.product.id }
      })
      
      if ((response as Record<string, unknown>).success) {
        isFavorite.value = true
        // console.log('محصول به لیست علاقه‌مندی‌ها اضافه شد:', props.product.id)
      }
    }
  } catch (error) {
    console.error('خطا در تغییر وضعیت علاقه‌مندی:', error)
  }
}

function addToList() {
  // console.log('محصول به لیست اضافه شد:', props.product.id)
  // TODO: نمایش مودال انتخاب لیست یا ارسال درخواست به سرور
}

// توابع ناوبری تصاویر
function previousImage() {
  if (currentImageIndex.value > 0) {
    currentImageIndex.value--
  } else {
    currentImageIndex.value = availableImages.value.length - 1
  }
}

function nextImage() {
  if (currentImageIndex.value < availableImages.value.length - 1) {
    currentImageIndex.value++
  } else {
    currentImageIndex.value = 0
  }
}
</script> 
