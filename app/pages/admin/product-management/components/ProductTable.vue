<template>
  <div class="products-table-wrapper">
    <!-- Table Content -->
    <div class="overflow-hidden">
      <!-- نوار عملیات گروهی -->
      <div v-if="selectedProducts.length > 0" class="px-4 py-2 bg-blue-50 border-b border-blue-100 flex items-center justify-between">
        <span class="text-xs text-blue-800">{{ selectedProducts.length }} مورد انتخاب شده</span>
        <div class="flex items-center gap-2">
          <button
            v-if="hasPermission('products_delete')"
            class="px-3 py-1.5 text-xs rounded-md bg-red-600 text-white hover:bg-red-700"
            @click="openBulkDeleteConfirm"
          >
            حذف انتخابشدهها
          </button>
          <button class="px-3 py-1.5 text-xs rounded-md bg-gray-100 text-gray-700 hover:bg-gray-200" @click="clearSelection">لغو انتخاب</button>
        </div>
      </div>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-300 table-fixed">
        <!-- Header جدول -->
        <thead class="bg-gray-50">
          <tr>
            <th scope="col" class="w-12 px-3 py-4 text-right text-xs font-semibold text-gray-700 uppercase tracking-wider">
              <input 
                type="checkbox" 
                :checked="selectAll"
                class="h-3 w-3 text-blue-600 focus:ring-blue-900 border-gray-300 rounded"
                @change="toggleSelectAll" 
              />
            </th>
            <th scope="col" class="w-16 px-3 py-4 text-center text-xs font-semibold text-gray-700 uppercase tracking-wider">ردیف</th>
            <th scope="col" class="w-20 px-3 py-4 text-center text-xs font-semibold text-gray-700 uppercase tracking-wider">تصویر</th>
            <th scope="col" class="px-4 py-4 text-right text-xs font-semibold text-gray-700 uppercase tracking-wider">نام محصول</th>
            <th scope="col" class="w-40 px-6 py-4 text-center text-xs font-semibold text-gray-700 uppercase tracking-wider">دسته‌بندی</th>
            <th scope="col" class="w-32 px-6 py-4 text-center text-xs font-semibold text-gray-700 uppercase tracking-wider">قیمت</th>
            <th scope="col" class="w-28 px-6 py-4 text-center text-xs font-semibold text-gray-700 uppercase tracking-wider">موجودی</th>
            <th scope="col" class="w-24 px-3 py-4 text-center text-xs font-semibold text-gray-700 uppercase tracking-wider">وضعیت</th>
            <th scope="col" class="w-36 px-3 py-4 text-center text-xs font-semibold text-gray-700 uppercase tracking-wider">عملیات</th>
          </tr>
        </thead>

        <tbody class="bg-white divide-y divide-gray-200">
          <!-- Loading State -->
          <tr v-if="isLoading">
            <td :colspan="9" class="px-4 py-8 text-center text-gray-500">
              <div class="flex items-center justify-center">
                <div class="animate-spin rounded-full h-6 w-6 border-b-2 border-blue-600 ml-2"></div>
                در حال بارگذاری...
              </div>
            </td>
          </tr>

          <!-- محصولات واقعی از API -->
          <tr v-for="(product, index) in products" v-else-if="products.length > 0" :key="product.id" class="hover:bg-gray-50 transition-colors">
            <!-- انتخاب -->
            <td class="w-12 px-3 py-4 whitespace-nowrap text-right">
              <input 
                v-model="selectedProducts" 
                type="checkbox" 
                :value="product.id"
                class="h-3 w-3 text-blue-600 focus:ring-blue-500 border-gray-300 rounded" 
              />
            </td>
            <!-- ردیف -->
            <td class="w-16 px-3 py-4 text-center">{{ (currentPage - 1) * perPage + index + 1 }}</td>

            <!-- تصویر -->
            <td class="w-20 px-3 py-4 text-center pl-4">
              <div class="flex justify-center">
                <img 
                  v-if="getProductThumbnail(product)"
                  :src="getProductThumbnail(product)"
                  :alt="product.name"
                  class="w-12 h-12 object-cover rounded-md cursor-pointer border border-gray-200 shadow-sm hover:shadow-md transition-shadow"
                  loading="lazy"
                  decoding="async"
                  @click="openModal(product.id)"
                  @error="handleImageError"
                />
                <img 
                  v-else 
                  src="/statics/images/default-image_100.png" 
                  alt="بدون تصویر" 
                  class="w-12 h-12 object-cover rounded-md opacity-60 border border-gray-200"
                />
                <ImagePreviewModal
                  v-if="modalProductId === product.id"
                  v-model="modalProductId"
                  :image="{ url: getProductOriginalImage(product), name: product.name }"
                />
              </div>
            </td>

            <!-- نام محصول -->
            <td class="px-4 py-4 text-right">
              <div class="min-w-0">
                <NuxtLink 
                  :to="productLink(product)"
                  class="text-gray-900 hover:text-blue-600 font-medium block truncate"
                  :title="product.name"
                >
                  {{ product.name }}
                </NuxtLink>
                <div class="text-sm text-gray-500 truncate mt-1">کد: {{ product.sku || product.id }}</div>
              </div>
            </td>

            <!-- دسته‌بندی -->
            <td class="w-40 px-6 py-4 text-center">
              <NuxtLink 
                v-if="product.category" 
                :to="buildCategoryLink(product.category)"
                target="_blank"
                class="inline-block px-2 py-1 bg-blue-100 text-blue-800 rounded-md text-xs font-medium truncate max-w-full hover:bg-blue-200 hover:text-blue-900 transition-colors cursor-pointer" 
                :title="`مشاهده دسته‌بندی: ${product.category.name}`"
              >
                {{ product.category.name }}
              </NuxtLink>
              <span v-else class="text-gray-400 text-xs">بدون دسته</span>
            </td>

            <!-- قیمت -->
            <td class="w-32 px-6 py-4 text-center">
              <span class="text-sm font-medium">{{ formatPrice(product.price) }}</span>
            </td>

            <!-- موجودی -->
            <td class="w-28 px-6 py-4 text-center">
              <span 
                :class="{
                  'bg-green-100 text-green-800': product.stock_quantity > 10,
                  'bg-yellow-100 text-yellow-800': product.stock_quantity <= 10 && product.stock_quantity > 0,
                  'bg-red-100 text-red-800': product.stock_quantity === 0
                }"
                class="inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium"
              >
                {{ product.stock_quantity || 0 }}
              </span>
            </td>

            <!-- وضعیت -->
            <td class="w-24 px-3 py-4 text-center">
              <span 
                :class="{
                  'bg-green-100 text-green-800': product.status === 'active',
                  'bg-red-100 text-red-800': product.status !== 'active'
                }"
                class="inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium"
              >
                {{ product.status === 'active' ? 'منتشر' : 'غیرفعال' }}
              </span>
            </td>

            <!-- عملیات -->
            <td class="w-36 px-3 py-4 text-center">
              <div class="flex justify-center gap-2">
                <NuxtLink 
                  v-if="hasPermission('product.update')"
                  :to="`/admin/product-management/products/edit?id=${product.id}`"
                  class="inline-flex items-center px-3 py-1.5 border border-blue-300 text-xs font-medium rounded text-blue-700 bg-blue-50 hover:bg-blue-100 transition-colors"
                >
                  ویرایش
                </NuxtLink>
                <button 
                  v-if="hasPermission('products_delete')"
                  class="inline-flex items-center px-3 py-1.5 border border-red-300 text-xs font-medium rounded text-red-700 bg-red-50 hover:bg-red-100 transition-colors"
                  @click="openDeleteConfirm(product.id)"
                >
                  حذف
                </button>
              </div>
            </td>
          </tr>

          <!-- Empty State -->
          <tr v-else>
            <td :colspan="9" class="px-4 py-12 text-center text-gray-500">
              <div class="text-center">
                <svg class="w-12 h-12 text-gray-300 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2 2v-5m16 0h-2M4 13h2"></path>
                </svg>
                <h3 class="text-lg font-medium text-gray-900 mb-2">هنوز محصولی اضافه نشده</h3>
                <p class="text-gray-600 mb-4">اولین محصول خود را اضافه کنید</p>
                <NuxtLink 
                  v-if="hasPermission('product.create')"
                  to="/admin/product-management/products/new"
                  class="inline-flex items-center px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-md hover:bg-blue-700"
                >
                  افزودن محصول
                </NuxtLink>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>

    <!-- Pagination -->
    <Pagination 
      v-if="products.length > 0"
      :current-page="currentPage"
      :total-pages="totalPages"
      :total="totalItems"
      :per-page="perPage"
      @page-changed="handlePageChange"
      @per-page-changed="handlePerPageChange"
    />

    <!-- مودال تایید حذف محصول -->
    <div v-if="deleteModalOpen" class="fixed inset-0 bg-transparent z-50 flex items-center justify-center">
      <div role="dialog" aria-modal="true" class="bg-blue-50 rounded-xl shadow-2xl ring-4 ring-purple-300 p-6 sm:p-7 w-full max-w-md mx-4">
        <h3 class="text-base font-semibold text-gray-900 mb-2 flex items-center gap-2">
          <span class="i-heroicons-exclamation-triangle text-red-500 text-xl"></span>
          تایید حذف محصول
        </h3>
        <p class="text-sm leading-6 text-gray-600 mb-6">آیا از حذف این محصول اطمینان دارید؟ این عملیات غیرقابل بازگشت است.</p>
        <div class="flex justify-end gap-2">
          <button class="px-4 py-2 text-sm rounded-md bg-gray-100 text-gray-700 hover:bg-gray-200" @click="closeDeleteConfirm">انصراف</button>
          <button :disabled="deleting" class="px-4 py-2 text-sm rounded-md bg-red-600 text-white hover:bg-red-700 disabled:opacity-60 disabled:cursor-not-allowed" @click="confirmDelete">
            <span v-if="deleting" class="i-heroicons-arrow-path animate-spin mr-1"></span>
            حذف
          </button>
        </div>
      </div>
    </div>

    <!-- مودال تایید حذف گروهی -->
    <div v-if="bulkDeleteModalOpen" class="fixed inset-0 bg-transparent z-50 flex items-center justify-center">
      <div role="dialog" aria-modal="true" class="bg-blue-50 rounded-xl shadow-2xl ring-4 ring-purple-300 p-6 sm:p-7 w-full max-w-md mx-4">
        <h3 class="text-base font-semibold text-gray-900 mb-2 flex items-center gap-2">
          <span class="i-heroicons-exclamation-triangle text-red-500 text-xl"></span>
          حذف گروهی محصولات
        </h3>
        <p class="text-sm leading-6 text-gray-600 mb-6">آیا از حذف {{ selectedProducts.length }} محصول انتخابشده اطمینان دارید؟ این عملیات غیرقابل بازگشت است.</p>
        <div class="flex justify-end gap-2">
          <button class="px-4 py-2 text-sm rounded-md bg-gray-100 text-gray-700 hover:bg-gray-200" @click="closeBulkDeleteConfirm">انصراف</button>
          <button :disabled="deletingBulk" class="px-4 py-2 text-sm rounded-md bg-red-600 text-white hover:bg-red-700 disabled:opacity-60 disabled:cursor-not-allowed" @click="confirmBulkDelete">
            <span v-if="deletingBulk" class="i-heroicons-arrow-path animate-spin mr-1"></span>
            حذف
          </button>
        </div>
      </div>
    </div>

  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue';
import Pagination from '~/components/admin/common/Pagination.vue';
import ImagePreviewModal from '~/components/media/ImagePreviewModal.vue';
import { useProductLink } from '~/composables/useProductLink';

// Import useAuth for permission checking
// Auth disabled
const props = defineProps<{ 
  searchTerm: string;
  filters: {
    productName?: string;
    categoryId?: string;
    brandId?: string;
    published?: string;
    stockStatus?: string;
    productId?: string;
  }
}>()

const emit = defineEmits(['stats-updated'])

// استفاده از سیستم پرمیژن محلی
 
const hasPermission = (_permission: string) => true
const { buildProductLink: productLink } = useProductLink()

// تابع برای ساخت لینک دسته‌بندی
function buildCategoryLink(category) {
  if (!category || !category.slug) {
    return '#'
  }
  return `/product-category/${category.slug}`
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const products = ref<any[]>([])
const isLoading = ref(false)
const modalProductId = ref(null)

// Pagination state
const currentPage = ref(1)
const perPage = ref(10)
const totalItems = ref(0)
const totalPages = computed(() => Math.ceil(totalItems.value / perPage.value))

const selectedProducts = ref<number[]>([])
const selectAll = computed({
  get() {
    return products.value.length > 0 && selectedProducts.value.length === products.value.length
  },
  set(value) {
    if (value) {
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      selectedProducts.value = products.value.map((p: any) => p.id as number)
    } else {
      selectedProducts.value = []
    }
  }
})

function toggleSelectAll() {
  selectAll.value = !selectAll.value
}

function getProductThumbnail(product) {
  // تابع برای ایجاد thumbnail با فرمت صحیح و بهینه‌سازی
  function toThumbnail(url) {
    if (!url) return null;
    
    // تشخیص فرمت اصلی
    const dotIdx = url.lastIndexOf('.');
    if (dotIdx === -1) return url + '_thumbnail.webp';
    
    const baseName = url.slice(0, dotIdx);
    const originalExt = url.slice(dotIdx);
    
    // اگر فرمت اصلی jpg یا jpeg است، thumbnail را با webp بساز
    if (originalExt.toLowerCase() === '.jpg' || originalExt.toLowerCase() === '.jpeg') {
      return baseName + '_thumbnail.webp';
    }
    
    // برای سایر فرمت‌ها، همان فرمت را حفظ کن
    return baseName + '_thumbnail' + originalExt;
  }
  
  // بررسی تصاویر محصول
  if (product.images && product.images.length > 0) {
    const firstImage = product.images[0];
    
    // اگر thumbnail مستقیماً موجود است
    if (firstImage.thumbnail_url) {
      return firstImage.thumbnail_url;
    }
    
    // اگر image_url موجود است، thumbnail را بساز
    if (firstImage.image_url) {
      return toThumbnail(firstImage.image_url);
    }
  }
  
  // اگر تصویر اصلی محصول موجود است
  if (product.image) {
    return toThumbnail(product.image);
  }
  
  return null;
}

function getProductOriginalImage(product) {
  if (product.images && product.images.length > 0 && product.images[0].image_url) return product.images[0].image_url;
  return null;
}

// دریافت لیست محصولات از API - صفحه‌بندی سرور
const fetchProducts = async () => {
  try {
    isLoading.value = true
    
    // ساخت URL با page و limit
    const params = new URLSearchParams()
    params.append('page', currentPage.value.toString())
    params.append('limit', perPage.value.toString())
    
    // افزودن search parameter اگر موجود باشد
    if (props.searchTerm && props.searchTerm.trim()) {
      params.append('search', props.searchTerm.trim())
    }
    
    // افزودن filter parameters اگر موجود باشند
    if (props.filters && Object.keys(props.filters).length > 0) {
      Object.entries(props.filters).forEach(([key, value]) => {
        if (value !== null && value !== undefined && value !== '') {
          params.append(`filter_${key}`, String(value))
        }
      })
    }
    
    // درخواست محصولات با pagination سرور
    const url = `/api/admin/products?${params.toString()}`
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const response: any = await $fetch(url)
    
    // بررسی فرمت پاسخ
    if (response && response.data && Array.isArray(response.data)) {
      products.value = response.data
      totalItems.value = response.total || 0
    } else if (Array.isArray(response)) {
      // فرمت قدیمی (برای سازگاری)
      products.value = response
      totalItems.value = response.length
    } else {
      products.value = []
      totalItems.value = 0
    }
    
    // آپدیت کردن آمار
    updateStats()
  } catch {
    products.value = []
    totalItems.value = 0
  } finally {
    isLoading.value = false
  }
}

// Enhanced filtering logic - حالا در سمت سرور انجام می‌شود
// اما برای محاسبه آمار محلی نگه می‌داریم
const filteredAll = computed(() => {
  const filtered = products.value

  // این فیلترها فقط برای نمایش محلی استفاده می‌شوند
  // فیلتر اصلی در سمت سرور انجام می‌شود
  
  return filtered
})

const updateStats = () => {
  // آمار را از محصولات صفحه جاری محاسبه می‌کنیم
  const total = products.value.length
  const inStock = products.value.filter(p => (parseInt(p.stock_quantity) || 0) > 10).length
  const lowStock = products.value.filter(p => {
    const stock = parseInt(p.stock_quantity) || 0
    return stock > 0 && stock <= 10
  }).length
  const outStock = products.value.filter(p => (parseInt(p.stock_quantity) || 0) === 0).length
  
  emit('stats-updated', { total, inStock, lowStock, outStock })
}

// فرمت کردن قیمت با جداکننده «،» فارسی
const formatPrice = (price) => {
  if (!price || price === 0) return '0 تومان'
  const formatted = Number(price).toLocaleString('en-US').replace(/,/g, '،')
  return formatted + ' تومان'
}

// نوتیفایر داخلی پروژه
const notifier = {
   
  success: (_message: string) => {
    // TODO: پیاده‌سازی نوتیفایر واقعی
  },
   
  error: (_message: string) => {
    // خطا را بدون چاپ کنسول مدیریت کن
  },
   
  warning: (_message: string) => {
    // TODO: پیاده‌سازی نوتیفایر واقعی
  }
}

// وضعیت تایید حذف
const deleteModalOpen = ref(false)
const deleting = ref(false)
const deleteProductId = ref<number | null>(null)
// وضعیت حذف گروهی
const bulkDeleteModalOpen = ref(false)
const deletingBulk = ref(false)

// باز کردن مودال تایید حذف
function openDeleteConfirm(productId: number) {
  deleteProductId.value = productId
  deleteModalOpen.value = true
}

// بستن مودال تایید حذف
function closeDeleteConfirm() {
  deleteModalOpen.value = false
  deleteProductId.value = null
}

// اجرای حذف محصول (بدون confirm مرورگر)
async function performDelete(productId: number) {
  try {
    // استفاده از API داخلی Nuxt که JWT را منتقل می‌کند
    await $fetch(`/api/admin/products/${productId}`, { method: 'DELETE' })
    await fetchProducts()
    notifier.success('محصول با موفقیت حذف شد')
  } catch {
    notifier.error('خطا در حذف محصول')
  }
}

// تایید حذف از طریق مودال داخلی
async function confirmDelete() {
  if (!deleteProductId.value) return
  try {
    deleting.value = true
    await performDelete(deleteProductId.value)
    closeDeleteConfirm()
  } finally {
    deleting.value = false
  }
}


// مدیریت انتخاب‌ها
function clearSelection() {
  selectedProducts.value = []
}

// باز/بستن مودال حذف گروهی
function openBulkDeleteConfirm() {
  if (selectedProducts.value.length === 0) return
  bulkDeleteModalOpen.value = true
}
function closeBulkDeleteConfirm() {
  bulkDeleteModalOpen.value = false
}

// اجرای حذف گروهی
async function confirmBulkDelete() {
  if (!selectedProducts.value.length) return
  try {
    deletingBulk.value = true
    const ids = [...selectedProducts.value]
    const results = await Promise.allSettled(ids.map((id) => $fetch(`/api/admin/products/${id}`, { method: 'DELETE' })))
    const deletedCount = results.filter(r => r.status === 'fulfilled').length
    const failedCount = results.length - deletedCount
    await fetchProducts()
    clearSelection()
    closeBulkDeleteConfirm()
    if (deletedCount > 0) notifier.success(`${deletedCount} محصول با موفقیت حذف شد`)
    if (failedCount > 0) notifier.warning(`${failedCount} حذف ناموفق بود`)
  } catch {
    notifier.error('خطا در حذف گروهی محصولات')
  } finally {
    deletingBulk.value = false
  }
}



// Pagination handlers - صفحه‌بندی سرور
const handlePageChange = (page) => {
  currentPage.value = page
  fetchProducts()
}

const handlePerPageChange = (newPerPage: number) => {
  perPage.value = Number(newPerPage)
  currentPage.value = 1
  fetchProducts()
}

function openModal(productId) {
  modalProductId.value = productId
}

function handleImageError(event) {
  // اگر تصویر لود نشد، تصویر پیش‌فرض را نمایش بده
  event.target.src = '/statics/images/default-image_100.png'
}

// Watchers for re-filtering when props change
watch(() => props.searchTerm, () => {
  currentPage.value = 1
  fetchProducts()
})

watch(() => props.filters, () => {
  currentPage.value = 1
  fetchProducts()
}, { deep: true })

// Watch products to update stats automatically
watch(() => products.value, () => {
  updateStats()
})

onMounted(() => {
  fetchProducts()
})

// متد عمومی برای دریافت داده خروجی اکسل
function getExportData() {
  if (selectedProducts.value.length > 0) {
    // فقط محصولات انتخاب شده
    return filteredAll.value.filter(p => selectedProducts.value.includes(p.id))
  } else {
    // همه محصولات فیلترشده
    return filteredAll.value
  }
}

defineExpose({ getExportData })
</script>

<style scoped>
/* استایل‌های بهبود یافته */
table {
  font-size: 0.95rem;
}

thead th {
  font-size: 0.85rem;
}

.truncate {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
</style> 

