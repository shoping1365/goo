<template>
  <div class="w-full">
    <!-- Title Section -->
    <div class="bg-white w-full shadow-sm border-2 border-blue-200 rounded-lg mt-4">
      <div class="flex justify-between items-center px-8 py-6">
        <h1 class="text-2xl font-bold text-gray-800">ایجاد بنر چهارتایی</h1>
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

    <!-- تنظیمات بنر -->
    <DeviceTabs
      ref="deviceTabsRef"
      v-model:banner-config="bannerConfig"
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
            <!-- نمایش عمودی در موبایل -->
            <div>
              <label class="block mb-2 text-sm font-medium text-gray-700">نمایش عمودی</label>
              <select
                v-model="bannerConfig.mobile_vertical_display"
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
              >
                <option :value="true">فعال</option>
                <option :value="false">غیرفعال</option>
              </select>
            </div>

            <!-- پدینگ بالا موبایل -->
            <div>
              <label class="block mb-2 text-sm font-medium text-gray-700">پدینگ بالا موبایل (px)</label>
              <input
                type="number"
                :value="bannerConfig.mobile_padding_top !== undefined ? bannerConfig.mobile_padding_top : ''"
                min="0"
                max="100"
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
                placeholder="0"
                @input="(e: Event) => bannerConfig.mobile_padding_top = (e.target as HTMLInputElement).value === '' ? undefined : Number((e.target as HTMLInputElement).value)"
              />
            </div>

            <!-- پدینگ پایین موبایل -->
            <div>
              <label class="block mb-2 text-sm font-medium text-gray-700">پدینگ پایین موبایل (px)</label>
              <input
                type="number"
                :value="bannerConfig.mobile_padding_bottom !== undefined ? bannerConfig.mobile_padding_bottom : ''"
                min="0"
                max="100"
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
                placeholder="0"
                @input="(e: Event) => bannerConfig.mobile_padding_bottom = (e.target as HTMLInputElement).value === '' ? undefined : Number((e.target as HTMLInputElement).value)"
              />
            </div>
          </div>
        </div>
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
    <BannerModal
      :is-visible="showBannerModal"
      :is-editing="editingBannerIndex !== null"
      :banner-data="editingBanner"
      :show-title="showTitleInBanner"
      @update:is-visible="showBannerModal = $event"
      @update:show-title="showTitleInBanner = $event"
      @update:banner-data="editingBanner = $event"
      @save="handleBannerSave"
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

    <!-- فاصله انتهای صفحه -->
    <div class="pb-16"></div>
  </div>
</template>

<script setup lang="ts">
import BannerModal from '@/components/common/BannerModal.vue'
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import TemplateButton from '~/components/common/TemplateButton.vue'
import MediaLibraryModal from '~/components/media/MediaLibraryModal.vue'
import { useToast } from '~/composables/useToast'
import { useWidget } from '~/composables/useWidget'
import type { BannerConfig, BannerItem, Widget, WidgetPage, WidgetStatus, WidgetType } from '~/types/widget'
import { WIDGET_TYPE_LABELS } from '~/types/widget'
import DeviceTabs from './components/DeviceTabs.vue'

interface MediaFile {
  id?: number | string
  url: string
  name?: string
  [key: string]: unknown
}

// تعریف definePageMeta و useHead برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
declare const useHead: (head: { title?: string }) => void

definePageMeta({ layout: 'admin-main', middleware: 'admin' })
useHead({ title: 'ایجاد بنر چهارتایی - پنل ادمین' })

// Router
const router = useRouter()

// Composables
const { createWidget, clearError } = useWidget()
const { showSuccess, showError } = useToast()

// Props
interface Props {
  widget?: Widget
}

const _props = defineProps<Props>()

// Emits
const _emit = defineEmits<{
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

// Editing banner data
const editingBanner = ref<BannerItem>({
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

// Banner config
const bannerConfig = ref<BannerConfig>({
  wide_bg: false,
  bg_color: '#ffffff',
  height: 400,
  banners: [],
  bg_enabled: false,
  banner_width: 800,
  center_width: 1000,
  banner1_ratio: 25,
  banner2_ratio: 25,
  banner3_ratio: 25,
  banner4_ratio: 25,
  show_title: true,
  show_description: true,
  padding_top: 0,
  padding_bottom: 0,
  margin_left: 0,
  margin_right: 0,
  // Mobile specific settings
  mobile_vertical_display: false,
  mobile_height: 150,
  // Mobile padding and margin settings
  mobile_padding_top: undefined,
  mobile_padding_bottom: undefined
})

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
  showTitleInBanner.value = true
  showBannerModal.value = true
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
  showTitleInBanner.value = banner.showTitle !== undefined ? banner.showTitle : true
  showBannerModal.value = true
}

const removeBanner = (index: number) => {
  bannerConfig.value.banners.splice(index, 1)
  bannerConfig.value.banners.forEach((banner, idx) => {
    banner.order = idx + 1
  })
}

// Handle banner save from modal component
const handleBannerSave = (bannerData: BannerItem, showTitle: boolean) => {
  const banner: BannerItem = {
    ...bannerData,
    title: bannerData.title?.trim() || '',
    description: bannerData.description?.trim() || '',
    link: bannerData.link?.trim() || '',
    order: editingBannerIndex.value !== null ? bannerConfig.value.banners[editingBannerIndex.value].order : bannerConfig.value.banners.length + 1,
    status: 'active',
    showTitle: showTitle
  }

  if (editingBannerIndex.value !== null) {
    bannerConfig.value.banners[editingBannerIndex.value] = banner
  } else {
    bannerConfig.value.banners.push(banner)
  }

  showBannerModal.value = false
  editingBannerIndex.value = null
}

// Media library functions
const openMediaLibrary = () => {
  showMediaLibrary.value = true
}

const onSelectFromLibrary = (files: MediaFile[]) => {
  if (files && files.length > 0) {
    const file = files[0]
    editingBanner.value.image = file.url
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
    bannerConfig.value[`${changedBanner}_ratio`] = 100
    bannerConfig.value.banner1_ratio = 0
    bannerConfig.value.banner2_ratio = 0
    bannerConfig.value.banner3_ratio = 0
    bannerConfig.value.banner4_ratio = 0
    return
  }
  
  const otherBanners = ['banner1', 'banner2', 'banner3', 'banner4'].filter(b => b !== changedBanner)
  
  if (otherBanners.length === 3) {
    const [bannerA, bannerB, bannerC] = otherBanners
    const currentA = bannerConfig.value[`${bannerA}_ratio`]
    const currentB = bannerConfig.value[`${bannerB}_ratio`]
    const currentC = bannerConfig.value[`${bannerC}_ratio`]
    const totalCurrent = currentA + currentB + currentC
    
    if (totalCurrent > 0) {
      const ratioA = currentA / totalCurrent
      const ratioB = currentB / totalCurrent
      const _ratioC = currentC / totalCurrent
      
      bannerConfig.value[`${bannerA}_ratio`] = Math.round(remaining * ratioA)
      bannerConfig.value[`${bannerB}_ratio`] = Math.round(remaining * ratioB)
      bannerConfig.value[`${bannerC}_ratio`] = remaining - bannerConfig.value[`${bannerA}_ratio`] - bannerConfig.value[`${bannerB}_ratio`]
    } else {
      const equalShare = Math.floor(remaining / 3)
      bannerConfig.value[`${bannerA}_ratio`] = equalShare
      bannerConfig.value[`${bannerB}_ratio`] = equalShare
      bannerConfig.value[`${bannerC}_ratio`] = remaining - equalShare - equalShare
    }
  }
}

// Save widget function
const saveWidget = async () => {
  validateTitle()

  if (titleError.value) {
    const titleInput = document.querySelector('input[placeholder="عنوان ابزارک را وارد کنید"]') as HTMLInputElement
    if (titleInput) {
      titleInput.focus()
      titleInput.scrollIntoView({ behavior: 'smooth', block: 'center' })
    }
    return
  }

  try {
    isSaving.value = true

    const configToSave = { ...bannerConfig.value }
    configToSave.padding_top = bannerConfig.value.padding_top
    configToSave.padding_bottom = bannerConfig.value.padding_bottom
    configToSave.margin_left = bannerConfig.value.margin_left
    configToSave.margin_right = bannerConfig.value.margin_right
    configToSave.mobile_padding_top = bannerConfig.value.mobile_padding_top
    configToSave.mobile_padding_bottom = bannerConfig.value.mobile_padding_bottom

    const widgetData = {
      title: formData.value.title,
      description: formData.value.description,
      type: formData.value.type as WidgetType,
      status: formData.value.status as WidgetStatus,
      page: formData.value.page as WidgetPage,
      config: configToSave
    }

    const createdWidget = await createWidget(widgetData)

    if (createdWidget) {
      showSuccess('ابزارک با موفقیت ایجاد شد!')
      await router.push(`/admin/content/banners/quadbanner/edit/${createdWidget.id}`)
    }
  } catch (_error) {
    showError('خطا در ذخیره ابزارک')
  } finally {
    isSaving.value = false
  }
}

// Initialize config
onMounted(async () => {
  clearError()
})
</script>
