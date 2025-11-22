<template>
      <!-- موجودی اصلی و انبارها -->
      <div class="bg-gradient-to-r from-emerald-50 to-teal-50 border border-emerald-200 rounded-xl shadow-lg p-6 text-right mb-6 transition-all duration-300 hover:shadow-xl">
        <div class="flex items-center justify-between cursor-pointer" @click="toggleSection('inventoryMain')">
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 bg-gradient-to-r from-emerald-600 to-teal-600 rounded-lg flex items-center justify-center shadow-md">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
              </svg>
            </div>
            <h3 class="text-lg font-bold text-emerald-800">موجودی اصلی و انبارها</h3>
          </div>
          <span class="text-emerald-600 text-xl font-bold bg-white rounded-full w-8 h-8 flex items-center justify-center shadow-md">{{ sections.inventoryMain ? '−' : '+' }}</span>
        </div>

        <div v-show="sections.inventoryMain" class="mt-6">
          <div class="grid grid-cols-1 gap-6">
            <!-- تنظیمات موجودی -->
            <div v-if="store.inventoryForm.track_inventory" class="grid grid-cols-1 lg:grid-cols-3 gap-6">
              <!-- موجودی کل -->
              <div class="bg-white rounded-xl p-5 shadow-md border border-blue-100 transition-all duration-300 hover:shadow-lg">
                <div class="flex items-center gap-2 mb-3">
                  <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
                  </svg>
                  <label class="text-sm font-bold text-blue-800">موجودی کل</label>
                </div>
                <input v-model="store.inventoryForm.stock_quantity" type="number" class="w-full border-2 border-blue-200 rounded-lg px-4 py-3 text-gray-800 focus:border-blue-500 focus:ring-2 focus:ring-blue-200 transition-all duration-300" min="0" placeholder="0" />
              </div>

              <!-- حداقل موجودی -->
              <div class="bg-white rounded-xl p-5 shadow-md border border-orange-100 transition-all duration-300 hover:shadow-lg">
                <div class="flex items-center gap-2 mb-3">
                  <svg class="w-5 h-5 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"></path>
                  </svg>
                  <label class="text-sm font-bold text-orange-800">حداقل موجودی (هشدار)</label>
                </div>
                <input v-model="store.inventoryForm.min_stock_quantity" type="number" class="w-full border-2 border-orange-200 rounded-lg px-4 py-3 text-gray-800 focus:border-orange-500 focus:ring-2 focus:ring-orange-200 transition-all duration-300" min="0" placeholder="5" />
              </div>

              <!-- حداکثر موجودی -->
              <div class="bg-white rounded-xl p-5 shadow-md border border-green-100 transition-all duration-300 hover:shadow-lg">
                <div class="flex items-center gap-2 mb-3">
                  <svg class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 11l5-5m0 0l5 5m-5-5v12"></path>
                  </svg>
                  <label class="text-sm font-bold text-green-800">حداکثر موجودی</label>
                </div>
                <input v-model="store.inventoryForm.max_stock_quantity" type="number" class="w-full border-2 border-green-200 rounded-lg px-4 py-3 text-gray-800 focus:border-green-500 focus:ring-2 focus:ring-green-200 transition-all duration-300" min="0" placeholder="1000" />
              </div>
            </div>

            <!-- وضعیت موجودی -->
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <label class="flex items-center gap-2">
                <input
                  v-model="store.inventoryForm.track_inventory"
                  type="checkbox"
                  class="checkbox"
                />
                <span class="text-xs text-gray-700 font-semibold">ردیابی موجودی</span>
              </label>
              <label class="flex items-center gap-2">
                <input type="checkbox" class="checkbox" />
                <span class="text-xs text-gray-700 font-semibold">فروش در صورت نبود موجودی</span>
              </label>
              <label class="flex items-center gap-2">
                <input
                  v-model="store.inventoryForm.show_stock_to_customer"
                  type="checkbox"
                  class="checkbox"
                />
                <span class="text-xs text-gray-700 font-semibold">نمایش تعداد موجودی به مشتری</span>
              </label>
              <label class="flex items-center gap-2">
                <input type="checkbox" class="checkbox" checked />
                <span class="text-xs text-gray-700 font-semibold">ارسال هشدار کمبود موجودی</span>
              </label>
            </div>

            <!-- جدول انبارها -->
            <div class="border-t pt-4 border-blue-200">
              <h4 class="text-xs font-semibold text-gray-700 mb-3">توزیع موجودی در انبارها</h4>
              <div class="overflow-x-auto">
                <table class="min-w-full text-xs text-right rtl border-collapse">
                  <thead>
                    <tr class="bg-gray-100 border-b border-blue-200">
                      <th class="px-3 py-2 text-gray-600 font-normal text-right">نام انبار</th>
                      <th class="px-3 py-2 text-gray-600 font-normal text-right">موجودی فعلی</th>
                      <th class="px-3 py-2 text-gray-600 font-normal text-right">حد آستانه</th>
                      <th class="px-3 py-2 text-gray-600 font-normal text-right">وضعیت</th>
                      <th class="px-3 py-2 text-gray-600 font-normal text-right w-20">ویرایش</th>
                      <th class="px-3 py-2 text-gray-600 font-normal text-right w-16">حذف</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr class="border-b hover:bg-gray-50 border-blue-200">
                      <td class="px-3 py-2 text-right">انبار مرکزی</td>
                      <td class="px-3 py-2 text-right">50</td>
                      <td class="px-3 py-2 text-right">5</td>
                      <td class="px-3 py-2 text-right">
                        <span class="bg-green-100 text-green-800 px-2 py-1 rounded text-xs">ایمن</span>
                      </td>
                      <td class="px-3 py-2 text-center">
                        <button class="text-blue-500 hover:text-blue-700">ویرایش</button>
                      </td>
                      <td class="px-3 py-2 text-center">
                        <button class="text-red-500 hover:text-red-700">حذف</button>
                      </td>
                    </tr>
                    <tr class="border-b hover:bg-gray-50 border-blue-200">
                      <td class="px-3 py-2 text-right">انبار فرعی</td>
                      <td class="px-3 py-2 text-right">20</td>
                      <td class="px-3 py-2 text-right">5</td>
                      <td class="px-3 py-2 text-right">
                        <span class="bg-yellow-100 text-yellow-800 px-2 py-1 rounded text-xs">کم</span>
                      </td>
                      <td class="px-3 py-2 text-center">
                        <button class="text-blue-500 hover:text-blue-700">ویرایش</button>
                      </td>
                      <td class="px-3 py-2 text-center">
                        <button class="text-red-500 hover:text-red-700">حذف</button>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>
                </div>
              </div>
            </div>

      <!-- گزارشات و هشدارها -->
      <div class="bg-gradient-to-r from-orange-50 to-amber-50 border border-orange-200 rounded-xl shadow-lg p-6 text-right mb-6 transition-all duration-300 hover:shadow-xl">
        <div class="flex items-center justify-between cursor-pointer" @click="toggleSection('inventoryReports')">
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 bg-gradient-to-r from-orange-600 to-amber-600 rounded-lg flex items-center justify-center shadow-md">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-bold text-orange-800">گزارشات و هشدارها</h3>
          </div>
          <span class="text-orange-600 text-xl font-bold bg-white rounded-full w-8 h-8 flex items-center justify-center shadow-md">{{ sections.inventoryReports ? '−' : '+' }}</span>
        </div>

        <div v-show="sections.inventoryReports" class="mt-4">
          <div class="grid grid-cols-1 gap-6">
            <!-- آمارهای کلی -->
            <div class="grid grid-cols-2 lg:grid-cols-4 gap-6">
              <div class="bg-white rounded-xl p-5 shadow-md border border-blue-100 text-center transition-all duration-300 hover:shadow-lg hover:scale-105">
                <div class="flex justify-center mb-3">
                  <div class="w-12 h-12 bg-gradient-to-r from-blue-500 to-blue-600 rounded-full flex items-center justify-center">
                    <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
                    </svg>
                  </div>
                </div>
                <div class="text-2xl font-bold text-blue-600 mb-1">70</div>
                <div class="text-sm text-gray-600 font-medium">موجودی کل</div>
              </div>
              <div class="bg-white rounded-xl p-5 shadow-md border border-emerald-100 text-center transition-all duration-300 hover:shadow-lg hover:scale-105">
                <div class="flex justify-center mb-3">
                  <div class="w-12 h-12 bg-gradient-to-r from-emerald-500 to-emerald-600 rounded-full flex items-center justify-center">
                    <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"></path>
                    </svg>
                  </div>
                </div>
                <div class="text-2xl font-bold text-emerald-600 mb-1">7</div>
                <div class="text-sm text-gray-600 font-medium">رزرو شده</div>
              </div>
              <div class="bg-white rounded-xl p-5 shadow-md border border-amber-100 text-center transition-all duration-300 hover:shadow-lg hover:scale-105">
                <div class="flex justify-center mb-3">
                  <div class="w-12 h-12 bg-gradient-to-r from-amber-500 to-amber-600 rounded-full flex items-center justify-center">
                    <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
                    </svg>
                  </div>
                </div>
                <div class="text-2xl font-bold text-amber-600 mb-1">63</div>
                <div class="text-sm text-gray-600 font-medium">قابل فروش</div>
              </div>
              <div class="bg-white rounded-xl p-5 shadow-md border border-red-100 text-center transition-all duration-300 hover:shadow-lg hover:scale-105">
                <div class="flex justify-center mb-3">
                  <div class="w-12 h-12 bg-gradient-to-r from-red-500 to-red-600 rounded-full flex items-center justify-center">
                    <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"></path>
                    </svg>
                  </div>
                </div>
                <div class="text-2xl font-bold text-red-600 mb-1">0</div>
                <div class="text-sm text-gray-600 font-medium">کمبود</div>
              </div>
            </div>

            <!-- تنظیمات هشدار -->
            <div class="border-t pt-4 border-blue-200">
              <h4 class="text-xs font-semibold text-gray-700 mb-3">تنظیمات هشدار</h4>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div class="flex flex-col gap-2">
                  <label class="text-xs text-gray-700 font-semibold">دریافت هشدار در</label>
                  <select class="w-full text-sm px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-orange-500 transition-all duration-200 bg-white text-gray-900">
                    <option value="email" class="text-gray-900 bg-white py-2">ایمیل</option>
                    <option value="sms" class="text-gray-900 bg-white py-2">پیامک</option>
                    <option value="both" class="text-gray-900 bg-white py-2">هر دو</option>
                  </select>
                </div>
                <div class="flex flex-col gap-2">
                  <label class="text-xs text-gray-700 font-semibold">دریافت‌کننده هشدار</label>
                  <input type="email" class="w-full text-sm px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-orange-500 transition-all duration-200 bg-white text-gray-900" placeholder="admin@example.com" />
                </div>
              </div>
                </div>

            <!-- سابقه تغییرات -->
            <div class="border-t pt-4 border-blue-200">
              <h4 class="text-xs font-semibold text-gray-700 mb-3">آخرین تغییرات موجودی</h4>
              <div class="overflow-x-auto">
                <table class="min-w-full text-xs text-right rtl border-collapse">
                  <thead>
                    <tr class="bg-gray-100 border-b border-blue-200">
                      <th class="px-3 py-2 text-gray-600 font-normal text-right">تاریخ</th>
                      <th class="px-3 py-2 text-gray-600 font-normal text-right">نوع تغییر</th>
                      <th class="px-3 py-2 text-gray-600 font-normal text-right">تعداد</th>
                      <th class="px-3 py-2 text-gray-600 font-normal text-right">علت</th>
                      <th class="px-3 py-2 text-gray-600 font-normal text-right">کاربر</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr class="border-b hover:bg-gray-50 border-blue-200">
                      <td class="px-3 py-2 text-right">1403/08/15</td>
                      <td class="px-3 py-2 text-right">
                        <span class="bg-green-100 text-green-800 px-2 py-1 rounded text-xs">افزایش</span>
                      </td>
                      <td class="px-3 py-2 text-right">+20</td>
                      <td class="px-3 py-2 text-right">خرید جدید</td>
                      <td class="px-3 py-2 text-right">مدیر انبار</td>
                    </tr>
                    <tr class="border-b hover:bg-gray-50 border-blue-200">
                      <td class="px-3 py-2 text-right">1403/08/14</td>
                      <td class="px-3 py-2 text-right">
                        <span class="bg-red-100 text-red-800 px-2 py-1 rounded text-xs">کاهش</span>
                      </td>
                      <td class="px-3 py-2 text-right">-3</td>
                      <td class="px-3 py-2 text-right">فروش</td>
                      <td class="px-3 py-2 text-right">سیستم</td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>
          </div>
        </div>
                </div>

      <!-- حرکات موجودی -->
      <div class="bg-gradient-to-r from-purple-50 to-pink-50 border border-purple-200 rounded-xl shadow-lg p-6 text-right mb-6 transition-all duration-300 hover:shadow-xl">
        <div class="flex items-center justify-between cursor-pointer" @click="toggleSection('inventoryMovements')">
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 bg-gradient-to-r from-purple-600 to-pink-600 rounded-lg flex items-center justify-center shadow-md">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7h12m0 0l-4-4m4 4l-4 4m0 6H4m0 0l4 4m-4-4l4-4"></path>
              </svg>
            </div>
            <h3 class="text-lg font-bold text-purple-800">حرکات موجودی و انتقالات</h3>
          </div>
          <span class="text-purple-600 text-xl font-bold bg-white rounded-full w-8 h-8 flex items-center justify-center shadow-md">{{ sections.inventoryMovements ? '−' : '+' }}</span>
        </div>

        <div v-show="sections.inventoryMovements" class="mt-4">
          <div class="grid grid-cols-1 gap-6">
            <!-- فرم حرکت جدید -->
            <div class="border border-blue-200 rounded p-3 bg-gray-50">
              <h4 class="text-xs font-semibold text-gray-700 mb-3">ثبت حرکت موجودی</h4>
              <div class="grid grid-cols-1 md:grid-cols-4 gap-2">
                <div class="flex flex-col gap-1">
                  <label class="text-xs text-gray-700 font-semibold">نوع حرکت</label>
                  <select class="w-full text-xs px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500 transition-all duration-200 bg-white text-gray-900">
                    <option value="in" class="text-gray-900 bg-white py-2">ورود</option>
                    <option value="out" class="text-gray-900 bg-white py-2">خروج</option>
                    <option value="transfer" class="text-gray-900 bg-white py-2">انتقال</option>
                    <option value="adjustment" class="text-gray-900 bg-white py-2">تعدیل</option>
                  </select>
                </div>
                <div class="flex flex-col gap-1">
                  <label class="text-xs text-gray-700 font-semibold">تعداد</label>
                  <input type="number" class="w-full text-xs px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500 transition-all duration-200 bg-white text-gray-900" min="1" placeholder="1" />
                </div>
                <div class="flex flex-col gap-1">
                  <label class="text-xs text-gray-700 font-semibold">انبار</label>
                  <select class="w-full text-xs px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500 transition-all duration-200 bg-white text-gray-900">
                    <option value="main" class="text-gray-900 bg-white py-2">انبار مرکزی</option>
                    <option value="secondary" class="text-gray-900 bg-white py-2">انبار فرعی</option>
                  </select>
                </div>
                <div class="flex items-end">
                  <button class="bg-blue-600 text-white rounded px-3 py-1 text-xs hover:bg-blue-700 transition w-full">ثبت</button>
                </div>
              </div>
              <div class="mt-2">
                <label class="text-xs text-gray-700 font-semibold">علت</label>
                <input type="text" class="w-full text-xs px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500 transition-all duration-200 bg-white text-gray-900 mt-1" placeholder="دلیل تغییر موجودی..." />
              </div>
                </div>

            <!-- جدول حرکات -->
            <div class="overflow-x-auto">
              <table class="min-w-full text-xs text-right rtl border-collapse">
                <thead>
                  <tr class="bg-gray-100 border-b border-blue-200">
                    <th class="px-3 py-2 text-gray-600 font-normal text-right">تاریخ/ساعت</th>
                    <th class="px-3 py-2 text-gray-600 font-normal text-right">نوع</th>
                    <th class="px-3 py-2 text-gray-600 font-normal text-right">انبار</th>
                    <th class="px-3 py-2 text-gray-600 font-normal text-right">مقدار</th>
                    <th class="px-3 py-2 text-gray-600 font-normal text-right">باقی‌مانده</th>
                    <th class="px-3 py-2 text-gray-600 font-normal text-right">علت</th>
                    <th class="px-3 py-2 text-gray-600 font-normal text-right">کاربر</th>
                  </tr>
                </thead>
                <tbody>
                  <tr class="border-b hover:bg-gray-50 border-blue-200">
                    <td class="px-3 py-2 text-right">1403/08/15 14:30</td>
                    <td class="px-3 py-2 text-right">
                      <span class="bg-green-100 text-green-800 px-2 py-1 rounded text-xs">ورود</span>
                    </td>
                    <td class="px-3 py-2 text-right">انبار مرکزی</td>
                    <td class="px-3 py-2 text-right">+20</td>
                    <td class="px-3 py-2 text-right">70</td>
                    <td class="px-3 py-2 text-right">خرید از تامین‌کننده</td>
                    <td class="px-3 py-2 text-right">مدیر انبار</td>
                  </tr>
                  <tr class="border-b hover:bg-gray-50 border-blue-200">
                    <td class="px-3 py-2 text-right">1403/08/14 16:45</td>
                    <td class="px-3 py-2 text-right">
                      <span class="bg-red-100 text-red-800 px-2 py-1 rounded text-xs">خروج</span>
                    </td>
                    <td class="px-3 py-2 text-right">انبار مرکزی</td>
                    <td class="px-3 py-2 text-right">-3</td>
                    <td class="px-3 py-2 text-right">50</td>
                    <td class="px-3 py-2 text-right">فروش آنلاین</td>
                    <td class="px-3 py-2 text-right">سیستم</td>
                  </tr>
                </tbody>
              </table>
            </div>
                </div>
              </div>
            </div>

      <!-- پیش‌بینی و برنامه‌ریزی -->
      <div class="bg-gradient-to-r from-indigo-50 to-blue-50 border border-indigo-200 rounded-xl shadow-lg p-6 text-right transition-all duration-300 hover:shadow-xl">
        <div class="flex items-center justify-between cursor-pointer" @click="toggleSection('inventoryPlanning')">
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 bg-gradient-to-r from-indigo-600 to-blue-600 rounded-lg flex items-center justify-center shadow-md">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-bold text-indigo-800">پیش‌بینی و برنامه‌ریزی موجودی</h3>
          </div>
          <span class="text-indigo-600 text-xl font-bold bg-white rounded-full w-8 h-8 flex items-center justify-center shadow-md">{{ sections.inventoryPlanning ? '−' : '+' }}</span>
        </div>

        <div v-show="sections.inventoryPlanning" class="mt-4">
          <div class="grid grid-cols-1 gap-6">
            <!-- تنظیمات پیش‌بینی -->
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div class="flex flex-col gap-2">
                <label class="text-xs text-gray-700 font-semibold">نوع پیش‌بینی</label>
                <select class="w-full text-sm px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-all duration-200 bg-white text-gray-900">
                  <option value="manual" class="text-gray-900 bg-white py-2">دستی</option>
                  <option value="auto" class="text-gray-900 bg-white py-2">خودکار (بر اساس فروش)</option>
                  <option value="seasonal" class="text-gray-900 bg-white py-2">فصلی</option>
                </select>
                </div>
              <div class="flex flex-col gap-2">
                <label class="text-xs text-gray-700 font-semibold">دوره بررسی</label>
                <select class="w-full text-sm px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-all duration-200 bg-white text-gray-900">
                  <option value="weekly" class="text-gray-900 bg-white py-2">هفتگی</option>
                  <option value="monthly" class="text-gray-900 bg-white py-2">ماهانه</option>
                  <option value="quarterly" class="text-gray-900 bg-white py-2">فصلی</option>
                </select>
              </div>
            </div>

            <!-- آمار فروش -->
            <div class="border-t pt-4 border-blue-200">
              <h4 class="text-xs font-semibold text-gray-700 mb-3">آمار فروش و مصرف</h4>
              <div class="grid grid-cols-2 md:grid-cols-4 gap-6">
                <div class="bg-gray-50 p-3 rounded border border-blue-200 text-center">
                  <div class="text-lg font-bold text-blue-600">15</div>
                  <div class="text-xs text-gray-600">فروش این ماه</div>
                </div>
                <div class="bg-green-50 p-3 rounded border border-blue-200 text-center">
                  <div class="text-lg font-bold text-green-600">12</div>
                  <div class="text-xs text-gray-600">میانگین ماهانه</div>
                </div>
                <div class="bg-yellow-50 p-3 rounded border border-blue-200 text-center">
                  <div class="text-lg font-bold text-yellow-600">45</div>
                  <div class="text-xs text-gray-600">روزهای تکفی</div>
                </div>
                <div class="bg-purple-50 p-3 rounded border border-blue-200 text-center">
                  <div class="text-lg font-bold text-purple-600">25</div>
                  <div class="text-xs text-gray-600">نقطه سفارش</div>
                </div>
              </div>
            </div>

            <!-- تنظیمات خودکار -->
            <div class="border-t pt-4 border-blue-200">
              <h4 class="text-xs font-semibold text-gray-700 mb-3">سفارش خودکار</h4>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-2">
                <label class="flex items-center gap-2">
                  <input type="checkbox" class="checkbox" />
                  <span class="text-xs text-gray-700">فعال‌سازی سفارش خودکار</span>
                </label>
                <label class="flex items-center gap-2">
                  <input type="checkbox" class="checkbox" />
                  <span class="text-xs text-gray-700">هشدار پیش از رسیدن به نقطه سفارش</span>
                </label>
              </div>
              <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mt-3">
                <div class="flex flex-col gap-2">
                  <label class="text-xs text-gray-700 font-semibold">نقطه سفارش</label>
                  <input type="number" class="w-full text-sm px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-all duration-200 bg-white text-gray-900" min="0" placeholder="25" />
                </div>
                <div class="flex flex-col gap-2">
                  <label class="text-xs text-gray-700 font-semibold">مقدار سفارش</label>
                  <input type="number" class="w-full text-sm px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-all duration-200 bg-white text-gray-900" min="1" placeholder="50" />
                </div>
                <div class="flex flex-col gap-2">
                  <label class="text-xs text-gray-700 font-semibold">تامین‌کننده</label>
                  <select class="w-full text-sm px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-all duration-200 bg-white text-gray-900">
                    <option value="main" class="text-gray-900 bg-white py-2">تامین‌کننده اصلی</option>
                    <option value="secondary" class="text-gray-900 bg-white py-2">تامین‌کننده دوم</option>
                  </select>
                </div>
              </div>
            </div>

            <!-- خلاصه برنامه‌ریزی -->
            <div class="border-t pt-4 border-blue-200">
              <div class="bg-gray-50 p-3 rounded border border-blue-200">
                <h4 class="text-xs font-semibold text-gray-700 mb-2">خلاصه وضعیت موجودی</h4>
                <div class="grid grid-cols-1 md:grid-cols-2 gap-2 text-xs">
                  <div>
                    <span class="text-gray-600">وضعیت فعلی:</span>
                    <span class="font-medium text-green-600"> مطلوب</span>
                  </div>
                  <div>
                    <span class="text-gray-600">تاریخ پیش‌بینی کمبود:</span>
                    <span class="font-medium text-blue-600"> 1403/10/15</span>
                  </div>
                  <div>
                    <span class="text-gray-600">پیشنهاد سفارش:</span>
                    <span class="font-medium text-orange-600"> 50 عدد تا 1403/09/30</span>
                  </div>
                  <div>
                    <span class="text-gray-600">اولویت:</span>
                    <span class="font-medium text-blue-600"> متوسط</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
</script>

<script setup lang="ts">
import { reactive } from 'vue';

import { useProductCreateStore } from '~/stores/productCreate';

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

const store = useProductCreateStore()

const sections = reactive({
  inventoryMain: true,
  inventoryReports: false,
  inventoryMovements: false,
  inventoryPlanning: false
})

const toggleSection = (section: keyof typeof sections) => {
  sections[section] = !sections[section]
}
</script>

<style scoped>
/* Custom Select Styles */
select {
  appearance: none;
  -webkit-appearance: none;
  -moz-appearance: none;
  background-image: url("data:image/svg+xml;charset=US-ASCII,<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 4 5'><path fill='%23666' d='M2 0L0 2h4zm0 5L0 3h4z'/></svg>");
  background-repeat: no-repeat;
  background-position: left 8px center;
  background-size: 12px;
  padding-left: 28px;
}

select:focus {
  outline: none;
  box-shadow: 0 0 0 3px rgba(147, 51, 234, 0.1);
}

/* Option Styles */
select option {
  background-color: white !important;
  color: #1f2937 !important;
  padding: 8px 12px;
  font-weight: normal;
}

select option:checked,
select option:hover {
  background-color: #f3f4f6 !important;
  color: #1f2937 !important;
}

/* Animation for cards */
.transition-all {
  transition: all 0.3s ease-in-out;
}

.hover\:transform:hover {
  transform: scale(1.02);
}

.shadow-lg {
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05);
}

.hover\:shadow-xl:hover {
  box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04);
}
</style>
