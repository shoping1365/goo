<template>
  <ClientOnly>
    <!-- Page Header -->
    <div class="bg-white shadow-sm border-b border-gray-200 mb-4">
      <div class="w-full px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center py-3">
        <div>
            <h1 class="text-lg font-bold text-gray-900">سفارشات کنسل شده</h1>
            <p class="text-xs text-gray-600 mt-1">مدیریت و بررسی سفارشات لغو شده</p>
        </div>
          <div class="flex space-x-2 space-x-reverse">
          <button
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-green-400 to-green-600 hover:from-green-500 hover:to-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105"
              @click="exportData"
          >
            <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
            </svg>
              خروجی Excel
          </button>
            <button 
              class="inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-lg text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 shadow-md transition-all duration-200"
              @click="refreshData"
            >
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
              </svg>
            بروزرسانی
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
              :class="[
                activeTab === tab.id
                  ? 'border-red-500 text-red-600'
                  : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300',
              'flex-1 whitespace-nowrap py-4 px-4 border-b-2 font-medium text-sm text-center'
              ]"
              @click="activeTab = tab.id"
            >
              {{ tab.name }}
            </button>
          </nav>
        </div>

        <div class="px-4 py-4">
          <!-- تب داشبورد -->
          <div v-if="activeTab === 'dashboard'">
            <!-- آمار کلی سفارشات کنسل شده -->
            <div class="rounded-lg shadow-md border border-gray-200 overflow-hidden mb-6">
              <div class="bg-gradient-to-r from-red-50 to-pink-50 px-4 py-3 border-b border-gray-200">
        <div class="flex items-center">
                  <div class="w-6 h-6 bg-red-500 rounded-md flex items-center justify-center ml-2">
                    <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                    </svg>
          </div>
                  <h3 class="text-sm font-semibold text-gray-900">آمار کلی سفارشات کنسل شده</h3>
                </div>
              </div>
              
              <div class="px-4 py-4">
                <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-3">
                  <!-- کل سفارشات کنسل شده -->
                  <div class="bg-gradient-to-br from-red-50 to-red-100 border border-red-200 rounded-md p-3 hover:shadow-sm transition-all duration-200">
                    <div class="flex items-center justify-between">
                      <div>
                        <p class="text-xs font-medium text-red-700 mb-0.5">کل سفارشات کنسل شده</p>
                        <p class="text-lg font-bold text-red-900">{{ stats.totalCancelled?.toLocaleString('fa-IR') || '0' }}</p>
                      </div>
                      <div class="w-8 h-8 bg-red-500 rounded-md flex items-center justify-center">
                        <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                        </svg>
          </div>
        </div>
      </div>

                  <!-- مجموع مبلغ کنسل شده -->
                  <div class="bg-gradient-to-br from-orange-50 to-orange-100 border border-orange-200 rounded-md p-3 hover:shadow-sm transition-all duration-200">
                    <div class="flex items-center justify-between">
                      <div>
                        <p class="text-xs font-medium text-orange-700 mb-0.5">مجموع مبلغ کنسل شده</p>
                        <p class="text-lg font-bold text-orange-900">{{ formatPrice(stats.totalAmount || 0) }}</p>
          </div>
                      <div class="w-8 h-8 bg-orange-500 rounded-md flex items-center justify-center">
                        <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
                        </svg>
          </div>
        </div>
      </div>

                  <!-- مشتریان متاثر -->
                  <div class="bg-gradient-to-br from-purple-50 to-purple-100 border border-purple-200 rounded-md p-3 hover:shadow-sm transition-all duration-200">
                    <div class="flex items-center justify-between">
                      <div>
                        <p class="text-xs font-medium text-purple-700 mb-0.5">مشتریان متاثر</p>
                        <p class="text-lg font-bold text-purple-900">{{ stats.affectedCustomers?.toLocaleString('fa-IR') || '0' }}</p>
          </div>
                      <div class="w-8 h-8 bg-purple-500 rounded-md flex items-center justify-center">
                        <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
                        </svg>
          </div>
        </div>
      </div>

                  <!-- نرخ کنسل -->
                  <div class="bg-gradient-to-br from-blue-50 to-blue-100 border border-blue-200 rounded-md p-3 hover:shadow-sm transition-all duration-200">
                    <div class="flex items-center justify-between">
                      <div>
                        <p class="text-xs font-medium text-blue-700 mb-0.5">نرخ کنسل</p>
                        <p class="text-lg font-bold text-blue-900">{{ stats.cancellationRate || 0 }}%</p>
          </div>
                      <div class="w-8 h-8 bg-blue-500 rounded-md flex items-center justify-center">
                        <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
                        </svg>
          </div>
        </div>
      </div>
    </div>
      </div>
    </div>



            <!-- آمار کنسل -->
            <div class="grid grid-cols-1 lg:grid-cols-2 gapx-4 py-4 mb-6">
              <!-- آمار کلی کنسل -->
              <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
                <h3 class="text-lg font-semibold text-gray-900 mb-4">آمار کلی کنسل</h3>
                <div class="space-y-4">
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-700">کل سفارشات بررسی شده</span>
                    <span class="text-sm font-medium">۲۳,۴۵۶</span>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-700">سفارشات کنسل شده</span>
                    <span class="text-sm font-medium text-red-600">۱,۲۳۴</span>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-700">نرخ کنسل</span>
                    <span class="text-sm font-medium text-red-600">۵.۳%</span>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-700">متوسط زمان کنسل</span>
                    <span class="text-sm font-medium text-blue-600">۲.۳ ساعت</span>
                  </div>
                </div>
              </div>
              
              <!-- دلایل کنسل -->
              <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
                <h3 class="text-lg font-semibold text-gray-900 mb-4">دلایل کنسل</h3>
                <div class="space-y-4">
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-700">درخواست مشتری</span>
                    <div class="flex items-center">
                      <div class="w-16 bg-gray-200 rounded-full h-2 ml-2">
                        <div class="bg-red-600 h-2 rounded-full" style="width: 45%"></div>
                      </div>
                      <span class="text-sm font-medium">۴۵%</span>
                    </div>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-700">عدم موجودی</span>
                    <div class="flex items-center">
                      <div class="w-16 bg-gray-200 rounded-full h-2 ml-2">
                        <div class="bg-orange-600 h-2 rounded-full" style="width: 28%"></div>
                      </div>
                      <span class="text-sm font-medium">۲۸%</span>
                    </div>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-700">مشکل پرداخت</span>
                    <div class="flex items-center">
                      <div class="w-16 bg-gray-200 rounded-full h-2 ml-2">
                        <div class="bg-yellow-600 h-2 rounded-full" style="width: 18%"></div>
                      </div>
                      <span class="text-sm font-medium">۱۸%</span>
                    </div>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-700">مشکل سیستمی</span>
                    <div class="flex items-center">
                      <div class="w-16 bg-gray-200 rounded-full h-2 ml-2">
                        <div class="bg-gray-600 h-2 rounded-full" style="width: 9%"></div>
                      </div>
                      <span class="text-sm font-medium">۹%</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <!-- آمار انواع کنسل -->
            <div class="grid grid-cols-1 lg:grid-cols-2 gapx-4 py-4">
              <!-- نمودار انواع کنسل -->
              <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
                <h3 class="text-lg font-semibold text-gray-900 mb-4">توزیع انواع کنسل</h3>
                <div class="space-y-4">
                  <div v-for="(count, type) in cancellationTypeStats" :key="type" class="flex items-center justify-between">
                    <div class="flex items-center">
                      <div :class="getReasonClass(type)" class="w-3 h-3 rounded-full ml-3"></div>
                      <span class="text-sm text-gray-700">{{ getReasonText(type) }}</span>
                    </div>
                    <div class="flex items-center">
                      <div class="w-24 bg-gray-200 rounded-full h-2 ml-3">
                        <div :class="getReasonClass(type).replace('text-', 'bg-').replace('-100', '-500')" class="h-2 rounded-full" :style="{ width: (count / cancelledOrders.length * 100) + '%' }"></div>
                      </div>
                      <span class="text-sm font-medium text-gray-900">{{ count }}</span>
                    </div>
                  </div>
                </div>
              </div>

              <!-- آمار روش‌های پرداخت -->
              <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
                <h3 class="text-lg font-semibold text-gray-900 mb-4">توزیع روش‌های پرداخت</h3>
                <div class="space-y-4">
                  <div v-for="(count, method) in paymentMethodStats" :key="method" class="flex items-center justify-between">
                    <div class="flex items-center">
                      <div :class="getPaymentMethodClass(method)" class="w-3 h-3 rounded-full ml-3"></div>
                      <span class="text-sm text-gray-700">{{ getPaymentMethodText(method) }}</span>
                    </div>
                    <div class="flex items-center">
                      <div class="w-24 bg-gray-200 rounded-full h-2 ml-3">
                        <div :class="getPaymentMethodClass(method).replace('text-', 'bg-').replace('-100', '-500')" class="h-2 rounded-full" :style="{ width: (count / cancelledOrders.length * 100) + '%' }"></div>
                      </div>
                      <span class="text-sm font-medium text-gray-900">{{ count }}</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <!-- نمودار کنسل -->
            <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4 mt-6">
              <h3 class="text-lg font-semibold text-gray-900 mb-4">نمودار کنسل سفارشات</h3>
              <div class="h-64 flex items-center justify-center bg-gray-50 rounded-lg">
                <p class="text-gray-500">نمودار کنسل سفارشات</p>
              </div>
            </div>

            <!-- جدول سفارشات کنسل شده اخیر -->
            <div class="bg-white rounded-lg shadow-md border border-gray-200 mt-6">
              <div class="px-4 py-4 border-b border-gray-200">
                <div class="flex items-center justify-between">
                  <h3 class="text-lg font-semibold text-gray-900">سفارشات کنسل شده اخیر</h3>
                </div>
              </div>

              <div class="overflow-x-auto">
                <table class="min-w-full divide-y divide-gray-200">
                  <thead class="bg-gray-50">
                    <tr>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">شماره سفارش</th>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مشتری</th>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مبلغ</th>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">دلیل کنسل</th>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ کنسل</th>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
                    </tr>
                  </thead>
                  <tbody class="bg-white divide-y divide-gray-200">
                    <tr v-for="order in paginatedCancelledOrders.slice(0, 5)" :key="order.id" class="hover:bg-gray-50">
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
                        {{ formatPrice(order.totalAmount || 0) }}
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap">
                        <span :class="getReasonClass(order.cancellationReason)" class="px-2 py-1 text-xs font-medium rounded-full">
                          {{ getReasonText(order.cancellationReason) }}
                        </span>
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                        {{ formatDate(order.cancelledAt) }}
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                        <div class="flex items-center space-x-2">
                          <button class="text-blue-600 hover:text-blue-900" @click="viewOrderDetails(order)">
                            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                            </svg>
                          </button>
                          <button class="text-green-600 hover:text-green-900">بازگردانی</button>
                        </div>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>
          </div>

          <!-- تب سفارشات کنسل شده -->
          <div v-if="activeTab === 'cancelled-orders'">
            <!-- آمار کلی سفارشات کنسل شده -->
            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gapx-4 py-4 mb-8">
              <!-- کل سفارشات کنسل شده -->
              <div class="bg-gradient-to-br from-red-50 to-red-100 border border-red-200 rounded-md p-3 hover:shadow-sm transition-all duration-200">
                <div class="flex items-center justify-between">
                  <div>
                    <p class="text-xs font-medium text-red-700 mb-0.5">کل سفارشات کنسل شده</p>
                    <p class="text-lg font-bold text-red-900">{{ stats.totalCancelled?.toLocaleString('fa-IR') || '0' }}</p>
                  </div>
                  <div class="w-8 h-8 bg-red-500 rounded-md flex items-center justify-center">
                    <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                    </svg>
                  </div>
                </div>
                <p class="text-xs text-red-600 mt-1">+۸.۵% از ماه گذشته</p>
              </div>

              <!-- مجموع مبلغ کنسل شده -->
              <div class="bg-gradient-to-br from-orange-50 to-orange-100 border border-orange-200 rounded-md p-3 hover:shadow-sm transition-all duration-200">
                <div class="flex items-center justify-between">
                  <div>
                    <p class="text-xs font-medium text-orange-700 mb-0.5">مجموع مبلغ کنسل شده</p>
                    <p class="text-lg font-bold text-orange-900">{{ formatPrice(stats.totalAmount || 0) }}</p>
                  </div>
                  <div class="w-8 h-8 bg-orange-500 rounded-md flex items-center justify-center">
                    <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
                    </svg>
                  </div>
                </div>
                <p class="text-xs text-orange-600 mt-1">+۱۲.۳% از ماه گذشته</p>
              </div>

              <!-- مشتریان متاثر -->
              <div class="bg-gradient-to-br from-purple-50 to-purple-100 border border-purple-200 rounded-md p-3 hover:shadow-sm transition-all duration-200">
                <div class="flex items-center justify-between">
                  <div>
                    <p class="text-xs font-medium text-purple-700 mb-0.5">مشتریان متاثر</p>
                    <p class="text-lg font-bold text-purple-900">{{ stats.affectedCustomers?.toLocaleString('fa-IR') || '0' }}</p>
                  </div>
                  <div class="w-8 h-8 bg-purple-500 rounded-md flex items-center justify-center">
                    <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
                    </svg>
                  </div>
                </div>
                <p class="text-xs text-purple-600 mt-1">+۵.۷% از ماه گذشته</p>
              </div>

              <!-- نرخ کنسل -->
              <div class="bg-gradient-to-br from-blue-50 to-blue-100 border border-blue-200 rounded-md p-3 hover:shadow-sm transition-all duration-200">
                <div class="flex items-center justify-between">
                  <div>
                    <p class="text-xs font-medium text-blue-700 mb-0.5">نرخ کنسل</p>
                    <p class="text-lg font-bold text-blue-900">{{ stats.cancellationRate || 0 }}%</p>
                  </div>
                  <div class="w-8 h-8 bg-blue-500 rounded-md flex items-center justify-center">
                    <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
                    </svg>
                  </div>
                </div>
                <p class="text-xs text-blue-600 mt-1">-۲.۱% از ماه گذشته</p>
              </div>
            </div>

            <!-- فیلترهای پیشرفته -->
            <div class="bg-gradient-to-br from-gray-50 to-gray-100 rounded-xl shadow-lg border border-gray-200 px-4 py-4 mb-6 w-full">
              <!-- Advanced Filters -->
              <div class="flex items-center justify-between">
                <div class="flex items-center">
                  <button class="w-8 h-8 bg-gradient-to-r from-red-500 to-pink-600 rounded-lg flex items-center justify-center ml-3 text-white hover:bg-red-600 transition-colors" @click="showAdvanced = !showAdvanced">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.207A1 1 0 013 6.5V4z"></path>
                    </svg>
                  </button>
                  <h3 class="text-lg font-semibold text-gray-900">فیلترهای سفارشات کنسل شده</h3>
                </div>
                <span class="text-sm text-gray-600 cursor-pointer hover:text-red-600 transition-colors" @click="showAdvanced = !showAdvanced">
                  {{ showAdvanced ? 'مخفی کردن' : 'نمایش' }}
                </span>
              </div>

              <div v-show="showAdvanced" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gapx-4 py-4">
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

                <!-- Cancellation Reason -->
                <div class="space-y-2">
                  <label class="block text-sm font-semibold text-gray-700">دلیل کنسل</label>
                  <select 
                    v-model="reasonFilter" 
                    class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-red-500 focus:border-red-500 transition-all duration-200 bg-white shadow-sm"
                  >
                    <option value="">همه دلایل</option>
                    <option value="customer">درخواست مشتری</option>
                    <option value="inventory">عدم موجودی</option>
                    <option value="payment">مشکل پرداخت</option>
                    <option value="system">مشکل سیستمی</option>
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
            </div>

            <!-- نمودارها و آمار -->
            <div class="grid grid-cols-1 lg:grid-cols-2 gapx-4 py-4 mb-6">
              <!-- نمودار دلایل کنسل -->
              <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
                <h3 class="text-lg font-semibold text-gray-900 mb-4">توزیع دلایل کنسل</h3>
                <div class="space-y-4">
                  <div class="flex items-center justify-between">
                    <div class="flex items-center">
                      <div class="w-3 h-3 bg-red-500 rounded-full ml-3"></div>
                      <span class="text-sm text-gray-700">درخواست مشتری</span>
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
                      <span class="text-sm text-gray-700">عدم موجودی</span>
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
                      <div class="w-3 h-3 bg-yellow-500 rounded-full ml-3"></div>
                      <span class="text-sm text-gray-700">مشکل پرداخت</span>
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
                      <div class="w-3 h-3 bg-gray-500 rounded-full ml-3"></div>
                      <span class="text-sm text-gray-700">مشکل سیستمی</span>
                    </div>
                    <div class="flex items-center">
                      <div class="w-24 bg-gray-200 rounded-full h-2 ml-3">
                        <div class="bg-gray-500 h-2 rounded-full" style="width: 9%"></div>
                      </div>
                      <span class="text-sm font-medium text-gray-900">۹%</span>
                    </div>
                  </div>
                </div>
              </div>

              <!-- آمار زمانی -->
              <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
                <h3 class="text-lg font-semibold text-gray-900 mb-4">آمار زمانی کنسل</h3>
                <div class="space-y-4">
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-700">سریع‌ترین کنسل</span>
                    <span class="text-sm font-medium text-green-600">۱۵ دقیقه</span>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-700">متوسط زمان کنسل</span>
                    <span class="text-sm font-medium text-blue-600">۲.۳ ساعت</span>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-700">دیرترین کنسل</span>
                    <span class="text-sm font-medium text-red-600">۲۴ ساعت</span>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-700">تعداد کنسل در روز</span>
                    <span class="text-sm font-medium text-purple-600">۱۲۳</span>
                  </div>
                </div>
              </div>
            </div>

            <!-- جدول سفارشات کنسل شده -->
            <div class="bg-white rounded-lg shadow-md border border-gray-200">
              <div class="px-4 py-4 border-b border-gray-200">
                <div class="flex items-center justify-between">
                  <h3 class="text-lg font-semibold text-gray-900">سفارشات کنسل شده</h3>
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <button class="bg-red-600 hover:bg-red-700 text-white px-4 py-2 rounded-lg text-sm font-medium transition-colors">
                      بازگردانی همه
                    </button>
                    <button class="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg text-sm font-medium transition-colors">
                      خروجی Excel
                    </button>
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
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">دلیل کنسل</th>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">روش پرداخت</th>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ کنسل</th>
                      <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
                    </tr>
                  </thead>
                  <tbody class="bg-white divide-y divide-gray-200">
                    <tr v-for="order in paginatedCancelledOrders" :key="order.id" class="hover:bg-gray-50">
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
                        {{ formatPrice(order.totalAmount || 0) }}
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap">
                        <span :class="getReasonClass(order.cancellationReason)" class="px-2 py-1 text-xs font-medium rounded-full">
                          {{ getReasonText(order.cancellationReason) }}
                        </span>
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap">
                        <span :class="getPaymentMethodClass(order.paymentMethod)" class="px-2 py-1 text-xs font-medium rounded-full">
                          {{ getPaymentMethodText(order.paymentMethod) }}
                        </span>
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                        {{ formatDate(order.cancelledAt) }}
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                        <div class="flex items-center space-x-2">
                          <button class="text-blue-600 hover:text-blue-900" @click="viewOrderDetails(order)">
                            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                            </svg>
                          </button>
                          <button class="text-green-600 hover:text-green-900">بازگردانی</button>
                        </div>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>

              <!-- صفحه‌بندی -->
              <div v-if="cancelledOrders.length > 0">
                <Pagination 
                  :current-page="currentPage"
                  :total-pages="totalPages"
                  :total="cancelledOrders.length"
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
            <div class="grid grid-cols-1 lg:grid-cols-2 gapx-4 py-4 mb-6">
              <!-- نمودار روند کنسل‌ها -->
              <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
                <h3 class="text-lg font-semibold text-gray-900 mb-4">روند کنسل‌ها در ۶ ماه گذشته</h3>
                <div class="space-y-4">
        <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">ماه اول</span>
                    <div class="flex items-center">
                      <div class="w-32 bg-gray-200 rounded-full h-3 ml-3">
                        <div class="bg-red-500 h-3 rounded-full" style="width: 35%"></div>
                      </div>
                      <span class="text-sm font-medium text-gray-900">۳۵</span>
                    </div>
                  </div>
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">ماه دوم</span>
                    <div class="flex items-center">
                      <div class="w-32 bg-gray-200 rounded-full h-3 ml-3">
                        <div class="bg-red-500 h-3 rounded-full" style="width: 42%"></div>
                      </div>
                      <span class="text-sm font-medium text-gray-900">۴۲</span>
                    </div>
                  </div>
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">ماه سوم</span>
                    <div class="flex items-center">
                      <div class="w-32 bg-gray-200 rounded-full h-3 ml-3">
                        <div class="bg-red-500 h-3 rounded-full" style="width: 28%"></div>
                      </div>
                      <span class="text-sm font-medium text-gray-900">۲۸</span>
                    </div>
                  </div>
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">ماه چهارم</span>
                    <div class="flex items-center">
                      <div class="w-32 bg-gray-200 rounded-full h-3 ml-3">
                        <div class="bg-red-500 h-3 rounded-full" style="width: 51%"></div>
                      </div>
                      <span class="text-sm font-medium text-gray-900">۵۱</span>
                    </div>
                  </div>
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">ماه پنجم</span>
                    <div class="flex items-center">
                      <div class="w-32 bg-gray-200 rounded-full h-3 ml-3">
                        <div class="bg-red-500 h-3 rounded-full" style="width: 33%"></div>
                      </div>
                      <span class="text-sm font-medium text-gray-900">۳۳</span>
                    </div>
                  </div>
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-600">ماه ششم</span>
                    <div class="flex items-center">
                      <div class="w-32 bg-gray-200 rounded-full h-3 ml-3">
                        <div class="bg-red-500 h-3 rounded-full" style="width: 45%"></div>
                      </div>
                      <span class="text-sm font-medium text-gray-900">۴۵</span>
                    </div>
                  </div>
                </div>
              </div>

              <!-- نمودار توزیع دلایل کنسل -->
              <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
                <h3 class="text-lg font-semibold text-gray-900 mb-4">توزیع دلایل کنسل</h3>
                <div class="space-y-4">
                  <div class="flex items-center justify-between">
                    <div class="flex items-center">
                      <div class="w-3 h-3 bg-red-500 rounded-full ml-3"></div>
                      <span class="text-sm text-gray-700">درخواست مشتری</span>
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
                      <span class="text-sm text-gray-700">عدم موجودی</span>
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
                      <div class="w-3 h-3 bg-yellow-500 rounded-full ml-3"></div>
                      <span class="text-sm text-gray-700">مشکل پرداخت</span>
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
                      <div class="w-3 h-3 bg-gray-500 rounded-full ml-3"></div>
                      <span class="text-sm text-gray-700">مشکل سیستمی</span>
                    </div>
                    <div class="flex items-center">
                      <div class="w-24 bg-gray-200 rounded-full h-2 ml-3">
                        <div class="bg-gray-500 h-2 rounded-full" style="width: 9%"></div>
                      </div>
                      <span class="text-sm font-medium text-gray-900">۹%</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <!-- فیلترهای گزارش -->
            <div class="bg-gradient-to-br from-gray-50 to-gray-100 rounded-xl shadow-lg border border-gray-200 px-4 py-4 mb-6">
              <div class="flex items-center justify-between mb-4">
                <div class="flex items-center">
                  <button class="w-8 h-8 bg-gradient-to-r from-red-500 to-pink-600 rounded-lg flex items-center justify-center ml-3 text-white hover:bg-red-600 transition-colors" @click="showReportFilters = !showReportFilters">
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
                <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gapx-4 py-4">
                  <!-- نوع گزارش -->
                  <div class="space-y-1">
                    <label class="block text-sm font-semibold text-gray-700">نوع گزارش</label>
                    <select v-model="reportType" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500 transition-all duration-200 bg-white shadow-sm">
                      <option value="comprehensive">گزارش جامع</option>
                      <option value="cancellation_summary">خلاصه کنسل‌ها</option>
                      <option value="customer_analysis">تحلیل مشتریان</option>
                      <option value="financial_impact">تأثیر مالی</option>
                      <option value="trend_analysis">تحلیل روند</option>
                      <option value="reason_analysis">تحلیل دلایل</option>
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
                <div class="grid grid-cols-1 md:grid-cols-4 gapx-4 py-4 mt-4">
                  <!-- دلیل کنسل -->
                  <div class="space-y-1">
                    <label class="block text-sm font-semibold text-gray-700">دلیل کنسل</label>
                    <select v-model="reportCancellationReason" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500 transition-all duration-200 bg-white shadow-sm">
                      <option value="">همه دلایل</option>
                      <option value="customer">درخواست مشتری</option>
                      <option value="inventory">عدم موجودی</option>
                      <option value="payment">مشکل پرداخت</option>
                      <option value="system">مشکل سیستمی</option>
                    </select>
                  </div>

                  <!-- روش پرداخت -->
                  <div class="space-y-1">
                    <label class="block text-sm font-semibold text-gray-700">روش پرداخت</label>
                    <select v-model="reportPaymentMethod" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500 transition-all duration-200 bg-white shadow-sm">
                      <option value="">همه روش‌ها</option>
                      <option value="online">آنلاین</option>
                      <option value="cash">نقدی</option>
                      <option value="bank_transfer">انتقال بانکی</option>
                      <option value="wallet">کیف پول</option>
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
                    <button class="w-full bg-gradient-to-r from-red-500 to-pink-600 text-white px-3 py-2 rounded-lg font-medium hover:from-red-600 hover:to-pink-700 transition-all duration-200 flex items-center justify-center" @click="generateReport">
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
            <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4 mb-6">
              <h3 class="text-lg font-semibold text-gray-900 mb-4">گزارشات سریع</h3>
              <div class="grid grid-cols-1 md:grid-cols-3 gapx-4 py-4">
                <div class="flex items-center justify-between px-4 py-4 bg-gradient-to-r from-blue-50 to-blue-100 rounded-lg border border-blue-200">
            <div>
                    <h4 class="text-sm font-medium text-blue-900">گزارش ماهانه</h4>
                    <p class="text-xs text-blue-700">آخرین ۳۰ روز</p>
            </div>
                  <div class="flex space-x-2">
                    <button class="bg-blue-600 hover:bg-blue-700 text-white px-3 py-2 rounded-lg text-xs font-medium transition-colors" @click="downloadQuickReport('monthly', 'pdf')">
                      PDF
                    </button>
                    <button class="bg-green-600 hover:bg-green-700 text-white px-3 py-2 rounded-lg text-xs font-medium transition-colors" @click="downloadQuickReport('monthly', 'excel')">
                      Excel
                    </button>
                  </div>
                </div>

                <div class="flex items-center justify-between px-4 py-4 bg-gradient-to-r from-green-50 to-green-100 rounded-lg border border-green-200">
            <div>
                    <h4 class="text-sm font-medium text-green-900">گزارش سه‌ماهه</h4>
                    <p class="text-xs text-green-700">آخرین ۹۰ روز</p>
                  </div>
                  <div class="flex space-x-2">
                    <button class="bg-blue-600 hover:bg-blue-700 text-white px-3 py-2 rounded-lg text-xs font-medium transition-colors" @click="downloadQuickReport('quarterly', 'pdf')">
                      PDF
                </button>
                    <button class="bg-green-600 hover:bg-green-700 text-white px-3 py-2 rounded-lg text-xs font-medium transition-colors" @click="downloadQuickReport('quarterly', 'excel')">
                      Excel
                </button>
            </div>
          </div>

                <div class="flex items-center justify-between px-4 py-4 bg-gradient-to-r from-purple-50 to-purple-100 rounded-lg border border-purple-200">
                  <div>
                    <h4 class="text-sm font-medium text-purple-900">گزارش سالانه</h4>
                    <p class="text-xs text-purple-700">آخرین ۳۶۵ روز</p>
        </div>
                  <div class="flex space-x-2">
                    <button class="bg-blue-600 hover:bg-blue-700 text-white px-3 py-2 rounded-lg text-xs font-medium transition-colors" @click="downloadQuickReport('yearly', 'pdf')">
                      PDF
                    </button>
                    <button class="bg-green-600 hover:bg-green-700 text-white px-3 py-2 rounded-lg text-xs font-medium transition-colors" @click="downloadQuickReport('yearly', 'excel')">
                      Excel
                    </button>
      </div>
    </div>
  </div>
            </div>

            <!-- تاریخچه گزارشات -->
            <div class="bg-white rounded-lg shadow-md border border-gray-200">
              <div class="px-4 py-4 border-b border-gray-200">
                <div class="flex items-center justify-between">
                  <h3 class="text-lg font-semibold text-gray-900">تاریخچه گزارشات</h3>
                  <button class="bg-gray-600 hover:bg-gray-700 text-white px-4 py-2 rounded-lg text-sm font-medium transition-colors" @click="refreshReportHistory">
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


            </div>
          </div>
        </div>
      </div>

      <!-- مودال جزئیات سفارش -->
      <OrderDetailsModal 
        :is-open="isModalOpen"
        :order="selectedOrder"
        @close="closeModal"
        @edit="editOrder"
      />
    </ClientOnly>
</template>

<script setup>
definePageMeta({
  layout: 'admin-main'
})

import { computed, ref } from 'vue'



// Import کامپوننت‌ها
import Pagination from '~/components/admin/common/Pagination.vue'
import OrderDetailsModal from '~/components/admin/modals/OrderDetailsModal.vue'

// تب‌های صفحه
const tabs = ref([
  { id: 'dashboard', name: 'داشبورد' },
  { id: 'cancelled-orders', name: 'سفارشات کنسل شده' },
  { id: 'reports', name: 'گزارشات' }
])

const activeTab = ref('dashboard')

// داده‌های واقعی از API
const stats = ref({})

// متغیرهای مدیریت وضعیت بارگذاری
const loading = ref(false)
const error = ref(null)

// تابع دریافت آمار کنسل
const fetchStats = async () => {
  try {
    loading.value = true
    error.value = null

    const { data } = await $fetch('/api/admin/orders/cancelled/stats', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    })

    // تبدیل داده‌های API به فرمت مورد نیاز کامپوننت
    stats.value = {
      totalCancelled: data.totalCancelled || 0,
      totalAmount: (data.todaySales || 0) + (data.weeklySales || 0) + (data.monthlySales || 0),
      affectedCustomers: data.totalCancelled || 0, // تعداد سفارشات کنسل شده = تعداد مشتریان متاثر
      cancellationRate: data.averageOrder || 0 // TODO: محاسبه واقعی نرخ کنسل
    }

  } catch (err) {
    console.error('خطا در دریافت آمار:', err)
    error.value = err.message || 'خطا در دریافت آمار'
  } finally {
    loading.value = false
  }
}

// مقداردهی اولیه و دریافت داده‌ها
onMounted(() => {
  fetchStats()
})

const searchQuery = ref('')
const reasonFilter = ref('')
const dateFilter = ref('')
const paymentMethodFilter = ref('')
const minAmountFilter = ref('')
const maxAmountFilter = ref('')
const dateFromFilter = ref('')
const dateToFilter = ref('')
const showAdvanced = ref(false)

// متغیرهای مودال
const isModalOpen = ref(false)
const selectedOrder = ref(null)

// متغیرهای گزارشات
const showReportFilters = ref(true)
const reportDateFrom = ref('')
const reportDateTo = ref('')
const reportFormat = ref('pdf')
const reportType = ref('comprehensive')
const reportPeriod = ref('last_30_days')
const reportCancellationReason = ref('')
const reportPaymentMethod = ref('')
const reportCurrentPage = ref(1)
const reportItemsPerPage = ref(5)
const reportHistory = ref([
  {
    id: 1,
    name: 'گزارش جامع کنسل ماه گذشته',
    type: 'جامع',
    period: '۳۰ روز گذشته',
    format: 'PDF',
    size: '۲.۴ MB',
    date: '۲ ساعت پیش'
  },
  {
    id: 2,
    name: 'تحلیل دلایل کنسل',
    type: 'تحلیل دلایل',
    period: '۹۰ روز گذشته',
    format: 'Excel',
    size: '۱.۸ MB',
    date: '۱ روز پیش'
  },
  {
    id: 3,
    name: 'آمار کنسل هفتگی',
    type: 'خلاصه کنسل‌ها',
    period: '۷ روز گذشته',
    format: 'CSV',
    size: '۰.۵ MB',
    date: '۳ روز پیش'
  },
  {
    id: 4,
    name: 'گزارش مشتریان متاثر',
    type: 'تحلیل مشتریان',
    period: '۶۰ روز گذشته',
    format: 'Excel',
    size: '۳.۱ MB',
    date: '۱ هفته پیش'
  },
  {
    id: 5,
    name: 'تحلیل روند کنسل',
    type: 'تحلیل روند',
    period: '۱۸۰ روز گذشته',
    format: 'PDF',
    size: '۲.۷ MB',
    date: '۲ هفته پیش'
  },
  {
    id: 6,
    name: 'گزارش مالی کنسل',
    type: 'تأثیر مالی',
    period: '۳۶۵ روز گذشته',
    format: 'Excel',
    size: '۱.۹ MB',
    date: '۱ ماه پیش'
  }
])

// داده‌های واقعی از API
const cancelledOrders = ref([])

// تابع دریافت لیست سفارشات کنسل شده
const fetchCancelledOrders = async () => {
  try {
    loading.value = true
    error.value = null

    const { data } = await $fetch('/api/admin/orders/cancelled', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    })

    cancelledOrders.value = (data || []).map(order => ({
      id: order.id || 0,
      orderNumber: order.orderNumber || 'نامشخص',
      customerName: order.customerName || 'نامشخص',
      customerPhone: order.customerPhone || 'نامشخص',
      totalAmount: order.totalAmount || 0,
      paymentMethod: order.paymentMethod || 'نامشخص',
      status: order.status || 'نامشخص',
      orderIntegrity: order.orderIntegrity || 'نامشخص',
      createdAt: order.createdAt || null,
      cancelledAt: order.createdAt || null, // TODO: اضافه کردن فیلد cancelledAt به مدل Order
      cancellationReason: 'نامشخص', // TODO: اضافه کردن فیلد cancellationReason به مدل Order
      items: [] // TODO: دریافت آیتم‌های سفارش
    }))

  } catch (err) {
    console.error('خطا در دریافت سفارشات کنسل شده:', err)
    error.value = err.message || 'خطا در دریافت سفارشات کنسل شده'
  } finally {
    loading.value = false
  }
}

// مقداردهی اولیه و دریافت داده‌ها
onMounted(() => {
  fetchStats()
  fetchCancelledOrders()
})

// داده‌های قدیمی برای سازگاری
const orders = ref([
  {
    id: 1,
    orderNumber: 'ORD-001',
    customerName: 'علی احمدی',
    cancelledAt: '2024-01-15',
    totalAmount: 1250000,
    cancellationReason: 'customer'
  },
  {
    id: 2,
    orderNumber: 'ORD-002',
    customerName: 'فاطمه محمدی',
    cancelledAt: '2024-01-14',
    totalAmount: 890000,
    cancellationReason: 'inventory'
  },
  {
    id: 3,
    orderNumber: 'ORD-003',
    customerName: 'محمد رضایی',
    cancelledAt: '2024-01-13',
    totalAmount: 2100000,
    cancellationReason: 'payment'
  }
])

// متغیرهای صفحه‌بندی
const currentPage = ref(1)
const itemsPerPage = ref(10)
const _totalItems = ref(cancelledOrders.value.length)

const totalPages = computed(() => Math.ceil(cancelledOrders.value.length / itemsPerPage.value))

// computed برای گزارشات
const _reportTotalPages = computed(() => {
  return Math.ceil(reportHistory.value.length / reportItemsPerPage.value)
})

const paginatedReportHistory = computed(() => {
  const start = (reportCurrentPage.value - 1) * reportItemsPerPage.value
  const end = start + reportItemsPerPage.value
  return reportHistory.value.slice(start, end)
})

// computed برای فیلتر کردن سفارشات کنسل شده
const filteredCancelledOrders = computed(() => {
  let filtered = cancelledOrders.value

  if (searchQuery.value) {
    filtered = filtered.filter(order => 
      order.orderNumber.includes(searchQuery.value) ||
      order.customerName.includes(searchQuery.value)
    )
  }

  if (reasonFilter.value) {
    filtered = filtered.filter(order => order.cancellationReason === reasonFilter.value)
  }

  if (paymentMethodFilter.value) {
    filtered = filtered.filter(order => order.paymentMethod === paymentMethodFilter.value)
  }

  if (minAmountFilter.value) {
    filtered = filtered.filter(order => order.totalAmount >= parseInt(minAmountFilter.value))
  }

  if (maxAmountFilter.value) {
    filtered = filtered.filter(order => order.totalAmount <= parseInt(maxAmountFilter.value))
  }

  if (dateFromFilter.value) {
    filtered = filtered.filter(order => new Date(order.cancelledAt) >= new Date(dateFromFilter.value))
  }

  if (dateToFilter.value) {
    filtered = filtered.filter(order => new Date(order.cancelledAt) <= new Date(dateToFilter.value))
  }

  return filtered
})

const paginatedCancelledOrders = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage.value
  const end = start + itemsPerPage.value
  return filteredCancelledOrders.value.slice(start, end)
})

// آمار انواع کنسل
const cancellationTypeStats = computed(() => {
  const stats = {}
  cancelledOrders.value.forEach(order => {
    stats[order.cancellationReason] = (stats[order.cancellationReason] || 0) + 1
  })
  return stats
})

// آمار روش‌های پرداخت
const paymentMethodStats = computed(() => {
  const stats = {}
  cancelledOrders.value.forEach(order => {
    stats[order.paymentMethod] = (stats[order.paymentMethod] || 0) + 1
  })
  return stats
})

// متغیر قدیمی برای سازگاری
const _pagination = ref({
  from: 1,
  to: 10,
  total: 45
})

// فیلتر کردن سفارشات
const _filteredOrders = computed(() => {
  let filtered = orders.value

  if (searchQuery.value) {
    filtered = filtered.filter(order => 
      order.orderNumber.includes(searchQuery.value) ||
      order.customerName.includes(searchQuery.value)
    )
  }

  if (reasonFilter.value) {
    filtered = filtered.filter(order => order.cancellationReason === reasonFilter.value)
  }

  if (dateFilter.value) {
    const today = new Date()
    const orderDate = new Date()
    
    switch (dateFilter.value) {
      case 'today':
        filtered = filtered.filter(order => {
          orderDate.setTime(new Date(order.cancelledAt).getTime())
          return orderDate.toDateString() === today.toDateString()
        })
        break
      case 'week':
        const weekAgo = new Date(today.getTime() - 7 * 24 * 60 * 60 * 1000)
        filtered = filtered.filter(order => new Date(order.cancelledAt) >= weekAgo)
        break
      case 'month':
        const monthAgo = new Date(today.getFullYear(), today.getMonth() - 1, today.getDate())
        filtered = filtered.filter(order => new Date(order.cancelledAt) >= monthAgo)
        break
    }
  }

  return filtered
})

// متدهای کمکی
const formatDate = (date) => {
  return new Date(date).toLocaleDateString('fa-IR')
}

const formatPrice = (price) => {
  return new Intl.NumberFormat('fa-IR').format(price) + ' تومان'
}

const getReasonClass = (reason) => {
  const classes = {
    customer: 'bg-red-100 text-red-800',
    inventory: 'bg-orange-100 text-orange-800',
    payment: 'bg-yellow-100 text-yellow-800',
    system: 'bg-gray-100 text-gray-800'
  }
  return classes[reason] || 'bg-gray-100 text-gray-800'
}

const getReasonText = (reason) => {
  const texts = {
    customer: 'درخواست مشتری',
    inventory: 'عدم موجودی',
    payment: 'مشکل پرداخت',
    system: 'مشکل سیستمی'
  }
  return texts[reason] || reason
}

const getPaymentMethodClass = (method) => {
  const classes = {
    online: 'bg-blue-100 text-blue-800',
    cash: 'bg-green-100 text-green-800',
    bank_transfer: 'bg-purple-100 text-purple-800',
    wallet: 'bg-yellow-100 text-yellow-800'
  }
  return classes[method] || 'bg-gray-100 text-gray-800'
}

const getPaymentMethodText = (method) => {
  const texts = {
    online: 'آنلاین',
    cash: 'نقدی',
    bank_transfer: 'انتقال بانکی',
    wallet: 'کیف پول'
  }
  return texts[method] || method
}

const refreshData = async () => {
  await fetchStats()
  await fetchCancelledOrders()
}

const exportData = () => {
  // خروجی اکسل
}

// متدهای مودال
const viewOrderDetails = (order) => {
  selectedOrder.value = order
  isModalOpen.value = true
}

const closeModal = () => {
  isModalOpen.value = false
  selectedOrder.value = null
}

const editOrder = (_order) => {
  // اینجا می‌توانید کاربر را به صفحه ویرایش هدایت کنید
}

const _viewOrder = (_order) => {
  // مشاهده جزئیات سفارش
}

const _restoreOrder = (_order) => {
  // بازگردانی سفارش
}



// متدهای صفحه‌بندی
const handlePageChange = (page) => {
  currentPage.value = page
  // اینجا می‌توانید داده‌های جدید را بارگذاری کنید
}

const handleItemsPerPageChange = (newItemsPerPage) => {
  itemsPerPage.value = newItemsPerPage
  currentPage.value = 1 // بازگشت به صفحه اول
  // اینجا می‌توانید داده‌های جدید را بارگذاری کنید
}

// متدهای گزارشات
const generateReport = () => {
}

const downloadQuickReport = (_period, _format) => {
}

const refreshReportHistory = () => {
}

const _handleReportPageChange = (page) => {
  reportCurrentPage.value = page
}

const _handleReportItemsPerPageChange = (newItemsPerPage) => {
  reportItemsPerPage.value = newItemsPerPage
  reportCurrentPage.value = 1
}
</script> 
