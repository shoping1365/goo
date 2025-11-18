<template>
  <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
    <!-- هدر -->
    <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-6 mb-6">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">مدیریت نرخ‌های مالیاتی</h3>
        <p class="text-sm text-gray-600">تعریف و مدیریت نرخ‌های مالیات بر ارزش افزوده و سایر مالیات‌ها</p>
      </div>
      
      <!-- دکمه افزودن نرخ جدید -->
      <button 
        @click="showAddRateModal = true"
        class="inline-flex items-center gap-2 px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg transition-colors duration-200"
      >
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
        </svg>
        افزودن نرخ جدید
      </button>
    </div>

    <!-- فیلترها -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
      <!-- فیلتر نوع مالیات -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">نوع مالیات</label>
        <select 
          v-model="filters.taxType"
          class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
        >
          <option value="all">همه انواع</option>
          <option value="vat">مالیات بر ارزش افزوده</option>
          <option value="income">مالیات بر درآمد</option>
          <option value="customs">مالیات گمرکی</option>
          <option value="excise">مالیات غیرمستقیم</option>
        </select>
      </div>

      <!-- فیلتر وضعیت -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
        <select 
          v-model="filters.status"
          class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
        >
          <option value="all">همه وضعیت‌ها</option>
          <option value="active">فعال</option>
          <option value="inactive">غیرفعال</option>
          <option value="expired">منقضی شده</option>
        </select>
      </div>

      <!-- فیلتر دسته‌بندی -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">دسته‌بندی</label>
        <select 
          v-model="filters.category"
          class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
        >
          <option value="all">همه دسته‌ها</option>
          <option value="products">محصولات</option>
          <option value="services">خدمات</option>
          <option value="digital">محصولات دیجیتال</option>
          <option value="food">مواد غذایی</option>
        </select>
      </div>

      <!-- جستجو -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">جستجو</label>
        <input 
          v-model="filters.search"
          type="text"
          placeholder="جستجو در نرخ‌ها..."
          class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
        />
      </div>
    </div>

    <!-- جدول نرخ‌های مالیاتی -->
    <div class="overflow-x-auto">
      <table class="w-full text-sm">
        <thead>
          <tr class="border-b border-gray-200 bg-gray-50">
            <th class="text-right py-3 px-4 font-medium text-gray-600">نوع مالیات</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">دسته‌بندی</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">نرخ (%)</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">تاریخ شروع</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">تاریخ پایان</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">وضعیت</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">عملیات</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="rate in filteredRates" :key="rate.id" class="border-b border-gray-100 hover:bg-gray-50">
            <td class="py-3 px-4">
              <div class="flex items-center gap-2">
                <div class="w-8 h-8 rounded-lg flex items-center justify-center" :class="getTaxTypeColor(rate?.taxType || '')">
                  <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
                  </svg>
                </div>
                <span class="font-medium text-gray-900">{{ getTaxTypeLabel(rate?.taxType || '') }}</span>
              </div>
            </td>
            <td class="py-3 px-4 text-gray-900">{{ rate.category }}</td>
            <td class="py-3 px-4">
              <span class="font-bold text-blue-600">{{ rate.rate }}%</span>
            </td>
            <td class="py-3 px-4 text-gray-600">{{ formatDate(rate.startDate) }}</td>
            <td class="py-3 px-4 text-gray-600">{{ formatDate(rate.endDate) }}</td>
            <td class="py-3 px-4">
              <span 
                :class="[
                  'px-2 py-1 rounded-full text-xs font-medium',
                  rate.status === 'active' ? 'bg-green-100 text-green-700' : 
                  rate.status === 'inactive' ? 'bg-red-100 text-red-700' : 
                  'bg-yellow-100 text-yellow-700'
                ]"
              >
                {{ getStatusLabel(rate.status) }}
              </span>
            </td>
            <td class="py-3 px-4">
              <div class="flex items-center gap-2">
                <button 
                  @click="editRate(rate)"
                  class="p-1 text-blue-600 hover:text-blue-800 transition-colors"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
                  </svg>
                </button>
                <button 
                  @click="toggleRateStatus(rate)"
                  :class="[
                    'p-1 transition-colors',
                    rate.status === 'active' ? 'text-red-600 hover:text-red-800' : 'text-green-600 hover:text-green-800'
                  ]"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="rate.status === 'active' ? 'M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728L5.636 5.636m12.728 12.728L18.364 5.636M5.636 18.364l12.728-12.728' : 'M5 13l4 4L19 7'"></path>
                  </svg>
                </button>
                <button 
                  v-if="canDeleteTaxRate"
                  @click="deleteRate(rate)"
                  class="p-1 text-red-600 hover:text-red-800 transition-colors"
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

    <!-- آمار نرخ‌ها -->
    <div class="mt-6 pt-6 border-t border-gray-200">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div class="text-center">
          <div class="text-2xl font-bold text-gray-900">{{ totalRates }}</div>
          <div class="text-sm text-gray-500">کل نرخ‌ها</div>
        </div>
        <div class="text-center">
          <div class="text-2xl font-bold text-green-600">{{ activeRates }}</div>
          <div class="text-sm text-gray-500">نرخ‌های فعال</div>
        </div>
        <div class="text-center">
          <div class="text-2xl font-bold text-red-600">{{ inactiveRates }}</div>
          <div class="text-sm text-gray-500">نرخ‌های غیرفعال</div>
        </div>
        <div class="text-center">
          <div class="text-2xl font-bold text-yellow-600">{{ expiredRates }}</div>
          <div class="text-sm text-gray-500">نرخ‌های منقضی شده</div>
        </div>
      </div>
    </div>

    <!-- مودال افزودن/ویرایش نرخ -->
    <div v-if="showAddRateModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-xl p-6 w-full max-w-md mx-4">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">
            {{ editingRate ? 'ویرایش نرخ مالیاتی' : 'افزودن نرخ جدید' }}
          </h3>
          <button @click="closeModal" class="text-gray-400 hover:text-gray-600">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>

        <form @submit.prevent="saveRate" class="space-y-4">
          <!-- نوع مالیات -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع مالیات</label>
            <select 
              v-model="rateForm.taxType"
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="">انتخاب کنید</option>
              <option value="vat">مالیات بر ارزش افزوده</option>
              <option value="income">مالیات بر درآمد</option>
              <option value="customs">مالیات گمرکی</option>
              <option value="excise">مالیات غیرمستقیم</option>
            </select>
          </div>

          <!-- دسته‌بندی -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">دسته‌بندی</label>
            <select 
              v-model="rateForm.category"
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="">انتخاب کنید</option>
              <option value="products">محصولات</option>
              <option value="services">خدمات</option>
              <option value="digital">محصولات دیجیتال</option>
              <option value="food">مواد غذایی</option>
            </select>
          </div>

          <!-- نرخ -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نرخ مالیات (%)</label>
            <input 
              v-model.number="rateForm.rate"
              type="number"
              step="0.1"
              min="0"
              max="100"
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
          </div>

          <!-- تاریخ شروع -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ شروع</label>
            <input 
              v-model="rateForm.startDate"
              type="date"
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
          </div>

          <!-- تاریخ پایان -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ پایان</label>
            <input 
              v-model="rateForm.endDate"
              type="date"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
          </div>

          <!-- توضیحات -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات</label>
            <textarea 
              v-model="rateForm.description"
              rows="3"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            ></textarea>
          </div>

          <!-- دکمه‌ها -->
          <div class="flex gap-3 pt-4">
            <button 
              type="submit"
              class="flex-1 px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg transition-colors duration-200"
            >
              {{ editingRate ? 'ویرایش' : 'افزودن' }}
            </button>
            <button 
              type="button"
              @click="closeModal"
              class="flex-1 px-4 py-2 bg-gray-200 hover:bg-gray-300 text-gray-700 rounded-lg transition-colors duration-200"
            >
              انصراف
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
</script>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import { useAuth } from '~/composables/useAuth';

// احراز هویت غیرفعال شده است - محدودیت دسترسی حذف شده
definePageMeta({
  // احراز هویت غیرفعال شده است
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { hasPermission } = useAuth()

// Computed برای چک کردن پرمیژن حذف
const canDeleteTaxRate = computed(() => hasPermission('tax_rate.delete'))

// نرخ‌های مالیاتی
const taxRates = ref([
  {
    id: 1,
    taxType: 'vat',
    category: 'محصولات',
    rate: 9,
    startDate: '2024-01-01',
    endDate: '2024-12-31',
    status: 'active',
    description: 'مالیات بر ارزش افزوده برای محصولات عمومی'
  },
  {
    id: 2,
    taxType: 'vat',
    category: 'خدمات',
    rate: 9,
    startDate: '2024-01-01',
    endDate: '2024-12-31',
    status: 'active',
    description: 'مالیات بر ارزش افزوده برای خدمات'
  },
  {
    id: 3,
    taxType: 'income',
    category: 'محصولات دیجیتال',
    rate: 15,
    startDate: '2024-01-01',
    endDate: '2024-06-30',
    status: 'expired',
    description: 'مالیات بر درآمد برای محصولات دیجیتال'
  },
  {
    id: 4,
    taxType: 'customs',
    category: 'مواد غذایی',
    rate: 5,
    startDate: '2024-01-01',
    endDate: '2024-12-31',
    status: 'inactive',
    description: 'مالیات گمرکی برای مواد غذایی'
  }
])

// فیلترها
const filters = ref({
  taxType: 'all',
  status: 'all',
  category: 'all',
  search: ''
})

// مودال
const showAddRateModal = ref(false)
const editingRate = ref(null)

// فرم نرخ
const rateForm = ref({
  taxType: '',
  category: '',
  rate: 0,
  startDate: '',
  endDate: '',
  description: ''
})

// فیلتر کردن نرخ‌ها
const filteredRates = computed(() => {
  return taxRates.value.filter(rate => {
    if ((filters.value?.taxType ?? 'all') !== 'all' && (rate?.taxType ?? '') !== (filters.value?.taxType ?? '')) return false
    if (filters.value.status !== 'all' && rate.status !== filters.value.status) return false
    if (filters.value.category !== 'all' && rate.category !== filters.value.category) return false
    if (filters.value.search && !rate.description.includes(filters.value.search)) return false
    return true
  })
})

// آمار نرخ‌ها
const totalRates = computed(() => taxRates.value.length)
const activeRates = computed(() => taxRates.value.filter(rate => rate.status === 'active').length)
const inactiveRates = computed(() => taxRates.value.filter(rate => rate.status === 'inactive').length)
const expiredRates = computed(() => taxRates.value.filter(rate => rate.status === 'expired').length)

// رنگ نوع مالیات
const getTaxTypeColor = (taxType: string) => {
  const colors = {
    vat: 'bg-blue-500',
    income: 'bg-green-500',
    customs: 'bg-purple-500',
    excise: 'bg-orange-500'
  }
  return colors[taxType] || 'bg-gray-500'
}

// برچسب نوع مالیات
const getTaxTypeLabel = (taxType: string) => {
  const labels = {
    vat: 'مالیات بر ارزش افزوده',
    income: 'مالیات بر درآمد',
    customs: 'مالیات گمرکی',
    excise: 'مالیات غیرمستقیم'
  }
  return labels[taxType] || taxType
}

// برچسب وضعیت
const getStatusLabel = (status: string) => {
  const labels = {
    active: 'فعال',
    inactive: 'غیرفعال',
    expired: 'منقضی شده'
  }
  return labels[status] || status
}

// فرمت تاریخ
const formatDate = (date: string) => {
  return new Date(date).toLocaleDateString('fa-IR')
}

// ویرایش نرخ
const editRate = (rate: any) => {
  editingRate.value = rate
  rateForm.value = { ...rate }
  showAddRateModal.value = true
}

// تغییر وضعیت نرخ
const toggleRateStatus = async (rate: any) => {
  try {
    // TODO: ارسال درخواست به API
    // await $fetch(`/api/admin/tax/rates/${rate.id}/toggle-status`, { method: 'PATCH' })
    
    rate.status = rate.status === 'active' ? 'inactive' : 'active'
  } catch (error) {
    console.error('خطا در تغییر وضعیت نرخ:', error)
  }
}

// حذف نرخ
const deleteRate = async (rate: any) => {
  if (confirm('آیا از حذف این نرخ اطمینان دارید؟')) {
    try {
      // TODO: ارسال درخواست به API
      // await $fetch(`/api/admin/tax/rates/${rate.id}`, { method: 'DELETE' })
      
      const index = taxRates.value.findIndex(r => r.id === rate.id)
      if (index > -1) {
        taxRates.value.splice(index, 1)
      }
    } catch (error) {
      console.error('خطا در حذف نرخ:', error)
    }
  }
}

// ذخیره نرخ
const saveRate = async () => {
  try {
    if (editingRate.value) {
      // ویرایش نرخ موجود
      // await $fetch(`/api/admin/tax/rates/${editingRate.value.id}`, {
      //   method: 'PUT',
      //   body: rateForm.value
      // })
      
      const index = taxRates.value.findIndex(r => r.id === editingRate.value.id)
      if (index > -1) {
        taxRates.value[index] = { ...editingRate.value, ...rateForm.value }
      }
    } else {
      // افزودن نرخ جدید
      // const response = await $fetch('/api/admin/tax/rates', {
      //   method: 'POST',
      //   body: rateForm.value
      // })
      
      const newRate = {
        id: Date.now(),
        ...rateForm.value,
        status: 'active'
      }
      taxRates.value.push(newRate)
    }
    
    closeModal()
  } catch (error) {
    console.error('خطا در ذخیره نرخ:', error)
  }
}

// بستن مودال
const closeModal = () => {
  showAddRateModal.value = false
  editingRate.value = null
  rateForm.value = {
    taxType: '',
    category: '',
    rate: 0,
    startDate: '',
    endDate: '',
    description: ''
  }
}

// بارگذاری اولیه
onMounted(() => {
  // TODO: بارگذاری نرخ‌ها از API
})
</script>

<!--
  کامپوننت مدیریت نرخ‌های مالیاتی
  شامل:
  1. جدول نرخ‌های مالیاتی
  2. فیلترهای پیشرفته
  3. افزودن/ویرایش نرخ‌ها
  4. تغییر وضعیت و حذف
  5. آمار نرخ‌ها
  6. طراحی مدرن و کاملاً ریسپانسیو
--> 
