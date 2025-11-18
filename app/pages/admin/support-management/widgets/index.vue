<template>
  <div class="min-h-screen bg-gradient-to-b from-purple-50 to-gray-100" dir="rtl">
    <!-- نوار ناوبری بالا -->
    <div class="bg-white shadow-sm border-b border-gray-200">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex space-x-8 space-x-reverse">
          <button 
            v-for="tab in tabs" 
            :key="tab.id"
            @click="activeTab = tab.id"
            :class="[
              'py-4 px-1 border-b-2 font-medium text-sm transition-colors',
              activeTab === tab.id 
                ? 'border-purple-500 text-purple-600' 
                : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
            ]"
          >
            {{ tab.title }}
          </button>
        </div>
      </div>
    </div>

    <!-- محتوای اصلی -->
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <div class="bg-white rounded-lg shadow-lg">
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-8 p-8">
          
          <!-- ستون چپ - پیش‌نمایش ابزارک -->
          <div class="space-y-6">
            <h2 class="text-lg font-semibold text-gray-900 mb-4">پیش‌نمایش ابزارک</h2>
            
            <!-- پیش‌نمایش ابزارک چت -->
            <div class="bg-blue-600 rounded-lg p-6 max-w-sm mx-auto">
              <!-- تب‌های داخلی -->
              <div class="flex bg-blue-700 rounded-t-lg mb-4">
                <button class="flex-1 py-2 px-4 text-white text-sm bg-blue-800 rounded-t-lg">صفحه فرود</button>
                <button class="flex-1 py-2 px-4 text-white text-sm">صفحه چت</button>
              </div>
              
              <!-- هدر ابزارک -->
              <div class="flex items-center justify-between mb-4">
                <div class="w-8 h-8 bg-white rounded-full flex items-center justify-center">
                  <svg class="w-5 h-5 text-blue-600" fill="currentColor" viewBox="0 0 20 20">
                    <path fill-rule="evenodd" d="M18 10c0 3.866-3.582 7-8 7a8.841 8.841 0 01-4.083-.98L2 17l1.338-3.123C2.493 12.767 2 11.434 2 10c0-3.866 3.582-7 8-7s8 3.134 8 7zM7 9H5v2h2V9zm8 0h-2v2h2V9zM9 9h2v2H9V9z" clip-rule="evenodd" />
                  </svg>
                </div>
                <div class="flex -space-x-2">
                  <div class="w-6 h-6 bg-pink-400 rounded-full border-2 border-white"></div>
                  <div class="w-6 h-6 bg-purple-400 rounded-full border-2 border-white"></div>
                  <div class="w-6 h-6 bg-blue-400 rounded-full border-2 border-white"></div>
                </div>
              </div>
              
              <!-- عنوان و توضیحات -->
              <div class="text-center mb-4">
                <h3 class="text-white font-bold text-lg mb-1">پشتیبانی ایرانشیا</h3>
                <p class="text-blue-100 text-sm">هر سوالی داری بپرس</p>
              </div>
              
              <!-- بخش ارسال پیام -->
              <div class="bg-white rounded-lg p-3 mb-3">
                <div class="flex items-center">
                  <svg class="w-4 h-4 text-gray-400 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
                  </svg>
                  <div>
                    <p class="text-gray-800 text-sm font-medium">ارسال پیام برای ما</p>
                    <p class="text-gray-500 text-xs">ما به زودی پاسخگوی شما هستیم</p>
                  </div>
                </div>
              </div>
              
              <!-- نوار جستجو -->
              <div class="bg-white rounded-lg p-3 mb-3">
                <div class="flex items-center">
                  <svg class="w-4 h-4 text-gray-400 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                  </svg>
                  <span class="text-gray-500 text-sm">جست و جو کنید</span>
                </div>
              </div>
              
              <!-- لیست سوالات -->
              <div class="space-y-2">
                <div v-for="i in 4" :key="i" class="flex items-center text-white text-sm">
                  <svg class="w-3 h-3 text-blue-200 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
                  </svg>
                  <span>سوال {{ i }}</span>
                </div>
              </div>
            </div>
            
            <!-- دکمه شناور -->
            <div class="flex justify-center">
              <div class="w-12 h-12 bg-blue-600 rounded-full flex items-center justify-center shadow-lg">
                <svg class="w-6 h-6 text-white" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M18 10c0 3.866-3.582 7-8 7a8.841 8.841 0 01-4.083-.98L2 17l1.338-3.123C2.493 12.767 2 11.434 2 10c0-3.866 3.582-7 8-7s8 3.134 8 7zM7 9H5v2h2V9zm8 0h-2v2h2V9zM9 9h2v2H9V9z" clip-rule="evenodd" />
                </svg>
              </div>
            </div>
          </div>
          
          <!-- ستون راست - محتوای تب‌ها -->
          <div class="space-y-8">
            <!-- تب شخصی‌سازی -->
            <div v-if="activeTab === 'personalization'">
              <div>
                <h2 class="text-lg font-semibold text-gray-900 mb-2">شخصی‌سازی</h2>
                <p class="text-gray-600 text-sm">ظاهر ابزارکتان را مطابق با هویت بصری برندتان طراحی کنید.</p>
              </div>
              
              <!-- بخش لوگو -->
              <div class="space-y-4">
                <div>
                  <h3 class="text-md font-medium text-gray-900 mb-1">لوگو</h3>
                  <p class="text-gray-500 text-sm">امکان بارگذاری لوگوی سفارشی شما. ابعاد پیشنهادی ۶۰×۶۰ پیکسل، فرمت PNG.</p>
                </div>
                
                <button class="flex items-center space-x-2 space-x-reverse bg-gray-100 hover:bg-gray-200 text-gray-700 px-4 py-2 rounded-lg transition-colors">
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12" />
                  </svg>
                  <span>بارگذاری لوگوی اختصاصی</span>
                </button>
                
                <!-- لوگوهای پیش‌فرض -->
                <div class="flex space-x-3 space-x-reverse">
                  <div v-for="i in 8" :key="i" 
                       :class="[
                         'w-12 h-12 rounded-full flex items-center justify-center cursor-pointer border-2',
                         i === 6 ? 'border-purple-500 bg-purple-100' : 'border-gray-200 bg-blue-100'
                       ]"
                  >
                    <svg class="w-6 h-6 text-blue-600" fill="currentColor" viewBox="0 0 20 20">
                      <path fill-rule="evenodd" d="M18 10c0 3.866-3.582 7-8 7a8.841 8.841 0 01-4.083-.98L2 17l1.338-3.123C2.493 12.767 2 11.434 2 10c0-3.866 3.582-7 8-7s8 3.134 8 7zM7 9H5v2h2V9zm8 0h-2v2h2V9zM9 9h2v2H9V9z" clip-rule="evenodd" />
                    </svg>
                  </div>
                </div>
              </div>
              
              <!-- بخش رنگ‌ها -->
              <div class="space-y-6">
                <!-- رنگ پس زمینه -->
                <div class="space-y-3">
                  <div>
                    <h3 class="text-md font-medium text-gray-900 mb-1">رنگ پس زمینه</h3>
                    <p class="text-gray-500 text-sm">انتخاب رنگ برای پس زمینه سرصفحه ها و صفحه فرود.</p>
                  </div>
                  <div class="flex space-x-3 space-x-reverse">
                    <div v-for="(color, index) in backgroundColors" :key="index" class="flex items-center">
                      <input type="radio" :id="`bg-${index}`" name="background" :value="color" 
                             :checked="index === 0" class="mr-2">
                      <label :for="`bg-${index}`" class="cursor-pointer">
                        <div :class="`w-8 h-8 rounded-full ${color}`"></div>
                      </label>
                    </div>
                  </div>
                </div>
                
                <!-- رنگ پس زمینه لوگو -->
                <div class="space-y-3">
                  <div>
                    <h3 class="text-md font-medium text-gray-900 mb-1">رنگ پس زمینه لوگو</h3>
                    <p class="text-gray-500 text-sm">تعیین رنگ پس زمینه برای لوگوی ابزارک و صفحه فرود.</p>
                  </div>
                  <div class="flex space-x-3 space-x-reverse">
                    <div v-for="(color, index) in logoBackgroundColors" :key="index" class="flex items-center">
                      <input type="radio" :id="`logo-bg-${index}`" name="logoBackground" :value="color" 
                             :checked="index === 0" class="mr-2">
                      <label :for="`logo-bg-${index}`" class="cursor-pointer">
                        <div :class="`w-8 h-8 rounded-full ${color}`"></div>
                      </label>
                    </div>
                  </div>
                </div>
                
                <!-- رنگ مکمل -->
                <div class="space-y-3">
                  <div>
                    <h3 class="text-md font-medium text-gray-900 mb-1">رنگ مکمل</h3>
                    <p class="text-gray-500 text-sm">انتخاب رنگ برای دکمه ها و نقاط کلیدی.</p>
                  </div>
                  <div class="flex space-x-3 space-x-reverse">
                    <div v-for="(color, index) in accentColors" :key="index" class="flex items-center">
                      <input type="radio" :id="`accent-${index}`" name="accent" :value="color" 
                             :checked="index === 0" class="mr-2">
                      <label :for="`accent-${index}`" class="cursor-pointer">
                        <div :class="`w-8 h-8 rounded-full ${color}`"></div>
                      </label>
                    </div>
                  </div>
                </div>
                
                <!-- رنگ تیترها -->
                <div class="space-y-3">
                  <div>
                    <h3 class="text-md font-medium text-gray-900 mb-1">رنگ تیترها</h3>
                    <p class="text-gray-500 text-sm">رنگ تیتر هدر و لندینگ</p>
                  </div>
                  <div class="flex space-x-3 space-x-reverse">
                    <div v-for="(color, index) in headerColors" :key="index" class="flex items-center">
                      <input type="radio" :id="`header-${index}`" name="header" :value="color" 
                             :checked="index === 0" class="mr-2">
                      <label :for="`header-${index}`" class="cursor-pointer">
                        <div :class="`w-8 h-8 rounded-full ${color}`"></div>
                      </label>
                    </div>
                  </div>
                </div>
              </div>
              
              <!-- طرح پس زمینه -->
              <div class="space-y-4">
                <div>
                  <h3 class="text-md font-medium text-gray-900 mb-1">طرح پس زمینه</h3>
                  <p class="text-gray-500 text-sm">تنظیم طرح پس زمینه.</p>
                </div>
                <div class="grid grid-cols-5 gap-2">
                  <div v-for="i in 20" :key="i" 
                       :class="[
                         'w-12 h-12 rounded border-2 cursor-pointer',
                         i === 5 ? 'border-purple-500' : 'border-gray-200'
                       ]"
                       :style="`background: ${getPatternBackground(i)}`"
                  ></div>
                </div>
              </div>
            </div>

            <!-- تب تنظیمات پایه -->
            <div v-if="activeTab === 'basic-settings'">
              <div>
                <h2 class="text-lg font-semibold text-gray-900 mb-2">تنظیمات عمومی</h2>
                <p class="text-gray-600 text-sm">تنظیمات ویژه وب سایتها را از اینجا مدیریت کنید. هر وب سایت تنظیمات مستقلی دارد.</p>
              </div>
              
              <!-- بخش آپلود لوگو -->
              <div class="space-y-4">
                <div class="w-24 h-24 bg-gray-200 rounded-lg flex items-center justify-center">
                  <svg class="w-8 h-8 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                  </svg>
                </div>
                <p class="text-gray-500 text-sm">حداکثر حجم مجاز: ۵۰۰ کیلوبایت.</p>
                <button class="flex items-center space-x-2 space-x-reverse bg-gray-800 hover:bg-gray-700 text-white px-4 py-2 rounded-lg transition-colors">
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12" />
                  </svg>
                  <span>بارگذاری لوگو</span>
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                </button>
              </div>
              
              <!-- تنظیمات با کلید -->
              <div class="space-y-6">
                <div v-for="setting in basicSettings" :key="setting.id" class="flex items-center justify-between">
                  <div class="flex-1">
                    <h3 class="text-md font-medium text-gray-900 mb-1">{{ setting.title }}</h3>
                    <p class="text-gray-500 text-sm">{{ setting.description }}</p>
                  </div>
                  <button 
                    @click="setting.enabled = !setting.enabled"
                    :class="[
                      'relative inline-flex h-6 w-11 items-center rounded-full transition-colors',
                      setting.enabled ? 'bg-purple-600' : 'bg-gray-200'
                    ]"
                  >
                    <span 
                      :class="[
                        'inline-block h-4 w-4 transform rounded-full bg-white transition-transform',
                        setting.enabled ? 'translate-x-6' : 'translate-x-1'
                      ]"
                    ></span>
                  </button>
                </div>
              </div>
            </div>

            <!-- تب صفحه فرود -->
            <div v-if="activeTab === 'landing-page'">
              <div class="space-y-6">
                <!-- زبان ابزارک -->
                <div class="space-y-3">
                  <label class="block text-md font-medium text-gray-900">نمایش ابزارک به زبان</label>
                  <div class="relative">
                    <select class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500">
                      <option>فارسی</option>
                      <option>English</option>
                      <option>العربية</option>
                    </select>
                    <div class="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none">
                      <svg class="w-4 h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                      </svg>
                    </div>
                  </div>
                </div>

                <!-- عنوان -->
                <div class="space-y-3">
                  <label class="block text-md font-medium text-gray-900">متنی که در ابزارک گفتگو نمایش داده میشود .</label>
                  <label class="block text-sm text-gray-500">فارسی (پیش فرض)</label>
                  <input 
                    type="text" 
                    v-model="landingPageSettings.title"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                    placeholder="پشتیبانی ایرانشیا"
                  >
                </div>

                <!-- معرفی تیم -->
                <div class="space-y-3">
                  <label class="block text-md font-medium text-gray-900">تیم خود را معرفی کنید بگویید چطور میتوانید به مشتری کمک بکنید .</label>
                  <label class="block text-sm text-gray-500">فارسی (پیش فرض)</label>
                  <textarea 
                    v-model="landingPageSettings.teamIntro"
                    rows="3"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                    placeholder="هر سوالی داری بپرس"
                  ></textarea>
                </div>

                <!-- شروع مکالمه -->
                <div class="space-y-4">
                  <h3 class="text-md font-medium text-gray-900">شروع مکالمه</h3>
                  
                  <!-- متن -->
                  <div class="space-y-3">
                    <label class="block text-md font-medium text-gray-900">متن</label>
                    <input 
                      type="text" 
                      v-model="landingPageSettings.startConversation.text"
                      class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                      placeholder="ارسال پیام برای ما"
                    >
                  </div>

                  <!-- توضیحات -->
                  <div class="space-y-3">
                    <label class="block text-md font-medium text-gray-900">توضیحات</label>
                    <input 
                      type="text" 
                      v-model="landingPageSettings.startConversation.description"
                      class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                      placeholder="ما به زودی پاسخگوی شما هستیم"
                    >
                  </div>
                </div>
              </div>
            </div>

            <!-- تب تنظیمات نمایش -->
            <div v-if="activeTab === 'display-settings'">
              <div>
                <h2 class="text-lg font-semibold text-gray-900 mb-2">تنظیم نمایش ابزارک</h2>
              </div>
              
              <div class="space-y-6">
                <!-- حالت تعطیلات -->
                <div class="space-y-3">
                  <div class="flex items-center justify-between">
                    <div>
                      <h3 class="text-md font-medium text-gray-900">حالت تعطیلات</h3>
                      <p class="text-gray-500 text-sm">(فعال سازی عدم نمایش ابزارک)</p>
                    </div>
                    <button 
                      @click="displaySettings.holidayMode = !displaySettings.holidayMode"
                      :class="[
                        'relative inline-flex h-6 w-11 items-center rounded-full transition-colors',
                        displaySettings.holidayMode ? 'bg-purple-600' : 'bg-gray-200'
                      ]"
                    >
                      <span 
                        :class="[
                          'inline-block h-4 w-4 transform rounded-full bg-white transition-transform',
                          displaySettings.holidayMode ? 'translate-x-6' : 'translate-x-1'
                        ]"
                      ></span>
                    </button>
                  </div>
                  <p class="text-gray-500 text-sm">توجه: وارد کردن پروتکل (http یا https) در آدرس صفحات ضروری است.</p>
                </div>

                <!-- نمایش ابزارک فقط در این صفحات -->
                <div class="space-y-3">
                  <div class="flex items-center justify-between">
                    <h3 class="text-md font-medium text-gray-900">نمایش ابزارک فقط در این صفحات</h3>
                    <button 
                      @click="displaySettings.showOnlyOnPages = !displaySettings.showOnlyOnPages"
                      :class="[
                        'relative inline-flex h-6 w-11 items-center rounded-full transition-colors',
                        displaySettings.showOnlyOnPages ? 'bg-purple-600' : 'bg-gray-200'
                      ]"
                    >
                      <span 
                        :class="[
                          'inline-block h-4 w-4 transform rounded-full bg-white transition-transform',
                          displaySettings.showOnlyOnPages ? 'translate-x-6' : 'translate-x-1'
                        ]"
                      ></span>
                    </button>
                  </div>
                  <div class="flex space-x-2 space-x-reverse">
                    <input 
                      type="text" 
                      placeholder="لطفا آدرس صفحه را وارد کنید. مثال: http://example.com"
                      class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                    >
                    <button class="flex items-center space-x-2 space-x-reverse bg-purple-600 hover:bg-purple-700 text-white px-4 py-2 rounded-lg transition-colors">
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
                      </svg>
                      <span>افزودن</span>
                    </button>
                  </div>
                </div>

                <!-- عدم نمایش ابزارک در این صفحات -->
                <div class="space-y-3">
                  <div class="flex items-center justify-between">
                    <h3 class="text-md font-medium text-gray-900">عدم نمایش ابزارک در این صفحات</h3>
                    <button 
                      @click="displaySettings.hideOnPages = !displaySettings.hideOnPages"
                      :class="[
                        'relative inline-flex h-6 w-11 items-center rounded-full transition-colors',
                        displaySettings.hideOnPages ? 'bg-purple-600' : 'bg-gray-200'
                      ]"
                    >
                      <span 
                        :class="[
                          'inline-block h-4 w-4 transform rounded-full bg-white transition-transform',
                          displaySettings.hideOnPages ? 'translate-x-6' : 'translate-x-1'
                        ]"
                      ></span>
                    </button>
                  </div>
                  <div class="flex space-x-2 space-x-reverse">
                    <input 
                      type="text" 
                      placeholder="لطفا آدرس صفحه را وارد کنید. مثال: http://example.com"
                      class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                    >
                    <button class="flex items-center space-x-2 space-x-reverse bg-purple-600 hover:bg-purple-700 text-white px-4 py-2 rounded-lg transition-colors">
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
                      </svg>
                      <span>افزودن</span>
                    </button>
                  </div>
                </div>
              </div>
            </div>

            <!-- تب موقعیت -->
            <div v-if="activeTab === 'position'">
              <div class="space-y-6">
                <!-- جهت وب‌سایت -->
                <div class="space-y-3">
                  <label class="block text-md font-medium text-gray-900">جهت وب سایت شما</label>
                  <div class="relative">
                    <select class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500">
                      <option>راست به چپ (پیش فرض)</option>
                      <option>چپ به راست</option>
                    </select>
                    <div class="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none">
                      <svg class="w-4 h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                      </svg>
                    </div>
                  </div>
                </div>

                <!-- موقعیت ابزارک گفتگو -->
                <div class="space-y-3">
                  <label class="block text-md font-medium text-gray-900">موقعیت ابزارک گفتگو</label>
                  <div class="relative">
                    <select class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500">
                      <option>پایین راست</option>
                      <option>پایین چپ</option>
                      <option>بالا راست</option>
                      <option>بالا چپ</option>
                    </select>
                    <div class="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none">
                      <svg class="w-4 h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                      </svg>
                    </div>
                  </div>
                </div>

                <!-- موقعیت (پیکسل) -->
                <div class="grid grid-cols-2 gap-6">
                  <div class="space-y-3">
                    <label class="block text-md font-medium text-gray-900">راست (پیکسل)</label>
                    <input 
                      type="number" 
                      v-model="positionSettings.right"
                      class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                      placeholder="5"
                    >
                  </div>
                  <div class="space-y-3">
                    <label class="block text-md font-medium text-gray-900">پایین (پیکسل)</label>
                    <input 
                      type="number" 
                      v-model="positionSettings.bottom"
                      class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                      placeholder="55"
                    >
                  </div>
                </div>

                <!-- نمایش ابزارک در موبایل -->
                <div class="space-y-3">
                  <div class="flex items-center justify-between">
                    <label class="block text-md font-medium text-gray-900">نمایش ابزارک در موبایل</label>
                    <button 
                      @click="positionSettings.showOnMobile = !positionSettings.showOnMobile"
                      :class="[
                        'relative inline-flex h-6 w-11 items-center rounded-full transition-colors',
                        positionSettings.showOnMobile ? 'bg-purple-600' : 'bg-gray-200'
                      ]"
                    >
                      <span 
                        :class="[
                          'inline-block h-4 w-4 transform rounded-full bg-white transition-transform',
                          positionSettings.showOnMobile ? 'translate-x-6' : 'translate-x-1'
                        ]"
                      ></span>
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string }) => void
</script>

<script setup lang="ts">
import { ref, onMounted } from 'vue'

definePageMeta({ layout: 'admin-main' })

// تب‌های ناوبری
const tabs = ref([
  { id: 'personalization', title: 'شخصی سازی' },
  { id: 'basic-settings', title: 'تنظیمات پایه' },
  { id: 'landing-page', title: 'صفحه فرود' },
  { id: 'display-settings', title: 'تنظیمات نمایش' },
  { id: 'position', title: 'موقعیت' }
])

const activeTab = ref('personalization')
// وضعیت ابزارک‌ها
const widgets = ref<any[]>([])
async function loadWidgets() {
  try {
    const res: any = await $fetch('/api/admin/chat/widgets')
    if (res?.status === 'success') widgets.value = res.data
  } catch (e) { console.error('loadWidgets failed', e) }
}

async function saveWidget(partial: any) {
  try {
    if (!partial?.id) {
      const res: any = await $fetch('/api/admin/chat/widgets', {
        method: 'POST',
        body: {
          name: 'Default',
          description: 'Chat widget',
          is_active: true,
          welcome_message: landingPageSettings.value.startConversation.text,
          offline_message: landingPageSettings.value.teamIntro,
          theme: 'light',
          position: 'bottom-right'
        }
      })
      if (res?.status === 'success' || res?.success) await loadWidgets()
    } else {
      const res: any = await $fetch(`/api/admin/chat/widgets/${partial.id}`, { method: 'PUT', body: partial })
      if (res?.success) await loadWidgets()
    }
    alert('ذخیره شد')
  } catch (e) { console.error('saveWidget failed', e); alert('خطا در ذخیره') }
}

onMounted(async () => { await loadWidgets() })

// رنگ‌های پس زمینه
const backgroundColors = ref([
  'bg-red-500',
  'bg-orange-500', 
  'bg-pink-500',
  'bg-purple-500',
  'bg-blue-500'
])

// رنگ‌های پس زمینه لوگو
const logoBackgroundColors = ref([
  'bg-pink-500',
  'bg-purple-500',
  'bg-purple-700',
  'bg-blue-500',
  'bg-blue-700'
])

// رنگ‌های مکمل
const accentColors = ref([
  'bg-red-500',
  'bg-orange-500',
  'bg-pink-500', 
  'bg-purple-500',
  'bg-blue-500'
])

// رنگ‌های تیتر
const headerColors = ref([
  'bg-pink-500',
  'bg-purple-500',
  'bg-purple-700',
  'bg-blue-500',
  'bg-blue-700'
])

// تنظیمات پایه
const basicSettings = ref([
  {
    id: 'typed-words',
    title: 'مشاهده کلمات تایپ شده',
    description: 'فعال/غیر فعال کردن نمایش جملات تایپ شده توسط کاربران قبل از ارسال.',
    enabled: true
  },
  {
    id: 'voice-access',
    title: 'دسترسی ارسال صدا (voice) در ابزارک',
    description: 'فعال/غیر فعال کردن امکان ارسال صدا (voice) توسط کاربران در ابزارک',
    enabled: true
  },
  {
    id: 'file-access',
    title: 'دسترسی ارسال فایل در ابزارک',
    description: 'فعال/غیر فعال کردن دسترسی کاربران به ارسال فایل در ابزارک',
    enabled: true
  },
  {
    id: 'hide-powered-by',
    title: 'مخفی سازی عبارت "قدرت گرفته از رایچت"',
    description: 'فعال/غیر فعال کردن نمایش این عبارت در ابزارک گفتگو',
    enabled: true
  },
  {
    id: 'info-form',
    title: 'ارسال فرم دریافت اطلاعات (تنظیمات)',
    description: 'تنظیم گرفتن ایمیل با شماره تلفن کاربران',
    enabled: false
  }
])

// تنظیمات صفحه فرود
const landingPageSettings = ref({
  title: 'پشتیبانی ایرانشیا',
  teamIntro: 'هر سوالی داری بپرس',
  startConversation: {
    text: 'ارسال پیام برای ما',
    description: 'ما به زودی پاسخگوی شما هستیم'
  }
})

// تنظیمات نمایش
const displaySettings = ref({
  holidayMode: false,
  showOnlyOnPages: false,
  hideOnPages: false
})

// تنظیمات موقعیت
const positionSettings = ref({
  right: 5,
  bottom: 55,
  showOnMobile: true
})

// تابع تولید طرح‌های پس زمینه
const getPatternBackground = (index) => {
  const patterns = [
    'linear-gradient(45deg, #f3f4f6 25%, transparent 25%), linear-gradient(-45deg, #f3f4f6 25%, transparent 25%), linear-gradient(45deg, transparent 75%, #f3f4f6 75%), linear-gradient(-45deg, transparent 75%, #f3f4f6 75%)',
    'radial-gradient(circle, #e5e7eb 1px, transparent 1px)',
    'linear-gradient(90deg, #f3f4f6 1px, transparent 1px), linear-gradient(0deg, #f3f4f6 1px, transparent 1px)',
    'linear-gradient(45deg, #e5e7eb 25%, transparent 25%), linear-gradient(-45deg, #e5e7eb 25%, transparent 25%)',
    '#8b5cf6', // solid purple
    'linear-gradient(135deg, #f3f4f6 25%, transparent 25%) -10px 0, linear-gradient(225deg, #f3f4f6 25%, transparent 25%) -10px 0, linear-gradient(315deg, #f3f4f6 25%, transparent 25%), linear-gradient(45deg, #f3f4f6 25%, transparent 25%)',
    'radial-gradient(circle at 20% 50%, #e5e7eb 1px, transparent 1px), radial-gradient(circle at 80% 50%, #e5e7eb 1px, transparent 1px)',
    'linear-gradient(45deg, #f3f4f6 25%, transparent 25%, transparent 75%, #f3f4f6 75%, #f3f4f6), linear-gradient(45deg, #f3f4f6 25%, transparent 25%, transparent 75%, #f3f4f6 75%, #f3f4f6)',
    'linear-gradient(90deg, #e5e7eb 1px, transparent 1px), linear-gradient(0deg, #e5e7eb 1px, transparent 1px)',
    'radial-gradient(circle, #f3f4f6 2px, transparent 2px)',
    'linear-gradient(45deg, #e5e7eb 25%, transparent 25%), linear-gradient(-45deg, #e5e7eb 25%, transparent 25%)',
    'linear-gradient(135deg, #f3f4f6 25%, transparent 25%) -10px 0, linear-gradient(225deg, #f3f4f6 25%, transparent 25%) -10px 0',
    'radial-gradient(circle at 25% 25%, #e5e7eb 1px, transparent 1px), radial-gradient(circle at 75% 75%, #e5e7eb 1px, transparent 1px)',
    'linear-gradient(45deg, #f3f4f6 25%, transparent 25%, transparent 75%, #f3f4f6 75%, #f3f4f6)',
    'linear-gradient(90deg, #e5e7eb 1px, transparent 1px)',
    'radial-gradient(circle, #f3f4f6 1px, transparent 1px)',
    'linear-gradient(45deg, #e5e7eb 25%, transparent 25%)',
    'linear-gradient(135deg, #f3f4f6 25%, transparent 25%)',
    'radial-gradient(circle at 50% 50%, #e5e7eb 1px, transparent 1px)'
  ]
  
  return patterns[index - 1] || '#f3f4f6'
}
</script> 
