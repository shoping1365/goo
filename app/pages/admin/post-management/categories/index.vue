<template>
  <div class="min-h-screen">
    <!-- Header -->
    <div class="bg-white shadow-sm border-b border-gray-200 sticky top-0 z-30">
      <div class="px-6 py-4">
        <div class="flex items-center justify-between">
          <div>
            <h1 class="text-2xl font-bold text-gray-900">دسته‌بندی‌های نوشته‌ها</h1>
            <p class="text-sm text-gray-500 mt-1">مدیریت دسته‌بندی‌های محتوا</p>
          </div>
          
          <div class="flex items-center gap-3">
            <!-- جستجو -->
            <div class="relative">
              <input
                v-model="searchQuery"
                type="text"
                placeholder="جستجو در دسته‌بندی‌ها..."
                class="w-64 pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
              <svg class="w-5 h-5 text-gray-400 absolute left-3 top-2.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
              </svg>
            </div>
            
            <!-- افزودن دسته‌بندی جدید -->
            <button 
              class="bg-blue-500 text-white px-4 py-2 rounded-lg hover:bg-blue-600 transition-colors flex items-center gap-2 font-medium"
              @click="openAddModal"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
              </svg>
              افزودن دسته‌بندی جدید
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Main Content -->
    <div class="p-6">
      <!-- آمار -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-6">
        <TemplateCard
          title="کل دسته‌بندی‌ها"
          :value="totalCategories"
          variant="blue"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"></path>
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
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
            </svg>
          </template>
        </TemplateCard>
        <TemplateCard
          title="کل نوشته‌ها"
          :value="totalPosts"
          variant="purple"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
            </svg>
          </template>
        </TemplateCard>
      </div>

      <!-- جدول دسته‌بندی‌ها -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نام دسته‌بندی</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">Slug</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">ترتیب</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تعداد نوشته‌ها</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ ایجاد</th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="category in filteredCategories" :key="category.id" class="hover:bg-gray-50">
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="flex items-center">
                    <div class="flex-shrink-0 h-8 w-8">
                      <div class="h-8 w-8 rounded-lg bg-blue-100 flex items-center justify-center">
                        <svg class="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"></path>
                        </svg>
                      </div>
                    </div>
                    <div class="mr-3">
                      <div class="text-sm font-medium text-gray-900">{{ category.name }}</div>
                      <div class="text-sm text-gray-500">{{ category.description }}</div>
                    </div>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ category.slug }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ category.order }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ category.posts_count }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span :class="category.is_active ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'" class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full">
                    {{ category.is_active ? 'فعال' : 'غیرفعال' }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {{ formatDate(category.created_at) }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                  <div class="flex items-center gap-2">
                    <button 
                      class="text-blue-600 hover:text-blue-900 transition-colors"
                      title="ویرایش"
                      @click="editCategory(category)"
                    >
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
                      </svg>
                    </button>
                    <button 
                      :class="category.is_active ? 'text-red-600 hover:text-red-900' : 'text-green-600 hover:text-green-900'"
                      class="transition-colors"
                      :title="category.is_active ? 'غیرفعال کردن' : 'فعال کردن'"
                      @click="toggleCategoryStatus(category)"
                    >
                      <svg v-if="category.is_active" class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728L5.636 5.636m12.728 12.728L18.364 5.636M5.636 18.364l12.728-12.728"></path>
                      </svg>
                      <svg v-else class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                      </svg>
                    </button>
                    <button 
                      v-if="canDeleteCategory"
                      class="text-red-600 hover:text-red-900 transition-colors"
                      title="حذف"
                      @click="deleteCategory(category)"
                    >
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                      </svg>
                    </button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>

    <!-- Modal افزودن/ویرایش دسته‌بندی -->
    <div v-if="showModal" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
      <div class="relative top-20 mx-auto p-5 border w-full max-w-md sm:max-w-lg md:max-w-xl shadow-lg rounded-md bg-white">
        <div class="mt-3">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-medium text-gray-900">
              {{ isEditing ? 'ویرایش دسته‌بندی' : 'افزودن دسته‌بندی جدید' }}
            </h3>
            <button 
              class="text-gray-400 hover:text-gray-600 transition-colors"
              @click="closeModal"
            >
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
              </svg>
            </button>
          </div>
          
          <form class="space-y-4" @submit.prevent="saveCategory">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">نام دسته‌بندی</label>
              <input 
                v-model="categoryForm.name"
                type="text"
                required
                placeholder="نام دسته‌بندی را وارد کنید..."
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors text-right"
                dir="rtl"
              >
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">توضیحات</label>
              <textarea 
                v-model="categoryForm.description"
                rows="3"
                placeholder="توضیحات دسته‌بندی..."
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors text-right resize-none"
                dir="rtl"
              ></textarea>
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">Slug</label>
              <input 
                v-model="categoryForm.slug"
                type="text"
                required
                placeholder="slug-dastebandi"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors"
              >
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">ترتیب نمایش</label>
              <input 
                v-model="categoryForm.order"
                type="number"
                min="1"
                placeholder="1"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors"
              >
            </div>
            
            <div class="flex items-center">
              <input 
                v-model="categoryForm.is_active"
                type="checkbox"
                class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
              >
              <label class="mr-2 text-sm text-gray-700">فعال</label>
            </div>
            
            <div class="flex items-center justify-end gap-3 pt-4">
              <button 
                type="button"
                class="px-4 py-2 text-gray-700 bg-gray-200 rounded-lg hover:bg-gray-300 transition-colors"
                @click="closeModal"
              >
                انصراف
              </button>
              <button 
                type="submit"
                :disabled="isSaving"
                class="px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 transition-colors disabled:opacity-50"
              >
                {{ isSaving ? 'در حال ذخیره...' : (isEditing ? 'ویرایش' : 'افزودن') }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
declare const useHead: (head: { title?: string }) => void
declare const useFetch: <T>(url: string, options?: unknown) => Promise<{ data: { value: T } }>
</script>

<script setup lang="ts">
import { ref, reactive, computed, onMounted } from 'vue'
import { useAuth } from '~/composables/useAuth'
// 
// تعریف interface برای دسته‌بندی
interface Category {
  id: number
  name: string
  slug: string
  description: string
  is_active: boolean
  posts_count: number
  order: number
  created_at: string
  updated_at: string
}

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

useHead({
  title: 'دسته‌بندی‌های نوشته‌ها - پنل مدیریت'
})

// Reactive data
const categories = ref<Category[]>([])
const searchQuery = ref('')
const showModal = ref(false)
const isEditing = ref(false)
const isSaving = ref(false)

// فرم دسته‌بندی
const categoryForm = reactive({
  id: null as number | null,
  name: '',
  slug: '',
  description: '',
  is_active: true,
  order: 1
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user, hasPermission } = useAuth()

// Computed برای چک کردن پرمیژن حذف
const canDeleteCategory = computed(() => hasPermission('post-category.delete'))

// Methods
const fetchCategories = async () => {
  try {
    const { data } = await useFetch('/api/post-categories?all=1')
    if (Array.isArray(data.value)) {
      categories.value = data.value
    } else {
      categories.value = []
    }
  } catch (error) {
    console.error('خطا در دریافت دسته‌بندی‌ها:', error)
    categories.value = []
  }
}

// Computed properties
const filteredCategories = computed(() => {
  if (!searchQuery.value) return categories.value
  
  return categories.value.filter(category => 
    category.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
    category.description.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
    category.slug.toLowerCase().includes(searchQuery.value.toLowerCase())
  )
})

const totalCategories = computed(() => categories.value.length)
const activeCategories = computed(() => categories.value.filter(c => c.is_active).length)
const totalPosts = computed(() => categories.value.reduce((sum, c) => sum + c.posts_count, 0))

// Methods
const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleDateString('fa-IR')
}

const openAddModal = () => {
  isEditing.value = false
  resetForm()
  showModal.value = true
}

const editCategory = (category: Category) => {
  isEditing.value = true
  Object.assign(categoryForm, {
    id: category.id,
    name: category.name,
    slug: category.slug,
    description: category.description,
    is_active: category.is_active,
    order: category.order
  })
  showModal.value = true
}

const closeModal = () => {
  showModal.value = false
  resetForm()
}

const resetForm = () => {
  Object.assign(categoryForm, {
    id: null,
    name: '',
    slug: '',
    description: '',
    is_active: true,
    order: 1
  })
}

const saveCategory = async () => {
  if (!categoryForm.name.trim() || !categoryForm.slug.trim()) {
    alert('لطفاً نام و slug دسته‌بندی را وارد کنید')
    return
  }

  isSaving.value = true
  
  try {
    const payload = {
      name: categoryForm.name,
      slug: categoryForm.slug,
      description: categoryForm.description,
      is_active: categoryForm.is_active,
      order: categoryForm.order
    }

    if (isEditing.value) {
      // ویرایش دسته‌بندی موجود
      await $fetch(`/api/post-categories/${categoryForm.id}`, {
        method: 'PUT',
        body: payload
      })
    } else {
      // افزودن دسته‌بندی جدید
      await $fetch('/api/post-categories', {
        method: 'POST',
        body: payload
      })
    }
    
    // دریافت مجدد لیست دسته‌بندی‌ها
    await fetchCategories()
    
    closeModal()
    alert(isEditing.value ? 'دسته‌بندی با موفقیت ویرایش شد' : 'دسته‌بندی با موفقیت افزوده شد')
    
  } catch (error) {
    console.error('خطا در ذخیره دسته‌بندی:', error)
    alert('خطا در ذخیره دسته‌بندی')
  } finally {
    isSaving.value = false
  }
}

const toggleCategoryStatus = async (category: Category) => {
  try {
    await $fetch(`/api/post-categories/${category.id}/toggle-status`, {
      method: 'PATCH'
    })
    
    // دریافت مجدد لیست دسته‌بندی‌ها
    await fetchCategories()
    
    alert(category.is_active ? 'دسته‌بندی غیرفعال شد' : 'دسته‌بندی فعال شد')
  } catch (error) {
    console.error('خطا در تغییر وضعیت دسته‌بندی:', error)
    alert('خطا در تغییر وضعیت دسته‌بندی')
  }
}

const deleteCategory = async (category: Category) => {
  if (category.posts_count > 0) {
    alert('این دسته‌بندی دارای نوشته است و قابل حذف نیست')
    return
  }
  
  if (confirm(`آیا از حذف دسته‌بندی "${category.name}" اطمینان دارید؟`)) {
    try {
      await $fetch(`/api/post-categories/${category.id}`, {
        method: 'DELETE'
      })
      
      // دریافت مجدد لیست دسته‌بندی‌ها
      await fetchCategories()
      
      alert('دسته‌بندی با موفقیت حذف شد')
    } catch (error) {
      console.error('خطا در حذف دسته‌بندی:', error)
      alert('خطا در حذف دسته‌بندی')
    }
  }
}

// Lifecycle
onMounted(() => {
  fetchCategories()
})
</script> 

