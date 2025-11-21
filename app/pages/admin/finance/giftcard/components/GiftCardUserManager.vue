<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">مدیریت استفاده‌کنندگان گیفت کارت</h3>
        <p class="text-sm text-gray-600">ثبت، پیگیری و مدیریت استفاده‌کنندگان کارت‌ها</p>
      </div>
      <div class="flex gap-2">
        <button
          class="inline-flex items-center px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2"
          @click="showNewUserModal = true"
        >
          <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
          </svg>
          ثبت استفاده‌کننده جدید
        </button>
        <button
          class="inline-flex items-center px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
          @click="showBulkImportModal = true"
        >
          <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M9 19l3 3m0 0l3-3m-3 3V10"></path>
          </svg>
          ورود گروهی
        </button>
        <button
          class="inline-flex items-center px-4 py-2 bg-purple-600 text-white text-sm font-medium rounded-lg hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-purple-500 focus:ring-offset-2"
          @click="exportUsers"
        >
          <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
          </svg>
          خروجی گزارش
        </button>
      </div>
    </div>

    <!-- آمار استفاده‌کنندگان -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="bg-white p-6 rounded-lg border border-gray-200">
        <div class="flex items-center">
          <div class="p-2 bg-blue-100 rounded-lg">
            <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path>
            </svg>
          </div>
          <div class="ml-3">
            <p class="text-sm font-medium text-gray-600">کل استفاده‌کنندگان</p>
            <p class="text-2xl font-bold text-gray-900">{{ totalUsers }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white p-6 rounded-lg border border-gray-200">
        <div class="flex items-center">
          <div class="p-2 bg-green-100 rounded-lg">
            <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
          <div class="ml-3">
            <p class="text-sm font-medium text-gray-600">کارت‌های فعال</p>
            <p class="text-2xl font-bold text-gray-900">{{ activeCards }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white p-6 rounded-lg border border-gray-200">
        <div class="flex items-center">
          <div class="p-2 bg-yellow-100 rounded-lg">
            <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
            </svg>
          </div>
          <div class="ml-3">
            <p class="text-sm font-medium text-gray-600">اعتبار باقی‌مانده</p>
            <p class="text-2xl font-bold text-gray-900">{{ formatCurrency(totalRemainingBalance) }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white p-6 rounded-lg border border-gray-200">
        <div class="flex items-center">
          <div class="p-2 bg-purple-100 rounded-lg">
            <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
            </svg>
          </div>
          <div class="ml-3">
            <p class="text-sm font-medium text-gray-600">استفاده امروز</p>
            <p class="text-2xl font-bold text-gray-900">{{ todayUsage }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- تب‌های مدیریت -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="border-b border-gray-200">
        <nav class="-mb-px flex space-x-8 space-x-reverse">
          <button
            :class="{
              'border-blue-500 text-blue-600': activeTab === 'users',
              'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300': activeTab !== 'users'
            }"
            class="whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm"
            @click="activeTab = 'users'"
          >
            استفاده‌کنندگان
          </button>
          <button
            :class="{
              'border-blue-500 text-blue-600': activeTab === 'usage_history',
              'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300': activeTab !== 'usage_history'
            }"
            class="whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm"
            @click="activeTab = 'usage_history'"
          >
            تاریخچه استفاده
          </button>
          <button
            :class="{
              'border-blue-500 text-blue-600': activeTab === 'balances',
              'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300': activeTab !== 'balances'
            }"
            class="whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm"
            @click="activeTab = 'balances'"
          >
            اعتبار باقی‌مانده
          </button>
          <button
            :class="{
              'border-blue-500 text-blue-600': activeTab === 'restrictions',
              'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300': activeTab !== 'restrictions'
            }"
            class="whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm"
            @click="activeTab = 'restrictions'"
          >
            محدودیت‌ها
          </button>
        </nav>
      </div>

      <!-- محتوای تب‌ها -->
      <div class="p-6">
        <!-- تب استفاده‌کنندگان -->
        <div v-if="activeTab === 'users'">
          <div class="flex justify-between items-center mb-4">
            <h4 class="text-lg font-semibold text-gray-900">لیست استفاده‌کنندگان</h4>
            <div class="flex gap-2">
              <input
                v-model="userSearch"
                type="text"
                placeholder="جستجو در استفاده‌کنندگان..."
                class="px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              />
              <select
                v-model="userStatusFilter"
                class="px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
                <option value="">همه وضعیت‌ها</option>
                <option value="active">فعال</option>
                <option value="inactive">غیرفعال</option>
                <option value="suspended">معلق</option>
                <option value="blocked">مسدود</option>
              </select>
            </div>
          </div>

          <div class="overflow-x-auto">
            <table class="min-w-full divide-y divide-gray-200">
              <thead class="bg-gray-50">
                <tr>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                    استفاده‌کننده
                  </th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                    کارت‌های فعال
                  </th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                    اعتبار کل
                  </th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                    آخرین استفاده
                  </th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                    وضعیت
                  </th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                    عملیات
                  </th>
                </tr>
              </thead>
              <tbody class="bg-white divide-y divide-gray-200">
                <tr
                  v-for="user in filteredUsers"
                  :key="user.id"
                  class="hover:bg-gray-50"
                >
                  <td class="px-6 py-4 whitespace-nowrap">
                    <div class="flex items-center">
                      <div class="">
                        <div class="h-10 w-10 rounded-full bg-gradient-to-r from-green-500 to-blue-600 flex items-center justify-center">
                          <span class="text-white font-medium">{{ user.name.charAt(0) }}</span>
                        </div>
                      </div>
                      <div class="mr-3">
                        <div class="text-sm font-medium text-gray-900">{{ user.name }}</div>
                        <div class="text-sm text-gray-500">{{ user.email }}</div>
                      </div>
                    </div>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap">
                    <div class="text-sm text-gray-900">{{ user.activeCards }} کارت</div>
                    <div class="text-sm text-gray-500">{{ user.totalCards }} کل</div>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap">
                    <div class="text-sm text-gray-900">{{ formatCurrency(user.totalBalance) }}</div>
                    <div class="text-sm text-gray-500">{{ formatCurrency(user.remainingBalance) }} باقی</div>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                    {{ formatDate(user.lastUsageDate) }}
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap">
                    <span
                      :class="{
                        'bg-green-100 text-green-800': user.status === 'active',
                        'bg-red-100 text-red-800': user.status === 'inactive',
                        'bg-yellow-100 text-yellow-800': user.status === 'suspended',
                        'bg-gray-100 text-gray-800': user.status === 'blocked'
                      }"
                      class="px-2 py-1 text-xs font-medium rounded-full"
                    >
                      {{ getStatusText(user.status) }}
                    </span>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                    <div class="flex gap-2">
                      <button
                        class="text-blue-600 hover:text-blue-900"
                        @click="viewUserDetails(user)"
                      >
                        جزئیات
                      </button>
                      <button
                        class="text-green-600 hover:text-green-900"
                        @click="editUser(user)"
                      >
                        ویرایش
                      </button>
                      <button
                        :class="{
                          'text-red-600 hover:text-red-900': user.status === 'active',
                          'text-green-600 hover:text-green-900': user.status === 'inactive'
                        }"
                        @click="toggleUserStatus(user)"
                      >
                        {{ user.status === 'active' ? 'غیرفعال' : 'فعال' }}
                      </button>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <!-- تب تاریخچه استفاده -->
        <div v-if="activeTab === 'usage_history'">
          <div class="flex justify-between items-center mb-4">
            <h4 class="text-lg font-semibold text-gray-900">تاریخچه استفاده از کارت‌ها</h4>
            <div class="flex gap-2">
              <input
                v-model="usageSearch"
                type="text"
                placeholder="جستجو در تاریخچه..."
                class="px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              />
              <select
                v-model="usageTypeFilter"
                class="px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
                <option value="">همه انواع</option>
                <option value="purchase">خرید</option>
                <option value="usage">استفاده</option>
                <option value="refund">بازپرداخت</option>
                <option value="expiry">انقضا</option>
              </select>
            </div>
          </div>

          <div class="space-y-4">
            <div
              v-for="usage in filteredUsageHistory"
              :key="usage.id"
              class="bg-white border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow"
            >
              <div class="flex items-center justify-between mb-3">
                <div class="flex items-center">
                  <div
class="p-2 rounded-lg mr-3"
                    :class="{
                      'bg-green-100': usage.type === 'usage',
                      'bg-blue-100': usage.type === 'purchase',
                      'bg-yellow-100': usage.type === 'refund',
                      'bg-red-100': usage.type === 'expiry'
                    }"
                  >
                    <svg
class="w-5 h-5"
                      :class="{
                        'text-green-600': usage.type === 'usage',
                        'text-blue-600': usage.type === 'purchase',
                        'text-yellow-600': usage.type === 'refund',
                        'text-red-600': usage.type === 'expiry'
                      }"
                      fill="none" stroke="currentColor" viewBox="0 0 24 24"
                    >
                      <path v-if="usage.type === 'usage'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
                      <path v-if="usage.type === 'purchase'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4m0 0L7 13m0 0l-2.5 5M7 13l2.5 5m6-5v6a2 2 0 01-2 2H9a2 2 0 01-2-2v-6m6 0V9a2 2 0 00-2-2H9a2 2 0 00-2 2v4.01"></path>
                      <path v-if="usage.type === 'refund'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h10a8 8 0 018 8v2M3 10l6 6m-6-6l6-6"></path>
                      <path v-if="usage.type === 'expiry'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                    </svg>
                  </div>
                  <div>
                    <h5 class="text-lg font-semibold text-gray-900">{{ usage.userName }}</h5>
                    <p class="text-sm text-gray-600">{{ usage.cardCode }}</p>
                  </div>
                </div>
                <span class="text-sm text-gray-500">{{ formatDate(usage.date) }}</span>
              </div>
              
              <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-3">
                <div>
                  <span class="text-sm text-gray-600">مبلغ:</span>
                  <span class="text-sm font-medium text-gray-900 mr-2">{{ formatCurrency(usage.amount) }}</span>
                </div>
                <div>
                  <span class="text-sm text-gray-600">نوع:</span>
                  <span class="text-sm font-medium text-gray-900 mr-2">{{ getUsageTypeText(usage.type) }}</span>
                </div>
                <div>
                  <span class="text-sm text-gray-600">وضعیت:</span>
                  <span
                    :class="{
                      'bg-green-100 text-green-800': usage.status === 'success',
                      'bg-red-100 text-red-800': usage.status === 'failed',
                      'bg-yellow-100 text-yellow-800': usage.status === 'pending'
                    }"
                    class="px-2 py-1 text-xs font-medium rounded-full"
                  >
                    {{ getUsageStatusText(usage.status) }}
                  </span>
                </div>
              </div>

              <div class="flex gap-2">
                <button
                  class="flex-1 px-3 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
                  @click="viewUsageDetails(usage)"
                >
                  جزئیات
                </button>
                <button
                  v-if="usage.type === 'usage' && usage.status === 'failed'"
                  class="flex-1 px-3 py-2 bg-green-600 text-white text-sm font-medium rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2"
                  @click="retryUsage(usage)"
                >
                  تلاش مجدد
                </button>
              </div>
            </div>
          </div>
        </div>

        <!-- تب اعتبار باقی‌مانده -->
        <div v-if="activeTab === 'balances'">
          <div class="flex justify-between items-center mb-4">
            <h4 class="text-lg font-semibold text-gray-900">اعتبار باقی‌مانده کارت‌ها</h4>
            <div class="flex gap-2">
              <input
                v-model="balanceSearch"
                type="text"
                placeholder="جستجو در کارت‌ها..."
                class="px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              />
              <select
                v-model="balanceStatusFilter"
                class="px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
                <option value="">همه وضعیت‌ها</option>
                <option value="active">فعال</option>
                <option value="expired">منقضی شده</option>
                <option value="used">استفاده شده</option>
                <option value="blocked">مسدود</option>
              </select>
            </div>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            <div
              v-for="card in filteredBalanceCards"
              :key="card.id"
              class="bg-white border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow"
            >
              <div class="flex items-center justify-between mb-3">
                <h5 class="text-lg font-semibold text-gray-900">{{ card.code }}</h5>
                <span
                  :class="{
                    'bg-green-100 text-green-800': card.status === 'active',
                    'bg-red-100 text-red-800': card.status === 'expired',
                    'bg-yellow-100 text-yellow-800': card.status === 'used',
                    'bg-gray-100 text-gray-800': card.status === 'blocked'
                  }"
                  class="px-2 py-1 text-xs font-medium rounded-full"
                >
                  {{ getStatusText(card.status) }}
                </span>
              </div>
              
              <div class="space-y-2 mb-4">
                <div class="flex justify-between">
                  <span class="text-sm text-gray-600">مبلغ اولیه:</span>
                  <span class="text-sm font-medium text-gray-900">{{ formatCurrency(card.initialAmount) }}</span>
                </div>
                <div class="flex justify-between">
                  <span class="text-sm text-gray-600">مبلغ استفاده شده:</span>
                  <span class="text-sm font-medium text-gray-900">{{ formatCurrency(card.usedAmount) }}</span>
                </div>
                <div class="flex justify-between border-t border-gray-200 pt-2">
                  <span class="text-sm font-medium text-gray-600">باقی‌مانده:</span>
                  <span class="text-sm font-bold text-gray-900">{{ formatCurrency(card.remainingAmount) }}</span>
                </div>
                <div class="flex justify-between">
                  <span class="text-sm text-gray-600">تاریخ انقضا:</span>
                  <span class="text-sm font-medium text-gray-900">{{ formatDate(card.expiryDate) }}</span>
                </div>
              </div>

              <div class="flex gap-2">
                <button
                  class="flex-1 px-3 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
                  @click="viewCardDetails(card)"
                >
                  جزئیات
                </button>
                <button
                  v-if="card.status === 'active'"
                  class="flex-1 px-3 py-2 bg-green-600 text-white text-sm font-medium rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2"
                  @click="extendCard(card)"
                >
                  تمدید
                </button>
              </div>
            </div>
          </div>
        </div>

        <!-- تب محدودیت‌ها -->
        <div v-if="activeTab === 'restrictions'">
          <div class="flex justify-between items-center mb-4">
            <h4 class="text-lg font-semibold text-gray-900">مدیریت محدودیت‌های استفاده</h4>
            <button
              class="px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2"
              @click="showNewRestrictionModal = true"
            >
              افزودن محدودیت جدید
            </button>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            <div
              v-for="restriction in restrictions"
              :key="restriction.id"
              class="bg-white border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow"
            >
              <div class="flex items-center justify-between mb-3">
                <h5 class="text-lg font-semibold text-gray-900">{{ restriction.name }}</h5>
                <span
                  :class="{
                    'bg-green-100 text-green-800': restriction.status === 'active',
                    'bg-red-100 text-red-800': restriction.status === 'inactive'
                  }"
                  class="px-2 py-1 text-xs font-medium rounded-full"
                >
                  {{ getStatusText(restriction.status) }}
                </span>
              </div>
              
              <div class="space-y-2 mb-4">
                <div class="flex justify-between">
                  <span class="text-sm text-gray-600">نوع:</span>
                  <span class="text-sm font-medium text-gray-900">{{ getRestrictionTypeText(restriction.type) }}</span>
                </div>
                <div class="flex justify-between">
                  <span class="text-sm text-gray-600">محدودیت:</span>
                  <span class="text-sm font-medium text-gray-900">{{ restriction.limit }}</span>
                </div>
                <div class="flex justify-between">
                  <span class="text-sm text-gray-600">دوره:</span>
                  <span class="text-sm font-medium text-gray-900">{{ getRestrictionPeriodText(restriction.period) }}</span>
                </div>
                <div class="flex justify-between">
                  <span class="text-sm text-gray-600">اعمال شده:</span>
                  <span class="text-sm font-medium text-gray-900">{{ restriction.appliedCount }} بار</span>
                </div>
              </div>

              <div class="flex gap-2">
                <button
                  class="flex-1 px-3 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
                  @click="editRestriction(restriction)"
                >
                  ویرایش
                </button>
                <button
                  :class="{
                    'flex-1 px-3 py-2 text-sm font-medium rounded-lg focus:outline-none focus:ring-2 focus:ring-offset-2': true,
                    'bg-red-600 text-white hover:bg-red-700 focus:ring-red-500': restriction.status === 'active',
                    'bg-green-600 text-white hover:bg-green-700 focus:ring-green-500': restriction.status === 'inactive'
                  }"
                  @click="toggleRestrictionStatus(restriction)"
                >
                  {{ restriction.status === 'active' ? 'غیرفعال' : 'فعال' }}
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- مودال‌ها -->
    <GiftCardUserCreateModal
      v-if="showNewUserModal"
      @close="showNewUserModal = false"
      @created="handleUserCreated"
    />

    <GiftCardRestrictionCreateModal
      v-if="showNewRestrictionModal"
      @close="showNewRestrictionModal = false"
      @created="handleRestrictionCreated"
    />
  </div>
</template>

<script setup lang="ts">
// تعریف interface ها (Rule 67 - Interface Organization and Reusability)
interface User {
  status?: string
  [key: string]: unknown
}

interface Usage {
  [key: string]: unknown
}

interface Card {
  [key: string]: unknown
}

interface Restriction {
  status?: string
  [key: string]: unknown
}

// تعریف متغیرهای reactive
const activeTab = ref('users')
const showNewUserModal = ref(false)
const showNewRestrictionModal = ref(false)
const showBulkImportModal = ref(false)

const userSearch = ref('')
const userStatusFilter = ref('')
const usageSearch = ref('')
const usageTypeFilter = ref('')
const balanceSearch = ref('')
const balanceStatusFilter = ref('')

// داده‌های نمونه استفاده‌کنندگان
const users = ref([
  {
    id: 1,
    name: 'فاطمه احمدی',
    email: 'fateme@example.com',
    activeCards: 3,
    totalCards: 5,
    totalBalance: 2500000,
    remainingBalance: 1800000,
    lastUsageDate: new Date(Date.now() - 86400000),
    status: 'active'
  },
  {
    id: 2,
    name: 'محمد رضایی',
    email: 'mohammad@example.com',
    activeCards: 1,
    totalCards: 2,
    totalBalance: 800000,
    remainingBalance: 0,
    lastUsageDate: new Date(Date.now() - 172800000),
    status: 'inactive'
  },
  {
    id: 3,
    name: 'سارا کریمی',
    email: 'sara@example.com',
    activeCards: 2,
    totalCards: 4,
    totalBalance: 1500000,
    remainingBalance: 1200000,
    lastUsageDate: new Date(Date.now() - 259200000),
    status: 'suspended'
  }
])

// داده‌های نمونه تاریخچه استفاده
const usageHistory = ref([
  {
    id: 1,
    userName: 'فاطمه احمدی',
    cardCode: 'GC-2024-001',
    type: 'usage',
    amount: 250000,
    status: 'success',
    date: new Date(Date.now() - 86400000)
  },
  {
    id: 2,
    userName: 'محمد رضایی',
    cardCode: 'GC-2024-002',
    type: 'purchase',
    amount: 500000,
    status: 'success',
    date: new Date(Date.now() - 172800000)
  },
  {
    id: 3,
    userName: 'سارا کریمی',
    cardCode: 'GC-2024-003',
    type: 'usage',
    amount: 300000,
    status: 'failed',
    date: new Date(Date.now() - 259200000)
  }
])

// داده‌های نمونه کارت‌های اعتباری
const balanceCards = ref([
  {
    id: 1,
    code: 'GC-2024-001',
    initialAmount: 500000,
    usedAmount: 250000,
    remainingAmount: 250000,
    expiryDate: new Date(Date.now() + 2592000000),
    status: 'active'
  },
  {
    id: 2,
    code: 'GC-2024-002',
    initialAmount: 800000,
    usedAmount: 800000,
    remainingAmount: 0,
    expiryDate: new Date(Date.now() + 5184000000),
    status: 'used'
  },
  {
    id: 3,
    code: 'GC-2024-003',
    initialAmount: 300000,
    usedAmount: 100000,
    remainingAmount: 200000,
    expiryDate: new Date(Date.now() - 86400000),
    status: 'expired'
  }
])

// داده‌های نمونه محدودیت‌ها
const restrictions = ref([
  {
    id: 1,
    name: 'محدودیت روزانه',
    type: 'daily',
    limit: 500000,
    period: 'day',
    appliedCount: 15,
    status: 'active'
  },
  {
    id: 2,
    name: 'محدودیت ماهانه',
    type: 'monthly',
    limit: 5000000,
    period: 'month',
    appliedCount: 3,
    status: 'active'
  },
  {
    id: 3,
    name: 'محدودیت تعداد تراکنش',
    type: 'transaction_count',
    limit: 10,
    period: 'day',
    appliedCount: 8,
    status: 'inactive'
  }
])

// محاسبات
const totalUsers = computed(() => users.value.length)
const activeCards = computed(() => users.value.reduce((sum, user) => sum + user.activeCards, 0))
const totalRemainingBalance = computed(() => users.value.reduce((sum, user) => sum + user.remainingBalance, 0))
const todayUsage = computed(() => {
  const today = new Date().toDateString()
  return usageHistory.value.filter(usage => 
    new Date(usage.date).toDateString() === today
  ).length
})

// فیلترها
const filteredUsers = computed(() => {
  return users.value.filter(user => {
    const matchesSearch = 
      user.name.toLowerCase().includes(userSearch.value.toLowerCase()) ||
      user.email.toLowerCase().includes(userSearch.value.toLowerCase())
    
    const matchesStatus = !userStatusFilter.value || user.status === userStatusFilter.value
    
    return matchesSearch && matchesStatus
  })
})

const filteredUsageHistory = computed(() => {
  return usageHistory.value.filter(usage => {
    const matchesSearch = 
      usage.userName.toLowerCase().includes(usageSearch.value.toLowerCase()) ||
      usage.cardCode.toLowerCase().includes(usageSearch.value.toLowerCase())
    
    const matchesType = !usageTypeFilter.value || usage.type === usageTypeFilter.value
    
    return matchesSearch && matchesType
  })
})

const filteredBalanceCards = computed(() => {
  return balanceCards.value.filter(card => {
    const matchesSearch = 
      card.code.toLowerCase().includes(balanceSearch.value.toLowerCase())
    
    const matchesStatus = !balanceStatusFilter.value || card.status === balanceStatusFilter.value
    
    return matchesSearch && matchesStatus
  })
})

// توابع کمکی
const getStatusText = (status) => {
  const statusMap = {
    active: 'فعال',
    inactive: 'غیرفعال',
    suspended: 'معلق',
    blocked: 'مسدود',
    expired: 'منقضی شده',
    used: 'استفاده شده'
  }
  return statusMap[status] || status
}

const getUsageTypeText = (type) => {
  const typeMap = {
    usage: 'استفاده',
    purchase: 'خرید',
    refund: 'بازپرداخت',
    expiry: 'انقضا'
  }
  return typeMap[type] || type
}

const getUsageStatusText = (status) => {
  const statusMap = {
    success: 'موفق',
    failed: 'ناموفق',
    pending: 'در انتظار'
  }
  return statusMap[status] || status
}

const getRestrictionTypeText = (type) => {
  const typeMap = {
    daily: 'روزانه',
    monthly: 'ماهانه',
    transaction_count: 'تعداد تراکنش',
    amount_limit: 'محدودیت مبلغ'
  }
  return typeMap[type] || type
}

const getRestrictionPeriodText = (period) => {
  const periodMap = {
    day: 'روزانه',
    week: 'هفتگی',
    month: 'ماهانه',
    year: 'سالانه'
  }
  return periodMap[period] || period
}

const formatCurrency = (amount) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

const formatDate = (date) => {
  return new Intl.DateTimeFormat('fa-IR', {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  }).format(date)
}

// توابع عملیات
const viewUserDetails = (_user: User) => {

  // اینجا کد نمایش جزئیات کاربر را اضافه کنید
}

const editUser = (_user: User) => {

  // اینجا کد ویرایش کاربر را اضافه کنید
}

const toggleUserStatus = (user: User) => {
  user.status = user.status === 'active' ? 'inactive' : 'active'
}

const viewUsageDetails = (_usage: Usage) => {

  // اینجا کد نمایش جزئیات استفاده را اضافه کنید
}

const retryUsage = (_usage: Usage) => {

  alert('تلاش مجدد برای استفاده از کارت انجام شد')
}

const viewCardDetails = (_card: Card) => {

  // اینجا کد نمایش جزئیات کارت را اضافه کنید
}

const extendCard = (_card: Card) => {

  alert('کارت با موفقیت تمدید شد')
}

const editRestriction = (_restriction: Restriction) => {

  // اینجا کد ویرایش محدودیت را اضافه کنید
}

const toggleRestrictionStatus = (restriction) => {
  restriction.status = restriction.status === 'active' ? 'inactive' : 'active'
}

const exportUsers = () => {

  alert('گزارش استفاده‌کنندگان در حال دانلود است')
}

// توابع مدیریت مودال‌ها
const handleUserCreated = (newUser) => {
  users.value.unshift({
    ...newUser,
    id: Date.now(),
    activeCards: 0,
    totalCards: 0,
    totalBalance: 0,
    remainingBalance: 0,
    lastUsageDate: null
  })
  showNewUserModal.value = false
}

const handleRestrictionCreated = (newRestriction) => {
  restrictions.value.unshift({
    ...newRestriction,
    id: Date.now(),
    appliedCount: 0
  })
  showNewRestrictionModal.value = false
}
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 
