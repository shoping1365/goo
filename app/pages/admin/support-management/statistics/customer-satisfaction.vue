<template>
  <div class="p-6" dir="rtl">
    <div class="mb-6 bg-white p-6 rounded-lg shadow-md border border-gray-200">
      <h1 class="text-2xl font-bold text-gray-900 mb-2">آمار رضایت مشتری</h1>
      <p class="text-gray-600">تحلیل و گزارش‌گیری از رضایت مشتریان</p>
    </div>

    <!-- آمار کلی -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
      <TemplateCard
        title="میانگین رضایت"
        :value="avgSatisfaction"
        variant="green"
      />
      <TemplateCard
        title="کل نظرات"
        :value="totalReviews"
        variant="blue"
      />
      <TemplateCard
        title="رضایت بالا (4-5 ستاره)"
        :value="highSatisfaction"
        variant="cyan"
      />
      <TemplateCard
        title="رضایت پایین (1-2 ستاره)"
        :value="lowSatisfaction"
        variant="red"
      />
    </div>

    <!-- کارت‌های آمار جدید رضایت مشتری -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-6 mb-8">
      <!-- درصد مشتریان خوشحال -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <div class="flex items-center space-x-2 space-x-reverse">
            <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
            </svg>
            <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
            </svg>
          </div>
        </div>
        <h3 class="text-lg font-semibold text-gray-900 mb-2">درصد مشتریان خوشحال (😊)</h3>
        <div class="text-2xl font-bold text-green-600 mb-2">درصد مشتریان : ۰</div>
        <div class="text-sm text-gray-500">نرخ رشد : ۰٪</div>
      </div>

      <!-- درصد رضایت مشتریان -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <div class="flex items-center space-x-2 space-x-reverse">
            <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
            </svg>
            <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
            </svg>
          </div>
        </div>
        <h3 class="text-lg font-semibold text-gray-900 mb-2">درصد رضایت مشتریان</h3>
        <div class="text-2xl font-bold text-blue-600 mb-2">درصد رضایت : ۰</div>
        <div class="text-sm text-gray-500">نرخ رشد : ۰٪</div>
      </div>

      <!-- تعداد امتیازهای ثبت شده -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <div class="flex items-center space-x-2 space-x-reverse">
            <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
            </svg>
            <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
            </svg>
          </div>
        </div>
        <h3 class="text-lg font-semibold text-gray-900 mb-2">تعداد امتیازهای ثبت شده</h3>
        <div class="text-2xl font-bold text-purple-600 mb-2">تعداد امتیازهای ثبت شده : ۰</div>
        <div class="text-sm text-gray-500">نرخ رشد : ۰٪</div>
      </div>

      <!-- درصد مشتریان عصبانی -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <div class="flex items-center space-x-2 space-x-reverse">
            <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
            </svg>
            <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
            </svg>
          </div>
        </div>
        <h3 class="text-lg font-semibold text-gray-900 mb-2">درصد مشتریان عصبانی (😡)</h3>
        <div class="text-2xl font-bold text-red-600 mb-2">درصد مشتریان : ۰</div>
        <div class="text-sm text-gray-500">نرخ رشد : ۰٪</div>
      </div>

      <!-- میانگین امتیازها -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-4">
          <div class="flex items-center space-x-2 space-x-reverse">
            <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
            </svg>
            <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
            </svg>
          </div>
        </div>
        <h3 class="text-lg font-semibold text-gray-900 mb-2">میانگین امتیازها</h3>
        <div class="text-2xl font-bold text-orange-600 mb-2">میانگین امتیازها ( از ۱ تا ۴ ) : ۰</div>
        <div class="text-sm text-gray-500">نرخ رشد : ۰٪</div>
      </div>
    </div>

    <!-- انتخاب بازه زمانی -->
    <div class="bg-gray-100 rounded-lg p-6 mb-8">
      <div class="text-sm text-gray-600">۱۴۰۴/۰۵/۰۱ - ۱۴۰۴/۰۴/۲۶</div>
    </div>

    <!-- امتیاز به تفکیک اپراتورها -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 mb-8">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">امتیاز به تفکیک اپراتورها</h3>
      </div>
      <div class="p-6">
        <div class="flex items-start space-x-6 space-x-reverse">
          <!-- لیست امتیازها -->
          <div class="flex flex-col space-y-4 space-y-reverse">
            <div class="flex items-center">
              <div class="w-3 h-3 bg-red-500 rounded mr-3"></div>
              <span class="text-sm text-gray-700">امتیاز ۱</span>
            </div>
            <div class="flex items-center">
              <div class="w-3 h-3 bg-red-500 rounded mr-3"></div>
              <span class="text-sm text-gray-700">امتیاز ۲</span>
            </div>
            <div class="flex items-center">
              <div class="w-3 h-3 bg-red-500 rounded mr-3"></div>
              <span class="text-sm text-gray-700">امتیاز ۳</span>
            </div>
            <div class="flex items-center">
              <div class="w-3 h-3 bg-red-500 rounded mr-3"></div>
              <span class="text-sm text-gray-700">امتیاز ۴</span>
            </div>
          </div>
          
          <!-- جدول امتیازات -->
          <div class="flex-1">
            <div class="border border-gray-200 rounded-lg p-8">
              <div class="text-center text-gray-500">
                <p class="text-lg">داده ای برای نمایش وجود ندارد !</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- فیلترها -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 mb-8">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-semibold text-gray-900">فیلترها و جستجو</h3>
      </div>
      <div class="p-6">
        <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">دوره زمانی</label>
            <select class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm">
              <option value="today">امروز</option>
              <option value="week">هفته گذشته</option>
              <option value="month">ماه گذشته</option>
              <option value="quarter">سه ماهه</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">سطح رضایت</label>
            <select class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm">
              <option value="">همه</option>
              <option value="5">5 ستاره</option>
              <option value="4">4 ستاره</option>
              <option value="3">3 ستاره</option>
              <option value="2">2 ستاره</option>
              <option value="1">1 ستاره</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع خدمات</label>
            <select class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm">
              <option value="">همه</option>
              <option value="chat">چت زنده</option>
              <option value="ticket">تیکت</option>
              <option value="call">تماس تلفنی</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">اپراتور</label>
            <select class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm">
              <option value="">همه</option>
              <option value="operator1">علی احمدی</option>
              <option value="operator2">فاطمه محمدی</option>
              <option value="operator3">محمد رضایی</option>
            </select>
          </div>
        </div>
      </div>
    </div>

    <!-- نمودارها -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8 mb-8">
      <!-- نمودار توزیع رضایت -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200">
        <div class="px-6 py-4 border-b border-gray-200">
          <h3 class="text-lg font-semibold text-gray-900">توزیع رضایت مشتریان</h3>
        </div>
        <div class="p-6">
          <div class="h-64 bg-gray-50 rounded-lg flex items-center justify-center">
            <div class="text-center">
              <svg class="w-16 h-16 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 3.055A9.001 9.001 0 1020.945 13H11V3.055z"></path>
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.488 9H15V3.512A9.025 9.025 0 0120.488 9z"></path>
              </svg>
              <p class="text-gray-500">نمودار دایره‌ای توزیع رضایت</p>
              <p class="text-sm text-gray-400">بر اساس تعداد ستاره‌ها</p>
            </div>
          </div>
        </div>
      </div>

      <!-- نمودار روند رضایت در طول زمان -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200">
        <div class="px-6 py-4 border-b border-gray-200">
          <h3 class="text-lg font-semibold text-gray-900">روند رضایت در طول زمان</h3>
        </div>
        <div class="p-6">
          <div class="h-64 bg-gray-50 rounded-lg flex items-center justify-center">
            <div class="text-center">
              <svg class="w-16 h-16 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
              </svg>
              <p class="text-gray-500">نمودار روند رضایت</p>
              <p class="text-sm text-gray-400">آخرین 30 روز</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- آمار تفصیلی -->
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-8 mb-8">
      <!-- توزیع ستاره‌ها -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200">
        <div class="px-6 py-4 border-b border-gray-200">
          <h3 class="text-lg font-semibold text-gray-900">توزیع ستاره‌ها</h3>
        </div>
        <div class="p-6">
          <div class="space-y-4">
            <div v-for="rating in ratingDistribution" :key="rating.stars" class="flex items-center justify-between">
              <div class="flex items-center">
                <div class="flex text-yellow-400 mr-3">
                  <svg v-for="i in 5" :key="i" class="w-4 h-4" :class="i <= rating.stars ? 'text-yellow-400' : 'text-gray-300'" fill="currentColor" viewBox="0 0 20 20">
                    <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"></path>
                  </svg>
                </div>
                <span class="text-sm text-gray-900">{{ rating.stars }} ستاره</span>
              </div>
              <div class="text-right">
                <div class="text-sm font-medium text-gray-900">{{ rating.count }}</div>
                <div class="text-xs text-gray-500">{{ rating.percentage }}%</div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- رضایت بر اساس نوع خدمات -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200">
        <div class="px-6 py-4 border-b border-gray-200">
          <h3 class="text-lg font-semibold text-gray-900">رضایت بر اساس نوع خدمات</h3>
        </div>
        <div class="p-6">
          <div class="space-y-4">
            <div v-for="service in serviceSatisfaction" :key="service.type" class="flex items-center justify-between">
              <div class="flex items-center">
                <div :class="service.color" class="w-3 h-3 rounded-full mr-3"></div>
                <span class="text-sm text-gray-900">{{ service.name }}</span>
              </div>
              <div class="text-right">
                <div class="text-sm font-medium text-gray-900">{{ service.satisfaction }}</div>
                <div class="text-xs text-gray-500">{{ service.count }} نظر</div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- بهترین اپراتورها -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200">
        <div class="px-6 py-4 border-b border-gray-200">
          <h3 class="text-lg font-semibold text-gray-900">بهترین اپراتورها</h3>
        </div>
        <div class="p-6">
          <div class="space-y-4">
            <div v-for="operator in topOperators" :key="operator.id" class="flex items-center justify-between">
              <div class="flex items-center">
                <div class="w-8 h-8 rounded-full bg-gray-300 flex items-center justify-center mr-3">
                  <span class="text-xs font-medium text-gray-700">{{ operator.name.charAt(0) }}</span>
                </div>
                <div>
                  <div class="text-sm font-medium text-gray-900">{{ operator.name }}</div>
                  <div class="text-xs text-gray-500">{{ operator.satisfaction }} ستاره</div>
                </div>
              </div>
              <div class="text-right">
                <div class="text-sm font-medium text-gray-900">{{ operator.reviews }}</div>
                <div class="text-xs text-gray-500">نظر</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- جدول نظرات -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200 flex justify-between items-center">
        <h3 class="text-lg font-semibold text-gray-900">لیست نظرات</h3>
        <button class="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600 transition-colors">
          دانلود گزارش
        </button>
      </div>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مشتری</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">اپراتور</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع خدمات</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">رضایت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نظر</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="review in reviews" :key="review.id">
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{{ review.customer }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ review.operator }}</td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getServiceClass(review.serviceType)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ getServiceText(review.serviceType) }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                <div class="flex items-center">
                  <span class="text-sm font-medium text-gray-900 mr-2">{{ review.rating }}</span>
                  <div class="flex text-yellow-400">
                    <svg v-for="i in 5" :key="i" class="w-4 h-4" :class="i <= review.rating ? 'text-yellow-400' : 'text-gray-300'" fill="currentColor" viewBox="0 0 20 20">
                      <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"></path>
                    </svg>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 text-sm text-gray-900">
                <div class="max-w-xs truncate">{{ review.comment }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ review.date }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <div class="flex space-x-2 space-x-reverse">
                  <button class="text-blue-600 hover:text-blue-900 transition-colors" title="مشاهده جزئیات">
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                    </svg>
                  </button>
                  <button class="text-green-600 hover:text-green-900 transition-colors" title="پاسخ">
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h10a8 8 0 018 8v2M3 10l6 6m-6-6l6-6"></path>
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup>
import TemplateCard from '@/components/common/TemplateCard.vue'

definePageMeta({ layout: 'admin-main' })

// داده‌های نمونه
const avgSatisfaction = ref('4.7/5')
const totalReviews = ref(1245)
const highSatisfaction = ref(892)
const lowSatisfaction = ref(45)

const ratingDistribution = ref([
  { stars: 5, count: 456, percentage: 37 },
  { stars: 4, count: 436, percentage: 35 },
  { stars: 3, count: 234, percentage: 19 },
  { stars: 2, count: 78, percentage: 6 },
  { stars: 1, count: 41, percentage: 3 }
])

const serviceSatisfaction = ref([
  { type: 'chat', name: 'چت زنده', satisfaction: '4.8', count: 456, color: 'bg-blue-500' },
  { type: 'ticket', name: 'تیکت', satisfaction: '4.6', count: 389, color: 'bg-green-500' },
  { type: 'call', name: 'تماس تلفنی', satisfaction: '4.7', count: 400, color: 'bg-purple-500' }
])

const topOperators = ref([
  { id: 1, name: 'علی احمدی', satisfaction: 4.9, reviews: 156 },
  { id: 2, name: 'فاطمه محمدی', satisfaction: 4.8, reviews: 134 },
  { id: 3, name: 'حسن نوری', satisfaction: 4.7, reviews: 98 },
  { id: 4, name: 'محمد رضایی', satisfaction: 4.6, reviews: 87 },
  { id: 5, name: 'سارا کریمی', satisfaction: 4.5, reviews: 76 }
])

const reviews = ref([
  {
    id: 1,
    customer: 'احمد محمدی',
    operator: 'علی احمدی',
    serviceType: 'chat',
    rating: 5,
    comment: 'خدمات عالی بود. اپراتور خیلی سریع مشکل من رو حل کرد.',
    date: '1402/10/15'
  },
  {
    id: 2,
    customer: 'فاطمه کریمی',
    operator: 'فاطمه محمدی',
    serviceType: 'ticket',
    rating: 4,
    comment: 'خوب بود ولی کمی طول کشید تا پاسخ دادند.',
    date: '1402/10/14'
  },
  {
    id: 3,
    customer: 'محمد رضایی',
    operator: 'حسن نوری',
    serviceType: 'call',
    rating: 5,
    comment: 'مکالمه خیلی مفید بود. مشکل کاملاً حل شد.',
    date: '1402/10/13'
  },
  {
    id: 4,
    customer: 'سارا احمدی',
    operator: 'محمد رضایی',
    serviceType: 'chat',
    rating: 3,
    comment: 'متوسط بود. انتظار بیشتری داشتم.',
    date: '1402/10/12'
  },
  {
    id: 5,
    customer: 'علی نوری',
    operator: 'سارا کریمی',
    serviceType: 'ticket',
    rating: 5,
    comment: 'عالی! خیلی راضی هستم از خدمات.',
    date: '1402/10/11'
  }
])

function getServiceClass(serviceType) {
  switch (serviceType) {
    case 'chat':
      return 'bg-blue-100 text-blue-800'
    case 'ticket':
      return 'bg-green-100 text-green-800'
    case 'call':
      return 'bg-purple-100 text-purple-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

function getServiceText(serviceType) {
  switch (serviceType) {
    case 'chat':
      return 'چت زنده'
    case 'ticket':
      return 'تیکت'
    case 'call':
      return 'تماس تلفنی'
    default:
      return 'نامشخص'
  }
}
</script> 
