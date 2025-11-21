<template>
  <div class="p-6">
    <!-- Header Section -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-8">
      <div class="flex justify-between items-start">
        <div>
          <h1 class="text-2xl font-bold text-gray-900 mb-2">انتقال محصولات و سفارشات از وردپرس</h1>
          <p class="text-gray-600">انتقال کامل محصولات، سفارشات و کاربران از سایت وردپرس</p>
        </div>
        <button 
          class="bg-green-600 text-white px-4 py-2 rounded-md hover:bg-green-700 transition-colors duration-200 flex items-center gap-2"
          @click="refreshPage"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
          </svg>
          بروزرسانی صفحه
        </button>
      </div>
    </div>

    <!-- تنظیمات اتصال به وردپرس -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6">
      <h2 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 mr-2 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
        </svg>
        تنظیمات اتصال به وردپرس
      </h2>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">آدرس سایت وردپرس</label>
          <input 
            v-model="wordpressConfig.siteUrl" 
            type="url" 
            placeholder="https://example.com"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">کلید API وردپرس (اختیاری)</label>
          <input 
            v-model="wordpressConfig.apiKey" 
            type="password" 
            placeholder="کلید API یا Application Password (در صورت استفاده)"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">WooCommerce Consumer Key</label>
          <input 
            v-model="wordpressConfig.consumerKey" 
            type="text" 
            placeholder="ck_..."
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">WooCommerce Consumer Secret</label>
          <input 
            v-model="wordpressConfig.consumerSecret" 
            type="password" 
            placeholder="cs_..."
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>
      </div>

      <div class="mt-4">
        <button 
          :disabled="isTesting"
          class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 disabled:opacity-50"
          @click="testConnection"
        >
          {{ isTesting ? 'در حال تست...' : 'تست اتصال' }}
        </button>
      </div>
    </div>

    <!-- انتخاب موارد برای انتقال -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6">
      <h2 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 mr-2 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
        </svg>
        انتخاب موارد برای انتقال
      </h2>
      
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <!-- کارت محصولات -->
        <div 
          class="border-2 rounded-lg p-6 cursor-pointer transition-all duration-200 hover:shadow-md"
          :class="selectedItem === 'products' ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-blue-300'"
          @click="selectItem('products')"
        >
          <div class="flex items-center justify-between mb-2">
            <h3 class="font-semibold text-gray-900">محصولات</h3>
            <input 
              v-model="migrationOptions.products" 
              type="checkbox" 
              class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              @click.stop
            />
          </div>
          <p class="text-sm text-gray-600">شامل تصاویر، دسته‌بندی‌ها و ویژگی‌ها</p>
        </div>

        <!-- کارت سفارشات -->
        <div 
          class="border-2 rounded-lg p-6 cursor-pointer transition-all duration-200 hover:shadow-md"
          :class="selectedItem === 'orders' ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-blue-300'"
          @click="selectItem('orders')"
        >
          <div class="flex items-center justify-between mb-2">
            <h3 class="font-semibold text-gray-900">سفارشات</h3>
            <input 
              v-model="migrationOptions.orders" 
              type="checkbox" 
              class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              @click.stop
            />
          </div>
          <p class="text-sm text-gray-600">شامل جزئیات سفارش و وضعیت‌ها</p>
        </div>

        <!-- کارت کاربران وردپرس -->
        <div 
          class="border-2 rounded-lg p-6 cursor-pointer transition-all duration-200 hover:shadow-md"
          :class="selectedItem === 'wordpressUsers' ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-blue-300'"
          @click="selectItem('wordpressUsers')"
        >
          <div class="flex items-center justify-between mb-2">
            <h3 class="font-semibold text-gray-900">کاربران وردپرس</h3>
            <input 
              v-model="migrationOptions.wordpressUsers" 
              type="checkbox" 
              class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              @click.stop
            />
          </div>
          <p class="text-sm text-gray-600">شامل اطلاعات پروفایل</p>
        </div>

        <!-- کارت کاربران دیجیتس -->
        <div 
          class="border-2 rounded-lg p-6 cursor-pointer transition-all duration-200 hover:shadow-md"
          :class="selectedItem === 'digitsUsers' ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-blue-300'"
          @click="selectItem('digitsUsers')"
        >
          <div class="flex items-center justify-between mb-2">
            <h3 class="font-semibold text-gray-900">کاربران دیجیتس</h3>
            <input 
              v-model="migrationOptions.digitsUsers" 
              type="checkbox" 
              class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
              @click.stop
            />
          </div>
          <p class="text-sm text-gray-600">شامل اطلاعات پروفایل</p>
        </div>

        <!-- کارت دسته‌بندی‌ها -->
        <div 
          class="border-2 rounded-lg p-6 cursor-pointer transition-all duration-200 hover:shadow-md"
          :class="selectedItem === 'categories' ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-blue-300'"
          @click="selectItem('categories')"
        >
          <div class="flex items-center justify-between mb-2">
            <h3 class="font-semibold text-gray-900">دسته‌بندی‌ها</h3>
            <input 
              v-model="migrationOptions.categories" 
              :disabled="migrationOptions.products"
              type="checkbox" 
              class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded disabled:opacity-50"
              @click.stop
            />
          </div>
          <p class="text-sm text-gray-600">به همراه رسانه‌ها و محتوای مرتبط</p>
        </div>

        <!-- کارت برندها -->
        <div 
          class="border-2 rounded-lg p-6 cursor-pointer transition-all duration-200 hover:shadow-md"
          :class="selectedItem === 'brands' ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-blue-300'"
          @click="selectItem('brands')"
        >
          <div class="flex items-center justify-between mb-2">
            <h3 class="font-semibold text-gray-900">برندها</h3>
            <input 
              v-model="migrationOptions.brands" 
              :disabled="migrationOptions.products"
              type="checkbox" 
              class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded disabled:opacity-50"
              @click.stop
            />
          </div>
          <p class="text-sm text-gray-600">به همراه رسانه‌ها و محتوای مرتبط</p>
        </div>

        <!-- کارت ویژگی‌ها -->
        <div 
          class="border-2 rounded-lg p-6 cursor-pointer transition-all duration-200 hover:shadow-md"
          :class="selectedItem === 'attributes' ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-blue-300'"
          @click="selectItem('attributes')"
        >
          <div class="flex items-center justify-between mb-2">
            <h3 class="font-semibold text-gray-900">ویژگی‌ها</h3>
            <input 
              v-model="migrationOptions.attributes"
              :disabled="migrationOptions.products"
              type="checkbox"
              class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded disabled:opacity-50"
              @click.stop
            />
          </div>
          <p class="text-sm text-gray-600">فقط ویژگی‌های موجود به سیستم اضافه شوند</p>
        </div>
      </div>
    </div>

    <!-- تنظیمات مربوط به آیتم انتخاب شده -->
    <div v-if="selectedItem && selectedItemSettings && selectedItem!=='attributes'" class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6">
      <h2 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 mr-2 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"></path>
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
        </svg>
        تنظیمات {{ selectedItemSettings.title }}
      </h2>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div v-for="setting in selectedItemSettings.settings" :key="setting.key" class="flex items-center">
          <input 
            :id="setting.key" 
            v-model="migrationOptions[setting.key]" 
            type="checkbox"
            class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
          />
          <label :for="setting.key" class="ml-3 text-sm font-medium text-gray-700">
            {{ setting.label }}
          </label>
        </div>
      </div>
    </div>

    <!-- تنظیمات پیشرفته مهاجرت -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6">
      <h2 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات پیشرفته مهاجرت</h2>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تعداد آیتم در هر مرحله (Batch Size)</label>
          <input 
            v-model.number="advancedOptions.batchSize"
            type="number"
            min="1"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
          <p class="text-xs text-gray-500 mt-1">مثلاً 20 یعنی در هر اتصال 20 آیتم واکشی و ثبت می‌شود و سپس به صفحه بعد می‌رود.</p>
        </div>
      </div>
    </div>

    <!-- استخراج توضیحات متا از محصولات وردپرس -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6">
      <h2 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 mr-2 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
        </svg>
        استخراج توضیحات متا از محصولات وردپرس
      </h2>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <!-- فیلد URL محصول وردپرس -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">URL محصول وردپرس</label>
          <div class="flex gap-2">
            <input 
              v-model="metaExtractionUrl" 
              type="url" 
              placeholder="https://example.com/product/product-name"
              class="flex-1 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500"
              dir="ltr"
            />
            <button 
              :disabled="isExtractingMeta || !metaExtractionUrl.trim()"
              class="bg-purple-600 text-white px-4 py-2 rounded-md hover:bg-purple-700 disabled:opacity-50 disabled:cursor-not-allowed"
              @click="extractMetaFromWordPress"
            >
              {{ isExtractingMeta ? 'در حال استخراج...' : 'استخراج متا' }}
            </button>
          </div>
          <p class="text-xs text-gray-500 mt-1">URL صفحه محصول وردپرس را وارد کنید تا متا تگ‌های آن استخراج شود</p>
        </div>

        <!-- نمایش نتایج استخراج -->
        <div v-if="extractedMeta" class="bg-gray-50 p-6 rounded-lg">
          <h4 class="font-medium text-gray-900 mb-3">نتایج استخراج شده:</h4>
          <div class="space-y-3">
            <!-- متا تگ‌های عمومی -->
            <div v-if="extractedMeta.title || extractedMeta.description || extractedMeta.keywords">
              <h5 class="text-sm font-semibold text-gray-800 mb-2">متا تگ‌های عمومی:</h5>
              <div class="space-y-2">
                <div v-if="extractedMeta.title">
                  <label class="text-xs font-medium text-gray-600">عنوان:</label>
                  <p class="text-sm text-gray-900 bg-white p-2 rounded border">{{ extractedMeta.title }}</p>
                </div>
                <div v-if="extractedMeta.description">
                  <label class="text-xs font-medium text-gray-600">توضیحات:</label>
                  <p class="text-sm text-gray-900 bg-white p-2 rounded border">{{ extractedMeta.description }}</p>
                </div>
                <div v-if="extractedMeta.keywords">
                  <label class="text-xs font-medium text-gray-600">کلمات کلیدی:</label>
                  <p class="text-sm text-gray-900 bg-white p-2 rounded border">{{ extractedMeta.keywords }}</p>
                </div>
              </div>
            </div>

            <!-- Open Graph -->
            <div v-if="extractedMeta.ogTitle || extractedMeta.ogDescription || extractedMeta.ogImage">
              <h5 class="text-sm font-semibold text-gray-800 mb-2">Open Graph:</h5>
              <div class="space-y-2">
                <div v-if="extractedMeta.ogTitle">
                  <label class="text-xs font-medium text-gray-600">عنوان OG:</label>
                  <p class="text-sm text-gray-900 bg-white p-2 rounded border">{{ extractedMeta.ogTitle }}</p>
                </div>
                <div v-if="extractedMeta.ogDescription">
                  <label class="text-xs font-medium text-gray-600">توضیحات OG:</label>
                  <p class="text-sm text-gray-900 bg-white p-2 rounded border">{{ extractedMeta.ogDescription }}</p>
                </div>
                <div v-if="extractedMeta.ogImage">
                  <label class="text-xs font-medium text-gray-600">تصویر OG:</label>
                  <p class="text-sm text-gray-900 bg-white p-2 rounded border break-all">{{ extractedMeta.ogImage }}</p>
                </div>
              </div>
            </div>

            <!-- Yoast SEO -->
            <div v-if="extractedMeta.yoastTitle || extractedMeta.yoastDesc">
              <h5 class="text-sm font-semibold text-gray-800 mb-2">Yoast SEO:</h5>
              <div class="space-y-2">
                <div v-if="extractedMeta.yoastTitle">
                  <label class="text-xs font-medium text-gray-600">عنوان Yoast:</label>
                  <p class="text-sm text-gray-900 bg-white p-2 rounded border">{{ extractedMeta.yoastTitle }}</p>
                </div>
                <div v-if="extractedMeta.yoastDesc">
                  <label class="text-xs font-medium text-gray-600">توضیحات Yoast:</label>
                  <p class="text-sm text-gray-900 bg-white p-2 rounded border">{{ extractedMeta.yoastDesc }}</p>
                </div>
              </div>
            </div>

            <!-- RankMath -->
            <div v-if="extractedMeta.rankMathTitle || extractedMeta.rankMathDesc">
              <h5 class="text-sm font-semibold text-gray-800 mb-2">RankMath:</h5>
              <div class="space-y-2">
                <div v-if="extractedMeta.rankMathTitle">
                  <label class="text-xs font-medium text-gray-600">عنوان RankMath:</label>
                  <p class="text-sm text-gray-900 bg-white p-2 rounded border">{{ extractedMeta.rankMathTitle }}</p>
                </div>
                <div v-if="extractedMeta.rankMathDesc">
                  <label class="text-xs font-medium text-gray-600">توضیحات RankMath:</label>
                  <p class="text-sm text-gray-900 bg-white p-2 rounded border">{{ extractedMeta.rankMathDesc }}</p>
                </div>
              </div>
            </div>

            <!-- سایر اطلاعات -->
            <div v-if="extractedMeta.canonical || extractedMeta.robots">
              <h5 class="text-sm font-semibold text-gray-800 mb-2">سایر اطلاعات:</h5>
              <div class="space-y-2">
                <div v-if="extractedMeta.canonical">
                  <label class="text-xs font-medium text-gray-600">Canonical URL:</label>
                  <p class="text-sm text-gray-900 bg-white p-2 rounded border break-all">{{ extractedMeta.canonical }}</p>
                </div>
                <div v-if="extractedMeta.robots">
                  <label class="text-xs font-medium text-gray-600">Robots:</label>
                  <p class="text-sm text-gray-900 bg-white p-2 rounded border">{{ extractedMeta.robots }}</p>
                </div>
              </div>
            </div>
          </div>
          
          <div class="mt-4 flex gap-2">
            <button 
              class="bg-green-600 text-white px-3 py-2 rounded-md hover:bg-green-700 text-sm"
              @click="applyExtractedMetaToMigration"
            >
              اعمال به تنظیمات انتقال
            </button>
            <button 
              class="bg-gray-600 text-white px-3 py-2 rounded-md hover:bg-gray-700 text-sm"
              @click="clearExtractedMeta"
            >
              پاک کردن
            </button>
          </div>
        </div>
      </div>

      <!-- تنظیمات استخراج متا -->
      <div class="mt-6 border-t pt-4">
        <h4 class="font-medium text-gray-900 mb-3">تنظیمات استخراج متا:</h4>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div class="flex items-center">
            <input 
              id="extractFromYoast" 
              v-model="metaExtractionOptions.extractFromYoast" 
              type="checkbox"
              class="h-4 w-4 text-purple-600 focus:ring-purple-500 border-gray-300 rounded"
            />
            <label for="extractFromYoast" class="ml-3 text-sm font-medium text-gray-700">
              استخراج از Yoast SEO
            </label>
          </div>
          
          <div class="flex items-center">
            <input 
              id="extractFromRankMath" 
              v-model="metaExtractionOptions.extractFromRankMath" 
              type="checkbox"
              class="h-4 w-4 text-purple-600 focus:ring-purple-500 border-gray-300 rounded"
            />
            <label for="extractFromRankMath" class="ml-3 text-sm font-medium text-gray-700">
              استخراج از RankMath
            </label>
          </div>
          
          <div class="flex items-center">
            <input 
              id="extractFromOpenGraph" 
              v-model="metaExtractionOptions.extractFromOpenGraph" 
              type="checkbox"
              class="h-4 w-4 text-purple-600 focus:ring-purple-500 border-gray-300 rounded"
            />
            <label for="extractFromOpenGraph" class="ml-3 text-sm font-medium text-gray-700">
              استخراج از Open Graph
            </label>
          </div>
          
          <div class="flex items-center">
            <input 
              id="extractFromMetaTags" 
              v-model="metaExtractionOptions.extractFromMetaTags" 
              type="checkbox"
              class="h-4 w-4 text-purple-600 focus:ring-purple-500 border-gray-300 rounded"
            />
            <label for="extractFromMetaTags" class="ml-3 text-sm font-medium text-gray-700">
              استخراج از Meta Tags عمومی
            </label>
          </div>
        </div>
      </div>
    </div>

    <!-- دکمه شروع انتقال -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6">
      <div class="flex justify-between items-center">
        <div>
          <h2 class="text-lg font-semibold text-gray-900">شروع انتقال</h2>
          <p class="text-sm text-gray-600 mt-1">پس از بررسی تنظیمات، انتقال را شروع کنید</p>
        </div>
        
        <div class="flex items-center space-x-2 space-x-reverse">
          <button 
            :disabled="!canStartMigration || isMigrating"
            class="bg-green-600 text-white px-6 py-3 rounded-md hover:bg-green-700 disabled:opacity-50 font-medium"
            @click="startMigration"
          >
            {{ isMigrating ? 'در حال انتقال...' : 'شروع انتقال' }}
          </button>
          <button 
            v-if="isMigrating && migrationProgress.isActive"
            class="bg-red-600 text-white px-4 py-3 rounded-md hover:bg-red-700 font-medium"
            @click="abortMigration"
          >
            توقف انتقال
          </button>
        </div>
      </div>
    </div>

    <!-- نمایش پیشرفت و وضعیت جزئی -->
    <div v-if="migrationProgress.isActive || migrationLogs.length > 0" class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center justify-between">
        <div class="flex items-center">
          <svg class="w-5 h-5 mr-2 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
          </svg>
          پیشرفت انتقال
        </div>
        <div class="flex items-center space-x-2 space-x-reverse">
          <div class="flex items-center">
            <div
:class="[
              'w-3 h-3 rounded-full mr-2',
              migrationProgress.isActive ? 'bg-green-500 animate-pulse' : 'bg-gray-400'
            ]"></div>
            <span class="text-sm font-medium" :class="migrationProgress.isActive ? 'text-green-600' : 'text-gray-600'">
              {{ migrationStatus }}
            </span>
          </div>
          <span v-if="migrationDuration" class="text-xs text-gray-500">{{ migrationDuration }}</span>
        </div>
      </h3>
      
      <div class="space-y-6">
        <!-- وضعیت کلی -->
        <div class="bg-gray-50 p-6 rounded-lg">
          <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-4">
            <div
class="text-center cursor-pointer hover:bg-green-50 p-2 rounded transition-colors" 
                 :class="{ 'bg-green-100': selectedFilter === 'success' }"
                 @click="showProductList('success')">
              <div class="text-2xl font-bold text-green-600">{{ migrationStats.success }}</div>
              <div class="text-sm text-green-700">موفق</div>
            </div>
            <div
class="text-center cursor-pointer hover:bg-red-50 p-2 rounded transition-colors" 
                 :class="{ 'bg-red-100': selectedFilter === 'failed' }"
                 @click="showProductList('failed')">
              <div class="text-2xl font-bold text-red-600">{{ migrationStats.failed }}</div>
              <div class="text-sm text-red-700">ناموفق</div>
            </div>
            <div
class="text-center cursor-pointer hover:bg-yellow-50 p-2 rounded transition-colors" 
                 :class="{ 'bg-yellow-100': selectedFilter === 'skipped' }"
                 @click="showProductList('skipped')">
              <div class="text-2xl font-bold text-yellow-600">{{ migrationStats.skipped }}</div>
              <div class="text-sm text-yellow-700">رد شده</div>
            </div>
            <div
class="text-center cursor-pointer hover:bg-blue-50 p-2 rounded transition-colors" 
                 :class="{ 'bg-blue-100': selectedFilter === 'all' }"
                 @click="showProductList('all')">
              <div class="text-2xl font-bold text-blue-600">{{ migrationStats.total }}</div>
              <div class="text-sm text-blue-700">کل</div>
            </div>
          </div>
          
          <!-- نوار پیشرفت کلی -->
          <div v-if="migrationProgress.isActive || migrationStats.total > 0">
            <div class="flex justify-between text-sm text-gray-600 mb-2">
              <span>پیشرفت کلی</span>
              <span>{{ Math.round(getAccurateProgress()) }}%</span>
            </div>
            <div class="w-full bg-gray-200 rounded-full h-3">
              <div
class="bg-gradient-to-r from-blue-500 to-green-500 h-3 rounded-full transition-all duration-500" 
                   :style="{ width: Math.round(getAccurateProgress()) + '%' }">
              </div>
            </div>
          </div>
        </div>

        <!-- لیست محصولات فیلتر شده -->
        <div v-if="selectedFilter && filteredProducts.length > 0" class="bg-white p-6 rounded-lg border">
          <div class="flex justify-between items-center mb-4">
            <h3 class="text-lg font-semibold">
              {{ getFilterTitle() }} ({{ filteredProducts.length }})
            </h3>
            <button
class="text-gray-500 hover:text-gray-700 text-sm" 
                    @click="selectedFilter = null">
              ✕ بستن
            </button>
          </div>
          
          <div class="max-h-96 overflow-y-auto">
            <div class="space-y-2">
              <div
v-for="(product, index) in filteredProducts" 
                   :key="index"
                   class="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
                <div class="flex-1">
                  <div class="font-medium text-gray-900">{{ product.name }}</div>
                  <div class="text-sm text-gray-500">{{ product.timestamp }}</div>
                </div>
                <div class="flex items-center space-x-2">
                  <span
:class="getStatusBadgeClass(product.status)" 
                        class="px-2 py-1 rounded-full text-xs font-medium">
                    {{ getStatusText(product.status) }}
                  </span>
                  <span
v-if="product.error" 
                        class="text-red-500 text-xs max-w-xs truncate" 
                        :title="product.error">
                    {{ product.error }}
                  </span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- وضعیت مراحل مختلف -->
        <div v-if="Object.keys(migrationStages).length > 0" class="space-y-3">
          <h4 class="font-medium text-gray-900">وضعیت مراحل:</h4>
          <div class="space-y-2">
            <div
v-for="(stage, name) in migrationStages" :key="name" 
                 class="flex items-center justify-between p-3 rounded-lg border" 
                 :class="getStageClass(stage.status)">
              <div class="flex items-center">
                <div :class="getStageIconClass(stage.status)" class="w-2 h-2 rounded-full mr-3"></div>
                <span class="font-medium">{{ getStageName(name) }}</span>
              </div>
              <div class="flex items-center space-x-4 space-x-reverse text-sm">
                <span v-if="stage.current" class="text-blue-600 font-medium">{{ stage.current }}</span>
                <span v-if="stage.duration" class="text-gray-500">{{ stage.duration }}</span>
                <div v-if="stage.stats" class="flex space-x-2 space-x-reverse">
                  <span class="text-green-600">✓{{ stage.stats.success }}</span>
                  <span class="text-red-600">✗{{ stage.stats.failed }}</span>
                  <span class="text-yellow-600">⊘{{ stage.stats.skipped }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- آیتم فعلی -->
        <div v-if="migrationProgress.currentItem" class="bg-blue-50 border border-blue-200 p-6 rounded-lg">
          <div class="flex items-start justify-between">
            <div>
              <div class="font-medium text-blue-900 mb-1">در حال پردازش:</div>
              <div class="text-blue-800">{{ migrationProgress.currentItem }}</div>
              <div v-if="currentItemStartTime" class="text-xs text-blue-600 mt-1">
                زمان صرف شده: {{ getCurrentItemDuration() }}
              </div>
            </div>
            <div class="animate-spin w-5 h-5 border-2 border-blue-500 border-t-transparent rounded-full"></div>
          </div>
        </div>

        <!-- لاگ‌های زنده -->
        <div class="bg-gray-50 p-6 rounded-md max-h-96 overflow-y-auto">
          <div class="text-sm font-medium text-gray-700 mb-2 flex items-center justify-between">
            <span>لاگ‌های زنده</span>
            <button class="text-red-600 hover:text-red-700 text-xs" @click="clearLogs">پاک کردن</button>
          </div>
          <div class="space-y-2">
            <div v-for="(log, index) in migrationLogs" :key="index" class="text-xs p-2 rounded" :class="getLogClass(log.type)">
              <div class="flex items-center justify-between">
                <span class="font-medium">{{ log.timestamp }}</span>
                <span class="text-xs px-2 py-1 rounded" :class="getLogBadgeClass(log.type)">
                  {{ getLogTypeText(log.type) }}
                </span>
              </div>
              <div class="mt-1">{{ log.message }}</div>
              <div v-if="log.duration" class="text-xs text-gray-500 mt-1">
                مدت زمان: {{ log.duration }}ms
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- نتایج انتقال -->
    <div v-if="migrationResults" class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mt-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">نتایج انتقال</h3>
      
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div class="bg-green-50 p-6 rounded-md">
          <div class="text-2xl font-bold text-green-600">{{ migrationResults.success }}</div>
          <div class="text-sm text-green-700">موفق</div>
        </div>
        
        <div class="bg-red-50 p-6 rounded-md">
          <div class="text-2xl font-bold text-red-600">{{ migrationResults.failed }}</div>
          <div class="text-sm text-red-700">ناموفق</div>
        </div>
        
        <div class="bg-blue-50 p-6 rounded-md">
          <div class="text-2xl font-bold text-blue-600">{{ migrationResults.skipped }}</div>
          <div class="text-sm text-blue-700">رد شده</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { onMounted, onUnmounted } from 'vue'

definePageMeta({
  layout: 'admin-main',
  middleware: ['developer-only']
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user: _user, hasPermission: _hasPermission } = useAuth()

// تنظیمات اتصال به وردپرس
const wordpressConfig = ref({
  siteUrl: '',
  apiKey: '',
  consumerKey: '',
  consumerSecret: ''
})

// گزینه‌های انتقال
const migrationOptions = ref({
  products: false,
  orders: false,
  wordpressUsers: false,
  digitsUsers: false,
  categories: false,
  brands: false,
  attributes: false,
  replaceUsernameWithMobile: false,
  skipEmails: false,
  includeAddresses: false,
  redirectProducts: true,
  redirectCode: 301,
  redirectGroupName: 'محصولات انتقالی',
  includeSlugInURL: true,
  transferPrices: true,
  transferStock: true,
  transferVariants: true,
  transferSEOfromYoast: true,
  transferSEOfromRankMath: true,
  transferDescriptions: true,
  transferShortDescription: true,
  transferImages: true,
  transferAttributesFields: true,
  createCategoriesFromProducts: true,
  draftMode: false
})

// تنظیمات پیشرفته
const advancedOptions = ref({
  mode: 'full',
  batchSize: 20,
  includeImages: true
})

// Persist settings in localStorage (client-side only)
const LS_KEYS = {
  config: 'wp_migration.config',
  options: 'wp_migration.options',
  advanced: 'wp_migration.advanced'
}

onMounted(async () => {
  try {
    const cfg = localStorage.getItem(LS_KEYS.config)
    if (cfg) Object.assign(wordpressConfig.value, JSON.parse(cfg))
  } catch {}
  try {
    const opt = localStorage.getItem(LS_KEYS.options)
    if (opt) Object.assign(migrationOptions.value, JSON.parse(opt))
  } catch {}
  try {
    const adv = localStorage.getItem(LS_KEYS.advanced)
    if (adv) Object.assign(advancedOptions.value, JSON.parse(adv))
  } catch {}
  
  // بارگذاری لاگ‌های موجود
  try {
    const response = await $fetch('/api/admin/wordpress-migration/logs')
    
    if (response && response.success && Array.isArray(response.logs)) {
      // اگر لاگ‌هایی موجود است، یعنی فرآیند در حال اجرا است
      if (response.logs.length > 0) {
        isMigrating.value = true
        migrationProgress.value.isActive = true
        migrationStartTime.value = new Date()
        
        response.logs.forEach((log) => {
          let message = log.message
          try {
            // بهبود decode برای کاراکترهای فارسی
            if (message.includes('Ø') || message.includes('Û') || message.includes('Ø§') || message.includes('Ø±')) {
              message = decodeURIComponent(escape(message))
            }
          } catch (_e) {
            // Failed to decode existing log
          }
          
          addLog(log.type, message)
        })
        
        // شروع polling برای لاگ‌های جدید
        startLogPolling(response.logs.length)
      }
    }
  } catch (error) {
    console.error('❌ Error loading existing logs:', error)
  }
  
  // شروع auto-refresh
  startAutoRefresh()
})

// Cleanup on unmount
onUnmounted(() => {
  stopAutoRefresh()
})

// تابع جداگانه برای polling
const startLogPolling = (initialLogCount = 0) => {
  let lastLogCount = initialLogCount
  let lastLogUpdateTime = Date.now()
  
  const fetchLogs = async () => {
    try {
      const response = await $fetch('/api/admin/wordpress-migration/logs')
      
      if (response && response.success && Array.isArray(response.logs)) {
        // همیشه آخرین لاگ را بررسی کن حتی اگر تعداد تغییر نکرده باشد
        if (response.logs.length > 0) {
          const lastLog = response.logs[response.logs.length - 1]
          
          // اگر لاگ جدیدی آمده یا آخرین لاگ تغییر کرده
          if (response.logs.length > lastLogCount || 
              (lastLogCount > 0 && lastLog.message !== migrationLogs.value[migrationLogs.value.length - 1]?.message)) {
            
            const newLogs = response.logs.slice(lastLogCount)
            
            newLogs.forEach((log) => {
              let message = log.message
              try {
                // بهبود decode برای کاراکترهای فارسی
                if (message.includes('Ø') || message.includes('Û') || message.includes('Ø§') || message.includes('Ø±')) {
                  message = decodeURIComponent(escape(message))
                }
              } catch (_e) {
                // Failed to decode new log
              }
              
              addLog(log.type, message)
              
              if (message.includes('=== انتقال تکمیل شد ===') || 
                  (message.includes('فرآیند انتقال با موفقیت تکمیل شد') && log.type === 'success')) {
                progressInternal.value = 100
                isMigrating.value = false
                migrationProgress.value.isActive = false
                migrationProgress.value.currentItem = 'انتقال تکمیل شد'
                
                setTimeout(() => {
                  alert('انتقال با موفقیت تکمیل شد!')
                }, 1000)
              } else if (message.includes('انتقال محصولات تکمیل شد')) {
                migrationProgress.value.currentItem = 'انتقال محصولات تکمیل شد'
              }
            })
            lastLogCount = response.logs.length
            lastLogUpdateTime = Date.now() // به‌روزرسانی زمان آخرین لاگ
          } else {
            // اگر لاگ جدیدی نیامده، اما تعداد کل تغییر کرده، به‌روزرسانی کن
            if (response.logs.length !== lastLogCount) {
              lastLogCount = response.logs.length
            }
          }
        }
      }
    } catch (error) {
      console.error('🚨 Polling error:', error)
    }
  }
  
  const logInterval = setInterval(async () => {
    if (!isMigrating.value) {
      clearInterval(logInterval)
      return
    }
    
    await fetchLogs()
    
    // اگر بیش از 30 ثانیه لاگ جدیدی نیامده، وضعیت را به‌روزرسانی کن
    if (Date.now() - lastLogUpdateTime > 30000) {
      migrationProgress.value.currentItem = 'در حال پردازش...'
      lastLogUpdateTime = Date.now()
    }
  }, 2000)
  
  return logInterval
}

watch(wordpressConfig, (v)=>{
  try { localStorage.setItem(LS_KEYS.config, JSON.stringify(v)) } catch {}
}, { deep: true })

watch(migrationOptions, (v)=>{
  try { localStorage.setItem(LS_KEYS.options, JSON.stringify(v)) } catch {}
}, { deep: true })

watch(advancedOptions, (v)=>{
  try { localStorage.setItem(LS_KEYS.advanced, JSON.stringify(v)) } catch {}
}, { deep: true })

// وضعیت‌های مختلف
const isTesting = ref(false)
const isMigrating = ref(false)
const selectedItem = ref(null)

// متغیرهای مربوط به استخراج متا
const metaExtractionUrl = ref('')
const isExtractingMeta = ref(false)
const extractedMeta = ref(null)
const metaExtractionOptions = ref({
  extractFromYoast: true,
  extractFromRankMath: true,
  extractFromOpenGraph: true,
  extractFromMetaTags: true
})

// پیشرفت انتقال
const migrationProgress = ref({
  isActive: false,
  overall: 0,
  currentItem: '',
  logs: []
})
// پیشرفت داخلی تدریجی (شناور)
const progressInternal = ref(0)

// نتایج انتقال
const migrationResults = ref(null)

// لاگ‌های زنده
const migrationLogs = ref([])

// آمار جزئی انتقال
const migrationStats = ref({
  success: 0,
  failed: 0,
  skipped: 0,
  total: 0
})

// وضعیت مراحل مختلف
const migrationStages = ref({})

// زمان‌سنجی
const migrationStartTime = ref(null)
const currentItemStartTime = ref(null)

// فیلتر و لیست محصولات
const selectedFilter = ref(null)
const productList = ref([])

// وضعیت فعلی
const migrationStatus = computed(() => {
  if (migrationProgress.value.isActive) {
    return 'در حال اجرا'
  } else if (migrationLogs.value.length > 0) {
    const lastLog = migrationLogs.value[0]
    if (lastLog.type === 'success' && lastLog.message.includes('تکمیل')) {
      return 'تکمیل شده'
    } else if (lastLog.type === 'error') {
      return 'متوقف شده (خطا)'
    } else if (lastLog.type === 'warning' && lastLog.message.includes('متوقف')) {
      return 'متوقف شده (کاربر)'
    }
    return 'آماده'
  }
  return 'آماده'
})

// فیلتر محصولات بر اساس انتخاب
const filteredProducts = computed(() => {
  if (!selectedFilter.value) return []
  
  return productList.value.filter(product => {
    switch (selectedFilter.value) {
      case 'success':
        return product.status === 'success'
      case 'failed':
        return product.status === 'failed'
      case 'skipped':
        return product.status === 'skipped'
      case 'all':
        return true
      default:
        return false
    }
  })
})

// مدت زمان کل
const migrationDuration = computed(() => {
  if (!migrationStartTime.value) return null
  const now = new Date()
  const diff = now - migrationStartTime.value
  return formatDuration(diff)
})

// محاسبه پیشرفت کلی (بر اساس پیشرفت داخلی)
const calculateOverallProgress = computed(() => {
  return Math.min(100, Math.round(progressInternal.value))
})

// محاسبه دقیق پیشرفت بر اساس تعداد کل محصولات
const getAccurateProgress = () => {
  // اگر تعداد کل مشخص شده، همیشه از محاسبه دقیق استفاده کن
  if (migrationStats.value.total > 0) {
    const processed = migrationStats.value.success + migrationStats.value.failed + migrationStats.value.skipped
    return Math.min(100, Math.round((processed / migrationStats.value.total) * 100))
  }
  
  // در غیر این صورت از پیشرفت داخلی استفاده کن
  return Math.min(100, Math.round(progressInternal.value))
}

// watch برای به‌روزرسانی پیشرفت کلی
watch([calculateOverallProgress, migrationLogs], () => {
  migrationProgress.value.overall = calculateOverallProgress.value
}, { immediate: true })

// تنظیمات مربوط به هر آیتم
const itemSettings = {
  products: {
    title: 'محصولات',
    settings: [
      { key: 'transferPrices', label: 'انتقال قیمت (قیمت اصلی/تخفیف)' },
      { key: 'transferStock', label: 'انتقال موجودی و وضعیت انبار' },
      { key: 'transferVariants', label: 'انتقال تنوع‌های محصول (Simple/Variable)' },
      { key: 'transferDescriptions', label: 'انتقال توضیحات کامل' },
      { key: 'transferShortDescription', label: 'انتقال توضیح کوتاه' },
      { key: 'transferImages', label: 'انتقال تصاویر محصول' },
      { key: 'transferAttributesFields', label: 'انتقال ساختار ویژگی‌ها از وردپرس' },
      { key: 'transferSEOfromYoast', label: 'انتقال SEO از Yoast' },
      { key: 'transferSEOfromRankMath', label: 'انتقال SEO از RankMath' },
      { key: 'createCategoriesFromProducts', label: 'ساخت/انتساب دسته‌بندی‌ها هنگام ایمپورت' },
      { key: 'redirectProducts', label: 'ریدایرکت به ساختار جدید' },
      { key: 'includeSlugInURL', label: 'نمایش اسلاگ در URL مقصد' },
      { key: 'draftMode', label: 'به عنوان پیش‌نویس ثبت شوند' }
    ]
  },
  orders: {
    title: 'سفارشات',
    settings: [
      { key: 'includeAddresses', label: 'ثبت آدرس ارسال از سفارش (برای هر کاربر)' }
    ]
  },
  wordpressUsers: {
    title: 'کاربران وردپرس',
    settings: [
      { key: 'replaceUsernameWithMobile', label: 'جایگزینی یوزرنیم با شماره موبایل' },
      { key: 'skipEmails', label: 'عدم انتقال ایمیل (استفاده از ایمیل سیستمی)' },
      { key: 'includeAddresses', label: 'انتقال آدرس‌های کاربر (billing/shipping)' }
    ]
  },
  digitsUsers: {
    title: 'کاربران دیجیتس',
    settings: [
      { key: 'replaceUsernameWithMobile', label: 'جایگزینی یوزرنیم با شماره موبایل' },
      { key: 'skipEmails', label: 'عدم انتقال ایمیل (استفاده از ایمیل سیستمی)' },
      { key: 'includeAddresses', label: 'انتقال آدرس‌های کاربر (billing)' }
    ]
  },
  categories: {
    title: 'دسته‌بندی‌ها',
    settings: [
      { key: 'includeSlugInURL', label: 'نمایش اسلاگ در URL مقصد' },
      { key: 'draftMode', label: 'به عنوان پیش‌نویس ثبت شوند' }
    ]
  },
  brands: {
    title: 'برندها',
    settings: [
      { key: 'includeSlugInURL', label: 'نمایش اسلاگ در URL مقصد' },
      { key: 'draftMode', label: 'به عنوان پیش‌نویس ثبت شوند' }
    ]
  },
  attributes: {
    title: 'ویژگی‌ها',
    settings: [
    ]
  }
}

// محاسبه تنظیمات آیتم انتخاب شده
const selectedItemSettings = computed(() => {
  return selectedItem.value ? itemSettings[selectedItem.value] : null
})

// انتخاب آیتم
const selectItem = (item) => {
  selectedItem.value = selectedItem.value === item ? null : item
}

// محاسبه امکان شروع انتقال
const canStartMigration = computed(() => {
  const hasWooKeys = !!wordpressConfig.value.consumerKey && !!wordpressConfig.value.consumerSecret
  const hasApiKey = !!wordpressConfig.value.apiKey
  const wantsAny = (
    migrationOptions.value.products ||
    migrationOptions.value.orders ||
    migrationOptions.value.wordpressUsers ||
    migrationOptions.value.digitsUsers ||
    migrationOptions.value.categories ||
    migrationOptions.value.brands ||
    migrationOptions.value.attributes
  )
  if (!wordpressConfig.value.siteUrl || !wantsAny) return false
  // اگر سفارشات/محصولات/کاربران دیجیتس/ویژگی‌ها انتخاب شده باشند، کلیدهای ووکامرس الزامی است
  const needsWoo = (
    migrationOptions.value.orders ||
    migrationOptions.value.products ||
    migrationOptions.value.digitsUsers ||
    migrationOptions.value.attributes
  )
  if (needsWoo) return hasWooKeys
  // برای سایر موارد (مثلاً فقط تست API عمومی) وجود یکی از کلیدها کفایت می‌کند
  return hasWooKeys || hasApiKey
})

// تست اتصال به وردپرس
const testConnection = async () => {
  const hasWooKeys = !!wordpressConfig.value.consumerKey && !!wordpressConfig.value.consumerSecret
  const hasApiKey = !!wordpressConfig.value.apiKey
  if (!wordpressConfig.value.siteUrl || (!hasApiKey && !hasWooKeys)) {
    alert('لطفاً آدرس سایت و یکی از گزینه‌ها را وارد کنید: (کلید API) یا (Consumer Key/Secret)')
    return
  }

  isTesting.value = true
  addLog('info', 'شروع تست اتصال به وردپرس...')
  
  try {
    const startTime = Date.now()
    const response = await $fetch('/api/admin/wordpress-migration/test-connection', {
      method: 'POST',
      body: {
        siteUrl: wordpressConfig.value.siteUrl,
        apiKey: wordpressConfig.value.apiKey,
        consumerKey: wordpressConfig.value.consumerKey,
        consumerSecret: wordpressConfig.value.consumerSecret
      }
    })
    
    const duration = Date.now() - startTime
    
    if (response.success) {
      addLog('success', 'اتصال موفقیت‌آمیز بود!', duration)
      alert('اتصال موفقیت‌آمیز بود!')
    } else {
      addLog('error', 'خطا در اتصال: ' + response.message, duration)
      alert('خطا در اتصال: ' + response.message)
    }
  } catch (error) {
    addLog('error', 'خطا در برقراری اتصال: ' + error.message)
    alert('خطا در برقراری اتصال: ' + error.message)
  } finally {
    isTesting.value = false
  }
}

// توابع کمکی
const formatDuration = (ms) => {
  const seconds = Math.floor(ms / 1000)
  const minutes = Math.floor(seconds / 60)
  const hours = Math.floor(minutes / 60)
  
  if (hours > 0) {
    return `${hours}س ${minutes % 60}د ${seconds % 60}ث`
  } else if (minutes > 0) {
    return `${minutes}د ${seconds % 60}ث`
  } else {
    return `${seconds}ث`
  }
}

const getCurrentItemDuration = () => {
  if (!currentItemStartTime.value) return '0ث'
  const now = new Date()
  const diff = now - currentItemStartTime.value
  return formatDuration(diff)
}

const getStageName = (stage) => {
  const names = {
    products: 'محصولات',
    orders: 'سفارشات',
    categories: 'دسته‌بندی‌ها',
    brands: 'برندها',
    attributes: 'ویژگی‌ها',
    wordpressUsers: 'کاربران وردپرس',
    digitsUsers: 'کاربران دیجیتس'
  }
  return names[stage] || stage
}

const getStageClass = (status) => {
  switch (status) {
    case 'active': return 'border-blue-300 bg-blue-50'
    case 'completed': return 'border-green-300 bg-green-50'
    case 'error': return 'border-red-300 bg-red-50'
    case 'waiting': return 'border-gray-300 bg-gray-50'
    default: return 'border-gray-200 bg-white'
  }
}

const getStageIconClass = (status) => {
  switch (status) {
    case 'active': return 'bg-blue-500 animate-pulse'
    case 'completed': return 'bg-green-500'
    case 'error': return 'bg-red-500'
    case 'waiting': return 'bg-yellow-500'
    default: return 'bg-gray-400'
  }
}

// شروع فرآیند انتقال
const startMigration = async () => {
  if (!canStartMigration.value) {
    const needsWoo = (
      migrationOptions.value.orders ||
      migrationOptions.value.products ||
      migrationOptions.value.digitsUsers ||
      migrationOptions.value.attributes
    )
    if (needsWoo && (!wordpressConfig.value.consumerKey || !wordpressConfig.value.consumerSecret)) {
      alert('برای انتقال سفارشات/محصولات/کاربران/ویژگی‌ها باید WooCommerce Consumer Key/Secret را وارد کنید')
    } else {
      alert('لطفاً تنظیمات را تکمیل کنید')
    }
    return
  }

  isMigrating.value = true
  migrationProgress.value.isActive = true
  migrationProgress.value.overall = 0
  progressInternal.value = 0
  migrationProgress.value.logs = []
  migrationResults.value = null
  migrationLogs.value = []
  
  // ریست کردن آمار و زمان‌سنجی
  migrationStats.value = { success: 0, failed: 0, skipped: 0, total: 0 }
  migrationStages.value = {}
  migrationStartTime.value = new Date()
  currentItemStartTime.value = null

  addLog('info', 'شروع فرآیند انتقال...')

  // شروع polling برای لاگ‌های زنده
  // شروع polling با تابع جداگانه
  const logInterval = startLogPolling(0)

  try {
    const startTime = Date.now()
    const controller = new AbortController()
    window.migrationController = controller

    const response = await $fetch('/api/admin/wordpress-migration/start', {
      method: 'POST',
      body: {
        config: wordpressConfig.value,
        options: migrationOptions.value,
        advanced: advancedOptions.value
      },
      signal: controller.signal
    })

    const duration = Date.now() - startTime

    if (response.success) {
      // اگر status برابر "started" است، یعنی فرآیند شروع شده نه تمام شده
      if (response.status === 'started') {
        addLog('info', response.message, duration)
        // فرآیند شروع شده، منتظر لاگ‌های بیشتر می‌مانیم
        migrationProgress.value.currentItem = 'در حال شروع فرآیند انتقال...'
      } else {
        // فرآیند واقعاً تمام شده
        migrationResults.value = response.results
        progressInternal.value = 100
        addLog('success', 'انتقال با موفقیت انجام شد!', duration)
        alert('انتقال با موفقیت انجام شد!')
      }
    } else {
      addLog('error', 'خطا در انتقال: ' + response.message, duration)
      alert('خطا در انتقال: ' + response.message)
      isMigrating.value = false
      migrationProgress.value.isActive = false
    }
  } catch (error) {
    if (error.name === 'AbortError') {
      addLog('warning', 'عملیات توسط کاربر متوقف شد')
      isMigrating.value = false
      migrationProgress.value.isActive = false
    } else {
      addLog('error', 'خطا در شروع انتقال: ' + error.message)
      alert('خطا در شروع انتقال: ' + error.message)
      isMigrating.value = false
      migrationProgress.value.isActive = false
    }
  } finally {
    // فقط در صورت خطا یا تکمیل، وضعیت را تغییر دهیم
    // اگر migration هنوز در حال اجرا است، وضعیت را تغییر ندهیم
    window.migrationController = null
    // clearInterval فقط در صورت خطا یا تکمیل
    if (!isMigrating.value) {
      clearInterval(logInterval)
    }
  }
}

// توقف دستی فرآیند
const abortMigration = async () => {
  try {
    await $fetch('/api/admin/wordpress-migration/abort', {
      method: 'POST'
    })
  } catch (_error) {
    // ignore abort errors
  }
  
  if (window.migrationController) {
    window.migrationController.abort()
  }
  isMigrating.value = false
  migrationProgress.value.isActive = false
  addLog('warning', 'عملیات توسط کاربر متوقف شد')
}

// اضافه کردن لاگ
const addLog = (type, message, duration = null) => {
  const timestamp = new Date().toLocaleTimeString('fa-IR')
  migrationLogs.value.unshift({
    type,
    message,
    timestamp,
    duration
  })

  // نگه داشتن فقط 100 لاگ آخر
  if (migrationLogs.value.length > 100) {
    migrationLogs.value = migrationLogs.value.slice(0, 100)
  }

  // تحلیل پیام و به‌روزرسانی آمار
  updateStatsFromMessage(message, type)
  updateStagesFromMessage(message, type)

  // به‌روزرسانی آیتم فعلی و زمان‌سنجی
  updateCurrentItem(message)

  // به‌روزرسانی پیشرفت داخلی
  let inc = 0
  if (type === 'success') inc = 2
  else if (type === 'info') inc = 1
  else if (type === 'warning') inc = 0.5
  else inc = 0.3

  if (message.includes('تکمیل شد') || message.includes('انجام شد') || message.includes('successfully')) {
    progressInternal.value = 100
  } else {
    if (migrationProgress.value.isActive) {
      // اگر تعداد کل محصولات مشخص شده، پیشرفت داخلی را محدود نکن
      if (migrationStats.value.total > 0) {
        // پیشرفت داخلی را بر اساس تعداد واقعی محصولات پردازش شده محاسبه کن
        const processed = migrationStats.value.success + migrationStats.value.failed + migrationStats.value.skipped
        progressInternal.value = Math.min(100, Math.round((processed / migrationStats.value.total) * 100))
      } else {
        progressInternal.value = Math.min(95, progressInternal.value + inc)
      }
    }
  }

  migrationProgress.value.overall = calculateOverallProgress.value
}

// تحلیل پیام و به‌روزرسانی آمار
const updateStatsFromMessage = (message, type) => {
  // استخراج آمار از پیام‌های گروهی (برای آمار کلی)
  const successMatch = message.match(/(\d+)\s*(?:محصول|کاربر|سفارش|ویژگی|دسته|برند).*(?:انتقال یافت|موفق|ایجاد شد)/)
  const failedMatch = message.match(/(\d+)\s*(?:محصول|کاربر|سفارش|ویژگی|دسته|برند).*(?:ناموفق|خطا|شکست)/)
  const skippedMatch = message.match(/(\d+)\s*(?:محصول|کاربر|سفارش|ویژگی|دسته|برند).*(?:رد شد|تکراری|موجود)/)
  
  if (successMatch) {
    migrationStats.value.success += parseInt(successMatch[1])
  }
  if (failedMatch) {
    migrationStats.value.failed += parseInt(failedMatch[1])
  }
  if (skippedMatch) {
    migrationStats.value.skipped += parseInt(skippedMatch[1])
  }
  
  // استخراج تعداد کل محصولات از پیام "تعداد کل محصولات برای انتقال"
  const totalProductsMatch = message.match(/تعداد کل محصولات برای انتقال: (\d+)/)
  if (totalProductsMatch) {
    migrationStats.value.total = parseInt(totalProductsMatch[1])
  }
  
  // استخراج پیشرفت دقیق از پیام‌های "در حال انتقال محصول X/Y (Z%)"
  const progressMatch = message.match(/در حال انتقال محصول (\d+)\/(\d+) \((\d+)%\)/)
  if (progressMatch) {
    const current = parseInt(progressMatch[1])
    const total = parseInt(progressMatch[2])
    const percent = parseInt(progressMatch[3])
    
    // به‌روزرسانی تعداد کل اگر هنوز تنظیم نشده
    if (migrationStats.value.total === 0) {
      migrationStats.value.total = total
    }
    
    // به‌روزرسانی پیشرفت داخلی بر اساس درصد واقعی
    progressInternal.value = percent
    
    // محاسبه آمار بر اساس پیشرفت واقعی
    // فرض می‌کنیم که محصولات پردازش شده شامل موفق + ناموفق + رد شده هستند
    const processed = current - 1 // محصول فعلی هنوز پردازش نشده
    const estimatedSuccess = Math.round(processed * 0.8) // فرض: 80% موفق
    const estimatedFailed = Math.round(processed * 0.1) // فرض: 10% ناموفق  
    const estimatedSkipped = processed - estimatedSuccess - estimatedFailed // بقیه رد شده
    
    // فقط اگر آمار هنوز صفر است، تخمین بزن
    if (migrationStats.value.success === 0 && migrationStats.value.failed === 0 && migrationStats.value.skipped === 0) {
      migrationStats.value.success = Math.max(0, estimatedSuccess)
      migrationStats.value.failed = Math.max(0, estimatedFailed)
      migrationStats.value.skipped = Math.max(0, estimatedSkipped)
    }
  }
  
  // بررسی پیام تکمیل انتقال محصولات
  const completionMatch = message.match(/انتقال محصولات تکمیل شد\. تعداد کل پردازش شده: (\d+)/)
  if (completionMatch) {
    progressInternal.value = 100
  }
  
  // تخمین کل بر اساس پیام‌های "در حال پردازش X محصول" (فقط اگر تعداد کل هنوز مشخص نشده)
  if (migrationStats.value.total === 0) {
    const processingMatch = message.match(/در حال پردازش (\d+)/)
    if (processingMatch) {
      const count = parseInt(processingMatch[1])
      migrationStats.value.total = Math.max(migrationStats.value.total, migrationStats.value.success + migrationStats.value.failed + migrationStats.value.skipped + count)
    }
  }
  
  // شمارش محصولات موفق بر اساس پیام‌های فردی (فقط برای لیست محصولات)
  if (type === 'success' && message.includes('محصول') && message.includes('با موفقیت انتقال یافت')) {
    // استخراج نام محصول از پیام با regex دقیق‌تر
    const productMatch = message.match(/محصول (.+?) با موفقیت انتقال یافت/)
    const productName = productMatch ? productMatch[1] : 'نامشخص'
    productList.value.push({
      name: productName,
      status: 'success',
      timestamp: timestamp,
      error: null
    })
    // افزایش آمار موفق
    migrationStats.value.success++
  }
  
  // شمارش محصولات رد شده بر اساس پیام‌های فردی (فقط برای لیست محصولات)
  if (type === 'info' && message.includes('محصول') && message.includes('قبلاً وجود دارد')) {
    // استخراج نام محصول از پیام با regex دقیق‌تر
    const productMatch = message.match(/محصول (.+?) قبلاً وجود دارد/)
    const productName = productMatch ? productMatch[1] : 'نامشخص'
    productList.value.push({
      name: productName,
      status: 'skipped',
      timestamp: timestamp,
      error: null
    })
    // افزایش آمار رد شده
    migrationStats.value.skipped++
  }
  
  // شمارش محصولات ناموفق بر اساس پیام‌های خطا (فقط برای لیست محصولات)
  if (type === 'error' && message.includes('محصول') && (message.includes('خطا در انتقال') || message.includes('خطا در commit'))) {
    // استخراج نام محصول و خطا از پیام با regex دقیق‌تر
    let productName = 'نامشخص'
    let error = message
    
    if (message.includes('خطا در انتقال محصول')) {
      const match = message.match(/خطا در انتقال محصول '(.+?)': (.+)/)
      if (match) {
        productName = match[1]
        error = match[2]
      }
    } else if (message.includes('خطا در commit محصول')) {
      const match = message.match(/خطا در commit محصول (.+?): (.+)/)
      if (match) {
        productName = match[1]
        error = match[2]
      }
    }
    
    productList.value.push({
      name: productName,
      status: 'failed',
      timestamp: timestamp,
      error: error
    })
    // افزایش آمار ناموفق
    migrationStats.value.failed++
  }
  
  // شمارش سفارشات موفق
  if (type === 'success' && message.includes('سفارشات:') && message.includes('انتقال یافت')) {
    const orderMatch = message.match(/سفارشات: (\d+) انتقال یافت/)
    if (orderMatch) {
      migrationStats.value.success += parseInt(orderMatch[1])
    }
  }
  
  // شمارش سفارشات رد شده
  if (type === 'success' && message.includes('سفارشات:') && message.includes('رد شد')) {
    const orderMatch = message.match(/سفارشات: \d+ انتقال یافت، (\d+) رد شد/)
    if (orderMatch) {
      migrationStats.value.skipped += parseInt(orderMatch[1])
    }
  }
  
  // شمارش سفارشات ناموفق
  if (type === 'success' && message.includes('سفارشات:') && message.includes('ناموفق')) {
    const orderMatch = message.match(/سفارشات: \d+ انتقال یافت، \d+ رد شد، (\d+) ناموفق/)
    if (orderMatch) {
      migrationStats.value.failed += parseInt(orderMatch[1])
    }
  }
  
  // شمارش کاربران دیجیتس
  if (type === 'success' && message.includes('کاربران دیجیتس:')) {
    const userMatch = message.match(/کاربران دیجیتس: (\d+) انتقال یافت، (\d+) رد شد، (\d+) ناموفق/)
    if (userMatch) {
      migrationStats.value.success += parseInt(userMatch[1])
      migrationStats.value.skipped += parseInt(userMatch[2])
      migrationStats.value.failed += parseInt(userMatch[3])
    }
  }
  
  // شمارش ویژگی‌ها
  if (type === 'success' && message.includes('ویژگی‌ها:')) {
    const attrMatch = message.match(/ویژگی‌ها: (\d+) ویژگی و (\d+) مقدار انتقال یافت \( (\d+) رد، (\d+) ناموفق \)/)
    if (attrMatch) {
      migrationStats.value.success += parseInt(attrMatch[1]) + parseInt(attrMatch[2])
      migrationStats.value.skipped += parseInt(attrMatch[3])
      migrationStats.value.failed += parseInt(attrMatch[4])
    }
  }
}

// به‌روزرسانی وضعیت مراحل
const updateStagesFromMessage = (message, type) => {
  const stages = ['products', 'orders', 'categories', 'brands', 'attributes', 'wordpressUsers', 'digitsUsers']
  
  stages.forEach(stage => {
    const stageName = getStageName(stage)
    if (message.includes(stageName) || message.includes(stage)) {
      if (!migrationStages.value[stage]) {
        migrationStages.value[stage] = {
          status: 'waiting',
          startTime: null,
          duration: null,
          stats: { success: 0, failed: 0, skipped: 0 },
          current: null
        }
      }
      
      const stageObj = migrationStages.value[stage]
      
      if (message.includes('شروع')) {
        stageObj.status = 'active'
        stageObj.startTime = new Date()
      } else if (message.includes('تکمیل') || message.includes('انجام شد')) {
        stageObj.status = 'completed'
        if (stageObj.startTime) {
          stageObj.duration = formatDuration(new Date() - stageObj.startTime)
        }
      } else if (type === 'error') {
        stageObj.status = 'error'
      }
      
      // استخراج آمار مرحله‌ای
      const stageSuccessMatch = message.match(/(\d+).*انتقال یافت/)
      const stageFailedMatch = message.match(/(\d+).*ناموفق/)
      const stageSkippedMatch = message.match(/(\d+).*رد شد/)
      
      if (stageSuccessMatch) stageObj.stats.success += parseInt(stageSuccessMatch[1])
      if (stageFailedMatch) stageObj.stats.failed += parseInt(stageFailedMatch[1])
      if (stageSkippedMatch) stageObj.stats.skipped += parseInt(stageSkippedMatch[1])
    }
  })
}

// به‌روزرسانی آیتم فعلی
const updateCurrentItem = (message) => {
  if (message.includes('در حال انتقال محصول') || message.includes('در حال پردازش')) {
    migrationProgress.value.currentItem = message
    currentItemStartTime.value = new Date()
  } else if (message.includes('شروع انتقال')) {
    const stage = message.replace('شروع انتقال ', '').replace('...', '')
    migrationProgress.value.currentItem = `آماده‌سازی ${stage}`
    currentItemStartTime.value = new Date()
  }
}

// نمایش لیست محصولات
const showProductList = (filter) => {
  selectedFilter.value = filter
}

// دریافت عنوان فیلتر
const getFilterTitle = () => {
  switch (selectedFilter.value) {
    case 'success':
      return 'محصولات موفق'
    case 'failed':
      return 'محصولات ناموفق'
    case 'skipped':
      return 'محصولات رد شده'
    case 'all':
      return 'همه محصولات'
    default:
      return ''
  }
}

// دریافت کلاس badge وضعیت
const getStatusBadgeClass = (status) => {
  switch (status) {
    case 'success':
      return 'bg-green-100 text-green-800'
    case 'failed':
      return 'bg-red-100 text-red-800'
    case 'skipped':
      return 'bg-yellow-100 text-yellow-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

// دریافت متن وضعیت
const getStatusText = (status) => {
  switch (status) {
    case 'success':
      return 'موفق'
    case 'failed':
      return 'ناموفق'
    case 'skipped':
      return 'رد شده'
    default:
      return 'نامشخص'
  }
}

// پاک کردن لاگ‌ها
const clearLogs = () => {
  migrationLogs.value = []
  productList.value = []
  selectedFilter.value = null
}

// تابع refresh برای بارگذاری مجدد لاگ‌ها و وضعیت
const refreshPage = async () => {
  // Reset states
  isMigrating.value = false
  migrationProgress.value.isActive = false
  migrationProgress.value.overall = 0
  migrationProgress.value.currentItem = ''
  migrationLogs.value = []
  migrationStats.value = { success: 0, failed: 0, skipped: 0, total: 0 }
  migrationStages.value = {}
  migrationStartTime.value = null
  progressInternal.value = 0
  productList.value = []
  selectedFilter.value = null
  
  // Reload existing logs
  try {
    const response = await $fetch('/api/admin/wordpress-migration/logs')
    
    if (response && response.success && Array.isArray(response.logs)) {
      // اگر لاگ‌هایی موجود است، یعنی فرآیند در حال اجرا است
      if (response.logs.length > 0) {
        isMigrating.value = true
        migrationProgress.value.isActive = true
        migrationStartTime.value = new Date()
        
        response.logs.forEach((log) => {
          let message = log.message
          try {
            if (message.includes('Ø') || message.includes('Û')) {
              message = decodeURIComponent(escape(message))
            }
          } catch (_e) {
            // Failed to decode refreshed log
          }
          
          addLog(log.type, message)
        })
        
        // شروع polling برای لاگ‌های جدید
        startLogPolling(response.logs.length)
      }
    }
  } catch (error) {
    console.error('❌ Error refreshing logs:', error)
  }
}

// Auto-refresh هر 30 ثانیه
let autoRefreshInterval = null

const startAutoRefresh = () => {
  if (autoRefreshInterval) {
    clearInterval(autoRefreshInterval)
  }
  
  autoRefreshInterval = setInterval(() => {
    if (!isMigrating.value) {
      // فقط اگر migration در حال اجرا نیست، auto-refresh کن
      refreshPage()
    }
  }, 30000) // هر 30 ثانیه
}

const stopAutoRefresh = () => {
  if (autoRefreshInterval) {
    clearInterval(autoRefreshInterval)
    autoRefreshInterval = null
  }
}

// توابع استخراج متا
const extractMetaFromWordPress = async () => {
  if (!metaExtractionUrl.value.trim()) {
    alert('لطفاً URL محصول وردپرس را وارد کنید')
    return
  }

  isExtractingMeta.value = true
  extractedMeta.value = null

  try {
    const response = await $fetch('/api/admin/wordpress-migration/extract-meta', {
      method: 'POST',
      body: {
        productUrl: metaExtractionUrl.value,
        options: {
          extractFromYoast: metaExtractionOptions.value.extractFromYoast,
          extractFromRankMath: metaExtractionOptions.value.extractFromRankMath,
          extractFromOpenGraph: metaExtractionOptions.value.extractFromOpenGraph,
          extractFromMetaTags: metaExtractionOptions.value.extractFromMetaTags
        }
      }
    })

    if (response.success) {
      extractedMeta.value = {
        title: response.data.title || '',
        description: response.data.description || '',
        keywords: response.data.keywords || '',
        ogTitle: response.data.ogTitle || '',
        ogDescription: response.data.ogDescription || '',
        ogImage: response.data.ogImage || '',
        canonical: response.data.canonical || '',
        robots: response.data.robots || '',
        yoastTitle: response.data.yoastTitle || '',
        yoastDesc: response.data.yoastDesc || '',
        rankMathTitle: response.data.rankMathTitle || '',
        rankMathDesc: response.data.rankMathDesc || ''
      }
      
      addLog('success', `متا تگ‌ها از ${metaExtractionUrl.value} استخراج شدند`)
    } else {
      addLog('error', 'خطا در استخراج متا تگ‌ها: ' + response.message)
      alert('خطا در استخراج متا تگ‌ها: ' + response.message)
    }
  } catch (error) {
    console.error('خطا در استخراج متا:', error)
    addLog('error', 'خطا در استخراج متا تگ‌ها: ' + error.message)
    alert('خطا در استخراج متا تگ‌ها: ' + error.message)
  } finally {
    isExtractingMeta.value = false
  }
}

const applyExtractedMetaToMigration = () => {
  if (!extractedMeta.value) return

  // فعال کردن گزینه‌های انتقال SEO
  migrationOptions.value.transferSEOfromYoast = true
  migrationOptions.value.transferSEOfromRankMath = true
  
  addLog('info', 'تنظیمات انتقال SEO بر اساس متا تگ‌های استخراج شده فعال شدند')
  alert('تنظیمات انتقال SEO فعال شدند!')
}

const clearExtractedMeta = () => {
  extractedMeta.value = null
  metaExtractionUrl.value = ''
}

// کلاس‌های CSS برای لاگ‌ها
const getLogClass = (type) => {
  switch (type) {
    case 'success':
      return 'bg-green-50 border border-green-200'
    case 'error':
      return 'bg-red-50 border border-red-200'
    case 'warning':
      return 'bg-yellow-50 border border-yellow-200'
    case 'info':
      return 'bg-blue-50 border border-blue-200'
    default:
      return 'bg-gray-50 border border-gray-200'
  }
}

const getLogBadgeClass = (type) => {
  switch (type) {
    case 'success':
      return 'bg-green-100 text-green-800'
    case 'error':
      return 'bg-red-100 text-red-800'
    case 'warning':
      return 'bg-yellow-100 text-yellow-800'
    case 'info':
      return 'bg-blue-100 text-blue-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getLogTypeText = (type) => {
  switch (type) {
    case 'success':
      return 'موفق'
    case 'error':
      return 'خطا'
    case 'warning':
      return 'هشدار'
    case 'info':
      return 'اطلاعات'
    default:
      return 'سایر'
  }
}
</script> 
