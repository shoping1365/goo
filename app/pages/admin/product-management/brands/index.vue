<template>
  <ClientOnly>
    <!-- Page Header -->
    <div class="bg-white shadow-sm border-b border-gray-200 mb-4">
      <div class="w-full px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center py-3">
  <div>
            <h1 class="text-lg font-bold text-gray-900">مدیریت برندها</h1>
            <p class="text-xs text-gray-600 mt-1">مدیریت کامل برندها و شرکت‌های تولیدکننده</p>
          </div>
          <div class="flex space-x-2 space-x-reverse">
            <NuxtLink v-if="hasPermission('brand.create')" to="/admin/product-management/brands/new">
              <TemplateButton
                bg-gradient="bg-gradient-to-r from-blue-400 to-blue-600"
                text-color="text-white"
                hover-class="hover:from-blue-500 hover:to-blue-700 hover:shadow-lg hover:scale-105"
                focus-class="focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
                size="medium"
              >
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5 ml-2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M12 4v16m8-8H4" />
                </svg>
                مورد جدید
              </TemplateButton>
            </NuxtLink>
            <TemplateButton
              bg-gradient="bg-gradient-to-r from-green-400 to-green-600"
              text-color="text-white"
              hover-class="hover:from-green-500 hover:to-green-700 hover:shadow-lg hover:scale-105"
              focus-class="focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
              size="medium"
            >
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
              </svg>
              خروجی Excel
            </TemplateButton>
            <TemplateButton
              bg-gradient="bg-gradient-to-r from-gray-200 to-gray-400"
              text-color="text-gray-800"
              hover-class="hover:from-gray-300 hover:to-gray-500 hover:shadow-lg hover:scale-105"
              focus-class="focus:ring-2 focus:ring-offset-2 focus:ring-gray-400"
              size="medium"
            >
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5 ml-2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M3 16.5v2.25A2.25 2.25 0 005.25 21h13.5A2.25 2.25 0 0021 18.75V16.5m-13.5-9L12 3m0 0l4.5 4.5M12 3v13.5" />
              </svg>
              ورود اطلاعات
            </TemplateButton>
          </div>
        </div>
      </div>
    </div>

    <div class="w-full px-4 py-4">
      
      <!-- آمار کلی - کارت اول -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-3 mb-6">
        <TemplateCard
          title="کل برندها"
          :value="brands.length"
          variant="purple"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"></path>
            </svg>
          </template>
        </TemplateCard>
        <TemplateCard
          title="منتشر شده"
          :value="publishedBrands.length"
          variant="green"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </template>
        </TemplateCard>
        <TemplateCard
          title="منتشر نشده"
          :value="unpublishedBrands.length"
          variant="orange"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5C2.962 18.333 3.924 20 5.464 20z"></path>
            </svg>
          </template>
        </TemplateCard>
        <TemplateCard
          title="با تصویر"
          :value="brandsWithImage.length"
          variant="blue"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
            </svg>
          </template>
        </TemplateCard>
      </div>

      <!-- جستجو -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden mb-6">
        <div class="bg-gradient-to-r from-slate-50 to-gray-50 px-4 py-3 border-b border-gray-200">
          <div class="flex items-center">
            <div class="w-6 h-6 bg-slate-600 rounded-md flex items-center justify-center ml-2">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
              </svg>
            </div>
            <h3 class="text-sm font-semibold text-gray-900">جستجوی برندها</h3>
          </div>
        </div>
        
        <div class="p-6">
      <div class="flex items-center gap-6">
        <div class="flex-1">
              <input id="brand-name" v-model="searchTerm" type="text" class="w-full rounded-lg border border-gray-200 shadow-sm bg-white focus:outline-none focus:ring-2 focus:ring-blue-400 py-3 px-6 text-gray-800 text-base transition-all duration-200" placeholder="نام برند یا نام رسمی را وارد کنید...">
            </div>
          </div>
        </div>
      </div>

      <!-- مدیریت برندها - کارت اصلی -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
        <div class="bg-gradient-to-r from-slate-50 to-gray-50 px-4 py-3 border-b border-gray-200">
          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-6 h-6 bg-slate-600 rounded-md flex items-center justify-center ml-2">
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16"></path>
                </svg>
              </div>
              <h3 class="text-sm font-semibold text-gray-900">فهرست برندها</h3>
            </div>
            
            <div class="flex items-center">
              <span class="text-xs text-gray-600">{{ filteredBrands.length }} برند</span>
        </div>
      </div>
    </div>

    <!-- Brands Table -->
        <div class="overflow-hidden">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">آیدی</th>
            <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تصویر</th>
            <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نام</th>
            <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">منتشر شده</th>
            <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">ترتیب نمایش</th>
            <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نام رسمی</th>
            <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-600">عملیات</th>
          </tr>
        </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="brand in filteredBrands" :key="brand.id" class="hover:bg-gray-50 transition-colors">
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900 text-right">{{ brand.id }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500 text-right">
                  <img :src="brand.image_url || '/statics/images/default-image_100.png'" class="w-12 h-12 object-contain border rounded-md shadow-sm" />
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-blue-700 hover:underline">
              <NuxtLink :to="`/brand/${brand.slug}`" target="_blank">{{ brand.name }}</NuxtLink>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500 text-right">
                  <div class="flex items-center">
                    <div v-if="brand.published" class="w-2 h-2 bg-green-500 rounded-full ml-2"></div>
                    <div v-else class="w-2 h-2 bg-red-500 rounded-full ml-2"></div>
              <svg v-if="brand.published" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-5 h-5 text-green-500">
                <path fill-rule="evenodd" d="M19.916 4.626a.75.75 0 01.208 1.153l-9.75 10.5a.75.75 0 01-1.31-.013l-5.25-7.5a.75.75 0 011.135-.967l4.723 6.741 9.137-9.91a.75.75 0 011.153-.208z" clip-rule="evenodd" />
              </svg>
              <svg v-else xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5 text-red-500">
                <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
              </svg>
                  </div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500 text-right">
                  <span v-if="editingOrderId !== brand.id" class="cursor-pointer hover:underline bg-gray-100 px-2 py-1 rounded text-xs" @click="startEditOrder(brand)">{{ brand.order || 0 }}</span>
                  <input v-else v-model="editingOrderValue" type="number" min="1" :max="filteredBrands.length" class="w-16 border rounded px-1 py-0.5 text-center text-xs ltr:text-left" dir="ltr" autofocus @blur="finishEditOrder(brand)" @keyup="handleOrderInputKey($event, brand)" />
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-700 text-right">{{ brand.official_name || '-' }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-center text-sm">
              <div class="flex items-center justify-center gap-2">
                <NuxtLink v-if="hasPermission('brand.update')" :to="`/admin/product-management/brands/${brand.id}/edit`">
                  <TemplateButton
                    text-color="text-blue-600"
                    bg-gradient="bg-blue-50"
                    hover-class="hover:bg-blue-100"
                    size="small"
                    :border-color="'border-0'"
                  >
                    ویرایش
                  </TemplateButton>
                </NuxtLink>
                <TemplateButton
                  v-if="hasPermission('brand.delete')"
                  text-color="text-red-600"
                  bg-gradient="bg-red-50"
                  hover-class="hover:bg-red-100"
                  size="small"
                  :border-color="'border-0'"
                  @click="deleteBrand(brand)"
                >
                  حذف
                </TemplateButton>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
    </div>
  </ClientOnly>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
</script>

<script setup lang="ts">
import TemplateButton from '@/components/common/TemplateButton.vue';
import { computed, onActivated, onMounted, provide, ref } from 'vue';

// Permission check (simple implementation without useAuth)
const hasPermission = (permission: string) => true

// Reactive list of brands
const brands = ref([])

// Fetch brands data
const fetchBrands = async () => {
  try {
    interface Brand {
      id?: number | string
      name?: string
      order?: number | null
      [key: string]: unknown
    }
    interface BrandsResponse {
      data?: Brand[]
      [key: string]: unknown
    }
    const response = await $fetch<Brand[] | BrandsResponse>('/api/brands')
    const raw = Array.isArray(response) ? response : (response.data || [])
    raw.forEach((b: Brand, idx: number) => {
      if (b.order === undefined || b.order === null) b.order = idx + 1
    })
    brands.value = raw
  } catch (error) {
    console.error('خطا در دریافت برندها:', error)
  }
}

// Load brands on mount
onMounted(() => {
  fetchBrands()
})

// Ensure list stays fresh when returning via keep-alive/navigation (if enabled later)
onActivated(() => {
  fetchBrands()
})

// Stats computeds
const publishedBrands = computed(() => brands.value.filter(b => b.published))
const unpublishedBrands = computed(() => brands.value.filter(b => !b.published))
const brandsWithImage = computed(() => brands.value.filter(b => b.image_url && b.image_url.trim() !== ''))

// Delete brand helper
async function deleteBrand(b) {
  const ok = confirm(`آیا از حذف برند «${b.name}» مطمئن هستید؟`)
  if (!ok) return

  try {
    await $fetch(`/api/brands/${b.id}`, { method: 'DELETE' })
    // Refresh list after deletion
    await fetchBrands()
  } catch (err) {
    console.error('خطا در حذف برند:', err)
    alert('خطا در حذف برند')
  }
}

// search term and computed filtered list
const searchTerm = ref('')

const filteredBrands = computed(() => {
  if (!searchTerm.value.trim()) return brands.value
  const term = searchTerm.value.toLowerCase()
  return brands.value.filter(b => {
    const n1 = (b.name || '').toLowerCase()
    const n2 = (b.official_name || '').toLowerCase()
    return n1.includes(term) || n2.includes(term)
  })
})

// order editing reactive vars
const editingOrderId = ref(null)
const editingOrderValue = ref('')

function startEditOrder(b) {
  editingOrderId.value = b.id
  editingOrderValue.value = b.order ? String(b.order) : ''
}

async function finishEditOrder(b) {
  const newOrder = parseInt(editingOrderValue.value)
  if (!isNaN(newOrder) && newOrder >= 1 && newOrder <= filteredBrands.value.length && newOrder !== b.order) {
    const oldIndex = brands.value.findIndex(x => x.id === b.id)
    const moved = brands.value.splice(oldIndex,1)[0]
    brands.value.splice(newOrder-1,0,moved)
    brands.value.forEach((x, idx) => { x.order = idx+1 })
    // TODO: Send to backend when endpoint ready
  }
  editingOrderId.value=null
  editingOrderValue.value=''
}

function handleOrderInputKey(e,b){ if(e.key==='Enter') finishEditOrder(b); else if(e.key==='Escape'){ editingOrderId.value=null; editingOrderValue.value=''} }

// Auth disabled
// eslint-disable-next-line @typescript-eslint/no-explicit-any
const confirmDialogRef = ref<any>(null)
provide('confirmDialogRef', confirmDialogRef)

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})
</script>

<style scoped>
/* Add any specific styles here */
</style> 
