<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="flex justify-between items-center">
      <div>
        <h2 class="text-xl font-semibold text-gray-900">تحلیل SEO و رتبه‌بندی</h2>
        <p class="text-sm text-gray-600 mt-1">بررسی رتبه SEO، کلمات کلیدی و عملکرد جستجو</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <select v-model="selectedDomain" class="px-3 py-2 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-blue-500">
          <option value="main">دامنه اصلی</option>
          <option value="blog">وبلاگ</option>
          <option value="shop">فروشگاه</option>
        </select>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 text-sm" @click="refreshSeoData">
          <svg class="w-4 h-4 ml-2 inline" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
          </svg>
          به‌روزرسانی
        </button>
      </div>
    </div>

    <!-- SEO Score Overview -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <div class="bg-gradient-to-br from-green-50 to-green-100 rounded-xl p-6 border border-green-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-green-600">امتیاز کلی SEO</p>
            <p class="text-3xl font-bold text-green-900">{{ seoData.overallScore }}/100</p>
            <div class="w-full bg-green-200 rounded-full h-2 mt-2">
              <div class="h-2 bg-green-500 rounded-full" :style="{ width: seoData.overallScore + '%' }"></div>
            </div>
          </div>
          <div class="w-12 h-12 bg-green-500 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-br from-blue-50 to-blue-100 rounded-xl p-6 border border-blue-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-blue-600">کلمات کلیدی رتبه 1-3</p>
            <p class="text-2xl font-bold text-blue-900">{{ seoData.topKeywords }}</p>
            <p class="text-sm text-blue-700 mt-1">
              <span :class="seoData.topKeywordsGrowth >= 0 ? 'text-green-600' : 'text-red-600'">
                {{ seoData.topKeywordsGrowth >= 0 ? '+' : '' }}{{ seoData.topKeywordsGrowth }}
              </span>
              نسبت به ماه قبل
            </p>
          </div>
          <div class="w-12 h-12 bg-blue-500 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-br from-purple-50 to-purple-100 rounded-xl p-6 border border-purple-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-purple-600">ترافیک ارگانیک</p>
            <p class="text-2xl font-bold text-purple-900">{{ formatNumber(seoData.organicTraffic) }}</p>
            <p class="text-sm text-purple-700 mt-1">
              <span :class="seoData.organicTrafficGrowth >= 0 ? 'text-green-600' : 'text-red-600'">
                {{ seoData.organicTrafficGrowth >= 0 ? '+' : '' }}{{ seoData.organicTrafficGrowth }}%
              </span>
              نسبت به ماه قبل
            </p>
          </div>
          <div class="w-12 h-12 bg-purple-500 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-br from-orange-50 to-orange-100 rounded-xl p-6 border border-orange-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-orange-600">نرخ کلیک (CTR)</p>
            <p class="text-2xl font-bold text-orange-900">{{ seoData.ctr }}%</p>
            <p class="text-sm text-orange-700 mt-1">
              <span :class="seoData.ctrGrowth >= 0 ? 'text-green-600' : 'text-red-600'">
                {{ seoData.ctrGrowth >= 0 ? '+' : '' }}{{ seoData.ctrGrowth }}%
              </span>
              نسبت به ماه قبل
            </p>
          </div>
          <div class="w-12 h-12 bg-orange-500 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 15l-2 5L9 9l11 4-5 2zm0 0l5 5M7.188 2.239l.777 2.322M2.46 7.2l5.514-.458M13.95 4.05l-.95 5.514"></path>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- Keyword Rankings -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex justify-between items-center mb-6">
        <h3 class="text-lg font-semibold text-gray-900">رتبه کلمات کلیدی اصلی</h3>
        <div class="flex items-center space-x-3 space-x-reverse">
          <input 
            v-model="keywordSearch" 
            type="text" 
            placeholder="جستجو در کلمات کلیدی..."
            class="px-3 py-2 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
          <button class="px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 text-sm" @click="addNewKeyword">
            <svg class="w-4 h-4 ml-2 inline" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
            </svg>
            افزودن
          </button>
        </div>
      </div>
      
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">کلمه کلیدی</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">رتبه فعلی</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تغییر</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">حجم جستجو</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">سطح دشواری</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملکرد</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="keyword in filteredKeywords" :key="keyword.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <span class="text-sm font-medium text-gray-900">{{ keyword.term }}</span>
                  <span v-if="keyword.isNew" class="ml-2 inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800">
                    جدید
                  </span>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="text-sm font-medium text-gray-900">{{ keyword.currentRank }}</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span
:class="[
                  keyword.rankChange > 0 ? 'text-green-600' : keyword.rankChange < 0 ? 'text-red-600' : 'text-gray-500',
                  'text-sm font-medium'
                ]">
                  {{ keyword.rankChange > 0 ? '+' : '' }}{{ keyword.rankChange }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="text-sm text-gray-900">{{ formatNumber(keyword.searchVolume) }}</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="w-16 bg-gray-200 rounded-full h-2">
                    <div class="h-2 rounded-full" :class="getDifficultyColor(keyword.difficulty)" :style="{ width: keyword.difficulty + '%' }"></div>
                  </div>
                  <span class="text-sm text-gray-500 mr-2">{{ keyword.difficulty }}%</span>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span
:class="[
                  keyword.performance === 'excellent' ? 'bg-green-100 text-green-800' :
                  keyword.performance === 'good' ? 'bg-blue-100 text-blue-800' :
                  keyword.performance === 'fair' ? 'bg-yellow-100 text-yellow-800' :
                  'bg-red-100 text-red-800',
                  'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium'
                ]">
                  {{ getPerformanceText(keyword.performance) }}
                </span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- SEO Issues -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">مشکلات SEO</h3>
        <div class="space-y-4">
          <div v-for="issue in seoIssues" :key="issue.id" class="flex items-start space-x-3 space-x-reverse">
            <div
:class="[
              issue.severity === 'high' ? 'bg-red-100 text-red-600' :
              issue.severity === 'medium' ? 'bg-yellow-100 text-yellow-600' :
              'bg-blue-100 text-blue-600',
              'w-8 h-8 rounded-full flex items-center justify-center flex-shrink-0'
            ]">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path v-if="issue.severity === 'high'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
                <path v-else-if="issue.severity === 'medium'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
                <path v-else stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
            <div class="flex-1 min-w-0">
              <p class="text-sm font-medium text-gray-900">{{ issue.title }}</p>
              <p class="text-sm text-gray-600 mt-1">{{ issue.description }}</p>
              <div class="flex items-center space-x-4 space-x-reverse mt-2">
                <span
:class="[
                  issue.severity === 'high' ? 'text-red-600' :
                  issue.severity === 'medium' ? 'text-yellow-600' :
                  'text-blue-600',
                  'text-xs font-medium'
                ]">
                  {{ getSeverityText(issue.severity) }}
                </span>
                <button class="text-xs text-blue-600 hover:text-blue-800" @click="fixIssue(issue.id)">
                  رفع مشکل
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">توصیه‌های بهبود</h3>
        <div class="space-y-4">
          <div v-for="recommendation in seoRecommendations" :key="recommendation.id" class="p-6 bg-blue-50 rounded-lg border border-blue-200">
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
                  <button class="text-xs text-blue-600 hover:text-blue-800" @click="implementRecommendation(recommendation.id)">
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
const selectedDomain = ref('main')
const keywordSearch = ref('')

// داده‌های SEO
const seoData = ref({
  overallScore: 78,
  topKeywords: 24,
  topKeywordsGrowth: 3,
  organicTraffic: 45678,
  organicTrafficGrowth: 12.5,
  ctr: 3.2,
  ctrGrowth: 0.8
})

// کلمات کلیدی
const keywords = ref([
  { id: 1, term: 'فروشگاه آنلاین', currentRank: 3, rankChange: 2, searchVolume: 12000, difficulty: 65, performance: 'excellent', isNew: false },
  { id: 2, term: 'خرید اینترنتی', currentRank: 7, rankChange: -1, searchVolume: 8900, difficulty: 72, performance: 'good', isNew: false },
  { id: 3, term: 'پرداخت آنلاین', currentRank: 15, rankChange: 5, searchVolume: 5600, difficulty: 45, performance: 'good', isNew: false },
  { id: 4, term: 'ارسال رایگان', currentRank: 22, rankChange: 0, searchVolume: 3400, difficulty: 38, performance: 'fair', isNew: false },
  { id: 5, term: 'تخفیف ویژه', currentRank: 8, rankChange: 3, searchVolume: 7800, difficulty: 58, performance: 'excellent', isNew: true }
])

// مشکلات SEO
const seoIssues = ref([
  { id: 1, title: 'صفحات بدون متا توضیحات', description: '15 صفحه فاقد متا توضیحات هستند', severity: 'medium' },
  { id: 2, title: 'تصاویر بدون alt text', description: '23 تصویر فاقد متن جایگزین هستند', severity: 'high' },
  { id: 3, title: 'سرعت بارگذاری کند', description: 'زمان بارگذاری صفحه اصلی بیش از 3 ثانیه', severity: 'high' },
  { id: 4, title: 'لینک‌های شکسته', description: '8 لینک داخلی شکسته یافت شد', severity: 'low' }
])

// توصیه‌های بهبود
const seoRecommendations = ref([
  { id: 1, title: 'بهینه‌سازی تصاویر', description: 'فشرده‌سازی تصاویر برای بهبود سرعت بارگذاری', priority: 'بالا' },
  { id: 2, title: 'افزودن Schema Markup', description: 'استفاده از Schema.org برای بهبود نمایش در نتایج جستجو', priority: 'متوسط' },
  { id: 3, title: 'بهینه‌سازی محتوای داخلی', description: 'افزودن کلمات کلیدی مرتبط به صفحات موجود', priority: 'بالا' }
])

// computed properties
const filteredKeywords = computed(() => {
  if (!keywordSearch.value) return keywords.value
  return keywords.value.filter(keyword => 
    keyword.term.toLowerCase().includes(keywordSearch.value.toLowerCase())
  )
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

const getDifficultyColor = (difficulty) => {
  if (difficulty <= 30) return 'bg-green-500'
  if (difficulty <= 60) return 'bg-yellow-500'
  return 'bg-red-500'
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

const getSeverityText = (severity) => {
  const texts = {
    high: 'بالا',
    medium: 'متوسط',
    low: 'پایین'
  }
  return texts[severity] || severity
}

const refreshSeoData = () => {
  console.log('Refreshing SEO data...')
  // منطق به‌روزرسانی داده‌های SEO
}

const addNewKeyword = () => {
  console.log('Adding new keyword...')
  // منطق افزودن کلمه کلیدی جدید
}

const fixIssue = (issueId) => {
  console.log('Fixing issue:', issueId)
  // منطق رفع مشکل
}

const implementRecommendation = (recommendationId) => {
  console.log('Implementing recommendation:', recommendationId)
  // منطق پیاده‌سازی توصیه
}
</script>
