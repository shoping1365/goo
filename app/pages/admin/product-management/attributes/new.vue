<template>
  <div class="bg-blue-50 min-h-screen">
    <!-- Header -->
    <div class="bg-white shadow-sm border-b border-gray-200">
      <div class="mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center py-6">
          <div class="flex items-center space-x-4 space-x-reverse">
            <button 
              class="p-2 rounded-md text-gray-600 hover:text-gray-900 hover:bg-gray-100"
              title="بازگشت"
              @click="goBack"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
              </svg>
            </button>
            <div>
              <h1 class="text-2xl font-bold text-gray-900">{{ pageTitle }}</h1>
              <p class="text-sm text-gray-600 mt-1">{{ formData.name ? 'ویرایش مشخصات فنی محصول' : 'ایجاد مشخصه فنی جدید برای محصولات' }}</p>
            </div>
          </div>
          <div class="flex gap-3">
            <button
              class="inline-flex items-center px-4 py-2 rounded-lg text-white bg-gradient-to-r from-purple-400 to-purple-600 hover:from-purple-500 hover:to-purple-700 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105 font-semibold"
              @click="saveAndContinueEdit"
            >
              ذخیره و ادامه ویرایش
            </button>
            <button
              class="inline-flex items-center px-4 py-2 rounded-lg text-white bg-gradient-to-r from-emerald-400 to-green-600 hover:from-emerald-500 hover:to-green-700 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105 font-semibold"
              @click="saveChanges"
            >
              ذخیره
            </button>
            <button
              v-if="formData.name"
              class="inline-flex items-center px-4 py-2 rounded-lg text-white bg-gradient-to-r from-red-400 to-red-600 hover:from-red-500 hover:to-red-700 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105 font-semibold"
              @click="deleteAttribute"
            >
              حذف
            </button>
            <button
              class="flex items-center px-4 py-2 border border-gray-300 rounded-md text-gray-600 hover:bg-gray-100 transition-colors font-semibold"
              @click="goBack"
            >
              {{ cancelLabel }}
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Content -->
    <div class="mx-auto px-4 sm:px-6 lg:px-8 py-8 space-y-6">

      <div class="space-y-6">

      <!-- اطلاعات اصلی ویژگی -->
      <div class="bg-white border border-blue-400 rounded shadow p-6 w-full mb-6">
        <div class="flex items-center justify-between cursor-pointer" @click="toggleSection('mainInfo')">
          <h3 class="text-sm font-semibold text-gray-700">اطلاعات ویژگی</h3>
          <span class="text-gray-500 text-lg">{{ expandedSections.mainInfo ? '−' : '+' }}</span>
        </div>
        <div v-show="expandedSections.mainInfo" class="mt-4">
          <div class="space-y-4">
            <div>
              <label class="block text-xs text-gray-700 font-semibold mb-2">نام ویژگی</label>
              <input
                v-model="formData.name"
                type="text"
                class="block w-full border border-gray-300 rounded-md shadow-sm px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                dir="rtl"
                placeholder="نام ویژگی را وارد کنید"
                @blur="!formData.displayText && (formData.displayText = formData.name)"
              />
            </div>
            <div>
              <label class="block text-xs text-gray-700 font-semibold mb-2">متن نمایشی</label>
              <input
                v-model="formData.displayText"
                type="text"
                class="block w-full border border-gray-300 rounded-md shadow-sm px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                dir="rtl"
                placeholder="متن نمایشی ویژگی"
              />
            </div>
            <div>
              <label class="block text-xs text-gray-700 font-semibold mb-2">واحد اندازه‌گیری</label>
              <select
                v-model="unitSelection"
                class="block w-full border border-gray-300 rounded-md shadow-sm px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="">بدون واحد</option>
                <option v-for="u in unitOptions" :key="u.value" :value="u.value">{{ u.label }}</option>
                <option value="_custom">سفارشی...</option>
                <option value="_manage">➕ / ✏️ مدیریت واحدها...</option>
              </select>
              <input
                v-if="unitSelection === '_custom'"
                v-model="customUnit"
                type="text"
                class="mt-2 block w-full border border-gray-300 rounded-md shadow-sm px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                placeholder="واحد دلخواه را وارد کنید"
              />
              <UnitManagerModal
                v-if="showUnitModal"
                :model-value="unitOptions"
                @close="showUnitModal=false"
                @saved="refreshUnits"
              />
            </div>
          </div>
        </div>
      </div>

      <!-- اطلاعات فنی و شناسایی -->
      <div class="bg-white border border-blue-400 rounded shadow p-6 w-full mb-6">
        <div class="flex items-center justify-between cursor-pointer" @click="toggleSection('technicalInfo')">
          <h3 class="text-sm font-semibold text-gray-700">اطلاعات فنی و شناسایی</h3>
          <span class="text-gray-500 text-lg">{{ expandedSections.technicalInfo ? '−' : '+' }}</span>
        </div>
        <div v-show="expandedSections.technicalInfo" class="mt-4 w-full">
          <!-- کانتینر جدا برای کد ویژگی و نوع داده -->
          <div class="bg-white border border-blue-200 rounded p-6 mt-6 flex flex-col md:flex-row gap-6">
            <div class="flex-1">
              <label class="block text-xs text-gray-700 font-semibold mb-2">کد ویژگی</label>
              <input
                v-model="formData.code"
                type="text"
                class="block w-full bg-gray-100 cursor-not-allowed border border-gray-300 rounded-md shadow-sm px-3 py-2 focus:outline-none"
                placeholder="ATTR-001"
                disabled
              />
              <div class="text-xs text-gray-500 mt-1">کد منحصر به فرد ویژگی</div>
            </div>
            <div class="flex-1">
              <label class="block text-xs text-gray-700 font-semibold mb-2">نوع داده</label>
              <div class="block w-full bg-gray-100 border border-gray-300 rounded-md px-3 py-2 text-sm text-gray-700">
                متن
              </div>
            </div>

            <!-- شناسه ویژگی -->
            <div class="flex-1">
              <label class="block text-xs text-gray-700 font-semibold mb-2">ID ویژگی</label>
              <input
                :value="editingId || '-'"
                type="text"
                class="block w-full bg-gray-100 cursor-not-allowed border border-gray-300 rounded-md shadow-sm px-3 py-2"
                disabled
              />
            </div>
          </div>
        </div>
      </div>

      </div>

      <!-- گزینه ها -->
      <div class="space-y-6">
        <div class="bg-white border border-blue-400 rounded shadow p-6 w-full">
          <div class="flex items-center justify-between cursor-pointer" @click="toggleSection('options')">
            <h3 class="text-sm font-semibold text-gray-700">گزینه ها</h3>
            <span class="text-gray-500 text-lg">{{ expandedSections.options ? '−' : '+' }}</span>
          </div>

          <div v-show="expandedSections.options" class="mt-4">
            <div class="flex justify-between items-center mb-4">
              <div class="text-sm text-gray-600">
                مقادیر مختلف این ویژگی
              </div>
              <button
                class="inline-flex items-center px-4 py-2 rounded-lg text-white bg-gradient-to-r from-emerald-400 to-green-600 hover:from-emerald-500 hover:to-green-700 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105 text-sm font-semibold"
                @click="openAddOptionModal"
              >
                افزودن گزینه جدید
              </button>
            </div>

            <!-- جدول گزینه ها -->
            <div class="overflow-x-auto">
              <table class="min-w-full divide-y divide-gray-200">
                <thead class="bg-gray-50">
                  <tr>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نام</th>
                   <th class="py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider pl-11 ">عملیات</th>
                  </tr>
                </thead>
                <tbody class="bg-white divide-y divide-gray-200">
                  <tr v-if="paginatedOptions.length === 0">
                    <td colspan="2" class="px-6 py-4 text-center text-sm text-gray-500">هیچ گزینه‌ای ثبت نشده است</td>
                  </tr>
                  <tr v-for="opt in paginatedOptions" :key="opt.id">
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ opt.name }}</td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900 text-left">
                      <div class="flex flex-row items-center space-x-2 space-x-reverse justify-end">
                        <button class="text-blue-600 hover:text-blue-900 w-max" title="ویرایش" @click="editOption(opt.name)">✏️ ویرایش</button>
                        <button class="text-red-600 hover:text-red-900 w-max" title="حذف" @click="deleteOption(opt.name)">🗑️ حذف</button>
                      </div>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>

            <!-- Pagination Component -->
            <Pagination
              :current-page="optionsPage"
              :total-pages="optionsTotalPages"
              :total="optionsTotal"
              :per-page="optionsPerPage"
              class="mt-4"
              @page-changed="handleOptionsPageChange"
              @per-page-changed="val => { optionsPerPage = val; optionsPage = 1 }"
            />
          </div>
        </div>

        <!-- توسط محصولات استفاده می شود -->
        <div class="bg-white border border-blue-400 rounded shadow p-6 w-full">
          <div class="flex items-center justify-between cursor-pointer" @click="toggleSection('usedByProducts')">
            <h3 class="text-sm font-semibold text-gray-700">توسط محصولات استفاده می شود</h3>
            <span class="text-gray-500 text-lg">{{ expandedSections.usedByProducts ? '−' : '+' }}</span>
          </div>

          <div v-show="expandedSections.usedByProducts" class="mt-4">
            <div class="overflow-x-auto">
              <table class="min-w-full divide-y divide-gray-200">
                <thead class="bg-gray-50">
                  <tr>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                      محتوا
                    </th>
                    <th class="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">
                      منتشر شده
                    </th>
                    <th class="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">
                      مشاهده
                    </th>
                  </tr>
                </thead>
                <tbody class="bg-white divide-y divide-gray-200">
                  <tr v-if="paginatedUsedProducts.length === 0">
                    <td colspan="3" class="px-6 py-4 text-center text-sm text-gray-500">هیچ محصولی یافت نشد</td>
                  </tr>
                  <tr v-for="prod in paginatedUsedProducts" :key="prod.id">
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ prod.name }}</td>
                    <td class="px-6 py-4 whitespace-nowrap text-center">
                      <input type="checkbox" :checked="prod.published" class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded" />
                    </td>
                    <td class="px-6 py-4 whitespace-nowrap text-center">
                      <button class="inline-flex items-center px-3 py-1 border border-gray-300 text-xs font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50" @click="viewProduct(prod.name)">👁️ مشاهده</button>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>

            <Pagination
              :current-page="usedProductsPage"
              :total-pages="usedProductsTotalPages"
              :total="usedProductsTotal"
              :per-page="usedProductsPerPage"
              class="mt-4 px-4"
              @page-changed="handleUsedProductsPageChange"
              @per-page-changed="val => { usedProductsPerPage = val; usedProductsPage = 1 }"
            />
          </div>
        </div>

        <!-- گروه های مشخصات فنی -->
        <div class="bg-white border border-blue-400 rounded shadow p-6 w-full">
          <div class="flex items-center justify-between cursor-pointer" @click="toggleSection('techSpecs')">
            <h3 class="text-sm font-semibold text-gray-700">گروه های مشخصات فنی</h3>
            <span class="text-gray-500 text-lg">{{ expandedSections.techSpecs ? '−' : '+' }}</span>
          </div>

          <div v-show="expandedSections.techSpecs" class="mt-4">
            <div class="text-sm text-gray-600">
              <p>گروه‌های مشخصات فنی اینجا قرار خواهد گرفت</p>
            </div>
          </div>
        </div>
      </div>
    </div>

  <!-- مودال افزودن گزینه جدید -->
  <transition name="fade-scale">
    <div v-if="showAddOptionModal" class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-40 backdrop-blur-sm">
      <div class="relative w-full max-w-md mx-auto p-0 md:p-0 animate-fade-in">
        <div class="bg-white rounded-2xl shadow-2xl border border-blue-200 px-6 py-7 md:px-8 md:py-8 transition-transform duration-200 hover:scale-[1.02]" dir="rtl">
          <!-- Close Button -->
          <button class="absolute left-4 top-6 text-gray-400 hover:text-red-500 transition-colors text-2xl focus:outline-none" @click="closeAddOptionModal">
            <span aria-hidden="true">×</span>
          </button>
          <!-- Title -->
          <div class="mb-6 text-center">
            <h3 class="text-2xl font-extrabold text-gray-800 tracking-tight">افزودن گزینه جدید</h3>
          </div>
          <hr class="mb-6 border-blue-100">
          <!-- Form -->
          <form class="space-y-6" @submit.prevent="saveNewOption">
            <div>
              <label class="block text-sm font-bold text-gray-700 mb-2">نام</label>
              <input
                v-model="newOption.name"
                type="text"
                class="block w-full border border-blue-200 bg-blue-50 rounded-lg shadow-sm px-4 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400 focus:border-blue-400 transition-all text-base placeholder-gray-400"
                dir="rtl"
                placeholder="نام گزینه را وارد کنید"
                autocomplete="off"
              />
            </div>
            <div class="flex flex-row-reverse gap-3 mt-8">
              <button
                type="submit"
                class="inline-flex items-center justify-center px-6 py-2 rounded-lg text-white bg-gradient-to-r from-blue-500 to-blue-700 hover:from-blue-600 hover:to-blue-800 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105 text-base font-bold gap-2"
              >
                <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7" /></svg>
                ذخیره
              </button>
              <button
                type="button"
                class="inline-flex items-center justify-center px-6 py-2 rounded-lg text-gray-700 bg-gradient-to-r from-gray-200 to-gray-300 hover:from-gray-300 hover:to-gray-400 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105 text-base font-bold gap-2"
                @click="closeAddOptionModal"
              >
                <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" /></svg>
                انصراف
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </transition>
</div>
</template>

<script setup lang="ts">
import { navigateTo } from '#imports'
import { computed, onMounted, ref, watch } from 'vue'
import Pagination from '~/components/admin/common/Pagination.vue'
import UnitManagerModal from '~/components/admin/modals/UnitManagerModal.vue'
import { useConfirmDialog } from '~/composables/useConfirmDialog'
import { useNotifier } from '~/composables/useNotifier'
import { definePageMeta } from '#imports'
import { useRoute } from 'vue-router'

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
// const { user, hasPermission } = useAuth()

// Get route for query parameters
const route = useRoute()

interface AttributeFormData {
  name: string
  displayText: string
  code: string
  unit: string
}

interface AttributeOption {
  id: number
  name: string
  hasColor: boolean
  colorValue: string | null
  colorName: string | null
}

interface UnitOption {
  value: string
  label: string
}

interface UsedProduct {
  id: number
  name: string
  published: boolean
}

interface RawAttributeValue {
  id: number | string
  value: string
  meta?: string | Record<string, unknown>
}

interface AttributeResponse {
  id: number
  name: string
  display_name: string
  data_type: string
  unit: string
  code?: string
  [key: string]: unknown
}

// Form data - will be populated dynamically
const formData = ref<AttributeFormData>({
  name: '',
  displayText: '',
  code: '',
  unit: '',
})

const unitSelection = ref('')
const customUnit = ref('')
const unitOptions = ref([
  { label: 'بدون واحد', value: '' },
  { label: 'سانتی‌متر (cm)', value: 'cm' },
  { label: 'کیلوگرم (kg)', value: 'kg' },
  { label: 'گرم (g)', value: 'g' },
  { label: 'لیتر (L)', value: 'L' },
  { label: 'میلی‌لیتر (ml)', value: 'ml' },
  { label: 'تعداد', value: 'count' },
  { label: 'سفارشی...', value: '_custom' }
])

// Expanded sections state with persistence in localStorage
const SECTION_KEY = 'attributeNewExpandedSections'

const defaultExpanded = {
  mainInfo: true,
  technicalInfo: true,
  options: true,
  usedByProducts: true,
  techSpecs: true
}

let initialExpanded = defaultExpanded

if (typeof window !== 'undefined') {
  try {
    const saved = localStorage.getItem(SECTION_KEY)
    if (saved) {
      initialExpanded = { ...defaultExpanded, ...JSON.parse(saved) }
    }
  } catch {
    // Failed to parse saved expanded sections
  }
}

const expandedSections = ref(initialExpanded)

// Modal state
const showAddOptionModal = ref(false)
const editingIndex = ref<number | null>(null) // null برای افزودن، اندیس برای ویرایش

// New option form data
const newOption = ref({
  name: '',
  hasColor: false,
  colorValue: '#000000',
  colorName: ''
})

// Options pagination
const options = ref<AttributeOption[]>([]) // will hold option objects when loaded
const originalOptionIds = ref(new Set<number>()) // نگه‌داری آیدی‌های اولیه گزینه‌ها
const optionsPage = ref(1)
const optionsPerPage = ref(10)

const optionsTotal = computed(() => options.value.length)
const optionsTotalPages = computed(() => Math.max(1, Math.ceil(optionsTotal.value / optionsPerPage.value)))

const paginatedOptions = computed(() => {
  const start = (optionsPage.value - 1) * optionsPerPage.value
  return options.value.slice(start, start + optionsPerPage.value)
})

const handleOptionsPageChange = (page: number) => {
  if (page >= 1 && page <= optionsTotalPages.value) {
    optionsPage.value = page
  }
}

// UsedByProducts pagination
const usedProducts = ref<UsedProduct[]>([])
const usedProductsPage = ref(1)
const usedProductsPerPage = ref(10)

const usedProductsTotal = computed(() => usedProducts.value.length)
const usedProductsTotalPages = computed(() => Math.max(1, Math.ceil(usedProductsTotal.value / usedProductsPerPage.value)))

const paginatedUsedProducts = computed(() => {
  const start = (usedProductsPage.value - 1) * usedProductsPerPage.value
  return usedProducts.value.slice(start, start + usedProductsPerPage.value)
})

const handleUsedProductsPageChange = (page: number) => {
  if (page >= 1 && page <= usedProductsTotalPages.value) {
    usedProductsPage.value = page
  }
}

// تبدیل کد رنگ به نام فارسی
const getColorName = (hexColor: string) => {
  const colorMap: Record<string, string> = {
    '#000000': 'مشکی',
    '#ffffff': 'سفید',
    '#ff0000': 'قرمز',
    '#00ff00': 'سبز',
    '#0000ff': 'آبی',
    '#ffff00': 'زرد',
    '#ff00ff': 'بنفش',
    '#00ffff': 'فیروزه‌ای',
    '#ffa500': 'نارنجی',
    '#800080': 'ارغوانی',
    '#ffc0cb': 'صورتی',
    '#a52a2a': 'قهوه‌ای',
    '#808080': 'خاکستری',
    '#008000': 'سبز تیره',
    '#000080': 'آبی تیره',
    '#800000': 'قرمز تیره'
  }

  return colorMap[hexColor.toLowerCase()] || hexColor
}

// تابع نرمال‌سازی برای تشخیص تکراری بودن
const normalizeValue = (str: string) => String(str).trim().replace(/\s+/g, ' ').toLowerCase()

// نظارت بر تغییر رنگ
watch(() => newOption.value.colorValue, (newColor) => {
  if (newOption.value.hasColor) {
    newOption.value.colorName = getColorName(newColor)
  }
})

// Computed property for dynamic title
const pageTitle = computed(() => {
  // Only show 'ویرایش ویژگی' if editing (id exists)
  if (route.query.id) {
    return formData.value.name ? `ویرایش ویژگی: ${formData.value.name}` : 'ویرایش ویژگی'
  }
  return 'ایجاد ویژگی جدید'
})

// Methods
const goBack = () => {
  // Going back to attributes list
  navigateTo('/admin/product-management/attributes')
}

const DRAFT_KEY = 'attributeNewDraft'

// Load draft on mount
onMounted(async () => {
  if (typeof window === 'undefined') return
  try {
    const draft = localStorage.getItem(DRAFT_KEY)
    if (draft) {
      const parsed = JSON.parse(draft)
      if (parsed.formData) Object.assign(formData.value, parsed.formData)
      if (Array.isArray(parsed.options)) options.value = parsed.options
    }
  } catch {
    // Failed to load attribute draft
  }

  // If coming from edit route with ?name=foo, prefill
  if (route.query.name && !formData.value.name) {
    formData.value.name = String(route.query.name)
  }

  // If coming from edit route with ?id=123, fetch details
  if (route.query.id) {
    try {
      const attrId = route.query.id
      const attribute = await $fetch<AttributeResponse>(`/api/attributes/${attrId}`)
      if (attribute) {
        formData.value.name = attribute.name || ''
        formData.value.displayText = attribute.display_name || ''
        formData.value.code = attribute.code || ''
        formData.value.unit = attribute.unit || ''

        unitSelection.value = attribute.unit || ''
        if (unitSelection.value && !unitOptions.value.some(u=>u.value===unitSelection.value)) {
          unitSelection.value = '_custom'
          customUnit.value = attribute.unit
        }
      }

      await fetchAttributeValues(String(attrId))
    } catch {
      // Failed to load attribute for edit
    }
  }
})

// Persist draft whenever formData or options change
watch([formData, options], () => {
  if (typeof window === 'undefined') return
  const draft = {
    formData: formData.value,
    options: options.value
  }
  try {
    localStorage.setItem(DRAFT_KEY, JSON.stringify(draft))
  } catch {
    // Cannot save draft to localStorage
  }
}, { deep: true })

const clearDraft = () => {
  if (typeof window !== 'undefined') localStorage.removeItem(DRAFT_KEY)
}

const showToast = (msg: string, type: 'success' | 'error' | 'warning' | 'info' = 'success') => {
  if (type === 'error') useNotifier().error(msg)
  else if (type === 'warning') useNotifier().warning(msg)
  else useNotifier().success(msg)
}

const editingId = computed(() => route.query.id || null)

const preparePayload = () => ({
  name: formData.value.name?.trim(),
  display_name: formData.value.displayText?.trim(),
  code: formData.value.code?.trim(),
  unit: formData.value.unit?.trim(),
  sort_order: 0,
  is_required: false,
  is_filterable: false,
  is_active: true
})

const fetchAttributeValues = async (attrId: string) => {
  try {
    const vals = await $fetch<RawAttributeValue[]>(`/api/attribute-values/by-attribute/${attrId}`)
    if (Array.isArray(vals)) {
      options.value = vals.map(v => {
        let metaObj: Record<string, unknown> = {}
        if (v.meta) {
          try { metaObj = typeof v.meta === 'string' ? JSON.parse(v.meta) : v.meta as Record<string, unknown> } catch { metaObj = {} }
        }
        return {
          id: Number(v.id),
          name: v.value,
          hasColor: !!metaObj.color,
          colorValue: (metaObj.color as string) || '#000000',
          colorName: (metaObj.color_name as string) || ''
        }
      })
      originalOptionIds.value = new Set(vals.map(v => Number(v.id)))
    }
  } catch {
    // Failed to fetch attribute values
  }
}

const syncOptions = async (attrId: string | number) => {
  if (!attrId) return

  const currentIds = new Set<number>()

  for (const [idx, opt] of options.value.entries()) {
    const metaObj: Record<string, unknown> = {}
    if (opt.hasColor) {
      metaObj.color = opt.colorValue
      if (opt.colorName) metaObj.color_name = opt.colorName
    }

    // تابع داخلی برای ساخت slug سازگار بدون وابستگی خارجی
    const makeSlug = (s: string) => {
      if (!s) return ''
      return String(s)
        .normalize('NFKD')
        .replace(/[\u064B-\u065F]/g, '') // حذف اعراب عربی
        .replace(/[^\p{L}\p{N}]+/gu, '-') // هر چیزی به جز حروف/اعداد → '-'
        .replace(/^-+|-+$/g, '')
        .toLowerCase()
    }

    const payload = {
      value: opt.name,
      sort_order: idx + 1,
      slug: makeSlug(opt.name),
      meta: Object.keys(metaObj).length ? JSON.stringify(metaObj) : undefined
    }

    if (opt.id != null && originalOptionIds.value.has(Number(opt.id))) {
      // update existing
      currentIds.add(opt.id)
      await $fetch(`/api/attribute-values/${opt.id}`, { method: 'PUT', body: payload })
    } else {
      const created = await $fetch<{ id: number }>(`/api/attribute-values/by-attribute/${attrId}`, { method: 'POST', body: payload })
      if (created?.id) {
        opt.id = Number(created.id)
        currentIds.add(Number(created.id))
      }
    }
  }

  // deletions if editing existing attribute
  if (editingId.value) {
    for (const oldId of originalOptionIds.value) {
      if (!currentIds.has(oldId)) {
        try {
          await $fetch(`/api/attribute-values/${oldId}`, { method: 'DELETE' })
        } catch {
          // Failed to delete option
        }
      }
    }
    originalOptionIds.value = currentIds
  }

  // --- sync from backend ---
  await fetchAttributeValues(String(attrId))
}

const savedContinue = ref(false)
const cancelLabel = computed(() => savedContinue.value ? 'بازگشت' : 'انصراف')

const saveChanges = async () => {
  try {
    const payload = preparePayload()
    const url = editingId.value ? `/api/attributes/${editingId.value}` : '/api/attributes'
    const method = editingId.value ? 'PUT' : 'POST'
    const resp = await $fetch<{ id: number }>(url, { method, body: payload })

    const attrId = editingId.value || resp?.id
    await syncOptions(Array.isArray(attrId) ? attrId[0] : attrId)

    clearDraft()
    showToast('✅ ویژگی با موفقیت ایجاد شد')
    navigateTo('/admin/product-management/attributes')
  } catch (err: unknown) {
    const e = err as { data?: { message?: string, msg?: string, error?: string }, message?: string }
    const backendMsg = e?.data?.message || e?.data?.msg || e?.data?.error
    const msg = backendMsg || e?.message || 'خطا در ذخیره ویژگی'
    showToast(msg, 'error')
  }
}

const saveAndContinueEdit = async () => {
  try {
    const payload = preparePayload()
    const url = editingId.value ? `/api/attributes/${editingId.value}` : '/api/attributes'
    const method = editingId.value ? 'PUT' : 'POST'
    const resp = await $fetch<{ id: number }>(url, { method, body: payload })
    const attrId = editingId.value || resp?.id
    await syncOptions(Array.isArray(attrId) ? attrId[0] : attrId)

    savedContinue.value = true
    clearDraft()
    showToast('✅ ویژگی ذخیره شد، می‌توانید ادامه دهید')
    // optionally set formData.id etc.
  } catch (err: unknown) {
    const e = err as { data?: { message?: string, msg?: string, error?: string }, message?: string }
    const backendMsg = e?.data?.message || e?.data?.msg || e?.data?.error
    const msg = backendMsg || e?.message || 'خطا در ذخیره ویژگی'
    showToast(msg, 'error')
  }
}

const deleteAttribute = async () => {
  if (!editingId.value) return
  const { confirm } = useConfirmDialog()
  const ok = await confirm({ title:'تأیید حذف', message:'آیا از حذف این ویژگی اطمینان دارید؟', confirmText:'حذف', cancelText:'انصراف', type:'danger' })
  if (ok) {
    try {
      await $fetch(`/api/attributes/${editingId.value}`, { method: 'DELETE' })
      showToast('🗑️ ویژگی با موفقیت حذف شد!')
      clearDraft()
      navigateTo('/admin/product-management/attributes')
    } catch {
      showToast('خطا در حذف ویژگی', 'error')
    }
  }
}

const toggleSection = (section: keyof typeof expandedSections.value) => {
  expandedSections.value[section] = !expandedSections.value[section]
  localStorage.setItem(SECTION_KEY, JSON.stringify(expandedSections.value))
}

const openAddOptionModal = () => {
  editingIndex.value = null
  newOption.value = {
    name: '',
    hasColor: false,
    colorValue: '#000000',
    colorName: ''
  }
  showAddOptionModal.value = true
}

const closeAddOptionModal = () => {
  showAddOptionModal.value = false
}

const saveNewOption = () => {
  const name = newOption.value.name.trim()
  if (!name) {
    useNotifier().warning('لطفاً نام گزینه را وارد کنید')
    return
  }

  // تکراری نباشد
  const normName = normalizeValue(name)
  const duplicate = options.value.some((o, idx) => idx !== editingIndex.value && normalizeValue(o.name) === normName)
  if (duplicate) {
    useNotifier().warning('این مقدار قبلاً ثبت شده است')
    return
  }

  const optionObj: AttributeOption = {
    id: editingIndex.value !== null ? options.value[editingIndex.value].id : Date.now(),
    name,
    hasColor: newOption.value.hasColor,
    colorValue: newOption.value.hasColor ? newOption.value.colorValue : null,
    colorName: newOption.value.hasColor ? newOption.value.colorName : null
  }

  if (editingIndex.value !== null) {
    options.value[editingIndex.value] = optionObj
  } else {
    options.value.push(optionObj)
  }

  // reset form
  newOption.value = {
    name: '',
    hasColor: false,
    colorValue: '#000000',
    colorName: ''
  }

  showAddOptionModal.value = false
}

const deleteOption = async (optionName: string) => {
  const { confirm } = useConfirmDialog()
  const ok = await confirm({ title:'تأیید حذف', message:`آیا از حذف گزینه "${optionName}" اطمینان دارید؟`, confirmText:'حذف', cancelText:'انصراف', type:'danger' })
  if (ok) {
    options.value = options.value.filter(o => o.name !== optionName)
  }
}

const editOption = (optionName: string) => {
  const index = options.value.findIndex(o => o.name === optionName)
  if (index === -1) return
  editingIndex.value = index
  const opt = options.value[index]
  newOption.value = {
    name: opt.name,
    hasColor: opt.hasColor,
    colorValue: opt.colorValue || '#000000',
    colorName: opt.colorName || ''
  }
  showAddOptionModal.value = true
}

const viewProduct = (productName: string) => {
  useNotifier().info(`مشاهده محصول "${productName}" (نمونه)`) // TODO: implement
}

// in <script> section additions
const DEFAULT_UNITS: UnitOption[] = [
  { value: 'kg', label: 'کیلوگرم' },
  { value: 'g', label: 'گرم' },
  { value: 'mg', label: 'میلی‌گرم' },
  { value: 'l', label: 'لیتر' },
  { value: 'ml', label: 'میلی‌لیتر' },
  { value: 'm', label: 'متر' },
  { value: 'cm', label: 'سانتی‌متر' },
  { value: 'mm', label: 'میلی‌متر' },
  { value: 'in', label: 'اینچ' },
  { value: 'ft', label: 'فوت' },
  { value: '°C', label: 'درجهٔ سلسیوس' },
  { value: '°F', label: 'درجهٔ فارنهایت' },
  { value: 'V', label: 'ولت' },
  { value: 'W', label: 'وات' },
  { value: 'kW', label: 'کیلووات' },
  { value: 'Ah', label: 'آمپر ساعت' },
  { value: 'mAh', label: 'میلی‌آمپر ساعت' },
  { value: 'MP', label: 'مگاپیکسل' },
  { value: 'Hz', label: 'هرتز' },
  { value: 'ppm', label: 'قسمت در میلیون (ppm)' }
]

const loadUnits = () => {
  try {
    const saved = localStorage.getItem('measurementUnits')
    unitOptions.value = saved ? JSON.parse(saved) : DEFAULT_UNITS
  } catch {
    unitOptions.value = DEFAULT_UNITS
  }
}
const refreshUnits = () => {
  loadUnits()
  if (unitSelection.value && !unitOptions.value.some(u=>u.value===unitSelection.value)) {
    unitSelection.value = ''
    formData.value.unit = ''
  }
}
onMounted(() => { loadUnits() })

const showUnitModal = ref(false)

watch(unitSelection, (val) => {
  if (val === '_custom') {
    formData.value.unit = customUnit.value
  } else if (val === '_manage') {
    unitSelection.value = formData.value.unit || ''
    showUnitModal.value = true
  } else {
    formData.value.unit = val
  }
})

</script>

