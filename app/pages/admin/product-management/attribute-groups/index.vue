<template>
  <ClientOnly>
    <!-- Page Header -->
    <div class="bg-white shadow-sm border-b border-gray-200 mb-4">
      <div class="w-full px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center py-3">
          <div>
            <h1 class="text-lg font-bold text-gray-900">مدیریت گروه‌های مشخصات فنی</h1>
            <p class="text-xs text-gray-600 mt-1">مدیریت و سازماندهی گروه‌های ویژگی‌ها و مشخصات فنی محصولات</p>
          </div>
          <div class="flex space-x-2 space-x-reverse">
            <button 
              v-if="selectedGroups.length > 0 && hasPermission('attribute-group.delete')"
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-red-600 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 shadow-md transition-all duration-200"
              @click="bulkDelete"
            >
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5 ml-2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M14.74 9l-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 01-2.244 2.077H8.084a2.25 2.25 0 01-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 00-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 013.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 00-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 00-7.5 0" />
              </svg>
              حذف ({{ selectedGroups.length }})
            </button>
            <button 
              v-if="hasPermission('attribute-group.create')"
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-blue-400 to-blue-600 hover:from-blue-500 hover:to-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105"
              @click="createNewGroup"
            >
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5 ml-2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M12 4v16m8-8H4" />
              </svg>
              گروه جدید
            </button>
          </div>
        </div>
      </div>
    </div>

    <div class="w-full px-4 py-4">
      
      <!-- آمار کلی - کارت اول -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-6">
        <TemplateCard 
          title="کل گروه‌ها" 
          :value="groupStats.total" 
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
          :value="groupStats.active" 
          variant="green"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </template>
        </TemplateCard>

        <TemplateCard 
          title="کل ویژگی‌ها" 
          :value="groupStats.totalAttributes" 
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
          :value="groupStats.used" 
          variant="orange"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
            </svg>
          </template>
        </TemplateCard>
      </div>

      <!-- مدیریت گروه‌ها - کارت دوم -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
        <div class="bg-gradient-to-r from-slate-50 to-gray-50 px-4 py-3 border-b border-gray-200">
          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-6 h-6 bg-slate-600 rounded-md flex items-center justify-center ml-2">
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16"></path>
                </svg>
              </div>
              <h3 class="text-sm font-semibold text-gray-900">مدیریت گروه‌های مشخصات فنی</h3>
            </div>
            
            <div class="flex items-center space-x-2 space-x-reverse">
              <!-- Search -->
              <div class="relative">
                <input 
                  v-model="searchQuery"
                  type="text" 
                  class="block w-56 pl-8 pr-3 py-1.5 border border-gray-300 rounded-md leading-5 bg-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-xs shadow-sm" 
                  placeholder="جستجو نام گروه"
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
                  <th scope="col" class="px-4 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">
                    ردیف
                  </th>
                  <th scope="col" class="px-4 py-2 text-right text-xs font-medium text-gray-600 uppercase tracking-wider">
                    نام گروه
                  </th>
                  <th scope="col" class="px-4 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">
                    تعداد ویژگی‌ها
                  </th>
                  <th scope="col" class="px-4 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">
                    دسته‌بندی
                  </th>
                  <th scope="col" class="px-4 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">
                    عملیات
                  </th>
                </tr>
              </thead>
              <tbody class="bg-white divide-y divide-gray-200">
                <tr v-for="(group, index) in paginatedGroups" :key="group.id" class="hover:bg-gray-50 transition-colors">
                  <!-- Checkbox -->
                  <td class="px-4 py-3 whitespace-nowrap text-center">
                    <input 
                      v-model="selectedGroups" 
                      type="checkbox"
                      :value="group.id"
                      class="h-3 w-3 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                    />
                  </td>

                  <!-- Row Number -->
                  <td class="px-4 py-3 whitespace-nowrap text-center">
                    <span class="inline-flex items-center justify-center w-6 h-6 bg-gray-100 text-gray-700 text-xs font-medium rounded-full">{{ (currentPage - 1) * itemsPerPage + index + 1 }}</span>
                  </td>

                  <!-- Name -->
                  <td class="px-4 py-3 whitespace-nowrap text-xs font-medium text-gray-900 text-right">
                    {{ group.name }}
                  </td>

                  <!-- Attributes Count -->
                  <td class="px-4 py-3 whitespace-nowrap text-center">
                    <span class="inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium bg-purple-100 text-purple-800">
                      {{ group.attributes_count }}
                    </span>
                  </td>

                  <!-- Category -->
                  <td class="px-4 py-3 whitespace-nowrap text-center text-xs text-gray-700">
                    {{ group.category_name || '—' }}
                  </td>

                  <!-- Actions -->
                  <td class="px-4 py-3 whitespace-nowrap text-center">
                    <div class="flex items-center justify-center space-x-1.5 space-x-reverse">
                      <button 
                        v-if="hasPermission('attribute-group.update')"
                        class="inline-flex items-center px-2 py-1 border border-blue-300 text-xs font-medium rounded-sm text-blue-700 bg-blue-50 hover:bg-blue-100 transition-colors"
                        @click="() => { 
                          // Edit button clicked for group 
                          // Current route
                          navigateToEdit(group.id); 
                        }"
                      >
                        ✏️ ویرایش گروه
                      </button>
                      <button 
                        v-if="hasPermission('attribute-group.delete')"
                        class="inline-flex items-center px-2 py-1 border border-red-300 text-xs font-medium rounded-sm text-red-700 bg-red-50 hover:bg-red-100 transition-colors"
                        @click="confirmDelete(group)"
                      >
                        🗑️ حذف
                      </button>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>

          <!-- Pagination component -->
          <Pagination
            v-if="filteredGroups && filteredGroups.length > 0"
            :current-page="currentPage"
            :total-pages="totalPages"
            :total="filteredGroups.length"
            :per-page="itemsPerPage"
            @page-changed="goToPage"
            @per-page-changed="handlePerPageChange"
          />

          <!-- Empty State -->
          <div v-if="!paginatedGroups || paginatedGroups.length === 0" class="text-center py-8">
            <div class="w-12 h-12 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-3">
              <svg class="w-6 h-6 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"></path>
              </svg>
            </div>
            <h3 class="text-sm font-medium text-gray-900 mb-1">هیچ گروهی یافت نشد</h3>
            <p class="text-xs text-gray-500 mb-4">برای شروع، یک گروه جدید اضافه کنید.</p>
            <button 
              v-if="hasPermission('attribute-group.create')"
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-blue-400 to-blue-600 hover:from-blue-500 hover:to-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105"
              @click="createNewGroup"
            >
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5 ml-2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M12 4v16m8-8H4" />
              </svg>
              گروه جدید
            </button>
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
declare const useHead: (head: { title?: string }) => void
</script>

<script setup lang="ts">
import Pagination from '@/components/admin/common/Pagination.vue'
import TemplateCard from '@/components/common/TemplateCard.vue'
import { computed, onMounted, ref, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useAuth } from '~/composables/useAuth'
import { useConfirmDialog } from '~/composables/useConfirmDialog'
import { useNotifier } from '~/composables/useNotifier'
// 
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// Import useAuth for permission checking
const { hasPermission } = useAuth()


// Set page title
useHead({
  title: 'مدیریت گروه‌های مشخصات فنی'
})

// Interfaces
interface AttributeGroup {
  id: string
  name: string
  attributes_count: number
  category_name?: string
  used?: boolean
  description?: string
  createdAt?: string
}

interface GroupStats {
  total: number
  active: number
  totalAttributes: number
  used: number
}

interface ApiResponse {
  data: AttributeGroup[]
  meta?: Record<string, unknown>
}

// State
const groups = ref<AttributeGroup[]>([])
const searchQuery = ref('')
const selectedGroups = ref<string[]>([])
const currentPage = ref(1)
const itemsPerPage = ref(10)
const router = useRouter()

// Computed with null checks
const groupStats = computed((): GroupStats => {
  if (!Array.isArray(groups.value)) {
    return { total: 0, active: 0, totalAttributes: 0, used: 0 }
  }

  const total = groups.value.length

  const totalAttributes = groups.value.reduce((sum, g) => sum + (g.attributes_count || 0), 0)

  const active = groups.value.filter(g => g.used).length

  return { total, active, totalAttributes, used: active }
})

const filteredGroups = computed((): AttributeGroup[] => {
  if (!groups.value || !Array.isArray(groups.value)) {
    return []
  }
  
  if (!searchQuery.value) {
    return groups.value
  }
  
  return groups.value.filter(group => 
    group && group.name && group.name.toLowerCase().includes(searchQuery.value.toLowerCase())
  )
})

// Pagination computed properties
const totalPages = computed(() => {
  return Math.ceil((filteredGroups.value?.length || 0) / itemsPerPage.value)
})

const paginatedGroups = computed((): AttributeGroup[] => {
  if (!filteredGroups.value || !Array.isArray(filteredGroups.value)) {
    return []
  }
  
  const start = (currentPage.value - 1) * itemsPerPage.value
  const end = start + itemsPerPage.value
  return filteredGroups.value.slice(start, end)
})

const isAllSelected = computed(() => {
  if (!paginatedGroups.value || !Array.isArray(paginatedGroups.value) || 
      !selectedGroups.value || !Array.isArray(selectedGroups.value)) {
    return false
  }
  
  return paginatedGroups.value.length > 0 && 
         selectedGroups.value.length === paginatedGroups.value.length
})

// Methods
const loadGroups = async (): Promise<void> => {
  try {
    // Call Nuxt server proxy → Go backend
    const res = await $fetch<ApiResponse | AttributeGroup[]>('/api/attribute-groups', {
      params: { per_page: 100, page: 1 },
    })

    if (res && 'data' in res && Array.isArray(res.data)) {
      groups.value = res.data
    } else if (Array.isArray(res)) {
      // In case backend returns raw array (no meta wrapper)
      groups.value = res
    } else {
      groups.value = []
    }
  } catch (error) {
    // خطا در دریافت گروه مشخصات
    groups.value = []
  }
}

const confirmDelete = async (group: AttributeGroup): Promise<void> => {
  if (!group || !group.name) return

  const { confirm } = useConfirmDialog()
  const ok = await confirm({ title:'تأیید حذف', message:`آیا از حذف گروه "${group.name}" اطمینان دارید؟`, confirmText:'حذف', cancelText:'انصراف', type:'danger' })
  if (!ok) return

  try {
    await $fetch(`/api/attribute-groups/${group.id}` , { method: 'DELETE' })
    // حذف از لیست محلی پس از موفقیت
    groups.value = groups.value.filter(g => g && g.id !== group.id)
    selectedGroups.value = selectedGroups.value.filter(id => id !== group.id)
    useNotifier().success('گروه حذف شد')
  } catch (err) {
    // خطا در حذف گروه
    useNotifier().error('خطا در حذف گروه')
  }
}

const toggleSelectAll = (): void => {
  if (!paginatedGroups.value || !Array.isArray(paginatedGroups.value) || 
      !selectedGroups.value) {
    return
  }
  
  if (isAllSelected.value) {
    selectedGroups.value = []
  } else {
    selectedGroups.value = paginatedGroups.value.map(group => group.id).filter(Boolean)
  }
}

 
const bulkDelete = async (): Promise<void> => {
  if (!selectedGroups.value || selectedGroups.value.length === 0) return

  const selectedCount = selectedGroups.value.length
  const { confirm } = useConfirmDialog()
  const ok = await confirm({ title:'حذف گروهی', message:`آیا از حذف ${selectedCount} مورد انتخاب شده اطمینان دارید؟`, confirmText:'حذف', cancelText:'انصراف', type:'danger' })
  if (!ok) return

  try {
    // استفاده از endpoint bulk-delete اگر وجود دارد در غیر اینصورت حلقه
    await Promise.all(selectedGroups.value.map(id => $fetch(`/api/attribute-groups/${id}`, { method:'DELETE' }).catch(()=>null)))
    groups.value = groups.value.filter(group => group && !selectedGroups.value.includes(group.id))
    selectedGroups.value = []
    useNotifier().success('حذف انجام شد')
  } catch (err) {
    // خطا در حذف گروه‌ها
    useNotifier().error('خطا در حذف گروه‌ها')
  }
}

const navigateToEdit = (groupId: string): void => {
  // Navigating to edit group
  // Navigate to the correct edit group page
  const editRoute = `/admin/product-management/attribute-groups/${groupId}/edit`
  // Target route
  try {
    router.push(editRoute)
    // Router.push called successfully
  } catch (error) {
    // Router.push failed
    window.location.href = editRoute
  }
}

const createNewGroup = (): void => {
  // Navigating to create new group
  
  // Clear any previously stored group name
  localStorage.removeItem('selectedGroupName')
  
  // Navigate to the correct new group page
  router.push('/admin/product-management/attribute-groups/new')
}

// Pagination methods
const goToPage = (page: number): void => {
  if (page >= 1 && page <= totalPages.value) {
    currentPage.value = page
    selectedGroups.value = [] // Clear selection when changing page
  }
}

const handlePerPageChange = (newPerPage: number): void => {
  itemsPerPage.value = newPerPage
  currentPage.value = 1
  selectedGroups.value = []
}

// Watch for search query changes to reset pagination
watch(searchQuery, () => {
  currentPage.value = 1
  selectedGroups.value = []
})

// Watch for items per page changes
watch(itemsPerPage, () => {
  currentPage.value = 1
  selectedGroups.value = []
})

// Lifecycle
onMounted(() => {
  loadGroups()
  // Groups loaded
  // Stats
  
  // Watch for changes
  watch(() => groups.value.length, (newLength) => {
    // Groups length changed
  })
  
  watch(() => filteredGroups.value.length, (newLength) => {
    // Filtered groups length
  })
  
  watch(() => totalPages.value, (newTotal) => {
    // Total pages
  })
})
</script> 

