<template>
  <div class="bg-blue-50 min-h-screen">
    <!-- Header -->
    <div class="bg-white shadow-sm border-b border-gray-200">
      <div class="mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center py-6">
          <div class="flex items-center space-x-4 space-x-reverse">
            <button
              @click="goBack"
              class="p-2 rounded-md text-gray-600 hover:text-gray-900 hover:bg-gray-100"
              title="بازگشت"
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
              @click="saveAndContinueEdit"
              class="inline-flex items-center px-4 py-2 rounded-lg text-white bg-gradient-to-r from-purple-400 to-purple-600 hover:from-purple-500 hover:to-purple-700 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105 font-semibold"
            >
              ذخیره و ادامه ویرایش
            </button>
            <button
              @click="saveChanges"
              class="inline-flex items-center px-4 py-2 rounded-lg text-white bg-gradient-to-r from-emerald-400 to-green-600 hover:from-emerald-500 hover:to-green-700 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105 font-semibold"
            >
              ذخیره
            </button>
            <button
              v-if="formData.name"
              @click="deleteAttribute"
              class="inline-flex items-center px-4 py-2 rounded-lg text-white bg-gradient-to-r from-red-400 to-red-600 hover:from-red-500 hover:to-red-700 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105 font-semibold"
            >
              حذف
            </button>
            <button
              @click="goBack"
              class="flex items-center px-4 py-2 border border-gray-300 rounded-md text-gray-600 hover:bg-gray-100 transition-colors font-semibold"
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
                :disabled="isColorAttr"
                class="block w-full border border-gray-300 rounded-md shadow-sm px-3 py-2 focus:outline-none focus:ring-blue-500 focus:border-blue-500 bg-gray-100"
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
              <!-- مدیریت واحدها -->
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
                {{ dataTypeLabel }}
              </div>
            </div>

            <!-- شناسه ویژگی -->
            <div class="flex-1">
              <label class="block text-xs text-gray-700 font-semibold mb-2">ID ویژگی</label>
              <input
                :value="attributeId || '-'"
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
                @click="openAddOptionModal"
                class="inline-flex items-center px-4 py-2 rounded-lg text-white bg-gradient-to-r from-emerald-400 to-green-600 hover:from-emerald-500 hover:to-green-700 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105 text-sm font-semibold"
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
                    <th v-if="isColorAttr" class="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">رنگ</th>
                    <th class="py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider pr-4">عملیات</th>
                  </tr>
                </thead>
                <tbody class="bg-white divide-y divide-gray-200">
                  <tr v-if="paginatedOptions.length === 0">
                    <td :colspan="isColorAttr ? 3 : 2" class="px-6 py-4 text-center text-sm text-gray-500">هیچ گزینه‌ای ثبت نشده است</td>
                  </tr>
                  <tr v-for="opt in paginatedOptions" :key="opt.id">
                    <!-- نام -->
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ opt.name }}</td>

                    <!-- رنگ -->
                    <td v-if="isColorAttr" class="px-6 py-4 whitespace-nowrap text-center">
                      <span class="inline-block w-6 h-6 rounded border border-gray-300 align-middle" :style="`background:${opt.colorValue}`"></span>
                      <code class="text-xs text-gray-600 mr-2">{{ opt.colorValue }}</code>
                    </td>

                    <!-- عملیات -->
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900 text-left">
                      <div class="flex flex-row items-center space-x-2 space-x-reverse justify-end">
                        <button @click="editOption(opt.name)" class="text-blue-600 hover:text-blue-900 w-max" title="ویرایش">✏️ ویرایش</button>
                        <button @click="deleteOption(opt.name)" class="text-red-600 hover:text-red-900 w-max" title="حذف">🗑️ حذف</button>
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
              @page-changed="handleOptionsPageChange"
              @per-page-changed="val => { optionsPerPage = val; optionsPage = 1 }"
              class="mt-4"
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
                      <button @click="viewProduct(prod.name)" class="inline-flex items-center px-3 py-1 border border-gray-300 text-xs font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50">👁️ مشاهده</button>
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
              @page-changed="handleUsedProductsPageChange"
              @per-page-changed="val => { usedProductsPerPage = val; usedProductsPage = 1 }"
              class="mt-4 px-4"
            />
          </div>
        </div>

        <!-- گروه های مشخصات فنی -->
        <div class="bg-white border border-blue-400 rounded shadow p-6 w-full">
          <div class="flex items-center justify-between cursor-pointer" @click="toggleSection('attrGroups')">
            <h3 class="text-sm font-semibold text-gray-700">گروه های مشخصات فنی</h3>
            <span class="text-gray-500 text-lg">{{ expandedSections.attrGroups ? '−' : '+' }}</span>
          </div>

          <div v-show="expandedSections.attrGroups" class="mt-4">
            <table class="min-w-full divide-y divide-gray-200" v-if="groups.length">
              <thead class="bg-gray-50">
                <tr>
                  <th class="px-4 py-2 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نام گروه</th>
                  <th class="px-4 py-2 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">دسته‌بندی</th>
                  <th class="px-4 py-2 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">مشاهده</th>
                </tr>
              </thead>
              <tbody class="bg-white divide-y divide-gray-200">
                <tr v-for="g in groups" :key="g.id">
                  <td class="px-4 py-3 text-sm text-gray-900 text-right">{{ g.name }}</td>
                  <td class="px-4 py-3 text-sm text-gray-700 text-right">{{ g.category?.name || '—' }}</td>
                  <td class="px-4 py-3 text-center">
                    <NuxtLink :to="`/admin/attribute-groups/${g.id}/edit`" class="text-blue-600 hover:text-blue-800 text-xs underline">ویرایش گروه</NuxtLink>
                  </td>
                </tr>
              </tbody>
            </table>
            <p v-else class="text-center text-sm text-gray-500 py-6">این ویژگی در هیچ گروهی قرار ندارد.</p>
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
          <button @click="closeAddOptionModal" class="absolute left-4 top-6 text-gray-400 hover:text-red-500 transition-colors text-2xl focus:outline-none">
            <span aria-hidden="true">×</span>
          </button>
          <!-- Title -->
          <div class="mb-6 text-center">
            <h3 class="text-2xl font-extrabold text-gray-800 tracking-tight">افزودن گزینه جدید</h3>
          </div>
          <hr class="mb-6 border-blue-100">
          <!-- Form -->
          <form @submit.prevent="saveNewOption" class="space-y-6">
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
            <div>
              <label class="inline-flex items-center cursor-pointer gap-2">
                <input
                  id="has-color"
                  type="checkbox"
                  v-model="newOption.hasColor"
                  class="form-checkbox accent-blue-500 w-5 h-5 rounded border-gray-300 focus:ring-2 focus:ring-blue-400 transition-all"
                />
                <span class="text-sm font-bold text-gray-700">انتخاب رنگ</span>
              </label>
            </div>
            <transition name="fade">
              <div v-if="newOption.hasColor" class="space-y-4">
                <div>
                  <label class="block text-sm font-bold text-gray-700 mb-2">انتخاب رنگ</label>
                  <div class="flex items-center gap-6">
                    <input
                      v-model="newOption.colorValue"
                      type="color"
                      class="h-10 w-16 border border-gray-300 rounded-lg cursor-pointer shadow-sm"
                    />
                    <span class="text-base text-gray-600 font-mono">{{ newOption.colorValue }}</span>
                  </div>
                </div>
                <div>
                  <label class="block text-sm font-bold text-gray-700 mb-2">نام رنگ</label>
                  <input
                    v-model="newOption.colorName"
                    type="text"
                    class="block w-full border border-blue-200 bg-blue-50 rounded-lg shadow-sm px-4 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400 focus:border-blue-400 transition-all text-base placeholder-gray-400"
                    dir="rtl"
                    placeholder="نام رنگ (مثال: قرمز، آبی)"
                  />
                </div>
              </div>
            </transition>
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
                @click="closeAddOptionModal"
                class="inline-flex items-center justify-center px-6 py-2 rounded-lg text-gray-700 bg-gradient-to-r from-gray-200 to-gray-300 hover:from-gray-300 hover:to-gray-400 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105 text-base font-bold gap-2"
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

<script setup>
import { ref, watch, computed, onMounted } from 'vue'
import Pagination from '~/components/admin/common/Pagination.vue'
import { navigateTo } from '#app'
import UnitManagerModal from '~/components/admin/modals/UnitManagerModal.vue'
// بارگذاری داینامیک slugify با فول‌بک داخلی (بدون TypeScript annotation)
let slugify = null
try {
  const mod = await import('slugify')
  slugify = (mod && (mod.default || mod.slugify)) ? (mod.default || mod.slugify) : mod
} catch (_) {
  slugify = null
}
const makeSlugFallback = (s, replacement = '-') => {
  if (!s) return ''
  return String(s)
    .normalize('NFKD')
    .replace(/[\u064B-\u065F]/g, '')
    .replace(/[^\p{L}\p{N}]+/gu, ' ')
    .trim()
    .replace(/\s+/g, replacement)
    .toLowerCase()
}
const toSlug = (s, replacement = '-') => {
  try { return slugify ? slugify(s, { lower: true, strict: true, replacement }) : makeSlugFallback(s, replacement) } catch { return makeSlugFallback(s, replacement) }
}

definePageMeta({
  layout: 'admin-main'
})

// Get route for query parameters
const route = useRoute()
const attributeId = computed(() => route.params.id)

// Form data - will be populated dynamically
const formData = ref({
  name: '',
  displayText: '',
  code: '',
  dataType: 'auto',
  unit: '',
  // remove obsolete fields
})

// Expanded sections state with persistence in localStorage
const SECTION_KEY = 'attributeNewExpandedSections'

const defaultExpanded = {
  mainInfo: true,
  technicalInfo: true,
  options: true,
  usedByProducts: true,
  attrGroups: true
}

let initialExpanded = defaultExpanded

if (typeof window !== 'undefined') {
  try {
    const saved = localStorage.getItem(SECTION_KEY)
    if (saved) {
      initialExpanded = { ...defaultExpanded, ...JSON.parse(saved) }
    }
  } catch (e) {
    console.error('Failed to parse saved expanded sections', e)
  }
}

const expandedSections = ref(initialExpanded)

// Modal state
const showAddOptionModal = ref(false)
const editingIndex = ref(null) // null برای افزودن، اندیس برای ویرایش

// New option form data
const newOption = ref({
  name: '',
  hasColor: false,
  colorValue: '#000000',
  colorName: ''
})

// Options pagination
const options = ref([]) // will hold option objects when loaded
const originalOptionIds = ref(new Set()) // نگه‌داری آیدی‌های اولیه گزینه‌ها برای تشخیص حذف‌ها
const optionsPage = ref(1)
const optionsPerPage = ref(10)

const optionsTotal = computed(() => options.value.length)
const optionsTotalPages = computed(() => Math.max(1, Math.ceil(optionsTotal.value / optionsPerPage.value)))

const paginatedOptions = computed(() => {
  const start = (optionsPage.value - 1) * optionsPerPage.value
  return options.value.slice(start, start + optionsPerPage.value)
})

const handleOptionsPageChange = (page) => {
  if (page >= 1 && page <= optionsTotalPages.value) {
    optionsPage.value = page
  }
}

// UsedByProducts pagination
const usedProducts = ref([])
const usedProductsPage = ref(1)
const usedProductsPerPage = ref(10)

const usedProductsTotal = computed(() => usedProducts.value.length)
const usedProductsTotalPages = computed(() => Math.max(1, Math.ceil(usedProductsTotal.value / usedProductsPerPage.value)))

const paginatedUsedProducts = computed(() => {
  const start = (usedProductsPage.value - 1) * usedProductsPerPage.value
  return usedProducts.value.slice(start, start + usedProductsPerPage.value)
})

const handleUsedProductsPageChange = (page) => {
  if (page >= 1 && page <= usedProductsTotalPages.value) {
    usedProductsPage.value = page
  }
}

// تبدیل کد رنگ به نام فارسی
const getColorName = (hexColor) => {
  const colorMap = {
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

// نرمال‌سازی متن برای مقایسهٔ دقیق (حذف فاصله‌های اضافی و حروف بزرگ/کوچک)
const normalizeValue = (str) => String(str).trim().replace(/\s+/g, ' ').toLowerCase()

// نظارت بر تغییر رنگ
watch(() => newOption.value.colorValue, (newColor) => {
  if (newOption.value.hasColor) {
    newOption.value.colorName = getColorName(newColor)
  }
})

// Computed property for dynamic title
const pageTitle = computed(() => {
  return attributeId.value && formData.value.name
    ? `ویرایش ویژگی: ${formData.value.name}`
    : 'ایجاد ویژگی جدید'
})

// آیا این ویژگی از نوع رنگ است؟
const isColorAttr = computed(() => formData.value.dataType === 'color')

// Methods
const goBack = () => {
  console.log('🔙 Going back to attributes list')
  navigateTo('/admin/product-management/attributes')
}

const DRAFT_KEY = 'attributeNewDraft'

// Load draft on mount
onMounted(async () => {
  if (typeof window === 'undefined') return
  
  loadUnits()

  // If editing (route.params.id exists), fetch details FIRST
  if (attributeId.value) {
    const attrId = attributeId.value

    // دریافت خودِ ویژگی برای پرکردن فرم
    try {
      const attribute = await $fetch(`/api/attributes/${attrId}`)
      if (attribute) {
        formData.value.name = attribute.name || ''
        formData.value.displayText = attribute.display_name || ''
        formData.value.code = attribute.code || ''
        formData.value.dataType = attribute.data_type || 'auto'
        formData.value.unit = attribute.unit || ''

        // هماهنگی انتخاب واحد
        unitSelection.value = attribute.unit || ''
        if (unitSelection.value && !unitOptions.value.some(u => u.value === unitSelection.value)) {
          unitSelection.value = '_custom'
          customUnit.value = attribute.unit
        }
      }
    } catch (e) {
      console.error('Failed to load attribute', e)
    }

    await fetchAttributeValues(attrId)
    await loadAttributeGroups()
  } else {
    // Only load draft for NEW attributes, not for editing
    try {
      const draft = localStorage.getItem(DRAFT_KEY)
      if (draft) {
        const parsed = JSON.parse(draft)
        if (parsed.formData) Object.assign(formData.value, parsed.formData)
        if (Array.isArray(parsed.options)) options.value = parsed.options
      }
    } catch (e) {
      console.error('Failed to load attribute draft', e)
    }

    // If coming from edit route with ?name=foo, prefill
    if (route.query.name && !formData.value.name) {
      formData.value.name = String(route.query.name)
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
  } catch (e) {
    console.warn('Cannot save draft to localStorage', e)
  }
}, { deep: true })

const clearDraft = () => {
  if (typeof window !== 'undefined') localStorage.removeItem(DRAFT_KEY)
}

const showToast = (msg, type = 'success') => {
  // You can replace with your own toast library
  alert(msg)
}

const editingId = computed(() => attributeId.value || null)

const preparePayload = () => ({
  name: formData.value.name?.trim(),
  display_name: formData.value.displayText?.trim(),
  code: formData.value.code?.trim(),
  data_type: formData.value.dataType || 'text',
  unit: formData.value.unit?.trim(),
  sort_order: 0,
  is_required: false,
  is_filterable: false,
  is_active: true
})

const syncOptions = async (attrId) => {
  if (!attrId) return

  // مجموعه آیدی‌های فعلی پس از ذخیره
  const currentIds = new Set()

  for (const [idx, opt] of options.value.entries()) {
    // آماده‌سازی meta (در حال حاضر فقط رنگ)
    const metaObj = {}
    if (opt.hasColor) {
      metaObj.color = opt.colorValue
      if (opt.colorName) metaObj.color_name = opt.colorName
    }

    const payload = {
      value: opt.name,
      sort_order: idx + 1,
      slug: toSlug(opt.name, '-'),
      meta: Object.keys(metaObj).length ? JSON.stringify(metaObj) : undefined
    }

    if (opt.id != null && originalOptionIds.value.has(Number(opt.id))) {
      // گزینه موجود – به‌روزرسانی
      currentIds.add(opt.id)
      await $fetch(`/api/attribute-values/${opt.id}`, { method: 'PUT', body: payload })
    } else {
      // گزینه جدید – ایجاد
      const created = await $fetch(`/api/attribute-values/by-attribute/${attrId}`, { method: 'POST', body: payload })
      if (created?.id) {
        opt.id = Number(created.id)
        currentIds.add(Number(created.id))
      }
    }
  }

  // گزینه‌هایی که حذف شده‌اند
  for (const oldId of originalOptionIds.value) {
    if (!currentIds.has(oldId)) {
      try {
        await $fetch(`/api/attribute-values/${oldId}`, { method: 'DELETE' })
      } catch (e) {
        console.warn('Failed to delete option', oldId, e)
      }
    }
  }

  // به‌روزرسانی نسخه اصلی برای دفعات بعدی
  originalOptionIds.value = currentIds

  // --- دریافت لیست تازه از سرور برای همگام‌سازی کامل ---
  await fetchAttributeValues(attrId)
  await fetchAttribute()
}

const savedContinue = ref(false)
const cancelLabel = computed(() => savedContinue.value ? 'بازگشت' : 'انصراف')

const saveChanges = async () => {
  try {
    const payload = preparePayload()
    const url = editingId.value ? `/api/attributes/${editingId.value}` : '/api/attributes'
    const method = editingId.value ? 'PUT' : 'POST'
    const resp = await $fetch(url, { method, body: payload })

    const attrId = editingId.value || resp?.id

    // همگام‌سازی گزینه‌ها
    await syncOptions(attrId)

    clearDraft()
    showToast('✅ ویژگی با موفقیت ایجاد شد')
    navigateTo('/admin/product-management/attributes')
  } catch (err) {
    console.error('Save error', err)
    const msg = err?.data?.error || err?.message || 'خطا در ذخیره ویژگی'
    showToast(msg, 'error')
  }
}

const saveAndContinueEdit = async () => {
  try {
    const payload = preparePayload()
    const url = editingId.value ? `/api/attributes/${editingId.value}` : '/api/attributes'
    const method = editingId.value ? 'PUT' : 'POST'
    const resp = await $fetch(url, { method, body: payload })

    const attrId = editingId.value || resp?.id
    await syncOptions(attrId)

    savedContinue.value = true
    clearDraft()
    showToast('✅ ویژگی ذخیره شد، می‌توانید ادامه دهید')
    // optionally set formData.id etc.
  } catch (err) {
    console.error('Save error', err)
    const msg = err?.data?.error || err?.message || 'خطا در ذخیره ویژگی'
    showToast(msg, 'error')
  }
}

const deleteAttribute = async () => {
  if (!editingId.value) return
  if (confirm('آیا از حذف این ویژگی اطمینان دارید؟')) {
    try {
      await $fetch(`/api/attributes/${editingId.value}`, { method: 'DELETE' })
      showToast('🗑️ ویژگی با موفقیت حذف شد!')
      clearDraft()
      navigateTo('/admin/product-management/attributes')
    } catch (e) {
      showToast('خطا در حذف ویژگی', 'error')
    }
  }
}

const toggleSection = (section) => {
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
    alert('لطفاً نام گزینه را وارد کنید')
    return
  }

  // جلوگیری از مقدار تکراری (کلاینت-ساید)
  const normName = normalizeValue(name)
  const duplicate = options.value.some((o, idx) => idx !== editingIndex.value && normalizeValue(o.name) === normName)
  if (duplicate) {
    alert('این مقدار قبلاً ثبت شده است')
    return
  }

  const optionObj = {
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

const deleteOption = (optionName) => {
  if (isColorAttr.value) {
    alert('حذف گزینه برای ویژگی رنگ مجاز نیست')
    return
  }
  if (confirm(`آیا از حذف گزینه "${optionName}" اطمینان دارید؟`)) {
    options.value = options.value.filter(o => o.name !== optionName)
  }
}

const editOption = (optionName) => {
  // ویرایش برای ویژگی رنگ مجاز است؛ فقط حذف محدود شده است.
}

const viewProduct = (productName) => {
  alert(`مشاهده محصول "${productName}" (نمونه)`) // TODO: implement
}

// Auto-generate attribute code when name changes (unless user manually edited)
watch(() => formData.value.name, (newName) => {
  if (!newName) return
  // If user hasn't typed a code manually or current code matches previous slug, update it
  const auto = toSlug(newName, '_')
  if (!formData.value.code || formData.value.code === auto || formData.value.code.startsWith('attr_')) {
    formData.value.code = auto
  }
})

// برچسب نمایشی نوع داده
const dataTypeLabel = computed(() => {
  const map = { text: 'متن', string: 'متن چندخطی', number: 'عدد', color: 'رنگ' }
  return map[formData.value.dataType] || 'نامشخص'
})

// in <script> section additions
const DEFAULT_UNITS = [
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

const unitOptions = ref([])

const loadUnits = () => {
  try {
    const saved = localStorage.getItem('measurementUnits')
    unitOptions.value = saved ? JSON.parse(saved) : DEFAULT_UNITS
  } catch (_) {
    unitOptions.value = DEFAULT_UNITS
  }
}

const saveUnits = () => {
  localStorage.setItem('measurementUnits', JSON.stringify(unitOptions.value))
}

const refreshUnits = () => {
  loadUnits()
  // اگر واحد انتخاب‌شده حذف شده بود، ریست شود
  if (unitSelection.value && !unitOptions.value.some(u=>u.value===unitSelection.value)) {
    unitSelection.value = ''
    formData.value.unit = ''
  }
}

onMounted(async () => {
  loadUnits()
  if (attributeId.value) {
    // Load main attribute data
    try {
      const attribute = await $fetch(`/api/attributes/${attributeId.value}`)
      if (attribute) {
        formData.value.name = attribute.name || ''
        formData.value.displayText = attribute.display_text || ''
        formData.value.code = attribute.code || ''
        formData.value.dataType = attribute.data_type || 'auto'
        formData.value.unit = attribute.unit || ''
        unitSelection.value = attribute.unit || ''
        
        // Load attribute values
        await fetchAttributeValues(attributeId.value)
      }
    } catch (err) {
      console.error('Failed to load attribute data:', err)
    }
    
    loadAttributeGroups()
  }
})

const showUnitModal = ref(false)

// واکنش به انتخاب «مدیریت واحد»
const unitSelection = ref(formData.value.unit || '')
const customUnit = ref('')

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

watch(customUnit, (val) => {
  if (unitSelection.value === '_custom') {
    formData.value.unit = val
  }
})

// ------------- Helper: Load attribute values from backend -------------
const fetchAttributeValues = async (attrId) => {
  try {
    const vals = await $fetch(`/api/attribute-values/by-attribute/${attrId}`)
    if (Array.isArray(vals)) {
      options.value = vals.map(v => {
        let metaObj = {}
        if (v.meta) {
          try { metaObj = typeof v.meta === 'string' ? JSON.parse(v.meta) : v.meta } catch (_) { metaObj = {} }
        }
        return {
          id: Number(v.id),
          name: v.value,
          hasColor: !!metaObj.color,
          colorValue: metaObj.color || '#000000',
          colorName: metaObj.color_name || ''
        }
      })
      originalOptionIds.value = new Set(vals.map(v => Number(v.id)))
    }
  } catch (e) {
    console.error('Failed to fetch attribute values', e)
  }
}

// --- Groups this attribute belongs to ---
const groups = ref([])

const loadAttributeGroups = async () => {
  if (!attributeId || !attributeId.value) return
  try {
    const res = await $fetch(`/api/attribute-groups/by-attribute/${attributeId.value}`)
    groups.value = Array.isArray(res) ? res : (res.data || [])
  } catch (err) {
    console.error('Failed to fetch attribute groups', err)
  }
}

// تابع دریافت مجدد اطلاعات ویژگی از سرور
const fetchAttribute = async () => {
  if (!attributeId.value) return;
  try {
    const attribute = await $fetch(`/api/attributes/${attributeId.value}`);
    console.log('attribute from API', attribute);
    if (attribute) {
      formData.value.name = attribute.name || '';
      formData.value.displayText = attribute.display_name || '';
      console.log('formData.displayText', formData.value.displayText);
      formData.value.dataType = attribute.data_type || 'auto';
      formData.value.unit = attribute.unit || '';
      // سایر فیلدها اگر لازم بود
    }
  } catch (e) {
    console.error('خطا در دریافت اطلاعات ویژگی', e);
  }
}

// مقداردهی خودکار displayText از name، مگر این که کاربر خودش displayText را تغییر دهد
let displayTextManuallyChanged = false

watch(() => formData.value.displayText, (val, oldVal) => {
  if (oldVal === '' && val !== '' && val !== formData.value.name) {
    displayTextManuallyChanged = true
  }
})

watch(() => formData.value.name, (newName, oldName) => {
  if (!displayTextManuallyChanged) {
    formData.value.displayText = newName || ''
  }
})

</script>