<template>
  <div class="services-slider-management-page p-6">
    <div class="max-w-7xl mx-auto">
      <!-- Page Header -->
      <div class="mb-8">
        <div class="flex items-center justify-between mb-4">
          <div>
            <h1 class="text-3xl font-bold text-gray-900">مدیریت اسلایدرهای خدمات</h1>
            <p class="text-gray-600">مدیریت و تنظیم اسلایدرهای خدمات</p>
          </div>
          <NuxtLink 
            to="/admin/content/banners/ServicesSlider/create"
            class="bg-blue-600 hover:bg-blue-700 text-white px-6 py-3 rounded-lg font-semibold transition-colors duration-200 inline-flex items-center gap-2"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/>
            </svg>
            ایجاد اسلایدر جدید
          </NuxtLink>
        </div>
      </div>

      <!-- Filters and Search -->
      <div class="bg-white rounded-lg shadow-sm p-6 mb-6">
        <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">جستجو</label>
            <input 
              v-model="filters.search"
              type="text"
              placeholder="جستجو در عنوان یا کد..."
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
            <select 
              v-model="filters.status"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
              <option value="">همه</option>
              <option value="active">فعال</option>
              <option value="inactive">غیرفعال</option>
              <option value="draft">پیش‌نویس</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">صفحه</label>
            <select 
              v-model="filters.page"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
              <option value="">همه صفحات</option>
              <option value="home">صفحه اصلی</option>
              <option value="about">درباره ما</option>
              <option value="services">صفحات خدمات</option>
              <option value="contact">تماس با ما</option>
            </select>
          </div>
          <div class="flex items-end">
            <button 
              class="w-full px-4 py-2 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors"
              @click="clearFilters"
            >
              پاک کردن فیلترها
            </button>
          </div>
        </div>
      </div>

      <!-- Services Sliders List -->
      <div class="bg-white rounded-lg shadow-sm overflow-hidden">
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                  ترتیب
                </th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                  پیش‌نمایش
                </th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                  عنوان
                </th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                  کد
                </th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                  صفحه
                </th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                  وضعیت
                </th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                  تنظیمات
                </th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                  عملیات
                </th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="slider in filteredSliders" :key="slider.id" class="hover:bg-gray-50">
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="flex items-center gap-2">
                    <button 
                      :disabled="slider.order === 1"
                      class="text-gray-400 hover:text-gray-600 disabled:opacity-50"
                      @click="moveUp(slider.id)"
                    >
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 15l7-7 7 7"/>
                      </svg>
                    </button>
                    <span class="text-sm font-medium text-gray-900">{{ slider.order }}</span>
                    <button 
                      :disabled="slider.order === sliders.length"
                      class="text-gray-400 hover:text-gray-600 disabled:opacity-50"
                      @click="moveDown(slider.id)"
                    >
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/>
                      </svg>
                    </button>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="w-16 h-16 bg-gray-100 rounded-lg flex items-center justify-center">
                    <svg class="w-8 h-8 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
                    </svg>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <div>
                    <div class="text-sm font-medium text-gray-900">{{ slider.title }}</div>
                    <div class="text-sm text-gray-500">{{ slider.servicesCount }} خدمت</div>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <code class="text-sm bg-gray-100 px-2 py-1 rounded">{{ slider.code }}</code>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800">
                    {{ getPageName(slider.page) }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span 
                    :class="[
                      'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium',
                      slider.status === 'active' ? 'bg-green-100 text-green-800' : 
                      slider.status === 'inactive' ? 'bg-red-100 text-red-800' : 
                      'bg-yellow-100 text-yellow-800'
                    ]"
                  >
                    {{ getStatusName(slider.status) }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  <div class="space-y-1">
                    <div>نمایش: {{ slider.slidesPerView }}</div>
                    <div>ارتفاع: {{ slider.carouselHeight }}px</div>
                    <div v-if="slider.autoplay">پخش خودکار: {{ slider.autoplayDelay }}ms</div>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                  <div class="flex items-center gap-2">
                    <button 
                      :class="[
                        'px-3 py-1 rounded-md text-xs font-medium transition-colors',
                        slider.status === 'active' 
                          ? 'bg-red-100 text-red-800 hover:bg-red-200' 
                          : 'bg-green-100 text-green-800 hover:bg-green-200'
                      ]"
                      @click="toggleStatus(slider.id)"
                    >
                      {{ slider.status === 'active' ? 'غیرفعال' : 'فعال' }}
                    </button>
                    <NuxtLink 
                      :to="`/admin/content/banners/ServicesSlider/edit/${slider.id}`"
                      class="px-3 py-1 bg-blue-100 text-blue-800 rounded-md text-xs font-medium hover:bg-blue-200 transition-colors"
                    >
                      ویرایش
                    </NuxtLink>
                    <button 
                      class="px-3 py-1 bg-purple-100 text-purple-800 rounded-md text-xs font-medium hover:bg-purple-200 transition-colors"
                      @click="duplicateSlider(slider.id)"
                    >
                      تکرار
                    </button>
                    <button 
                      class="px-3 py-1 bg-red-100 text-red-800 rounded-md text-xs font-medium hover:bg-red-200 transition-colors"
                      @click="deleteSlider(slider.id)"
                    >
                      حذف
                    </button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <!-- Empty State -->
        <div v-if="filteredSliders.length === 0" class="text-center py-12">
          <div class="w-24 h-24 mx-auto mb-4 bg-gray-100 rounded-full flex items-center justify-center">
            <svg class="w-12 h-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-800 mb-2">هیچ اسلایدری یافت نشد</h3>
          <p class="text-gray-600 mb-4">
            {{ filters.search || filters.status || filters.page ? 'فیلترهای خود را تغییر دهید یا' : '' }}
            اولین اسلایدر خدمات خود را ایجاد کنید
          </p>
          <NuxtLink 
            to="/admin/content/banners/ServicesSlider/create"
            class="bg-blue-600 hover:bg-blue-700 text-white px-6 py-3 rounded-lg font-semibold transition-colors duration-200 inline-flex items-center gap-2"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/>
            </svg>
            ایجاد اسلایدر جدید
          </NuxtLink>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// Page meta definition
definePageMeta({
  layout: 'admin',
  middleware: 'admin'
})

import { ref, computed } from 'vue'

// تعریف definePageMeta و useHead برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
declare const useHead: (head: { title?: string }) => void

// Page title
useHead({
  title: 'مدیریت اسلایدرهای خدمات - پنل ادمین'
})

// State
const filters = ref({
  search: '',
  status: '',
  page: ''
})

// Mock data for demonstration
const sliders = ref([
  {
    id: '1',
    title: 'اسلایدر خدمات اصلی',
    code: 'main-services',
    page: 'home',
    order: 1,
    status: 'active',
    slidesPerView: 3,
    carouselHeight: 200,
    autoplay: true,
    autoplayDelay: 5000,
    servicesCount: 6
  },
  {
    id: '2',
    title: 'اسلایدر خدمات ویژه',
    code: 'featured-services',
    page: 'services',
    order: 2,
    status: 'active',
    slidesPerView: 4,
    carouselHeight: 180,
    autoplay: false,
    autoplayDelay: 0,
    servicesCount: 8
  },
  {
    id: '3',
    title: 'اسلایدر خدمات تماس',
    code: 'contact-services',
    page: 'contact',
    order: 3,
    status: 'draft',
    slidesPerView: 2,
    carouselHeight: 250,
    autoplay: true,
    autoplayDelay: 7000,
    servicesCount: 4
  }
])

// Computed
const filteredSliders = computed(() => {
  return sliders.value.filter(slider => {
    const matchesSearch = !filters.value.search || 
      slider.title.toLowerCase().includes(filters.value.search.toLowerCase()) ||
      slider.code.toLowerCase().includes(filters.value.search.toLowerCase())
    
    const matchesStatus = !filters.value.status || slider.status === filters.value.status
    const matchesPage = !filters.value.page || slider.page === filters.value.page
    
    return matchesSearch && matchesStatus && matchesPage
  })
})

// Methods
const getPageName = (page: string) => {
  const pageNames: Record<string, string> = {
    'home': 'صفحه اصلی',
    'about': 'درباره ما',
    'services': 'صفحات خدمات',
    'contact': 'تماس با ما'
  }
  return pageNames[page] || page
}

const getStatusName = (status: string) => {
  const statusNames: Record<string, string> = {
    'active': 'فعال',
    'inactive': 'غیرفعال',
    'draft': 'پیش‌نویس'
  }
  return statusNames[status] || status
}

const clearFilters = () => {
  filters.value = {
    search: '',
    status: '',
    page: ''
  }
}

const moveUp = (id: string) => {
  const index = sliders.value.findIndex(s => s.id === id)
  if (index > 0) {
    const temp = sliders.value[index]
    sliders.value[index] = sliders.value[index - 1]
    sliders.value[index - 1] = temp
    
    // Update order
    sliders.value[index].order = index + 1
    sliders.value[index - 1].order = index
  }
}

const moveDown = (id: string) => {
  const index = sliders.value.findIndex(s => s.id === id)
  if (index < sliders.value.length - 1) {
    const temp = sliders.value[index]
    sliders.value[index] = sliders.value[index + 1]
    sliders.value[index + 1] = temp
    
    // Update order
    sliders.value[index].order = index + 1
    sliders.value[index + 1].order = index + 2
  }
}

const toggleStatus = (id: string) => {
  const slider = sliders.value.find(s => s.id === id)
  if (slider) {
    slider.status = slider.status === 'active' ? 'inactive' : 'active'
  }
}

const duplicateSlider = (id: string) => {
  const original = sliders.value.find(s => s.id === id)
  if (original) {
    const duplicated = {
      ...original,
      id: Date.now().toString(),
      title: `${original.title} (کپی)`,
      code: `${original.code}-copy`,
      order: sliders.value.length + 1,
      status: 'draft' as const
    }
    sliders.value.push(duplicated)
    
    // Reorder all items
    sliders.value.forEach((item, index) => {
      item.order = index + 1
    })
  }
}

const deleteSlider = (id: string) => {
  if (confirm('آیا از حذف این اسلایدر اطمینان دارید؟')) {
    const index = sliders.value.findIndex(s => s.id === id)
    if (index > -1) {
      sliders.value.splice(index, 1)
      
      // Reorder remaining items
      sliders.value.forEach((item, index) => {
        item.order = index + 1
      })
    }
  }
}
</script>

<style scoped>
.services-slider-management-page {
  min-height: 100vh;
}
</style>

