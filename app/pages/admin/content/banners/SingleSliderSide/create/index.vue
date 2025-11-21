<template>
  <div class="w-full">
    <!-- Title Section -->
    <div class="bg-white w-full shadow-sm border-2 border-blue-200 rounded-lg mt-4">
      <div class="flex justify-between items-center px-8 py-6">
        <h1 class="text-2xl font-bold text-gray-800">ایجاد {{ widget?.title || 'ابزارک' }}</h1>
        <div class="flex items-center gap-6">
          <button
            :disabled="isSaving"
            class="bg-green-500 text-white px-6 py-2 rounded-md font-bold hover:bg-green-600 transition-colors disabled:opacity-50 disabled:cursor-not-allowed flex items-center gap-2"
            @click="saveWidget"
          >
            <svg v-if="isSaving" class="w-5 h-5 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
            </svg>
            <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
            </svg>
            {{ isSaving ? 'در حال ذخیره...' : 'ذخیره' }}
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
      <form class="grid grid-cols-1 md:grid-cols-4 gap-6 items-end">
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
          <label class="block mb-2 text-sm font-medium text-gray-700" style="font-family: 'Yekan', sans-serif;">عنوان ابزارک *</label>
          <input
            v-model="formData.title"
            type="text"
            class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400 h-11"
            style="font-family: 'Yekan', sans-serif;"
            :class="{ 'border-red-500 focus:ring-red-500': titleError }"
            placeholder="عنوان ابزارک را وارد کنید"
            required
            @input="validateTitle"
            @blur="validateTitle"
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
            <option value="category">دسته‌بندی</option>
            <option value="product">محصول</option>
            <option value="about">درباره ما</option>
            <option value="contact">تماس با ما</option>
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

    <!-- تنظیمات اسلایدر -->
    <DeviceTabs
      ref="deviceTabsRef"
      :slider-config="sliderConfig"
      :current-preview-slide="currentPreviewSlide"
      :open-add-slider-modal="openAddSliderModal"
      :edit-slide="editSlide"
      :remove-slide="removeSlide"
      :get-slider-classes="() => ''"
      :get-slider-width="() => ''"
    >
      <!-- محتوای تب دسکتاپ -->
      <template #desktop-content>
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
          
          <div class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 gap-6">
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
            </div>
          </div>
        </template>

      <!-- محتوای تب موبایل -->
      <template #mobile-content>
        <div class="bg-white rounded-xl shadow p-8 mt-6 mx-4">
          <div class="flex justify-between items-center mb-6">
            <h3 class="text-lg font-bold text-gray-700">تنظیمات موبایل</h3>
          </div>
          
          <div class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 gap-6">
            <!-- پخش خودکار موبایل -->
            <div>
              <label class="block mb-2 text-sm font-medium text-gray-700">پخش خودکار</label>
              <select
                v-model="sliderConfig.mobile_autoplay_enabled"
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
              >
                <option :value="true">فعال</option>
                <option :value="false">غیرفعال</option>
              </select>
            </div>

            <!-- تاخیر پخش خودکار موبایل -->
            <div>
              <label class="block mb-2 text-sm font-medium text-gray-700">تاخیر پخش خودکار (ثانیه)</label>
              <select
                v-model="sliderConfig.mobile_autoplay_delay"
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
              >
                <option :value="5">5 ثانیه</option>
                <option :value="7">7 ثانیه</option>
                <option :value="10">10 ثانیه</option>
                <option :value="15">15 ثانیه</option>
              </select>
            </div>

            <!-- حلقه اسلایدر موبایل -->
            <div>
              <label class="block mb-2 text-sm font-medium text-gray-700">حلقه اسلایدر</label>
              <select
                v-model="sliderConfig.mobile_loop_enabled"
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
              >
                <option :value="true">فعال</option>
                <option :value="false">غیرفعال</option>
              </select>
            </div>

            <!-- نمایش دکمه‌های ناوبری موبایل -->
            <div>
              <label class="block mb-2 text-sm font-medium text-gray-700">نمایش دکمه‌های ناوبری</label>
              <select
                v-model="sliderConfig.mobile_show_navigation"
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
              >
                <option :value="true">نمایش</option>
                <option :value="false">مخفی</option>
              </select>
            </div>

            <!-- نمایش نقطه‌های ناوبری موبایل -->
            <div>
              <label class="block mb-2 text-sm font-medium text-gray-700">نمایش نقطه‌های ناوبری</label>
              <select
                v-model="sliderConfig.mobile_show_pagination"
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
              >
                <option :value="true">نمایش</option>
                <option :value="false">مخفی</option>
              </select>
            </div>

            <!-- نمایش عمودی در موبایل -->
            <div>
              <label class="block mb-2 text-sm font-medium text-gray-700">نمایش عمودی</label>
              <select
                v-model="sliderConfig.mobile_vertical_display"
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
              >
                <option :value="true">فعال</option>
                <option :value="false">غیرفعال</option>
              </select>
            </div>
          </div>
        </div>
      </template>
    </DeviceTabs>


            <!-- پیش‌نمایش زنده -->
        <div v-if="deviceTabsRef?.activeTab === 'desktop'" class="bg-white rounded-xl shadow p-8 mt-6 mx-4">
      <h3 class="text-lg font-bold text-gray-700 mb-6">پیش‌نمایش زنده</h3>
      
      <div class="bg-gray-50 rounded-lg p-6 border-2 border-dashed border-gray-300">
        <div 
          class="relative overflow-hidden rounded-lg"
          :style="{
            height: `${sliderConfig.height}px`,
            backgroundColor: sliderConfig.bg_enabled ? sliderConfig.bg_color : 'transparent',
            maxWidth: sliderConfig.slider_width === 800 ? '100%' : '600px',
            margin: '0 auto'
          }"
        >
          <!-- نمایش اسلایدها -->
          <div v-if="sliderConfig.slides.length > 0" class="relative h-full">
            <div 
              v-for="(slide, index) in sliderConfig.slides" 
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
                v-for="(slide, index) in sliderConfig.slides" 
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
        
        <!-- اطلاعات تنظیمات -->
        <div class="mt-4 text-sm text-gray-600 text-center">
          <p>ارتفاع: {{ sliderConfig.height }}px | 
              عرض: {{ sliderConfig.slider_width === 800 ? 'تمام عرض' : 'وسط صفحه' }} | 
              پس‌زمینه: {{ sliderConfig.bg_enabled ? 'فعال' : 'غیرفعال' }}</p>
        </div>
      </div>
    </div>

            <!-- نمایش و مدیریت اسلایدها -->
        <div v-if="deviceTabsRef?.activeTab === 'desktop'" class="bg-white rounded-xl shadow p-8 mt-6 mx-4">
      <div class="flex justify-between items-center mb-6">
        <h3 class="text-lg font-bold text-gray-700">اسلایدر اصلی</h3>
        <button
          class="bg-cyan-500 text-white px-4 py-2 rounded-md font-bold hover:bg-cyan-600 transition-colors"
          @click="openAddSliderModal"
        >
          افزودن اسلایدر جدید
        </button>
      </div>

      <div v-if="sliderConfig.slides.length === 0" class="text-gray-400 text-center py-8">
        چیزی برای نمایش وجود ندارد!
      </div>

      <div v-else class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div
          v-for="(slide, idx) in sliderConfig.slides"
          :key="idx"
          class="flex items-center gap-6 p-3 border rounded-lg bg-gray-50"
        >
          <img
            :src="slide.image"
            alt="اسلایدر"
            class="w-32 h-20 object-cover rounded border"
          />
          <div class="flex flex-col flex-1">
            <div class="font-bold text-sm text-gray-700 mb-1">{{ slide.title }}</div>
            <div v-if="slide.description" class="text-xs text-gray-600 mb-1">{{ slide.description }}</div>
            <div v-if="slide.link" class="text-xs text-blue-600 break-all">{{ slide.link }}</div>
          </div>
          <div class="flex gap-2">
            <button
              class="text-blue-500 hover:text-blue-700 p-1"
              @click="editSlide(idx)"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536M9 11l6 6M3 17v4h4l10.293-10.293a1 1 0 00-1.414-1.414L3 17z"></path>
              </svg>
            </button>
            <button
              class="text-red-500 hover:text-red-700 p-1"
              @click="removeSlide(idx)"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
              </svg>
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Slide Modal Component -->
    <SlideModal
      :is-visible="showSliderModal"
      :is-editing="editingSlideIndex !== null"
      :slide-data="editingSlide"
      :show-title="showTitleInSlide"
      @update:is-visible="showSliderModal = $event"
      @update:show-title="showTitleInSlide = $event"
      @update:slide-data="editingSlide = $event"
      @save="handleSlideSave"
      @open-media-library="openMediaLibrary"
      @remove-image="removeImage"
    />

    <!-- Media Library Modal -->
    <MediaLibraryModal
      v-model="showMediaLibrary"
      file-type="image"
      default-category="banners"
      @confirm="onSelectFromLibrary"
    />

    <!-- Page content will be added here -->

    <!-- فاصله انتهای صفحه -->
    <div class="pb-16"></div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, onUnmounted, ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import SlideModal from '~/components/common/SlideModal.vue'
import TemplateButton from '~/components/common/TemplateButton.vue'
import MediaLibraryModal from '~/components/media/MediaLibraryModal.vue'
import { useWidget } from '~/composables/useWidget'
import type { SlideItem, SliderConfig, Widget } from '~/types/widget'
import { WIDGET_TYPE_LABELS } from '~/types/widget'
import DeviceTabs from './components/DeviceTabs.vue'

// تعریف definePageMeta و useHead برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
declare const useHead: (head: { title?: string }) => void

definePageMeta({ layout: 'admin-main' })
useHead({ title: 'ایجاد اسلایدر تکی - پنل ادمین' })

// Route params
const route = useRoute()
const router = useRouter()
const widgetId = parseInt(route.params.id as string)

// Composables
const { fetchWidget, createWidget, loading, error, clearError, widget: fetchedWidget } = useWidget()

// Props
interface Props {
  widget?: Widget
}

const props = defineProps<Props>()

// Emits
const emit = defineEmits<{
  updated: [widget: Widget]
}>()

// State
const isSaving = ref(false)
const showSliderModal = ref(false)
const showMediaLibrary = ref(false)
const editingSlideIndex = ref<number | null>(null)
const currentPreviewSlide = ref(0)
const deviceTabsRef = ref()

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

// Editing slide data
const editingSlide = ref<SlideItem>({
  title: '',
  description: '',
  image: '',
  link: '',
  order: 1,
  status: 'active'
})

// Title display control
const showTitleInSlide = ref(true)

// Form data
const formData = ref({
  title: '',
  description: '',
  type: 'single-slider-side',
  status: 'active',
  page: 'home'
})

// Initialize form data when widget is available
const initializeFormData = () => {
  if (fetchedWidget.value) {
    console.log('Widget data:', fetchedWidget.value) // Debug log
    formData.value = {
      title: fetchedWidget.value.title || '',
      description: fetchedWidget.value.description || '',
      type: fetchedWidget.value.type || 'single-slider-side',
      status: fetchedWidget.value.status || 'active',
      page: fetchedWidget.value.page || 'home'
    }
    console.log('Form data initialized:', formData.value) // Debug log
  }
}

// Watch for widget changes
watch(fetchedWidget, (newWidget) => {
  if (newWidget) {
    initializeFormData()
  }
}, { immediate: true })

// Computed properties for reactive form data
const widgetTitle = computed(() => fetchedWidget.value?.title || '')
const widgetType = computed(() => fetchedWidget.value?.type || 'single-slider-side')
const widgetStatus = computed(() => fetchedWidget.value?.status || 'active')
const widgetPage = computed(() => fetchedWidget.value?.page || 'home')

// Slider config
const sliderConfig = ref<SliderConfig>({
  slider_count: 5,
  wide_bg: false,
  bg_color: '#ffffff',
  display_order: 'asc',
  height: 400,
  slides: [],
  bg_enabled: false,
  slider_width: 800,
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
  // Mobile specific settings
  mobile_autoplay_enabled: true,
  mobile_autoplay_delay: 5,
  mobile_loop_enabled: true,
  mobile_show_navigation: true,
  mobile_show_pagination: true,
  mobile_vertical_display: false,
  mobile_slider_width: 300,
  mobile_height: 150
})

// Methods for managing slides
const openAddSliderModal = () => {
  editingSlideIndex.value = null
  editingSlide.value = {
    title: '',
    description: '',
    image: '',
    link: '',
    order: sliderConfig.value.slides.length + 1,
    status: 'active'
  }
  showTitleInSlide.value = true // Reset to default
  showSliderModal.value = true
}

const editSlide = (index: number) => {
  editingSlideIndex.value = index
  const slide = sliderConfig.value.slides[index]
  editingSlide.value = {
    title: slide.title || '',
    description: slide.description || '',
    image: slide.image || '',
    link: slide.link || '',
    order: slide.order || index + 1,
    status: slide.status || 'active'
  }
  // Set showTitleInSlide based on existing slide data or default to true
  showTitleInSlide.value = slide.showTitle !== undefined ? slide.showTitle : true
  showSliderModal.value = true
}

const removeSlide = (index: number) => {
  sliderConfig.value.slides.splice(index, 1)
  // Re-order remaining slides
  sliderConfig.value.slides.forEach((slide, idx) => {
    slide.order = idx + 1
  })
}

// Handle slide save from modal component
const handleSlideSave = (slideData: SlideItem, showTitle: boolean) => {
  const slide: SlideItem = {
    ...slideData,
    title: slideData.title?.trim() || '',
    description: slideData.description?.trim() || '',
    link: slideData.link?.trim() || '',
    order: editingSlideIndex.value !== null ? sliderConfig.value.slides[editingSlideIndex.value].order : sliderConfig.value.slides.length + 1,
    status: 'active',
    showTitle: showTitle
  }

  if (editingSlideIndex.value !== null) {
    // ویرایش اسلاید موجود
    sliderConfig.value.slides[editingSlideIndex.value] = slide
  } else {
    // افزودن اسلاید جدید
    sliderConfig.value.slides.push(slide)
  }

  showSliderModal.value = false
  editingSlideIndex.value = null
}

// Media library functions
const openMediaLibrary = () => {
  showMediaLibrary.value = true
}

const onSelectFromLibrary = (files: any[]) => {
  if (files && files.length > 0) {
    const file = files[0]
    editingSlide.value.image = file.url
    // Clear any previous image error
    // Note: This will be handled by the child component's reactive system
  }
  showMediaLibrary.value = false
}

// Image functions
const removeImage = () => {
  editingSlide.value.image = ''
}

// Auto-rotation for live preview
let previewInterval: NodeJS.Timeout | null = null

const startPreviewRotation = () => {
  if (previewInterval) clearInterval(previewInterval)
  
  if (sliderConfig.value.autoplay_enabled && sliderConfig.value.slides.length > 1) {
    previewInterval = setInterval(() => {
      currentPreviewSlide.value = currentPreviewSlide.value === 0 
        ? sliderConfig.value.slides.length - 1 
        : currentPreviewSlide.value - 1
    }, (sliderConfig.value.autoplay_delay || 5) * 1000)
  }
}

const stopPreviewRotation = () => {
  if (previewInterval) {
    clearInterval(previewInterval)
    previewInterval = null
  }
}

// Watch for autoplay changes
watch(() => [sliderConfig.value.autoplay_enabled, sliderConfig.value.autoplay_delay, sliderConfig.value.slides.length], () => {
  if (sliderConfig.value.autoplay_enabled && sliderConfig.value.slides.length > 1) {
    startPreviewRotation()
  } else {
    stopPreviewRotation()
  }
}, { immediate: true })

// Cleanup on unmount
onUnmounted(() => {
  stopPreviewRotation()
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

    // Prepare widget data for creation
    const widgetData = {
      title: formData.value.title,
      description: formData.value.description,
      type: formData.value.type as any, // Type casting for compatibility
      status: formData.value.status as any, // Type casting for compatibility
      page: formData.value.page as any, // Type casting for compatibility
      config: sliderConfig.value
    }
    
    // Create new widget
    const createdWidget = await createWidget(widgetData)
    
    if (createdWidget) {
      // Redirect to widgets list page after successful creation
      await router.push('/admin/content/banners')
    }
    
  } catch (error) {
    console.error('خطا در ذخیره ابزارک:', error)
  } finally {
    isSaving.value = false
  }
}

// Initialize config from widget
onMounted(async () => {
  clearError()
  if (widgetId) {
    await fetchWidget(widgetId)
  }
  initializeFormData()
  
  // Only copy specific config fields, don't overwrite defaults
  if (fetchedWidget.value?.config) {
    const config = fetchedWidget.value.config as SliderConfig
    if (config.slides) sliderConfig.value.slides = config.slides
    if (config.height) sliderConfig.value.height = config.height
    if (config.slider_width) sliderConfig.value.slider_width = config.slider_width
    if (config.bg_enabled !== undefined) sliderConfig.value.bg_enabled = config.bg_enabled
    if (config.bg_color) sliderConfig.value.bg_color = config.bg_color
    if (config.wide_bg !== undefined) sliderConfig.value.wide_bg = config.wide_bg
    if (config.autoplay_enabled !== undefined) sliderConfig.value.autoplay_enabled = config.autoplay_enabled
    if (config.autoplay_delay) sliderConfig.value.autoplay_delay = config.autoplay_delay
    if (config.loop_enabled !== undefined) sliderConfig.value.loop_enabled = config.loop_enabled
    if (config.show_navigation !== undefined) sliderConfig.value.show_navigation = config.show_navigation
    if (config.show_pagination !== undefined) sliderConfig.value.show_pagination = config.show_pagination
    if (config.easy_load_enabled !== undefined) sliderConfig.value.easy_load_enabled = config.easy_load_enabled
    // Mobile specific settings
    if (config.mobile_autoplay_enabled !== undefined) sliderConfig.value.mobile_autoplay_enabled = config.mobile_autoplay_enabled
    if (config.mobile_autoplay_delay) sliderConfig.value.mobile_autoplay_delay = config.mobile_autoplay_delay
    if (config.mobile_loop_enabled !== undefined) sliderConfig.value.mobile_loop_enabled = config.mobile_loop_enabled
    if (config.mobile_show_navigation !== undefined) sliderConfig.value.mobile_show_navigation = config.mobile_show_navigation
    if (config.mobile_show_pagination !== undefined) sliderConfig.value.mobile_show_pagination = config.mobile_show_pagination
    if (config.mobile_vertical_display !== undefined) sliderConfig.value.mobile_vertical_display = config.mobile_vertical_display
  }
})
</script>
