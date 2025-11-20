<template>
  <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
    <!-- هدر -->
    <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-6 mb-6">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">گزارش‌های مالیاتی</h3>
        <p class="text-sm text-gray-600">گزارش‌های جامع مالیات بر ارزش افزوده، درآمد و معافیت‌ها</p>
      </div>
      
      <!-- دکمه‌های عملیاتی -->
      <div class="flex gap-2">
        <button 
          class="inline-flex items-center gap-2 px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg transition-colors duration-200"
          @click="generateAllReports"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
          </svg>
          تولید تمام گزارش‌ها
        </button>
        <button 
          class="inline-flex items-center gap-2 px-4 py-2 bg-green-600 hover:bg-green-700 text-white rounded-lg transition-colors duration-200"
          @click="exportReports"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
          </svg>
          خروجی Excel
        </button>
      </div>
    </div>

    <!-- فیلترهای گزارش -->
    <div class="bg-gray-50 rounded-lg p-6 mb-6">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">دوره زمانی</label>
          <select 
            v-model="reportFilters.period"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            @change="updateReports"
          >
            <option value="monthly">ماهانه</option>
            <option value="quarterly">فصلی</option>
            <option value="yearly">سالانه</option>
            <option value="custom">سفارشی</option>
          </select>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">از تاریخ</label>
          <input 
            v-model="reportFilters.dateFrom"
            type="date"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            @change="updateReports"
          />
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تا تاریخ</label>
          <input 
            v-model="reportFilters.dateTo"
            type="date"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            @change="updateReports"
          />
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">نوع گزارش</label>
          <select 
            v-model="reportFilters.reportType"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            @change="updateReports"
          >
            <option value="all">همه گزارش‌ها</option>
            <option value="vat">مالیات بر ارزش افزوده</option>
            <option value="income">مالیات بر درآمد</option>
            <option value="exemptions">معافیت‌ها</option>
            <option value="detailed">تفصیلی</option>
          </select>
        </div>
      </div>
    </div>

    <!-- تب‌های گزارش‌ها -->
    <div class="mb-6">
      <div class="border-b border-gray-200">
        <nav class="-mb-px flex space-x-8">
          <button 
            v-for="tab in reportTabs" 
            :key="tab.id"
            :class="[
              'py-2 px-1 border-b-2 font-medium text-sm transition-colors duration-200',
              activeReportTab === tab.id 
                ? 'border-blue-500 text-blue-600' 
                : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
            ]"
            @click="activeReportTab = tab.id"
          >
            {{ tab.label }}
          </button>
        </nav>
      </div>
    </div>

    <!-- گزارش مالیات بر ارزش افزوده -->
    <div v-if="activeReportTab === 'vat'">
      <div class="space-y-6">
        <!-- خلاصه VAT -->
        <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
          <div class="bg-blue-50 rounded-lg p-6">
            <div class="flex items-center justify-between">
              <div>
                <div class="text-2xl font-bold text-blue-600">{{ formatCurrency(vatReport.totalVat) }}</div>
                <div class="text-sm text-blue-700">کل مالیات VAT</div>
              </div>
              <div class="w-10 h-10 bg-blue-500 rounded-lg flex items-center justify-center">
                <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
                </svg>
              </div>
            </div>
          </div>

          <div class="bg-green-50 rounded-lg p-6">
            <div class="flex items-center justify-between">
              <div>
                <div class="text-2xl font-bold text-green-600">{{ vatReport.totalTransactions }}</div>
                <div class="text-sm text-green-700">کل تراکنش‌ها</div>
              </div>
              <div class="w-10 h-10 bg-green-500 rounded-lg flex items-center justify-center">
                <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"></path>
                </svg>
              </div>
            </div>
          </div>

          <div class="bg-purple-50 rounded-lg p-6">
            <div class="flex items-center justify-between">
              <div>
                <div class="text-2xl font-bold text-purple-600">{{ vatReport.avgVatRate }}%</div>
                <div class="text-sm text-purple-700">میانگین نرخ VAT</div>
              </div>
              <div class="w-10 h-10 bg-purple-500 rounded-lg flex items-center justify-center">
                <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
                </svg>
              </div>
            </div>
          </div>

          <div class="bg-orange-50 rounded-lg p-6">
            <div class="flex items-center justify-between">
              <div>
                <div class="text-2xl font-bold text-orange-600">{{ formatCurrency(vatReport.totalAmount) }}</div>
                <div class="text-sm text-orange-700">کل مبلغ پایه</div>
              </div>
              <div class="w-10 h-10 bg-orange-500 rounded-lg flex items-center justify-center">
                <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
                </svg>
              </div>
            </div>
          </div>
        </div>

        <!-- نمودار VAT -->
        <div class="bg-gray-50 rounded-lg p-6">
          <h4 class="font-medium text-gray-900 mb-4">روند مالیات بر ارزش افزوده</h4>
          <div class="h-64 flex items-end justify-center gap-2">
            <div 
              v-for="(item, index) in vatReport.monthlyData" 
              :key="index"
              class="flex flex-col items-center"
            >
              <div 
                class="w-8 bg-blue-500 rounded-t-lg transition-all duration-300 hover:bg-blue-600"
                :style="{ height: `${(item.amount / vatReport.maxAmount) * 200}px` }"
                :title="`${item.month}: ${formatCurrency(item.amount)}`"
              ></div>
              <div class="mt-2 text-xs text-gray-600">{{ item.month }}</div>
            </div>
          </div>
        </div>

        <!-- جدول تفصیلی VAT -->
        <div class="overflow-x-auto">
          <table class="w-full text-sm">
            <thead>
              <tr class="border-b border-gray-200 bg-gray-50">
                <th class="text-right py-3 px-4 font-medium text-gray-600">دسته‌بندی</th>
                <th class="text-right py-3 px-4 font-medium text-gray-600">تعداد تراکنش</th>
                <th class="text-right py-3 px-4 font-medium text-gray-600">مبلغ پایه</th>
                <th class="text-right py-3 px-4 font-medium text-gray-600">مالیات VAT</th>
                <th class="text-right py-3 px-4 font-medium text-gray-600">نرخ متوسط</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="category in vatReport.categories" :key="category.name" class="border-b border-gray-100">
                <td class="py-3 px-4 text-gray-900">{{ category.name }}</td>
                <td class="py-3 px-4 text-gray-900">{{ category.transactions }}</td>
                <td class="py-3 px-4 text-gray-900">{{ formatCurrency(category.baseAmount) }}</td>
                <td class="py-3 px-4 text-blue-600 font-medium">{{ formatCurrency(category.vatAmount) }}</td>
                <td class="py-3 px-4 text-gray-900">{{ category.avgRate }}%</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>

    <!-- گزارش مالیات بر درآمد -->
    <div v-else-if="activeReportTab === 'income'">
      <div class="space-y-6">
        <!-- خلاصه مالیات بر درآمد -->
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div class="bg-green-50 rounded-lg p-6">
            <div class="flex items-center justify-between">
              <div>
                <div class="text-2xl font-bold text-green-600">{{ formatCurrency(incomeReport.totalIncomeTax) }}</div>
                <div class="text-sm text-green-700">کل مالیات بر درآمد</div>
              </div>
              <div class="w-10 h-10 bg-green-500 rounded-lg flex items-center justify-center">
                <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path>
                </svg>
              </div>
            </div>
          </div>

          <div class="bg-blue-50 rounded-lg p-6">
            <div class="flex items-center justify-between">
              <div>
                <div class="text-2xl font-bold text-blue-600">{{ incomeReport.totalTaxpayers }}</div>
                <div class="text-sm text-blue-700">تعداد مودیان</div>
              </div>
              <div class="w-10 h-10 bg-blue-500 rounded-lg flex items-center justify-center">
                <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
                </svg>
              </div>
            </div>
          </div>

          <div class="bg-purple-50 rounded-lg p-6">
            <div class="flex items-center justify-between">
              <div>
                <div class="text-2xl font-bold text-purple-600">{{ formatCurrency(incomeReport.avgTaxPerPerson) }}</div>
                <div class="text-sm text-purple-700">میانگین مالیات هر نفر</div>
              </div>
              <div class="w-10 h-10 bg-purple-500 rounded-lg flex items-center justify-center">
                <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
                </svg>
              </div>
            </div>
          </div>
        </div>

        <!-- توزیع مالیات بر درآمد -->
        <div class="bg-gray-50 rounded-lg p-6">
          <h4 class="font-medium text-gray-900 mb-4">توزیع مالیات بر درآمد بر اساس طبقه‌بندی</h4>
          <div class="space-y-3">
            <div v-for="bracket in incomeReport.taxBrackets" :key="bracket.range" class="flex items-center justify-between">
              <div class="flex items-center gap-3">
                <div class="w-4 h-4 bg-green-500 rounded-full"></div>
                <span class="text-sm text-gray-700">{{ bracket.range }}</span>
              </div>
              <div class="flex items-center gap-6">
                <span class="text-sm text-gray-600">{{ bracket.count }} نفر</span>
                <span class="text-sm font-medium text-green-600">{{ formatCurrency(bracket.totalTax) }}</span>
                <span class="text-sm text-gray-500">{{ bracket.percentage }}%</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- گزارش معافیت‌ها -->
    <div v-else-if="activeReportTab === 'exemptions'">
      <div class="space-y-6">
        <!-- خلاصه معافیت‌ها -->
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div class="bg-yellow-50 rounded-lg p-6">
            <div class="flex items-center justify-between">
              <div>
                <div class="text-2xl font-bold text-yellow-600">{{ formatCurrency(exemptionsReport.totalExemptedAmount) }}</div>
                <div class="text-sm text-yellow-700">کل مبلغ معاف شده</div>
              </div>
              <div class="w-10 h-10 bg-yellow-500 rounded-lg flex items-center justify-center">
                <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
              </div>
            </div>
          </div>

          <div class="bg-green-50 rounded-lg p-6">
            <div class="flex items-center justify-between">
              <div>
                <div class="text-2xl font-bold text-green-600">{{ exemptionsReport.totalExemptions }}</div>
                <div class="text-sm text-green-700">تعداد معافیت‌ها</div>
              </div>
              <div class="w-10 h-10 bg-green-500 rounded-lg flex items-center justify-center">
                <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
              </div>
            </div>
          </div>

          <div class="bg-blue-50 rounded-lg p-6">
            <div class="flex items-center justify-between">
              <div>
                <div class="text-2xl font-bold text-blue-600">{{ exemptionsReport.exemptionRate }}%</div>
                <div class="text-sm text-blue-700">نرخ معافیت</div>
              </div>
              <div class="w-10 h-10 bg-blue-500 rounded-lg flex items-center justify-center">
                <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
                </svg>
              </div>
            </div>
          </div>
        </div>

        <!-- نمودار معافیت‌ها -->
        <div class="bg-gray-50 rounded-lg p-6">
          <h4 class="font-medium text-gray-900 mb-4">توزیع معافیت‌ها بر اساس نوع</h4>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div class="space-y-3">
              <div v-for="exemption in exemptionsReport.exemptionTypes" :key="exemption.type" class="flex items-center justify-between">
                <div class="flex items-center gap-3">
                  <div class="w-3 h-3 rounded-full" :style="{ backgroundColor: exemption.color }"></div>
                  <span class="text-sm text-gray-700">{{ exemption.name }}</span>
                </div>
                <div class="flex items-center gap-6">
                  <span class="text-sm font-medium text-gray-900">{{ formatCurrency(exemption.amount) }}</span>
                  <span class="text-sm text-gray-500">{{ exemption.percentage }}%</span>
                </div>
              </div>
            </div>
            <div class="flex items-center justify-center">
              <!-- نمودار دایره‌ای ساده -->
              <div class="relative w-32 h-32">
                <svg class="w-32 h-32 transform -rotate-90">
                  <circle
                    cx="64"
                    cy="64"
                    r="54"
                    stroke="currentColor"
                    stroke-width="8"
                    fill="transparent"
                    class="text-gray-200"
                  />
                  <circle
                    v-for="(exemption, index) in exemptionsReport.exemptionTypes"
                    :key="exemption.type"
                    cx="64"
                    cy="64"
                    r="54"
                    stroke="currentColor"
                    stroke-width="8"
                    fill="transparent"
                    :stroke-dasharray="`${exemption.percentage * 3.39} 339`"
                    :stroke-dashoffset="getExemptionOffset(index)"
                    :style="{ color: exemption.color }"
                  />
                </svg>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- گزارش تفصیلی -->
    <div v-else-if="activeReportTab === 'detailed'">
      <div class="space-y-6">
        <!-- فیلترهای پیشرفته -->
        <div class="bg-gray-50 rounded-lg p-6">
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">دسته‌بندی محصول</label>
              <select 
                v-model="detailedFilters.category"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm"
                @change="updateDetailedReport"
              >
                <option value="">همه دسته‌ها</option>
                <option value="electronics">الکترونیک</option>
                <option value="clothing">پوشاک</option>
                <option value="food">مواد غذایی</option>
                <option value="services">خدمات</option>
              </select>
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نوع مشتری</label>
              <select 
                v-model="detailedFilters.customerType"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm"
                @change="updateDetailedReport"
              >
                <option value="">همه انواع</option>
                <option value="individual">شخصی</option>
                <option value="business">شرکتی</option>
                <option value="government">دولتی</option>
              </select>
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">منطقه</label>
              <select 
                v-model="detailedFilters.region"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm"
                @change="updateDetailedReport"
              >
                <option value="">همه مناطق</option>
                <option value="tehran">تهران</option>
                <option value="isfahan">اصفهان</option>
                <option value="mashhad">مشهد</option>
              </select>
            </div>
          </div>
        </div>

        <!-- جدول تفصیلی -->
        <div class="overflow-x-auto">
          <table class="w-full text-sm">
            <thead>
              <tr class="border-b border-gray-200 bg-gray-50">
                <th class="text-right py-3 px-4 font-medium text-gray-600">تاریخ</th>
                <th class="text-right py-3 px-4 font-medium text-gray-600">شماره فاکتور</th>
                <th class="text-right py-3 px-4 font-medium text-gray-600">مشتری</th>
                <th class="text-right py-3 px-4 font-medium text-gray-600">دسته‌بندی</th>
                <th class="text-right py-3 px-4 font-medium text-gray-600">مبلغ پایه</th>
                <th class="text-right py-3 px-4 font-medium text-gray-600">مالیات VAT</th>
                <th class="text-right py-3 px-4 font-medium text-gray-600">مالیات درآمد</th>
                <th class="text-right py-3 px-4 font-medium text-gray-600">معافیت</th>
                <th class="text-right py-3 px-4 font-medium text-gray-600">کل</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="item in detailedReport.items" :key="item.id" class="border-b border-gray-100">
                <td class="py-3 px-4 text-gray-900">{{ formatDate(item.date) }}</td>
                <td class="py-3 px-4 text-gray-900">{{ item.invoiceNumber }}</td>
                <td class="py-3 px-4 text-gray-900">{{ item.customer }}</td>
                <td class="py-3 px-4 text-gray-900">{{ item.category }}</td>
                <td class="py-3 px-4 text-gray-900">{{ formatCurrency(item.baseAmount) }}</td>
                <td class="py-3 px-4 text-blue-600">{{ formatCurrency(item.vatAmount) }}</td>
                <td class="py-3 px-4 text-green-600">{{ formatCurrency(item.incomeTaxAmount) }}</td>
                <td class="py-3 px-4 text-yellow-600">{{ formatCurrency(item.exemptionAmount) }}</td>
                <td class="py-3 px-4 font-medium text-gray-900">{{ formatCurrency(item.totalAmount) }}</td>
              </tr>
            </tbody>
          </table>
        </div>

        <!-- خلاصه تفصیلی -->
        <div class="bg-gray-50 rounded-lg p-6">
          <div class="grid grid-cols-1 md:grid-cols-4 gap-6 text-sm">
            <div class="text-center">
              <div class="text-lg font-bold text-gray-900">{{ formatCurrency(detailedReport.summary.totalBaseAmount) }}</div>
              <div class="text-gray-600">کل مبلغ پایه</div>
            </div>
            <div class="text-center">
              <div class="text-lg font-bold text-blue-600">{{ formatCurrency(detailedReport.summary.totalVatAmount) }}</div>
              <div class="text-gray-600">کل مالیات VAT</div>
            </div>
            <div class="text-center">
              <div class="text-lg font-bold text-green-600">{{ formatCurrency(detailedReport.summary.totalIncomeTaxAmount) }}</div>
              <div class="text-gray-600">کل مالیات درآمد</div>
            </div>
            <div class="text-center">
              <div class="text-lg font-bold text-gray-900">{{ formatCurrency(detailedReport.summary.totalAmount) }}</div>
              <div class="text-gray-600">کل مبلغ</div>
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
  { id: 'vat', label: 'مالیات بر ارزش افزوده' },
  { id: 'income', label: 'مالیات بر درآمد' },
  { id: 'exemptions', label: 'معافیت‌ها' },
  { id: 'detailed', label: 'گزارش تفصیلی' }
]

const activeReportTab = ref('vat')

// فیلترهای گزارش
const reportFilters = ref({
  period: 'monthly',
  dateFrom: '',
  dateTo: '',
  reportType: 'all'
})

// فیلترهای گزارش تفصیلی
const detailedFilters = ref({
  category: '',
  customerType: '',
  region: ''
})

// گزارش مالیات بر ارزش افزوده
const vatReport = ref({
  totalVat: 1250000000,
  totalTransactions: 15420,
  avgVatRate: 9.2,
  totalAmount: 13500000000,
  maxAmount: 150000000,
  monthlyData: [
    { month: 'فروردین', amount: 120000000 },
    { month: 'اردیبهشت', amount: 135000000 },
    { month: 'خرداد', amount: 110000000 },
    { month: 'تیر', amount: 145000000 },
    { month: 'مرداد', amount: 130000000 },
    { month: 'شهریور', amount: 140000000 }
  ],
  categories: [
    { name: 'الکترونیک', transactions: 5200, baseAmount: 4500000000, vatAmount: 405000000, avgRate: 9 },
    { name: 'پوشاک', transactions: 3800, baseAmount: 2800000000, vatAmount: 252000000, avgRate: 9 },
    { name: 'مواد غذایی', transactions: 4200, baseAmount: 3200000000, vatAmount: 288000000, avgRate: 9 },
    { name: 'خدمات', transactions: 2200, baseAmount: 3000000000, vatAmount: 270000000, avgRate: 9 }
  ]
})

// گزارش مالیات بر درآمد
const incomeReport = ref({
  totalIncomeTax: 850000000,
  totalTaxpayers: 12500,
  avgTaxPerPerson: 68000,
  taxBrackets: [
    { range: 'کمتر از ۵۰ میلیون', count: 8500, totalTax: 425000000, percentage: 50 },
    { range: '۵۰ تا ۱۰۰ میلیون', count: 3200, totalTax: 288000000, percentage: 34 },
    { range: 'بیش از ۱۰۰ میلیون', count: 800, totalTax: 137000000, percentage: 16 }
  ]
})

// گزارش معافیت‌ها
const exemptionsReport = ref({
  totalExemptedAmount: 320000000,
  totalExemptions: 8500,
  exemptionRate: 23.7,
  exemptionTypes: [
    { type: 'food', name: 'مواد غذایی', amount: 180000000, percentage: 56.3, color: '#10B981' },
    { type: 'medical', name: 'پزشکی', amount: 80000000, percentage: 25, color: '#3B82F6' },
    { type: 'education', name: 'آموزشی', amount: 40000000, percentage: 12.5, color: '#F59E0B' },
    { type: 'digital', name: 'دیجیتال', amount: 20000000, percentage: 6.2, color: '#8B5CF6' }
  ]
})

// گزارش تفصیلی
const detailedReport = ref({
  items: [
    {
      id: 1,
      date: '2024-01-15',
      invoiceNumber: 'INV-001',
      customer: 'علی احمدی',
      category: 'الکترونیک',
      baseAmount: 1000000,
      vatAmount: 90000,
      incomeTaxAmount: 150000,
      exemptionAmount: 0,
      totalAmount: 1240000
    },
    {
      id: 2,
      date: '2024-01-14',
      invoiceNumber: 'INV-002',
      customer: 'شرکت فناوری',
      category: 'خدمات',
      baseAmount: 2000000,
      vatAmount: 180000,
      incomeTaxAmount: 300000,
      exemptionAmount: 50000,
      totalAmount: 2430000
    }
  ],
  summary: {
    totalBaseAmount: 3000000,
    totalVatAmount: 270000,
    totalIncomeTaxAmount: 450000,
    totalAmount: 3670000
  }
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

// فرمت تاریخ
const formatDate = (dateString: string): string => {
  return new Date(dateString).toLocaleDateString('fa-IR')
}

// محاسبه offset برای نمودار دایره‌ای
const getExemptionOffset = (index: number): number => {
  let offset = 0
  for (let i = 0; i < index; i++) {
    offset += exemptionsReport.value.exemptionTypes[i].percentage * 3.39
  }
  return offset
}

// بروزرسانی گزارش‌ها
const updateReports = () => {
  // TODO: بروزرسانی گزارش‌ها بر اساس فیلترها
  console.log('گزارش‌ها بروزرسانی شد:', reportFilters.value)
}

// بروزرسانی گزارش تفصیلی
const updateDetailedReport = () => {
  // TODO: بروزرسانی گزارش تفصیلی بر اساس فیلترها
  console.log('گزارش تفصیلی بروزرسانی شد:', detailedFilters.value)
}

// تولید تمام گزارش‌ها
const generateAllReports = async () => {
  try {
    // TODO: تولید تمام گزارش‌ها
    console.log('تمام گزارش‌ها تولید شدند')
  } catch (error) {
    console.error('خطا در تولید گزارش‌ها:', error)
  }
}

// خروجی گزارش‌ها
const exportReports = () => {
  // TODO: خروجی Excel
  console.log('گزارش‌ها به Excel صادر شد')
}

// بارگذاری اولیه
onMounted(() => {
  // تنظیم تاریخ‌های پیش‌فرض
  const now = new Date()
  const firstDay = new Date(now.getFullYear(), now.getMonth(), 1)
  const lastDay = new Date(now.getFullYear(), now.getMonth() + 1, 0)
  
  reportFilters.value.dateFrom = firstDay.toISOString().split('T')[0]
  reportFilters.value.dateTo = lastDay.toISOString().split('T')[0]
  
  // TODO: بارگذاری گزارش‌ها از API
})
</script>

<!--
  کامپوننت گزارش‌های مالیاتی
  شامل:
  1. گزارش مالیات بر ارزش افزوده
  2. گزارش مالیات بر درآمد
  3. گزارش معافیت‌ها
  4. گزارش تفصیلی
  5. فیلترهای پیشرفته
  6. نمودارها و آمار
  7. طراحی مدرن و کاملاً ریسپانسیو
--> 
