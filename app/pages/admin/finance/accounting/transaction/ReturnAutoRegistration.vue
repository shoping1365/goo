<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">ثبت خودکار برگشت‌ها</h4>
        <p class="text-sm text-gray-600 mt-1">تنظیمات ثبت خودکار برگشت‌ها و بازپرداخت‌ها در نرم‌افزار حسابداری</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          :class="[
            autoRegistrationEnabled ? 'bg-green-600 hover:bg-green-700' : 'bg-gray-600 hover:bg-gray-700',
            'inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500'
          ]"
          @click="toggleAutoRegistration"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path v-if="autoRegistrationEnabled" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
            <path v-else stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
          {{ autoRegistrationEnabled ? 'فعال' : 'غیرفعال' }}
        </button>
        <button
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="showAddForm = true"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
          </svg>
          افزودن قوانین جدید
        </button>
      </div>
    </div>

    <!-- تنظیمات کلی -->
    <div class="bg-gray-50 rounded-lg p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">تنظیمات کلی</h5>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <!-- حساب برگشت فروش -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حساب برگشت فروش</label>
          <select
            v-model="settings.returnSalesAccount"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">انتخاب حساب</option>
            <option value="return-sales">برگشت فروش</option>
            <option value="sales-returns">بازگشت از فروش</option>
          </select>
        </div>

        <!-- حساب بازپرداخت -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حساب بازپرداخت</label>
          <select
            v-model="settings.refundAccount"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">انتخاب حساب</option>
            <option value="refund-payable">بازپرداخت قابل پرداخت</option>
            <option value="customer-refunds">بازپرداخت مشتریان</option>
          </select>
        </div>

        <!-- حساب مالیات برگشتی -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حساب مالیات برگشتی</label>
          <select
            v-model="settings.returnTaxAccount"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">انتخاب حساب</option>
            <option value="tax-refund">مالیات قابل بازگشت</option>
            <option value="vat-refund">مالیات بر ارزش افزوده قابل بازگشت</option>
          </select>
        </div>

        <!-- زمان ثبت -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">زمان ثبت</label>
          <select
            v-model="settings.registrationTime"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="immediate">فوری</option>
            <option value="hourly">ساعتی</option>
            <option value="daily">روزانه</option>
            <option value="manual">دستی</option>
          </select>
        </div>

        <!-- وضعیت برگشت -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت برگشت</label>
          <select
            v-model="settings.returnStatus"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="approved">تایید شده</option>
            <option value="completed">تکمیل شده</option>
            <option value="all">همه</option>
          </select>
        </div>

        <!-- حداقل مبلغ -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداقل مبلغ (تومان)</label>
          <input
            v-model="settings.minAmount"
            type="number"
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            placeholder="0"
          />
        </div>
      </div>
    </div>

    <!-- قوانین ثبت خودکار -->
    <div class="bg-white rounded-lg border border-gray-200">
      <div class="px-6 py-4 border-b border-gray-200">
        <h5 class="text-md font-medium text-gray-900">قوانین ثبت خودکار</h5>
      </div>
      <div class="overflow-x-auto">
        <table class="w-full text-sm">
          <thead>
            <tr class="border-b border-gray-200 bg-gray-50">
              <th class="text-right py-3 px-4 font-medium text-gray-600">نام قانون</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">شرایط</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">حساب مقصد</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">وضعیت</th>
              <th class="text-right py-3 px-4 font-medium text-gray-600">عملیات</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="rule in autoRegistrationRules" :key="rule.id" class="border-b border-gray-100 hover:bg-gray-50">
              <td class="py-3 px-4">
                <div class="font-medium text-gray-900">{{ rule.name }}</div>
                <div class="text-xs text-gray-500">{{ rule.description }}</div>
              </td>
              <td class="py-3 px-4">
                <div class="text-sm text-gray-900">{{ rule.conditions }}</div>
              </td>
              <td class="py-3 px-4">
                <div class="text-sm text-gray-900">{{ rule.targetAccount }}</div>
              </td>
              <td class="py-3 px-4">
                <span
                  :class="[
                    rule.active ? 'bg-green-100 text-green-700' : 'bg-red-100 text-red-700',
                    'px-2 py-1 rounded-full text-xs font-medium'
                  ]"
                >
                  {{ rule.active ? 'فعال' : 'غیرفعال' }}
                </span>
              </td>
              <td class="py-3 px-4">
                <div class="flex items-center space-x-2 space-x-reverse">
                  <button
                    class="p-1 text-blue-600 hover:text-blue-800"
                    title="ویرایش"
                    @click="editRule(rule)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                    </svg>
                  </button>
                  <button
                    :class="[
                      rule.active ? 'text-red-600 hover:text-red-800' : 'text-green-600 hover:text-green-800'
                    ]"
                    class="p-1"
                    :title="rule.active ? 'غیرفعال کردن' : 'فعال کردن'"
                    @click="toggleRule(rule)"
                  >
                    <svg v-if="rule.active" class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728L5.636 5.636m12.728 12.728L18.364 5.636M5.636 18.364l12.728-12.728" />
                    </svg>
                    <svg v-else class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                    </svg>
                  </button>
                  <button
                    class="p-1 text-red-600 hover:text-red-800"
                    title="حذف"
                    @click="deleteRule(rule)"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- آمار ثبت خودکار -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
      <div class="bg-blue-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-blue-600">{{ stats.totalReturns }}</div>
            <div class="text-sm text-blue-700">کل برگشت‌ها</div>
          </div>
          <div class="w-10 h-10 bg-blue-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h10a8 8 0 018 8v2M3 10l6 6m-6-6l6-6" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-green-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-green-600">{{ stats.autoRegistered }}</div>
            <div class="text-sm text-green-700">ثبت خودکار</div>
          </div>
          <div class="w-10 h-10 bg-green-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-yellow-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-yellow-600">{{ stats.pendingRegistration }}</div>
            <div class="text-sm text-yellow-700">در انتظار ثبت</div>
          </div>
          <div class="w-10 h-10 bg-yellow-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// وضعیت ثبت خودکار
const autoRegistrationEnabled = ref(true)
const showAddForm = ref(false)

// تنظیمات کلی
const settings = ref({
  returnSalesAccount: 'return-sales',
  refundAccount: 'refund-payable',
  returnTaxAccount: 'tax-refund',
  registrationTime: 'immediate',
  returnStatus: 'approved',
  minAmount: 50000
})

// قوانین ثبت خودکار
const autoRegistrationRules = ref([
  {
    id: 1,
    name: 'برگشت‌های کامل',
    description: 'ثبت خودکار برگشت‌های کامل کالا',
    conditions: 'نوع برگشت = کامل',
    targetAccount: 'برگشت فروش',
    active: true
  },
  {
    id: 2,
    name: 'برگشت‌های جزئی',
    description: 'ثبت خودکار برگشت‌های جزئی کالا',
    conditions: 'نوع برگشت = جزئی',
    targetAccount: 'برگشت فروش',
    active: true
  },
  {
    id: 3,
    name: 'بازپرداخت‌ها',
    description: 'ثبت خودکار بازپرداخت‌های نقدی',
    conditions: 'نوع = بازپرداخت نقدی',
    targetAccount: 'بازپرداخت قابل پرداخت',
    active: false
  }
])

// آمار
const stats = ref({
  totalReturns: 45,
  autoRegistered: 38,
  pendingRegistration: 7
})

// تغییر وضعیت ثبت خودکار
const toggleAutoRegistration = () => {
  autoRegistrationEnabled.value = !autoRegistrationEnabled.value
  // TODO: ذخیره در سرور
}

interface Rule {
  id?: number | string
  active?: boolean
  [key: string]: unknown
}

// ویرایش قانون
const editRule = (_rule: Rule) => {
  // TODO: باز کردن فرم ویرایش

}

// تغییر وضعیت قانون
const toggleRule = (rule: Rule) => {
  rule.active = !rule.active
  // TODO: ذخیره در سرور
}

// حذف قانون
const deleteRule = (_rule: Rule) => {
  // TODO: حذف قانون

}
</script>

<!--
  کامپوننت ثبت خودکار برگشت‌ها
  شامل:
  1. تنظیمات کلی ثبت خودکار
  2. مدیریت قوانین ثبت خودکار
  3. آمار و گزارش‌گیری
  4. طراحی مدرن و کاملاً ریسپانسیو
--> 
