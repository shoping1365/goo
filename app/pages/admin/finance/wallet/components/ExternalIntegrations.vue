<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between">
        <div>
          <h2 class="text-xl font-bold text-gray-900">اتصال به سیستم‌های خارجی</h2>
          <p class="text-gray-600 mt-1">یکپارچه‌سازی با سیستم‌های بانکی، پرداخت و حسابداری</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'banking'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
            @click="activeTab = 'banking'"
          >
            سیستم‌های بانکی
          </button>
          <button
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'payment'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
            @click="activeTab = 'payment'"
          >
            درگاه‌های پرداخت
          </button>
          <button
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'accounting'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
            @click="activeTab = 'accounting'"
          >
            سیستم‌های حسابداری
          </button>
          <button
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'api'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
            @click="activeTab = 'api'"
          >
            API و وب‌هوک
          </button>
        </div>
      </div>
    </div>

    <!-- محتوای تب‌ها -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200">
      <!-- سیستم‌های بانکی -->
      <div v-if="activeTab === 'banking'" class="p-6">
        <div class="space-y-6">
          <!-- اتصالات بانکی -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">اتصالات بانکی</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              <div v-for="bank in bankingConnections" :key="bank.id" class="border border-gray-200 rounded-lg p-6">
                <div class="flex items-center justify-between mb-3">
                  <div class="flex items-center space-x-3 space-x-reverse">
                    <div class="w-10 h-10 bg-blue-100 rounded-lg flex items-center justify-center">
                      <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z" />
                      </svg>
                    </div>
                    <div>
                      <h4 class="font-medium text-gray-900">{{ bank.name }}</h4>
                      <p class="text-sm text-gray-600">{{ bank.type }}</p>
                    </div>
                  </div>
                  <label class="relative inline-flex items-center cursor-pointer">
                    <input v-model="bank.connected" type="checkbox" class="sr-only peer">
                    <div class="w-9 h-5 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-4 after:w-4 after:transition-all peer-checked:bg-blue-600"></div>
                  </label>
                </div>
                <div class="space-y-2">
                  <div class="flex justify-between text-sm">
                    <span class="text-gray-600">وضعیت:</span>
                    <span :class="`${bank.connected ? 'text-green-600' : 'text-red-600'}`">
                      {{ bank.connected ? 'متصل' : 'قطع' }}
                    </span>
                  </div>
                  <div class="flex justify-between text-sm">
                    <span class="text-gray-600">آخرین همگام‌سازی:</span>
                    <span class="text-gray-900">{{ bank.lastSync }}</span>
                  </div>
                  <div class="flex justify-between text-sm">
                    <span class="text-gray-600">تعداد تراکنش:</span>
                    <span class="text-gray-900">{{ bank.transactionCount }}</span>
                  </div>
                </div>
                <div class="mt-3 flex space-x-2 space-x-reverse">
                  <button class="flex-1 bg-blue-600 text-white px-3 py-2 rounded text-sm hover:bg-blue-700 transition-colors">
                    تنظیمات
                  </button>
                  <button class="flex-1 bg-gray-600 text-white px-3 py-2 rounded text-sm hover:bg-gray-700 transition-colors">
                    تست اتصال
                  </button>
                </div>
              </div>
            </div>
          </div>

          <!-- تنظیمات بانکی -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات اتصال بانکی</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <h4 class="font-medium text-gray-900 mb-3">تنظیمات عمومی</h4>
                <div class="space-y-3">
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">فاصله همگام‌سازی</label>
                    <select v-model="bankingSettings.syncInterval" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                      <option value="5">5 دقیقه</option>
                      <option value="15">15 دقیقه</option>
                      <option value="30">30 دقیقه</option>
                      <option value="60">1 ساعت</option>
                    </select>
                  </div>
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">حداکثر تلاش اتصال</label>
                    <input v-model="bankingSettings.maxRetries" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500" min="1" max="10">
                  </div>
                  <div class="flex items-center">
                    <input v-model="bankingSettings.autoSync" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                    <label class="mr-2 text-sm text-gray-700">همگام‌سازی خودکار</label>
                  </div>
                </div>
              </div>
              <div>
                <h4 class="font-medium text-gray-900 mb-3">تنظیمات امنیتی</h4>
                <div class="space-y-3">
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">کلید API</label>
                    <input v-model="bankingSettings.apiKey" type="password" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="کلید API بانک">
                  </div>
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">رمز عبور</label>
                    <input v-model="bankingSettings.password" type="password" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="رمز عبور">
                  </div>
                  <div class="flex items-center">
                    <input v-model="bankingSettings.encryption" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                    <label class="mr-2 text-sm text-gray-700">رمزنگاری داده‌ها</label>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- درگاه‌های پرداخت -->
      <div v-if="activeTab === 'payment'" class="p-6">
        <div class="space-y-6">
          <!-- درگاه‌های فعال -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">درگاه‌های پرداخت</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              <div v-for="gateway in paymentGateways" :key="gateway.id" class="border border-gray-200 rounded-lg p-6">
                <div class="flex items-center justify-between mb-3">
                  <div class="flex items-center space-x-3 space-x-reverse">
                    <div class="w-10 h-10 bg-green-100 rounded-lg flex items-center justify-center">
                      <svg class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z" />
                      </svg>
                    </div>
                    <div>
                      <h4 class="font-medium text-gray-900">{{ gateway.name }}</h4>
                      <p class="text-sm text-gray-600">{{ gateway.type }}</p>
                    </div>
                  </div>
                  <span :class="`px-2 py-1 text-xs rounded-full ${gateway.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'}`">
                    {{ gateway.status === 'active' ? 'فعال' : 'غیرفعال' }}
                  </span>
                </div>
                <div class="space-y-2">
                  <div class="flex justify-between text-sm">
                    <span class="text-gray-600">کارمزد:</span>
                    <span class="text-gray-900">{{ gateway.fee }}%</span>
                  </div>
                  <div class="flex justify-between text-sm">
                    <span class="text-gray-600">حداکثر مبلغ:</span>
                    <span class="text-gray-900">{{ formatCurrency(gateway.maxAmount) }}</span>
                  </div>
                  <div class="flex justify-between text-sm">
                    <span class="text-gray-600">نرخ موفقیت:</span>
                    <span class="text-gray-900">{{ gateway.successRate }}%</span>
                  </div>
                </div>
                <div class="mt-3 flex space-x-2 space-x-reverse">
                  <button class="flex-1 bg-blue-600 text-white px-3 py-2 rounded text-sm hover:bg-blue-700 transition-colors">
                    تنظیمات
                  </button>
                  <button class="flex-1 bg-gray-600 text-white px-3 py-2 rounded text-sm hover:bg-gray-700 transition-colors">
                    آمار
                  </button>
                </div>
              </div>
            </div>
          </div>

          <!-- تنظیمات درگاه -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات درگاه پرداخت</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <h4 class="font-medium text-gray-900 mb-3">تنظیمات عمومی</h4>
                <div class="space-y-3">
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">درگاه پیش‌فرض</label>
                    <select v-model="paymentSettings.defaultGateway" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                      <option value="">انتخاب کنید</option>
                      <option value="mellat">ملت</option>
                      <option value="parsian">پارسیان</option>
                      <option value="saman">سامان</option>
                    </select>
                  </div>
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">زمان انتظار (ثانیه)</label>
                    <input v-model="paymentSettings.timeout" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500" min="30" max="300">
                  </div>
                  <div class="flex items-center">
                    <input v-model="paymentSettings.autoRetry" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                    <label class="mr-2 text-sm text-gray-700">تلاش مجدد خودکار</label>
                  </div>
                </div>
              </div>
              <div>
                <h4 class="font-medium text-gray-900 mb-3">تنظیمات امنیتی</h4>
                <div class="space-y-3">
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">کلید API</label>
                    <input v-model="paymentSettings.apiKey" type="password" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="کلید API درگاه">
                  </div>
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">کلید امنیتی</label>
                    <input v-model="paymentSettings.secretKey" type="password" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="کلید امنیتی">
                  </div>
                  <div class="flex items-center">
                    <input v-model="paymentSettings.sslVerification" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                    <label class="mr-2 text-sm text-gray-700">تایید SSL</label>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- سیستم‌های حسابداری -->
      <div v-if="activeTab === 'accounting'" class="p-6">
        <div class="space-y-6">
          <!-- اتصالات حسابداری -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">سیستم‌های حسابداری</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div v-for="system in accountingSystems" :key="system.id" class="border border-gray-200 rounded-lg p-6">
                <div class="flex items-center justify-between mb-3">
                  <div class="flex items-center space-x-3 space-x-reverse">
                    <div class="w-10 h-10 bg-purple-100 rounded-lg flex items-center justify-center">
                      <svg class="w-5 h-5 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 7h6m0 10v-3m-3 3h.01M9 17h.01M9 14h.01M12 14h.01M15 11h.01M12 11h.01M9 11h.01M7 21h10a2 2 0 002-2V5a2 2 0 00-2-2H7a2 2 0 00-2 2v14a2 2 0 002 2z" />
                      </svg>
                    </div>
                    <div>
                      <h4 class="font-medium text-gray-900">{{ system.name }}</h4>
                      <p class="text-sm text-gray-600">{{ system.version }}</p>
                    </div>
                  </div>
                  <label class="relative inline-flex items-center cursor-pointer">
                    <input v-model="system.connected" type="checkbox" class="sr-only peer">
                    <div class="w-9 h-5 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-4 after:w-4 after:transition-all peer-checked:bg-blue-600"></div>
                  </label>
                </div>
                <div class="space-y-2">
                  <div class="flex justify-between text-sm">
                    <span class="text-gray-600">نوع اتصال:</span>
                    <span class="text-gray-900">{{ system.connectionType }}</span>
                  </div>
                  <div class="flex justify-between text-sm">
                    <span class="text-gray-600">آخرین همگام‌سازی:</span>
                    <span class="text-gray-900">{{ system.lastSync }}</span>
                  </div>
                  <div class="flex justify-between text-sm">
                    <span class="text-gray-600">تعداد رکورد:</span>
                    <span class="text-gray-900">{{ system.recordCount }}</span>
                  </div>
                </div>
                <div class="mt-3 flex space-x-2 space-x-reverse">
                  <button class="flex-1 bg-blue-600 text-white px-3 py-2 rounded text-sm hover:bg-blue-700 transition-colors">
                    تنظیمات
                  </button>
                  <button class="flex-1 bg-green-600 text-white px-3 py-2 rounded text-sm hover:bg-green-700 transition-colors">
                    همگام‌سازی
                  </button>
                </div>
              </div>
            </div>
          </div>

          <!-- تنظیمات همگام‌سازی -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات همگام‌سازی</h3>
            <div class="space-y-4">
              <div class="flex items-center justify-between">
                <div>
                  <h4 class="font-medium text-gray-900">همگام‌سازی خودکار</h4>
                  <p class="text-sm text-gray-600">همگام‌سازی خودکار داده‌ها با سیستم‌های حسابداری</p>
                </div>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input v-model="accountingSettings.autoSync" type="checkbox" class="sr-only peer">
                  <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
                </label>
              </div>
              <div class="flex items-center justify-between">
                <div>
                  <h4 class="font-medium text-gray-900">همگام‌سازی دوطرفه</h4>
                  <p class="text-sm text-gray-600">اجازه تغییر داده‌ها از سیستم‌های خارجی</p>
                </div>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input v-model="accountingSettings.bidirectionalSync" type="checkbox" class="sr-only peer">
                  <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
                </label>
              </div>
              <div class="flex items-center justify-between">
                <div>
                  <h4 class="font-medium text-gray-900">تایید تراکنش‌ها</h4>
                  <p class="text-sm text-gray-600">تایید دستی تراکنش‌ها قبل از همگام‌سازی</p>
                </div>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input v-model="accountingSettings.manualApproval" type="checkbox" class="sr-only peer">
                  <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
                </label>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- API و وب‌هوک -->
      <div v-if="activeTab === 'api'" class="p-6">
        <div class="space-y-6">
          <!-- وب‌هوک‌های فعال -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">وب‌هوک‌های فعال</h3>
            <div class="space-y-4">
              <div v-for="webhook in webhooks" :key="webhook.id" class="border border-gray-200 rounded-lg p-6">
                <div class="flex items-center justify-between mb-3">
                  <div>
                    <h4 class="font-medium text-gray-900">{{ webhook.name }}</h4>
                    <p class="text-sm text-gray-600">{{ webhook.url }}</p>
                  </div>
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <span :class="`px-2 py-1 text-xs rounded-full ${webhook.status === 'active' ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'}`">
                      {{ webhook.status === 'active' ? 'فعال' : 'غیرفعال' }}
                    </span>
                    <button class="text-blue-600 hover:text-blue-800">
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                      </svg>
                    </button>
                  </div>
                </div>
                <div class="grid grid-cols-2 md:grid-cols-4 gap-6 text-sm">
                  <div>
                    <span class="text-gray-600">رویداد:</span>
                    <span class="text-gray-900">{{ webhook.event }}</span>
                  </div>
                  <div>
                    <span class="text-gray-600">آخرین ارسال:</span>
                    <span class="text-gray-900">{{ webhook.lastSent }}</span>
                  </div>
                  <div>
                    <span class="text-gray-600">نرخ موفقیت:</span>
                    <span class="text-gray-900">{{ webhook.successRate }}%</span>
                  </div>
                  <div>
                    <span class="text-gray-600">تعداد ارسال:</span>
                    <span class="text-gray-900">{{ webhook.sentCount }}</span>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- تنظیمات API -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات API</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <h4 class="font-medium text-gray-900 mb-3">تنظیمات عمومی</h4>
                <div class="space-y-3">
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">محدودیت نرخ درخواست</label>
                    <input v-model="apiSettings.rateLimit" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="درخواست در دقیقه">
                  </div>
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">زمان انتظار (ثانیه)</label>
                    <input v-model="apiSettings.timeout" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500" min="5" max="60">
                  </div>
                  <div class="flex items-center">
                    <input v-model="apiSettings.corsEnabled" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                    <label class="mr-2 text-sm text-gray-700">فعال‌سازی CORS</label>
                  </div>
                </div>
              </div>
              <div>
                <h4 class="font-medium text-gray-900 mb-3">تنظیمات امنیتی</h4>
                <div class="space-y-3">
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">کلید API</label>
                    <div class="flex space-x-2 space-x-reverse">
                      <input v-model="apiSettings.apiKey" type="password" class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500" readonly>
                      <button class="bg-blue-600 text-white px-3 py-2 rounded hover:bg-blue-700 transition-colors">
                        تولید
                      </button>
                    </div>
                  </div>
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">کلید امنیتی</label>
                    <div class="flex space-x-2 space-x-reverse">
                      <input v-model="apiSettings.secretKey" type="password" class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500" readonly>
                      <button class="bg-blue-600 text-white px-3 py-2 rounded hover:bg-blue-700 transition-colors">
                        تولید
                      </button>
                    </div>
                  </div>
                  <div class="flex items-center">
                    <input v-model="apiSettings.ipWhitelist" type="checkbox" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                    <label class="mr-2 text-sm text-gray-700">محدودیت IP</label>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
// تعریف reactive data
const activeTab = ref('banking')

// اتصالات بانکی
const bankingConnections = ref([
  {
    id: 1,
    name: 'بانک ملت',
    type: 'شتاب',
    connected: true,
    lastSync: '2024-01-31 14:30',
    transactionCount: 1250
  },
  {
    id: 2,
    name: 'بانک پارسیان',
    type: 'شتاب',
    connected: true,
    lastSync: '2024-01-31 14:25',
    transactionCount: 890
  },
  {
    id: 3,
    name: 'بانک سامان',
    type: 'شتاب',
    connected: false,
    lastSync: '2024-01-30 16:45',
    transactionCount: 650
  }
])

// تنظیمات بانکی
const bankingSettings = ref({
  syncInterval: 15,
  maxRetries: 3,
  autoSync: true,
  apiKey: '',
  password: '',
  encryption: true
})

// درگاه‌های پرداخت
const paymentGateways = ref([
  {
    id: 1,
    name: 'ملت',
    type: 'شتاب',
    status: 'active',
    fee: 1.5,
    maxAmount: 50000000,
    successRate: 98.5
  },
  {
    id: 2,
    name: 'پارسیان',
    type: 'شتاب',
    status: 'active',
    fee: 1.8,
    maxAmount: 100000000,
    successRate: 97.2
  },
  {
    id: 3,
    name: 'سامان',
    type: 'شتاب',
    status: 'active',
    fee: 1.2,
    maxAmount: 25000000,
    successRate: 99.1
  }
])

// تنظیمات درگاه
const paymentSettings = ref({
  defaultGateway: 'mellat',
  timeout: 60,
  autoRetry: true,
  apiKey: '',
  secretKey: '',
  sslVerification: true
})

// سیستم‌های حسابداری
const accountingSystems = ref([
  {
    id: 1,
    name: 'هلو',
    version: 'v5.0',
    connected: true,
    connectionType: 'API',
    lastSync: '2024-01-31 14:30',
    recordCount: 15420
  },
  {
    id: 2,
    name: 'پارسیان',
    version: 'v3.2',
    connected: false,
    connectionType: 'فایل',
    lastSync: '2024-01-30 16:45',
    recordCount: 8900
  }
])

// تنظیمات حسابداری
const accountingSettings = ref({
  autoSync: true,
  bidirectionalSync: false,
  manualApproval: true
})

// وب‌هوک‌ها
const webhooks = ref([
  {
    id: 1,
    name: 'اعلان تراکنش جدید',
    url: 'https://api.example.com/webhooks/transaction',
    status: 'active',
    event: 'transaction.created',
    lastSent: '2024-01-31 14:30',
    successRate: 99.5,
    sentCount: 15420
  },
  {
    id: 2,
    name: 'اعلان خطا',
    url: 'https://api.example.com/webhooks/error',
    status: 'active',
    event: 'error.occurred',
    lastSent: '2024-01-31 14:25',
    successRate: 100,
    sentCount: 23
  }
])

// تنظیمات API
const apiSettings = ref({
  rateLimit: 1000,
  timeout: 30,
  corsEnabled: true,
  apiKey: 'sk_live_1234567890abcdef',
  secretKey: 'sk_secret_abcdef1234567890',
  ipWhitelist: false
})

// تابع فرمت‌بندی ارز
const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR', {
    style: 'currency',
    currency: 'IRR',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount)
}
</script> 
