<template>
  <div class="space-y-8">
    <!-- هدر بخش -->
    <div class="flex justify-between items-center">
      <div>
        <h2 class="text-2xl font-bold text-gray-900">شخصی‌سازی پیشنهادات</h2>
        <p class="mt-1 text-sm text-gray-600">تنظیم و مدیریت شخصی‌سازی پیشنهادات برای کاربران</p>
      </div>
      <button class="bg-pink-600 text-white px-4 py-2 rounded-lg hover:bg-pink-700 transition-colors">
        <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
        </svg>
        ایجاد پروفایل جدید
      </button>
    </div>

    <!-- آمار شخصی‌سازی -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gapx-4 py-4">
      <div class="bg-gradient-to-r from-pink-500 to-pink-600 rounded-xl px-4 py-4 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">کاربران شخصی‌سازی شده</p>
            <p class="text-2xl font-bold">8,234</p>
            <p class="text-xs opacity-75">از 12,847 کاربر</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-blue-500 to-blue-600 rounded-xl px-4 py-4 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">میانگین رضایت</p>
            <p class="text-2xl font-bold">4.7/5</p>
            <p class="text-xs opacity-75">+0.3 از ماه گذشته</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-green-500 to-green-600 rounded-xl px-4 py-4 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">نرخ تعامل</p>
            <p class="text-2xl font-bold">34.2%</p>
            <p class="text-xs opacity-75">+5.8% از ماه گذشته</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-purple-500 to-purple-600 rounded-xl px-4 py-4 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">پروفایل‌های فعال</p>
            <p class="text-2xl font-bold">156</p>
            <p class="text-xs opacity-75">+12 از هفته گذشته</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیم شدت پیشنهادات -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 px-4 py-4">
      <div class="flex items-center justify-between mb-6">
        <div class="flex items-center">
          <div class="p-2 bg-pink-100 rounded-lg">
            <svg class="w-6 h-6 text-pink-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900 mr-3">تنظیم شدت پیشنهادات</h3>
        </div>
        <label class="relative inline-flex items-center cursor-pointer">
          <input v-model="settings.intensity.enabled" type="checkbox" class="sr-only peer">
          <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-pink-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-pink-600"></div>
        </label>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
        <div class="space-y-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تعداد پیشنهادات در صفحه اصلی</label>
            <input v-model="settings.intensity.homePage" type="range" min="1" max="20" class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
            <div class="flex justify-between text-xs text-gray-500 mt-1">
              <span>1</span>
              <span>{{ settings.intensity.homePage }}</span>
              <span>20</span>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تعداد پیشنهادات در صفحه محصول</label>
            <input v-model="settings.intensity.productPage" type="range" min="1" max="15" class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
            <div class="flex justify-between text-xs text-gray-500 mt-1">
              <span>1</span>
              <span>{{ settings.intensity.productPage }}</span>
              <span>15</span>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تعداد پیشنهادات در سبد خرید</label>
            <input v-model="settings.intensity.cartPage" type="range" min="1" max="10" class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
            <div class="flex justify-between text-xs text-gray-500 mt-1">
              <span>1</span>
              <span>{{ settings.intensity.cartPage }}</span>
              <span>10</span>
            </div>
          </div>
        </div>
        
        <div class="space-y-6">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">فراوانی نمایش</label>
            <select v-model="settings.intensity.frequency" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-pink-500 focus:border-pink-500">
              <option value="always">همیشه</option>
              <option value="sometimes">گاهی اوقات</option>
              <option value="rarely">به ندرت</option>
              <option value="smart">هوشمند</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداقل فاصله بین پیشنهادات</label>
            <input v-model="settings.intensity.minInterval" type="range" min="1" max="60" class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
            <div class="flex justify-between text-xs text-gray-500 mt-1">
              <span>1 دقیقه</span>
              <span>{{ settings.intensity.minInterval }} دقیقه</span>
              <span>60 دقیقه</span>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر پیشنهادات در روز</label>
            <input v-model="settings.intensity.maxDaily" type="range" min="5" max="50" class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
            <div class="flex justify-between text-xs text-gray-500 mt-1">
              <span>5</span>
              <span>{{ settings.intensity.maxDaily }}</span>
              <span>50</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- فیلترهای شخصی -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <!-- فیلتر دسته‌بندی -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 px-4 py-4">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-blue-100 rounded-lg">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">فیلتر دسته‌بندی</h3>
          </div>
          <label class="relative inline-flex items-center cursor-pointer">
            <input v-model="settings.categoryFilter.enabled" type="checkbox" class="sr-only peer">
            <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
          </label>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">دسته‌بندی‌های حذف شده</label>
            <div class="space-y-2 max-h-32 overflow-y-auto">
              <div v-for="category in settings.categoryFilter.excludedCategories" :key="category.id" class="flex items-center justify-between p-2 bg-gray-50 rounded-lg">
                <span class="text-sm text-gray-700">{{ category.name }}</span>
                <button class="text-red-600 hover:text-red-800" @click="removeCategory(category.id)">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                  </svg>
                </button>
              </div>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">افزودن دسته‌بندی جدید</label>
            <div class="flex space-x-2 space-x-reverse">
              <select v-model="newCategory" class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                <option value="">انتخاب دسته‌بندی</option>
                <option v-for="cat in availableCategories" :key="cat.id" :value="cat.id">{{ cat.name }}</option>
              </select>
              <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700" @click="addCategory">افزودن</button>
            </div>
          </div>
        </div>
      </div>

      <!-- فیلتر برند -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 px-4 py-4">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-green-100 rounded-lg">
              <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">فیلتر برند</h3>
          </div>
          <label class="relative inline-flex items-center cursor-pointer">
            <input v-model="settings.brandFilter.enabled" type="checkbox" class="sr-only peer">
            <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-green-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-green-600"></div>
          </label>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">برندهای حذف شده</label>
            <div class="space-y-2 max-h-32 overflow-y-auto">
              <div v-for="brand in settings.brandFilter.excludedBrands" :key="brand.id" class="flex items-center justify-between p-2 bg-gray-50 rounded-lg">
                <span class="text-sm text-gray-700">{{ brand.name }}</span>
                <button class="text-red-600 hover:text-red-800" @click="removeBrand(brand.id)">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                  </svg>
                </button>
              </div>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">افزودن برند جدید</label>
            <div class="flex space-x-2 space-x-reverse">
              <select v-model="newBrand" class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500">
                <option value="">انتخاب برند</option>
                <option v-for="brand in availableBrands" :key="brand.id" :value="brand.id">{{ brand.name }}</option>
              </select>
              <button class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700" @click="addBrand">افزودن</button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- اولویت‌بندی -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 px-4 py-4">
      <div class="flex items-center justify-between mb-6">
        <div class="flex items-center">
          <div class="p-2 bg-purple-100 rounded-lg">
            <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 4V2a1 1 0 011-1h8a1 1 0 011 1v2m-9 0h10m-9 0a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V6a2 2 0 00-2-2m-9 0h10"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900 mr-3">اولویت‌بندی پیشنهادات</h3>
        </div>
        <label class="relative inline-flex items-center cursor-pointer">
          <input v-model="settings.priority.enabled" type="checkbox" class="sr-only peer">
          <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-purple-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-purple-600"></div>
        </label>
      </div>
      
      <div class="space-y-6">
        <div v-for="(priority, index) in settings.priority.types" :key="priority.id" class="flex items-center justify-between px-4 py-4 border border-gray-200 rounded-lg">
          <div class="flex items-center">
            <div class="w-8 h-8 bg-purple-100 rounded-lg flex items-center justify-center mr-4">
              <span class="text-purple-600 font-bold">{{ index + 1 }}</span>
            </div>
            <div>
              <h4 class="text-sm font-medium text-gray-900">{{ priority.name }}</h4>
              <p class="text-xs text-gray-500">{{ priority.description }}</p>
            </div>
          </div>
          <div class="flex items-center space-x-4 space-x-reverse">
            <input v-model="priority.weight" type="range" min="0" max="100" class="w-32 h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
            <span class="text-sm font-medium text-gray-900 w-12">{{ priority.weight }}%</span>
            <div class="flex space-x-1 space-x-reverse">
              <button :disabled="index === 0" class="p-1 text-gray-400 hover:text-gray-600 disabled:opacity-50" @click="movePriority(index, 'up')">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 15l7-7 7 7"></path>
                </svg>
              </button>
              <button :disabled="index === settings.priority.types.length - 1" class="p-1 text-gray-400 hover:text-gray-600 disabled:opacity-50" @click="movePriority(index, 'down')">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
                </svg>
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- زبان و فرهنگ -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 px-4 py-4">
      <div class="flex items-center justify-between mb-6">
        <div class="flex items-center">
          <div class="p-2 bg-yellow-100 rounded-lg">
            <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5h12M9 3v2m1.048 9.5A18.022 18.022 0 016.412 9m6.088 9h7M11 21l5-10 5 10M12.751 5C11.783 10.77 8.07 15.61 3 18.129"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900 mr-3">زبان و فرهنگ</h3>
        </div>
        <label class="relative inline-flex items-center cursor-pointer">
          <input v-model="settings.culture.enabled" type="checkbox" class="sr-only peer">
          <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-yellow-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-yellow-600"></div>
        </label>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">زبان پیشنهادات</label>
            <select v-model="settings.culture.language" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500">
              <option value="fa">فارسی</option>
              <option value="en">انگلیسی</option>
              <option value="ar">عربی</option>
              <option value="auto">خودکار</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">منطقه جغرافیایی</label>
            <select v-model="settings.culture.region" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500">
              <option value="iran">ایران</option>
              <option value="middle_east">خاورمیانه</option>
              <option value="asia">آسیا</option>
              <option value="global">جهانی</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">مناسبت‌های فرهنگی</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input v-model="settings.culture.culturalEvents" type="checkbox" value="nowruz" class="rounded border-gray-300 text-yellow-600 focus:ring-yellow-500">
                <span class="mr-2 text-sm text-gray-700">نوروز</span>
              </label>
              <label class="flex items-center">
                <input v-model="settings.culture.culturalEvents" type="checkbox" value="ramadan" class="rounded border-gray-300 text-yellow-600 focus:ring-yellow-500">
                <span class="mr-2 text-sm text-gray-700">رمضان</span>
              </label>
              <label class="flex items-center">
                <input v-model="settings.culture.culturalEvents" type="checkbox" value="christmas" class="rounded border-gray-300 text-yellow-600 focus:ring-yellow-500">
                <span class="mr-2 text-sm text-gray-700">کریسمس</span>
              </label>
              <label class="flex items-center">
                <input v-model="settings.culture.culturalEvents" type="checkbox" value="valentine" class="rounded border-gray-300 text-yellow-600 focus:ring-yellow-500">
                <span class="mr-2 text-sm text-gray-700">ولنتاین</span>
              </label>
            </div>
          </div>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">سبک پیشنهادات</label>
            <select v-model="settings.culture.style" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500">
              <option value="formal">رسمی</option>
              <option value="casual">غیررسمی</option>
              <option value="friendly">دوستانه</option>
              <option value="professional">حرفه‌ای</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تنوع فرهنگی</label>
            <input v-model="settings.culture.diversity" type="range" min="0" max="100" class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
            <div class="flex justify-between text-xs text-gray-500 mt-1">
              <span>کم</span>
              <span>{{ settings.culture.diversity }}%</span>
              <span>زیاد</span>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">حساسیت فرهنگی</label>
            <select v-model="settings.culture.sensitivity" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500">
              <option value="low">کم</option>
              <option value="medium">متوسط</option>
              <option value="high">زیاد</option>
              <option value="strict">سختگیرانه</option>
            </select>
          </div>
        </div>
      </div>
    </div>

    <!-- سطح کاربر -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 px-4 py-4">
      <div class="flex items-center justify-between mb-6">
        <div class="flex items-center">
          <div class="p-2 bg-indigo-100 rounded-lg">
            <svg class="w-6 h-6 text-indigo-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900 mr-3">سطح کاربر</h3>
        </div>
        <label class="relative inline-flex items-center cursor-pointer">
          <input v-model="settings.userLevel.enabled" type="checkbox" class="sr-only peer">
          <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-indigo-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-indigo-600"></div>
        </label>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-3 gapx-4 py-4">
        <div v-for="level in settings.userLevel.levels" :key="level.id" class="border border-gray-200 rounded-lg px-4 py-4">
          <div class="flex items-center justify-between mb-3">
            <h4 class="text-sm font-medium text-gray-900">{{ level.name }}</h4>
                          <span class="text-xs px-2 py-1 rounded-full" :class="[getLevelColor(level.type).bg, getLevelColor(level.type).text]">
                {{ level.count }} کاربر
              </span>
          </div>
          
          <div class="space-y-3">
            <div>
              <label class="block text-xs font-medium text-gray-700 mb-1">تعداد پیشنهادات</label>
              <input v-model="level.recommendationCount" type="range" min="1" max="20" class="w-full h-1 bg-gray-200 rounded-lg appearance-none cursor-pointer">
              <div class="flex justify-between text-xs text-gray-500 mt-1">
                <span>1</span>
                <span>{{ level.recommendationCount }}</span>
                <span>20</span>
              </div>
            </div>
            
            <div>
              <label class="block text-xs font-medium text-gray-700 mb-1">فراوانی نمایش</label>
              <select v-model="level.frequency" class="w-full px-2 py-1 text-xs border border-gray-300 rounded focus:ring-1 focus:ring-indigo-500 focus:border-indigo-500">
                <option value="high">زیاد</option>
                <option value="medium">متوسط</option>
                <option value="low">کم</option>
              </select>
            </div>
            
            <div>
              <label class="block text-xs font-medium text-gray-700 mb-1">نوع پیشنهادات</label>
              <div class="space-y-1">
                <label class="flex items-center">
                  <input v-model="level.types" type="checkbox" value="popular" class="rounded border-gray-300 text-indigo-600 focus:ring-indigo-500 w-3 h-3">
                  <span class="mr-1 text-xs text-gray-700">محبوب</span>
                </label>
                <label class="flex items-center">
                  <input v-model="level.types" type="checkbox" value="trending" class="rounded border-gray-300 text-indigo-600 focus:ring-indigo-500 w-3 h-3">
                  <span class="mr-1 text-xs text-gray-700">ترند</span>
                </label>
                <label class="flex items-center">
                  <input v-model="level.types" type="checkbox" value="personalized" class="rounded border-gray-300 text-indigo-600 focus:ring-indigo-500 w-3 h-3">
                  <span class="mr-1 text-xs text-gray-700">شخصی</span>
                </label>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
// تنظیمات شخصی‌سازی
const settings = ref({
  intensity: {
    enabled: true,
    homePage: 8,
    productPage: 6,
    cartPage: 4,
    frequency: 'smart',
    minInterval: 5,
    maxDaily: 25
  },
  categoryFilter: {
    enabled: true,
    excludedCategories: [
      { id: 1, name: 'الکل و مشروبات' },
      { id: 2, name: 'محصولات غیرمجاز' }
    ]
  },
  brandFilter: {
    enabled: true,
    excludedBrands: [
      { id: 1, name: 'برند غیرمجاز 1' },
      { id: 2, name: 'برند غیرمجاز 2' }
    ]
  },
  priority: {
    enabled: true,
    types: [
      { id: 1, name: 'پیشنهادات شخصی', description: 'بر اساس تاریخچه کاربر', weight: 40 },
      { id: 2, name: 'محصولات محبوب', description: 'محصولات با امتیاز بالا', weight: 25 },
      { id: 3, name: 'محصولات جدید', description: 'محصولات تازه اضافه شده', weight: 20 },
      { id: 4, name: 'تخفیف‌ها', description: 'محصولات با تخفیف', weight: 15 }
    ]
  },
  culture: {
    enabled: true,
    language: 'fa',
    region: 'iran',
    culturalEvents: ['nowruz', 'ramadan'],
    style: 'friendly',
    diversity: 70,
    sensitivity: 'medium'
  },
  userLevel: {
    enabled: true,
    levels: [
      {
        id: 1,
        name: 'کاربران جدید',
        type: 'new',
        count: 3247,
        recommendationCount: 5,
        frequency: 'medium',
        types: ['popular', 'trending']
      },
      {
        id: 2,
        name: 'کاربران فعال',
        type: 'active',
        count: 5678,
        recommendationCount: 8,
        frequency: 'high',
        types: ['popular', 'personalized']
      },
      {
        id: 3,
        name: 'کاربران قدیمی',
        type: 'veteran',
        count: 3922,
        recommendationCount: 12,
        frequency: 'high',
        types: ['personalized', 'trending']
      }
    ]
  }
})

// داده‌های کمکی
const newCategory = ref('')
const newBrand = ref('')

const availableCategories = ref([
  { id: 3, name: 'الکترونیک' },
  { id: 4, name: 'پوشاک' },
  { id: 5, name: 'کتاب' },
  { id: 6, name: 'ورزشی' }
])

const availableBrands = ref([
  { id: 3, name: 'سامسونگ' },
  { id: 4, name: 'اپل' },
  { id: 5, name: 'شیائومی' },
  { id: 6, name: 'هوآوی' }
])

// توابع کمکی
const getLevelColor = (type: string) => {
  const colors = {
    new: { bg: 'bg-blue-100', text: 'text-blue-800' },
    active: { bg: 'bg-green-100', text: 'text-green-800' },
    veteran: { bg: 'bg-purple-100', text: 'text-purple-800' }
  }
  return colors[type] || colors.new
}

const removeCategory = (id: number) => {
  settings.value.categoryFilter.excludedCategories = settings.value.categoryFilter.excludedCategories.filter(cat => cat.id !== id)
}

const addCategory = () => {
  if (newCategory.value) {
    const category = availableCategories.value.find(cat => cat.id === parseInt(newCategory.value))
    if (category) {
      settings.value.categoryFilter.excludedCategories.push(category)
      newCategory.value = ''
    }
  }
}

const removeBrand = (id: number) => {
  settings.value.brandFilter.excludedBrands = settings.value.brandFilter.excludedBrands.filter(brand => brand.id !== id)
}

const addBrand = () => {
  if (newBrand.value) {
    const brand = availableBrands.value.find(b => b.id === parseInt(newBrand.value))
    if (brand) {
      settings.value.brandFilter.excludedBrands.push(brand)
      newBrand.value = ''
    }
  }
}

const movePriority = (index: number, direction: 'up' | 'down') => {
  const types = settings.value.priority.types
  if (direction === 'up' && index > 0) {
    [types[index], types[index - 1]] = [types[index - 1], types[index]]
  } else if (direction === 'down' && index < types.length - 1) {
    [types[index], types[index + 1]] = [types[index + 1], types[index]]
  }
}
</script> 
