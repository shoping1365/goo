<template>

    <!-- Page Header -->
    <div class="bg-white shadow-sm border-b border-gray-200 mb-4">
      <div class="w-full px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center py-3">
          <div>
            <h1 class="text-lg font-bold text-gray-900">مدیریت نظرات محصولات</h1>
            <p class="text-xs text-gray-600 mt-1">مدیریت و بررسی نظرات مشتریان</p>
          </div>
          <div class="flex space-x-2 space-x-reverse">
            <button
              @click="showFilters = !showFilters"
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-cyan-400 to-cyan-600 hover:from-cyan-500 hover:to-cyan-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-cyan-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105"
            >
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.707A1 1 0 013 7V4z"></path>
              </svg>
              فیلترها
            </button>
            <button 
              @click="showSettingsModal = true" 
              class="inline-flex items-center px-3 py-2 border border-gray-200 rounded-lg bg-white hover:bg-gray-50 transition-all shadow-md"
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
            <button
              @click="exportReviews"
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-green-400 to-green-600 hover:from-green-500 hover:to-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105"
            >
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
              </svg>
              خروجی Excel
            </button>

          </div>
        </div>
      </div>
    </div>

    <div class="w-full px-4 py-4">

    <!-- Stats Cards -->
    <div class="mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-6 mb-6">
          <div
            class="cursor-pointer transition-all duration-200 hover:scale-105"
            @click="setStatusFilter('')"
          >
            <TemplateCard 
              title="کل نظرات" 
              :value="reviewsStats.total" 
              variant="blue"
            >
              <template #icon>
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
              </template>
            </TemplateCard>
          </div>

          <div
            class="cursor-pointer transition-all duration-200 hover:scale-105"
            @click="setStatusFilter('approved')"
          >
            <TemplateCard 
              title="تایید شده" 
              :value="reviewsStats.approved" 
              variant="green"
            >
              <template #icon>
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                </svg>
              </template>
            </TemplateCard>
          </div>

          <div
            class="cursor-pointer transition-all duration-200 hover:scale-105"
            @click="setStatusFilter('pending')"
          >
            <TemplateCard 
              title="در انتظار" 
              :value="reviewsStats.pending" 
              variant="amber"
            >
              <template #icon>
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
              </template>
            </TemplateCard>
          </div>

          <div
            class="cursor-pointer transition-all duration-200 hover:scale-105"
            @click="setStatusFilter('rejected')"
          >
            <TemplateCard 
              title="رد شده" 
              :value="reviewsStats.rejected" 
              variant="red"
            >
              <template #icon>
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                </svg>
              </template>
            </TemplateCard>
          </div>

          <div
            class="cursor-pointer transition-all duration-200 hover:scale-105"
            @click="setStatusFilter('average')"
          >
            <TemplateCard 
              title="میانگین امتیاز" 
              :value="reviewsStats.averageRating.toFixed(1)" 
              variant="purple"
            >
              <template #icon>
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z"></path>
                </svg>
              </template>
            </TemplateCard>
          </div>
        </div>
      </div>

      <!-- Filters -->
      <div v-if="showFilters" class="bg-gradient-to-br from-blue-50 to-cyan-50 shadow-lg rounded-2xl mb-8 border border-blue-100">
        <div class="px-6 py-4 border-b border-blue-100 flex items-center gap-2">
          <svg class="w-6 h-6 text-blue-400" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.707A1 1 0 013 7V4z"></path></svg>
          <h3 class="text-lg font-bold text-blue-900">فیلترها و جستجو</h3>
        </div>
        <div class="p-8">
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">جستجو در نظرات</label>
              <div class="relative">
                <input 
                  v-model="filters.search"
                  type="text"
                  class="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md leading-5 bg-white placeholder-gray-500 focus:outline-none focus:placeholder-gray-400 focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm"
                  placeholder="نام مشتری، محصول یا متن نظر..."
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
              <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت نظر</label>
              <select v-model="filters.status" class="block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-blue-500 focus:border-blue-500 text-sm">
                <option value="">همه وضعیت‌ها</option>
                <option value="pending">در انتظار بررسی</option>
                <option value="approved">تایید شده</option>
                <option value="rejected">رد شده</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">امتیاز</label>
              <select v-model="filters.rating" class="block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-white focus:outline-none focus:ring-blue-500 focus:border-blue-500 text-sm">
                <option value="">همه امتیازها</option>
                <option value="5">⭐⭐⭐⭐⭐ (5 ستاره)</option>
                <option value="4">⭐⭐⭐⭐ (4 ستاره)</option>
                <option value="3">⭐⭐⭐ (3 ستاره)</option>
                <option value="2">⭐⭐ (2 ستاره)</option>
                <option value="1">⭐ (1 ستاره)</option>
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
                <option value="6months">6 ماه گذشته</option>
                <option value="year">سال گذشته</option>
              </select>
            </div>
          </div>
          <div class="flex justify-end gap-6 mt-8">
          
          </div>
        </div>
      </div>

      <!-- Bulk Actions -->
      <div v-if="selectedReviews.length > 0" class="bg-blue-50 border border-blue-200 rounded-lg p-6 mb-6">
        <div class="flex items-center justify-between">
          <div class="flex items-center">
            <span class="text-sm font-medium text-blue-800">{{ selectedReviews.length }} نظر انتخاب شده</span>
          </div>
          <div class="flex space-x-2 space-x-reverse">
            <button 
              @click="bulkApprove"
              class="inline-flex items-center px-3 py-2 border border-transparent text-xs font-medium rounded text-white bg-green-600 hover:bg-green-700"
            >
              تایید همه
            </button>
            <button 
              @click="bulkReject"
              class="inline-flex items-center px-3 py-2 border border-transparent text-xs font-medium rounded text-white bg-red-600 hover:bg-red-700"
            >
              رد همه
            </button>
            <button 
              @click="bulkDelete"
              class="inline-flex items-center px-3 py-2 border border-transparent text-xs font-medium rounded text-white bg-gray-600 hover:bg-gray-700"
            >
              حذف همه
            </button>
          </div>
        </div>
      </div>

      <!-- Reviews Table -->
      <div class="bg-white shadow rounded-lg overflow-hidden">
        <div class="px-6 py-4 border-b border-gray-200">
          <div class="flex items-center justify-between">
            <h3 class="text-lg font-medium text-gray-900">لیست نظرات</h3>
            <div class="flex items-center space-x-2 space-x-reverse">
              <span class="text-sm text-gray-500">{{ paginatedReviews.length }} از {{ filteredReviews.length }} نظر</span>
            </div>
          </div>
        </div>

        <!-- Table -->
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                  <input 
                    type="checkbox" 
                    @change="toggleSelectAll" 
                    :checked="allSelected"
                    class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                  />
                </th>
                <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                  مشتری
                </th>
                <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                  محصول
                </th>
                <th scope="col" class="px-6 py-3 text-center text-xs font-bold tracking-wider">
                  امتیاز
                </th>
                <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                  نظر
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
              <tr v-for="review in paginatedReviews" :key="review.id" class="hover:bg-gray-50">
                <!-- Checkbox -->
                <td class="px-6 py-4 whitespace-nowrap">
                  <input 
                    type="checkbox" 
                    v-model="selectedReviews" 
                    :value="review.id"
                    class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                  />
                </td>

                <!-- Customer -->
                <td class="px-6 py-4 whitespace-nowrap text-right align-middle">
                  <div>
                    <div class="text-sm font-medium text-gray-900">{{ review.customer.name }}</div>
                    <div class="text-sm text-gray-500">{{ review.customer.email }}</div>
                  </div>
                </td>

                <!-- Product -->
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="flex items-center">
                    <div class="">
                      <img
                        :src="review.product.thumbnail || review.product.image"
                        alt="عکس محصول"
                        class="w-10 h-10 object-cover rounded-lg cursor-pointer hover:opacity-80 transition-opacity"
                        @error="e => ((e.target as HTMLImageElement).src = '/default-product.svg')"
                        @click="previewImage(getFullSizeImage(review.product), review.product.name)"
                      />
                    </div>
                    <div class="mr-4">
                      <a
                        :href="`/product/sku-${review.product.sku || review.product.id}`"
                        target="_blank"
                        rel="noopener noreferrer"
                        class="text-sm font-medium text-blue-700 hover:underline cursor-pointer"
                      >
                        {{ truncateText(review.product.name, 30) }}
                      </a>
                      <div class="text-sm text-gray-500">#{{ review.product.id }}</div>
                    </div>
                  </div>
                </td>

                <!-- Rating -->
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="flex items-center justify-center">
                    <span :style="getRatingStyle(review.rating)">
                      {{ review.rating }}
                    </span>
                  </div>
                </td>

                <!-- Comment -->
                <td class="px-6 py-4">
                  <div class="text-sm text-gray-900 max-w-xs" dir="rtl">
                    <p class="line-clamp-2">{{ truncateText(review.comment, 100) }}</p>
                    <div v-if="review.images && review.images.length > 0" class="flex space-x-1 space-x-reverse mt-2">
                      <img v-for="(image, index) in review.images.slice(0, 2)" :key="index" :src="image" class="h-6 w-6 rounded object-cover" />
                      <span v-if="review.images.length > 2" class="text-xs text-gray-500 flex items-center">+{{ review.images.length - 2 }}</span>
                    </div>
                  </div>
                </td>

                <!-- Date -->
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  <div>{{ formatDate(review.createdAt) }}</div>
                  <div class="text-xs text-gray-500">{{ formatTime(review.createdAt) }}</div>
                </td>

                <!-- Status -->
                <td class="px-6 py-4 whitespace-nowrap">
                  <span :class="getStatusClass(review.status)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                    {{ getStatusLabel(review.status) }}
                  </span>
                </td>

                <!-- Actions -->
                <td class="px-6 py-4 whitespace-nowrap text-left text-sm font-medium">
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <button 
                      @click="viewReview(review)"
                      class="text-blue-600 hover:text-blue-900 p-1 rounded-full hover:bg-blue-100"
                      title="مشاهده جزئیات"
                    >
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                      </svg>
                    </button>
                    <button 
                      v-if="review.status === 'pending' && hasPermission('review.approve')"
                      @click="approveReview(review.id)"
                      class="text-green-600 hover:text-green-900 p-1 rounded-full hover:bg-green-100"
                      title="تایید"
                    >
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                      </svg>
                    </button>
                    <button 
                      v-if="review.status === 'pending' && hasPermission('review.reject')"
                      @click="rejectReview(review.id)"
                      class="text-red-600 hover:text-red-900 p-1 rounded-full hover:bg-red-100"
                      title="رد"
                    >
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                      </svg>
                    </button>
                    
                    <button 
                      v-if="hasPermission('review.delete')"
                      @click="deleteReview(review.id)"
                      class="text-red-600 hover:text-red-900 p-1 rounded-full hover:bg-red-100"
                      title="حذف"
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

        <!-- Pagination -->
        <Pagination
          :current-page="currentPage"
          :total-pages="totalPages"
          :total="filteredReviews.length"
          @page-changed="goToPage"
          class="mt-6 mb-4"
        />
      </div>
    </div>

    <!-- Image Preview Modal -->
    <ImagePreviewModal
      v-if="showImagePreview && selectedImage"
      :show="showImagePreview"
      :image="selectedImage"
      @close="showImagePreview = false"
    />

    <!-- Review Detail Modal -->
    <div v-if="showReviewModal && selectedReview" class="fixed inset-0 z-50 overflow-y-auto">
      <div class="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
        <!-- Overlay fully removed -->
        <span class="hidden sm:inline-block sm:align-middle sm:h-screen">&#8203;</span>

        <div class="inline-block align-bottom bg-white rounded-xl text-right overflow-hidden shadow-2xl transform transition-all sm:my-8 sm:align-middle sm:max-w-7xl sm:w-full border border-indigo-100" dir="rtl">
          <div class="bg-gradient-to-r from-indigo-500 to-purple-600 px-4 py-4">
            <div class="flex items-center justify-between">
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-white/20 backdrop-blur-sm rounded-lg flex items-center justify-center">
                  <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V4a2 2 0 10-4 0v1.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"/></svg>
                </div>
                <div><h3 class="text-lg font-bold text-white">جزئیات نظر مشتری</h3></div>
              </div>
              <button @click="showReviewModal = false" class="text-white/80 hover:text-white hover:bg-white/10 rounded-lg p-1 transition-all">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path></svg>
              </button>
            </div>
          </div>

          <div class="px-4 pt-4 pb-2 max-h-[75vh] overflow-y-auto">
            <div v-if="selectedReview" class="space-y-4">
              <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 items-stretch">
                <div class="bg-gradient-to-br from-blue-50 to-indigo-100 rounded-xl p-6 border border-blue-200 shadow-sm">
                  <div class="flex items-center justify-between mb-4">
                    <div class="flex items-center">
                      <div class="w-12 h-12 rounded-full bg-blue-500 flex items-center justify-center ml-3">
                        <span class="text-white font-bold text-lg">{{ selectedReview.customer.name.charAt(0) }}</span>
                      </div>
                      <div>
                        <div class="text-sm font-semibold text-gray-900">
                          {{ selectedReview.customer.name }}
                        </div>
                        <div class="text-xs text-gray-500">
                          کاربر عضو
                        </div>
                      </div>
                    </div>
                  </div>
                  
                  <div class="grid grid-cols-2 gap-3">
                    <div class="bg-white rounded-lg p-3 flex items-center shadow-sm">
                      <div class="w-8 h-8 rounded-full bg-green-100 flex items-center justify-center ml-2">
                        <svg class="w-4 h-4 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z"></path>
                        </svg>
                      </div>
                      <div class="flex-1">
                        <div class="text-xs text-gray-500">موبایل</div>
                        <div class="text-sm font-medium text-gray-900">{{ selectedReview.customer.phone || 'نامشخص' }}</div>
                      </div>
                    </div>
                    
                    <div class="bg-white rounded-lg p-3 flex flex-col sm:flex-row items-center shadow-sm text-center sm:text-right mb-2">
                      <div class="w-8 h-8 rounded-full bg-orange-100 flex items-center justify-center mb-2 sm:mb-0 sm:ml-2">
                        <svg class="w-4 h-4 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3a1 1 0 011-1h6a1 1 0 011 1v4m4 0V9a2 2 0 00-2-2H6a2 2 0 00-2 2v9a2 2 0 002 2h12a2 2 0 002-2V9a2 2 0 00-2-2m-6 0V3"></path>
                        </svg>
                      </div>
                      <div class="flex-1">
                        <div class="text-xs text-gray-500">تاریخ نظر</div>
                        <div class="text-sm font-medium text-gray-900">{{ formatDate(selectedReview.createdAt) }}</div>
                      </div>
                    </div>
                    <div class="bg-white rounded-lg p-3 flex flex-col sm:flex-row items-center shadow-sm text-center sm:text-right mb-2">
                      <div class="w-8 h-8 rounded-full bg-purple-100 flex items-center justify-center mb-2 sm:mb-0 sm:ml-2">
                        <svg class="w-4 h-4 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                        </svg>
                      </div>
                      <div class="flex-1">
                        <div class="text-xs text-gray-500">تعداد نظرات</div>
                        <div class="text-sm font-medium text-gray-900">{{ selectedReview.customer.reviewCount || 0 }}</div>
                      </div>
                    </div>
                    <div class="bg-white rounded-lg p-3 flex flex-col sm:flex-row items-center shadow-sm text-center sm:text-right mb-2">
                      <div class="w-8 h-8 rounded-full bg-blue-100 flex items-center justify-center mb-2 sm:mb-0 sm:ml-2">
                        <svg class="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3a1 1 0 011-1h6a1 1 0 011 1v4m4 0V9a2 2 0 00-2-2H6a2 2 0 00-2 2v9a2 2 0 002 2h12a2 2 0 002-2V9a2 2 0 00-2-2m-6 0V3"></path>
                        </svg>
                      </div>
                      <div class="flex-1">
                        <div class="text-xs text-gray-500">تاریخ عضویت</div>
                        <div class="text-sm font-medium text-gray-900">{{ formatDate(selectedReview.customer.joinDate) }}</div>
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
                  
                  <div class="space-y-3">
                    <div class="bg-white rounded-lg p-3 border border-emerald-200">
                      <div class="flex items-center justify-between">
                        <div class="flex-1">
                          <div class="text-xs text-emerald-600 mb-1">نام محصول</div>
                          <NuxtLink :to="`/admin/product-management/products/edit?id=${selectedReview.product.id}`" class="text-base font-bold text-blue-600 hover:text-blue-800 hover:underline transition-colors">
                            {{ selectedReview.product.name }}
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
                        <div class="text-sm font-bold text-gray-900">#{{ selectedReview.product.id }}</div>
                      </div>
                      
                      <div class="bg-white rounded-lg p-3 border border-emerald-200">
                        <div class="text-xs text-emerald-600 mb-1">قیمت</div>
                        <div class="text-sm font-bold text-green-700">{{ selectedReview.product.price.toLocaleString() }} تومان</div>
                      </div>
                    </div>
                    
                    <!-- تصویر محصول حذف شد -->
                  </div>
                </div>
              </div>

              <div class="bg-gradient-to-br from-amber-50 to-orange-100 rounded-lg p-6 border border-amber-200 shadow-sm">
                <!-- ستاره‌ها بالای هدر و وسط (حذف شد) -->
                <div class="relative flex items-center justify-between mb-3">
                  <!-- عنوان راست -->
                  <div class="flex items-center gap-3">
                    <h4 class="text-base font-semibold text-gray-900">نظر و امتیاز مشتری</h4>
                    <span :class="getStatusClass(selectedReview.status)" class="inline-flex px-2 py-1 text-xs font-semibold rounded-lg shadow-sm ml-2">
                      {{ getStatusLabel(selectedReview.status) }}
                    </span>
                  </div>
                  <!-- ستاره‌ها وسط -->
                  <div class="absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 flex items-center">
                    <span v-for="n in 5" :key="n" :class="n <= selectedReview.rating ? 'text-yellow-400' : 'text-gray-300'" class="text-2xl">⭐</span>
                  </div>
                  <!-- دکمه پاسخ چپ -->
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <button @click="showReplyModal = true" class="inline-flex items-center gap-1 px-3 py-1 text-xs font-bold text-white rounded-lg shadow-lg bg-gradient-to-r from-blue-400 to-blue-600 hover:from-blue-500">
                      <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h10a8 8 0 018 8v2M3 10l6 6m-6-6l6-6"/>
                      </svg>
                      <span>پاسخ</span>
                    </button>
                  </div>
                </div>
                
                <!-- امتیاز (حذف شد) -->
                <!-- متن نظر -->
                <div class="bg-white/50 border border-amber-200 rounded-xl p-6 mb-4 min-h-32">
                  <p class="text-base text-gray-900 leading-relaxed">"{{ selectedReview.comment }}"</p>
                </div>

                <div v-if="isAnsweringInline" class="mt-4 p-6 bg-green-50 rounded-xl border border-green-200">
                  <h5 class="text-lg font-bold text-green-900 mb-3">{{ selectedReview.adminReply ? 'ویرایش پاسخ' : 'ثبت پاسخ جدید' }}</h5>
                  <textarea v-model="inlineAnswerText" rows="5" class="block w-full border-green-300 rounded-lg shadow-sm p-3" placeholder="پاسخ خود را اینجا بنویسید..."></textarea>
                  <div class="flex justify-end gap-2 mt-3">
                    <button @click="isAnsweringInline = false" class="px-4 py-2 rounded-lg text-sm bg-gray-200 hover:bg-gray-300">انصراف</button>
                    <button @click="submitInlineAnswer" class="px-4 py-2 rounded-lg text-sm text-white bg-green-600 hover:bg-green-700">ارسال پاسخ</button>
                  </div>
                </div>

                <div v-if="selectedReview.adminReply && !isAnsweringInline" class="p-6 bg-green-50 rounded-xl border border-green-200">
                  <h5 class="text-base font-bold text-green-900 mb-2">پاسخ مدیریت</h5>
                  <div class="bg-white/70 border border-green-300 rounded-xl p-6"><p class="text-gray-900 leading-relaxed">{{ selectedReview.adminReply }}</p></div>
                </div>
              </div>

              <!-- تصاویر ضمیمه -->
              <!-- این بخش به طور کامل حذف شد -->

              <div class="bg-slate-50 rounded-xl border border-slate-200 shadow-sm">
                <button @click="isAdditionalInfoOpen = !isAdditionalInfoOpen" class="w-full p-6 flex items-center justify-between cursor-pointer">
                  <h4 class="text-base font-semibold text-gray-900">اطلاعات اضافی</h4>
                  <svg class="w-5 h-5 text-gray-600 transition-transform" :class="{'rotate-180': isAdditionalInfoOpen}" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path></svg>
                </button>
                <div v-if="isAdditionalInfoOpen" class="p-6 border-t border-slate-200">
                  <div class="flex items-center justify-between gap-6">
                    <!-- اطلاعات اضافی: ۴ بخش با فاصله خیلی زیاد و جداکننده -->
                    <div class="flex items-center w-full gap-x-16">
                      <!-- وضعیت فعلی -->
                      <div class="flex items-center pl-16 border-l border-gray-200">
                        <span class="text-gray-500 whitespace-nowrap">وضعیت فعلی:</span>
                        <span :class="getStatusClass(selectedReview.status)" class="px-2 py-1 rounded-full whitespace-nowrap ml-2">{{ getStatusLabel(selectedReview.status) }}</span>
                      </div>
                      <!-- وضعیت خرید -->
                      <div class="flex items-center pl-16 border-l border-gray-200">
                        <span class="text-gray-500 whitespace-nowrap">وضعیت خرید:</span>
                        <span class="font-bold ml-2" :class="selectedReview.hasPurchased ? 'text-green-600' : 'text-red-600'">
                          {{ selectedReview.hasPurchased ? '✅ خریداری شده' : '❌ خریداری نشده' }}
                        </span>
                      </div>
                      <!-- آدرس IP -->
                      <div class="flex items-center pl-16 border-l border-gray-200">
                        <span class="text-gray-500 whitespace-nowrap">آدرس IP:</span>
                        <span class="font-medium ml-2 whitespace-nowrap">{{ selectedReview.ipAddress || 'نامشخص' }}</span>
                      </div>
                      <!-- دکمه رهگیری آی پی -->
                      <div class="flex items-center ml-16">
                        <button
                          @click="selectedReview.ipAddress && trackIP(selectedReview.ipAddress)"
                          :disabled="!selectedReview.ipAddress"
                          class="bg-blue-500 hover:bg-blue-600 text-white text-xs px-3 py-1 rounded-lg transition-colors duration-200 flex items-center gap-1 whitespace-nowrap disabled:opacity-50 disabled:cursor-not-allowed"
                        >
                          رهگیری آی پی
                        </button>
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              <div class="bg-blue-50 rounded-xl border border-blue-200 shadow-sm">
                <button @click="isStatusSectionOpen = !isStatusSectionOpen" class="w-full p-6 flex items-center justify-between cursor-pointer">
                  <h4 class="text-base font-semibold text-gray-900">تغییر وضعیت</h4>
                  <svg class="w-5 h-5 text-gray-600 transition-transform" :class="{'rotate-180': isStatusSectionOpen}" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path></svg>
                </button>
                <div v-if="isStatusSectionOpen" class="p-6 border-t border-blue-200">
                  <div class="flex flex-wrap gap-3">
                    <button @click="approveReview(selectedReview.id)" :disabled="selectedReview.status === 'approved'" :class="selectedReview.status === 'approved' ? 'bg-green-600 text-white border-2 border-green-700 font-bold' : 'bg-green-100 text-green-800 hover:bg-green-200 border border-green-300'" class="flex-1 btn-status-no-shadow disabled:opacity-50 transition-all duration-150">
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg>
                      تایید شده
                    </button>
                    <button @click="rejectReview(selectedReview.id)" :disabled="selectedReview.status === 'rejected'" :class="selectedReview.status === 'rejected' ? 'bg-red-600 text-white border-2 border-red-700 font-bold' : 'bg-red-100 text-red-800 hover:bg-red-200 border border-red-300'" class="flex-1 btn-status-no-shadow disabled:opacity-50 transition-all duration-150">
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path></svg>
                      رد شده
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <div class="bg-gray-50 px-4 py-3 border-t">
            <div class="flex gap-2 justify-end">
              <button v-if="!isAnsweringInline" @click="startInlineAnswer" class="btn bg-green-600 text-white hover:bg-green-700">{{ selectedReview && selectedReview.adminReply ? 'ویرایش پاسخ' : 'پاسخ دادن' }}</button>
              <button @click="showReviewModal = false" class="btn bg-gray-200 text-gray-700 hover:bg-gray-300">بستن</button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Reply Modal -->
    <div v-if="showReplyModal" class="fixed inset-0 z-50 overflow-y-auto">
      <div class="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
        <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" @click="showReplyModal = false"></div>
        
        <span class="hidden sm:inline-block sm:align-middle sm:h-screen">&#8203;</span>
        
        <div class="inline-block align-bottom bg-white rounded-lg text-right overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-2xl sm:w-full" dir="rtl">
          <div class="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
            <!-- Modal Header -->
            <div class="flex items-center justify-between pb-4 border-b border-gray-200">
              <h3 class="text-lg font-medium text-gray-900">پاسخ به نظر</h3>
              <button @click="showReplyModal = false" class="text-gray-400 hover:text-gray-600">
                <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                </svg>
              </button>
            </div>

            <!-- Review Summary -->
            <div v-if="selectedReview" class="mt-4 bg-gray-50 rounded-lg p-6">
              <div class="flex items-center space-x-3 space-x-reverse">
                <div class="flex-1">
                  <div class="text-sm font-medium text-gray-900">{{ selectedReview.customer.name }}</div>
                  <div class="text-sm text-gray-500">{{ selectedReview.product.name }}</div>
                </div>
                <div class="flex items-center">
                  <span v-for="n in 5" :key="n" :class="n <= selectedReview.rating ? 'text-yellow-400' : 'text-gray-300'" class="text-sm">
                    ⭐
                  </span>
                </div>
              </div>
              <div class="mt-3 text-sm text-gray-700 border-r-4 border-blue-200 pr-4">
                "{{ selectedReview.comment }}"
              </div>
            </div>

            <!-- Reply Form -->
            <div class="mt-6">
              <label for="reply-text" class="block text-sm font-medium text-gray-700 mb-2">
                پاسخ شما
              </label>
              <textarea
                id="reply-text"
                v-model="replyText"
                rows="6"
                class="block w-full border border-gray-300 rounded-md shadow-sm px-3 py-2 text-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                placeholder="پاسخ خود را به نظر مشتری بنویسید..."
                dir="rtl"
              ></textarea>
              <div class="flex justify-between items-center mt-2">
                <span class="text-xs text-gray-500">{{ replyText.length }}/500 کاراکتر</span>
                <div class="flex space-x-2 space-x-reverse text-xs">
                  <button 
                    @click="insertTemplate('thanks')"
                    class="text-blue-600 hover:text-blue-700"
                  >
                    + تشکر
                  </button>
                  <button 
                    @click="insertTemplate('apology')"
                    class="text-blue-600 hover:text-blue-700"
                  >
                    + عذرخواهی
                  </button>
                  <button 
                    @click="insertTemplate('followup')"
                    class="text-blue-600 hover:text-blue-700"
                  >
                    + پیگیری
                  </button>
                </div>
              </div>
            </div>
          </div>

          <!-- Reply Modal Actions -->
          <div class="bg-gray-50 px-4 py-3 sm:px-6 sm:flex sm:flex-row-reverse space-y-2 sm:space-y-0 sm:space-x-3 sm:space-x-reverse">
            <button 
              @click="submitReply"
              :disabled="!replyText.trim()"
              class="w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-blue-600 text-base font-medium text-white hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:bg-gray-400 disabled:cursor-not-allowed sm:ml-3 sm:w-auto sm:text-sm"
            >
              ارسال پاسخ
            </button>
            <button 
              @click="showReplyModal = false"
              class="w-full inline-flex justify-center rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 sm:ml-3 sm:w-auto sm:text-sm"
            >
              انصراف
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Settings Modal -->
    <ViewAllModal v-model="showSettingsModal" title="تنظیمات نظرات">
      <review-settings @loading="settingsLoading = $event" />
    </ViewAllModal>

    <!-- Image Preview Modal -->
    <ImagePreviewModal v-model="showImagePreview" :image="selectedImage" />
  
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
</script>

<script setup lang="ts">
import TemplateCard from '@/components/common/TemplateCard.vue'
import { computed, onMounted, provide, ref, watch } from 'vue'
import Pagination from '~/components/admin/common/Pagination.vue'
import ViewAllModal from '~/components/admin/modals/ViewAllModal.vue'
import ImagePreviewModal from '~/components/media/ImagePreviewModal.vue'
import { useAuth } from '~/composables/useAuth'
import { useConfirmDialog } from '~/composables/useConfirmDialog'
import { useNotifier } from '~/composables/useNotifier'
import ReviewSettings from './review-settings.vue'

// Import useAuth for permission checking
const { hasPermission } = useAuth()

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// Provide ConfirmDialog ref for this page
const confirmDialogRef = ref<any>(null)
provide('confirmDialogRef', confirmDialogRef)

// Interfaces
interface Customer {
  name: string
  email: string
  phone?: string
  joinDate: string
  reviewCount?: number
}

interface Product {
  id: string
  name: string
  sku?: string
  image: string
  price: number
  thumbnail?: string
}

interface Review {
  id: string
  customer: Customer
  product: Product
  rating: number
  comment: string
  status: 'pending' | 'approved' | 'rejected'
  createdAt: string
  images?: string[]
  ipAddress?: string
  hasPurchased: boolean
  adminReply?: string
}

// State
const reviews = ref<Review[]>([])
const selectedReviews = ref<string[]>([])
const showFilters = ref(false)
const showReviewModal = ref(false)
const showReplyModal = ref(false)
const selectedReview = ref<Review | null>(null)
const replyText = ref('')
const currentPage = ref(1)
const itemsPerPage = ref(10)
const showSettingsModal = ref(false)
const settingsLoading = ref(false)
const showImagePreview = ref(false)
const selectedImage = ref<{ url?: string; src?: string; name?: string } | null>(null)

// Modal-specific state for inline answering
const isAnsweringInline = ref(false)
const inlineAnswerText = ref('')
const isAdditionalInfoOpen = ref(true)
const isStatusSectionOpen = ref(true)

// Filters
const filters = ref({
  search: '',
  status: '',
  rating: '',
  dateRange: ''
})

// Stats
const reviewsStats = ref({
  total: 0,
  approved: 0,
  pending: 0,
  rejected: 0,
  averageRating: 0
})

// Computed
const allSelected = computed(() => {
  return paginatedReviews.value.length > 0 && selectedReviews.value.length === paginatedReviews.value.length
})

const filteredReviews = computed(() => {
  let filtered = [...reviews.value]

  // جستجو
  if (filters.value.search) {
    const search = filters.value.search.toLowerCase()
    filtered = filtered.filter(review => 
      review.customer.name.toLowerCase().includes(search) ||
      review.customer.email.toLowerCase().includes(search) ||
      review.product.name.toLowerCase().includes(search) ||
      review.comment.toLowerCase().includes(search)
    )
  }

  // فیلتر وضعیت
  if (filters.value.status) {
    filtered = filtered.filter(review => review.status === filters.value.status)
  }

  // فیلتر امتیاز
  if (filters.value.rating) {
    filtered = filtered.filter(review => review.rating === parseInt(filters.value.rating))
  }

  // فیلتر تاریخ
  if (filters.value.dateRange) {
    const now = new Date()
    const filterDate = new Date()
    
    switch (filters.value.dateRange) {
      case 'today':
        filterDate.setHours(0, 0, 0, 0)
        break
      case 'week':
        filterDate.setDate(now.getDate() - 7)
        break
      case 'month':
        filterDate.setMonth(now.getMonth() - 1)
        break
      case '3months':
        filterDate.setMonth(now.getMonth() - 3)
        break
      case '6months':
        filterDate.setMonth(now.getMonth() - 6)
        break
      case 'year':
        filterDate.setFullYear(now.getFullYear() - 1)
        break
    }
    
    filtered = filtered.filter(review => new Date(review.createdAt) >= filterDate)
  }

  return filtered
})

const totalPages = computed(() => {
  return Math.ceil(filteredReviews.value.length / itemsPerPage.value)
})

const paginatedReviews = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage.value
  const end = start + itemsPerPage.value
  return filteredReviews.value.slice(start, end)
})

const visiblePages = computed(() => {
  const total = totalPages.value
  const current = currentPage.value
  const delta = 2
  const range = []
  const rangeWithDots = []

  for (let i = Math.max(2, current - delta); i <= Math.min(total - 1, current + delta); i++) {
    range.push(i)
  }

  if (current - delta > 2) {
    rangeWithDots.push(1, '...')
  } else {
    rangeWithDots.push(1)
  }

  rangeWithDots.push(...range)

  if (current + delta < total - 1) {
    rangeWithDots.push('...', total)
  } else if (total > 1) {
    rangeWithDots.push(total)
  }

  return rangeWithDots
})

// Methods
// Local alias to avoid deep type instantiation on $fetch
const $f: any = (globalThis as any).$fetch || (window as any).$fetch
const loadReviews = async () => {
  try {
    const params = new URLSearchParams()
    if (filters.value.search) params.append('search', filters.value.search)
    if (filters.value.status) params.append('status', filters.value.status)
    if (filters.value.rating) params.append('rating', String(filters.value.rating))
    params.append('page', String(currentPage.value))
    params.append('per_page', String(itemsPerPage.value))

    const response: any = await $f(`/api/admin/reviews?${params.toString()}`)
    // Map ID, createdAt, joinDate, and product.image for frontend compatibility
    const list = Array.isArray(response?.reviews) ? response.reviews : []
    reviews.value = list.map((r: any) => ({
      ...r,
      id: r.id || r.ID,
      createdAt: r.createdAt || r.created_at,
      updatedAt: r.updatedAt || r.updated_at,
      customer: {
        ...r.customer,
        joinDate: r.customer?.joinDate || r.customer?.join_date
      },
      product: {
        ...r.product,
        image: r.product?.image || r.product?.image_url || r.product?.main_image || '/default-product.svg',
        thumbnail: r.product?.thumbnail || r.product?.image || r.product?.image_url || r.product?.main_image || '/default-product.svg'
      }
    }))
    updateStats()
  } catch (error) {
    // Failed to load reviews
    reviews.value = []
  }
}

const updateStats = () => {
  const total = reviews.value.length
  const approved = reviews.value.filter(r => r.status === 'approved').length
  const pending = reviews.value.filter(r => r.status === 'pending').length
  const rejected = reviews.value.filter(r => r.status === 'rejected').length
  const averageRating = reviews.value.length > 0 
    ? reviews.value.reduce((sum, r) => sum + r.rating, 0) / reviews.value.length 
    : 0

  reviewsStats.value = {
    total,
    approved,
    pending,
    rejected,
    averageRating
  }
}

const formatDate = (dateStr: string) => {
  if (!dateStr) return '-'
  const d = new Date(dateStr)
  if (isNaN(d.getTime())) return '-'
  return d.toLocaleDateString('fa-IR', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}

const formatTime = (dateStr: string) => {
  if (!dateStr) return '-'
  const d = new Date(dateStr)
  if (isNaN(d.getTime())) return '-'
  return d.toLocaleTimeString('fa-IR', { 
    hour: '2-digit', 
    minute: '2-digit',
    hour12: false
  })
}

const truncateText = (text: string, maxLength: number) => {
  if (text.length <= maxLength) return text
  return text.substring(0, maxLength) + '...'
}

const getStatusClass = (status: string) => {
  const classes = {
    pending: 'bg-yellow-100 text-yellow-800',
    approved: 'bg-green-100 text-green-800',
    rejected: 'bg-red-100 text-red-800'
  }
  return classes[status as keyof typeof classes] || 'bg-gray-100 text-gray-800'
}

const getStatusLabel = (status: string) => {
  const labels = {
    pending: 'در انتظار بررسی',
    approved: 'تایید شده',
    rejected: 'رد شده'
  }
  return labels[status as keyof typeof labels] || status
}

const toggleSelectAll = () => {
  if (allSelected.value) {
    selectedReviews.value = []
  } else {
    selectedReviews.value = paginatedReviews.value.map(r => r.id)
  }
}

const applyFilters = () => {
  currentPage.value = 1
  // فیلترها خودکار اعمال می‌شوند
  loadReviews()
}

const clearFilters = () => {
  filters.value = {
    search: '',
    status: '',
    rating: '',
    dateRange: ''
  }
  currentPage.value = 1
  loadReviews()
}

const viewReview = (review: Review) => {
  selectedReview.value = review
  showReviewModal.value = true
}

const editReview = (review: Review) => {
  // پیاده‌سازی ویرایش نظر - در آینده با modal یا صفحه جدید
  useNotifier().info(`ویرایش نظر ${review.id} - این قابلیت در نسخه بعدی اضافه خواهد شد`)
}

const approveReview = async (reviewId: string) => {
  try {
    await $f(`/api/admin/reviews/${reviewId}/approve`, { method: 'POST' })
    
    // به‌روزرسانی وضعیت در local state
    const review = reviews.value.find(r => r.id === reviewId)
    if (review) {
      review.status = 'approved'
    }
    
    // به‌روزرسانی selectedReview اگر همان نظر باشد
    if (selectedReview.value && selectedReview.value.id === reviewId) {
      selectedReview.value.status = 'approved'
    }
    
    updateStats()
    // نظر تایید شد
  } catch (error) {
    // خطا در تایید نظر
    useNotifier().error('خطا در تایید نظر. لطفاً دوباره تلاش کنید.')
  }
}

const rejectReview = async (reviewId: string) => {
  try {
    await $f(`/api/admin/reviews/${reviewId}/reject`, { method: 'POST' })
    
    // به‌روزرسانی وضعیت در local state
    const review = reviews.value.find(r => r.id === reviewId)
    if (review) {
      review.status = 'rejected'
    }
    
    // به‌روزرسانی selectedReview اگر همان نظر باشد
    if (selectedReview.value && selectedReview.value.id === reviewId) {
      selectedReview.value.status = 'rejected'
    }
    
    updateStats()
    // نظر رد شد
  } catch (error) {
    // خطا در رد نظر
    useNotifier().error('خطا در رد نظر. لطفاً دوباره تلاش کنید.')
  }
}

const deleteReview = async (reviewId: string) => {
  // deleteReview called
  if (!reviewId || reviewId === 'undefined') {
    useNotifier().error('خطا: شناسه نظر نامعتبر است!')
    return
  }
  {
    const { confirm } = useConfirmDialog()
    const ok = await confirm({ title:'تأیید حذف', message:'آیا از حذف این نظر اطمینان دارید؟ این عمل قابل بازگشت نیست.', confirmText:'حذف', cancelText:'انصراف', type:'danger' })
    if (!ok) return
  }
  try {
    await $f(`/api/admin/reviews/${reviewId}`, { method: 'DELETE' })
    // حذف از local state
    reviews.value = reviews.value.filter(r => r.id !== reviewId)
    selectedReviews.value = selectedReviews.value.filter(id => id !== reviewId)
    // نظر حذف شد
  } catch (error) {
    // خطا در حذف نظر
    useNotifier().error('خطا در حذف نظر. لطفاً دوباره تلاش کنید.')
  }
}

const replyToReview = (review: Review) => {
  // پیاده‌سازی پاسخ به نظر - در آینده با modal مخصوص
  useNotifier().info(`پاسخ به نظر ${review.customer.name} - این قابلیت در نسخه بعدی اضافه خواهد شد`)
}

const bulkApprove = () => {
  if (selectedReviews.value.length === 0) return
  
  selectedReviews.value.forEach(reviewId => {
    const review = reviews.value.find(r => r.id === reviewId)
    if (review) review.status = 'approved'
  })
  
  // نظرات تایید شدند
  selectedReviews.value = []
}

const bulkReject = async () => {
  if (selectedReviews.value.length === 0) return
  
  const { confirm } = useConfirmDialog()
  const ok = await confirm({ title:'رد گروهی', message:`آیا از رد ${selectedReviews.value.length} نظر انتخاب شده اطمینان دارید؟`, confirmText:'تأیید', cancelText:'انصراف', type:'warning' })
  if (!ok) return

  selectedReviews.value.forEach(reviewId => {
    const review = reviews.value.find(r => r.id === reviewId)
    if (review) review.status = 'rejected'
  })

  // نظرات رد شدند
  selectedReviews.value = []
}

const bulkDelete = async () => {
  if (selectedReviews.value.length === 0) return
  
  const { confirm } = useConfirmDialog()
  const ok = await confirm({ title:'حذف گروهی', message:`آیا از حذف ${selectedReviews.value.length} نظر انتخاب شده اطمینان دارید؟ این عمل قابل بازگشت نیست.`, confirmText:'حذف', cancelText:'انصراف', type:'danger' })
  if (!ok) return

  reviews.value = reviews.value.filter(r => !selectedReviews.value.includes(r.id))
  selectedReviews.value = []
  // نظرات انتخاب شده حذف شدند
}

const exportReviews = () => {
  // پیاده‌سازی خروجی Excel
  const csvContent = "data:text/csv;charset=utf-8," 
    + "نام مشتری,ایمیل,محصول,امتیاز,نظر,وضعیت,تاریخ\n"
    + filteredReviews.value.map(review => 
        [
          review.customer.name,
          review.customer.email,
          review.product.name,
          review.rating,
          review.comment.replace(/,/g, ';'),
          getStatusLabel(review.status),
          formatDate(review.createdAt)
        ].join(",")
      ).join("\n")

  const encodedUri = encodeURI(csvContent)
  const link = document.createElement("a")
  link.setAttribute("href", encodedUri)
  link.setAttribute("download", `reviews_${new Date().toISOString().split('T')[0]}.csv`)
  document.body.appendChild(link)
  link.click()
  document.body.removeChild(link)
}

const viewImage = (imageUrl: string) => {
  window.open(imageUrl, '_blank', 'width=800,height=600')
}

const getFullSizeImage = (product: Product) => {
  // اگر تصویر اصلی موجود است همان را برمی‌گردانیم
  if (product.image && product.image !== '/default-product.svg') return product.image

  // تلاش برای تبدیل thumbnail به نسخه بزرگ‌تر با حذف پسوندهای _thumb/_thumbnail
  const thumb = product.thumbnail || ''
  const full = thumb.replace(/(_thumb(nail)?)(?=\.)/i, '')
  return full || '/default-product.svg'
}

const previewImage = (imageUrl: string, imageName?: string) => {
  if (imageUrl && imageUrl !== '/default-product.svg') {
    selectedImage.value = {
      url: imageUrl,
      src: imageUrl,
      name: imageName || 'عکس محصول'
    }
    showImagePreview.value = true
  }
}

const goToPage = (page: number | string) => {
  if (typeof page === 'number' && page >= 1 && page <= totalPages.value) {
    currentPage.value = page
    loadReviews()
  }
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

const insertTemplate = (template: string) => {
  const templates = {
    'thanks': 'از شما بابت نظر ارزشمندتان سپاسگزاریم. نظرات شما به بهبود کیفیت خدمات ما کمک می‌کند.',
    'apology': 'از ناراحتی شما عذرخواهی می‌کنیم. مشکل شما برای ما اهمیت دارد و در تلاش برای بهبود هستیم.',
    'followup': 'جهت بررسی بیشتر موضوع و حل مشکل، لطفاً با پشتیبانی ما تماس بگیرید. شماره تماس: ۰۲۱-۸۸۸۸۸۸۸۸'
  }
  
  const templateText = templates[template as keyof typeof templates]
  if (templateText) {
    if (replyText.value) {
      replyText.value += '\n\n' + templateText
    } else {
      replyText.value = templateText
    }
  }
}

const submitReply = async () => {
  if (!selectedReview.value || !replyText.value.trim()) return
  
  try {
    await $f(`/api/admin/reviews/${selectedReview.value.id}/reply`, {
      method: 'POST',
      body: {
        reply: replyText.value.trim()
      }
    })
    
    // اضافه کردن پاسخ به نظر در local state
    selectedReview.value.adminReply = replyText.value.trim()
    
    // پاسخ به نظر ارسال شد
    
    // ریست کردن فرم و بستن modal
    replyText.value = ''
    showReplyModal.value = false
    
    // نمایش پیام موفقیت
  useNotifier().success('پاسخ شما با موفقیت ارسال شد و برای مشتری نمایش داده خواهد شد.')
  } catch (error) {
    // خطا در ارسال پاسخ
  useNotifier().error('خطا در ارسال پاسخ. لطفاً دوباره تلاش کنید.')
  }
}

// Adding setStatusFilter function to fix linter errors
function setStatusFilter(status: string) {
  filters.value.status = status;
}

// Inline answer functions
const startInlineAnswer = () => {
  isAnsweringInline.value = true
  inlineAnswerText.value = selectedReview.value?.adminReply || ''
}

const submitInlineAnswer = async () => {
  if (!selectedReview.value || !inlineAnswerText.value.trim()) return
  
  try {
    await $f(`/api/admin/reviews/${selectedReview.value.id}/reply`, {
      method: 'POST',
      body: {
        reply: inlineAnswerText.value.trim()
      }
    })
    
    // اضافه کردن پاسخ به نظر در local state
    selectedReview.value.adminReply = inlineAnswerText.value.trim()
    
    // پاسخ به نظر ارسال شد
    
    // ریست کردن فرم
    isAnsweringInline.value = false
    inlineAnswerText.value = ''
    
    // نمایش پیام موفقیت
  useNotifier().success('پاسخ شما با موفقیت ارسال شد و برای مشتری نمایش داده خواهد شد.')
  } catch (error) {
    // خطا در ارسال پاسخ
  useNotifier().error('خطا در ارسال پاسخ. لطفاً دوباره تلاش کنید.')
  }
}

// IP tracking function
const trackIP = (ipAddress: string) => {
  if (ipAddress) {
    window.open(`https://www.whatismyipaddress.com/ip/${ipAddress}`, '_blank')
  }
}

// Watchers
watch(filters, () => {
  currentPage.value = 1
}, { deep: true })

// Lifecycle
onMounted(() => {
  loadReviews()
})

function getRatingStyle(rating) {
  switch (rating) {
    case 1:
      return {
        background: 'linear-gradient(90deg, #ff4e50 0%, #f9d423 100%)',
        color: '#fff',
        borderRadius: '12px',
        padding: '6px 18px',
        fontWeight: 'bold',
        fontSize: '1.1rem',
        boxShadow: '0 2px 8px 0 rgba(255,78,80,0.10)'
      };
    case 2:
      return {
        background: 'linear-gradient(90deg, #ffb199 0%, #ff0844 100%)',
        color: '#fff',
        borderRadius: '12px',
        padding: '6px 18px',
        fontWeight: 'bold',
        fontSize: '1.1rem',
        boxShadow: '0 2px 8px 0 rgba(255,176,153,0.10)'
      };
    case 3:
      return {
        background: 'linear-gradient(90deg, #fbc2eb 0%, #f9d423 100%)',
        color: '#444',
        borderRadius: '12px',
        padding: '6px 18px',
        fontWeight: 'bold',
        fontSize: '1.1rem',
        boxShadow: '0 2px 8px 0 rgba(251,194,235,0.10)'
      };
    case 4:
      return {
        background: 'linear-gradient(90deg, #a8ff78 0%, #78ffd6 100%)',
        color: '#222',
        borderRadius: '12px',
        padding: '6px 18px',
        fontWeight: 'bold',
        fontSize: '1.1rem',
        boxShadow: '0 2px 8px 0 rgba(120,255,214,0.10)'
      };
    case 5:
      return {
        background: 'linear-gradient(90deg, #43e97b 0%, #38f9d7 100%)',
        color: '#fff',
        borderRadius: '12px',
        padding: '6px 18px',
        fontWeight: 'bold',
        fontSize: '1.1rem',
        boxShadow: '0 2px 8px 0 rgba(67,233,123,0.10)'
      };
    default:
      return {
        background: '#eee',
        color: '#444',
        borderRadius: '12px',
        padding: '6px 18px',
        fontWeight: 'bold',
        fontSize: '1.1rem',
        boxShadow: '0 2px 8px 0 rgba(200,200,200,0.10)'
      };
  }
}
</script>

<style scoped>
.line-clamp-2 {
  line-clamp: 2; /* compatibility */
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

/* Custom scrollbar for modal */
.overflow-y-auto::-webkit-scrollbar {
  width: 6px;
}

.overflow-y-auto::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 3px;
}

.overflow-y-auto::-webkit-scrollbar-thumb {
  background: #c1c1c1;
  border-radius: 3px;
}

.overflow-y-auto::-webkit-scrollbar-thumb:hover {
  background: #a8a8a8;
}

/* Animation for modal */
.modal-enter-active, .modal-leave-active {
  transition: opacity 0.3s ease;
}

.modal-enter-from, .modal-leave-to {
  opacity: 0;
}
</style> 
<!-- Global Confirm Dialog instance for this page -->
<ConfirmDialog ref="confirmDialogRef" />

