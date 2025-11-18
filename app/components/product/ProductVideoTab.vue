<template>
  <!-- افزودن ویدیو جدید -->
  <div class="bg-gradient-to-br from-purple-50 to-pink-50 rounded-2xl border border-purple-200 shadow-lg overflow-hidden mb-8">
    <div class="bg-gradient-to-r from-purple-600 to-pink-600 px-6 py-4 cursor-pointer" @click="toggleSection('videoForm')">
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-3">
          <div class="p-2 bg-white/20 rounded-xl">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
            </svg>
          </div>
          <h3 class="text-xl font-bold text-white">افزودن ویدیو جدید</h3>
        </div>
        <div class="p-2 bg-white/20 rounded-lg">
          <span class="text-white font-bold text-lg">{{ sections.videoForm ? '−' : '+' }}</span>
        </div>
      </div>
    </div>

    <div v-show="sections.videoForm" class="p-8">
      <div class="grid grid-cols-1 gapx-4 py-4">
        <!-- اطلاعات ویدیو -->
        <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
          <!-- عنوان ویدیو -->
          <div class="bg-white rounded-xl shadow-sm border border-purple-100 px-4 py-4">
            <label class="flex items-center gap-2 text-sm font-semibold text-gray-900 mb-4">
              <div class="p-2 bg-purple-100 rounded-lg">
                <svg class="w-4 h-4 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"></path>
                </svg>
              </div>
              عنوان ویدیو
            </label>
            <input 
              v-model="currentForm.title"
              type="text" 
              class="w-full px-4 py-3 bg-white border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500 text-gray-900 transition-all duration-200" 
              dir="rtl" 
              placeholder="عنوان ویدیو را وارد کنید" 
            />
          </div>

                  <!-- توضیحات ویدیو -->
        <div class="bg-white rounded-xl shadow-sm border border-blue-100 px-4 py-4">
          <label class="flex items-center gap-2 text-sm font-semibold text-gray-900 mb-4">
            <div class="p-2 bg-blue-100 rounded-lg">
              <svg class="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h7"></path>
              </svg>
            </div>
            توضیحات ویدیو
          </label>
          <textarea 
            v-model="currentForm.description"
            class="w-full px-4 py-3 bg-white border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-gray-900 transition-all duration-200 min-h-[100px] resize-none" 
            dir="rtl" 
            placeholder="توضیحات ویدیو را وارد کنید"
          ></textarea>
        </div>
        </div>



        <!-- آپلود و انتخاب فایل -->
        <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
          <!-- تصویر بندانگشتی -->
          <div class="bg-white rounded-xl shadow-sm border border-orange-100 px-4 py-4">
            <label class="flex items-center gap-2 text-sm font-semibold text-gray-900 mb-4">
              <div class="p-2 bg-orange-100 rounded-lg">
                <svg class="w-4 h-4 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                </svg>
              </div>
              تصویر بندانگشتی
            </label>
            
            <div class="border-2 border-dashed border-gray-300 rounded-lg px-4 py-4 text-center hover:border-orange-400 transition-colors cursor-pointer">
              <div class="w-16 h-16 mx-auto bg-gradient-to-br from-orange-100 to-orange-200 rounded-lg flex items-center justify-center mb-3">
                <svg class="w-6 h-6 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 00-2-2V6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                </svg>
              </div>
              <p class="text-sm text-gray-600 mb-3">تصویر بندانگشتی را آپلود کنید</p>
              <div class="flex flex-col gap-2">
                <button 
                  @click="openMediaLibrary('thumbnail')"
                  class="bg-gradient-to-r from-blue-500 to-indigo-500 text-white rounded-lg px-4 py-2 text-sm font-medium hover:from-blue-600 hover:to-indigo-600 transition-all duration-200"
                >
                  انتخاب از رسانه
                </button>
              </div>
            </div>
            <p class="text-xs text-gray-500 mt-3 text-center">ابعاد توصیه شده: 16:9 (مثلاً 1280x720)</p>
          </div>

          <!-- آپلود ویدیو مستقیم -->
          <div class="bg-white rounded-xl shadow-sm border border-indigo-100 px-4 py-4">
            <label class="flex items-center gap-2 text-sm font-semibold text-gray-900 mb-4">
              <div class="p-2 bg-indigo-100 rounded-lg">
                <svg class="w-4 h-4 text-indigo-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
                </svg>
              </div>
              آپلود ویدیو مستقیم
            </label>
            
            <!-- input type=file مخفی -->
            <input ref="videoFileInput" type="file" accept="video/*" class="hidden" @change="handleVideoFileChange" />
            
            <div class="border-2 border-dashed border-gray-300 rounded-lg px-4 py-4 text-center hover:border-indigo-400 transition-colors cursor-pointer">
              <div class="w-16 h-16 mx-auto bg-gradient-to-br from-indigo-100 to-indigo-200 rounded-lg flex items-center justify-center mb-3">
                <svg class="w-6 h-6 text-indigo-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
                </svg>
              </div>
              <p class="text-sm text-gray-600 mb-3">فایل ویدیو را آپلود کنید</p>
              <div class="flex flex-col gap-2">
                <button 
                  @click="openMediaLibrary('video')"
                  class="bg-gradient-to-r from-green-500 to-teal-500 text-white rounded-lg px-4 py-2 text-sm font-medium hover:from-green-600 hover:to-teal-600 transition-all duration-200"
                >
                  انتخاب از رسانه
                </button>
              </div>
            </div>
            <div v-if="uploadingVideo" class="mt-2 text-xs text-blue-600">در حال آپلود ویدیو...</div>
            <div v-if="uploadError" class="mt-2 text-xs text-red-600">{{ uploadError }}</div>
            <div v-if="currentForm.video_url" class="mt-2 text-xs text-green-600 break-all">لینک ویدیو: {{ currentForm.video_url }}</div>
            <div class="mt-3 text-xs text-gray-500 text-center space-y-1">
              <p>فرمت‌های مجاز: MP4, AVI, MOV, WebM</p>
              <p>حداکثر حجم: 100MB</p>
            </div>
          </div>
        </div>

        <!-- Video Preview Section -->
        <div v-if="currentForm.video_url" class="bg-white rounded-xl shadow-sm border border-green-100 px-4 py-4">
          <label class="flex items-center gap-2 text-sm font-semibold text-gray-900 mb-4">
            <div class="p-2 bg-green-100 rounded-lg">
              <svg class="w-4 h-4 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
              </svg>
            </div>
            پیش‌نمایش ویدیو
          </label>

          <div class="bg-gray-50 rounded-lg px-4 py-4">
            <ProductVideoPlayer 
              :video-url="currentForm.video_url"
              :poster-image="currentForm.thumbnail_url"
              :title="currentForm.title || 'ویدیو جدید'"
              :description="currentForm.description"
              :autoplay="false"
              :show-controls="true"
              :show-in-gallery="true"
            />
          </div>
        </div>

        <!-- تنظیمات نمایش -->
        <div class="bg-white rounded-xl shadow-sm border border-slate-100 px-4 py-4">
          <h4 class="text-lg font-semibold text-gray-800 mb-6">تنظیمات نمایش ویدیو</h4>
          <div class="grid grid-cols-1 md:grid-cols-3 gapx-4 py-4">
            <label class="flex items-center gap-3 px-4 py-4 rounded-lg border border-gray-200 hover:bg-gray-50 cursor-pointer transition-all duration-200">
              <input 
                v-model="currentForm.show_in_gallery"
                type="checkbox" 
                class="w-5 h-5 text-purple-600 bg-gray-100 border-gray-300 rounded focus:ring-purple-500 focus:ring-2" 
              />
              <div class="flex items-center gap-2">
                <div class="p-1 bg-purple-100 rounded">
                  <svg class="w-3 h-3 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                  </svg>
                </div>
                <span class="text-sm text-gray-700 font-medium">نمایش در گالری</span>
              </div>
            </label>

            <label class="flex items-center gap-3 px-4 py-4 rounded-lg border border-gray-200 hover:bg-gray-50 cursor-pointer transition-all duration-200">
              <input 
                v-model="currentForm.autoplay"
                type="checkbox" 
                class="w-5 h-5 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 focus:ring-2" 
              />
              <div class="flex items-center gap-2">
                <div class="p-1 bg-blue-100 rounded">
                  <svg class="w-3 h-3 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.828 14.828a4 4 0 01-5.656 0M9 10h1.586a1 1 0 01.707.293l2.414 2.414a1 1 0 00.707.293H15M9 10V9a2 2 0 012-2h2a2 2 0 012 2v1M9 10v5a2 2 0 002 2h2a2 2 0 002-2v-5"></path>
                  </svg>
                </div>
                <span class="text-sm text-gray-700 font-medium">پخش خودکار</span>
              </div>
            </label>

            <label class="flex items-center gap-3 px-4 py-4 rounded-lg border border-gray-200 hover:bg-gray-50 cursor-pointer transition-all duration-200">
              <input 
                v-model="currentForm.show_controls"
                type="checkbox" 
                class="w-5 h-5 text-green-600 bg-gray-100 border-gray-300 rounded focus:ring-green-500 focus:ring-2" 
              />
              <div class="flex items-center gap-2">
                <div class="p-1 bg-green-100 rounded">
                  <svg class="w-3 h-3 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6V4m0 2a2 2 0 100 4m0-4a2 2 0 110 4m-6 8a2 2 0 100-4m0 4a2 2 0 100 4m0-4v2m0-6V4m6 6v10m6-2a2 2 0 100-4m0 4a2 2 0 100 4m0-4v2m0-6V4"></path>
                  </svg>
                </div>
                <span class="text-sm text-gray-700 font-medium">نمایش کنترل‌ها</span>
              </div>
            </label>
          </div>
        </div>

        <!-- دکمه ذخیره -->
        <div class="flex justify-end gap-3">
          <button 
            v-if="isEditing"
            @click="cancelEdit"
            class="bg-gray-500 text-white rounded-lg px-6 py-3 font-semibold hover:bg-gray-600 transition-all duration-200 shadow-md hover:shadow-lg"
          >
            انصراف
          </button>
          <button 
            @click="isEditing ? handleUpdateVideo() : handleAddVideo()"
            :disabled="isLoading"
            class="bg-gradient-to-r from-purple-600 to-pink-600 text-white rounded-lg px-8 py-3 font-semibold hover:from-purple-700 hover:to-pink-700 transition-all duration-200 shadow-md hover:shadow-lg disabled:opacity-50 disabled:cursor-not-allowed"
          >
            <span v-if="isLoading">در حال پردازش...</span>
            <span v-else>{{ isEditing ? 'ویرایش ویدیو' : 'افزودن ویدیو' }}</span>
          </button>
        </div>
      </div>
    </div>
  </div>

  <!-- لیست ویدیوها -->
  <div class="bg-gradient-to-br from-blue-50 to-indigo-50 rounded-2xl border border-blue-200 shadow-lg overflow-hidden mb-8">
    <div class="bg-gradient-to-r from-blue-600 to-indigo-600 px-6 py-4 cursor-pointer" @click="toggleSection('videoList')">
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-3">
          <div class="p-2 bg-white/20 rounded-xl">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              
            </svg>
          </div>
          <h3 class="text-xl font-bold text-white">مدیریت ویدیوهای محصول</h3>
        </div>
        <div class="p-2 bg-white/20 rounded-lg">
          <span class="text-white font-bold text-lg">{{ sections.videoList ? '−' : '+' }}</span>
        </div>
      </div>
    </div>

    <div v-show="sections.videoList" class="p-8">
      
      <div class="space-y-4">
        <!-- Loading State -->
        <div v-if="isLoading" class="bg-white rounded-xl p-12 text-center">
          <div class="animate-spin w-8 h-8 border-4 border-blue-600 border-t-transparent rounded-full mx-auto mb-4"></div>
          <p class="text-gray-600">در حال بارگذاری ویدیوها...</p>
        </div>

        <!-- Error State -->
        <div v-else-if="error" class="bg-red-50 border border-red-200 rounded-xl px-4 py-4">
          <div class="flex items-center gap-3">
            <svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
            <p class="text-red-800">{{ error }}</p>
            <button @click="clearError" class="text-red-600 hover:text-red-800">×</button>
          </div>
        </div>

        <!-- لیست ویدیوها -->
        <div v-else class="space-y-4">
          <div 
            v-for="video in sortedVideos" 
            :key="video.id"
            class="bg-white rounded-xl shadow-sm border border-gray-100 px-4 py-4"
          >
            <!-- Video Player Section -->
            <div class="mb-4 relative">
              <ProductVideoPlayer 
                :video-url="video.video_url"
                :poster-image="video.thumbnail_url"
                :title="video.title"
                :description="video.description"
                :autoplay="video.autoplay"
                :show-controls="video.show_controls"
                :show-in-gallery="video.show_in_gallery"
              />
              <!-- overlay clickable layer to capture clicks even over iframes -->
              <button type="button" class="absolute inset-0" aria-label="پیش‌نمایش ویدیو" @click="openPreview(video)"></button>
            </div>
            
            <div class="flex items-center gapx-4 py-4">
              <!-- Thumbnail -->
              <div class="w-24 h-16 bg-gradient-to-br from-gray-100 to-gray-200 rounded-lg flex items-center justify-center flex-shrink-0 overflow-hidden">
                <img 
                  v-if="video.thumbnail_url" 
                  :src="video.thumbnail_url" 
                  :alt="video.title"
                  class="w-full h-full object-cover"
                />
                <svg v-else class="w-6 h-6 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
                </svg>
              </div>
              
              <!-- Video Info -->
              <div class="flex-1">
                <h5 class="font-semibold text-gray-900 mb-1">{{ video.title || 'بدون عنوان' }}</h5>
                <p class="text-sm text-gray-600 mb-2">{{ video.description || 'بدون توضیحات' }}</p>
                <div class="flex items-center gapx-4 py-4 text-xs text-gray-500">
                  <span>اولویت: {{ video.sort_order + 1 }}</span>
                  <span v-if="video.show_in_gallery" class="text-green-600">نمایش در گالری</span>
                  <span v-if="video.autoplay" class="text-blue-600">پخش خودکار</span>
                  <span v-if="video.show_controls" class="text-purple-600">کنترل‌ها فعال</span>
                </div>
              </div>
              
              <!-- Actions -->
              <div class="flex items-center gap-2">
                <button 
                  @click="openPreview(video)"
                  class="p-2 text-gray-600 hover:text-gray-800 hover:bg-gray-50 rounded-lg transition-colors"
                  title="پیش‌نمایش ویدیو"
                >
                  <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-4 h-4">
                    <path d="M12 5c-7 0-11 7-11 7s4 7 11 7 11-7 11-7-4-7-11-7zm0 12a5 5 0 110-10 5 5 0 010 10zm0-2a3 3 0 100-6 3 3 0 000 6z"/>
                  </svg>
                </button>
                <button 
                  @click="startEditVideo(video)"
                  class="p-2 text-blue-600 hover:text-blue-800 hover:bg-blue-50 rounded-lg transition-colors"
                  title="ویرایش ویدیو"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
                  </svg>
                </button>
                <button 
                  @click="handleDeleteVideo(video.id!)"
                  class="p-2 text-red-600 hover:text-red-800 hover:bg-red-50 rounded-lg transition-colors"
                  title="حذف ویدیو"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- ویدیوهای آپارات -->
  <div class="bg-gradient-to-br from-orange-50 to-red-50 rounded-2xl border border-orange-200 shadow-lg overflow-hidden">
    <div class="bg-gradient-to-r from-orange-600 to-red-600 px-6 py-4 cursor-pointer" @click="toggleSection('aparatVideos')">
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-3">
          <div class="p-2 bg-white/20 rounded-xl">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
            </svg>
          </div>
          <h3 class="text-xl font-bold text-white">ویدیوهای آپارات</h3>
          <span class="bg-white/20 text-white px-3 py-1 rounded-full text-sm font-medium">اختصاصی</span>
        </div>
        <div class="p-2 bg-white/20 rounded-lg">
          <span class="text-white font-bold text-lg">{{ sections.aparatVideos ? '−' : '+' }}</span>
        </div>
      </div>
    </div>

    <div v-show="sections.aparatVideos" class="p-8">
      <div class="grid grid-cols-1 gapx-4 py-4">
        <!-- دکمه افزودن -->
        <div class="flex justify-between items-center">
          <button class="bg-gradient-to-r from-orange-600 to-red-600 text-white rounded-lg px-6 py-3 font-semibold hover:from-orange-700 hover:to-red-700 transition-all duration-200 shadow-md hover:shadow-lg" onclick="addAparatVideo()">
            <svg class="w-5 h-5 inline-block mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
            </svg>
            افزودن ویدیو آپارات
          </button>
        </div>

        <!-- لیست ویدیوهای آپارات -->
        <div id="aparatVideosList" class="space-y-4">
          <!-- آیتم نمونه ویدیو آپارات -->
          <div class="bg-white rounded-xl shadow-sm border border-orange-100 px-4 py-4">
            <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
              <div class="space-y-4">
                <div>
                  <label class="flex items-center gap-2 text-sm font-semibold text-gray-900 mb-3">
                    <div class="p-1 bg-orange-100 rounded">
                      <svg class="w-3 h-3 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1"></path>
                      </svg>
                    </div>
                    لینک آپارات
                  </label>
                  <input type="url" class="w-full px-4 py-3 bg-white border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-orange-500 text-gray-900 transition-all duration-200" dir="ltr" placeholder="https://www.aparat.com/v/..." />
                </div>
              </div>
              
              <div class="space-y-4">
                <div>
                  <label class="flex items-center gap-2 text-sm font-semibold text-gray-900 mb-3">
                    <div class="p-1 bg-red-100 rounded">
                      <svg class="w-3 h-3 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 20l4-16m2 16l4-16M6 9h14M4 15h14"></path>
                      </svg>
                    </div>
                    یا شناسه ویدیو
                  </label>
                  <input type="text" class="w-full px-4 py-3 bg-white border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500 text-gray-900 transition-all duration-200" dir="ltr" placeholder="abc123def" />
                </div>
              </div>
            </div>
            
            <div class="flex justify-end mt-4">
              <button class="bg-red-500 text-white rounded-lg px-4 py-2 text-sm font-medium hover:bg-red-600 transition-colors" onclick="removeAparatVideo(this)">
                <svg class="w-4 h-4 inline-block mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                </svg>
                حذف ویدیو
              </button>
            </div>
          </div>
        </div>

        <!-- راهنما -->
        <div class="bg-gradient-to-br from-orange-100 to-red-100 border border-orange-300 rounded-xl px-4 py-4">
          <div class="flex items-start gap-3">
            <div class="p-2 bg-orange-200 rounded-lg flex-shrink-0">
              <svg class="w-5 h-5 text-orange-700" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
            <div class="flex-1">
              <h4 class="text-sm font-semibold text-orange-800 mb-3">راهنمای استفاده از آپارات</h4>
              <div class="space-y-2 text-xs text-orange-700">
                <div class="flex items-center gap-2">
                  <div class="w-1.5 h-1.5 bg-orange-600 rounded-full"></div>
                  <span>لینک کامل ویدیو آپارات را کپی کرده و در فیلد "لینک آپارات" قرار دهید</span>
                </div>
                <div class="flex items-center gap-2">
                  <div class="w-1.5 h-1.5 bg-orange-600 rounded-full"></div>
                  <span>یا فقط شناسه ویدیو (بخش انتهایی URL بعد از /v/) را در فیلد "شناسه ویدیو" وارد کنید</span>
                </div>
                <div class="flex items-center gap-2">
                  <div class="w-1.5 h-1.5 bg-orange-600 rounded-full"></div>
                  <span><strong>مثال لینک:</strong> https://www.aparat.com/v/abc123def</span>
                </div>
                <div class="flex items-center gap-2">
                  <div class="w-1.5 h-1.5 bg-orange-600 rounded-full"></div>
                  <span><strong>مثال شناسه:</strong> abc123def</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- Media Library Modal -->
  <!-- 
    برای ویدیوها: category="products" و file-type="video" 
    این باعث می‌شود که ویدیوها در پوشه products/videos ذخیره شوند
    برای تصاویر: category="products" و file-type="image"
    این باعث می‌شود که تصاویر در پوشه products/images ذخیره شوند
  -->
  <MediaLibraryModal 
    v-if="showMediaLibrary"
    v-model="showMediaLibrary"
    :default-category="'products'"
    :file-type="mediaLibraryType === 'video' ? 'video' : 'image'"
    @confirm="onMediaSelected"
  />

  <!-- Preview Modal -->
  <MediaPreviewModal v-if="previewShow" v-model="previewShow" :file="previewFile" :can-delete="false" />
  <!-- Local delete confirm modal instance -->
  <DeleteConfirmModal ref="deleteModalRef" @single-delete="onSingleDelete" />
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed, watch, nextTick } from 'vue'
import DeleteConfirmModal from '~/components/common/DeleteConfirmModal.vue'
import { useDeleteConfirmModal } from '~/composables/useDeleteConfirmModal'
import { useNotifier } from '~/composables/useNotifier'
import { useProductVideos, type ProductVideo } from '~/composables/useProductVideos'
import { useProductCreateStore } from '~/stores/productCreate'
import ProductVideoPlayer from '~/components/product/ProductVideo.vue'
import MediaPreviewModal from '~/components/media/MediaPreviewModal.vue'
import MediaLibraryModal from '~/components/media/MediaLibraryModal.vue'

// استفاده از composable
const {
  videos,
  isLoading,
  error,
  currentProductId,
  fetchVideos,
  addVideo,
  updateVideo,
  deleteVideo,
  clearError
} = useProductVideos()

// DeleteConfirm modal controller
const { deleteModalRef, openDeleteConfirm } = useDeleteConfirmModal()

// استفاده مستقیم از store
const pStore = useProductCreateStore()
console.log('🎬 ProductVideoTab - pStore initialized:', { 
  pStore, 
  pStoreType: typeof pStore,
  editingProductId: pStore.editingProductId,
  isEditMode: pStore.isEditMode
})

// State management
const sections = reactive({
  videoForm: false,
  videoList: false,
  aparatVideos: false
})

// فرم افزودن ویدیو
const videoForm = reactive({
  title: '',
  description: '',
  video_url: '',
  thumbnail_url: '',
  show_in_gallery: true,
  autoplay: false,
  show_controls: true,
  sort_order: 0
})

// فرم ویرایش ویدیو
const editingVideo = ref<ProductVideo | null>(null)
const isEditing = ref(false)

// Computed برای فرم فعلی
const currentForm = computed(() => {
  return isEditing.value ? editingVideo.value! : videoForm
})

// Media Library Modal
const showMediaLibrary = ref(false)
const mediaLibraryType = ref<'video' | 'thumbnail'>('video')

// Computed properties
const sortedVideos = computed(() => {
  return [...videos.value].sort((a, b) => a.sort_order - b.sort_order)
})

const hasVideos = computed(() => videos.value.length > 0)

// Methods
const toggleSection = (section: keyof typeof sections) => {
  sections[section] = !sections[section]
}

const resetVideoForm = () => {
  videoForm.title = ''
  videoForm.description = ''
  videoForm.video_url = ''
  videoForm.thumbnail_url = ''
  videoForm.show_in_gallery = true
  videoForm.autoplay = false
  videoForm.show_controls = true
  videoForm.sort_order = videos.value.length
}

const openMediaLibrary = (type: 'video' | 'thumbnail') => {
  mediaLibraryType.value = type
  // برای ویدیوها، اطمینان حاصل می‌کنیم که فقط ویدیوها نمایش داده شوند
  // و در پوشه products/videos ذخیره شوند
  showMediaLibrary.value = true
}

const onMediaSelected = async (payload: any) => {
  const media = Array.isArray(payload) ? payload[0] : payload
  if (!media) {
    showMediaLibrary.value = false
    return
  }

  console.log('🎬 Media selected:', media)
  console.log('📹 Media type:', mediaLibraryType.value)
  console.log('✏️ Is editing:', isEditing.value)

  let url = media.url || media.path || media.file_path || media.src || ''
  // normalize path for public serving
  url = (url || '').toString().replace(/\\\\/g,'/').replace(/\/public\//,'/')
  if (!/^https?:\/\//i.test(url)) { if (!url.startsWith('/')) url = '/' + url }
  // map bare /products/* to /uploads/media/products/* (actual public path)
  if (!/\/uploads\//.test(url) && /^\/products\//.test(url)) {
    url = '/uploads/media' + url
  }
  if (!url) {
    showMediaLibrary.value = false
    return
  }

  if (mediaLibraryType.value === 'video') {
    if (isEditing.value && editingVideo.value) {
      editingVideo.value.video_url = url
      console.log('✅ Set video_url on editingVideo:', url)
      // ذخیره فوری در حالت ویرایش
      try {
        await updateVideo(editingVideo.value.id!, {
          title: editingVideo.value.title,
          description: editingVideo.value.description,
          video_url: editingVideo.value.video_url,
          thumbnail_url: editingVideo.value.thumbnail_url,
          show_in_gallery: editingVideo.value.show_in_gallery,
          autoplay: editingVideo.value.autoplay,
          show_controls: editingVideo.value.show_controls,
          sort_order: editingVideo.value.sort_order
        })
        await fetchVideos()
      } catch {}
    } else {
      // در حالت افزودن: پس از انتخاب رسانه، فوراً رکورد را ایجاد کن
      videoForm.video_url = url
      console.log('✅ Set video_url on videoForm:', url)
      try {
        const ok = await addVideo({
          title: videoForm.title,
          description: videoForm.description,
          video_url: videoForm.video_url,
          thumbnail_url: videoForm.thumbnail_url,
          show_in_gallery: videoForm.show_in_gallery,
          autoplay: videoForm.autoplay,
          show_controls: videoForm.show_controls,
          sort_order: videos.value.length
        })
        if (ok) {
          resetVideoForm()
          await nextTick()
          await fetchVideos()
        }
      } catch {}
    }
  } else {
    if (isEditing.value && editingVideo.value) {
      editingVideo.value.thumbnail_url = url
      console.log('✅ Set thumbnail_url on editingVideo:', url)
    } else {
      videoForm.thumbnail_url = url
      console.log('✅ Set thumbnail_url on videoForm:', url)
    }
  }

  console.log('📊 Current form state:')
  console.log('  - videoForm.video_url:', videoForm.video_url)
  console.log('  - editingVideo?.video_url:', editingVideo.value?.video_url)
  console.log('  - currentForm.video_url:', currentForm.value.video_url)

  showMediaLibrary.value = false
}

const handleAddVideo = async () => {
  if (!videoForm.video_url) {
    const notifier = useNotifier()
    notifier.warning('لینک ویدیو الزامی است')
    return
  }

  const success = await addVideo({
    title: videoForm.title,
    description: videoForm.description,
    video_url: videoForm.video_url,
    thumbnail_url: videoForm.thumbnail_url,
    show_in_gallery: videoForm.show_in_gallery,
    autoplay: videoForm.autoplay,
    show_controls: videoForm.show_controls,
    sort_order: videoForm.sort_order
  })

  if (success) {
    resetVideoForm()
    sections.videoForm = false
    const notifier = useNotifier()
    notifier.success('ویدیو با موفقیت اضافه شد')
  }
}

const startEditVideo = (video: ProductVideo) => {
  editingVideo.value = { ...video }
  isEditing.value = true
  sections.videoForm = true
}

const handleUpdateVideo = async () => {
  if (!editingVideo.value?.id) return

  const success = await updateVideo(editingVideo.value.id, {
    title: editingVideo.value.title,
    description: editingVideo.value.description,
    video_url: editingVideo.value.video_url,
    thumbnail_url: editingVideo.value.thumbnail_url,
    show_in_gallery: editingVideo.value.show_in_gallery,
    autoplay: editingVideo.value.autoplay,
    show_controls: editingVideo.value.show_controls,
    sort_order: editingVideo.value.sort_order
  })

  if (success) {
    editingVideo.value = null
    isEditing.value = false
    sections.videoForm = false
    const notifier = useNotifier()
    notifier.success('ویدیو با موفقیت ویرایش شد')
  }
}

const pendingDeleteId = ref<number | null>(null)
const handleDeleteVideo = (videoId: number) => {
  pendingDeleteId.value = videoId
  openDeleteConfirm(videoId)
}

async function onSingleDelete(id: number | string) {
  const success = await deleteVideo(Number(id))
  if (success) useNotifier().success('ویدیو با موفقیت حذف شد')
}

const cancelEdit = () => {
  editingVideo.value = null
  isEditing.value = false
  resetVideoForm()
}

const videoFileInput = ref<HTMLInputElement | null>(null)
const uploadingVideo = ref(false)
const uploadError = ref('')

function triggerVideoFileInput() {
  videoFileInput.value?.click()
}

async function handleVideoFileChange(e: Event) {
  const target = e.target as HTMLInputElement
  const files = target.files
  if (!files || files.length === 0) return
  const file = files[0]
  if (!file) return
  uploadingVideo.value = true
  uploadError.value = ''
  try {
    const formData = new FormData()
    formData.append('video', file)
    formData.append('category', 'products')
    // فرض بر این است که API آپلود ویدیو در /api/media/upload وجود دارد
    const res = await fetch('/api/media/upload', {
      method: 'POST',
      body: formData
    })
    const data = await res.json()
    if (data.success && (data.url || data.video_url)) {
      currentForm.value.video_url = data.url || data.video_url
    } else {
      uploadError.value = data.message || 'خطا در آپلود ویدیو'
    }
  } catch (err: any) {
    uploadError.value = err.message || 'خطا در آپلود ویدیو'
  } finally {
    uploadingVideo.value = false
  }
}

// Lifecycle
onMounted(async () => {
  if (currentProductId.value) {
    await fetchVideos()
  }
})

// واکنش به آماده‌شدن شناسه محصول برای بارگذاری لیست ویدیوها
watch(currentProductId, async (newId, oldId) => {
  if (newId && newId !== oldId) {
    await fetchVideos(Number(newId))
  }
})

// Preview modal state and opener
const previewShow = ref(false)
const previewFile = ref<any>(null)
function openPreview(video: ProductVideo) {
  previewFile.value = {
    id: video.id,
    url: video.video_url,
    preview: video.video_url,
    name: video.title || 'video',
    type: 'video',
    mime_type: 'video/mp4',
    caption: video.description || ''
  }
  previewShow.value = true
}
</script>

<style scoped>
/* Custom Select Arrow */
select {
  background-image: url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' fill='none' viewBox='0 0 20 20'%3e%3cpath stroke='%236b7280' stroke-linecap='round' stroke-linejoin='round' stroke-width='1.5' d='m6 8 4 4 4-4'/%3e%3c/svg%3e");
  background-position: right 0.5rem center;
  background-repeat: no-repeat;
  background-size: 1.5em 1.5em;
  padding-right: 2.5rem;
}

/* Hover transitions */
.transition-all {
  transition: all 0.3s ease;
}

/* Custom focus ring */
.focus\:ring-2:focus {
  --tw-ring-offset-shadow: var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color);
  --tw-ring-shadow: var(--tw-ring-inset) 0 0 0 calc(2px + var(--tw-ring-offset-width)) var(--tw-ring-color);
  box-shadow: var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000);
}

/* Upload areas hover effects */
.border-dashed:hover {
  border-style: solid;
  transform: translateY(-1px);
}
</style>
