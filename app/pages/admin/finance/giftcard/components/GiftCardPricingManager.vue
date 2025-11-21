<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">مدیریت قیمت‌گذاری گیفت کارت</h3>
        <p class="text-sm text-gray-600">تعیین قیمت، تخفیف‌ها و کارمزدهای گیفت کارت</p>
      </div>
      <div class="flex gap-2">
        <button
          class="inline-flex items-center px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2"
          @click="showNewPricingModal = true"
        >
          <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
          </svg>
          تعریف قیمت جدید
        </button>
        <button
          class="inline-flex items-center px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
          @click="showDiscountModal = true"
        >
          <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 7h6m0 10v-3m-3 3h.01M9 17h.01M9 14h.01M12 14h.01M15 11h.01M12 11h.01M9 11h.01M7 21h10a2 2 0 002-2V5a2 2 0 00-2-2H7a2 2 0 00-2 2v14a2 2 0 002 2z"></path>
          </svg>
          تخفیف ویژه
        </button>
        <button
          class="inline-flex items-center px-4 py-2 bg-purple-600 text-white text-sm font-medium rounded-lg hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-purple-500 focus:ring-offset-2"
          @click="exportPricing"
        >
          <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
          </svg>
          خروجی گزارش
        </button>
      </div>
    </div>

    <!-- آمار قیمت‌گذاری -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
      <div class="bg-white p-6 rounded-lg border border-gray-200">
        <div class="flex items-center">
          <div class="p-2 bg-green-100 rounded-lg">
            <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
            </svg>
          </div>
          <div class="ml-3">
            <p class="text-sm font-medium text-gray-600">قیمت‌های فعال</p>
            <p class="text-2xl font-bold text-gray-900">{{ activePricing.length }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white p-6 rounded-lg border border-gray-200">
        <div class="flex items-center">
          <div class="p-2 bg-blue-100 rounded-lg">
            <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 7h6m0 10v-3m-3 3h.01M9 17h.01M9 14h.01M12 14h.01M15 11h.01M12 11h.01M9 11h.01M7 21h10a2 2 0 002-2V5a2 2 0 00-2-2H7a2 2 0 00-2 2v14a2 2 0 002 2z"></path>
            </svg>
          </div>
          <div class="ml-3">
            <p class="text-sm font-medium text-gray-600">تخفیف‌های فعال</p>
            <p class="text-2xl font-bold text-gray-900">{{ activeDiscounts.length }}</p>
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
            <p class="text-sm font-medium text-gray-600">میانگین کارمزد</p>
            <p class="text-2xl font-bold text-gray-900">{{ averageCommission }}%</p>
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
            <p class="text-sm font-medium text-gray-600">درآمد کارمزد</p>
            <p class="text-2xl font-bold text-gray-900">{{ formatCurrency(totalCommissionRevenue) }}</p>
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
              'border-blue-500 text-blue-600': activeTab === 'pricing',
              'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300': activeTab !== 'pricing'
            }"
            class="whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm"
            @click="activeTab = 'pricing'"
          >
            قیمت‌گذاری
          </button>
          <button
            :class="{
              'border-blue-500 text-blue-600': activeTab === 'discounts',
              'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300': activeTab !== 'discounts'
            }"
            class="whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm"
            @click="activeTab = 'discounts'"
          >
            تخفیف‌ها
          </button>
          <button
            :class="{
              'border-blue-500 text-blue-600': activeTab === 'commissions',
              'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300': activeTab !== 'commissions'
            }"
            class="whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm"
            @click="activeTab = 'commissions'"
          >
            کارمزدها
          </button>
          <button
            :class="{
              'border-blue-500 text-blue-600': activeTab === 'occasions',
              'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300': activeTab !== 'occasions'
            }"
            class="whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm"
            @click="activeTab = 'occasions'"
          >
            مناسبت‌ها
          </button>
        </nav>
      </div>

      <!-- محتوای تب‌ها -->
      <div class="p-6">
        <!-- تب قیمت‌گذاری -->
        <div v-if="activeTab === 'pricing'">
          <div class="flex justify-between items-center mb-4">
            <h4 class="text-lg font-semibold text-gray-900">قیمت‌گذاری گیفت کارت‌ها</h4>
            <div class="flex gap-2">
              <input
                v-model="pricingSearch"
                type="text"
                placeholder="جستجو در قیمت‌ها..."
                class="px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              />
              <select
                v-model="pricingTypeFilter"
                class="px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
                <option value="">همه انواع</option>
                <option value="fixed">مبلغ ثابت</option>
                <option value="variable">مبلغ متغیر</option>
                <option value="themed">موضوعی</option>
                <option value="corporate">شرکتی</option>
              </select>
            </div>
          </div>

          <div class="overflow-x-auto">
            <table class="min-w-full divide-y divide-gray-200">
              <thead class="bg-gray-50">
                <tr>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                    نوع کارت
                  </th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                    مبلغ پایه
                  </th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                    کارمزد
                  </th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                    قیمت نهایی
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
                  v-for="pricing in filteredPricing"
                  :key="pricing.id"
                  class="hover:bg-gray-50"
                >
                  <td class="px-6 py-4 whitespace-nowrap">
                    <div class="flex items-center">
                      <div class="">
                        <div class="h-10 w-10 rounded-lg bg-gradient-to-r from-purple-500 to-pink-600 flex items-center justify-center">
                          <span class="text-white font-medium">{{ pricing.type.charAt(0) }}</span>
                        </div>
                      </div>
                      <div class="mr-3">
                        <div class="text-sm font-medium text-gray-900">{{ pricing.name }}</div>
                        <div class="text-sm text-gray-500">{{ getTypeText(pricing.type) }}</div>
                      </div>
                    </div>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                    {{ formatCurrency(pricing.baseAmount) }}
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                    {{ pricing.commission }}%
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                    {{ formatCurrency(pricing.finalPrice) }}
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap">
                    <span
                      :class="{
                        'bg-green-100 text-green-800': pricing.status === 'active',
                        'bg-red-100 text-red-800': pricing.status === 'inactive',
                        'bg-yellow-100 text-yellow-800': pricing.status === 'pending'
                      }"
                      class="px-2 py-1 text-xs font-medium rounded-full"
                    >
                      {{ getStatusText(pricing.status) }}
                    </span>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                    <div class="flex gap-2">
                      <button
                        class="text-blue-600 hover:text-blue-900"
                        @click="editPricing(pricing)"
                      >
                        ویرایش
                      </button>
                      <button
                        :class="{
                          'text-red-600 hover:text-red-900': pricing.status === 'active',
                          'text-green-600 hover:text-green-900': pricing.status === 'inactive'
                        }"
                        @click="togglePricingStatus(pricing)"
                      >
                        {{ pricing.status === 'active' ? 'غیرفعال' : 'فعال' }}
                      </button>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <!-- تب تخفیف‌ها -->
        <div v-if="activeTab === 'discounts'">
          <div class="flex justify-between items-center mb-4">
            <h4 class="text-lg font-semibold text-gray-900">تخفیف‌های ویژه</h4>
            <div class="flex gap-2">
              <input
                v-model="discountSearch"
                type="text"
                placeholder="جستجو در تخفیف‌ها..."
                class="px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              />
              <select
                v-model="discountTypeFilter"
                class="px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
                <option value="">همه انواع</option>
                <option value="percentage">درصدی</option>
                <option value="fixed">مبلغ ثابت</option>
                <option value="buy_one_get_one">یک به یک</option>
              </select>
            </div>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            <div
              v-for="discount in filteredDiscounts"
              :key="discount.id"
              class="bg-white border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow"
            >
              <div class="flex items-center justify-between mb-3">
                <h5 class="text-lg font-semibold text-gray-900">{{ discount.name }}</h5>
                <span
                  :class="{
                    'bg-green-100 text-green-800': discount.status === 'active',
                    'bg-red-100 text-red-800': discount.status === 'inactive',
                    'bg-yellow-100 text-yellow-800': discount.status === 'scheduled'
                  }"
                  class="px-2 py-1 text-xs font-medium rounded-full"
                >
                  {{ getStatusText(discount.status) }}
                </span>
              </div>
              
              <div class="space-y-2 mb-4">
                <div class="flex justify-between">
                  <span class="text-sm text-gray-600">نوع تخفیف:</span>
                  <span class="text-sm font-medium text-gray-900">{{ getDiscountTypeText(discount.type) }}</span>
                </div>
                <div class="flex justify-between">
                  <span class="text-sm text-gray-600">مقدار:</span>
                  <span class="text-sm font-medium text-gray-900">
                    {{ discount.type === 'percentage' ? discount.value + '%' : formatCurrency(discount.value) }}
                  </span>
                </div>
                <div class="flex justify-between">
                  <span class="text-sm text-gray-600">حداقل خرید:</span>
                  <span class="text-sm font-medium text-gray-900">{{ formatCurrency(discount.minAmount) }}</span>
                </div>
                <div class="flex justify-between">
                  <span class="text-sm text-gray-600">تاریخ انقضا:</span>
                  <span class="text-sm font-medium text-gray-900">{{ formatDate(discount.expiresAt) }}</span>
                </div>
              </div>

              <div class="flex gap-2">
                <button
                  class="flex-1 px-3 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
                  @click="editDiscount(discount)"
                >
                  ویرایش
                </button>
                <button
                  :class="{
                    'flex-1 px-3 py-2 text-sm font-medium rounded-lg focus:outline-none focus:ring-2 focus:ring-offset-2': true,
                    'bg-red-600 text-white hover:bg-red-700 focus:ring-red-500': discount.status === 'active',
                    'bg-green-600 text-white hover:bg-green-700 focus:ring-green-500': discount.status === 'inactive'
                  }"
                  @click="toggleDiscountStatus(discount)"
                >
                  {{ discount.status === 'active' ? 'غیرفعال' : 'فعال' }}
                </button>
              </div>
            </div>
          </div>
        </div>

        <!-- تب کارمزدها -->
        <div v-if="activeTab === 'commissions'">
          <div class="flex justify-between items-center mb-4">
            <h4 class="text-lg font-semibold text-gray-900">مدیریت کارمزدها</h4>
            <button
              class="px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2"
              @click="showCommissionModal = true"
            >
              تعریف کارمزد جدید
            </button>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            <div
              v-for="commission in commissions"
              :key="commission.id"
              class="bg-white border border-gray-200 rounded-lg p-6"
            >
              <div class="flex items-center justify-between mb-3">
                <h5 class="text-lg font-semibold text-gray-900">{{ commission.name }}</h5>
                <span class="text-2xl font-bold text-blue-600">{{ commission.rate }}%</span>
              </div>
              
              <div class="space-y-2 mb-4">
                <div class="flex justify-between">
                  <span class="text-sm text-gray-600">نوع:</span>
                  <span class="text-sm font-medium text-gray-900">{{ getCommissionTypeText(commission.type) }}</span>
                </div>
                <div class="flex justify-between">
                  <span class="text-sm text-gray-600">حداقل مبلغ:</span>
                  <span class="text-sm font-medium text-gray-900">{{ formatCurrency(commission.minAmount) }}</span>
                </div>
                <div class="flex justify-between">
                  <span class="text-sm text-gray-600">حداکثر مبلغ:</span>
                  <span class="text-sm font-medium text-gray-900">{{ formatCurrency(commission.maxAmount) }}</span>
                </div>
                <div class="flex justify-between">
                  <span class="text-sm text-gray-600">وضعیت:</span>
                  <span
                    :class="{
                      'bg-green-100 text-green-800': commission.status === 'active',
                      'bg-red-100 text-red-800': commission.status === 'inactive'
                    }"
                    class="px-2 py-1 text-xs font-medium rounded-full"
                  >
                    {{ getStatusText(commission.status) }}
                  </span>
                </div>
              </div>

              <div class="flex gap-2">
                <button
                  class="flex-1 px-3 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
                  @click="editCommission(commission)"
                >
                  ویرایش
                </button>
                <button
                  :class="{
                    'flex-1 px-3 py-2 text-sm font-medium rounded-lg focus:outline-none focus:ring-2 focus:ring-offset-2': true,
                    'bg-red-600 text-white hover:bg-red-700 focus:ring-red-500': commission.status === 'active',
                    'bg-green-600 text-white hover:bg-green-700 focus:ring-green-500': commission.status === 'inactive'
                  }"
                  @click="toggleCommissionStatus(commission)"
                >
                  {{ commission.status === 'active' ? 'غیرفعال' : 'فعال' }}
                </button>
              </div>
            </div>
          </div>
        </div>

        <!-- تب مناسبت‌ها -->
        <div v-if="activeTab === 'occasions'">
          <div class="flex justify-between items-center mb-4">
            <h4 class="text-lg font-semibold text-gray-900">قیمت‌گذاری بر اساس مناسبت‌ها</h4>
            <button
              class="px-4 py-2 bg-green-600 text-white text-sm font-medium rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2"
              @click="showOccasionModal = true"
            >
              تعریف مناسبت جدید
            </button>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            <div
              v-for="occasion in occasions"
              :key="occasion.id"
              class="bg-white border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow"
            >
              <div class="flex items-center justify-between mb-3">
                <h5 class="text-lg font-semibold text-gray-900">{{ occasion.name }}</h5>
                <span
                  :class="{
                    'bg-green-100 text-green-800': occasion.status === 'active',
                    'bg-red-100 text-red-800': occasion.status === 'inactive',
                    'bg-yellow-100 text-yellow-800': occasion.status === 'upcoming'
                  }"
                  class="px-2 py-1 text-xs font-medium rounded-full"
                >
                  {{ getStatusText(occasion.status) }}
                </span>
              </div>
              
              <div class="space-y-2 mb-4">
                <div class="flex justify-between">
                  <span class="text-sm text-gray-600">تاریخ شروع:</span>
                  <span class="text-sm font-medium text-gray-900">{{ formatDate(occasion.startDate) }}</span>
                </div>
                <div class="flex justify-between">
                  <span class="text-sm text-gray-600">تاریخ پایان:</span>
                  <span class="text-sm font-medium text-gray-900">{{ formatDate(occasion.endDate) }}</span>
                </div>
                <div class="flex justify-between">
                  <span class="text-sm text-gray-600">تخفیف:</span>
                  <span class="text-sm font-medium text-gray-900">{{ occasion.discount }}%</span>
                </div>
                <div class="flex justify-between">
                  <span class="text-sm text-gray-600">کارمزد ویژه:</span>
                  <span class="text-sm font-medium text-gray-900">{{ occasion.specialCommission }}%</span>
                </div>
              </div>

              <div class="flex gap-2">
                <button
                  class="flex-1 px-3 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
                  @click="editOccasion(occasion)"
                >
                  ویرایش
                </button>
                <button
                  :class="{
                    'flex-1 px-3 py-2 text-sm font-medium rounded-lg focus:outline-none focus:ring-2 focus:ring-offset-2': true,
                    'bg-red-600 text-white hover:bg-red-700 focus:ring-red-500': occasion.status === 'active',
                    'bg-green-600 text-white hover:bg-green-700 focus:ring-green-500': occasion.status === 'inactive'
                  }"
                  @click="toggleOccasionStatus(occasion)"
                >
                  {{ occasion.status === 'active' ? 'غیرفعال' : 'فعال' }}
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- مودال‌ها -->
    <GiftCardPricingCreateModal
      v-if="showNewPricingModal"
      @close="showNewPricingModal = false"
      @created="handlePricingCreated"
    />

    <GiftCardDiscountCreateModal
      v-if="showDiscountModal"
      @close="showDiscountModal = false"
      @created="handleDiscountCreated"
    />

    <GiftCardCommissionCreateModal
      v-if="showCommissionModal"
      @close="showCommissionModal = false"
      @created="handleCommissionCreated"
    />

    <GiftCardOccasionCreateModal
      v-if="showOccasionModal"
      @close="showOccasionModal = false"
      @created="handleOccasionCreated"
    />
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from "vue"

// تعریف interface ها (Rule 67 - Interface Organization and Reusability)
interface Pricing {
  id?: number | string
  status?: string
  [key: string]: unknown
}

interface Discount {
  id?: number | string
  status?: string
  [key: string]: unknown
}

interface Commission {
  id?: number | string
  status?: string
  [key: string]: unknown
}

interface Occasion {
  id?: number | string
  status?: string
  [key: string]: unknown
}

// تعریف متغیرهای reactive
const activeTab = ref('pricing')
const showNewPricingModal = ref(false)
const showDiscountModal = ref(false)
const showCommissionModal = ref(false)
const showOccasionModal = ref(false)

const pricingSearch = ref('')
const pricingTypeFilter = ref('')
const discountSearch = ref('')
const discountTypeFilter = ref('')

// داده‌های نمونه قیمت‌گذاری
const pricingData = ref([
  {
    id: 1,
    name: 'کارت تولد',
    type: 'themed',
    baseAmount: 500000,
    commission: 5,
    finalPrice: 525000,
    status: 'active'
  },
  {
    id: 2,
    name: 'کارت عروسی',
    type: 'themed',
    baseAmount: 1000000,
    commission: 7,
    finalPrice: 1070000,
    status: 'active'
  },
  {
    id: 3,
    name: 'کارت سال نو',
    type: 'themed',
    baseAmount: 750000,
    commission: 6,
    finalPrice: 795000,
    status: 'active'
  },
  {
    id: 4,
    name: 'کارت شرکتی',
    type: 'corporate',
    baseAmount: 2000000,
    commission: 3,
    finalPrice: 2060000,
    status: 'inactive'
  }
])

// داده‌های نمونه تخفیف‌ها
const discounts = ref([
  {
    id: 1,
    name: 'تخفیف ویژه تولد',
    type: 'percentage',
    value: 15,
    minAmount: 500000,
    expiresAt: new Date(Date.now() + 2592000000),
    status: 'active'
  },
  {
    id: 2,
    name: 'تخفیف عروسی',
    type: 'fixed',
    value: 200000,
    minAmount: 1000000,
    expiresAt: new Date(Date.now() + 5184000000),
    status: 'active'
  },
  {
    id: 3,
    name: 'تخفیف سال نو',
    type: 'percentage',
    value: 20,
    minAmount: 750000,
    expiresAt: new Date(Date.now() + 7776000000),
    status: 'scheduled'
  }
])

// داده‌های نمونه کارمزدها
const commissions = ref([
  {
    id: 1,
    name: 'کارمزد استاندارد',
    type: 'percentage',
    rate: 5,
    minAmount: 0,
    maxAmount: 1000000,
    status: 'active'
  },
  {
    id: 2,
    name: 'کارمزد ویژه',
    type: 'percentage',
    rate: 3,
    minAmount: 1000000,
    maxAmount: 5000000,
    status: 'active'
  },
  {
    id: 3,
    name: 'کارمزد شرکتی',
    type: 'fixed',
    rate: 50000,
    minAmount: 5000000,
    maxAmount: 10000000,
    status: 'inactive'
  }
])

// داده‌های نمونه مناسبت‌ها
const occasions = ref([
  {
    id: 1,
    name: 'تخفیف تولد',
    startDate: new Date(Date.now() - 86400000),
    endDate: new Date(Date.now() + 2592000000),
    discount: 15,
    specialCommission: 2,
    status: 'active'
  },
  {
    id: 2,
    name: 'تخفیف عروسی',
    startDate: new Date(Date.now() + 86400000),
    endDate: new Date(Date.now() + 5184000000),
    discount: 20,
    specialCommission: 3,
    status: 'upcoming'
  },
  {
    id: 3,
    name: 'تخفیف سال نو',
    startDate: new Date(Date.now() + 2592000000),
    endDate: new Date(Date.now() + 7776000000),
    discount: 25,
    specialCommission: 1,
    status: 'upcoming'
  }
])

// محاسبات
const activePricing = computed(() => 
  pricingData.value.filter(p => p.status === 'active')
)

const activeDiscounts = computed(() => 
  discounts.value.filter(d => d.status === 'active')
)

const averageCommission = computed(() => {
  const activeCommissions = commissions.value.filter(c => c.status === 'active')
  if (activeCommissions.length === 0) return 0
  const total = activeCommissions.reduce((sum, c) => sum + c.rate, 0)
  return (total / activeCommissions.length).toFixed(1)
})

const totalCommissionRevenue = computed(() => {
  // محاسبه درآمد کارمزد بر اساس تراکنش‌های نمونه
  return 2340000
})

// فیلترها
const filteredPricing = computed(() => {
  return pricingData.value.filter(pricing => {
    const matchesSearch = 
      pricing.name.toLowerCase().includes(pricingSearch.value.toLowerCase())
    
    const matchesType = !pricingTypeFilter.value || pricing.type === pricingTypeFilter.value
    
    return matchesSearch && matchesType
  })
})

const filteredDiscounts = computed(() => {
  return discounts.value.filter(discount => {
    const matchesSearch = 
      discount.name.toLowerCase().includes(discountSearch.value.toLowerCase())
    
    const matchesType = !discountTypeFilter.value || discount.type === discountTypeFilter.value
    
    return matchesSearch && matchesType
  })
})

// توابع کمکی
const getTypeText = (type) => {
  const typeMap = {
    fixed: 'مبلغ ثابت',
    variable: 'مبلغ متغیر',
    themed: 'موضوعی',
    corporate: 'شرکتی'
  }
  return typeMap[type] || type
}

const getStatusText = (status) => {
  const statusMap = {
    active: 'فعال',
    inactive: 'غیرفعال',
    pending: 'در انتظار',
    scheduled: 'برنامه‌ریزی شده',
    upcoming: 'آینده'
  }
  return statusMap[status] || status
}

const getDiscountTypeText = (type) => {
  const typeMap = {
    percentage: 'درصدی',
    fixed: 'مبلغ ثابت',
    buy_one_get_one: 'یک به یک'
  }
  return typeMap[type] || type
}

const getCommissionTypeText = (type) => {
  const typeMap = {
    percentage: 'درصدی',
    fixed: 'مبلغ ثابت'
  }
  return typeMap[type] || type
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
const editPricing = (_pricing: Pricing) => {

  // اینجا کد ویرایش قیمت را اضافه کنید
}

const togglePricingStatus = (pricing: Pricing) => {
  pricing.status = pricing.status === 'active' ? 'inactive' : 'active'
}

const editDiscount = (_discount: Discount) => {

  // اینجا کد ویرایش تخفیف را اضافه کنید
}

const toggleDiscountStatus = (discount: Discount) => {
  discount.status = discount.status === 'active' ? 'inactive' : 'active'
}

const editCommission = (_commission: Commission) => {

  // اینجا کد ویرایش کارمزد را اضافه کنید
}

const toggleCommissionStatus = (commission: Commission) => {
  commission.status = commission.status === 'active' ? 'inactive' : 'active'
}

const editOccasion = (_occasion: Occasion) => {

  // اینجا کد ویرایش مناسبت را اضافه کنید
}

const toggleOccasionStatus = (occasion) => {
  occasion.status = occasion.status === 'active' ? 'inactive' : 'active'
}

const exportPricing = () => {

  alert('گزارش قیمت‌گذاری در حال دانلود است')
}

// توابع مدیریت مودال‌ها
const handlePricingCreated = (newPricing) => {
  pricingData.value.unshift({
    ...newPricing,
    id: Date.now(),
    finalPrice: newPricing.baseAmount * (1 + newPricing.commission / 100)
  })
  showNewPricingModal.value = false
}

const handleDiscountCreated = (newDiscount) => {
  discounts.value.unshift({
    ...newDiscount,
    id: Date.now()
  })
  showDiscountModal.value = false
}

const handleCommissionCreated = (newCommission) => {
  commissions.value.unshift({
    ...newCommission,
    id: Date.now()
  })
  showCommissionModal.value = false
}

const handleOccasionCreated = (newOccasion) => {
  occasions.value.unshift({
    ...newOccasion,
    id: Date.now()
  })
  showOccasionModal.value = false
}
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 
