<template>
  <div class="space-y-6">
    <!-- هدر صفحه -->
    <div class="bg-white rounded-lg shadow-sm p-6">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-2xl font-bold text-gray-900">مدیریت اتصال حسابداری</h1>
          <p class="text-gray-600 mt-1">مدیریت کامل اتصال به سیستم‌های حسابداری و همگام‌سازی داده‌ها</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
            <i class="fas fa-sync ml-2"></i>
            همگام‌سازی
          </button>
          <button class="bg-green-600 text-white px-4 py-2 rounded-lg hover:bg-green-700 transition-colors">
            <i class="fas fa-download ml-2"></i>
            گزارش
          </button>
        </div>
      </div>
    </div>

    <!-- منوی تب‌ها -->
    <div class="bg-white rounded-lg shadow-sm">
      <div class="border-b border-gray-200">
        <nav class="flex space-x-6 space-x-reverse overflow-x-auto px-6">
          <button
            v-for="tab in tabs"
            :key="tab.id"
            @click="activeTab = tab.id"
            class="flex items-center space-x-2 space-x-reverse px-4 py-4 transition-colors whitespace-nowrap border-b-2"
            :class="[
              activeTab === tab.id
                ? 'text-blue-600 border-blue-500'
                : 'text-gray-600 hover:text-gray-900 border-transparent'
            ]"
          >
            <i :class="tab.icon + ' text-lg'"></i>
            <span class="font-medium">{{ tab.title }}</span>
          </button>
        </nav>
      </div>

      <!-- محتوای تب‌ها -->
      <div class="p-6">
        <!-- داشبورد -->
        <div v-if="activeTab === 'dashboard'" class="space-y-6">
          <ConnectionDashboard />
        </div>

        <!-- اتصالات -->
        <div v-if="activeTab === 'connections'" class="space-y-6">
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <div>
              <h3 class="text-lg font-semibold mb-4">اتصالات حسابداری</h3>
              <AccountingConnections />
            </div>
            <div>
              <h3 class="text-lg font-semibold mb-4">نظارت بر اتصالات</h3>
              <ConnectionMonitoring />
            </div>
          </div>
        </div>

        <!-- همگام‌سازی -->
        <div v-if="activeTab === 'sync'" class="space-y-6">
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <div>
              <h3 class="text-lg font-semibold mb-4">همگام‌سازی داده‌ها</h3>
              <DataSynchronization />
            </div>
            <div>
              <h3 class="text-lg font-semibold mb-4">مدیریت داده‌ها</h3>
              <DataManagement />
            </div>
          </div>
        </div>

        <!-- نگاشت حساب‌ها -->
        <div v-if="activeTab === 'mapping'" class="space-y-6">
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <div>
              <h3 class="text-lg font-semibold mb-4">نگاشت حساب‌ها</h3>
              <AccountMapping />
            </div>
            <div>
              <h3 class="text-lg font-semibold mb-4">نگاشت خودکار</h3>
              <AutoAccountMapping />
            </div>
          </div>
          <div class="mt-6">
            <h3 class="text-lg font-semibold mb-4">مدیریت کدهای حساب</h3>
            <AccountCodesManagement />
          </div>
        </div>

        <!-- حساب‌های مختلف -->
        <div v-if="activeTab === 'accounts'" class="space-y-6">
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <div>
              <h3 class="text-lg font-semibold mb-4">حساب‌های بانکی</h3>
              <BankAccounts />
            </div>
            <div>
              <h3 class="text-lg font-semibold mb-4">حساب‌های فروش</h3>
              <SalesAccounts />
            </div>
          </div>
          <div class="mt-6">
            <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
              <div>
                <h3 class="text-lg font-semibold mb-4">حساب‌های خرید</h3>
                <PurchaseAccounts />
              </div>
              <div>
                <h3 class="text-lg font-semibold mb-4">حساب‌های مالیات</h3>
                <TaxAccounts />
              </div>
            </div>
          </div>
        </div>

        <!-- تراکنش‌ها -->
        <div v-if="activeTab === 'transactions'" class="space-y-6">
          <TransactionManagement />
        </div>

        <!-- گزارش‌گیری -->
        <div v-if="activeTab === 'reporting'" class="space-y-6">
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <div>
              <h3 class="text-lg font-semibold mb-4">گزارش‌گیری و نظارت</h3>
              <ReportingAndMonitoring />
            </div>
            <div>
              <h3 class="text-lg font-semibold mb-4">تست و اعتبارسنجی</h3>
              <TestingValidation />
            </div>
          </div>
        </div>

        <!-- تنظیمات -->
        <div v-if="activeTab === 'settings'" class="space-y-6">
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <div>
              <h3 class="text-lg font-semibold mb-4">تنظیمات پیشرفته</h3>
              <AdvancedSettings />
            </div>
            <div>
              <h3 class="text-lg font-semibold mb-4">رابط کاربری</h3>
              <UserInterface />
            </div>
          </div>
        </div>

        <!-- امنیت -->
        <div v-if="activeTab === 'security'" class="space-y-6">
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <div>
              <h3 class="text-lg font-semibold mb-4">مدیریت امنیت</h3>
              <SecurityManagement />
            </div>
            <div>
              <h3 class="text-lg font-semibold mb-4">مدیریت خطاها</h3>
              <ErrorManagement />
            </div>
          </div>
        </div>

        <!-- پشتیبانی -->
        <div v-if="activeTab === 'support'" class="space-y-6">
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <div>
              <h3 class="text-lg font-semibold mb-4">پشتیبانی و راهنما</h3>
              <SupportAndHelp />
            </div>
            <div>
              <h3 class="text-lg font-semibold mb-4">مدیریت نسخه‌ها</h3>
              <VersionManagement />
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// Import کامپوننت‌های حسابداری
import AccountCodesManagement from './AccountCodesManagement.vue'
import AccountMapping from './AccountMapping.vue'
import AccountingConnections from './AccountingConnections.vue'
import AdvancedSettings from './AdvancedSettings.vue'
import AutoAccountMapping from './AutoAccountMapping.vue'
import BankAccounts from './BankAccounts.vue'
import ConnectionMonitoring from './ConnectionMonitoring.vue'
import DataSynchronization from './DataSynchronization.vue'
import ErrorManagement from './ErrorManagement.vue'
import PurchaseAccounts from './PurchaseAccounts.vue'
import ReportingAndMonitoring from './ReportingAndMonitoring.vue'
import SalesAccounts from './SalesAccounts.vue'
import TaxAccounts from './TaxAccounts.vue'
import TransactionManagement from './TransactionManagement.vue'
import ConnectionDashboard from './components/DashboardStatistics.vue'
import SecurityManagement from './components/SecurityManagement.vue'
import SupportAndHelp from './components/SupportHelp.vue'
import DataManagement from './components/data/DataManagement.vue'
import TestingValidation from './components/testing/TestingValidation.vue'
import UserInterface from './components/ui/UserInterface.vue'
import VersionManagement from './components/version/VersionManagement.vue'

import { useAuth } from '~/composables/useAuth'

// تعریف definePageMeta برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void

// تعریف متا برای صفحه
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user, hasPermission } = useAuth()

// تعریف تب‌های اصلی - با توجه به ماهیت کانتینرها دسته‌بندی شده
const tabs = [
  // بخش اصلی
  { id: 'dashboard', title: 'داشبورد', icon: 'fas fa-tachometer-alt' },
  { id: 'connections', title: 'اتصالات', icon: 'fas fa-plug' },
  { id: 'sync', title: 'همگام‌سازی', icon: 'fas fa-sync' },
  { id: 'mapping', title: 'نگاشت حساب‌ها', icon: 'fas fa-map' },
  { id: 'accounts', title: 'حساب‌های مختلف', icon: 'fas fa-book' },
  
  // بخش تراکنش‌ها و گزارش‌گیری
  { id: 'transactions', title: 'تراکنش‌ها', icon: 'fas fa-exchange-alt' },
  { id: 'reporting', title: 'گزارش‌گیری', icon: 'fas fa-chart-bar' },
  
  // بخش تنظیمات و امنیت
  { id: 'settings', title: 'تنظیمات', icon: 'fas fa-cog' },
  { id: 'security', title: 'امنیت', icon: 'fas fa-shield-alt' },
  
  // بخش پشتیبانی
  { id: 'support', title: 'پشتیبانی', icon: 'fas fa-life-ring' }
]

const activeTab = ref('dashboard')
</script>

<style scoped>
/* استایل‌های سفارشی برای صفحه حسابداری */
.space-x-reverse > * + * {
  margin-right: 0.75rem;
  margin-left: 0;
}

/* انیمیشن hover */
.transition-colors {
  transition: all 0.2s ease-in-out;
}
</style> 
