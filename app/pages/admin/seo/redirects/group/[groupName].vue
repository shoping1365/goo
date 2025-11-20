<template>
  <div class="group-redirects-management">
    <!-- Header -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-2xl font-bold text-gray-900 mb-2">ریدایرکت‌های گروه: {{ groupName }}</h1>
          <p class="text-gray-600">مدیریت ریدایرکت‌های گروه {{ groupName }}</p>
        </div>
        <button 
          class="bg-gray-500 hover:bg-gray-600 text-white px-4 py-2 rounded-lg flex items-center gap-2"
          @click="router.push('/admin/seo/redirects')"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
          </svg>
          بازگشت
        </button>
      </div>
    </div>

    <!-- Stats Cards -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-6">
      <TemplateCard title="کل ریدایرکت‌ها" :value="stats.total" variant="blue">
        <template #icon>
          <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7l5 5m0 0l-5 5m5-5H6"></path>
          </svg>
        </template>
      </TemplateCard>

      <TemplateCard title="فعال" :value="stats.active" variant="green">
        <template #icon>
          <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
          </svg>
        </template>
      </TemplateCard>

      <TemplateCard title="نیازمند بررسی" :value="stats.needsReview" variant="yellow">
        <template #icon>
          <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
          </svg>
        </template>
      </TemplateCard>

      <TemplateCard title="خطا" :value="stats.errors" variant="red">
        <template #icon>
          <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
        </template>
      </TemplateCard>
    </div>

    <!-- Redirects Table -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 w-full">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-medium text-gray-900">لیست ریدایرکت‌ها</h3>
      </div>
      
      <div class="overflow-x-auto">
        <table class="w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                مسیر مبدا
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                مسیر مقصد
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                نوع
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                وضعیت
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                تعداد بازدید
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                آخرین بازدید
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                عملیات
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="redirect in redirects" :key="redirect.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm font-medium text-gray-900">{{ redirect.sourcePath || 'نامشخص' }}</div>
                <div class="text-sm text-gray-500">{{ redirect.sourcePath ? `https://example.com${redirect.sourcePath}` : 'نامشخص' }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm font-medium text-gray-900">{{ redirect.targetPath || 'نامشخص' }}</div>
                <div class="text-sm text-gray-500">{{ redirect.targetPath ? `https://example.com${redirect.targetPath}` : 'نامشخص' }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getTypeClass(redirect.code)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ redirect.code || 'نامشخص' }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusClass(redirect.status)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ getStatusText(redirect.status) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ formatNumber(redirect.visitCount) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ formatDate(redirect.lastVisitedAt) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <div class="flex gap-2">
                  <button 
                    v-if="hasPermission('seo.update')"
                    :class="redirect.status === 'active' ? 'text-red-600 hover:text-red-900' : 'text-green-600 hover:text-green-900'"
                    @click="toggleRedirect(redirect)"
                  >
                    {{ redirect.status === 'active' ? 'غیرفعال' : 'فعال' }}
                  </button>
                  <button 
                    v-if="canDeleteRedirect"
                    class="text-red-600 hover:text-red-900"
                    @click="deleteRedirect(redirect)"
                  >
                    حذف
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      
      <!-- Pagination -->
      <AdminCommonPagination
        v-if="totalPages > 1"
        :current-page="currentPage"
        :total-pages="totalPages"
        :total="totalItems"
        :items-per-page="itemsPerPage"
        @page-changed="handlePageChange"
        @items-per-page-changed="handleItemsPerPageChange"
      />
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
// eslint-disable-next-line @typescript-eslint/no-explicit-any
declare const useAuth: () => { user: any; hasPermission: (perm: string) => boolean }
// eslint-disable-next-line @typescript-eslint/no-explicit-any
declare const useRouter: () => any
// eslint-disable-next-line @typescript-eslint/no-explicit-any
declare const useRoute: () => any
</script>

<script setup lang="ts">
import { computed, onMounted, reactive, ref } from 'vue';

const router = useRouter()

const { hasPermission } = useAuth()

const canDeleteRedirect = computed(() => hasPermission('seo.delete'))

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// Get group name from route
const route = useRoute()
const groupName = computed(() => decodeURIComponent(route.params.groupName as string))

// State
const redirects = ref([])

// Pagination state
const currentPage = ref(1)
const itemsPerPage = ref(20)
const totalItems = ref(0)
const totalPages = computed(() => Math.ceil(totalItems.value / itemsPerPage.value))

// Stats
const stats = reactive({
  total: 0,
  active: 0,
  needsReview: 0,
  errors: 0
})

// Helper functions
const getTypeClass = (code: number | string) => {
  if (!code) {
    return 'bg-gray-100 text-gray-800'
  }
  
  const type = code.toString()
  switch (type) {
    case '301': return 'bg-blue-100 text-blue-800'
    case '302': return 'bg-yellow-100 text-yellow-800'
    case '307': return 'bg-purple-100 text-purple-800'
    case '308': return 'bg-indigo-100 text-indigo-800'
    case 'meta': return 'bg-green-100 text-green-800'
    case 'js': return 'bg-pink-100 text-pink-800'
    default: return 'bg-gray-100 text-gray-800'
  }
}

const getStatusClass = (status: string) => {
  if (!status || typeof status !== 'string') {
    return 'bg-gray-100 text-gray-800'
  }
  
  switch (status) {
    case 'active':
      return 'bg-green-100 text-green-800'
    case 'inactive':
      return 'bg-red-100 text-red-800'
    case 'error':
      return 'bg-orange-100 text-orange-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getStatusText = (status: string) => {
  if (!status || typeof status !== 'string') {
    return 'نامشخص'
  }
  
  switch (status) {
    case 'active':
      return 'فعال'
    case 'inactive':
      return 'غیرفعال'
    case 'error':
      return 'خطا'
    default:
      return 'نامشخص'
  }
}

const formatDate = (date: string | null) => {
  if (!date) return 'هیچ‌وقت'
  return new Date(date).toLocaleDateString('fa-IR')
}

const formatNumber = (num: number | undefined | null) => {
  return (num || 0).toLocaleString()
}

const updateStats = () => {
  stats.total = redirects.value.length
  stats.active = redirects.value.filter(r => r.status === 'active').length
  stats.needsReview = redirects.value.filter(r => (r.visitCount || 0) === 0 && r.status === 'active').length
  stats.errors = redirects.value.filter(r => r.status === 'error').length
}

// Load data from API
const loadRedirects = async () => {
  try {
    const offset = (currentPage.value - 1) * itemsPerPage.value
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const response: any = await $fetch('/api/admin/seo/redirects', {
      query: {
        limit: itemsPerPage.value,
        offset: offset,
        group: groupName.value
      }
    })
    if (response.success) {
      redirects.value = response.data
      totalItems.value = response.total || response.data.length
      updateStats()
    }
  } catch (error) {
    console.error('خطا در دریافت ریدایرکت‌ها:', error)
  }
}

// Pagination handlers
const handlePageChange = (page: number) => {
  currentPage.value = page
  loadRedirects()
}

const handleItemsPerPageChange = (perPage: number) => {
  itemsPerPage.value = perPage
  currentPage.value = 1
  loadRedirects()
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const toggleRedirect = async (redirect: any) => {
  try {
    // TODO: Implement toggle endpoint
    alert('قابلیت تغییر وضعیت ریدایرکت در حال توسعه است')
  } catch (error) {
    console.error('خطا در تغییر وضعیت ریدایرکت:', error)
  }
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const deleteRedirect = async (redirect: any) => {
  if (confirm('آیا از حذف این تغییر مسیر اطمینان دارید؟')) {
    try {
      // TODO: Implement delete endpoint
      alert('قابلیت حذف ریدایرکت در حال توسعه است')
    } catch (error) {
      console.error('خطا در حذف ریدایرکت:', error)
    }
  }
}

// Initialize on mount
onMounted(() => {
  loadRedirects()
})
</script>

<style scoped>
.group-redirects-management {
  max-width: 1200px;
  margin: 0 auto;
  padding: 1rem;
}
</style>
