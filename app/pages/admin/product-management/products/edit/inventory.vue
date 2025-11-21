<template v-else-if="activeTab === 'inventory'">
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
            <!-- کارت‌های حداقل/حداکثر موجودی حذف شدند؛ مقادیر آستانه فقط در جدول انبار مدیریت می‌شوند. -->

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
                <input
                  v-model="store.inventoryForm.allow_reservation"
                  type="checkbox"
                  class="checkbox"
                />
                <span class="text-xs text-gray-700 font-semibold">اجازه رزرو (در صورت فعال)</span>
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
                <div class="grid grid-cols-1 md:grid-cols-3 gap-3 mb-3">
                  <div class="md:col-span-3 flex items-center gap-2">
                    <label class="flex items-center gap-2 text-xs text-gray-700 font-semibold">
                      <input v-model="store.inventoryForm.shipping_enabled" type="checkbox" />
                      <span>ارسال از انبار پیش‌فرض</span>
                    </label>
                  </div>
                </div>
                <table class="min-w-full text-sm text-right rtl border-collapse">
                  <thead>
                    <tr class="bg-gray-100 border-b border-blue-200 text-sm">
                      <th class="px-3 py-3 text-gray-600 font-medium text-right">نام انبار</th>
                      <th class="px-3 py-3 text-gray-600 font-medium text-right">موجودی فعلی</th>
                      <th class="px-3 py-3 text-gray-600 font-medium text-right">حد آستانه</th>
                      <th class="px-3 py-3 text-gray-600 font-medium text-right">حداکثر موجودی</th>
                      <th class="px-3 py-3 text-gray-600 font-medium text-right">وضعیت</th>
                      <th class="px-3 py-3 text-gray-600 font-medium text-right">عملیات</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr v-for="row in tableRows" :key="row.warehouse_id" class="border-b hover:bg-gray-50 border-blue-200">
                      <td class="px-3 py-2 text-right">{{ row.warehouse_name || warehouseName(row.warehouse_id) }}</td>
                      <td class="px-3 py-2 text-right">
                        <template v-if="editingId === row.warehouse_id">
                          <input v-model.number="editBuffer.quantity" type="number" class="w-24 border-2 border-white rounded px-2 py-1 text-right" />
                        </template>
                        <template v-else>
                          {{ row.quantity }}
                        </template>
                      </td>
                      <td class="px-3 py-2 text-right">
                        <template v-if="editingId === row.warehouse_id">
                          <input v-model.number="editBuffer.min_qty" type="number" class="w-24 border-2 border-white rounded px-2 py-1 text-right" />
                        </template>
                        <template v-else>
                          {{ row.min_qty || 0 }}
                        </template>
                      </td>
                      <td class="px-3 py-2 text-right">
                        <template v-if="editingId === row.warehouse_id">
                          <input v-model.number="editBuffer.max_qty" type="number" class="w-24 border-2 border-white rounded px-2 py-1 text-right" />
                        </template>
                        <template v-else>
                          {{ row.max_qty || 0 }}
                        </template>
                      </td>
                      <td class="px-3 py-2 text-right">
                        <span :class="getStockStatusClass(row)" class="px-2 py-1 rounded text-xs">
                          {{ getStockStatusText(row) }}
                        </span>
                      </td>
                      <td class="px-3 py-2 text-left">
                        <template v-if="editingId === row.warehouse_id">
                          <button class="text-white bg-emerald-600 rounded px-2 py-1 text-xs ml-1" @click="saveRow(row.warehouse_id)">ذخیره</button>
                          <button class="text-gray-700 border border-gray-300 rounded px-2 py-1 text-xs" @click="cancelEdit">انصراف</button>
                        </template>
                        <template v-else>
                          <button class="text-blue-600" @click="startEdit(row)">ویرایش</button>
                        </template>
                      </td>
                    </tr>
                  </tbody>
                </table>
                <div class="mt-3 flex gap-2"></div>

                
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
                <div class="text-2xl font-bold text-blue-600 mb-1">{{ totalQty }}</div>
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
                <div class="text-2xl font-bold text-emerald-600 mb-1">{{ totalReserved }}</div>
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
                <div class="text-2xl font-bold text-amber-600 mb-1">{{ totalQty - totalReserved }}</div>
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
                <div class="text-2xl font-bold text-red-600 mb-1">{{ Math.max(0, totalReserved - totalQty) }}</div>
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


<script setup lang="ts">
import { computed, reactive, ref, watch } from 'vue'
import { useProductCreateStore } from '~/stores/productCreate'

interface Warehouse {
  id: number | string
  code: string
  name: string
  city?: string
  is_default?: boolean
  IsDefault?: boolean
  isDefault?: boolean
  priority?: number
}

interface WarehouseStockRaw {
  warehouse_id?: number
  WarehouseID?: number
  warehouseId?: number
  warehouse_name?: string
  WarehouseName?: string
  warehouseName?: string
  quantity?: number
  Quantity?: number
  reserved?: number
  Reserved?: number
  min_qty?: number
  MinQty?: number
  max_qty?: number
  MaxQty?: number
}

interface StockRow {
  warehouse_id: number
  warehouse_name: string
  quantity: number
  reserved: number
  min_qty: number
  max_qty: number
}

// Store for inventory data
const store = useProductCreateStore()

// Collapsible sections
const sections = reactive({
  inventoryMain: true,
  inventoryReports: false,
  inventoryMovements: false,
  inventoryPlanning: false
})

const toggleSection = (section) => {
  sections[section] = !sections[section]
}

// دریافت لیست انبارها برای نمایش در انتخاب‌گر (بدون await useAsyncData برای سادگی)
const warehouses = ref<Warehouse[]>([])
try {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  warehouses.value = await $fetch<any[]>('/api/admin/warehouses')
  // اگر هیچ انباری وجود ندارد، یک انبار پیشفرض ایجاد کن
  if (warehouses.value.length === 0) {
    try {
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      const defaultWarehouse = await $fetch<any>('/api/admin/warehouses', {
        method: 'POST',
        body: {
          code: 'DEFAULT',
          name: 'انبار پیشفرض',
          city: 'تهران',
          is_default: true,
          priority: 100
        }
      })
      warehouses.value = [defaultWarehouse]
    } catch {
      // حتی اگر ایجاد انبار شکست خورد، حداقل یک fallback داشته باش
      warehouses.value = [{
        id: 1,
        code: 'DEFAULT',
        name: 'انبار پیشفرض',
        is_default: true
      }]
    }
  }
} catch {
  warehouses.value = []
}

// وضعیت قابل ویرایش برای هر انبار
const stocks = ref<StockRow[]>([])
const tableRows = computed(() => {
  let normalized = (stocks.value || []).map((r) => ({
    warehouse_id: r.warehouse_id,
    warehouse_name: r.warehouse_name, // اضافه کردن نام انبار
    quantity: Number(r.quantity ?? 0),
    reserved: Number(r.reserved ?? 0),
    min_qty: Number(r.min_qty ?? 0),
    max_qty: Number(r.max_qty ?? 0)
  }))

  // اگر هیچ انباری وجود ندارد، حداقل انبار پیشفرض را اضافه کن
  if (normalized.length === 0 && warehouses.value.length > 0) {
    const defaultWarehouse = warehouses.value.find(w => w.is_default)
    if (defaultWarehouse) {
      normalized = [{
        warehouse_id: Number(defaultWarehouse.id),
        warehouse_name: defaultWarehouse.name, // اضافه کردن نام انبار پیشفرض
        quantity: Number(store.inventoryForm?.stock_quantity || 0),
        reserved: 0,
        min_qty: 0,
        max_qty: 0
      }]
    }
  }

  return normalized
})

// موجودی کل (جمع روی انبارها) — فقط خواندنی
const totalQty = computed(() => (tableRows.value || []).reduce((a, r) => a + Number(r.quantity || 0), 0))
const totalReserved = computed(() => (tableRows.value || []).reduce((a, r) => a + Number(r.reserved || 0), 0))
// همگام‌سازی فرم برای استفاده‌های دیگر (مثلاً گزارشات)
watch(totalQty, (qty) => {
  store.inventoryForm.stock_quantity = Number(qty)
}, { immediate: true })
async function loadStocks(productId: string | number) {
  try {
    const res = await $fetch<WarehouseStockRaw[] | { data: WarehouseStockRaw[] }>(`/api/product-warehouse-stocks/${productId}?_=${Date.now()}`)
    const arr: WarehouseStockRaw[] = (res && 'data' in res && Array.isArray(res.data) ? res.data : (Array.isArray(res) ? res : []))
    
    stocks.value = arr.map((r) => ({
      warehouse_id: Number(r.warehouse_id || r.WarehouseID || r.warehouseId || 0),
      warehouse_name: r.warehouse_name || r.WarehouseName || r.warehouseName || '',
      quantity: Number(r.quantity || r.Quantity || 0),
      reserved: Number(r.reserved || r.Reserved || 0),
      min_qty: Number(r.min_qty || r.MinQty || 0),
      max_qty: Number(r.max_qty || r.MaxQty || 0)
    }))
  } catch (error) {
    console.error('خطا در بارگذاری موجودی انبارها:', error)
    // اگر خطا رخ داد، stocks را خالی کن
    stocks.value = []
  }
}

// انبار پیش‌فرض خرید
const defaultWarehouseId = ref(null)
// const selectedWarehouseName = computed(() => {
//   const w = (warehouses.value || []).find((x) => Number(x.id) === Number(defaultWarehouseId.value || 0))
//   return w ? w.name : '—'
// })

function warehouseName(id: number | string) {
  // ابتدا از نام انبار موجود در stocks استفاده کن
  const stock = stocks.value.find(s => s.warehouse_id === Number(id))
  if (stock && stock.warehouse_name) {
    return stock.warehouse_name
  }

  // اگر در stocks نبود، از لیست انبارها پیدا کن
  const w = (warehouses.value || []).find(x => Number(x.id) === Number(id))
  if (w) return w.name

  // اگر انبار پیدا نشد، چک کن آیا انبار پیشفرض است
  const defaultWarehouse = (warehouses.value || []).find(x => x.is_default)
  if (defaultWarehouse && Number(defaultWarehouse.id) === Number(id)) {
    return defaultWarehouse.name || 'انبار پیشفرض'
  }

  // fallback
  return `انبار ${id}`
}

const editingId = ref<number | null>(null)
const editBuffer = reactive({ quantity: 0, min_qty: 0, max_qty: 0 })
function startEdit(row: StockRow) {
  // اگر این انبار در stocks نیست، آن را اضافه کن
  const existingStockIndex = stocks.value.findIndex(s => s.warehouse_id === row.warehouse_id)
  if (existingStockIndex === -1) {
    stocks.value.push({
      warehouse_id: row.warehouse_id,
      warehouse_name: row.warehouse_name, // حفظ نام انبار
      quantity: Number(row.quantity || 0),
      reserved: 0,
      min_qty: Number(row.min_qty || 0),
      max_qty: Number(row.max_qty || 0)
    })
  }

  editingId.value = row.warehouse_id
  // استفاده از مقادیر واقعی از row به جای 0
  editBuffer.quantity = Number(row.quantity) || 0
  editBuffer.min_qty = Number(row.min_qty) || 0
  editBuffer.max_qty = Number(row.max_qty) || 0
}
function cancelEdit() {
  // اگر انباری که در حال ویرایش بود در stocks وجود نداشت، آن را حذف کن
  if (editingId.value) {
    const existingStockIndex = stocks.value.findIndex(s => s.warehouse_id === editingId.value)
    if (existingStockIndex !== -1) {
      const stock = stocks.value[existingStockIndex]
      // اگر این انبار جدید اضافه شده بود (یعنی قبلاً در stocks نبود)، آن را حذف کن
      if (!stock.quantity && !stock.min_qty) {
        stocks.value.splice(existingStockIndex, 1)
      }
    }
  }
  editingId.value = null
}
async function saveRow(warehouseId: number) {
  if (!store.editingProductId) return

  try {
    // اگر این انبار در stocks نیست، آن را اضافه کن
    const existingStockIndex = stocks.value.findIndex(s => s.warehouse_id === warehouseId)
    if (existingStockIndex === -1) {
      // پیدا کردن نام انبار از لیست انبارها
      const warehouse = warehouses.value.find(w => Number(w.id) === warehouseId)
      stocks.value.push({
        warehouse_id: warehouseId,
        warehouse_name: warehouse ? warehouse.name : `انبار ${warehouseId}`, // اضافه کردن نام انبار
        quantity: Number(editBuffer.quantity ?? 0),
        reserved: 0,
        min_qty: Number(editBuffer.min_qty ?? 0),
        max_qty: Number(editBuffer.max_qty ?? 0)
      })
    } else {
      // اگر وجود دارد، به‌روزرسانی کن
      stocks.value[existingStockIndex].quantity = Number(editBuffer.quantity ?? 0)
      stocks.value[existingStockIndex].min_qty = Number(editBuffer.min_qty ?? 0)
      stocks.value[existingStockIndex].max_qty = Number(editBuffer.max_qty ?? 0)
    }

    const entries = [{ warehouse_id: Number(warehouseId), quantity: Number(editBuffer.quantity ?? 0), min_qty: Number(editBuffer.min_qty ?? 0), max_qty: Number(editBuffer.max_qty ?? 0) }]
     
    await $fetch(`/api/product-warehouse-stocks/${store.editingProductId}`, { method: 'PUT', body: { entries, default_warehouse_id: defaultWarehouseId.value } })
    await loadStocks(store.editingProductId)
    // بارگذاری مجدد موجودی کلی بعد از ذخیره انبارها
    await loadInventoryData(store.editingProductId)
    // مجموع موجودی کل در هدر به صورت computed از stocks نمایش داده می‌شود
    editingId.value = null

    // همچنین موجودی کلی را در store به‌روزرسانی می‌کنیم
    store.inventoryForm.stock_quantity = totalQty.value
  } catch (error) {
    console.error('خطا در ذخیره موجودی انبار:', error)
    // در صورت خطا، editingId را reset کن تا حالت ویرایش بسته شود
    editingId.value = null
  }
}

async function loadInventoryData(productId: string | number) {
  try {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const response = await $fetch<any>(`/api/product-inventories/${productId}`)
    if (response) {
      store.inventoryForm.stock_quantity = response.stock_quantity || 0
      store.inventoryForm.min_stock_quantity = response.min_stock_quantity || 0
      store.inventoryForm.max_stock_quantity = response.max_stock_quantity || 0
      store.inventoryForm.stock_status = response.stock_status || 'in_stock'
      store.inventoryForm.show_stock_to_customer = !!response.show_stock_to_customer
      store.inventoryForm.track_inventory = response.track_inventory !== undefined ? !!response.track_inventory : true
      store.inventoryForm.allow_reservation = !!response.allow_reservation
      // اگر بک‌اند انبار پیش‌فرض را برگرداند، همان را ست کن؛ وگرنه 1st warehouse
      if (response.default_warehouse_id) {
        store.inventoryForm.warehouse_id = Number(response.default_warehouse_id)
        store.inventoryForm.shipping_enabled = true
      }
    }
  } catch (error: unknown) {
    console.error('خطا در بارگذاری اطلاعات موجودی:', error)
    // اگر رکورد موجودی هنوز ایجاد نشده باشد، 404 حالت طبیعی است
    const err = error as { status?: number; response?: { status?: number; statusCode?: number } }
    const status = err?.status || err?.response?.status || err?.response?.statusCode
    if (status === 404) {
      // مقداردهی پیش‌فرض برای فرم موجودی
      store.inventoryForm.stock_quantity = 0
      store.inventoryForm.min_stock_quantity = 0
      store.inventoryForm.max_stock_quantity = 0
      store.inventoryForm.stock_status = 'in_stock'
      store.inventoryForm.show_stock_to_customer = false
      store.inventoryForm.track_inventory = true
    }
  }
}

// توابع وضعیت موجودی
function getStockStatusClass(row) {
  const quantity = Number(row.quantity || 0)
  const minQty = Number(row.min_qty || 0)
  const maxQty = Number(row.max_qty || 0)
  
  if (quantity <= minQty) {
    return 'bg-yellow-100 text-yellow-800'
  } else if (maxQty > 0 && quantity >= maxQty) {
    return 'bg-red-100 text-red-800'
  } else {
    return 'bg-green-100 text-green-800'
  }
}

function getStockStatusText(row) {
  const quantity = Number(row.quantity || 0)
  const minQty = Number(row.min_qty || 0)
  const maxQty = Number(row.max_qty || 0)
  
  if (quantity <= minQty) {
    return 'کم'
  } else if (maxQty > 0 && quantity >= maxQty) {
    return 'بیش از حد'
  } else {
    return 'ایمن'
  }
}

function getDefaultWarehouseId() {
  const list = (warehouses.value || [])
  const def = list.find((w) => w.is_default || w.IsDefault || w.isDefault)
  return def ? Number(def.id) : (list[0] ? Number(list[0].id) : null)
}

// اگر انبار پیش‌فرض محصول از بک‌اند برگردد، انتخاب را ست کنیم؛ در غیراین‌صورت انبار پیش‌فرض سیستم یا اولین انبار
async function loadDefaultWarehouse(productId) {
  try {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const inv = await $fetch<any>(`/api/product-inventories/${productId}`)
    if (inv && inv.default_warehouse_id) {
      defaultWarehouseId.value = Number(inv.default_warehouse_id)
      store.inventoryForm.warehouse_id = Number(inv.default_warehouse_id)
    } else if ((warehouses.value || []).length > 0) {
      const d = getDefaultWarehouseId()
      defaultWarehouseId.value = d
      store.inventoryForm.warehouse_id = d
    }
  } catch (error) {
    console.error('خطا در بارگذاری انبار پیشفرض:', error)
    // در صورت خطا، انبار پیشفرض سیستم را انتخاب کن
    if ((warehouses.value || []).length > 0) {
      const d = getDefaultWarehouseId()
      defaultWarehouseId.value = d
      store.inventoryForm.warehouse_id = d
    }
  }
}

// اگر موجودی کل > 0 است ولی هیچ ردیف انباری نداریم، یکبار ردیف انبار پیش‌فرض را ایجاد/همگام می‌کنیم
const didEnsureDefaultRow = ref(false)
async function ensureDefaultRow(productId: string | number) {
  try {
    // اگر defaultWarehouseId تنظیم نشده، سعی کن آن را تنظیم کن
    if (!defaultWarehouseId.value && warehouses.value.length > 0) {
      const d = getDefaultWarehouseId()
      if (d) {
        defaultWarehouseId.value = d
        store.inventoryForm.warehouse_id = d
      }
    }

    // اگر stocks خالی است اما انبارها وجود دارند، حداقل انبار پیشفرض را اضافه کن
    if (stocks.value.length === 0 && warehouses.value.length > 0 && !didEnsureDefaultRow.value) {
      const defaultWarehouse = warehouses.value.find(w => w.is_default)
      if (defaultWarehouse && defaultWarehouseId.value) {
        stocks.value = [{
          warehouse_id: Number(defaultWarehouseId.value),
          warehouse_name: defaultWarehouse.name, // اضافه کردن نام انبار پیشفرض
          quantity: Number(store.inventoryForm?.stock_quantity || 0),
          reserved: 0,
          min_qty: Number(store.inventoryForm?.min_stock_quantity || 0),
          max_qty: Number(store.inventoryForm?.max_stock_quantity || 0)
        }]
        didEnsureDefaultRow.value = true
      }
    }

    if (
      !didEnsureDefaultRow.value &&
      Array.isArray(stocks.value) &&
      Number(store.inventoryForm?.stock_quantity || 0) > 0 &&
      Number(defaultWarehouseId.value || 0) > 0 &&
      !stocks.value.find(s => s.warehouse_id === Number(defaultWarehouseId.value))
    ) {
      try {
         
        await $fetch('/api/product-warehouse-stocks/adjust', {
          method: 'POST',
          body: {
            product_id: Number(productId),
            warehouse_id: Number(defaultWarehouseId.value),
            delta: Number(store.inventoryForm.stock_quantity),
            reason: 'initial_backfill_default_warehouse'
          }
        })
        // بارگذاری مجدد موجودی انبارها بعد از تنظیم
        await loadStocks(productId)
        didEnsureDefaultRow.value = true
      } catch (adjustError) {
        console.error('خطا در تنظیم موجودی انبار پیشفرض:', adjustError)
        didEnsureDefaultRow.value = true // جلوگیری از تکرار تلاش
      }
    }
  } catch (error) {
    console.error('خطا در ensureDefaultRow:', error)
  }
}

watch(() => store.editingProductId, async (pid) => {
  if (!pid) return
  try {
    // ابتدا اطلاعات موجودی کلی را بارگذاری کن
    await loadInventoryData(pid)
    // سپس انبار پیش‌فرض را تنظیم کن
    await loadDefaultWarehouse(pid)
    // در نهایت موجودی انبارها را بارگذاری کن
    await loadStocks(pid)
    await ensureDefaultRow(pid)
  } catch (error) {
    console.error('خطا در بارگذاری داده‌های موجودی:', error)
  }
}, { immediate: true })

// دیگر نیازی به توابع تغییر مجموع نیست؛ موجودی فقط از ردیف‌های انبار می‌آید.
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
  background-color: white;
  color: #1f2937 ;
  padding: 8px 12px;
  font-weight: normal;
}

select option:checked,
select option:hover {
  background-color: #f3f4f6 ;
  color: #1f2937 ;
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
