<template>
  <div class="redirects-management">
    <ToastContainer />
    <!-- Header -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-6 flex flex-col md:flex-row md:items-center md:justify-between gap-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-900 mb-2">مدیریت تغییر مسیر</h1>
        <p class="text-gray-600">مدیریت و تنظیم تغییر مسیرهای 301، 302 و سایر انواع ریدایرکت برای بهبود سئو</p>
      </div>
      <div class="flex gap-2 flex-wrap">
        <button
          class="bg-green-600 text-white px-4 py-2 rounded-lg hover:bg-green-700 transition-colors flex items-center gap-2"
          @click="importRedirects"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M9 19l3 3m0 0l3-3m-3 3V10"></path>
          </svg>
          وارد کردن از فایل
        </button>
        <button
          class="bg-purple-600 text-white px-4 py-2 rounded-lg hover:bg-purple-700 transition-colors flex items-center gap-2"
          @click="exportRedirects"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
          </svg>
          صادر کردن
        </button>
      </div>
    </div>

    <!-- Stats Cards -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-6">
      <TemplateCard title="کل تغییر مسیرها" :value="stats.total" variant="blue">
        <template #icon>
          <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7l5 5m0 0l-5 5m5-5H6"></path>
          </svg>
        </template>
      </TemplateCard>

      <TemplateCard title="فعال" :value="stats.active" variant="green">
        <template #icon>
          <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
          </svg>
        </template>
      </TemplateCard>

      <TemplateCard title="نیازمند بررسی" :value="stats.needsReview" variant="yellow">
        <template #icon>
          <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
          </svg>
        </template>
      </TemplateCard>

      <TemplateCard title="خطا" :value="stats.errors" variant="red">
        <template #icon>
          <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
        </template>
      </TemplateCard>
    </div>

    <!-- Redirect Types Info -->
    <div class="bg-green-50 border border-green-200 rounded-lg p-6 mb-6">
      <button class="flex items-center gap-2 text-green-900 font-semibold focus:outline-none" @click="showTypesInfo = !showTypesInfo">
        <svg :class="{'rotate-90': showTypesInfo, 'rotate-0': !showTypesInfo}" class="w-4 h-4 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
        </svg>
        انواع تغییر مسیر و خطاها
        <span class="text-xs text-green-600">({{ showTypesInfo ? 'بستن' : 'نمایش' }})</span>
      </button>
      <div v-show="showTypesInfo" class="mt-4">
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 text-sm">
          <div class="bg-white p-3 rounded border">
            <div class="font-medium text-blue-800 mb-1">301 - انتقال دائمی</div>
            <p class="text-gray-600">برای تغییر مسیرهای دائمی، بهترین گزینه برای سئو</p>
          </div>
          <div class="bg-white p-3 rounded border">
            <div class="font-medium text-yellow-800 mb-1">302 - انتقال موقت</div>
            <p class="text-gray-600">برای تغییر مسیرهای موقت، قدرت سئو منتقل نمی‌شود</p>
          </div>
          <div class="bg-white p-3 rounded border">
            <div class="font-medium text-orange-800 mb-1">303 - See Other</div>
            <p class="text-gray-600">برای تغییر مسیر به روش GET، معمولاً برای فرم‌ها</p>
          </div>
          <div class="bg-white p-3 rounded border">
            <div class="font-medium text-purple-800 mb-1">307 - انتقال موقت (حفظ متد)</div>
            <p class="text-gray-600">حفظ متد HTTP در تغییر مسیر موقت</p>
          </div>
          <div class="bg-white p-3 rounded border">
            <div class="font-medium text-indigo-800 mb-1">308 - انتقال دائمی (حفظ متد)</div>
            <p class="text-gray-600">حفظ متد HTTP در تغییر مسیر دائمی</p>
          </div>
          <div class="bg-white p-3 rounded border">
            <div class="font-medium text-red-800 mb-1">404 - Not Found</div>
            <p class="text-gray-600">صفحه یافت نشد، رایج‌ترین خطای HTTP</p>
          </div>
          <div class="bg-white p-3 rounded border">
            <div class="font-medium text-red-800 mb-1">403 - Forbidden</div>
            <p class="text-gray-600">دسترسی ممنوع، کاربر مجاز به دسترسی نیست</p>
          </div>
          <div class="bg-white p-3 rounded border">
            <div class="font-medium text-red-800 mb-1">405 - Method Not Allowed</div>
            <p class="text-gray-600">متد HTTP غیرمجاز (مثل POST روی GET-only endpoint)</p>
          </div>
          <div class="bg-white p-3 rounded border">
            <div class="font-medium text-red-800 mb-1">500 - Internal Server Error</div>
            <p class="text-gray-600">خطای داخلی سرور، مشکل در کد یا دیتابیس</p>
          </div>
          <div class="bg-white p-3 rounded border">
            <div class="font-medium text-orange-800 mb-1">503 - Service Unavailable</div>
            <p class="text-gray-600">سرویس در دسترس نیست، معمولاً برای نگهداری</p>
          </div>
          <div class="bg-white p-3 rounded border">
            <div class="font-medium text-orange-800 mb-1">429 - Too Many Requests</div>
            <p class="text-gray-600">تعداد درخواست‌ها بیش از حد، Rate Limiting</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Actions Bar -->
    <div class="bg-white rounded-lg shadow border border-gray-200 p-6 mb-6">
      <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-6">
        <div class="flex flex-col sm:flex-row gap-3">
          <button 
            v-if="hasPermission('seo.create')"
            class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors flex items-center gap-2"
            @click="showAddModal = true"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
            </svg>
            افزودن تغییر مسیر جدید
          </button>
        </div>

        <div class="flex gap-2">
          <!-- فیلتر گروه -->
          <select v-model="filterGroup" class="border border-gray-300 rounded-lg px-3 py-2 text-sm">
            <option value="">همه گروه‌ها</option>
            <option v-for="group in groups" :key="group.groupName" :value="group.groupName">
              {{ group.groupName }} ({{ group.cnt }})
            </option>
          </select>
          
          <!-- فیلتر نوع -->
          <select v-model="filterType" class="border border-gray-300 rounded-lg px-3 py-2 text-sm">
             <option value="">همه انواع</option>
             <optgroup label="تغییر مسیرها">
               <option value="301">301 - انتقال دائمی</option>
               <option value="302">302 - انتقال موقت</option>
               <option value="303">303 - See Other</option>
               <option value="307">307 - انتقال موقت (حفظ متد)</option>
               <option value="308">308 - انتقال دائمی (حفظ متد)</option>
             </optgroup>
             <optgroup label="خطاهای رایج">
               <option value="404">404 - Not Found</option>
               <option value="403">403 - Forbidden</option>
               <option value="500">500 - Internal Server Error</option>
               <option value="503">503 - Service Unavailable</option>
             </optgroup>
             <optgroup label="سایر">
               <option value="meta">Meta Refresh</option>
               <option value="js">JavaScript Redirect</option>
             </optgroup>
           </select>
        </div>
      </div>
    </div>

    <!-- Redirects Table -->
    <div class="bg-white rounded-lg shadow border border-gray-200 overflow-x-auto w-full pb-24 mb-10" style="min-height: 120px; margin-bottom: 80px;">
      <!-- دکمه تغییر گروه -->
      <div v-if="selectedRedirects.length > 0" class="p-3 bg-blue-50 flex items-center gap-3 border-b border-blue-200">
        <span class="text-blue-800 text-sm">{{ selectedRedirects.length }} ریدایرکت انتخاب شده</span>
        <button class="bg-blue-600 text-white px-3 py-1 rounded hover:bg-blue-700 text-sm" @click="openGroupModal">تغییر گروه</button>
        <button v-if="canDeleteRedirect" class="bg-red-600 text-white px-3 py-1 rounded hover:bg-red-700 text-sm" @click="deleteSelectedRedirects">حذف</button>
        <button class="text-gray-500 hover:text-red-600 text-xs" @click="selectedRedirects = []; selectAll = false">لغو انتخاب</button>
      </div>

      <!-- Empty State -->
      <div v-if="redirects.length === 0" class="text-center py-12">
        <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7l5 5m0 0l-5 5m5-5H6"></path>
        </svg>
        <h3 class="mt-2 text-sm font-medium text-gray-900">هیچ ریدایرکتی یافت نشد</h3>
        <p class="mt-1 text-sm text-gray-500">برای شروع، یک ریدایرکت جدید ایجاد کنید.</p>
        <div class="mt-6">
          <button 
            v-if="hasPermission('seo.create')"
            class="inline-flex items-center px-4 py-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
            @click="showAddModal = true"
          >
            <svg class="-ml-1 mr-2 h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
            </svg>
            افزودن ریدایرکت جدید
          </button>
        </div>
      </div>

      <!-- Table - فقط وقتی داده وجود دارد -->
      <div v-if="redirects.length > 0" class="overflow-x-auto">
        <table class="w-full min-w-[900px] divide-y divide-gray-200 text-xs">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-1 py-2 text-center w-6">
                <input v-model="selectAll" type="checkbox" :checked="selectAll" @change="toggleSelectAll" />
              </th>
              <th class="px-1 py-2 text-center w-10">ردیف</th>
              <th class="px-2 py-2 text-right font-medium text-gray-500 uppercase tracking-wider max-w-[120px] truncate">مسیر مبدا</th>
              <th class="px-2 py-2 text-right font-medium text-gray-500 uppercase tracking-wider max-w-[120px] truncate">مسیر مقصد</th>
              <th class="px-2 py-2 text-right font-medium text-gray-500 uppercase tracking-wider w-14">نوع</th>
              <th class="px-2 py-2 text-right font-medium text-gray-500 uppercase tracking-wider w-14">وضعیت</th>
              <th class="px-2 py-2 text-right font-medium text-gray-500 uppercase tracking-wider w-16">بازدید</th>
              <th class="px-2 py-2 text-right font-medium text-gray-500 uppercase tracking-wider w-20">آخرین بازدید</th>
              <th class="px-1 py-2 text-right font-medium text-gray-500 uppercase tracking-wider w-20">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="(redirect, idx) in paginatedRedirects" :key="redirect.id || idx" class="hover:bg-gray-50">
              <td class="px-1 py-2 text-center w-6">
                <input type="checkbox" :value="redirect.id" :checked="selectedRedirects.includes(redirect.id)" @change="toggleSelectOne(redirect.id)" />
              </td>
              <td class="px-1 py-2 text-center w-10">{{ (currentPage - 1) * itemsPerPage + idx + 1 }}</td>
              <td class="px-2 py-2 max-w-[120px] truncate" :title="redirect.source_path">{{ redirect.source_path && redirect.source_path !== '' ? redirect.source_path : '---' }}</td>
              <td class="px-2 py-2 max-w-[120px] truncate" :title="redirect.target_path">{{ redirect.target_path && redirect.target_path !== '' ? redirect.target_path : '---' }}</td>
              <td class="px-2 py-2 w-14">
                <span :class="getTypeClass(redirect.code)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ redirect.code || 'نامشخص' }}
                </span>
              </td>
              <td class="px-2 py-2 w-14">
                <span :class="getStatusClass(redirect.redirect_type)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ getStatusText(redirect.redirect_type) }}
                </span>
              </td>
              <td class="px-2 py-2 w-16 text-center">{{ formatNumber(redirect.visit_count) }}</td>
              <td class="px-2 py-2 w-20 text-center">{{ formatDate(redirect.last_visited_at) }}</td>
              <td class="px-1 py-2 w-20 text-center">
                <div class="flex gap-1 justify-center">
                  <button v-if="hasPermission('seo.update')" class="text-blue-600 hover:text-blue-900" @click="editRedirect(redirect)">ویرایش</button>
                  <button v-if="hasPermission('seo.update')" :class="redirect.status === 'active' ? 'text-red-600 hover:text-red-900' : 'text-green-600 hover:text-green-900'" @click="toggleRedirect(redirect)">
                    {{ redirect.status === 'active' ? 'غیرفعال' : 'فعال' }}
                  </button>
                  <button v-if="canDeleteRedirect" class="text-red-600 hover:text-red-900" @click="deleteRedirect(redirect)">حذف</button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Pagination -->
      <template v-if="totalPages > 1 && redirects.length > 0">
        <div class="pt-6 pb-8 flex justify-center">
          <Pagination
            :current-page="currentPage"
            :total-pages="totalPages"
            :total="totalItems"
            :items-per-page="itemsPerPage"
            @page-changed="handlePageChange"
            @items-per-page-changed="handleItemsPerPageChange"
          />
        </div>
      </template>
    </div>

    <!-- Add/Edit Modal -->
    <div v-if="showAddModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg shadow-xl p-6 w-full max-w-2xl mx-4">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-medium text-gray-900">
            {{ editingRedirect ? 'ویرایش تغییر مسیر' : 'افزودن تغییر مسیر جدید' }}
          </h3>
          <button class="text-gray-400 hover:text-gray-600" @click="closeModal">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>

        <form class="space-y-4" @submit.prevent="saveRedirect">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">مسیر مبدا *</label>
              <input 
                v-model="redirectForm.sourcePath"
                type="text"
                class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="/old-page"
                required
              />
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">مسیر مقصد *</label>
              <input 
                v-model="redirectForm.targetPath"
                type="text"
                class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="/new-page"
                required
              />
            </div>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">گروه</label>
              <select 
                v-model="redirectForm.groupName"
                class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="">بدون گروه</option>
                <option value="انتقال محصولات">انتقال محصولات</option>
                <option value="تغییرات URL">تغییرات URL</option>
                <option value="حذف صفحات">حذف صفحات</option>
                <option value="SEO Optimization">SEO Optimization</option>
                <option value="Mobile Redirects">Mobile Redirects</option>
                <option value="Landing Pages">Landing Pages</option>
                <option value="سایر">سایر</option>
              </select>
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">نوع تغییر مسیر *</label>
                             <select 
                 v-model="redirectForm.type"
                 class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                 required
               >
                 <optgroup label="تغییر مسیرها">
                   <option value="301">301 - انتقال دائمی</option>
                   <option value="302">302 - انتقال موقت</option>
                   <option value="303">303 - See Other</option>
                   <option value="304">304 - Not Modified</option>
                   <option value="305">305 - Use Proxy</option>
                   <option value="307">307 - انتقال موقت (حفظ متد)</option>
                   <option value="308">308 - انتقال دائمی (حفظ متد)</option>
                   <option value="310">310 - Too Many Redirects</option>
                 </optgroup>
                 <optgroup label="خطاهای HTTP">
                   <option value="400">400 - Bad Request</option>
                   <option value="401">401 - Unauthorized</option>
                   <option value="403">403 - Forbidden</option>
                   <option value="404">404 - Not Found</option>
                   <option value="405">405 - Method Not Allowed</option>
                   <option value="406">406 - Not Acceptable</option>
                   <option value="408">408 - Request Timeout</option>
                   <option value="409">409 - Conflict</option>
                   <option value="410">410 - Gone</option>
                   <option value="411">411 - Length Required</option>
                   <option value="412">412 - Precondition Failed</option>
                   <option value="413">413 - Payload Too Large</option>
                   <option value="414">414 - URI Too Long</option>
                   <option value="415">415 - Unsupported Media Type</option>
                   <option value="416">416 - Range Not Satisfiable</option>
                   <option value="417">417 - Expectation Failed</option>
                   <option value="418">418 - I'm a teapot</option>
                   <option value="421">421 - Misdirected Request</option>
                   <option value="422">422 - Unprocessable Entity</option>
                   <option value="423">423 - Locked</option>
                   <option value="424">424 - Failed Dependency</option>
                   <option value="425">425 - Too Early</option>
                   <option value="426">426 - Upgrade Required</option>
                   <option value="428">428 - Precondition Required</option>
                   <option value="429">429 - Too Many Requests</option>
                   <option value="431">431 - Request Header Fields Too Large</option>
                   <option value="451">451 - Unavailable For Legal Reasons</option>
                   <option value="500">500 - Internal Server Error</option>
                   <option value="501">501 - Not Implemented</option>
                   <option value="502">502 - Bad Gateway</option>
                   <option value="503">503 - Service Unavailable</option>
                   <option value="504">504 - Gateway Timeout</option>
                   <option value="505">505 - HTTP Version Not Supported</option>
                   <option value="506">506 - Variant Also Negotiates</option>
                   <option value="507">507 - Insufficient Storage</option>
                   <option value="508">508 - Loop Detected</option>
                   <option value="510">510 - Not Extended</option>
                   <option value="511">511 - Network Authentication Required</option>
                 </optgroup>
                 <optgroup label="سایر روش‌ها">
                   <option value="meta">Meta Refresh</option>
                   <option value="js">JavaScript Redirect</option>
                 </optgroup>
               </select>
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">وضعیت</label>
              <select 
                v-model="redirectForm.status"
                class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="active">فعال</option>
                <option value="inactive">غیرفعال</option>
              </select>
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">توضیحات</label>
            <textarea 
              v-model="redirectForm.description"
              rows="3"
              class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="توض��حات اختیاری دربار�� این تغییر مسیر..."
            ></textarea>
          </div>

          <div class="flex justify-end gap-3 pt-4">
            <button 
              type="button"
              class="px-4 py-2 border border-gray-300 rounded-lg text-gray-700 hover:bg-gray-50"
              @click="closeModal"
            >
              انصراف
            </button>
            <button 
              type="submit"
              class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700"
            >
              {{ editingRedirect ? 'ویرایش' : 'افزودن' }}
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- Group Change Modal -->
    <div v-if="showGroupModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg shadow-xl p-6 w-full max-w-md mx-4">
        <div class="mb-4">
          <h3 class="text-lg font-medium text-gray-900">تغییر گروه ریدایرکت‌ها</h3>
        </div>

        <div class="mb-4">
          <label class="block text-sm font-medium text-gray-700 mb-1">نام گروه جدید *</label>
          <input
            v-model="groupModalName"
            type="text"
            class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            placeholder="گروه جدید"
            required
          />
        </div>

        <div class="flex justify-end gap-3">
          <button
            type="button"
            class="px-4 py-2 border border-gray-300 rounded-lg text-gray-700 hover:bg-gray-50"
            @click="closeGroupModal"
          >
            انصراف
          </button>
          <button
            class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700"
            @click="changeGroupForSelected"
          >
            تغییر گروه
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
declare const navigateTo: (to: string) => Promise<void>
interface User {
  id?: number | string
  name?: string
  [key: string]: unknown
}
declare const useAuth: () => { user: User | null; hasPermission: (perm: string) => boolean }
</script>

<script setup lang="ts">
import Pagination from '@/components/common/Pagination.vue'
import TemplateCard from '@/components/common/TemplateCard.vue'
import ToastContainer from '@/components/common/ToastContainer.vue'
import { useToast } from '@/composables/useToast'
import { computed, onMounted, reactive, ref, watch } from 'vue'

const { hasPermission } = useAuth()
const { showSuccess, showError } = useToast()

const canDeleteRedirect = computed(() => hasPermission('seo.delete'))

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// State
const showAddModal = ref(false)
const showGroupModal = ref(false)
// eslint-disable-next-line @typescript-eslint/no-explicit-any
const editingRedirect = ref<any>(null)
const filterType = ref('')
// const filterStatus = ref('')
const filterGroup = ref('')
const selectedRedirects = ref<number[]>([])
const selectAll = ref(false)
const groupModalName = ref('')
const showTypesInfo = ref(false)

// Form data
const redirectForm = reactive({
  sourcePath: '',
  targetPath: '',
  groupName: '',
  type: '301',
  status: 'active',
  description: ''
})

// Stats
const stats = reactive({
  total: 0,
  active: 0,
  needsReview: 0,
  errors: 0
})

// Empty data array - will be populated from API
const redirects = ref([])
const groups = ref([])

// Pagination state
const currentPage = ref(1)
const itemsPerPage = ref(10)

// محاسبه داده‌های فیلترشده
const filteredRedirects = computed(() => {
  let filtered = redirects.value

  if (filterGroup.value) {
    filtered = filtered.filter(r => r.group_name === filterGroup.value)
  }

  if (filterType.value) {
    filtered = filtered.filter(r => r.code?.toString() === filterType.value)
  }

  return filtered
})

const totalItems = computed(() => filteredRedirects.value.length)
const totalPages = computed(() => Math.ceil(totalItems.value / itemsPerPage.value))

const paginatedRedirects = computed(() => {
  const filtered = filteredRedirects.value
  
  const total = filtered.length;
  if (total === 0) return [];
  
  const start = (currentPage.value - 1) * itemsPerPage.value;
  // اگر داده‌ای وجود ندارد یا slice خارج از بازه است، آرایه خالی برگردان
  if (start >= total) return [];
  return filtered.slice(start, start + itemsPerPage.value).filter(r => r && r.id);
})

watch(totalPages, (newTotalPages) => {
  if (currentPage.value > newTotalPages) {
    currentPage.value = 1;
  }
})


// توابع کمکی با فیلد snake_case
const getTypeClass = (code: number | string) => {
  // بررسی null/undefined
  if (!code) {
    return 'bg-gray-100 text-gray-800'
  }
  
  const type = code.toString()

  // Redirects
  if (['301', '302', '303', '307', '308'].includes(type)) {
    switch (type) {
      case '301': return 'bg-blue-100 text-blue-800'
      case '302': return 'bg-yellow-100 text-yellow-800'
      case '303': return 'bg-orange-100 text-orange-800'
      case '307': return 'bg-purple-100 text-purple-800'
      case '308': return 'bg-indigo-100 text-indigo-800'
    }
  }
  
  // Client Errors (4xx)
  if (type.startsWith('4')) {
    switch (type) {
      case '400': return 'bg-red-100 text-red-800'
      case '401': return 'bg-red-100 text-red-800'
      case '403': return 'bg-red-100 text-red-800'
      case '404': return 'bg-red-100 text-red-800'
      case '405': return 'bg-red-100 text-red-800'
      case '429': return 'bg-orange-100 text-orange-800'
      case '451': return 'bg-red-100 text-red-800'
      default: return 'bg-red-100 text-red-800'
    }
  }
  
  // Server Errors (5xx)
  if (type.startsWith('5')) {
    switch (type) {
      case '500': return 'bg-red-100 text-red-800'
      case '502': return 'bg-red-100 text-red-800'
      case '503': return 'bg-orange-100 text-orange-800'
      case '504': return 'bg-orange-100 text-orange-800'
      default: return 'bg-red-100 text-red-800'
    }
  }
  
  // Other types
  switch (type) {
    case '310': return 'bg-red-100 text-red-800'
    case 'meta': return 'bg-green-100 text-green-800'
    case 'js': return 'bg-pink-100 text-pink-800'
    default: return 'bg-gray-100 text-gray-800'
  }
}

const getStatusClass = (redirectType: string) => {
  if (!redirectType || typeof redirectType !== 'string') {
    return 'bg-gray-100 text-gray-800'
  }
  
  switch (redirectType) {
    case 'permanent':
      return 'bg-green-100 text-green-800'
    case 'temporary':
      return 'bg-yellow-100 text-yellow-800'
    case 'error':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const getStatusText = (redirectType: string) => {
  if (!redirectType || typeof redirectType !== 'string') {
    return 'نامشخص'
  }
  
  switch (redirectType) {
    case 'permanent':
      return 'دائمی'
    case 'temporary':
      return 'موقت'
    case 'error':
      return 'خطا'
    default:
      return 'نامشخص'
  }
}

const formatDate = (date: string | null) => {
  if (!date) return 'هیچ‌وقت'
  return new Date(date).toLocaleDateString('fa-IR')
}

const formatNumber = (num: number | undefined | null) => {
  return (num || 0).toLocaleString()
}

const updateStats = () => {
  stats.total = redirects.value.length
  stats.active = redirects.value.filter(r => r.redirect_type === 'permanent').length
  stats.needsReview = redirects.value.filter(r => (r.visit_count || 0) === 0 && r.redirect_type === 'permanent').length
  stats.errors = redirects.value.filter(r => r.redirect_type === 'error').length
}

const closeModal = () => {
  showAddModal.value = false
  editingRedirect.value = null
  resetForm()
}

const closeGroupModal = () => {
  showGroupModal.value = false
  groupModalName.value = ''
}

const resetForm = () => {
  redirectForm.sourcePath = ''
  redirectForm.targetPath = ''
  redirectForm.groupName = ''
  redirectForm.type = '301'
  redirectForm.status = 'active'
  redirectForm.description = ''
}

interface Redirect {
  id?: number | string
  from?: string
  to?: string
  [key: string]: unknown
}
const editRedirect = (redirect: Redirect) => {
  // Navigate to group redirects page
  navigateTo(`/admin/seo/redirects/group/${encodeURIComponent(redirect.group_name || 'بدون گروه')}`)
}

const saveRedirect = async () => {
  try {
    // اعتبارسنجی فرم
    if (!redirectForm.sourcePath.trim()) {
      showError('مسیر مبدا الزامی است')
      return
    }
    if (!redirectForm.targetPath.trim()) {
      showError('مسیر مقصد الزامی است')
      return
    }

    if (editingRedirect.value) {
      showError('قابلیت ویرایش ریدایرکت در حال توسعه است')
      return
    } else {
      // Add new redirect
      // تبدیل URL کامل به path
      const sourcePath = redirectForm.sourcePath.trim().startsWith('http') 
        ? new URL(redirectForm.sourcePath.trim()).pathname 
        : redirectForm.sourcePath.trim()
      
      const targetPath = redirectForm.targetPath.trim().startsWith('http') 
        ? new URL(redirectForm.targetPath.trim()).pathname 
        : redirectForm.targetPath.trim()

      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      const response: any = await $fetch('/api/admin/seo/redirects', {
        method: 'POST',
        body: {
          sourcePath,
          targetPath,
          code: parseInt(redirectForm.type),
          groupName: redirectForm.groupName || 'دسته‌بندی جدید',
          status: redirectForm.status
        }
      })

      if (response.success) {
        await loadRedirects()
        closeModal()
        showSuccess('ریدایرکت با موفقیت ایجاد شد')
      } else {
        showError('خطا در ایجاد ریدایرکت: ' + response.message)
        return
      }
    }
    updateStats()
  } catch (error: any) { // eslint-disable-line @typescript-eslint/no-explicit-any
    showError('خطا در ذخیره ریدایرکت: ' + error.message)
  }
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const toggleRedirect = async (redirect: any) => {
  try {
    const newStatus = redirect.status === 'active' ? 'inactive' : 'active'
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const response: any = await $fetch(`/api/admin/seo/redirects/${redirect.id}/status`, {
      method: 'PATCH',
      body: { status: newStatus }
    })
    if (response.success) {
      await loadRedirects()
      showSuccess(`ریدایرکت با موفقیت ${newStatus === 'active' ? 'فعال' : 'غیرفعال'} شد`)
    } else {
      showError('خطا در تغییر وضعیت ریدایرکت: ' + response.message)
    }
    updateStats()
  } catch (error: any) { // eslint-disable-line @typescript-eslint/no-explicit-any
    showError('خطا در تغییر وضعیت ریدایرکت: ' + error.message)
  }
}

// تابع بارگذاری ریدایرکت‌ها (در صورت نبود، اضافه شود)
const loadRedirects = async () => {
  try {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const response: any = await $fetch('/api/admin/seo/redirects')
    if (response && Array.isArray(response)) {
      redirects.value = response
    } else if (response && Array.isArray(response.data)) {
      redirects.value = response.data
    } else {
      redirects.value = []
    }
    if (redirects.value.length === 0) {
      currentPage.value = 1;
    }
    updateStats()
  } catch (error: any) { // eslint-disable-line @typescript-eslint/no-explicit-any
    showError('خطا در بارگذاری ریدایرکت‌ها: ' + error.message)
    redirects.value = []
    currentPage.value = 1;
  }
}

// Watch for filter changes to reset pagination
watch([filterGroup, filterType], () => {
  currentPage.value = 1
})

onMounted(() => {
  loadRedirects()
})

// صفحه‌بندی
const handlePageChange = (page: number) => {
  currentPage.value = page
}
const handleItemsPerPageChange = (count: number) => {
  itemsPerPage.value = count
  currentPage.value = 1 // بازگشت به صفحه اول
}

// انتخاب همه
const toggleSelectAll = () => {
  if (selectAll.value) {
    selectedRedirects.value = paginatedRedirects.value.map(r => r.id)
  } else {
    selectedRedirects.value = []
  }
}

// انتخاب تکی
const toggleSelectOne = (id: number) => {
  if (selectedRedirects.value.includes(id)) {
    selectedRedirects.value = selectedRedirects.value.filter(i => i !== id)
  } else {
    selectedRedirects.value.push(id)
  }
}

// باز کردن مدال گروه‌بندی
const openGroupModal = () => {
  groupModalName.value = ''
  showGroupModal.value = true
}

// تغییر گروه گروهی
const changeGroupForSelected = async () => {
  if (!groupModalName.value.trim()) {
    showError('نام گروه را وارد کنید')
    return
  }
  try {
    await $fetch('/api/admin/seo/redirects/group', {
      method: 'PATCH',
      body: {
        ids: selectedRedirects.value,
        group_name: groupModalName.value.trim()
      }
    })
    showSuccess('گروه همه ریدایرکت‌های انتخاب‌شده با موفقیت تغییر کرد')
    closeGroupModal()
    selectedRedirects.value = []
    selectAll.value = false
    await loadRedirects()
  } catch (e: any) { // eslint-disable-line @typescript-eslint/no-explicit-any
    showError('خطا در تغییر گروه: ' + (e.message || e))
  }
}

// حذف گروهی ریدایرکت‌ها
const deleteSelectedRedirects = async () => {
  if (!selectedRedirects.value.length) return;
  if (!confirm('آیا از حذف ریدایرکت‌های انتخاب‌شده مطمئن هستید؟')) return;
  try {
    await $fetch('/api/admin/seo/redirects/bulk-delete', {
      method: 'DELETE',
      body: { ids: selectedRedirects.value },
      headers: { 'Content-Type': 'application/json' }
    });
    showSuccess('ریدایرکت‌های انتخاب‌شده با موفقیت حذف شدند');
    selectedRedirects.value = [];
    selectAll.value = false;
    currentPage.value = 1; // بعد از حذف، صفحه‌بندی به اول برگردد
    await loadRedirects();
  } catch (e: any) { // eslint-disable-line @typescript-eslint/no-explicit-any
    showError('خطا در حذف گروهی: ' + (e.message || e));
  }
}

const deleteRedirect = async (redirect) => {
  if (!redirect?.id) return;
  if (!confirm('آیا از حذف این ریدایرکت مطمئن هستید؟')) return;
  try {
    await $fetch(`/api/admin/seo/redirects/${redirect.id}`, {
      method: 'DELETE',
    });
    showSuccess('ریدایرکت با موفقیت حذف شد');
    currentPage.value = 1; // بعد از حذف، صفحه‌بندی به اول برگردد
    await loadRedirects();
  } catch (e) {
    showError('خطا در حذف ریدایرکت: ' + (e.message || e));
  }
}

// رفع خطا: توابع import/export اولیه برای جلوگیری از خطا
const importRedirects = () => { showError('این قابلیت هنوز پیاده‌سازی نشده است'); }
const exportRedirects = () => { showError('این قابلیت هنوز پیاده‌سازی نشده است'); }
</script>
