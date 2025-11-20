<template>
  <div class="space-y-6">
    <!-- اطلاعات فاکتور -->
    <div class="px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-blue-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
        </svg>
        اطلاعات فاکتور
      </h3>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
        <!-- اطلاعات اصلی فاکتور -->
        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">اطلاعات اصلی</h4>
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">شماره فاکتور</label>
              <input 
                v-model="invoice.invoiceNumber"
                type="text" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="INV-2024-001"
              />
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">تاریخ فاکتور</label>
              <input 
                v-model="invoice.invoiceDate"
                type="date" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">تاریخ سررسید</label>
              <input 
                v-model="invoice.dueDate"
                type="date" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">وضعیت فاکتور</label>
              <select 
                v-model="invoice.status"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="draft">پیش‌نویس</option>
                <option value="sent">ارسال شده</option>
                <option value="paid">پرداخت شده</option>
                <option value="overdue">معوق</option>
                <option value="cancelled">لغو شده</option>
              </select>
            </div>
          </div>
        </div>

        <!-- اطلاعات مالی -->
        <div>
          <h4 class="text-md font-medium text-gray-700 mb-3">اطلاعات مالی</h4>
          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">نوع ارز</label>
              <select 
                v-model="invoice.currency"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="IRR">ریال ایران</option>
                <option value="USD">دلار آمریکا</option>
                <option value="EUR">یورو</option>
                <option value="GBP">پوند انگلیس</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">نرخ تبدیل</label>
              <input 
                v-model="invoice.exchangeRate"
                type="number" 
                step="0.01"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="1.00"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">مبلغ کل (ریال)</label>
              <input 
                v-model="invoice.totalAmount"
                type="number" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="0"
                readonly
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-600 mb-2">مبلغ پرداخت شده</label>
              <input 
                v-model="invoice.paidAmount"
                type="number" 
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                placeholder="0"
              />
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- جزئیات فاکتور -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-green-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"></path>
        </svg>
        جزئیات فاکتور
      </h3>
      
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                شرح
              </th>
              <th class="px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                تعداد
              </th>
              <th class="px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                قیمت واحد
              </th>
              <th class="px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                تخفیف
              </th>
              <th class="px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                مالیات
              </th>
              <th class="px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                جمع
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="(item, index) in invoice.items" :key="index" class="hover:bg-gray-50">
              <td class="px-4 py-4">
                <input 
                  v-model="item.description"
                  type="text" 
                  class="w-full px-2 py-1 border border-gray-300 rounded text-sm focus:ring-1 focus:ring-blue-500 focus:border-blue-500"
                  placeholder="شرح کالا/خدمات"
                />
              </td>
              <td class="px-4 py-4">
                <input 
                  v-model="item.quantity"
                  type="number" 
                  class="w-full px-2 py-1 border border-gray-300 rounded text-sm focus:ring-1 focus:ring-blue-500 focus:border-blue-500"
                  placeholder="1"
                  @input="calculateItemTotal(index)"
                />
              </td>
              <td class="px-4 py-4">
                <input 
                  v-model="item.unitPrice"
                  type="number" 
                  class="w-full px-2 py-1 border border-gray-300 rounded text-sm focus:ring-1 focus:ring-blue-500 focus:border-blue-500"
                  placeholder="0"
                  @input="calculateItemTotal(index)"
                />
              </td>
              <td class="px-4 py-4">
                <input 
                  v-model="item.discount"
                  type="number" 
                  class="w-full px-2 py-1 border border-gray-300 rounded text-sm focus:ring-1 focus:ring-blue-500 focus:border-blue-500"
                  placeholder="0"
                  @input="calculateItemTotal(index)"
                />
              </td>
              <td class="px-4 py-4">
                <input 
                  v-model="item.tax"
                  type="number" 
                  class="w-full px-2 py-1 border border-gray-300 rounded text-sm focus:ring-1 focus:ring-blue-500 focus:border-blue-500"
                  placeholder="0"
                  @input="calculateItemTotal(index)"
                />
              </td>
              <td class="px-4 py-4">
                <div class="text-sm font-medium text-gray-900">
                  {{ formatPrice(item.total) }} تومان
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- دکمه افزودن آیتم -->
      <div class="mt-4">
        <button 
          class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors flex items-center text-sm"
          @click="addInvoiceItem"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
          </svg>
          افزودن آیتم
        </button>
      </div>
    </div>

    <!-- خلاصه مالی فاکتور -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-purple-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 7h6m0 10v-3m-3 3h.01M9 17h.01M9 14h.01M12 14h.01M15 11h.01M12 11h.01M9 11h.01M7 21h10a2 2 0 002-2V5a2 2 0 00-2-2H7a2 2 0 00-2 2v14a2 2 0 002 2z"></path>
        </svg>
        خلاصه مالی
      </h3>
      
      <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
        <!-- محاسبات -->
        <div class="space-y-3">
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">جمع کل:</span>
            <span class="font-medium">{{ formatPrice(subtotal) }} تومان</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">تخفیف کل:</span>
            <span class="font-medium text-green-600">-{{ formatPrice(totalDiscount) }} تومان</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">مالیات کل:</span>
            <span class="font-medium">{{ formatPrice(totalTax) }} تومان</span>
          </div>
          <div class="border-t pt-2 flex justify-between text-lg font-bold">
            <span>مبلغ نهایی:</span>
            <span class="text-blue-600">{{ formatPrice(totalAmount) }} تومان</span>
          </div>
        </div>

        <!-- وضعیت پرداخت -->
        <div class="space-y-3">
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">مبلغ پرداخت شده:</span>
            <span class="font-medium">{{ formatPrice(invoice.paidAmount) }} تومان</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">مبلغ باقی‌مانده:</span>
            <span class="font-medium text-red-600">{{ formatPrice(remainingAmount) }} تومان</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">درصد پیشرفت:</span>
            <span class="font-medium">{{ paymentProgress }}%</span>
          </div>
          <div class="w-full bg-gray-200 rounded-full h-2">
            <div 
              class="bg-blue-600 h-2 rounded-full transition-all duration-300"
              :style="{ width: paymentProgress + '%' }"
            ></div>
          </div>
        </div>
      </div>
    </div>

    <!-- یادداشت فاکتور -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
      <h3 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
        <svg class="w-5 h-5 text-orange-600 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
        </svg>
        یادداشت و شرایط
      </h3>
      
      <div class="space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-600 mb-2">یادداشت فاکتور</label>
          <textarea 
            v-model="invoice.notes"
            rows="3"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            placeholder="یادداشت‌های مربوط به فاکتور..."
          ></textarea>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-600 mb-2">شرایط پرداخت</label>
          <textarea 
            v-model="invoice.paymentTerms"
            rows="2"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            placeholder="شرایط پرداخت..."
          ></textarea>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'

// تعریف props
const props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({
      invoiceNumber: '',
      invoiceDate: '',
      dueDate: '',
      status: 'draft',
      currency: 'IRR',
      exchangeRate: 1,
      totalAmount: 0,
      paidAmount: 0,
      notes: '',
      paymentTerms: '',
      items: []
    })
  }
})

// تعریف emit
const emit = defineEmits(['update:modelValue'])

// متغیر محلی
const invoice = ref({ ...props.modelValue })

// محاسبات
const subtotal = computed(() => {
  return invoice.value.items.reduce((sum, item) => sum + (item.unitPrice * item.quantity), 0)
})

const totalDiscount = computed(() => {
  return invoice.value.items.reduce((sum, item) => sum + (item.discount || 0), 0)
})

const totalTax = computed(() => {
  return invoice.value.items.reduce((sum, item) => sum + (item.tax || 0), 0)
})

const totalAmount = computed(() => {
  return subtotal.value - totalDiscount.value + totalTax.value
})

const remainingAmount = computed(() => {
  return totalAmount.value - invoice.value.paidAmount
})

const paymentProgress = computed(() => {
  if (totalAmount.value === 0) return 0
  return Math.round((invoice.value.paidAmount / totalAmount.value) * 100)
})

// توابع
const calculateItemTotal = (index) => {
  const item = invoice.value.items[index]
  const subtotal = item.unitPrice * item.quantity
  const discount = item.discount || 0
  const tax = item.tax || 0
  item.total = subtotal - discount + tax
}

const addInvoiceItem = () => {
  invoice.value.items.push({
    description: '',
    quantity: 1,
    unitPrice: 0,
    discount: 0,
    tax: 0,
    total: 0
  })
}

// فرمت قیمت
const formatPrice = (price) => {
  if (!price) return '0'
  return new Intl.NumberFormat('fa-IR').format(price)
}

// نظارت بر تغییرات
watch(invoice, (newValue) => {
  emit('update:modelValue', newValue)
}, { deep: true })

// به‌روزرسانی مبلغ کل
watch(totalAmount, (newValue) => {
  invoice.value.totalAmount = newValue
})
</script> 
