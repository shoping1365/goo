<template>
  <div class="min-h-screen">
    <!-- هدر صفحه -->
    <div class="header-bg">
      <div class="page-header-flex">
        <h1 class="page-title">مدیریت بنرها و ابزارک‌ها</h1>
        <button class="btn btn-primary new-item-btn" @click="addNewWidget">ایجاد ابزارک</button>
      </div>
    </div>

    <!-- لودینگ -->
    <div v-if="loading" class="flex justify-center items-center py-12">
      <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
    </div>

    <!-- خطا -->
    <div v-if="error" class="mx-6 mt-4 p-6 bg-red-50 border border-red-200 rounded-lg">
      <p class="text-red-800">{{ error }}</p>
      <button 
        class="mt-2 text-red-600 hover:text-red-800"
        @click="clearError"
      >
        بستن
      </button>
    </div>

    <!-- جدول ابزارک‌ها -->
    <div class="p-6">
      <div class="bg-white rounded-xl shadow-lg overflow-hidden">
        <!-- منو پایین -->
        <div class="px-6 py-4 bg-gray-50 border-b border-gray-200">
          <!-- نمایش صفحه انتخاب شده -->
          <div class="mb-4 text-sm text-gray-600">
            <span>صفحه فعلی: </span>
            <span class="font-medium text-gray-900">
              {{ selectedPage === 'home' ? 'صفحه اصلی' :
                 selectedPage === 'other' ? 'سایر بخش‌ها' : 'نامشخص' }}
            </span>
          </div>
          <div class="flex gap-3">
            <button 
              :class="[
                'inline-flex items-center font-medium rounded-lg focus:outline-none transition-all duration-200 shadow-md px-5 py-3 text-sm',
                selectedPage === 'home'
                  ? 'bg-gradient-to-r from-blue-400 to-blue-600 text-white border border-blue-500 hover:from-blue-500 hover:to-blue-700'
                  : 'bg-gradient-to-r from-gray-100 to-gray-200 text-gray-700 border border-gray-300 hover:from-gray-200 hover:to-gray-300'
              ]"
              @click="fetchPageWidgets('home')"
            >
              صفحه اصلی
            </button>
            <button
              :class="[
                'inline-flex items-center font-medium rounded-lg focus:outline-none transition-all duration-200 shadow-md px-5 py-3 text-sm',
                selectedPage === 'other'
                  ? 'bg-gradient-to-r from-blue-400 to-blue-600 text-white border border-blue-500 hover:from-blue-500 hover:to-blue-700'
                  : 'bg-gradient-to-r from-gray-100 to-gray-200 text-gray-700 border border-gray-300 hover:from-gray-200 hover:to-gray-300'
              ]"
              @click="fetchOtherPagesWidgets"
            >
              سایر بخش‌ها
            </button>
          </div>
        </div>
        <!-- فیلتر و جستجو -->
        <div class="px-6 py-4 bg-gray-50 border-b border-gray-200">
          <div class="flex flex-col md:flex-row gap-6 items-center">
            <div class="flex-1">
              <input 
                v-model="searchQuery"
                type="text"
                placeholder="جستجو در ابزارک‌ها..."
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              />
            </div>
            <div class="flex gap-2">
              <select 
                v-model="statusFilter"
                class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
                <option value="">همه وضعیت‌ها</option>
                <option value="active">فعال</option>
                <option value="inactive">غیرفعال</option>
                <option value="draft">پیش‌نویس</option>
              </select>
              <select 
                v-model="typeFilter"
                class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
                <option value="">همه انواع</option>
                <option v-for="(label, type) in WIDGET_TYPE_LABELS" :key="type" :value="type">
                  {{ label }}
                </option>
              </select>
            </div>
          </div>
        </div>
        <!-- هدر جدول -->
        <div class="bg-gradient-to-r from-blue-50 to-indigo-50 px-6 py-4 border-b border-gray-200">
          <div class="grid grid-cols-12 gap-6 text-sm font-bold text-gray-700" dir="rtl">
            <span class="col-span-1 text-center">جابجایی</span>
            <span class="col-span-2 text-right">عنوان</span>
            <span class="col-span-1 text-center">کد</span>
            <span class="col-span-2 text-center">نوع</span>
            <span class="col-span-2 text-center">وضعیت</span>
            <span class="col-span-2 text-center">نمایش در موبایل</span>
            <span class="col-span-2 text-center">عملیات</span>
          </div>
        </div>

        <!-- لیست ابزارک‌ها -->
        <div class="divide-y divide-gray-200">
          <div
            v-for="(widget, idx) in filteredWidgets"
            :key="widget.id"
            class="px-6 py-4 hover:bg-gray-50 transition-colors duration-200"
            draggable="true"
            @dragstart="onDragStart(Number(idx))"
            @dragover.prevent
            @drop="onDrop(Number(idx))"
          >
            <div class="grid grid-cols-12 gap-6 items-center" dir="rtl">
              <!-- آیکون جابجایی -->
              <div class="col-span-1 flex justify-center">
                <div class="cursor-grab text-gray-400 hover:text-gray-600 transition-colors duration-300">
                  <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
                    <path d="M3 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 10a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 16a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z"></path>
                  </svg>
                </div>
              </div>

              <!-- عنوان -->
              <div class="col-span-2">
                <div class="font-semibold text-gray-900">{{ widget.title }}</div>
                <div class="text-sm text-gray-500 mt-1">{{ widget.description }}</div>
              </div>

              <!-- کد -->
              <div class="col-span-1 flex justify-center">
                <span class="px-3 py-1 text-xs font-semibold rounded-full bg-gray-100 text-gray-800 font-mono">
                  {{ widget.code }}
                </span>
              </div>

              <!-- نوع ابزارک -->
              <div class="col-span-2 flex justify-center">
                <span class="px-3 py-1 text-xs font-semibold rounded-full" :class="getTypeBadgeClass(widget.type)">
                  {{ getWidgetTypeLabel(widget.type) }}
                </span>
              </div>

              <!-- وضعیت -->
              <div class="col-span-2 flex justify-center">
                <button 
                  class="px-3 py-1 text-xs font-semibold rounded-full transition-colors"
                  :class="getStatusBadgeClass(widget.status)" 
                  @click="toggleStatus(widget.id)"
                >
                  {{ getWidgetStatusLabel(widget.status) }}
                </button>
              </div>

              <!-- نمایش در موبایل -->
              <div class="col-span-2 flex justify-center">
                <span class="px-3 py-1 text-xs font-semibold rounded-full" :class="getMobileDisplayBadgeClass(widget.show_on_mobile)">
                  {{ getMobileDisplayLabel(widget.show_on_mobile) }}
                </span>
              </div>

              <!-- عملیات -->
              <div class="col-span-2 flex justify-center gap-2">
                <button 
                  v-if="hasPermission('widget.update')"
                  class="w-8 h-8 bg-blue-100 hover:bg-blue-200 text-blue-600 rounded-lg transition-colors duration-200 flex items-center justify-center"
                  @click="editWidget(widget.id)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
                  </svg>
                </button>
                <button 
                  v-if="hasPermission('widget.update')"
                  class="w-8 h-8 bg-green-100 hover:bg-green-200 text-green-600 rounded-lg transition-colors duration-200 flex items-center justify-center"
                  @click="duplicateWidget(widget.id)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
                  </svg>
                </button>
                <button 
                  v-if="hasPermission('widget.delete')"
                  class="w-8 h-8 bg-red-100 hover:bg-red-200 text-red-600 rounded-lg transition-colors duration-200 flex items-center justify-center"
                  @click="confirmDelete(widget.id)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>

        <!-- پیام خالی بودن -->
        <div v-if="!widgets.length && !loading" class="text-center py-12">
          <div class="w-24 h-24 mx-auto mb-4 bg-gray-100 rounded-full flex items-center justify-center">
            <svg class="w-12 h-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-800 mb-2">هنوز ابزارکی ایجاد نشده</h3>
          <p class="text-gray-600 mb-4">برای شروع، اولین ابزارک خود را ایجاد کنید</p>
                     <NuxtLink 
             v-if="hasPermission('widget.create')"
             to="/admin/content/banners/create" 
             class="bg-blue-600 hover:bg-blue-700 text-white px-6 py-3 rounded-lg font-semibold transition-colors duration-200 inline-flex items-center gap-2"
           >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
            </svg>
            ایجاد ابزارک
          </NuxtLink>
        </div>
      </div>
    </div>


  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuth } from '~/composables/useAuth'
import { useWidget } from '~/composables/useWidget'
import type { WidgetStatus, WidgetType, WidgetPage } from '~/types/widget'
import { WIDGET_TYPE_LABELS, getWidgetStatusLabel, getWidgetTypeLabel } from '~/types/widget'

// تعریف definePageMeta برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void

definePageMeta({
  layout: 'admin-main'
})

const router = useRouter()

// استفاده از composable
const { 
  widgets, 
  loading, 
  error, 
  fetchWidgets: _fetchWidgets, 
  fetchWidgetsByPage,
  updateWidgetOrder, 
  toggleWidgetStatus, 
  deleteWidget,
  duplicateWidget,
  clearError 
} = useWidget()

// استفاده از composable auth
const { hasPermission } = useAuth()

// State
const selectedPage = ref<string>('home')
const searchQuery = ref('')
const statusFilter = ref<WidgetStatus | ''>('')
const typeFilter = ref<WidgetType | ''>('')
const _showAddPageModal = ref(false)

// Drag & Drop
const dragIdx = ref<number | null>(null)

// Computed
const filteredWidgets = computed(() => {
  let filtered = widgets.value

  if (searchQuery.value) {
    const search = searchQuery.value.toLowerCase()
    filtered = filtered.filter(w => 
      w.title.toLowerCase().includes(search) ||
      w.description?.toLowerCase().includes(search) ||
      w.code.toLowerCase().includes(search)
    )
  }

  if (statusFilter.value) {
    filtered = filtered.filter(w => w.status === statusFilter.value)
  }

  if (typeFilter.value) {
    filtered = filtered.filter(w => w.type === typeFilter.value)
  }

  return filtered
})

// Methods
const fetchPageWidgets = async (page: string) => {
  selectedPage.value = page
  await fetchWidgetsByPage(page as WidgetPage)
}

const fetchOtherPagesWidgets = async () => {
  selectedPage.value = 'other'
  await fetchWidgetsByPage('other' as WidgetPage)
}

const _updateOrder = async (id: number, event: Event) => {
  const target = event.target as HTMLInputElement
  const newOrder = parseInt(target.value)
  
  if (isNaN(newOrder) || newOrder < 1) return
  
  await updateWidgetOrder({
    items: [{ id, order: newOrder }]
  })
}

const toggleStatus = async (id: number) => {
  try {
    await toggleWidgetStatus(id)
  } catch (error) {
    console.error('خطا در تغییر وضعیت:', error)
  }
}

const editWidget = (id: number) => {
  const widget = widgets.value.find(w => w.id === id)
  if (widget?.type === 'main-slider-side-banner') {
    router.push(`/admin/content/banners/MainSliderSideBanner/edit/${id}`)
  } else if (widget?.type === 'single-slider-side') {
    router.push(`/admin/content/banners/SingleSliderSide/edit/${id}`)
  } else if (widget?.type === 'full-slider') {
    router.push(`/admin/content/banners/full-slider/edit/${id}`)
  } else if (widget?.type === 'full-banner') {
    router.push(`/admin/content/banners/FullBanner/edit/${id}`)
  } else if (widget?.type === 'double-banner') {
    router.push(`/admin/content/banners/DoubleBanner/edit/${id}`)
  } else if (widget?.type === 'triple-banner') {
    router.push(`/admin/content/banners/TripleBanner/edit/${id}`)
  } else if (widget?.type === 'quad-banner') {
    router.push(`/admin/content/banners/QuadBanner/edit/${id}`)
  } else if (widget?.type === 'penta-banner') {
    router.push(`/admin/content/banners/PentaBanner/edit/${id}`)
  } else if (widget?.type === 'brands-slider') {
    router.push(`/admin/content/banners/BrandsSlider/edit/${id}`)
  } else if (widget?.type === 'services-slider') {
    router.push(`/admin/content/banners/ServicesSlider/edit/${id}`)
  } else if (widget?.type === 'category-box') {
    router.push(`/admin/content/banners/CategoryBox/edit/${id}`)
  } else if (widget?.type === 'blog-posts') {
    router.push(`/admin/content/banners/BlogPosts/edit/${id}`)
  } else if (widget?.type === 'product-carousel') {
    router.push(`/admin/content/banners/ProductCarousel/edit/${id}`)
  } else {
    router.push(`/admin/content/banners/${id}`)
  }
}

const confirmDelete = (id: number) => {
  if (confirm('آیا از حذف این ابزارک اطمینان دارید؟')) {
    deleteWidget(id)
  }
}

// Drag & Drop handlers
const onDragStart = (index: number) => {
  dragIdx.value = index
}

const onDrop = async (dropIndex: number) => {
  if (dragIdx.value === null) return
  
  const draggedWidget = filteredWidgets.value[dragIdx.value]
  const targetWidget = filteredWidgets.value[dropIndex]
  
  if (!draggedWidget || !targetWidget) return
  
  // Find the actual indices in the original widgets array
  const draggedIndex = widgets.value.findIndex(w => w.id === draggedWidget.id)
  const targetIndex = widgets.value.findIndex(w => w.id === targetWidget.id)
  
  if (draggedIndex === -1 || targetIndex === -1) return
  
  // Swap orders
  await updateWidgetOrder({
    items: [
      { id: draggedWidget.id, order: targetWidget.order },
      { id: targetWidget.id, order: draggedWidget.order }
    ]
  })
  
  // Refresh the widgets list to reflect the changes
  await fetchPageWidgets(selectedPage.value)
  
  dragIdx.value = null
}

// Style helpers
const getTypeBadgeClass = (type: WidgetType): string => {
  const classes = {
    'main-slider-side-banner': 'bg-blue-100 text-blue-800',
    'single-slider-side': 'bg-blue-100 text-blue-800',

    'products-box-bg': 'bg-purple-100 text-purple-800',
    'full-banner': 'bg-green-100 text-green-800',

    'double-banner': 'bg-green-100 text-green-800',
    'triple-banner': 'bg-green-100 text-green-800',
    'quad-banner': 'bg-green-100 text-green-800',
    'penta-banner': 'bg-green-100 text-green-800',
    'brands-slider': 'bg-orange-100 text-orange-800',
    'services-slider': 'bg-orange-100 text-orange-800',
    'category-box': 'bg-indigo-100 text-indigo-800',
    'blog-posts': 'bg-pink-100 text-pink-800'
  }
  return classes[type] || 'bg-gray-100 text-gray-800'
}

const getStatusBadgeClass = (status: WidgetStatus): string => {
  const classes = {
    active: 'bg-green-100 text-green-800 hover:bg-green-200',
    inactive: 'bg-red-100 text-red-800 hover:bg-red-200',
    draft: 'bg-yellow-100 text-yellow-800 hover:bg-yellow-200'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const getMobileDisplayBadgeClass = (showOnMobile: boolean): string => {
  const classes = {
    true: 'bg-green-100 text-green-800 hover:bg-green-200',
    false: 'bg-red-100 text-red-800 hover:bg-red-200'
  }
  return classes[showOnMobile ? 'true' : 'false'] || 'bg-gray-100 text-gray-800'
}

const getMobileDisplayLabel = (showOnMobile: boolean): string => {
  return showOnMobile ? 'نمایش در موبایل' : 'عدم نمایش در موبایل'
}

// افزودن ابزارک جدید
const addNewWidget = () => {
  router.push('/admin/content/banners/create')
}

// Initialize
onMounted(() => {
  fetchPageWidgets('home')
})
</script>

<style scoped>
.header-bg {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  padding: 30px;
  border-radius: 12px;
  margin-bottom: 30px;
}

.page-header-flex {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.page-title {
  font-size: 2rem;
  font-weight: bold;
  margin: 0;
}

.new-item-btn {
  background: rgba(255, 255, 255, 0.2);
  border: 2px solid rgba(255, 255, 255, 0.3);
  color: white;
  padding: 12px 24px;
  border-radius: 8px;
  font-weight: 600;
  transition: all 0.3s ease;
}

.new-item-btn:hover {
  background: rgba(255, 255, 255, 0.3);
  border-color: rgba(255, 255, 255, 0.5);
}
</style> 
