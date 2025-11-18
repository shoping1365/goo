<template>
  <ClientOnly>
    <!-- Page Header -->
    <div class="bg-white shadow-sm border-b border-gray-200 mb-4">
      <div class="w-full px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center py-3">
          <div>
            <h1 class="text-lg font-bold text-gray-900">ایجاد گروه مشخصات</h1>
            <p class="text-xs text-gray-600 mt-1">ایجاد و مدیریت گروه ویژگی‌ها و مشخصات فنی</p>
          </div>
          <div class="flex items-center gap-x-4">
            <button
              @click="saveAndContinue"
              class="inline-flex items-center px-4 py-2 bg-gradient-to-r from-blue-500 to-blue-700 text-sm font-medium rounded-lg text-white shadow-lg"
            >
              ذخیره و ادامه ویرایش
            </button>
            <button
              @click="saveAndExit"
              class="inline-flex items-center px-4 py-2 bg-gradient-to-r from-green-500 to-green-700 text-sm font-medium rounded-lg text-white shadow-lg"
            >
              ذخیره
            </button>
            <button
              v-if="isSaved"
              @click="deleteGroup"
              class="inline-flex items-center px-4 py-2 bg-gradient-to-r from-red-500 to-red-700 text-sm font-medium rounded-lg text-white shadow-lg"
            >
              حذف
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Back Button -->
    <div class="hidden"></div>

    <!-- Main Content -->
    <div class="w-full px-4 py-4">

      <!-- گروه ویژگی -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6 mb-6">
        <div class="flex flex-row items-end gap-x-4 justify-between">
          <div class="flex flex-row items-end gap-x-4">
            <div class="inline-block">
              <label class="block text-xs font-semibold text-gray-700 mb-1">نام گروه ویژگی</label>
              <input
                v-model="groupName"
                type="text"
                class="input-select-uniform"
                placeholder="نام گروه ویژگی را وارد کنید"
              />
            </div>
            <div class="inline-block">
              <label class="block text-xs font-semibold text-gray-700 mb-1">دسته‌بندی گروه ویژگی</label>
              <select
                v-model="selectedCategory"
                class="input-select-uniform"
              >
                <option disabled value="">انتخاب کنید...</option>
                <option v-for="cat in categories" :key="cat.id" :value="cat.id">{{ cat.name }}</option>
              </select>
            </div>
          </div>
          <button
            @click="$router.push('/admin/product-management/attribute-groups')"
            class="inline-flex items-center px-3 py-1.5 border border-blue-300 text-xs font-medium rounded-md text-blue-700 bg-blue-50 hover:bg-blue-100 shadow-sm"
          >
            🔙 بازگشت به لیست گروه ویژگی ها
          </button>
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
                @click="addNewAttribute"
                class="inline-flex items-center px-4 py-2 bg-gradient-to-r from-purple-500 to-purple-700 text-sm font-medium rounded-lg text-white shadow-lg"
              >
                افزودن ویژگی به لیست
              </button>
            </div>
          </div>
        </div>

        <!-- Modal Selector (displays above everything) -->
        <div v-if="showAddModal" class="fixed z-50" :style="modalStyle">
          <div class="bg-white rounded-lg shadow-md w-full max-w-md sm:max-w-lg md:max-w-xl max-h-[80vh] flex flex-col border border-gray-200">
            <!-- header -->
            <div class="px-4 py-2 bg-purple-50 border-b border-purple-100 rounded-t-lg cursor-move select-none" @mousedown.prevent="startDrag">
              <h3 class="text-sm font-semibold text-gray-800">افزودن ویژگی به لیست</h3>
            </div>
            <div class="p-6 flex-1 overflow-y-auto">
              <!-- list of attributes -->
              <ul class="mb-3 space-y-1 max-h-64 overflow-y-auto pr-1 scrollbar-thin scrollbar-thumb-gray-300 scrollbar-track-gray-100">
                <li v-for="attr in filteredModalAttributes" :key="attr.id"
                    @dblclick="addAttributeById(attr.id)"
                    class="cursor-pointer text-xs px-2 py-1 rounded hover:bg-blue-50 border border-transparent hover:border-blue-200"
                    :class="{'bg-gray-200': isAlreadySelected(attr)}">
                  {{ attr.name }}
                </li>
              </ul>
              <!-- search box -->
              <input v-model="modalSearch" type="text" placeholder="جستجو..." class="w-full mb-4 border border-gray-300 rounded px-2 py-1 text-xs focus:ring-2 focus:ring-blue-500 focus:border-blue-500" />
            </div>
            <!-- footer -->
            <div class="px-4 py-2 bg-red-50 border-t border-red-100 rounded-b-lg flex justify-between">
              <button @click="promptNewAttribute" class="inline-flex items-center px-4 py-2 bg-gradient-to-r from-green-500 to-green-700 text-sm font-medium rounded-lg text-white shadow-lg">
                ایجاد ویژگی
              </button>
              <button @click="closeModal"
                :class="finishButtonClass"
                class="inline-flex items-center px-4 py-2 text-sm font-medium rounded-lg text-white shadow-lg">
                {{ finishButtonLabel }}
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
                <tr v-for="(attribute, index) in attributes" :key="attribute.id" class="hover:bg-gray-50 transition-colors">

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
                        class="w-12 border border-gray-300 rounded px-1.5 py-1 text-xs text-center focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                      />
                    </td>

                    <!-- Type (Editable) -->
                    <td class="px-3 py-3 whitespace-nowrap text-center">
                      <select
                        v-model="editingAttribute!.type"
                        class="w-28 border border-gray-300 rounded px-1.5 py-1 text-xs bg-white focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
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
                        <option v-for="opt in controlOptions(editingAttribute!.type)" :key="opt" :value="opt">{{ opt }}</option>
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
                      <div class="flex items-center justify-center space-x-4 space-x-reverse">
                        <button
                          @click="saveAttribute"
                          class="inline-flex items-center p-1.5 border border-transparent text-xs font-medium rounded text-green-700 bg-green-100 hover:bg-green-200 transition-colors"
                          title="ذخیره"
                        >
                          ✓
                        </button>
                        <button
                          @click="cancelEdit"
                          class="inline-flex items-center p-1.5 border border-transparent text-xs font-medium rounded text-red-700 bg-red-100 hover:bg-red-200 transition-colors"
                          title="انصراف"
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
                      <div class="flex items-center justify-center space-x-4 space-x-reverse">
                        <button
                          @click="editAttribute(attribute)"
                          class="inline-flex items-center p-1.5 border border-transparent text-xs font-medium rounded text-amber-700 bg-amber-100 hover:bg-amber-200 transition-colors"
                          title="ویرایش"
                        >
                          ✏️
                        </button>
                        <button
                          @click="deleteAttribute(attribute.id)"
                          class="inline-flex items-center p-1.5 border border-transparent text-xs font-medium rounded text-red-700 bg-red-100 hover:bg-red-200 transition-colors"
                          title="حذف"
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

        <!-- Pagination Component -->
        <Pagination
          :current-page="currentPage"
          :total-pages="totalPages"
          :total="attributes.length"
          :per-page="itemsPerPage"
          @page-changed="goToPage"
          @per-page-changed="handlePerPageChange"
          class="bg-gray-50 border-t border-gray-200"
        />
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
  <!-- Global Confirm Dialog for this page -->
  <ConfirmDialog ref="confirmDialogRef" />
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
declare const useHead: (head: { title?: string }) => void
</script>

<script setup lang="ts">
import Pagination from '@/components/admin/common/Pagination.vue'
import { slugControl, slugType } from '@/utils/attributeLabels'
import { computed, onMounted, provide, ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAuth } from '~/composables/useAuth'
import { useConfirmDialog } from '~/composables/useConfirmDialog'
import { useNotifier } from '~/composables/useNotifier'

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

useHead({
  title: 'ایجاد گروه مشخصات فنی'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user, hasPermission } = useAuth()

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

interface Category { id: number; name: string }

const route = useRoute()
const router = useRouter()
interface ConfirmDialogOptions {
  title?: string
  message: string
  confirmText?: string
  cancelText?: string
  type?: 'danger' | 'warning' | 'info'
}
interface ConfirmDialogInstance {
  show: (options: ConfirmDialogOptions) => Promise<boolean>
}
const confirmDialogRef = ref<ConfirmDialogInstance | null>(null)
provide('confirmDialogRef', confirmDialogRef)

const groupId = computed(() => (route.params.id as string) || (route.query.id as string) || '')

// Pagination variables
const currentPage = ref(1)
const itemsPerPage = ref(10)

// Editing state
const editingAttributeId = ref<number | null>(null)
const editingAttribute = ref<Attribute | null>(null)

// Attribute list (initially خالی – کاربر با دکمه «افزودن ویژگی» اضافه می‌کند)
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
    title: 'تأیید حذف',
    message: 'آیا از حذف این گروه اطمینان دارید؟',
    confirmText: 'حذف',
    cancelText: 'انصراف',
    type: 'danger'
  })
  if (ok) router.push('/admin/product-management/attribute-groups')
}

const groupName = ref('')
const selectedCategory = ref<number | ''>('')
const categories = ref<Category[]>([])
const existingGroups = ref<{ id?: number; name: string; category_id: number }[]>([])

// نشان می‌دهد که آیا گروه حداقل یکبار ذخیره شده است یا خیر
const isSaved = ref(false)
const createdGroupId = ref<number | null>(null)

onMounted(async () => {
  try {
    const catData: any = await $fetch('/api/product-categories')
    categories.value = (catData as any[]).filter((c: any) => !c.parent_id)
  } catch (err) {
    // در صورت خطا، لیست خالی می‌ماند تا کاربر نتواند ذخیره کند
    categories.value = []
  }

  // Fetch existing groups list for duplicate checks
  try {
    const groupsRes: any = await $fetch('/api/attribute-groups', { params: { per_page: 1000, page: 1 } })
    existingGroups.value = Array.isArray(groupsRes.data) ? groupsRes.data : groupsRes
  } catch (err) {
    existingGroups.value = []
  }

  window.addEventListener('mousemove', drag)
  window.addEventListener('mouseup', endDrag)

  // بارگذاری لیست ویژگی‌ها
  try {
    const res: any = await $fetch('/api/attributes')
    availableAttributes.value = res.data || res
  } catch {}
})

const validateGroup = (): boolean => {
  if (!groupName.value.trim() || !selectedCategory.value) {
    useNotifier().warning('لطفاً نام گروه ویژگی و دسته‌بندی را وارد/انتخاب کنید')
    return false
  }

  const nameDup = existingGroups.value.some(g => g.name.trim() === groupName.value.trim())
  const categoryDup = existingGroups.value.some(g => g.category_id === Number(selectedCategory.value))

  if (nameDup && categoryDup) {
    useNotifier().warning('نام گروه ویژگی و دسته بندی گروه ویژگی قبلا استفاده شده است')
    return false
  } else if (nameDup) {
    useNotifier().warning('این نام گروه ویژگی قبلا استفاده شده است')
    return false
  } else if (categoryDup) {
    useNotifier().warning('این دسته بندی قبلا انتخاب شده است')
    return false
  }
  return true
}

// -------------------------
// Save helpers
// -------------------------
const createGroup = async () => {
  const payload = {
    name: groupName.value.trim(),
    category_id: Number(selectedCategory.value),
    description: ''
  }
  const res: any = await $fetch('/api/attribute-groups', {
    method: 'POST' as any,
    body: payload
  })
  return res
}

const updateGroup = async () => {
  if (!createdGroupId.value) return
  const payload = {
    name: groupName.value.trim(),
    description: ''
  }
  await $fetch(`/api/attribute-groups/${createdGroupId.value}`, {
    method: 'PUT' as any,
    body: payload
  })
}

const upsertGroupAttributes = async () => {
  if (!createdGroupId.value || attributes.value.length === 0) return
  const payload = {
    attributes: attributes.value.map(a => ({
      attribute_id: a.id,
      sort_order: a.display_order,
      type: slugType(a.type),
      control_type: slugControl(a.control_type),
      has_filter: a.has_filter,
      is_key: a.is_key,
      show_on_product: a.show_on_product,
      is_required: a.is_required
    }))
  }
  await $fetch(`/api/attribute-groups/${createdGroupId.value}/attributes`, {
    method: 'PUT' as any,
    body: payload
  })
}

const saveAndContinue = async () => {
  if (!validateGroup()) return

  try {
    const firstTimeSave = !isSaved.value

    if (firstTimeSave) {
      const res = await createGroup()
      createdGroupId.value = res.id || res.data?.id || null
    } else {
      await updateGroup()
    }

    await upsertGroupAttributes()

    // Mark as saved once operations succeed
    isSaved.value = true

    // If this was the first save (i.e. we just created the record), navigate
    // to the dedicated edit route so that a full-page refresh keeps the data.
    if (firstTimeSave && createdGroupId.value) {
      router.push(`/admin/product-management/attribute-groups/${createdGroupId.value}/edit`)
      return // Navigation will change page, no need to show alert here
    }

    useNotifier().success('تغییرات ذخیره شد')
  } catch (err) {
    console.error('خطا در ذخیره گروه مشخصات:', err)
    useNotifier().error('خطا در ذخیره گروه مشخصات')
  }
}

const saveAndExit = async () => {
  if (!validateGroup()) return
  try {
    if (!isSaved.value) {
      const res = await createGroup()
      createdGroupId.value = res.id || res.data?.id || null
    } else {
      await updateGroup()
    }

    await upsertGroupAttributes()

    router.push('/admin/product-management/attribute-groups')
  } catch (err) {
    console.error('خطا در ذخیره گروه مشخصات:', err)
    useNotifier().error('خطا در ذخیره گروه مشخصات')
  }
}

const addNewAttribute = () => {
  showAddModal.value = true
  modalPos.value = { top: window.innerHeight / 2, left: window.innerWidth / 2, centered: true }
}

const editAttribute = (attribute: Attribute) => {
  console.log('Editing attribute:', attribute)
  editingAttributeId.value = attribute.id
  editingAttribute.value = { ...attribute } // Create a copy to avoid direct mutation
}

const deleteAttribute = async (attributeId: number) => {
  const { confirm } = useConfirmDialog()
  const ok = await confirm({ title:'تأیید حذف', message:'آیا از حذف این مشخصه اطمینان دارید؟', confirmText:'حذف', cancelText:'انصراف', type:'danger' })
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
    console.log('Attribute saved:', editingAttribute.value)
  }
  editingAttributeId.value = null
  editingAttribute.value = null
}

const cancelEdit = () => {
  console.log('Edit canceled')
  editingAttributeId.value = null
  editingAttribute.value = null
}

// Pagination methods
const goToPage = (page: number) => {
  if (page >= 1 && page <= totalPages.value) {
    currentPage.value = page
  }
}

const handlePerPageChange = (val: number) => {
  itemsPerPage.value = val
  currentPage.value = 1
}

// State for modal
const showAddModal = ref(false)
const selectedAttributeId = ref<number | null>(null)

// لیست ویژگی‌های موجود از سرور بارگذاری می‌شود
const availableAttributes = ref<{id:number,name:string}[]>([])

const modalSearch = ref('')

// mapping for control types by attribute type
const CONTROL_TYPE_MAP: Record<string, string[]> = {
  'انتخاب': ['لیست باز شونده تک انتخابی', 'لیست باز شونده چند انتخابی'],
  'متن سفارشی': ['تکست باکس متنی تک خطی', 'تکست باکس متنی چند خطی']
}

const getDefaultControl = (type: string) => (CONTROL_TYPE_MAP[type] || [])[0] || ''

const controlOptions = (type: string) => CONTROL_TYPE_MAP[type] || []

// Ensure control_type always valid when type changes during edit
watch(
  () => editingAttribute.value?.type,
  (newType, _, onCleanup) => {
    if (editingAttribute.value) {
      const opts = CONTROL_TYPE_MAP[newType as string] || []
      if (!opts.includes(editingAttribute.value.control_type)) {
        editingAttribute.value.control_type = opts[0] || ''
      }
    }
  }
)

const selectedNames = computed(() => attributes.value.map(a => a.name))

const isAlreadySelected = (attr: {id:number,name:string}) => {
  return selectedNames.value.includes(attr.name)
}

const filteredModalAttributes = computed(() => {
  let arr = availableAttributes.value
  if (modalSearch.value) arr = arr.filter(a => a.name.includes(modalSearch.value))
  return arr
})

const addAttributeById = (id: number) => {
  selectedAttributeId.value = id
  confirmAddAttribute()
}

const confirmAddAttribute = () => {
  if (selectedAttributeId.value === null) return
  const attr = availableAttributes.value.find(a => a.id === selectedAttributeId.value)
  if (!attr) return
  // Prevent duplicates
  if (attributes.value.some(a => a.name === attr.name)) {
    useNotifier().warning('این مشخصه قبلاً اضافه شده است')
    return
  }
  attributes.value.push({
    id: Date.now(),
    name: attr.name,
    display_order: attributes.value.length + 1,
    type: 'انتخاب',
    control_type: getDefaultControl('انتخاب'),
    has_filter: false,
    is_key: false,
    show_on_product: true,
    is_required: false
  })
  selectedAttributeId.value = null
}

const closeModal = () => {
  showAddModal.value = false
  selectedAttributeId.value = null
}

// Drag and drop logic
const modalPos = ref<{ top: number; left: number; centered: boolean }>({ top: 0, left: 0, centered: true })

const modalStyle = computed(() => ({
  top: modalPos.value.top + 'px',
  left: modalPos.value.left + 'px',
  transform: modalPos.value.centered ? 'translate(-50%, -50%)' : ''
}))

const isDragging = ref(false)
const startDrag = () => {
  isDragging.value = true
  modalPos.value.centered = false
}

const drag = (event: MouseEvent) => {
  if (isDragging.value) {
    modalPos.value.left += event.movementX
    modalPos.value.top += event.movementY
  }
}

const endDrag = () => {
  isDragging.value = false
}

const promptNewAttribute = async () => {
  const name = prompt('نام ویژگی جدید را وارد کنید:')?.trim()
  if (!name) return
  // Check duplicate in available list
  if (availableAttributes.value.some(a => a.name === name)) {
    useNotifier().warning('این ویژگی قبلاً وجود دارد')
    return
  }
  try {
    const res: any = await $fetch('/api/attributes', {
      method: 'POST' as any,
      body: { name, display_name: name, data_type: 'text' }
    })
    // Add to local lists
    const newAttr = { id: res.id || res.data?.id, name }
    availableAttributes.value.push(newAttr)
    attributes.value.push({
      id: newAttr.id,
      name,
      display_order: attributes.value.length + 1,
      type: 'انتخاب',
      control_type: getDefaultControl('انتخاب'),
      has_filter: false,
      is_key: false,
      show_on_product: true,
      is_required: false
    })
  } catch (err: any) {
    console.error('Failed to create attribute on server', err)
  }
}

const finishButtonLabel = computed(() => attributes.value.length > 0 ? 'اتمام' : 'انصراف')

const finishButtonClass = computed(() => {
  return attributes.value.length > 0
    ? 'bg-gradient-to-r from-blue-500 to-blue-700'
    : 'bg-gradient-to-r from-red-500 to-red-700'
})

</script>

<style scoped>
.input-select-uniform {
  width: 16rem !important;
  min-width: 0;
  box-sizing: border-box;
  padding: 0.5rem 0.75rem;
  font-size: 0.875rem;
  border-radius: 0.5rem;
  border: 1px solid #d1d5db;
  background: #fff;
  appearance: none;
  direction: rtl;
}
</style> 

