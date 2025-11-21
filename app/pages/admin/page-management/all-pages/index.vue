<template>
  <div class="space-y-6">
    <!-- هدر صفحه -->
    <div class="bg-white rounded-lg shadow-sm p-6">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-2xl font-bold text-gray-900">همه صفحات</h1>
          <p class="text-gray-600 mt-1">مدیریت و ویرایش صفحات سایت</p>
        </div>
        <NuxtLink to="/admin/page-management/create-page">
          <TemplateButton
            bg-gradient="bg-gradient-to-r from-blue-400 to-blue-600"
            text-color="text-white"
            hover-class="hover:from-blue-500 hover:to-blue-700 hover:shadow-lg hover:scale-105"
            focus-class="focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
            size="medium"
          >
            ایجاد صفحه جدید
          </TemplateButton>
        </NuxtLink>
      </div>
    </div>

    <!-- فیلترها و جستجو -->
    <div class="bg-white rounded-lg shadow-sm p-6">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div>
          <label for="search" class="block text-sm font-medium text-gray-700 mb-1">جستجو</label>
          <input
            id="search"
            v-model="filters.search"
            type="text"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            placeholder="جستجو در عنوان یا محتوا..."
          />
        </div>
        <div>
          <label for="statusFilter" class="block text-sm font-medium text-gray-700 mb-1">وضعیت</label>
          <select
            id="statusFilter"
            v-model="filters.status"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه وضعیت‌ها</option>
            <option value="published">منتشر شده</option>
            <option value="draft">پیش‌نویس</option>
            <option value="private">خصوصی</option>
          </select>
        </div>
        <div>
          <label for="categoryFilter" class="block text-sm font-medium text-gray-700 mb-1">دسته‌بندی</label>
          <select
            id="categoryFilter"
            v-model="filters.category"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه دسته‌بندی‌ها</option>
            <option value="1">صفحات عمومی</option>
            <option value="2">صفحات اطلاعاتی</option>
            <option value="3">صفحات تماس</option>
          </select>
        </div>
        <div class="flex items-end">
          <TemplateButton
            bg-gradient="bg-gradient-to-r from-gray-400 to-gray-600"
            text-color="text-white"
            hover-class="hover:from-gray-500 hover:to-gray-700 hover:shadow-lg hover:scale-105"
            focus-class="focus:ring-2 focus:ring-offset-2 focus:ring-gray-500"
            size="medium"
            class="w-full"
            @click="applyFilters"
          >
            اعمال فیلتر
          </TemplateButton>
        </div>
      </div>
    </div>

    <!-- جدول صفحات -->
    <div class="bg-white rounded-lg shadow-sm overflow-hidden">
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                عنوان
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                دسته‌بندی
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                وضعیت
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                تاریخ ایجاد
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                بازدید
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                عملیات
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="page in pages" :key="page.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="">
                    <img v-if="page.featuredImage" :src="page.featuredImage" alt="" class="h-10 w-10 rounded-lg object-cover" />
                    <div v-else class="h-10 w-10 bg-gray-200 rounded-lg flex items-center justify-center">
                      <svg class="h-6 w-6 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                      </svg>
                    </div>
                  </div>
                  <div class="mr-4">
                    <div class="text-sm font-medium text-gray-900">{{ page.title }}</div>
                    <div class="text-sm text-gray-500">{{ page.excerpt }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-blue-100 text-blue-800">
                  {{ page.category }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusClass(page.status)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                  {{ getStatusText(page.status) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ formatDate(page.createdAt) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ page.views }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <div class="flex space-x-2 space-x-reverse">
                  <TemplateButton
                    text-color="text-blue-600"
                    bg-gradient="bg-blue-50"
                    hover-class="hover:bg-blue-100"
                    size="small"
                    :border-color="'border-0'"
                    @click="editPage(page.id)"
                  >
                    ویرایش
                  </TemplateButton>
                  <TemplateButton
                    text-color="text-green-600"
                    bg-gradient="bg-green-50"
                    hover-class="hover:bg-green-100"
                    size="small"
                    :border-color="'border-0'"
                    @click="previewPage(page.id)"
                  >
                    پیش‌نمایش
                  </TemplateButton>
                  <TemplateButton
                    v-if="canDeletePage"
                    text-color="text-red-600"
                    bg-gradient="bg-red-50"
                    hover-class="hover:bg-red-100"
                    size="small"
                    :border-color="'border-0'"
                    @click="deletePage(page.id)"
                  >
                    حذف
                  </TemplateButton>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- صفحه‌بندی -->
      <div class="bg-white px-4 py-3 flex items-center justify-between border-t border-gray-200 sm:px-6">
        <div class="flex-1 flex justify-between sm:hidden">
          <TemplateButton
            :disabled="currentPage === 1"
            bg-gradient="bg-white"
            text-color="text-gray-700"
            border-color="border border-gray-300"
            hover-class="hover:bg-gray-50"
            size="small"
            class="disabled:opacity-50 disabled:cursor-not-allowed"
            @click="previousPage"
          >
            قبلی
          </TemplateButton>
          <TemplateButton
            :disabled="currentPage === totalPages"
            bg-gradient="bg-white"
            text-color="text-gray-700"
            border-color="border border-gray-300"
            hover-class="hover:bg-gray-50"
            size="small"
            class="mr-3 disabled:opacity-50 disabled:cursor-not-allowed"
            @click="nextPage"
          >
            بعدی
          </TemplateButton>
        </div>
        <div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
          <div>
            <p class="text-sm text-gray-700">
              نمایش
              <span class="font-medium">{{ (currentPage - 1) * perPage + 1 }}</span>
              تا
              <span class="font-medium">{{ Math.min(currentPage * perPage, totalItems) }}</span>
              از
              <span class="font-medium">{{ totalItems }}</span>
              نتیجه
            </p>
          </div>
          <div>
            <nav class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px" aria-label="Pagination">
              <TemplateButton
                :disabled="currentPage === 1"
                bg-gradient="bg-white"
                text-color="text-gray-500"
                border-color="border border-gray-300"
                hover-class="hover:bg-gray-50"
                size="small"
                class="rounded-r-md px-2 py-2 disabled:opacity-50 disabled:cursor-not-allowed"
                @click="previousPage"
              >
                <span class="sr-only">قبلی</span>
                <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                  <path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd" />
                </svg>
              </TemplateButton>
              <TemplateButton
                v-for="page in visiblePages"
                :key="page"
                :bg-gradient="page === currentPage ? 'bg-blue-50' : 'bg-white'"
                :text-color="page === currentPage ? 'text-blue-600' : 'text-gray-500'"
                :border-color="page === currentPage ? 'border-blue-500' : 'border-gray-300'"
                :hover-class="page === currentPage ? '' : 'hover:bg-gray-50'"
                size="small"
                class="relative inline-flex items-center px-4 py-2 border text-sm font-medium"
                :class="page === currentPage ? 'z-10' : ''"
                @click="goToPage(page)"
              >
                {{ page }}
              </TemplateButton>
              <TemplateButton
                :disabled="currentPage === totalPages"
                bg-gradient="bg-white"
                text-color="text-gray-500"
                border-color="border border-gray-300"
                hover-class="hover:bg-gray-50"
                size="small"
                class="rounded-l-md px-2 py-2 disabled:opacity-50 disabled:cursor-not-allowed"
                @click="nextPage"
              >
                <span class="sr-only">بعدی</span>
                <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                  <path fill-rule="evenodd" d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z" clip-rule="evenodd" />
                </svg>
              </TemplateButton>
            </nav>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
</script>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useAuth } from '~/composables/useAuth'

definePageMeta({
  layout: 'admin-main'
})

// Loading state
const _isLoading = ref(false)

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user: _user, hasPermission } = useAuth()

// Computed برای چک کردن پرمیژن حذف
const canDeletePage = computed(() => hasPermission('page.delete'))

// تعریف interface برای صفحه
interface Page {
  id: number
  title: string
  excerpt: string
  status: 'published' | 'draft' | 'private'
  category: string
  categoryId: number
  featuredImage?: string
  createdAt: string
  views: number
}

// فیلترها
const filters = ref({
  search: '',
  status: '',
  category: ''
})

// داده‌های نمونه
const pages = ref<Page[]>([
  {
    id: 1,
    title: 'درباره ما',
    excerpt: 'صفحه معرفی شرکت و خدمات ما',
    status: 'published',
    category: 'صفحات عمومی',
    categoryId: 1,
    featuredImage: '/images/about-us.jpg',
    createdAt: '2024-01-15T10:30:00Z',
    views: 1250
  },
  {
    id: 2,
    title: 'تماس با ما',
    excerpt: 'اطلاعات تماس و فرم ارتباطی',
    status: 'published',
    category: 'صفحات تماس',
    categoryId: 3,
    featuredImage: '/images/contact.jpg',
    createdAt: '2024-01-10T14:20:00Z',
    views: 890
  },
  {
    id: 3,
    title: 'قوانین و مقررات',
    excerpt: 'قوانین استفاده از خدمات سایت',
    status: 'draft',
    category: 'صفحات اطلاعاتی',
    categoryId: 2,
    createdAt: '2024-01-20T09:15:00Z',
    views: 0
  }
])

// صفحه‌بندی
const currentPage = ref(1)
const perPage = ref(10)
const totalItems = ref(25)
const totalPages = computed(() => Math.ceil(totalItems.value / perPage.value))

// محاسبه صفحات قابل مشاهده
const visiblePages = computed(() => {
  const pages = []
  const start = Math.max(1, currentPage.value - 2)
  const end = Math.min(totalPages.value, currentPage.value + 2)
  
  for (let i = start; i <= end; i++) {
    pages.push(i)
  }
  
  return pages
})

// اعمال فیلترها
const applyFilters = () => {
  // اینجا کد فیلتر کردن صفحات قرار می‌گیرد
}

// ویرایش صفحه
const editPage = (_pageId: number) => {
  // اینجا کد انتقال به صفحه ویرایش قرار می‌گیرد
}

// پیش‌نمایش صفحه
const previewPage = (_pageId: number) => {
  // اینجا کد باز کردن پیش‌نمایش قرار می‌گیرد
}

// حذف صفحه
const deletePage = (_pageId: number) => {
  if (confirm('آیا از حذف این صفحه اطمینان دارید؟')) {
    // اینجا کد حذف صفحه قرار می‌گیرد
  }
}

// صفحه‌بندی
const previousPage = () => {
  if (currentPage.value > 1) {
    currentPage.value--
  }
}

const nextPage = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value++
  }
}

const goToPage = (page: number) => {
  currentPage.value = page
}

// دریافت کلاس وضعیت
const getStatusClass = (status: string) => {
  switch (status) {
    case 'published':
      return 'bg-green-100 text-green-800'
    case 'draft':
      return 'bg-yellow-100 text-yellow-800'
    case 'private':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

// دریافت متن وضعیت
const getStatusText = (status: string) => {
  switch (status) {
    case 'published':
      return 'منتشر شده'
    case 'draft':
      return 'پیش‌نویس'
    case 'private':
      return 'خصوصی'
    default:
      return 'نامشخص'
  }
}

// فرمت تاریخ
const formatDate = (dateString: string) => {
  const date = new Date(dateString)
  return date.toLocaleDateString('fa-IR')
}

// SSR hydration fix
onMounted(() => {
  // این کد فقط در سمت کلاینت اجرا می‌شود
})
</script> 
