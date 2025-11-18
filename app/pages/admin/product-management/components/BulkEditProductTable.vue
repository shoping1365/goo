<template>
  <div class="overflow-x-auto">
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
              @click="bulkIncreasePrice"
              :disabled="!priceIncreasePercent || priceIncreasePercent <= 0"
              class="px-2 py-1 bg-green-600 hover:bg-green-700 disabled:bg-gray-400 text-white text-xs font-medium rounded transition-colors duration-200 flex items-center gap-1"
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
              @click="bulkDecreasePrice"
              :disabled="!priceDecreasePercent || priceDecreasePercent <= 0"
              class="px-2 py-1 bg-orange-600 hover:bg-orange-700 disabled:bg-gray-400 text-white text-xs font-medium rounded transition-colors duration-200 flex items-center gap-1"
            >
              <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 13l-5 5m0 0l-5-5m5 5V6"></path>
              </svg>
              <span class="hidden lg:inline">کاهش</span>
            </button>
          </div>
          
          <button @click="bulkCallForPrice" class="px-3 py-2 bg-blue-600 hover:bg-blue-700 text-white text-xs lg:text-sm font-medium rounded-md transition-colors duration-200 flex items-center gap-1 lg:gap-2">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z"></path>
            </svg>
            <span class="hidden sm:inline">تماس برای قیمت</span>
            <span class="sm:hidden">تماس</span>
          </button>
          
          <button @click="bulkDisableBuyButton" class="px-3 py-2 bg-red-600 hover:bg-red-700 text-white text-xs lg:text-sm font-medium rounded-md transition-colors duration-200 flex items-center gap-1 lg:gap-2">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728L5.636 5.636m12.728 12.728L18.364 5.636"></path>
            </svg>
            <span class="hidden sm:inline">دکمه خرید غیرفعال</span>
            <span class="sm:hidden">غیرفعال</span>
          </button>
          

          <button @click="clearSelection" class="px-2 py-2 bg-gray-300 hover:bg-gray-400 text-gray-700 text-xs lg:text-sm font-medium rounded-md transition-colors duration-200" title="لغو انتخاب">
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
              @change="toggleSelectAll"
              :checked="allSelected"
              class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 focus:ring-2"
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
              type="checkbox" 
              v-model="selectedProducts[product.id]"
              @change="updateSelection"
              class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 focus:ring-2"
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
              v-model="modalProductId"
              v-if="modalProductId === product.id"
              :image="{ url: getProductOriginalImage(product), name: product.name }"
            />
          </td>
          <td v-if="visibleColumns.includes('name')" class="px-3 py-2 border-b border-gray-200 text-xs font-normal cursor-pointer" @click="startEditing(product.id, 'name')">
            <span v-if="!isEditing[product.id]?.name">{{ editedProducts[product.id]?.name || product.name }}</span>
            <input v-else type="text" v-model="editedProducts[product.id].name" :class="['w-full px-2 py-1 border border-gray-300 rounded text-xs font-normal', isEditing[product.id]?.name ? 'bg-white' : '']" @blur="stopEditing(product.id, 'name')" @keyup.enter="saveEdit(product.id, 'name')">
          </td>
          <td v-if="visibleColumns.includes('sku')" class="px-3 py-2 border-b border-gray-200 text-xs font-normal">{{ product.sku }}</td>
          <td v-if="visibleColumns.includes('price')" class="px-3 py-2 border-b border-gray-200 text-xs font-normal cursor-pointer" @click="startEditing(product.id, 'price')">
            <div v-if="!isEditing[product.id]?.price">
              <span>{{ editedProducts[product.id]?.price || product.price }}</span>
            </div>
            <input v-else type="number" v-model="editedProducts[product.id].price" :class="['w-24 px-2 py-1 border border-gray-300 rounded text-xs font-normal', isEditing[product.id]?.price ? 'bg-white' : '']" @blur="stopEditing(product.id, 'price')" @keyup.enter="saveEdit(product.id, 'price')">
          </td>
          <td v-if="visibleColumns.includes('old_price')" class="px-3 py-2 border-b border-gray-200 text-xs font-normal cursor-pointer" @click="startEditing(product.id, 'old_price')">
              <span v-if="!isEditing[product.id]?.old_price">{{ editedProducts[product.id]?.old_price || 0 }}</span>
              <input v-else type="number" v-model="editedProducts[product.id].old_price" :class="['w-24 px-2 py-1 border border-gray-300 rounded text-xs font-normal', isEditing[product.id]?.old_price ? 'bg-white' : '']" @blur="stopEditing(product.id, 'old_price')" @keyup.enter="saveEdit(product.id, 'old_price')">
          </td>
          <td v-if="visibleColumns.includes('cost')" class="px-3 py-2 border-b border-gray-200 text-xs font-normal cursor-pointer" @click="startEditing(product.id, 'cost')">
              <span v-if="!isEditing[product.id]?.cost">{{ editedProducts[product.id]?.cost || 0 }}</span>
              <input v-else type="number" v-model="editedProducts[product.id].cost" :class="['w-24 px-2 py-1 border border-gray-300 rounded text-xs font-normal', isEditing[product.id]?.cost ? 'bg-white' : '']" @blur="stopEditing(product.id, 'cost')" @keyup.enter="saveEdit(product.id, 'cost')">
          </td>
          <td v-if="visibleColumns.includes('track_inventory')" class="px-3 py-2 border-b border-gray-200 text-xs font-normal text-center cursor-pointer" @click="toggleCheckbox(product, 'trackInventory')">
            <svg v-if="editedProducts[product.id]?.track_inventory" class="h-5 w-5 text-green-600 inline-block" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg>
            <svg v-else class="h-5 w-5 text-red-600 inline-block" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path></svg>
          </td>
          <td v-if="visibleColumns.includes('stock_quantity')" class="px-3 py-2 border-b border-gray-200 text-xs font-normal cursor-pointer" @click="startEditing(product.id, 'stock_quantity')">
              <span v-if="!isEditing[product.id]?.stock_quantity">{{ editedProducts[product.id]?.stock_quantity || 0 }}</span>
              <input v-else type="number" v-model="editedProducts[product.id].stock_quantity" :class="['w-24 px-2 py-1 border border-gray-300 rounded text-xs font-normal', isEditing[product.id]?.stock_quantity ? 'bg-white' : '']" @blur="stopEditing(product.id, 'stock_quantity')" @keyup.enter="saveEdit(product.id, 'stock_quantity')">
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

<script setup>
import { ref, reactive, onMounted, computed, watch, nextTick } from 'vue'
import ImagePreviewModal from '~/components/media/ImagePreviewModal.vue'

const props = defineProps({
  products: {
    type: Array,
    required: true,
    default: () => []
  },
  visibleColumns: {
    type: Array,
    default: () => []
  }
})

const emit = defineEmits(['stats-updated'])

const products = ref([])
const isEditing = reactive({})
const editedProducts = reactive({})
const modalProductId = ref(null)
const selectedProducts = reactive({})
const selectAllCheckbox = ref(null)
const priceIncreasePercent = ref(null)
const priceDecreasePercent = ref(null)
const hasChanges = ref(false)

// متابع اضافی برای دریافت داده‌ها از API با sync
const fetchProducts = async () => {
  // این تابع برای sync کردن editedProducts با محصولات جدید است
  const allProducts = props.products || []
  
  // ایجاد نسخه‌ی editable برای هر محصول
  allProducts.forEach(p => {
    if (!editedProducts[p.id]) {
      editedProducts[p.id] = {
        name: p.name,
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
      if (!isEditing[p.id]) {
        isEditing[p.id] = {}
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

function getProductThumbnail(product) {
  // Return thumbnail version if available, otherwise fallback to original or placeholder
  function toThumbnail(url) {
    if (!url) return 'https://via.placeholder.com/50';
    const dotIdx = url.lastIndexOf('.')
    if (dotIdx === -1) return url + '_thumbnail';
    return url.slice(0, dotIdx) + '_thumbnail' + url.slice(dotIdx)
  }
  
  const imageUrl = getImageUrlFromProduct(product)
  if (imageUrl) return toThumbnail(imageUrl)
  return 'https://via.placeholder.com/50'
}

function getProductOriginalImage(product) {
  const imageUrl = getImageUrlFromProduct(product)
  if (imageUrl) return imageUrl
  return 'https://via.placeholder.com/400'
}

function getImageUrlFromProduct(product) {
  if (!product) return null
  
  // سعی مختلف راه‌های ممکن برای یافتن URL تصویر
  
  // راه 1: product.images[0].image_url (ساختار معمول API)
  if (product.images && Array.isArray(product.images) && product.images.length > 0) {
    if (product.images[0].image_url) return product.images[0].image_url
    if (product.images[0].url) return product.images[0].url
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
    if (product.primary_image.url) return product.primary_image.url
    if (product.primary_image.image_url) return product.primary_image.image_url
  }
  
  return null
}

function openModal(productId) {
  modalProductId.value = productId
}

function emitStats() {
  const selected = selectedCount.value
  // تعداد محصولاتی که تغییر کرده‌اند (چه انتخاب شده باشند چه نه)
  const changed = (props.products || []).filter(p => {
    const orig = p
    const edited = editedProducts[p.id]
    if (!edited) return false
    // بررسی تغییرات فیلد به فیلد
    const hasChanges = (
      edited.name !== orig.name ||
      (edited.price || 0) !== (orig.price || 0) ||
      (edited.old_price || 0) !== (orig.old_price || 0) ||
      (edited.cost || 0) !== (orig.cost || 0) ||
      (edited.stock_quantity || 0) !== (orig.stock_quantity || 0) ||
      !!edited.track_inventory !== !!orig.track_inventory ||
      !!edited.isPublished !== (orig.status === 'active' || orig.status === 'published') ||
      !!edited.disableBuyButton !== !!orig.disable_buy_button ||
      !!edited.callForPrice !== !!orig.call_for_price
    )
    
    if (hasChanges) {
      // Product changed
    }
    
    return hasChanges
  }).length
  
  emit('stats-updated', {
    selected,
    changed,
    saved: 0,
    errors: 0
  })
}

function startEditing(productId, field) {
  const product = (props.products || []).find(p => p.id === productId)
  if (product && !editedProducts[productId]) {
    editedProducts[productId] = { ...product }
    editedProducts[productId].isPublished = product.status === 'published' || product.status === 'active'
    editedProducts[productId].disableBuyButton = !!product.disable_buy_button
    editedProducts[productId].callForPrice = !!product.call_for_price
    editedProducts[productId].track_inventory = !!product.track_inventory
  }
  // Close all other editors
  Object.keys(isEditing).forEach(id => {
    Object.keys(isEditing[id] || {}).forEach(f => {
      isEditing[id][f] = false
    })
  })
  if (!isEditing[productId]) isEditing[productId] = {}
  isEditing[productId][field] = true
}

function stopEditing(productId, field) {
  if (isEditing[productId]?.[field]) {
    isEditing[productId][field] = false
    saveEdit(productId, field)
  }
}

function saveEdit(productId, field) {
  // تغییرات در editedProducts ذخیره می‌شود
  emitStats()
}

function toggleCheckbox(product, field) {
  if (!editedProducts[product.id]) {
    editedProducts[product.id] = { ...product }
    editedProducts[product.id].isPublished = product.status === 'published' || product.status === 'active'
    editedProducts[product.id].disableBuyButton = !!product.disable_buy_button
    editedProducts[product.id].callForPrice = !!product.call_for_price
    editedProducts[product.id].track_inventory = !!product.track_inventory
  }
  if (field === 'isPublished') {
    editedProducts[product.id].isPublished = !editedProducts[product.id].isPublished
  } else if (field === 'disableBuyButton') {
    editedProducts[product.id].disableBuyButton = !editedProducts[product.id].disableBuyButton
    if (editedProducts[product.id].disableBuyButton) {
      editedProducts[product.id].callForPrice = false
    }
  } else if (field === 'callForPrice') {
    editedProducts[product.id].callForPrice = !editedProducts[product.id].callForPrice
    if (editedProducts[product.id].callForPrice) {
      editedProducts[product.id].disableBuyButton = false
    }
  } else if (field === 'trackInventory') {
    editedProducts[product.id].track_inventory = !editedProducts[product.id].track_inventory
  } else {
    editedProducts[product.id][field] = !editedProducts[product.id][field]
  }
  saveEdit(product.id, field)
}

function toggleSelectAll() {
  const newState = !allSelected.value
  if (props.products && Array.isArray(props.products)) {
    props.products.forEach(product => {
      selectedProducts[product.id] = newState
    })
  }
  updateSelection()
}

function updateSelection() {
  emitStats()
}

const saveAllEdits = async () => {
  // همه محصولاتی که تغییر کرده‌اند (چه انتخاب شده باشند چه نه)
  const allProducts = props.products || []
  const changed = allProducts.filter(p => {
    const edited = editedProducts[p.id]
    if (!edited) return false
    // بررسی تغییرات فیلد به فیلد
    const hasChanges = (
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
    
    if (hasChanges) {
      // Product changed
    }
    
    return hasChanges
  })
  
  if (changed.length === 0) {
    return
  }
  const tasks = []
  changed.forEach(orig => {
    const edited = editedProducts[orig.id]
    
    const productPayload = {}
    if (!!edited.isPublished !== (orig.status === 'active' || orig.status === 'published')) {
      productPayload.status = edited.isPublished ? 'active' : 'draft'
    }
    if (edited.name !== orig.name) productPayload.name = edited.name
    if (!!edited.disableBuyButton !== !!orig.disable_buy_button) productPayload.disable_buy_button = !!edited.disableBuyButton
    if (!!edited.callForPrice !== !!orig.call_for_price) productPayload.call_for_price = !!edited.callForPrice
    if (Object.keys(productPayload).length > 0) {
      delete productPayload.images
      tasks.push($fetch(`/api/products/${orig.id}`, { method: 'PUT', body: productPayload }))
    }
    
    const pricingPayload = {}
    if ((edited.price || 0) !== (orig.price || 0)) pricingPayload.price = edited.price || 0
    if ((edited.old_price || 0) !== (orig.old_price || 0)) pricingPayload.old_price = edited.old_price || 0
    if ((edited.cost || 0) !== (orig.cost || 0)) pricingPayload.cost = edited.cost || 0
    if (Object.keys(pricingPayload).length > 0) {
      delete pricingPayload.images
      delete pricingPayload.id
      delete pricingPayload.created_at
      delete pricingPayload.updated_at
      tasks.push($fetch(`/api/product-prices/${orig.id}`, { method: 'PUT', body: pricingPayload }))
    }
    
    const inventoryPayload = {}
    if ((edited.stock_quantity || 0) !== (orig.stock_quantity || 0)) inventoryPayload.stock_quantity = edited.stock_quantity || 0
    if (!!edited.track_inventory !== !!orig.track_inventory) inventoryPayload.track_inventory = !!edited.track_inventory
    if (Object.keys(inventoryPayload).length > 0) {
      tasks.push($fetch(`/api/product-inventories/${orig.id}`, { method: 'PUT', body: inventoryPayload }))
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
  const allProducts = props.products || []
  allProducts.forEach(p => {
    if (editedProducts[p.id]) {
      editedProducts[p.id] = { ...p }
      editedProducts[p.id].isPublished = p.status === 'published' || p.status === 'active'
      editedProducts[p.id].disableBuyButton = !!p.disable_buy_button
      editedProducts[p.id].callForPrice = !!p.call_for_price
      editedProducts[p.id].track_inventory = !!p.track_inventory
    }
  })
  // بازنشانی selection state
  Object.keys(selectedProducts).forEach(id => {
    selectedProducts[id] = false
  })
  // Close all editing cells
  Object.keys(isEditing).forEach(id => {
    Object.keys(isEditing[id] || {}).forEach(f => {
      isEditing[id][f] = false
    })
  })
  hasChanges.value = false // Reset changes flag
  priceIncreasePercent.value = null // Reset price inputs
  priceDecreasePercent.value = null
  emitStats() // بروزرسانی آمار پس از بازنشانی
}

const bulkDisableBuyButton = () => {
  const selectedProductIds = Object.keys(selectedProducts).filter(id => selectedProducts[id])
  selectedProductIds.forEach(id => {
    editedProducts[id].disableBuyButton = true
    editedProducts[id].callForPrice = false // mutual exclusion
  })
  hasChanges.value = true
  emitStats()
}

const bulkCallForPrice = () => {
  const selectedProductIds = Object.keys(selectedProducts).filter(id => selectedProducts[id])
  selectedProductIds.forEach(id => {
    editedProducts[id].callForPrice = true
    editedProducts[id].disableBuyButton = false // mutual exclusion
  })
  hasChanges.value = true
  emitStats()
}

const bulkClearFlags = () => {
  const selectedProductIds = Object.keys(selectedProducts).filter(id => selectedProducts[id])
  selectedProductIds.forEach(id => {
    editedProducts[id].disableBuyButton = false
    editedProducts[id].callForPrice = false
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
    
  } catch (error) {
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
    
  } catch (error) {
    // Show error message (you can add a toast notification here)
  }
}

defineExpose({ saveAllEdits, resetEdits })

// Watch for products changes and update editedProducts
watch(() => props.products, () => {
  fetchProducts()
}, { deep: true, immediate: true })

onMounted(fetchProducts)
</script>

<style scoped>
/* استایل‌های خاص جدول در صورت نیاز */
input:focus, select:focus {
  background: #fff !important;
}
</style> 
