<template>
  <div class="bg-white" dir="rtl">
    <!-- Header -->
    <div class="bg-white border-b border-gray-200">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center py-6">
          <div>
            <h1 class="text-2xl font-bold text-gray-900">سیستم نظرسنجی سفارشات</h1>
            <p class="text-gray-600 text-sm">مدیریت نظرسنجی‌ها، پیامک‌ها و گزارش‌ها</p>
          </div>
          <div class="flex items-center space-x-4 space-x-reverse">
            <button class="px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg text-sm transition-colors">
              راهنما
            </button>
            <button class="px-4 py-2 bg-green-600 hover:bg-green-700 text-white rounded-lg text-sm transition-colors">
              پشتیبان‌گیری
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Main Content with Sidebar -->
    <div class="flex">
      <!-- Sidebar Navigation -->
      <div class="w-64 bg-white shadow-lg border-l border-gray-200">
        <div class="px-4 py-4">
          <h3 class="text-lg font-semibold text-gray-900 mb-4">منوی نظرسنجی</h3>
          <nav class="space-y-2">
          <button
              v-for="(tab, index) in tabs" 
              :key="index"
            :class="[
                'w-full text-right px-4 py-3 rounded-lg transition-all duration-200 flex items-center justify-between',
              activeTab === tab.id
                  ? 'bg-blue-100 text-blue-700 border-r-4 border-blue-500' 
                  : 'text-gray-700 hover:bg-gray-100'
            ]"
            @click="activeTab = tab.id"
          >
              <span class="text-sm font-medium">{{ tab.name }}</span>
              <component :is="tab.icon" class="w-4 h-4" />
          </button>
        </nav>
      </div>
    </div>

      <!-- Main Content Area -->
      <div class="flex-1 max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        
      <!-- Dashboard Tab -->
      <div v-if="activeTab === 'dashboard'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
            <h2 class="text-xl font-bold text-gray-900 mb-4">داشبورد</h2>
        <StatsDashboard />
          </div>
      </div>

      <!-- SMS Management Tab -->
      <div v-if="activeTab === 'sms-management'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
            <h2 class="text-xl font-bold text-gray-900 mb-4">مدیریت پیامک</h2>
        <SMSManagement />
          </div>
      </div>

      <!-- Order List Tab -->
      <div v-if="activeTab === 'order-list'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
            <h2 class="text-xl font-bold text-gray-900 mb-4">لیست سفارشات</h2>
        <OrderListTable />
          </div>
      </div>

      <!-- Response Management Tab -->
      <div v-if="activeTab === 'response-management'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
            <h2 class="text-xl font-bold text-gray-900 mb-4">مدیریت پاسخ‌ها</h2>
        <OrderSurveyResponses />
          </div>
      </div>

      <!-- Advanced Reports Tab -->
      <div v-if="activeTab === 'advanced-reports'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
            <h2 class="text-xl font-bold text-gray-900 mb-4">گزارشات پیشرفته</h2>
        <AdvancedReports />
          </div>
      </div>

      <!-- Advanced Analytics Tab -->
      <div v-if="activeTab === 'advanced-analytics'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
            <h2 class="text-xl font-bold text-gray-900 mb-4">تحلیل پیشرفته</h2>
        <AdvancedAnalytics />
          </div>
      </div>

      <!-- SMS Settings Tab -->
      <div v-if="activeTab === 'sms-settings'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
            <h2 class="text-xl font-bold text-gray-900 mb-4">تنظیمات پیامک</h2>
        <SMSSettings />
          </div>
      </div>

      <!-- SMS Gateway Status Tab -->
      <div v-if="activeTab === 'sms-gateway-status'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
            <h2 class="text-xl font-bold text-gray-900 mb-4">وضعیت درگاه‌ها</h2>
            <SMSGatewayStatus :gateway="sampleGateway" @gateway-deleted="handleGatewayDeleted" />
          </div>
      </div>

      <!-- Bulk Order Manager Tab -->
      <div v-if="activeTab === 'bulk-order-manager'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
            <h2 class="text-xl font-bold text-gray-900 mb-4">مدیریت گروهی</h2>
        <BulkOrderManager />
          </div>
      </div>

      <!-- Notifications Settings Tab -->
      <div v-if="activeTab === 'notifications-settings'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
            <h2 class="text-xl font-bold text-gray-900 mb-4">تنظیمات اعلان</h2>
        <NotificationsSettings />
          </div>
      </div>

      <!-- Backup & Export Tab -->
      <div v-if="activeTab === 'backup-export'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
            <h2 class="text-xl font-bold text-gray-900 mb-4">پشتیبان‌گیری</h2>
        <BackupExport />
          </div>
      </div>

      <!-- Access Control Tab -->
      <div v-if="activeTab === 'access-control'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
            <h2 class="text-xl font-bold text-gray-900 mb-4">کنترل دسترسی</h2>
        <AccessControl />
          </div>
      </div>

      <!-- UI Personalization Tab -->
      <div v-if="activeTab === 'ui-personalization'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
            <h2 class="text-xl font-bold text-gray-900 mb-4">شخصی‌سازی</h2>
        <UIPersonalization />
          </div>
      </div>

      <!-- Survey Templates Tab -->
      <div v-if="activeTab === 'survey-templates'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 px-4 py-4">
            <h2 class="text-xl font-bold text-gray-900 mb-4">تمپلیت‌های نظرسنجی</h2>
        <SurveyTemplates />
          </div>
        </div>
        
        <!-- Default message when no section is selected -->
        <div v-if="!activeTab" class="flex items-center justify-center h-64">
          <div class="text-center">
            <svg class="w-16 h-16 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
            </svg>
            <h3 class="text-lg font-medium text-gray-900 mb-2">بخش مورد نظر را انتخاب کنید</h3>
            <p class="text-gray-500">برای مشاهده محتوای هر بخش، آن را از منوی سمت راست انتخاب کنید.</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string }) => void
</script>

<script setup lang="ts">
import { ref } from 'vue';
import SMSGatewayStatus from '~/components/admin/sms/SMSGatewayStatus.vue';
import AccessControl from './AccessControl.vue';
import AdvancedAnalytics from './AdvancedAnalytics.vue';
import AdvancedReports from './AdvancedReports.vue';
import BackupExport from './BackupExport.vue';
import BulkOrderManager from './BulkOrderManager.vue';
import NotificationsSettings from './NotificationsSettings.vue';
import OrderListTable from './OrderListTable.vue';
import OrderSurveyResponses from './OrderSurveyResponses.vue';
import SMSManagement from './SMSManagement.vue';
import SMSSettings from './SMSSettings.vue';
import StatsDashboard from './StatsDashboard.vue';
import SurveyTemplates from './SurveyTemplates.vue';
import UIPersonalization from './UIPersonalization.vue';

definePageMeta({
  layout: 'admin-main'
})

const activeTab = ref('dashboard')

// Sample data for SMS Gateway Status
const sampleGateway = ref({
  id: 1,
  name: 'درگاه نمونه',
  status: 'active' as const,
  balance: 1000,
  lastActivity: '2024-01-15T10:30:00Z'
})

const handleGatewayDeleted = (id: number) => {
}

// Tab icons components
const DashboardIcon = {
  template: `
    <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
    </svg>
  `
}

const SMSIcon = {
  template: `
    <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
    </svg>
  `
}

const OrderIcon = {
  template: `
    <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"></path>
    </svg>
  `
}

const ResponseIcon = {
  template: `
    <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-5 5v-5z"></path>
    </svg>
  `
}

const ReportsIcon = {
  template: `
    <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 17v-2m3 2v-4m3 4v-6m2 10H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
    </svg>
  `
}

const AnalyticsIcon = {
  template: `
    <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
    </svg>
  `
}

const SettingsIcon = {
  template: `
    <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"></path>
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
    </svg>
  `
}

const GatewayStatusIcon = {
  template: `
    <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
    </svg>
  `
}

const BulkIcon = {
  template: `
    <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
      
    </svg>
  `
}

const NotificationIcon = {
  template: `
    <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-5 5v-5zM4.19 4.19A2 2 0 004 5v12a2 2 0 002 2h12a2 2 0 002-2V5a2 2 0 00-2-2H6a2 2 0 00-1.81 1.19z"></path>
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01"></path>
    </svg>
  `
}

const BackupIcon = {
  template: `
    <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7H5a2 2 0 00-2 2v9a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-3m-1 4l-3 3m0 0l-3-3m3 3V4"></path>
    </svg>
  `
}

const SecurityIcon = {
  template: `
    <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"></path>
    </svg>
  `
}

const UIIcon = {
  template: `
    <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4zm0 0h12a2 2 0 002-2v-4a2 2 0 00-2-2h-2.343M11 7.343l1.657-1.657a2 2 0 012.828 0l2.829 2.829a2 2 0 010 2.828l-8.486 8.485M7 17h.01"></path>
    </svg>
  `
}

const TemplateIcon = {
  template: `
    <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 5a1 1 0 011-1h14a1 1 0 011 1v2a1 1 0 01-1 1H5a1 1 0 01-1-1V5zM4 13a1 1 0 011-1h6a1 1 0 011 1v6a1 1 0 01-1 1H5a1 1 0 01-1-1v-6zM16 13a1 1 0 011-1h2a1 1 0 011 1v6a1 1 0 01-1 1h-2a1 1 0 01-1-1v-6z"></path>
    </svg>
  `
}

const tabs = ref([
  { id: 'dashboard', name: 'داشبورد', icon: DashboardIcon },
  { id: 'sms-management', name: 'مدیریت پیامک', icon: SMSIcon },
  { id: 'order-list', name: 'لیست سفارشات', icon: OrderIcon },
  { id: 'response-management', name: 'مدیریت پاسخ‌ها', icon: ResponseIcon },
  { id: 'advanced-reports', name: 'گزارشات پیشرفته', icon: ReportsIcon },
  { id: 'advanced-analytics', name: 'تحلیل پیشرفته', icon: AnalyticsIcon },
  { id: 'sms-settings', name: 'تنظیمات پیامک', icon: SettingsIcon },
  { id: 'sms-gateway-status', name: 'وضعیت درگاه‌ها', icon: GatewayStatusIcon },
  { id: 'bulk-order-manager', name: 'مدیریت گروهی', icon: BulkIcon },
  { id: 'notifications-settings', name: 'تنظیمات اعلان', icon: NotificationIcon },
  { id: 'backup-export', name: 'پشتیبان‌گیری', icon: BackupIcon },
  { id: 'access-control', name: 'کنترل دسترسی', icon: SecurityIcon },
  { id: 'ui-personalization', name: 'شخصی‌سازی', icon: UIIcon },
  { id: 'survey-templates', name: 'تمپلیت‌های نظرسنجی', icon: TemplateIcon }
])
</script>

<style scoped>
.space-x-reverse > :not([hidden]) ~ :not([hidden]) {
  --tw-space-x-reverse: 1;
}
</style>
