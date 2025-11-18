<template>
  <div class="space-y-8">
    <!-- بخش هوش مصنوعی -->
    <section>
      <div class="bg-gradient-to-r from-blue-50 to-indigo-50 rounded-lg p-6 mb-6">
        <h2 class="text-2xl font-bold text-gray-900 mb-2">🤖 هوش مصنوعی</h2>
        <p class="text-gray-600">استفاده از الگوریتم‌های هوش مصنوعی برای بهینه‌سازی و بهبود تجربه کاربری</p>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div v-for="feature in aiFeatures" :key="feature.title" class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-start space-x-4 space-x-reverse">
            <div class="flex-shrink-0">
              <div class="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center">
                <span class="text-2xl">{{ feature.icon }}</span>
              </div>
            </div>
            <div class="flex-1">
              <h3 class="text-lg font-semibold text-gray-900 mb-2">{{ feature.title }}</h3>
              <p class="text-gray-600 mb-4">{{ feature.description }}</p>
              
              <!-- وضعیت و آمار -->
              <div class="space-y-3">
                <div class="flex items-center justify-between">
                  <span class="text-sm text-gray-500">وضعیت:</span>
                  <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">فعال</span>
                </div>
                
                <div class="flex items-center justify-between">
                  <span class="text-sm text-gray-500">دقت:</span>
                  <span class="text-sm font-medium text-gray-900">{{ feature.accuracy }}%</span>
                </div>
                
                <div class="flex items-center justify-between">
                  <span class="text-sm text-gray-500">استفاده امروز:</span>
                  <span class="text-sm font-medium text-gray-900">{{ feature.usageToday }}</span>
                </div>
              </div>
              
              <!-- دکمه‌های عملیاتی -->
              <div class="mt-4 flex space-x-2 space-x-reverse">
                <button class="px-4 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
                  تنظیمات
                </button>
                <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
                  گزارش
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- بخش گیمیفیکیشن -->
    <section>
      <div class="bg-gradient-to-r from-purple-50 to-pink-50 rounded-lg p-6 mb-6">
        <h2 class="text-2xl font-bold text-gray-900 mb-2">🎮 گیمیفیکیشن</h2>
        <p class="text-gray-600">افزایش تعامل کاربران از طریق سیستم‌های امتیازدهی، چالش‌ها و جوایز</p>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div v-for="feature in gamificationFeatures" :key="feature.title" class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-start space-x-4 space-x-reverse">
            <div class="flex-shrink-0">
              <div class="w-12 h-12 bg-purple-100 rounded-lg flex items-center justify-center">
                <span class="text-2xl">{{ feature.icon }}</span>
              </div>
            </div>
            <div class="flex-1">
              <h3 class="text-lg font-semibold text-gray-900 mb-2">{{ feature.title }}</h3>
              <p class="text-gray-600 mb-4">{{ feature.description }}</p>
              
              <!-- آمار و اطلاعات -->
              <div class="space-y-3">
                <div class="flex items-center justify-between">
                  <span class="text-sm text-gray-500">شرکت‌کنندگان:</span>
                  <span class="text-sm font-medium text-gray-900">{{ feature.participants }}</span>
                </div>
                
                <div class="flex items-center justify-between">
                  <span class="text-sm text-gray-500">جوایز توزیع شده:</span>
                  <span class="text-sm font-medium text-gray-900">{{ feature.rewardsDistributed }}</span>
                </div>
                
                <div class="flex items-center justify-between">
                  <span class="text-sm text-gray-500">وضعیت:</span>
                  <span class="text-xs bg-green-100 text-green-800 rounded-full px-3 py-1">{{ feature.status }}</span>
                </div>
              </div>
              
              <!-- نوار پیشرفت -->
              <div class="mt-4">
                <div class="flex justify-between text-sm text-gray-500 mb-1">
                  <span>پیشرفت کلی</span>
                  <span>{{ feature.progress }}%</span>
                </div>
                <div class="w-full bg-gray-200 rounded-full h-2">
                  <div class="bg-gradient-to-r from-purple-500 to-pink-500 h-2 rounded-full" :style="{ width: feature.progress + '%' }"></div>
                </div>
              </div>
              
              <!-- دکمه‌های عملیاتی -->
              <div class="mt-4 flex space-x-2 space-x-reverse">
                <button class="px-4 py-2 bg-purple-600 text-white text-sm rounded-lg hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-purple-500">
                  مدیریت
                </button>
                <button class="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-lg hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-500">
                  آمار
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- داشبورد کلی -->
    <section>
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h3 class="text-xl font-semibold text-gray-900 mb-4">📊 آمار کلی ویژگی‌های پیشرفته</h3>
        
        <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
          <div class="text-center p-6 bg-blue-50 rounded-lg">
            <div class="text-2xl font-bold text-blue-600">{{ overallStats.aiUsage }}</div>
            <div class="text-sm text-gray-600">استفاده از AI</div>
          </div>
          
          <div class="text-center p-6 bg-purple-50 rounded-lg">
            <div class="text-2xl font-bold text-purple-600">{{ overallStats.activeUsers }}</div>
            <div class="text-sm text-gray-600">کاربران فعال</div>
          </div>
          
          <div class="text-center p-6 bg-green-50 rounded-lg">
            <div class="text-2xl font-bold text-green-600">{{ overallStats.rewardsGiven }}</div>
            <div class="text-sm text-gray-600">جوایز اهدا شده</div>
          </div>
          
          <div class="text-center p-6 bg-yellow-50 rounded-lg">
            <div class="text-2xl font-bold text-yellow-600">{{ overallStats.accuracy }}%</div>
            <div class="text-sm text-gray-600">دقت کلی</div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
// ویژگی‌های هوش مصنوعی
const aiFeatures = [
  {
    title: 'پیشنهاد مبلغ بهینه',
    description: 'تحلیل رفتار کاربران و پیشنهاد مبالغ مناسب بر اساس الگوهای خرید',
    icon: '💰',
    accuracy: 94,
    usageToday: '1,247 بار'
  },
  {
    title: 'تشخیص تقلب',
    description: 'شناسایی فعالیت‌های مشکوک و جلوگیری از سوء استفاده از گیفت کارت‌ها',
    icon: '🛡️',
    accuracy: 98,
    usageToday: '3,891 بار'
  },
  {
    title: 'پیش‌بینی تقاضا',
    description: 'پیش‌بینی نیاز به گیفت کارت در مناسبت‌های مختلف و مدیریت موجودی',
    icon: '📈',
    accuracy: 87,
    usageToday: '156 بار'
  },
  {
    title: 'شخصی‌سازی پیشنهادات',
    description: 'ارائه پیشنهادات شخصی‌سازی شده بر اساس سوابق و ترجیحات کاربر',
    icon: '🎯',
    accuracy: 91,
    usageToday: '2,134 بار'
  }
]

// ویژگی‌های گیمیفیکیشن
const gamificationFeatures = [
  {
    title: 'سیستم امتیازدهی',
    description: 'امتیازدهی به کاربران بر اساس فعالیت‌های مختلف و خریدهای انجام شده',
    icon: '⭐',
    participants: '12,847',
    rewardsDistributed: '8,234',
    status: 'فعال',
    progress: 85
  },
  {
    title: 'چالش‌های خرید',
    description: 'چالش‌های ماهانه و فصلی برای تشویق کاربران به خرید بیشتر',
    icon: '🏆',
    participants: '5,432',
    rewardsDistributed: '2,156',
    status: 'فعال',
    progress: 72
  },
  {
    title: 'رتبه‌بندی کاربران',
    description: 'سیستم رتبه‌بندی کاربران بر اساس امتیازات و فعالیت‌های انجام شده',
    icon: '👑',
    participants: '15,234',
    rewardsDistributed: '12,789',
    status: 'فعال',
    progress: 93
  },
  {
    title: 'جوایز ویژه',
    description: 'جوایز ویژه برای کاربران برتر و مناسبت‌های خاص',
    icon: '🎁',
    participants: '8,765',
    rewardsDistributed: '4,321',
    status: 'فعال',
    progress: 68
  }
]

// آمار کلی
const overallStats = {
  aiUsage: '7,428',
  activeUsers: '15,234',
  rewardsGiven: '27,500',
  accuracy: 92
}
</script>

<!--
  مستندسازی:
  این کامپوننت شامل دو بخش اصلی است:
  1. ویژگی‌های هوش مصنوعی: پیشنهاد مبلغ، تشخیص تقلب، پیش‌بینی تقاضا، شخصی‌سازی
  2. گیمیفیکیشن: امتیازدهی، چالش‌ها، رتبه‌بندی، جوایز
  
  تمام توضیحات به فارسی و با طراحی ریسپانسیو ارائه شده است.
  آمار و اعداد نمونه برای نمایش قابلیت‌ها استفاده شده‌اند.
--> 
