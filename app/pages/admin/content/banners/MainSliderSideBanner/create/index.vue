<template>
  <div class="w-full">
    <!-- Title Section -->
    <div class="bg-white w-full shadow-sm border-2 border-blue-200 rounded-lg mt-4">
      <div class="flex justify-between items-center px-8 py-6">
        <h1 class="text-2xl font-bold text-gray-800">ایجاد اسلایدر اصلی و بنر کناری</h1>
        <div class="flex items-center gap-6">
          <button
            :disabled="isSaving"
            class="bg-green-500 text-white px-6 py-3 rounded-md font-bold hover:bg-green-600 transition-colors disabled:opacity-50 disabled:cursor-not-allowed flex items-center gap-2"
            @click="saveWidget"
          >
            <svg v-if="isSaving" class="w-5 h-5 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
            </svg>
            <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
            </svg>
            {{ isSaving ? 'در حال ذخیره...' : 'ذخیره ابزارک' }}
          </button>
          <TemplateButton
            bg-gradient="bg-gradient-to-r from-purple-400 to-purple-600"
            text-color="text-white"
            border-color="border border-purple-500"
            hover-class="hover:from-purple-500 hover:to-purple-700"
            size="medium"
            class="flex items-center gap-2"
            @click="$router.push('/admin/content/banners')"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
            </svg>
            بازگشت به ابزارک ها
          </TemplateButton>
        </div>
      </div>
    </div>

    <!-- Form Section -->
    <div class="bg-white rounded-xl shadow p-8 mt-6 mx-4">
      <form class="grid grid-cols-1 md:grid-cols-5 gap-6 items-end">
        <!-- نوع ابزارک -->
        <div>
          <label class="block mb-2 text-sm font-medium text-gray-700">نوع ابزارک</label>
          <select
            v-model="formData.type"
            class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400 bg-gray-100 cursor-not-allowed"
            disabled
            required
          >
            <option value="">انتخاب کنید</option>
            <option v-for="(label, type) in WIDGET_TYPE_LABELS" :key="type" :value="type">
              {{ label }}
            </option>
          </select>
        </div>

        <!-- عنوان ابزارک -->
        <div>
          <label class="block mb-2 text-sm font-medium text-gray-700">عنوان ابزارک</label>
          <input
            v-model="formData.title"
            type="text"
            class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
            placeholder="عنوان ابزارک را وارد کنید"
            required
          />
        </div>

        <!-- وضعیت -->
        <div>
          <label class="block mb-2 text-sm font-medium text-gray-700">وضعیت</label>
          <select
            v-model="formData.status"
            class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
          >
            <option value="active">فعال</option>
            <option value="inactive">غیرفعال</option>
            <option value="draft">پیش‌نویس</option>
          </select>
        </div>

        <!-- صفحه -->
        <div>
          <label class="block mb-2 text-sm font-medium text-gray-700">صفحه</label>
          <select
            v-model="formData.page"
            class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
          >
            <option value="home">صفحه اصلی</option>
            <option value="other">سایر بخش‌ها</option>
          </select>
        </div>

        <!-- نمایش در موبایل -->
        <div>
          <label class="block mb-2 text-sm font-medium text-gray-700">نمایش در موبایل</label>
          <select
            v-model="formData.show_on_mobile"
            class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
          >
            <option :value="true">فعال</option>
            <option :value="false">غیرفعال</option>
          </select>
        </div>
      </form>

      <!-- Error Message for Title Field -->
      <div v-if="titleError" class="mt-2">
        <p class="text-sm text-red-600" style="font-family: 'Yekan', sans-serif;">
          {{ titleError }}
        </p>
      </div>
    </div>

    <!-- کامپوننت تب‌های دسکتاپ و موبایل -->
    <DeviceTabs
      ref="deviceTabsRef"
      :slider-config="sliderConfig"
      :current-preview-slide="currentPreviewSlide"
      :is-mobile-enabled="formData.show_on_mobile"
      :open-add-slider-modal="openAddSliderModal"
      :edit-slide="editSlide"
      :remove-slide="removeSlide"
      :open-add-banner-modal="openAddBannerModal"
      :edit-banner="editBanner"
      :remove-banner="removeBanner"
      :get-slider-classes="getSliderClasses"
      :get-slider-width="getSliderWidth"
      :get-banner-classes="getBannerClasses"
      :get-banner-width="getBannerWidth"
      @change-device="activeDeviceTab = $event"
    >
      <template #desktop-content>
        <!-- تنظیمات اسلایدر -->
        <div class="bg-white rounded-xl shadow p-8 mt-6 mx-4">
          <div class="flex justify-between items-center mb-6">
            <h3 class="text-lg font-bold text-gray-700">تنظیمات دسکتاپ</h3>
            <div class="flex items-center gap-2 border-2 border-blue-200 rounded-lg p-1 bg-blue-50">
              <input
                id="easyLoad"
                v-model="sliderConfig.easy_load_enabled"
                type="checkbox"
                class="w-4 h-4 text-blue-600 bg-blue-100 border-blue-300 rounded focus:ring-blue-500 focus:ring-2"
              />
              <label for="easyLoad" class="text-sm font-medium text-blue-700">لیزی لود</label>
            </div>
          </div>
          
          <div class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-6 gap-6">
            <!-- پدینگ بالا -->
            <div>
              <label class="block mb-2 text-sm font-medium text-gray-700">پدینگ بالا (px)</label>
              <input
                v-model="sliderConfig.padding_top"
                type="number"
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
                v-model="sliderConfig.padding_bottom"
                type="number"
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
                v-model="sliderConfig.margin_right"
                type="number"
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
                v-model="sliderConfig.margin_left"
                type="number"
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
                v-model="sliderConfig.autoplay_enabled"
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
                v-model="sliderConfig.autoplay_delay"
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
                v-model="sliderConfig.loop_enabled"
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
                v-model="sliderConfig.show_navigation"
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
                v-model="sliderConfig.show_pagination"
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
                v-model="sliderConfig.bg_enabled"
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
              >
                <option :value="true">فعال</option>
                <option :value="false">غیرفعال</option>
              </select>
            </div>

            <!-- عریض پس‌زمینه -->
            <div v-if="sliderConfig.bg_enabled">
              <label class="block mb-2 text-sm font-medium text-gray-700">عریض پس‌زمینه</label>
              <select
                v-model="sliderConfig.wide_bg"
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
              >
                <option :value="true">بله</option>
                <option :value="false">خیر</option>
              </select>
            </div>

            <!-- رنگ پس‌زمینه -->
            <div v-if="sliderConfig.bg_enabled">
              <label class="block mb-2 text-sm font-medium text-gray-700">رنگ پس‌زمینه</label>
              <input
                v-model="sliderConfig.bg_color"
                type="color"
                class="w-full h-10 border border-gray-300 rounded-md"
              />
            </div>

            <!-- ارتفاع اسلایدر -->
            <div>
              <label class="block mb-2 text-sm font-medium text-gray-700">ارتفاع اسلایدر (پیکسل)</label>
              <select
                v-model="sliderConfig.height"
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
              >
                <option :value="300">300 پیکسل</option>
                <option :value="400">400 پیکسل</option>
                <option :value="500">500 پیکسل</option>
                <option :value="600">600 پیکسل</option>
                <option :value="700">700 پیکسل</option>
              </select>
            </div>

            <!-- عرض اسلایدر -->
            <div>
              <label class="block mb-2 text-sm font-medium text-gray-700">عرض اسلایدر</label>
              <select
                v-model="sliderConfig.slider_width"
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
              >
                <option :value="800">تمام عرض</option>
                <option :value="600">وسط</option>
              </select>
            </div>

            <!-- عرض دقیق وسط -->
            <div v-if="sliderConfig.slider_width === 600">
              <label class="block mb-2 text-sm font-medium text-gray-700">عرض دقیق وسط (px)</label>
              <input
                v-model="sliderConfig.center_width"
                type="number"
                min="200"
                max="2000"
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
                placeholder="1000"
              />
            </div>

            <!-- نسبت اسلایدر و بنر -->
            <div>
              <label class="block mb-2 text-sm font-medium text-gray-700">نسبت اسلایدر و بنر</label>
              <select
                v-model="sliderConfig.slider_banner_ratio"
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
              >
                <option value="8-2">80% اسلایدر - 20% بنر</option>
                <option value="8-4">70% اسلایدر - 30% بنر</option>
                <option value="7-5">58% اسلایدر - 42% بنر</option>
                <option value="6-6">50% اسلایدر - 50% بنر</option>
                <option value="5-7">42% اسلایدر - 58% بنر</option>
                <option value="4-8">30% اسلایدر - 70% بنر</option>
              </select>
            </div>

            <!-- جایگاه بنر -->
            <div>
              <label class="block mb-2 text-sm font-medium text-gray-700">جایگاه بنر</label>
              <select
                v-model="sliderConfig.banner_position"
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
              >
                <option value="left">سمت چپ</option>
                <option value="right">سمت راست</option>
              </select>
            </div>
          </div>
        </div>

        <!-- پیش‌نمایش زنده -->
        <div class="rounded-xl shadow p-8 mt-6 mx-4">
          <h3 class="text-lg font-bold text-gray-700 mb-6">پیش‌نمایش زنده</h3>
          
          <div class="bg-white rounded-lg p-6 border-2 border-dashed border-gray-300">
            <!-- Layout container for slider and banner -->
            <div
              class="flex gap-6 p-2 min-h-[400px] relative"
              :class="sliderConfig.slider_width === 800 ? 'w-full' : 'max-w-4xl mx-auto'"
            >
              <!-- Slider Section -->
              <div
                class="relative overflow-hidden rounded-lg min-w-[200px]"
                :class="getSliderClasses()"
                :style="{
                  height: `${sliderConfig.height}px`,
                  width: getSliderWidth(),
                  minWidth: '200px',
                  position: 'relative',
                }"
              >
              <!-- نمایش اسلایدها -->
              <div v-if="desktopSlides.length > 0" class="relative h-full">
                <div 
                  v-for="(slide, index) in desktopSlides" 
                  :key="index"
                  class="absolute inset-0 transition-opacity duration-500"
                  :class="{ 'opacity-100': index === currentPreviewSlide, 'opacity-0': index !== currentPreviewSlide }"
                >
                  <img 
                    :src="slide.image" 
                    :alt="slide.title"
                    class="w-full h-full object-cover"
                  />
                  <div
                    v-if="sliderConfig.show_title || sliderConfig.show_description"
                    class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-6"
                  >
                    <h4
                      v-if="sliderConfig.show_title && slide.title && (slide.showTitle !== false)"
                      class="text-white text-lg font-bold mb-2"
                    >
                      {{ slide.title }}
                    </h4>
                    <p
                      v-if="sliderConfig.show_description && slide.description"
                      class="text-white/90 text-sm"
                    >
                      {{ slide.description }}
                    </p>
                  </div>
                </div>
                
                <!-- دکمه‌های ناوبری -->
                <div v-if="sliderConfig.show_navigation" class="absolute inset-y-0 left-0 right-0 flex items-center justify-between px-4">
                  <button class="bg-white/20 hover:bg-white/30 text-white rounded-full p-2 transition-colors">
                    <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"></path>
                    </svg>
                  </button>
                  <button class="bg-white/20 hover:bg-white/30 text-white rounded-full p-2 transition-colors">
                    <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
                    </svg>
                  </button>
                </div>
                
                <!-- نقطه‌های ناوبری -->
                <div v-if="sliderConfig.show_pagination" class="absolute bottom-4 left-1/2 transform -translate-x-1/2 flex gap-2">
                  <div 
                    v-for="(slide, index) in desktopSlides" 
                    :key="index"
                    class="w-3 h-3 rounded-full transition-colors"
                    :class="index === currentPreviewSlide ? 'bg-white' : 'bg-white/50'"
                  ></div>
                </div>
              </div>
              
              <!-- پیام عدم وجود اسلاید -->
              <div v-else class="flex items-center justify-center h-full text-gray-400">
                <div class="text-center">
                  <svg class="w-16 h-16 mx-auto mb-4 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                  </svg>
                  <p class="text-lg">هیچ اسلایدی برای نمایش وجود ندارد</p>
                  <p class="text-sm text-gray-500">برای شروع، یک اسلاید جدید اضافه کنید</p>
                </div>
              </div>
            </div>
              <!-- Banner Section -->
              <div
                class="relative overflow-hidden rounded-lg min-w-[200px]"
                :class="getBannerClasses()"
                :style="{
                  height: `${sliderConfig.height}px`,
                  width: getBannerWidth(),
                  minWidth: '200px',
                  position: 'relative',
                }"
              >
                <!-- نمایش بنرها -->
                <div v-if="desktopBanners.length > 0" class="relative h-full flex flex-col gap-6">
                  <div
                    v-for="(banner, index) in desktopBanners"
                    :key="index"
                    class="relative rounded-lg overflow-hidden flex-1"
                  >
                    <img
                      :src="banner.image"
                      :alt="banner.title"
                      class="w-full h-full object-cover"
                    />
                    <div class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-3">
                      <h4 class="text-white text-sm font-bold mb-1">{{ banner.title }}</h4>
                      <p v-if="banner.description" class="text-white/90 text-xs">{{ banner.description }}</p>
                    </div>
                  </div>
                </div>

                <!-- پیام عدم وجود بنر -->
                <div v-else class="flex items-center justify-center h-full text-gray-400">
                  <div class="text-center">
                    <svg class="w-12 h-12 mx-auto mb-3 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      
                    </svg>
                    <p class="text-sm">هیچ بنری اضافه نشده است</p>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- اطلاعات تنظیمات -->
          <div class="mt-6 text-sm text-gray-600 text-center border-t border-gray-200 pt-4">
            <p>ارتفاع: {{ sliderConfig.height }}px |
                عرض: {{ sliderConfig.slider_width === 800 ? 'تمام عرض' : 'وسط صفحه' }} |
                پس‌زمینه: {{ sliderConfig.bg_enabled ? 'فعال' : 'غیرفعال' }}</p>
          </div>
        </div>

        <!-- نمایش و مدیریت اسلایدها -->
        <div class="bg-white rounded-xl shadow p-8 mt-6 mx-4">
          <div class="flex justify-between items-center mb-6">
            <h3 class="text-lg font-bold text-gray-700">اسلایدر اصلی</h3>
            <button
              class="bg-cyan-500 text-white px-4 py-2 rounded-md font-bold hover:bg-cyan-600 transition-colors"
              @click="openAddSliderModal"
            >
              افزودن اسلایدر جدید
            </button>
          </div>

          <div v-if="desktopSlides.length === 0" class="text-gray-400 text-center py-8">
            چیزی برای نمایش وجود ندارد!
          </div>

          <div v-else class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div
              v-for="(slide, idx) in desktopSlides"
              :key="slide.id || slide.order || idx"
              class="flex items-center gap-6 p-3 rounded-lg bg-gray-50"
            >
              <img
                :src="slide.image"
                alt="اسلایدر"
                class="w-32 h-20 object-cover rounded border-2 border-purple-200"
              />
              <div class="flex flex-col flex-1">
                <div class="font-bold text-sm text-gray-700 mb-1">{{ slide.title }}</div>
                <div v-if="slide.description" class="text-xs text-gray-600 mb-1">{{ slide.description }}</div>
                <div v-if="slide.link" class="text-xs text-blue-600 break-all">{{ slide.link }}</div>
              </div>
              <div class="flex gap-2">
                <button
                  class="text-blue-500 hover:text-blue-700 p-1"
                  @click="editSlide(sliderConfig.slides.indexOf(slide))"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536M9 11l6 6M3 17v4h4l10.293-10.293a1 1 0 00-1.414-1.414L3 17z"></path>
                  </svg>
                </button>
                <button
                  class="text-red-500 hover:text-red-700 p-1"
                  @click="removeSlide(sliderConfig.slides.indexOf(slide))"
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
        <div class="bg-white rounded-xl shadow p-8 mt-6 mx-4">
          <div class="flex justify-between items-center mb-6">
            <div class="flex flex-col">
              <h3 class="text-lg font-bold text-gray-700">بنرهای کناری</h3>
            </div>
            <div class="flex flex-col items-end gap-2">
              <button
                :disabled="sliderConfig.side_banners.length >= 2"
                class="bg-cyan-500 text-white px-4 py-2 rounded-md font-bold transition-colors disabled:bg-gray-400 disabled:cursor-not-allowed disabled:opacity-50"
                :class="sliderConfig.side_banners.length >= 2 ? 'cursor-not-allowed' : 'hover:bg-cyan-600'"
                @click="openAddBannerModal"
              >
                افزودن بنر جدید
              </button>
              <p v-if="sliderConfig.side_banners.length >= 2" class="text-sm text-red-600 font-medium">
                حداکثر ۲ بنر قابل اضافه کردن است
              </p>
            </div>
          </div>

          <div v-if="desktopBanners.length === 0" class="text-gray-400 text-center py-8">
            هیچ بنری اضافه نشده است!
          </div>
          <div v-else class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div
              v-for="(banner, idx) in desktopBanners"
              :key="banner.id || banner.order || idx"
              class="flex items-center gap-6 p-3 rounded-lg bg-gray-50"
            >
              <img
                :src="banner.image"
                alt="بنر"
                class="w-32 h-20 object-cover rounded border-2 border-purple-200"
              />
              <div class="flex flex-col flex-1">
                <div class="font-bold text-sm text-gray-700 mb-1">{{ banner.title }}</div>
                <div v-if="banner.description" class="text-xs text-gray-600 mb-1">{{ banner.description }}</div>
                <div v-if="banner.link" class="text-xs text-blue-600 break-all">{{ banner.link }}</div>
              </div>
              <div class="flex gap-2">
                <button
                  class="text-blue-500 hover:text-blue-700 p-1"
                  @click="editBanner(sliderConfig.side_banners.indexOf(banner))"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536M9 11l6 6M3 17v4h4l10.293-10.293a1 1 0 00-1.414-1.414L3 17z"></path>
                  </svg>
                </button>
                <button
                  class="text-red-500 hover:text-red-700 p-1"
                  @click="removeBanner(sliderConfig.side_banners.indexOf(banner))"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>
      </template>
      
      <template #mobile-content>
        <!-- محتوای تب موبایل حالا در کامپوننت DeviceTabs قرار دارد -->
        <!-- تنظیمات موبایل -->
        <div class="space-y-4">
          <h3 class="text-lg font-semibold text-gray-800 mb-4">تنظیمات موبایل</h3>
          
          <!-- ارتفاع اسلایدر در موبایل -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              ارتفاع اسلایدر (پیکسل)
            </label>
            <select
              v-model="sliderConfig.mobile_height"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
            >
              <option :value="100">100 پیکسل</option>
              <option :value="150">150 پیکسل</option>
              <option :value="200">200 پیکسل</option>
              <option :value="250">250 پیکسل</option>
              <option :value="300">300 پیکسل</option>
              <option :value="350">350 پیکسل</option>
            </select>
          </div>

          <!-- ارتفاع بنر در موبایل -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              ارتفاع بنر در موبایل (پیکسل)
            </label>
            <select
              v-model="sliderConfig.mobile_banner_height"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
            >
              <option :value="60">60 پیکسل</option>
              <option :value="80">80 پیکسل</option>
              <option :value="100">100 پیکسل</option>
              <option :value="120">120 پیکسل</option>
            </select>
          </div>

          <!-- جایگاه بنر در موبایل -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              جایگاه بنر
            </label>
            <select
              v-model="sliderConfig.banner_position"
              class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
            >
              <option value="left">چپ</option>
              <option value="right">راست</option>
            </select>
          </div>
        </div>
        <!-- پیش نمایش موبایل -->
        <div class="mt-8">
          <h3 class="text-lg font-semibold text-gray-800 mb-4">پیش نمایش موبایل</h3>
          <div class="max-w-xs mx-auto">
            <div class="w-64 mx-auto bg-gray-100 rounded-lg p-2 min-h-[500px] flex flex-col">
              <!-- اسلایدر اصلی -->
              <div 
                class="w-full bg-blue-200 rounded mb-2 flex items-center justify-center text-xs text-blue-800 font-medium"
                :style="{ height: `${sliderConfig.mobile_height}px` }"
              >
                اسلایدر اصلی
              </div>

              <!-- بنر کناری -->
              <div 
                class="w-full bg-purple-200 rounded mb-2 flex items-center justify-center text-xs text-purple-800 font-medium"
                :style="{ height: `${sliderConfig.mobile_banner_height}px` }"
              >
                بنر کناری
              </div>

              <!-- دکمه‌های کنترل -->
              <div class="flex justify-center space-x-2 mt-2">
                <button class="w-3 h-3 bg-blue-400 rounded-full"></button>
                <button class="w-3 h-3 bg-gray-300 rounded-full"></button>
                <button class="w-3 h-3 bg-gray-300 rounded-full"></button>
              </div>

              <!-- متن توضیحات -->
              <div class="text-center mt-2 text-xs text-gray-600">
                <p>عنوان محصول</p>
                <p class="text-gray-500">توضیحات کوتاه</p>
              </div>
            </div>
          </div>
        </div>
      </template>
    </DeviceTabs>

    <!-- Slide Modal Component -->
    <SlideModal
      :is-visible="showSliderModal"
      :is-editing="editingSlideIndex !== null"
      :slide-data="editingSlide"
      :show-title="showTitleInSlide"
      :device-type="activeDeviceTab"
      @update:is-visible="showSliderModal = $event"
      @update:show-title="showTitleInSlide = $event"
      @update:slide-data="editingSlide = $event"
      @save="handleSlideSave"
      @open-media-library="openMediaLibrary"
      @remove-image="removeImage"
    />

    <!-- Banner Modal Component -->
    <SlideModal
      :is-visible="showBannerModal"
      :is-editing="editingBannerIndex !== null"
      :slide-data="editingBanner"
      :show-title="showTitleInBanner"
      :device-type="activeDeviceTab"
      @update:is-visible="showBannerModal = $event"
      @update:show-title="showTitleInBanner = $event"
      @update:slide-data="editingBanner = $event"
      @save="handleBannerSave"
      @open-media-library="openMediaLibrary"
      @remove-image="removeImage"
    />

    <!-- Media Library Modal -->
    <MediaLibraryModal
      v-model="showMediaLibrary"
      file-type="image"
      default-category="banners"
      :max-select="1"
      @confirm="onSelectFromLibrary"
    />

    <!-- Page content will be added here -->

    <!-- فاصله انتهای صفحه -->
    <div class="pb-16"></div>
  </div>


</template>

<script setup lang="ts">
import { WIDGET_TYPE_LABELS } from '~/types/widget'
import type { Widget, SliderConfig, SlideItem } from '~/types/widget'
import TemplateButton from '~/components/common/TemplateButton.vue'
import MediaLibraryModal from '~/components/media/MediaLibraryModal.vue'
import SlideModal from '~/components/common/SlideModal.vue'
import DeviceTabs from './components/DeviceTabs.vue'
import { useToast } from '~/composables/useToast'
import { ref, computed, onMounted, watch, onUnmounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useWidget } from '~/composables/useWidget'

// تعریف definePageMeta و useHead برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
declare const useHead: (head: { title?: string }) => void

definePageMeta({ layout: 'admin-main', middleware: 'admin' })
useHead({ title: 'ایجاد اسلایدر همراه با بنر - پنل ادمین' })

// Route params
const route = useRoute()
const router = useRouter()
const widgetId = computed(() => route.params.id ? parseInt(route.params.id as string) : null)

// Composables
const { createWidget, fetchWidget, updateWidget, loading: _widgetLoading, error: _widgetError, clearError, widget: widgetData } = useWidget()
const { showSuccess, showError } = useToast()

// Props
interface Props {
  widget?: Widget
}

const _props = defineProps<Props>()

// Emits
const emit = defineEmits<{
  updated: [widget: Widget]
}>()

// State
const isSaving = ref(false)
const showSliderModal = ref(false)
const showBannerModal = ref(false)
const showMediaLibrary = ref(false)
const editingSlideIndex = ref<number | null>(null)
const editingBannerIndex = ref<number | null>(null)
const currentPreviewSlide = ref(0)
const currentPreviewBanner = ref(0)

// DeviceTabs ref to access exposed APIs
const deviceTabsRef = ref<InstanceType<typeof DeviceTabs> | null>(null)

// Track which device tab is active (desktop or mobile)
const activeDeviceTab = ref<'desktop' | 'mobile'>('desktop')

const desktopConfig = ref({
  enabled: true,
  minWidth: 768
})
const mobileConfig = ref({
  enabled: true,
  maxWidth: 767
})

// Error states
const titleError = ref<string>('')

// Validation functions
const validateTitle = () => {
  if (!formData.value.title?.trim()) {
    titleError.value = 'لطفاً عنوان ابزارک را وارد کنید.'
  } else if (formData.value.title.trim().length < 2) {
    titleError.value = 'عنوان ابزارک باید حداقل ۲ کاراکتر باشد.'
  } else if (formData.value.title.trim().length > 100) {
    titleError.value = 'عنوان ابزارک نمی‌تواند بیشتر از ۱۰۰ کاراکتر باشد.'
  } else {
    titleError.value = ''
  }
}

// Helper functions for preview
const getSliderWidth = () => {
  const ratio = sliderConfig.value.slider_banner_ratio || '6-6'
  const [slider, banner] = ratio.split('-').map(Number)
  const total = slider + banner
  const sliderPercent = (slider / total) * 100
  return `${sliderPercent}%`
}

const getBannerWidth = () => {
  const ratio = sliderConfig.value.slider_banner_ratio || '6-6'
  const [slider, banner] = ratio.split('-').map(Number)
  const total = slider + banner
  const bannerPercent = (banner / total) * 100
  return `${bannerPercent}%`
}

const getSliderClasses = () => {
  if (sliderConfig.value.banner_position === 'left') {
    return `order-2 flex-none`
  } else {
    return `order-1 flex-none`
  }
}

const getBannerClasses = () => {
  if (sliderConfig.value.banner_position === 'left') {
    return `order-1 flex-none`
  } else {
    return `order-2 flex-none`
  }
}

// Banner navigation functions
const _nextBanner = () => {
  if (desktopBanners.value.length > 1) {
    currentPreviewBanner.value = currentPreviewBanner.value === desktopBanners.value.length - 1
      ? 0
      : currentPreviewBanner.value + 1
  }
}

const _previousBanner = () => {
  if (desktopBanners.value.length > 1) {
    currentPreviewBanner.value = currentPreviewBanner.value === 0
      ? desktopBanners.value.length - 1
      : currentPreviewBanner.value - 1
  }
}

// Editing slide data
const editingSlide = ref<SlideItem>({
  title: '',
  description: '',
  image: '',
  mobile_image: '',
  link: '',
  openInNewTab: false,
  order: 1,
  status: 'active'
})

const editingBanner = ref<SlideItem>({
  title: '',
  description: '',
  image: '',
  mobile_image: '',
  link: '',
  openInNewTab: false,
  order: 1,
  status: 'active'
})

// Title display control
const showTitleInSlide = ref(false)
const showTitleInBanner = ref(false)

// Form data
const formData = ref({
  title: '',
  description: '',
  type: 'main-slider-side-banner',
  status: 'active',
  page: 'home',
  show_on_mobile: true
})

// Initialize form data when widget is available
const initializeFormData = () => {
  if (widgetData.value) {
    formData.value = {
      title: widgetData.value.title || '',
      description: widgetData.value.description || '',
      type: widgetData.value.type || 'main-slider-side-banner',
      status: widgetData.value.status || 'active',
      page: widgetData.value.page || 'home',
      show_on_mobile: widgetData.value.show_on_mobile !== undefined ? widgetData.value.show_on_mobile : true
    }
  }
}

// Watch for widget changes
watch(widgetData, (newWidget) => {
  if (newWidget) {
    initializeFormData()
  }
}, { immediate: true, deep: true })

// Watch for show_on_mobile changes - if mobile is disabled while on mobile tab, switch to desktop
watch(() => formData.value.show_on_mobile, (isEnabled) => {
  if (!isEnabled && activeDeviceTab.value === 'mobile') {
    activeDeviceTab.value = 'desktop'
    // Force DeviceTabs component to switch to desktop
    if (deviceTabsRef.value && deviceTabsRef.value.activeTab) {
      deviceTabsRef.value.activeTab = 'desktop'
    }
  }
})

// Computed properties for reactive form data - currently unused but kept for future use
const _widgetTitle = computed(() => widgetData.value?.title || '')
const _widgetType = computed(() => widgetData.value?.type || 'single-slider-side')
const _widgetStatus = computed(() => widgetData.value?.status || 'active')
const _widgetPage = computed(() => widgetData.value?.page || 'home')

// Slider config
const sliderConfig = ref<SliderConfig>({
  slider_count: 5,
  wide_bg: false,
  bg_color: '#ffffff',
  banner_position: 'right',
  display_order: 'asc',
  height: 400,
  mobile_height: 150,
  mobile_banner_height: 'auto',
  mobile_slider_width: 'auto',
  mobile_banner_width: 'auto',
  mobile_banner_position: 'top',
  slides: [],
  side_banners: [],
  mobile_slides: [],
  mobile_side_banners: [],
  bg_enabled: false,
  slider_width: 800,
  banner_width: 400,
  slider_banner_ratio: '6-6',
  navigation_enabled: true,
  autoplay_enabled: true,
  autoplay_delay: 5,
  loop_enabled: true,
  show_navigation: true,
  show_pagination: true,
  show_title: true,
  show_description: true,
  transition_duration: 500,
  pause_on_hover: true,
  navigation_style: 'default',
  navigation_size: 40,
  pagination_style: 'default',
  pagination_color: '#ffffff',
  padding_top: 0,
  padding_bottom: 0,
  margin_left: 0,
  margin_right: 0,
  border_radius: 8,
  shadow_enabled: false,
  shadow_style: 'default',
  easy_load_enabled: false,
  // تنظیمات جداگانه موبایل
  mobile_padding_top: 0,
  mobile_padding_bottom: 0,
  mobile_margin_left: 0,
  mobile_margin_right: 0,
  mobile_autoplay_enabled: true,
  mobile_autoplay_delay: 5,
  mobile_loop_enabled: true,
  mobile_show_navigation: true,
  mobile_show_pagination: true,
  mobile_show_title: true,
  mobile_show_description: true,
  mobile_bg_enabled: false,
  mobile_bg_color: '#ffffff',
  mobile_wide_bg: false
})

// Cached collections for cleaner templates
const desktopSlides = computed(() => sliderConfig.value.slides.filter(slide => Boolean(slide.image)))
const desktopBanners = computed(() => sliderConfig.value.side_banners.filter(banner => Boolean(banner.image)))

// Methods for managing slides
const openAddSliderModal = () => {
  // تعیین آرایه بر اساس تب فعال
  const slides = activeDeviceTab.value === 'mobile' 
    ? sliderConfig.value.mobile_slides 
    : sliderConfig.value.slides
    
  editingSlideIndex.value = null
  editingSlide.value = {
    title: '',
    description: '',
    image: '',
    mobile_image: '',
    link: '',
    openInNewTab: false,
    order: slides.length + 1,
    status: 'active'
  }
  showTitleInSlide.value = false // Reset to default
  showSliderModal.value = true
}

const editSlide = (index: number) => {
  // تعیین آرایه بر اساس تب فعال
  const slides = activeDeviceTab.value === 'mobile' 
    ? sliderConfig.value.mobile_slides 
    : sliderConfig.value.slides
    
  editingSlideIndex.value = index
  const slide = slides[index]
  editingSlide.value = {
    title: slide.title || '',
    description: slide.description || '',
    image: slide.image || '',
    mobile_image: slide.mobile_image || '',
    link: slide.link || '',
    openInNewTab: slide.openInNewTab || false,
    order: slide.order || index + 1,
    status: slide.status || 'active'
  }
  // Set showTitleInSlide based on existing slide data or default to false
  showTitleInSlide.value = slide.showTitle !== undefined ? slide.showTitle : false
  showSliderModal.value = true
}

const removeSlide = (index: number) => {
  // تعیین آرایه بر اساس تب فعال
  const slides = activeDeviceTab.value === 'mobile' 
    ? sliderConfig.value.mobile_slides 
    : sliderConfig.value.slides
    
  slides.splice(index, 1)
  // Re-order remaining slides
  slides.forEach((slide, idx) => {
    slide.order = idx + 1
  })
}

// Handle slide save from modal component
const handleSlideSave = (slideData: SlideItem, showTitle: boolean) => {
  // تعیین آرایه بر اساس تب فعال
  const slides = activeDeviceTab.value === 'mobile' 
    ? sliderConfig.value.mobile_slides 
    : sliderConfig.value.slides
  
  const slide: SlideItem = {
    ...slideData,
    title: slideData.title?.trim() || '',
    description: slideData.description?.trim() || '',
    link: slideData.link?.trim() || '',
    order: editingSlideIndex.value !== null ? slides[editingSlideIndex.value].order : slides.length + 1,
    status: 'active',
    showTitle: showTitle
  }

  if (editingSlideIndex.value !== null) {
    // ویرایش اسلاید موجود
    slides[editingSlideIndex.value] = slide
  } else {
    // افزودن اسلاید جدید
    slides.push(slide)
  }

  showSliderModal.value = false
  editingSlideIndex.value = null
}

// Methods for managing banners
const openAddBannerModal = () => {
  // تعیین آرایه بر اساس تب فعال
  const banners = activeDeviceTab.value === 'mobile' 
    ? sliderConfig.value.mobile_side_banners 
    : sliderConfig.value.side_banners
  
  // Check if maximum banners limit reached
  if (banners.length >= 2) {
    return // Don't open modal if already have 2 banners
  }

  editingBannerIndex.value = null
  editingBanner.value = {
    title: '',
    description: '',
    image: '',
    mobile_image: '',
    link: '',
    openInNewTab: false,
    order: banners.length + 1,
    status: 'active'
  }
  showTitleInBanner.value = false // Reset to default
  showBannerModal.value = true
}

const editBanner = (index: number) => {
  // تعیین آرایه بر اساس تب فعال
  const banners = activeDeviceTab.value === 'mobile' 
    ? sliderConfig.value.mobile_side_banners 
    : sliderConfig.value.side_banners
    
  editingBannerIndex.value = index
  const banner = banners[index]
  editingBanner.value = {
    title: banner.title || '',
    description: banner.description || '',
    image: banner.image || '',
    mobile_image: banner.mobile_image || '',
    link: banner.link || '',
    openInNewTab: banner.openInNewTab || false,
    order: banner.order || index + 1,
    status: banner.status || 'active',
    showTitle: banner.showTitle
  }
  // Set showTitleInBanner based on existing banner data or default to false
  showTitleInBanner.value = banner.showTitle !== undefined ? banner.showTitle : false
  showBannerModal.value = true
}

const removeBanner = (index: number) => {
  // تعیین آرایه بر اساس تب فعال
  const banners = activeDeviceTab.value === 'mobile' 
    ? sliderConfig.value.mobile_side_banners 
    : sliderConfig.value.side_banners
    
  banners.splice(index, 1)
  // Re-order remaining banners
  banners.forEach((banner, idx) => {
    banner.order = idx + 1
  })
}

// Handle banner save from modal component
const handleBannerSave = (bannerData: SlideItem, showTitle: boolean) => {
  // تعیین آرایه بر اساس تب فعال
  const banners = activeDeviceTab.value === 'mobile' 
    ? sliderConfig.value.mobile_side_banners 
    : sliderConfig.value.side_banners
  
  const banner: SlideItem = {
    ...bannerData,
    title: bannerData.title?.trim() || '',
    description: bannerData.description?.trim() || '',
    link: bannerData.link?.trim() || '',
    order: editingBannerIndex.value !== null ? banners[editingBannerIndex.value].order : banners.length + 1,
    status: 'active',
    showTitle: showTitle
  }

  if (editingBannerIndex.value !== null) {
    // ویرایش بنر موجود
    banners[editingBannerIndex.value] = banner
  } else {
    // افزودن بنر جدید
    banners.push(banner)
  }

  // Force reactivity update
  sliderConfig.value.side_banners = [...sliderConfig.value.side_banners]
  
  // Start banner rotation after adding/editing
  if (sliderConfig.value.side_banners.length > 0) {
    startBannerRotation()
  }

  showBannerModal.value = false
  editingBannerIndex.value = null
}

// Media library functions
const openMediaLibrary = () => {
  showMediaLibrary.value = true
}

interface MediaFile {
  url?: string
  id?: string | number
  [key: string]: unknown
}

const onSelectFromLibrary = (files: MediaFile[]) => {
  if (files && files.length > 0) {
    const file = files[0]
    // Check if we're editing a slide or banner
    if (showSliderModal.value) {
      // بسته به تب فعلی، عکس را در فیلد مناسب ذخیره کن
  if (activeDeviceTab.value === 'mobile') {
        editingSlide.value.mobile_image = file.url
      } else {
        editingSlide.value.image = file.url
      }
    } else if (showBannerModal.value) {
      // بسته به تب فعلی، عکس را در فیلد مناسب ذخیره کن
  if (activeDeviceTab.value === 'mobile') {
        editingBanner.value.mobile_image = file.url
      } else {
        editingBanner.value.image = file.url
      }
    }
    // Clear any previous image error
    // Note: This will be handled by the child component's reactive system
  }
  showMediaLibrary.value = false
}

// Image functions
const removeImage = () => {
  // Check if we're editing a slide or banner
  if (showSliderModal.value) {
    // بسته به تب فعلی، عکس مناسب را حذف کن
  if (activeDeviceTab.value === 'mobile') {
      editingSlide.value.mobile_image = ''
    } else {
      editingSlide.value.image = ''
    }
  } else if (showBannerModal.value) {
    // بسته به تب فعلی، عکس مناسب را حذف کن
  if (activeDeviceTab.value === 'mobile') {
      editingBanner.value.mobile_image = ''
    } else {
      editingBanner.value.image = ''
    }
  }
}

// Auto-rotation for live preview
let previewInterval: NodeJS.Timeout | null = null
let bannerInterval: NodeJS.Timeout | null = null

const startPreviewRotation = () => {
  if (previewInterval) clearInterval(previewInterval)

  if (sliderConfig.value.autoplay_enabled && desktopSlides.value.length > 1) {
    previewInterval = setInterval(() => {
      if (desktopSlides.value.length <= 1) {
        currentPreviewSlide.value = 0
        return
      }
      currentPreviewSlide.value = currentPreviewSlide.value === 0
        ? desktopSlides.value.length - 1
        : currentPreviewSlide.value - 1
    }, (sliderConfig.value.autoplay_delay || 5) * 1000)
  }
}

const startBannerRotation = () => {
  if (bannerInterval) clearInterval(bannerInterval)

  if (desktopBanners.value.length > 1) {
    bannerInterval = setInterval(() => {
      if (desktopBanners.value.length <= 1) {
        currentPreviewBanner.value = 0
        return
      }
      currentPreviewBanner.value = currentPreviewBanner.value === desktopBanners.value.length - 1
        ? 0
        : currentPreviewBanner.value + 1
    }, 4000) // 4 seconds for banners
  } else if (desktopBanners.value.length === 1) {
    // برای یک بنر، همیشه index 0 را نمایش بده
    currentPreviewBanner.value = 0
  }
}

const stopPreviewRotation = () => {
  if (previewInterval) {
    clearInterval(previewInterval)
    previewInterval = null
  }
}

const stopBannerRotation = () => {
  if (bannerInterval) {
    clearInterval(bannerInterval)
    bannerInterval = null
  }
}

// Watch for autoplay changes
watch(() => [sliderConfig.value.autoplay_enabled, sliderConfig.value.autoplay_delay, desktopSlides.value.length], () => {
  if (sliderConfig.value.autoplay_enabled && desktopSlides.value.length > 1) {
    startPreviewRotation()
  } else {
    stopPreviewRotation()
  }
}, { immediate: true, deep: true })

// Watch for banner changes
watch(() => desktopBanners.value.length, () => {
  if (desktopBanners.value.length > 0) {
    startBannerRotation()
  } else {
    stopBannerRotation()
  }
}, { immediate: true })

watch(desktopSlides, (slides) => {
  const count = slides.length
  if (count === 0) {
    currentPreviewSlide.value = 0
    stopPreviewRotation()
    return
  }

  if (currentPreviewSlide.value >= count) {
    currentPreviewSlide.value = 0
  }
}, { immediate: true })

watch(desktopBanners, (banners) => {
  const count = banners.length
  if (count === 0) {
    currentPreviewBanner.value = 0
    stopBannerRotation()
    return
  }

  if (currentPreviewBanner.value >= count) {
    currentPreviewBanner.value = 0
  }
}, { immediate: true })

// Cleanup on unmount
onUnmounted(() => {
  stopPreviewRotation()
  stopBannerRotation()
})

// Save widget function
const saveWidget = async () => {
  // Validate title field
  validateTitle()

  // Check if there are validation errors
  if (titleError.value) {
    // Scroll to the title field
    const titleInput = document.querySelector('input[placeholder="عنوان ابزارک را وارد کنید"]') as HTMLInputElement
    if (titleInput) {
      titleInput.focus()
      titleInput.scrollIntoView({ behavior: 'smooth', block: 'center' })
    }
    return
  }

  try {
    isSaving.value = true

    // Update form data from slider config
    const updatedConfig = {
      ...sliderConfig.value,
      slides: sliderConfig.value.slides,
      side_banners: sliderConfig.value.side_banners,
      mobile_slides: sliderConfig.value.mobile_slides,
      mobile_side_banners: sliderConfig.value.mobile_side_banners,
      desktop: desktopConfig.value,
      mobile: mobileConfig.value
    }
    
    // Prepare widget data for update
    const widgetData: Partial<Widget> = {
      title: formData.value.title,
      description: formData.value.description,
      type: formData.value.type as Widget['type'],
      status: formData.value.status as Widget['status'],
      page: formData.value.page as Widget['page'],
      show_on_mobile: formData.value.show_on_mobile,
      config: updatedConfig
    }
    
    if (widgetId.value) {
      // Update existing widget
      const updatedWidget = await updateWidget(widgetId.value, widgetData)
      showSuccess('ابزارک با موفقیت ذخیره شد')
      if (updatedWidget) {
        emit('updated', updatedWidget)
      }
    } else {
      // Create new widget
      const createData: import('~/types/widget').CreateWidgetRequest = {
        title: widgetData.title || '',
        description: widgetData.description,
        type: widgetData.type,
        status: widgetData.status,
        page: widgetData.page,
        show_on_mobile: widgetData.show_on_mobile,
        config: widgetData.config
      }
      const newWidget = await createWidget(createData)
      if (newWidget) {
        showSuccess('ابزارک با موفقیت ایجاد شد')
        // Redirect to the banner list page after creation
        setTimeout(() => {
          router.push('/admin/content/banners')
        }, 1500)
      }
    }
    
  } catch {
    showError('خطا در ذخیره ابزارک. لطفاً دوباره تلاش کنید.')
  } finally {
    isSaving.value = false
  }
}

// Initialize config from widget
onMounted(async () => {
  clearError()
  if (widgetId.value) {
    await fetchWidget(widgetId.value)
  }
  initializeFormData()
  
  // Only copy specific config fields, don't overwrite defaults
  if (widgetData.value?.config) {
    const config = widgetData.value.config as SliderConfig
    if (config.slides) sliderConfig.value.slides = config.slides
    if (config.side_banners) sliderConfig.value.side_banners = config.side_banners
    if (config.height) sliderConfig.value.height = config.height
    if (config.mobile_height) sliderConfig.value.mobile_height = config.mobile_height
    if (config.mobile_banner_height) sliderConfig.value.mobile_banner_height = config.mobile_banner_height
    if (config.slider_width) sliderConfig.value.slider_width = config.slider_width
    if (config.slider_banner_ratio) sliderConfig.value.slider_banner_ratio = config.slider_banner_ratio
    if (config.banner_position) sliderConfig.value.banner_position = config.banner_position
    if (config.bg_enabled !== undefined) sliderConfig.value.bg_enabled = config.bg_enabled
    if (config.bg_color) sliderConfig.value.bg_color = config.bg_color
    if (config.wide_bg !== undefined) sliderConfig.value.wide_bg = config.wide_bg
    if (config.autoplay_enabled !== undefined) sliderConfig.value.autoplay_enabled = config.autoplay_enabled
    if (config.autoplay_delay) sliderConfig.value.autoplay_delay = config.autoplay_delay
    if (config.loop_enabled !== undefined) sliderConfig.value.loop_enabled = config.loop_enabled
    if (config.show_navigation !== undefined) sliderConfig.value.show_navigation = config.show_navigation
    if (config.show_pagination !== undefined) sliderConfig.value.show_pagination = config.show_pagination
    if (config.easy_load_enabled !== undefined) sliderConfig.value.easy_load_enabled = config.easy_load_enabled
    
    // Add these lines to load padding and margin
    if (config.padding_top) sliderConfig.value.padding_top = config.padding_top
    if (config.padding_bottom) sliderConfig.value.padding_bottom = config.padding_bottom
    if (config.margin_left) sliderConfig.value.margin_left = config.margin_left
    if (config.margin_right) sliderConfig.value.margin_right = config.margin_right

    // Load desktop and mobile configurations
  }

  // Set default values if they are not loaded from the config
  if (sliderConfig.value.padding_top === 0) sliderConfig.value.padding_top = 0
  if (sliderConfig.value.padding_bottom === 0) sliderConfig.value.padding_bottom = 0
  if (sliderConfig.value.margin_left === 0) sliderConfig.value.margin_left = 0
  if (sliderConfig.value.margin_right === 0) sliderConfig.value.margin_right = 0

  // Initialize banner rotation
  if (sliderConfig.value.side_banners.length > 0) {
    startBannerRotation()
  }
})
</script>
