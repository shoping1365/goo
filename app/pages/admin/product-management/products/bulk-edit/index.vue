<template>
  <ClientOnly>
    <!-- Notification -->
    <div
v-if="notification.show" :class="[
      'fixed top-6 right-4 z-50 max-w-sm w-full rounded-lg shadow-lg transition-all duration-300',
      notification.type === 'success' ? 'bg-green-50 border border-green-200' : 'bg-red-50 border border-red-200'
    ]">
      <div class="p-6">
        <div class="flex items-start">
          <div class="flex-shrink-0">
            <svg v-if="notification.type === 'success'" class="h-5 w-5 text-green-400" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
            </svg>
            <svg v-else class="h-5 w-5 text-red-400" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd" />
            </svg>
          </div>
          <div class="mr-3 w-0 flex-1">
            <p
:class="[
              'text-sm font-medium',
              notification.type === 'success' ? 'text-green-800' : 'text-red-800'
            ]">
              {{ notification.message }}
            </p>
          </div>
          <div class="mr-4 flex-shrink-0 flex">
            <button
:class="[
              'rounded-md inline-flex text-sm focus:outline-none focus:ring-2 focus:ring-offset-2',
              notification.type === 'success' ? 'text-green-500 hover:text-green-600 focus:ring-green-500' : 'text-red-500 hover:text-red-600 focus:ring-red-500'
            ]" @click="notification.show = false">
              <svg class="h-5 w-5" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd" />
              </svg>
            </button>
          </div>
        </div>
      </div>
    </div>
    <!-- Page Header -->
    <div class="bg-white shadow-sm border-b border-gray-200 mb-4">
      <div class="w-full px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center py-3">
          <div>
            <h1 class="text-lg font-bold text-gray-900">ویرایش کلی محصولات</h1>
            <p class="text-xs text-gray-600 mt-1">ویرایش و بروزرسانی انبوه محصولات</p>
          </div>
        </div>
      </div>
    </div>
    <div class="w-full px-4 py-4">
      <!-- بخش فیلتر -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden mb-6">
        <div class="bg-gradient-to-r from-slate-50 to-gray-50 px-4 py-3 border-b border-gray-200">
          <div class="flex items-center">
            <div class="w-6 h-6 bg-slate-600 rounded-md flex items-center justify-center ml-2">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.707A1 1 0 013 7V4z"></path>
              </svg>
            </div>
            <h3 class="text-sm font-semibold text-gray-900">فیلتر و جستجوی محصولات</h3>
          </div>
        </div>
        <div class="p-6">
         <ProductFilters />
         </div>
    </div>
      <!-- جدول ویرایش کلی محصولات -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
        <div class="bg-gradient-to-r from-slate-50 to-gray-50 px-4 py-3 border-b border-gray-200">
          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-6 h-6 bg-slate-600 rounded-md flex items-center justify-center ml-2 cursor-pointer" @click="showColumnsModal = true">
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <circle cx="12" cy="12" r="10" stroke="currentColor" stroke-width="2" fill="none" />
                  <path d="M8 12h8M8 16h8M8 8h8" stroke="currentColor" stroke-width="2" stroke-linecap="round" />
                </svg>
              </div>
              <h3 class="text-sm font-semibold text-gray-900">جدول ویرایش انبوه</h3>
            </div>
            <!-- Action buttons -->
            <div class="flex gap-2">
              <button v-if="stats.changed > 0" class="inline-flex items-center px-4 py-1.5 rounded text-xs font-medium text-gray-800 bg-gradient-to-r from-gray-200 to-gray-400 hover:from-gray-300 hover:to-gray-500 shadow transition" @click="bulkEditTableRef.resetEdits()">
                لغو ویرایش
              </button>
              <button 
                v-if="stats.changed > 0" 
                :disabled="saving" 
                class="inline-flex items-center px-4 py-1.5 rounded text-xs font-medium text-white bg-gradient-to-r from-blue-400 to-blue-600 hover:from-blue-500 hover:to-blue-700 disabled:opacity-50 disabled:cursor-not-allowed shadow transition"
                @click="saveChanges"
              >
                <svg v-if="saving" class="animate-spin -mr-1 ml-2 h-3 w-3 text-white" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
                ذخیره تغییرات
              </button>
            </div>
          </div>
        </div>
        <div class="p-6">
          <!-- Loading State -->
          <div v-if="loading" class="flex justify-center items-center py-12">
            <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
            <span class="mr-2 text-gray-600">در حال بارگذاری محصولات...</span>
          </div>
          <!-- Products Table -->
          <BulkEditProductTable 
            ref="bulkEditTableRef" 
            :products="products"
            :visible-columns="productTableVisibleColumns"
            @stats-updated="updateStats" 
            @products-updated="fetchProducts"
          />
          
          <!-- Pagination Component -->
          <Pagination 
            :current-page="currentPage"
            :total-pages="totalPages"
            :total="totalProducts"
            :items-per-page="itemsPerPage"
            @page-changed="handlePageChange"
            @items-per-page-changed="handleItemsPerPageChange"
          />
        </div>
      </div>
    </div>
    <!-- Columns Modal -->
    <ColumnSettingsModal
      :is-open="showColumnsModal"
      :visible-columns="productTableVisibleColumns"
      @close="showColumnsModal = false"
      @update:visible-columns="handleColumnUpdate"
    />
  </ClientOnly>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
</script>

<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue';
import Pagination from '~/components/admin/common/Pagination.vue';
import BulkEditProductTable from '~/pages/admin/product-management/components/BulkEditProductTable.vue';
import ColumnSettingsModal from '~/pages/admin/product-management/components/ColumnSettingsModal.vue';
import ProductFilters from '~/pages/admin/product-management/components/ProductFilters.vue';

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
});

// استفاده از useAuth برای چک کردن پرمیژن‌ها
// const { user, hasPermission } = useAuth()

const bulkEditTableRef = ref(null)
const products = ref([])
const loading = ref(false)
const saving = ref(false)
const showColumnsModal = ref(false)

// Pagination variables
const currentPage = ref(1)
const totalPages = ref(1)
const totalProducts = ref(0)
const itemsPerPage = ref(20)

const productTableVisibleColumns = ref([
  'select',
  'index',
  'image',
  'name',
  'sku',
  'price',
  'old_price',
  'cost',
  'track_inventory',
  'stock_quantity',
  'isNew',
  'isPublished',
  'showOnHomepage',
  'disableBuyButton',
  'callForPrice',
])

function handleColumnUpdate(newColumns) {
  productTableVisibleColumns.value = newColumns
}

const notification = reactive({
  show: false,
  type: 'success',
  message: ''
})

const stats = reactive({
  selected: 0,
  changed: 0,
  saved: 0,
  errors: 0
})

// Add a function to load column settings from the backend
async function loadColumnSettings() {
  try {
    const response = await $fetch<{ value: string }>('/api/admin/admin-settings/bulk-edit-columns', { method: 'GET' })
    if (response && response.value) {
      // If a value is returned and it's not an empty string, parse it
      productTableVisibleColumns.value = JSON.parse(response.value)

    } else {
      // If no value is returned, or if it's an empty string, use default and save them
      productTableVisibleColumns.value = [
        'select',
        'index',
        'image',
        'name',
        'sku',
        'price',
        'old_price',
        'cost',
        'track_inventory',
        'stock_quantity',
        'isNew',
        'isPublished',
        'showOnHomepage',
        'disableBuyButton',
        'callForPrice',
      ]
      
      await $fetch('/api/admin/admin-settings/bulk-edit-columns', {
        method: 'PUT',
        body: {
          value: JSON.stringify(productTableVisibleColumns.value)
        }
      })
    }
  } catch {

    // Fallback to default if loading fails
    productTableVisibleColumns.value = [
      'select',
      'index',
      'image',
      'name',
      'sku',
      'price',
      'old_price',
      'cost',
      'track_inventory',
      'stock_quantity',
      'isNew',
      'isPublished',
      'showOnHomepage',
      'disableBuyButton',
      'callForPrice',
    ]

    // تلاش برای ذخیره‌ی پیش‌فرض‌ها حتی در صورت خطا (fallback)
    try {
      await $fetch('/api/admin/admin-settings/bulk-edit-columns', {
        method: 'PUT',
        body: { value: JSON.stringify(productTableVisibleColumns.value) }
      })
    } catch {}
  }
}

function updateStats() {
  // Stats are updated internally by the component
  // No parameters needed
}

function showNotification(type, message) {
  notification.type = type
  notification.message = message
  notification.show = true
  setTimeout(() => { notification.show = false }, 5000)
}

async function fetchProducts() {
  loading.value = true
  try {
    // ساخت URL با page و limit - دقیقاً مثل ProductTable
    const params = new URLSearchParams()
    params.append('page', currentPage.value.toString())
    params.append('limit', itemsPerPage.value.toString())
    
    // درخواست محصولات با pagination سرور
    const url = `/api/admin/products?${params.toString()}`
    const response = await $fetch<{ data: unknown[], total: number, total_pages: number } | unknown[]>(url)
    
    // بررسی فرمت پاسخ - دقیقاً مثل ProductTable
    if (response && 'data' in response && Array.isArray(response.data)) {
      products.value = response.data
      totalProducts.value = response.total || 0
      totalPages.value = response.total_pages || Math.ceil(totalProducts.value / itemsPerPage.value)
    } else if (Array.isArray(response)) {
      // فرمت قدیمی (برای سازگاری)
      products.value = response
      totalProducts.value = response.length
      totalPages.value = Math.ceil(totalProducts.value / itemsPerPage.value)
    } else {
      products.value = []
      totalProducts.value = 0
      totalPages.value = 1
    }
  } catch {
    showNotification('error', 'خطا در بارگذاری محصولات')
    products.value = []
    totalProducts.value = 0
    totalPages.value = 1
  } finally {
    loading.value = false
  }
}

async function saveChanges() {
  saving.value = true
  try {
    // تعداد محصولات تغییر کرده قبل از ذخیره
    const changedCount = stats.changed
    await bulkEditTableRef.value.saveAllEdits()
    await fetchProducts()
    showNotification('success', `${changedCount} محصول با موفقیت ذخیره شد`)
    Object.assign(stats, { selected: 0, changed: 0, saved: changedCount, errors: 0 })
  } catch {
    showNotification('error', 'خطا در ذخیره تغییرات. لطفاً دوباره تلاش کنید.')
    stats.errors = 1
  } finally {
    saving.value = false
  }
}

// Pagination handlers
function handlePageChange(page) {
  currentPage.value = page
  fetchProducts()
}

function handleItemsPerPageChange(perPage) {
  itemsPerPage.value = perPage
  currentPage.value = 1
  fetchProducts()
}

onMounted(async () => {
  try {
    await loadColumnSettings()
    await fetchProducts()
  } catch {
    // حتی در صورت خطا، سعی کن محصولات را بارگذاری کن
    await fetchProducts()
  }
})
</script>

<style>
/* Override all border colors inside bulk edit products page to blue */
.admin-bulk-edit-products-page *,
.admin-bulk-edit-products-page *::before,
.admin-bulk-edit-products-page *::after {
  border-color: #bfdbfe;
}

/* Maintain blue border on focus for form elements */
.admin-bulk-edit-products-page input:focus,
.admin-bulk-edit-products-page textarea:focus,
.admin-bulk-edit-products-page select:focus {
  border-color: #60a5fa;
  outline: none;
  box-shadow: 0 0 0 1px #60a5fa;
}
</style>
