<template>
  <div class="full-banner-management-page p-6">
    <div class="max-w-7xl mx-auto">
      <!-- Page Header -->
      <div class="mb-8">
        <div class="flex items-center justify-between mb-4">
          <div>
            <h1 class="text-3xl font-bold text-gray-900">مدیریت بنرهای تمام عرض</h1>
            <p class="text-gray-600">مدیریت و تنظیم بنرهای تمام عرض</p>
          </div>
          <NuxtLink 
            to="/admin/content/banners/FullBanner/create"
            class="bg-blue-600 hover:bg-blue-700 text-white px-6 py-3 rounded-lg font-semibold transition-colors duration-200 inline-flex items-center gap-2"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/>
            </svg>
            ایجاد بنر جدید
          </NuxtLink>
        </div>
      </div>

      <!-- Filters -->
      <div class="bg-white p-6 rounded-lg border border-gray-200 mb-6">
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">جستجو</label>
            <input 
              v-model="filters.search" 
              type="text" 
              placeholder="جستجو در عنوان یا کد..."
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
            <select 
              v-model="filters.status" 
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
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
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="">همه</option>
              <option value="home">صفحه اصلی</option>
              <option value="products">محصولات</option>
              <option value="categories">دسته‌بندی‌ها</option>
              <option value="about">درباره ما</option>
            </select>
          </div>
        </div>
        <div class="flex justify-end mt-4">
          <button 
            @click="clearFilters"
            class="px-4 py-2 text-gray-600 hover:text-gray-800 transition-colors"
          >
            پاک کردن فیلترها
          </button>
        </div>
      </div>

      <!-- Banners List -->
      <div class="bg-white rounded-lg border border-gray-200 overflow-hidden">
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">ترتیب</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">پیش‌نمایش</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عنوان</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کد</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">صفحه</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تنظیمات</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="banner in filteredBanners" :key="banner.id" class="hover:bg-gray-50">
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="flex flex-col gap-1">
                    <button 
                      @click="moveUp(banner.id)"
                      :disabled="banner.order === 1"
                      class="w-6 h-6 bg-gray-100 hover:bg-gray-200 rounded flex items-center justify-center disabled:opacity-50 disabled:cursor-not-allowed"
                    >
                      <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 15l7-7 7 7"/>
                      </svg>
                    </button>
                    <span class="text-sm font-medium text-gray-900">{{ banner.order }}</span>
                    <button 
                      @click="moveDown(banner.id)"
                      :disabled="banner.order === banners.length"
                      class="w-6 h-6 bg-gray-100 hover:bg-gray-200 rounded flex items-center justify-center disabled:opacity-50 disabled:cursor-not-allowed"
                    >
                      <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/>
                      </svg>
                    </button>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="w-20 h-16 bg-gray-100 rounded-lg overflow-hidden">
                    <img 
                      v-if="banner.bannerImage" 
                      :src="banner.bannerImage" 
                      :alt="banner.bannerTitle"
                      class="w-full h-full object-cover"
                    />
                    <div v-else class="w-full h-full flex items-center justify-center text-gray-400">
                      <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"/>
                      </svg>
                    </div>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <div>
                    <div class="text-sm font-medium text-gray-900">{{ banner.title }}</div>
                    <div class="text-sm text-gray-500">{{ banner.bannerTitle }}</div>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <code class="text-sm bg-gray-100 px-2 py-1 rounded">{{ banner.code }}</code>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span class="text-sm text-gray-900">{{ getPageName(banner.page) }}</span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span 
                    :class="getStatusBadgeClass(banner.status)"
                    class="inline-flex px-2 py-1 text-xs font-semibold rounded-full"
                  >
                    {{ getStatusName(banner.status) }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  <div class="space-y-1">
                    <div>ارتفاع: {{ banner.bannerHeight }}px</div>
                    <div>موقعیت متن: {{ banner.textPosition === 'left' ? 'چپ' : 'راست' }}</div>
                    <div>رنگ پس‌زمینه: {{ banner.textBackgroundColor }}</div>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                  <div class="flex items-center gap-2">
                    <NuxtLink 
                      :to="`/admin/content/banners/FullBanner/edit/${banner.id}`"
                      class="text-blue-600 hover:text-blue-900 transition-colors"
                    >
                      ویرایش
                    </NuxtLink>
                    <button 
                      @click="toggleStatus(banner.id)"
                      :class="banner.status === 'active' ? 'text-red-600 hover:text-red-900' : 'text-green-600 hover:text-green-900'"
                      class="transition-colors"
                    >
                      {{ banner.status === 'active' ? 'غیرفعال' : 'فعال' }}
                    </button>
                    <button 
                      @click="duplicateBanner(banner.id)"
                      class="text-purple-600 hover:text-purple-900 transition-colors"
                    >
                      تکرار
                    </button>
                    <button 
                      @click="deleteBanner(banner.id)"
                      class="text-red-600 hover:text-red-900 transition-colors"
                    >
                      حذف
                    </button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Empty State -->
      <div v-if="filteredBanners.length === 0" class="text-center py-12">
        <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"/>
        </svg>
        <h3 class="mt-2 text-sm font-medium text-gray-900">هیچ بنری یافت نشد</h3>
        <p class="mt-1 text-sm text-gray-500">با تغییر فیلترها یا ایجاد بنر جدید شروع کنید.</p>
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

definePageMeta({ layout: 'admin', middleware: 'admin' })

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user, hasPermission } = useAuth()

useHead({ title: 'مدیریت بنرهای تمام عرض - پنل ادمین' })

// تعریف نوع داده برای بنر
interface Banner {
  id: string
  title: string
  code: string
  page: string
  order: number
  status: string
  bannerTitle: string
  bannerImage: string
  bannerLink: string
  buttonText: string
  bannerDescription: string
  bannerHeight: number
  textPosition: string
  textBackgroundColor: string
  showTitle: boolean
  showDescription: boolean
  showButton: boolean
  parallaxEffect: boolean
}

// تعریف نوع برای فیلترها
interface Filters {
  search: string
  status: string
  page: string
}

// فیلترها
const filters = ref<Filters>({
  search: '',
  status: '',
  page: ''
})

// داده‌های نمونه
const banners = ref<Banner[]>([
  {
    id: '1',
    title: 'بنر اصلی صفحه اصلی',
    code: 'main-banner',
    page: 'home',
    order: 1,
    status: 'active',
    bannerTitle: 'فروش ویژه تابستانه',
    bannerImage: '/images/banner1.jpg',
    bannerLink: '/sale',
    buttonText: 'مشاهده محصولات',
    bannerDescription: 'تا 50% تخفیف روی محصولات تابستانه',
    bannerHeight: 400,
    textPosition: 'left',
    textBackgroundColor: '#000000',
    showTitle: true,
    showDescription: true,
    showButton: true,
    parallaxEffect: false
  },
  {
    id: '2',
    title: 'بنر محصولات',
    code: 'products-banner',
    page: 'products',
    order: 2,
    status: 'active',
    bannerTitle: 'محصولات جدید',
    bannerImage: '/images/banner2.jpg',
    bannerLink: '/new-products',
    buttonText: 'مشاهده',
    bannerDescription: 'آخرین محصولات اضافه شده به فروشگاه',
    bannerHeight: 350,
    textPosition: 'right',
    textBackgroundColor: '#1f2937',
    showTitle: true,
    showDescription: true,
    showButton: true,
    parallaxEffect: true
  },
  {
    id: '3',
    title: 'بنر دسته‌بندی‌ها',
    code: 'categories-banner',
    page: 'categories',
    order: 3,
    status: 'draft',
    bannerTitle: 'دسته‌بندی‌های محبوب',
    bannerImage: '/images/banner3.jpg',
    bannerLink: '/categories',
    buttonText: 'مشاهده همه',
    bannerDescription: 'محبوب‌ترین دسته‌بندی‌های محصولات',
    bannerHeight: 300,
    textPosition: 'center',
    textBackgroundColor: '#374151',
    showTitle: true,
    showDescription: false,
    showButton: true,
    parallaxEffect: false
  }
])

// فیلتر کردن بنرها
const filteredBanners = computed(() => {
  return banners.value.filter(banner => {
    const matchesSearch = !filters.value.search || 
      banner.title.toLowerCase().includes(filters.value.search.toLowerCase()) ||
      banner.code.toLowerCase().includes(filters.value.search.toLowerCase())
    
    const matchesStatus = !filters.value.status || banner.status === filters.value.status
    const matchesPage = !filters.value.page || banner.page === filters.value.page
    
    return matchesSearch && matchesStatus && matchesPage
  })
})

// دریافت نام صفحه
const getPageName = (page: string) => {
  const pageNames: Record<string, string> = {
    home: 'صفحه اصلی',
    products: 'محصولات',
    categories: 'دسته‌بندی‌ها',
    about: 'درباره ما'
  }
  return pageNames[page] || page
}

// دریافت نام وضعیت
const getStatusName = (status: string) => {
  const statusNames: Record<string, string> = {
    active: 'فعال',
    inactive: 'غیرفعال',
    draft: 'پیش‌نویس'
  }
  return statusNames[status] || status
}

// دریافت کلاس badge وضعیت
const getStatusBadgeClass = (status: string) => {
  const statusClasses: Record<string, string> = {
    active: 'bg-green-100 text-green-800',
    inactive: 'bg-red-100 text-red-800',
    draft: 'bg-yellow-100 text-yellow-800'
  }
  return statusClasses[status] || 'bg-gray-100 text-gray-800'
}

// پاک کردن فیلترها
const clearFilters = () => {
  filters.value.search = ''
  filters.value.status = ''
  filters.value.page = ''
}

// جابجایی به بالا
const moveUp = (id: string) => {
  const banner = banners.value.find(b => b.id === id)
  if (banner && banner.order > 1) {
    const prevBanner = banners.value.find(b => b.order === banner.order - 1)
    if (prevBanner) {
      [banner.order, prevBanner.order] = [prevBanner.order, banner.order]
      // مرتب کردن بر اساس ترتیب
      banners.value.sort((a, b) => a.order - b.order)
    }
  }
}

// جابجایی به پایین
const moveDown = (id: string) => {
  const banner = banners.value.find(b => b.id === id)
  if (banner && banner.order < banners.value.length) {
    const nextBanner = banners.value.find(b => b.order === banner.order + 1)
    if (nextBanner) {
      [banner.order, nextBanner.order] = [nextBanner.order, banner.order]
      // مرتب کردن بر اساس ترتیب
      banners.value.sort((a, b) => a.order - b.order)
    }
  }
}

// تغییر وضعیت
const toggleStatus = (id: string) => {
  const banner = banners.value.find(b => b.id === id)
  if (banner) {
    if (banner.status === 'active') {
      banner.status = 'inactive'
    } else if (banner.status === 'inactive') {
      banner.status = 'active'
    } else {
      banner.status = 'active'
    }
  }
}

// تکرار بنر
const duplicateBanner = (id: string) => {
  const banner = banners.value.find(b => b.id === id)
  if (banner) {
    const newBanner: Banner = {
      ...banner,
      id: Date.now().toString(),
      title: `${banner.title} (کپی)`,
      code: `${banner.code}_copy`,
      order: banners.value.length + 1,
      status: 'draft'
    }
    banners.value.push(newBanner)
    // مرتب کردن بر اساس ترتیب
    banners.value.sort((a, b) => a.order - b.order)
  }
}

// حذف بنر
const deleteBanner = (id: string) => {
  if (confirm('آیا از حذف این بنر اطمینان دارید؟')) {
    const index = banners.value.findIndex(b => b.id === id)
    if (index > -1) {
      banners.value.splice(index, 1)
      // بروزرسانی ترتیب
      banners.value.forEach((banner, idx) => {
        banner.order = idx + 1
      })
    }
  }
}
</script>

