<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between">
        <div>
          <h2 class="text-xl font-bold text-gray-900">مدیریت مالیات</h2>
          <p class="text-gray-600 mt-1">محاسبه، مدیریت و گزارش‌گیری مالیات کیف پول</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button
            @click="activeTab = 'calculation'"
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'calculation'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
          >
            محاسبه مالیات
          </button>
          <button
            @click="activeTab = 'rates'"
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'rates'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
          >
            نرخ‌های مالیاتی
          </button>
          <button
            @click="activeTab = 'exemptions'"
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'exemptions'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
          >
            معافیت‌ها
          </button>
          <button
            @click="activeTab = 'reports'"
            :class="[
              'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
              activeTab === 'reports'
                ? 'bg-blue-100 text-blue-700 border border-blue-200'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
          >
            گزارش‌های مالیاتی
          </button>
        </div>
      </div>
    </div>

    <!-- محتوای تب‌ها -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200">
      <!-- محاسبه مالیات -->
      <div v-if="activeTab === 'calculation'" class="p-6">
        <div class="space-y-6">
          <!-- آمار مالیاتی -->
          <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
            <div class="bg-white border border-gray-200 rounded-lg p-6">
              <div class="flex items-center">
                <div class="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center">
                  <svg class="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1" />
                  </svg>
                </div>
                <div class="mr-3">
                  <p class="text-sm font-medium text-gray-900">کل درآمد مشمول</p>
                  <p class="text-lg font-bold text-blue-600">{{ formatCurrency(taxStats.totalIncome) }}</p>
                </div>
              </div>
            </div>
            <div class="bg-white border border-gray-200 rounded-lg p-6">
              <div class="flex items-center">
                <div class="w-8 h-8 bg-red-100 rounded-full flex items-center justify-center">
                  <svg class="w-4 h-4 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 7h6m0 10v-3m-3 3h.01M9 17h.01M9 14h.01M12 14h.01M15 11h.01M12 11h.01M9 11h.01M7 21h10a2 2 0 002-2V5a2 2 0 00-2-2H7a2 2 0 00-2 2v14a2 2 0 002 2z" />
                  </svg>
                </div>
                <div class="mr-3">
                  <p class="text-sm font-medium text-gray-900">کل مالیات</p>
                  <p class="text-lg font-bold text-red-600">{{ formatCurrency(taxStats.totalTax) }}</p>
                </div>
              </div>
            </div>
            <div class="bg-white border border-gray-200 rounded-lg p-6">
              <div class="flex items-center">
                <div class="w-8 h-8 bg-green-100 rounded-full flex items-center justify-center">
                  <svg class="w-4 h-4 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                </div>
                <div class="mr-3">
                  <p class="text-sm font-medium text-gray-900">معافیت‌ها</p>
                  <p class="text-lg font-bold text-green-600">{{ formatCurrency(taxStats.exemptions) }}</p>
                </div>
              </div>
            </div>
            <div class="bg-white border border-gray-200 rounded-lg p-6">
              <div class="flex items-center">
                <div class="w-8 h-8 bg-purple-100 rounded-full flex items-center justify-center">
                  <svg class="w-4 h-4 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
                  </svg>
                </div>
                <div class="mr-3">
                  <p class="text-sm font-medium text-gray-900">نرخ متوسط</p>
                  <p class="text-lg font-bold text-purple-600">{{ taxStats.averageRate }}%</p>
                </div>
              </div>
            </div>
          </div>

          <!-- محاسبه‌گر مالیات -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">محاسبه‌گر مالیات</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <h4 class="font-medium text-gray-900 mb-3">اطلاعات ورودی</h4>
                <div class="space-y-3">
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">درآمد سالانه</label>
                    <input v-model="taxCalculator.annualIncome" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="درآمد سالانه به تومان">
                  </div>
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">نوع فعالیت</label>
                    <select v-model="taxCalculator.activityType" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                      <option value="">انتخاب کنید</option>
                      <option value="individual">شخص حقیقی</option>
                      <option value="corporate">شخص حقوقی</option>
                      <option value="freelancer">فریلنسر</option>
                    </select>
                  </div>
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">استان</label>
                    <select v-model="taxCalculator.province" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                      <option value="">انتخاب کنید</option>
                      <option value="tehran">تهران</option>
                      <option value="isfahan">اصفهان</option>
                      <option value="shiraz">شیراز</option>
                      <option value="mashhad">مشهد</option>
                    </select>
                  </div>
                  <div class="flex items-center">
                    <input type="checkbox" v-model="taxCalculator.hasExemptions" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                    <label class="mr-2 text-sm text-gray-700">دارای معافیت مالیاتی</label>
                  </div>
                  <button @click="calculateTax" class="w-full bg-blue-600 text-white py-2 px-4 rounded-lg hover:bg-blue-700 transition-colors">
                    محاسبه مالیات
                  </button>
                </div>
              </div>
              <div>
                <h4 class="font-medium text-gray-900 mb-3">نتایج محاسبه</h4>
                <div class="space-y-3">
                  <div class="bg-gray-50 rounded-lg p-6">
                    <div class="flex justify-between text-sm">
                      <span class="text-gray-600">درآمد مشمول:</span>
                      <span class="text-gray-900">{{ formatCurrency(taxCalculation.taxableIncome) }}</span>
                    </div>
                  </div>
                  <div class="bg-gray-50 rounded-lg p-6">
                    <div class="flex justify-between text-sm">
                      <span class="text-gray-600">معافیت‌ها:</span>
                      <span class="text-gray-900">{{ formatCurrency(taxCalculation.exemptions) }}</span>
                    </div>
                  </div>
                  <div class="bg-gray-50 rounded-lg p-6">
                    <div class="flex justify-between text-sm">
                      <span class="text-gray-600">مالیات محاسبه شده:</span>
                      <span class="text-red-600 font-medium">{{ formatCurrency(taxCalculation.calculatedTax) }}</span>
                    </div>
                  </div>
                  <div class="bg-gray-50 rounded-lg p-6">
                    <div class="flex justify-between text-sm">
                      <span class="text-gray-600">نرخ مالیات:</span>
                      <span class="text-gray-900">{{ taxCalculation.taxRate }}%</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- نرخ‌های مالیاتی -->
      <div v-if="activeTab === 'rates'" class="p-6">
        <div class="space-y-6">
          <!-- جدول نرخ‌های مالیاتی -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">نرخ‌های مالیاتی شخص حقیقی</h3>
            <div class="overflow-x-auto">
              <table class="min-w-full divide-y divide-gray-200">
                <thead class="bg-gray-50">
                  <tr>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">ردیف</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">درآمد مشمول</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نرخ مالیات</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مالیات اضافی</th>
                    <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
                  </tr>
                </thead>
                <tbody class="bg-white divide-y divide-gray-200">
                  <tr v-for="(rate, index) in taxRates" :key="rate.id">
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ index + 1 }}</td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ formatCurrency(rate.minIncome) }} - {{ formatCurrency(rate.maxIncome) }}</td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ rate.rate }}%</td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ formatCurrency(rate.additionalTax) }}</td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                      <button class="text-blue-600 hover:text-blue-800 ml-2">ویرایش</button>
                      <button class="text-red-600 hover:text-red-800">حذف</button>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>

          <!-- نرخ‌های مالیاتی حقوقی -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">نرخ‌های مالیاتی شخص حقوقی</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div v-for="corporateRate in corporateTaxRates" :key="corporateRate.id" class="border border-gray-200 rounded-lg p-6">
                <div class="flex items-center justify-between mb-3">
                  <h4 class="font-medium text-gray-900">{{ corporateRate.type }}</h4>
                  <span class="text-lg font-bold text-blue-600">{{ corporateRate.rate }}%</span>
                </div>
                <p class="text-sm text-gray-600 mb-3">{{ corporateRate.description }}</p>
                <div class="flex items-center justify-between">
                  <span class="text-xs text-gray-500">{{ corporateRate.updatedAt }}</span>
                  <button class="text-blue-600 hover:text-blue-800 text-sm">ویرایش</button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- معافیت‌ها -->
      <div v-if="activeTab === 'exemptions'" class="p-6">
        <div class="space-y-6">
          <!-- لیست معافیت‌ها -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">معافیت‌های مالیاتی</h3>
            <div class="space-y-4">
              <div v-for="exemption in taxExemptions" :key="exemption.id" class="border border-gray-200 rounded-lg p-6">
                <div class="flex items-center justify-between mb-3">
                  <div>
                    <h4 class="font-medium text-gray-900">{{ exemption.name }}</h4>
                    <p class="text-sm text-gray-600">{{ exemption.description }}</p>
                  </div>
                  <label class="relative inline-flex items-center cursor-pointer">
                    <input type="checkbox" v-model="exemption.active" class="sr-only peer">
                    <div class="w-9 h-5 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-4 after:w-4 after:transition-all peer-checked:bg-blue-600"></div>
                  </label>
                </div>
                <div class="grid grid-cols-2 md:grid-cols-4 gap-6 text-sm">
                  <div>
                    <span class="text-gray-600">مبلغ معافیت:</span>
                    <span class="text-gray-900">{{ formatCurrency(exemption.amount) }}</span>
                  </div>
                  <div>
                    <span class="text-gray-600">نوع:</span>
                    <span class="text-gray-900">{{ exemption.type }}</span>
                  </div>
                  <div>
                    <span class="text-gray-600">شرایط:</span>
                    <span class="text-gray-900">{{ exemption.conditions }}</span>
                  </div>
                  <div>
                    <span class="text-gray-600">آخرین به‌روزرسانی:</span>
                    <span class="text-gray-900">{{ exemption.updatedAt }}</span>
                  </div>
                </div>
                <div class="mt-3 flex space-x-2 space-x-reverse">
                  <button class="bg-blue-600 text-white px-3 py-2 rounded text-sm hover:bg-blue-700 transition-colors">
                    ویرایش
                  </button>
                  <button class="bg-gray-600 text-white px-3 py-2 rounded text-sm hover:bg-gray-700 transition-colors">
                    جزئیات
                  </button>
                </div>
              </div>
            </div>
          </div>

          <!-- افزودن معافیت جدید -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">افزودن معافیت جدید</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">نام معافیت</label>
                <input v-model="newExemption.name" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="نام معافیت">
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">مبلغ معافیت</label>
                <input v-model="newExemption.amount" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="مبلغ به تومان">
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">نوع معافیت</label>
                <select v-model="newExemption.type" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                  <option value="">انتخاب کنید</option>
                  <option value="fixed">ثابت</option>
                  <option value="percentage">درصدی</option>
                  <option value="conditional">مشروط</option>
                </select>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">شرایط</label>
                <input v-model="newExemption.conditions" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="شرایط معافیت">
              </div>
              <div class="md:col-span-2">
                <label class="block text-sm font-medium text-gray-700 mb-1">توضیحات</label>
                <textarea v-model="newExemption.description" rows="3" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="توضیحات معافیت"></textarea>
              </div>
              <div class="md:col-span-2">
                <button class="bg-green-600 text-white px-6 py-2 rounded-lg hover:bg-green-700 transition-colors">
                  افزودن معافیت
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- گزارش‌های مالیاتی -->
      <div v-if="activeTab === 'reports'" class="p-6">
        <div class="space-y-6">
          <!-- گزارش‌های آماده -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">گزارش‌های مالیاتی</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div v-for="report in taxReports" :key="report.id" class="border border-gray-200 rounded-lg p-6">
                <h4 class="font-medium text-gray-900 mb-2">{{ report.name }}</h4>
                <p class="text-sm text-gray-600 mb-3">{{ report.description }}</p>
                <div class="flex items-center justify-between">
                  <span class="text-xs text-gray-500">{{ report.lastGenerated }}</span>
                  <button class="bg-blue-600 text-white px-3 py-1 rounded text-sm hover:bg-blue-700 transition-colors">
                    ایجاد گزارش
                  </button>
                </div>
              </div>
            </div>
          </div>

          <!-- گزارش سفارشی -->
          <div class="border border-gray-200 rounded-lg p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">گزارش سفارشی مالیاتی</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <h4 class="font-medium text-gray-900 mb-3">پارامترهای گزارش</h4>
                <div class="space-y-3">
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">نوع گزارش</label>
                    <select v-model="customTaxReport.type" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                      <option value="">انتخاب کنید</option>
                      <option value="income">گزارش درآمد</option>
                      <option value="tax">گزارش مالیات</option>
                      <option value="exemption">گزارش معافیت‌ها</option>
                      <option value="comprehensive">گزارش جامع</option>
                    </select>
                  </div>
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">بازه زمانی</label>
                    <select v-model="customTaxReport.timeRange" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                      <option value="month">ماه جاری</option>
                      <option value="quarter">سه ماهه</option>
                      <option value="year">سال مالی</option>
                      <option value="custom">سفارشی</option>
                    </select>
                  </div>
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">فرمت گزارش</label>
                    <select v-model="customTaxReport.format" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                      <option value="pdf">PDF</option>
                      <option value="excel">Excel</option>
                      <option value="csv">CSV</option>
                    </select>
                  </div>
                </div>
              </div>
              <div>
                <h4 class="font-medium text-gray-900 mb-3">تنظیمات اضافی</h4>
                <div class="space-y-3">
                  <div class="flex items-center">
                    <input type="checkbox" v-model="customTaxReport.includeDetails" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                    <label class="mr-2 text-sm text-gray-700">شامل جزئیات کامل</label>
                  </div>
                  <div class="flex items-center">
                    <input type="checkbox" v-model="customTaxReport.includeCharts" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                    <label class="mr-2 text-sm text-gray-700">شامل نمودارها</label>
                  </div>
                  <div class="flex items-center">
                    <input type="checkbox" v-model="customTaxReport.sendEmail" class="rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                    <label class="mr-2 text-sm text-gray-700">ارسال به ایمیل</label>
                  </div>
                  <div v-if="customTaxReport.sendEmail">
                    <label class="block text-sm font-medium text-gray-700 mb-1">آدرس ایمیل</label>
                    <input v-model="customTaxReport.email" type="email" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500" placeholder="example@domain.com">
                  </div>
                </div>
              </div>
            </div>
            <div class="mt-6">
              <button class="bg-green-600 text-white px-6 py-2 rounded-lg hover:bg-green-700 transition-colors">
                ایجاد گزارش سفارشی
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// تعریف reactive data
const activeTab = ref('calculation')

// آمار مالیاتی
const taxStats = ref({
  totalIncome: 12500000000, // 12.5 میلیارد تومان
  totalTax: 1875000000, // 1.875 میلیارد تومان
  exemptions: 2500000000, // 2.5 میلیارد تومان
  averageRate: 15
})

// محاسبه‌گر مالیات
const taxCalculator = ref({
  annualIncome: 0,
  activityType: '',
  province: '',
  hasExemptions: false
})

// نتایج محاسبه
const taxCalculation = ref({
  taxableIncome: 0,
  exemptions: 0,
  calculatedTax: 0,
  taxRate: 0
})

// نرخ‌های مالیاتی شخص حقیقی
const taxRates = ref([
  {
    id: 1,
    minIncome: 0,
    maxIncome: 56000000,
    rate: 0,
    additionalTax: 0
  },
  {
    id: 2,
    minIncome: 56000000,
    maxIncome: 150000000,
    rate: 10,
    additionalTax: 0
  },
  {
    id: 3,
    minIncome: 150000000,
    maxIncome: 500000000,
    rate: 15,
    additionalTax: 9400000
  },
  {
    id: 4,
    minIncome: 500000000,
    maxIncome: 1000000000,
    rate: 20,
    additionalTax: 61900000
  },
  {
    id: 5,
    minIncome: 1000000000,
    maxIncome: 999999999999,
    rate: 25,
    additionalTax: 161900000
  }
])

// نرخ‌های مالیاتی شخص حقوقی
const corporateTaxRates = ref([
  {
    id: 1,
    type: 'شرکت‌های تولیدی',
    rate: 15,
    description: 'نرخ مالیات برای شرکت‌های تولیدی و صنعتی',
    updatedAt: '2024-01-15'
  },
  {
    id: 2,
    type: 'شرکت‌های خدماتی',
    rate: 25,
    description: 'نرخ مالیات برای شرکت‌های خدماتی و تجاری',
    updatedAt: '2024-01-15'
  },
  {
    id: 3,
    type: 'شرکت‌های دانش‌بنیان',
    rate: 10,
    description: 'نرخ مالیات برای شرکت‌های دانش‌بنیان',
    updatedAt: '2024-01-15'
  },
  {
    id: 4,
    type: 'سایر شرکت‌ها',
    rate: 22,
    description: 'نرخ مالیات برای سایر انواع شرکت‌ها',
    updatedAt: '2024-01-15'
  }
])

// معافیت‌های مالیاتی
const taxExemptions = ref([
  {
    id: 1,
    name: 'معافیت پایه',
    description: 'معافیت پایه سالانه برای تمام افراد',
    active: true,
    amount: 56000000,
    type: 'ثابت',
    conditions: 'برای تمام افراد',
    updatedAt: '2024-01-15'
  },
  {
    id: 2,
    name: 'معافیت بیمه',
    description: 'معافیت حق بیمه پرداختی',
    active: true,
    amount: 0,
    type: 'درصدی',
    conditions: '10% حق بیمه',
    updatedAt: '2024-01-15'
  },
  {
    id: 3,
    name: 'معافیت مسکن',
    description: 'معافیت اجاره مسکن',
    active: false,
    amount: 120000000,
    type: 'ثابت',
    conditions: 'برای مستأجرین',
    updatedAt: '2024-01-15'
  }
])

// معافیت جدید
const newExemption = ref({
  name: '',
  amount: 0,
  type: '',
  conditions: '',
  description: ''
})

// گزارش‌های مالیاتی
const taxReports = ref([
  {
    id: 1,
    name: 'گزارش مالیات سالانه',
    description: 'گزارش کامل مالیات سال مالی جاری',
    lastGenerated: '2024-01-31 14:00'
  },
  {
    id: 2,
    name: 'گزارش معافیت‌ها',
    description: 'گزارش معافیت‌های مالیاتی اعمال شده',
    lastGenerated: '2024-01-31 13:30'
  },
  {
    id: 3,
    name: 'گزارش نرخ‌های مالیاتی',
    description: 'گزارش نرخ‌های مالیاتی و تغییرات',
    lastGenerated: '2024-01-31 12:00'
  },
  {
    id: 4,
    name: 'گزارش تطبیق مالیاتی',
    description: 'گزارش تطبیق با قوانین مالیاتی',
    lastGenerated: '2024-01-31 11:00'
  }
])

// گزارش سفارشی
const customTaxReport = ref({
  type: '',
  timeRange: 'month',
  format: 'pdf',
  includeDetails: true,
  includeCharts: false,
  sendEmail: false,
  email: ''
})

// تابع محاسبه مالیات
const calculateTax = () => {
  const income = taxCalculator.value.annualIncome
  const exemptions = taxCalculator.value.hasExemptions ? 56000000 : 0
  const taxableIncome = Math.max(0, income - exemptions)
  
  let calculatedTax = 0
  let taxRate = 0
  
  // محاسبه مالیات بر اساس نرخ‌های پلکانی
  for (const rate of taxRates.value) {
    if (taxableIncome > rate.minIncome && taxableIncome <= rate.maxIncome) {
      const baseAmount = rate.minIncome
      const excessAmount = taxableIncome - baseAmount
      calculatedTax = rate.additionalTax + (excessAmount * rate.rate / 100)
      taxRate = rate.rate
      break
    }
  }
  
  taxCalculation.value = {
    taxableIncome,
    exemptions,
    calculatedTax,
    taxRate
  }
}

// تابع فرمت‌بندی ارز
const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR', {
    style: 'currency',
    currency: 'IRR',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount)
}
</script>

<!--
  مستندسازی:
  این کامپوننت شامل مدیریت مالیات کیف پول است که شامل:
  1. آمار مالیاتی: مالیات بر ارزش افزوده، مالیات بر درآمد، عوارض، کل مالیات
  2. تنظیمات مالیاتی: نرخ‌های مالیاتی و معافیت‌ها
  3. گزارش‌های مالیاتی: خرید و فروش، مالیات بر ارزش افزوده، اظهارنامه
  4. جدول تراکنش‌های مالیاتی: نمایش و محاسبه مالیات
  5. نمودار روند مالیاتی: نمایش روند مالیات در زمان
  
  تمام بخش‌ها به صورت ریسپانسیو و با طراحی مدرن ارائه شده‌اند.
--> 
