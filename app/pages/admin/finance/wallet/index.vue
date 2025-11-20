<template>
  <div class="max-w-full mx-auto space-y-6">
    <!-- هدر صفحه -->
    <div class="p-6">
      <div class="flex flex-col lg:flex-row lg:items-center lg:justify-between gap-6">
        <div>
          <h1 class="text-2xl font-bold text-gray-900">مدیریت کیف پول</h1>
          <p class="text-gray-600 mt-1">مدیریت کامل کیف پول کاربران و تراکنش‌های مالی</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors" @click="showNewOperationModal = true">
            <i class="fas fa-plus ml-2"></i>
            عملیات جدید
          </button>
          <button class="bg-green-600 text-white px-4 py-2 rounded-lg hover:bg-green-700 transition-colors">
            <i class="fas fa-download ml-2"></i>
            گزارش
          </button>
        </div>
      </div>
    </div>

    <!-- منوی تب‌ها -->
    <div>
      <div class="border-b border-gray-200">
        <nav class="flex space-x-6 space-x-reverse overflow-x-auto px-6">
          <button
            v-for="tab in tabs"
            :key="tab.id"
            class="flex items-center space-x-2 space-x-reverse px-4 py-4 transition-colors whitespace-nowrap border-b-2"
            :class="[
              activeTab === tab.id
                ? 'text-blue-600 border-blue-500'
                : 'text-gray-600 hover:text-gray-900 border-transparent'
            ]"
            @click="activeTab = tab.id"
          >
            <i :class="tab.icon + ' text-lg'"></i>
            <span class="font-medium">{{ tab.title }}</span>
          </button>
        </nav>
      </div>

      <!-- محتوای تب‌ها -->
      <div class="p-6 overflow-x-auto">
        <!-- مدیریت موجودی و تراز -->
        <div v-if="activeTab === 'balance'" class="space-y-6 min-w-0">
          <WalletBalanceManagement />
        </div>

        <!-- مدیریت تراکنش‌ها -->
        <div v-if="activeTab === 'transactions'" class="space-y-6 min-w-0">
          <TransactionManagement />
        </div>

        <!-- مدیریت برداشت -->
        <div v-if="activeTab === 'withdrawals'" class="space-y-6 min-w-0">
          <WithdrawalManagement />
        </div>

        <!-- مدیریت شارژ -->
        <div v-if="activeTab === 'charges'" class="space-y-6 min-w-0">
          <ChargeManagement />
        </div>

        <!-- مدیریت کارت‌های بانکی -->
        <div v-if="activeTab === 'bankcards'" class="space-y-6 min-w-0">
          <BankCardManagement />
        </div>

        <!-- گزارش‌ها و تحلیل -->
        <div v-if="activeTab === 'reports'" class="space-y-6 min-w-0">
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <div class="min-w-0">
              <h3 class="text-lg font-semibold mb-4">گزارش‌های مالی</h3>
              <FinancialReports />
            </div>
            <div class="min-w-0">
              <h3 class="text-lg font-semibold mb-4">داشبورد تحلیلی</h3>
              <AnalyticsDashboard />
            </div>
          </div>
          <div class="mt-6 min-w-0">
            <h3 class="text-lg font-semibold mb-4">گزارش‌های پیشرفته</h3>
            <AdvancedReports />
          </div>
          <div class="mt-6 min-w-0">
            <h3 class="text-lg font-semibold mb-4">تحلیل‌های مالی</h3>
            <FinancialAnalytics />
          </div>
        </div>

        <!-- اتصالات و API -->
        <div v-if="activeTab === 'integrations'" class="space-y-6 min-w-0">
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <div class="min-w-0">
              <h3 class="text-lg font-semibold mb-4">اتصال حسابداری</h3>
              <AccountingIntegration />
            </div>
            <div class="min-w-0">
              <h3 class="text-lg font-semibold mb-4">اتصال سیستم‌های خارجی</h3>
              <ExternalIntegrations />
            </div>
          </div>
          <div class="mt-6 min-w-0">
            <h3 class="text-lg font-semibold mb-4">API و وب‌هوک</h3>
            <ApiWebhookManagement />
          </div>
        </div>

        <!-- مدیریت و کنترل -->
        <div v-if="activeTab === 'management'" class="space-y-6 min-w-0">
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <div class="min-w-0">
              <h3 class="text-lg font-semibold mb-4">مدیریت کاربران</h3>
              <UserManagement />
            </div>
            <div class="min-w-0">
              <h3 class="text-lg font-semibold mb-4">تنظیمات کیف پول</h3>
              <WalletSettings />
            </div>
          </div>
          <div class="mt-6 min-w-0">
            <h3 class="text-lg font-semibold mb-4">کنترل‌های داخلی</h3>
            <InternalControls />
          </div>
          <div class="mt-6 min-w-0">
            <h3 class="text-lg font-semibold mb-4">مدیریت مالیات</h3>
            <TaxManagement />
          </div>
          <div class="mt-6 min-w-0">
            <h3 class="text-lg font-semibold mb-4">اعلان‌ها</h3>
            <NotificationManagement />
          </div>
          <div class="mt-6 min-w-0">
            <h3 class="text-lg font-semibold mb-4">مستندات و پشتیبانی</h3>
            <DocumentationSupport />
          </div>
        </div>

        <!-- امنیت و نظارت -->
        <div v-if="activeTab === 'security'" class="space-y-6 min-w-0">
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <div class="min-w-0">
              <h3 class="text-lg font-semibold mb-4">امنیت و لاگ</h3>
              <SecurityLogs />
            </div>
            <div class="min-w-0">
              <h3 class="text-lg font-semibold mb-4">احراز هویت</h3>
              <AuthenticationSecurity />
            </div>
          </div>
          <div class="mt-6 min-w-0">
            <h3 class="text-lg font-semibold mb-4">پشتیبان‌گیری</h3>
            <BackupManagement />
          </div>
          <div class="mt-6 min-w-0">
            <h3 class="text-lg font-semibold mb-4">مدیریت خطاها</h3>
            <ErrorManagement />
          </div>
          <div class="mt-6 min-w-0">
            <h3 class="text-lg font-semibold mb-4">خطاهای حسابداری</h3>
            <AccountingErrorManagement />
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- مودال عملیات جدید -->
  <Teleport to="body">
    <div v-if="showNewOperationModal" class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-40">
      <div class="rounded-lg shadow-lg w-full max-w-2xl p-6 relative bg-white mx-4">
        <button class="absolute left-4 top-6 text-gray-400 hover:text-gray-600 text-2xl" @click="showNewOperationModal = false">
          <i class="fas fa-times"></i>
        </button>
        <ManualWalletManagement />
      </div>
    </div>
  </Teleport>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
</script>

<script setup lang="ts">
// صفحه مدیریت کیف پول
import { ref } from 'vue'
import AccountingErrorManagement from './components/AccountingErrorManagement.vue'
import AccountingIntegration from './components/AccountingIntegration.vue'
import AdvancedReports from './components/AdvancedReports.vue'
import AnalyticsDashboard from './components/AnalyticsDashboard.vue'
import ApiWebhookManagement from './components/ApiWebhookManagement.vue'
import AuthenticationSecurity from './components/AuthenticationSecurity.vue'
import BackupManagement from './components/BackupManagement.vue'
import BankCardManagement from './components/BankCardManagement.vue'
import ChargeManagement from './components/ChargeManagement.vue'
import DocumentationSupport from './components/DocumentationSupport.vue'
import ErrorManagement from './components/ErrorManagement.vue'
import ExternalIntegrations from './components/ExternalIntegrations.vue'
import FinancialAnalytics from './components/FinancialAnalytics.vue'
import FinancialReports from './components/FinancialReports.vue'
import InternalControls from './components/InternalControls.vue'
import ManualWalletManagement from './components/ManualWalletManagement.vue'
import NotificationManagement from './components/NotificationManagement.vue'
import SecurityLogs from './components/SecurityLogs.vue'
import TaxManagement from './components/TaxManagement.vue'
import TransactionManagement from './components/TransactionManagement.vue'
import UserManagement from './components/UserManagement.vue'
import WalletBalanceManagement from './components/WalletBalanceManagement.vue'
import WalletSettings from './components/WalletSettings.vue'
import WithdrawalManagement from './components/WithdrawalManagement.vue'

definePageMeta({ layout: 'admin-main' })

// تعریف تب‌های اصلی - با توجه به ماهیت کانتینرها دسته‌بندی شده
const tabs = [
  // بخش اصلی کیف پول
  { id: 'balance', title: 'موجودی و تراز', icon: 'fas fa-wallet' },
  { id: 'transactions', title: 'تراکنش‌ها', icon: 'fas fa-exchange-alt' },
  { id: 'withdrawals', title: 'برداشت', icon: 'fas fa-arrow-up' },
  { id: 'charges', title: 'شارژ', icon: 'fas fa-arrow-down' },
  { id: 'bankcards', title: 'کارت‌های بانکی', icon: 'fas fa-credit-card' },
  
  // بخش گزارش‌گیری و تحلیل
  { id: 'reports', title: 'گزارش‌ها و تحلیل', icon: 'fas fa-chart-bar' },
  
  // بخش اتصالات و یکپارچه‌سازی
  { id: 'integrations', title: 'اتصالات و API', icon: 'fas fa-plug' },
  
  // بخش مدیریت و کنترل
  { id: 'management', title: 'مدیریت و کنترل', icon: 'fas fa-cog' },
  
  // بخش امنیت و نظارت
  { id: 'security', title: 'امنیت و نظارت', icon: 'fas fa-shield-alt' }
]

// تب فعال
const activeTab = ref('balance')

// وضعیت نمایش مودال عملیات جدید
const showNewOperationModal = ref(false)
</script>

<style scoped>
/* استایل‌های سفارشی برای صفحه کیف پول */
.space-x-reverse > * + * {
  margin-right: 0.75rem;
  margin-left: 0;
}

/* انیمیشن hover */
.transition-colors {
  transition: all 0.2s ease-in-out;
}
</style> 
