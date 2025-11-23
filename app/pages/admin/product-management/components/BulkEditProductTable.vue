<template>
  <div v-if="hasAccess" class="overflow-x-auto">
    <!-- Bulk Actions Panel -->
    <div v-if="selectedCount > 0" class="mb-4 px-4 py-4 bg-blue-50 border border-blue-200 rounded-lg shadow-sm">
      <div class="flex flex-col lg:flex-row lg:items-center lg:justify-between gapx-4 py-4">
        <div class="flex items-center gap-2">
          <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
          </svg>
          <span class="text-sm font-medium text-blue-800">{{ selectedCount }} محصول انتخاب شده</span>
        </div>
        <div class="flex flex-wrap items-center gap-2 lg:gap-3">
          <!-- Price Increase -->
          <div class="flex items-center gap-1 bg-green-50 border border-green-200 rounded-md px-2 py-1">
            <input
              v-model.number="priceIncreasePercent"
              type="number"
              min="0"
              max="1000"
              placeholder="10"
              class="w-12 lg:w-16 px-2 py-1 text-xs border border-gray-300 rounded text-center focus:ring-2 focus:ring-green-500 focus:border-green-500"
            />
            <span class="text-xs text-green-700">%</span>
            <button
              :disabled="!priceIncreasePercent || priceIncreasePercent <= 0"
              class="px-2 py-1 bg-green-600 hover:bg-green-700 disabled:bg-gray-400 text-white text-xs font-medium rounded transition-colors duration-200 flex items-center gap-1"
              @click="bulkIncreasePrice"
            >
              <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 11l5-5m0 0l5 5m-5-5v12"></path>
              </svg>
              <span class="hidden lg:inline">افزایش</span>
            </button>
          </div>
          
          <!-- Price Decrease -->
          <div class="flex items-center gap-1 bg-orange-50 border border-orange-200 rounded-md px-2 py-1">
            <input
              v-model.number="priceDecreasePercent"
              type="number"
              min="0"
              max="100"
              placeholder="10"
              class="w-12 lg:w-16 px-2 py-1 text-xs border border-gray-300 rounded text-center focus:ring-2 focus:ring-orange-500 focus:border-orange-500"
            />
            <span class="text-xs text-orange-700">%</span>
            <button
              :disabled="!priceDecreasePercent || priceDecreasePercent <= 0"
              class="px-2 py-1 bg-orange-600 hover:bg-orange-700 disabled:bg-gray-400 text-white text-xs font-medium rounded transition-colors duration-200 flex items-center gap-1"
              @click="bulkDecreasePrice"
            >
              <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 13l-5 5m0 0l-5-5m5 5V6"></path>
              </svg>
              <span class="hidden lg:inline">کاهش</span>
            </button>
          </div>
          
          <button class="px-3 py-2 bg-blue-600 hover:bg-blue-700 text-white text-xs lg:text-sm font-medium rounded-md transition-colors duration-200 flex items-center gap-1 lg:gap-2" @click="bulkCallForPrice">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z"></path>
            </svg>
            <span class="hidden sm:inline">تماس برای قیمت</span>
            <span class="sm:hidden">تماس</span>
          </button>
          
          <button class="px-3 py-2 bg-red-600 hover:bg-red-700 text-white text-xs lg:text-sm font-medium rounded-md transition-colors duration-200 flex items-center gap-1 lg:gap-2" @click="bulkDisableBuyButton">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728L5.636 5.636m12.728 12.728L18.364 5.636"></path>
            </svg>
            <span class="hidden sm:inline">دکمه خرید غیرفعال</span>
            <span class="sm:hidden">غیرفعال</span>
          </button>
          

          <button class="px-2 py-2 bg-gray-300 hover:bg-gray-400 text-gray-700 text-xs lg:text-sm font-medium rounded-md transition-colors duration-200" title="لغو انتخاب" @click="clearSelection">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>
      </div>
    </div>
    <table class="min-w-full border border-gray-200 rtl:text-right">
      <thead>
        <tr>
          <th v-if="visibleColumns.includes('select')" class="px-3 py-2 bg-gray-100 border-b border-gray-200 text-gray-600 font-normal text-xs">
            <input 
              ref="selectAllCheckbox"
              type="checkbox" 
              :checked="allSelected"
              class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 focus:ring-2"
              @change="toggleSelectAll"
            />
          </th>
          <th v-if="visibleColumns.includes('index')" class="px-3 py-2 bg-gray-100 border-b border-gray-200 text-gray-600 font-normal text-xs">#</th>
          <th v-if="visibleColumns.includes('image')" class="px-3 py-2 bg-gray-100 border-b border-gray-200 text-gray-600 font-normal text-xs">تصویر</th>
          <th v-if="visibleColumns.includes('name')" class="px-3 py-2 bg-gray-100 border-b border-gray-200 text-gray-600 font-normal text-xs">نام</th>
          <th v-if="visibleColumns.includes('sku')" class="px-3 py-2 bg-gray-100 border-b border-gray-200 text-gray-600 font-normal text-xs">شناسه کالا در انبار</th>
          <th v-if="visibleColumns.includes('price')" class="px-3 py-2 bg-gray-100 border-b border-gray-200 text-gray-600 font-normal text-xs">قیمت فروش</th>
          <th v-if="visibleColumns.includes('old_price')" class="px-3 py-2 bg-gray-100 border-b border-gray-200 text-gray-600 font-normal text-xs">قیمت خط خورده</th>
          <th v-if="visibleColumns.includes('cost')" class="px-3 py-2 bg-gray-100 border-b border-gray-200 text-gray-600 font-normal text-xs">قیمت خرید</th>
          <th v-if="visibleColumns.includes('track_inventory')" class="px-3 py-2 bg-gray-100 border-b border-gray-200 text-gray-600 font-normal text-xs">مدیریت موجودی</th>
          <th v-if="visibleColumns.includes('stock_quantity')" class="px-3 py-2 bg-gray-100 border-b border-gray-200 text-gray-600 font-normal text-xs">تعداد انبار</th>
          <th v-if="visibleColumns.includes('isNew')" class="px-3 py-2 bg-gray-100 border-b border-gray-200 text-gray-600 font-normal text-xs">علامت گذاری به عنوان جدید</th>
          <th v-if="visibleColumns.includes('isPublished')" class="px-3 py-2 bg-gray-100 border-b border-gray-200 text-gray-600 font-normal text-xs">منتشر شده</th>
          <th v-if="visibleColumns.includes('showOnHomepage')" class="px-3 py-2 bg-gray-100 border-b border-gray-200 text-gray-600 font-normal text-xs">نمایش در صفحه اصلی</th>
          <th v-if="visibleColumns.includes('disableBuyButton')" class="px-3 py-2 bg-gray-100 border-b border-gray-200 text-gray-600 font-normal text-xs">دکمه خرید غیرفعال باشد - ناموجود</th>
          <th v-if="visibleColumns.includes('callForPrice')" class="px-3 py-2 bg-gray-100 border-b border-gray-200 text-gray-600 font-normal text-xs">تماس برای قیمت</th>
          <th class="px-3 py-2 bg-gray-100 border-b border-gray-200 text-gray-600 font-normal text-xs">مشاهده</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(product, idx) in (props.products || [])" :key="product.id" class="hover:bg-gray-100">
          <td v-if="visibleColumns.includes('select')" class="px-3 py-2 border-b border-gray-200 text-xs font-normal text-center">
            <input 
              v-model="selectedProducts[product.id]" 
              type="checkbox"
              class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 focus:ring-2"
              @change="updateSelection"
            />
          </td>
          <td v-if="visibleColumns.includes('index')" class="px-3 py-2 border-b border-gray-200 text-xs font-normal">{{ idx + 1 }}</td>
          <td v-if="visibleColumns.includes('image')" class="px-3 py-2 border-b border-gray-200 text-xs font-normal">
            <img 
              :src="getProductThumbnail(product)" 
              alt="Product Image" 
              class="w-12 h-12 object-cover rounded-md cursor-pointer border border-gray-200 shadow-sm hover:shadow-md transition-shadow"
              @click="openModal(product.id)"
            />
            <ImagePreviewModal
              v-if="modalProductId === product.id"
              :model-value="Boolean(modalProductId === product.id)"
              :image="{ url: getProductOriginalImage(product) || '', name: product.name || '' }"
              @update:model-value="modalProductId = null"
            />
          </td>
          <td v-if="visibleColumns.includes('name')" class="px-3 py-2 border-b border-gray-200 text-xs font-normal cursor-pointer" @click="startEditing(product.id, 'name')">
            <span v-if="!isEditing[product.id]?.name">{{ editedProducts[product.id]?.name || product.name }}</span>
            <input v-else v-model="editedProducts[product.id].name" type="text" :class="['w-full px-2 py-1 border border-gray-300 rounded text-xs font-normal', isEditing[product.id]?.name ? 'bg-white' : '']" @blur="stopEditing(product.id, 'name')" @keyup.enter="saveEdit(product.id, 'name')">
          </td>
          <td v-if="visibleColumns.includes('sku')" class="px-3 py-2 border-b border-gray-200 text-xs font-normal">{{ product.sku }}</td>
          <td v-if="visibleColumns.includes('price')" class="px-3 py-2 border-b border-gray-200 text-xs font-normal cursor-pointer" @click="startEditing(product.id, 'price')">
            <div v-if="!isEditing[product.id]?.price">
              <span>{{ editedProducts[product.id]?.price || product.price }}</span>
            </div>
            <input v-else v-model="editedProducts[product.id].price" type="number" :class="['w-24 px-2 py-1 border border-gray-300 rounded text-xs font-normal', isEditing[product.id]?.price ? 'bg-white' : '']" @blur="stopEditing(product.id, 'price')" @keyup.enter="saveEdit(product.id, 'price')">
          </td>
          <td v-if="visibleColumns.includes('old_price')" class="px-3 py-2 border-b border-gray-200 text-xs font-normal cursor-pointer" @click="startEditing(product.id, 'old_price')">
              <span v-if="!isEditing[product.id]?.old_price">{{ editedProducts[product.id]?.old_price || 0 }}</span>
              <input v-else v-model="editedProducts[product.id].old_price" type="number" :class="['w-24 px-2 py-1 border border-gray-300 rounded text-xs font-normal', isEditing[product.id]?.old_price ? 'bg-white' : '']" @blur="stopEditing(product.id, 'old_price')" @keyup.enter="saveEdit(product.id, 'old_price')">
          </td>
          <td v-if="visibleColumns.includes('cost')" class="px-3 py-2 border-b border-gray-200 text-xs font-normal cursor-pointer" @click="startEditing(product.id, 'cost')">
              <span v-if="!isEditing[product.id]?.cost">{{ editedProducts[product.id]?.cost || 0 }}</span>
              <input v-else v-model="editedProducts[product.id].cost" type="number" :class="['w-24 px-2 py-1 border border-gray-300 rounded text-xs font-normal', isEditing[product.id]?.cost ? 'bg-white' : '']" @blur="stopEditing(product.id, 'cost')" @keyup.enter="saveEdit(product.id, 'cost')">
          </td>
          <td v-if="visibleColumns.includes('track_inventory')" class="px-3 py-2 border-b border-gray-200 text-xs font-normal text-center cursor-pointer" @click="toggleCheckbox(product, 'trackInventory')">
            <svg v-if="editedProducts[product.id]?.track_inventory" class="h-5 w-5 text-green-600 inline-block" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg>
            <svg v-else class="h-5 w-5 text-red-600 inline-block" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path></svg>
          </td>
          <td v-if="visibleColumns.includes('stock_quantity')" class="px-3 py-2 border-b border-gray-200 text-xs font-normal cursor-pointer" @click="startEditing(product.id, 'stock_quantity')">
              <span v-if="!isEditing[product.id]?.stock_quantity">{{ editedProducts[product.id]?.stock_quantity || 0 }}</span>
              <input v-else v-model="editedProducts[product.id].stock_quantity" type="number" :class="['w-24 px-2 py-1 border border-gray-300 rounded text-xs font-normal', isEditing[product.id]?.stock_quantity ? 'bg-white' : '']" @blur="stopEditing(product.id, 'stock_quantity')" @keyup.enter="saveEdit(product.id, 'stock_quantity')">
          </td>
          <td v-if="visibleColumns.includes('isNew')" class="px-3 py-2 border-b border-gray-200 text-xs font-normal text-center cursor-pointer" @click="toggleCheckbox(product, 'isNew')">
               <svg v-if="editedProducts[product.id]?.isNew" class="h-5 w-5 text-blue-600 inline-block" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg>
               <input v-else type="checkbox" :checked="editedProducts[product.id]?.isNew" class="form-checkbox h-4 w-4 text-blue-600 rounded pointer-events-none">
          </td>
          <td v-if="visibleColumns.includes('isPublished')" class="px-3 py-2 border-b border-gray-200 text-xs font-normal text-center cursor-pointer" @click="toggleCheckbox(product, 'isPublished')">
               <svg v-if="editedProducts[product.id]?.isPublished" class="h-5 w-5 text-green-600 inline-block" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg>
               <svg v-else class="h-5 w-5 text-red-600 inline-block" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path></svg>
          </td>
          <td v-if="visibleColumns.includes('showOnHomepage')" class="px-3 py-2 border-b border-gray-200 text-xs font-normal text-center cursor-pointer" @click="toggleCheckbox(product, 'showOnHomepage')">
              <svg v-if="editedProducts[product.id]?.showOnHomepage" class="h-5 w-5 text-blue-600 inline-block" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg>
              <svg v-else class="h-5 w-5 text-red-600 inline-block" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path></svg>
          </td>
           <td v-if="visibleColumns.includes('disableBuyButton')" class="px-3 py-2 border-b border-gray-200 text-xs font-normal text-center cursor-pointer" @click="toggleCheckbox(product, 'disableBuyButton')">
              <svg v-if="editedProducts[product.id]?.disableBuyButton" class="h-5 w-5 text-blue-600 inline-block" fill="none" viewBox="0 0 24 24" stroke="currentColor" @click.stop="toggleCheckbox(product, 'disableBuyButton')"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg>
              <svg v-else class="h-5 w-5 text-red-600 inline-block" fill="none" viewBox="0 0 24 24" stroke="currentColor" @click.stop="toggleCheckbox(product, 'disableBuyButton')"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path></svg>
          </td>
          <td v-if="visibleColumns.includes('callForPrice')" class="px-3 py-2 border-b border-gray-200 text-xs font-normal text-center cursor-pointer" @click="toggleCheckbox(product, 'callForPrice')">
               <svg v-if="editedProducts[product.id]?.callForPrice" class="h-5 w-5 text-blue-600 inline-block" fill="none" viewBox="0 0 24 24" stroke="currentColor" @click.stop="toggleCheckbox(product, 'callForPrice')"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg>
               <svg v-else class="h-5 w-5 text-red-600 inline-block" fill="none" viewBox="0 0 24 24" stroke="currentColor" @click.stop="toggleCheckbox(product, 'callForPrice')"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path></svg>
          </td>
          <td class="px-3 py-2 border-b border-gray-200 text-xs font-normal text-center">
            <a :href="`/product/sku-${product.sku || product.id}`" target="_blank" rel="noopener">
              <svg class="h-5 w-5 text-blue-600 inline-block cursor-pointer" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
              </svg>
            </a>
          </td>
        </tr>
        <!-- Empty State -->
        <tr v-if="!products || products.length === 1 || typeof products === 'string'">
          <td :colspan="visibleColumns.length + 1" class="px-4 py-12 text-center text-gray-500">
            <div class="text-center">
              <svg class="w-12 h-12 text-gray-300 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2 2v-5m16 0h-2M4 13h2"></path>
              </svg>
              <h3 class="text-lg font-medium text-gray-900 mb-2">هیچ محصولی یافت نشد</h3>
              <p class="text-gray-600 mb-4">محصولاتی برای ویرایش وجود ندارد</p>
            </div>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
</script>

<script setup lang="ts">
import { computed, nextTick, onMounted, reactive, ref, watch } from 'vue'
import { useAuth } from '~/composables/useAuth'
import ImagePreviewModal from '~/components/media/ImagePreviewModal.vue'

// احراز هویت
const { user, isAuthenticated } = useAuth()

// بررسی دسترسی admin
const hasAccess = computed(() => {
  if (!isAuthenticated.value) {
    return false
  }

  const userRole = user.value?.role?.toLowerCase() || ''
  const adminRoles = ['admin', 'developer']
  return adminRoles.includes(userRole)
})

// بررسی احراز هویت و دسترسی admin - نمایش 404 در صورت عدم دسترسی
const checkAuth = async () => {
  if (!hasAccess.value) {
    await navigateTo('/404', { external: false })
  }
}

// بررسی احراز هویت هنگام تغییر وضعیت احراز هویت
watch([isAuthenticated, hasAccess], async () => {
  if (!hasAccess.value) {
    await checkAuth()
  }
})

// تعریف interface برای Product
interface Product {
  id: number | string
  name?: string
  price?: number
  old_price?: number
  cost?: number
  stock_quantity?: number
  track_inventory?: boolean
  status?: string
  disable_buy_button?: boolean
  call_for_price?: boolean
  isNew?: boolean
  showOnHomepage?: boolean
  sku?: string
  image?: string | { url?: string; thumbnail?: string }
  image_url?: string
  thumbnail?: string
  featured_image?: string
  primary_image?: string
  images?: Array<{ url?: string; thumbnail?: string; image_url?: string; thumbnail_url?: string }>
  [key: string]: unknown
}

// تعریف interface برای EditedProduct
interface EditedProduct {
  name?: string
  price?: number
  old_price?: number
  cost?: number
  stock_quantity?: number
  track_inventory?: boolean
  isPublished?: boolean
  disableBuyButton?: boolean
  callForPrice?: boolean
  isNew?: boolean
  showOnHomepage?: boolean
  [key: string]: unknown
}

// تعریف interface برای Payload
interface ProductPayload {
  status?: string
  name?: string
  disable_buy_button?: boolean
  call_for_price?: boolean
  [key: string]: unknown
}

interface PricingPayload {
  price?: number
  old_price?: number
  cost?: number
  [key: string]: unknown
}

interface InventoryPayload {
  stock_quantity?: number
  track_inventory?: boolean
  [key: string]: unknown
}

const props = defineProps<{
  products?: Product[]
  visibleColumns?: string[]
}>()

const emit = defineEmits<{
  (e: 'stats-updated'): void
}>()

const isEditing = reactive<Record<string | number, Record<string, boolean>>>({})
const editedProducts = reactive<Record<string | number, EditedProduct>>({})
const modalProductId = ref<number | string | null>(null)
const selectedProducts = reactive<Record<string | number, boolean>>({})
const selectAllCheckbox = ref<HTMLInputElement | null>(null)
const priceIncreasePercent = ref<number | null>(null)
const priceDecreasePercent = ref<number | null>(null)
const hasChanges = ref(false)

// متابع اضافی برای دریافت داده‌ها از API با sync
const fetchProducts = async () => {
  // این تابع برای sync کردن editedProducts با محصولات جدید است
  const allProducts = (props.products || []) as Product[]
  
  // ایجاد نسخه‌ی editable برای هر محصول
  allProducts.forEach((p: Product) => {
    const productId = p.id
    if (!editedProducts[productId]) {
      editedProducts[productId] = {
        name: p.name || '',
        price: p.price || 0,
        old_price: p.old_price || 0,
        cost: p.cost || 0,
        stock_quantity: p.stock_quantity || 0,
        track_inventory: !!p.track_inventory,
        isPublished: p.status === 'published' || p.status === 'active',
        disableBuyButton: !!p.disable_buy_button,
        callForPrice: !!p.call_for_price,
        isNew: !!p.isNew,
        showOnHomepage: !!p.showOnHomepage
      }
      
      // تهیه‌سازی isEditing state
      if (!isEditing[productId]) {
        isEditing[productId] = {}
      }
    }
  })
}

const selectedCount = computed(() => {
  return Object.values(selectedProducts).filter(Boolean).length
})
const allSelected = computed(() => {
  return props.products && props.products.length > 0 && selectedCount.value === props.products.length
})
const someSelected = computed(() => {
  return selectedCount.value > 0 && props.products && selectedCount.value < props.products.length
})

watch(someSelected, async (newVal) => {
  await nextTick()
  if (selectAllCheckbox.value) {
    selectAllCheckbox.value.indeterminate = newVal
  }
})

function getProductThumbnail(product: Product): string {
  // Return thumbnail version if available, otherwise fallback to original or placeholder
  function toThumbnail(url: string): string {
    if (!url) return 'https://via.placeholder.com/50';
    const dotIdx = url.lastIndexOf('.')
    if (dotIdx === -1) return url + '_thumbnail';
    return url.slice(0, dotIdx) + '_thumbnail' + url.slice(dotIdx)
  }
  
  const imageUrl = getImageUrlFromProduct(product)
  if (imageUrl) return toThumbnail(imageUrl)
  return 'https://via.placeholder.com/50'
}

function getProductOriginalImage(product: Product): string {
  const imageUrl = getImageUrlFromProduct(product)
  if (imageUrl) return imageUrl
  return 'https://via.placeholder.com/400'
}

function getImageUrlFromProduct(product: Product): string | null {
  if (!product) return null
  
  // سعی مختلف راه‌های ممکن برای یافتن URL تصویر
  
  // راه 1: product.images[0].image_url (ساختار معمول API)
  if (product.images && Array.isArray(product.images) && product.images.length > 0) {
    const firstImage = product.images[0]
    if (firstImage && typeof firstImage === 'object') {
      if ('image_url' in firstImage && typeof firstImage.image_url === 'string') return firstImage.image_url
      if ('url' in firstImage && typeof firstImage.url === 'string') return firstImage.url
      if ('thumbnail_url' in firstImage && typeof firstImage.thumbnail_url === 'string') return firstImage.thumbnail_url
    }
  }
  
  // راه 2: product.image (فیلد مستقیم)
  if (product.image && typeof product.image === 'string') return product.image
  
  // راه 3: product.image_url (فیلد مستقیم)
  if (product.image_url && typeof product.image_url === 'string') return product.image_url
  
  // راه 4: product.thumbnail (فیلد مستقیم)
  if (product.thumbnail && typeof product.thumbnail === 'string') return product.thumbnail
  
  // راه 5: product.featured_image (فیلد مستقیم)
  if (product.featured_image && typeof product.featured_image === 'string') return product.featured_image
  
  // راه 6: محصول شاید یک شی تصویر دارد
  if (product.primary_image) {
    if (typeof product.primary_image === 'string') return product.primary_image
    if (typeof product.primary_image === 'object' && product.primary_image !== null) {
      const primaryImage = product.primary_image as Record<string, unknown>
      if ('url' in primaryImage && typeof primaryImage.url === 'string') return primaryImage.url
      if ('image_url' in primaryImage && typeof primaryImage.image_url === 'string') return primaryImage.image_url
    }
  }
  
  return null
}

function openModal(productId: number | string): void {
  modalProductId.value = productId
}

function emitStats(): void {
  // بررسی آمار برای emit
  // تعداد محصولاتی که تغییر کرده‌اند (چه انتخاب شده باشند چه نه)
  const allProducts = (props.products || []) as Product[]
  allProducts.filter((p: Product) => {
    const productId = p.id
    const edited = editedProducts[productId]
    if (!edited) return false
    // بررسی تغییرات فیلد به فیلد
    const hasChangesValue = (
      edited.name !== p.name ||
      (edited.price || 0) !== (p.price || 0) ||
      (edited.old_price || 0) !== (p.old_price || 0) ||
      (edited.cost || 0) !== (p.cost || 0) ||
      (edited.stock_quantity || 0) !== (p.stock_quantity || 0) ||
      !!edited.track_inventory !== !!p.track_inventory ||
      !!edited.isPublished !== (p.status === 'active' || p.status === 'published') ||
      !!edited.disableBuyButton !== !!p.disable_buy_button ||
      !!edited.callForPrice !== !!p.call_for_price
    )
    
    return hasChangesValue
  })
  
  emit('stats-updated')
}

function startEditing(productId: number | string, field: string): void {
  const allProducts = (props.products || []) as Product[]
  const product = allProducts.find((p: Product) => p.id === productId)
  if (product && !editedProducts[productId]) {
    editedProducts[productId] = {
      name: product.name || '',
      price: product.price || 0,
      old_price: product.old_price || 0,
      cost: product.cost || 0,
      stock_quantity: product.stock_quantity || 0,
      track_inventory: !!product.track_inventory,
      isPublished: product.status === 'published' || product.status === 'active',
      disableBuyButton: !!product.disable_buy_button,
      callForPrice: !!product.call_for_price,
      isNew: !!product.isNew,
      showOnHomepage: !!product.showOnHomepage
    }
  }
  // Close all other editors
  Object.keys(isEditing).forEach(id => {
    if (isEditing[id]) {
      Object.keys(isEditing[id] || {}).forEach(f => {
        if (isEditing[id]) {
          isEditing[id][f] = false
        }
      })
    }
  })
  if (!isEditing[productId]) {
    isEditing[productId] = {}
  }
  if (isEditing[productId]) {
    isEditing[productId][field] = true
  }
}

function stopEditing(productId: number | string, field: string): void {
  if (isEditing[productId]?.[field]) {
    if (isEditing[productId]) {
      isEditing[productId][field] = false
    }
    saveEdit(productId, field)
  }
}

function saveEdit(_productId: number | string, _field: string): void {
  // تغییرات در editedProducts ذخیره می‌شود
  hasChanges.value = true
  emitStats()
}

function toggleCheckbox(product: Product, field: string): void {
  const productId = product.id
  if (!editedProducts[productId]) {
    editedProducts[productId] = {
      name: product.name || '',
      price: product.price || 0,
      old_price: product.old_price || 0,
      cost: product.cost || 0,
      stock_quantity: product.stock_quantity || 0,
      track_inventory: !!product.track_inventory,
      isPublished: product.status === 'published' || product.status === 'active',
      disableBuyButton: !!product.disable_buy_button,
      callForPrice: !!product.call_for_price,
      isNew: !!product.isNew,
      showOnHomepage: !!product.showOnHomepage
    }
  }
  const edited = editedProducts[productId]
  if (!edited) return
  
  if (field === 'isPublished') {
    edited.isPublished = !edited.isPublished
  } else if (field === 'disableBuyButton') {
    edited.disableBuyButton = !edited.disableBuyButton
    if (edited.disableBuyButton) {
      edited.callForPrice = false
    }
  } else if (field === 'callForPrice') {
    edited.callForPrice = !edited.callForPrice
    if (edited.callForPrice) {
      edited.disableBuyButton = false
    }
  } else if (field === 'trackInventory') {
    edited.track_inventory = !edited.track_inventory
  } else {
    (edited as Record<string, unknown>)[field] = !(edited as Record<string, unknown>)[field]
  }
  saveEdit(productId, field)
}

function toggleSelectAll(): void {
  const newState = !allSelected.value
  const allProducts = (props.products || []) as Product[]
  allProducts.forEach((product: Product) => {
    selectedProducts[product.id] = newState
  })
  updateSelection()
}

function updateSelection() {
  emitStats()
}

const saveAllEdits = async () => {
  // همه محصولاتی که تغییر کرده‌اند (چه انتخاب شده باشند چه نه)
  const allProducts = (props.products || []) as Product[]
  const changed = allProducts.filter((p: Product) => {
    const productId = p.id
    const edited = editedProducts[productId]
    if (!edited) return false
    // بررسی تغییرات فیلد به فیلد
    const hasChangesValue = (
      edited.name !== p.name ||
      (edited.price || 0) !== (p.price || 0) ||
      (edited.old_price || 0) !== (p.old_price || 0) ||
      (edited.cost || 0) !== (p.cost || 0) ||
      (edited.stock_quantity || 0) !== (p.stock_quantity || 0) ||
      !!edited.track_inventory !== !!p.track_inventory ||
      !!edited.isPublished !== (p.status === 'active' || p.status === 'published') ||
      !!edited.disableBuyButton !== !!p.disable_buy_button ||
      !!edited.callForPrice !== !!p.call_for_price
    )
    
    return hasChangesValue
  })
  
  if (changed.length === 0) {
    return
  }
  const tasks: Promise<unknown>[] = []
  changed.forEach((orig: Product) => {
    const productId = orig.id
    const edited = editedProducts[productId]
    if (!edited) return
    
    const productPayload: ProductPayload = {}
    if (!!edited.isPublished !== (orig.status === 'active' || orig.status === 'published')) {
      productPayload.status = edited.isPublished ? 'active' : 'draft'
    }
    if (edited.name !== orig.name) productPayload.name = edited.name
    if (!!edited.disableBuyButton !== !!orig.disable_buy_button) productPayload.disable_buy_button = !!edited.disableBuyButton
    if (!!edited.callForPrice !== !!orig.call_for_price) productPayload.call_for_price = !!edited.callForPrice
    if (Object.keys(productPayload).length > 0) {
      delete productPayload.images
      tasks.push($fetch(`/api/products/${productId}`, { method: 'PUT', body: productPayload }))
    }
    
    const pricingPayload: PricingPayload = {}
    if ((edited.price || 0) !== (orig.price || 0)) pricingPayload.price = edited.price || 0
    if ((edited.old_price || 0) !== (orig.old_price || 0)) pricingPayload.old_price = edited.old_price || 0
    if ((edited.cost || 0) !== (orig.cost || 0)) pricingPayload.cost = edited.cost || 0
    if (Object.keys(pricingPayload).length > 0) {
      delete pricingPayload.images
      delete pricingPayload.id
      delete pricingPayload.created_at
      delete pricingPayload.updated_at
      tasks.push($fetch(`/api/product-prices/${productId}`, { method: 'PUT', body: pricingPayload }))
    }
    
    const inventoryPayload: InventoryPayload = {}
    if ((edited.stock_quantity || 0) !== (orig.stock_quantity || 0)) inventoryPayload.stock_quantity = edited.stock_quantity || 0
    if (!!edited.track_inventory !== !!orig.track_inventory) inventoryPayload.track_inventory = !!edited.track_inventory
    if (Object.keys(inventoryPayload).length > 0) {
      tasks.push($fetch(`/api/product-inventories/${productId}`, { method: 'PUT', body: inventoryPayload }))
    }
  })
  
  if (tasks.length > 0) {
    try {
      await Promise.all(tasks)
    } catch (error) {
      throw error
    }
  }
  
  hasChanges.value = false // Reset after saving
  await fetchProducts()
}

const resetEdits = () => {
  // بازنشانی همه محصولات تغییر کرده
  const allProducts = (props.products || []) as Product[]
  allProducts.forEach((p: Product) => {
    const productId = p.id
    if (editedProducts[productId]) {
      editedProducts[productId] = {
        name: p.name || '',
        price: p.price || 0,
        old_price: p.old_price || 0,
        cost: p.cost || 0,
        stock_quantity: p.stock_quantity || 0,
        track_inventory: !!p.track_inventory,
        isPublished: p.status === 'published' || p.status === 'active',
        disableBuyButton: !!p.disable_buy_button,
        callForPrice: !!p.call_for_price,
        isNew: !!p.isNew,
        showOnHomepage: !!p.showOnHomepage
      }
    }
  })
  // بازنشانی selection state
  Object.keys(selectedProducts).forEach(id => {
    selectedProducts[id] = false
  })
  // Close all editing cells
  Object.keys(isEditing).forEach(id => {
    Object.keys(isEditing[id] || {}).forEach(f => {
      if (isEditing[id]) {
        isEditing[id][f] = false
      }
    })
  })
  hasChanges.value = false // Reset changes flag
  priceIncreasePercent.value = null // Reset price inputs
  priceDecreasePercent.value = null
  emitStats() // بروزرسانی آمار پس از بازنشانی
}

const bulkDisableBuyButton = (): void => {
  const selectedProductIds = Object.keys(selectedProducts).filter(id => selectedProducts[id])
  selectedProductIds.forEach(id => {
    const productId = id as string | number
    if (!editedProducts[productId]) {
      const allProducts = (props.products || []) as Product[]
      const product = allProducts.find((p: Product) => p.id === productId)
      if (product) {
        editedProducts[productId] = {
          name: product.name || '',
          price: product.price || 0,
          old_price: product.old_price || 0,
          cost: product.cost || 0,
          stock_quantity: product.stock_quantity || 0,
          track_inventory: !!product.track_inventory,
          isPublished: product.status === 'published' || product.status === 'active',
          disableBuyButton: !!product.disable_buy_button,
          callForPrice: !!product.call_for_price,
          isNew: !!product.isNew,
          showOnHomepage: !!product.showOnHomepage
        }
      }
    }
    if (editedProducts[productId]) {
      editedProducts[productId].disableBuyButton = true
      editedProducts[productId].callForPrice = false // mutual exclusion
    }
  })
  hasChanges.value = true
  emitStats()
}

const bulkCallForPrice = (): void => {
  const selectedProductIds = Object.keys(selectedProducts).filter(id => selectedProducts[id])
  selectedProductIds.forEach(id => {
    const productId = id as string | number
    if (!editedProducts[productId]) {
      const allProducts = (props.products || []) as Product[]
      const product = allProducts.find((p: Product) => p.id === productId)
      if (product) {
        editedProducts[productId] = {
          name: product.name || '',
          price: product.price || 0,
          old_price: product.old_price || 0,
          cost: product.cost || 0,
          stock_quantity: product.stock_quantity || 0,
          track_inventory: !!product.track_inventory,
          isPublished: product.status === 'published' || product.status === 'active',
          disableBuyButton: !!product.disable_buy_button,
          callForPrice: !!product.call_for_price,
          isNew: !!product.isNew,
          showOnHomepage: !!product.showOnHomepage
        }
      }
    }
    if (editedProducts[productId]) {
      editedProducts[productId].callForPrice = true
      editedProducts[productId].disableBuyButton = false // mutual exclusion
    }
  })
  hasChanges.value = true
  emitStats()
}

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const bulkClearFlags = (): void => {
  const selectedProductIds = Object.keys(selectedProducts).filter(id => selectedProducts[id])
  selectedProductIds.forEach(id => {
    const productId = id as string | number
    if (editedProducts[productId]) {
      editedProducts[productId].disableBuyButton = false
      editedProducts[productId].callForPrice = false
    }
  })
  hasChanges.value = true
  emitStats()
}

const clearSelection = () => {
  Object.keys(selectedProducts).forEach(id => {
    selectedProducts[id] = false
  })
  emitStats()
}

const bulkIncreasePrice = async () => {
  if (!priceIncreasePercent.value || priceIncreasePercent.value <= 0) {
    return
  }
  
  const selectedProductIds = Object.keys(selectedProducts).filter(id => selectedProducts[id])
  
  if (selectedProductIds.length === 0) return
  
  try {
    // Use bulk API for better performance
    await $fetch('/api/product-prices/bulk-update', {
      method: 'POST',
      body: {
        product_ids: selectedProductIds.map(id => parseInt(id)),
        action: 'increase',
        percent: priceIncreasePercent.value
      }
    })
    
    // Refresh products to get updated data
    await fetchProducts()
    
    // Reset inputs
    priceIncreasePercent.value = null
    hasChanges.value = false
    
  } catch {
    // Show error message (you can add a toast notification here)
  }
}

const bulkDecreasePrice = async () => {
  if (!priceDecreasePercent.value || priceDecreasePercent.value <= 0) {
    return
  }
  
  const selectedProductIds = Object.keys(selectedProducts).filter(id => selectedProducts[id])
  
  if (selectedProductIds.length === 0) return
  
  try {
    // Use bulk API for better performance
    await $fetch('/api/product-prices/bulk-update', {
      method: 'POST',
      body: {
        product_ids: selectedProductIds.map(id => parseInt(id)),
        action: 'decrease',
        percent: priceDecreasePercent.value
      }
    })
    
    // Refresh products to get updated data
    await fetchProducts()
    
    // Reset inputs
    priceDecreasePercent.value = null
    hasChanges.value = false
    
  } catch {
    // Show error message (you can add a toast notification here)
  }
}

defineExpose({ saveAllEdits, resetEdits })

// Watch for products changes and update editedProducts
watch(() => props.products, () => {
  fetchProducts()
}, { deep: true, immediate: true })

onMounted(async () => {
  await checkAuth()
  if (!hasAccess.value) {
    return
  }
  fetchProducts()
})
</script>

<style scoped>
/* استایل‌های خاص جدول در صورت نیاز */
input:focus, select:focus {
  background: #fff !important;
}
</style> 
