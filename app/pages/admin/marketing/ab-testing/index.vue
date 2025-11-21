<template>
  <div class="p-6">
    <div class="mb-6">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-2xl font-bold text-gray-900">A/B تست</h1>
          <p class="text-gray-600">مدیریت و تحلیل تست‌های A/B برای بهبود عملکرد سایت</p>
        </div>
        <button class="px-4 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200 flex items-center" @click="openSettingsModal">
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
          </svg>
          تنظیمات
        </button>
      </div>
    </div>

    <!-- داشبورد -->
    <ABTestDashboard />

    <!-- فیلترها و جستجو -->
    <ABTestFilters 
      v-model="filters"
      @filter-change="handleFilterChange"
    />

    <!-- مدیریت تست‌ها -->
    <ABTestManagement 
      :filters="filters"
      @create-test="openFormModal"
      @edit-test="editTest"
      @view-results="viewResults"
      @duplicate-test="duplicateTest"
      @delete-test="deleteTest"
    />

    <!-- کنترل‌های مدیریتی -->
    <ABTestAccessControl 
      @permissions-changed="handlePermissionsChanged"
      @limits-changed="handleLimitsChanged"
    />

    <!-- یکپارچه‌سازی -->
    <ABTestIntegrations 
      @integration-changed="handleIntegrationChanged"
    />

    <!-- گزارش‌گیری -->
    <ABTestReports />

    <!-- Modal فرم ایجاد/ویرایش -->
    <ABTestFormModal 
      :is-open="isFormModalOpen"
      :editing-test="editingTest"
      @close="closeFormModal"
      @submit="handleFormSubmit"
    />

    <!-- Modal نتایج -->
    <ABTestResultsModal 
      :is-open="isResultsModalOpen"
      :test-id="selectedTestId"
      @close="closeResultsModal"
    />

    <!-- Modal تنظیمات -->
    <ABTestSettingsModal 
      :is-open="isSettingsModalOpen"
      @close="closeSettingsModal"
      @save="handleSettingsSave"
    />

    <!-- اعلان‌ها -->
    <ABTestNotifications ref="notificationsRef" />
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
</script>

<script setup lang="ts">
import { ref } from 'vue'
import { useAuth } from '~/composables/useAuth'

// Import components
import ABTestAccessControl from '~/pages/admin/marketing/ab-testing/components/ABTestAccessControl.vue'
import ABTestDashboard from '~/pages/admin/marketing/ab-testing/components/ABTestDashboard.vue'
import ABTestFilters from '~/pages/admin/marketing/ab-testing/components/ABTestFilters.vue'
import ABTestFormModal from '~/pages/admin/marketing/ab-testing/components/ABTestFormModal.vue'
import ABTestIntegrations from '~/pages/admin/marketing/ab-testing/components/ABTestIntegrations.vue'
import ABTestManagement from '~/pages/admin/marketing/ab-testing/components/ABTestManagement.vue'
import ABTestNotifications from '~/pages/admin/marketing/ab-testing/components/ABTestNotifications.vue'
import ABTestReports from '~/pages/admin/marketing/ab-testing/components/ABTestReports.vue'
import ABTestResultsModal from '~/pages/admin/marketing/ab-testing/components/ABTestResultsModal.vue'
import ABTestSettingsModal from '~/pages/admin/marketing/ab-testing/components/ABTestSettingsModal.vue'

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user: _user, hasPermission: _hasPermission } = useAuth()

// State management
const isFormModalOpen = ref(false)
const isResultsModalOpen = ref(false)
const isSettingsModalOpen = ref(false)
const editingTest = ref(null)
const selectedTestId = ref(null)
const notificationsRef = ref()
const filters = ref({})

// باز کردن modal فرم
const openFormModal = () => {
  editingTest.value = null
  isFormModalOpen.value = true
}

// بستن modal فرم
const closeFormModal = () => {
  isFormModalOpen.value = false
  editingTest.value = null
}

// باز کردن modal تنظیمات
const openSettingsModal = () => {
  isSettingsModalOpen.value = true
}

// ویرایش تست
const editTest = (testId: number) => {
  // در حالت واقعی، تست را از API دریافت می‌کنیم
  editingTest.value = { id: testId, name: 'تست نمونه' }
  isFormModalOpen.value = true
}

// مشاهده نتایج
const viewResults = (testId: number) => {
  selectedTestId.value = testId
  isResultsModalOpen.value = true
}

// کپی کردن تست
const duplicateTest = (testId: number) => {
  // منطق کپی کردن تست

  notificationsRef.value?.addNotification({
    type: 'completion',
    title: 'تست کپی شد',
    message: 'تست با موفقیت کپی شد و آماده ویرایش است.'
  })
}

// حذف تست
const deleteTest = (testId: number) => {
  if (confirm('آیا از حذف این تست اطمینان دارید؟')) {
    // منطق حذف تست

    notificationsRef.value?.addNotification({
      type: 'completion',
      title: 'تست حذف شد',
      message: 'تست با موفقیت حذف شد.'
    })
  }
}

interface ABTestFormData {
  name?: string
  description?: string
  [key: string]: unknown
}

interface ABTestSettings {
  [key: string]: unknown
}

// ارسال فرم
const handleFormSubmit = (formData: ABTestFormData) => {

  closeFormModal()
  notificationsRef.value?.addNotification({
    type: 'completion',
    title: 'تست ذخیره شد',
    message: 'تست با موفقیت ذخیره شد.'
  })
}

// بستن modal نتایج
const closeResultsModal = () => {
  isResultsModalOpen.value = false
  selectedTestId.value = null
}

// بستن modal تنظیمات
const closeSettingsModal = () => {
  isSettingsModalOpen.value = false
}

// ذخیره تنظیمات
const handleSettingsSave = (settings: ABTestSettings) => {

  closeSettingsModal()
  notificationsRef.value?.addNotification({
    type: 'completion',
    title: 'تنظیمات ذخیره شد',
    message: 'تنظیمات با موفقیت ذخیره شد.'
  })
}

interface Filters {
  [key: string]: unknown
}

interface Permissions {
  [key: string]: unknown
}

interface Limits {
  [key: string]: unknown
}

interface IntegrationData {
  [key: string]: unknown
}

// مدیریت فیلترها
const handleFilterChange = (newFilters: Filters) => {

  filters.value = newFilters
}

// مدیریت مجوزها
const handlePermissionsChanged = (permissions: Permissions) => {

  notificationsRef.value?.addNotification({
    type: 'completion',
    title: 'مجوزها ذخیره شد',
    message: 'مجوزهای کاربران با موفقیت ذخیره شد.'
  })
}

// مدیریت محدودیت‌ها
const handleLimitsChanged = (limits: Limits) => {

  notificationsRef.value?.addNotification({
    type: 'completion',
    title: 'محدودیت‌ها ذخیره شد',
    message: 'محدودیت‌های سیستم با موفقیت ذخیره شد.'
  })
}

// مدیریت یکپارچه‌سازی
const handleIntegrationChanged = (integration: string, data: IntegrationData) => {

  notificationsRef.value?.addNotification({
    type: 'completion',
    title: 'یکپارچه‌سازی ذخیره شد',
    message: `تنظیمات ${integration} با موفقیت ذخیره شد.`
  })
}
</script> 

