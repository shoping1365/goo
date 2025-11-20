<template>
  <ClientOnly>
    <!-- Page Header -->
    <div class="bg-white shadow-sm border-b border-gray-200 mb-4">
      <div class="w-full px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center py-3">
          <div>
            <h1 class="text-lg font-bold text-gray-900">مدیریت محصولات</h1>
            <p class="text-xs text-gray-600 mt-1">مدیریت کامل کالاها و خدمات فروشگاه</p>
          </div>
          <div class="flex gap-2">
            <NuxtLink 
              v-if="hasPermission('product.create')"
              to="/admin/product-management/products/new"
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-blue-400 to-blue-600 hover:from-blue-500 hover:to-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105"
            >
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5 ml-2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M12 4v16m8-8H4" />
              </svg>
              مورد جدید
            </NuxtLink>
            <ExportExcelButton :data="getExportData()" filename="products.csv" />
          </div>
        </div>
      </div>
    </div>

    <div class="w-full px-4 py-4">
      
      <!-- آمار کلی - کارت اول -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden mb-6">
        <div class="bg-gradient-to-r from-blue-50 to-indigo-50 px-4 py-3 border-b border-gray-200">
          <div class="flex items-center">
            <div class="w-6 h-6 bg-blue-500 rounded-md flex items-center justify-center ml-2">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
              </svg>
            </div>
            <h3 class="text-sm font-semibold text-gray-900">آمار و اطلاعات کلی</h3>
          </div>
        </div>
        
        <div class="p-6">
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
            <TemplateCard 
              title="کل محصولات" 
              :value="stats.total.toLocaleString('en-US')" 
              variant="blue"
            >
              <template #icon>
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
                </svg>
              </template>
            </TemplateCard>

            <TemplateCard 
              title="موجود" 
              :value="stats.inStock.toLocaleString('en-US')" 
              variant="green"
            >
              <template #icon>
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
              </template>
            </TemplateCard>

            <TemplateCard 
              title="کم‌موجود" 
              :value="stats.lowStock.toLocaleString('en-US')" 
              variant="yellow"
            >
              <template #icon>
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5C2.962 18.333 3.924 20 5.464 20z"></path>
                </svg>
              </template>
            </TemplateCard>

            <TemplateCard 
              title="ناموجود" 
              :value="stats.outStock.toLocaleString('en-US')" 
              variant="red"
            >
              <template #icon>
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
              </template>
            </TemplateCard>
          </div>
        </div>
      </div>

      <!-- فیلترهای جستجو -->
      <ProductFilters @filters-changed="handleFiltersChanged" />
      
      <!-- مدیریت محصولات - کارت اصلی -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
        <div class="bg-gradient-to-r from-slate-50 to-gray-50 px-4 py-3 border-b border-gray-200">
          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-6 h-6 bg-slate-600 rounded-md flex items-center justify-center ml-2">
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16"></path>
                </svg>
              </div>
              <h3 class="text-sm font-semibold text-gray-900">مدیریت محصولات</h3>
            </div>
            
            <div class="flex items-center space-x-2 space-x-reverse">
              <!-- Bulk Actions -->

              <!-- Search -->
              <div class="relative">
                <input 
                  v-model="searchTerm"
                  type="text" 
                  class="block w-56 pl-8 pr-3 py-1.5 border border-gray-300 rounded-md leading-5 bg-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-xs shadow-sm" 
                  placeholder="جستجو سریع نام یا کد محصول"
                  dir="rtl"
                />
                <div class="absolute inset-y-0 left-0 pl-2 flex items-center pointer-events-none">
                  <svg class="h-4 w-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
                  </svg>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Table Content -->
        <ProductTable 
          ref="productTableRef"
          :search-term="searchTerm" 
          :filters="activeFilters"
          @stats-updated="updateStats" 
        />
      </div>
    </div>
  </ClientOnly>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
</script>

<script setup lang="ts">
import TemplateCard from '@/components/common/TemplateCard.vue'
import { useAuth } from '@/composables/useAuth'
import { ref } from 'vue'
import ExportExcelButton from '~/components/common/ExportExcelButton.vue'
import ProductFilters from '~/pages/admin/product-management/components/ProductFilters.vue'
import ProductTable from '~/pages/admin/product-management/components/ProductTable.vue'

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// Search term binding
const searchTerm = ref('')

// Stats object reactive
const stats = ref({ total: 0, inStock: 0, lowStock: 0, outStock: 0 })

// Filters object reactive
const activeFilters = ref({})

const productTableRef = ref(null)

function updateStats(newStats) {
  stats.value = newStats
}

function handleFiltersChanged(newFilters) {
  activeFilters.value = newFilters
}

function getExportData() {
  if (productTableRef.value && productTableRef.value.getExportData) {
    return productTableRef.value.getExportData()
  }
  return []
}

// استفاده از سیستم پرمیژن‌های پروژه
const { hasPermission } = useAuth()

</script>

<style scoped>
/* حذف تمام انیمیشن‌ها و پس‌زمینه‌های اضافی */
</style> 

