<template>
  <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
    <!-- هدر -->
    <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-6 mb-6">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">مدیریت فاکتورها و رسیدها</h3>
        <p class="text-sm text-gray-600">تولید، چاپ و مدیریت فاکتورهای مالیاتی</p>
      </div>
      
      <!-- دکمه‌های عملیاتی -->
      <div class="flex gap-2">
        <button 
          class="inline-flex items-center gap-2 px-4 py-2 bg-green-600 hover:bg-green-700 text-white rounded-lg transition-colors duration-200"
          @click="showCreateInvoiceModal = true"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
          </svg>
          فاکتور جدید
        </button>
        <button 
          class="inline-flex items-center gap-2 px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg transition-colors duration-200"
          @click="bulkPrintInvoices"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 17h2a2 2 0 002-2v-4a2 2 0 00-2-2H5a2 2 0 00-2 2v4a2 2 0 002 2h2m2 4h6a2 2 0 002-2v-4a2 2 0 00-2-2H9a2 2 0 00-2 2v4a2 2 0 002 2zm8-12V5a2 2 0 00-2-2H9a2 2 0 00-2 2v4h10z"></path>
          </svg>
          چاپ دسته‌ای
        </button>
      </div>
    </div>

    <!-- آمار فاکتورها -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
      <div class="bg-blue-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-blue-600">{{ invoiceStats.total }}</div>
            <div class="text-sm text-blue-700">کل فاکتورها</div>
          </div>
          <div class="w-10 h-10 bg-blue-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-green-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-green-600">{{ invoiceStats.paid }}</div>
            <div class="text-sm text-green-700">پرداخت شده</div>
          </div>
          <div class="w-10 h-10 bg-green-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-yellow-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-yellow-600">{{ invoiceStats.pending }}</div>
            <div class="text-sm text-yellow-700">در انتظار</div>
          </div>
          <div class="w-10 h-10 bg-yellow-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-red-50 rounded-lg p-6">
        <div class="flex items-center justify-between">
          <div>
            <div class="text-2xl font-bold text-red-600">{{ invoiceStats.overdue }}</div>
            <div class="text-sm text-red-700">معوق</div>
          </div>
          <div class="w-10 h-10 bg-red-500 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- فیلترهای فاکتور -->
    <div class="bg-gray-50 rounded-lg p-6 mb-6">
      <div class="grid grid-cols-1 md:grid-cols-5 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
          <select 
            v-model="invoiceFilters.status"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            @change="filterInvoices"
          >
            <option value="">همه وضعیت‌ها</option>
            <option value="paid">پرداخت شده</option>
            <option value="pending">در انتظار</option>
            <option value="overdue">معوق</option>
            <option value="cancelled">لغو شده</option>
          </select>
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">از تاریخ</label>
          <input 
            v-model="invoiceFilters.dateFrom"
            type="date"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            @change="filterInvoices"
          />
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">تا تاریخ</label>
          <input 
            v-model="invoiceFilters.dateTo"
            type="date"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            @change="filterInvoices"
          />
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">مشتری</label>
          <input 
            v-model="invoiceFilters.customer"
            type="text"
            placeholder="جستجو در مشتریان..."
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            @input="filterInvoices"
          />
        </div>
        
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">شماره فاکتور</label>
          <input 
            v-model="invoiceFilters.invoiceNumber"
            type="text"
            placeholder="شماره فاکتور..."
            class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            @input="filterInvoices"
          />
        </div>
      </div>
    </div>

    <!-- جدول فاکتورها -->
    <div class="overflow-x-auto">
      <table class="w-full text-sm">
        <thead>
          <tr class="border-b border-gray-200 bg-gray-50">
            <th class="text-right py-3 px-4 font-medium text-gray-600">
              <input 
                v-model="selectAll"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                @change="toggleSelectAll"
              />
            </th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">شماره فاکتور</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">تاریخ</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">مشتری</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">مبلغ</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">وضعیت</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">تاریخ سررسید</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">عملیات</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="invoice in filteredInvoices" :key="invoice.id" class="border-b border-gray-100 hover:bg-gray-50">
            <td class="py-3 px-4">
              <input 
                v-model="selectedInvoices"
                :value="invoice.id"
                type="checkbox"
                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
              />
            </td>
            <td class="py-3 px-4 text-gray-900 font-medium">{{ invoice.invoiceNumber }}</td>
            <td class="py-3 px-4 text-gray-900">{{ formatDate(invoice.date) }}</td>
            <td class="py-3 px-4 text-gray-900">{{ invoice.customer }}</td>
            <td class="py-3 px-4 text-gray-900">{{ formatCurrency(invoice.amount) }}</td>
            <td class="py-3 px-4">
              <span 
                :class="[
                  'px-2 py-1 rounded-full text-xs font-medium',
                  getStatusClass(invoice.status)
                ]"
              >
                {{ getStatusLabel(invoice.status) }}
              </span>
            </td>
            <td class="py-3 px-4 text-gray-900">{{ formatDate(invoice.dueDate) }}</td>
            <td class="py-3 px-4">
              <div class="flex items-center gap-2">
                <button 
                  class="p-1 text-blue-600 hover:text-blue-800"
                  title="مشاهده"
                  @click="viewInvoice(invoice)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                  </svg>
                </button>
                <button 
                  class="p-1 text-green-600 hover:text-green-800"
                  title="چاپ"
                  @click="printInvoice(invoice)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 17h2a2 2 0 002-2v-4a2 2 0 00-2-2H5a2 2 0 00-2 2v4a2 2 0 002 2h2m2 4h6a2 2 0 002-2v-4a2 2 0 00-2-2H9a2 2 0 00-2 2v4a2 2 0 002 2zm8-12V5a2 2 0 00-2-2H9a2 2 0 00-2 2v4h10z"></path>
                  </svg>
                </button>
                <button 
                  class="p-1 text-yellow-600 hover:text-yellow-800"
                  title="ویرایش"
                  @click="editInvoice(invoice)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
                  </svg>
                </button>
                <button 
                  v-if="canDeleteTaxInvoice"
                  class="p-1 text-red-600 hover:text-red-800"
                  title="حذف"
                  @click="deleteInvoice(invoice)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                  </svg>
                </button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- صفحه‌بندی -->
    <div class="flex items-center justify-between mt-6">
      <div class="text-sm text-gray-700">
        نمایش {{ pagination.start }} تا {{ pagination.end }} از {{ pagination.total }} فاکتور
      </div>
      <div class="flex items-center gap-2">
        <button 
          :disabled="pagination.currentPage === 1"
          class="px-3 py-1 border border-gray-300 rounded-lg text-sm disabled:opacity-50 disabled:cursor-not-allowed"
          @click="previousPage"
        >
          قبلی
        </button>
        <span class="px-3 py-1 text-sm text-gray-700">
          صفحه {{ pagination.currentPage }} از {{ pagination.totalPages }}
        </span>
        <button 
          :disabled="pagination.currentPage === pagination.totalPages"
          class="px-3 py-1 border border-gray-300 rounded-lg text-sm disabled:opacity-50 disabled:cursor-not-allowed"
          @click="nextPage"
        >
          بعدی
        </button>
      </div>
    </div>

    <!-- مودال ایجاد فاکتور -->
    <div v-if="showCreateInvoiceModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-xl p-6 w-full max-w-4xl mx-4 max-h-[90vh] overflow-y-auto">
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-semibold text-gray-900">ایجاد فاکتور جدید</h3>
          <button class="text-gray-400 hover:text-gray-600" @click="showCreateInvoiceModal = false">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <!-- اطلاعات مشتری -->
          <div class="space-y-4">
            <h4 class="font-medium text-gray-900">اطلاعات مشتری</h4>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نام مشتری</label>
              <input 
                v-model="newInvoice.customerName"
                type="text"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">کد ملی/شناسه ملی</label>
              <input 
                v-model="newInvoice.customerId"
                type="text"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">آدرس</label>
              <textarea 
                v-model="newInvoice.address"
                rows="3"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              ></textarea>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">شماره تماس</label>
              <input 
                v-model="newInvoice.phone"
                type="tel"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
          </div>

          <!-- اطلاعات فاکتور -->
          <div class="space-y-4">
            <h4 class="font-medium text-gray-900">اطلاعات فاکتور</h4>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ فاکتور</label>
              <input 
                v-model="newInvoice.date"
                type="date"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ سررسید</label>
              <input 
                v-model="newInvoice.dueDate"
                type="date"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نوع فاکتور</label>
              <select 
                v-model="newInvoice.type"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="sale">فروش</option>
                <option value="service">خدمات</option>
                <option value="export">صادرات</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات</label>
              <textarea 
                v-model="newInvoice.description"
                rows="3"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              ></textarea>
            </div>
          </div>
        </div>

        <!-- آیتم‌های فاکتور -->
        <div class="mt-6">
          <h4 class="font-medium text-gray-900 mb-4">آیتم‌های فاکتور</h4>
          
          <div class="space-y-3">
            <div v-for="(item, index) in newInvoice.items" :key="index" class="grid grid-cols-1 md:grid-cols-5 gap-3">
              <div>
                <input 
                  v-model="item.description"
                  type="text"
                  placeholder="توضیحات"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm"
                />
              </div>
              <div>
                <input 
                  v-model.number="item.quantity"
                  type="number"
                  placeholder="تعداد"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm"
                />
              </div>
              <div>
                <input 
                  v-model.number="item.unitPrice"
                  type="number"
                  placeholder="قیمت واحد"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm"
                />
              </div>
              <div>
                <input 
                  v-model.number="item.taxRate"
                  type="number"
                  placeholder="نرخ مالیات %"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm"
                />
              </div>
              <div class="flex items-center gap-2">
                <span class="text-sm font-medium">{{ formatCurrency(item.total) }}</span>
                <button 
                  class="p-1 text-red-600 hover:text-red-800"
                  @click="removeInvoiceItem(index)"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                  </svg>
                </button>
              </div>
            </div>
            
            <button 
              class="w-full px-4 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200"
              @click="addInvoiceItem"
            >
              افزودن آیتم
            </button>
          </div>
        </div>

        <!-- خلاصه فاکتور -->
        <div class="mt-6 bg-gray-50 rounded-lg p-6">
          <div class="flex justify-between items-center">
            <div class="text-lg font-medium text-gray-900">مجموع: {{ formatCurrency(newInvoice.total) }}</div>
            <div class="flex gap-3">
              <button 
                class="px-4 py-2 bg-gray-200 text-gray-700 rounded-lg hover:bg-gray-300"
                @click="showCreateInvoiceModal = false"
              >
                انصراف
              </button>
              <button 
                class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700"
                @click="createInvoice"
              >
                ایجاد فاکتور
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { useAuth } from '~/composables/useAuth'

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { hasPermission } = useAuth()

// Computed برای چک کردن پرمیژن حذف
const canDeleteTaxInvoice = computed(() => hasPermission('tax_invoice.delete'))

// آمار فاکتورها
const invoiceStats = ref({
  total: 1250,
  paid: 890,
  pending: 280,
  overdue: 80
})

// فیلترهای فاکتور
const invoiceFilters = ref({
  status: '',
  dateFrom: '',
  dateTo: '',
  customer: '',
  invoiceNumber: ''
})

// صفحه‌بندی
const pagination = ref({
  currentPage: 1,
  totalPages: 50,
  total: 1250,
  start: 1,
  end: 25,
  perPage: 25
})

// انتخاب فاکتورها
const selectAll = ref(false)
const selectedInvoices = ref<number[]>([])

// مودال ایجاد فاکتور
const showCreateInvoiceModal = ref(false)

// فاکتور جدید
const newInvoice = ref({
  customerName: '',
  customerId: '',
  address: '',
  phone: '',
  date: '',
  dueDate: '',
  type: 'sale',
  description: '',
  items: [
    { description: '', quantity: 1, unitPrice: 0, taxRate: 9, total: 0 }
  ],
  total: 0
})

// فاکتورهای نمونه
const invoices = ref([
  {
    id: 1,
    invoiceNumber: 'INV-2024-001',
    date: '2024-01-15',
    customer: 'علی احمدی',
    amount: 1250000,
    status: 'paid',
    dueDate: '2024-02-15'
  },
  {
    id: 2,
    invoiceNumber: 'INV-2024-002',
    date: '2024-01-14',
    customer: 'شرکت فناوری',
    amount: 2500000,
    status: 'pending',
    dueDate: '2024-02-14'
  },
  {
    id: 3,
    invoiceNumber: 'INV-2024-003',
    date: '2024-01-13',
    customer: 'فاطمه محمدی',
    amount: 850000,
    status: 'overdue',
    dueDate: '2024-01-13'
  }
])

// فاکتورهای فیلتر شده
const filteredInvoices = computed(() => {
  return invoices.value.filter(invoice => {
    if (invoiceFilters.value.status && invoice.status !== invoiceFilters.value.status) return false
    if (invoiceFilters.value.customer && !invoice.customer.includes(invoiceFilters.value.customer)) return false
    if (invoiceFilters.value.invoiceNumber && !invoice.invoiceNumber.includes(invoiceFilters.value.invoiceNumber)) return false
    // TODO: فیلتر تاریخ
    return true
  })
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

// کلاس وضعیت
const getStatusClass = (status: string) => {
  const classes = {
    paid: 'bg-green-100 text-green-700',
    pending: 'bg-yellow-100 text-yellow-700',
    overdue: 'bg-red-100 text-red-700',
    cancelled: 'bg-gray-100 text-gray-700'
  }
  return classes[status] || 'bg-gray-100 text-gray-700'
}

// برچسب وضعیت
const getStatusLabel = (status: string) => {
  const labels = {
    paid: 'پرداخت شده',
    pending: 'در انتظار',
    overdue: 'معوق',
    cancelled: 'لغو شده'
  }
  return labels[status] || status
}

// فیلتر کردن فاکتورها
const filterInvoices = () => {
  // TODO: اعمال فیلترها
  console.log('فاکتورها فیلتر شدند:', invoiceFilters.value)
}

// انتخاب همه
const toggleSelectAll = () => {
  if (selectAll.value) {
    selectedInvoices.value = filteredInvoices.value.map(invoice => invoice.id)
  } else {
    selectedInvoices.value = []
  }
}

// مشاهده فاکتور
const viewInvoice = (invoice: any) => {
  // TODO: نمایش جزئیات فاکتور
  console.log('مشاهده فاکتور:', invoice)
}

// چاپ فاکتور
const printInvoice = (invoice: any) => {
  // TODO: چاپ فاکتور
  console.log('چاپ فاکتور:', invoice)
}

// ویرایش فاکتور
const editInvoice = (invoice: any) => {
  // TODO: ویرایش فاکتور
  console.log('ویرایش فاکتور:', invoice)
}

// حذف فاکتور
const deleteInvoice = (invoice: any) => {
  if (confirm('آیا از حذف این فاکتور اطمینان دارید؟')) {
    // TODO: حذف فاکتور
    console.log('حذف فاکتور:', invoice)
  }
}

// چاپ دسته‌ای
const bulkPrintInvoices = () => {
  if (selectedInvoices.value.length === 0) {
    alert('لطفاً فاکتورهایی را انتخاب کنید')
    return
  }
  // TODO: چاپ دسته‌ای
  console.log('چاپ دسته‌ای فاکتورها:', selectedInvoices.value)
}

// صفحه بعدی
const nextPage = () => {
  if (pagination.value.currentPage < pagination.value.totalPages) {
    pagination.value.currentPage++
    // TODO: بارگذاری فاکتورهای صفحه جدید
  }
}

// صفحه قبلی
const previousPage = () => {
  if (pagination.value.currentPage > 1) {
    pagination.value.currentPage--
    // TODO: بارگذاری فاکتورهای صفحه جدید
  }
}

// افزودن آیتم فاکتور
const addInvoiceItem = () => {
  newInvoice.value.items.push({
    description: '',
    quantity: 1,
    unitPrice: 0,
    taxRate: 9,
    total: 0
  })
}

// حذف آیتم فاکتور
const removeInvoiceItem = (index: number) => {
  newInvoice.value.items.splice(index, 1)
  calculateInvoiceTotal()
}

// محاسبه مجموع فاکتور
const calculateInvoiceTotal = () => {
  newInvoice.value.total = newInvoice.value.items.reduce((sum, item) => {
    const subtotal = item.quantity * item.unitPrice
    const tax = subtotal * (item.taxRate / 100)
    item.total = subtotal + tax
    return sum + item.total
  }, 0)
}

// ایجاد فاکتور
const createInvoice = async () => {
  try {
    // TODO: ارسال درخواست به API
    console.log('فاکتور جدید ایجاد شد:', newInvoice.value)
    showCreateInvoiceModal.value = false
    
    // ریست فرم
    newInvoice.value = {
      customerName: '',
      customerId: '',
      address: '',
      phone: '',
      date: '',
      dueDate: '',
      type: 'sale',
      description: '',
      items: [{ description: '', quantity: 1, unitPrice: 0, taxRate: 9, total: 0 }],
      total: 0
    }
  } catch (error) {
    console.error('خطا در ایجاد فاکتور:', error)
  }
}

// بارگذاری اولیه
onMounted(() => {
  // تنظیم تاریخ‌های پیش‌فرض
  const now = new Date()
  const firstDay = new Date(now.getFullYear(), now.getMonth(), 1)
  const lastDay = new Date(now.getFullYear(), now.getMonth() + 1, 0)
  
  invoiceFilters.value.dateFrom = firstDay.toISOString().split('T')[0]
  invoiceFilters.value.dateTo = lastDay.toISOString().split('T')[0]
  
  // TODO: بارگذاری فاکتورها از API
})
</script>

<!--
  کامپوننت مدیریت فاکتورها و رسیدها
  شامل:
  1. نمایش لیست فاکتورها
  2. فیلترهای پیشرفته
  3. ایجاد فاکتور جدید
  4. چاپ و ویرایش فاکتورها
  5. مدیریت وضعیت فاکتورها
  6. چاپ دسته‌ای
  7. طراحی مدرن و کاملاً ریسپانسیو
--> 
