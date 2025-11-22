<template>
  <ClientOnly>
    <!-- هدر صفحه -->
    <OrderHeader 
      :order-number="orderData.orderNumber"
      @save-order="saveOrder"
      @print-order="printOrder"
    />

    <!-- محتوای اصلی -->
    <div class="w-full max-w-7xl mx-auto">
      <!-- تب‌ها -->
      <div class="rounded-lg shadow-md border border-gray-200">
        <!-- هدر تب‌ها -->
        <div class="border-b border-gray-200">
          <div class="relative">
            <!-- دکمه‌های اسکرول -->
            <button 
              v-if="canScrollLeft"
              class="absolute left-0 top-0 z-10 h-full w-8 bg-gradient-to-r from-white to-transparent flex items-center justify-center text-gray-500 hover:text-gray-700"
              @click="scrollTabs('left')"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"></path>
              </svg>
            </button>
            
            <button 
              v-if="canScrollRight"
              class="absolute right-0 top-0 z-10 h-full w-8 bg-gradient-to-l from-white to-transparent flex items-center justify-center text-gray-500 hover:text-gray-700"
              @click="scrollTabs('right')"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
              </svg>
            </button>
            
            <!-- نوار تب‌ها -->
            <nav 
              ref="tabsContainer"
              class="flex space-x-6 space-x-reverse px-6 overflow-x-auto" 
              aria-label="Tabs"
              @scroll="checkScrollButtons"
            >
              <button
                v-for="tab in tabs"
                :key="tab.id"
                :class="[
                  activeTab === tab.id
                    ? 'border-blue-500 text-blue-600'
                    : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300',
                  'whitespace-nowrap py-4 px-2 border-b-2 font-medium text-sm transition-colors flex-shrink-0'
                ]"
                @click="activeTab = tab.id"
              >
                <div class="flex items-center">
                  <component :is="tab.icon" class="w-4 h-4 ml-2" />
                  <span class="hidden sm:inline">{{ tab.name }}</span>
                  <span class="sm:hidden">{{ tab.shortName || tab.name }}</span>
                </div>
              </button>
            </nav>
          </div>
        </div>

        <!-- محتوای تب‌ها -->
        <div class="px-4 py-4">
          <!-- تب اطلاعات مشتری -->
          <div v-if="activeTab === 'customer'" class="space-y-6 max-h-[70vh] overflow-y-auto">
            <CustomerInfo 
              v-model="orderData.customer"
            />
          </div>

          <!-- تب محصولات سفارش -->
          <div v-if="activeTab === 'items'" class="space-y-6 max-h-[70vh] overflow-y-auto">
            <OrderItems 
              v-model="orderData.items"
              :shipping-cost="orderData.shippingCost"
              @update:shipping-cost="orderData.shippingCost = $event"
            />
          </div>

          <!-- تب وضعیت سفارش -->
          <div v-if="activeTab === 'status'" class="space-y-6 max-h-[70vh] overflow-y-auto">
            <OrderStatus 
              v-model="orderData.status"
              :order-id="orderId"
              @status-updated="handleStatusUpdate"
            />
          </div>

          <!-- تب حمل و نقل -->
          <div v-if="activeTab === 'shipping'" class="space-y-6 max-h-[70vh] overflow-y-auto">
            <ShippingInfo 
              v-model="orderData.shipping"
            />
          </div>

          <!-- تب فاکتور -->
          <div v-if="activeTab === 'invoice'" class="space-y-6 max-h-[70vh] overflow-y-auto">
            <Invoice 
              v-model="orderData.invoice"
            />
          </div>

          <!-- تب مالیات و حسابداری -->
          <div v-if="activeTab === 'tax'" class="space-y-6 max-h-[70vh] overflow-y-auto">
            <TaxAccounting 
              v-model="orderData.taxAccounting"
            />
          </div>

          <!-- تب کیف پول و کارت هدیه -->
          <div v-if="activeTab === 'wallet'" class="space-y-6 max-h-[70vh] overflow-y-auto">
            <WalletGiftCard 
              v-model="orderData.walletGiftCard"
            />
          </div>

          <!-- تب درگاه‌های پرداخت -->
          <div v-if="activeTab === 'gateways'" class="space-y-6 max-h-[70vh] overflow-y-auto">
            <PaymentGateways 
              v-model="orderData.paymentGateways"
            />
          </div>

          <!-- تب لجستیک پیشرفته -->
          <div v-if="activeTab === 'logistics'" class="space-y-6 max-h-[70vh] overflow-y-auto">
            <AdvancedLogistics 
              v-model="orderData.advancedLogistics"
            />
          </div>

          <!-- تب پشتیبانی و ارتباطات -->
          <div v-if="activeTab === 'support'" class="space-y-6 max-h-[70vh] overflow-y-auto">
            <SupportCommunication 
              v-model="orderData.supportCommunication"
            />
          </div>

          <!-- تب آمار و گزارشات -->
          <div v-if="activeTab === 'statistics'" class="space-y-6 max-h-[70vh] overflow-y-auto">
            <StatisticsReports 
              v-model="orderData.statisticsReports"
            />
          </div>

          <!-- تب خلاصه سفارش -->
          <div v-if="activeTab === 'summary'" class="space-y-6">
            <OrderSummary 
              :order-data="orderData"
            />
          </div>
        </div>
      </div>
    </div>
  </ClientOnly>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
</script>

<script setup>
import { nextTick, onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'

definePageMeta({ layout: 'admin-main', middleware: 'admin' });

// Import کامپوننت‌ها
import AdvancedLogistics from './components/AdvancedLogistics.vue'
import CustomerInfo from './components/CustomerInfo.vue'
import Invoice from './components/Invoice.vue'
import OrderHeader from './components/OrderHeader.vue'
import OrderItems from './components/OrderItems.vue'
import OrderStatus from './components/OrderStatus.vue'
import OrderSummary from './components/OrderSummary.vue'
import PaymentGateways from './components/PaymentGateways.vue'
import ShippingInfo from './components/ShippingInfo.vue'
import StatisticsReports from './components/StatisticsReports.vue'
import SupportCommunication from './components/SupportCommunication.vue'
import TaxAccounting from './components/TaxAccounting.vue'
import WalletGiftCard from './components/WalletGiftCard.vue'

definePageMeta({
  layout: 'admin-main'
})

const route = useRoute()
const orderId = route.params.id

// تب فعال
const activeTab = ref('summary')

// متغیرهای اسکرول
const tabsContainer = ref(null)
const canScrollLeft = ref(false)
const canScrollRight = ref(false)

// تعریف تب‌ها
const tabs = [
  {
    id: 'summary',
    name: 'خلاصه سفارش',
    shortName: 'خلاصه',
    icon: 'SummaryIcon'
  },
  {
    id: 'customer',
    name: 'اطلاعات مشتری',
    shortName: 'مشتری',
    icon: 'CustomerIcon'
  },
  {
    id: 'items',
    name: 'محصولات سفارش',
    shortName: 'محصولات',
    icon: 'ItemsIcon'
  },
  {
    id: 'status',
    name: 'وضعیت سفارش',
    shortName: 'وضعیت',
    icon: 'StatusIcon'
  },
  {
    id: 'shipping',
    name: 'حمل و نقل',
    shortName: 'حمل',
    icon: 'ShippingIcon'
  },
  {
    id: 'invoice',
    name: 'فاکتور',
    shortName: 'فاکتور',
    icon: 'InvoiceIcon'
  },
  {
    id: 'tax',
    name: 'مالیات و حسابداری',
    shortName: 'مالیات',
    icon: 'TaxIcon'
  },
  {
    id: 'wallet',
    name: 'کیف پول و کارت هدیه',
    shortName: 'کیف پول',
    icon: 'WalletIcon'
  },
  {
    id: 'gateways',
    name: 'درگاه‌های پرداخت',
    shortName: 'درگاه‌ها',
    icon: 'GatewayIcon'
  },
  {
    id: 'logistics',
    name: 'لجستیک پیشرفته',
    shortName: 'لجستیک',
    icon: 'LogisticsIcon'
  },
  {
    id: 'support',
    name: 'پشتیبانی و ارتباطات',
    shortName: 'پشتیبانی',
    icon: 'SupportIcon'
  },
  {
    id: 'statistics',
    name: 'آمار و گزارشات',
    shortName: 'آمار',
    icon: 'StatisticsIcon'
  }
]

// آیکون‌های تب‌ها
const _CustomerIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
  </svg>`
}

const _ItemsIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path>
  </svg>`
}

const _StatusIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
  </svg>`
}

const _ShippingIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 8h14M5 8a2 2 0 110-4h14a2 2 0 110 4M5 8v10a2 2 0 002 2h10a2 2 0 002-2V8m-9 4h4"></path>
  </svg>`
}

const _InvoiceIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
  </svg>`
}

const _TaxIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 7h6m0 10v-3m-3 3h.01M9 17h.01M9 14h.01M12 14h.01M15 11h.01M12 11h.01M9 11h.01M7 21h10a2 2 0 002-2V5a2 2 0 00-2-2H7a2 2 0 00-2 2v14a2 2 0 002 2z"></path>
  </svg>`
}

const _WalletIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"></path>
  </svg>`
}

const _GatewayIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"></path>
  </svg>`
}

const _LogisticsIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7h12m0 0l-4-4m4 4l-4 4m0 6H4m0 0l4 4m-4-4l4-4"></path>
  </svg>`
}

const _SupportIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
  </svg>`
}

const _StatisticsIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
  </svg>`
}

const _SummaryIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
  </svg>`
}

// داده‌های سفارش
const orderData = ref({
  id: orderId,
  orderNumber: 'ORD-2024-001',
  createdAt: '2024-01-15T10:30:00',
  customer: {
    name: 'علی احمدی',
    phone: '09123456789',
    email: 'ali.ahmadi@example.com',
    address: {
      province: 'tehran',
      city: 'تهران',
      fullAddress: 'خیابان ولیعصر، پلاک 123',
      postalCode: '1234567890'
    },
    joinDate: '2023-01-15',
    totalOrders: 5,
    rating: 4.5
  },
  items: [
    {
      name: 'لپ‌تاپ اپل مک‌بوک پرو',
      sku: 'MBP-001',
      category: 'لپ‌تاپ',
      image: '/default-product.svg',
      unitPrice: 2500000,
      quantity: 1,
      discount: 0,
      totalPrice: 2500000
    },
    {
      name: 'گوشی سامسونگ گلکسی S24',
      sku: 'S24-001',
      category: 'گوشی موبایل',
      image: '/default-product.svg',
      unitPrice: 1800000,
      quantity: 1,
      discount: 100000,
      totalPrice: 1700000
    }
  ],
  shippingCost: 50000,
  status: {
    status: 'processing_queue',
    priority: 'normal',
    statusReason: '',
    paymentMethod: 'online',
    paymentStatus: 'awaiting_payment',
    paymentDate: '2024-01-15T10:35:00',
    returnReason: '',
    returnStatus: '',
    returnRequestDate: '',
    refundAmount: 0,
    refundMethod: 'original_method',
    refundDate: '',
    internalNotes: 'سفارش VIP - اولویت بالا',
    customerNotes: 'سفارش شما در حال پردازش است',
    notes: 'سفارش در حال پردازش است',
    history: [
      {
        status: 'processing_queue',
        date: '2024-01-15T10:30:00',
        user: 'سیستم',
        note: 'سفارش ثبت شد'
      },
      {
        status: 'paid',
        date: '2024-01-15T10:35:00',
        user: 'سیستم',
        note: 'پرداخت تایید شد'
      },
      {
        status: 'processing',
        date: '2024-01-16T09:00:00',
        user: 'مدیر فروش',
        note: 'سفارش در حال آماده‌سازی'
      }
    ]
  },
  shipping: {
    method: 'express',
    company: 'iran_post',
    cost: 50000,
    trackingNumber: 'TRK123456789',
    shippingDate: '2024-01-20T10:00:00',
    estimatedDelivery: '2024-01-25T18:00:00',
    actualDelivery: '',
    deliveryStatus: 'shipped',
    deliveryNotes: 'ارسال با پست پیشتاز - تحویل در محل',
    weight: 2.5,
    dimensions: {
      length: 30,
      width: 20,
      height: 15
    },
    packageCount: 1,
    status: 'shipped',
    lastUpdate: '2024-01-20T10:00:00',
    notes: 'ارسال با پست پیشتاز',
    history: [
      {
        status: 'packaged',
        date: '2024-01-19T14:00:00',
        location: 'انبار مرکزی',
        description: 'بسته‌بندی تکمیل شد'
      },
      {
        status: 'shipped',
        date: '2024-01-20T10:00:00',
        location: 'دفتر پست',
        description: 'تحویل به پست'
      }
    ]
  },
      invoice: {
      invoiceNumber: 'INV-2024-001',
      invoiceDate: '2024-01-15',
      dueDate: '2024-01-30',
      status: 'paid',
      currency: 'IRR',
      exchangeRate: 1,
      totalAmount: 4250000,
      paidAmount: 4250000,
      notes: 'فاکتور پرداخت شده',
      paymentTerms: 'پرداخت نقدی',
      items: [
        {
          description: 'لپ‌تاپ اپل مک‌بوک پرو',
          quantity: 1,
          unitPrice: 2500000,
          discount: 0,
          tax: 0,
          total: 2500000
        },
        {
          description: 'گوشی سامسونگ گلکسی S24',
          quantity: 1,
          unitPrice: 1800000,
          discount: 100000,
          tax: 0,
          total: 1700000
        },
        {
          description: 'هزینه ارسال',
          quantity: 1,
          unitPrice: 50000,
          discount: 0,
          tax: 0,
          total: 50000
        }
      ]
    },
    walletGiftCard: {
      currentBalance: 500000,
      usedAmount: 200000,
      walletStatus: 'active',
      recentTransactions: [
        {
          type: 'شارژ کیف پول',
          date: '2024-01-10T14:30:00',
          amount: 500000,
          description: 'شارژ از طریق درگاه بانکی'
        },
        {
          type: 'استفاده از کیف پول',
          date: '2024-01-15T10:30:00',
          amount: -200000,
          description: 'استفاده در سفارش #12345'
        }
      ],
      usedGiftCards: [
        {
          code: 'GIFT-2024-001',
          usedAmount: 100000,
          status: 'active',
          expiryDate: '2024-12-31',
          remainingBalance: 400000
        }
      ],
      transactionHistory: [
        {
          type: 'wallet_use',
          date: '2024-01-15T10:30:00',
          user: 'سیستم',
          amount: -200000,
          note: 'استفاده از کیف پول در سفارش'
        },
        {
          type: 'gift_card_use',
          date: '2024-01-15T10:31:00',
          user: 'سیستم',
          amount: -100000,
          note: 'استفاده از کارت هدیه GIFT-2024-001'
        }
      ]
    },
    taxAccounting: {
      regionalTaxRate: 9,
      taxableAmount: 4100000,
      calculatedTax: 369000,
      taxExemption: '',
      taxInvoiceNumber: 'TAX-2024-001',
      taxInvoiceDate: '2024-01-15',
      taxInvoiceStatus: 'issued',
      economicCode: '123456789',
      salesAccount: '4001',
      taxAccount: '2201',
      bankAccount: '1011',
      costCenter: 'CC001',
      accountingDocumentNumber: 'DOC-2024-001',
      accountingDate: '2024-01-15',
      accountingStatus: 'posted',
      accountingNotes: 'ثبت خودکار سفارش',
      history: [
        {
          type: 'tax_calculation',
          date: '2024-01-15T10:30:00',
          user: 'سیستم',
          note: 'محاسبه مالیات اولیه'
        },
        {
          type: 'tax_invoice_issued',
          date: '2024-01-15T10:35:00',
          user: 'سیستم',
          note: 'صدور فاکتور مالیاتی'
        },
        {
          type: 'accounting_entry',
          date: '2024-01-15T10:40:00',
          user: 'سیستم',
          note: 'ثبت تراکنش حسابداری'
        }
      ]
    },
    paymentGateways: {
      primaryGateway: 'zarinpal',
      transactionId: 'ZP123456789',
      referenceNumber: 'REF987654321',
      transactionStatus: 'completed',
      transactionAmount: 4250000,
      gatewayFee: 85000,
      transactionDate: '2024-01-15T10:35:00',
      errorCode: '',
      alternativeGateways: [
        {
          name: 'mellat',
          status: 'available',
          priority: 2,
          notes: 'درگاه جایگزین اول'
        },
        {
          name: 'parsian',
          status: 'available',
          priority: 3,
          notes: 'درگاه جایگزین دوم'
        }
      ],
      authenticationType: 'sms',
      authenticationStatus: 'verified',
      authenticationDate: '2024-01-15T10:30:00',
      securityLevel: 'high',
      securityFilterStatus: 'active',
      clientIP: '192.168.1.100',
      transactionLogs: [
        {
          action: 'شروع تراکنش',
          date: '2024-01-15T10:30:00',
          gateway: 'zarinpal',
          status: 'success',
          details: 'ارسال درخواست به درگاه'
        },
        {
          action: 'احراز هویت',
          date: '2024-01-15T10:32:00',
          gateway: 'zarinpal',
          status: 'success',
          details: 'تایید کد پیامکی'
        },
        {
          action: 'پرداخت موفق',
          date: '2024-01-15T10:35:00',
          gateway: 'zarinpal',
          status: 'success',
          details: 'تراکنش با موفقیت انجام شد'
        }
      ]
    },
    advancedLogistics: {
      driverName: 'علی احمدی',
      driverPhone: '09123456789',
      vehiclePlate: '12-345-67',
      driverStatus: 'busy',
      vehicleType: 'car',
      vehicleCapacity: 100,
      vehicleStatus: 'active',
      lastServiceDate: '2024-01-10',
      currentLatitude: 35.6892,
      currentLongitude: 51.3890,
      currentAddress: 'تهران، خیابان ولیعصر، پلاک 123',
      lastLocationUpdate: '2024-01-15T10:30:00',
      estimatedArrival: '2024-01-15T14:00:00',
      remainingDistance: 5.2,
      currentSpeed: 45,
      trafficStatus: 'moderate',
      selectedTimeSlot: 'afternoon',
      timeSlotStart: '12:00',
      timeSlotEnd: '16:00',
      timeSlotStatus: 'booked',
      deliveryType: 'door',
      deliveryPriority: 'normal',
      specialDeliveryFee: 0,
      deliveryNotes: 'تحویل درب منزل، تماس قبل از تحویل',
      deliveryZone: 'zone1',
      trafficRestriction: 'none',
      restrictionHours: '',
      restrictionStatus: 'inactive',
      distributionCenter: 'center1',
      optimizedRoute: 'fastest',
      maxStops: 3,
                logisticsNotes: 'مسیر بهینه انتخاب شده',
          logisticsHistory: [
            {
              action: 'تخصیص راننده',
              date: '2024-01-15T08:00:00',
              location: 'مرکز توزیع',
              status: 'completed',
              details: 'راننده علی احمدی تخصیص داده شد'
            },
            {
              action: 'شروع تحویل',
              date: '2024-01-15T10:00:00',
              location: 'مرکز توزیع',
              status: 'started',
              details: 'شروع فرآیند تحویل سفارش'
            },
            {
              action: 'در مسیر',
              date: '2024-01-15T10:30:00',
              location: 'تهران، خیابان ولیعصر',
              status: 'in_progress',
              details: 'راننده در مسیر تحویل است'
            }
          ]
        },
        supportCommunication: {
          tickets: [
            {
              ticketNumber: 'TKT-2024-001',
              subject: 'سوال در مورد زمان تحویل',
              status: 'in_progress',
              priority: 'normal',
              category: 'delivery',
              description: 'مشتری در مورد زمان دقیق تحویل سوال دارد'
            }
          ],
          chatInfo: {
            chatId: 'CHAT-2024-001',
            status: 'active',
            agent: 'سارا احمدی',
            startDate: '2024-01-15T10:00:00',
            messageCount: 8,
            duration: 25,
            satisfactionRating: '5',
            notes: 'چت فعال با مشتری'
          },
          chatMessages: [
            {
              sender: 'customer',
              content: 'سلام، سوالی در مورد سفارشم دارم',
              timestamp: '2024-01-15T10:00:00'
            },
            {
              sender: 'support',
              content: 'سلام، چطور می‌تونم کمکتون کنم؟',
              timestamp: '2024-01-15T10:01:00'
            },
            {
              sender: 'customer',
              content: 'زمان تحویل سفارشم کی هست؟',
              timestamp: '2024-01-15T10:02:00'
            },
            {
              sender: 'support',
              content: 'سفارش شما امروز بین ساعت 12 تا 16 تحویل داده خواهد شد',
              timestamp: '2024-01-15T10:03:00'
            }
          ],
          phoneCalls: [
            {
              phoneNumber: '09123456789',
              callType: 'incoming',
              status: 'answered',
              callDate: '2024-01-15T09:30:00',
              duration: 180,
              notes: 'تماس مشتری برای پیگیری سفارش'
            }
          ]
        },
        statisticsReports: {
          orderStats: {
            totalOrders: 15,
            completedOrders: 12,
            pendingOrders: 2,
            cancelledOrders: 1,
            averageOrderValue: 2850000,
            totalRevenue: 42750000,
            conversionRate: 85.7,
            returnRate: 6.7
          },
          customerStats: {
            totalCustomers: 8,
            newCustomers: 2,
            repeatCustomers: 6,
            averageCustomerValue: 5343750,
            customerSatisfaction: 4.6,
            customerLifetimeValue: 8500000,
            churnRate: 12.5
          },
          productStats: {
            totalProducts: 25,
            topSellingProduct: 'لپ‌تاپ اپل مک‌بوک پرو',
            topSellingCategory: 'لپ‌تاپ',
            averageProductRating: 4.4,
            productsInStock: 20,
            lowStockProducts: 3,
            outOfStockProducts: 2
          },
          timeAnalysis: {
            peakOrderTime: '14:00-16:00',
            peakOrderDay: 'دوشنبه',
            averageProcessingTime: '2.5 ساعت',
            averageDeliveryTime: '3.2 روز',
            seasonalTrend: 'افزایشی',
            monthlyGrowth: 15.3
          },
          geographicAnalysis: {
            topCity: 'تهران',
            topProvince: 'تهران',
            deliveryZones: 5,
            averageDeliveryDistance: 12.5,
            internationalOrders: 0,
            localOrders: 15
          },
          financialReports: {
            totalSales: 42750000,
            totalCost: 32000000,
            grossProfit: 10750000,
            profitMargin: 25.1,
            averageOrderValue: 2850000,
            paymentMethods: {
              online: 80,
              cash: 20
            }
          },
          reportOperations: {
            lastReportGenerated: '2024-01-15T10:00:00',
            reportFrequency: 'daily',
            nextReportDate: '2024-01-16T10:00:00',
            reportFormat: 'PDF',
            autoExport: true,
            exportEmail: 'reports@example.com'
          }
        }
})

// توابع اسکرول تب‌ها
const scrollTabs = (direction) => {
  if (!tabsContainer.value) return
  
  const container = tabsContainer.value
  const scrollAmount = 200
  
  if (direction === 'left') {
    container.scrollLeft -= scrollAmount
  } else {
    container.scrollLeft += scrollAmount
  }
  
  // بررسی وضعیت دکمه‌ها بعد از اسکرول
  setTimeout(checkScrollButtons, 100)
}

const checkScrollButtons = () => {
  if (!tabsContainer.value) return
  
  const container = tabsContainer.value
  canScrollLeft.value = container.scrollLeft > 0
  canScrollRight.value = container.scrollLeft < (container.scrollWidth - container.clientWidth)
}

// دریافت اطلاعات سفارش از API
const fetchOrderData = async () => {
  try {
    // دریافت اطلاعات سفارش از API
    const response = await $fetch(`/api/admin/orders/${orderId}`)
    
    if (response && response.success && response.data) {
      const orderInfo = response.data
      
      // به‌روزرسانی وضعیت سفارش با داده‌های واقعی
      orderData.value.status.status = orderInfo.status || 'processing_queue'
      orderData.value.status.paymentStatus = orderInfo.paymentStatus || 'awaiting_payment'
      orderData.value.status.paymentMethod = orderInfo.paymentMethod || 'online'
      
      // به‌روزرسانی اطلاعات مشتری
      if (orderInfo.customerName) {
        orderData.value.customer.name = orderInfo.customerName
      }
      if (orderInfo.customerPhone) {
        orderData.value.customer.phone = orderInfo.customerPhone
      }
      
      // به‌روزرسانی مبلغ کل
      if (orderInfo.totalAmount) {
        orderData.value.items.forEach(item => {
          item.totalPrice = orderInfo.totalAmount / orderData.value.items.length
        })
      }
    }
  } catch {
    // console.error('خطا در دریافت اطلاعات سفارش:', error)
    // در صورت خطا، از داده‌های پیش‌فرض استفاده می‌شود
  }
}

// ذخیره تغییرات سفارش
const saveOrder = async () => {
  try {
    // به‌روزرسانی وضعیت سفارش
    if (orderData.value.status && orderData.value.status.status) {
      const updateData = {
        status: orderData.value.status.status,
        notes: orderData.value.status.internalNotes || ''
      }
      
      const _response = await $fetch(`/api/admin/orders/${orderId}/status`, {
        method: 'PUT',
        body: updateData
      })
    }
    
    // نمایش پیام موفقیت
    alert('تغییرات با موفقیت ذخیره شد')
  } catch {
    // console.error('خطا در ذخیره تغییرات:', error)
    // console.error('جزئیات خطا:', error.response?.data || error.message)
    alert('خطا در ذخیره تغییرات')
  }
}

// مدیریت به‌روزرسانی وضعیت از کامپوننت OrderStatus
const handleStatusUpdate = (updateData) => {
  // به‌روزرسانی داده‌های محلی
  orderData.value.status.status = updateData.status
  orderData.value.status.paymentStatus = updateData.paymentStatus
  
  // نمایش پیام موفقیت
  alert('وضعیت سفارش با موفقیت به‌روزرسانی شد')
}

// چاپ سفارش
const printOrder = () => {
  // در اینجا منطق چاپ سفارش پیاده‌سازی می‌شود
  window.print()
}

onMounted(() => {
  fetchOrderData()
  
  // بررسی وضعیت دکمه‌های اسکرول بعد از لود صفحه
  nextTick(() => {
    checkScrollButtons()
    
    // بررسی مجدد در صورت تغییر اندازه پنجره
    window.addEventListener('resize', checkScrollButtons)
  })
})
  </script>

<style scoped>
/* استایل‌های سفارشی برای scrollbar */
nav::-webkit-scrollbar {
  height: 6px;
}

nav::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 3px;
}

nav::-webkit-scrollbar-thumb {
  background: #c1c1c1;
  border-radius: 3px;
}

nav::-webkit-scrollbar-thumb:hover {
  background: #a8a8a8;
}

/* برای Firefox */
nav {
  scrollbar-width: thin;
  scrollbar-color: #c1c1c1 #f1f1f1;
}
</style>
