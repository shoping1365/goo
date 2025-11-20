<template>
  <div class="brands-slider-management-page p-6">
    <div class="max-w-7xl mx-auto">
      <!-- Page Header -->
      <div class="mb-8">
        <div class="flex items-center justify-between">
          <div>
            <h1 class="text-3xl font-bold text-gray-900">مدیریت اسلایدر برندها</h1>
            <p class="text-gray-600 mt-1">مدیریت و تنظیم ابزارک‌های اسلایدر برندها</p>
          </div>
          <div class="flex gap-3">
            <NuxtLink 
              to="/admin/content/banners/BrandsSlider/create"
              class="bg-green-600 hover:bg-green-700 text-white px-6 py-3 rounded-lg font-semibold transition-colors duration-200 flex items-center gap-2"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/>
              </svg>
              ایجاد اسلایدر جدید
            </NuxtLink>
            <NuxtLink 
              to="/admin/content/banners"
              class="bg-gray-600 hover:bg-gray-700 text-white px-6 py-3 rounded-lg font-semibold transition-colors duration-200 flex items-center gap-2"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"/>
              </svg>
              بازگشت به ابزارک‌ها
            </NuxtLink>
          </div>
        </div>
      </div>

      <!-- Filters and Search -->
      <div class="bg-white rounded-lg shadow-sm p-6 mb-6">
        <div class="flex flex-col md:flex-row gap-6 items-center">
          <div class="flex-1">
            <input 
              v-model="searchQuery"
              type="text"
              placeholder="جستجو در اسلایدرها..."
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
              v-model="pageFilter"
              class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
              <option value="">همه صفحات</option>
              <option value="home">صفحه اصلی</option>
              <option value="category">دسته‌بندی</option>
              <option value="product">محصول</option>
              <option value="brands">برندها</option>
            </select>
          </div>
        </div>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="flex justify-center items-center py-12">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
      </div>

      <!-- Error State -->
      <div v-else-if="error" class="bg-red-50 border border-red-200 rounded-lg p-6 mb-6">
        <p class="text-red-800">{{ error }}</p>
        <button 
          class="mt-2 text-red-600 hover:text-red-800"
          @click="clearError"
        >
          بستن
        </button>
      </div>

      <!-- Sliders List -->
      <div v-else-if="filteredSliders.length > 0" class="bg-white rounded-xl shadow-lg overflow-hidden">
        <!-- Table Header -->
        <div class="bg-gradient-to-r from-blue-50 to-indigo-50 px-6 py-4 border-b border-gray-200">
          <div class="grid grid-cols-12 gap-6 text-sm font-bold text-gray-700">
            <span class="col-span-1 text-center">ترتیب</span>
            <span class="col-span-1 text-center">تصویر</span>
            <span class="col-span-2">عنوان</span>
            <span class="col-span-1 text-center">کد</span>
            <span class="col-span-1 text-center">صفحه</span>
            <span class="col-span-1 text-center">وضعیت</span>
            <span class="col-span-2 text-center">تنظیمات</span>
            <span class="col-span-3 text-center">عملیات</span>
          </div>
        </div>

        <!-- Sliders List -->
        <div class="divide-y divide-gray-200">
          <div
            v-for="(slider, idx) in filteredSliders"
            :key="slider.id"
            class="px-6 py-4 hover:bg-gray-50 transition-colors duration-200"
          >
            <div class="grid grid-cols-12 gap-6 items-center">
              <!-- Order -->
              <div class="col-span-1 flex justify-center">
                <input 
                  type="number" 
                  :value="slider.order" 
                  min="1" 
                  class="w-16 h-8 bg-white border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 p-2 text-center font-semibold" 
                  @change="updateOrder(slider.id, $event)"
                />
              </div>

              <!-- Image -->
              <div class="col-span-1 flex justify-center">
                <div class="w-12 h-12 bg-gray-100 rounded-lg border border-gray-300 overflow-hidden">
                  <img 
                    v-if="slider.image"
                    :src="slider.image" 
                    :alt="slider.title"
                    class="w-full h-full object-cover"
                  />
                  <div v-else class="w-full h-full bg-gray-100 flex items-center justify-center">
                    <svg class="w-6 h-6 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"/>
                    </svg>
                  </div>
                </div>
              </div>

              <!-- Title -->
              <div class="col-span-2">
                <div class="font-semibold text-gray-900">{{ slider.title }}</div>
                <div class="text-sm text-gray-500 mt-1">{{ slider.description }}</div>
              </div>

              <!-- Code -->
              <div class="col-span-1 flex justify-center">
                <span class="px-3 py-1 text-xs font-semibold rounded-full bg-gray-100 text-gray-800 font-mono">
                  {{ slider.code }}
                </span>
              </div>

              <!-- Page -->
              <div class="col-span-1 flex justify-center">
                <span class="px-3 py-1 text-xs font-semibold rounded-full bg-blue-100 text-blue-800">
                  {{ slider.page }}
                </span>
              </div>

              <!-- Status -->
              <div class="col-span-1 flex justify-center">
                <button 
                  class="px-3 py-1 text-xs font-semibold rounded-full transition-colors"
                  :class="getStatusBadgeClass(slider.status)" 
                  @click="toggleStatus(slider.id)"
                >
                  {{ getStatusLabel(slider.status) }}
                </button>
              </div>

              <!-- Configuration -->
              <div class="col-span-2 text-center text-sm text-gray-600">
                <div>{{ slider.slidesPerView }} برند</div>
                <div v-if="slider.autoplay" class="text-green-600">پخش خودکار</div>
              </div>

              <!-- Actions -->
              <div class="col-span-3 flex justify-center gap-2">
                <NuxtLink 
                  :to="`/admin/content/banners/BrandsSlider/edit/${slider.id}`"
                  class="w-8 h-8 bg-blue-100 hover:bg-blue-200 text-blue-600 rounded-lg transition-colors duration-200 flex items-center justify-center"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"/>
                  </svg>
                </NuxtLink>
                <button 
                  class="w-8 h-8 bg-green-100 hover:bg-green-200 text-green-600 rounded-lg transition-colors duration-200 flex items-center justify-center"
                  @click="duplicateSlider(slider.id)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"/>
                  </svg>
                </button>
                <button 
                  class="w-8 h-8 bg-red-100 hover:bg-red-200 text-red-600 rounded-lg transition-colors duration-200 flex items-center justify-center"
                  @click="confirmDelete(slider.id)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"/>
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Empty State -->
      <div v-else-if="!loading" class="text-center py-12 bg-white rounded-lg shadow-sm">
        <div class="w-24 h-24 mx-auto mb-4 bg-gray-100 rounded-full flex items-center justify-center">
          <svg class="w-12 h-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"/>
          </svg>
        </div>
        <h3 class="text-lg font-semibold text-gray-800 mb-2">هنوز اسلایدر برندی ایجاد نشده</h3>
        <p class="text-gray-600 mb-4">با ایجاد اولین اسلایدر برند شروع کنید</p>
        <NuxtLink 
          to="/admin/content/banners/BrandsSlider/create"
          class="bg-blue-600 hover:bg-blue-700 text-white px-6 py-3 rounded-lg font-semibold transition-colors duration-200 inline-flex items-center gap-2"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/>
          </svg>
          ایجاد اولین اسلایدر
        </NuxtLink>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useAuth } from '~/composables/useAuth'

// تعریف definePageMeta و useHead برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
declare const useHead: (head: { title?: string }) => void

// Page meta definition
definePageMeta({
  layout: 'admin',
  middleware: 'admin'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user, hasPermission } = useAuth()

// Page title
useHead({
  title: 'مدیریت اسلایدر برندها - پنل ادمین'
})

// Mock data for demonstration
const sliders = ref([
  {
    id: 1,
    title: 'اسلایدر برندهای محبوب',
    description: 'نمایش برندهای محبوب و پرطرفدار',
    code: 'popular-brands',
    page: 'home',
    status: 'active',
    order: 1,
    image: null,
    slidesPerView: 6,
    autoplay: true
  },
  {
    id: 2,
    title: 'اسلایدر برندهای دسته‌بندی',
    description: 'برندهای مربوط به دسته‌بندی خاص',
    code: 'category-brands',
    page: 'category',
    status: 'active',
    order: 2,
    image: null,
    slidesPerView: 4,
    autoplay: false
  }
])

// State
const loading = ref(false)
const error = ref('')
const searchQuery = ref('')
const statusFilter = ref('')
const pageFilter = ref('')

// Computed
const filteredSliders = computed(() => {
  let filtered = sliders.value

  if (searchQuery.value) {
    const search = searchQuery.value.toLowerCase()
    filtered = filtered.filter(s => 
      s.title.toLowerCase().includes(search) ||
      s.description.toLowerCase().includes(search) ||
      s.code.toLowerCase().includes(search)
    )
  }

  if (statusFilter.value) {
    filtered = filtered.filter(s => s.status === statusFilter.value)
  }

  if (pageFilter.value) {
    filtered = filtered.filter(s => s.page === pageFilter.value)
  }

  return filtered.sort((a, b) => a.order - b.order)
})

// Methods
const updateOrder = async (id: number, event: Event) => {
  const target = event.target as HTMLInputElement
  const newOrder = parseInt(target.value)
  
  if (isNaN(newOrder) || newOrder < 1) return
  
  const slider = sliders.value.find(s => s.id === id)
  if (slider) {
    slider.order = newOrder
  }
  
  // Here you would call your API to update the order
  console.log('Updating order for slider:', id, 'to:', newOrder)
}

const toggleStatus = async (id: number) => {
  const slider = sliders.value.find(s => s.id === id)
  if (slider) {
    slider.status = slider.status === 'active' ? 'inactive' : 'active'
  }
  
  // Here you would call your API to update the status
  console.log('Toggling status for slider:', id)
}

const duplicateSlider = async (id: number) => {
  const slider = sliders.value.find(s => s.id === id)
  if (slider) {
    const newSlider = {
      ...slider,
      id: Date.now(),
      title: `${slider.title} (کپی)`,
      code: `${slider.code}-copy`,
      status: 'draft' as const,
      order: sliders.value.length + 1
    }
    sliders.value.push(newSlider)
  }
  
  // Here you would call your API to duplicate the slider
  console.log('Duplicating slider:', id)
}

const confirmDelete = async (id: number) => {
  if (confirm('آیا مطمئن هستید که می‌خواهید این اسلایدر را حذف کنید؟')) {
    sliders.value = sliders.value.filter(s => s.id !== id)
    
    // Here you would call your API to delete the slider
    console.log('Deleting slider:', id)
  }
}

const clearError = () => {
  error.value = ''
}

const getStatusLabel = (status: string): string => {
  const labels = {
    active: 'فعال',
    inactive: 'غیرفعال',
    draft: 'پیش‌نویس'
  }
  return labels[status] || status
}

const getStatusBadgeClass = (status: string): string => {
  const classes = {
    active: 'bg-green-100 text-green-800 hover:bg-green-200',
    inactive: 'bg-red-100 text-red-800 hover:bg-red-200',
    draft: 'bg-yellow-100 text-yellow-800 hover:bg-yellow-200'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}
</script>

<style scoped>
.brands-slider-management-page {
  background-color: #f9fafb;
  min-height: 100vh;
}
</style>


