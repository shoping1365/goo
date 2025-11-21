<template>
  <div class="bg-white rounded-xl shadow-lg border border-gray-200 overflow-hidden" dir="rtl">
    <!-- Header Section -->
    <div class="bg-gradient-to-r from-blue-600 to-purple-600 px-6 py-4">
      <div class="flex items-center justify-between">
        <div class="flex items-center space-x-3 space-x-reverse">
          <div class="p-2 bg-white/20 rounded-lg">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"></path>
            </svg>
          </div>
          <div>
            <h2 class="text-xl font-bold text-white">لیست سفارشات نظرسنجی</h2>
            <p class="text-blue-100 text-sm">مدیریت و پیگیری ارسال SMS نظرسنجی</p>
          </div>
        </div>
        
        <div class="flex items-center space-x-3 space-x-reverse">
          <button 
            class="px-4 py-2 bg-white/20 hover:bg-white/30 text-white rounded-lg font-medium transition-colors flex items-center space-x-2 space-x-reverse"
            @click="exportToExcel"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
            </svg>
            <span>خروجی اکسل</span>
          </button>
          
          <button 
            class="px-4 py-2 bg-white/20 hover:bg-white/30 text-white rounded-lg font-medium transition-colors flex items-center space-x-2 space-x-reverse"
            @click="refreshData"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
            </svg>
            <span>بروزرسانی</span>
          </button>
        </div>
      </div>
    </div>

    <!-- Advanced Filters Section -->
    <div class="p-6 border-b border-gray-200 bg-gray-50">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
        <!-- Gender Filter -->
        <div class="space-y-2">
          <label class="block text-sm font-medium text-gray-700">جنسیت</label>
          <select 
            v-model="filters.gender" 
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه</option>
            <option value="male">مرد</option>
            <option value="female">زن</option>
          </select>
        </div>

        <!-- Date Range Filter -->
        <div class="space-y-2">
          <label class="block text-sm font-medium text-gray-700">بازه زمانی</label>
          <div class="flex space-x-2 space-x-reverse">
            <input 
              v-model="filters.dateFrom" 
              type="date" 
              class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
            <input 
              v-model="filters.dateTo" 
              type="date" 
              class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
          </div>
        </div>

        <!-- Amount Range Filter -->
        <div class="space-y-2">
          <label class="block text-sm font-medium text-gray-700">مبلغ سفارش</label>
          <div class="flex space-x-2 space-x-reverse">
            <input 
              v-model="filters.amountFrom" 
              type="number" 
              placeholder="از"
              class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
            <input 
              v-model="filters.amountTo" 
              type="number" 
              placeholder="تا"
              class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            >
          </div>
        </div>

        <!-- SMS Status Filter -->
        <div class="space-y-2">
          <label class="block text-sm font-medium text-gray-700">وضعیت SMS</label>
          <select 
            v-model="filters.smsStatus" 
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه</option>
            <option value="sent">ارسال شده</option>
            <option value="delivered">تحویل شده</option>
            <option value="failed">ناموفق</option>
            <option value="pending">در انتظار</option>
            <option value="not_sent">ارسال نشده</option>
          </select>
        </div>

        <!-- Geographic Region Filter -->
        <div class="space-y-2">
          <label class="block text-sm font-medium text-gray-700">منطقه جغرافیایی</label>
          <select 
            v-model="filters.region" 
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه مناطق</option>
            <option value="tehran">تهران</option>
            <option value="isfahan">اصفهان</option>
            <option value="shiraz">شیراز</option>
            <option value="mashhad">مشهد</option>
            <option value="tabriz">تبریز</option>
            <option value="other">سایر</option>
          </select>
        </div>

        <!-- Product Filter -->
        <div class="space-y-2">
          <label class="block text-sm font-medium text-gray-700">محصول</label>
          <select 
            v-model="filters.product" 
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه محصولات</option>
            <option value="electronics">الکترونیک</option>
            <option value="clothing">پوشاک</option>
            <option value="books">کتاب</option>
            <option value="home">خانه و آشپزخانه</option>
            <option value="sports">ورزشی</option>
          </select>
        </div>

        <!-- Response Status Filter -->
        <div class="space-y-2">
          <label class="block text-sm font-medium text-gray-700">وضعیت پاسخ</label>
          <select 
            v-model="filters.responseStatus" 
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه</option>
            <option value="responded">پاسخ داده</option>
            <option value="not_responded">پاسخ نداده</option>
            <option value="partial">پاسخ ناقص</option>
          </select>
        </div>

        <!-- Search -->
        <div class="space-y-2">
          <label class="block text-sm font-medium text-gray-700">جستجو</label>
          <input 
            v-model="filters.search" 
            type="text" 
            placeholder="نام، شماره سفارش یا تلفن..."
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
        </div>

        <!-- Filter Actions -->
        <div class="space-y-2">
          <label class="block text-sm font-medium text-gray-700">&nbsp;</label>
          <div class="flex space-x-2 space-x-reverse">
            <button 
              class="flex-1 px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg font-medium transition-colors"
              @click="applyFilters"
            >
              اعمال فیلتر
            </button>
            <button 
              class="px-4 py-2 bg-gray-500 hover:bg-gray-600 text-white rounded-lg font-medium transition-colors"
              @click="clearFilters"
            >
              پاک کردن
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Bulk Actions -->
    <div class="p-6 bg-yellow-50 border-b border-yellow-200">
      <div class="flex items-center justify-between">
        <div class="flex items-center space-x-4 space-x-reverse">
          <label class="flex items-center space-x-2 space-x-reverse">
            <input 
              v-model="selectAll" 
              type="checkbox"
              class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500" 
              @change="toggleSelectAll"
            >
            <span class="text-sm font-medium text-gray-700">انتخاب همه</span>
          </label>
          
          <span v-if="selectedOrders.length > 0" class="text-sm text-gray-600">
            {{ selectedOrders.length }} سفارش انتخاب شده
          </span>
        </div>
        
        <div v-if="selectedOrders.length > 0" class="flex items-center space-x-3 space-x-reverse">
          <button 
            class="px-4 py-2 bg-green-600 hover:bg-green-700 text-white rounded-lg font-medium transition-colors flex items-center space-x-2 space-x-reverse"
            @click="sendBulkSMS"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8"></path>
            </svg>
            <span>ارسال SMS گروهی</span>
          </button>
          
          <button 
            class="px-4 py-2 bg-orange-600 hover:bg-orange-700 text-white rounded-lg font-medium transition-colors flex items-center space-x-2 space-x-reverse"
            @click="resendFailedSMS"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
            </svg>
            <span>ارسال مجدد ناموفق‌ها</span>
          </button>
        </div>
      </div>
    </div>

    <!-- Table -->
    <div class="overflow-x-auto">
      <table class="w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              <input 
                v-model="selectAll" 
                type="checkbox"
                class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500" 
                @change="toggleSelectAll"
              >
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              شماره سفارش
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              مشتری
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              جنسیت
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              تاریخ سفارش
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              مبلغ
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              وضعیت SMS
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              تاریخ ارسال
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              وضعیت پاسخ
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              امتیاز
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              عملیات
            </th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr 
            v-for="order in filteredOrders" 
            :key="order.id"
            class="hover:bg-gray-50 transition-colors"
          >
            <td class="px-6 py-4 whitespace-nowrap">
              <input 
                v-model="selectedOrders" 
                :value="order.id"
                type="checkbox" 
                class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
              >
            </td>
            
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="text-sm font-medium text-gray-900">#{{ order.orderNumber }}</div>
            </td>
            
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="flex items-center">
                <div class="">
                  <div class="h-10 w-10 rounded-full bg-gradient-to-r from-blue-400 to-purple-500 flex items-center justify-center">
                    <span class="text-white font-medium text-sm">{{ getInitials(order.customerName) }}</span>
                  </div>
                </div>
                <div class="mr-4">
                  <div class="text-sm font-medium text-gray-900">{{ order.customerName }}</div>
                  <div class="text-sm text-gray-500">{{ order.phoneNumber }}</div>
                </div>
              </div>
            </td>
            
            <td class="px-6 py-4 whitespace-nowrap">
              <span 
                :class="[
                  'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium',
                  order.gender === 'male' 
                    ? 'bg-blue-100 text-blue-800' 
                    : 'bg-pink-100 text-pink-800'
                ]"
              >
                <svg 
                  v-if="order.gender === 'male'" 
                  class="w-3 h-3 ml-1" 
                  fill="currentColor" 
                  viewBox="0 0 20 20"
                >
                  <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM9.555 7.168A1 1 0 008 8v4a1 1 0 001.555.832l3-2a1 1 0 000-1.664l-3-2z" clip-rule="evenodd"></path>
                </svg>
                <svg 
                  v-else 
                  class="w-3 h-3 ml-1" 
                  fill="currentColor" 
                  viewBox="0 0 20 20"
                >
                  <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM9.555 7.168A1 1 0 008 8v4a1 1 0 001.555.832l3-2a1 1 0 000-1.664l-3-2z" clip-rule="evenodd"></path>
                </svg>
                {{ order.gender === 'male' ? 'مرد' : 'زن' }}
              </span>
            </td>
            
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
              {{ formatDate(order.orderDate) }}
            </td>
            
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
              {{ formatCurrency(order.amount) }}
            </td>
            
            <td class="px-6 py-4 whitespace-nowrap">
              <span 
                :class="[
                  'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium',
                  getSMSStatusClass(order.smsStatus)
                ]"
              >
                <svg 
                  :class="getSMSStatusIcon(order.smsStatus)" 
                  class="w-3 h-3 ml-1" 
                  fill="currentColor" 
                  viewBox="0 0 20 20"
                >
                  <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"></path>
                </svg>
                {{ getSMSStatusText(order.smsStatus) }}
              </span>
            </td>
            
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
              {{ order.smsSentDate ? formatDate(order.smsSentDate) : '-' }}
            </td>
            
            <td class="px-6 py-4 whitespace-nowrap">
              <span 
                :class="[
                  'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium',
                  getResponseStatusClass(order.responseStatus)
                ]"
              >
                {{ getResponseStatusText(order.responseStatus) }}
              </span>
            </td>
            
            <td class="px-6 py-4 whitespace-nowrap">
              <div v-if="order.rating" class="flex items-center">
                <div class="flex items-center space-x-1 space-x-reverse">
                  <svg 
                    v-for="star in 5" 
                    :key="star"
                    :class="star <= order.rating ? 'text-yellow-400' : 'text-gray-300'"
                    class="w-4 h-4" 
                    fill="currentColor" 
                    viewBox="0 0 20 20"
                  >
                    <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"></path>
                  </svg>
                </div>
                <span class="mr-2 text-sm text-gray-600">({{ order.rating }}/5)</span>
              </div>
              <span v-else class="text-gray-400 text-sm">-</span>
            </td>
            
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
              <div class="flex items-center space-x-2 space-x-reverse">
                <button 
                  class="text-blue-600 hover:text-blue-900 transition-colors"
                  title="ارسال SMS"
                  @click="sendSMS(order.id)"
                >
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8"></path>
                  </svg>
                </button>
                
                <button 
                  class="text-green-600 hover:text-green-900 transition-colors"
                  title="مشاهده جزئیات"
                  @click="viewDetails(order.id)"
                >
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                  </svg>
                </button>
                
                <button 
                  class="text-orange-600 hover:text-orange-900 transition-colors"
                  title="ویرایش"
                  @click="editOrder(order.id)"
                >
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
                  </svg>
                </button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Pagination -->
    <div class="bg-white px-4 py-3 flex items-center justify-between border-t border-gray-200 sm:px-6">
      <div class="flex-1 flex justify-between sm:hidden">
        <button 
          :disabled="currentPage === 1"
          class="relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
          @click="previousPage"
        >
          قبلی
        </button>
        <button 
          :disabled="currentPage >= totalPages"
          class="mr-3 relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
          @click="nextPage"
        >
          بعدی
        </button>
      </div>
      <div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
        <div>
          <p class="text-sm text-gray-700">
            نمایش 
            <span class="font-medium">{{ (currentPage - 1) * itemsPerPage + 1 }}</span>
            تا 
            <span class="font-medium">{{ Math.min(currentPage * itemsPerPage, totalItems) }}</span>
            از 
            <span class="font-medium">{{ totalItems }}</span>
            نتیجه
          </p>
        </div>
        <div>
          <nav class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px" aria-label="Pagination">
            <button 
              :disabled="currentPage === 1"
              class="relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
              @click="previousPage"
            >
              <span class="sr-only">قبلی</span>
              <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                <path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd" />
              </svg>
            </button>
            
            <button 
              v-for="page in visiblePages" 
              :key="page"
              :class="[
                'relative inline-flex items-center px-4 py-2 border text-sm font-medium',
                page === currentPage
                  ? 'z-10 bg-blue-50 border-blue-500 text-blue-600'
                  : 'bg-white border-gray-300 text-gray-500 hover:bg-gray-50'
              ]"
              @click="goToPage(page)"
            >
              {{ page }}
            </button>
            
            <button 
              :disabled="currentPage >= totalPages"
              class="relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
              @click="nextPage"
            >
              <span class="sr-only">بعدی</span>
              <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                <path fill-rule="evenodd" d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z" clip-rule="evenodd" />
              </svg>
            </button>
          </nav>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, reactive, ref } from 'vue'

// Types
interface Order {
  id: number
  orderNumber: string
  customerName: string
  phoneNumber: string
  gender: 'male' | 'female'
  orderDate: string
  amount: number
  smsStatus: 'sent' | 'delivered' | 'failed' | 'pending' | 'not_sent'
  smsSentDate?: string
  responseStatus: 'responded' | 'not_responded' | 'partial'
  rating?: number
  region: string
  product: string
}

interface Filters {
  gender: string
  dateFrom: string
  dateTo: string
  amountFrom: string
  amountTo: string
  smsStatus: string
  region: string
  product: string
  responseStatus: string
  search: string
}

// Reactive data
const orders = ref<Order[]>([])
const selectedOrders = ref<number[]>([])
const selectAll = ref(false)
const currentPage = ref(1)
const itemsPerPage = ref(20)
const totalItems = ref(0)
const loading = ref(false)

const filters = reactive<Filters>({
  gender: '',
  dateFrom: '',
  dateTo: '',
  amountFrom: '',
  amountTo: '',
  smsStatus: '',
  region: '',
  product: '',
  responseStatus: '',
  search: ''
})

// Computed properties
const filteredOrders = computed(() => {
  let filtered = [...orders.value]
  
  if (filters.gender) {
    filtered = filtered.filter(order => order.gender === filters.gender)
  }
  
  if (filters.dateFrom) {
    filtered = filtered.filter(order => new Date(order.orderDate) >= new Date(filters.dateFrom))
  }
  
  if (filters.dateTo) {
    filtered = filtered.filter(order => new Date(order.orderDate) <= new Date(filters.dateTo))
  }
  
  if (filters.amountFrom) {
    filtered = filtered.filter(order => order.amount >= parseFloat(filters.amountFrom))
  }
  
  if (filters.amountTo) {
    filtered = filtered.filter(order => order.amount <= parseFloat(filters.amountTo))
  }
  
  if (filters.smsStatus) {
    filtered = filtered.filter(order => order.smsStatus === filters.smsStatus)
  }
  
  if (filters.region) {
    filtered = filtered.filter(order => order.region === filters.region)
  }
  
  if (filters.product) {
    filtered = filtered.filter(order => order.product === filters.product)
  }
  
  if (filters.responseStatus) {
    filtered = filtered.filter(order => order.responseStatus === filters.responseStatus)
  }
  
  if (filters.search) {
    const searchTerm = filters.search.toLowerCase()
    filtered = filtered.filter(order => 
      order.customerName.toLowerCase().includes(searchTerm) ||
      order.orderNumber.toLowerCase().includes(searchTerm) ||
      order.phoneNumber.includes(searchTerm)
    )
  }
  
  return filtered
})

const totalPages = computed(() => Math.ceil(filteredOrders.value.length / itemsPerPage.value))

const visiblePages = computed(() => {
  const pages = []
  const start = Math.max(1, currentPage.value - 2)
  const end = Math.min(totalPages.value, currentPage.value + 2)
  
  for (let i = start; i <= end; i++) {
    pages.push(i)
  }
  
  return pages
})

// Methods
const loadOrders = async () => {
  loading.value = true
  try {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    // Mock data
    orders.value = [
      {
        id: 1,
        orderNumber: 'ORD-001',
        customerName: 'علی احمدی',
        phoneNumber: '09123456789',
        gender: 'male',
        orderDate: '2024-01-15',
        amount: 250000,
        smsStatus: 'delivered',
        smsSentDate: '2024-01-18',
        responseStatus: 'responded',
        rating: 5,
        region: 'tehran',
        product: 'electronics'
      },
      {
        id: 2,
        orderNumber: 'ORD-002',
        customerName: 'فاطمه محمدی',
        phoneNumber: '09187654321',
        gender: 'female',
        orderDate: '2024-01-16',
        amount: 180000,
        smsStatus: 'sent',
        smsSentDate: '2024-01-19',
        responseStatus: 'not_responded',
        region: 'isfahan',
        product: 'clothing'
      },
      {
        id: 3,
        orderNumber: 'ORD-003',
        customerName: 'محمد رضایی',
        phoneNumber: '09111111111',
        gender: 'male',
        orderDate: '2024-01-17',
        amount: 320000,
        smsStatus: 'failed',
        responseStatus: 'not_responded',
        region: 'shiraz',
        product: 'electronics'
      },
      {
        id: 4,
        orderNumber: 'ORD-004',
        customerName: 'زهرا کریمی',
        phoneNumber: '09222222222',
        gender: 'female',
        orderDate: '2024-01-18',
        amount: 95000,
        smsStatus: 'not_sent',
        responseStatus: 'not_responded',
        region: 'mashhad',
        product: 'books'
      },
      {
        id: 5,
        orderNumber: 'ORD-005',
        customerName: 'حسین نوری',
        phoneNumber: '09333333333',
        gender: 'male',
        orderDate: '2024-01-19',
        amount: 450000,
        smsStatus: 'delivered',
        smsSentDate: '2024-01-22',
        responseStatus: 'responded',
        rating: 4,
        region: 'tehran',
        product: 'home'
      }
    ]
    
    totalItems.value = orders.value.length
  } catch {
    // console.error('Error loading orders:', error)
  } finally {
    loading.value = false
  }
}

const applyFilters = () => {
  currentPage.value = 1
  // Filters are applied automatically through computed property
}

const clearFilters = () => {
  Object.keys(filters).forEach(key => {
    filters[key as keyof Filters] = ''
  })
  currentPage.value = 1
}

const toggleSelectAll = () => {
  if (selectAll.value) {
    selectedOrders.value = filteredOrders.value.map(order => order.id)
  } else {
    selectedOrders.value = []
  }
}

const sendSMS = async (orderId: number) => {
  try {
    // Simulate SMS sending
    await new Promise(resolve => setTimeout(resolve, 1000))
    // Update order status
    const order = orders.value.find(o => o.id === orderId)
    if (order) {
      order.smsStatus = 'sent'
      order.smsSentDate = new Date().toISOString().split('T')[0]
    }
  } catch {
    // console.error('Error sending SMS:', error)
  }
}

const sendBulkSMS = async () => {
  try {
    // Simulate bulk SMS sending
    await new Promise(resolve => setTimeout(resolve, 2000))
    // Update all selected orders
    selectedOrders.value.forEach(orderId => {
      const order = orders.value.find(o => o.id === orderId)
      if (order) {
        order.smsStatus = 'sent'
        order.smsSentDate = new Date().toISOString().split('T')[0]
      }
    })
    selectedOrders.value = []
    selectAll.value = false
  } catch {
    // console.error('Error sending bulk SMS:', error)
  }
}

const resendFailedSMS = async () => {
  try {
    const failedOrders = selectedOrders.value.filter(orderId => {
      const order = orders.value.find(o => o.id === orderId)
      return order?.smsStatus === 'failed'
    })
    
    await new Promise(resolve => setTimeout(resolve, 1500))
    
    failedOrders.forEach(orderId => {
      const order = orders.value.find(o => o.id === orderId)
      if (order) {
        order.smsStatus = 'sent'
        order.smsSentDate = new Date().toISOString().split('T')[0]
      }
    })
    
    selectedOrders.value = []
    selectAll.value = false
  } catch {
    // console.error('Error resending failed SMS:', error)
  }
}

const viewDetails = (_orderId: number) => {
  // Navigate to order details page
}

const editOrder = (_orderId: number) => {
  // Navigate to edit order page
}

const exportToExcel = () => {
  // Implement Excel export functionality
}

const refreshData = () => {
  loadOrders()
}

const previousPage = () => {
  if (currentPage.value > 1) {
    currentPage.value--
  }
}

const nextPage = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value++
  }
}

const goToPage = (page: number) => {
  currentPage.value = page
}

// Utility functions
const getInitials = (name: string) => {
  return name.split(' ').map(n => n[0]).join('').toUpperCase()
}

const formatDate = (date: string) => {
  return new Date(date).toLocaleDateString('fa-IR')
}

const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

const getSMSStatusClass = (status: string) => {
  const classes = {
    sent: 'bg-blue-100 text-blue-800',
    delivered: 'bg-green-100 text-green-800',
    failed: 'bg-red-100 text-red-800',
    pending: 'bg-yellow-100 text-yellow-800',
    not_sent: 'bg-gray-100 text-gray-800'
  }
  return classes[status as keyof typeof classes] || 'bg-gray-100 text-gray-800'
}

const getSMSStatusText = (status: string) => {
  const texts = {
    sent: 'ارسال شده',
    delivered: 'تحویل شده',
    failed: 'ناموفق',
    pending: 'در انتظار',
    not_sent: 'ارسال نشده'
  }
  return texts[status as keyof typeof texts] || status
}

const getSMSStatusIcon = (status: string) => {
  const icons = {
    sent: 'w-3 h-3 ml-1 text-blue-600',
    delivered: 'w-3 h-3 ml-1 text-green-600',
    failed: 'w-3 h-3 ml-1 text-red-600',
    pending: 'w-3 h-3 ml-1 text-yellow-600',
    not_sent: 'w-3 h-3 ml-1 text-gray-600'
  }
  return icons[status as keyof typeof icons] || 'w-3 h-3 ml-1 text-gray-600'
}

const getResponseStatusClass = (status: string) => {
  const classes = {
    responded: 'bg-green-100 text-green-800',
    not_responded: 'bg-gray-100 text-gray-800',
    partial: 'bg-orange-100 text-orange-800'
  }
  return classes[status as keyof typeof classes] || 'bg-gray-100 text-gray-800'
}

const getResponseStatusText = (status: string) => {
  const texts = {
    responded: 'پاسخ داده',
    not_responded: 'پاسخ نداده',
    partial: 'پاسخ ناقص'
  }
  return texts[status as keyof typeof texts] || status
}

// Lifecycle
onMounted(() => {
  loadOrders()
})
</script>

<style scoped>
/* Custom styles for better RTL support */
.space-x-reverse > :not([hidden]) ~ :not([hidden]) {
  --tw-space-x-reverse: 1;
}

/* Smooth transitions */
.transition-colors {
  transition-property: color, background-color, border-color, text-decoration-color, fill, stroke;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  transition-duration: 150ms;
}

/* Hover effects */
.hover\:bg-gray-50:hover {
  background-color: rgb(249 250 251);
}

.hover\:bg-blue-50:hover {
  background-color: rgb(239 246 255);
}

.hover\:bg-green-50:hover {
  background-color: rgb(240 253 244);
}

.hover\:bg-orange-50:hover {
  background-color: rgb(255 247 237);
}

/* Focus states */
.focus\:ring-2:focus {
  --tw-ring-offset-shadow: var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color);
  --tw-ring-shadow: var(--tw-ring-inset) 0 0 0 calc(2px + var(--tw-ring-offset-width)) var(--tw-ring-color);
  box-shadow: var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000);
}

.focus\:ring-blue-500:focus {
  --tw-ring-color: rgb(59 130 246);
}

.focus\:border-blue-500:focus {
  border-color: rgb(59 130 246);
}
</style> 
