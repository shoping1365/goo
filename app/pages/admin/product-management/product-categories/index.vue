<template>
  <ClientOnly>
    <!-- Page Header -->
    <div class="bg-white shadow-sm border-b border-gray-200 mb-4">
      <div class="w-full px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center py-3">
  <div>
            <h1 class="text-lg font-bold text-gray-900">دسته‌بندی‌های محصولات</h1>
            <p class="text-xs text-gray-600 mt-1">مدیریت کامل دسته‌بندی‌ها و ساختار محصولات</p>
          </div>
          <div class="flex space-x-2 space-x-reverse">
          <button v-if="hasPermission('categories_manage')" class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-blue-400 to-blue-600 hover:from-blue-500 hover:to-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105" @click="showTypeModal = true">
            <svg fill='none' viewBox='0 0 24 24' stroke-width='1.5' stroke='currentColor' class='w-5 h-5 ml-2'>
              <path stroke-linecap='round' stroke-linejoin='round' d='M12 4v16m8-8H4' />
            </svg>
            مورد جدید
          </button>
            <button class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-green-400 to-green-600 hover:from-green-500 hover:to-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105">
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
            </svg>
              خروجی Excel
        </button>
          <button class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-gray-800 bg-gradient-to-r from-gray-200 to-gray-400 hover:from-gray-300 hover:to-gray-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-400 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105">
            <svg fill='none' viewBox='0 0 24 24' stroke-width='1.5' stroke='currentColor' class='w-5 h-5 ml-2'>
              <path stroke-linecap='round' stroke-linejoin='round' d='M3 16v2a2 2 0 002 2h14a2 2 0 002-2v-2M16 10l-4-4m0 0l-4 4m4-4v12' />
            </svg>
              ورود اطلاعات
        </button>
          </div>
        </div>
      </div>
    </div>

    <div class="w-full px-4 py-4">
      
      <!-- آمار کلی - کارت اول -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-6">
        <TemplateCard 
          title="کل دسته‌بندی‌ها" 
          :value="categories.length" 
          variant="indigo"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              
            </svg>
          </template>
        </TemplateCard>

        <TemplateCard 
          title="دسته‌بندی اصلی" 
          :value="mainCategories.length" 
          variant="purple"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7v10c0 2.21 3.582 4 8 4s8-1.79 8-4V7M4 7c0 2.21 3.582 4 8 4s8-1.79 8-4M4 7c0-2.21 3.582-4 8-4s8 1.79 8 4"></path>
            </svg>
          </template>
        </TemplateCard>

        <TemplateCard 
          title="دسته‌بندی فرعی" 
          :value="subCategories.length" 
          variant="green"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
            </svg>
          </template>
        </TemplateCard>

        <TemplateCard 
          title="منتشر شده" 
          :value="publishedCategories.length" 
          variant="blue"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </template>
        </TemplateCard>
      </div>

      <!-- Bulk Actions Toolbar -->
      <div v-if="canDeleteSelected()" class="bg-blue-50 border border-blue-200 rounded-lg p-6 mb-6">
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-6">
            <span class="text-sm font-medium text-blue-800">
              {{ selectedCategories.size }} دسته‌بندی انتخاب شده
            </span>
            <button 
              class="text-sm text-blue-600 hover:text-blue-800 underline"
              @click="selectedCategories.clear(); isAllSelected = false"
            >
              لغو انتخاب
            </button>
          </div>
          <div class="flex items-center gap-2">
            <button 
              v-if="canDeleteCategory"
              class="inline-flex items-center px-4 py-2 bg-red-600 hover:bg-red-700 text-white text-sm font-medium rounded-lg transition-colors"
              @click="bulkDeleteCategories"
            >
              <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
              </svg>
              حذف انتخاب شده‌ها
            </button>
          </div>
        </div>
      </div>

      <!-- جستجو و فیلتر -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden mb-6">
        <div class="bg-gradient-to-r from-slate-50 to-gray-50 px-4 py-3 border-b border-gray-200">
          <div class="flex items-center">
            <div class="w-6 h-6 bg-slate-600 rounded-md flex items-center justify-center ml-2">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
              </svg>
            </div>
            <h3 class="text-sm font-semibold text-gray-900">جستجو و فیلتر</h3>
          </div>
        </div>
        
        <div class="p-6">
      <div class="flex flex-col md:flex-row items-center gap-6">
        <!-- Search input -->
        <div class="flex-1 w-full">
              <input id="category-name" v-model="searchTerm" type="text" class="w-full rounded-lg border border-gray-200 shadow-sm bg-white focus:outline-none focus:ring-2 focus:ring-blue-400 py-3 px-6 text-gray-800 text-base transition-all duration-200" placeholder="نام دسته‌بندی را وارد کنید..." @keyup.enter="applySearch">
        </div>

        <!-- Filter select -->
            <div class="flex items-center gap-2">
              <label class="text-sm font-medium text-gray-700">فیلتر:</label>
              <select v-model="filterOption" class="rounded-lg border border-gray-200 py-3 px-4 bg-white text-gray-800 text-sm focus:outline-none focus:ring-2 focus:ring-blue-400 shadow-sm">
          <option value="all">همه</option>
          <option value="main">دسته‌بندی اصلی</option>
          <option value="sub">دسته‌بندی فرعی</option>
          <option value="published">منتشر شده</option>
          <option value="unpublished">منتشر نشده</option>
        </select>
            </div>
          </div>
        </div>
      </div>

      <!-- مدیریت دسته‌بندی‌ها - کارت اصلی -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
        <div class="bg-gradient-to-r from-slate-50 to-gray-50 px-4 py-3 border-b border-gray-200">
          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="w-6 h-6 bg-slate-600 rounded-md flex items-center justify-center ml-2">
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16"></path>
                </svg>
              </div>
              <h3 class="text-sm font-semibold text-gray-900">فهرست دسته‌بندی‌ها</h3>
            </div>
            
            <div class="flex items-center">
              <span class="text-xs text-gray-600">{{ displayedCategories.length }} دسته‌بندی</span>
            </div>
      </div>
    </div>

    <!-- Categories Table -->
        <div class="overflow-hidden">
      <table class="min-w-full divide-y divide-gray-200 table-fixed">
        <thead class="bg-gray-50">
          <tr>
            <th class="w-1/12 px-4 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">
              <input 
                type="checkbox" 
                :checked="isAllSelected"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                @change="toggleSelectAll"
              />
            </th>
            <th class="w-1/12 px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تصویر</th>
            <th class="w-1/12 px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">آیدی</th>
            <th class="w-2/12 px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نام</th>
            <th class="w-1/12 px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
            <th class="w-1/12 px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">محصولات</th>
            <th class="w-1/12 px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">منتشر شده</th>
            <th class="w-1/12 px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">ترتیب</th>
            <th class="w-3/12 px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">دسته پدر</th>
            <th class="w-1/12 px-4 py-3 text-center text-xs font-medium text-gray-600">عملیات</th>
          </tr>
        </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="cat in paginatedCategories" :key="cat.id" class="hover:bg-gray-50 transition-colors">
            <td class="w-1/12 py-4 px-4 text-center">
              <input 
                v-if="canDeleteCategory && cat.slug !== 'default'"
                type="checkbox" 
                :checked="isCategorySelected(cat.id)"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                @change="toggleCategorySelection(cat.id)"
              />
            </td>
            <td class="w-1/12 py-4 pr-6 whitespace-nowrap text-sm text-gray-900 text-right">
              <img 
                :src="getCategoryThumbnail(cat)" 
                alt="img" 
                class="object-cover rounded-md shadow-sm border cursor-pointer w-12 h-12 md:w-14 md:h-14 ms-2"
                @error="onImgError($event)"
                @click="openModal(cat.id)"
              />
              <ImagePreviewModal
                v-if="isModalOpen(cat.id)"
                :model-value="isModalOpen(cat.id)"
                :image="{ url: getCategoryOriginalImage(cat), name: cat.name }"
                @update:model-value="val => { if (!val) closeModal() }"
              />
            </td>
            <td class="w-1/12 px-4 py-4 whitespace-nowrap text-sm text-gray-900 text-right">{{ cat.id }}</td>
            <td class="w-2/12 px-4 py-4 whitespace-nowrap text-sm font-medium text-gray-900 text-right">
              <a 
                :href="categoryLink(cat)" 
                target="_blank" 
                class="text-blue-600 hover:text-blue-800 hover:underline cursor-pointer transition-colors"
                :title="`مشاهده دسته‌بندی ${cat.name}`"
              >
                {{ cat.name }}
              </a>
            </td>
            <td class="w-1/12 px-4 py-4 whitespace-nowrap text-sm text-gray-700 text-right">
                  <span class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium" :class="typeLabel(cat) === 'اصلی' ? 'bg-purple-100 text-purple-800' : 'bg-green-100 text-green-800'">
              {{ typeLabel(cat) }}
                  </span>
            </td>
            <td class="w-1/12 px-4 py-4 whitespace-nowrap text-sm text-gray-900 text-right">
              <span
                v-if="cat.product_count && cat.product_count>0"
                class="inline-flex items-center px-2 py-1 rounded-md bg-blue-100 text-blue-800 cursor-pointer hover:bg-blue-200 transition-colors text-xs font-medium"
                    @click="openProductsModal(cat)"
              >{{ cat.product_count }}</span>
                  <span v-else class="text-gray-400 text-xs">0</span>
            </td>
            <td class="w-1/12 px-4 py-4 whitespace-nowrap text-sm text-right">
                  <div class="flex items-center">
                    <div v-if="cat.published === true" class="w-2 h-2 bg-green-500 rounded-full ml-2"></div>
                    <div v-else class="w-2 h-2 bg-red-500 rounded-full ml-2"></div>
                    <span v-if="cat.published === true" class="text-green-700 text-xs font-medium">منتشر شده</span>
                    <span v-else class="text-red-600 text-xs font-medium">منتشر نشده</span>
                  </div>
            </td>
            <td class="w-1/12 px-4 py-4 whitespace-nowrap text-sm text-gray-500 text-right">
                  <span v-if="editingOrderId !== cat.id" class="cursor-pointer hover:underline bg-gray-100 px-2 py-1 rounded text-xs" @click="startEditOrder(cat)">{{ cat.order || 0 }}</span>
              <input
                v-else
                v-model="editingOrderValue"
                type="number"
                min="1"
                :max="categories.length"
                    class="w-16 border rounded px-1 py-0.5 text-center text-xs"
                dir="ltr"
                autofocus
                @blur="finishEditOrder(cat)"
                @keyup="handleOrderInputKey($event, cat)"
              />
            </td>
                <td class="w-3/12 px-4 py-4 whitespace-nowrap text-sm text-gray-700 text-right">
                  <a 
                    v-if="cat.parent_name && cat.parent_name !== '-' && cat.parent_slug" 
                    :href="`/product-category/${cat.parent_slug}`"
                    target="_blank"
                    class="inline-flex items-center px-2 py-1 rounded-md bg-gray-100 text-gray-800 hover:bg-gray-200 hover:text-gray-900 cursor-pointer transition-colors text-xs"
                    :title="`مشاهده دسته‌بندی پدر ${cat.parent_name}`"
                  >
                    {{ cat.parent_name }}
                  </a>
                  <span v-else-if="cat.parent_name && cat.parent_name !== '-'" class="inline-flex items-center px-2 py-1 rounded-md bg-gray-100 text-gray-800 text-xs">{{ cat.parent_name }}</span>
                  <span v-else class="text-gray-400 text-xs">-</span>
                </td>
            <td class="w-1/12 px-4 py-4 text-center">
              <div class="flex items-center justify-center gap-2">
                    <NuxtLink v-if="hasPermission('categories_manage')" :to="`/admin/product-management/product-categories/${cat.id}/edit`" class="inline-flex items-center px-3 py-1 rounded-md text-xs font-semibold text-blue-600 bg-blue-50 hover:bg-blue-100 transition-colors">
                  ویرایش
                </NuxtLink>
                    <button 
  v-if="canDeleteCategory && cat.slug !== 'default'"
  class="inline-flex items-center px-3 py-1 rounded-md bg-red-50 text-red-600 hover:bg-red-100 transition-colors text-xs font-semibold"
  @click="deleteCategory(cat)"
>
  حذف
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
          :total="displayedCategories.length"
          :per-page="itemsPerPage"
          @page-changed="goToPage"
          @per-page-changed="val => { itemsPerPage = val; currentPage = 1 }"
        />
      </div>
    </div>

    <!-- Choose Type Modal -->
    <div v-if="showTypeModal" class="fixed inset-0 z-50 flex items-center justify-center bg-black/50 backdrop-blur-sm">
      <div class="bg-white rounded-xl shadow-2xl w-full max-w-md sm:max-w-lg md:max-w-xl p-6 text-center space-y-6 border border-gray-200">
        <div class="flex items-center justify-center w-16 h-16 mx-auto bg-gradient-to-r from-blue-100 to-indigo-100 rounded-full">
          <svg class="w-8 h-8 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
          </svg>
        </div>
        
        <div>
          <h2 class="text-lg font-bold text-gray-900 mb-2">ایجاد دسته‌بندی جدید</h2>
          <p class="text-sm text-gray-600">نوع دسته‌بندی جدید را انتخاب کنید</p>
        </div>
        
        <div class="space-y-3">
        <!-- Main category button -->
        <button
          class="w-full flex items-center justify-center py-3 px-4 rounded-lg text-sm font-semibold text-white bg-gradient-to-r from-blue-500 to-blue-700 shadow-md hover:from-blue-600 hover:to-blue-800 transition-all duration-200 hover:shadow-lg transform hover:scale-105"
            @click="createMain"
        >
            <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7v10c0 2.21 3.582 4 8 4s8-1.79 8-4V7M4 7c0 2.21 3.582 4 8 4s8-1.79 8-4M4 7c0-2.21 3.582-4 8-4s8 1.79 8 4"></path>
            </svg>
          دسته‌بندی اصلی
        </button>

        <!-- Sub category button -->
        <button
          class="w-full flex items-center justify-center py-3 px-4 rounded-lg text-sm font-semibold text-white bg-gradient-to-r from-green-500 to-emerald-600 shadow-md hover:from-green-600 hover:to-emerald-700 transition-all duration-200 hover:shadow-lg transform hover:scale-105"
            @click="createSub"
        >
            <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
            </svg>
          دسته‌بندی فرعی
        </button>

        <!-- Cancel button -->
        <button
          class="w-full py-2 px-4 rounded-lg text-sm font-semibold text-gray-700 bg-gray-100 hover:bg-gray-200 transition-colors border border-gray-300"
            @click="showTypeModal = false"
        >
          انصراف
        </button>
        </div>
      </div>
    </div>

    <!-- Products Modal -->
    <div v-if="showProductsModal" class="fixed inset-0 z-50 flex items-center justify-center bg-black/50 backdrop-blur-sm">
      <div class="bg-white rounded-xl shadow-2xl w-11/12 md:w-2/3 lg:w-1/2 max-h-[80vh] flex flex-col border border-gray-200">
        <div class="flex justify-between items-center px-6 py-4 border-b border-gray-200 bg-gradient-to-r from-blue-50 to-indigo-50">
          <div class="flex items-center">
            <div class="w-8 h-8 bg-blue-500 rounded-md flex items-center justify-center ml-3">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
              </svg>
            </div>
            <h2 class="text-lg font-semibold text-gray-900">محصولات «{{ selectedCategory?.name }}»</h2>
          </div>
          <button class="w-8 h-8 rounded-lg bg-gray-100 hover:bg-gray-200 text-gray-500 hover:text-gray-700 transition-colors flex items-center justify-center" @click="showProductsModal=false">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>
        <div class="overflow-y-auto p-6 flex-1">
          <table v-if="modalProducts.length" class="min-w-full bg-white border border-gray-200 rounded-lg overflow-hidden">
            <thead>
              <tr class="bg-gray-50 text-right text-xs">
                <th class="py-3 px-4 font-medium text-gray-700">ID</th>
                <th class="py-3 px-4 font-medium text-gray-700">نام</th>
                <th class="py-3 px-4 font-medium text-gray-700">قیمت</th>
                <th class="py-3 px-4 font-medium text-gray-700">منتشر شده</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-gray-200">
              <tr v-for="p in modalProducts" :key="p.id" class="hover:bg-gray-50 transition-colors">
                <td class="py-3 px-4 text-sm text-gray-900">{{ p.id }}</td>
                <td class="py-3 px-4 text-sm text-gray-900 font-medium">{{ p.name }}</td>
                <td class="py-3 px-4 text-sm text-gray-600">{{ p.price || '-' }}</td>
                <td class="py-3 px-4 text-sm">
                  <span :class="isProductPublished(p) ? 'inline-flex items-center px-2 py-1 rounded-full text-xs font-medium bg-green-100 text-green-800' : 'inline-flex items-center px-2 py-1 rounded-full text-xs font-medium bg-red-100 text-red-800'">
                    {{ isProductPublished(p) ? 'منتشر شده' : 'منتشر نشده' }}
                  </span>
                </td>
              </tr>
            </tbody>
          </table>
          <div v-else class="text-center py-12">
            <svg class="w-12 h-12 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4"></path>
            </svg>
            <p class="text-gray-500">محصولی یافت نشد.</p>
          </div>
        </div>
        <div class="px-6 py-4 border-t border-gray-200 bg-gray-50 text-right">
          <button class="px-4 py-2 rounded-lg bg-gray-600 hover:bg-gray-700 text-white font-medium transition-colors" @click="showProductsModal=false">بستن</button>
        </div>
      </div>
    </div>
  </ClientOnly>
  
  <!-- Delete Confirm Modal -->
  <DeleteConfirmModal 
    ref="deleteModalRef"
    single-delete-title="تأیید حذف دسته‌بندی"
    :single-delete-message="deleteMessage"
    @single-delete="handleDeleteConfirm"
    @close-single="handleDeleteCancel"
  />
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
declare const navigateTo: (to: string) => Promise<void>
</script>

<script setup lang="ts">
import { computed, onActivated, onMounted, ref, watch } from 'vue'
import Pagination from '~/components/admin/common/Pagination.vue'
import DeleteConfirmModal from '~/components/common/DeleteConfirmModal.vue'
import { useAuth } from '~/composables/useAuth'
import { useNotifier } from '~/composables/useNotifier'

interface Category {
  id?: number
  name?: string
  slug?: string
  parent_id?: number
  parent_name?: string
  parent_slug?: string
  image_url?: string
  banner_url?: string
  published?: boolean
  order?: number
  product_count?: number
  [key: string]: unknown
}

interface Product {
  id?: number
  name?: string
  price?: number
  published?: boolean
  status?: string
  category_id?: number
  category?: {
    id?: number
    [key: string]: unknown
  }
  [key: string]: unknown
}

interface DeleteModalRef {
  openDeleteConfirm?: (id: number | string) => void
}

const { hasPermission } = useAuth()

const canDeleteCategory = computed(() => hasPermission('categories_manage'))

const deleteModalRef = ref<DeleteModalRef | null>(null)
const deleteMessage = ref('')
const categoryToDelete = ref<Category | null>(null)

const categories = ref<Category[]>([])
const editingOrderId = ref<number | null>(null)
const editingOrderValue = ref('')
// Modal visibility for new category type chooser
const showTypeModal = ref(false)

// Row selection for bulk operations
const selectedCategories = ref<Set<number>>(new Set())
const isAllSelected = ref(false)

// Search & filter
const searchTerm = ref('')
const filterOption = ref<'all'|'main'|'sub'|'published'|'unpublished'>('all')

// Pagination
const currentPage = ref(1)
const itemsPerPage = ref(10)

// Fetch all categories, including unpublished ones (all=1)
const fetchCategories = async () => {
  try {
    interface ApiResponse {
      data?: unknown[]
    }
    const response = await $fetch<unknown[] | ApiResponse>('/api/product-categories?all=1')
    const raw = Array.isArray(response) ? response as Category[] : ((response as ApiResponse)?.data || []) as Category[];
    // Add parent_name & parent_slug for display/link
    raw.forEach((cat: Category) => {
      if (cat.parent_id) {
        const parent = raw.find((c: Category) => c.id === cat.parent_id);
        cat.parent_name = parent ? (parent.name || '-') : '-';
        cat.parent_slug = parent ? (parent.slug || '') : '';
      } else {
        cat.parent_name = '-';
        cat.parent_slug = '';
      }
    });
    categories.value = raw;
  } catch (error) {
    console.error('خطا در دریافت دسته‌بندی‌ها:', error);
    categories.value = [];
  }
}

onMounted(() => {
  fetchCategories()
})

// Refresh list whenever returning to this page (e.g., after navigating away)
onActivated(() => {
  fetchCategories()
})

// Stats computeds
const mainCategories = computed(() => categories.value.filter(c => !String(c.slug || '').includes('/')))
const subCategories = computed(() => categories.value.filter(c => String(c.slug || '').includes('/')))
const publishedCategories = computed(() => categories.value.filter(c => c.published === true))

function startEditOrder(cat: Category) {
  editingOrderId.value = cat.id
  editingOrderValue.value = cat.order ? String(cat.order) : ''
}

async function finishEditOrder(cat: Category) {
  const newOrder = parseInt(editingOrderValue.value)
  if (!isNaN(newOrder) && newOrder >= 1 && newOrder <= categories.value.length && newOrder !== cat.order) {
    // Remove the category from its current position
    const oldIndex = categories.value.findIndex(c => c.id === cat.id)
    const moved = categories.value.splice(oldIndex, 1)[0]
    // Insert at new position (1-based to 0-based)
    categories.value.splice(newOrder - 1, 0, moved)
    // Update order fields
    categories.value.forEach((c, idx) => { c.order = idx + 1 })
    // ارسال ترتیب جدید به سرور
    try {
      await $fetch('/api/product-categories/reorder', {
        method: 'POST',
        body: categories.value.map(c => ({ id: c.id, order: c.order }))
      })
      // alert('ترتیب جدید ذخیره شد')
    } catch (e) {
      alert('خطا در ذخیره ترتیب جدید!')
    }
  }
  editingOrderId.value = null
  editingOrderValue.value = ''
}

function handleOrderInputKey(e: KeyboardEvent, cat: Category) {
  if (e.key === 'Enter') {
    finishEditOrder(cat)
  } else if (e.key === 'Escape') {
    editingOrderId.value = null
    editingOrderValue.value = ''
  }
}

function createMain() {
  showTypeModal.value = false
      navigateTo('/admin/product-management/product-categories/new')
}

function createSub() {
  showTypeModal.value = false
  navigateTo('/admin/category-brand-pages/new')
}

// Delete category
function deleteCategory(cat: Category) {
  if (!canDeleteCategory.value) {
    console.error('شما مجوز حذف دسته‌بندی را ندارید')
    return
  }
  
  categoryToDelete.value = cat
  let msg = `آیا از حذف «${cat.name || ''}» مطمئن هستید؟`
  if (cat.product_count && cat.product_count > 0) {
    msg = `هشدار: این دسته‌بندی دارای ${cat.product_count} محصول است. در صورت حذف، محصولات به دسته‌بندی پیش‌فرض منتقل خواهند شد.\n\n` + msg
  }
  deleteMessage.value = msg
  if (cat.id !== undefined) {
    deleteModalRef.value?.openDeleteConfirm(cat.id)
  }
}

// Handle delete confirmation
async function handleDeleteConfirm(id: number | string) {
  if (id === 'bulk') {
    await handleBulkDeleteConfirm()
    return
  }
  
  if (!categoryToDelete.value || !canDeleteCategory.value) {
    console.error('شما مجوز حذف دسته‌بندی را ندارید')
    return
  }
  
  try {
    // Hard delete via maintenance endpoint
    await $fetch('/api/dev/hard-delete', {
      method: 'POST' as const,
      body: { table: 'categories', id: categoryToDelete.value.id }
    })
    // Remove locally
    categories.value = categories.value.filter(c => c.id !== categoryToDelete.value.id)
    const notifier = useNotifier()
    notifier.success('دسته‌بندی با موفقیت حذف شد')
  } catch (e) {
    console.error('خطا در حذف دسته‌بندی:', e)
    const notifier = useNotifier()
    notifier.error('خطا در حذف دسته‌بندی')
  }
}

// Handle delete cancellation
function handleDeleteCancel() {
  categoryToDelete.value = null
  deleteMessage.value = ''
}


// Apply button (no-op because computed is reactive)
function applySearch() {
  /* intentionally empty: computed property reacts automatically */
}

// Watch for search/filter changes to reset pagination
watch([searchTerm, filterOption], () => {
  currentPage.value = 1
})

// Watch for items per page changes
watch(itemsPerPage, () => {
  currentPage.value = 1
})

// Computed list based on search/filter
const displayedCategories = computed(() => {
  let list = [...categories.value]

  switch (filterOption.value) {
    case 'main':
      list = list.filter(c => !c.parent_id)
      break
    case 'sub':
      list = list.filter(c => !!c.parent_id)
      break
    case 'published':
      list = list.filter(c => c.published === true)
      break
    case 'unpublished':
      list = list.filter(c => c.published === false)
      break
  }

  if (searchTerm.value.trim()) {
    const term = searchTerm.value.trim().toLowerCase()
    list = list.filter(c => String(c.name).toLowerCase().includes(term))
  }

  return list
})

// Pagination computed properties
const totalPages = computed(() => {
  return Math.ceil((displayedCategories.value?.length || 0) / itemsPerPage.value)
})

const paginatedCategories = computed(() => {
  if (!displayedCategories.value || !Array.isArray(displayedCategories.value)) {
    return []
  }
  
  const start = (currentPage.value - 1) * itemsPerPage.value
  const end = start + itemsPerPage.value
  return displayedCategories.value.slice(start, end)
})

// Pagination methods
const goToPage = (page: number) => {
  if (page >= 1 && page <= totalPages.value) {
    currentPage.value = page
  }
}

function onImgError(e: Event){
  const target = e.target as HTMLImageElement
  if (target) {
    target.src = '/statics/images/default-image_100.png'
  }
}

const showProductsModal = ref(false)
const selectedCategory = ref<Category | null>(null)
const modalProducts = ref<Product[]>([])

async function openProductsModal(cat: Category){
  selectedCategory.value = cat
  showProductsModal.value = true
  modalProducts.value = []
  try{
    interface ProductsResponse {
      data?: Product[]
    }
    const response = await $fetch<Product[] | ProductsResponse>('/api/products')
    const all = Array.isArray(response) ? response : (response as ProductsResponse)?.data || []
    modalProducts.value = all.filter((p: Product) => {
      if (p.category_id && String(p.category_id) === String(cat.id)) return true
      if (p.category && String(p.category.id) === String(cat.id)) return true
      return false
    }).slice(0,5)
  }catch(e){
    modalProducts.value = []
  }
}

function isProductPublished(p: Product){
  if (typeof p.published === 'boolean') return p.published
  const status = (p.status || '').toString().toLowerCase()
  return status === 'active' || status === 'published' || status === 'enabled'
}

const typeLabel = (cat: Category):string => {
  // همه دسته‌بندی‌ها "اصلی" هستند
  // "فرعی" فقط برای CategoryBrandPage است (ترکیب برند + دسته‌بندی)
  // که در slug با اسلش مشخص می‌شود
  return String(cat.slug || '').includes('/') ? 'فرعی' : 'اصلی'
}

function categoryLink(cat: Category){
  return cat.parent_slug ? `/product-category/${cat.parent_slug}/${cat.slug}` : `/product-category/${cat.slug}`
}

const modalCategoryId = ref<number|null>(null)

function getCategoryThumbnail(cat: Category) {
  // Try to get *_category thumbnail, fallback to original or placeholder
  if (cat.image_url) {
    const url = cat.image_url.toString()
    const dotIdx = url.lastIndexOf('.')
    if (dotIdx !== -1) {
      return url.slice(0, dotIdx) + '_category' + url.slice(dotIdx)
    }
    return url + '_category'
  }
  return '/statics/images/default-image_100.png'
}

function getCategoryOriginalImage(cat: Category) {
  return cat.image_url || '/statics/images/default-image_100.png'
}

function openModal(categoryId: number) {
  modalCategoryId.value = categoryId
}

function closeModal() {
  modalCategoryId.value = null
}

function isModalOpen(categoryId: number) {
  return modalCategoryId.value === categoryId
}

// Row selection functions
function toggleCategorySelection(categoryId: number) {
  if (selectedCategories.value.has(categoryId)) {
    selectedCategories.value.delete(categoryId)
  } else {
    selectedCategories.value.add(categoryId)
  }
  updateSelectAllState()
}

function toggleSelectAll() {
  if (isAllSelected.value) {
    selectedCategories.value.clear()
  } else {
    paginatedCategories.value.forEach(cat => {
      if (canDeleteCategory.value && cat.slug !== 'default') {
        selectedCategories.value.add(cat.id)
      }
    })
  }
  isAllSelected.value = !isAllSelected.value
}

function updateSelectAllState() {
  const selectableCategories = paginatedCategories.value.filter(cat => 
    canDeleteCategory.value && cat.slug !== 'default'
  )
  isAllSelected.value = selectableCategories.length > 0 && 
    selectableCategories.every(cat => selectedCategories.value.has(cat.id))
}

function isCategorySelected(categoryId: number) {
  return selectedCategories.value.has(categoryId)
}

function canDeleteSelected() {
  return selectedCategories.value.size > 0 && canDeleteCategory.value
}

// Bulk delete function
async function bulkDeleteCategories() {
  if (selectedCategories.value.size === 0 || !canDeleteCategory.value) return
  
  const selectedIds = Array.from(selectedCategories.value)
  const selectedCats = categories.value.filter(cat => selectedIds.includes(cat.id))
  
  let message = `آیا از حذف ${selectedIds.length} دسته‌بندی انتخاب شده مطمئن هستید؟\n\n`
  message += `دسته‌بندی‌های انتخاب شده:\n`
  selectedCats.forEach(cat => {
    message += `• ${cat.name}`
    if (cat.product_count && cat.product_count > 0) {
      message += ` (${cat.product_count} محصول)`
    }
    message += '\n'
  })
  
  deleteMessage.value = message
  deleteModalRef.value?.openDeleteConfirm('bulk')
}

// Handle bulk delete confirmation
async function handleBulkDeleteConfirm() {
  if (!canDeleteCategory.value) {
    console.error('شما مجوز حذف دسته‌بندی را ندارید')
    return
  }
  
  const selectedIds = Array.from(selectedCategories.value)
  
  try {
    for (const id of selectedIds) {
      await $fetch('/api/dev/hard-delete', {
        method: 'POST' as const,
        body: { table: 'categories', id }
      })
    }
    
    // Remove locally
    categories.value = categories.value.filter(cat => !selectedIds.includes(cat.id))
    selectedCategories.value.clear()
    isAllSelected.value = false
    
    const notifier = useNotifier()
    notifier.success(`${selectedIds.length} دسته‌بندی با موفقیت حذف شد`)
  } catch (e) {
    console.error('خطا در حذف گروهی دسته‌بندی‌ها:', e)
    const notifier = useNotifier()
    notifier.error('خطا در حذف گروهی دسته‌بندی‌ها')
  }
}

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})


// Expose helpers to template
defineExpose({ getCategoryThumbnail, getCategoryOriginalImage, openModal, closeModal, isModalOpen, modalCategoryId })
</script>

<style scoped>
/* Add any specific styles here */
</style> 

