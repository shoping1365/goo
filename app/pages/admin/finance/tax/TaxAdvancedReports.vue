<template>
  <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
    <!-- هدر -->
    <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-6 mb-6">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">گزارش‌گیری پیشرفته</h3>
        <p class="text-sm text-gray-600">گزارش‌های دوره‌ای، مقایسه‌ای و تحلیلی پیشرفته</p>
      </div>
      
      <!-- دکمه‌های عملیاتی -->
      <div class="flex gap-2">
        <button 
          @click="generateReport"
          class="inline-flex items-center gap-2 px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg transition-colors duration-200"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
          </svg>
          تولید گزارش
        </button>
        <button 
          @click="exportToExcel"
          class="inline-flex items-center gap-2 px-4 py-2 bg-green-600 hover:bg-green-700 text-white rounded-lg transition-colors duration-200"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
          </svg>
          خروجی Excel
        </button>
        <button 
          @click="exportToPDF"
          class="inline-flex items-center gap-2 px-4 py-2 bg-red-600 hover:bg-red-700 text-white rounded-lg transition-colors duration-200"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21h10a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v14a2 2 0 002 2z"></path>
          </svg>
          خروجی PDF
        </button>
      </div>
    </div>

    <!-- تنظیمات گزارش -->
    <div class="bg-gray-50 rounded-lg p-6 mb-6">
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نوع گزارش</label>
          <select 
            v-model="reportConfig.type"
            @change="updateReportConfig"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="periodic">دوره‌ای</option>
            <option value="comparative">مقایسه‌ای</option>
            <option value="analytical">تحلیلی</option>
            <option value="custom">سفارشی</option>
          </select>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">دوره زمانی</label>
          <select 
            v-model="reportConfig.period"
            @change="updateReportConfig"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="daily">روزانه</option>
            <option value="weekly">هفتگی</option>
            <option value="monthly">ماهانه</option>
            <option value="quarterly">فصلی</option>
            <option value="yearly">سالانه</option>
          </select>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">سطح جزئیات</label>
          <select 
            v-model="reportConfig.detailLevel"
            @change="updateReportConfig"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="summary">خلاصه</option>
            <option value="detailed">تفصیلی</option>
            <option value="comprehensive">جامع</option>
          </select>
        </div>
      </div>

      <!-- تنظیمات اضافی -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mt-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">از تاریخ</label>
          <input 
            v-model="reportConfig.dateFrom"
            type="date"
            @change="updateReportConfig"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تا تاریخ</label>
          <input 
            v-model="reportConfig.dateTo"
            type="date"
            @change="updateReportConfig"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">دسته‌بندی</label>
          <select 
            v-model="reportConfig.category"
            @change="updateReportConfig"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه دسته‌ها</option>
            <option value="products">محصولات</option>
            <option value="services">خدمات</option>
            <option value="digital">دیجیتال</option>
          </select>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">منطقه</label>
          <select 
            v-model="reportConfig.region"
            @change="updateReportConfig"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">همه مناطق</option>
            <option value="tehran">تهران</option>
            <option value="isfahan">اصفهان</option>
            <option value="mashhad">مشهد</option>
          </select>
        </div>
      </div>

      <!-- گزینه‌های اضافی -->
      <div class="mt-4 space-y-3">
        <div class="flex flex-wrap gap-6">
          <label class="flex items-center">
            <input 
              v-model="reportConfig.includeCharts"
              type="checkbox"
              class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
            />
            <span class="mr-2 text-sm text-gray-700">شامل نمودارها</span>
          </label>
          
          <label class="flex items-center">
            <input 
              v-model="reportConfig.includeTables"
              type="checkbox"
              class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
            />
            <span class="mr-2 text-sm text-gray-700">شامل جداول</span>
          </label>
          
          <label class="flex items-center">
            <input 
              v-model="reportConfig.includeSummary"
              type="checkbox"
              class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
            />
            <span class="mr-2 text-sm text-gray-700">شامل خلاصه</span>
          </label>
          
          <label class="flex items-center">
            <input 
              v-model="reportConfig.includeRecommendations"
              type="checkbox"
              class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
          />
            <span class="mr-2 text-sm text-gray-700">شامل توصیه‌ها</span>
          </label>
        </div>
      </div>
    </div>

    <!-- تب‌های گزارش -->
    <div class="mb-6">
      <div class="border-b border-gray-200">
        <nav class="-mb-px flex space-x-8">
          <button 
            v-for="tab in reportTabs" 
            :key="tab.id"
            @click="activeReportTab = tab.id"
            :class="[
              'py-2 px-1 border-b-2 font-medium text-sm transition-colors duration-200',
              activeReportTab === tab.id 
                ? 'border-blue-500 text-blue-600' 
                : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
            ]"
          >
            {{ tab.label }}
          </button>
        </nav>
      </div>
    </div>

    <!-- گزارش دوره‌ای -->
    <div v-if="activeReportTab === 'periodic'">
      <div class="space-y-6">
        <!-- خلاصه دوره‌ای -->
        <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
          <div class="bg-blue-50 rounded-lg p-6">
            <div class="text-2xl font-bold text-blue-600">{{ formatCurrency(periodicReport.totalRevenue) }}</div>
            <div class="text-sm text-blue-700">کل درآمد</div>
            <div class="text-xs text-blue-600 mt-1">{{ periodicReport.revenueGrowth > 0 ? '+' : '' }}{{ periodicReport.revenueGrowth }}% نسبت به دوره قبل</div>
          </div>

          <div class="bg-green-50 rounded-lg p-6">
            <div class="text-2xl font-bold text-green-600">{{ formatCurrency(periodicReport.totalTax) }}</div>
            <div class="text-sm text-green-700">کل مالیات</div>
            <div class="text-xs text-green-600 mt-1">{{ periodicReport.taxGrowth > 0 ? '+' : '' }}{{ periodicReport.taxGrowth }}% نسبت به دوره قبل</div>
          </div>

          <div class="bg-purple-50 rounded-lg p-6">
            <div class="text-2xl font-bold text-purple-600">{{ periodicReport.totalTransactions }}</div>
            <div class="text-sm text-purple-700">کل تراکنش‌ها</div>
            <div class="text-xs text-purple-600 mt-1">{{ periodicReport.transactionGrowth > 0 ? '+' : '' }}{{ periodicReport.transactionGrowth }}% نسبت به دوره قبل</div>
          </div>

          <div class="bg-orange-50 rounded-lg p-6">
            <div class="text-2xl font-bold text-orange-600">{{ periodicReport.avgTransactionValue }}</div>
            <div class="text-sm text-orange-700">میانگین تراکنش</div>
            <div class="text-xs text-orange-600 mt-1">{{ periodicReport.avgGrowth > 0 ? '+' : '' }}{{ periodicReport.avgGrowth }}% نسبت به دوره قبل</div>
          </div>
        </div>

        <!-- نمودار روند دوره‌ای -->
        <div class="bg-gray-50 rounded-lg p-6">
          <h4 class="font-medium text-gray-900 mb-4">روند دوره‌ای</h4>
          <div class="h-64 flex items-end justify-center gap-2">
            <div 
              v-for="(period, index) in periodicReport.periods" 
              :key="index"
              class="flex flex-col items-center"
            >
              <div class="flex flex-col gap-1">
                <div 
                  class="w-6 bg-blue-500 rounded-t-lg transition-all duration-300 hover:bg-blue-600"
                  :style="{ height: `${(period.revenue / periodicReport.maxRevenue) * 150}px` }"
                  :title="`درآمد: ${formatCurrency(period.revenue)}`"
                ></div>
                <div 
                  class="w-6 bg-green-500 rounded-t-lg transition-all duration-300 hover:bg-green-600"
                  :style="{ height: `${(period.tax / periodicReport.maxTax) * 150}px` }"
                  :title="`مالیات: ${formatCurrency(period.tax)}`"
                ></div>
              </div>
              <div class="mt-2 text-xs text-gray-600">{{ period.name }}</div>
            </div>
          </div>
        </div>

        <!-- جدول تفصیلی دوره‌ای -->
        <div class="overflow-x-auto">
          <table class="w-full text-sm">
            <thead>
              <tr class="border-b border-gray-200 bg-gray-50">
                <th class="text-right py-3 px-4 font-medium text-gray-600">دوره</th>
                <th class="text-right py-3 px-4 font-medium text-gray-600">درآمد</th>
                <th class="text-right py-3 px-4 font-medium text-gray-600">مالیات VAT</th>
                <th class="text-right py-3 px-4 font-medium text-gray-600">مالیات درآمد</th>
                <th class="text-right py-3 px-4 font-medium text-gray-600">معافیت‌ها</th>
                <th class="text-right py-3 px-4 font-medium text-gray-600">تراکنش‌ها</th>
                <th class="text-right py-3 px-4 font-medium text-gray-600">رشد</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="period in periodicReport.periods" :key="period.name" class="border-b border-gray-100">
                <td class="py-3 px-4 text-gray-900 font-medium">{{ period.name }}</td>
                <td class="py-3 px-4 text-gray-900">{{ formatCurrency(period.revenue) }}</td>
                <td class="py-3 px-4 text-blue-600">{{ formatCurrency(period.vatTax) }}</td>
                <td class="py-3 px-4 text-green-600">{{ formatCurrency(period.incomeTax) }}</td>
                <td class="py-3 px-4 text-yellow-600">{{ formatCurrency(period.exemptions) }}</td>
                <td class="py-3 px-4 text-gray-900">{{ period.transactions }}</td>
                <td class="py-3 px-4">
                  <span :class="period.growth >= 0 ? 'text-green-600' : 'text-red-600'">
                    {{ period.growth >= 0 ? '+' : '' }}{{ period.growth }}%
                  </span>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>

    <!-- گزارش مقایسه‌ای -->
    <div v-else-if="activeReportTab === 'comparative'">
      <div class="space-y-6">
        <!-- انتخاب دوره‌های مقایسه -->
        <div class="bg-gray-50 rounded-lg p-6">
          <h4 class="font-medium text-gray-900 mb-4">انتخاب دوره‌های مقایسه</h4>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">دوره اول</label>
              <div class="grid grid-cols-2 gap-2">
                <input 
                  v-model="comparativeReport.period1.from"
                  type="date"
                  class="px-3 py-2 border border-gray-300 rounded-lg text-sm"
                />
                <input 
                  v-model="comparativeReport.period1.to"
                  type="date"
                  class="px-3 py-2 border border-gray-300 rounded-lg text-sm"
                />
              </div>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">دوره دوم</label>
              <div class="grid grid-cols-2 gap-2">
                <input 
                  v-model="comparativeReport.period2.from"
                  type="date"
                  class="px-3 py-2 border border-gray-300 rounded-lg text-sm"
                />
                <input 
                  v-model="comparativeReport.period2.to"
                  type="date"
                  class="px-3 py-2 border border-gray-300 rounded-lg text-sm"
                />
              </div>
            </div>
          </div>
        </div>

        <!-- مقایسه شاخص‌ها -->
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div class="bg-blue-50 rounded-lg p-6">
            <div class="text-lg font-medium text-blue-900 mb-2">درآمد</div>
            <div class="space-y-2">
              <div class="flex justify-between">
                <span class="text-sm text-blue-700">دوره اول:</span>
                <span class="font-medium">{{ formatCurrency(comparativeReport.revenue.period1) }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-sm text-blue-700">دوره دوم:</span>
                <span class="font-medium">{{ formatCurrency(comparativeReport.revenue.period2) }}</span>
              </div>
              <div class="border-t pt-2">
                <div class="flex justify-between">
                  <span class="text-sm font-medium">تغییر:</span>
                  <span :class="comparativeReport.revenue.change >= 0 ? 'text-green-600' : 'text-red-600'">
                    {{ comparativeReport.revenue.change >= 0 ? '+' : '' }}{{ comparativeReport.revenue.change }}%
                  </span>
                </div>
              </div>
            </div>
          </div>

          <div class="bg-green-50 rounded-lg p-6">
            <div class="text-lg font-medium text-green-900 mb-2">مالیات</div>
            <div class="space-y-2">
              <div class="flex justify-between">
                <span class="text-sm text-green-700">دوره اول:</span>
                <span class="font-medium">{{ formatCurrency(comparativeReport.tax.period1) }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-sm text-green-700">دوره دوم:</span>
                <span class="font-medium">{{ formatCurrency(comparativeReport.tax.period2) }}</span>
              </div>
              <div class="border-t pt-2">
                <div class="flex justify-between">
                  <span class="text-sm font-medium">تغییر:</span>
                  <span :class="comparativeReport.tax.change >= 0 ? 'text-green-600' : 'text-red-600'">
                    {{ comparativeReport.tax.change >= 0 ? '+' : '' }}{{ comparativeReport.tax.change }}%
                  </span>
                </div>
              </div>
            </div>
          </div>

          <div class="bg-purple-50 rounded-lg p-6">
            <div class="text-lg font-medium text-purple-900 mb-2">تراکنش‌ها</div>
            <div class="space-y-2">
              <div class="flex justify-between">
                <span class="text-sm text-purple-700">دوره اول:</span>
                <span class="font-medium">{{ comparativeReport.transactions.period1 }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-sm text-purple-700">دوره دوم:</span>
                <span class="font-medium">{{ comparativeReport.transactions.period2 }}</span>
              </div>
              <div class="border-t pt-2">
                <div class="flex justify-between">
                  <span class="text-sm font-medium">تغییر:</span>
                  <span :class="comparativeReport.transactions.change >= 0 ? 'text-green-600' : 'text-red-600'">
                    {{ comparativeReport.transactions.change >= 0 ? '+' : '' }}{{ comparativeReport.transactions.change }}%
                  </span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- نمودار مقایسه‌ای -->
        <div class="bg-gray-50 rounded-lg p-6">
          <h4 class="font-medium text-gray-900 mb-4">نمودار مقایسه‌ای</h4>
          <div class="h-64 flex items-end justify-center gap-8">
            <div class="flex flex-col items-center">
              <div class="text-sm font-medium text-gray-700 mb-2">دوره اول</div>
              <div class="space-y-2">
                <div class="flex items-center gap-2">
                  <div class="w-4 h-4 bg-blue-500 rounded"></div>
                  <div class="w-20 bg-blue-500 rounded-t-lg h-[120px]"></div>
                </div>
                <div class="flex items-center gap-2">
                  <div class="w-4 h-4 bg-green-500 rounded"></div>
                  <div class="w-20 bg-green-500 rounded-t-lg h-[80px]"></div>
                </div>
              </div>
            </div>
            <div class="flex flex-col items-center">
              <div class="text-sm font-medium text-gray-700 mb-2">دوره دوم</div>
              <div class="space-y-2">
                <div class="flex items-center gap-2">
                  <div class="w-4 h-4 bg-blue-500 rounded"></div>
                  <div class="w-20 bg-blue-500 rounded-t-lg h-[140px]"></div>
                </div>
                <div class="flex items-center gap-2">
                  <div class="w-4 h-4 bg-green-500 rounded"></div>
                  <div class="w-20 bg-green-500 rounded-t-lg h-[100px]"></div>
                </div>
              </div>
            </div>
          </div>
          <div class="flex justify-center gap-8 mt-4">
            <div class="flex items-center gap-2">
              <div class="w-4 h-4 bg-blue-500 rounded"></div>
              <span class="text-sm text-gray-600">درآمد</span>
            </div>
            <div class="flex items-center gap-2">
              <div class="w-4 h-4 bg-green-500 rounded"></div>
              <span class="text-sm text-gray-600">مالیات</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- گزارش تحلیلی -->
    <div v-else-if="activeReportTab === 'analytical'">
      <div class="space-y-6">
        <!-- شاخص‌های کلیدی -->
        <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
          <div class="bg-indigo-50 rounded-lg p-6">
            <div class="text-2xl font-bold text-indigo-600">{{ analyticalReport.taxEfficiency }}%</div>
            <div class="text-sm text-indigo-700">کارایی مالیاتی</div>
          </div>

          <div class="bg-pink-50 rounded-lg p-6">
            <div class="text-2xl font-bold text-pink-600">{{ analyticalReport.complianceRate }}%</div>
            <div class="text-sm text-pink-700">نرخ انطباق</div>
          </div>

          <div class="bg-teal-50 rounded-lg p-6">
            <div class="text-2xl font-bold text-teal-600">{{ analyticalReport.riskScore }}</div>
            <div class="text-sm text-teal-700">امتیاز ریسک</div>
          </div>

          <div class="bg-amber-50 rounded-lg p-6">
            <div class="text-2xl font-bold text-amber-600">{{ analyticalReport.optimizationScore }}%</div>
            <div class="text-sm text-amber-700">امتیاز بهینه‌سازی</div>
          </div>
        </div>

        <!-- تحلیل روندها -->
        <div class="bg-gray-50 rounded-lg p-6">
          <h4 class="font-medium text-gray-900 mb-4">تحلیل روندها</h4>
          <div class="space-y-4">
            <div v-for="trend in analyticalReport.trends" :key="trend.name" class="bg-white rounded-lg p-3">
              <div class="flex items-center justify-between mb-2">
                <span class="font-medium text-gray-900">{{ trend.name }}</span>
                <span :class="trend.direction === 'up' ? 'text-green-600' : 'text-red-600'">
                  {{ trend.direction === 'up' ? '↗' : '↘' }} {{ trend.change }}%
                </span>
              </div>
              <div class="text-sm text-gray-600">{{ trend.description }}</div>
            </div>
          </div>
        </div>

        <!-- توصیه‌های تحلیلی -->
        <div class="bg-gray-50 rounded-lg p-6">
          <h4 class="font-medium text-gray-900 mb-4">توصیه‌های تحلیلی</h4>
          <div class="space-y-3">
            <div v-for="recommendation in analyticalReport.recommendations" :key="recommendation.id" class="bg-white rounded-lg p-3 border-l-4" :class="getRecommendationBorderClass(recommendation.priority)">
              <div class="flex items-start justify-between">
                <div>
                  <div class="font-medium text-gray-900">{{ recommendation.title }}</div>
                  <div class="text-sm text-gray-600 mt-1">{{ recommendation.description }}</div>
                </div>
                <span :class="getRecommendationBadgeClass(recommendation.priority)" class="px-2 py-1 rounded-full text-xs font-medium">
                  {{ getRecommendationPriorityLabel(recommendation.priority) }}
                </span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- گزارش سفارشی -->
    <div v-else-if="activeReportTab === 'custom'">
      <div class="space-y-6">
        <!-- سازنده گزارش سفارشی -->
        <div class="bg-gray-50 rounded-lg p-6">
          <h4 class="font-medium text-gray-900 mb-4">سازنده گزارش سفارشی</h4>
          
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">عنوان گزارش</label>
              <input 
                v-model="customReport.title"
                type="text"
                placeholder="عنوان گزارش را وارد کنید..."
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات</label>
              <textarea 
                v-model="customReport.description"
                rows="3"
                placeholder="توضیحات گزارش..."
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              ></textarea>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">فیلدهای مورد نیاز</label>
              <div class="space-y-2">
                <label v-for="field in customReport.availableFields" :key="field.id" class="flex items-center">
                  <input 
                    v-model="customReport.selectedFields"
                    :value="field.id"
                    type="checkbox"
                    class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                  />
                  <span class="mr-2 text-sm text-gray-700">{{ field.name }}</span>
                </label>
              </div>
            </div>

            <div class="flex gap-3">
              <button 
                @click="saveCustomReport"
                class="px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg"
              >
                ذخیره گزارش
              </button>
              <button 
                @click="loadCustomReport"
                class="px-4 py-2 bg-gray-200 hover:bg-gray-300 text-gray-700 rounded-lg"
              >
                بارگذاری گزارش
              </button>
            </div>
          </div>
        </div>

        <!-- پیش‌نمایش گزارش سفارشی -->
        <div v-if="customReport.preview" class="bg-gray-50 rounded-lg p-6">
          <h4 class="font-medium text-gray-900 mb-4">پیش‌نمایش گزارش</h4>
          <div class="bg-white rounded-lg p-6">
            <div class="text-lg font-medium text-gray-900 mb-2">{{ customReport.title }}</div>
            <div class="text-sm text-gray-600 mb-4">{{ customReport.description }}</div>
            <div class="text-sm text-gray-700">
              فیلدهای انتخاب شده: {{ customReport.selectedFields.join(', ') }}
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'

// تب‌های گزارش
const reportTabs = [
  { id: 'periodic', label: 'دوره‌ای' },
  { id: 'comparative', label: 'مقایسه‌ای' },
  { id: 'analytical', label: 'تحلیلی' },
  { id: 'custom', label: 'سفارشی' }
]

const activeReportTab = ref('periodic')

// تنظیمات گزارش
const reportConfig = ref({
  type: 'periodic',
  period: 'monthly',
  detailLevel: 'detailed',
  dateFrom: '',
  dateTo: '',
  category: '',
  region: '',
  includeCharts: true,
  includeTables: true,
  includeSummary: true,
  includeRecommendations: true
})

// گزارش دوره‌ای
const periodicReport = ref({
  totalRevenue: 15000000000,
  totalTax: 1350000000,
  totalTransactions: 15420,
  avgTransactionValue: 972000,
  revenueGrowth: 12.5,
  taxGrowth: 8.3,
  transactionGrowth: 15.2,
  avgGrowth: -2.1,
  maxRevenue: 3000000000,
  maxTax: 300000000,
  periods: [
    { name: 'فروردین', revenue: 2500000000, tax: 225000000, vatTax: 180000000, incomeTax: 45000000, exemptions: 50000000, transactions: 2500, growth: 12.5 },
    { name: 'اردیبهشت', revenue: 2800000000, tax: 252000000, vatTax: 201600000, incomeTax: 50400000, exemptions: 55000000, transactions: 2800, growth: 15.2 },
    { name: 'خرداد', revenue: 2600000000, tax: 234000000, vatTax: 187200000, incomeTax: 46800000, exemptions: 48000000, transactions: 2600, growth: 8.7 },
    { name: 'تیر', revenue: 3000000000, tax: 270000000, vatTax: 216000000, incomeTax: 54000000, exemptions: 60000000, transactions: 3000, growth: 18.3 },
    { name: 'مرداد', revenue: 2750000000, tax: 247500000, vatTax: 198000000, incomeTax: 49500000, exemptions: 52000000, transactions: 2750, growth: 10.1 },
    { name: 'شهریور', revenue: 2900000000, tax: 261000000, vatTax: 208800000, incomeTax: 52200000, exemptions: 58000000, transactions: 2900, growth: 14.2 }
  ]
})

// گزارش مقایسه‌ای
const comparativeReport = ref({
  period1: { from: '2024-01-01', to: '2024-03-31' },
  period2: { from: '2024-04-01', to: '2024-06-30' },
  revenue: { period1: 7500000000, period2: 8500000000, change: 13.3 },
  tax: { period1: 675000000, period2: 765000000, change: 13.3 },
  transactions: { period1: 7500, period2: 8500, change: 13.3 }
})

// گزارش تحلیلی
const analyticalReport = ref({
  taxEfficiency: 92.5,
  complianceRate: 98.7,
  riskScore: 7.2,
  optimizationScore: 85.3,
  trends: [
    { name: 'افزایش نرخ مالیات بر ارزش افزوده', direction: 'up', change: 15.2, description: 'رشد قابل توجه در جمع‌آوری مالیات بر ارزش افزوده' },
    { name: 'کاهش معافیت‌های مالیاتی', direction: 'down', change: 8.7, description: 'کاهش در تعداد و مبلغ معافیت‌های اعمال شده' },
    { name: 'بهبود انطباق مالیاتی', direction: 'up', change: 12.3, description: 'افزایش نرخ انطباق با قوانین مالیاتی' }
  ],
  recommendations: [
    { id: 1, title: 'بهینه‌سازی نرخ‌های مالیاتی', description: 'بررسی و تنظیم نرخ‌های مالیاتی برای بهبود کارایی', priority: 'high' },
    { id: 2, title: 'افزایش معافیت‌های هدفمند', description: 'اعمال معافیت‌های بیشتر برای بخش‌های خاص', priority: 'medium' },
    { id: 3, title: 'بهبود سیستم نظارت', description: 'تقویت سیستم نظارت و کنترل مالیاتی', priority: 'low' }
  ]
})

// گزارش سفارشی
const customReport = ref({
  title: '',
  description: '',
  selectedFields: [],
  availableFields: [
    { id: 'revenue', name: 'درآمد' },
    { id: 'tax', name: 'مالیات' },
    { id: 'transactions', name: 'تراکنش‌ها' },
    { id: 'exemptions', name: 'معافیت‌ها' },
    { id: 'customers', name: 'مشتریان' },
    { id: 'regions', name: 'مناطق' },
    { id: 'categories', name: 'دسته‌بندی‌ها' }
  ],
  preview: false
})

// فرمت مبلغ
const formatCurrency = (amount: number): string => {
  return new Intl.NumberFormat('fa-IR', {
    style: 'currency',
    currency: 'IRR',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount)
}

// بروزرسانی تنظیمات گزارش
const updateReportConfig = () => {
  // TODO: بروزرسانی گزارش بر اساس تنظیمات
  console.log('تنظیمات گزارش بروزرسانی شد:', reportConfig.value)
}

// تولید گزارش
const generateReport = async () => {
  try {
    // TODO: تولید گزارش بر اساس تنظیمات
    console.log('گزارش تولید شد:', reportConfig.value)
  } catch (error) {
    console.error('خطا در تولید گزارش:', error)
  }
}

// خروجی Excel
const exportToExcel = () => {
  // TODO: خروجی Excel
  console.log('گزارش به Excel صادر شد')
}

// خروجی PDF
const exportToPDF = () => {
  // TODO: خروجی PDF
  console.log('گزارش به PDF صادر شد')
}

// کلاس border توصیه
const getRecommendationBorderClass = (priority: string) => {
  const classes = {
    high: 'border-red-500',
    medium: 'border-yellow-500',
    low: 'border-green-500'
  }
  return classes[priority] || 'border-gray-500'
}

// کلاس badge توصیه
const getRecommendationBadgeClass = (priority: string) => {
  const classes = {
    high: 'bg-red-100 text-red-700',
    medium: 'bg-yellow-100 text-yellow-700',
    low: 'bg-green-100 text-green-700'
  }
  return classes[priority] || 'bg-gray-100 text-gray-700'
}

// برچسب اولویت توصیه
const getRecommendationPriorityLabel = (priority: string) => {
  const labels = {
    high: 'بالا',
    medium: 'متوسط',
    low: 'پایین'
  }
  return labels[priority] || priority
}

// ذخیره گزارش سفارشی
const saveCustomReport = () => {
  // TODO: ذخیره گزارش سفارشی
  console.log('گزارش سفارشی ذخیره شد:', customReport.value)
  customReport.value.preview = true
}

// بارگذاری گزارش سفارشی
const loadCustomReport = () => {
  // TODO: بارگذاری گزارش سفارشی
  console.log('گزارش سفارشی بارگذاری شد')
}

// بارگذاری اولیه
onMounted(() => {
  // تنظیم تاریخ‌های پیش‌فرض
  const now = new Date()
  const firstDay = new Date(now.getFullYear(), now.getMonth(), 1)
  const lastDay = new Date(now.getFullYear(), now.getMonth() + 1, 0)
  
  reportConfig.value.dateFrom = firstDay.toISOString().split('T')[0]
  reportConfig.value.dateTo = lastDay.toISOString().split('T')[0]
  
  // TODO: بارگذاری گزارش‌ها از API
})
</script>

<!--
  کامپوننت گزارش‌گیری پیشرفته
  شامل:
  1. گزارش‌های دوره‌ای (ماهانه، فصلی، سالانه)
  2. گزارش‌های مقایسه‌ای
  3. گزارش‌های تحلیلی پیشرفته
  4. گزارش‌های سفارشی
  5. خروجی Excel/PDF
  6. نمودارها و آمار پیشرفته
  7. طراحی مدرن و کاملاً ریسپانسیو
--> 
