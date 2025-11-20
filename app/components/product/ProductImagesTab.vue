  <template>
   <!-- تصاویر اصلی محصول -->
   <div class="bg-gradient-to-r from-blue-50 to-cyan-50 border border-blue-200 rounded-xl shadow-lg px-4 py-4 text-right mb-6 transition-all duration-300 hover:shadow-xl">
     <div class="flex items-center justify-between cursor-pointer" @click="toggleSection('mainImages')">
       <div class="flex items-center gap-3">
         <div class="w-10 h-10 bg-gradient-to-r from-blue-600 to-cyan-600 rounded-lg flex items-center justify-center shadow-md">
           <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
             <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
           </svg>
         </div>
         <h3 class="text-lg font-bold text-blue-800">آپلود و مدیریت تصاویر اصلی</h3>
       </div>
       <span class="text-blue-600 text-xl font-bold bg-white rounded-full w-8 h-8 flex items-center justify-center shadow-md">{{ sections.mainImages ? '−' : '+' }}</span>
     </div>

     <div v-show="sections.mainImages" class="mt-6">
       <div class="grid grid-cols-1 gapx-4 py-4">
         <!-- آپلود تصاویر -->
         <div
           ref="dropzone"
           class="border-2 border-dashed border-blue-300 bg-gradient-to-br from-blue-50 to-cyan-50 rounded-xl p-8 text-center transition-all duration-300 hover:border-blue-400 hover:shadow-md"
           @dragover.prevent
           @drop.prevent="onDropFiles"
         >
           <div class="text-blue-600">
             <div class="text-6xl mb-4">
               <svg class="w-16 h-16 mx-auto text-blue-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                 <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"></path>
               </svg>
             </div>
             <div class="text-lg font-semibold text-blue-800 mb-2">تصاویر محصول را بکشید و رها کنید</div>
             <div class="text-sm text-blue-600 mb-4">یا روی دکمه‌های زیر کلیک کنید</div>
             <div class="text-xs text-blue-500 bg-blue-100 rounded-lg px-4 py-2 inline-block">
               فرمت‌های مجاز: JPG, PNG, WebP • حداکثر 5 مگابایت
             </div>
           </div>
           <input ref="fileInput" type="file" class="hidden" multiple accept="image/*" @change="onFileInput" />
           
           <!-- دکمه‌های عملیات -->
            <div class="flex flex-wrap justify-center gapx-4 py-4 mt-8">
              <!-- انتخاب از رسانه -->
             <div class="flex flex-col items-center gap-3">
               <button class="bg-gradient-to-r from-green-600 to-emerald-600 text-white rounded-xl px-6 py-3 font-semibold shadow-lg hover:from-green-700 hover:to-emerald-700 transition-all duration-300 flex items-center gap-2" @click.prevent="openMedia">
                 <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                   
                 </svg>
                 انتخاب از کتابخانه رسانه
               </button>
             </div>
           </div>
         </div>
       </div>
     </div>
   </div>

   <!-- گالری تصاویر -->
   <div class="bg-gradient-to-r from-emerald-50 to-teal-50 border border-emerald-200 rounded-xl shadow-lg px-4 py-4 text-right mb-6 transition-all duration-300 hover:shadow-xl">
     <div class="flex items-center justify-between cursor-pointer" @click="toggleSection('galleryList')">
       <div class="flex items-center gap-3">
         <div class="w-10 h-10 bg-gradient-to-r from-emerald-600 to-teal-600 rounded-lg flex items-center justify-center shadow-md">
           <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
             
           </svg>
         </div>
         <h3 class="text-lg font-bold text-emerald-800">گالری و مدیریت تصاویر</h3>
       </div>
       <span class="text-emerald-600 text-xl font-bold bg-white rounded-full w-8 h-8 flex items-center justify-center shadow-md">{{ sections.galleryList ? '−' : '+' }}</span>
     </div>

     <div v-show="sections.galleryList" class="mt-6">
       <div v-if="store.images.length" class="space-y-4">
         <!-- آمار سریع -->
         <div class="bg-white rounded-xl px-4 py-4 shadow-md border border-emerald-200">
           <div class="flex items-center justify-between">
             <div class="flex items-center gap-3">
               <div class="w-8 h-8 bg-emerald-100 rounded-lg flex items-center justify-center">
                 <svg class="w-4 h-4 text-emerald-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                   <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v4"></path>
                 </svg>
               </div>
               <span class="text-sm font-semibold text-emerald-800">مجموع تصاویر: {{ store.images.length }} عدد</span>
             </div>
             <div class="text-xs text-emerald-600 bg-emerald-100 px-3 py-1 rounded-full">
               تصاویر محصول
             </div>
           </div>
         </div>

          <!-- header جدول -->
          <div class="bg-gradient-to-r from-emerald-100 to-teal-100 rounded-xl px-4 py-4 shadow-md">
            <div class="grid grid-cols-12 gapx-4 py-4 text-sm font-bold text-emerald-800" dir="rtl">
              <span class="col-span-1 text-center">جابجایی</span>
              <span class="col-span-2 text-center">تصویر</span>
              <span class="col-span-1 text-center">ترتیب</span>
              <span class="col-span-4 text-right">عنوان تصویر</span>
              <span class="col-span-3 text-right">متن جایگزین (Alt)</span>
              <span class="col-span-1 text-center">عملیات</span>
            </div>
          </div>

         <!-- لیست تصاویر -->
         <div class="space-y-3">
           <div
             v-for="(img, idx) in store.images"
             :key="img.id"
             class="bg-white rounded-xl px-4 py-4 shadow-md border border-emerald-200 hover:shadow-lg transition-all duration-300 relative"
             draggable="true"
             @dragstart="onDragStart(idx)"
             @dragover.prevent
             @drop="onDrop(idx)"
           >
             <div class="grid grid-cols-12 gapx-4 py-4 items-center" dir="rtl">
               <!-- drag handle -->
               <div class="col-span-1 flex justify-center">
                 <div class="cursor-grab text-emerald-400 hover:text-emerald-600 transition-colors duration-300">
                   <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
                     <path d="M3 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 10a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 16a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z"></path>
                   </svg>
                 </div>
               </div>
               
               <!-- تصویر -->
               <div class="col-span-2 flex justify-center">
                 <div class="w-20 h-20 bg-white rounded-xl border-2 border-gray-300 overflow-hidden">
                   <!-- تصویر ساده -->
                   <img 
                     v-if="img.thumbnail || img.url"
                     :src="(img.thumbnail || img.url) as string" 
                     :alt="(img.alt as string) || 'تصویر محصول'" 
                     class="w-full h-full object-cover cursor-pointer hover:opacity-80 transition-opacity duration-300" 
                     @click.prevent="previewImage=img;previewShow=true"
                     @error="onImageError($event)"
                   />
                   
                   <!-- placeholder -->
                   <div v-else class="w-full h-full bg-gray-100 flex items-center justify-center">
                     <svg class="w-8 h-8 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                       <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                     </svg>
                   </div>
                 </div>
               </div>
               
               <!-- ترتیب -->
               <div class="col-span-1 flex justify-center">
                 <input type="number" :value="idx+1" min="1" :max="store.images.length" class="w-16 h-10 bg-white border border-emerald-300 text-gray-900 text-sm rounded-lg focus:ring-emerald-500 focus:border-emerald-500 p-2 text-center font-semibold" @change="onOrderChange(idx, $event)" />
               </div>
               
               <!-- عنوان -->
               <div class="col-span-4">
                 <input v-model="img.title" type="text" placeholder="عنوان تصویر..." class="w-full bg-white border border-emerald-300 text-gray-900 text-sm rounded-lg focus:ring-emerald-500 focus:border-emerald-500 p-2.5" />
               </div>
               
               <!-- متن جایگزین -->
               <div class="col-span-3">
                 <input v-model="img.alt" type="text" placeholder="توضیح تصویر برای SEO..." class="w-full bg-white border border-emerald-300 text-gray-900 text-sm rounded-lg focus:ring-emerald-500 focus:border-emerald-500 p-2.5" />
               </div>
               
               <!-- عملیات -->
               <div class="col-span-1 flex justify-center">
                 <button class="w-10 h-10 bg-red-100 hover:bg-red-200 text-red-600 rounded-lg transition-colors duration-300 flex items-center justify-center" @click.prevent="store.removeImage(img.id)">
                   <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                     <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                   </svg>
                 </button>
               </div>
             </div>
           </div>
         </div>
       </div>
       
       <!-- پیام خالی بودن -->
       <div v-else class="text-center py-12">
         <div class="w-24 h-24 mx-auto mb-4 bg-emerald-100 rounded-full flex items-center justify-center">
           <svg class="w-12 h-12 text-emerald-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
             <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
           </svg>
         </div>
         <h3 class="text-lg font-semibold text-emerald-800 mb-2">هنوز تصویری آپلود نشده</h3>
         <p class="text-emerald-600 mb-4">برای شروع، تصاویر محصول خود را از کتابخانه انتخاب کنید</p>
       </div>
     </div>
   </div>

   <!-- تنظیمات گالری و نمایش -->
   <div class="bg-gradient-to-r from-purple-50 to-pink-50 border border-purple-200 rounded-xl shadow-lg px-4 py-4 text-right transition-all duration-300 hover:shadow-xl">
     <div class="flex items-center justify-between cursor-pointer" @click="toggleSection('gallerySettings')">
       <div class="flex items-center gap-3">
         <div class="w-10 h-10 bg-gradient-to-r from-purple-600 to-pink-600 rounded-lg flex items-center justify-center shadow-md">
           <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
             <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"></path>
             <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
           </svg>
         </div>
         <h3 class="text-lg font-bold text-purple-800">تنظیمات نمایش و گالری</h3>
       </div>
       <span class="text-purple-600 text-xl font-bold bg-white rounded-full w-8 h-8 flex items-center justify-center shadow-md">{{ sections.gallerySettings ? '−' : '+' }}</span>
     </div>

     <div v-show="sections.gallerySettings" class="mt-6">
       <div class="grid grid-cols-1 gapx-4 py-4">
         <!-- تنظیمات ابعاد تصاویر -->
         <div class="bg-white rounded-xl px-4 py-4 shadow-md border border-purple-200">
           <h4 class="text-base font-bold text-purple-800 mb-4 flex items-center gap-2">
             <svg class="w-5 h-5 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
               <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 8V4m0 0h4M4 4l5 5m11-1V4m0 0h-4m4 0l-5 5M4 16v4m0 0h4m-4 0l5-5m11 5v-4m0 4h-4m4 0l-5-5"></path>
             </svg>
             تنظیمات ابعاد تصاویر
           </h4>
           <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
             <div class="bg-purple-50 rounded-lg px-4 py-4">
               <label class="text-sm text-purple-700 font-semibold mb-2 block">ابعاد تصویر اصلی</label>
               <select class="w-full bg-white border border-purple-300 text-gray-900 text-sm rounded-lg focus:ring-purple-500 focus:border-purple-500 p-2.5">
                 <option value="800x800">800×800 پیکسل (استاندارد)</option>
                 <option value="1000x1000">1000×1000 پیکسل (کیفیت بالا)</option>
                 <option value="1200x1200">1200×1200 پیکسل (بالا)</option>
                 <option value="original">اندازه اصلی فایل</option>
               </select>
             </div>
             <div class="bg-pink-50 rounded-lg px-4 py-4">
               <label class="text-sm text-pink-700 font-semibold mb-2 block">ابعاد تصاویر کوچک</label>
               <select class="w-full bg-white border border-pink-300 text-gray-900 text-sm rounded-lg focus:ring-pink-500 focus:border-pink-500 p-2.5">
                 <option value="150x150">150×150 پیکسل</option>
                 <option value="200x200">200×200 پیکسل</option>
                 <option value="300x300">300×300 پیکسل</option>
               </select>
             </div>
           </div>
         </div>

         <!-- تنظیمات عملکرد گالری -->
         <div class="bg-white rounded-xl px-4 py-4 shadow-md border border-purple-200">
           <h4 class="text-base font-bold text-purple-800 mb-4 flex items-center gap-2">
             <svg class="w-5 h-5 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
               <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
             </svg>
             قابلیت‌های نمایش گالری
           </h4>
           <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
             <div class="space-y-3">
               <label class="flex items-center gap-3 p-3 bg-purple-50 rounded-lg cursor-pointer hover:bg-purple-100 transition-colors duration-300">
                 <input type="checkbox" class="w-5 h-5 text-purple-600 bg-white border-purple-300 rounded focus:ring-purple-500" checked />
                 <div class="flex items-center gap-2">
                   <svg class="w-4 h-4 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                     <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0zM10 7v3m0 0v3m0-3h3m-3 0H7"></path>
                   </svg>
                   <span class="text-sm text-purple-700 font-medium">نمایش بزرگ‌نمایی</span>
                 </div>
               </label>
               <label class="flex items-center gap-3 p-3 bg-purple-50 rounded-lg cursor-pointer hover:bg-purple-100 transition-colors duration-300">
                 <input type="checkbox" class="w-5 h-5 text-purple-600 bg-white border-purple-300 rounded focus:ring-purple-500" checked />
                 <div class="flex items-center gap-2">
                   <svg class="w-4 h-4 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                     <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.828 14.828a4 4 0 01-5.656 0M9 10h1m4 0h1m-6 4h8m-4-8v1m0 4v1"></path>
                   </svg>
                   <span class="text-sm text-purple-700 font-medium">امکان اسلایدشو</span>
                 </div>
               </label>
             </div>
             <div class="space-y-3">
               <label class="flex items-center gap-3 p-3 bg-pink-50 rounded-lg cursor-pointer hover:bg-pink-100 transition-colors duration-300">
                 <input type="checkbox" class="w-5 h-5 text-pink-600 bg-white border-pink-300 rounded focus:ring-pink-500" />
                 <div class="flex items-center gap-2">
                   <svg class="w-4 h-4 text-pink-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                     <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 8V4m0 0h4M4 4l5 5m11-1V4m0 0h-4m4 0l-5 5M4 16v4m0 0h4m-4 0l5-5m11 5v-4m0 4h-4m4 0l-5-5"></path>
                   </svg>
                   <span class="text-sm text-pink-700 font-medium">نمایش تمام‌صفحه</span>
                 </div>
               </label>
               <label class="flex items-center gap-3 p-3 bg-pink-50 rounded-lg cursor-pointer hover:bg-pink-100 transition-colors duration-300">
                 <input type="checkbox" class="w-5 h-5 text-pink-600 bg-white border-pink-300 rounded focus:ring-pink-500" />
                 <div class="flex items-center gap-2">
                   <svg class="w-4 h-4 text-pink-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                     <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                   </svg>
                   <span class="text-sm text-pink-700 font-medium">دانلود تصاویر</span>
                 </div>
               </label>
             </div>
           </div>
         </div>

         <!-- آمار کامل تصاویر -->
         <div class="bg-gradient-to-r from-purple-100 to-pink-100 rounded-xl px-4 py-4 shadow-md border border-purple-200">
           <h4 class="text-base font-bold text-purple-800 mb-4 flex items-center gap-2">
             <svg class="w-5 h-5 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
               <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v4"></path>
             </svg>
             آمار و اطلاعات تصاویر
           </h4>
           <div class="grid grid-cols-2 md:grid-cols-4 gapx-4 py-4">
             <div class="bg-white rounded-lg px-4 py-4 text-center shadow-sm">
               <div class="text-2xl font-bold text-purple-600">{{ store.images.length }}</div>
               <div class="text-xs text-purple-700 mt-1">تعداد تصاویر</div>
             </div>
             <div class="bg-white rounded-lg px-4 py-4 text-center shadow-sm">
               <div class="text-2xl font-bold text-pink-600">{{ store.images.filter(img => img.title).length }}</div>
               <div class="text-xs text-pink-700 mt-1">دارای عنوان</div>
             </div>
             <div class="bg-white rounded-lg px-4 py-4 text-center shadow-sm">
               <div class="text-2xl font-bold text-indigo-600">{{ store.images.filter(img => img.alt).length }}</div>
               <div class="text-xs text-indigo-700 mt-1">دارای Alt Text</div>
             </div>
             <div class="bg-white rounded-lg px-4 py-4 text-center shadow-sm">
               <div class="text-2xl font-bold text-teal-600">100%</div>
               <div class="text-xs text-teal-700 mt-1">آماده نمایش</div>
             </div>
           </div>
         </div>
       </div>
     </div>
   </div>

   <!-- Media Library Modal -->
    <MediaLibraryModal
     v-model="showMediaModal"
     :model-selected="store.images.map(i=>i.id)"
      :default-category="defaultMediaCategory"
     @confirm="onImagesSelected"
   />
   <MediaPreviewModal
     v-model="previewShow"
     :file="previewImage"
     @delete="onDeleteImage"
     @save="onSaveImageMeta"
   />
 </template>

  <script setup lang="ts">
  import { ref } from 'vue'
import MediaLibraryModal from '~/components/media/MediaLibraryModal.vue'
import MediaPreviewModal from '~/components/media/MediaPreviewModal.vue'
import { useProductCreateStore } from '~/stores/productCreate'

 // Props for edit mode
 const props = defineProps({
   isEditMode: {
     type: Boolean,
     default: false
   },
   productId: {
     type: [String, Number],
     default: null
    },
    // پیش‌فرض دسته‌بندی کتابخانه رسانه فقط در صفحات محصول (ایجاد/ویرایش)
    defaultMediaCategory: {
      type: String,
      default: undefined
   }
 })

 const store = useProductCreateStore()

 // Use store sections
 const sections = store.sections
 const toggleSection = store.toggleSection

 // Media modal state
 const showMediaModal = ref(false)
 const previewImage = ref<any|null>(null)
 const previewShow = ref(false)

 // drag and drop ordering
 const dragIdx = ref<number|null>(null)
 function onDragStart(i:number){ dragIdx.value = i }
 function onDrop(i:number){
   if(dragIdx.value===null) return
   const arr = [...store.images]
   const item = arr.splice(dragIdx.value,1)[0]
   arr.splice(i,0,item)
   store.images = arr
   dragIdx.value=null
 }

 function onOrderChange(oldIdx:number, e:Event){
   const input = e.target as HTMLInputElement
   let newPos = parseInt(input.value,10)
   if(isNaN(newPos)) return
   const len = store.images.length
   if(newPos < 1) newPos = 1
   if(newPos > len) newPos = len
   newPos = newPos - 1 // zero-based
   if(newPos === oldIdx) return
   const arr = [...store.images]
   const item = arr.splice(oldIdx,1)[0]
   arr.splice(newPos,0,item)
   store.images = arr
 }

 function openMedia(){
   showMediaModal.value = true
 }

 function onImagesSelected(files:any[]){
   if(files && files.length){
     store.addImages(files)
   }
   showMediaModal.value = false
 }

 // Handle delete emitted from preview modal
 async function onDeleteImage(id:number){
   console.log('delete event received', id)
   store.removeImage(id)
   previewShow.value = false
   try{
     await $fetch(`/api/media/${id}`,{
       method:'DELETE'
     })
   }catch(err){ console.error('delete media error', err) }
 }

 async function onSaveImageMeta(payload:any){
   // optimistic update
   const idx = store.images.findIndex(i=>i.id===payload.id)
   if(idx>-1){
     store.images[idx] = { ...store.images[idx], ...payload }
   }
   console.log('save event received', payload)
   try{
     await $fetch(`/api/media/${payload.id}`,{
       method:'PUT',
       body:payload
     })
   }catch(err){ console.error('save media error', err) }
 }

 // --------------------------------------------------
 // File upload logic for drag-and-drop & file picker
 // --------------------------------------------------

// منطق آپلود غیرفعال شد؛ فقط انتخاب از کتابخانه استفاده می‌شود
// این توابع برای سازگاری با template نگه داشته شده‌اند
function onDropFiles(event: DragEvent) {
  // منطق آپلود غیرفعال است - از کتابخانه رسانه استفاده کنید
  event.preventDefault()
  console.log('Drag and drop disabled - please use media library')
}

function onFileInput(event: Event) {
  // منطق آپلود غیرفعال است - از کتابخانه رسانه استفاده کنید
  const target = event.target as HTMLInputElement
  if (target.files && target.files.length > 0) {
    console.log('File input disabled - please use media library')
    // Reset input
    target.value = ''
  }
}

 // Simple error handling for images
 function onImageError(event: Event) {
   const target = event.target as HTMLImageElement
   console.log('Image failed to load:', target.src)
   
   // Hide broken image and show placeholder icon instead
   target.style.display = 'none'
   
   const container = target.parentElement
   if (container && !container.querySelector('.error-placeholder')) {
     const placeholder = document.createElement('div')
     placeholder.className = 'error-placeholder w-full h-full bg-gray-100 flex items-center justify-center'
     placeholder.innerHTML = '<svg class="w-8 h-8 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.864-.833-2.634 0L4.18 16.5c-.77.833.192 2.5 1.732 2.5z"></path></svg>'
     container.appendChild(placeholder)
   }
 }

 // expose to template
 // (nothing to return in <script setup>)
 </script>
