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

    <!-- تنظیمات بنر -->
    <DeviceTabs
      ref="deviceTabsRef"
      v-model:active-tab="activeDeviceTab"
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
            <h3 class="text-lg font-bold text-gray-700">تنظیمات بنر</h3>
          </div>
          
                     <div class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 gap-6">

            <!-- پس‌زمینه فعال -->
            <div>
              <label class="block mb-2 text-sm font-medium text-gray-700">پس‌زمینه فعال</label>
              <select
                v-model="bannerConfig.bg_enabled"
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
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
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
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
                class="w-full h-10 border border-gray-300 rounded-md"
              />
            </div>

            <!-- ارتفاع بنر -->
            <div>
              <label class="block mb-2 text-sm font-medium text-gray-700">ارتفاع بنر (پیکسل)</label>
              <select
                v-model="bannerConfig.height"
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
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
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
              >
                <option :value="800">تمام عرض</option>
                <option :value="600">وسط</option>
              </select>
            </div>

            <!-- عرض دقیق وسط -->
            <div v-if="bannerConfig.banner_width === 600">
              <label class="block mb-2 text-sm font-medium text-gray-700">عرض دقیق وسط (px)</label>
              <input
                type="number"
                :value="bannerConfig.center_width !== undefined ? bannerConfig.center_width : ''"
                min="200"
                max="2000"
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
                placeholder="1000"
                @input="(e: any) => bannerConfig.center_width = e.target.value === '' ? undefined : Number(e.target.value)"
              />
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
                @input="(e: any) => bannerConfig.padding_top = e.target.value === '' ? undefined : Number(e.target.value)"
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
                @input="(e: any) => bannerConfig.padding_bottom = e.target.value === '' ? undefined : Number(e.target.value)"
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
                @input="(e: any) => bannerConfig.margin_right = e.target.value === '' ? undefined : Number(e.target.value)"
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
                @input="(e: any) => bannerConfig.margin_left = e.target.value === '' ? undefined : Number(e.target.value)"
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
            <button
              class="px-4 py-2 rounded-md font-bold transition-colors"
              :class="bannerConfig.banners.length >= 1 
                ? 'bg-gray-400 text-gray-600 cursor-not-allowed' 
                : 'bg-cyan-500 text-white hover:bg-cyan-600'"
              :disabled="bannerConfig.banners.length >= 1"
              @click="openAddBannerModal"
            >
              {{ bannerConfig.banners.length >= 1 ? 'حداکثر یک بنر مجاز است' : 'افزودن بنر جدید' }}
            </button>
          </div>
          
          <div class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 gap-6">
            <!-- ارتفاع موبایل -->
            <div>
              <label class="block mb-2 text-sm font-medium text-gray-700">ارتفاع موبایل (px)</label>
              <select
                v-model="bannerConfig.mobile_height"
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
              >
                <option :value="100">100 پیکسل</option>
                <option :value="150" selected>150 پیکسل</option>
                <option :value="200">200 پیکسل</option>
                <option :value="250">250 پیکسل</option>
                <option :value="300">300 پیکسل</option>
                <option :value="400">400 پیکسل</option>
                <option :value="500">500 پیکسل</option>
              </select>
            </div>

            <!-- عرض موبایل -->
            <div>
              <label class="block mb-2 text-sm font-medium text-gray-700">عرض موبایل (px)</label>
              <select
                v-model.number="bannerConfig.mobile_banner_width"
                class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-400"
              >
                <option :value="400">400 پیکسل</option>
                <option :value="500">500 پیکسل</option>
                <option :value="600">600 پیکسل</option>
                <option :value="700">700 پیکسل</option>
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
                        min="100"
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
                @input="(e: any) => bannerConfig.mobile_padding_top = e.target.value === '' ? undefined : Number(e.target.value)"
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
                @input="(e: any) => bannerConfig.mobile_padding_bottom = e.target.value === '' ? undefined : Number(e.target.value)"
              />
            </div>
          </div>

          <!-- پیش نمایش موبایل -->
          <div class="mt-8">
            <h4 class="text-lg font-bold text-gray-700 mb-4">پیش نمایش موبایل</h4>
            <div class="bg-gray-100 p-6 rounded-lg">
              <div class="max-w-sm mx-auto bg-white rounded-lg shadow-lg overflow-hidden">
                <div 
                  class="relative overflow-hidden rounded-lg"
                  :style="{
                    height: `${bannerConfig.mobile_height}px`,
                    backgroundColor: bannerConfig.bg_enabled ? bannerConfig.bg_color : 'transparent',
                    paddingTop: bannerConfig.mobile_padding_top ? `${bannerConfig.mobile_padding_top}px` : '0',
                    paddingBottom: bannerConfig.mobile_padding_bottom ? `${bannerConfig.mobile_padding_bottom}px` : '0'
                  }"
                >
                  <!-- نمایش بنرها -->
                  <div v-if="bannerConfig.banners.length > 0" class="relative h-full">
                    <div 
                      v-for="(banner, index) in bannerConfig.banners" 
                      :key="index"
                      class="absolute inset-0 transition-opacity duration-500"
                      :class="{ 'opacity-100': index === currentPreviewBanner, 'opacity-0': index !== currentPreviewBanner }"
                    >
                      <!-- نمایش عکس موبایل -->
                      <img 
                        :src="banner.image"
                        :alt="banner.title"
                        class="w-full h-full"
                        :class="getMobileImageClass()"
                      />
                    </div>
                  </div>
                  <div v-else class="flex items-center justify-center h-full text-gray-400 text-sm">
                    تصویر بنر انتخاب نشده است
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- مدیریت بنرها در موبایل -->
          <div class="mt-8">
            <h4 class="text-lg font-bold text-gray-700 mb-4">مدیریت بنرها</h4>
            <div v-if="bannerConfig.banners.length === 0" class="text-gray-400 text-center py-8">
              چیزی برای نمایش وجود ندارد!
            </div>
            <div v-else class="grid grid-cols-1 gap-6">
              <div
                v-for="(banner, idx) in bannerConfig.banners"
                :key="idx"
                class="flex items-center gap-6 p-3 border rounded-lg bg-gray-50"
              >
                <img
                  :src="banner.image"
                  alt="بنر"
                  class="w-24 h-16 object-cover rounded border"
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
            height: `${bannerConfig.height}px`,
            backgroundColor: bannerConfig.bg_enabled ? bannerConfig.bg_color : 'transparent',
            maxWidth: bannerConfig.banner_width === 800 ? '100%' : '600px',
            margin: '0 auto'
          }"
        >
          <!-- نمایش بنرها -->
          <div v-if="bannerConfig.banners.length > 0" class="relative h-full">
            <div 
              v-for="(banner, index) in bannerConfig.banners" 
              :key="index"
              class="absolute inset-0 transition-opacity duration-500"
              :class="{ 'opacity-100': index === currentPreviewBanner, 'opacity-0': index !== currentPreviewBanner }"
            >
              <img 
                :src="banner.image" 
                :alt="banner.title"
                class="w-full h-full object-cover"
              />
              <div
                v-if="bannerConfig.show_title || bannerConfig.show_description"
                class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-6"
              >
                <h4
                  v-if="bannerConfig.show_title && banner.title && (banner.showTitle !== false)"
                  class="text-white text-lg font-bold mb-2"
                >
                  {{ banner.title }}
                </h4>
                <p
                  v-if="bannerConfig.show_description && banner.description"
                  class="text-white/90 text-sm"
                >
                  {{ banner.description }}
                </p>
              </div>
            </div>
          </div>
          
          <!-- پیام عدم وجود بنر -->
          <div v-else class="flex items-center justify-center h-full text-gray-400">
            <div class="text-center">
              <svg class="w-16 h-16 mx-auto mb-4 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
              </svg>
              <p class="text-lg">هیچ بنری برای نمایش وجود ندارد</p>
              <p class="text-sm text-gray-500">برای شروع، یک بنر جدید اضافه کنید</p>
            </div>
          </div>
        </div>
        
        <!-- اطلاعات تنظیمات -->
        <div class="mt-4 text-sm text-gray-600 text-center">
          <p>ارتفاع: {{ bannerConfig.height }}px | 
              عرض: {{ bannerConfig.banner_width === 800 ? 'تمام عرض' : 'وسط صفحه' }} | 
              پس‌زمینه: {{ bannerConfig.bg_enabled ? 'فعال' : 'غیرفعال' }}</p>
        </div>
      </div>
    </div>

            <!-- نمایش و مدیریت اسلایدها -->
        <div v-if="deviceTabsRef?.activeTab === 'desktop'" class="bg-white rounded-xl shadow p-8 mt-6 mx-4">
      <div class="flex justify-between items-center mb-6">
        <h3 class="text-lg font-bold text-gray-700">بنر اصلی</h3>
        <button
          class="px-4 py-2 rounded-md font-bold transition-colors"
          :class="bannerConfig.banners.length >= 1 
            ? 'bg-gray-400 text-gray-600 cursor-not-allowed' 
            : 'bg-cyan-500 text-white hover:bg-cyan-600'"
          :disabled="bannerConfig.banners.length >= 1"
          @click="openAddBannerModal"
        >
          {{ bannerConfig.banners.length >= 1 ? 'حداکثر یک بنر مجاز است' : 'افزودن بنر جدید' }}
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
      :device-type="activeDeviceTab"
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

    <!-- Page content will be added here -->

    <!-- فاصله انتهای صفحه -->
    <div class="pb-16"></div>
  </div>
</template>

<script setup lang="ts">
import { WIDGET_TYPE_LABELS } from '~/types/widget'
import type { Widget, BannerConfig, BannerItem } from '~/types/widget'
import TemplateButton from '~/components/common/TemplateButton.vue'
import MediaLibraryModal from '~/components/media/MediaLibraryModal.vue'
import BannerModal from '@/components/common/BannerModal.vue'
import DeviceTabs from './components/DeviceTabs.vue'
import { useToast } from '~/composables/useToast'
import { useWidget } from '~/composables/useWidget'
import { ref, computed, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'

// تعریف definePageMeta و useHead برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
declare const useHead: (head: { title?: string }) => void

definePageMeta({ layout: 'admin-main', middleware: 'admin' })
useHead({ title: 'ویرایش بنر کامل - پنل ادمین' })

// Route params
const route = useRoute()
const widgetId = parseInt(route.params.id as string)

// Composables
const { fetchWidget, updateWidget, loading, error, clearError, widget } = useWidget()
const { showSuccess, showError } = useToast()

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
const currentPreviewBanner = ref(0)
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

// Editing banner data
const editingBanner = ref<BannerItem>({
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
const showTitleInBanner = ref(false)

// Form data
const formData = ref({
  title: '',
  description: '',
  type: 'full-banner',
  status: 'active',
  page: 'home',
  show_on_mobile: true
})

// Initialize form data when widget is available
const initializeFormData = () => {
  if (widget.value) {
    console.log('Widget data:', widget.value) // Debug log
    formData.value = {
      title: widget.value.title || '',
      description: widget.value.description || '',
      type: widget.value.type || 'full-banner',
      status: widget.value.status || 'active',
      page: widget.value.page || 'home',
      show_on_mobile: widget.value.show_on_mobile !== undefined ? widget.value.show_on_mobile : true
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

// Computed properties for reactive form data
const widgetTitle = computed(() => widget.value?.title || '')
const widgetType = computed(() => widget.value?.type || 'full-banner')
const widgetStatus = computed(() => widget.value?.status || 'active')
const widgetPage = computed(() => widget.value?.page || 'home')

// Active device tab
const activeDeviceTab = ref<'desktop' | 'mobile'>('desktop')

// Banner config
const bannerConfig = ref<BannerConfig>({
  wide_bg: false,
  bg_color: '#ffffff',
  height: 400,
  banners: [],
  mobile_banners: [],
  bg_enabled: false,
  banner_width: 800,
  show_title: true,
  show_description: true,
  padding_top: 0,
  padding_bottom: 0,
  margin_left: 0,
  margin_right: 0,
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
    showError('ابتدا یک بنر اضافه کنید')
    return
  }
  
  const banner = bannerConfig.value.banners[0]
  if (!banner.image) {
    showError('عکس بنر انتخاب نشده است')
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
      } else {
        throw new Error('API response failed')
      }
    } catch (error) {
      console.error('Backend crop API error:', error)
      // اگر API کراپ کار نکرد، هیچ fallback استفاده نکن
      return
    }
    
    showSuccess('برش موبایل با موفقیت اعمال شد')
  } catch (error) {
    showError('خطا در اعمال برش موبایل')
    console.error('Mobile crop error:', error)
  }
}

// Helper functions for mobile image handling
const getMobileImageUrl = (banner) => {
  // اگر حالت عکس جداگانه انتخاب شده و عکس موبایل وجود دارد
  if (bannerConfig.value.mobile_image_mode === 'separate' && banner.mobile_image) {
    return banner.mobile_image
  }
  
  // اگر حالت برش خودکار انتخاب شده و عکس برش شده وجود دارد
  if (bannerConfig.value.mobile_image_mode === 'auto' && bannerConfig.value.mobile_cropped_image) {
    return bannerConfig.value.mobile_cropped_image
  }
  
  // در غیر این صورت از عکس اصلی استفاده کن
  return banner.image
}

const getMobileImageClass = () => {
  // برای عکس‌های برش داده شده یا جداگانه، همیشه object-cover استفاده می‌کنیم
  return 'object-cover'
}

// Mobile image selector - removed as it now uses openAddBannerModal

// Methods for managing banners
const openAddBannerModal = () => {
  // بررسی محدودیت یک بنر برای بنر تکی برای هر دستگاه
  const currentBanners = activeDeviceTab.value === 'desktop' 
    ? bannerConfig.value.banners 
    : bannerConfig.value.mobile_banners || []
    
  if (currentBanners.length >= 1) {
    showError('در بنر تکی فقط یک بنر مجاز است')
    return
  }
  
  editingBannerIndex.value = null
  editingBanner.value = {
    title: '',
    description: '',
    image: '',
    mobile_image: '',
    link: '',
    openInNewTab: false,
    order: currentBanners.length + 1,
    status: 'active'
  }
  showTitleInBanner.value = false // Reset to default
  showBannerModal.value = true
}

const editBanner = (index: number) => {
  editingBannerIndex.value = index
  const currentBanners = activeDeviceTab.value === 'desktop' 
    ? bannerConfig.value.banners 
    : bannerConfig.value.mobile_banners || []
  const banner = currentBanners[index]
  editingBanner.value = {
    title: banner.title || '',
    description: banner.description || '',
    image: banner.image || '',
    mobile_image: banner.mobile_image || '',
    link: banner.link || '',
    openInNewTab: banner.openInNewTab || false,
    order: banner.order || index + 1,
    status: banner.status || 'active'
  }
  // Set showTitleInBanner based on existing banner data or default to false
  showTitleInBanner.value = banner.showTitle !== undefined ? banner.showTitle : false
  showBannerModal.value = true
}

const removeBanner = (index: number) => {
  if (activeDeviceTab.value === 'desktop') {
    bannerConfig.value.banners.splice(index, 1)
    // Re-order remaining banners
    bannerConfig.value.banners.forEach((banner, idx) => {
      banner.order = idx + 1
    })
  } else {
    if (!bannerConfig.value.mobile_banners) {
      bannerConfig.value.mobile_banners = []
    }
    bannerConfig.value.mobile_banners.splice(index, 1)
    // Re-order remaining banners
    bannerConfig.value.mobile_banners.forEach((banner, idx) => {
      banner.order = idx + 1
    })
  }
}

// Handle banner save from modal component
const handleBannerSave = (bannerData: BannerItem, showTitle: boolean) => {
  const currentBanners = activeDeviceTab.value === 'desktop' 
    ? bannerConfig.value.banners 
    : bannerConfig.value.mobile_banners || []
  
  // برای موبایل، اگر mobile_image وجود دارد، آن را به image کپی کن
  const imageUrl = activeDeviceTab.value === 'mobile' && bannerData.mobile_image 
    ? bannerData.mobile_image 
    : bannerData.image
    
  const banner: BannerItem = {
    ...bannerData,
    image: imageUrl, // استفاده از imageUrl به جای bannerData.image
    title: bannerData.title?.trim() || '',
    description: bannerData.description?.trim() || '',
    link: bannerData.link?.trim() || '',
    mobile_image: bannerData.mobile_image || '',
    order: editingBannerIndex.value !== null ? currentBanners[editingBannerIndex.value].order : currentBanners.length + 1,
    status: 'active',
    showTitle: showTitle
  }

  if (editingBannerIndex.value !== null) {
    // ویرایش بنر موجود
    if (activeDeviceTab.value === 'desktop') {
      bannerConfig.value.banners[editingBannerIndex.value] = banner
    } else {
      if (!bannerConfig.value.mobile_banners) {
        bannerConfig.value.mobile_banners = []
      }
      bannerConfig.value.mobile_banners[editingBannerIndex.value] = banner
    }
  } else {
    // افزودن بنر جدید
    if (activeDeviceTab.value === 'desktop') {
      bannerConfig.value.banners.push(banner)
    } else {
      if (!bannerConfig.value.mobile_banners) {
        bannerConfig.value.mobile_banners = []
      }
      bannerConfig.value.mobile_banners.push(banner)
    }
  }

  showBannerModal.value = false
  editingBannerIndex.value = null
}

// Media library functions
const openMediaLibrary = () => {
  showMediaLibrary.value = true
}

const onSelectFromLibrary = (files: any[]) => {
  if (files && files.length > 0) {
    const file = files[0]
    if (activeDeviceTab.value === 'mobile') {
      editingBanner.value.mobile_image = file.url
    } else {
      editingBanner.value.image = file.url
    }
  }
  showMediaLibrary.value = false
}

// Image functions
const removeImage = () => {
  if (activeDeviceTab.value === 'mobile') {
    editingBanner.value.mobile_image = ''
  } else {
    editingBanner.value.image = ''
  }
}


// Save widget function
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

    // Prepare widget data for update
    const configToSave = { ...bannerConfig.value }
    // Explicitly save both arrays
    configToSave.banners = bannerConfig.value.banners
    configToSave.mobile_banners = bannerConfig.value.mobile_banners || []
    
    const widgetData = {
      title: formData.value.title,
      description: formData.value.description,
      type: formData.value.type as any, // Type casting for compatibility
      status: formData.value.status as any, // Type casting for compatibility
      page: formData.value.page as any, // Type casting for compatibility
      show_on_mobile: formData.value.show_on_mobile,
      config: configToSave
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

// Initialize config from widget
onMounted(async () => {
  clearError()
  if (widgetId) {
    await fetchWidget(widgetId)
  }
  initializeFormData()
  
  // Only copy specific config fields, don't overwrite defaults
  if (widget.value?.config) {
    const config = widget.value.config as BannerConfig
    if (config.banners) bannerConfig.value.banners = config.banners
    if (config.mobile_banners) bannerConfig.value.mobile_banners = config.mobile_banners
    if (config.height) bannerConfig.value.height = config.height
    if (config.banner_width) bannerConfig.value.banner_width = config.banner_width
    if (config.center_width) bannerConfig.value.center_width = config.center_width
    if (config.bg_enabled !== undefined) bannerConfig.value.bg_enabled = config.bg_enabled
    if (config.bg_color) bannerConfig.value.bg_color = config.bg_color
    if (config.wide_bg !== undefined) bannerConfig.value.wide_bg = config.wide_bg
    if (config.padding_top !== undefined) bannerConfig.value.padding_top = config.padding_top
    if (config.padding_bottom !== undefined) bannerConfig.value.padding_bottom = config.padding_bottom
    if (config.margin_left !== undefined) bannerConfig.value.margin_left = config.margin_left
    if (config.margin_right !== undefined) bannerConfig.value.margin_right = config.margin_right
    // Mobile specific settings
    if (config.mobile_height !== undefined) bannerConfig.value.mobile_height = config.mobile_height
    if (config.mobile_banner_width !== undefined) bannerConfig.value.mobile_banner_width = config.mobile_banner_width
    if (config.mobile_image_mode !== undefined) bannerConfig.value.mobile_image_mode = config.mobile_image_mode
    if (config.mobile_crop_width !== undefined) bannerConfig.value.mobile_crop_width = config.mobile_crop_width
    if (config.mobile_crop_height !== undefined) bannerConfig.value.mobile_crop_height = config.mobile_crop_height
    if (config.mobile_image !== undefined) bannerConfig.value.mobile_image = config.mobile_image
    if (config.mobile_vertical_display !== undefined) bannerConfig.value.mobile_vertical_display = config.mobile_vertical_display
    if (config.mobile_padding_top !== undefined) bannerConfig.value.mobile_padding_top = config.mobile_padding_top
    if (config.mobile_padding_bottom !== undefined) bannerConfig.value.mobile_padding_bottom = config.mobile_padding_bottom
  }
})
</script>