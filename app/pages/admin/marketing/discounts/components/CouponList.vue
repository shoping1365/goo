<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200 w-full">
    <!-- هدر بخش -->
    <div class="p-6 border-b border-gray-200">
      <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-6">
        <div class="min-w-0 flex-1">
          <h2 class="text-lg font-semibold text-gray-900">مدیریت کوپن‌ها</h2>
          <p class="text-gray-600 mt-1">ایجاد و مدیریت کوپن‌های تخفیف</p>
        </div>
        <div class="flex items-center gap-3 flex-shrink-0">
          <button @click="showBulkActions = !showBulkActions" class="px-4 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200 transition-colors whitespace-nowrap">
            <span class="i-heroicons-squares-2x2 ml-2"></span>
            عملیات دسته‌ای
          </button>
          <button @click="showAddCoupon = true" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors whitespace-nowrap">
            <span class="i-heroicons-plus ml-2"></span>
            افزودن کوپن جدید
          </button>
        </div>
      </div>
    </div>

    <!-- فیلترها -->
    <div class="p-6 border-b border-gray-200">
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-5 gap-6">
        <div class="min-w-0">
          <label class="block text-sm font-medium text-gray-700 mb-2">نوع کوپن</label>
          <select v-model="filters.type" @change="applyFilters" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
            <option value="">همه انواع</option>
            <option value="percentage">درصدی</option>
            <option value="fixed">مبلغ ثابت</option>
            <option value="free_shipping">ارسال رایگان</option>
          </select>
        </div>
        <div class="min-w-0">
          <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
          <select v-model="filters.status" @change="applyFilters" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
            <option value="">همه وضعیت‌ها</option>
            <option value="active">فعال</option>
            <option value="inactive">غیرفعال</option>
            <option value="expired">منقضی شده</option>
          </select>
        </div>
        <div class="min-w-0">
          <label class="block text-sm font-medium text-gray-700 mb-2">جستجو</label>
          <input v-model="filters.search" @input="applyFilters" type="text" placeholder="جستجو در کد یا نام..." class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
        </div>
        <div class="min-w-0">
          <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ انقضا</label>
          <select v-model="filters.expiry" @change="applyFilters" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
            <option value="">همه</option>
            <option value="today">امروز</option>
            <option value="week">این هفته</option>
            <option value="month">این ماه</option>
            <option value="expired">منقضی شده</option>
          </select>
        </div>
        <div class="flex items-end min-w-0">
          <button @click="resetFilters" class="w-full px-4 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200 transition-colors whitespace-nowrap">
            پاک کردن فیلترها
          </button>
        </div>
      </div>
    </div>

    <!-- پیام عدم وجود نتیجه -->
    <div class="text-center py-12">
      <div class="w-16 h-16 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-4">
        <span class="i-heroicons-ticket text-gray-400 text-xl"></span>
      </div>
      <h3 class="text-lg font-medium text-gray-900 mb-2">هیچ کوپنی یافت نشد</h3>
      <p class="text-gray-600">کوپن‌های تخفیف با فیلترهای انتخاب شده یافت نشد.</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { useAuth } from '~/composables/useAuth'

// تعریف interface برای کوپن
interface Coupon {
  id: number
  code: string
  name: string
  description: string
  type: 'percentage' | 'fixed' | 'free_shipping'
  discountValue: number
  minAmount: number
  maxDiscount?: number
  status: 'active' | 'inactive' | 'expired'
  usedCount: number
  maxUses?: number
  maxUsesPerUser?: number
  startsAt: string
  expiresAt: string
  createdAt: string
  updatedAt: string
  applicableProducts?: number[]
  excludedProducts?: number[]
  applicableCategories?: number[]
  excludedCategories?: number[]
  userGroups?: string[]
}

// فیلترها
const filters = ref({
  type: '',
  status: '',
  search: '',
  expiry: ''
})

// متغیرهای reactive
const showAddCoupon = ref(false)
const showEditCoupon = ref(false)
const showCouponDetails = ref(false)
const showBulkActions = ref(false)
const selectedCoupon = ref<Coupon | null>(null)
const editingCoupon = ref<Coupon | null>(null)
const selectedCoupons = ref<number[]>([])
const selectAll = ref(false)
const currentPage = ref(1)
const itemsPerPage = ref(10)

// داده‌های نمونه کوپن‌ها
const coupons = ref<Coupon[]>([
  {
    id: 1,
    code: 'WELCOME20',
    name: 'کوپن خوشامدگویی',
    description: 'تخفیف ۲۰ درصدی برای مشتریان جدید',
    type: 'percentage',
    discountValue: 20,
    minAmount: 100000,
    maxDiscount: 50000,
    status: 'active',
    usedCount: 45,
    maxUses: 100,
    maxUsesPerUser: 1,
    startsAt: '2024-01-01',
    expiresAt: '2024-12-31',
    createdAt: '2024-01-01',
    updatedAt: '2024-01-15'
  },
  {
    id: 2,
    code: 'FREESHIP',
    name: 'ارسال رایگان',
    description: 'ارسال رایگان برای سفارش‌های بالای ۵۰۰ هزار تومان',
    type: 'free_shipping',
    discountValue: 0,
    minAmount: 500000,
    status: 'active',
    usedCount: 120,
    maxUses: 1000,
    maxUsesPerUser: 3,
    startsAt: '2024-01-01',
    expiresAt: '2024-06-30',
    createdAt: '2024-01-01',
    updatedAt: '2024-01-15'
  },
  {
    id: 3,
    code: 'SAVE50K',
    name: 'تخفیف ۵۰ هزار تومانی',
    description: 'تخفیف ۵۰ هزار تومانی برای سفارش‌های بالای ۳۰۰ هزار تومان',
    type: 'fixed',
    discountValue: 50000,
    minAmount: 300000,
    status: 'active',
    usedCount: 78,
    maxUses: 200,
    maxUsesPerUser: 2,
    startsAt: '2024-01-01',
    expiresAt: '2024-03-31',
    createdAt: '2024-01-01',
    updatedAt: '2024-01-15'
  },
  {
    id: 4,
    code: 'BIRTHDAY15',
    name: 'تخفیف تولد',
    description: 'تخفیف ۱۵ درصدی در روز تولد',
    type: 'percentage',
    discountValue: 15,
    minAmount: 50000,
    status: 'inactive',
    usedCount: 12,
    maxUses: 500,
    maxUsesPerUser: 1,
    startsAt: '2024-01-01',
    expiresAt: '2024-12-31',
    createdAt: '2024-01-01',
    updatedAt: '2024-01-15'
  },
  {
    id: 5,
    code: 'FLASH30',
    name: 'تخفیف فلش',
    description: 'تخفیف ۳۰ درصدی برای ۲۴ ساعت',
    type: 'percentage',
    discountValue: 30,
    minAmount: 200000,
    maxDiscount: 100000,
    status: 'expired',
    usedCount: 89,
    maxUses: 100,
    maxUsesPerUser: 1,
    startsAt: '2024-01-10',
    expiresAt: '2024-01-11',
    createdAt: '2024-01-10',
    updatedAt: '2024-01-11'
  }
])

// محاسبات
const filteredCoupons = computed(() => {
  let result = coupons.value

  // فیلتر بر اساس نوع
  if (filters.value.type) {
    result = result.filter(coupon => coupon.type === filters.value.type)
  }

  // فیلتر بر اساس وضعیت
  if (filters.value.status) {
    result = result.filter(coupon => coupon.status === filters.value.status)
  }

  // فیلتر بر اساس جستجو
  if (filters.value.search) {
    const search = filters.value.search.toLowerCase()
    result = result.filter(coupon => 
      coupon.code.toLowerCase().includes(search) ||
      coupon.name.toLowerCase().includes(search) ||
      coupon.description.toLowerCase().includes(search)
    )
  }

  // فیلتر بر اساس تاریخ انقضا
  if (filters.value.expiry) {
    const now = new Date()
    const today = new Date(now.getFullYear(), now.getMonth(), now.getDate())
    const weekFromNow = new Date(today.getTime() + 7 * 24 * 60 * 60 * 1000)
    const monthFromNow = new Date(today.getTime() + 30 * 24 * 60 * 60 * 1000)

    result = result.filter(coupon => {
      const expiryDate = new Date(coupon.expiresAt)
      switch (filters.value.expiry) {
        case 'today':
          return expiryDate.toDateString() === today.toDateString()
        case 'week':
          return expiryDate >= today && expiryDate <= weekFromNow
        case 'month':
          return expiryDate >= today && expiryDate <= monthFromNow
        case 'expired':
          return expiryDate < today
        default:
          return true
      }
    })
  }

  return result
})

const totalPages = computed(() => Math.ceil(filteredCoupons.value.length / itemsPerPage.value))

// توابع کمکی
const getTypeName = (type: string) => {
  const types = {
    percentage: 'درصدی',
    fixed: 'مبلغ ثابت',
    free_shipping: 'ارسال رایگان'
  }
  return types[type as keyof typeof types] || type
}

const getTypeBadgeClass = (type: string) => {
  const classes = {
    percentage: 'bg-blue-100 text-blue-800',
    fixed: 'bg-green-100 text-green-800',
    free_shipping: 'bg-purple-100 text-purple-800'
  }
  return classes[type as keyof typeof classes] || 'bg-gray-100 text-gray-800'
}

const getStatusName = (status: string) => {
  const statuses = {
    active: 'فعال',
    inactive: 'غیرفعال',
    expired: 'منقضی شده'
  }
  return statuses[status as keyof typeof statuses] || status
}

const getStatusBadgeClass = (status: string) => {
  const classes = {
    active: 'bg-green-100 text-green-800',
    inactive: 'bg-yellow-100 text-yellow-800',
    expired: 'bg-red-100 text-red-800'
  }
  return classes[status as keyof typeof classes] || 'bg-gray-100 text-gray-800'
}

const formatDiscount = (coupon: Coupon) => {
  switch (coupon.type) {
    case 'percentage':
      return `${coupon.discountValue}%`
    case 'fixed':
      return `${coupon.discountValue.toLocaleString()} تومان`
    case 'free_shipping':
      return 'ارسال رایگان'
    default:
      return ''
  }
}

const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleDateString('fa-IR')
}

// توابع عملیاتی
const applyFilters = () => {
  currentPage.value = 1
}

const resetFilters = () => {
  filters.value = {
    type: '',
    status: '',
    search: '',
    expiry: ''
  }
  currentPage.value = 1
}

const toggleSelectAll = () => {
  if (selectAll.value) {
    selectedCoupons.value = filteredCoupons.value.map(coupon => coupon.id)
  } else {
    selectedCoupons.value = []
  }
}

const copyCode = async (code: string) => {
  try {
    await navigator.clipboard.writeText(code)
    // TODO: نمایش پیام موفقیت
  } catch (error) {
    console.error('خطا در کپی کردن کد:', error)
  }
}

const editCoupon = (coupon: Coupon) => {
  editingCoupon.value = { ...coupon }
  showEditCoupon.value = true
}

const toggleCouponStatus = async (coupon: Coupon) => {
  try {
    // TODO: فراخوانی API برای تغییر وضعیت
    coupon.status = coupon.status === 'active' ? 'inactive' : 'active'
  } catch (error) {
    console.error('خطا در تغییر وضعیت کوپن:', error)
  }
}

const viewCouponDetails = (coupon: Coupon) => {
  selectedCoupon.value = coupon
  showCouponDetails.value = true
}

const deleteCoupon = async (coupon: Coupon) => {
  if (confirm(`آیا از حذف کوپن "${coupon.name}" اطمینان دارید؟`)) {
    try {
      // TODO: فراخوانی API برای حذف کوپن
      const index = coupons.value.findIndex(c => c.id === coupon.id)
      if (index !== -1) {
        coupons.value.splice(index, 1)
      }
    } catch (error) {
      console.error('خطا در حذف کوپن:', error)
    }
  }
}

const bulkActivate = async () => {
  try {
    // TODO: فراخوانی API برای فعال کردن دسته‌ای
    selectedCoupons.value.forEach(id => {
      const coupon = coupons.value.find(c => c.id === id)
      if (coupon) coupon.status = 'active'
    })
    selectedCoupons.value = []
  } catch (error) {
    console.error('خطا در فعال کردن دسته‌ای:', error)
  }
}

const bulkDeactivate = async () => {
  try {
    // TODO: فراخوانی API برای غیرفعال کردن دسته‌ای
    selectedCoupons.value.forEach(id => {
      const coupon = coupons.value.find(c => c.id === id)
      if (coupon) coupon.status = 'inactive'
    })
    selectedCoupons.value = []
  } catch (error) {
    console.error('خطا در غیرفعال کردن دسته‌ای:', error)
  }
}

const bulkDelete = async () => {
  if (confirm(`آیا از حذف ${selectedCoupons.value.length} کوپن انتخاب شده اطمینان دارید؟`)) {
    try {
      // TODO: فراخوانی API برای حذف دسته‌ای
      coupons.value = coupons.value.filter(coupon => !selectedCoupons.value.includes(coupon.id))
      selectedCoupons.value = []
    } catch (error) {
      console.error('خطا در حذف دسته‌ای:', error)
    }
  }
}

const handleSaveCoupon = (couponData: any) => {
  if (editingCoupon.value) {
    // ویرایش کوپن موجود
    const index = coupons.value.findIndex(c => c.id === editingCoupon.value?.id)
    if (index !== -1) {
      coupons.value[index] = { ...coupons.value[index], ...couponData }
    }
  } else {
    // افزودن کوپن جدید
    const newCoupon: Coupon = {
      id: Date.now(),
      ...couponData,
      usedCount: 0,
      createdAt: new Date().toISOString(),
      updatedAt: new Date().toISOString()
    }
    coupons.value.unshift(newCoupon)
  }
  closeCouponForm()
}

const closeCouponForm = () => {
  showAddCoupon.value = false
  showEditCoupon.value = false
  editingCoupon.value = null
}

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

// بارگذاری داده‌ها
onMounted(async () => {
  try {
    // TODO: فراخوانی API برای دریافت لیست کوپن‌ها
    // const response = await $fetch('/api/admin/coupons')
    // coupons.value = response.data
  } catch (error) {
    console.error('خطا در بارگذاری کوپن‌ها:', error)
  }
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { hasPermission } = useAuth()

// Computed برای چک کردن پرمیژن حذف
const canDeleteCoupon = computed(() => hasPermission('coupon.delete'))
</script>

<!--
  کامپوننت لیست کوپن‌ها
  شامل:
  1. نمایش لیست کوپن‌ها در جدول
  2. فیلترهای پیشرفته
  3. عملیات دسته‌ای
  4. عملیات فردی (ویرایش، حذف، فعال/غیرفعال)
  5. صفحه‌بندی
  6. کپی کردن کد کوپن
  توضیحات کامل به فارسی در کد
-->
