<template>
  <ClientOnly>
    <!-- Page Header -->
    <div class="bg-white shadow-sm border-b border-gray-200 mb-4">
      <div class="w-full px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center py-3">
          <div>
            <h1 class="text-lg font-bold text-gray-900">تشخیص تقلب در سفارشات</h1>
            <p class="text-xs text-gray-600 mt-1">مدیریت و نظارت بر فعالیت‌های مشکوک و تقلبی در سفارشات</p>
          </div>
          <div class="flex space-x-2 space-x-reverse">


            <button 
              @click="showSettings"
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-blue-400 to-blue-600 hover:from-blue-500 hover:to-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105"
            >
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"></path>
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
              </svg>
              تنظیمات
            </button>
          </div>
        </div>
      </div>
    </div>

    <div class="w-full px-4 py-4">
      
      <!-- تب‌های اصلی -->
      <div class="bg-white border-b border-gray-200 mb-6">
        <nav class="-mb-px flex px-6 overflow-x-auto">
            <button
              v-for="tab in tabs"
              :key="tab.id"
              @click="activeTab = tab.id as 'dashboard' | 'bot-detection' | 'reports' | 'settings' | 'fraud-orders'"
              :class="[
                activeTab === tab.id
                  ? 'border-red-500 text-red-600'
                  : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300',
              'flex-1 whitespace-nowrap py-4 px-4 border-b-2 font-medium text-sm text-center'
              ]"
            >
              {{ tab.name }}
            </button>
          </nav>
        </div>

        <div class="p-6">
          <!-- تب داشبورد -->
          <div v-if="activeTab === 'dashboard'">
            <!-- آمار کلی تشخیص تقلب -->
            <div class="rounded-lg shadow-md border border-gray-200 overflow-hidden mb-6">
              <div class="bg-gradient-to-r from-red-50 to-pink-50 px-4 py-3 border-b border-gray-200">
                <div class="flex items-center">
                  <div class="w-6 h-6 bg-red-500 rounded-md flex items-center justify-center ml-2">
                    <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
                    </svg>
                  </div>
                  <h3 class="text-sm font-semibold text-gray-900">آمار کلی تشخیص تقلب</h3>
                </div>
              </div>
              
              <div class="p-6">
                <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-3">
                  <!-- سفارشات مشکوک -->
                  <div class="bg-gradient-to-br from-red-50 to-red-100 border border-red-200 rounded-md p-3 hover:shadow-sm transition-all duration-200">
                    <div class="flex items-center justify-between">
                      <div>
                        <p class="text-xs font-medium text-red-700 mb-0.5">سفارشات مشکوک</p>
                        <p class="text-lg font-bold text-red-900">{{ stats.suspiciousOrders.toLocaleString('en-US') }}</p>
                      </div>
                      <div class="w-8 h-8 bg-red-500 rounded-md flex items-center justify-center">
                        <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
                        </svg>
                      </div>
                    </div>
                  </div>

                  <!-- تقلب تایید شده -->
                  <div class="bg-gradient-to-br from-orange-50 to-orange-100 border border-orange-200 rounded-md p-3 hover:shadow-sm transition-all duration-200">
                    <div class="flex items-center justify-between">
                      <div>
                        <p class="text-xs font-medium text-orange-700 mb-0.5">تقلب تایید شده</p>
                        <p class="text-lg font-bold text-orange-900">{{ stats.confirmedFraud.toLocaleString('en-US') }}</p>
                      </div>
                      <div class="w-8 h-8 bg-orange-500 rounded-md flex items-center justify-center">
                        <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
                        </svg>
                      </div>
                    </div>
                  </div>

                  <!-- دقت تشخیص -->
                  <div class="bg-gradient-to-br from-green-50 to-green-100 border border-green-200 rounded-md p-3 hover:shadow-sm transition-all duration-200">
                    <div class="flex items-center justify-between">
                      <div>
                        <p class="text-xs font-medium text-green-700 mb-0.5">دقت تشخیص</p>
                        <p class="text-lg font-bold text-green-900">{{ stats.accuracy }}%</p>
                      </div>
                      <div class="w-8 h-8 bg-green-500 rounded-md flex items-center justify-center">
                        <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                        </svg>
                      </div>
                    </div>
                  </div>

                  <!-- IP های مسدود شده -->
                  <div class="bg-gradient-to-br from-purple-50 to-purple-100 border border-purple-200 rounded-md p-3 hover:shadow-sm transition-all duration-200">
                    <div class="flex items-center justify-between">
                      <div>
                        <p class="text-xs font-medium text-purple-700 mb-0.5">IP های مسدود شده</p>
                        <p class="text-lg font-bold text-purple-900">{{ stats.blockedIPs.toLocaleString('en-US') }}</p>
                      </div>
                      <div class="w-8 h-8 bg-purple-500 rounded-md flex items-center justify-center">
                        <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728L5.636 5.636m12.728 12.728L18.364 5.636M5.636 18.364l12.728-12.728"></path>
                        </svg>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <!-- فیلترهای پیشرفته -->
            <div class="bg-gradient-to-br from-gray-50 to-gray-100 rounded-xl shadow-lg border border-gray-200 p-6 mb-6 w-full">
              <!-- Advanced Filters -->
              <div class="flex items-center justify-between">
                <div class="flex items-center">
                  <button @click="showAdvanced = !showAdvanced" class="w-8 h-8 bg-gradient-to-r from-red-500 to-pink-600 rounded-lg flex items-center justify-center ml-3 text-white hover:bg-red-600 transition-colors">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.207A1 1 0 013 6.5V4z"></path>
                    </svg>
                  </button>
                  <h3 class="text-lg font-semibold text-gray-900">فیلترهای سفارشات تقلبی</h3>
                </div>
                <span class="text-sm text-gray-600 cursor-pointer hover:text-red-600 transition-colors" @click="showAdvanced = !showAdvanced">
                  {{ showAdvanced ? 'مخفی کردن' : 'نمایش' }}
                </span>
              </div>

              <div v-show="showAdvanced" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
                <!-- Search -->
                <div class="space-y-2">
                  <label class="block text-sm font-semibold text-gray-700">جستجو</label>
                  <div class="relative">
                    <input
                      v-model="searchQuery"
                      type="text"
                      placeholder="شماره سفارش، نام مشتری..."
                      class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-red-500 focus:border-red-500 transition-all duration-200 bg-white shadow-sm"
                    />
                    <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                      <svg class="h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
                      </svg>
                    </div>
                  </div>
                </div>

                <!-- Fraud Type -->
                <div class="space-y-2">
                  <label class="block text-sm font-semibold text-gray-700">نوع تقلب</label>
                  <select
                    v-model="fraudTypeFilter"
                    class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-red-500 focus:border-red-500 transition-all duration-200 bg-white shadow-sm"
                  >
                    <option value="">همه انواع</option>
                    <option value="fake_address">آدرس جعلی</option>
                    <option value="stolen_card">کارت دزدیده شده</option>
                    <option value="bot_activity">فعالیت بوت</option>
                    <option value="multiple_accounts">حساب‌های متعدد</option>
                    <option value="fake_identity">هویت جعلی</option>
                  </select>
                </div>

                <!-- Risk Level -->
                <div class="space-y-2">
                  <label class="block text-sm font-semibold text-gray-700">سطح ریسک</label>
                  <select
                    v-model="riskLevelFilter"
                    class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-red-500 focus:border-red-500 transition-all duration-200 bg-white shadow-sm"
                  >
                    <option value="">همه سطوح</option>
                    <option value="high">ریسک بالا</option>
                    <option value="medium">ریسک متوسط</option>
                    <option value="low">ریسک پایین</option>
                  </select>
                </div>

                <!-- Payment Method -->
                <div class="space-y-2">
                  <label class="block text-sm font-semibold text-gray-700">روش پرداخت</label>
                  <select
                    v-model="paymentMethodFilter"
                    class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-red-500 focus:border-red-500 transition-all duration-200 bg-white shadow-sm"
                  >
                    <option value="">همه روش‌ها</option>
                    <option value="online">آنلاین</option>
                    <option value="cash">نقدی</option>
                    <option value="bank_transfer">انتقال بانکی</option>
                    <option value="wallet">کیف پول</option>
                  </select>
                </div>

                <!-- Fraud Score -->
                <div class="space-y-2">
                  <label class="block text-sm font-semibold text-gray-700">امتیاز تقلب</label>
                  <select
                    v-model="fraudScoreFilter"
                    class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-red-500 focus:border-red-500 transition-all duration-200 bg-white shadow-sm"
                  >
                    <option value="">همه امتیازات</option>
                    <option value="90-100">90-100% (خیلی بالا)</option>
                    <option value="80-89">80-89% (بالا)</option>
                    <option value="70-79">70-79% (متوسط)</option>
                    <option value="60-69">60-69% (پایین)</option>
                  </select>
                </div>

                <!-- Amount Range -->
                <div class="space-y-2">
                  <label class="block text-sm font-semibold text-gray-700">حداقل مبلغ (تومان)</label>
                  <input
                    v-model="minAmountFilter"
                    type="number"
                    placeholder="0"
                    class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-red-500 focus:border-red-500 transition-all duration-200 bg-white shadow-sm"
                  />
                </div>

                <div class="space-y-2">
                  <label class="block text-sm font-semibold text-gray-700">حداکثر مبلغ (تومان)</label>
                  <input
                    v-model="maxAmountFilter"
                    type="number"
                    placeholder="نامحدود"
                    class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-red-500 focus:border-red-500 transition-all duration-200 bg-white shadow-sm"
                  />
                </div>

                <!-- Date Range -->
                <div class="space-y-2">
                  <label class="block text-sm font-semibold text-gray-700">از تاریخ</label>
                  <input
                    v-model="dateFromFilter"
                    type="date"
                    class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-red-500 focus:border-red-500 transition-all duration-200 bg-white shadow-sm"
                  />
                </div>

                <div class="space-y-2">
                  <label class="block text-sm font-semibold text-gray-700">تا تاریخ</label>
                  <input
                    v-model="dateToFilter"
                    type="date"
                    class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-red-500 focus:border-red-500 transition-all duration-200 bg-white shadow-sm"
                  />
                </div>
              </div>

              <!-- Active Filters -->
              <div v-if="hasActiveFilters" class="mt-6 pt-6 border-t border-gray-200">
                <div class="flex items-center space-x-2 flex-wrap gap-2">
                  <span class="text-sm font-semibold text-gray-700">فیلترهای انتخاب شده:</span>
                  <span
                    v-for="(value, key) in activeFilters"
                    :key="key"
                    class="inline-flex items-center px-3 py-1.5 rounded-full text-xs font-medium bg-gradient-to-r from-red-100 to-red-200 text-red-800 border border-red-300 shadow-sm"
                  >
                    {{ getFilterLabel(key, value) }}
                    <button @click="removeFilter(key)" class="mr-1 text-red-600 hover:text-red-800 transition-colors">
                      <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                      </svg>
                    </button>
                  </span>
                </div>
              </div>
            </div>

            <!-- بخش سفارشات تقلبی -->
            <div class="bg-white rounded-lg shadow-md border border-gray-200 mb-6">
              <div class="p-6 border-b border-gray-200">
                <div class="flex items-center justify-between">
                  <h3 class="text-lg font-semibold text-gray-900">سفارشات تقلبی</h3>
                  <div class="flex items-center gap-2">
                    <button @click="fetchFraudCases" class="text-sm bg-gray-100 px-3 py-1.5 rounded">بروزرسانی</button>
                    <button @click="openRules" class="text-sm bg-blue-600 text-white px-3 py-1.5 rounded">قوانین</button>
                  </div>
                </div>
              </div>

              <div class="overflow-x-auto">
                <table class="min-w-full divide-y divide-gray-200">
                  <thead class="bg-gray-50">
                    <tr>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">شماره سفارش</th>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مشتری</th>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مبلغ</th>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع تقلب</th>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">امتیاز تقلب</th>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت بلاک</th>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ</th>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
                    </tr>
                  </thead>
                  <tbody class="bg-white divide-y divide-gray-200">
                    <tr v-for="order in paginatedFraudOrders" :key="order.id" class="hover:bg-gray-50">
                      <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                        #{{ order.orderNumber }}
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap">
                        <div>
                          <div class="text-sm font-medium text-gray-900">{{ order.customerName }}</div>
                          <div class="text-sm text-gray-500">{{ order.customerPhone }}</div>
                        </div>
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                        {{ formatPrice(order.totalAmount) }}
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap">
                        <span :class="getFraudTypeClass(order.fraudType)" class="px-2 py-1 text-xs font-medium rounded-full">
                          {{ getFraudTypeText(order.fraudType) }}
                        </span>
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap">
                        <div class="flex items-center">
                          <div class="w-16 bg-gray-200 rounded-full h-2 ml-2">
                            <div class="bg-red-600 h-2 rounded-full" :style="{ width: order.fraudScore + '%' }"></div>
                          </div>
                          <span class="text-sm font-medium text-red-600">{{ order.fraudScore }}%</span>
                        </div>
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap">
                        <div class="flex flex-col">
                          <span :class="getBlockStatusClass(order.blockStatus)" class="px-2 py-1 text-xs font-medium rounded-full mb-1">
                            {{ getBlockStatusText(order.blockStatus) }}
                          </span>
                          <span v-if="order.blockStatus === 'manual' && order.blockedBy" class="text-xs text-gray-500">
                            توسط: {{ order.blockedBy }}
                          </span>
                        </div>
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                        {{ formatDate(order.createdAt) }}
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                        <div class="flex items-center space-x-2">
                          <button class="text-blue-600 hover:text-blue-900" @click="viewFraudOrder(order)">جزئیات</button>
                          <button class="text-green-600 hover:text-green-900" @click="approveCustomer(order)">تایید</button>
                          <button class="text-yellow-600 hover:text-yellow-900" @click="markReview(order)">ارسال به بررسی</button>
                          <button class="text-red-600 hover:text-red-900" @click="blockCustomer(order)">مسدود</button>
                        </div>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>

              <!-- صفحه‌بندی -->
              <div v-if="filteredFraudOrders.length > 0">
                <AdminPagination 
                  :current-page="currentPage"
                  :total-pages="totalPages"
                  :total="filteredFraudOrders.length"
                  :items-per-page="itemsPerPage"
                  @page-changed="handlePageChange"
                  @items-per-page-changed="handleItemsPerPageChange"
                />
              </div>
            </div>

            <!-- آمار خریدهای مشکوک -->
            <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-6">
              <!-- آمار کلی خریدهای مشکوک -->
              <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
                <h3 class="text-lg font-semibold text-gray-900 mb-4">آمار خریدهای مشکوک</h3>
                <div class="space-y-4">
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-700">کل خریدهای بررسی شده</span>
                    <span class="text-sm font-medium">۱۵,۶۷۸</span>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-700">خریدهای مشکوک</span>
                    <span class="text-sm font-medium text-red-600">۱۲۳</span>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-700">خریدهای تایید شده جعلی</span>
                    <span class="text-sm font-medium text-red-600">۸۹</span>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-700">دقت تشخیص</span>
                    <span class="text-sm font-medium text-green-600">۹۲.۳%</span>
                  </div>
                </div>
              </div>
              
              <!-- معیارهای تشخیص -->
              <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
                <h3 class="text-lg font-semibold text-gray-900 mb-4">معیارهای تشخیص</h3>
                <div class="space-y-4">
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-700">سرعت خرید غیرعادی</span>
                    <div class="flex items-center">
                      <div class="w-16 bg-gray-200 rounded-full h-2 ml-2">
                        <div class="bg-red-600 h-2 rounded-full" style="width: 85%"></div>
                      </div>
                      <span class="text-sm font-medium">۸۵%</span>
                    </div>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-700">آدرس‌های غیرواقعی</span>
                    <div class="flex items-center">
                      <div class="w-16 bg-gray-200 rounded-full h-2 ml-2">
                        <div class="bg-orange-600 h-2 rounded-full" style="width: 72%"></div>
                      </div>
                      <span class="text-sm font-medium">۷۲%</span>
                    </div>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-700">کارت‌های اعتباری مشکوک</span>
                    <div class="flex items-center">
                      <div class="w-16 bg-gray-200 rounded-full h-2 ml-2">
                        <div class="bg-yellow-600 h-2 rounded-full" style="width: 68%"></div>
                      </div>
                      <span class="text-sm font-medium">۶۸%</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <!-- آمار انواع تقلب -->
            <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
              <!-- نمودار انواع تقلب -->
              <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
                <h3 class="text-lg font-semibold text-gray-900 mb-4">توزیع انواع تقلب</h3>
                <div class="space-y-4">
                  <div v-for="(count, type) in fraudTypeStats" :key="type" class="flex items-center justify-between">
                    <div class="flex items-center">
                      <div :class="getFraudTypeClass(type)" class="w-3 h-3 rounded-full ml-3"></div>
                      <span class="text-sm text-gray-700">{{ getFraudTypeText(type) }}</span>
                    </div>
                    <div class="flex items-center">
                      <div class="w-24 bg-gray-200 rounded-full h-2 ml-3">
                        <div :class="getFraudTypeClass(type).replace('text-', 'bg-').replace('-100', '-500')" class="h-2 rounded-full" :style="{ width: (count / fraudOrders.length * 100) + '%' }"></div>
                      </div>
                      <span class="text-sm font-medium text-gray-900">{{ count }}</span>
                    </div>
                  </div>
                </div>
              </div>

              <!-- آمار سطح ریسک -->
              <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
                <h3 class="text-lg font-semibold text-gray-900 mb-4">توزیع سطح ریسک</h3>
                <div class="space-y-4">
                  <div v-for="(count, level) in riskLevelStats" :key="level" class="flex items-center justify-between">
                    <div class="flex items-center">
                      <div :class="getRiskLevelClass(level)" class="w-3 h-3 rounded-full ml-3"></div>
                      <span class="text-sm text-gray-700">{{ getRiskLevelText(level) }}</span>
                    </div>
                    <div class="flex items-center">
                      <div class="w-24 bg-gray-200 rounded-full h-2 ml-3">
                        <div :class="getRiskLevelClass(level).replace('text-', 'bg-').replace('-100', '-500')" class="h-2 rounded-full" :style="{ width: (count / fraudOrders.length * 100) + '%' }"></div>
                      </div>
                      <span class="text-sm font-medium text-gray-900">{{ count }}</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <!-- نمودار تشخیص تقلب -->
            <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6 mt-6">
              <h3 class="text-lg font-semibold text-gray-900 mb-4">نمودار تشخیص تقلب</h3>
              <div class="h-64 flex items-center justify-center bg-gray-50 rounded-lg">
                <p class="text-gray-500">نمودار تشخیص تقلب در سفارشات</p>
              </div>
            </div>
          </div>

          <!-- تب تشخیص بوت -->
          <div v-if="activeTab === 'bot-detection'">
            <!-- آمار کلی تشخیص بوت -->
            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
              <!-- بوت‌های تشخیص داده شده -->
              <div class="bg-gradient-to-br from-red-50 to-red-100 border border-red-200 rounded-md p-3 hover:shadow-sm transition-all duration-200">
                <div class="flex items-center justify-between">
                  <div>
                    <p class="text-xs font-medium text-red-700 mb-0.5">بوت تشخیص داده شده</p>
                    <p class="text-lg font-bold text-red-900">۲۳۴</p>
                  </div>
                  <div class="w-8 h-8 bg-red-500 rounded-md flex items-center justify-center">
                    <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.75 17L9 20l-1 1h8l-1-1-.75-3M3 13h18M5 17h14a2 2 0 002-2V5a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"></path>
                    </svg>
                  </div>
                </div>
                <p class="text-xs text-red-600 mt-1">+۱۲% از ماه گذشته</p>
              </div>

              <!-- دقت تشخیص -->
              <div class="bg-gradient-to-br from-green-50 to-green-100 border border-green-200 rounded-md p-3 hover:shadow-sm transition-all duration-200">
                <div class="flex items-center justify-between">
                  <div>
                    <p class="text-xs font-medium text-green-700 mb-0.5">دقت تشخیص</p>
                    <p class="text-lg font-bold text-green-900">۹۶.۸%</p>
                  </div>
                  <div class="w-8 h-8 bg-green-500 rounded-md flex items-center justify-center">
                    <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                    </svg>
                  </div>
                </div>
                <p class="text-xs text-green-600 mt-1">+۲.۳% از ماه گذشته</p>
              </div>

              <!-- زمان تشخیص -->
              <div class="bg-gradient-to-br from-blue-50 to-blue-100 border border-blue-200 rounded-md p-3 hover:shadow-sm transition-all duration-200">
                <div class="flex items-center justify-between">
                  <div>
                    <p class="text-xs font-medium text-blue-700 mb-0.5">متوسط زمان تشخیص</p>
                    <p class="text-lg font-bold text-blue-900">۱.۲ ثانیه</p>
                  </div>
                  <div class="w-8 h-8 bg-blue-500 rounded-md flex items-center justify-center">
                    <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                    </svg>
                  </div>
                </div>
                <p class="text-xs text-blue-600 mt-1">-۰.۳ ثانیه از ماه گذشته</p>
              </div>

              <!-- IP های مسدود شده -->
              <div class="bg-gradient-to-br from-purple-50 to-purple-100 border border-purple-200 rounded-md p-3 hover:shadow-sm transition-all duration-200">
                <div class="flex items-center justify-between">
                  <div>
                    <p class="text-xs font-medium text-purple-700 mb-0.5">IP مسدود شده</p>
                    <p class="text-lg font-bold text-purple-900">۱۸۹</p>
                  </div>
                  <div class="w-8 h-8 bg-purple-500 rounded-md flex items-center justify-center">
                    <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 5.636l-3.536 3.536m0 5.656l3.536 3.536M9.172 9.172L5.636 5.636m3.536 9.192L5.636 18.364M12 2.25a9.75 9.75 0 109.75 9.75A9.75 9.75 0 0012 2.25z"></path>
                    </svg>
                  </div>
                </div>
                <p class="text-xs text-purple-600 mt-1">+۸% از ماه گذشته</p>
              </div>
            </div>

            <!-- نمودارها و آمار -->
            <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-6">
              <!-- نمودار انواع بوت -->
              <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
                <h3 class="text-lg font-semibold text-gray-900 mb-4">توزیع انواع بوت</h3>
                <div class="space-y-4">
                  <div class="flex items-center justify-between">
                    <div class="flex items-center">
                      <div class="w-3 h-3 bg-red-500 rounded-full ml-3"></div>
                      <span class="text-sm text-gray-700">خرید خودکار</span>
                    </div>
                    <div class="flex items-center">
                      <div class="w-24 bg-gray-200 rounded-full h-2 ml-3">
                        <div class="bg-red-500 h-2 rounded-full" style="width: 45%"></div>
                      </div>
                      <span class="text-sm font-medium text-gray-900">۴۵%</span>
                    </div>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <div class="flex items-center">
                      <div class="w-3 h-3 bg-orange-500 rounded-full ml-3"></div>
                      <span class="text-sm text-gray-700">کلیک خودکار</span>
                    </div>
                    <div class="flex items-center">
                      <div class="w-24 bg-gray-200 rounded-full h-2 ml-3">
                        <div class="bg-orange-500 h-2 rounded-full" style="width: 32%"></div>
                      </div>
                      <span class="text-sm font-medium text-gray-900">۳۲%</span>
                    </div>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <div class="flex items-center">
                      <div class="w-3 h-3 bg-yellow-500 rounded-full ml-3"></div>
                      <span class="text-sm text-gray-700">اسکرپینگ</span>
                    </div>
                    <div class="flex items-center">
                      <div class="w-24 bg-gray-200 rounded-full h-2 ml-3">
                        <div class="bg-yellow-500 h-2 rounded-full" style="width: 18%"></div>
                      </div>
                      <span class="text-sm font-medium text-gray-900">۱۸%</span>
                    </div>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <div class="flex items-center">
                      <div class="w-3 h-3 bg-blue-500 rounded-full ml-3"></div>
                      <span class="text-sm text-gray-700">سایر</span>
                    </div>
                    <div class="flex items-center">
                      <div class="w-24 bg-gray-200 rounded-full h-2 ml-3">
                        <div class="bg-blue-500 h-2 rounded-full" style="width: 5%"></div>
                      </div>
                      <span class="text-sm font-medium text-gray-900">۵%</span>
                    </div>
                  </div>
                </div>
              </div>

              <!-- آمار زمانی -->
              <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
                <h3 class="text-lg font-semibold text-gray-900 mb-4">آمار زمانی تشخیص</h3>
                <div class="space-y-4">
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-700">سریع‌ترین تشخیص</span>
                    <span class="text-sm font-medium text-green-600">۰.۳ ثانیه</span>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-700">متوسط زمان تشخیص</span>
                    <span class="text-sm font-medium text-blue-600">۱.۲ ثانیه</span>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-700">کندترین تشخیص</span>
                    <span class="text-sm font-medium text-red-600">۳.۸ ثانیه</span>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-700">تعداد تشخیص در ساعت</span>
                    <span class="text-sm font-medium text-purple-600">۲,۴۵۶</span>
                  </div>
                </div>
              </div>
            </div>

            <!-- جدول بوت‌های تشخیص داده شده -->
            <div class="bg-white rounded-lg shadow-md border border-gray-200">
              <div class="p-6 border-b border-gray-200">
                <div class="flex items-center justify-between">
                  <h3 class="text-lg font-semibold text-gray-900">بوت‌های تشخیص داده شده اخیر</h3>
                  <button class="bg-red-600 hover:bg-red-700 text-white px-4 py-2 rounded-lg text-sm font-medium transition-colors">
                    مسدود کردن همه
                  </button>
                </div>
              </div>

              <div class="overflow-x-auto">
                <table class="min-w-full divide-y divide-gray-200">
                  <thead class="bg-gray-50">
                    <tr>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">IP آدرس</th>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع فعالیت</th>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">امتیاز ریسک</th>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">زمان تشخیص</th>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
                    </tr>
                  </thead>
                  <tbody class="bg-white divide-y divide-gray-200">
                    <tr class="hover:bg-gray-50">
                      <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                        192.168.1.100
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap">
                        <span class="px-2 py-1 text-xs font-medium bg-red-100 text-red-800 rounded-full">
                          خرید خودکار
                        </span>
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap">
                        <div class="flex items-center">
                          <div class="w-16 bg-gray-200 rounded-full h-2 ml-2">
                            <div class="bg-red-600 h-2 rounded-full" style="width: 95%"></div>
                          </div>
                          <span class="text-sm font-medium text-red-600">۹۵%</span>
                        </div>
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                        ۲ دقیقه پیش
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap">
                        <span class="px-2 py-1 text-xs font-medium bg-red-100 text-red-800 rounded-full">
                          مسدود شده
                        </span>
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                        <div class="flex items-center space-x-2">
                          <button class="text-blue-600 hover:text-blue-900">جزئیات</button>
                          <button class="text-red-600 hover:text-red-900">حذف</button>
                        </div>
                      </td>
                    </tr>
                    
                    <tr class="hover:bg-gray-50">
                      <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                        10.0.0.45
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap">
                        <span class="px-2 py-1 text-xs font-medium bg-orange-100 text-orange-800 rounded-full">
                          کلیک خودکار
                        </span>
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap">
                        <div class="flex items-center">
                          <div class="w-16 bg-gray-200 rounded-full h-2 ml-2">
                            <div class="bg-orange-600 h-2 rounded-full" style="width: 87%"></div>
                          </div>
                          <span class="text-sm font-medium text-orange-600">۸۷%</span>
                        </div>
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                        ۵ دقیقه پیش
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap">
                        <span class="px-2 py-1 text-xs font-medium bg-red-100 text-red-800 rounded-full">
                          مسدود شده
                        </span>
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                        <div class="flex items-center space-x-2">
                          <button class="text-blue-600 hover:text-blue-900">جزئیات</button>
                          <button class="text-red-600 hover:text-red-900">حذف</button>
                        </div>
                      </td>
                    </tr>
                    
                    <tr class="hover:bg-gray-50">
                      <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                        172.16.0.23
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap">
                        <span class="px-2 py-1 text-xs font-medium bg-yellow-100 text-yellow-800 rounded-full">
                          اسکرپینگ
                        </span>
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap">
                        <div class="flex items-center">
                          <div class="w-16 bg-gray-200 rounded-full h-2 ml-2">
                            <div class="bg-yellow-600 h-2 rounded-full" style="width: 78%"></div>
                          </div>
                          <span class="text-sm font-medium text-yellow-600">۷۸%</span>
                        </div>
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                        ۱۲ دقیقه پیش
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap">
                        <span class="px-2 py-1 text-xs font-medium bg-red-100 text-red-800 rounded-full">
                          مسدود شده
                        </span>
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                        <div class="flex items-center space-x-2">
                          <button class="text-blue-600 hover:text-blue-900">جزئیات</button>
                          <button class="text-red-600 hover:text-red-900">حذف</button>
                        </div>
                      </td>
                    </tr>
                    
                    <tr class="hover:bg-gray-50">
                      <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                        203.0.113.67
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap">
                        <span class="px-2 py-1 text-xs font-medium bg-blue-100 text-blue-800 rounded-full">
                          رفتار مشکوک
                        </span>
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap">
                        <div class="flex items-center">
                          <div class="w-16 bg-gray-200 rounded-full h-2 ml-2">
                            <div class="bg-blue-600 h-2 rounded-full" style="width: 65%"></div>
                          </div>
                          <span class="text-sm font-medium text-blue-600">۶۵%</span>
                        </div>
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                        ۱۸ دقیقه پیش
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap">
                        <span class="px-2 py-1 text-xs font-medium bg-yellow-100 text-yellow-800 rounded-full">
                          تحت نظارت
                        </span>
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                        <div class="flex items-center space-x-2">
                          <button class="text-blue-600 hover:text-blue-900">جزئیات</button>
                          <button class="text-red-600 hover:text-red-900">مسدود</button>
                        </div>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>

              <!-- صفحه‌بندی بوت‌ها -->
              <div v-if="botData.length > 0">
                <AdminPagination 
                  :current-page="botCurrentPage"
                  :total-pages="botTotalPages"
                  :total="botData.length"
                  :items-per-page="botItemsPerPage"
                  @page-changed="handleBotPageChange"
                  @items-per-page-changed="handleBotItemsPerPageChange"
                />
              </div>
            </div>
          </div>



          <!-- تب سفارشات تقلبی -->
          <div v-if="activeTab === 'fraud-orders'">
            <!-- فیلترهای پیشرفته -->
            <div class="bg-white rounded-lg shadow-md border border-gray-200 mb-6">
              <div class="p-6 border-b border-gray-200">
                <div class="flex items-center justify-between">
                  <h3 class="text-lg font-semibold text-gray-900">فیلترهای پیشرفته</h3>
                  <button @click="resetFilters" class="text-red-600 hover:text-red-800 text-sm font-medium">
                    پاک کردن فیلترها
                  </button>
                </div>
              </div>

              <div class="p-6">
                <!-- ردیف اول فیلترها -->
                <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-4">
                  <!-- جستجو -->
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-2">جستجو</label>
                    <input 
                      v-model="searchQuery" 
                      type="text" 
                      placeholder="شماره سفارش، نام یا شماره تلفن..." 
                      class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-transparent text-sm"
                    >
                  </div>

                  <!-- نوع تقلب -->
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-2">نوع تقلب</label>
                    <select 
                      v-model="fraudTypeFilter" 
                      class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-transparent text-sm"
                    >
                      <option value="">همه انواع</option>
                      <option value="fake_address">آدرس جعلی</option>
                      <option value="stolen_card">کارت دزدیده شده</option>
                      <option value="bot_activity">فعالیت بوت</option>
                      <option value="multiple_accounts">حساب‌های متعدد</option>
                      <option value="fake_identity">هویت جعلی</option>
                    </select>
                  </div>

                  <!-- سطح ریسک -->
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-2">سطح ریسک</label>
                    <select 
                      v-model="riskLevelFilter" 
                      class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-transparent text-sm"
                    >
                      <option value="">همه سطوح</option>
                      <option value="high">ریسک بالا</option>
                      <option value="medium">ریسک متوسط</option>
                      <option value="low">ریسک پایین</option>
                    </select>
                  </div>

                  <!-- روش پرداخت -->
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-2">روش پرداخت</label>
                    <select 
                      v-model="paymentMethodFilter" 
                      class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-transparent text-sm"
                    >
                      <option value="">همه روش‌ها</option>
                      <option value="online">آنلاین</option>
                      <option value="cash">نقدی</option>
                      <option value="bank_transfer">انتقال بانکی</option>
                      <option value="wallet">کیف پول</option>
                    </select>
                  </div>
                </div>

                <!-- ردیف دوم فیلترها -->
                <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-4">
                  <!-- بازه امتیاز تقلب -->
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-2">امتیاز تقلب</label>
                    <select 
                      v-model="fraudScoreFilter" 
                      class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-transparent text-sm"
                    >
                      <option value="">همه امتیازات</option>
                      <option value="90-100">90-100% (خیلی بالا)</option>
                      <option value="80-89">80-89% (بالا)</option>
                      <option value="70-79">70-79% (متوسط)</option>
                      <option value="60-69">60-69% (پایین)</option>
                    </select>
                  </div>

                  <!-- بازه مبلغ -->
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-2">بازه مبلغ (تومان)</label>
                    <select 
                      v-model="amountRangeFilter" 
                      class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-transparent text-sm"
                    >
                      <option value="">همه مبالغ</option>
                      <option value="0-1000000">کمتر از 1 میلیون</option>
                      <option value="1000000-3000000">1 تا 3 میلیون</option>
                      <option value="3000000-5000000">3 تا 5 میلیون</option>
                      <option value="5000000+">بیش از 5 میلیون</option>
                    </select>
                  </div>

                  <!-- تاریخ از -->
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-2">از تاریخ</label>
                    <input 
                      v-model="dateFromFilter" 
                      type="date" 
                      class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-transparent text-sm"
                    >
                  </div>

                  <!-- تاریخ تا -->
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-2">تا تاریخ</label>
                    <input 
                      v-model="dateToFilter" 
                      type="date" 
                      class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-transparent text-sm"
                    >
                  </div>
                </div>

                <!-- دکمه‌های عملیات -->
                <div class="flex items-center justify-between pt-4 border-t border-gray-200">
                  <div class="flex items-center space-x-3">
                    <button @click="applyFilters" class="bg-red-600 hover:bg-red-700 text-white px-4 py-2 rounded-lg text-sm font-medium transition-colors">
                      اعمال فیلترها
                    </button>
                    <button @click="exportFilteredData" class="bg-green-600 hover:bg-green-700 text-white px-4 py-2 rounded-lg text-sm font-medium transition-colors">
                      خروجی اکسل
                    </button>
                  </div>
                  <div class="text-sm text-gray-600">
                    {{ filteredFraudOrders.length }} سفارش از {{ fraudOrders.length }} سفارش
                  </div>
                </div>
              </div>
            </div>

            <!-- جدول سفارشات تقلبی -->
            <div class="bg-white rounded-lg shadow-md border border-gray-200">
              <div class="p-6 border-b border-gray-200">
                <h3 class="text-lg font-semibold text-gray-900">لیست سفارشات تقلبی</h3>
              </div>

              <div class="overflow-x-auto">
                <table class="min-w-full divide-y divide-gray-200">
                  <thead class="bg-gray-50">
                    <tr>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">شماره سفارش</th>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مشتری</th>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مبلغ</th>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع تقلب</th>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">امتیاز تقلب</th>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">سطح ریسک</th>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ</th>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
                    </tr>
                  </thead>
                  <tbody class="bg-white divide-y divide-gray-200">
                    <tr v-for="order in paginatedFraudOrders" :key="order.id" class="hover:bg-gray-50">
                      <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                        #{{ order.orderNumber }}
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap">
                        <div>
                          <div class="text-sm font-medium text-gray-900">{{ order.customerName }}</div>
                          <div class="text-sm text-gray-500">{{ order.customerPhone }}</div>
                        </div>
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                        {{ formatPrice(order.totalAmount) }}
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap">
                        <span :class="getFraudTypeClass(order.fraudType)" class="px-2 py-1 text-xs font-medium rounded-full">
                          {{ getFraudTypeText(order.fraudType) }}
                        </span>
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap">
                        <div class="flex items-center">
                          <div class="w-16 bg-gray-200 rounded-full h-2 ml-2">
                            <div class="bg-red-600 h-2 rounded-full" :style="{ width: order.fraudScore + '%' }"></div>
                          </div>
                          <span class="text-sm font-medium text-red-600">{{ order.fraudScore }}%</span>
                        </div>
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap">
                        <span :class="getRiskLevelClass(order.riskLevel)" class="px-2 py-1 text-xs font-medium rounded-full">
                          {{ getRiskLevelText(order.riskLevel) }}
                        </span>
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                        {{ formatDate(order.createdAt) }}
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                        <div class="flex items-center space-x-2">
                          <button @click="viewFraudOrder(order)" class="text-blue-600 hover:text-blue-900">مشاهده</button>
                          <button @click="blockCustomer(order)" class="text-red-600 hover:text-red-900">مسدود</button>
                          <button @click="reportToAuthorities(order)" class="text-orange-600 hover:text-orange-900">گزارش</button>
                        </div>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>

              <!-- Empty State -->
              <div v-if="filteredFraudOrders.length === 0" class="text-center py-12">
                <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
                </svg>
                <h3 class="mt-2 text-sm font-medium text-gray-900">هیچ سفارش تقلبی یافت نشد</h3>
                <p class="mt-1 text-sm text-gray-500">فیلترها را تغییر دهید یا سفارش جدیدی بررسی کنید.</p>
              </div>

              <!-- صفحه‌بندی -->
              <div v-if="filteredFraudOrders.length > 0">
                <AdminPagination 
                  :current-page="currentPage"
                  :total-pages="totalPages"
                  :total="filteredFraudOrders.length"
                  :items-per-page="itemsPerPage"
                  @page-changed="handlePageChange"
                  @items-per-page-changed="handleItemsPerPageChange"
                />
              </div>
            </div>
          </div>

          <!-- تب گزارشات -->
          <div v-if="activeTab === 'reports'">
            <!-- نمودارهای گزارشات -->
            <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-6">
              <!-- نمودار روند تقلب‌ها -->
              <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
                <h3 class="text-lg font-semibold text-gray-900 mb-4">روند تقلب‌ها در ۶ ماه گذشته</h3>
                <div class="space-y-4">
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">ماه اول</span>
                    <div class="flex items-center">
                      <div class="w-32 bg-gray-200 rounded-full h-3 ml-3">
                        <div class="bg-red-500 h-3 rounded-full" style="width: 45%"></div>
                      </div>
                      <span class="text-sm font-medium text-gray-900">۴۵</span>
                    </div>
                  </div>
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">ماه دوم</span>
                    <div class="flex items-center">
                      <div class="w-32 bg-gray-200 rounded-full h-3 ml-3">
                        <div class="bg-red-500 h-3 rounded-full" style="width: 62%"></div>
                      </div>
                      <span class="text-sm font-medium text-gray-900">۶۲</span>
                    </div>
                  </div>
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">ماه سوم</span>
                    <div class="flex items-center">
                      <div class="w-32 bg-gray-200 rounded-full h-3 ml-3">
                        <div class="bg-red-500 h-3 rounded-full" style="width: 38%"></div>
                      </div>
                      <span class="text-sm font-medium text-gray-900">۳۸</span>
                    </div>
                  </div>
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">ماه چهارم</span>
                    <div class="flex items-center">
                      <div class="w-32 bg-gray-200 rounded-full h-3 ml-3">
                        <div class="bg-red-500 h-3 rounded-full" style="width: 71%"></div>
                      </div>
                      <span class="text-sm font-medium text-gray-900">۷۱</span>
                    </div>
                  </div>
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">ماه پنجم</span>
                    <div class="flex items-center">
                      <div class="w-32 bg-gray-200 rounded-full h-3 ml-3">
                        <div class="bg-red-500 h-3 rounded-full" style="width: 53%"></div>
                      </div>
                      <span class="text-sm font-medium text-gray-900">۵۳</span>
                    </div>
                  </div>
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">ماه ششم</span>
                    <div class="flex items-center">
                      <div class="w-32 bg-gray-200 rounded-full h-3 ml-3">
                        <div class="bg-red-500 h-3 rounded-full" style="width: 89%"></div>
                      </div>
                      <span class="text-sm font-medium text-gray-900">۸۹</span>
                    </div>
                  </div>
                </div>
              </div>

              <!-- نمودار توزیع انواع تقلب -->
              <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
                <h3 class="text-lg font-semibold text-gray-900 mb-4">توزیع انواع تقلب</h3>
                <div class="space-y-4">
                  <div class="flex items-center justify-between">
                    <div class="flex items-center">
                      <div class="w-3 h-3 bg-red-500 rounded-full ml-3"></div>
                      <span class="text-sm text-gray-700">آدرس جعلی</span>
                    </div>
                    <div class="flex items-center">
                      <div class="w-24 bg-gray-200 rounded-full h-2 ml-3">
                        <div class="bg-red-500 h-2 rounded-full" style="width: 35%"></div>
                      </div>
                      <span class="text-sm font-medium text-gray-900">۳۵%</span>
                    </div>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <div class="flex items-center">
                      <div class="w-3 h-3 bg-orange-500 rounded-full ml-3"></div>
                      <span class="text-sm text-gray-700">کارت دزدیده شده</span>
                    </div>
                    <div class="flex items-center">
                      <div class="w-24 bg-gray-200 rounded-full h-2 ml-3">
                        <div class="bg-orange-500 h-2 rounded-full" style="width: 28%"></div>
                      </div>
                      <span class="text-sm font-medium text-gray-900">۲۸%</span>
                    </div>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <div class="flex items-center">
                      <div class="w-3 h-3 bg-purple-500 rounded-full ml-3"></div>
                      <span class="text-sm text-gray-700">فعالیت بوت</span>
                    </div>
                    <div class="flex items-center">
                      <div class="w-24 bg-gray-200 rounded-full h-2 ml-3">
                        <div class="bg-purple-500 h-2 rounded-full" style="width: 22%"></div>
                      </div>
                      <span class="text-sm font-medium text-gray-900">۲۲%</span>
                    </div>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <div class="flex items-center">
                      <div class="w-3 h-3 bg-yellow-500 rounded-full ml-3"></div>
                      <span class="text-sm text-gray-700">حساب‌های متعدد</span>
                    </div>
                    <div class="flex items-center">
                      <div class="w-24 bg-gray-200 rounded-full h-2 ml-3">
                        <div class="bg-yellow-500 h-2 rounded-full" style="width: 12%"></div>
                      </div>
                      <span class="text-sm font-medium text-gray-900">۱۲%</span>
                    </div>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <div class="flex items-center">
                      <div class="w-3 h-3 bg-pink-500 rounded-full ml-3"></div>
                      <span class="text-sm text-gray-700">هویت جعلی</span>
                    </div>
                    <div class="flex items-center">
                      <div class="w-24 bg-gray-200 rounded-full h-2 ml-3">
                        <div class="bg-pink-500 h-2 rounded-full" style="width: 3%"></div>
                      </div>
                      <span class="text-sm font-medium text-gray-900">۳%</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <!-- فیلترهای گزارش -->
            <div class="bg-gradient-to-br from-gray-50 to-gray-100 rounded-xl shadow-lg border border-gray-200 p-6 mb-6">
              <div class="flex items-center justify-between mb-4">
                <div class="flex items-center">
                  <button @click="showReportFilters = !showReportFilters" class="w-8 h-8 bg-gradient-to-r from-red-500 to-pink-600 rounded-lg flex items-center justify-center ml-3 text-white hover:bg-red-600 transition-colors">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.207A1 1 0 013 6.5V4z"></path>
                    </svg>
                  </button>
                  <h3 class="text-lg font-semibold text-gray-900">فیلترهای گزارش</h3>
                </div>
                <div class="flex items-center">
                  <span class="text-sm text-gray-600 cursor-pointer hover:text-red-600 transition-colors ml-4" @click="showReportFilters = !showReportFilters">
                    {{ showReportFilters ? 'مخفی کردن' : 'نمایش' }}
                  </span>
                </div>
              </div>

              <div v-show="showReportFilters">
                <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
                  <!-- نوع گزارش -->
                  <div class="space-y-1">
                    <label class="block text-sm font-semibold text-gray-700">نوع گزارش</label>
                    <select v-model="reportType" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500 transition-all duration-200 bg-white shadow-sm">
                      <option value="comprehensive">گزارش جامع</option>
                      <option value="fraud_summary">خلاصه تقلب‌ها</option>
                      <option value="bot_activity">فعالیت بوت‌ها</option>
                      <option value="risk_analysis">تحلیل ریسک</option>
                      <option value="financial_impact">تأثیر مالی</option>
                      <option value="trend_analysis">تحلیل روند</option>
                    </select>
                  </div>

                  <!-- بازه زمانی -->
                  <div class="space-y-1">
                    <label class="block text-sm font-semibold text-gray-700">بازه زمانی</label>
                    <select v-model="reportPeriod" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500 transition-all duration-200 bg-white shadow-sm">
                      <option value="last_7_days">۷ روز گذشته</option>
                      <option value="last_30_days">۳۰ روز گذشته</option>
                      <option value="last_3_months">۳ ماه گذشته</option>
                      <option value="last_6_months">۶ ماه گذشته</option>
                      <option value="last_year">۱ سال گذشته</option>
                      <option value="custom">انتخاب دستی</option>
                    </select>
                  </div>

                  <!-- از تاریخ -->
                  <div class="space-y-1">
                    <label class="block text-sm font-semibold text-gray-700">از تاریخ</label>
                    <input v-model="reportDateFrom" type="date" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500 transition-all duration-200 bg-white shadow-sm" />
                  </div>

                  <!-- تا تاریخ -->
                  <div class="space-y-1">
                    <label class="block text-sm font-semibold text-gray-700">تا تاریخ</label>
                    <input v-model="reportDateTo" type="date" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500 transition-all duration-200 bg-white shadow-sm" />
                  </div>
                </div>

                <!-- فیلترهای اضافی -->
                <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mt-4">
                  <!-- نوع تقلب -->
                  <div class="space-y-1">
                    <label class="block text-sm font-semibold text-gray-700">نوع تقلب</label>
                    <select v-model="reportFraudType" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500 transition-all duration-200 bg-white shadow-sm">
                      <option value="">همه انواع</option>
                      <option value="fake_address">آدرس جعلی</option>
                      <option value="stolen_card">کارت دزدیده شده</option>
                      <option value="bot_activity">فعالیت بوت</option>
                      <option value="multiple_accounts">حساب‌های متعدد</option>
                      <option value="fake_identity">هویت جعلی</option>
                    </select>
                  </div>

                  <!-- سطح ریسک -->
                  <div class="space-y-1">
                    <label class="block text-sm font-semibold text-gray-700">سطح ریسک</label>
                    <select v-model="reportRiskLevel" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500 transition-all duration-200 bg-white shadow-sm">
                      <option value="">همه سطوح</option>
                      <option value="high">ریسک بالا</option>
                      <option value="medium">ریسک متوسط</option>
                      <option value="low">ریسک پایین</option>
                    </select>
                  </div>

                  <!-- فرمت خروجی -->
                  <div class="space-y-1">
                    <label class="block text-sm font-semibold text-gray-700">فرمت خروجی</label>
                    <select v-model="reportFormat" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500 transition-all duration-200 bg-white shadow-sm">
                      <option value="pdf">PDF</option>
                      <option value="excel">Excel</option>
                      <option value="csv">CSV</option>
                      <option value="json">JSON</option>
                    </select>
                  </div>

                  <!-- دکمه تولید گزارش -->
                  <div class="space-y-1">
                    <label class="block text-sm font-semibold text-gray-700 opacity-0">دکمه</label>
                    <button @click="generateReport" class="w-full bg-gradient-to-r from-red-500 to-pink-600 text-white px-3 py-2 rounded-lg font-medium hover:from-red-600 hover:to-pink-700 transition-all duration-200 flex items-center justify-center">
                      <svg class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                      </svg>
                      تولید گزارش
                    </button>
                  </div>
                </div>
              </div>
            </div>

            <!-- گزارشات سریع -->
            <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6 mb-6">
              <h3 class="text-lg font-semibold text-gray-900 mb-4">گزارشات سریع</h3>
              <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
                <div class="flex items-center justify-between p-6 bg-gradient-to-r from-blue-50 to-blue-100 rounded-lg border border-blue-200">
                  <div>
                    <h4 class="text-sm font-medium text-blue-900">گزارش ماهانه</h4>
                    <p class="text-xs text-blue-700">آخرین ۳۰ روز</p>
                  </div>
                  <div class="flex space-x-2">
                    <button @click="downloadQuickReport('monthly', 'pdf')" class="bg-blue-600 hover:bg-blue-700 text-white px-3 py-2 rounded-lg text-xs font-medium transition-colors">
                      PDF
                    </button>
                    <button @click="downloadQuickReport('monthly', 'excel')" class="bg-green-600 hover:bg-green-700 text-white px-3 py-2 rounded-lg text-xs font-medium transition-colors">
                      Excel
                    </button>
                  </div>
                </div>

                <div class="flex items-center justify-between p-6 bg-gradient-to-r from-green-50 to-green-100 rounded-lg border border-green-200">
                  <div>
                    <h4 class="text-sm font-medium text-green-900">گزارش سه‌ماهه</h4>
                    <p class="text-xs text-green-700">آخرین ۹۰ روز</p>
                  </div>
                  <div class="flex space-x-2">
                    <button @click="downloadQuickReport('quarterly', 'pdf')" class="bg-blue-600 hover:bg-blue-700 text-white px-3 py-2 rounded-lg text-xs font-medium transition-colors">
                      PDF
                    </button>
                    <button @click="downloadQuickReport('quarterly', 'excel')" class="bg-green-600 hover:bg-green-700 text-white px-3 py-2 rounded-lg text-xs font-medium transition-colors">
                      Excel
                    </button>
                  </div>
                </div>

                <div class="flex items-center justify-between p-6 bg-gradient-to-r from-purple-50 to-purple-100 rounded-lg border border-purple-200">
                  <div>
                    <h4 class="text-sm font-medium text-purple-900">گزارش سالانه</h4>
                    <p class="text-xs text-purple-700">آخرین ۳۶۵ روز</p>
                  </div>
                  <div class="flex space-x-2">
                    <button @click="downloadQuickReport('yearly', 'pdf')" class="bg-blue-600 hover:bg-blue-700 text-white px-3 py-2 rounded-lg text-xs font-medium transition-colors">
                      PDF
                    </button>
                    <button @click="downloadQuickReport('yearly', 'excel')" class="bg-green-600 hover:bg-green-700 text-white px-3 py-2 rounded-lg text-xs font-medium transition-colors">
                      Excel
                    </button>
                  </div>
                </div>
              </div>
            </div>

            <!-- تاریخچه گزارشات -->
            <div class="bg-white rounded-lg shadow-md border border-gray-200">
              <div class="p-6 border-b border-gray-200">
                <div class="flex items-center justify-between">
                  <h3 class="text-lg font-semibold text-gray-900">تاریخچه گزارشات</h3>
                  <button @click="refreshReportHistory" class="bg-gray-600 hover:bg-gray-700 text-white px-4 py-2 rounded-lg text-sm font-medium transition-colors">
                    بروزرسانی
                  </button>
                </div>
              </div>

              <div class="overflow-x-auto">
                <table class="min-w-full divide-y divide-gray-200">
                  <thead class="bg-gray-50">
                    <tr>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نام گزارش</th>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">بازه زمانی</th>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">فرمت</th>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">حجم</th>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ تولید</th>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
                    </tr>
                  </thead>
                  <tbody class="bg-white divide-y divide-gray-200">
                    <tr v-for="report in paginatedReportHistory" :key="report.id" class="hover:bg-gray-50">
                      <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                        {{ report.name }}
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap">
                        <span class="px-2 py-1 text-xs font-medium bg-blue-100 text-blue-800 rounded-full">
                          {{ report.type }}
                        </span>
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                        {{ report.period }}
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap">
                        <span class="px-2 py-1 text-xs font-medium bg-red-100 text-red-800 rounded-full">
                          {{ report.format }}
                        </span>
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                        {{ report.size }}
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                        {{ report.date }}
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                        <div class="flex items-center space-x-2">
                          <button class="text-blue-600 hover:text-blue-900">دانلود</button>
                          <button class="text-red-600 hover:text-red-900">حذف</button>
                        </div>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>

              <!-- صفحه‌بندی تاریخچه گزارشات -->
              <div v-if="reportHistory.length > 0">
                <AdminPagination 
                  :current-page="reportCurrentPage"
                  :total-pages="reportTotalPages"
                  :total="reportHistory.length"
                  :items-per-page="reportItemsPerPage"
                  @page-changed="handleReportPageChange"
                  @items-per-page-changed="handleReportItemsPerPageChange"
                />
              </div>
            </div>
          </div>

          <!-- تب تنظیمات -->
          <div v-if="activeTab === 'settings'">
            <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
              <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات سیستم تشخیص تقلب</h3>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <!-- تنظیمات سرویس GeoIP/IP-Intel -->
                <div>
                  <h4 class="font-medium text-gray-800 mb-2">سرویس GeoIP / IP-Intel</h4>
                  <div class="space-y-3">
                    <div>
                      <label class="block text-sm text-gray-700 mb-1">Provider URL</label>
                      <input v-model="settings.ipintel.url" type="text" class="w-full px-3 py-2 border rounded" placeholder="https://provider.example.com/lookup" />
                    </div>
                    <div>
                      <label class="block text-sm text-gray-700 mb-1">Provider Token</label>
                      <input v-model="settings.ipintel.token" type="password" class="w-full px-3 py-2 border rounded" placeholder="توکن سرویس" />
                    </div>
                    <div class="flex items-center gap-2">
                      <input id="ipintel-cache" v-model="settings.ipintel.cacheEnabled" type="checkbox" class="border rounded" />
                      <label for="ipintel-cache" class="text-sm text-gray-700">فعال‌سازی کش نتایج</label>
                    </div>
                  </div>
                </div>

                <!-- تنظیمات Rate Limit -->
                <div>
                  <h4 class="font-medium text-gray-800 mb-2">Rate Limit</h4>
                  <div class="grid grid-cols-1 md:grid-cols-2 gap-3">
                    <div>
                      <label class="block text-sm text-gray-700 mb-1">ادمین (req/min)</label>
                      <input v-model.number="settings.ratelimit.adminRPM" type="number" min="1" class="w-full px-3 py-2 border rounded" />
                    </div>
                    <div>
                      <label class="block text-sm text-gray-700 mb-1">ارزیابی (req/min)</label>
                      <input v-model.number="settings.ratelimit.evaluateRPM" type="number" min="1" class="w-full px-3 py-2 border rounded" />
                    </div>
                  </div>
                </div>

                <!-- تنظیمات امتیازدهی و قوانین -->
                <div class="md:col-span-2">
                  <h4 class="font-medium text-gray-800 mb-2">قوانین و وزن‌ها</h4>
                  <div class="overflow-x-auto border rounded">
                    <table class="min-w-full divide-y divide-gray-200">
                      <thead class="bg-gray-50">
                        <tr>
                          <th class="px-4 py-2 text-right text-xs font-medium text-gray-500">کلید قانون</th>
                          <th class="px-4 py-2 text-right text-xs font-medium text-gray-500">توضیح</th>
                          <th class="px-4 py-2 text-right text-xs font-medium text-gray-500">وزن</th>
                          <th class="px-4 py-2 text-right text-xs font-medium text-gray-500">فعال</th>
                          <th class="px-4 py-2 text-right text-xs font-medium text-gray-500">پارامترها (JSON)</th>
                        </tr>
                      </thead>
                      <tbody class="bg-white divide-y divide-gray-200">
                        <tr v-for="(r, idx) in settings.rules" :key="r.key">
                          <td class="px-4 py-2 text-sm">{{ r.key }}</td>
                          <td class="px-4 py-2 text-sm">{{ r.description }}</td>
                          <td class="px-4 py-2 text-sm"><input v-model.number="r.weight" type="number" class="w-24 px-2 py-1 border rounded" /></td>
                          <td class="px-4 py-2 text-sm"><input v-model="r.enabled" type="checkbox" /></td>
                          <td class="px-4 py-2 text-sm">
                            <textarea v-model="r.paramsJson" class="w-full px-2 py-1 border rounded" rows="2" placeholder="{\n  &quot;threshold&quot;: 3\n}"></textarea>
                          </td>
                        </tr>
                      </tbody>
                    </table>
                  </div>
                </div>

                <!-- لیست‌ها -->
                <div class="md:col-span-2">
                  <h4 class="font-medium text-gray-800 mb-2">لیست سفید / سیاه</h4>
                  <div class="grid grid-cols-1 md:grid-cols-3 gap-3">
                    <div>
                      <label class="block text-sm text-gray-700 mb-1">نوع</label>
                      <select v-model="listForm.kind" class="w-full px-3 py-2 border rounded">
                        <option value="ip">IP</option>
                        <option value="email">Email</option>
                        <option value="card">Card</option>
                        <option value="device">Device</option>
                      </select>
                    </div>
                    <div>
                      <label class="block text-sm text-gray-700 mb-1">مقدار</label>
                      <input v-model="listForm.value" type="text" class="w-full px-3 py-2 border rounded" placeholder="مانند IP یا ایمیل" />
                    </div>
                    <div class="flex items-end">
                      <div class="flex gap-2">
                        <button @click="addToWhitelist" class="px-3 py-2 bg-green-600 text-white rounded text-sm">افزودن به سفید</button>
                        <button @click="addToBlacklist" class="px-3 py-2 bg-red-600 text-white rounded text-sm">افزودن به سیاه</button>
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              <div class="mt-6 flex gap-2">
                <button @click="saveSettings" class="px-4 py-2 bg-blue-600 text-white rounded">ذخیره تنظیمات (Placeholder)</button>
                <button @click="reEvaluateOpen" class="px-4 py-2 bg-gray-100 rounded">بازارزیابی پرونده‌های باز</button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </ClientOnly>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
declare const useAuth: () => { user: { id?: number; name?: string; email?: string } | null; hasPermission: (perm: string) => boolean }
declare const $fetch: <T = unknown>(url: string, options?: { method?: string; body?: unknown; query?: Record<string, string | number> }) => Promise<T>
</script>

<script setup lang="ts">
import AdminPagination from '@/components/admin/common/Pagination.vue';
import { computed, onMounted, reactive, ref, watch } from 'vue';

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user, hasPermission } = useAuth()

// متغیرهای صفحه‌بندی سفارشات تقلبی (باید در ابتدا تعریف شوند)
const currentPage = ref(1)
const itemsPerPage = ref(5)

// فیلترهای سفارشات تقلبی (باید در ابتدا تعریف شوند)
const fraudTypeFilter = ref('')
const riskLevelFilter = ref('')
const searchQuery = ref('')
const paymentMethodFilter = ref('')
const fraudScoreFilter = ref('')
const minAmountFilter = ref('')
const maxAmountFilter = ref('')
const dateFromFilter = ref('')
const dateToFilter = ref('')
const amountRangeFilter = ref('')
const showAdvanced = ref(false)

// تب‌های صفحه
const tabs = ref([
  { id: 'dashboard', name: 'داشبورد' },
  { id: 'bot-detection', name: 'تشخیص بوت' },
  { id: 'reports', name: 'گزارشات' },
  { id: 'settings', name: 'تنظیمات' }
])

const activeTab = ref<'dashboard' | 'bot-detection' | 'reports' | 'settings' | 'fraud-orders'>('dashboard')

// آمار تشخیص تقلب
const stats = ref({
  suspiciousOrders: 156,
  confirmedFraud: 89,
  accuracy: 92.3,
  blockedIPs: 234
})

// داده‌های بوت‌های تشخیص داده شده
const botData = ref([
  {
    id: 1,
    ipAddress: '192.168.1.100',
    activityType: 'خرید خودکار',
    riskScore: 95,
    detectedTime: '۲ دقیقه پیش',
    status: 'مسدود شده'
  },
  {
    id: 2,
    ipAddress: '10.0.0.45',
    activityType: 'کلیک خودکار',
    riskScore: 87,
    detectedTime: '۵ دقیقه پیش',
    status: 'مسدود شده'
  },
  {
    id: 3,
    ipAddress: '172.16.0.23',
    activityType: 'اسکرپینگ',
    riskScore: 78,
    detectedTime: '۱۲ دقیقه پیش',
    status: 'مسدود شده'
  },
  {
    id: 4,
    ipAddress: '203.0.113.67',
    activityType: 'رفتار مشکوک',
    riskScore: 65,
    detectedTime: '۱۸ دقیقه پیش',
    status: 'تحت نظارت'
  },
  {
    id: 5,
    ipAddress: '198.51.100.34',
    activityType: 'حمله DDoS',
    riskScore: 92,
    detectedTime: '۲۵ دقیقه پیش',
    status: 'مسدود شده'
  },
  {
    id: 6,
    ipAddress: '203.0.113.89',
    activityType: 'اسپم',
    riskScore: 73,
    detectedTime: '۳۰ دقیقه پیش',
    status: 'مسدود شده'
  }
])

// داده‌های تاریخچه گزارشات
const reportHistory = ref([
  {
    id: 1,
    name: 'گزارش جامع تشخیص تقلب',
    type: 'جامع',
    period: '۳۰ روز گذشته',
    format: 'PDF',
    size: '۲.۴ MB',
    date: '۲ ساعت پیش'
  },
  {
    id: 2,
    name: 'تحلیل ریسک سه‌ماهه',
    type: 'تحلیل ریسک',
    period: '۹۰ روز گذشته',
    format: 'Excel',
    size: '۱.۸ MB',
    date: '۱ روز پیش'
  },
  {
    id: 3,
    name: 'گزارش فعالیت بوت‌ها',
    type: 'فعالیت بوت',
    period: '۷ روز گذشته',
    format: 'CSV',
    size: '۰.۹ MB',
    date: '۲ روز پیش'
  },
  {
    id: 4,
    name: 'گزارش ماهانه تقلب',
    type: 'ماهانه',
    period: '۳۰ روز گذشته',
    format: 'PDF',
    size: '۳.۲ MB',
    date: '۳ روز پیش'
  },
  {
    id: 5,
    name: 'تحلیل الگوهای مشکوک',
    type: 'تحلیل الگو',
    period: '۱۴ روز گذشته',
    format: 'JSON',
    size: '۱.۵ MB',
    date: '۴ روز پیش'
  },
  {
    id: 6,
    name: 'گزارش IP های مسدود شده',
    type: 'IP مسدود',
    period: '۷ روز گذشته',
    format: 'Excel',
    size: '۰.۷ MB',
    date: '۵ روز پیش'
  }
])

// داده سفارشات از API
const fraudOrders = ref([])
const totalCount = ref(0)
const isLoading = ref(false)

const fetchFraudCases = async () => {
  try {
    isLoading.value = true
    const res: any = await $fetch('/api/admin/transactions/fraud-detection', {
      query: {
        risk: riskLevelFilter.value,
        status: '',
        page: currentPage.value,
        pageSize: itemsPerPage.value
      }
    })
    totalCount.value = Number(res?.total || 0)
    fraudOrders.value = (res?.data || []).map((row: any) => ({
      id: row.id,
      orderNumber: row.orderNumber || '-',
      customerName: row.customerName || '-',
      customerPhone: row.customerPhone || '-',
      totalAmount: row.totalAmount || 0,
      paymentMethod: row.paymentMethod || '-',
      status: row.status || '-',
      createdAt: row.createdAt,
      fraudType: '',
      fraudScore: row.fraudScore,
      riskLevel: row.riskLevel,
      blockStatus: row.status === 'blocked' ? 'system' : row.status,
      blockedBy: null,
      items: []
    }))
  } finally {
    isLoading.value = false
  }
}

onMounted(fetchFraudCases)

// واکنش به تغییر صفحه و اندازه صفحه
watch(currentPage, () => { fetchFraudCases() })
watch(itemsPerPage, () => { currentPage.value = 1; fetchFraudCases() })
// واکنش به تغییر سطح ریسک (فیلتر سمت سرور)
watch(riskLevelFilter, () => { currentPage.value = 1; fetchFraudCases() })

// فیلترهای سفارشات تقلبی (منتقل به ابتدای فایل)

// متغیرهای صفحه‌بندی سفارشات تقلبی (منتقل به ابتدای فایل)

// محاسبه تعداد صفحات بر اساس total سمت سرور
const totalPages = computed(() => {
  return Math.max(1, Math.ceil((totalCount.value || 0) / itemsPerPage.value))
})

// سفارشات صفحه‌بندی شده
const paginatedFraudOrders = computed(() => {
  // paginatedFraudOrders - currentPage
  // paginatedFraudOrders - itemsPerPage
  // paginatedFraudOrders - filteredFraudOrders
  const start = (currentPage.value - 1) * itemsPerPage.value
  const end = start + itemsPerPage.value
  const result = filteredFraudOrders.value.slice(start, end)
  // paginatedFraudOrders - result
  return result
})

// متد تغییر صفحه سفارشات تقلبی
const handlePageChange = (page) => {
  currentPage.value = page
}

// متد تغییر تعداد آیتم در هر صفحه سفارشات تقلبی
const handleItemsPerPageChange = (newItemsPerPage) => {
  itemsPerPage.value = newItemsPerPage
  currentPage.value = 1 // برگشت به صفحه اول
}

// متغیرهای صفحه‌بندی بوت‌ها
const botCurrentPage = ref(1)
const botItemsPerPage = ref(5)

// محاسبه تعداد صفحات بوت‌ها
const botTotalPages = computed(() => {
  return Math.ceil(botData.value.length / botItemsPerPage.value)
})

// بوت‌های صفحه‌بندی شده
const paginatedBotData = computed(() => {
  const start = (botCurrentPage.value - 1) * botItemsPerPage.value
  const end = start + botItemsPerPage.value
  return botData.value.slice(start, end)
})

// متد تغییر صفحه بوت‌ها
const handleBotPageChange = (page) => {
  botCurrentPage.value = page
}

// وضعیت تنظیمات (Placeholder – بعداً به API متصل می‌شود)
const settings = reactive({
  ipintel: { url: '', token: '', cacheEnabled: true },
  ratelimit: { adminRPM: 60, evaluateRPM: 60 },
  rules: [
    { key: 'geoip_mismatch_city', description: 'عدم تطابق شهر GeoIP با شهر ارسال', weight: 30, enabled: true, paramsJson: '{"threshold":0}' },
    { key: 'new_account_high_amount', description: 'حساب جدید با مبلغ بالا', weight: 25, enabled: true, paramsJson: '{}' },
    { key: 'rapid_failed_payments', description: 'خطای مکرر پرداخت از یک IP', weight: 20, enabled: true, paramsJson: '{"threshold":3}' },
    { key: 'datacenter_or_proxy_ip', description: 'IP دیتاسنتر یا پروکسی', weight: 40, enabled: true, paramsJson: '{}' },
    { key: 'whitelist_hit', description: 'مطابقت با لیست سفید', weight: -40, enabled: true, paramsJson: '{}' }
  ]
})

const listForm = reactive({ kind: 'ip', value: '' })

const saveSettings = async () => {
  // Placeholder: بعداً به API متصل می‌شود
  // save settings
}

const reEvaluateOpen = async () => {
  try {
    await $fetch('/api/admin/fraud/re-evaluate-open', { method: 'POST' })
  } catch (e) {
    // re-evaluate failed
  }
}

const addToWhitelist = async () => {
  try {
    await $fetch('/api/admin/fraud/list/whitelist', { method: 'POST', body: { kind: listForm.kind, value: listForm.value } })
  } catch (e) {
    // whitelist failed
  }
}

const addToBlacklist = async () => {
  try {
    await $fetch('/api/admin/fraud/list/blacklist', { method: 'POST', body: { kind: listForm.kind, value: listForm.value } })
  } catch (e) {
    // blacklist failed
  }
}

// متد تغییر تعداد آیتم در هر صفحه بوت‌ها
const handleBotItemsPerPageChange = (newItemsPerPage) => {
  botItemsPerPage.value = newItemsPerPage
  botCurrentPage.value = 1
}

// متغیرهای صفحه‌بندی گزارشات
const reportCurrentPage = ref(1)
const reportItemsPerPage = ref(5)

// محاسبه تعداد صفحات گزارشات
const reportTotalPages = computed(() => {
  return Math.ceil(reportHistory.value.length / reportItemsPerPage.value)
})

// گزارشات صفحه‌بندی شده
const paginatedReportHistory = computed(() => {
  const start = (reportCurrentPage.value - 1) * reportItemsPerPage.value
  const end = start + reportItemsPerPage.value
  return reportHistory.value.slice(start, end)
})

// متد تغییر صفحه گزارشات
const handleReportPageChange = (page) => {
  reportCurrentPage.value = page
}

// متد تغییر تعداد آیتم در هر صفحه گزارشات
const handleReportItemsPerPageChange = (newItemsPerPage) => {
  reportItemsPerPage.value = newItemsPerPage
  reportCurrentPage.value = 1
}

// فیلتر کردن سفارشات تقلبی
const filteredFraudOrders = computed(() => {
  // fraudOrders.value
  let filtered = fraudOrders.value

  // فیلتر جستجو
  if (searchQuery.value) {
    filtered = filtered.filter(order => 
      order.orderNumber.includes(searchQuery.value) ||
      order.customerName.includes(searchQuery.value) ||
      order.customerPhone.includes(searchQuery.value)
    )
  }

  // فیلتر نوع تقلب
  if (fraudTypeFilter.value) {
    filtered = filtered.filter(order => order.fraudType === fraudTypeFilter.value)
  }

  // فیلتر سطح ریسک
  if (riskLevelFilter.value) {
    filtered = filtered.filter(order => order.riskLevel === riskLevelFilter.value)
  }

  // فیلتر روش پرداخت
  if (paymentMethodFilter.value) {
    filtered = filtered.filter(order => order.paymentMethod === paymentMethodFilter.value)
  }

  // فیلتر امتیاز تقلب
  if (fraudScoreFilter.value) {
    const [min, max] = fraudScoreFilter.value.split('-').map(Number)
    if (fraudScoreFilter.value === '5000000+') {
      filtered = filtered.filter(order => order.fraudScore >= 90)
    } else {
      filtered = filtered.filter(order => order.fraudScore >= min && order.fraudScore <= max)
    }
  }

  // فیلتر حداقل مبلغ
  if (minAmountFilter.value) {
    filtered = filtered.filter(order => order.totalAmount >= Number(minAmountFilter.value))
  }

  // فیلتر حداکثر مبلغ
  if (maxAmountFilter.value) {
    filtered = filtered.filter(order => order.totalAmount <= Number(maxAmountFilter.value))
  }

  // فیلتر تاریخ از
  if (dateFromFilter.value) {
    filtered = filtered.filter(order => new Date(order.createdAt) >= new Date(dateFromFilter.value))
  }

  // فیلتر تاریخ تا
  if (dateToFilter.value) {
    filtered = filtered.filter(order => new Date(order.createdAt) <= new Date(dateToFilter.value))
  }

  return filtered
})

// آمار انواع تقلب
const fraudTypeStats = computed(() => {
  const stats = {}
  fraudOrders.value.forEach(order => {
    stats[order.fraudType] = (stats[order.fraudType] || 0) + 1
  })
  return stats
})

// آمار سطح ریسک
const riskLevelStats = computed(() => {
  const stats = {}
  fraudOrders.value.forEach(order => {
    stats[order.riskLevel] = (stats[order.riskLevel] || 0) + 1
  })
  return stats
})

// متغیرهای گزارشات
const showReportFilters = ref(true)
const reportType = ref('comprehensive')
const reportPeriod = ref('last_30_days')
const reportDateFrom = ref('')
const reportDateTo = ref('')
const reportFraudType = ref('')
const reportRiskLevel = ref('')
const reportFormat = ref('pdf')

// متدهای گزارشات
const generateReport = () => {
  // تولید گزارش با تنظیمات
  const reportData = {
    type: reportType.value,
    period: reportPeriod.value,
    dateFrom: reportDateFrom.value,
    dateTo: reportDateTo.value,
    fraudType: reportFraudType.value,
    riskLevel: reportRiskLevel.value,
    format: reportFormat.value
  }
  
  // اینجا می‌توانید منطق تولید گزارش را اضافه کنید
  console.log('Generating report with data:', reportData)
}

const downloadQuickReport = (period, format) => {
  // دانلود گزارش سریع
}

const refreshReportHistory = () => {
  // بروزرسانی تاریخچه گزارشات
}

// متدهای مدیریت
const exportFraudReport = () => {
  // خروجی گزارش تشخیص تقلب
}

const refreshData = () => {
  // بروزرسانی داده‌های تشخیص تقلب
}

const showSettings = () => {
  // نمایش تنظیمات تشخیص تقلب
}

// متدهای کمکی برای سفارشات تقلبی
const getFraudTypeText = (type) => {
  const types = {
    'fake_address': 'آدرس جعلی',
    'stolen_card': 'کارت دزدیده شده',
    'bot_activity': 'فعالیت بوت',
    'multiple_accounts': 'حساب‌های متعدد',
    'fake_identity': 'هویت جعلی'
  }
  return types[type] || type
}

const getFraudTypeClass = (type) => {
  const classes = {
    'fake_address': 'bg-red-100 text-red-800',
    'stolen_card': 'bg-orange-100 text-orange-800',
    'bot_activity': 'bg-purple-100 text-purple-800',
    'multiple_accounts': 'bg-yellow-100 text-yellow-800',
    'fake_identity': 'bg-pink-100 text-pink-800'
  }
  return classes[type] || 'bg-gray-100 text-gray-800'
}

const getRiskLevelClass = (level) => {
  const classes = {
    'high': 'bg-red-100 text-red-800',
    'medium': 'bg-yellow-100 text-yellow-800',
    'low': 'bg-green-100 text-green-800'
  }
  return classes[level] || 'bg-gray-100 text-gray-800'
}

const getRiskLevelText = (level) => {
  const levels = {
    'high': 'بالا',
    'medium': 'متوسط',
    'low': 'پایین'
  }
  return levels[level] || level
}

const getBlockStatusText = (status) => {
  const statuses = {
    'system': 'سیستمی',
    'manual': 'دستی',
    'pending': 'در انتظار',
    'none': 'بلاک نشده'
  }
  return statuses[status] || status
}

const getBlockStatusClass = (status) => {
  const classes = {
    'system': 'bg-red-100 text-red-800',
    'manual': 'bg-blue-100 text-blue-800',
    'pending': 'bg-yellow-100 text-yellow-800',
    'none': 'bg-gray-100 text-gray-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}

const formatPrice = (price) => {
  return new Intl.NumberFormat('fa-IR').format(price) + ' تومان'
}

const formatDate = (date) => {
  return new Date(date).toLocaleDateString('fa-IR')
}

type FraudCaseDetail = {
  data?: unknown
  [key: string]: unknown
}

const viewFraudOrder = async (order: { id: number | string }) => {
  try {
    const detail = await $fetch<FraudCaseDetail>(`/api/admin/fraud/cases/${order.id}`)
    // جزئیات پرونده تقلب
    const detailData = (detail as { data?: unknown }).data || detail
    alert('جزئیات پرونده:\n' + JSON.stringify(detailData, null, 2))
  } catch (e) {
    // load case detail failed
    console.error('Error loading fraud case detail:', e)
  }
}

const blockCustomer = async (order) => {
  try {
    await $fetch(`/api/admin/transactions/${order.id}/fraud-decision`, {
      method: 'POST',
      body: { action: 'block', note: '' }
    })
    await fetchFraudCases()
  } catch (e) {
    // block failed
  }
}

const approveCustomer = async (order) => {
  try {
    await $fetch(`/api/admin/transactions/${order.id}/fraud-decision`, {
      method: 'POST',
      body: { action: 'approve', note: '' }
    })
    await fetchFraudCases()
  } catch (e) {
    // approve failed
  }
}

const markReview = async (order) => {
  try {
    await $fetch(`/api/admin/transactions/${order.id}/fraud-decision`, {
      method: 'POST',
      body: { action: 'review', note: '' }
    })
    await fetchFraudCases()
  } catch (e) {
    // review failed
  }
}

const reportToAuthorities = (order) => {
  // گزارش به مقامات
}

const openRules = async () => {
  try {
    const rules: any = await $fetch('/api/admin/fraud/rules')
    console.table(rules?.data || [])
  } catch (e) {
    // load rules failed
  }
}

// متدهای فیلتر
const resetFilters = () => {
  searchQuery.value = ''
  fraudTypeFilter.value = ''
  riskLevelFilter.value = ''
  paymentMethodFilter.value = ''
  fraudScoreFilter.value = ''
      minAmountFilter.value = ''
    maxAmountFilter.value = ''
  dateFromFilter.value = ''
  dateToFilter.value = ''
}

const applyFilters = () => {
  // فیلترها به صورت reactive اعمال می‌شوند
  // فیلترها اعمال شدند
}

const exportFilteredData = () => {
  // خروجی اکسل از داده‌های فیلتر شده
}

// computed برای فیلترهای فعال
const activeFilters = computed(() => {
  const active: Record<string, string> = {}
  if (searchQuery.value) active.search = searchQuery.value
  if (fraudTypeFilter.value) active.fraudType = fraudTypeFilter.value
  if (riskLevelFilter.value) active.riskLevel = riskLevelFilter.value
  if (paymentMethodFilter.value) active.paymentMethod = paymentMethodFilter.value
  if (fraudScoreFilter.value) active.fraudScore = fraudScoreFilter.value
  if (minAmountFilter.value) active.minAmount = minAmountFilter.value
  if (maxAmountFilter.value) active.maxAmount = maxAmountFilter.value
  if (dateFromFilter.value) active.dateFrom = dateFromFilter.value
  if (dateToFilter.value) active.dateTo = dateToFilter.value
  return active
})

const hasActiveFilters = computed(() => {
  return Object.keys(activeFilters.value).length > 0
})

const getFilterLabel = (key, value) => {
  const labels = {
    search: `جستجو: ${value}`,
    fraudType: `نوع تقلب: ${getFraudTypeText(value)}`,
    riskLevel: `سطح ریسک: ${getRiskLevelText(value)}`,
    paymentMethod: `روش پرداخت: ${getPaymentMethodText(value)}`,
    fraudScore: `امتیاز تقلب: ${value}`,
    minAmount: `حداقل مبلغ: ${value} تومان`,
    maxAmount: `حداکثر مبلغ: ${value} تومان`,
    dateFrom: `از تاریخ: ${value}`,
    dateTo: `تا تاریخ: ${value}`
  }
  return labels[key] || `${key}: ${value}`
}

const getPaymentMethodText = (method) => {
  const methodMap = {
    'online': 'آنلاین',
    'cash': 'نقدی',
    'bank_transfer': 'انتقال بانکی',
    'wallet': 'کیف پول'
  }
  return methodMap[method] || method
}

const removeFilter = (key) => {
  switch (key) {
    case 'search':
      searchQuery.value = ''
      break
    case 'fraudType':
      fraudTypeFilter.value = ''
      break
    case 'riskLevel':
      riskLevelFilter.value = ''
      break
    case 'paymentMethod':
      paymentMethodFilter.value = ''
      break
    case 'fraudScore':
      fraudScoreFilter.value = ''
      break
    case 'minAmount':
      minAmountFilter.value = ''
      break
    case 'maxAmount':
      maxAmountFilter.value = ''
      break
    case 'dateFrom':
      dateFromFilter.value = ''
      break
    case 'dateTo':
      dateToFilter.value = ''
      break
  }
}
</script> 

