<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="flex justify-between items-center">
      <div>
        <h2 class="text-xl font-semibold text-gray-900">پایش شبکه‌های اجتماعی</h2>
        <p class="text-sm text-gray-600 mt-1">نرخ تعامل و رشد شبکه‌های اجتماعی</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <select v-model="selectedPlatform" class="px-3 py-2 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-blue-500">
          <option value="all">همه پلتفرم‌ها</option>
          <option value="instagram">Instagram</option>
          <option value="telegram">Telegram</option>
          <option value="twitter">Twitter</option>
          <option value="linkedin">LinkedIn</option>
        </select>
        <button class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 text-sm" @click="refreshSocialData">
          <svg class="w-4 h-4 ml-2 inline" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
          </svg>
          به‌روزرسانی
        </button>
      </div>
    </div>

    <!-- Social Media Overview -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <div class="bg-gradient-to-br from-pink-50 to-pink-100 rounded-xl p-6 border border-pink-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-pink-600">کل فالوورها</p>
            <p class="text-2xl font-bold text-pink-900">{{ formatNumber(socialData.totalFollowers) }}</p>
            <p class="text-sm text-pink-700 mt-1">
              <span :class="socialData.followersGrowth >= 0 ? 'text-green-600' : 'text-red-600'">
                {{ socialData.followersGrowth >= 0 ? '+' : '' }}{{ socialData.followersGrowth }}%
              </span>
              نسبت به ماه قبل
            </p>
          </div>
          <div class="w-12 h-12 bg-pink-500 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-br from-blue-50 to-blue-100 rounded-xl p-6 border border-blue-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-blue-600">نرخ تعامل</p>
            <p class="text-2xl font-bold text-blue-900">{{ socialData.engagementRate }}%</p>
            <p class="text-sm text-blue-700 mt-1">
              <span :class="socialData.engagementGrowth >= 0 ? 'text-green-600' : 'text-red-600'">
                {{ socialData.engagementGrowth >= 0 ? '+' : '' }}{{ socialData.engagementGrowth }}%
              </span>
              نسبت به ماه قبل
            </p>
          </div>
          <div class="w-12 h-12 bg-blue-500 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-br from-green-50 to-green-100 rounded-xl p-6 border border-green-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-green-600">پست‌های منتشر شده</p>
            <p class="text-2xl font-bold text-green-900">{{ formatNumber(socialData.totalPosts) }}</p>
            <p class="text-sm text-green-700 mt-1">
              <span :class="socialData.postsGrowth >= 0 ? 'text-green-600' : 'text-red-600'">
                {{ socialData.postsGrowth >= 0 ? '+' : '' }}{{ socialData.postsGrowth }}%
              </span>
              نسبت به ماه قبل
            </p>
          </div>
          <div class="w-12 h-12 bg-green-500 rounded-lg flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-br from-purple-50 to-purple-100 rounded-xl p-6 border border-purple-200">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-purple-600">تعداد بازدید</p>
            <p class="text-2xl font-bold text-purple-900">{{ formatNumber(socialData.totalViews) }}</p>
            <p class="text-sm text-purple-700 mt-1">
              <span :class="socialData.viewsGrowth >= 0 ? 'text-green-600' : 'text-red-600'">
                {{ socialData.viewsGrowth >= 0 ? '+' : '' }}{{ socialData.viewsGrowth }}%
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
              <span class="text-gray-500">فالوور:</span>
              <span class="font-medium">{{ formatNumber(platform.followers) }}</span>
            </div>
            <div class="flex justify-between text-xs">
              <span class="text-gray-500">تعامل:</span>
              <span class="font-medium">{{ platform.engagement }}%</span>
            </div>
            <div class="flex justify-between text-xs">
              <span class="text-gray-500">پست:</span>
              <span class="font-medium">{{ platform.posts }}</span>
            </div>
            <div class="flex justify-between text-xs">
              <span class="text-gray-500">رشد:</span>
              <span class="font-medium" :class="platform.growth >= 0 ? 'text-green-600' : 'text-red-600'">
                {{ platform.growth >= 0 ? '+' : '' }}{{ platform.growth }}%
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Top Performing Posts -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex justify-between items-center mb-6">
        <h3 class="text-lg font-semibold text-gray-900">پست‌های برتر</h3>
        <div class="flex items-center space-x-3 space-x-reverse">
          <select v-model="selectedPostType" class="px-3 py-2 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option value="all">همه انواع</option>
            <option value="image">تصویر</option>
            <option value="video">ویدیو</option>
            <option value="story">استوری</option>
            <option value="reel">ریل</option>
          </select>
          <button class="px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 text-sm" @click="createPost">
            <svg class="w-4 h-4 ml-2 inline" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
            </svg>
            پست جدید
          </button>
        </div>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div v-for="post in filteredPosts" :key="post.id" class="bg-gray-50 rounded-lg p-6 border border-gray-200">
          <div class="flex items-start space-x-3 space-x-reverse mb-3">
            <div class="w-12 h-12 bg-gray-200 rounded-lg flex items-center justify-center flex-shrink-0">
              <svg class="w-6 h-6 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path v-if="post.type === 'image'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                <path v-else-if="post.type === 'video'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
                <path v-else stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
              </svg>
            </div>
            <div class="flex-1 min-w-0">
              <p class="text-sm font-medium text-gray-900">{{ post.title }}</p>
              <p class="text-xs text-gray-500">{{ getPlatformName(post.platform) }} • {{ formatDate(post.date) }}</p>
            </div>
            <span
:class="[
              post.performance === 'excellent' ? 'bg-green-100 text-green-800' :
              post.performance === 'good' ? 'bg-blue-100 text-blue-800' :
              post.performance === 'fair' ? 'bg-yellow-100 text-yellow-800' :
              'bg-red-100 text-red-800',
              'inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium'
            ]">
              {{ getPerformanceText(post.performance) }}
            </span>
          </div>
          
          <div class="grid grid-cols-3 gap-2 text-center">
            <div class="bg-white rounded p-2">
              <p class="text-xs text-gray-500">لایک</p>
              <p class="text-sm font-semibold text-gray-900">{{ formatNumber(post.likes) }}</p>
            </div>
            <div class="bg-white rounded p-2">
              <p class="text-xs text-gray-500">کامنت</p>
              <p class="text-sm font-semibold text-gray-900">{{ formatNumber(post.comments) }}</p>
            </div>
            <div class="bg-white rounded p-2">
              <p class="text-xs text-gray-500">اشتراک</p>
              <p class="text-sm font-semibold text-gray-900">{{ formatNumber(post.shares) }}</p>
            </div>
          </div>
          
          <div class="mt-3 pt-3 border-t border-gray-200">
            <div class="flex justify-between items-center text-xs">
              <span class="text-gray-500">نرخ تعامل:</span>
              <span class="font-medium text-blue-600">{{ post.engagementRate }}%</span>
            </div>
            <div class="flex justify-between items-center text-xs mt-1">
              <span class="text-gray-500">دسترسی:</span>
              <span class="font-medium text-gray-900">{{ formatNumber(post.reach) }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Engagement Analytics -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">تحلیل تعامل</h3>
        <div class="space-y-4">
          <div v-for="metric in engagementMetrics" :key="metric.name" class="flex items-center justify-between">
            <div class="flex items-center space-x-3 space-x-reverse">
              <div class="w-3 h-3 rounded-full" :style="{ backgroundColor: metric.color }"></div>
              <span class="text-sm font-medium text-gray-700">{{ metric.name }}</span>
            </div>
            <div class="flex items-center space-x-3 space-x-reverse">
              <span class="text-sm text-gray-600">{{ metric.value }}</span>
              <div class="w-24 bg-gray-200 rounded-full h-2">
                <div class="h-2 rounded-full" :style="{ width: metric.percentage + '%', backgroundColor: metric.color }"></div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <h3 class="text-lg font-semibold text-gray-900 mb-4">بهترین زمان‌های انتشار</h3>
        <div class="space-y-4">
          <div v-for="timeSlot in bestPostingTimes" :key="timeSlot.day" class="flex items-center justify-between">
            <div class="flex items-center space-x-3 space-x-reverse">
              <div class="w-8 h-8 bg-blue-100 rounded-lg flex items-center justify-center">
                <span class="text-xs font-medium text-blue-600">{{ timeSlot.day }}</span>
              </div>
              <span class="text-sm font-medium text-gray-700">{{ timeSlot.time }}</span>
            </div>
            <div class="text-right">
              <p class="text-sm font-semibold text-gray-900">{{ timeSlot.engagement }}%</p>
              <p class="text-xs text-gray-500">نرخ تعامل</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Competitor Analysis -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-6">تحلیل رقبا</h3>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">رقبا</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">فالوور</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نرخ تعامل</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">فرکانس پست</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملکرد</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">اقدام</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="competitor in competitors" :key="competitor.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="w-8 h-8 bg-gray-200 rounded-full flex items-center justify-center">
                    <span class="text-sm font-medium text-gray-600">{{ competitor.name.charAt(0) }}</span>
                  </div>
                  <span class="text-sm font-medium text-gray-900 mr-3">{{ competitor.name }}</span>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="text-sm text-gray-900">{{ formatNumber(competitor.followers) }}</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="text-sm text-gray-900">{{ competitor.engagement }}%</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="text-sm text-gray-900">{{ competitor.postingFrequency }}/هفته</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span
:class="[
                  competitor.performance === 'excellent' ? 'bg-green-100 text-green-800' :
                  competitor.performance === 'good' ? 'bg-blue-100 text-blue-800' :
                  competitor.performance === 'fair' ? 'bg-yellow-100 text-yellow-800' :
                  'bg-red-100 text-red-800',
                  'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium'
                ]">
                  {{ getPerformanceText(competitor.performance) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <button class="text-sm text-blue-600 hover:text-blue-800" @click="analyzeCompetitor(competitor.id)">
                  تحلیل
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup>
// متغیرهای reactive
const selectedPlatform = ref('all')
const selectedPostType = ref('all')

// داده‌های شبکه‌های اجتماعی
const socialData = ref({
  totalFollowers: 125430,
  followersGrowth: 8.5,
  engagementRate: 4.2,
  engagementGrowth: 12.3,
  totalPosts: 456,
  postsGrowth: 15.7,
  totalViews: 234567,
  viewsGrowth: 18.9
})

// عملکرد پلتفرم‌ها
const platformPerformance = ref([
  { name: 'Instagram', followers: 67890, engagement: 5.2, posts: 234, growth: 12.5, color: '#E4405F', icon: 'InstagramIcon' },
  { name: 'Telegram', followers: 34567, engagement: 3.8, posts: 123, growth: 8.7, color: '#0088CC', icon: 'TelegramIcon' },
  { name: 'Twitter', followers: 15678, engagement: 4.1, posts: 67, growth: 6.3, color: '#1DA1F2', icon: 'TwitterIcon' },
  { name: 'LinkedIn', followers: 7295, engagement: 2.9, posts: 32, growth: 4.2, color: '#0A66C2', icon: 'LinkedInIcon' }
])

// پست‌های برتر
const posts = ref([
  { id: 1, title: 'معرفی محصول جدید', platform: 'instagram', type: 'image', date: '2024-01-15', likes: 1234, comments: 89, shares: 45, engagementRate: 6.8, reach: 18900, performance: 'excellent' },
  { id: 2, title: 'ویدیو آموزشی', platform: 'instagram', type: 'video', date: '2024-01-14', likes: 987, comments: 67, shares: 34, engagementRate: 5.4, reach: 15600, performance: 'good' },
  { id: 3, title: 'استوری روزانه', platform: 'instagram', type: 'story', date: '2024-01-13', likes: 456, comments: 23, shares: 12, engagementRate: 4.2, reach: 8900, performance: 'fair' },
  { id: 4, title: 'پست محصول', platform: 'telegram', type: 'image', date: '2024-01-12', likes: 234, comments: 15, shares: 8, engagementRate: 3.8, reach: 6700, performance: 'good' },
  { id: 5, title: 'ریل سرگرمی', platform: 'instagram', type: 'reel', date: '2024-01-11', likes: 876, comments: 45, shares: 23, engagementRate: 5.1, reach: 12300, performance: 'good' },
  { id: 6, title: 'پست اطلاع‌رسانی', platform: 'twitter', type: 'text', date: '2024-01-10', likes: 123, comments: 12, shares: 5, engagementRate: 2.9, reach: 4500, performance: 'fair' }
])

// متریک‌های تعامل
const engagementMetrics = ref([
  { name: 'لایک', value: '45.2%', percentage: 45.2, color: '#3B82F6' },
  { name: 'کامنت', value: '28.7%', percentage: 28.7, color: '#10B981' },
  { name: 'اشتراک', value: '18.3%', percentage: 18.3, color: '#8B5CF6' },
  { name: 'ذخیره', value: '7.8%', percentage: 7.8, color: '#F59E0B' }
])

// بهترین زمان‌های انتشار
const bestPostingTimes = ref([
  { day: 'شنبه', time: '18:00-20:00', engagement: 6.8 },
  { day: 'یکشنبه', time: '19:00-21:00', engagement: 6.5 },
  { day: 'دوشنبه', time: '17:00-19:00', engagement: 6.2 },
  { day: 'سه‌شنبه', time: '18:00-20:00', engagement: 6.0 },
  { day: 'چهارشنبه', time: '19:00-21:00', engagement: 5.8 },
  { day: 'پنج‌شنبه', time: '16:00-18:00', engagement: 5.5 },
  { day: 'جمعه', time: '15:00-17:00', engagement: 5.2 }
])

// رقبا
const competitors = ref([
  { id: 1, name: 'رقبای اصلی', followers: 234567, engagement: 4.8, postingFrequency: 5, performance: 'excellent' },
  { id: 2, name: 'رقبای متوسط', followers: 123456, engagement: 3.9, postingFrequency: 4, performance: 'good' },
  { id: 3, name: 'رقبای کوچک', followers: 67890, engagement: 3.2, postingFrequency: 3, performance: 'fair' }
])

// computed properties
const filteredPosts = computed(() => {
  let filtered = posts.value
  
  if (selectedPlatform.value !== 'all') {
    filtered = filtered.filter(post => post.platform === selectedPlatform.value)
  }
  
  if (selectedPostType.value !== 'all') {
    filtered = filtered.filter(post => post.type === selectedPostType.value)
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

const formatDate = (dateString) => {
  const date = new Date(dateString)
  return date.toLocaleDateString('fa-IR')
}

const getPlatformName = (platform) => {
  const names = {
    instagram: 'Instagram',
    telegram: 'Telegram',
    twitter: 'Twitter',
    linkedin: 'LinkedIn'
  }
  return names[platform] || platform
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

const refreshSocialData = () => {
  console.log('Refreshing social media data...')
  // منطق به‌روزرسانی داده‌های شبکه‌های اجتماعی
}

const createPost = () => {
  console.log('Creating new post...')
  // منطق ایجاد پست جدید
}

const analyzeCompetitor = (competitorId) => {
  console.log('Analyzing competitor:', competitorId)
  // منطق تحلیل رقیب
}

// آیکون‌های پلتفرم‌ها
const InstagramIcon = defineComponent({
  template: `
    <svg class="w-8 h-8" viewBox="0 0 24 24" fill="currentColor">
      <path d="M12 2.163c3.204 0 3.584.012 4.85.07 3.252.148 4.771 1.691 4.919 4.919.058 1.265.069 1.645.069 4.849 0 3.205-.012 3.584-.069 4.849-.149 3.225-1.664 4.771-4.919 4.919-1.266.058-1.644.07-4.85.07-3.204 0-3.584-.012-4.849-.07-3.26-.149-4.771-1.699-4.919-4.92-.058-1.265-.07-1.644-.07-4.849 0-3.204.013-3.583.07-4.849.149-3.227 1.664-4.771 4.919-4.919 1.266-.057 1.645-.069 4.849-.069zm0-2.163c-3.259 0-3.667.014-4.947.072-4.358.2-6.78 2.618-6.98 6.98-.059 1.281-.073 1.689-.073 4.948 0 3.259.014 3.668.072 4.948.2 4.358 2.618 6.78 6.98 6.98 1.281.058 1.689.072 4.948.072 3.259 0 3.668-.014 4.948-.072 4.354-.2 6.782-2.618 6.979-6.98.059-1.28.073-1.689.073-4.948 0-3.259-.014-3.667-.072-4.947-.196-4.354-2.617-6.78-6.979-6.98-1.281-.059-1.69-.073-4.949-.073zm0 5.838c-3.403 0-6.162 2.759-6.162 6.162s2.759 6.163 6.162 6.163 6.162-2.759 6.162-6.163c0-3.403-2.759-6.162-6.162-6.162zm0 10.162c-2.209 0-4-1.79-4-4 0-2.209 1.791-4 4-4s4 1.791 4 4c0 2.21-1.791 4-4 4zm6.406-11.845c-.796 0-1.441.645-1.441 1.44s.645 1.44 1.441 1.44c.795 0 1.439-.645 1.439-1.44s-.644-1.44-1.439-1.44z"/>
    </svg>
  `
})

const TelegramIcon = defineComponent({
  template: `
    <svg class="w-8 h-8" viewBox="0 0 24 24" fill="currentColor">
      <path d="M11.944 0A12 12 0 0 0 0 12a12 12 0 0 0 12 12 12 12 0 0 0 12-12A12 12 0 0 0 12 0a12 12 0 0 0-.056 0zm4.962 7.224c.1-.002.321.023.465.14a.506.506 0 0 1 .171.325c.016.093.036.306.02.472-.18 1.898-.962 6.502-1.36 8.627-.168.9-.499 1.201-.82 1.23-.696.065-1.225-.46-1.9-.902-1.056-.693-1.653-1.124-2.678-1.8-1.185-.78-.417-1.21.258-1.91.177-.184 3.247-2.977 3.307-3.23.007-.032.014-.15-.056-.212s-.174-.041-.249-.024c-.106.024-1.793 1.14-5.061 3.345-.48.33-.913.49-1.302.48-.428-.008-1.252-.241-1.865-.44-.752-.245-1.349-.374-1.297-.789.027-.216.325-.437.893-.663 3.498-1.524 5.83-2.529 6.998-3.014 3.332-1.386 4.025-1.627 4.476-1.635z"/>
    </svg>
  `
})

const TwitterIcon = defineComponent({
  template: `
    <svg class="w-8 h-8" viewBox="0 0 24 24" fill="currentColor">
      <path d="M23.953 4.57a10 10 0 01-2.825.775 4.958 4.958 0 002.163-2.723c-.951.555-2.005.959-3.127 1.184a4.92 4.92 0 00-8.384 4.482C7.69 8.095 4.067 6.13 1.64 3.162a4.822 4.822 0 00-.666 2.475c0 1.71.87 3.213 2.188 4.096a4.904 4.904 0 01-2.228-.616v.06a4.923 4.923 0 003.946 4.827 4.996 4.996 0 01-2.212.085 4.936 4.936 0 004.604 3.417 9.867 9.867 0 01-6.102 2.105c-.39 0-.779-.023-1.17-.067a13.995 13.995 0 007.557 2.209c9.053 0 13.998-7.496 13.998-13.985 0-.21 0-.42-.015-.63A9.935 9.935 0 0024 4.59z"/>
    </svg>
  `
})

const LinkedInIcon = defineComponent({
  template: `
    <svg class="w-8 h-8" viewBox="0 0 24 24" fill="currentColor">
      <path d="M20.447 20.452h-3.554v-5.569c0-1.328-.027-3.047-1.852-3.047-1.853 0-2.136 1.445-2.136 2.939v5.677H9.351V9h3.414v1.561h.046c.477-.9 1.637-1.85 3.37-1.85 3.601 0 4.267 2.37 4.267 5.455v6.286zM5.337 7.433c-1.144 0-2.063-.926-2.063-2.065 0-1.138.92-2.063 2.063-2.063 1.14 0 2.064.925 2.064 2.063 0 1.139-.925 2.065-2.064 2.065zm1.782 13.019H3.555V9h3.564v11.452zM22.225 0H1.771C.792 0 0 .774 0 1.729v20.542C0 23.227.792 24 1.771 24h20.451C23.2 24 24 23.227 24 22.271V1.729C24 .774 23.2 0 22.222 0h.003z"/>
    </svg>
  `
})
</script>
