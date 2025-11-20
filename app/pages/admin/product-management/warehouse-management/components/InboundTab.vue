<template>
  <div class="space-y-6">
    <!-- هدر -->
    <div class="flex flex-col lg:flex-row lg:items-center lg:justify-between gap-6">
      <div>
        <h2 class="text-2xl font-bold text-gray-900">ورودی کالا</h2>
        <p class="text-gray-600">ثبت ورودی کالاهای جدید به انبار</p>
      </div>
      
      <button 
        class="bg-gradient-to-r from-green-600 to-green-700 text-white px-6 py-2 rounded-xl hover:from-green-700 hover:to-green-800 transition-all duration-200 shadow-lg hover:shadow-xl flex items-center gap-2"
        @click="showForm = !showForm"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
        </svg>
        {{ showForm ? 'بستن فرم' : 'ثبت ورودی جدید' }}
      </button>
    </div>

    <!-- فرم ثبت ورودی -->
    <div v-show="showForm" class="bg-white rounded-2xl shadow-lg p-6 border border-gray-200">
      <h3 class="text-lg font-bold text-gray-900 mb-4">اطلاعات ورودی کالا</h3>
      
      <form class="space-y-4" @submit.prevent="submitInbound">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <!-- شماره فاکتور -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">شماره فاکتور</label>
            <input 
              v-model="form.invoiceNumber"
              type="text" 
              required
              class="w-full bg-white border border-gray-300 rounded-xl px-4 py-2 text-gray-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent"
              placeholder="مثال: INV-2024-001"
            >
          </div>
          
          <!-- تأمین‌کننده -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تأمین‌کننده</label>
            <select 
              v-model="form.supplier"
              required
              class="w-full bg-white border border-gray-300 rounded-xl px-4 py-2 text-gray-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent"
            >
              <option value="">انتخاب تأمین‌کننده</option>
              <option value="شرکت الکترونیک تهران">شرکت الکترونیک تهران</option>
              <option value="کارخانه پوشاک اصفهان">کارخانه پوشاک اصفهان</option>
              <option value="نشر کتاب‌های آموزشی">نشر کتاب‌های آموزشی</option>
            </select>
          </div>
          
          <!-- تاریخ ورود -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ ورود</label>
            <input 
              v-model="form.entryDate"
              type="date" 
              required
              class="w-full bg-white border border-gray-300 rounded-xl px-4 py-2 text-gray-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent"
            >
          </div>
          
          <!-- نوع ورودی -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نوع ورودی</label>
            <select 
              v-model="form.entryType"
              required
              class="w-full bg-white border border-gray-300 rounded-xl px-4 py-2 text-gray-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent"
            >
              <option value="">انتخاب نوع</option>
              <option value="خرید">خرید</option>
              <option value="مرجوعی">مرجوعی</option>
              <option value="انتقال از انبار دیگر">انتقال از انبار دیگر</option>
              <option value="سایر">سایر</option>
            </select>
          </div>
        </div>
        
        <!-- کالاها -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">کالاها</label>
          <div class="space-y-3">
            <div v-for="(item, index) in form.items" :key="index" class="flex items-center gap-3 p-3 bg-gray-50 rounded-xl">
              <div class="flex-1">
                <select 
                  v-model="item.productId"
                  required
                  class="w-full bg-white border border-gray-300 rounded-lg px-3 py-2 text-gray-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent"
                >
                  <option value="">انتخاب کالا</option>
                  <option value="1">لپ‌تاپ Dell XPS 13</option>
                  <option value="2">کتاب React.js در عمل</option>
                  <option value="3">تی‌شرت نخی مردانه</option>
                  <option value="4">ماوس بی‌سیم Logitech</option>
                </select>
              </div>
              <div class="w-32">
                <input 
                  v-model.number="item.quantity"
                  type="number" 
                  required
                  min="1"
                  class="w-full bg-white border border-gray-300 rounded-lg px-3 py-2 text-gray-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent"
                  placeholder="تعداد"
                >
              </div>
              <div class="w-32">
                <input 
                  v-model.number="item.unitPrice"
                  type="number" 
                  required
                  min="0"
                  step="0.01"
                  class="w-full bg-white border border-gray-300 rounded-lg px-3 py-2 text-gray-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent"
                  placeholder="قیمت واحد"
                >
              </div>
              <button 
                type="button"
                class="p-2 text-red-600 hover:bg-red-50 rounded-lg transition-colors"
                title="حذف کالا"
                @click="removeItem(index)"
              >
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                </svg>
              </button>
            </div>
          </div>
          
          <button 
            type="button"
            class="mt-3 text-green-600 hover:text-green-700 font-medium flex items-center gap-2"
            @click="addItem"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
            </svg>
            افزودن کالای دیگر
          </button>
        </div>
        
        <!-- یادداشت -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">یادداشت</label>
          <textarea 
            v-model="form.notes"
            rows="3"
            class="w-full bg-white border border-gray-300 rounded-xl px-4 py-2 text-gray-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent"
            placeholder="توضیحات اضافی..."
          ></textarea>
        </div>
        
        <!-- دکمه‌های عملیات -->
        <div class="flex gap-3 pt-4">
          <button 
            type="submit"
            class="bg-gradient-to-r from-green-600 to-green-700 text-white px-8 py-3 rounded-xl hover:from-green-700 hover:to-green-800 transition-all duration-200 shadow-lg hover:shadow-xl flex items-center gap-2"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
            </svg>
            ثبت ورودی
          </button>
          
          <button 
            type="button"
            class="bg-gray-500 text-white px-6 py-3 rounded-xl hover:bg-gray-600 transition-all duration-200"
            @click="resetForm"
          >
            پاک کردن
          </button>
        </div>
      </form>
    </div>

    <!-- لیست ورودی‌های اخیر -->
    <div class="bg-white rounded-2xl shadow-lg overflow-hidden">
      <div class="px-6 py-4 border-b border-gray-200">
        <h3 class="text-lg font-bold text-gray-900">ورودی‌های اخیر</h3>
      </div>
      
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-4 text-right text-sm font-medium text-gray-900">شماره فاکتور</th>
              <th class="px-6 py-4 text-right text-sm font-medium text-gray-900">تأمین‌کننده</th>
              <th class="px-6 py-4 text-right text-sm font-medium text-gray-900">تاریخ ورود</th>
              <th class="px-6 py-4 text-right text-sm font-medium text-gray-900">نوع ورودی</th>
              <th class="px-6 py-4 text-right text-sm font-medium text-gray-900">تعداد کالا</th>
              <th class="px-6 py-4 text-right text-sm font-medium text-gray-900">مبلغ کل</th>
              <th class="px-6 py-4 text-right text-sm font-medium text-gray-900">عملیات</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="inbound in recentInbounds" :key="inbound.id" class="hover:bg-gray-50 transition-colors">
              <td class="px-6 py-4">
                <span class="font-medium text-gray-900">{{ inbound.invoiceNumber }}</span>
              </td>
              <td class="px-6 py-4">
                <div class="text-sm text-gray-900">{{ inbound.supplier }}</div>
              </td>
              <td class="px-6 py-4">
                <div class="text-sm text-gray-900">{{ formatDate(inbound.entryDate) }}</div>
              </td>
              <td class="px-6 py-4">
                <span 
                  :class="[
                    'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium',
                    getEntryTypeColor(inbound.entryType)
                  ]"
                >
                  {{ inbound.entryType }}
                </span>
              </td>
              <td class="px-6 py-4">
                <div class="text-sm text-gray-900">{{ inbound.itemCount }}</div>
              </td>
              <td class="px-6 py-4">
                <div class="text-sm font-medium text-gray-900">{{ formatPrice(inbound.totalAmount) }}</div>
              </td>
              <td class="px-6 py-4">
                <div class="flex items-center gap-2">
                  <button class="p-1.5 text-blue-600 hover:bg-blue-50 rounded-lg transition-colors" title="مشاهده">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                    </svg>
                  </button>
                  <button class="p-1.5 text-green-600 hover:bg-green-50 rounded-lg transition-colors" title="ویرایش">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useNotifier } from '~/composables/useNotifier';

const notifier = useNotifier()

// Props
interface Props {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  warehouse: any
}
defineProps<Props>()

// نمایش فرم
const showForm = ref(false)

// فرم ورودی
const form = ref({
  invoiceNumber: '',
  supplier: '',
  entryDate: new Date().toISOString().split('T')[0],
  entryType: '',
  items: [{ productId: '', quantity: 1, unitPrice: 0 }],
  notes: ''
})

// ورودی‌های اخیر
const recentInbounds = ref([
  {
    id: 1,
    invoiceNumber: 'INV-2024-001',
    supplier: 'شرکت الکترونیک تهران',
    entryDate: '2024-01-15',
    entryType: 'خرید',
    itemCount: 3,
    totalAmount: 12500000
  },
  {
    id: 2,
    invoiceNumber: 'INV-2024-002',
    supplier: 'کارخانه پوشاک اصفهان',
    entryDate: '2024-01-14',
    entryType: 'خرید',
    itemCount: 5,
    totalAmount: 8500000
  },
  {
    id: 3,
    invoiceNumber: 'INV-2024-003',
    supplier: 'انبار مشهد',
    entryDate: '2024-01-13',
    entryType: 'انتقال از انبار دیگر',
    itemCount: 2,
    totalAmount: 0
  }
])

// توابع
const addItem = () => {
  form.value.items.push({ productId: '', quantity: 1, unitPrice: 0 })
}

const removeItem = (index: number) => {
  if (form.value.items.length > 1) {
    form.value.items.splice(index, 1)
  }
}

const submitInbound = () => {
  // اینجا منطق ثبت ورودی قرار می‌گیرد
  notifier.success('ورودی کالا با موفقیت ثبت شد!')
  resetForm()
  showForm.value = false
}

const resetForm = () => {
  form.value = {
    invoiceNumber: '',
    supplier: '',
    entryDate: new Date().toISOString().split('T')[0],
    entryType: '',
    items: [{ productId: '', quantity: 1, unitPrice: 0 }],
    notes: ''
  }
}

const getEntryTypeColor = (type: string) => {
  switch (type) {
    case 'خرید': return 'bg-blue-100 text-blue-800'
    case 'مرجوعی': return 'bg-green-100 text-green-800'
    case 'انتقال از انبار دیگر': return 'bg-purple-100 text-purple-800'
    default: return 'bg-gray-100 text-gray-800'
  }
}

const formatDate = (dateString: string) => {
  const date = new Date(dateString)
  return date.toLocaleDateString('fa-IR')
}

const formatPrice = (price: number) => {
  return new Intl.NumberFormat('fa-IR').format(price) + ' تومان'
}
</script>
