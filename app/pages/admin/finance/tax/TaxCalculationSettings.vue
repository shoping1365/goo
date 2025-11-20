<template>
  <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
    <!-- هدر -->
    <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-6 mb-6">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">تنظیمات محاسباتی مالیات</h3>
        <p class="text-sm text-gray-600">تنظیم فرمول‌ها، قوانین و پارامترهای محاسباتی</p>
      </div>
      
      <!-- دکمه ذخیره تنظیمات -->
      <button 
        class="inline-flex items-center gap-2 px-4 py-2 bg-green-600 hover:bg-green-700 text-white rounded-lg transition-colors duration-200"
        @click="saveSettings"
      >
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
        </svg>
        ذخیره تنظیمات
      </button>
    </div>

    <!-- تب‌های تنظیمات -->
    <div class="mb-6">
      <div class="border-b border-gray-200">
        <nav class="-mb-px flex space-x-8">
          <button 
            v-for="tab in tabs" 
            :key="tab.id"
            :class="[
              'py-2 px-1 border-b-2 font-medium text-sm transition-colors duration-200',
              activeTab === tab.id 
                ? 'border-green-500 text-green-600' 
                : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
            ]"
            @click="activeTab = tab.id"
          >
            {{ tab.label }}
          </button>
        </nav>
      </div>
    </div>

    <!-- تنظیمات عمومی -->
    <div v-if="activeTab === 'general'">
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <!-- تنظیمات محاسباتی -->
        <div class="bg-gray-50 rounded-lg p-6">
          <h4 class="font-medium text-gray-900 mb-4">تنظیمات محاسباتی</h4>
          
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">دقت محاسبات</label>
              <select 
                v-model="settings.calculationPrecision"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-green-500 focus:border-green-500"
              >
                <option value="0">بدون اعشار</option>
                <option value="1">یک رقم اعشار</option>
                <option value="2">دو رقم اعشار</option>
                <option value="3">سه رقم اعشار</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">روش گرد کردن</label>
              <select 
                v-model="settings.roundingMethod"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-green-500 focus:border-green-500"
              >
                <option value="nearest">نزدیک‌ترین عدد</option>
                <option value="up">گرد به بالا</option>
                <option value="down">گرد به پایین</option>
                <option value="bankers">گرد بانکی</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">حداقل مبلغ برای محاسبه مالیات</label>
              <input 
                v-model.number="settings.minTaxAmount"
                type="number"
                min="0"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-green-500 focus:border-green-500"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر مبلغ برای معافیت</label>
              <input 
                v-model.number="settings.maxExemptionAmount"
                type="number"
                min="0"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-green-500 focus:border-green-500"
              />
            </div>
          </div>
        </div>

        <!-- تنظیمات خودکار -->
        <div class="bg-gray-50 rounded-lg p-6">
          <h4 class="font-medium text-gray-900 mb-4">تنظیمات خودکار</h4>
          
          <div class="space-y-4">
            <label class="flex items-center">
              <input 
                v-model="settings.autoCalculate"
                type="checkbox"
                class="rounded border-gray-300 text-green-600 focus:ring-green-500"
              />
              <span class="mr-2 text-sm text-gray-700">محاسبه خودکار مالیات</span>
            </label>

            <label class="flex items-center">
              <input 
                v-model="settings.autoApplyExemptions"
                type="checkbox"
                class="rounded border-gray-300 text-green-600 focus:ring-green-500"
              />
              <span class="mr-2 text-sm text-gray-700">اعمال خودکار معافیت‌ها</span>
            </label>

            <label class="flex items-center">
              <input 
                v-model="settings.autoValidateDiscounts"
                type="checkbox"
                class="rounded border-gray-300 text-green-600 focus:ring-green-500"
              />
              <span class="mr-2 text-sm text-gray-700">اعتبارسنجی خودکار تخفیف‌ها</span>
            </label>

            <label class="flex items-center">
              <input 
                v-model="settings.autoSaveCalculations"
                type="checkbox"
                class="rounded border-gray-300 text-green-600 focus:ring-green-500"
              />
              <span class="mr-2 text-sm text-gray-700">ذخیره خودکار محاسبات</span>
            </label>

            <label class="flex items-center">
              <input 
                v-model="settings.autoGenerateReports"
                type="checkbox"
                class="rounded border-gray-300 text-green-600 focus:ring-green-500"
              />
              <span class="mr-2 text-sm text-gray-700">تولید خودکار گزارش‌ها</span>
            </label>
          </div>
        </div>
      </div>
    </div>

    <!-- فرمول‌های محاسباتی -->
    <div v-else-if="activeTab === 'formulas'">
      <div class="space-y-6">
        <!-- فرمول مالیات بر ارزش افزوده -->
        <div class="bg-blue-50 rounded-lg p-6">
          <h4 class="font-medium text-blue-900 mb-4">فرمول مالیات بر ارزش افزوده</h4>
          
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-blue-700 mb-2">فرمول محاسبه</label>
              <select 
                v-model="settings.vatFormula"
                class="w-full px-3 py-2 border border-blue-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="simple">ساده: مبلغ × نرخ</option>
                <option value="inclusive">شامل: مبلغ × (نرخ / (1 + نرخ))</option>
                <option value="exclusive">خارج از: مبلغ × نرخ</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-blue-700 mb-2">نرخ پیش‌فرض</label>
              <input 
                v-model.number="settings.defaultVatRate"
                type="number"
                step="0.1"
                min="0"
                max="100"
                class="w-full px-3 py-2 border border-blue-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
          </div>

          <div class="mt-4">
            <label class="block text-sm font-medium text-blue-700 mb-2">شرایط اعمال</label>
            <textarea 
              v-model="settings.vatConditions"
              rows="3"
              placeholder="شرایط اعمال مالیات بر ارزش افزوده..."
              class="w-full px-3 py-2 border border-blue-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            ></textarea>
          </div>
        </div>

        <!-- فرمول مالیات بر درآمد -->
        <div class="bg-green-50 rounded-lg p-6">
          <h4 class="font-medium text-green-900 mb-4">فرمول مالیات بر درآمد</h4>
          
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-green-700 mb-2">نوع محاسبه</label>
              <select 
                v-model="settings.incomeTaxType"
                class="w-full px-3 py-2 border border-green-300 rounded-lg text-sm focus:ring-2 focus:ring-green-500 focus:border-green-500"
              >
                <option value="flat">نرخ ثابت</option>
                <option value="progressive">نرخ تصاعدی</option>
                <option value="bracket">طبقه‌بندی</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-green-700 mb-2">نرخ پیش‌فرض</label>
              <input 
                v-model.number="settings.defaultIncomeTaxRate"
                type="number"
                step="0.1"
                min="0"
                max="100"
                class="w-full px-3 py-2 border border-green-300 rounded-lg text-sm focus:ring-2 focus:ring-green-500 focus:border-green-500"
              />
            </div>
          </div>

          <!-- طبقه‌بندی مالیات بر درآمد -->
          <div v-if="settings.incomeTaxType === 'bracket'" class="mt-4">
            <label class="block text-sm font-medium text-green-700 mb-2">طبقه‌بندی مالیات</label>
            <div class="space-y-2">
              <div v-for="(bracket, index) in settings.incomeTaxBrackets" :key="index" class="flex gap-2">
                <input 
                  v-model.number="bracket.from"
                  type="number"
                  placeholder="از"
                  class="flex-1 px-3 py-2 border border-green-300 rounded-lg text-sm"
                />
                <input 
                  v-model.number="bracket.to"
                  type="number"
                  placeholder="تا"
                  class="flex-1 px-3 py-2 border border-green-300 rounded-lg text-sm"
                />
                <input 
                  v-model.number="bracket.rate"
                  type="number"
                  step="0.1"
                  placeholder="نرخ %"
                  class="flex-1 px-3 py-2 border border-green-300 rounded-lg text-sm"
                />
                <button 
                  class="px-3 py-2 bg-red-100 text-red-600 rounded-lg hover:bg-red-200"
                  @click="removeBracket(index)"
                >
                  حذف
                </button>
              </div>
              <button 
                class="w-full px-3 py-2 bg-green-100 text-green-600 rounded-lg hover:bg-green-200"
                @click="addBracket"
              >
                افزودن طبقه جدید
              </button>
            </div>
          </div>
        </div>

        <!-- فرمول مالیات گمرکی -->
        <div class="bg-purple-50 rounded-lg p-6">
          <h4 class="font-medium text-purple-900 mb-4">فرمول مالیات گمرکی</h4>
          
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-purple-700 mb-2">نرخ پیش‌فرض</label>
              <input 
                v-model.number="settings.defaultCustomsTaxRate"
                type="number"
                step="0.1"
                min="0"
                max="100"
                class="w-full px-3 py-2 border border-purple-300 rounded-lg text-sm focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-purple-700 mb-2">حداقل مبلغ</label>
              <input 
                v-model.number="settings.minCustomsTaxAmount"
                type="number"
                min="0"
                class="w-full px-3 py-2 border border-purple-300 rounded-lg text-sm focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
              />
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- قوانین و استثناها -->
    <div v-else-if="activeTab === 'rules'">
      <div class="space-y-6">
        <!-- قوانین معافیت -->
        <div class="bg-yellow-50 rounded-lg p-6">
          <h4 class="font-medium text-yellow-900 mb-4">قوانین معافیت</h4>
          
          <div class="space-y-4">
            <div v-for="(rule, index) in settings.exemptionRules" :key="index" class="border border-yellow-200 rounded-lg p-3">
              <div class="grid grid-cols-1 md:grid-cols-3 gap-3">
                <div>
                  <label class="block text-sm font-medium text-yellow-700 mb-1">نوع محصول</label>
                  <select 
                    v-model="rule.productType"
                    class="w-full px-3 py-2 border border-yellow-300 rounded-lg text-sm"
                  >
                    <option value="food">مواد غذایی</option>
                    <option value="medical">پزشکی</option>
                    <option value="education">آموزشی</option>
                    <option value="digital">دیجیتال</option>
                  </select>
                </div>
                <div>
                  <label class="block text-sm font-medium text-yellow-700 mb-1">درصد معافیت</label>
                  <input 
                    v-model.number="rule.percentage"
                    type="number"
                    min="0"
                    max="100"
                    class="w-full px-3 py-2 border border-yellow-300 rounded-lg text-sm"
                  />
                </div>
                <div class="flex items-end">
                  <button 
                    class="w-full px-3 py-2 bg-red-100 text-red-600 rounded-lg hover:bg-red-200"
                    @click="removeExemptionRule(index)"
                  >
                    حذف
                  </button>
                </div>
              </div>
            </div>
            <button 
              class="w-full px-4 py-2 bg-yellow-100 text-yellow-700 rounded-lg hover:bg-yellow-200"
              @click="addExemptionRule"
            >
              افزودن قانون معافیت جدید
            </button>
          </div>
        </div>

        <!-- قوانین تخفیف -->
        <div class="bg-orange-50 rounded-lg p-6">
          <h4 class="font-medium text-orange-900 mb-4">قوانین تخفیف</h4>
          
          <div class="space-y-4">
            <div v-for="(rule, index) in settings.discountRules" :key="index" class="border border-orange-200 rounded-lg p-3">
              <div class="grid grid-cols-1 md:grid-cols-4 gap-3">
                <div>
                  <label class="block text-sm font-medium text-orange-700 mb-1">شرط</label>
                  <select 
                    v-model="rule.condition"
                    class="w-full px-3 py-2 border border-orange-300 rounded-lg text-sm"
                  >
                    <option value="amount">مبلغ سفارش</option>
                    <option value="quantity">تعداد محصول</option>
                    <option value="customer_type">نوع مشتری</option>
                    <option value="region">منطقه</option>
                  </select>
                </div>
                <div>
                  <label class="block text-sm font-medium text-orange-700 mb-1">مقدار</label>
                  <input 
                    v-model.number="rule.value"
                    type="number"
                    class="w-full px-3 py-2 border border-orange-300 rounded-lg text-sm"
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium text-orange-700 mb-1">درصد تخفیف</label>
                  <input 
                    v-model.number="rule.discountPercentage"
                    type="number"
                    min="0"
                    max="100"
                    class="w-full px-3 py-2 border border-orange-300 rounded-lg text-sm"
                  />
                </div>
                <div class="flex items-end">
                  <button 
                    class="w-full px-3 py-2 bg-red-100 text-red-600 rounded-lg hover:bg-red-200"
                    @click="removeDiscountRule(index)"
                  >
                    حذف
                  </button>
                </div>
              </div>
            </div>
            <button 
              class="w-full px-4 py-2 bg-orange-100 text-orange-700 rounded-lg hover:bg-orange-200"
              @click="addDiscountRule"
            >
              افزودن قانون تخفیف جدید
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- پیش‌نمایش و تست -->
    <div v-else-if="activeTab === 'preview'">
      <div class="space-y-6">
        <!-- پیش‌نمایش فرمول -->
        <div class="bg-gray-50 rounded-lg p-6">
          <h4 class="font-medium text-gray-900 mb-4">پیش‌نمایش فرمول‌ها</h4>
          
          <div class="space-y-3 text-sm">
            <div class="bg-white rounded-lg p-3">
              <div class="font-medium text-gray-900 mb-2">مالیات بر ارزش افزوده:</div>
              <div class="text-gray-600 font-mono">{{ getVatFormulaPreview() }}</div>
            </div>
            
            <div class="bg-white rounded-lg p-3">
              <div class="font-medium text-gray-900 mb-2">مالیات بر درآمد:</div>
              <div class="text-gray-600 font-mono">{{ getIncomeTaxFormulaPreview() }}</div>
            </div>
            
            <div class="bg-white rounded-lg p-3">
              <div class="font-medium text-gray-900 mb-2">مالیات گمرکی:</div>
              <div class="text-gray-600 font-mono">{{ getCustomsTaxFormulaPreview() }}</div>
            </div>
          </div>
        </div>

        <!-- تست محاسبات -->
        <div class="bg-gray-50 rounded-lg p-6">
          <h4 class="font-medium text-gray-900 mb-4">تست محاسبات</h4>
          
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">مبلغ تست</label>
              <input 
                v-model.number="testAmount"
                type="number"
                min="0"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm"
              />
            </div>
            <div class="flex items-end">
              <button 
                class="w-full px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg"
                @click="testCalculation"
              >
                تست محاسبه
              </button>
            </div>
          </div>

          <div v-if="testResult" class="mt-4 bg-white rounded-lg p-3">
            <div class="text-sm space-y-1">
              <div class="flex justify-between">
                <span class="text-gray-600">مالیات بر ارزش افزوده:</span>
                <span class="font-medium">{{ formatCurrency(testResult.vat) }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">مالیات بر درآمد:</span>
                <span class="font-medium">{{ formatCurrency(testResult.incomeTax) }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">مالیات گمرکی:</span>
                <span class="font-medium">{{ formatCurrency(testResult.customsTax) }}</span>
              </div>
              <div class="border-t pt-2">
                <div class="flex justify-between font-medium">
                  <span>کل مالیات:</span>
                  <span>{{ formatCurrency(testResult.totalTax) }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'

// تب‌های تنظیمات
const tabs = [
  { id: 'general', label: 'تنظیمات عمومی' },
  { id: 'formulas', label: 'فرمول‌های محاسباتی' },
  { id: 'rules', label: 'قوانین و استثناها' },
  { id: 'preview', label: 'پیش‌نمایش و تست' }
]

const activeTab = ref('general')

// تنظیمات
const settings = ref({
  // تنظیمات عمومی
  calculationPrecision: 2,
  roundingMethod: 'nearest',
  minTaxAmount: 1000,
  maxExemptionAmount: 1000000,
  
  // تنظیمات خودکار
  autoCalculate: true,
  autoApplyExemptions: true,
  autoValidateDiscounts: true,
  autoSaveCalculations: false,
  autoGenerateReports: false,
  
  // فرمول‌های مالیاتی
  vatFormula: 'simple',
  defaultVatRate: 9,
  vatConditions: 'اعمال بر تمام محصولات و خدمات',
  
  incomeTaxType: 'flat',
  defaultIncomeTaxRate: 15,
  incomeTaxBrackets: [
    { from: 0, to: 50000000, rate: 10 },
    { from: 50000000, to: 100000000, rate: 15 },
    { from: 100000000, to: 0, rate: 20 }
  ],
  
  defaultCustomsTaxRate: 5,
  minCustomsTaxAmount: 50000,
  
  // قوانین معافیت
  exemptionRules: [
    { productType: 'food', percentage: 100 },
    { productType: 'medical', percentage: 25 },
    { productType: 'education', percentage: 50 }
  ],
  
  // قوانین تخفیف
  discountRules: [
    { condition: 'amount', value: 1000000, discountPercentage: 5 },
    { condition: 'customer_type', value: 'business', discountPercentage: 10 }
  ]
})

// تست محاسبات
const testAmount = ref(1000000)
const testResult = ref(null)

// فرمت مبلغ
const formatCurrency = (amount: number): string => {
  return new Intl.NumberFormat('fa-IR', {
    style: 'currency',
    currency: 'IRR',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount)
}

// افزودن طبقه مالیاتی
const addBracket = () => {
  settings.value.incomeTaxBrackets.push({ from: 0, to: 0, rate: 0 })
}

// حذف طبقه مالیاتی
const removeBracket = (index: number) => {
  settings.value.incomeTaxBrackets.splice(index, 1)
}

// افزودن قانون معافیت
const addExemptionRule = () => {
  settings.value.exemptionRules.push({ productType: 'food', percentage: 0 })
}

// حذف قانون معافیت
const removeExemptionRule = (index: number) => {
  settings.value.exemptionRules.splice(index, 1)
}

// افزودن قانون تخفیف
const addDiscountRule = () => {
  settings.value.discountRules.push({ condition: 'amount', value: 0, discountPercentage: 0 })
}

// حذف قانون تخفیف
const removeDiscountRule = (index: number) => {
  settings.value.discountRules.splice(index, 1)
}

// پیش‌نمایش فرمول مالیات بر ارزش افزوده
const getVatFormulaPreview = () => {
  const formulas = {
    simple: `مالیات = مبلغ × ${settings.value.defaultVatRate}%`,
    inclusive: `مالیات = مبلغ × (${settings.value.defaultVatRate}% / (1 + ${settings.value.defaultVatRate}%))`,
    exclusive: `مالیات = مبلغ × ${settings.value.defaultVatRate}%`
  }
  return formulas[settings.value.vatFormula]
}

// پیش‌نمایش فرمول مالیات بر درآمد
const getIncomeTaxFormulaPreview = () => {
  if (settings.value.incomeTaxType === 'flat') {
    return `مالیات = مبلغ × ${settings.value.defaultIncomeTaxRate}%`
  } else if (settings.value.incomeTaxType === 'progressive') {
    return 'مالیات = محاسبه تصاعدی بر اساس طبقه‌بندی'
  } else {
    return 'مالیات = محاسبه بر اساس طبقه‌بندی تعریف شده'
  }
}

// پیش‌نمایش فرمول مالیات گمرکی
const getCustomsTaxFormulaPreview = () => {
  return `مالیات = مبلغ × ${settings.value.defaultCustomsTaxRate}% (حداقل ${formatCurrency(settings.value.minCustomsTaxAmount)})`
}

// تست محاسبه
const testCalculation = () => {
  const amount = testAmount.value || 0
  
  // محاسبه مالیات بر ارزش افزوده
  let vat = 0
  if (settings.value.vatFormula === 'simple') {
    vat = (amount * settings.value.defaultVatRate) / 100
  }
  
  // محاسبه مالیات بر درآمد
  let incomeTax = 0
  if (settings.value.incomeTaxType === 'flat') {
    incomeTax = (amount * settings.value.defaultIncomeTaxRate) / 100
  }
  
  // محاسبه مالیات گمرکی
  let customsTax = (amount * settings.value.defaultCustomsTaxRate) / 100
  if (customsTax < settings.value.minCustomsTaxAmount) {
    customsTax = settings.value.minCustomsTaxAmount
  }
  
  const totalTax = vat + incomeTax + customsTax
  
  testResult.value = {
    vat,
    incomeTax,
    customsTax,
    totalTax
  }
}

// ذخیره تنظیمات
const saveSettings = async () => {
  try {
    // TODO: ارسال درخواست به API
    console.log('تنظیمات ذخیره شد:', settings.value)
  } catch (error) {
    console.error('خطا در ذخیره تنظیمات:', error)
  }
}

// بارگذاری اولیه
onMounted(() => {
  // TODO: بارگذاری تنظیمات از API
})
</script>

<!--
  کامپوننت تنظیمات محاسباتی مالیات
  شامل:
  1. تنظیمات عمومی محاسبات
  2. فرمول‌های محاسباتی مختلف
  3. قوانین معافیت و تخفیف
  4. پیش‌نمایش و تست فرمول‌ها
  5. طراحی مدرن و کاملاً ریسپانسیو
--> 
