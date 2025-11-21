<template>
  <ClientOnly>
    <div class="min-h-screen bg-blue-50">
      <!-- Header -->
      <div class="bg-white shadow-sm border-b border-gray-200">
        <div class="w-full px-4 sm:px-6 lg:px-8">
          <div class="flex justify-between items-center py-3">
            <div>
              <h1 class="text-lg font-bold text-gray-900">مشخصات فنی محصول</h1>
              <p class="text-xs text-gray-600 mt-1">مدیریت ویژگی‌ها و مشخصات فنی محصولات</p>
            </div>
            <div class="flex space-x-2 space-x-reverse">
              <NuxtLink 
                v-if="hasPermission('attribute.create')"
                to="/admin/product-management/attributes/new"
                class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-blue-400 to-blue-600 hover:from-blue-500 hover:to-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105"
                @click.prevent="navigateTo('/admin/product-management/attributes/new')"
              >
                مورد جدید
              </NuxtLink>
            </div>
          </div>
        </div>
      </div>

      <!-- Main Content Container with Large Spacing -->
      <div class="w-full px-4 py-4">
        <!-- Delete Confirm Modal -->
        <DeleteConfirmModal
          ref="deleteModalRef"
          single-delete-title="تایید حذف ویژگی"
          single-delete-message="آیا از حذف این ویژگی اطمینان دارید؟ این عملیات غیرقابل بازگشت است."
          bulk-delete-title="حذف گروهی ویژگی‌ها"
          bulk-delete-message="آیا از حذف ویژگی‌های انتخاب شده اطمینان دارید؟ این عملیات غیرقابل بازگشت است."
          :selected-count="selectedAttributes.length"
          @single-delete="handleSingleDelete"
          @bulk-delete="handleBulkDelete"
          @close-single="() => {}"
          @close-bulk="() => {}"
        />
        
        <!-- آمار کلی - کارت اول -->
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-6">
          <TemplateCard 
            title="کل ویژگی‌ها" 
            :value="attributeStats.total" 
            variant="blue"
          >
            <template #icon>
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"></path>
              </svg>
            </template>
          </TemplateCard>

          <TemplateCard 
            title="فعال" 
            :value="attributeStats.active" 
            variant="green"
          >
            <template #icon>
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </template>
          </TemplateCard>

          <TemplateCard 
            title="کل مقادیر" 
            :value="attributeStats.totalValues" 
            variant="purple"
          >
            <template #icon>
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
              </svg>
            </template>
          </TemplateCard>

          <TemplateCard 
            title="استفاده شده" 
            :value="attributeStats.used" 
            variant="orange"
          >
            <template #icon>
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
              </svg>
            </template>
          </TemplateCard>
        </div>

        <!-- مدیریت ویژگی‌ها - کارت دوم -->
        <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
          <div class="bg-gradient-to-r from-slate-50 to-gray-50 px-4 py-3 border-b border-gray-200">
            <div class="flex items-center justify-between">
              <div class="flex items-center">
                <div class="w-6 h-6 bg-slate-600 rounded-md flex items-center justify-center mr-2 ml-2">
                  <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16"></path>
                  </svg>
                </div>
                <h3 class="text-sm font-semibold text-gray-900">مدیریت مشخصات فنی</h3>
              </div>
              
              <div class="flex items-center space-x-2 space-x-reverse">
                <div v-if="selectedAttributes && selectedAttributes.length > 0" class="flex items-center space-x-2 space-x-reverse bg-red-50 rounded-md px-2 py-1.5 border border-red-200">
                  <span class="text-xs text-red-700 font-medium">{{ selectedAttributes.length }} مورد انتخاب شده</span>
                  <button
                    v-if="hasPermission('attribute.delete')"
                    class="inline-flex items-center px-2 py-1 border border-red-300 text-xs font-medium rounded-sm text-red-700 bg-red-100 hover:bg-red-200 transition-colors"
                    @click="openBulkDeleteConfirm()"
                  >
                    🗑️ حذف انتخاب شده‌ها
                  </button>
                </div>
                
                <!-- Search -->
                <div class="relative">
                  <input 
                    v-model="searchQuery"
                    type="text" 
                    class="block w-56 pl-8 pr-3 py-1.5 border border-gray-300 rounded-md leading-5 bg-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-xs shadow-sm" 
                    placeholder="جستجو نام مشخصه فنی"
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
          <div class="overflow-hidden">
            <div class="overflow-x-auto">
              <div dir="rtl">
                <table class="min-w-full divide-y divide-gray-200">
                  <thead class="bg-gray-50">
                    <tr>
                      <th scope="col" class="px-4 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">
                        <input 
                          type="checkbox" 
                          :checked="isAllSelected"
                          class="h-3 w-3 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                          @change="toggleSelectAll"
                        />
                      </th>
                      <th scope="col" class="py-2 text-right text-xs font-medium text-gray-600 uppercase tracking-wider pr-4">
                        ردیف
                      </th>
                      <th scope="col" class="px-4 py-2 text-right text-xs font-medium text-gray-600 uppercase tracking-wider">
                        نام مشخصه فنی
                      </th>
                      <th scope="col" class="px-4 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">
                        عملیات
                      </th>
                    </tr>
                  </thead>
                  <tbody class="bg-white divide-y divide-gray-200">
                    <!-- Loading State -->
                    <tr v-if="isLoading">
                      <td :colspan="4" class="px-4 py-8 text-center text-gray-500">
                        <div class="flex items-center justify-center">
                          <div class="animate-spin rounded-full h-6 w-6 border-b-2 border-blue-600 ml-2"></div>
                          در حال بارگذاری...
                        </div>
                      </td>
                    </tr>
                    
                    <!-- Empty State -->
                    <tr v-else-if="paginatedAttributes.length === 0">
                      <td :colspan="4" class="px-4 py-12 text-center text-gray-500">
                        <div class="text-center">
                          <svg class="w-12 h-12 text-gray-300 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"></path>
                          </svg>
                          <h3 class="text-lg font-medium text-gray-900 mb-2">هیچ ویژگی‌ای یافت نشد</h3>
                          <p class="text-gray-600 mb-4">اولین ویژگی خود را اضافه کنید</p>
                          <NuxtLink 
                            v-if="hasPermission('attribute.create')"
                            to="/admin/product-management/attributes/new"
                            class="inline-flex items-center px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-md hover:bg-blue-700"
                          >
                            افزودن ویژگی
                          </NuxtLink>
                        </div>
                      </td>
                    </tr>
                    
                    <!-- Attributes Data -->
                    <tr v-for="(attribute, index) in paginatedAttributes" v-else :key="attribute.id" class="hover:bg-gray-50 transition-colors">
                      <!-- Checkbox -->
                      <td class="px-4 py-3 whitespace-nowrap text-center">
                        <input 
                          v-model="selectedAttributes" 
                          type="checkbox"
                          :value="attribute.id"
                          class="h-3 w-3 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                        />
                      </td>

                      <!-- Row Number -->
                      <td class="py-3 whitespace-nowrap text-right pr-4">
                        <span class="inline-flex items-center w-6 h-6 bg-gray-100 text-gray-700 text-xs font-medium rounded-full text-right pr-4">{{ (currentPage - 1) * itemsPerPage + index + 1 }}</span>
                      </td>

                      <!-- Name -->
                      <td class="py-3 whitespace-nowrap text-xs font-medium text-gray-900 text-right pr-4">
                        {{ attribute.name }}
                      </td>

                      <!-- Actions -->
                      <td class="px-4 py-3 whitespace-nowrap text-center">
                        <div class="flex flex-row items-center justify-center gap-x-4">
                          <button 
                            v-if="hasPermission('attribute.update')"
                            class="inline-flex items-center px-2 py-1 border border-blue-300 text-xs font-medium rounded-sm text-blue-700 bg-blue-50 hover:bg-blue-100 transition-colors"
                            @click="navigateToEdit(attribute)"
                          >
                            ✏️ ویرایش
                          </button>
                          <button 
                            v-if="hasPermission('attribute.delete')"
                            class="inline-flex items-center px-2 py-1 border border-red-300 text-xs font-medium rounded-sm text-red-700 bg-red-50 hover:bg-red-100 transition-colors mr-2"
                            @click="confirmDelete(attribute)"
                          >
                            🗑️ حذف
                          </button>
                        </div>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>

              <!-- Pagination -->
              <Pagination
                :current-page="currentPage"
                :total-pages="totalPages"
                :total="totalCount"
                :per-page="itemsPerPage"
                @page-changed="goToPage"
                @per-page-changed="val => { itemsPerPage = val; currentPage = 1 }"
              />
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <template #fallback>
      <div class="min-h-screen bg-blue-50 flex items-center justify-center">
        <div class="text-center">
          <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600 mx-auto mb-3"></div>
          <div class="text-gray-600 text-sm">در حال بارگذاری...</div>
        </div>
      </div>
    </template>
  </ClientOnly>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
</script>

<script setup lang="ts">
import TemplateCard from '@/components/common/TemplateCard.vue'
import { navigateTo } from 'nuxt/app'
import { computed, onActivated, onMounted, ref, watch } from 'vue'
import Pagination from '~/components/admin/common/Pagination.vue'
import DeleteConfirmModal from '~/components/common/DeleteConfirmModal.vue'
import { useAuth } from '~/composables/useAuth'
import { useDeleteConfirmModal } from '~/composables/useDeleteConfirmModal'
import { useNotifier } from '~/composables/useNotifier'

// استفاده از useAuth برای چک کردن پرمیژن
const { hasPermission } = useAuth()


definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// Interfaces
interface Attribute {
  id: string
  name: string
  order: number
  description?: string
  isActive: boolean
  createdAt: string
}

interface AttributeResponse {
  id: number | string
  name: string
  sort_order?: number
  order?: number
  display_name?: string
  is_active?: boolean
  created_at?: string
}

interface ApiResponse {
  data?: AttributeResponse[]
  attributes?: AttributeResponse[]
  items?: AttributeResponse[]
  meta?: { total: number; per_page: number; total_pages: number }
  pagination?: { total: number; per_page: number; total_pages: number }
  [key: string]: unknown
}

// State
const attributes = ref<Attribute[]>([])
const searchQuery = ref('')
const selectedAttributes = ref<string[]>([])
const currentPage = ref(1)
const itemsPerPage = ref(10)
const serverMeta = ref<{ total: number; per_page: number; total_pages: number } | null>(null)

// Delete modal controller
const deleteModalCtl = useDeleteConfirmModal()
const deleteModalRef = deleteModalCtl.deleteModalRef
const openDeleteConfirm = deleteModalCtl.openDeleteConfirm
const openBulkDeleteModal = deleteModalCtl.openBulkDeleteConfirm

// ───────────────────────────────────────────────────────────────
// Fetch data from backend (Go API via Nuxt server proxy)
const fetched = ref<ApiResponse | null>(null)
const fetchError = ref<unknown>(null)
const isLoading = ref(false)

const refresh = async () => {
  try {
    isLoading.value = true
    const params = new URLSearchParams()
    params.set('page', String(currentPage.value))
    params.set('per_page', String(itemsPerPage.value))
    if (searchQuery.value) params.set('search', String(searchQuery.value))
    
    // اضافه کردن timestamp برای جلوگیری از cache
    params.set('_t', String(Date.now()))
    
    const res = await $fetch<ApiResponse>(`/api/attributes?${params.toString()}`)
    fetched.value = res
    fetchError.value = null
  } catch (e) {
    fetchError.value = e
    fetched.value = null
  } finally {
    isLoading.value = false
  }
}

// بارگذاری اولیه داده‌ها
await refresh()

watch(fetched, (val) => {
  if (fetchError.value) {
    // Failed to fetch attributes
    attributes.value = []
    serverMeta.value = null
    return
  }

  if (!val) {
    attributes.value = []
    serverMeta.value = null
    return
  }

  // Normalize response shape
  let list: AttributeResponse[] = []
  if (Array.isArray(val)) {
    list = val
  } else if (val.data && Array.isArray(val.data)) {
    list = val.data
  } else if (val.attributes && Array.isArray(val.attributes)) {
    list = val.attributes
  } else if (val.items && Array.isArray(val.items)) {
    list = val.items
  }

  const normalized: Attribute[] = list.map((attr) => ({
    id: String(attr?.id ?? ''),
    name: String(attr?.name ?? ''),
    order: Number(attr?.sort_order ?? attr?.order ?? 0),
    description: String(attr?.display_name ?? ''),
    isActive: Boolean(attr?.is_active ?? true),
    createdAt: String(attr?.created_at ?? '')
  }))
  
  // به‌روزرسانی داده‌ها با force reactivity
  attributes.value = [...normalized]

  // Capture server-side pagination meta if present
  if (val && typeof val === 'object' && val.meta) {
    const meta = val.meta
    const total = Number(meta.total ?? 0)
    const per_page = Number(meta.per_page ?? itemsPerPage.value)
    const total_pages = Number(meta.total_pages ?? Math.ceil(total / Math.max(per_page, 1)))
    serverMeta.value = { total, per_page, total_pages }
  } else if (val && typeof val === 'object' && val.pagination) {
    const pagination = val.pagination
    const total = Number(pagination.total ?? 0)
    const per_page = Number(pagination.per_page ?? itemsPerPage.value)
    const total_pages = Number(pagination.total_pages ?? Math.ceil(total / Math.max(per_page, 1)))
    serverMeta.value = { total, per_page, total_pages }
  } else {
    serverMeta.value = null
  }
}, { immediate: true, deep: true })

// Refresh list whenever page is activated (e.g., after navigating back)
onActivated(() => {
  refresh()
})

// Stats
type StatsApiResponse = { total: number; active: number; total_values: number; used: number }
const statsData = ref<StatsApiResponse | null>(null)
const statsError = ref<unknown>(null)
try {
  statsData.value = await $fetch<StatsApiResponse>('/api/attributes/stats')
  statsError.value = null
} catch (e: unknown) {
  statsError.value = e
  statsData.value = null
}

// Computed stats (fallback to derived values if stats API fails)
interface AttributeStats { total: number; active: number; totalValues: number; used: number }
const attributeStats = computed<AttributeStats>(() => {
  if (!statsError.value && statsData.value) {
    return {
      total: statsData.value.total || 0,
      active: statsData.value.active || 0,
      totalValues: statsData.value.total_values || 0,
      used: statsData.value.used || 0
    }
  }
  // Fallback using attributes list (less precise)
  const total = attributes.value.length
  const active = attributes.value.filter(attr => attr && attr.isActive).length
  const totalValues = total * 3 // Estimate
  const used = attributes.value.filter(attr => attr && attr.order > 0).length
  return { total, active, totalValues, used }
})

const filteredAttributes = computed(() => {
  if (!attributes.value || !Array.isArray(attributes.value)) {
    return []
  }
  
  // For server-side pagination, we don't filter on client side
  // The server handles both pagination and search filtering
  if (serverMeta.value) {
    return attributes.value
  }
  
  // For client-side filtering (when no server meta)
  if (!searchQuery.value) {
    return attributes.value
  }
  
  return attributes.value.filter(attr => 
    attr && attr.name && attr.name.toLowerCase().includes(searchQuery.value.toLowerCase())
  )
})

// Pagination computed properties
const totalPages = computed(() => {
  return serverMeta.value?.total_pages ?? Math.ceil((filteredAttributes.value?.length || 0) / itemsPerPage.value)
})

const totalCount = computed(() => serverMeta.value?.total ?? (filteredAttributes.value?.length || 0))

const paginatedAttributes = computed(() => {
  if (!filteredAttributes.value || !Array.isArray(filteredAttributes.value)) {
    return []
  }
  
  // Server-side pagination: data is already paginated by server
  // Client-side pagination: slice the filtered data
  if (serverMeta.value) {
    // Server already sent the correct page data
    return filteredAttributes.value
  } else {
    // Client-side pagination for local filtering
    const start = (currentPage.value - 1) * itemsPerPage.value
    const end = start + itemsPerPage.value
    return filteredAttributes.value.slice(start, end)
  }
})



const isAllSelected = computed(() => {
  if (!paginatedAttributes.value || !Array.isArray(paginatedAttributes.value) || 
      !selectedAttributes.value || !Array.isArray(selectedAttributes.value)) {
    return false
  }
  
  return paginatedAttributes.value.length > 0 && 
         selectedAttributes.value.length === paginatedAttributes.value.length
})

// Methods
const confirmDelete = (attribute: Attribute) => {
  if (!attribute || !attribute.id) return
  openDeleteConfirm(attribute.id)
}

// Handle single delete (from modal)
const handleSingleDelete = async (attributeId: number | string) => {
  try {
    const numericId = parseInt(String(attributeId), 10)
    if (isNaN(numericId)) return
    await $fetch(`/api/attributes/${numericId}`, { method: 'DELETE' })
    localStorage.removeItem('attributeNewDraft')
    
    // Refresh the list after deletion
    await refresh()
    
    // cleanup selection
    selectedAttributes.value = selectedAttributes.value.filter(id => id !== String(numericId))
  } catch (e) {
    useNotifier().error('خطا در حذف ویژگی')
  }
}

const toggleSelectAll = () => {
  if (!paginatedAttributes.value || !Array.isArray(paginatedAttributes.value) || 
      !selectedAttributes.value) {
    return
  }
  
  if (isAllSelected.value) {
    selectedAttributes.value = []
  } else {
    selectedAttributes.value = paginatedAttributes.value.map(attr => attr.id).filter(Boolean)
  }
}

const handleBulkDelete = async () => {
  if (!selectedAttributes.value || !Array.isArray(selectedAttributes.value) || selectedAttributes.value.length === 0) return
  try {
    const numericIds = selectedAttributes.value.map(id => parseInt(String(id), 10)).filter(id => !isNaN(id))
    if (numericIds.length === 0) return
    await $fetch('/api/attributes/bulk-delete', { method: 'POST', body: { ids: numericIds } })
    localStorage.removeItem('attributeNewDraft')
    
    // Refresh the list after bulk deletion
    await refresh()
    
    selectedAttributes.value = []
  } catch (e) {
    useNotifier().error('خطا در حذف گروهی ویژگی‌ها')
  }
}

const openBulkDeleteConfirm = () => {
  if (!selectedAttributes.value || selectedAttributes.value.length === 0) return
  openBulkDeleteModal()
}

const navigateToEdit = (attr: Attribute) => {
  if (!attr || !attr.id) return
  // Navigating to edit id
  navigateTo(`/admin/product-management/attributes/edit/${attr.id}`)
}



// Pagination methods
const goToPage = async (page: number) => {
  if (page >= 1 && page <= totalPages.value && page !== currentPage.value) {
    selectedAttributes.value = [] // Clear selection when changing page
    currentPage.value = page
    await refresh()
  }
}



// Watch for search query changes to reset pagination
watch(searchQuery, async () => {
  currentPage.value = 1
  selectedAttributes.value = []
  await refresh()
})

// Watch for items per page changes
watch(itemsPerPage, async () => {
  currentPage.value = 1
  selectedAttributes.value = []
  await refresh()
})

// Clamp current page when filtered list or items per page changes
watch([filteredAttributes, itemsPerPage], () => {
  const tp = totalPages.value
  if (tp === 0) {
    currentPage.value = 1
  } else if (currentPage.value > tp) {
    currentPage.value = tp
  }
})

// Lifecycle
onMounted(() => {
  // Initial load handled by useFetch; nothing else required here.
  // Attributes page mounted
})
</script> 

