<template>

    <!-- Page Header -->
    <div class="bg-white shadow-sm border-b border-gray-200 mb-4">
      <div class="w-full px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center py-3">
          <div>
            <h1 class="text-lg font-bold text-gray-900">مدیریت پرسش و پاسخ</h1>
            <p class="text-xs text-gray-600 mt-1">پاسخ به سوالات مشتریان و مدیریت پرسش‌ها</p>
          </div>
          <div class="flex space-x-2 space-x-reverse">
            <button
                class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-purple-400 to-purple-600 hover:from-purple-500 hover:to-purple-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-purple-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105"
                @click="showCategoriesModal = true"
            >
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"></path>
              </svg>
              مدیریت دسته‌بندی
            </button>
            <ExportExcelButton :data="exportQuestionsData" filename="questions.csv" />
            <button
                class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-cyan-400 to-cyan-600 hover:from-cyan-500 hover:to-cyan-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-cyan-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105"
                @click="showFilters = !showFilters"
            >
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.707A1 1 0 013 7V4z"></path>
              </svg>
              فیلترها
            </button>
            <button 
              class="inline-flex items-center px-3 py-2 border border-gray-200 rounded-lg bg-white hover:bg-gray-50 transition-all shadow-md" 
              @click="showSettingsModal = true"
            >
              <svg v-if="settingsLoading" class="w-5 h-5 text-gray-500 animate-spin" fill="none" viewBox="0 0 24 24">
                <circle cx="12" cy="12" r="4" stroke="currentColor" stroke-width="2" fill="none" />
                <path stroke="currentColor" stroke-width="2" fill="none" d="M4.93 4.93l2.12 2.12M2 12h3M4.93 19.07l2.12-2.12M12 22v-3M19.07 19.07l-2.12-2.12M22 12h-3M19.07 4.93l-2.12 2.12M12 2v3" />
              </svg>
              <svg v-else class="w-5 h-5 text-gray-500" fill="none" viewBox="0 0 24 24">
                <circle cx="12" cy="12" r="4" stroke="currentColor" stroke-width="2" fill="none" />
                <path stroke="currentColor" stroke-width="2" fill="none" d="M4.93 4.93l2.12 2.12M2 12h3M4.93 19.07l2.12-2.12M12 22v-3M19.07 19.07l-2.12-2.12M22 12h-3M19.07 4.93l-2.12 2.12M12 2v3" />
              </svg>
            </button>
          </div>
        </div>
      </div>
    </div>

    <div class="w-full px-4 py-4">

    <div class="mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <div class="relative">
        <div class="absolute left-0 right-0 top-2 bottom-2 bg-white rounded-2xl shadow-lg py-6" style="z-index:0;"></div>
        <div class="relative grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-6 mb-6" style="z-index:1;">
          <div
              class="bg-gradient-to-br from-blue-50 to-blue-100 overflow-hidden shadow-md rounded-lg cursor-pointer border border-blue-200 transition-all duration-200 hover:shadow-lg hover:scale-105"
              :class="filters.status === '' ? 'ring-2 ring-blue-400 shadow-lg' : ''"
              @click="setStatusFilter('')"
          >
            <div class="p-3">
              <div class="flex items-center justify-between">
                <div class="flex-1">
                  <dl>
                    <dt class="text-xs font-medium text-blue-600 truncate">کل پرسش‌ها</dt>
                    <dd class="text-base font-bold text-blue-800">{{ stats.total }}</dd>
                  </dl>
                </div>
                <div class="flex-shrink-0 mr-3">
                  <div class="w-8 h-8 bg-gradient-to-br from-blue-400 to-blue-500 rounded-lg flex items-center justify-center shadow-md">
                    <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                    </svg>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <div
              class="bg-gradient-to-br from-green-50 to-emerald-100 overflow-hidden shadow-md rounded-lg cursor-pointer border border-green-200 transition-all duration-200 hover:shadow-lg hover:scale-105"
              :class="filters.status === 'answered' ? 'ring-2 ring-green-400 shadow-lg' : ''"
              @click="setStatusFilter('answered')"
          >
            <div class="p-3">
              <div class="flex items-center justify-between">
                <div class="flex-1">
                  <dl>
                    <dt class="text-xs font-medium text-green-600 truncate">پاسخ داده شده</dt>
                    <dd class="text-base font-bold text-green-800">{{ stats.answered }}</dd>
                  </dl>
                </div>
                <div class="flex-shrink-0 mr-3">
                  <div class="w-8 h-8 bg-gradient-to-br from-green-400 to-emerald-500 rounded-lg flex items-center justify-center shadow-md">
                    <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                    </svg>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <div
              class="bg-gradient-to-br from-amber-50 to-yellow-100 overflow-hidden shadow-md rounded-lg cursor-pointer border border-amber-200 transition-all duration-200 hover:shadow-lg hover:scale-105"
              :class="filters.status === 'pending' ? 'ring-2 ring-amber-400 shadow-lg' : ''"
              @click="setStatusFilter('pending')"
          >
            <div class="p-3">
              <div class="flex items-center justify-between">
                <div class="flex-1">
                  <dl>
                    <dt class="text-xs font-medium text-amber-600 truncate">در انتظار پاسخ</dt>
                    <dd class="text-base font-bold text-amber-800">{{ stats.pending }}</dd>
                  </dl>
                </div>
                <div class="flex-shrink-0 mr-3">
                  <div class="w-8 h-8 bg-gradient-to-br from-amber-400 to-yellow-500 rounded-lg flex items-center justify-center shadow-md">
                    <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                    </svg>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <div
              class="bg-gradient-to-br from-rose-50 to-red-100 overflow-hidden shadow-md rounded-lg cursor-pointer border border-rose-200 transition-all duration-200 hover:shadow-lg hover:scale-105"
              :class="filters.status === 'rejected' ? 'ring-2 ring-rose-400 shadow-lg' : ''"
              @click="setStatusFilter('rejected')"
          >
            <div class="p-3">
              <div class="flex items-center justify-between">
                <div class="flex-1">
                  <dl>
                    <dt class="text-xs font-medium text-rose-600 truncate">رد شده</dt>
                    <dd class="text-base font-bold text-rose-800">{{ stats.rejected }}</dd>
                  </dl>
                </div>
                <div class="flex-shrink-0 mr-3">
                  <div class="w-8 h-8 bg-gradient-to-br from-rose-400 to-red-500 rounded-lg flex items-center justify-center shadow-md">
                    <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                    </svg>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <div
              class="bg-gradient-to-br from-cyan-50 to-cyan-100 overflow-hidden shadow-md rounded-lg cursor-pointer border border-cyan-200 transition-all duration-200 hover:shadow-lg hover:scale-105"
              :class="filters.status === 'reviewing' ? 'ring-2 ring-cyan-400 shadow-lg' : ''"
              @click="setStatusFilter('reviewing')"
          >
            <div class="p-3">
              <div class="flex items-center justify-between">
                <div class="flex-1">
                  <dl>
                    <dt class="text-xs font-medium text-cyan-600 truncate">در حال بررسی</dt>
                    <dd class="text-base font-bold text-cyan-800">{{ stats.reviewing || 0 }}</dd>
                  </dl>
                </div>
                <div class="flex-shrink-0 mr-3">
                  <div class="w-8 h-8 bg-gradient-to-br from-cyan-400 to-cyan-500 rounded-lg flex items-center justify-center shadow-md">
                    <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
                    </svg>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div v-if="showFilters" class="bg-gradient-to-br from-blue-50 to-cyan-50 shadow-lg rounded-2xl mb-8 border border-blue-100">
        <div class="px-6 py-4 border-b border-blue-100 flex items-center gap-2">
          <svg class="w-6 h-6 text-blue-400" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.707A1 1 0 013 7V4z"></path></svg>
          <h3 class="text-lg font-bold text-blue-900">فیلترها و جستجو</h3>
        </div>
        <div class="p-8">
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">جستجو در پرسش‌ها</label>
              <div class="relative">
                <input
                    v-model="filters.search"
                    type="text"
                    class="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md leading-5 bg-white placeholder-gray-500 focus:outline-none focus:placeholder-gray-400 focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm"
                    placeholder="نام مشتری، محصول یا متن سوال..."
                    dir="rtl"
                />
                <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <svg class="h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
                  </svg>
                </div>
              </div>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت پرسش</label>
              <select v-model="filters.status" class="block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-blue-500 focus:border-blue-500 text-sm">
                <option value="">همه وضعیت‌ها</option>
                <option value="pending">در انتظار پاسخ</option>
                <option value="reviewing">در حال بررسی</option>
                <option value="answered">پاسخ داده شده</option>
                <option value="rejected">رد شده</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نوع سوال</label>
              <select v-model="filters.category" class="block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-blue-500 focus:border-blue-500 text-sm">
                <option value="">همه دسته‌ها</option>
                <option v-for="category in categories" :key="category.key" :value="category.key">
                  {{ category.name }}
                </option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">بازه زمانی</label>
              <select v-model="filters.dateRange" class="block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-blue-500 focus:border-blue-500 text-sm">
                <option value="">همه تاریخ‌ها</option>
                <option value="today">امروز</option>
                <option value="week">هفته گذشته</option>
                <option value="month">ماه گذشته</option>
                <option value="3months">3 ماه گذشته</option>
              </select>
            </div>
          </div>
          <div class="flex gap-6 justify-end mt-8">
          
          </div>
        </div>
      </div>

      <div v-if="selectedQuestions.length > 0" class="bg-blue-50 border border-blue-200 rounded-lg p-6 mb-6">
        <div class="flex items-center justify-between">
          <div class="flex items-center">
            <span class="text-sm font-medium text-blue-800">{{ selectedQuestions.length }} پرسش انتخاب شده</span>
          </div>
          <div class="flex space-x-2 space-x-reverse">
            <div class="relative">
              <button
                  class="inline-flex items-center px-3 py-2 border border-transparent text-xs font-medium rounded text-white bg-orange-600 hover:bg-orange-700"
                  @click="showTransferDropdown = !showTransferDropdown"
              >
                <svg class="w-4 h-4 ml-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7h12m0 0l-4-4m4 4l-4 4m0 6H4m0 0l4 4m-4-4l4-4"></path>
                </svg>
                انتقال به
                <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
                </svg>
              </button>
              
              <!-- Transfer Dropdown -->
              <div v-if="showTransferDropdown" class="absolute right-0 mt-2 w-56 rounded-md shadow-lg bg-white ring-1 ring-black ring-opacity-5 z-10">
                <div class="py-1">
                  <div class="px-4 py-2 text-xs text-gray-500 border-b">انتقال {{ selectedQuestions.length }} پرسش به:</div>
                  <button
                      v-for="category in categories"
                      :key="category.key"
                      class="block w-full text-right px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                      @click="bulkTransferToCategory(category.key)"
                  >
                    {{ category.name }}
                  </button>
                  <div class="border-t">
                    <button
                        class="block w-full text-right px-4 py-2 text-sm text-gray-500 hover:bg-gray-100"
                        @click="showTransferDropdown = false"
                    >
                      لغو
                    </button>
                  </div>
                </div>
              </div>
            </div>
            <button
                class="inline-flex items-center px-3 py-2 border border-transparent text-xs font-medium rounded text-white bg-red-600 hover:bg-red-700"
                @click="bulkReject"
            >
              رد {{ selectedQuestions.length }} آیتم
            </button>
            <button
                class="inline-flex items-center px-3 py-2 border border-transparent text-xs font-medium rounded text-white bg-gray-600 hover:bg-gray-700"
                @click="bulkDelete"
            >
              حذف {{ selectedQuestions.length }} آیتم
            </button>
          </div>
        </div>
      </div>

      <div class="bg-white shadow rounded-lg overflow-hidden">
        <div class="px-6 py-4 border-b border-gray-200">
          <div class="flex items-center justify-between">
            <h3 class="text-lg font-medium text-gray-900">لیست پرسش‌ها</h3>
            <div class="flex items-center space-x-2 space-x-reverse">
              <span class="text-sm text-gray-500">{{ paginatedQuestions.length }} از {{ filteredQuestions.length }} پرسش</span>
            </div>
          </div>
        </div>

        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
            <tr>
              <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                <input
                    type="checkbox"
                    :checked="allSelected"
                    class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                    @change="toggleSelectAll"
                />
              </th>
              <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                مشتری
              </th>
              <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider w-64 sm:w-80 md:w-96">
                محصول
              </th>
              <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                دسته‌بندی
              </th>
              <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                سوال
              </th>
              <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                تاریخ
              </th>
              <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                وضعیت
              </th>
              <th scope="col" class="relative px-6 py-3">
                <span class="sr-only">عملیات</span>
              </th>
            </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="question in paginatedQuestions" :key="question.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <input
                    v-model="selectedQuestions"
                    type="checkbox"
                    :value="question.id"
                    class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                />
              </td>

              <td class="px-6 py-4 whitespace-nowrap text-right">
                <div>
                  <div class="text-sm font-medium text-gray-900">{{ question.customer.name }}</div>
                  <div class="text-sm text-gray-500">{{ question.customer.email }}</div>
                </div>
              </td>

              <td class="px-6 py-4 text-right">
                <NuxtLink v-if="question.product" :to="`/admin/product-management/products/edit?id=${question.product.id}`" class="text-sm font-medium text-blue-600 hover:underline block text-right">
                  {{ question.product.name }}
                </NuxtLink>
                <div v-if="question.product" class="text-sm text-gray-500 text-right">#{{ question.product.id }}</div>
                <span v-else class="text-sm text-gray-400 text-right">بدون محصول</span>
              </td>

              <td class="px-6 py-4 whitespace-nowrap text-right">
                  <span :class="getCategoryClass(question.category)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                    {{ getCategoryLabel(question.category) }}
                  </span>
              </td>

              <td class="px-6 py-4 text-right">
                <div class="text-sm text-gray-900 max-w-xs text-right" dir="rtl">
                  <p class="line-clamp-2">{{ truncateText(question.question, 100) }}</p>
                </div>
              </td>

              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900 text-right">
                <div class="text-right">{{ formatDate(question.createdAt) }}</div>
                <div class="text-xs text-gray-500 text-right">{{ formatTime(question.createdAt) }}</div>
              </td>

              <td class="px-6 py-4 whitespace-nowrap text-right">
                  <span :class="getStatusClass(question.status)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                    {{ getStatusLabel(question.status) }}
                  </span>
              </td>

              <td class="px-6 py-4 whitespace-nowrap text-left text-sm font-medium">
                <div class="flex items-center space-x-2 space-x-reverse">
                  <button
                      class="text-blue-600 hover:text-blue-900 p-1 rounded-full hover:bg-blue-100"
                      title="ویرایش"
                      @click="openEditModal(question)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
                    </svg>
                  </button>
                  <button
                      class="text-red-600 hover:text-red-900 p-1 rounded-full hover:bg-red-100"
                      title="حذف"
                      @click="deleteQuestion(question.id)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
            </tbody>
          </table>
        </div>

        <div class="bg-white px-4 py-3 border-t border-gray-200 sm:px-6">
          <div class="flex items-center justify-between">
            <div class="flex-1 flex justify-between sm:hidden">
              <button
                  :disabled="currentPage === 1"
                  class="relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50"
                  @click="previousPage"
              >
                قبلی
              </button>
              <button
                  :disabled="currentPage >= totalPages"
                  class="mr-3 relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 disabled:opacity-50"
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
                  <span class="font-medium">{{ Math.min(currentPage * itemsPerPage, filteredQuestions.length) }}</span>
                  از
                  <span class="font-medium">{{ filteredQuestions.length }}</span>
                  نتیجه
                </p>
              </div>
              <div>
                <nav class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px">
                  <button
                      :disabled="currentPage === 1"
                      class="relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50"
                      @click="previousPage"
                  >
                    <svg class="h-5 w-5" fill="currentColor" viewBox="0 0 20 20">
                      <path fill-rule="evenodd" d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z" clip-rule="evenodd" />
                    </svg>
                  </button>

                  <span v-for="page in visiblePages" :key="page">
                    <button
                        v-if="page !== '...'"
                        :class="page === currentPage ? 'bg-blue-50 border-blue-500 text-blue-600' : 'bg-white border-gray-300 text-gray-500 hover:bg-gray-50'"
                        class="relative inline-flex items-center px-4 py-2 border text-sm font-medium"
                        @click="goToPage(Number(page))"
                    >
                      {{ page }}
                    </button>
                    <span v-else class="relative inline-flex items-center px-4 py-2 border border-gray-300 bg-white text-sm font-medium text-gray-700">
                      ...
                    </span>
                  </span>

                  <button
                      :disabled="currentPage >= totalPages"
                      class="relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50"
                      @click="nextPage"
                  >
                    <svg class="h-5 w-5" fill="currentColor" viewBox="0 0 20 20">
                      <path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd" />
                    </svg>
                  </button>
                </nav>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div v-if="showQuestionModal" class="fixed inset-0 z-50 overflow-y-auto">
      <div class="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
        <!-- Overlay fully removed -->
        <span class="hidden sm:inline-block sm:align-middle sm:h-screen">&#8203;</span>

        <div class="inline-block align-bottom bg-white rounded-xl text-right overflow-hidden shadow-2xl transform transition-all sm:my-8 sm:align-middle sm:max-w-7xl sm:w-full border border-indigo-100" dir="rtl">
          <div class="bg-gradient-to-r from-indigo-500 to-purple-600 px-4 py-4">
            <div class="flex items-center justify-between">
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-white/20 backdrop-blur-sm rounded-lg flex items-center justify-center">
                  <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
                </div>
                <div><h3 class="text-lg font-bold text-white">جزئیات پرسش مشتری</h3></div>
              </div>
              <button class="text-white/80 hover:text-white hover:bg-white/10 rounded-lg p-1 transition-all" @click="closeDetailModal">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path></svg>
              </button>
            </div>
          </div>

          <div class="px-4 pt-4 pb-2 max-h-[75vh] overflow-y-auto">
            <div v-if="selectedQuestion" class="space-y-4">
              <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 items-stretch">
                                <div class="bg-gradient-to-br from-blue-50 to-indigo-100 rounded-xl p-6 border border-blue-200 shadow-sm">
                  <div class="flex items-center justify-between mb-4">
                    <div class="flex items-center">
                      <div class="w-12 h-12 rounded-full bg-blue-500 flex items-center justify-center ml-3">
                        <span class="text-white font-bold text-lg">{{ (selectedQuestion.customer_name || selectedQuestion.customer?.name || 'مهمان').charAt(0) }}</span>
                      </div>
                      <div>
                        <div class="text-sm font-semibold text-gray-900">
                          {{ selectedQuestion.customer_name || selectedQuestion.customer?.name || 'کاربر مهمان' }}
                          <span v-if="selectedQuestion.isAnonymous" class="text-xs bg-orange-100 text-orange-800 px-2 py-1 rounded-full mr-2">درخواست ناشناس</span>
                        </div>
                        <div class="text-xs text-gray-500">
                          {{ selectedQuestion.customer_id ? 'کاربر عضو' : 'کاربر مهمان' }}
                        </div>
                      </div>
                    </div>
                  </div>
                  
                  <div class="grid grid-cols-2 gap-3">
                    <div class="bg-white rounded-lg p-3 flex items-center shadow-sm">
                      <div class="w-8 h-8 rounded-full bg-green-100 flex items-center justify-center ml-2">
                        <svg class="w-4 h-4 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a2 2 0 002-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z"></path>
                        </svg>
                      </div>
                      <div class="flex-1">
                        <div class="text-xs text-gray-500">موبایل</div>
                        <div class="text-sm font-medium text-gray-900">{{ selectedQuestion.customer_mobile || 'نامشخص' }}</div>
                      </div>
                    </div>
                    
                <div class="bg-white rounded-lg p-3 flex items-center shadow-sm">
                  <div class="w-8 h-8 rounded-full bg-indigo-100 flex items-center justify-center ml-2">
                    <svg class="w-4 h-4 text-indigo-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 15a4 4 0 004 4h10a4 4 0 004-4M7 10a5 5 0 1110 0v2H7v-2z" />
                    </svg>
                  </div>
                  <div class="flex-1 min-w-0">
                    <div class="text-xs text-gray-500">نام کاربر</div>
                    <div class="text-sm font-medium text-gray-900 truncate">{{ selectedQuestion.customer_name || selectedQuestion.customer?.name || 'کاربر مهمان' }}</div>
                  </div>
                </div>

                    <div class="bg-white rounded-lg p-3 flex items-center shadow-sm">
                      <div class="w-8 h-8 rounded-full bg-purple-100 flex items-center justify-center ml-2">
                        <svg class="w-4 h-4 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                        </svg>
                      </div>
                      <div class="flex-1">
                        <div class="text-xs text-gray-500">تعداد پرسش</div>
                        <div class="text-sm font-medium text-gray-900">1</div>
                      </div>
                    </div>
                    
                    <div class="bg-white rounded-lg p-3 flex items-center shadow-sm">
                      <div class="w-8 h-8 rounded-full bg-orange-100 flex items-center justify-center ml-2">
                        <svg class="w-4 h-4 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3a1 1 0 011-1h6a1 1 0 011 1v4m4 0V9a2 2 0 00-2-2H6a2 2 0 00-2 2v9a2 2 0 002 2h12a2 2 0 002-2V9a2 2 0 00-2-2m-6 0V3"></path>
                        </svg>
                      </div>
                      <div class="flex-1">
                        <div class="text-xs text-gray-500">تاریخ پرسش</div>
                        <div class="text-sm font-medium text-gray-900">{{ formatDate(selectedQuestion.created_at) || formatDate(selectedQuestion.createdAt) || 'نامشخص' }}</div>
                      </div>
                    </div>

                <div class="bg-white rounded-lg p-3 flex items-center shadow-sm">
                  <div class="w-8 h-8 rounded-full bg-sky-100 flex items-center justify-center ml-2">
                    <svg class="w-4 h-4 text-sky-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a2 2 0 01-2.828 0l-4.243-4.243a8 8 0 1111.314 0z" />
                    </svg>
                  </div>
                  <div class="flex-1">
                    <div class="text-xs text-gray-500">آی‌پی</div>
                    <div class="text-sm font-medium text-gray-900">{{ selectedQuestion.ip_address || 'نامشخص' }}</div>
                  </div>
                </div>

                <div class="bg-white rounded-lg p-3 flex items-center shadow-sm md:col-span-2">
                  <div class="w-8 h-8 rounded-full bg-emerald-100 flex items-center justify-center ml-2">
                    <svg class="w-4 h-4 text-emerald-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.75 17L15 12l-5.25-5" />
                    </svg>
                  </div>
                  <div class="flex-1 min-w-0">
                    <div class="text-xs text-gray-500">دستگاه / مرورگر</div>
                    <div class="text-sm font-medium text-gray-900 truncate" :title="selectedQuestion?.user_agent || (selectedQuestion?.device_info && JSON.stringify(selectedQuestion.device_info)) || ''">
                      {{ selectedQuestion?.user_agent || 'نامشخص' }}
                    </div>
                  </div>
                </div>
                    
                    <div v-if="selectedQuestion.customer_id" class="bg-white rounded-lg p-3 flex items-center shadow-sm">
                      <div class="w-8 h-8 rounded-full bg-blue-100 flex items-center justify-center ml-2">
                        <svg class="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3a1 1 0 011-1h6a1 1 0 011 1v4m4 0V9a2 2 0 00-2-2H6a2 2 0 00-2 2v9a2 2 0 002 2h12a2 2 0 002-2V9a2 2 0 00-2-2m-6 0V3"></path>
                        </svg>
                      </div>
                      <div class="flex-1">
                        <div class="text-xs text-gray-500">تاریخ عضویت</div>
                        <div class="text-sm font-medium text-gray-900">{{ formatDate(selectedQuestion.customer?.joinDate) || 'نامشخص' }}</div>
                      </div>
                    </div>
                    
                    <div v-else class="bg-white rounded-lg p-3 flex items-center shadow-sm">
                      <div class="w-8 h-8 rounded-full bg-gray-100 flex items-center justify-center ml-2">
                        <svg class="w-4 h-4 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
                        </svg>
                      </div>
                      <div class="flex-1">
                        <div class="text-xs text-gray-500">وضعیت</div>
                        <div class="text-sm font-medium text-gray-900">کاربر مهمان</div>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="bg-gradient-to-br from-emerald-50 to-green-100 rounded-xl p-6 border border-emerald-200 shadow-sm">
                  <div class="flex items-center gap-3 mb-4">
                    <div class="w-10 h-10 rounded-lg bg-emerald-500 flex items-center justify-center">
                      <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
                      </svg>
                    </div>
                    <h4 class="text-lg font-bold text-gray-900">اطلاعات محصول</h4>
                  </div>
                  
                  <div v-if="selectedQuestion.product" class="space-y-3">
                    <div class="bg-white rounded-lg p-3 border border-emerald-200">
                      <div class="flex items-center justify-between">
                        <div class="flex-1">
                          <div class="text-xs text-emerald-600 mb-1">نام محصول</div>
                          <NuxtLink :to="`/admin/product-management/products/edit?id=${selectedQuestion.product.id}`" class="text-base font-bold text-blue-600 hover:text-blue-800 hover:underline transition-colors">
                            {{ selectedQuestion.product.name }}
                          </NuxtLink>
                        </div>
                        <svg class="w-4 h-4 text-blue-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"></path>
                        </svg>
                      </div>
                    </div>
                    
                    <div class="grid grid-cols-2 gap-3">
                      <div class="bg-white rounded-lg p-3 border border-emerald-200">
                        <div class="text-xs text-emerald-600 mb-1">کد محصول</div>
                        <div class="text-sm font-bold text-gray-900">#{{ selectedQuestion.product.id }}</div>
                      </div>
                      
                      <div class="bg-white rounded-lg p-3 border border-emerald-200">
                        <div class="text-xs text-emerald-600 mb-1">قیمت</div>
                        <div class="text-sm font-bold text-green-700">{{ selectedQuestion.product.price.toLocaleString() }} تومان</div>
                      </div>
                    </div>
                  </div>
                  
                  <div v-else class="text-center py-4">
                    <div class="w-12 h-12 rounded-lg bg-gray-200 flex items-center justify-center mx-auto mb-2">
                      <svg class="w-6 h-6 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2 2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4"></path>
                      </svg>
                    </div>
                    <div class="text-sm text-gray-500">بدون محصول</div>
                  </div>
                </div>
              </div>

              <div class="bg-gradient-to-br from-amber-50 to-orange-100 rounded-lg p-6 border border-amber-200 shadow-sm">
                <div class="flex items-center justify-between mb-3">
                  <div class="flex items-center gap-3">
                    <h4 class="text-base font-semibold text-gray-900">پرسش مشتری</h4>
                    <span :class="getStatusClass(selectedQuestion.status)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-lg shadow-sm ml-2">
                      {{ getStatusLabel(selectedQuestion.status) }}
                    </span>
                  </div>
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <button v-if="!isEditingQuestionInline" class="inline-flex items-center gap-1 px-3 py-1 text-xs font-bold text-white rounded-lg shadow-lg bg-gradient-to-r from-yellow-400 to-orange-500 hover:from-yellow-500" @click="startInlineQuestionEdit"><svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.536L16.732 3.732z"></path></svg><span>ویرایش</span></button>
                    <span :class="getCategoryClass(selectedQuestion.category)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-lg shadow-sm">{{ getCategoryLabel(selectedQuestion.category) }}</span>
                  </div>
                </div>
                
                <!-- Static question display -->
                <div v-if="!isEditingQuestionInline" class="mt-3 bg-white/50 border border-amber-200 rounded-xl p-6 mb-4">
                  <p class="text-base text-gray-900 leading-relaxed">"{{ selectedQuestion.question }}"</p>
                </div>
                
                <!-- Inline question editing -->
                <div v-if="isEditingQuestionInline" class="mt-3 bg-white/70 border border-amber-300 rounded-xl p-6 mb-4">
                  <textarea 
                    v-model="inlineQuestionText" 
                    rows="4" 
                    class="block w-full border-amber-300 rounded-lg shadow-sm p-3 text-base text-gray-900 leading-relaxed resize-none focus:border-amber-500 focus:ring-amber-500" 
                    placeholder="متن پرسش مشتری را ویرایش کنید..."
                    dir="rtl"
                  ></textarea>
                  <div class="flex justify-end gap-2 mt-3">
                    <button class="px-4 py-2 rounded-lg text-sm bg-gray-200 hover:bg-gray-300 text-gray-700" @click="cancelInlineQuestionEdit">انصراف</button>
                    <button class="px-4 py-2 rounded-lg text-sm text-white bg-green-600 hover:bg-green-700" @click="saveInlineQuestionEdit">ذخیره</button>
                  </div>
                </div>

                <div v-if="isAnsweringInline" class="mt-4 p-6 bg-green-50 rounded-xl border border-green-200">
                  <h5 class="text-lg font-bold text-green-900 mb-3">{{ selectedQuestion.answer ? 'ویرایش پاسخ' : 'ثبت پاسخ جدید' }}</h5>
                  <textarea v-model="inlineAnswerText" rows="5" class="block w-full border-green-300 rounded-lg shadow-sm p-3" placeholder="پاسخ خود را اینجا بنویسید..."></textarea>
                  <div class="flex justify-end gap-2 mt-3">
                    <button class="px-4 py-2 rounded-lg text-sm bg-gray-200 hover:bg-gray-300" @click="isAnsweringInline = false">انصراف</button>
                    <button class="px-4 py-2 rounded-lg text-sm text-white bg-green-600 hover:bg-green-700" @click="submitInlineAnswer">ارسال پاسخ</button>
                  </div>
                </div>

                <div v-if="selectedQuestion.answer && !isAnsweringInline" class="p-6 bg-green-50 rounded-xl border border-green-200">
                  <h5 class="text-base font-bold text-green-900 mb-2">پاسخ مدیریت</h5>
                  <div class="bg-white/70 border border-green-300 rounded-xl p-6"><p class="text-gray-900 leading-relaxed">{{ selectedQuestion.answer }}</p></div>
                </div>
              </div>

              <div class="bg-slate-50 rounded-xl border border-slate-200 shadow-sm">
                <button class="w-full p-6 flex items-center justify-between cursor-pointer" @click="isAdditionalInfoOpen = !isAdditionalInfoOpen">
                  <h4 class="text-base font-semibold text-gray-900">اطلاعات اضافی</h4>
                  <svg class="w-5 h-5 text-gray-600 transition-transform" :class="{'rotate-180': isAdditionalInfoOpen}" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path></svg>
                </button>
                                  <div v-if="isAdditionalInfoOpen" class="p-6 border-t border-slate-200">
                    <div class="grid grid-cols-1 md:grid-cols-2 gap-6 text-sm">
                        <div class="flex items-center">
                          <!-- وضعیت فعلی - سمت راست -->
                          <div class="flex-1 flex items-center justify-end gap-12">
                            <span class="text-gray-500 whitespace-nowrap">وضعیت فعلی:</span>
                            <span :class="getStatusClass(selectedQuestion.status)" class="px-2 py-1 rounded-full whitespace-nowrap">{{ getStatusLabel(selectedQuestion.status) }}</span>
                            <span class="text-gray-500">|</span>
                            <span class="text-gray-500 whitespace-nowrap">اولویت:</span>
                            <span :class="getPriorityClass(selectedQuestion.priority)" class="px-2 py-1 rounded-full whitespace-nowrap ml-1" style="margin-right: 0; margin-left: 5px;">{{ getPriorityLabel(selectedQuestion.priority) }}</span>
                            <span class="text-gray-500">|</span>
                            <span class="text-gray-500 whitespace-nowrap">آدرس IP:</span>
                            <span class="font-medium ml-2 whitespace-nowrap">{{ selectedQuestion.ip_address || 'نامشخص' }}</span>
                            <span class="text-gray-500">|</span>

                          </div>
                          
                          <!-- رهگیری IP - کاملاً چپ -->
                          <div class="flex-1 flex justify-start">
                            <button v-if="selectedQuestion.ip_address" class="bg-blue-500 hover:bg-blue-600 text-white text-xs px-3 py-1 rounded-lg transition-colors duration-200 flex items-center gap-1 whitespace-nowrap ml-4"   @click="trackIP(selectedQuestion.ip_address)"
                            >
                              <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"></path>
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"></path>
                              </svg>
                              رهگیری IP
                            </button>
                          </div>
                        </div>
                      <div v-if="selectedQuestion.isAnonymous" class="md:col-span-2 text-orange-700">این کاربر درخواست ناشناس بودن کرده است.</div>
                  </div>
                </div>
              </div>

              <div class="bg-blue-50 rounded-xl border border-blue-200 shadow-sm">
                <button class="w-full p-6 flex items-center justify-between cursor-pointer" @click="isStatusSectionOpen = !isStatusSectionOpen">
                  <h4 class="text-base font-semibold text-gray-900">تغییر وضعیت</h4>
                  <svg class="w-5 h-5 text-gray-600 transition-transform" :class="{'rotate-180': isStatusSectionOpen}" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path></svg>
                </button>
                <div v-if="isStatusSectionOpen" class="p-6 border-t border-blue-200">
                  <div class="flex flex-wrap gap-3">
                    <button :disabled="selectedQuestion.status === 'pending'" :class="selectedQuestion.status === 'pending' ? 'bg-yellow-400 text-yellow-900 border-2 border-yellow-500 font-bold' : 'bg-amber-100 text-amber-800 hover:bg-amber-200 border border-amber-300'" class="flex-1 btn-status-no-shadow disabled:opacity-50 transition-all duration-150" @click="changeQuestionStatus('pending')"><svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>در انتظار پاسخ</button>
                    <button :disabled="selectedQuestion.status === 'reviewing'" :class="selectedQuestion.status === 'reviewing' ? 'bg-cyan-600 text-white border-2 border-cyan-700 font-bold' : 'bg-cyan-100 text-cyan-800 hover:bg-cyan-200 border border-cyan-300'" class="flex-1 btn-status-no-shadow disabled:opacity-50 transition-all duration-150" @click="changeQuestionStatus('reviewing')"><svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path></svg>در حال بررسی</button>
                    <button :disabled="selectedQuestion.status === 'answered' || !selectedQuestion.answer" :class="selectedQuestion.status === 'answered' ? 'bg-green-600 text-white border-2 border-green-700 font-bold' : 'bg-green-100 text-green-800 hover:bg-green-200 border border-green-300'" class="flex-1 btn-status-no-shadow disabled:opacity-50 transition-all duration-150" @click="changeQuestionStatus('answered')"><svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg>پاسخ داده شده</button>
                    <button :disabled="selectedQuestion.status === 'rejected'" :class="selectedQuestion.status === 'rejected' ? 'bg-red-600 text-white border-2 border-red-700 font-bold' : 'bg-red-100 text-red-800 hover:bg-red-200 border border-red-300'" class="flex-1 btn-status-no-shadow disabled:opacity-50 transition-all duration-150" @click="changeQuestionStatus('rejected')"><svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path></svg>رد شده</button>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <div class="bg-gray-50 px-4 py-3 border-t">
            <div class="flex gap-2 justify-end">
              <button v-if="!isAnsweringInline && selectedQuestion && selectedQuestion.status !== 'rejected'" class="btn bg-green-600 text-white hover:bg-green-700" @click="startInlineAnswer">{{ selectedQuestion && selectedQuestion.answer ? 'ویرایش پاسخ' : 'پاسخ دادن' }}</button>
              <button class="btn bg-gray-200 text-gray-700 hover:bg-gray-300" @click="closeDetailModal">بستن</button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div v-if="showEditModal" class="fixed inset-0 z-50 overflow-y-auto">
      <div class="flex items-center justify-center min-h-screen p-6">
        <div class="fixed inset-0 bg-gray-500 bg-opacity-75" @click="showEditModal = false"></div>
        <div class="relative bg-white rounded-lg shadow-xl w-full max-w-2xl" dir="rtl">
          <div class="p-6 border-b"><h3 class="text-lg font-medium">ویرایش پرسش و پاسخ</h3></div>
          <div v-if="editQuestionData" class="p-6 space-y-4">
            <div><label class="block text-sm font-bold text-gray-700 mb-1">متن پرسش</label><textarea v-model="editQuestionData.question" rows="4" class="form-input"></textarea></div>
            <div><label class="block text-sm font-bold text-purple-700 mb-1">پاسخ مدیریت</label><textarea v-model="editQuestionData.answer" rows="4" class="form-input"></textarea></div>
            <div class="flex gap-6">
              <div class="flex-1"><label class="block text-sm font-bold text-gray-700 mb-1">وضعیت</label>
                <select v-model="editQuestionData.status" class="form-input">
                  <option value="pending">در انتظار پاسخ</option>
                  <option value="reviewing">در حال بررسی</option>
                  <option value="answered">پاسخ داده شده</option>
                  <option value="rejected">رد شده</option>
                </select>
              </div>
              <div class="flex-1"><label class="block text-sm font-bold text-gray-700 mb-1">اولویت</label>
                <select v-model="editQuestionData.priority" class="form-input">
                  <option value="low">پایین</option>
                  <option value="medium">متوسط</option>
                  <option value="high">فوری</option>
                </select>
              </div>
            </div>
          </div>
          <div class="p-6 bg-gray-50 flex justify-end gap-3">
            <button class="btn bg-green-600 text-white hover:bg-green-700" @click="saveEdit">ذخیره</button>
            <button class="btn bg-gray-200 text-gray-700 hover:bg-gray-300" @click="showEditModal = false">انصراف</button>
          </div>
        </div>
      </div>
    </div>

    <!-- Categories Management Modal -->
    <div v-if="showCategoriesModal" class="fixed inset-0 z-50 overflow-y-auto">
      <div class="flex items-center justify-center min-h-screen p-6">
        <div class="fixed inset-0 bg-gray-500 bg-opacity-75" @click="showCategoriesModal = false"></div>
        <div class="relative bg-gradient-to-br from-purple-50 to-pink-50 rounded-2xl shadow-2xl w-full max-w-4xl border border-purple-200" dir="rtl">
          <div class="p-6 border-b border-purple-200 flex items-center gap-2">
            <svg class="w-6 h-6 text-purple-400" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"></path>
            </svg>
            <h3 class="text-lg font-bold text-purple-900">مدیریت دسته‌بندی‌های پرسش و پاسخ</h3>
          </div>
          <div class="p-6">
            <!-- Add New Category Form -->
            <div class="bg-gradient-to-r from-purple-100 to-pink-100 rounded-xl p-6 mb-6 border border-purple-200 shadow flex flex-col gap-2">
              <h4 class="text-md font-bold text-purple-800 mb-4 flex items-center gap-2">
                <svg class="w-5 h-5 text-purple-400" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path></svg>
                افزودن دسته‌بندی جدید
              </h4>
              <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
                <div>
                  <label class="block text-sm font-bold text-purple-700 mb-1">کلید (انگلیسی)</label>
                  <input v-model="newCategory.key" type="text" class="form-input rounded-lg border-purple-300 focus:border-purple-500 focus:ring-purple-500" placeholder="technical" />
                </div>
                <div>
                  <label class="block text-sm font-bold text-purple-700 mb-1">نام (فارسی)</label>
                  <input v-model="newCategory.name" type="text" class="form-input rounded-lg border-purple-300 focus:border-purple-500 focus:ring-purple-500" placeholder="فنی" />
                </div>
                <div class="flex items-end">
                  <button class="inline-flex items-center gap-2 px-6 py-2 rounded-lg bg-gradient-to-r from-purple-500 to-pink-500 text-white font-bold shadow hover:from-purple-600 hover:to-pink-600 transition w-full" @click="createCategory">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path></svg>
                    افزودن
                  </button>
                </div>
              </div>
            </div>
            <!-- Categories List -->
            <div class="bg-white border border-purple-200 rounded-xl overflow-hidden shadow">
              <table class="min-w-full divide-y divide-purple-100">
                <thead class="bg-cyan-100">
                  <tr>
                    <th class="px-6 py-3 text-right text-xs font-bold text-cyan-700 uppercase tracking-wider">کلید</th>
                    <th class="px-6 py-3 text-right text-xs font-bold text-cyan-700 uppercase tracking-wider">نام</th>
                    <th class="px-6 py-3 text-right text-xs font-bold text-cyan-700 uppercase tracking-wider">تعداد سوالات</th>
                    <th class="px-6 py-3 text-right text-xs font-bold text-cyan-700 uppercase tracking-wider">عملیات</th>
                  </tr>
                </thead>
                <tbody class="bg-white divide-y divide-purple-50">
                  <tr v-for="category in categories" :key="category.key">
                    <td class="px-6 py-4 whitespace-nowrap text-sm font-bold text-purple-900">{{ category.key }}</td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-purple-700">{{ category.name }}</td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-purple-700">{{ category.question_count || 0 }}</td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                      <button 
                        :disabled="(category.question_count || 0) > 0" 
                        :class="(category.question_count || 0) > 0 ? 'text-gray-400 cursor-not-allowed' : 'inline-flex items-center gap-1 font-bold bg-gradient-to-r from-rose-400 to-red-400 text-white px-4 py-1 rounded-lg shadow hover:from-rose-500 hover:to-red-500 transition'"
                        @click="deleteCategory(category.key)"
                      >
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path></svg>
                        حذف
                      </button>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
          <div class="p-6 bg-gradient-to-r from-purple-50 to-pink-50 flex justify-end rounded-b-2xl border-t border-purple-100">
            <button class="inline-flex items-center gap-2 px-6 py-2 rounded-lg bg-gradient-to-r from-blue-500 to-cyan-500 text-white font-bold shadow hover:from-blue-600 hover:to-cyan-600 transition" @click="showCategoriesModal = false">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path></svg>
              بستن
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Settings Modal -->
    <ViewAllModal v-model="showSettingsModal" title="تنظیمات پرسش و پاسخ">
      <qa-settings @loading="settingsLoading = $event" />
    </ViewAllModal>

    <!-- Missing utility functions
    function exportQuestions() {
      // Export functionality - can be implemented later
      alert('قابلیت صادرات در حال توسعه است');
    }
    -->

  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
</script>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'
import ViewAllModal from '~/components/admin/modals/ViewAllModal.vue'
import ExportExcelButton from '~/components/common/ExportExcelButton.vue'
import QaSettings from './qa-settings.vue'

// Import useAuth for permission checking
// Auth disabled
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// Interfaces & Types
interface Customer { 
  name: string; 
  email: string; 
  phone?: string; 
  joinDate: string; 
  questionCount?: number; 
  deletedQuestionCount?: number;
}
interface Product { 
  id: string; 
  name: string; 
  image?: string; 
  price: number; 
}
type Status = 'pending' | 'reviewing' | 'answered' | 'rejected';
interface Question {
  id: string; customer: Customer; product: Product; question: string; category: string;
  status: Status; createdAt: string; answer?: string; answeredAt?: string; answeredBy?: string;
  ip_address?: string; priority: 'low' | 'medium' | 'high'; isAnonymous?: boolean;
  customer_name?: string; customer_mobile?: string; customer_id?: number | null;
  created_at?: string;
  user_agent?: string;
  device_info?: Record<string, unknown>;
}

// Component State
const questions = ref<Question[]>([]);
const selectedQuestions = ref<string[]>([]);
const showFilters = ref(false);
const showQuestionModal = ref(false);
const showCategoriesModal = ref(false);
const selectedQuestion = ref<Question | null>(null);
const currentPage = ref(1);
const itemsPerPage = ref(10);

// Modal-specific state
const isAnsweringInline = ref(false);
const inlineAnswerText = ref('');
const isAdditionalInfoOpen = ref(true);
const isStatusSectionOpen = ref(true);
const showEditModal = ref(false);
const editQuestionData = ref<Partial<Question>>({});

// Add inline question editing state
const isEditingQuestionInline = ref(false);
const inlineQuestionText = ref('');

// Categories State
const categories = ref<Array<{id?: number, key: string, name: string, question_count?: number}>>([]);
const newCategory = ref({ key: '', name: '' });
const showTransferDropdown = ref(false);

// Filters & Stats
const filters = ref({ search: '', status: '', category: '', dateRange: '' });
// حذف useQAStats و fetchStats
// const { stats, fetchStats } = useQAStats();

// آمار را مستقیماً از questions محاسبه کن
const stats = computed(() => ({
  total: questions.value.length,
  pending: questions.value.filter(q => q.status === 'pending').length,
  answered: questions.value.filter(q => q.status === 'answered').length,
  rejected: questions.value.filter(q => q.status === 'rejected').length,
  reviewing: questions.value.filter(q => q.status === 'reviewing').length,
}));

// Computed Properties
const allSelected = computed(() => paginatedQuestions.value.length > 0 && selectedQuestions.value.length === paginatedQuestions.value.length);

const filteredQuestions = computed(() => {
  return questions.value
      .filter(q => {
        const search = filters.value.search.toLowerCase();
        const statusMatch = !filters.value.status || q.status === filters.value.status;
        const categoryMatch = !filters.value.category || q.category === filters.value.category;
        const searchMatch = !search ||
            q.customer.name.toLowerCase().includes(search) ||
            (q.product && q.product.name.toLowerCase().includes(search)) ||
            q.question.toLowerCase().includes(search);
        // Date range filtering can be added here if needed
        return statusMatch && categoryMatch && searchMatch;
      })
      .sort((a, b) => new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime());
});

const totalPages = computed(() => Math.ceil(filteredQuestions.value.length / itemsPerPage.value));
const paginatedQuestions = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage.value;
  return filteredQuestions.value.slice(start, start + itemsPerPage.value);
});

const visiblePages = computed(() => {
  const total = totalPages.value;
  const current = currentPage.value;
  if (total <= 7) return Array.from({ length: total }, (_, i) => i + 1);
  if (current < 5) return [1, 2, 3, 4, 5, '...', total];
  if (current > total - 4) return [1, '...', total - 4, total - 3, total - 2, total - 1, total];
  return [1, '...', current - 1, current, current + 1, '...', total];
});

const exportQuestionsData = computed(() => questions.value.map(q => ({
  'شناسه': q.id,
  'نام مشتری': q.customer?.name || '',
  'موبایل مشتری': q.customer?.phone || q.customer_mobile || '',
  'ایمیل مشتری': q.customer?.email || '',
  'محصول': q.product?.name || '',
  'شناسه محصول': q.product?.id || '',
  'دسته‌بندی': getCategoryLabel(q.category),
  'متن سوال': q.question,
  'پاسخ': q.answer || '',
  'وضعیت': getStatusLabel(q.status),
  'اولویت': getPriorityLabel(q.priority),
  'تاریخ ثبت': formatDate(q.createdAt),
  'ساعت ثبت': formatTime(q.createdAt),
  'آی‌پی': q.ip_address || ''
})));

// Data Fetching & Handling
async function loadData() {
  await Promise.all([loadQuestions(), loadCategories()]);
}

async function loadQuestions() {
  try {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const data = await $fetch<any[]>('/api/questions/admin');
    questions.value = data.map((q): Question => {
      const mapped = {
        ...q,
        id: q.id.toString(),
        status: q.status || 'pending', // Ensure status is always mapped
        customer_name: q.customer_name || 'کاربر مهمان',
        customer_mobile: q.customer_mobile || '',
        customer_id: q.customer_id,
        created_at: q.created_at,
        createdAt: q.created_at,
        product: q.product_name ? { 
          id: q.product_id, 
          name: q.product_name, 
          price: q.product_price || 0, 
          image: '' 
        } : null,
        customer: { 
          name: q.customer_name || 'کاربر مهمان', 
          email: 'نامشخص', 
          phone: q.customer_mobile || 'نامشخص',
          joinDate: q.created_at || new Date().toISOString(),
          deletedQuestionCount: 0
        },
      };
      return mapped;
    });
  } catch (error) {
    // Error loading questions
    questions.value = [];
  }
}

async function loadCategories() {
  try {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    categories.value = await $fetch<any[]>('/api/categories-qa');
  } catch (error) {
    // Error loading categories
  }
}

// UI Methods
const formatDate = (d?: string) => d ? new Date(d).toLocaleDateString('fa-IR', { year: 'numeric', month: 'long', day: 'numeric' }) : '';
const formatTime = (d?: string) => d ? new Date(d).toLocaleTimeString('fa-IR', { hour: '2-digit', minute: '2-digit' }).replace(/\d/g, c => '۰۱۲۳۴۵۶۷۸۹'[+c]) : '';
const truncateText = (t: string, len: number) => t && t.length > len ? `${t.substring(0, len)}...` : t;

const statusMap: Record<Status, { label: string; class: string }> = {
  pending: { label: 'در انتظار پاسخ', class: 'bg-yellow-100 text-yellow-800' },
  reviewing: { label: 'در حال بررسی', class: 'bg-cyan-100 text-cyan-800' },
  answered: { label: 'پاسخ داده شده', class: 'bg-green-100 text-green-800' },
  rejected: { label: 'رد شده', class: 'bg-red-100 text-red-800' },
};
const getStatusLabel = (status: Status) => statusMap[status]?.label || status;
const getStatusClass = (status: Status) => statusMap[status]?.class || 'bg-gray-100 text-gray-800';

const getCategoryLabel = (val: string | number | null | undefined) => {
  if (val === null || typeof val === 'undefined' || val === '') return 'بدون دسته'
  // اگر مقدار عددی باشد، بر اساس شناسه جستجو کن
  const isNumeric = typeof val === 'number' || (/^\d+$/.test(String(val)))
  const cat = isNumeric
    ? categories.value.find(c => c.id === Number(val))
    : categories.value.find(c => c.key === String(val))
  return cat?.name || String(val)
}
const getCategoryClass = (val: string | number | null | undefined) => {
  if (val === null || typeof val === 'undefined' || val === '') return 'bg-gray-100 text-gray-800'
  const key = String(val)
  let hash = 0
  for (let i = 0; i < key.length; i++) { hash = key.charCodeAt(i) + ((hash << 5) - hash) }
  const colors = ['bg-blue-100 text-blue-800', 'bg-purple-100 text-purple-800', 'bg-orange-100 text-orange-800', 'bg-pink-100 text-pink-800']
  return colors[Math.abs(hash % colors.length)]
}

const toggleSelectAll = (e: Event) => {
  selectedQuestions.value = (e.target as HTMLInputElement).checked ? paginatedQuestions.value.map(q => q.id) : [];
};

// Modal Controls

function closeDetailModal() { showQuestionModal.value = false; isAnsweringInline.value = false; isAdditionalInfoOpen.value = true; isStatusSectionOpen.value = true; }
function trackIP(ipAddress: string) {
  // Copy IP to clipboard
  if (navigator.clipboard) {
    navigator.clipboard.writeText(ipAddress);
  } else {
    // Fallback for older browsers
    const textarea = document.createElement('textarea');
    textarea.value = ipAddress;
    document.body.appendChild(textarea);
    textarea.select();
    document.execCommand('copy');
    document.body.removeChild(textarea);
  }
  // Open check-host.net with IP as parameter
  window.open(`https://check-host.net/ip-info?host=${encodeURIComponent(ipAddress)}`, '_blank');
}
function startInlineAnswer() { if (selectedQuestion.value) { isAnsweringInline.value = true; inlineAnswerText.value = selectedQuestion.value.answer || ''; } }

// Inline Question Editing Functions
function startInlineQuestionEdit() { 
  if (selectedQuestion.value) { 
    isEditingQuestionInline.value = true; 
    inlineQuestionText.value = selectedQuestion.value.question || ''; 
  } 
}

function cancelInlineQuestionEdit() {
  isEditingQuestionInline.value = false;
  inlineQuestionText.value = '';
}

async function saveInlineQuestionEdit() {
  const q = selectedQuestion.value;
  if (!q || !inlineQuestionText.value.trim()) return;
  
  try {
    await $fetch(`/api/questions/${q.id}`, {
      method: 'PUT',
      body: {
        question: inlineQuestionText.value.trim(),
        status: q.status,
        category_id: resolveCategoryId(q.category),
        priority: q.priority,
        answer: q.answer,
        product_id: q.product?.id,
        customer_id: q.customer_id,
        customer_name: q.customer_name,
        customer_mobile: q.customer_mobile,
        is_anonymous: q.isAnonymous
      }
    });
    
    isEditingQuestionInline.value = false;
    await loadQuestions(); // Reload all questions
    selectedQuestion.value = questions.value.find(item => item.id === q.id) || null; // Re-select the updated question
  } catch (error) { 
    // Error saving question edit
    // useNotifier().error('خطا در ذخیره تغییرات');
  }
}

// CRUD Operations
async function submitInlineAnswer() {
  const q = selectedQuestion.value;
  if (!q || !inlineAnswerText.value.trim()) return;
  try {
    await $fetch(`/api/questions/${q.id}`, {
      method: 'PUT',
      body: {
        question: q.question,
        status: 'answered',
        category_id: resolveCategoryId(q.category),
        priority: q.priority,
        answer: inlineAnswerText.value.trim(),
        product_id: q.product?.id,
        customer_id: q.customer_id,
        customer_name: q.customer_name,
        customer_mobile: q.customer_mobile,
        is_anonymous: q.isAnonymous,
        answeredAt: new Date().toISOString(),
        answeredBy: 'مدیر سیستم'
      }
    });
    isAnsweringInline.value = false;
    await loadQuestions(); // Reload all questions
    selectedQuestion.value = questions.value.find(item => item.id === q.id) || null; // Re-select the updated question
  } catch (error) { console.error('Error submitting answer:', error); }
  // Error submitting answer handled above
}

async function changeQuestionStatus(status: Status) {
  const q = selectedQuestion.value;
  if (!q || q.status === status) return;
  if (status === 'answered' && !q.answer) { 
    // useNotifier().warning('ابتدا باید پاسخی برای سوال ثبت کنید.'); 
    alert('ابتدا باید پاسخی برای سوال ثبت کنید.');
    return; 
  }
  {
    // const { confirm } = useConfirmDialog();
    // const ok = await confirm({ title:'تأیید عملیات', message:`آیا از تغییر وضعیت به «${getStatusLabel(status)}» اطمینان دارید؟`, confirmText:'تأیید', cancelText:'انصراف', type:'warning' });
    const ok = confirm(`آیا از تغییر وضعیت به «${getStatusLabel(status)}» اطمینان دارید؟`);
    if (!ok) return;
  }

  try {
    await $fetch(`/api/questions/${q.id}`, {
      method: 'PUT',
      body: {
        question: q.question,
        status,
        category_id: resolveCategoryId(q.category),
        priority: q.priority,
        answer: q.answer,
        product_id: q.product?.id,
        customer_id: q.customer_id,
        customer_name: q.customer_name,
        customer_mobile: q.customer_mobile,
        is_anonymous: q.isAnonymous
      }
    });
    await loadQuestions(); // Reload all questions
    selectedQuestion.value = questions.value.find(item => item.id === q.id) || null; // Re-select the updated question
    // useNotifier().success('وضعیت با موفقیت تغییر کرد.');
    alert('وضعیت با موفقیت تغییر کرد.');
  } catch (error) { console.error('Error changing status:', error); }
  // Error changing status handled above
}

async function deleteQuestion(id: string) {
  const ok = confirm('آیا از حذف این پرسش اطمینان دارید؟');
  if (ok) {
    try {
       
      await $fetch(`/api/questions/${id}`, { method: 'DELETE' as const });
      await loadData();
    } catch (error) { console.error('Error deleting question:', error); }
    // Error deleting question handled above
  }
}

// Edit Modal
function openEditModal(q: Question) { editQuestionData.value = JSON.parse(JSON.stringify(q)); showEditModal.value = true; }

async function saveEdit() {
  const qd = editQuestionData.value;
  if (!qd || !qd.id) return;
  try {
     
    await $fetch(`/api/questions/${qd.id}`, { method: 'PUT', body: qd });
    showEditModal.value = false;
    await loadQuestions(); // Reload all questions
    // If the edited question was the one being viewed, update it
    if (selectedQuestion.value?.id === qd.id) {
      selectedQuestion.value = questions.value.find(item => item.id === qd.id) || null;
    }
  } catch (error) { console.error('Error saving edit:', error); }
  // Error saving edit handled above
}

// Bulk Actions
async function bulkReject() {
  if (!selectedQuestions.value.length) return;
  const ok = confirm(`آیا از رد کردن ${selectedQuestions.value.length} پرسش اطمینان دارید؟`);
   
  await Promise.all(selectedQuestions.value.map(id => $fetch(`/api/questions/${id}`, { method: 'PUT', body: { status: 'rejected' } })));
  selectedQuestions.value = [];
  await loadData();
}

async function bulkDelete() {
  if (!selectedQuestions.value.length) return;
  const ok = confirm(`آیا از حذف ${selectedQuestions.value.length} پرسش اطمینان دارید؟`);
  if (!ok) return;
   
  await Promise.all(selectedQuestions.value.map(id => $fetch(`/api/questions/${id}`, { method: 'DELETE' as const })));
  selectedQuestions.value = [];
  await loadData();
}

// Category Management Functions
async function createCategory() {
  if (!newCategory.value.key || !newCategory.value.name) {
    alert('لطفاً کلید و نام دسته‌بندی را وارد کنید');
    return;
  }
  
  try {
     
    await $fetch('/api/categories-qa', { method: 'POST', body: { key: newCategory.value.key, name: newCategory.value.name } });
    
    newCategory.value = { key: '', name: '' };
    await loadCategories();
    alert('دسته‌بندی با موفقیت اضافه شد');
  } catch (error) {
    // Error creating category
    alert('خطا در ایجاد دسته‌بندی');
  }
}

async function deleteCategory(key: string) {
  const ok = confirm('آیا از حذف این دسته‌بندی اطمینان دارید؟');
  if (!ok) return;
  
  try {
     
    await $fetch(`/api/categories-qa/${key}`, {
      method: 'DELETE' as const
    });
    
    await loadCategories();
    alert('دسته‌بندی با موفقیت حذف شد');
  } catch (error) {
    // Error deleting category
    alert('خطا در حذف دسته‌بندی');
  }
}

async function bulkTransferToCategory(categoryKey: string) {
  if (!selectedQuestions.value.length) return;
  
  const categoryName = categories.value.find(c => c.key === categoryKey)?.name;
  const ok = confirm(`آیا از انتقال ${selectedQuestions.value.length} پرسش انتخاب شده به دسته‌بندی "${categoryName}" اطمینان دارید؟`);
  if (!ok) return;
  
  try {
     
    await $fetch(`/api/admin/qa/questions/bulk-transfer`, {
      method: 'POST',
      body: {
        question_ids: selectedQuestions.value,
        to_category: categoryKey
      }
    });
    
    selectedQuestions.value = [];
    showTransferDropdown.value = false;
    await loadData();
    alert(`پرسش‌ها با موفقیت به دسته‌بندی "${categoryName}" منتقل شدند`);
  } catch (error) {
    // Error transferring questions
    alert('خطا در انتقال پرسش‌ها');
  }
}

// Missing utility functions


function setStatusFilter(status: string) {
  filters.value.status = status;
}





function previousPage() {
  if (currentPage.value > 1) {
    currentPage.value--;
  }
}

function nextPage() {
  if (currentPage.value < totalPages.value) {
    currentPage.value++;
  }
}

function goToPage(page: number) {
  if (page >= 1 && page <= totalPages.value) {
    currentPage.value = page;
  }
}

// Lifecycle and Watchers
watch(filters, () => { currentPage.value = 1; }, { deep: true });
onMounted(loadData);

function getPriorityLabel(priority: string) {
  switch(priority) {
    case 'high': return 'فوری';
    case 'medium': return 'متوسط';
    case 'low': return 'کم';
    default: return 'نامشخص';
  }
}
function getPriorityClass(priority: string) {
  switch(priority) {
    case 'high': return 'bg-red-100 text-red-800';
    case 'medium': return 'bg-yellow-100 text-yellow-800';
    case 'low': return 'bg-blue-100 text-blue-800';
    default: return 'bg-gray-100 text-gray-600';
  }
}

const showSettingsModal = ref(false)
const settingsLoading = ref(false)

// Helper: resolve category id from key or id
function resolveCategoryId(val: string | number): number | undefined {
  if (typeof val === 'number') return val
  const found = categories.value.find(c => c.key === String(val))
  return found?.id
}
</script>

<style scoped>
.line-clamp-2 { display: -webkit-box; -webkit-line-clamp: 2; line-clamp: 2; -webkit-box-orient: vertical; overflow: hidden; }
.form-input { 
  display: block; 
  width: 100%; 
  border-radius: 0.5rem; 
  box-shadow: 0 1px 2px 0 rgb(0 0 0 / 0.05); 
}
.btn { 
  padding: 0.5rem 1rem; 
  font-size: 0.875rem; 
  font-weight: 500; 
  border-radius: 0.5rem; 
  transition: color 0.2s, background-color 0.2s, border-color 0.2s, text-decoration-color 0.2s, fill 0.2s, stroke 0.2s, opacity 0.2s, box-shadow 0.2s, transform 0.2s, filter 0.2s, backdrop-filter 0.2s; 
}
.btn-status { 
  display: inline-flex; 
  align-items: center; 
  justify-content: center; 
  gap: 0.5rem; 
  padding: 0.5rem 1rem; 
  border-radius: 0.5rem; 
  font-size: 0.875rem; 
  font-weight: 700; 
}
.btn-status-no-shadow { 
  display: inline-flex; 
  align-items: center; 
  justify-content: center; 
  gap: 0.5rem; 
  padding: 0.5rem 1rem; 
  border-radius: 0.5rem; 
  font-size: 0.875rem; 
  font-weight: 700; 
  box-shadow: none !important; 
}
/* Add or update the style for the question detail modal */
.qna-modal, .modal, .fixed.inset-0.z-50 {
  z-index: 9999;
  position: fixed;
}
</style>

