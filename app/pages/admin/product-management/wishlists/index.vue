<template>
  <ClientOnly>
    <!-- Page Header -->
    <div class="mb-4">
      <div class="w-full px-4 sm:px-6 lg:px-8 bg-white">
        <div class="flex justify-between items-center py-3">
          <div>
            <h1 class="text-lg font-bold text-gray-900">مدیریت لیست علاقمندی</h1>
            <p class="text-xs text-gray-600 mt-1">مشاهده و مدیریت محصولات اضافه شده به لیست علاقمندی کاربران</p>
          </div>
          <div class="flex space-x-2 space-x-reverse">
            <button
                class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-cyan-400 to-cyan-600 hover:from-cyan-500 hover:to-cyan-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-cyan-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105"
                @click="showFilters = !showFilters"
            >
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.707A1 1 0 013 7V4z"></path>
              </svg>
              فیلترها
            </button>
            <ExportExcelButton
              :data="exportData"
              filename="wishlists.csv"
              :button-class="'inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-blue-400 to-blue-600 hover:from-blue-500 hover:to-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105'"
            >
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01-.293.707V19a2 2 0 01-2 2z"></path>
              </svg>
              خروجی اکسل
            </ExportExcelButton>
            <button 
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-green-400 to-green-600 hover:from-green-500 hover:to-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105"
              @click="sendWishlistEmail"
            >
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 4.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"></path>
              </svg>
              ارسال ایمیل تشویقی
            </button>
            <button 
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-indigo-400 to-indigo-600 hover:from-indigo-500 hover:to-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105"
              @click="openBulkMessageModal"
            >
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 8h10M7 12h8m-5 8l-4-4h8l-4 4z"/>
              </svg>
              ارسال پیام گروهی
            </button>
            <div v-if="showBulkMessageModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
              <div class="bg-white rounded-lg shadow-xl p-6 w-full max-w-lg mx-4">
                <h3 class="text-lg font-semibold text-gray-900 mb-4">ارسال پیام گروهی به علاقه‌مندان</h3>
                <div class="space-y-4">
                  <div class="text-sm text-gray-700">گیرندگان: {{ bulkRecipientsCount }} کاربر</div>
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-2">نوع پیام</label>
                    <select v-model="smsType" class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500">
                      <option value="discount">اطلاع‌رسانی تخفیف</option>
                      <option value="availability">اطلاع‌رسانی موجودی</option>
                      <option value="price_drop">کاهش قیمت</option>
                      <option value="custom">پیام سفارشی</option>
                    </select>
                  </div>
                  <div v-if="smsType === 'discount'">
                    <label class="block text-sm font-medium text-gray-700 mb-2">درصد تخفیف</label>
                    <input v-model="discountPercent" type="number" min="1" max="99" class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500" />
                  </div>
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-2">متن پیام</label>
                    <textarea v-model="smsMessage" rows="4" class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500" :placeholder="getBulkSmsTemplate()"></textarea>
                    <div class="text-xs text-gray-500 mt-1">تعداد کاراکتر: {{ smsMessage.length }}/160</div>
                  </div>
                </div>
                <div class="flex justify-end space-x-2 space-x-reverse mt-6">
                  <button class="px-4 py-2 bg-gray-300 text-gray-700 text-sm font-medium rounded-lg hover:bg-gray-400 transition-colors" @click="showBulkMessageModal = false">انصراف</button>
                  <button class="px-4 py-2 bg-indigo-600 text-white text-sm font-medium rounded-lg hover:bg-indigo-700 transition-colors" @click="sendBulkMessage">ارسال</button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="w-full px-4 py-4">

      <div class="mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div class="relative">
          <div class="relative grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-6 mb-6">
            <div
                class="cursor-pointer transition-all duration-200 hover:scale-105"
                @click="viewMode = 'product'"
            >
              <TemplateCard 
                title="کل علاقمندی‌ها" 
                :value="totalWishlists" 
                variant="blue"
              >
                <template #icon>
                  <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"></path>
                  </svg>
                </template>
              </TemplateCard>
            </div>

            <div
                class="cursor-pointer transition-all duration-200 hover:scale-105"
                @click="viewMode = 'user'"
            >
              <TemplateCard 
                title="محصولات منحصر" 
                :value="uniqueProducts" 
                variant="green"
              >
                <template #icon>
                  <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
                  </svg>
                </template>
              </TemplateCard>
            </div>

            <TemplateCard 
              title="کاربران فعال" 
              :value="activeUsers" 
              variant="purple"
            >
              <template #icon>
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197m13.5-9a2.5 2.5 0 11-5 0 2.5 2.5 0 015 0z"></path>
                </svg>
              </template>
            </TemplateCard>

            <TemplateCard 
              title="پردرخواست‌ترین" 
              :value="mostRequestedProduct?.users.length || 0" 
              variant="red"
            >
              <template #icon>
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
                </svg>
              </template>
            </TemplateCard>

            <TemplateCard 
              title="تبدیل به خرید" 
              :value="`${conversionRate}%`" 
              variant="amber"
            >
              <template #icon>
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4m0 0L7 13m0 0l-2.5 5M7 13l2.5 5m6-5v6a2 2 0 01-2 2H9a2 2 0 01-2-2v-6m6 0V9a2 2 0 00-2-2H9a2 2 0 00-2 2v4.01"></path>
                </svg>
              </template>
            </TemplateCard>
          </div>
        </div>
      </div>
    </div>
    <!-- این سه تگ بسته اضافه شد -->

    <div v-if="showFilters" class="shadow-lg rounded-2xl mb-8 border border-blue-100">
        <div class="px-6 py-4 border-b border-blue-100 flex items-center gap-2">
          <svg class="w-6 h-6 text-blue-400" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.707A1 1 0 013 7V4z"></path></svg>
          <h3 class="text-lg font-bold text-blue-900">فیلترها و جستجو</h3>
        </div>
        <div class="p-8">
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">جستجو در محصولات</label>
              <div class="relative">
                <input
                    v-model="searchQuery"
                    type="text"
                    class="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md leading-5 bg-white placeholder-gray-500 focus:outline-none focus:placeholder-gray-400 focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm"
                    placeholder="نام محصول، برند یا دسته‌بندی..."
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
              <label class="block text-sm font-medium text-gray-700 mb-2">دسته‌بندی</label>
              <select 
                v-model="selectedCategory"
                class="block w-full px-3 py-2 border border-gray-300 rounded-md leading-5 bg-white focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm"
              >
                <option value="">همه دسته‌ها</option>
                <option value="الکترونیکی">الکترونیکی</option>
                <option value="پوشاک">پوشاک</option>
                <option value="خانه و آشپزخانه">خانه و آشپزخانه</option>
                <option value="کتاب">کتاب</option>
                <option value="ورزشی">ورزشی</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">بازه قیمت</label>
              <select 
                v-model="priceRange"
                class="block w-full px-3 py-2 border border-gray-300 rounded-md leading-5 bg-white focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm"
              >
                <option value="">همه قیمت‌ها</option>
                <option value="0-1000000">زیر 1 میلیون تومان</option>
                <option value="1000000-5000000">1 تا 5 میلیون تومان</option>
                <option value="5000000-10000000">5 تا 10 میلیون تومان</option>
                <option value="10000000+">بالای 10 میلیون</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ اضافه شدن</label>
              <select 
                v-model="dateRange"
                class="block w-full px-3 py-2 border border-gray-300 rounded-md leading-5 bg-white focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm"
              >
                <option value="">همه تاریخ‌ها</option>
                <option value="today">امروز</option>
                <option value="week">هفته گذشته</option>
                <option value="month">ماه گذشته</option>
                <option value="3months">3 ماه گذشته</option>
              </select>
            </div>
          </div>

          <div class="flex justify-between items-center mt-6 pt-4 border-t border-blue-100">
            <button 
              class="text-sm text-blue-600 hover:text-blue-800 font-medium"
              @click="clearFilters"
            >
              پاک کردن فیلترها
            </button>
            <button 
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-blue-400 to-blue-600 hover:from-blue-500 hover:to-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105"
              @click="performSearch"
            >
              <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
              </svg>
              اعمال فیلترها
            </button>
          </div>
        </div>
      </div>

      <!-- Wishlist Table -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
        <div class="bg-gradient-to-r from-slate-50 to-gray-50 px-4 py-3 border-b border-gray-200">
          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-6 h-6 bg-slate-600 rounded-md flex items-center justify-center ml-2">
                <span class="text-white text-sm">💙</span>
              </div>
              <h3 class="text-sm font-semibold text-gray-900">لیست علاقمندی‌ها</h3>
            </div>
          </div>
        </div>

        <!-- Table -->
        <div class="overflow-hidden">
          <div class="overflow-x-auto">
            
            <!-- Product-centric View -->
            <table v-if="viewMode === 'product'" class="min-w-full divide-y divide-gray-200">
              <thead class="bg-gray-50">
                <tr>
                  <th scope="col" class="px-3 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">
                    <input 
                      type="checkbox" 
                      :checked="isAllProductsSelected"
                      class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                      @change="toggleSelectAllProducts"
                    />
                  </th>
                  <th scope="col" class="px-3 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">تصویر</th>
                  <th 
                    scope="col" 
                    class="px-3 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider cursor-pointer hover:bg-gray-100 transition-colors"
                    @click="sortByColumn('name')"
                  >
                    <div class="flex items-center justify-center">
                      نام محصول
                      <span v-if="currentSortColumn === 'name'" class="mr-1">
                        {{ sortDirection === 'asc' ? '↑' : '↓' }}
                      </span>
                    </div>
                  </th>
                  <th 
                    scope="col" 
                    class="px-3 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider cursor-pointer hover:bg-gray-100 transition-colors"
                    @click="sortByColumn('price')"
                  >
                    <div class="flex items-center justify-center">
                      قیمت
                      <span v-if="currentSortColumn === 'price'" class="mr-1">
                        {{ sortDirection === 'asc' ? '↑' : '↓' }}
                      </span>
                    </div>
                  </th>
                  <th scope="col" class="px-3 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">دسته‌بندی</th>
                  <th 
                    scope="col" 
                    class="px-3 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider cursor-pointer hover:bg-gray-100 transition-colors"
                    @click="sortByColumn('wishlist_count')"
                  >
                    <div class="flex items-center justify-center">
                      تعداد علاقمندان
                      <span v-if="currentSortColumn === 'wishlist_count'" class="mr-1">
                        {{ sortDirection === 'asc' ? '↑' : '↓' }}
                      </span>
                    </div>
                  </th>
                  <th 
                    scope="col" 
                    class="px-3 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider cursor-pointer hover:bg-gray-100 transition-colors"
                    @click="sortByColumn('latest_date')"
                  >
                    <div class="flex items-center justify-center">
                      آخرین علاقمندی
                      <span v-if="currentSortColumn === 'latest_date'" class="mr-1">
                        {{ sortDirection === 'asc' ? '↑' : '↓' }}
                      </span>
                    </div>
                  </th>
                  <th 
                    scope="col" 
                    class="px-3 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider cursor-pointer hover:bg-gray-100 transition-colors"
                    @click="sortByColumn('stock_status')"
                  >
                    <div class="flex items-center justify-center">
                      وضعیت موجودی
                      <span v-if="currentSortColumn === 'stock_status'" class="mr-1">
                        {{ sortDirection === 'asc' ? '↑' : '↓' }}
                      </span>
                    </div>
                  </th>
                  <th scope="col" class="px-3 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">عملیات</th>
                </tr>
                                </thead>
                  <tbody class="bg-white divide-y divide-gray-200">
                    <template v-for="productGroup in paginatedProductGroups" :key="productGroup.product.id">
                  <!-- Main Product Row -->
                  <tr class="hover:bg-gray-50 transition-colors" :class="{ 'bg-blue-50': expandedProducts.includes(productGroup.product.id) }">
                    <!-- Checkbox -->
                    <td class="px-3 py-3 whitespace-nowrap text-center">
                      <input 
                        v-model="selectedProductIds" 
                        type="checkbox"
                        :value="productGroup.product.id"
                        class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                      />
                    </td>

                    <!-- Product Image -->
                    <td class="px-3 py-3 whitespace-nowrap text-center">
                      <img 
                        :src="productGroup.product.image" 
                        :alt="productGroup.product.name"
                        class="h-12 w-12 rounded-md object-cover mx-auto border border-gray-200"
                      />
                    </td>

                    <!-- Product Name -->
                    <td class="px-3 py-3 whitespace-nowrap text-right">
                      <div class="flex items-center">
                        <button 
                          class="mr-2 p-1 hover:bg-gray-200 rounded transition-colors"
                          @click="toggleProductExpansion(productGroup.product.id)"
                        >
                          <span v-if="expandedProducts.includes(productGroup.product.id)" class="text-blue-600">▼</span>
                          <span v-else class="text-gray-400">◄</span>
                        </button>
                        <div>
                          <div class="text-xs font-medium text-gray-900">{{ productGroup.product.name }}</div>
                          <div class="text-xs text-gray-500">کد: {{ productGroup.product.sku }}</div>
                        </div>
                      </div>
                    </td>

                    <!-- Price -->
                    <td class="px-3 py-3 whitespace-nowrap text-center">
                      <div class="text-xs font-medium text-gray-900">{{ formatPrice(productGroup.product.price) }}</div>
                      <div v-if="productGroup.product.discount" class="text-xs text-red-600">
                        {{ productGroup.product.discount }}% تخفیف
                      </div>
                    </td>

                    <!-- Category -->
                    <td class="px-3 py-3 whitespace-nowrap text-center">
                      <span class="inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium bg-gray-100 text-gray-800">
                        {{ productGroup.product.category }}
                      </span>
                    </td>

                    <!-- Users Count -->
                    <td class="px-3 py-3 whitespace-nowrap text-center">
                      <div class="flex flex-col items-center">
                      <button 
                        :class="[
                            'inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium transition-colors cursor-pointer',
                            productGroup.users.length >= 5 ? 'bg-red-100 text-red-800 hover:bg-red-200' :
                            productGroup.users.length >= 3 ? 'bg-orange-100 text-orange-800 hover:bg-orange-200' :
                            'bg-blue-100 text-blue-800 hover:bg-blue-200'
                          ]"
                          :title="'مدیریت کاربران علاقمند - ' + getRequestLevelText(productGroup.users.length)"
                          @click="manageProductUsers(productGroup.product.id)"
                      >
                          <span class="ml-1">
                            {{ productGroup.users.length >= 5 ? '🔥' : productGroup.users.length >= 3 ? '📈' : '💙' }}
                          </span>
                        {{ productGroup.users.length }} کاربر
                      </button>
                        <div class="text-xs text-gray-500 mt-1">
                          {{ getRequestLevelText(productGroup.users.length) }}
                        </div>
                      </div>
                    </td>

                    <!-- Latest Addition -->
                    <td class="px-3 py-3 whitespace-nowrap text-center">
                      <div class="text-xs text-gray-900">{{ formatDate(productGroup.latestDate) }}</div>
                      <div class="text-xs text-gray-500">{{ getRelativeTime(productGroup.latestDate) }}</div>
                    </td>

                    <!-- Stock Status -->
                    <td class="px-3 py-3 whitespace-nowrap text-center">
                      <span 
                        :class="[
                          'inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium',
                          productGroup.product.in_stock 
                            ? 'bg-green-100 text-green-800' 
                            : 'bg-red-100 text-red-800'
                        ]"
                      >
                        {{ productGroup.product.in_stock ? 'موجود' : 'ناموجود' }}
                      </span>
                    </td>

                    <!-- Actions -->
                    <td class="px-3 py-3 whitespace-nowrap text-center">
                      <div class="flex items-center justify-center space-x-1 space-x-reverse">
                        <button 
                          class="inline-flex items-center p-1.5 border border-transparent text-xs font-medium rounded text-blue-700 bg-blue-100 hover:bg-blue-200 transition-colors"
                          title="مشاهده محصول"
                          @click="viewProduct(productGroup.product.id)"
                        >
                          👁️
                        </button>
                        <button 
                          class="inline-flex items-center p-1.5 border border-transparent text-xs font-medium rounded text-green-700 bg-green-100 hover:bg-green-200 transition-colors"
                          title="اطلاع‌رسانی همه"
                          @click="notifyAllUsers(productGroup.product.id)"
                        >
                          📧
                        </button>
                        <button 
                          class="inline-flex items-center p-1.5 border border-transparent text-xs font-medium rounded text-red-700 bg-red-100 hover:bg-red-200 transition-colors"
                          title="حذف از همه لیست‌ها"
                          @click="removeProductFromAllWishlists(productGroup.product.id)"
                        >
                          🗑️
                        </button>
                      </div>
                    </td>
                  </tr>

                  <!-- Expanded Users List -->
                  <tr v-if="expandedProducts.includes(productGroup.product.id)" class="bg-blue-25">
                    <td colspan="9" class="px-6 py-4">
                      <div class="bg-gray-50 rounded-lg p-6">
                        <h4 class="text-sm font-medium text-gray-900 mb-3">کاربران علاقمند ({{ productGroup.users.length }} نفر):</h4>
                        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-3">
                          <div
v-for="userWishlist in productGroup.users" :key="userWishlist.user.id" 
                               class="bg-white rounded-md p-3 border border-gray-200 hover:border-blue-300 transition-colors">
                            <div class="flex items-center justify-between">
                              <div>
                                <div class="text-xs font-medium text-gray-900">{{ userWishlist.user.name }}</div>
                                <div class="text-xs text-gray-500">{{ userWishlist.user.email }}</div>
                                <div class="text-xs text-gray-400 mt-1">{{ getRelativeTime(userWishlist.created_at) }}</div>
                              </div>
                              <div class="flex space-x-1 space-x-reverse">
                                <button 
                                  class="p-1 text-green-600 hover:bg-green-100 rounded transition-colors"
                                  title="ارسال اعلان"
                                  @click="sendNotification(userWishlist.user.id, productGroup.product.id)"
                                >
                                  📧
                                </button>
                                <button 
                                  class="p-1 text-red-600 hover:bg-red-100 rounded transition-colors"
                                  title="حذف از لیست"
                                  @click="removeFromWishlist(userWishlist.id)"
                                >
                                  🗑️
                                </button>
                              </div>
                            </div>
                          </div>
                        </div>
                      </div>
                    </td>
                  </tr>
                </template>
              </tbody>
            </table>

            <!-- User-centric View (Original) -->
            <table v-else class="min-w-full divide-y divide-gray-200">
              <thead class="bg-gray-50">
                <tr>
                  <th scope="col" class="px-3 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">
                    <input 
                      type="checkbox" 
                      :checked="isAllSelected"
                      class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                      @change="toggleSelectAll"
                    />
                  </th>
                  <th scope="col" class="px-3 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">تصویر</th>
                  <th scope="col" class="px-3 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">نام محصول</th>
                  <th scope="col" class="px-3 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">قیمت</th>
                  <th scope="col" class="px-3 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">دسته‌بندی</th>
                  <th scope="col" class="px-3 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">کاربر</th>
                  <th scope="col" class="px-3 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">تاریخ اضافه</th>
                  <th scope="col" class="px-3 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">تعداد علاقمندی</th>
                  <th scope="col" class="px-3 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">وضعیت موجودی</th>
                  <th scope="col" class="px-3 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">عملیات</th>
                </tr>
                                </thead>
                  <tbody class="bg-white divide-y divide-gray-200">
                    <tr v-for="wishlist in paginatedWishlists" :key="wishlist.id" class="hover:bg-gray-50 transition-colors">
                  <!-- Original table content for user view -->
                  <!-- [Keep the existing user-centric table content] -->
                  <!-- Checkbox -->
                  <td class="px-3 py-3 whitespace-nowrap text-center">
                    <input 
                      v-model="selectedItems" 
                      type="checkbox"
                      :value="wishlist.id"
                      class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                    />
                  </td>

                  <!-- Product Image -->
                  <td class="px-3 py-3 whitespace-nowrap text-center">
                    <img 
                      :src="wishlist.product.image" 
                      :alt="wishlist.product.name"
                      class="h-12 w-12 rounded-md object-cover mx-auto border border-gray-200"
                    />
                  </td>

                  <!-- Product Name -->
                  <td class="px-3 py-3 whitespace-nowrap text-right">
                    <div class="text-xs font-medium text-gray-900">{{ wishlist.product.name }}</div>
                    <div class="text-xs text-gray-500">کد: {{ wishlist.product.sku }}</div>
                  </td>

                  <!-- Price -->
                  <td class="px-3 py-3 whitespace-nowrap text-center">
                    <div class="text-xs font-medium text-gray-900">{{ formatPrice(wishlist.product.price) }}</div>
                    <div v-if="wishlist.product.discount" class="text-xs text-red-600">
                      {{ wishlist.product.discount }}% تخفیف
                    </div>
                  </td>

                  <!-- Category -->
                  <td class="px-3 py-3 whitespace-nowrap text-center">
                    <span class="inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium bg-gray-100 text-gray-800">
                      {{ wishlist.product.category }}
                    </span>
                  </td>

                  <!-- User -->
                  <td class="px-3 py-3 whitespace-nowrap text-center">
                    <div class="text-xs font-medium text-gray-900">{{ wishlist.user.name }}</div>
                    <div class="text-xs text-gray-500">{{ wishlist.user.email }}</div>
                  </td>

                  <!-- Date Added -->
                  <td class="px-3 py-3 whitespace-nowrap text-center">
                    <div class="text-xs text-gray-900">{{ formatDate(wishlist.created_at) }}</div>
                    <div class="text-xs text-gray-500">{{ getRelativeTime(wishlist.created_at) }}</div>
                  </td>

                  <!-- Wishlist Count -->
                  <td class="px-3 py-3 whitespace-nowrap text-center">
                    <span class="inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800">
                      {{ wishlist.product.wishlist_count }}
                    </span>
                  </td>

                  <!-- Stock Status -->
                  <td class="px-3 py-3 whitespace-nowrap text-center">
                    <span 
                      :class="[
                        'inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium',
                        wishlist.product.in_stock 
                          ? 'bg-green-100 text-green-800' 
                          : 'bg-red-100 text-red-800'
                      ]"
                    >
                      {{ wishlist.product.in_stock ? 'موجود' : 'ناموجود' }}
                    </span>
                  </td>

                  <!-- Actions -->
                  <td class="px-3 py-3 whitespace-nowrap text-center">
                    <div class="flex items-center justify-center space-x-1 space-x-reverse">
                      <button 
                        class="inline-flex items-center p-1.5 border border-transparent text-xs font-medium rounded text-blue-700 bg-blue-100 hover:bg-blue-200 transition-colors"
                        title="مشاهده محصول"
                        @click="viewProduct(wishlist.product.id)"
                      >
                        👁️
                      </button>
                      <button 
                        class="inline-flex items-center p-1.5 border border-transparent text-xs font-medium rounded text-green-700 bg-green-100 hover:bg-green-200 transition-colors"
                        title="ارسال اعلان"
                        @click="sendNotification(wishlist.user.id, wishlist.product.id)"
                      >
                        📧
                      </button>
                      <button 
                        class="inline-flex items-center p-1.5 border border-transparent text-xs font-medium rounded text-red-700 bg-red-100 hover:bg-red-200 transition-colors"
                        title="حذف از لیست"
                        @click="removeFromWishlist(wishlist.id)"
                      >
                        🗑️
                      </button>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>

          </div>
        </div>

        <!-- Pagination -->
        <div class="bg-gray-50 px-4 py-3 border-t border-gray-200">
          <div class="flex items-center justify-between">
            <div class="flex items-center text-sm text-gray-700">
              <span class="font-medium">{{ paginationInfo.start }}</span>
              <span class="mx-1">تا</span>
              <span class="font-medium">{{ paginationInfo.end }}</span>
              <span class="mx-1">از</span>
              <span class="font-medium">{{ paginationInfo.total }}</span>
              <span class="mr-1">علاقمندی</span>
            </div>
            
            <div class="flex items-center space-x-2 space-x-reverse">
              <div class="flex items-center space-x-1 space-x-reverse text-sm text-gray-600">
                <span>نمایش</span>
                <select 
                  v-model="itemsPerPage" 
                  class="border border-gray-300 rounded px-2 py-1 text-xs bg-white focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  @change="updatePagination"
                >
                  <option value="10">10</option>
                  <option value="25">25</option>
                  <option value="50">50</option>
                  <option value="100">100</option>
                </select>
                <span>در هر صفحه</span>
              </div>
              
              <div class="flex items-center space-x-1 space-x-reverse">
                <button 
                  :disabled="currentPage <= 1"
                  class="px-2 py-1 text-xs border rounded transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
                  @click="goToPage(currentPage - 1)"
                >
                  قبلی
                </button>
                
                <div class="flex items-center space-x-1 space-x-reverse">
                  <template v-for="page in visiblePages" :key="page">
                    <button 
                      v-if="page !== '...'"
                      :class="[
                        'px-2 py-1 text-xs border rounded transition-colors',
                        currentPage === page 
                          ? 'border-blue-500 text-white bg-blue-600' 
                          : 'border-gray-300 text-gray-700 bg-white hover:bg-gray-50'
                      ]"
                      @click="goToPage(Number(page))"
                    >
                      {{ page }}
                    </button>
                    <span v-else class="px-2 py-1 text-xs text-gray-400">...</span>
                  </template>
                </div>
                
                <button 
                  :disabled="currentPage >= totalPages"
                  class="px-2 py-1 text-xs border rounded transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
                  @click="goToPage(currentPage + 1)"
                >
                  بعدی
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <template #fallback>
      <div class="min-h-screen flex items-center justify-center">
        <div class="text-center">
          <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600 mx-auto mb-3"></div>
          <div class="text-gray-600 text-sm">در حال بارگذاری...</div>
        </div>
      </div>
    </template>
  </ClientOnly>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
declare const useHead: (head: { title?: string }) => void
declare const navigateTo: (to: string) => Promise<void>
declare const useAsyncData: <T>(key: string, fn: () => Promise<T>, options?: { watch?: unknown[] }) => Promise<{ data: { value: T } }>
declare const $fetch: <T = unknown>(url: string, options?: { method?: string; body?: unknown; params?: Record<string, unknown> }) => Promise<T>
</script>

<script setup lang="ts">
import ExportExcelButton from '@/components/common/ExportExcelButton.vue'
import { computed, ref, watch } from 'vue'
import { useConfirmDialog } from '~/composables/useConfirmDialog'
import { useNotifier } from '~/composables/useNotifier'

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

useHead({
  title: 'مدیریت لیست علاقمندی'
})

const notifier = useNotifier()
const { confirm } = useConfirmDialog()

const showFilters = ref(false)

interface Product {
  id: number
  name: string
  sku: string
  price: number
  discount?: number
  category: string
  image: string
  wishlist_count: number
  in_stock: boolean
}

interface User {
  id: number
  name: string
  email: string
}

interface Wishlist {
  id: number
  product: Product
  user: User
  created_at: string
}

// Reactive data
const searchQuery = ref('')
const selectedCategory = ref('')
const priceRange = ref('')
const dateRange = ref('')
const selectedItems = ref<number[]>([])
const selectedProductIds = ref<number[]>([])
const expandedProducts = ref<number[]>([])
const currentPage = ref(1)
const itemsPerPage = ref(25)
const viewMode = ref('product')
const currentSortColumn = ref('wishlist_count')
const sortDirection = ref('desc')
const sortKey = ref(0) // Add a key to force reactivity
const showBulkMessageModal = ref(false)

// داده‌های مربوط به ارسال پیام گروهی
const smsType = ref('discount')
const discountPercent = ref<number>(20)
const smsMessage = ref('')

// محاسبه گیرندگان بر اساس انتخاب‌ها
const bulkRecipientUserIds = computed<number[]>(() => {
  if (viewMode.value === 'product') {
    const selectedSet = new Set(selectedProductIds.value)
    const userIds = new Set<number>()
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    productGroups.value.forEach((group: any) => {
      if (selectedSet.has(group.product.id)) {
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        ;(group.users || []).forEach((uw: any) => {
          if (uw?.user?.id) userIds.add(uw.user.id)
        })
      }
    })
    return Array.from(userIds)
  } else {
    const selectedWishlistSet = new Set(selectedItems.value)
    const userIds = new Set<number>()
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    wishlists.value.forEach((w: any) => {
      if (selectedWishlistSet.has(w.id) && w?.user?.id) userIds.add(w.user.id)
    })
    return Array.from(userIds)
  }
})

const bulkRecipientsCount = computed(() => bulkRecipientUserIds.value.length)

const getBulkSmsTemplate = () => {
  switch (smsType.value) {
    case 'discount':
      return `سلام، روی برخی از محصولات مورد علاقه شما ${discountPercent.value}% تخفیف اعمال شد!`
    case 'availability':
      return 'سلام، برخی از محصولات مورد علاقه شما موجود شدند!'
    case 'price_drop':
      return 'سلام، قیمت برخی از محصولات مورد علاقه شما کاهش یافت!'
    default:
      return ''
  }
}

watch([smsType, discountPercent], () => {
  if (smsType.value !== 'custom') {
    smsMessage.value = getBulkSmsTemplate()
  }
})

const openBulkMessageModal = () => {
  if (bulkRecipientsCount.value === 0) {
    notifier.warning('هیچ گیرنده‌ای انتخاب نشده است.')
    return
  }
  smsMessage.value = getBulkSmsTemplate()
  showBulkMessageModal.value = true
}

const sendBulkMessage = async () => {
  try {
    if (bulkRecipientsCount.value === 0) {
      notifier.warning('هیچ گیرنده‌ای انتخاب نشده است.')
      return
    }
    // اتصال به سرویس اعلان مرکزی
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const res: any = await $fetch('/api/admin/notifications/jobs', {
      method: 'POST',
      body: {
        channel: 'sms',
        templateId: null,
        message: smsMessage.value,
        variables: {
          discountPercent: smsType.value === 'discount' ? discountPercent.value : undefined,
          productIds: viewMode.value === 'product' ? selectedProductIds.value : []
        },
        recipients: bulkRecipientUserIds.value,
        scheduleAt: null,
        meta: { origin: 'admin-wishlists' }
      }
    })
    showBulkMessageModal.value = false
    if (res?.jobId) {
      const go = await confirm({ title:'ثبت شد', message:`ارسال پیام ثبت شد. کد پیگیری: ${res.jobId}\nآیا می‌خواهید وضعیت ارسال را مشاهده کنید؟`, confirmText:'بله', cancelText:'خیر', type:'info' })
      if (go) {
        await navigateTo(`/admin/notifications/monitoring?jobId=${encodeURIComponent(res.jobId)}`)
      }
    } else {
      notifier.success('ارسال پیام گروهی با موفقیت ثبت شد.')
    }
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  } catch (e: any) {
    notifier.error(e?.data?.message || 'خطا در ارسال پیام گروهی')
  }
}

// Sample data - with more realistic data showing multiple users per product
const wishlists = ref<Wishlist[]>([
  {
    id: 1,
    product: {
      id: 101,
      name: 'گوشی سامسونگ گلکسی A54',
      sku: 'SAM-A54-128',
      price: 8500000,
      discount: 15,
      category: 'الکترونیکی',
      image: 'https://via.placeholder.com/100x100?text=Phone',
      wishlist_count: 245,
      in_stock: true
    },
    user: {
      id: 1001,
      name: 'علی احمدی',
      email: 'ali@example.com'
    },
    created_at: '2024-01-15T10:30:00Z'
  },
  {
    id: 2,
    product: {
      id: 101, // Same product as above
      name: 'گوشی سامسونگ گلکسی A54',
      sku: 'SAM-A54-128',
      price: 8500000,
      discount: 15,
      category: 'الکترونیکی',
      image: 'https://via.placeholder.com/100x100?text=Phone',
      wishlist_count: 245,
      in_stock: true
    },
    user: {
      id: 1006,
      name: 'فاطمه رضایی',
      email: 'fatemeh@example.com'
    },
    created_at: '2024-01-16T14:20:00Z'
  },
  {
    id: 3,
    product: {
      id: 101, // Same product as above
      name: 'گوشی سامسونگ گلکسی A54',
      sku: 'SAM-A54-128',
      price: 8500000,
      discount: 15,
      category: 'الکترونیکی',
      image: 'https://via.placeholder.com/100x100?text=Phone',
      wishlist_count: 245,
      in_stock: true
    },
    user: {
      id: 1007,
      name: 'حمید کریمی',
      email: 'hamid@example.com'
    },
    created_at: '2024-01-14T09:15:00Z'
  },
  {
    id: 4,
    product: {
      id: 102,
      name: 'لپ‌تاپ ایسوس VivoBook',
      sku: 'ASUS-VB15-512',
      price: 25000000,
      category: 'الکترونیکی',
      image: 'https://via.placeholder.com/100x100?text=Laptop',
      wishlist_count: 189,
      in_stock: false
    },
    user: {
      id: 1002,
      name: 'مریم محمدی',
      email: 'maryam@example.com'
    },
    created_at: '2024-01-14T14:20:00Z'
  },
  {
    id: 5,
    product: {
      id: 102, // Same product as above
      name: 'لپ‌تاپ ایسوس VivoBook',
      sku: 'ASUS-VB15-512',
      price: 25000000,
      category: 'الکترونیکی',
      image: 'https://via.placeholder.com/100x100?text=Laptop',
      wishlist_count: 189,
      in_stock: false
    },
    user: {
      id: 1008,
      name: 'سارا احمدی',
      email: 'sara@example.com'
    },
    created_at: '2024-01-13T16:45:00Z'
  },
  {
    id: 6,
    product: {
      id: 103,
      name: 'کفش ورزشی نایک ایر مکس',
      sku: 'NIKE-AM90-42',
      price: 3200000,
      discount: 20,
      category: 'ورزشی',
      image: 'https://via.placeholder.com/100x100?text=Shoes',
      wishlist_count: 456,
      in_stock: true
    },
    user: {
      id: 1003,
      name: 'حسن رضایی',
      email: 'hassan@example.com'
    },
    created_at: '2024-01-13T09:15:00Z'
  },
  {
    id: 7,
    product: {
      id: 104,
      name: 'کتاب برنامه‌نویسی JavaScript',
      sku: 'BOOK-JS-2024',
      price: 450000,
      category: 'کتاب',
      image: 'https://via.placeholder.com/100x100?text=Book',
      wishlist_count: 78,
      in_stock: true
    },
    user: {
      id: 1004,
      name: 'زهرا حسینی',
      email: 'zahra@example.com'
    },
    created_at: '2024-01-12T16:45:00Z'
  },
  {
    id: 8,
    product: {
      id: 105,
      name: 'ساعت هوشمند اپل واچ',
      sku: 'APPLE-WATCH-S9',
      price: 12000000,
      category: 'الکترونیکی',
      image: 'https://via.placeholder.com/100x100?text=Watch',
      wishlist_count: 322,
      in_stock: true
    },
    user: {
      id: 1005,
      name: 'محمد کریمی',
      email: 'mohammad@example.com'
    },
    created_at: '2024-01-11T11:30:00Z'
  },
  {
    id: 9,
    product: {
      id: 103, // Same as Nike Air Max
      name: 'کفش ورزشی نایک ایر مکس',
      sku: 'NIKE-AM90-42',
      price: 3200000,
      discount: 20,
      category: 'ورزشی',
      image: 'https://via.placeholder.com/100x100?text=Shoes',
      wishlist_count: 456,
      in_stock: true
    },
    user: {
      id: 1009,
      name: 'پویا مرادی',
      email: 'pouya@example.com'
    },
    created_at: '2024-01-10T08:20:00Z'
  },
  {
    id: 10,
    product: {
      id: 103, // Same as Nike Air Max
      name: 'کفش ورزشی نایک ایر مکس',
      sku: 'NIKE-AM90-42',
      price: 3200000,
      discount: 20,
      category: 'ورزشی',
      image: 'https://via.placeholder.com/100x100?text=Shoes',
      wishlist_count: 456,
      in_stock: true
    },
    user: {
      id: 1010,
      name: 'نیلوفر حسنی',
      email: 'niloofar@example.com'
    },
    created_at: '2024-01-09T12:30:00Z'
  },
  {
    id: 11,
    product: {
      id: 103, // Same as Nike Air Max
      name: 'کفش ورزشی نایک ایر مکس',
      sku: 'NIKE-AM90-42',
      price: 3200000,
      discount: 20,
      category: 'ورزشی',
      image: 'https://via.placeholder.com/100x100?text=Shoes',
      wishlist_count: 456,
      in_stock: true
    },
    user: {
      id: 1011,
      name: 'امین صادقی',
      email: 'amin@example.com'
    },
    created_at: '2024-01-08T15:45:00Z'
  },
  {
    id: 12,
    product: {
      id: 103, // Same as Nike Air Max
      name: 'کفش ورزشی نایک ایر مکس',
      sku: 'NIKE-AM90-42',
      price: 3200000,
      discount: 20,
      category: 'ورزشی',
      image: 'https://via.placeholder.com/100x100?text=Shoes',
      wishlist_count: 456,
      in_stock: true
    },
    user: {
      id: 1012,
      name: 'مهسا رحیمی',
      email: 'mahsa@example.com'
    },
    created_at: '2024-01-07T10:15:00Z'
  },
  {
    id: 13,
    product: {
      id: 103, // Same as Nike Air Max
      name: 'کفش ورزشی نایک ایر مکس',
      sku: 'NIKE-AM90-42',
      price: 3200000,
      discount: 20,
      category: 'ورزشی',
      image: 'https://via.placeholder.com/100x100?text=Shoes',
      wishlist_count: 456,
      in_stock: true
    },
    user: {
      id: 1013,
      name: 'رضا نوری',
      email: 'reza@example.com'
    },
    created_at: '2024-01-06T14:25:00Z'
  },
  {
    id: 14,
    product: {
      id: 106,
      name: 'تبلت سامسونگ Galaxy Tab',
      sku: 'SAM-TAB-64',
      price: 6500000,
      category: 'الکترونیکی',
      image: 'https://via.placeholder.com/100x100?text=Tablet',
      wishlist_count: 89,
      in_stock: false
    },
    user: {
      id: 1014,
      name: 'الهام حسنی',
      email: 'elham@example.com'
    },
    created_at: '2024-01-18T16:30:00Z'
  },
  {
    id: 15,
    product: {
      id: 107,
      name: 'هدفون بی‌سیم AirPods',
      sku: 'APPLE-AIRPODS-3',
      price: 4800000,
      discount: 10,
      category: 'الکترونیکی',
      image: 'https://via.placeholder.com/100x100?text=AirPods',
      wishlist_count: 278,
      in_stock: true
    },
    user: {
      id: 1015,
      name: 'کیان احمدی',
      email: 'kian@example.com'
    },
    created_at: '2024-01-19T08:45:00Z'
  },
  {
    id: 16,
    product: {
      id: 107, // Same as AirPods
      name: 'هدفون بی‌سیم AirPods',
      sku: 'APPLE-AIRPODS-3',
      price: 4800000,
      discount: 10,
      category: 'الکترونیکی',
      image: 'https://via.placeholder.com/100x100?text=AirPods',
      wishlist_count: 278,
      in_stock: true
    },
    user: {
      id: 1016,
      name: 'شیدا رضایی',
      email: 'sheida@example.com'
    },
    created_at: '2024-01-20T11:20:00Z'
  },
  {
    id: 17,
    product: {
      id: 108,
      name: 'خودکار ساده',
      sku: 'PEN-SIMPLE-01',
      price: 5000, // Very cheap product
      category: 'لوازم‌التحریر',
      image: 'https://via.placeholder.com/100x100?text=Pen',
      wishlist_count: 12,
      in_stock: true
    },
    user: {
      id: 1017,
      name: 'احمد نوری',
      email: 'ahmad@example.com'
    },
    created_at: '2024-01-21T09:30:00Z'
  }
])

// Computed properties
const totalWishlists = computed(() => wishlists.value.length)
const uniqueProducts = computed(() => {
  const productIds = new Set(wishlists.value.map(w => w.product.id))
  return productIds.size
})
const activeUsers = computed(() => new Set(wishlists.value.map(w => w.user.id)).size)
const conversionRate = computed(() => 23.5) // Mock conversion rate

const mostRequestedProduct = computed(() => {
  if (productGroups.value.length === 0) return null
  
  return productGroups.value.reduce((mostRequested, current) => {
    return current.users.length > mostRequested.users.length ? current : mostRequested
  })
})

const filteredWishlists = computed(() => {
  let filtered = wishlists.value

  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(w => 
      w.product.name.toLowerCase().includes(query) ||
      w.product.sku.toLowerCase().includes(query) ||
      w.user.name.toLowerCase().includes(query)
    )
  }

  if (selectedCategory.value) {
    filtered = filtered.filter(w => w.product.category === selectedCategory.value)
  }

  return filtered
})

// Product-centric computed properties
const productGroups = computed(() => {
  // اگر داده‌ی سروری برای گروه محصولات داریم از آن استفاده کن
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const data = wishlistProductsData.value as { items?: any[] } | null
  if (data && Array.isArray(data.items) && data.items.length > 0) {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    return data.items.map((item: any) => ({
      product: {
        id: item.product?.id,
        name: item.product?.name,
        sku: item.product?.sku,
        price: item.product?.price,
        discount: item.product?.discount,
        category: item.product?.category,
        image: item.product?.image,
        wishlist_count: item.users?.length || 0,
        in_stock: item.product?.in_stock,
      },
      users: item.users || [],
      latestDate: item.latestDate || item.latest_date || ''
    }))
  }

  // حالت fallback: محاسبه از روی داده‌های Mock
  const groups = new Map()
  wishlists.value.forEach(wishlist => {
    const productId = wishlist.product.id
    if (!groups.has(productId)) {
      groups.set(productId, {
        product: wishlist.product,
        users: [],
        latestDate: wishlist.created_at
      })
    }
    const group = groups.get(productId)
    group.users.push(wishlist)
    if (new Date(wishlist.created_at) > new Date(group.latestDate)) {
      group.latestDate = wishlist.created_at
    }
  })
  return Array.from(groups.values())
})

const filteredProductGroups = computed(() => {
  // Access sortKey to force reactivity
  sortKey.value
  
  let filtered = productGroups.value

  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(group => 
      group.product.name.toLowerCase().includes(query) ||
      group.product.sku.toLowerCase().includes(query) ||
      group.users.some((userWishlist: Wishlist) => userWishlist.user.name.toLowerCase().includes(query))
    )
  }

  if (selectedCategory.value) {
    filtered = filtered.filter(group => group.product.category === selectedCategory.value)
  }

  // Clone the array to avoid mutating the original
  filtered = [...filtered]

  // Always use column-based sorting
  filtered.sort((a, b) => {
    let aValue, bValue
    
    switch (currentSortColumn.value) {
      case 'name':
        aValue = a.product.name.toLowerCase()
        bValue = b.product.name.toLowerCase()
        break
      case 'price':
        aValue = a.product.price
        bValue = b.product.price
        break
      case 'wishlist_count':
        aValue = a.users.length
        bValue = b.users.length
        break
      case 'latest_date':
        aValue = new Date(a.latestDate).getTime()
        bValue = new Date(b.latestDate).getTime()
        break
      case 'stock_status':
        aValue = a.product.in_stock ? 1 : 0
        bValue = b.product.in_stock ? 1 : 0
        break
      default:
        aValue = a.users.length
        bValue = b.users.length
    }

    // String comparison for text fields
    if (typeof aValue === 'string' && typeof bValue === 'string') {
      if (sortDirection.value === 'asc') {
        return aValue.localeCompare(bValue, 'fa')
      } else {
        return bValue.localeCompare(aValue, 'fa')
      }
    }
    
    // Numeric comparison
    if (sortDirection.value === 'asc') {
      return aValue - bValue
    } else {
      return bValue - aValue
    }
  })

  return filtered
})

const totalPages = computed(() => {
  if (viewMode.value === 'product') {
    return Math.ceil(filteredProductGroups.value.length / itemsPerPage.value)
  } else {
    return Math.ceil(filteredWishlists.value.length / itemsPerPage.value)
  }
})

const paginatedWishlists = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage.value
  const end = start + itemsPerPage.value
  return filteredWishlists.value.slice(start, end)
})

const paginatedProductGroups = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage.value
  const end = start + itemsPerPage.value
  return filteredProductGroups.value.slice(start, end)
})

const paginationInfo = computed(() => {
  if (viewMode.value === 'product') {
    const start = (currentPage.value - 1) * itemsPerPage.value + 1
    const end = Math.min(start + itemsPerPage.value - 1, filteredProductGroups.value.length)
    return {
      start,
      end,
      total: filteredProductGroups.value.length
    }
  } else {
    const start = (currentPage.value - 1) * itemsPerPage.value + 1
    const end = Math.min(start + itemsPerPage.value - 1, filteredWishlists.value.length)
    return {
      start,
      end,
      total: filteredWishlists.value.length
    }
  }
})

const visiblePages = computed(() => {
  const pages = []
  const total = totalPages.value
  const current = currentPage.value
  
  if (total <= 7) {
    for (let i = 1; i <= total; i++) {
      pages.push(i)
    }
  } else {
    if (current <= 4) {
      for (let i = 1; i <= 5; i++) pages.push(i)
      pages.push('...')
      pages.push(total)
    } else if (current >= total - 3) {
      pages.push(1)
      pages.push('...')
      for (let i = total - 4; i <= total; i++) pages.push(i)
    } else {
      pages.push(1)
      pages.push('...')
      for (let i = current - 1; i <= current + 1; i++) pages.push(i)
      pages.push('...')
      pages.push(total)
    }
  }
  
  return pages
})

const isAllSelected = computed(() => {
  return paginatedWishlists.value.length > 0 && 
         selectedItems.value.length === paginatedWishlists.value.length
})

const isAllProductsSelected = computed(() => {
  return paginatedProductGroups.value.length > 0 && 
         selectedProductIds.value.length === paginatedProductGroups.value.length
})

// Methods
const performSearch = () => {
  currentPage.value = 1
}

const clearFilters = () => {
  searchQuery.value = ''
  selectedCategory.value = ''
  priceRange.value = ''
  dateRange.value = ''
  currentSortColumn.value = 'wishlist_count'
  sortDirection.value = 'desc'
  currentPage.value = 1
}

const toggleSelectAll = () => {
  if (isAllSelected.value) {
    selectedItems.value = []
  } else {
    selectedItems.value = paginatedWishlists.value.map(w => w.id)
  }
}

// const selectAll = () => {
//   selectedItems.value = filteredWishlists.value.map(w => w.id)
// }

// const bulkRemove = async () => {
//   const ok = await confirm({ title:'حذف گروهی', message:`آیا از حذف ${selectedItems.value.length} آیتم انتخاب شده اطمینان دارید؟`, confirmText:'حذف', cancelText:'انصراف', type:'danger' })
//   if (!ok) return
//   wishlists.value = wishlists.value.filter(w => !selectedItems.value.includes(w.id))
//   selectedItems.value = []
// }

const goToPage = (page: number) => {
  if (page >= 1 && page <= totalPages.value) {
    currentPage.value = page
  }
}

const updatePagination = () => {
  currentPage.value = 1
}

const formatPrice = (price: number) => {
  return new Intl.NumberFormat('fa-IR').format(price) + ' تومان'
}

const formatDate = (dateString: string) => {
  const date = new Date(dateString)
  return new Intl.DateTimeFormat('fa-IR', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit'
  }).format(date)
}

const getRelativeTime = (dateString: string) => {
  const date = new Date(dateString)
  const now = new Date()
  const diffTime = Math.abs(now.getTime() - date.getTime())
  const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24))
  
  if (diffDays === 1) return 'دیروز'
  if (diffDays <= 7) return `${diffDays} روز پیش`
  if (diffDays <= 30) return `${Math.ceil(diffDays / 7)} هفته پیش`
  return `${Math.ceil(diffDays / 30)} ماه پیش`
}

const viewProduct = (productId: number) => {
  window.open(`/products/${productId}`, '_blank')
}

const sendNotification = (userId: number, productId: number) => {
  notifier.success('اعلان ارسال شد!')
}

const removeFromWishlist = async (wishlistId: number) => {
  const ok = await confirm({ title:'تأیید حذف', message:'آیا از حذف این آیتم از لیست علاقمندی اطمینان دارید؟', confirmText:'حذف', cancelText:'انصراف', type:'danger' })
  if (!ok) return
  const index = wishlists.value.findIndex(w => w.id === wishlistId)
  if (index > -1) {
    wishlists.value.splice(index, 1)
  }
}

// const exportWishlists = () => {
//   notifier.info('خروجی اکسل در حال تهیه است...')
// }

const sendWishlistEmail = () => {
  // هدایت به سیستم ارسال پیام/اعلان موجود
  navigateTo('/admin/transactions/order-surveys2')
}

const toggleSelectAllProducts = () => {
  if (isAllProductsSelected.value) {
    selectedProductIds.value = []
  } else {
    selectedProductIds.value = paginatedProductGroups.value.map(group => group.product.id)
  }
}

const toggleProductExpansion = (productId: number) => {
  const index = expandedProducts.value.indexOf(productId)
  if (index > -1) {
    expandedProducts.value.splice(index, 1)
  } else {
    expandedProducts.value.push(productId)
  }
}

const notifyAllUsers = (productId: number) => {
  // Implement notification logic here
}

const removeProductFromAllWishlists = (productId: number) => {
  // Implement removal logic here
  wishlists.value = wishlists.value.filter(w => w.product.id !== productId)
}

const manageProductUsers = (productId: number) => {
  // هدایت به صفحه مدیریت کاربران علاقه‌مند برای یک محصول
  navigateTo(`/admin/product-management/wishlists/product/${productId}`)
}

const getRequestLevelText = (count: number) => {
  if (count >= 5) return 'درخواست بالا'
  if (count >= 3) return 'درخواست متوسط'
  return 'درخواست پایین'
}

const sortByColumn = (column: string) => {
  if (currentSortColumn.value === column) {
    // Toggle direction if same column
    sortDirection.value = sortDirection.value === 'asc' ? 'desc' : 'asc'
  } else {
    // Set new column and default to descending
    currentSortColumn.value = column
    sortDirection.value = 'desc'
  }
  
  // Force reactivity
  sortKey.value++
  
  // Reset page to 1 when sorting
  currentPage.value = 1
}

// بارگذاری داده‌های تجمیعی علاقه‌مندی‌ها (نمای محصول) از Nuxt API (Mock فعلی)
const { data: wishlistProductsData } = await useAsyncData(
  'admin-wishlist-products',
  () => $fetch('/api/admin/wishlists/products', {
    params: {
      page: currentPage.value,
      pageSize: itemsPerPage.value,
      sort: currentSortColumn.value,
      order: sortDirection.value,
      q: searchQuery.value,
      category: selectedCategory.value
    }
  }),
  { watch: [currentPage, itemsPerPage, currentSortColumn, sortDirection, searchQuery, selectedCategory] }
)

// داده خروجی اکسل بر اساس نمای فعال
const exportData = computed(() => {
  if (viewMode.value === 'product') {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    return productGroups.value.map((g: any) => ({
      id: g.product.id,
      name: g.product.name,
      sku: g.product.sku,
      category: g.product.category,
      price: g.product.price,
      wishlist_count: g.users.length,
      latest_date: g.latestDate,
      in_stock: g.product.in_stock ? 'موجود' : 'ناموجود'
    }))
  }
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  return filteredWishlists.value.map((w: any) => ({
    id: w.id,
    product_id: w.product.id,
    product_name: w.product.name,
    sku: w.product.sku,
    user_id: w.user.id,
    user_name: w.user.name,
    user_email: w.user.email,
    price: w.product.price,
    category: w.product.category,
    created_at: w.created_at,
    in_stock: w.product.in_stock ? 'موجود' : 'ناموجود'
  }))
})
</script> 

<style scoped>
.bulk-modal-backdrop { position: fixed; inset: 0; background: rgba(0,0,0,0.5); display: flex; align-items: center; justify-content: center; z-index: 50; }
.bulk-modal { background: #fff; border-radius: 0.5rem; box-shadow: 0 10px 25px rgba(0,0,0,0.1); padding: 1.5rem; width: 100%; max-width: 36rem; margin: 1rem; }
</style>
