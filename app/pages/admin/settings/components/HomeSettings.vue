<template>
  <div v-if="hasAccess">
    <!-- Header Section -->
    <div class="relative overflow-hidden bg-gradient-to-br from-blue-50 via-indigo-50 to-purple-50 rounded-2xl p-8 mb-8">
      <div class="absolute top-0 right-0 w-32 h-32 bg-gradient-to-bl from-blue-200/30 to-purple-200/30 rounded-full -translate-y-16 translate-x-16"></div>
      <div class="absolute bottom-0 left-0 w-24 h-24 bg-gradient-to-tr from-indigo-200/30 to-pink-200/30 rounded-full translate-y-12 -translate-x-12"></div>
      
      <div class="relative z-10">
        <div class="flex items-center mb-4">
          <div class="w-12 h-12 bg-gradient-to-br from-blue-500 to-purple-600 rounded-xl flex items-center justify-center mr-4">
            <i class="i-heroicons-home text-white text-xl"></i>
          </div>
          <div>
            <h2 class="text-3xl font-bold bg-gradient-to-r from-gray-800 to-gray-600 bg-clip-text text-transparent">تنظیمات اصلی فروشگاه</h2>
            <p class="text-gray-600 mt-1">تنظیمات پایه و اطلاعات اصلی فروشگاه اینترنتی شما</p>
          </div>
        </div>
      </div>
    </div>

    <form class="space-y-8" @submit.prevent="saveSettings">
      <!-- اطلاعات پایه -->
      <div class="bg-white rounded-2xl shadow-lg border border-gray-100 overflow-hidden">
        <div class="bg-gradient-to-r from-blue-500 to-indigo-600 px-6 py-4">
          <div class="flex items-center">
            <i class="i-heroicons-building-storefront text-white text-xl mr-3"></i>
            <h3 class="text-xl font-bold text-white">اطلاعات پایه فروشگاه</h3>
          </div>
        </div>
        
        <div class="p-6 space-y-6">
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <div class="group">
              <label class="block text-sm font-semibold text-gray-700 mb-3 flex items-center">
                <i class="i-heroicons-language text-blue-500 mr-2"></i>
                نام فروشگاه (فارسی)
              </label>
              <div class="relative">
                <input 
                  v-model="localSettings.shopNameFa" 
                  type="text" 
                  class="w-full px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:border-blue-500 focus:ring-4 focus:ring-blue-100 transition-all duration-200 group-hover:border-gray-300"
                  placeholder="نام فروشگاه به فارسی"
                >
                <div class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400">
                  <i class="i-heroicons-flag text-sm"></i>
                </div>
              </div>
            </div>
            
            <div class="group">
              <label class="block text-sm font-semibold text-gray-700 mb-3 flex items-center">
                <i class="i-heroicons-globe-alt text-green-500 mr-2"></i>
                نام فروشگاه (انگلیسی)
              </label>
              <div class="relative">
                <input 
                  v-model="localSettings.shopNameEn" 
                  type="text" 
                  class="w-full px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:border-green-500 focus:ring-4 focus:ring-green-100 transition-all duration-200 group-hover:border-gray-300"
                  placeholder="Shop name in English"
                >
                <div class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400">
                  <i class="i-heroicons-flag text-sm"></i>
                </div>
              </div>
            </div>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div class="group">
              <label class="block text-sm font-semibold text-gray-700 mb-3 flex items-center">
                <i class="i-heroicons-photo text-purple-500 mr-2"></i>
                لوگو معمولی
              </label>
              <div class="relative">
                <div class="border-2 border-dashed border-gray-300 rounded-xl p-6 text-center hover:border-purple-400 transition-colors duration-200 group-hover:bg-purple-50">
                  <div v-if="localSettings.logo" class="mb-3 cursor-pointer" @click="openLogoUploader('logo')">
                    <img :src="localSettings.logo" alt="لوگو" class="w-16 h-16 object-contain mx-auto rounded-lg shadow-md hover:scale-110 transition-transform duration-200">
                    <p class="text-xs text-gray-500 mt-1">برای تغییر کلیک کنید</p>
                  </div>
                  <div v-else class="mb-3">
                    <i class="i-heroicons-photo text-4xl text-gray-400 mx-auto"></i>
                  </div>
                                     <div class="space-y-2">
                     <button type="button" class="w-full px-4 py-2 bg-gradient-to-r from-indigo-500 to-indigo-600 text-white rounded-lg hover:from-indigo-600 hover:to-indigo-700 transition-all duration-200 shadow-md hover:shadow-lg" @click="openMediaLibrary('logo')">
                       <i class="i-heroicons-arrow-up-tray mr-2"></i>
                       آپلود
                     </button>
                   </div>
                </div>
              </div>
            </div>
            
            <div class="group">
              <label class="block text-sm font-semibold text-gray-700 mb-3 flex items-center">
                <i class="i-heroicons-device-phone-mobile text-indigo-500 mr-2"></i>
                لوگو رتینا
              </label>
              <div class="relative">
                <div class="border-2 border-dashed border-gray-300 rounded-xl p-6 text-center hover:border-indigo-400 transition-colors duration-200 group-hover:bg-indigo-50">
                  <div v-if="localSettings.logoRetina" class="mb-3 cursor-pointer" @click="openLogoUploader('logoRetina')">
                    <img :src="localSettings.logoRetina" alt="لوگو رتینا" class="w-16 h-16 object-contain mx-auto rounded-lg shadow-md hover:scale-110 transition-transform duration-200">
                    <p class="text-xs text-gray-500 mt-1">برای تغییر کلیک کنید</p>
                  </div>
                  <div v-else class="mb-3">
                    <i class="i-heroicons-device-phone-mobile text-4xl text-gray-400 mx-auto"></i>
                  </div>
                                     <div class="space-y-2">
                     <button type="button" class="w-full px-4 py-2 bg-gradient-to-r from-purple-500 to-purple-600 text-white rounded-lg hover:from-purple-600 hover:to-purple-700 transition-all duration-200 shadow-md hover:shadow-lg" @click="openMediaLibrary('logoRetina')">
                       <i class="i-heroicons-arrow-up-tray mr-2"></i>
                       آپلود
                     </button>
                   </div>
                </div>
              </div>
            </div>
            
            <div class="group">
              <label class="block text-sm font-semibold text-gray-700 mb-3 flex items-center">
                <i class="i-heroicons-heart text-pink-500 mr-2"></i>
                فاویکون
              </label>
              <div class="relative">
                <div class="border-2 border-dashed border-gray-300 rounded-xl p-6 text-center hover:border-pink-400 transition-colors duration-200 group-hover:bg-pink-50">
                  <div v-if="localSettings.favicon" class="mb-3 cursor-pointer" @click="openLogoUploader('favicon')">
                    <img :src="localSettings.favicon" alt="فاویکون" class="w-12 h-12 object-contain mx-auto rounded-lg shadow-md hover:scale-110 transition-transform duration-200">
                    <p class="text-xs text-gray-500 mt-1">برای تغییر کلیک کنید</p>
                  </div>
                  <div v-else class="mb-3">
                    <i class="i-heroicons-heart text-3xl text-gray-400 mx-auto"></i>
                  </div>
                                     <div class="space-y-2">
                     <button type="button" class="w-full px-4 py-2 bg-gradient-to-r from-purple-500 to-purple-600 text-white rounded-lg hover:from-purple-600 hover:to-purple-700 transition-all duration-200 text-sm" @click="openMediaLibrary('favicon')">
                       <i class="i-heroicons-arrow-up-tray mr-2"></i>
                       آپلود
                     </button>
                   </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- تنظیمات منطقه‌ای -->
      <div class="bg-white rounded-2xl shadow-lg border border-gray-100 overflow-hidden">
        <div class="bg-gradient-to-r from-green-500 to-emerald-600 px-6 py-4">
          <div class="flex items-center">
            <i class="i-heroicons-globe-americas text-white text-xl mr-3"></i>
            <h3 class="text-xl font-bold text-white">تنظیمات منطقه‌ای</h3>
          </div>
        </div>
        
        <div class="p-6">
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div class="group">
              <label class="block text-sm font-semibold text-gray-700 mb-3 flex items-center">
                <i class="i-heroicons-language text-green-500 mr-2"></i>
                زبان پیش‌فرض
              </label>
              <div class="relative">
                <select v-model="localSettings.defaultLanguage" class="w-full px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:border-green-500 focus:ring-4 focus:ring-green-100 transition-all duration-200 group-hover:border-gray-300 appearance-none bg-white">
                  <option value="fa">🇮🇷 فارسی</option>
                  <option value="en">🇺🇸 English</option>
                  <option value="ar">🇸🇦 العربية</option>
                </select>
                <div class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 pointer-events-none">
                  <i class="i-heroicons-chevron-down text-sm"></i>
                </div>
              </div>
            </div>
            
            <div class="group">
              <label class="block text-sm font-semibold text-gray-700 mb-3 flex items-center">
                <i class="i-heroicons-clock text-blue-500 mr-2"></i>
                منطقه زمانی
              </label>
              <div class="relative">
                <select v-model="localSettings.timezone" class="w-full px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:border-blue-500 focus:ring-4 focus:ring-blue-100 transition-all duration-200 group-hover:border-gray-300 appearance-none bg-white">
                  <option value="Asia/Tehran">🇮🇷 تهران (UTC+3:30)</option>
                  <option value="UTC">🌍 UTC</option>
                  <option value="Europe/London">🇬🇧 لندن (UTC+0)</option>
                  <option value="America/New_York">🇺🇸 نیویورک (UTC-5)</option>
                </select>
                <div class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 pointer-events-none">
                  <i class="i-heroicons-chevron-down text-sm"></i>
                </div>
              </div>
            </div>
            
            <div class="group">
              <label class="block text-sm font-semibold text-gray-700 mb-3 flex items-center">
                <i class="i-heroicons-currency-dollar text-yellow-500 mr-2"></i>
                واحد پول پیش‌فرض
              </label>
              <div class="relative">
                <select v-model="localSettings.defaultCurrency" class="w-full px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:border-yellow-500 focus:ring-4 focus:ring-yellow-100 transition-all duration-200 group-hover:border-gray-300 appearance-none bg-white">
                  <option value="IRR">🇮🇷 ریال ایران</option>
                  <option value="USD">🇺🇸 دلار آمریکا</option>
                  <option value="EUR">🇪🇺 یورو</option>
                  <option value="GBP">🇬🇧 پوند انگلیس</option>
                </select>
                <div class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 pointer-events-none">
                  <i class="i-heroicons-chevron-down text-sm"></i>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- وضعیت فروشگاه -->
      <div class="bg-white rounded-2xl shadow-lg border border-gray-100 overflow-hidden">
        <div class="bg-gradient-to-r from-orange-500 to-red-600 px-6 py-4">
          <div class="flex items-center">
            <i class="i-heroicons-wrench-screwdriver text-white text-xl mr-3"></i>
            <h3 class="text-xl font-bold text-white">وضعیت فروشگاه</h3>
          </div>
        </div>
        
        <div class="p-6">
          <div class="flex items-center p-6 bg-gradient-to-r from-orange-50 to-red-50 rounded-xl border border-orange-200">
            <input 
              id="maintenanceMode" 
              v-model="localSettings.maintenanceMode" 
              type="checkbox"
              class="w-5 h-5 text-orange-600 border-gray-300 rounded focus:ring-orange-500"
            >
            <label for="maintenanceMode" class="mr-3 text-sm font-semibold text-gray-700 flex items-center">
              <i class="i-heroicons-exclamation-triangle text-orange-500 mr-2"></i>
              حالت تعمیر
            </label>
            <div class="mr-auto">
              <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-orange-100 text-orange-800">
                <i class="i-heroicons-information-circle mr-1"></i>
                فروشگاه در حالت تعمیر نمایش داده می‌شود
              </span>
            </div>
          </div>
          
          <div v-if="localSettings.maintenanceMode" class="mt-4">
            <label class="block text-sm font-semibold text-gray-700 mb-3 flex items-center">
              <i class="i-heroicons-chat-bubble-left-right text-red-500 mr-2"></i>
              پیام حالت تعمیر
            </label>
            <textarea 
              v-model="localSettings.maintenanceMessage" 
              rows="3"
              class="w-full px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:border-red-500 focus:ring-4 focus:ring-red-100 transition-all duration-200 resize-none"
              placeholder="پیام نمایش داده شده در حالت تعمیر..."
            ></textarea>
          </div>
        </div>
      </div>

      <!-- اطلاعات تماس -->
      <div class="bg-white rounded-2xl shadow-lg border border-gray-100 overflow-hidden">
        <div class="bg-gradient-to-r from-teal-500 to-cyan-600 px-6 py-4">
          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <i class="i-heroicons-phone text-white text-xl mr-3"></i>
              <h3 class="text-xl font-bold text-white">اطلاعات تماس</h3>
            </div>
            <button
              type="button"
              class="px-4 py-2 bg-white/20 hover:bg-white/30 text-white rounded-lg transition-all duration-200 text-sm flex items-center backdrop-blur-sm"
              @click="addLocation"
            >
              <i class="i-heroicons-plus mr-2"></i>
              افزودن آدرس جدید
            </button>
          </div>
        </div>
        
        <div class="p-6 space-y-8">
          <!-- آدرس‌ها و شماره‌های تماس -->
          <div class="space-y-6">

            <div class="space-y-6">
              <div
                v-for="(location, index) in localSettings.locations"
                :key="location.id || index"
                class="border border-gray-200 rounded-xl bg-gray-50/60 p-5 space-y-5"
              >
                <div class="flex items-center justify-between">
                  <div class="flex items-center text-xs font-semibold text-gray-600 gap-2">
                    <i class="i-heroicons-map-pin text-indigo-500 text-base"></i>
                    <span>آدرس {{ index + 1 }}</span>
                  </div>
                  <button
                    v-if="localSettings.locations.length > 1"
                    type="button"
                    class="text-sm text-red-600 hover:text-red-700 flex items-center gap-1"
                    @click="removeLocation(index)"
                  >
                    <i class="i-heroicons-trash text-base"></i>
                    حذف آدرس
                  </button>
                </div>

                <div class="space-y-4">
                  <div class="group">
                    <label class="block text-sm font-semibold text-gray-700 mb-3 flex items-center">
                      <i class="i-heroicons-tag text-purple-500 mr-2"></i>
                      عنوان آدرس (اختیاری)
                    </label>
                    <input
                      v-model="location.title"
                      type="text"
                      class="w-full px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:border-purple-500 focus:ring-4 focus:ring-purple-100 transition-all duration-200 group-hover:border-gray-300"
                      placeholder="مثلاً: فروشگاه مرکزی"
                    >
                  </div>

                  <div class="group">
                    <label class="block text-sm font-semibold text-gray-700 mb-3 flex items-center">
                      <i class="i-heroicons-home-modern text-blue-500 mr-2"></i>
                      آدرس کامل
                    </label>
                    <input
                      v-model="location.address"
                      type="text"
                      class="w-full px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:border-blue-500 focus:ring-4 focus:ring-blue-100 transition-all duration-200 group-hover:border-gray-300"
                      placeholder="آدرس کامل این شعبه"
                    >
                  </div>
                </div>

                <div class="group">
                  <label class="block text-sm font-semibold text-gray-700 mb-3 flex items-center">
                    <i class="i-heroicons-phone text-teal-500 mr-2"></i>
                    شماره تلفن‌های این آدرس
                  </label>
                  <div class="space-y-3">
                    <div
                      v-for="(phone, phoneIndex) in location.phones"
                      :key="`${location.id || index}-phone-${phoneIndex}`"
                      class="flex gap-2"
                    >
                      <input
                        v-model="location.phones[phoneIndex]"
                        type="tel"
                        class="flex-1 px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:border-teal-500 focus:ring-4 focus:ring-teal-100 transition-all duration-200"
                        placeholder="021-12345678"
                      >
                      <button
                        v-if="location.phones.length > 1"
                        type="button"
                        class="px-3 py-3 bg-gradient-to-r from-red-500 to-red-600 text-white rounded-xl hover:from-red-600 hover:to-red-700 transition-all duration-200"
                        @click="removeLocationPhone(index, phoneIndex)"
                      >
                        <i class="i-heroicons-trash text-sm"></i>
                      </button>
                    </div>
                    <button
                      type="button"
                      class="w-full px-4 py-2 bg-gradient-to-r from-teal-500 to-teal-600 text-white rounded-lg hover:from-teal-600 hover:to-teal-700 transition-all duration-200 text-sm"
                      @click="addLocationPhone(index)"
                    >
                      <i class="i-heroicons-plus mr-2"></i>
                      افزودن شماره برای این آدرس
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- شماره موبایل‌های مدیر -->
          <div class="border-t border-gray-200 pt-6 space-y-3">
            <label class="block text-sm font-semibold text-gray-700 flex items-center">
              <i class="i-heroicons-device-phone-mobile text-green-500 mr-2"></i>
              شماره موبایل‌های مدیر
            </label>
            <div class="space-y-3">
              <div
                v-for="(phone, index) in localSettings.adminPhones"
                :key="`admin-phone-${index}`"
                class="flex gap-2"
              >
                <input
                  v-model="localSettings.adminPhones[index]"
                  type="tel"
                  class="flex-1 px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:border-green-500 focus:ring-4 focus:ring-green-100 transition-all duration-200"
                  placeholder="09123456789"
                >
                <button
                  v-if="localSettings.adminPhones.length > 1"
                  type="button"
                  class="px-3 py-3 bg-gradient-to-r from-red-500 to-red-600 text-white rounded-xl hover:from-red-600 hover:to-red-700 transition-all duration-200"
                  @click="removeAdminPhone(index)"
                >
                  <i class="i-heroicons-trash text-sm"></i>
                </button>
              </div>
              <button
                type="button"
                class="w-full px-4 py-2 bg-gradient-to-r from-green-500 to-green-600 text-white rounded-lg hover:from-green-600 hover:to-green-700 transition-all duration-200 text-sm"
                @click="addAdminPhone"
              >
                <i class="i-heroicons-plus mr-2"></i>
                افزودن شماره جدید
              </button>
            </div>
          </div>

          <!-- ایمیل و مختصات -->
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <div class="group">
              <label class="block text-sm font-semibold text-gray-700 mb-3 flex items-center">
                <i class="i-heroicons-envelope text-blue-500 mr-2"></i>
                ایمیل فروشگاه
              </label>
              <input
                v-model="localSettings.email"
                type="email"
                class="w-full px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:border-blue-500 focus:ring-4 focus:ring-blue-100 transition-all duration-200 group-hover:border-gray-300"
                placeholder="info@shop.com"
              >
            </div>
            <div class="group">
              <label class="block text-sm font-semibold text-gray-700 mb-3 flex items-center">
                <i class="i-heroicons-map text-orange-500 mr-2"></i>
                مختصات جغرافیایی
              </label>
              <input
                v-model="localSettings.coordinates"
                type="text"
                class="w-full px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:border-orange-500 focus:ring-4 focus:ring-orange-100 transition-all duration-200 group-hover:border-gray-300"
                placeholder="مثلاً: 35.6892, 51.3890"
              >
              <p class="text-xs text-gray-500 mt-1">فرمت: عرض جغرافیایی, طول جغرافیایی</p>
            </div>
          </div>
        </div>
      </div>

      <!-- اطلاعات اضافی -->
      <div class="bg-white rounded-2xl shadow-lg border border-gray-100 overflow-hidden">
        <div class="bg-gradient-to-r from-purple-500 to-pink-600 px-6 py-4">
          <div class="flex items-center">
            <i class="i-heroicons-clock text-white text-xl mr-3"></i>
            <h3 class="text-xl font-bold text-white">اطلاعات اضافی</h3>
          </div>
        </div>
        
        <div class="p-6 space-y-6">
          <!-- ساعات کاری -->
          <div class="group">
            <label class="block text-sm font-semibold text-gray-700 mb-3 flex items-center">
              <i class="i-heroicons-clock text-purple-500 mr-2"></i>
              ساعات کاری
            </label>
            <div class="space-y-3">
              <div v-for="(hour, index) in localSettings.workingHours" :key="index" class="flex gap-2">
                <input 
                  v-model="localSettings.workingHours[index]" 
                  type="text" 
                  class="flex-1 px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:border-purple-500 focus:ring-4 focus:ring-purple-100 transition-all duration-200"
                  placeholder="مثلاً: شنبه تا چهارشنبه: 9 صبح تا 6 عصر"
                >
                <button 
                  v-if="localSettings.workingHours.length > 1"
                  type="button" 
                  class="px-3 py-3 bg-gradient-to-r from-red-500 to-red-600 text-white rounded-xl hover:from-red-600 hover:to-red-700 transition-all duration-200"
                  @click="removeWorkingHour(index)"
                >
                  <i class="i-heroicons-trash text-sm"></i>
                </button>
              </div>
              <button 
                type="button" 
                class="w-full px-4 py-3 bg-gradient-to-r from-purple-500 to-purple-600 text-white rounded-xl hover:from-purple-600 hover:to-purple-700 transition-all duration-200 shadow-md hover:shadow-lg"
                @click="addWorkingHour"
              >
                <i class="i-heroicons-plus mr-2"></i>
                افزودن ساعت کاری جدید
              </button>
            </div>
          </div>

          <!-- توضیحات کوتاه -->
          <div class="group">
            <label class="block text-sm font-semibold text-gray-700 mb-3 flex items-center">
              <i class="i-heroicons-document-text text-pink-500 mr-2"></i>
              توضیحات کوتاه
            </label>
            <textarea 
              v-model="localSettings.shortDescription" 
              rows="4"
              class="w-full px-4 py-3 border-2 border-gray-200 rounded-xl focus:outline-none focus:border-pink-500 focus:ring-4 focus:ring-pink-100 transition-all duration-200 group-hover:border-gray-300 resize-none"
              placeholder="توضیحات کوتاه درباره فروشگاه..."
            ></textarea>
          </div>
        </div>
      </div>

      <!-- دکمه‌های عملیات -->
      <div class="bg-gradient-to-r from-gray-50 to-gray-100 rounded-2xl p-6">
        <div class="flex flex-col sm:flex-row justify-end space-y-3 sm:space-y-0 sm:space-x-3 sm:space-x-reverse">
          <button 
            type="button" 
            class="px-8 py-3 border-2 border-gray-300 text-gray-700 rounded-xl hover:bg-gray-50 hover:border-gray-400 transition-all duration-200 font-semibold flex items-center justify-center"
            @click="resetSettings"
          >
            <i class="i-heroicons-arrow-path mr-2"></i>
            بازنشانی تنظیمات
          </button>
          <button 
            type="submit" 
            :disabled="saving"
            class="px-8 py-3 bg-gradient-to-r from-blue-500 to-indigo-600 text-white rounded-xl hover:from-blue-600 hover:to-indigo-700 disabled:opacity-50 transition-all duration-200 shadow-lg hover:shadow-xl font-semibold flex items-center justify-center"
          >
            <i v-if="saving" class="i-heroicons-arrow-path animate-spin mr-2"></i>
            <i v-else class="i-heroicons-check mr-2"></i>
            {{ saving ? 'در حال ذخیره...' : 'ذخیره تنظیمات' }}
          </button>
        </div>
      </div>
    </form>

    

    <!-- مودال کتابخانه رسانه -->
    <MediaLibraryModal
      v-model="showMediaLibrary"
      :model-selected="[]"
      default-category="library"
      @confirm="handleMediaSelected"
    />
  </div>
</template>

<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
</script>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue';
import MediaLibraryModal from '~/components/media/MediaLibraryModal.vue';
import { useAuth } from '~/composables/useAuth';


// احراز هویت
const { user, isAuthenticated } = useAuth();

// بررسی دسترسی admin
const hasAccess = computed(() => {
  if (!isAuthenticated.value) {
    return false;
  }

  const userRole = user.value?.role?.toLowerCase() || '';
  const adminRoles = ['admin', 'developer'];
  return adminRoles.includes(userRole);
});

// بررسی احراز هویت و دسترسی admin - نمایش 404 در صورت عدم دسترسی
const checkAuth = async (): Promise<void> => {
  if (!hasAccess.value) {
    await navigateTo('/404', { external: false });
  }
};

// بررسی احراز هویت در هنگام mount
onMounted(async () => {
  await checkAuth();
});

// بررسی احراز هویت هنگام تغییر وضعیت احراز هویت
watch([isAuthenticated, hasAccess], async () => {
  if (!hasAccess.value) {
    await checkAuth();
  }
});

interface Location {
  id?: string | number
  title: string
  address: string
  phones: string[]
}

interface HomeSettings {
  shopNameFa: string
  shopNameEn: string
  logo: string
  logoRetina: string
  favicon: string
  defaultLanguage: string
  timezone: string
  defaultCurrency: string
  maintenanceMode: boolean
  maintenanceMessage: string
  locations: Location[]
  adminPhones: string[]
  email: string
  coordinates: string
  workingHours: string[]
  shortDescription: string
}

const props = defineProps<{
  settings: HomeSettings
  saving: boolean
}>()

const emit = defineEmits(['save', 'reset', 'selectImage', 'addPhone', 'removePhone', 'addAdminPhone', 'removeAdminPhone', 'addWorkingHour', 'removeWorkingHour', 'addLocation', 'removeLocation', 'addLocationPhone', 'removeLocationPhone', 'update:settings'])

const localSettings = ref<HomeSettings>({ ...props.settings })

watch(() => props.settings, (newVal) => {
  localSettings.value = { ...newVal }
}, { deep: true })

watch(localSettings, (newVal) => {
  emit('update:settings', newVal)
}, { deep: true })

// State for modals
const showMediaLibrary = ref(false)
const currentLogoType = ref('')

// Methods
const saveSettings = () => {
  emit('save')
}

const resetSettings = () => {
  emit('reset')
}

// Logo upload methods
const openLogoUploader = (type: string) => {
  currentLogoType.value = type
  showMediaLibrary.value = true
}

const openMediaLibrary = (type: string) => {
  currentLogoType.value = type
  showMediaLibrary.value = true
}

interface MediaItem {
  url?: string
  path?: string
  [key: string]: unknown
}

const handleLogoUploaded = (imageData: MediaItem) => {
  if (imageData && imageData.url) {
    // به‌روزرسانی لوگو در تنظیمات
    switch (currentLogoType.value) {
      case 'logo':
        localSettings.value.logo = imageData.url
        break
      case 'logoRetina':
        localSettings.value.logoRetina = imageData.url
        break
      case 'favicon':
        localSettings.value.favicon = imageData.url
        break
    }
    
    // ارسال به کامپوننت والد برای ذخیره
    emit('selectImage', currentLogoType.value, imageData.url)
  }
}

const handleMediaSelected = (selectedMedia: MediaItem[]) => {
  if (selectedMedia && selectedMedia.length > 0) {
    const media = selectedMedia[0]
    handleLogoUploaded({ url: media.url || media.path })
  }
}

const addAdminPhone = () => {
  emit('addAdminPhone')
}

const removeAdminPhone = (index: number) => {
  emit('removeAdminPhone', index)
}

const addWorkingHour = () => {
  emit('addWorkingHour')
}

const removeWorkingHour = (index: number) => {
  emit('removeWorkingHour', index)
}

const addLocation = () => {
  emit('addLocation')
}

const removeLocation = (index: number) => {
  emit('removeLocation', index)
}

const addLocationPhone = (locationIndex: number) => {
  emit('addLocationPhone', locationIndex)
}

const removeLocationPhone = (locationIndex: number, phoneIndex: number) => {
  emit('removeLocationPhone', locationIndex, phoneIndex)
}
</script>
