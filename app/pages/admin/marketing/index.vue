<template>
  <div class="space-y-6">
    <!-- هدر صفحه -->
    <div class="flex justify-between items-center">
      <h1 class="text-2xl font-bold text-gray-900">بازاریابی و فروش</h1>
      <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
        ایجاد کمپین جدید
      </button>
    </div>

    <!-- آمار کلی -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="bg-white p-6 rounded-lg shadow">
        <div class="flex items-center">
          <div class="p-2 bg-blue-100 rounded-lg">
            <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
            </svg>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">کل فروش</p>
            <p class="text-2xl font-semibold text-gray-900">{{ marketingStats.totalSales }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white p-6 rounded-lg shadow">
        <div class="flex items-center">
          <div class="p-2 bg-green-100 rounded-lg">
            <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">کمپین‌های فعال</p>
            <p class="text-2xl font-semibold text-gray-900">{{ marketingStats.activeCampaigns }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white p-6 rounded-lg shadow">
        <div class="flex items-center">
          <div class="p-2 bg-yellow-100 rounded-lg">
            <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">نرخ تبدیل</p>
            <p class="text-2xl font-semibold text-gray-900">{{ marketingStats.conversionRate }}%</p>
          </div>
        </div>
      </div>

      <div class="bg-white p-6 rounded-lg shadow">
        <div class="flex items-center">
          <div class="p-2 bg-purple-100 rounded-lg">
            <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path>
            </svg>
          </div>
          <div class="mr-4">
            <p class="text-sm font-medium text-gray-600">مشتریان جدید</p>
            <p class="text-2xl font-semibold text-gray-900">{{ marketingStats.newCustomers }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- کارت‌های ابزار -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <!-- کوپن و کمپین -->
      <div class="bg-white rounded-lg shadow p-6">
        <div class="flex items-center mb-4">
          <div class="p-2 bg-red-100 rounded-lg">
            <svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900 mr-3">کوپن و کمپین</h3>
        </div>
        <p class="text-gray-600 mb-4">مدیریت کوپن‌های تخفیف و کمپین‌های بازاریابی</p>
        <NuxtLink to="/admin/marketing/discounts" class="inline-flex items-center text-blue-600 hover:text-blue-800">
          مدیریت کوپن‌ها
          <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
          </svg>
        </NuxtLink>
      </div>

      <!-- بنرها و اسلایدر -->
      <div class="bg-white rounded-lg shadow p-6">
        <div class="flex items-center mb-4">
          <div class="p-2 bg-blue-100 rounded-lg">
            <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900 mr-3">بنرها و اسلایدر</h3>
        </div>
        <p class="text-gray-600 mb-4">مدیریت بنرهای تبلیغاتی و اسلایدرهای سایت</p>
        <NuxtLink to="/admin/content/banners" class="inline-flex items-center text-blue-600 hover:text-blue-800">
          مدیریت بنرها
          <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
          </svg>
        </NuxtLink>
      </div>

      <!-- ترب / ایمالز -->
      <div class="bg-white rounded-lg shadow p-6">
        <div class="flex items-center mb-4">
          <div class="p-2 bg-green-100 rounded-lg">
            <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900 mr-3">ترب / ایمالز</h3>
        </div>
        <p class="text-gray-600 mb-4">اتصال به سرویس‌های قیمت‌گذاری و مقایسه</p>
        <NuxtLink to="/admin/marketing/integrations" class="inline-flex items-center text-blue-600 hover:text-blue-800">
          مدیریت اتصالات
          <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
          </svg>
        </NuxtLink>
      </div>

      <!-- برنامه ارجاع -->
      <div class="bg-white rounded-lg shadow p-6">
        <div class="flex items-center mb-4">
          <div class="p-2 bg-purple-100 rounded-lg">
            <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900 mr-3">برنامه ارجاع</h3>
        </div>
        <p class="text-gray-600 mb-4">مدیریت سیستم ارجاع و معرفی مشتریان</p>
        <NuxtLink to="/admin/marketing/referral" class="inline-flex items-center text-blue-600 hover:text-blue-800">
          مدیریت ارجاع
          <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
          </svg>
        </NuxtLink>
      </div>

      <!-- برنامه وفاداری -->
      <div class="bg-white rounded-lg shadow p-6">
        <div class="flex items-center mb-4">
          <div class="p-2 bg-yellow-100 rounded-lg">
            <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900 mr-3">برنامه وفاداری</h3>
        </div>
        <p class="text-gray-600 mb-4">مدیریت سیستم امتیازدهی و وفاداری مشتریان</p>
        <NuxtLink to="/admin/marketing/loyalty" class="inline-flex items-center text-blue-600 hover:text-blue-800">
          مدیریت وفاداری
          <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
          </svg>
        </NuxtLink>
      </div>


    </div>

    <!-- کمپین‌های فعال -->
    <div class="bg-white rounded-lg shadow overflow-hidden">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">کمپین‌های فعال</h3>
      </div>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                نام کمپین
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                نوع
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                تخفیف
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                تاریخ شروع
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                تاریخ پایان
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                وضعیت
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                عملیات
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="campaign in activeCampaigns" :key="campaign.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm font-medium text-gray-900">{{ campaign.name }}</div>
                <div class="text-sm text-gray-500">{{ campaign.description }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span class="inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-blue-100 text-blue-800">
                  {{ campaign.type }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ campaign.discount }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ formatDate(campaign.startDate) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ formatDate(campaign.endDate) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span
class="inline-flex px-2 py-1 text-xs font-semibold rounded-full" 
                      :class="getStatusClass(campaign.status)">
                  {{ getStatusText(campaign.status) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">ویرایش</button>
                <button class="text-green-600 hover:text-green-900 ml-3">مشاهده</button>
                <button class="text-red-600 hover:text-red-900">غیرفعال</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup>
// آمار بازاریابی
definePageMeta({ layout: 'admin-main', middleware: 'admin' })
const marketingStats = ref({
  totalSales: '125,000,000 تومان',
  activeCampaigns: 8,
  conversionRate: 3.2,
  newCustomers: 156
})

// کمپین‌های فعال
const activeCampaigns = ref([
  {
    id: 1,
    name: 'تخفیف تابستانه',
    description: 'تخفیف ویژه محصولات تابستانه',
    type: 'تخفیف درصدی',
    discount: '20%',
    startDate: '2024-06-01',
    endDate: '2024-08-31',
    status: 'active'
  },
  {
    id: 2,
    name: 'ارسال رایگان',
    description: 'ارسال رایگان برای خریدهای بالای 500 هزار تومان',
    type: 'ارسال رایگان',
    discount: 'رایگان',
    startDate: '2024-07-01',
    endDate: '2024-07-31',
    status: 'active'
  },
  {
    id: 3,
    name: 'کد تخفیف جدید',
    description: 'کد تخفیف برای مشتریان جدید',
    type: 'کد تخفیف',
    discount: '50,000 تومان',
    startDate: '2024-07-15',
    endDate: '2024-08-15',
    status: 'active'
  }
])

// تابع فرمت کردن تاریخ
const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString('fa-IR')
}

// تابع کلاس‌های وضعیت
const getStatusClass = (status) => {
  const classes = {
    active: 'bg-green-100 text-green-800',
    paused: 'bg-yellow-100 text-yellow-800',
    ended: 'bg-red-100 text-red-800'
  }
  return classes[status] || classes.ended
}

// تابع متن وضعیت
const getStatusText = (status) => {
  const texts = {
    active: 'فعال',
    paused: 'متوقف شده',
    ended: 'پایان یافته'
  }
  return texts[status] || 'نامشخص'
}
</script> 
