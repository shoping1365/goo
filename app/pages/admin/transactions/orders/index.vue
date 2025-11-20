<template>
  <ClientOnly>
    <!-- Page Header -->
    <div class="bg-white shadow-sm border-b border-gray-200 mb-4">
      <div class="w-full px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center py-3">
          <div>
            <h1 class="text-lg font-bold text-gray-900">مدیریت سفارشات</h1>
            <p class="text-xs text-gray-600 mt-1">مدیریت کامل سفارشات و پیگیری وضعیت</p>
          </div>
          <div class="flex space-x-2 space-x-reverse">
            <button 
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-blue-400 to-blue-600 hover:from-blue-500 hover:to-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105"
              @click="showCreateOrderModal = true"
            >
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5 ml-2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M12 4v16m8-8H4" />
              </svg>
              سفارش دستی
            </button>
            <button 
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-green-400 to-green-600 hover:from-green-500 hover:to-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105"
              @click="exportOrders"
            >
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
              </svg>
              خروجی Excel
            </button>
            <button 
              class="inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-lg text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 shadow-md transition-all duration-200"
              @click="printOrders"
            >
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 17h2a2 2 0 002-2v-4a2 2 0 00-2-2H5a2 2 0 00-2 2v4a2 2 0 002 2h2m2 4h6a2 2 0 002-2v-4a2 2 0 00-2-2H9a2 2 0 00-2 2v4a2 2 0 002 2zm8-12V5a2 2 0 00-2-2H9a2 2 0 00-2 2v4h10z"></path>
              </svg>
              چاپ
            </button>
            <NuxtLink 
              to="/admin/order-surveys2"
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-purple-400 to-purple-600 hover:from-purple-500 hover:to-purple-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-purple-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105"
            >
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
              نظرسنجی 2
            </NuxtLink>
          </div>
        </div>
      </div>
    </div>

    <div class="w-full px-4 py-4">
      <!-- پیام موفقیت -->
      <div v-if="showSuccessMessage" class="mb-4 bg-green-50 border border-green-200 rounded-md p-6">
        <div class="flex">
          <div class="flex-shrink-0">
            <svg class="h-5 w-5 text-green-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
          <div class="mr-3">
            <p class="text-sm text-green-800">{{ successMessage }}</p>
          </div>
          <div class="mr-auto pr-3">
            <button class="text-green-400 hover:text-green-600" @click="showSuccessMessage = false">
              <svg class="h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
              </svg>
            </button>
          </div>
        </div>
      </div>
      
      <!-- تب‌های اصلی -->
      <div class="bg-white border-b border-gray-200 mb-6">
        <nav class="-mb-px flex px-6 overflow-x-auto">
            <button
              v-for="tab in tabs"
              :key="tab.id"
              :class="[
                activeTab === tab.id
                  ? 'border-blue-500 text-blue-600'
                  : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300',
              'flex-1 whitespace-nowrap py-4 px-4 border-b-2 font-medium text-sm text-center'
              ]"
              @click="activeTab = tab.id"
            >
              {{ tab.name }}
            </button>
          </nav>
        </div>

        <div class="p-6">
          <!-- تب لیست سفارشات -->
          <div v-if="activeTab === 'list'">
            <!-- آمار کلی سفارشات -->
          <div class="rounded-lg shadow-md border border-gray-200 overflow-hidden mb-6">
              <div class="bg-gradient-to-r from-blue-50 to-indigo-50 px-4 py-3 border-b border-gray-200">
                <div class="flex items-center">
                  <div class="w-6 h-6 bg-blue-500 rounded-md flex items-center justify-center ml-2">
                    <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"></path>
                    </svg>
                  </div>
                  <h3 class="text-sm font-semibold text-gray-900">آمار و اطلاعات کلی سفارشات</h3>
                </div>
              </div>
              
              <div class="p-6">
                <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-6 gap-3">
                  <!-- کل سفارشات -->
                  <div class="bg-gradient-to-br from-blue-50 to-blue-100 border border-blue-200 rounded-md p-3 hover:shadow-sm transition-all duration-200">
                    <div class="flex items-center justify-between">
                      <div>
                        <p class="text-xs font-medium text-blue-700 mb-0.5">کل سفارشات</p>
                        <p class="text-lg font-bold text-blue-900">{{ stats.total.toLocaleString('en-US') }}</p>
                      </div>
                      <div class="w-8 h-8 bg-blue-500 rounded-md flex items-center justify-center">
                        <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"></path>
                        </svg>
                      </div>
                    </div>
                  </div>

                  <!-- در انتظار تایید -->
                  <div class="bg-gradient-to-br from-yellow-50 to-yellow-100 border border-yellow-200 rounded-md p-3 hover:shadow-sm transition-all duration-200">
                    <div class="flex items-center justify-between">
                      <div>
                        <p class="text-xs font-medium text-yellow-700 mb-0.5">در انتظار تایید</p>
                        <p class="text-lg font-bold text-yellow-900">{{ stats.pending.toLocaleString('en-US') }}</p>
                      </div>
                      <div class="w-8 h-8 bg-yellow-500 rounded-md flex items-center justify-center">
                        <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                        </svg>
                      </div>
                    </div>
                  </div>

                  <!-- در حال آماده‌سازی -->
                  <div class="bg-gradient-to-br from-purple-50 to-purple-100 border border-purple-200 rounded-md p-3 hover:shadow-sm transition-all duration-200">
                    <div class="flex items-center justify-between">
                      <div>
                        <p class="text-xs font-medium text-purple-700 mb-0.5">در حال آماده‌سازی</p>
                        <p class="text-lg font-bold text-purple-900">{{ stats.processing.toLocaleString('en-US') }}</p>
                      </div>
                      <div class="w-8 h-8 bg-purple-500 rounded-md flex items-center justify-center">
                        <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
                        </svg>
                      </div>
                    </div>
                  </div>

                  <!-- ارسال شده -->
                  <div class="bg-gradient-to-br from-blue-50 to-blue-100 border border-blue-200 rounded-md p-3 hover:shadow-sm transition-all duration-200">
                    <div class="flex items-center justify-between">
                      <div>
                        <p class="text-xs font-medium text-blue-700 mb-0.5">ارسال شده</p>
                        <p class="text-lg font-bold text-blue-900">{{ stats.shipped.toLocaleString('en-US') }}</p>
                      </div>
                      <div class="w-8 h-8 bg-blue-500 rounded-md flex items-center justify-center">
                        <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4"></path>
                        </svg>
                      </div>
                    </div>
                  </div>

                  <!-- تحویل داده شده -->
                  <div class="bg-gradient-to-br from-green-50 to-green-100 border border-green-200 rounded-md p-3 hover:shadow-sm transition-all duration-200">
                    <div class="flex items-center justify-between">
                      <div>
                        <p class="text-xs font-medium text-green-700 mb-0.5">تحویل شده</p>
                        <p class="text-lg font-bold text-green-900">{{ stats.delivered.toLocaleString('en-US') }}</p>
                      </div>
                      <div class="w-8 h-8 bg-green-500 rounded-md flex items-center justify-center">
                        <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                        </svg>
                      </div>
                    </div>
                  </div>

                  <!-- لغو شده -->
                  <div class="bg-gradient-to-br from-red-50 to-red-100 border border-red-200 rounded-md p-3 hover:shadow-sm transition-all duration-200">
                    <div class="flex items-center justify-between">
                      <div>
                        <p class="text-xs font-medium text-red-700 mb-0.5">لغو شده</p>
                        <p class="text-lg font-bold text-red-900">{{ stats.cancelled.toLocaleString('en-US') }}</p>
                      </div>
                      <div class="w-8 h-8 bg-red-500 rounded-md flex items-center justify-center">
                        <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                        </svg>
                      </div>
                    </div>
                  </div>
                </div>

                <!-- آمار مالی -->
                <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-3 mt-4 pt-4 border-t border-gray-200">
                  <!-- کل فروش امروز -->
                  <div class="bg-gradient-to-br from-emerald-50 to-emerald-100 border border-emerald-200 rounded-md p-3 hover:shadow-sm transition-all duration-200">
                    <div class="flex items-center justify-between">
                      <div>
                        <p class="text-xs font-medium text-emerald-700 mb-0.5">فروش امروز</p>
                        <p class="text-lg font-bold text-emerald-900">{{ formatPrice(stats.todaySales) }} تومان</p>
                      </div>
                      <div class="w-8 h-8 bg-emerald-500 rounded-md flex items-center justify-center">
                        <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
                        </svg>
                      </div>
                    </div>
                  </div>

                  <!-- فروش هفته -->
                  <div class="bg-gradient-to-br from-teal-50 to-teal-100 border border-teal-200 rounded-md p-3 hover:shadow-sm transition-all duration-200">
                    <div class="flex items-center justify-between">
                      <div>
                        <p class="text-xs font-medium text-teal-700 mb-0.5">فروش هفته</p>
                        <p class="text-lg font-bold text-teal-900">{{ formatPrice(stats.weeklySales) }} تومان</p>
                      </div>
                      <div class="w-8 h-8 bg-teal-500 rounded-md flex items-center justify-center">
                        <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
                        </svg>
                      </div>
                    </div>
                  </div>

                  <!-- فروش ماه -->
                  <div class="bg-gradient-to-br from-indigo-50 to-indigo-100 border border-indigo-200 rounded-md p-3 hover:shadow-sm transition-all duration-200">
                    <div class="flex items-center justify-between">
                      <div>
                        <p class="text-xs font-medium text-indigo-700 mb-0.5">فروش ماه</p>
                        <p class="text-lg font-bold text-indigo-900">{{ formatPrice(stats.monthlySales) }} تومان</p>
                      </div>
                      <div class="w-8 h-8 bg-indigo-500 rounded-md flex items-center justify-center">
                        <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
                        </svg>
                      </div>
                    </div>
                  </div>

                  <!-- متوسط سفارش -->
                  <div class="bg-gradient-to-br from-pink-50 to-pink-100 border border-pink-200 rounded-md p-3 hover:shadow-sm transition-all duration-200">
                    <div class="flex items-center justify-between">
                      <div>
                        <p class="text-xs font-medium text-pink-700 mb-0.5">متوسط سفارش</p>
                        <p class="text-lg font-bold text-pink-900">{{ formatPrice(stats.averageOrder) }} تومان</p>
                      </div>
                      <div class="w-8 h-8 bg-pink-500 rounded-md flex items-center justify-center">
                        <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 13v-1a4 4 0 014-4 4 4 0 014 4v1m0 0h-3m3 0h3m-9 7a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                        </svg>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <!-- فیلترهای پیشرفته -->
          <div class="rounded-lg shadow-md border border-gray-200 overflow-hidden mb-6">
            <div class="bg-gradient-to-r from-green-50 to-emerald-50 px-4 py-3 border-b border-gray-200">
              <div class="flex items-center justify-between">
                <div class="flex items-center">
                  <div class="w-6 h-6 bg-green-500 rounded-md flex items-center justify-center ml-2">
                    <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.207A1 1 0 013 6.5V4z"></path>
                    </svg>
                  </div>
                  <h3 class="text-sm font-semibold text-gray-900">فیلترهای پیشرفته سفارشات</h3>
                </div>
                <button class="text-sm text-green-600 hover:text-green-800 transition-colors font-medium hover:bg-green-50 px-3 py-1 rounded-lg" @click="showFilters = !showFilters">
                  {{ showFilters ? 'مخفی کردن' : 'نمایش' }}
                </button>
              </div>
            </div>
            
            <div v-if="showFilters" class="p-6">
            <OrderFilters @filter-change="handleFiltersChange" />
            </div>
          </div>
            
            <!-- لیست سفارشات -->
          <div class="rounded-lg shadow-md border border-gray-200 overflow-hidden">
              <div class="bg-gradient-to-r from-slate-50 to-gray-50 px-4 py-3 border-b border-gray-200">
                <div class="flex items-center justify-between">
                  <div class="flex items-center">
                    <div class="w-6 h-6 bg-slate-600 rounded-md flex items-center justify-center ml-2">
                      <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16"></path>
                      </svg>
                    </div>
                    <div>
                      <h3 class="text-sm font-semibold text-gray-900">لیست سفارشات</h3>
                      <p class="text-xs text-gray-500 mt-1">
                        نمایش {{ totalItems }} از {{ sampleOrders.length }} سفارش
                        <span v-if="searchTerm || Object.values(filters).some(v => v && v !== 'all')" class="text-blue-600">
                          (فیلتر شده)
                        </span>
                      </p>
                    </div>
                  </div>
                  
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <!-- Bulk Actions -->
                    <div v-if="selectedOrders.length > 0" class="flex items-center space-x-2 space-x-reverse bg-blue-50 rounded-md px-2 py-1.5 border border-blue-200">
                      <span class="text-xs text-blue-700 font-medium">{{ selectedOrders.length }} مورد انتخاب شده</span>
                      <select 
                        v-model="bulkAction"
                        class="text-xs border border-blue-300 rounded px-2 py-1 bg-blue-50"
                        @change="executeBulkAction"
                      >
                        <option value="">عملیات گروهی</option>
                        <option value="confirm">تایید سفارشات</option>
                        <option value="process">آماده‌سازی</option>
                        <option value="ship">ارسال</option>
                        <option value="deliver">تحویل</option>
                        <option value="cancel">لغو</option>
                        <option value="print">چاپ فاکتور</option>
                        <option value="export">خروجی Excel</option>
                      </select>
                    </div>
                    

                    
                    <!-- Search -->
                    <div class="relative">
                      <input 
                        v-model="searchTerm"
                        type="text" 
                        class="block w-56 pl-8 pr-3 py-1.5 border border-gray-300 rounded-md leading-5 bg-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-xs shadow-sm" 
                        placeholder="جستجو شماره سفارش، مشتری یا کد پیگیری"
                        dir="rtl"
                      />
                      <div class="absolute inset-y-0 left-0 pl-2 flex items-center pointer-events-none">
                        <svg class="h-4 w-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
                        </svg>
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              <!-- Table Content -->
              <OrderTable 
                :orders="paginatedOrders"
                :loading="ordersPending"
                @view-details="openOrderDetail"
                @quick-actions="openQuickActions"
                @refresh="loadOrders"
                @export="exportOrders"
                @bulk-delete="handleBulkDelete"
              />
            
            <!-- صفحه‌بندی -->
            <Pagination
              :current-page="currentPage"
              :total-pages="totalPages"
              :total="totalItems"
              :items-per-page="itemsPerPage"
              @page-changed="handlePageChange"
              @items-per-page-changed="handleItemsPerPageChange"
            />
            </div>
          </div>

          <!-- تب اتصال حسابداری -->
          <div v-if="activeTab === 'accounting'">
            <OrderAccountingIntegration />
          </div>



        <!-- تب گزارشات -->
        <div v-if="activeTab === 'reports'">
          <!-- آمار کلی گزارشات -->
          <div class="rounded-lg shadow-md border border-gray-200 overflow-hidden mb-6">
            <div class="bg-gradient-to-r from-indigo-50 to-purple-50 px-4 py-3 border-b border-gray-200">
              <div class="flex items-center">
                <div class="w-6 h-6 bg-indigo-500 rounded-md flex items-center justify-center ml-2">
                  <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
                  </svg>
                </div>
                <h3 class="text-sm font-semibold text-gray-900">گزارشات جامع و تحلیلی سفارشات</h3>
              </div>
            </div>

            <div class="p-6">
              <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
                <!-- کل سفارشات -->
                <div class="bg-gradient-to-br from-blue-50 to-blue-100 border border-blue-200 rounded-lg p-6 hover:shadow-md transition-all duration-200">
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm font-medium text-blue-700 mb-1">کل سفارشات</p>
                      <p class="text-2xl font-bold text-blue-900">{{ comprehensiveStats.totalOrders.toLocaleString('fa-IR') }}</p>
                      <p class="text-xs text-blue-600 mt-1">در تمام زمان‌ها</p>
                    </div>
                    <div class="w-12 h-12 bg-blue-500 rounded-lg flex items-center justify-center">
                      <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"></path>
                      </svg>
                    </div>
                  </div>
                </div>

                <!-- ارزش کل فروش -->
                <div class="bg-gradient-to-br from-green-50 to-green-100 border border-green-200 rounded-lg p-6 hover:shadow-md transition-all duration-200">
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm font-medium text-green-700 mb-1">ارزش کل فروش</p>
                      <p class="text-2xl font-bold text-green-900">{{ formatPrice(comprehensiveStats.totalRevenue) }}</p>
                      <p class="text-xs text-green-600 mt-1">در تمام زمان‌ها</p>
                    </div>
                    <div class="w-12 h-12 bg-green-500 rounded-lg flex items-center justify-center">
                      <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
                      </svg>
                    </div>
                  </div>
                </div>

                <!-- نرخ تبدیل -->
                <div class="bg-gradient-to-br from-purple-50 to-purple-100 border border-purple-200 rounded-lg p-6 hover:shadow-md transition-all duration-200">
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm font-medium text-purple-700 mb-1">نرخ تبدیل</p>
                      <p class="text-2xl font-bold text-purple-900">{{ comprehensiveStats.conversionRate }}%</p>
                      <p class="text-xs text-purple-600 mt-1">از بازدید به خرید</p>
                    </div>
                    <div class="w-12 h-12 bg-purple-500 rounded-lg flex items-center justify-center">
                      <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
                      </svg>
                    </div>
                  </div>
                </div>

                <!-- متوسط سبد خرید -->
                <div class="bg-gradient-to-br from-orange-50 to-orange-100 border border-orange-200 rounded-lg p-6 hover:shadow-md transition-all duration-200">
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm font-medium text-orange-700 mb-1">متوسط سبد خرید</p>
                      <p class="text-2xl font-bold text-orange-900">{{ formatPrice(comprehensiveStats.avgOrderValue) }}</p>
                      <p class="text-xs text-orange-600 mt-1">به ازای هر سفارش</p>
                    </div>
                    <div class="w-12 h-12 bg-orange-500 rounded-lg flex items-center justify-center">
                      <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
                      </svg>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <!-- نمودارهای تحلیلی -->
            <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-6">
              <!-- نمودار روند فروش سالانه -->
              <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
                <div class="flex items-center justify-between mb-4">
                  <h3 class="text-lg font-semibold text-gray-900">روند فروش سالانه</h3>
                  <select v-model="selectedYear" class="text-sm border border-gray-300 rounded-md px-3 py-1">
                    <option value="2024">2024</option>
                    <option value="2023">2023</option>
                    <option value="2022">2022</option>
                  </select>
                </div>
                
                <!-- نمودار بهبود یافته -->
                <div class="relative h-64">
                  <!-- خطوط راهنما -->
                  <div class="absolute inset-0 flex flex-col justify-between text-xs text-gray-400">
                    <div class="border-b border-gray-100 pb-1">{{ formatPrice(maxYearlyValue) }}</div>
                    <div class="border-b border-gray-100 pb-1">{{ formatPrice(maxYearlyValue * 0.75) }}</div>
                    <div class="border-b border-gray-100 pb-1">{{ formatPrice(maxYearlyValue * 0.5) }}</div>
                    <div class="border-b border-gray-100 pb-1">{{ formatPrice(maxYearlyValue * 0.25) }}</div>
                    <div class="pb-1">0</div>
                  </div>
                  
                  <!-- ستون‌های نمودار -->
                  <div class="absolute inset-0 flex items-end justify-between space-x-1 space-x-reverse pr-16">
                    <div
                      v-for="(month, index) in yearlySalesData"
                      :key="index"
                      class="flex-1 flex flex-col items-center relative group"
                    >
                      <!-- ستون -->
                      <div
                        class="w-full bg-gradient-to-t from-indigo-500 to-indigo-400 rounded-t transition-all duration-200 hover:from-indigo-600 hover:to-indigo-500"
                        :style="{ height: `${(month.value / maxYearlyValue) * 240}px` }"
                      ></div>
                      
                      <!-- نام ماه -->
                      <div class="text-xs text-gray-600 mt-2 text-center">{{ month.name }}</div>
                      
                      <!-- مقدار در tooltip -->
                      <div class="absolute bottom-full mb-2 opacity-0 group-hover:opacity-100 transition-opacity duration-200 bg-gray-800 text-white text-xs px-2 py-1 rounded whitespace-nowrap z-10">
                        {{ formatPrice(month.value) }}
                      </div>
                  </div>
                </div>
              </div>

                <!-- راهنمای نمودار -->
                <div class="mt-4 text-xs text-gray-500 text-center">
                  مقادیر در tooltip نمایش داده می‌شوند
                </div>
              </div>

              <!-- نمودار مقایسه وضعیت‌ها -->
              <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
                <div class="flex items-center justify-between mb-4">
                  <h3 class="text-lg font-semibold text-gray-900">مقایسه وضعیت سفارشات</h3>
                  <span class="text-sm text-gray-500">آخرین 30 روز</span>
                </div>
                <div class="space-y-4">
                  <div
                    v-for="status in orderStatusComparison"
                    :key="status.name"
                    class="flex items-center justify-between"
                  >
                    <div class="flex items-center">
                      <div class="w-4 h-4 rounded-full ml-3" :style="{ backgroundColor: status.color }"></div>
                      <span class="text-sm text-gray-700">{{ status.name }}</span>
                  </div>
                    <div class="flex items-center space-x-3 space-x-reverse">
                      <div class="w-32 bg-gray-200 rounded-full h-3">
                        <div
                          class="h-3 rounded-full"
                          :style="{ width: `${Math.min(Math.max(status.percentage, 0), 100)}%`, backgroundColor: status.color }"
                        ></div>
                      </div>
                      <span class="text-sm font-medium text-gray-900 w-16 text-left">{{ Math.round(Math.max(status.percentage, 0)) }}%</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <!-- تحلیل‌های پیشرفته -->
            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-6">
              <!-- تحلیل جغرافیایی -->
              <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
                <div class="flex items-center justify-between mb-3">
                  <h4 class="text-sm font-semibold text-gray-900">تحلیل جغرافیایی</h4>
                  <svg class="w-5 h-5 text-blue-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"></path>
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"></path>
                </svg>
                </div>
                <div class="space-y-2">
                  <div
                    v-for="geo in geoAnalysis.slice(0, 4)"
                    :key="geo.cityName"
                    class="flex items-center justify-between text-xs"
                  >
                    <span class="text-gray-600">{{ geo.cityName }}</span>
                    <span class="font-medium">{{ geo.percentage.toFixed(1) }}%</span>
                  </div>
                </div>
              </div>

              <!-- تحلیل زمانی -->
              <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
                <div class="flex items-center justify-between mb-3">
                  <h4 class="text-sm font-semibold text-gray-900">تحلیل زمانی</h4>
                  <svg class="w-5 h-5 text-green-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
                </div>
                <div class="space-y-2">
                  <div class="flex items-center justify-between text-xs">
                    <span class="text-gray-600">ساعات اوج</span>
                    <span class="font-medium">{{ timeAnalysis.peakHour }}</span>
                  </div>
                  <div class="flex items-center justify-between text-xs">
                    <span class="text-gray-600">روزهای پرفروش</span>
                    <span class="font-medium">{{ timeAnalysis.busiestDay }}</span>
                  </div>
                  <div class="flex items-center justify-between text-xs">
                    <span class="text-gray-600">فصل پرفروش</span>
                    <span class="font-medium">{{ timeAnalysis.busiestSeason }}</span>
                  </div>
                  <div class="flex items-center justify-between text-xs">
                    <span class="text-gray-600">متوسط زمان خرید</span>
                    <span class="font-medium">{{ Math.round(timeAnalysis.avgPurchaseTime) }} دقیقه</span>
                  </div>
                </div>
              </div>

              <!-- تحلیل مشتریان -->
              <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
                <div class="flex items-center justify-between mb-3">
                  <h4 class="text-sm font-semibold text-gray-900">تحلیل مشتریان</h4>
                  <svg class="w-5 h-5 text-purple-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
                </svg>
                </div>
                <div class="space-y-2">
                  <div class="flex items-center justify-between text-xs">
                    <span class="text-gray-600">مشتریان جدید</span>
                    <span class="font-medium">{{ customerAnalysis.newCustomers }}</span>
                  </div>
                  <div class="flex items-center justify-between text-xs">
                    <span class="text-gray-600">مشتریان وفادار</span>
                    <span class="font-medium">{{ customerAnalysis.returnCustomers }}</span>
                  </div>
                  <div class="flex items-center justify-between text-xs">
                    <span class="text-gray-600">متوسط ارزش سفارش</span>
                    <span class="font-medium">{{ formatPrice(customerAnalysis.avgOrderValue) }}</span>
                  </div>
                  <div class="flex items-center justify-between text-xs">
                    <span class="text-gray-600">متوسط سن</span>
                    <span class="font-medium">{{ customerAnalysis.avgAge }} سال</span>
                  </div>
                </div>
              </div>

              <!-- تحلیل محصولات -->
              <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
                <div class="flex items-center justify-between mb-3">
                  <h4 class="text-sm font-semibold text-gray-900">تحلیل محصولات</h4>
                  <svg class="w-5 h-5 text-orange-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
                </svg>
                </div>
                <div class="space-y-2">
                  <div
                    v-for="product in productAnalysis.slice(0, 4)"
                    :key="product.categoryName"
                    class="flex items-center justify-between text-xs"
                  >
                    <span class="text-gray-600">{{ product.categoryName }}</span>
                    <span class="font-medium">{{ product.percentage.toFixed(1) }}%</span>
                  </div>
                </div>
              </div>
            </div>

            <!-- فیلترهای پیشرفته -->
            <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden mb-6">
              <div class="bg-gradient-to-r from-indigo-50 to-purple-50 px-4 py-3 border-b border-gray-200">
                <div class="flex items-center justify-between">
                  <div class="flex items-center">
                    <div class="w-6 h-6 bg-indigo-500 rounded-md flex items-center justify-center ml-2">
                      <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.207A1 1 0 013 6.5V4z"></path>
                    </svg>
                </div>
                  <h3 class="text-sm font-semibold text-gray-900">فیلترهای پیشرفته گزارشات</h3>
                </div>
                <button class="text-sm text-indigo-600 hover:text-indigo-800 transition-colors font-medium hover:bg-indigo-50 px-3 py-1 rounded-lg" @click="showReportFilters = !showReportFilters">
                  {{ showReportFilters ? 'مخفی کردن' : 'نمایش' }}
                </button>
              </div>
            </div>

            <div v-if="showReportFilters" class="p-6">
              <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
                <!-- فیلتر دوره زمانی -->
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">دوره زمانی</label>
                  <select v-model="reportFilters.timePeriod" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500">
                    <option value="">همه دوره‌ها</option>
                    <option value="today">امروز</option>
                    <option value="week">هفته جاری</option>
                    <option value="month">ماه جاری</option>
                    <option value="quarter">فصل جاری</option>
                    <option value="year">سال جاری</option>
                    <option value="custom">انتخاب دستی</option>
                </select>
              </div>

                <!-- فیلتر محدوده مبلغ -->
                  <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">محدوده مبلغ</label>
                  <select v-model="reportFilters.amountRange" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500">
                    <option value="">همه مبالغ</option>
                    <option value="0-1000000">کمتر از 1 میلیون تومان</option>
                    <option value="1000000-5000000">1 تا 5 میلیون تومان</option>
                    <option value="5000000-10000000">5 تا 10 میلیون تومان</option>
                    <option value="10000000+">بیش از 10 میلیون تومان</option>
                  </select>
                  </div>

                <!-- فیلتر دسته‌بندی محصول -->
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">دسته‌بندی محصول</label>
                  <select v-model="reportFilters.category" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500">
                    <option value="">همه دسته‌ها</option>
                    <option value="digital">محصولات دیجیتال</option>
                    <option value="home">لوازم خانگی</option>
                    <option value="clothing">پوشاک</option>
                    <option value="other">سایر</option>
                  </select>
                  </div>

                <!-- فیلتر منطقه جغرافیایی -->
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">منطقه جغرافیایی</label>
                  <select v-model="reportFilters.geographicRegion" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500">
                    <option value="">همه مناطق</option>
                    <option value="tehran">تهران</option>
                    <option value="isfahan">اصفهان</option>
                    <option value="mashhad">مشهد</option>
                    <option value="shiraz">شیراز</option>
                    <option value="tabriz">تبریز</option>
                    <option value="other">سایر شهرها</option>
                  </select>
                </div>
              </div>

              <!-- ردیف دوم فیلترها -->
              <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mt-4">
                <!-- فیلتر نوع مشتری -->
                  <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">نوع مشتری</label>
                  <select v-model="reportFilters.customerType" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500">
                    <option value="">همه مشتریان</option>
                    <option value="new">مشتریان جدید</option>
                    <option value="loyal">مشتریان وفادار</option>
                    <option value="returning">مشتریان بازگشتی</option>
                    <option value="vip">مشتریان VIP</option>
                  </select>
                  </div>

                <!-- فیلتر نرخ تبدیل -->
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">نرخ تبدیل</label>
                  <select v-model="reportFilters.conversionRate" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500">
                    <option value="">همه نرخ‌ها</option>
                    <option value="0-1">کمتر از 1%</option>
                    <option value="1-2">1 تا 2%</option>
                    <option value="2-3">2 تا 3%</option>
                    <option value="3-4">3 تا 4%</option>
                    <option value="4+">بیش از 4%</option>
                  </select>
                  </div>

                <!-- فیلتر وضعیت سفارش -->
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت سفارش</label>
                  <select v-model="reportFilters.orderStatus" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500">
                    <option value="">همه وضعیت‌ها</option>
                    <option value="pending">در انتظار پرداخت</option>
                    <option value="processing">در حال پردازش</option>
                    <option value="shipped">ارسال شده</option>
                    <option value="delivered">تحویل شده</option>
                    <option value="returned">مرجوع شده</option>
                    <option value="cancelled">لغو شده</option>
                  </select>
                </div>

                <!-- فیلتر روش پرداخت -->
                  <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">روش پرداخت</label>
                  <select v-model="reportFilters.paymentMethod" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500">
                    <option value="">همه روش‌ها</option>
                    <option value="online">پرداخت آنلاین</option>
                    <option value="cash">پرداخت نقدی</option>
                    <option value="bank_transfer">انتقال بانکی</option>
                    <option value="wallet">کیف پول</option>
                    <option value="gift_card">گیفت کارت</option>
                  </select>
                  </div>
                  </div>

              <!-- ردیف سوم فیلترها -->
              <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mt-4">
                <!-- فیلتر بازه سنی -->
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">بازه سنی</label>
                  <select v-model="reportFilters.ageRange" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500">
                    <option value="">همه سنین</option>
                    <option value="18-25">18 تا 25 سال</option>
                    <option value="26-35">26 تا 35 سال</option>
                    <option value="36-45">36 تا 45 سال</option>
                    <option value="46-55">46 تا 55 سال</option>
                    <option value="55+">بیش از 55 سال</option>
                  </select>
                </div>

                <!-- فیلتر نرخ بازگشت -->
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">نرخ بازگشت</label>
                  <select v-model="reportFilters.returnRate" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500">
                    <option value="">همه نرخ‌ها</option>
                    <option value="0-5">کمتر از 5%</option>
                    <option value="5-10">5 تا 10%</option>
                    <option value="10-15">10 تا 15%</option>
                    <option value="15-20">15 تا 20%</option>
                    <option value="20+">بیش از 20%</option>
                  </select>
              </div>

                <!-- فیلتر نوع گزارش -->
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">نوع گزارش</label>
                  <select v-model="reportFilters.reportType" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500">
                    <option value="">همه گزارشات</option>
                    <option value="sales">گزارش فروش</option>
                    <option value="customer">گزارش مشتریان</option>
                    <option value="product">گزارش محصولات</option>
                    <option value="geographic">گزارش جغرافیایی</option>
                    <option value="temporal">گزارش زمانی</option>
                  </select>
                </div>
              </div>

              <!-- دکمه‌های عملیات -->
              <div class="mt-6 flex justify-between items-center">
                <div class="flex space-x-2 space-x-reverse">
                  <button class="px-4 py-2 bg-gray-100 text-gray-700 rounded-md hover:bg-gray-200 transition-colors" @click="clearReportFilters">
                    پاک کردن فیلترها
                  </button>
                  <button class="px-4 py-2 bg-blue-100 text-blue-700 rounded-md hover:bg-blue-200 transition-colors" @click="resetReportFilters">
                    بازنشانی فیلترها
                  </button>
                </div>
                
                <div class="flex space-x-2 space-x-reverse">
                  <button class="px-4 py-2 bg-green-500 text-white rounded-md hover:bg-green-600 transition-colors flex items-center" @click="exportFilteredReports">
                    <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                    </svg>
                    خروجی Excel
                  </button>
                  <button class="px-4 py-2 bg-red-500 text-white rounded-md hover:bg-red-600 transition-colors flex items-center" @click="exportFilteredReportsPDF">
                    <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 17h2a2 2 0 002-2v-4a2 2 0 00-2-2H5a2 2 0 00-2 2v4a2 2 0 002 2h2m2 4h6a2 2 0 002-2v-4a2 2 0 00-2-2H9a2 2 0 00-2 2v4a2 2 0 002 2zm8-12V5a2 2 0 00-2-2H9a2 2 0 00-2 2v4h10z"></path>
                    </svg>
                    خروجی PDF
                  </button>
                </div>
              </div>
            </div>
          </div>

          <!-- جدول گزارشات تفصیلی -->
          <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
            <div class="bg-gradient-to-r from-slate-50 to-gray-50 px-4 py-3 border-b border-gray-200">
              <div class="flex items-center justify-between">
                <div class="flex items-center">
                  <div class="w-6 h-6 bg-slate-600 rounded-md flex items-center justify-center ml-2">
                    <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16"></path>
                    </svg>
                  </div>
                  <h3 class="text-sm font-semibold text-gray-900">گزارشات تفصیلی و تحلیلی</h3>
                </div>
                
                <div class="flex items-center space-x-2 space-x-reverse">
                  <!-- Search -->
                  <div class="relative">
                    <input 
                      v-model="reportSearchTerm"
                      type="text" 
                      class="block w-56 pl-8 pr-3 py-1.5 border border-gray-300 rounded-md leading-5 bg-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 text-xs shadow-sm" 
                      placeholder="جستجو در گزارشات..."
                      dir="rtl"
                    />
                    <div class="absolute inset-y-0 left-0 pl-2 flex items-center pointer-events-none">
                      <svg class="h-4 w-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
                      </svg>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <!-- Table Content -->
            <div class="overflow-x-auto">
              <table class="w-full border-collapse">
                <thead>
                  <tr class="bg-gray-50 border-b border-gray-200">
                    <th class="px-4 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider border-r border-gray-200" style="width: 20%;">
                      <button 
                        class="flex items-center justify-center gap-1 w-full cursor-pointer hover:bg-gray-100 transition-colors select-none py-1" 
                        @click="sortReports('period')"
                      >
                        <span :class="getSortClass('period')">دوره</span>
                        <svg class="w-3 h-3 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="getSortIcon('period')"></path>
                        </svg>
                      </button>
                    </th>
                    <th class="px-4 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider border-r border-gray-200" style="width: 15%;">
                      <button 
                        class="flex items-center justify-center gap-1 w-full cursor-pointer hover:bg-gray-100 transition-colors select-none py-1" 
                        @click="sortReports('orderCount')"
                      >
                        <span :class="getSortClass('orderCount')">تعداد سفارش</span>
                        <svg class="w-3 h-3 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="getSortIcon('orderCount')"></path>
                        </svg>
                      </button>
                    </th>
                    <th class="px-4 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider border-r border-gray-200" style="width: 20%;">
                      <button 
                        class="flex items-center justify-center gap-1 w-full cursor-pointer hover:bg-gray-100 transition-colors select-none py-1" 
                        @click="sortReports('totalRevenue')"
                      >
                        <span :class="getSortClass('totalRevenue')">ارزش فروش</span>
                        <svg class="w-3 h-3 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="getSortIcon('totalRevenue')"></path>
                        </svg>
                      </button>
                    </th>
                    <th class="px-4 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider border-r border-gray-200" style="width: 20%;">
                      <button 
                        class="flex items-center justify-center gap-1 w-full cursor-pointer hover:bg-gray-100 transition-colors select-none py-1" 
                        @click="sortReports('avgOrderValue')"
                      >
                        <span :class="getSortClass('avgOrderValue')">متوسط سبد خرید</span>
                        <svg class="w-3 h-3 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="getSortIcon('avgOrderValue')"></path>
                        </svg>
                      </button>
                    </th>
                    <th class="px-4 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider border-r border-gray-200" style="width: 12%;">
                      <button 
                        class="flex items-center justify-center gap-1 w-full cursor-pointer hover:bg-gray-100 transition-colors select-none py-1" 
                        @click="sortReports('conversionRate')"
                      >
                        <span :class="getSortClass('conversionRate')">نرخ تبدیل</span>
                        <svg class="w-3 h-3 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="getSortIcon('conversionRate')"></path>
                        </svg>
                      </button>
                    </th>
                    <th class="px-4 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider" style="width: 13%;">
                      <button 
                        class="flex items-center justify-center gap-1 w-full cursor-pointer hover:bg-gray-100 transition-colors select-none py-1" 
                        @click="sortReports('change')"
                      >
                        <span :class="getSortClass('change')">تغییر</span>
                        <svg class="w-3 h-3 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="getSortIcon('change')"></path>
                        </svg>
                      </button>
                    </th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="report in paginatedDetailedReports" :key="report.id" class="hover:bg-gray-50 border-b border-gray-200">
                <td class="px-4 py-4 text-sm font-medium text-gray-900 text-center border-r border-gray-200" style="width: 20%;">
                  {{ formatPersianDate(report.period) }}
                </td>
                    <td class="px-4 py-4 text-sm text-gray-900 text-center border-r border-gray-200" style="width: 15%;">
                      {{ report.orderCount.toLocaleString('fa-IR') }}
                    </td>
                    <td class="px-4 py-4 text-sm text-gray-900 text-center border-r border-gray-200" style="width: 20%;">
                      {{ formatPrice(report.totalRevenue) }}
                    </td>
                    <td class="px-4 py-4 text-sm text-gray-900 text-center border-r border-gray-200" style="width: 20%;">
                      {{ formatPrice(report.avgOrderValue) }}
                    </td>
                    <td class="px-4 py-4 text-center border-r border-gray-200" style="width: 12%;">
                      <span :class="getConversionRateClass(report.conversionRate)" class="px-2 py-1 text-xs font-medium rounded-full">
                        {{ report.conversionRate }}%
                      </span>
                    </td>
                    <td class="px-4 py-4 text-center" style="width: 13%;">
                      <span :class="getChangeClass(report.change)" class="px-2 py-1 text-xs font-medium rounded-full">
                        {{ report.change > 0 ? '+' : '' }}{{ report.change }}%
                      </span>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>

            <!-- Pagination -->
            <Pagination
              :current-page="reportCurrentPage"
              :total-pages="reportTotalPages"
              :total="filteredDetailedReports.length"
              :items-per-page="reportItemsPerPage"
              @page-changed="handleReportPageChange"
              @items-per-page-changed="handleReportItemsPerPageChange"
            />
          </div>
        </div>
      </div>
    </div>
    </div>

    <!-- Order Detail Modal -->
    <OrderDetailModal 
      :is-open="showOrderDetail"
      :order="selectedOrder"
      @close="showOrderDetail = false"
    />

    <!-- Create Order Modal -->
    <CreateOrderModal 
      :is-open="showCreateOrderModal"
      @close="showCreateOrderModal = false"
      @created="handleOrderCreated"
    />

    <!-- Quick Actions Modal -->
    <QuickActionsModal 
      :is-open="showQuickActions"
      :order="selectedOrder"
      @close="showQuickActions = false"
      @status-change="handleStatusChange"
      @notification="handleNotification"
      @print="handlePrint"
      @duplicate="handleDuplicate"
      @export="handleExport"
      @cancel="handleCancel"
      @delete="handleDelete"
    />
  </ClientOnly>
  
  <!-- Confirm Dialog -->
  <ConfirmDialog ref="confirmDialog" />
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
</script>

<script setup lang="ts">
import { useAsyncData } from 'nuxt/app'
import { computed, ref, watch } from 'vue'
import Pagination from '~/components/admin/common/Pagination.vue'
import CreateOrderModal from './CreateOrderModal.vue'
import OrderDetailModal from './OrderDetailModal.vue'
import OrderFilters from './OrderFilters.vue'
import OrderTable from './OrderTable.vue'
import QuickActionsModal from './QuickActionsModal.vue'
import OrderAccountingIntegration from './components/OrderAccountingIntegration.vue'

interface Order {
  id: number
  orderNumber: string
  customerName: string
  customerPhone?: string
  trackingCode?: string
  status: string
  createdAt: string
  totalAmount: number
  [key: string]: unknown
}

interface AnalyticsData {
  comprehensiveStats: {
    totalOrders: number
    totalRevenue: number
    conversionRate: number
    avgOrderValue: number
  }
  yearlySalesData: Record<string, unknown>[]
  orderStatusComparison: Record<string, unknown>[]
  paymentMethodStats: Record<string, unknown>[]
  detailedReports: Record<string, unknown>[]
  productAnalysis: Record<string, unknown>[]
   customerAnalysis: Record<string, unknown>
  timeAnalysis: Record<string, unknown>
  geoAnalysis: Record<string, unknown>[]
}

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
// Auth disabled
// Reactive data
const searchTerm = ref('')
const selectedOrders = ref([])
const selectedOrder = ref(null)
const showOrderDetail = ref(false)
const showSuccessMessage = ref(false)
const successMessage = ref('')
const confirmDialog = ref()

// متغیرهای گزارشات جامع - داده‌های واقعی از API
const comprehensiveStats = ref({
  totalOrders: 0,
  totalRevenue: 0,
  conversionRate: 0,
  avgOrderValue: 0
})

const selectedYear = ref('2024')

// داده‌های نمودار فروش سالانه - از API دریافت می‌شود
const yearlySalesData = ref([])

// مقایسه وضعیت سفارشات - از API دریافت می‌شود
const orderStatusComparison = ref([
  { name: 'ارسال شده', percentage: 45, color: '#10B981' },
  { name: 'لغو شده', percentage: 20, color: '#6B7280' },
  { name: 'در حال پردازش', percentage: 25, color: '#3B82F6' },
  { name: 'بازپرداخت شده', percentage: 10, color: '#F59E0B' }
])

// آمار روش‌های پرداخت - از API دریافت می‌شود
const paymentMethodStats = ref([])

// تحلیل‌های پیشرفته - از API دریافت می‌شود
const productAnalysis = ref([])
const customerAnalysis = ref({
  newCustomers: 0,
  returnCustomers: 0,
  avgOrderValue: 0,
  avgAge: 0
})
const timeAnalysis = ref({
  peakHour: '',
  busiestDay: '',
  busiestMonth: '',
  busiestSeason: '',
  avgPurchaseTime: 0
})
const geoAnalysis = ref([])

// متغیرهای فیلتر گزارشات
const showReportFilters = ref(false)
const reportSearchTerm = ref('')
const reportFilters = ref({
  timePeriod: '',
  amountRange: '',
  category: '',
  geographicRegion: '',
  customerType: '',
  conversionRate: '',
  orderStatus: '',
  paymentMethod: '',
  ageRange: '',
  returnRate: '',
  reportType: ''
})

// متغیرهای صفحه‌بندی گزارشات
const reportCurrentPage = ref(1)
const reportItemsPerPage = ref(10)

// متغیرهای مرتب‌سازی گزارشات
const reportSortField = ref('period')
const reportSortDirection = ref('desc') // پیش‌فرض: نزولی (آخرین تاریخ)

// تابع مرتب‌سازی گزارشات
const sortReports = (field) => {
  if (reportSortField.value === field) {
    reportSortDirection.value = reportSortDirection.value === 'asc' ? 'desc' : 'asc'
  } else {
    reportSortField.value = field
    reportSortDirection.value = 'asc'
  }
}

// تابع دریافت کلاس مرتب‌سازی
const getSortClass = (field) => {
  if (reportSortField.value !== field) return ''
  return reportSortDirection.value === 'asc' ? 'text-blue-600' : 'text-blue-600'
}

// تابع دریافت آیکون مرتب‌سازی
const getSortIcon = (field) => {
  if (reportSortField.value !== field) {
    return 'M7 16V4m0 0L3 8m4-4l4 4m6 0v12m0 0l4-4m-4 4l-4-4'
  }
  return reportSortDirection.value === 'asc' 
    ? 'M5 15l7-7 7 7'
    : 'M19 9l-7 7-7-7'
}

// داده‌های گزارشات تفصیلی - از API دریافت می‌شود
const detailedReports = ref([])
const showCreateOrderModal = ref(false)
const showQuickActions = ref(false)

// Computed properties برای گزارشات
const maxYearlyValue = computed(() => {
  return Math.max(...yearlySalesData.value.map(item => item.value))
})

const filteredDetailedReports = computed(() => {
  let filtered = detailedReports.value

  // فیلتر بر اساس جستجو
  if (reportSearchTerm.value) {
    const search = reportSearchTerm.value.toLowerCase()
    filtered = filtered.filter(report => 
      report.period.toLowerCase().includes(search)
    )
  }

  // فیلتر بر اساس دوره زمانی
  if (reportFilters.value.timePeriod) {
    filtered = filtered.filter(report => 
      report.period.includes(reportFilters.value.timePeriod)
    )
  }

  // فیلتر بر اساس محدوده مبلغ
  if (reportFilters.value.amountRange) {
    const [min, max] = reportFilters.value.amountRange.split('-').map(Number)
    filtered = filtered.filter(report => {
      if (max) {
        return report.totalRevenue >= min && report.totalRevenue <= max
      } else {
        return report.totalRevenue >= min
      }
    })
  }

  // مرتب‌سازی
  filtered.sort((a, b) => {
    let aValue, bValue
    
    switch (reportSortField.value) {
      case 'period':
        // برای دوره، مرتب‌سازی بر اساس تاریخ
        aValue = new Date(a.period)
        bValue = new Date(b.period)
        break
      case 'orderCount':
        aValue = a.orderCount
        bValue = b.orderCount
        break
      case 'totalRevenue':
        aValue = a.totalRevenue
        bValue = b.totalRevenue
        break
      case 'avgOrderValue':
        aValue = a.avgOrderValue
        bValue = b.avgOrderValue
        break
      case 'conversionRate':
        aValue = a.conversionRate
        bValue = b.conversionRate
        break
      case 'change':
        aValue = a.change
        bValue = b.change
        break
      default:
        aValue = a.period
        bValue = b.period
    }

    if (reportSortDirection.value === 'asc') {
      return aValue > bValue ? 1 : -1
    } else {
      return aValue < bValue ? 1 : -1
    }
  })

  return filtered
})

const reportTotalPages = computed(() => {
  return Math.ceil(filteredDetailedReports.value.length / reportItemsPerPage.value)
})

const paginatedDetailedReports = computed(() => {
  const start = (reportCurrentPage.value - 1) * reportItemsPerPage.value
  const end = start + reportItemsPerPage.value
  return filteredDetailedReports.value.slice(start, end)
})

// توابع گزارشات
const clearReportFilters = () => {
  reportFilters.value = {
    timePeriod: '',
    amountRange: '',
    category: '',
    geographicRegion: '',
    customerType: '',
    conversionRate: '',
    orderStatus: '',
    paymentMethod: '',
    ageRange: '',
    returnRate: '',
    reportType: ''
  }
  reportSearchTerm.value = ''
}

const resetReportFilters = () => {
  clearReportFilters()
  reportCurrentPage.value = 1
}

const exportFilteredReports = () => {
  // خروجی Excel گزارشات فیلتر شده
  // اینجا می‌توانید منطق خروجی Excel را پیاده‌سازی کنید
}

const exportFilteredReportsPDF = () => {
  // خروجی PDF گزارشات فیلتر شده
  // اینجا می‌توانید منطق خروجی PDF را پیاده‌سازی کنید
}

const handleReportPageChange = (page) => {
  reportCurrentPage.value = page
}

// دریافت تحلیل کامل سفارشات از API
const fetchAnalytics = async () => {
  try {
    const response = await $fetch('/api/admin/orders/real-reports') as { success: boolean, data: AnalyticsData }
    
    if (response && response.success && response.data) {
      const data: AnalyticsData = response.data
      
      // به‌روزرسانی آمار جامع
      comprehensiveStats.value = {
        totalOrders: data.comprehensiveStats?.totalOrders || 0,
        totalRevenue: data.comprehensiveStats?.totalRevenue || 0,
        conversionRate: data.comprehensiveStats?.conversionRate || 0,
        avgOrderValue: data.comprehensiveStats?.avgOrderValue || 0
      }
      
      // به‌روزرسانی داده‌های نمودار سالانه
      yearlySalesData.value = data.yearlySalesData || []
      
      // به‌روزرسانی آمار وضعیت‌ها
      orderStatusComparison.value = (data.orderStatusComparison || []).map((item: Record<string, unknown>) => ({
        name: getStatusText(item.status as string),
        percentage: item.percentage as number,
        color: getStatusColor(item.status as string)
      }))
      
      // به‌روزرسانی آمار روش‌های پرداخت
      paymentMethodStats.value = data.paymentMethodStats || []
      
      // به‌روزرسانی گزارشات تفصیلی
      detailedReports.value = (data.detailedReports || []).map((report: Record<string, unknown>, index: number) => ({
        id: index + 1,
        period: report.period as string,
        orderCount: report.orderCount as number,
        totalRevenue: report.totalRevenue as number,
        avgOrderValue: report.avgOrderValue as number,
        conversionRate: report.conversionRate as number,
        change: report.change as number
      }))
    }
  } catch (error) {
    console.error('خطا در دریافت تحلیل سفارشات:', error)
  }
}

// دریافت تحلیل پیشرفته سفارشات از API
const fetchAdvancedAnalytics = async () => {
  try {
    const response = await $fetch('/api/admin/orders/advanced-analytics') as { success: boolean, data: AnalyticsData }
    
    if (response && response.success && response.data) {
      const data: AnalyticsData = response.data
      
      // به‌روزرسانی تحلیل محصولات
      productAnalysis.value = data.productAnalysis || []
      
      // به‌روزرسانی تحلیل مشتریان
      customerAnalysis.value = {
        newCustomers: (data.customerAnalysis?.newCustomers as number) || 0,
        returnCustomers: (data.customerAnalysis?.returnCustomers as number) || 0,
        avgOrderValue: (data.customerAnalysis?.avgOrderValue as number) || 0,
        avgAge: (data.customerAnalysis?.avgAge as number) || 0
      }
      
      // به‌روزرسانی تحلیل زمانی
      timeAnalysis.value = {
        peakHour: (data.timeAnalysis?.peakHour as string) || '',
        busiestDay: (data.timeAnalysis?.busiestDay as string) || '',
        busiestMonth: (data.timeAnalysis?.busiestMonth as string) || '',
        busiestSeason: (data.timeAnalysis?.busiestSeason as string) || '',
        avgPurchaseTime: (data.timeAnalysis?.avgPurchaseTime as number) || 0
      }
      
      // به‌روزرسانی تحلیل جغرافیایی
      geoAnalysis.value = data.geoAnalysis || []
    }
  } catch (error) {
    console.error('خطا در دریافت تحلیل پیشرفته:', error)
  }
}

// تابع کمکی برای دریافت رنگ وضعیت
const getStatusColor = (status) => {
  const colors = {
    'awaiting_payment': '#F59E0B',
    'processing': '#3B82F6',
    'shipped': '#10B981',
    'delivered': '#059669',
    'cancelled': '#EF4444',
    'refunded': '#6B7280',
    'pending': '#F59E0B',
    'paid': '#10B981',
    'failed': '#EF4444',
    'completed': '#059669'
  }
  return colors[status] || '#6B7280'
}

// تابع تبدیل تاریخ میلادی به شمسی
const formatPersianDate = (dateString) => {
  if (!dateString) return 'نامشخص'
  
  try {
    const PersianDate = require('persian-date')
    const date = new Date(dateString)
    const persianDate = new PersianDate(date)
    return `${persianDate.year()}/${persianDate.month().toString().padStart(2, '0')}/${persianDate.date().toString().padStart(2, '0')}`
  } catch (error) {
    return 'نامشخص'
  }
}

// تابع تبدیل تاریخ و زمان میلادی به شمسی
const _formatPersianDateTime = (dateString) => {
  if (!dateString) return 'نامشخص'
  
  try {
    const PersianDate = require('persian-date')
    const date = new Date(dateString)
    const persianDate = new PersianDate(date)
    const hours = date.getHours().toString().padStart(2, '0')
    const minutes = date.getMinutes().toString().padStart(2, '0')
    return `${persianDate.year()}/${persianDate.month().toString().padStart(2, '0')}/${persianDate.date().toString().padStart(2, '0')} - ${hours}:${minutes}`
  } catch (error) {
    return 'نامشخص'
  }
}

// تابع کمکی برای دریافت متن وضعیت
const getStatusText = (status) => {
  const statusTexts = {
    'awaiting_payment': 'در انتظار پرداخت',
    'processing': 'در حال پردازش',
    'shipped': 'ارسال شده',
    'delivered': 'تحویل شده',
    'returned': 'مرجوع شده',
    'cancelled': 'لغو شده'
  }
  return statusTexts[status] || status
}

const handleReportItemsPerPageChange = (newItemsPerPage) => {
  reportItemsPerPage.value = newItemsPerPage
  reportCurrentPage.value = 1
}

const _viewDetailedReport = (_report) => {
  // مشاهده گزارش تفصیلی
}

const getConversionRateClass = (rate) => {
  if (rate >= 4) return 'bg-green-100 text-green-800'
  if (rate >= 3) return 'bg-blue-100 text-blue-800'
  if (rate >= 2) return 'bg-yellow-100 text-yellow-800'
  return 'bg-red-100 text-red-800'
}

const getChangeClass = (change) => {
  if (change > 0) return 'bg-green-100 text-green-800'
  if (change < 0) return 'bg-red-100 text-red-800'
  return 'bg-gray-100 text-gray-800'
}
// متغیرهای صفحه‌بندی
const currentPage = ref(1)
const itemsPerPage = ref(10)
const totalItems = computed(() => filteredOrders.value?.length || 0)
const totalPages = computed(() => {
  return Math.max(1, Math.ceil(totalItems.value / itemsPerPage.value))
})

// تب‌های صفحه
const tabs = [
  { id: 'list', name: 'لیست سفارشات' },
  { id: 'accounting', name: 'اتصال حسابداری' },
  { id: 'reports', name: 'گزارشات' }
]
const activeTab = ref('list')

// Watch برای تغییر سال و بارگذاری مجدد گزارشات
watch(selectedYear, (newYear) => {
  if (activeTab.value === 'reports') {
    fetchAnalytics()
    fetchAdvancedAnalytics()
  }
})

// Watch برای تغییر تب و بارگذاری گزارشات
watch(activeTab, (newTab) => {
  if (newTab === 'reports') {
    fetchAnalytics()
    fetchAdvancedAnalytics()
  }
})

// --- دریافت سفارشات از API ادمین ---
const {
  data: ordersData,
  refresh: _refreshOrders,
  pending: ordersPending
} = useAsyncData('admin-orders', () =>
  $fetch('/api/admin/orders?pageSize=10000').then((res: { data: Record<string, unknown>[] })=> res.data || [])
)

// مقداردهی لیست سفارشات فقط از API
const sampleOrders = ref<Order[]>([])
watch(() => ordersData.value, (newData) => {
  if (newData) {
    sampleOrders.value = newData as unknown as Order[]
    currentPage.value = 1
  }
}, { immediate: true })

async function loadOrders() {
  try {
    // حالت لودینگ
    ordersPending.value = true

    // آخرین شناسه موجود در لیست فعلی
    const lastId = sampleOrders.value.length
      ? Math.max(...sampleOrders.value.map((o: { id: number }) => o.id))
      : 0

    // فراخوانی API فقط برای سفارشات جدید
    const res = await $fetch<{ data: Record<string, unknown>[] }>(
      `/api/admin/orders${lastId ? `?afterId=${lastId}` : ''}`
    )

    const newOrders = res.data || []

    if (newOrders.length) {
      // افزودن به ابتدای آرایه (جدیدترین بالا)
      sampleOrders.value = [...(newOrders as unknown as Order[]), ...sampleOrders.value]
      successMessage.value = `${newOrders.length} سفارش جدید اضافه شد`
    } else {
      successMessage.value = 'سفارش جدیدی وجود ندارد'
    }

    // بازگشت به صفحه اول
    currentPage.value = 1

    showSuccessMessage.value = true
    setTimeout(() => {
      showSuccessMessage.value = false
    }, 3000)
  } catch (error) {
    // خطا در بارگذاری سفارشات
    successMessage.value = 'خطا در بارگذاری سفارشات'
    showSuccessMessage.value = true
    setTimeout(() => {
      showSuccessMessage.value = false
    }, 3000)
  } finally {
    ordersPending.value = false
  }
}

// آمار واقعی سفارشات (داده‌های کارت‌ها)
const stats = ref({
  total: 0,
  pending: 0,
  processing: 0,
  shipped: 0,
  delivered: 0,
  cancelled: 0,
  todaySales: 0,
  weeklySales: 0,
  monthlySales: 0,
  averageOrder: 0
})

// دریافت آمار کارت‌ها از API ادمین
const { data: statsData, refresh: _refreshOrderStats } = useAsyncData('admin-orders-stats', () =>
  $fetch('/api/admin/orders/stats')
)

watch(() => statsData.value, (val: Record<string, unknown> | null) => {
  if (val) {
    stats.value = { ...stats.value, ...val }
  }
}, { immediate: true })

// Methods
const _updateStats = (newStats: Partial<typeof stats.value>) => {
  stats.value = { ...stats.value, ...newStats }
}

const formatPrice = (price: number): string => {
  return new Intl.NumberFormat('fa-IR').format(price)
}

// تعریف متغیرهای اضافی
const showFilters = ref(false)
const bulkAction = ref('')
const _filteredOrdersCount = ref(0)

// تعریف متغیر فیلترها
const filters = ref({
  status: '',
  dateFrom: '',
  dateTo: '',
  minAmount: null as number | null,
  maxAmount: null as number | null
})

const handleFiltersChange = (newFilters: typeof filters.value) => {
  filters.value = newFilters
  currentPage.value = 1
}

const openOrderDetail = (order: Record<string, unknown>) => {
  selectedOrder.value = order
  showOrderDetail.value = true
}

const openQuickActions = (order: Record<string, unknown>) => {
  selectedOrder.value = order
  showQuickActions.value = true
}



const _handleOrderUpdate = (_updatedOrder: Record<string, unknown>) => {
  // Update order in the list and refresh stats
  // Order updated
}

const handleOrderCreated = (_newOrder: Record<string, unknown>) => {
  // Add new order to the list and refresh stats
  // Order created
}

const handleStatusChange = (data: { orderId: number, status: string }) => {
  // Status changed
  // Update order status in the list
  const order = sampleOrders.value.find(o => o.id === data.orderId)
  if (order) {
    order.status = data.status
  }
}

const handleNotification = (order: Record<string, unknown>) => {
  // Sending notification for order
}

const handlePrint = (order: Record<string, unknown>) => {
  // Printing order
}

const handleDuplicate = (order: Record<string, unknown>) => {
  // Duplicating order
}

const handleExport = (order: Record<string, unknown>) => {
  // Exporting order
}

const handleCancel = (order: { id: number }) => {
  // Cancelling order
  // Update order status to cancelled
  const orderIndex = sampleOrders.value.findIndex(o => o.id === order.id)
  if (orderIndex !== -1) {
    sampleOrders.value[orderIndex].status = 'cancelled'
  }
}

const handleDelete = async (order: { id: number, orderNumber: string }) => {
  const confirmed = await confirmDialog.value?.show({
    title: 'حذف سفارش',
    message: `آیا از حذف سفارش ${order.orderNumber} مطمئن هستید؟`,
    confirmText: 'حذف',
    cancelText: 'انصراف',
    type: 'danger'
  })
  
  if (confirmed) {
    try {
      await $fetch(`/api/admin/orders/${order.id}`, { method: 'DELETE' })
      
      // حذف از آرایه محلی
      const orderIndex = sampleOrders.value.findIndex(o => o.id === order.id)
      if (orderIndex !== -1) {
        sampleOrders.value.splice(orderIndex, 1)
      }
      
      // نمایش پیام موفقیت
      successMessage.value = 'سفارش با موفقیت حذف شد'
      setTimeout(() => {
        successMessage.value = ''
      }, 3000)
    } catch (error) {
      // console.error('خطا در حذف سفارش:', error)
      alert('خطا در حذف سفارش. لطفاً دوباره تلاش کنید.')
    }
  }
}

const executeBulkAction = () => {
  if (bulkAction.value && selectedOrders.value.length > 0) {
    // Execute the selected bulk action
    // Executing bulk action
    bulkAction.value = ''
  }
}

const exportOrders = () => {
  // Export orders to Excel
  // Exporting orders to Excel
}

const printOrders = () => {
  // Print selected or filtered orders
  // Printing orders
}

  const handlePageChange = (page: number) => {
    currentPage.value = page
    // صفحه‌بندی در سمت کلاینت انجام می‌شود، نیازی به بارگذاری مجدد از سرور نیست
  }

  const handleItemsPerPageChange = (newItemsPerPage: number) => {
    itemsPerPage.value = newItemsPerPage
    currentPage.value = 1
  }

async function handleBulkDelete(ids: number[]) {
  if (!ids || ids.length === 0) return;
  
  const confirmed = await confirmDialog.value?.show({
    title: 'حذف چندگانه',
    message: `آیا از حذف ${ids.length} سفارش انتخاب شده مطمئن هستید؟`,
    confirmText: 'حذف همه',
    cancelText: 'انصراف',
    type: 'danger'
  })
  
  if (confirmed) {
    try {
      // حذف از دیتابیس
      for (const id of ids) {
        await $fetch(`/api/admin/orders/${id}`, { method: 'DELETE' })
      }
      
      // حذف از آرایه محلی
      sampleOrders.value = sampleOrders.value.filter(order => !ids.includes(order.id));
      
      // نمایش پیام موفقیت
      successMessage.value = `${ids.length} سفارش با موفقیت حذف شد`
      setTimeout(() => {
        successMessage.value = ''
      }, 3000)
    } catch (error) {
      // console.error('خطا در حذف سفارشات:', error)
      alert('خطا در حذف سفارشات. لطفاً دوباره تلاش کنید.')
    }
  }
}

// فیلتر کردن سفارشات بر اساس جستجو و فیلترها
const filteredOrders = computed(() => {
  let filtered = sampleOrders.value || []

  // فیلتر بر اساس جستجو
  if (searchTerm.value) {
    const search = searchTerm.value.toLowerCase()
    filtered = filtered.filter(order => 
      order.orderNumber?.toLowerCase().includes(search) ||
      order.customerName?.toLowerCase().includes(search) ||
      order.trackingCode?.toLowerCase().includes(search) ||
      order.customerPhone?.includes(search)
    )
  }

  // فیلتر بر اساس وضعیت سفارش
  if (filters.value.status && filters.value.status !== 'all') {
    filtered = filtered.filter(order => order.status === filters.value.status)
  }

  // فیلتر بر اساس تاریخ
  if (filters.value.dateFrom) {
    const fromDate = new Date(filters.value.dateFrom)
    filtered = filtered.filter(order => new Date(order.createdAt) >= fromDate)
  }

  if (filters.value.dateTo) {
    const toDate = new Date(filters.value.dateTo)
    filtered = filtered.filter(order => new Date(order.createdAt) <= toDate)
  }

  // فیلتر بر اساس مبلغ
  if (filters.value.minAmount) {
    filtered = filtered.filter(order => order.totalAmount >= filters.value.minAmount)
  }

  if (filters.value.maxAmount) {
    filtered = filtered.filter(order => order.totalAmount <= filters.value.maxAmount)
  }

  return filtered
})

// سفارشات صفحه‌بندی شده
const paginatedOrders = computed(() => {
  if (!filteredOrders.value || filteredOrders.value.length === 0) {
    return []
  }
  
  const start = (currentPage.value - 1) * itemsPerPage.value
  const end = start + itemsPerPage.value
  
  return filteredOrders.value.slice(start, end)
})

// بازگشت به صفحه اول هنگام تغییر فیلترها
watch([filters, searchTerm], () => {
  currentPage.value = 1
})
</script>

<style scoped>
/* Additional styles if needed */
</style>
