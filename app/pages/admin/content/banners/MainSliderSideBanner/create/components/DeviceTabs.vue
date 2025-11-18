<template>
  <!-- کانتینر تب‌های دسکتاپ و موبایل -->
  <div class="mt-6 mx-4">
    <!-- تب‌ها -->
    <div class="mb-6">
      <nav class="flex rounded-lg overflow-hidden">
        <button
          @click="switchTab('desktop')"
          :class="[
            'w-1/2 text-center py-3 font-medium text-sm transition-colors',
            activeTab === 'desktop'
              ? 'bg-purple-200 text-purple-800'
              : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
          ]"
        >
          دسکتاپ
        </button>
        <button
          @click="switchTab('mobile')"
          :disabled="!isMobileEnabled"
          :class="[
            'w-1/2 text-center py-3 font-medium text-sm transition-colors',
            !isMobileEnabled 
              ? 'bg-gray-100 text-gray-400 cursor-not-allowed opacity-50'
              : activeTab === 'mobile'
              ? 'bg-purple-200 text-purple-800'
              : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
          ]"
          :title="!isMobileEnabled ? 'نمایش در موبایل غیرفعال است' : ''"
        >
          موبایل
          <span v-if="!isMobileEnabled" class="text-xs block">(غیرفعال)</span>
        </button>
      </nav>
    </div>

    <!-- محتوای تب دسکتاپ -->
    <div v-if="activeTab === 'desktop'">
      <slot name="desktop-content"></slot>
    </div>

    <!-- محتوای تب موبایل -->
    <div v-if="activeTab === 'mobile'">
      <!-- تنظیمات اسلایدر -->
      <div class="bg-white rounded-xl shadow p-8 mt-6 mx-4">
        <div class="flex justify-between items-center mb-6">
          <h3 class="text-lg font-bold text-gray-700">تنظیمات موبایل</h3>
          <div class="flex items-center gap-2 border-2 border-blue-200 rounded-lg p-1 bg-blue-50">
            <input
              type="checkbox"
              id="easyLoadMobile"
              v-model="props.sliderConfig.easy_load_enabled"
              class="w-4 h-4 text-blue-600 bg-blue-100 border-blue-300 rounded focus:ring-blue-500 focus:ring-2"
            />
            <label for="easyLoadMobile" class="text-sm font-medium text-blue-700">لیزی لود</label>
          </div>
        </div>
        
        <div class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-6 gap-6">
          <!-- پدینگ بالا -->
          <div>
            <label class="block mb-2 text-sm font-medium text-gray-700">پدینگ بالا (px)</label>
            <input
              type="number"
              :value="props.sliderConfig.mobile_padding_top !== undefined ? props.sliderConfig.mobile_padding_top : ''"
              @input="e => props.sliderConfig.mobile_padding_top = (e.target as HTMLInputElement).value === '' ? undefined : Number((e.target as HTMLInputElement).value)"
              min="0"
              max="100"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
              placeholder="0"
            />
          </div>

          <!-- پدینگ پایین -->
          <div>
            <label class="block mb-2 text-sm font-medium text-gray-700">پدینگ پایین (px)</label>
            <input
              type="number"
              :value="props.sliderConfig.mobile_padding_bottom !== undefined ? props.sliderConfig.mobile_padding_bottom : ''"
              @input="e => props.sliderConfig.mobile_padding_bottom = (e.target as HTMLInputElement).value === '' ? undefined : Number((e.target as HTMLInputElement).value)"
              min="0"
              max="100"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
              placeholder="0"
            />
          </div>

          <!-- مارجین راست -->
          <div>
            <label class="block mb-2 text-sm font-medium text-gray-700">مارجین راست (px)</label>
            <input
              type="number"
              :value="props.sliderConfig.mobile_margin_right !== undefined ? props.sliderConfig.mobile_margin_right : ''"
              @input="e => props.sliderConfig.mobile_margin_right = (e.target as HTMLInputElement).value === '' ? undefined : Number((e.target as HTMLInputElement).value)"
              min="0"
              max="100"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
              placeholder="0"
            />
          </div>

          <!-- مارجین چپ -->
          <div>
            <label class="block mb-2 text-sm font-medium text-gray-700">مارجین چپ (px)</label>
            <input
              type="number"
              :value="props.sliderConfig.mobile_margin_left !== undefined ? props.sliderConfig.mobile_margin_left : ''"
              @input="e => props.sliderConfig.mobile_margin_left = (e.target as HTMLInputElement).value === '' ? undefined : Number((e.target as HTMLInputElement).value)"
              min="0"
              max="100"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
              placeholder="0"
            />
          </div>

          <!-- پخش خودکار -->
          <div>
            <label class="block mb-2 text-sm font-medium text-gray-700">پخش خودکار</label>
            <select
              v-model="props.sliderConfig.mobile_autoplay_enabled"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
            >
              <option :value="true">فعال</option>
              <option :value="false">غیرفعال</option>
            </select>
          </div>

          <!-- تاخیر پخش خودکار -->
          <div>
            <label class="block mb-2 text-sm font-medium text-gray-700">تاخیر پخش خودکار (ثانیه)</label>
            <select
              v-model="props.sliderConfig.mobile_autoplay_delay"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
            >
              <option :value="3">3 ثانیه</option>
              <option :value="5">5 ثانیه</option>
              <option :value="7">7 ثانیه</option>
              <option :value="10">10 ثانیه</option>
              <option :value="15">15 ثانیه</option>
            </select>
          </div>

          <!-- حلقه اسلایدر -->
          <div>
            <label class="block mb-2 text-sm font-medium text-gray-700">حلقه اسلایدر</label>
            <select
              v-model="props.sliderConfig.mobile_loop_enabled"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
            >
              <option :value="true">فعال</option>
              <option :value="false">غیرفعال</option>
            </select>
          </div>

          <!-- نمایش دکمه‌های ناوبری -->
          <div>
            <label class="block mb-2 text-sm font-medium text-gray-700">نمایش دکمه‌های ناوبری</label>
            <select
              v-model="props.sliderConfig.mobile_show_navigation"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
            >
              <option :value="true">نمایش</option>
              <option :value="false">مخفی</option>
            </select>
          </div>

          <!-- نمایش نقطه‌های ناوبری -->
          <div>
            <label class="block mb-2 text-sm font-medium text-gray-700">نمایش نقطه‌های ناوبری</label>
            <select
              v-model="props.sliderConfig.mobile_show_pagination"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
            >
              <option :value="true">نمایش</option>
              <option :value="false">مخفی</option>
            </select>
          </div>

          <!-- پس‌زمینه فعال -->
          <div>
            <label class="block mb-2 text-sm font-medium text-gray-700">پس‌زمینه فعال</label>
            <select
              v-model="props.sliderConfig.bg_enabled"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
            >
              <option :value="true">فعال</option>
              <option :value="false">غیرفعال</option>
            </select>
          </div>

          <!-- عریض پس‌زمینه -->
          <div v-if="props.sliderConfig.bg_enabled">
            <label class="block mb-2 text-sm font-medium text-gray-700">عریض پس‌زمینه</label>
            <select
              v-model="props.sliderConfig.wide_bg"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
            >
              <option :value="true">بله</option>
              <option :value="false">خیر</option>
            </select>
          </div>

          <!-- رنگ پس‌زمینه -->
                    <div v-if="props.sliderConfig.bg_enabled">
                    <label class="block mb-2 text-sm font-medium text-gray-700">رنگ پس‌زمینه</label>
                    <input
                    type="color"
                    v-model="props.sliderConfig.bg_color"
                    class="w-full h-10 border border-gray-300 rounded-md"
                    />
                    </div>

                                                  <!-- عرض اسلایدر در موبایل -->
                    <div>
                      <label class="block mb-2 text-sm font-medium text-gray-700">عرض اسلایدر در موبایل</label>
                      <select
                        v-model="props.sliderConfig.mobile_slider_width"
                        class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
                      >
                        <option :value="'auto'">اتوماتیک</option>
                        <option :value="100">کوچک (100px)</option>
                        <option :value="300">متوسط (300px)</option>
                        <option :value="500">بزرگ (500px)</option>
                      </select>
                    </div>

                    <!-- عرض بنر در موبایل -->
                    <div>
                      <label class="block mb-2 text-sm font-medium text-gray-700">عرض بنر در موبایل</label>
                      <select
                        v-model="props.sliderConfig.mobile_banner_width"
                        class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-purple-400"
                      >
                        <option :value="'auto'">اتوماتیک</option>
                        <option :value="100">کوچک (100px)</option>
                        <option :value="300">متوسط (300px)</option>
                        <option :value="500">بزرگ (500px)</option>
                      </select>
                    </div>

                    <!-- ارتفاع اسلایدر -->
                    <div>
                    <label class="block mb-2 text-sm font-medium text-gray-700">ارتفاع اسلایدر (پیکسل)</label>
                    <select
                         v-model="mobileSliderHeight"
                         class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
                    >
                         <option :value="'auto'">اتوماتیک</option>
                         <option :value="100">100 پیکسل</option>
                         <option :value="150">150 پیکسل</option>
                         <option :value="200">200 پیکسل</option>
                         <option :value="250">250 پیکسل</option>
                    </select>
                    </div>

                    <!-- ارتفاع بنر در موبایل -->
                    <div>
                    <label class="block mb-2 text-sm font-medium text-gray-700">ارتفاع بنر در موبایل (پیکسل)</label>
                    <select
                         v-model="mobileBannerHeight"
                         class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
                    >
                         <option :value="'auto'">اتوماتیک</option>
                         <option :value="100">100 پیکسل</option>
                         <option :value="150">150 پیکسل</option>
                         <option :value="200">200 پیکسل</option>
                         <option :value="250">250 پیکسل</option>
                    </select>
                    </div>

                              
               </div>
               </div>

                                                                                                                        <!-- کانتینر پیش‌نمایش و مدیریت -->
                    <div class="flex gap-6 mt-4 w-full">
                                    <!-- پیش‌نمایش زنده -->
                                    <div class="rounded-xl shadow p-2 ml-auto max-w-xs">
                                    <h3 class="text-lg font-bold text-gray-700 mb-2">پیش‌نمایش زنده موبایل</h3>
                                    
                                              <div class="bg-white rounded-lg p-2 border-2 border-dashed border-gray-300">
                                    <!-- Layout container for mobile - vertical stack -->
                                                  <div
                                        class="flex flex-col gap-3 p-2 min-h-[500px] relative w-64 mx-auto"
                                    >
                                                  <!-- Slider Section - Full width for mobile -->
                                                  <div
                                        class="relative overflow-hidden rounded-lg w-full"
                                                        :style="{
                                        height: `${props.sliderConfig.mobile_height || 150}px`,
                                        position: 'relative',
                                        }"
                                    >
                  <!-- نمایش اسلایدها -->
                  <div v-if="mobileSlides.length > 0" class="relative h-full">
                  <div
                    v-for="(slide, index) in mobileSlides"
                    :key="slide.id || slide.order || index"
                    class="absolute inset-0 transition-opacity duration-500"
                    :class="{ 'opacity-100': index === activeMobileSlideIndex, 'opacity-0': index !== activeMobileSlideIndex }"
                  >
                    <img
                    :src="slide.mobile_image"
                    :alt="slide.title"
                    class="w-full h-full object-cover"
                    />
                                        <div
                                        v-if="props.sliderConfig.show_title || props.sliderConfig.show_description"
                                        class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-3"
                                        >
                                        <h4
                                              v-if="props.sliderConfig.show_title && slide.title && (slide.showTitle !== false)"
                                              class="text-white text-base font-bold mb-1"
                                        >
                                              {{ slide.title }}
                                        </h4>
                                        <p
                                              v-if="props.sliderConfig.show_description && slide.description"
                                              class="text-white/90 text-xs"
                                        >
                                              {{ slide.description }}
                                        </p>
                                        </div>
                                    </div>
                                    
                                    <!-- دکمه‌های ناوبری - کوچکتر برای موبایل -->
                                    <div v-if="props.sliderConfig.show_navigation" class="absolute inset-y-0 left-0 right-0 flex items-center justify-between px-2">
                                        <button class="bg-white/20 hover:bg-white/30 text-white rounded-full p-1.5 transition-colors">
                                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"></path>
                                        </svg>
                                        </button>
                                        <button class="bg-white/20 hover:bg-white/30 text-white rounded-full p-1.5 transition-colors">
                                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
                                        </svg>
                                        </button>
                                    </div>
                                    
                                    <!-- نقطه‌های ناوبری - کوچکتر برای موبایل -->
                  <div v-if="props.sliderConfig.show_pagination" class="absolute bottom-2 left-1/2 transform -translate-x-1/2 flex gap-1.5">
                    <div 
                    v-for="(slide, index) in mobileSlides" 
                    :key="slide.id || slide.order || index"
                    class="w-2.5 h-2.5 rounded-full transition-colors"
                    :class="index === activeMobileSlideIndex ? 'bg-white' : 'bg-white/50'"
                    ></div>
                                    </div>
                                    </div>
                                    
                                    <!-- پیام عدم وجود اسلاید -->
                  <div v-else class="flex items-center justify-center h-full text-gray-400">
                                    <div class="text-center">
                                        <svg class="w-12 h-12 mx-auto mb-3 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                                        </svg>
                    <p class="text-sm">هیچ تصویر موبایلی برای نمایش وجود ندارد</p>
                    <p class="text-xs text-gray-500">برای فعال شدن پیش‌نمایش موبایل، برای هر اسلاید تصویر موبایل بارگذاری کنید</p>
                                    </div>
                                    </div>
                                    </div>
                                    
                                                  <!-- Banner Section - Full width below slider for mobile -->
                                                  <div
                                        class="relative overflow-hidden rounded-lg w-full"
                                                        :style="{
                                        height: `${props.sliderConfig.mobile_banner_height || 80}px`,
                                        position: 'relative',
                                        }"
                                    >
                                    <!-- نمایش بنرها - عمودی برای موبایل -->
            <div v-if="mobileBanners.length > 0" class="relative h-full">
                                        <!-- نمایش بنرها به صورت عمودی -->
                                        <div class="h-full flex flex-col gap-2">
                                        <div
                  v-for="(banner, index) in mobileBanners"
                  :key="banner.id || banner.order || index"
                                              class="relative rounded-lg overflow-hidden flex-1"
                                                                  :style="{
                                                                  height: `${(props.sliderConfig.mobile_banner_height || 80) / mobileBanners.length - 4}px`
                                                                  }"
                                        >
                                              <img
                  :src="banner.mobile_image"
                                              :alt="banner.title"
                                              class="w-full h-full object-cover"
                                              />
                                              <div class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-2">
                                              <h4 class="text-white text-sm font-bold mb-1">{{ banner.title }}</h4>
                                              <p v-if="banner.description" class="text-white/90 text-xs">{{ banner.description }}</p>
                                              </div>
                                        </div>
                                        </div>
                                    </div>

                                    <!-- پیام عدم وجود بنر -->
                  <div v-else class="flex items-center justify-center h-full text-gray-400">
                                        <div class="text-center">
                                        <svg class="w-8 h-8 mx-auto mb-2 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                              
                                        </svg>
                    <p class="text-xs">هیچ بنر موبایلی اضافه نشده است</p>
                                        </div>
                                    </div>
                                    </div>
                                    </div>
                              </div>

                                                            <!-- اطلاعات تنظیمات موبایل -->
                                    <div class="mt-4 text-sm text-gray-600 text-center border-t border-gray-200 pt-4">
                                              <p>ارتفاع اسلایدر: {{ props.sliderConfig.mobile_height || 150 }}px |
                                        ارتفاع بنر: {{ props.sliderConfig.mobile_banner_height || 80 }}px |
                                        پس‌زمینه: {{ props.sliderConfig.bg_enabled ? 'فعال' : 'غیرفعال' }}</p>
                                    </div>
                                    </div>

                                                            <!-- نمایش و مدیریت اسلایدها -->
                                    <div class="bg-white rounded-xl shadow p-6 flex-1 min-h-[400px]">
                              <div class="flex justify-between items-center mb-6">
                                    <h3 class="text-lg font-bold text-gray-700">اسلایدر اصلی</h3>
                                    <button
                                    class="bg-cyan-500 text-white px-4 py-2 rounded-md font-bold hover:bg-cyan-600 transition-colors"
                                    @click="props.openAddSliderModal"
                                    >
                                    افزودن اسلایدر جدید
                                    </button>
                              </div>

          <div v-if="mobileSlides.length === 0" class="text-gray-400 text-center py-8">
            هیچ اسلایدی برای موبایل ثبت نشده است!
                              </div>

                              <div v-else class="flex flex-col gap-6">
                                              <div
                                    v-for="(slide, idx) in mobileSlides"
                                    :key="slide.id || slide.order || idx"
                                    class="flex items-center gap-3 p-2 rounded-lg bg-gray-50 w-full"
                                    >
                                             <div class="w-28 h-20 flex items-center justify-center rounded border-2 border-purple-200 bg-white overflow-hidden relative">
                <img
                v-if="slide.mobile_image"
                :src="slide.mobile_image"
                alt="اسلایدر موبایل"
                class="w-full h-full object-cover"
                />
                <span v-else class="text-xs text-gray-500 text-center px-2">بدون تصویر موبایل</span>
                </div>
                                    <div class="flex flex-col flex-1">
                                    <div class="font-bold text-sm text-gray-700 mb-1">{{ slide.title }}</div>
                                    <div v-if="slide.description" class="text-xs text-gray-600 mb-1">{{ slide.description }}</div>
                                    <div v-if="slide.link" class="text-xs text-blue-600 break-all">{{ slide.link }}</div>
                                    </div>
                                    <div class="flex gap-2">
                                    <button
                                        @click="props.editSlide(idx)"
                                        class="text-blue-500 hover:text-blue-700 p-1"
                                    >
                                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536M9 11l6 6M3 17v4h4l10.293-10.293a1 1 0 00-1.414-1.414L3 17z"></path>
                                        </svg>
                                    </button>
                                    <button
                                        @click="props.removeSlide(idx)"
                                        class="text-red-500 hover:text-red-700 p-1"
                                    >
                                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                                        </svg>
                                    </button>
                                    </div>
                                    </div>
                              </div>
                              </div>

                                                            <!-- کانتینر بنرها -->
                                    <div class="bg-white rounded-xl shadow p-6 flex-1 min-h-[400px]">
                              <div class="flex justify-between items-center mb-6">
                                    <div class="flex flex-col">
                                    <h3 class="text-lg font-bold text-gray-700">بنرهای کناری</h3>
                                    </div>
                                    <div class="flex flex-col items-end gap-2">
                                    <button
                                    :disabled="mobileBanners.length >= 2"
                                    class="bg-cyan-500 text-white px-4 py-2 rounded-md font-bold transition-colors disabled:bg-gray-400 disabled:cursor-not-allowed disabled:opacity-50"
                                    :class="mobileBanners.length >= 2 ? 'cursor-not-allowed' : 'hover:bg-cyan-600'"
                                    @click="props.openAddBannerModal"
                                    >
                                    افزودن بنر جدید
                                    </button>
                                    <p v-if="mobileBanners.length >= 2" class="text-sm text-red-600 font-medium">
                                    حداکثر ۲ بنر قابل اضافه کردن است
                                    </p>
                                    </div>
                              </div>

          <div v-if="mobileBanners.length === 0" class="text-gray-400 text-center py-8">
            هیچ بنری برای موبایل ثبت نشده است!
                              </div>
                              <div v-else class="flex flex-col gap-6">
                                              <div
                                    v-for="(banner, idx) in mobileBanners"
                                    :key="banner.id || banner.order || idx"
                                    class="flex items-center gap-3 p-2 rounded-lg bg-gray-50 w-full"
                                    >
                                             <div class="w-28 h-20 flex items-center justify-center rounded border-2 border-purple-200 bg-white overflow-hidden relative">
                <img
                v-if="banner.mobile_image"
                :src="banner.mobile_image"
                alt="بنر موبایل"
                class="w-full h-full object-cover"
                />
                <span v-else class="text-xs text-gray-500 text-center px-2">بدون تصویر موبایل</span>
                </div>
                                    <div class="flex flex-col flex-1">
                                    <div class="font-bold text-sm text-gray-700 mb-1">{{ banner.title }}</div>
                                    <div v-if="banner.description" class="text-xs text-gray-600 mb-1">{{ banner.description }}</div>
                                    <div v-if="banner.link" class="text-xs text-blue-600 break-all">{{ banner.link }}</div>
                                    </div>
                                    <div class="flex gap-2">
                                    <button
                                        @click="props.editBanner(idx)"
                                        class="text-blue-500 hover:text-blue-700 p-1"
                                    >
                                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536M9 11l6 6M3 17v4h4l10.293-10.293a1 1 0 00-1.414-1.414L3 17z"></path>
                                        </svg>
                                    </button>
                                    <button
                                        @click="props.removeBanner(idx)"
                                        class="text-red-500 hover:text-red-700 p-1"
                                    >
                                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                                        </svg>
                                    </button>
                                    </div>
                                    </div>
                                        </div>
                              </div>
                              </div>
                              </div>
          </div>
          </template>

          <script setup lang="ts">
          // Vue composables
          import { ref, computed, onMounted } from 'vue'

          // Props
          interface Props {
          sliderConfig: any
          currentPreviewSlide: number
          isMobileEnabled: boolean
          openAddSliderModal: () => void
          editSlide: (index: number) => void
          removeSlide: (index: number) => void
          openAddBannerModal: () => void
          editBanner: (index: number) => void
          removeBanner: (index: number) => void
          getSliderClasses: () => string
          getSliderWidth: () => string
          getBannerClasses: () => string
          getBannerWidth: () => string
          }

          const props = defineProps<Props>()

          const emit = defineEmits<{ (e: 'change-device', value: 'desktop' | 'mobile'): void }>()

          // Tab state
          const activeTab = ref<'desktop' | 'mobile'>('desktop')

          // helper to switch tabs and notify parent
          const switchTab = (tab: 'desktop' | 'mobile') => {
            // اگر تب موبایل غیرفعال است، اجازه سوییچ نده
            if (tab === 'mobile' && !props.isMobileEnabled) {
              return
            }
            activeTab.value = tab
            emit('change-device', tab)
          }

          // Expose activeTab for parent component
          defineExpose({
          activeTab
          })

          onMounted(() => {
            emit('change-device', activeTab.value)
          })

          // Collections split per device
          const desktopSlides = computed(() => (props.sliderConfig?.slides ?? []))
          const desktopBanners = computed(() => (props.sliderConfig?.side_banners ?? []))
          // آرایه‌های جداگانه برای موبایل
          const mobileSlides = computed(() => (props.sliderConfig?.mobile_slides ?? []))
          const mobileBanners = computed(() => (props.sliderConfig?.mobile_side_banners ?? []))

          const activeMobileSlideIndex = computed(() => {
            const count = mobileSlides.value.length
            if (count === 0) {
              return 0
            }
            const raw = props.currentPreviewSlide % count
            return raw < 0 ? raw + count : raw
          })

          // Computed values with default fallbacks
          const mobileSliderHeight = computed({
            get: () => props.sliderConfig.mobile_height,
            set: val => {
              if (props.sliderConfig) {
                props.sliderConfig.mobile_height = val
              }
            }
          })

          const mobileBannerHeight = computed({
            get: () => props.sliderConfig.mobile_banner_height,
            set: val => {
              if (props.sliderConfig) {
                props.sliderConfig.mobile_banner_height = val
              }
            }
          })
          </script>
