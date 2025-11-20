<template>
  <!-- اطلاعات اصلی محصول -->
  <div v-if="sectionSettings.mainInfo" class="bg-white border border-gray-200 rounded-xl shadow-lg overflow-hidden text-right mb-8">
    <!-- Header -->
    <div class="bg-gradient-to-r from-blue-600 to-cyan-600 p-6">
      <div class="flex items-center justify-between cursor-pointer" @click="toggleSection('mainInfo')">
        <div class="flex items-center gap-3">
          <div class="bg-white/20 p-2 rounded-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <div>
            <h3 class="text-xl font-bold text-white">اطلاعات اصلی محصول</h3>
            <p class="text-blue-100 text-sm mt-1">نام، توضیحات و معرفی محصول</p>
          </div>
        </div>
        <div class="bg-white/20 p-2 rounded-lg">
          <svg class="w-5 h-5 text-white transition-transform duration-200" :class="{ 'rotate-180': sections.mainInfo }" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
          </svg>
        </div>
      </div>
    </div>

    <div v-show="sections.mainInfo" class="p-6">
      <!-- نام محصول و نام انگلیسی -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-8">
        <div class="space-y-2">
          <label class="block text-sm font-semibold text-gray-700 flex items-center gap-2">
            <svg class="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z" />
            </svg>
            نام محصول
          </label>
          <input
            v-model="productForm.name"
            type="text"
            class="w-full px-4 py-3 text-right border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors"
            dir="rtl"
            placeholder="نام محصول را وارد کنید..."
            required
          />
        </div>
        <div class="space-y-2">
          <label class="block text-sm font-semibold text-gray-700 flex items-center gap-2">
            <svg class="w-4 h-4 text-cyan-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5h12M9 3v2m1.048 9.5A18.022 18.022 0 016.412 9m6.088 9h7M11 21l5-10 5 10M12.751 5C11.783 10.77 8.07 15.61 3 18.129" />
            </svg>
            Product English Name
          </label>
          <input
            v-model="productForm.englishName"
            type="text"
            class="w-full px-4 py-3 text-left border border-gray-300 rounded-lg focus:ring-2 focus:ring-cyan-500 focus:border-cyan-500 transition-colors"
            dir="ltr"
            placeholder="Enter product name in English..."
          />
        </div>
      </div>

      <!-- توضیحات کوتاه -->
      <div class="bg-gradient-to-br from-emerald-50 to-teal-50 border border-emerald-200 rounded-xl p-6 mb-8">
        <div class="flex items-center justify-between mb-4">
          <div class="flex items-center gap-3">
            <div class="bg-gradient-to-r from-emerald-500 to-teal-500 p-2 rounded-lg">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
              </svg>
            </div>
            <div>
              <label class="text-lg font-bold text-gray-800">توضیحات کوتاه</label>
            </div>
          </div>
          <button
              class="flex items-center gap-2 bg-gradient-to-r from-purple-500 to-indigo-500 text-white rounded-lg px-4 py-2 text-sm font-medium hover:from-purple-600 hover:to-indigo-600 transition-all duration-200 shadow-lg hover:shadow-xl"
              title="تولید محتوا با هوش مصنوعی"
              @click="generateAIContent('short')"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
            </svg>
            <span>تولید با AI</span>
          </button>
        </div>
        <div class="border border-emerald-300 rounded-lg overflow-hidden bg-white">
          <ClientOnly>
            <RichTextEditor
                v-model="productForm.description"
                :lang="'fa'"
                :direction="'rtl'"
                :height="300"
            />
          </ClientOnly>
        </div>
      </div>

      <!-- توضیحات کامل -->
      <div class="bg-gradient-to-br from-amber-50 to-orange-50 border border-amber-200 rounded-xl p-6">
        <div class="flex items-center justify-between mb-4">
          <div class="flex items-center gap-3">
            <div class="bg-gradient-to-r from-amber-500 to-orange-500 p-2 rounded-lg">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253" />
              </svg>
            </div>
            <div>
              <label class="text-lg font-bold text-gray-800">توضیحات کامل</label>
            </div>
          </div>
          <button
              class="flex items-center gap-2 bg-gradient-to-r from-purple-500 to-indigo-500 text-white rounded-lg px-4 py-2 text-sm font-medium hover:from-purple-600 hover:to-indigo-600 transition-all duration-200 shadow-lg hover:shadow-xl"
              title="تولید محتوا با هوش مصنوعی"
              @click="generateAIContent('full')"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
            </svg>
            <span>تولید با AI</span>
          </button>
        </div>
        <div class="border border-amber-300 rounded-lg overflow-hidden bg-white">
          <ClientOnly>
            <RichTextEditor
                v-model="productForm.fullDescription"
                :lang="'fa'"
                :direction="'rtl'"
                :height="500"
            />
          </ClientOnly>
        </div>
      </div>
    </div>
  </div>

  <!-- اطلاعات فنی و شناسایی -->
  <div v-if="sectionSettings.technicalInfo" class="bg-white border border-gray-200 rounded-xl shadow-lg overflow-hidden text-right mb-8">
    <!-- Header -->
    <div class="bg-gradient-to-r from-purple-600 to-pink-600 p-6">
      <div class="flex items-center justify-between cursor-pointer" @click="toggleSection('technicalInfo')">
        <div class="flex items-center gap-3">
          <div class="bg-white/20 p-2 rounded-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-3 7h3m-3 4h3m-6-4h.01M9 16h.01" />
            </svg>
          </div>
          <div>
            <h3 class="text-xl font-bold text-white">اطلاعات فنی و شناسایی</h3>
            <p class="text-purple-100 text-sm mt-1">کدها، دسته‌بندی و مشخصات فنی</p>
          </div>
        </div>
        <div class="bg-white/20 p-2 rounded-lg">
          <svg class="w-5 h-5 text-white transition-transform duration-200" :class="{ 'rotate-180': sections.technicalInfo }" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
          </svg>
        </div>
      </div>
    </div>

    <div v-show="sections.technicalInfo" class="p-6">
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- شناسه کالا (SKU) -->
        <div class="space-y-2">
          <label class="block text-sm font-semibold text-gray-700 flex items-center gap-2">
            <svg class="w-4 h-4 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 20l4-16m2 16l4-16M6 9h14M4 15h14" />
            </svg>
            شناسه کالا (SKU)
          </label>
          <input
            v-model="productForm.sku"
            type="text"
            class="w-full px-4 py-3 border border-gray-300 rounded-lg bg-gray-50 cursor-not-allowed text-gray-500"
            placeholder="پس از ذخیره ایجاد می‌شود"
            readonly
          />
          <div class="text-sm text-purple-600 flex items-center gap-2">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            پس از ذخیره محصول، این شناسه به‌صورت خودکار ایجاد می‌شود
          </div>
        </div>

        <!-- کد محصول در انبار -->
        <div class="space-y-2">
          <label class="block text-sm font-semibold text-gray-700 flex items-center gap-2">
            <svg class="w-4 h-4 text-indigo-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
            </svg>
            کد محصول در انبار
          </label>
          <input 
            type="text" 
            class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-colors" 
            placeholder="مثال: PROD-001" 
          />
          <div class="text-sm text-indigo-600 flex items-center gap-2">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            کد نمایشی محصول برای مشتریان
          </div>
        </div>

        <!-- دسته‌بندی اصلی -->
        <div class="space-y-2">
          <label class="block text-sm font-semibold text-gray-700 flex items-center gap-2">
            <svg class="w-4 h-4 text-emerald-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            </svg>
            دسته‌بندی اصلی
          </label>
          <select 
            :key="mainCategories.length"
            v-model="productForm.category_id"
            class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 transition-colors"
          >
            <option :value="''">انتخاب دسته‌بندی</option>
            <option v-for="cat in mainCategories" :key="cat.id" :value="cat.id">
              {{ cat.name }}
            </option>
          </select>
        </div>

        <!-- دسته‌بندی فرعی -->
        <div class="space-y-2">
          <label class="block text-sm font-semibold text-gray-700 flex items-center gap-2">
            <svg class="w-4 h-4 text-teal-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z" />
            </svg>
            دسته‌بندی فرعی
          </label>
           <select 
            v-model="productForm.sub_category_id"
            :disabled="!productForm.category_id"
            class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-teal-500 focus:border-teal-500 transition-colors disabled:bg-gray-100 disabled:cursor-not-allowed"
          >
            <option :value="''">انتخاب دسته‌بندی فرعی</option>
            <option v-for="cat in subCategories" :key="cat.id" :value="cat.id">
              {{ cat.name }}
            </option>
          </select>
        </div>

        <!-- برند -->
        <div class="space-y-2">
          <label class="block text-sm font-semibold text-gray-700 flex items-center gap-2">
            <svg class="w-4 h-4 text-rose-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z" />
            </svg>
            برند
          </label>
          <select :key="brands.length" v-model="productForm.brand_id" class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-rose-500 focus:border-rose-500 transition-colors">
            <option :value="''">انتخاب برند</option>
            <option v-for="b in brands" :key="b.id" :value="b.id">{{ b.name }}</option>
          </select>
        </div>

        <!-- واحد اندازه‌گیری -->
        <div class="space-y-2">
          <label class="block text-sm font-semibold text-gray-700 flex items-center gap-2">
            <svg class="w-4 h-4 text-amber-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
            </svg>
            واحد اندازه‌گیری
          </label>
          <select class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-amber-500 focus:border-amber-500 transition-colors">
            <option value="piece">عدد</option>
            <option value="kg">کیلوگرم</option>
            <option value="meter">متر</option>
            <option value="liter">لیتر</option>
            <option value="pack">بسته</option>
          </select>
        </div>
      </div>
    </div>
  </div>

  <!-- تنظیمات نمایش و انتشار -->
  <div v-if="sectionSettings.displaySettings" class="bg-white border border-gray-200 rounded-xl shadow-lg overflow-hidden text-right mb-8">
    <!-- Header -->
    <div class="bg-gradient-to-r from-indigo-600 to-purple-600 p-6">
      <div class="flex items-center justify-between cursor-pointer" @click="toggleSection('displaySettings')">
        <div class="flex items-center gap-3">
          <div class="bg-white/20 p-2 rounded-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
            </svg>
          </div>
          <div>
            <h3 class="text-xl font-bold text-white">تنظیمات نمایش و انتشار</h3>
            <p class="text-indigo-100 text-sm mt-1">کنترل نمایش محصول در بخش‌های مختلف سایت</p>
          </div>
        </div>
        <div class="bg-white/20 p-2 rounded-lg">
          <svg class="w-5 h-5 text-white transition-transform duration-200" :class="{ 'rotate-180': sections.displaySettings }" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
          </svg>
        </div>
      </div>
    </div>

    <div v-show="sections.displaySettings" class="p-6">
      <div class="space-y-8">
        <!-- وضعیت انتشار -->
        <div class="bg-gradient-to-br from-green-50 to-emerald-50 border border-green-200 rounded-xl p-6">
          <h4 class="text-lg font-bold text-green-800 mb-4 flex items-center gap-2">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            وضعیت انتشار و نمایش
          </h4>
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            <label class="bg-white border border-green-300 rounded-lg p-3 cursor-pointer hover:shadow-md transition-all duration-200 group">
              <div class="flex items-center gap-3">
                <input type="checkbox" class="w-5 h-5 text-green-600 border-green-300 rounded focus:ring-green-500" checked />
                <div>
                  <div class="text-sm font-semibold text-gray-800">منتشر شده</div>
                  <div class="text-xs text-green-600">قابل مشاهده مشتریان</div>
                </div>
              </div>
            </label>
            
            <label class="bg-white border border-blue-300 rounded-lg p-3 cursor-pointer hover:shadow-md transition-all duration-200 group">
              <div class="flex items-center gap-3">
                <input type="checkbox" class="w-5 h-5 text-blue-600 border-blue-300 rounded focus:ring-blue-500" />
                <div>
                  <div class="text-sm font-semibold text-gray-800">صفحه اصلی</div>
                  <div class="text-xs text-blue-600">نمایش در هوم پیج</div>
                </div>
              </div>
            </label>

            <label class="bg-white border border-purple-300 rounded-lg p-3 cursor-pointer hover:shadow-md transition-all duration-200 group">
              <div class="flex items-center gap-3">
                <input type="checkbox" class="w-5 h-5 text-purple-600 border-purple-300 rounded focus:ring-purple-500" />
                <div>
                  <div class="text-sm font-semibold text-gray-800">محصول ویژه</div>
                  <div class="text-xs text-purple-600">برچسب ویژه</div>
                </div>
              </div>
            </label>

            <label class="bg-white border border-orange-300 rounded-lg p-3 cursor-pointer hover:shadow-md transition-all duration-200 group">
              <div class="flex items-center gap-3">
                <input type="checkbox" class="w-5 h-5 text-orange-600 border-orange-300 rounded focus:ring-orange-500" />
                <div>
                  <div class="text-sm font-semibold text-gray-800">پیشنهاد ویژه</div>
                  <div class="text-xs text-orange-600">تخفیف خاص</div>
                </div>
              </div>
            </label>

            <label class="bg-white border border-red-300 rounded-lg p-3 cursor-pointer hover:shadow-md transition-all duration-200 group">
              <div class="flex items-center gap-3">
                <input type="checkbox" class="w-5 h-5 text-red-600 border-red-300 rounded focus:ring-red-500" />
                <div>
                  <div class="text-sm font-semibold text-gray-800">محصول پرفروش</div>
                  <div class="text-xs text-red-600">برچسب پرفروش</div>
                </div>
              </div>
            </label>

            <label class="bg-white border border-cyan-300 rounded-lg p-3 cursor-pointer hover:shadow-md transition-all duration-200 group">
              <div class="flex items-center gap-3">
                <input type="checkbox" class="w-5 h-5 text-cyan-600 border-cyan-300 rounded focus:ring-cyan-500" />
                <div>
                  <div class="text-sm font-semibold text-gray-800">نمایش در اسلایدر</div>
                  <div class="text-xs text-cyan-600">اسلایدر اصلی</div>
                </div>
              </div>
            </label>
          </div>
        </div>

        <!-- اولویت نمایش -->
        <div class="bg-gradient-to-br from-blue-50 to-indigo-50 border border-blue-200 rounded-xl p-6">
          <h4 class="text-lg font-bold text-blue-800 mb-4 flex items-center gap-2">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 11.5V14m0-2.5v-6a1.5 1.5 0 113 0m-3 6a1.5 1.5 0 00-3 0v2a7.5 7.5 0 0015 0v-5a1.5 1.5 0 00-3 0m-6-3V11m0-5.5v-1a1.5 1.5 0 013 0v1m0 0V11m0-5.5a1.5 1.5 0 013 0v3M7 15h3m6-3h3" />
            </svg>
            اولویت و دسته‌بندی نمایش
          </h4>
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <div class="space-y-2">
              <label class="block text-sm font-semibold text-gray-700 flex items-center gap-2">
                <svg class="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
                </svg>
                اولویت نمایش
              </label>
              <select class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors">
                <option value="1">بسیار بالا (1)</option>
                <option value="2">بالا (2)</option>
                <option value="3" selected>متوسط (3)</option>
                <option value="4">کم (4)</option>
                <option value="5">بسیار کم (5)</option>
              </select>
            </div>

            <div class="space-y-2">
              <label class="block text-sm font-semibold text-gray-700 flex items-center gap-2">
                <svg class="w-4 h-4 text-indigo-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                </svg>
                محدوده نمایش
              </label>
              <select class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-colors">
                <option value="">همه صفحات</option>
                <option value="homepage">فقط صفحه اصلی</option>
                <option value="category">فقط صفحه دسته‌بندی</option>
                <option value="search">فقط نتایج جستجو</option>
              </select>
            </div>
          </div>
        </div>

        <!-- محدودیت‌های دسترسی -->
        <div class="bg-gradient-to-br from-amber-50 to-yellow-50 border border-amber-200 rounded-xl p-6">
          <h4 class="text-lg font-bold text-amber-800 mb-4 flex items-center gap-2">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
            </svg>
            محدودیت‌های دسترسی
          </h4>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <label class="bg-white border border-amber-300 rounded-lg p-3 cursor-pointer hover:shadow-md transition-all duration-200">
              <div class="flex items-center gap-3">
                <input type="checkbox" class="w-5 h-5 text-amber-600 border-amber-300 rounded focus:ring-amber-500" />
                <div>
                  <div class="text-sm font-semibold text-gray-800">اعضای ثبت‌نام شده</div>
                  <div class="text-xs text-amber-600">نیاز به حساب کاربری</div>
                </div>
              </div>
            </label>
            
            <label class="bg-white border border-yellow-300 rounded-lg p-3 cursor-pointer hover:shadow-md transition-all duration-200">
              <div class="flex items-center gap-3">
                <input type="checkbox" class="w-5 h-5 text-yellow-600 border-yellow-300 rounded focus:ring-yellow-500" />
                <div>
                  <div class="text-sm font-semibold text-gray-800">اعضای VIP</div>
                  <div class="text-xs text-yellow-600">فقط مشتریان ویژه</div>
                </div>
              </div>
            </label>

            <label class="bg-white border border-orange-300 rounded-lg p-3 cursor-pointer hover:shadow-md transition-all duration-200">
              <div class="flex items-center gap-3">
                <input type="checkbox" class="w-5 h-5 text-orange-600 border-orange-300 rounded focus:ring-orange-500" />
                <div>
                  <div class="text-sm font-semibold text-gray-800">تایید مدیر</div>
                  <div class="text-xs text-orange-600">نیاز به مجوز خرید</div>
                </div>
              </div>
            </label>

            <label class="bg-white border border-red-300 rounded-lg p-3 cursor-pointer hover:shadow-md transition-all duration-200">
              <div class="flex items-center gap-3">
                <input type="checkbox" class="w-5 h-5 text-red-600 border-red-300 rounded focus:ring-red-500" />
                <div>
                  <div class="text-sm font-semibold text-gray-800">محدودیت سنی</div>
                  <div class="text-xs text-red-600">بالای 18 سال</div>
                </div>
              </div>
            </label>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- زمان‌بندی و تاریخ‌ها -->
  <div v-if="sectionSettings.scheduling" class="bg-white border border-gray-200 rounded-xl shadow-lg overflow-hidden text-right mb-8">
    <!-- Header -->
    <div class="bg-gradient-to-r from-emerald-600 to-teal-600 p-6">
      <div class="flex items-center justify-between cursor-pointer" @click="toggleSection('scheduling')">
        <div class="flex items-center gap-3">
          <div class="bg-white/20 p-2 rounded-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
            </svg>
          </div>
          <div>
            <h3 class="text-xl font-bold text-white">زمان‌بندی و تاریخ‌ها</h3>
            <p class="text-emerald-100 text-sm mt-1">برنامه‌ریزی انتشار و موجودی محصول</p>
          </div>
        </div>
        <div class="bg-white/20 p-2 rounded-lg">
          <svg class="w-5 h-5 text-white transition-transform duration-200" :class="{ 'rotate-180': sections.scheduling }" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
          </svg>
        </div>
      </div>
    </div>

    <div v-show="sections.scheduling" class="p-6">
      <div class="space-y-8">
        <!-- تاریخ‌های موجودی -->
        <div class="bg-gradient-to-br from-blue-50 to-cyan-50 border border-blue-200 rounded-xl p-6">
          <h4 class="text-lg font-bold text-blue-800 mb-4 flex items-center gap-2">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            دوره فروش محصول
          </h4>
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <div class="space-y-2">
              <label class="block text-sm font-semibold text-gray-700 flex items-center gap-2">
                <svg class="w-4 h-4 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z" />
                </svg>
                تاریخ شروع فروش
              </label>
              <input 
                type="datetime-local" 
                class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500 transition-colors" 
              />
              <div class="text-sm text-green-600 flex items-center gap-2">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                محصول از این تاریخ قابل خرید می‌شود
              </div>
            </div>
            <div class="space-y-2">
              <label class="block text-sm font-semibold text-gray-700 flex items-center gap-2">
                <svg class="w-4 h-4 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z" />
                </svg>
                تاریخ پایان فروش
              </label>
              <input 
                type="datetime-local" 
                class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500 transition-colors" 
              />
              <div class="text-sm text-red-600 flex items-center gap-2">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                فروش محصول در این تاریخ متوقف می‌شود
              </div>
            </div>
          </div>
        </div>

        <!-- علامت‌گذاری کالای جدید -->
        <div class="bg-gradient-to-br from-purple-50 to-pink-50 border border-purple-200 rounded-xl p-6">
          <div class="flex items-center gap-3 mb-6">
            <input id="newProduct" type="checkbox" class="w-5 h-5 text-purple-600 border-purple-300 rounded focus:ring-purple-500" />
            <label for="newProduct" class="text-lg font-bold text-purple-800 flex items-center gap-2 cursor-pointer">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z" />
              </svg>
              علامت‌گذاری به عنوان کالای جدید
            </label>
          </div>

          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <div class="space-y-2">
              <label class="block text-sm font-semibold text-gray-700 flex items-center gap-2">
                <svg class="w-4 h-4 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                </svg>
                تاریخ شروع برچسب "جدید"
              </label>
              <input 
                type="date" 
                class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500 transition-colors" 
              />
            </div>
            <div class="space-y-2">
              <label class="block text-sm font-semibold text-gray-700 flex items-center gap-2">
                <svg class="w-4 h-4 text-pink-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                </svg>
                تاریخ پایان برچسب "جدید"
              </label>
              <input 
                type="date" 
                class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-pink-500 focus:border-pink-500 transition-colors" 
              />
            </div>
          </div>
        </div>

        <!-- فصلی بودن محصول -->
        <div class="bg-gradient-to-br from-orange-50 to-yellow-50 border border-orange-200 rounded-xl p-6">
          <h4 class="text-lg font-bold text-orange-800 mb-4 flex items-center gap-2">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z" />
            </svg>
            تنظیمات فصلی و نوع محصول
          </h4>
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <div class="space-y-2">
              <label class="block text-sm font-semibold text-gray-700 flex items-center gap-2">
                <svg class="w-4 h-4 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z" />
                </svg>
                نوع محصول
              </label>
              <select class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-orange-500 transition-colors">
                <option value="regular">عادی</option>
                <option value="seasonal">فصلی</option>
                <option value="limited">محدود</option>
                <option value="preorder">پیش‌خرید</option>
              </select>
            </div>
            <div class="space-y-2">
              <label class="block text-sm font-semibold text-gray-700 flex items-center gap-2">
                <svg class="w-4 h-4 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z" />
                </svg>
                فصل مرتبط
              </label>
              <select class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500 transition-colors">
                <option value="">همه فصل‌ها</option>
                <option value="spring">🌸 بهار</option>
                <option value="summer">☀️ تابستان</option>
                <option value="autumn">🍂 پاییز</option>
                <option value="winter">❄️ زمستان</option>
              </select>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- مدیریت و یادداشت‌ها -->
  <div v-if="sectionSettings.management" class="bg-white border border-gray-200 rounded-xl shadow-lg overflow-hidden text-right mb-8">
    <!-- Header -->
    <div class="bg-gradient-to-r from-slate-600 to-gray-700 p-6">
      <div class="flex items-center justify-between cursor-pointer" @click="toggleSection('management')">
        <div class="flex items-center gap-3">
          <div class="bg-white/20 p-2 rounded-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-3 7h3m-3 4h3m-6-4h.01M9 16h.01" />
            </svg>
          </div>
          <div>
            <h3 class="text-xl font-bold text-white">مدیریت و یادداشت‌ها</h3>
            <p class="text-slate-100 text-sm mt-1">اطلاعات اداری و یادداشت‌های داخلی</p>
          </div>
        </div>
        <div class="bg-white/20 p-2 rounded-lg">
          <svg class="w-5 h-5 text-white transition-transform duration-200" :class="{ 'rotate-180': sections.management }" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
          </svg>
        </div>
      </div>
    </div>

    <div v-show="sections.management" class="p-6">
      <div class="space-y-8">
        <!-- یادداشت‌های مدیریتی -->
        <div class="bg-gradient-to-br from-blue-50 to-indigo-50 border border-blue-200 rounded-xl p-6">
          <h4 class="text-lg font-bold text-blue-800 mb-4 flex items-center gap-2">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
            </svg>
            یادداشت‌های مدیریتی
          </h4>
          <div class="space-y-4">
            <textarea 
              class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors min-h-[120px] text-right resize-none" 
              dir="rtl" 
              placeholder="یادداشت‌ها و توضیحات مدیریتی که فقط برای مدیران قابل مشاهده است..."
            ></textarea>
            <div class="flex items-center gap-2 text-sm text-blue-600">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
              </svg>
              این یادداشت‌ها محرمانه بوده و فقط برای مدیران سایت قابل مشاهده است
            </div>
          </div>
        </div>

        <!-- تنظیمات پیشرفته -->
        <div class="bg-gradient-to-br from-purple-50 to-violet-50 border border-purple-200 rounded-xl p-6">
          <h4 class="text-lg font-bold text-purple-800 mb-4 flex items-center gap-2">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
            </svg>
            تنظیمات پیشرفته
          </h4>
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <div class="space-y-2">
              <label class="block text-sm font-semibold text-gray-700 flex items-center gap-2">
                <svg class="w-4 h-4 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                </svg>
                مدیر مسئول محصول
              </label>
              <select class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500 transition-colors">
                <option value="">انتخاب مدیر</option>
                <option value="admin1">👤 مدیر فروش</option>
                <option value="admin2">📦 مدیر انبار</option>
                <option value="admin3">🛍️ مدیر محصول</option>
              </select>
            </div>

            <div class="space-y-2">
              <label class="block text-sm font-semibold text-gray-700 flex items-center gap-2">
                <svg class="w-4 h-4 text-violet-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4M7.835 4.697a3.42 3.42 0 001.946-.806 3.42 3.42 0 014.438 0 3.42 3.42 0 001.946.806 3.42 3.42 0 013.138 3.138 3.42 3.42 0 00.806 1.946 3.42 3.42 0 010 4.438 3.42 3.42 0 00-.806 1.946 3.42 3.42 0 01-3.138 3.138 3.42 3.42 0 00-1.946.806 3.42 3.42 0 01-4.438 0 3.42 3.42 0 00-1.946-.806 3.42 3.42 0 01-3.138-3.138 3.42 3.42 0 00-.806-1.946 3.42 3.42 0 010-4.438 3.42 3.42 0 00.806-1.946 3.42 3.42 0 013.138-3.138z" />
                </svg>
                وضعیت تولید
              </label>
              <select class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-violet-500 focus:border-violet-500 transition-colors">
                <option value="active">✅ فعال</option>
                <option value="discontinued">⏹️ متوقف شده</option>
                <option value="development">🔧 در حال توسعه</option>
                <option value="testing">🧪 در حال تست</option>
              </select>
            </div>
          </div>
        </div>

        <!-- رتبه‌بندی و امتیاز -->
        <div class="bg-gradient-to-br from-amber-50 to-orange-50 border border-amber-200 rounded-xl p-6">
          <h4 class="text-lg font-bold text-amber-800 mb-4 flex items-center gap-2">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z" />
            </svg>
            رتبه‌بندی و کیفیت اولیه
          </h4>
          <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
            <div class="space-y-2">
              <label class="block text-sm font-semibold text-gray-700 flex items-center gap-2">
                <svg class="w-4 h-4 text-amber-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z" />
                </svg>
                امتیاز اولیه
              </label>
              <select class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-amber-500 focus:border-amber-500 transition-colors">
                <option value="5">⭐⭐⭐⭐⭐ 5 ستاره</option>
                <option value="4">⭐⭐⭐⭐ 4 ستاره</option>
                <option value="3">⭐⭐⭐ 3 ستاره</option>
                <option value="2">⭐⭐ 2 ستاره</option>
                <option value="1">⭐ 1 ستاره</option>
              </select>
            </div>

            <div class="space-y-2">
              <label class="block text-sm font-semibold text-gray-700 flex items-center gap-2">
                <svg class="w-4 h-4 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                درجه کیفیت
              </label>
              <select class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-orange-500 transition-colors">
                <option value="premium">💎 پریمیوم</option>
                <option value="high">🔷 بالا</option>
                <option value="medium">🔶 متوسط</option>
                <option value="basic">🔸 پایه</option>
              </select>
            </div>

            <div class="space-y-2">
              <label class="block text-sm font-semibold text-gray-700 flex items-center gap-2">
                <svg class="w-4 h-4 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
                </svg>
                گارانتی (ماه)
              </label>
              <input 
                type="number" 
                class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500 transition-colors" 
                min="0" 
                placeholder="12" 
              />
            </div>
          </div>
        </div>

        <!-- خلاصه وضعیت -->
        <div class="bg-gradient-to-br from-gray-50 to-slate-50 border border-gray-200 rounded-xl p-6">
          <h4 class="text-lg font-bold text-gray-800 mb-4 flex items-center gap-2">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
            </svg>
            خلاصه وضعیت محصول
          </h4>
          <div class="grid grid-cols-2 lg:grid-cols-4 gap-6">
            <div class="bg-white border border-gray-200 rounded-lg p-6 text-center">
              <div class="text-2xl font-bold text-green-600 mb-1">پیش‌نویس</div>
              <div class="text-sm text-gray-600">وضعیت</div>
            </div>
            <div class="bg-white border border-gray-200 rounded-lg p-6 text-center">
              <div class="text-2xl font-bold text-blue-600 mb-1">65%</div>
              <div class="text-sm text-gray-600">تکمیل شده</div>
            </div>
            <div class="bg-white border border-gray-200 rounded-lg p-6 text-center">
              <div class="text-2xl font-bold text-orange-600 mb-1">امروز</div>
              <div class="text-sm text-gray-600">آخرین ویرایش</div>
            </div>
            <div class="bg-white border border-gray-200 rounded-lg p-6 text-center">
              <div class="text-2xl font-bold text-purple-600 mb-1">مدیر سیستم</div>
              <div class="text-sm text-gray-600">ویرایشگر</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- نقاط قوت و ضعف محصول -->
  <div v-if="sectionSettings.strengthsWeaknesses" class="bg-white border border-gray-200 rounded-xl shadow-lg overflow-hidden text-right mb-8">
    <!-- Header -->
    <div class="bg-gradient-to-r from-teal-600 to-cyan-600 p-6">
      <div class="flex items-center justify-between cursor-pointer" @click="toggleSection('strengthsWeaknesses')">
        <div class="flex items-center gap-3">
          <div class="bg-white/20 p-2 rounded-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
            </svg>
          </div>
          <div>
            <h3 class="text-xl font-bold text-white">نقاط قوت و ضعف محصول</h3>
            <p class="text-teal-100 text-sm mt-1">تحلیل جامع مزایا و معایب محصول برای مشتریان</p>
          </div>
        </div>
        <div class="bg-white/20 p-2 rounded-lg">
          <svg class="w-5 h-5 text-white transition-transform duration-200" :class="{ 'rotate-180': sections.strengthsWeaknesses }" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
          </svg>
        </div>
      </div>
    </div>

    <div v-show="sections.strengthsWeaknesses" class="p-6">
      <div class="space-y-8">
        <div class="grid grid-cols-1 xl:grid-cols-2 gap-8">
          <!-- نقاط قوت -->
          <div class="bg-gradient-to-br from-green-50 to-emerald-50 border border-green-200 rounded-xl p-6">
            <div class="flex items-center justify-between mb-6">
              <h4 class="text-lg font-bold text-green-800 flex items-center gap-2">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                نقاط قوت محصول
              </h4>
              <span class="text-sm text-green-600 bg-green-100 px-3 py-1 rounded-full font-medium">مزایا</span>
            </div>

            <!-- فرم افزودن نقطه قوت -->
            <div class="bg-white border-2 border-green-300 rounded-lg p-6 mb-6">
              <div class="flex gap-3">
                <input 
                  type="text" 
                  class="flex-1 px-4 py-3 border border-green-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500 transition-colors text-right" 
                  dir="rtl" 
                  placeholder="نقطه قوت جدید را وارد کنید..." 
                />
                <button class="bg-gradient-to-r from-green-600 to-emerald-600 text-white rounded-lg px-6 py-3 text-sm hover:from-green-700 hover:to-emerald-700 transition-all duration-200 font-medium shadow-md hover:shadow-lg">
                  افزودن
                </button>
              </div>
            </div>

            <!-- لیست نقاط قوت -->
            <div class="space-y-3 mb-6">
              <div class="bg-white border border-green-300 rounded-lg p-6 hover:shadow-md transition-all duration-200">
                <div class="flex items-center justify-between">
                  <div class="flex items-center gap-3">
                    <div class="w-6 h-6 bg-green-100 rounded-full flex items-center justify-center">
                      <svg class="w-4 h-4 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                      </svg>
                    </div>
                    <span class="text-sm font-medium text-gray-800">کیفیت بالا و دوام زیاد</span>
                  </div>
                  <div class="flex gap-2">
                    <button class="text-blue-600 hover:text-blue-800 text-sm font-medium transition-colors">ویرایش</button>
                    <button class="text-red-600 hover:text-red-800 text-sm font-medium transition-colors">حذف</button>
                  </div>
                </div>
              </div>

              <div class="bg-white border border-green-300 rounded-lg p-6 hover:shadow-md transition-all duration-200">
                <div class="flex items-center justify-between">
                  <div class="flex items-center gap-3">
                    <div class="w-6 h-6 bg-green-100 rounded-full flex items-center justify-center">
                      <svg class="w-4 h-4 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                      </svg>
                    </div>
                    <span class="text-sm font-medium text-gray-800">قیمت مناسب و رقابتی</span>
                  </div>
                  <div class="flex gap-2">
                    <button class="text-blue-600 hover:text-blue-800 text-sm font-medium transition-colors">ویرایش</button>
                    <button class="text-red-600 hover:text-red-800 text-sm font-medium transition-colors">حذف</button>
                  </div>
                </div>
              </div>

              <div class="bg-white border border-green-300 rounded-lg p-6 hover:shadow-md transition-all duration-200">
                <div class="flex items-center justify-between">
                  <div class="flex items-center gap-3">
                    <div class="w-6 h-6 bg-green-100 rounded-full flex items-center justify-center">
                      <svg class="w-4 h-4 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                      </svg>
                    </div>
                    <span class="text-sm font-medium text-gray-800">گارانتی 2 ساله</span>
                  </div>
                  <div class="flex gap-2">
                    <button class="text-blue-600 hover:text-blue-800 text-sm font-medium transition-colors">ویرایش</button>
                    <button class="text-red-600 hover:text-red-800 text-sm font-medium transition-colors">حذف</button>
                  </div>
                </div>
              </div>
            </div>

            <!-- آمار نقاط قوت -->
            <div class="bg-gradient-to-r from-green-100 to-emerald-100 border border-green-300 p-6 rounded-lg">
              <div class="flex items-center justify-between">
                <span class="text-green-800 font-bold">تعداد نقاط قوت:</span>
                <span class="text-2xl font-bold text-green-600">3</span>
              </div>
            </div>
          </div>

          <!-- نقاط ضعف -->
          <div class="bg-gradient-to-br from-red-50 to-rose-50 border border-red-200 rounded-xl p-6">
            <div class="flex items-center justify-between mb-6">
              <h4 class="text-lg font-bold text-red-800 flex items-center gap-2">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                </svg>
                نقاط ضعف محصول
              </h4>
              <span class="text-sm text-red-600 bg-red-100 px-3 py-1 rounded-full font-medium">معایب</span>
            </div>

            <!-- فرم افزودن نقطه ضعف -->
            <div class="bg-white border-2 border-red-300 rounded-lg p-6 mb-6">
              <div class="flex gap-3">
                <input 
                  type="text" 
                  class="flex-1 px-4 py-3 border border-red-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500 transition-colors text-right" 
                  dir="rtl" 
                  placeholder="نقطه ضعف جدید را وارد کنید..." 
                />
                <button class="bg-gradient-to-r from-red-600 to-rose-600 text-white rounded-lg px-6 py-3 text-sm hover:from-red-700 hover:to-rose-700 transition-all duration-200 font-medium shadow-md hover:shadow-lg">
                  افزودن
                </button>
              </div>
            </div>

            <!-- لیست نقاط ضعف -->
            <div class="space-y-3 mb-6">
              <div class="bg-white border border-red-300 rounded-lg p-6 hover:shadow-md transition-all duration-200">
                <div class="flex items-center justify-between">
                  <div class="flex items-center gap-3">
                    <div class="w-6 h-6 bg-red-100 rounded-full flex items-center justify-center">
                      <svg class="w-4 h-4 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                      </svg>
                    </div>
                    <span class="text-sm font-medium text-gray-800">وزن بالای محصول</span>
                  </div>
                  <div class="flex gap-2">
                    <button class="text-blue-600 hover:text-blue-800 text-sm font-medium transition-colors">ویرایش</button>
                    <button class="text-red-600 hover:text-red-800 text-sm font-medium transition-colors">حذف</button>
                  </div>
                </div>
              </div>

              <div class="bg-white border border-red-300 rounded-lg p-6 hover:shadow-md transition-all duration-200">
                <div class="flex items-center justify-between">
                  <div class="flex items-center gap-3">
                    <div class="w-6 h-6 bg-red-100 rounded-full flex items-center justify-center">
                      <svg class="w-4 h-4 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                      </svg>
                    </div>
                    <span class="text-sm font-medium text-gray-800">نیاز به نگهداری مداوم</span>
                  </div>
                  <div class="flex gap-2">
                    <button class="text-blue-600 hover:text-blue-800 text-sm font-medium transition-colors">ویرایش</button>
                    <button class="text-red-600 hover:text-red-800 text-sm font-medium transition-colors">حذف</button>
                  </div>
                </div>
              </div>
            </div>

            <!-- آمار نقاط ضعف -->
            <div class="bg-gradient-to-r from-red-100 to-rose-100 border border-red-300 p-6 rounded-lg">
              <div class="flex items-center justify-between">
                <span class="text-red-800 font-bold">تعداد نقاط ضعف:</span>
                <span class="text-2xl font-bold text-red-600">2</span>
              </div>
            </div>
          </div>
        </div>

        <!-- تنظیمات نمایش -->
        <div class="bg-gradient-to-br from-indigo-50 to-purple-50 border border-indigo-200 rounded-xl p-6">
          <h4 class="text-lg font-bold text-indigo-800 mb-4 flex items-center gap-2">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
            </svg>
            تنظیمات نمایش نقاط قوت و ضعف
          </h4>
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <label class="bg-white border border-indigo-300 rounded-lg p-6 cursor-pointer hover:shadow-md transition-all duration-200">
              <div class="flex items-center gap-3">
                <input type="checkbox" class="w-5 h-5 text-indigo-600 border-indigo-300 rounded focus:ring-indigo-500" />
                <div>
                  <div class="text-sm font-semibold text-gray-800">نمایش در مقایسه محصولات</div>
                  <div class="text-xs text-indigo-600">قابل مشاهده در صفحه مقایسه</div>
                </div>
              </div>
            </label>
            <label class="bg-white border border-purple-300 rounded-lg p-6 cursor-pointer hover:shadow-md transition-all duration-200">
              <div class="flex items-center gap-3">
                <input type="checkbox" class="w-5 h-5 text-purple-600 border-purple-300 rounded focus:ring-purple-500" checked />
                <div>
                  <div class="text-sm font-semibold text-gray-800">نمایش در صفحه محصول</div>
                  <div class="text-xs text-purple-600">نمایش در جزئیات محصول</div>
                </div>
              </div>
            </label>
            <label class="bg-white border border-pink-300 rounded-lg p-6 cursor-pointer hover:shadow-md transition-all duration-200">
              <div class="flex items-center gap-3">
                <input type="checkbox" class="w-5 h-5 text-pink-600 border-pink-300 rounded focus:ring-pink-500" />
                <div>
                  <div class="text-sm font-semibold text-gray-800">مخفی کردن نقاط ضعف</div>
                  <div class="text-xs text-pink-600">فقط نقاط قوت نمایش داده شود</div>
                </div>
              </div>
            </label>
            <label class="bg-white border border-blue-300 rounded-lg p-6 cursor-pointer hover:shadow-md transition-all duration-200">
              <div class="flex items-center gap-3">
                <input type="checkbox" class="w-5 h-5 text-blue-600 border-blue-300 rounded focus:ring-blue-500" checked />
                <div>
                  <div class="text-sm font-semibold text-gray-800">نمایش آیکون کنار هر مورد</div>
                  <div class="text-xs text-blue-600">آیکون‌های بصری برای هر نکته</div>
                </div>
              </div>
            </label>
          </div>
        </div>

        <!-- خلاصه تحلیل -->
        <div class="bg-gradient-to-br from-slate-50 to-gray-50 border border-gray-200 rounded-xl p-6">
          <h4 class="text-lg font-bold text-gray-800 mb-4 flex items-center gap-2">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
            </svg>
            خلاصه تحلیل نقاط قوت و ضعف
          </h4>
          <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
            <div class="bg-white border border-gray-200 rounded-lg p-6 text-center hover:shadow-md transition-all duration-200">
              <div class="text-4xl font-bold text-green-600 mb-2">3</div>
              <div class="text-sm text-gray-600 font-medium">نقاط قوت</div>
              <div class="text-xs text-green-600 mt-1">مزایای محصول</div>
            </div>
            <div class="bg-white border border-gray-200 rounded-lg p-6 text-center hover:shadow-md transition-all duration-200">
              <div class="text-4xl font-bold text-red-600 mb-2">2</div>
              <div class="text-sm text-gray-600 font-medium">نقاط ضعف</div>
              <div class="text-xs text-red-600 mt-1">معایب محصول</div>
            </div>
            <div class="bg-white border border-gray-200 rounded-lg p-6 text-center hover:shadow-md transition-all duration-200">
              <div class="text-4xl font-bold text-blue-600 mb-2">60%</div>
              <div class="text-sm text-gray-600 font-medium">امتیاز کلی</div>
              <div class="text-xs text-blue-600 mt-1">بر اساس نقاط قوت و ضعف</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, inject, ref, watch, type Ref } from 'vue';
import RichTextEditor from '~/components/common/RichTextEditor.vue';
import { useProductCreateStore } from '~/stores/productCreate';

const store = useProductCreateStore()

interface Category {
  id: number | string;
  name: string;
  parent_id?: number | string | null;
}

interface Brand {
  id: number | string;
  name: string;
}

// دریافت تنظیمات بخش‌ها از کامپوننت والد
const sectionSettings = inject('sectionSettings', {
  mainInfo: true,
  technicalInfo: true,
  displaySettings: true,
  scheduling: true,
  management: true,
  strengthsWeaknesses: true
})

// Use store state and actions directly
const sections = store.sections
const productForm = store.productForm
const brands = inject<Ref<Brand[]>>('brands', ref([]))
const tinyApiKey = store.tinyApiKey
const categories = computed(() => store.categories as Category[])

// دسته‌بندی‌های اصلی (parent_id == null)
const mainCategories = computed(() => categories.value.filter(c => !c.parent_id))

// زیردسته‌ها بر اساس انتخاب دسته اصلی
const subCategories = computed(() => {
  const pid = productForm.category_id
  if (!pid) return []
  return categories.value.filter(c => c.parent_id == pid)
})

// ریست زیردسته در صورت تغییر دسته اصلی
watch(() => productForm.category_id, () => {
  productForm.sub_category_id = ''
})

const toggleSection = store.toggleSection
const generateAIContent = store.generateAIContent
</script>
