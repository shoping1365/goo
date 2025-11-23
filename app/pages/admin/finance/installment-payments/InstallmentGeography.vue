<template>
  <div v-if="hasAccess" class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-6">
      <h3 class="text-lg font-semibold text-gray-900">توزیع جغرافیایی</h3>
      <div class="flex items-center space-x-2 space-x-reverse">
        <select v-model="selectedMetric" class="text-sm border border-gray-300 rounded-md px-3 py-1">
          <option value="count">تعداد درخواست</option>
          <option value="amount">مبلغ کل</option>
          <option value="approval_rate">نرخ تایید</option>
          <option value="avg_amount">میانگین مبلغ</option>
        </select>
        <select v-model="selectedPeriod" class="text-sm border border-gray-300 rounded-md px-3 py-1">
          <option value="7">7 روز گذشته</option>
          <option value="30">30 روز گذشته</option>
          <option value="90">90 روز گذشته</option>
          <option value="365">یک سال گذشته</option>
        </select>
        <button class="text-blue-600 hover:text-blue-800" @click="refreshData">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
        </button>
      </div>
    </div>

    <!-- Summary Cards -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
      <div class="bg-blue-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-blue-600 mb-1">{{ summary.totalProvinces }}</div>
        <div class="text-sm text-gray-600">استان‌های فعال</div>
      </div>
      <div class="bg-green-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-green-600 mb-1">{{ summary.totalCities }}</div>
        <div class="text-sm text-gray-600">شهرهای فعال</div>
      </div>
      <div class="bg-yellow-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-yellow-600 mb-1">{{ summary.topProvince }}</div>
        <div class="text-sm text-gray-600">استان برتر</div>
      </div>
      <div class="bg-purple-50 rounded-lg p-6 text-center">
        <div class="text-2xl font-bold text-purple-600 mb-1">{{ summary.coverage }}%</div>
        <div class="text-sm text-gray-600">پوشش جغرافیایی</div>
      </div>
    </div>

    <!-- Map Visualization -->
    <div class="mb-6">
      <h4 class="text-md font-semibold text-gray-900 mb-4">نقشه ایران</h4>
      <div class="bg-gray-100 rounded-lg p-6 text-center">
        <div class="text-gray-500 mb-4">
          <svg class="w-16 h-16 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 20l-5.447-2.724A1 1 0 013 16.382V5.618a1 1 0 011.447-.894L9 7m0 13l6-3m-6 3V7m6 10l4.553 2.276A1 1 0 0021 18.382V7.618a1 1 0 00-.553-.894L15 4m0 13V4m0 0L9 7" />
          </svg>
        </div>
        <p class="text-gray-600">نقشه تعاملی ایران</p>
        <p class="text-sm text-gray-500 mt-2">در اینجا نقشه تعاملی ایران با داده‌های اقساط نمایش داده می‌شود</p>
      </div>
    </div>

    <!-- Top Provinces -->
    <div class="mb-6">
      <h4 class="text-md font-semibold text-gray-900 mb-4">استان‌های برتر</h4>
      <div class="space-y-3">
        <div v-for="(province, index) in topProvinces" :key="province.name" class="bg-gray-50 rounded-lg p-6">
          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-8 h-8 rounded-full flex items-center justify-center mr-3" :class="getRankColor(index)">
                <span class="text-sm font-bold text-white">{{ index + 1 }}</span>
              </div>
              <div>
                <h5 class="text-lg font-semibold text-gray-900">{{ province.name }}</h5>
                <p class="text-sm text-gray-600">{{ province.cities }} شهر فعال</p>
              </div>
            </div>
            <div class="text-right">
              <div class="text-lg font-bold text-gray-900">{{ getMetricValue(province) }}</div>
              <div class="text-sm text-gray-500">{{ getMetricLabel() }}</div>
            </div>
          </div>
          
          <div class="mt-3 grid grid-cols-2 md:grid-cols-4 gap-6 text-sm">
            <div>
              <span class="text-gray-600 block">درخواست‌ها</span>
              <span class="font-medium">{{ province.requests }}</span>
            </div>
            <div>
              <span class="text-gray-600 block">مبلغ کل</span>
              <span class="font-medium">{{ formatCurrency(province.totalAmount) }}</span>
            </div>
            <div>
              <span class="text-gray-600 block">نرخ تایید</span>
              <span class="font-medium">{{ province.approvalRate }}%</span>
            </div>
            <div>
              <span class="text-gray-600 block">میانگین مبلغ</span>
              <span class="font-medium">{{ formatCurrency(province.avgAmount) }}</span>
            </div>
          </div>

          <div class="mt-3">
            <div class="flex items-center justify-between text-sm mb-1">
              <span class="text-gray-600">سهم از کل</span>
              <span class="font-medium">{{ province.percentage }}%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div class="bg-blue-500 h-2 rounded-full" :style="{ width: province.percentage + '%' }"></div>
            </div>
          </div>

          <div class="mt-3 flex items-center justify-between">
            <button class="text-blue-600 hover:text-blue-800 text-sm font-medium" @click="viewProvinceDetails(province)">
              جزئیات استان
            </button>
            <div class="flex items-center" :class="province.trend > 0 ? 'text-green-600' : province.trend < 0 ? 'text-red-600' : 'text-gray-600'">
              <svg v-if="province.trend > 0" class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
              </svg>
              <svg v-else-if="province.trend < 0" class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 17h8m0 0V9m0 8l-8-8-4 4-6-6" />
              </svg>
              <svg v-else class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h8" />
              </svg>
              <span class="text-sm font-medium">
                {{ province.trend > 0 ? '+' : '' }}{{ province.trend }}%
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Top Cities -->
    <div class="mb-6">
      <h4 class="text-md font-semibold text-gray-900 mb-4">شهرهای برتر</h4>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div v-for="city in topCities" :key="city.name" class="bg-gray-50 rounded-lg p-6">
          <div class="flex items-center justify-between mb-3">
            <div>
              <h5 class="text-md font-semibold text-gray-900">{{ city.name }}</h5>
              <p class="text-sm text-gray-600">{{ city.province }}</p>
            </div>
            <div class="text-right">
              <div class="text-lg font-bold text-gray-900">{{ city.requests }}</div>
              <div class="text-sm text-gray-500">درخواست</div>
            </div>
          </div>
          
          <div class="grid grid-cols-2 gap-6 text-sm mb-3">
            <div>
              <span class="text-gray-600 block">مبلغ کل</span>
              <span class="font-medium">{{ formatCurrency(city.totalAmount) }}</span>
            </div>
            <div>
              <span class="text-gray-600 block">نرخ تایید</span>
              <span class="font-medium">{{ city.approvalRate }}%</span>
            </div>
          </div>

          <div class="flex items-center justify-between">
            <div class="flex items-center justify-between text-sm">
              <span class="text-gray-600">سهم از استان:</span>
              <span class="font-medium">{{ city.provincePercentage }}%</span>
            </div>
            <button class="text-blue-600 hover:text-blue-800 text-sm font-medium" @click="viewCityDetails(city)">
              جزئیات
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Regional Comparison -->
    <div class="mb-6">
      <h4 class="text-md font-semibold text-gray-900 mb-4">مقایسه مناطق</h4>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-6">
        <div v-for="region in regionalData" :key="region.name" class="bg-gray-50 rounded-lg p-6 text-center">
          <h5 class="text-md font-semibold text-gray-900 mb-2">{{ region.name }}</h5>
          <div class="text-2xl font-bold text-blue-600 mb-1">{{ region.requests }}</div>
          <div class="text-sm text-gray-600 mb-3">درخواست</div>
          
          <div class="space-y-2 text-sm">
            <div class="flex justify-between">
              <span class="text-gray-600">مبلغ:</span>
              <span class="font-medium">{{ formatCurrency(region.amount) }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-gray-600">تایید:</span>
              <span class="font-medium">{{ region.approvalRate }}%</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Growth Analysis -->
    <div>
      <h4 class="text-md font-semibold text-gray-900 mb-4">تحلیل رشد</h4>
      <div class="bg-gray-50 rounded-lg p-6">
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div class="text-center">
            <div class="text-lg font-bold text-green-600 mb-1">{{ growthAnalysis.fastestGrowing }}</div>
            <div class="text-sm text-gray-600">سریع‌ترین رشد</div>
          </div>
          <div class="text-center">
            <div class="text-lg font-bold text-red-600 mb-1">{{ growthAnalysis.slowestGrowing }}</div>
            <div class="text-sm text-gray-600">کندترین رشد</div>
          </div>
          <div class="text-center">
            <div class="text-lg font-bold text-blue-600 mb-1">{{ growthAnalysis.mostPotential }}</div>
            <div class="text-sm text-gray-600">بیشترین پتانسیل</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
</script>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useAuth } from '~/composables/useAuth'

// احراز هویت
const { user, isAuthenticated } = useAuth()

// بررسی دسترسی admin
const hasAccess = computed(() => {
  if (!isAuthenticated.value) {
    return false
  }

  const userRole = user.value?.role?.toLowerCase() || ''
  const adminRoles = ['admin', 'developer']
  return adminRoles.includes(userRole)
})

// بررسی احراز هویت و دسترسی admin - نمایش 404 در صورت عدم دسترسی
const checkAuth = async () => {
  if (!hasAccess.value) {
    await navigateTo('/404', { external: false })
  }
}

// بررسی احراز هویت هنگام تغییر وضعیت احراز هویت
watch([isAuthenticated, hasAccess], async () => {
  if (!hasAccess.value) {
    await checkAuth()
  }
})

interface Province {
  name: string
  cities: number
  requests: number
  totalAmount: number
  approvalRate: number
  avgAmount: number
  percentage: number
  trend: number
}

interface City {
  name: string
  province: string
  requests: number
  totalAmount: number
  approvalRate: number
  provincePercentage: number
}

interface RegionalData {
  name: string
  requests: number
  amount: number
  approvalRate: number
}

interface Summary {
  totalProvinces: number
  totalCities: number
  topProvince: string
  coverage: number
}

interface GrowthAnalysis {
  fastestGrowing: string
  slowestGrowing: string
  mostPotential: string
}

const router = useRouter()

const selectedMetric = ref('count')
const selectedPeriod = ref('30')

const summary = ref<Summary>({
  totalProvinces: 0,
  totalCities: 0,
  topProvince: '',
  coverage: 0
})

const topProvinces = ref<Province[]>([])
const topCities = ref<City[]>([])
const regionalData = ref<RegionalData[]>([])
const growthAnalysis = ref<GrowthAnalysis>({
  fastestGrowing: '',
  slowestGrowing: '',
  mostPotential: ''
})

const formatCurrency = (amount: number): string => {
  return new Intl.NumberFormat('fa-IR', {
    style: 'currency',
    currency: 'IRR',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount)
}

const getRankColor = (index: number): string => {
  switch (index) {
    case 0: return 'bg-yellow-500' // Gold
    case 1: return 'bg-gray-400'   // Silver
    case 2: return 'bg-yellow-600' // Bronze
    default: return 'bg-blue-500'
  }
}

const getMetricValue = (province: Province): string => {
  switch (selectedMetric.value) {
    case 'count': return province.requests.toString()
    case 'amount': return formatCurrency(province.totalAmount)
    case 'approval_rate': return province.approvalRate + '%'
    case 'avg_amount': return formatCurrency(province.avgAmount)
    default: return ''
  }
}

const getMetricLabel = (): string => {
  switch (selectedMetric.value) {
    case 'count': return 'درخواست'
    case 'amount': return 'مبلغ کل'
    case 'approval_rate': return 'نرخ تایید'
    case 'avg_amount': return 'میانگین مبلغ'
    default: return ''
  }
}

const fetchGeographyData = async () => {
  try {
    interface GeographyData {
      summary: Summary
      topProvinces: Province[]
      topCities: City[]
      regionalData: RegionalData[]
      growthAnalysis: GrowthAnalysis
    }
    interface ApiResponse {
      data?: GeographyData
    }
    const response = await $fetch<ApiResponse>('/api/admin/installment-payments/geography', {
      query: { 
        period: selectedPeriod.value,
        metric: selectedMetric.value
      }
    })
    
    if (response.data) {
      summary.value = response.data.summary
      topProvinces.value = response.data.topProvinces
      topCities.value = response.data.topCities
      regionalData.value = response.data.regionalData
      growthAnalysis.value = response.data.growthAnalysis
    }
  } catch (error) {
    console.error('خطا در دریافت داده‌های جغرافیایی:', error)
    // در حالت واقعی، اینجا از createError استفاده می‌کنیم
  }
}

const refreshData = () => {
  fetchGeographyData()
}

const viewProvinceDetails = (province: Province) => {
  router.push(`/admin/finance/installment-payments/geography/province/${province.name}`)
}

const viewCityDetails = (city: City) => {
  router.push(`/admin/finance/installment-payments/geography/city/${city.name}`)
}

onMounted(async () => {
  await checkAuth()
  if (!hasAccess.value) {
    return
  }
  fetchGeographyData()
})

// Watch for changes
watch([selectedMetric, selectedPeriod], () => {
  fetchGeographyData()
})
</script>
