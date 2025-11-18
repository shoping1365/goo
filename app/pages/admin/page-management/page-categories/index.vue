<template>
  <div class="space-y-6">
    <!-- هدر صفحه -->
    <div class="bg-white rounded-lg shadow-sm p-6">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-2xl font-bold text-gray-900">دسته بندی صفحات</h1>
          <p class="text-gray-600 mt-1">مدیریت دسته‌بندی‌های صفحات سایت</p>
        </div>
        <TemplateButton
          @click="showAddModal = true"
          bgGradient="bg-gradient-to-r from-blue-400 to-blue-600"
          textColor="text-white"
          hoverClass="hover:from-blue-500 hover:to-blue-700 hover:shadow-lg hover:scale-105"
          focusClass="focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          size="medium"
        >
          افزودن دسته‌بندی جدید
        </TemplateButton>
      </div>
    </div>

    <!-- آمار کلی -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <TemplateCard
        title="کل دسته‌بندی‌ها"
        :value="totalCategories"
        variant="blue"
      >
        <template #icon>
          <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          </svg>
        </template>
      </TemplateCard>

      <TemplateCard
        title="دسته‌بندی‌های فعال"
        :value="activeCategories"
        variant="green"
      >
        <template #icon>
          <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
        </template>
      </TemplateCard>

      <TemplateCard
        title="دسته‌بندی‌های غیرفعال"
        :value="inactiveCategories"
        variant="orange"
      >
        <template #icon>
          <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
        </template>
      </TemplateCard>

      <TemplateCard
        title="کل صفحات"
        :value="totalPages"
        variant="purple"
      >
        <template #icon>
          <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
          </svg>
        </template>
      </TemplateCard>
    </div>

    <!-- جدول دسته‌بندی‌ها -->
    <div class="bg-white rounded-lg shadow-sm overflow-hidden">
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                نام دسته‌بندی
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                توضیحات
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                تعداد صفحات
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                وضعیت
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                تاریخ ایجاد
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                عملیات
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="category in categories" :key="category.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="">
                    <div class="h-10 w-10 rounded-lg bg-blue-100 flex items-center justify-center">
                      <svg class="h-6 w-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      </svg>
                    </div>
                  </div>
                  <div class="mr-4">
                    <div class="text-sm font-medium text-gray-900">{{ category.name }}</div>
                    <div class="text-sm text-gray-500">{{ category.slug }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4">
                <div class="text-sm text-gray-900 max-w-xs truncate">{{ category.description }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-blue-100 text-blue-800">
                  {{ category.pageCount }} صفحه
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="category.isActive ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                  {{ category.isActive ? 'فعال' : 'غیرفعال' }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ formatDate(category.createdAt) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <div class="flex space-x-2 space-x-reverse">
                  <TemplateButton
                    @click="editCategory(category)"
                    textColor="text-blue-600"
                    bgGradient="bg-blue-50"
                    hoverClass="hover:bg-blue-100"
                    size="small"
                    :borderColor="'border-0'"
                  >
                    ویرایش
                  </TemplateButton>
                  <TemplateButton
                    @click="toggleCategoryStatus(category.id)"
                    :textColor="category.isActive ? 'text-red-600' : 'text-green-600'"
                    :bgGradient="category.isActive ? 'bg-red-50' : 'bg-green-50'"
                    :hoverClass="category.isActive ? 'hover:bg-red-100' : 'hover:bg-green-100'"
                    size="small"
                    :borderColor="'border-0'"
                  >
                    {{ category.isActive ? 'غیرفعال' : 'فعال' }}
                  </TemplateButton>
                  <TemplateButton
                    v-if="canDeletePageCategory"
                    @click="deleteCategory(category.id)"
                    textColor="text-red-600"
                    bgGradient="bg-red-50"
                    hoverClass="hover:bg-red-100"
                    size="small"
                    :borderColor="'border-0'"
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

    <!-- مودال افزودن/ویرایش دسته‌بندی -->
    <div v-if="showAddModal || showEditModal" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
      <div class="relative top-20 mx-auto p-5 border w-full max-w-md sm:max-w-lg md:max-w-xl shadow-lg rounded-md bg-white">
        <div class="mt-3">
          <h3 class="text-lg font-medium text-gray-900 mb-4">
            {{ showEditModal ? 'ویرایش دسته‌بندی' : 'افزودن دسته‌بندی جدید' }}
          </h3>
          <form @submit.prevent="saveCategory" class="space-y-4">
            <div>
              <label for="categoryName" class="block text-sm font-medium text-gray-700 mb-1">نام دسته‌بندی *</label>
              <input
                id="categoryName"
                v-model="categoryForm.name"
                type="text"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="نام دسته‌بندی را وارد کنید"
              />
            </div>
            <div>
              <label for="categorySlug" class="block text-sm font-medium text-gray-700 mb-1">نامک (Slug)</label>
              <input
                id="categorySlug"
                v-model="categoryForm.slug"
                type="text"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="نامک دسته‌بندی"
              />
            </div>
            <div>
              <label for="categoryDescription" class="block text-sm font-medium text-gray-700 mb-1">توضیحات</label>
              <textarea
                id="categoryDescription"
                v-model="categoryForm.description"
                rows="3"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="توضیحات دسته‌بندی"
              ></textarea>
            </div>
            <div class="flex items-center">
              <input
                id="categoryActive"
                v-model="categoryForm.isActive"
                type="checkbox"
                class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              />
              <label for="categoryActive" class="mr-2 block text-sm text-gray-900">
                فعال
              </label>
            </div>
            <div class="flex justify-end space-x-3 space-x-reverse">
              <TemplateButton
                type="button"
                @click="closeModal"
                bgGradient="bg-gradient-to-r from-gray-200 to-gray-400"
                textColor="text-gray-800"
                hoverClass="hover:from-gray-300 hover:to-gray-500 hover:shadow-lg hover:scale-105"
                focusClass="focus:ring-2 focus:ring-offset-2 focus:ring-gray-400"
                size="medium"
              >
                انصراف
              </TemplateButton>
              <TemplateButton
                type="submit"
                bgGradient="bg-gradient-to-r from-blue-400 to-blue-600"
                textColor="text-white"
                hoverClass="hover:from-blue-500 hover:to-blue-700 hover:shadow-lg hover:scale-105"
                focusClass="focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
                size="medium"
              >
                {{ showEditModal ? 'ویرایش' : 'افزودن' }}
              </TemplateButton>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
</script>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useAuth } from '~/composables/useAuth'

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// تعریف interface برای دسته‌بندی
interface Category {
  id: number
  name: string
  slug: string
  description: string
  isActive: boolean
  pageCount: number
  createdAt: string
}

// تعریف interface برای فرم دسته‌بندی
interface CategoryForm {
  name: string
  slug: string
  description: string
  isActive: boolean
}

// متغیرهای reactive
const showAddModal = ref(false)
const showEditModal = ref(false)
const editingCategoryId = ref<number | null>(null)

// فرم دسته‌بندی
const categoryForm = ref<CategoryForm>({
  name: '',
  slug: '',
  description: '',
  isActive: true
})

// داده‌های نمونه
const categories = ref<Category[]>([
  {
    id: 1,
    name: 'صفحات عمومی',
    slug: 'general-pages',
    description: 'صفحات عمومی و اطلاعاتی سایت',
    isActive: true,
    pageCount: 5,
    createdAt: '2024-01-15T10:30:00Z'
  },
  {
    id: 2,
    name: 'صفحات اطلاعاتی',
    slug: 'information-pages',
    description: 'صفحات حاوی اطلاعات و راهنما',
    isActive: true,
    pageCount: 3,
    createdAt: '2024-01-10T14:20:00Z'
  },
  {
    id: 3,
    name: 'صفحات تماس',
    slug: 'contact-pages',
    description: 'صفحات مربوط به تماس و ارتباط',
    isActive: false,
    pageCount: 2,
    createdAt: '2024-01-20T09:15:00Z'
  }
])

// محاسبه آمار
const totalCategories = computed(() => categories.value.length)
const activeCategories = computed(() => categories.value.filter(cat => cat.isActive).length)
const inactiveCategories = computed(() => categories.value.filter(cat => !cat.isActive).length)
const totalPages = computed(() => categories.value.reduce((sum, cat) => sum + cat.pageCount, 0))

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user, hasPermission } = useAuth()

// Computed برای چک کردن پرمیژن حذف
const canDeletePageCategory = computed(() => hasPermission('page-category.delete'))

// باز کردن مودال افزودن
const openAddModal = () => {
  showAddModal.value = true
  resetForm()
}

// باز کردن مودال ویرایش
const editCategory = (category: Category) => {
  showEditModal.value = true
  editingCategoryId.value = category.id
  categoryForm.value = {
    name: category.name,
    slug: category.slug,
    description: category.description,
    isActive: category.isActive
  }
}

// بستن مودال
const closeModal = () => {
  showAddModal.value = false
  showEditModal.value = false
  editingCategoryId.value = null
  resetForm()
}

// ریست کردن فرم
const resetForm = () => {
  categoryForm.value = {
    name: '',
    slug: '',
    description: '',
    isActive: true
  }
}

// ذخیره دسته‌بندی
const saveCategory = async () => {
  try {
    if (showEditModal.value && editingCategoryId.value) {
      // ویرایش دسته‌بندی موجود
      const index = categories.value.findIndex(cat => cat.id === editingCategoryId.value)
      if (index !== -1) {
        categories.value[index] = {
          ...categories.value[index],
          ...categoryForm.value
        }
      }
    } else {
      // افزودن دسته‌بندی جدید
      const newCategory: Category = {
        id: Date.now(),
        ...categoryForm.value,
        pageCount: 0,
        createdAt: new Date().toISOString()
      }
      categories.value.push(newCategory)
    }
    
    closeModal()
    console.log('دسته‌بندی ذخیره شد:', categoryForm.value)
  } catch (error) {
    console.error('خطا در ذخیره دسته‌بندی:', error)
  }
}

// تغییر وضعیت دسته‌بندی
const toggleCategoryStatus = (categoryId: number) => {
  const category = categories.value.find(cat => cat.id === categoryId)
  if (category) {
    category.isActive = !category.isActive
    console.log('وضعیت دسته‌بندی تغییر کرد:', category.name, category.isActive)
  }
}

// حذف دسته‌بندی
const deleteCategory = (categoryId: number) => {
  if (confirm('آیا از حذف این دسته‌بندی اطمینان دارید؟')) {
    const index = categories.value.findIndex(cat => cat.id === categoryId)
    if (index !== -1) {
      categories.value.splice(index, 1)
      console.log('دسته‌بندی حذف شد:', categoryId)
    }
  }
}

// فرمت تاریخ
const formatDate = (dateString: string) => {
  const date = new Date(dateString)
  return date.toLocaleDateString('fa-IR')
}
</script> 
