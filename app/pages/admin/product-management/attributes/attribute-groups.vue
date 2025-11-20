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
              class="inline-flex items-center px-3 py-1.5 border border-green-300 text-xs font-medium rounded-md text-white bg-green-600 hover:bg-green-700 shadow-sm"
              @click="createNewGroup"
            >
              گروه جدید
            </button>
          </div>
        </div>
      </div>
    </div>

    <div class="w-full px-4 py-4">
      
      <!-- آمار کلی - کارت اول -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden mb-6">
        <div class="bg-gradient-to-r from-blue-50 to-indigo-50 px-4 py-3 border-b border-gray-200">
          <div class="flex items-center">
            <div class="w-6 h-6 bg-blue-500 rounded-md flex items-center justify-center mr-2">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
              </svg>
            </div>
            <h3 class="text-sm font-semibold text-gray-900">آمار و اطلاعات کلی</h3>
          </div>
        </div>
        
        <div class="p-6">
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-3">
            <!-- کل گروه‌ها -->
            <div class="bg-gradient-to-br from-blue-50 to-blue-100 border border-blue-200 rounded-md p-3 hover:shadow-sm transition-all duration-200">
              <div class="flex items-center justify-between">
                <div>
                  <p class="text-xs font-medium text-blue-700 mb-0.5">کل گروه‌ها</p>
                  <p class="text-lg font-bold text-blue-900">{{ groupStats.total }}</p>
                </div>
                <div class="w-8 h-8 bg-blue-500 rounded-md flex items-center justify-center">
                  <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"></path>
                  </svg>
                </div>
              </div>
            </div>

            <!-- فعال -->
            <div class="bg-gradient-to-br from-green-50 to-green-100 border border-green-200 rounded-md p-3 hover:shadow-sm transition-all duration-200">
              <div class="flex items-center justify-between">
                <div>
                  <p class="text-xs font-medium text-green-700 mb-0.5">فعال</p>
                  <p class="text-lg font-bold text-green-900">{{ groupStats.active }}</p>
                </div>
                <div class="w-8 h-8 bg-green-500 rounded-md flex items-center justify-center">
                  <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                  </svg>
                </div>
              </div>
            </div>

            <!-- کل ویژگی‌ها -->
            <div class="bg-gradient-to-br from-purple-50 to-purple-100 border border-purple-200 rounded-md p-3 hover:shadow-sm transition-all duration-200">
              <div class="flex items-center justify-between">
                <div>
                  <p class="text-xs font-medium text-purple-700 mb-0.5">کل ویژگی‌ها</p>
                  <p class="text-lg font-bold text-purple-900">{{ groupStats.totalAttributes }}</p>
                </div>
                <div class="w-8 h-8 bg-purple-500 rounded-md flex items-center justify-center">
                  <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
                  </svg>
                </div>
              </div>
            </div>

            <!-- استفاده شده -->
            <div class="bg-gradient-to-br from-orange-50 to-orange-100 border border-orange-200 rounded-md p-3 hover:shadow-sm transition-all duration-200">
              <div class="flex items-center justify-between">
                <div>
                  <p class="text-xs font-medium text-orange-700 mb-0.5">استفاده شده</p>
                  <p class="text-lg font-bold text-orange-900">{{ groupStats.used }}</p>
                </div>
                <div class="w-8 h-8 bg-orange-500 rounded-md flex items-center justify-center">
                  <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
                  </svg>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- مدیریت گروه‌ها - کارت دوم -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
        <div class="bg-gradient-to-r from-slate-50 to-gray-50 px-4 py-3 border-b border-gray-200">
          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-6 h-6 bg-slate-600 rounded-md flex items-center justify-center mr-2">
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16"></path>
                </svg>
              </div>
              <h3 class="text-sm font-semibold text-gray-900">مدیریت گروه‌های مشخصات فنی</h3>
            </div>
            
            <div class="flex items-center space-x-2 space-x-reverse">
              <!-- Bulk Actions -->
              <div v-if="selectedGroups && selectedGroups.length > 0" class="flex items-center space-x-2 space-x-reverse bg-red-50 rounded-md px-2 py-1.5 border border-red-200">
                <span class="text-xs text-red-700 font-medium">{{ selectedGroups.length }} مورد انتخاب شده</span>
                <button 
                  class="inline-flex items-center px-2 py-1 border border-red-300 text-xs font-medium rounded-sm text-red-700 bg-red-100 hover:bg-red-200 transition-colors"
                  @click="bulkDelete"
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
                  placeholder="جستجو نام گروه"
                  dir="rtl"
                />
                <div class="absolute inset-y-0 left-0 pl-2 flex items-center pointer-events-none">
                  <svg class="h-4 w-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
                  </svg>
                </div>
              </div>
              
              <!-- Search Button -->
              <button 
                class="inline-flex items-center px-3 py-1.5 border border-blue-300 text-xs font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 shadow-sm transition-colors"
                @click="performSearch"
              >
                جستجو
              </button>
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
                    ترتیب نمایش
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

                  <!-- Order -->
                  <td class="px-4 py-3 whitespace-nowrap text-center">
                    <span class="inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800">
                      {{ group.order }}
                    </span>
                  </td>

                  <!-- Actions -->
                  <td class="px-4 py-3 whitespace-nowrap text-center">
                    <div class="flex items-center justify-center space-x-1.5 space-x-reverse">
                      <button 
                        class="inline-flex items-center px-2 py-1 border border-blue-300 text-xs font-medium rounded-sm text-blue-700 bg-blue-50 hover:bg-blue-100 transition-colors"
                        @click="() => { 
                          console.log('🖱️ Edit button clicked for group:', group.name, 'ID:', group.id); 
                          console.log('📍 Current route:', $route.path);
                          navigateToEdit(group.id); 
                        }"
                      >
                        ✏️ ویرایش گروه
                      </button>
                      <button 
                        v-if="canDeleteAttributeGroup"
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

          <!-- Pagination -->
          <div v-if="filteredGroups && filteredGroups.length > 0" class="bg-white px-4 py-3 border-t border-gray-200 sm:px-6">
            <div class="flex items-center justify-between">
              <!-- Pagination Info -->
              <div class="flex items-center">
                <p class="text-xs text-gray-700">
                  نمایش 
                  <span class="font-medium">{{ paginationInfo.start }}</span>
                  تا 
                  <span class="font-medium">{{ paginationInfo.end }}</span>
                  از 
                  <span class="font-medium">{{ paginationInfo.total }}</span>
                  نتیجه
                </p>
                
                <!-- Items per page selector -->
                <div class="mr-4 flex items-center">
                  <label class="text-xs text-gray-700 ml-2">تعداد در صفحه:</label>
                  <select 
                    v-model="itemsPerPage" 
                    class="text-xs border border-gray-300 rounded px-2 py-1 focus:outline-none focus:ring-1 focus:ring-blue-500"
                  >
                    <option :value="5">5</option>
                    <option :value="10">10</option>
                    <option :value="15">15</option>
                    <option :value="25">25</option>
                    <option :value="50">50</option>
                  </select>
                </div>
              </div>

              <!-- Pagination Controls -->
              <div class="flex items-center space-x-1">
                <!-- Previous Button -->
                <button 
                  :disabled="currentPage === 1"
                  class="inline-flex items-center px-2 py-1 border border-gray-300 text-xs font-medium rounded-md text-gray-500 bg-white hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
                  @click="prevPage"
                >
                  <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"></path>
                  </svg>
                  قبلی
                </button>

                <!-- Page Numbers -->
                <div class="flex items-center space-x-1">
                  <template v-for="page in totalPages" :key="page">
                    <button 
                      v-if="page === 1 || page === totalPages || (page >= currentPage - 1 && page <= currentPage + 1)"
                      :class="[
                        'inline-flex items-center px-2 py-1 border text-xs font-medium rounded-md',
                        page === currentPage 
                          ? 'border-blue-500 bg-blue-50 text-blue-600' 
                          : 'border-gray-300 bg-white text-gray-500 hover:bg-gray-50'
                      ]"
                      @click="goToPage(page)"
                    >
                      {{ page }}
                    </button>
                    
                    <!-- Show dots for gaps -->
                    <span 
                      v-else-if="page === currentPage - 2 || page === currentPage + 2"
                      class="inline-flex items-center px-1 text-xs text-gray-500"
                    >
                      ...
                    </span>
                  </template>
                </div>

                <!-- Next Button -->
                <button 
                  :disabled="currentPage === totalPages"
                  class="inline-flex items-center px-2 py-1 border border-gray-300 text-xs font-medium rounded-md text-gray-500 bg-white hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
                  @click="nextPage"
                >
                  بعدی
                  <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
                  </svg>
                </button>
              </div>
            </div>
          </div>

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
              class="inline-flex items-center px-3 py-1.5 border border-green-300 text-xs font-medium rounded-md text-white bg-green-600 hover:bg-green-700 shadow-sm transition-colors"
              @click="createNewGroup"
            >
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
declare const navigateTo: (to: string) => Promise<void>
</script>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue';
import { useRouter } from 'vue-router';
import { useAuth } from '~/composables/useAuth';

definePageMeta({
  layout: 'admin-main'
})

// Set page title
useHead({
  title: 'مدیریت گروه‌های مشخصات فنی'
})

// Interfaces
interface AttributeGroup {
  id: string
  name: string
  order: number
  description?: string
  isActive: boolean
  createdAt: string
  attributesCount: number
}

// State
const groups = ref<AttributeGroup[]>([])
const searchQuery = ref('')
const selectedGroups = ref<string[]>([])
const currentPage = ref(1)
const itemsPerPage = ref(5)
const router = useRouter()

const { user, hasPermission } = useAuth()

const canDeleteAttributeGroup = computed(() => hasPermission('attribute-group.delete'))

// Computed with null checks
const groupStats = computed(() => {
  if (!groups.value || !Array.isArray(groups.value)) {
    return { total: 0, active: 0, totalAttributes: 0, used: 0 }
  }
  
  const total = groups.value.length
  const active = groups.value.filter(group => group && group.isActive).length
  const totalAttributes = groups.value.reduce((sum, group) => sum + (group.attributesCount || 0), 0)
  const used = groups.value.filter(group => group && group.order > 0).length

  return { total, active, totalAttributes, used }
})

const filteredGroups = computed(() => {
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

const paginatedGroups = computed(() => {
  if (!filteredGroups.value || !Array.isArray(filteredGroups.value)) {
    return []
  }
  
  const start = (currentPage.value - 1) * itemsPerPage.value
  const end = start + itemsPerPage.value
  return filteredGroups.value.slice(start, end)
})

const paginationInfo = computed(() => {
  const total = filteredGroups.value?.length || 0
  const start = total === 0 ? 0 : (currentPage.value - 1) * itemsPerPage.value + 1
  const end = Math.min(currentPage.value * itemsPerPage.value, total)
  
  return { start, end, total }
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
const loadGroups = () => {
  // Sample data for attribute groups (keep only the general specs group)
  groups.value = [
    {
      id: '1',
      name: 'مشخصات کلی',
      order: 1,
      description: 'مشخصات عمومی محصولات',
      isActive: true,
      createdAt: '2024-01-15T10:30:00Z',
      attributesCount: 8
    }
  ]
}

const confirmDelete = (group: AttributeGroup) => {
  if (!group || !group.name) return
  
  if (confirm(`آیا از حذف گروه "${group.name}" اطمینان دارید؟`)) {
    if (groups.value && Array.isArray(groups.value)) {
      groups.value = groups.value.filter(g => g && g.id !== group.id)
    }
    
    // Remove from selected if it was selected
    if (selectedGroups.value && Array.isArray(selectedGroups.value)) {
      selectedGroups.value = selectedGroups.value.filter(id => id !== group.id)
    }
    console.log(`گروه "${group.name}" حذف شد`)
  }
}

const toggleSelectAll = () => {
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

const bulkDelete = () => {
  if (!selectedGroups.value || !Array.isArray(selectedGroups.value) || 
      selectedGroups.value.length === 0) {
    return
  }
  
  const selectedCount = selectedGroups.value.length
  if (confirm(`آیا از حذف ${selectedCount} مورد انتخاب شده اطمینان دارید؟`)) {
    if (groups.value && Array.isArray(groups.value)) {
      groups.value = groups.value.filter(group => 
        group && !selectedGroups.value.includes(group.id)
      )
    }
    selectedGroups.value = []
    console.log(`${selectedCount} مورد حذف شد`)
  }
}

const navigateToEdit = (groupId: string) => {
  console.log(`🔄 Navigating to edit group with ID: ${groupId}`)
  
  // Navigate to the simple edit page with query parameter
  const editRoute = `/admin/attribute-groups-edit?id=${groupId}`
  console.log(`🎯 Target route: ${editRoute}`)
  
  try {
    // Use router.push instead of navigateTo for better debugging
    router.push(editRoute)
    console.log('✅ Router.push called successfully')
  } catch (error) {
    console.error('❌ Router.push failed:', error)
    // Fallback to window.location
    console.log('🔄 Trying window.location.href as fallback')
    window.location.href = editRoute
  }
}

const createNewGroup = () => {
  console.log('🔄 Navigating to create new group')
  
  // Clear any previously stored group name
  localStorage.removeItem('selectedGroupName')
  
  // Navigate to the edit-simple page without any query parameters for creating new group
  navigateTo('/admin/edit-simple')
}

const performSearch = () => {
  console.log('🔍 Performing search with query:', searchQuery.value)
  // Focus on search input if it exists
  const searchInput = document.querySelector('input[placeholder="جستجو نام گروه"]')
  if (searchInput) {
    (searchInput as HTMLInputElement).focus()
  }
}

// Pagination methods
const goToPage = (page: number) => {
  if (page >= 1 && page <= totalPages.value) {
    currentPage.value = page
    selectedGroups.value = [] // Clear selection when changing page
  }
}

const nextPage = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value++
    selectedGroups.value = []
  }
}

const prevPage = () => {
  if (currentPage.value > 1) {
    currentPage.value--
    selectedGroups.value = []
  }
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
  console.log('🔄 Groups loaded:', groups.value.length)
  console.log('📊 Stats:', groupStats.value)
  
  // Watch for changes
  watch(() => groups.value.length, (newLength) => {
    console.log('📈 Groups length changed to:', newLength)
  })
  
  watch(() => filteredGroups.value.length, (newLength) => {
    console.log('🔍 Filtered groups length:', newLength)
  })
  
  watch(() => totalPages.value, (newTotal) => {
    console.log('📄 Total pages:', newTotal)
  })
})
</script> 
