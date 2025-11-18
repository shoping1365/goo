<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="flex justify-between items-center">
      <div>
        <h2 class="text-xl font-semibold text-gray-900">تبلیغات دیجیتال</h2>
        <p class="text-sm text-gray-600 mt-1">پایش عملکرد تبلیغات Google Ads و Social Ads</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <select v-model="selectedPlatform" class="px-3 py-2 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-blue-500">
          <option value="all">همه پلتفرم‌ها</option>
          <option value="google">Google Ads</option>
          <option value="facebook">Facebook Ads</option>
          <option value="instagram">Instagram Ads</option>
          <option value="linkedin">LinkedIn Ads</option>
        </select>
        <button @click="refreshAdsData" class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 text-sm">
          <svg class="w-4 h-4 ml-2 inline" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
          </svg>
          به‌روزرسانی
        </button>
      </div>
    </div>

    <!-- Performance Overview -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <div class="bg-gradient-to-br from-blue-50 to-blue-100 rounded-xl p-6 border border-blue-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-blue-600">هزینه کل</p>
            <p class="text-2xl font-bold text-blue-900">{{ formatCurrency(adsData.totalSpent) }}</p>
            <p class="text-sm text-blue-700 mt-1">
              <span :class="adsData.spentGrowth >= 0 ? 'text-red-600' : 'text-green-600'">
                {{ adsData.spentGrowth >= 0 ? '+' : '' }}{{ adsData.spentGrowth }}%
              </span>
              نسبت به ماه قبل
            </p>
          </div>
          <div class="w-12 h-12 bg-blue-500 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-br from-green-50 to-green-100 rounded-xl p-6 border border-green-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-green-600">کلیک‌ها</p>
            <p class="text-2xl font-bold text-green-900">{{ formatNumber(adsData.totalClicks) }}</p>
            <p class="text-sm text-green-700 mt-1">
              <span :class="adsData.clicksGrowth >= 0 ? 'text-green-600' : 'text-red-600'">
                {{ adsData.clicksGrowth >= 0 ? '+' : '' }}{{ adsData.clicksGrowth }}%
              </span>
              نسبت به ماه قبل
            </p>
          </div>
          <div class="w-12 h-12 bg-green-500 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 15l-2 5L9 9l11 4-5 2zm0 0l5 5M7.188 2.239l.777 2.322M2.46 7.2l5.514-.458M13.95 4.05l-.95 5.514"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-br from-purple-50 to-purple-100 rounded-xl p-6 border border-purple-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-purple-600">نمایش‌ها</p>
            <p class="text-2xl font-bold text-purple-900">{{ formatNumber(adsData.totalImpressions) }}</p>
            <p class="text-sm text-purple-700 mt-1">
              <span :class="adsData.impressionsGrowth >= 0 ? 'text-green-600' : 'text-red-600'">
                {{ adsData.impressionsGrowth >= 0 ? '+' : '' }}{{ adsData.impressionsGrowth }}%
              </span>
              نسبت به ماه قبل
            </p>
          </div>
          <div class="w-12 h-12 bg-purple-500 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-br from-orange-50 to-orange-100 rounded-xl p-6 border border-orange-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-orange-600">نرخ کلیک (CTR)</p>
            <p class="text-2xl font-bold text-orange-900">{{ adsData.ctr }}%</p>
            <p class="text-sm text-orange-700 mt-1">
              <span :class="adsData.ctrGrowth >= 0 ? 'text-green-600' : 'text-red-600'">
                {{ adsData.ctrGrowth >= 0 ? '+' : '' }}{{ adsData.ctrGrowth }}%
              </span>
              نسبت به ماه قبل
            </p>
          </div>
          <div class="w-12 h-12 bg-orange-500 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- Platform Performance -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-6">عملکرد پلتفرم‌ها</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <div v-for="platform in platformPerformance" :key="platform.name" class="text-center">
          <div class="w-16 h-16 mx-auto mb-3 rounded-full flex items-center justify-center" :style="{ backgroundColor: platform.color + '20' }">
            <component :is="platform.icon" class="w-8 h-8" :style="{ color: platform.color }" />
          </div>
          <h4 class="text-sm font-medium text-gray-900 mb-2">{{ platform.name }}</h4>
          <div class="space-y-2">
            <div class="flex justify-between text-xs">
              <span class="text-gray-500">هزینه:</span>
              <span class="font-medium">{{ formatCurrency(platform.spent) }}</span>
            </div>
            <div class="flex justify-between text-xs">
              <span class="text-gray-500">کلیک:</span>
              <span class="font-medium">{{ formatNumber(platform.clicks) }}</span>
            </div>
            <div class="flex justify-between text-xs">
              <span class="text-gray-500">CTR:</span>
              <span class="font-medium">{{ platform.ctr }}%</span>
            </div>
            <div class="flex justify-between text-xs">
              <span class="text-gray-500">CPC:</span>
              <span class="font-medium">{{ formatCurrency(platform.cpc) }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Campaign Performance -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex justify-between items-center mb-6">
        <h3 class="text-lg font-semibold text-gray-900">عملکرد کمپین‌ها</h3>
        <div class="flex items-center space-x-3 space-x-reverse">
          <select v-model="selectedCampaignType" class="px-3 py-2 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option value="all">همه کمپین‌ها</option>
            <option value="brand">برندینگ</option>
            <option value="performance">عملکرد</option>
            <option value="retargeting">بازاریابی مجدد</option>
          </select>
          <button @click="createCampaign" class="px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 text-sm">
            <svg class="w-4 h-4 ml-2 inline" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
            </svg>
            کمپین جدید
          </button>
        </div>
      </div>
      
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نام کمپین</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">پلتفرم</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">هزینه</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کلیک‌ها</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">CTR</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">CPC</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملکرد</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="campaign in filteredCampaigns" :key="campaign.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <span class="text-sm font-medium text-gray-900">{{ campaign.name }}</span>
                  <span v-if="campaign.isNew" class="ml-2 inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800">
                    جدید
                  </span>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="w-6 h-6 rounded-full flex items-center justify-center" :style="{ backgroundColor: getPlatformColor(campaign.platform) + '20' }">
                    <component :is="getPlatformIcon(campaign.platform)" class="w-3 h-3" :style="{ color: getPlatformColor(campaign.platform) }" />
                  </div>
                  <span class="text-sm text-gray-900 mr-2">{{ getPlatformName(campaign.platform) }}</span>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="[
                  campaign.status === 'active' ? 'bg-green-100 text-green-800' :
                  campaign.status === 'paused' ? 'bg-yellow-100 text-yellow-800' :
                  'bg-red-100 text-red-800',
                  'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium'
                ]">
                  {{ getStatusText(campaign.status) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="text-sm text-gray-900">{{ formatCurrency(campaign.spent) }}</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="text-sm text-gray-900">{{ formatNumber(campaign.clicks) }}</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="text-sm text-gray-900">{{ campaign.ctr }}%</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="text-sm text-gray-900">{{ formatCurrency(campaign.cpc) }}</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="[
                  campaign.performance === 'excellent' ? 'bg-green-100 text-green-800' :
                  campaign.performance === 'good' ? 'bg-blue-100 text-blue-800' :
                  campaign.performance === 'fair' ? 'bg-yellow-100 text-yellow-800' :
                  'bg-red-100 text-red-800',
                  'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium'
                ]">
                  {{ getPerformanceText(campaign.performance) }}
                </span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Ad Creative Performance -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">عملکرد خلاقیت‌های تبلیغاتی</h3>
        <div class="space-y-4">
          <div v-for="creative in topCreatives" :key="creative.id" class="flex items-center space-x-3 space-x-reverse">
            <div class="w-16 h-16 bg-gray-200 rounded-lg flex items-center justify-center flex-shrink-0">
              <svg class="w-8 h-8 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
              </svg>
            </div>
            <div class="flex-1 min-w-0">
              <p class="text-sm font-medium text-gray-900">{{ creative.title }}</p>
              <p class="text-xs text-gray-500">{{ creative.platform }}</p>
              <div class="flex items-center space-x-4 space-x-reverse mt-1">
                <span class="text-xs text-gray-600">{{ creative.ctr }}% CTR</span>
                <span class="text-xs text-gray-600">{{ formatCurrency(creative.cpc) }} CPC</span>
              </div>
            </div>
            <div class="text-right">
              <span :class="[
                creative.performance === 'excellent' ? 'text-green-600' :
                creative.performance === 'good' ? 'text-blue-600' :
                creative.performance === 'fair' ? 'text-yellow-600' :
                'text-red-600',
                'text-xs font-medium'
              ]">
                {{ getPerformanceText(creative.performance) }}
              </span>
            </div>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">توصیه‌های بهینه‌سازی</h3>
        <div class="space-y-4">
          <div v-for="recommendation in adRecommendations" :key="recommendation.id" class="p-6 bg-blue-50 rounded-lg border border-blue-200">
            <div class="flex items-start space-x-3 space-x-reverse">
              <div class="w-6 h-6 bg-blue-500 rounded-full flex items-center justify-center flex-shrink-0">
                <svg class="w-3 h-3 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"></path>
                </svg>
              </div>
              <div class="flex-1">
                <p class="text-sm font-medium text-blue-900">{{ recommendation.title }}</p>
                <p class="text-sm text-blue-700 mt-1">{{ recommendation.description }}</p>
                <div class="flex items-center space-x-3 space-x-reverse mt-2">
                  <span class="text-xs text-blue-600 bg-blue-100 px-2 py-1 rounded-full">
                    {{ recommendation.priority }}
                  </span>
                  <button @click="implementRecommendation(recommendation.id)" class="text-xs text-blue-600 hover:text-blue-800">
                    پیاده‌سازی
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// متغیرهای reactive
const selectedPlatform = ref('all')
const selectedCampaignType = ref('all')

// داده‌های تبلیغات
const adsData = ref({
  totalSpent: 45678.90,
  spentGrowth: 15.3,
  totalClicks: 12345,
  clicksGrowth: 8.7,
  totalImpressions: 234567,
  impressionsGrowth: 12.4,
  ctr: 5.3,
  ctrGrowth: 2.1
})

// عملکرد پلتفرم‌ها
const platformPerformance = ref([
  { name: 'Google Ads', spent: 23456.78, clicks: 6789, ctr: 6.2, cpc: 3.45, color: '#4285F4', icon: 'GoogleIcon' },
  { name: 'Facebook Ads', spent: 12345.67, clicks: 3456, ctr: 4.8, cpc: 3.57, color: '#1877F2', icon: 'FacebookIcon' },
  { name: 'Instagram Ads', spent: 6789.12, clicks: 1234, ctr: 3.9, cpc: 5.50, color: '#E4405F', icon: 'InstagramIcon' },
  { name: 'LinkedIn Ads', spent: 3087.33, clicks: 866, ctr: 4.2, cpc: 3.57, color: '#0A66C2', icon: 'LinkedInIcon' }
])

// کمپین‌ها
const campaigns = ref([
  { id: 1, name: 'کمپین برندینگ Q1', platform: 'google', status: 'active', spent: 12345.67, clicks: 3456, ctr: 5.2, cpc: 3.57, performance: 'excellent', isNew: false },
  { id: 2, name: 'تبلیغات محصولات جدید', platform: 'facebook', status: 'active', spent: 8901.23, clicks: 2345, ctr: 4.8, cpc: 3.80, performance: 'good', isNew: false },
  { id: 3, name: 'بازاریابی مجدد', platform: 'instagram', status: 'paused', spent: 5678.90, clicks: 1234, ctr: 3.9, cpc: 4.60, performance: 'fair', isNew: false },
  { id: 4, name: 'کمپین B2B', platform: 'linkedin', status: 'active', spent: 3456.78, clicks: 789, ctr: 4.2, cpc: 4.38, performance: 'good', isNew: true }
])

// خلاقیت‌های برتر
const topCreatives = ref([
  { id: 1, title: 'تصویر محصول A', platform: 'Google Ads', ctr: 6.8, cpc: 3.20, performance: 'excellent' },
  { id: 2, title: 'ویدیو معرفی', platform: 'Facebook Ads', ctr: 5.4, cpc: 3.80, performance: 'good' },
  { id: 3, title: 'کاروسل محصولات', platform: 'Instagram Ads', ctr: 4.2, cpc: 4.50, performance: 'fair' }
])

// توصیه‌های بهینه‌سازی
const adRecommendations = ref([
  { id: 1, title: 'بهینه‌سازی کلمات کلیدی', description: 'حذف کلمات کلیدی با عملکرد ضعیف و افزودن کلمات جدید', priority: 'بالا' },
  { id: 2, title: 'تست A/B خلاقیت‌ها', description: 'آزمایش انواع مختلف تصاویر و متن‌های تبلیغاتی', priority: 'متوسط' },
  { id: 3, title: 'بهبود هدف‌گذاری', description: 'تنظیم مجدد مخاطبان هدف بر اساس داده‌های عملکرد', priority: 'بالا' }
])

// computed properties
const filteredCampaigns = computed(() => {
  let filtered = campaigns.value
  
  if (selectedPlatform.value !== 'all') {
    filtered = filtered.filter(campaign => campaign.platform === selectedPlatform.value)
  }
  
  if (selectedCampaignType.value !== 'all') {
    // منطق فیلتر بر اساس نوع کمپین
  }
  
  return filtered
})

// توابع
const formatNumber = (num) => {
  if (num >= 1000000) {
    return (num / 1000000).toFixed(1) + 'M'
  } else if (num >= 1000) {
    return (num / 1000).toFixed(1) + 'K'
  }
  return num.toLocaleString('fa-IR')
}

const formatCurrency = (amount) => {
  return new Intl.NumberFormat('fa-IR', {
    style: 'currency',
    currency: 'IRR',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount)
}

const getPlatformColor = (platform) => {
  const colors = {
    google: '#4285F4',
    facebook: '#1877F2',
    instagram: '#E4405F',
    linkedin: '#0A66C2'
  }
  return colors[platform] || '#6B7280'
}

const getPlatformName = (platform) => {
  const names = {
    google: 'Google Ads',
    facebook: 'Facebook Ads',
    instagram: 'Instagram Ads',
    linkedin: 'LinkedIn Ads'
  }
  return names[platform] || platform
}

const getPlatformIcon = (platform) => {
  const icons = {
    google: 'GoogleIcon',
    facebook: 'FacebookIcon',
    instagram: 'InstagramIcon',
    linkedin: 'LinkedInIcon'
  }
  return icons[platform] || 'DefaultIcon'
}

const getStatusText = (status) => {
  const texts = {
    active: 'فعال',
    paused: 'متوقف',
    inactive: 'غیرفعال'
  }
  return texts[status] || status
}

const getPerformanceText = (performance) => {
  const texts = {
    excellent: 'عالی',
    good: 'خوب',
    fair: 'متوسط',
    poor: 'ضعیف'
  }
  return texts[performance] || performance
}

const refreshAdsData = () => {
  console.log('Refreshing ads data...')
  // منطق به‌روزرسانی داده‌های تبلیغات
}

const createCampaign = () => {
  console.log('Creating new campaign...')
  // منطق ایجاد کمپین جدید
}

const implementRecommendation = (recommendationId) => {
  console.log('Implementing recommendation:', recommendationId)
  // منطق پیاده‌سازی توصیه
}

// آیکون‌های پلتفرم‌ها
const GoogleIcon = defineComponent({
  template: `
    <svg class="w-6 h-6" viewBox="0 0 24 24" fill="currentColor">
      <path d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"/>
      <path d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"/>
      <path d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"/>
      <path d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"/>
    </svg>
  `
})

const FacebookIcon = defineComponent({
  template: `
    <svg class="w-6 h-6" viewBox="0 0 24 24" fill="currentColor">
      <path d="M24 12.073c0-6.627-5.373-12-12-12s-12 5.373-12 12c0 5.99 4.388 10.954 10.125 11.854v-8.385H7.078v-3.47h3.047V9.43c0-3.007 1.792-4.669 4.533-4.669 1.312 0 2.686.235 2.686.235v2.953H15.83c-1.491 0-1.956.925-1.956 1.874v2.25h3.328l-.532 3.47h-2.796v8.385C19.612 23.027 24 18.062 24 12.073z"/>
    </svg>
  `
})

const InstagramIcon = defineComponent({
  template: `
    <svg class="w-6 h-6" viewBox="0 0 24 24" fill="currentColor">
      <path d="M12 2.163c3.204 0 3.584.012 4.85.07 3.252.148 4.771 1.691 4.919 4.919.058 1.265.069 1.645.069 4.849 0 3.205-.012 3.584-.069 4.849-.149 3.225-1.664 4.771-4.919 4.919-1.266.058-1.644.07-4.85.07-3.204 0-3.584-.012-4.849-.07-3.26-.149-4.771-1.699-4.919-4.92-.058-1.265-.07-1.644-.07-4.849 0-3.204.013-3.583.07-4.849.149-3.227 1.664-4.771 4.919-4.919 1.266-.057 1.645-.069 4.849-.069zm0-2.163c-3.259 0-3.667.014-4.947.072-4.358.2-6.78 2.618-6.98 6.98-.059 1.281-.073 1.689-.073 4.948 0 3.259.014 3.668.072 4.948.2 4.358 2.618 6.78 6.98 6.98 1.281.058 1.689.072 4.948.072 3.259 0 3.668-.014 4.948-.072 4.354-.2 6.782-2.618 6.979-6.98.059-1.28.073-1.689.073-4.948 0-3.259-.014-3.667-.072-4.947-.196-4.354-2.617-6.78-6.979-6.98-1.281-.059-1.69-.073-4.949-.073zm0 5.838c-3.403 0-6.162 2.759-6.162 6.162s2.759 6.163 6.162 6.163 6.162-2.759 6.162-6.163c0-3.403-2.759-6.162-6.162-6.162zm0 10.162c-2.209 0-4-1.79-4-4 0-2.209 1.791-4 4-4s4 1.791 4 4c0 2.21-1.791 4-4 4zm6.406-11.845c-.796 0-1.441.645-1.441 1.44s.645 1.44 1.441 1.44c.795 0 1.439-.645 1.439-1.44s-.644-1.44-1.439-1.44z"/>
    </svg>
  `
})

const LinkedInIcon = defineComponent({
  template: `
    <svg class="w-6 h-6" viewBox="0 0 24 24" fill="currentColor">
      <path d="M20.447 20.452h-3.554v-5.569c0-1.328-.027-3.047-1.852-3.047-1.853 0-2.136 1.445-2.136 2.939v5.677H9.351V9h3.414v1.561h.046c.477-.9 1.637-1.85 3.37-1.85 3.601 0 4.267 2.37 4.267 5.455v6.286zM5.337 7.433c-1.144 0-2.063-.926-2.063-2.065 0-1.138.92-2.063 2.063-2.063 1.14 0 2.064.925 2.064 2.063 0 1.139-.925 2.065-2.064 2.065zm1.782 13.019H3.555V9h3.564v11.452zM22.225 0H1.771C.792 0 0 .774 0 1.729v20.542C0 23.227.792 24 1.771 24h20.451C23.2 24 24 23.227 24 22.271V1.729C24 .774 23.2 0 22.222 0h.003z"/>
    </svg>
  `
})
</script>
