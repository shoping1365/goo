<template>
  <!-- قیمت‌های اصلی -->
  <div v-if="sectionSettings.mainPrices" class="bg-gradient-to-r from-green-50 to-emerald-50 border border-green-200 rounded-xl shadow-lg p-6 text-right mb-6 transition-all duration-300 hover:shadow-xl">
    <div class="flex items-center justify-between cursor-pointer" @click="toggleSection('mainPrices')">
      <div class="flex items-center gap-3">
        <div class="w-10 h-10 bg-gradient-to-r from-green-600 to-emerald-600 rounded-lg flex items-center justify-center shadow-md">
          <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
          </svg>
        </div>
        <h3 class="text-lg font-bold text-green-800">قیمت‌های اصلی</h3>
      </div>
      <span class="text-green-600 text-xl font-bold bg-white rounded-full w-8 h-8 flex items-center justify-center shadow-md">{{ sections.mainPrices ? '−' : '+' }}</span>
    </div>

    <div v-show="sections.mainPrices" class="mt-6">
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- قیمت فروش -->
        <div class="bg-white rounded-xl p-5 shadow-md border border-green-100 transition-all duration-300 hover:shadow-lg">
          <div class="flex items-center gap-2 mb-3">
            <svg class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 7h6m0 10v-3m-3 3h.01M9 17h.01M9 14h.01M12 14h.01M15 11h.01M12 11h.01M9 11h.01M7 21h10a2 2 0 002-2V5a2 2 0 00-2-2H7a2 2 0 00-2 2v14a2 2 0 002 2z"></path>
            </svg>
            <label class="text-sm font-bold text-green-800">قیمت فروش</label>
          </div>
          <div class="flex items-center gap-3 relative">
            <!-- نمایش قیمت فرمت شده -->
            <span
              v-if="store.pricingForm.price"
              class="absolute -top-7 left-0 text-lg font-extrabold text-green-700"
            >
              {{ formattedSalePrice }}
            </span>
            <span
              v-if="salePriceWarning"
              class="absolute -top-6 left-1/2 -translate-x-1/2 text-sm font-semibold text-red-600"
            >
              {{ salePriceWarning }}
            </span>
            <input v-model="store.pricingForm.price" type="number" :disabled="isSaleActive" class="flex-1 border-2 border-green-200 rounded-lg px-4 py-3 text-gray-800 focus:border-green-500 focus:ring-2 focus:ring-green-200 transition-all duration-300 disabled:bg-gray-100 disabled:text-gray-500" min="0" placeholder="0" />
            <span class="text-sm font-medium text-green-600 bg-green-100 px-3 py-2 rounded-lg">تومان</span>
          </div>
        </div>

        <!-- قیمت خط خورده -->
        <div class="bg-white rounded-xl p-5 shadow-md border border-orange-100 transition-all duration-300 hover:shadow-lg">
          <div class="flex items-center gap-2 mb-3">
            <svg class="w-5 h-5 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
            </svg>
            <label class="text-sm font-bold text-orange-800">قیمت خط خورده (قیمت قبلی)</label>
          </div>
          <div class="flex items-center gap-3 relative">
            <span
              v-if="store.pricingForm.old_price"
              class="absolute -top-6 left-0 text-lg font-extrabold text-orange-700"
            >
              {{ formattedOldPrice }}
            </span>
            <input v-model="store.pricingForm.old_price" type="number" :disabled="isSaleActive" class="flex-1 border-2 border-orange-200 rounded-lg px-4 py-3 text-gray-800 focus:border-orange-500 focus:ring-2 focus:ring-orange-200 transition-all duration-300 disabled:bg-gray-100 disabled:text-gray-500" min="0" placeholder="0" @input="updatePricingField('old_price', $event)" />
            <span class="text-sm font-medium text-orange-600 bg-orange-100 px-3 py-2 rounded-lg">تومان</span>
          </div>
        </div>

        <!-- هزینه خرید -->
        <div class="bg-white rounded-xl p-5 shadow-md border border-blue-100 transition-all duration-300 hover:shadow-lg">
          <div class="flex items-center gap-2 mb-3">
            <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4m0 0L7 13m0 0l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17M17 13v4a2 2 0 01-2 2H9a2 2 0 01-2-2v-4m8 0V9a2 2 0 00-2-2H9a2 2 0 00-2 2v4.01"></path>
            </svg>
            <label class="text-sm font-bold text-blue-800">هزینه خرید محصول (برای شما)</label>
          </div>
          <div class="flex items-center gap-3 relative">
            <span
              v-if="store.pricingForm.cost"
              class="absolute -top-6 left-0 text-lg font-extrabold text-blue-700"
            >
              {{ formattedCost }}
            </span>
            <input v-model="store.pricingForm.cost" type="number" class="flex-1 border-2 border-blue-200 rounded-lg px-4 py-3 text-gray-800 focus:border-blue-500 focus:ring-2 focus:ring-blue-200 transition-all duration-300" min="0" placeholder="0" @input="updatePricingField('cost', $event)" />
            <span class="text-sm font-medium text-blue-600 bg-blue-100 px-3 py-2 rounded-lg">تومان</span>
          </div>
        </div>

        <!-- سود خالص -->
        <div class="bg-gradient-to-r from-emerald-50 to-green-50 rounded-xl p-5 shadow-md border border-emerald-200 transition-all duration-300 hover:shadow-lg">
          <div class="flex items-center gap-2 mb-3">
            <svg class="w-5 h-5 text-emerald-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
            </svg>
            <label class="text-sm font-bold text-emerald-800">سود خالص (محاسبه شده)</label>
          </div>
          <div class="bg-white p-6 rounded-lg border border-emerald-200">
            <span class="text-lg font-bold text-emerald-600">{{ formatPrice(store.computedProfit) }} تومان ({{ store.computedProfitPercent.toFixed(1) }}%)</span>
          </div>
        </div>
      </div>

      <!-- محاسبه تخفیف -->
      <div class="mt-6 pt-6 border-t border-green-200">
        <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
          <div class="bg-gradient-to-r from-red-50 to-pink-50 rounded-xl p-5 shadow-md border border-red-200">
            <div class="flex items-center gap-2 mb-3">
              <svg class="w-5 h-5 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 8v8m-4-5v5m-4-2v2m-2 4h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
              </svg>
              <label class="text-sm font-bold text-red-800">درصد تخفیف</label>
            </div>
            <div class="bg-white p-6 rounded-lg border border-red-200">
              <input v-model.number="store.pricingForm.discount_percent" type="number" class="w-full text-xl font-bold text-red-600 bg-transparent border-0 outline-none text-center" min="0" max="100" step="0.1" placeholder="0" @input="updatePricingField('discount_percent', $event)" />
            </div>
          </div>
          <div class="bg-gradient-to-r from-red-50 to-pink-50 rounded-xl p-5 shadow-md border border-red-200">
            <div class="flex items-center gap-2 mb-3">
              <svg class="w-5 h-5 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 9V7a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2m2 4h10a2 2 0 002-2v-6a2 2 0 00-2-2H9a2 2 0 00-2 2v6a2 2 0 002 2zm7-5a2 2 0 11-4 0 2 2 0 014 0z"></path>
              </svg>
              <label class="text-sm font-bold text-red-800">مبلغ تخفیف</label>
            </div>
            <div class="bg-white p-6 rounded-lg border border-red-200">
              <input v-model="store.pricingForm.discount_amount" type="number" class="w-full text-xl font-bold text-red-600 bg-transparent border-0 outline-none text-center" min="0" placeholder="0" @input="updatePricingField('discount_amount', $event)" />
            </div>
          </div>
          <div class="bg-gradient-to-r from-blue-50 to-indigo-50 rounded-xl p-5 shadow-md border border-blue-200">
            <div class="flex items-center gap-2 mb-3">
              <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 12l3-3 3 3 4-4M8 21l4-4 4 4M3 4h18M4 4h16v12a1 1 0 01-1 1H5a1 1 0 01-1-1V4z"></path>
              </svg>
              <label class="text-sm font-bold text-blue-800">درصد سود</label>
            </div>
            <div class="bg-white p-6 rounded-lg border border-blue-200 text-center">
              <span class="text-xl font-bold text-blue-600">{{ store.computedProfitPercent.toFixed(1) }}%</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- تنظیمات قیمت و فروش -->
  <div v-if="sectionSettings.priceSettings" class="bg-gradient-to-r from-blue-50 to-indigo-50 border border-blue-200 rounded-xl shadow-lg p-6 text-right mb-6 transition-all duration-300 hover:shadow-xl">
    <div class="flex items-center justify-between cursor-pointer" @click="toggleSection('priceSettings')">
      <div class="flex items-center gap-3">
        <div class="w-10 h-10 bg-gradient-to-r from-blue-600 to-indigo-600 rounded-lg flex items-center justify-center shadow-md">
          <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"></path>
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
          </svg>
        </div>
        <h3 class="text-lg font-bold text-blue-800">تنظیمات قیمت و فروش</h3>
      </div>
      <span class="text-blue-600 text-xl font-bold bg-white rounded-full w-8 h-8 flex items-center justify-center shadow-md">{{ sections.priceSettings ? '−' : '+' }}</span>
    </div>

    <div v-show="sections.priceSettings" class="mt-6">
      <div class="grid grid-cols-1 gap-6">
        <!-- گزینه‌های اصلی -->
        <div class="bg-white rounded-xl p-5 shadow-md border border-blue-100">
          <h4 class="text-md font-bold text-blue-800 mb-4 flex items-center gap-2">
            <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
            گزینه‌های اصلی فروش
          </h4>
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <label class="flex items-center gap-3 p-3 bg-gray-50 rounded-lg hover:bg-blue-50 transition-all duration-300 cursor-pointer">
              <input 
                v-model="store.pricingForm.disableBuyButton" 
                type="checkbox" 
                class="w-5 h-5 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 focus:ring-2"
                @change="onDisableBuyButtonChange" 
              />
              <span class="text-sm text-gray-800 font-medium">دکمه خرید غیرفعال باشد - ناموجود</span>
            </label>
            <label class="flex items-center gap-3 p-3 bg-gray-50 rounded-lg hover:bg-blue-50 transition-all duration-300 cursor-pointer">
              <input type="checkbox" class="w-5 h-5 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 focus:ring-2" />
              <span class="text-sm text-gray-800 font-medium">امکان پیش خرید وجود دارد</span>
            </label>
            <label class="flex items-center gap-3 p-3 bg-gray-50 rounded-lg hover:bg-blue-50 transition-all duration-300 cursor-pointer">
              <input 
                v-model="store.pricingForm.callForPrice" 
                type="checkbox" 
                class="w-5 h-5 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 focus:ring-2"
                @change="onCallForPriceChange" 
              />
              <span class="text-sm text-gray-800 font-medium">تماس برای قیمت</span>
            </label>
            <label class="flex items-center gap-3 p-3 bg-gray-50 rounded-lg hover:bg-blue-50 transition-all duration-300 cursor-pointer">
              <input type="checkbox" class="w-5 h-5 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 focus:ring-2" />
              <span class="text-sm text-gray-800 font-medium">فقط برای اعضای VIP</span>
            </label>
          </div>
        </div>

        <!-- زمان‌بندی قیمت -->
        <div class="bg-white rounded-xl p-5 shadow-md border border-purple-100"> 
          <h4 class="text-md font-bold text-purple-800 mb-4 flex items-center justify-between">
            <div class="flex items-center gap-2">
              <svg class="w-5 h-5 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
              زمان‌بندی قیمت ویژه
            </div>
            <span :class="['px-3 py-1 text-xs font-bold rounded-full', campaignStatus.class]">
              وضعیت: {{ campaignStatus.text }}
            </span>
          </h4>
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <div class="space-y-3">
              <div class="flex items-center gap-2">
                <svg class="w-4 h-4 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                </svg>
                <label class="text-sm font-bold text-purple-800">تاریخ شروع قیمت ویژه</label>
              </div>
              <ClientOnly>
                <date-picker
                  v-model="store.pricingForm.sale_start_jalali"
                  model-type="format"
                  format="yyyy/MM/dd HH:mm"
                  :enable-time-picker="true"
                  :is-24="true"
                  placeholder="تاریخ شروع را انتخاب کنید"
                />
              </ClientOnly>
            </div>
            <div class="space-y-3">
              <div class="flex items-center gap-2">
                <svg class="w-4 h-4 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                </svg>
                <label class="text-sm font-bold text-purple-800">تاریخ پایان قیمت ویژه</label>
              </div>
              <ClientOnly>
                <date-picker
                  v-model="store.pricingForm.sale_end_jalali"
                  model-type="format"
                  format="yyyy/MM/dd HH:mm"
                  :enable-time-picker="true"
                  :is-24="true"
                  placeholder="تاریخ پایان را انتخاب کنید"
                />
              </ClientOnly>
            </div>
          </div>

          <!-- طرح‌های فروش ویژه مرحله‌ای (قیمت/تعداد) -->
          <div class="mt-6 border-t pt-4">
            <div class="flex items-center justify-between mb-3">
              <h5 class="text-sm font-bold text-purple-800">طرح‌های فروش ویژه (قیمت/تعداد)</h5>
              <button type="button" class="px-3 py-1 text-xs rounded bg-purple-600 text-white hover:bg-purple-700" @click="addSpecialOffer()">+ افزودن طرح</button>
            </div>
            <div v-if="!specialOffers || specialOffers.length === 0" class="text-xs text-gray-500">هنوز طرحی اضافه نشده است.</div>
            <div v-for="(offer, idx) in (specialOffers || [])" :key="idx" class="grid grid-cols-1 md:grid-cols-6 gap-3 items-end mb-3">
              <!-- قیمت فروش (پایه) -->
              <div class="md:col-span-2">
                <label class="text-xs font-semibold text-gray-700">قیمت فروش</label>
                <input v-model.number="offer.base_price" type="number" min="0" class="w-full border-2 border-green-200 rounded-lg px-3 py-2" placeholder="مثال: 900000" />
              </div>
              <!-- قیمت فروش ویژه -->
              <div class="md:col-span-2">
                <label class="text-xs font-semibold text-gray-700">قیمت فروش ویژه</label>
                <input v-model.number="offer.price" type="number" min="0" class="w-full border-2 border-purple-200 rounded-lg px-3 py-2" placeholder="مثال: 800000" />
              </div>
              <!-- تعداد -->
              <div class="md:col-span-1">
                <label class="text-xs font-semibold text-gray-700">تعداد با این قیمت</label>
                <input v-model.number="offer.quantity" type="number" min="1" class="w-full border-2 border-purple-200 rounded-lg px-3 py-2" placeholder="مثال: 10" />
              </div>
              <div class="md:col-span-1 flex gap-2">
                <button type="button" class="px-3 py-2 text-xs rounded bg-red-100 text-red-700 hover:bg-red-200" @click="removeSpecialOffer(idx)">حذف</button>
              </div>
            </div>
            <p class="text-[11px] text-gray-500 mt-1">منطق: ابتدا از قیمت‌های ویژه استفاده می‌شود (به ترتیب فهرست). با اتمام تعداد هر پله یا پایان بازه زمانی، قیمت به حالت بعدی/قیمت اصلی برمی‌گردد.</p>
          </div>
        </div>

        <!-- تنظیمات نمایش قیمت -->
        <div v-if="!store.pricingForm?.callForPrice && !store.pricingForm?.disableBuyButton" class="bg-white rounded-xl p-5 shadow-md border border-teal-100">
          <h4 class="text-md font-bold text-teal-800 mb-4 flex items-center gap-2">
            <svg class="w-5 h-5 text-teal-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
            </svg>
            تنظیمات نمایش قیمت
          </h4>
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <label class="flex items-center gap-3 p-3 bg-gray-50 rounded-lg transition-all duration-300 cursor-pointer">
              <input type="checkbox" class="w-5 h-5 text-teal-600 bg-gray-100 border-gray-300 rounded focus:ring-teal-500 focus:ring-2" checked />
              <span class="text-sm text-gray-800 font-medium">نمایش قیمت به مشتریان</span>
            </label>
            <label class="flex items-center gap-3 p-3 bg-gray-50 rounded-lg transition-all duration-300 cursor-pointer">
              <input type="checkbox" class="w-5 h-5 text-teal-600 bg-gray-100 border-gray-300 rounded focus:ring-teal-500 focus:ring-2" />
              <span class="text-sm text-gray-800 font-medium">مخفی کردن قیمت تا ورود</span>
            </label>
            <label class="flex items-center gap-3 p-3 bg-gray-50 rounded-lg transition-all duration-300 cursor-pointer">
              <input type="checkbox" class="w-5 h-5 text-teal-600 bg-gray-100 border-gray-300 rounded focus:ring-teal-500 focus:ring-2" />
              <span class="text-sm text-gray-800 font-medium">نمایش قیمت بر حسب واحد</span>
            </label>
            <label class="flex items-center gap-3 p-3 bg-gray-50 rounded-lg transition-all duration-300 cursor-pointer">
              <input type="checkbox" class="w-5 h-5 text-teal-600 bg-gray-100 border-gray-300 rounded focus:ring-teal-500 focus:ring-2" />
              <span class="text-sm text-gray-800 font-medium">نمایش درصد تخفیف</span>
            </label>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- مالیات و هزینه‌های اضافی -->
  <div v-if="sectionSettings.taxSettings" class="bg-gradient-to-r from-orange-50 to-amber-50 border border-orange-200 rounded-xl shadow-lg p-6 text-right mb-6 transition-all duration-300 hover:shadow-xl">
    <div class="flex items-center justify-between cursor-pointer" @click="toggleSection('taxSettings')">
      <div class="flex items-center gap-3">
        <div class="w-10 h-10 bg-gradient-to-r from-orange-600 to-amber-600 rounded-lg flex items-center justify-center shadow-md">
          <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 7h6m0 10v-3m-3 3h.01M9 17h.01M9 14h.01M12 14h.01M15 11h.01M12 11h.01M9 11h.01M7 21h10a2 2 0 002-2V5a2 2 0 00-2-2H7a2 2 0 00-2 2v14a2 2 0 002 2z"></path>
          </svg>
        </div>
        <h3 class="text-lg font-bold text-orange-800">مالیات و هزینه‌های اضافی</h3>
      </div>
      <span class="text-orange-600 text-xl font-bold bg-white rounded-full w-8 h-8 flex items-center justify-center shadow-md">{{ sections.taxSettings ? '−' : '+' }}</span>
    </div>

    <div v-show="sections.taxSettings" class="mt-4">
      <div class="grid grid-cols-1 gap-6">
        <!-- تنظیمات مالیات -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div class="flex flex-col gap-2">
            <label class="text-xs text-gray-700 font-semibold">دسته مالیات</label>
            <select class="w-full bg-white border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-orange-500 focus:border-orange-500 p-2.5">
              <option value="">بدون مالیات</option>
              <option value="standard">مالیات استاندارد (9%)</option>
              <option value="reduced">مالیات کاهش یافته (5%)</option>
              <option value="food">مواد غذایی (5%)</option>
              <option value="services">خدمات (9%)</option>
            </select>
          </div>

          <div class="flex flex-col gap-2">
            <label class="text-xs text-gray-700 font-semibold">درصد مالیات سفارشی</label>
            <div class="flex items-center gap-2">
              <input type="number" class="w-24 bg-white border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-orange-500 focus:border-orange-500 p-2.5" min="0" max="100" step="0.1" placeholder="0" />
              <span class="text-xs text-gray-500">درصد</span>
            </div>
          </div>
        </div>

        <!-- چک‌باکس‌های مالیات -->
        <div class="border-t pt-4 border-blue-200">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-2">
            <label class="flex items-center gap-2">
              <input type="checkbox" class="checkbox" />
              <span class="text-xs text-gray-700">شامل مالیات نمی‌باشد</span>
            </label>
            <label class="flex items-center gap-2">
              <input type="checkbox" class="checkbox" />
              <span class="text-xs text-gray-700">مالیات در قیمت لحاظ شده</span>
            </label>
            <label class="flex items-center gap-2">
              <input type="checkbox" class="checkbox" />
              <span class="text-xs text-gray-700">معاف از مالیات</span>
            </label>
            <label class="flex items-center gap-2">
              <input type="checkbox" class="checkbox" />
              <span class="text-xs text-gray-700">محاسبه مالیات در checkout</span>
            </label>
          </div>
        </div>

        <!-- هزینه‌های اضافی -->
        <div class="border-t pt-4 border-blue-200">
          <h4 class="text-xs font-semibold text-gray-700 mb-3">هزینه‌های اضافی</h4>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div class="flex flex-col gap-2">
              <label class="text-xs text-gray-700 font-semibold">هزینه بسته‌بندی</label>
              <div class="flex items-center gap-2">
                <input type="number" class="w-full bg-white border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-orange-500 focus:border-orange-500 p-2.5" min="0" placeholder="0" />
                <span class="text-xs text-gray-500">تومان</span>
              </div>
            </div>
            <div class="flex flex-col gap-2">
              <label class="text-xs text-gray-700 font-semibold">هزینه خدمات اضافی</label>
              <div class="flex items-center gap-2">
                <input type="number" class="w-full bg-white border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-orange-500 focus:border-orange-500 p-2.5" min="0" placeholder="0" />
                <span class="text-xs text-gray-500">تومان</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- قیمت‌گذاری پلکانی و عمده -->
  <div v-if="sectionSettings.tierPricing" class="bg-gradient-to-r from-purple-50 to-pink-50 border border-purple-200 rounded-xl shadow-lg p-6 text-right mb-6 transition-all duration-300 hover:shadow-xl">
    <div class="flex items-center justify-between cursor-pointer" @click="toggleSection('tierPricing')">
      <div class="flex items-center gap-3">
        <div class="w-10 h-10 bg-gradient-to-r from-purple-600 to-pink-600 rounded-lg flex items-center justify-center shadow-md">
          <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 8v8m-4-5v5m-4-2v2m-2 4h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
          </svg>
        </div>
        <h3 class="text-lg font-bold text-purple-800">قیمت‌گذاری پلکانی و خرید عمده</h3>
      </div>
      <span class="text-purple-600 text-xl font-bold bg-white rounded-full w-8 h-8 flex items-center justify-center shadow-md">{{ sections.tierPricing ? '−' : '+' }}</span>
    </div>

    <div v-show="sections.tierPricing" class="mt-4">
      <!-- فرم افزودن قیمت پلکانی -->
      <div class="border rounded p-3 bg-gray-50 mb-4 border-blue-200">
        <div class="grid grid-cols-1 md:grid-cols-5 gap-2 items-end">
          <div class="flex flex-col gap-1">
            <label class="text-xs text-gray-700 font-semibold">نقش مشتری</label>
            <select class="w-full bg-white border border-gray-300 text-gray-900 text-xs rounded-lg focus:ring-purple-500 focus:border-purple-500 p-2.5">
              <option>همه مشتریان</option>
              <option>مشتریان VIP</option>
              <option>مشتریان عادی</option>
              <option>عمده فروشان</option>
            </select>
          </div>
          <div class="flex flex-col gap-1">
            <label class="text-xs text-gray-700 font-semibold">حداقل تعداد</label>
            <input type="number" class="w-full bg-white border border-gray-300 text-gray-900 text-xs rounded-lg focus:ring-purple-500 focus:border-purple-500 p-2.5" min="1" placeholder="1" />
          </div>
          <div class="flex flex-col gap-1">
            <label class="text-xs text-gray-700 font-semibold">قیمت واحد</label>
            <input type="number" class="w-full bg-white border border-gray-300 text-gray-900 text-xs rounded-lg focus:ring-purple-500 focus:border-purple-500 p-2.5" min="0" placeholder="0" />
          </div>
          <div class="flex flex-col gap-1">
            <label class="text-xs text-gray-700 font-semibold">تاریخ انقضا</label>
            <input type="date" class="w-full bg-white border border-gray-300 text-gray-900 text-xs rounded-lg focus:ring-purple-500 focus:border-purple-500 p-2.5" />
          </div>
          <div>
            <button class="bg-blue-600 text-white rounded px-3 py-1 text-xs hover:bg-blue-700 transition">افزودن</button>
          </div>
        </div>
      </div>

      <!-- جدول قیمت‌های پلکانی -->
      <div class="overflow-x-auto">
        <table class="min-w-full text-xs text-right rtl border-collapse">
          <thead>
          <tr class="bg-gray-100 border-b border-blue-200">
            <th class="px-3 py-2 text-gray-600 font-normal text-right">نقش مشتری</th>
            <th class="px-3 py-2 text-gray-600 font-normal text-right">حداقل تعداد</th>
            <th class="px-3 py-2 text-gray-600 font-normal text-right">قیمت واحد</th>
            <th class="px-3 py-2 text-gray-600 font-normal text-right">تخفیف</th>
            <th class="px-3 py-2 text-gray-600 font-normal text-right">تاریخ شروع</th>
            <th class="px-3 py-2 text-gray-600 font-normal text-right">تاریخ پایان</th>
            <th class="px-3 py-2 text-gray-600 font-normal text-right w-20">ویرایش</th>
            <th class="px-3 py-2 text-gray-600 font-normal text-right w-16">حذف</th>
          </tr>
          </thead>
          <tbody>
          <tr class="border-b hover:bg-gray-50 border-blue-200">
            <td class="px-3 py-2 text-right">عمده فروشان</td>
            <td class="px-3 py-2 text-right">10</td>
            <td class="px-3 py-2 text-right">45,000</td>
            <td class="px-3 py-2 text-right">10%</td>
            <td class="px-3 py-2 text-right">1403/08/01</td>
            <td class="px-3 py-2 text-right">1403/12/29</td>
            <td class="px-3 py-2 text-center">
              <button class="text-blue-500 hover:text-blue-700">ویرایش</button>
            </td>
            <td class="px-3 py-2 text-center">
              <button class="text-red-500 hover:text-red-700">حذف</button>
            </td>
          </tr>
          <tr class="border-b hover:bg-gray-50 border-blue-200">
            <td class="px-3 py-2 text-right">مشتریان VIP</td>
            <td class="px-3 py-2 text-right">5</td>
            <td class="px-3 py-2 text-right">47,500</td>
            <td class="px-3 py-2 text-right">5%</td>
            <td class="px-3 py-2 text-right">1403/08/01</td>
            <td class="px-3 py-2 text-right">1403/12/29</td>
            <td class="px-3 py-2 text-center">
              <button class="text-blue-500 hover:text-blue-700">ویرایش</button>
            </td>
            <td class="px-3 py-2 text-center">
              <button class="text-red-500 hover:text-red-700">حذف</button>
            </td>
          </tr>
          </tbody>
        </table>
      </div>

      <!-- تنظیمات قیمت‌گذاری پلکانی -->
      <div class="mt-4 pt-4 border-t border-blue-200">
        <h4 class="text-xs font-semibold text-gray-700 mb-3">تنظیمات قیمت‌گذاری پلکانی</h4>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-2">
          <label class="flex items-center gap-2">
            <input type="checkbox" class="checkbox" checked />
            <span class="text-xs text-gray-700">فعال‌سازی قیمت‌گذاری پلکانی</span>
          </label>
          <label class="flex items-center gap-2">
            <input type="checkbox" class="checkbox" />
            <span class="text-xs text-gray-700">نمایش جدول قیمت به مشتریان</span>
          </label>
          <label class="flex items-center gap-2">
            <input type="checkbox" class="checkbox" />
            <span class="text-xs text-gray-700">اعمال خودکار بهترین قیمت</span>
          </label>
          <label class="flex items-center gap-2">
            <input type="checkbox" class="checkbox" />
            <span class="text-xs text-gray-700">ترکیب با کدهای تخفیف</span>
          </label>
        </div>
      </div>
    </div>
  </div>

  <!-- کدهای تخفیف و کوپن -->
  <div v-if="sectionSettings.discountCodes" class="bg-gradient-to-r from-red-50 to-rose-50 border border-red-200 rounded-xl shadow-lg p-6 text-right transition-all duration-300 hover:shadow-xl mb-6">
    <div class="flex items-center justify-between cursor-pointer" @click="toggleSection('discountCodes')">
      <div class="flex items-center gap-3">
        <div class="w-10 h-10 bg-gradient-to-r from-red-600 to-rose-600 rounded-lg flex items-center justify-center shadow-md">
          <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"></path>
          </svg>
        </div>
        <h3 class="text-lg font-bold text-red-800">کدهای تخفیف و کوپن‌های مخصوص</h3>
      </div>
      <span class="text-red-600 text-xl font-bold bg-white rounded-full w-8 h-8 flex items-center justify-center shadow-md">{{ sections.discountCodes ? '−' : '+' }}</span>
    </div>

    <div v-show="sections.discountCodes" class="mt-4">
      <!-- ایجاد کد تخفیف -->
      <div class="border rounded p-3 bg-gray-50 mb-4 border-blue-200">
        <div class="grid grid-cols-1 md:grid-cols-4 gap-2 items-end">
          <div class="flex flex-col gap-1">
            <label class="text-xs text-gray-700 font-semibold">کد تخفیف</label>
            <input type="text" class="w-full bg-white border border-gray-300 text-gray-900 text-xs rounded-lg focus:ring-red-500 focus:border-red-500 p-2.5" placeholder="SAVE20" />
          </div>
          <div class="flex flex-col gap-1">
            <label class="text-xs text-gray-700 font-semibold">نوع تخفیف</label>
            <select class="w-full bg-white border border-gray-300 text-gray-900 text-xs rounded-lg focus:ring-red-500 focus:border-red-500 p-2.5">
              <option>درصدی</option>
              <option>مبلغی</option>
            </select>
          </div>
          <div class="flex flex-col gap-1">
            <label class="text-xs text-gray-700 font-semibold">مقدار تخفیف</label>
            <input type="number" class="w-full bg-white border border-gray-300 text-gray-900 text-xs rounded-lg focus:ring-red-500 focus:border-red-500 p-2.5" min="0" placeholder="20" />
          </div>
          <div>
            <button class="bg-green-600 text-white rounded px-3 py-1 text-xs hover:bg-green-700 transition">ایجاد کد</button>
          </div>
        </div>
      </div>

      <!-- لیست کدهای تخفیف -->
      <div class="overflow-x-auto">
        <table class="min-w-full text-xs text-right rtl border-collapse">
          <thead>
          <tr class="bg-gray-100 border-b border-blue-200">
            <th class="px-3 py-2 text-gray-600 font-normal text-right">کد تخفیف</th>
            <th class="px-3 py-2 text-gray-600 font-normal text-right">نوع</th>
            <th class="px-3 py-2 text-gray-600 font-normal text-right">مقدار</th>
            <th class="px-3 py-2 text-gray-600 font-normal text-right">استفاده شده</th>
            <th class="px-3 py-2 text-gray-600 font-normal text-right">وضعیت</th>
            <th class="px-3 py-2 text-gray-600 font-normal text-right w-20">ویرایش</th>
            <th class="px-3 py-2 text-gray-600 font-normal text-right w-16">حذف</th>
          </tr>
          </thead>
          <tbody>
          <tr class="border-b hover:bg-gray-50 border-blue-200">
            <td class="px-3 py-2 text-right font-mono">SAVE20</td>
            <td class="px-3 py-2 text-right">درصدی</td>
            <td class="px-3 py-2 text-right">20%</td>
            <td class="px-3 py-2 text-right">3/100</td>
            <td class="px-3 py-2 text-right">
              <span class="bg-green-100 text-green-800 px-2 py-1 rounded text-xs">فعال</span>
            </td>
            <td class="px-3 py-2 text-center">
              <button class="text-blue-500 hover:text-blue-700">ویرایش</button>
            </td>
            <td class="px-3 py-2 text-center">
              <button class="text-red-500 hover:text-red-700">حذف</button>
            </td>
          </tr>
          </tbody>
        </table>
      </div>

      <!-- خلاصه قیمت نهایی -->
      <div class="mt-4 pt-4 border-t border-blue-200">
        <div class="bg-gray-50 p-3 rounded border-blue-200 border">
          <h4 class="text-xs font-semibold text-gray-700 mb-2">خلاصه قیمت نهایی محصول</h4>
          <div class="grid grid-cols-2 md:grid-cols-4 gap-2 text-xs">
            <div>
              <span class="text-gray-600">قیمت پایه:</span>
              <span class="font-medium text-blue-600"> {{ formatPrice(store.pricingForm.price) }} تومان</span>
            </div>
            <div>
              <span class="text-gray-600">مالیات:</span>
              <span class="font-medium text-blue-600"> 0 تومان</span>
            </div>
            <div>
              <span class="text-gray-600">هزینه اضافی:</span>
              <span class="font-medium text-blue-600"> 0 تومان</span>
            </div>
            <div>
              <span class="text-gray-600">قیمت نهایی:</span>
              <span class="font-bold text-green-600"> {{ formatPrice(store.pricingForm.price - (store.pricingForm.discount_amount || 0)) }} تومان</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- تحلیل قیمت رقبا -->
  <div v-if="sectionSettings.priceAnalysis" class="bg-gradient-to-r from-yellow-50 to-amber-50 border border-amber-200 rounded-xl shadow-lg p-6 text-right transition-all duration-300 hover:shadow-xl mt-6">
    <div class="flex items-center justify-between cursor-pointer" @click="toggleSection('priceAnalysis')">
      <div class="flex items-center gap-3">
        <div class="w-10 h-10 bg-gradient-to-r from-amber-600 to-yellow-600 rounded-lg flex items-center justify-center shadow-md">
          <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
          </svg>
        </div>
        <h3 class="text-lg font-bold text-amber-800">تحلیل قیمت رقبا</h3>
      </div>
      <span class="text-amber-600 text-xl font-bold bg-white rounded-full w-8 h-8 flex items-center justify-center shadow-md">{{ sections.priceAnalysis ? '−' : '+' }}</span>
    </div>

    <div v-show="sections.priceAnalysis" class="mt-6">
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <div class="bg-white rounded-xl p-6 border border-amber-200">
          <label class="text-xs font-semibold text-gray-700 mb-2 block">نام یا کلمه کلیدی محصول</label>
          <input v-model.trim="keyword" type="text" class="w-full border-2 border-amber-200 rounded-lg px-3 py-2 focus:border-amber-500 focus:ring-2 focus:ring-amber-200" placeholder="مثال: گوشی سامسونگ A55 128" />
          <p class="text-[11px] text-gray-500 mt-1">برای دقت بهتر، برند/مدل/ظرفیت را هم وارد کنید.</p>
        </div>
        <div class="bg-white rounded-xl p-6 border border-amber-200">
          <label class="text-xs font-semibold text-gray-700 mb-2 block">تعداد نتایج</label>
          <input v-model.number="limit" type="number" min="3" max="15" class="w-full border-2 border-amber-200 rounded-lg px-3 py-2 focus:border-amber-500 focus:ring-2 focus:ring-amber-200" />
        </div>
        <div class="bg-white rounded-xl p-6 border border-amber-200 flex items-end">
          <button type="button" class="w-full bg-amber-600 text-white rounded-lg px-4 py-2 hover:bg-amber-700 disabled:opacity-50" :disabled="loading" @click="runPriceAnalysis">
            <span v-if="loading">در حال تحلیل...</span>
            <span v-else>تحلیل قیمت رقبا</span>
          </button>
        </div>
      </div>

      <div v-if="errorMsg" class="mt-4 text-sm text-red-600">{{ errorMsg }}</div>

      <div v-if="results.length" class="mt-6 overflow-x-auto">
        <table class="min-w-full text-xs text-right rtl border-collapse">
          <thead>
            <tr class="bg-amber-100 border-b border-amber-200">
              <th class="px-3 py-2">سایت</th>
              <th class="px-3 py-2">عنوان</th>
              <th class="px-3 py-2">قیمت</th>
              <th class="px-3 py-2">واحد</th>
              <th class="px-3 py-2">اعتماد</th>
              <th class="px-3 py-2">لینک</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="row in results" :key="row.url" class="border-b border-amber-100 hover:bg-amber-50">
              <td class="px-3 py-2 whitespace-nowrap">{{ row.domain }}</td>
              <td class="px-3 py-2 max-w-[360px] truncate" :title="row.title">{{ row.title }}</td>
              <td class="px-3 py-2 font-semibold">{{ row.priceFormatted }}</td>
              <td class="px-3 py-2">{{ row.currency || '-' }}</td>
              <td class="px-3 py-2">{{ (row.confidence ?? 0).toFixed(2) }}</td>
              <td class="px-3 py-2"><a :href="row.url" target="_blank" class="text-blue-600 hover:underline">مشاهده</a></td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
</script>

<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import { storeToRefs } from 'pinia'
import { useProductCreateStore } from '~/stores/productCreate'

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// Props for edit mode
defineProps({
  isEditMode: {
    type: Boolean,
    default: true // Always in edit mode here
  },
  productId: {
    type: [String, Number],
    default: null
  },
  sectionSettings: {
    type: Object,
    default: () => ({
      mainPrices: true,
      priceSettings: true,
      taxSettings: true,
      tierPricing: true,
      discountCodes: true
    })
  }
})

const store = useProductCreateStore()
const { specialOffers } = storeToRefs(store)
// const notifier = useNotifier()

// const { sections: sectionSettings } = store

// Formatting numbers for display
const formatPrice = (price) => new Intl.NumberFormat('fa-IR').format(price)
const formattedSalePrice = computed(() => formatPrice(store.pricingForm.price))
const formattedOldPrice = computed(() => formatPrice(store.pricingForm.old_price))
const formattedCost = computed(() => formatPrice(store.pricingForm.cost))

// Computed property to check if a sale is active (disables main price fields)
const isSaleActive = computed(() => {
  return (specialOffers.value && specialOffers.value.length > 0) || (store.pricingForm.sale_price && store.pricingForm.sale_price > 0)
})

const salePriceWarning = computed(() => {
  if (store.pricingForm.price <= 0 && isSaleActive.value) {
    return 'قیمت اصلی برای محاسبه تخفیف الزامی است'
  }
  return ''
})

function updatePricingField(field: string, event: Event): void {
  const target = event.target as HTMLInputElement
  const value = target.value
  const numericValue = Number(value)

  if (field === 'price' || field === 'cost' || field === 'old_price' || field === 'discount_amount') {
    (store.pricingForm)[field] = numericValue
    if (store.pricingForm.price > 0 && store.pricingForm.old_price > 0) {
      store.pricingForm.discount_percent = ((store.pricingForm.old_price - store.pricingForm.price) / store.pricingForm.old_price) * 100
      store.pricingForm.discount_amount = store.pricingForm.old_price - store.pricingForm.price
    }
  }

  if (field === 'discount_percent') {
    store.pricingForm.discount_percent = numericValue
    if (store.pricingForm.price > 0) {
      store.pricingForm.discount_amount = (store.pricingForm.price * numericValue) / 100
    }
  }
}

function addSpecialOffer() {
  const item = { base_price: store.pricingForm.price, price: 0, quantity: 1 }
  // انتساب آرایه‌ای برای تضمین تریگر رندر در تمام سناریوها
  specialOffers.value = [...(specialOffers.value || []), item]
}

function removeSpecialOffer(index: number): void {
  if (specialOffers.value) {
    specialOffers.value.splice(index, 1)
  }
}

function onDisableBuyButtonChange(event: Event): void {
  const target = event.target as HTMLInputElement
  const isChecked = target.checked
  if (isChecked) {
    store.pricingForm.callForPrice = false
  }
}

function onCallForPriceChange(event: Event): void {
  const target = event.target as HTMLInputElement
  const isChecked = target.checked
  if (isChecked) {
    store.pricingForm.disableBuyButton = false
  }
}

// function prefillIfEmpty() {
//   if (!store.pricingForm.sale_start_jalali) {
//     const now = new Date()
//     const year = now.getFullYear()
//     const month = (now.getMonth() + 1).toString().padStart(2, '0')
//     const day = now.getDate().toString().padStart(2, '0')
//     const hours = now.getHours().toString().padStart(2, '0')
//     const minutes = now.getMinutes().toString().padStart(2, '0')
//     store.pricingForm.sale_start_jalali = `${year}/${month}/${day} ${hours}:${minutes}`
//   }
// }

const { toggleSection, sections } = store

// ---------------------------
// Price Analysis (Competitors)
// ---------------------------
const keyword = ref('')
const limit = ref(10)
const loading = ref(false)
const errorMsg = ref('')
const results = ref([])

function formatToman(val) {
  const num = Number(val || 0)
  return new Intl.NumberFormat('fa-IR').format(num)
}

async function runPriceAnalysis() {
  try {
    errorMsg.value = ''
    results.value = []
    if (!keyword.value || !keyword.value.trim()) {
      errorMsg.value = 'کلیدواژه محصول را وارد کنید'
      return
    }
    loading.value = true
    const res = await $fetch('/api/admin/price-analysis', {
      method: 'POST',
      body: { q: keyword.value.trim(), limit: Number(limit.value) || 10 }
    })
    interface PriceAnalysisResponse {
      data?: unknown[]
      results?: unknown[]
      [key: string]: unknown
    }
    const response = res as PriceAnalysisResponse
    const rows = Array.isArray(response?.data) ? response.data : (Array.isArray(response?.results) ? response.results : [])
    results.value = rows.map((r: { price?: number; [key: string]: unknown }) => ({
      ...r,
      priceFormatted: r?.price ? `${formatToman(r.price)} تومان` : '-'
    }))
  } catch (err) {
    errorMsg.value = err?.data?.message || err?.message || 'خطا در تحلیل قیمت'
  } finally {
    loading.value = false
  }
}

// Prefill keyword from current product titles (fa/en)
function buildKeywordFromTitles() {
  const fa = (store.productForm?.name || '').toString().trim()
  const en = (store.productForm?.englishName || '').toString().trim()
  const parts = []
  if (fa) parts.push(fa)
  if (en && en.toLowerCase() !== fa.toLowerCase()) parts.push(en)
  return parts.join(' ')
}

watch(() => [store.productForm.name, store.productForm.englishName], () => {
  if (!keyword.value || keyword.value.trim().length < 2) {
    keyword.value = buildKeywordFromTitles()
  }
}, { immediate: true })

const campaignStatus = computed(() => {
  const now = new Date()
  const start = store.pricingForm.sale_start_at ? new Date(store.pricingForm.sale_start_at) : null
  const end = store.pricingForm.sale_end_at ? new Date(store.pricingForm.sale_end_at) : null
  const hasOffers = (specialOffers.value || []).length > 0 && specialOffers.value.some(o => o.price > 0)

  if (!hasOffers && !store.pricingForm.sale_price) {
      return { text: 'بدون قیمت ویژه', class: 'bg-gray-100 text-gray-800' }
  }
  
  if (start && end && end < start) {
      return { text: 'بازه نامعتبر', class: 'bg-yellow-100 text-yellow-800' }
  }

  if (!start && !end) {
    return { text: 'فعال (همیشه)', class: 'bg-green-100 text-green-800' }
  }

  if (start && now < start) {
    return { text: 'زمان‌بندی شده', class: 'bg-blue-100 text-blue-800' }
  }

  if (end && now > end) {
    return { text: 'منقضی شده', class: 'bg-red-100 text-red-800' }
  }

  if ((start && now >= start && (!end || now <= end)) || (!start && end && now <= end)) {
    return { text: 'فعال', class: 'bg-green-100 text-green-800' }
  }

  return { text: 'نامشخص', class: 'bg-gray-100 text-gray-800' }
})
</script>
