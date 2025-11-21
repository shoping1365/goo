<template>
  <ClientOnly>
    <!-- Page Header -->
    <div class="bg-white shadow-sm border-b border-gray-200 mb-4">
      <div class="w-full px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center py-3">
          <div>
            <h1 class="text-lg font-bold text-gray-900">جزییات گروه مشخصه</h1>
            <p class="text-xs text-gray-600 mt-1">ویرایش و مدیریت گروه ویژگی‌ها و مشخصات فنی</p>
          </div>
          <div class="flex flex-row-reverse gap-6">
            <button 
              class="inline-flex items-center px-3 py-1.5 border border-red-300 text-xs font-medium rounded-md text-white bg-red-600 hover:bg-red-700 shadow-sm"
              @click="deleteGroup"
            >
              حذف
            </button>
            <button 
              class="inline-flex items-center px-3 py-1.5 border border-yellow-300 text-xs font-medium rounded-md text-white bg-yellow-600 hover:bg-yellow-700 shadow-sm"
              @click="saveAndContinue"
            >
              ذخیره و ادامه ویرایش
            </button>
            <button 
              class="inline-flex items-center px-3 py-1.5 border border-green-300 text-xs font-medium rounded-md text-white bg-green-600 hover:bg-green-700 shadow-sm"
              @click="saveAndExit"
            >
              ذخیره
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Back Button -->
    <div class="w-full px-4 py-2">
      <button 
        class="inline-flex items-center px-3 py-1.5 border border-blue-300 text-xs font-medium rounded-md text-blue-700 bg-blue-50 hover:bg-blue-100 shadow-sm"
        @click="$router.push('/admin/attribute-groups')"
      >
        🔙 بازگشت به لیست گروه مشخصه ها
      </button>
    </div>

    <!-- گروه ویژگی (نام و دسته‌بندی) -->
    <div class="w-full px-4 py-4">
      <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6 mb-6 flex flex-row items-end gap-x-4 justify-between">
        <div class="flex flex-row items-end gap-x-4">
          <div class="inline-block">
            <label class="block text-xs font-semibold text-gray-700 mb-1">نام گروه ویژگی</label>
            <input v-model="groupName" type="text" class="input-select-uniform" placeholder="نام گروه ویژگی را وارد کنید" />
          </div>
          <div class="inline-block">
            <label class="block text-xs font-semibold text-gray-700 mb-1">دسته‌بندی گروه ویژگی</label>
            <select v-model="selectedCategory" class="input-select-uniform">
              <option disabled value="">انتخاب کنید...</option>
              <option v-for="cat in categories" :key="cat.id" :value="cat.id">{{ cat.name }}</option>
            </select>
          </div>
        </div>
      </div>

      <!-- Attributes Management Table -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
        <div class="bg-gradient-to-r from-slate-50 to-gray-50 px-4 py-3 border-b border-gray-200">
          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-6 h-6 bg-slate-600 rounded-md flex items-center justify-center ml-2">
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16"></path>
                </svg>
              </div>
              <h3 class="text-sm font-semibold text-gray-900">لیست مشخصات</h3>
            </div>
            
            <div class="flex items-center space-x-2 space-x-reverse">
              <button 
                class="inline-flex items-center px-3 py-1.5 border border-green-300 text-xs font-medium rounded-md text-white bg-green-600 hover:bg-green-700 shadow-sm"
                @click="addNewAttribute"
              >
                افزودن ویژگی جدید
              </button>
            </div>
          </div>
        </div>

        <!-- Attributes Table -->
        <div class="overflow-hidden">
          <div class="overflow-x-auto">
            <table class="min-w-full divide-y divide-gray-200">
              <thead class="bg-gray-50">
                <tr>
                  <th scope="col" class="px-3 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">نام ویژگی</th>
                  <th scope="col" class="px-3 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">ترتیب نمایش</th>
                  <th scope="col" class="px-3 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">نوع</th>
                  <th scope="col" class="px-3 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">نوع کنترل نمایش</th>
                  <th scope="col" class="px-3 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">ابزار فیلتر</th>
                  <th scope="col" class="px-3 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">مشخصه کلیدی</th>
                  <th scope="col" class="px-3 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">نمایش بروی صفحه محصول</th>
                  <th scope="col" class="px-3 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">ضروری</th>
                  <th scope="col" class="px-3 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">عملیات</th>
                </tr>
              </thead>
              <tbody class="bg-white divide-y divide-gray-200">
                <tr v-for="attribute in attributes" :key="attribute.id" class="hover:bg-gray-50 transition-colors">
                  
                  <!-- Editing Mode -->
                  <template v-if="editingAttributeId === attribute.id">
                    <!-- Name (Editable) -->
                    <td class="px-3 py-3 whitespace-nowrap text-right">
                      <input 
                        v-model="editingAttribute!.name"
                        type="text"
                        class="w-full border border-gray-300 rounded px-2 py-1 text-xs focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                      />
                    </td>

                    <!-- Display Order (Editable) -->
                    <td class="px-3 py-3 whitespace-nowrap text-center">
                      <input 
                        v-model.number="editingAttribute!.display_order"
                        type="number"
                        min="1"
                        class="w-16 border border-gray-300 rounded px-2 py-1 text-xs text-center focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                      />
                    </td>

                    <!-- Type (Editable) -->
                    <td class="px-3 py-3 whitespace-nowrap text-center">
                      <select 
                        v-model="editingAttribute!.type"
                        class="border border-gray-300 rounded px-2 py-1 text-xs bg-white focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                      >
                        <option value="انتخاب">انتخاب</option>
                        <option value="متن سفارشی">متن سفارشی</option>
                      </select>
                    </td>

                    <!-- Control Type (Editable) -->
                    <td class="px-3 py-3 whitespace-nowrap text-center">
                      <select 
                        v-model="editingAttribute!.control_type"
                        class="border border-gray-300 rounded px-2 py-1 text-xs bg-white focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                      >
                        <option value="لیست باز شونده تک انتخابی">لیست باز شونده تک انتخابی</option>
                        <option value="تکست باکس متنی تک خطی">تکست باکس متنی تک خطی</option>
                        <option value="تکست باکس متنی چند خطی">تکست باکس متنی چند خطی</option>
                      </select>
                    </td>

                    <!-- Filter Tool (Editable) -->
                    <td class="px-3 py-3 whitespace-nowrap text-center">
                      <input 
                        v-model="editingAttribute!.has_filter"
                        type="checkbox"
                        class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                      />
                    </td>

                    <!-- Key Attribute (Editable) -->
                    <td class="px-3 py-3 whitespace-nowrap text-center">
                      <input 
                        v-model="editingAttribute!.is_key"
                        type="checkbox"
                        class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                      />
                    </td>

                    <!-- Show On Product Page (Editable) -->
                    <td class="px-3 py-3 whitespace-nowrap text-center">
                      <input 
                        v-model="editingAttribute!.show_on_product"
                        type="checkbox"
                        class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                      />
                    </td>

                    <!-- Required (Editable) -->
                    <td class="px-3 py-3 whitespace-nowrap text-center">
                      <input 
                        v-model="editingAttribute!.is_required"
                        type="checkbox"
                        class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                      />
                    </td>

                    <!-- Save/Cancel Actions -->
                    <td class="px-3 py-3 whitespace-nowrap text-center">
                      <div class="flex items-center justify-center space-x-1 space-x-reverse">
                        <button 
                          class="inline-flex items-center p-1.5 border border-transparent text-xs font-medium rounded text-green-700 bg-green-100 hover:bg-green-200 transition-colors"
                          title="ذخیره"
                          @click="saveAttribute"
                        >
                          ✓
                        </button>
                        <button 
                          class="inline-flex items-center p-1.5 border border-transparent text-xs font-medium rounded text-red-700 bg-red-100 hover:bg-red-200 transition-colors"
                          title="انصراف"
                          @click="cancelEdit"
                        >
                          ✗
                        </button>
                      </div>
                    </td>
                  </template>

                  <!-- View Mode -->
                  <template v-else>
                    <!-- Name -->
                    <td class="px-3 py-3 whitespace-nowrap text-xs font-medium text-gray-900 text-right">
                      {{ attribute.name }}
                    </td>

                    <!-- Display Order -->
                    <td class="px-3 py-3 whitespace-nowrap text-center">
                      <span class="inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800">
                        {{ attribute.display_order }}
                      </span>
                    </td>

                    <!-- Type -->
                    <td class="px-3 py-3 whitespace-nowrap text-center">
                      <span class="text-xs text-gray-600">{{ attribute.type }}</span>
                    </td>

                    <!-- Control Type -->
                    <td class="px-3 py-3 whitespace-nowrap text-center">
                      <span class="text-xs text-gray-600">{{ attribute.control_type }}</span>
                    </td>

                    <!-- Filter Tool -->
                    <td class="px-3 py-3 whitespace-nowrap text-center">
                      <span v-if="attribute.has_filter" class="text-green-600 text-lg">✓</span>
                      <span v-else class="text-red-600 text-lg">✗</span>
                    </td>

                    <!-- Key Attribute -->
                    <td class="px-3 py-3 whitespace-nowrap text-center">
                      <span v-if="attribute.is_key" class="text-green-600 text-lg">✓</span>
                      <span v-else class="text-red-600 text-lg">✗</span>
                    </td>

                    <!-- Show On Product Page -->
                    <td class="px-3 py-3 whitespace-nowrap text-center">
                      <span v-if="attribute.show_on_product" class="text-green-600 text-lg">✓</span>
                      <span v-else class="text-red-600 text-lg">✗</span>
                    </td>

                    <!-- Required -->
                    <td class="px-3 py-3 whitespace-nowrap text-center">
                      <span v-if="attribute.is_required" class="text-green-600 text-lg">✓</span>
                      <span v-else class="text-red-600 text-lg">✗</span>
                    </td>

                    <!-- Actions -->
                    <td class="px-3 py-3 whitespace-nowrap text-center">
                      <div class="flex items-center justify-center space-x-1 space-x-reverse">
                        <button 
                          class="inline-flex items-center p-1.5 border border-transparent text-xs font-medium rounded text-amber-700 bg-amber-100 hover:bg-amber-200 transition-colors"
                          title="ویرایش"
                          @click="editAttribute(attribute)"
                        >
                          ✏️
                        </button>
                        <button 
                          class="inline-flex items-center p-1.5 border border-transparent text-xs font-medium rounded text-red-700 bg-red-100 hover:bg-red-200 transition-colors"
                          title="حذف"
                          @click="deleteAttribute(attribute.id)"
                        >
                          ✗
                        </button>
                      </div>
                    </td>
                  </template>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <!-- Counter and Pagination -->
        <div class="bg-gray-50 px-4 py-3 border-t border-gray-200">
          <div class="flex items-center justify-between">
            <div class="flex items-center text-sm text-gray-700">
              <span class="font-medium">{{ paginationInfo.start }}</span>
              <span class="mx-1">تا</span>
              <span class="font-medium">{{ paginationInfo.end }}</span>
              <span class="mx-1">از</span>
              <span class="font-medium">{{ paginationInfo.total }}</span>
              <span class="mr-1">مشخصه</span>
            </div>
            
            <div class="flex items-center space-x-2 space-x-reverse">
              <div class="flex items-center space-x-1 space-x-reverse text-sm text-gray-600">
                <span>نمایش</span>
                <select 
                  v-model="itemsPerPage" 
                  class="border border-gray-300 rounded px-2 py-1 text-xs bg-white focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  @change="updatePagination"
                >
                  <option value="5">5</option>
                  <option value="10">10</option>
                  <option value="25">25</option>
                  <option value="50">50</option>
                </select>
                <span>در هر صفحه</span>
              </div>
              
              <div class="flex items-center space-x-1 space-x-reverse">
                <button 
                  :disabled="currentPage <= 1"
                  :class="[
                    'px-2 py-1 text-xs border rounded transition-colors',
                    currentPage <= 1 
                      ? 'border-gray-300 text-gray-400 bg-gray-100 cursor-not-allowed' 
                      : 'border-gray-300 text-gray-700 bg-white hover:bg-gray-50'
                  ]"
                  @click="goToPage(currentPage - 1)"
                >
                  قبلی
                </button>
                
                <div class="flex items-center space-x-1 space-x-reverse">
                  <template v-for="page in visiblePages" :key="page">
                    <button 
                      v-if="page !== '...'"
                      :class="[
                        'px-2 py-1 text-xs border rounded transition-colors',
                        currentPage === page 
                          ? 'border-blue-500 text-white bg-blue-600' 
                          : 'border-gray-300 text-gray-700 bg-white hover:bg-gray-50'
                      ]"
                      @click="goToPage(Number(page))"
                    >
                      {{ page }}
                    </button>
                    <span v-else class="px-2 py-1 text-xs text-gray-400">...</span>
                  </template>
                </div>
                
                <button 
                  :disabled="currentPage >= totalPages"
                  :class="[
                    'px-2 py-1 text-xs border rounded transition-colors',
                    currentPage >= totalPages 
                      ? 'border-gray-300 text-gray-400 bg-gray-100 cursor-not-allowed' 
                      : 'border-gray-300 text-gray-700 bg-white hover:bg-gray-50'
                  ]"
                  @click="goToPage(currentPage + 1)"
                >
                  بعدی
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <template #fallback>
      <div class="min-h-screen bg-blue-50 flex items-center justify-center">
        <div class="text-center">
          <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600 mx-auto mb-3"></div>
          <div class="text-gray-600 text-sm">در حال بارگذاری...</div>
        </div>
      </div>
    </template>
  </ClientOnly>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useConfirmDialog } from '~/composables/useConfirmDialog';
import { useNotifier } from '~/composables/useNotifier';

definePageMeta({
  layout: 'admin-main'
})

useHead({
  title: 'ویرایش گروه مشخصات فنی'
})

interface Attribute {
  id: number
  name: string
  display_order: number
  type: string
  control_type: string
  has_filter: boolean
  is_key: boolean
  show_on_product: boolean
  is_required: boolean
}

interface Category {
  id: number
  name: string
  parent_id?: number | null
}

interface AttributeGroup {
  id: number
  name: string
  category_id?: number | null
}

const route = useRoute()
const router = useRouter()

const groupId = computed(() => route.query.id as string)

// Pagination variables
const currentPage = ref(1)
const itemsPerPage = ref(5)

// Editing state
const editingAttributeId = ref<number | null>(null)
const editingAttribute = ref<Attribute | null>(null)

// Attribute list (to be loaded from backend)
const attributes = ref<Attribute[]>([])

// Computed properties for pagination
const totalPages = computed(() => Math.ceil(attributes.value.length / itemsPerPage.value))

const paginationInfo = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage.value + 1
  const end = Math.min(start + itemsPerPage.value - 1, attributes.value.length)
  return {
    start,
    end,
    total: attributes.value.length
  }
})

const visiblePages = computed(() => {
  const pages = []
  const total = totalPages.value
  const current = currentPage.value
  
  if (total <= 7) {
    for (let i = 1; i <= total; i++) {
      pages.push(i)
    }
  } else {
    if (current <= 4) {
      for (let i = 1; i <= 5; i++) pages.push(i)
      pages.push('...')
      pages.push(total)
    } else if (current >= total - 3) {
      pages.push(1)
      pages.push('...')
      for (let i = total - 4; i <= total; i++) pages.push(i)
    } else {
      pages.push(1)
      pages.push('...')
      for (let i = current - 1; i <= current + 1; i++) pages.push(i)
      pages.push('...')
      pages.push(total)
    }
  }
  
  return pages
})

// Methods
const deleteGroup = async () => {
  const { confirm } = useConfirmDialog()
  const ok = await confirm({
    title: 'حذف گروه',
    message: 'آیا از حذف این گروه اطمینان دارید؟',
    confirmText: 'حذف',
    cancelText: 'انصراف',
    type: 'danger'
  })
  if (ok) {
    router.push('/admin/attribute-groups')
  }
}

const saveAndContinue = () => {
  useNotifier().success('تغییرات ذخیره شد')
}

const saveAndExit = () => {
  router.push('/admin/attribute-groups')
}

const addNewAttribute = () => {
  // Open modal or navigate to add form
}

const editAttribute = (attribute: Attribute) => {
  editingAttributeId.value = attribute.id
  editingAttribute.value = { ...attribute } // Create a copy to avoid direct mutation
}

const deleteAttribute = async (attributeId: number) => {
  const { confirm } = useConfirmDialog()
  const ok = await confirm({
    title: 'حذف مشخصه',
    message: 'آیا از حذف این مشخصه اطمینان دارید؟',
    confirmText: 'حذف',
    cancelText: 'انصراف',
    type: 'danger'
  })
  if (ok) {
    const index = attributes.value.findIndex(attr => attr.id === attributeId)
    if (index > -1) {
      attributes.value.splice(index, 1)
    }
  }
}

const saveAttribute = () => {
  if (editingAttribute.value && editingAttributeId.value) {
    const index = attributes.value.findIndex(attr => attr.id === editingAttributeId.value)
    if (index > -1) {
      attributes.value[index] = { ...editingAttribute.value }
    }
  }
  editingAttributeId.value = null
  editingAttribute.value = null
}

const cancelEdit = () => {
  editingAttributeId.value = null
  editingAttribute.value = null
}

// Pagination methods
const goToPage = (page: number) => {
  if (page >= 1 && page <= totalPages.value) {
    currentPage.value = page
  }
}

const updatePagination = () => {
  currentPage.value = 1
}

// New refs
const groupName = ref('')
const selectedCategory = ref<number | ''>('')
const categories = ref<Category[]>([])

onMounted(async () => {
  try {
    const catData = await $fetch<Category[]>('/api/product-categories')
    if (Array.isArray(catData)) {
      categories.value = catData.filter((c) => !c.parent_id)
    }
  } catch {}

  if (groupId.value) {
    try {
      const g = await $fetch<AttributeGroup>(`/api/attribute-groups/${groupId.value}`)
      groupName.value = g.name || ''
      selectedCategory.value = g.category_id || ''
    } catch {}
  }
})
</script> 
