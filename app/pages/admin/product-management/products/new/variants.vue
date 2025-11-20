<template>
  <!-- ویژگی‌های تنوع -->
  <div class="bg-gradient-to-br from-emerald-50 to-teal-50 rounded-2xl border border-emerald-200 shadow-lg overflow-hidden mb-8">
    <div class="bg-gradient-to-r from-emerald-600 to-teal-600 px-6 py-4 cursor-pointer" @click="toggleSection('variantAttributes')">
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-3">
          <div class="p-2 bg-white/20 rounded-xl">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4M7.835 4.697a3.42 3.42 0 001.946-.806 3.42 3.42 0 014.438 0 3.42 3.42 0 001.946.806 3.42 3.42 0 013.138 3.138 3.42 3.42 0 00.806 1.946 3.42 3.42 0 010 4.438 3.42 3.42 0 00-.806 1.946 3.42 3.42 0 01-3.138 3.138 3.42 3.42 0 00-1.946.806 3.42 3.42 0 01-4.438 0 3.42 3.42 0 00-1.946-.806 3.42 3.42 0 01-3.138-3.138 3.42 3.42 0 00-.806-1.946 3.42 3.42 0 010-4.438 3.42 3.42 0 00.806-1.946 3.42 3.42 0 013.138-3.138z"></path>
            </svg>
          </div>
          <h3 class="text-xl font-bold text-white">ویژگی‌های تنوع محصول</h3>
        </div>
        <div class="p-2 bg-white/20 rounded-lg">
          <span class="text-white font-bold text-lg">{{ sections.variantAttributes ? '−' : '+' }}</span>
        </div>
      </div>
    </div>

    <div v-show="sections.variantAttributes" class="p-8">
      <div class="grid grid-cols-1 gap-6">
        <!-- افزودن ویژگی جدید -->
        <div class="bg-white rounded-xl shadow-sm border border-emerald-100 p-6">
          <h4 class="text-lg font-semibold text-gray-800 mb-6">افزودن ویژگی تنوع جدید</h4>
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div class="space-y-3">
              <label class="flex items-center gap-2 text-sm font-semibold text-gray-900">
                <div class="p-1 bg-emerald-100 rounded">
                  <svg class="w-3 h-3 text-emerald-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"></path>
                  </svg>
                </div>
                نام ویژگی
              </label>
              <input v-model="variantName" type="text" class="w-full px-4 py-3 bg-white border border-gray-300 rounded-lg focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 text-gray-900 transition-all duration-200" placeholder="مثال: رنگ" />
            </div>
            
            <div class="space-y-3">
              <label class="flex items-center gap-2 text-sm font-semibold text-gray-900">
                <div class="p-1 bg-teal-100 rounded">
                  <svg class="w-3 h-3 text-teal-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16"></path>
                  </svg>
                </div>
                نوع ویژگی
              </label>
              <select class="w-full px-4 py-3 bg-white border border-gray-300 rounded-lg focus:ring-2 focus:ring-teal-500 focus:border-teal-500 text-gray-900 transition-all duration-200">
                <option value="text">متنی</option>
                <option value="color">رنگ</option>
                <option value="image">تصویری</option>
                <option value="dropdown">لیست کشویی</option>
              </select>
            </div>
            
            <div class="flex items-end">
              <button class="w-full bg-gradient-to-r from-emerald-600 to-teal-600 text-white rounded-lg px-6 py-3 font-semibold hover:from-emerald-700 hover:to-teal-700 transition-all duration-200 shadow-md hover:shadow-lg" @click.prevent="addVariant">
                <svg class="w-4 h-4 inline-block mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
                </svg>
                افزودن ویژگی
              </button>
            </div>
          </div>
        </div>

        <!-- لیست ویژگی‌ها -->
        <div class="overflow-x-auto">
          <table class="min-w-full text-xs text-right rtl border-collapse">
            <thead>
            <tr class="bg-gray-100 border-b border-blue-200">
              <th class="px-3 py-2 text-gray-600 font-normal text-right">نام ویژگی</th>
              <th class="px-3 py-2 text-gray-600 font-normal text-right">نوع</th>
              <th class="px-3 py-2 text-gray-600 font-normal text-right">مقادیر</th>
              <th class="px-3 py-2 text-gray-600 font-normal text-right">تاثیر بر قیمت</th>
              <th class="px-3 py-2 text-gray-600 font-normal text-right w-20">ویرایش</th>
              <th class="px-3 py-2 text-gray-600 font-normal text-right w-16">حذف</th>
            </tr>
            </thead>
            <tbody>
              <tr v-for="v in variants" :key="v.id" class="border-b hover:bg-gray-50 border-blue-200">
                <td class="px-3 py-2 text-right">—</td>
                <td class="px-3 py-2 text-right">{{ v.id }}</td>
                <td class="px-3 py-2 text-right">{{ v.name }}</td>
                <td class="px-3 py-2 text-right">{{ v.value }}</td>
                <td class="px-3 py-2 text-right">{{ v.price_adjustment || 0 }}</td>
                <td class="px-3 py-2 text-right">{{ v.stock_quantity || 0 }}</td>
                <td class="px-3 py-2 text-right">
                  <span class="bg-green-100 text-green-800 px-2 py-1 rounded text-xs">فعال</span>
                </td>
                <td class="px-3 py-2 text-center">
                  <button class="text-blue-500 hover:text-blue-700">ویرایش</button>
                </td>
                <td class="px-3 py-2 text-center">
                  <button class="text-red-500 hover:text-red-700" @click="removeVariant(v.id)">حذف</button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <!-- تنظیمات ویژگی -->
        <div class="border-t pt-4 border-blue-200">
          <h4 class="text-xs font-semibold text-gray-700 mb-3">تنظیمات کلی ویژگی‌ها</h4>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-2">
            <label class="flex items-center gap-2">
              <input type="checkbox" class="checkbox" checked />
              <span class="text-xs text-gray-700">نمایش ویژگی‌ها در صفحه محصول</span>
            </label>
            <label class="flex items-center gap-2">
              <input type="checkbox" class="checkbox" />
              <span class="text-xs text-gray-700">اجباری بودن انتخاب تمام ویژگی‌ها</span>
            </label>
            <label class="flex items-center gap-2">
              <input type="checkbox" class="checkbox" />
              <span class="text-xs text-gray-700">نمایش قیمت تغییریافته برای هر تنوع</span>
            </label>
            <label class="flex items-center gap-2">
              <input type="checkbox" class="checkbox" checked />
              <span class="text-xs text-gray-700">نمایش تصویر مخصوص هر تنوع</span>
            </label>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- مدیریت تنوع‌ها -->
  <div class="bg-gradient-to-br from-blue-50 to-indigo-50 rounded-2xl border border-blue-200 shadow-lg overflow-hidden mb-8">
    <div class="bg-gradient-to-r from-blue-600 to-indigo-600 px-6 py-4 cursor-pointer" @click="toggleSection('variantManagement')">
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-3">
          <div class="p-2 bg-white/20 rounded-xl">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              
            </svg>
          </div>
          <h3 class="text-xl font-bold text-white">مدیریت تنوع‌های محصول</h3>
        </div>
        <div class="p-2 bg-white/20 rounded-lg">
          <span class="text-white font-bold text-lg">{{ sections.variantManagement ? '−' : '+' }}</span>
        </div>
      </div>
    </div>

    <div v-show="sections.variantManagement" class="mt-4">
      <div class="grid grid-cols-1 gap-6">
        <!-- ایجاد تنوع جدید -->
        <div class="border border-blue-200 rounded p-3 bg-gray-50">
          <h4 class="text-xs font-semibold text-gray-700 mb-3">ایجاد تنوع جدید</h4>
          <div class="grid grid-cols-1 md:grid-cols-5 gap-2">
            <div class="flex flex-col gap-1">
              <label class="text-xs text-gray-700 font-semibold">رنگ</label>
              <select class="input input-bordered w-full text-xs">
                <option>قرمز</option>
                <option>آبی</option>
                <option>سبز</option>
              </select>
            </div>
            <div class="flex flex-col gap-1">
              <label class="text-xs text-gray-700 font-semibold">سایز</label>
              <select class="input input-bordered w-full text-xs">
                <option>S</option>
                <option>M</option>
                <option>L</option>
                <option>XL</option>
              </select>
            </div>
            <div class="flex flex-col gap-1">
              <label class="text-xs text-gray-700 font-semibold">قیمت</label>
              <input type="number" class="input input-bordered w-full text-xs" min="0" placeholder="50000" />
            </div>
            <div class="flex flex-col gap-1">
              <label class="text-xs text-gray-700 font-semibold">موجودی</label>
              <input type="number" class="input input-bordered w-full text-xs" min="0" placeholder="10" />
            </div>
            <div class="flex items-end">
              <button class="bg-green-600 text-white rounded px-3 py-1 text-xs hover:bg-green-700 transition w-full">ایجاد</button>
            </div>
          </div>
        </div>

        <!-- جدول تنوع‌ها -->
        <div class="overflow-x-auto">
          <table class="min-w-full text-xs text-right rtl border-collapse">
            <thead>
            <tr class="bg-gray-100 border-b border-blue-200">
              <th class="px-3 py-2 text-gray-600 font-normal text-right">تصویر</th>
              <th class="px-3 py-2 text-gray-600 font-normal text-right">کد تنوع</th>
              <th class="px-3 py-2 text-gray-600 font-normal text-right">رنگ</th>
              <th class="px-3 py-2 text-gray-600 font-normal text-right">سایز</th>
              <th class="px-3 py-2 text-gray-600 font-normal text-right">قیمت</th>
              <th class="px-3 py-2 text-gray-600 font-normal text-right">موجودی</th>
              <th class="px-3 py-2 text-gray-600 font-normal text-right">وضعیت</th>
              <th class="px-3 py-2 text-gray-600 font-normal text-right w-20">ویرایش</th>
              <th class="px-3 py-2 text-gray-600 font-normal text-right w-16">حذف</th>
            </tr>
            </thead>
            <tbody>
            <tr class="border-b hover:bg-gray-50 border-blue-200">
              <td class="px-3 py-2 text-center">
                <div class="w-8 h-8 bg-red-500 rounded mx-auto"></div>
              </td>
              <td class="px-3 py-2 text-right">VAR-001</td>
              <td class="px-3 py-2 text-right">قرمز</td>
              <td class="px-3 py-2 text-right">M</td>
              <td class="px-3 py-2 text-right">50,000</td>
              <td class="px-3 py-2 text-right">15</td>
              <td class="px-3 py-2 text-right">
                <span class="bg-green-100 text-green-800 px-2 py-1 rounded text-xs">فعال</span>
              </td>
              <td class="px-3 py-2 text-center">
                <button class="text-blue-500 hover:text-blue-700">ویرایش</button>
              </td>
              <td class="px-3 py-2 text-center">
                <button class="text-red-500 hover:text-red-700">حذف</button>
              </td>
            </tr>
            <tr class="border-b hover:bg-gray-50 border-blue-200">
              <td class="px-3 py-2 text-center">
                <div class="w-8 h-8 bg-blue-500 rounded mx-auto"></div>
              </td>
              <td class="px-3 py-2 text-right">VAR-002</td>
              <td class="px-3 py-2 text-right">آبی</td>
              <td class="px-3 py-2 text-right">L</td>
              <td class="px-3 py-2 text-right">55,000</td>
              <td class="px-3 py-2 text-right">8</td>
              <td class="px-3 py-2 text-right">
                <span class="bg-green-100 text-green-800 px-2 py-1 rounded text-xs">فعال</span>
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

        <!-- عملیات کلی -->
        <div class="border-t pt-4 border-blue-200">
          <h4 class="text-xs font-semibold text-gray-700 mb-3">عملیات کلی</h4>
          <div class="flex flex-wrap gap-2">
            <button class="bg-blue-600 text-white rounded px-3 py-1 text-xs hover:bg-blue-700 transition">تولید خودکار تنوع‌ها</button>
            <button class="bg-green-600 text-white rounded px-3 py-1 text-xs hover:bg-green-700 transition">کپی قیمت از محصول اصلی</button>
            <button class="bg-orange-600 text-white rounded px-3 py-1 text-xs hover:bg-orange-700 transition">به‌روزرسانی موجودی</button>
            <button class="bg-purple-600 text-white rounded px-3 py-1 text-xs hover:bg-purple-700 transition">تنظیم تخفیف گروهی</button>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- قیمت‌گذاری تنوع‌ها -->
  <div class="bg-gradient-to-br from-orange-50 to-red-50 rounded-2xl border border-orange-200 shadow-lg overflow-hidden mb-8">
    <div class="bg-gradient-to-r from-orange-600 to-red-600 px-6 py-4 cursor-pointer" @click="toggleSection('variantPricing')">
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-3">
          <div class="p-2 bg-white/20 rounded-xl">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
            </svg>
          </div>
          <h3 class="text-xl font-bold text-white">قیمت‌گذاری تنوع‌ها</h3>
        </div>
        <div class="p-2 bg-white/20 rounded-lg">
          <span class="text-white font-bold text-lg">{{ sections.variantPricing ? '−' : '+' }}</span>
        </div>
      </div>
    </div>

    <div v-show="sections.variantPricing" class="mt-4">
      <div class="grid grid-cols-1 gap-6">
        <!-- قوانین قیمت‌گذاری -->
        <div class="border border-blue-200 rounded p-3 bg-gray-50">
          <h4 class="text-xs font-semibold text-gray-700 mb-3">قوانین قیمت‌گذاری</h4>
          <div class="grid grid-cols-1 md:grid-cols-4 gap-2">
            <div class="flex flex-col gap-1">
              <label class="text-xs text-gray-700 font-semibold">ویژگی</label>
              <select class="input input-bordered w-full text-xs">
                <option>رنگ</option>
                <option>سایز</option>
              </select>
            </div>
            <div class="flex flex-col gap-1">
              <label class="text-xs text-gray-700 font-semibold">مقدار</label>
              <select class="input input-bordered w-full text-xs">
                <option>XL</option>
                <option>قرمز</option>
              </select>
            </div>
            <div class="flex flex-col gap-1">
              <label class="text-xs text-gray-700 font-semibold">تغییر قیمت</label>
              <div class="flex gap-1">
                <select class="input input-bordered text-xs w-16">
                  <option>+</option>
                  <option>-</option>
                  <option>%+</option>
                  <option>%-</option>
                </select>
                <input type="number" class="input input-bordered text-xs flex-1" min="0" placeholder="5000" />
              </div>
            </div>
            <div class="flex items-end">
              <button class="bg-blue-600 text-white rounded px-3 py-1 text-xs hover:bg-blue-700 transition w-full">افزودن</button>
            </div>
          </div>
        </div>

        <!-- جدول قوانین قیمت -->
        <div class="overflow-x-auto">
          <table class="min-w-full text-xs text-right rtl border-collapse">
            <thead>
            <tr class="bg-gray-100 border-b border-blue-200">
              <th class="px-3 py-2 text-gray-600 font-normal text-right">ویژگی</th>
              <th class="px-3 py-2 text-gray-600 font-normal text-right">مقدار</th>
              <th class="px-3 py-2 text-gray-600 font-normal text-right">تغییر قیمت</th>
              <th class="px-3 py-2 text-gray-600 font-normal text-right">اولویت</th>
              <th class="px-3 py-2 text-gray-600 font-normal text-right w-20">ویرایش</th>
              <th class="px-3 py-2 text-gray-600 font-normal text-right w-16">حذف</th>
            </tr>
            </thead>
            <tbody>
            <tr class="border-b hover:bg-gray-50 border-blue-200">
              <td class="px-3 py-2 text-right">سایز</td>
              <td class="px-3 py-2 text-right">XL</td>
              <td class="px-3 py-2 text-right">+5,000 تومان</td>
              <td class="px-3 py-2 text-right">1</td>
              <td class="px-3 py-2 text-center">
                <button class="text-blue-500 hover:text-blue-700">ویرایش</button>
              </td>
              <td class="px-3 py-2 text-center">
                <button class="text-red-500 hover:text-red-700">حذف</button>
              </td>
            </tr>
            <tr class="border-b hover:bg-gray-50 border-blue-200">
              <td class="px-3 py-2 text-right">رنگ</td>
              <td class="px-3 py-2 text-right">طلایی</td>
              <td class="px-3 py-2 text-right">+15%</td>
              <td class="px-3 py-2 text-right">2</td>
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

        <!-- خلاصه قیمت‌گذاری -->
        <div class="border-t pt-4 border-blue-200">
          <div class="bg-gray-50 p-3 rounded border border-blue-200">
            <h4 class="text-xs font-semibold text-gray-700 mb-2">خلاصه قیمت‌گذاری تنوع‌ها</h4>
            <div class="grid grid-cols-2 md:grid-cols-4 gap-2 text-xs">
              <div>
                <span class="text-gray-600">کمترین قیمت:</span>
                <span class="font-medium text-green-600"> 50,000 تومان</span>
              </div>
              <div>
                <span class="text-gray-600">بیشترین قیمت:</span>
                <span class="font-medium text-red-600"> 65,000 تومان</span>
              </div>
              <div>
                <span class="text-gray-600">میانگین قیمت:</span>
                <span class="font-medium text-blue-600"> 57,500 تومان</span>
              </div>
              <div>
                <span class="text-gray-600">تعداد تنوع:</span>
                <span class="font-medium text-purple-600"> 12 تنوع</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- تصاویر تنوع‌ها -->
  <div class="bg-gradient-to-br from-purple-50 to-pink-50 rounded-2xl border border-purple-200 shadow-lg overflow-hidden">
    <div class="bg-gradient-to-r from-purple-600 to-pink-600 px-6 py-4 cursor-pointer" @click="toggleSection('variantImages')">
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-3">
          <div class="p-2 bg-white/20 rounded-xl">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
            </svg>
          </div>
          <h3 class="text-xl font-bold text-white">تصاویر تنوع‌های محصول</h3>
        </div>
        <div class="p-2 bg-white/20 rounded-lg">
          <span class="text-white font-bold text-lg">{{ sections.variantImages ? '−' : '+' }}</span>
        </div>
      </div>
    </div>

    <div v-show="sections.variantImages" class="mt-4">
      <div class="grid grid-cols-1 gap-6">
        <!-- آپلود تصویر برای تنوع -->
        <div class="border border-blue-200 rounded p-3 bg-gray-50">
          <h4 class="text-xs font-semibold text-gray-700 mb-3">تخصیص تصویر به تنوع</h4>
          <div class="grid grid-cols-1 md:grid-cols-3 gap-2">
            <div class="flex flex-col gap-1">
              <label class="text-xs text-gray-700 font-semibold">انتخاب تنوع</label>
              <select class="input input-bordered w-full text-xs">
                <option>قرمز - M</option>
                <option>آبی - L</option>
                <option>سبز - S</option>
              </select>
            </div>
            <div class="flex flex-col gap-1">
              <label class="text-xs text-gray-700 font-semibold">آپلود تصویر</label>
              <input type="file" class="input input-bordered w-full text-xs" accept="image/*" />
            </div>
            <div class="flex items-end">
              <button class="bg-blue-600 text-white rounded px-3 py-1 text-xs hover:bg-blue-700 transition w-full">تخصیص</button>
            </div>
          </div>
        </div>

        <!-- گالری تصاویر تنوع‌ها -->
        <div class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-6">
          <div class="border rounded p-2 text-center border-blue-200">
            <div class="w-full h-20 bg-red-100 rounded mb-2 flex items-center justify-center">
              <span class="text-xs text-gray-500">قرمز - M</span>
            </div>
            <div class="text-xs text-gray-700">قرمز - M</div>
            <button class="text-red-500 hover:text-red-700 text-xs mt-1">حذف</button>
          </div>
          <div class="border rounded p-2 text-center border-blue-200">
            <div class="w-full h-20 bg-blue-100 rounded mb-2 flex items-center justify-center">
              <span class="text-xs text-gray-500">آبی - L</span>
            </div>
            <div class="text-xs text-gray-700">آبی - L</div>
            <button class="text-red-500 hover:text-red-700 text-xs mt-1">حذف</button>
          </div>
          <div class="border rounded p-2 text-center border-blue-200">
            <div class="w-full h-20 bg-green-100 rounded mb-2 flex items-center justify-center">
              <span class="text-xs text-gray-500">سبز - S</span>
            </div>
            <div class="text-xs text-gray-700">سبز - S</div>
            <button class="text-red-500 hover:text-red-700 text-xs mt-1">حذف</button>
          </div>
          <div class="border-2 border-dashed border-blue-200 rounded p-2 text-center flex items-center justify-center">
            <div class="text-center">
              <div class="text-2xl text-gray-400">+</div>
              <div class="text-xs text-gray-500">افزودن تصویر</div>
            </div>
          </div>
        </div>

        <!-- تنظیمات تصاویر -->
        <div class="border-t pt-4 border-blue-200">
          <h4 class="text-xs font-semibold text-gray-700 mb-3">تنظیمات نمایش تصاویر</h4>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-2">
            <label class="flex items-center gap-2">
              <input type="checkbox" class="checkbox" checked />
              <span class="text-xs text-gray-700">نمایش تصویر تنوع در لیست محصولات</span>
            </label>
            <label class="flex items-center gap-2">
              <input type="checkbox" class="checkbox" checked />
              <span class="text-xs text-gray-700">تغییر تصویر اصلی با انتخاب تنوع</span>
            </label>
            <label class="flex items-center gap-2">
              <input type="checkbox" class="checkbox" />
              <span class="text-xs text-gray-700">نمایش همه تصاویر تنوع‌ها در گالری</span>
            </label>
            <label class="flex items-center gap-2">
              <input type="checkbox" class="checkbox" />
              <span class="text-xs text-gray-700">ضغط خودکار تصاویر</span>
            </label>
          </div>
        </div>

        <!-- خلاصه تصاویر -->
        <div class="border-t pt-4 border-blue-200">
          <div class="bg-gray-50 p-3 rounded border border-blue-200">
            <h4 class="text-xs font-semibold text-gray-700 mb-2">خلاصه تصاویر تنوع‌ها</h4>
            <div class="grid grid-cols-2 md:grid-cols-4 gap-2 text-xs">
              <div>
                <span class="text-gray-600">تصاویر آپلود شده:</span>
                <span class="font-medium text-blue-600"> 3 تصویر</span>
              </div>
              <div>
                <span class="text-gray-600">تنوع‌های بدون تصویر:</span>
                <span class="font-medium text-orange-600"> 9 تنوع</span>
              </div>
              <div>
                <span class="text-gray-600">حجم کل:</span>
                <span class="font-medium text-green-600"> 2.4 مگابایت</span>
              </div>
              <div>
                <span class="text-gray-600">فرمت پیشنهادی:</span>
                <span class="font-medium text-purple-600"> JPG, PNG</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useProductCreateStore } from '~/stores/productCreate'
import { useNotifier } from '~/composables/useNotifier'

const store = useProductCreateStore()
const notifier = useNotifier()
// استفاده از وضعیت بخش‌های قابل جمع/باز از استور سراسری
const sections = store.sections
const toggleSection = store.toggleSection
const variantName = ref('')
const variantValue = ref('')
const variantPrice = ref<number | null>(null)
const variantStock = ref<number | null>(null)
const variants = ref<any[]>([])

const productId = computed(()=> store.editingProductId)

async function loadVariants(){
  if (!store.isEditMode || !store.editingProductId) { variants.value = []; return }
  try{
    const res:any = await $fetch(`/api/product-variants/${store.editingProductId}`)
    variants.value = Array.isArray(res?.data) ? res.data : []
  }catch(e:any){
    // silent error to avoid double toasts
    console.error('loadVariants failed', e)
  }
}

async function addVariant(){
  if (!store.isEditMode || !store.editingProductId) return
  const body:any = { name: variantName.value, value: variantValue.value }
  if (variantPrice.value != null) body.price_adjustment = Number(variantPrice.value)
  if (variantStock.value != null) body.stock_quantity = Number(variantStock.value)
  try {
    await $fetch(`/api/product-variants/${store.editingProductId}`, { method:'POST', body })
    await loadVariants()
    variantName.value = ''
    variantValue.value = ''
    variantPrice.value = null
    variantStock.value = null
    notifier.success('تنوع با موفقیت اضافه شد')
  } catch(e){ /* ignore */ }
}

async function removeVariant(id:number){
  if (!id) return
  try{
    await $fetch(`/api/product-variants/${id}`, { method:'DELETE' })
    await loadVariants()
    notifier.success('تنوع حذف شد')
  }catch(e:any){ console.error('removeVariant failed', e) }
}

onMounted(()=>{ loadVariants() })
</script>
