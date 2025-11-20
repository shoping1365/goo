<template>
  <div class="w-full">
    <!-- Title Section -->
    <div class="bg-white w-full shadow-sm border-2 border-blue-200 rounded-lg mt-4">
      <div class="flex justify-between items-center px-8 py-6">
        <h1 class="text-2xl font-bold text-gray-800">ویرایش {{ widget?.title || 'ابزارک' }}</h1>
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

    <!-- تنظیمات بنر -->
    <DeviceTabs
      ref="deviceTabsRef"
      :banner-config="bannerConfig"
      :current-preview-banner="currentPreviewBanner"
      :open-add-banner-modal="openAddBannerModal"
      :edit-banner="editBanner"
      :remove-banner="removeBanner"
      :get-banner-classes="() => ''"
      :get-banner-width="() => ''"
    >
      <!-- محتوای تب دسکتاپ -->
      <template #desktop-content>
        <div class="bg-white rounded-xl shadow p-8 mt-6 mx-4">
          <div class="flex justify-between items-center mb-6">
            <h3 class="text-lg font-bold text-gray-700">تنظیمات دسکتاپ</h3>
          </div>
          
          <div class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-6">
        <!-- پس‌زمینه فعال -->
        <div>
          <label class="block mb-2 text-sm font-medium text-gray-700">پس‌زمینه فعال</label>
          <select
            v-model="bannerConfig.bg_enabled"
            class="w-full border border-gray-300 rounded-md px-2 py-1.5 focus:outline-none focus:ring-2 focus:ring-purple-400 text-sm"
          >
            <option :value="true">فعال</option>
            <option :value="false">غیرفعال</option>
          </select>
        </div>

        <!-- عریض پس‌زمینه -->
        <div v-if="bannerConfig.bg_enabled">
          <label class="block mb-2 text-sm font-medium text-gray-700">عریض پس‌زمینه</label>
          <select
            v-model="bannerConfig.wide_bg"
            class="w-full border border-gray-300 rounded-md px-2 py-1.5 focus:outline-none focus:ring-2 focus:ring-purple-400 text-sm"
          >
            <option :value="true">بله</option>
            <option :value="false">خیر</option>
          </select>
        </div>

        <!-- رنگ پس‌زمینه -->
        <div v-if="bannerConfig.bg_enabled">
          <label class="block mb-2 text-sm font-medium text-gray-700">رنگ پس‌زمینه</label>
          <input
            v-model="bannerConfig.bg_color"
            type="color"
            class="w-full h-8 border border-gray-300 rounded-md"
          />
        </div>

        <!-- ارتفاع بنر -->
        <div>
          <label class="block mb-2 text-sm font-medium text-gray-700">ارتفاع بنر (پیکسل)</label>
          <select
            v-model="bannerConfig.height"
            class="w-full border border-gray-300 rounded-md px-2 py-1.5 focus:outline-none focus:ring-2 focus:ring-purple-400 text-sm"
          >
            <option :value="300">300 پیکسل</option>
            <option :value="400">400 پیکسل</option>
            <option :value="500">500 پیکسل</option>
            <option :value="600">600 پیکسل</option>
            <option :value="700">700 پیکسل</option>
          </select>
        </div>

        <!-- عرض بنر -->
        <div>
          <label class="block mb-2 text-sm font-medium text-gray-700">عرض بنر</label>
          <select
            v-model="bannerConfig.banner_width"
            class="w-full border border-gray-300 rounded-md px-2 py-1.5 focus:outline-none focus:ring-2 focus:ring-purple-400 text-sm"
          >
            <option :value="800">تمام عرض</option>
            <option :value="600">وسط</option>
          </select>
        </div>

        <!-- عرض مرکز -->
        <div>
          <label class="block mb-2 text-sm font-medium text-gray-700">عرض مرکز (پیکسل)</label>
          <input
            v-model.number="bannerConfig.center_width"
            type="number"
            min="400"
            max="1400"
            step="50"
            class="w-full border border-gray-300 rounded-md px-2 py-1.5 focus:outline-none focus:ring-2 focus:ring-purple-400 text-sm"
            placeholder="1000"
          />
          <p class="text-xs text-gray-500 mt-1">عرض بنر در وسط صفحه (400-1400 پیکسل)</p>
        </div>

        <!-- نسبت بنر 1 -->
        <div>
          <label class="block mb-2 text-sm font-medium text-gray-700">نسبت بنر 1</label>
          <select
            v-model="bannerConfig.banner1_ratio"
            class="w-full border border-gray-300 rounded-md px-2 py-1.5 focus:outline-none focus:ring-2 focus:ring-purple-400 text-sm"
            @change="updateBannerRatios('banner1')"
          >
            <option :value="20">20%</option>
            <option :value="25">25%</option>
            <option :value="30">30%</option>
            <option :value="40">40%</option>
            <option :value="50">50%</option>
          </select>
        </div>

        <!-- نسبت بنر 2 -->
        <div>
          <label class="block mb-2 text-sm font-medium text-gray-700">نسبت بنر 2</label>
          <select
            v-model="bannerConfig.banner2_ratio"
            class="w-full border border-gray-300 rounded-md px-2 py-1.5 focus:outline-none focus:ring-2 focus:ring-purple-400 text-sm"
            @change="updateBannerRatios('banner2')"
          >
            <option :value="15">15%</option>
            <option :value="20">20%</option>
            <option :value="25">25%</option>
            <option :value="30">30%</option>
            <option :value="40">40%</option>
          </select>
        </div>

        <!-- نسبت بنر 3 -->
        <div>
          <label class="block mb-2 text-sm font-medium text-gray-700">نسبت بنر 3</label>
          <select
            v-model="bannerConfig.banner3_ratio"
            class="w-full border border-gray-300 rounded-md px-2 py-1.5 focus:outline-none focus:ring-2 focus:ring-purple-400 text-sm"
            @change="updateBannerRatios('banner3')"
          >
            <option :value="15">15%</option>
            <option :value="20">20%</option>
            <option :value="25">25%</option>
            <option :value="30">30%</option>
            <option :value="40">40%</option>
          </select>
        </div>

        <!-- نسبت بنر 4 -->
        <div>
          <label class="block mb-2 text-sm font-medium text-gray-700">نسبت بنر 4</label>
          <select
            v-model="bannerConfig.banner4_ratio"
            class="w-full border border-gray-300 rounded-md px-2 py-1.5 focus:outline-none focus:ring-2 focus:ring-purple-400 text-sm"
            @change="updateBannerRatios('banner4')"
          >
            <option :value="15">15%</option>
            <option :value="20">20%</option>
            <option :value="25">25%</option>
            <option :value="30">30%</option>
            <option :value="40">40%</option>
          </select>
        </div>

        <!-- پدینگ بالا -->
        <div>
          <label class="block mb-2 text-sm font-medium text-gray-700">پدینگ بالا (px)</label>
          <input
            type="number"
            :value="bannerConfig.padding_top !== undefined ? bannerConfig.padding_top : ''"
            min="0"
            max="100"
            class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
            placeholder="0"
            @input="(e: Event) => bannerConfig.padding_top = (e.target as HTMLInputElement).value === '' ? undefined : Number((e.target as HTMLInputElement).value)"
          />
        </div>

        <!-- پدینگ پایین -->
        <div>
          <label class="block mb-2 text-sm font-medium text-gray-700">پدینگ پایین (px)</label>
          <input
            type="number"
            :value="bannerConfig.padding_bottom !== undefined ? bannerConfig.padding_bottom : ''"
            min="0"
            max="100"
            class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
            placeholder="0"
            @input="(e: Event) => bannerConfig.padding_bottom = (e.target as HTMLInputElement).value === '' ? undefined : Number((e.target as HTMLInputElement).value)"
          />
        </div>

        <!-- مارجین راست -->
        <div>
          <label class="block mb-2 text-sm font-medium text-gray-700">مارجین راست (px)</label>
          <input
            type="number"
            :value="bannerConfig.margin_right !== undefined ? bannerConfig.margin_right : ''"
            min="0"
            max="100"
            class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
            placeholder="0"
            @input="(e: Event) => bannerConfig.margin_right = (e.target as HTMLInputElement).value === '' ? undefined : Number((e.target as HTMLInputElement).value)"
          />
        </div>

        <!-- مارجین چپ -->
        <div>
          <label class="block mb-2 text-sm font-medium text-gray-700">مارجین چپ (px)</label>
          <input
            type="number"
            :value="bannerConfig.margin_left !== undefined ? bannerConfig.margin_left : ''"
            min="0"
            max="100"
            class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
            placeholder="0"
            @input="(e: Event) => bannerConfig.margin_left = (e.target as HTMLInputElement).value === '' ? undefined : Number((e.target as HTMLInputElement).value)"
          />
        </div>

        <!-- نمایش عمودی در موبایل -->
        <div>
          <label class="block mb-2 text-sm font-medium text-gray-700">نمایش عمودی موبایل</label>
          <select
            v-model="bannerConfig.mobile_vertical_display"
            class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
          >
            <option :value="true">فعال</option>
            <option :value="false">غیرفعال</option>
          </select>
        </div>

        <!-- تنظیمات عکس موبایل -->
        <div class="col-span-full">
          <div class="bg-gray-50 rounded-lg p-6">
            <div class="flex items-center gap-6 mb-4">
              <label class="flex items-center gap-2">
                <input
                  v-model="bannerConfig.mobile_image_mode"
                  type="radio"
                  value="auto"
                  class="w-4 h-4 text-blue-600"
                />
                <span class="text-sm">برش خودکار از عکس دسکتاپ</span>
              </label>
              <label class="flex items-center gap-2">
                <input
                  v-model="bannerConfig.mobile_image_mode"
                  type="radio"
                  value="separate"
                  class="w-4 h-4 text-blue-600"
                />
                <span class="text-sm">عکس جداگانه برای موبایل</span>
              </label>
              <button
                type="button"
                class="bg-green-500 text-white px-3 py-2 rounded-md text-sm hover:bg-green-600 transition-colors"
                @click="applyMobileCrop"
              >
                اعمال برش موبایل
              </button>
              <span v-if="bannerConfig.mobile_cropped_image" class="text-xs text-green-600">
                ✓ برش اعمال شده
              </span>
              <span v-else class="text-xs text-gray-500">
                برش اعمال نشده
              </span>
            </div>
            
            <!-- تنظیمات برش خودکار -->
            <div v-if="bannerConfig.mobile_image_mode === 'auto'" class="space-y-3">
              <div class="grid grid-cols-2 gap-6">
                <div>
                  <label class="block mb-2 text-sm font-medium text-gray-700">عرض موبایل (پیکسل)</label>
                  <input
                    v-model="bannerConfig.mobile_crop_width"
                    type="number"
                    min="200"
                    max="800"
                    class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
                    placeholder="375"
                  />
                </div>
                <div>
                  <label class="block mb-2 text-sm font-medium text-gray-700">ارتفاع موبایل (پیکسل)</label>
                  <input
                    v-model="bannerConfig.mobile_crop_height"
                    type="number"
                    min="100"
                    max="600"
                    class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
                    placeholder="150"
                  />
                </div>
              </div>
            </div>
            
            <!-- آپلود عکس جداگانه -->
            <div v-if="bannerConfig.mobile_image_mode === 'separate'" class="space-y-3">
            </div>
          </div>
        </div>
      </div>
    </div>
    </template>
      
      <!-- محتوای تب موبایل -->
      <template #mobile-content>
        <!-- محتوای موبایل توسط کامپوننت DeviceTabs مدیریت می‌شود -->
      </template>
    </DeviceTabs>

    <!-- پیش‌نمایش زنده دسکتاپ -->
    <div v-if="deviceTabsRef?.activeTab === 'desktop'" class="bg-white rounded-xl shadow p-8 mt-6 mx-4">
      <h3 class="text-lg font-bold text-gray-700 mb-6">پیش‌نمایش زنده</h3>
       
      <!-- پیش‌نمایش بنرهای چهارتایی -->
      <div class="bg-gray-50 rounded-lg p-6 border-2 border-dashed border-gray-300">
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
          <!-- بنر اول -->
          <div class="relative overflow-hidden rounded-lg shadow-lg">
              <div 
                class="relative"
                :style="{
                  height: `${bannerConfig.height}px`,
                  backgroundColor: bannerConfig.bg_enabled ? bannerConfig.bg_color : 'transparent'
                }"
              >
                <img
                  v-if="bannerConfig.banners && bannerConfig.banners.length > 0"
                  :src="bannerConfig.banners[0].image"
                  :alt="bannerConfig.banners[0].title"
                  class="w-full h-full object-cover"
                />
                <div v-else class="flex items-center justify-center h-full text-gray-400">
                  <div class="text-center">
                    <svg class="w-12 h-12 mx-auto mb-2 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                    </svg>
                    <p class="text-sm">بنر اول</p>
                  </div>
                </div>
                <div
                  v-if="bannerConfig.show_title || bannerConfig.show_description"
                  class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-6"
                >
                  <h4
                    v-if="bannerConfig.show_title && bannerConfig.banners && bannerConfig.banners[0] && bannerConfig.banners[0].title"
                    class="text-white text-lg font-bold mb-2"
                  >
                    {{ bannerConfig.banners[0].title }}
                  </h4>
                  <p
                    v-if="bannerConfig.show_description && bannerConfig.banners && bannerConfig.banners[0] && bannerConfig.banners[0].description"
                    class="text-white/90 text-sm"
                  >
                    {{ bannerConfig.banners[0].description }}
                  </p>
                </div>
              </div>
            </div>

          <!-- بنر دوم -->
          <div class="relative overflow-hidden rounded-lg shadow-lg">
              <div 
                class="relative"
                :style="{
                  height: `${bannerConfig.height}px`,
                  backgroundColor: bannerConfig.bg_enabled ? bannerConfig.bg_color : 'transparent'
                }"
              >
                <img
                  v-if="bannerConfig.banners && bannerConfig.banners.length > 1"
                  :src="bannerConfig.banners[1].image"
                  :alt="bannerConfig.banners[1].title"
                  class="w-full h-full object-cover"
                />
                <div v-else class="flex items-center justify-center h-full text-gray-400">
                  <div class="text-center">
                    <svg class="w-12 h-12 mx-auto mb-2 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                    </svg>
                    <p class="text-sm">بنر دوم</p>
                  </div>
                </div>
                <div
                  v-if="bannerConfig.show_title || bannerConfig.show_description"
                  class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-6"
                >
                  <h4
                    v-if="bannerConfig.show_title && bannerConfig.banners && bannerConfig.banners[1] && bannerConfig.banners[1].title"
                    class="text-white text-lg font-bold mb-2"
                  >
                    {{ bannerConfig.banners[1].title }}
                  </h4>
                  <p
                    v-if="bannerConfig.show_description && bannerConfig.banners && bannerConfig.banners[1] && bannerConfig.banners[1].description"
                    class="text-white/90 text-sm"
                  >
                    {{ bannerConfig.banners[1].description }}
                  </p>
                </div>
              </div>
            </div>

          <!-- بنر سوم -->
          <div class="relative overflow-hidden rounded-lg shadow-lg">
              <div 
                class="relative"
                :style="{
                  height: `${bannerConfig.height}px`,
                  backgroundColor: bannerConfig.bg_enabled ? bannerConfig.bg_color : 'transparent'
                }"
              >
                <img
                  v-if="bannerConfig.banners && bannerConfig.banners.length > 2"
                  :src="bannerConfig.banners[2].image"
                  :alt="bannerConfig.banners[2].title"
                  class="w-full h-full object-cover"
                />
                <div v-else class="flex items-center justify-center h-full text-gray-400">
                  <div class="text-center">
                    <svg class="w-12 h-12 mx-auto mb-2 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                    </svg>
                    <p class="text-sm">بنر سوم</p>
                  </div>
                </div>
                <div
                  v-if="bannerConfig.show_title || bannerConfig.show_description"
                  class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-6"
                >
                  <h4
                    v-if="bannerConfig.show_title && bannerConfig.banners && bannerConfig.banners[2] && bannerConfig.banners[2].title"
                    class="text-white text-lg font-bold mb-2"
                  >
                    {{ bannerConfig.banners[2].title }}
                  </h4>
                  <p
                    v-if="bannerConfig.show_description && bannerConfig.banners && bannerConfig.banners[2] && bannerConfig.banners[2].description"
                    class="text-white/90 text-sm"
                  >
                    {{ bannerConfig.banners[2].description }}
                  </p>
                </div>
              </div>
            </div>

          <!-- بنر چهارم -->
          <div class="relative overflow-hidden rounded-lg shadow-lg">
              <div 
                class="relative"
                :style="{
                  height: `${bannerConfig.height}px`,
                  backgroundColor: bannerConfig.bg_enabled ? bannerConfig.bg_color : 'transparent'
                }"
              >
                <img
                  v-if="bannerConfig.banners && bannerConfig.banners.length > 3"
                  :src="bannerConfig.banners[3].image"
                  :alt="bannerConfig.banners[3].title"
                  class="w-full h-full object-cover"
                />
                <div v-else class="flex items-center justify-center h-full text-gray-400">
                  <div class="text-center">
                    <svg class="w-12 h-12 mx-auto mb-2 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                    </svg>
                    <p class="text-sm">بنر چهارم</p>
                  </div>
                </div>
                <div
                  v-if="bannerConfig.show_title || bannerConfig.show_description"
                  class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-6"
                >
                  <h4
                    v-if="bannerConfig.show_title && bannerConfig.banners && bannerConfig.banners[3] && bannerConfig.banners[3].title"
                    class="text-white text-lg font-bold mb-2"
                  >
                    {{ bannerConfig.banners[3].title }}
                  </h4>
                  <p
                    v-if="bannerConfig.show_description && bannerConfig.banners && bannerConfig.banners[3] && bannerConfig.banners[3].description"
                    class="text-white/90 text-sm"
                  >
                    {{ bannerConfig.banners[3].description }}
                  </p>
                </div>
              </div>
            </div>
          </div>
          
          <!-- اطلاعات تنظیمات -->
          <div class="mt-4 text-sm text-gray-600 text-center">
            <p>ارتفاع: {{ bannerConfig.height }}px | 
                عرض: {{ bannerConfig.banner_width === 800 ? 'تمام عرض' : 'وسط صفحه' }} | 
                پس‌زمینه: {{ bannerConfig.bg_enabled ? 'فعال' : 'غیرفعال' }} | 
                نسبت‌ها: {{ bannerConfig.banner1_ratio }}% / {{ bannerConfig.banner2_ratio }}% / {{ bannerConfig.banner3_ratio }}% / {{ bannerConfig.banner4_ratio }}%</p>
          </div>
          </div>
    </div>

    <!-- نمایش و مدیریت بنرها -->
    <div class="bg-white rounded-xl shadow p-8 mt-6 mx-4">
      <div class="flex justify-between items-center mb-6">
        <h3 class="text-lg font-bold text-gray-700">بنرها</h3>
        <button
          class="bg-cyan-500 text-white px-4 py-2 rounded-md font-bold hover:bg-cyan-600 transition-colors"
          @click="openAddBannerModal"
        >
          افزودن بنر جدید
        </button>
      </div>

      <div v-if="bannerConfig.banners.length === 0" class="text-gray-400 text-center py-8">
        چیزی برای نمایش وجود ندارد!
      </div>

      <div v-else class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div
          v-for="(banner, idx) in bannerConfig.banners"
          :key="idx"
          class="flex items-center gap-6 p-3 border rounded-lg bg-gray-50"
        >
          <img
            :src="banner.image"
            alt="بنر"
            class="w-32 h-20 object-cover rounded border"
          />
          <div class="flex flex-col flex-1">
            <div class="font-bold text-sm text-gray-700 mb-1">{{ banner.title }}</div>
            <div v-if="banner.description" class="text-xs text-gray-600 mb-1">{{ banner.description }}</div>
            <div v-if="banner.link" class="text-xs text-blue-600 break-all">{{ banner.link }}</div>
          </div>
          <div class="flex gap-2">
            <button
              class="text-blue-500 hover:text-blue-700 p-1"
              @click="editBanner(idx)"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536M9 11l6 6M3 17v4h4l10.293-10.293a1 1 0 00-1.414-1.414L3 17z"></path>
              </svg>
            </button>
            <button
              class="text-red-500 hover:text-red-700 p-1"
              @click="removeBanner(idx)"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
              </svg>
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Banner Modal Component -->
    <div v-if="showBannerModal" class="fixed inset-0 flex items-center justify-center z-50 p-6" @click="closeBannerModal">
      <div class="bg-white rounded-lg border-4 border-blue-300 max-w-2xl w-full max-h-[90vh] overflow-y-auto" style="box-shadow: 0 0 20px 8px rgba(173, 216, 230, 0.3), 0 0 40px 12px rgba(173, 216, 230, 0.15);" @click.stop>
        <div class="p-6">
          <div class="flex justify-between items-center mb-6">
            <h3 class="text-xl font-bold text-gray-800" style="font-family: 'Yekan', sans-serif;">
              {{ editingBannerIndex !== null ? 'ویرایش بنر' : 'افزودن بنر جدید' }}
            </h3>
            <button class="text-gray-400 hover:text-gray-600" @click="closeBannerModal">
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
              </svg>
            </button>
          </div>

          <form class="space-y-4" @submit.prevent="handleBannerSave">
            <!-- عنوان بنر -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1" style="font-family: 'Yekan', sans-serif;">عنوان بنر *</label>
              <input
                v-model="editingBanner.title"
                type="text"
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
                placeholder="عنوان بنر را وارد کنید"
                style="font-family: 'Yekan', sans-serif;"
                required
              />
            </div>

            <!-- توضیحات -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1" style="font-family: 'Yekan', sans-serif;">توضیحات</label>
              <textarea
                v-model="editingBanner.description"
                rows="3"
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
                placeholder="توضیحات بنر را وارد کنید"
                style="font-family: 'Yekan', sans-serif;"
              ></textarea>
            </div>

            <!-- لینک -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1" style="font-family: 'Yekan', sans-serif;">لینک</label>
              <input
                v-model="editingBanner.link"
                type="url"
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
                placeholder="https://example.com"
                style="font-family: 'Yekan', sans-serif;"
              />
            </div>

            <!-- تصویر -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1" style="font-family: 'Yekan', sans-serif;">تصویر بنر *</label>
              <div class="flex items-center gap-6">
                <div v-if="editingBanner.image" class="flex-1">
                  <img :src="editingBanner.image" alt="بنر" class="w-full h-32 object-cover rounded border" />
                </div>
                <div v-else class="flex-1 border-2 border-dashed border-gray-300 rounded p-8 text-center">
                  <svg class="w-12 h-12 mx-auto mb-2 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                  </svg>
                  <p class="text-gray-500 text-sm">تصویری انتخاب نشده</p>
                </div>
                <div class="flex flex-col gap-2">
                  <button
                    type="button"
                    class="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600 transition-colors"
                    @click="openMediaLibrary"
                  >
                    انتخاب تصویر
                  </button>
                  <button
                    v-if="editingBanner.image"
                    type="button"
                    class="bg-red-500 text-white px-4 py-2 rounded-md hover:bg-red-600 transition-colors"
                    @click="removeImage"
                  >
                    حذف تصویر
                  </button>
                </div>
              </div>
            </div>

            <!-- نمایش عنوان -->
            <div class="flex items-center gap-2">
              <input
                id="showTitle"
                v-model="showTitleInBanner"
                type="checkbox"
                class="w-4 h-4 text-blue-600 bg-blue-100 border-blue-300 rounded focus:ring-blue-500 focus:ring-2"
              />
              <label for="showTitle" class="text-sm font-medium text-gray-700">نمایش عنوان روی بنر</label>
            </div>

            <!-- دکمه‌های عملیات -->
            <div class="flex justify-end gap-6 pt-4">
              <button
                type="button"
                class="px-6 py-2 border border-gray-300 rounded-md text-gray-700 hover:bg-gray-50 transition-colors"
                @click="closeBannerModal"
              >
                انصراف
              </button>
              <button
                type="submit"
                :disabled="!editingBanner.title || !editingBanner.image"
                class="px-6 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
              >
                {{ editingBannerIndex !== null ? 'ویرایش' : 'افزودن' }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>

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
import { WIDGET_TYPE_LABELS } from '~/types/widget'
import type { Widget, SlideItem } from '~/types/widget'
import TemplateButton from '~/components/common/TemplateButton.vue'
import MediaLibraryModal from '~/components/media/MediaLibraryModal.vue'
import DeviceTabs from './components/DeviceTabs.vue'
import { ref, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'
import { useWidget } from '~/composables/useWidget'

// تعریف definePageMeta و useHead برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
declare const useHead: (head: { title?: string }) => void

definePageMeta({ layout: 'admin-main', middleware: 'admin' })
useHead({ title: 'ویرایش بنر چهارتایی - پنل ادمین' })

// Route params
const route = useRoute()
const widgetId = parseInt(route.params.id as string)

// Composables
const { fetchWidget, updateWidget, loading, error, clearError, widget } = useWidget()

// Notification helpers
const showSuccess = (message: string) => {
  console.log('Success:', message)
  // You can add toast notification logic here
}

const showError = (message: string) => {
  console.error('Error:', message)
  // You can add toast notification logic here
}

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
const showBannerModal = ref(false)
const showMediaLibrary = ref(false)
const editingBannerIndex = ref<number | null>(null)
const deviceTabsRef = ref()
const currentPreviewBanner = ref(0)

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
const editingBanner = ref<SlideItem>({
  title: '',
  description: '',
  image: '',
  link: '',
  order: 1,
  status: 'active'
})

// Title display control
const showTitleInBanner = ref(true)

// Form data
const formData = ref({
  title: '',
  description: '',
  type: 'quad-banner',
  status: 'active',
  page: 'home'
})

// Initialize form data when widget is available
const initializeFormData = () => {
  if (widget.value) {
    console.log('Widget data:', widget.value) // Debug log
    formData.value = {
      title: widget.value.title || '',
      description: widget.value.description || '',
      type: widget.value.type || 'quad-banner',
      status: widget.value.status || 'active',
      page: widget.value.page || 'home'
    }
    console.log('Form data initialized:', formData.value) // Debug log
  }
}

// Watch for widget changes
watch(widget, (newWidget) => {
  if (newWidget) {
    initializeFormData()
  }
}, { immediate: true })


// Banner config
const bannerConfig = ref({
  height: 400,
  banners: [],
  bg_enabled: false,
  bg_color: '#ffffff',
  wide_bg: false,
  banner_width: 800,
  center_width: 1000,
  banner1_ratio: 25,
  banner2_ratio: 25,
  banner3_ratio: 25,
  banner4_ratio: 25,
  padding_top: 0,
  padding_bottom: 0,
  margin_left: 0,
  margin_right: 0,
  show_title: true,
  show_description: true,
  // Mobile specific settings
  mobile_vertical_display: false,
  mobile_height: 150,
  mobile_banner_width: 500, // عرض بنر موبایل
  mobile_image_mode: 'auto', // 'auto' or 'separate'
  mobile_crop_width: 375, // عرض پیش‌فرض برای برش موبایل
  mobile_crop_height: 150, // ارتفاع پیش‌فرض برای برش موبایل
  mobile_cropped_image: '', // URL عکس برش داده شده برای موبایل
  mobile_image: '', // URL for separate mobile image
  // Mobile padding and margin settings
  mobile_padding_top: undefined,
  mobile_padding_bottom: undefined
})

// Apply mobile crop function
const applyMobileCrop = async () => {
  if (!bannerConfig.value.banners.length) {
    alert('ابتدا یک بنر اضافه کنید')
    return
  }
  
  const banner = bannerConfig.value.banners[0]
  if (!banner.image) {
    alert('عکس بنر انتخاب نشده است')
    return
  }
  
  try {
    // برش عکس موبایل بر اساس ابعاد مشخص شده
    const cropWidth = bannerConfig.value.mobile_crop_width || 375
    const cropHeight = bannerConfig.value.mobile_crop_height || 150
    
    // فراخوانی API برش عکس موبایل در بک‌اند
    try {
      const response = await $fetch('/api/media/image-crop', {
        method: 'POST',
        body: {
          src: banner.image,
          crop_width: cropWidth,
          crop_height: cropHeight,
          mode: 'cover',
          device: 'mobile',
          quality: 85
        }
      }) as any
      
      if (response.success) {
        bannerConfig.value.mobile_cropped_image = response.data.cropped_url
        alert('برش موبایل با موفقیت اعمال شد')
      } else {
        throw new Error('API response failed')
      }
    } catch (error) {
      console.error('Backend crop API error:', error)
      alert('خطا در برش عکس موبایل')
      return
    }
  } catch (error) {
    console.error('خطا در اعمال برش موبایل:', error)
    alert('خطا در اعمال برش موبایل')
  }
}

// Methods for managing banners
const openAddBannerModal = () => {
  editingBannerIndex.value = null
  editingBanner.value = {
    title: '',
    description: '',
    image: '',
    link: '',
    order: bannerConfig.value.banners.length + 1,
    status: 'active'
  }
  showTitleInBanner.value = true // Reset to default
  showBannerModal.value = true
}

const closeBannerModal = () => {
  showBannerModal.value = false
  editingBannerIndex.value = null
}

const editBanner = (index: number) => {
  editingBannerIndex.value = index
  const banner = bannerConfig.value.banners[index]
  editingBanner.value = {
    title: banner.title || '',
    description: banner.description || '',
    image: banner.image || '',
    link: banner.link || '',
    order: banner.order || index + 1,
    status: banner.status || 'active'
  }
  // Set showTitleInBanner based on existing banner data or default to true
  showTitleInBanner.value = banner.showTitle !== undefined ? banner.showTitle : true
  showBannerModal.value = true
}

const removeBanner = (index: number) => {
  bannerConfig.value.banners.splice(index, 1)
  // Re-order remaining banners
  bannerConfig.value.banners.forEach((banner, idx) => {
    banner.order = idx + 1
  })
}

// Handle banner save from modal component
const handleBannerSave = () => {
  const banner: SlideItem = {
    title: editingBanner.value.title?.trim() || '',
    description: editingBanner.value.description?.trim() || '',
    link: editingBanner.value.link?.trim() || '',
    image: editingBanner.value.image || '',
    order: editingBannerIndex.value !== null && editingBannerIndex.value >= 0 ? bannerConfig.value.banners[editingBannerIndex.value].order : bannerConfig.value.banners.length + 1,
    status: 'active',
    showTitle: showTitleInBanner.value
  }

  if (editingBannerIndex.value !== null && editingBannerIndex.value >= 0) {
    // ویرایش بنر موجود
    bannerConfig.value.banners[editingBannerIndex.value] = banner
  } else {
    // افزودن بنر جدید
    bannerConfig.value.banners.push(banner)
  }

  closeBannerModal()
}

// Media library functions
const openMediaLibrary = () => {
  showMediaLibrary.value = true
}

const onSelectFromLibrary = (files: any[]) => {
  if (files && files.length > 0) {
    const file = files[0]
    editingBanner.value.image = file.url
    // Clear any previous image error
    // Note: This will be handled by the child component's reactive system
  }
  showMediaLibrary.value = false
}

// Image functions
const removeImage = () => {
  editingBanner.value.image = ''
}

// Banner ratio management functions
const updateBannerRatios = (changedBanner: 'banner1' | 'banner2' | 'banner3' | 'banner4') => {
  const changedValue = bannerConfig.value[`${changedBanner}_ratio`]
  const remaining = 100 - changedValue
  
  if (remaining < 0) {
    // اگر مقدار انتخاب شده بیشتر از 100 است، آن را به 100 محدود کن
    bannerConfig.value[`${changedBanner}_ratio`] = 100
    bannerConfig.value.banner1_ratio = 0
    bannerConfig.value.banner2_ratio = 0
    bannerConfig.value.banner3_ratio = 0
    bannerConfig.value.banner4_ratio = 0
    return
  }
  
  // نسبت‌های باقیمانده را بین سه بنر دیگر تقسیم کن
  const otherBanners = ['banner1', 'banner2', 'banner3', 'banner4'].filter(b => b !== changedBanner)
  
  if (otherBanners.length === 3) {
    const [bannerA, bannerB, bannerC] = otherBanners
    const currentA = bannerConfig.value[`${bannerA}_ratio`]
    const currentB = bannerConfig.value[`${bannerB}_ratio`]
    const currentC = bannerConfig.value[`${bannerC}_ratio`]
    const totalCurrent = currentA + currentB + currentC
    
    if (totalCurrent > 0) {
      // نسبت‌ها را بر اساس نسبت فعلی تقسیم کن
      const ratioA = currentA / totalCurrent
      const ratioB = currentB / totalCurrent
      const ratioC = currentC / totalCurrent
      
      bannerConfig.value[`${bannerA}_ratio`] = Math.round(remaining * ratioA)
      bannerConfig.value[`${bannerB}_ratio`] = Math.round(remaining * ratioB)
      bannerConfig.value[`${bannerC}_ratio`] = remaining - bannerConfig.value[`${bannerA}_ratio`] - bannerConfig.value[`${bannerB}_ratio`]
    } else {
      // اگر هر سه صفر هستند، به طور مساوی تقسیم کن
      const equalShare = Math.floor(remaining / 3)
      bannerConfig.value[`${bannerA}_ratio`] = equalShare
      bannerConfig.value[`${bannerB}_ratio`] = equalShare
      bannerConfig.value[`${bannerC}_ratio`] = remaining - equalShare - equalShare
    }
  }
}


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

    // اعمال برش خودکار موبایل قبل از ذخیره
    await applyAutoMobileCrop()

    // Prepare widget data for update
    const widgetData = {
      title: formData.value.title,
      description: formData.value.description,
      type: formData.value.type as any, // Type casting for compatibility
      status: formData.value.status as any, // Type casting for compatibility
      page: formData.value.page as any, // Type casting for compatibility
      config: bannerConfig.value
    }
    
    // Update existing widget
    const updatedWidget = await updateWidget(widgetId, widgetData)
    
    if (updatedWidget) {
      showSuccess('ابزارک با موفقیت به‌روزرسانی شد!')
    }
    
  } catch (error) {
    console.error('خطا در ذخیره ابزارک:', error)
  } finally {
    isSaving.value = false
  }
}

// اعمال برش خودکار موبایل هنگام ذخیره
const applyAutoMobileCrop = async () => {
  // فقط اگر حالت برش خودکار انتخاب شده باشد
  if (bannerConfig.value.mobile_image_mode !== 'auto') {
    return
  }
  
  if (!bannerConfig.value.banners.length) {
    return
  }
  
  const banner = bannerConfig.value.banners[0]
  if (!banner.image) {
    return
  }
  
  // همیشه برش جدید اعمال کن (حتی اگر قبلاً برش شده باشد)
  
  try {
    // برش عکس موبایل بر اساس ابعاد مشخص شده
    const cropWidth = bannerConfig.value.mobile_crop_width || 375
    const cropHeight = bannerConfig.value.mobile_crop_height || 150
    
    // فراخوانی API برش عکس موبایل در بک‌اند
    try {
      const response = await $fetch('/api/media/image-crop', {
        method: 'POST',
        body: {
          src: banner.image,
          crop_width: cropWidth,
          crop_height: cropHeight,
          mode: 'cover',
          device: 'mobile',
          quality: 85
        }
      }) as any
      
      if (response.success) {
        bannerConfig.value.mobile_cropped_image = response.data.cropped_url
      } else {
        throw new Error('API response failed')
      }
    } catch (error) {
      console.error('Backend crop API error:', error)
      // اگر API کراپ کار نکرد، هیچ fallback استفاده نکن
      return
    }
    
    console.log('برش خودکار موبایل اعمال شد')
  } catch (error) {
    console.error('خطا در اعمال برش خودکار موبایل:', error)
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
  if (widget.value?.config) {
    const config = widget.value.config as any
    if (config.banners) bannerConfig.value.banners = config.banners
    else if (config.slides) bannerConfig.value.banners = config.slides
    if (config.height) bannerConfig.value.height = config.height
    if (config.banner_width) bannerConfig.value.banner_width = config.banner_width
    else if (config.slider_width) bannerConfig.value.banner_width = config.slider_width
    if (config.bg_enabled !== undefined) bannerConfig.value.bg_enabled = config.bg_enabled
    if (config.bg_color) bannerConfig.value.bg_color = config.bg_color
    if (config.wide_bg !== undefined) bannerConfig.value.wide_bg = config.wide_bg
    // Mobile specific settings
    if (config.mobile_vertical_display !== undefined) bannerConfig.value.mobile_vertical_display = config.mobile_vertical_display
    if (config.mobile_height !== undefined) bannerConfig.value.mobile_height = config.mobile_height
    if (config.mobile_banner_width !== undefined) bannerConfig.value.mobile_banner_width = config.mobile_banner_width
    if (config.mobile_image_mode !== undefined) bannerConfig.value.mobile_image_mode = config.mobile_image_mode
    if (config.mobile_crop_width !== undefined) bannerConfig.value.mobile_crop_width = config.mobile_crop_width
    if (config.mobile_crop_height !== undefined) bannerConfig.value.mobile_crop_height = config.mobile_crop_height
    if (config.mobile_cropped_image !== undefined) bannerConfig.value.mobile_cropped_image = config.mobile_cropped_image
    if (config.mobile_image !== undefined) bannerConfig.value.mobile_image = config.mobile_image
    if (config.mobile_padding_top !== undefined) bannerConfig.value.mobile_padding_top = config.mobile_padding_top
    if (config.mobile_padding_bottom !== undefined) bannerConfig.value.mobile_padding_bottom = config.mobile_padding_bottom
    // Banner ratios
    if (config.banner1_ratio) bannerConfig.value.banner1_ratio = config.banner1_ratio
    if (config.banner2_ratio) bannerConfig.value.banner2_ratio = config.banner2_ratio
    if (config.banner3_ratio) bannerConfig.value.banner3_ratio = config.banner3_ratio
    if (config.banner4_ratio) bannerConfig.value.banner4_ratio = config.banner4_ratio
    // Padding and margins
    if (config.padding_top !== undefined) bannerConfig.value.padding_top = config.padding_top
    if (config.padding_bottom !== undefined) bannerConfig.value.padding_bottom = config.padding_bottom
    if (config.margin_left !== undefined) bannerConfig.value.margin_left = config.margin_left
    if (config.margin_right !== undefined) bannerConfig.value.margin_right = config.margin_right
    // Show title and description
    if (config.show_title !== undefined) bannerConfig.value.show_title = config.show_title
    if (config.show_description !== undefined) bannerConfig.value.show_description = config.show_description
  }
})
</script>