<template>
  <div class="min-h-screen p-6">
    <div class="max-w-7xl mx-auto space-y-6">
      <!-- هدر صفحه -->
      <div class="rounded-lg p-6 bg-white shadow-sm">
        <div>
          <h1 class="text-2xl font-bold text-gray-900">مدیریت درگاه های پرداخت</h1>
          <p class="text-gray-600 mt-1">تنظیم و مدیریت درگاه‌های پرداخت آنلاین</p>
        </div>
      </div>

      <!-- تب‌های اصلی -->
      <div class="rounded-xl">
        <div class="border-b border-gray-200 bg-white">
          <nav class="flex justify-start space-x-4 space-x-reverse px-6 overflow-x-auto whitespace-nowrap scrollbar-thin scrollbar-thumb-gray-300">
            <button 
              v-for="(tab, index) in tabs" 
              :key="tab.id"
              @click="activeTab = tab.id"
              :class="[
                'py-4 px-1 border-b-2 font-medium text-sm',
                activeTab === tab.id
                  ? 'border-blue-500 text-blue-600'
                  : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300',
                index === tabs.length - 1 ? 'mr-8' : ''
              ]"
            >
              {{ tab.name }}
            </button>
          </nav>
        </div>
        
        <div class="p-6">
          <!-- تب داشبورد -->
          <div v-if="activeTab === 'dashboard'" class="space-y-6">
            <GatewayStats />
            <GatewayList />
          </div>
          

          
          <!-- تب گزارشات -->
          <div v-if="activeTab === 'reports'">
            <PaymentReports />
          </div>
          
          <!-- تب تراکنش‌ها -->
          <div v-if="activeTab === 'transactions'">
            <GatewayTransactions />
          </div>
          
          <!-- تب لاگ‌ها -->
          <div v-if="activeTab === 'logs'">
            <LogContainer />
          </div>
          
          <!-- تب تنظیمات -->
          <div v-if="activeTab === 'settings'">
            <GlobalSettings />
          </div>
        </div>
      </div>

      <!-- مودال‌ها -->
      <GatewayForm v-if="showAddGatewayModal" @close="showAddGatewayModal = false" />
      <GatewayDetails v-if="selectedGateway" :gateway="selectedGateway" @close="selectedGateway = null" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

// Type declaration for Nuxt 4 auto-imported functions
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
declare const useHead: (head: { title?: string }) => void

// آیکون‌ها به صورت مستقیم استفاده می‌شوند

// Import کامپوننت‌های درگاه‌های پرداخت
import GatewayDetails from './components/GatewayDetails.vue';
import GatewayForm from './components/GatewayForm.vue';
import GatewayList from './components/GatewayList.vue';
import GatewayStats from './components/GatewayStats.vue';
import GatewayTransactions from './components/GatewayTransactions.vue';
import GlobalSettings from './components/GlobalSettings.vue';
import LogContainer from './components/LogContainer.vue';
import PaymentReports from './components/PaymentReports.vue';


// متغیرهای reactive
const showAddGatewayModal = ref(false)
const selectedGateway = ref(null)
const activeTab = ref('dashboard')

// تب‌های اصلی
const tabs = [
  { id: 'dashboard', name: 'داشبورد' },
  { id: 'transactions', name: 'تراکنش‌ها' },
  { id: 'reports', name: 'گزارشات پرداخت' },
  { id: 'logs', name: 'گزارش لاگ‌ها' },
  { id: 'settings', name: 'تنظیمات کلی' }
]

// تعریف متا برای صفحه
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// عنوان صفحه
useHead({
  title: 'مدیریت درگاه های پرداخت - پنل مدیریت'
})
</script>

<!--
  صفحه اصلی مدیریت درگاه های پرداخت
  شامل:
  1. آمار کلی درگاه‌ها
  2. لیست درگاه‌های پرداخت
  3. تنظیمات کلی درگاه‌ها
  4. طراحی مدرن و کاملاً ریسپانسیو
  توضیحات کامل به فارسی در کد
--> 
