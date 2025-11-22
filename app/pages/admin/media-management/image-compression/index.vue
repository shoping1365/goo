<template>
  <div class="min-h-screen">
    <!-- Header -->
    <div class="bg-white shadow-sm border-b border-gray-200 sticky top-0 z-10">
      <div class="px-3 py-2">
        <div class="flex items-center justify-between">
          <div>
            <h1 class="text-lg font-bold text-gray-900">تنظیمات تصاویر</h1>
            <p class="text-xs text-gray-500 mt-0.5">بهینه‌سازی و کاهش حجم تصاویر</p>
          </div>

          <div class="flex items-center gap-3">
            <NuxtLink
                to="/admin/media"
                class="bg-gradient-to-l from-pink-200 via-purple-200 to-blue-200 text-gray-800 px-4 py-2 rounded-lg shadow-md hover:from-pink-300 hover:to-blue-300 hover:shadow-lg transition-all flex items-center gap-2 font-bold border border-blue-100"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
              </svg>
              بازگشت به کتابخانه
            </NuxtLink>
          </div>
        </div>
      </div>
    </div>

    <!-- moved cards will be placed at bottom -->

    <!-- Main Content -->
    <div class="px-0 max-w-none">
      <!-- Image SEO (AI) Settings & Jobs -->
      <div class="bg-white rounded-2xl shadow-lg border border-gray-100 mb-6 mt-6 mx-4">
        <div class="flex items-center justify-between px-3 py-4 border-b border-blue-300 relative">
          <div class="flex items-center gap-2 cursor-pointer pr-0" @click="toggleSection('imageSEO')">
            <span class="text-base font-bold bg-emerald-100 text-emerald-700 px-4 py-1 rounded-lg ml-auto">هوش‌مصنوعی تصاویر (ALT/CAPTION)</span>
          </div>
          <span class="absolute left-4 top-1/2 -translate-y-1/2 text-gray-500 text-2xl select-none cursor-pointer" @click="toggleSection('imageSEO')">{{ sections.imageSEO ? '−' : '+' }}</span>
        </div>
        <div v-show="sections.imageSEO" class="p-6 space-y-6">
          <!-- Settings Row -->
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div class="space-y-3">
              <h3 class="font-semibold text-gray-900 mb-2">تنظیمات</h3>
              <label class="flex items-center gap-2">
                <input v-model="imageSeo.enabled" type="checkbox" class="w-4 h-4" />
                <span class="text-sm">فعال‌سازی تولید خودکار ALT/CAPTION</span>
              </label>
              <label class="flex items-center gap-2">
                <input v-model="imageSeo.generateTitle" type="checkbox" class="w-4 h-4" />
                <span class="text-sm">تولید Title (اختیاری)</span>
              </label>
              <label class="flex items-center gap-2">
                <input v-model="imageSeo.overwriteExisting" type="checkbox" class="w-4 h-4" />
                <span class="text-sm">بازنویسی مقادیر موجود (پیشنهادی: خاموش)</span>
              </label>
            </div>
              <div class="space-y-3">
                <h3 class="font-semibold text-gray-900 mb-2">مدل و زبان</h3>
                <div class="flex items-center gap-2">
                  <label class="text-sm w-24">مدل</label>
                  <template v-if="availableModels.length">
                    <select v-model="imageSeo.model" class="px-3 py-1 border rounded w-full">
                      <option v-for="m in availableModels" :key="m.id" :value="m.id">{{ m.name || m.id }}</option>
                    </select>
                  </template>
                  <template v-else>
                    <input v-model="imageSeo.model" class="px-3 py-1 border rounded w-full" placeholder="gpt-4o-mini" />
                  </template>
                </div>
              <div class="flex items-center gap-2">
                <label class="text-sm w-24">زبان</label>
                <input v-model="imageSeo.lang" class="px-3 py-1 border rounded w-full" placeholder="fa" />
              </div>
            </div>
            <div class="space-y-3">
              <h3 class="font-semibold text-gray-900 mb-2">زمان‌بندی</h3>
              <div class="flex items-center gap-2">
                <label class="text-sm w-36">تاخیر پس از آپلود (ثانیه)</label>
                <input v-model.number="imageSeo.delaySeconds" type="number" min="5" class="px-3 py-1 border rounded w-full" />
              </div>
              <div class="flex items-center gap-2">
                <label class="text-sm w-36">Batch Size</label>
                <input v-model.number="imageSeo.batchSize" type="number" min="1" class="px-3 py-1 border rounded w-full" />
              </div>
            </div>
          </div>
          <div class="flex items-center gap-3">
            <button class="bg-green-500 hover:bg-green-600 text-white text-sm px-4 py-1.5 rounded-lg shadow" @click="saveImageSeoSettings">ذخیره تنظیمات</button>
            <span v-if="imageSeoSaved" class="text-xs text-green-700">ذخیره شد</span>
          </div>

          <!-- Scan and Jobs Row -->
          <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
            <!-- Scanner -->
            <div class="border rounded-lg p-6">
              <div class="flex items-center justify-between mb-3">
                <h4 class="font-semibold">اسکن تصاویر بدون ALT/Title</h4>
                <button class="bg-purple-100 text-purple-700 px-3 py-1 rounded-lg hover:bg-purple-200 text-sm" @click="scanMissingMeta">اسکن</button>
              </div>
              <p class="text-sm text-gray-600 mb-2">یافت‌شده: {{ missingList.length }}</p>
              <div class="max-h-56 overflow-auto text-sm space-y-1">
                <div v-for="m in missingList" :key="m.id" class="flex items-center justify-between border-b py-1">
                  <span class="truncate mr-2">#{{ m.id }} · {{ m.name }}</span>
                  <button class="bg-blue-100 text-blue-700 px-2 py-0.5 rounded hover:bg-blue-200 text-xs" @click="retryForMedia(m.id)">ایجاد Job</button>
                </div>
              </div>
            </div>

            <!-- Jobs Table -->
            <div class="lg:col-span-2 border rounded-lg p-6">
              <div class="flex items-center justify-between mb-3">
                <h4 class="font-semibold">صف پردازش ALT/CAPTION</h4>
                <div class="flex items-center gap-2">
                  <select v-model="jobFilter" class="px-2 py-1 border rounded text-sm">
                    <option value="">همه</option>
                    <option value="pending">pending</option>
                    <option value="processing">processing</option>
                    <option value="completed">completed</option>
                    <option value="error">error</option>
                  </select>
                  <button class="bg-gray-100 text-gray-800 px-3 py-1 rounded text-sm" @click="loadJobs">رفرش</button>
                </div>
              </div>
              <div class="overflow-auto">
                <table class="w-full text-sm">
                  <thead class="bg-gray-50">
                    <tr class="text-right">
                      <th class="px-3 py-2">ID</th>
                      <th class="px-3 py-2">Media</th>
                      <th class="px-3 py-2">Status</th>
                      <th class="px-3 py-2">Scheduled</th>
                      <th class="px-3 py-2">Finished</th>
                      <th class="px-3 py-2">Action</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr v-if="jobs.length===0"><td class="px-3 py-3 text-gray-400" colspan="6">موردی یافت نشد</td></tr>
                    <tr v-for="j in jobs" :key="j.id" class="border-b">
                      <td class="px-3 py-2">{{ j.id }}</td>
                      <td class="px-3 py-2">{{ j.media_id }}</td>
                      <td class="px-3 py-2">{{ j.status }}</td>
                      <td class="px-3 py-2">{{ formatDateTime(j.scheduled_at as string | number | Date | null | undefined) }}</td>
                      <td class="px-3 py-2">{{ formatDateTime(j.finished_at as string | number | Date | null | undefined) }}</td>
                      <td class="px-3 py-2 flex items-center gap-2">
                        <button v-if="j.status==='error' || j.status==='completed'" class="bg-blue-100 text-blue-700 px-2 py-0.5 rounded hover:bg-blue-200" @click="retryJob(j.id)">Retry</button>
                        <span v-if="j.error_message" class="text-xs text-red-600 truncate max-w-[240px]">{{ j.error_message }}</span>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>
          </div>
        </div>
      </div>
      <!-- Compression Settings -->
      <div class="bg-white rounded-2xl shadow-lg border border-gray-100 mb-6 mt-6 mx-4">
        <div class="flex items-center justify-between px-3 py-4 border-b border-blue-300 relative">
          <!-- عنوان تاشو -->
          <div class="flex items-center gap-2 cursor-pointer pr-0" @click="toggleSection('compressionSettings')">
            <span class="text-base font-bold bg-pink-100 text-rose-700 px-4 py-1 rounded-lg ml-auto">تنظیمات فشرده‌سازی تصاویر</span>
          </div>
          <!-- toggle icon moved to far left -->
          <span class="absolute left-4 top-1/2 -translate-y-1/2 text-gray-500 text-2xl select-none cursor-pointer" @click="toggleSection('compressionSettings')">{{ sections.compressionSettings ? '−' : '+' }}</span>

          <!-- سوییچ فعال‌سازی سراسری + دکمه ذخیره -->
          <div class="flex items-center gap-6">
            <!-- Toggle -->
            <span class="text-sm text-gray-700">فشرده‌سازی خودکار</span>
            <label class="inline-flex items-center cursor-pointer select-none">
              <input v-model="compressionSettings.enabled" type="checkbox" class="sr-only peer">
              <div class="w-11 h-6 bg-gray-200 rounded-full relative transition peer-checked:bg-green-500">
                <span class="absolute left-1 top-1 w-4 h-4 bg-white rounded-full transition-transform peer-checked:translate-x-5"></span>
              </div>
            </label>

            <!-- Save button (shown only on unsaved changes) -->
            <button
                v-if="unsavedChanges"
                class="bg-green-500 hover:bg-green-600 text-white text-sm px-4 py-1.5 rounded-lg shadow ml-8"
                @click="saveCompressionSettings"
            >
              ذخیره تغییرات
            </button>
          </div>
        </div>
        <div v-show="sections.compressionSettings" class="p-6">
          <div :class="[ 'grid grid-cols-1 md:grid-cols-3 gap-8', !compressionSettings.enabled && 'opacity-40 pointer-events-none']">
            <!-- Quality Settings -->
            <div class="space-y-6">
              <h3 class="font-semibold text-gray-900 mb-2">کیفیت خروجی</h3>
              <div class="space-y-4">
                <!-- گزینه‌های کیفیت خروجی، غیرفعال در حالت هوشمند -->
                <!-- Output quality options, disabled in smart mode -->
                <label class="flex items-center">
                  <input
                      v-model="compressionSettings.quality"
                      type="radio"
                      value="high"
                      :disabled="isLoadingCompressionSettings"
                      class="w-4 h-4 text-blue-600 border-gray-300 focus:ring-blue-500"
                  >
                  <span class="mr-2 text-sm">بالا (90%) - کیفیت عالی، حجم بیشتر</span>
                </label>
                <label class="flex items-center">
                  <input
                      v-model="compressionSettings.quality"
                      type="radio"
                      value="medium"
                      :disabled="isLoadingCompressionSettings"
                      class="w-4 h-4 text-blue-600 border-gray-300 focus:ring-blue-500"
                  >
                  <span class="mr-2 text-sm">متوسط (75%) - تعادل بین کیفیت و حجم</span>
                </label>
                <label class="flex items-center">
                  <input
                      v-model="compressionSettings.quality"
                      type="radio"
                      value="low"
                      :disabled="isLoadingCompressionSettings"
                      class="w-4 h-4 text-blue-600 border-gray-300 focus:ring-blue-500"
                  >
                  <span class="mr-2 text-sm">پایین (60%) - حجم کم، کیفیت قابل قبول</span>
                </label>

                <!-- گزینه هوشمند -->
                <label class="flex items-center">
                  <input
                      v-model="compressionSettings.quality"
                      type="radio"
                      value="smart"
                      :disabled="isLoadingCompressionSettings"
                      class="w-4 h-4 text-blue-600 border-gray-300 focus:ring-blue-500"
                  >
                  <span class="mr-2 text-sm text-blue-700 font-semibold cursor-pointer">فشرده‌سازی هوشمند (پیشنهادی)</span>
                </label>

                <label class="flex items-center">
                  <input
                      v-model="compressionSettings.quality"
                      type="radio"
                      value="custom"
                      :disabled="isLoadingCompressionSettings"
                      class="w-4 h-4 text-blue-600 border-gray-300 focus:ring-blue-500"
                  >
                  <span class="mr-2 text-sm">سفارشی</span>
                </label>
                <div v-if="compressionSettings.quality === 'custom'" class="mt-3">
                  <label class="block text-sm font-medium text-gray-700 mb-2">
                    کیفیت: {{ compressionSettings.customQuality }}%
                  </label>
                  <input
                      v-model="compressionSettings.customQuality"
                      type="range"
                      min="10"
                      max="100"
                      class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer"
                      :disabled="compressionSettings.quality !== 'custom'"
                  >
                </div>
              </div>
            </div>
            <!-- Format Settings -->
            <div class="space-y-6">
              <h3 class="font-semibold text-gray-900 mb-2">فرمت خروجی</h3>
              <div class="space-y-4">
                <label class="flex items-center">
                  <input
                      v-model="compressionSettings.format"
                      type="radio"
                      value="original"
                      class="w-4 h-4 text-blue-600 border-gray-300 focus:ring-blue-500"
                  >
                  <span class="mr-2 text-sm">حفظ فرمت اصلی</span>
                </label>
                <label class="flex items-center">
                  <input
                      v-model="compressionSettings.format"
                      type="radio"
                      value="webp"
                      class="w-4 h-4 text-blue-600 border-gray-300 focus:ring-blue-500"
                  >
                  <span class="mr-2 text-sm">WebP (بهترین فشرده‌سازی)</span>
                </label>
                <label class="flex items-center">
                  <input
                      v-model="compressionSettings.format"
                      type="radio"
                      value="jpeg"
                      class="w-4 h-4 text-blue-600 border-gray-300 focus:ring-blue-500"
                  >
                  <span class="mr-2 text-sm">JPEG (سازگاری بالا)</span>
                </label>
                <label class="flex items-center">
                  <input
                      v-model="compressionSettings.format"
                      type="radio"
                      value="png"
                      class="w-4 h-4 text-blue-600 border-gray-300 focus:ring-blue-500"
                  >
                  <span class="mr-2 text-sm">PNG (شفافیت)</span>
                </label>
              </div>
            </div>
            <!-- Advanced Options -->
            <div class="space-y-6">
              <h3 class="font-semibold text-gray-900 mb-2">تنظیمات پیشرفته</h3>
              <div class="grid grid-cols-1 gap-3">
                <label class="flex items-center">
                  <input
                      v-model="compressionSettings.progressive"
                      type="checkbox"
                      :disabled="compressionSettings.format !== 'jpeg'"
                      class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500 disabled:opacity-40 disabled:cursor-not-allowed"
                  >
                  <span class="mr-2 text-sm">JPEG پیش‌رونده</span>
                </label>
                <label class="flex items-center">
                  <input
                      v-model="compressionSettings.removeMetadata"
                      type="checkbox"
                      class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                  >
                  <span class="mr-2 text-sm">حذف متادیتا</span>
                </label>
                <label class="flex items-center">
                  <input
                      v-model="compressionSettings.optimizeColors"
                      type="checkbox"
                      class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                  >
                  <span class="mr-2 text-sm">بهینه‌سازی رنگ‌ها</span>
                </label>
                <label class="flex items-center">
                  <input
                      v-model="compressionSettings.createBackup"
                      type="checkbox"
                      class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                  >
                  <span class="mr-2 text-sm">ایجاد نسخه پشتیبان</span>
                </label>
                <label class="flex items-center">
                  <input
                      v-model="compressionSettings.keepOriginalDate"
                      type="checkbox"
                      class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                  >
                  <span class="mr-2 text-sm">حفظ تاریخ اصلی فایل</span>
                </label>
                <label class="flex items-center">
                  <input
                      v-model="compressionSettings.autoOrient"
                      type="checkbox"
                      class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                  >
                  <span class="mr-2 text-sm">اصلاح خودکار جهت تصویر (Auto-Orientation)</span>
                </label>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Image Selection -->
      <div class="bg-white rounded-2xl shadow-lg border border-gray-200 mb-6 mx-4">
        <div class="flex items-center cursor-pointer px-6 py-4 relative">
          <!-- Toggle icon -->
          <span class="absolute left-4 top-1/2 -translate-y-1/2 text-gray-500 text-2xl select-none cursor-pointer" @click="toggleSection('imageSelection')">{{ sections.imageSelection ? '−' : '+' }}</span>
          <!-- Scan + dynamic action buttons container -->
          <div class="absolute left-12 top-1/2 -translate-y-1/2 flex items-center gap-2 select-none">
            <!-- Scan -->
            <button class="bg-purple-100 text-purple-700 text-xs md:text-sm px-3 py-1 rounded-lg hover:bg-purple-200 transition-colors" @click.stop="scanImages">اسکن</button>
            <!-- Select All (shown when at least one selected & not all) -->
            <button v-if="selectedImages.length > 0 && selectedImages.length < images.length" class="bg-indigo-100 text-indigo-700 text-xs md:text-sm px-3 py-1 rounded-lg hover:bg-indigo-200 transition-colors" @click.stop="selectAllImages">انتخاب همه</button>
            <!-- Compress -->
            <button v-if="selectedImages.length > 0" :disabled="isCompressing" class="bg-pink-100 text-pink-700 text-xs md:text-sm px-3 py-1 rounded-lg hover:bg-pink-200 transition-colors disabled:opacity-50" @click.stop="startCompression">فشرده‌سازی</button>
          </div>
          <!-- Title pill -->
          <span class="text-base font-bold bg-purple-100 text-purple-800 px-4 py-1 rounded-lg ml-auto">فشرده‌سازی دستی</span>
          <!-- Inline no-image notice, centered -->
          <span
              v-if="scanPerformed && images.length === 0"
              class="absolute left-1/2 -translate-x-1/2 text-sm text-red-600 whitespace-nowrap"
          >تصویر فشرده‌نشده‌ای یافت نشد</span>
        </div>
        <div class="border-b border-gray-200 mx-6"></div>
        <div v-show="sections.imageSelection">
          <!-- نوار دکمه‌های انتخاب/فشرده‌سازی حذف شد طبق درخواست -->
          <!-- این div که قبلاً border-b داشت، حالا بدون border است -->
          <div class="px-6 py-4 flex items-center justify-between">
            <!-- نوار عنوان و دکمه‌های انتخاب تصاویر حذف شد -->
          </div>

          <!-- Images Grid -->
          <div class="p-6">
            <!-- Empty state permanently removed -->
            <div v-if="false" class="hidden">
              <svg class="w-24 h-24 text-gray-300 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
              </svg>
              <h3 class="text-lg font-medium text-gray-900 mb-2">تصویری برای فشرده‌سازی یافت نشد</h3>
            </div>

            <div v-if="images.length > 0" class="grid grid-cols-3 md:grid-cols-6 lg:grid-cols-8 xl:grid-cols-10 gap-3">
              <div
                  v-for="image in images"
                  :key="image.id"
                  :class="[
                 'relative group cursor-pointer rounded-xl border-2 transition-all duration-200 overflow-hidden shadow-md hover:shadow-lg',
                 selectedImages.includes(image.id)
                   ? 'border-blue-500 bg-blue-50'
                   : 'border-gray-200 hover:border-gray-300'
               ]"
              >
                <!-- Image Preview -->
                <div class="aspect-square bg-gray-100 flex items-center justify-center relative" @click="openImagePreview(image)">
                  <img
                      :src="image.thumbnail"
                      :alt="image.name"
                      class="w-full h-full object-cover"
                  >
                  <!-- Eye icon ... -->
                  <button class="absolute bottom-2 left-2 bg-white bg-opacity-70 hover:bg-opacity-100 rounded-full p-1 shadow opacity-0 group-hover:opacity-100 transition-opacity duration-200" @click.stop="openImagePreview(image)">
                    <svg class="w-4 h-4 text-gray-700" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                    </svg>
                  </button>
                </div>

                <!-- Image Info (کلیک روی این بخش فقط انتخاب را تغییر دهد) -->
                <div
                    class="p-3 bg-white relative"
                    style="cursor:pointer"
                    @click.stop="handleImageSelect(image.id, $event)"
                >
                  <p class="text-sm font-medium text-gray-900 truncate" :title="image.name">
                    {{ image.name }}
                  </p>
                  <div class="flex items-center justify-between mt-1">
                    <p class="text-xs text-gray-500">{{ formatFileSize(image.size) }}</p>
                    <p class="text-xs text-gray-500">{{ image.dimensions }}</p>
                  </div>
                  <!-- هیچ چک‌باکسی نمایش داده نشود -->
                </div>

                <!-- Compression Status -->
                <div v-if="image.compressionStatus" class="absolute inset-0 bg-black bg-opacity-75 flex items-center justify-center">
                  <div class="text-center text-white">
                    <svg v-if="image.compressionStatus === 'processing'" class="animate-spin w-8 h-8 mx-auto mb-2" fill="none" viewBox="0 0 24 24">
                      <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                      <path class="opacity-75" fill="currentColor" d="m4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                    </svg>
                    <svg v-else-if="image.compressionStatus === 'completed'" class="w-8 h-8 mx-auto mb-2 text-green-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                    </svg>
                    <svg v-else-if="image.compressionStatus === 'error'" class="w-8 h-8 mx-auto mb-2 text-red-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                    </svg>

                    <p class="text-sm font-medium">
                      <span v-if="image.compressionStatus === 'processing'">در حال پردازش...</span>
                      <span v-else-if="image.compressionStatus === 'completed'">تکمیل شد</span>
                      <span v-else-if="image.compressionStatus === 'error'">خطا</span>
                    </p>

                    <p v-if="image.compressionStatus === 'completed' && image.compressionResult" class="text-xs mt-1">
                      {{ image.compressionResult.reduction }}% کاهش حجم
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Compression Results -->
    <div class="bg-white rounded-2xl shadow-lg border border-gray-200 mb-6 mx-4 p-6 pb-[10px] flex flex-col justify-between">
      <div class="flex items-center justify-between flex-wrap mb-4 w-full">
        <!-- Title (right side in RTL) -->
        <div class="flex items-center flex-wrap gap-2 order-last">
          <!-- Filter dropdown -->
          <select v-model="filterStatus" class="px-4 py-2 rounded-lg border border-gray-300 focus:ring-2 focus:ring-blue-400 text-sm order-last">
            <option value="">همه وضعیت‌ها</option>
            <option value="success">موفق</option>
            <option value="error">خطا</option>
            <option value="processing">در حال پردازش</option>
          </select>
          <!-- Clear table button -->
          <button v-if="compressionResults.length > 0" class="bg-red-100 text-red-700 px-4 py-2 rounded-lg shadow hover:bg-red-200 transition-all font-bold text-sm" @click="clearCompressionResults">
            پاکسازی جدول
          </button>
          <!-- Batch actions (shown when results exist) -->
          <template v-if="compressionResults.length > 0">
            <button v-if="selectedResults.length >= 3" class="bg-yellow-100 text-yellow-800 px-4 py-2 rounded-lg shadow hover:bg-yellow-200 transition-all font-bold border border-yellow-200 flex items-center gap-2" @click="batchRestore">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path></svg>
              بازگردانی انتخاب‌شده‌ها
            </button>
            <button v-if="selectedResults.length > 0" class="bg-blue-100 text-blue-800 px-4 py-2 rounded-lg shadow hover:bg-blue-200 transition-all font-bold border border-blue-200 flex items-center gap-2" @click="clearResultsSelection">
              لغو انتخاب
            </button>
            <button v-if="selectedResults.length >= 3" class="bg-red-100 text-red-800 px-4 py-2 rounded-lg shadow hover:bg-red-200 transition-all font-bold border border-red-200 flex items-center gap-2" @click="deleteSelectedResults">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path></svg>
              حذف انتخاب‌شده‌ها
            </button>
          </template>
        </div>
        <!-- Title -->
        <span class="text-base font-bold bg-blue-100 text-blue-800 px-3 py-1 rounded-lg ml-auto">آخرین تصاویر فشرده</span>
      </div>
      <!-- (Old filter row removed; actions consolidated) -->
      <div class="flex flex-col lg:flex-row gap-8 flex-1">



        <!-- جدول نتایج -->
        <div class="flex-[2] overflow-x-auto">
          <div class="overflow-x-auto">
            <table class="w-full">
              <thead class="bg-gray-50 border-b border-gray-200">
              <tr class="text-right">
                <th class="px-6 py-3"><input v-model="selectAllResults" type="checkbox" @change="toggleSelectAllResults" /></th>
                <th class="px-6 py-3 w-32 min-w-[8rem] text-xs font-medium text-gray-500 uppercase tracking-wider text-right">تصویر</th>
                <th class="px-6 py-3 text-xs font-medium text-gray-500 uppercase tracking-wider text-right">حجم اصلی</th>
                <th class="px-6 py-3 text-xs font-medium text-gray-500 uppercase tracking-wider text-right">حجم جدید</th>
                <th class="px-6 py-3 text-xs font-medium text-gray-500 uppercase tracking-wider text-right">کاهش حجم</th>
                <th class="px-6 py-3 text-xs font-medium text-gray-500 uppercase tracking-wider text-right">وضعیت</th>
                <th class="px-6 py-3 text-xs font-medium text-gray-500 uppercase tracking-wider text-right">عملیات</th>
              </tr>
              </thead>

              <tbody class="bg-white divide-y divide-gray-200">
              <tr v-if="filteredResults.length === 0">
                <td colspan="7" class="text-center py-8 text-gray-400">هنوز تصویری فشرده نشده است</td>
              </tr>
              <tr v-for="result in filteredResults" :key="result.id" class="hover:bg-gray-50">
                <td class="px-6 py-4"><input v-model="selectedResults" type="checkbox" :value="result.id" /></td>
                <td class="px-6 py-4 w-32 min-w-[8rem]">
                  <div class="flex items-center">
                    <div class="flex-shrink-0 h-16 w-16">
                      <img
                          :src="result.thumbnail"
                          :alt="result.name"
                          class="h-16 w-16 rounded-lg object-cover cursor-pointer"
                          @click="openImagePreview({
                          id: result.id,
                          name: result.name,
                          size: result.compressedSize,
                          url: result.compressedUrl || result.thumbnail,
                          thumbnail: result.compressedUrl || result.thumbnail,
                          dimensions: result.dimensions || '',
                          extension: result.name.split('.').pop() || '',
                          type: 'image'
                        } as ImageFile)"
                      >
                    </div>
                    <div class="mr-4">
                      <div class="text-sm font-medium text-gray-900">{{ result.name }}</div>
                      <div class="text-sm text-gray-500">{{ result.dimensions }}</div>
                      <div v-if="result.status === 'error'" class="text-xs text-red-600 mt-1">خطا در فشرده‌سازی این تصویر</div>
                    </div>
                  </div>
                </td>
                <td class="px-6 py-4 text-sm text-gray-900">{{ formatFileSize(result.originalSize) }}</td>
                <td class="px-6 py-4 text-sm text-gray-900">{{ formatFileSize(result.compressedSize) }}</td>
                <td class="px-6 py-4 text-sm">
                  <span
:class="[
                    'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium',
                    result.reduction > 50 ? 'bg-green-100 text-green-800' :
                    result.reduction > 25 ? 'bg-yellow-100 text-yellow-800' :
                    'bg-red-100 text-red-800'
                  ]">
                    {{ result.reduction }}%
                  </span>
                </td>
                <td class="px-6 py-4 text-sm">
                  <span
:class="[
                    'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium',
                    result.status === 'success' ? 'bg-green-100 text-green-800' :
                    result.status === 'error' ? 'bg-red-100 text-red-800' :
                    'bg-yellow-100 text-yellow-800'
                  ]">
                    {{ result.status === 'success' ? 'موفق' : result.status === 'error' ? 'خطا' : 'در حال پردازش' }}
                  </span>
                </td>
                <td class="px-6 py-4 text-sm font-medium">
                  <div class="flex flex-wrap items-center gap-2">
                    <button
                        class="bg-blue-100 text-blue-800 px-3 py-1 rounded-lg hover:bg-blue-200 transition-colors flex items-center gap-1"
                        title="دانلود"
                        @click="downloadCompressed(result)"
                    >
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                      </svg>
                      دانلود
                    </button>
                    <button
                        v-if="compressionSettings.createBackup"
                        class="bg-yellow-100 text-yellow-800 px-3 py-1 rounded-lg hover:bg-yellow-200 transition-colors flex items-center gap-1"
                        title="بازگردانی"
                        @click="replaceOriginal(result)"
                    >
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
                      </svg>
                      بازگردانی
                    </button>
                    <button
                        class="bg-purple-100 text-purple-800 px-3 py-1 rounded-lg hover:bg-purple-200 transition-colors flex items-center gap-1"
                        title="مقایسه"
                        @click="compareImages(result)"
                    >
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                      </svg>
                      مقایسه
                    </button>
                    <button
                        v-if="canDeleteCompressionResult"
                        class="bg-red-100 text-red-800 px-3 py-1 rounded-lg hover:bg-red-200 transition-colors flex items-center gap-1"
                        title="حذف"
                        @click="removeResult(result)"
                    >
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                      </svg>
                      حذف
                    </button>
                  </div>
                </td>
              </tr>
              </tbody>
            </table>
          </div>
          <!-- آمار زیر جدول (لبه پایین باکس، با فاصله 10px) -->
          <div class="flex flex-row justify-between items-center gap-6 mt-0 mb-0 px-8 py-3 bg-gradient-to-l from-pink-50 via-blue-50 to-purple-50 rounded-xl border border-gray-100 overflow-x-auto w-max min-w-[500px] mx-auto">
            <div class="flex flex-row items-center gap-1">
              <span class="text-xs text-gray-500">تعداد تصاویر:</span>
              <span class="font-bold text-base text-blue-700">{{ compressionResults.length }}</span>
            </div>
            <div class="flex flex-row items-center gap-1">
              <span class="text-xs text-gray-500">حجم اصلی کل:</span>
              <span class="font-bold text-base text-gray-800">{{ formatFileSize(totalOriginalSize) }}</span>
            </div>
            <div class="flex flex-row items-center gap-1">
              <span class="text-xs text-gray-500">حجم جدید کل:</span>
              <span class="font-bold text-base text-green-700">{{ formatFileSize(totalCompressedSize) }}</span>
            </div>
            <div class="flex flex-row items-center gap-1">
              <span class="text-xs text-gray-500">صرفه‌جویی کل:</span>
              <span class="font-bold text-base text-emerald-700">{{ formatFileSize(totalSavedSize) }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- پیش نمایش فشرده‌سازی تصویر -->
    <div class="bg-white rounded-2xl shadow-lg border border-gray-200 mb-6 mx-4">
      <div class="flex items-center justify-between cursor-pointer px-6 py-4 border-b border-blue-300" @click="toggleSection('preview')">
        <span class="text-base font-bold bg-purple-100 text-purple-800 px-4 py-1 rounded-lg ml-auto">پیش نمایش</span>
        <span class="text-gray-500 text-2xl">{{ sections.preview ? '−' : '+' }}</span>
      </div>
      <div v-show="sections.preview" class="p-10">
        <div class="flex flex-col md:flex-row gap-10 items-center justify-center">
          <!-- تصویر اصلی (قبل) -->
          <div class="flex-1 flex flex-col items-center min-h-[240px] md:min-h-[320px]">
            <div class="relative group w-full h-72 flex items-center justify-center">
              <img v-if="previewImageFile" :src="previewImageUrl || ''" class="w-full h-full object-contain rounded-lg shadow border border-gray-200 bg-gray-50" />
              <button v-if="previewImageFile" class="absolute bottom-2 left-2 bg-white bg-opacity-70 hover:bg-opacity-100 rounded-full p-1 shadow opacity-0 group-hover:opacity-100 transition-opacity duration-200 z-10" @click.stop="openPreviewUrl(previewImageUrl || '', 'عکس قبل')">
                <svg class="w-4 h-4 text-gray-700" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                </svg>
              </button>
            </div>
            <span class="text-xs text-gray-500 mb-2">عکس قبل</span>
            <div v-if="previewUrl" class="mt-2 bg-blue-50 rounded-lg px-3 py-2 text-xs text-gray-700 space-y-1 w-full max-w-xs">
              <div>فرمت ورودی: <span class="font-bold">{{ previewInfo.originalFormat }}</span></div>
              <div>سایز تصویر: <span class="font-bold">{{ previewInfo.originalDimensions }}</span></div>
              <div>حجم تصویر: <span class="font-bold">{{ formatFileSize(previewInfo.originalSize) }}</span></div>
            </div>
          </div>
          <!-- تصویر فشرده‌شده (بعد) -->
          <div class="flex-1 flex flex-col items-center min-h-[240px] md:min-h-[320px]">
            <div class="relative group w-full h-72 flex items-center justify-center">
              <img v-if="previewUrl" :src="previewUrl" class="w-full h-full object-contain rounded-lg shadow border border-green-200 bg-green-50" />
              <button v-if="previewUrl" class="absolute bottom-2 left-2 bg-white bg-opacity-70 hover:bg-opacity-100 rounded-full p-1 shadow opacity-0 group-hover:opacity-100 transition-opacity duration-200 z-10" @click.stop="openPreviewUrl(previewUrl, 'عکس بعد')">
                <svg class="w-4 h-4 text-gray-700" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                </svg>
              </button>
              <div v-else class="w-full h-72 flex items-center justify-center text-gray-400"></div>
            </div>
            <span class="text-xs text-gray-500 mb-2">عکس بعد</span>
            <div v-if="previewUrl" class="mt-2 bg-purple-50 rounded-lg px-3 py-2 text-xs text-gray-700 space-y-1 w-full max-w-xs">
              <div>فرمت خروجی: <span class="font-bold">{{ previewInfo.outputFormat }}</span></div>
              <div>سایز تصویر: <span class="font-bold">{{ previewInfo.compressedDimensions }}</span></div>
              <div>حجم تصویر: <span class="font-bold">{{ formatFileSize(previewInfo.compressedSize) }}</span></div>
            </div>
          </div>
        </div>
        <!-- کنترل‌های پایین: همه در یک ردیف و وسط چین -->
        <div class="flex flex-wrap flex-row gap-6 items-center justify-center my-8 bg-gradient-to-l from-purple-50 via-blue-50 to-pink-50 rounded-xl p-6">
          <!-- دکمه انتخاب تصویر -->
          <input ref="fileInput" type="file" accept="image/*" class="hidden" @change="onPreviewFileChange" />
          <button class="flex items-center gap-2 font-bold px-4 py-2 rounded-lg shadow transition-colors bg-gradient-to-l from-pink-100 via-purple-100 to-blue-100 text-blue-800 hover:from-pink-200 hover:to-blue-200" @click="fileInput && fileInput.click()">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path></svg>
            انتخاب تصویر
          </button>
          <!-- انتخاب فرمت خروجی -->
          <div class="flex items-center gap-2 bg-blue-50 rounded-lg px-3 py-1">
            <label class="text-sm font-medium text-gray-700">فرمت خروجی</label>
            <select v-model="previewFormat" class="px-2 py-1 rounded border border-gray-300 focus:ring-2 focus:ring-blue-400 text-xs w-auto bg-white">
              <option value="jpeg">JPEG</option>
              <option value="webp">WebP</option>
              <option value="png">PNG</option>
            </select>
          </div>
          <!-- کیفیت خروجی عددی -->
          <div v-if="!smartCompression" class="flex items-center gap-2 bg-gradient-to-l from-pink-100 via-blue-50 to-purple-50 rounded-lg px-3 py-1">
            <label class="text-sm font-medium text-gray-700">کیفیت خروجی</label>
            <input v-model="previewCompressionPercent" type="number" min="10" max="100" class="w-16 px-2 py-1 rounded border border-gray-300 text-center text-xs focus:ring-2 focus:ring-blue-400 bg-white" />
            <span class="text-xs text-gray-500">%</span>
          </div>
          <!-- سوییچ هوشمند -->
          <div class="flex items-center gap-2 bg-purple-50 rounded-lg px-3 py-1">
            <label class="flex items-center">
              <input
                  v-model="smartCompression"
                  type="checkbox"
                  class="w-4 h-4 text-blue-600 bg-white border-gray-300 rounded-full focus:ring-blue-500"
              >
              <span class="mr-1 text-sm text-blue-700 cursor-pointer">فشرده‌سازی هوشمند</span>
            </label>
          </div>
          <!-- دکمه پیش‌نمایش -->
          <button :disabled="isPreviewLoading" class="flex items-center gap-2 font-bold px-6 py-2 rounded-lg shadow transition-colors bg-gradient-to-l from-blue-400 to-purple-400 text-white hover:from-blue-500 hover:to-purple-500 disabled:bg-gray-400 disabled:cursor-not-allowed" @click="onPreviewButtonClick">
            <span v-if="isPreviewLoading">در حال تولید پیش‌نمایش...</span>
            <span v-else>پیش‌نمایش</span>
          </button>
        </div>
      </div>
    </div>
  </div>





  <!-- Toast Notification -->
  <div v-if="toast.message" :class="['fixed top-8 left-1/2 z-50 transform -translate-x-1/2 px-6 py-3 rounded-lg shadow-lg flex items-center gap-2', toast.type === 'success' ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800']">
    <svg v-if="toast.type === 'success'" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg>
    <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path></svg>
    <span>{{ toast.message }}</span>
  </div>

  <!-- Loader Animation -->
  <div v-if="isCompressing" class="fixed inset-0 z-40 flex items-center justify-center bg-black bg-opacity-30">
    <div class="bg-white rounded-xl shadow-lg px-8 py-6 flex flex-col items-center gap-6">
      <svg class="animate-spin w-10 h-10 text-blue-500" fill="none" viewBox="0 0 24 24"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle><path class="opacity-75" fill="currentColor" d="m4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path></svg>
      <span class="text-blue-700 font-bold">در حال فشرده‌سازی تصاویر...</span>
    </div>
  </div>






  <!-- MediaPreviewModal modal instance -->
  <ImagePreviewModal
      :model-value="previewModalVisible"
      :image="previewFile"
      @update:model-value="previewModalVisible = $event"
      @close="previewModalVisible = false"
  />

  <!-- Backup Restore Section (moved to the end) -->
  <div class="bg-white rounded-2xl shadow-lg border border-yellow-200 mb-6 mx-4">
    <div class="flex items-center justify-between px-3 py-4 border-b border-blue-300 relative cursor-pointer" @click="toggleSection('backup')">
      <span class="text-base font-bold bg-yellow-100 text-yellow-800 px-4 py-1 rounded-lg ml-auto">بازگردانی بک‌آپ</span>
      <button :disabled="backupState.loading" class="bg-blue-100 text-blue-800 px-3 py-1 rounded-lg hover:bg-blue-200 transition-all ml-8" @click.stop="loadBackupPeriods">
        <span v-if="backupState.loading">در حال دریافت...</span>
        <span v-else>دریافت تاریخ‌ها</span>
      </button>
      <span class="absolute left-4 top-1/2 -translate-y-1/2 text-gray-500 text-2xl select-none">{{ sections.backup ? '−' : '+' }}</span>
    </div>
    <div v-show="sections.backup" class="p-6 space-y-4">
      <!-- Periods selector & restore -->
      <div v-if="backupState.periods.length" class="flex flex-wrap items-center gap-6">
        <select v-model="backupState.selected" class="border border-gray-300 rounded-lg px-3 py-2 text-sm">
          <option disabled value="">انتخاب تاریخ</option>
          <option v-for="p in backupState.periods" :key="p" :value="p">{{ p }}</option>
        </select>
        <button :disabled="backupState.restoring || !backupState.selected" class="bg-yellow-100 text-yellow-800 px-4 py-2 rounded-lg hover:bg-yellow-200 transition-all disabled:opacity-50" @click="restoreBackup">
          <span v-if="backupState.restoring">در حال بازگردانی...</span>
          <span v-else>بازگردانی</span>
        </button>
      </div>
      <!-- Report -->
      <div v-if="backupState.report" class="border-t pt-4 text-sm text-gray-700 space-y-1">
        <div>دوره: <span class="font-bold">{{ backupState.report.period }}</span></div>
        <div>تعداد فایل‌های بازگردانی‌شده: <span class="font-bold">{{ backupState.report.restored_files }}</span></div>
        <div>حجم کل بازگردانی‌شده: <span class="font-bold">{{ formatFileSize(backupState.report.total_size) }}</span></div>
      </div>
    </div>
  </div>

  <!-- مدیریت سایز تصاویر (در پایین‌ترین بخش صفحه) -->
  <div class="bg-gradient-to-l from-pink-50 via-blue-50 to-purple-50 rounded-2xl shadow-2xl border border-purple-200 mb-10 mx-4 p-10">
    <div class="flex items-center justify-between cursor-pointer px-6 py-4 border-b-2 border-purple-100" @click="toggleSection('imageSizes')">
      <span class="text-base font-bold text-purple-700 ml-auto">مدیریت سایز تصاویر</span>
      <span class="text-purple-400 text-2xl">{{ sections.imageSizes ? '−' : '+' }}</span>
    </div>
    <div v-show="sections.imageSizes">
      <form @submit.prevent="saveSizes">
        <div class="grid grid-cols-1 md:grid-cols-3 gap-8 mt-8">
          <div
              v-for="entry in orderedSizes"
              :key="entry.key"
              class="flex flex-col items-center bg-white/80 rounded-2xl shadow-lg border-2 border-blue-100 hover:border-pink-300 hover:shadow-2xl transition-all duration-200 p-6 group relative overflow-hidden"
          >
            <span class="font-bold text-right w-full mb-3 text-blue-700 group-hover:text-pink-600 transition">{{ entry.label }}</span>
            <div class="flex items-center gap-2 mb-3">
              <input
                  v-model.number="sizes[entry.key].width"
                  type="number"
                  class="w-20 px-2 py-1 border-2 border-blue-200 rounded-lg text-center text-base font-semibold focus:ring-2 focus:ring-pink-300 focus:border-pink-400 transition group-hover:border-pink-300 group-hover:bg-pink-50"
                  :min="entry.key === 'original' ? 0 : 1"
                  :placeholder="'عرض'"
              />
              <span class="text-lg text-purple-400 font-bold">×</span>
              <input
                  v-model.number="sizes[entry.key].height"
                  type="number"
                  class="w-20 px-2 py-1 border-2 border-blue-200 rounded-lg text-center text-base font-semibold focus:ring-2 focus:ring-pink-300 focus:border-pink-400 transition group-hover:border-pink-300 group-hover:bg-pink-50"
                  :min="entry.key === 'original' ? 0 : 1"
                  :placeholder="'ارتفاع'"
              />
            </div>
            <span class="text-xs text-purple-400 font-bold">px</span>
            <div class="absolute -top-3 -left-3 bg-gradient-to-tr from-pink-200 via-purple-200 to-blue-200 w-10 h-10 rounded-full blur-xl opacity-40 group-hover:opacity-70 transition-all"></div>
          </div>
        </div>
        <div class="flex justify-end mt-10">
          <button
              type="submit"
              class="bg-gradient-to-l from-pink-400 via-purple-500 to-blue-500 text-white px-10 py-3 rounded-xl font-extrabold text-lg shadow-lg hover:from-pink-500 hover:to-blue-600 hover:scale-105 transition-all duration-200"
          >
            ذخیره تغییرات
          </button>
        </div>
      </form>
    </div>
  </div>

  <!-- Toast -->
  <div v-if="toast.message" :class="['fixed bottom-6 right-6 px-4 py-2 rounded-lg shadow-lg text-sm text-white', toast.type === 'success' ? 'bg-green-500' : 'bg-red-500']">
    {{ toast.message }}
  </div>

  <!-- Compare Modal -->
  <div v-if="showCompareModal" class="fixed inset-0 bg-black bg-opacity-75 flex items-center justify-center z-50" @click="closeCompareModal">
    <div class="relative bg-white rounded-lg overflow-hidden shadow-lg w-[95vw] max-w-none max-h-[95vh] overflow-auto" @click.stop>
      <!-- Close btn -->
      <button class="absolute left-4 top-6 bg-white bg-opacity-80 rounded-full p-1 shadow hover:bg-opacity-100 transition-all" @click="closeCompareModal">
        <svg class="w-6 h-6 text-gray-700" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
        </svg>
      </button>
      <div class="flex flex-col sm:flex-row">
        <!-- Original -->
        <div class="flex-1 p-6 flex flex-col items-center justify-start border-b sm:border-b-0 sm:border-r">
          <span class="font-bold text-gray-700 mb-2">عکس اصلی</span>
          <img :src="compareModalData?.originalUrl" :alt="compareModalData?.name" class="w-full max-h-[80vh] object-contain" />
          <span v-if="compareModalData" class="text-sm text-gray-600 mt-2">{{ formatFileSize(compareModalData.originalSize) }} - {{ compareModalData.originalFormat }}</span>
        </div>
        <!-- Compressed -->
        <div class="flex-1 p-6 flex flex-col items-center justify-start">
          <span class="font-bold text-gray-700 mb-2">عکس فشرده‌شده</span>
          <img :src="compareModalData?.compressedUrl" :alt="compareModalData?.name" class="w-full max-h-[80vh] object-contain" />
          <span v-if="compareModalData" class="text-sm text-gray-600 mt-2">{{ formatFileSize(compareModalData.compressedSize) }} - {{ compareModalData.compressedFormat }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, reactive, ref, watch } from 'vue'



const canDeleteCompressionResult = computed(() => true)

interface ImageFile {
  id: number
  name: string
  url: string
  thumbnail: string
  size: number
  dimensions: string
  compressionStatus?: 'processing' | 'completed' | 'error'
  compressionResult?: {
    reduction: number
    newSize: number
  }
  // Optional preview URL for display purposes
  preview?: string
  extension?: string
  type?: string
}

interface CompressionResult {
  id: number
  name: string
  thumbnail: string
  dimensions: string
  originalSize: number
  compressedSize: number
  reduction: number
  status: 'success' | 'error' | 'processing'
  compressedUrl?: string
  outputFormat?: string
  errorMessage?: string
  originalUrl?: string
}

declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
declare const useHead: (head: { title?: string }) => void

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

useHead({
  title: 'تنظیمات تصاویر - پنل مدیریت'
})

// Reactive data
const smartCompression = computed({
  get() {
    return compressionSettings.value.quality === 'smart'
  },
  set(val: boolean) {
    // هنگام انتخاب/لغو، مقدار quality تنظیم می‌شود
    compressionSettings.value.quality = val ? 'smart' : 'medium'
  }
})
const compressionSettings = ref({
  quality: 'medium',
  customQuality: 75,
  format: 'original',
  progressive: true,
  removeMetadata: true,
  optimizeColors: true,
  createBackup: true,
  keepOriginalDate: true,
  autoOrient: true,
  enabled: true
})

// وقتی تنظیمات از سرور در حال بارگذاری است، رادیوها را غیرفعال نگه می‌داریم تا انتخاب کاربر با مقدار سرور بلافاصله بازنویسی نشود
const isLoadingCompressionSettings = ref(true)

const selectedImages = ref<number[]>([])
const isCompressing = ref(false)
const compressionResults = ref<CompressionResult[]>([])

// Compression Queue State
const compressionQueue = reactive({
  active: false,
  total: 0,
  inProgress: 0,
  completed: 0,
  errors: 0,
  progress: 0
})

const images = ref<ImageFile[]>([])
const isScanning = ref(false)
// Flag to know if کاربر دکمه اسکن را زده است
const scanPerformed = ref(false)

// Comparison Modal State
const showCompareModal = ref(false)
const compareModalData = ref<{ name: string, originalUrl: string, compressedUrl: string, originalSize: number, compressedSize: number, originalFormat: string, compressedFormat: string } | null>(null)



// Methods
const _toggleImageSelection = (imageId: number) => {
  const index = selectedImages.value.indexOf(imageId)
  if (index > -1) {
    selectedImages.value.splice(index, 1)
  } else {
    selectedImages.value.push(imageId)
  }
}

const selectAllImages = () => {
  selectedImages.value = images.value.map(img => img.id)
}

const _clearSelection = () => {
  selectedImages.value = []
}

// Update queue state after each compression
function updateCompressionQueue() {
  const total = selectedImages.value.length || images.value.length;
  const completed = images.value.filter(img => img.compressionStatus === 'completed').length;
  const inProgress = images.value.filter(img => img.compressionStatus === 'processing').length;
  const errors = images.value.filter(img => img.compressionStatus === 'error').length;
  const done = completed + errors;
  compressionQueue.total = total;
  compressionQueue.completed = completed;
  compressionQueue.inProgress = inProgress;
  compressionQueue.errors = errors;
  compressionQueue.progress = total ? Math.round((done / total) * 100) : 0;
  compressionQueue.active = inProgress > 0 || (done > 0 && done < total);
}

const startCompression = async () => {
  if (selectedImages.value.length === 0 || isCompressing.value) return
  // اگر تغییری ذخیره نشده باشد ابتدا ذخیره می‌کنیم
  if (unsavedChanges.value) {
    await saveCompressionSettings()
    // اگر هنوز ذخیره‌نشده (به دلیل خطا) متوقف شویم
    if (unsavedChanges.value) {
      showToast('ابتدا خطای ذخیره تنظیمات را برطرف کنید', 'error')
      return
    }
  }

  isCompressing.value = true
  // حذف نشود: compressionResults.value = []
  // فقط نتایج جدید اضافه یا آپدیت شوند

  // Set processing status for selected images
  selectedImages.value.forEach(imageId => {
    const img = images.value.find(i => i.id === imageId)
    if (img) img.compressionStatus = 'processing'
  })
  updateCompressionQueue()

  // آماده‌سازی پارامترهای کوئری برای کیفیت/حالت هوشمند
  const buildQuery = () => {
    const q: string[] = []
    if (compressionSettings.value.quality === 'smart') {
      q.push('smart=true')
    } else if (compressionSettings.value.quality === 'custom') {
      q.push('quality=' + compressionSettings.value.customQuality)
    } else {
      // high / medium / low → درصد متناظر
      const map: Record<string, number> = { high: 90, medium: 75, low: 60 }
      const pct = map[compressionSettings.value.quality] || 75
      q.push('quality=' + pct)
    }
    return q.length ? '?' + q.join('&') : ''
  }
  const queryStr = buildQuery()

  const headers: Record<string, string> = {}

  for (const imageId of selectedImages.value) {
    const img = images.value.find(i => i.id === imageId)
    if (!img) continue
    // Determine output format for row display
    const outFmt = compressionSettings.value.format === 'original'
        ? (img.name.split('.').pop() || '').toLowerCase()
        : compressionSettings.value.format
    // Insert or update processing row
    const existing = compressionResults.value.find(r => r.id === img.id)
    if (!existing) {
      compressionResults.value.push({
        id: img.id,
        name: img.name,
        thumbnail: img.thumbnail,
        dimensions: img.dimensions,
        originalSize: img.size,
        compressedSize: 0,
        reduction: 0,
        status: 'processing',
        outputFormat: outFmt,
        originalUrl: img.url
      })
    } else {
      existing.status = 'processing'
    }
    try {
      const res = await fetch(`/api/media/compress/${imageId}${queryStr}`, {
        method: 'POST',
        headers,
        credentials: 'include'
      })
      if (!res.ok) {
        const errText = await res.text()
        throw new Error(errText || 'server')
      }
      const data = await res.json()
      const compressedSize = data.size || img.size
      const compressedUrl = data.url || data.compressedUrl || data.path || data.compressed_path || img.url
      const reduction = Math.round(((img.size - compressedSize) / img.size) * 100)
      img.compressionStatus = 'completed'
      img.compressionResult = { reduction, newSize: compressedSize }
      // Update/replace result row
      const row = compressionResults.value.find(r => r.id === img.id)
      if (row) {
        Object.assign(row, {
          compressedSize,
          reduction,
          status: 'success',
          compressedUrl: compressedUrl,
          outputFormat: outFmt,
          originalUrl: img.url
        })
      }
    } catch (e: unknown) {
      img.compressionStatus = 'error'
      const row = compressionResults.value.find(r => r.id === img.id)
      if (row) {
        const errorMessage = e instanceof Error ? e.message : 'خطا'
        Object.assign(row, {
          status: 'error',
          errorMessage,
          outputFormat: outFmt,
          originalUrl: img.url
        })
      } else {
        compressionResults.value.push({
          id: img.id,
          name: img.name,
          thumbnail: img.thumbnail,
          dimensions: img.dimensions,
          originalSize: img.size,
          compressedSize: 0,
          reduction: 0,
          status: 'error',
          errorMessage: (e as { message?: string })?.message || 'خطا',
          outputFormat: outFmt,
          originalUrl: img.url
        })
      }
    }
    updateCompressionQueue()
  }

  isCompressing.value = false
  // Remove compressed images from list
  const toRemove = [...selectedImages.value]
  images.value = images.value.filter(i => !toRemove.includes(i.id) || i.compressionStatus === 'error')
  selectedImages.value = []
  updateCompressionQueue()
}

const _getQualityMultiplier = (): number => {
  switch (compressionSettings.value.quality) {
    case 'high': return 0.85
    case 'medium': return 0.65
    case 'low': return 0.45
    case 'custom': return (100 - compressionSettings.value.customQuality) / 100 + 0.3
    default: return 0.65
  }
}

const _getFormatMultiplier = (): number => {
  switch (compressionSettings.value.format) {
    case 'webp': return 0.7
    case 'jpeg': return 0.8
    case 'png': return 1.1
    default: return 1
  }
}

const downloadCompressed = (result: CompressionResult) => {
  // Create download link
  const link = document.createElement('a')
  link.href = result.compressedUrl || '#'
  link.download = `compressed-${result.name}`
  link.click()
}

const replaceOriginal = (result: CompressionResult) => {
  if (confirm(`آیا از جایگزینی فایل اصلی "${result.name}" اطمینان دارید؟`)) {
    // Replace original file with compressed version
    const image = images.value.find(img => img.id === result.id)
    if (image) {
      image.size = result.compressedSize
      image.url = result.compressedUrl || image.url
    }
    showToast('فایل اصلی با نسخه فشرده جایگزین شد', 'success')
  }
}

const openCompareModal = (result: CompressionResult) => {
  // Find original image url
  const _original = images.value.find(img => img.id === result.id)
  const originalFormat = (result.name.split('.').pop() || '').toUpperCase()
  const compressedFormat = (result.outputFormat || originalFormat).toUpperCase()
  compareModalData.value = {
    name: result.name,
    originalUrl: result.originalUrl || result.thumbnail,
    compressedUrl: result.compressedUrl || result.thumbnail,
    originalSize: result.originalSize || 0,
    compressedSize: result.compressedSize || 0,
    originalFormat,
    compressedFormat
  }
  showCompareModal.value = true
}

const closeCompareModal = () => {
  showCompareModal.value = false
  compareModalData.value = null
}

const compareImages = (result: CompressionResult) => {
  openCompareModal(result)
}

const formatFileSize = (bytes: number): string => {
  if (bytes === 0) return '0 B'
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

// Simulate download all and restore actions
const _downloadAllCompressed = () => {
  showToast('دانلود همه تصاویر فشرده‌شده (شبیه‌سازی)', 'success')
}

// Technical summary computed
const totalOriginalSize = computed(() => compressionResults.value.reduce((sum, r) => sum + (r.originalSize || 0), 0))
const totalCompressedSize = computed(() => compressionResults.value.reduce((sum, r) => sum + (r.compressedSize || 0), 0))
const totalSavedSize = computed(() => totalOriginalSize.value - totalCompressedSize.value)

interface DetailsModalData {
  name: string
  dimensions: string
  format: string
  originalSize: number
  compressedSize: number
  path: string
  date: string
}

// Technical details modal state
const detailsModalData = ref<DetailsModalData | null>(null)
const _showDetails = (result: CompressionResult) => {
  // Simulate technical details
  detailsModalData.value = {
    name: result.name,
    dimensions: result.dimensions,
    format: result.name.split('.').pop()?.toUpperCase() || 'JPG',
    originalSize: result.originalSize,
    compressedSize: result.compressedSize,
    path: '/images/' + result.name,
    date: '1403/03/25'
  }
}
const _closeDetailsModal = () => {
  detailsModalData.value = null
}

// Toast notification state
const toast = reactive({ message: '', type: 'success' })
function showToast(message: string, type: 'success' | 'error' = 'success') {
  toast.message = message
  toast.type = type
  setTimeout(() => { toast.message = '' }, 3000)
}

// Auto-scroll to results after compression
watch(isCompressing, (val) => {
  if (!val) {
    setTimeout(() => {
      const el = document.querySelector('.bg-white.rounded-xl.shadow-lg.border-gray-200.overflow-hidden')
      if (el) el.scrollIntoView({ behavior: 'smooth' })
    }, 300)
  }
})

// Filter & search state
const searchQuery = ref('')
const filterStatus = ref('')
const selectedResults = ref<number[]>([])
const selectAllResults = ref(false)

const showOnlyFailed = ref(false)

const filteredResults = computed(() => {
  let results = compressionResults.value
  if (searchQuery.value) {
    results = results.filter(r => r.name.toLowerCase().includes(searchQuery.value.toLowerCase()))
  }
  if (filterStatus.value) {
    results = results.filter(r => r.status === filterStatus.value)
  }
  if (showOnlyFailed.value) {
    results = results.filter(r => r.status === 'error')
  }
  return results
})

function toggleSelectAllResults() {
  if (selectAllResults.value) {
    selectedResults.value = filteredResults.value.map(r => r.id)
  } else {
    selectedResults.value = []
  }
}

function deleteSelectedResults() {
  if (selectedResults.value.length === 0) {
    showToast('هیچ نتیجه‌ای انتخاب نشده است', 'error')
    return
  }
  compressionResults.value = compressionResults.value.filter(r => !selectedResults.value.includes(r.id))
  selectedResults.value = []
  selectAllResults.value = false
  showToast('نتایج انتخاب‌شده حذف شدند', 'success')
}

function batchRestore() {
  if (selectedResults.value.length === 0) {
    showToast('هیچ نتیجه‌ای انتخاب نشده است', 'error')
    return
  }
  // Simulate restore for all selected
  selectedResults.value.forEach(id => {
    const result = compressionResults.value.find(r => r.id === id)
    if (result) {
      const image = images.value.find(img => img.id === result.id)
      if (image) {
        image.size = result.compressedSize
        image.url = result.compressedUrl || image.url
      }
    }
  })
  showToast('همه تصاویر انتخاب‌شده بازگردانی شدند', 'success')
}

const _batchCompare = () => {
  if (selectedResults.value.length === 0) {
    showToast('هیچ نتیجه‌ای انتخاب نشده است', 'error')
    return
  }
  // For demo, just show compare modal for the first selected
  const first = compressionResults.value.find(r => r.id === selectedResults.value[0])
  if (first) openCompareModal(first)
}

const _exportCSV = () => {
  // Simulate CSV export
  let csv = 'نام,ابعاد,فرمت,حجم اصلی,حجم جدید,کاهش حجم,وضعیت\n'
  compressionResults.value.forEach(r => {
    csv += `${r.name},${r.dimensions},${r.name.split('.').pop()?.toUpperCase() || ''},${r.originalSize},${r.compressedSize},${r.reduction},${r.status}\n`
  })
  const blob = new Blob([csv], { type: 'text/csv' })
  const link = document.createElement('a')
  link.href = URL.createObjectURL(blob)
  link.download = 'compression-results.csv'
  link.click()
  showToast('خروجی CSV دانلود شد', 'success')
}

// Preview section state
const previewImageFile = ref<File|null>(null)
const previewUrl = ref<string|null>(null)
const previewCompressionPercent = ref(75)
const previewFormat = ref('jpeg')
const isPreviewLoading = ref(false)

const previewImageUrl = computed(() => previewImageFile.value ? URL.createObjectURL(previewImageFile.value) : '')

const previewInfo = reactive({
  originalFormat: '',
  originalSize: 0,
  originalDimensions: '',
  outputFormat: '',
  compressedSize: 0,
  compressedDimensions: ''
})

function onPreviewFileChange(e: Event) {
  const files = (e.target as HTMLInputElement).files
  if (files && files[0]) {
    previewImageFile.value = files[0]
    // set original info
    previewInfo.originalFormat = (files[0].name.split('.').pop() || '').toUpperCase()
    previewInfo.originalSize = files[0].size
    // get dimensions
    const img = new Image()
    img.onload = () => {
      previewInfo.originalDimensions = `${img.width}x${img.height}`
    }
    img.src = URL.createObjectURL(files[0])
  }
}

async function updatePreview() {
  if (!previewImageFile.value) {
    isPreviewLoading.value = false;
    return;
  }
  // اگر آدرس قبلی وجود دارد، آن را آزاد کن تا کش مرورگر پاک شود و تصویر جدید نمایش داده شود
  if (previewUrl.value) {
    URL.revokeObjectURL(previewUrl.value);
    previewUrl.value = null; // باعث رفرش شدن کامپوننت می‌شود
  }
  const formData = new FormData();
  formData.append('file', previewImageFile.value);
  formData.append('format', previewFormat.value);
  if (smartCompression.value) {
    formData.append('smart', 'true')
  }
  // اگر فشرده‌سازی هوشمند فعال باشد کیفیت ارسالی اهمیتی ندارد، اما برای سازگاری یک مقدار پیش‌فرض می‌فرستیم
  formData.append('quality', String(smartCompression.value ? 82 : previewCompressionPercent.value));
  const res = await fetch('/api/media/preview-compress', {
    method: 'POST',
    body: formData
  });
  if (res.ok) {
    const blob = await res.blob();
    previewUrl.value = URL.createObjectURL(blob);
    // compressed info
    previewInfo.outputFormat = previewFormat.value.toUpperCase()
    previewInfo.compressedSize = blob.size
    const img2 = new Image()
    img2.onload = () => {
      previewInfo.compressedDimensions = `${img2.width}x${img2.height}`
    }
    img2.src = previewUrl.value
    isPreviewLoading.value = false;
  } else {
    previewUrl.value = null;
    isPreviewLoading.value = false;
  }
}

function onPreviewButtonClick() {
  if (!previewImageFile.value) {
    showToast('ابتدا یک تصویر انتخاب کنید', 'error')
    return
  }
  isPreviewLoading.value = true
  updatePreview()
}

// ------------------------------ Collapsible Sections ------------------------------
const defaultSections = {
  imageSEO: true,
  compressionSettings: true,
  imageSelection: true,
  compressionResults: true,
  stats: true,
  preview: true,
  backup: true,
  imageSizes: true
} as const

// reactive state (will be overwritten from localStorage in onMounted)
const sections = reactive<Record<keyof typeof defaultSections, boolean>>({ ...defaultSections })

onMounted(() => {
  try {
    const saved = localStorage.getItem('imgCompSections')
    if (saved) {
      const parsed = JSON.parse(saved)
      Object.keys(defaultSections).forEach(k => {
        if (k in parsed) {
          // @ts-ignore dynamic index
          sections[k] = !!parsed[k]
        }
      })
    }
  } catch {}
})

// save whenever user toggles
watch(sections, () => {
  try {
    localStorage.setItem('imgCompSections', JSON.stringify(sections))
  } catch {}
}, { deep: true })

function toggleSection(sectionName: keyof typeof sections) {
  sections[sectionName] = !sections[sectionName]
}

const fileInput = ref<HTMLInputElement | null>(null)
defineExpose({ fileInput })

// اگر فرمت خروجی از JPEG تغییر کند، گزینه Progressive غیرفعال/ریست می‌شود
watch(() => compressionSettings.value.format, (newVal) => {
  if (newVal !== 'jpeg' && compressionSettings.value.progressive) {
    compressionSettings.value.progressive = false
  }
})

// اگر کاربر یکی از پارامترهای پیش‌نمایش را تغییر دهد، پیش‌نمایش جدید به‌صورت خودکار درخواست می‌شود
watch([previewFormat, smartCompression, previewCompressionPercent], () => {
  if (!previewImageFile.value) return
  // ریست پیش‌نمایش قبلی
  if (previewUrl.value) {
    URL.revokeObjectURL(previewUrl.value)
    previewUrl.value = null
  }
  isPreviewLoading.value = true
  updatePreview()
})

// تغییر فرمت پیش‌نمایش فقط پیش‌نمایش فعلی را پاک می‌کند؛
// تا دوباره دکمه «پیش‌نمایش» زده نشود درخواستی ارسال نمی‌شود.
watch(previewFormat, () => {
  if (previewUrl.value) {
    URL.revokeObjectURL(previewUrl.value)
    previewUrl.value = null
  }
  isPreviewLoading.value = false
})

// After reactive data definitions add unsavedChanges ref
const unsavedChanges = ref(false)

// Watch for changes in compressionSettings (deep) to activate save button
watch(compressionSettings, () => {
  unsavedChanges.value = true
}, { deep: true })

// Function to map settings and send to backend
async function saveCompressionSettings() {
  const payload = [
    { key: 'compression.quality', value: compressionSettings.value.quality, category: 'compression', type: 'string' },
    { key: 'compression.custom_quality', value: String(compressionSettings.value.customQuality), category: 'compression', type: 'number' },
    { key: 'compression.format', value: compressionSettings.value.format, category: 'compression', type: 'string' },
    { key: 'compression.progressive', value: String(compressionSettings.value.progressive), category: 'compression', type: 'boolean' },
    { key: 'compression.remove_metadata', value: String(compressionSettings.value.removeMetadata), category: 'compression', type: 'boolean' },
    { key: 'compression.optimize_colors', value: String(compressionSettings.value.optimizeColors), category: 'compression', type: 'boolean' },
    { key: 'compression.create_backup', value: String(compressionSettings.value.createBackup), category: 'compression', type: 'boolean' },
    { key: 'compression.keep_original_date', value: String(compressionSettings.value.keepOriginalDate), category: 'compression', type: 'boolean' },
    { key: 'compression.auto_orient', value: String(compressionSettings.value.autoOrient), category: 'compression', type: 'boolean' },
    { key: 'compression.enabled', value: String(compressionSettings.value.enabled), category: 'compression', type: 'boolean' }
  ]
  try {
    const headers: Record<string, string> = { 'Content-Type': 'application/json' }
    const res = await fetch('/api/admin/settings', {
      method: 'PUT',
      credentials: 'include',
      headers,
      body: JSON.stringify(payload)
    })
    if (res.ok) {
      unsavedChanges.value = false
      lastLoadedSettings.value = { ...compressionSettings.value }
      showToast('تنظیمات با موفقیت ذخیره شد')
    } else {
      const data = await res.json()
      showToast(data?.error || 'خطا در ذخیره تنظیمات', 'error')
    }
  } catch (_err) {
    showToast('ارتباط با سرور برقرار نشد', 'error')
  }
}

// هنگام بارگذاری صفحه، تنظیمات ذخیره‌شده را از سرور دریافت می‌کنیم
onMounted(async () => {
  try {
    const headers: Record<string, string> = {}
    const res = await fetch(`/api/admin/settings?t=${Date.now()}` , {
      headers,
      credentials: 'include'
    })
    if (!res.ok) { isLoadingCompressionSettings.value = false; return }
    const raw = await res.json()
    interface SettingItem {
      key?: string
      Key?: string
      value?: string
      Value?: string
    }
    interface SettingsResponse {
      data?: SettingItem[]
    }
    const list = Array.isArray(raw) ? raw : (Array.isArray((raw as SettingsResponse)?.data) ? (raw as SettingsResponse).data : [])
    const map: Record<string, string> = {}
    // settings می‌تواند آرایه‌ای از آبجکت‌ها باشد (مستقیم یا داخل data)
    for (const s of list) {
      map[s.key || (s.Key as string)] = (s.value ?? (s.Value as string)) as string
    }

    // کمک‌کننده برای تبدیل رشته بولین به مقدار منطقی
    const toBool = (v: unknown) => {
      if (typeof v === 'boolean') return v
      return String(v).toLowerCase() === 'true' || v === '1'
    }

    if (map['compression.quality']) compressionSettings.value.quality = map['compression.quality']
    if (map['compression.custom_quality']) {
      const num = Number(map['compression.custom_quality'])
      if (!isNaN(num)) compressionSettings.value.customQuality = num
    }
    if (map['compression.format']) compressionSettings.value.format = map['compression.format']
    if (map['compression.progressive']) compressionSettings.value.progressive = toBool(map['compression.progressive'])
    if (map['compression.remove_metadata']) compressionSettings.value.removeMetadata = toBool(map['compression.remove_metadata'])
    if (map['compression.optimize_colors']) compressionSettings.value.optimizeColors = toBool(map['compression.optimize_colors'])
    if (map['compression.create_backup']) compressionSettings.value.createBackup = toBool(map['compression.create_backup'])
    if (map['compression.keep_original_date']) compressionSettings.value.keepOriginalDate = toBool(map['compression.keep_original_date'])
    if (map['compression.auto_orient']) compressionSettings.value.autoOrient = toBool(map['compression.auto_orient'])
    if (map['compression.enabled']) compressionSettings.value.enabled = toBool(map['compression.enabled'])

    lastLoadedSettings.value = { ...compressionSettings.value }
    unsavedChanges.value = false
  } catch (_e) {
    // silently fail
  }
  finally {
    isLoadingCompressionSettings.value = false
  }
})

async function scanImages() {
  if (isScanning.value) return
  scanPerformed.value = true
  try {
    isScanning.value = true
    showToast('در حال اسکن تصاویر...')
    const headers: Record<string, string> = {}
    const res = await fetch('/api/media/list', {
      headers,
      credentials: 'include'
    })
    if (!res.ok) throw new Error('server')
    const json = await res.json()
    interface MediaItem {
      id: number | string
      compressed?: boolean | number | string
      is_compressed?: boolean | number | string
      compressed_flag?: boolean | number | string
      compressed_size?: number | string
      file_name?: string
      name?: string
      url?: string
      size?: number
      width?: number
      height?: number
    }
    const list = (json.data || []) as MediaItem[]
    images.value = list.filter((m: MediaItem) => {
      // consider various representations of compressed flag
      const flag = m.compressed ?? m.is_compressed ?? m.compressed_flag
      const isCompressed = flag === true || flag === 1 || flag === '1' || String(flag).toLowerCase() === 'true'
      const hasCompressedSize = m.compressed_size && Number(m.compressed_size) > 0
      return !(isCompressed || hasCompressedSize)
    }).map(m => ({
      id: Number(m.id) || 0,
      name: m.file_name || m.name || '',
      url: m.url || '',
      thumbnail: m.url || '',
      size: Number(m.size) || 0,
      dimensions: m.width && m.height ? `${m.width}x${m.height}` : ''
    }))
    if (images.value.length > 0) {
      // فقط در صورت یافتن تصویر، Toast نمایش داده شود
      showToast(`${images.value.length} تصویر یافت شد`)
    }
  } catch (_e) {
    showToast('خطا در ارتباط با سرور', 'error')
  } finally {
    isScanning.value = false
  }
}

// Import MediaPreviewModal and add modal instance near root template bottom (before closing main div)  -- I'll append after preview section
import ImagePreviewModal from '~/components/media/ImagePreviewModal.vue'
const previewModalVisible = ref(false)
interface PreviewFile {
  id: number | string
  name: string
  url: string
  thumbnail: string
  size: number
  dimensions: string
}
const previewFile = ref<PreviewFile|null>(null)

function openImagePreview(image: ImageFile) {
  previewFile.value = {
    id: image.id,
    name: image.name,
    url: image.url,
    thumbnail: image.thumbnail,
    size: image.size,
    dimensions: image.dimensions
  }
  previewModalVisible.value = true
}

// ------------------------------ Backup Restore State ------------------------------
const backupState = reactive({
  loading: false,
  periods: [] as string[],
  selected: '',
  restoring: false,
  report: null as { period?: string; file_count?: number; total_size?: number; [key: string]: unknown } | null
})

async function loadBackupPeriods() {
  try {
    backupState.loading = true
    const headers: Record<string, string> = {}
    const res = await fetch('/api/admin/media/backup/periods', {
      headers,
      credentials: 'include'
    })
    if (!res.ok) throw new Error('server')
    const data = await res.json()
    backupState.periods = data.periods || []
    if (backupState.periods.length === 0) {
      showToast('هیچ دوره‌ای پیدا نشد', 'error')
    }
  } catch (_e) {
    showToast('خطا در دریافت دوره‌ها', 'error')
  } finally {
    backupState.loading = false
  }
}

async function restoreBackup() {
  if (!backupState.selected) return
  try {
    backupState.restoring = true
    const headers: Record<string, string> = {}
    const res = await fetch(`/api/admin/media/backup/restore/${backupState.selected}`, {
      method: 'POST',
      headers,
      credentials: 'include'
    })
    if (!res.ok) throw new Error('server')
    const data = await res.json()
    backupState.report = data
    showToast('بک‌آپ با موفقیت بازگردانی شد')
    // Refresh uncompressed images list so newly restored originals appear in manual compression section
    await scanImages()
    selectedImages.value = []
  } catch (_e) {
    showToast('خطا در بازگردانی', 'error')
  } finally {
    backupState.restoring = false
  }
}

function openPreviewUrl(url: string, title: string) {
  previewFile.value = {
    id: 0,
    name: title,
    size: 0,
    url,
    thumbnail: url,
    dimensions: ''
  }
  // اطمینان از reactive بودن
  previewFile.value = { ...previewFile.value }
  previewModalVisible.value = true
}

// در بخش <script setup lang="ts"> این متد را اضافه کن:
function handleImageSelect(imageId: number, event: MouseEvent) {
  if (event.ctrlKey || event.metaKey) {
    // multi-select
    if (selectedImages.value.includes(imageId)) {
      selectedImages.value = selectedImages.value.filter(id => id !== imageId)
    } else {
      selectedImages.value = [...selectedImages.value, imageId]
    }
  } else {
    // single-select
    selectedImages.value = [imageId]
  }
}

// Persist compressionResults in localStorage
watch(compressionResults, (val) => {
  try {
    localStorage.setItem('imgCompResults', JSON.stringify(val))
  } catch {}
}, { deep: true })

onMounted(() => {
  try {
    const saved = localStorage.getItem('imgCompResults')
    if (saved) {
      compressionResults.value = JSON.parse(saved)
    }
  } catch {}
})

function clearCompressionResults() {
  compressionResults.value = []
  try { localStorage.removeItem('imgCompResults') } catch {}
}

const lastLoadedSettings = ref({ ...compressionSettings.value })

interface SettingsObject {
  [key: string]: unknown
}
function isSettingsChanged(current: SettingsObject, last: SettingsObject) {
  return JSON.stringify(current) !== JSON.stringify(last)
}

watch(compressionSettings, (val) => {
  unsavedChanges.value = isSettingsChanged(val, lastLoadedSettings.value)
}, { deep: true })

const sizes = reactive({
  thumbnail: { label: 'بند انگشتی (thumbnail)', width: 150, height: 150 },
  small: { label: 'کوچک (small)', width: 400, height: 400 },
  medium: { label: 'متوسط (medium)', width: 800, height: 800 },
  large: { label: 'بزرگ (large)', width: 1600, height: 1600 },
  avatar: { label: 'آواتار/پروفایل (avatar)', width: 256, height: 256 },
  category: { label: 'دسته‌بندی (category)', width: 400, height: 400 },
  brand: { label: 'برند (brand)', width: 200, height: 100 },
  original: { label: 'سایز اصلی (original)', width: 0, height: 0 },
})

async function loadImageSizes() {
  try {
    const res = await fetch('/api/admin/settings', {
      credentials: 'include'
    })
    if (!res.ok) {
      return
    }
    
    const settings = await res.json()
    
    const map: Record<string, string> = {}
    // Handle the response structure: {data: [...], success: true}
    const settingsArray = settings.data || settings
    if (Array.isArray(settingsArray)) {
      for (const s of settingsArray) {
        map[s.key || s.Key] = s.value ?? s.Value
      }
    }
    
    if (map['image_sizes']) {
      const savedSizes = JSON.parse(map['image_sizes'])
      // Update the reactive sizes object with saved values
      for (const [key, value] of Object.entries(savedSizes)) {
        if (sizes[key] && typeof value === 'object' && value !== null) {
          interface SizeValue {
            width?: number | string
            height?: number | string
          }
          const sizeValue = value as SizeValue
          const newWidth = Number(sizeValue.width)
          const newHeight = Number(sizeValue.height)
          // Only update if the values are valid numbers (including 0)
          if (!isNaN(newWidth)) {
            sizes[key].width = newWidth
          }
          if (!isNaN(newHeight)) {
            sizes[key].height = newHeight
          }
        }
      }
    }
  } catch (e) {
    console.error('خطا در بارگذاری سایزهای تصاویر:', e)
  }
}

interface SizeData {
  [key: string]: { width: number; height: number }
}
async function saveSizes() {
  // Build plain object { thumbnail:{width,height}, ... }
  const data: SizeData = {}
  for (const k of Object.keys(sizes)) {
    data[k] = { width: Number(sizes[k].width)||0, height: Number(sizes[k].height)||0 }
  }

  const payload = [{ key: 'image_sizes', value: JSON.stringify(data), category: 'media', type: 'json' }]

  try {
    const headers:Record<string,string> = { 'Content-Type':'application/json' }
    const res = await fetch('/api/admin/settings', {
      method:'PUT',
      credentials:'include',
      headers,
      body: JSON.stringify(payload)
    })
    if(res.ok){
      showToast('سایزها با موفقیت ذخیره شدند!')
      // Reload sizes after saving to verify
      await loadImageSizes()
    }else{
      const j=await res.json().catch(()=>({}))
      showToast(j?.error||'خطا در ذخیره سایزها','error')
    }
  }catch(_e){
    showToast('ارتباط با سرور برقرار نشد','error')
  }
}

function clearResultsSelection() {
  selectedResults.value = []
  selectAllResults.value = false
}

function removeResult(result: CompressionResult) {
  if (confirm(`آیا از حذف این تصویر "${result.name}" اطمینان دارید؟`)) {
    // Remove the result from the list
    compressionResults.value = compressionResults.value.filter(r => r.id !== result.id)
    // Remove from selected list if present
    selectedResults.value = selectedResults.value.filter(id => id !== result.id)
    // Ensure select all checkbox reflects current state
    selectAllResults.value = selectedResults.value.length === compressionResults.value.length && compressionResults.value.length > 0
    showToast(`تصویر "${result.name}" با موفقیت حذف شد`, 'success')
  }
}

const orderedKeys = ['thumbnail','small','medium','large','category','brand','avatar','original']
const orderedSizes = computed(()=> orderedKeys.filter(k=>sizes[k]).map(k=>({ key:k, ...sizes[k] })))

// ---------------- Image SEO (AI) client logic ----------------
const imageSeo = reactive({
  enabled: true,
  generateTitle: false,
  overwriteExisting: false,
  delaySeconds: 60,
  batchSize: 10,
  model: 'gpt-4o-mini',
  lang: 'fa'
})
const availableModels = ref<{id:string,name?:string}[]>([])
const imageSeoSaved = ref(false)
interface ImageSeoJob {
  id: number
  image_id: number
  status: string
  created_at?: string
  updated_at?: string
  [key: string]: unknown
}
interface MissingImageSeo {
  id: number
  alt_text?: string
  short_description?: string
  [key: string]: unknown
}
const jobs = ref<ImageSeoJob[]>([])
const jobFilter = ref('')
const missingList = ref<MissingImageSeo[]>([])

async function loadImageSeoSettings(){
  try{
    const res = await fetch('/api/admin/settings')
    const arr = await res.json().catch(()=>[])
    const map:Record<string,string> = {}
    if(Array.isArray(arr)) for(const s of arr){ map[s.key||s.Key]=s.value??s.Value }
    if(map['image_seo.enabled']) imageSeo.enabled = ['1','true','TRUE'].includes(String(map['image_seo.enabled']))
    if(map['image_seo.generate_title']) imageSeo.generateTitle = ['1','true','TRUE'].includes(String(map['image_seo.generate_title']))
    if(map['image_seo.overwrite_existing']) imageSeo.overwriteExisting = ['1','true','TRUE'].includes(String(map['image_seo.overwrite_existing']))
    if(map['image_seo.delay_seconds']) imageSeo.delaySeconds = Number(map['image_seo.delay_seconds'])||60
    if(map['image_seo.batch_size']) imageSeo.batchSize = Number(map['image_seo.batch_size'])||10
    if(map['image_seo.model']) imageSeo.model = map['image_seo.model']
    if(map['image_seo.lang']) imageSeo.lang = map['image_seo.lang']
    // دریافت مدل‌های در دسترس و مدل پیش‌فرض سکشن از تنظیمات API
    try{
      const apiRes = await fetch('/api/admin/api-settings', { credentials: 'include' })
      interface ApiSettingsResponse {
        data?: {
          openai?: {
            available_models?: Array<{id: string; name?: string}>
          }
        }
        openai?: {
          available_models?: Array<{id: string; name?: string}>
        }
      }
      const apiJson: ApiSettingsResponse = await apiRes.json().catch(()=>null) as ApiSettingsResponse
      const openai = apiJson?.data?.openai || apiJson?.openai
      if (openai?.available_models && Array.isArray(openai.available_models)) {
        availableModels.value = openai.available_models
      }
      const sectionModel = (openai as { section_models?: { image_seo?: string }; [key: string]: unknown })?.section_models?.image_seo
      if (sectionModel) imageSeo.model = sectionModel
    }catch{}
    // در صورت نبود مقدار، روی gpt-4o-mini
    if(!imageSeo.model){ imageSeo.model = 'gpt-4o-mini' }
  }catch{}
}

async function saveImageSeoSettings(){
  imageSeoSaved.value=false
  const payload = [
    { key:'image_seo.enabled', value:String(imageSeo.enabled), category:'image_seo', type:'boolean' },
    { key:'image_seo.generate_title', value:String(imageSeo.generateTitle), category:'image_seo', type:'boolean' },
    { key:'image_seo.overwrite_existing', value:String(imageSeo.overwriteExisting), category:'image_seo', type:'boolean' },
    { key:'image_seo.delay_seconds', value:String(imageSeo.delaySeconds), category:'image_seo', type:'number' },
    { key:'image_seo.batch_size', value:String(imageSeo.batchSize), category:'image_seo', type:'number' },
    { key:'image_seo.model', value:imageSeo.model, category:'image_seo', type:'string' },
    { key:'image_seo.lang', value:imageSeo.lang, category:'image_seo', type:'string' }
  ]
  await fetch('/api/admin/settings',{ method:'PUT', credentials:'include', headers:{'Content-Type':'application/json'}, body: JSON.stringify(payload) })
  // همگام‌سازی با سکشن مدل‌ها در تنظیمات API
  try{
    // ابتدا تنظیمات فعلی را می‌گیریم تا فیلدهای الزامی (api_key و ...) حفظ شوند
    interface CurrentApiSettings {
      data?: {
        openai?: {
          section_models?: Record<string, string>
          [key: string]: unknown
        }
      }
      openai?: {
        section_models?: Record<string, string>
        [key: string]: unknown
      }
    }
    const current: CurrentApiSettings = await fetch('/api/admin/api-settings', { credentials:'include' }).then(r=>r.json()).catch(()=>null) as CurrentApiSettings
    const openai = current?.data?.openai || current?.openai || {}
    const section_models = { ...(openai.section_models||{}), image_seo: imageSeo.model }
    const body = { openai: {
      enabled: openai.enabled ?? true,
      api_key: openai.api_key || '',
      api_url: openai.api_url || 'https://api.openai.com/v1',
      default_model: openai.default_model || 'gpt-4.1-nano-2025-04-14',
      temperature: typeof openai.temperature === 'number' ? openai.temperature : 0.7,
      rate_limit: typeof openai.rate_limit === 'number' ? openai.rate_limit : 60,
      timeout: typeof openai.timeout === 'number' ? openai.timeout : 30,
      max_daily_cost: typeof openai.max_daily_cost === 'number' ? openai.max_daily_cost : 10.0,
      consuming_pages: Array.isArray(openai.consuming_pages) ? openai.consuming_pages : [],
      section_models,
      available_models: Array.isArray(openai.available_models) ? openai.available_models : []
    } }
    await fetch('/api/admin/api-settings',{
      method:'PUT',
      credentials:'include',
      headers:{'Content-Type':'application/json'},
      body: JSON.stringify(body)
    })
  }catch{}
  imageSeoSaved.value=true
}

async function loadJobs(){
  const q = jobFilter.value? ('?status='+jobFilter.value):''
  const res = await fetch('/api/admin/image-seo/jobs'+q)
  const json = await res.json().catch(()=>({}))
  jobs.value = json?.data || []
}
function formatDateTime(v: string | number | Date | null | undefined): string {
  if(!v) return ''
  try{ 
    return new Date(v).toLocaleString('fa-IR') 
  }catch{ 
    return '' 
  }
}

async function retryJob(id:number){
  await fetch('/api/admin/image-seo/retry',{ method:'POST', headers:{'Content-Type':'application/json'}, body: JSON.stringify({ job_id:id }) })
  await loadJobs()
}
async function retryForMedia(mediaId:number){
  await fetch('/api/admin/image-seo/retry',{ method:'POST', headers:{'Content-Type':'application/json'}, body: JSON.stringify({ media_id:mediaId }) })
  await loadJobs()
}
async function scanMissingMeta(){
  const res = await fetch('/api/media/list')
  const j = await res.json().catch(()=>({}))
  const arr = Array.isArray(j?.data)? j.data:[]
  missingList.value = arr.filter((m: MissingImageSeo)=> !m?.alt_text || !m?.short_description)
}

onMounted(async()=>{
  await loadImageSeoSettings();
  await loadJobs();
  await loadImageSizes();
})
</script>

<style scoped>
/* تغییر رنگ accent چک‌باکس‌ها و رادیوها به سبز */
input[type="checkbox"],
input[type="radio"] {
  accent-color: #86efac; /* pastel green */
}

/* custom tick darker green */
input[type="checkbox"] {
  position: relative;
}
input[type="checkbox"]:checked::after {
  content: "";
  position: absolute;
  inset: 0;
  background-image: url("data:image/svg+xml,%3Csvg viewBox='0 0 16 16' fill='none' stroke='%23047a48' stroke-width='3' stroke-linecap='round' stroke-linejoin='round' xmlns='http://www.w3.org/2000/svg'%3E%3Cpolyline points='3 9 6 12 13 5'/%3E%3C/svg%3E");
  background-repeat: no-repeat;
  background-position: center;
  background-size: 70% 70%;
}
</style>
