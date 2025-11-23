<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between">
        <div>
          <h2 class="text-xl font-bold text-gray-900">محصولات قابل خرید اقساطی</h2>
          <p class="text-sm text-gray-500 mt-1">مدیریت محصولاتی که امکان خرید اقساطی دارند</p>
        </div>
        <button
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700"
          @click="openAddProductModal"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
          </svg>
          افزودن محصول
        </button>
      </div>
    </div>

    <!-- فیلترها -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">جستجو</label>
          <input
            v-model="filters.search"
            type="text"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            placeholder="نام محصول، SKU یا برند"
          />
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">دسته‌بندی</label>
          <select
            v-model="filters.category"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            <option value="">همه دسته‌ها</option>
            <option v-for="category in categories" :key="category.id" :value="category.id">
              {{ category.name }}
            </option>
          </select>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
          <select
            v-model="filters.status"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            <option value="">همه وضعیت‌ها</option>
            <option value="active">فعال</option>
            <option value="inactive">غیرفعال</option>
            <option value="pending">در انتظار تایید</option>
          </select>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">مرتب‌سازی</label>
          <select
            v-model="filters.sort"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            <option value="name">نام محصول</option>
            <option value="price">قیمت</option>
            <option value="installment_count">تعداد اقساط</option>
            <option value="created_at">تاریخ افزودن</option>
          </select>
        </div>
      </div>
      
      <div class="mt-4 flex justify-end">
        <button
          class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500"
          @click="applyFilters"
        >
          اعمال فیلتر
        </button>
      </div>
    </div>

    <!-- جدول محصولات -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200">
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">محصول</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">قیمت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">اقساط</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">آمار فروش</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="product in filteredProducts" :key="product.id">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="">
                    <img class="h-10 w-10 rounded-lg object-cover" :src="product.image" :alt="product.name" />
                  </div>
                  <div class="mr-4">
                    <div class="text-sm font-medium text-gray-900">{{ product.name }}</div>
                    <div class="text-sm text-gray-500">{{ product.sku }}</div>
                    <div class="text-sm text-gray-500">{{ product.brand }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-900">{{ formatCurrency(product.price) }}</div>
                <div class="text-sm text-gray-500">{{ formatCurrency(product.installmentPrice) }} اقساطی</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="space-y-1">
                  <div v-for="installment in product.installments" :key="installment.count" class="text-sm">
                    <span class="text-gray-900">{{ installment.count }} قسط:</span>
                    <span class="text-gray-600">{{ formatCurrency(installment.amount) }}</span>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="px-2 py-1 text-xs font-medium rounded-full" :class="getStatusClass(product.status)">
                  {{ getStatusText(product.status) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-900">{{ product.salesCount }} فروش</div>
                <div class="text-sm text-gray-500">{{ formatCurrency(product.totalSales) }} درآمد</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button
                  class="text-blue-600 hover:text-blue-900 ml-3"
                  @click="editProduct(product)"
                >
                  ویرایش
                </button>
                <button
                  :class="product.status === 'active' ? 'text-red-600 hover:text-red-900' : 'text-green-600 hover:text-green-900'"
                  @click="toggleProductStatus(product)"
                >
                  {{ product.status === 'active' ? 'غیرفعال' : 'فعال' }}
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- مودال افزودن/ویرایش محصول -->
    <div v-if="showProductModal" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
      <div class="relative top-20 mx-auto p-5 border w-11/12 md:w-3/4 lg:w-1/2 shadow-lg rounded-md bg-white">
        <div class="mt-3">
          <h3 class="text-lg font-medium text-gray-900 mb-4">
            {{ editingProduct ? 'ویرایش محصول' : 'افزودن محصول جدید' }}
          </h3>
          
          <form class="space-y-6" @submit.prevent="saveProduct">
            <!-- اطلاعات پایه -->
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">نام محصول</label>
                <input
                  v-model="productForm.name"
                  type="text"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                  required
                />
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">SKU</label>
                <input
                  v-model="productForm.sku"
                  type="text"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                  required
                />
              </div>
            </div>

            <!-- قیمت‌ها -->
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">قیمت نقدی (تومان)</label>
                <input
                  v-model="productForm.price"
                  type="number"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                  required
                />
              </div>
              
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">قیمت اقساطی (تومان)</label>
                <input
                  v-model="productForm.installmentPrice"
                  type="number"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                  required
                />
              </div>
            </div>

            <!-- تنظیمات اقساط -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">گزینه‌های اقساط</label>
              <div class="space-y-3">
                <div v-for="(installment, index) in productForm.installments" :key="index" class="flex items-center space-x-3 space-x-reverse">
                  <input
                    v-model="installment.count"
                    type="number"
                    placeholder="تعداد اقساط"
                    class="w-24 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                  />
                  <span class="text-gray-500">قسط</span>
                  <input
                    v-model="installment.amount"
                    type="number"
                    placeholder="مبلغ هر قسط"
                    class="w-32 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                  />
                  <span class="text-gray-500">تومان</span>
                  <button
                    type="button"
                    class="text-red-600 hover:text-red-800"
                    @click="removeInstallment(index)"
                  >
                    حذف
                  </button>
                </div>
                <button
                  type="button"
                  class="text-blue-600 hover:text-blue-800 text-sm"
                  @click="addInstallment"
                >
                  + افزودن گزینه اقساط
                </button>
              </div>
            </div>

            <!-- شرایط اعتبارسنجی -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">شرایط اعتبارسنجی</label>
              <div class="space-y-3">
                <div class="flex items-center">
                  <input
                    v-model="productForm.creditCheck.required"
                    type="checkbox"
                    class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                  />
                  <label class="mr-2 text-sm text-gray-700">نیاز به اعتبارسنجی</label>
                </div>
                
                <div v-if="productForm.creditCheck.required" class="grid grid-cols-1 md:grid-cols-2 gap-6">
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-2">حداقل امتیاز اعتباری</label>
                    <input
                      v-model="productForm.creditCheck.minScore"
                      type="number"
                      min="0"
                      max="1000"
                      class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                    />
                  </div>
                  
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-2">حداقل سن</label>
                    <input
                      v-model="productForm.creditCheck.minAge"
                      type="number"
                      min="18"
                      max="80"
                      class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                    />
                  </div>
                </div>
              </div>
            </div>

            <!-- دکمه‌ها -->
            <div class="flex justify-end space-x-3 space-x-reverse">
              <button
                type="button"
                class="px-4 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400 focus:outline-none focus:ring-2 focus:ring-gray-500"
                @click="closeProductModal"
              >
                انصراف
              </button>
              <button
                type="submit"
                class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500"
              >
                {{ editingProduct ? 'ویرایش' : 'افزودن' }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
</script>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'
import { useAuth } from '~/composables/useAuth'

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

// تعریف interface ها
interface ApiResponse<T> {
  data?: T
  success?: boolean
  message?: string
}

interface Category {
  id: string
  name: string
}

interface Product {
  id: string
  name: string
  sku: string
  brand: string
  image: string
  price: number
  installmentPrice: number
  installments: Array<{
    count: number
    amount: number
  }>
  status: string
  salesCount: number
  totalSales: number
  categoryId?: string
  createdAt?: string
}

interface ProductForm {
  name: string
  sku: string
  price: number
  installmentPrice: number
  installments: Array<{
    count: number
    amount: number
  }>
  creditCheck: {
    required: boolean
    minScore: number
    minAge: number
  }
}

// داده‌های reactive
const products = ref<Product[]>([])
const categories = ref<Category[]>([])
const showProductModal = ref(false)
const editingProduct = ref<Product | null>(null)

const filters = ref({
  search: '',
  category: '',
  status: '',
  sort: 'name'
})

const productForm = ref<ProductForm>({
  name: '',
  sku: '',
  price: 0,
  installmentPrice: 0,
  installments: [],
  creditCheck: {
    required: false,
    minScore: 500,
    minAge: 18
  }
})

// Computed
const filteredProducts = computed(() => {
  let filtered = products.value

  if (filters.value.search) {
    const search = filters.value.search.toLowerCase()
    filtered = filtered.filter(product => 
      product.name.toLowerCase().includes(search) ||
      product.sku.toLowerCase().includes(search) ||
      product.brand.toLowerCase().includes(search)
    )
  }

  if (filters.value.category) {
    filtered = filtered.filter(product => product.categoryId === filters.value.category)
  }

  if (filters.value.status) {
    filtered = filtered.filter(product => product.status === filters.value.status)
  }

  // مرتب‌سازی
  filtered.sort((a, b) => {
    switch (filters.value.sort) {
      case 'price':
        return b.price - a.price
      case 'installment_count':
        return b.installments.length - a.installments.length
      case 'created_at':
        return new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime()
      default:
        return a.name.localeCompare(b.name)
    }
  })

  return filtered
})

// متدها
const openAddProductModal = () => {
  editingProduct.value = null
  resetProductForm()
  showProductModal.value = true
}

const editProduct = (product: Product) => {
  editingProduct.value = product
  productForm.value = {
    name: product.name,
    sku: product.sku,
    price: product.price,
    installmentPrice: product.installmentPrice,
    installments: [...product.installments],
    creditCheck: {
      required: true,
      minScore: 500,
      minAge: 18
    }
  }
  showProductModal.value = true
}

const closeProductModal = () => {
  showProductModal.value = false
  editingProduct.value = null
  resetProductForm()
}

const resetProductForm = () => {
  productForm.value = {
    name: '',
    sku: '',
    price: 0,
    installmentPrice: 0,
    installments: [],
    creditCheck: {
      required: false,
      minScore: 500,
      minAge: 18
    }
  }
}

const addInstallment = () => {
  productForm.value.installments.push({
    count: 0,
    amount: 0
  })
}

const removeInstallment = (index: number) => {
  productForm.value.installments.splice(index, 1)
}

const saveProduct = async () => {
  try {
    if (editingProduct.value) {
      await $fetch(`/api/admin/installment-payments/products/${editingProduct.value.id}`, {
        method: 'PUT',
        body: productForm.value
      })
    } else {
      await $fetch('/api/admin/installment-payments/products', {
        method: 'POST',
        body: productForm.value
      })
    }
    
    closeProductModal()
    fetchProducts()
  } catch (error) {
    console.error('خطا در ذخیره محصول:', error)
  }
}

const toggleProductStatus = async (product: Product) => {
  try {
    await $fetch(`/api/admin/installment-payments/products/${product.id}/toggle-status`, {
      method: 'POST'
    })
    await fetchProducts()
  } catch (error) {
    console.error('خطا در تغییر وضعیت محصول:', error)
  }
}

const applyFilters = () => {
  // فیلترها اعمال می‌شوند

}

// Utility functions
const getStatusClass = (status: string) => {
  switch (status) {
    case 'active':
      return 'bg-green-100 text-green-800'
    case 'inactive':
      return 'bg-red-100 text-red-800'
    case 'pending':
      return 'bg-yellow-100 text-yellow-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getStatusText = (status: string) => {
  switch (status) {
    case 'active':
      return 'فعال'
    case 'inactive':
      return 'غیرفعال'
    case 'pending':
      return 'در انتظار تایید'
    default:
      return 'نامشخص'
  }
}

const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

// دریافت داده‌ها
const fetchProducts = async () => {
  try {
    const response: ApiResponse<Product[]> = await $fetch('/api/admin/installment-payments/products', {
      method: 'GET'
    })
    products.value = response.data || []
  } catch (error) {
    console.error('خطا در دریافت محصولات:', error)
  }
}

const fetchCategories = async () => {
  try {
    const response: ApiResponse<Category[]> = await $fetch('/api/admin/product-categories', {
      method: 'GET'
    })
    categories.value = response.data || []
  } catch (error) {
    console.error('خطا در دریافت دسته‌بندی‌ها:', error)
  }
}

onMounted(async () => {
  await checkAuth()
  if (!hasAccess.value) {
    return
  }
  fetchProducts()
  fetchCategories()
})
</script> 
