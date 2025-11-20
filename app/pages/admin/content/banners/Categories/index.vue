<template>
  <div class="categories-management-page p-6">
    <div class="max-w-7xl mx-auto">
      <!-- Page Header -->
      <div class="mb-8">
        <div class="flex items-center justify-between">
          <div>
            <h1 class="text-3xl font-bold text-gray-900">مدیریت باکس دسته‌بندی</h1>
            <p class="text-gray-600 mt-1">مدیریت و تنظیم ابزارک‌های نمایش دسته‌بندی‌ها</p>
          </div>
          <div class="flex gap-3">
            <button
              :disabled="loading"
              class="bg-blue-600 hover:bg-blue-700 text-white px-6 py-3 rounded-lg font-semibold transition-colors duration-200 flex items-center gap-2 disabled:opacity-50"
              @click="fetchCategories"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"/>
              </svg>
              {{ loading ? 'در حال بارگذاری...' : 'بروزرسانی' }}
            </button>
            <NuxtLink
              to="/admin/content/banners/Categories/create"
              class="bg-green-600 hover:bg-green-700 text-white px-6 py-3 rounded-lg font-semibold transition-colors duration-200 flex items-center gap-2"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/>
              </svg>
              ایجاد باکس جدید
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
              placeholder="جستجو در باکس‌ها..."
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
              <option value="sidebar">نوار کناری</option>
              <option value="footer">فوتر</option>
            </select>
          </div>
        </div>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="flex justify-center items-center py-12">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
        <span class="mr-4 text-gray-600">در حال بارگذاری...</span>
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

      <!-- Categories List -->
      <div v-else-if="filteredCategories.length > 0" class="bg-white rounded-xl shadow-lg overflow-hidden">
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

        <!-- Categories List -->
        <div class="divide-y divide-gray-200">
          <div
            v-for="(category, idx) in filteredCategories"
            :key="category.id"
            class="px-6 py-4 hover:bg-gray-50 transition-colors duration-200"
          >
            <div class="grid grid-cols-12 gap-6 items-center">
              <!-- Order -->
              <div class="col-span-1 flex justify-center">
                <input
                  type="number"
                  :value="category.order"
                  min="1"
                  class="w-16 h-8 bg-white border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 p-2 text-center font-semibold"
                  @change="updateOrder(category.id, $event)"
                />
              </div>

              <!-- Image -->
              <div class="col-span-1 flex justify-center">
                <div class="w-12 h-12 bg-gray-100 rounded-lg border border-gray-300 overflow-hidden">
                  <img
                    v-if="category.image"
                    :src="category.image"
                    :alt="category.title"
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
                <div class="font-semibold text-gray-900">{{ category.title }}</div>
                <div class="text-sm text-gray-500 mt-1">{{ category.description }}</div>
              </div>

              <!-- Code -->
              <div class="col-span-1 flex justify-center">
                <span class="px-3 py-1 text-xs font-semibold rounded-full bg-gray-100 text-gray-800 font-mono">
                  {{ category.code }}
                </span>
              </div>

              <!-- Page -->
              <div class="col-span-1 flex justify-center">
                <span class="px-3 py-1 text-xs font-semibold rounded-full bg-blue-100 text-blue-800">
                  {{ category.page }}
                </span>
              </div>

              <!-- Status -->
              <div class="col-span-1 flex justify-center">
                <button
                  class="px-3 py-1 text-xs font-semibold rounded-full transition-colors"
                  :class="getStatusBadgeClass(category.status)"
                  @click="toggleStatus(category.id)"
                >
                  {{ getStatusLabel(category.status) }}
                </button>
              </div>

              <!-- Configuration -->
              <div class="col-span-2 text-center text-sm text-gray-600">
                <div>{{ category.columns }} ستون</div>
                <div>{{ category.maxCategories }} دسته‌بندی</div>
              </div>

              <!-- Actions -->
              <div class="col-span-3 flex justify-center gap-2">
                <NuxtLink
                  :to="`/admin/content/banners/Categories/edit/${category.id}`"
                  class="w-8 h-8 bg-blue-100 hover:bg-blue-200 text-blue-600 rounded-lg transition-colors duration-200 flex items-center justify-center"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"/>
                  </svg>
                </NuxtLink>
                <button
                  class="w-8 h-8 bg-green-100 hover:bg-green-200 text-green-600 rounded-lg transition-colors duration-200 flex items-center justify-center"
                  @click="duplicateCategory(category.id)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"/>
                  </svg>
                </button>
                <button
                  class="w-8 h-8 bg-red-100 hover:bg-red-200 text-red-600 rounded-lg transition-colors duration-200 flex items-center justify-center"
                  @click="confirmDelete(category.id)"
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
        <h3 class="text-lg font-semibold text-gray-800 mb-2">هنوز باکس دسته‌بندی ایجاد نشده</h3>
        <p class="text-gray-600 mb-4">با ایجاد اولین باکس دسته‌بندی شروع کنید</p>
        <NuxtLink 
          to="/admin/content/banners/Categories/create"
          class="bg-blue-600 hover:bg-blue-700 text-white px-6 py-3 rounded-lg font-semibold transition-colors duration-200 inline-flex items-center gap-2"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/>
          </svg>
          ایجاد اولین باکس
        </NuxtLink>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'

// تعریف definePageMeta و useHead برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
declare const useHead: (head: { title?: string }) => void
// Types
interface Category {
  id: number
  title: string
  description: string
  code: string
  page: string
  status: 'active' | 'inactive' | 'draft'
  order: number
  image: string | null
  columns: number
  maxCategories: number
}

// Page meta definition
definePageMeta({
  layout: 'admin',
  middleware: 'admin'
})

// Page title
useHead({
  title: 'مدیریت باکس دسته‌بندی - پنل ادمین'
})

// State
const loading = ref(false)
const error = ref<string | null>(null)
const searchQuery = ref('')
const statusFilter = ref('')
const pageFilter = ref('')

// Categories data
const categories = ref<Category[]>([
  {
    id: 1,
    title: 'دسته‌بندی‌های اصلی',
    description: 'نمایش دسته‌بندی‌های اصلی سایت',
    code: 'main-categories',
    page: 'home',
    status: 'active',
    order: 1,
    image: null,
    columns: 3,
    maxCategories: 12
  },
  {
    id: 2,
    title: 'دسته‌بندی‌های نوار کناری',
    description: 'دسته‌بندی‌های نمایش داده شده در نوار کناری',
    code: 'sidebar-categories',
    page: 'sidebar',
    status: 'active',
    order: 2,
    image: null,
    columns: 1,
    maxCategories: 8
  }
])

// تابع دریافت لیست باکس‌های دسته‌بندی
const fetchCategories = async () => {
  try {
    loading.value = true
    error.value = null

    // TODO: جایگزین با API واقعی
    // const { data } = await $fetch('/api/admin/content/banners/categories', {
    //   method: 'GET',
    //   headers: {
    //     'Content-Type': 'application/json'
    //   }
    // })

    // categories.value = data || []

  } catch (err) {
    console.error('خطا در دریافت باکس‌های دسته‌بندی:', err)
    error.value = err.message || 'خطا در دریافت باکس‌های دسته‌بندی'
  } finally {
    loading.value = false
  }
}

// مقداردهی اولیه و دریافت داده‌ها
onMounted(() => {
  fetchCategories()
})

// Computed
const filteredCategories = computed(() => {
  let filtered = categories.value

  if (searchQuery.value) {
    const search = searchQuery.value.toLowerCase()
    filtered = filtered.filter(c => 
      c.title.toLowerCase().includes(search) ||
      c.description.toLowerCase().includes(search) ||
      c.code.toLowerCase().includes(search)
    )
  }

  if (statusFilter.value) {
    filtered = filtered.filter(c => c.status === statusFilter.value)
  }

  if (pageFilter.value) {
    filtered = filtered.filter(c => c.page === pageFilter.value)
  }

  return filtered.sort((a, b) => a.order - b.order)
})

// Methods
const updateOrder = async (id: number, event: Event) => {
  const target = event.target as HTMLInputElement
  const newOrder = parseInt(target.value)

  if (isNaN(newOrder) || newOrder < 1) return

  const category = categories.value.find(c => c.id === id)
  if (category) {
    category.order = newOrder
  }

  // TODO: جایگزین با API واقعی
  // try {
  //   await $fetch(`/api/admin/content/banners/categories/${id}/order`, {
  //     method: 'PUT',
  //     headers: {
  //       'Content-Type': 'application/json'
  //     },
  //     body: { order: newOrder }
  //   })
  // } catch (err) {
  //   console.error('خطا در بروزرسانی ترتیب:', err)
  // }

  console.log('Updating order for category:', id, 'to:', newOrder)
}

const toggleStatus = async (id: number) => {
  const category = categories.value.find(c => c.id === id)
  if (category) {
    category.status = category.status === 'active' ? 'inactive' : 'active'
  }

  // TODO: جایگزین با API واقعی
  // try {
  //   await $fetch(`/api/admin/content/banners/categories/${id}/status`, {
  //     method: 'PUT',
  //     headers: {
  //       'Content-Type': 'application/json'
  //     },
  //     body: { status: category.status }
  //   })
  // } catch (err) {
  //   console.error('خطا در تغییر وضعیت:', err)
  // }

  console.log('Toggling status for category:', id)
}

const duplicateCategory = async (id: number) => {
  const category = categories.value.find(c => c.id === id)
  if (category) {
    const newCategory: Category = {
      ...category,
      id: Date.now(),
      title: `${category.title} (کپی)`,
      code: `${category.code}-copy`,
      status: 'draft',
      order: categories.value.length + 1
    }
    categories.value.push(newCategory)
  }

  // TODO: جایگزین با API واقعی
  // try {
  //   const { data } = await $fetch('/api/admin/content/banners/categories/duplicate', {
  //     method: 'POST',
  //     headers: {
  //       'Content-Type': 'application/json'
  //     },
  //     body: { id }
  //   })
  //   categories.value.push(data)
  // } catch (err) {
  //   console.error('خطا در کپی کردن:', err)
  // }

  console.log('Duplicating category:', id)
}

const confirmDelete = async (id: number) => {
  if (confirm('آیا مطمئن هستید که می‌خواهید این باکس دسته‌بندی را حذف کنید؟')) {
    categories.value = categories.value.filter(c => c.id !== id)

    // TODO: جایگزین با API واقعی
    // try {
    //   await $fetch(`/api/admin/content/banners/categories/${id}`, {
    //     method: 'DELETE',
    //     headers: {
    //       'Content-Type': 'application/json'
    //     }
    //   })
    // } catch (err) {
    //   console.error('خطا در حذف:', err)
    // }

    console.log('Deleting category:', id)
  }
}

const clearError = () => {
  error.value = null
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
.categories-management-page {
  background-color: #f9fafb;
  min-height: 100vh;
}
</style>

