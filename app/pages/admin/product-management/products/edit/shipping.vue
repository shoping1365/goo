
   <template v-else-if="activeTab === 'shipping'">
      <!-- ابعاد و وزن محصول -->
      <div class="bg-gradient-to-r from-blue-50 to-cyan-50 border border-blue-200 rounded-xl shadow-lg p-6 text-right mb-6 transition-all duration-300 hover:shadow-xl">
        <div class="flex items-center justify-between cursor-pointer" @click="toggleSection('dimensions')">
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 bg-gradient-to-r from-blue-600 to-cyan-600 rounded-lg flex items-center justify-center shadow-md">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 7h6m0 10v-3m-3 3h.01M9 17h.01M9 14h.01M12 14h.01M15 11h.01M12 11h.01M9 11h.01M7 21h10a2 2 0 002-2V5a2 2 0 00-2-2H7a2 2 0 00-2 2v14a2 2 0 002 2z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-bold text-blue-800">ابعاد و وزن محصول</h3>
          </div>
          <span class="text-blue-600 text-xl font-bold bg-white rounded-full w-8 h-8 flex items-center justify-center shadow-md">{{ sections.dimensions ? '−' : '+' }}</span>
        </div>

        <div v-show="sections.dimensions" class="mt-6">
          <div class="grid grid-cols-1 lg:grid-cols-4 gap-6">
            <!-- وزن -->
            <div class="bg-white rounded-xl p-5 shadow-md border border-purple-100 transition-all duration-300 hover:shadow-lg">
              <div class="flex items-center gap-2 mb-3">
                <svg class="w-5 h-5 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6l3 1m0 0l-3 9a5.002 5.002 0 006.001 0M6 7l3 9M6 7l6-2m6 2l3-1m-3 1l-3 9a5.002 5.002 0 006.001 0M18 7l3 9m-3-9l-6-2m0-2v2m0 16V5m0 16l-3-3m3 3l3-3"></path>
                </svg>
                <label class="text-sm font-bold text-purple-800">وزن</label>
              </div>
              <div class="flex items-center gap-3">
                <input type="number" class="flex-1 border-2 border-purple-200 rounded-lg px-4 py-3 text-gray-800 focus:border-purple-500 focus:ring-2 focus:ring-purple-200 transition-all duration-300 bg-white" min="0" step="0.1" placeholder="0" />
                <span class="text-sm font-medium text-purple-600 bg-purple-100 px-3 py-2 rounded-lg">کیلوگرم</span>
              </div>
            </div>

            <!-- طول -->
            <div class="bg-white rounded-xl p-5 shadow-md border border-green-100 transition-all duration-300 hover:shadow-lg">
              <div class="flex items-center gap-2 mb-3">
                <svg class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21l3-3m-3 3l-3-3m3 3V8a4 4 0 014-4h4a4 4 0 014 4v13"></path>
                </svg>
                <label class="text-sm font-bold text-green-800">طول</label>
              </div>
              <div class="flex items-center gap-3">
                <input type="number" class="flex-1 border-2 border-green-200 rounded-lg px-4 py-3 text-gray-800 focus:border-green-500 focus:ring-2 focus:ring-green-200 transition-all duration-300 bg-white" min="0" placeholder="0" />
                <span class="text-sm font-medium text-green-600 bg-green-100 px-3 py-2 rounded-lg">سانتی‌متر</span>
              </div>
            </div>

            <!-- عرض -->
            <div class="bg-white rounded-xl p-5 shadow-md border border-orange-100 transition-all duration-300 hover:shadow-lg">
              <div class="flex items-center gap-2 mb-3">
                <svg class="w-5 h-5 text-orange-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 3.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"></path>
                </svg>
                <label class="text-sm font-bold text-orange-800">عرض</label>
              </div>
              <div class="flex items-center gap-3">
                <input type="number" class="flex-1 border-2 border-orange-200 rounded-lg px-4 py-3 text-gray-800 focus:border-orange-500 focus:ring-2 focus:ring-orange-200 transition-all duration-300 bg-white" min="0" placeholder="0" />
                <span class="text-sm font-medium text-orange-600 bg-orange-100 px-3 py-2 rounded-lg">سانتی‌متر</span>
              </div>
            </div>

            <!-- ارتفاع -->
            <div class="bg-white rounded-xl p-5 shadow-md border border-blue-100 transition-all duration-300 hover:shadow-lg">
              <div class="flex items-center gap-2 mb-3">
                <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 21v-6m0 0l2.5-2.5m0 0L8 15m-2.5-2.5L8 10m-2.5 2.5L3 15m18 6v-6m0 0l-2.5-2.5m0 0L16 15m2.5-2.5L16 10m2.5 2.5L21 15"></path>
                </svg>
                <label class="text-sm font-bold text-blue-800">ارتفاع</label>
              </div>
              <div class="flex items-center gap-3">
                <input type="number" class="flex-1 border-2 border-blue-200 rounded-lg px-4 py-3 text-gray-800 focus:border-blue-500 focus:ring-2 focus:ring-blue-200 transition-all duration-300 bg-white" min="0" placeholder="0" />
                <span class="text-sm font-medium text-blue-600 bg-blue-100 px-3 py-2 rounded-lg">سانتی‌متر</span>
              </div>
            </div>
          </div>

          <!-- اطلاعات اضافی ابعاد -->
          <div class="mt-4 pt-4 border-t">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <!-- حجم محاسبه شده -->
              <div class="flex flex-col gap-2">
                <label class="text-xs text-gray-700 font-semibold">حجم محاسبه شده</label>
                <div class="bg-gray-100 p-2 rounded text-xs text-gray-700">
                  <span class="font-medium">0 سانتی‌متر مکعب</span>
                </div>
              </div>

              <!-- دسته‌بندی ابعادی -->
              <div class="flex flex-col gap-2">
                <label class="text-xs text-gray-700 font-semibold">دسته‌بندی ابعادی</label>
                <select class="w-full text-sm px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all duration-200 bg-white text-gray-900">
                  <option value="" class="text-gray-900 bg-white py-2">خودکار</option>
                  <option value="small" class="text-gray-900 bg-white py-2">کوچک (زیر 30 سانتی‌متر)</option>
                  <option value="medium" class="text-gray-900 bg-white py-2">متوسط (30-100 سانتی‌متر)</option>
                  <option value="large" class="text-gray-900 bg-white py-2">بزرگ (بالای 100 سانتی‌متر)</option>
                  <option value="oversized" class="text-gray-900 bg-white py-2">فوق‌العاده بزرگ</option>
                </select>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- تنظیمات ارسال اصلی -->
      <div class="bg-gradient-to-r from-green-50 to-emerald-50 border border-green-200 rounded-xl shadow-lg p-6 text-right mb-6 transition-all duration-300 hover:shadow-xl">
        <div class="flex items-center justify-between cursor-pointer" @click="toggleSection('shippingSettings')">
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 bg-gradient-to-r from-green-600 to-emerald-600 rounded-lg flex items-center justify-center shadow-md">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-bold text-green-800">تنظیمات ارسال اصلی</h3>
          </div>
          <span class="text-green-600 text-xl font-bold bg-white rounded-full w-8 h-8 flex items-center justify-center shadow-md">{{ sections.shippingSettings ? '−' : '+' }}</span>
        </div>

        <div v-show="sections.shippingSettings" class="mt-4">
          <div class="grid grid-cols-1 gap-6">
            <!-- گزینه‌های اصلی ارسال -->
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <label class="flex items-center gap-2">
                <input id="isFreeShipping" type="checkbox" class="checkbox" />
                <span class="text-xs text-gray-700 font-semibold">ارسال رایگان</span>
              </label>
              <label class="flex items-center gap-2">
                <input id="isSeparateShipping" type="checkbox" class="checkbox" />
                <span class="text-xs text-gray-700 font-semibold">ارسال به صورت جداگانه</span>
              </label>
              <label class="flex items-center gap-2">
                <input type="checkbox" class="checkbox" />
                <span class="text-xs text-gray-700 font-semibold">نیاز به امضا هنگام تحویل</span>
              </label>
              <label class="flex items-center gap-2">
                <input type="checkbox" class="checkbox" />
                <span class="text-xs text-gray-700 font-semibold">قابل ارسال به تمام شهرها</span>
              </label>
            </div>

            <!-- روش‌های ارسال مجاز -->
            <div class="border-t pt-4">
              <h4 class="text-xs font-semibold text-gray-700 mb-3">روش‌های ارسال مجاز</h4>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-2">
                <label class="flex items-center gap-2">
                  <input type="checkbox" class="checkbox" checked />
                  <span class="text-xs text-gray-700">پست پیشتاز</span>
                </label>
                <label class="flex items-center gap-2">
                  <input type="checkbox" class="checkbox" checked />
                  <span class="text-xs text-gray-700">پیک موتوری</span>
                </label>
                <label class="flex items-center gap-2">
                  <input type="checkbox" class="checkbox" />
                  <span class="text-xs text-gray-700">باربری</span>
                </label>
                <label class="flex items-center gap-2">
                  <input type="checkbox" class="checkbox" />
                  <span class="text-xs text-gray-700">حمل ویژه</span>
                </label>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- هزینه‌های ارسال -->
      <div class="bg-gradient-to-r from-orange-50 to-amber-50 border border-orange-200 rounded-xl shadow-lg p-6 text-right mb-6 transition-all duration-300 hover:shadow-xl">
        <div class="flex items-center justify-between cursor-pointer" @click="toggleSection('shippingCosts')">
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 bg-gradient-to-r from-orange-600 to-amber-600 rounded-lg flex items-center justify-center shadow-md">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 9V7a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2m2 4h10a2 2 0 002-2v-6a2 2 0 00-2-2H9a2 2 0 00-2 2v6a2 2 0 002 2zm7-5a2 2 0 11-4 0 2 2 0 014 0z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-bold text-orange-800">هزینه‌های ارسال</h3>
          </div>
          <span class="text-orange-600 text-xl font-bold bg-white rounded-full w-8 h-8 flex items-center justify-center shadow-md">{{ sections.shippingCosts ? '−' : '+' }}</span>
        </div>

        <div v-show="sections.shippingCosts" class="mt-4">
          <div class="grid grid-cols-1 gap-6">
            <!-- تنظیمات هزینه -->
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <!-- مبلغ اضافی -->
              <div class="flex flex-col gap-2">
                <label class="text-xs text-gray-700 font-semibold">مبلغ اضافه به هزینه ارسال</label>
                <div class="flex items-center gap-2">
                  <input type="number" class="w-32 text-sm px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-orange-500 transition-all duration-200 bg-white text-gray-900" min="0" placeholder="0" />
                  <span class="text-xs text-gray-500">تومان</span>
                </div>
              </div>

              <!-- نوع محاسبه هزینه -->
              <div class="flex flex-col gap-2">
                <label class="text-xs text-gray-700 font-semibold">نوع محاسبه هزینه</label>
                <select class="w-full text-sm px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-orange-500 transition-all duration-200 bg-white text-gray-900">
                  <option value="fixed" class="text-gray-900 bg-white py-2">مبلغ ثابت</option>
                  <option value="weight" class="text-gray-900 bg-white py-2">بر اساس وزن</option>
                  <option value="volume" class="text-gray-900 bg-white py-2">بر اساس حجم</option>
                  <option value="distance" class="text-gray-900 bg-white py-2">بر اساس مسافت</option>
                </select>
              </div>
            </div>

            <!-- جدول هزینه‌های ارسال -->
            <div class="border-t pt-4">
              <h4 class="text-xs font-semibold text-gray-700 mb-3">هزینه‌های ارسال بر اساس منطقه</h4>
              <div class="overflow-x-auto">
                <table class="min-w-full text-xs text-right rtl border-collapse">
                  <thead>
                    <tr class="bg-gray-100 border-b">
                      <th class="px-3 py-2 text-gray-600 font-normal text-right">منطقه</th>
                      <th class="px-3 py-2 text-gray-600 font-normal text-right">روش ارسال</th>
                      <th class="px-3 py-2 text-gray-600 font-normal text-right">هزینه (تومان)</th>
                      <th class="px-3 py-2 text-gray-600 font-normal text-right">حداقل خرید رایگان</th>
                      <th class="px-3 py-2 text-gray-600 font-normal text-right w-20">ویرایش</th>
                      <th class="px-3 py-2 text-gray-600 font-normal text-right w-16">حذف</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr class="border-b hover:bg-gray-50">
                      <td class="px-3 py-2 text-right">تهران</td>
                      <td class="px-3 py-2 text-right">پیک موتوری</td>
                      <td class="px-3 py-2 text-right">25,000</td>
                      <td class="px-3 py-2 text-right">500,000</td>
                      <td class="px-3 py-2 text-center">
                        <button class="text-blue-500 hover:text-blue-700">ویرایش</button>
                      </td>
                      <td class="px-3 py-2 text-center">
                        <button class="text-red-500 hover:text-red-700">حذف</button>
                      </td>
                    </tr>
                    <tr class="border-b hover:bg-gray-50">
                      <td class="px-3 py-2 text-right">سایر شهرها</td>
                      <td class="px-3 py-2 text-right">پست پیشتاز</td>
                      <td class="px-3 py-2 text-right">35,000</td>
                      <td class="px-3 py-2 text-right">1,000,000</td>
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
              <div class="mt-3">
                <button class="bg-blue-600 text-white rounded px-4 py-2 text-xs hover:bg-blue-700 transition">افزودن منطقه جدید</button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- زمان‌بندی و تحویل -->
      <div class="bg-gradient-to-r from-purple-50 to-pink-50 border border-purple-200 rounded-xl shadow-lg p-6 text-right mb-6 transition-all duration-300 hover:shadow-xl">
        <div class="flex items-center justify-between cursor-pointer" @click="toggleSection('deliverySchedule')">
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 bg-gradient-to-r from-purple-600 to-pink-600 rounded-lg flex items-center justify-center shadow-md">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-bold text-purple-800">زمان‌بندی و تحویل</h3>
          </div>
          <span class="text-purple-600 text-xl font-bold bg-white rounded-full w-8 h-8 flex items-center justify-center shadow-md">{{ sections.deliverySchedule ? '−' : '+' }}</span>
        </div>

        <div v-show="sections.deliverySchedule" class="mt-4">
          <div class="grid grid-cols-1 gap-6">
            <!-- تنظیمات زمانی -->
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <!-- مدت تحویل -->
              <div class="flex flex-col gap-2">
                <label class="text-xs text-gray-700 font-semibold">مدت تخمینی تحویل</label>
                <select class="w-full text-sm px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500 transition-all duration-200 bg-white text-gray-900">
                  <option value="" class="text-gray-900 bg-white py-2">هیچکدام</option>
                  <option value="same_day" class="text-gray-900 bg-white py-2">همان روز</option>
                  <option value="next_day" class="text-gray-900 bg-white py-2">روز بعد</option>
                  <option value="2_days" class="text-gray-900 bg-white py-2">2 روز کاری</option>
                  <option value="3_days" class="text-gray-900 bg-white py-2">3 روز کاری</option>
                  <option value="1_week" class="text-gray-900 bg-white py-2">1 هفته</option>
                  <option value="2_weeks" class="text-gray-900 bg-white py-2">2 هفته</option>
                  <option value="custom" class="text-gray-900 bg-white py-2">سفارشی</option>
                </select>
              </div>

              <!-- مدت آماده‌سازی -->
              <div class="flex flex-col gap-2">
                <label class="text-xs text-gray-700 font-semibold">مدت آماده‌سازی</label>
                <select class="w-full text-sm px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500 transition-all duration-200 bg-white text-gray-900">
                  <option value="0" class="text-gray-900 bg-white py-2">آماده ارسال</option>
                  <option value="1" class="text-gray-900 bg-white py-2">1 روز کاری</option>
                  <option value="2" class="text-gray-900 bg-white py-2">2 روز کاری</option>
                  <option value="3" class="text-gray-900 bg-white py-2">3 روز کاری</option>
                  <option value="7" class="text-gray-900 bg-white py-2">1 هفته</option>
                </select>
              </div>
            </div>

            <!-- روزهای ارسال -->
            <div class="border-t pt-4">
              <h4 class="text-xs font-semibold text-gray-700 mb-3">روزهای مجاز ارسال</h4>
              <div class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-7 gap-2">
                <label class="flex items-center gap-2">
                  <input type="checkbox" class="checkbox" checked />
                  <span class="text-xs text-gray-700">شنبه</span>
                </label>
                <label class="flex items-center gap-2">
                  <input type="checkbox" class="checkbox" checked />
                  <span class="text-xs text-gray-700">یکشنبه</span>
                </label>
                <label class="flex items-center gap-2">
                  <input type="checkbox" class="checkbox" checked />
                  <span class="text-xs text-gray-700">دوشنبه</span>
                </label>
                <label class="flex items-center gap-2">
                  <input type="checkbox" class="checkbox" checked />
                  <span class="text-xs text-gray-700">سه‌شنبه</span>
                </label>
                <label class="flex items-center gap-2">
                  <input type="checkbox" class="checkbox" checked />
                  <span class="text-xs text-gray-700">چهارشنبه</span>
                </label>
                <label class="flex items-center gap-2">
                  <input type="checkbox" class="checkbox" />
                  <span class="text-xs text-gray-700">پنج‌شنبه</span>
                </label>
                <label class="flex items-center gap-2">
                  <input type="checkbox" class="checkbox" />
                  <span class="text-xs text-gray-700">جمعه</span>
                </label>
              </div>
            </div>

            <!-- ساعات ارسال -->
            <div class="border-t pt-4">
              <h4 class="text-xs font-semibold text-gray-700 mb-3">ساعات مجاز ارسال</h4>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div class="flex flex-col gap-2">
                  <label class="text-xs text-gray-700 font-semibold">از ساعت</label>
                  <input type="time" class="w-full text-sm px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500 transition-all duration-200 bg-white text-gray-900" value="08:00" />
                </div>
                <div class="flex flex-col gap-2">
                  <label class="text-xs text-gray-700 font-semibold">تا ساعت</label>
                  <input type="time" class="w-full text-sm px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500 transition-all duration-200 bg-white text-gray-900" value="18:00" />
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- تنظیمات ویژه و محدودیت‌ها -->
      <div class="bg-gradient-to-r from-red-50 to-rose-50 border border-red-200 rounded-xl shadow-lg p-6 text-right mb-6 transition-all duration-300 hover:shadow-xl">
        <div class="flex items-center justify-between cursor-pointer" @click="toggleSection('specialSettings')">
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 bg-gradient-to-r from-red-600 to-rose-600 rounded-lg flex items-center justify-center shadow-md">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-bold text-red-800">تنظیمات ویژه و محدودیت‌ها</h3>
          </div>
          <span class="text-red-600 text-xl font-bold bg-white rounded-full w-8 h-8 flex items-center justify-center shadow-md">{{ sections.specialSettings ? '−' : '+' }}</span>
        </div>

        <div v-show="sections.specialSettings" class="mt-4">
          <div class="grid grid-cols-1 gap-6">
            <!-- محدودیت‌های جغرافیایی -->
            <div>
              <h4 class="text-xs font-semibold text-gray-700 mb-3">محدودیت‌های جغرافیایی</h4>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-2">
                <label class="flex items-center gap-2">
                  <input type="checkbox" class="checkbox" />
                  <span class="text-xs text-gray-700">محدود به تهران</span>
                </label>
                <label class="flex items-center gap-2">
                  <input type="checkbox" class="checkbox" />
                  <span class="text-xs text-gray-700">محدود به کلان‌شهرها</span>
                </label>
                <label class="flex items-center gap-2">
                  <input type="checkbox" class="checkbox" />
                  <span class="text-xs text-gray-700">نیاز به مجوز ویژه</span>
                </label>
              </div>
            </div>

            <!-- شرایط خاص -->
            <div class="border-t pt-4">
              <h4 class="text-xs font-semibold text-gray-700 mb-3">شرایط خاص ارسال</h4>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-2">
                <label class="flex items-center gap-2">
                  <input type="checkbox" class="checkbox" />
                  <span class="text-xs text-gray-700">نیاز به نگهداری در سردخانه</span>
                </label>
                <label class="flex items-center gap-2">
                  <input type="checkbox" class="checkbox" />
                  <span class="text-xs text-gray-700">شکستنی - احتیاط در حمل</span>
                </label>
                <label class="flex items-center gap-2">
                  <input type="checkbox" class="checkbox" />
                  <span class="text-xs text-gray-700">مواد خطرناک</span>
                </label>
                <label class="flex items-center gap-2">
                  <input type="checkbox" class="checkbox" />
                  <span class="text-xs text-gray-700">ارسال با بیمه</span>
                </label>
              </div>
            </div>

            <!-- لیست شهرهای غیرمجاز -->
            <div class="border-t pt-4">
              <label class="text-xs text-gray-700 font-semibold">شهرهای غیرمجاز برای ارسال (اختیاری)</label>
              <textarea class="w-full text-sm px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-red-500 transition-all duration-200 bg-white text-gray-900 mt-2 min-h-[80px]" placeholder="نام شهرها را با کاما جدا کنید..."></textarea>
            </div>
          </div>
        </div>
      </div>

      <!-- بسته‌بندی و برچسب‌گذاری -->
      <div class="bg-gradient-to-r from-teal-50 to-cyan-50 border border-teal-200 rounded-xl shadow-lg p-6 text-right transition-all duration-300 hover:shadow-xl">
        <div class="flex items-center justify-between cursor-pointer" @click="toggleSection('packaging')">
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 bg-gradient-to-r from-teal-600 to-cyan-600 rounded-lg flex items-center justify-center shadow-md">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
              </svg>
            </div>
            <h3 class="text-lg font-bold text-teal-800">بسته‌بندی و برچسب‌گذاری</h3>
          </div>
          <span class="text-teal-600 text-xl font-bold bg-white rounded-full w-8 h-8 flex items-center justify-center shadow-md">{{ sections.packaging ? '−' : '+' }}</span>
        </div>

        <div v-show="sections.packaging" class="mt-4">
          <div class="grid grid-cols-1 gap-6">
            <!-- نوع بسته‌بندی -->
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div class="flex flex-col gap-2">
                <label class="text-xs text-gray-700 font-semibold">نوع بسته‌بندی</label>
                <select class="w-full text-sm px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-teal-500 focus:border-teal-500 transition-all duration-200 bg-white text-gray-900">
                  <option value="standard" class="text-gray-900 bg-white py-2">استاندارد</option>
                  <option value="bubble" class="text-gray-900 bg-white py-2">حبابی</option>
                  <option value="wooden" class="text-gray-900 bg-white py-2">جعبه چوبی</option>
                  <option value="custom" class="text-gray-900 bg-white py-2">سفارشی</option>
                </select>
              </div>

              <div class="flex flex-col gap-2">
                <label class="text-xs text-gray-700 font-semibold">هزینه بسته‌بندی</label>
                <div class="flex items-center gap-2">
                  <input type="number" class="w-32 text-sm px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-teal-500 focus:border-teal-500 transition-all duration-200 bg-white text-gray-900" min="0" placeholder="0" />
                  <span class="text-xs text-gray-500">تومان</span>
                </div>
              </div>
              </div>

            <!-- برچسب‌های ویژه -->
            <div class="border-t pt-4">
              <h4 class="text-xs font-semibold text-gray-700 mb-3">برچسب‌های ویژه</h4>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-2">
                <label class="flex items-center gap-2">
                  <input type="checkbox" class="checkbox" />
                  <span class="text-xs text-gray-700">این رو بالا (TOP)</span>
                </label>
                <label class="flex items-center gap-2">
                  <input type="checkbox" class="checkbox" />
                  <span class="text-xs text-gray-700">شکستنی (FRAGILE)</span>
                </label>
                <label class="flex items-center gap-2">
                  <input type="checkbox" class="checkbox" />
                  <span class="text-xs text-gray-700">خشک نگهداری شود</span>
                </label>
                <label class="flex items-center gap-2">
                  <input type="checkbox" class="checkbox" />
                  <span class="text-xs text-gray-700">حمل عمودی</span>
                </label>
              </div>
                </div>

            <!-- دستورالعمل‌های خاص -->
            <div class="border-t pt-4">
              <label class="text-xs text-gray-700 font-semibold">دستورالعمل‌های خاص برای حمل</label>
              <textarea class="w-full text-sm px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-teal-500 focus:border-teal-500 transition-all duration-200 bg-white text-gray-900 mt-2 min-h-[80px]" placeholder="دستورالعمل‌های خاص برای بسته‌بندی یا حمل این محصول..."></textarea>
                </div>

            <!-- خلاصه هزینه‌های ارسال -->
            <div class="border-t pt-4">
              <div class="bg-gray-50 p-3 rounded border border-blue-200">
                <h4 class="text-xs font-semibold text-gray-700 mb-2">خلاصه هزینه‌های ارسال</h4>
                <div class="grid grid-cols-2 md:grid-cols-4 gap-2 text-xs">
                  <div>
                    <span class="text-gray-600">هزینه پایه:</span>
                    <span class="font-medium text-blue-600"> 35,000 تومان</span>
                  </div>
                  <div>
                    <span class="text-gray-600">بسته‌بندی:</span>
                    <span class="font-medium text-blue-600"> 5,000 تومان</span>
                  </div>
                  <div>
                    <span class="text-gray-600">بیمه:</span>
                    <span class="font-medium text-blue-600"> 2,000 تومان</span>
                  </div>
                  <div>
                    <span class="text-gray-600">مجموع:</span>
                    <span class="font-bold text-green-600"> 42,000 تومان</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
</template>

<script lang="ts">
// definePageMeta در <script setup> استفاده می‌شود
// eslint-disable-next-line @typescript-eslint/no-unused-vars
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
</script>

<script setup lang="ts">
import { reactive } from 'vue';

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

const sections = reactive<Record<string, boolean>>({
  dimensions: true,
  shippingSettings: true,
  shippingCosts: true,
  deliverySchedule: true,
  specialSettings: true,
  packaging: true
})

// Toggle section visibility
const toggleSection = (section) => {
  sections[section] = !sections[section]
}
</script>
<script setup lang="ts">
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

.hover\:shadow-xl:hover {
  box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04);
}
</style>
