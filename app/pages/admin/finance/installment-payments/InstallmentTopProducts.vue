<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-6">
      <h3 class="text-lg font-semibold text-gray-900">محصولات برتر</h3>
      <div class="flex items-center space-x-2 space-x-reverse">
        <select v-model="selectedPeriod" class="text-sm border border-gray-300 rounded-md px-3 py-1">
          <option value="7">7 روز گذشته</option>
          <option value="30">30 روز گذشته</option>
          <option value="90">90 روز گذشته</option>
          <option value="365">یک سال گذشته</option>
        </select>
        <select v-model="sortBy" class="text-sm border border-gray-300 rounded-md px-3 py-1">
          <option value="requests">بیشترین درخواست</option>
          <option value="amount">بیشترین مبلغ</option>
          <option value="approval_rate">بالاترین نرخ تایید</option>
          <option value="avg_installments">بیشترین اقساط</option>
        </select>
        <button class="text-blue-600 hover:text-blue-800" @click="refreshData">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
        </button>
      </div>
    </div>

    <!-- Summary Cards -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
      <div class="bg-blue-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-blue-600 mb-1">{{ summary.totalProducts }}</div>
        <div class="text-sm text-gray-600">محصولات فعال</div>
      </div>
      <div class="bg-green-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-green-600 mb-1">{{ summary.totalRequests }}</div>
        <div class="text-sm text-gray-600">کل درخواست‌ها</div>
      </div>
      <div class="bg-yellow-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-yellow-600 mb-1">{{ formatCurrency(summary.totalAmount) }}</div>
        <div class="text-sm text-gray-600">کل مبلغ</div>
      </div>
      <div class="bg-purple-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-purple-600 mb-1">{{ summary.avgApprovalRate }}%</div>
        <div class="text-sm text-gray-600">میانگین نرخ تایید</div>
      </div>
    </div>

    <!-- Top Products List -->
    <div class="space-y-4">
      <div v-for="(product, index) in topProducts" :key="product.id" class="bg-gray-50 rounded-lg p-6">
        <div class="flex items-start space-x-4 space-x-reverse">
          <!-- Rank -->
          <div class="flex-shrink-0">
            <div class="w-10 h-10 rounded-full flex items-center justify-center" :class="getRankColor(index)">
              <span class="text-sm font-bold text-white">{{ index + 1 }}</span>
            </div>
          </div>

          <!-- Product Image -->
          <div class="flex-shrink-0">
            <img 
              :src="product.image || '/default-product.svg'" 
              :alt="product.name"
              class="w-16 h-16 object-cover rounded-lg border border-gray-200"
            >
          </div>

          <!-- Product Info -->
          <div class="flex-1 min-w-0">
            <div class="flex items-start justify-between">
              <div class="flex-1">
                <h4 class="text-lg font-semibold text-gray-900 truncate">{{ product.name }}</h4>
                <p class="text-sm text-gray-500 mb-2">{{ product.category }}</p>
                
                <div class="flex items-center space-x-4 space-x-reverse text-sm">
                  <span class="text-gray-600">SKU: {{ product.sku }}</span>
                  <span class="text-gray-600">قیمت: {{ formatCurrency(product.price) }}</span>
                  <div class="flex items-center">
                    <div class="w-2 h-2 rounded-full mr-1" :class="product.status === 'active' ? 'bg-green-500' : 'bg-red-500'"></div>
                    <span class="text-gray-600">{{ product.status === 'active' ? 'فعال' : 'غیرفعال' }}</span>
                  </div>
                </div>
              </div>

              <!-- Stats -->
              <div class="flex-shrink-0 text-right">
                <div class="text-lg font-bold text-gray-900">{{ product.requests }}</div>
                <div class="text-sm text-gray-500">درخواست</div>
              </div>
            </div>

            <!-- Progress Bars -->
            <div class="mt-4 space-y-2">
              <div class="flex items-center justify-between text-sm">
                <span class="text-gray-600">نرخ تایید</span>
                <span class="font-medium">{{ product.approvalRate }}%</span>
              </div>
              <div class="w-full bg-gray-200 rounded-full h-2">
                <div class="bg-green-500 h-2 rounded-full" :style="{ width: product.approvalRate + '%' }"></div>
              </div>
            </div>

            <!-- Installment Info -->
            <div class="mt-3 grid grid-cols-2 md:grid-cols-4 gap-6 text-sm">
              <div>
                <span class="text-gray-600 block">مبلغ کل</span>
                <span class="font-medium">{{ formatCurrency(product.totalAmount) }}</span>
              </div>
              <div>
                <span class="text-gray-600 block">میانگین اقساط</span>
                <span class="font-medium">{{ product.avgInstallments }} قسط</span>
              </div>
              <div>
                <span class="text-gray-600 block">میانگین امتیاز</span>
                <span class="font-medium">{{ product.avgCreditScore }}/100</span>
              </div>
              <div>
                <span class="text-gray-600 block">آخرین درخواست</span>
                <span class="font-medium">{{ product.lastRequest }}</span>
              </div>
            </div>

            <!-- Performance Trend -->
            <div class="mt-3 flex items-center justify-between text-sm">
              <span class="text-gray-600">روند عملکرد:</span>
              <div class="flex items-center" :class="product.trend > 0 ? 'text-green-600' : product.trend < 0 ? 'text-red-600' : 'text-gray-600'">
                <svg v-if="product.trend > 0" class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
                </svg>
                <svg v-else-if="product.trend < 0" class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 17h8m0 0V9m0 8l-8-8-4 4-6-6" />
                </svg>
                <svg v-else class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h8" />
                </svg>
                <span class="font-medium">
                  {{ product.trend > 0 ? '+' : '' }}{{ product.trend }}%
                </span>
              </div>
            </div>

            <!-- Action Buttons -->
            <div class="mt-4 flex items-center space-x-2 space-x-reverse">
              <button class="text-blue-600 hover:text-blue-800 text-sm font-medium" @click="viewDetails(product)">
                مشاهده جزئیات
              </button>
              <button class="text-green-600 hover:text-green-800 text-sm font-medium" @click="editInstallmentSettings(product)">
                تنظیمات اقساط
              </button>
              <button class="text-purple-600 hover:text-purple-800 text-sm font-medium" @click="viewHistory(product)">
                تاریخچه
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Load More -->
    <div v-if="hasMore" class="text-center mt-6">
      <button class="bg-blue-600 text-white px-6 py-2 rounded-md hover:bg-blue-700" @click="loadMore">
        نمایش بیشتر
      </button>
    </div>

    <!-- Empty State -->
    <div v-if="topProducts.length === 0" class="text-center py-12">
      <svg class="w-16 h-16 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
      </svg>
      <h3 class="text-lg font-medium text-gray-900 mb-2">هیچ محصولی یافت نشد</h3>
      <p class="text-gray-500">در بازه زمانی انتخاب شده هیچ محصولی با درخواست اقساط وجود ندارد.</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'
import { useRouter } from 'vue-router'

interface Product {
  id: number
  name: string
  sku: string
  category: string
  price: number
  image: string | null
  status: 'active' | 'inactive'
  requests: number
  approvalRate: number
  totalAmount: number
  avgInstallments: number
  avgCreditScore: number
  lastRequest: string
  trend: number
}

interface Summary {
  totalProducts: number
  totalRequests: number
  totalAmount: number
  avgApprovalRate: number
}

const selectedPeriod = ref('30')
const sortBy = ref('requests')
const hasMore = ref(false)
const page = ref(1)

const topProducts = ref<Product[]>([])
const summary = ref<Summary>({
  totalProducts: 0,
  totalRequests: 0,
  totalAmount: 0,
  avgApprovalRate: 0
})

const formatCurrency = (amount: number): string => {
  return new Intl.NumberFormat('fa-IR', {
    style: 'currency',
    currency: 'IRR',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount)
}

const getRankColor = (index: number): string => {
  switch (index) {
    case 0: return 'bg-yellow-500' // Gold
    case 1: return 'bg-gray-400'   // Silver
    case 2: return 'bg-yellow-600' // Bronze
    default: return 'bg-blue-500'
  }
}

const router = useRouter()

const fetchTopProducts = async (loadMoreData = false) => {
  try {
    const currentPage = loadMoreData ? page.value + 1 : 1
    
    interface TopProductsData {
      products: Product[]
      summary: Summary
      hasMore: boolean
    }
    interface ApiResponse {
      data?: TopProductsData
    }
    const response = await $fetch<ApiResponse>('/api/admin/installment-payments/top-products', {
      query: { 
        period: selectedPeriod.value,
        sort: sortBy.value,
        page: currentPage,
        limit: 10
      }
    })
    
    if (response.data) {
      if (loadMoreData) {
        topProducts.value = [...topProducts.value, ...response.data.products]
        page.value = currentPage
      } else {
        topProducts.value = response.data.products
        summary.value = response.data.summary
        page.value = 1
      }
      hasMore.value = response.data.hasMore
    }
  } catch (error) {
    console.error('خطا در دریافت محصولات برتر:', error)
    // در حالت واقعی، اینجا از createError استفاده می‌کنیم
  }
}

const refreshData = () => {
  fetchTopProducts()
}

const loadMore = () => {
  fetchTopProducts(true)
}

const viewDetails = (product: Product) => {
  // نمایش جزئیات محصول
  router.push(`/admin/product-management/products/${product.id}`)
}

const editInstallmentSettings = (product: Product) => {
  // ویرایش تنظیمات اقساط محصول
  router.push(`/admin/finance/installment-payments/products/${product.id}/settings`)
}

const viewHistory = (product: Product) => {
  // نمایش تاریخچه درخواست‌های اقساط محصول
  router.push(`/admin/finance/installment-payments/products/${product.id}/history`)
}

onMounted(() => {
  fetchTopProducts()
})

// Watch for changes
watch([selectedPeriod, sortBy], () => {
  fetchTopProducts()
})
</script>
