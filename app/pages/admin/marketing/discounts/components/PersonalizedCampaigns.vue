<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200">
    <!-- هدر بخش -->
    <div class="p-6 border-b border-gray-200">
      <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-6">
        <div class="min-w-0 flex-1">
          <h2 class="text-lg font-semibold text-gray-900">کمپین‌های شخصی‌سازی شده</h2>
          <p class="text-gray-600 mt-1">مدیریت کمپین‌های تخفیف شخصی‌سازی شده بر اساس رفتار کاربران</p>
        </div>
        <div class="flex items-center gap-3 flex-shrink-0">
          <button @click="showAddCampaign = true" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors whitespace-nowrap">
            <span class="i-heroicons-plus ml-2"></span>
            افزودن کمپین جدید
          </button>
        </div>
      </div>
    </div>

    <!-- فیلترها -->
    <div class="p-6 border-b border-gray-200">
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
        <div class="min-w-0">
          <label class="block text-sm font-medium text-gray-700 mb-2">نوع شخصی‌سازی</label>
          <select v-model="filters.personalizationType" @change="applyFilters" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
            <option value="">همه انواع</option>
            <option value="behavioral">رفتاری</option>
            <option value="demographic">جمعیت‌شناختی</option>
            <option value="purchase_history">تاریخچه خرید</option>
            <option value="location">مکانی</option>
            <option value="device">دستگاه</option>
          </select>
        </div>
        <div class="min-w-0">
          <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
          <select v-model="filters.status" @change="applyFilters" class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
            <option value="">همه وضعیت‌ها</option>
            <option value="active">فعال</option>
            <option value="paused">متوقف</option>
            <option value="draft">پیش‌نویس</option>
            <option value="ended">پایان یافته</option>
          </select>
        </div>
        <div class="min-w-0">
          <label class="block text-sm font-medium text-gray-700 mb-2">جستجو</label>
          <input v-model="filters.search" @input="applyFilters" type="text" placeholder="جستجو در نام یا توضیحات..." class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
        </div>
        <div class="flex items-end min-w-0">
          <button @click="resetFilters" class="w-full px-4 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200 transition-colors whitespace-nowrap">
            پاک کردن فیلترها
          </button>
        </div>
      </div>
    </div>

    <!-- لیست کمپین‌ها -->
    <div class="overflow-x-auto">
      <table class="w-full min-w-max divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              <input v-model="selectAll" @change="toggleSelectAll" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نام کمپین</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع شخصی‌سازی</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">معیار هدف</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کاربران هدف</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نرخ تبدیل</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-for="campaign in filteredCampaigns" :key="campaign.id" class="hover:bg-gray-50">
            <td class="px-6 py-4 whitespace-nowrap">
              <input v-model="selectedCampaigns" :value="campaign.id" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="flex items-center">
                <div class="">
                  <div class="h-10 w-10 rounded-full bg-gradient-to-r from-purple-400 to-pink-400 flex items-center justify-center">
                    <span class="text-white font-semibold text-sm">{{ campaign.name.charAt(0) }}</span>
                  </div>
                </div>
                <div class="mr-4">
                  <div class="text-sm font-medium text-gray-900">{{ campaign.name }}</div>
                  <div class="text-sm text-gray-500">{{ campaign.description }}</div>
                </div>
              </div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span :class="`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${getPersonalizationTypeBadgeClass(campaign.personalizationType)}`">
                {{ getPersonalizationTypeName(campaign.personalizationType) }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
              {{ campaign.targetCriteria }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span :class="`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${getStatusBadgeClass(campaign.status)}`">
                {{ getStatusName(campaign.status) }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
              {{ formatNumber(campaign.targetUsers) }} کاربر
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
              <div class="flex items-center">
                <span class="text-green-600 font-medium">{{ campaign.conversionRate }}%</span>
                <span class="text-xs text-gray-500 mr-2">({{ campaign.conversions }}/{{ campaign.impressions }})</span>
              </div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
              <div class="flex items-center gap-2">
                <button @click="viewCampaign(campaign)" class="text-blue-600 hover:text-blue-900">
                  <span class="i-heroicons-eye"></span>
                </button>
                <button @click="editCampaign(campaign)" class="text-indigo-600 hover:text-indigo-900">
                  <span class="i-heroicons-pencil-square"></span>
                </button>
                <button @click="duplicateCampaign(campaign)" class="text-green-600 hover:text-green-900">
                  <span class="i-heroicons-document-duplicate"></span>
                </button>
                                  <button 
                    v-if="canDeleteCampaign"
                    @click="deleteCampaign(campaign)" 
                    class="text-red-600 hover:text-red-900"
                  >
                  <span class="i-heroicons-trash"></span>
                </button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- پیام عدم وجود نتیجه -->
    <div v-if="filteredCampaigns.length === 0" class="text-center py-12">
      <div class="w-16 h-16 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-4">
        <span class="i-heroicons-user-group text-gray-400 text-xl"></span>
      </div>
      <h3 class="text-lg font-medium text-gray-900 mb-2">هیچ کمپین شخصی‌سازی شده‌ای یافت نشد</h3>
      <p class="text-gray-600">کمپین‌های شخصی‌سازی شده با فیلترهای انتخاب شده یافت نشد.</p>
    </div>

    <!-- عملیات دسته‌ای -->
    <div v-if="selectedCampaigns.length > 0" class="p-6 border-t border-gray-200 bg-gray-50">
      <div class="flex items-center justify-between">
        <span class="text-sm text-gray-700">{{ selectedCampaigns.length }} کمپین انتخاب شده</span>
        <div class="flex items-center gap-2">
          <button @click="bulkActivate" class="px-3 py-1 bg-green-600 text-white text-sm rounded hover:bg-green-700 transition-colors">
            فعال‌سازی
          </button>
          <button @click="bulkPause" class="px-3 py-1 bg-yellow-600 text-white text-sm rounded hover:bg-yellow-700 transition-colors">
            توقف
          </button>
          <button @click="bulkDelete" class="px-3 py-1 bg-red-600 text-white text-sm rounded hover:bg-red-700 transition-colors">
            حذف
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { useAuth } from '~/composables/useAuth'

// تعریف interface برای کمپین شخصی‌سازی شده
interface PersonalizedCampaign {
  id: number
  name: string
  description: string
  personalizationType: 'behavioral' | 'demographic' | 'purchase_history' | 'location' | 'device'
  targetCriteria: string
  status: 'active' | 'paused' | 'draft' | 'ended'
  targetUsers: number
  conversionRate: number
  conversions: number
  impressions: number
  createdAt: string
  updatedAt: string
}

// فیلترها
const filters = ref({
  personalizationType: '',
  status: '',
  search: ''
})

// متغیرهای reactive
const campaigns = ref<PersonalizedCampaign[]>([])
const selectedCampaigns = ref<number[]>([])
const selectAll = ref(false)
const showAddCampaign = ref(false)

// داده‌های نمونه کمپین‌های شخصی‌سازی شده
const sampleCampaigns: PersonalizedCampaign[] = [
  {
    id: 1,
    name: 'تخفیف برای خریداران مکرر',
    description: 'تخفیف ویژه برای کاربرانی که بیش از 3 بار خرید کرده‌اند',
    personalizationType: 'purchase_history',
    targetCriteria: 'بیش از 3 خرید',
    status: 'active',
    targetUsers: 1250,
    conversionRate: 8.5,
    conversions: 106,
    impressions: 1245,
    createdAt: '2024-01-01',
    updatedAt: '2024-01-15'
  },
  {
    id: 2,
    name: 'تخفیف برای کاربران تهران',
    description: 'تخفیف ویژه برای کاربران ساکن تهران',
    personalizationType: 'location',
    targetCriteria: 'شهر تهران',
    status: 'active',
    targetUsers: 3200,
    conversionRate: 6.2,
    conversions: 198,
    impressions: 3195,
    createdAt: '2024-01-01',
    updatedAt: '2024-01-15'
  },
  {
    id: 3,
    name: 'تخفیف برای کاربران موبایل',
    description: 'تخفیف ویژه برای کاربرانی که از موبایل استفاده می‌کنند',
    personalizationType: 'device',
    targetCriteria: 'دستگاه موبایل',
    status: 'paused',
    targetUsers: 8500,
    conversionRate: 4.8,
    conversions: 408,
    impressions: 8492,
    createdAt: '2024-01-01',
    updatedAt: '2024-01-15'
  },
  {
    id: 4,
    name: 'تخفیف برای کاربران 25-35 سال',
    description: 'تخفیف ویژه برای کاربران در رده سنی 25 تا 35 سال',
    personalizationType: 'demographic',
    targetCriteria: 'سن 25-35 سال',
    status: 'active',
    targetUsers: 2100,
    conversionRate: 7.1,
    conversions: 149,
    impressions: 2095,
    createdAt: '2024-01-01',
    updatedAt: '2024-01-15'
  },
  {
    id: 5,
    name: 'تخفیف برای کاربران علاقه‌مند به الکترونیک',
    description: 'تخفیف ویژه برای کاربرانی که محصولات الکترونیک را مشاهده کرده‌اند',
    personalizationType: 'behavioral',
    targetCriteria: 'مشاهده محصولات الکترونیک',
    status: 'draft',
    targetUsers: 1800,
    conversionRate: 0,
    conversions: 0,
    impressions: 0,
    createdAt: '2024-01-01',
    updatedAt: '2024-01-15'
  }
]

// فیلتر کردن کمپین‌ها
const filteredCampaigns = computed(() => {
  return campaigns.value.filter(campaign => {
    const matchesType = !filters.value.personalizationType || campaign.personalizationType === filters.value.personalizationType
    const matchesStatus = !filters.value.status || campaign.status === filters.value.status
    const matchesSearch = !filters.value.search || 
      campaign.name.toLowerCase().includes(filters.value.search.toLowerCase()) ||
      campaign.description.toLowerCase().includes(filters.value.search.toLowerCase())
    
    return matchesType && matchesStatus && matchesSearch
  })
})

// اعمال فیلترها
function applyFilters() {
  // فیلترها به صورت reactive اعمال می‌شوند
}

// پاک کردن فیلترها
function resetFilters() {
  filters.value = {
    personalizationType: '',
    status: '',
    search: ''
  }
}

// انتخاب همه
function toggleSelectAll() {
  if (selectAll.value) {
    selectedCampaigns.value = filteredCampaigns.value.map(campaign => campaign.id)
  } else {
    selectedCampaigns.value = []
  }
}

// مشاهده کمپین
function viewCampaign(campaign: PersonalizedCampaign) {
  console.log('مشاهده کمپین:', campaign)
  // TODO: نمایش جزئیات کمپین
}

// ویرایش کمپین
function editCampaign(campaign: PersonalizedCampaign) {
  console.log('ویرایش کمپین:', campaign)
  // TODO: باز کردن فرم ویرایش
}

// کپی کردن کمپین
function duplicateCampaign(campaign: PersonalizedCampaign) {
  console.log('کپی کردن کمپین:', campaign)
  // TODO: ایجاد کپی از کمپین
}

// حذف کمپین
function deleteCampaign(campaign: PersonalizedCampaign) {
  if (confirm(`آیا از حذف کمپین "${campaign.name}" اطمینان دارید؟`)) {
    campaigns.value = campaigns.value.filter(c => c.id !== campaign.id)
    console.log('کمپین حذف شد:', campaign)
  }
}

// عملیات دسته‌ای - فعال‌سازی
function bulkActivate() {
  if (confirm(`آیا از فعال‌سازی ${selectedCampaigns.value.length} کمپین اطمینان دارید؟`)) {
    campaigns.value.forEach(campaign => {
      if (selectedCampaigns.value.includes(campaign.id)) {
        campaign.status = 'active'
      }
    })
    selectedCampaigns.value = []
    console.log('کمپین‌ها فعال شدند')
  }
}

// عملیات دسته‌ای - توقف
function bulkPause() {
  if (confirm(`آیا از توقف ${selectedCampaigns.value.length} کمپین اطمینان دارید؟`)) {
    campaigns.value.forEach(campaign => {
      if (selectedCampaigns.value.includes(campaign.id)) {
        campaign.status = 'paused'
      }
    })
    selectedCampaigns.value = []
    console.log('کمپین‌ها متوقف شدند')
  }
}

// عملیات دسته‌ای - حذف
function bulkDelete() {
  if (confirm(`آیا از حذف ${selectedCampaigns.value.length} کمپین اطمینان دارید؟`)) {
    campaigns.value = campaigns.value.filter(campaign => !selectedCampaigns.value.includes(campaign.id))
    selectedCampaigns.value = []
    console.log('کمپین‌ها حذف شدند')
  }
}

// توابع کمکی
function getPersonalizationTypeBadgeClass(type: string) {
  switch (type) {
    case 'behavioral':
      return 'bg-purple-100 text-purple-800'
    case 'demographic':
      return 'bg-blue-100 text-blue-800'
    case 'purchase_history':
      return 'bg-green-100 text-green-800'
    case 'location':
      return 'bg-yellow-100 text-yellow-800'
    case 'device':
      return 'bg-indigo-100 text-indigo-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

function getPersonalizationTypeName(type: string) {
  switch (type) {
    case 'behavioral':
      return 'رفتاری'
    case 'demographic':
      return 'جمعیت‌شناختی'
    case 'purchase_history':
      return 'تاریخچه خرید'
    case 'location':
      return 'مکانی'
    case 'device':
      return 'دستگاه'
    default:
      return 'نامشخص'
  }
}

function getStatusBadgeClass(status: string) {
  switch (status) {
    case 'active':
      return 'bg-green-100 text-green-800'
    case 'paused':
      return 'bg-yellow-100 text-yellow-800'
    case 'draft':
      return 'bg-gray-100 text-gray-800'
    case 'ended':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

function getStatusName(status: string) {
  switch (status) {
    case 'active':
      return 'فعال'
    case 'paused':
      return 'متوقف'
    case 'draft':
      return 'پیش‌نویس'
    case 'ended':
      return 'پایان یافته'
    default:
      return 'نامشخص'
  }
}

function formatNumber(num: number) {
  return new Intl.NumberFormat('fa-IR').format(num)
}

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { hasPermission } = useAuth()

// Computed برای چک کردن پرمیژن حذف
const canDeleteCampaign = computed(() => hasPermission('campaign.delete'))

// مقداردهی اولیه
onMounted(() => {
  campaigns.value = sampleCampaigns
})
</script> 
