     <template>
     <div class="min-h-screen bg-gray-50">
     <!-- Header -->
     <div class="bg-white shadow-sm border-b border-gray-200 sticky top-0 z-10">
          <div class="px-4 py-3">
          <div class="flex items-center justify-between">
               <div>
               <h1 class="text-2xl font-bold text-gray-900">مدیریت رسانه</h1>
               <p class="text-sm text-gray-500 mt-1">مدیریت فایل‌های تصویری و ویدیویی</p>
               </div>
          </div>
          </div>
     </div>

     <!-- Content -->
     <div class="p-6">
          <!-- Media Management Cards -->
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
          
          <!-- Media Library -->
          <NuxtLink to="/admin/media-management/library" class="block">
               <div class="bg-white rounded-lg shadow-md hover:shadow-lg transition-shadow p-6 border border-gray-200 hover:border-blue-300">
               <div class="flex items-center mb-4">
               <div class="bg-blue-100 p-3 rounded-lg">
                    <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    
                    </svg>
               </div>
               <h3 class="text-lg font-semibold text-gray-900 mr-3">کتابخانه رسانه</h3>
               </div>
               <p class="text-gray-600 text-sm">مشاهده و مدیریت فایل‌های آپلود شده</p>
               </div>
          </NuxtLink>

          <!-- Upload Media -->
          <NuxtLink to="/admin/media-management/upload" class="block">
               <div class="bg-white rounded-lg shadow-md hover:shadow-lg transition-shadow p-6 border border-gray-200 hover:border-green-300">
               <div class="flex items-center mb-4">
               <div class="bg-green-100 p-3 rounded-lg">
                    <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"/>
                    </svg>
               </div>
               <h3 class="text-lg font-semibold text-gray-900 mr-3">آپلود فایل</h3>
               </div>
               <p class="text-gray-600 text-sm">آپلود فایل‌های جدید تصویری و ویدیویی</p>
               </div>
          </NuxtLink>

          <!-- Image Compression -->
          <NuxtLink to="/admin/media-management/image-compression" class="block">
               <div class="bg-white rounded-lg shadow-md hover:shadow-lg transition-shadow p-6 border border-gray-200 hover:border-purple-300">
               <div class="flex items-center mb-4">
               <div class="bg-purple-100 p-3 rounded-lg">
                    <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"/>
                    </svg>
               </div>
               <h3 class="text-lg font-semibold text-gray-900 mr-3">فشرده‌سازی تصاویر</h3>
               </div>
               <p class="text-gray-600 text-sm">تنظیمات و بهینه‌سازی تصاویر</p>
               </div>
          </NuxtLink>

          <!-- Video Compression -->
          <NuxtLink to="/admin/media-management/video-compression" class="block">
          <div class="bg-white rounded-lg shadow-md hover:shadow-lg transition-shadow p-6 border border-gray-200 hover:border-red-300">
               <div class="flex items-center mb-4">
               <div class="bg-red-100 p-3 rounded-lg">
                    <svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z"/>
                    </svg>
               </div>
               <h3 class="text-lg font-semibold text-gray-900 mr-3">فشرده‌سازی ویدیو</h3>
               </div>
               <p class="text-gray-600 text-sm">تنظیمات و بهینه‌سازی ویدیوها</p>
               </div>
          </NuxtLink>

          </div>

          <!-- Statistics Section -->
          <div class="mt-8">
          <h2 class="text-xl font-semibold text-gray-900 mb-4">آمار کلی</h2>
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
               
               <div class="bg-white rounded-lg shadow p-6 border border-gray-200">
               <div class="flex items-center">
               <div class="bg-blue-100 p-2 rounded-lg">
                    <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"/>
                         </svg>
               </div>
               <div class="mr-3">
                    <h3 class="text-sm font-medium text-gray-500">کل تصاویر</h3>
                    <p class="text-2xl font-semibold text-gray-900">-</p>
               </div>
               </div>
               </div>

               <div class="bg-white rounded-lg shadow p-6 border border-gray-200">
               <div class="flex items-center">
               <div class="bg-green-100 p-2 rounded-lg">
                    <svg class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z"/>
                    </svg>
               </div>
               <div class="mr-3">
                    <h3 class="text-sm font-medium text-gray-500">کل ویدیوها</h3>
                    <p class="text-2xl font-semibold text-gray-900">-</p>
               </div>
               </div>
               </div>

          <div class="bg-white rounded-lg shadow p-6 border border-gray-200">
               <div class="flex items-center">
               <div class="bg-purple-100 p-2 rounded-lg">
                    <svg class="w-5 h-5 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7v10c0 2.21 3.582 4 8 4s8-1.79 8-4V7M4 7c0 2.21 3.582 4 8 4s8-1.79 8-4M4 7c0-2.21 3.582-4 8-4s8 1.79 8 4"/>
                    </svg>
               </div>
               <div class="mr-3">
                    <h3 class="text-sm font-medium text-gray-500">فضای استفاده شده</h3>
                    <p class="text-2xl font-semibold text-gray-900">-</p>
               </div>
               </div>
               </div>

          </div>
          </div>
     </div>
     </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
declare const useHead: (head: { title?: string }) => void
</script>

<script setup lang="ts">
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

useHead({
  title: 'مدیریت رسانه - پنل مدیریت'
})
</script>
