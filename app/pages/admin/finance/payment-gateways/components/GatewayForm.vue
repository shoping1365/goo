<template>
  <div class="max-w-6xl mx-auto">
    <div class="bg-white rounded-2xl shadow-xl border border-gray-100 overflow-hidden" style="box-shadow: 0 10px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04);">
      <!-- هدر زیبا با پالت رنگ آمریکایی -->
      <div class="bg-gradient-to-br from-indigo-500 via-purple-500 to-pink-500 p-8 relative overflow-hidden">
        <div class="absolute inset-0 bg-black opacity-10"></div>
        <div class="relative flex items-center justify-between">
          <div class="flex items-center space-x-6 space-x-reverse">
            <div class="w-16 h-16 bg-white/20 backdrop-blur-sm rounded-2xl flex items-center justify-center border border-white/30 shadow-lg">
              <span class="i-heroicons-credit-card text-white text-2xl"></span>
            </div>
            <div>
              <h3 class="text-2xl font-bold text-white mb-2">{{ props.mode === 'edit' ? `ویرایش درگاه پرداخت ${props.gateway?.name || ''}` : 'افزودن درگاه پرداخت جدید' }}</h3>
              <p class="text-indigo-100 text-lg">تنظیم اطلاعات و پارامترهای درگاه پرداخت</p>
            </div>
          </div>
          <TemplateButton 
            bg-gradient="bg-gradient-to-r from-white/20 to-white/10" 
            text-color="text-white"
            border-color="border border-white/30"
            hover-class="hover:from-white/30 hover:to-white/20 hover:border-white/50"
            focus-class="focus:ring-2 focus:ring-white/50 focus:ring-offset-2"
            size="large"
            custom-class="backdrop-blur-sm font-bold text-lg"
            @click="$emit('cancel')"
          >
            بازگشت
          </TemplateButton>
        </div>
        <!-- دکوراسیون پس‌زمینه -->
        <div class="absolute top-0 right-0 w-32 h-32 bg-white/10 rounded-full -translate-y-16 translate-x-16"></div>
        <div class="absolute bottom-0 left-0 w-24 h-24 bg-white/10 rounded-full translate-y-12 -translate-x-12"></div>
      </div>
    
      <form class="p-8 bg-gradient-to-br from-gray-50 to-white" @submit.prevent="handleSubmit">
        <!-- نمایش خطاهای اعتبارسنجی -->
        <div
v-if="showErrors && validationErrors.length > 0" 
          class="mb-8 p-6 bg-red-50 border-2 border-red-200 rounded-2xl shadow-lg">
          <div class="flex items-center mb-4">
            <div class="w-10 h-10 bg-red-500 rounded-xl flex items-center justify-center mr-4">
              <span class="i-heroicons-exclamation-triangle text-white text-lg"></span>
            </div>
            <h4 class="text-lg font-bold text-red-800">خطاهای اعتبارسنجی</h4>
          </div>
          <ul class="space-y-2">
            <li
v-for="(error, index) in validationErrors" :key="index" 
              class="flex items-center text-red-700 text-sm">
              <span class="i-heroicons-x-circle text-red-500 mr-2"></span>
              {{ error }}
            </li>
          </ul>
          <TemplateButton 
            bg-gradient="bg-gradient-to-r from-red-100 to-red-200" 
            text-color="text-red-700"
            border-color="border border-red-300"
            hover-class="hover:from-red-200 hover:to-red-300"
            focus-class="focus:ring-2 focus:ring-red-500 focus:ring-offset-2"
            size="small"
            custom-class="font-bold"
            @click="showErrors = false"
          >
            بستن
          </TemplateButton>
        </div>
        
        <!-- تب‌های فرم با طراحی مدرن -->
        <div class="flex bg-white rounded-2xl p-2 shadow-lg border border-gray-100 mb-8">
          <button
v-for="tab in formTabs" :key="tab.value" type="button" :class="[
              'px-8 py-4 font-semibold text-sm focus:outline-none transition-all rounded-xl relative flex-1 flex items-center justify-center',
              activeTab === tab.value 
                ? 'bg-gradient-to-r from-blue-500 to-indigo-600 text-white shadow-lg transform scale-105' 
                : 'text-gray-600 hover:text-blue-600 hover:bg-blue-50'
            ]"
            @click="activeTab = tab.value">
            <span v-if="activeTab === tab.value" class="absolute inset-0 bg-gradient-to-r from-blue-500 to-indigo-600 rounded-xl opacity-90"></span>
            <span class="relative z-10 flex items-center">
              {{ tab.label }}
              <!-- نشانگر تکمیل -->
              <span v-if="isTabComplete(tab.value)" class="mr-2 text-green-300">
                <span class="i-heroicons-check-circle"></span>
              </span>
            </span>
            <div v-if="activeTab === tab.value" class="absolute -bottom-1 left-1/2 transform -translate-x-1/2 w-8 h-1 bg-white rounded-full"></div>
          </button>
        </div>

        <!-- محتوای تب‌ها -->
        <div class="space-y-8">
          <!-- اطلاعات پایه -->
          <div v-if="activeTab === 'basic'" class="space-y-8">
            <div class="bg-white rounded-2xl p-8 shadow-lg border border-gray-100">
              <h4 class="text-xl font-bold text-gray-900 mb-6">
                اطلاعات پایه درگاه
              </h4>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-8">

                <div class="space-y-3">
                  <label class="block text-sm font-bold text-gray-700">نوع درگاه *</label>
                  <select
v-model="form.type" required class="w-full px-5 py-4 border-2 border-gray-200 rounded-xl text-sm focus:ring-4 focus:ring-blue-100 focus:border-blue-500 transition-all bg-white hover:bg-gray-50 shadow-sm" 
                    @change="onGatewayTypeChange">
                    <option value="">انتخاب کنید</option>
                    <option value="zarinpal" :disabled="isGatewayExists('zarinpal')" :class="{ 'text-gray-400': isGatewayExists('zarinpal') }">
                      زرین‌پال {{ isGatewayExists('zarinpal') ? '(قبلاً ثبت شده)' : '' }}
                    </option>
                    <option value="mellat" :disabled="isGatewayExists('mellat')" :class="{ 'text-gray-400': isGatewayExists('mellat') }">
                      ملت {{ isGatewayExists('mellat') ? '(قبلاً ثبت شده)' : '' }}
                    </option>
                    <option value="melli" :disabled="isGatewayExists('melli')" :class="{ 'text-gray-400': isGatewayExists('melli') }">
                      ملی {{ isGatewayExists('melli') ? '(قبلاً ثبت شده)' : '' }}
                    </option>
                    <option value="saman" :disabled="isGatewayExists('saman')" :class="{ 'text-gray-400': isGatewayExists('saman') }">
                      سامان {{ isGatewayExists('saman') ? '(قبلاً ثبت شده)' : '' }}
                    </option>
                    <option value="parsian" :disabled="isGatewayExists('parsian')" :class="{ 'text-gray-400': isGatewayExists('parsian') }">
                      پارسیان {{ isGatewayExists('parsian') ? '(قبلاً ثبت شده)' : '' }}
                    </option>
                    <option value="paypal" :disabled="isGatewayExists('paypal')" :class="{ 'text-gray-400': isGatewayExists('paypal') }">
                      پی‌پال {{ isGatewayExists('paypal') ? '(قبلاً ثبت شده)' : '' }}
                    </option>
                    <option value="stripe" :disabled="isGatewayExists('stripe')" :class="{ 'text-gray-400': isGatewayExists('stripe') }">
                      استرایپ {{ isGatewayExists('stripe') ? '(قبلاً ثبت شده)' : '' }}
                    </option>
                  </select>
                </div>

                <div class="space-y-3">
                  <label class="block text-sm font-bold text-gray-700">کارمزد (%)</label>
                  <input
v-model.number="form.fee" type="number" min="0" max="100" step="0.1" 
                    class="w-full px-5 py-4 border-2 border-gray-200 rounded-xl text-sm focus:ring-4 focus:ring-blue-100 focus:border-blue-500 transition-all bg-white hover:bg-gray-50 shadow-sm"
                    placeholder="0 = بدون کارمزد">
                  <p class="text-xs text-gray-500">صفر یعنی بدون کارمزد</p>
                </div>
                <div class="space-y-3">
                  <label class="block text-sm font-bold text-gray-700">حداقل مبلغ (تومان)</label>
                  <input
v-model.number="form.minAmount" type="number" min="0" 
                    class="w-full px-5 py-4 border-2 border-gray-200 rounded-xl text-sm focus:ring-4 focus:ring-blue-100 focus:border-blue-500 transition-all bg-white hover:bg-gray-50 shadow-sm"
                    placeholder="0 = بدون محدودیت">
                  <p class="text-xs text-gray-500">صفر یعنی بدون محدودیت حداقل</p>
                </div>
                <div class="space-y-3">
                  <label class="block text-sm font-bold text-gray-700">حداکثر مبلغ (تومان)</label>
                  <input
v-model.number="form.maxAmount" type="number" min="0" 
                    class="w-full px-5 py-4 border-2 border-gray-200 rounded-xl text-sm focus:ring-4 focus:ring-blue-100 focus:border-blue-500 transition-all bg-white hover:bg-gray-50 shadow-sm"
                    placeholder="0 = بدون محدودیت">
                  <p class="text-xs text-gray-500">صفر یعنی بدون محدودیت حداکثر</p>
                </div>

                <div class="space-y-3">
                  <label class="block text-sm font-bold text-gray-700">آیکون درگاه</label>
                  <div v-if="form.type" class="w-20 h-20 rounded-xl border-2 border-gray-200 bg-white flex items-center justify-center shadow-sm">
                    <!-- استفاده از آیکون SVG به جای تصویر برای سرعت بیشتر -->
                    <div class="w-12 h-12 flex items-center justify-center">
                      <span class="i-heroicons-credit-card text-blue-600 text-2xl"></span>
                    </div>
                  </div>
                  <div v-else class="w-20 h-20 rounded-xl border-2 border-gray-200 bg-gray-50 flex items-center justify-center">
                    <span class="text-gray-400 text-sm">انتخاب نشده</span>
                  </div>
                  <p class="text-xs text-gray-500">آیکون درگاه به صورت خودکار نمایش داده می‌شود</p>
                </div>
              </div>
              <div class="mt-8 space-y-3">
                <label class="block text-sm font-bold text-gray-700">توضیحات</label>
                <textarea
v-model="form.description" rows="4" 
                  class="w-full px-5 py-4 border-2 border-gray-200 rounded-xl text-sm focus:ring-4 focus:ring-blue-100 focus:border-blue-500 transition-all bg-white hover:bg-gray-50 shadow-sm resize-none"></textarea>
              </div>
            </div>
          </div>

          <!-- تنظیمات API -->
          <div v-if="activeTab === 'api'" class="space-y-8">
            <!-- فیلدهای عمومی API -->
            <div class="bg-gradient-to-br from-blue-50 to-indigo-50 rounded-2xl p-8 shadow-lg border border-blue-100">
              <h4 class="text-xl font-bold text-gray-900 mb-6 flex items-center">
                <div class="w-10 h-10 bg-gradient-to-r from-blue-500 to-indigo-600 rounded-xl flex items-center justify-center mr-4">
                  <span class="i-heroicons-key text-white text-lg"></span>
                </div>
                تنظیمات عمومی API
              </h4>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
                <div class="space-y-3">
                  <label class="block text-sm font-bold text-gray-700">کلید عمومی (Public Key)</label>
                  <input
v-model="form.apiKeys.publicKey" type="text" 
                    class="w-full px-5 py-4 border-2 border-blue-200 rounded-xl text-sm focus:ring-4 focus:ring-blue-100 focus:border-blue-500 transition-all bg-white hover:bg-blue-50 shadow-sm">
                </div>
                <div class="space-y-3">
                  <label class="block text-sm font-bold text-gray-700">کلید خصوصی (Private Key)</label>
                  <input
v-model="form.apiKeys.privateKey" type="text" 
                    class="w-full px-5 py-4 border-2 border-blue-200 rounded-xl text-sm focus:ring-4 focus:ring-blue-100 focus:border-blue-500 transition-all bg-white hover:bg-blue-50 shadow-sm">
                </div>
                <div class="space-y-3">
                  <label class="block text-sm font-bold text-gray-700">کلید تست (Test Key)</label>
                  <input
v-model="form.apiKeys.testKey" type="text" 
                    class="w-full px-5 py-4 border-2 border-blue-200 rounded-xl text-sm focus:ring-4 focus:ring-blue-100 focus:border-blue-500 transition-all bg-white hover:bg-blue-50 shadow-sm">
                </div>
                <div class="space-y-3">
                  <label class="block text-sm font-bold text-gray-700">شماره ترمینال</label>
                  <input
v-model="form.terminalId" type="text" 
                    class="w-full px-5 py-4 border-2 border-blue-200 rounded-xl text-sm focus:ring-4 focus:ring-blue-100 focus:border-blue-500 transition-all bg-white hover:bg-blue-50 shadow-sm">
                </div>
                <div class="space-y-3">
                  <label class="block text-sm font-bold text-gray-700">آدرس API پرداخت</label>
                  <input
v-model="form.apiEndpoints.payment" type="url" 
                    class="w-full px-5 py-4 border-2 border-blue-200 rounded-xl text-sm focus:ring-4 focus:ring-blue-100 focus:border-blue-500 transition-all bg-white hover:bg-blue-50 shadow-sm">
                </div>
                <div class="space-y-3">
                  <label class="block text-sm font-bold text-gray-700">آدرس API تایید</label>
                  <input
v-model="form.apiEndpoints.verification" type="url" 
                    class="w-full px-5 py-4 border-2 border-blue-200 rounded-xl text-sm focus:ring-4 focus:ring-blue-100 focus:border-blue-500 transition-all bg-white hover:bg-blue-50 shadow-sm">
                </div>
              </div>
            </div>

            <!-- فیلدهای خاص هر درگاه -->
            <div v-if="form.type" class="bg-gradient-to-br from-purple-50 to-pink-50 rounded-2xl p-8 shadow-lg border border-purple-100">
              <h4 class="text-xl font-bold text-gray-900 mb-6 flex items-center">
                <div class="w-10 h-10 bg-gradient-to-r from-purple-500 to-pink-600 rounded-xl flex items-center justify-center mr-4">
                  <span class="i-heroicons-cog text-white text-lg"></span>
                </div>
                تنظیمات خاص {{ getGatewayDisplayName(form.type) }}
              </h4>
              
              <!-- فیلدهای زرین‌پال -->
              <div v-if="form.type === 'zarinpal'" class="grid grid-cols-1 md:grid-cols-2 gap-8">
                <div class="space-y-3">
                  <label class="block text-sm font-bold text-gray-700">مرچنت ID *</label>
                  <input
v-model="form.merchantId" type="text" required 
                    class="w-full px-5 py-4 border-2 border-purple-200 rounded-xl text-sm focus:ring-4 focus:ring-purple-100 focus:border-purple-500 transition-all bg-white hover:bg-purple-50 shadow-sm" 
                    placeholder="مثال: 00000000-0000-0000-0000-000000000000">
                  <p class="mt-2 text-xs text-purple-600 bg-purple-50 px-3 py-2 rounded-lg">مرچنت ID دریافتی از زرین‌پال</p>
                </div>
                <div class="space-y-3">
                  <label class="block text-sm font-bold text-gray-700">آدرس Callback</label>
                  <input
v-model="form.apiEndpoints.callback" type="url" 
                    class="w-full px-5 py-4 border-2 border-purple-200 rounded-xl text-sm focus:ring-4 focus:ring-purple-100 focus:border-purple-500 transition-all bg-white hover:bg-purple-50 shadow-sm" 
                    placeholder="https://yourdomain.com/payment/callback/zarinpal">
                </div>
              </div>

              <!-- فیلدهای ملت -->
              <div v-if="form.type === 'mellat'" class="grid grid-cols-1 md:grid-cols-2 gap-8">
                <div class="space-y-3">
                  <label class="block text-sm font-bold text-gray-700">شماره ترمینال *</label>
                  <input
v-model="form.terminalId" type="text" required 
                    class="w-full px-5 py-4 border-2 border-purple-200 rounded-xl text-sm focus:ring-4 focus:ring-purple-100 focus:border-purple-500 transition-all bg-white hover:bg-purple-50 shadow-sm" 
                    placeholder="مثال: 12345678">
                  <p class="mt-2 text-xs text-purple-600 bg-purple-50 px-3 py-2 rounded-lg">شماره ترمینال دریافتی از بانک ملت</p>
                </div>
                <div class="space-y-3">
                  <label class="block text-sm font-bold text-gray-700">نام کاربری *</label>
                  <input
v-model="form.apiKeys.publicKey" type="text" required 
                    class="w-full px-5 py-4 border-2 border-purple-200 rounded-xl text-sm focus:ring-4 focus:ring-purple-100 focus:border-purple-500 transition-all bg-white hover:bg-purple-50 shadow-sm" 
                    placeholder="نام کاربری دریافتی از بانک ملت">
                </div>
                <div class="space-y-3">
                  <label class="block text-sm font-bold text-gray-700">رمز عبور *</label>
                  <input
v-model="form.apiKeys.privateKey" type="password" required 
                    class="w-full px-5 py-4 border-2 border-purple-200 rounded-xl text-sm focus:ring-4 focus:ring-purple-100 focus:border-purple-500 transition-all bg-white hover:bg-purple-50 shadow-sm" 
                    placeholder="رمز عبور دریافتی از بانک ملت">
                </div>
                <div class="space-y-3">
                  <label class="block text-sm font-bold text-gray-700">آدرس Callback</label>
                  <input
v-model="form.apiEndpoints.callback" type="url" 
                    class="w-full px-5 py-4 border-2 border-purple-200 rounded-xl text-sm focus:ring-4 focus:ring-purple-100 focus:border-purple-500 transition-all bg-white hover:bg-purple-50 shadow-sm" 
                    placeholder="https://yourdomain.com/payment/callback/mellat">
                </div>
              </div>

              <!-- فیلدهای ملی -->
              <div v-if="form.type === 'melli'" class="grid grid-cols-1 md:grid-cols-2 gap-8">
                <div class="space-y-3">
                  <label class="block text-sm font-bold text-gray-700">شماره ترمینال *</label>
                  <input
v-model="form.terminalId" type="text" required 
                    class="w-full px-5 py-4 border-2 border-purple-200 rounded-xl text-sm focus:ring-4 focus:ring-purple-100 focus:border-purple-500 transition-all bg-white hover:bg-purple-50 shadow-sm" 
                    placeholder="شماره ترمینال دریافتی از بانک ملی">
                  <p class="mt-2 text-xs text-purple-600 bg-purple-50 px-3 py-2 rounded-lg">شماره ترمینال دریافتی از بانک ملی</p>
                </div>
                <div class="space-y-3">
                  <label class="block text-sm font-bold text-gray-700">مرچنت ID *</label>
                  <input
v-model="form.merchantId" type="text" required 
                    class="w-full px-5 py-4 border-2 border-purple-200 rounded-xl text-sm focus:ring-4 focus:ring-purple-100 focus:border-purple-500 transition-all bg-white hover:bg-purple-50 shadow-sm" 
                    placeholder="مرچنت ID دریافتی از بانک ملی">
                </div>
                <div class="space-y-3">
                  <label class="block text-sm font-bold text-gray-700">کلید رمزنگاری *</label>
                  <input
v-model="form.apiKeys.privateKey" type="password" required 
                    class="w-full px-5 py-4 border-2 border-purple-200 rounded-xl text-sm focus:ring-4 focus:ring-purple-100 focus:border-purple-500 transition-all bg-white hover:bg-purple-50 shadow-sm" 
                    placeholder="کلید رمزنگاری دریافتی از بانک ملی">
                </div>
                <div class="space-y-3">
                  <label class="block text-sm font-bold text-gray-700">آدرس Callback</label>
                  <input
v-model="form.apiEndpoints.callback" type="url" 
                    class="w-full px-5 py-4 border-2 border-purple-200 rounded-xl text-sm focus:ring-4 focus:ring-purple-100 focus:border-purple-500 transition-all bg-white hover:bg-purple-50 shadow-sm" 
                    placeholder="https://yourdomain.com/payment/callback/melli">
                </div>
              </div>

              <!-- فیلدهای سامان -->
              <div v-if="form.type === 'saman'" class="grid grid-cols-1 md:grid-cols-2 gap-8">
                <div class="space-y-3">
                  <label class="block text-sm font-bold text-gray-700">مرچنت ID *</label>
                  <input
v-model="form.merchantId" type="text" required 
                    class="w-full px-5 py-4 border-2 border-purple-200 rounded-xl text-sm focus:ring-4 focus:ring-purple-100 focus:border-purple-500 transition-all bg-white hover:bg-purple-50 shadow-sm" 
                    placeholder="مرچنت ID دریافتی از بانک سامان">
                  <p class="mt-2 text-xs text-purple-600 bg-purple-50 px-3 py-2 rounded-lg">مرچنت ID دریافتی از بانک سامان</p>
                </div>
                <div class="space-y-3">
                  <label class="block text-sm font-bold text-gray-700">کلید خصوصی *</label>
                  <input
v-model="form.apiKeys.privateKey" type="password" required 
                    class="w-full px-5 py-4 border-2 border-purple-200 rounded-xl text-sm focus:ring-4 focus:ring-purple-100 focus:border-purple-500 transition-all bg-white hover:bg-purple-50 shadow-sm" 
                    placeholder="کلید خصوصی دریافتی از بانک سامان">
                </div>
                <div class="space-y-3">
                  <label class="block text-sm font-bold text-gray-700">آدرس Callback</label>
                  <input
v-model="form.apiEndpoints.callback" type="url" 
                    class="w-full px-5 py-4 border-2 border-purple-200 rounded-xl text-sm focus:ring-4 focus:ring-purple-100 focus:border-purple-500 transition-all bg-white hover:bg-purple-50 shadow-sm" 
                    placeholder="https://yourdomain.com/payment/callback/saman">
                </div>
              </div>
            </div>
          </div>

          <!-- تنظیمات امنیت -->
          <div v-if="activeTab === 'security'" class="space-y-8">
            <div class="bg-gradient-to-br from-yellow-50 to-orange-50 rounded-2xl p-8 shadow-lg border border-yellow-100">
              <h4 class="text-xl font-bold text-gray-900 mb-6 flex items-center">
                <div class="w-10 h-10 bg-gradient-to-r from-yellow-500 to-orange-600 rounded-xl flex items-center justify-center mr-4">
                  <span class="i-heroicons-shield-check text-white text-lg"></span>
                </div>
                تنظیمات امنیتی
              </h4>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
                <div class="space-y-3">
                  <label class="block text-sm font-bold text-gray-700">لیست IP های مجاز</label>
                  <textarea
v-model="form.security.ipWhitelist" rows="4" 
                    class="w-full px-5 py-4 border-2 border-yellow-200 rounded-xl text-sm focus:ring-4 focus:ring-yellow-100 focus:border-yellow-500 transition-all bg-white hover:bg-yellow-50 shadow-sm resize-none"
                    placeholder="هر IP در یک خط جداگانه"></textarea>
                  <p class="mt-2 text-xs text-yellow-700 bg-yellow-50 px-3 py-2 rounded-lg">IP های مجاز برای دسترسی به API</p>
                </div>
                <div class="space-y-6">
                  <div class="flex items-center space-x-4 space-x-reverse p-6 bg-white rounded-xl border-2 border-yellow-200">
                    <input
id="httpsRequired" v-model="form.security.httpsRequired" type="checkbox" 
                      class="w-5 h-5 text-yellow-600 border-yellow-300 rounded focus:ring-yellow-500">
                    <label for="httpsRequired" class="text-sm font-bold text-gray-700">اجبار استفاده از HTTPS</label>
                  </div>
                  <div class="space-y-3">
                    <label class="block text-sm font-bold text-gray-700">حداکثر تلاش</label>
                    <input
v-model.number="form.security.maxAttempts" type="number" min="1" max="10" 
                      class="w-full px-5 py-4 border-2 border-yellow-200 rounded-xl text-sm focus:ring-4 focus:ring-yellow-100 focus:border-yellow-500 transition-all bg-white hover:bg-yellow-50 shadow-sm">
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- تنظیمات تست -->
          <div v-if="activeTab === 'test'" class="space-y-8">
            <div class="bg-gradient-to-br from-green-50 to-emerald-50 rounded-2xl p-8 shadow-lg border border-green-100">
              <h4 class="text-xl font-bold text-gray-900 mb-6 flex items-center">
                <div class="w-10 h-10 bg-gradient-to-r from-green-500 to-emerald-600 rounded-xl flex items-center justify-center mr-4">
                  <span class="i-heroicons-beaker text-white text-lg"></span>
                </div>
                تنظیمات تست
              </h4>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
                <div class="flex items-center space-x-4 space-x-reverse p-6 bg-white rounded-xl border-2 border-green-200">
                  <input
id="testMode" v-model="form.testMode" type="checkbox" 
                    class="w-5 h-5 text-green-600 border-green-300 rounded focus:ring-green-500">
                  <label for="testMode" class="text-sm font-bold text-gray-700">فعال کردن حالت تست</label>
                </div>
                <div class="space-y-3">
                  <label class="block text-sm font-bold text-gray-700">شماره کارت تست</label>
                  <input
v-model="form.testCard" type="text" 
                    class="w-full px-5 py-4 border-2 border-green-200 rounded-xl text-sm focus:ring-4 focus:ring-green-100 focus:border-green-500 transition-all bg-white hover:bg-green-50 shadow-sm"
                    placeholder="شماره کارت تست">
                </div>
                <div class="space-y-3">
                  <label class="block text-sm font-bold text-gray-700">مبلغ تست (تومان)</label>
                  <input
v-model.number="form.testAmount" type="number" min="1000" 
                    class="w-full px-5 py-4 border-2 border-green-200 rounded-xl text-sm focus:ring-4 focus:ring-green-100 focus:border-green-500 transition-all bg-white hover:bg-green-50 shadow-sm">
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- دکمه‌های عملیات -->
        <div class="flex items-center justify-between pt-8 border-t border-gray-200 mt-8">
          <!-- دکمه انصراف -->
          <TemplateButton 
            type="button" 
            bg-gradient="bg-gradient-to-r from-gray-100 to-gray-200" 
            text-color="text-gray-700"
            border-color="border border-gray-300"
            hover-class="hover:from-gray-200 hover:to-gray-300"
            focus-class="focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
            size="large"
            custom-class="font-bold shadow-lg hover:shadow-xl transform hover:scale-105"
            @click="$emit('cancel')"
          >
            انصراف
          </TemplateButton>
          
          <!-- دکمه‌های ناوبری و ذخیره -->
          <div class="flex items-center space-x-4 space-x-reverse">
            <!-- دکمه قبلی -->
            <TemplateButton 
              v-if="hasPreviousTab()" 
              type="button" 
              bg-gradient="bg-gradient-to-r from-blue-50 to-blue-100"
              text-color="text-blue-600"
              border-color="border border-blue-300"
              hover-class="hover:from-blue-100 hover:to-blue-200"
              focus-class="focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
              size="large"
              custom-class="font-bold shadow-lg hover:shadow-xl transform hover:scale-105"
              @click="previousTab"
            >
              <span class="i-heroicons-arrow-right ml-2"></span>
              قبلی
            </TemplateButton>
            
            <!-- دکمه تست -->
            <TemplateButton 
              type="button" 
              bg-gradient="bg-gradient-to-r from-yellow-500 to-orange-500" 
              text-color="text-white"
              border-color="border border-yellow-500"
              hover-class="hover:from-yellow-600 hover:to-orange-600"
              focus-class="focus:ring-2 focus:ring-yellow-500 focus:ring-offset-2"
              size="large"
              custom-class="font-bold flex items-center shadow-lg mr-4"
              @click="handleSubmit"
            >
              <span class="i-heroicons-bug ml-2"></span>
              تست Submit
            </TemplateButton>
            
            <!-- دکمه ذخیره -->
            <TemplateButton 
              type="submit" 
              :disabled="saving" 
              bg-gradient="bg-gradient-to-r from-blue-500 to-indigo-600"
              text-color="text-white"
              border-color="border border-blue-500"
              hover-class="hover:from-blue-600 hover:to-indigo-700"
              focus-class="focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
              size="large"
              custom-class="font-bold flex items-center shadow-lg hover:shadow-xl transform hover:scale-105"
              style="box-shadow: 0 10px 25px -5px rgba(59, 130, 246, 0.3);"
            >
              <span v-if="saving" class="i-heroicons-arrow-path animate-spin ml-3"></span>
              <span v-else class="i-heroicons-check ml-3"></span>
              {{ saving ? 'در حال ذخیره...' : (props.mode === 'edit' ? 'بروزرسانی درگاه' : 'ذخیره درگاه') }}
            </TemplateButton>
          </div>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'
import TemplateButton from '~/components/common/TemplateButton.vue'

// تعریف interface برای فرم درگاه
interface GatewayFormData {
  type: string
  description: string
  fee: number
  minAmount: number
  maxAmount: number
  icon: string
  merchantId: string
  terminalId: string
  apiKeys: {
    publicKey: string
    privateKey: string
    testKey: string
    secretKey: string
  }
  apiEndpoints: {
    payment: string
    verification: string
    callback: string
    webhook: string
  }
  security: {
    ipWhitelist: string
    httpsRequired: boolean
    maxAttempts: number
  }
  testMode: boolean
  testCard: string
  testAmount: number
}

// تعریف props
interface Props {
  gateway?: any
  mode?: 'create' | 'edit'
}

const props = withDefaults(defineProps<Props>(), {
  mode: 'create'
})

// تعریف emit events
const emit = defineEmits<{
  save: [data: GatewayFormData]
  cancel: []
}>()

// متغیرهای reactive
const activeTab = ref('basic')
const saving = ref(false)
const validationErrors = ref<string[]>([])
const showErrors = ref(false)

// تب‌های فرم
const formTabs = [
  { value: 'basic', label: 'اطلاعات پایه' },
  { value: 'api', label: 'کلیدها و API' },
  { value: 'security', label: 'امنیت' },
  { value: 'test', label: 'تست' }
]





// درگاه‌های موجود
const existingGateways = ref<string[]>([])

// دریافت درگاه‌های موجود - بهینه‌سازی شده
const fetchExistingGateways = async () => {
  try {
    const response: any = await $fetch('/api/payment-gateways', {
      method: 'GET',
      headers: {
        'Cache-Control': 'max-age=300' // کش 5 دقیقه
      }
    })
    
    if (response.data) {
      existingGateways.value = response.data.map((gateway: any) => gateway.type)
    }
  } catch (error) {
    existingGateways.value = []
  }
}

// بررسی اینکه آیا درگاه قبلاً ثبت شده
const isGatewayExists = (type: string) => {
  // در حالت edit نیازی به بررسی نیست
  if (props.mode === 'edit') {
    return false
  }
  
  return existingGateways.value.includes(type)
}

// فرم درگاه
const form = ref<GatewayFormData>({
  type: '',
  description: '',
  fee: 0,
  minAmount: 0,
  maxAmount: 0,
  icon: '',
  merchantId: '',
  terminalId: '',
  apiKeys: {
    publicKey: '',
    privateKey: '',
    testKey: '',
    secretKey: ''
  },
  apiEndpoints: {
    payment: '',
    verification: '',
    callback: '',
    webhook: ''
  },
  security: {
    ipWhitelist: '',
    httpsRequired: true,
    maxAttempts: 5
  },
  testMode: false,
  testCard: '',
  testAmount: 10000
})

// تابع مقداردهی اولیه فرم در حالت ویرایش
const initializeForm = () => {
  if (props.mode === 'edit' && props.gateway) {
    
    // مقداردهی مستقیم فیلدها
    form.value.type = props.gateway.type || ''
    form.value.description = props.gateway.description || ''
    form.value.fee = props.gateway.fee || 0
    form.value.minAmount = props.gateway.min_amount || 0
    form.value.maxAmount = props.gateway.max_amount || 0
    form.value.icon = props.gateway.icon || ''
    form.value.merchantId = props.gateway.merchant_id || ''
    form.value.terminalId = props.gateway.terminal_id || ''
    form.value.testMode = props.gateway.is_test_mode || false
    
    // مقداردهی API Keys
    if (props.gateway.api_keys) {
      form.value.apiKeys.publicKey = props.gateway.api_keys.public_key || ''
      form.value.apiKeys.privateKey = props.gateway.api_keys.private_key || ''
      form.value.apiKeys.testKey = props.gateway.api_keys.test_key || ''
      form.value.apiKeys.secretKey = props.gateway.api_keys.secret_key || ''
    }
    
    // مقداردهی API Endpoints
    if (props.gateway.api_endpoints) {
      form.value.apiEndpoints.payment = props.gateway.api_endpoints.payment || ''
      form.value.apiEndpoints.verification = props.gateway.api_endpoints.verification || ''
      form.value.apiEndpoints.callback = props.gateway.api_endpoints.callback || ''
      form.value.apiEndpoints.webhook = props.gateway.api_endpoints.webhook || ''
    }
    
    // مقداردهی Security Settings
    if (props.gateway.settings) {
      form.value.security.ipWhitelist = props.gateway.settings.ip_whitelist || ''
      form.value.security.httpsRequired = props.gateway.settings.https_required || true
      form.value.security.maxAttempts = props.gateway.settings.max_attempts || 5
    }
    
    console.log('فرم مقداردهی شد')
  }
}

// مقداردهی اولیه در زمان mount
onMounted(() => {
  initializeForm()
  
  // فقط در حالت create درگاه‌های موجود را دریافت کن
  if (props.mode === 'create') {
    fetchExistingGateways()
  }
})

// پاک کردن خطاها هنگام تغییر فیلدها
// فقط وقتی فرم تغییر می‌کند، خطاها را پاک کن
watch(form, () => {
  if (showErrors.value) {
    showErrors.value = false
    validationErrors.value = []
  }
}, { deep: true, flush: 'post' })

// مقداردهی مجدد فرم وقتی gateway تغییر می‌کند
watch(() => props.gateway, (newGateway, oldGateway) => {
  // فقط اگر gateway واقعاً تغییر کرده باشد
  if (props.mode === 'edit' && newGateway && JSON.stringify(newGateway) !== JSON.stringify(oldGateway)) {
    initializeForm()
  }
}, { deep: true })

// تعریف interface برای تنظیمات پیش‌فرض درگاه
interface GatewayDefault {
  description: string
  apiEndpoints: {
    payment: string
    verification: string
  }
}

// تنظیم مقادیر پیش‌فرض بر اساس نوع درگاه
const gatewayDefaults: Record<string, GatewayDefault> = {
  zarinpal: {
    description: 'درگاه پرداخت زرین‌پال - ارائه خدمات پرداخت آنلاین امن',
    apiEndpoints: {
      payment: 'https://api.zarinpal.com/pg/v4/payment/request.json',
      verification: 'https://api.zarinpal.com/pg/v4/payment/verify.json'
    }
  },
  mellat: {
    description: 'درگاه پرداخت بانک ملت - یکی از معتبرترین درگاه‌های پرداخت ایران',
    apiEndpoints: {
      payment: 'https://bpm.shaparak.ir/pgwchannel/startpay.mellat',
      verification: 'https://bpm.shaparak.ir/pgwchannel/verifytransaction.mellat'
    }
  },
  melli: {
    description: 'درگاه پرداخت بانک ملی - ارائه خدمات پرداخت آنلاین امن',
    apiEndpoints: {
      payment: 'https://sadad.shaparak.ir/vpg/api/v0/Request/PaymentRequest',
      verification: 'https://sadad.shaparak.ir/vpg/api/v0/Request/PaymentVerification'
    }
  },
  saman: {
    description: 'درگاه پرداخت بانک سامان - ارائه خدمات پرداخت امن و سریع',
    apiEndpoints: {
      payment: 'https://sep.shaparak.ir/api/v1/Payment/GetToken',
      verification: 'https://sep.shaparak.ir/api/v1/Payment/VerifyPayment'
    }
  },
  parsian: {
    description: 'درگاه پرداخت بانک پارسیان - ارائه خدمات پرداخت آنلاین',
    apiEndpoints: {
      payment: 'https://pec.shaparak.ir/pecpaymentgateway/',
      verification: 'https://pec.shaparak.ir/pecpaymentgateway/'
    }
  },
  paypal: {
    description: 'درگاه پرداخت پی‌پال - ارائه خدمات پرداخت بین‌المللی',
    apiEndpoints: {
      payment: 'https://api-m.paypal.com/v1/payments/payment',
      verification: 'https://api-m.paypal.com/v1/payments/payment'
    }
  },
  stripe: {
    description: 'درگاه پرداخت استرایپ - ارائه خدمات پرداخت بین‌المللی',
    apiEndpoints: {
      payment: 'https://api.stripe.com/v1/payment_intents',
      verification: 'https://api.stripe.com/v1/payment_intents'
    }
  }
}

// تابع تغییر نوع درگاه
const onGatewayTypeChange = () => {
  if (form.value.type && gatewayDefaults[form.value.type]) {
    const defaults = gatewayDefaults[form.value.type]
    form.value.description = defaults.description
    form.value.apiEndpoints.payment = defaults.apiEndpoints.payment
    form.value.apiEndpoints.verification = defaults.apiEndpoints.verification
  }
  
  // پاک کردن خطاها
  showErrors.value = false
  validationErrors.value = []
}

// تابع دریافت نام نمایشی درگاه
const getGatewayDisplayName = (type: string) => {
  const names: Record<string, string> = {
    zarinpal: 'زرین‌پال',
    mellat: 'ملت',
    melli: 'ملی',
    saman: 'سامان',
    parsian: 'پارسیان',
    paypal: 'پی‌پال',
    stripe: 'استرایپ'
  }
  return names[type] || type
}

// توابع مدیریت تصاویر - حذف شده برای بهینه‌سازی

// توابع اعتبارسنجی تب‌ها
const validateBasicTab = () => {
  return form.value.type && form.value.fee > 0 && form.value.minAmount > 0 && form.value.maxAmount > 0
}

const validateApiTab = () => {
  if (!form.value.type) return false
  
  // اعتبارسنجی بر اساس نوع درگاه
  switch (form.value.type) {
    case 'zarinpal':
      return form.value.merchantId && form.value.apiEndpoints.payment && form.value.apiEndpoints.verification
    case 'mellat':
      return form.value.terminalId && form.value.apiKeys.publicKey && form.value.apiKeys.privateKey
    case 'melli':
      return form.value.terminalId && form.value.merchantId && form.value.apiKeys.privateKey
    case 'saman':
      return form.value.merchantId && form.value.apiKeys.privateKey
    case 'parsian':
      return form.value.terminalId && form.value.apiKeys.publicKey && form.value.apiKeys.privateKey
    case 'paypal':
    case 'stripe':
      return form.value.apiKeys.publicKey && form.value.apiKeys.privateKey
    default:
      return true
  }
}

const validateSecurityTab = () => {
  return form.value.security.maxAttempts > 0
}

const validateTestTab = () => {
  return true // تب تست اختیاری است
}

// بررسی تکمیل تمام تب‌ها
const isFormComplete = () => {
  return validateBasicTab() && validateApiTab() && validateSecurityTab() && validateTestTab()
}

// بررسی تکمیل تب فعلی
const isCurrentTabComplete = () => {
  switch (activeTab.value) {
    case 'basic':
      return validateBasicTab()
    case 'api':
      return validateApiTab()
    case 'security':
      return validateSecurityTab()
    case 'test':
      return validateTestTab()
    default:
      return false
  }
}

// بررسی تکمیل تب خاص
const isTabComplete = (tabValue: string) => {
  switch (tabValue) {
    case 'basic':
      return validateBasicTab()
    case 'api':
      return validateApiTab()
    case 'security':
      return validateSecurityTab()
    case 'test':
      return validateTestTab()
    default:
      return false
  }
}

// بررسی وجود تب بعدی
const hasNextTab = () => {
  const currentIndex = formTabs.findIndex(tab => tab.value === activeTab.value)
  return currentIndex < formTabs.length - 1
}

// بررسی وجود تب قبلی
const hasPreviousTab = () => {
  const currentIndex = formTabs.findIndex(tab => tab.value === activeTab.value)
  return currentIndex > 0
}

// توابع عملیات
const nextTab = () => {
  if (isCurrentTabComplete() && hasNextTab()) {
    const currentIndex = formTabs.findIndex(tab => tab.value === activeTab.value)
    activeTab.value = formTabs[currentIndex + 1].value
  }
}

const previousTab = () => {
  if (hasPreviousTab()) {
    const currentIndex = formTabs.findIndex(tab => tab.value === activeTab.value)
    activeTab.value = formTabs[currentIndex - 1].value
  }
}

// تابع اعتبارسنجی کامل فرم - بهینه‌سازی شده
const validateForm = () => {
  const errors: string[] = []
  
  // اعتبارسنجی اطلاعات پایه
  if (!form.value.type) {
    errors.push('لطفاً نوع درگاه را انتخاب کنید')
  }
  
  // اعتبارسنجی کارمزد - اگر صفر باشد نادیده گرفته می‌شود
  if (form.value.fee && form.value.fee < 0) {
    errors.push('کارمزد نمی‌تواند منفی باشد')
  }
  
  // اعتبارسنجی حداقل مبلغ - اگر صفر باشد نادیده گرفته می‌شود
  if (form.value.minAmount && form.value.minAmount < 0) {
    errors.push('حداقل مبلغ نمی‌تواند منفی باشد')
  }
  
  // اعتبارسنجی حداکثر مبلغ - اگر صفر باشد نادیده گرفته می‌شود
  if (form.value.maxAmount && form.value.maxAmount < 0) {
    errors.push('حداکثر مبلغ نمی‌تواند منفی باشد')
  }
  
  // اعتبارسنجی منطقی حداقل و حداکثر مبلغ - فقط اگر هر دو مقدار داشته باشند
  if (form.value.minAmount > 0 && form.value.maxAmount > 0 && form.value.minAmount >= form.value.maxAmount) {
    errors.push('حداقل مبلغ باید کمتر از حداکثر مبلغ باشد')
  }
  
  // اعتبارسنجی بر اساس نوع درگاه - فقط برای درگاه‌های جدید
  if (form.value.type && props.mode === 'create') {
    switch (form.value.type) {
      case 'zarinpal':
        if (!form.value.merchantId) {
          errors.push('مرچنت ID زرین‌پال الزامی است')
        }
        if (!form.value.apiEndpoints.payment) {
          errors.push('آدرس API پرداخت زرین‌پال الزامی است')
        }
        if (!form.value.apiEndpoints.verification) {
          errors.push('آدرس API تایید زرین‌پال الزامی است')
        }
        break
        
      case 'mellat':
        if (!form.value.terminalId) {
          errors.push('شماره ترمینال ملت الزامی است')
        }
        if (!form.value.apiKeys.publicKey) {
          errors.push('نام کاربری ملت الزامی است')
        }
        if (!form.value.apiKeys.privateKey) {
          errors.push('رمز عبور ملت الزامی است')
        }
        break
        
      case 'melli':
        if (!form.value.terminalId) {
          errors.push('شماره ترمینال ملی الزامی است')
        }
        if (!form.value.merchantId) {
          errors.push('مرچنت ID ملی الزامی است')
        }
        if (!form.value.apiKeys.privateKey) {
          errors.push('کلید رمزنگاری ملی الزامی است')
        }
        break
        
      case 'saman':
        if (!form.value.merchantId) {
          errors.push('مرچنت ID سامان الزامی است')
        }
        if (!form.value.apiKeys.privateKey) {
          errors.push('کلید خصوصی سامان الزامی است')
        }
        break
        
      case 'parsian':
        if (!form.value.terminalId) {
          errors.push('شماره ترمینال پارسیان الزامی است')
        }
        if (!form.value.apiKeys.publicKey) {
          errors.push('کلید عمومی پارسیان الزامی است')
        }
        if (!form.value.apiKeys.privateKey) {
          errors.push('کلید خصوصی پارسیان الزامی است')
        }
        break
        
      case 'paypal':
      case 'stripe':
        if (!form.value.apiKeys.publicKey) {
          errors.push('کلید عمومی الزامی است')
        }
        if (!form.value.apiKeys.privateKey) {
          errors.push('کلید خصوصی الزامی است')
        }
        break
    }
  }
  
  // اعتبارسنجی امنیت - فقط اگر مقدار داشته باشد
  if (form.value.security.maxAttempts && form.value.security.maxAttempts <= 0) {
    errors.push('حداکثر تعداد تلاش باید بزرگتر از صفر باشد')
  }
  
  return errors
}

// تابع تست submit - حذف شده برای بهینه‌سازی

const handleSubmit = async () => {
  // اعتبارسنجی فرم
  const errors = validateForm()
  
  if (errors.length > 0) {
    validationErrors.value = errors
    showErrors.value = true
    
    // اسکرول به بالای فرم برای نمایش خطاها
    setTimeout(() => {
      window.scrollTo({ top: 0, behavior: 'smooth' })
    }, 100)
    
    return
  }
  
  saving.value = true
  
  try {
    // ارسال داده‌ها به parent component
    emit('save', form.value)
  } catch (error) {
    validationErrors.value = ['خطا در ذخیره درگاه. لطفاً دوباره تلاش کنید.']
    showErrors.value = true
  } finally {
    saving.value = false
  }
}
</script>

<!--
  کامپوننت فرم افزودن درگاه پرداخت جدید
  شامل:
  1. فرم چندمرحله‌ای با تب‌های مختلف
  2. فیلدهای خاص هر درگاه بر اساس نوع انتخاب شده
  3. تنظیمات خودکار مقادیر پیش‌فرض
  4. اعتبارسنجی فیلدهای اجباری
  5. طراحی ریسپانسیو
  6. مدیریت state و validation
  توضیحات کامل به فارسی در کد
--> 
