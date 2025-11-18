<template>
  <client-only>
  <div class="min-h-screen">
    <!-- Header -->
    <div class="bg-white shadow-sm border-b border-gray-200 sticky top-0 z-30">
      <div class="px-3 py-2">
        <div class="flex flex-col lg:flex-row lg:items-center lg:justify-between gap-3">
          <div>
            <h1 class="text-lg font-bold text-gray-900">کتابخانه رسانه</h1>
            <p class="text-xs text-gray-500 mt-0.5">مدیریت فایل‌ها، تصاویر و ویدیوها</p>
          </div>
          
          <div class="flex flex-col sm:flex-row items-stretch sm:items-center gap-3">
            <!-- Search -->
            <div class="relative">
              <input
                v-model="searchQuery"
                type="text"
                placeholder="جستجو در فایل‌ها..."
                class="w-full sm:w-64 pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
              <svg class="w-5 h-5 text-gray-400 absolute left-3 top-2.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
              </svg>
            </div>
            
            <!-- Toggle Compression Tools Button -->
            <TemplateButton
              @click="showCompressionTools = !showCompressionTools"
              :bgGradient="showCompressionTools ? 'bg-amber-100' : 'bg-teal-100'"
              :textColor="showCompressionTools ? 'text-amber-800' : 'text-teal-800'"
              :hoverClass="showCompressionTools ? 'hover:bg-amber-200' : 'hover:bg-teal-200'"
              size="medium"
              title="ابزار فشرده‌سازی"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 12h16M12 4v16" />
              </svg>
              <span class="hidden sm:inline">فشرده‌سازی</span>
            </TemplateButton>

            <!-- Upload Button -->
            <NuxtLink 
              to="/admin/media-management/upload"
            >
              <TemplateButton
                bgGradient="bg-gradient-to-r from-green-400 to-emerald-600"
                textColor="text-white"
                hoverClass="hover:from-green-500 hover:to-emerald-700 hover:shadow-lg hover:scale-105"
                focusClass="focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
                size="medium"
              >
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
                </svg>
                <span class="hidden sm:inline">افزودن جدید</span>
                <span class="sm:hidden">افزودن</span>
              </TemplateButton>
            </NuxtLink>
          </div>
        </div>
      </div>
    </div>

    <!-- Main Content -->
    <div class="p-3">
      <!-- Stats Cards -->
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6 mb-6">
        <TemplateCard
          title="کل فایل‌ها"
          :value="totalFiles"
          variant="blue"
          @click="filterByType('all')"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
            </svg>
          </template>
        </TemplateCard>
        <TemplateCard
          title="تصاویر"
          :value="imageCount"
          variant="purple"
          @click="filterByType('images')"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
            </svg>
          </template>
        </TemplateCard>
        <TemplateCard
          title="ویدیوها"
          :value="videoCount"
          variant="orange"
          @click="filterByType('videos')"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
            </svg>
          </template>
        </TemplateCard>
        <TemplateCard
          title="فضای استفاده شده"
          :value="formatFileSize(filteredSize)"
          variant="green"
        >
          <template #icon>
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7v10c0 2.21 3.582 4 8 4s8-1.79 8-4V7M4 7c0 2.21 3.582 4 8 4s8-1.79 8-4M4 7c0-2.21 3.582-4 8-4s8 1.79 8 4"></path>
            </svg>
          </template>
        </TemplateCard>
      </div>

      <!-- Compression Tools Panel -->
      <div v-if="showCompressionTools" class="bg-gradient-to-r from-amber-50 to-orange-50 border border-amber-200 rounded-lg p-6 mb-4">
        <div class="flex flex-col lg:flex-row lg:items-center lg:justify-between gap-6">
          <div>
            <h3 class="text-lg font-bold text-amber-800 mb-2">ابزار فشرده‌سازی</h3>
            <p class="text-sm text-amber-700">مدیریت فشرده‌سازی تصاویر و ویدیوها</p>
          </div>
          
          <div class="flex flex-col sm:flex-row gap-3">
            <!-- فشرده‌سازی تصاویر -->
            <NuxtLink 
              to="/admin/media-management/image-compression"
            >
              <TemplateButton
                bgGradient="bg-purple-500"
                textColor="text-white"
                hoverClass="hover:bg-purple-600"
                size="medium"
              >
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                </svg>
                <span>فشرده‌سازی تصاویر</span>
              </TemplateButton>
            </NuxtLink>
            
            <!-- فشرده‌سازی ویدیو -->
            <NuxtLink 
              to="/admin/media-management/video-compression"
            >
              <TemplateButton
                bgGradient="bg-orange-500"
                textColor="text-white"
                hoverClass="hover:bg-orange-600"
                size="medium"
              >
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
                </svg>
                <span>فشرده‌سازی ویدیو</span>
              </TemplateButton>
            </NuxtLink>
          </div>
        </div>
      </div>

      <!-- Filter and Bulk Actions -->
      <div class="flex flex-col lg:flex-row lg:items-center lg:justify-between gap-6 mt-3 pt-3 border-t border-gray-100 bg-white p-3 rounded-lg">
        <div class="flex flex-col sm:flex-row items-start sm:items-center gap-3">
          <!-- Add Select All checkbox for grid view -->
          <div v-if="viewMode === 'grid'" class="flex items-center gap-2">
            <input
              type="checkbox"
              @change="toggleSelectAll"
              :checked="selectedItems.length === filteredFiles.length && filteredFiles.length > 0"
              class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500"
            >
            <span class="text-sm text-gray-600">انتخاب همه</span>
          </div>
          
          <div class="flex flex-col sm:flex-row gap-2 w-full sm:w-auto">
            <select v-model="selectedCategory" class="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 text-sm">
              <option value="library">کتابخانه</option>
              <option value="products">محصولات</option>
              <option value="product-categories">دسته‌بندی محصولات</option>
              <option value="brands">برندها</option>
              <option value="banners">بنرهای تبلیغاتی</option>
              <option value="customer">مشتریان</option>
            </select>
            <select 
              v-model="selectedFilter"
              class="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 text-sm"
            >
              <option value="all">همه فایل‌ها</option>
              <option value="images">تصاویر</option>
              <option value="videos">ویدیوها</option>
              <option value="documents">اسناد</option>
              <option value="audio">صوتی</option>
            </select>
            <select 
              v-model="sortBy"
              class="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 text-sm"
            >
              <option value="date">تاریخ آپلود</option>
              <option value="name">نام فایل</option>
              <option value="size">اندازه فایل</option>
              <option value="type">نوع فایل</option>
            </select>
          </div>
          
          <div class="flex border border-gray-300 rounded-lg overflow-hidden">
            <TemplateButton
              @click="viewMode = 'grid'"
              :bgGradient="viewMode === 'grid' ? 'bg-blue-500' : 'bg-white'"
              :textColor="viewMode === 'grid' ? 'text-white' : 'text-gray-700'"
              :hoverClass="viewMode === 'grid' ? '' : 'hover:bg-gray-50'"
              size="small"
              :borderColor="'border-0'"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z"></path>
              </svg>
            </TemplateButton>
            <TemplateButton
              @click="viewMode = 'list'"
              :bgGradient="viewMode === 'list' ? 'bg-blue-500' : 'bg-white'"
              :textColor="viewMode === 'list' ? 'text-white' : 'text-gray-700'"
              :hoverClass="viewMode === 'list' ? '' : 'hover:bg-gray-50'"
              size="small"
              :borderColor="'border-0'"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16"></path>
              </svg>
            </TemplateButton>
          </div>
        </div>
        
        <!-- Bulk Actions -->
        <div v-if="selectedItems.length > 0" class="flex flex-col sm:flex-row items-start sm:items-center gap-3">
          <span class="text-sm text-gray-600">{{ selectedItems.length }} مورد انتخاب شده</span>
          <div class="flex flex-wrap gap-2">
            <TemplateButton 
              v-if="canDeleteMedia"
              @click="bulkDelete"
              bgGradient="bg-red-500"
              textColor="text-white"
              hoverClass="hover:bg-red-600"
              size="small"
            >
              حذف
            </TemplateButton>
            <TemplateButton 
              v-if="hasUncompressedSelected"
              @click="bulkCompress"
              bgGradient="bg-green-500"
              textColor="text-white"
              hoverClass="hover:bg-green-600"
              size="small"
            >
              فشرده‌سازی
            </TemplateButton>
            <TemplateButton 
              @click="selectedItems = []"
              bgGradient="bg-gray-500"
              textColor="text-white"
              hoverClass="hover:bg-gray-600"
              size="small"
            >
              لغو انتخاب
            </TemplateButton>
          </div>
        </div>
      </div>

      <!-- Media Grid/List -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
        <!-- Grid View -->
        <div v-if="viewMode === 'grid'" class="p-3">
          <div v-if="filteredFiles.length === 0" class="text-center py-8">
            <svg class="w-24 h-24 text-gray-300 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21h10a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v14a2 2 0 002 2z"></path>
            </svg>
            <h3 class="text-lg font-medium text-gray-900 mb-2">هنوز فایلی آپلود نشده</h3>
            <p class="text-gray-600 mb-6">اولین فایل خود را آپلود کنید</p>
            <NuxtLink 
              to="/admin/media-management/upload"
            >
              <TemplateButton
                bgGradient="bg-blue-500"
                textColor="text-white"
                hoverClass="hover:bg-blue-600"
                size="large"
              >
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
                </svg>
                آپلود فایل
              </TemplateButton>
            </NuxtLink>
          </div>

          <div v-else>
            <ImageSelector
              :images="visibleFiles"
              :selected-images="selectedItems"
              :format-file-size="formatFileSize"
              @preview="viewFile"
              @select="handleSelect"
            />
          </div>
        </div>

        <!-- List View -->
        <div v-else class="overflow-x-auto">
          <table class="w-full min-w-full">
            <thead class="bg-gray-50 border-b border-gray-200">
              <tr class="text-right">
                <th class="px-3 lg:px-6 py-3 text-right">
                  <input
                    type="checkbox"
                    @change="toggleSelectAll"
                    :checked="selectedItems.length === filteredFiles.length && filteredFiles.length > 0"
                    class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500"
                  >
                </th>
                <th class="px-3 lg:px-6 py-3 text-xs font-medium text-gray-500 uppercase tracking-wider text-right">فایل</th>
                <th class="hidden md:table-cell px-3 lg:px-6 py-3 text-xs font-medium text-gray-500 uppercase tracking-wider text-right">نوع</th>
                <th class="hidden sm:table-cell px-3 lg:px-6 py-3 text-xs font-medium text-gray-500 uppercase tracking-wider text-right">اندازه</th>
                <th class="hidden lg:table-cell px-3 lg:px-6 py-3 text-xs font-medium text-gray-500 uppercase tracking-wider text-right">تاریخ آپلود</th>
                <th class="px-3 lg:px-6 py-3 text-xs font-medium text-gray-500 uppercase tracking-wider text-right">عملیات</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="file in visibleFiles" :key="file.id" class="hover:bg-gray-50">
                <td class="px-3 lg:px-6 py-4">
                  <input
                    type="checkbox"
                    :checked="selectedItems.includes(file.id)"
                    @change="toggleSelection(file.id)"
                    class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500"
                  >
                </td>
                <td class="px-3 lg:px-6 py-4">
                  <div class="flex items-center">
                    <div class="flex-shrink-0 h-12 w-12">
                      <img 
                        v-if="file.type === 'image'" 
                        :src="file.thumbnail" 
                        v-bind="{
                          ...(file.alt ? { alt: file.alt } : {}),
                          ...(file.title ? { title: file.title } : {}),
                          ...(file.caption ? { 'data-caption': file.caption } : {}),
                          ...(file.description ? { 'data-description': file.description } : {})
                        }"
                        class="h-12 w-12 rounded-lg object-cover"
                      >
                      <div v-else class="h-12 w-12 bg-gray-100 rounded-lg flex items-center justify-center">
                        <svg class="w-6 h-6 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                        </svg>
                      </div>
                    </div>
                    <div class="mr-4 min-w-0 flex-1">
                      <div class="text-sm font-medium text-gray-900 truncate">{{ file.name }}</div>
                      <div class="text-xs text-gray-500 mt-1">
                        <span class="md:hidden">{{ file.extension.toUpperCase() }} • {{ formatFileSize(file.compressed_size ?? file.size) }}</span>
                        <span class="hidden md:inline lg:hidden">{{ formatFileSize(file.compressed_size ?? file.size) }}</span>
                      </div>
                    </div>
                  </div>
                </td>
                <td class="hidden md:table-cell px-3 lg:px-6 py-4 text-sm text-gray-900">
                  <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-gray-100 text-gray-800">
                    {{ file.extension.toUpperCase() }}
                  </span>
                </td>
                <td class="hidden sm:table-cell px-3 lg:px-6 py-4 text-sm text-gray-900">{{ formatFileSize(file.compressed_size ?? file.size) }}</td>
                <td class="hidden lg:table-cell px-3 lg:px-6 py-4 text-sm text-gray-900">{{ formatDate(file.uploadDate) }}</td>
                <td class="px-3 lg:px-6 py-4 text-sm font-medium">
                  <div class="flex items-center gap-1 lg:gap-2">
                    <button 
                      @click="viewFile(file)"
                      class="text-blue-600 hover:text-blue-900 transition-colors p-1 lg:p-2"
                      title="مشاهده"
                    >
                      <svg class="w-4 h-4 lg:w-5 lg:h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                      </svg>
                    </button>

                    <!-- Quick compress / revert -->
                    <template v-if="showCompressionTools && file.type === 'image'">
                      <button
                        v-if="!(file.compressed_size && file.size && file.compressed_size < file.size)"
                        @click="quickCompress(file.id)"
                        class="text-emerald-600 hover:text-emerald-900 transition-colors p-1 lg:p-2"
                        title="فشرده سازی سریع"
                      >
                        <svg class="w-4 h-4 lg:w-5 lg:h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
                        </svg>
                      </button>
                      <button
                        v-else
                        @click="handleRevert(file.id)"
                        class="text-rose-600 hover:text-rose-900 transition-colors p-1 lg:p-2"
                        title="بازگردانی"
                      >
                        <svg class="w-4 h-4 lg:w-5 lg:h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
                        </svg>
                      </button>
                    </template>

                    <button 
                      v-if="canDeleteMedia"
                      @click="deleteFile(file)"
                      class="text-red-600 hover:text-red-900 transition-colors p-1 lg:p-2"
                      title="حذف"
                    >
                      <svg class="w-4 h-4 lg:w-5 lg:h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                      </svg>
                    </button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- بارگذاری بیشتر -->
      <div v-if="itemsToShow < filteredFiles.length" class="flex justify-center my-6">
        <button @click="itemsToShow += 36" class="bg-red-500 text-white px-6 py-2 rounded-lg hover:bg-red-600 transition-colors">
          بارگذاری بیشتر
        </button>
      </div>
    </div>

    <!-- File Viewer Modal -->
    <MediaPreviewModal
      v-model:modelValue="previewModalVisible"
      v-if="selectedFile"
      :file="selectedFile"
      @close="closeViewer"
      @save="handleSaveDetail"
      @delete="handleDeleteDetail"
      @compress="handleCompress"
      @revert="handleRevert"
    />
  </div>
  </client-only>
  <!-- Global Delete Confirm Modal (standard project component) -->
  <DeleteConfirmModal
    ref="deleteModalRef"
    single-delete-title="تایید حذف"
    single-delete-message="آیا از حذف این مورد اطمینان دارید؟"
    bulk-delete-title="حذف گروهی"
    :selected-count="selectedItems.length"
    @single-delete="handleSingleDelete"
    @bulk-delete="handleBulkDelete"
  />

  <!-- Global Confirm Dialog for this page -->
  <ConfirmDialog ref="confirmDialogRef" />
</template>

<script setup lang="ts">
import { computed, onMounted, provide, reactive, ref, watch } from 'vue'
import ConfirmDialog from '~/components/common/ConfirmDialog.vue'
import DeleteConfirmModal from '~/components/common/DeleteConfirmModal.vue'
import TemplateButton from '~/components/common/TemplateButton.vue'
import TemplateCard from '~/components/common/TemplateCard.vue'
import ImageSelector from '~/components/media/ImageSelector.vue'
import MediaPreviewModal from '~/components/media/MediaPreviewModal.vue'
import { useDeleteConfirmModal } from '~/composables/useDeleteConfirmModal'
import { useToast } from '~/composables/useToast'

// استفاده از useAuth برای چک کردن پرمیژن‌ها
// @ts-ignore
const { hasPermission } = useAuth()

interface MediaFile {
  id: number
  name: string
  url: string
  thumbnail: string
  type: 'image' | 'video' | 'document' | 'audio'
  extension: string
  size: number
  compressed_size?: number | null
  uploadDate: string
  alt?: string
  title?: string
  description?: string
  duration?: string
  width?: number
  height?: number
  caption?: string
  mime_type?: string
  author?: string
  category: string
  uploader_name?: string
  uploader_username?: string
  uploader_role?: string
  uploader_first_name?: string
  uploader_last_name?: string
}

declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
declare const useHead: (head: { title?: string }) => void

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

useHead({
  title: 'کتابخانه رسانه - پنل مدیریت'
})

// Reactive data
const searchQuery = ref('')
const selectedFilter = ref('all')
const sortBy = ref('date')
const viewMode = ref<'grid' | 'list'>('grid')
const selectedItems = ref<number[]>([])
const selectedFile = ref<MediaFile | null>(null)
const currentPage = ref(1)
const perPage = ref(24)
const files = ref<MediaFile[]>([])
const selectedCategory = ref('library')
const itemsToShow = ref(36)
const previewModalVisible = ref(false)
const showCompressionTools = ref(false)
const notifier = useToast()
const { deleteModalRef, openDeleteConfirm, openBulkDeleteConfirm } = useDeleteConfirmModal()

// Provide confirm dialog for this page
const confirmDialogRef = ref<any>(null)
provide('confirmDialogRef', confirmDialogRef)

const editFields = reactive({
  alt: '',
  title: '',
  caption: '',
  description: ''
})

watch(selectedFile, (file) => {
  if (file) {
    editFields.alt = file.alt || ''
    editFields.title = file.title || ''
    editFields.caption = file.caption || ''
    editFields.description = file.description || ''
  }
})

// Add compressionConfig
const compressionConfig = reactive({
  quality: 'medium' as 'low' | 'medium' | 'high' | 'smart' | 'custom',
  customQuality: 75,
  format: 'original' as 'original' | 'webp' | 'jpeg' | 'png'
})

async function loadCompressionSettings() {
  try {
    const res = await fetch('/api/admin/settings')
    if (!res.ok) return
    const arr = await res.json()
    if (Array.isArray(arr)) {
      const map: Record<string,string> = {}
      for (const s of arr) map[s.key || s.Key] = s.value ?? s.Value
      if (map['compression.quality']) compressionConfig.quality = map['compression.quality'] as any
      if (map['compression.custom_quality']) {
        const n = Number(map['compression.custom_quality']); if(!isNaN(n)) compressionConfig.customQuality = n
      }
      if (map['compression.format']) compressionConfig.format = map['compression.format'] as any
    }
  } catch {}
}

onMounted(() => {
  fetchMediaList()
  loadCompressionSettings()
  
  // بررسی اینکه آیا نیاز به refresh داریم (مثلاً بعد از آپلود فایل جدید)
  const shouldRefresh = sessionStorage.getItem('refreshMediaLibrary')
  if (shouldRefresh === 'true') {
    sessionStorage.removeItem('refreshMediaLibrary')
    setTimeout(() => {
      fetchMediaList()
    }, 100)
  }
})

function calcQuality(): number {
  if (compressionConfig.quality === 'smart') return 75 // backend decides
  if (compressionConfig.quality === 'custom') return compressionConfig.customQuality
  const map: Record<string, number> = { high: 90, medium: 75, low: 60 }
  return map[compressionConfig.quality] || 75
}

function calcFormat(): string | undefined {
  return compressionConfig.format === 'original' ? undefined : compressionConfig.format
}

function quickCompress(id: number) {
  if (compressionConfig.quality === 'smart') {
    handleCompress({ id, smart: true, format: calcFormat() })
  } else {
    handleCompress({ id, quality: calcQuality(), format: calcFormat() })
  }
}

// تابع fetchMediaList را به صورت جدا تعریف کن تا هم در onMounted و هم در watcher استفاده شود
const fetchMediaList = async () => {
  try {
    let response
    let json
    
    // اگر بخش products انتخاب شده، هر دو منبع را ترکیب کن
    if (selectedCategory.value === 'products') {
      // دریافت تصاویر محصولات از جدول product_images
      const productImagesResponse = await fetch('/api/products/images')
      let productImages = []
      if (productImagesResponse.ok) {
        const productImagesJson = await productImagesResponse.json()
        productImages = Array.isArray(productImagesJson.data) ? productImagesJson.data : []
      }
      
      // دریافت تصاویر معمولی از جدول media_files
      const mediaResponse = await fetch('/api/media/list?category=products')
      let mediaFiles = []
      if (mediaResponse.ok) {
        const mediaJson = await mediaResponse.json()
        mediaFiles = Array.isArray(mediaJson.data) ? mediaJson.data : []
      }
      
      // ترکیب هر دو لیست
      const allItems = [...productImages, ...mediaFiles]
      
      // دریافت سایز فایل‌ها برای تصاویر محصولات
      const itemsWithSize = await Promise.all(allItems.map(async (item: any) => {
        const isProductImage = 'image_url' in item
        if (isProductImage) {
          try {
            const sizeResponse = await fetch(`/api/media/file-size?url=${encodeURIComponent(item.image_url)}`)
            const sizeData = await sizeResponse.json()
            return { ...item, fileSize: sizeData.size || 0 }
          } catch (error) {
            return { ...item, fileSize: 0 }
          }
        }
        return item
      }))
      
      files.value = itemsWithSize.map((item: any) => {
        // تشخیص نوع آیتم (از product_images یا media_files)
        const isProductImage = 'image_url' in item
        
        const rawPath = isProductImage ? item.image_url : (item.url || item.file_path || '');
        let normPath = (rawPath || '').replace(/\\/g, '/');
        // ترازسازی مسیر برای سروینگ public
        if (!/^https?:\/\//i.test(normPath)) {
          if (!normPath.startsWith('/')) normPath = '/' + normPath;
        }
        let publicPath = normPath;
        // اگر مسیر از قبل شامل uploads است، همان را استفاده کن
        if (!/^https?:\/\//i.test(normPath)) {
          if (/\/uploads\//.test(normPath)) {
            publicPath = normPath;
          } else if (/^\/(products|product-categories|brands|banners|customer|library)\//i.test(normPath)) {
            // مسیرهای استاندارد دایرکتوری‌های ما زیر /uploads/media
            publicPath = '/uploads/media' + normPath;
          } else {
            // سایر مسیرها را زیر دسته انتخاب‌شده قرار بده
            const baseCat = (selectedCategory.value || 'library').replace(/[^a-zA-Z0-9-_]/g, '');
            publicPath = `/uploads/media/${baseCat}/${normPath.replace(/^\//, '')}`;
          }
        }
        
        // MIME/Extension detection
        let ext = ''
        if (!ext) {
          const m = publicPath.match(/\.([a-z0-9]+)(?:\?|#|$)/i)
          if (m) ext = m[1].toLowerCase()
        }
        const isImage = ['jpg','jpeg','png','gif','webp','svg','avif'].includes(ext)
        
        if (isProductImage) {
          // آیتم از جدول product_images
          return {
            id: item.id,
            name: item.product_name || `Product Image ${item.id}`,
            url: publicPath,
            thumbnail: isImage ? publicPath : '',
            type: 'image' as const,
            extension: ext,
            size: item.fileSize || 0, // اندازه فایل از سیستم فایل دریافت شده
            compressed_size: undefined,
            uploadDate: item.created_at || new Date().toISOString(),
            alt: `تصویر محصول ${item.product_name || 'نامشخص'}`,
            title: `تصویر محصول: ${item.product_name || 'نامشخص'}`,
            description: `تصویر محصول: ${item.product_name || 'نامشخص'} (SKU: ${item.product_id})`,
            duration: undefined,
            width: undefined,
            height: undefined,
            caption: `تصویر محصول ${item.product_name || 'نامشخص'}`,
            mime_type: ext === 'webp' ? 'image/webp' : ext === 'jpg' || ext === 'jpeg' ? 'image/jpeg' : ext === 'png' ? 'image/png' : 'image/webp',
            author: 'System Import',
            category: 'products' as const,
            uploader_name: 'System Import',
            uploader_username: 'system',
            uploader_role: 'admin',
          }
        } else {
          // آیتم از جدول media_files
          const mime = (item.mime_type || item.file_type || '').toLowerCase()
          let ext = (mime.split('/').pop() || item.extension || '').toLowerCase()
          if (!ext) {
            const m = publicPath.match(/\.([a-z0-9]+)(?:\?|#|$)/i)
            if (m) ext = m[1].toLowerCase()
          }
          const isImage = (mime.startsWith('image')) || ['jpg','jpeg','png','gif','webp','svg'].includes(ext)
          const isVideo = (mime.startsWith('video')) || ['mp4','webm','mov','avi','mkv','flv','3gp','wmv','m4v','ogv'].includes(ext)
          const isAudio = (mime.startsWith('audio')) || ['mp3','wav','ogg'].includes(ext)
          return {
            id: item.id,
            name: item.original_name || item.file_name,
            url: publicPath,
            thumbnail: isImage ? publicPath : '',
            type: isImage ? 'image' : (isVideo ? 'video' : (isAudio ? 'audio' : 'document')),
            extension: ext,
            size: item.size ?? item.file_size ?? 0,
            compressed_size: item.compressed_size,
            uploadDate: item.created_at || item.uploaded_at || item.uploadDate,
            alt: item.alt_text,
            title: item.title,
            description: item.description,
            duration: item.duration ? item.duration.toString() : undefined,
            width: item.width,
            height: item.height,
            caption: item.caption,
            mime_type: item.mime_type,
            author: item.author,
            category: item.category || 'products',
            uploader_name: item.uploader_name,
            uploader_username: item.uploader_username,
            uploader_role: item.uploader_role,
          }
        }
      }) as MediaFile[]
    } else {
      // برای سایر بخش‌ها از API معمولی استفاده کن
      response = await fetch(`/api/media/list?category=${selectedCategory.value}`)
      if (!response.ok) throw new Error('Failed to fetch media list')
      json = await response.json()
      const data = Array.isArray(json.data) ? json.data : json
      if (!Array.isArray(data)) {
        files.value = []
        return
      }
      files.value = data.map((item: any) => {
        const rawPath = item.url || item.file_path || '';
        let normPath = (rawPath || '').replace(/\\/g, '/');
        // ترازسازی مسیر برای سروینگ public
        if (!/^https?:\/\//i.test(normPath)) {
          if (!normPath.startsWith('/')) normPath = '/' + normPath;
        }
        let publicPath = normPath;
        // اگر مسیر از قبل شامل uploads است، همان را استفاده کن
        if (!/^https?:\/\//i.test(normPath)) {
          if (/\/uploads\//.test(normPath)) {
            publicPath = normPath;
          } else if (/^\/(products|product-categories|brands|banners|customer|library)\//i.test(normPath)) {
            // مسیرهای استاندارد دایرکتوری‌های ما زیر /uploads/media
            publicPath = '/uploads/media' + normPath;
          } else {
            // سایر مسیرها را زیر دسته انتخاب‌شده قرار بده (مثل ویدیوها: /uploads/media/library/videos/..)
            const baseCat = (selectedCategory.value || 'library').replace(/[^a-zA-Z0-9-_]/g, '');
            publicPath = `/uploads/media/${baseCat}/${normPath.replace(/^\//, '')}`;
          }
        }
        // MIME/Extension detection
        const mime = (item.mime_type || item.file_type || '').toLowerCase()
        let ext = (mime.split('/').pop() || item.extension || '').toLowerCase()
        if (!ext) {
          const m = publicPath.match(/\.([a-z0-9]+)(?:\?|#|$)/i)
          if (m) ext = m[1].toLowerCase()
        }
        const isImage = (mime.startsWith('image')) || ['jpg','jpeg','png','gif','webp','svg'].includes(ext)
        const isVideo = (mime.startsWith('video')) || ['mp4','webm','mov','avi','mkv','flv','3gp','wmv','m4v','ogv'].includes(ext)
        const isAudio = (mime.startsWith('audio')) || ['mp3','wav','ogg'].includes(ext)
        return {
          id: item.id,
          name: item.original_name || item.file_name,
          url: publicPath,
          thumbnail: isImage ? publicPath : '',
          type: isImage ? 'image' : (isVideo ? 'video' : (isAudio ? 'audio' : 'document')),
          extension: ext,
          size: item.size ?? item.file_size ?? 0,
          compressed_size: item.compressed_size,
          uploadDate: item.created_at || item.uploaded_at || item.uploadDate,
          alt: item.alt_text,
          title: item.title,
          description: item.description,
          duration: item.duration ? item.duration.toString() : undefined,
          width: item.width,
          height: item.height,
          caption: item.caption,
          mime_type: item.mime_type,
          author: item.author,
          category: item.category || (
            rawPath.includes('/product-categories/') || rawPath.includes('\\product-categories\\')
              ? 'product-categories'
              : rawPath.includes('/brands/') || rawPath.includes('\\brands\\')
              ? 'brands'
              : rawPath.includes('/banners/') || rawPath.includes('\\banners\\')
              ? 'banners'
              : rawPath.includes('/products/') || rawPath.includes('\\products\\')
              ? 'products'
              : rawPath.includes('/customer/') || rawPath.includes('\\customer\\')
              ? 'customer'
              : (selectedCategory.value || 'library')
          ),
          uploader_name: item.uploader_name,
          uploader_username: item.uploader_username,
          uploader_role: item.uploader_role,
        }
      }) as MediaFile[]
    }
  } catch (error: unknown) {
    console.error('Error fetching media list:', error as any)
  }
}

watch(selectedCategory, () => {
  files.value = []
  fetchMediaList()
})

// Computed properties
const filteredFiles = computed(() => {
  let mediaFiles: MediaFile[] = files.value

  // همیشه فیلتر بر اساس دسته انتخاب‌شده
  mediaFiles = mediaFiles.filter(f => f.category === selectedCategory.value)

  // Filter by type
  if (selectedFilter.value !== 'all') {
    if (selectedFilter.value === 'images') {
      mediaFiles = mediaFiles.filter((f: MediaFile) => f.type === 'image')
    } else if (selectedFilter.value === 'videos') {
      mediaFiles = mediaFiles.filter((f: MediaFile) => f.type === 'video')
    } else if (selectedFilter.value === 'documents') {
      mediaFiles = mediaFiles.filter((f: MediaFile) => f.type === 'document')
    } else if (selectedFilter.value === 'audio') {
      mediaFiles = mediaFiles.filter((f: MediaFile) => f.type === 'audio')
    }
  }

  // Filter by search
  if (searchQuery.value) {
    mediaFiles = mediaFiles.filter((f: MediaFile) => 
      f.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      (f.alt && f.alt.toLowerCase().includes(searchQuery.value.toLowerCase()))
    )
  }

  // Sort files
  mediaFiles.sort((a: MediaFile, b: MediaFile) => {
    switch (sortBy.value) {
      case 'name':
        return a.name.localeCompare(b.name)
      case 'size':
        return b.size - a.size
      case 'type':
        return a.extension.localeCompare(b.extension)
      case 'date':
      default:
        return new Date(b.uploadDate).getTime() - new Date(a.uploadDate).getTime()
    }
  })

  return mediaFiles
})

const totalFiles = computed(() => filteredFiles.value.length)
const totalSize = computed(() => filteredFiles.value.reduce((sum: number, file: MediaFile) => sum + (file.compressed_size ?? file.size), 0))
const imageCount = computed(() => filteredFiles.value.filter((f: MediaFile) => f.type === 'image').length)
const videoCount = computed(() => filteredFiles.value.filter((f: MediaFile) => f.type === 'video').length)

const totalPages = computed(() => Math.ceil(filteredFiles.value.length / perPage.value))
const startIndex = computed(() => (currentPage.value - 1) * perPage.value)
const paginatedFiles = computed(() => 
  filteredFiles.value.slice(startIndex.value, startIndex.value + perPage.value)
)

const visibleFiles = computed(() => filteredFiles.value.slice(0, itemsToShow.value))

// Computed برای چک کردن پرمیژن حذف
const canDeleteMedia = computed(() => hasPermission('media_library.delete'))

// Methods
const formatFileSize = (bytes: number): string => {
  if (bytes === 0) return '0 B'
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

const formatDate = (dateString: string): string => {
  return new Date(dateString).toLocaleDateString('fa-IR', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

const selectFile = (file: MediaFile) => {
  selectedFile.value = file
}

const viewFile = (file: MediaFile) => {
  selectedFile.value = file
  previewModalVisible.value = true
}

const closeViewer = () => {
  previewModalVisible.value = false
  selectedFile.value = null
}

const toggleSelection = (id: number) => {
  const idx = selectedItems.value.indexOf(id)
  if (idx === -1) {
    selectedItems.value.push(id)
  } else {
    selectedItems.value.splice(idx, 1)
  }
}

const handleSelect = (id: number, event: MouseEvent) => {
  if (event.ctrlKey || event.metaKey) {
    // multi-select toggle
    const idx = selectedItems.value.indexOf(id)
    if (idx === -1) {
      selectedItems.value.push(id)
    } else {
      selectedItems.value.splice(idx, 1)
    }
  } else {
    // single select
    selectedItems.value = [id]
  }
}

const toggleSelectAll = () => {
  if (selectedItems.value.length === filteredFiles.value.length) {
    selectedItems.value = []
  } else {
    selectedItems.value = filteredFiles.value.map(f => f.id)
  }
}

const editFile = (file: MediaFile) => {
  // Navigate to edit page or open edit modal
  console.log('Edit file:', file)
}

const deleteFile = async (file: MediaFile) => {
  try {
    console.log('Attempting to delete file:', file);
    const response = await fetch(`/api/media/delete?id=${file.id}`, { method: 'DELETE' })
    
    console.log('Delete response status:', response.status);
    const responseText = await response.text();
    console.log('Delete response text:', responseText);
    
    if (!response.ok) {
      throw new Error(`خطا در حذف فایل: ${response.status} ${responseText}`);
    }
    
    // Remove the file from the files array
    const index = files.value.findIndex((f: MediaFile) => f.id === file.id);
    if (index > -1) {
      files.value.splice(index, 1);
    }
    
    // Close viewer if the deleted file was being viewed
    if (selectedFile.value && selectedFile.value.id === file.id) {
      selectedFile.value = null;
    }
    
    // Remove from selected items if present
    const selectedIndex = selectedItems.value.indexOf(file.id);
    if (selectedIndex > -1) {
      selectedItems.value.splice(selectedIndex, 1);
    }
    
    notifier.showSuccess('فایل با موفقیت حذف شد')
  } catch (error: any) {
    console.error('Error deleting file:', error);
    notifier.showError(`خطا در حذف فایل: ${error.message}`)
  }
}

const bulkDelete = async () => {
  openBulkDeleteConfirm()
}

function isAlreadyCompressed(f: MediaFile) {
  const flag = (f as any).compressed === true
  const sizeFlag = f.compressed_size !== null && f.compressed_size !== undefined && f.size && f.compressed_size < f.size
  return flag || sizeFlag
}

const bulkCompress = async () => {
  const targets = selectedItems.value
    .map(id => files.value.find(f => f.id === id))
    .filter((f): f is MediaFile => !!f && f.type === 'image' && !isAlreadyCompressed(f))

  if (targets.length === 0) {
    notifier.showInfo('همه تصاویر انتخاب‌شده قبلاً فشرده شده‌اند')
    return
  }

  for (const f of targets) {
    await handleCompress({ id: f.id, quality: calcQuality(), format: calcFormat() })
  }
  // پس از اتمام، فقط موارد فشرده‌نشده قبلی را از انتخاب خارج کنیم
  selectedItems.value = selectedItems.value.filter(id => !targets.some(t => t.id === id))
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

const copyFileUrl = () => {
  if (selectedFile.value) {
    window.navigator.clipboard.writeText(selectedFile.value.url)
  }
}

const handleSaveDetail = (fields: { id: number, title: string, caption: string, alt: string, description: string }) => {
  if (!fields.id) return;
  const idx = files.value.findIndex(f => f.id === fields.id);
  if (idx === -1) return;
  fetch(`/api/media/${fields.id}`, {
      method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        title: fields.title,
        caption: fields.caption,
        alt_text: fields.alt,
        description: fields.description,
      mime_type: files.value[idx].mime_type || '',
      author: files.value[idx].author || ''
      })
  }).then(res => {
    if (!res.ok) throw new Error('خطا در ذخیره اطلاعات');
    files.value[idx].title = fields.title;
    files.value[idx].caption = fields.caption;
    files.value[idx].alt = fields.alt;
    files.value[idx].description = fields.description;
    if (selectedFile.value && selectedFile.value.id === fields.id) {
    selectedFile.value.title = fields.title;
    selectedFile.value.caption = fields.caption;
    selectedFile.value.alt = fields.alt;
    selectedFile.value.description = fields.description;
    }
  }).catch(() => {
    notifier.showError('خطا در ذخیره اطلاعات')
  });
}

const handleDeleteDetail = (id: number) => {
  openDeleteConfirm(id)
}

async function handleSingleDelete(id: number | string) {
  const file = files.value.find(f => f.id === id)
  if (!file) return
  await deleteFile(file)
  if (selectedFile.value && selectedFile.value.id === id) {
    previewModalVisible.value = false
    selectedFile.value = null
  }
}

async function handleBulkDelete() {
  try {
    for (const fileId of selectedItems.value) {
      const response = await fetch(`/api/media/delete?id=${fileId}`, { method: 'DELETE' })
      if (!response.ok) {
        const errorText = await response.text()
        console.error(`Delete response for file ${fileId}:`, response.status, errorText)
        throw new Error(`خطا در حذف فایل با شناسه ${fileId}: ${response.status} ${errorText}`)
      }
    }
    files.value = files.value.filter((f: MediaFile) => !selectedItems.value.includes(f.id))
    selectedItems.value = []
    if (selectedFile.value && selectedItems.value.includes(selectedFile.value.id)) {
      selectedFile.value = null
    }
  } catch (error: any) {
    console.error('Error in bulk delete:', error)
    notifier.showError('خطا در حذف فایل‌ها')
  }
}

interface CompressPayload { id: number; quality?: number; format?: string; smart?: boolean }

async function handleCompress(payload: CompressPayload) {
  try {
    const formatQuery = payload.format ? `&format=${payload.format}` : '';
    const baseQuery = payload.smart ? 'smart=true' : `quality=${payload.quality ?? calcQuality()}`;
    const res = await fetch(`/api/media/compress/${payload.id}?${baseQuery}${formatQuery}`)
    if (!res.ok) throw new Error('خطا در فشرده‌سازی');
    const data = await res.json();
    // Update file in array and selectedFile
    const idx = files.value.findIndex(f => f.id === payload.id);
    if (idx !== -1) {
      files.value[idx].compressed_size = data.compressed_size ?? data.size;
      (files.value[idx] as any).compressed = true;
      if (data.size) {
        files.value[idx].size = data.size;
      }
      if (data.url) {
        files.value[idx].url = data.url;
      }
      else {
        files.value[idx].url = files.value[idx].url.split('?')[0];
      }
      const extMatch = data.url.match(/\.([a-zA-Z0-9]+)$/);
      if (extMatch) {
        files.value[idx].extension = extMatch[1].toLowerCase();
      }
    }
    if (selectedFile.value && selectedFile.value.id === payload.id) {
      selectedFile.value.compressed_size = data.compressed_size ?? data.size;
      (selectedFile.value as any).compressed = true;
      if (data.size) {
        selectedFile.value.size = data.size;
      }
      if (data.url) {
        selectedFile.value.url = data.url;
      } else {
        selectedFile.value.url = selectedFile.value.url.split('?')[0];
      }
    }
  } catch (e:any) {
    notifier.showError(e.message || 'خطا در فشرده‌سازی')
  }
}

async function handleRevert(id: number) {
  try {
    const res = await fetch(`/api/media/revert/${id}`)
    if (!res.ok) throw new Error('خطا در بازگردانی');
    const data = await res.json();
    const idx = files.value.findIndex(f => f.id === id);
    if (idx !== -1) {
      files.value[idx].compressed_size = null;
      (files.value[idx] as any).compressed = false;
      if (data.url) {
        files.value[idx].url = data.url;
        const extMatch = data.url.match(/\.([a-zA-Z0-9]+)$/);
        if (extMatch) files.value[idx].extension = extMatch[1].toLowerCase();
      } else {
        // If the URL has not changed (same path), append a timestamp to force reload
        files.value[idx].url = files.value[idx].url.split('?')[0] + '?t=' + Date.now();
      }
      if (data.size) files.value[idx].size = data.size;
    }
    if (selectedFile.value && selectedFile.value.id === id) {
      selectedFile.value.compressed_size = null;
      (selectedFile.value as any).compressed = false;
      if (data.url) {
        selectedFile.value.url = data.url;
      } else if (selectedFile.value.url) {
        selectedFile.value.url = selectedFile.value.url.split('?')[0] + '?t=' + Date.now();
      }
      if (data.size) selectedFile.value.size = data.size;
    }
  } catch (e:any) {
    notifier.showError(e.message || 'خطا در بازگردانی')
  }
}

const filterByType = (type: string) => {
  selectedFilter.value = type;
};

const filteredSize = computed(() => {
  const getSize = (f: MediaFile) => (f.compressed_size ?? f.size);
  if (selectedFilter.value === 'images') {
    return files.value.filter(f => f.type === 'image').reduce((sum, f) => sum + getSize(f), 0);
  } else if (selectedFilter.value === 'videos') {
    return files.value.filter(f => f.type === 'video').reduce((sum, f) => sum + getSize(f), 0);
  }
  return files.value.reduce((sum, f) => sum + getSize(f), 0);
});

const reloadCategory = (category: string) => {
  selectedCategory.value = category
  files.value = []
  fetchMediaList()
}

const handleFileClick = (file: MediaFile, event: MouseEvent) => {
  if (event.ctrlKey) {
    // حالت چند انتخابی
    const idx = selectedItems.value.indexOf(file.id)
    if (idx === -1) {
      selectedItems.value.push(file.id)
    } else {
      selectedItems.value.splice(idx, 1)
    }
  } else {
    // انتخاب تکی
    selectedItems.value = [file.id]
  }
}

const hasUncompressedSelected = computed(() => {
  return selectedItems.value.some(id => {
    const file = files.value.find(f => f.id === id)
    if (!file || file.type !== 'image') return false
    return !isAlreadyCompressed(file)
  })
})
</script> 

